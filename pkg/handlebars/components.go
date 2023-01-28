/* cSpell:disable */

package handlebars

import "github.com/gogoracer/racer/pkg/engine"

// <a> tag defines a hyperlink, which is used to link from one page to another.
func HyperlinkComponent(elements ...any) *engine.Component {
	return engine.NewComponent("a", elements...)
}

// <abbr> tag defines an abbreviation or an acronym, like "Mr.", "Dec.", "ASAP", "ATM".
func AbbreviationComponent(elements ...any) *engine.Component {
	return engine.NewComponent("abbr", elements...)
}

// <address> tag defines the contact information for its nearest article or body element ancestor.
func AddressComponent(elements ...any) *engine.Component {
	return engine.NewComponent("address", elements...)
}

// <area> tag defines a hot-spot region on an image, and optionally associates it with a hypertext link.
func AreaComponent(elements ...any) *engine.Component {
	return engine.NewComponent("area", elements...)
}

// <article> element represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable.
func ArticleComponent(elements ...any) *engine.Component {
	return engine.NewComponent("article", elements...)
}

// <aside> element represents a section of the web page that encloses content which is tangentially related to the content around it.
func AsideComponent(elements ...any) *engine.Component {
	return engine.NewComponent("aside", elements...)
}

// <audio> element is used to embed audio content in an HTML document without requiring any additional plug-in like Flash player.
func AudioComponent(elements ...any) *engine.Component {
	return engine.NewComponent("audio", elements...)
}

// <b> (short for bold) tag displays text in a bold style. This element typically renders the text it encloses in a bold typeface without conveying any extra importance.
func BoldComponent(elements ...any) *engine.Component { return engine.NewComponent("b", elements...) }

// <base> tag defines the base URL and a common target to for all relative URLs contained within a document. There must be no more than one <base> tag per document.
func BaseComponent(elements ...any) *engine.Component {
	return engine.NewComponent("base", elements...)
}

// <bdi> (stands for bi-directional isolation) element represents a span of text that is to be isolated from its surroundings for the purposes of bidirectional text formatting.
func BidirectionalIsolationComponent(elements ...any) *engine.Component {
	return engine.NewComponent("bdi", elements...)
}

// <bdo> (stands for bi-directional override) element represents the override of the current directionality.
func BidirectionalOverrideComponent(elements ...any) *engine.Component {
	return engine.NewComponent("bdo", elements...)
}

// <blockquote> tag defines a section that is quoted from another source.
func BlockQuoteComponent(elements ...any) *engine.Component {
	return engine.NewComponent("blockquote", elements...)
}

// <br> tag produces a line break in text.
func LineBreakComponent(elements ...any) *engine.Component {
	return engine.NewComponent("br", elements...)
}

// <button> tag defines a clickable button.
func ButtonComponent(elements ...any) *engine.Component {
	return engine.NewComponent("button", elements...)
}

// <canvas> tag is used to draw graphics, on the fly, via scripting (usually JavaScript).
func CanvasComponent(elements ...any) *engine.Component {
	return engine.NewComponent("canvas", elements...)
}

// <caption> tag defines a table caption.
func CaptionComponent(elements ...any) *engine.Component {
	return engine.NewComponent("caption", elements...)
}

// <cite> tag defines the title of a work (e.g. a book, a song, a movie, a painting, a sculpture, etc.).
func CitationComponent(elements ...any) *engine.Component {
	return engine.NewComponent("cite", elements...)
}

// <code> tag defines a piece of computer code.
func CodeComponent(elements ...any) *engine.Component {
	return engine.NewComponent("code", elements...)
}

// <col> tag defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.
func ColumnComponent(elements ...any) *engine.Component {
	return engine.NewComponent("col", elements...)
}

// <colgroup> tag defines a group of columns within a table.
func ColumnGroupComponent(elements ...any) *engine.Component {
	return engine.NewComponent("colgroup", elements...)
}

// <data> tag links a given content with a machine-readable translation. If the content is time- or date-related, the <time> element must be used.
func DataComponent(elements ...any) *engine.Component {
	return engine.NewComponent("data", elements...)
}

// <datalist> tag specifies a list of pre-defined options for input controls.
func DataListComponent(elements ...any) *engine.Component {
	return engine.NewComponent("datalist", elements...)
}

// <dd> tag defines a description/value of a term in a description list.
func DefinitionDescriptionComponent(elements ...any) *engine.Component {
	return engine.NewComponent("dd", elements...)
}

