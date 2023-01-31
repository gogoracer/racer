/* cSpell:disable */

package handlebars

import "github.com/gogoracer/racer/pkg/engine"

type InputAttributes struct {

	// Accept is the "accept"" attribute.
	// Hint for expected file type in
	// Valid values are constrained to the following:
	//   * set-of-comma-separated-tokens
	//   * valid-mime-type-with-no-parameters
	//   * audio/*
	//   * video/*
	//   * image/*
	Accept string `json:"accept"`

	// Accesskey is the "accesskey"" attribute.
	// Keyboard shortcut to activate or focus element
	// Valid values are constrained to the following:
	//   * ordered-set-of-unique-space-separated-tokens
	//   * string-is
	Accesskey string `json:"accesskey"`

	// Alt is the "alt"" attribute.
	// Replacement text for use when images are not available
	// Valid values are constrained to the following:
	//   * attribute-text
	Alt string `json:"alt"`

	// Autocapitalize is the "autocapitalize"" attribute.
	// Recommended autocapitalization behavior (for supported input methods)
	// Valid values are constrained to the following:
	//   *
	//   *
	//   *
	//   *
	//   *
	//   *
	Autocapitalize string `json:"autocapitalize"`

	// Autocomplete is the "autocomplete"" attribute.
	// Hint for form autofill feature
	// Valid values are constrained to the following:
	//   * autofill-field
	Autocomplete string `json:"autocomplete"`

	// Autofocus is the "autofocus"" attribute.
	// Automatically focus the element when the page is loaded
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Autofocus string `json:"autofocus"`

	// Checked is the "checked"" attribute.
	// Whether the control is checked
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Checked string `json:"checked"`

	// Class is the "class"" attribute.
	// Classes to which the element belongs
	// Valid values are constrained to the following:
	//   * set-of-space-separated-tokens
	Class string `json:"class"`

	// Contenteditable is the "contenteditable"" attribute.
	// Whether the element is editable
	// Valid values are constrained to the following:
	//   * true
	//   * false
	Contenteditable string `json:"contenteditable"`

	// Dir is the "dir"" attribute.
	//
	// Valid values are constrained to the following:
	//   *
	//   *
	//   *
	Dir string `json:"dir"`

	// Dirname is the "dirname"" attribute.
	// Name of form control to use for sending the element&#39;s
	// Valid values are constrained to the following:
	//   * attribute-text
	Dirname string `json:"dirname"`

	// Disabled is the "disabled"" attribute.
	// Whether the form control is disabled
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Disabled string `json:"disabled"`

	// Draggable is the "draggable"" attribute.
	// Whether the element is draggable
	// Valid values are constrained to the following:
	//   * true
	//   * false
	Draggable string `json:"draggable"`

	// Enterkeyhint is the "enterkeyhint"" attribute.
	// Hint for selecting an enter key action
	// Valid values are constrained to the following:
	//   *
	//   *
	//   *
	//   *
	//   *
	//   *
	//   *
	Enterkeyhint string `json:"enterkeyhint"`

	// Form is the "form"" attribute.
	// Associates the element with a
	// Valid values are constrained to the following:
	//   * concept-id
	Form string `json:"form"`

	// Formaction is the "formaction"" attribute.
	//
	// Valid values are constrained to the following:
	//   * valid-non-empty-url-potentially-surrounded-by-spaces
	Formaction string `json:"formaction"`

	// Formenctype is the "formenctype"" attribute.
	//
	// Valid values are constrained to the following:
	//   *
	//   *
	//   *
	Formenctype string `json:"formenctype"`

	// Formmethod is the "formmethod"" attribute.
	// Variant to use for
	// Valid values are constrained to the following:
	//   * GET
	//   * POST
	//   * dialog
	Formmethod string `json:"formmethod"`

	// Formnovalidate is the "formnovalidate"" attribute.
	// Bypass form control validation for
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Formnovalidate string `json:"formnovalidate"`

	// Formtarget is the "formtarget"" attribute.
	//
	// Valid values are constrained to the following:
	//   * valid-navigable-target-name-or-keyword
	Formtarget string `json:"formtarget"`

	// Height is the "height"" attribute.
	// Vertical dimension
	// Valid values are constrained to the following:
	//   * valid-non-negative-integer
	Height string `json:"height"`

	// Hidden is the "hidden"" attribute.
	// Whether the element is relevant
	// Valid values are constrained to the following:
	//   *
	//   *
	Hidden string `json:"hidden"`

	// Id is the "id"" attribute.
	// The element&#39;s
	// Valid values are constrained to the following:
	//   * attribute-text
	Id string `json:"id"`

	// Inert is the "inert"" attribute.
	// Whether the element is
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Inert string `json:"inert"`

	// Inputmode is the "inputmode"" attribute.
	// Hint for selecting an input modality
	// Valid values are constrained to the following:
	//   *
	//   *
	//   *
	//   *
	//   *
	//   *
	//   *
	//   *
	Inputmode string `json:"inputmode"`

	// Is is the "is"" attribute.
	// Creates a
	// Valid values are constrained to the following:
	//   * valid-custom-element-name
	//   * customized-built-in-element
	Is string `json:"is"`

	// Itemid is the "itemid"" attribute.
	//
	// Valid values are constrained to the following:
	//   * valid-url-potentially-surrounded-by-spaces
	Itemid string `json:"itemid"`

	// Itemprop is the "itemprop"" attribute.
	//
	// Valid values are constrained to the following:
	//   * unordered-set-of-unique-space-separated-tokens
	//   * syntax-url-absolute
	//   * defined-property-name
	Itemprop string `json:"itemprop"`

	// Itemref is the "itemref"" attribute.
	//
	// Valid values are constrained to the following:
	//   * unordered-set-of-unique-space-separated-tokens
	Itemref string `json:"itemref"`

	// Itemscope is the "itemscope"" attribute.
	// Introduces a microdata item
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Itemscope string `json:"itemscope"`

	// Itemtype is the "itemtype"" attribute.
	//
	// Valid values are constrained to the following:
	//   * unordered-set-of-unique-space-separated-tokens
	//   * syntax-url-absolute
	Itemtype string `json:"itemtype"`

	// Lang is the "lang"" attribute.
	//
	// Valid values are constrained to the following:
	Lang string `json:"lang"`

	// List is the "list"" attribute.
	// List of autocomplete options
	// Valid values are constrained to the following:
	//   * concept-id
	List string `json:"list"`

	// Max is the "max"" attribute.
	// Maximum value
	// Valid values are constrained to the following:
	Max string `json:"max"`

	// Maxlength is the "maxlength"" attribute.
	// Maximum
	// Valid values are constrained to the following:
	//   * valid-non-negative-integer
	Maxlength string `json:"maxlength"`

	// Min is the "min"" attribute.
	// Minimum value
	// Valid values are constrained to the following:
	Min string `json:"min"`

	// Minlength is the "minlength"" attribute.
	// Minimum
	// Valid values are constrained to the following:
	//   * valid-non-negative-integer
	Minlength string `json:"minlength"`

	// Multiple is the "multiple"" attribute.
	// Whether to allow multiple values
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Multiple string `json:"multiple"`

	// Name is the "name"" attribute.
	// Name of the element to use for
	// Valid values are constrained to the following:
	//   * attribute-text
	Name string `json:"name"`

	// Nonce is the "nonce"" attribute.
	// Cryptographic nonce used in
	// Valid values are constrained to the following:
	//   * attribute-text
	Nonce string `json:"nonce"`

	// Pattern is the "pattern"" attribute.
	// Pattern to be matched by the form control&#39;s value
	// Valid values are constrained to the following:
	Pattern string `json:"pattern"`

	// Placeholder is the "placeholder"" attribute.
	// User-visible label to be placed within the form control
	// Valid values are constrained to the following:
	//   * attribute-text
	Placeholder string `json:"placeholder"`

	// Popover is the "popover"" attribute.
	// Makes the element a
	// Valid values are constrained to the following:
	//   *
	//   *
	Popover string `json:"popover"`

	// Popoverhidetarget is the "popoverhidetarget"" attribute.
	// Hides the specified
	// Valid values are constrained to the following:
	Popoverhidetarget string `json:"popoverhidetarget"`

	// Popovershowtarget is the "popovershowtarget"" attribute.
	// Shows the specified
	// Valid values are constrained to the following:
	Popovershowtarget string `json:"popovershowtarget"`

	// Popovertoggletarget is the "popovertoggletarget"" attribute.
	// Toggles the specified
	// Valid values are constrained to the following:
	Popovertoggletarget string `json:"popovertoggletarget"`

	// Readonly is the "readonly"" attribute.
	// Whether to allow the value to be edited by the user
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Readonly string `json:"readonly"`

	// Required is the "required"" attribute.
	// Whether the control is required for
	// Valid values are constrained to the following:
	//   * boolean-attribute
	Required string `json:"required"`

	// Size is the "size"" attribute.
	// Size of the control
	// Valid values are constrained to the following:
	//   * valid-non-negative-integer
	Size string `json:"size"`

	// Slot is the "slot"" attribute.
	// The element&#39;s desired slot
	// Valid values are constrained to the following:
	//   * attribute-text
	Slot string `json:"slot"`

	// Spellcheck is the "spellcheck"" attribute.
	// Whether the element is to have its spelling and grammar checked
	// Valid values are constrained to the following:
	//   * true
	//   * false
	Spellcheck string `json:"spellcheck"`

	// Src is the "src"" attribute.
	// Address of the resource
	// Valid values are constrained to the following:
	//   * valid-non-empty-url-potentially-surrounded-by-spaces
	Src string `json:"src"`

	// Step is the "step"" attribute.
	// Granularity to be matched by the form control&#39;s value
	// Valid values are constrained to the following:
	//   * valid-floating-point-number
	//   * any
	Step string `json:"step"`

	// Style is the "style"" attribute.
	// Presentational and formatting instructions
	// Valid values are constrained to the following:
	Style string `json:"style"`

	// Tabindex is the "tabindex"" attribute.
	// Whether the element is
	// Valid values are constrained to the following:
	//   * valid-integer
	Tabindex string `json:"tabindex"`

	// Title is the "title"" attribute.
	// Description of pattern (when used with
	// Valid values are constrained to the following:
	//   * attribute-text
	Title string `json:"title"`

	// Translate is the "translate"" attribute.
	// Whether the element is to be translated when the page is localized
	// Valid values are constrained to the following:
	//   * yes
	//   * no
	Translate string `json:"translate"`

	// Type is the "type"" attribute.
	// Type of form control
	// Valid values are constrained to the following:
	//   * attr-input-type
	Type string `json:"type"`

	// Value is the "value"" attribute.
	// Value of the form control
	// Valid values are constrained to the following:
	Value string `json:"value"`

	// Width is the "width"" attribute.
	// Horizontal dimension
	// Valid values are constrained to the following:
	//   * valid-non-negative-integer
	Width string `json:"width"`

	// <code id="attributes-3:event-auxclick"><a data-x-internal="event-auxclick" href="https://w3c.github.io/uievents/#event-type-auxclick">auxclick</a></code>  event handler
	ClientSideAuxclick string

	// <code id="attributes-3:event-beforematch"><a href="#event-beforematch">beforematch</a></code>  event handler
	ClientSideBeforematch string

	// <code id="attributes-3:event-beforetoggle"><a href="#event-beforetoggle">beforetoggle</a></code>  event handler
	ClientSideBeforetoggle string

	// <code id="attributes-3:event-blur"><a href="#event-blur">blur</a></code>  event handler
	ClientSideBlur string

	// <code id="attributes-3:event-cancel"><a href="#event-cancel">cancel</a></code>  event handler
	ClientSideCancel string

	// <code id="attributes-3:event-media-canplay"><a href="media.html#event-media-canplay">canplay</a></code>  event handler
	ClientSideCanplay string

	// <code id="attributes-3:event-media-canplaythrough"><a href="media.html#event-media-canplaythrough">canplaythrough</a></code>  event handler
	ClientSideCanplaythrough string

	// <code id="attributes-3:event-change"><a href="#event-change">change</a></code>  event handler
	ClientSideChange string

	// <code id="attributes-3:event-click"><a data-x-internal="event-click" href="https://w3c.github.io/uievents/#event-type-click">click</a></code>  event handler
	ClientSideClick string

	// <code id="attributes-3:event-close"><a href="#event-close">close</a></code>  event handler
	ClientSideClose string

	// <code id="attributes-3:event-contextlost"><a href="#event-contextlost">contextlost</a></code>  event handler
	ClientSideContextlost string

	// <code id="attributes-3:event-contextmenu"><a data-x-internal="event-contextmenu" href="https://w3c.github.io/uievents/#event-type-contextmenu">contextmenu</a></code>  event handler
	ClientSideContextmenu string

	// <code id="attributes-3:event-contextrestored"><a href="#event-contextrestored">contextrestored</a></code>  event handler
	ClientSideContextrestored string

	// <code id="attributes-3:event-copy"><a data-x-internal="event-copy" href="https://w3c.github.io/clipboard-apis/#clipboard-event-copy">copy</a></code>  event handler
	ClientSideCopy string

	// <code id="attributes-3:event-media-cuechange"><a href="media.html#event-media-cuechange">cuechange</a></code>  event handler
	ClientSideCuechange string

	// <code id="attributes-3:event-cut"><a data-x-internal="event-cut" href="https://w3c.github.io/clipboard-apis/#clipboard-event-cut">cut</a></code>  event handler
	ClientSideCut string

	// <code id="attributes-3:event-dblclick"><a data-x-internal="event-dblclick" href="https://w3c.github.io/uievents/#event-type-dblclick">dblclick</a></code>  event handler
	ClientSideDblclick string

	// <code id="attributes-3:event-dnd-drag"><a href="dnd.html#event-dnd-drag">drag</a></code>  event handler
	ClientSideDrag string

	// <code id="attributes-3:event-dnd-dragend"><a href="dnd.html#event-dnd-dragend">dragend</a></code>  event handler
	ClientSideDragend string

	// <code id="attributes-3:event-dnd-dragenter"><a href="dnd.html#event-dnd-dragenter">dragenter</a></code>  event handler
	ClientSideDragenter string

	// <code id="attributes-3:event-dnd-dragleave"><a href="dnd.html#event-dnd-dragleave">dragleave</a></code>  event handler
	ClientSideDragleave string

	// <code id="attributes-3:event-dnd-dragover"><a href="dnd.html#event-dnd-dragover">dragover</a></code>  event handler
	ClientSideDragover string

	// <code id="attributes-3:event-dnd-dragstart"><a href="dnd.html#event-dnd-dragstart">dragstart</a></code>  event handler
	ClientSideDragstart string

	// <code id="attributes-3:event-dnd-drop"><a href="dnd.html#event-dnd-drop">drop</a></code>  event handler
	ClientSideDrop string

	// <code id="attributes-3:event-media-durationchange"><a href="media.html#event-media-durationchange">durationchange</a></code>  event handler
	ClientSideDurationchange string

	// <code id="attributes-3:event-media-emptied"><a href="media.html#event-media-emptied">emptied</a></code>  event handler
	ClientSideEmptied string

	// <code id="attributes-3:event-media-ended"><a href="media.html#event-media-ended">ended</a></code>  event handler
	ClientSideEnded string

	// <code id="attributes-3:event-error"><a href="#event-error">error</a></code>  event handler
	ClientSideError string

	// <code id="attributes-3:event-focus"><a href="#event-focus">focus</a></code>  event handler
	ClientSideFocus string

	// <code id="attributes-3:event-formdata"><a href="#event-formdata">formdata</a></code>  event handler
	ClientSideFormdata string

	// <code id="attributes-3:event-input"><a data-x-internal="event-input" href="https://w3c.github.io/uievents/#event-type-input">input</a></code>  event handler
	ClientSideInput string

	// <code id="attributes-3:event-invalid"><a href="#event-invalid">invalid</a></code>  event handler
	ClientSideInvalid string

	// <code id="attributes-3:event-keydown"><a data-x-internal="event-keydown" href="https://w3c.github.io/uievents/#event-type-keydown">keydown</a></code>  event handler
	ClientSideKeydown string

	// <code id="attributes-3:event-keypress"><a data-x-internal="event-keypress" href="https://w3c.github.io/uievents/#event-type-keypress">keypress</a></code>  event handler
	ClientSideKeypress string

	// <code id="attributes-3:event-keyup"><a data-x-internal="event-keyup" href="https://w3c.github.io/uievents/#event-type-keyup">keyup</a></code>  event handler
	ClientSideKeyup string

	// <code id="attributes-3:event-load"><a href="#event-load">load</a></code>  event handler
	ClientSideLoad string

	// <code id="attributes-3:event-media-loadeddata"><a href="media.html#event-media-loadeddata">loadeddata</a></code>  event handler
	ClientSideLoadeddata string

	// <code id="attributes-3:event-media-loadedmetadata"><a href="media.html#event-media-loadedmetadata">loadedmetadata</a></code>  event handler
	ClientSideLoadedmetadata string

	// <code id="attributes-3:event-media-loadstart"><a href="media.html#event-media-loadstart">loadstart</a></code>  event handler
	ClientSideLoadstart string

	// <code id="attributes-3:event-mousedown"><a data-x-internal="event-mousedown" href="https://w3c.github.io/uievents/#event-type-mousedown">mousedown</a></code>  event handler
	ClientSideMousedown string

	// <code id="attributes-3:event-mouseenter"><a data-x-internal="event-mouseenter" href="https://w3c.github.io/uievents/#event-type-mouseenter">mouseenter</a></code>  event handler
	ClientSideMouseenter string

	// <code id="attributes-3:event-mouseleave"><a data-x-internal="event-mouseleave" href="https://w3c.github.io/uievents/#event-type-mouseleave">mouseleave</a></code>  event handler
	ClientSideMouseleave string

	// <code id="attributes-3:event-mousemove"><a data-x-internal="event-mousemove" href="https://w3c.github.io/uievents/#event-type-mousemove">mousemove</a></code>  event handler
	ClientSideMousemove string

	// <code id="attributes-3:event-mouseout"><a data-x-internal="event-mouseout" href="https://w3c.github.io/uievents/#event-type-mouseout">mouseout</a></code>  event handler
	ClientSideMouseout string

	// <code id="attributes-3:event-mouseover"><a data-x-internal="event-mouseover" href="https://w3c.github.io/uievents/#event-type-mouseover">mouseover</a></code>  event handler
	ClientSideMouseover string

	// <code id="attributes-3:event-mouseup"><a data-x-internal="event-mouseup" href="https://w3c.github.io/uievents/#event-type-mouseup">mouseup</a></code>  event handler
	ClientSideMouseup string

	// <code id="attributes-3:event-paste"><a data-x-internal="event-paste" href="https://w3c.github.io/clipboard-apis/#clipboard-event-paste">paste</a></code>  event handler
	ClientSidePaste string

	// <code id="attributes-3:event-media-pause"><a href="media.html#event-media-pause">pause</a></code>  event handler
	ClientSidePause string

	// <code id="attributes-3:event-media-play"><a href="media.html#event-media-play">play</a></code>  event handler
	ClientSidePlay string

	// <code id="attributes-3:event-media-playing"><a href="media.html#event-media-playing">playing</a></code>  event handler
	ClientSidePlaying string

	// <code id="attributes-3:event-media-progress"><a href="media.html#event-media-progress">progress</a></code>  event handler
	ClientSideProgress string

	// <code id="attributes-3:event-media-ratechange"><a href="media.html#event-media-ratechange">ratechange</a></code>  event handler
	ClientSideRatechange string

	// <code id="attributes-3:event-reset"><a href="#event-reset">reset</a></code>  event handler
	ClientSideReset string

	// <code id="attributes-3:event-resize"><a data-x-internal="event-resize" href="https://drafts.csswg.org/cssom-view/#eventdef-window-resize">resize</a></code>  event handler
	ClientSideResize string

	// <code id="attributes-3:event-scroll"><a data-x-internal="event-scroll" href="https://drafts.csswg.org/cssom-view/#eventdef-document-scroll">scroll</a></code>  event handler
	ClientSideScroll string

	// <code id="attributes-3:event-scrollend"><a data-x-internal="event-scrollend" href="https://drafts.csswg.org/cssom-view/#eventdef-document-scrollend">scrollend</a></code>  event handler
	ClientSideScrollend string

	// <code id="attributes-3:event-securitypolicyviolation"><a data-x-internal="event-securitypolicyviolation" href="https://w3c.github.io/webappsec-csp/#eventdef-globaleventhandlers-securitypolicyviolation">securitypolicyviolation</a></code>  event handler
	ClientSideSecuritypolicyviolation string

	// <code id="attributes-3:event-media-seeked"><a href="media.html#event-media-seeked">seeked</a></code>  event handler
	ClientSideSeeked string

	// <code id="attributes-3:event-media-seeking"><a href="media.html#event-media-seeking">seeking</a></code>  event handler
	ClientSideSeeking string

	// <code id="attributes-3:event-select"><a href="#event-select">select</a></code>  event handler
	ClientSideSelect string

	// <code id="attributes-3:event-slotchange"><a data-x-internal="event-slotchange" href="https://dom.spec.whatwg.org/#eventdef-htmlslotelement-slotchange">slotchange</a></code>  event handler
	ClientSideSlotchange string

	// <code id="attributes-3:event-media-stalled"><a href="media.html#event-media-stalled">stalled</a></code>  event handler
	ClientSideStalled string

	// <code id="attributes-3:event-submit"><a href="#event-submit">submit</a></code>  event handler
	ClientSideSubmit string

	// <code id="attributes-3:event-media-suspend"><a href="media.html#event-media-suspend">suspend</a></code>  event handler
	ClientSideSuspend string

	// <code id="attributes-3:event-media-timeupdate"><a href="media.html#event-media-timeupdate">timeupdate</a></code>  event handler
	ClientSideTimeupdate string

	// <code id="attributes-3:event-toggle"><a href="#event-toggle">toggle</a></code>  event handler
	ClientSideToggle string

	// <code id="attributes-3:event-media-volumechange"><a href="media.html#event-media-volumechange">volumechange</a></code>  event handler
	ClientSideVolumechange string

	// <code id="attributes-3:event-media-waiting"><a href="media.html#event-media-waiting">waiting</a></code>  event handler
	ClientSideWaiting string

	// <code id="attributes-3:event-wheel"><a data-x-internal="event-wheel" href="https://w3c.github.io/uievents/#event-type-wheel">wheel</a></code>  event handler
	ClientSideWheel string
}

