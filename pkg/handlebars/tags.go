package handlebars

import "github.com/gogoracer/racer/pkg/gas"

// <a> tag defines a hyperlink, which is used to link from one page to another.
func Hyperlink(elements ...any) *gas.Tag { return gas.NewTag("a", elements...) }

// <abbr> tag defines an abbreviation or an acronym, like "Mr.", "Dec.", "ASAP", "ATM".
func Abbreviation(elements ...any) *gas.Tag { return gas.NewTag("abbr", elements...) }

// <address> tag defines the contact information for its nearest article or body element ancestor.
func Address(elements ...any) *gas.Tag { return gas.NewTag("address", elements...) }

// <area> tag defines a hot-spot region on an image, and optionally associates it with a hypertext link.
func Area(elements ...any) *gas.Tag { return gas.NewTag("area", elements...) }

// <article> element represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable.
func Article(elements ...any) *gas.Tag { return gas.NewTag("article", elements...) }

// <aside> element represents a section of the web page that encloses content which is tangentially related to the content around it.
func Aside(elements ...any) *gas.Tag { return gas.NewTag("aside", elements...) }

// <audio> element is used to embed audio content in an HTML document without requiring any additional plug-in like Flash player.
func Audio(elements ...any) *gas.Tag { return gas.NewTag("audio", elements...) }

// <b> (short for bold) tag displays text in a bold style. This element typically renders the text it encloses in a bold typeface without conveying any extra importance.
func Bold(elements ...any) *gas.Tag { return gas.NewTag("b", elements...) }

// <base> tag defines the base URL and a common target to for all relative URLs contained within a document. There must be no more than one <base> tag per document.
func Base(elements ...any) *gas.Tag { return gas.NewTag("base", elements...) }

// <bdi> (stands for bi-directional isolation) element represents a span of text that is to be isolated from its surroundings for the purposes of bidirectional text formatting.
func BidirectionalIsolation(elements ...any) *gas.Tag { return gas.NewTag("bdi", elements...) }

// <bdo> (stands for bi-directional override) element represents the override of the current directionality.
func BidirectionalOverride(elements ...any) *gas.Tag { return gas.NewTag("bdo", elements...) }

// <blockquote> tag defines a section that is quoted from another source.
func BlockQuote(elements ...any) *gas.Tag { return gas.NewTag("blockquote", elements...) }

// <br> tag produces a line break in text.
func LineBreak(elements ...any) *gas.Tag { return gas.NewTag("br", elements...) }

// <button> tag defines a clickable button.
func Button(elements ...any) *gas.Tag { return gas.NewTag("button", elements...) }

// <canvas> tag is used to draw graphics, on the fly, via scripting (usually JavaScript).
func Canvas(elements ...any) *gas.Tag { return gas.NewTag("canvas", elements...) }

// <caption> tag defines a table caption.
func Caption(elements ...any) *gas.Tag { return gas.NewTag("caption", elements...) }

// <cite> tag defines the title of a work (e.g. a book, a song, a movie, a painting, a sculpture, etc.).
func Citation(elements ...any) *gas.Tag { return gas.NewTag("cite", elements...) }

// <code> tag defines a piece of computer code.
func Code(elements ...any) *gas.Tag { return gas.NewTag("code", elements...) }

// <col> tag defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.
func Column(elements ...any) *gas.Tag { return gas.NewTag("col", elements...) }

// <colgroup> tag defines a group of columns within a table.
func ColumnGroup(elements ...any) *gas.Tag { return gas.NewTag("colgroup", elements...) }

// <data> tag links a given content with a machine-readable translation. If the content is time- or date-related, the <time> element must be used.
func Data(elements ...any) *gas.Tag { return gas.NewTag("data", elements...) }

// <datalist> tag specifies a list of pre-defined options for input controls.
func DataList(elements ...any) *gas.Tag { return gas.NewTag("datalist", elements...) }

// <dd> tag defines a description/value of a term in a description list.
func DefinitionDescription(elements ...any) *gas.Tag { return gas.NewTag("dd", elements...) }

