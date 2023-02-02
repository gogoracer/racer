/* cSpell:disable */
package quill

import "github.com/gogoracer/racer/pkg/goggles"

func Activity() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16h6l4-11l6 22l4-11h6"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Add() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 25V7m-9 9h18"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Alarm() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M25.9 7.1a1 1 0 1 0 1.414-1.415L25.899 7.1Zm-5.76-5.557a1 1 0 1 0-.517 1.932l.517-1.932Zm-7.764 1.932a1 1 0 1 0-.518-1.932l.518 1.932Zm-7.69 2.21A1 1 0 1 0 6.1 7.098L4.686 5.685ZM17 9a1 1 0 1 0-2 0h2Zm-1 8l.707.707A1 1 0 0 0 17 17h-1Zm-4.707 3.293a1 1 0 0 0 1.414 1.414l-1.414-1.414ZM26 17c0 5.523-4.477 10-10 10v2c6.627 0 12-5.373 12-12h-2ZM16 27c-5.523 0-10-4.477-10-10H4c0 6.627 5.373 12 12 12v-2ZM6 17c0-5.523 4.477-10 10-10V5C9.373 5 4 10.373 4 17h2ZM16 7c5.523 0 10 4.477 10 10h2c0-6.627-5.373-12-12-12v2Zm11.314-1.315a16 16 0 0 0-7.173-4.14l-.517 1.93A14 14 0 0 1 25.9 7.1l1.414-1.414Zm-15.455-4.14a16 16 0 0 0-7.173 4.14L6.101 7.1a14 14 0 0 1 6.276-3.623l-.518-1.932ZM15 9v8h2V9h-2Zm.293 7.293l-4 4l1.414 1.414l4-4l-1.414-1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Alt() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 9h5.485a1 1 0 0 1 .814.419l9.402 13.162a1 1 0 0 0 .814.419H27M21 9h6"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowDown() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M15 26a1 1 0 1 0 2 0h-2Zm2-21a1 1 0 1 0-2 0h2Zm-1 22l-.707.707a1 1 0 0 0 1.414 0L16 27Zm-7.293-8.707a1 1 0 0 0-1.414 1.414l1.414-1.414Zm16 1.414a1 1 0 0 0-1.414-1.414l1.414 1.414ZM17 26V5h-2v21h2Zm-.293.293l-8-8l-1.414 1.414l8 8l1.414-1.414Zm0 1.414l8-8l-1.414-1.414l-8 8l1.414 1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowLeft() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M6 15a1 1 0 1 0 0 2v-2Zm21 2a1 1 0 1 0 0-2v2ZM5 16l-.707-.707a1 1 0 0 0 0 1.414L5 16Zm8.707-7.293a1 1 0 0 0-1.414-1.414l1.414 1.414Zm-1.414 16a1 1 0 0 0 1.414-1.414l-1.414 1.414ZM6 17h21v-2H6v2Zm-.293-.293l8-8l-1.414-1.414l-8 8l1.414 1.414Zm-1.414 0l8 8l1.414-1.414l-8-8l-1.414 1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowRight() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M5 15a1 1 0 1 0 0 2v-2Zm21 2a1 1 0 1 0 0-2v2Zm1-1l.707.707a1 1 0 0 0 0-1.414L27 16Zm-7.293-8.707a1 1 0 1 0-1.414 1.414l1.414-1.414Zm-1.414 16a1 1 0 0 0 1.414 1.414l-1.414-1.414ZM5 17h21v-2H5v2Zm22.707-1.707l-8-8l-1.414 1.414l8 8l1.414-1.414Zm-1.414 0l-8 8l1.414 1.414l8-8l-1.414-1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ArrowUp() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M15 27a1 1 0 1 0 2 0h-2Zm2-21a1 1 0 1 0-2 0h2Zm-1-1l.707-.707a1 1 0 0 0-1.414 0L16 5Zm7.293 8.707a1 1 0 0 0 1.414-1.414l-1.414 1.414Zm-16-1.414a1 1 0 1 0 1.414 1.414l-1.414-1.414ZM17 27V6h-2v21h2ZM15.293 5.707l8 8l1.414-1.414l-8-8l-1.414 1.414Zm0-1.414l-8 8l1.414 1.414l8-8l-1.414-1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func At() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m23.363 5.286l.567-.824l-.567.824Zm-9.228-2.151l.143.99l-.143-.99ZM5.897 7.819l-.777-.63l.777.63Zm-2.87 9.031l-.997.066l.998-.066Zm22.34 9.554a1 1 0 1 0-1.337-1.486l1.338 1.486ZM22 10.5a1 1 0 1 0-2 0h2Zm-.744 9.896l.8-.6l-.8.6ZM24.5 22v1v-1ZM20 16a5 5 0 0 1-5 5v2a7 7 0 0 0 7-7h-2Zm-5 5a5 5 0 0 1-5-5H8a7 7 0 0 0 7 7v-2Zm-5-5a5 5 0 0 1 5-5V9a7 7 0 0 0-7 7h2Zm5-5a5 5 0 0 1 5 5h2a7 7 0 0 0-7-7v2Zm14.657 1.919c-.735-3.386-2.827-6.463-5.727-8.457l-1.133 1.649c2.491 1.712 4.281 4.356 4.905 7.232l1.955-.424ZM23.93 4.462a14 14 0 0 0-9.939-2.317l.287 1.98a12 12 0 0 1 8.519 1.986l1.133-1.649Zm-9.94-2.317A14 14 0 0 0 5.12 7.19l1.554 1.258a12 12 0 0 1 7.604-4.324l-.287-1.98ZM5.12 7.19a14 14 0 0 0-3.09 9.726l1.996-.131a12 12 0 0 1 2.648-8.337L5.12 7.19Zm-3.09 9.726a14 14 0 0 0 4.333 9.24l1.377-1.451a12 12 0 0 1-3.714-7.92l-1.996.13Zm4.333 9.24a14 14 0 0 0 9.454 3.843l.026-2a12 12 0 0 1-8.103-3.294l-1.377 1.45Zm9.454 3.843a14 14 0 0 0 9.55-3.595l-1.337-1.486a12 12 0 0 1-8.187 3.081l-.026 2ZM20 10.5v9.028h2V10.5h-2Zm.456 10.495c.54.721 1.823 2.005 4.044 2.005v-2a2.97 2.97 0 0 1-2.444-1.204l-1.6 1.2ZM20 19.529c0 .449.1.992.456 1.468l1.6-1.2c-.016-.022-.056-.096-.056-.268h-2ZM24.5 23c1.357 0 2.465-.44 3.314-1.2c.83-.744 1.354-1.737 1.673-2.764c.632-2.034.535-4.434.17-6.117l-1.955.424c.314 1.447.378 3.482-.125 5.1c-.248.8-.616 1.436-1.097 1.866c-.462.414-1.087.691-1.98.691v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Attachment() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.49 19.182l7.778-7.778c.976-.976 2.382-1.153 3.359-.177c.976.976.8 2.383-.177 3.359l-8.132 8.132c-1.414 1.414-4.243 1.414-6.01-.354c-1.768-1.768-1.768-4.596-.354-6.01l8.132-8.132a6.5 6.5 0 0 1 9.192 9.192l-4.596 4.596"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Breather() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25.192 25.192A13 13 0 1 0 6.808 6.808a13 13 0 0 0 18.384 18.384Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Broadcast() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 20.66a10 10 0 0 0 0-17.32M11 20.66a10 10 0 0 1 0-17.32m8 13.856a6 6 0 0 0 0-10.392m-6 10.392a6 6 0 0 1 0-10.392M13.09 26L16 18l2.91 8m-5.82 0L12 29m1.09-3h5.82M20 29l-1.09-3M18 12a2 2 0 1 0-4 0a2 2 0 0 0 4 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Calendar() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M14.5 15.5a1 1 0 1 0 0 2v-2Zm2 1h1a1 1 0 0 0-1-1v1Zm-1 6.5a1 1 0 1 0 2 0h-2ZM5 11a1 1 0 1 0 0 2v-2Zm22 2a1 1 0 1 0 0-2v2Zm-7-5a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2ZM10 8a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2Zm2.5 13.5h2v-2h-2v2Zm1-1V23h2v-6.5h-2ZM7 7h18V5H7v2Zm19 1v18h2V8h-2Zm-1 19H7v2h18v-2ZM6 26V8H4v18h2Zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm19-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 7a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM7 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm-2 8h22v-2H5v2Zm17-5V4h-2v4h2ZM12 8V4h-2v4h2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CalendarAdd() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M15 23.5a1 1 0 1 0 2 0h-2Zm2-7a1 1 0 1 0-2 0h2ZM12.5 19a1 1 0 1 0 0 2v-2Zm7 2a1 1 0 1 0 0-2v2ZM5 11a1 1 0 1 0 0 2v-2Zm22 2a1 1 0 1 0 0-2v2Zm-7-5a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2ZM10 8a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2Zm5 19.5V20h-2v3.5h2Zm0-3.5v-3.5h-2V20h2Zm-4.5 1H16v-2h-3.5v2Zm3.5 0h3.5v-2H16v2ZM7 7h18V5H7v2Zm19 1v18h2V8h-2Zm-1 19H7v2h18v-2ZM6 26V8H4v18h2Zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm19-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 7a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM7 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm-2 8h22v-2H5v2Zm17-5V4h-2v4h2ZM12 8V4h-2v4h2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CalendarMore() *goggles.ElementSVG {
	return goggles.SVG(`<defs><path id="quillCalendarMore0" fill-rule="evenodd" d="M12.5 19.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM21 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/></defs><mask id="quillCalendarMore1" fill="#fff"><use href="#quillCalendarMore0" fill-rule="evenodd" clip-rule="evenodd"/></mask><g fill="currentColor"><path d="M5 11a1 1 0 1 0 0 2v-2Zm22 2a1 1 0 1 0 0-2v2Zm-7-5a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2ZM10 8a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2ZM7 7h18V5H7v2Zm19 1v18h2V8h-2Zm-1 19H7v2h18v-2ZM6 26V8H4v18h2Zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm19-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 7a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM7 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm-2 8h22v-2H5v2Zm17-5V4h-2v4h2ZM12 8V4h-2v4h2Z"/><use href="#quillCalendarMore0" fill-rule="evenodd" clip-rule="evenodd"/><path d="M11 22a2.5 2.5 0 0 0 2.5-2.5h-2a.5.5 0 0 1-.5.5v2Zm-2.5-2.5A2.5 2.5 0 0 0 11 22v-2a.5.5 0 0 1-.5-.5h-2ZM11 17a2.5 2.5 0 0 0-2.5 2.5h2a.5.5 0 0 1 .5-.5v-2Zm2.5 2.5A2.5 2.5 0 0 0 11 17v2a.5.5 0 0 1 .5.5h2ZM16 22a2.5 2.5 0 0 0 2.5-2.5h-2a.5.5 0 0 1-.5.5v2Zm-2.5-2.5A2.5 2.5 0 0 0 16 22v-2a.5.5 0 0 1-.5-.5h-2ZM16 17a2.5 2.5 0 0 0-2.5 2.5h2a.5.5 0 0 1 .5-.5v-2Zm2.5 2.5A2.5 2.5 0 0 0 16 17v2a.5.5 0 0 1 .5.5h2Zm3 0a.5.5 0 0 1-.5.5v2a2.5 2.5 0 0 0 2.5-2.5h-2ZM21 19a.5.5 0 0 1 .5.5h2A2.5 2.5 0 0 0 21 17v2Zm-.5.5a.5.5 0 0 1 .5-.5v-2a2.5 2.5 0 0 0-2.5 2.5h2Zm.5.5a.5.5 0 0 1-.5-.5h-2A2.5 2.5 0 0 0 21 22v-2Z" mask="url(#quillCalendarMore1)"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CalendarSomeday() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="currentColor"><path stroke="currentColor" d="M16.75 23.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path d="M16 19.462v-1a1 1 0 0 0-1 1h1Zm1.962-2.39l-.981.195l.98-.196Zm-2.727-1.458l.382.924l-.382-.924ZM13 17.462a1 1 0 1 0 2 0h-2Zm2 3.038a1 1 0 1 0 2 0h-2ZM5 11a1 1 0 1 0 0 2v-2Zm22 2a1 1 0 1 0 0-2v2Zm-7-5a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2ZM10 8a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2Zm4 16.462a3 3 0 0 0 1.667-.506l-1.111-1.663a1 1 0 0 1-.556.169v2Zm1.667-.506a3 3 0 0 0 1.105-1.346l-1.848-.766a1 1 0 0 1-.368.449l1.11 1.663Zm1.105-1.346a3 3 0 0 0 .17-1.734l-1.961.39a1 1 0 0 1-.057.578l1.848.766Zm.17-1.734a3 3 0 0 0-.82-1.536l-1.415 1.414a1 1 0 0 1 .274.512l1.961-.39Zm-.82-1.536a3 3 0 0 0-1.537-.82l-.39 1.96a1 1 0 0 1 .512.274l1.414-1.414Zm-1.537-.82a3 3 0 0 0-1.733.17l.765 1.848a1 1 0 0 1 .578-.057l.39-1.962Zm-1.733.17a3 3 0 0 0-1.346 1.105l1.662 1.111a1 1 0 0 1 .45-.368l-.766-1.848Zm-1.346 1.105A3 3 0 0 0 13 17.462h2a1 1 0 0 1 .168-.556l-1.662-1.111ZM15 19.462V20.5h2v-1.038h-2ZM7 7h18V5H7v2Zm19 1v18h2V8h-2Zm-1 19H7v2h18v-2ZM6 26V8H4v18h2Zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm19-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 7a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM7 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm-2 8h22v-2H5v2Zm17-5V4h-2v4h2ZM12 8V4h-2v4h2Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Catchup() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 27H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6m0 14h2a2 2 0 0 0 2-2V15a2 2 0 0 0-2-2h-2m0 14V13M9 10h4m-4 7h10M9 21h6"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Chat() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M25 5H7a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h11l6 4v-4h1a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4Z"/><path d="M10 15a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm6 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm6 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Checkmark() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.5 17l6 6l13-13"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CheckmarkDouble() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 17l5 5l12-12m-5 10l2 2l12-12"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CheckmarkTodo() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" fill-rule="evenodd" d="M26.207 9.293a1 1 0 0 1 0 1.414l-1.625 1.625a1 1 0 0 1-1.414-1.414l1.625-1.625a1 1 0 0 1 1.414 0Zm-4.875 4.875a1 1 0 0 1 0 1.414l-3.25 3.25a1 1 0 0 1-1.414-1.414l3.25-3.25a1 1 0 0 1 1.414 0Zm-15.54 2.125a1 1 0 0 1 1.415 0l1.5 1.5a1 1 0 1 1-1.414 1.414l-1.5-1.5a1 1 0 0 1 0-1.414Zm9.04 4.375a1 1 0 0 1 0 1.414l-1.625 1.625a1 1 0 0 1-1.414 0l-1.5-1.5a1 1 0 1 1 1.414-1.414l.793.793l.918-.918a1 1 0 0 1 1.414 0Z" clip-rule="evenodd"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronDown() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 12l10 10l10-10"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronLeft() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6L10 16l10 10"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronRight() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 26l10-10L12 6"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func ChevronUp() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M26 20L16 10L6 20"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Clock() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8v8l4 4m9-4c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cloudia() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor"><path fill="currentColor" d="M11.5 18a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm10 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 21h-2m12.708-4.717a9 9 0 1 0-17.66-3.218A7.001 7.001 0 0 0 10 27h13a6 6 0 0 0 3.708-10.717Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Cog() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m12.874 6.498l.313.95l-.313-.95Zm.342-.362l-.97-.242l.97.242Zm-2.222.921l-.514.857l.514-.857Zm.498.014l.451.893l-.451-.893ZM7.946 5.67l-.707-.707l.707.707Zm.61-.075l.515-.857l-.514.857ZM5.595 8.557l-.857.514l.857-.514Zm.075-.611l.708.707l-.708-.707Zm1.402 3.546l-.892-.451l.892.451Zm-.014-.498l.857-.514l-.857.514Zm-.92 2.222l-.243-.97l.242.97Zm.361-.341l.95.312l-.95-.313Zm-3.12 1.03l.243.97l-.242-.97Zm0 4.19l.243-.97l-.242.97Zm3.12 1.03l-.95.313l.95-.313Zm-.362-.341l-.242.97l.242-.97Zm.921 2.222l.857.514l-.857-.514Zm.014-.498l.893-.451l-.893.451ZM5.67 24.054l-.707.707l.707-.707Zm-.075-.61l-.857-.515l.857.514Zm2.963 2.962l-.515-.858l.515.858Zm-.611-.075l.707-.708l-.707.708Zm3.546-1.402l.451-.893l-.451.893Zm-.498.014l.515.857l-.515-.857Zm1.914.78l-.949.317l.949-.316Zm-.315-.318l-.34.94l.34-.94Zm1.293 3.253l.949-.316l-.949.316Zm4.228 0l-.949-.316l.949.316Zm1.293-3.253l.34.94l-.34-.94Zm-.315.319l.949.316l-.949-.316Zm1.597-.662l.447-.895l-.447.895Zm-.449-.003l-.424-.905l.424.905Zm3.216 1.386l-.447.895l.447-.895Zm2.99-2.99l-.895.448l.894-.447ZM25.06 20.24l.905.425l-.905-.425Zm.003.449l.894-.447l-.894.447Zm.662-1.597l.316.949l-.316-.949Zm-.32.315l-.94-.34l.94.34Zm3.254-1.293l-.316-.949l.316.949Zm0-4.228l.316-.949l-.316.949Zm-3.253-1.293l.94-.34l-.94.34Zm.319.315l-.317.949l.317-.95Zm-.662-1.597l.894.447l-.894-.447Zm-.003.449l.905-.425l-.905.425Zm1.293-3.793l-.707.707l.707-.707Zm.093.577l-.894-.447l.894.447Zm-2.99-2.99l-.446-.894l.447.895Zm.578.094l.707-.707l-.707.707ZM20.24 6.941l.425-.906l-.425.906Zm.449-.003l.447.894l-.447-.894Zm-1.905-.802l.97-.242l-.97.242Zm.341.362l-.312.95l.313-.95Zm-5.22-3.12l.97.243l-.97-.242Zm4.19 0l-.97.243l.97-.242ZM17.61 2h-3.22v2h3.22V2Zm2.144 3.894l-.69-2.758l-1.94.485l.69 2.758l1.94-.485Zm-.94 1.554c.343.113.678.246 1.002.398l.849-1.81a10.933 10.933 0 0 0-1.227-.488l-.625 1.9Zm4.195-2.788l-2.767 1.384l.894 1.788l2.767-1.383l-.894-1.789Zm4.05 2.6L24.74 4.94l-1.414 1.414l2.319 2.32l1.414-1.415Zm-1.103 4.498l1.384-2.767l-1.789-.894l-1.383 2.767l1.788.894Zm-1.802.426c.114.244.218.494.31.75l1.88-.682a10.95 10.95 0 0 0-.38-.917l-1.81.85Zm4.82.753l-2.934-.978l-.633 1.898l2.935.978l.632-1.898ZM30 17.64v-3.28h-2v3.28h2Zm-3.96 2.4l2.934-.977l-.632-1.898l-2.934.978l.632 1.898Zm-1.576-.974a8.94 8.94 0 0 1-.31.75l1.81.849c.14-.299.268-.605.38-.917l-1.88-.682Zm2.876 3.943l-1.384-2.767l-1.788.894l1.383 2.767l1.789-.894Zm-2.6 4.05l2.319-2.319l-1.414-1.414l-2.32 2.319l1.415 1.414Zm-4.498-1.103l2.767 1.384l.894-1.789l-2.767-1.383l-.894 1.788Zm-.426-1.802a8.94 8.94 0 0 1-.75.31l.682 1.88c.312-.113.618-.24.917-.38l-.85-1.81Zm-.753 4.82l.978-2.934l-1.898-.633l-.978 2.935l1.898.632ZM14.36 30h3.28v-2h-3.28v2Zm-2.4-3.96l.977 2.934l1.898-.632l-.978-2.934l-1.898.632Zm.974-1.576a8.922 8.922 0 0 1-.99-.428l-.903 1.785c.39.198.795.373 1.211.524l.682-1.88Zm-3.863 2.8L11.51 25.8l-1.03-1.715l-2.437 1.462l1.03 1.715ZM4.962 24.76l2.277 2.277l1.414-1.415l-2.276-2.276l-1.415 1.414Zm1.237-4.27l-1.462 2.44l1.715 1.029l1.462-2.438l-1.714-1.03Zm1.765-.434a8.938 8.938 0 0 1-.516-1.244l-1.9.625c.173.526.384 1.034.631 1.521l1.785-.902Zm-4.828-.992l2.758.69l.485-1.941l-2.758-.69l-.485 1.94ZM2 14.39v3.22h2v-3.22H2Zm3.894-2.144l-2.758.69l.485 1.94l2.758-.69l-.485-1.94Zm1.554.94c.141-.429.314-.844.516-1.243l-1.785-.902c-.247.487-.458.995-.63 1.521l1.9.625ZM4.737 9.072l1.462 2.438l1.715-1.03l-1.462-2.437l-1.715 1.03Zm2.502-4.109L4.962 7.24l1.415 1.414l2.276-2.276l-1.414-1.415Zm4.27 1.237L9.07 4.737L8.042 6.452l2.438 1.462L11.509 6.2Zm.434 1.765a8.938 8.938 0 0 1 1.244-.516l-.625-1.9c-.526.173-1.034.384-1.521.631l.902 1.785Zm.992-4.828l-.69 2.758l1.941.485l.69-2.758l-1.94-.485Zm.252 4.312c.479-.157.87-.55 1-1.07l-1.941-.484a.485.485 0 0 1 .316-.346l.625 1.9Zm-2.707.466c.46.276 1.013.277 1.463.05l-.902-1.785a.485.485 0 0 1 .468.02l-1.03 1.715ZM8.653 6.377a.5.5 0 0 1-.61.075L9.07 4.737a1.5 1.5 0 0 0-1.832.225l1.414 1.415ZM6.452 8.042a.5.5 0 0 1-.075.61L4.962 7.24a1.5 1.5 0 0 0-.225 1.832l1.715-1.029Zm1.512 3.901a1.515 1.515 0 0 0-.05-1.463L6.2 11.509a.485.485 0 0 1-.02-.468l1.785.902Zm-1.585 2.243c.52-.13.912-.52 1.07-1l-1.9-.624a.485.485 0 0 1 .345-.316l.485 1.94ZM4 14.39a.5.5 0 0 1-.379.486l-.485-1.94A1.5 1.5 0 0 0 2 14.39h2Zm-.379 2.735A.5.5 0 0 1 4 17.61H2a1.5 1.5 0 0 0 1.136 1.455l.485-1.94Zm3.827 1.688a1.515 1.515 0 0 0-1.07-1l-.484 1.941a.485.485 0 0 1-.346-.316l1.9-.625Zm.466 2.707c.276-.46.277-1.013.05-1.463l-1.785.902a.485.485 0 0 1 .02-.468l1.715 1.03Zm-1.537 1.827a.5.5 0 0 1 .075.61L4.737 22.93a1.5 1.5 0 0 0 .225 1.832l1.415-1.414Zm1.665 2.201a.5.5 0 0 1 .61.075L7.24 27.038a1.5 1.5 0 0 0 1.832.225l-1.029-1.715Zm3.901-1.512a1.515 1.515 0 0 0-1.463.05l1.029 1.715a.485.485 0 0 1-.468.02l.902-1.785Zm1.914 1.371a1.519 1.519 0 0 0-.923-.943l-.682 1.88a.485.485 0 0 1-.293-.304l1.898-.633ZM14.36 28a.5.5 0 0 1 .475.342l-1.898.632A1.5 1.5 0 0 0 14.36 30v-2Zm2.805.342A.5.5 0 0 1 17.64 28v2a1.5 1.5 0 0 0 1.423-1.026l-1.898-.632Zm1.901-3.878a1.515 1.515 0 0 0-.923.944l1.898.632a.485.485 0 0 1-.293.305l-.682-1.88Zm2.07-.296a1.514 1.514 0 0 0-1.32-.014l.849 1.81a.485.485 0 0 1-.423-.008l.894-1.788Zm2.19 1.477a.5.5 0 0 1 .577-.094l-.894 1.789a1.5 1.5 0 0 0 1.731-.281l-1.414-1.414Zm2.225-1.742a.5.5 0 0 1 .094-.577l1.414 1.414a1.5 1.5 0 0 0 .28-1.731l-1.788.894Zm-1.397-4.087c-.191.407-.2.894.014 1.32l1.788-.894a.485.485 0 0 1 .009.423l-1.811-.85Zm1.253-1.673c-.451.151-.79.501-.943.923l1.88.682a.485.485 0 0 1-.304.293l-.633-1.898ZM28 17.64a.5.5 0 0 1 .342-.475l.632 1.898A1.5 1.5 0 0 0 30 17.64h-2Zm.342-2.805A.5.5 0 0 1 28 14.36h2a1.5 1.5 0 0 0-1.026-1.423l-.632 1.898Zm-3.878-1.901c.154.422.492.772.944.923l.632-1.898a.485.485 0 0 1 .305.293l-1.88.682Zm-.296-2.07a1.514 1.514 0 0 0-.014 1.32l1.81-.849a.485.485 0 0 1-.008.423l-1.788-.894Zm1.477-2.19a.5.5 0 0 1-.094-.577l1.789.894a1.5 1.5 0 0 0-.281-1.731l-1.414 1.414ZM23.903 6.45a.5.5 0 0 1-.577-.094l1.414-1.414a1.5 1.5 0 0 0-1.731-.28l.894 1.788Zm-4.087 1.397c.407.191.894.2 1.32-.014l-.894-1.788a.486.486 0 0 1 .423-.009l-.85 1.811ZM17.814 6.38c.13.52.52.912 1 1.07l.624-1.9a.485.485 0 0 1 .316.345l-1.94.485ZM14.39 2a1.5 1.5 0 0 0-1.455 1.136l1.94.485A.5.5 0 0 1 14.39 4V2Zm3.22 2a.5.5 0 0 1-.485-.379l1.94-.485A1.5 1.5 0 0 0 17.61 2v2ZM20 16a4 4 0 0 1-4 4v2a6 6 0 0 0 6-6h-2Zm-4 4a4 4 0 0 1-4-4h-2a6 6 0 0 0 6 6v-2Zm-4-4a4 4 0 0 1 4-4v-2a6 6 0 0 0-6 6h2Zm4-4a4 4 0 0 1 4 4h2a6 6 0 0 0-6-6v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func CogAlt() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M19.19 3.757A1 1 0 0 0 18.22 3h-4.44a1 1 0 0 0-.97.757l-.66 2.644a1 1 0 0 1-.47.623l-1.291.747a1 1 0 0 1-.776.095l-2.62-.75a1 1 0 0 0-1.142.462l-2.219 3.844a1 1 0 0 0 .171 1.219l1.96 1.895a1 1 0 0 1 .305.719v1.49a1 1 0 0 1-.305.72l-1.96 1.894a1 1 0 0 0-.17 1.22l2.218 3.843a1 1 0 0 0 1.141.461l2.61-.746a1 1 0 0 1 .793.106l.963.584a1 1 0 0 1 .43.54l.984 2.95a1 1 0 0 0 .949.683h4.558a1 1 0 0 0 .949-.684l.982-2.947a1 1 0 0 1 .435-.542l.982-.587a1 1 0 0 1 .79-.103l2.59.745a1 1 0 0 0 1.142-.461l2.222-3.848a1 1 0 0 0-.166-1.214l-1.933-1.896a1 1 0 0 1-.3-.713v-1.5a1 1 0 0 1 .3-.713l1.933-1.896a1 1 0 0 0 .166-1.214l-2.222-3.848a1 1 0 0 0-1.142-.46l-2.6.747a1 1 0 0 1-.774-.093l-1.31-.75a1 1 0 0 1-.474-.625l-.66-2.64Z"/><path d="M21 16a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Collapse() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m23 26l-7-7l-7 7M9 6l7 7l7-7"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Command() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-width="2" d="M19 19v4a4 4 0 1 0 4-4h-4Zm0 0v-6m0 6h-6m6-6V9a4 4 0 1 1 4 4h-4Zm0 0h-6m-4 0h-.007m0 0A4 4 0 1 1 13 9v4m-4.007 0H13m0 0v6m0 0v4a4 4 0 1 1-4-4h4Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Compose() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27 5L17 15m0-10H8a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-9"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Creditcard() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M3 23h1h-1Zm8-4a1 1 0 1 0 0 2v-2Zm8 2a1 1 0 1 0 0-2v2ZM3 12a1 1 0 1 0 0 2v-2Zm26 2a1 1 0 1 0 0-2v2Zm-6 5a1 1 0 1 0 0 2v-2Zm2 2a1 1 0 1 0 0-2v2Zm3-12v14h2V9h-2ZM5 24c-.175 0-.433-.098-.668-.332C4.097 23.433 4 23.175 4 23H2c0 .825.403 1.567.918 2.082C3.433 25.597 4.175 26 5 26v-2Zm23-1c0 .175-.098.433-.332.668c-.235.235-.493.332-.668.332v2c.825 0 1.567-.402 2.082-.918c.515-.515.918-1.257.918-2.082h-2ZM4 23V9H2v14h2ZM4 9c0-.175.097-.433.332-.668C4.567 8.097 4.825 8 5 8V6c-.825 0-1.567.403-2.082.918C2.403 7.433 2 8.175 2 9h2Zm1-1h22V6H5v2Zm22 0c.175 0 .433.097.668.332c.235.235.332.493.332.668h2c0-.825-.402-1.567-.918-2.082C28.567 6.403 27.825 6 27 6v2ZM11 21h8v-2h-8v2Zm-8-7h26v-2H3v2Zm20 7h2v-2h-2v2Zm4 3H5v2h22v-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Desktop() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 27h4m-4 0l.5-4h3l.5 4m-4 0h-3m7 0h3M5 5h22a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Discard() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.5 23l13-13.5M29 16c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Download() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M15 22a1 1 0 1 0 2 0h-2Zm2-17a1 1 0 1 0-2 0h2Zm-1 18l-.707.707a1 1 0 0 0 1.414 0L16 23Zm-6.293-7.707a1 1 0 0 0-1.414 1.414l1.414-1.414Zm14 1.414a1 1 0 0 0-1.414-1.414l1.414 1.414ZM9 26a1 1 0 1 0 0 2v-2Zm14 2a1 1 0 1 0 0-2v2Zm-6-6V5h-2v17h2Zm-.293.293l-7-7l-1.414 1.414l7 7l1.414-1.414Zm0 1.414l7-7l-1.414-1.414l-7 7l1.414 1.414ZM9 28h14v-2H9v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Earth() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-width="2" d="M21.5 4.5c0 1.167-1.735 1.5-3 1.5c-5.217 0-10.705 3.48-11.421 8.004C6.992 14.549 6.552 15 6 15H3m1 5.5c.785.262 2.126 1.285 3.44 1.517c.57.101 1.153.464 1.299 1.023c.303 1.16.548 1.992-.239 3.96m8.725-1.707c-1 0-5-2.2-5-9c0-5.5 3.5-7 8.5-7c3 0 5 1.5 5 5s-3.5 2.707-5 4.207s0 6.793-3.5 6.793ZM29 16c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Escape() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 23L23 9m0 14L9 9"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Expand() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 20l7 7l7-7m0-8l-7-7l-7 7"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Eye() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M29 16c0 3-5.82 9-13 9S3 19 3 16s5.82-9 13-9s13 6 13 9Z"/><path d="M21 16a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func EyeClosed() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M22 16a1 1 0 1 0-2 0h2Zm-6 4a1 1 0 1 0 0 2v-2Zm-6-4a1 1 0 1 0 2 0h-2Zm6-4a1 1 0 1 0 0-2v2Zm-2.776 11.68a1 1 0 0 0-.448 1.95l.448-1.95Zm6.017-15.244a1 1 0 0 0 .518-1.932l-.518 1.932Zm7.358 1.822a1 1 0 1 0-1.34 1.484l1.34-1.484ZM5.325 21.673a1 1 0 0 0 1.35-1.475l-1.35 1.475Zm.968 2.62a1 1 0 1 0 1.414 1.414l-1.414-1.414ZM25.707 7.707a1 1 0 0 0-1.414-1.414l1.414 1.414ZM28 16c0 .464-.243 1.203-.853 2.116c-.593.888-1.471 1.845-2.578 2.727C22.351 22.611 19.314 24 16 24v2c3.866 0 7.329-1.611 9.816-3.593c1.246-.993 2.271-2.099 2.994-3.18C29.515 18.172 30 17.037 30 16h-2ZM4 16c0-.464.243-1.203.853-2.116c.593-.888 1.471-1.845 2.578-2.727C9.649 9.389 12.686 8 16 8V6c-3.866 0-7.329 1.611-9.816 3.593c-1.246.993-2.271 2.098-2.994 3.18C2.485 13.828 2 14.963 2 16h2Zm16 0a4 4 0 0 1-4 4v2a6 6 0 0 0 6-6h-2Zm-8 0a4 4 0 0 1 4-4v-2a6 6 0 0 0-6 6h2Zm4 8c-.952 0-1.881-.114-2.776-.32l-.448 1.95c1.031.236 2.111.37 3.224.37v-2Zm0-16c1.118 0 2.205.158 3.24.436l.52-1.932A14.489 14.489 0 0 0 16 6v2Zm9.258 3.742c.899.812 1.6 1.655 2.071 2.427c.482.79.671 1.423.671 1.831h2c0-.928-.389-1.93-.963-2.872c-.586-.962-1.42-1.95-2.438-2.87l-1.34 1.484ZM6.675 20.198c-.878-.804-1.563-1.636-2.021-2.395C4.184 17.024 4 16.403 4 16H2c0 .917.38 1.906.941 2.836c.573.95 1.389 1.926 2.384 2.837l1.35-1.475Zm1.032 5.51l18-18l-1.414-1.415l-18 18l1.414 1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Filter() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22M7 13h18M9 19h14m-12 6h10"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Focus() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm9 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path fill="currentColor" d="M9.5 16a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Folder() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M28 11v13a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6c3 0 3 3 5 3h9.003C27.108 9 28 9.895 28 11Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderAdd() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M15 21a1 1 0 1 0 2 0h-2Zm2-8a1 1 0 1 0-2 0h2Zm-5 3a1 1 0 1 0 0 2v-2Zm8 2a1 1 0 1 0 0-2v2Zm9 6V11h-2v13h2ZM26.003 8H17v2h9.003V8ZM17 8c-.208 0-.36-.063-.552-.228c-.245-.21-.442-.477-.805-.912c-.327-.393-.755-.872-1.35-1.24C13.678 5.236 12.933 5 12 5v2c.568 0 .947.138 1.238.318c.311.194.571.465.869.822c.262.315.627.797 1.04 1.15c.463.398 1.06.71 1.853.71V8Zm-5-3H6v2h6V5ZM3 8v16h2V8H3Zm3 19h20v-2H6v2ZM6 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm23 6c0-1.655-1.338-3-2.997-3v2c.55 0 .997.446.997 1h2ZM3 24a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1H3Zm24 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm-10-3v-8h-2v8h2Zm-5-3h8v-2h-8v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderArchive() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M5 11H4a1 1 0 0 0 1 1v-1Zm2 0h1a1 1 0 0 0-1-1v1Zm18 0v-1a1 1 0 0 0-1 1h1Zm2 0v1a1 1 0 0 0 1-1h-1ZM6 10a1 1 0 1 0 0 2v-2Zm20 2a1 1 0 1 0 0-2v2Zm-14 3a1 1 0 1 0 0 2v-2Zm8 2a1 1 0 1 0 0-2v2ZM4 7v4h2V7H4Zm1 5h2v-2H5v2Zm1-1v14.002h2V11H6Zm20 14.002V11h-2v14.002h2ZM25 12h2v-2h-2v2Zm3-1V7h-2v4h2Zm-5 15H9v2h14v-2Zm2-22H7v2h18V4Zm3 3a3 3 0 0 0-3-3v2a1 1 0 0 1 1 1h2Zm-4 18.002A.999.999 0 0 1 23 26v2c1.656 0 3-1.34 3-2.998h-2Zm-18 0A2.999 2.999 0 0 0 9 28v-2c-.553 0-1-.447-1-.998H6ZM6 6.998A1 1 0 0 1 7 6V4a3 3 0 0 0-3 3h2ZM6 12h20v-2H6v2Zm6 5h8v-2h-8v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderDownload() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m16 21l-.707.707a1 1 0 0 0 1.414 0L16 21Zm1-8a1 1 0 1 0-2 0h2Zm3.207 5.207a1 1 0 0 0-1.414-1.414l1.414 1.414Zm-7-1.414a1 1 0 0 0-1.414 1.414l1.414-1.414ZM29 24V11h-2v13h2ZM26.003 8H17v2h9.003V8ZM17 8c-.208 0-.36-.063-.552-.228c-.245-.21-.442-.477-.805-.912c-.327-.393-.755-.872-1.35-1.24C13.678 5.236 12.933 5 12 5v2c.568 0 .947.138 1.238.318c.311.194.571.465.869.822c.262.315.627.797 1.04 1.15c.463.398 1.06.71 1.853.71V8Zm-5-3H6v2h6V5ZM3 8v16h2V8H3Zm3 19h20v-2H6v2ZM6 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm23 6c0-1.655-1.338-3-2.997-3v2c.55 0 .997.446.997 1h2ZM3 24a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1H3Zm24 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm-10-3v-8h-2v8h2Zm1.793-4.207l-3.5 3.5l1.414 1.414l3.5-3.5l-1.414-1.414Zm-2.086 3.5l-3.5-3.5l-1.414 1.414l3.5 3.5l1.414-1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderDrafts() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h8m-8 4h22M5 16h6m-6 4h12m-3-4h13m-6 4h6M5 24h2m4 0h4"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderList() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 15h8m-8 4h8m8 5V11c0-1.105-.892-2-1.997-2H17c-2 0-2-3-5-3H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderOpen() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 26V8a2 2 0 0 1 2-2h6c3 0 3 3 5 3h7a2 2 0 0 1 2 2v2M4 26l3.783-12.294A1 1 0 0 1 8.739 13H26M4 26h19.523a2 2 0 0 0 1.911-1.412l3.168-10.294A1 1 0 0 0 27.646 13H26"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderPut() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M16.168 9.555a1 1 0 0 0 1.664-1.11l-1.664 1.11ZM25 8a1 1 0 1 0 0 2V8Zm-4 10l-.707.707a1 1 0 0 0 1.414 0L21 18Zm1-12a1 1 0 1 0-2 0h2Zm3.207 9.207a1 1 0 0 0-1.414-1.414l1.414 1.414Zm-7-1.414a1 1 0 0 0-1.414 1.414l1.414-1.414ZM29 24V11h-2v13h2ZM17 9l.832-.555v-.001l-.002-.002a.164.164 0 0 1-.01-.014a3.047 3.047 0 0 0-.107-.15a9.757 9.757 0 0 0-1.437-1.537C15.313 5.915 13.838 5 12 5v2c1.162 0 2.187.585 2.974 1.26a7.753 7.753 0 0 1 1.13 1.205a3.88 3.88 0 0 1 .065.091c0-.001-.001-.001.831-.556Zm-5-4H6v2h6V5ZM3 8v16h2V8H3Zm3 19h20v-2H6v2ZM26 8h-1v2h1V8ZM6 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm23 6a3 3 0 0 0-3-3v2a1 1 0 0 1 1 1h2ZM3 24a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1H3Zm24 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm-5-6V6h-2v12h2Zm1.793-4.207l-3.5 3.5l1.414 1.414l3.5-3.5l-1.414-1.414Zm-2.086 3.5l-3.5-3.5l-1.414 1.414l3.5 3.5l1.414-1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderSpam() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.5 12.5c1.5 0 5.5 0 5.5 7M11 5c4 0 8.5 2.8 8.5 6c0 5-9.5 3-9.5 8c0 5.5 17-5 17 4c0 4-8.5 4-13 4c-12 0-11-10.5-3-10.5C15.4 12.1 12.5 7 11 5Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderTodo() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M12.707 16.293a1 1 0 0 0-1.414 1.414l1.414-1.414ZM15 20l-.707.707a1 1 0 0 0 1.414 0L15 20Zm6.207-4.793a1 1 0 0 0-1.414-1.414l1.414 1.414Zm-9.914 2.5l3 3l1.414-1.414l-3-3l-1.414 1.414Zm4.414 3l5.5-5.5l-1.414-1.414l-5.5 5.5l1.414 1.414ZM29 24V11h-2v13h2ZM26.003 8H17v2h9.003V8ZM17 8c-.208 0-.36-.063-.552-.228c-.245-.21-.442-.477-.805-.912c-.327-.393-.755-.872-1.35-1.24C13.678 5.236 12.933 5 12 5v2c.568 0 .947.138 1.238.318c.311.194.571.465.869.822c.262.315.627.797 1.04 1.15c.463.398 1.06.71 1.853.71V8Zm-5-3H6v2h6V5ZM3 8v16h2V8H3Zm3 19h20v-2H6v2ZM6 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm23 6c0-1.655-1.338-3-2.997-3v2c.55 0 .997.446.997 1h2ZM3 24a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1H3Zm24 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func FolderTrash() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m8 26l-.986.164a1 1 0 0 0 .092.283L8 26Zm16 0l.894.447a1 1 0 0 0 .092-.283L24 26Zm-4.793-8.793a1 1 0 0 0-1.414-1.414l1.414 1.414Zm-6.414 3.586a1 1 0 0 0 1.414 1.414l-1.414-1.414Zm1.414-5a1 1 0 0 0-1.414 1.414l1.414-1.414Zm3.586 6.414a1 1 0 0 0 1.414-1.414l-1.414 1.414ZM26 8c0-.115.073.018-.406.327c-.435.28-1.13.57-2.079.83C21.632 9.67 18.975 10 16 10v2c3.1 0 5.943-.342 8.041-.914c1.042-.284 1.958-.64 2.636-1.078C27.311 9.6 28 8.944 28 8h-2Zm-10 2c-2.975 0-5.632-.33-7.515-.843c-.949-.26-1.644-.55-2.079-.83C5.926 8.018 6 7.885 6 8H4c0 .944.689 1.6 1.323 2.008c.678.438 1.594.794 2.636 1.078c2.098.572 4.94.914 8.041.914v-2ZM6 8c0 .115-.073-.018.406-.327c.435-.28 1.13-.57 2.079-.83C10.368 6.33 13.025 6 16 6V4c-3.1 0-5.943.342-8.041.914c-1.042.284-1.958.64-2.636 1.078C4.689 6.4 4 7.056 4 8h2Zm10-2c2.975 0 5.632.33 7.515.843c.949.26 1.644.55 2.079.83c.48.309.406.442.406.327h2c0-.944-.689-1.6-1.323-2.008c-.678-.438-1.594-.794-2.636-1.078C21.943 4.342 19.101 4 16 4v2ZM4.014 8.164l3 18l1.972-.328l-3-18l-1.972.328Zm20.972 18l3-18l-1.972-.328l-3 18l1.972.328ZM8 26c-.894.447-.894.448-.893.45l.001.002a.858.858 0 0 0 .054.095c.025.04.056.087.095.139c.077.104.185.227.33.361c.292.27.73.577 1.382.867C10.264 28.489 12.415 29 16 29v-2c-3.415 0-5.264-.49-6.219-.914c-.474-.21-.723-.403-.837-.508a.742.742 0 0 1-.087-.092c-.003-.003 0 0 .007.012a.509.509 0 0 1 .026.045l.002.005l.001.002c0 .002.001.003-.893.45Zm8 3c3.585 0 5.736-.51 7.031-1.086c.651-.29 1.09-.597 1.382-.867a2.64 2.64 0 0 0 .33-.361a1.826 1.826 0 0 0 .146-.229l.003-.005l.001-.002c0-.002.001-.003-.893-.45a36.234 36.234 0 0 1-.893-.45l.001-.002l.002-.005l.005-.008a.47.47 0 0 1 .02-.037l.008-.012a.75.75 0 0 1-.087.092c-.114.105-.363.298-.837.508c-.955.425-2.804.914-6.219.914v2Zm1.793-13.207l-5 5l1.414 1.414l5-5l-1.414-1.414Zm-5 1.414l5 5l1.414-1.414l-5-5l-1.414 1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Forcebatch() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M24 14a1 1 0 1 0 0 2v-2ZM8 16a1 1 0 1 0 0-2v2Zm6 2a1 1 0 1 0 0 2v-2Zm4 2a1 1 0 1 0 0-2v2Zm-2-5l-.707.707a1 1 0 0 0 1.414 0L16 15Zm1-12a1 1 0 1 0-2 0h2Zm3.707 8.707a1 1 0 0 0-1.414-1.414l1.414 1.414Zm-8-1.414a1 1 0 1 0-1.414 1.414l1.414-1.414ZM24 16h1v-2h-1v2Zm2 1v8h2v-8h-2Zm-1 9H7v2h18v-2ZM6 25v-8H4v8h2Zm1-9h1v-2H7v2Zm-1 1a1 1 0 0 1 1-1v-2a3 3 0 0 0-3 3h2Zm1 9a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm19-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm-1-9a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2Zm-11 4h4v-2h-4v2Zm3-5V3h-2v12h2Zm-.293.707l4-4l-1.414-1.414l-4 4l1.414 1.414Zm0-1.414l-4-4l-1.414 1.414l4 4l1.414-1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Formatting() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.75 18L10 8L6.25 18m7.5 0L16 24m-2.25-6h-7.5M4 24l2.25-6M28 16c-.333.833-1.8 2-5 2s-3.5 2-3.5 3s.7 3 3.5 3s5-2.5 5-4.5m0-3.5v3.5m0-3.5c0-1.333-.8-4-4-4c-2.4 0-3.667 1.333-4 2m8 5.5V24"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Forward() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13S3 13 3 16s14 3 14 3v3.453c0 1.74 2.069 2.65 3.351 1.475l7.04-6.454a2 2 0 0 0 0-2.948l-7.04-6.454C19.07 6.896 17 7.806 17 9.546V13Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Fullscreen() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21v3.5a.5.5 0 0 0 .5.5H7M3 11V7.5a.5.5 0 0 1 .5-.5H7m18 0h3.5a.5.5 0 0 1 .5.5V11m0 10v3.5a.5.5 0 0 1-.5.5H25"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Gift() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M5 15H4a1 1 0 0 0 1 1v-1Zm22 0v1a1 1 0 0 0 1-1h-1ZM15 27a1 1 0 1 0 2 0h-2Zm7.5-23.5l.707-.707l-.707.707ZM7 8c-.825 0-1.567.403-2.082.918C4.403 9.433 4 10.175 4 11h2c0-.175.097-.433.332-.668c.235-.235.493-.332.668-.332V8Zm-3 3v4h2v-4H4Zm1 5h2v-2H5v2Zm1-1v10h2V15H6Zm0 10c0 .825.403 1.567.918 2.082C7.433 27.597 8.175 28 9 28v-2c-.175 0-.433-.097-.668-.332C8.097 25.433 8 25.175 8 25H6Zm3 3h14v-2H9v2Zm14 0c.825 0 1.567-.402 2.082-.918c.515-.515.918-1.257.918-2.082h-2c0 .175-.098.433-.332.668c-.235.235-.493.332-.668.332v2Zm3-3V15h-2v10h2Zm-1-9h2v-2h-2v2Zm3-1v-4h-2v4h2Zm0-4c0-.825-.402-1.567-.918-2.082C26.567 8.403 25.825 8 25 8v2c.175 0 .433.098.668.332c.235.235.332.493.332.668h2Zm-3-3H7v2h18V8ZM7 16h18v-2H7v2Zm8-7v18h2V9h-2Zm1.905-.426c-.678-1.44-1.76-3.413-3.037-4.732c-.633-.654-1.414-1.252-2.314-1.449c-.99-.216-1.95.088-2.761.9l1.414 1.414c.389-.388.674-.413.92-.36c.335.073.78.344 1.305.886c1.039 1.073 2.007 2.799 2.663 4.193l1.81-.852ZM8.793 3.293c-.398.398-.688.854-.816 1.368a2.46 2.46 0 0 0 .1 1.492c.338.876 1.11 1.571 1.922 2.09C11.647 9.295 14.06 10 16 10V8c-1.559 0-3.597-.595-4.924-1.443c-.676-.431-1.022-.836-1.132-1.122a.456.456 0 0 1-.027-.289c.022-.087.088-.237.29-.439L8.793 3.293ZM17 9c0-1.914.96-3.507 2.111-4.396c1.227-.948 2.223-.856 2.682-.397l1.414-1.414c-1.541-1.541-3.795-.949-5.319.228C16.29 4.257 15 6.414 15 9h2Zm4.793-4.793c.405.405.414.697.347.934c-.094.33-.412.773-1.047 1.242C19.834 7.313 17.794 8 16 8v2c2.206 0 4.666-.814 6.282-2.008c.802-.593 1.516-1.369 1.782-2.305c.292-1.028-.012-2.05-.857-2.894l-1.414 1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Hamburger() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h22M5 16h22M5 24h22"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func HamburgerSidebar() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8h15M5 16h22M5 24h22M5 11l3-3l-3-3"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Inbox() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m5 16l-.94-.342A1 1 0 0 0 4 16h1Zm22 0h1a1 1 0 0 0-.06-.342L27 16Zm-5.5 0v-1a1 1 0 0 0-.92.606l.92.394Zm-11 0l.92-.394A1 1 0 0 0 10.5 15v1ZM23 5l.94-.342A1 1 0 0 0 23 4v1Zm3 14.5a1 1 0 1 0 2 0h-2ZM9 5V4a1 1 0 0 0-.94.658L9 5ZM4 19.5a1 1 0 1 0 2 0H4ZM27 16h-1v8h2v-7.999L27 16Zm-1 8c0 .425-.223.933-.645 1.355c-.422.422-.93.645-1.355.645v2c1.075 0 2.067-.527 2.77-1.23c.703-.703 1.23-1.695 1.23-2.77h-2ZM8 26c-.425 0-.933-.223-1.355-.645C6.222 24.933 6 24.425 6 24H4c0 1.075.528 2.067 1.23 2.77C5.933 27.473 6.925 28 8 28v-2Zm-2-2v-8H4v8h2Zm10-3.5a6.85 6.85 0 0 0 4.957-2.043a7.718 7.718 0 0 0 1.096-1.384a6.57 6.57 0 0 0 .355-.655l.007-.014l.002-.006l.001-.002v-.001c.001 0 .001-.001-.918-.395c-.92-.394-.919-.394-.919-.395l.001-.002v-.001c.001-.001.001 0 0 0l-.006.014l-.039.08a4.591 4.591 0 0 1-.184.324a5.72 5.72 0 0 1-.81 1.023c-.735.735-1.868 1.457-3.543 1.457v2Zm5.5-3.5H27v-2h-5.5v2ZM5 17h5.5v-2H5v2Zm5.5-1c-.92.394-.919.394-.919.395l.002.003l.002.006l.006.014a2.994 2.994 0 0 0 .09.187a7.718 7.718 0 0 0 1.361 1.852C12.059 19.472 13.676 20.5 16 20.5v-2c-1.675 0-2.808-.723-3.543-1.457a5.718 5.718 0 0 1-.81-1.023a4.591 4.591 0 0 1-.223-.404l-.007-.014s0-.001 0 0v.001l.002.002L10.5 16ZM24 26H8v2h16v-2ZM22.06 5.342l4 11l1.88-.684l-4-11l-1.88.684ZM26 16v3.5h2V16h-2ZM8.06 4.658l-4 11l1.88.684l4-11l-1.88-.684ZM4 16v3.5h2V16H4ZM9 6h14V4H9v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func InboxAdd() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 27H8c-1.5 0-3-1.5-3-3v-8m0 0h5.5s1 3.5 5.5 3.5s5.5-3.5 5.5-3.5H27L23 5H9L5 16Zm20 12v-8m-4 4h8"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func InboxDouble() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M21.5 11v-1a1 1 0 0 0-.962.725L21.5 11Zm5.5 0h1a.999.999 0 0 0-.106-.447L27 11ZM5 11l-.894-.447A1 1 0 0 0 4 11h1Zm5.5 0l.961-.275A1 1 0 0 0 10.5 10v1ZM5 19H4a1 1 0 0 0 1 1v-1Zm22 0v1a1 1 0 0 0 1-1h-1ZM8 5V4a1 1 0 0 0-.894.553L8 5Zm16 0l.894-.447A1 1 0 0 0 24 4v1Zm2 9.5a1 1 0 1 0 2 0h-2ZM4 14a1 1 0 1 0 2 0H4Zm0-1v11h2V13H4Zm0 11c0 1.075.528 2.067 1.23 2.77C5.933 27.473 6.925 28 8 28v-2c-.425 0-.933-.223-1.355-.645C6.222 24.933 6 24.425 6 24H4Zm4 4h16v-2H8v2Zm16 0c1.075 0 2.067-.527 2.77-1.23c.703-.703 1.23-1.695 1.23-2.77h-2c0 .425-.223.933-.645 1.355c-.422.422-.93.645-1.355.645v2Zm4-4V13h-2v11h2Zm-12-8.5c2.547 0 4.182-1.005 5.17-2.07c.484-.52.8-1.041.998-1.436a5.28 5.28 0 0 0 .285-.691a1.307 1.307 0 0 0 .007-.023v-.003l.001-.001L21.5 11a87.01 87.01 0 0 1-.961-.276v-.003l.001-.002v-.002l-.002.01a3.259 3.259 0 0 1-.159.373c-.13.261-.346.615-.674.97c-.637.685-1.752 1.43-3.705 1.43v2Zm5.5-3.5H27v-2h-5.5v2ZM5 12h5.5v-2H5v2Zm5.5-1a78.07 78.07 0 0 0-.961.276v.004l.003.006a1.573 1.573 0 0 0 .02.066c.013.038.03.09.054.152c.047.124.117.292.216.49a6.2 6.2 0 0 0 .998 1.437c.988 1.064 2.623 2.069 5.17 2.069v-2c-1.953 0-3.068-.745-3.705-1.43a4.196 4.196 0 0 1-.674-.97a3.259 3.259 0 0 1-.162-.383v.002l.001.002v.002l.001.001L10.5 11ZM16 21.5c-2.236 0-3.323-.768-3.866-1.4a3.104 3.104 0 0 1-.543-.91a2.405 2.405 0 0 1-.102-.338v-.006v.011s0 .002-.989.143a68.748 68.748 0 0 0-.99.144a.23.23 0 0 0 0 .004l.002.007a1.565 1.565 0 0 0 .012.071a4.387 4.387 0 0 0 .197.676c.158.414.43.957.895 1.499c.957 1.117 2.62 2.099 5.384 2.099v-2ZM10.5 18H5v2h5.5v-2ZM27 18h-5.5v2H27v-2Zm-5.5 1a76.41 76.41 0 0 1-.99-.144v-.002l.001-.004a.1.1 0 0 1 0-.004v.006l-.014.06c-.014.06-.042.158-.088.279a3.104 3.104 0 0 1-.543.908c-.543.633-1.63 1.401-3.866 1.401v2c2.764 0 4.427-.982 5.384-2.1a5.101 5.101 0 0 0 .894-1.497a4.36 4.36 0 0 0 .207-.73a.946.946 0 0 0 .003-.018l.001-.008v-.004s.001-.002-.989-.143Zm4.5-8v8h2v-8h-2ZM6 19v-8H4v8h2ZM8 6h16V4H8v2Zm15.106-.553l3 6l1.788-.894l-3-6l-1.788.894ZM26 11v3.5h2V11h-2ZM7.106 4.553l-3 6l1.788.894l3-6l-1.788-.894ZM4 11v3h2v-3H4Zm6.5 9h11v-2h-11v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func InboxList() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M17 28a1 1 0 1 0 0-2v2ZM5 16l-.94-.342A1 1 0 0 0 4 16h1Zm5.5 0l.961-.275A1 1 0 0 0 10.5 15v1Zm11 0v-1a1 1 0 0 0-.962.725L21.5 16Zm5.5 0v1a1 1 0 0 0 .94-1.342L27 16ZM23 5l.94-.342A1 1 0 0 0 23 4v1ZM9 5V4a1 1 0 0 0-.94.658L9 5Zm12 17a1 1 0 1 0 0 2v-2Zm6 2a1 1 0 1 0 0-2v2Zm-6 2a1 1 0 1 0 0 2v-2Zm6 2a1 1 0 1 0 0-2v2Zm-10-2H8v2h9v-2Zm-9 0c-.425 0-.933-.223-1.355-.645C6.222 24.933 6 24.425 6 24H4c0 1.075.528 2.067 1.23 2.77C5.933 27.473 6.925 28 8 28v-2Zm-2-2v-8H4v8h2Zm-1-7h5.5v-2H5v2Zm5.5-1a78.07 78.07 0 0 0-.961.276v.004l.003.006a1.573 1.573 0 0 0 .02.066c.013.038.03.09.054.152c.047.124.117.292.216.49a6.2 6.2 0 0 0 .998 1.436c.988 1.065 2.623 2.07 5.17 2.07v-2c-1.953 0-3.068-.745-3.705-1.43a4.196 4.196 0 0 1-.674-.97a3.259 3.259 0 0 1-.162-.383v.002l.001.002v.002l.001.001L10.5 16Zm5.5 4.5c2.547 0 4.182-1.005 5.17-2.07c.484-.52.8-1.041.998-1.436a5.28 5.28 0 0 0 .285-.691a1.307 1.307 0 0 0 .007-.023v-.003l.001-.001L21.5 16a87.01 87.01 0 0 1-.961-.276v-.003l.001-.002v-.002l-.002.01a3.259 3.259 0 0 1-.159.372c-.13.262-.346.616-.674.97c-.637.686-1.752 1.431-3.705 1.431v2Zm5.5-3.5H27v-2h-5.5v2Zm6.44-1.342l-4-11l-1.88.684l4 11l1.88-.684ZM23 4H9v2h14V4Zm-14.94.658l-4 11l1.88.684l4-11l-1.88-.684ZM21 24h6v-2h-6v2Zm0 4h6v-2h-6v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func InboxNewsletter() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m5 16l-.98-.196A1 1 0 0 0 4 16h1Zm22 0h1a1 1 0 0 0-.02-.196L27 16Zm-5.5 0v-1a1 1 0 0 0-.962.725L21.5 16Zm-11 0l.961-.275A1 1 0 0 0 10.5 15v1Zm16.48-5.196a1 1 0 0 0-1.96.392l1.96-.392ZM4 19.5a1 1 0 1 0 2 0H4Zm2.98-8.304a1 1 0 0 0-1.96-.392l1.96.392ZM13.5 8a1 1 0 1 0 0 2V8Zm5 2a1 1 0 1 0 0-2v2Zm-5 2a1 1 0 1 0 0 2v-2Zm5 2a1 1 0 1 0 0-2v2Zm-10-1a1 1 0 1 0 2 0h-2Zm1-8V4a1 1 0 0 0-1 1h1Zm13 0h1a1 1 0 0 0-1-1v1Zm-1 8a1 1 0 1 0 2 0h-2Zm5.5 3h-1v8h2v-7.999L27 16Zm-1 8c0 .425-.223.933-.645 1.355c-.422.422-.93.645-1.355.645v2c1.075 0 2.067-.527 2.77-1.23c.703-.703 1.23-1.695 1.23-2.77h-2ZM8 26c-.425 0-.933-.223-1.355-.645C6.222 24.933 6 24.425 6 24H4c0 1.075.528 2.067 1.23 2.77C5.933 27.473 6.925 28 8 28v-2Zm-2-2v-8H4v8h2Zm10-3.5c2.547 0 4.182-1.005 5.17-2.07c.484-.52.8-1.041.998-1.436a5.28 5.28 0 0 0 .285-.691a1.307 1.307 0 0 0 .007-.023v-.003l.001-.001L21.5 16a87.01 87.01 0 0 1-.961-.276v-.003l.001-.002v-.002l-.002.01a3.259 3.259 0 0 1-.159.372c-.13.262-.346.616-.674.97c-.637.686-1.752 1.431-3.705 1.431v2Zm5.5-3.5H27v-2h-5.5v2ZM5 17h5.5v-2H5v2Zm5.5-1a78.07 78.07 0 0 0-.961.276v.004l.003.006a1.573 1.573 0 0 0 .02.066c.013.038.03.09.054.152c.047.124.117.292.216.49a6.2 6.2 0 0 0 .998 1.436c.988 1.065 2.623 2.07 5.17 2.07v-2c-1.953 0-3.068-.745-3.705-1.43a4.196 4.196 0 0 1-.674-.97a3.259 3.259 0 0 1-.162-.383v.002l.001.002v.002l.001.001L10.5 16ZM24 26H8v2h16v-2Zm1.02-14.804l1 5l1.96-.392l-1-5l-1.96.392ZM4 16v3.5h2V16H4Zm1.02-5.196l-1 5l1.96.392l1-5l-1.96-.392ZM13.5 10h5V8h-5v2Zm0 4h5v-2h-5v2Zm-3-1V5h-2v8h2Zm-1-7h13V4h-13v2Zm12-1v8h2V5h-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Info() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 14h1v9h1m12-7a13 13 0 1 1-26 0a13 13 0 0 1 26 0Z"/><path fill="currentColor" d="M17 9.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func InlineDown() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 13l8 8l8-8"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func InlineLeft() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 8l-8 8l8 8"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func InlineRight() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 24l8-8l-8-8"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func InlineUp() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m24 19l-8-8l-8 8"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Jump() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M26.622 5.17c.096-.043.185.072.119.154l-7.355 9.193a.5.5 0 0 0 .167.76l3.563 1.781a1 1 0 0 1-.037 1.806L5.38 26.83c-.096.043-.185-.072-.12-.154l7.355-9.193a.5.5 0 0 0-.167-.76l-3.563-1.781a1 1 0 0 1 .037-1.806l17.7-7.966Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func JumpAlt() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M25.657 19.707a1 1 0 0 0-1.414-1.414l1.414 1.414ZM18.95 25l-.707.707a1 1 0 0 0 1.414 0L18.95 25Zm-5.293-6.707a1 1 0 0 0-1.414 1.414l1.414-1.414ZM18 24.5a1 1 0 1 0 2 0h-2ZM9 6a1 1 0 0 0 0 2V6Zm1-1a1 1 0 0 0-2 0h2ZM8 9a1 1 0 0 0 2 0H8Zm16.243 9.293l-6 6l1.414 1.414l6-6l-1.414-1.414Zm-4.586 6l-6-6l-1.414 1.414l6 6l1.414-1.414ZM20 24.5V9h-2v15.5h2ZM17 6H9v2h8V6Zm3 3a3 3 0 0 0-3-3v2a1 1 0 0 1 1 1h2ZM8 5v4h2V5H8Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Label() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 19l5.5-5.5M13 22l2.5-2.5M4.414 16.586l11-11A2 2 0 0 1 16.828 5H25a2 2 0 0 1 2 2v8.172a2 2 0 0 1-.586 1.414l-11 11a2 2 0 0 1-2.828 0l-8.172-8.172a2 2 0 0 1 0-2.828Z"/><path fill="currentColor" d="M23 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LabelMini() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 14.172V10a2 2 0 0 0-2-2h-4.172a2 2 0 0 0-1.414.586l-8 8a2 2 0 0 0 0 2.828l4.172 4.172a2 2 0 0 0 2.828 0l8-8A2 2 0 0 0 24 14.172Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Link() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.9 19.142l-.708-.707c-2.12-2.121-2.12-4.95 0-7.071l2.829-2.828c2.121-2.122 4.95-2.122 7.07 0c2.122 2.12 2.122 4.95 0 7.07m-6.363-2.12l.707.706c2.121 2.122 2.121 4.95 0 7.072l-2.828 2.828c-2.122 2.121-4.95 2.121-7.071 0c-2.122-2.121-2.122-4.95 0-7.071"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LinkOut() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 3h7v7m-1.5-5.5L20 12m-3-7H8a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-9"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func List() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M12 7a1 1 0 1 0 0 2V7Zm15 2a1 1 0 1 0 0-2v2Zm-15 6a1 1 0 1 0 0 2v-2Zm9 2a1 1 0 1 0 0-2v2Zm-9 6a1 1 0 1 0 0 2v-2Zm15 2a1 1 0 1 0 0-2v2ZM12 9h15V7H12v2Zm0 8h9v-2h-9v2Zm0 8h15v-2H12v2Zm-6-1v2a2 2 0 0 0 2-2H6Zm0 0H4a2 2 0 0 0 2 2v-2Zm0 0v-2a2 2 0 0 0-2 2h2Zm0 0h2a2 2 0 0 0-2-2v2Zm0-8v2a2 2 0 0 0 2-2H6Zm0 0H4a2 2 0 0 0 2 2v-2Zm0 0v-2a2 2 0 0 0-2 2h2Zm0 0h2a2 2 0 0 0-2-2v2Zm0-8v2a2 2 0 0 0 2-2H6Zm0 0H4a2 2 0 0 0 2 2V8Zm0 0V6a2 2 0 0 0-2 2h2Zm0 0h2a2 2 0 0 0-2-2v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LoadingSpin() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.5 4.742A13 13 0 1 0 29 16"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Lock() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 24v-4m5-5V8a5 5 0 0 0-10 0v7M6 27V17a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func LockWindow() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor"><circle cx="22" cy="23" r=".5" fill="currentColor"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 19v-2.917c0-1.15-.895-2.083-2-2.083s-2 .933-2 2.083V19m-7 8H8a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v4M17 25v-4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2Z"/><path fill="currentColor" d="M9.5 9a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm4 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm4 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Mail() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m14.862 17.212l-.57.822l.57-.822Zm2.276 0l.57.822l-.57-.822ZM5 8h22V6H5v2Zm23 1v14h2V9h-2Zm-1 15H5v2h22v-2ZM4 23V9H2v14h2Zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3v-2Zm23-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM27 8a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM5 6a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V6ZM2.43 9.822l11.862 8.212l1.139-1.644L3.569 8.178L2.431 9.822Zm15.278 8.212l11.861-8.212l-1.138-1.644l-11.862 8.212l1.139 1.644Zm-3.416 0a3 3 0 0 0 3.416 0l-1.139-1.644a1 1 0 0 1-1.138 0l-1.139 1.644Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MailList() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h14m4 0h4M5 13h10m8 0h4M5 19h14m4 0h4M5 25h12m6 0h4"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MailOpen() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m15.118 3.271l-.294-.955l.294.955Zm1.764 0l-.294.956l.294-.956Zm11.412 3.512l.294-.956l-.294.956Zm-24.588 0l-.294-.956l.294.956Zm.908.428a1 1 0 1 0-1.228 1.578l1.228-1.578ZM13 15l-.614.79A1 1 0 0 0 13 16v-1Zm6 0v1a1 1 0 0 0 .614-.21L19 15Zm9.614-6.21a1 1 0 1 0-1.228-1.58l1.228 1.58ZM28 7.738V23h2V7.739h-2ZM27 24H5v2h22v-2ZM4 23V7.739H2V23h2ZM4 7.739l11.412-3.512l-.588-1.911L3.412 5.827L4 7.739Zm12.588-3.512L28 7.74l.588-1.912l-11.412-3.511l-.588 1.911Zm-1.176 0a2 2 0 0 1 1.176 0l.588-1.911a4 4 0 0 0-2.352 0l.588 1.911ZM5 24a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3v-2Zm23-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm2-15.261a2 2 0 0 0-1.412-1.912L28 7.739h2Zm-26 0l-.588-1.912A2 2 0 0 0 2 7.739h2Zm-.614 1.05l9 7l1.228-1.578l-9-7l-1.228 1.578ZM13 16h6v-2h-6v2Zm6.614-.21l9-7l-1.228-1.58l-9 7l1.228 1.58Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MailPlus() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25 28v-8m-4 4h8m0-15v7m0-7a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2m26 0l-11.862 8.212a2 2 0 0 1-2.276 0L3 9m0 0v14a2 2 0 0 0 2 2h10"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MailSubbed() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 9v7m0-7c0-1-1-2-2-2H5C4 7 3 8 3 9m26 0l-11.862 8.212a2 2 0 0 1-2.276 0L3 9m13 16H5c-1 0-2-1-2-2V9m16.5 14l3 3l6-6"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MailUnsub() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 9v7m0-7c0-1-1-2-2-2H5C4 7 3 8 3 9m26 0l-11.862 8.212a2 2 0 0 1-2.276 0L3 9m13 16H5c-1 0-2-1-2-2V9m18 17l6-6m2 3a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Markdown() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20v-8l-4 4l-4-4v8m12-3.5l3.5 3.5m0 0l3.5-3.5M22.5 20v-9M5 7h22a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MeatballsH() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func MeatballsV() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-width="2" d="M16 24a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Moon() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.294 5A11.19 11.19 0 1 0 27 18.706s-5.723 2.19-10.81-2.897C11.105 10.723 13.295 5 13.295 5Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Mute() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 21H8a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h7l6.586-6.586C22.846 3.154 25 4.047 25 5.828V6m0 8.5v11.672c0 1.781-2.154 2.674-3.414 1.414L17 23M7 28L29 6"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Notifications() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27 25c0 1.657-4.925 3-11 3S5 26.657 5 25m22 0c0-1.47-3.88-2.694-9-2.95M27 25c0-2-.2-6.2-3-9c-2.986-2.986-.513-7.427-5-8.668M5 25c0-1.47 3.88-2.694 9-2.95M5 25c0-2 .2-6.2 3-9c2.986-2.986.513-7.427 5-8.668m1 14.717a40.015 40.015 0 0 1 4 0m-4 0V23a2 2 0 1 0 4 0v-.95M13 7.331C13.773 7.12 14.751 7 16 7s2.227.119 3 .332m-6 0V7a3 3 0 0 1 6 0v.332"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Nuclear() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor" stroke-width="2"><path fill="currentColor" d="M19.316 19.743A4.996 4.996 0 0 0 20.9 17h4.045a8.992 8.992 0 0 1-3.607 6.247l-2.023-3.504ZM11.1 17a4.996 4.996 0 0 0 1.585 2.743l-2.024 3.504A8.991 8.991 0 0 1 7.056 17H11.1Zm4.9-6c-.553 0-1.086.09-1.584.256l-2.023-3.504A8.966 8.966 0 0 1 16 7c1.284 0 2.504.268 3.607.752l-2.023 3.504A4.994 4.994 0 0 0 16 11Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 16a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Off() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.5 4.742a13 13 0 1 1-13 0M16 3v10"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Outbox() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m5 16l-.923-.385A1 1 0 0 0 4 16h1Zm22 0h1a1 1 0 0 0-.077-.385L27 16Zm-5.5 0v-1a1 1 0 0 0-.962.725L21.5 16Zm-11 0l.961-.275A1 1 0 0 0 10.5 15v1Zm14-6l.923-.385A1 1 0 0 0 24.5 9v1Zm1.5 9.5a1 1 0 1 0 2 0h-2Zm-22 0a1 1 0 1 0 2 0H4ZM22.5 9a1 1 0 1 0 0 2V9Zm-15 1V9a1 1 0 0 0-.923.615L7.5 10Zm2 1a1 1 0 1 0 0-2v2Zm5.5 4a1 1 0 1 0 2 0h-2Zm2-11a1 1 0 1 0-2 0h2Zm-1-1l.707-.707a1 1 0 0 0-1.414 0L16 3Zm3.293 4.707a1 1 0 1 0 1.414-1.414l-1.414 1.414Zm-8-1.414a1 1 0 0 0 1.414 1.414l-1.414-1.414ZM27 16h-1v8h2v-7.999L27 16Zm-1 8c0 .425-.223.933-.645 1.355c-.422.422-.93.645-1.355.645v2c1.075 0 2.067-.527 2.77-1.23c.703-.703 1.23-1.695 1.23-2.77h-2ZM8 26c-.425 0-.933-.223-1.355-.645C6.222 24.933 6 24.425 6 24H4c0 1.075.528 2.067 1.23 2.77C5.933 27.473 6.925 28 8 28v-2Zm-2-2v-8H4v8h2Zm10-3.5c2.547 0 4.182-1.005 5.17-2.07c.484-.52.8-1.041.998-1.436a5.28 5.28 0 0 0 .285-.691a1.307 1.307 0 0 0 .007-.023v-.003l.001-.001L21.5 16a87.01 87.01 0 0 1-.961-.276v-.003l.001-.002v-.002l-.002.01a3.259 3.259 0 0 1-.159.372c-.13.262-.346.616-.674.97c-.637.686-1.752 1.431-3.705 1.431v2Zm5.5-3.5H27v-2h-5.5v2ZM5 17h5.5v-2H5v2Zm5.5-1a78.07 78.07 0 0 0-.961.276v.004l.003.006a1.573 1.573 0 0 0 .02.066c.013.038.03.09.054.152c.047.124.117.292.216.49a6.2 6.2 0 0 0 .998 1.436c.988 1.065 2.623 2.07 5.17 2.07v-2c-1.953 0-3.068-.745-3.705-1.43a4.196 4.196 0 0 1-.674-.97a3.259 3.259 0 0 1-.162-.383v.002l.001.002v.002l.001.001L10.5 16ZM24 26H8v2h16v-2Zm-.423-15.615l2.5 6l1.846-.77l-2.5-6l-1.846.77ZM26 16v3.5h2V16h-2ZM4 16v3.5h2V16H4Zm18.5-5h2V9h-2v2Zm-15 0h2V9h-2v2Zm-.923-1.385l-2.5 6l1.846.77l2.5-6l-1.846-.77ZM17 15V4h-2v11h2ZM15.293 3.707l4 4l1.414-1.414l-4-4l-1.414 1.414Zm0-1.414l-4 4l1.414 1.414l4-4l-1.414-1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Paper() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="currentColor"><path d="M10 8a1 1 0 0 0 0 2V8Zm4 2a1 1 0 1 0 0-2v2Zm-4 5a1 1 0 1 0 0 2v-2Zm12 2a1 1 0 1 0 0-2v2Zm-12 2a1 1 0 1 0 0 2v-2Zm12 2a1 1 0 1 0 0-2v2Zm-12 2a1 1 0 1 0 0 2v-2Zm4 2a1 1 0 1 0 0-2v2ZM8 4h16V2H8v2Zm17 1v22h2V5h-2Zm-1 23H8v2h16v-2ZM7 27V5H5v22h2Zm1 1a1 1 0 0 1-1-1H5a3 3 0 0 0 3 3v-2Zm17-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM24 4a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM8 2a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V2Zm2 8h4V8h-4v2Zm0 7h12v-2H10v2Zm0 4h12v-2H10v2Zm0 4h4v-2h-4v2Z"/><circle cx="22" cy="9" r=".5" stroke="currentColor"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pause() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 7v18m9-18v18"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Phone() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 25h2M11 3h10a2 2 0 0 1 2 2v22a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Pin() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 18l-8 8M20.667 4L28 11.333l-6.38 6.076a2 2 0 0 0-.62 1.448v3.729c0 .89-1.077 1.337-1.707.707L8.707 12.707c-.63-.63-.184-1.707.707-1.707h3.729a2 2 0 0 0 1.448-.62L20.667 4Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Play() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 24.414V7.586c0-1.746 2.081-2.653 3.36-1.465l9.062 8.413a2 2 0 0 1 0 2.932l-9.061 8.413C13.08 27.067 11 26.16 11 24.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func PrintAlt() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 22h-3.621c-.252 0-.504-.022-.752-.067l-.92-.164c-3.383-.604-5.244-3.076-5.62-5.769c-.484-3.466 1.494-7.3 5.861-7.992a.114.114 0 0 0 .08-.053l.636-1.049C13.44 3.982 16.582 3.56 19 4.687c2.378 1.11 4.055 3.722 3.037 6.93a.11.11 0 0 0 .084.142c2.952.586 4.372 3.143 4.173 5.567a5.188 5.188 0 0 1-.915 2.566c-1.353 1.92-4.028 2.1-6.377 2.115c-1 .006-2-.007-3.002-.007Z"/><path d="m14.5 22l-.5 6h4l-.5-6h-3Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Printer() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M27 12h1h-1Zm0 10h-1h1ZM5 22h1h-1Zm0-10H4h1Zm17 11a1 1 0 1 0 0 2v-2Zm-12 2a1 1 0 1 0 0-2v2ZM9 10a1 1 0 1 0 2 0H9Zm1-5V4a1 1 0 0 0-1 1h1Zm12 0h1a1 1 0 0 0-1-1v1Zm-1 5a1 1 0 1 0 2 0h-2ZM9 16a1 1 0 1 0 0 2v-2Zm14 2a1 1 0 1 0 0-2v2Zm-12-1a1 1 0 1 0-2 0h2Zm-1 10H9a1 1 0 0 0 1 1v-1Zm12 0v1a1 1 0 0 0 1-1h-1Zm1-10a1 1 0 1 0-2 0h2Zm3-5v10h2V12h-2ZM6 22V12H4v10h2Zm1-11h18V9H7v2Zm18 12h-3v2h3v-2Zm-15 0H7v2h3v-2Zm-6-1a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1H4Zm22 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm2-10a3 3 0 0 0-3-3v2a1 1 0 0 1 1 1h2ZM6 12a1 1 0 0 1 1-1V9a3 3 0 0 0-3 3h2Zm5-2V5H9v5h2Zm-1-4h12V4H10v2Zm11-1v5h2V5h-2ZM9 18h14v-2H9v2Zm0-1v10h2V17H9Zm1 11h12v-2H10v2Zm13-1V17h-2v10h2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Queue() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="m9.83 8.465l.69-.724l-.69.724Zm7.632 7.269l.69-.724l-.69.724Zm0 2.532l-.69-.724l.69.724ZM9.83 25.535l.69.724l-.69-.724ZM24 5a1 1 0 1 0-2 0h2Zm-2 8a1 1 0 1 0 2 0h-2Zm-3-5a1 1 0 1 0 0 2V8Zm8 2a1 1 0 1 0 0-2v2ZM6 9.731V24.27h2V9.73H6Zm3.14-.542l7.632 7.269l1.38-1.448l-7.632-7.27l-1.38 1.449Zm7.632 8.353l-7.631 7.269l1.38 1.448l7.63-7.269l-1.379-1.448Zm0-1.084a.758.758 0 0 1 0 1.084l1.38 1.448a2.758 2.758 0 0 0 0-3.98l-1.38 1.448ZM6 24.268c0 2.335 2.767 3.66 4.52 1.991l-1.38-1.448c-.402.383-1.14.14-1.14-.542H6ZM8 9.732c0-.683.738-.925 1.14-.542l1.38-1.448C8.767 6.071 6 7.397 6 9.731h2ZM22 5v8h2V5h-2Zm-3 5h8V8h-8v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Remind() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M15 8c-2 2-6.4 5.6-8 4c-3.5-3.5 4.667-3.333 8-4Zm0 0c2.833 0 12-1 8.5 2.5C21 13 17.833 9.667 15 8Zm-2.172 17.828L7.5 20.5a2.121 2.121 0 1 1 3-3l1.83 1.83c.247.247.67.072.67-.277V7a2 2 0 1 1 4 0v7.997c0 .333.419.48.626.22a1.759 1.759 0 0 1 2.747 0l.398.497a.453.453 0 0 0 .557.122l.666-.333a2.249 2.249 0 0 1 2.012 0c.609.304.994.927.994 1.609V23a4 4 0 0 1-4 4h-5.343a4 4 0 0 1-2.829-1.172Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Reply() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.608 12.526l7.04-6.454C13.931 4.896 16 5.806 16 7.546V11c13 0 11 16 11 16s-4-10-11-10v3.453c0 1.74-2.069 2.65-3.351 1.475l-7.04-6.454a2 2 0 0 1 0-2.948Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Replyall() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 19.141v1.313c0 1.74-2.069 2.65-3.351 1.474l-7.04-6.454a2 2 0 0 1 0-2.948l7.04-6.454C11.93 4.896 14 5.806 14 7.546V8.86m3.04-2.787L10 12.526a2 2 0 0 0 0 2.948l7.04 6.454c1.283 1.176 3.352.266 3.352-1.475V17c5 0 8.5 10 8.5 10s1.5-16-8.5-16V7.546c0-1.74-2.069-2.65-3.352-1.474Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Search() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 27l7.5-7.5M28 13a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SearchAlt() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m28 27l-7.5-7.5M5 13a9 9 0 1 0 18 0a9 9 0 0 0-18 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Send() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 3L3 15l12 2.5M29 3L19 29l-4-11.5M29 3L15 17.5"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SendCancelled() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 19l4 4m0 0l4 4m-4-4l-4 4m4-4l4-4m1-16L3 15l12 2.5M29 3L15 17.5M29 3l-4.375 11.375M15 17.5l.625 2.875"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SendLater() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 21v2l-1.5 1.5M29 3L3 15l12 2.5M29 3L15 17.5M29 3l-4.375 11.375M15 17.5l.625 2.875M29 23a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SendStop() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 26l6-6m2-17L3 15l12 2.5M29 3L15 17.5M29 3l-4.375 11.375M15 17.5l.625 2.875M29 23a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Share() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M6 17a1 1 0 1 0-2 0h2Zm22 0a1 1 0 1 0-2 0h2Zm-13 3a1 1 0 1 0 2 0h-2Zm2-16.5a1 1 0 1 0-2 0h2ZM16 3l.707-.707a1 1 0 0 0-1.414 0L16 3Zm5.293 6.707a1 1 0 0 0 1.414-1.414l-1.414 1.414Zm-12-1.414a1 1 0 0 0 1.414 1.414L9.293 8.293ZM4 17v8h2v-8H4Zm3 11h18v-2H7v2Zm21-3v-8h-2v8h2Zm-3 3a3 3 0 0 0 3-3h-2a1 1 0 0 1-1 1v2ZM4 25a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1H4Zm13-5V3.5h-2V20h2ZM15.293 3.707l6 6l1.414-1.414l-6-6l-1.414 1.414Zm0-1.414l-6 6l1.414 1.414l6-6l-1.414-1.414Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Sign() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 25h-5L7 15c2-.667 6-4 6-12h6c0 7.6 4 11.167 6 12l-4 10h-5Zm0 0v-6m0-6V3M9 25h14v4H9v-4Zm9-9a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Signature() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 27C6 19.333 13.5 5 18.5 5C23 5 10 26 15 26c3.5 0 6-10.5 8.5-10.5s-1 9 1 9c2.5 0 4-4 4-4"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Skip() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 5v22M8 7.586v16.828c0 1.746 2.081 2.653 3.36 1.465l9.062-8.413a2 2 0 0 0 0-2.932L11.36 6.121C10.08 4.933 8 5.84 8 7.586Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SnoozeMonth() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M17.207 16.293a1 1 0 0 0-1.414 1.414l1.414-1.414ZM19.5 20l.707.707a1 1 0 0 0 0-1.414L19.5 20Zm-3.707 2.293a1 1 0 0 0 1.414 1.414l-1.414-1.414Zm-3.086-6a1 1 0 0 0-1.414 1.414l1.414-1.414ZM15 20l.707.707a1 1 0 0 0 0-1.414L15 20Zm-3.707 2.293a1 1 0 0 0 1.414 1.414l-1.414-1.414ZM5 11a1 1 0 1 0 0 2v-2Zm22 2a1 1 0 1 0 0-2v2Zm-7-5a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2ZM10 8a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2Zm3.793 13.707l3 3l1.414-1.414l-3-3l-1.414 1.414Zm3 1.586l-3 3l1.414 1.414l3-3l-1.414-1.414Zm-7.5-1.586l3 3l1.414-1.414l-3-3l-1.414 1.414Zm3 1.586l-3 3l1.414 1.414l3-3l-1.414-1.414ZM7 7h18V5H7v2Zm19 1v18h2V8h-2Zm-1 19H7v2h18v-2ZM6 26V8H4v18h2Zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm19-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 7a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM7 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm-2 8h22v-2H5v2Zm17-5V4h-2v4h2ZM12 8V4h-2v4h2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SnoozeTomorrow() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21h22M7 26h18M10 16a6 6 0 0 1 12 0M6.5 16h-2M9 9L7.5 7.5M16 6V4m7 5l1.485-1.485M28 16h-2.5"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SnoozeWeek() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M15.707 16.293a1 1 0 0 0-1.414 1.414l1.414-1.414ZM18 20l.707.707a1 1 0 0 0 0-1.414L18 20Zm-3.707 2.293a1 1 0 0 0 1.414 1.414l-1.414-1.414ZM5 11a1 1 0 1 0 0 2v-2Zm22 2a1 1 0 1 0 0-2v2Zm-7-5a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2ZM10 8a1 1 0 1 0 2 0h-2Zm2-4a1 1 0 1 0-2 0h2Zm2.293 13.707l3 3l1.414-1.414l-3-3l-1.414 1.414Zm3 1.586l-3 3l1.414 1.414l3-3l-1.414-1.414ZM7 7h18V5H7v2Zm19 1v18h2V8h-2Zm-1 19H7v2h18v-2ZM6 26V8H4v18h2Zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm19-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 7a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM7 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V5Zm-2 8h22v-2H5v2Zm17-5V4h-2v4h2ZM12 8V4h-2v4h2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SnoozeWeekend() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.5 14v-4a3 3 0 0 1 3-3h15a3 3 0 0 1 3 3v4m-21 0A2.5 2.5 0 0 0 3 16.5V25h5M5.5 14A2.5 2.5 0 0 1 8 16.5V20h16v-3.5a2.5 2.5 0 0 1 2.5-2.5m0 0a2.5 2.5 0 0 1 2.5 2.5V25h-5m0 0H8m16 0v2M8 25v2"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Sort() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="currentColor"><circle cx="18.078" cy="8.286" r="1" transform="rotate(-9 18.078 8.286)"/><path d="M8.045 8.862a1 1 0 0 0 .313 1.976l-.313-1.976Zm4.263 1.35a1 1 0 0 0-.313-1.976l.313 1.976Zm-3.325 4.576a1 1 0 1 0 .313 1.976l-.313-1.976Zm10.19.411a1 1 0 1 0-.313-1.975l.313 1.975ZM9.61 18.74a1 1 0 0 0 .313 1.976L9.61 18.74Zm8.214.724a1 1 0 1 0-.312-1.975l.313 1.975ZM8.73 25.966l-.157-.988l.157.988Zm-2.29-1.663l-.988.157l.988-.157Zm17.778-2.816l-.988.157l.988-.156Zm-1.663 2.289l.157.987l-.157-.987ZM19.271 3.034l-.156-.987l.156.987Zm2.288 1.663l.988-.157l-.988.157Zm-16.115.527l.156.988l-.156-.988ZM3.78 7.513l.988-.157l-.988.157ZM26 10v17h2V10h-2Zm-1 18H11v2h14v-2Zm-14 0a1 1 0 0 1-1-1H8a3 3 0 0 0 3 3v-2Zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 9a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM10 27v-2H8v2h2ZM25 7h-3v2h3V7ZM8.358 10.838l3.95-.626l-.313-1.976l-3.95.626l.313 1.976Zm.938 5.926l9.877-1.565l-.313-1.975l-9.877 1.564l.313 1.976Zm.626 3.95l7.901-1.251l-.312-1.975l-7.902 1.251l.313 1.976ZM5.6 6.213l13.828-2.19l-.313-1.975l-13.828 2.19l.313 1.975Zm14.972-1.359l2.66 16.79l1.975-.312l-2.66-16.79l-1.975.312ZM22.4 22.788l-13.828 2.19l.313 1.975l13.828-2.19l-.313-1.975ZM7.428 24.147l-2.66-16.79l-1.975.312l2.66 16.79l1.975-.312Zm1.144.831a1 1 0 0 1-1.144-.831l-1.975.313a3 3 0 0 0 3.432 2.493l-.313-1.975Zm14.66-3.334a1 1 0 0 1-.832 1.144l.313 1.975a3 3 0 0 0 2.494-3.432l-1.976.313ZM19.427 4.022a1 1 0 0 1 1.144.831l1.975-.313a3 3 0 0 0-3.432-2.493l.313 1.975Zm-14.14.215a3 3 0 0 0-2.495 3.432l1.976-.313A1 1 0 0 1 5.6 6.212l-.313-1.975Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func SortAlt() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="currentColor" d="M5 5.015v-1a1 1 0 0 0-1 1h1Zm8 1a1 1 0 1 0 0-2v2ZM4 13a1 1 0 1 0 2 0H4Zm23-7.985h1a1 1 0 0 0-1-1v1ZM26 13a1 1 0 1 0 2 0h-2Zm-7-8.985a1 1 0 1 0 0 2v-2Zm-13.293.278a1 1 0 0 0-1.414 1.414l1.414-1.414Zm22 1.414a1 1 0 0 0-1.414-1.414l1.414 1.414ZM11 26a1 1 0 1 0 0 2v-2Zm10 2a1 1 0 1 0 0-2v2ZM5 6.015h.015v-2H5v2Zm.015 0H13v-2H5.015v2ZM4 5.015V13h2V5.015H4Zm22 0V13h2V5.015h-2Zm1-1h-.015v2H27v-2Zm-.015 0H19v2h7.985v-2ZM16.707 15.293L5.722 4.308L4.308 5.722l10.985 10.985l1.414-1.414ZM5.722 4.308l-.015-.015l-1.414 1.414l.015.015l1.414-1.414Zm10.985 12.4L27.692 5.721l-1.414-1.414l-10.985 10.985l1.414 1.414ZM27.692 5.721l.015-.015l-1.414-1.414l-.015.015l1.414 1.414ZM17 27V16h-2v11h2Zm-6 1h5v-2h-5v2Zm5 0h5v-2h-5v2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Sound() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12v8a1 1 0 0 0 1 1h7l6.586 6.586c1.26 1.26 3.414.367 3.414-1.414V5.828c0-1.781-2.154-2.674-3.414-1.414L15 11H8a1 1 0 0 0-1 1Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Stack() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="currentColor"><circle cx="19" cy="9" r=".5" stroke="currentColor"/><path d="M9 8a1 1 0 0 0 0 2V8Zm4 2a1 1 0 1 0 0-2v2Zm-4 4a1 1 0 1 0 0 2v-2Zm10 2a1 1 0 1 0 0-2v2ZM9 18a1 1 0 1 0 0 2v-2Zm8 2a1 1 0 1 0 0-2v2Zm9-10v17h2V10h-2Zm-1 18H11v2h14v-2Zm-14 0a1 1 0 0 1-1-1H8a3 3 0 0 0 3 3v-2Zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM25 9a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM10 27v-2H8v2h2ZM25 7h-2v2h2V7ZM9 10h4V8H9v2Zm0 6h10v-2H9v2Zm0 4h8v-2H9v2ZM7 5h14V3H7v2Zm15 1v17h2V6h-2Zm-1 18H7v2h14v-2ZM6 23V6H4v17h2Zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2ZM21 5a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2ZM7 3a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1V3Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func StackAlt() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 15.5l13 6l13-6M3 20l13 6l13-6M3 11l13 6l13-6l-13-6l-13 6Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Star() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.546 4.984a.5.5 0 0 1 .908 0l3.097 6.714a.5.5 0 0 0 .395.287l7.341.87a.5.5 0 0 1 .28.864l-5.427 5.02a.5.5 0 0 0-.15.464l1.44 7.251a.5.5 0 0 1-.735.534l-6.45-3.611a.5.5 0 0 0-.49 0l-6.45 3.61a.5.5 0 0 1-.735-.533l1.44-7.251a.5.5 0 0 0-.15-.465l-5.428-5.02a.5.5 0 0 1 .28-.863l7.342-.87a.5.5 0 0 0 .396-.287l3.096-6.714Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Stopwatch() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27.9 7.869a15 15 0 0 0-2.769-2.77m-7.434-.988a13 13 0 0 0-3.394 0m12.304 2.282l-2.829 2.829M16 17l5 5m6-5c0 6.075-4.925 11-11 11S5 23.075 5 17S9.925 6 16 6s11 4.925 11 11Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Sun() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 6V3M8.929 8.929L6.808 6.808M6 16H3m13 13v-3m9.192-.808l-2.121-2.12M29 16h-3M8.929 23.071l-2.121 2.121M25.192 6.808l-2.12 2.121M22 16a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TextCenter() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22M9 13h14M5 19h22M9 25h14"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TextJustify() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22M5 13h22M5 19h22M5 25h22"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TextLeft() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22M5 13h14M5 19h22M5 25h14"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func TextRight() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22m-14 6h14M5 19h22m-14 6h14"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Todo() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.5 16l3.5 3.5l6-6m8 2.5a13 13 0 1 1-26 0a13 13 0 0 1 26 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Userhappy() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 23s1 1 2.5 1s2.5-1 2.5-1m4.5-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-12 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm18-2c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Userneutral() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 23h5m4.5-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-12 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm18-2c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Usersad() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 24s1-1 2.5-1s2.5 1 2.5 1M29 16c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13Zm-6 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-12 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Vip() *goggles.ElementSVG {
	return goggles.SVG(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 13l3 6.5l3-6.5m3.5 0v6.5m4-2.225h1.48c.651 0 1.277-.275 1.721-.75a2.338 2.338 0 0 0 .215-2.932a1.394 1.394 0 0 0-1.14-.593H20.5v4.275Zm0 0V19.5M5 7h22a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2Z"/>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func Warning() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="currentColor"><path d="m4.6 24.047l-.88-.476l.88.476Zm22.8 0l-.879.477l.88-.477Zm-9.641-17.8l.879-.477l-.88.477Zm-3.518 0l.88.476l-.88-.476ZM15 18a1 1 0 1 0 2 0h-2Zm2-6a1 1 0 1 0-2 0h2Zm-.12-5.277l9.641 17.8l1.759-.952l-9.642-17.8l-1.759.952ZM25.641 26H6.358v2h19.284v-2ZM5.48 24.524l9.64-17.801l-1.759-.953L3.72 23.571l1.759.953ZM6.358 26a1 1 0 0 1-.88-1.476L3.72 23.57C2.637 25.571 4.085 28 6.358 28v-2Zm20.163-1.476A1 1 0 0 1 25.642 26v2c2.273 0 3.72-2.43 2.638-4.429l-1.759.953ZM18.638 5.77c-1.135-2.095-4.141-2.095-5.276 0l1.759.953a1 1 0 0 1 1.758 0l1.759-.953ZM17 18v-6h-2v6h2Z"/><path stroke="currentColor" d="M17 21.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
func WarningAlt() *goggles.ElementSVG {
	return goggles.SVG(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 17.5v-8M29 16a13 13 0 1 1-26 0a13 13 0 0 1 26 0Z"/><path fill="currentColor" d="M17 22a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`).
		Custom("viewBox", "0,0,32,32").
		Custom("width", "1em").
		Custom("height", "1em")
}
