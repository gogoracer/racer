// Code generated by qtc from "autogen.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pkg/goggles/gen/autogen.qtpl:3
package gen

//line pkg/goggles/gen/autogen.qtpl:3
import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

//line pkg/goggles/gen/autogen.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pkg/goggles/gen/autogen.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pkg/goggles/gen/autogen.qtpl:10
func snake(s string) string {
	return strcase.ToSnake(s)
}

func camel(s string) string {
	return strcase.ToLowerCamel(s)
}

func pascal(s string) string {
	return strcase.ToCamel(s)
}

func upper(s string) string {
	return strings.ToUpper(s)
}

func lower(s string) string {
	return strings.ToLower(s)
}

//line pkg/goggles/gen/autogen.qtpl:32
func StreamGenerateElement(qw422016 *qt422016.Writer, e *Element) {
//line pkg/goggles/gen/autogen.qtpl:36
	tagPascal := pascal(e.Tag)
	elementName := "Element" + tagPascal
	// attributesName := tagPascal + "Attributes"
	// eventHandlersName := tagPascal + "EventHandlers"

//line pkg/goggles/gen/autogen.qtpl:40
	qw422016.N().S(`
/* cSpell:disable */

package  goggles

import (
    "github.com/gogoracer/racer/pkg/engine"
)

type `)
//line pkg/goggles/gen/autogen.qtpl:49
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:49
	qw422016.N().S(` struct {
    *baseElement
}

func (e *`)
//line pkg/goggles/gen/autogen.qtpl:53
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:53
	qw422016.N().S(`) IsGogglesElement() {}

func `)
//line pkg/goggles/gen/autogen.qtpl:55
	qw422016.E().S(upper(tagPascal))
//line pkg/goggles/gen/autogen.qtpl:55
	qw422016.N().S(`(children ...GogglesElement) *`)
//line pkg/goggles/gen/autogen.qtpl:55
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:55
	qw422016.N().S(` {
    return &`)
//line pkg/goggles/gen/autogen.qtpl:56
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:56
	qw422016.N().S(`{
        baseElement: newBaseElement("`)
//line pkg/goggles/gen/autogen.qtpl:57
	qw422016.E().S(e.Tag)
//line pkg/goggles/gen/autogen.qtpl:57
	qw422016.N().S(`", children...),
    }
}

func (e *`)
//line pkg/goggles/gen/autogen.qtpl:61
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:61
	qw422016.N().S(`) Add(children ...GogglesElement) *`)
//line pkg/goggles/gen/autogen.qtpl:61
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:61
	qw422016.N().S(` {
	e.baseElement.add(children...)
	return e
}

func (e *`)
//line pkg/goggles/gen/autogen.qtpl:66
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:66
	qw422016.N().S(`) Custom(k, v string, dontEscape ...bool) *`)
//line pkg/goggles/gen/autogen.qtpl:66
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:66
	qw422016.N().S(` {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *`)
//line pkg/goggles/gen/autogen.qtpl:71
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:71
	qw422016.N().S(`) BindCustom(k string, v string,  dontEscape ...bool) *`)
//line pkg/goggles/gen/autogen.qtpl:71
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:71
	qw422016.N().S(` {
    e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *`)
//line pkg/goggles/gen/autogen.qtpl:76
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:76
	qw422016.N().S(`) appendAttribute(k string, v string, dontEscape ...bool) *`)
//line pkg/goggles/gen/autogen.qtpl:76
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:76
	qw422016.N().S(` {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}


func (e *`)
//line pkg/goggles/gen/autogen.qtpl:82
	qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:82
	qw422016.N().S(`) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

`)
//line pkg/goggles/gen/autogen.qtpl:86
	for _, a := range e.Attributes {
//line pkg/goggles/gen/autogen.qtpl:86
		qw422016.N().S(`
`)
//line pkg/goggles/gen/autogen.qtpl:87
		attrPascal := pascal(a.Name)

//line pkg/goggles/gen/autogen.qtpl:87
		qw422016.N().S(`// `)
//line pkg/goggles/gen/autogen.qtpl:88
		qw422016.E().S(attrPascal)
//line pkg/goggles/gen/autogen.qtpl:88
		qw422016.N().S(` is the "`)
//line pkg/goggles/gen/autogen.qtpl:88
		qw422016.E().S(a.Name)
//line pkg/goggles/gen/autogen.qtpl:88
		qw422016.N().S(`" attribute.
// `)
//line pkg/goggles/gen/autogen.qtpl:89
		qw422016.E().S(strings.ReplaceAll(a.Description, "\n", "  "))
//line pkg/goggles/gen/autogen.qtpl:89
		qw422016.N().S(`
// Valid values are constrained to the following:
`)
//line pkg/goggles/gen/autogen.qtpl:91
		for _, v := range a.ValidValueTypes {
//line pkg/goggles/gen/autogen.qtpl:91
			qw422016.N().S(`//   * `)
//line pkg/goggles/gen/autogen.qtpl:92
			qw422016.E().S(v)
//line pkg/goggles/gen/autogen.qtpl:92
			qw422016.N().S(`
`)
//line pkg/goggles/gen/autogen.qtpl:93
		}
//line pkg/goggles/gen/autogen.qtpl:93
		qw422016.N().S(`func (element *`)
//line pkg/goggles/gen/autogen.qtpl:94
		qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:94
		qw422016.N().S(`) `)
//line pkg/goggles/gen/autogen.qtpl:94
		qw422016.E().S(attrPascal)
//line pkg/goggles/gen/autogen.qtpl:94
		qw422016.N().S(`(v string, dontEscape ...bool) *`)
//line pkg/goggles/gen/autogen.qtpl:94
		qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:94
		qw422016.N().S(` {
    element.appendAttribute("`)
//line pkg/goggles/gen/autogen.qtpl:95
		qw422016.E().S(a.Name)
//line pkg/goggles/gen/autogen.qtpl:95
		qw422016.N().S(`", v, dontEscape...)
    return element
}
`)
//line pkg/goggles/gen/autogen.qtpl:98
	}
//line pkg/goggles/gen/autogen.qtpl:98
	qw422016.N().S(`
`)
//line pkg/goggles/gen/autogen.qtpl:100
	for _, eh := range e.EventHandlers {
//line pkg/goggles/gen/autogen.qtpl:100
		qw422016.N().S(`
`)
//line pkg/goggles/gen/autogen.qtpl:101
		evt := "On" + pascal(eh.Name[2:])

//line pkg/goggles/gen/autogen.qtpl:101
		qw422016.N().S(`// `)
//line pkg/goggles/gen/autogen.qtpl:102
		qw422016.E().S(eh.Description)
//line pkg/goggles/gen/autogen.qtpl:102
		qw422016.N().S(`
func (e *`)
//line pkg/goggles/gen/autogen.qtpl:103
		qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:103
		qw422016.N().S(`) `)
//line pkg/goggles/gen/autogen.qtpl:103
		qw422016.E().S(evt)
//line pkg/goggles/gen/autogen.qtpl:103
		qw422016.N().S(`(fn engine.EventHandler) (*`)
//line pkg/goggles/gen/autogen.qtpl:103
		qw422016.E().S(elementName)
//line pkg/goggles/gen/autogen.qtpl:103
		qw422016.N().S(`) {
    if fn == nil {
        return e
    }

    e.shouldBeComponent = true
    e.children = append(
        e.children,
        EventHandler("`)
//line pkg/goggles/gen/autogen.qtpl:111
		qw422016.E().S(eh.Name[2:])
//line pkg/goggles/gen/autogen.qtpl:111
		qw422016.N().S(`", fn),
    )
    return e
}
`)
//line pkg/goggles/gen/autogen.qtpl:115
	}
//line pkg/goggles/gen/autogen.qtpl:115
	qw422016.N().S(`
`)
//line pkg/goggles/gen/autogen.qtpl:117
}

