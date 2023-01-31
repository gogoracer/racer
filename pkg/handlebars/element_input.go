/* cSpell:disable */

package handlebars

import "github.com/gogoracer/racer/pkg/engine"

type ElementInput struct {
	shouldBeComponent bool
	attrs             map[string]interface{}
	children          []any
}

func INPUT(children ...any) *ElementInput {
	return &ElementInput{
		attrs:    map[string]interface{}{},
		children: children,
	}
}

func (e *ElementInput) Add(children ...any) *ElementInput {
	e.children = append(e.children, children...)
	return e
}

func (e *ElementInput) Custom(k, v string) *ElementInput {
	e.attrs[k] = v
	return e
}

func (e *ElementInput) BindCustom(k string, v bool) *ElementInput {
	e.shouldBeComponent = true
	return e
}

func (e ElementInput) HandlebarElement() {}

func (e ElementInput) GenerateVDOM() interface{} {
	all := append([]any{e.attrs}, e.children...)
	if e.shouldBeComponent {
		return engine.NewComponent("input", all...)
	} else {
		return engine.NewTag("input", all...)
	}
}

// Accept is the "accept"" attribute.
// Hint for expected file type in
// Valid values are constrained to the following:
//   - set-of-comma-separated-tokens
//   - valid-mime-type-with-no-parameters
//   - audio/*
//   - video/*
//   - image/*
func (e *ElementInput) Accept(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["accept"] = v
	return e
}

// Accesskey is the "accesskey"" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered-set-of-unique-space-separated-tokens
//   - string-is
func (e *ElementInput) Accesskey(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["accesskey"] = v
	return e
}

// Alt is the "alt"" attribute.
// Replacement text for use when images are not available
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementInput) Alt(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["alt"] = v
	return e
}

// Autocapitalize is the "autocapitalize"" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Valid values are constrained to the following:
//
//	*
//	*
//	*
//	*
//	*
//	*
func (e *ElementInput) Autocapitalize(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["autocapitalize"] = v
	return e
}

// Autocomplete is the "autocomplete"" attribute.
// Hint for form autofill feature
// Valid values are constrained to the following:
//   - autofill-field
func (e *ElementInput) Autocomplete(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["autocomplete"] = v
	return e
}

// Autofocus is the "autofocus"" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Autofocus(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["autofocus"] = v
	return e
}

// Checked is the "checked"" attribute.
// Whether the control is checked
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Checked(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["checked"] = v
	return e
}

// Class is the "class"" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set-of-space-separated-tokens
func (e *ElementInput) Class(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["class"] = v
	return e
}

// Contenteditable is the "contenteditable"" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (e *ElementInput) Contenteditable(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["contenteditable"] = v
	return e
}

// Dir is the "dir"" attribute.
//
// Valid values are constrained to the following:
//
//	*
//	*
//	*
func (e *ElementInput) Dir(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["dir"] = v
	return e
}

// Dirname is the "dirname"" attribute.
// Name of form control to use for sending the element&#39;s
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementInput) Dirname(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["dirname"] = v
	return e
}

// Disabled is the "disabled"" attribute.
// Whether the form control is disabled
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Disabled(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["disabled"] = v
	return e
}

// Draggable is the "draggable"" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (e *ElementInput) Draggable(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["draggable"] = v
	return e
}

// Enterkeyhint is the "enterkeyhint"" attribute.
// Hint for selecting an enter key action
// Valid values are constrained to the following:
//
//	*
//	*
//	*
//	*
//	*
//	*
//	*
func (e *ElementInput) Enterkeyhint(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["enterkeyhint"] = v
	return e
}

// Form is the "form"" attribute.
// Associates the element with a
// Valid values are constrained to the following:
//   - concept-id
func (e *ElementInput) Form(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["form"] = v
	return e
}

// Formaction is the "formaction"" attribute.
//
// Valid values are constrained to the following:
//   - valid-non-empty-url-potentially-surrounded-by-spaces
func (e *ElementInput) Formaction(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["formaction"] = v
	return e
}

// Formenctype is the "formenctype"" attribute.
//
// Valid values are constrained to the following:
//
//	*
//	*
//	*
func (e *ElementInput) Formenctype(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["formenctype"] = v
	return e
}

