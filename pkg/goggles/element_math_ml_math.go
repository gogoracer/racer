/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementMathMLMath struct {
	*baseElement
}

func MATHMLMATH(children ...any) *ElementMathMLMath {
	return &ElementMathMLMath{
		baseElement: newBaseElement("MathML math", children...),
	}
}

func (e *ElementMathMLMath) Add(children ...any) *ElementMathMLMath {
	e.baseElement.add(children...)
	return e
}

func (e *ElementMathMLMath) Custom(k, v string, dontEscape ...bool) *ElementMathMLMath {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementMathMLMath) BindCustom(k string, v string, dontEscape ...bool) *ElementMathMLMath {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementMathMLMath) appendAttribute(k string, v string, dontEscape ...bool) *ElementMathMLMath {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementMathMLMath) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementMathMLMath) Hidden(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("hidden", v, dontEscape...)
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
func (element *ElementMathMLMath) Inputmode(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementMathMLMath) Itemid(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMathMLMath) Contenteditable(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMathMLMath) Itemscope(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("itemscope", v, dontEscape...)
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
func (element *ElementMathMLMath) Autocapitalize(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMathMLMath) Draggable(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementMathMLMath) Popover(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementMathMLMath) Tabindex(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementMathMLMath) Class(v string, dontEscape ...bool) *ElementMathMLMath {
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
func (element *ElementMathMLMath) Dir(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementMathMLMath) Id(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMathMLMath) Inert(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementMathMLMath) Is(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementMathMLMath) Style(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementMathMLMath) Lang(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMathMLMath) Spellcheck(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("spellcheck", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementMathMLMath) Title(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementMathMLMath) Accesskey(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("accesskey", v, dontEscape...)
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
func (element *ElementMathMLMath) Enterkeyhint(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementMathMLMath) Itemtype(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementMathMLMath) Slot(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementMathMLMath) Translate(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMathMLMath) Autofocus(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementMathMLMath) Itemprop(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementMathMLMath) Itemref(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementMathMLMath) Nonce(v string, dontEscape ...bool) *ElementMathMLMath {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementMathMLMath) OnAuxclick(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnBeforematch(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnBeforetoggle(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnBlur(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnCancel(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnCanplay(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnCanplaythrough(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnChange(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnClick(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnClose(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnContextlost(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnContextmenu(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnContextrestored(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnCopy(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnCuechange(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnCut(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDblclick(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDrag(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDragend(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDragenter(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDragleave(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDragover(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDragstart(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDrop(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnDurationchange(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnEmptied(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnEnded(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnError(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnFocus(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnFormdata(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnInput(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnInvalid(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnKeydown(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnKeypress(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnKeyup(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnLoad(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnLoadeddata(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnLoadedmetadata(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnLoadstart(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnMousedown(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnMouseenter(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnMouseleave(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnMousemove(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnMouseout(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnMouseover(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnMouseup(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnPaste(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnPause(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnPlay(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnPlaying(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnProgress(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnRatechange(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnReset(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnResize(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnScroll(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnScrollend(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnSeeked(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnSeeking(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnSelect(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnSlotchange(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnStalled(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnSubmit(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnSuspend(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnTimeupdate(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnToggle(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnVolumechange(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnWaiting(fn engine.EventHandler) *ElementMathMLMath {
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
func (e *ElementMathMLMath) OnWheel(fn engine.EventHandler) *ElementMathMLMath {
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
