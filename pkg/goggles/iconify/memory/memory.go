/* cSpell:disable */
package memory

import "github.com/gogoracer/racer/pkg/handlebars"

func Stop() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M16 5v1h1v10h-1v1H6v-1H5V6h1V5h10M7 7v8h8V7H7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 4h-2V2h-2V0h4v4m-6-2h-4V0h4v2m-6 0H6V0h4v2M4 2H1V0h3v2m18 4v4h-2V6h2m0 6v3h-2v-3h2m0 5v4h-2v-4h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CheckboxBlank() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 4h1V3h14v1h1v14h-1v1H4v-1H3V4m2 13h12V5H5v12Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func File() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M13 1v1h1v1h1v1h1v1h1v1h1v1h1v13h-1v1H4v-1H3V2h1V1h9m0 3h-1v4h4V7h-1V6h-1V5h-1V4M5 3v16h12v-9h-6V9h-1V3H5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Speaker() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 1h14v1h1v18h-1v1H4v-1H3V2h1V1m1 2v16h12V3H5m4 2h1V4h2v1h1v2h-1v1h-2V7H9V5m0 13v-1H8v-1H7v-4h1v-1h1v-1h4v1h1v1h1v4h-1v1h-1v1H9m1-5H9v2h1v1h2v-1h1v-2h-1v-1h-2v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaM() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h8v1h1v9h-2V8h-2v7h-2V8H8v8H6V7h1V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CheckboxMarked() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 4h1V3h14v1h-1v1H5v12h12v-6h1v-1h1v8h-1v1H4v-1H3V4m3 5h2v1h1v1h1v1h2v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6V9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Email() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 5h1V4h18v1h1v13h-1v1H2v-1H1V5m2 12h16V9h-1v1h-2v1h-2v1h-2v1h-2v-1H8v-1H6v-1H4V9H3v8M19 6H3v1h2v1h2v1h2v1h4V9h2V8h2V7h2V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Label() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 4h13v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-1H1V5h1V4m14 9h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H3v10h11v-1h1v-1h1v-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MessageProcessing() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m2 14h15V3H3v13h1v-1m2-7h2v2H6V8m4 0h2v2h-2V8m4 0h2v2h-2V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PlusCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-7 0h2v4h4v2h-4v4h-2v-4H6v-2h4V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Water() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 21H8v-1H6v-1H5v-1H4v-2H3v-3h1v-2h1V9h1V7h1V6h1V4h1V3h1V1h2v2h1v1h1v2h1v1h1v2h1v2h1v2h1v3h-1v2h-1v1h-1v1h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaA() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 3v1h1v9h-2v-4h-2v4H8V7h1V6h4m-1 2h-2v2h2V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaCFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1v-2h-2v1h-2V8h2v1h2V7h-1V6H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 12v-2h1V9h1V8h1V7h1V6h2v2h-1v1H9v1h9v2H9v1h1v1h1v2H9v-1H8v-1H7v-1H6v-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDownRightCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 9v4h-1v1h-1v1H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7h1V0h2v7h1v1h1v1h1m-3 4v-1h1v-2h-1V9h-2v1H9v2h1v1h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowTopRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 5h9v9h-2V9h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v-1H6v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7H8V5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightRoundUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 10h4v1h3v1h1v1h1v1h1v1h1v3h1v4h-2v-3H9v-3H8v-1H7v-1H6v-1H3v-1H0v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightRoundUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 0v2h-5v1h-4v1h-2v1H9v1H8v1H7v1H6v1H5v2H4v2H3v4H2v5H.004v-6H1v-4h1v-2h1V8h1V7h1V6h1V5h1V4h1V3h2V2h2V1h4V0h6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Clock() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 5h2v6h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-1V5m5-4v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedDownLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 18h2v2h2v2H0v-4m6 2h4v2H6v-2m6 0h4v2h-4v-2m6 0h2v-2h2v4h-4v-2M0 16v-4h2v4H0m0-6V7h2v3H0m0-5V1h2v4H0m22-4v4h-2V1h2m0 6v3h-2V7h2m0 5v4h-2v-4h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightRoundUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 22h-2v-5h-1v-4h-1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-2V4H9V3H5V2H.004V0H6v1h4v1h2v1h2v1h1v1h1v1h1v1h1v1h1v2h1v2h1v4h1v6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9m2-6v4h4V3H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaC() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v2h-2V8h-2v6h2v-1h2v2h-1v1H9v-1H8V7h1V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaMFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v1H6v9h2V8h2v7h2V8h2v8h2V7h-1V6H7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 20v-2h2v2h-2m0-16V2h2v2h-2m8 8v-2h2v2h-2m0 4v-2h2v2h-2m-4 4v-2h2v2h-2m4 0v-2h2v2h-2M6 20v-2h2v2H6M6 4V2h2v2H6m12 4V6h2v2h-2m0-4V2h2v2h-2m-4 0V2h2v2h-2M2 20V2h2v18H2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderTopRightBottom() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 12H2v-2h2v2m0 4H2v-2h2v2m0-8H2V6h2v2M2 4V2h18v18H2v-2h16V4H2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaLFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h6v-2h-4V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowTopLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 14V5h9v2H9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7v5H5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 5h2v1h1v1h1v1h1v1h1v2h-2v-1h-1V9h-1v9h-2V9H9v1H8v1H6V9h1V8h1V7h1V6h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pencil() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M16 2h1v1h1v1h1v1h1v1h-1v1h-1v1h-1V7h-1V6h-1V5h-1V4h1V3h1m-4 3h2v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H2v-4h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FlaskRoundBottom() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M11 11h2v2h-2v-2m2-10v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1V1h4m-1 4h-2v4H9v1H8v1H7v1H6v2h1v-1h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h-1v-1h-1v-1h-1V9h-1V5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Lightbulb() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 19h6v2H8v-2m0-1v-4H7v-1H6v-1H5v-1H4V5h1V4h1V3h1V2h1V1h6v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v4H8m5-6h1v-1h1v-1h1V6h-1V5h-1V4h-1V3H9v1H8v1H7v1H6v4h1v1h1v1h1v1h1v3h2v-3h1v-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TextImage() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 4h10v10H2V4m2 2v6h6V6H4m10-2h6v2h-6V4m0 4h6v2h-6V8m0 4h6v2h-6v-2M2 16h16v2H2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipEnd() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M21 20V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1h1m-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12v16Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CardText() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M17 8v2H5V8h12M5 12h10v2H5v-2M2 3h18v1h1v14h-1v1H2v-1H1V4h1V3m1 2v12h16V5H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Device() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 1h18v1h.94v18H20v1H2v-1h-.94V2H2V1m1 2v16h16V3H3m1 1h14v8H4V4m1 10h3v3H5v-3m7 1h2v2h-2v-2m3-1h2v2h-2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuUpFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 13v1h12v-1h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-2v1H9v1H8v1H7v1H6v1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Monitor() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 2h18v1h1v12h-1v1h-7v2h2v2H7v-2h2v-2H2v-1H1V3h1V2m1 2v10h16V4H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaEFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h6v-2h-4v-2h4v-2h-4V8h4V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaUFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v9h1v1h4v-1h1V6h-2v8h-2V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BagPersonal() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M17 15H9v2H7v-2H5v4h12v-4m0-6h-1V8h-1V7H7v1H6v1H5v4h12V9m-4 2H9v-1h1V9h2v1h1v1M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3V8m6-5v2h4V3H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 21v-4h2v4h-2m0-6v-3h2v3h-2m0-5V6h2v4h-2m0-6V1h2v3h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func VolumeHigh() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M13 2h2v1h1v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-1V4h-1V2m1 5v1h1v2h1v2h-1v2h-1v1h-1V7h1M2 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1H9v-1H8v-1H7v-1H6v-1H2V8m2 2v2h3v1h1v1h1V8H8v1H7v1H4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CompassSouthEast() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 6h5v2H6v2h3v1h1v4H9v1H4v-2h4v-2H5v-1H4V7h1m7-1h6v2h-4v2h4v2h-4v2h4v2h-6"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadCenter() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 7H9v1H8v4h1v1h4v-1h1V9h-1V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipAboveText() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3m2 3h12v2H5V6m0 4h10v2H5v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronUpCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M6 13v-2h1v-1h1V9h1V8h1V7h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6m9-12v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuUpDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 10V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v2H4m0 2h14v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-2m4-4h6V7h-1V6h-1V5h-2v1H9v1H8v1m0 6v1h1v1h1v1h2v-1h1v-1h1v-1H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MinusBox() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M16 12H6v-2h10Zm2 8H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 2v2h-2V2h2m0 16v2h-2v-2h2m-8-8v2H2v-2h2m0-4v2H2V6h2m4-4v2H6V2h2M4 2v2H2V2h2m12 0v2h-2V2h2m0 16v2h-2v-2h2M4 14v2H2v-2h2m0 4v2H2v-2h2m4 0v2H6v-2h2M20 2v18h-2V2h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedUpLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 4h-2V2h-2V0h4v4m-6-2h-4V0h4v2m-6 0H6V0h4v2M4 2H2v2H0V0h4v2m18 4v4h-2V6h2m0 6v3h-2v-3h2m0 5v4h-2v-4h2M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BugFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 7h3V6h1V5h1V4H7V3H6V2h1V1h1v1h1v1h1v1h2V3h1V2h1V1h1v1h1v1h-1v1h-1v1h1v1h1v1h3v2h-2v2h2v2h-2v2h2v2h-3v1h-1v1h-1v1H8v-1H7v-1H6v-1H3v-2h2v-2H3v-2h2V9H3V7m6 6v2h4v-2H9m0-4v2h4V9H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cart() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 14v2H6v-1H5v-4H4V8H3V3H1V1h4v3h16v4h-1v3h-1v1H7v2h12M5 7h1v3h12V7h1V6H5v1m2 10h2v1h1v2H9v1H7v-1H6v-2h1v-1m8 0h2v1h1v2h-1v1h-2v-1h-1v-2h1v-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightUpDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 0h22v2H0V0m0 20h22v2H0v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaG() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h5v2h-4v6h2v-2h-1v-2h3v5h-1v1H9v-1H8V7h1V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaL() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v8h4v2H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Battery100() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7m-6 1h2v6h-2V8m3 0h2v6h-2V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 22V10h12v2H12v10h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightVertical() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 0v22h-2V0h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CheckboxCross() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M13 12h1v1h1v2h-2v-1h-1v-1h-2v1H9v1H7v-2h1v-1h1v-2H8V9H7V7h2v1h1v1h2V8h1V7h2v2h-1v1h-1v2m5 7H4v-1H3V4h1V3h14v1h1v14h-1v1M5 5v12h12V5H5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CommentText() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1V2m1 2v12h5v3h1v-1h1v-1h1v-1h8V4H3m2 3h12v2H5V7m0 4h10v2H5v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuBottomLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M7 4h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2H7V4m2 4v5h5v-1h-1v-1h-1v-1h-1V9h-1V8H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaY() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h2v2h1v2h2V8h1V6h2v3h-1v2h-1v2h-1v3h-2v-3H9v-2H8V9H7V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Axe() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M11 3h2v1h1v1h1v1h2v1h2v1h1v1h1v2h-1v2h-1v1h-1v1h-2v1h-1v-1h-1v-1h-1v-2h-1v-1h-1v-1h-1V9H9V8H8V6h1V5h1V4h1m-1 6v1h1v2h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v-1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderBottom() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 12h-2v-2h2v2M4 12H2v-2h2v2m8-8h-2V2h2v2m4 0h-2V2h2v2m4 4h-2V6h2v2m0-4h-2V2h2v2m0 12h-2v-2h2v2M4 16H2v-2h2v2M8 4H6V2h2v2M4 4H2V2h2v2m0 4H2V6h2v2m16 12H2v-2h18v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderOutside() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 2h18v18H2V2m2 2v14h14V4H4m6 2h2v2h-2V6m0 4h2v2h-2v-2m-4 0h2v2H6v-2m8 0h2v2h-2v-2m-4 4h2v2h-2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MonitorImage() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 6h2v1h1v2h-1v1h-2V9h-1V7h1V6M2 2h18v1h1v12h-1v1h-7v2h2v2H7v-2h2v-2H2v-1H1V3h1V2m1 2v5h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h4V4H3m7 7H9v-1H8V9H6v1H5v1H4v1H3v2h9v-1h-1v-1h-1v-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipStartText() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 8V6h8v2H5m0 8v-2h6v2H5m0-4v-2h8v2H5m-4 8V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1H1m2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3v16Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlertCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 12h-2V6h2Zm0 4h-2v-2h2Zm3 5H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-1-2v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaP() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v4h-1v1h-3v4H8V6m2 2v2h2V8h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightRoundUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 22v-4h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h4v2h-3v1h-3v1h-1v1h-1v1h-1v3h-1v3h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 22v-2h2v-2h2v4h-4m2-6v-4h2v4h-2m0-6V6h2v4h-2m0-6V1h2v3h-2m-4 18h-4v-2h4v2m-6 0H7v-2h3v2m-5 0H1v-2h4v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Crown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 17h18v2H2v-2M4 6v1h1v1h1V7h1V6h1V5h1V4h1V3h2v1h1v1h1v1h1v1h1v1h1V7h1V6h1V5h1v11H2V5h1v1h1m3 8h11v-4h-3V9h-1V8h-1V7h-1V6h-2v1H9v1H8v1H7v1H4v4h3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Folder() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4h1V3m1 4v10h16V7H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Microphone() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 21v-3H8v-1H6v-1H5v-1H4v-2H3V9h2v3h1v2h1v1h2v1h4v-1h2v-1h1v-2h1V9h2v4h-1v2h-1v1h-1v1h-2v1h-2v3h-2m-2-8v-1H7V3h1V2h1V1h4v1h1v1h1v9h-1v1h-1v1H9v-1H8m1-2h1v1h2v-1h1V4h-1V3h-2v1H9v7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Necklace() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M9 17h1v-1h2v1h1v2h-1v1h-2v-1H9v-2m1-2v-1H9v-1h-.91v-1H7v-2H6V8H5V6H4V3h1V2h12v1h1v3h-1v2h-1v2h-1v2h-1v1h-1v1h-1v1h-2M7 5v2h1v2h1v2h1.09v1H12v-1h1V9h1V7h1V5h1V4H6v1h1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaHFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h2v4h2V6h-2v4h-2V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Archive() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 2h18v6h-1v12H3V8H2V2m15 16V8H5v10h12M4 4v2h14V4H4m3 6h8v2H7v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDownCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M16 10v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V6h2v6h1v-1h1v-1h2m5-3v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0m0-6V1h2v3H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Trash() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 7v9H8V7h2m2 0h2v9h-2V7M8 2h6v1h5v2h-1v14h-1v1H5v-1H4V5H3V3h5V2M6 5v13h10V5H6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipEndAlert() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 15h-2v-2h2v2m0-3h-2V7h2v5m7-10v18h-1v1H6v-1H5v-5H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V2h1V1h14v1h1m-2 1H7v5H6v1H5v1H4v2h1v1h1v1h1v5h12V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaFFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h3v-2h-3V8h4V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 4v14h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h1V5h1V4h2m-2 4h-1v1h-1v1h-1v2h1v1h1v1h1V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Message() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m1 2v13h1v-1h15V3H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MinusBoxFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-2-8v-2H6v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M6 12H5v1H4v2h2v-1h1v-1h1v-1h1v-1h1v-1h2v1h1v1h1v1h1v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-1V7h-2v1H9v1H8v1H7v1H6"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MessageText() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m1 2v13h1v-1h15V3H3m2 3h12v2H5V6m0 4h9v2H5v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func RotateClockwise() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 11v1h-1v1h-1v1h-1v1h-2v-1h-1v-1h-1v-1h-1v-1h3V9h-1V7h-1V6h-2V5H9v1H7v1H6v2H5v4h1v2h1v1h2v1h4v-1h3v2h-2v1H8v-1H6v-1H5v-1H4v-2H3V8h1V6h1V5h1V4h2V3h6v1h2v1h1v1h1v2h1v3h3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Sword() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 3v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-1h2v2h-1v1h-1v1h1v1h1v1h1v2h-1v1h-2v-1h-1v-1h-1v-1h-1v1h-1v1h-2v-2h1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V8H3V7H2V2h5v1h1M7 5H6V4H4v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h-1v-1h-1v-1h-1V9h-1V8H9V7H8V6H7V5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AccountBox() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m0 14h1v-1h2v-1h8v1h2v1h1V5h-1V4H5v1H4v11m12 2v-1h-2v-1H8v1H6v1h10M9 5h4v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V7h1V6h1V5m3 3V7h-2v1H9v2h1v1h2v-1h1V8h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaI() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 3v2h-1v6h1v2H9v-2h1V8H9V6h4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Arrow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 2h-3v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H3v1H2v1H1v1h1v1h1v1h1v1h1v-1h1v-1h1v-3h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h2"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowBottomLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 17H5V8h2v5h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1v1h1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1h5v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipStart() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 20V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1H1m2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3v16Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TextBox() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m13 3V4H5v1H4v12h1v1h12v-1h1V5h-1M6 8h10v2H6V8m0 4h7v2H6v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderTop() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2m-8 8h2v2h-2v-2m-4 0h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 6h2v2H2V6m16 0h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2M2 2h18v2H2V2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Box() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m13 3V4H5v1H4v12h1v1h12v-1h1V5h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Fire() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 20H7v-1H6v-1H5v-1H4v-5h1v-2h1V9h1V8h1v1h1v2h1V9h1V5h-1V4H9V3H8V2h3v1h2v1h1v1h1v1h1v1h1v2h1v7h-1v2h-1v1h-2m-2-1v-1h2v-1h1v-2h1v-4h-1V8h-1V7h-1v4h-1v2h-1v1h-1v1H9v-1H8v-3H7v1H6v4h1v1h1v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuDownFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M17 9V8H5v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1V9"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 0h2v20h20v2H0V0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 22v-2h20V0h2v22H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cancel() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-2h1V8h-1V6h-1m-3 10v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7H5V6H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h-1v-1h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Image() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 4h1V3h18v1h1v14h-1v1H2v-1H1V4m2 10h1v-1h1v-1h1v-1h1v-1h1V9h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2V5H3v9m11 3v-1h-1v-1h-1v-1h-1v-1h-1v-1H8v1H7v1H6v1H5v1H4v1h10m-1-9h1V7h2v1h1v2h-1v1h-2v-1h-1V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaZFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v2h4v1h-1v1h-1v1H9v1H8v4h6v-2h-4v-1h1v-1h1v-1h1v-1h1V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowBottomRightCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1m1-4h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1m-1-9v7H8v-2h3v-1h-1v-1H9v-1H8V9H7V8h1V7h1v1h1v1h1v1h1v1h1V8h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowTopLeftCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M7 1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1M6 5H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1m1 9V7h7v2h-3v1h1v1h1v1h1v1h1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v-1H9v3H7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderTopRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 18v2h-2v-2h2m-8-8v2H2v-2h2m0-4v2H2V6h2m12 12v2h-2v-2h2M4 14v2H2v-2h2m0 4v2H2v-2h2m4 0v2H6v-2h2M20 2v18h-2V4H2V2h18Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PlusBox() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 16h-2v-4H6v-2h4V6h2v4h4v2h-4Zm6 4H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Terminal() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 3v16h-1v1H3v-1H2V3h1V2h16v1h1m-2 3H4v12h14V6M9 9v1h1v1h1v2h-1v1H9v1H8v1H6v-2h1v-1h1v-2H7v-1H6V8h2v1h1m2 7v-2h5v2h-5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func VolumeLow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M6 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H6V8m2 2v2h3v1h1v1h1V8h-1v1h-1v1H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedAll() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v2h2v2H0v-4h2M6 0h4v2H6V0m6 0h4v2h-4V0m6 0h4v4h-2V2h-2V0m0 22v-2h2v-2h2v4h-4m-2 0h-4v-2h4v2m-6 0H6v-2h4v2M20 6h2v4h-2V6m0 6h2v4h-2v-4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 0h22v2H0V0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightUpLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 0h22v22h-2V2H2v20H0V0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Coffee() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 20v-2h16v2H1M2 3h17v1h1v6h-1v1h-3v3h-1v1h-1v1H4v-1H3v-1H2V3m14 2v4h2V5h-2M4 5v8h1v1h8v-1h1V5H4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaS() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V7h1V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ApplicationCode() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M11 16H9v-1H8v-4h1v-1h2v2h-1v2h1m4 2h-2v-2h1v-2h-1v-2h2v1h1v4h-1m4 5H3v-1H2V3h1V2h16v1h1v16h-1M18 6V4H4v2m14 12V8H4v10Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 12H10V0h2v10h10v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightHorizontal() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 10h22v2H0v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Target() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 13h-2v-1H9v-2h1V9h2v1h1v2h-1Zm2 4H8v-1H7v-1H6v-1H5V8h1V7h1V6h1V5h6v1h1v1h1v1h1v6h-1v1h-1v1h-1Zm1 4H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-2-6v-1h1v-1h1V9h-1V8h-1V7H9v1H8v1H7v4h1v1h1v1Zm1 4v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipEndText() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M17 8V6H9v2h8m-2 8v-2H9v2h6m2-4v-2H9v2h8m4 8V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1h1m-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12v16Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Multiply() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M16 16h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6v-2h1v-1h1v-1h1v-2H8V9H7V8H6V6h2v1h1v1h1v1h2V8h1V7h1V6h2v2h-1v1h-1v1h-1v2h1v1h1v1h1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 22V0h2v22H0m20 0V0h2v22h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Comment() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1V2m1 2v12h5v3h1v-1h1v-1h1v-1h8V4H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadEmpty() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 2H9v6H3v4h6v6h4v-6h6V9h-6V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MinusCircleFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2m1-8v-2H6v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowRightDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 19h2v-1h1v-1h1v-1h1v-1h1v-2h-2v1h-1v1h-1v-4h-1V9h-1V8h-2V7H3v2h6v1h2v2h1v3h-1v-1h-1v-1H8v2h1v1h1v1h1v1h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 22h-2V2H.002V0H22v22Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 22V0h2v22H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaR() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v4h-1v2h1v3h-2v-2h-1v-2h-1v4H8V6m2 2v2h2V8h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDownBold() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 10v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h5V2h8v8h5m-4 2h-3V4H9v8H6v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDoubleRoundUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 22v-3h-1v-2h-1v-2H9v-1H8v-1H7v-1H5v-1H3v-1H0V8h4v1h2v1h2v1h1v1h1v1h1v1h1v2h1v2h1v4h-2M0 12h2v1h3v1h1v1h1v1h1v1h1v3h1v2H8v-1H7v-3H6v-1H5v-1H4v-1H1v-1H0v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 10h12v12h-2V12H0v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaT() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-2v8h-2V8H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowLeftCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 16h-2v-1H9v-1H8v-1H7v-1H6v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h6v2h-6v1h1v1h1v2m3 5H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1m1-4h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipBelowAlert() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 17h2v-2h-2v2m0-3h2V9h-2v5m-8 7h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Logout() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M17 1v1h1v4h-1V5h-1V3H6v16h10v-2h1v-1h1v4h-1v1H5v-1H4V2h1V1h12m-4 5h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H8v-2h7V9h-1V8h-1V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Octagon() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1V6m13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1v-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlertBox() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 12h-2V6h2m0 10h-2v-2h2m6 6H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlertRhombus() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 12h-2V6h2Zm0 4h-2v-2h2Zm0 5h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm0-3v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-2v1H9v1H8v1H7v1H6v1H5v1H4v2h1v1h1v1h1v1h1v1h1v1h1v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Database() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M7 2h8v1h2v1h1v1h1v12h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3V5h1V4h1V3h2V2m1 14v-1H6v-1H5v2h1v1h2v1h6v-1h2v-1h1v-2h-1v1h-2v1H8m0-5v-1H6V9H5v3h2v1h2v1h4v-1h2v-1h2V9h-1v1h-2v1H8m1-3v1h4V8h2V7h2V6h-1V5h-2V4H8v1H6v1H5v1h2v1h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Diamond() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M6 2h10v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2V6h1V5h1V4h1V3h1V2m9 3V4h-1v2h1v1h2V6h-1V5h-1m-3 1V4h-2v2H9v1h4V6h-1M8 6V4H7v1H6v1H5v1h2V6h1m-4 5h1v1h1v1h1v1h1v-2H7V9H4v2m6 1v4h2v-4h1V9H9v3h1m4 0v2h1v-1h1v-1h1v-1h1V9h-3v3h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaZ() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v4h-1v1h-1v1h-1v1h-1v1h4v2H8v-4h1v-1h1v-1h1V9h1V8H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CalendarMonth() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1v1M4 4v2h14V4H4m0 4v10h14V8H4m10 6h2v2h-2v-2m-4 0h2v2h-2v-2m-4 0h2v2H6v-2m0-4h2v2H6v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pause() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M6 4h3v14H6V4m7 14V4h3v14h-3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Battery0() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 5h15v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5m1 2v8h13V7H4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedUpDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M21 2h-4V0h4v2m-6 0h-3V0h3v2m-5 0H6V0h4v2M4 2H1V0h3v2m17 20h-4v-2h4v2m-6 0h-3v-2h3v2m-5 0H6v-2h4v2m-6 0H1v-2h3v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CreditCard() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 4h18v1h1v12h-1v1H2v-1H1V5h1V4m1 2v2h16V6H3m0 10h16v-5H3v5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Map() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 4h2V3h2V2h4v1h3v1h2V3h2V2h3v16h-2v1h-2v1h-4v-1H9v-1H7v1H5v1H2V4m2 2v11h2v-1h1V5H5v1H4m8-1H9v11h1v1h3V6h-1V5m4 1h-1v11h2v-1h1V5h-2v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CompassSouthArrow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 6h2V5h1V4h1V3h1V2h1V0h-2v1h-1v1h-1v1h-2V2H9V1H8V0H6v2h1v1h1v1h1v1h1M9 8h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V9h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Heart() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2H1V5h1V4h1V3h1V2h5v1h1v1h2V3h1V2h5v1h1v1h1v1h1v5h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1m-7-9v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V6h-1V5h-1V4h-3v1h-1v1h-1v1h-2V6H9V5H8V4H5v1H4v1H3v3h1v2h1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Radiobox() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2V2m1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 12v2h1v1h1v1h1v1h1v1h2v-2H8v-1H7v-1h4v-1h2v-1h1v-2h1V3h-2v6h-1v2h-2v1H7v-1h1v-1h1V8H7v1H6v1H5v1H4v1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderBottomLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 4V2h2v2h-2m8 8v-2h2v2h-2m0 4v-2h2v2h-2M6 4V2h2v2H6m12 4V6h2v2h-2m0-4V2h2v2h-2m-4 0V2h2v2h-2M2 20V2h2v16h16v2H2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderTopLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 18v2h-2v-2h2m4 0v2h-2v-2h2m-8 0v2H6v-2h2m-4 2H2V2h18v18h-2V4H4v16Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChartBar() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 2h2v16h16v2H2V2m4 14V8h4v8H6m5 0V4h4v12h-4m5 0v-5h4v5h-4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 8h14v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4V8m4 2v1h1v1h1v1h2v-1h1v-1h1v-1H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipAboveAlert() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 5h2v5h-2V5m0 6h2v2h-2v-2M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaOFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1V7h-1V6H9m1 2h2v6h-2V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaQ() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v8h1v1h-1v1h-1v-1h-1v-1h-1v1H9v-1H8V7h1V6m1 2v6h1v-1h1V8h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BattleAxe() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1h-4v1h-1v1H9v1H8v4h4v1h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1H2v1H1v1h1v1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v4h4v-1h1v-1h1v-1h1V7h-4V5h-2"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderTopLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 10h2v2h-2v-2m-8 8h2v2h-2v-2m-4 0h2v2H6v-2M18 6h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2M2 2h18v2H4v16H2V2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUpBold() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 12v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-5v8H7v-8H2m4-2h3v8h4v-8h3V9h-1V8h-1V7h-1V6h-1V5h-2v1H9v1H8v1H7v1H6v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDashedDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 12h-2v-2h2v2m-4 0h-3v-2h3v2m-5 0h-3V9h2v1h1v2M12 0v2h-2V0h2m0 4v3h-2V4h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 21v-4h2v4h-2m0-6v-3h2v3h-2m0-5V6h2v4h-2m0-6V1h2v3h-2M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0m0-6V1h2v3H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v3H0v-3h2M6 0h4v2H6V0m6 0h3v2h-3V0m5 0h4v2h-4V0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaBFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h5v-1h1v-3h-1v-2h1V7h-1V6H8m2 2h2v2h-2V8m2 4v2h-2v-2h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaIFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 5H9v2h1v6H9v2h4v-2h-1V8h1V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaQFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h2v-1h1v1h1v1h1v-1h1v-1h-1V7h-1V6H9m1 2h2v5h-1v1h-1V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaW() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M6 6h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1H8v-1H7v-2H6V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Plus() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 17h-2v-5H5v-2h5V5h2v5h5v2h-5Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func RemoveCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-3 1v1h1v1h-1v1h-1v2h1v1h1v1h-1v1h-1v-1h-1v-1h-2v1H9v1H8v-1H7v-1h1v-1h1v-2H8V9H7V8h1V7h1v1h1v1h2V8h1V7h1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M17 10v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H4v-2h9V9h-1V8h-1V6h2v1h1v1h1v1h1v1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowRightBold() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-5H2V7h8V2m2 4v3H4v4h8v3h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowRightUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v4h-1v2h-1v1h-2v1H3v-2h6v-1h2v-2h1V7h-1v1h-1v1H8V7h1V6h1V5h1V4h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUpDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v8h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V7H9v1H8v1H6V7h1V6h1V5h1V4h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 14H4v-2h1v-1h1v-1h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v1h1v1h1v2m-4-2v-1h-1v-1h-1V9h-2v1H9v1H8v1h6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowLeftBold() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v5h8v8h-8v5m-2-4v-3h8V9h-8V6H9v1H8v1H7v1H6v1H5v2h1v1h1v1h1v1h1v1h1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Battery75() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7m-6 1h2v6h-2V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Book() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 2h1V1h14v1h1v18h-1v1H4v-1H3V2m8 7h-1V8H9v1H8v1H7V3H5v16h12V3h-5v7h-1V9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDoubleRoundDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 12h3v-1h2v-1h2V9h1V8h1V7h1V5h1V3h1V0h2v4h-1v2h-1v2h-1v1h-1v1H9v1H8v1H6v1H4v1H0v-2M10 0v2H9v3H8v1H7v1H6v1H5v1H2v1H0V8h1V7h3V6h1V5h1V4h1V1h1V0h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Check() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 11h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CompassEastArrow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M6 10v2H5v1H4v1H3v1H2v1H0v-2h1v-1h1v-1h1v-2H2V9H1V8H0V6h2v1h1v1h1v1h1v1m5-4h6v2h-4v2h4v2h-4v2h4v2h-6"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaO() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v8h-1v1H9v-1H8V7h1V6m1 2v6h2V8h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BagPersonalFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3V8m6-5v2h4V3H9m8 11H5v1h2v2h1v-2h9v-1m-5-3h1V9h-1V8h-2v1H9v2h1v1h2v-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderTopLeftBottom() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 10h2v2h-2v-2m0-4h2v2h-2V6m0 8h2v2h-2v-2m2 4v2H2V2h18v2H4v14h16Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDownLeftCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M13 15H9v-1H8v-1H7V9h1V8h1V7h1V0h2v7h1v1h1v1h1v1h7v2h-7v1h-1v1h-1v1m-4-3h1v1h2v-1h1v-2h-1V9h-2v1H9v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 18h2v2h2v2H0v-4m6 2h4v2H6v-2m6 0h4v2h-4v-2m6 0h3v2h-3v-2M0 16v-4h2v4H0m0-6V7h2v3H0m0-5V1h2v4H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDownLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 22H0V0h2v20h18V0h2v22Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func DoorBox() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M13 14h-2v-2h2Zm3 4h1v-1h1V5h-1V4H5v1H4v12h1v1h1V6h10Zm2 2H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-4-2V8H8v10Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipAbove() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlertBoxFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-6-8V6h-2v6Zm0 4v-2h-2v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 10v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1H3v-2h1V9h1V8h1V7h1V6h2v2H8v1H7v1h8V9h-1V8h-1V6h2v1h1v1h1v1h1v1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Battery50() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Clipboard() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 5h1V4h4V2h2V1h4v1h2v2h4v1h1v15h-1v1H3v-1H2V5m8-2v2h2V3h-2m8 3h-2v2H6V6H4v13h14V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 20h22v2H0v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Download() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 17v2H4v-2h14M14 2v6h4v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V8h4V2h6m-2 2h-2v6H9v1h1v1h2v-1h1v-1h-1V4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Led() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 21v-6H4v-2h2V6h1V4h1V3h1V2h4v1h1v1h1v2h1v7h2v2h-4v6h-2v-6h-2v6H8m4-16V4h-2v1H9v2H8v6h1v-1h4v1h1V7h-1V5h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaRFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h1v2h1v2h2v-3h-1v-2h1V7h-1V6H8m2 2h2v2h-2V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDoubleVertical() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 22V0h2v22h-2m-4 0V0h2v22H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LockOpen() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 13h2v1h1v2h-1v1h-2v-1H9v-2h1v-1m4-11v1h1v1h1v5h1v1h1v10h-1v1H5v-1H4V10h1V9h9V5h-1V4H9v1H8v2H6V4h1V3h1V2h6m2 9H6v8h10v-8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MusicNote() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M11 2h7v5h-5v11h-1v1h-1v1H7v-1H6v-1H5v-4h1v-1h1v-1h4V2m0 13h-1v-1H8v1H7v2h1v1h2v-1h1v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PlusBoxFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-6-4v-4h4v-2h-4V6h-2v4H6v2h4v4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaVFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v4h1v2h1v2h1v2h2v-2h1v-2h1v-2h1V6h-2v3h-1v2h-2V9H9V6H7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightUpRightCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M9 7h4v1h1v1h1v4h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7m4 3h-1V9h-2v1H9v2h1v1h2v-1h1v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 22V0h2v22h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FlaskEmpty() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 1h6v1h1v3h-1v2h1v2h1v2h1v2h1v1h1v2h1v4h-1v1H3v-1H2v-4h1v-2h1v-1h1v-2h1V9h1V7h1V5H7V2h1V1m2 2v5H9v2H8v2H7v1H6v2H5v2H4v2h14v-2h-1v-2h-1v-1h-1v-2h-1v-2h-1V8h-1V3h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaF() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-4v2h3v2h-3v4H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 6V5H9V4H7v2h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1H9v1H8v1H7v2h2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CubeUnfolded() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 3v5h10v7h-5v5h-7v-5H0V8h5V3h7m-2 2H7v3h3V5m-3 5v3h3v-3H7m-2 0H2v3h3v-3m12 0v3h3v-3h-3m-2 5h-3v3h3v-3m-3-5v3h3v-3h-3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Tag() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 2h1V1h9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V2m2 8h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H3v7m2-6h2v1h1v2H7v1H5V7H4V5h1V4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightUpDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 22V0h22v2H2v18h20v2H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Bug() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 7h3V6h1V5h1V4H7V3H6V2h1V1h1v1h1v1h1v1h2V3h1V2h1V1h1v1h1v1h-1v1h-1v1h1v1h1v1h3v2h-2v2h2v2h-2v2h2v2h-3v1h-1v1h-1v1H8v-1H7v-1H6v-1H3v-2h2v-2H3v-2h2V9H3V7m10 11v-1h1v-1h1V8h-1V7h-1V6H9v1H8v1H7v8h1v1h1v1h4m-4-5h4v2H9v-2m0-4h4v2H9V9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlertCircleFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-3-9V6h-2v6Zm0 4v-2h-2v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaNFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h1v2h1v2h2V6h-2v4h-1V8h-1V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaV() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h2v3h1v2h2V9h1V6h2v4h-1v2h-1v2h-1v2h-2v-2H9v-2H8v-2H7V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 17h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V4h2v9h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDoubleRoundUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 10h-3v1h-2v1h-2v1h-1v1h-1v1h-1v2h-1v2h-1v3H8v-4h1v-2h1v-2h1v-1h1v-1h1v-1h1v-1h2V9h2V8h4v2M12 22v-2h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h2v2h-1v1h-3v1h-1v1h-1v1h-1v3h-1v1h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 0h4v2H1V0m6 0h3v2H7V0m5 0h4v2h-4V0m6 0h3v2h-3V0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Briefcase() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1V6m7 0h4V4H9v2m10 2H3v10h16V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CompassNorthArrow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 4h2v1.5h1V8h1V4h2v10h-2v-2h-1v-2h-1v4H8m2 3v-1h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6v-2h1v-1h1v-1h1v-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Alert() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 20H2v-1H1v-4h1v-2h1v-2h1V9h1V7h1V5h1V3h1V2h6v1h1v2h1v2h1v2h1v2h1v2h1v2h1v4h-1v1M9 6H8v2H7v2H6v2H5v2H4v2H3v2h16v-2h-1v-2h-1v-2h-1v-2h-1V8h-1V6h-1V4H9v2m1 1h2v6h-2V7m0 7h2v2h-2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaD() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v1h1v6h-1v1h-1v1H8V6m2 2v6h2v-1h1V9h-1V8h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaWFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M6 6v7h1v2h1v1h2v-1h2v1h2v-1h1v-2h1V6h-2v6h-1v1h-1V8h-2v5H9v-1H8V6H6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowBottomLeftCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 15V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1m4 1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1m9-1H7V8h2v3h1v-1h1V9h1V8h1V7h1v1h1v1h-1v1h-1v1h-1v1h-1v1h3v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuLeftFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M13 17h1V5h-1v1h-1v1h-1v1h-1v1H9v1H8v2h1v1h1v1h1v1h1v1h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 4h2v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2V4m-2 0v14H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4h2m4 4v6h1v-1h1v-1h1v-2h-1V9h-1V8h-1M8 8H7v1H6v1H5v2h1v1h1v1h1V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Upload() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 17v2H4v-2h14M8 15V9H4V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h-4v6H8m2-2h2V7h1V6h-1V5h-2v1H9v1h1v6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Well() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 21H3v-5H1v-2h20v2h-2v5M5 16v3h12v-3H5M2 7V5h1V4h1V3h1V2h1V1h10v1h1v1h1v1h1v1h1v2h-2v6h-2V7h-4v2h2v4H8V9h2V7H6v6H4V7H2m5-4v1H6v1h10V4h-1V3H7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuRightFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M9 5H8v12h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H9"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlertRhombusFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 21h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm0-9V6h-2v6Zm0 4v-2h-2v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 10V8h1V7h1V6h1V5h1V4h2v2H8v1H7v1h4v1h2v1h1v2h1v7h-2v-6h-1v-2h-2v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDashedDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 0v2h-2V0h2m0 4v3h-2V4h2m0 5v3H9v-2h1V9h2M.002 10H2v2H.002v-2M4 10h3v2H4v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 16v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1v-1h-1v-2h1V9h1V8h1V7h1V6h1V4h-2v1h-1v1h-1v1h-1v1H9v1H8v1H7v2h1v1h1v1h1v1h1v1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Circle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LabelVariant() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 4v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-2H6V9H5V8H4V7H3V6H2V4h13m-1 12v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H6v1h1v1h1v1h1v1h1v2H9v1H8v1H7v1H6v1h8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CompassSouthWest() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 6h5v2H4v2h3v1h1v4H7v1H2v-2h4v-2H3v-1H2V7h1m7-1h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderOpen() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 4h1V3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4m2 5h16V7H9V6H8V5H3v4m0 8h16v-6H3v6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuTopLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 7v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7V7h11m-4 2H9v5h1v-1h1v-1h1v-1h1v-1h1V9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipBelowText() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3m2-3h10v-2H5v2m0-4h12v-2H5v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderBottomLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 4V2h2v2h-2M6 4V2h2v2H6m8 0V2h2v2h-2m4-2h2v18H2V2h2v16h14V2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDoubleHorizontal() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 12h22v2H0v-2m0-4h22v2H0V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Chat() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M7 4h9v1h2v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-2v1H8v1H4v1H1v-2h2v-1h1v-2H3v-1H2V8h1V7h1V6h1V5h2V4m10 4V7h-2V6H8v1H6v1H5v1H4v4h1v1h1v1h2v1h7v-1h2v-1h1v-1h1V9h-1V8h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M16 10h1V9h1V7h-2v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8V9H7V8H6V7H4v2h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDoubleRoundDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 0v3h1v2h1v2h1v1h1v1h1v1h2v1h2v1h3v2h-4v-1h-2v-1h-2v-1h-1v-1h-1V9h-1V8h-1V6H9V4H8V0h2m12 10h-2V9h-3V8h-1V7h-1V6h-1V5h-1V2h-1V0h2v1h1v3h1v1h1v1h1v1h3v1h1v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedUpDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v2h2v2H0v-4h2M6 0h4v2H6V0m6 0h3v2h-3V0m5 0h4v2h-4V0m4 22h-4v-2h4v2m-6 0h-3v-2h3v2m-5 0H6v-2h4v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaAFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 8h2v2h-2V8m5-7v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 5H9v1H8v9h2v-4h2v4h2V7h-1V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaJ() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-2 3h2v9h-1v1H9v-1H8v-2h2v1h2V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaN() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v2h1v2h1V6h2v10h-2v-2h-1v-2h-1v4H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderBottomRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 12H2v-2h2v2m8-8h-2V2h2v2m4 0h-2V2h2v2M4 16H2v-2h2v2M8 4H6V2h2v2M4 4H2V2h2v2m0 4H2V6h2v2m16 12H2v-2h16V2h2v18Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Play() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 5v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7V4h2v1h1m2 5h-1V9h-1V8H9v6h1v-1h1v-1h1v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Toolbox() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1V6m7 0h4V4H9v2m10 2H3v4h3v-2h3v2h4v-2h3v2h3V8M3 18h16v-4h-3v2h-3v-2H9v2H6v-2H3v4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 0v12H0v-2h10V0h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedUpDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 22v-2h2v-2h2v4h-4m2-6v-4h2v4h-2m0-6V6h2v4h-2m0-6V2h-2V0h4v4h-2m-4 18h-4v-2h4v2m-6 0H7v-2h3v2m-5 0H1v-2h4v2M1 0h4v2H1V0m6 0h3v2H7V0m5 0h4v2h-4V0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 14H9v4h4v-4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9M9 15v4h4v-4H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Bow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 3h10v1h2v1h2v1h1v1h1v2h1v2h1v10h-2v-2h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6V9H5V8H4V7H3V5H1m15 13h1v-6h-1v-2h-1V8h-1V7h-2V6h-2V5H4v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDashedUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 22v-2h2v2h-2m0-4v-3h2v3h-2m0-5v-3h3v2h-1v1h-2m12-1h-2v-2h2v2m-4 0h-3v-2h3v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronDownCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M16 9v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6V9h2v1h1v1h1v1h2v-1h1v-1h1V9h2m-1-8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func VolumeMedium() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M16 7v1h1v2h1v2h-1v2h-1v1h-1V7h1M8 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1H9v-1H8v-1H4V8h4m-2 2v2h3v1h1v1h1V8h-1v1H9v1H6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaGFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1v-5h-3v2h1v2h-2V8h4V6H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaSFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v4h1v1h3v2H8v2h5v-1h1v-4h-1v-1h-3V8h4V6H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderInside() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 10h8V2h2v8h8v2h-8v8h-2v-8H2v-2m4 8h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 2h2v2H2V2m0 4h2v2H2V6m4-4h2v2H6V2m8 0h2v2h-2V2m4 0h2v2h-2V2m0 4h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderNone() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2m-8-8h2v2h-2V2m0 16h2v2h-2v-2m-4 0h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 2h2v2H2V2m0 4h2v2H2V6m4-4h2v2H6V2m8 0h2v2h-2V2m4 0h2v2h-2V2m0 4h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func WallFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 3h6v4H5m8-4h5v4h-5m-8 8v4h6v-4m2 0h5v4h-5M8 9H3v4h5m2-4h10v4H10"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipStartAlert() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 15h2v-2H8v2m0-3h2V7H8v5M1 2v18h1v1h14v-1h1v-5h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V2h-1V1H2v1H1m2 1h12v5h1v1h1v1h1v2h-1v1h-1v1h-1v5H3V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 0v2H2v20H0V0h22Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Login() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 1h12v1h1v18h-1v1H5v-1H4v-6h2v5h10V3H6v5H4V2h1V1m3 5h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1H8v-2h1v-1h1v-1H2v-2h8V9H9V8H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pictogrammers() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 0h14v1h1v1h1v18h-1v1h-1v1H4v-1H3v-1H2V2h1V1h1V0m0 18v1h1v1h12v-1h1v-1H4M17 2H5v1H4v12h1v1h12v-1h1V3h-1V2m-4 2v1h1v1h1v2h-1v1h-1v1h-3v4H8V4h5m-1 2h-2v2h2V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M7 18V4h2v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7m2-4h1v-1h1v-1h1v-2h-1V9h-1V8H9v6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaDFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h5v-1h1v-1h1V8h-1V7h-1V6H8m2 2h2v1h1v4h-1v1h-2V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Battery25() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M7 8v6H5V8h2m11-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightDashedDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 20h4v2H1v-2m6 0h3v2H7v-2m5 0h4v2h-4v-2m6 0h3v2h-3v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9M9 3v4h4V3H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Minus() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M17 12H5v-2h12Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TagText() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 2h1V1h9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V2m2 8h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H3v7m11 1h1v2h-1v-1h-1v-1h-1v-1h-1V9h-1V7h1v1h1v1h1v1h1v1m-4 1h1v1h1v2h-1v-1h-1v-1H9v-1H8v-2h1v1h1v1M5 4h2v1h1v2H7v1H5V7H4V5h1V4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaE() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-4v2h4v2h-4v2h4v2H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Compass() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-4 6V8h2V7h2V6h2v2h-1v2h-1v2h-1v1h-1v1h-2v1H8v1H6v-2h1v-2h1v-2h1V9h1m2 1h-2v2h2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Flask() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M11 12h2v2h-2v-2m3-11v1h1v3h-1v2h1v2h1v2h1v2h1v1h1v2h1v4h-1v1H3v-1H2v-4h1v-2h1v-1h1v-2h1V9h1V7h1V5H7V2h1V1h6m-2 2h-2v5H9v2H8v2H7v1H6v2H5v2h1v-1h1v-1h1v-1h1v1h1v1h1v1h1v1h2v-1h1v-2h1v-1h-1v-2h-1v-2h-1V8h-1V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Lock() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 12h2v1h1v2h-1v1h-2v-1H9v-2h1v-1M8 2h6v1h1v1h1v4h1v1h1v10h-1v1H5v-1H4V9h1V8h1V4h1V3h1V2m1 2v1H8v3h6V5h-1V4H9m7 6H6v8h10v-8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightRoundDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 0v4h-1v3h-1v1H9v1H8v1H7v1H4v1H0v-2h3V9h3V8h1V7h1V6h1V3h1V0h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Calculator() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M18 21H4v-1H3V3h1V2h14v1h1v17h-1ZM16 7V4H6v3Zm-8 4V9H6v2Zm4 0V9h-2v2Zm4 0V9h-2v2Zm-8 4v-2H6v2Zm4 0v-2h-2v2Zm4 0v-2h-2v2Zm-8 4v-2H6v2Zm4 0v-2h-2v2Zm4 0v-2h-2v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Calendar() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1v1M4 4v2h14V4H4m0 4v10h14V8H4m8 4h4v4h-4v-4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CompassNorthEast() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 6h2v1.5h1V10h1V6h2v10H8v-2H7v-2H6v4H4m8-10h6v2h-4v2h4v2h-4v2h4v2h-6"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaK() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v2h1V7h1V6h2v2h-1v1h-1v1h-1v1h1v1h1v1h1v3h-2v-2h-1v-1h-1v3H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowLeftDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 19H8v-1H7v-1H6v-1H5v-1H4v-2h2v1h1v1h1v-4h1V9h1V8h2V7h7v2h-6v1h-2v2h-1v3h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowLeftUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 3H8v1H7v1H6v1H5v1H4v2h2V8h1V7h1v4h1v2h1v1h2v1h7v-2h-6v-1h-2v-2h-1V7h1v1h1v1h2V7h-1V6h-1V5h-1V4h-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderTopBottom() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2M2 14h2v2H2v-2m0-8h2v2H2V6m16 0h2v2h-2V6m0 8h2v2h-2v-2M2 18h18v2H2v-2M2 4V2h18v2H2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FloppyDisk() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 3h1V2h13v1h1v1h1v1h1v1h1v13h-1v1H3v-1H2V3m16 4h-1V6h-1V5h-1v4H6V4H4v14h2v-5h10v5h2V7m-7-3v3h2V4h-2m3 14v-3H8v3h6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TooltipBelow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaPFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h3v-1h1V7h-1V6H8m2 2h2v2h-2V8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaYFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v3h1v2h1v2h1v3h2v-3h1v-2h1V9h1V6h-2v2h-1v2h-2V8H9V6H7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Application() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h16v1h1v16h-1ZM18 6V4H4v2Zm0 12V8H4v10Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightAll() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 0h22v22H0V0m2 2v18h18V2H2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaH() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v4h2V6h2v10h-2v-4h-2v4H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowRightCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M7 1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1M6 5H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1m4 1h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H6v-2h6V9h-1V8h-1V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightDashedUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 10h2v2H0v-2m4 0h3v2H4v-2m5 0h3v3h-2v-1H9v-2m1 12v-2h2v2h-2m0-4v-3h2v3h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Skull() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M6 2h2V1h6v1h2v1h1v1h1v1h1v2h1v7h-1v2h-1v4h-1v1H5v-1H4v-4H3v-2H2V8h1V5h1V4h1V3h1V2m9 3V4h-2V3H9v1H7v1H6v1H5v3H4v4h1v2h1v4h2v-2h2v2h2v-2h2v2h2v-4h1v-2h1V8h-1V6h-1V5h-1M7 8h3v3H7V8m5 3V8h3v3h-3m-2 2h2v2h-2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func OctagonAlert() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1V6m13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1v-1M10 6h2v7h-2V6m0 8h2v2h-2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func RadioboxMarked() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2V2m1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4H9m4 3v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V9h1V8h1V7h4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightRoundDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 12h-4v-1h-3v-1h-1V9h-1V8h-1V7h-1V4h-1V0h2v3h1v3h1v1h1v1h1v1h3v1h3v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxLightUpLeftCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M7 13V9h1V8h1V7h4v1h1v1h1v1h7v2h-7v1h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1H7m3-4v1H9v2h1v1h2v-1h1v-2h-1V9h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightUpDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M22 0v22H0v-2h20V2H0V0h22Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Account() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M9 3h4v1h1v1h1v4h-1v1h-1v1H9v-1H8V9H7V5h1V4h1V3m1 5v1h2V8h1V6h-1V5h-2v1H9v2h1m-3 4h8v1h2v1h1v1h1v4H3v-4h1v-1h1v-1h2v-1m-1 4H5v1h12v-1h-1v-1h-2v-1H8v1H6v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaJFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-3 5v8h-2v-1H8v2h1v1h4v-1h1V6h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowBottomRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M17 8v9H8v-2h5v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7h1V6h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1V8h2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowTopRightCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1M8 7h7v7h-2v-3h-1v1h-1v1h-1v1H9v1H8v-1H7v-1h1v-1h1v-1h1v-1h1V9H8V7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaB() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v3h-1v2h1v3h-1v1H8V6m2 2v2h2V8h-2m2 4h-2v2h2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 2H9v4h4V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ToggleSwitchOn() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 5h16v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5m10 3v1h-1v4h1v1h4v-1h1V9h-1V8h-4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Note() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 3v1h1v1h1v1h1v1h1v1h1v1h1v9h-1v1H2v-1H1V4h1V3h13m0 3h-1v4h4V9h-1V8h-1V7h-1V6M3 5v12h16v-5h-6v-1h-1V5H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaXFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 15h2v-3h-1v-1h-1v-2h1V9h1V6h-2v2h-1v1h-2V8H9V6H7v3h1v1h1v2H8v1H7v3h2v-2h1v-1h2v1h1v2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 12v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h-4v-1H9v-1H8v-2H7V3h2v6h1v2h2v1h3v-1h-1v-1h-1V8h2v1h1v1h1v1h1v1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Card() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 3h18v1h1v14h-1v1H2v-1H1V4h1V3m1 2v12h16V5H3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronLeftCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M13 16h-2v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h-1v2h1v1h1v1h1v2m2-15v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CompassWestArrow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 6h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1H8v1H6v-1H5v-2H4m12-3v2h1v1h1v1h1v1h1v1h2v-2h-1v-1h-1v-1h-1v-2h1V9h1V8h1V6h-2v1h-1v1h-1v1h-1v1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func DoorOpen() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M10 10v2H9v-2h1M6 2h10v1h1v15h2v2H3v-2h2V3h1V2m1 2v14h4V4H7m6 0v1h1V4h-1m1 1v1h1V5h-1m0 1h-1v1h1V6m0 1v1h1V7h-1m0 1h-1v1h1V8m0 1v1h1V9h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuTopRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 18h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V7h11v11m-2-4V9H8v1h1v1h1v1h1v1h1v1h1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Notification() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 20v1h-2v-1H8v-2h6v2h-2M2 15h1v-1h1V6h1V4h1V3h2V2h2V1h2v1h2v1h2v1h1v2h1v8h1v1h1v2H2v-2m4 0h10V7h-1V5h-2V4H9v1H7v2H6v8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaKFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-3h1v1h1v2h2v-3h-1v-1h-1v-1h-1v-1h1V9h1V8h1V6h-2v1h-1v1h-1V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 10V8h-1V7h-1V6h-1V5h-1V4h-2v2h1v1h1v1h-4v1H9v1H8v2H7v7h2v-6h1v-2h2v-1h3v1h-1v1h-1v2h2v-1h1v-1h1v-1h1v-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightRoundDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 22v-2h5v-1h4v-1h2v-1h2v-1h1v-1h1v-1h1v-1h1v-2h1V9h1V5h1V0h2v6h-1v4h-1v2h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-2v1h-2v1H6v1H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CompassNorthWest() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M2 6h2v1.5h1V10h1V6h2v10H6v-2H5v-2H4v4H2m8-10h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pickaxe() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M8 2h3v1h2v1h2v1h2v2h1v2h1v2h1v3h-2v-2h-1v-1h-1v-1h-2V9h-1V8h-1V6h-1V5h-1V4H8m3 5h1v1h1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1H2v-1H1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PlusCircleFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-3-5v-4h4v-2h-4V6h-2v4H6v2h4v4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func VolumeMute() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 7h2v1h1v1h2V8h1V7h2v2h-1v1h-1v2h1v1h1v2h-2v-1h-1v-1h-2v1h-1v1h-2v-2h1v-1h1v-2h-1V9h-1V7M6 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1H9v-1H8v-1H7v-1H6v-1H2V8h4m1 2H4v2h3v1h1v1h1V8H8v1H7v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaU() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v8h2V6h2v9h-1v1H9v-1H8V6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Bookmark() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 2h12v1h1v17h-2v-1h-2v-1h-2v-1h-2v1H8v1H6v1H4V3h1V2m1 2v13h1v-1h2v-1h1v-1h2v1h1v1h2v1h1V4H6Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FlaskRoundBottomEmpty() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M9 1h4v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1V1m1 4v4H9v1H8v1H7v1H6v4h1v1h1v1h1v1h4v-1h1v-1h1v-1h1v-4h-1v-1h-1v-1h-1V9h-1V5h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MinusCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-1 4v2H6v-2h10Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Script() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M20 1H5v1H4v13h2V3h9v16h-2v-1h-1v-1H1v3h1v1h14v-1h1V3h2v2h2V2h-1"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ToggleSwitchOff() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M5 8h4v1h1v4H9v1H5v-1H4V9h1V8m14-3v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5h16m-1 2H4v1H3v6h1v1h14v-1h1V8h-1V7Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaX() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 13v-2h-1v-1h-2v1H9v2H7v-3h1v-1h1v-2H8V9H7V6h2v2h1v1h2V8h1V6h2v3h-1v1h-1v2h1v1h1v3h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Division() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 8h-2V7H9V5h1V4h2v1h1v2h-1Zm5 4H5v-2h12Zm-5 6h-2v-1H9v-2h1v-1h2v1h1v2h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Door() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 10h2v2h-2v-2m4-8v1h1v15h2v2H3v-2h2V3h1V2h10m-1 2H7v14h8V4Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Notebook() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 2v18h-1v1H4v-1H3v-2H1v-2h2v-4H1v-2h2V6H1V4h2V2h1V1h14v1h1m-5 7h-1V8h-1v1h-1v1h-1V3H7v16h10V3h-2v7h-1V9M3 4v2h2V4H3m2 6H3v2h2v-2m0 6H3v2h2v-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GamepadDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9m2 6v4h4v-4H9Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Shield() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M3 4h2V3h2V2h2V1h4v1h2v1h2v1h2v10h-1v2h-1v2h-1v1h-1v1h-2v1H9v-1H7v-1H6v-1H5v-2H4v-2H3V4m7-1v1H8v1H6v1H5v7h1v2h1v2h1v1h2v1h2v-1h2v-1h1v-2h1v-2h1V6h-1V5h-2V4h-2V3h-2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Wall() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M19 20H4v-5H1V7h3V2h15v5h2v8h-2ZM12 7V4H6v3Zm5 0V4h-3v3Zm-9 6V9H3v4Zm11 0V9h-9v4Zm-8 5v-3H6v3Zm6 0v-3h-4v3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUpCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M6 12v-2h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v6h-2v-6H9v1H8v1H6m-5 3V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1m4 1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BorderLeftRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M12 2v2h-2V2h2m0 16v2h-2v-2h2M8 2v2H6V2h2m8 0v2h-2V2h2m0 16v2h-2v-2h2m-8 0v2H6v-2h2M4 2v18H2V2h2m14 0h2v18h-2V2Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BoxOuterLightRoundDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 0h2v5h1v4h1v2h1v2h1v1h1v1h1v1h1v1h2v1h2v1h4v1h5v2h-6v-1h-4v-1h-2v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-2H2v-2H1V6H0V0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronRightCircle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M9 16h2v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H9v2h1v1h1v1h1v2h-1v1h-1v1H9v2m6-15v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BowArrow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M1 3h10v1h2v1h3V4h1V2h3v3h-2v1h-1v3h1v2h1v10h-2v-2h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v1H8v1H7v3H6v1H5v1H4v-1H3v-1H2v-1H1v-1h1v-1h1v-1h3v-1h1v-1h1v-2H7v-1H6V9H5V8H4V7H3V5H1V3m15 15h1v-6h-1v-2h-1V9h-1v1h-1v1h-1v1h-1v1h1v1h1v1h1v1h1v1h1v1M12 7V6h-2V5H4v1h1v1h1v1h1v1h1v1h1v1h1v-1h1V9h1V8h1V7h-1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuBottomRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M4 15v-2h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v11H4m4-2h5V8h-1v1h-1v1h-1v1H9v1H8v1Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func RotateCounterclockwise() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M0 11v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1H5V9h1V7h1V6h2V5h4v1h2v1h1v2h1v4h-1v2h-1v1h-2v1H9v-1H6v2h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v3H0Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AlphaTFill() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v2h2v8h2V8h2V6H8Z"/>`).
		Custom("viewBox", "0,0,22,22").
		Custom("width", "1em").
		Custom("height", "1em")
}