// <del> tag defines text that has been deleted from a document.
func Deleted(elements ...any) *gas.Tag { return gas.NewTag("del", elements...) }

// <details> tag specifies additional details that the user can view or hide on demand.
func Details(elements ...any) *gas.Tag { return gas.NewTag("details", elements...) }

// <dfn> (short for definition) tag represents the defining instance of a term.
func Definition(elements ...any) *gas.Tag { return gas.NewTag("dfn", elements...) }

// <dialog> tag defines a dialog box or window.
func Dialog(elements ...any) *gas.Tag { return gas.NewTag("dialog", elements...) }

// <div> tag defines a section in a document.
func Division(elements ...any) *gas.Tag { return gas.NewTag("div", elements...) }

// <dl> (short for definition list) tag defines a description list.
func DefinitionList(elements ...any) *gas.Tag { return gas.NewTag("dl", elements...) }

// <dt> (short for definition term) tag represent a term or an item in a definition list.
func DefinitionTerm(elements ...any) *gas.Tag { return gas.NewTag("dt", elements...) }

// <em> (short for emphasis) tag is used to emphasize the text content.
func Emphasis(elements ...any) *gas.Tag { return gas.NewTag("em", elements...) }

// <embed> tag defines a container for an external application or interactive content (a plug-in).
func Embed(elements ...any) *gas.Tag { return gas.NewTag("embed", elements...) }

// <fieldset> tag is used to group several controls as well as labels (<label>) within a web form.
func FieldSet(elements ...any) *gas.Tag { return gas.NewTag("fieldset", elements...) }

// <figcaption> tag represents a caption or legend describing the rest of the contents of its parent <figure> element.
func FigureCaption(elements ...any) *gas.Tag { return gas.NewTag("figcaption", elements...) }

// <figure> tag specifies self-contained content, potentially with an optional caption, which is specified using the <figcaption> element.
func Figure(elements ...any) *gas.Tag { return gas.NewTag("figure", elements...) }

// <footer> tag represents a footer for its nearest sectioning content or sectioning root element.
func Footer(elements ...any) *gas.Tag { return gas.NewTag("footer", elements...) }

// <form> tag represents a document section containing interactive controls for submitting information to a web server.
func Form(elements ...any) *gas.Tag { return gas.NewTag("form", elements...) }

func HeadingGroup(elements ...any) *gas.Tag { return gas.NewTag("hgroup", elements...) }

// <h1> is the highest section level
func Heading1(elements ...any) *gas.Tag { return gas.NewTag("h1", elements...) }

// <h2> is the second highest section level
func Heading2(elements ...any) *gas.Tag { return gas.NewTag("h2", elements...) }

// <h3> is the third highest section level
func Heading3(elements ...any) *gas.Tag { return gas.NewTag("h3", elements...) }

// <h4> is the fourth highest section level
func Heading4(elements ...any) *gas.Tag { return gas.NewTag("h4", elements...) }

// <h5> is the fifth highest section level
func Heading5(elements ...any) *gas.Tag { return gas.NewTag("h5", elements...) }

// <h6> is the sixth highest section level
func Heading6(elements ...any) *gas.Tag { return gas.NewTag("h6", elements...) }

// <hr> tag defines a thematic break in an HTML page (e.g. a shift of topic).
func HorizontalRule(elements ...any) *gas.Tag { return gas.NewTag("hr", elements...) }

// <i> tag defines a part of text in an alternate voice or mood. The content inside is typically displayed in italic type.
func Italic(elements ...any) *gas.Tag { return gas.NewTag("i", elements...) }

// <iframe> tag defines an inline frame used to embed another document within the current HTML document.
func InlineFrame(elements ...any) *gas.Tag { return gas.NewTag("iframe", elements...) }

// <img> tag defines an image in an HTML page.
func Image(elements ...any) *gas.Tag { return gas.NewTag("img", elements...) }

// <input> tag specifies an input field where the user can enter data.
func Input(elements ...any) *gas.Tag { return gas.NewTag("input", elements...) }

