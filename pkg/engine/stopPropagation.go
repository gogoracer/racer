package engine

import (
	_ "embed"
)

const StopPropagationAttributeName = "data-hlive-sp"

//go:embed js/stopPropagation.js
var StopPropagationJavaScript []byte

func StopPropagation() Attributer {
	attr := &StopPropagationAttribute{
		NewAttribute(StopPropagationAttributeName, ""),
	}

	return attr
}

func StopPropagationRemove(tag Adder) {
	tag.Add(AttrsOff{StopPropagationAttributeName})
}

type StopPropagationAttribute struct {
	*Attribute
}

func (a *StopPropagationAttribute) Initialize(page *Page) {
	page.DOM().Head().Add(NewTag("script", HTML(StopPropagationJavaScript)))
}

func (a *StopPropagationAttribute) InitializeSSR(page *Page) {
	page.DOM().Head().Add(NewTag("script", HTML(StopPropagationJavaScript)))
}