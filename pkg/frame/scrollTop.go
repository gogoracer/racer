package frame

import (
	_ "embed"
	"strconv"

	"github.com/gogoracer/racer/pkg/engine"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"
)

const ScrollTopAttributeName = "data-scrollTop"

//go:embed scrollTop.js
var ScrollTopJavaScript []byte

func ScrollTop(val int) engine.Attributer {
	attr := &ScrollTopAttribute{
		Attribute: engine.NewAttribute(ScrollTopAttributeName, strconv.Itoa(val)),
	}

	return attr
}

func ScrollTopRemove(tag engine.Adder) {
	tag.Add(engine.AttrsOff{ScrollTopAttributeName})
}

type ScrollTopAttribute struct {
	*engine.Attribute

	rendered bool
}

func (a *ScrollTopAttribute) Initialize(page engine.Pager) {
	if a.rendered {
		return
	}

	page.(*Page).DOM().Head().Element(SCRIPT().HTML(string(ScrollTopJavaScript)))
}

func (a *ScrollTopAttribute) InitializeSSR(page engine.Pager) {
	a.rendered = true
	page.(*Page).DOM().Head().Element(SCRIPT().HTML(string(ScrollTopJavaScript)))
}
