/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementMeta struct {
	*baseElement
}

func META(children ...any) *ElementMeta {
	return &ElementMeta{
		baseElement: newBaseElement("meta", children...),
	}
}

func (e *ElementMeta) Add(children ...any) *ElementMeta {
	e.baseElement.add(children...)
	return e
}

func (e *ElementMeta) Custom(k, v string, dontEscape ...bool) *ElementMeta {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementMeta) BindCustom(k string, v string, dontEscape ...bool) *ElementMeta {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementMeta) appendAttribute(k string, v string, dontEscape ...bool) *ElementMeta {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementMeta) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementMeta) Itemprop(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementMeta) Slot(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementMeta) Id(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementMeta) Style(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementMeta) Accesskey(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMeta) Contenteditable(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementMeta) Hidden(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMeta) Itemscope(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMeta) Spellcheck(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("spellcheck", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementMeta) Class(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("class", v, dontEscape...)
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
func (element *ElementMeta) Inputmode(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementMeta) Itemref(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Media is the "media" attribute.
// Applicable media
// Valid values are constrained to the following:
//   - valid_media_query_list
func (element *ElementMeta) Media(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("media", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMeta) Autofocus(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Charset is the "charset" attribute.
// Character encoding declaration
// Valid values are constrained to the following:
//   - utf_8
func (element *ElementMeta) Charset(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("charset", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementMeta) Draggable(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   - text
func (element *ElementMeta) Name(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("name", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementMeta) Popover(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementMeta) Title(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementMeta) Translate(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// HttpEquiv is the "http-equiv" attribute.
// Pragma directive
// Valid values are constrained to the following:
//   - content_type
//   - content_type
//   - default_style
//   - default_style
//   - refresh
//   - refresh
//   - x_ua_compatible
//   - x_ua_compatible
//   - content_security_policy
//   - content_security_policy
func (element *ElementMeta) HttpEquiv(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("http-equiv", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementMeta) Is(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementMeta) Itemtype(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementMeta) Nonce(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("nonce", v, dontEscape...)
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
func (element *ElementMeta) Autocapitalize(v string, dontEscape ...bool) *ElementMeta {
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
func (element *ElementMeta) Enterkeyhint(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementMeta) Lang(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementMeta) Tabindex(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Content is the "content" attribute.
// Value of the element
// Valid values are constrained to the following:
//   - text
func (element *ElementMeta) Content(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("content", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementMeta) Dir(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementMeta) Inert(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementMeta) Itemid(v string, dontEscape ...bool) *ElementMeta {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementMeta) OnAuxclick(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnBeforematch(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnBeforetoggle(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnBlur(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnCancel(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnCanplay(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnCanplaythrough(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnChange(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnClick(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnClose(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnContextlost(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnContextmenu(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnContextrestored(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnCopy(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnCuechange(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnCut(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDblclick(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDrag(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDragend(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDragenter(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDragleave(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDragover(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDragstart(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDrop(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnDurationchange(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnEmptied(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnEnded(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnError(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnFocus(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnFormdata(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnInput(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnInvalid(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnKeydown(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnKeypress(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnKeyup(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnLoad(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnLoadeddata(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnLoadedmetadata(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnLoadstart(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnMousedown(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnMouseenter(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnMouseleave(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnMousemove(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnMouseout(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnMouseover(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnMouseup(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnPaste(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnPause(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnPlay(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnPlaying(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnProgress(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnRatechange(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnReset(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnResize(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnScroll(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnScrollend(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnSeeked(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnSeeking(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnSelect(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnSlotchange(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnStalled(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnSubmit(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnSuspend(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnTimeupdate(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnToggle(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnVolumechange(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnWaiting(fn engine.EventHandler) *ElementMeta {
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
func (e *ElementMeta) OnWheel(fn engine.EventHandler) *ElementMeta {
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