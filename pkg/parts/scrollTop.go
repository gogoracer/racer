package parts

import (
	_ "embed"
	"strconv"

	"github.com/gogoracer/racer/pkg/engine"
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

func (a *ScrollTopAttribute) Initialize(page *engine.Page) {
	if a.rendered {
		return
	}

	page.DOM().Head().Add(engine.NewTag("script", engine.HTML(ScrollTopJavaScript)))
}

func (a *ScrollTopAttribute) InitializeSSR(page *engine.Page) {
	a.rendered = true
	page.DOM().Head().Add(engine.NewTag("script", engine.HTML(ScrollTopJavaScript)))
}
