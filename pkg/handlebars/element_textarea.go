/* cSpell:disable */

package handlebars

import "github.com/gogoracer/racer/pkg/engine"

type ElementTextarea struct {
	shouldBeComponent bool
	attrs             map[string]interface{}
	children          []any
}

func TEXTAREA(children ...any) *ElementTextarea {
	return &ElementTextarea{
		attrs:    map[string]interface{}{},
		children: children,
	}
}

func (e *ElementTextarea) Add(children ...any) *ElementTextarea {
	e.children = append(e.children, children...)
	return e
}

func (e *ElementTextarea) Custom(k, v string) *ElementTextarea {
	e.attrs[k] = v
	return e
}

func (e *ElementTextarea) BindCustom(k string, v bool) *ElementTextarea {
	e.shouldBeComponent = true
	return e
}

func (e ElementTextarea) HandlebarElement() {}

func (e ElementTextarea) GenerateVDOM() interface{} {
	all := append([]any{e.attrs}, e.children...)
	if e.shouldBeComponent {
		return engine.NewComponent("textarea", all...)
	} else {
		return engine.NewTag("textarea", all...)
	}
}

// Accesskey is the "accesskey"" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered-set-of-unique-space-separated-tokens
//   - string-is
func (e *ElementTextarea) Accesskey(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["accesskey"] = v
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
func (e *ElementTextarea) Autocapitalize(v string) *ElementTextarea {
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
func (e *ElementTextarea) Autocomplete(v string) *ElementTextarea {
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
func (e *ElementTextarea) Autofocus(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["autofocus"] = v
	return e
}

// Class is the "class"" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set-of-space-separated-tokens
func (e *ElementTextarea) Class(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["class"] = v
	return e
}

// Cols is the "cols"" attribute.
// Maximum number of characters per line
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementTextarea) Cols(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["cols"] = v
	return e
}

// Contenteditable is the "contenteditable"" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   - true
//   - false
func (e *ElementTextarea) Contenteditable(v string) *ElementTextarea {
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
func (e *ElementTextarea) Dir(v string) *ElementTextarea {
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
func (e *ElementTextarea) Dirname(v string) *ElementTextarea {
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
func (e *ElementTextarea) Disabled(v string) *ElementTextarea {
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
func (e *ElementTextarea) Draggable(v string) *ElementTextarea {
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
func (e *ElementTextarea) Enterkeyhint(v string) *ElementTextarea {
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
func (e *ElementTextarea) Form(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["form"] = v
	return e
}

// Hidden is the "hidden"" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//
//	*
//	*
func (e *ElementTextarea) Hidden(v string) *ElementTextarea {
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
func (e *ElementTextarea) Id(v string) *ElementTextarea {
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
func (e *ElementTextarea) Inert(v string) *ElementTextarea {
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
func (e *ElementTextarea) Inputmode(v string) *ElementTextarea {
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
func (e *ElementTextarea) Is(v string) *ElementTextarea {
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
func (e *ElementTextarea) Itemid(v string) *ElementTextarea {
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
func (e *ElementTextarea) Itemprop(v string) *ElementTextarea {
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
func (e *ElementTextarea) Itemref(v string) *ElementTextarea {
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
func (e *ElementTextarea) Itemscope(v string) *ElementTextarea {
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
func (e *ElementTextarea) Itemtype(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["itemtype"] = v
	return e
}

// Lang is the "lang"" attribute.
//
// Valid values are constrained to the following:
func (e *ElementTextarea) Lang(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["lang"] = v
	return e
}

// Maxlength is the "maxlength"" attribute.
// Maximum
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementTextarea) Maxlength(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["maxlength"] = v
	return e
}

// Minlength is the "minlength"" attribute.
// Minimum
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementTextarea) Minlength(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["minlength"] = v
	return e
}

// Name is the "name"" attribute.
// Name of the element to use for
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementTextarea) Name(v string) *ElementTextarea {
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
func (e *ElementTextarea) Nonce(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["nonce"] = v
	return e
}

// Placeholder is the "placeholder"" attribute.
// User-visible label to be placed within the form control
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementTextarea) Placeholder(v string) *ElementTextarea {
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
func (e *ElementTextarea) Popover(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["popover"] = v
	return e
}

// Readonly is the "readonly"" attribute.
// Whether to allow the value to be edited by the user
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementTextarea) Readonly(v string) *ElementTextarea {
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
func (e *ElementTextarea) Required(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["required"] = v
	return e
}

// Rows is the "rows"" attribute.
// Number of lines to show
// Valid values are constrained to the following:
//   - valid-non-negative-integer
func (e *ElementTextarea) Rows(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["rows"] = v
	return e
}

// Slot is the "slot"" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementTextarea) Slot(v string) *ElementTextarea {
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
func (e *ElementTextarea) Spellcheck(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["spellcheck"] = v
	return e
}

// Style is the "style"" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *ElementTextarea) Style(v string) *ElementTextarea {
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
func (e *ElementTextarea) Tabindex(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["tabindex"] = v
	return e
}

// Title is the "title"" attribute.
// Advisory information for the element
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementTextarea) Title(v string) *ElementTextarea {
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
func (e *ElementTextarea) Translate(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["translate"] = v
	return e
}

// Wrap is the "wrap"" attribute.
// How the value of the form control is to be wrapped for
// Valid values are constrained to the following:
//
//	*
//	*
func (e *ElementTextarea) Wrap(v string) *ElementTextarea {
	if v == "" {
		return e
	}
	e.attrs["wrap"] = v
	return e
}

// &lt;code id=&quot;attributes-3:event-auxclick&quot;&gt;&lt;a data-x-internal=&quot;event-auxclick&quot; href=&quot;https://w3c.github.io/uievents/#event-type-auxclick&quot;&gt;auxclick&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-beforematch&quot;&gt;&lt;a href=&quot;#event-beforematch&quot;&gt;beforematch&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-beforetoggle&quot;&gt;&lt;a href=&quot;#event-beforetoggle&quot;&gt;beforetoggle&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-blur&quot;&gt;&lt;a href=&quot;#event-blur&quot;&gt;blur&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-cancel&quot;&gt;&lt;a href=&quot;#event-cancel&quot;&gt;cancel&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-canplay&quot;&gt;&lt;a href=&quot;media.html#event-media-canplay&quot;&gt;canplay&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-canplaythrough&quot;&gt;&lt;a href=&quot;media.html#event-media-canplaythrough&quot;&gt;canplaythrough&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-change&quot;&gt;&lt;a href=&quot;#event-change&quot;&gt;change&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-click&quot;&gt;&lt;a data-x-internal=&quot;event-click&quot; href=&quot;https://w3c.github.io/uievents/#event-type-click&quot;&gt;click&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-close&quot;&gt;&lt;a href=&quot;#event-close&quot;&gt;close&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-contextlost&quot;&gt;&lt;a href=&quot;#event-contextlost&quot;&gt;contextlost&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-contextmenu&quot;&gt;&lt;a data-x-internal=&quot;event-contextmenu&quot; href=&quot;https://w3c.github.io/uievents/#event-type-contextmenu&quot;&gt;contextmenu&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-contextrestored&quot;&gt;&lt;a href=&quot;#event-contextrestored&quot;&gt;contextrestored&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-copy&quot;&gt;&lt;a data-x-internal=&quot;event-copy&quot; href=&quot;https://w3c.github.io/clipboard-apis/#clipboard-event-copy&quot;&gt;copy&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-cuechange&quot;&gt;&lt;a href=&quot;media.html#event-media-cuechange&quot;&gt;cuechange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-cut&quot;&gt;&lt;a data-x-internal=&quot;event-cut&quot; href=&quot;https://w3c.github.io/clipboard-apis/#clipboard-event-cut&quot;&gt;cut&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dblclick&quot;&gt;&lt;a data-x-internal=&quot;event-dblclick&quot; href=&quot;https://w3c.github.io/uievents/#event-type-dblclick&quot;&gt;dblclick&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-drag&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-drag&quot;&gt;drag&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragend&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragend&quot;&gt;dragend&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragenter&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragenter&quot;&gt;dragenter&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragleave&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragleave&quot;&gt;dragleave&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragover&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragover&quot;&gt;dragover&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-dragstart&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-dragstart&quot;&gt;dragstart&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-dnd-drop&quot;&gt;&lt;a href=&quot;dnd.html#event-dnd-drop&quot;&gt;drop&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-durationchange&quot;&gt;&lt;a href=&quot;media.html#event-media-durationchange&quot;&gt;durationchange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-emptied&quot;&gt;&lt;a href=&quot;media.html#event-media-emptied&quot;&gt;emptied&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-ended&quot;&gt;&lt;a href=&quot;media.html#event-media-ended&quot;&gt;ended&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-error&quot;&gt;&lt;a href=&quot;#event-error&quot;&gt;error&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-focus&quot;&gt;&lt;a href=&quot;#event-focus&quot;&gt;focus&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-formdata&quot;&gt;&lt;a href=&quot;#event-formdata&quot;&gt;formdata&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-input&quot;&gt;&lt;a data-x-internal=&quot;event-input&quot; href=&quot;https://w3c.github.io/uievents/#event-type-input&quot;&gt;input&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-invalid&quot;&gt;&lt;a href=&quot;#event-invalid&quot;&gt;invalid&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-keydown&quot;&gt;&lt;a data-x-internal=&quot;event-keydown&quot; href=&quot;https://w3c.github.io/uievents/#event-type-keydown&quot;&gt;keydown&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-keypress&quot;&gt;&lt;a data-x-internal=&quot;event-keypress&quot; href=&quot;https://w3c.github.io/uievents/#event-type-keypress&quot;&gt;keypress&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-keyup&quot;&gt;&lt;a data-x-internal=&quot;event-keyup&quot; href=&quot;https://w3c.github.io/uievents/#event-type-keyup&quot;&gt;keyup&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-load&quot;&gt;&lt;a href=&quot;#event-load&quot;&gt;load&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-loadeddata&quot;&gt;&lt;a href=&quot;media.html#event-media-loadeddata&quot;&gt;loadeddata&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-loadedmetadata&quot;&gt;&lt;a href=&quot;media.html#event-media-loadedmetadata&quot;&gt;loadedmetadata&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-loadstart&quot;&gt;&lt;a href=&quot;media.html#event-media-loadstart&quot;&gt;loadstart&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mousedown&quot;&gt;&lt;a data-x-internal=&quot;event-mousedown&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mousedown&quot;&gt;mousedown&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseenter&quot;&gt;&lt;a data-x-internal=&quot;event-mouseenter&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseenter&quot;&gt;mouseenter&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseleave&quot;&gt;&lt;a data-x-internal=&quot;event-mouseleave&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseleave&quot;&gt;mouseleave&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mousemove&quot;&gt;&lt;a data-x-internal=&quot;event-mousemove&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mousemove&quot;&gt;mousemove&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseout&quot;&gt;&lt;a data-x-internal=&quot;event-mouseout&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseout&quot;&gt;mouseout&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseover&quot;&gt;&lt;a data-x-internal=&quot;event-mouseover&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseover&quot;&gt;mouseover&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-mouseup&quot;&gt;&lt;a data-x-internal=&quot;event-mouseup&quot; href=&quot;https://w3c.github.io/uievents/#event-type-mouseup&quot;&gt;mouseup&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-paste&quot;&gt;&lt;a data-x-internal=&quot;event-paste&quot; href=&quot;https://w3c.github.io/clipboard-apis/#clipboard-event-paste&quot;&gt;paste&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-pause&quot;&gt;&lt;a href=&quot;media.html#event-media-pause&quot;&gt;pause&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-play&quot;&gt;&lt;a href=&quot;media.html#event-media-play&quot;&gt;play&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-playing&quot;&gt;&lt;a href=&quot;media.html#event-media-playing&quot;&gt;playing&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-progress&quot;&gt;&lt;a href=&quot;media.html#event-media-progress&quot;&gt;progress&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-ratechange&quot;&gt;&lt;a href=&quot;media.html#event-media-ratechange&quot;&gt;ratechange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-reset&quot;&gt;&lt;a href=&quot;#event-reset&quot;&gt;reset&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-resize&quot;&gt;&lt;a data-x-internal=&quot;event-resize&quot; href=&quot;https://drafts.csswg.org/cssom-view/#eventdef-window-resize&quot;&gt;resize&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-scroll&quot;&gt;&lt;a data-x-internal=&quot;event-scroll&quot; href=&quot;https://drafts.csswg.org/cssom-view/#eventdef-document-scroll&quot;&gt;scroll&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-scrollend&quot;&gt;&lt;a data-x-internal=&quot;event-scrollend&quot; href=&quot;https://drafts.csswg.org/cssom-view/#eventdef-document-scrollend&quot;&gt;scrollend&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-securitypolicyviolation&quot;&gt;&lt;a data-x-internal=&quot;event-securitypolicyviolation&quot; href=&quot;https://w3c.github.io/webappsec-csp/#eventdef-globaleventhandlers-securitypolicyviolation&quot;&gt;securitypolicyviolation&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-seeked&quot;&gt;&lt;a href=&quot;media.html#event-media-seeked&quot;&gt;seeked&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-seeking&quot;&gt;&lt;a href=&quot;media.html#event-media-seeking&quot;&gt;seeking&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-select&quot;&gt;&lt;a href=&quot;#event-select&quot;&gt;select&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-slotchange&quot;&gt;&lt;a data-x-internal=&quot;event-slotchange&quot; href=&quot;https://dom.spec.whatwg.org/#eventdef-htmlslotelement-slotchange&quot;&gt;slotchange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-stalled&quot;&gt;&lt;a href=&quot;media.html#event-media-stalled&quot;&gt;stalled&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-submit&quot;&gt;&lt;a href=&quot;#event-submit&quot;&gt;submit&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-suspend&quot;&gt;&lt;a href=&quot;media.html#event-media-suspend&quot;&gt;suspend&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-timeupdate&quot;&gt;&lt;a href=&quot;media.html#event-media-timeupdate&quot;&gt;timeupdate&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-toggle&quot;&gt;&lt;a href=&quot;#event-toggle&quot;&gt;toggle&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-volumechange&quot;&gt;&lt;a href=&quot;media.html#event-media-volumechange&quot;&gt;volumechange&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-media-waiting&quot;&gt;&lt;a href=&quot;media.html#event-media-waiting&quot;&gt;waiting&lt;/a&gt;&lt;/code&gt;  event handler
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

// &lt;code id=&quot;attributes-3:event-wheel&quot;&gt;&lt;a data-x-internal=&quot;event-wheel&quot; href=&quot;https://w3c.github.io/uievents/#event-type-wheel&quot;&gt;wheel&lt;/a&gt;&lt;/code&gt;  event handler
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
