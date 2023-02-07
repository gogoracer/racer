package parts

import (
	_ "embed"

	"github.com/gogoracer/racer/pkg/engine"
)

const FocusAttributeName = "data-hlive-focus"

//go:embed focus.js
var FocusJavaScript []byte

func Focus() engine.Attributer {
	attr := &FocusAttribute{
		Attribute: engine.NewAttribute(FocusAttributeName, ""),
	}

	return attr
}

func FocusRemove(tag engine.Adder) {
	tag.Add(engine.AttrsOff{FocusAttributeName})
}

type FocusAttribute struct {
	*engine.Attribute

	rendered bool
}

func (a *FocusAttribute) Initialize(page *engine.Page) {
	if a.rendered {
		return
	}

	page.DOM().Head().Add(engine.NewTag("script", engine.HTML(FocusJavaScript)))
}

func (a *FocusAttribute) InitializeSSR(page *engine.Page) {
	a.rendered = true
	page.DOM().Head().Add(engine.NewTag("script", engine.HTML(FocusJavaScript)))
}
