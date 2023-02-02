/* cSpell:disable */
package charm

import "github.com/gogoracer/racer/pkg/handlebars"

func ClipboardTick() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 1.75h4.5v3.5h-4.5zm4 11.05l1.5 1.5l3-2.5m-9-9h-2.5v11.5h4.5m6-5V2.8h-2.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Gamepad() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 3.75c-2 5-2 9 0 9.5s2.5-2 2.5-2h4.5s.5 2.5 2.5 2s2-4.5 0-9.5h-2l-1 1h-3.5l-1-1z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Music() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="12" r="2.25"/><circle cx="12" cy="11" r="2.25"/><path d="M6.25 12V2.75l8-1V11"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func OctagonWarning() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 1.75h5.5l3.5 3.5v5.5l-3.5 3.5h-5.5l-3.5-3.5v-5.5zM8 11.25v0m0-6.5v3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func WifiPoor() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/><path d="M5 9c.75-1.75 5.25-1.75 6 0"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDownLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 11.75h-6.5v-6.5m7.5-1l-7.5 7.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PaperPlane() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 1.75l12.5 6l-12.5 6.5l1.5-6.5zm2 6h3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func People() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5" cy="9" r="2.25"/><circle cx="11" cy="4" r="2.25"/><path d="M7.75 9.25c0-1 .75-3 3.25-3s3.25 2 3.25 3m-12.5 5c0-1 .75-3 3.25-3s3.25 2 3.25 3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PhoneCross() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm11.5 1l-3.5 3.5m0-3.5l3.5 3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PlantPot() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.75 6.75c0 1.25-.75 3-.75 3m.25-2.5s.75-2-1-3.5s-4.5-1-4.5-1s0 2 1.5 3.5s4 1 4 1zm.5-1s-.75-2 1-3.5s4.5-1 4.5-1s0 2-1.5 3.5s-4 1-4 1zm-4 3.5h6.5s.5 4.5-3.25 4.5s-3.25-4.5-3.25-4.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUpRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 4.25h6.5v6.5m-7.5 1l7.5-7.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Camera() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75v8.5h12.5v-8.5h-3l-1.5-2h-3.5l-1.5 2z"/><circle cx="8" cy="8.5" r="2.25"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronsUpDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 10.75L8 14.25l-3.25-3.5m6.5-5.5L8 1.75l-3.25 3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Extensions() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 8.75h5.5v5.5m5-12.5v4m-2-2h4m-12.5-1v11.5h11.5v-5.5h-6v-6z"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MonitorArrow() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.2 7.75v3.5H1.7v-9.5h6.5M4.75 14.2h6.5M8 11.7v2.5m1.75-7.95l4.5-4.5m-3.5-.5h4v4"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TreeFir() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l-4.25 5.5h2.5l-3.5 4h4v3h2.5v-3h4l-3.5-4h2.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Database() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 1.75c-3.75 0-5.25 2-5.25 2v8.5s1.5 2 5.25 2s5.25-2 5.25-2v-8.5s-1.5-2-5.25-2z"/><path d="M2.75 8.25s1.5 2 5.25 2s5.25-2 5.25-2m-10.5-4s1.5 2 5.25 2s5.25-2 5.25-2"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MediaFastForward() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 3.75v8.5l6-4.25zm-6.5 0v8.5l6-4.25z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MediaPlay() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75v10.5L12.25 8z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Message() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 14.25V2.75h12.5v8.5h-8.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Power() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 1.75v6.5m4.25-5s2 1.298 2 4.75a6.25 6.25 0 1 1-12.5 0c0-3.452 2-4.75 2-4.75"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Sun() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="3.25"/><path d="m2.75 13.25l.5-.5m9.5 0l.5.5m-.5-10l.5-.5m-10 .5l-.5-.5M2.25 8h-1m13.5 0h-1M8 13.75v1m0-13.5v1"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func StackPush() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 7.25L1.75 8L8 11.25L14.25 8l-1.5-.75M1.75 11L8 14.25L14.25 11"/><path d="M8 8.25v-6.5m-2.25 4.5l2.25 2l2.25-2"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ZoomIn() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="7.5" cy="7.5" r="4.75"/><path d="M9.25 7.5h-3.5M7.5 5.75v3.5m3.75 2l3 3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AppsMinus() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm5.05-6h-5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Atom() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><ellipse cx="11.3" rx="8.28" ry="3.17" transform="rotate(45)"/><ellipse cy="11.3" rx="8.28" ry="3.17" transform="rotate(315)"/><path d="M8 8v0"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Chip() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 2.75h10.5v10.5H2.75z"/><path d="M6.25 6.25h3.5v3.5h-3.5zm-4 4h-1m1-4.5h-1m13.5 4.5h-1m1-4.5h-1m-3.5 8v1m-4.5-1v1m4.5-13.5v1m-4.5-1v1"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Map() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 5.25v8.5m-4.5-10.5v8.5m-4 2.5v-9.5l4-2l4.5 2l4-2v9.5l-4 2l-4.5-2z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ShieldKeyhole() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm0 5.5v3"/><circle cx="8" cy="6.5" r=".75" fill="currentColor"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SquareTick() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.25 2.75h-7.5v10.5h10.5v-3.5"/><path d="m5.75 7.75l2.5 2.5l6-6.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitCommit() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="2.25"/><path d="M8 10.75v3.5m0-12.5v3.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuMeatball() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="2.5" cy="8" r=".75"/><circle cx="8" cy="8" r=".75"/><circle cx="13.5" cy="8" r=".75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Signpost() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 9.25h10.5l2-2.25l-2-2.25H1.75zm5.5.5v4.5m0-12.5v2.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Eraser() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 13.25h-9.5l-3-3l7.5-7.5l5 5l-5.5 5.5m-3.5-6.5l5 5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Flame() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 7.75c2 2 2.5-2.5 3.5-2s1.5 2 1.5 3.25c0 3.25-2.35 5.25-5.25 5.25s-5.25-2.5-5.25-6s3.5-7 5.5-7c0 0-2 4.5 0 6.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Move() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.25 10.25l2-2.25l-2-2.25m-2-2L8 1.75l-2.25 2m-2 2L1.75 8l2 2.25m2 2l2.25 2l2.25-2M8 1.75v12M13.75 8h-12"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Share() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="8" r="2.25"/><circle cx="12" cy="12" r="2.25"/><circle cx="12" cy="4" r="2.25"/><path d="m6 9l4 2M6 7l4-2"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.25 8.75l4.5 4.5l4.5-4.5m-4.5-6v10.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ClockAlarm() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.75 1.75l2.5 2m-10-2l-2.5 2m10.5 9.5l1 1m-9.5-1l-1 1m5.5-7.5v2.5l-1.5 1"/><circle cx="8" cy="9" r="5.25"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuHamburger() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 12.25h10.5m-10.5-4h10.5m-10.5-4h10.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SquareCross() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.25 5.75l-4.5 4.5m0-4.5l4.5 4.5m-7.5-7.5h10.5v10.5H2.75z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 10.25L8 5.75l-4.25 4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Eye() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 8s2-4.25 6.25-4.25S14.25 8 14.25 8s-2 4.25-6.25 4.25S1.75 8 1.75 8z"/><circle cx="8" cy="8" r="1.25" fill="currentColor"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Mail() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 3.75h12.5v9.5H1.75z"/><path d="m2.25 4.25l5.75 5l5.75-5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Monitor() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h12.5v9.5H1.75zm3 12.5h6.5M8 11.75v2.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FileCode() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5h-2"/><path d="M7.75 2.25v5h5.05M6.75 10.8l2 1.75l-2 1.75m-3-3.5l-2 1.75l2 1.75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitBranch() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="11.5" cy="3.5" r="1.75"/><circle cx="4.5" cy="12.5" r="1.75"/><path d="M5.25 8.25c3 0 6 .5 6-2.5m-6.5 4.5v-4.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Padlock() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 6.75h10.5v7.5H2.75zm2-.5s-1-4.5 3.25-4.5s3.25 4.5 3.25 4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Certificate() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.25 1.75h-8.5v11.5h2.5m3.5-3.5l-.5 4.5l2.25-1l2.25 1l-.5-4.5"/><circle cx="10.5" cy="7.5" r="2.75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Filter() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h12.5v1.5l-5 5.5v4l-2.5 1.5v-5.5l-5-5.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Forward() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 13.25c.5-6 5.5-7.5 8-7v-3.5L14.25 8l-4.5 5.25v-3.5c-2.5-.5-6.5.5-8 3.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func NorthStar() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.75 7.75h-12m6-6v12m-3.5-2.5l7-7m0 7l-7-7"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func NotesCross() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 14.25h-5.5V1.75h10.5v6.5m1 2.5l-3.5 3.5m-5-6.5h4.5m-4.5 3h1.5m-1.5-6h4.5m.5 6l3.5 3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Nut() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5v-6.5z"/><circle cx="8" cy="8" r="2.25"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ShieldTick() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5z"/><path d="m5.75 7.75l1.5 1.5l3-3.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 3.75L5.75 8l4.5 4.25"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func EyeSlash() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.75 3.75c3.5.5 5.5 4.25 5.5 4.25s-.5 1.25-1.5 2.25m-2.5 1.5c-6 2-8.5-3.75-8.5-3.75s.5-1.75 3-3.25"/><path fill="currentColor" d="M8.625 9.083a1.25 1.25 0 0 1-1.649-.366a1.25 1.25 0 0 1 .22-1.675L8 8z"/><path d="m3.75 1.75l8.5 12.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Id() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><circle cx="8" cy="7.5" r="2.25"/><path d="M4.75 12.75c0-1 .75-3 3.25-3s3.25 2 3.25 3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LinkSlash() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.75 1.75l-5.5 12.5m4.5-9.5c3 0 4.5 1.5 4.5 3.25s-1.5 3.25-4.5 3.25m-3.5-6.5c-3 0-4.5 1.5-4.5 3.25s1.5 3.25 4.5 3.25"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MapPin() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.25 7c0 3.75-5.25 7.25-5.25 7.25S2.75 10.75 2.75 7a5.25 5.25 0 0 1 10.5 0z"/><circle cx="8" cy="7" r="1.25" fill="currentColor"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Person() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="6" r="3.25"/><path d="M2.75 14.25c0-2.5 2-5 5.25-5s5.25 2.5 5.25 5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Globe() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M2 8.25h12M8.25 14.2C11 11 11 5 8.25 1.8m-.5 12.4C5 11 5 5 7.75 1.8"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Printer() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.75 9.75h6.5v4.5h-6.5z"/><path d="M4.75 4.25v-2.5h6.5v2.5m-7 8h-2.5v-7.5h12.5v7.5h-2.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ThumbUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 5.75c1.5 0 3-4 4.5-4v4h4.5s-.5 7.5-3.5 7.5h-5.5zm0 0h-3.5v7.5h3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Copyright() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M10 6.75s-.75-1-2-1s-2.25 1-2.25 2.25s1 2.25 2.25 2.25s2-1 2-1"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FaceNeutral() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 4.5h4.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitRequestCross() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M12.25 7.25v3m-8.5-4.5v4.5m10.5-8.5l-3.5 3.5m0-3.5l3.5 3.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitRequestDraft() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M7.75 2.75h.5m2.5 0h.5m1.5 2.5v-.5m0 3v.5m-9-2.5v4.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func RotateClockwise() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 5.25h3m0 3.5c0 2.5-2.798 5.5-6.25 5.5a6.25 6.25 0 1 1 0-12.5c3.75 0 6.25 3.5 6.25 3.5v-3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Speaker() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 1.75h9.5v12.5h-9.5zm5 2.5h-.5"/><circle cx="8" cy="9.5" r="2.25"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronsUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 12.25L8 7.75l-4.25 4.5m8.5-5L8 2.75l-4.25 4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cross() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.25 4.75l-6.5 6.5m0-6.5l6.5 6.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SignIn() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 2.25h-3.5v12h3.5m4-9.5l-3.5 3.5l3.5 3.5m5-3.5h-8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Candy() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="3.25"/><path d="M7.25 11.25c0 1-.5 2.5-1.5 3c-.75 0-1.5-1-2-2c-1-.5-2-1.5-2-2c.5-1 2-1.5 3-1.5m4-4c0-1 .5-2.5 1.5-3c.75 0 1.5 1 2 2c1 .5 2 1.5 2 2c-.5 1-2 1.5-3 1.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 5.75L8 10.25l4.25-4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitCherryPick() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5" cy="8" r="2.25"/><path d="M5 10.75v3.5m0-12.5v3.5M11.75 8h1.5m-4.5-3.25h1.5l1 3.25l-1 3.25h-1.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ShoppingBag() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 4.75h10.5v9.5H2.75z"/><path d="M5.75 7.75c0 1.5 1 2.5 2.25 2.5s2.25-1 2.25-2.5m-7.5-3l1.5-3h7.5l1.5 3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func WifiWarning() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 4.75C11 2 5 2 1.75 4.75L8 13.25l1-1.5m3.25 2v0m0-6v3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Link() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 4.75c3 0 4.5 1.5 4.5 3.25s-1.5 3.25-4.5 3.25M5.75 8h4.5m-4-3.25c-3 0-4.5 1.5-4.5 3.25s1.5 3.25 4.5 3.25"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Messages() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 14.25v-9h-9.5v6h6z"/><path d="m4.75 7.25l-3 3v-8.5h10v3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SwapVertical() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.75 5.75l-3-3l-3 3m3 7.5V2.75m9.5 7.5l-3 3l-3-3m3-7.5v10.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Umbrella() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8.25s.5-6.5 6.25-6.5s6.25 6.5 6.25 6.5zm6 .5v4s0 1.5 1.5 1.5s1.5-1.5 1.5-1.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FileBinary() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5M1.75 10.8h3v3.5h-3z"/><path d="M7.25 14.2h3m-3-3.5h1.5v3m-1-11.45v5h5.05"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pin() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.25 10.25l4 4m-12.5-7.5l5-5s1 2 2 3s4.5 2 4.5 2l-6.5 6.5s-1-3.5-2-4.5s-3-2-3-2z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Shield() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Stack() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11L8 14.25L14.25 11M1.75 8L8 11.25L14.25 8M8 1.75L1.75 5L8 8.25L14.25 5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Archive() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v3.5H1.75zm5 6.5h2.5m-6.5-2.5v7.5h10.5v-7.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cards() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75H10v11.5H1.75zm8.25 1l4.25 2l-4.25 7.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Folder() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75v10.5h12.5v-8.5h-6l-1.5-2z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Coffee() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 11.25c4.5 0 4.5-5.5 0-5.5h-9v5c0 5 8.5 5 8.5 0v-5m-1.5-4v1.5m-3-1.5v1.5m-3-1.5v1.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GraduateCap() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 9.25V6L8 2.75L1.75 6L8 9.25l3.25-1.5v3.5c0 1-1.5 2-3.25 2s-3.25-1-3.25-2v-3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Trophy() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 10.75h6.5v3.5h-6.5zm3.25-2v2m-3.25-9c-1.5 0-3 .5-3 2.25s1.5 2.25 3 2.25m6.5-4.5c1.5 0 3 .5 3 2.25s-1.5 2.25-3 2.25m-6.5-4.5h6.5v3.5c0 1.5-1 3-3.25 3s-3.25-1.5-3.25-3z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Organisation() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 1.75h3.5v3.5h-3.5zm4 9h3.5v3.5h-3.5zm-8 0h3.5v3.5h-3.5zm5.75-5v2m-3.75 2.5v-2h7.5v2"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TickDouble() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 9.75l2.5 2.5m3.5-4l2.5-2.5m-4.5 4l2.5 2.5l6-6.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Bin() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 4.25v-2.5h4.5v2.5m-6.5 1v9h8.5v-9m-9.5-.5h10.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MediaPause() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h3.5v10.5h-3.5zm7 0h3.5v10.5h-3.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Briefcase() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75h12.5v9.5H1.75z"/><path d="M1.75 6.25s-.5 3.5 3 3.5h6.5c3.5 0 3-3.5 3-3.5m-8.5-2v-2.5h4.5v2.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Phone() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PhoneIncoming() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5z"/><path d="m13.25 2.75l-3.5 3.5m0-3v3h3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SoundUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm9 .5s1 .5 1 1.75s-1 1.75-1 1.75m1-6.5c2 1 3 2.5 3 4.75s-1 3.75-3 4.75"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LayoutDashboard() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.5 4h5.5m-11.5 2.5h5.5m.25-6v9.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BluetoothSearching() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.25L10.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25m11.5 1.5s1 .5 1 1.75s-1 1.75-1 1.75m-2-1.75v0"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Bug() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="10" r="4.25"/><path d="M14.25 10.25h-1.5m-1 2.5l1.5 1.5m0-8.5l-1.5 1.5m-10 3h1.5m1 2.5l-1.5 1.5m0-8.5l1.5 1.5m1.5-1.5s-.75-3 2.25-3s2.25 3 2.25 3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Paperclip() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 10.25v-7s0-1.5-1.75-1.5s-1.75 1.5-1.75 1.5v8s0 3 3.25 3s3.25-3 3.25-3v-4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Tent() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.25 14.25L8 10l2.75 4.25"/><path d="m9.75 1.75l-8 12.5h12.5l-8-12.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Bluetooth() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 11.25L12.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BookOpen() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 3.75c-1.75-1-2.25-1-6.25-1v9.5c4 0 4.5 0 6.25 1c1.75-1 3.25-1 6.25-1v-9.5c-4 0-4.5 0-6.25 1zm0 .5v8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FileSymlink() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5h-4"/><path d="M7.75 2.25v5h5m-10 7l3.5-3.5m0 3v-3h-3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GrabHorizontal() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="2.5" cy="5.5" r=".75"/><circle cx="8" cy="5.5" r=".75"/><circle cx="13.5" cy="5.5" r=".75"/><circle cx="2.5" cy="10.5" r=".75"/><circle cx="8" cy="10.5" r=".75"/><circle cx="13.5" cy="10.5" r=".75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Notes() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 1.75h10.5v12.5H2.75zm3 6h4.5m-4.5 3h2.5m-2.5-6h4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Triangle() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 2.75l-6.25 11.5h12.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUp() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.75 7.25l4.5-4.5l4.5 4.5m-4.5 6V2.75"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AtSign() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.25 8c0 3.25 4 3.25 4 0A6.25 6.25 0 1 0 8 14.25c2.25 0 3.25-1 3.25-1"/><circle cx="8" cy="8" r="2.25"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Clipboard() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.75 1.75h4.5v3.5h-4.5z"/><path d="M5.25 2.75h-2.5v11.5h10.5V2.75h-2.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func WifiFair() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/><path d="M4.25 8c2-1.75 5.5-1.75 7.5 0"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Newspaper() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.2 14.2h.5c1.5 0 2.5-1 2.5-2.5v-6h-3m-9.5-4h9.5v12.5h-7c-1.5 0-2.5-1-2.5-2.5V2.26zm3.05 9.5h3.5"/><path d="M4.75 4.75h3.5v3.5h-3.5z"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pencil() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.25v3h3l9.5-9.5l-3-3zm7-6.5l2.5 2.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Anchor() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 5.75V14M3.25 7.75h-1.5c0 4 2.5 6.5 6.25 6.5s6.25-2.5 6.25-6.5h-1.5"/><circle cx="8" cy="3.5" r="1.75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Circle() *handlebars.ElementSVG {
	return handlebars.SVG(`<circle cx="8" cy="8" r="6.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CreditCard() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v9.5H1.75zm8 6.5h1.5m-9-3h11.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Crop() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.25 1.75v10h10M11.8 14.2v-2.5m0-2.5v-5h-5m-2.5 0H1.8"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Download() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 13.25h9m-8.5-6.5l4 3.5l4-3.5m-4-5v8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Infinity() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 5c2.5 1 3.5 5 6 6s3.25-1.25 3.25-3S13.5 4 11 5s-3.5 5-6 6s-3.25-1.25-3.25-3S2.5 4 5 5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LayoutRows() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LightningBolt() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.25 1.75l-6.5 7.5l4.5.5l-.5 4.5l6.5-7.5l-4.5-.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CameraVideo() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75h7.5v7.5h-7.5zm8 2.5l4.5-2.5v7.5l-4.5-2.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CircleCross() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.25 5.75l-4.5 4.5m0-4.5l4.5 4.5"/><circle cx="8" cy="8" r="6.25"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ConicalFlask() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 1.75h6.5m-6.5 8h6.5m-5.5-7.5v4.5l-4 7.5h12.5l-4-7.5v-4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Flag() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 14.25v-11s2-1.5 4-1.5s2.5 1.5 4.5 1.5s4-1.5 4-1.5v7s-2 1.5-4 1.5s-2.5-1.5-4.5-1.5s-4 1.5-4 1.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Github() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 14.25s-.5-2 .5-3c0 0-2 0-3.5-1.5s-1-4.5 0-5.5c-.5-1.5.5-2.5.5-2.5s1.5 0 2.5 1c1-.5 3.5-.5 4.5 0c1-1 2.5-1 2.5-1s1 1 .5 2.5c1 1 1.5 4 0 5.5s-3.5 1.5-3.5 1.5c1 1 .5 3 .5 3m-5-.5c-1.5.5-3-.5-3.5-1"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Home() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 5.75v7.5h8.5v-7.5m-10.5 1.5L8 1.75l6.25 5.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func RotateAntiClockwise() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 5.25h-3m0 3.5c0 2.5 2.798 5.5 6.25 5.5a6.25 6.25 0 1 0 0-12.5c-3.75 0-6.25 3.5-6.25 3.5v-3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PhoneOutgoing() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 4.5l3.5-3.5m0 3v-3h-3"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ThumbDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 10.25c1.5 0 3 4 4.5 4v-4h4.5s-.5-7.5-3.5-7.5h-5.5zm0 0h-3.5v-7.5h3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.25 3.75l-4.5 4.5l4.5 4.5m6-4.5H2.75"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronsRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 12.25L8.25 8l-4.5-4.25m5 8.5L13.25 8l-4.5-4.25"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Crosshair() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 5.25v-3m0 11.5v-3M10.75 8h3M2.25 8h3"/><circle cx="8" cy="8" r="6.25"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Diamond() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.25 8L8 14.75L14.75 8L8 1.25z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitFork() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="12.5" r="1.75"/><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="11.5" cy="3.5" r="1.75"/><path d="M8 8.75v1.5m-3.25-4.5c0 3.5 6.5 3.5 6.5 0"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Moon() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8c0 3.45 2.8 6.25 6.25 6.25c3.41-.003 6.25-3 6.25-6c-1 .5-4 1.5-6-.5s-1-5-.5-6c-3 0-6 2.84-6 6.25z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Wifi() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 12.25L10.25 8l-4.5-4.25"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronsLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 3.75L7.75 8l4.5 4.25m-5-8.5L2.75 8l4.5 4.25"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PhoneCall() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 0c2.5 0 4.5 2 4.5 4.5m-4.5-2c1 0 2 1 2 2"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ScreenMinimise() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 10.75h3.5v3.5m5.5 0v-3.5h3.5m0-5.5h-3.5v-3.5m-5.5 0v3.5h-3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SwapHorizontal() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.75 8.25l-3 3l3 3m7.5-3H2.75m7.5-9.5l3 3l-3 3m-7.5-3h10.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Ticket() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v3s-2 0-2 1.75s2 1.75 2 1.75v3H1.75v-3s2 0 2-1.75s-2-1.75-2-1.75zm7 8v1.5m0-9.5v1.5m0 2.5v1.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Container() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.75 12.2l5.5 2l7-4.5v-6l-5.5-2l-7 4.5z"/><path d="M10.8 6.25v5.5m-3.5-3.5v6m-5.5-8l5.5 2l7-4.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FaceFrown() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 5s.5-1 2.25-1s2.25 1 2.25 1"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderSymlink() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.75 13.25l3.5-3.5m0 3v-3h-3"/><path d="M8.25 13.25h6v-8.5h-6l-1.5-2h-5v4"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Glasses() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="11" r="2.25"/><circle cx="12" cy="11" r="2.25"/><path d="M14.25 10.75c-1.5-6-2-6.5-3.5-7m-9 7c1.5-6 2-6.5 3.5-7m1 7c1-1 2.5-1 3.5 0"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Terminal() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><path d="M8.75 10.25h2.5m-6.5-4.5L7.25 8l-2.5 2.25"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Block() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="m4.25 11.75l8-8"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Copyleft() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M6 6.75s.75-1 2-1s2.25 1 2.25 2.25s-1 2.25-2.25 2.25s-2-1-2-1"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Files() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9.25 1.75l4 4v5.5h-7.5v-9.5z"/><path d="M9.25 2.25v3.5h3.5m-2.5 6v2.5h-7.5v-9.5h2.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GrabVertical() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5.5" cy="2.5" r=".75"/><circle cx="5.5" cy="8" r=".75"/><circle cx="5.5" cy="13.5" r=".75"/><circle cx="10.496" cy="2.5" r=".75"/><circle cx="10.496" cy="8" r=".75"/><circle cx="10.496" cy="13.5" r=".75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Microphone() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 1.75c-2.25 0-2.25 2-2.25 3v1.5c0 1 0 3 2.25 3s2.25-2 2.25-3v-1.5c0-1 0-3-2.25-3z"/><path d="M8 13v1.25m-5.25-6.5s0 4.5 5.25 4.508c5.25.008 5.25-4.508 5.25-4.508"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Star() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l-2.25 4l-4 .5l3 3.5l-1 4.5l4.25-2l4.25 2l-1-4.5l3-3.5l-4-.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MediaEject() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 11.25h10.5L8 2.75zm10.5 3H2.75"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Droplet() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 9a5.25 5.25 0 1 0 10.5 0C13.25 5.75 8 1.75 8 1.75S2.75 5.75 2.75 9z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CircleWarning() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8 10.75v0m0-6v3.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cube() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5zM8 14V8m5.75-3L8 8M2 5l6 3"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Gem() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 2.75h6.5l3 3.5l-6.25 7l-6.25-7z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Headphones() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.75c0-2.5 3.5-2 3.5-2v4.5s-3.5.5-3.5-2.5v-3.5c0-3 .5-6.5 6.25-6.5s6.25 3.5 6.25 6.5v3.5c0 3-3.5 2.5-3.5 2.5v-4.5s3.5-.5 3.5 2"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Hourglass() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.75 13.75c0-5-2-4-2-5.75s2-.75 2-5.75m-7.5 11.5c0-5 2-4 2-5.75s-2-.75-2-5.75m-1.5-.5h10.5m-10.5 12.5h10.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LayoutList() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm3.5.5v9.5m-3-6.5h11.5m-11.5 3.5h11.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LayoutColumns() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.25.5v9.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ScreenMaximise() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 14.25h-3.5v-3.5m12.5 0v3.5h-3.5m0-12.5h3.5v3.5m-12.5 0v-3.5h3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ShieldCross() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm1.75 4l-3.5 3.5m0-3.5l3.5 3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func WifiSlash() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 3.25c-1.5 0-3.5 1.5-3.5 1.5L8 13.25l2.25-3m-1.5-7.5s2.977-.135 5.5 2l-2 2.5m-8-5.5l8 10.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Apps() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm0-8h4.5v4.5h-4.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Compass() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="m6.75 6.75l-1 4l3.5-1.5l1-4z"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FloppyDisk() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 2.75v10.5h10.5v-7.5l-3-3z"/><path d="M5.75 13.25v-3.5h4.5v3.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MediaSkip() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 13.25L11.25 8l-8.5-5.25zm11.5-9.5v8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Scales() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 3.75c1 1 2.5 1.5 4 0h4.5c1.5 1.5 3 1 4 0M8 1.75v12m-3.25.5h6.5"/><path d="m12.75 4.75l-2 5c.5 1 3.5 1 4 0zm-9.5 0l-2 5c.5 1 3.5 1 4 0z"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Server() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.25h12.5v10H1.75zm.5 5h11.5m-9 2.5v0m0-5v0m6.5 0h-3m3 5h-3"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Bell() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 1.75c-2.468 0-4.25 1.5-4.25 3.5v3l-2 3.5h12.5l-2-3.5v-3c0-2-1.166-3.5-4.25-3.5zm-2.25 10.5c0 3 4.5 3 4.5 0"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BellSlash() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.75 12.25c0 3 4.5 3 4.5 0m2-4v-3c0-2-1.166-3.5-4.25-3.5m-3.75 2c-.53.585-.5.674-.5 1.5v3l-2 3.5h8.5"/><path d="M5.75 12.25c0 3 4.5 3 4.5 0m-7.5-10.5l10.5 12.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BluetoothSlash() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 6.25L12.25 5l-4.5-3.25v2.5m4.5 6.75l-4.5 3.25v-6l-4 3m-2-8l12.5 9"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PhoneForward() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 3h4.5m-2 2l2-2l-2-2"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SoundDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm9 .5s1 .5 1 1.75s-1 1.75-1 1.75"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Tag() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7.25 14.25l-5.5-5.5l7-7h5.5v5.5z"/><circle cx="11" cy="5" r=".5" fill="currentColor"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Inbox() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 13.25h12.5v-5l-2.5-5.5h-7.5l-2.5 5.5z"/><path d="M2.25 8.75h3l1 1.5h3.5l1-1.5h3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Mobile() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 1.75h8.5v12.5h-8.5zm4.5 10h-.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ShieldWarning() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm0 9v0m0-5.5v3"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SignOut() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 2.25h-3.5v12h3.5m5.5-9.5l3.5 3.5l-3.5 3.5m-5-3.5h8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Bookmark() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 1.75h8.5v12.5L8 9.75l-4.25 4.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Clover() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 2.75C4.25 4.25 6 6 8 8c2-2 3.75-3.75 3.25-5.25s-2.5-1-3.25.5c-.75-1.5-2.75-2-3.25-.5zM8 8c2 2 3.75 3.75 5.25 3.25s1-2.5-.5-3.25c1.5-.75 2-2.75.5-3.25S10 6 8 8zm0 0c-2 2-3.75 3.75-3.25 5.25s2.5 1 3.25-.5c.75 1.5 2.75 2 3.25.5S10 10 8 8zm0 0C6 6 4.25 4.25 2.75 4.75s-1 2.5.5 3.25c-1.5.75-2 2.75-.5 3.25S6 10 8 8z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitCompare() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M3.75 5.75v5c0 1 .5 1.5 1.5 1.5h2m-.5 2l1.5-2l-1.5-2m5.5 0v-5c0-1-.5-1.5-1.5-1.5h-2m.5-2l-1.5 2l1.5 2"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Hash() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 10.25h9.5m-8.5-4.5h9.5m-2.5-4l-1.5 12.5m-2.5-12.5l-1.5 12.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func NotesTick() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.25 14.25h-4.5V1.75h10.5v7.5m-3.5 3.5l1.5 1.5l3-2.5m-8.5-4h4.5m-4.5 3h1.5m-1.5-6h4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Plus() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.75 7.75h-10m5-5v10"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDownRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 11.75h6.5v-6.5m-7.5-1l7.5 7.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowRight() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.75 3.25l4.5 4.5l-4.5 4.5m-6-4.5h10.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Clock() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8.25 4.75v3.5l-2.5 2"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Laptop() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h10.5v7.5H2.75zm0 7.5l-1 3h12.5l-1-3"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MediaRewind() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.75 3.75v8.5L1.75 8zm6.5 0v8.5L8.25 8z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Swords() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2.75 9.25l1.5 2.5l2 1.5m-4.5 0l1 1m1.5-2.5l-1.5 1.5m3-1l8.5-8.5v-2h-2l-8.5 8.5"/><path d="M10.25 12.25L8 10m2-2l2.25 2.25m1-1l-1.5 2.5l-2 1.5m4.5 0l-1 1m-1.5-2.5l1.5 1.5M6 8L1.75 3.75v-2h2L8 6"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func BluetoothConnected() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 11.25L12.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25M1.75 8h1.5m9.5 0h1.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cog() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="1.75"/><path d="m6.75 1.75l-.5 1.5l-1.5 1l-2-.5l-1 2l1.5 1.5v1.5l-1.5 1.5l1 2l2-.5l1.5 1l.5 1.5h2.5l.5-1.5l1.5-1l2 .5l1-2l-1.5-1.5v-1.5l1.5-1.5l-1-2l-2 .5l-1.5-1l-.5-1.5z"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SoundMute() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm12.5 0l-3.5 4.5m0-4.5l3.5 4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Tablet() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 1.75h10.5v12.5H2.75zm5.5 10h-.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Diff() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 13.75h8m0-7.5h-8m4-4v8"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Refresh() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 10.75h-3m12.5-2c0 3-2.798 5.5-6.25 5.5c-3.75 0-6.25-3.5-6.25-3.5v3.5m9.5-9h3m-12.5 2c0-3 2.798-5.5 6.25-5.5c3.75 0 6.25 3.5 6.25 3.5v-3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronsDown() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 3.75L8 8.25l4.25-4.5m-8.5 5L8 13.25l4.25-4.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CircleTick() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 8.75c-.5 2.5-2.385 4.854-5.03 5.38A6.25 6.25 0 0 1 3.373 3.798C5.187 1.8 8.25 1.25 10.75 2.25"/><path d="m5.75 7.75l2.5 2.5l6-6.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Gift() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75h12.5v3.5H1.75z"/><path d="M10.25 4.75H8c0-2 .5-3 2.25-3c2 0 2 3 0 3zm-4.5 0H8c0-2-.5-3-2.25-3c-2 0-2 3 0 3zm2.25 9V5M2.75 8.75v5.5h10.5v-5.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Quote() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.25 3.75h-4.5v5.5c0 3.5 2.5 4.5 4.5 4c-1.5-1.5-1.5-2.5-1.5-4h1.5zm7 0h-4.5v5.5c0 3.5 2.5 4.5 4.5 4c-1.5-1.5-1.5-2.5-1.5-4h1.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Telescope() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 5.75l1 2.5m3.5-4.5l1.5 3.5m-9 0l1 2.5l11.5-3.5l-2-4.5zm6 3.95v3m-3-.5L7 11.2l1.75-.5l2.5 3"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CircleMinus() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M4.75 8h6.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Disc() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><circle cx="8" cy="8" r="1.75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitRequest() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="m9.25 1.75l-1.5 2l1.5 2m3 4.5v-5c0-1-.5-1.5-1.5-1.5h-2m-5 2v4.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Help() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M5.75 6.75c0-1 1-2 2.25-2s2.25 1.034 2.25 2c0 1.5-1.5 1.5-2.25 2m0 2.5v0"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LayoutStackV() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.25.5v9.5M8 8h6"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func StackPop() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.25 6.75L1.75 8L8 11.25L14.25 8l-2.5-1.25M1.75 11L8 14.25L14.25 11M8 8.25v-6.5m-2.25 2l2.25-2l2.25 2"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Heart() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 9.75c3 3.5 4.75 4.5 4.75 4.5s1.75-1 4.75-4.5s1-7-1.5-7s-3.25 3-3.25 3s-.75-3-3.25-3s-4.5 3.5-1.5 7z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MonitorCross() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.2 7.75v3.5H1.7v-9.5h6.5M4.75 14.2h6.5M8 11.7v2.5m6.2-12.45l-3.5 3.5m0-3.5l3.5 3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Octagon() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 1.75h5.5l3.5 3.5v5.5l-3.5 3.5h-5.5l-3.5-3.5v-5.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Sword() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 9.25l1.5 2.5l2 1.5m-4.5 0l1 1m1.5-2.5l-1.5 1.5m3-1l8.5-8.5v-2h-2l-8.5 8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Tick() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 8.75l3.5 3.5l7-7.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func AppsPlus() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm5.05-6h-5m2.5-2.5v5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Calendar() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v10.5H1.75zm9.5-2v1.5m-6.5-1.5v1.5m-2.5 4h11.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Code() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 11.25L1.75 8l3.5-3.25m5.5 6.5L14.25 8l-3.5-3.25"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Hexagon() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Image() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><path d="m3.75 13.2l6.5-5.5l4 3"/><circle cx="5.25" cy="6.25" r=".5" fill="currentColor"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LayoutSidebar() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm4.5.25v9.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUpLeft() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 4.25h-6.5v6.5m7.5 1l-7.5-7.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cast() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.25v-2.5h12.5v10.5h-4.5m-8-5a5 5 0 0 1 5 5m-5-2.5a2.5 2.5 0 0 1 2.5 2.5m-2.5 0v0"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MenuKebab() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="2.5" r=".75"/><circle cx="8" cy="8" r=".75"/><circle cx="8" cy="13.5" r=".75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Square() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h10.5v10.5H2.75z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ZoomOut() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="7.5" cy="7.5" r="4.75"/><path d="M9.25 7.5h-3.5m5.5 3.75l3 3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Binary() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 1.75h3v4.5h-3zm6.5 4.5h3m-3-4.5h1.5v4m-1.5 4h3v4.5h-3zm-6.5 4.5h3m-3-4.5h1.5v4"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func File() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 1.75h5.5l5 5v7.5H2.75z"/><path d="M7.75 2.25v5h5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func GitMerge() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="4.5" cy="12.5" r="1.75"/><circle cx="12.5" cy="8.5" r="1.75"/><path d="M4.75 10.25v-4.5c1 2 2 3 5.5 3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LayoutGrid() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12m-3.75-4.75v9.5m-4.5-9.5v9.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cloud() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 3.75A3.25 3.25 0 0 0 3.75 7c.002.255.033.508.094.756h-.002A2.25 2.25 0 0 0 4 12.25h7.5a2.75 2.75 0 1 0-1.252-5.2L10.25 7A3.25 3.25 0 0 0 7 3.75z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Copy() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 4.25v-2.5h-9.5v9.5h2.5m.5-6.5v9.5h9.5v-9.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Rocket() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.25 9.75l-2-.5s0-1.5.5-3s4-1.5 4-1.5m-.5 7l.5 2s1.5 0 3-.5s1.5-4 1.5-4m-7 .5l2 2s5-2 6.5-4.5s1.5-5.5 1.5-5.5s-3 0-5.5 1.5s-4.5 6.5-4.5 6.5z"/><path fill="currentColor" d="m1.75 14.25l2-1l-1-1z"/><circle cx="10.25" cy="5.75" r=".5" fill="currentColor"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Book() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.75 11.75v2m1.5.5h-9c-.75 0-1.5-.5-1.5-1.5s.75-1.5 1.5-1.5h9v-9.5h-9c-.75 0-1.5.75-1.5 1.5v9.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Search() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.25 11.25l3 3"/><circle cx="7.5" cy="7.5" r="4.75"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Upload() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 2.75h9m-8.5 6.5l4-3.5l4 3.5m-4 5v-8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CameraVideoSlash() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.25 10.75l3 1.5v-7.5l-5 2.5v-2.5h-2.5m1.5 7.5h-6.5v-7.5h1.5m-1.5-2.5l8.5 12"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Gitlab() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 14.25l-6.25-4.5l2-8l2 5.5h4.5l2-5.5l2 8z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Reply() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 13.25c-.5-6-5.5-7.5-8-7v-3.5L1.75 8l4.5 5.25v-3.5c2.5-.5 6.5.5 8 3.5z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Snowflake() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.75 7.75h-12m6-6v12m-2.5-1l2.5-2.5l2.5 2.5m-7.5-7.5l2.5 2.5l-2.5 2.5m7.5-7.5l-2.5 2.5l-2.5-2.5m7.5 7.5l-2.5-2.5l2.5-2.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func StickyNote() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.25 13.25h-6.5V2.75h10.5v6.5z"/><path d="M8.75 13.25v-4.5h4.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cursor() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 1.75l4.5 12.5l2.5-5.5l5.5-2.5zm7.5 7.5l4 4"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FaceSmile() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 4s.5 1.5 2.25 1.5s2.25-1.5 2.25-1.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Package() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v8.5h12.5v-8.5l-3.5-4h-5.5zm6.25-4v3.5m-5.75.5h11.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pulse() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8.25h2.5l2-4.5l3.5 8.5l2-4h2.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Robot() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 5.75h12.5v7.5H1.75z"/><path d="M10.75 8.75v1.5m-5.5-1.5v1.5m-.5-7.5l3.25 3l3.25-3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Skull() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 11.25h3v3h6.5v-3h3s1-9.5-6.25-9.5s-6.25 9.5-6.25 9.5z"/><circle cx="5.25" cy="7.75" r=".5" fill="currentColor"/><circle cx="10.75" cy="7.75" r=".5" fill="currentColor"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChartLine() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 11.25l2.5-4.5l2.5 2.5l3.5-6m-11.5-1.5v12.5h12.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Key() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 1.75a4.25 4.25 0 0 0-4.104 5.354L1.75 11.25v3h3v-1.5h1.5v-1.5h1.5L8.9 10.1c.359.098.728.148 1.1.15a4.25 4.25 0 0 0 0-8.5z"/><circle cx="10.75" cy="5.25" r=".5" fill="currentColor"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Lightbulb() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 14.25h2.5M8 1.75c-2.75 0-4.25 2-4.25 4s2 2.5 2 4.5v1h4.5v-1c0-2 2-2.5 2-4.5s-1.5-4-4.25-4z"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Minus() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.25 7.75H2.75"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChartBar() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75v12.5h12.5m-9-3v-2.5m4 2.5v-5.5m4 5.5v-8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LayoutStackH() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12M8 8v4.75"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LinkExternal() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 2.75h-5.5v10.5h10.5v-5.5m0-5l-5.5 5.5m3-6.5h3.5v3.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MediaBack() *handlebars.ElementSVG {
	return handlebars.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.25 13.25L4.75 8l8.5-5.25zm-11.5-9.5v8.5"/>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Folders() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.75 2.25v8h9.5v-6.5h-5l-1.5-1.5z"/><path d="M4.75 5.25h-3v8h9.5v-3"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Info() *handlebars.ElementSVG {
	return handlebars.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8 5.25v0m0 6v-3.5"/></g>`).
		Custom("viewBox", "0,0,0,0").
		Custom("width", "1em").
		Custom("height", "1em")
}