type InputEventHandlers struct {

	// <code id="attributes-3:event-auxclick"><a data-x-internal="event-auxclick" href="https://w3c.github.io/uievents/#event-type-auxclick">auxclick</a></code>  event handler
	OnAuxclick engine.EventHandler

	// <code id="attributes-3:event-beforematch"><a href="#event-beforematch">beforematch</a></code>  event handler
	OnBeforematch engine.EventHandler

	// <code id="attributes-3:event-beforetoggle"><a href="#event-beforetoggle">beforetoggle</a></code>  event handler
	OnBeforetoggle engine.EventHandler

	// <code id="attributes-3:event-blur"><a href="#event-blur">blur</a></code>  event handler
	OnBlur engine.EventHandler

	// <code id="attributes-3:event-cancel"><a href="#event-cancel">cancel</a></code>  event handler
	OnCancel engine.EventHandler

	// <code id="attributes-3:event-media-canplay"><a href="media.html#event-media-canplay">canplay</a></code>  event handler
	OnCanplay engine.EventHandler

	// <code id="attributes-3:event-media-canplaythrough"><a href="media.html#event-media-canplaythrough">canplaythrough</a></code>  event handler
	OnCanplaythrough engine.EventHandler

	// <code id="attributes-3:event-change"><a href="#event-change">change</a></code>  event handler
	OnChange engine.EventHandler

	// <code id="attributes-3:event-click"><a data-x-internal="event-click" href="https://w3c.github.io/uievents/#event-type-click">click</a></code>  event handler
	OnClick engine.EventHandler

	// <code id="attributes-3:event-close"><a href="#event-close">close</a></code>  event handler
	OnClose engine.EventHandler

	// <code id="attributes-3:event-contextlost"><a href="#event-contextlost">contextlost</a></code>  event handler
	OnContextlost engine.EventHandler

	// <code id="attributes-3:event-contextmenu"><a data-x-internal="event-contextmenu" href="https://w3c.github.io/uievents/#event-type-contextmenu">contextmenu</a></code>  event handler
	OnContextmenu engine.EventHandler

	// <code id="attributes-3:event-contextrestored"><a href="#event-contextrestored">contextrestored</a></code>  event handler
	OnContextrestored engine.EventHandler

	// <code id="attributes-3:event-copy"><a data-x-internal="event-copy" href="https://w3c.github.io/clipboard-apis/#clipboard-event-copy">copy</a></code>  event handler
	OnCopy engine.EventHandler

	// <code id="attributes-3:event-media-cuechange"><a href="media.html#event-media-cuechange">cuechange</a></code>  event handler
	OnCuechange engine.EventHandler

	// <code id="attributes-3:event-cut"><a data-x-internal="event-cut" href="https://w3c.github.io/clipboard-apis/#clipboard-event-cut">cut</a></code>  event handler
	OnCut engine.EventHandler

	// <code id="attributes-3:event-dblclick"><a data-x-internal="event-dblclick" href="https://w3c.github.io/uievents/#event-type-dblclick">dblclick</a></code>  event handler
	OnDblclick engine.EventHandler

	// <code id="attributes-3:event-dnd-drag"><a href="dnd.html#event-dnd-drag">drag</a></code>  event handler
	OnDrag engine.EventHandler

	// <code id="attributes-3:event-dnd-dragend"><a href="dnd.html#event-dnd-dragend">dragend</a></code>  event handler
	OnDragend engine.EventHandler

	// <code id="attributes-3:event-dnd-dragenter"><a href="dnd.html#event-dnd-dragenter">dragenter</a></code>  event handler
	OnDragenter engine.EventHandler

	// <code id="attributes-3:event-dnd-dragleave"><a href="dnd.html#event-dnd-dragleave">dragleave</a></code>  event handler
	OnDragleave engine.EventHandler

	// <code id="attributes-3:event-dnd-dragover"><a href="dnd.html#event-dnd-dragover">dragover</a></code>  event handler
	OnDragover engine.EventHandler

	// <code id="attributes-3:event-dnd-dragstart"><a href="dnd.html#event-dnd-dragstart">dragstart</a></code>  event handler
	OnDragstart engine.EventHandler

	// <code id="attributes-3:event-dnd-drop"><a href="dnd.html#event-dnd-drop">drop</a></code>  event handler
	OnDrop engine.EventHandler

	// <code id="attributes-3:event-media-durationchange"><a href="media.html#event-media-durationchange">durationchange</a></code>  event handler
	OnDurationchange engine.EventHandler

	// <code id="attributes-3:event-media-emptied"><a href="media.html#event-media-emptied">emptied</a></code>  event handler
	OnEmptied engine.EventHandler

	// <code id="attributes-3:event-media-ended"><a href="media.html#event-media-ended">ended</a></code>  event handler
	OnEnded engine.EventHandler

	// <code id="attributes-3:event-error"><a href="#event-error">error</a></code>  event handler
	OnError engine.EventHandler

	// <code id="attributes-3:event-focus"><a href="#event-focus">focus</a></code>  event handler
	OnFocus engine.EventHandler

	// <code id="attributes-3:event-formdata"><a href="#event-formdata">formdata</a></code>  event handler
	OnFormdata engine.EventHandler

	// <code id="attributes-3:event-input"><a data-x-internal="event-input" href="https://w3c.github.io/uievents/#event-type-input">input</a></code>  event handler
	OnInput engine.EventHandler

	// <code id="attributes-3:event-invalid"><a href="#event-invalid">invalid</a></code>  event handler
	OnInvalid engine.EventHandler

	// <code id="attributes-3:event-keydown"><a data-x-internal="event-keydown" href="https://w3c.github.io/uievents/#event-type-keydown">keydown</a></code>  event handler
	OnKeydown engine.EventHandler

	// <code id="attributes-3:event-keypress"><a data-x-internal="event-keypress" href="https://w3c.github.io/uievents/#event-type-keypress">keypress</a></code>  event handler
	OnKeypress engine.EventHandler

	// <code id="attributes-3:event-keyup"><a data-x-internal="event-keyup" href="https://w3c.github.io/uievents/#event-type-keyup">keyup</a></code>  event handler
	OnKeyup engine.EventHandler

	// <code id="attributes-3:event-load"><a href="#event-load">load</a></code>  event handler
	OnLoad engine.EventHandler

	// <code id="attributes-3:event-media-loadeddata"><a href="media.html#event-media-loadeddata">loadeddata</a></code>  event handler
	OnLoadeddata engine.EventHandler

	// <code id="attributes-3:event-media-loadedmetadata"><a href="media.html#event-media-loadedmetadata">loadedmetadata</a></code>  event handler
	OnLoadedmetadata engine.EventHandler

	// <code id="attributes-3:event-media-loadstart"><a href="media.html#event-media-loadstart">loadstart</a></code>  event handler
	OnLoadstart engine.EventHandler

	// <code id="attributes-3:event-mousedown"><a data-x-internal="event-mousedown" href="https://w3c.github.io/uievents/#event-type-mousedown">mousedown</a></code>  event handler
	OnMousedown engine.EventHandler

	// <code id="attributes-3:event-mouseenter"><a data-x-internal="event-mouseenter" href="https://w3c.github.io/uievents/#event-type-mouseenter">mouseenter</a></code>  event handler
	OnMouseenter engine.EventHandler

	// <code id="attributes-3:event-mouseleave"><a data-x-internal="event-mouseleave" href="https://w3c.github.io/uievents/#event-type-mouseleave">mouseleave</a></code>  event handler
	OnMouseleave engine.EventHandler

	// <code id="attributes-3:event-mousemove"><a data-x-internal="event-mousemove" href="https://w3c.github.io/uievents/#event-type-mousemove">mousemove</a></code>  event handler
	OnMousemove engine.EventHandler

	// <code id="attributes-3:event-mouseout"><a data-x-internal="event-mouseout" href="https://w3c.github.io/uievents/#event-type-mouseout">mouseout</a></code>  event handler
	OnMouseout engine.EventHandler

	// <code id="attributes-3:event-mouseover"><a data-x-internal="event-mouseover" href="https://w3c.github.io/uievents/#event-type-mouseover">mouseover</a></code>  event handler
	OnMouseover engine.EventHandler

	// <code id="attributes-3:event-mouseup"><a data-x-internal="event-mouseup" href="https://w3c.github.io/uievents/#event-type-mouseup">mouseup</a></code>  event handler
	OnMouseup engine.EventHandler

	// <code id="attributes-3:event-paste"><a data-x-internal="event-paste" href="https://w3c.github.io/clipboard-apis/#clipboard-event-paste">paste</a></code>  event handler
	OnPaste engine.EventHandler

	// <code id="attributes-3:event-media-pause"><a href="media.html#event-media-pause">pause</a></code>  event handler
	OnPause engine.EventHandler

	// <code id="attributes-3:event-media-play"><a href="media.html#event-media-play">play</a></code>  event handler
	OnPlay engine.EventHandler

	// <code id="attributes-3:event-media-playing"><a href="media.html#event-media-playing">playing</a></code>  event handler
	OnPlaying engine.EventHandler

	// <code id="attributes-3:event-media-progress"><a href="media.html#event-media-progress">progress</a></code>  event handler
	OnProgress engine.EventHandler

	// <code id="attributes-3:event-media-ratechange"><a href="media.html#event-media-ratechange">ratechange</a></code>  event handler
	OnRatechange engine.EventHandler

	// <code id="attributes-3:event-reset"><a href="#event-reset">reset</a></code>  event handler
	OnReset engine.EventHandler

	// <code id="attributes-3:event-resize"><a data-x-internal="event-resize" href="https://drafts.csswg.org/cssom-view/#eventdef-window-resize">resize</a></code>  event handler
	OnResize engine.EventHandler

	// <code id="attributes-3:event-scroll"><a data-x-internal="event-scroll" href="https://drafts.csswg.org/cssom-view/#eventdef-document-scroll">scroll</a></code>  event handler
	OnScroll engine.EventHandler

	// <code id="attributes-3:event-scrollend"><a data-x-internal="event-scrollend" href="https://drafts.csswg.org/cssom-view/#eventdef-document-scrollend">scrollend</a></code>  event handler
	OnScrollend engine.EventHandler

	// <code id="attributes-3:event-securitypolicyviolation"><a data-x-internal="event-securitypolicyviolation" href="https://w3c.github.io/webappsec-csp/#eventdef-globaleventhandlers-securitypolicyviolation">securitypolicyviolation</a></code>  event handler
	OnSecuritypolicyviolation engine.EventHandler

	// <code id="attributes-3:event-media-seeked"><a href="media.html#event-media-seeked">seeked</a></code>  event handler
	OnSeeked engine.EventHandler

	// <code id="attributes-3:event-media-seeking"><a href="media.html#event-media-seeking">seeking</a></code>  event handler
	OnSeeking engine.EventHandler

	// <code id="attributes-3:event-select"><a href="#event-select">select</a></code>  event handler
	OnSelect engine.EventHandler

	// <code id="attributes-3:event-slotchange"><a data-x-internal="event-slotchange" href="https://dom.spec.whatwg.org/#eventdef-htmlslotelement-slotchange">slotchange</a></code>  event handler
	OnSlotchange engine.EventHandler

	// <code id="attributes-3:event-media-stalled"><a href="media.html#event-media-stalled">stalled</a></code>  event handler
	OnStalled engine.EventHandler

	// <code id="attributes-3:event-submit"><a href="#event-submit">submit</a></code>  event handler
	OnSubmit engine.EventHandler

	// <code id="attributes-3:event-media-suspend"><a href="media.html#event-media-suspend">suspend</a></code>  event handler
	OnSuspend engine.EventHandler

	// <code id="attributes-3:event-media-timeupdate"><a href="media.html#event-media-timeupdate">timeupdate</a></code>  event handler
	OnTimeupdate engine.EventHandler

	// <code id="attributes-3:event-toggle"><a href="#event-toggle">toggle</a></code>  event handler
	OnToggle engine.EventHandler

	// <code id="attributes-3:event-media-volumechange"><a href="media.html#event-media-volumechange">volumechange</a></code>  event handler
	OnVolumechange engine.EventHandler

	// <code id="attributes-3:event-media-waiting"><a href="media.html#event-media-waiting">waiting</a></code>  event handler
	OnWaiting engine.EventHandler

	// <code id="attributes-3:event-wheel"><a data-x-internal="event-wheel" href="https://w3c.github.io/uievents/#event-type-wheel">wheel</a></code>  event handler
	OnWheel engine.EventHandler
}

