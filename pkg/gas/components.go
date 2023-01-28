package gas

import (
	"context"
	"strconv"
	"strings"
)

// Componenter builds on UniqueTagger and adds the ability to handle events.
type Componenter interface {
	UniqueTagger
	// GetEventBinding returns a binding by its id
	GetEventBinding(id string) *EventBinding
	// GetEventBindings returns all event bindings for this tag
	GetEventBindings() []*EventBinding
	// RemoveEventBinding remove an event binding from this component
	RemoveEventBinding(id string)
	// IsAutoRender indicates if the page should rerender after an event binding on this tag is called
	IsAutoRender() bool
}

// Component is the default implementation of Componenter.
type Component struct {
	*Tag

	AutoRender bool
	id         string
	bindingID  uint32
	bindings   []*EventBinding
}

// NewComponent is a constructor for Component.
//
// You can add zero or many Attributes and Tags.
func NewComponent(name string, elements ...any) *Component {
	c := &Component{
		Tag:        NewTag(name),
		AutoRender: true,
	}

	c.Add(elements...)

	return c
}

// W is a shortcut for Wrap.
//
// Wrap takes a Tag and creates a Component with it.
func W(tag *Tag, elements ...any) *Component {
	return Wrap(tag, elements...)
}

// Wrap takes a Tag and creates a Component with it.
func Wrap(tag *Tag, elements ...any) *Component {
	c := &Component{
		Tag:        tag,
		AutoRender: true,
	}

	c.Add(elements...)

	return c
}

// GetID returns this component's unique ID
func (c *Component) GetID() string {
	c.Tag.mu.RLock()
	defer c.Tag.mu.RUnlock()

	return c.id
}

// SetID component's unique ID
func (c *Component) SetID(id string) {
	c.Tag.mu.Lock()
	defer c.Tag.mu.Unlock()

	c.id = id
	c.Tag.addAttributes(NewAttribute(AttrID, id))

	if value := c.bindingAttrValue(); value != "" {
		c.Tag.addAttributes(NewAttribute(AttrOn, value))
	}
}

func (c *Component) bindingAttrValue() string {
	var value string
	for i := 0; i < len(c.bindings); i++ {
		if c.bindings[i].ID == "" {
			c.bindingID++
			c.bindings[i].ID = c.id + "-" + strconv.FormatUint(uint64(c.bindingID), 10)
		}

		value += c.bindings[i].ID + "|" + c.bindings[i].Name + ","
	}

	return strings.TrimRight(value, ",")
}

// IsAutoRender indicates if this component should trigger "Auto Render"
func (c *Component) IsAutoRender() bool {
	c.Tag.mu.RLock()
	defer c.Tag.mu.RUnlock()

	return c.AutoRender
}

// GetEventBinding will return an EventBinding that exists directly on this element, it doesn't check its children.
// Returns nil is not found.
func (c *Component) GetEventBinding(id string) *EventBinding {
	c.Tag.mu.RLock()
	defer c.Tag.mu.RUnlock()

	for i := 0; i < len(c.bindings); i++ {
		if c.bindings[i].ID == id {
			return c.bindings[i]
		}
	}

	return nil
}

// GetEventBindings returns all EventBindings for this component, not it's children.
func (c *Component) GetEventBindings() []*EventBinding {
	c.Tag.mu.RLock()
	defer c.Tag.mu.RUnlock()

	return append([]*EventBinding{}, c.bindings...)
}

// RemoveEventBinding removes an EventBinding that matches the passed ID.
//
// No error if the passed id doesn't match an EventBinding.
// It doesn't check its children.
func (c *Component) RemoveEventBinding(id string) {
	c.Tag.mu.Lock()
	defer c.Tag.mu.Unlock()

	var newList []*EventBinding
	for i := 0; i < len(c.bindings); i++ {
		if c.bindings[i].ID == id {
			continue
		}

		newList = append(newList, c.bindings[i])
	}
	c.bindings = newList

	// Reset attribute
	if value := c.bindingAttrValue(); value == "" {
		c.removeAttributes(AttrOn)
	} else {
		c.Tag.addAttributes(NewAttribute(AttrOn, value))
	}
}

