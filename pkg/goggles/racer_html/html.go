package racer_html

import (
	"github.com/gogoracer/racer/pkg/engine"
)

func LINK(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("link").Element(children...)
}

func HEADER(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("header")
}

func H1(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("h1").Element(children...)
}

func H2(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("h2").Element(children...)
}

func H3(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("h3").Element(children...)
}

func H4(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("h4").Element(children...)
}

func P(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("p").Element(children...)
}

func MAIN(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("main").Element(children...)
}

func BUTTON(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("button").Element(children...)
}

func SCRIPT(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("script").Element(children...)
}

func STYLETAG(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("style").Element(children...)
}

func DIV(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("div").Element(children...)
}

func SPAN(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("span").Element(children...)
}

func PRE(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("pre").Element(children...)
}

func CODE(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("code").Element(children...)
}

func HR(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("hr").Element(children...)
}

func HYPERLINK(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("a").Element(children...)
}

func INPUT(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("input").Element(children...)
}

func FORM(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("form").Element(children...)
}

func BR(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("br").Element(children...)
}

func LABEL(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("label").Element(children...)
}

func IMG(children ...*engine.UberElement) *engine.UberElement {
	return engine.Uber("img").Element(children...)
}
