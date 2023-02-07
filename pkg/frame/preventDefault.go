package frame

import (
	_ "embed"
	"github.com/gogoracer/racer/pkg/engine"
)

const PreventDefaultAttributeName = "data-hlive-pd"

//go:embed js/preventDefault.js
var PreventDefaultJavaScript []byte

func PreventDefault() *PreventDefaultAttribute {
	attr := &PreventDefaultAttribute{
		engine.NewAttribute(PreventDefaultAttributeName, ""),
	}

	return attr
}

func PreventDefaultRemove(tag engine.Adder) {
	tag.Add(engine.AttrsOff{PreventDefaultAttributeName})
}

type PreventDefaultAttribute struct {
	*engine.Attribute
}

func (a *PreventDefaultAttribute) Initialize(page engine.Pager) {
	page.(*Page).DOM().Head().Element(engine.Uber("script").HTML(string(PreventDefaultJavaScript)))
}

func (a *PreventDefaultAttribute) InitializeSSR(page engine.Pager) {
	page.(*Page).DOM().Head().Element(engine.Uber("script").HTML(string(PreventDefaultJavaScript)))
}
