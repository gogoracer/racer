package racer_attribute

import (
	"fmt"
	"strings"

	"github.com/gogoracer/racer/pkg/engine"
)

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

func DEFER() *engine.Attribute {
	return engine.NewAttribute("defer", "true")
}

func DISABLED() *engine.Attribute {
	return engine.NewAttribute("disabled", "true")
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

func ALPINE_CLICK(value string) *engine.Attribute {
	return engine.NewAttribute("@click", value)
}

func ALPINE_SHOW_ON(value string) *engine.Attribute {
	return engine.NewAttribute("x-show", value)
}

func ALPINE_ON_CLICK_AWAY(value string) *engine.Attribute {
	return engine.NewAttribute("@click.away", value)
}

func ALPINE_START_HIDDEN() *engine.Attribute {
	return engine.NewAttribute("style", "display: none;")
}

func ALPINE_TRANSITION() *engine.Attribute {
	return engine.NewAttribute("x-transition", "")
}

func ALPINE_KEYDOWN(keyName, fnName string) *engine.Attribute {
	return engine.NewAttribute(fmt.Sprintf("@keydown.%s.prevent.stop", keyName), fnName)
}

type ALPINE_DATA_ENTRY string

func ALPINE_DATA(data ...ALPINE_DATA_ENTRY) *engine.Attribute {
	sb := strings.Builder{}
	sb.WriteString("{\n")
	for _, datum := range data {
		sb.WriteString("  ")
		sb.WriteString(string(datum))
		sb.WriteString(",\n")
	}
	sb.WriteString(" }")
	return engine.NewAttribute("x-data", sb.String())
}

func ALPINE_DATA_STRING(name, value string) ALPINE_DATA_ENTRY {
	return ALPINE_DATA_ENTRY(fmt.Sprintf("%s: '%s'", name, value))
}

func ALPINE_DATA_BOOL(name string, value bool) ALPINE_DATA_ENTRY {
	return ALPINE_DATA_ENTRY(fmt.Sprintf("%s: %t", name, value))
}

func ALPINE_DATA_FUNCTION(name, contents string) ALPINE_DATA_ENTRY {
	return ALPINE_DATA_ENTRY(fmt.Sprintf("%s() { %s }", name, contents))
}
