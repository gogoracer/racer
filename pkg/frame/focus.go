package frame

import (
	_ "embed"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"

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

func (a *FocusAttribute) Initialize(page engine.Pager) {
	if a.rendered {
		return
	}

	page.(*Page).DOM().Head().Element(SCRIPT().HTML(string(FocusJavaScript)))
}

func (a *FocusAttribute) InitializeSSR(page engine.Pager) {
	a.rendered = true
	page.(*Page).DOM().Head().Element(SCRIPT().HTML(string(FocusJavaScript)))
}
