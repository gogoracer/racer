package engine

import (
	"context"
)

func Uber(tagName string) *UberElement {
	ue := &UberElement{
		c: NewComponentPubSub(tagName),
	}

	return ue
}

type UberChild interface {
	isValidUberChild()
}

type UberElement struct {
	c *ComponentPubSub
}

func (ue *UberElement) isValidUberChild() {}

func (ue *UberElement) GetComponent() Componenter {
	return ue.c
}

// TODO: do something better
func (ue *UberElement) GetComponentPubSub() *ComponentPubSub {
	return ue.c
}

func (ue *UberElement) Attr(attrs ...Attributer) *UberElement {
	ue.c.AddAttributes(typeToAny(attrs)...)
	return ue
}

func (ue *UberElement) AttrRemove(keys ...string) *UberElement {
	ue.c.RemoveAttributes(keys...)
	return ue
}

func (ue *UberElement) Element(childElements ...UberChild) *UberElement {
	for _, element := range childElements {
		ue.c.Add(element)
	}
	return ue
}

func (ue *UberElement) Component(components ...GetComponenter) *UberElement {
	ue.c.Add(typeToAny(components)...)
	return ue
}

func (ue *UberElement) LockBox(box LockBoxer) *UberElement {
	ue.c.Add(box)
	return ue
}

func (ue *UberElement) Box(box NodeBoxer) *UberElement {
	ue.c.Add(box)
	return ue
}

type UberElementBox struct {
	*NodeBox[*UberElement]
}

func (ueb UberElementBox) isValidUberChild() {}

func (ue *UberElement) IntoBox() UberElementBox {
	return UberElementBox{NodeBox: Box(ue)}
}

func (ue *UberElement) On(eventBindings ...*EventBinding) *UberElement {
	ue.c.Add(typeToAny(eventBindings)...)
	return ue
}

func (ue *UberElement) RemoveEventBinding(eventBindingID string) *UberElement {
	ue.c.RemoveEventBinding(eventBindingID)

	return ue
}

func (ue *UberElement) Class(class string) *UberElement {
	ue.c.Add(Class(class))
	return ue
}

func (ue *UberElement) ClassOff(class string) *UberElement {
	ue.c.Add(ClassOff(class))
	return ue
}

func (ue *UberElement) Style(key, value string) *UberElement {
	ue.c.Add(Style{key: value})
	return ue
}

func (ue *UberElement) Styles(styles Style) *UberElement {
	ue.c.Add(styles)
	return ue
}

func (ue *UberElement) HTML(htmlVal string) *UberElement {
	ue.c.Add(HTML(htmlVal))
	return ue
}

func (ue *UberElement) SetMount(mount func(ctx context.Context)) *UberElement {
	ue.c.SetMount(mount)
	return ue
}

func (ue *UberElement) SetUnmount(unmount func(ctx context.Context)) *UberElement {
	ue.c.SetUnmount(unmount)
	return ue
}

func (ue *UberElement) SetMountPubSub(f func(ctx context.Context, pubSub *PubSub)) *UberElement {
	ue.c.SetMountPubSub(f)
	return ue
}

func (ue *UberElement) SetAfterMountPubSub(f func(ctx context.Context, pubSub *PubSub)) *UberElement {
	ue.c.SetAfterMountPubSub(f)
	return ue
}

////////

func (ue *UberElement) Val() *UberValue {
	return &UberValue{ue: ue}
}

type UberValue struct {
	ue *UberElement
}

func (uv *UberValue) Str(val string) *UberElement {
	uv.ue.c.Add(val)
	return uv.ue
}

func (uv *UberValue) BindStr(val *string) *UberElement {
	uv.ue.c.Add(val)
	return uv.ue
}

func (uv *UberValue) Int(val int) *UberElement {
	uv.ue.c.Add(val)
	return uv.ue
}

// TODO: int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64,

func typeToAny[T any](tt []T) []interface{} {
	anyVals := make([]interface{}, len(tt))
	for i, v := range tt {
		anyVals[i] = v
	}
	return anyVals
}
