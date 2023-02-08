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
