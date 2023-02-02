/* cSpell:disable */

package handlebars

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementFigcaption struct {
	*baseElement
}

func FIGCAPTION(children ...any) *ElementFigcaption {
	return &ElementFigcaption{
		baseElement: newBaseElement("figcaption", children...),
	}
}

func (e *ElementFigcaption) Add(children ...any) *ElementFigcaption {
	e.baseElement.add(children...)
	return e
}

func (e *ElementFigcaption) Custom(k, v string, dontEscape ...bool) *ElementFigcaption {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementFigcaption) BindCustom(k string, v string, dontEscape ...bool) *ElementFigcaption {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementFigcaption) setAttribute(k string, v string, dontEscape ...bool) *ElementFigcaption {
	e.baseElement.setAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementFigcaption) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementFigcaption) Draggable(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("draggable", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementFigcaption) Itemid(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("itemid", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementFigcaption) Popover(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("popover", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementFigcaption) Accesskey(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("accesskey", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementFigcaption) Dir(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("dir", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementFigcaption) Hidden(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("hidden", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementFigcaption) Slot(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("slot", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementFigcaption) Itemprop(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("itemprop", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementFigcaption) Itemref(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("itemref", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementFigcaption) Style(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("style", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementFigcaption) Translate(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("translate", v, dontEscape...)
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
func (element *ElementFigcaption) Autocapitalize(v string, dontEscape ...bool) *ElementFigcaption {
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
func (element *ElementFigcaption) Enterkeyhint(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementFigcaption) Inert(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("inert", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementFigcaption) Lang(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("lang", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementFigcaption) Title(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("title", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementFigcaption) Class(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("class", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementFigcaption) Is(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("is", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementFigcaption) Nonce(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("nonce", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementFigcaption) Autofocus(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("autofocus", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementFigcaption) Contenteditable(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("contenteditable", v, dontEscape...)
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
func (element *ElementFigcaption) Inputmode(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementFigcaption) Itemtype(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("itemtype", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementFigcaption) Tabindex(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("tabindex", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementFigcaption) Itemscope(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("itemscope", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementFigcaption) Id(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("id", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementFigcaption) Spellcheck(v string, dontEscape ...bool) *ElementFigcaption {
	element.setAttribute("spellcheck", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementFigcaption) OnAuxclick(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnBeforematch(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnBeforetoggle(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnBlur(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnCancel(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnCanplay(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnCanplaythrough(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnChange(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnClick(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnClose(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnContextlost(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnContextmenu(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnContextrestored(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnCopy(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnCuechange(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnCut(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDblclick(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDrag(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDragend(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDragenter(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDragleave(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDragover(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDragstart(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDrop(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnDurationchange(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnEmptied(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnEnded(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnError(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnFocus(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnFormdata(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnInput(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnInvalid(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnKeydown(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnKeypress(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnKeyup(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnLoad(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnLoadeddata(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnLoadedmetadata(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnLoadstart(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnMousedown(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnMouseenter(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnMouseleave(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnMousemove(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnMouseout(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnMouseover(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnMouseup(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnPaste(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnPause(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnPlay(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnPlaying(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnProgress(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnRatechange(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnReset(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnResize(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnScroll(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnScrollend(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnSeeked(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnSeeking(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnSelect(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnSlotchange(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnStalled(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnSubmit(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnSuspend(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnTimeupdate(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnToggle(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnVolumechange(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnWaiting(fn engine.EventHandler) *ElementFigcaption {
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
func (e *ElementFigcaption) OnWheel(fn engine.EventHandler) *ElementFigcaption {
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
