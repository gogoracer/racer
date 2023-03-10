package engine

import (
	"context"
	"io"
	"sync"

	"github.com/teris-io/shortid"
)

type QueueMessage struct {
	Topic string
	Value any
}

type QueueSubscriber interface {
	GetID() string
	OnMessage(message QueueMessage)
}

type PubSubMounter interface {
	GetID() string
	PubSubMount(context.Context, *PubSub)
}

type PubSubAfterMounter interface {
	PubSubMounter
	AfterPubSubMount(context.Context, *PubSub)
}

type PubSubSSRMounter interface {
	GetID() string
	PubSubSSRMount(context.Context, *PubSub)
}

type PubSub struct {
	mu          sync.RWMutex
	subscribers map[string][]QueueSubscriber
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: map[string][]QueueSubscriber{},
	}
}

func (ps *PubSub) Subscribe(sub QueueSubscriber, topics ...string) {
	if len(topics) == 0 {
		// TODO: fix
		// engine.LoggerDev.Warn().Str("callers", engine.CallerStackStr()).Msg("no topics passed")

		return
	}

	if sub == nil {
		// TODO: fix
		// engine.LoggerDev.Warn().Str("callers", engine.CallerStackStr()).Msg("sub nil")

		return
	}

	// A Publish can trigger a Subscribe. Subscribe will be added after the Publish
	go func() {
		ps.mu.Lock()
		defer ps.mu.Unlock()

		for i := 0; i < len(topics); i++ {
			ps.subscribers[topics[i]] = append(ps.subscribers[topics[i]], sub)
		}
	}()
}

func (ps *PubSub) SubscribeWait(sub QueueSubscriber, topics ...string) {
	if len(topics) == 0 {
		// TODO: fix
		// engine.LoggerDev.Warn().Str("callers", engine.CallerStackStr()).Msg("no topics passed")

		return
	}

	if sub == nil {
		// TODO: fix
		// engine.LoggerDev.Warn().Str("callers", engine.CallerStackStr()).Msg("sub nil")

		return
	}

	ps.mu.Lock()
	defer ps.mu.Unlock()

	for i := 0; i < len(topics); i++ {
		ps.subscribers[topics[i]] = append(ps.subscribers[topics[i]], sub)
	}
}

func (ps *PubSub) SubscribeWaitFunc(subFunc func(message QueueMessage), topics ...string) SubscribeFunc {
	sub := NewSub(subFunc)

	ps.SubscribeWait(sub, topics...)

	return sub
}

func (ps *PubSub) SubscribeFunc(subFunc func(message QueueMessage), topics ...string) SubscribeFunc {
	sub := NewSub(subFunc)

	ps.Subscribe(sub, topics...)

	return sub
}

func (ps *PubSub) UnsubscribeWait(sub QueueSubscriber, topics ...string) {
	if len(topics) == 0 {
		panic("no topics passed")
	}

	if sub == nil {
		// TODO: fix
		// engine.LoggerDev.Warn().Str("callers", engine.CallerStackStr()).Msg("sub when nil")

		return
	}

	ps.mu.Lock()
	defer ps.mu.Unlock()

	for i := 0; i < len(topics); i++ {
		var newList []QueueSubscriber

		for j := 0; j < len(ps.subscribers[topics[i]]); j++ {
			if ps.subscribers[topics[i]][j].GetID() == sub.GetID() {
				continue
			}

			newList = append(newList, ps.subscribers[topics[i]][j])
		}

		ps.subscribers[topics[i]] = newList
	}
}

func (ps *PubSub) Unsubscribe(sub QueueSubscriber, topics ...string) {
	if len(topics) == 0 {
		panic("no topics passed")
	}

	if sub == nil {
		// TODO: fix
		// engine.LoggerDev.Warn().Str("callers", engine.CallerStackStr()).Msg("sub when nil")

		return
	}

	go func() {
		ps.mu.Lock()
		defer ps.mu.Unlock()

		for i := 0; i < len(topics); i++ {
			var newList []QueueSubscriber

			for j := 0; j < len(ps.subscribers[topics[i]]); j++ {
				if ps.subscribers[topics[i]][j].GetID() == sub.GetID() {
					continue
				}

				newList = append(newList, ps.subscribers[topics[i]][j])
			}

			ps.subscribers[topics[i]] = newList
		}
	}()
}

func (ps *PubSub) Publish(topic string, value any) {
	// Multiple Publish calls can run concurrently
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	item := QueueMessage{topic, value}
	for i := 0; i < len(ps.subscribers[topic]); i++ {
		ps.subscribers[topic][i].OnMessage(item)
	}
}

type SubscribeFunc struct {
	fn func(message QueueMessage)
	id string
}

