package engine

import (
	"context"
	"io"
	"strconv"
	"sync/atomic"

	"github.com/cornelk/hashmap"
)

const (
	PipelineProcessorKeyStripHLiveAttrs      = "hlive_strip_hlive_attr"
	PipelineProcessorKeyRenderer             = "hlive_renderer"
	PipelineProcessorKeyEventBindingCache    = "hlive_eb"
	PipelineProcessorKeyAttributePluginMount = "hlive_attr_mount"
	PipelineProcessorKeyMount                = "hlive_mount"
	PipelineProcessorKeyUnmount              = "hlive_unmount"
	PipelineProcessorKeyConvertToString      = "hlive_conv_str"
)

type PipelineProcessor struct {
	// Will replace an existing processor with the same key. An empty string won't error.
	Key             string
	Disabled        bool
	BeforeWalk      PipeNodeGroupHandler
	OnSimpleNode    PipeNodeHandler
	BeforeTagger    PipeTaggerHandler
	BeforeAttribute PipeAttributerHandler
	AfterAttribute  PipeAttributerHandler
	AfterTagger     PipeTagHandler
	AfterWalk       PipeNodeGroupHandler
}

func NewPipelineProcessor(key string) *PipelineProcessor {
	return &PipelineProcessor{Key: key}
}

func PipelineProcessorEventBindingCache(cache *hashmap.Map[string, *EventBinding]) *PipelineProcessor {
	pp := NewPipelineProcessor(PipelineProcessorKeyEventBindingCache)

	pp.BeforeTagger = func(ctx context.Context, w io.Writer, tag Tagger) (Tagger, error) {
		if comp, ok := tag.(Componenter); ok {
			bindings := comp.GetEventBindings()

			for i := 0; i < len(bindings); i++ {
				cache.Set(bindings[i].ID, bindings[i])
			}
		}

		return tag, nil
	}

	return pp
}

func PipelineProcessorMount() *PipelineProcessor {
	var compID uint64

	pp := NewPipelineProcessor(PipelineProcessorKeyMount)

	pp.BeforeTagger = func(ctx context.Context, w io.Writer, tag Tagger) (Tagger, error) {
		if comp, ok := tag.(UniqueTagger); ok && comp.GetID() == "" {
			comp.SetID(strconv.FormatUint(atomic.AddUint64(&compID, 1), 10))

			if comp, ok := tag.(Mounter); ok {
				comp.Mount(ctx)
			}
		}

		return tag, nil
	}

	return pp
}

func PipelineProcessorUnmount(page *Page) *PipelineProcessor {
	cache := hashmap.New[string, Unmounter]()

	page.hookClose = append(page.hookClose, func(ctx context.Context, page *Page) {
		cache.Range(func(key string, value Unmounter) bool {
			if value != nil {
				value.Unmount(ctx)
			}

			return true
		})
	})

	pp := NewPipelineProcessor(PipelineProcessorKeyUnmount)

	pp.BeforeTagger = func(ctx context.Context, w io.Writer, tag Tagger) (Tagger, error) {
		if comp, ok := tag.(Unmounter); ok {
			id := comp.GetID()

			if id == "" {
				return tag, nil
			}

			if cache.Insert(id, comp) {
				// A way to remove the key when you delete a Component
				if comp, ok := tag.(Teardowner); ok {
					comp.AddTeardown(func() {
						cache.Del(id)
					})
				}
			}
		}

		return tag, nil
	}

	return pp
}

func PipelineProcessorRenderer(renderer *Renderer) *PipelineProcessor {
	pp := NewPipelineProcessor(PipelineProcessorKeyRenderer)

	pp.AfterWalk = func(ctx context.Context, w io.Writer, node *NodeGroup) (*NodeGroup, error) {
		return node, renderer.HTML(w, node)
	}

	return pp
}

func PipelineProcessorConvertToString() *PipelineProcessor {
	pp := NewPipelineProcessor(PipelineProcessorKeyConvertToString)

	pp.OnSimpleNode = func(ctx context.Context, w io.Writer, node any) (any, error) {
		if node == nil {
			return nil, nil
		}

		parseString := func(s string) (any, error) {
			if len(s) == 0 {
				return nil, nil
			}
			return s, nil
		}

		parseInt64 := func(i int64) (any, error) {
			return strconv.FormatInt(i, base10), nil
		}

		parseUint64 := func(i uint64) (any, error) {
			return strconv.FormatUint(i, base10), nil
		}

		parseFloat64 := func(f float64) (any, error) {
			return strconv.FormatFloat(f, 'f', -1, bit64), nil
		}

		switch v := node.(type) {
		case nil:
			return nil, nil

		case string:
			return parseString(v)
		case *string:
			return parseString(*v)

		case int:
			return parseInt64(int64(v))
		case *int:
			return parseInt64(int64(*v))
		case int8:
			return parseInt64(int64(v))
		case *int8:
			return parseInt64(int64(*v))
		case int16:
			return parseInt64(int64(v))
		case *int16:
			return parseInt64(int64(*v))
		case int32:
			return parseInt64(int64(v))
		case *int32:
			return parseInt64(int64(*v))
		case int64:
			return parseInt64(v)
		case *int64:
			return parseInt64(*v)

		case uint:
			return parseUint64(uint64(v))
		case *uint:
			return parseUint64(uint64(*v))
		case uint8:
			return parseUint64(uint64(v))
		case *uint8:
			return parseUint64(uint64(*v))
		case uint16:
			return parseUint64(uint64(v))
		case *uint16:
			return parseUint64(uint64(*v))
		case uint32:
			return parseUint64(uint64(v))
		case *uint32:
			return parseUint64(uint64(*v))
		case uint64:
			return parseUint64(v)
		case *uint64:
			return parseUint64(*v)

		case float32:
			return parseFloat64(float64(v))
		case *float32:
			return parseFloat64(float64(*v))
		case float64:
			return parseFloat64(v)
		case *float64:
			return parseFloat64(*v)

		// HTML need to be a pointer to allow for msgpack to keep its type
		case HTML:
			return &v, nil
		default:
			return v, nil
		}
	}

	return pp
}

func PipelineProcessorAttributePluginMount(page *Page) *PipelineProcessor {
	cache := hashmap.New[string, *struct{}]()

	pp := NewPipelineProcessor(PipelineProcessorKeyAttributePluginMount)

	pp.BeforeAttribute = func(ctx context.Context, w io.Writer, attr Attributer) (Attributer, error) {
		var err error
		if ap, ok := attr.(AttributePluginer); ok {
			if set := cache.Insert(ap.GetName(), nil); set {
				ap.Initialize(page)

				err = ErrDOMInvalidated
			}
		}

		return attr, err
	}

	return pp
}

func PipelineProcessorAttributePluginMountSSR(page *Page) *PipelineProcessor {
	cache := hashmap.New[string, *struct{}]()

	pp := NewPipelineProcessor(PipelineProcessorKeyAttributePluginMount)

	pp.BeforeAttribute = func(ctx context.Context, w io.Writer, attr Attributer) (Attributer, error) {
		var err error
		if ap, ok := attr.(AttributePluginer); ok {
			if _, loaded := cache.GetOrInsert(ap.GetName(), nil); !loaded {
				ap.InitializeSSR(page)

				err = ErrDOMInvalidated
			}
		}

		return attr, err
	}

	return pp
}