// <del> tag defines text that has been deleted from a document.
func DeletedComponent(elements ...any) *engine.Component {
	return engine.NewComponent("del", elements...)
}

// <details> tag specifies additional details that the user can view or hide on demand.
func DetailsComponent(elements ...any) *engine.Component {
	return engine.NewComponent("details", elements...)
}

// <dfn> (short for definition) tag represents the defining instance of a term.
func DefinitionComponent(elements ...any) *engine.Component {
	return engine.NewComponent("dfn", elements...)
}

// <dialog> tag defines a dialog box or window.
func DialogComponent(elements ...any) *engine.Component {
	return engine.NewComponent("dialog", elements...)
}

// <div> tag defines a section in a document.
func DivisionComponent(elements ...any) *engine.Component {
	return engine.NewComponent("div", elements...)
}

// <dl> (short for definition list) tag defines a description list.
func DefinitionListComponent(elements ...any) *engine.Component {
	return engine.NewComponent("dl", elements...)
}

// <dt> (short for definition term) tag represent a term or an item in a definition list.
func DefinitionTermComponent(elements ...any) *engine.Component {
	return engine.NewComponent("dt", elements...)
}

// <em> (short for emphasis) tag is used to emphasize the text content.
func EmphasisComponent(elements ...any) *engine.Component {
	return engine.NewComponent("em", elements...)
}

// <embed> tag defines a container for an external application or interactive content (a plug-in).
func EmbedComponent(elements ...any) *engine.Component {
	return engine.NewComponent("embed", elements...)
}

// <fieldset> tag is used to group several controls as well as labels (<label>) within a web form.
func FieldSetComponent(elements ...any) *engine.Component {
	return engine.NewComponent("fieldset", elements...)
}

// <figcaption> tag represents a caption or legend describing the rest of the contents of its parent <figure> element.
func FigureCaptionComponent(elements ...any) *engine.Component {
	return engine.NewComponent("figcaption", elements...)
}

// <figure> tag specifies self-contained content, potentially with an optional caption, which is specified using the <figcaption> element.
func FigureComponent(elements ...any) *engine.Component {
	return engine.NewComponent("figure", elements...)
}

// <footer> tag represents a footer for its nearest sectioning content or sectioning root element.
func FooterComponent(elements ...any) *engine.Component {
	return engine.NewComponent("footer", elements...)
}

// <form> tag represents a document section containing interactive controls for submitting information to a web server.
func FormComponent(elements ...any) *engine.Component {
	return engine.NewComponent("form", elements...)
}

// <hgroup> tag represents a multi-level heading for a section of a document.
func HeadingGroupComponent(elements ...any) *engine.Component {
	return engine.NewComponent("hgroup", elements...)
}

// <h1> is the highest section level
func Heading1Component(elements ...any) *engine.Component {
	return engine.NewComponent("h1", elements...)
}

// <h2> is the second highest section level
func Heading2Component(elements ...any) *engine.Component {
	return engine.NewComponent("h2", elements...)
}

// <h3> is the third highest section level
func Heading3Component(elements ...any) *engine.Component {
	return engine.NewComponent("h3", elements...)
}

// <h4> is the fourth highest section level
func Heading4Component(elements ...any) *engine.Component {
	return engine.NewComponent("h4", elements...)
}

// <h5> is the fifth highest section level
func Heading5Component(elements ...any) *engine.Component {
	return engine.NewComponent("h5", elements...)
}

// <h6> is the sixth highest section level
func Heading6Component(elements ...any) *engine.Component {
	return engine.NewComponent("h6", elements...)
}

// <hr> tag defines a thematic break in an HTML page (e.g. a shift of topic).
func HorizontalRuleComponent(elements ...any) *engine.Component {
	return engine.NewComponent("hr", elements...)
}

// <i> tag defines a part of text in an alternate voice or mood. The content inside is typically displayed in italic type.
func ItalicComponent(elements ...any) *engine.Component { return engine.NewComponent("i", elements...) }

// <iframe> tag defines an inline frame used to embed another document within the current HTML document.
func InlineFrameComponent(elements ...any) *engine.Component {
	return engine.NewComponent("iframe", elements...)
}

// <img> tag defines an image in an HTML page.
func ImageComponent(elements ...any) *engine.Component {
	return engine.NewComponent("img", elements...)
}

// <input> tag specifies an input field where the user can enter data.
func InputComponent(elements ...any) *engine.Component {
	return engine.NewComponent("input", elements...)
}

