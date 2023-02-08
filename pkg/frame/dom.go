package frame

import "github.com/gogoracer/racer/pkg/engine"

type DOM struct {
	// Root DOM elements
	docType engine.HTML
	html    *engine.UberElement
	head    *engine.UberElement
	meta    *engine.UberElement
	title   *engine.UberElement
	body    *engine.UberElement
}

func NewDOM() DOM {
	dom := DOM{
		docType: "<!doctype html>",
		html:    engine.Uber("html").Attr(engine.NewAttribute("lang", "en")),
		head:    engine.Uber("head"),
		meta:    engine.Uber("meta").Attr(engine.NewAttribute("charset", "utf-8")),
		title:   engine.Uber("title"),
		body:    engine.Uber("body"),
	}

	dom.head.Element(dom.meta).Element(dom.title)
	dom.html.Element(dom.head).Element(dom.body)

	return dom
}

func (dom DOM) DocType() engine.HTML {
	return dom.docType
}

func (dom DOM) HTML() *engine.UberElement {
	return dom.html
}

func (dom DOM) Head(children ...*engine.UberElement) *engine.UberElement {
	head := dom.head
	head.Element(children...)
	return dom.head
}

func (dom DOM) Meta(children ...*engine.UberElement) *engine.UberElement {
	meta := dom.meta
	meta.Element(children...)
	return dom.meta
}

func (dom DOM) Title() *engine.UberElement {
	return dom.title
}

func (dom DOM) Body(children ...*engine.UberElement) *engine.UberElement {
	body := dom.body
	body.Element(children...)
	return body
}