// Add an element to this Component.
//
// This is an easy way to add anything.
func (c *Component) Add(elements ...any) {
	if c.IsNil() {
		return
	}

	for i := 0; i < len(elements); i++ {
		switch v := elements[i].(type) {
		// NoneNodeElements
		case []any:
			for j := 0; j < len(v); j++ {
				c.Add(v[j])
			}
		case *NodeGroup:
			if v == nil {
				continue
			}

			list := v.Get()
			for j := 0; j < len(list); j++ {
				c.Add(list[j])
			}
		case *ElementGroup:
			if v == nil {
				continue
			}

			list := v.Get()
			for j := 0; j < len(list); j++ {
				c.Add(list[j])
			}
		case *EventBinding:
			if v == nil {
				continue
			}

			c.Tag.mu.Lock()
			c.on(v)
			c.Tag.mu.Unlock()
		default:
			c.Tag.Add(v)
		}
	}
}

func (c *Component) on(binding *EventBinding) {
	binding.Component = c
	c.bindings = append(c.bindings, binding)

	// See Component.SetID
	if c.id == "" {
		return
	}

	if value := c.bindingAttrValue(); value != "" {
		c.Tag.addAttributes(NewAttribute(AttrOn, value))
	}
}

// Mounter wants to be notified after it's mounted.
type Mounter interface {
	UniqueTagger
	// Mount is called after a component is mounted
	Mount(ctx context.Context)
}

// Unmounter wants to be notified before it's unmounted.
type Unmounter interface {
	UniqueTagger
	// Unmount is called before a component is unmounted
	Unmount(ctx context.Context)
}

// Teardowner wants to have manual control when it needs to be removed from a Page.
// If you have a Mounter or Unmounter that will be permanently removed from a Page they must call the passed
// function to clean up their references.
type Teardowner interface {
	UniqueTagger
	// AddTeardown adds a teardown function
	AddTeardown(teardown func())
	// Teardown call the set teardown function passed in SetTeardown
	Teardown()
}

type ComponentMountable struct {
	*Component

	mountFunc   func(ctx context.Context)
	unmountFunc func(ctx context.Context)
	teardowns   []func()
}

// CM is a shortcut for NewComponentMountable
func CM(name string, elements ...any) *ComponentMountable {
	return NewComponentMountable(name, elements...)
}

func NewComponentMountable(name string, elements ...any) *ComponentMountable {
	return &ComponentMountable{
		Component: NewComponent(name, elements...),
	}
}

func (c *ComponentMountable) Mount(ctx context.Context) {
	if c == nil {
		return
	}

	c.Tag.mu.RLock()
	f := c.mountFunc
	c.Tag.mu.RUnlock()

	if f != nil {
		f(ctx)
	}
}

func (c *ComponentMountable) Unmount(ctx context.Context) {
	if c == nil {
		return
	}

	c.Tag.mu.RLock()
	f := c.unmountFunc
	c.Tag.mu.RUnlock()

	if c.unmountFunc != nil {
		f(ctx)
	}
}

func (c *ComponentMountable) SetMount(mount func(ctx context.Context)) {
	c.Tag.mu.Lock()
	c.mountFunc = mount
	c.Tag.mu.Unlock()
}

func (c *ComponentMountable) SetUnmount(unmount func(ctx context.Context)) {
	c.Tag.mu.Lock()
	c.unmountFunc = unmount
	c.Tag.mu.Unlock()
}

func (c *ComponentMountable) AddTeardown(teardown func()) {
	c.Tag.mu.Lock()
	c.teardowns = append(c.teardowns, teardown)
	c.Tag.mu.Unlock()
}

func (c *ComponentMountable) Teardown() {
	c.Tag.mu.RLock()
	teardowns := c.teardowns
	c.Tag.mu.RUnlock()

	for i := 0; i < len(teardowns); i++ {
		teardowns[i]()
	}
}

// WM is a shortcut for WrapMountable.
func WM(tag *Tag, elements ...any) *ComponentMountable {
	return WrapMountable(tag, elements...)
}

// WrapMountable takes a Tag and creates a Component with it.
func WrapMountable(tag *Tag, elements ...any) *ComponentMountable {
	return &ComponentMountable{
		Component: Wrap(tag, elements),
	}
}
