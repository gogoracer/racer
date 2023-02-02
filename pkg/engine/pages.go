package engine

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/cornelk/hashmap"
	"github.com/gogoracer/racer/pkg/gas"
	"github.com/gogoracer/racer/pkg/headlamp"
	"github.com/rs/zerolog"
	"github.com/valyala/bytebufferpool"
)

const EventBindingsCacheDefault = 10 // Default for a small page

type PageOption func(*Page)

func PageOptionCache(cache Cache) func(*Page) {
	return func(p *Page) {
		p.cache = cache
	}
}

func PageOptionEventBindingCache(m *hashmap.Map[string, *EventBinding]) func(*Page) {
	return func(page *Page) {
		page.eventBindings = m
	}
}

type Page struct {
	// HTML renderer
	renderer *Renderer
	// DOM differ
	differ *Differ
	// Page rendering pipelines
	pipelineDiff *Pipeline
	// Page HTML rendering pipeline
	pipelineSSR *Pipeline
	// Internal debug logger
	logger zerolog.Logger
	// Virtual DOM
	dom DOM
	// What we think is the browser DOM is
	domBrowser any
	// Lock the page for writes
	mu sync.RWMutex
	// sessID is the WebSocket connection session id
	// only one page will have this at a time but is can be passed from page to page if connection is kept open
	sessID uint64
	// Component caches, to prevent walking to tree to find something
	eventBindings *hashmap.Map[string, *EventBinding]
	// Channel of outbound messages.
	send chan<- *gas.ToClient
	// Channel of inbound messages.
	receive <-chan *gas.FromClient
	// cache async safe
	cache Cache
	//
	// Hooks
	//
	// Before each event
	hookBeforeEvent []func(ctx context.Context, e Event) (context.Context, Event)
	// After each event
	hookAfterEvent []func(ctx context.Context, e Event) (context.Context, Event)
	// After each render
	hookAfterRender []func(context.Context, []*gas.Diff, chan<- *gas.ToClient)
	hookClose       []func(context.Context, *Page)
	// Before we do the initial render and send to the browser
	hookBeforeMount []func(context.Context, *Page)
	// After we do the initial render and send to the browser
	hookMount []func(context.Context, *Page)
	// When we close the page
	hookUnmount []func(context.Context, *Page)
}

var frontendShim *Tag

func headlampShim() *Tag {
	if frontendShim == nil {
		dev, err := headlamp.DistFS.ReadFile("dist/gogoracer-headlamp-lib-dev.es.js")
		if err != nil {
			panic(err)
		}

		frontendShim = NewTag(
			"script",
			Attrs{"type": "module"},
			HTML(dev),
		)
	}
	return frontendShim
}

func NewPage(options ...PageOption) *Page {
	p := &Page{
		dom:    NewDOM(),
		logger: zerolog.Nop(),
	}

	for i := 0; i < len(options); i++ {
		options[i](p)
	}

	if p.eventBindings == nil {
		p.eventBindings = hashmap.NewSized[string, *EventBinding](EventBindingsCacheDefault)
	}

	if p.renderer == nil {
		p.renderer = NewRenderer()
	}

	if p.differ == nil {
		p.differ = NewDiffer()
		p.dom.head.Add(headlampShim())
	}

	// Differ Pipeline
	if p.pipelineDiff == nil {
		p.pipelineDiff = NewPipeline(
			PipelineProcessorAttributePluginMount(p),
			PipelineProcessorMount(),
			PipelineProcessorEventBindingCache(p.eventBindings),
			PipelineProcessorUnmount(p),
			PipelineProcessorConvertToString(),
		)
	}
	// Server Side Render Pipeline
	if p.pipelineSSR == nil {
		p.pipelineSSR = NewPipeline(
			PipelineProcessorAttributePluginMountSSR(p),
			PipelineProcessorConvertToString(),
		)
	}

	if p.cache != nil {
		p.pipelineSSR.Add(PipelineProcessorRenderHashAndCache(p.logger, p.renderer, p.cache))
		p.DOM().HTML().Add(Attrs{PageHashAttr: PageHashAttrTmpl})
	} else {
		p.pipelineSSR.Add(PipelineProcessorRenderer(p.renderer))
	}

	return p
}

// type fromClientInfo struct {
// 	fromClient *gas.FromClient

// 	// Typ        string            `json:"t"`
// 	// ID         string            `json:"i,omitempty"`
// 	// Data       map[string]string `json:"d,omitempty"`
// 	File *File `json:"file,omitempty"`
// 	// ValueMulti []string          `json:"vm,omitempty"`
// 	// Selected   bool              `json:"s,omitempty"`
// 	// Extra      map[string]string `json:"e,omitempty"`
// }

