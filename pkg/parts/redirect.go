package parts

import (
	_ "embed"

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

func (a *RedirectAttribute) Initialize(page *engine.Page) {
	if a.rendered {
		return
	}

	page.DOM().Head().Add(engine.NewTag("script", engine.HTML(RedirectJavaScript)))
}

func (a *RedirectAttribute) InitializeSSR(page *engine.Page) {
	a.rendered = true
	page.DOM().Head().Add(engine.NewTag("script", engine.HTML(RedirectJavaScript)))
}
