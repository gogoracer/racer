package frame

import (
	"context"
	_ "embed"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"

	"github.com/gogoracer/racer/pkg/engine"
)

const (
	PreemptDisableAttributeName = "data-hlive-pre-disable"
)

//go:embed preemptDisableOnClick.js
var PreemptDisableOnClickJavaScript []byte

// TODO: Once?
func PreemptDisableOn(eb *engine.EventBinding) *engine.ElementGroup {
	sourceAttr := &PreemptDisableAttribute{
		Attribute: engine.NewAttribute(PreemptDisableAttributeName, eb.Name),
	}

	ogHandler := eb.Handler

	eb.Handler = func(ctx context.Context, e engine.Event) {
		// Update the Browser DOM with what we've done client first
		if sourceAttr.page != nil {
			if browserTag := sourceAttr.page.GetPage().GetBrowserNodeByID(e.Binding.Component.GetID()); browserTag != nil {
				browserTag.Add(engine.Attrs{"disabled": ""})
			}
		}
		// Update the Page DOM
		if adder, ok := e.Binding.Component.(engine.Adder); ok {
			adder.Add(engine.Attrs{"disabled": ""})
		} else {
			panic("PreemptDisableOn: bound Component must be an Adder")
			// engine.LoggerDev.Error().Msg("PreemptDisableOn: bound Component must be an Adder")
		}

		// Call original handler
		if ogHandler != nil {
			ogHandler(ctx, e)
		}
	}

	return engine.Elements(eb, sourceAttr)
}

type PreemptDisableAttribute struct {
	*engine.Attribute

	page     *Page
	rendered bool
}

func (a *PreemptDisableAttribute) Initialize(page engine.Pager) {
	if a.rendered {
		return
	}

	a.page = page.(*Page)
	a.page.DOM().Head().Element(SCRIPT().HTML(string(PreemptDisableOnClickJavaScript)))
}

func (a *PreemptDisableAttribute) InitializeSSR(pager engine.Pager) {
	a.rendered = true

	a.page = pager.(*Page)
	a.page.DOM().Head().Element(SCRIPT().HTML(string(PreemptDisableOnClickJavaScript)))
}