func (p *Page) SetLogger(logger zerolog.Logger) {
	p.logger = logger
	p.renderer.SetLogger(p.logger)
	p.differ.SetLogger(p.logger)
}

func (p *Page) GetSessionID() uint64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.sessID
}

func (p *Page) Close(ctx context.Context) {
	for i := 0; i < len(p.hookClose); i++ {
		p.hookClose[i](ctx, p)
	}
}

func (p *Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.mu.Lock()
	if err := p.serverHTTP(w, r); err != nil {
		p.logger.Err(err).Msg("server http")
	}
	p.mu.Unlock()
}

func (p *Page) serverHTTP(w http.ResponseWriter, r *http.Request) error {
	_, err := p.runRenderPipeline(r.Context(), w)

	return err
}

func (p *Page) RunRenderPipeline(ctx context.Context, w io.Writer) (*NodeGroup, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.runRenderPipeline(ctx, w)
}

func (p *Page) runRenderPipeline(ctx context.Context, w io.Writer) (*NodeGroup, error) {
start:
	tree, err := p.pipelineSSR.run(ctx, w, p.getNodes())

	// A plugin being added will do this
	if errors.Is(err, ErrDOMInvalidated) {
		goto start
	}

	return tree, err
}

func (p *Page) RunDiffPipeline(ctx context.Context, w io.Writer) (*NodeGroup, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.runDiffPipeline(ctx, w)
}

func (p *Page) runDiffPipeline(ctx context.Context, w io.Writer) (*NodeGroup, error) {
start:
	tree, err := p.pipelineDiff.run(ctx, w, p.getNodes())

	// Plugins will do this
	if errors.Is(err, ErrDOMInvalidated) {
		goto start
	}

	return tree, err
}

func (p *Page) ServeWS(ctx context.Context, sessID uint64, send chan<- *gas.ToClient, receive <-chan *gas.FromClient) error {
	var err error
	p.mu.Lock()

	if p.send == nil {
		p.send = send
	}

	if p.receive == nil {
		p.receive = receive
	}

	// Send session id, it may be new or have changed
	p.sessID = sessID

	p.wsSend(ctx, &gas.ToClient{
		Message: &gas.ToClient_SessionInfo{
			SessionInfo: &gas.SessionInfo{
				Id: sessID,
			},
		},
	})

	// Do an initial render
	if p.domBrowser == nil {
		p.logger.Debug().Msg("ServeWS: browser render")
		// We need a static render
		p.domBrowser, err = p.runRenderPipeline(ctx, io.Discard)
		if err != nil {
			p.logger.Err(err).Msg("ServeWS: render pipeline")
		}
	}

	// Add render function to context
	ctx = context.WithValue(ctx, CtxRenderKey, p.executeRenderWS)
	ctx = context.WithValue(ctx, CtxRenderComponentKey, p.renderComponentWS)

	p.mu.Unlock()

	// TODO: add tests
	for i := 0; i < len(p.hookBeforeMount); i++ {
		p.hookBeforeMount[i](ctx, p)
	}

	// Do a dynamic render
	p.executeRenderWS(ctx)

	// TODO: add tests
	for i := 0; i < len(p.hookMount); i++ {
		p.hookMount[i](ctx, p)
	}

	defer func() {
		for i := 0; i < len(p.hookUnmount); i++ {
			p.hookUnmount[i](ctx, p)
		}
	}()

	taskQueue := make(chan func())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case task := <-taskQueue:
				task()
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case fromClient, ok := <-p.receive:
			if !ok {
				return nil
			}

			// We can't block here else we can't close and events here can trigger a close
			go func() {
				f := func() {

					p.logger.Trace().Str("msg", fromClient.String()).Msg("ws msg recv")

					switch fromClient.Type {
					// log
					case gas.FromClient_LOG:
						p.logger.Info().Str("log", fromClient.Data["m"]).Uint64("sess", sessID).Msg("ws log")
					// Event
					case gas.FromClient_EVENT:
						// Call handler
						go p.processMsgEvent(ctx, fromClient)
					default:
						p.logger.Error().Str("msg", fromClient.String()).Msg("ws msg recv: unexpected message format")
					}
				}

				select {
				case <-ctx.Done():
					return
				case taskQueue <- f:
				}
			}()
		}
	}
}

func (p *Page) executeRenderWS(ctx context.Context) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Do a dynamic render
	diffs, err := p.renderWS(ctx)
	if err != nil {
		p.logger.Err(err).Msg("ws render")
	}
	// Any DOM updates?

	if len(diffs) > 0 {
		toClient := &gas.ToClient{
			Message: &gas.ToClient_Diffs{
				Diffs: &gas.Diffs{
					Values: diffs,
				},
			},
		}
		p.wsSend(ctx, toClient)
	}

	for i := 0; i < len(p.hookAfterRender); i++ {
		p.hookAfterRender[i](ctx, diffs, p.send)
	}
}