// Formmethod is the "formmethod"" attribute.
// Variant to use for
// Valid values are constrained to the following:
//   - GET
//   - POST
//   - dialog
func (e *ElementInput) Formmethod(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["formmethod"] = v
	return e
}

// Formnovalidate is the "formnovalidate"" attribute.
// Bypass form control validation for
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Formnovalidate(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["formnovalidate"] = v
	return e
}

// Formtarget is the "formtarget"" attribute.
//
// Valid values are constrained to the following:
//   - valid-navigable-target-name-or-keyword
func (e *ElementInput) Formtarget(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["formtarget"] = v
	return e
}

// Height is the "height"" attribute.
// Vertical dimension
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementInput) Height(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["height"] = v
	return e
}

// Hidden is the "hidden"" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//
//	*
//	*
func (e *ElementInput) Hidden(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["hidden"] = v
	return e
}

// Id is the "id"" attribute.
// The element&#39;s
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementInput) Id(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["id"] = v
	return e
}

// Inert is the "inert"" attribute.
// Whether the element is
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Inert(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["inert"] = v
	return e
}

// Inputmode is the "inputmode"" attribute.
// Hint for selecting an input modality
// Valid values are constrained to the following:
//
//	*
//	*
//	*
//	*
//	*
//	*
//	*
//	*
func (e *ElementInput) Inputmode(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["inputmode"] = v
	return e
}

// Is is the "is"" attribute.
// Creates a
// Valid values are constrained to the following:
//   - valid-custom-element-name
//   - customized-built-in-element
func (e *ElementInput) Is(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["is"] = v
	return e
}

// Itemid is the "itemid"" attribute.
//
// Valid values are constrained to the following:
//   - valid-url-potentially-surrounded-by-spaces
func (e *ElementInput) Itemid(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["itemid"] = v
	return e
}

// Itemprop is the "itemprop"" attribute.
//
// Valid values are constrained to the following:
//   - unordered-set-of-unique-space-separated-tokens
//   - syntax-url-absolute
//   - defined-property-name
func (e *ElementInput) Itemprop(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["itemprop"] = v
	return e
}

// Itemref is the "itemref"" attribute.
//
// Valid values are constrained to the following:
//   - unordered-set-of-unique-space-separated-tokens
func (e *ElementInput) Itemref(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["itemref"] = v
	return e
}

// Itemscope is the "itemscope"" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Itemscope(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["itemscope"] = v
	return e
}

// Itemtype is the "itemtype"" attribute.
//
// Valid values are constrained to the following:
//   - unordered-set-of-unique-space-separated-tokens
//   - syntax-url-absolute
func (e *ElementInput) Itemtype(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["itemtype"] = v
	return e
}

// Lang is the "lang"" attribute.
//
// Valid values are constrained to the following:
func (e *ElementInput) Lang(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["lang"] = v
	return e
}

// List is the "list"" attribute.
// List of autocomplete options
// Valid values are constrained to the following:
//   - concept-id
func (e *ElementInput) List(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["list"] = v
	return e
}

// Max is the "max"" attribute.
// Maximum value
// Valid values are constrained to the following:
func (e *ElementInput) Max(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["max"] = v
	return e
}

// Maxlength is the "maxlength"" attribute.
// Maximum
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementInput) Maxlength(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["maxlength"] = v
	return e
}

// Min is the "min"" attribute.
// Minimum value
// Valid values are constrained to the following:
func (e *ElementInput) Min(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["min"] = v
	return e
}

// Minlength is the "minlength"" attribute.
// Minimum
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementInput) Minlength(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["minlength"] = v
	return e
}

// Multiple is the "multiple"" attribute.
// Whether to allow multiple values
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Multiple(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["multiple"] = v
	return e
}

// Name is the "name"" attribute.
// Name of the element to use for
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementInput) Name(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["name"] = v
	return e
}

// Nonce is the "nonce"" attribute.
// Cryptographic nonce used in
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementInput) Nonce(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["nonce"] = v
	return e
}

// Pattern is the "pattern"" attribute.
// Pattern to be matched by the form control&#39;s value
// Valid values are constrained to the following:
func (e *ElementInput) Pattern(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["pattern"] = v
	return e
}