type ElementInput struct {
	Attributes    InputAttributes
	EventHandlers InputEventHandlers

	// Children is the list of child nodes.
	Children []any
}

func INPUT(children ...any) *ElementInput {
	return &ElementInput{
		Children: children,
	}
}

func INPUT_A(attributes InputAttributes, children ...any) *ElementInput {
	return &ElementInput{
		Attributes: attributes,
		Children:   children,
	}
}

func INPUT_E(eventHandlers InputEventHandlers, children ...any) *ElementInput {
	return &ElementInput{
		EventHandlers: eventHandlers,
		Children:      children,
	}
}

func INPUT_AE(attributes InputAttributes, eventHandlers InputEventHandlers, children ...any) *ElementInput {
	return &ElementInput{
		Attributes:    attributes,
		EventHandlers: eventHandlers,
		Children:      children,
	}
}

func (element ElementInput) HandlebarElement() {}

func (element ElementInput) GenerateVDOM() interface{} {
	eventHandlers := map[string]engine.EventHandler{}
	if element.EventHandlers.OnAuxclick != nil {
		eventHandlers["auxclick"] = element.EventHandlers.OnAuxclick
	}
	if element.EventHandlers.OnBeforematch != nil {
		eventHandlers["beforematch"] = element.EventHandlers.OnBeforematch
	}
	if element.EventHandlers.OnBeforetoggle != nil {
		eventHandlers["beforetoggle"] = element.EventHandlers.OnBeforetoggle
	}
	if element.EventHandlers.OnBlur != nil {
		eventHandlers["blur"] = element.EventHandlers.OnBlur
	}
	if element.EventHandlers.OnCancel != nil {
		eventHandlers["cancel"] = element.EventHandlers.OnCancel
	}
	if element.EventHandlers.OnCanplay != nil {
		eventHandlers["canplay"] = element.EventHandlers.OnCanplay
	}
	if element.EventHandlers.OnCanplaythrough != nil {
		eventHandlers["canplaythrough"] = element.EventHandlers.OnCanplaythrough
	}
	if element.EventHandlers.OnChange != nil {
		eventHandlers["change"] = element.EventHandlers.OnChange
	}
	if element.EventHandlers.OnClick != nil {
		eventHandlers["click"] = element.EventHandlers.OnClick
	}
	if element.EventHandlers.OnClose != nil {
		eventHandlers["close"] = element.EventHandlers.OnClose
	}
	if element.EventHandlers.OnContextlost != nil {
		eventHandlers["contextlost"] = element.EventHandlers.OnContextlost
	}
	if element.EventHandlers.OnContextmenu != nil {
		eventHandlers["contextmenu"] = element.EventHandlers.OnContextmenu
	}
	if element.EventHandlers.OnContextrestored != nil {
		eventHandlers["contextrestored"] = element.EventHandlers.OnContextrestored
	}
	if element.EventHandlers.OnCopy != nil {
		eventHandlers["copy"] = element.EventHandlers.OnCopy
	}
	if element.EventHandlers.OnCuechange != nil {
		eventHandlers["cuechange"] = element.EventHandlers.OnCuechange
	}
	if element.EventHandlers.OnCut != nil {
		eventHandlers["cut"] = element.EventHandlers.OnCut
	}
	if element.EventHandlers.OnDblclick != nil {
		eventHandlers["dblclick"] = element.EventHandlers.OnDblclick
	}
	if element.EventHandlers.OnDrag != nil {
		eventHandlers["drag"] = element.EventHandlers.OnDrag
	}
	if element.EventHandlers.OnDragend != nil {
		eventHandlers["dragend"] = element.EventHandlers.OnDragend
	}
	if element.EventHandlers.OnDragenter != nil {
		eventHandlers["dragenter"] = element.EventHandlers.OnDragenter
	}
	if element.EventHandlers.OnDragleave != nil {
		eventHandlers["dragleave"] = element.EventHandlers.OnDragleave
	}
	if element.EventHandlers.OnDragover != nil {
		eventHandlers["dragover"] = element.EventHandlers.OnDragover
	}
	if element.EventHandlers.OnDragstart != nil {
		eventHandlers["dragstart"] = element.EventHandlers.OnDragstart
	}
	if element.EventHandlers.OnDrop != nil {
		eventHandlers["drop"] = element.EventHandlers.OnDrop
	}
	if element.EventHandlers.OnDurationchange != nil {
		eventHandlers["durationchange"] = element.EventHandlers.OnDurationchange
	}
	if element.EventHandlers.OnEmptied != nil {
		eventHandlers["emptied"] = element.EventHandlers.OnEmptied
	}
	if element.EventHandlers.OnEnded != nil {
		eventHandlers["ended"] = element.EventHandlers.OnEnded
	}
	if element.EventHandlers.OnError != nil {
		eventHandlers["error"] = element.EventHandlers.OnError
	}
	if element.EventHandlers.OnFocus != nil {
		eventHandlers["focus"] = element.EventHandlers.OnFocus
	}
	if element.EventHandlers.OnFormdata != nil {
		eventHandlers["formdata"] = element.EventHandlers.OnFormdata
	}
	if element.EventHandlers.OnInput != nil {
		eventHandlers["input"] = element.EventHandlers.OnInput
	}
	if element.EventHandlers.OnInvalid != nil {
		eventHandlers["invalid"] = element.EventHandlers.OnInvalid
	}
	if element.EventHandlers.OnKeydown != nil {
		eventHandlers["keydown"] = element.EventHandlers.OnKeydown
	}
	if element.EventHandlers.OnKeypress != nil {
		eventHandlers["keypress"] = element.EventHandlers.OnKeypress
	}
	if element.EventHandlers.OnKeyup != nil {
		eventHandlers["keyup"] = element.EventHandlers.OnKeyup
	}
	if element.EventHandlers.OnLoad != nil {
		eventHandlers["load"] = element.EventHandlers.OnLoad
	}
	if element.EventHandlers.OnLoadeddata != nil {
		eventHandlers["loadeddata"] = element.EventHandlers.OnLoadeddata
	}
	if element.EventHandlers.OnLoadedmetadata != nil {
		eventHandlers["loadedmetadata"] = element.EventHandlers.OnLoadedmetadata
	}
	if element.EventHandlers.OnLoadstart != nil {
		eventHandlers["loadstart"] = element.EventHandlers.OnLoadstart
	}
	if element.EventHandlers.OnMousedown != nil {
		eventHandlers["mousedown"] = element.EventHandlers.OnMousedown
	}
	if element.EventHandlers.OnMouseenter != nil {
		eventHandlers["mouseenter"] = element.EventHandlers.OnMouseenter
	}
	if element.EventHandlers.OnMouseleave != nil {
		eventHandlers["mouseleave"] = element.EventHandlers.OnMouseleave
	}
	if element.EventHandlers.OnMousemove != nil {
		eventHandlers["mousemove"] = element.EventHandlers.OnMousemove
	}
	if element.EventHandlers.OnMouseout != nil {
		eventHandlers["mouseout"] = element.EventHandlers.OnMouseout
	}
	if element.EventHandlers.OnMouseover != nil {
		eventHandlers["mouseover"] = element.EventHandlers.OnMouseover
	}
	if element.EventHandlers.OnMouseup != nil {
		eventHandlers["mouseup"] = element.EventHandlers.OnMouseup
	}
	if element.EventHandlers.OnPaste != nil {
		eventHandlers["paste"] = element.EventHandlers.OnPaste
	}
	if element.EventHandlers.OnPause != nil {
		eventHandlers["pause"] = element.EventHandlers.OnPause
	}
	if element.EventHandlers.OnPlay != nil {
		eventHandlers["play"] = element.EventHandlers.OnPlay
	}
	if element.EventHandlers.OnPlaying != nil {
		eventHandlers["playing"] = element.EventHandlers.OnPlaying
	}
	if element.EventHandlers.OnProgress != nil {
		eventHandlers["progress"] = element.EventHandlers.OnProgress
	}
	if element.EventHandlers.OnRatechange != nil {
		eventHandlers["ratechange"] = element.EventHandlers.OnRatechange
	}
	if element.EventHandlers.OnReset != nil {
		eventHandlers["reset"] = element.EventHandlers.OnReset
	}
	if element.EventHandlers.OnResize != nil {
		eventHandlers["resize"] = element.EventHandlers.OnResize
	}
	if element.EventHandlers.OnScroll != nil {
		eventHandlers["scroll"] = element.EventHandlers.OnScroll
	}
	if element.EventHandlers.OnScrollend != nil {
		eventHandlers["scrollend"] = element.EventHandlers.OnScrollend
	}
	if element.EventHandlers.OnSecuritypolicyviolation != nil {
		eventHandlers["securitypolicyviolation"] = element.EventHandlers.OnSecuritypolicyviolation
	}
	if element.EventHandlers.OnSeeked != nil {
		eventHandlers["seeked"] = element.EventHandlers.OnSeeked
	}
	if element.EventHandlers.OnSeeking != nil {
		eventHandlers["seeking"] = element.EventHandlers.OnSeeking
	}
	if element.EventHandlers.OnSelect != nil {
		eventHandlers["select"] = element.EventHandlers.OnSelect
	}
	if element.EventHandlers.OnSlotchange != nil {
		eventHandlers["slotchange"] = element.EventHandlers.OnSlotchange
	}
	if element.EventHandlers.OnStalled != nil {
		eventHandlers["stalled"] = element.EventHandlers.OnStalled
	}
	if element.EventHandlers.OnSubmit != nil {
		eventHandlers["submit"] = element.EventHandlers.OnSubmit
	}
	if element.EventHandlers.OnSuspend != nil {
		eventHandlers["suspend"] = element.EventHandlers.OnSuspend
	}
	if element.EventHandlers.OnTimeupdate != nil {
		eventHandlers["timeupdate"] = element.EventHandlers.OnTimeupdate
	}
	if element.EventHandlers.OnToggle != nil {
		eventHandlers["toggle"] = element.EventHandlers.OnToggle
	}
	if element.EventHandlers.OnVolumechange != nil {
		eventHandlers["volumechange"] = element.EventHandlers.OnVolumechange
	}
	if element.EventHandlers.OnWaiting != nil {
		eventHandlers["waiting"] = element.EventHandlers.OnWaiting
	}
	if element.EventHandlers.OnWheel != nil {
		eventHandlers["wheel"] = element.EventHandlers.OnWheel
	}

	attrs := engine.Attrs{}
	if element.Attributes.Accept != "" {
		attrs["accept"] = string(element.Attributes.Accept)
	}
	if element.Attributes.Accesskey != "" {
		attrs["accesskey"] = string(element.Attributes.Accesskey)
	}
	if element.Attributes.Alt != "" {
		attrs["alt"] = string(element.Attributes.Alt)
	}
	if element.Attributes.Autocapitalize != "" {
		attrs["autocapitalize"] = string(element.Attributes.Autocapitalize)
	}
	if element.Attributes.Autocomplete != "" {
		attrs["autocomplete"] = string(element.Attributes.Autocomplete)
	}
	if element.Attributes.Autofocus != "" {
		attrs["autofocus"] = string(element.Attributes.Autofocus)
	}
	if element.Attributes.Checked != "" {
		attrs["checked"] = string(element.Attributes.Checked)
	}
	if element.Attributes.Class != "" {
		attrs["class"] = string(element.Attributes.Class)
	}
	if element.Attributes.Contenteditable != "" {
		attrs["contenteditable"] = string(element.Attributes.Contenteditable)
	}
	if element.Attributes.Dir != "" {
		attrs["dir"] = string(element.Attributes.Dir)
	}
	if element.Attributes.Dirname != "" {
		attrs["dirname"] = string(element.Attributes.Dirname)
	}
	if element.Attributes.Disabled != "" {
		attrs["disabled"] = string(element.Attributes.Disabled)
	}
	if element.Attributes.Draggable != "" {
		attrs["draggable"] = string(element.Attributes.Draggable)
	}
	if element.Attributes.Enterkeyhint != "" {
		attrs["enterkeyhint"] = string(element.Attributes.Enterkeyhint)
	}
	if element.Attributes.Form != "" {
		attrs["form"] = string(element.Attributes.Form)
	}
	if element.Attributes.Formaction != "" {
		attrs["formaction"] = string(element.Attributes.Formaction)
	}
	if element.Attributes.Formenctype != "" {
		attrs["formenctype"] = string(element.Attributes.Formenctype)
	}
	if element.Attributes.Formmethod != "" {
		attrs["formmethod"] = string(element.Attributes.Formmethod)
	}
	if element.Attributes.Formnovalidate != "" {
		attrs["formnovalidate"] = string(element.Attributes.Formnovalidate)
	}
	if element.Attributes.Formtarget != "" {
		attrs["formtarget"] = string(element.Attributes.Formtarget)
	}
	if element.Attributes.Height != "" {
		attrs["height"] = string(element.Attributes.Height)
	}
	if element.Attributes.Hidden != "" {
		attrs["hidden"] = string(element.Attributes.Hidden)
	}
	if element.Attributes.Id != "" {
		attrs["id"] = string(element.Attributes.Id)
	}
	if element.Attributes.Inert != "" {
		attrs["inert"] = string(element.Attributes.Inert)
	}
	if element.Attributes.Inputmode != "" {
		attrs["inputmode"] = string(element.Attributes.Inputmode)
	}
	if element.Attributes.Is != "" {
		attrs["is"] = string(element.Attributes.Is)
	}
	if element.Attributes.Itemid != "" {
		attrs["itemid"] = string(element.Attributes.Itemid)
	}
	if element.Attributes.Itemprop != "" {
		attrs["itemprop"] = string(element.Attributes.Itemprop)
	}
	if element.Attributes.Itemref != "" {
		attrs["itemref"] = string(element.Attributes.Itemref)
	}
	if element.Attributes.Itemscope != "" {
		attrs["itemscope"] = string(element.Attributes.Itemscope)
	}
	if element.Attributes.Itemtype != "" {
		attrs["itemtype"] = string(element.Attributes.Itemtype)
	}
	if element.Attributes.Lang != "" {
		attrs["lang"] = string(element.Attributes.Lang)
	}
	if element.Attributes.List != "" {
		attrs["list"] = string(element.Attributes.List)
	}
	if element.Attributes.Max != "" {
		attrs["max"] = string(element.Attributes.Max)
	}
	if element.Attributes.Maxlength != "" {
		attrs["maxlength"] = string(element.Attributes.Maxlength)
	}
	if element.Attributes.Min != "" {
		attrs["min"] = string(element.Attributes.Min)
	}
	if element.Attributes.Minlength != "" {
		attrs["minlength"] = string(element.Attributes.Minlength)
	}
	if element.Attributes.Multiple != "" {
		attrs["multiple"] = string(element.Attributes.Multiple)
	}
	if element.Attributes.Name != "" {
		attrs["name"] = string(element.Attributes.Name)
	}
	if element.Attributes.Nonce != "" {
		attrs["nonce"] = string(element.Attributes.Nonce)
	}
	if element.Attributes.Pattern != "" {
		attrs["pattern"] = string(element.Attributes.Pattern)
	}
	if element.Attributes.Placeholder != "" {
		attrs["placeholder"] = string(element.Attributes.Placeholder)
	}
	if element.Attributes.Popover != "" {
		attrs["popover"] = string(element.Attributes.Popover)
	}
	if element.Attributes.Popoverhidetarget != "" {
		attrs["popoverhidetarget"] = string(element.Attributes.Popoverhidetarget)
	}
	if element.Attributes.Popovershowtarget != "" {
		attrs["popovershowtarget"] = string(element.Attributes.Popovershowtarget)
	}
	if element.Attributes.Popovertoggletarget != "" {
		attrs["popovertoggletarget"] = string(element.Attributes.Popovertoggletarget)
	}
	if element.Attributes.Readonly != "" {
		attrs["readonly"] = string(element.Attributes.Readonly)
	}
	if element.Attributes.Required != "" {
		attrs["required"] = string(element.Attributes.Required)
	}
	if element.Attributes.Size != "" {
		attrs["size"] = string(element.Attributes.Size)
	}
	if element.Attributes.Slot != "" {
		attrs["slot"] = string(element.Attributes.Slot)
	}
	if element.Attributes.Spellcheck != "" {
		attrs["spellcheck"] = string(element.Attributes.Spellcheck)
	}
	if element.Attributes.Src != "" {
		attrs["src"] = string(element.Attributes.Src)
	}
	if element.Attributes.Step != "" {
		attrs["step"] = string(element.Attributes.Step)
	}
	if element.Attributes.Style != "" {
		attrs["style"] = string(element.Attributes.Style)
	}
	if element.Attributes.Tabindex != "" {
		attrs["tabindex"] = string(element.Attributes.Tabindex)
	}
	if element.Attributes.Title != "" {
		attrs["title"] = string(element.Attributes.Title)
	}
	if element.Attributes.Translate != "" {
		attrs["translate"] = string(element.Attributes.Translate)
	}
	if element.Attributes.Type != "" {
		attrs["type"] = string(element.Attributes.Type)
	}
	if element.Attributes.Value != "" {
		attrs["value"] = string(element.Attributes.Value)
	}
	if element.Attributes.Width != "" {
		attrs["width"] = string(element.Attributes.Width)
	}

	if element.Attributes.ClientSideAuxclick != "" {
		attrs["auxclick"] = string(element.Attributes.ClientSideAuxclick)
	}

	if element.Attributes.ClientSideBeforematch != "" {
		attrs["beforematch"] = string(element.Attributes.ClientSideBeforematch)
	}

	if element.Attributes.ClientSideBeforetoggle != "" {
		attrs["beforetoggle"] = string(element.Attributes.ClientSideBeforetoggle)
	}

	if element.Attributes.ClientSideBlur != "" {
		attrs["blur"] = string(element.Attributes.ClientSideBlur)
	}

	if element.Attributes.ClientSideCancel != "" {
		attrs["cancel"] = string(element.Attributes.ClientSideCancel)
	}

	if element.Attributes.ClientSideCanplay != "" {
		attrs["canplay"] = string(element.Attributes.ClientSideCanplay)
	}

	if element.Attributes.ClientSideCanplaythrough != "" {
		attrs["canplaythrough"] = string(element.Attributes.ClientSideCanplaythrough)
	}

	if element.Attributes.ClientSideChange != "" {
		attrs["change"] = string(element.Attributes.ClientSideChange)
	}

	if element.Attributes.ClientSideClick != "" {
		attrs["click"] = string(element.Attributes.ClientSideClick)
	}

	if element.Attributes.ClientSideClose != "" {
		attrs["close"] = string(element.Attributes.ClientSideClose)
	}

	if element.Attributes.ClientSideContextlost != "" {
		attrs["contextlost"] = string(element.Attributes.ClientSideContextlost)
	}

	if element.Attributes.ClientSideContextmenu != "" {
		attrs["contextmenu"] = string(element.Attributes.ClientSideContextmenu)
	}

	if element.Attributes.ClientSideContextrestored != "" {
		attrs["contextrestored"] = string(element.Attributes.ClientSideContextrestored)
	}

	if element.Attributes.ClientSideCopy != "" {
		attrs["copy"] = string(element.Attributes.ClientSideCopy)
	}

	if element.Attributes.ClientSideCuechange != "" {
		attrs["cuechange"] = string(element.Attributes.ClientSideCuechange)
	}

	if element.Attributes.ClientSideCut != "" {
		attrs["cut"] = string(element.Attributes.ClientSideCut)
	}

	if element.Attributes.ClientSideDblclick != "" {
		attrs["dblclick"] = string(element.Attributes.ClientSideDblclick)
	}

	if element.Attributes.ClientSideDrag != "" {
		attrs["drag"] = string(element.Attributes.ClientSideDrag)
	}

	if element.Attributes.ClientSideDragend != "" {
		attrs["dragend"] = string(element.Attributes.ClientSideDragend)
	}

	if element.Attributes.ClientSideDragenter != "" {
		attrs["dragenter"] = string(element.Attributes.ClientSideDragenter)
	}

	if element.Attributes.ClientSideDragleave != "" {
		attrs["dragleave"] = string(element.Attributes.ClientSideDragleave)
	}

	if element.Attributes.ClientSideDragover != "" {
		attrs["dragover"] = string(element.Attributes.ClientSideDragover)
	}

	if element.Attributes.ClientSideDragstart != "" {
		attrs["dragstart"] = string(element.Attributes.ClientSideDragstart)
	}

	if element.Attributes.ClientSideDrop != "" {
		attrs["drop"] = string(element.Attributes.ClientSideDrop)
	}

	if element.Attributes.ClientSideDurationchange != "" {
		attrs["durationchange"] = string(element.Attributes.ClientSideDurationchange)
	}

	if element.Attributes.ClientSideEmptied != "" {
		attrs["emptied"] = string(element.Attributes.ClientSideEmptied)
	}

	if element.Attributes.ClientSideEnded != "" {
		attrs["ended"] = string(element.Attributes.ClientSideEnded)
	}

	if element.Attributes.ClientSideError != "" {
		attrs["error"] = string(element.Attributes.ClientSideError)
	}

	if element.Attributes.ClientSideFocus != "" {
		attrs["focus"] = string(element.Attributes.ClientSideFocus)
	}

	if element.Attributes.ClientSideFormdata != "" {
		attrs["formdata"] = string(element.Attributes.ClientSideFormdata)
	}

	if element.Attributes.ClientSideInput != "" {
		attrs["input"] = string(element.Attributes.ClientSideInput)
	}

	if element.Attributes.ClientSideInvalid != "" {
		attrs["invalid"] = string(element.Attributes.ClientSideInvalid)
	}

	if element.Attributes.ClientSideKeydown != "" {
		attrs["keydown"] = string(element.Attributes.ClientSideKeydown)
	}

	if element.Attributes.ClientSideKeypress != "" {
		attrs["keypress"] = string(element.Attributes.ClientSideKeypress)
	}

	if element.Attributes.ClientSideKeyup != "" {
		attrs["keyup"] = string(element.Attributes.ClientSideKeyup)
	}

	if element.Attributes.ClientSideLoad != "" {
		attrs["load"] = string(element.Attributes.ClientSideLoad)
	}

	if element.Attributes.ClientSideLoadeddata != "" {
		attrs["loadeddata"] = string(element.Attributes.ClientSideLoadeddata)
	}

	if element.Attributes.ClientSideLoadedmetadata != "" {
		attrs["loadedmetadata"] = string(element.Attributes.ClientSideLoadedmetadata)
	}

	if element.Attributes.ClientSideLoadstart != "" {
		attrs["loadstart"] = string(element.Attributes.ClientSideLoadstart)
	}

	if element.Attributes.ClientSideMousedown != "" {
		attrs["mousedown"] = string(element.Attributes.ClientSideMousedown)
	}

	if element.Attributes.ClientSideMouseenter != "" {
		attrs["mouseenter"] = string(element.Attributes.ClientSideMouseenter)
	}

	if element.Attributes.ClientSideMouseleave != "" {
		attrs["mouseleave"] = string(element.Attributes.ClientSideMouseleave)
	}

	if element.Attributes.ClientSideMousemove != "" {
		attrs["mousemove"] = string(element.Attributes.ClientSideMousemove)
	}

	if element.Attributes.ClientSideMouseout != "" {
		attrs["mouseout"] = string(element.Attributes.ClientSideMouseout)
	}

	if element.Attributes.ClientSideMouseover != "" {
		attrs["mouseover"] = string(element.Attributes.ClientSideMouseover)
	}

	if element.Attributes.ClientSideMouseup != "" {
		attrs["mouseup"] = string(element.Attributes.ClientSideMouseup)
	}

	if element.Attributes.ClientSidePaste != "" {
		attrs["paste"] = string(element.Attributes.ClientSidePaste)
	}

	if element.Attributes.ClientSidePause != "" {
		attrs["pause"] = string(element.Attributes.ClientSidePause)
	}

	if element.Attributes.ClientSidePlay != "" {
		attrs["play"] = string(element.Attributes.ClientSidePlay)
	}

	if element.Attributes.ClientSidePlaying != "" {
		attrs["playing"] = string(element.Attributes.ClientSidePlaying)
	}

	if element.Attributes.ClientSideProgress != "" {
		attrs["progress"] = string(element.Attributes.ClientSideProgress)
	}

	if element.Attributes.ClientSideRatechange != "" {
		attrs["ratechange"] = string(element.Attributes.ClientSideRatechange)
	}

	if element.Attributes.ClientSideReset != "" {
		attrs["reset"] = string(element.Attributes.ClientSideReset)
	}

	if element.Attributes.ClientSideResize != "" {
		attrs["resize"] = string(element.Attributes.ClientSideResize)
	}

	if element.Attributes.ClientSideScroll != "" {
		attrs["scroll"] = string(element.Attributes.ClientSideScroll)
	}

	if element.Attributes.ClientSideScrollend != "" {
		attrs["scrollend"] = string(element.Attributes.ClientSideScrollend)
	}

	if element.Attributes.ClientSideSecuritypolicyviolation != "" {
		attrs["securitypolicyviolation"] = string(element.Attributes.ClientSideSecuritypolicyviolation)
	}

	if element.Attributes.ClientSideSeeked != "" {
		attrs["seeked"] = string(element.Attributes.ClientSideSeeked)
	}

	if element.Attributes.ClientSideSeeking != "" {
		attrs["seeking"] = string(element.Attributes.ClientSideSeeking)
	}

	if element.Attributes.ClientSideSelect != "" {
		attrs["select"] = string(element.Attributes.ClientSideSelect)
	}

	if element.Attributes.ClientSideSlotchange != "" {
		attrs["slotchange"] = string(element.Attributes.ClientSideSlotchange)
	}

	if element.Attributes.ClientSideStalled != "" {
		attrs["stalled"] = string(element.Attributes.ClientSideStalled)
	}

	if element.Attributes.ClientSideSubmit != "" {
		attrs["submit"] = string(element.Attributes.ClientSideSubmit)
	}

	if element.Attributes.ClientSideSuspend != "" {
		attrs["suspend"] = string(element.Attributes.ClientSideSuspend)
	}

	if element.Attributes.ClientSideTimeupdate != "" {
		attrs["timeupdate"] = string(element.Attributes.ClientSideTimeupdate)
	}

	if element.Attributes.ClientSideToggle != "" {
		attrs["toggle"] = string(element.Attributes.ClientSideToggle)
	}

	if element.Attributes.ClientSideVolumechange != "" {
		attrs["volumechange"] = string(element.Attributes.ClientSideVolumechange)
	}

	if element.Attributes.ClientSideWaiting != "" {
		attrs["waiting"] = string(element.Attributes.ClientSideWaiting)
	}

	if element.Attributes.ClientSideWheel != "" {
		attrs["wheel"] = string(element.Attributes.ClientSideWheel)
	}

	var virtualDOMElement interface{}
	if len(eventHandlers) == 0 {
		tag := engine.NewTag("input", attrs)
		virtualDOMElement = tag
	} else {
		component := engine.NewComponent("input", attrs)

		for evt, h := range eventHandlers {
			component.Add(engine.On(evt, h))
		}
		virtualDOMElement = component
	}

	return virtualDOMElement

}