// <ins> (short for insert) tag defines a range of text that has been added to a document.
func Inserted(elements ...any) *gas.Tag { return gas.NewTag("ins", elements...) }

// <kbd> tag defines keyboard text.
func Keyboard(elements ...any) *gas.Tag { return gas.NewTag("kbd", elements...) }

// <keygen> tag defines a key-pair generator field (for forms).
func KeyGen(elements ...any) *gas.Tag { return gas.NewTag("keygen", elements...) }

// <label> tag defines a label for an <input> element.
func Label(elements ...any) *gas.Tag { return gas.NewTag("label", elements...) }

// <legend> tag defines a caption for the <fieldset> element.
func Legend(elements ...any) *gas.Tag { return gas.NewTag("legend", elements...) }

// <li> tag defines a list item.
func ListItem(elements ...any) *gas.Tag { return gas.NewTag("li", elements...) }

// <main> tag specifies the main content of a document.
func Main(elements ...any) *gas.Tag { return gas.NewTag("main", elements...) }

// <map> tag is used with <area> elements to define an image map (a clickable link area).
func Map(elements ...any) *gas.Tag { return gas.NewTag("map", elements...) }

// <mark> tag defines marked/highlighted text.
func Marked(elements ...any) *gas.Tag { return gas.NewTag("mark", elements...) }

// <menu> tag defines a list/menu of commands.
func Menu(elements ...any) *gas.Tag { return gas.NewTag("menu", elements...) }

// <meta> tag provides metadata about the HTML document. Metadata will not be displayed on the page, but will be machine parsable.
func Meta(elements ...any) *gas.Tag { return gas.NewTag("meta", elements...) }

// <meter> tag defines a scalar measurement within a known range (a gauge).
func Meter(elements ...any) *gas.Tag { return gas.NewTag("meter", elements...) }

// <nav> tag defines a set of navigation links.
func Navigation(elements ...any) *gas.Tag { return gas.NewTag("nav", elements...) }

// <noscript> tag defines an alternate content for users that do not support client-side scripts.
func NoScript(elements ...any) *gas.Tag { return gas.NewTag("noscript", elements...) }

// <object> tag defines an embedded object within an HTML document.
func Object(elements ...any) *gas.Tag { return gas.NewTag("object", elements...) }

// <ol> tag defines an ordered list.
func OrderedList(elements ...any) *gas.Tag { return gas.NewTag("ol", elements...) }

// <optgroup> tag defines a group of related options in a drop-down list.
func OptionGroup(elements ...any) *gas.Tag { return gas.NewTag("optgroup", elements...) }

// <option> tag defines an option in a drop-down list.
func Option(elements ...any) *gas.Tag { return gas.NewTag("option", elements...) }

// <output> tag represents the result of a calculation.
func Output(elements ...any) *gas.Tag { return gas.NewTag("output", elements...) }

// <p> tag defines a paragraph.
func Paragraph(elements ...any) *gas.Tag { return gas.NewTag("p", elements...) }

// <param> tag defines a parameter for an object.
func Parameter(elements ...any) *gas.Tag { return gas.NewTag("param", elements...) }

// <picture> tag specifies multiple source for img element
func Picture(elements ...any) *gas.Tag { return gas.NewTag("picture", elements...) }

// <pre> tag defines preformatted text.
func Preformatted(elements ...any) *gas.Tag { return gas.NewTag("pre", elements...) }

// <progress> tag represents the progress of a task.
func Progress(elements ...any) *gas.Tag { return gas.NewTag("progress", elements...) }

// <q> tag defines a short quotation.
func Quotation(elements ...any) *gas.Tag { return gas.NewTag("q", elements...) }

// <rp> tag defines what to show in browsers that do not support ruby annotations.
func RubyParenthesis(elements ...any) *gas.Tag { return gas.NewTag("rp", elements...) }

// <rt> tag defines an explanation/pronunciation of characters (for East Asian typography).
func RubyText(elements ...any) *gas.Tag { return gas.NewTag("rt", elements...) }

