/* cSpell:disable */

package handlebars

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementButton struct {
	*baseElement
}

func BUTTON(children ...any) *ElementButton {
	return &ElementButton{
		baseElement: newBaseElement("button", children...),
	}
}

func (e *ElementButton) Add(children ...any) *ElementButton {
	e.baseElement.add(children...)
	return e
}

func (e *ElementButton) Custom(k, v string, dontEscape ...bool) *ElementButton {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementButton) BindCustom(k string, v string, dontEscape ...bool) *ElementButton {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementButton) setAttribute(k string, v string, dontEscape ...bool) *ElementButton {
	e.baseElement.setAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementButton) GenerateVDOM() interface{} {
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
func (element *ElementButton) Inputmode(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("inputmode", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementButton) Is(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("is", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementButton) Itemprop(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("itemprop", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementButton) Lang(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("lang", v, dontEscape...)
	return element
}

// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   - text
func (element *ElementButton) Name(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("name", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementButton) Slot(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("slot", v, dontEscape...)
	return element
}

// Formenctype is the "formenctype" attribute.
// Entry list encoding type to use for form submission
// Valid values are constrained to the following:
//   - application/x_www_form_urlencoded
//   - application/x_www_form_urlencoded
//   - multipart/form_data
//   - multipart/form_data
//   - text/plain
//   - text/plain
func (element *ElementButton) Formenctype(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("formenctype", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementButton) Inert(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("inert", v, dontEscape...)
	return element
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementButton) Style(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("style", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementButton) Accesskey(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("accesskey", v, dontEscape...)
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
func (element *ElementButton) Autocapitalize(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementButton) Spellcheck(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("spellcheck", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementButton) Dir(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("dir", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementButton) Nonce(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("nonce", v, dontEscape...)
	return element
}

// Formtarget is the "formtarget" attribute.
// Navigable for form submission
// Valid values are constrained to the following:
//   - valid_navigable_target_name_or_keyword
func (element *ElementButton) Formtarget(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("formtarget", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementButton) Itemref(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("itemref", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementButton) Itemscope(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("itemscope", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementButton) Popover(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("popover", v, dontEscape...)
	return element
}

// Popovertoggletarget is the "popovertoggletarget" attribute.
// Toggles the specified popover element when clicked
// Valid values are constrained to the following:
func (element *ElementButton) Popovertoggletarget(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("popovertoggletarget", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementButton) Autofocus(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("autofocus", v, dontEscape...)
	return element
}

// Formmethod is the "formmethod" attribute.
// Variant to use for form submission
// Valid values are constrained to the following:
//   - get
//   - post
//   - dialog
func (element *ElementButton) Formmethod(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("formmethod", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementButton) Itemtype(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("itemtype", v, dontEscape...)
	return element
}

// Popoverhidetarget is the "popoverhidetarget" attribute.
// Hides the specified popover element when clicked
// Valid values are constrained to the following:
func (element *ElementButton) Popoverhidetarget(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("popoverhidetarget", v, dontEscape...)
	return element
}

// Type is the "type" attribute.
// Type of script
// Valid values are constrained to the following:
//   - module
//   - valid_mime_type_string
//   - java_script_mime_type_essence_match
func (element *ElementButton) Type(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("type", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementButton) Contenteditable(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("contenteditable", v, dontEscape...)
	return element
}

// Disabled is the "disabled" attribute.
// Whether the link is disabled
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementButton) Disabled(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("disabled", v, dontEscape...)
	return element
}

// Formnovalidate is the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementButton) Formnovalidate(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("formnovalidate", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementButton) Hidden(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("hidden", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementButton) Id(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("id", v, dontEscape...)
	return element
}

// Popovershowtarget is the "popovershowtarget" attribute.
// Shows the specified popover element when clicked
// Valid values are constrained to the following:
func (element *ElementButton) Popovershowtarget(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("popovershowtarget", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementButton) Tabindex(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("tabindex", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementButton) Title(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("title", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementButton) Draggable(v string, dontEscape ...bool) *ElementButton {
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
func (element *ElementButton) Enterkeyhint(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Formaction is the "formaction" attribute.
// URL to use for form submission
// Valid values are constrained to the following:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (element *ElementButton) Formaction(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("formaction", v, dontEscape...)
	return element
}

// Value is the "value" attribute.
// Current value of the element
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementButton) Value(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("value", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementButton) Class(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("class", v, dontEscape...)
	return element
}

// Form is the "form" attribute.
// Associates the element with a form element
// Valid values are constrained to the following:
//   - id
func (element *ElementButton) Form(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("form", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementButton) Itemid(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("itemid", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementButton) Translate(v string, dontEscape ...bool) *ElementButton {
	element.setAttribute("translate", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementButton) OnAuxclick(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnBeforematch(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnBeforetoggle(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnBlur(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnCancel(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnCanplay(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnCanplaythrough(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnChange(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnClick(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnClose(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnContextlost(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnContextmenu(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnContextrestored(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnCopy(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnCuechange(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnCut(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDblclick(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDrag(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDragend(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDragenter(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDragleave(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDragover(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDragstart(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDrop(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnDurationchange(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnEmptied(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnEnded(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnError(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnFocus(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnFormdata(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnInput(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnInvalid(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnKeydown(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnKeypress(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnKeyup(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnLoad(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnLoadeddata(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnLoadedmetadata(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnLoadstart(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnMousedown(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnMouseenter(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnMouseleave(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnMousemove(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnMouseout(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnMouseover(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnMouseup(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnPaste(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnPause(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnPlay(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnPlaying(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnProgress(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnRatechange(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnReset(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnResize(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnScroll(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnScrollend(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnSeeked(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnSeeking(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnSelect(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnSlotchange(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnStalled(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnSubmit(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnSuspend(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnTimeupdate(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnToggle(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnVolumechange(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnWaiting(fn engine.EventHandler) *ElementButton {
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
func (e *ElementButton) OnWheel(fn engine.EventHandler) *ElementButton {
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
