package handlebars

import "github.com/gogoracer/racer/pkg/engine"

type Page struct {
	headElements []IsHandlebarElement
	bodyElements []IsHandlebarElement
	// page *engine.Page
}

func NewPage() *Page {
	return &Page{}
}

func (p *Page) Head(children ...IsHandlebarElement) *Page {
	p.headElements = append(p.headElements, children...)
	return p
}

func (p *Page) Body(children ...IsHandlebarElement) *Page {
	p.bodyElements = append(p.bodyElements, children...)
	return p
}

func (p *Page) GenerateVDOM() *engine.Page {
	page := engine.NewPage()

	for _, element := range p.headElements {
		page.DOM().Head().Add(element.GenerateVDOM())
	}

	for _, element := range p.bodyElements {
		page.DOM().Body().Add(element.GenerateVDOM())
	}

	return page
}
