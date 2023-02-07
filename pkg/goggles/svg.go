package goggles

type ElementSVG struct {
	*baseElement
}

func SVG(children ...GogglesElement) *ElementSVG {
	return &ElementSVG{
		baseElement: newBaseElement("svg", children...),
	}
}

func (e *ElementSVG) Custom(k, v string, dontEscape ...bool) *ElementSVG {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}