// Placeholder is the "placeholder"" attribute.
// User-visible label to be placed within the form control
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementInput) Placeholder(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["placeholder"] = v
	return e
}

// Popover is the "popover"" attribute.
// Makes the element a
// Valid values are constrained to the following:
//
//	*
//	*
func (e *ElementInput) Popover(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["popover"] = v
	return e
}

// Popoverhidetarget is the "popoverhidetarget"" attribute.
// Hides the specified
// Valid values are constrained to the following:
func (e *ElementInput) Popoverhidetarget(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["popoverhidetarget"] = v
	return e
}

// Popovershowtarget is the "popovershowtarget"" attribute.
// Shows the specified
// Valid values are constrained to the following:
func (e *ElementInput) Popovershowtarget(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["popovershowtarget"] = v
	return e
}

// Popovertoggletarget is the "popovertoggletarget"" attribute.
// Toggles the specified
// Valid values are constrained to the following:
func (e *ElementInput) Popovertoggletarget(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["popovertoggletarget"] = v
	return e
}

// Readonly is the "readonly"" attribute.
// Whether to allow the value to be edited by the user
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Readonly(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["readonly"] = v
	return e
}

// Required is the "required"" attribute.
// Whether the control is required for
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementInput) Required(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["required"] = v
	return e
}

// Size is the "size"" attribute.
// Size of the control
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementInput) Size(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["size"] = v
	return e
}

// Slot is the "slot"" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementInput) Slot(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["slot"] = v
	return e
}

// Spellcheck is the "spellcheck"" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   - true
//   - false
func (e *ElementInput) Spellcheck(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["spellcheck"] = v
	return e
}

// Src is the "src"" attribute.
// Address of the resource
// Valid values are constrained to the following:
//   - valid-non-empty-url-potentially-surrounded-by-spaces
func (e *ElementInput) Src(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["src"] = v
	return e
}

// Step is the "step"" attribute.
// Granularity to be matched by the form control&#39;s value
// Valid values are constrained to the following:
//   - valid-floating-point-number
//   - any
func (e *ElementInput) Step(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["step"] = v
	return e
}

// Style is the "style"" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *ElementInput) Style(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["style"] = v
	return e
}

// Tabindex is the "tabindex"" attribute.
// Whether the element is
// Valid values are constrained to the following:
//   - valid-integer
func (e *ElementInput) Tabindex(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["tabindex"] = v
	return e
}

// Title is the "title"" attribute.
// Description of pattern (when used with
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementInput) Title(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["title"] = v
	return e
}

// Translate is the "translate"" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   - yes
//   - no
func (e *ElementInput) Translate(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["translate"] = v
	return e
}

// Type is the "type"" attribute.
// Type of form control
// Valid values are constrained to the following:
//   - attr-input-type
func (e *ElementInput) Type(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["type"] = v
	return e
}

// Value is the "value"" attribute.
// Value of the form control
// Valid values are constrained to the following:
func (e *ElementInput) Value(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["value"] = v
	return e
}

// Width is the "width"" attribute.
// Horizontal dimension
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementInput) Width(v string) *ElementInput {
	if v == "" {
		return e
	}
	e.attrs["width"] = v
	return e
}

