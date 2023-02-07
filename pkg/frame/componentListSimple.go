package frame

import (
	"sync"

	"github.com/gogoracer/racer/pkg/engine"
)

// ComponentListSimple is a version of ComponentList that doesn't have the Teardown logic
type ComponentListSimple struct {
	*engine.ComponentMountable

	items []engine.UniqueTagger
	mu    sync.RWMutex
}

// NewComponentListSimple creates a ComponentListSimple value
func NewComponentListSimple(name string, elements ...any) *ComponentListSimple {
	list := &ComponentListSimple{
		ComponentMountable: engine.CM(name),
	}

	list.Add(elements...)

	return list
}

// GetNodes returns the list of items.
func (list *ComponentListSimple) GetNodes() *engine.NodeGroup {
	list.mu.RLock()
	defer list.mu.RUnlock()

	return engine.G(list.items)
}

// Add an element to this ComponentListSimple.
//
// You can add Groups, UniqueTagger, EventBinding, or None Node Elements
func (list *ComponentListSimple) Add(elements ...any) {
	for i := 0; i < len(elements); i++ {
		switch v := elements[i].(type) {
		case *engine.NodeGroup:
			list.Add(v.Get()...)
		case engine.UniqueTagger:
			list.items = append(list.items, v)
		default:
			if engine.IsNonNodeElement(v) {
				list.Component.Add(v)
			} else {
				panic("invalid element type")
				// engine.LoggerDev.Warn().Str("callers", engine.CallerStackStr()).
				// 	Str("element", fmt.Sprintf("%#v", v)).
				// 	Msg("invalid element type")
			}
		}
	}
}

func (list *ComponentListSimple) AddItems(items ...engine.UniqueTagger) {
	list.mu.Lock()
	list.items = append(list.items, items...)
	list.mu.Unlock()
}

func (list *ComponentListSimple) RemoveItems(items ...engine.UniqueTagger) {
	list.mu.Lock()
	defer list.mu.Unlock()

	var newList []engine.UniqueTagger

	for i := 0; i < len(list.items); i++ {
		hit := false

		for j := 0; j < len(items); j++ {
			item := items[j]

			if item.GetID() == list.items[i].GetID() {
				hit = true

				break
			}
		}

		if !hit {
			newList = append(newList, list.items[i])
		}
	}

	list.items = newList
}

func (list *ComponentListSimple) RemoveAllItems() {
	list.mu.Lock()
	list.items = nil
	list.mu.Unlock()
}