func (s SubscribeFunc) OnMessage(message QueueMessage) {
	s.fn(message)
}

func (s SubscribeFunc) GetID() string {
	return s.id
}

func NewSub(onMessageFn func(message QueueMessage)) SubscribeFunc {
	return SubscribeFunc{onMessageFn, shortid.MustGenerate()}
}

const PipelineProcessorKeyPubSubMount = "hlivekit_ps_mount"

func (a *PubSubAttribute) PipelineProcessorPubSub() *PipelineProcessor {
	pp := NewPipelineProcessor(PipelineProcessorKeyPubSubMount)

	pp.BeforeTagger = func(ctx context.Context, w io.Writer, tag Tagger) (Tagger, error) {
		if comp, ok := tag.(PubSubMounter); ok {
			a.mu.Lock()
			defer a.mu.Unlock()

			if _, exists := a.mountedMap[comp.GetID()]; !exists {
				a.mountedMap[comp.GetID()] = struct{}{}
				comp.PubSubMount(ctx, a.pubSub)

				if afterComp, ok := comp.(PubSubAfterMounter); ok {
					a.afterMountQueue = append(a.afterMountQueue, afterComp)
				}
			}

			// A way to remove the key when you delete a Component
			if comp, ok := tag.(Teardowner); ok {
				comp.AddTeardown(func() {
					a.mu.Lock()
					delete(a.mountedMap, comp.GetID())
					a.mu.Unlock()
				})
			}
		}

		return tag, nil
	}

	pp.AfterWalk = func(ctx context.Context, _ io.Writer, node *NodeGroup) (*NodeGroup, error) {
		a.mu.Lock()
		defer a.mu.Unlock()

		queue := append([]PubSubAfterMounter{}, a.afterMountQueue...)
		a.afterMountQueue = nil

		// People will want to be able to publish and render
		go func() {
			for i := 0; i < len(queue); i++ {
				queue[i].AfterPubSubMount(ctx, a.pubSub)
			}
		}()

		return node, nil
	}

	return pp
}

const PubSubAttributeName = "data-hlive-pubsub"

func InstallPubSub(pubSub *PubSub) Attributer {
	attr := &PubSubAttribute{
		Attribute:  NewAttribute(PubSubAttributeName, ""),
		pubSub:     pubSub,
		mountedMap: map[string]struct{}{},
	}

	return attr
}

type PubSubAttribute struct {
	*Attribute

	pubSub *PubSub
	// Will memory leak is you don't use a Teardowner when deleting Components
	mountedMap      map[string]struct{}
	afterMountQueue []PubSubAfterMounter
	mu              sync.Mutex
	rendered        bool
}

func (a *PubSubAttribute) Initialize(page *Page) {
	if a.rendered {
		return
	}

	page.PipelineDiff().Add(a.PipelineProcessorPubSub())
}

func (a *PubSubAttribute) InitializeSSR(page *Page) {
	a.rendered = true
	page.PipelineDiff().Add(a.PipelineProcessorPubSub())
}

// ComponentPubSub add PubSub to ComponentMountable
type ComponentPubSub struct {
	*ComponentMountable

	mountPubSubFunc      *LockBox[func(ctx context.Context, pubSub *PubSub)]
	afterMountPubSubFunc *LockBox[func(ctx context.Context, pubSub *PubSub)]
}

// CPS is a shortcut for NewComponentPubSub
func CPS(name string, elements ...any) *ComponentPubSub {
	return NewComponentPubSub(name, elements...)
}

func NewComponentPubSub(name string, elements ...any) *ComponentPubSub {
	return &ComponentPubSub{
		ComponentMountable:   NewComponentMountable(name, elements...),
		mountPubSubFunc:      NewLockBox[func(ctx context.Context, pubSub *PubSub)](nil),
		afterMountPubSubFunc: NewLockBox[func(ctx context.Context, pubSub *PubSub)](nil),
	}
}

func (c *ComponentPubSub) PubSubMount(ctx context.Context, pubSub *PubSub) {
	f := c.mountPubSubFunc.Get()
	if f == nil {
		return
	}

	f(ctx, pubSub)
}

func (c *ComponentPubSub) AfterPubSubMount(ctx context.Context, pubSub *PubSub) {
	f := c.afterMountPubSubFunc.Get()
	if f == nil {
		return
	}

	f(ctx, pubSub)
}

func (c *ComponentPubSub) SetMountPubSub(f func(ctx context.Context, pubSub *PubSub)) {
	c.mountPubSubFunc.Set(f)
}

func (c *ComponentPubSub) SetAfterMountPubSub(f func(ctx context.Context, pubSub *PubSub)) {
	c.afterMountPubSubFunc.Set(f)
}
