/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementPicture struct {
	*baseElement
}

func PICTURE(children ...any) *ElementPicture {
	return &ElementPicture{
		baseElement: newBaseElement("picture", children...),
	}
}

func (e *ElementPicture) Add(children ...any) *ElementPicture {
	e.baseElement.add(children...)
	return e
}

func (e *ElementPicture) Custom(k, v string, dontEscape ...bool) *ElementPicture {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementPicture) BindCustom(k string, v string, dontEscape ...bool) *ElementPicture {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementPicture) appendAttribute(k string, v string, dontEscape ...bool) *ElementPicture {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementPicture) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementPicture) Nonce(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementPicture) Popover(v string, dontEscape ...bool) *ElementPicture {
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
func (element *ElementPicture) Autocapitalize(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementPicture) Tabindex(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementPicture) Translate(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementPicture) Hidden(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementPicture) Itemid(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementPicture) Dir(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Height is the "height" attribute.
// Vertical dimension
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementPicture) Height(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("height", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementPicture) Itemprop(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementPicture) Autofocus(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementPicture) Contenteditable(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("contenteditable", v, dontEscape...)
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
func (element *ElementPicture) Inputmode(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementPicture) Itemscope(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementPicture) Slot(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Width is the "width" attribute.
// Horizontal dimension
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementPicture) Width(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("width", v, dontEscape...)
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
func (element *ElementPicture) Enterkeyhint(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementPicture) Id(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementPicture) Draggable(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementPicture) Is(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementPicture) Itemref(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementPicture) Itemtype(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementPicture) Lang(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Media is the "media" attribute.
// Applicable media
// Valid values are constrained to the following:
//   - valid_media_query_list
func (element *ElementPicture) Media(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("media", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementPicture) Accesskey(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementPicture) Class(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementPicture) Spellcheck(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("spellcheck", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementPicture) Style(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementPicture) Inert(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementPicture) Title(v string, dontEscape ...bool) *ElementPicture {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementPicture) OnAuxclick(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnBeforematch(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnBeforetoggle(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnBlur(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnCancel(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnCanplay(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnCanplaythrough(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnChange(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnClick(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnClose(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnContextlost(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnContextmenu(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnContextrestored(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnCopy(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnCuechange(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnCut(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDblclick(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDrag(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDragend(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDragenter(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDragleave(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDragover(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDragstart(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDrop(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnDurationchange(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnEmptied(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnEnded(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnError(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnFocus(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnFormdata(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnInput(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnInvalid(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnKeydown(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnKeypress(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnKeyup(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnLoad(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnLoadeddata(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnLoadedmetadata(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnLoadstart(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnMousedown(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnMouseenter(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnMouseleave(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnMousemove(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnMouseout(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnMouseover(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnMouseup(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnPaste(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnPause(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnPlay(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnPlaying(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnProgress(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnRatechange(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnReset(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnResize(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnScroll(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnScrollend(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnSeeked(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnSeeking(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnSelect(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnSlotchange(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnStalled(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnSubmit(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnSuspend(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnTimeupdate(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnToggle(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnVolumechange(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnWaiting(fn engine.EventHandler) *ElementPicture {
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
func (e *ElementPicture) OnWheel(fn engine.EventHandler) *ElementPicture {
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
