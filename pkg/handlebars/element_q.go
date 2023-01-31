/* cSpell:disable */

package handlebars

import "github.com/gogoracer/racer/pkg/engine"

type ElementQ struct {
	shouldBeComponent bool
	attrs             map[string]interface{}
	children          []any
}

func Q(children ...any) *ElementQ {
	return &ElementQ{
		attrs:    map[string]interface{}{},
		children: children,
	}
}

func (e *ElementQ) Add(children ...any) *ElementQ {
	e.children = append(e.children, children...)
	return e
}

func (e *ElementQ) Custom(k, v string) *ElementQ {
	e.attrs[k] = v
	return e
}

func (e *ElementQ) BindCustom(k string, v bool) *ElementQ {
	e.shouldBeComponent = true
	return e
}

func (e ElementQ) HandlebarElement() {}

func (e ElementQ) GenerateVDOM() interface{} {
	all := append([]any{e.attrs}, e.children...)
	if e.shouldBeComponent {
		return engine.NewComponent("q", all...)
	} else {
		return engine.NewTag("q", all...)
	}
}

// Accesskey is the "accesskey"" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered-set-of-unique-space-separated-tokens
//   - string-is
func (e *ElementQ) Accesskey(v string) *ElementQ {
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
func (e *ElementQ) Autocapitalize(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["autocapitalize"] = v
	return e
}

// Autofocus is the "autofocus"" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   - boolean-attribute
func (e *ElementQ) Autofocus(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["autofocus"] = v
	return e
}

// Cite is the "cite"" attribute.
// Link to the source of the quotation or more information about the edit
// Valid values are constrained to the following:
//   - valid-url-potentially-surrounded-by-spaces
func (e *ElementQ) Cite(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["cite"] = v
	return e
}

// Class is the "class"" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   - set-of-space-separated-tokens
func (e *ElementQ) Class(v string) *ElementQ {
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
func (e *ElementQ) Contenteditable(v string) *ElementQ {
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
func (e *ElementQ) Dir(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["dir"] = v
	return e
}

// Draggable is the "draggable"" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   - true
//   - false
func (e *ElementQ) Draggable(v string) *ElementQ {
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
func (e *ElementQ) Enterkeyhint(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["enterkeyhint"] = v
	return e
}

// Hidden is the "hidden"" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//
//	*
//	*
func (e *ElementQ) Hidden(v string) *ElementQ {
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
func (e *ElementQ) Id(v string) *ElementQ {
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
func (e *ElementQ) Inert(v string) *ElementQ {
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
func (e *ElementQ) Inputmode(v string) *ElementQ {
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
func (e *ElementQ) Is(v string) *ElementQ {
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
func (e *ElementQ) Itemid(v string) *ElementQ {
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
func (e *ElementQ) Itemprop(v string) *ElementQ {
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
func (e *ElementQ) Itemref(v string) *ElementQ {
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
func (e *ElementQ) Itemscope(v string) *ElementQ {
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
func (e *ElementQ) Itemtype(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["itemtype"] = v
	return e
}

// Lang is the "lang"" attribute.
//
// Valid values are constrained to the following:
func (e *ElementQ) Lang(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["lang"] = v
	return e
}

// Nonce is the "nonce"" attribute.
// Cryptographic nonce used in
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementQ) Nonce(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["nonce"] = v
	return e
}

// Popover is the "popover"" attribute.
// Makes the element a
// Valid values are constrained to the following:
//
//	*
//	*
func (e *ElementQ) Popover(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["popover"] = v
	return e
}

// Slot is the "slot"" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   - attribute-text
func (e *ElementQ) Slot(v string) *ElementQ {
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
func (e *ElementQ) Spellcheck(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["spellcheck"] = v
	return e
}

// Style is the "style"" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *ElementQ) Style(v string) *ElementQ {
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
func (e *ElementQ) Tabindex(v string) *ElementQ {
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
func (e *ElementQ) Title(v string) *ElementQ {
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
func (e *ElementQ) Translate(v string) *ElementQ {
	if v == "" {
		return e
	}
	e.attrs["translate"] = v
	return e
}

// &lt;code id=&quot;attributes-3:event-auxclick&quot;&gt;&lt;a data-x-internal=&quot;event-auxclick&quot; href=&quot;https://w3c.github.io/uievents/#event-type-auxclick&quot;&gt;auxclick&lt;/a&gt;&lt;/code&gt;  event handler
func (e *ElementQ) OnAuxclick(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnBeforematch(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnBeforetoggle(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnBlur(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnCancel(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnCanplay(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnCanplaythrough(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnChange(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnClick(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnClose(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnContextlost(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnContextmenu(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnContextrestored(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnCopy(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnCuechange(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnCut(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDblclick(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDrag(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDragend(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDragenter(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDragleave(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDragover(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDragstart(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDrop(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnDurationchange(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnEmptied(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnEnded(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnError(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnFocus(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnFormdata(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnInput(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnInvalid(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnKeydown(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnKeypress(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnKeyup(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnLoad(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnLoadeddata(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnLoadedmetadata(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnLoadstart(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnMousedown(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnMouseenter(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnMouseleave(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnMousemove(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnMouseout(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnMouseover(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnMouseup(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnPaste(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnPause(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnPlay(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnPlaying(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnProgress(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnRatechange(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnReset(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnResize(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnScroll(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnScrollend(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnSeeked(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnSeeking(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnSelect(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnSlotchange(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnStalled(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnSubmit(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnSuspend(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnTimeupdate(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnToggle(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnVolumechange(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnWaiting(fn engine.EventHandler) *ElementQ {
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
func (e *ElementQ) OnWheel(fn engine.EventHandler) *ElementQ {
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