//line pkg/goggles/gen/autogen.qtpl:117
func WriteGenerateElement(qq422016 qtio422016.Writer, e *Element) {
//line pkg/goggles/gen/autogen.qtpl:117
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/goggles/gen/autogen.qtpl:117
	StreamGenerateElement(qw422016, e)
//line pkg/goggles/gen/autogen.qtpl:117
	qt422016.ReleaseWriter(qw422016)
//line pkg/goggles/gen/autogen.qtpl:117
}

//line pkg/goggles/gen/autogen.qtpl:117
func GenerateElement(e *Element) string {
//line pkg/goggles/gen/autogen.qtpl:117
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/goggles/gen/autogen.qtpl:117
	WriteGenerateElement(qb422016, e)
//line pkg/goggles/gen/autogen.qtpl:117
	qs422016 := string(qb422016.B)
//line pkg/goggles/gen/autogen.qtpl:117
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/goggles/gen/autogen.qtpl:117
	return qs422016
//line pkg/goggles/gen/autogen.qtpl:117
}

//line pkg/goggles/gen/autogen.qtpl:119
func StreamGenerateIcon(qw422016 *qt422016.Writer, pkg iconPackage) {
//line pkg/goggles/gen/autogen.qtpl:119
	qw422016.N().S(`package `)
//line pkg/goggles/gen/autogen.qtpl:120
	qw422016.E().S(pkg.Name)
//line pkg/goggles/gen/autogen.qtpl:120
	qw422016.N().S(`

import "github.com/gogoracer/racer/pkg/engine"

const (
`)
//line pkg/goggles/gen/autogen.qtpl:125
	for _, namedIcon := range pkg.Icons {
//line pkg/goggles/gen/autogen.qtpl:126
		cn := camel(namedIcon.Name) + "InnerSVG"

//line pkg/goggles/gen/autogen.qtpl:127
		qw422016.E().S(cn)
//line pkg/goggles/gen/autogen.qtpl:127
		qw422016.N().S(` = `)
//line pkg/goggles/gen/autogen.qtpl:127
		qw422016.N().S("`")
//line pkg/goggles/gen/autogen.qtpl:127
		qw422016.N().S(namedIcon.Icon.SvgBody)
//line pkg/goggles/gen/autogen.qtpl:127
		qw422016.N().S(``)
//line pkg/goggles/gen/autogen.qtpl:127
		qw422016.N().S("`")
//line pkg/goggles/gen/autogen.qtpl:127
		qw422016.N().S(`
`)
//line pkg/goggles/gen/autogen.qtpl:128
	}
//line pkg/goggles/gen/autogen.qtpl:128
	qw422016.N().S(`)

var sharedIconAttrs =  engine.Attrs{"width": "1em","height": "1em"}

`)
//line pkg/goggles/gen/autogen.qtpl:133
	for _, namedIcon := range pkg.Icons {
//line pkg/goggles/gen/autogen.qtpl:135
		cn := camel(namedIcon.Name) + "InnerSVG"
		viewBox := fmt.Sprintf(`0 0 %d %d`, pkg.Width, pkg.Height)

//line pkg/goggles/gen/autogen.qtpl:137
		qw422016.N().S(`
func `)
//line pkg/goggles/gen/autogen.qtpl:138
		qw422016.E().S(namedIcon.Name)
//line pkg/goggles/gen/autogen.qtpl:138
		qw422016.N().S(`(children ...any) *engine.HTMLElement {
    return engine.Element(
        "svg",
        sharedIconAttrs,
        engine.Attrs{
            "viewBox": "`)
//line pkg/goggles/gen/autogen.qtpl:143
		qw422016.E().S(viewBox)
//line pkg/goggles/gen/autogen.qtpl:143
		qw422016.N().S(`",
        },
        `)
//line pkg/goggles/gen/autogen.qtpl:145
		qw422016.E().S(cn)
//line pkg/goggles/gen/autogen.qtpl:145
		qw422016.N().S(`,
        children,
    )
}
`)
//line pkg/goggles/gen/autogen.qtpl:149
	}
//line pkg/goggles/gen/autogen.qtpl:149
	qw422016.N().S(`
`)
//line pkg/goggles/gen/autogen.qtpl:160
	qw422016.N().S(`
`)
//line pkg/goggles/gen/autogen.qtpl:161
}

//line pkg/goggles/gen/autogen.qtpl:161
func WriteGenerateIcon(qq422016 qtio422016.Writer, pkg iconPackage) {
//line pkg/goggles/gen/autogen.qtpl:161
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/goggles/gen/autogen.qtpl:161
	StreamGenerateIcon(qw422016, pkg)
//line pkg/goggles/gen/autogen.qtpl:161
	qt422016.ReleaseWriter(qw422016)
//line pkg/goggles/gen/autogen.qtpl:161
}

//line pkg/goggles/gen/autogen.qtpl:161
func GenerateIcon(pkg iconPackage) string {
//line pkg/goggles/gen/autogen.qtpl:161
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/goggles/gen/autogen.qtpl:161
	WriteGenerateIcon(qb422016, pkg)
//line pkg/goggles/gen/autogen.qtpl:161
	qs422016 := string(qb422016.B)
//line pkg/goggles/gen/autogen.qtpl:161
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/goggles/gen/autogen.qtpl:161
	return qs422016
//line pkg/goggles/gen/autogen.qtpl:161
}