// &lt;code id=&quot;attributes-3:event-auxclick&quot;&gt;&lt;a data-x-internal=&quot;event-auxclick&quot; href=&quot;https://w3c.github.io/uievents/#event-type-auxclick&quot;&gt;auxclick&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-beforematch&quot;&gt;&lt;a href=&quot;#event-beforematch&quot;&gt;beforematch&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-beforetoggle&quot;&gt;&lt;a href=&quot;#event-beforetoggle&quot;&gt;beforetoggle&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-blur&quot;&gt;&lt;a href=&quot;#event-blur&quot;&gt;blur&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-cancel&quot;&gt;&lt;a href=&quot;#event-cancel&quot;&gt;cancel&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-canplay&quot;&gt;&lt;a href=&quot;media.html#event-media-canplay&quot;&gt;canplay&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-canplaythrough&quot;&gt;&lt;a href=&quot;media.html#event-media-canplaythrough&quot;&gt;canplaythrough&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-change&quot;&gt;&lt;a href=&quot;#event-change&quot;&gt;change&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-click&quot;&gt;&lt;a data-x-internal=&quot;event-click&quot; href=&quot;https://w3c.github.io/uievents/#event-type-click&quot;&gt;click&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-close&quot;&gt;&lt;a href=&quot;#event-close&quot;&gt;close&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-contextlost&quot;&gt;&lt;a href=&quot;#event-contextlost&quot;&gt;contextlost&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-contextmenu&quot;&gt;&lt;a data-x-internal=&quot;event-contextmenu&quot; href=&quot;https://w3c.github.io/uievents/#event-type-contextmenu&quot;&gt;contextmenu&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-contextrestored&quot;&gt;&lt;a href=&quot;#event-contextrestored&quot;&gt;contextrestored&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-copy&quot;&gt;&lt;a data-x-internal=&quot;event-copy&quot; href=&quot;https://w3c.github.io/clipboard-apis/#clipboard-event-copy&quot;&gt;copy&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-cuechange&quot;&gt;&lt;a href=&quot;media.html#event-media-cuechange&quot;&gt;cuechange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-cut&quot;&gt;&lt;a data-x-internal=&quot;event-cut&quot; href=&quot;https://w3c.github.io/clipboard-apis/#clipboard-event-cut&quot;&gt;cut&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dblclick&quot;&gt;&lt;a data-x-internal=&quot;event-dblclick&quot; href=&quot;https://w3c.github.io/uievents/#event-type-dblclick&quot;&gt;dblclick&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-drag&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-drag&quot;&gt;drag&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragend&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragend&quot;&gt;dragend&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragenter&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragenter&quot;&gt;dragenter&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragleave&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragleave&quot;&gt;dragleave&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragover&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragover&quot;&gt;dragover&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragstart&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragstart&quot;&gt;dragstart&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-drop&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-drop&quot;&gt;drop&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-durationchange&quot;&gt;&lt;a href=&quot;media.html#event-media-durationchange&quot;&gt;durationchange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-emptied&quot;&gt;&lt;a href=&quot;media.html#event-media-emptied&quot;&gt;emptied&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-ended&quot;&gt;&lt;a href=&quot;media.html#event-media-ended&quot;&gt;ended&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-error&quot;&gt;&lt;a href=&quot;#event-error&quot;&gt;error&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-focus&quot;&gt;&lt;a href=&quot;#event-focus&quot;&gt;focus&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-formdata&quot;&gt;&lt;a href=&quot;#event-formdata&quot;&gt;formdata&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-input&quot;&gt;&lt;a data-x-internal=&quot;event-input&quot; href=&quot;https://w3c.github.io/uievents/#event-type-input&quot;&gt;input&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-invalid&quot;&gt;&lt;a href=&quot;#event-invalid&quot;&gt;invalid&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-keydown&quot;&gt;&lt;a data-x-internal=&quot;event-keydown&quot; href=&quot;https://w3c.github.io/uievents/#event-type-keydown&quot;&gt;keydown&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-keypress&quot;&gt;&lt;a data-x-internal=&quot;event-keypress&quot; href=&quot;https://w3c.github.io/uievents/#event-type-keypress&quot;&gt;keypress&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-keyup&quot;&gt;&lt;a data-x-internal=&quot;event-keyup&quot; href=&quot;https://w3c.github.io/uievents/#event-type-keyup&quot;&gt;keyup&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-load&quot;&gt;&lt;a href=&quot;#event-load&quot;&gt;load&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-loadeddata&quot;&gt;&lt;a href=&quot;media.html#event-media-loadeddata&quot;&gt;loadeddata&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-loadedmetadata&quot;&gt;&lt;a href=&quot;media.html#event-media-loadedmetadata&quot;&gt;loadedmetadata&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-loadstart&quot;&gt;&lt;a href=&quot;media.html#event-media-loadstart&quot;&gt;loadstart&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mousedown&quot;&gt;&lt;a data-x-internal=&quot;event-mousedown&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mousedown&quot;&gt;mousedown&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseenter&quot;&gt;&lt;a data-x-internal=&quot;event-mouseenter&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseenter&quot;&gt;mouseenter&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseleave&quot;&gt;&lt;a data-x-internal=&quot;event-mouseleave&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseleave&quot;&gt;mouseleave&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mousemove&quot;&gt;&lt;a data-x-internal=&quot;event-mousemove&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mousemove&quot;&gt;mousemove&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseout&quot;&gt;&lt;a data-x-internal=&quot;event-mouseout&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseout&quot;&gt;mouseout&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseover&quot;&gt;&lt;a data-x-internal=&quot;event-mouseover&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseover&quot;&gt;mouseover&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseup&quot;&gt;&lt;a data-x-internal=&quot;event-mouseup&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseup&quot;&gt;mouseup&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-paste&quot;&gt;&lt;a data-x-internal=&quot;event-paste&quot; href=&quot;https://w3c.github.io/clipboard-apis/#clipboard-event-paste&quot;&gt;paste&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-pause&quot;&gt;&lt;a href=&quot;media.html#event-media-pause&quot;&gt;pause&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-play&quot;&gt;&lt;a href=&quot;media.html#event-media-play&quot;&gt;play&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-playing&quot;&gt;&lt;a href=&quot;media.html#event-media-playing&quot;&gt;playing&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-progress&quot;&gt;&lt;a href=&quot;media.html#event-media-progress&quot;&gt;progress&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-ratechange&quot;&gt;&lt;a href=&quot;media.html#event-media-ratechange&quot;&gt;ratechange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-reset&quot;&gt;&lt;a href=&quot;#event-reset&quot;&gt;reset&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-resize&quot;&gt;&lt;a data-x-internal=&quot;event-resize&quot; href=&quot;https://drafts.csswg.org/cssom-view/#eventdef-window-resize&quot;&gt;resize&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-scroll&quot;&gt;&lt;a data-x-internal=&quot;event-scroll&quot; href=&quot;https://drafts.csswg.org/cssom-view/#eventdef-document-scroll&quot;&gt;scroll&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-scrollend&quot;&gt;&lt;a data-x-internal=&quot;event-scrollend&quot; href=&quot;https://drafts.csswg.org/cssom-view/#eventdef-document-scrollend&quot;&gt;scrollend&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-securitypolicyviolation&quot;&gt;&lt;a data-x-internal=&quot;event-securitypolicyviolation&quot; href=&quot;https://w3c.github.io/webappsec-csp/#eventdef-globaleventhandlers-securitypolicyviolation&quot;&gt;securitypolicyviolation&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-seeked&quot;&gt;&lt;a href=&quot;media.html#event-media-seeked&quot;&gt;seeked&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-seeking&quot;&gt;&lt;a href=&quot;media.html#event-media-seeking&quot;&gt;seeking&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-select&quot;&gt;&lt;a href=&quot;#event-select&quot;&gt;select&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-slotchange&quot;&gt;&lt;a data-x-internal=&quot;event-slotchange&quot; href=&quot;https://dom.spec.whatwg.org/#eventdef-htmlslotelement-slotchange&quot;&gt;slotchange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-stalled&quot;&gt;&lt;a href=&quot;media.html#event-media-stalled&quot;&gt;stalled&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-submit&quot;&gt;&lt;a href=&quot;#event-submit&quot;&gt;submit&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-suspend&quot;&gt;&lt;a href=&quot;media.html#event-media-suspend&quot;&gt;suspend&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-timeupdate&quot;&gt;&lt;a href=&quot;media.html#event-media-timeupdate&quot;&gt;timeupdate&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-toggle&quot;&gt;&lt;a href=&quot;#event-toggle&quot;&gt;toggle&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-volumechange&quot;&gt;&lt;a href=&quot;media.html#event-media-volumechange&quot;&gt;volumechange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-waiting&quot;&gt;&lt;a href=&quot;media.html#event-media-waiting&quot;&gt;waiting&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-wheel&quot;&gt;&lt;a data-x-internal=&quot;event-wheel&quot; href=&quot;https://w3c.github.io/uievents/#event-type-wheel&quot;&gt;wheel&lt;/a&gt;&lt;/code&gt;  event handler
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