func (p *Page) wsSend(ctx context.Context, toClient *gas.ToClient) {
	p.logger.Trace().Str("msg", toClient.String()).Msg("ws send")

	select {
	case <-ctx.Done():
	case p.send <- toClient:
	}
}

// Create and deletes should only happen at the end of a tag or attr list?
func (p *Page) diffInfosToProtobuf(diffInfos []*diffInfo) []*gas.Diff {
	diffValues := make([]*gas.Diff, len(diffInfos))

	// Reverse order to help with offset changes when adding new nodes
	for i := 0; i < len(diffInfos); i++ {
		diffInfo := diffInfos[i]

		diff := &gas.Diff{
			DiffType: diffInfo.Type,
		}

		if diffInfo.Root == "doc" {
			diff.Root = &gas.Diff_Document{}
		} else {
			diff.Root = &gas.Diff_ElementSelector{
				ElementSelector: diffInfo.Root,
			}
		}
		if cap(diff.PathIndicies) < len(diffInfo.PathIndicies) {
			diff.PathIndicies = make([]uint32, 0, len(diffInfo.PathIndicies))
		}
		for _, index := range diffInfo.PathIndicies {
			diff.PathIndicies = append(diff.PathIndicies, uint32(index))
		}

		p.updateDiffContents(diffInfo, diff)
		diffValues[i] = diff
	}

	return diffValues
}

func (p *Page) updateDiffContents(diffInfo *diffInfo, diff *gas.Diff) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	var el any
	if diffInfo.Type == gas.DiffType_DELETE && diffInfo.Attribute == nil {
		diff.ContentType = gas.ContentType_EMPTY_FROM_REMOVAL
	} else if diffInfo.Text != nil {
		// We use node.textContent in the browser which doesn't require us to encode
		diff.ContentType = gas.ContentType_TEXT
		el = HTML(*diffInfo.Text)
	} else if diffInfo.HTML != nil {
		diff.ContentType = gas.ContentType_HTML
		el = *diffInfo.HTML
	} else if diffInfo.Attribute != nil {
		diff.ContentType = gas.ContentType_ATTRIBUTE
		if err := p.renderer.Attribute([]Attributer{diffInfo.Attribute}, buf); err != nil {
			p.logger.Err(err).Msg("diffs to msg: render attribute")
		}
	} else if diffInfo.Tag != nil {
		diff.ContentType = gas.ContentType_HTML
		el = diffInfo.Tag
	}

	if err := p.renderer.HTML(buf, el); err != nil {
		p.logger.Err(err).Msg("diffs to msg: render children")
	}
	diff.Contents = string(buf.Bytes())

}

func (p *Page) processMsgEvent(ctx context.Context, msg *gas.FromClient) {
	keyCode, _ := strconv.Atoi(msg.Data["keyCode"])
	charCode, _ := strconv.Atoi(msg.Data["charCode"])
	shiftKey, _ := strconv.ParseBool(msg.Data["shiftKey"])
	altKey, _ := strconv.ParseBool(msg.Data["altKey"])
	ctrlKey, _ := strconv.ParseBool(msg.Data["ctrlKey"])
	isInitial, _ := strconv.ParseBool(msg.Data["init"])

	e := Event{
		Value:     msg.Data["value"],
		Values:    msg.ValueMulti,
		IsInitial: isInitial,
		Key:       msg.Data["key"],
		KeyCode:   keyCode,
		CharCode:  charCode,
		ShiftKey:  shiftKey,
		AltKey:    altKey,
		CtrlKey:   ctrlKey,
		File:      msg.File,
		Selected:  msg.Selected,
		Extra:     msg.Extra,
	}

	for _, id := range msg.Ids {
		p.logger.Trace().Str("id", id).Msg("call event handler")

		binding, _ := p.eventBindings.Get(id)
		if binding == nil {
			p.logger.Error().Str("id", id).Msg("unable to find binding")

			continue
		}

		e.Binding = binding

		if binding.Handler == nil {
			p.logger.Error().Str("id", id).Msg("binding handler nil")

			p.eventBindings.Del(id)

			break
		}

		// Hook
		for _, hook := range p.hookBeforeEvent {
			ctx, e = hook(ctx, e)
		}

		e.Binding.Handler(ctx, e)

		// Hook
		for _, hook := range p.hookAfterEvent {
			ctx, e = hook(ctx, e)
		}

		// Once, do this after calling the handler so the developer can change their mind
		if e.Binding.Once {
			p.eventBindings.Del(id)
			binding.Component.RemoveEventBinding(id)
		}

		// Auto Render?
		if e.Binding.Component.IsAutoRender() {
			p.executeRenderWS(ctx)
		}
	}
}

