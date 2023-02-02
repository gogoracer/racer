/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementAbbr struct {
	*baseElement
}

func ABBR(children ...any) *ElementAbbr {
	return &ElementAbbr{
		baseElement: newBaseElement("abbr", children...),
	}
}

func (e *ElementAbbr) Add(children ...any) *ElementAbbr {
	e.baseElement.add(children...)
	return e
}

func (e *ElementAbbr) Custom(k, v string, dontEscape ...bool) *ElementAbbr {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementAbbr) BindCustom(k string, v string, dontEscape ...bool) *ElementAbbr {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementAbbr) appendAttribute(k string, v string, dontEscape ...bool) *ElementAbbr {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementAbbr) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementAbbr) Autofocus(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementAbbr) Draggable(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementAbbr) Id(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("id", v, dontEscape...)
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
func (element *ElementAbbr) Inputmode(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementAbbr) Itemid(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementAbbr) Itemprop(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementAbbr) Lang(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementAbbr) Translate(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementAbbr) Tabindex(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("tabindex", v, dontEscape...)
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
func (element *ElementAbbr) Autocapitalize(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementAbbr) Itemscope(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementAbbr) Itemtype(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementAbbr) Nonce(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementAbbr) Accesskey(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementAbbr) Contenteditable(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementAbbr) Inert(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementAbbr) Popover(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementAbbr) Spellcheck(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("spellcheck", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementAbbr) Itemref(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementAbbr) Class(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementAbbr) Dir(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementAbbr) Slot(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementAbbr) Title(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("title", v, dontEscape...)
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
func (element *ElementAbbr) Enterkeyhint(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementAbbr) Hidden(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementAbbr) Is(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementAbbr) Style(v string, dontEscape ...bool) *ElementAbbr {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementAbbr) OnAuxclick(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnBeforematch(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnBeforetoggle(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnBlur(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnCancel(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnCanplay(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnCanplaythrough(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnChange(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnClick(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnClose(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnContextlost(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnContextmenu(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnContextrestored(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnCopy(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnCuechange(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnCut(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDblclick(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDrag(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDragend(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDragenter(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDragleave(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDragover(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDragstart(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDrop(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnDurationchange(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnEmptied(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnEnded(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnError(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnFocus(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnFormdata(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnInput(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnInvalid(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnKeydown(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnKeypress(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnKeyup(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnLoad(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnLoadeddata(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnLoadedmetadata(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnLoadstart(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnMousedown(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnMouseenter(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnMouseleave(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnMousemove(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnMouseout(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnMouseover(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnMouseup(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnPaste(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnPause(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnPlay(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnPlaying(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnProgress(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnRatechange(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnReset(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnResize(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnScroll(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnScrollend(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnSeeked(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnSeeking(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnSelect(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnSlotchange(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnStalled(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnSubmit(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnSuspend(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnTimeupdate(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnToggle(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnVolumechange(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnWaiting(fn engine.EventHandler) *ElementAbbr {
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
func (e *ElementAbbr) OnWheel(fn engine.EventHandler) *ElementAbbr {
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
