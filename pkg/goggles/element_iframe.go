/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementIframe struct {
	*baseElement
}

func IFRAME(children ...any) *ElementIframe {
	return &ElementIframe{
		baseElement: newBaseElement("iframe", children...),
	}
}

func (e *ElementIframe) Add(children ...any) *ElementIframe {
	e.baseElement.add(children...)
	return e
}

func (e *ElementIframe) Custom(k, v string, dontEscape ...bool) *ElementIframe {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementIframe) BindCustom(k string, v string, dontEscape ...bool) *ElementIframe {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementIframe) appendAttribute(k string, v string, dontEscape ...bool) *ElementIframe {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementIframe) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementIframe) Hidden(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementIframe) Is(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Src is the "src" attribute.
// Address of the resource
// Valid values are constrained to the following:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (element *ElementIframe) Src(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("src", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementIframe) Translate(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Allow is the "allow" attribute.
// Permissions policy to be applied to the iframe&#39;s contents
// Valid values are constrained to the following:
//   - serialized_permissions_policy
func (element *ElementIframe) Allow(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("allow", v, dontEscape...)
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
func (element *ElementIframe) Autocapitalize(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Height is the "height" attribute.
// Vertical dimension
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementIframe) Height(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("height", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementIframe) Itemscope(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementIframe) Style(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementIframe) Dir(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("dir", v, dontEscape...)
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
func (element *ElementIframe) Enterkeyhint(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementIframe) Itemref(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Srcdoc is the "srcdoc" attribute.
// A document to render in the iframe
// Valid values are constrained to the following:
//   - an_iframe_srcdoc_document
//   - iframe
//   - srcdoc
func (element *ElementIframe) Srcdoc(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("srcdoc", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementIframe) Title(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementIframe) Draggable(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   - text
func (element *ElementIframe) Name(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("name", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementIframe) Spellcheck(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("spellcheck", v, dontEscape...)
	return element
}

// Referrerpolicy is the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Valid values are constrained to the following:
//   - referrer_policy
func (element *ElementIframe) Referrerpolicy(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("referrerpolicy", v, dontEscape...)
	return element
}

// Width is the "width" attribute.
// Horizontal dimension
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementIframe) Width(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("width", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementIframe) Autofocus(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementIframe) Contenteditable(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementIframe) Itemid(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementIframe) Class(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Loading is the "loading" attribute.
// Used when determining loading deferral
// Valid values are constrained to the following:
//   - lazy
//   - lazy
//   - eager
//   - eager
func (element *ElementIframe) Loading(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("loading", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementIframe) Slot(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementIframe) Tabindex(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementIframe) Accesskey(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Allowfullscreen is the "allowfullscreen" attribute.
// Whether to allow the iframe&#39;s contents to use requestFullscreen()
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementIframe) Allowfullscreen(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("allowfullscreen", v, dontEscape...)
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
func (element *ElementIframe) Inputmode(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementIframe) Itemtype(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementIframe) Lang(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementIframe) Nonce(v string, dontEscape ...bool) *ElementIframe {
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
func (element *ElementIframe) Popover(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Sandbox is the "sandbox" attribute.
// Security rules for nested content
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - ascii_case_insensitive
//   - allow_downloads
//   - allow_downloads
//   - allow_forms
//   - allow_forms
//   - allow_modals
//   - allow_modals
//   - allow_orientation_lock
//   - allow_orientation_lock
//   - allow_pointer_lock
//   - allow_pointer_lock
//   - allow_popups
//   - allow_popups
//   - allow_popups_to_escape_sandbox
//   - allow_popups_to_escape_sandbox
//   - allow_presentation
//   - allow_presentation
//   - allow_same_origin
//   - allow_same_origin
//   - allow_scripts
//   - allow_scripts
//   - allow_top_navigation
//   - allow_top_navigation
//   - allow_top_navigation_by_user_activation
//   - allow_top_navigation_by_user_activation
//   - allow_top_navigation_to_custom_protocols
//   - allow_top_navigation_to_custom_protocols
func (element *ElementIframe) Sandbox(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("sandbox", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementIframe) Id(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementIframe) Inert(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementIframe) Itemprop(v string, dontEscape ...bool) *ElementIframe {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementIframe) OnAuxclick(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnBeforematch(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnBeforetoggle(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnBlur(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnCancel(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnCanplay(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnCanplaythrough(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnChange(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnClick(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnClose(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnContextlost(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnContextmenu(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnContextrestored(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnCopy(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnCuechange(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnCut(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDblclick(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDrag(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDragend(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDragenter(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDragleave(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDragover(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDragstart(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDrop(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnDurationchange(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnEmptied(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnEnded(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnError(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnFocus(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnFormdata(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnInput(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnInvalid(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnKeydown(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnKeypress(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnKeyup(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnLoad(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnLoadeddata(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnLoadedmetadata(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnLoadstart(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnMousedown(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnMouseenter(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnMouseleave(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnMousemove(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnMouseout(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnMouseover(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnMouseup(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnPaste(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnPause(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnPlay(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnPlaying(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnProgress(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnRatechange(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnReset(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnResize(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnScroll(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnScrollend(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnSeeked(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnSeeking(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnSelect(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnSlotchange(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnStalled(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnSubmit(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnSuspend(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnTimeupdate(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnToggle(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnVolumechange(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnWaiting(fn engine.EventHandler) *ElementIframe {
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
func (e *ElementIframe) OnWheel(fn engine.EventHandler) *ElementIframe {
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
