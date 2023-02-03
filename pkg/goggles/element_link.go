/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementLink struct {
	*baseElement
}

func LINK(children ...any) *ElementLink {
	return &ElementLink{
		baseElement: newBaseElement("link", children...),
	}
}

func (e *ElementLink) Add(children ...any) *ElementLink {
	e.baseElement.add(children...)
	return e
}

func (e *ElementLink) Custom(k, v string, dontEscape ...bool) *ElementLink {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementLink) BindCustom(k string, v string, dontEscape ...bool) *ElementLink {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementLink) appendAttribute(k string, v string, dontEscape ...bool) *ElementLink {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementLink) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementLink) Contenteditable(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Imagesrcset is the "imagesrcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc. (for rel=&quot;preload&quot;)
// Valid values are constrained to the following:
//   - image_candidate_strings
func (element *ElementLink) Imagesrcset(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("imagesrcset", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementLink) Is(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementLink) Itemtype(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Rel is the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementLink) Rel(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("rel", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementLink) Accesskey(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Color is the "color" attribute.
// Color to use when customizing a site&#39;s icon (for rel=&quot;mask-icon&quot;)
// Valid values are constrained to the following:
//   - &lt;color&gt;
func (element *ElementLink) Color(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("color", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementLink) Hidden(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementLink) Itemid(v string, dontEscape ...bool) *ElementLink {
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
func (element *ElementLink) Popover(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementLink) Slot(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementLink) Tabindex(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// As is the "as" attribute.
// Potential destination for a preload request (for rel=&quot;preload&quot; and rel=&quot;modulepreload&quot;)
// Valid values are constrained to the following:
//   - potential_destination
//   - rel
//   - rel
//   - preload
//   - preload
//   - script_like_destination
//   - rel
//   - rel
//   - modulepreload
//   - modulepreload
func (element *ElementLink) As(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("as", v, dontEscape...)
	return element
}

// Imagesizes is the "imagesizes" attribute.
// Image sizes for different page layouts (for rel=&quot;preload&quot;)
// Valid values are constrained to the following:
//   - valid_source_size_list
func (element *ElementLink) Imagesizes(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("imagesizes", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementLink) Lang(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementLink) Nonce(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementLink) Title(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("title", v, dontEscape...)
	return element
}

// Blocking is the "blocking" attribute.
// Whether the element is potentially render-blocking
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementLink) Blocking(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("blocking", v, dontEscape...)
	return element
}

// Disabled is the "disabled" attribute.
// Whether the link is disabled
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementLink) Disabled(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("disabled", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementLink) Draggable(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Hreflang is the "hreflang" attribute.
// Language of the linked resource
// Valid values are constrained to the following:
func (element *ElementLink) Hreflang(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("hreflang", v, dontEscape...)
	return element
}

// Media is the "media" attribute.
// Applicable media
// Valid values are constrained to the following:
//   - valid_media_query_list
func (element *ElementLink) Media(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("media", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementLink) Translate(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Href is the "href" attribute.
// Document base URL
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementLink) Href(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("href", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementLink) Id(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementLink) Itemprop(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Referrerpolicy is the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Valid values are constrained to the following:
//   - referrer_policy
func (element *ElementLink) Referrerpolicy(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("referrerpolicy", v, dontEscape...)
	return element
}

// Sizes is the "sizes" attribute.
// Image sizes for different page layouts
// Valid values are constrained to the following:
//   - valid_source_size_list
func (element *ElementLink) Sizes(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("sizes", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementLink) Class(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementLink) Dir(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Integrity is the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Valid values are constrained to the following:
//   - text
func (element *ElementLink) Integrity(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("integrity", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementLink) Spellcheck(v string, dontEscape ...bool) *ElementLink {
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
func (element *ElementLink) Enterkeyhint(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementLink) Itemref(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementLink) Itemscope(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("itemscope", v, dontEscape...)
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
func (element *ElementLink) Autocapitalize(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementLink) Autofocus(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("autofocus", v, dontEscape...)
	return element
}

// Crossorigin is the "crossorigin" attribute.
// How the element handles crossorigin requests
// Valid values are constrained to the following:
//   - anonymous
//   - anonymous
//   - use_credentials
//   - use_credentials
func (element *ElementLink) Crossorigin(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("crossorigin", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementLink) Inert(v string, dontEscape ...bool) *ElementLink {
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
func (element *ElementLink) Inputmode(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementLink) Style(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Type is the "type" attribute.
// Type of script
// Valid values are constrained to the following:
//   - module
//   - valid_mime_type_string
//   - java_script_mime_type_essence_match
func (element *ElementLink) Type(v string, dontEscape ...bool) *ElementLink {
	element.appendAttribute("type", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementLink) OnAuxclick(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnBeforematch(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnBeforetoggle(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnBlur(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnCancel(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnCanplay(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnCanplaythrough(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnChange(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnClick(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnClose(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnContextlost(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnContextmenu(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnContextrestored(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnCopy(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnCuechange(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnCut(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDblclick(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDrag(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDragend(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDragenter(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDragleave(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDragover(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDragstart(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDrop(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnDurationchange(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnEmptied(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnEnded(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnError(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnFocus(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnFormdata(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnInput(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnInvalid(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnKeydown(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnKeypress(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnKeyup(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnLoad(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnLoadeddata(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnLoadedmetadata(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnLoadstart(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnMousedown(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnMouseenter(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnMouseleave(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnMousemove(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnMouseout(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnMouseover(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnMouseup(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnPaste(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnPause(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnPlay(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnPlaying(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnProgress(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnRatechange(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnReset(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnResize(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnScroll(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnScrollend(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnSeeked(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnSeeking(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnSelect(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnSlotchange(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnStalled(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnSubmit(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnSuspend(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnTimeupdate(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnToggle(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnVolumechange(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnWaiting(fn engine.EventHandler) *ElementLink {
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
func (e *ElementLink) OnWheel(fn engine.EventHandler) *ElementLink {
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
