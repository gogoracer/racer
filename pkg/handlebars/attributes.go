/* cSpell:disable */

package handlebars

import (
	"fmt"

	"github.com/gogoracer/racer/pkg/engine"
	"k8s.io/apimachinery/pkg/util/sets"
)

func attributes(keyValues ...string) engine.Attrs {
	attr := engine.Attrs{}

	if len(keyValues)%2 != 0 {
		panic("Attributes must be called with an even number of arguments")
	}

	for i := 0; i < len(keyValues); i += 2 {
		attr[keyValues[i]] = keyValues[i+1]
	}
	return attr
}

const (
	AttrAccept           = "accept"           // <input>	Specifies the types of files that the server accepts (only for type="file")
	AttrAcceptCharset    = "accept-charset"   // <form>	Specifies the character encodings that are to be used for the form submission
	AttrAccesskey        = "accesskey"        // Global Attributes	Specifies a shortcut key to activate/focus an element
	AttrAction           = "action"           // <form>	Specifies where to send the form-data when a form is submitted
	AttrAlign            = "align"            // Not supported in HTML 5.	Specifies the alignment according to surrounding elements. Use CSS instead
	AttrAlt              = "alt"              // <area>, <img>, <input>	Specifies an alternate text when the original element fails to display
	AttrAsync            = "async"            // <script>	Specifies that the script is executed asynchronously (only for external scripts)
	AttrAutocomplete     = "autocomplete"     // <form>, <input>	Specifies whether the <form> or the <input> element should have autocomplete enabled
	AttrAutofocus        = "autofocus"        // <button>, <input>, <select>, <textarea>	Specifies that the element should automatically get focus when the page loads
	AttrAutoplay         = "autoplay"         // <audio>, <video>	Specifies that the audio/video will start playing as soon as it is ready
	AttrBgcolor          = "bgcolor"          // Not supported in HTML 5.	Specifies the background color of an element. Use CSS instead
	AttrBorder           = "border"           // Not supported in HTML 5.	Specifies the width of the border of an element. Use CSS instead
	AttrCharset          = "charset"          // <meta>, <script>	Specifies the character encoding
	AttrChecked          = "checked"          // <input>	Specifies that an <input> element should be pre-selected when the page loads (for type="checkbox" or type="radio")
	AttrCite             = "cite"             // <blockquote>, <del>, <ins>, <q>	Specifies a URL which explains the quote/deleted/inserted text
	AttrClass            = "class"            // Global Attributes	Specifies one or more class names for an element (refers to a class in a style sheet)
	AttrColor            = "color"            // Not supported in HTML 5.	Specifies the text color of an element. Use CSS instead
	AttrCols             = "cols"             // <textarea>	Specifies the visible width of a text area
	AttrColspan          = "colspan"          // <td>, <th>	Specifies the number of columns a table cell should span
	AttrContent          = "content"          // <meta>	Gives the value associated with the http-equiv or name attribute
	AttrContenteditable  = "contenteditable"  // Global Attributes	Specifies whether the content of an element is editable or not
	AttrControls         = "controls"         // <audio>, <video>	Specifies that audio/video controls should be displayed (such as a play/pause button etc)
	AttrCoords           = "coords"           // <area>	Specifies the coordinates of the area
	AttrData             = "data"             // <object>	Specifies the URL of the resource to be used by the object
	AttrDatetime         = "datetime"         // <del>, <ins>, <time>	Specifies the date and time
	AttrDefault          = "default"          // <track>	Specifies that the track is to be enabled if the user's preferences do not indicate that another track would be more appropriate
	AttrDefer            = "defer"            // <script>	Specifies that the script is executed when the page has finished parsing (only for external scripts)
	AttrDir              = "dir"              // Global Attributes	Specifies the text direction for the content in an element
	AttrDirname          = "dirname"          // <input>, <textarea>	Specifies that the text direction will be submitted
	AttrDisabled         = "disabled"         // <button>, <fieldset>, <input>, <optgroup>, <option>, <select>, <textarea>	Specifies that the specified element/group of elements should be disabled
	AttrDownload         = "download"         // <a>, <area>	Specifies that the target will be downloaded when a user clicks on the hyperlink
	AttrDraggable        = "draggable"        // Global Attributes	Specifies whether an element is draggable or not
	AttrEnctype          = "enctype"          // <form>	Specifies how the form-data should be encoded when submitting it to the server (only for method="post")
	AttrFor              = "for"              // <label>, <output>	Specifies which form element(s) a label/calculation is bound to
	AttrForm             = "form"             // <button>, <fieldset>, <input>, <label>, <meter>, <object>, <output>, <select>, <textarea>	Specifies the name of the form the element belongs to
	AttrFormaction       = "formaction"       // <button>, <input>	Specifies where to send the form-data when a form is submitted. Only for type="submit"
	AttrFormenctype      = "formenctype"      // <button>, <input>	Specifies how form-data should be encoded before sending it to a server. Only for type="submit"
	AttrFormmethod       = "formmethod"       // <button>, <input>	Specifies how to send the form-data (which HTTP method to use). Only for type="submit"
	AttrFormnovalidate   = "formnovalidate"   // <button>, <input>	Specifies that the form-data should not be validated on submission. Only for type="submit"
	AttrFormtarget       = "formtarget"       // <button>, <input>	Specifies where to display the response after submitting the form. Only for type="submit"
	AttrHeaders          = "headers"          // <td>, <th>	Specifies one or more headers cells a cell is related to
	AttrHeight           = "height"           // <canvas>, <embed>, <iframe>, <img>, <input>, <object>, <video>	Specifies the height of the element
	AttrHidden           = "hidden"           // Global Attributes	Specifies that an element is not yet, or is no longer, relevant
	AttrHigh             = "high"             // <meter>	Specifies the range that is considered to be a high value
	AttrHref             = "href"             // <a>, <area>, <base>, <link>	Specifies the URL of the page the link goes to
	AttrHreflang         = "hreflang"         // <a>, <area>, <link>	Specifies the language of the linked document
	AttrHttpEquiv        = "http-equiv"       //	<meta>	Provides an HTTP header for the information/value of the content attribute
	AttrId               = "id"               // Global Attributes	Specifies a unique id for an element
	AttrIsmap            = "ismap"            // <img>	Specifies an image as a server-side image map
	AttrKind             = "kind"             // <track>	Specifies the kind of text track
	AttrLabel            = "label"            // <track>, <option>, <optgroup>	Specifies the title of the text track
	AttrLang             = "lang"             // Global Attributes	Specifies the language of the element's content
	AttrList             = "list"             // <input>	Refers to a <datalist> element that contains pre-defined options for an <input> element
	AttrLoop             = "loop"             // <audio>, <video>	Specifies that the audio/video will start over again, every time it is finished
	AttrLow              = "low"              // <meter>	Specifies the range that is considered to be a low value
	AttrMax              = "max"              // <input>, <meter>, <progress>	Specifies the maximum value
	AttrMaxlength        = "maxlength"        // <input>, <textarea>	Specifies the maximum number of characters allowed in an element
	AttrMedia            = "media"            // <a>, <area>, <link>, <source>, <style>	Specifies what media/device the linked document is optimized for
	AttrMethod           = "method"           // <form>	Specifies the HTTP method to use when sending form-data
	AttrMin              = "min"              // <input>, <meter>	Specifies a minimum value
	AttrMinlength        = "minlength"        // <input>, <textarea>	Specifies a minimum number of characters allowed in an element
	AttrMultiple         = "multiple"         // <input>, <select>	Specifies that a user can enter more than one value
	AttrMuted            = "muted"            // <video>, <audio>	Specifies that the audio output of the video should be muted
	AttrName             = "name"             // <button>, <fieldset>, <form>, <iframe>, <input>, <map>, <meta>, <object>, <output>, <param>, <select>, <textarea>	Specifies the name of the element
	AttrNovalidate       = "novalidate"       // <form>	Specifies that the form should not be validated when submitted
	AttrOnAbort          = "onabort"          // <audio>, <embed>, <img>, <object>, <video>	Script to be run on abort
	AttrOnAfterprint     = "onafterprint"     // <body>	Script to be run after the document is printed
	AttrOnBeforeprint    = "onbeforeprint"    // <body>	Script to be run before the document is printed
	AttrOnBeforeunload   = "onbeforeunload"   // <body>	Script to be run when the document is about to be unloaded
	AttrOnBlur           = "onblur"           // All visible elements.	Script to be run when the element loses focus
	AttrOnCanplay        = "oncanplay"        // <audio>, <embed>, <object>, <video>	Script to be run when a file is ready to start playing (when it has buffered enough to begin)
	AttrOnCanplaythrough = "oncanplaythrough" // <audio>, <video>	Script to be run when a file can be played all the way to the end without pausing for buffering
	AttrOnChange         = "onchange"         // All visible elements.	Script to be run when the value of the element is changed
	AttrOnClick          = "onclick"          // All visible elements.	Script to be run when the element is being clicked
	AttrOnContextmenu    = "oncontextmenu"    // All visible elements.	Script to be run when a context menu is triggered
	AttrOnCopy           = "oncopy"           // All visible elements.	Script to be run when the content of the element is being copied
	AttrOnCuechange      = "oncuechange"      // <track>	Script to be run when the cue changes in a <track> element
	AttrOnCut            = "oncut"            // All visible elements.	Script to be run when the content of the element is being cut
	AttrOnDblclick       = "ondblclick"       // All visible elements.	Script to be run when the element is being double-clicked
	AttrOnDrag           = "ondrag"           // All visible elements.	Script to be run when the element is being dragged
	AttrOnDragend        = "ondragend"        // All visible elements.	Script to be run at the end of a drag operation
	AttrOnDragenter      = "ondragenter"      // All visible elements.	Script to be run when an element has been dragged to a valid drop target
	AttrOnDragleave      = "ondragleave"      // All visible elements.	Script to be run when an element leaves a valid drop target
	AttrOnDragover       = "ondragover"       // All visible elements.	Script to be run when an element is being dragged over a valid drop target
	AttrOnDragstart      = "ondragstart"      // All visible elements.	Script to be run at the start of a drag operation
	AttrOnDrop           = "ondrop"           // All visible elements.	Script to be run when dragged element is being dropped
	AttrOnDurationchange = "ondurationchange" // <audio>, <video>	Script to be run when the length of the media changes
	AttrOnEmptied        = "onemptied"        // <audio>, <video>	Script to be run when something bad happens and the file is suddenly unavailable (like unexpectedly disconnects)
	AttrOnEnded          = "onended"          // <audio>, <video>	Script to be run when the media has reach the end (a useful event for messages like "thanks for listening")
	AttrOnError          = "onerror"          // <audio>, <body>, <embed>, <img>, <object>, <script>, <style>, <video>	Script to be run when an error occurs
	AttrOnFocus          = "onfocus"          // All visible elements.	Script to be run when the element gets focus
	AttrOnHashchange     = "onhashchange"     // <body>	Script to be run when there has been changes to the anchor part of the a URL
	AttrOnInput          = "oninput"          // All visible elements.	Script to be run when the element gets user input
	AttrOnInvalid        = "oninvalid"        // All visible elements.	Script to be run when the element is invalid
	AttrOnKeydown        = "onkeydown"        // All visible elements.	Script to be run when a user is pressing a key
	AttrOnKeypress       = "onkeypress"       // All visible elements.	Script to be run when a user presses a key
	AttrOnKeyup          = "onkeyup"          // All visible elements.	Script to be run when a user releases a key
	AttrOnLoad           = "onload"           // <body>, <iframe>, <img>, <input>, <link>, <script>, <style>	Script to be run when the element is finished loading
	AttrOnLoadeddata     = "onloadeddata"     // <audio>, <video>	Script to be run when media data is loaded
	AttrOnLoadedmetadata = "onloadedmetadata" // <audio>, <video>	Script to be run when meta data (like dimensions and duration) are loaded
	AttrOnLoadstart      = "onloadstart"      // <audio>, <video>	Script to be run just as the file begins to load before anything is actually loaded
	AttrOnMousedown      = "onmousedown"      // All visible elements.	Script to be run when a mouse button is pressed down on an element
	AttrOnMousemove      = "onmousedown"      // All visible elements.	Script to be run as long as the  mouse pointer is moving over an element
	AttrOnMouseout       = "onmouseout"       // All visible elements.	Script to be run when a mouse pointer moves out of an element
	AttrOnMouseover      = "onmouseover"      // All visible elements.	Script to be run when a mouse pointer moves over an element
	AttrOnMouseup        = "onmouseup"        // All visible elements.	Script to be run when a mouse button is released over an element
	AttrOnMousewheel     = "onmousewheel"     // All visible elements.	Script to be run when a mouse wheel is being scrolled over an element
	AttrOnOffline        = "onoffline"        // <body>	Script to be run when the browser starts to work offline
	AttrOnOnline         = "ononline"         // <body>	Script to be run when the browser starts to work online
	AttrOnPagehide       = "onpagehide"       // <body>	Script to be run when a user navigates away from a page
	AttrOnPageshow       = "onpageshow"       // <body>	Script to be run when a user navigates to a page
	AttrOnPaste          = "onpaste"          // All visible elements.	Script to be run when the user pastes some content in an element
	AttrOnPause          = "onpause"          // <audio>, <video>	Script to be run when the media is paused either by the user or programmatically
	AttrOnPlay           = "onplay"           // <audio>, <video>	Script to be run when the media has started playing
	AttrOnPlaying        = "onplaying"        // <audio>, <video>	Script to be run when the media has started playing
	AttrOnPopstate       = "onpopstate"       // <body>	Script to be run when the window's history changes.
	AttrOnProgress       = "onprogress"       // <audio>, <video>	Script to be run when the browser is in the process of getting the media data
	AttrOnRatechange     = "onratechange"     // <audio>, <video>	Script to be run each time the playback rate changes (like when a user switches to a slow motion or fast forward mode).
	AttrOnReset          = "onreset"          // <form>	Script to be run when a reset button in a form is clicked.
	AttrOnResize         = "onresize"         // <body>	Script to be run when the browser window is being resized.
	AttrOnScroll         = "onscroll"         // All visible elements.	Script to be run when an element's scrollbar is being scrolled
	AttrOnSearch         = "onsearch"         // <input>	Script to be run when the user writes something in a search field (for <input type="search">)
	AttrOnSeeked         = "onseeked"         // <audio>, <video>	Script to be run when the seeking attribute is set to false indicating that seeking has ended
	AttrOnSeeking        = "onseeking"        // <audio>, <video>	Script to be run when the seeking attribute is set to true indicating that seeking is active
	AttrOnSelect         = "onselect"         // All visible elements.	Script to be run when the element gets selected
	AttrOnStalled        = "onstalled"        // <audio>, <video>	Script to be run when the browser is unable to fetch the media data for whatever reason
	AttrOnStorage        = "onstorage"        // <body>	Script to be run when a Web Storage area is updated
	AttrOnSubmit         = "onsubmit"         // <form>	Script to be run when a form is submitted
	AttrOnSuspend        = "onsuspend"        // <audio>, <video>	Script to be run when fetching the media data is stopped before it is completely loaded for whatever reason
	AttrOnTimeupdate     = "ontimeupdate"     // <audio>, <video>	Script to be run when the playing position has changed (like when the user fast forwards to a different point in the media)
	AttrOnToggle         = "ontoggle"         // <details>	Script to be run when the user opens or closes the <details> element
	AttrOnUnload         = "onunload"         // <body>	Script to be run when a page has unloaded (or the browser window has been closed)
	AttrOnVolumechange   = "onvolumechange"   // <audio>, <video>	Script to be run each time the volume of a video/audio has been changed
	AttrOnWaiting        = "onwaiting"        // <audio>, <video>	Script to be run when the media has paused but is expected to resume (like when the media pauses to buffer more data)
	AttrOnWheel          = "onwheel"          // All visible elements.	Script to be run when the mouse wheel rolls up or down over an element
	AttrOpen             = "open"             // <details>	Specifies that the details should be visible (open) to the user
	AttrOptimum          = "optimum"          // <meter>	Specifies what value is the optimal value for the gauge
	AttrPattern          = "pattern"          // <input>	Specifies a regular expression that an <input> element's value is checked against
	AttrPing             = "ping"             // <a>, <area>	Specifies a space-separated list of URLs to ping
	AttrPlaceholder      = "placeholder"      // <input>, <textarea>	Specifies a short hint that describes the expected value of the element
	AttrPoster           = "poster"           // <video>	Specifies an image to be shown while the video is downloading, or until the user hits the play button
	AttrPreload          = "preload"          // <audio>, <video>	Specifies if and how the author thinks the audio/video should be loaded when the page loads
	AttrReadonly         = "readonly"         // <input>, <textarea>	Specifies that the element is read-only
	AttrReferrerpolicy   = "referrerpolicy"   // <a>, <area>, <iframe>, <img>, <link>, <script>	Specifies which referrer to use when fetching the resource
	AttrRel              = "rel"              // <a>, <area>, <form>, <link>	Specifies the relationship between the current document and the linked document
	AttrRequired         = "required"         // <input>, <select>, <textarea>	Specifies that the element must be filled out before submitting the form
	AttrReversed         = "reversed"         // <ol>	Specifies that the list order should be descending (9,8,7...)
	AttrRows             = "rows"             // <textarea>	Specifies the visible number of lines in a text area
	AttrRowspan          = "rowspan"          // <td>, <th>	Specifies the number of rows a table cell should span
	AttrSandbox          = "sandbox"          // <iframe>	Enables an extra set of restrictions for the content in an <iframe>
	AttrScope            = "scope"            // <th>	Specifies whether a header cell is a header for a column, row, or group of columns or rows
	AttrSelected         = "selected"         // <option>	Specifies that an option should be pre-selected when the page loads
	AttrShape            = "shape"            // <area>	Specifies the shape of the area
	AttrSize             = "size"             // <input>, <select>	Specifies the width, in characters (for <input>) or specifies the number of visible options (for <select>)
	AttrSizes            = "sizes"            // <img>, <link>, <source>	Specifies the size of the linked resource
	AttrSpan             = "span"             // <col>, <colgroup>	Specifies the number of columns to span
	AttrSpellcheck       = "spellcheck"       // Global Attributes	Specifies whether the element is to have its spelling and grammar checked or not
	AttrSrc              = "src"              // <audio>, <embed>, <iframe>, <img>, <input>, <script>, <source>, <track>, <video>	Specifies the URL of the media file
	AttrSrcdoc           = "srcdoc"           // <iframe>	Specifies the HTML content of the page to show in the <iframe>
	AttrSrclang          = "srclang"          // <track>	Specifies the language of the track text data (required if kind="subtitles")
	AttrSrcset           = "srcset"           // <img>, <source>	Specifies the URL of the image to use in different situations
	AttrStart            = "start"            // <ol>	Specifies the start value of an ordered list
	AttrStep             = "step"             // <input>	Specifies the legal number intervals for an input field
	AttrStyle            = "style"            // Global Attributes	Specifies an inline CSS style for an element
	AttrTabindex         = "tabindex"         // Global Attributes	Specifies the tabbing order of an element
	AttrTarget           = "target"           // <a>, <area>, <base>, <form>	Specifies the target for where to open the linked document or where to submit the form
	AttrTitle            = "title"            // Global Attributes	Specifies extra information about an element
	AttrTranslate        = "translate"        // Global Attributes	Specifies whether the content of an element should be translated or not
	AttrType             = "type"             // <a>, <button>, <embed>, <input>, <link>, <menu>, <object>, <script>, <source>, <style>	Specifies the type of element
	AttrUsemap           = "usemap"           // <img>, <object>	Specifies an image as a client-side image map
	AttrValue            = "value"            // <button>, <input>, <li>, <option>, <meter>, <progress>, <param>	Specifies the value of the element
	AttrWidth            = "width"            // <canvas>, <embed>, <iframe>, <img>, <input>, <object>, <video>	Specifies the width of the element
	AttrWrap             = "wrap"             // <textarea>	Specifies how the text in a text area is to be wrapped when submitted in a form
	AttrManifest         = "manifest"         // <html>	Specifies the address of the document's cache manifest (a file that provides information about the resources of the page and lists the resources that need to be cached)
	AttrXMLNS            = "xmlns"            // <html>	Specifies the namespace of an element
	AttrVersion          = "version"          // <html>	Specifies the version of the document type definition to use
	AttrSeamless         = "seamless"         // <iframe>	Specifies that the iframe should be seamless, that is, without borders
	AttrCrossorigin      = "crossorigin"      // <img>, <link>, <script>, <audio>, <video>, <source>, <track>	Specifies how the element handles crossorigin requests
	AttrChallenge        = "challenge"        // <keygen>	Specifies the challenge that the server will accept in order to authenticate the user
	AttrKeytype          = "keytype"          // <keygen>	Specifies the type of key generated
	AttrDeclare          = "declare"          // <object>	Specifies that the object is declared, but is not yet instantiated, until script code requests it
	AttrScoped           = "scoped"           // <style>	Specifies that the style should only be applied to this element's parent element and that element's child elements
	AttrSortable         = "sortable"         // <table>	Specifies that the content of the element should be editable
	AttrSorted           = "sorted"           // <table>	Specifies that the content of the element should be editable
	AttrAbbr             = "abbr"             // <td>, <th>	Specifies an abbreviated version of the content in a cell
)

