/* cSpell:disable */

package handlebars

import (
	"github.com/gogoracer/racer/pkg/engine"
)

type ElementInput struct {
	*baseElement
}

func INPUT(children ...any) *ElementInput {
	return &ElementInput{
		baseElement: newBaseElement("input", children...),
	}
}

func (e *ElementInput) Add(children ...any) *ElementInput {
	e.baseElement.add(children...)
	return e
}

func (e *ElementInput) Custom(k, v string, dontEscape ...bool) *ElementInput {
	e.baseElement.custom(k, v, dontEscape...)
	return e
}

func (e *ElementInput) BindCustom(k string, v string, dontEscape ...bool) *ElementInput {
	e.baseElement.bindCustom(k, v, dontEscape...)
	return e
}

func (e *ElementInput) setAttribute(k string, v string, dontEscape ...bool) *ElementInput {
	e.baseElement.setAttribute(k, v, dontEscape...)
	return e
}

func (e *ElementInput) GenerateVDOM() interface{} {
	return e.baseElement.generateVDOM()
}

// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (element *ElementInput) Style(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("style", v, dontEscape...)
	return element
}

// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (element *ElementInput) Lang(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("lang", v, dontEscape...)
	return element
}

// Pattern is the "pattern" attribute.
// Pattern to be matched by the form control&#39;s value
// Valid values are constrained to the following:
//   - pattern
func (element *ElementInput) Pattern(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("pattern", v, dontEscape...)
	return element
}

// Readonly is the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Readonly(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("readonly", v, dontEscape...)
	return element
}

// Step is the "step" attribute.
// Granularity to be matched by the form control&#39;s value
// Valid values are constrained to the following:
//   - valid_floating_point_number
//   - any
func (element *ElementInput) Step(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("step", v, dontEscape...)
	return element
}

// Disabled is the "disabled" attribute.
// Whether the link is disabled
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Disabled(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("disabled", v, dontEscape...)
	return element
}

// Alt is the "alt" attribute.
// Replacement text for use when images are not available
// Valid values are constrained to the following:
//   - text
func (element *ElementInput) Alt(v string, dontEscape ...bool) *ElementInput {
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
func (element *ElementInput) Autocapitalize(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("autocapitalize", v, dontEscape...)
	return element
}

// Checked is the "checked" attribute.
// Whether the control is checked
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Checked(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("checked", v, dontEscape...)
	return element
}

// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (element *ElementInput) Dir(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("dir", v, dontEscape...)
	return element
}

// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set_of_space_separated_tokens
func (element *ElementInput) Class(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("class", v, dontEscape...)
	return element
}

// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   - text
func (element *ElementInput) Nonce(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("nonce", v, dontEscape...)
	return element
}

// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - text
func (element *ElementInput) Slot(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("slot", v, dontEscape...)
	return element
}

// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   - text
func (element *ElementInput) Title(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("title", v, dontEscape...)
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
func (element *ElementInput) Enterkeyhint(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("enterkeyhint", v, dontEscape...)
	return element
}

// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementInput) Contenteditable(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("contenteditable", v, dontEscape...)
	return element
}

// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   - text
func (element *ElementInput) Id(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("id", v, dontEscape...)
	return element
}

// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Inert(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("inert", v, dontEscape...)
	return element
}

// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Itemscope(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("itemscope", v, dontEscape...)
	return element
}

// Max is the "max" attribute.
// Upper bound of range
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementInput) Max(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("max", v, dontEscape...)
	return element
}

// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (element *ElementInput) Accesskey(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("accesskey", v, dontEscape...)
	return element
}

// Dirname is the "dirname" attribute.
// Name of form control to use for sending the element&#39;s directionality in form submission
// Valid values are constrained to the following:
//   - text
func (element *ElementInput) Dirname(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("dirname", v, dontEscape...)
	return element
}

// Form is the "form" attribute.
// Associates the element with a form element
// Valid values are constrained to the following:
//   - id
func (element *ElementInput) Form(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("form", v, dontEscape...)
	return element
}

// Formnovalidate is the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Formnovalidate(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("formnovalidate", v, dontEscape...)
	return element
}

// Src is the "src" attribute.
// Address of the resource
// Valid values are constrained to the following:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (element *ElementInput) Src(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("src", v, dontEscape...)
	return element
}

// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (element *ElementInput) Translate(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("translate", v, dontEscape...)
	return element
}

// Multiple is the "multiple" attribute.
// Whether to allow multiple values
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Multiple(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("multiple", v, dontEscape...)
	return element
}

// Popoverhidetarget is the "popoverhidetarget" attribute.
// Hides the specified popover element when clicked
// Valid values are constrained to the following:
func (element *ElementInput) Popoverhidetarget(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("popoverhidetarget", v, dontEscape...)
	return element
}

// Popovertoggletarget is the "popovertoggletarget" attribute.
// Toggles the specified popover element when clicked
// Valid values are constrained to the following:
func (element *ElementInput) Popovertoggletarget(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("popovertoggletarget", v, dontEscape...)
	return element
}

// Size is the "size" attribute.
// Size of the control
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementInput) Size(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("size", v, dontEscape...)
	return element
}

// Required is the "required" attribute.
// Whether the control is required for form submission
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Required(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("required", v, dontEscape...)
	return element
}

// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementInput) Spellcheck(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("spellcheck", v, dontEscape...)
	return element
}

