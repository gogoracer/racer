package parts

import (
	"github.com/gogoracer/racer/pkg/engine"
)

// ComponentGetNodes add a custom GetNodes function to ComponentMountable
type ComponentGetNodes struct {
	*engine.ComponentMountable

	GetNodesFunc func() *engine.NodeGroup
}

// CGN is a shortcut for NewComponentGetNodes
func CGN(name string, getNodesFunc func() *engine.NodeGroup, elements ...any) *ComponentGetNodes {
	return NewComponentGetNodes(name, getNodesFunc, elements...)
}

func NewComponentGetNodes(name string, getNodesFunc func() *engine.NodeGroup, elements ...any) *ComponentGetNodes {
	return &ComponentGetNodes{
		ComponentMountable: engine.NewComponentMountable(name, elements...),
		GetNodesFunc:       getNodesFunc,
	}
}

func (c *ComponentGetNodes) GetNodes() *engine.NodeGroup {
	return c.GetNodesFunc()
}
