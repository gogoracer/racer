/* cSpell:disable */

package handlebars

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementMeter struct {
	*baseElement
}

func METER(children ...any) *ElementMeter {
	return &ElementMeter{
		baseElement: newBaseElement("meter", children...),
	}
}

func (e *ElementMeter) Add(children ...any) *ElementMeter {
	e.baseElement.add(children...)
	return e
}

func (e *ElementMeter) Custom(k, v string, dontEscape ...bool) *ElementMeter {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementMeter) BindCustom(k string, v string, dontEscape ...bool) *ElementMeter {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementMeter) setAttribute(k string, v string, dontEscape ...bool) *ElementMeter {
	e.baseElement.setAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementMeter) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementMeter) Hidden(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("hidden", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementMeter) Is(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("is", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMeter) Itemscope(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("itemscope", v, dontEscape...)
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
func (element *ElementMeter) Enterkeyhint(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("enterkeyhint", v, dontEscape...)
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
func (element *ElementMeter) Inputmode(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementMeter) Itemid(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("itemid", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementMeter) Nonce(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("nonce", v, dontEscape...)
	return element
}

// Optimum is the "optimum" attribute.
// Optimum value in gauge
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementMeter) Optimum(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("optimum", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementMeter) Slot(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("slot", v, dontEscape...)
	return element
}

// High is the "high" attribute.
// Low limit of high range
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementMeter) High(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("high", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMeter) Inert(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("inert", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementMeter) Id(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("id", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementMeter) Dir(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("dir", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementMeter) Itemref(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("itemref", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementMeter) Lang(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("lang", v, dontEscape...)
	return element
}

// Value is the "value" attribute.
// Current value of the element
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementMeter) Value(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("value", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMeter) Contenteditable(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("contenteditable", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementMeter) Itemprop(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("itemprop", v, dontEscape...)
	return element
}

// Low is the "low" attribute.
// High limit of low range
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementMeter) Low(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("low", v, dontEscape...)
	return element
}

// Max is the "max" attribute.
// Upper bound of range
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementMeter) Max(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("max", v, dontEscape...)
	return element
}

// Min is the "min" attribute.
// Lower bound of range
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementMeter) Min(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("min", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMeter) Draggable(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("draggable", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementMeter) Tabindex(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("tabindex", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementMeter) Popover(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("popover", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMeter) Autofocus(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("autofocus", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementMeter) Class(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("class", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementMeter) Itemtype(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("itemtype", v, dontEscape...)
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
func (element *ElementMeter) Autocapitalize(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMeter) Spellcheck(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("spellcheck", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementMeter) Style(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("style", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementMeter) Title(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("title", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementMeter) Translate(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("translate", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementMeter) Accesskey(v string, dontEscape ...bool) *ElementMeter {
	element.setAttribute("accesskey", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementMeter) OnAuxclick(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnBeforematch(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnBeforetoggle(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnBlur(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnCancel(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnCanplay(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnCanplaythrough(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnChange(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnClick(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnClose(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnContextlost(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnContextmenu(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnContextrestored(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnCopy(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnCuechange(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnCut(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDblclick(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDrag(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDragend(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDragenter(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDragleave(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDragover(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDragstart(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDrop(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnDurationchange(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnEmptied(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnEnded(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnError(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnFocus(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnFormdata(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnInput(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnInvalid(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnKeydown(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnKeypress(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnKeyup(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnLoad(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnLoadeddata(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnLoadedmetadata(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnLoadstart(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnMousedown(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnMouseenter(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnMouseleave(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnMousemove(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnMouseout(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnMouseover(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnMouseup(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnPaste(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnPause(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnPlay(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnPlaying(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnProgress(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnRatechange(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnReset(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnResize(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnScroll(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnScrollend(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnSeeked(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnSeeking(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnSelect(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnSlotchange(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnStalled(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnSubmit(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnSuspend(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnTimeupdate(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnToggle(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnVolumechange(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnWaiting(fn engine.EventHandler) *ElementMeter {
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
func (e *ElementMeter) OnWheel(fn engine.EventHandler) *ElementMeter {
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