// Value is the "value" attribute.
// Current value of the element
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementInput) Value(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("value", v, dontEscape...)
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
func (element *ElementInput) Formenctype(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("formenctype", v, dontEscape...)
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
func (element *ElementInput) Inputmode(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("inputmode", v, dontEscape...)
	return element
}

// Minlength is the "minlength" attribute.
// Minimum length of value
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementInput) Minlength(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("minlength", v, dontEscape...)
	return element
}

// Placeholder is the "placeholder" attribute.
// User-visible label to be placed within the form control
// Valid values are constrained to the following:
//   - text
func (element *ElementInput) Placeholder(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("placeholder", v, dontEscape...)
	return element
}

// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   - valid_url_potentially_surrounded_by_spaces
func (element *ElementInput) Itemid(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("itemid", v, dontEscape...)
	return element
}

// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
func (element *ElementInput) Itemref(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("itemref", v, dontEscape...)
	return element
}

// Maxlength is the "maxlength" attribute.
// Maximum length of value
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementInput) Maxlength(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("maxlength", v, dontEscape...)
	return element
}

// Formtarget is the "formtarget" attribute.
// Navigable for form submission
// Valid values are constrained to the following:
//   - valid_navigable_target_name_or_keyword
func (element *ElementInput) Formtarget(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("formtarget", v, dontEscape...)
	return element
}

// Accept is the "accept" attribute.
// Hint for expected file type in file upload controls
// Valid values are constrained to the following:
//   - set_of_comma_separated_tokens
//   - valid_mime_type_strings_with_no_parameters
//   - audio/*
//   - video/*
//   - image/*
func (element *ElementInput) Accept(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("accept", v, dontEscape...)
	return element
}

// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (element *ElementInput) Draggable(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("draggable", v, dontEscape...)
	return element
}

// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean_attribute
func (element *ElementInput) Autofocus(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("autofocus", v, dontEscape...)
	return element
}

// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   - valid_custom_element_name
//   - customized_built_in_element
func (element *ElementInput) Is(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("is", v, dontEscape...)
	return element
}

// Popovershowtarget is the "popovershowtarget" attribute.
// Shows the specified popover element when clicked
// Valid values are constrained to the following:
func (element *ElementInput) Popovershowtarget(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("popovershowtarget", v, dontEscape...)
	return element
}

// Formmethod is the "formmethod" attribute.
// Variant to use for form submission
// Valid values are constrained to the following:
//   - get
//   - post
//   - dialog
func (element *ElementInput) Formmethod(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("formmethod", v, dontEscape...)
	return element
}

// Min is the "min" attribute.
// Lower bound of range
// Valid values are constrained to the following:
//   - valid_floating_point_number
func (element *ElementInput) Min(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("min", v, dontEscape...)
	return element
}

// Width is the "width" attribute.
// Horizontal dimension
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementInput) Width(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("width", v, dontEscape...)
	return element
}

// List is the "list" attribute.
// List of autocomplete options
// Valid values are constrained to the following:
//   - id
func (element *ElementInput) List(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("list", v, dontEscape...)
	return element
}

// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   - auto
//   - auto
//   - manual
//   - manual
func (element *ElementInput) Popover(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("popover", v, dontEscape...)
	return element
}

// Type is the "type" attribute.
// Type of script
// Valid values are constrained to the following:
//   - module
//   - valid_mime_type_string
//   - java_script_mime_type_essence_match
func (element *ElementInput) Type(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("type", v, dontEscape...)
	return element
}

// Autocomplete is the "autocomplete" attribute.
// Hint for form autofill feature
// Valid values are constrained to the following:
//   - autofill_field
func (element *ElementInput) Autocomplete(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("autocomplete", v, dontEscape...)
	return element
}

// Formaction is the "formaction" attribute.
// URL to use for form submission
// Valid values are constrained to the following:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (element *ElementInput) Formaction(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("formaction", v, dontEscape...)
	return element
}

// Height is the "height" attribute.
// Vertical dimension
// Valid values are constrained to the following:
//   - valid_non_negative_integer
func (element *ElementInput) Height(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("height", v, dontEscape...)
	return element
}

// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (element *ElementInput) Itemtype(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("itemtype", v, dontEscape...)
	return element
}

// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (element *ElementInput) Hidden(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("hidden", v, dontEscape...)
	return element
}

// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (element *ElementInput) Itemprop(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("itemprop", v, dontEscape...)
	return element
}

// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   - text
func (element *ElementInput) Name(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("name", v, dontEscape...)
	return element
}

// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   - valid_integer
func (element *ElementInput) Tabindex(v string, dontEscape ...bool) *ElementInput {
	element.setAttribute("tabindex", v, dontEscape...)
	return element
}

// auxclick event handler
func (e *ElementInput) OnAuxclick(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnBeforematch(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnBeforetoggle(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnBlur(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnCancel(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnCanplay(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnCanplaythrough(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnChange(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnClick(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnClose(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnContextlost(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnContextmenu(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnContextrestored(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnCopy(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnCuechange(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnCut(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDblclick(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDrag(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDragend(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDragenter(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDragleave(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDragover(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDragstart(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDrop(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnDurationchange(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnEmptied(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnEnded(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnError(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnFocus(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnFormdata(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnInput(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnInvalid(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnKeydown(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnKeypress(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnKeyup(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnLoad(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnLoadeddata(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnLoadedmetadata(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnLoadstart(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnMousedown(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnMouseenter(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnMouseleave(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnMousemove(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnMouseout(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnMouseover(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnMouseup(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnPaste(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnPause(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnPlay(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnPlaying(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnProgress(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnRatechange(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnReset(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnResize(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnScroll(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnScrollend(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnSeeked(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnSeeking(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnSelect(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnSlotchange(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnStalled(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnSubmit(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnSuspend(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnTimeupdate(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnToggle(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnVolumechange(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnWaiting(fn engine.EventHandler) *ElementInput {
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
func (e *ElementInput) OnWheel(fn engine.EventHandler) *ElementInput {
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
