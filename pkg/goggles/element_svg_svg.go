/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementSVGSvg struct {
	*baseElement
}

func SVGSVG(children ...any) *ElementSVGSvg {
	return &ElementSVGSvg{
		baseElement: newBaseElement("SVG svg", children...),
	}
}

func (e *ElementSVGSvg) Add(children ...any) *ElementSVGSvg {
	e.baseElement.add(children...)
	return e
}

func (e *ElementSVGSvg) Custom(k, v string, dontEscape ...bool) *ElementSVGSvg {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementSVGSvg) BindCustom(k string, v string, dontEscape ...bool) *ElementSVGSvg {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementSVGSvg) appendAttribute(k string, v string, dontEscape ...bool) *ElementSVGSvg {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementSVGSvg) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementSVGSvg) Accesskey(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementSVGSvg) Contenteditable(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementSVGSvg) Dir(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementSVGSvg) Hidden(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementSVGSvg) Translate(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementSVGSvg) Itemid(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementSVGSvg) Itemprop(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementSVGSvg) Itemref(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementSVGSvg) Itemscope(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementSVGSvg) Nonce(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementSVGSvg) Style(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementSVGSvg) Draggable(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementSVGSvg) Id(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementSVGSvg) Popover(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Autocapitalize is the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Valid values are constrained to the following:
//   - on
//   - on
//   - off
//   - off
//   - none
//   - none
//   - sentences
//   - sentences
//   - words
//   - words
//   - characters
//   - characters
func (element *ElementSVGSvg) Autocapitalize(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Enterkeyhint is the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Valid values are constrained to the following:
//   - enter
//   - enter
//   - done
//   - done
//   - go
//   - go
//   - next
//   - next
//   - previous
//   - previous
//   - search
//   - search
//   - send
//   - send
func (element *ElementSVGSvg) Enterkeyhint(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementSVGSvg) Is(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementSVGSvg) Lang(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementSVGSvg) Spellcheck(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("spellcheck", v, dontEscape...)
	return element
}

// Inputmode is the "inputmode" attribute.
// Hint for selecting an input modality
// Valid values are constrained to the following:
//   - none
//   - none
//   - text
//   - text
//   - tel
//   - tel
//   - email
//   - email
//   - url
//   - url
//   - numeric
//   - numeric
//   - decimal
//   - decimal
//   - search
//   - search
func (element *ElementSVGSvg) Inputmode(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementSVGSvg) Autofocus(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementSVGSvg) Itemtype(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementSVGSvg) Slot(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementSVGSvg) Tabindex(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementSVGSvg) Class(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementSVGSvg) Inert(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementSVGSvg) Title(v string, dontEscape ...bool) *ElementSVGSvg {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementSVGSvg) OnAuxclick(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("auxclick", fn),
	)
	return e
}

// beforematch event handler
func (e *ElementSVGSvg) OnBeforematch(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("beforematch", fn),
	)
	return e
}

// beforetoggle event handler
func (e *ElementSVGSvg) OnBeforetoggle(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("beforetoggle", fn),
	)
	return e
}

// blur event handler
func (e *ElementSVGSvg) OnBlur(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("blur", fn),
	)
	return e
}

// cancel event handler
func (e *ElementSVGSvg) OnCancel(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("cancel", fn),
	)
	return e
}

// canplay event handler
func (e *ElementSVGSvg) OnCanplay(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("canplay", fn),
	)
	return e
}

// canplaythrough event handler
func (e *ElementSVGSvg) OnCanplaythrough(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("canplaythrough", fn),
	)
	return e
}

// change event handler
func (e *ElementSVGSvg) OnChange(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("change", fn),
	)
	return e
}

// click event handler
func (e *ElementSVGSvg) OnClick(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("click", fn),
	)
	return e
}

// close event handler
func (e *ElementSVGSvg) OnClose(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("close", fn),
	)
	return e
}

// contextlost event handler
func (e *ElementSVGSvg) OnContextlost(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("contextlost", fn),
	)
	return e
}

// contextmenu event handler
func (e *ElementSVGSvg) OnContextmenu(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("contextmenu", fn),
	)
	return e
}

// contextrestored event handler
func (e *ElementSVGSvg) OnContextrestored(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("contextrestored", fn),
	)
	return e
}

// copy event handler
func (e *ElementSVGSvg) OnCopy(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("copy", fn),
	)
	return e
}

// cuechange event handler
func (e *ElementSVGSvg) OnCuechange(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("cuechange", fn),
	)
	return e
}

// cut event handler
func (e *ElementSVGSvg) OnCut(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("cut", fn),
	)
	return e
}

// dblclick event handler
func (e *ElementSVGSvg) OnDblclick(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("dblclick", fn),
	)
	return e
}

// drag event handler
func (e *ElementSVGSvg) OnDrag(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("drag", fn),
	)
	return e
}

// dragend event handler
func (e *ElementSVGSvg) OnDragend(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("dragend", fn),
	)
	return e
}

// dragenter event handler
func (e *ElementSVGSvg) OnDragenter(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("dragenter", fn),
	)
	return e
}

// dragleave event handler
func (e *ElementSVGSvg) OnDragleave(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("dragleave", fn),
	)
	return e
}

// dragover event handler
func (e *ElementSVGSvg) OnDragover(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("dragover", fn),
	)
	return e
}

// dragstart event handler
func (e *ElementSVGSvg) OnDragstart(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("dragstart", fn),
	)
	return e
}