// <ins> (short for insert) tag defines a range of text that has been added to a document.
func InsertedComponent(elements ...any) *engine.Component {
	return engine.NewComponent("ins", elements...)
}

// <kbd> tag defines keyboard text.
func KeyboardComponent(elements ...any) *engine.Component {
	return engine.NewComponent("kbd", elements...)
}

// <keygen> tag defines a key-pair generator field (for forms).
func KeygenComponent(elements ...any) *engine.Component {
	return engine.NewComponent("keygen", elements...)
}

// <label> tag defines a label for an <input> element.
func LabelComponent(elements ...any) *engine.Component {
	return engine.NewComponent("label", elements...)
}

// <legend> tag defines a caption for the <fieldset> element.
func LegendComponent(elements ...any) *engine.Component {
	return engine.NewComponent("legend", elements...)
}

// <li> tag defines a list item.
func ListItemComponent(elements ...any) *engine.Component {
	return engine.NewComponent("li", elements...)
}

// <main> tag specifies the main content of a document.
func MainComponent(elements ...any) *engine.Component {
	return engine.NewComponent("main", elements...)
}

// <map> tag is used with <area> elements to define an image map (a clickable link area).
func MapComponent(elements ...any) *engine.Component { return engine.NewComponent("map", elements...) }

// <mark> tag defines marked/highlighted text.
func MarkedComponent(elements ...any) *engine.Component {
	return engine.NewComponent("mark", elements...)
}

// <menu> tag defines a list/menu of commands.
func MenuComponent(elements ...any) *engine.Component {
	return engine.NewComponent("menu", elements...)
}

// <meta> tag provides metadata about the HTML document. Metadata will not be displayed on the page, but will be machine parsable.
func MetaComponent(elements ...any) *engine.Component {
	return engine.NewComponent("meta", elements...)
}

// <meter> tag defines a scalar measurement within a known range (a gauge).
func MeterComponent(elements ...any) *engine.Component {
	return engine.NewComponent("meter", elements...)
}

// <nav> tag defines a set of navigation links.
func NavigationComponent(elements ...any) *engine.Component {
	return engine.NewComponent("nav", elements...)
}

// <noscript> tag defines an alternate content for users that do not support client-side scripts.
func NoScriptComponent(elements ...any) *engine.Component {
	return engine.NewComponent("noscript", elements...)
}

// <object> tag defines an embedded object within an HTML document.
func ObjectComponent(elements ...any) *engine.Component {
	return engine.NewComponent("object", elements...)
}

// <ol> tag defines an ordered list.
func OrderedListComponent(elements ...any) *engine.Component {
	return engine.NewComponent("ol", elements...)
}

// <optgroup> tag defines a group of related options in a drop-down list.
func OptionGroupComponent(elements ...any) *engine.Component {
	return engine.NewComponent("optgroup", elements...)
}

// <option> tag defines an option in a drop-down list.
func OptionComponent(elements ...any) *engine.Component {
	return engine.NewComponent("option", elements...)
}

// <output> tag represents the result of a calculation.
func OutputComponent(elements ...any) *engine.Component {
	return engine.NewComponent("output", elements...)
}

// <p> tag defines a paragraph.
func ParagraphComponent(elements ...any) *engine.Component {
	return engine.NewComponent("p", elements...)
}

// <param> tag defines a parameter for an object.
func ParameterComponent(elements ...any) *engine.Component {
	return engine.NewComponent("param", elements...)
}

// <picture> tag specifies multiple source for img element
func PictureComponent(elements ...any) *engine.Component {
	return engine.NewComponent("picture", elements...)
}

// <pre> tag defines preformatted text.
func PreformattedComponent(elements ...any) *engine.Component {
	return engine.NewComponent("pre", elements...)
}

// <progress> tag represents the progress of a task.
func ProgressComponent(elements ...any) *engine.Component {
	return engine.NewComponent("progress", elements...)
}

// <q> tag defines a short quotation.
func QuotationComponent(elements ...any) *engine.Component {
	return engine.NewComponent("q", elements...)
}

// <rp> tag defines what to show in browsers that do not support ruby annotations.
func RubyParenthesisComponent(elements ...any) *engine.Component {
	return engine.NewComponent("rp", elements...)
}

// <rt> tag defines an explanation/pronunciation of characters (for East Asian typography).
func RubyTextComponent(elements ...any) *engine.Component {
	return engine.NewComponent("rt", elements...)
}

