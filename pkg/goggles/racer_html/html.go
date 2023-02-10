package racer_html

import (
	"fmt"

	"github.com/gogoracer/racer/pkg/engine"
	"github.com/gogoracer/racer/pkg/goggles/racer_attribute"
)

func LINK(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("link").Element(children...)
}

func HEADER(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("header")
}

func H1(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("h1").Element(children...)
}

func H2(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("h2").Element(children...)
}

func H3(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("h3").Element(children...)
}

func H4(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("h4").Element(children...)
}

func P(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("p").Element(children...)
}

func MAIN(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("main").Element(children...)
}

func BUTTON(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("button").Element(children...)
}

func SCRIPT(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("script").Element(children...)
}

func STYLETAG(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("style").Element(children...)
}

func DIV(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("div").Element(children...)
}

func SPAN(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("span").Element(children...)
}

func PRE(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("pre").Element(children...)
}

func CODE(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("code").Element(children...)
}

func HR(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("hr").Element(children...)
}

func HYPERLINK(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("a").Element(children...)
}

func INPUT(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("input").Element(children...)
}

func FORM(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("form").Element(children...)
}

func BR(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("br").Element(children...)
}

func LABEL(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("label").Element(children...)
}

func IMG(children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("img").Element(children...)
}

func DIV_X(class string, children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("div").Class(class).Element(children...)
}

func BUTTON_X(class string, children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("button").Class(class).Element(children...)
}

func IMG_X(class string, src string, children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("img").Class(class).Attr(racer_attribute.SRC(src)).Element(children...)
}

func PRE_X(class string, children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("pre").Class(class).Element(children...)
}

func HYPERLINK_X(class, href string, children ...engine.UberChild) *engine.UberElement {
	return engine.Uber("a").Class(class).Attr(racer_attribute.HREF(href)).Element(children...)
}

func BUTTON_ALPINE_REDIRECT_X(class string, href string, children ...engine.UberChild) *engine.UberElement {
	return BUTTON_X(class, children...).Attr(
		engine.NewAttribute("x-data", fmt.Sprintf(`{
			redirect(){
				window.location.href = "%s"
			}
		}`, href)),
		racer_attribute.ALPINE_CLICK("redirect"),
	)
}

func SPAN_STRING(s string) *engine.UberElement {
	return SPAN().Val().Str(s)
}