func (p *Page) renderWS(ctx context.Context) ([]*gas.Diff, error) {
	// TODO: replace discard with something useful
	tree, err := p.runDiffPipeline(ctx, io.Discard)
	if err != nil {
		return nil, fmt.Errorf("run pipeline: %w", err)
	}

	diffInfos, err := p.differ.Trees("doc", nil, p.domBrowser, tree)
	if err != nil {
		return nil, fmt.Errorf("diff old and new tag trees: %w", err)
	}
	p.domBrowser = tree

	return p.diffInfosToProtobuf(diffInfos), nil
}

func (p *Page) GetNodes() *NodeGroup {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.getNodes()
}

func (p *Page) getNodes() *NodeGroup {
	return G(p.dom.docType, p.dom.html)
}

// Gets a lock as it acts like a public function
// Will fail if plugins invalidate the tree
func (p *Page) renderComponentWS(ctx context.Context, comp Componenter) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	oldTag := p.findComponentInTree(comp.GetID())
	if oldTag == nil {
		p.logger.Error().Str("id", comp.GetID()).Str("name", comp.GetName()).
			Msg("render component ws: can't find component in tree")

		return
	}

	// TODO: replace discard
	newTreeNode, err := p.pipelineDiff.runNode(ctx, io.Discard, comp)
	if err != nil {
		p.logger.Error().Str("id", comp.GetID()).Str("name", comp.GetName()).
			Msg("render component ws: pipeline run node")

		return
	}

	var newTag *Tag

	newTags, ok := newTreeNode.(*NodeGroup)
	if ok && len(newTags.Get()) != 0 {
		newTag, ok = newTags.Get()[0].(*Tag)
	} else {
		newTag, ok = newTreeNode.(*Tag)
	}

	if !ok {
		p.logger.Error().Str("id", comp.GetID()).Str("name", comp.GetName()).
			Msg("render component ws: new node is not a tag, component, or group")

		return
	}

	diffInfos, err := p.differ.Trees(comp.GetID(), nil, oldTag, newTag)
	if err != nil {
		p.logger.Err(err).Str("id", comp.GetID()).Str("name", comp.GetName()).Msg("diff trees")

		return
	}

	if len(diffInfos) > 0 {
		toClient := &gas.ToClient{
			Message: &gas.ToClient_Diffs{
				Diffs: &gas.Diffs{
					Values: p.diffInfosToProtobuf(diffInfos),
				},
			},
		}
		p.wsSend(ctx, toClient)

		oldTag.name = newTag.name
		oldTag.void = newTag.void
		oldTag.attributes = newTag.attributes
		oldTag.nodes = newTag.nodes
	}
}

func (p *Page) GetBrowserNodeByID(id string) *Tag {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.findComponent(id, p.domBrowser)
}

func (p *Page) findComponentInTree(id string) *Tag {
	return p.findComponent(id, p.domBrowser)
}

func (p *Page) findComponent(id string, tree any) *Tag {
	switch v := tree.(type) {
	case *NodeGroup:
		g := v.Get()
		for i := 0; i < len(g); i++ {
			c := p.findComponent(id, g[i])
			if c != nil {
				return c
			}
		}
	case *Tag:
		if v.GetAttributeValue(AttrID) == id {
			return v
		}

		return p.findComponent(id, v.nodes)
	}

	return nil
}

func (p *Page) DOM() DOM {
	return p.dom
}

func (p *Page) PipelineDiff() *Pipeline {
	return p.pipelineDiff
}

func (p *Page) PipelineSSR() *Pipeline {
	return p.pipelineSSR
}

func (p *Page) HookBeforeEventAdd(hook func(context.Context, Event) (context.Context, Event)) {
	p.hookBeforeEvent = append(p.hookBeforeEvent, hook)
}

func (p *Page) HookAfterRenderAdd(hook func(ctx context.Context, diffs []*gas.Diff, send chan<- *gas.ToClient)) {
	p.hookAfterRender = append(p.hookAfterRender, hook)
}

func (p *Page) HookCloseAdd(hook func(context.Context, *Page)) {
	p.hookClose = append(p.hookClose, hook)
}

func (p *Page) HookMountAdd(hook func(context.Context, *Page)) {
	p.hookMount = append(p.hookMount, hook)
}

func (p *Page) HookUnmountAdd(hook func(context.Context, *Page)) {
	p.hookUnmount = append(p.hookUnmount, hook)
}

func (p *Page) HookBeforeMountAdd(hook func(context.Context, *Page)) {
	p.hookBeforeMount = append(p.hookBeforeMount, hook)
}

func (p *Page) DOMBrowser() any {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.domBrowser
}

func (p *Page) SetDOMBrowser(dom any) {
	p.mu.Lock()
	p.domBrowser = dom
	p.mu.Unlock()
}
