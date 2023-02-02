/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementOutput struct {
	*baseElement
}

func OUTPUT(children ...any) *ElementOutput {
	return &ElementOutput{
		baseElement: newBaseElement("output", children...),
	}
}

func (e *ElementOutput) Add(children ...any) *ElementOutput {
	e.baseElement.add(children...)
	return e
}

func (e *ElementOutput) Custom(k, v string, dontEscape ...bool) *ElementOutput {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementOutput) BindCustom(k string, v string, dontEscape ...bool) *ElementOutput {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementOutput) appendAttribute(k string, v string, dontEscape ...bool) *ElementOutput {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementOutput) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementOutput) Autofocus(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementOutput) Class(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Form is the "form" attribute.
// Associates the element with a form element
// Valid values are constrained to the following:
//   - id
func (element *ElementOutput) Form(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("form", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementOutput) Id(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementOutput) Inert(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// For is the "for" attribute.
// Specifies controls from which the output was calculated
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementOutput) For(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("for", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementOutput) Is(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementOutput) Popover(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("popover", v, dontEscape...)
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
func (element *ElementOutput) Enterkeyhint(v string, dontEscape ...bool) *ElementOutput {
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
func (element *ElementOutput) Hidden(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementOutput) Itemscope(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementOutput) Slot(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementOutput) Accesskey(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementOutput) Contenteditable(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementOutput) Draggable(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementOutput) Itemref(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementOutput) Itemtype(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementOutput) Spellcheck(v string, dontEscape ...bool) *ElementOutput {
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
func (element *ElementOutput) Inputmode(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementOutput) Itemprop(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("itemprop", v, dontEscape...)
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
func (element *ElementOutput) Autocapitalize(v string, dontEscape ...bool) *ElementOutput {
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
func (element *ElementOutput) Dir(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementOutput) Itemid(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementOutput) Translate(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementOutput) Lang(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   - text
func (element *ElementOutput) Name(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("name", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementOutput) Nonce(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementOutput) Tabindex(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementOutput) Title(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementOutput) Style(v string, dontEscape ...bool) *ElementOutput {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementOutput) OnAuxclick(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnBeforematch(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnBeforetoggle(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnBlur(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnCancel(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnCanplay(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnCanplaythrough(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnChange(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnClick(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnClose(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnContextlost(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnContextmenu(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnContextrestored(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnCopy(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnCuechange(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnCut(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDblclick(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDrag(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDragend(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDragenter(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDragleave(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDragover(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDragstart(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDrop(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnDurationchange(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnEmptied(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnEnded(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnError(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnFocus(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnFormdata(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnInput(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnInvalid(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnKeydown(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnKeypress(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnKeyup(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnLoad(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnLoadeddata(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnLoadedmetadata(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnLoadstart(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnMousedown(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnMouseenter(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnMouseleave(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnMousemove(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnMouseout(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnMouseover(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnMouseup(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnPaste(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnPause(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnPlay(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnPlaying(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnProgress(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnRatechange(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnReset(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnResize(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnScroll(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnScrollend(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnSeeked(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnSeeking(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnSelect(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnSlotchange(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnStalled(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnSubmit(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnSuspend(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnTimeupdate(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnToggle(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnVolumechange(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnWaiting(fn engine.EventHandler) *ElementOutput {
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
func (e *ElementOutput) OnWheel(fn engine.EventHandler) *ElementOutput {
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
