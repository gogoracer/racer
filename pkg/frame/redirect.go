package frame

import (
	_ "embed"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"

	"github.com/gogoracer/racer/pkg/engine"
)

const RedirectAttributeName = "data-redirect"

//go:embed redirect.js
var RedirectJavaScript []byte

func Redirect(url string) engine.Attributer {
	attr := &RedirectAttribute{
		Attribute: engine.NewAttribute(RedirectAttributeName, url),
	}

	return attr
}

type RedirectAttribute struct {
	*engine.Attribute

	rendered bool
}

func (a *RedirectAttribute) Initialize(page engine.Pager) {
	if a.rendered {
		return
	}

	page.(*Page).DOM().Head().Element(SCRIPT().HTML(string(RedirectJavaScript)))
}

func (a *RedirectAttribute) InitializeSSR(page engine.Pager) {
	a.rendered = true
	page.(*Page).DOM().Head().Element(SCRIPT().HTML(string(RedirectJavaScript)))
}
