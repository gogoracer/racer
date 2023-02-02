/* cSpell:disable */

package goggles

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementTextarea struct {
	*baseElement
}

func TEXTAREA(children ...any) *ElementTextarea {
	return &ElementTextarea{
		baseElement: newBaseElement("textarea", children...),
	}
}

func (e *ElementTextarea) Add(children ...any) *ElementTextarea {
	e.baseElement.add(children...)
	return e
}

func (e *ElementTextarea) Custom(k, v string, dontEscape ...bool) *ElementTextarea {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementTextarea) BindCustom(k string, v string, dontEscape ...bool) *ElementTextarea {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementTextarea) appendAttribute(k string, v string, dontEscape ...bool) *ElementTextarea {
	e.baseElement.appendAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementTextarea) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTextarea) Autofocus(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("autofocus", v, dontEscape...)
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
func (element *ElementTextarea) Enterkeyhint(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementTextarea) Itemid(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("itemid", v, dontEscape...)
	return element
}

// Maxlength is the "maxlength" attribute.
// Maximum length of value
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementTextarea) Maxlength(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("maxlength", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementTextarea) Tabindex(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("tabindex", v, dontEscape...)
	return element
}

// Disabled is the "disabled" attribute.
// Whether the link is disabled
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTextarea) Disabled(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("disabled", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementTextarea) Draggable(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("draggable", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementTextarea) Lang(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("lang", v, dontEscape...)
	return element
}

// Placeholder is the "placeholder" attribute.
// User-visible label to be placed within the form control
// Valid values are constrained to the following:
//   - text
func (element *ElementTextarea) Placeholder(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("placeholder", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementTextarea) Title(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("title", v, dontEscape...)
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
func (element *ElementTextarea) Inputmode(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("inputmode", v, dontEscape...)
	return element
}

// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   - text
func (element *ElementTextarea) Name(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("name", v, dontEscape...)
	return element
}

// Readonly is the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTextarea) Readonly(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("readonly", v, dontEscape...)
	return element
}

// Required is the "required" attribute.
// Whether the control is required for form submission
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTextarea) Required(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("required", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementTextarea) Spellcheck(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("spellcheck", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementTextarea) Class(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("class", v, dontEscape...)
	return element
}

// Form is the "form" attribute.
// Associates the element with a form element
// Valid values are constrained to the following:
//   - id
func (element *ElementTextarea) Form(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("form", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementTextarea) Itemtype(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("itemtype", v, dontEscape...)
	return element
}

// Wrap is the "wrap" attribute.
// How the value of the form control is to be wrapped for form submission
// Valid values are constrained to the following:
//   - soft
//   - soft
//   - hard
//   - hard
func (element *ElementTextarea) Wrap(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("wrap", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementTextarea) Accesskey(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("accesskey", v, dontEscape...)
	return element
}

// Autocomplete is the "autocomplete" attribute.
// Hint for form autofill feature
// Valid values are constrained to the following:
//   - autofill_field
func (element *ElementTextarea) Autocomplete(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("autocomplete", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementTextarea) Dir(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("dir", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementTextarea) Nonce(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("nonce", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementTextarea) Translate(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("translate", v, dontEscape...)
	return element
}

// Rows is the "rows" attribute.
// Number of lines to show
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementTextarea) Rows(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("rows", v, dontEscape...)
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
func (element *ElementTextarea) Autocapitalize(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Dirname is the "dirname" attribute.
// Name of form control to use for sending the element&#39;s directionality in form submission
// Valid values are constrained to the following:
//   - text
func (element *ElementTextarea) Dirname(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("dirname", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementTextarea) Hidden(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("hidden", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementTextarea) Itemref(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("itemref", v, dontEscape...)
	return element
}

// Minlength is the "minlength" attribute.
// Minimum length of value
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementTextarea) Minlength(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("minlength", v, dontEscape...)
	return element
}

// Cols is the "cols" attribute.
// Maximum number of characters per line
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementTextarea) Cols(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("cols", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementTextarea) Id(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("id", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTextarea) Inert(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("inert", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementTextarea) Itemscope(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("itemscope", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementTextarea) Style(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("style", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementTextarea) Contenteditable(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("contenteditable", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementTextarea) Is(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("is", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementTextarea) Itemprop(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("itemprop", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementTextarea) Popover(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("popover", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementTextarea) Slot(v string, dontEscape ...bool) *ElementTextarea {
	element.appendAttribute("slot", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementTextarea) OnAuxclick(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnBeforematch(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnBeforetoggle(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnBlur(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnCancel(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnCanplay(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnCanplaythrough(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnChange(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnClick(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnClose(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnContextlost(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnContextmenu(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnContextrestored(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnCopy(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnCuechange(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnCut(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDblclick(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDrag(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDragend(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDragenter(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDragleave(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDragover(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDragstart(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDrop(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnDurationchange(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnEmptied(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnEnded(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnError(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnFocus(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnFormdata(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnInput(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnInvalid(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnKeydown(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnKeypress(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnKeyup(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnLoad(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnLoadeddata(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnLoadedmetadata(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnLoadstart(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnMousedown(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnMouseenter(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnMouseleave(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnMousemove(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnMouseout(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnMouseover(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnMouseup(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnPaste(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnPause(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnPlay(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnPlaying(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnProgress(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnRatechange(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnReset(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnResize(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnScroll(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnScrollend(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnSeeked(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnSeeking(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnSelect(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnSlotchange(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnStalled(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnSubmit(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnSuspend(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnTimeupdate(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnToggle(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnVolumechange(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnWaiting(fn engine.EventHandler) *ElementTextarea {
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
func (e *ElementTextarea) OnWheel(fn engine.EventHandler) *ElementTextarea {
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
