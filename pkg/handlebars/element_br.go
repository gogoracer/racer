/* cSpell:disable */

package handlebars

import "github.com/gogoracer/racer/pkg/engine"

type ElementBr struct {
	shouldBeComponent bool
	attrs             map[string]interface{}
	children          []any
}

func BR(children ...any) *ElementBr {
	return &ElementBr{
		attrs:    map[string]interface{}{},
		children: children,
	}
}

func (e *ElementBr) Add(children ...any) *ElementBr {
	e.children = append(e.children, children...)
	return e
}

func (e *ElementBr) Custom(k, v string) *ElementBr {
	e.attrs[k] = v
	return e
}

func (e *ElementBr) BindCustom(k string, v bool) *ElementBr {
	e.shouldBeComponent = true
	return e
}

func (e ElementBr) HandlebarElement() {}

func (e ElementBr) GenerateVDOM() interface{} {
	all := append([]any{e.attrs}, e.children...)
	if e.shouldBeComponent {
		return engine.NewComponent("br", all...)
	} else {
		return engine.NewTag("br", all...)
	}
}

// Accesskey is the "accesskey"" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   - ordered-set-of-unique-space-separated-tokens
//   - string-is
func (e *ElementBr) Accesskey(v string) *ElementBr {
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
func (e *ElementBr) Autocapitalize(v string) *ElementBr {
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
func (e *ElementBr) Autofocus(v string) *ElementBr {
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
func (e *ElementBr) Class(v string) *ElementBr {
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
func (e *ElementBr) Contenteditable(v string) *ElementBr {
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
func (e *ElementBr) Dir(v string) *ElementBr {
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
func (e *ElementBr) Draggable(v string) *ElementBr {
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
func (e *ElementBr) Enterkeyhint(v string) *ElementBr {
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
func (e *ElementBr) Hidden(v string) *ElementBr {
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
func (e *ElementBr) Id(v string) *ElementBr {
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
func (e *ElementBr) Inert(v string) *ElementBr {
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
func (e *ElementBr) Inputmode(v string) *ElementBr {
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
func (e *ElementBr) Is(v string) *ElementBr {
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
func (e *ElementBr) Itemid(v string) *ElementBr {
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
func (e *ElementBr) Itemprop(v string) *ElementBr {
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
func (e *ElementBr) Itemref(v string) *ElementBr {
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
func (e *ElementBr) Itemscope(v string) *ElementBr {
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
func (e *ElementBr) Itemtype(v string) *ElementBr {
	if v == "" {
		return e
	}
	e.attrs["itemtype"] = v
	return e
}

// Lang is the "lang"" attribute.
//
// Valid values are constrained to the following:
func (e *ElementBr) Lang(v string) *ElementBr {
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
func (e *ElementBr) Nonce(v string) *ElementBr {
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
func (e *ElementBr) Popover(v string) *ElementBr {
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
func (e *ElementBr) Slot(v string) *ElementBr {
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
func (e *ElementBr) Spellcheck(v string) *ElementBr {
	if v == "" {
		return e
	}
	e.attrs["spellcheck"] = v
	return e
}

// Style is the "style"" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *ElementBr) Style(v string) *ElementBr {
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
func (e *ElementBr) Tabindex(v string) *ElementBr {
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
func (e *ElementBr) Title(v string) *ElementBr {
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
func (e *ElementBr) Translate(v string) *ElementBr {
	if v == "" {
		return e
	}
	e.attrs["translate"] = v
	return e
}

// &lt;code id=&quot;attributes-3:event-auxclick&quot;&gt;&lt;a data-x-internal=&quot;event-auxclick&quot; href=&quot;https://w3c.github.io/uievents/#event-type-auxclick&quot;&gt;auxclick&lt;/a&gt;&lt;/code&gt;  event handler
func (e *ElementBr) OnAuxclick(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnBeforematch(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnBeforetoggle(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnBlur(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnCancel(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnCanplay(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnCanplaythrough(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnChange(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnClick(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnClose(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnContextlost(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnContextmenu(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnContextrestored(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnCopy(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnCuechange(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnCut(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDblclick(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDrag(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDragend(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDragenter(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDragleave(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDragover(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDragstart(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDrop(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnDurationchange(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnEmptied(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnEnded(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnError(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnFocus(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnFormdata(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnInput(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnInvalid(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnKeydown(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnKeypress(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnKeyup(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnLoad(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnLoadeddata(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnLoadedmetadata(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnLoadstart(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnMousedown(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnMouseenter(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnMouseleave(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnMousemove(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnMouseout(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnMouseover(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnMouseup(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnPaste(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnPause(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnPlay(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnPlaying(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnProgress(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnRatechange(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnReset(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnResize(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnScroll(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnScrollend(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnSecuritypolicyviolation(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnSeeked(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnSeeking(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnSelect(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnSlotchange(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnStalled(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnSubmit(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnSuspend(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnTimeupdate(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnToggle(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnVolumechange(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnWaiting(fn engine.EventHandler) *ElementBr {
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
func (e *ElementBr) OnWheel(fn engine.EventHandler) *ElementBr {
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
