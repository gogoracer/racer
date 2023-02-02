/* cSpell:disable */

package handlebars

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementWbr struct {
	*baseElement
}

func WBR(children ...any) *ElementWbr {
	return &ElementWbr{
		baseElement: newBaseElement("wbr", children...),
	}
}

func (e *ElementWbr) Add(children ...any) *ElementWbr {
	e.baseElement.add(children...)
	return e
}

func (e *ElementWbr) Custom(k, v string, dontEscape ...bool) *ElementWbr {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementWbr) BindCustom(k string, v string, dontEscape ...bool) *ElementWbr {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementWbr) setAttribute(k string, v string, dontEscape ...bool) *ElementWbr {
	e.baseElement.setAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementWbr) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementWbr) Slot(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("slot", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementWbr) Itemscope(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("itemscope", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementWbr) Itemtype(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("itemtype", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementWbr) Class(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("class", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementWbr) Dir(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("dir", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementWbr) Is(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("is", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementWbr) Hidden(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("hidden", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementWbr) Spellcheck(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("spellcheck", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementWbr) Tabindex(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("tabindex", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementWbr) Title(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("title", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementWbr) Translate(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("translate", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementWbr) Accesskey(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("accesskey", v, dontEscape...)
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
func (element *ElementWbr) Autocapitalize(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("autocapitalize", v, dontEscape...)
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
func (element *ElementWbr) Enterkeyhint(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementWbr) Nonce(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("nonce", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementWbr) Style(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("style", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementWbr) Contenteditable(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("contenteditable", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementWbr) Itemid(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("itemid", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementWbr) Lang(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("lang", v, dontEscape...)
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
func (element *ElementWbr) Inputmode(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("inputmode", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementWbr) Popover(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("popover", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementWbr) Autofocus(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("autofocus", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementWbr) Draggable(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("draggable", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementWbr) Id(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("id", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementWbr) Inert(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("inert", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementWbr) Itemprop(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("itemprop", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementWbr) Itemref(v string, dontEscape ...bool) *ElementWbr {
	element.setAttribute("itemref", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementWbr) OnAuxclick(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnBeforematch(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnBeforetoggle(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnBlur(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnCancel(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnCanplay(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnCanplaythrough(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnChange(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnClick(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnClose(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnContextlost(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnContextmenu(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnContextrestored(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnCopy(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnCuechange(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnCut(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDblclick(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDrag(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDragend(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDragenter(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDragleave(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDragover(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDragstart(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDrop(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnDurationchange(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnEmptied(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnEnded(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnError(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnFocus(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnFormdata(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnInput(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnInvalid(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnKeydown(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnKeypress(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnKeyup(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnLoad(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnLoadeddata(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnLoadedmetadata(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnLoadstart(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnMousedown(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnMouseenter(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnMouseleave(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnMousemove(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnMouseout(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnMouseover(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnMouseup(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnPaste(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnPause(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnPlay(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnPlaying(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnProgress(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnRatechange(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnReset(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnResize(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnScroll(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnScrollend(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnSeeked(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnSeeking(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnSelect(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnSlotchange(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnStalled(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnSubmit(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnSuspend(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnTimeupdate(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnToggle(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnVolumechange(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnWaiting(fn engine.EventHandler) *ElementWbr {
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
func (e *ElementWbr) OnWheel(fn engine.EventHandler) *ElementWbr {
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
