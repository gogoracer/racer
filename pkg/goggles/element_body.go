/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementBody struct {
	*baseElement
}

func BODY(children ...any) *ElementBody {
	return &ElementBody{
		baseElement: newBaseElement("body", children...),
	}
}

func (e *ElementBody) Add(children ...any) *ElementBody {
	e.baseElement.add(children...)
	return e
}

func (e *ElementBody) Custom(k, v string, dontEscape ...bool) *ElementBody {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementBody) BindCustom(k string, v string, dontEscape ...bool) *ElementBody {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementBody) appendAttribute(k string, v string, dontEscape ...bool) *ElementBody {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementBody) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementBody) Style(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementBody) Class(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementBody) Itemid(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementBody) Spellcheck(v string, dontEscape ...bool) *ElementBody {
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
func (element *ElementBody) Inputmode(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementBody) Itemtype(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementBody) Lang(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementBody) Accesskey(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementBody) Autofocus(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementBody) Is(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementBody) Nonce(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementBody) Translate(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementBody) Itemprop(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementBody) Dir(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementBody) Hidden(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementBody) Inert(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementBody) Itemscope(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementBody) Tabindex(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementBody) Popover(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementBody) Slot(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementBody) Title(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementBody) Itemref(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("itemref", v, dontEscape...)
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
func (element *ElementBody) Autocapitalize(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementBody) Contenteditable(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementBody) Draggable(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("draggable", v, dontEscape...)
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
func (element *ElementBody) Enterkeyhint(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementBody) Id(v string, dontEscape ...bool) *ElementBody {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementBody) OnAuxclick(fn engine.EventHandler) *ElementBody {
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

// afterprint event handler for Window object
func (e *ElementBody) OnAfterprint(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("afterprint", fn),
	)
	return e
}

// beforematch event handler
func (e *ElementBody) OnBeforematch(fn engine.EventHandler) *ElementBody {
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

// beforeprint event handler for Window object
func (e *ElementBody) OnBeforeprint(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("beforeprint", fn),
	)
	return e
}

// beforeunload event handler for Window object
func (e *ElementBody) OnBeforeunload(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("beforeunload", fn),
	)
	return e
}

// beforetoggle event handler
func (e *ElementBody) OnBeforetoggle(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnBlur(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnCancel(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnCanplay(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnCanplaythrough(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnChange(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnClick(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnClose(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnContextlost(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnContextmenu(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnContextrestored(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnCopy(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnCuechange(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnCut(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDblclick(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDrag(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDragend(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDragenter(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDragleave(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDragover(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDragstart(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDrop(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnDurationchange(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnEmptied(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnEnded(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnError(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnFocus(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnFormdata(fn engine.EventHandler) *ElementBody {
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

// hashchange event handler for Window object
func (e *ElementBody) OnHashchange(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("hashchange", fn),
	)
	return e
}

// input event handler
func (e *ElementBody) OnInput(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnInvalid(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnKeydown(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnKeypress(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnKeyup(fn engine.EventHandler) *ElementBody {
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

// languagechange event handler for Window object
func (e *ElementBody) OnLanguagechange(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("languagechange", fn),
	)
	return e
}

// load event handler
func (e *ElementBody) OnLoad(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnLoadeddata(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnLoadedmetadata(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnLoadstart(fn engine.EventHandler) *ElementBody {
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

// message event handler for Window object
func (e *ElementBody) OnMessage(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("message", fn),
	)
	return e
}

// messageerror event handler for Window object
func (e *ElementBody) OnMessageerror(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("messageerror", fn),
	)
	return e
}

// mousedown event handler
func (e *ElementBody) OnMousedown(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnMouseenter(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnMouseleave(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnMousemove(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnMouseout(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnMouseover(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnMouseup(fn engine.EventHandler) *ElementBody {
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

// offline event handler for Window object
func (e *ElementBody) OnOffline(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("offline", fn),
	)
	return e
}

// online event handler for Window object
func (e *ElementBody) OnOnline(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("online", fn),
	)
	return e
}

// pagehide event handler for Window object
func (e *ElementBody) OnPagehide(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("pagehide", fn),
	)
	return e
}

// pageshow event handler for Window object
func (e *ElementBody) OnPageshow(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("pageshow", fn),
	)
	return e
}

// paste event handler
func (e *ElementBody) OnPaste(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnPause(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnPlay(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnPlaying(fn engine.EventHandler) *ElementBody {
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

// popstate event handler for Window object
func (e *ElementBody) OnPopstate(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("popstate", fn),
	)
	return e
}

// progress event handler
func (e *ElementBody) OnProgress(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnRatechange(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnReset(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnResize(fn engine.EventHandler) *ElementBody {
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

// rejectionhandled event handler for Window object
func (e *ElementBody) OnRejectionhandled(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("rejectionhandled", fn),
	)
	return e
}

// scroll event handler
func (e *ElementBody) OnScroll(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnScrollend(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnSeeked(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnSeeking(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnSelect(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnSlotchange(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnStalled(fn engine.EventHandler) *ElementBody {
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

// storage event handler for Window object
func (e *ElementBody) OnStorage(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("storage", fn),
	)
	return e
}

// submit event handler
func (e *ElementBody) OnSubmit(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnSuspend(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnTimeupdate(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnToggle(fn engine.EventHandler) *ElementBody {
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

// unhandledrejection event handler for Window object
func (e *ElementBody) OnUnhandledrejection(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("unhandledrejection", fn),
	)
	return e
}

// unload event handler for Window object
func (e *ElementBody) OnUnload(fn engine.EventHandler) *ElementBody {
	if fn == nil {
		return e
	}

	e.shouldBeComponent = true
	e.children = append(
		e.children,
		engine.On("unload", fn),
	)
	return e
}

// volumechange event handler
func (e *ElementBody) OnVolumechange(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnWaiting(fn engine.EventHandler) *ElementBody {
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
func (e *ElementBody) OnWheel(fn engine.EventHandler) *ElementBody {
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
