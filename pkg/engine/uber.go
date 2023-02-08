package engine

import "context"

func Uber(tagName string) *UberElement {
	ue := &UberElement{
		c: NewComponentPubSub(tagName),
	}

	return ue
}

type UberElement struct {
	c *ComponentPubSub
}

func (ue *UberElement) GetComponent() Componenter {
	return ue.c
}

func (ue *UberElement) Attr(attr Attributer) *UberElement {
	ue.c.AddAttributes(attr)

	return ue
}

func (ue *UberElement) AttrRemove(key string) *UberElement {
	ue.c.RemoveAttributes(key)

	return ue
}

func (ue *UberElement) Element(element *UberElement) *UberElement {
	ue.c.Add(element)

	return ue
}

func (ue *UberElement) Component(component GetComponenter) *UberElement {
	ue.c.Add(component)

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

func (ue *UberElement) On(eventBinding *EventBinding) *UberElement {
	ue.c.Add(eventBinding)

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

func (ue *UberElement) SetMount(mount func(ctx context.Context)) {
	ue.c.SetMount(mount)
}

func (ue *UberElement) SetUnmount(unmount func(ctx context.Context)) {
	ue.c.SetUnmount(unmount)
}

func (ue *UberElement) SetMountPubSub(f func(ctx context.Context, pubSub *PubSub)) {
	ue.c.SetMountPubSub(f)
}

func (ue *UberElement) SetAfterMountPubSub(f func(ctx context.Context, pubSub *PubSub)) {
	ue.c.SetAfterMountPubSub(f)
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

func (uv *UberValue) Int(val int) *UberElement {
	uv.ue.c.Add(val)

	return uv.ue
}

// TODO: int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64,
