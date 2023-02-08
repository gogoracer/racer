package racer_html

import (
	"github.com/gogoracer/racer/pkg/engine"
)

func LINK() *engine.UberElement {
	return engine.Uber("link")
}

func HEADER() *engine.UberElement {
	return engine.Uber("header")
}

func H1() *engine.UberElement {
	return engine.Uber("h1")
}

func H2() *engine.UberElement {
	return engine.Uber("h2")
}

func H3() *engine.UberElement {
	return engine.Uber("h3")
}

func H4() *engine.UberElement {
	return engine.Uber("h4")
}

func P() *engine.UberElement {
	return engine.Uber("p")
}

func MAIN() *engine.UberElement {
	return engine.Uber("main")
}

func BUTTON() *engine.UberElement {
	return engine.Uber("button")
}

func SCRIPT() *engine.UberElement {
	return engine.Uber("script")
}

func STYLE() *engine.UberElement {
	return engine.Uber("style")
}

func DIV() *engine.UberElement {
	return engine.Uber("div")
}

func PRE() *engine.UberElement {
	return engine.Uber("pre")
}

func CODE() *engine.UberElement {
	return engine.Uber("code")
}

func HR() *engine.UberElement {
	return engine.Uber("hr")
}
