/* cSpell:disable */

package handlebars

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementImg struct {
	*baseElement
}

func IMG(children ...any) *ElementImg {
	return &ElementImg{
		baseElement: newBaseElement("img", children...),
	}
}

func (e *ElementImg) Add(children ...any) *ElementImg {
	e.baseElement.add(children...)
	return e
}

func (e *ElementImg) Custom(k, v string, dontEscape ...bool) *ElementImg {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementImg) BindCustom(k string, v string, dontEscape ...bool) *ElementImg {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementImg) setAttribute(k string, v string, dontEscape ...bool) *ElementImg {
	e.baseElement.setAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementImg) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Ismap is the "ismap" attribute.
// Whether the image is a server-side image map
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementImg) Ismap(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("ismap", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementImg) Style(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("style", v, dontEscape...)
	return element
}

// Decoding is the "decoding" attribute.
// Decoding hint to use when processing this image for presentation
// Valid values are constrained to the following:
//   - sync
//   - sync
//   - async
//   - async
//   - auto
//   - auto
func (element *ElementImg) Decoding(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("decoding", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementImg) Inert(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("inert", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementImg) Tabindex(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("tabindex", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementImg) Hidden(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("hidden", v, dontEscape...)
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
func (element *ElementImg) Inputmode(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("inputmode", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementImg) Class(v string, dontEscape ...bool) *ElementImg {
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
func (element *ElementImg) Dir(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("dir", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementImg) Itemref(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("itemref", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementImg) Itemtype(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("itemtype", v, dontEscape...)
	return element
}

// Alt is the "alt" attribute.
// Replacement text for use when images are not available
// Valid values are constrained to the following:
//   - text
func (element *ElementImg) Alt(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("alt", v, dontEscape...)
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
func (element *ElementImg) Autocapitalize(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementImg) Itemscope(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("itemscope", v, dontEscape...)
	return element
}

// Referrerpolicy is the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Valid values are constrained to the following:
//   - referrer_policy
func (element *ElementImg) Referrerpolicy(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("referrerpolicy", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementImg) Spellcheck(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("spellcheck", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementImg) Translate(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("translate", v, dontEscape...)
	return element
}

// Width is the "width" attribute.
// Horizontal dimension
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementImg) Width(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("width", v, dontEscape...)
	return element
}

// Height is the "height" attribute.
// Vertical dimension
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementImg) Height(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("height", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementImg) Itemprop(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("itemprop", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementImg) Slot(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("slot", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementImg) Id(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("id", v, dontEscape...)
	return element
}

// Loading is the "loading" attribute.
// Used when determining loading deferral
// Valid values are constrained to the following:
//   - lazy
//   - lazy
//   - eager
//   - eager
func (element *ElementImg) Loading(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("loading", v, dontEscape...)
	return element
}

// Srcset is the "srcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc.
// Valid values are constrained to the following:
//   - image_candidate_strings
func (element *ElementImg) Srcset(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("srcset", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementImg) Draggable(v string, dontEscape ...bool) *ElementImg {
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
func (element *ElementImg) Enterkeyhint(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementImg) Nonce(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("nonce", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementImg) Popover(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("popover", v, dontEscape...)
	return element
}

// Sizes is the "sizes" attribute.
// Image sizes for different page layouts
// Valid values are constrained to the following:
//   - valid_source_size_list
func (element *ElementImg) Sizes(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("sizes", v, dontEscape...)
	return element
}

// Src is the "src" attribute.
// Address of the resource
// Valid values are constrained to the following:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (element *ElementImg) Src(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("src", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementImg) Title(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("title", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementImg) Contenteditable(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("contenteditable", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementImg) Itemid(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("itemid", v, dontEscape...)
	return element
}

// Crossorigin is the "crossorigin" attribute.
// How the element handles crossorigin requests
// Valid values are constrained to the following:
//   - anonymous
//   - anonymous
//   - use_credentials
//   - use_credentials
func (element *ElementImg) Crossorigin(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("crossorigin", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementImg) Is(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("is", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementImg) Lang(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("lang", v, dontEscape...)
	return element
}

// Usemap is the "usemap" attribute.
// Name of image map to use
// Valid values are constrained to the following:
//   - valid_hash_name_reference
func (element *ElementImg) Usemap(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("usemap", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementImg) Accesskey(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("accesskey", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementImg) Autofocus(v string, dontEscape ...bool) *ElementImg {
	element.setAttribute("autofocus", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementImg) OnAuxclick(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnBeforematch(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnBeforetoggle(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnBlur(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnCancel(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnCanplay(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnCanplaythrough(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnChange(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnClick(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnClose(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnContextlost(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnContextmenu(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnContextrestored(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnCopy(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnCuechange(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnCut(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDblclick(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDrag(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDragend(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDragenter(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDragleave(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDragover(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDragstart(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDrop(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnDurationchange(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnEmptied(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnEnded(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnError(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnFocus(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnFormdata(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnInput(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnInvalid(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnKeydown(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnKeypress(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnKeyup(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnLoad(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnLoadeddata(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnLoadedmetadata(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnLoadstart(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnMousedown(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnMouseenter(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnMouseleave(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnMousemove(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnMouseout(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnMouseover(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnMouseup(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnPaste(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnPause(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnPlay(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnPlaying(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnProgress(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnRatechange(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnReset(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnResize(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnScroll(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnScrollend(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnSeeked(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnSeeking(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnSelect(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnSlotchange(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnStalled(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnSubmit(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnSuspend(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnTimeupdate(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnToggle(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnVolumechange(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnWaiting(fn engine.EventHandler) *ElementImg {
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
func (e *ElementImg) OnWheel(fn engine.EventHandler) *ElementImg {
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