// <ruby> tag defines a ruby annotation (for East Asian typography).
func RubyComponent(elements ...any) *engine.Component {
	return engine.NewComponent("ruby", elements...)
}

// <samp> tag defines sample output from a computer program.
func SampleComponent(elements ...any) *engine.Component {
	return engine.NewComponent("samp", elements...)
}

// <script> tag is used to define a client-side script (JavaScript).
func ScriptComponent(elements ...any) *engine.Component {
	return engine.NewComponent("script", elements...)
}

// <section> tag defines sections in a document, such as chapters, headers, footers, or any other sections of the document.
func SectionComponent(elements ...any) *engine.Component {
	return engine.NewComponent("section", elements...)
}

// <select> tag defines a drop-down list.
func SelectComponent(elements ...any) *engine.Component {
	return engine.NewComponent("select", elements...)
}

// <small> tag defines smaller text.
func SmallComponent(elements ...any) *engine.Component {
	return engine.NewComponent("small", elements...)
}

// <source> tag defines multiple media resources for media elements (<video> and <audio>).
func SourceComponent(elements ...any) *engine.Component {
	return engine.NewComponent("source", elements...)
}

// <span> tag is used to group inline-elements in a document.
func SpanComponent(elements ...any) *engine.Component {
	return engine.NewComponent("span", elements...)
}

// <strong> tag defines important text
func StrongComponent(elements ...any) *engine.Component {
	return engine.NewComponent("strong", elements...)
}

// <style> tag defines style information for a document
func StyleComponent(elements ...any) *engine.Component {
	return engine.NewComponent("style", elements...)
}

// <sub> tag defines subscripted text
func SubscriptComponent(elements ...any) *engine.Component {
	return engine.NewComponent("sub", elements...)
}

// <summary> tag defines a visible heading for a <details> element
func SummaryComponent(elements ...any) *engine.Component {
	return engine.NewComponent("summary", elements...)
}

// <sup> tag defines superscripted text
func SuperscriptComponent(elements ...any) *engine.Component {
	return engine.NewComponent("sup", elements...)
}

// <table> tag defines an HTML table
func TableComponent(elements ...any) *engine.Component {
	return engine.NewComponent("table", elements...)
}

// <tbody> tag groups the body content in a table
func TableBodyComponent(elements ...any) *engine.Component {
	return engine.NewComponent("tbody", elements...)
}

// <td> tag defines a standard cell in an HTML table
func TableDataComponent(elements ...any) *engine.Component {
	return engine.NewComponent("td", elements...)
}

// <textarea> tag defines a multi-line text input control
func TextAreaComponent(elements ...any) *engine.Component {
	return engine.NewComponent("textarea", elements...)
}

// <tfoot> tag groups the footer content in a table
func TableFooterComponent(elements ...any) *engine.Component {
	return engine.NewComponent("tfoot", elements...)
}

// <th> tag defines a header cell in an HTML table
func TableHeadComponent(elements ...any) *engine.Component {
	return engine.NewComponent("th", elements...)
}

// <thead> tag groups the header content in a table
func TableHeaderComponent(elements ...any) *engine.Component {
	return engine.NewComponent("thead", elements...)
}

// <time> tag defines a date/time
func TimeComponent(elements ...any) *engine.Component {
	return engine.NewComponent("time", elements...)
}

// <title> tag defines a title for the document
func TitleComponent(elements ...any) *engine.Component {
	return engine.NewComponent("title", elements...)
}

// <tr> tag defines a row in an HTML table
func TableRowComponent(elements ...any) *engine.Component {
	return engine.NewComponent("tr", elements...)
}

// <track> tag defines text tracks for media elements (<video> and <audio>)
func TrackComponent(elements ...any) *engine.Component {
	return engine.NewComponent("track", elements...)
}

// <u> tag defines text that should be stylistically different from normal text
func UnderlineComponent(elements ...any) *engine.Component {
	return engine.NewComponent("u", elements...)
}

// <ul> tag defines an unordered list
func UnorderedListComponent(elements ...any) *engine.Component {
	return engine.NewComponent("ul", elements...)
}

// <var> tag defines a variable
func VariableComponent(elements ...any) *engine.Component {
	return engine.NewComponent("var", elements...)
}

// <video> tag defines a video or movie
func VideoComponent(elements ...any) *engine.Component {
	return engine.NewComponent("video", elements...)
}

// <wbr> tag defines a possible line-break
func WordBreakOpportunityComponent(elements ...any) *engine.Component {
	return engine.NewComponent("wbr", elements...)
}
