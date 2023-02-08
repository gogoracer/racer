package frame

import (
	"github.com/cornelk/hashmap"
	"github.com/dgraph-io/ristretto"
	"github.com/gogoracer/racer/pkg/engine"
	"github.com/gogoracer/racer/pkg/headlamp"
	"github.com/gogoracer/racer/pkg/parts"
	"github.com/rs/zerolog/log"
)

var frontendShim *engine.UberElement

func PageOptionCache(cache Cache) func(*Page) {
	return func(p *Page) {
		p.cache = cache
	}
}

func headlampShim() *engine.UberElement {
	if frontendShim == nil {
		dev, err := headlamp.DistFS.ReadFile("dist/gogoracer-headlamp-lib-dev.es.js")
		if err != nil {
			panic(err)
		}

		frontendShim = engine.Uber("script").Attr(engine.NewAttribute("type", "module")).HTML(string(dev))
	}
	return frontendShim
}

type Page struct {
	page *engine.Page

	// Virtual DOM
	dom DOM
	// cache async safe
	cache         Cache
	eventBindings *hashmap.Map[string, *engine.EventBinding]
}

func (p *Page) DOM() DOM {
	return p.dom
}

func (p *Page) GetCache() engine.Cache {
	return p.cache
}

func NewPage() *Page {
	p := &Page{
		eventBindings: hashmap.NewSized[string, *engine.EventBinding](engine.EventBindingsCacheDefault),
	}

	renderer := engine.NewRenderer()

	p.page = engine.NewPage(
		engine.PageOptionDOMFunc(
			func() *engine.NodeGroup {
				return engine.G(p.dom.docType, p.dom.html)
			}),
		engine.PageOptionRenderer(renderer),
		engine.PageOptionEventBindingCache(p.eventBindings),
	)

	p.dom = NewDOM()
	p.dom.head.Element(headlampShim())

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	p.cache = parts.NewCacheRistretto(cache)

	// Differ Pipeline
	p.page.PipelineDiff().RemoveAll()
	p.page.PipelineDiff().Add(
		engine.PipelineProcessorAttributePluginMount(p),
		engine.PipelineProcessorMount(),
		engine.PipelineProcessorEventBindingCache(p.eventBindings),
		engine.PipelineProcessorUnmount(p),
		engine.PipelineProcessorConvertToString())

	// Server Side Render Pipeline
	p.page.PipelineSSR().RemoveAll()
	p.page.PipelineSSR().Add(
		engine.PipelineProcessorAttributePluginMountSSR(p),
		engine.PipelineProcessorConvertToString(),
		PipelineProcessorRenderHashAndCache(log.Logger, renderer, p.cache),
	)

	return p
}

func (p *Page) GetPage() *engine.Page {
	return p.page
}
