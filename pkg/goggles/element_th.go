/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementTh struct {
	*baseElement
}

func TH(children ...any) *ElementTh {
	return &ElementTh{
		baseElement: newBaseElement("th", children...),
	}
}

func (e *ElementTh) Add(children ...any) *ElementTh {
	e.baseElement.add(children...)
	return e
}

func (e *ElementTh) Custom(k, v string, dontEscape ...bool) *ElementTh {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementTh) BindCustom(k string, v string, dontEscape ...bool) *ElementTh {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementTh) appendAttribute(k string, v string, dontEscape ...bool) *ElementTh {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementTh) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementTh) Dir(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementTh) Lang(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTh) Itemscope(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementTh) Itemprop(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementTh) Itemref(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementTh) Style(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Abbr is the "abbr" attribute.
// Alternative label to use for the header cell when referencing the cell in other contexts
// Valid values are constrained to the following:
//   - text
func (element *ElementTh) Abbr(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("abbr", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementTh) Draggable(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementTh) Nonce(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Rowspan is the "rowspan" attribute.
// Number of rows that the cell is to span
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementTh) Rowspan(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("rowspan", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementTh) Slot(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementTh) Title(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementTh) Translate(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Colspan is the "colspan" attribute.
// Number of columns that the cell is to span
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementTh) Colspan(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("colspan", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTh) Inert(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementTh) Contenteditable(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementTh) Class(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Headers is the "headers" attribute.
// The header cells for this cell
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementTh) Headers(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("headers", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementTh) Hidden(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementTh) Itemtype(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Scope is the "scope" attribute.
// Specifies which cells the header cell applies to
// Valid values are constrained to the following:
//   - row
//   - row
//   - col
//   - col
//   - rowgroup
//   - rowgroup
//   - colgroup
//   - colgroup
func (element *ElementTh) Scope(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("scope", v, dontEscape...)
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
func (element *ElementTh) Autocapitalize(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTh) Autofocus(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementTh) Accesskey(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementTh) Tabindex(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("tabindex", v, dontEscape...)
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
func (element *ElementTh) Inputmode(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementTh) Is(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementTh) Itemid(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementTh) Popover(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementTh) Spellcheck(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("spellcheck", v, dontEscape...)
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
func (element *ElementTh) Enterkeyhint(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementTh) Id(v string, dontEscape ...bool) *ElementTh {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementTh) OnAuxclick(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnBeforematch(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnBeforetoggle(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnBlur(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnCancel(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnCanplay(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnCanplaythrough(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnChange(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnClick(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnClose(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnContextlost(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnContextmenu(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnContextrestored(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnCopy(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnCuechange(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnCut(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDblclick(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDrag(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDragend(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDragenter(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDragleave(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDragover(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDragstart(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDrop(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnDurationchange(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnEmptied(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnEnded(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnError(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnFocus(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnFormdata(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnInput(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnInvalid(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnKeydown(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnKeypress(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnKeyup(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnLoad(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnLoadeddata(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnLoadedmetadata(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnLoadstart(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnMousedown(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnMouseenter(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnMouseleave(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnMousemove(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnMouseout(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnMouseover(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnMouseup(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnPaste(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnPause(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnPlay(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnPlaying(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnProgress(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnRatechange(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnReset(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnResize(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnScroll(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnScrollend(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnSeeked(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnSeeking(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnSelect(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnSlotchange(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnStalled(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnSubmit(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnSuspend(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnTimeupdate(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnToggle(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnVolumechange(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnWaiting(fn engine.EventHandler) *ElementTh {
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
func (e *ElementTh) OnWheel(fn engine.EventHandler) *ElementTh {
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
