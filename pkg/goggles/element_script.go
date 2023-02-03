/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementScript struct {
	*baseElement
}

func SCRIPT(children ...any) *ElementScript {
	return &ElementScript{
		baseElement: newBaseElement("script", children...),
	}
}

func (e *ElementScript) Add(children ...any) *ElementScript {
	e.baseElement.add(children...)
	return e
}

func (e *ElementScript) Custom(k, v string, dontEscape ...bool) *ElementScript {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementScript) BindCustom(k string, v string, dontEscape ...bool) *ElementScript {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementScript) appendAttribute(k string, v string, dontEscape ...bool) *ElementScript {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementScript) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Integrity is the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Valid values are constrained to the following:
//   - text
func (element *ElementScript) Integrity(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("integrity", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementScript) Slot(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementScript) Translate(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementScript) Autofocus(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementScript) Is(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementScript) Itemprop(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementScript) Itemref(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Nomodule is the "nomodule" attribute.
// Prevents execution in user agents that support module scripts
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementScript) Nomodule(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("nomodule", v, dontEscape...)
	return element
}

// Src is the "src" attribute.
// Address of the resource
// Valid values are constrained to the following:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (element *ElementScript) Src(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("src", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementScript) Accesskey(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Async is the "async" attribute.
// Execute script when available, without blocking while fetching
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementScript) Async(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("async", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementScript) Popover(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementScript) Spellcheck(v string, dontEscape ...bool) *ElementScript {
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
func (element *ElementScript) Enterkeyhint(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementScript) Lang(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Referrerpolicy is the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Valid values are constrained to the following:
//   - referrer_policy
func (element *ElementScript) Referrerpolicy(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("referrerpolicy", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementScript) Tabindex(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Type is the "type" attribute.
// Type of script
// Valid values are constrained to the following:
//   - module
//   - valid_mime_type_string
//   - java_script_mime_type_essence_match
func (element *ElementScript) Type(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("type", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementScript) Dir(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementScript) Itemid(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementScript) Itemtype(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementScript) Title(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("title", v, dontEscape...)
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
func (element *ElementScript) Autocapitalize(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementScript) Class(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementScript) Contenteditable(v string, dontEscape ...bool) *ElementScript {
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
func (element *ElementScript) Hidden(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementScript) Nonce(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementScript) Style(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementScript) Draggable(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Blocking is the "blocking" attribute.
// Whether the element is potentially render-blocking
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementScript) Blocking(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("blocking", v, dontEscape...)
	return element
}

// Crossorigin is the "crossorigin" attribute.
// How the element handles crossorigin requests
// Valid values are constrained to the following:
//   - anonymous
//   - anonymous
//   - use_credentials
//   - use_credentials
func (element *ElementScript) Crossorigin(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("crossorigin", v, dontEscape...)
	return element
}

// Defer is the "defer" attribute.
// Defer script execution
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementScript) Defer(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("defer", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementScript) Id(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementScript) Inert(v string, dontEscape ...bool) *ElementScript {
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
func (element *ElementScript) Inputmode(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementScript) Itemscope(v string, dontEscape ...bool) *ElementScript {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementScript) OnAuxclick(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnBeforematch(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnBeforetoggle(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnBlur(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnCancel(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnCanplay(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnCanplaythrough(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnChange(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnClick(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnClose(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnContextlost(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnContextmenu(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnContextrestored(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnCopy(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnCuechange(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnCut(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDblclick(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDrag(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDragend(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDragenter(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDragleave(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDragover(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDragstart(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDrop(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnDurationchange(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnEmptied(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnEnded(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnError(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnFocus(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnFormdata(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnInput(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnInvalid(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnKeydown(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnKeypress(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnKeyup(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnLoad(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnLoadeddata(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnLoadedmetadata(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnLoadstart(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnMousedown(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnMouseenter(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnMouseleave(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnMousemove(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnMouseout(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnMouseover(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnMouseup(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnPaste(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnPause(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnPlay(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnPlaying(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnProgress(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnRatechange(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnReset(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnResize(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnScroll(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnScrollend(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnSeeked(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnSeeking(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnSelect(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnSlotchange(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnStalled(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnSubmit(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnSuspend(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnTimeupdate(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnToggle(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnVolumechange(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnWaiting(fn engine.EventHandler) *ElementScript {
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
func (e *ElementScript) OnWheel(fn engine.EventHandler) *ElementScript {
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
