package engine

import (
	_ "embed"
)

const PreventDefaultAttributeName = "data-hlive-pd"

//go:embed js/preventDefault.js
var PreventDefaultJavaScript []byte

func PreventDefault() *PreventDefaultAttribute {
	attr := &PreventDefaultAttribute{
		NewAttribute(PreventDefaultAttributeName, ""),
	}

	return attr
}

func PreventDefaultRemove(tag Adder) {
	tag.Add(AttrsOff{PreventDefaultAttributeName})
}

type PreventDefaultAttribute struct {
	*Attribute
}

func (a *PreventDefaultAttribute) Initialize(page *Page) {
	page.DOM().Head().Add(NewTag("script", HTML(PreventDefaultJavaScript)))
}

func (a *PreventDefaultAttribute) InitializeSSR(page *Page) {
	page.DOM().Head().Add(NewTag("script", HTML(PreventDefaultJavaScript)))
}
