package racer_attribute

import "github.com/gogoracer/racer/pkg/engine"

func REL(value string) *engine.Attribute {
	return engine.NewAttribute("rel", value)
}

func HREF(value string) *engine.Attribute {
	return engine.NewAttribute("href", value)
}

func TYPE(value string) *engine.Attribute {
	return engine.NewAttribute("type", value)
}

func CLASS(value string) *engine.Attribute {
	return engine.NewAttribute("class", value)
}

func DEFER(value bool) *engine.Attribute {
	s := "false"
	if value {
		s = "true"
	}
	return engine.NewAttribute("defer", s)
}

func DISABLED(value bool) *engine.Attribute {
	s := "false"
	if value {
		s = "true"
	}
	return engine.NewAttribute("disabled", s)
}

func SRC(value string) *engine.Attribute {
	return engine.NewAttribute("src", value)
}

func VALUE(value string) *engine.Attribute {
	return engine.NewAttribute("value", value)
}

func STYLE(value string) *engine.Attribute {
	return engine.NewAttribute("style", value)
}
