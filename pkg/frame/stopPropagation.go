package frame

import (
	_ "embed"
	"github.com/gogoracer/racer/pkg/engine"
)

const StopPropagationAttributeName = "data-hlive-sp"

//go:embed js/stopPropagation.js
var StopPropagationJavaScript []byte

func StopPropagation() engine.Attributer {
	attr := &StopPropagationAttribute{
		engine.NewAttribute(StopPropagationAttributeName, ""),
	}

	return attr
}

func StopPropagationRemove(tag engine.Adder) {
	tag.Add(engine.AttrsOff{StopPropagationAttributeName})
}

type StopPropagationAttribute struct {
	*engine.Attribute
}

func (a *StopPropagationAttribute) Initialize(pager engine.Pager) {
	page := pager.(*Page)

	page.DOM().Head().Element(engine.Uber("script").HTML(string(StopPropagationJavaScript)))
}

func (a *StopPropagationAttribute) InitializeSSR(pager engine.Pager) {
	page := pager.(*Page)

	page.DOM().Head().Element(engine.Uber("script").HTML(string(StopPropagationJavaScript)))
}
