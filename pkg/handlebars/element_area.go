/* cSpell:disable */

package handlebars

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementArea struct {
	*baseElement
}

func AREA(children ...any) *ElementArea {
	return &ElementArea{
		baseElement: newBaseElement("area", children...),
	}
}

func (e *ElementArea) Add(children ...any) *ElementArea {
	e.baseElement.add(children...)
	return e
}

func (e *ElementArea) Custom(k, v string, dontEscape ...bool) *ElementArea {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementArea) BindCustom(k string, v string, dontEscape ...bool) *ElementArea {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementArea) setAttribute(k string, v string, dontEscape ...bool) *ElementArea {
	e.baseElement.setAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementArea) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementArea) Dir(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("dir", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementArea) Class(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("class", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementArea) Itemprop(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("itemprop", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementArea) Tabindex(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("tabindex", v, dontEscape...)
	return element
}

// Target is the "target" attribute.
// Navigable for form submission
// Valid values are constrained to the following:
//   - valid_navigable_target_name_or_keyword
func (element *ElementArea) Target(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("target", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementArea) Draggable(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("draggable", v, dontEscape...)
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
func (element *ElementArea) Enterkeyhint(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementArea) Inert(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("inert", v, dontEscape...)
	return element
}

// Referrerpolicy is the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Valid values are constrained to the following:
//   - referrer_policy
func (element *ElementArea) Referrerpolicy(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("referrerpolicy", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementArea) Slot(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("slot", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementArea) Contenteditable(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("contenteditable", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementArea) Itemid(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("itemid", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementArea) Itemref(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("itemref", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementArea) Nonce(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("nonce", v, dontEscape...)
	return element
}

// Alt is the "alt" attribute.
// Replacement text for use when images are not available
// Valid values are constrained to the following:
//   - text
func (element *ElementArea) Alt(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("alt", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementArea) Is(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("is", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementArea) Itemtype(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("itemtype", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementArea) Popover(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("popover", v, dontEscape...)
	return element
}

// Shape is the "shape" attribute.
// The kind of shape to be created in an image map
// Valid values are constrained to the following:
//   - circle
//   - circle
//   - default
//   - default
//   - poly
//   - poly
//   - rect
//   - rect
func (element *ElementArea) Shape(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("shape", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementArea) Spellcheck(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("spellcheck", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementArea) Translate(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("translate", v, dontEscape...)
	return element
}

// Href is the "href" attribute.
// Document base URL
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementArea) Href(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("href", v, dontEscape...)
	return element
}

// Ping is the "ping" attribute.
// URLs to ping
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
//   - valid_non_empty_ur_ls
func (element *ElementArea) Ping(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("ping", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementArea) Title(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("title", v, dontEscape...)
	return element
}

// Coords is the "coords" attribute.
// Coordinates for the shape to be created in an image map
// Valid values are constrained to the following:
//   - valid_list_of_floating_point_numbers
func (element *ElementArea) Coords(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("coords", v, dontEscape...)
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
func (element *ElementArea) Autocapitalize(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementArea) Autofocus(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("autofocus", v, dontEscape...)
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
func (element *ElementArea) Inputmode(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementArea) Itemscope(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("itemscope", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementArea) Lang(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("lang", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementArea) Accesskey(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("accesskey", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementArea) Hidden(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("hidden", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementArea) Id(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("id", v, dontEscape...)
	return element
}

// Rel is the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementArea) Rel(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("rel", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementArea) Style(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("style", v, dontEscape...)
	return element
}

// Download is the "download" attribute.
// Whether to download the resource instead of navigating to it, and its filename if so
// Valid values are constrained to the following:
func (element *ElementArea) Download(v string, dontEscape ...bool) *ElementArea {
	element.setAttribute("download", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementArea) OnAuxclick(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnBeforematch(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnBeforetoggle(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnBlur(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnCancel(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnCanplay(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnCanplaythrough(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnChange(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnClick(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnClose(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnContextlost(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnContextmenu(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnContextrestored(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnCopy(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnCuechange(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnCut(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDblclick(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDrag(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDragend(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDragenter(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDragleave(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDragover(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDragstart(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDrop(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnDurationchange(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnEmptied(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnEnded(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnError(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnFocus(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnFormdata(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnInput(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnInvalid(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnKeydown(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnKeypress(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnKeyup(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnLoad(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnLoadeddata(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnLoadedmetadata(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnLoadstart(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnMousedown(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnMouseenter(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnMouseleave(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnMousemove(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnMouseout(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnMouseover(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnMouseup(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnPaste(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnPause(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnPlay(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnPlaying(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnProgress(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnRatechange(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnReset(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnResize(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnScroll(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnScrollend(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnSeeked(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnSeeking(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnSelect(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnSlotchange(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnStalled(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnSubmit(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnSuspend(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnTimeupdate(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnToggle(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnVolumechange(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnWaiting(fn engine.EventHandler) *ElementArea {
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
func (e *ElementArea) OnWheel(fn engine.EventHandler) *ElementArea {
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
