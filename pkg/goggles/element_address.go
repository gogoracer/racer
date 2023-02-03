/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementAddress struct {
	*baseElement
}

func ADDRESS(children ...any) *ElementAddress {
	return &ElementAddress{
		baseElement: newBaseElement("address", children...),
	}
}

func (e *ElementAddress) Add(children ...any) *ElementAddress {
	e.baseElement.add(children...)
	return e
}

func (e *ElementAddress) Custom(k, v string, dontEscape ...bool) *ElementAddress {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementAddress) BindCustom(k string, v string, dontEscape ...bool) *ElementAddress {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementAddress) appendAttribute(k string, v string, dontEscape ...bool) *ElementAddress {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementAddress) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementAddress) Draggable(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementAddress) Id(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementAddress) Itemprop(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementAddress) Itemscope(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementAddress) Itemtype(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementAddress) Nonce(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementAddress) Tabindex(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementAddress) Accesskey(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementAddress) Autofocus(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementAddress) Inert(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("inert", v, dontEscape...)
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
func (element *ElementAddress) Inputmode(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("inputmode", v, dontEscape...)
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
func (element *ElementAddress) Enterkeyhint(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementAddress) Itemid(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementAddress) Lang(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementAddress) Popover(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementAddress) Slot(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementAddress) Spellcheck(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("spellcheck", v, dontEscape...)
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
func (element *ElementAddress) Autocapitalize(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementAddress) Dir(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementAddress) Itemref(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementAddress) Title(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementAddress) Hidden(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementAddress) Style(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementAddress) Class(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementAddress) Contenteditable(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementAddress) Is(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementAddress) Translate(v string, dontEscape ...bool) *ElementAddress {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementAddress) OnAuxclick(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnBeforematch(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnBeforetoggle(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnBlur(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnCancel(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnCanplay(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnCanplaythrough(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnChange(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnClick(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnClose(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnContextlost(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnContextmenu(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnContextrestored(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnCopy(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnCuechange(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnCut(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDblclick(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDrag(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDragend(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDragenter(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDragleave(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDragover(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDragstart(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDrop(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnDurationchange(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnEmptied(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnEnded(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnError(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnFocus(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnFormdata(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnInput(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnInvalid(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnKeydown(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnKeypress(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnKeyup(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnLoad(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnLoadeddata(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnLoadedmetadata(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnLoadstart(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnMousedown(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnMouseenter(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnMouseleave(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnMousemove(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnMouseout(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnMouseover(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnMouseup(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnPaste(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnPause(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnPlay(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnPlaying(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnProgress(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnRatechange(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnReset(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnResize(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnScroll(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnScrollend(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnSeeked(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnSeeking(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnSelect(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnSlotchange(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnStalled(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnSubmit(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnSuspend(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnTimeupdate(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnToggle(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnVolumechange(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnWaiting(fn engine.EventHandler) *ElementAddress {
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
func (e *ElementAddress) OnWheel(fn engine.EventHandler) *ElementAddress {
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