// drop event handler
func (e *ElementSVGSvg) OnDrop(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("drop", fn),
	)
	return e
}

// durationchange event handler
func (e *ElementSVGSvg) OnDurationchange(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("durationchange", fn),
	)
	return e
}

// emptied event handler
func (e *ElementSVGSvg) OnEmptied(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("emptied", fn),
	)
	return e
}

// ended event handler
func (e *ElementSVGSvg) OnEnded(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("ended", fn),
	)
	return e
}

// error event handler
func (e *ElementSVGSvg) OnError(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("error", fn),
	)
	return e
}

// focus event handler
func (e *ElementSVGSvg) OnFocus(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("focus", fn),
	)
	return e
}

// formdata event handler
func (e *ElementSVGSvg) OnFormdata(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("formdata", fn),
	)
	return e
}

// input event handler
func (e *ElementSVGSvg) OnInput(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("input", fn),
	)
	return e
}

// invalid event handler
func (e *ElementSVGSvg) OnInvalid(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("invalid", fn),
	)
	return e
}

// keydown event handler
func (e *ElementSVGSvg) OnKeydown(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("keydown", fn),
	)
	return e
}

// keypress event handler
func (e *ElementSVGSvg) OnKeypress(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("keypress", fn),
	)
	return e
}

// keyup event handler
func (e *ElementSVGSvg) OnKeyup(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("keyup", fn),
	)
	return e
}

// load event handler
func (e *ElementSVGSvg) OnLoad(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("load", fn),
	)
	return e
}

// loadeddata event handler
func (e *ElementSVGSvg) OnLoadeddata(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("loadeddata", fn),
	)
	return e
}

// loadedmetadata event handler
func (e *ElementSVGSvg) OnLoadedmetadata(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("loadedmetadata", fn),
	)
	return e
}

// loadstart event handler
func (e *ElementSVGSvg) OnLoadstart(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("loadstart", fn),
	)
	return e
}

// mousedown event handler
func (e *ElementSVGSvg) OnMousedown(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("mousedown", fn),
	)
	return e
}

// mouseenter event handler
func (e *ElementSVGSvg) OnMouseenter(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("mouseenter", fn),
	)
	return e
}

// mouseleave event handler
func (e *ElementSVGSvg) OnMouseleave(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("mouseleave", fn),
	)
	return e
}

// mousemove event handler
func (e *ElementSVGSvg) OnMousemove(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("mousemove", fn),
	)
	return e
}

// mouseout event handler
func (e *ElementSVGSvg) OnMouseout(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("mouseout", fn),
	)
	return e
}

// mouseover event handler
func (e *ElementSVGSvg) OnMouseover(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("mouseover", fn),
	)
	return e
}

// mouseup event handler
func (e *ElementSVGSvg) OnMouseup(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("mouseup", fn),
	)
	return e
}

// paste event handler
func (e *ElementSVGSvg) OnPaste(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("paste", fn),
	)
	return e
}

// pause event handler
func (e *ElementSVGSvg) OnPause(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("pause", fn),
	)
	return e
}

// play event handler
func (e *ElementSVGSvg) OnPlay(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("play", fn),
	)
	return e
}

// playing event handler
func (e *ElementSVGSvg) OnPlaying(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("playing", fn),
	)
	return e
}

// progress event handler
func (e *ElementSVGSvg) OnProgress(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("progress", fn),
	)
	return e
}

// ratechange event handler
func (e *ElementSVGSvg) OnRatechange(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("ratechange", fn),
	)
	return e
}

// reset event handler
func (e *ElementSVGSvg) OnReset(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("reset", fn),
	)
	return e
}

// resize event handler
func (e *ElementSVGSvg) OnResize(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("resize", fn),
	)
	return e
}

// scroll event handler
func (e *ElementSVGSvg) OnScroll(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("scroll", fn),
	)
	return e
}

// scrollend event handler
func (e *ElementSVGSvg) OnScrollend(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("scrollend", fn),
	)
	return e
}

// securitypolicyviolation event handler
func (e *ElementSVGSvg) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("securitypolicyviolation", fn),
	)
	return e
}

// seeked event handler
func (e *ElementSVGSvg) OnSeeked(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("seeked", fn),
	)
	return e
}

// seeking event handler
func (e *ElementSVGSvg) OnSeeking(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("seeking", fn),
	)
	return e
}

// select event handler
func (e *ElementSVGSvg) OnSelect(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("select", fn),
	)
	return e
}

// slotchange event handler
func (e *ElementSVGSvg) OnSlotchange(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("slotchange", fn),
	)
	return e
}

// stalled event handler
func (e *ElementSVGSvg) OnStalled(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("stalled", fn),
	)
	return e
}

// submit event handler
func (e *ElementSVGSvg) OnSubmit(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("submit", fn),
	)
	return e
}

// suspend event handler
func (e *ElementSVGSvg) OnSuspend(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("suspend", fn),
	)
	return e
}

// timeupdate event handler
func (e *ElementSVGSvg) OnTimeupdate(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("timeupdate", fn),
	)
	return e
}

// toggle event handler
func (e *ElementSVGSvg) OnToggle(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("toggle", fn),
	)
	return e
}

// volumechange event handler
func (e *ElementSVGSvg) OnVolumechange(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("volumechange", fn),
	)
	return e
}

// waiting event handler
func (e *ElementSVGSvg) OnWaiting(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("waiting", fn),
	)
	return e
}

// wheel event handler
func (e *ElementSVGSvg) OnWheel(fn engine.EventHandler) *ElementSVGSvg {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("wheel", fn),
	)
	return e
}