func AttrDataValue(val string) string {
	return "data-" + val
}

func attrKeys(keyValues ...string) []string {
	if len(keyValues)%2 != 0 {
		panic("attributes must be key value pairs")
	}
	var keys []string
	for i := 0; i < len(keyValues); i += 2 {
		keys = append(keys, keyValues[i])
	}
	return keys
}

var globalAttrs = sets.NewString(
	AttrAccesskey,
	AttrClass,
	AttrContenteditable,
	AttrDir,
	AttrDraggable,
	AttrHidden,
	AttrId,
	AttrLang,
	AttrSpellcheck,
	AttrStyle,
	AttrTabindex,
	AttrTitle,
	AttrTranslate,
)

func globalAttributes(keyValues ...string) engine.Attrs {
	if !globalAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(globalAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var hyperlinkAttrs = sets.NewString(
	AttrDownload,
	AttrHref,
	AttrHreflang,
	AttrMedia,
	AttrPing,
	AttrRel,
	AttrTarget,
	AttrType,
).Union(globalAttrs)

func HyperlinkAttributes(keyValues ...string) engine.Attrs {
	if !hyperlinkAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(hyperlinkAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func AbbreviationAttributes(keyValues ...string) engine.Attrs {
	if !globalAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(globalAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func AddressAttributes(keyValues ...string) engine.Attrs {
	if !globalAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(globalAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var areaAttrs = sets.NewString(
	AttrAlt,
	AttrCoords,
	AttrDownload,
	AttrHref,
	AttrHreflang,
	AttrMedia,
	AttrReferrerpolicy,
	AttrRel,
	AttrPing,
	AttrRel,
	AttrShape,
	AttrTarget,
	AttrType,
).Union(globalAttrs)

func AreaAttributes(keyValues ...string) engine.Attrs {
	if !areaAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(areaAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func ArticleAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func AsideAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var audioAttrs = sets.NewString(
	AttrAutoplay,
	AttrControls,
	AttrLoop,
	AttrMuted,
	AttrPreload,
	AttrSrc,
).Union(globalAttrs)

func AudioAttributes(keyValues ...string) engine.Attrs {
	if !audioAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(audioAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func BoldAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var baseAttrs = sets.NewString(
	AttrHref,
	AttrTarget,
).Union(globalAttrs)

func BaseAttributes(keyValues ...string) engine.Attrs {
	if !baseAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(baseAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func BidirectionalIsolationAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var biDirectionalOverrideAttrs = sets.NewString(
	AttrDir,
).Union(globalAttrs)

func BiDirectionalOverrideAttributes(keyValues ...string) engine.Attrs {
	if !biDirectionalOverrideAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(biDirectionalOverrideAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var blockquoteAttrs = sets.NewString(
	AttrCite,
).Union(globalAttrs)

func BlockquoteAttributes(keyValues ...string) engine.Attrs {
	if !blockquoteAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(blockquoteAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func BodyAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var buttonAttrs = sets.NewString(
	AttrAutofocus,
	AttrDisabled,
	AttrForm,
	AttrFormaction,
	AttrFormenctype,
	AttrFormmethod,
	AttrFormnovalidate,
	AttrFormtarget,
	AttrName,
	AttrType,
	AttrValue,
	AttrOnClick,
).Union(globalAttrs)

func ButtonAttributes(keyValues ...string) engine.Attrs {
	if !buttonAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(buttonAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var canvasAttrs = sets.NewString(
	AttrHeight,
	AttrWidth,
).Union(globalAttrs)

func CanvasAttributes(keyValues ...string) engine.Attrs {
	if !canvasAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(canvasAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var captionAttrs = sets.NewString(
	AttrAlign,
).Union(globalAttrs)

func CaptionAttributes(keyValues ...string) engine.Attrs {
	if !captionAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(captionAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func CiteAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func CodeAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func ColAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func ColgroupAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var dataAttrs = sets.NewString(
	AttrValue,
).Union(globalAttrs)

func DataAttributes(keyValues ...string) engine.Attrs {
	if !dataAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(dataAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func DataListAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func DefinitionDescriptionAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var delAttrs = sets.NewString(
	AttrCite,
	AttrDatetime,
).Union(globalAttrs)

func DeletedAttributes(keyValues ...string) engine.Attrs {
	if !delAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(delAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var detailsAttrs = sets.NewString(
	AttrOpen,
).Union(globalAttrs)

func DetailsAttributes(keyValues ...string) engine.Attrs {
	if !detailsAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(detailsAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func DefinitionAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var dialogAttrs = sets.NewString(
	AttrOpen,
).Union(globalAttrs)

func DialogAttributes(keyValues ...string) engine.Attrs {
	if !dialogAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(dialogAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func DivAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func DefinitionListAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func DefinitionTermAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func EmphasisAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var embedAttrs = sets.NewString(
	AttrHeight,
	AttrSrc,
	AttrType,
	AttrWidth,
).Union(globalAttrs)

func EmbedAttributes(keyValues ...string) engine.Attrs {
	if !embedAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(embedAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func FieldsetAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func FigureCaptionAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func FigureAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func FooterAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var formAttrs = sets.NewString(
	AttrAccept,
	AttrAcceptCharset,
	AttrAction,
	AttrAutocomplete,
	AttrEnctype,
	AttrMethod,
	AttrName,
	AttrNovalidate,
	AttrTarget,
).Union(globalAttrs)

func FormAttributes(keyValues ...string) engine.Attrs {
	if !formAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(formAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func HeadAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func HeaderAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func Heading1GroupAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func Heading2Attributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func Heading3Attributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func Heading4Attributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func Heading5Attributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func Heading6Attributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func HorizontalRuleAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var htmlAttrs = sets.NewString(
	AttrManifest,
	AttrXMLNS,
	AttrVersion,
).Union(globalAttrs)

func HTMLAttributes(keyValues ...string) engine.Attrs {
	if !htmlAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(htmlAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func ItalicAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var iframeAttrs = sets.NewString(
	AttrName,
	AttrSandbox,
	AttrSeamless,
	AttrSrc,
	AttrSrcdoc,
	AttrWidth,
	AttrHeight,
).Union(globalAttrs)

func IFrameAttributes(keyValues ...string) engine.Attrs {
	if !iframeAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(iframeAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var imgAttrs = sets.NewString(
	AttrAlt,
	AttrSrc,
	AttrCrossorigin,
	AttrHeight,
	AttrIsmap,
	AttrWidth,
).Union(globalAttrs)

func ImageAttributes(keyValues ...string) engine.Attrs {
	if !imgAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(imgAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var inputAttrs = sets.NewString(
	AttrAccept,
	AttrAlt,
	AttrAutocomplete,
	AttrAutofocus,
	AttrChecked,
	AttrDisabled,
	AttrForm,
	AttrFormaction,
	AttrFormenctype,
	AttrFormmethod,
	AttrFormnovalidate,
	AttrFormtarget,
	AttrHeight,
	AttrList,
	AttrMax,
	AttrMaxlength,
	AttrMin,
	AttrMinlength,
	AttrMultiple,
	AttrName,
	AttrPattern,
	AttrPlaceholder,
	AttrReadonly,
	AttrRequired,
	AttrSize,
	AttrSrc,
	AttrStep,
	AttrType,
	AttrValue,
	AttrWidth,
).Union(globalAttrs)

func InputAttributes(keyValues ...string) engine.Attrs {
	if !inputAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(inputAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var insertAttrs = sets.NewString(
	AttrCite,
	AttrDatetime,
).Union(globalAttrs)

func InsertAttributes(keyValues ...string) engine.Attrs {
	if !insertAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(insertAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func KeyboardAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var keygenAttrs = sets.NewString(
	AttrAutofocus,
	AttrChallenge,
	AttrDisabled,
	AttrForm,
	AttrKeytype,
	AttrName,
).Union(globalAttrs)

func KeygenAttributes(keyValues ...string) engine.Attrs {
	if !keygenAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(keygenAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var labelAttrs = sets.NewString(
	AttrForm,
	AttrFor,
).Union(globalAttrs)

func LabelAttributes(keyValues ...string) engine.Attrs {
	if !labelAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(labelAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func LegendAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var liAttrs = sets.NewString(
	AttrValue,
).Union(globalAttrs)

func ListItemAttributes(keyValues ...string) engine.Attrs {
	if !liAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(liAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var linkAttrs = sets.NewString(
	AttrCrossorigin,
	AttrHref,
	AttrHreflang,
	AttrMedia,
	AttrRel,
	AttrSizes,
	AttrType,
).Union(globalAttrs)

func LinkAttributes(keyValues ...string) engine.Attrs {
	if !linkAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(linkAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func MainAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var mapAttrs = sets.NewString(
	AttrName,
).Union(globalAttrs)

func MapAttributes(keyValues ...string) engine.Attrs {
	if !mapAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(mapAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func MarkedAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var menuAttrs = sets.NewString(
	AttrLabel,
	AttrType,
).Union(globalAttrs)

func MenuAttributes(keyValues ...string) engine.Attrs {
	if !menuAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(menuAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var metaAttrs = sets.NewString(
	AttrCharset,
	AttrContent,
	AttrHttpEquiv,
	AttrName,
).Union(globalAttrs)

func MetaAttributes(keyValues ...string) engine.Attrs {
	if !metaAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(metaAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var meterAttrs = sets.NewString(
	AttrForm,
	AttrHigh,
	AttrLow,
	AttrMax,
	AttrMin,
	AttrOptimum,
	AttrValue,
).Union(globalAttrs)

func MeterAttributes(keyValues ...string) engine.Attrs {
	if !meterAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(meterAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func NavigationAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func NoScriptAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var objectAttrs = sets.NewString(
	AttrData,
	AttrDeclare,
	AttrForm,
	AttrHeight,
	AttrName,
	AttrType,
	AttrUsemap,
	AttrWidth,
).Union(globalAttrs)

func ObjectAttributes(keyValues ...string) engine.Attrs {
	if !objectAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(objectAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var olAttrs = sets.NewString(
	AttrReversed,
	AttrStart,
	AttrType,
).Union(globalAttrs)

func OrderedListAttributes(keyValues ...string) engine.Attrs {
	if !olAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(olAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var optgroupAttrs = sets.NewString(
	AttrDisabled,
	AttrLabel,
).Union(globalAttrs)

func OptGroupAttributes(keyValues ...string) engine.Attrs {
	if !optgroupAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(optgroupAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var optionAttrs = sets.NewString(
	AttrDisabled,
	AttrLabel,
	AttrSelected,
	AttrValue,
).Union(globalAttrs)

func OptionAttributes(keyValues ...string) engine.Attrs {
	if !optionAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(optionAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var outputAttrs = sets.NewString(
	AttrFor,
	AttrForm,
	AttrName,
).Union(globalAttrs)

func OutputAttributes(keyValues ...string) engine.Attrs {
	if !outputAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(outputAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func ParagraphAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var paramAttrs = sets.NewString(
	AttrName,
	AttrValue,
).Union(globalAttrs)

func ParamAttributes(keyValues ...string) engine.Attrs {
	if !paramAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(paramAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func PictureAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func PreformattedAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var progressAttrs = sets.NewString(
	AttrMax,
	AttrValue,
).Union(globalAttrs)

func ProgressAttributes(keyValues ...string) engine.Attrs {
	if !progressAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(progressAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

var qAttrs = sets.NewString(
	AttrCite,
).Union(globalAttrs)

func QuoteAttributes(keyValues ...string) engine.Attrs {
	if !qAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(qAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func RubyParenthesisAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func RubyTextAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func RubyAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var scriptAttrs = sets.NewString(
	AttrAsync,
	AttrCharset,
	AttrDefer,
	AttrSrc,
	AttrType,
).Union(globalAttrs)

func ScriptAttributes(keyValues ...string) engine.Attrs {
	if !scriptAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(scriptAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func SectionAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var selectAttrs = sets.NewString(
	AttrAutofocus,
	AttrDisabled,
	AttrForm,
	AttrMultiple,
	AttrName,
	AttrRequired,
	AttrSize,
).Union(globalAttrs)

func SelectAttributes(keyValues ...string) engine.Attrs {
	if !selectAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(selectAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func SmallAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var sourceAttrs = sets.NewString(
	AttrMedia,
	AttrSrc,
	AttrType,
).Union(globalAttrs)

func SourceAttributes(keyValues ...string) engine.Attrs {
	if !sourceAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(sourceAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func SpanAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func SampleAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func StrongAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var styleAttrs = sets.NewString(
	AttrMedia,
	AttrType,
	AttrScoped,
).Union(globalAttrs)

func StyleAttributes(keyValues ...string) engine.Attrs {
	if !styleAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(styleAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func SubscriptAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func SummaryAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func SuperscriptAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func SvgAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var tableAttrs = sets.NewString(
	AttrSortable,
).Union(globalAttrs)

func TableAttributes(keyValues ...string) engine.Attrs {
	if !tableAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(tableAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func TableBodyAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var tableDataAttrs = sets.NewString(
	AttrColspan,
	AttrHeaders,
	AttrRowspan,
).Union(globalAttrs)

func TableDataAttributes(keyValues ...string) engine.Attrs {
	if !tableDataAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(tableDataAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func TemplateAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var textAreaAttrs = sets.NewString(
	AttrAutofocus,
	AttrCols,
	AttrDisabled,
	AttrForm,
	AttrMaxlength,
	AttrMinlength,
	AttrName,
	AttrPlaceholder,
	AttrReadonly,
	AttrRequired,
	AttrRows,
	AttrWrap,
).Union(globalAttrs)

func TextAreaAttributes(keyValues ...string) engine.Attrs {
	if !textAreaAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(textAreaAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func TableFooterAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var tableHeaderAttrs = sets.NewString(
	AttrAbbr,
	AttrColspan,
	AttrHeaders,
	AttrRowspan,
	AttrScope,
	AttrSorted,
	AttrScope,
).Union(globalAttrs)

func TableHeadAttributes(keyValues ...string) engine.Attrs {
	if !tableHeaderAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(tableHeaderAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func TableHeaderAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var timeAttrs = sets.NewString(
	AttrDatetime,
).Union(globalAttrs)

func TimeAttributes(keyValues ...string) engine.Attrs {
	if !timeAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(timeAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func TitleAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func TableRowAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var trackAttrs = sets.NewString(
	AttrDefault,
	AttrKind,
	AttrLabel,
	AttrSrc,
	AttrSrclang,
).Union(globalAttrs)

func TrackAttributes(keyValues ...string) engine.Attrs {
	if !trackAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(trackAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func UnderlineAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func UnorderedListAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

func VariableAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}

var videoAttrs = sets.NewString(
	AttrAutoplay,
	AttrControls,
	AttrHeight,
	AttrLoop,
	AttrMuted,
	AttrPoster,
	AttrPreload,
	AttrSrc,
	AttrWidth,
).Union(globalAttrs)

func VideoAttributes(keyValues ...string) engine.Attrs {
	if !videoAttrs.HasAll(attrKeys(keyValues...)...) {
		diff := sets.NewString(keyValues...).Difference(videoAttrs)
		panic(fmt.Sprintf("unknown attribute %q", diff.List()))
	}
	return attributes(keyValues...)
}

func WordBreakOpportunityAttributes(keyValues ...string) engine.Attrs {
	return globalAttributes(keyValues...)
}
