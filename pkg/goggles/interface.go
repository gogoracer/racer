package goggles

import "github.com/gogoracer/racer/pkg/engine"

type GogglesElement interface {
	IsGogglesElement()
	GenerateVDOM() interface{}
}

type baseElement struct {
	name              string
	shouldBeComponent bool
	attrs             map[string]*engine.Attribute
	children          []GogglesElement
}

func newBaseElement(name string, children ...GogglesElement) *baseElement {
	return &baseElement{
		name:     name,
		attrs:    map[string]*engine.Attribute{},
		children: children,
	}
}

func (e *baseElement) add(children ...GogglesElement) *baseElement {
	e.children = append(e.children, children...)
	return e
}

func (e *baseElement) custom(k, v string, dontEscape ...bool) *baseElement {
	e.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *baseElement) bindCustom(k, v string, dontEscape ...bool) *baseElement {
	e.custom(k, v)
	e.shouldBeComponent = true
	return e
}

func (e *baseElement) appendAttribute(k string, v string, dontEscape ...bool) *baseElement {
	if k == "" || v == "" {
		return e
	}
	attr, exists := e.attrs[k]
	if exists {
		attr.SetValue(attr.GetValue() + " " + v)
	} else {
		attr = engine.NewAttribute(k, v)
	}
	attr.SetNoEscapeString(len(dontEscape) > 0 && dontEscape[0])
	e.attrs[k] = attr
	return e
}

func (e *baseElement) generateVDOM() interface{} {
	childVDOMs := make([]interface{}, 0, len(e.attrs)+len(e.children))
	for _, attr := range e.attrs {
		childVDOMs = append(childVDOMs, attr)
	}
	for _, child := range e.children {
		switch c := child.(type) {
		case GogglesElement:
			childVDOMs = append(childVDOMs, c.GenerateVDOM())
		default:
			childVDOMs = append(childVDOMs, c)
		}
	}
	if e.shouldBeComponent {
		return engine.NewComponent(e.name, childVDOMs...)
	} else {
		return engine.NewTag(e.name, childVDOMs...)
	}
}

type EventHandlerWrapper struct {
	Name    string
	Handler engine.EventHandler
}

func EventHandler(name string, handler engine.EventHandler) EventHandlerWrapper {
	return EventHandlerWrapper{Name: name, Handler: handler}
}

func (e EventHandlerWrapper) IsGogglesElement() {}

func (e EventHandlerWrapper) GenerateVDOM() interface{} {
	return engine.On(e.Name, e.Handler)
}
