package frame

import (
	_ "embed"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"

	"github.com/gogoracer/racer/pkg/engine"
)

//go:embed diffapply.js
var DiffApplyScript []byte

// DiffApply is a special event that will trigger when a diff is applied.
// This means that it will trigger itself when first added. This will allow you to know when a change in the tree has
// made it to the browser. You can then, if you wish, immediately remove it from the tree to prevent more triggers.
// You can also add it as a OnOnce and it wil remove itself.

func OnDiffApply(handler engine.EventHandler) (*engine.EventBinding, *DiffApplyAttribute) {
	eb := engine.On(DiffApplyEvent, handler)
	attr := &DiffApplyAttribute{
		Attribute: engine.NewAttribute(DiffApplyAttributeName, eb.ID),
	}

	return eb, attr
}

// TODO: how we remove the attribute once done?
func OnDiffApplyOnce(handler engine.EventHandler) *engine.ElementGroup {
	eb := engine.OnOnce(DiffApplyEvent, handler)
	attr := &DiffApplyAttribute{
		Attribute: engine.NewAttribute(DiffApplyAttributeName, eb.ID),
	}

	return engine.Elements(eb, attr)
}

const (
	DiffApplyEvent         = "diffapply"
	DiffApplyAttributeName = "data-hlive-on-diffapply"
)

type DiffApplyAttribute struct {
	*engine.Attribute

	rendered bool
}

func (a *DiffApplyAttribute) Initialize(page engine.Pager) {
	if a.rendered {
		return
	}

	page.(*Page).DOM().Head().Element(SCRIPT().HTML(string(DiffApplyScript)))
}

func (a *DiffApplyAttribute) InitializeSSR(page engine.Pager) {
	a.rendered = true

	page.(*Page).DOM().Head().Element(SCRIPT().HTML(string(DiffApplyScript)))
}
