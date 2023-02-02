/* cSpell:disable */

package handlebars

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementForm struct {
	*baseElement
}

func FORM(children ...any) *ElementForm {
	return &ElementForm{
		baseElement: newBaseElement("form", children...),
	}
}

func (e *ElementForm) Add(children ...any) *ElementForm {
	e.baseElement.add(children...)
	return e
}

func (e *ElementForm) Custom(k, v string, dontEscape ...bool) *ElementForm {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementForm) BindCustom(k string, v string, dontEscape ...bool) *ElementForm {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementForm) setAttribute(k string, v string, dontEscape ...bool) *ElementForm {
	e.baseElement.setAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementForm) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
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
func (element *ElementForm) Inputmode(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("inputmode", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementForm) Itemprop(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("itemprop", v, dontEscape...)
	return element
}

// Method is the "method" attribute.
// Variant to use for form submission
// Valid values are constrained to the following:
//   - get
//   - get
//   - post
//   - post
//   - dialog
//   - dialog
func (element *ElementForm) Method(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("method", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementForm) Class(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("class", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementForm) Itemscope(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("itemscope", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementForm) Lang(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("lang", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementForm) Spellcheck(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("spellcheck", v, dontEscape...)
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
func (element *ElementForm) Autocapitalize(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementForm) Draggable(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("draggable", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementForm) Slot(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("slot", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementForm) Translate(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("translate", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementForm) Autofocus(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("autofocus", v, dontEscape...)
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
func (element *ElementForm) Enterkeyhint(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementForm) Is(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("is", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementForm) Dir(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("dir", v, dontEscape...)
	return element
}

// Enctype is the "enctype" attribute.
// Entry list encoding type to use for form submission
// Valid values are constrained to the following:
//   - application/x_www_form_urlencoded
//   - application/x_www_form_urlencoded
//   - multipart/form_data
//   - multipart/form_data
//   - text/plain
//   - text/plain
func (element *ElementForm) Enctype(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("enctype", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementForm) Id(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("id", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementForm) Inert(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("inert", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementForm) Itemid(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("itemid", v, dontEscape...)
	return element
}

// Novalidate is the "novalidate" attribute.
// Bypass form control validation for form submission
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementForm) Novalidate(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("novalidate", v, dontEscape...)
	return element
}

// Autocomplete is the "autocomplete" attribute.
// Hint for form autofill feature
// Valid values are constrained to the following:
//   - autofill_field
func (element *ElementForm) Autocomplete(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("autocomplete", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementForm) Popover(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("popover", v, dontEscape...)
	return element
}

// Action is the "action" attribute.
// URL to use for form submission
// Valid values are constrained to the following:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (element *ElementForm) Action(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("action", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementForm) Contenteditable(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("contenteditable", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementForm) Itemref(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("itemref", v, dontEscape...)
	return element
}

// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   - text
func (element *ElementForm) Name(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("name", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementForm) Nonce(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("nonce", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementForm) Style(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("style", v, dontEscape...)
	return element
}

// AcceptCharset is the "accept-charset" attribute.
// Character encodings to use for form submission
// Valid values are constrained to the following:
//   - ascii_case_insensitive
//   - utf_8
func (element *ElementForm) AcceptCharset(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("accept-charset", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementForm) Accesskey(v string, dontEscape ...bool) *ElementForm {
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
func (element *ElementForm) Hidden(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("hidden", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementForm) Itemtype(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("itemtype", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementForm) Tabindex(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("tabindex", v, dontEscape...)
	return element
}

// Target is the "target" attribute.
// Navigable for form submission
// Valid values are constrained to the following:
//   - valid_navigable_target_name_or_keyword
func (element *ElementForm) Target(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("target", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementForm) Title(v string, dontEscape ...bool) *ElementForm {
	element.setAttribute("title", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementForm) OnAuxclick(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnBeforematch(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnBeforetoggle(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnBlur(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnCancel(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnCanplay(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnCanplaythrough(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnChange(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnClick(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnClose(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnContextlost(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnContextmenu(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnContextrestored(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnCopy(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnCuechange(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnCut(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDblclick(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDrag(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDragend(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDragenter(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDragleave(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDragover(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDragstart(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDrop(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnDurationchange(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnEmptied(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnEnded(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnError(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnFocus(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnFormdata(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnInput(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnInvalid(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnKeydown(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnKeypress(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnKeyup(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnLoad(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnLoadeddata(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnLoadedmetadata(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnLoadstart(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnMousedown(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnMouseenter(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnMouseleave(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnMousemove(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnMouseout(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnMouseover(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnMouseup(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnPaste(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnPause(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnPlay(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnPlaying(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnProgress(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnRatechange(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnReset(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnResize(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnScroll(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnScrollend(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnSeeked(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnSeeking(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnSelect(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnSlotchange(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnStalled(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnSubmit(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnSuspend(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnTimeupdate(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnToggle(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnVolumechange(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnWaiting(fn engine.EventHandler) *ElementForm {
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
func (e *ElementForm) OnWheel(fn engine.EventHandler) *ElementForm {
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