// <ruby> tag defines a ruby annotation (for East Asian typography).
func Ruby(elements ...any) *gas.Tag { return gas.NewTag("ruby", elements...) }

// <samp> tag defines sample output from a computer program.
func Sample(elements ...any) *gas.Tag { return gas.NewTag("samp", elements...) }

// <script> tag is used to define a client-side script (JavaScript).
func Script(elements ...any) *gas.Tag { return gas.NewTag("script", elements...) }

// <section> tag defines sections in a document, such as chapters, headers, footers, or any other sections of the document.
func Section(elements ...any) *gas.Tag { return gas.NewTag("section", elements...) }

// <select> tag defines a drop-down list.
func Select(elements ...any) *gas.Tag { return gas.NewTag("select", elements...) }

// <small> tag defines smaller text.
func Small(elements ...any) *gas.Tag { return gas.NewTag("small", elements...) }

// <source> tag defines multiple media resources for media elements (<video> and <audio>).
func Source(elements ...any) *gas.Tag { return gas.NewTag("source", elements...) }

// <span> tag is used to group inline-elements in a document.
func Span(elements ...any) *gas.Tag { return gas.NewTag("span", elements...) }

// <strong> tag defines important text
func Strong(elements ...any) *gas.Tag { return gas.NewTag("strong", elements...) }

// <style> tag defines style information for a document
func Style(elements ...any) *gas.Tag { return gas.NewTag("style", elements...) }

// <sub> tag defines subscripted text
func Subscript(elements ...any) *gas.Tag { return gas.NewTag("sub", elements...) }

// <summary> tag defines a visible heading for a <details> element
func Summary(elements ...any) *gas.Tag { return gas.NewTag("summary", elements...) }

// <sup> tag defines superscripted text
func Superscript(elements ...any) *gas.Tag { return gas.NewTag("sup", elements...) }

// <table> tag defines an HTML table
func Table(elements ...any) *gas.Tag { return gas.NewTag("table", elements...) }

// <tbody> tag groups the body content in a table
func TableBody(elements ...any) *gas.Tag { return gas.NewTag("tbody", elements...) }

// <td> tag defines a standard cell in an HTML table
func TableData(elements ...any) *gas.Tag { return gas.NewTag("td", elements...) }

// <textarea> tag defines a multi-line text input control
func TextArea(elements ...any) *gas.Tag { return gas.NewTag("textarea", elements...) }

// <tfoot> tag groups the footer content in a table
func TableFooter(elements ...any) *gas.Tag { return gas.NewTag("tfoot", elements...) }

// <th> tag defines a header cell in an HTML table
func TableHead(elements ...any) *gas.Tag { return gas.NewTag("th", elements...) }

// <thead> tag groups the header content in a table
func TableHeader(elements ...any) *gas.Tag { return gas.NewTag("thead", elements...) }

// <time> tag defines a date/time
func Time(elements ...any) *gas.Tag { return gas.NewTag("time", elements...) }

// <title> tag defines a title for the document
func Title(elements ...any) *gas.Tag { return gas.NewTag("title", elements...) }

// <tr> tag defines a row in an HTML table
func TableRow(elements ...any) *gas.Tag { return gas.NewTag("tr", elements...) }

// <track> tag defines text tracks for media elements (<video> and <audio>)
func Track(elements ...any) *gas.Tag { return gas.NewTag("track", elements...) }

// <u> tag defines text that should be stylistically different from normal text
func Underline(elements ...any) *gas.Tag { return gas.NewTag("u", elements...) }

// <ul> tag defines an unordered list
func UnorderedList(elements ...any) *gas.Tag { return gas.NewTag("ul", elements...) }

// <var> tag defines a variable
func Variable(elements ...any) *gas.Tag { return gas.NewTag("var", elements...) }

// <video> tag defines a video or movie
func Video(elements ...any) *gas.Tag { return gas.NewTag("video", elements...) }

// <wbr> tag defines a possible line-break
func WordBreakOpportunity(elements ...any) *gas.Tag { return gas.NewTag("wbr", elements...) }
