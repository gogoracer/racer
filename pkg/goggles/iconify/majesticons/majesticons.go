package majesticons

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	academicCapInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11.514 3.126a1 1 0 0 1 .972 0l9 5A1 1 0 0 1 22 9v7a1 1 0 1 1-2 0v-5.3l-1 .555v.004l-6.067 3.016a2 2 0 0 1-1.848-.035L2.357 9.479a1 1 0 0 0-.284-.103a1 1 0 0 1 .441-1.25l9-5zM5 13.199V17a1 1 0 0 0 .553.894l6 3a1 1 0 0 0 .894 0l6-3A1 1 0 0 0 19 17v-3.256l-6.083 2.844a2 2 0 0 1-1.805-.056L5 13.2z" fill="currentColor"/></g>`
	academicCapLineInnerSVG           = `<g fill="currentColor"><path d="M12.486 5.414a1 1 0 0 0-.972 0L5.06 9l6.455 3.586a1 1 0 0 0 .972 0L18.94 9l-6.455-3.586zm-1.943-1.749a3 3 0 0 1 2.914 0l8.029 4.46a1 1 0 0 1 0 1.75l-8.03 4.46a3 3 0 0 1-2.913 0l-8.029-4.46a1 1 0 0 1 0-1.75l8.03-4.46z"/><path d="M21 8a1 1 0 0 1 1 1v7a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1zM6 10a1 1 0 0 1 1 1v5.382l4.553 2.276a1 1 0 0 0 .894 0L17 16.382V11a1 1 0 1 1 2 0v6a1 1 0 0 1-.553.894l-5.105 2.553a3 3 0 0 1-2.684 0l-5.105-2.553A1 1 0 0 1 5 17v-6a1 1 0 0 1 1-1z"/></g>`
	addColumnInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h3M3 21h3m0 0h4a2 2 0 0 0 2-2V9M6 21V9m0-6h4a2 2 0 0 1 2 2v4M6 3v6M3 9h3m0 0h6m-9 6h9m3-3h3m0 0h3m-3 0v3m0-3V9"/>`
	addColumnLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h3M3 21h3m0 0h4a2 2 0 0 0 2-2V9M6 21V9m0-6h4a2 2 0 0 1 2 2v4M6 3v6M3 9h3m0 0h6m-9 6h9m3-3h3m0 0h3m-3 0v3m0-3V9"/>`
	addRowInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v3m18-3v3m0 0v4a2 2 0 0 1-2 2H9m12-6H9M3 6v4a2 2 0 0 0 2 2h4M3 6h6m0-3v3m0 0v6m6-9v9m-3 3v3m0 0v3m0-3h3m-3 0H9"/>`
	addRowLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v3m18-3v3m0 0v4a2 2 0 0 1-2 2H9m12-6H9M3 6v4a2 2 0 0 0 2 2h4M3 6h6m0-3v3m0 0v6m6-9v9m-3 3v3m0 0v3m0-3h3m-3 0H9"/>`
	adjustmentsInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M20 7a1 1 0 1 0 0-2h-9.17a3.001 3.001 0 0 0-5.66 0H4a1 1 0 1 0 0 2h1.17a3.001 3.001 0 0 0 5.66 0H20zm0 6a1 1 0 1 0 0-2h-1.17a3.001 3.001 0 0 0-5.66 0H4a1 1 0 1 0 0 2h9.17a3.001 3.001 0 0 0 5.66 0H20zm0 6a1 1 0 1 0 0-2h-9.17a3.001 3.001 0 0 0-5.66 0H4a1 1 0 1 0 0 2h1.17a3.001 3.001 0 0 0 5.66 0H20z" fill="currentColor"/></g>`
	adjustmentsLineInnerSVG           = `<g fill="currentColor"><path d="M8 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM5 6a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm11 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm-5 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0z"/><path d="M3 6a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H4a1 1 0 0 1-1-1zm6 0a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1zm-6 6a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1zm14 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1zM3 18a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1zm6 0a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1z"/></g>`
	airplaneInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M16.48 14h4.02a2.5 2.5 0 1 0 0-5H6.618a1 1 0 0 1-.894-.553l-.448-.894A1 1 0 0 0 4.382 7H2.517a1 1 0 0 0-.92 1.394l2.143 5a1 1 0 0 0 .92.606h3.863a1 1 0 0 1 .928 1.371L8.55 17.63A1 1 0 0 0 9.477 19h2.042a1 1 0 0 0 .781-.375l3.4-4.25a1 1 0 0 1 .78-.375zM9.5 8h4.75L12.3 5.4a1 1 0 0 0-.8-.4H9.618a1 1 0 0 0-.894 1.447L9.5 8z" clip-rule="evenodd"/>`
	airplaneFlightTwoInnerSVG         = `<g fill="none"><g clip-path="url(#majesticonsAirplaneFlight20)"><path fill="currentColor" d="M19.02 9.121v4.271a1.346 1.346 0 0 1-2.57.557l-.965-2.12l-1.414 1.414a2 2 0 0 1-2.828 0l-1.829-1.829a1.414 1.414 0 0 1 1-2.414h1.414a2 2 0 0 0 1.415-.586l.828-.828l-2.02-.674a1.448 1.448 0 0 1 .554-2.818l4.648.31l1.768-1.768c.471-.471 1.697-1.131 2.828 0c1.132 1.131.472 2.357 0 2.828l-2.242 2.243a2 2 0 0 0-.586 1.414z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 14.905c.705-1.234 1.825-2.32 3-3.204M2 22.404c1.072-3.002 3.055-5.564 5.023-7.5m.477 6.5c.721-1.442 1.96-3.077 3.087-4.405"/></g><defs><clipPath id="majesticonsAirplaneFlight20"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	airplaneFlightTwoLineInnerSVG     = `<g fill="none"><g clip-path="url(#majesticonsAirplaneFlight2Line0)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.071 7.586l-.828.828A2 2 0 0 1 11.828 9h-1.414c-.375 0-.735.149-1 .414v0a1.414 1.414 0 0 0 0 2l1.829 1.829a2 2 0 0 0 2.828 0l1.414-1.415l.964 2.121a1.346 1.346 0 0 0 2.178.395v0c.252-.252.394-.595.394-.952v-4.27a2 2 0 0 1 .586-1.415l2.242-2.243c.472-.47 1.132-1.697 0-2.828c-1.131-1.131-2.357-.471-2.828 0l-1.768 1.768m-3.182 3.182l-2.02-.674a1.448 1.448 0 0 1-.566-2.397v0a1.448 1.448 0 0 1 1.12-.421l4.648.31m-3.182 3.182l3.182-3.182M2 14.905c.705-1.234 1.825-2.32 3-3.204M2 22.404c1.072-3.002 3.055-5.564 5.023-7.5m.477 6.5c.721-1.442 1.96-3.077 3.087-4.405"/></g><defs><clipPath id="majesticonsAirplaneFlight2Line0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	airplaneLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 9h5.5a2.5 2.5 0 0 1 2.5 2.5v0a2.5 2.5 0 0 1-2.5 2.5h-4.02a1 1 0 0 0-.78.375l-3.4 4.25a1 1 0 0 1-.78.375H9.476a1 1 0 0 1-.928-1.371l.902-2.258A1 1 0 0 0 8.523 14H4.659a1 1 0 0 1-.919-.606l-2.143-5A1 1 0 0 1 2.517 7h1.865a1 1 0 0 1 .894.553l.448.894A1 1 0 0 0 6.618 9H10m5 0l-2.7-3.6a1 1 0 0 0-.8-.4H9.618a1 1 0 0 0-.894 1.447L10 9m5 0h-5"/>`
	alertCircleInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm0 5a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1zm1 9a1 1 0 1 0-2 0a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`
	alertCircleLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 8v5m0 3v0"/></g>`
	alignBottomInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 20h18"/><rect width="13" height="4" x="6" y="17" fill="currentColor" rx="2" transform="rotate(-90 6 17)"/><rect width="9" height="4" x="14" y="17" fill="currentColor" rx="2" transform="rotate(-90 14 17)"/></g>`
	alignBottomLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 20h18"/><rect width="13" height="4" x="6" y="17" rx="2" transform="rotate(-90 6 17)"/><rect width="9" height="4" x="14" y="17" rx="2" transform="rotate(-90 14 17)"/></g>`
	alignHorizontalCenterInnerSVG     = `<g fill="none"><path fill="currentColor" d="M4 8a2 2 0 0 1 2-2h12a2 2 0 1 1 0 4H6a2 2 0 0 1-2-2zm3 8a2 2 0 0 1 2-2h6a2 2 0 1 1 0 4H9a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6H6a2 2 0 1 0 0 4h6m0-4h6a2 2 0 1 1 0 4h-6m0-4V3m0 7v4m0 0H9a2 2 0 1 0 0 4h3m0-4h3a2 2 0 1 1 0 4h-3m0 0v3"/></g>`
	alignHorizontalCenterLineInnerSVG = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6H6a2 2 0 1 0 0 4h6m0-4h6a2 2 0 1 1 0 4h-6m0-4V3m0 7v4m0 0H9a2 2 0 1 0 0 4h3m0-4h3a2 2 0 1 1 0 4h-3m0 0v3"/>`
	alignLeftInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3v18"/><rect width="13" height="4" x="7" y="6" fill="currentColor" rx="2"/><rect width="9" height="4" x="7" y="14" fill="currentColor" rx="2"/></g>`
	alignLeftLineInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3v18"/><rect width="13" height="4" x="7" y="6" rx="2"/><rect width="9" height="4" x="7" y="14" rx="2"/></g>`
	alignRightInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 3v18"/><rect width="13" height="4" x="4" y="6" fill="currentColor" rx="2"/><rect width="9" height="4" x="8" y="14" fill="currentColor" rx="2"/></g>`
	alignRightLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 3v18"/><rect width="13" height="4" x="4" y="6" rx="2"/><rect width="9" height="4" x="8" y="14" rx="2"/></g>`
	alignTopInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 4h18"/><rect width="13" height="4" x="6" y="20" fill="currentColor" rx="2" transform="rotate(-90 6 20)"/><rect width="9" height="4" x="14" y="16" fill="currentColor" rx="2" transform="rotate(-90 14 16)"/></g>`
	alignTopLineInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 4h18"/><rect width="13" height="4" x="6" y="20" rx="2" transform="rotate(-90 6 20)"/><rect width="9" height="4" x="14" y="16" rx="2" transform="rotate(-90 14 16)"/></g>`
	alignVerticalCenterInnerSVG       = `<g fill="none"><path fill="currentColor" d="M8 20a2 2 0 0 1-2-2V6a2 2 0 1 1 4 0v12a2 2 0 0 1-2 2zm8-3a2 2 0 0 1-2-2V9a2 2 0 1 1 4 0v6a2 2 0 0 1-2 2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12v6a2 2 0 1 0 4 0v-6m-4 0V6a2 2 0 1 1 4 0v6m-4 0H3m7 0h4m0 0v3a2 2 0 1 0 4 0v-3m-4 0V9a2 2 0 1 1 4 0v3m0 0h3"/></g>`
	alignVerticalCenterLineInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12v6a2 2 0 1 0 4 0v-6m-4 0V6a2 2 0 1 1 4 0v6m-4 0H3m7 0h4m0 0v3a2 2 0 1 0 4 0v-3m-4 0V9a2 2 0 1 1 4 0v3m0 0h3"/>`
	analyticsInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6zm10 2a1 1 0 1 0-2 0v8a1 1 0 1 0 2 0V8zm-4 3a1 1 0 1 0-2 0v5a1 1 0 1 0 2 0v-5zm8 3a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2z" clip-rule="evenodd"/>`
	analyticsDeleteInnerSVG           = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 5l2-2m-2 2l-2-2m2 2l2 2m-2-2l-2 2"/><path fill="currentColor" fill-rule="evenodd" d="M21 10.659A6 6 0 0 1 13.341 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-7.341zM12 7a1 1 0 0 1 1 1v8a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1zm-4 3a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1zm8 3a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1z" clip-rule="evenodd"/></g>`
	analyticsDeleteLineInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16v-5m4 5V8m4 8v-2M13 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7m-1-6l2-2m-2 2l-2-2m2 2l2 2m-2-2l-2 2"/>`
	analyticsLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16v-5m4 5V8m4 8v-2m2-10H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/>`
	analyticsPlusInnerSVG             = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 2v3m0 3V5m0 0h3m-3 0h-3"/><path fill="currentColor" fill-rule="evenodd" d="M21 10.659A6 6 0 0 1 13.341 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-7.341zM12 7a1 1 0 0 1 1 1v8a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1zm-4 3a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1zm8 3a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1z" clip-rule="evenodd"/></g>`
	analyticsPlusLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16v-5m4 5V8m4 8v-2M12 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6M19 2v3m0 3V5m0 0h3m-3 0h-3"/>`
	analyticsRestrictedInnerSVG       = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.975 7.975a3.5 3.5 0 1 0-4.95-4.95m4.95 4.95a3.5 3.5 0 1 1-4.95-4.95m4.95 4.95L18.5 5.5l-2.475-2.475"/><path fill="currentColor" fill-rule="evenodd" d="M21 11.502a6.48 6.48 0 0 1-2.5.498A6.495 6.495 0 0 1 13 8.966V16a1 1 0 1 1-2 0V8a1 1 0 0 1 1.178-.984A6.515 6.515 0 0 1 12.498 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-6.498zM8 10a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1zm8 3a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1z" clip-rule="evenodd"/></g>`
	analyticsRestrictedLineInnerSVG   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 16v-5m4 5V8m4 8v-2M12 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6"/><path d="M20.975 7.975a3.5 3.5 0 1 0-4.95-4.95m4.95 4.95a3.5 3.5 0 1 1-4.95-4.95m4.95 4.95L18.5 5.5l-2.475-2.475"/></g>`
	annotationInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-2.172a1 1 0 0 0-.707.293l-2 2a3 3 0 0 1-4.242 0l-2-2A1 1 0 0 0 7.172 18H5a3 3 0 0 1-3-3V6zm5 1a1 1 0 0 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2H7z" fill="currentColor"/></g>`
	annotationLineInnerSVG            = `<g fill="currentColor"><path d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-2.172a1 1 0 0 0-.707.293l-2 2a3 3 0 0 1-4.242 0l-2-2A1 1 0 0 0 7.172 18H5a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h2.172a3 3 0 0 1 2.12.879l2 2a1 1 0 0 0 1.415 0l2-2A3 3 0 0 1 16.828 16H19a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H5zm1 3a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1z"/></g>`
	applicationsInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6zm0 10a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3H6zm10 0a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3h-2zm0-10a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-2z" fill="currentColor"/></g>`
	applicationsAddInnerSVG           = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6zm0 10a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-2zM13 6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3h-2a3 3 0 0 1-3-3V6zm5 8a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-2z" fill="currentColor"/></g>`
	applicationsAddLineInnerSVG       = `<g fill="currentColor"><path d="M6 3a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6zM5 6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6zm1 7a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3H6zm-1 3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-2zm8-10a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3h-2a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-2zm2 9a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-2z"/></g>`
	applicationsLineInnerSVG          = `<g fill="currentColor"><path d="M6 3a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6zM5 6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6zm1 7a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3H6zm-1 3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-2zm8 0a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3h-2a3 3 0 0 1-3-3v-2zm3-1a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-2zm0-12a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-2zm-1 3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1V6z"/></g>`
	archiveInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v1c0 .364-.097.706-.268 1H2.268A1.99 1.99 0 0 1 2 7V6zm1 4v8a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-8H3zm7 2a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4z" fill="currentColor"/></g>`
	archiveLineInnerSVG               = `<g fill="currentColor"><path d="M5 5a1 1 0 0 0-1 1v1h16V6a1 1 0 0 0-1-1H5zM2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v1a2 2 0 0 1-1 1.732V18a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V8.732A2 2 0 0 1 2 7V6zm3 3v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V9H5zm4 3a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1z"/></g>`
	arrowCircleDownInnerSVG           = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm10.707 3.707l3-3a1 1 0 0 0-1.414-1.414L13 12.586V9a1 1 0 1 0-2 0v3.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0z" fill="currentColor"/></g>`
	arrowCircleDownLineInnerSVG       = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm.707 13.707l3-3a1 1 0 0 0-1.414-1.414L13 12.586V9a1 1 0 1 0-2 0v3.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0z"/></g>`
	arrowCircleLeftInnerSVG           = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm6.293-.707l3-3a1 1 0 1 1 1.414 1.414L11.414 11H15a1 1 0 1 1 0 2h-3.586l1.293 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414z" fill="currentColor"/></g>`
	arrowCircleLeftLineInnerSVG       = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm-3.707 9.293l3-3a1 1 0 1 1 1.414 1.414L11.414 11H15a1 1 0 1 1 0 2h-3.586l1.293 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414z"/></g>`
	arrowCircleRightInnerSVG          = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm13.707-.707l-3-3a1 1 0 1 0-1.414 1.414L12.586 11H9a1 1 0 1 0 0 2h3.586l-1.293 1.293a1 1 0 0 0 1.414 1.414l3-3a1 1 0 0 0 0-1.414z" fill="currentColor"/></g>`
	arrowCircleRightLineInnerSVG      = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm3.707 9.293l-3-3a1 1 0 1 0-1.414 1.414L12.586 11H9a1 1 0 1 0 0 2h3.586l-1.293 1.293a1 1 0 0 0 1.414 1.414l3-3a1 1 0 0 0 0-1.414z"/></g>`
	arrowCircleUpInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm9.293-3.707l-3 3a1 1 0 1 0 1.414 1.414L11 11.414V15a1 1 0 1 0 2 0v-3.586l1.293 1.293a1 1 0 0 0 1.414-1.414l-3-3a1 1 0 0 0-1.414 0z" fill="currentColor"/></g>`
	arrowCircleUpLineInnerSVG         = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm-.707 6.293l-3 3a1 1 0 1 0 1.414 1.414L11 11.414V15a1 1 0 1 0 2 0v-3.586l1.293 1.293a1 1 0 0 0 1.414-1.414l-3-3a1 1 0 0 0-1.414 0z"/></g>`
	arrowDownInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 19l6-6m-6 6l-6-6m6 6V5"/>`
	arrowDownCircleInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm.707-12.707l3 3a1 1 0 0 1-1.414 1.414L13 11.414V15a1 1 0 1 1-2 0v-3.586l-1.293 1.293a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.414 0z" clip-rule="evenodd"/>`
	arrowDownCircleLineInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 15V9m0 0l3 3m-3-3l-3 3"/></g>`
	arrowDownLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 19l6-6m-6 6l-6-6m6 6V5"/>`
	arrowLeftInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 12l6-6m-6 6l6 6m-6-6h14"/>`
	arrowLeftCircleInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm.707-5.293l3-3a1 1 0 0 0-1.414-1.414L13 12.586V9a1 1 0 1 0-2 0v3.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`
	arrowLeftCircleLineInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 9v6m0 0l3-3m-3 3l-3-3"/></g>`
	arrowLeftLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 12l6-6m-6 6l6 6m-6-6h14"/>`
	arrowRightInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 12l-6-6m6 6l-6 6m6-6H5"/>`
	arrowRightCircleInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm-.707-12.707a1 1 0 0 1 1.414 1.415L11.414 11H15a1 1 0 1 1 0 2h-3.585l1.292 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.415l3-3z" clip-rule="evenodd"/>`
	arrowRightCircleLineInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M15 12H9m0 0l3-3m-3 3l3 3"/></g>`
	arrowRightLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 12l-6-6m6 6l-6 6m6-6H5"/>`
	arrowUpInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 5l6 6m-6-6l-6 6m6-6v14"/>`
	arrowUpCircleInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm.707-5.293a1 1 0 0 1-1.415-1.414L12.585 13H9a1 1 0 1 1 0-2h3.586l-1.293-1.293a1 1 0 0 1 1.415-1.414l3 3a1 1 0 0 1 0 1.414l-3 3z" clip-rule="evenodd"/>`
	arrowUpCircleLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M9 12h6m0 0l-3 3m3-3l-3-3"/></g>`
	arrowUpLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 5l6 6m-6-6l-6 6m6-6v14"/>`
	arrowsCollapseFullInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 15h-4m0 0v4m0-4l4 4M5 9h4m0 0V5m0 4L5 5m14 4h-4m0 0V5m0 4l4-4M5 15h4m0 0v4m0-4l-4 4"/>`
	arrowsCollapseFullLineInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 15h-4m0 0v4m0-4l4 4M5 9h4m0 0V5m0 4L5 5m14 4h-4m0 0V5m0 4l4-4M5 15h4m0 0v4m0-4l-4 4"/>`
	arrowsExpandInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 3a1 1 0 0 1 .707 1.707L7.414 6l2.293 2.293a1 1 0 0 1-1.414 1.414L6 7.414L4.707 8.707A1 1 0 0 1 3 8V4a1 1 0 0 1 1-1h4zm6.293 12.707a1 1 0 0 1 1.414-1.414L18 16.586l1.293-1.293A1 1 0 0 1 21 16v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-.707-1.707L16.586 18l-2.293-2.293zM8 21a1 1 0 0 0 .707-1.707L7.414 18l2.293-2.293a1 1 0 0 0-1.414-1.414L6 16.586l-1.293-1.293A1 1 0 0 0 3 16v4a1 1 0 0 0 1 1h4zm6.293-12.707a1 1 0 0 0 1.414 1.414L18 7.414l1.293 1.293A1 1 0 0 0 21 8V4a1 1 0 0 0-1-1h-4a1 1 0 0 0-.707 1.707L16.586 6l-2.293 2.293z" fill="currentColor"/></g>`
	arrowsExpandFullInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19h4m0 0v-4m0 4l-4-4M9 5H5m0 0v4m0-4l4 4m6-4h4m0 0v4m0-4l-4 4M9 19H5m0 0v-4m0 4l4-4"/>`
	arrowsExpandFullLineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19h4m0 0v-4m0 4l-4-4M9 5H5m0 0v4m0-4l4 4m6-4h4m0 0v4m0-4l-4 4M9 19H5m0 0v-4m0 4l4-4"/>`
	arrowsExpandLineInnerSVG          = `<g fill="currentColor"><path d="M9 4a1 1 0 0 1-1 1H6.414l3.293 3.293a1 1 0 0 1-1.414 1.414L5 6.414V8a1 1 0 1 1-2 0V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1zm8.586 15l-3.293-3.293a1 1 0 0 1 1.414-1.414L19 17.586V16a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1h-4a1 1 0 1 1 0-2h1.586zM9 20a1 1 0 0 0-1-1H6.414l3.293-3.293a1 1 0 0 0-1.414-1.414L5 17.586V16a1 1 0 1 0-2 0v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1zm8.586-15l-3.293 3.293a1 1 0 0 0 1.414 1.414L19 6.414V8a1 1 0 1 0 2 0V4a1 1 0 0 0-1-1h-4a1 1 0 1 0 0 2h1.586z"/></g>`
	articleInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm5 1a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H7zm1 4V9h2v2H8zm7-4a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2zm0 4a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2zm-8 4a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	articleLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8h2m-2 4h2m0 4H7m0-8v4h4V8H7zM5 20h14a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2z"/>`
	articleSearchInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M5 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h8.257a5.477 5.477 0 0 1-1.235-4H7a1 1 0 1 1 0-2h5.6a5.525 5.525 0 0 1 1.92-2.123A1 1 0 0 1 15 11h2a1 1 0 0 1 1 1.022a5.496 5.496 0 0 1 4 2.315V6a3 3 0 0 0-3-3H5zm1 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V8zm2 1v2h2V9H8zm6-1a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`
	articleSearchLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 8h2m-2 4h2m-5 4H7m14-4V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7M7 8v4h4V8H7z"/><path d="M19.268 19.268a2.5 2.5 0 1 0-3.536-3.536a2.5 2.5 0 0 0 3.536 3.536zm0 0L21 21"/></g>`
	atSymbolInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11zM4.788 6.416C6.174 4.492 8.474 3 12 3s5.826 1.492 7.212 3.416C20.56 8.289 21 10.506 21 12c0 .631-.152 1.603-.743 2.448C19.625 15.351 18.563 16 17 16c-1.037 0-1.74-.531-2.186-1.157A4 4 0 1 1 16 12c0 .403.086.993.302 1.442c.204.427.423.558.698.558c.937 0 1.375-.351 1.618-.698c.284-.405.382-.933.382-1.302c0-1.173-.36-2.956-1.412-4.416C16.575 6.175 14.874 5 12 5S7.426 6.175 6.412 7.584C5.36 9.044 5 10.827 5 12c0 1.173.36 2.956 1.412 4.416C7.426 17.825 9.126 19 12 19c1.98 0 3.385-.558 4.388-1.334a1 1 0 1 1 1.224 1.581C16.229 20.317 14.385 21 12 21c-3.526 0-5.826-1.491-7.212-3.416C3.44 15.711 3 13.494 3 12c0-1.493.44-3.711 1.788-5.584zM12 10a2 2 0 1 1 0 4a2 2 0 0 1 0-4z" fill="currentColor"/></g>`
	atSymbolLineInnerSVG              = `<g fill="currentColor"><path d="M6.412 7.584C5.36 9.044 5 10.827 5 12c0 1.173.36 2.956 1.412 4.416C7.426 17.825 9.126 19 12 19c1.98 0 3.385-.558 4.388-1.334a1 1 0 1 1 1.224 1.581C16.229 20.317 14.385 21 12 21c-3.526 0-5.826-1.491-7.212-3.416C3.44 15.711 3 13.494 3 12c0-1.493.44-3.711 1.788-5.584C6.174 4.492 8.474 3 12 3s5.826 1.492 7.212 3.416C20.56 8.289 21 10.506 21 12a3.79 3.79 0 0 1-.689 2.147c-.56.778-1.485 1.353-2.811 1.353c-1.326 0-2.251-.575-2.812-1.353A3.791 3.791 0 0 1 14 12a1 1 0 1 1 2 0c0 .257.085.664.311.978c.19.264.515.522 1.189.522s.999-.258 1.189-.522c.226-.314.311-.721.311-.978c0-1.173-.36-2.956-1.412-4.416C16.575 6.175 14.874 5 12 5S7.426 6.175 6.412 7.584z"/><path d="M12 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0z"/></g>`
	atomTwoInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.893 7.936a8.003 8.003 0 0 1-7.774 12.016m-6.012-3.888a8.003 8.003 0 0 1 7.774-12.016"/><circle cx="17.657" cy="6.343" r="2" fill="currentColor" transform="rotate(45 17.657 6.343)"/><circle cx="6.343" cy="17.657" r="2" fill="currentColor" transform="rotate(45 6.343 17.657)"/><circle cx="12" cy="12" r="2" fill="currentColor" transform="rotate(45 12 12)"/></g>`
	atomTwoLineInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.893 7.936a8.003 8.003 0 0 1-7.774 12.016m-6.012-3.888a8.003 8.003 0 0 1 7.774-12.016"/><circle cx="17.657" cy="6.343" r="2" transform="rotate(45 17.657 6.343)"/><circle cx="6.343" cy="17.657" r="2" transform="rotate(45 6.343 17.657)"/><circle cx="12" cy="12" r="2" transform="rotate(45 12 12)"/></g>`
	attachmentInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v5a2 2 0 0 0 2 2v0a2 2 0 0 0 2-2V7a4 4 0 0 0-4-4v0a4 4 0 0 0-4 4v8a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6V5"/>`
	attachmentLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v5a2 2 0 0 0 2 2v0a2 2 0 0 0 2-2V7a4 4 0 0 0-4-4v0a4 4 0 0 0-4 4v8a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6V5"/>`
	awardInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="9" r="7"/><path fill="currentColor" d="M7 20.234V14c.667.667 2.6 2 5 2s4.333-1.333 5-2v6.234a1 1 0 0 1-1.514.857l-2.972-1.782a1 1 0 0 0-1.028 0L8.514 21.09A1 1 0 0 1 7 20.234z"/></g>`
	awardLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="9" r="7"/><path d="M7 14v6.234a1 1 0 0 0 1.514.857l2.972-1.782a1 1 0 0 1 1.028 0l2.972 1.782A1 1 0 0 0 17 20.234V14"/></g>`
	backCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zm-7.496-4.868A1 1 0 0 1 17 8v8a1 1 0 0 1-1.496.868L9 13.152V16a1 1 0 1 1-2 0V8a1 1 0 1 1 2 0v2.848l6.504-3.716zM15 9.723L11.016 12L15 14.277V9.723z" clip-rule="evenodd"/>`
	backCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 8v8m8-8v8l-7-4l7-4z"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`
	backspaceInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M22 8a3 3 0 0 0-3-3H8.828a3 3 0 0 0-2.12.879l-4 4a3 3 0 0 0 0 4.242l4 4a3 3 0 0 0 2.12.879H19a3 3 0 0 0 3-3V8zm-6.707 1.293a1 1 0 1 1 1.414 1.414L15.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L14 13.414l-1.293 1.293a1 1 0 0 1-1.414-1.414L12.586 12l-1.293-1.293a1 1 0 0 1 1.414-1.414L14 10.586l1.293-1.293z" fill="currentColor"/></g>`
	backspaceLineInnerSVG             = `<g fill="currentColor"><path d="M19 7a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H8.828a1 1 0 0 1-.707-.293l-4-4a1 1 0 0 1 0-1.414l4-4A1 1 0 0 1 8.828 7H19zm3 1a3 3 0 0 0-3-3H8.828a3 3 0 0 0-2.12.879l-4 4a3 3 0 0 0 0 4.242l4 4a3 3 0 0 0 2.12.879H19a3 3 0 0 0 3-3V8zm-5.293 1.293a1 1 0 0 0-1.414 0L14 10.586l-1.293-1.293a1 1 0 1 0-1.414 1.414L12.586 12l-1.293 1.293a1 1 0 0 0 1.414 1.414L14 13.414l1.293 1.293a1 1 0 0 0 1.414-1.414L15.414 12l1.293-1.293a1 1 0 0 0 0-1.414z"/></g>`
	backwardCircleInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1zm5.493 7.13a1 1 0 0 0-1.008.013L12 10.833V9a1 1 0 0 0-1.514-.857l-5 3a1 1 0 0 0 0 1.714l5 3A1 1 0 0 0 12 15v-1.834l4.485 2.691A1 1 0 0 0 18 15V9a1 1 0 0 0-.507-.87zM16 13.234L13.944 12L16 10.766v2.468zM7.944 12L10 13.234v-2.468L7.944 12z" clip-rule="evenodd"/>`
	backwardCircleLineInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle r="10" transform="matrix(-1 0 0 1 12 12)"/><path d="M17 15V9l-5 3l5 3zm-6 0V9l-5 3l5 3z"/></g>`
	backwardStartCircleInnerSVG       = `<path fill="currentColor" fill-rule="evenodd" d="M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1zm5.493 7.13a1 1 0 0 0-1.008.013L13 10.233V9a1 1 0 0 0-1.514-.857L8 10.233V9a1 1 0 0 0-2 0v6a1 1 0 1 0 2 0v-1.234l3.486 2.091A1 1 0 0 0 13 15v-1.234l3.485 2.091A1 1 0 0 0 18 15V9a1 1 0 0 0-.507-.87zM11 12v-1.234L8.944 12L11 13.234V12zm5 1.234L13.944 12L16 10.766v2.468z" clip-rule="evenodd"/>`
	backwardStartCircleLineInnerSVG   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle r="10" transform="matrix(-1 0 0 1 12 12)"/><path d="M17 15V9l-5 3l5 3zM7 12l5 3V9l-5 3zm0 0V9m0 3v3"/></g>`
	badgeCheckInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10.054 2.344a3 3 0 0 1 3.892 0l1.271 1.084a1 1 0 0 0 .57.236l1.665.133a3 3 0 0 1 2.751 2.751l.133 1.666a1 1 0 0 0 .236.569l1.084 1.271a3 3 0 0 1 0 3.892l-1.084 1.271a1 1 0 0 0-.236.57l-.133 1.665a3 3 0 0 1-2.751 2.751l-1.666.133a1 1 0 0 0-.569.236l-1.271 1.084a3 3 0 0 1-3.892 0l-1.271-1.084a1 1 0 0 0-.57-.236l-1.665-.133a3 3 0 0 1-2.751-2.751l-.133-1.666a1 1 0 0 0-.236-.569l-1.084-1.271a3 3 0 0 1 0-3.892l1.084-1.271a1 1 0 0 0 .236-.57l.133-1.665a3 3 0 0 1 2.751-2.751l1.666-.133a1 1 0 0 0 .569-.236l1.271-1.084zm5.653 8.363a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4z" fill="currentColor"/></g>`
	badgeCheckLineInnerSVG            = `<g fill="currentColor"><path d="M13.946 2.344a3 3 0 0 0-3.892 0L8.783 3.428a1 1 0 0 1-.57.236l-1.665.132a3 3 0 0 0-2.751 2.752l-.133 1.666a1 1 0 0 1-.236.569l-1.084 1.271a3 3 0 0 0 0 3.892l1.084 1.271a1 1 0 0 1 .236.57l.133 1.665a3 3 0 0 0 2.751 2.751l1.666.133c.21.017.409.1.569.236l1.271 1.084a3 3 0 0 0 3.892 0l1.271-1.084c.16-.136.36-.219.57-.236l1.665-.133a3 3 0 0 0 2.751-2.751l.133-1.666a1 1 0 0 1 .236-.569l1.084-1.271a3 3 0 0 0 0-3.892l-1.084-1.271a1 1 0 0 1-.236-.57l-.133-1.665a3 3 0 0 0-2.751-2.752l-1.666-.132a1 1 0 0 1-.569-.236l-1.271-1.084zM11.35 3.867a1 1 0 0 1 1.298 0L13.92 4.95a3 3 0 0 0 1.707.707l1.665.133a1 1 0 0 1 .918.917l.133 1.666a3 3 0 0 0 .707 1.707l1.083 1.271a1 1 0 0 1 0 1.298L19.05 13.92a3 3 0 0 0-.707 1.707l-.133 1.665a1 1 0 0 1-.918.918l-1.665.133a3 3 0 0 0-1.707.707l-1.271 1.083a1 1 0 0 1-1.298 0L10.08 19.05a3 3 0 0 0-1.707-.707l-1.666-.133a1 1 0 0 1-.917-.918l-.133-1.665a3 3 0 0 0-.707-1.707l-1.083-1.271a1 1 0 0 1 0-1.298L4.95 10.08a3 3 0 0 0 .707-1.707l.133-1.666a1 1 0 0 1 .917-.917l1.666-.133a3 3 0 0 0 1.707-.707l1.271-1.083zm4.356 6.84a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 1 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4z"/></g>`
	banInnerSVG                       = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm6 10c0 1.296-.41 2.496-1.11 3.476L8.524 7.11A6 6 0 0 1 18 12zM6 12c0-1.296.41-2.496 1.11-3.477l8.366 8.368A6 6 0 0 1 6 12z" fill="currentColor"/></g>`
	banLineInnerSVG                   = `<g fill="currentColor"><path d="M7.094 5.68L18.32 16.906A8 8 0 0 0 7.094 5.68zm9.812 12.64L5.68 7.094A8 8 0 0 0 16.906 18.32zM4.929 4.929A9.972 9.972 0 0 1 12 2c5.523 0 10 4.477 10 10a9.972 9.972 0 0 1-2.929 7.071A9.972 9.972 0 0 1 12 22C6.477 22 2 17.523 2 12a9.972 9.972 0 0 1 2.929-7.071z"/></g>`
	bandAidsInnerSVG                  = `<g fill="none"><rect width="9" height="19" x="2" y="8.364" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="4.5" transform="rotate(-45 2 8.364)"/><path fill="currentColor" d="m11.546 18.617l.354-.353L5.535 11.9l-.354.353a4.5 4.5 0 0 0 6.364 6.364zm.707-13.435l-.354.354l6.365 6.364l.353-.354a4.5 4.5 0 1 0-6.364-6.364z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.546 18.617l.354-.353L5.535 11.9l-.354.353a4.5 4.5 0 0 0 6.364 6.364zm.707-13.435l-.354.354l6.365 6.364l.353-.354a4.5 4.5 0 1 0-6.364-6.364zM9.779 11.9h0M11.9 9.778h0m2.121 2.122h0M11.9 14.021h0"/></g>`
	bandAidsLineInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="9" height="19" x="2" y="8.364" rx="4.5" transform="rotate(-45 2 8.364)"/><path d="m11.9 18.264l-.354.353a4.5 4.5 0 0 1-6.364 0v0a4.5 4.5 0 0 1 0-6.364l.354-.353M11.9 5.536l.353-.354a4.5 4.5 0 0 1 6.364 0v0a4.5 4.5 0 0 1 0 6.364l-.354.354m-8.484 0h0M11.9 9.778h0m2.121 2.122h0M11.9 14.021h0"/></g>`
	barcodeTwoInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6v12m7-12v12m7-12v12M10 6.5v11a.5.5 0 0 0 1 0v-11a.5.5 0 0 0-1 0zm7 0v11a.5.5 0 0 0 1 0v-11a.5.5 0 0 0-1 0zM3.5 6v0a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5v0a.5.5 0 0 0 .5-.5v-11a.5.5 0 0 0-.5-.5z"/>`
	barcodeTwoLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6v12m7-12v12m7-12v12M10 6.5v11a.5.5 0 0 0 1 0v-11a.5.5 0 0 0-1 0zm7 0v11a.5.5 0 0 0 1 0v-11a.5.5 0 0 0-1 0zM3.5 6v0a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5v0a.5.5 0 0 0 .5-.5v-11a.5.5 0 0 0-.5-.5z"/>`
	basketTwoInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M9.612 6.084C9.16 6.711 9 7.494 9 8v1h6V8c0-.507-.16-1.289-.611-1.916C13.974 5.508 13.274 5 12 5c-1.274 0-1.974.508-2.388 1.084zM17 9V8c0-.827-.24-2.044-.988-3.084C15.226 3.825 13.926 3 12 3c-1.926 0-3.226.825-4.012 1.916C7.24 5.956 7 7.173 7 8v1H3a1 1 0 0 0 0 2h.095l.91 9.1A1 1 0 0 0 5 21h14a1 1 0 0 0 .995-.9l.91-9.1H21a1 1 0 1 0 0-2h-4zm-8 5a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2zm4 0a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2zm4 0a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2z" clip-rule="evenodd"/>`
	basketTwoLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h1m17 0h-1m0 0l-1 10H5L4 10m16 0h-4M4 10h4m4 4v2m3-2v2m-6-2v2m-1-6h8m-8 0V8c0-1.333.8-4 4-4s4 2.667 4 4v2"/>`
	bathShowerInnerSVG                = `<g fill="none"><path fill="currentColor" d="M20 14v-2H4v2c0 1.138.583 3.248 2.745 3.841c.37.102.787.159 1.255.159h8a4.71 4.71 0 0 0 1.255-.159C19.417 17.248 20 15.138 20 14z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h1m16 0v0a1 1 0 0 0 1-1V7c0-1-.6-3-3-3s-3 2-3 3m5 5v2c0 1.138-.583 3.248-2.745 3.841M20 12H4m0 0v2c0 1.138.583 3.248 2.745 3.841M6 20l.745-2.159m0 0c.37.102.787.159 1.255.159h8a4.71 4.71 0 0 0 1.255-.159M18 20l-.745-2.159M15 7h-2m2 0h2"/></g>`
	bathShowerLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h1m16 0v0a1 1 0 0 0 1-1V7c0-1-.6-3-3-3s-3 2-3 3m5 5v2c0 1.138-.583 3.248-2.745 3.841M20 12H4m0 0v2c0 1.138.583 3.248 2.745 3.841M6 20l.745-2.159m0 0c.37.102.787.159 1.255.159h8a4.71 4.71 0 0 0 1.255-.159M18 20l-.745-2.159M15 7h-2m2 0h2"/>`
	batteryInnerSVG                   = `<g fill="none"><path fill="currentColor" d="M16 6H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zm6 7v-2a1 1 0 0 0-1-1v4a1 1 0 0 0 1-1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 6H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zm6 7v-2a1 1 0 0 0-1-1v4a1 1 0 0 0 1-1z"/></g>`
	batteryFullInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1 8a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8zm19 2a1 1 0 0 1 1-1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2a1 1 0 0 1-1-1v-4zM7 10a1 1 0 0 0-2 0v4a1 1 0 1 0 2 0v-4zm3-1a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1zm5 1a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4z" clip-rule="evenodd"/>`
	batteryFullLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 10v4m4 0v-4m4 0v4m4-4V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2m0-4v0a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2v0m0-4v4"/>`
	batteryHalfInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M4 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3H4zm17 4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2zM6 9a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1zm5 1a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4z" clip-rule="evenodd"/>`
	batteryHalfLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 10v4m4 0v-4m8 0V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2m0-4v0a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2v0m0-4v4"/>`
	batteryLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 10V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2m0-4v0a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2v0m0-4v4"/>`
	batteryLowInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1 8a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8zm19 2a1 1 0 0 1 1-1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2a1 1 0 0 1-1-1v-4zM7 10a1 1 0 0 0-2 0v4a1 1 0 1 0 2 0v-4z" clip-rule="evenodd"/>`
	batteryLowLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 10v4m12-4V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2m0-4v0a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2v0m0-4v4"/>`
	beachInnerSVG                     = `<g fill="none"><path fill="currentColor" d="M10 5.196c1.5-2.598 5.098-2.83 7.696-1.33s4.196 4.732 2.696 7.33l-3.464-2l-1.732-1l-1.732-1l-3.464-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.696 3.866C15.098 2.366 11.5 2.598 10 5.196l3.464 2m4.232-3.33c2.598 1.5 4.196 4.732 2.696 7.33l-5.196-3m2.5-4.33l.5-.866m-.5.866c-1.821.488-2.982 1.165-4.232 3.33m4.232-3.33c.488 1.821.482 3.165-.768 5.33m-1.732-1l-1.732-1m1.732 1l-3 5.196M3 21l.88-1.056a2.001 2.001 0 0 1 3.139.08v0a2.001 2.001 0 0 0 3.107.118l.19-.218a2.236 2.236 0 0 1 3.367 0l.191.218c.838.957 2.344.9 3.107-.117v0a2.001 2.001 0 0 1 3.14-.08L21 21M6.708 16A7.97 7.97 0 0 1 12 14a7.97 7.97 0 0 1 5.292 2"/></g>`
	beachLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.696 3.866C15.098 2.366 11.5 2.598 10 5.196l3.464 2m4.232-3.33c2.598 1.5 4.196 4.732 2.696 7.33l-5.196-3m2.5-4.33l.5-.866m-.5.866c-1.821.488-2.982 1.165-4.232 3.33m4.232-3.33c.488 1.821.482 3.165-.768 5.33m-1.732-1l-1.732-1m1.732 1l-3 5.196M3 21l.88-1.056a2.001 2.001 0 0 1 3.139.08v0a2.001 2.001 0 0 0 3.107.118l.19-.218a2.236 2.236 0 0 1 3.367 0l.191.218c.838.957 2.344.9 3.107-.117v0a2.001 2.001 0 0 1 3.14-.08L21 21M6.708 16A7.97 7.97 0 0 1 12 14a7.97 7.97 0 0 1 5.292 2"/>`
	beakerInnerSVG                    = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 3a1 1 0 0 0 0 2v5.62l-4.789 5.387C1.491 17.942 2.865 21 5.454 21h13.092c2.589 0 3.962-3.058 2.242-4.993L16 10.62V5a1 1 0 1 0 0-2H8z" fill="currentColor"/></g>`
	beakerLineInnerSVG                = `<g fill="currentColor"><path d="M7 4a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2v5.62l4.788 5.387c1.72 1.935.347 4.993-2.242 4.993H5.454c-2.589 0-3.962-3.058-2.243-4.993L8 10.62V5a1 1 0 0 1-1-1zm3 1v6a1 1 0 0 1-.253.664l-5.04 5.672C4.132 17.98 4.59 19 5.453 19h13.092c.863 0 1.321-1.02.748-1.664l-5.041-5.672A1 1 0 0 1 14 11V5h-4z"/></g>`
	bellInnerSVG                      = `<g fill="none"><path fill="currentColor" d="M6 11c0-4.8 4-6 6-6c4.8 0 6 4 6 6v4l2 2H4l2-2v-4z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5c-2 0-6 1.2-6 6v4l-2 2h16l-2-2v-4c0-2-1.2-6-6-6zm0 0V3M9 18c0 1 .6 3 3 3s3-2 3-3"/></g>`
	bellLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5c-2 0-6 1.2-6 6v4l-2 2h5m3-12c4.8 0 6 4 6 6v4l2 2h-5M12 5V3M9 17v1c0 1 .6 3 3 3s3-2 3-3v-1m-6 0h6"/>`
	bitcoinCircleInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm10-6a1 1 0 1 0-2 0v1H8a1 1 0 0 0 0 2h1v6H8a1 1 0 1 0 0 2h1v1a1 1 0 1 0 2 0v-1h1v1a1 1 0 1 0 2 0v-1c.493 0 1.211-.14 1.834-.588C16.51 15.925 17 15.126 17 14c0-.851-.281-1.516-.71-2c.429-.484.71-1.149.71-2c0-1.126-.491-1.926-1.166-2.412A3.233 3.233 0 0 0 14 7V6a1 1 0 1 0-2 0v1h-1V6zm0 5V9h3c.173 0 .456.06.666.212c.159.114.334.314.334.788c0 .474-.175.674-.334.789A1.25 1.25 0 0 1 14 11h-3zm0 2h3c.173 0 .456.06.666.211c.159.115.334.315.334.789c0 .474-.175.674-.334.789A1.25 1.25 0 0 1 14 15h-3v-2z" clip-rule="evenodd"/>`
	bitcoinCircleLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M8 8h5m1 4c.667 0 2-.4 2-2s-1.333-2-2-2h-1m1 4c.667 0 2 .4 2 2s-1.333 2-2 2h-1m1-4h-4m-2 4h5M10 6v6m0 6v-6m3-6v2m0 8v2"/></g>`
	bluetoothInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 12l5-4l-5-4v8zm0 0l5 4l-5 4v-8zm0 0L7 8m5 4l-5 4"/>`
	bluetoothLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 12l5-4l-5-4v8zm0 0l5 4l-5 4v-8zm0 0L7 8m5 4l-5 4"/>`
	boldInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 6H6v6h7a3 3 0 1 0 0-6zm2 6H6v6h9a3 3 0 1 0 0-6z"/>`
	boldLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 6H6v6h7a3 3 0 1 0 0-6zm2 6H6v6h9a3 3 0 1 0 0-6z"/>`
	bookInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M3 5a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V5zm5 0v7l2.293-2.293a1 1 0 0 1 1.414 0L14 12V5a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1z" clip-rule="evenodd"/>`
	bookLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-4M8 3v9l3-3l3 3V3M8 3h6"/>`
	bookMinusInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 19h-6"/><path fill="currentColor" fill-rule="evenodd" d="M6 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7.803A6 6 0 0 1 21 13.341V5a3 3 0 0 0-3-3H6zm2 10V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v7l-2.293-2.293a1 1 0 0 0-1.414 0L8 12z" clip-rule="evenodd"/></g>`
	bookMinusLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6M8 3v9l3-3l3 3V3M8 3h6m0 0h4a2 2 0 0 1 2 2v7m2 7h-6"/>`
	bookOpenInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M11 4.85c-2.195-1.022-5.52-1.565-8.633.979A1 1 0 0 0 2 6.603V19a1 1 0 0 0 1.633.774c2.736-2.236 5.734-1.31 7.367-.255V4.849zm2 0v14.669c1.633-1.056 4.63-1.981 7.367.255A1 1 0 0 0 22 19V6.603a1 1 0 0 0-.367-.774C18.52 3.285 15.195 3.828 13 4.849z" clip-rule="evenodd"/>`
	bookOpenLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.603c1.667-1.271 5.5-2.86 9 0V19c-3.5-2.86-7.333-1.271-9 0m0-12.397c-1.667-1.271-5.5-2.86-9 0V19c3.5-2.86 7.333-1.271 9 0m0-12.397V19"/>`
	bookPlusInnerSVG                  = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 16v3m0 3v-3m0 0h3m-3 0h-3"/><path fill="currentColor" fill-rule="evenodd" d="M6 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7.803A6 6 0 0 1 21 13.341V5a3 3 0 0 0-3-3H6zm2 10V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v7l-2.293-2.293a1 1 0 0 0-1.414 0L8 12z" clip-rule="evenodd"/></g>`
	bookPlusLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6M8 3v9l3-3l3 3V3M8 3h6m0 0h4a2 2 0 0 1 2 2v7m-1 4v3m0 3v-3m0 0h3m-3 0h-3"/>`
	bookmarkInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7 2a3 3 0 0 0-3 3v15.138a1.5 1.5 0 0 0 2.244 1.303l5.26-3.006a1 1 0 0 1 .992 0l5.26 3.006A1.5 1.5 0 0 0 20 20.138V5a3 3 0 0 0-3-3H7z" clip-rule="evenodd"/>`
	bookmarkBookInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-6v8.755a.5.5 0 0 1-.832.374l-1.836-1.632a.5.5 0 0 0-.664 0L7.832 12.13A.5.5 0 0 1 7 11.755V3H5z" fill="currentColor"/></g>`
	bookmarkBookLineInnerSVG          = `<g fill="currentColor"><path d="M6 5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-3v9a1 1 0 0 1-1.555.832L11 13.202l-2.445 1.63A1 1 0 0 1 7 14V5H6zm3 0v7.132l1.445-.964a1 1 0 0 1 1.11 0l1.445.964V5H9zM3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6z"/></g>`
	bookmarkLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3H7a2 2 0 0 0-2 2v15.138a.5.5 0 0 0 .748.434l5.26-3.005a2 2 0 0 1 1.984 0l5.26 3.006a.5.5 0 0 0 .748-.435V5a2 2 0 0 0-2-2z"/>`
	bookmarkMinusInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M4 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v15.138a1.5 1.5 0 0 1-2.244 1.303l-5.26-3.006a1 1 0 0 0-.992 0l-5.26 3.006A1.5 1.5 0 0 1 4 20.138V5zm11 4a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h6z" clip-rule="evenodd"/>`
	bookmarkMinusLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3H7a2 2 0 0 0-2 2v15.138a.5.5 0 0 0 .748.434l5.26-3.005a2 2 0 0 1 1.984 0l5.26 3.006a.5.5 0 0 0 .748-.435V5a2 2 0 0 0-2-2zm-8 7h6"/>`
	bookmarkPlusInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M4 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v15.138a1.5 1.5 0 0 1-2.244 1.303l-5.26-3.006a1 1 0 0 0-.992 0l-5.26 3.006A1.5 1.5 0 0 1 4 20.138V5zm7 2a1 1 0 1 1 2 0v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9a1 1 0 1 1 0-2h2V7z" clip-rule="evenodd"/>`
	bookmarkPlusLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3H7a2 2 0 0 0-2 2v15.138a.5.5 0 0 0 .748.434l5.26-3.005a2 2 0 0 1 1.984 0l5.26 3.006a.5.5 0 0 0 .748-.435V5a2 2 0 0 0-2-2zm-5 4v3m0 3v-3m0 0H9m3 0h3"/>`
	boxInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M11.47 2.152a1 1 0 0 1 1.06 0l6.904 4.315L12 10.84L4.566 6.467l6.904-4.315zM3.008 7.871A1.001 1.001 0 0 0 3 8v8a1 1 0 0 0 .47.848L11 21.554v-8.982L3.008 7.87zM13 21.554l7.53-4.706A1 1 0 0 0 21 16V8c0-.043-.003-.087-.008-.129L13 12.571v8.983z" clip-rule="evenodd"/>`
	boxLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 8l8-5l4 2.5M4 8v8l8 5M4 8l4 2.5m4 2.5l8-5m-8 5v8m0-8l-4-2.5M20 8v8l-8 5m8-13l-4-2.5m-8 5l8-5"/>`
	briefcaseInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M8 6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v1h3a3 3 0 0 1 3 3v2h-7v-1a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v1H2v-2a3 3 0 0 1 3-3h3V6zm-6 8v4a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4h-7v1a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-1H2zm8-7h4V6a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v1zm3 7h-2v-2h2v2z" clip-rule="evenodd"/>`
	briefcaseLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8V6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v2m6 0h4a2 2 0 0 1 2 2v3m-6-5H9m0 0H5a2 2 0 0 0-2 2v3m0 0v5a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5M3 13h7m11 0h-7m0 0v-2h-4v2m4 0v2h-4v-2"/>`
	browserInnerSVG                   = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8M3 8h18"/><path fill="currentColor" d="M3 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2H3V6z"/></g>`
	browserCookieInnerSVG             = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v10a2 2 0 0 0 2 2h4M3 8h18m0 0v2"/><path fill="currentColor" fill-rule="evenodd" d="M18 12v1h2a1 1 0 0 1 1 1v2h1a1 1 0 0 1 1 1a6 6 0 1 1-6-6a1 1 0 0 1 1 1zm-4 4a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H15a1 1 0 0 1-1-1zm4 2a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H18z" clip-rule="evenodd"/><path fill="currentColor" d="M3 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2H3V6z"/></g>`
	browserCookieLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v10a2 2 0 0 0 2 2h4M3 8h18m0 0v2m-6 7h.001M18 19h.001M22 17a5 5 0 1 1-5-5v2h3v3h2z"/>`
	browserLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8M3 8h18"/>`
	bugTwoInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.293 3.293a1 1 0 0 1 1.414 0l1.876 1.876A6.39 6.39 0 0 1 12 5c.515 0 .996.049 1.445.14l1.848-1.847a1 1 0 1 1 1.414 1.414L15.45 5.965A5.5 5.5 0 0 1 17.249 8H18c.173 0 .456-.06.666-.212c.159-.114.334-.314.334-.788a1 1 0 1 1 2 0c0 1.126-.491 1.926-1.166 2.412A3.233 3.233 0 0 1 18 10h-.086c.06.36.086.7.086 1v1h2a1 1 0 1 1 0 2h-2v1c0 .3-.026.64-.086 1H18c.493 0 1.211.14 1.834.588C20.51 17.075 21 17.875 21 19a1 1 0 1 1-2 0c0-.474-.175-.674-.334-.788A1.239 1.239 0 0 0 18 18h-.751a5.537 5.537 0 0 1-1.552 1.857C14.766 20.563 13.543 21 12 21c-1.543 0-2.765-.437-3.697-1.143c-.7-.53-1.2-1.188-1.552-1.857H6c-.173 0-.456.06-.666.212c-.159.114-.334.314-.334.788a1 1 0 1 1-2 0c0-1.126.492-1.926 1.166-2.412A3.233 3.233 0 0 1 6 16h.086c-.06-.36-.086-.7-.086-1v-1H4a1 1 0 1 1 0-2h2v-1c0-.349.022-.682.065-1H6c-.493 0-1.211-.14-1.834-.588C3.492 8.926 3 8.126 3 7a1 1 0 0 1 2 0c0 .474.175.674.334.788c.21.152.493.212.666.212h.696A5.34 5.34 0 0 1 8.58 5.994L7.293 4.707a1 1 0 0 1 0-1.414zM12 9a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H12zm-3 4a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1zm5-1a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H14z" clip-rule="evenodd"/>`
	bugTwoLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 4l-2.251 2.251m0 0A5.782 5.782 0 0 0 12 6a5.56 5.56 0 0 0-1.711.289m3.46-.038c1.551.49 2.417 1.623 2.858 2.749M8 4l2.289 2.289m0 0C9.135 6.67 7.914 7.479 7.339 9m9.268 0A5.66 5.66 0 0 1 17 11v2m-.393-4H18c.667 0 2-.4 2-2m-3 6v2a5.66 5.66 0 0 1-.393 2M17 13h3m-3.393 4c-.585 1.494-1.918 3-4.607 3s-4.022-1.506-4.607-3m9.214 0H18c.667 0 2 .4 2 2M7.338 9C7.125 9.564 7 10.226 7 11v2m.338-4H6c-.667 0-2-.4-2-2m3 6v2c0 .546.107 1.272.393 2M7 13H4m3.393 4H6c-.667 0-2 .4-2 2m8-9h.001M10 13h.001M14 13h.001"/>`
	burgerInnerSVG                    = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 15l3.172-3.172a2.83 2.83 0 0 1 2-.828H20a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v0a2 2 0 0 1 2-2h3.657c1.5 0 2.939.596 4 1.657L14 15z"/><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15h18v2a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-2z"/><path fill="currentColor" fill-rule="evenodd" d="M5.467 4.554C7.27 3.457 9.58 3 12 3s4.73.457 6.533 1.554c1.844 1.121 3.109 2.886 3.402 5.328C22.082 11.106 21.067 12 20 12H4c-1.067 0-2.082-.894-1.935-2.118c.293-2.442 1.558-4.207 3.402-5.328zM7 6a1 1 0 0 0 0 2h.001a1 1 0 0 0 0-2H7zm8 1a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H16a1 1 0 0 1-1-1zm-4 1a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H11z" clip-rule="evenodd"/></g>`
	burgerLineInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 15l3.379-3.379a2.121 2.121 0 0 1 1.5-.621H20a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v0a2 2 0 0 1 2-2h4.515a6 6 0 0 1 4.242 1.757L15 15zM3 15h18v2a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-2z"/><path d="M12 4c-4.623 0-8.432 1.756-8.942 6c-.066.55.39 1 .942 1h16c.552 0 1.008-.45.942-1c-.51-4.244-4.319-6-8.942-6zM7.001 8H7m8.001-1H15m-2.999 1H12"/></g>`
	busInnerSVG                       = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 19v0a2 2 0 0 1-2-2v-4m2 6v1.5a.5.5 0 0 0 .5.5v0a.5.5 0 0 0 .5-.5V19m-1 0h1m11 0v0a2 2 0 0 0 2-2v-4m-2 6v1.5a.5.5 0 0 1-.5.5v0a.5.5 0 0 1-.5-.5V19m1 0h-1M4 13V7m0 6h8m8 0V7m0 6h-8m-5 6h10M4 7V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2M4 7h8m8 0h-8m0 0v6"/><path fill="currentColor" fill-rule="evenodd" d="M4 13h16v6H4v-6zm3 3a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1zm9-1a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H16z" clip-rule="evenodd"/><path fill="currentColor" d="M4 5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2H4V5z"/></g>`
	busLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 19v0a2 2 0 0 1-2-2v-4m2 6v1.5a.5.5 0 0 0 .5.5v0a.5.5 0 0 0 .5-.5V19m-1 0h1m11 0v0a2 2 0 0 0 2-2v-4m-2 6v1.5a.5.5 0 0 1-.5.5v0a.5.5 0 0 1-.5-.5V19m1 0h-1M4 13V7m0 6h8m8 0V7m0 6h-8m-5 6h10M4 7v0a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v0M4 7h8m8 0h-8m0 0v6m-4 3h.001M16 16h.001"/>`
	cakeInnerSVG                      = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3h.01M7 3h.01M17 3h.01"/><path fill="currentColor" fill-rule="evenodd" d="M7 5a1 1 0 0 1 1 1v2h3V6a1 1 0 1 1 2 0v2h3V6a1 1 0 1 1 2 0v2.1a5.005 5.005 0 0 1 3.937 4.102c-.4.155-.75.383-1.047.63c-.532.44-.966.994-1.302 1.46c-.265.367-.714.708-1.588.708c-.874 0-1.324-.342-1.588-.71A2.355 2.355 0 0 1 16 13a1 1 0 1 0-2 0c0 .34-.11.872-.412 1.29c-.264.368-.714.71-1.588.71c-.874 0-1.324-.342-1.588-.71A2.355 2.355 0 0 1 10 13a1 1 0 1 0-2 0c0 .34-.11.872-.412 1.29c-.264.368-.714.71-1.588.71c-.874 0-1.323-.341-1.588-.709c-.336-.465-.77-1.019-1.302-1.46a3.822 3.822 0 0 0-1.047-.629A5.005 5.005 0 0 1 6 8.1V6a1 1 0 0 1 1-1zm-5 9.52V19a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4.48c-.27.256-.532.583-.79.94c-.635.883-1.685 1.54-3.21 1.54c-1.367 0-2.353-.529-3-1.273c-.647.744-1.633 1.273-3 1.273s-2.353-.529-3-1.273C8.353 16.47 7.367 17 6 17c-1.525 0-2.575-.657-3.21-1.54a6.756 6.756 0 0 0-.79-.94z" clip-rule="evenodd"/></g>`
	cakeLineInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 9v0a4 4 0 0 0-4 4v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6a4 4 0 0 0-4-4v0M7 9h10M7 9V6m10 3V6m-5 0v3m0-6h.01M7 3h.01M17 3h.01"/><path d="M3 13c0 1 .6 3 3 3s3-2 3-3c0 1 .6 3 3 3s3-2 3-3c0 1 .6 3 3 3s3-2 3-3"/></g>`
	calculatorInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M7 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3H7zm0 4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V6zm1 7a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H8zm3 1a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H12a1 1 0 0 1-1-1zm1 2a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H12zm-5 1a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1zm10-3a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0v-3z" clip-rule="evenodd"/>`
	calculatorLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 14v3M9 7v4h6V7H9zM7 21h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2zm2-7h.01M12 14h.01M12 17h.01M9 17h.01"/>`
	calendarInnerSVG                  = `<g fill="none"><path fill="currentColor" d="M4 7v2h16V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 5h2a2 2 0 0 1 2 2v2H4V7a2 2 0 0 1 2-2h2m8 0V3m0 2H8m0-2v2M4 9.5V19a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9.5"/></g>`
	calendarLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9v10a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9M4 9V7a2 2 0 0 1 2-2h2M4 9h16m0 0V7a2 2 0 0 0-2-2h-2m0 0V3m0 2H8m0-2v2"/>`
	calendarPlusInnerSVG              = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 16v3m0 3v-3m0 0h3m-3 0h-3"/><path fill="currentColor" d="M4 7v2h16V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 5h2a2 2 0 0 1 2 2v2H4V7a2 2 0 0 1 2-2h2m8 0V3m0 2H8m0-2v2M4 9.5V19a2 2 0 0 0 2 2h7m7-11.5v2.75"/></g>`
	calendarPlusLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9v10a2 2 0 0 0 2 2h6M4 9V7a2 2 0 0 1 2-2h2M4 9h16m0 0V7a2 2 0 0 0-2-2h-2m4 4v3m-4-7V3m0 2H8m0-2v2m11 11v3m0 3v-3m0 0h3m-3 0h-3"/>`
	cameraInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.574 4.336A3 3 0 0 1 10.07 3h3.86a3 3 0 0 1 2.496 1.336l.812 1.219A1 1 0 0 0 18.07 6H19a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h.93a1 1 0 0 0 .832-.445l.812-1.22zM10 13a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm2-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8z" clip-rule="evenodd"/>`
	cameraLineInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 18V9a2 2 0 0 1 2-2h.93a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 10.07 4h3.86a2 2 0 0 1 1.664.89l.812 1.22A2 2 0 0 0 18.07 7H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><circle cx="12" cy="13" r="3"/></g>`
	cameraOffInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18"/><path fill="currentColor" fill-rule="evenodd" d="M5.172 6H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14c.35 0 .688-.06 1-.17l-5.087-5.089a4 4 0 1 1-5.654-5.654L5.172 6zm5.502 5.503a2 2 0 1 0 2.823 2.823l-2.823-2.823zM22 17.172V9a3 3 0 0 0-3-3h-.93a1 1 0 0 1-.832-.445l-.812-1.22A3 3 0 0 0 13.93 3h-3.86a3 3 0 0 0-1.708.534L22 17.172z" clip-rule="evenodd"/></g>`
	cameraOffLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h2m2-3h4.93a2 2 0 0 1 1.664.89l.812 1.22A2 2 0 0 0 18.07 7H19a2 2 0 0 1 2 2v6.5M9.5 9.877a4 4 0 1 0 5.5 5.768M3 3l18 18"/>`
	carInnerSVG                       = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 11l2.48-5.788A2 2 0 0 1 7.32 4h9.362a2 2 0 0 1 1.838 1.212L21 11M3 11h18M3 11v7m18-7v7m-3 0v.5a1.5 1.5 0 0 0 1.5 1.5v0a1.5 1.5 0 0 0 1.5-1.5V18m-3 0H6m12 0h3M6 18v.5A1.5 1.5 0 0 1 4.5 20v0A1.5 1.5 0 0 1 3 18.5V18m3 0H3"/><path fill="currentColor" fill-rule="evenodd" d="M3 11h18v7H3v-7zm3 3a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1zm11-1a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H17z" clip-rule="evenodd"/></g>`
	carLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 11l2.48-5.788A2 2 0 0 1 7.32 4h9.362a2 2 0 0 1 1.838 1.212L21 11M3 11h18M3 11v7m18-7v7m-3 0v.5a1.5 1.5 0 0 0 1.5 1.5v0a1.5 1.5 0 0 0 1.5-1.5V18m-3 0H6m12 0h3M6 18v.5A1.5 1.5 0 0 1 4.5 20v0A1.5 1.5 0 0 1 3 18.5V18m3 0H3m4-4h.001M17 14h.001"/>`
	cashInnerSVG                      = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M15 4c1.306 0 2.418.835 2.83 2H9a5 5 0 0 0-5 5v4.83A3.001 3.001 0 0 1 2 13V7a3 3 0 0 1 3-3h10zm4 4H9a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3zm-3 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0z" fill="currentColor"/></g>`
	cashLineInnerSVG                  = `<g fill="currentColor"><path d="M6 11a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3v-6zm3-1a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H9z"/><path d="M5 6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2v2H5a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v2h-2V7a1 1 0 0 0-1-1H5z"/><path d="M14 13a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0z"/></g>`
	centCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm12-6a1 1 0 1 0-2 0v1.1a5.002 5.002 0 0 0 0 9.8V18a1 1 0 1 0 2 0v-1.1c.98-.2 1.856-.685 2.536-1.364a1 1 0 1 0-1.415-1.415a3 3 0 1 1 0-4.243a1 1 0 1 0 1.415-1.414A4.993 4.993 0 0 0 13 7.1V6z" clip-rule="evenodd"/>`
	centCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M12 16a4 4 0 0 1 0-8m0 8v2m0-2a3.987 3.987 0 0 0 2.828-1.172M12 8V6m0 2c1.105 0 2.105.448 2.828 1.172"/></g>`
	chartBarInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M18 4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-2zm-8 7a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-8zm-7 6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-2z" fill="currentColor"/><path d="M18 5V3a2 2 0 0 0-2 2h2zm0 14V5h-2v14h2zm0 0h-2a2 2 0 0 0 2 2v-2zm2 0h-2v2h2v-2zm0 0v2a2 2 0 0 0 2-2h-2zm0-14v14h2V5h-2zm0 0h2a2 2 0 0 0-2-2v2zm-2 0h2V3h-2v2zm-7 4a2 2 0 0 0-2 2h2V9zm2 0h-2v2h2V9zm2 2a2 2 0 0 0-2-2v2h2zm0 8v-8h-2v8h2zm-2 2a2 2 0 0 0 2-2h-2v2zm-2 0h2v-2h-2v2zm-2-2a2 2 0 0 0 2 2v-2H9zm0-8v8h2v-8H9zm-5 4a2 2 0 0 0-2 2h2v-2zm2 0H4v2h2v-2zm2 2a2 2 0 0 0-2-2v2h2zm0 2v-2H6v2h2zm-2 2a2 2 0 0 0 2-2H6v2zm-2 0h2v-2H4v2zm-2-2a2 2 0 0 0 2 2v-2H2zm0-2v2h2v-2H2z" fill="currentColor"/></g>`
	chartBarLineInnerSVG              = `<g fill="currentColor"><path d="M19 3a3 3 0 0 0-3 3v12a3 3 0 1 0 6 0V6a3 3 0 0 0-3-3zm-1 3a1 1 0 1 1 2 0v12a1 1 0 1 1-2 0V6zm-9 6a3 3 0 1 1 6 0v6a3 3 0 1 1-6 0v-6zm3-1a1 1 0 0 0-1 1v6a1 1 0 1 0 2 0v-6a1 1 0 0 0-1-1zM2 18a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm3-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/></g>`
	chartPieInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 21a9 9 0 0 1-9-9c0-4.248 2.943-7.854 6.9-8.808c.58-.14 1.1.329 1.1.925V11a2 2 0 0 0 2 2h6.883c.596 0 1.064.52.925 1.1c-.954 3.957-4.56 6.9-8.808 6.9zm3-17v5h5c-.333-1.667-1.8-5-5-5z" fill="currentColor"/><path d="M9.9 3.192l.235.973l-.235-.973zM20.808 14.1l.972.234l-.972-.234zM15 9h-1a1 1 0 0 0 1 1V9zm0-5V3a1 1 0 0 0-1 1h1zm5 5v1a1 1 0 0 0 .98-1.196L20 9zM2 12c0 5.523 4.477 10 10 10v-2a8 8 0 0 1-8-8H2zm7.666-9.78C5.264 3.28 2 7.285 2 12h2c0-3.78 2.62-6.989 6.135-7.835L9.666 2.22zM12 4.117c0-1.15-1.037-2.21-2.334-1.897l.469 1.945a.146.146 0 0 1-.12-.028C10 4.123 10 4.112 10 4.117h2zM12 11V4.117h-2V11h2zm1 1a1 1 0 0 1-1-1h-2a3 3 0 0 0 3 3v-2zm6.883 0H13v2h6.883v-2zm1.897 2.334c.312-1.296-.748-2.334-1.897-2.334v2c.005 0-.006.001-.02-.016a.145.145 0 0 1-.027-.119l1.944.469zM12 22c4.715 0 8.72-3.264 9.78-7.666l-1.945-.469C18.99 17.38 15.78 20 12 20v2zm4-13V4h-2v5h2zm4-1h-5v2h5V8zm-5-3c1.203 0 2.105.611 2.785 1.488c.694.895 1.091 1.993 1.234 2.708l1.962-.392c-.19-.952-.694-2.354-1.616-3.542C18.429 4.055 16.997 3 15 3v2z" fill="currentColor"/></g>`
	chartPieLineInnerSVG              = `<g fill="currentColor"><path d="M4 12c0-3.732 2.554-6.907 6-7.802V11a3 3 0 0 0 3 3h6.802c-.895 3.446-4.07 6-7.802 6a8 8 0 0 1-8-8zm8-7.883c0-1.15-1.037-2.21-2.334-1.897C5.264 3.28 2 7.285 2 12c0 5.523 4.477 10 10 10c4.715 0 8.72-3.264 9.78-7.666c.312-1.296-.748-2.334-1.897-2.334H13a1 1 0 0 1-1-1V4.117zM15 3a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 .98-1.196c-.19-.952-.693-2.354-1.615-3.542C18.429 4.055 16.997 3 15 3zm1 5V5.155c.722.234 1.308.718 1.785 1.333c.37.476.654 1.01.862 1.512H16z"/></g>`
	chatInnerSVG                      = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-7.605-4.185L3 21l4.185-1.395A8.958 8.958 0 0 0 12 21z"/>`
	chatLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-7.605-4.185L3 21l4.185-1.395A8.958 8.958 0 0 0 12 21z"/>`
	chatSignalInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10a9.96 9.96 0 0 1-4.935-1.3l-3.749 1.249a1 1 0 0 1-1.265-1.265l1.25-3.749A9.959 9.959 0 0 1 2 12zm10.221-7.752a1 1 0 0 0-.442 1.95a8.013 8.013 0 0 1 6.023 6.023a1 1 0 1 0 1.95-.442a10.013 10.013 0 0 0-7.53-7.531zM9 9a1 1 0 0 1 1-1a6 6 0 0 1 6 6a1 1 0 1 1-2 0a4 4 0 0 0-4-4a1 1 0 0 1-1-1zm1 3a1 1 0 1 0 0 2a1 1 0 1 0 2 0a2 2 0 0 0-2-2z" clip-rule="evenodd"/>`
	chatSignalLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3.223a9.003 9.003 0 0 0-5.605 13.592L3 21l4.185-1.395A9.003 9.003 0 0 0 20.777 14m0-4A9.013 9.013 0 0 0 14 3.223M17 12a5 5 0 0 0-5-5m1 5a1 1 0 0 0-1-1"/>`
	chatStatusInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12a9.97 9.97 0 0 0 1.3 4.935l-1.249 3.749a1 1 0 0 0 1.265 1.265l3.749-1.25A9.96 9.96 0 0 0 12 22c5.523 0 10-4.477 10-10S17.523 2 12 2zm0 6c-.902 0-1.731.297-2.4.8a1 1 0 1 1-1.2-1.6a6 6 0 0 1 8.4 8.4a1 1 0 0 1-1.598-1.2A4 4 0 0 0 12 8zm-5 3a1 1 0 0 1 1 1a4 4 0 0 0 4 4a1 1 0 1 1 0 2a6 6 0 0 1-6-6a1 1 0 0 1 1-1zm5-1a2 2 0 1 0 0 4a2 2 0 0 0 0-4z" clip-rule="evenodd"/>`
	chatStatusLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 3.512a9.025 9.025 0 0 1 5.5 5.523M11 3.055a9.001 9.001 0 0 0-6.605 13.76L3 21l4.185-1.395A9.001 9.001 0 0 0 20.945 13"/><path d="M12 17a5 5 0 0 1-5-5m2-4a5 5 0 0 1 7 7"/><circle cx="12" cy="12" r="1"/></g>`
	chatTextInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10a9.96 9.96 0 0 1-4.935-1.3l-3.749 1.249a1 1 0 0 1-1.265-1.265l1.25-3.749A9.959 9.959 0 0 1 2 12zm6-5a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`
	chatTextLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9h8m-8 3h8m-8 3h3m10-3a9 9 0 0 1-13.815 7.605L3 21l1.395-4.185A9 9 0 1 1 21 12z"/>`
	chatTwoInnerSVG                   = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19c5.523 0 10-3.582 10-8s-4.477-8-10-8S2 6.582 2 11c0 2.157 1.067 4.114 2.801 5.553C4.271 18.65 3 20 2 21c3 0 4.527-.979 6.32-2.559c1.14.36 2.38.559 3.68.559z"/>`
	chatTwoLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19c5.523 0 10-3.582 10-8s-4.477-8-10-8S2 6.582 2 11c0 2.157 1.067 4.114 2.801 5.553C4.271 18.65 3 20 2 21c3 0 4.527-.979 6.32-2.559c1.14.36 2.38.559 3.68.559z"/>`
	chatTwoTextInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1 11c0-5.167 5.145-9 11-9s11 3.833 11 9s-5.145 9-11 9c-1.198 0-2.354-.156-3.437-.447c-.785.662-1.59 1.244-2.54 1.672C4.894 21.735 3.617 22 2 22a1 1 0 0 1-.707-1.707c.876-.876 1.843-1.914 2.368-3.416C2.029 15.327 1 13.28 1 11zm7-5a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`
	chatTwoTextLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8h8m-8 3h8m-8 3h3m11-3c0 4.418-4.477 8-10 8c-1.3 0-2.54-.198-3.68-.559C6.528 20.021 5 21 2 21c1-1 2.27-2.35 2.801-4.447C3.067 15.114 2 13.157 2 11c0-4.418 4.477-8 10-8s10 3.582 10 8z"/>`
	chatsInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13.584 7a5.001 5.001 0 1 0-8.809 4.675L4 14l2.325-.775c.214.136.44.256.675.359"/><path fill="currentColor" d="M15 20a5 5 0 1 1 4.225-2.325L20 20l-2.325-.775A4.976 4.976 0 0 1 15 20z"/></g>`
	chatsLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.584 7a5.001 5.001 0 1 0-8.809 4.675L4 14l2.325-.775c.214.136.44.256.675.359M15 20a5 5 0 1 1 4.225-2.325L20 20l-2.325-.775A4.976 4.976 0 0 1 15 20z"/>`
	chatsTwoInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 18.72C6.339 20.134 4.82 21 2 21c1-1 2.27-2.35 2.801-4.447C3.067 15.114 2 13.157 2 11c0-4.418 4.477-8 10-8c5.1 0 9.308 3.054 9.923 7"/><path fill="currentColor" d="M16 19.889c-3.314 0-6-1.99-6-4.445C10 12.99 12.686 11 16 11s6 1.99 6 4.444c0 1.199-.64 2.286-1.68 3.085c.317 1.165 1.08 1.915 1.68 2.471c-1.8 0-2.716-.544-3.792-1.422c-.684.2-1.428.31-2.208.31z"/></g>`
	chatsTwoLineInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 18.72C6.339 20.134 4.82 21 2 21c1-1 2.27-2.35 2.801-4.447C3.067 15.114 2 13.157 2 11c0-4.418 4.477-8 10-8c5.1 0 9.308 3.054 9.923 7"/><path d="M16 19.889c-3.314 0-6-1.99-6-4.445C10 12.99 12.686 11 16 11s6 1.99 6 4.444c0 1.199-.64 2.286-1.68 3.085c.317 1.165 1.08 1.915 1.68 2.471c-1.8 0-2.716-.544-3.792-1.422c-.684.2-1.428.31-2.208.31z"/></g>`
	checkInnerSVG                     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5zm14.707 6.707a1 1 0 0 0-1.414-1.414L9 15.586l-3.293-3.293a1 1 0 0 0-1.414 1.414l4 4a1 1 0 0 0 1.414 0l10-10z" fill="currentColor"/></g>`
	checkCircleInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm13.707-1.293a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4z" fill="currentColor"/></g>`
	checkCircleLineInnerSVG           = `<g fill="currentColor"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12z"/><path d="M15.707 9.293a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 1 1 1.414-1.414L11 12.586l3.293-3.293a1 1 0 0 1 1.414 0z"/></g>`
	checkLineInnerSVG                 = `<g fill="currentColor"><path d="M19.707 6.293a1 1 0 0 1 0 1.414l-10 10a1 1 0 0 1-1.414 0l-4-4a1 1 0 1 1 1.414-1.414L9 15.586l9.293-9.293a1 1 0 0 1 1.414 0z"/></g>`
	checkboxListInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h10m-10 7h10m-10 7h10"/><rect width="4" height="4" x="3" y="3" fill="currentColor" rx="1"/><rect width="4" height="4" x="3" y="10" fill="currentColor" rx="1"/><rect width="4" height="4" x="3" y="17" fill="currentColor" rx="1"/></g>`
	checkboxListDetailInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h10M11 9h5"/><rect width="4" height="4" x="3" y="5" fill="currentColor" rx="1"/><path d="M11 15h10m-10 4h5"/><rect width="4" height="4" x="3" y="15" fill="currentColor" rx="1"/></g>`
	checkboxListDetailLineInnerSVG    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h10M11 9h5"/><rect width="4" height="4" x="3" y="5" rx="1"/><path d="M11 15h10m-10 4h5"/><rect width="4" height="4" x="3" y="15" rx="1"/></g>`
	checkboxListLineInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h10m-10 7h10m-10 7h10"/><rect width="4" height="4" x="3" y="3" rx="1"/><rect width="4" height="4" x="3" y="10" rx="1"/><rect width="4" height="4" x="3" y="17" rx="1"/></g>`
	cheeseInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M10.783 3.024a1 1 0 0 1 .664.082l10 5A1 1 0 0 1 22 9v4a1 1 0 0 1-1 1c-.173 0-.456.06-.666.211c-.159.115-.334.315-.334.789c0 .474.175.674.334.789c.21.15.493.211.666.211a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H4.5a1 1 0 0 1-.894-.553c-1.34-2.679-2.02-6.427-1.136-9.824c.903-3.475 3.434-6.515 8.313-7.6zM8 17a1 1 0 1 1 0-2h.001a1 1 0 1 1 0 2H8zm4-4a1 1 0 0 0 1 1h.001a1 1 0 1 0 0-2H13a1 1 0 0 0-1 1zm3 4a1 1 0 1 1 0-2h.001a1 1 0 1 1 0 2H15z" clip-rule="evenodd"/>`
	cheeseLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 9v4c-.667 0-2 .4-2 2s1.333 2 2 2v3H4.5C2 15 2 6 11 4l10 5zm0 0H5m3 7h.001M13 13h.001M15 16h.001"/>`
	chevronDoubleDownInnerSVG         = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 5a1 1 0 0 0-.707 1.707l4 4a1 1 0 0 0 1.414 0l4-4A1 1 0 0 0 16 5H8zm0 8a1 1 0 0 0-.707 1.707l4 4a1 1 0 0 0 1.414 0l4-4A1 1 0 0 0 16 13H8z" fill="currentColor"/></g>`
	chevronDoubleDownLineInnerSVG     = `<g fill="currentColor"><path d="M5.293 12.293a1 1 0 0 1 1.414 0L12 17.586l5.293-5.293a1 1 0 0 1 1.414 1.414l-6 6a1 1 0 0 1-1.414 0l-6-6a1 1 0 0 1 0-1.414z"/><path d="M5.293 4.293a1 1 0 0 1 1.414 0L12 9.586l5.293-5.293a1 1 0 1 1 1.414 1.414l-6 6a1 1 0 0 1-1.414 0l-6-6a1 1 0 0 1 0-1.414z"/></g>`
	chevronDoubleLeftInnerSVG         = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11 8a1 1 0 0 0-1.707-.707l-4 4a1 1 0 0 0 0 1.414l4 4A1 1 0 0 0 11 16V8zm8 0a1 1 0 0 0-1.707-.707l-4 4a1 1 0 0 0 0 1.414l4 4A1 1 0 0 0 19 16V8z" fill="currentColor"/></g>`
	chevronDoubleLeftLineInnerSVG     = `<g fill="currentColor"><path d="M11.707 5.293a1 1 0 0 1 0 1.414L6.414 12l5.293 5.293a1 1 0 0 1-1.414 1.414l-6-6a1 1 0 0 1 0-1.414l6-6a1 1 0 0 1 1.414 0z"/><path d="M19.707 5.293a1 1 0 0 1 0 1.414L14.414 12l5.293 5.293a1 1 0 0 1-1.414 1.414l-6-6a1 1 0 0 1 0-1.414l6-6a1 1 0 0 1 1.414 0z"/></g>`
	chevronDoubleRightInnerSVG        = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 16a1 1 0 0 0 1.707.707l4-4a1 1 0 0 0 0-1.414l-4-4A1 1 0 0 0 5 8v8zm8 0a1 1 0 0 0 1.707.707l4-4a1 1 0 0 0 0-1.414l-4-4A1 1 0 0 0 13 8v8z" fill="currentColor"/></g>`
	chevronDoubleRightLineInnerSVG    = `<g fill="currentColor"><path d="M12.293 5.293a1 1 0 0 0 0 1.414L17.586 12l-5.293 5.293a1 1 0 0 0 1.414 1.414l6-6a1 1 0 0 0 0-1.414l-6-6a1 1 0 0 0-1.414 0z"/><path d="M4.293 5.293a1 1 0 0 0 0 1.414L9.586 12l-5.293 5.293a1 1 0 1 0 1.414 1.414l6-6a1 1 0 0 0 0-1.414l-6-6a1 1 0 0 0-1.414 0z"/></g>`
	chevronDoubleUpInnerSVG           = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 11a1 1 0 0 1-.707-1.707l4-4a1 1 0 0 1 1.414 0l4 4A1 1 0 0 1 16 11H8zm0 8a1 1 0 0 1-.707-1.707l4-4a1 1 0 0 1 1.414 0l4 4A1 1 0 0 1 16 19H8z" fill="currentColor"/></g>`
	chevronDoubleUpLineInnerSVG       = `<g fill="currentColor"><path d="M5.293 11.707a1 1 0 0 0 1.414 0L12 6.414l5.293 5.293a1 1 0 0 0 1.414-1.414l-6-6a1 1 0 0 0-1.414 0l-6 6a1 1 0 0 0 0 1.414z"/><path d="M5.293 19.707a1 1 0 0 0 1.414 0L12 14.414l5.293 5.293a1 1 0 0 0 1.414-1.414l-6-6a1 1 0 0 0-1.414 0l-6 6a1 1 0 0 0 0 1.414z"/></g>`
	chevronDownInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 10l-5 5l-5-5"/>`
	chevronDownCircleInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zM9.707 10.293a1 1 0 1 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L12 12.586l-2.293-2.293z" clip-rule="evenodd"/>`
	chevronDownCircleLineInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m9 11l3 3l3-3"/></g>`
	chevronDownLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 10l-5 5l-5-5"/>`
	chevronLeftInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 7l-5 5l5 5"/>`
	chevronLeftCircleInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm1.707-11.293a1 1 0 0 0-1.414-1.414l-3 3a1 1 0 0 0 0 1.414l3 3a1 1 0 0 0 1.414-1.414L11.414 12l2.293-2.293z" clip-rule="evenodd"/>`
	chevronLeftCircleLineInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13 9l-3 3l3 3"/><circle cx="12" cy="12" r="9"/></g>`
	chevronLeftLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 7l-5 5l5 5"/>`
	chevronRightInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 7l5 5l-5 5"/>`
	chevronRightCircleInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zM10.293 9.707a1 1 0 1 1 1.414-1.414l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414-1.414L12.586 12l-2.293-2.293z" clip-rule="evenodd"/>`
	chevronRightCircleLineInnerSVG    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m11 9l3 3l-3 3"/></g>`
	chevronRightLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 7l5 5l-5 5"/>`
	chevronUpInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 14l-5-5l-5 5"/>`
	chevronUpCircleInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm-2.293-7.293a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1-1.414 1.414L12 11.414l-2.293 2.293z" clip-rule="evenodd"/>`
	chevronUpCircleLineInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m9 13l3-3l3 3"/></g>`
	chevronUpLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 14l-5-5l-5 5"/>`
	chipInnerSVG                      = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10 3a1 1 0 0 0-2 0v1H7a3 3 0 0 0-3 3v1H3a1 1 0 0 0 0 2h1v4H3a1 1 0 1 0 0 2h1v1a3 3 0 0 0 3 3h1v1a1 1 0 1 0 2 0v-1h4v1a1 1 0 1 0 2 0v-1h1a3 3 0 0 0 3-3v-1h1a1 1 0 1 0 0-2h-1v-4h1a1 1 0 1 0 0-2h-1V7a3 3 0 0 0-3-3h-1V3a1 1 0 1 0-2 0v1h-4V3zm-.25 6a.75.75 0 0 0-.75.75v4.5c0 .414.336.75.75.75h4.5a.75.75 0 0 0 .75-.75v-4.5a.75.75 0 0 0-.75-.75h-4.5zM12 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2z" fill="currentColor"/></g>`
	chipLineInnerSVG                  = `<g fill="currentColor"><path d="M9 2a1 1 0 0 1 1 1v1h4V3a1 1 0 1 1 2 0v1h1a3 3 0 0 1 3 3v1h1a1 1 0 1 1 0 2h-1v4h1a1 1 0 1 1 0 2h-1v1a3 3 0 0 1-3 3h-1v1a1 1 0 1 1-2 0v-1h-4v1a1 1 0 1 1-2 0v-1H7a3 3 0 0 1-3-3v-1H3a1 1 0 1 1 0-2h1v-4H3a1 1 0 0 1 0-2h1V7a3 3 0 0 1 3-3h1V3a1 1 0 0 1 1-1zm8 16a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10zM8 9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V9zm2 1v4h4v-4h-4z"/></g>`
	chromecastInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 20h.01M7 20a4 4 0 0 0-4-4m8 4a8 8 0 0 0-8-8"/><path fill="currentColor" d="M19 4H5a2 2 0 0 0-2 2v2.5c4 .167 12 2.7 12 11.5h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/></g>`
	chromecastLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 20h.01M7 20a4 4 0 0 0-4-4m8 4a8 8 0 0 0-8-8"/><path d="M3 8.5V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-4"/></g>`
	churchInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M12 3a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2h-1v1.5l7.6 5.7A1 1 0 0 1 20 16h-1v3h2a1 1 0 1 1 0 2h-7v-3a2 2 0 1 0-4 0v3H3a1 1 0 1 1 0-2h2v-3H4a1 1 0 0 1-.6-1.8L11 8.5V7h-1a1 1 0 0 1 0-2h1V4a1 1 0 0 1 1-1z" clip-rule="evenodd"/>`
	churchLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 9l-8 6h2v5m6-11l8 6h-2v5M12 9V4m6 16h3m-3 0h-4M3 20h3m0 0h4m0-14h4m-4 14v-3a2 2 0 0 1 2-2v0a2 2 0 0 1 2 2v3m-4 0h4"/>`
	clipboardInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M7.17 4A3.001 3.001 0 0 1 10 2h4c1.306 0 2.418.835 2.83 2H18a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h1.17zM10 4a1 1 0 0 0 0 2h4a1 1 0 1 0 0-2h-4z" clip-rule="evenodd"/>`
	clipboardCheckInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M10 2a3.001 3.001 0 0 0-2.83 2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3h-1.17A3.001 3.001 0 0 0 14 2h-4zM9 5a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1zm6.78 6.625a1 1 0 1 0-1.56-1.25l-3.303 4.128l-1.21-1.21a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.488-.082l4-5z" clip-rule="evenodd"/>`
	clipboardCheckLineInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 5H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2M8 5v0a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v0M8 5v0a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v0"/><path d="m9 14l2 2l4-5"/></g>`
	clipboardCopyInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10 4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2h-4zm0-2a3.001 3.001 0 0 0-2.83 2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-4h-8.586l1.293 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 1 1 1.414 1.414L12.414 13H21V7a3 3 0 0 0-3-3h-1.17A3.001 3.001 0 0 0 14 2h-4z" fill="currentColor"/></g>`
	clipboardCopyLineInnerSVG         = `<g fill="currentColor"><path d="M10 4a1 1 0 0 0 0 2h4a1 1 0 1 0 0-2h-4zM7.17 4A3.001 3.001 0 0 1 10 2h4c1.306 0 2.418.835 2.83 2H18a3 3 0 0 1 3 3v3a1 1 0 1 1-2 0V7a1 1 0 0 0-1-1h-1.17A3.001 3.001 0 0 1 14 8h-4a3.001 3.001 0 0 1-2.83-2H6a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1a1 1 0 1 1 2 0v1a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h1.17zm7.537 6.293a1 1 0 0 1 0 1.414L13.414 13H20a1 1 0 1 1 0 2h-6.586l1.293 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 0 1 1.414 0z"/></g>`
	clipboardLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2M8 5v0a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v0M8 5v0a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v0"/>`
	clipboardListInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10 4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2h-4zm0-2a3.001 3.001 0 0 0-2.83 2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3h-1.17A3.001 3.001 0 0 0 14 2h-4zm-2 8a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H8z" fill="currentColor"/></g>`
	clipboardListLineInnerSVG         = `<g fill="currentColor"><path d="M10 4a1 1 0 0 0 0 2h4a1 1 0 1 0 0-2h-4zM7.17 4A3.001 3.001 0 0 1 10 2h4c1.306 0 2.418.835 2.83 2H18a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h1.17zm0 2H6a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-1.17A3.001 3.001 0 0 1 14 8h-4a3.001 3.001 0 0 1-2.83-2zM7 11a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1z"/></g>`
	clipboardMinusInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M10 2a3.001 3.001 0 0 0-2.83 2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3h-1.17A3.001 3.001 0 0 0 14 2h-4zM9 5a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1zm6 8a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h6z" clip-rule="evenodd"/>`
	clipboardMinusLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2M8 5v0a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v0M8 5v0a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v0m-7 9h6"/>`
	clipboardPlusInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M10 2a3.001 3.001 0 0 0-2.83 2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3h-1.17A3.001 3.001 0 0 0 14 2h-4zM9 5a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1zm2 6a1 1 0 1 1 2 0v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9a1 1 0 1 1 0-2h2v-2z" clip-rule="evenodd"/>`
	clipboardPlusLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2M8 5v0a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v0M8 5v0a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v0m-4 6v3m0 3v-3m0 0H9m3 0h3"/>`
	clockInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm11-5a1 1 0 1 0-2 0v3.764a3 3 0 0 0 1.658 2.683l2.895 1.447a1 1 0 1 0 .894-1.788l-2.894-1.448a1 1 0 0 1-.553-.894V7z" clip-rule="evenodd"/>`
	clockLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 7v3.764a2 2 0 0 0 1.106 1.789L16 14"/></g>`
	clockPlusInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 16v3m0 3v-3m0 0h-3m3 0h3"/><path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c.587 0 1.161-.05 1.72-.147A5.973 5.973 0 0 1 13 19c0-1.746.746-3.318 1.936-4.414l-2.278-1.139A3 3 0 0 1 11 10.764V7a1 1 0 1 1 2 0v3.764a1 1 0 0 0 .553.894l2.894 1.448c.149.074.271.18.362.306A5.986 5.986 0 0 1 19 13c1.033 0 2.004.26 2.853.72A10.15 10.15 0 0 0 22 12c0-5.523-4.477-10-10-10z" clip-rule="evenodd"/></g>`
	clockPlusLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 1 0-9 9m0-14v3.764a2 2 0 0 0 1.106 1.789L16 14m3 2v3m0 3v-3m0 0h-3m3 0h3"/>`
	closeInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12L7 7m5 5l5 5m-5-5l5-5m-5 5l-5 5"/>`
	closeCircleInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm8.207-3.207a1 1 0 0 0-1.414 1.414L10.586 12l-1.793 1.793a1 1 0 1 0 1.414 1.414L12 13.414l1.793 1.793a1 1 0 0 0 1.414-1.414L13.414 12l1.793-1.793a1 1 0 0 0-1.414-1.414L12 10.586l-1.793-1.793z" fill="currentColor"/></g>`
	closeCircleLineInnerSVG           = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm-1.793 6.793a1 1 0 0 0-1.414 1.414L10.586 12l-1.793 1.793a1 1 0 1 0 1.414 1.414L12 13.414l1.793 1.793a1 1 0 0 0 1.414-1.414L13.414 12l1.793-1.793a1 1 0 0 0-1.414-1.414L12 10.586l-1.793-1.793z"/></g>`
	closeLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12L7 7m5 5l5 5m-5-5l5-5m-5 5l-5 5"/>`
	cloudInnerSVG                     = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 13.5C2 17.9 5.667 19 7.5 19h10c1.5 0 4.5-.9 4.5-4.5S19 10 17.5 10c0-1.5-1.5-5-5-5c-2.8 0-4.5 2-5 3C5.667 8 2 9.1 2 13.5z"/>`
	cloudDownloadInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6.913 7.029C7.751 5.772 9.626 4 12.5 4c2.13 0 3.65 1.08 4.607 2.33a7.133 7.133 0 0 1 1.285 2.745c.785.127 1.695.43 2.505 1.014C22.092 10.948 23 12.373 23 14.5c0 2.126-.908 3.551-2.103 4.412C19.753 19.735 18.41 20 17.5 20h-10c-1.077 0-2.67-.315-4.022-1.288C2.075 17.701 1 16.026 1 13.5s1.075-4.201 2.478-5.212c1.124-.809 2.413-1.163 3.435-1.26zM13 10a1 1 0 1 0-2 0v4.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L13 14.586V10z" fill="currentColor"/></g>`
	cloudDownloadLineInnerSVG         = `<g fill="currentColor"><path d="M12.5 6c-2.294 0-3.71 1.655-4.106 2.447A1 1 0 0 1 7.5 9c-.757 0-1.914.235-2.853.912C3.758 10.552 3 11.626 3 13.5c0 1.458.459 2.415 1.05 3.06c.607.663 1.418 1.062 2.204 1.269a1 1 0 0 1-.508 1.934c-1.049-.276-2.238-.833-3.171-1.852C1.624 16.873 1 15.423 1 13.5c0-2.526 1.075-4.201 2.478-5.212c1.124-.809 2.413-1.163 3.435-1.26C7.751 5.773 9.626 4 12.5 4c2.13 0 3.65 1.08 4.607 2.33a7.133 7.133 0 0 1 1.285 2.745c.785.127 1.695.43 2.505 1.014C22.092 10.948 23 12.373 23 14.5c0 1.516-.462 2.697-1.196 3.571c-.72.86-1.65 1.362-2.498 1.634a1 1 0 1 1-.612-1.904c.586-.188 1.157-.513 1.578-1.015c.408-.486.728-1.202.728-2.286c0-1.474-.592-2.299-1.272-2.789c-.73-.526-1.638-.711-2.228-.711a1 1 0 0 1-1-1c0-.502-.284-1.543-.982-2.455C14.85 6.67 13.87 6 12.5 6zm-.5 4a1 1 0 0 1 1 1v5.586l1.293-1.293a1 1 0 0 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L11 16.586V11a1 1 0 0 1 1-1z"/></g>`
	cloudLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 13.5C2 17.9 5.667 19 7.5 19h10c1.5 0 4.5-.9 4.5-4.5S19 10 17.5 10c0-1.5-1.5-5-5-5c-2.8 0-4.5 2-5 3C5.667 8 2 9.1 2 13.5z"/>`
	cloudUploadInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6.913 7.029C7.751 5.772 9.626 4 12.5 4c2.13 0 3.65 1.08 4.607 2.33a7.133 7.133 0 0 1 1.285 2.745c.785.127 1.695.43 2.505 1.014C22.092 10.948 23 12.373 23 14.5c0 2.126-.908 3.551-2.103 4.412C19.753 19.735 18.41 20 17.5 20H13v-6.586l1.293 1.293a1 1 0 0 0 1.414-1.414l-3-3a1 1 0 0 0-1.414 0l-3 3a1 1 0 1 0 1.414 1.414L11 13.414V20H7.5c-1.077 0-2.67-.315-4.022-1.288C2.075 17.701 1 16.026 1 13.5s1.075-4.201 2.478-5.212c1.124-.809 2.413-1.163 3.435-1.26z" fill="currentColor"/></g>`
	cloudUploadLineInnerSVG           = `<g fill="currentColor"><path d="M12.5 6c-2.294 0-3.71 1.655-4.106 2.447A1 1 0 0 1 7.5 9c-.757 0-1.914.235-2.853.912C3.758 10.552 3 11.626 3 13.5s.758 2.949 1.647 3.588c.94.677 2.096.912 2.853.912a1 1 0 1 1 0 2c-1.077 0-2.67-.315-4.022-1.288C2.075 17.701 1 16.026 1 13.5s1.075-4.201 2.478-5.212c1.124-.809 2.413-1.163 3.435-1.26C7.751 5.773 9.626 4 12.5 4c2.13 0 3.65 1.08 4.607 2.33a7.133 7.133 0 0 1 1.285 2.745c.785.127 1.695.43 2.505 1.014C22.092 10.948 23 12.373 23 14.5c0 2.126-.908 3.551-2.103 4.412C19.753 19.735 18.41 20 17.5 20a1 1 0 1 1 0-2c.59 0 1.497-.185 2.228-.712c.68-.49 1.272-1.314 1.272-2.788c0-1.474-.592-2.299-1.272-2.789c-.73-.526-1.638-.711-2.228-.711a1 1 0 0 1-1-1c0-.502-.284-1.543-.982-2.455C14.85 6.67 13.87 6 12.5 6zm-1.207 4.293a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1-1.414 1.414L13 13.414V19a1 1 0 1 1-2 0v-5.586l-1.293 1.293a1 1 0 0 1-1.414-1.414l3-3z"/></g>`
	codeInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 7l-5 5l5 5m8 0l5-5l-5-5"/>`
	codeBlockInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm5 1a1 1 0 0 0 0 2h5a1 1 0 1 0 0-2H7zm8 0a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2zm-8 4a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H7zm4 0a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6zm-4 4a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2H7zm8 0a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2z" clip-rule="evenodd"/>`
	codeBlockLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h5m3 0h2m0 4h-6m-3 0H7m0 4h5m3 0h2M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8z"/>`
	codeLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 7l-5 5l5 5m8 0l5-5l-5-5"/>`
	cogInnerSVG                       = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11 2a2 2 0 0 0-2 2v.757l-.707.707L9 4.757l-.536-.535a2 2 0 0 0-2.828 0L4.222 5.636a2 2 0 0 0 0 2.828L4.757 9H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h.757l-.535.536a2 2 0 0 0 0 2.828l1.414 1.414a2 2 0 0 0 2.828 0l-.707-.707l.707.707l.536-.535V20a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-.757l.536.535a2 2 0 0 0 2.828 0l1.414-1.414a2 2 0 0 0 0-2.829L19.243 15H20a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-.757l.535-.536a2 2 0 0 0 0-2.828l-1.414-1.414a2 2 0 0 0-2.828 0L15 4.757V4a2 2 0 0 0-2-2h-2zm5 10a4 4 0 1 1-8 0a4 4 0 0 1 8 0zm-4 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4z" fill="currentColor"/></g>`
	cogLineInnerSVG                   = `<g fill="currentColor"><path d="M11 2a2 2 0 0 0-2 2v.757l-.707.707L9 4.757l-.536-.535a2 2 0 0 0-2.828 0L4.222 5.636a2 2 0 0 0 0 2.828L4.757 9H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h.757l-.535.536a2 2 0 0 0 0 2.828l1.414 1.414a2 2 0 0 0 2.828 0L9 19.243V20a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-.757l.536.535a2 2 0 0 0 2.828 0l1.414-1.414a2 2 0 0 0 0-2.829L19.243 15H20a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-.757l.535-.536a2 2 0 0 0 0-2.828l-1.414-1.414a2 2 0 0 0-2.828 0L15 4.757V4a2 2 0 0 0-2-2h-2zm0 2h2v.757c0 1.782 2.154 2.675 3.414 1.415l.536-.536l1.414 1.414l-.536.536C16.568 8.846 17.461 11 19.243 11H20v2h-.757c-1.782 0-2.674 2.154-1.415 3.414l.536.536l-1.414 1.414l-.536-.536c-1.26-1.26-3.414-.367-3.414 1.415V20h-2v-.757c0-1.782-2.154-2.674-3.414-1.415l-.536.536l-1.414-1.414l.536-.536C7.432 15.154 6.539 13 4.757 13H4v-2h.757c1.782 0 2.675-2.154 1.415-3.414l-.536-.536L7.05 5.636l.536.536C8.846 7.432 11 6.539 11 4.757V4zm-1 8a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm2-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8z"/></g>`
	coinsInnerSVG                     = `<g fill="none"><path fill="currentColor" d="M21 8c0 1.02-1.186 1.92-3 2.462c-1.134.34-2.513.538-4 .538s-2.866-.199-4-.538C8.187 9.92 7 9.02 7 8c0-1.657 3.134-3 7-3s7 1.343 7 3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 8c0-1.657-3.134-3-7-3S7 6.343 7 8m14 0v4c0 1.02-1.186 1.92-3 2.462c-1.134.34-2.513.538-4 .538s-2.866-.199-4-.538C8.187 13.92 7 13.02 7 12V8m14 0c0 1.02-1.186 1.92-3 2.462c-1.134.34-2.513.538-4 .538s-2.866-.199-4-.538C8.187 9.92 7 9.02 7 8"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12v4c0 1.02 1.187 1.92 3 2.462c1.134.34 2.513.538 4 .538s2.866-.199 4-.538c1.813-.542 3-1.442 3-2.462v-1M3 12c0-1.197 1.635-2.23 4-2.711M3 12c0 1.02 1.187 1.92 3 2.462c1.134.34 2.513.538 4 .538c.695 0 1.366-.043 2-.124"/></g>`
	coinsLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 8c0-1.657-3.134-3-7-3S7 6.343 7 8m14 0v4c0 1.02-1.186 1.92-3 2.462c-1.134.34-2.513.538-4 .538s-2.866-.199-4-.538C8.187 13.92 7 13.02 7 12V8m14 0c0 1.02-1.186 1.92-3 2.462c-1.134.34-2.513.538-4 .538s-2.866-.199-4-.538C8.187 9.92 7 9.02 7 8"/><path d="M3 12v4c0 1.02 1.187 1.92 3 2.462c1.134.34 2.513.538 4 .538s2.866-.199 4-.538c1.813-.542 3-1.442 3-2.462v-1M3 12c0-1.197 1.635-2.23 4-2.711M3 12c0 1.02 1.187 1.92 3 2.462c1.134.34 2.513.538 4 .538c.695 0 1.366-.043 2-.124"/></g>`
	collectionInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 21a2 2 0 0 1-2-2h12a2 2 0 0 1-2 2H8zm-4-5a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2H4zm0-1a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4z" fill="currentColor"/></g>`
	collectionLineInnerSVG            = `<g fill="currentColor"><path d="M6 18a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3v-.17A3.001 3.001 0 0 0 20 15v-.17A3.001 3.001 0 0 0 22 12V6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v6c0 1.306.835 2.418 2 2.83V15c0 1.306.835 2.418 2 2.83V18zm2 0h8a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1zm-2-3h12a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1zm-1-2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H5z"/></g>`
	colorSwatchInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M19 22a3 3 0 0 0 3-3v-4a3.001 3.001 0 0 0-2.361-2.932L13 18.708v.792c0 .58-.17 1.439-.694 2.167a3.212 3.212 0 0 1-.275.333H19zm-6-4.707l5.738-5.738a3.001 3.001 0 0 0-.445-3.676L16.12 5.707A3 3 0 0 0 13 5v12.294zM9 2a3 3 0 0 1 3 3v14.5c0 .42-.13 1.061-.506 1.583c-.357.496-.957.917-1.994.917H5a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h4zM7 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2z" fill="currentColor"/></g>`
	colorSwatchLineInnerSVG           = `<g fill="currentColor"><path d="M2 5a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v16a1 1 0 0 1-1 1H5a3 3 0 0 1-3-3V5zm3-1a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h5V5a1 1 0 0 0-1-1H5z"/><path d="M14.707 7.121a1 1 0 0 0-1.414 0l-1.586 1.586a1 1 0 1 1-1.414-1.414l1.586-1.586a3 3 0 0 1 4.242 0l2.172 2.172a3 3 0 0 1 .115 4.12H19a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3h-8a1 1 0 1 1 0-2h8a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2.586l-4.707 4.708a1 1 0 1 1-1.414-1.414l6.586-6.586a1 1 0 0 0 0-1.414L14.707 7.12zM7 15a1 1 0 0 1 1 1v.001a1 1 0 1 1-2 0V16a1 1 0 0 1 1-1z"/></g>`
	cometInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9.879 5.636l.707-.707a6 6 0 0 1 8.485 8.485l-2.121 2.122M7.757 7.757L5.636 9.88M3.515 12l-.707.707m12.02 4.95l-1.414 1.414m-.707-4.95l-.707.707M9.879 16.95l-2.122 2.12m2.122-7.778l-6.364 6.364"/><circle cx="14.828" cy="9.172" r="2" fill="currentColor" transform="rotate(45 14.828 9.172)"/></g>`
	cometLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9.879 5.636l.707-.707a6 6 0 0 1 8.485 8.485l-2.121 2.122M7.757 7.757L5.636 9.88M3.515 12l-.707.707m12.02 4.95l-1.414 1.414m-.707-4.95l-.707.707M9.879 16.95l-2.122 2.12m2.122-7.778l-6.364 6.364"/><circle cx="14.828" cy="9.172" r="2" transform="rotate(45 14.828 9.172)"/></g>`
	commentInnerSVG                   = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4H5a2 2 0 0 0-2 2v15l3.467-2.6a2 2 0 0 1 1.2-.4H19a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/>`
	commentLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4H5a2 2 0 0 0-2 2v15l3.467-2.6a2 2 0 0 1 1.2-.4H19a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/>`
	commentTextInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H7.667a1 1 0 0 0-.6.2L3.6 21.8A1 1 0 0 1 2 21V6zm5 0a1 1 0 0 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	commentTextLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h10M7 11h10M7 14h4m-8 4V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.667a2 2 0 0 0-1.2.4L3 21v-3z"/>`
	commentTwoInnerSVG                = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4H5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.188c1 0 1.812.811 1.812 1.812c0 .808.976 1.212 1.547.641l1.867-1.867A2 2 0 0 1 14.828 18H19a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/>`
	commentTwoLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4H5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.188c1 0 1.812.811 1.812 1.812c0 .808.976 1.212 1.547.641l1.867-1.867A2 2 0 0 1 14.828 18H19a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/>`
	commentTwoTextInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3h-4.172a1 1 0 0 0-.707.293l-1.867 1.867C11.054 22.361 9 21.51 9 19.812A.812.812 0 0 0 8.188 19H5a3 3 0 0 1-3-3V6zm5 0a1 1 0 0 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	commentTwoTextLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h10M7 11h10M7 14h4m3.828 4H19a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.188c1 0 1.812.811 1.812 1.812v0c0 .808.976 1.212 1.547.641l1.867-1.867A2 2 0 0 1 14.828 18z"/>`
	commentsInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 12h-.394a2 2 0 0 0-1.11.336L3 14V5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v2"/><path fill="currentColor" d="M12 10h7a2 2 0 0 1 2 2v9l-2.496-1.664a2 2 0 0 0-1.11-.336H12a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2z"/></g>`
	commentsLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12h-.394a2 2 0 0 0-1.11.336L3 14V5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v2m-2 3h7a2 2 0 0 1 2 2v9l-2.496-1.664a2 2 0 0 0-1.11-.336H12a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2z"/>`
	commentsTwoInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 14v-1a2 2 0 0 0-2-2v0a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v3"/><path fill="currentColor" d="M12 11h7a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-.64A1.36 1.36 0 0 0 17 20.36a.68.68 0 0 1-1.16.48l-1.254-1.254A2 2 0 0 0 13.172 19H12a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2z"/></g>`
	commentsTwoLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 14v-1a2 2 0 0 0-2-2v0a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v3m-2 3h7a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-.64A1.36 1.36 0 0 0 17 20.36a.68.68 0 0 1-1.16.48l-1.254-1.254A2 2 0 0 0 13.172 19H12a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2z"/>`
	communityInnerSVG                 = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 6a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3h-3.51a3.82 3.82 0 0 0 .51-1.911v-5.438a3.867 3.867 0 0 0-1.172-2.766l-4-3.911C9.52 5.694 7.53 5.487 6 6.351V6zm11 1a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H18a1 1 0 0 1-1-1zm-3-1a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H14zm3 6a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H18a1 1 0 0 1-1-1zm1 4a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H18z"/><path d="M5.879 8.707a3 3 0 0 1 4.242 0l3 3A3 3 0 0 1 14 13.828V18a3 3 0 0 1-3 3H9v-3a1 1 0 1 0-2 0v3H5a3 3 0 0 1-3-3v-4.172a3 3 0 0 1 .879-2.12l3-3z"/></g>`
	communityLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 9V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-6m0 0v-6.172a2 2 0 0 0-.586-1.414l-3-3a2 2 0 0 0-2.828 0l-3 3A2 2 0 0 0 3 13.828V18a2 2 0 0 0 2 2h3m5 0H8m0-4v4m9.001-12H17m-3.999 0H13m4.001 4H17m.001 4H17"/>`
	compassInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm5.191 6.074a1 1 0 0 0-1.265-1.265L9.562 8.93a1 1 0 0 0-.632.632l-2.121 6.364a1 1 0 0 0 1.265 1.265l6.364-2.121a1 1 0 0 0 .632-.633l2.121-6.363zM12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2z" fill="currentColor"/></g>`
	compassLineInnerSVG               = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm5.191 6.074a1 1 0 0 0-1.265-1.265L9.562 8.93a1 1 0 0 0-.632.632l-2.121 6.364a1 1 0 0 0 1.265 1.265l6.364-2.121a1 1 0 0 0 .632-.632l2.121-6.364zM9.34 14.662l1.33-3.993l3.992-1.33l-1.33 3.992l-3.992 1.33z"/></g>`
	compassTwoInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm14.243-4.243L9.879 9.88l-2.122 6.364l6.364-2.122l2.122-6.364z" clip-rule="evenodd"/>`
	compassTwoLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m9.879 9.879l6.364-2.122l-2.122 6.364l-6.364 2.122L9.88 9.879z"/></g>`
	contactInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H5zm7 7c0 .662-.215 1.275-.578 1.77c.62.465 1.073 1.088 1.444 1.73a1 1 0 1 1-1.732 1c-.313-.541-.61-.908-.93-1.142C9.909 14.141 9.54 14 9 14c-.54 0-.908.14-1.205.358c-.32.234-.616.6-.93 1.143a1 1 0 0 1-1.73-1.002c.37-.641.824-1.264 1.443-1.728A3 3 0 1 1 12 11zm3-3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3zm0 3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3zm0 3a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2h-1zm-6-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2z" fill="currentColor"/></g>`
	contactLineInnerSVG               = `<g fill="currentColor"><path d="M5 4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H5zM4 7a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7zm11 1a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3zm0 3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3zm0 3a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2h-1z"/><path d="M9 10a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0z"/><path d="M9 14c-.99 0-1.707.65-2.106 1.447a1 1 0 1 1-1.788-.894C5.707 13.349 6.99 12 9 12c2.01 0 3.293 1.35 3.894 2.553a1 1 0 1 1-1.788.894C10.707 14.651 9.99 14 9 14z"/></g>`
	cookieInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M6.85 4.787C8.731 3.44 10.958 3 12.456 3h1.005a1 1 0 0 1 1 1v3h2.014a1 1 0 0 1 1 1v2h3.019a1 1 0 0 1 1 1v1c0 3.503-1.395 5.808-3.297 7.206c-1.85 1.36-4.087 1.794-5.74 1.794c-1.654 0-3.89-.434-5.742-1.794C4.814 17.808 3.419 15.503 3.419 12c0-3.528 1.5-5.828 3.431-7.213zM9.442 7a1 1 0 1 0 0 2h.001a1 1 0 0 0 0-2zm-3.01 5a1 1 0 0 1 1-1h.002a1 1 0 1 1 0 2h-.001a1 1 0 0 1-1-1zm6.024 0a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2zm-3.01 4a1 1 0 0 1 1-1h.002a1 1 0 1 1 0 2h-.001a1 1 0 0 1-1-1zm7.029-2a1 1 0 1 0 0 2a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`
	cookieLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.419 12c0-6.4 5.358-8 8.037-8h1.005v4h3.014v3h4.019v1c0 6.4-5.024 8-8.038 8c-3.014 0-8.037-1.6-8.037-8zm5.023-4h.001m-2.01 4h0m5.023 1h.001m-2.01 3h0m6.028-1h0"/>`
	covidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M10 1a1 1 0 0 0 0 2h1v1.062A7.96 7.96 0 0 0 7.094 5.68L5.914 4.5l.793-.793a1 1 0 0 0-1.414-1.414l-3 3a1 1 0 0 0 1.414 1.414l.793-.793l1.18 1.18A7.96 7.96 0 0 0 4.062 11H3v-1a1 1 0 0 0-2 0v4a1 1 0 1 0 2 0v-1h1.062c.182 1.46.759 2.8 1.618 3.906l-1.18 1.18l-.793-.793a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414-1.414l-.793-.793l1.18-1.18A7.965 7.965 0 0 0 11 19.938V21h-1a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-1v-1.062a7.965 7.965 0 0 0 3.906-1.618l1.18 1.18l-.793.793a1 1 0 0 0 1.414 1.414l3-3a1 1 0 0 0-1.414-1.414l-.793.793l-1.18-1.18A7.965 7.965 0 0 0 19.938 13H21v1a1 1 0 1 0 2 0v-4a1 1 0 1 0-2 0v1h-1.062a7.965 7.965 0 0 0-1.618-3.906l1.18-1.18l.793.793a1 1 0 1 0 1.414-1.414l-3-3a1 1 0 1 0-1.414 1.414l.793.793l-1.18 1.18A7.966 7.966 0 0 0 13 4.062V3h1a1 1 0 1 0 0-2h-4zm3 13a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H14a1 1 0 0 1-1-1zm-3-3.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0zm.5-2.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5z" clip-rule="evenodd"/>`
	covidExclamationInnerSVG          = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9 2a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-1v1.062c1.46.182 2.8.759 3.906 1.618l1.18-1.18l-.793-.793a1 1 0 0 1 1.414-1.414l3 3a1 1 0 0 1-1.414 1.414l-.793-.793l-1.345 1.346A9 9 0 0 0 7.26 18.155L5.914 19.5l.793.793a1 1 0 1 1-1.414 1.414l-3-3a1 1 0 1 1 1.414-1.414l.793.793l1.18-1.18A7.966 7.966 0 0 1 4.062 13H3v1a1 1 0 1 1-2 0v-4a1 1 0 0 1 2 0v1h1.062A7.96 7.96 0 0 1 5.68 7.094L4.5 5.914l-.793.793a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.414 1.414l-.793.793l1.18 1.18A7.96 7.96 0 0 1 11 4.062V3h-1a1 1 0 0 1-1-1z"/><path d="M16 9a7 7 0 1 0 0 14a7 7 0 0 0 0-14zm-1 10a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H16a1 1 0 0 1-1-1zm2-6a1 1 0 1 0-2 0l.001 3a1 1 0 1 0 2 0L17 13z"/></g>`
	covidExclamationLineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12a7 7 0 0 1 7-7m-7 7H2m3 0c0 1.933.784 3.683 2.05 4.95L4.5 19.5M12 5V2m0 3c1.933 0 3.683.784 4.95 2.05L19.5 4.5M12 2h2m-2 0h-2M2 12v-2m0 2v2m17.5-9.5L18 3m1.5 1.5L21 6M4.5 4.5L6 3M4.5 4.5L3 6m1.5-1.5L7 7M4.5 19.5L6 21m-1.5-1.5L3 18m13 1h.001M9 10.5A1.5 1.5 0 0 1 10.5 9m5.501 7L16 13m6 3a6 6 0 1 1-12 0a6 6 0 0 1 12 0z"/>`
	covidLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 12h3m-3 0a6.978 6.978 0 0 1-2.05 4.95M19 12a6.978 6.978 0 0 0-2.05-4.95M12 19v3m0-3a6.978 6.978 0 0 0 4.95-2.05M12 19a6.978 6.978 0 0 1-4.95-2.05M5 12a7 7 0 0 1 7-7m-7 7H2m3 0c0 1.933.784 3.683 2.05 4.95M12 5V2m0 3c1.933 0 3.683.784 4.95 2.05M12 2h2m-2 0h-2m2 20h2m-2 0h-2m12-10v-2m0 2v2M2 12v-2m0 2v2m17.5-9.5L18 3m1.5 1.5L21 6m-1.5-1.5l-2.55 2.55M4.5 4.5L6 3M4.5 4.5L3 6m1.5-1.5L7 7M4.5 19.5L6 21m-1.5-1.5L3 18m1.5 1.5l2.55-2.55M19.5 19.5L18 21m1.5-1.5L21 18m-1.5 1.5l-2.55-2.55M14 14h.001"/><circle cx="10.5" cy="10.5" r="1.5"/></g>`
	covidOffInnerSVG                  = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18"/><path fill="currentColor" fill-rule="evenodd" d="M3.379 4.207L2.293 5.293a1 1 0 0 0 1.414 1.414l.793-.793l1.18 1.18A7.96 7.96 0 0 0 4.062 11H3v-1a1 1 0 0 0-2 0v4a1 1 0 1 0 2 0v-1h1.062c.182 1.46.759 2.8 1.618 3.906l-1.18 1.18l-.793-.793a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414-1.414l-.793-.793l1.18-1.18A7.965 7.965 0 0 0 11 19.938V21h-1a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-1v-1.062a7.965 7.965 0 0 0 3.906-1.618l1.18 1.18l-.793.793a1 1 0 0 0 1.414 1.414l1.086-1.086l-5.634-5.633a1.017 1.017 0 0 1-.158.012H14a1 1 0 0 1-.987-1.159l-1.208-1.208a2.5 2.5 0 0 1-3.438-3.438L3.379 4.207zm16.153 10.496A7.945 7.945 0 0 0 19.938 13H21v1a1 1 0 1 0 2 0v-4a1 1 0 1 0-2 0v1h-1.062a7.965 7.965 0 0 0-1.618-3.906l1.18-1.18l.793.793a1 1 0 1 0 1.414-1.414l-3-3a1 1 0 1 0-1.414 1.414l.793.793l-1.18 1.18A7.966 7.966 0 0 0 13 4.062V3h1a1 1 0 1 0 0-2h-4a1 1 0 0 0 0 2h1v1.062a7.95 7.95 0 0 0-1.703.406l10.235 10.235z" clip-rule="evenodd"/></g>`
	covidOffLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 12h3m-3 0a6.978 6.978 0 0 0-2.05-4.95M19 12v1m-7 6v3m0-3a6.978 6.978 0 0 0 4.95-2.05l2.55 2.55L18 21m-6-2a6.978 6.978 0 0 1-4.95-2.05M5 12H2m3 0c0 1.933.784 3.683 2.05 4.95M5 12c0-1.352.47-3.488 1.955-5.045M12 5V2m0 3c1.933 0 3.683.784 4.95 2.05M12 5h-1m1-3h2m-2 0h-2m2 20h2m-2 0h-2m12-10v-2m0 2v2M2 12v-2m0 2v2m17.5-9.5L18 3m1.5 1.5L21 6m-1.5-1.5l-2.55 2.55M3 6l1.5-1.5l2.455 2.455M4.5 19.5L6 21m-1.5-1.5L3 18m1.5 1.5l2.55-2.55M7 7l-.045-.045M14 14h.001M9.382 9.5a1.5 1.5 0 0 0 2.01 2.206M3 3l18 18"/>`
	cpuInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M9 2a1 1 0 0 1 1 1v1h4V3a1 1 0 1 1 2 0v1h1a3 3 0 0 1 3 3v1h1a1 1 0 1 1 0 2h-1v4h1a1 1 0 1 1 0 2h-1v1a3 3 0 0 1-3 3h-1v1a1 1 0 1 1-2 0v-1h-4v1a1 1 0 1 1-2 0v-1H7a3 3 0 0 1-3-3v-1H3a1 1 0 1 1 0-2h1v-4H3a1 1 0 0 1 0-2h1V7a3 3 0 0 1 3-3h1V3a1 1 0 0 1 1-1zm2 8h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1z" clip-rule="evenodd"/>`
	cpuLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 0 0-2 2v2m4-4V3m0 2h6m0 0h2a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-2M15 5V3M9 21v-2m6 2v-2M5 15H3m2 0V9m0 0H3m18 6h-2m2-6h-2m-6 1h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1z"/>`
	creditCardInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 7a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v1H2V7zm0 3v7a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-7H2zm5 2a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2H7z" fill="currentColor"/></g>`
	creditCardLineInnerSVG            = `<g fill="currentColor"><path d="M2 7a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V7zm3-1a1 1 0 0 0-1 1v1h16V7a1 1 0 0 0-1-1H5zm15 4H4v7a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-7zM6 13a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1z"/></g>`
	creditcardInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2 7a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v1H2V7zm0 3v7a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-7H2zm5 2a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	creditcardHandInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 13V5a2 2 0 0 0-2-2h-2m0 0H7a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7M13 3v10.5M9 7v3"/><path fill="currentColor" d="M18.948 9.95L16.998 8v6.587c0 .89-1.077 1.337-1.707.707L13.996 14c-.5-.5-1.701-.8-2.502 0c-.8.8-.5 2 0 2.5l4.918 4.915a2 2 0 0 0 1.414.585H20a1 1 0 0 0 1-1v-6.1a7 7 0 0 0-2.052-4.95z"/></g>`
	creditcardHandLineInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 13V5a2 2 0 0 0-2-2h-2m0 0H7a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7M13 3v10.5M9 7v3"/><path d="M21 22v-7.1a7 7 0 0 0-2.052-4.95L16.998 8v6.587c0 .89-1.077 1.337-1.707.707L13.996 14c-.5-.5-1.701-.8-2.502 0c-.8.8-.5 2 0 2.5l5.503 5.5"/></g>`
	creditcardLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9v8a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9M3 9V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 9h18M7 13h5"/>`
	creditcardPlusInnerSVG            = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 15v3m0 3v-3m0 0h-3m3 0h3"/><path fill="currentColor" fill-rule="evenodd" d="M5 4a3 3 0 0 0-3 3v1h20V7a3 3 0 0 0-3-3H5zM2 17v-7h20v3.528A6 6 0 0 0 12.341 20H5a3 3 0 0 1-3-3zm4-4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1z" clip-rule="evenodd"/></g>`
	creditcardPlusLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9v8a2 2 0 0 0 2 2h6M3 9V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 9h18m0 0v3M7 13h5m6 2v3m0 3v-3m0 0h-3m3 0h3"/>`
	crownInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M12 3a1 1 0 0 1 .832.445l3.471 5.207l4.182-2.51a1 1 0 0 1 1.503 1.01l-2 13A1 1 0 0 1 19 21H5a1 1 0 0 1-.988-.848l-2-13a1 1 0 0 1 1.503-1.01l4.182 2.51l3.471-5.207A1 1 0 0 1 12 3zm-1 11a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm1-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6z" clip-rule="evenodd"/>`
	crownLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 7l2 13h14l2-13l-5 3l-4-6l-4 6l-5-3z"/><circle cx="12" cy="14" r="2"/></g>`
	cubeInnerSVG                      = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10.658 2.553a3 3 0 0 1 2.684 0l7.105 3.553A1 1 0 0 1 21 7v.382l-9 4.5l-9-4.5V7a1 1 0 0 1 .553-.894l7.105-3.553zM3 9.618v6.146a3 3 0 0 0 1.658 2.683L11 21.618v-8l-8-4zm10 12l6.342-3.17A3 3 0 0 0 21 15.763V9.618l-8 4v8z" fill="currentColor"/></g>`
	cubeLineInnerSVG                  = `<g fill="currentColor"><path d="M12.447 4.342a1 1 0 0 0-.894 0L6.236 7L12 9.882L17.764 7l-5.317-2.658zM19 8.618l-6 3v7.764l5.447-2.724a1 1 0 0 0 .553-.894V8.618zm-8 10.764v-7.764l-6-3v7.146a1 1 0 0 0 .553.894L11 19.382zm-.342-16.83a3 3 0 0 1 2.684 0l7.105 3.554A1 1 0 0 1 21 7v8.764a3 3 0 0 1-1.658 2.683l-6.895 3.447a1 1 0 0 1-.894 0l-6.895-3.447A3 3 0 0 1 3 15.764V7a1 1 0 0 1 .553-.894l7.105-3.553z"/></g>`
	cupInnerSVG                       = `<g fill="none"><path fill="currentColor" d="M4 5h12v7a4 4 0 0 1-4 4H8a4 4 0 0 1-4-4V5z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 5H4v7a4 4 0 0 0 4 4h4a4 4 0 0 0 4-4V5zm0 0h2v0a2 2 0 0 1 2 2v4M4 19h14"/></g>`
	cupLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 5H4v7a4 4 0 0 0 4 4h4a4 4 0 0 0 4-4v-1m0-6v6m0-6h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2M4 19h14"/>`
	curlyBracesInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 5H9a2 2 0 0 0-2 2v2c0 1-.6 3-3 3c1 0 3 .6 3 3v2a2 2 0 0 0 2 2h.5m5-14h.5a2 2 0 0 1 2 2v2c0 1 .6 3 3 3c-1 0-3 .6-3 3v2a2 2 0 0 1-2 2h-.5"/>`
	curlyBracesLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 5H9a2 2 0 0 0-2 2v2c0 1-.6 3-3 3c1 0 3 .6 3 3v2a2 2 0 0 0 2 2h.5m5-14h.5a2 2 0 0 1 2 2v2c0 1 .6 3 3 3c-1 0-3 .6-3 3v2a2 2 0 0 1-2 2h-.5"/>`
	cursorClickInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8.949 2.684a1 1 0 0 0-1.898.632l1 3a1 1 0 1 0 1.898-.632l-1-3zm6.758 3.023a1 1 0 0 0-1.414-1.414l-2 2a1 1 0 0 0 1.414 1.414l2-2zM3.317 7.051a1 1 0 0 0-.633 1.898l3 1a1 1 0 1 0 .632-1.898l-3-1zm7.025 2.01a1 1 0 0 0-1.282 1.28l4 11a1 1 0 0 0 1.868.03l1.437-3.591l3.928 3.927a1 1 0 1 0 1.414-1.414l-3.928-3.928l3.592-1.436a1 1 0 0 0-.03-1.869l-11-4zm-2.635 4.646a1 1 0 1 0-1.414-1.414l-2 2a1 1 0 1 0 1.414 1.414l2-2z" fill="currentColor"/></g>`
	cursorClickLineInnerSVG           = `<g fill="currentColor"><path d="M7.684 2.051a1 1 0 0 1 1.265.633l1 3a1 1 0 0 1-1.898.632l-1-3a1 1 0 0 1 .633-1.265zm8.023 2.242a1 1 0 0 1 0 1.414l-2 2a1 1 0 1 1-1.414-1.414l2-2a1 1 0 0 1 1.414 0zM2.051 7.683a1 1 0 0 1 1.265-.632l3 1a1 1 0 1 1-.632 1.898l-3-1a1 1 0 0 1-.633-1.265zm7.242 1.61a1 1 0 0 1 1.049-.233l11 4a1 1 0 0 1 .03 1.868l-3.593 1.437l3.928 3.928a1 1 0 0 1-1.414 1.414l-3.928-3.928l-1.436 3.592a1 1 0 0 1-1.869-.03l-4-11a1 1 0 0 1 .233-1.048zm2.38 2.38l2.371 6.523l1.027-2.567a1 1 0 0 1 .558-.557l2.567-1.027l-6.524-2.373zm-3.966.62a1 1 0 0 1 0 1.414l-2 2a1 1 0 1 1-1.414-1.414l2-2a1 1 0 0 1 1.414 0z"/></g>`
	dashboardInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 9a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5zm0 13a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5zm10 0a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-4zm0-8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4z" fill="currentColor"/></g>`
	dashboardLineInnerSVG             = `<g fill="currentColor"><path d="M5 9a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5zm0-2h4V4H5v3zm0 15a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5zm0-2h4v-8H5v8zm8 0a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2zm6 0h-4v-2h4v2zm-4-6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4zm0-2h4V4h-4v8z"/></g>`
	dataInnerSVG                      = `<g fill="none"><path fill="currentColor" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7s4.03-4 9-4s9 1.79 9 4z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7m18 0c0-2.21-4.03-4-9-4S3 4.79 3 7m18 0v5M3 7v5m18 0c0 2.21-4.03 4-9 4s-9-1.79-9-4m18 0v5c0 2.21-4.03 4-9 4s-9-1.79-9-4v-5"/></g>`
	dataLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7m18 0c0-2.21-4.03-4-9-4S3 4.79 3 7m18 0v5M3 7v5m18 0c0 2.21-4.03 4-9 4s-9-1.79-9-4m18 0v5c0 2.21-4.03 4-9 4s-9-1.79-9-4v-5"/>`
	dataMinusInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7s4.03-4 9-4s9 1.79 9 4z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7m18 0c0-2.21-4.03-4-9-4S3 4.79 3 7m18 0v5M3 7v5m9 4c-4.97 0-9-1.79-9-4m0 0v5c0 2.21 4.03 4 9 4m9-3h-6"/></g>`
	dataMinusLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7m18 0c0-2.21-4.03-4-9-4S3 4.79 3 7m18 0v5M3 7v5m9 4c-4.97 0-9-1.79-9-4m0 0v5c0 2.21 4.03 4 9 4m9-3h-6"/>`
	dataPlusInnerSVG                  = `<g fill="none"><path fill="currentColor" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7s4.03-4 9-4s9 1.79 9 4z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7m18 0c0-2.21-4.03-4-9-4S3 4.79 3 7m18 0v5M3 7v5m9 4c-4.97 0-9-1.79-9-4m0 0v5c0 2.21 4.03 4 9 4m6-6v3m0 3v-3m0 0h3m-3 0h-3"/></g>`
	dataPlusLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7c0 2.21-4.03 4-9 4S3 9.21 3 7m18 0c0-2.21-4.03-4-9-4S3 4.79 3 7m18 0v5M3 7v5m9 4c-4.97 0-9-1.79-9-4m0 0v5c0 2.21 4.03 4 9 4m6-6v3m0 3v-3m0 0h3m-3 0h-3"/>`
	databaseInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5.23 3.258C7.014 2.465 9.408 2 12 2c2.592 0 4.986.465 6.77 1.258c.89.396 1.674.893 2.247 1.496C21.596 5.362 22 6.122 22 7v1c0 .151-.081.425-.516.799c-.432.371-1.116.75-2.048 1.09C17.581 10.563 14.952 11 12 11s-5.581-.437-7.436-1.111c-.932-.34-1.616-.719-2.048-1.09C2.081 8.425 2 8.15 2 8V7c0-.878.404-1.638.983-2.246c.573-.603 1.356-1.1 2.247-1.496zM2 10.884V13c0 .151.081.425.516.799c.432.371 1.116.75 2.048 1.09C6.419 15.563 9.048 16 12 16s5.581-.437 7.436-1.111c.932-.34 1.616-.719 2.048-1.09c.435-.374.516-.648.516-.799v-2.116a10.37 10.37 0 0 1-1.88.884C17.994 12.541 15.123 13 12 13c-3.123 0-5.994-.459-8.12-1.232A10.371 10.371 0 0 1 2 10.884zm20 5a10.37 10.37 0 0 1-1.88.884C17.994 17.541 15.123 18 12 18c-3.123 0-5.994-.459-8.12-1.232A10.371 10.371 0 0 1 2 15.884V17c0 .878.404 1.639.983 2.246c.573.603 1.356 1.1 2.247 1.496C7.014 21.535 9.408 22 12 22c2.592 0 4.986-.465 6.77-1.258c.89-.396 1.674-.893 2.247-1.496c.579-.607.983-1.368.983-2.246v-1.116z" fill="currentColor"/></g>`
	databaseLineInnerSVG              = `<g fill="currentColor"><path d="M4.432 6.132C4.099 6.482 4 6.773 4 7c0 .227.1.518.432.868c.337.354.872.719 1.61 1.047C7.516 9.569 9.622 10 12 10c2.379 0 4.484-.43 5.958-1.085c.738-.328 1.273-.693 1.61-1.047c.333-.35.432-.641.432-.868c0-.227-.1-.518-.432-.868c-.337-.354-.872-.719-1.61-1.047C16.484 4.431 14.378 4 12 4c-2.379 0-4.484.43-5.958 1.085c-.738.328-1.273.693-1.61 1.047zM20 10.08c-.375.246-.79.466-1.23.662C16.986 11.535 14.592 12 12 12c-2.592 0-4.986-.465-6.77-1.258A8.689 8.689 0 0 1 4 10.08V12c0 .227.1.518.432.868c.337.354.872.719 1.61 1.047C7.516 14.57 9.622 15 12 15c2.379 0 4.484-.43 5.958-1.085c.738-.328 1.273-.693 1.61-1.047c.333-.35.432-.641.432-.868v-1.92zm0 5c-.375.246-.79.466-1.23.662C16.986 16.535 14.592 17 12 17c-2.592 0-4.986-.465-6.77-1.258A8.689 8.689 0 0 1 4 15.08V17c0 .227.1.518.432.868c.337.354.872.719 1.61 1.047C7.516 19.57 9.622 20 12 20c2.379 0 4.484-.43 5.958-1.085c.738-.328 1.273-.693 1.61-1.047c.333-.35.432-.641.432-.868v-1.92zM22 17c0 .878-.404 1.639-.983 2.246c-.573.603-1.356 1.1-2.247 1.496C16.986 21.535 14.592 22 12 22c-2.592 0-4.986-.465-6.77-1.258c-.89-.396-1.674-.893-2.247-1.496C2.404 18.64 2 17.878 2 17V7c0-.878.404-1.638.983-2.246c.573-.603 1.356-1.1 2.247-1.496C7.014 2.465 9.408 2 12 2c2.592 0 4.986.465 6.77 1.258c.89.396 1.674.893 2.247 1.496C21.596 5.362 22 6.122 22 7v10z"/></g>`
	deleteBinInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M9 7h9v11a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V7h3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7h-2M4 7h2m0 0h12M6 7v11a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7m-9-.5v0A2.5 2.5 0 0 1 11.5 4h1A2.5 2.5 0 0 1 15 6.5v0"/></g>`
	deleteBinLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7v0a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v0M9 7h6M9 7H6m9 0h3m2 0h-2M4 7h2m0 0v11a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7"/>`
	desktopComputerInnerSVG           = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-4v2h1a1 1 0 1 1 0 2H8a1 1 0 1 1 0-2h1v-2H5a3 3 0 0 1-3-3V6zm11 11h-2v2h2v-2z" fill="currentColor"/></g>`
	desktopComputerLineInnerSVG       = `<g fill="currentColor"><path d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-4v2h1a1 1 0 1 1 0 2H8a1 1 0 1 1 0-2h1v-2H5a3 3 0 0 1-3-3V6zm9 11v2h2v-2h-2zM5 5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H5z"/></g>`
	deviceMobileInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H8zm3-4a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2z" fill="currentColor"/></g>`
	deviceMobileLineInnerSVG          = `<g fill="currentColor"><path d="M8 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H8zm-1-3a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v14zm3-1a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1z"/></g>`
	deviceTabletInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H7zm4-5a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2z" fill="currentColor"/></g>`
	deviceTabletLineInnerSVG          = `<g fill="currentColor"><path d="M7 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H7zm-1-3a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v14zm4-1a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1z"/></g>`
	distributeHorizontalInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 20V4m14 16V4"/><rect width="10" height="4" x="10" y="17" fill="currentColor" rx="2" transform="rotate(-90 10 17)"/></g>`
	distributeHorizontalLineInnerSVG  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 20V4m14 16V4"/><rect width="10" height="4" x="10" y="17" rx="2" transform="rotate(-90 10 17)"/></g>`
	distributeVerticalInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 5h16M4 19h16"/><rect width="10" height="4" x="7" y="10" fill="currentColor" rx="2"/></g>`
	distributeVerticalLineInnerSVG    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 5h16M4 19h16"/><rect width="10" height="4" x="7" y="10" rx="2"/></g>`
	divideInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12h14"/><circle cx="12" cy="6" r="1" fill="currentColor"/><circle cx="12" cy="18" r="1" fill="currentColor"/></g>`
	divideLineInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12h14"/><circle cx="12" cy="6" r="2"/><circle cx="12" cy="18" r="2"/></g>`
	documentInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M19 21H5c-1.126 0-1.926-.491-2.412-1.166A3.233 3.233 0 0 1 2 18V6a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v5h3a1 1 0 0 1 1 1v6c0 .493-.14 1.211-.588 1.834C20.925 20.51 20.125 21 19 21zm1-7a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4z" clip-rule="evenodd"/>`
	documentAwardInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M5 21h14c1.126 0 1.926-.491 2.412-1.166c.448-.623.588-1.34.588-1.834v-6a1 1 0 0 0-1-1h-3V6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v12c0 .493.14 1.211.588 1.834C3.074 20.51 3.874 21 5 21zm14-8a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1zM9 10a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm1-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm-3 8a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	documentAwardLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 20H5c-1.6 0-2-1.333-2-2V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6m2 8c-.667 0-2-.4-2-2v-6m2 8c1.6 0 2-1.333 2-2v-6h-4"/><path d="M7 16h6m-1-6a2 2 0 1 1-4 0a2 2 0 0 1 4 0z"/></g>`
	documentLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5c-1.6 0-2-1.333-2-2V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6m2 8c-.667 0-2-.4-2-2v-6m2 8c1.6 0 2-1.333 2-2v-6h-4"/>`
	dollarCircleInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm12-6a1 1 0 1 0-2 0v1a3 3 0 0 0 0 6h2a1 1 0 1 1 0 2H9a1 1 0 1 0 0 2h2v1a1 1 0 1 0 2 0v-1a3 3 0 1 0 0-6h-2a1 1 0 1 1 0-2h4a1 1 0 1 0 0-2h-2V6z" clip-rule="evenodd"/>`
	dollarCircleLineInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M15 8h-3m-3 8h3m0 0h1a2 2 0 0 0 2-2v0a2 2 0 0 0-2-2h-2a2 2 0 0 1-2-2v0a2 2 0 0 1 2-2h1m0 8v2m0-10V6"/></g>`
	doorEnterInnerSVG                 = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M10.138 1.815A3 3 0 0 1 14 4.688v14.624a3 3 0 0 1-3.862 2.873l-6-1.8A3 3 0 0 1 2 17.512V6.488a3 3 0 0 1 2.138-2.873l6-1.8zM15 4a1 1 0 0 1 1-1h3a3 3 0 0 1 3 3v1a1 1 0 1 1-2 0V6a1 1 0 0 0-1-1h-3a1 1 0 0 1-1-1zm6 12a1 1 0 0 1 1 1v1a3 3 0 0 1-3 3h-3a1 1 0 1 1 0-2h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1zM9 11a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H9z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12h-5m0 0l2-2m-2 2l2 2"/></g>`
	doorEnterLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4h3a2 2 0 0 1 2 2v1m-5 13h3a2 2 0 0 0 2-2v-1M4.425 19.428l6 1.8A2 2 0 0 0 13 19.312V4.688a2 2 0 0 0-2.575-1.916l-6 1.8A2 2 0 0 0 3 6.488v11.024a2 2 0 0 0 1.425 1.916zM9.001 12H9m12 0h-5m0 0l2-2m-2 2l2 2"/>`
	doorExitInnerSVG                  = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M10.138 1.815A3 3 0 0 1 14 4.688v14.624a3 3 0 0 1-3.862 2.873l-6-1.8A3 3 0 0 1 2 17.512V6.488a3 3 0 0 1 2.138-2.873l6-1.8zM15 4a1 1 0 0 1 1-1h3a3 3 0 0 1 3 3v1a1 1 0 1 1-2 0V6a1 1 0 0 0-1-1h-3a1 1 0 0 1-1-1zm6 12a1 1 0 0 1 1 1v1a3 3 0 0 1-3 3h-3a1 1 0 1 1 0-2h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1zM9 11a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H9z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12h5m0 0l-2-2m2 2l-2 2"/></g>`
	doorExitLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4h3a2 2 0 0 1 2 2v1m-5 13h3a2 2 0 0 0 2-2v-1M4.425 19.428l6 1.8A2 2 0 0 0 13 19.312V4.688a2 2 0 0 0-2.575-1.916l-6 1.8A2 2 0 0 0 3 6.488v11.024a2 2 0 0 0 1.425 1.916zM9.001 12H9m7 0h5m0 0l-2-2m2 2l-2 2"/>`
	dotsCircleInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm5.5 1.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zm3-1.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0zm4.5 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0z" fill="currentColor"/></g>`
	dotsCircleLineInnerSVG            = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm0 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm3 2a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm-8-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4z"/></g>`
	dotsHorizontalInnerSVG            = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4z" fill="currentColor"/></g>`
	dotsHorizontalLineInnerSVG        = `<g fill="currentColor"><path d="M2 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm3-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm4 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm3-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm7-2a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm-1 3a1 1 0 1 1 2 0a1 1 0 0 1-2 0z"/></g>`
	dotsVerticalInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10 5a2 2 0 1 0 4 0a2 2 0 0 0-4 0zm0 7a2 2 0 1 0 4 0a2 2 0 0 0-4 0zm0 7a2 2 0 1 0 4 0a2 2 0 0 0-4 0z" fill="currentColor"/></g>`
	dotsVerticalLineInnerSVG          = `<g fill="currentColor"><path d="M9 5a3 3 0 1 0 6 0a3 3 0 0 0-6 0zm3 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm-1-3a1 1 0 1 0 2 0a1 1 0 0 0-2 0zm1 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm-1-3a1 1 0 1 0 2 0a1 1 0 0 0-2 0z"/></g>`
	duplicateInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M16 6a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v6a4 4 0 0 0 4 4h1v-4.5c0-1.763.746-2.913 1.708-3.606C9.644 7.22 10.753 7 11.5 7H16V6zm.5 2h-5c-.587 0-1.478.18-2.208.706C8.588 9.213 8 10.063 8 11.5V18a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-6a4 4 0 0 0-4-4h-1.5z" fill="currentColor"/></g>`
	duplicateLineInnerSVG             = `<g fill="currentColor"><path d="M16 6a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v6a4 4 0 0 0 4 4h2v2a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-6a4 4 0 0 0-4-4h-2V6zm-2 2h-2a4 4 0 0 0-4 4v2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2zm4 2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h6z"/></g>`
	earthSphereInnerSVG               = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12a9.966 9.966 0 0 1-.832 4M12 22a9.966 9.966 0 0 0 7.071-2.929M2 12a9.966 9.966 0 0 1 2.929-7.071M12 2a9.966 9.966 0 0 0-4 .832m0 18.336A10.02 10.02 0 0 1 2.832 16m13-13A10.02 10.02 0 0 1 21 8.168"/><path fill="currentColor" fill-rule="evenodd" d="M9.06 5.646A6.977 6.977 0 0 1 12 5a6.998 6.998 0 0 1 6.184 3.718l-2.015.671a3 3 0 0 0-1.851 3.923l1.7 4.42A6.97 6.97 0 0 1 12 19a6.985 6.985 0 0 1-2.519-.467l1.564-1.369a2.44 2.44 0 0 0-.515-4.017a.44.44 0 0 1-.067-.744l.473-.355a2.804 2.804 0 0 0 .3-4.226L9.06 5.646zm-1.708 1.12A6.984 6.984 0 0 0 5 12c0 2.221 1.035 4.2 2.646 5.481l2.082-1.822a.44.44 0 0 0-.093-.723c-1.621-.811-1.823-3.045-.372-4.133l.473-.355a.804.804 0 0 0 .086-1.212l-2.47-2.47zm10.23 9.46A6.973 6.973 0 0 0 19 12c0-.479-.048-.947-.14-1.4l-2.058.687a1 1 0 0 0-.617 1.307l1.396 3.632z" clip-rule="evenodd"/></g>`
	earthSphereLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12a9.966 9.966 0 0 1-.832 4M12 22a9.966 9.966 0 0 0 7.071-2.929M2 12a9.966 9.966 0 0 1 2.929-7.071M12 2a9.966 9.966 0 0 0-4 .832m0 18.336A10.02 10.02 0 0 1 2.832 16m13-13A10.02 10.02 0 0 1 21 8.168M9 6.803A5.998 5.998 0 0 0 6 12c0 2.22 1.207 4.16 3 5.197M9 6.803A6.003 6.003 0 0 1 17.659 10M9 6.803l1.603 1.708a1.83 1.83 0 0 1-.236 2.714l-.59.442a1.373 1.373 0 0 0 .21 2.327v0c.92.46 1.027 1.729.198 2.336L9 17.197m0 0c.883.51 1.907.803 3 .803c1.792 0 3.4-.786 4.5-2.031M17.659 10A5.99 5.99 0 0 1 18 12a5.977 5.977 0 0 1-1.5 3.969M17.659 10l-1.101.349a2 2 0 0 0-1.252 2.653L16.5 15.97"/>`
	editPenFourInnerSVG               = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M17.586 2a2 2 0 0 1 2.828 0L22 3.586a2 2 0 0 1 0 2.828L20.414 8L16 3.586L17.586 2zm-3 3l-5 5A2 2 0 0 0 9 11.414V13a2 2 0 0 0 2 2h1.586A2 2 0 0 0 14 14.414l5-5L14.586 5z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 14H5a2 2 0 0 0-2 2v0a2 2 0 0 0 2 2h14a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2h-4"/></g>`
	editPenFourLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 5l2.293-2.293a1 1 0 0 1 1.414 0l1.586 1.586a1 1 0 0 1 0 1.414L19 8m-3-3l-5.707 5.707a1 1 0 0 0-.293.707V13a1 1 0 0 0 1 1h1.586a1 1 0 0 0 .707-.293L19 8m-3-3l3 3M6 14H5a2 2 0 0 0-2 2v0a2 2 0 0 0 2 2h14a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2h-4"/>`
	editPenTwoInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M15.586 3a2 2 0 0 1 2.828 0L21 5.586a2 2 0 0 1 0 2.828L19.414 10L14 4.586L15.586 3zm-3 3l-9 9A2 2 0 0 0 3 16.414V19a2 2 0 0 0 2 2h2.586A2 2 0 0 0 9 20.414l9-9L12.586 6z" clip-rule="evenodd"/>`
	editPenTwoLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 6l2.293-2.293a1 1 0 0 1 1.414 0l2.586 2.586a1 1 0 0 1 0 1.414L18 10m-4-4l-9.707 9.707a1 1 0 0 0-.293.707V19a1 1 0 0 0 1 1h2.586a1 1 0 0 0 .707-.293L18 10m-4-4l4 4"/>`
	ejectInnerSVG                     = `<g fill="none"><path fill="currentColor" d="m5 15l7-10l7 10H5z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19h14M12 5L5 15h14L12 5z"/></g>`
	ejectLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19h14M12 5L5 15h14L12 5z"/>`
	emojiHappyInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm7-1a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm.757 2.1a1 1 0 0 0-1.514 1.307C9.053 15.344 10.283 16 12 16c1.716 0 2.947-.656 3.757-1.593a1 1 0 1 0-1.514-1.307c-.419.485-1.091.9-2.243.9s-1.824-.415-2.243-.9zM16 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0z" fill="currentColor"/></g>`
	emojiHappyLineInnerSVG            = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zM9.757 13.1a1 1 0 1 0-1.514 1.307C9.053 15.344 10.283 16 12 16c1.716 0 2.947-.656 3.757-1.593a1 1 0 0 0-1.514-1.307c-.419.485-1.091.9-2.243.9s-1.824-.415-2.243-.9zM8 10a1 1 0 0 1 1-1h.01a1 1 0 0 1 0 2H9a1 1 0 0 1-1-1zm7-1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H15z"/></g>`
	emojiSadInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm7-1a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm.757 4.654a1 1 0 1 1-1.514-1.308c.81-.937 2.04-1.592 3.757-1.592c1.716 0 2.947.655 3.757 1.592a1 1 0 0 1-1.514 1.308c-.419-.485-1.091-.9-2.243-.9s-1.824.415-2.243.9zM16 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0z" fill="currentColor"/></g>`
	emojiSadLineInnerSVG              = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zM9.757 15.654a1 1 0 1 1-1.514-1.308c.81-.937 2.04-1.592 3.757-1.592c1.716 0 2.947.655 3.757 1.592a1 1 0 1 1-1.514 1.308c-.419-.485-1.091-.9-2.243-.9s-1.824.415-2.243.9zM8 10a1 1 0 0 1 1-1h.01a1 1 0 0 1 0 2H9a1 1 0 0 1-1-1zm7-1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H15z"/></g>`
	eraserInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M3.121 17.85a3 3 0 0 1 0-4.243l7.071-7.072l8.486 8.486l-3.243 3.242H20a1 1 0 1 1 0 2H6.778a3 3 0 0 1-2.121-.878L3.12 17.849zm16.97-4.243l1.415-1.415a3 3 0 0 0 0-4.242l-4.243-4.243a3 3 0 0 0-4.242 0l-1.414 1.414l8.485 8.486z" clip-rule="evenodd"/>`
	eraserLineInnerSVG                = `<g fill="none"><g clip-path="url(#majesticonsEraserLine0)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.607 6.535l-7.779 7.779a2 2 0 0 0 0 2.828l1.536 1.536a2 2 0 0 0 1.414.585h6.243M11.607 6.535l2.12-2.12a2 2 0 0 1 2.83 0l4.242 4.242a2 2 0 0 1 0 2.828l-2.121 2.122m-7.071-7.072l7.07 7.072m0 0l-5.656 5.656m0 0H20"/></g><defs><clipPath id="majesticonsEraserLine0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	etheriumCircleInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm11.814-7.581a1 1 0 0 0-1.628 0l-5 7a1 1 0 0 0 0 1.162l5 7a1 1 0 0 0 1.628 0l5-7a1 1 0 0 0 0-1.162l-5-7zM12 12.923L9.693 12L12 11.077l2.307.923l-2.307.923zm.371 2.005l1.832-.732L12 17.279l-2.203-3.083l1.832.732a1 1 0 0 0 .742 0zM12 6.72L9.797 9.804l1.832-.732a1 1 0 0 1 .742 0l1.832.732L12 6.72z" clip-rule="evenodd"/>`
	etheriumCircleLineInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="m7 12l5-7l5 7M7 12l5 7l5-7M7 12l5-2l5 2M7 12l5 2l5-2"/></g>`
	euroCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm8 0c0-.17.01-.336.031-.5H12a1 1 0 1 0 0-2H9.877A3.993 3.993 0 0 1 13 8c.902 0 1.731.297 2.4.8a1 1 0 0 0 1.2-1.6a6.001 6.001 0 0 0-9.057 2.3H7a1 1 0 0 0 0 2h.02a6.081 6.081 0 0 0 0 1H7a1 1 0 1 0 0 2h.544a6.001 6.001 0 0 0 9.057 2.3a1 1 0 0 0-1.201-1.6c-.669.503-1.498.8-2.4.8a3.992 3.992 0 0 1-3.123-1.5H12a1 1 0 1 0 0-2H9.031A4.039 4.039 0 0 1 9 12z" clip-rule="evenodd"/>`
	euroCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M16 8a5 5 0 1 0 0 8m-9-5.5h5m0 3H7"/></g>`
	exclamationInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M9.378 4.661c1.143-2.057 4.101-2.057 5.245 0l6.6 11.882c1.111 2-.335 4.457-2.622 4.457H5.399c-2.287 0-3.733-2.457-2.622-4.457l6.6-11.882zM12 8a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1zm1 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0z" fill="currentColor"/></g>`
	exclamationCircleInnerSVG         = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm10-5a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1zm1 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0z" fill="currentColor"/></g>`
	exclamationCircleLineInnerSVG     = `<g fill="currentColor"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm10-5a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1zm0 8a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H12z"/></g>`
	exclamationLineInnerSVG           = `<g fill="currentColor"><path d="M11.126 5.633a1 1 0 0 1 1.748 0l6.601 11.881A1 1 0 0 1 18.601 19H5.399a1 1 0 0 1-.874-1.486l6.6-11.881zm3.497-.972c-1.143-2.057-4.102-2.057-5.245 0L2.777 16.543C1.666 18.543 3.112 21 5.399 21h13.202c2.287 0 3.733-2.457 2.622-4.457l-6.6-11.882zM12 8a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1zm-1 8a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H12a1 1 0 0 1-1-1z"/></g>`
	externalLinkInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6zm10-1a1 1 0 1 0 0 2h2.586l-4.293 4.293a1 1 0 0 0 1.414 1.414L17 8.414V11a1 1 0 1 0 2 0V6a1 1 0 0 0-1-1h-5z" fill="currentColor"/></g>`
	externalLinkLineInnerSVG          = `<g fill="currentColor"><path d="M5 6a1 1 0 0 1 1-1h4a1 1 0 1 0 0-2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-4a1 1 0 1 0-2 0v4a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6zm10-3a1 1 0 1 0 0 2h2.586l-6.293 6.293a1 1 0 0 0 1.414 1.414L19 6.414V9a1 1 0 1 0 2 0V4a1 1 0 0 0-1-1h-5z"/></g>`
	eyeInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M4.19 7.262C5.94 5.577 8.517 4 12 4c3.483 0 6.06 1.577 7.81 3.262a15.086 15.086 0 0 1 3.001 4.11c.193.399.193.857 0 1.255a15.086 15.086 0 0 1-3 4.111C18.06 18.423 15.483 20 12 20c-3.483 0-6.06-1.577-7.81-3.262a15.088 15.088 0 0 1-3.001-4.11a1.435 1.435 0 0 1 0-1.255a15.088 15.088 0 0 1 3-4.111zM12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6z" clip-rule="evenodd"/>`
	eyeLineInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 5c-6.307 0-9.367 5.683-9.91 6.808a.435.435 0 0 0 0 .384C2.632 13.317 5.692 19 12 19s9.367-5.683 9.91-6.808a.435.435 0 0 0 0-.384C21.368 10.683 18.308 5 12 5z"/><circle cx="12" cy="12" r="3"/></g>`
	eyeOffInnerSVG                    = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18"/><path fill="currentColor" fill-rule="evenodd" d="M5.4 6.23c-.44.33-.843.678-1.21 1.032a15.088 15.088 0 0 0-3.001 4.11a1.435 1.435 0 0 0 0 1.255a15.088 15.088 0 0 0 3 4.111C5.94 18.423 8.518 20 12 20c2.236 0 4.1-.65 5.61-1.562l-3.944-3.943a3 3 0 0 1-4.161-4.161L5.401 6.229zm15.266 9.608a15.06 15.06 0 0 0 2.145-3.21a1.435 1.435 0 0 0 0-1.255a15.086 15.086 0 0 0-3-4.111C18.06 5.577 15.483 4 12 4a10.83 10.83 0 0 0-2.808.363l11.474 11.475z" clip-rule="evenodd"/></g>`
	eyeOffLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6.362A9.707 9.707 0 0 1 12 5c6.307 0 9.367 5.683 9.91 6.808c.06.123.06.261 0 .385c-.352.728-1.756 3.362-4.41 5.131M14 18.8a9.93 9.93 0 0 1-2 .2c-6.307 0-9.367-5.683-9.91-6.808a.44.44 0 0 1 0-.386c.219-.452.84-1.632 1.91-2.885m6 .843A3 3 0 0 1 14.236 14M3 3l18 18"/>`
	fastForwardInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M21.067 10.8l-6.667-5C13.411 5.058 12 5.764 12 7v3L6.4 5.8C5.411 5.058 4 5.764 4 7v10c0 1.236 1.411 1.942 2.4 1.2L12 14v3c0 1.236 1.411 1.942 2.4 1.2l6.667-5c.8-.6.8-1.8 0-2.4z" fill="currentColor"/></g>`
	fastForwardLineInnerSVG           = `<g fill="currentColor"><path d="M14 8v8l5.333-4L14 8zm-2-1c0-1.236 1.411-1.942 2.4-1.2l6.667 5c.8.6.8 1.8 0 2.4l-6.667 5c-.989.742-2.4.036-2.4-1.2v-3l-5.6 4.2c-.989.742-2.4.036-2.4-1.2V7c0-1.236 1.411-1.942 2.4-1.2L12 10V7zm-.667 5L6 8v8l5.333-4z"/></g>`
	ferrisWheelInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="6"/><circle cx="12" cy="4" r="2" fill="currentColor"/><circle cx="19" cy="8" r="2" fill="currentColor"/><circle cx="5" cy="8" r="2" fill="currentColor"/><circle cx="5" cy="16" r="2" fill="currentColor"/><circle cx="19" cy="16" r="2" fill="currentColor"/><path d="m8 22l4-10l4 10"/></g>`
	ferrisWheelLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="6"/><circle cx="12" cy="4" r="2"/><circle cx="19" cy="8" r="2"/><circle cx="5" cy="8" r="2"/><circle cx="5" cy="16" r="2"/><circle cx="19" cy="16" r="2"/><path d="m8 22l4-10l4 10"/></g>`
	fileInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M12 2H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-8h-6a3 3 0 0 1-3-3V2zm9 7v-.172a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H14v6a1 1 0 0 0 1 1h6z" clip-rule="evenodd"/>`
	fileAddInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H6a3 3 0 0 0-3 3v14zm10-7a1 1 0 1 0-2 0v2H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-2z" fill="currentColor"/></g>`
	fileAddLineInnerSVG               = `<g fill="currentColor"><path d="M3 19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H6a3 3 0 0 0-3 3v14zm3 1a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h7v3a3 3 0 0 0 3 3h3v9a1 1 0 0 1-1 1H6zM18.586 8H16a1 1 0 0 1-1-1V4.414L18.586 8z"/><path d="M12 11a1 1 0 0 1 1 1v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9a1 1 0 1 1 0-2h2v-2a1 1 0 0 1 1-1z"/></g>`
	fileDownloadInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H6a3 3 0 0 0-3 3v14zm10-7a1 1 0 1 0-2 0v3.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L13 15.586V12z" fill="currentColor"/></g>`
	fileDownloadLineInnerSVG          = `<g fill="currentColor"><path d="M3 19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H6a3 3 0 0 0-3 3v14zm3 1a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h7v3a3 3 0 0 0 3 3h3v9a1 1 0 0 1-1 1H6zM18.586 8H16a1 1 0 0 1-1-1V4.414L18.586 8z"/><path d="M12 11a1 1 0 0 1 1 1v3.586l1.293-1.293a1 1 0 0 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L11 15.586V12a1 1 0 0 1 1-1z"/></g>`
	fileDuplicateInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M16.83 20H8a3 3 0 0 1-3-3V6.17A3.001 3.001 0 0 0 3 9v10a3 3 0 0 0 3 3h8a3.001 3.001 0 0 0 2.83-2zM7 5v12a1 1 0 0 0 1 1h10a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H10a3 3 0 0 0-3 3z" fill="currentColor"/></g>`
	fileDuplicateLineInnerSVG         = `<g fill="currentColor"><path d="M10 4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V8.828a1 1 0 0 0-.293-.707L14.88 4.293A1 1 0 0 0 14.172 4H10zM7 5a3 3 0 0 1 3-3h4.172a3 3 0 0 1 2.12.879l3.83 3.828A3 3 0 0 1 21 8.828V15a3 3 0 0 1-3 3h-8a3 3 0 0 1-3-3V5z"/><path d="M3 9a3 3 0 0 1 3-3h2a1 1 0 0 1 0 2H6a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-2a1 1 0 1 1 2 0v2a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V9zm11-7a1 1 0 0 1 1 1v4a1 1 0 0 0 1 1h4a1 1 0 1 1 0 2h-4a3 3 0 0 1-3-3V3a1 1 0 0 1 1-1z"/></g>`
	fileLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 9v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h8m6 6v-.172a2 2 0 0 0-.586-1.414l-3.828-3.828A2 2 0 0 0 14.172 3H14m6 6h-4a2 2 0 0 1-2-2V3"/>`
	fileMinusInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 19h-6"/><path fill="currentColor" fill-rule="evenodd" d="M6 2h6v6a3 3 0 0 0 3 3h6v2.803A6 6 0 0 0 12.803 22H6a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3zm15 6.828V9h-6a1 1 0 0 1-1-1V2h.172a3 3 0 0 1 2.12.879l3.83 3.828A3 3 0 0 1 21 8.828z" clip-rule="evenodd"/></g>`
	fileMinusLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 19h-6m5-10v-.172a2 2 0 0 0-.586-1.414l-3.828-3.828A2 2 0 0 0 14.172 3H14m6 6h-4a2 2 0 0 1-2-2V3m6 6v3m-6-9H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7"/>`
	filePlusInnerSVG                  = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 16v3m0 0v3m0-3h3m-3 0h-3"/><path fill="currentColor" fill-rule="evenodd" d="M6 2h6v6a3 3 0 0 0 3 3h6v2.341A6 6 0 0 0 13.803 22H6a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3zm15 6.828V9h-6a1 1 0 0 1-1-1V2h.172a3 3 0 0 1 2.12.879l3.83 3.828A3 3 0 0 1 21 8.828z" clip-rule="evenodd"/></g>`
	filePlusLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 16v3m0 0v3m0-3h3m-3 0h-3m4-10v-.172a2 2 0 0 0-.586-1.414l-3.828-3.828A2 2 0 0 0 14.172 3H14m6 6h-4a2 2 0 0 1-2-2V3m6 6v3m-6-9H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"/>`
	fileRemoveInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H6a3 3 0 0 0-3 3v14zm6-5a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H9z" fill="currentColor"/></g>`
	fileRemoveLineInnerSVG            = `<g fill="currentColor"><path d="M6 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8.172a3 3 0 0 1 2.12.879l3.83 3.828A3 3 0 0 1 21 8.828V19a3 3 0 0 1-3 3H6zm-1-3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-9h-3a3 3 0 0 1-3-3V4H6a1 1 0 0 0-1 1v14zM16 8h2.586L15 4.414V7a1 1 0 0 0 1 1zm-7 6a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H9z"/></g>`
	fileReportInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H6a3 3 0 0 0-3 3v14zm14-8a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0v-6zm-4 2a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4zm-4 1.995a1 1 0 0 0-2 .01l.01 2a1 1 0 0 0 2-.01l-.01-2z" fill="currentColor"/></g>`
	fileReportLineInnerSVG            = `<g fill="currentColor"><path d="M6 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8.172a3 3 0 0 1 2.12.879l3.83 3.828A3 3 0 0 1 21 8.828V19a3 3 0 0 1-3 3H6zm-1-3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-9h-3a3 3 0 0 1-3-3V4H6a1 1 0 0 0-1 1v14zM16 8h2.586L15 4.414V7a1 1 0 0 0 1 1zm0 5a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4zm-3 2a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2zm-4 1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H9z"/></g>`
	fileSearchInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H6a3 3 0 0 0-3 3v14zm7-4.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0zm1.5-3.5a3.5 3.5 0 1 0 1.665 6.58l1.128 1.127a1 1 0 0 0 1.414-1.414l-1.128-1.128A3.5 3.5 0 0 0 11.5 11z" fill="currentColor"/></g>`
	fileSearchLineInnerSVG            = `<g fill="currentColor"><path d="M6 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8.172a3 3 0 0 1 2.12.879l3.83 3.828A3 3 0 0 1 21 8.828V19a3 3 0 0 1-3 3H6zm-1-3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-9h-3a3 3 0 0 1-3-3V4H6a1 1 0 0 0-1 1v14zM16 8h2.586L15 4.414V7a1 1 0 0 0 1 1zm-6 6.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0zm1.5-3.5a3.5 3.5 0 1 0 1.665 6.58l1.128 1.127a1 1 0 0 0 1.414-1.414l-1.128-1.128A3.5 3.5 0 0 0 11.5 11z"/></g>`
	fileTextInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.828a3 3 0 0 0-.879-2.12l-3.828-3.83A3 3 0 0 0 14.172 2H6a3 3 0 0 0-3 3v14zm5-7a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H8z" fill="currentColor"/></g>`
	fileTextLineInnerSVG              = `<g fill="currentColor"><path d="M6 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8.172a3 3 0 0 1 2.12.879l3.83 3.828A3 3 0 0 1 21 8.828V19a3 3 0 0 1-3 3H6zm-1-3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-9h-3a3 3 0 0 1-3-3V4H6a1 1 0 0 0-1 1v14zM16 8h2.586L15 4.414V7a1 1 0 0 0 1 1zm-8 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H8z"/></g>`
	filmInnerSVG                      = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6zM5 6a1 1 0 0 1 1-1h1v2H5V6zm14 0v1h-2V5h1a1 1 0 0 1 1 1zm-2 5V9h2v2h-2zm0 4v-2h2v2h-2zm1 4h-1v-2h2v1a1 1 0 0 1-1 1zM5 18v-1h2v2H6a1 1 0 0 1-1-1zm2-3H5v-2h2v2zm0-4H5V9h2v2z" fill="currentColor"/></g>`
	filmLineInnerSVG                  = `<g fill="currentColor"><path d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v1h2V5H6zm3 0v6h6V5H9zm8 0v2h2V6a1 1 0 0 0-1-1h-1zm2 4h-2v2h2V9zm0 4h-2v2h2v-2zm0 4h-2v2h1a1 1 0 0 0 1-1v-1zm-4 2v-6H9v6h6zm-8 0v-2H5v1a1 1 0 0 0 1 1h1zm-2-4h2v-2H5v2zm0-4h2V9H5v2z"/></g>`
	filterInnerSVG                    = `<path fill="currentColor" stroke="currentColor" stroke-width="2" d="M18 4H6c-1.105 0-2.026.91-1.753 1.98a8.018 8.018 0 0 0 4.298 5.238c.823.394 1.455 1.168 1.455 2.08v6.084a1 1 0 0 0 1.447.894l2-1a1 1 0 0 0 .553-.894v-5.084c0-.912.632-1.686 1.454-2.08a8.017 8.017 0 0 0 4.3-5.238C20.025 4.91 19.103 4 18 4z"/>`
	filterLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 4H6c-1.105 0-2.026.91-1.753 1.98a8.018 8.018 0 0 0 4.298 5.238c.823.394 1.455 1.168 1.455 2.08v6.084a1 1 0 0 0 1.447.894l2-1a1 1 0 0 0 .553-.894v-5.084c0-.912.632-1.686 1.454-2.08a8.017 8.017 0 0 0 4.3-5.238C20.025 4.91 19.103 4 18 4z"/>`
	fireInnerSVG                      = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11.445 2.168a1 1 0 0 1 1.325.194c.468.565 1.852 1.93 3.447 2.687A8.998 8.998 0 0 1 21 13A9 9 0 1 1 5.4 6.882a1 1 0 0 1 1.644.267c.127.282.471.663.992 1.066l.083.064c.2-1.41.632-2.592 1.145-3.538c.686-1.268 1.53-2.139 2.181-2.573zm.476 8.017a1 1 0 0 1 1.369.202c.042.054.158.186.323.334a2.536 2.536 0 0 0 .57.396c1.1.633 1.817 1.85 1.817 3.216C16 16.297 14.493 18 12.5 18S9 16.297 9 14.333c0-.943.341-1.812.912-2.469a1 1 0 0 1 .994-.314c.039-.092.08-.18.122-.264c.253-.498.582-.88.893-1.101z" fill="currentColor"/></g>`
	fireLineInnerSVG                  = `<g fill="currentColor"><path d="M12.77 2.362a1 1 0 0 0-1.325-.194c-.65.434-1.495 1.305-2.181 2.573a10.504 10.504 0 0 0-1.145 3.538a7.593 7.593 0 0 1-.083-.064c-.52-.403-.865-.784-.992-1.066A1 1 0 0 0 5.4 6.882a9 9 0 1 0 10.817-1.833c-1.595-.756-2.98-2.122-3.447-2.687zm-1.748 3.33c.3-.553.623-.984.91-1.3c.789.774 2.008 1.786 3.404 2.452a7 7 0 1 1-9.214 2.354c.225.223.463.423.688.598c.613.475 1.27.863 1.743 1.099A1 1 0 0 0 10 10c0-1.845.465-3.277 1.022-4.307zm2.268 4.695a1 1 0 0 0-1.37-.202c-.31.22-.639.603-.892 1.101a4.017 4.017 0 0 0-.122.264a1 1 0 0 0-.994.315A3.756 3.756 0 0 0 9 14.332C9 16.297 10.507 18 12.5 18s3.5-1.703 3.5-3.667c0-1.366-.717-2.583-1.818-3.216a2.54 2.54 0 0 1-.57-.396a2.895 2.895 0 0 1-.322-.334zM11 14.333c0-.055.002-.11.007-.164l.085.047a1 1 0 0 0 1.47-.883c0-.318.047-.584.113-.8c.163.119.34.232.529.329c.457.271.796.814.796 1.471c0 .983-.731 1.667-1.5 1.667s-1.5-.684-1.5-1.667z"/></g>`
	fishInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M14.346 6.005A13.18 13.18 0 0 0 14 6c-3.633 0-7.031 2.294-9.46 4.574a20.998 20.998 0 0 1-1.716-2.14a1 1 0 1 0-1.648 1.133A23.129 23.129 0 0 0 3.133 12c-.813.89-1.48 1.74-1.957 2.434a1 1 0 1 0 1.648 1.133c.41-.596.994-1.345 1.716-2.141C6.97 15.706 10.367 18 14 18c.123 0 .245-.002.365-.005a10.755 10.755 0 0 1-.592-1.732a19.124 19.124 0 0 1-.545-3.332c-.167-2.218.022-4.825 1.118-6.926zm2.122 11.742c.09-.02.18-.04.269-.062c2.017-.492 3.559-1.843 4.562-2.985a13.663 13.663 0 0 0 1.47-2.029c.036-.06.063-.109.082-.143l.023-.042l.007-.013l.002-.004v-.001l.001-.001l-.787-.416l.787.415l.247-.466l-.247-.466l-.787.415l.787-.416l-.001-.002l-.002-.005l-.007-.012l-.023-.042a11.436 11.436 0 0 0-.387-.637A13.66 13.66 0 0 0 21.3 9.3c-1.003-1.142-2.546-2.493-4.562-2.985a10.316 10.316 0 0 0-.27-.062a.998.998 0 0 1-.12.278c-1.033 1.648-1.294 4.005-1.125 6.25a17.16 17.16 0 0 0 .486 2.978c.226.867.47 1.447.632 1.698c.059.092.101.19.128.29zm-5.26-7.04a1 1 0 0 0-1.415-1.414c-.272.271-.47.662-.606 1.015a5.387 5.387 0 0 0-.318 1.32c-.104.928.016 2.172.924 3.08a1 1 0 0 0 1.414-1.415c-.292-.293-.422-.8-.35-1.445c.033-.306.11-.594.197-.822c.073-.19.134-.286.152-.313c.006-.01.007-.011.001-.006zM17 11a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H18a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`
	fishLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M2 15c1.833-2.667 6.8-8 12-8c.923 0 1.754.105 2.5.287M2 9c1.833 2.667 6.8 8 12 8c.923 0 1.754-.105 2.5-.287m0 0C19.96 15.87 22 12 22 12s-2.04-3.87-5.5-4.713m0 9.426c-1-1.546-2.4-5.597 0-9.426M12 10.5c-.5.5-1.2 1.8 0 3m6-2.5h.001"/>`
	flagInnerSVG                      = `<g fill="none"><path fill="currentColor" d="M19 5H5v9h14V5z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 20v-6m0-9h14v9H5m0-9v9m0-9V4"/></g>`
	flagLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 20v-6m0-9h14v9H5m0-9v9m0-9V4"/>`
	flaskInnerSVG                     = `<g fill="none"><path fill="currentColor" d="M14 8.46V3h-4v5.46a2 2 0 0 1-.272 1.007L8.25 12l-3.495 5.992C3.977 19.326 4.938 21 6.482 21h11.036c1.543 0 2.505-1.674 1.727-3.008L15.75 12l-1.478-2.533A2 2 0 0 1 14 8.459z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3h2m6 0h-2m0 0v5.46a2 2 0 0 0 .272 1.007L15.75 12l3.495 5.992c.778 1.334-.184 3.008-1.727 3.008H6.482c-1.544 0-2.505-1.674-1.727-3.008L8.25 12l1.478-2.533A2 2 0 0 0 10 8.459V3m4 0h-4"/></g>`
	flaskLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3h2m6 0h-2m0 0v5.46a2 2 0 0 0 .272 1.007L15.75 12l3.495 5.992c.778 1.334-.184 3.008-1.727 3.008H6.482c-1.544 0-2.505-1.674-1.727-3.008L8.25 12l1.478-2.533A2 2 0 0 0 10 8.459V3m4 0h-4"/>`
	flowerTwoInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M4 14c0 2.333 1.4 7 7 7c0-2.333-1.4-7-7-7zm3-6V4l2.5 2L12 3l2.5 3L17 4v4c0 1.667-1 5-5 5S7 9.667 7 8zm13 6c0 2.333-1.4 7-7 7c0-2.333 1.4-7 7-7z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 21c-5.6 0-7-4.667-7-7c5.6 0 7 4.667 7 7zm0 0h1m0 0v-8m0 8h1m-1-8c-4 0-5-3.333-5-5V4l2.5 2L12 3l2.5 3L17 4v4c0 1.667-1 5-5 5zm1 8c5.6 0 7-4.667 7-7c-5.6 0-7 4.667-7 7z"/></g>`
	flowerTwoLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13c-4 0-5-3.333-5-5V4l2.5 2L12 3l2.5 3L17 4v4c0 1.667-1 5-5 5zm0 0v8m1 0c5.6 0 7-4.667 7-7c-5.6 0-7 4.667-7 7zm0 0h-1m-1 0c-5.6 0-7-4.667-7-7c5.6 0 7 4.667 7 7zm0 0h1"/>`
	folderInnerSVG                    = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6a2 2 0 0 1 2-2h3.93a2 2 0 0 1 1.664.89l.812 1.22A2 2 0 0 0 13.07 7H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V6z"/>`
	folderAddInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h3.93a3 3 0 0 1 2.496 1.336l.812 1.219A1 1 0 0 0 13.07 6H19a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm11 5a1 1 0 1 0-2 0v2H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-2z" fill="currentColor"/></g>`
	folderAddLineInnerSVG             = `<g fill="currentColor"><path d="M4 6a1 1 0 0 1 1-1h3.93a1 1 0 0 1 .832.445l.812 1.22A3 3 0 0 0 13.07 8H19a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V6zm1-3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-5.93a1 1 0 0 1-.832-.445l-.812-1.22A3 3 0 0 0 8.93 3H5zm8 8a1 1 0 1 0-2 0v2H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-2z"/></g>`
	folderCheckInnerSVG               = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M5 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7.174A6.5 6.5 0 0 1 21 12.498V9a2 2 0 0 0-2-2h-5.93a2 2 0 0 1-1.664-.89l-.812-1.22A2 2 0 0 0 8.93 4H5z" clip-rule="evenodd"/><path fill="currentColor" d="M12.174 20v1a1 1 0 0 0 .973-1.23l-.973.23zM21 12.498l-.385.923A1 1 0 0 0 22 12.498h-1zM11.406 6.11l.832-.554l-.832.554zm-.812-1.218l-.832.554l.832-.554zM4 6a1 1 0 0 1 1-1V3a3 3 0 0 0-3 3h2zm0 12V6H2v12h2zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3v-2zm7.174 0H5v2h7.174v-2zm.973.77A5.522 5.522 0 0 1 13 18.5h-2c0 .594.07 1.174.2 1.73l1.947-.46zM13 18.5a5.5 5.5 0 0 1 5.5-5.5v-2a7.5 7.5 0 0 0-7.5 7.5h2zm5.5-5.5c.752 0 1.465.15 2.115.421l.77-1.846A7.48 7.48 0 0 0 18.5 11v2zM20 9v3.498h2V9h-2zm-1-1a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2zm-5.93 0H19V6h-5.93v2zm-2.496-1.336A3 3 0 0 0 13.07 8V6a1 1 0 0 1-.832-.445l-1.664 1.11zm-.812-1.219l.812 1.22l1.664-1.11l-.812-1.22l-1.664 1.11zM8.93 5a1 1 0 0 1 .832.445l1.664-1.11A3 3 0 0 0 8.93 3v2zM5 5h3.93V3H5v2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 19l2 2l4-5"/></g>`
	folderCheckLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 11.5V9a2 2 0 0 0-2-2h-5.93a2 2 0 0 1-1.664-.89l-.812-1.22A2 2 0 0 0 8.93 4H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7m4-1l2 2l4-5"/>`
	folderDownloadInnerSVG            = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h3.93a3 3 0 0 1 2.496 1.336l.812 1.219A1 1 0 0 0 13.07 6H19a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm11 5a1 1 0 1 0-2 0v3.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L13 14.586V11z" fill="currentColor"/></g>`
	folderDownloadLineInnerSVG        = `<g fill="currentColor"><path d="M4 6a1 1 0 0 1 1-1h3.93a1 1 0 0 1 .832.445l.812 1.22A3 3 0 0 0 13.07 8H19a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V6zm1-3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-5.93a1 1 0 0 1-.832-.445l-.812-1.22A3 3 0 0 0 8.93 3H5zm8 8a1 1 0 1 0-2 0v3.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L13 14.586V11z"/></g>`
	folderLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6a2 2 0 0 1 2-2h3.93a2 2 0 0 1 1.664.89l.812 1.22A2 2 0 0 0 13.07 7H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V6z"/>`
	folderMinusInnerSVG               = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M5 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6.498A6.5 6.5 0 0 1 21 12.022V9a2 2 0 0 0-2-2h-5.93a2 2 0 0 1-1.664-.89l-.812-1.22A2 2 0 0 0 8.93 4H5z" clip-rule="evenodd"/><path fill="currentColor" d="M11.498 20v1a1 1 0 0 0 .923-1.385l-.923.385zM21 12.022l-.54.842a1 1 0 0 0 1.54-.842h-1zm-9.594-5.913l.832-.554l-.832.554zm-.812-1.218l-.832.554l.832-.554zM4 6a1 1 0 0 1 1-1V3a3 3 0 0 0-3 3h2zm0 12V6H2v12h2zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3v-2zm6.498 0H5v2h6.498v-2zm.923.615A5.48 5.48 0 0 1 12 17.5h-2c0 1.02.204 1.995.575 2.885l1.846-.77zM12 17.5a5.5 5.5 0 0 1 5.5-5.5v-2a7.5 7.5 0 0 0-7.5 7.5h2zm5.5-5.5a5.47 5.47 0 0 1 2.96.864l1.08-1.684A7.47 7.47 0 0 0 17.5 10v2zM20 9v3.022h2V9h-2zm-1-1a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2zm-5.93 0H19V6h-5.93v2zm-2.496-1.336A3 3 0 0 0 13.07 8V6a1 1 0 0 1-.832-.445l-1.664 1.11zm-.812-1.219l.812 1.22l1.664-1.11l-.812-1.22l-1.664 1.11zM8.93 5a1 1 0 0 1 .832.445l1.664-1.11A3 3 0 0 0 8.93 3v2zM5 5h3.93V3H5v2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 18h6"/></g>`
	folderMinusLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13V9a2 2 0 0 0-2-2h-5.93a2 2 0 0 1-1.664-.89l-.812-1.22A2 2 0 0 0 8.93 4H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6m4-2h6"/>`
	folderOpenInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M21 10.17c1.165.412 2 1.524 2 2.83v5a3 3 0 0 1-3 3H6v-8a3 3 0 0 1 3-3h12v.17zM20.83 8A3.001 3.001 0 0 0 18 6h-5.93a1 1 0 0 1-.832-.445l-.812-1.22A3 3 0 0 0 7.93 3H5a3 3 0 0 0-3 3v12c0 .493.14 1.211.588 1.834A2.73 2.73 0 0 0 4 20.854V13a5 5 0 0 1 5-5h11.83z" fill="currentColor"/></g>`
	folderOpenLineInnerSVG            = `<g fill="currentColor"><path d="M2 6a3 3 0 0 1 3-3h2.93a3 3 0 0 1 2.496 1.336l.812 1.219A1 1 0 0 0 12.07 6H18a3 3 0 0 1 3 3v2h-2V9a1 1 0 0 0-1-1h-5.93a3 3 0 0 1-2.496-1.336l-.812-1.219A1 1 0 0 0 7.93 5H5a1 1 0 0 0-1 1v12a1 1 0 1 0 2 0v-2h2v2a3 3 0 1 1-6 0V6z"/><path d="M6 13a3 3 0 0 1 3-3h11a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3H5v-2c.173 0 .456-.06.666-.212c.159-.114.334-.314.334-.788v-5zm1.853 6H20a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v5c0 .368-.053.701-.147 1z"/></g>`
	folderPlusInnerSVG                = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M5 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6.29A7 7 0 0 1 21 11.674V9a2 2 0 0 0-2-2h-5.93a2 2 0 0 1-1.664-.89l-.812-1.22A2 2 0 0 0 8.93 4H5z" clip-rule="evenodd"/><path fill="currentColor" d="M11.29 20v1a1 1 0 0 0 .958-1.285L11.29 20zM21 11.674l-.43.903a1 1 0 0 0 1.43-.903h-1zm-9.594-5.565l.832-.554l-.832.554zm-.812-1.218l-.832.554l.832-.554zM4 6a1 1 0 0 1 1-1V3a3 3 0 0 0-3 3h2zm0 12V6H2v12h2zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3v-2zm6.29 0H5v2h6.29v-2zm.958.715A6.002 6.002 0 0 1 12 18h-2c0 .793.116 1.56.331 2.285l1.917-.57zM12 18a6 6 0 0 1 6-6v-2a8 8 0 0 0-8 8h2zm6-6a5.96 5.96 0 0 1 2.57.577l.86-1.807A7.973 7.973 0 0 0 18 10v2zm2-3v2.674h2V9h-2zm-1-1a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2zm-5.93 0H19V6h-5.93v2zm-2.496-1.336A3 3 0 0 0 13.07 8V6a1 1 0 0 1-.832-.445l-1.664 1.11zm-.812-1.219l.812 1.22l1.664-1.11l-.812-1.22l-1.664 1.11zM8.93 5a1 1 0 0 1 .832.445l1.664-1.11A3 3 0 0 0 8.93 3v2zM5 5h3.93V3H5v2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 15v3m0 3v-3m0 0h-3m3 0h3"/></g>`
	folderPlusLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13V9a2 2 0 0 0-2-2h-5.93a2 2 0 0 1-1.664-.89l-.812-1.22A2 2 0 0 0 8.93 4H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7m6-5v3m0 3v-3m0 0h-3m3 0h3"/>`
	folderRemoveInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h3.93a3 3 0 0 1 2.496 1.336l.812 1.219A1 1 0 0 0 13.07 6H19a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm7 7a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H9z" fill="currentColor"/></g>`
	folderRemoveLineInnerSVG          = `<g fill="currentColor"><path d="M5 5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-5.93a3 3 0 0 1-2.496-1.336l-.812-1.219A1 1 0 0 0 8.93 5H5zM2 6a3 3 0 0 1 3-3h3.93a3 3 0 0 1 2.496 1.336l.812 1.219A1 1 0 0 0 13.07 6H19a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6z"/><path d="M8 14a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1z"/></g>`
	fontSizeInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6h4m4 0h-4m0 0v12M4 11h3m3 0H7m0 0v7"/>`
	fontSizeLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6h4m4 0h-4m0 0v12M4 11h3m3 0H7m0 0v7"/>`
	forwardInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 5a1 1 0 0 1 1.707-.707l7 7a1 1 0 0 1 0 1.414l-7 7A1 1 0 0 1 12 19v-3H4a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h8V5z" fill="currentColor"/></g>`
	forwardCircleInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1zM6.507 8.13a1 1 0 0 1 1.008.013L12 10.833V9a1 1 0 0 1 1.514-.857l5 3a1 1 0 0 1 0 1.714l-5 3A1 1 0 0 1 12 15v-1.834l-4.486 2.691A1 1 0 0 1 6 15V9a1 1 0 0 1 .507-.87zM8 13.234L10.056 12L8 10.766v2.468zM16.056 12L14 13.234v-2.468L16.056 12z" clip-rule="evenodd"/>`
	forwardCircleLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M7 15V9l5 3l-5 3zm6 0V9l5 3l-5 3z"/></g>`
	forwardEndCircleInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1zM6.507 8.13a1 1 0 0 1 1.008.013L11 10.233V9a1 1 0 0 1 1.514-.857L16 10.233V9a1 1 0 1 1 2 0v6a1 1 0 1 1-2 0v-1.234l-3.486 2.091A1 1 0 0 1 11 15v-1.234l-3.486 2.091A1 1 0 0 1 6 15V9a1 1 0 0 1 .507-.87zM13 12v1.234L15.056 12L13 10.766V12zm-5 1.234L10.056 12L8 10.766v2.468z" clip-rule="evenodd"/>`
	forwardEndCircleLineInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M7 15V9l5 3l-5 3zm10-3l-5 3V9l5 3zm0 0V9m0 3v3"/></g>`
	forwardLineInnerSVG               = `<g fill="currentColor"><path d="M12.617 4.076A1 1 0 0 0 12 5v3H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h8v3a1 1 0 0 0 1.707.707l7-7a1 1 0 0 0 0-1.414l-7-7a1 1 0 0 0-1.09-.217zM18.586 12L14 16.586V15a1 1 0 0 0-1-1H5v-4h8a1 1 0 0 0 1-1V7.414L18.586 12z"/></g>`
	ghostInnerSVG                     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6.416 3.788C8.289 2.44 10.506 2 12 2c3.526 0 5.826 1.492 7.212 3.416C20.56 7.289 21 9.506 21 11v9a1 1 0 0 1-1.707.707L18 19.414L16.414 21a2 2 0 0 1-2.828 0L12 19.414L10.414 21a2 2 0 0 1-2.828 0L6 19.414l-1.293 1.293A1 1 0 0 1 3 20v-9c0-3.526 1.492-5.826 3.416-7.212zM7 10a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm6 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0z" fill="currentColor"/></g>`
	ghostLineInnerSVG                 = `<g fill="currentColor"><path d="M6.416 3.788C8.289 2.44 10.506 2 12 2c3.526 0 5.826 1.492 7.212 3.416C20.56 7.289 21 9.506 21 11v9a1 1 0 0 1-1.707.707L18 19.414L16.414 21a2 2 0 0 1-2.828 0L12 19.414L10.414 21a2 2 0 0 1-2.828 0L6 19.414l-1.293 1.293A1 1 0 0 1 3 20v-9c0-3.526 1.492-5.826 3.416-7.212zm1.168 1.624C6.175 6.426 5 8.126 5 11v6.682A2 2 0 0 1 7.414 18L9 19.586L10.586 18a2 2 0 0 1 2.828 0L15 19.586L16.586 18A2 2 0 0 1 19 17.682V11c0-1.173-.36-2.956-1.412-4.416C16.575 5.175 14.874 4 12 4c-1.173 0-2.956.36-4.416 1.412zM7 10a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm8-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4z"/></g>`
	gitBranchInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 21v-6m0-6v6m0 0h8a2 2 0 0 0 2-2v-2"/><circle cx="7" cy="6" r="3" fill="currentColor"/><circle cx="17" cy="8" r="3" fill="currentColor"/></g>`
	gitBranchLineInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 21v-6m0-6v6m0 0h8a2 2 0 0 0 2-2v-2"/><circle cx="7" cy="6" r="3"/><circle cx="17" cy="8" r="3"/></g>`
	gitCommitInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M12 16a4 4 0 1 1 0-8a4 4 0 0 1 0 8z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16a4 4 0 0 1 0-8m0 8a4 4 0 0 0 0-8m0 8v5m0-13V3"/></g>`
	gitCommitLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16a4 4 0 0 1 0-8m0 8a4 4 0 0 0 0-8m0 8v5m0-13V3"/>`
	gitCompareInnerSVG                = `<g fill="none"><path fill="currentColor" d="M21 17a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 14a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm0 0V9a2 2 0 0 0-2-2h-1m-2 0l2-2v2m-2 0h2m-2 0l2 2V7"/><path fill="currentColor" d="M3 8a3 3 0 1 1 6 0a3 3 0 0 1-6 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm0 0v5a2 2 0 0 0 2 2h1m2 0l-2 2v-2m2 0H9m2 0l-2-2v2"/></g>`
	gitCompareLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 14a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm0 0V9a2 2 0 0 0-2-2h-1m-2 0l2-2v2m-2 0h2m-2 0l2 2V7m-9 4a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm0 0v5a2 2 0 0 0 2 2h1m2 0l-2 2v-2m2 0H9m2 0l-2-2v2"/>`
	gitForkInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="6" cy="6" r="3" fill="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18" cy="6" r="3" fill="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12" cy="18" r="3" fill="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path d="M6 9v1a2 2 0 0 0 2 2h4m6-3v1a2 2 0 0 1-2 2h-4m0 0v3"/></g>`
	gitForkLineInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="6" cy="6" r="3" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18" cy="6" r="3" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12" cy="18" r="3" stroke-linecap="round" stroke-linejoin="round"/><path d="M6 9v1a2 2 0 0 0 2 2h4m6-3v1a2 2 0 0 1-2 2h-4m0 0v3"/></g>`
	gitMergeInnerSVG                  = `<g fill="none"><path fill="currentColor" d="M10 7a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm0 0v10m0-10c0 2.333 1.4 7 7 7"/><circle cx="17" cy="17" r="3" fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g>`
	gitMergeLineInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm0 0v10m0-10c0 2.333 1.4 7 7 7"/><circle cx="17" cy="17" r="3"/></g>`
	gitPullInnerSVG                   = `<g fill="none"><path fill="currentColor" d="M9 7a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm0 0v10"/><path fill="currentColor" d="M21 17a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 14a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm0 0V9a2 2 0 0 0-2-2h-1m-2 0l2-2v2m-2 0h2m-2 0l2 2V7"/></g>`
	gitPullLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm0 0v10m12-6a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm0 0V9a2 2 0 0 0-2-2h-1m-2 0l2-2v2m-2 0h2m-2 0l2 2V7"/>`
	glasWaterInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M15.234 20H8.766a2 2 0 0 1-1.985-1.752L6 12l2.072-.69a2 2 0 0 1 1.742.233l1.077.717a2 2 0 0 0 2.218 0l1.077-.717a2 2 0 0 1 1.742-.234L18 12l-.781 6.248A2 2 0 0 1 15.234 20z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.5 4h15M6 12l.781 6.248A2 2 0 0 0 8.766 20h6.468a2 2 0 0 0 1.985-1.752L18 12M6 12l2.072-.69a2 2 0 0 1 1.742.233l1.077.717a2 2 0 0 0 2.218 0l1.077-.717a2 2 0 0 1 1.742-.234L18 12M6 12l-.938-7.5M18 12l.938-7.5"/></g>`
	glasWaterLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h14M5 4h-.5M5 4l1 8m13-8h.5M19 4l-1 8M6 12l.781 6.248A2 2 0 0 0 8.766 20h6.468a2 2 0 0 0 1.985-1.752L18 12M6 12l2.072-.69a2 2 0 0 1 1.742.233l1.077.717a2 2 0 0 0 2.218 0l1.077-.717a2 2 0 0 1 1.742-.234L18 12"/>`
	globeInnerSVG                     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10.059 2.188a10.016 10.016 0 0 0-7.242 5.847h4.616a16.98 16.98 0 0 1 2.19-5.162l.436-.685zM2.202 9.99a10.04 10.04 0 0 0 0 4.021c.094-.029.194-.044.298-.044h4.579a17.026 17.026 0 0 1 0-3.931H2.5a.978.978 0 0 1-.298-.046zm.615 5.976a10.017 10.017 0 0 0 7.242 5.847l-.436-.685a16.98 16.98 0 0 1-2.19-5.161H2.817zm11.124 5.847a10.017 10.017 0 0 0 7.242-5.847h-4.616a16.98 16.98 0 0 1-2.19 5.162l-.436.685zm7.857-7.802a10.037 10.037 0 0 0 0-4.02a1.009 1.009 0 0 1-.298.044h-4.579a17.023 17.023 0 0 1 0 3.931H21.5c.104 0 .204.016.298.046zm-.615-5.975h-4.616a16.98 16.98 0 0 0-2.19-5.162l-.436-.685a10.017 10.017 0 0 1 7.242 5.847zM12.69 3.947a14.987 14.987 0 0 1 1.812 4.088H9.498c.39-1.424.994-2.803 1.812-4.088L12 2.863l.69 1.084zM9.094 13.966a15.023 15.023 0 0 1 0-3.931h5.812a15.02 15.02 0 0 1 0 3.93H9.094zm2.216 6.087a14.986 14.986 0 0 1-1.812-4.088h5.004a14.986 14.986 0 0 1-1.812 4.088L12 21.137l-.69-1.084z" fill="currentColor"/></g>`
	globeEarthInnerSVG                = `<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" d="M8 6V4c2.5-1.167 8.4-2 12 4h-1a2 2 0 0 0-2 2a2 2 0 1 1-4 0a2 2 0 0 0-2-2h-1a2 2 0 0 1-2-2zm9 10h3c-1.2 1.6-3.833 3.333-5 4v-2a2 2 0 0 1 2-2zm-6 2v2c-6.4-.4-8-5.5-8-8h2a2 2 0 0 1 2 2a2 2 0 0 0 2 2a2 2 0 0 1 2 2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6V4c2.5-1.167 8.4-2 12 4h-1a2 2 0 0 0-2 2a2 2 0 1 1-4 0a2 2 0 0 0-2-2h-1a2 2 0 0 1-2-2zm9 10h3c-1.2 1.6-3.833 3.333-5 4v-2a2 2 0 0 1 2-2zm-6 2v2c-6.4-.4-8-5.5-8-8h2a2 2 0 0 1 2 2a2 2 0 0 0 2 2a2 2 0 0 1 2 2z"/></g>`
	globeEarthLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M8 4v2a2 2 0 0 0 2 2h1a2 2 0 0 1 2 2v0a2 2 0 0 0 2 2v0a2 2 0 0 0 2-2v0a2 2 0 0 1 2-2h1m0 8h-3a2 2 0 0 0-2 2v2m-4 0v-2a2 2 0 0 0-2-2v0a2 2 0 0 1-2-2v0a2 2 0 0 0-2-2H3"/></g>`
	globeEarthTwoInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path fill="currentColor" d="M13.5 6h.5a2 2 0 0 0 2-2C8.4 1.2 4.5 6.5 3.5 9.5l4.167 5.093c.215.263.333.593.333.933C8 16.34 8.673 17 9.487 17c.828 0 1.513.672 1.513 1.5V21c2.167-.167 6.8-1.2 8-4h-1a2 2 0 0 1-2-2a2 2 0 0 0-2-2h-2a2 2 0 0 1-2-2v-1a1 1 0 0 1 1-1a1 1 0 0 0 1-1v-.5A1.5 1.5 0 0 1 13.5 6z"/></g>`
	globeEarthTwoLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M16 4v0a2 2 0 0 1-2 2h-.5A1.5 1.5 0 0 0 12 7.5V8a1 1 0 0 1-1 1v0a1 1 0 0 0-1 1v1a2 2 0 0 0 2 2h2a2 2 0 0 1 2 2v0a2 2 0 0 0 2 2h1m-8 4v-2.5c0-.828-.685-1.5-1.513-1.5v0C8.673 17 8 16.34 8 15.526v0c0-.34-.118-.67-.333-.933L3.5 9.5"/></g>`
	globeGridInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.5 9.035a9.004 9.004 0 0 0-17 0m17 0c.324.928.5 1.926.5 2.965a8.988 8.988 0 0 1-.5 2.966m0-5.931h-17m0 0A8.987 8.987 0 0 0 3 12a8.99 8.99 0 0 0 .5 2.966m0 0a9.004 9.004 0 0 0 17 0m-17 0h17"/><path d="M12 21c4.97-4.97 4.97-13.03 0-18c-4.97 4.97-4.97 13.03 0 18z"/></g>`
	globeGridLineInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.5 9.035a9.004 9.004 0 0 0-17 0m17 0c.324.928.5 1.926.5 2.965a8.988 8.988 0 0 1-.5 2.966m0-5.931h-17m0 0A8.987 8.987 0 0 0 3 12a8.99 8.99 0 0 0 .5 2.966m0 0a9.004 9.004 0 0 0 17 0m-17 0h17"/><path d="M12 21c4.97-4.97 4.97-13.03 0-18c-4.97 4.97-4.97 13.03 0 18z"/></g>`
	globeLineInnerSVG                 = `<g fill="currentColor"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12z"/><path d="M11.293 2.293c-5.361 5.36-5.361 14.053 0 19.414h1.414c5.361-5.361 5.361-14.053 0-19.414h-1.414zM12 4.479c3.637 4.343 3.637 10.7 0 15.042c-3.637-4.343-3.637-10.7 0-15.042z"/><path d="M3 9a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1zm0 6a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1z"/></g>`
	handInnerSVG                      = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 22C4.7 22 3 16.333 3 13.5V10c0-.167.1-.5.5-.5s.5.333.5.5v2c1.215.722 1.877.699 3 0V4c0-.167.1-.5.5-.5s.5.333.5.5v8c1.93-.882 2.73-1.484 3-3V3c0-.167.1-.5.5-.5s.5.333.5.5v6c.622 1.739 1.26 2.168 3 2V4c0-.167.1-.5.5-.5s.5.333.5.5v7c.728.963 1.455 1.09 3 1V9c0-.167.1-.5.5-.5s.5.333.5.5v4.5c0 2.833-1.7 8.5-8.5 8.5z"/>`
	handLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M13 6V4.5A1.5 1.5 0 0 0 11.5 3v0A1.5 1.5 0 0 0 10 4.5V12m3-1V5.5A1.5 1.5 0 0 1 14.5 4v0A1.5 1.5 0 0 1 16 5.5V10m-6-3V5.5A1.5 1.5 0 0 0 8.5 4v0A1.5 1.5 0 0 0 7 5.5v6M7 14v-2.5A1.5 1.5 0 0 0 5.5 10v0A1.5 1.5 0 0 0 4 11.5v2C4 16 5.5 21 11.5 21c2.5 0 7.5-1.5 7.5-7.5v-4A1.5 1.5 0 0 0 17.5 8v0A1.5 1.5 0 0 0 16 9.5V12"/>`
	handPointerInnerSVG               = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 17v1a2 2 0 0 1-2 2H8.236a2 2 0 0 1-1.789-1.106l-2.276-4.552A1.618 1.618 0 0 1 5.618 12H6a2 2 0 0 1 1.6.8L10 16V6a2 2 0 1 1 4 0v6a1 1 0 0 0 1 1h1a4 4 0 0 1 4 4z"/>`
	handPointerEventInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M15 15h1a4 4 0 0 1 4 4a2 2 0 0 1-2 2H8.132a2 2 0 0 1-1.715-.971l-2.194-3.657A1.566 1.566 0 0 1 5.566 14H6a2 2 0 0 1 1.6.8L10 18V9a2 2 0 1 1 4 0v5a1 1 0 0 0 1 1z"/><path d="M12 4V3m6 7h1M5 10h1m1.343-4.657l-.707-.707m10.021.707l.707-.707"/></g>`
	handPointerEventLineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 21v-2a4 4 0 0 0-4-4h-1a1 1 0 0 1-1-1V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v9l-2.4-3.2A2 2 0 0 0 6 14h-.434C4.701 14 4 14.701 4 15.566v0c0 .284.077.563.223.806L7 21m5-17V3m6 7h1M5 10h1m1.343-4.657l-.707-.707m10.021.707l.707-.707"/>`
	handPointerLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20v-3a4 4 0 0 0-4-4h-1a1 1 0 0 1-1-1V6a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v10l-2.4-3.2A2 2 0 0 0 6 12h-.382C4.724 12 4 12.724 4 13.618v0c0 .251.058.499.17.724L7 20"/>`
	handPointerTwoInnerSVG            = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 22C4.7 22 3 16.333 3 13.5V10c0-.333.2-1 1-1s1 .667 1 1v2c0 .5.3 1.5 1.5 1.5S8 12.5 8 12V3c0-.333.2-1 1-1s1 .667 1 1v7c.5.5.8 1.2 2 0V8c0-.333.2-1 1-1s1 .667 1 1v1h1c0-.333.2-1 1-1s1 .667 1 1v1h1c0-.333.2-1 1-1s1 .667 1 1v3.5c0 2.833-1.7 8.5-8.5 8.5z"/>`
	handPointerTwoLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M13 10V8.5A1.5 1.5 0 0 0 11.5 7v0A1.5 1.5 0 0 0 10 8.5V11m3 0V9.5A1.5 1.5 0 0 1 14.5 8v0A1.5 1.5 0 0 1 16 9.5v.5m-9 1.5v-7A1.5 1.5 0 0 1 8.5 3v0A1.5 1.5 0 0 1 10 4.5V11m0 1v-1m-3 3v-2.5A1.5 1.5 0 0 0 5.5 10v0A1.5 1.5 0 0 0 4 11.5v2C4 16 5.5 21 11.5 21c2.5 0 7.5-1.5 7.5-7.5V12a2 2 0 0 0-2-2h-1m0 0v2"/>`
	hardDriveInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M4 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v9H4V5zm0 11v3a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-3H4zm9 2a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-3z" clip-rule="evenodd"/>`
	hardDriveLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15h14M5 15v4a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-4M5 15V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10m-6 3h3"/>`
	hashtagInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5zm7.316 4.051a1 1 0 0 1 .633 1.265L12.054 9h1.892l1.105-3.316a1 1 0 0 1 1.898.632L16.054 9H19a1 1 0 1 1 0 2h-3.613l-.666 2H17a1 1 0 1 1 0 2h-2.946l-1.105 3.316a1 1 0 0 1-1.898-.632L11.946 15h-1.892L8.95 18.316a1 1 0 1 1-1.898-.632L7.946 15H5a1 1 0 1 1 0-2h3.613l.666-2H7a1 1 0 1 1 0-2h2.946l1.105-3.316a1 1 0 0 1 1.265-.633zM10.721 13l.666-2h1.892l-.666 2H10.72z" fill="currentColor"/></g>`
	hashtagLineInnerSVG               = `<g fill="currentColor"><path d="M11.242 3.03a1 1 0 0 1 .728 1.213l-4 16a1 1 0 1 1-1.94-.485l4-16a1 1 0 0 1 1.213-.728zm6 0a1 1 0 0 1 .728 1.213l-4 16a1 1 0 1 1-1.94-.485l4-16a1 1 0 0 1 1.212-.728z"/><path d="M4 9a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm-2 6a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1z"/></g>`
	headsetInnerSVG                   = `<g fill="none"><path fill="currentColor" d="M4 18v-7c0-2.333 2.4-7 8-7s8 4.667 8 7v7a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h3v-3c0-2.333-2.4-7-8-7s-8 4.667-8 7v3h3a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14v4a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2H4zm0 0v-3c0-2.333 2.4-7 8-7s8 4.667 8 7v3m0 0v4a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h3z"/></g>`
	headsetLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14v4a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2H4zm0 0v-3c0-2.333 2.4-7 8-7s8 4.667 8 7v3m0 0v4a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h3z"/>`
	heartInnerSVG                     = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 4c-3.2 0-5 2.667-5 4c0-1.333-1.8-4-5-4S3 6.667 3 8c0 7 9 12 9 12s9-5 9-12c0-1.333-.8-4-4-4z"/>`
	heartLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 4c-3.2 0-5 2.667-5 4c0-1.333-1.8-4-5-4S3 6.667 3 8c0 7 9 12 9 12s9-5 9-12c0-1.333-.8-4-4-4z"/>`
	homeInnerSVG                      = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 19v-8.5a1 1 0 0 0-.4-.8l-7-5.25a1 1 0 0 0-1.2 0l-7 5.25a1 1 0 0 0-.4.8V19a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v3a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1z"/>`
	homeAnalyticsInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M10.8 3.65a2 2 0 0 1 2.4 0l7 5.25l-.6.8l.6-.8a2 2 0 0 1 .8 1.6V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8.5a2 2 0 0 1 .8-1.6l7-5.25zM13 10a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0v-6zm-4 3a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0v-3zm8 2a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0v-1z" clip-rule="evenodd"/>`
	homeAnalyticsLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 19v-8.5a1 1 0 0 0-.4-.8l-7-5.25a1 1 0 0 0-1.2 0l-7 5.25a1 1 0 0 0-.4.8V19a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1zM8 13v3m4-6v6m4-1v1"/>`
	homeLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 19v-8.5a1 1 0 0 0-.4-.8l-7-5.25a1 1 0 0 0-1.2 0l-7 5.25a1 1 0 0 0-.4.8V19a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v3a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1z"/>`
	homeSimpleInnerSVG                = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 19v-8.5a1 1 0 0 0-.4-.8l-7-5.25a1 1 0 0 0-1.2 0l-7 5.25a1 1 0 0 0-.4.8V19a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1z"/>`
	homeSimpleLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 19v-8.5a1 1 0 0 0-.4-.8l-7-5.25a1 1 0 0 0-1.2 0l-7 5.25a1 1 0 0 0-.4.8V19a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1z"/>`
	imageInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M2 5a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v6.5a1 1 0 0 1-.032.25A1 1 0 0 1 22 12v7a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-3a1 1 0 0 1 .032-.25A1.002 1.002 0 0 1 2 15.5V5zm2.994 9.83c-.348.006-.68.022-.994.046V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v6.016c-4.297.139-7.4 1.174-9.58 2.623c.826.293 1.75.71 2.656 1.256c1.399.84 2.821 2.02 3.778 3.583a1 1 0 1 1-1.706 1.044c-.736-1.203-1.878-2.178-3.102-2.913c-1.222-.734-2.465-1.192-3.327-1.392a15.466 15.466 0 0 0-3.703-.386h-.022zm1.984-8.342A2.674 2.674 0 0 1 8.5 6c.41 0 1.003.115 1.522.488c.57.41.978 1.086.978 2.012c0 .926-.408 1.601-.978 2.011A2.674 2.674 0 0 1 8.5 11c-.41 0-1.003-.115-1.522-.489C6.408 10.101 6 9.427 6 8.5c0-.926.408-1.601.978-2.012z" clip-rule="evenodd"/>`
	imageCircleInnerSVG               = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M6.978 6.488A2.674 2.674 0 0 1 8.5 6c.41 0 1.003.115 1.522.488c.57.41.978 1.086.978 2.012c0 .926-.408 1.601-.978 2.011A2.674 2.674 0 0 1 8.5 11c-.41 0-1.003-.115-1.522-.489C6.408 10.101 6 9.427 6 8.5c0-.926.408-1.601.978-2.012zm9.353 15.456C18.611 21.177 23 18.143 23 12a1 1 0 0 0-1-1h-1c-4.803 0-8.21 1.072-10.555 2.622c2.035.662 4.076 1.82 5.63 3.751a1 1 0 0 1-1.56 1.254c-1.515-1.884-3.65-2.912-5.796-3.41a15.464 15.464 0 0 0-3.531-.388c-.784.003-1.477.066-2.024.157a1.005 1.005 0 0 1-.232.012l-.096.016a1 1 0 0 0-.76 1.367c.652 1.584 2.135 3.723 4.51 5.097c2.42 1.399 5.684 1.958 9.745.466z" clip-rule="evenodd"/><circle cx="12" cy="12" r="10" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g>`
	imageCircleLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 16c1.403-.234 3.637-.293 5.945.243M16 21c-1.704-2.768-4.427-4.148-7.055-4.757m0 0C10.895 13.985 14.558 12 21 12h1M8.5 7C8 7 7 7.3 7 8.5S8 10 8.5 10S10 9.7 10 8.5S9 7 8.5 7z"/><circle cx="12" cy="12" r="10"/></g>`
	imageCircleOffInnerSVG            = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.859 5A9.968 9.968 0 0 0 2 12c0 5.523 4.477 10 10 10a9.967 9.967 0 0 0 7-2.859M21.542 15A9.997 9.997 0 0 0 22 12c0-5.523-4.477-10-10-10a9.966 9.966 0 0 0-4 .832M2 2l20 20"/><path fill="currentColor" fill-rule="evenodd" d="M6.346 7.174C6.13 7.534 6 7.977 6 8.5c0 .926.408 1.601.978 2.011A2.674 2.674 0 0 0 8.5 11c.357 0 .852-.087 1.316-.355l-3.47-3.47zm5.598 5.599c-.542.261-1.04.546-1.5.849c2.036.662 4.077 1.82 5.63 3.751a1 1 0 0 1-1.558 1.254c-1.516-1.884-3.65-2.912-5.797-3.41a15.464 15.464 0 0 0-3.531-.388c-.784.003-1.477.066-2.024.157a1.005 1.005 0 0 1-.232.012l-.096.016a1 1 0 0 0-.76 1.367c.652 1.584 2.135 3.723 4.51 5.097c2.42 1.399 5.684 1.958 9.745.465c.829-.278 1.936-.856 3.01-1.774l-7.397-7.396zm9.898 4.24C22.541 15.65 23 13.988 23 12a1 1 0 0 0-1-1h-1c-1.772 0-3.355.146-4.763.408l5.605 5.606z" clip-rule="evenodd"/></g>`
	imageCircleOffLineInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 16c1.403-.234 3.637-.293 5.945.243M16 21c-1.704-2.768-4.427-4.148-7.055-4.757m0 0c.927-1.073 2.24-2.084 4.055-2.85M22 12h-1c-1.475 0-2.804.104-4 .291M7.341 7.5C7.14 7.728 7 8.051 7 8.5C7 9.7 8 10 8.5 10a1.66 1.66 0 0 0 1-.348"/><path d="M4.859 5A9.968 9.968 0 0 0 2 12c0 5.523 4.477 10 10 10a9.967 9.967 0 0 0 7-2.859M8 2.832A9.966 9.966 0 0 1 12 2c5.523 0 10 4.477 10 10a9.997 9.997 0 0 1-.832 4M2 2l20 20"/></g>`
	imageCirclePlusInnerSVG           = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2m7 0v3m0 3V5m0 0h3m-3 0h-3"/><path fill="currentColor" fill-rule="evenodd" d="M6.978 6.488A2.674 2.674 0 0 1 8.5 6c.41 0 1.003.115 1.522.488c.57.41.978 1.086.978 2.012c0 .926-.408 1.601-.978 2.011A2.674 2.674 0 0 1 8.5 11c-.41 0-1.003-.115-1.522-.489C6.408 10.101 6 9.427 6 8.5c0-.926.408-1.601.978-2.012zm9.353 15.456C18.611 21.177 23 18.143 23 12a1 1 0 0 0-1-1h-1c-4.803 0-8.21 1.072-10.555 2.622c2.035.662 4.076 1.82 5.63 3.751a1 1 0 0 1-1.56 1.254c-1.515-1.884-3.65-2.912-5.796-3.41a15.464 15.464 0 0 0-3.531-.388c-.784.003-1.477.066-2.024.157a1.005 1.005 0 0 1-.232.012l-.096.016a1 1 0 0 0-.76 1.367c.652 1.584 2.135 3.723 4.51 5.097c2.42 1.399 5.684 1.958 9.745.466z" clip-rule="evenodd"/></g>`
	imageCirclePlusLineInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 16c1.403-.234 3.637-.293 5.945.243M16 21c-1.704-2.768-4.427-4.148-7.055-4.757m0 0C10.895 13.985 14.558 12 21 12h1M8.5 7C8 7 7 7.3 7 8.5S8 10 8.5 10S10 9.7 10 8.5S9 7 8.5 7z"/><path d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2m7 0v3m0 3V5m0 0h3m-3 0h-3"/></g>`
	imageCircleStoryInnerSVG          = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12a9.966 9.966 0 0 1-.832 4M12 22a9.966 9.966 0 0 0 7.071-2.929M2 12a9.966 9.966 0 0 1 2.929-7.071M12 2a9.966 9.966 0 0 0-4 .832m0 18.336A10.02 10.02 0 0 1 2.832 16m13-13A10.02 10.02 0 0 1 21 8.168"/><path fill="currentColor" fill-rule="evenodd" d="M7.545 12.095a4.55 4.55 0 0 1 9.068-.54c-2.271.051-3.942.557-5.143 1.284c1.085.378 2.17 1.018 3.007 2.059a.75.75 0 0 1-1.169.94c-.79-.981-1.907-1.524-3.049-1.789a8.265 8.265 0 0 0-2.359-.19a4.535 4.535 0 0 1-.355-1.764zm4.55-6.36a6.36 6.36 0 1 0 0 12.72a6.36 6.36 0 0 0 0-12.72zm-1.909 3.087c-.224 0-.547.063-.83.267c-.31.224-.534.592-.534 1.097s.223.873.534 1.097c.283.204.606.267.83.267c.224 0 .547-.063.83-.267c.311-.224.534-.592.534-1.097s-.223-.873-.534-1.097a1.459 1.459 0 0 0-.83-.267z" clip-rule="evenodd"/></g>`
	imageCircleStoryLineInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11.23 13.92c1.598.614 2.56 1.587 3.77 3.08m-3.77-3.08c-1.943-.747-3.405-.942-4.73-.92m4.73.92c1.515-1.666 2.51-2.867 6.27-3.42M10 10h.001"/><circle cx="12" cy="12" r="6"/><path d="M22 12a9.966 9.966 0 0 1-.832 4M12 22a9.966 9.966 0 0 0 7.071-2.929M2 12a9.966 9.966 0 0 1 2.929-7.071M12 2a9.966 9.966 0 0 0-4 .832m0 18.336A10.02 10.02 0 0 1 2.832 16m13-13A10.02 10.02 0 0 1 21 8.168"/></g>`
	imageFrameInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M12.707 4.121a1 1 0 0 0-1.414 0L10.414 5h3.172l-.879-.879zm3.66.879a1.003 1.003 0 0 0-.16-.207l-2.086-2.086a3 3 0 0 0-4.242 0L7.793 4.793a1.001 1.001 0 0 0-.16.207H5a3 3 0 0 0-3 3v6.5c0 .086.01.17.031.25c-.02.08-.031.164-.031.25v3a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-5a1 1 0 0 0-.032-.25a.985.985 0 0 0 .032-.25V8a3 3 0 0 0-3-3h-2.634zM20 11.986V8a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v6.016A1 1 0 0 0 5 15c2.689 0 4.861.353 6.585.87c2.34.702 3.853 1.708 4.717 2.551a1 1 0 1 0 1.396-1.431c-1.12-1.094-2.925-2.252-5.539-3.036a19.62 19.62 0 0 0-.773-.215c1.176-.573 2.472-.967 3.737-1.235A25.149 25.149 0 0 1 20 11.986zM5.978 8.488A2.674 2.674 0 0 1 7.5 8c.41 0 1.003.115 1.522.488c.57.41.978 1.086.978 2.012c0 .926-.408 1.601-.978 2.011A2.674 2.674 0 0 1 7.5 13c-.41 0-1.003-.115-1.522-.489C5.408 12.101 5 11.427 5 10.5c0-.926.408-1.601.978-2.012z" clip-rule="evenodd"/>`
	imageFrameLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 20h3a2 2 0 0 0 2-2v-5m-5 7H5a2 2 0 0 1-2-2v-3m13 5c-.558-1.116-2.43-2.979-6.128-4.088M21 13V8a2 2 0 0 0-2-2h-3m5 7c-3.076-.118-8.335.435-11.128 2.912M3 15V8a2 2 0 0 1 2-2h3m-5 9c2.776 0 5.047.364 6.872.912M8 6l2.586-2.586a2 2 0 0 1 2.828 0L16 6M8 6h8M7.5 9C7 9 6 9.3 6 10.5S7 12 7.5 12S9 11.7 9 10.5S8 9 7.5 9z"/>`
	imageInPictureInnerSVG            = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h.001M20 8h.001M20 12h.001M20 16h.001M20 20h.001M16 20h.001M4 4h.001M8 4h.001M12 4h.001M16 4h.001M20 4h.001"/><path fill="currentColor" fill-rule="evenodd" d="M3 13a3 3 0 0 1 3-3h5a3 3 0 0 1 3 3v1h-1c-1.14 0-2.624.28-3.85 1.2C7.876 16.156 7 17.715 7 20v1H6a3 3 0 0 1-3-3v-5zm6 8h2a3 3 0 0 0 3-3v-2h-1c-.86 0-1.876.22-2.65.8C9.624 17.344 9 18.285 9 20v1zm-3-9a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H6z" clip-rule="evenodd"/></g>`
	imageInPictureLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 15v-2a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h2m5-5v3a2 2 0 0 1-2 2H8m5-5c-2 0-5 1-5 5m-1-6h.01M4 8h.001M20 8h.001M20 12h.001M20 16h.001M20 20h.001M16 20h.001M4 4h.001M8 4h.001M12 4h.001M16 4h.001M20 4h.001"/>`
	imageLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v11m18-4v7a2 2 0 0 1-2 2h-3m5-9c-6.442 0-10.105 1.985-12.055 4.243M3 16v3a2 2 0 0 0 2 2v0h11M3 16c1.403-.234 3.637-.293 5.945.243M16 21c-1.704-2.768-4.427-4.148-7.055-4.757M8.5 7C8 7 7 7.3 7 8.5S8 10 8.5 10S10 9.7 10 8.5S9 7 8.5 7z"/>`
	imageMultipleInnerSVG             = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 6v12a4 4 0 0 1-4 4H6"/><path fill="currentColor" fill-rule="evenodd" d="M1 3.7A2.7 2.7 0 0 1 3.7 1h12.6A2.7 2.7 0 0 1 19 3.7v5.85c0 .078-.01.153-.028.225A.897.897 0 0 1 19 10v6.3a2.7 2.7 0 0 1-2.7 2.7H3.7A2.7 2.7 0 0 1 1 16.3v-2.7a.9.9 0 0 1 .028-.225A.902.902 0 0 1 1 13.15V3.7zm2.695 8.848a13.81 13.81 0 0 0-.895.04V3.7a.9.9 0 0 1 .9-.9h12.6a.9.9 0 0 1 .9.9v5.414c-3.868.125-6.66 1.057-8.623 2.36c.745.265 1.575.64 2.391 1.131c1.26.756 2.54 1.819 3.4 3.225a.9.9 0 1 1-1.535.94c-.663-1.083-1.69-1.96-2.792-2.622c-1.1-.66-2.218-1.073-2.994-1.253a13.924 13.924 0 0 0-3.333-.348h-.02zM5.48 5.04a2.403 2.403 0 0 1 1.37-.44c.369 0 .903.103 1.37.44c.513.369.88.977.88 1.81c0 .833-.367 1.441-.88 1.81a2.403 2.403 0 0 1-1.37.44c-.369 0-.903-.103-1.37-.44c-.513-.369-.88-.977-.88-1.81c0-.833.367-1.441.88-1.81z" clip-rule="evenodd"/></g>`
	imageMultipleLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 9V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v9m16-4v7a2 2 0 0 1-2 2h-2m4-9c-6.442 0-9.105 1.985-11.055 4.243M2 13v4a1 1 0 0 0 1 1v0h11M2 13c1.403-.234 2.637-.293 4.945.243M14 18c-1.704-2.768-4.427-4.148-7.055-4.757M6.5 5C6 5 5 5.3 5 6.5S6 8 6.5 8S8 7.7 8 6.5S7 5 6.5 5z"/><path d="M22 6v12a4 4 0 0 1-4 4H6"/></g>`
	imageOffInnerSVG                  = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20"/><path fill="currentColor" fill-rule="evenodd" d="M2.505 3.333A2.986 2.986 0 0 0 2 5v10.5c0 .086.011.17.032.25A1 1 0 0 0 2 16v3a3 3 0 0 0 3 3h14c.617 0 1.19-.186 1.666-.505l-8.722-8.723c-.552.267-1.06.558-1.525.867c.827.293 1.75.71 2.657 1.256c1.399.84 2.821 2.02 3.778 3.583a1 1 0 1 1-1.706 1.044c-.736-1.203-1.878-2.178-3.102-2.913c-1.222-.734-2.465-1.192-3.327-1.392a15.466 15.466 0 0 0-3.703-.386h-.022c-.348.005-.68.02-.994.045V5a.93.93 0 0 1 .013-.159L2.505 3.333zm13.732 8.075A25.003 25.003 0 0 1 20 11.016V5a1 1 0 0 0-1-1H8.828l-2-2H19a3 3 0 0 1 3 3v6.5a1 1 0 0 1-.032.25A1 1 0 0 1 22 12v5.172l-5.763-5.764zM6.346 7.174C6.13 7.534 6 7.977 6 8.5c0 .926.408 1.601.978 2.011A2.674 2.674 0 0 0 8.5 11c.357 0 .852-.087 1.316-.355l-3.47-3.47z" clip-rule="evenodd"/></g>`
	imageOffLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 16v-4m0 0V5a2 2 0 0 0-2-2H8m13 9c-1.475 0-2.804.104-4 .291M3 16v3a2 2 0 0 0 2 2v0h11M3 16c1.403-.234 3.637-.293 5.945.243M3 16V5c0-.368.122-.939.5-1.377M16 21h3c.368 0 .939-.122 1.377-.5M16 21c-1.704-2.768-4.427-4.148-7.055-4.757m0 0c.927-1.073 2.24-2.084 4.055-2.85M7.341 7.5C7.14 7.728 7 8.051 7 8.5C7 9.7 8 10 8.5 10a1.66 1.66 0 0 0 1-.348M2 2l20 20"/>`
	imagePhotographyInnerSVG          = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10V5a1 1 0 0 1 1-1h1l2-2h4l2 2h1a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1z"/><circle cx="16" cy="7" r="1" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" fill-rule="evenodd" d="M7 2H5a3 3 0 0 0-3 3v10.5c0 .086.011.17.032.25A1 1 0 0 0 2 16v3a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-5H11.335c.567.246 1.157.544 1.741.895c1.399.84 2.821 2.02 3.778 3.583a1 1 0 1 1-1.706 1.044c-.736-1.203-1.878-2.178-3.102-2.913c-1.222-.734-2.465-1.192-3.327-1.392a15.466 15.466 0 0 0-3.703-.386h-.022c-.348.005-.68.02-.994.045V5a1 1 0 0 1 1-1h2V2zm0 4.473a2.23 2.23 0 0 0-.022.015C6.408 6.898 6 7.574 6 8.5c0 .926.408 1.601.978 2.011l.022.016V6.473z" clip-rule="evenodd"/></g>`
	imagePhotographyLineInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 8c-.5 0-1.5.3-1.5 1.5S6.5 11 7 11m-5 6v3a2 2 0 0 0 2 2v0h11M2 17V6a2 2 0 0 1 2-2h3M2 17c1.403-.234 3.637-.293 5.945.243M15 22h3a2 2 0 0 0 2-2v-6m-5 8c-1.704-2.768-4.427-4.148-7.055-4.757m0 0c1.095-1.268 2.73-2.45 5.096-3.243M10 10V5a1 1 0 0 1 1-1h1l2-2h4l2 2h1a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1z"/><circle cx="16" cy="7" r="1"/></g>`
	imagePlusInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 2v3m0 3V5m0 0h3m-3 0h-3"/><path fill="currentColor" fill-rule="evenodd" d="M13 2H5a3 3 0 0 0-3 3v10.5c0 .086.011.17.032.25A1 1 0 0 0 2 16v3a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-7a1 1 0 0 0-.032-.25A1 1 0 0 0 22 11.5V11h-2v.016c-4.297.139-7.4 1.174-9.58 2.623c.826.293 1.75.71 2.656 1.256c1.399.84 2.821 2.02 3.778 3.583a1 1 0 1 1-1.706 1.044c-.736-1.203-1.878-2.178-3.102-2.913c-1.222-.734-2.465-1.192-3.327-1.392a15.466 15.466 0 0 0-3.703-.386h-.022c-.348.005-.68.02-.994.045V5a1 1 0 0 1 1-1h8V2zM8.5 6a2.68 2.68 0 0 0-1.522.488C6.408 6.898 6 7.574 6 8.5c0 .926.408 1.601.978 2.011A2.674 2.674 0 0 0 8.5 11c.41 0 1.003-.115 1.522-.489c.57-.41.978-1.085.978-2.011c0-.926-.408-1.601-.978-2.012A2.674 2.674 0 0 0 8.5 6z" clip-rule="evenodd"/></g>`
	imagePlusLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12v7a2 2 0 0 1-2 2h-3m5-9c-6.442 0-10.105 1.985-12.055 4.243M21 12v-1.5M3 16v3a2 2 0 0 0 2 2v0h11M3 16V5a2 2 0 0 1 2-2h8M3 16c1.403-.234 3.637-.293 5.945.243M16 21c-1.704-2.768-4.427-4.148-7.055-4.757M8.5 7C8 7 7 7.3 7 8.5S8 10 8.5 10S10 9.7 10 8.5S9 7 8.5 7zM19 2v3m0 3V5m0 0h3m-3 0h-3"/>`
	inboxInnerSVG                     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 21a3 3 0 0 1-3-3v-3h5a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2h5v3a3 3 0 0 1-3 3H6zm15-8h-5a2 2 0 0 0-2 2h-4a2 2 0 0 0-2-2H3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v7z" fill="currentColor"/></g>`
	inboxInInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M9.707 7.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L13 8.586V3h5a3 3 0 0 1 3 3v7h-5a2 2 0 0 0-2 2h-4a2 2 0 0 0-2-2H3V6a3 3 0 0 1 3-3h5v5.586L9.707 7.293zM3 18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-3h-5a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2H3v3z" fill="currentColor"/></g>`
	inboxInLineInnerSVG               = `<g fill="currentColor"><path d="M8.293 7.293a1 1 0 0 1 1.414 0L11 8.586V3a1 1 0 1 1 2 0v5.586l1.293-1.293a1 1 0 1 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 0-1.414z"/><path d="M9 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-3v2h3a1 1 0 0 1 1 1v7h-3a2 2 0 0 0-2 2h-4a2 2 0 0 0-2-2H5V6a1 1 0 0 1 1-1h3V3zm7 12h3v3a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-3h3a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2z"/></g>`
	inboxLineInnerSVG                 = `<g fill="currentColor"><path d="M5 13h3a2 2 0 0 1 2 2h4a2 2 0 0 1 2-2h3V6a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v7zm14 2h-3a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2H5v3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-3zM3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6z"/></g>`
	incognitoInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M19 6v5H5V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11h2m16.5 0H19m0 0V6a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v5m14 0H5"/><circle cx="7" cy="17" r="3" fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="17" cy="17" r="3" fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 16h4"/></g>`
	incognitoLineInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 11h2m16.5 0H19m0 0V6a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v5m14 0H5"/><circle cx="7" cy="17" r="3"/><circle cx="17" cy="17" r="3"/><path d="M10 16h4"/></g>`
	infoCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm0 14a1 1 0 0 0 1-1v-3a1 1 0 1 0-2 0v3a1 1 0 0 0 1 1zm1-7a1 1 0 1 1-2 0a1 1 0 1 1 2 0z" clip-rule="evenodd"/>`
	infoCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 15v-3m0-3v0"/></g>`
	informationCircleInnerSVG         = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm11-4a1 1 0 1 0-2 0a1 1 0 0 0 2 0zm-1 9a1 1 0 0 0 1-1v-4a1 1 0 1 0-2 0v4a1 1 0 0 0 1 1z" fill="currentColor"/></g>`
	informationCircleLineInnerSVG     = `<g fill="currentColor"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm10 5a1 1 0 0 0 1-1v-4a1 1 0 1 0-2 0v4a1 1 0 0 0 1 1zm0-10a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H12z"/></g>`
	iphoneOldAppsInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M5 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V5zm3 1.001a1 1 0 0 0 2 0V6a1 1 0 1 0-2 0v.001zm4 1a1 1 0 0 1-1-1V6a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm2-1a1 1 0 1 0 2 0V6a1 1 0 1 0-2 0v.001zm-5 4a1 1 0 0 1-1-1V9a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm2-1a1 1 0 1 0 2 0V9a1 1 0 1 0-2 0v.001zm4 1a1 1 0 0 1-1-1V9a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm-7 2a1 1 0 1 0 2 0V12a1 1 0 1 0-2 0v.001zm4 1a1 1 0 0 1-1-1V12a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm-1 5a1 1 0 1 0 2 0V18a1 1 0 1 0-2 0v.001z" clip-rule="evenodd"/>`
	iphoneOldAppsLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 19V5a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zM9 6.001V6m3 .001V6m3 .001V6M9 9.001V9m3 .001V9m3 .001V9m-6 3.001V12m3 .001V12m0 6.001V18"/>`
	iphoneXAppsInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3H8a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-2m-4 0h4m-4 0v1m4-1v1m0 0v0a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v0m4 0h-4M9 8.001V8m0 3.001V11m3 .001V11m0-2.999V8m3 .001V8m0 3.001V11m-6 3.001V14m3 .001V14"/>`
	iphoneXAppsLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3H8a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-2m-4 0h4m-4 0v1m4-1v1m0 0v0a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v0m4 0h-4M9 8.001V8m0 3.001V11m3 .001V11m0-2.999V8m3 .001V8m0 3.001V11m-6 3.001V14m3 .001V14"/>`
	italicInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6h4m4 0h-4m0 0l-4 12m0 0h4m-4 0H6"/>`
	italicLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6h4m4 0h-4m0 0l-4 12m0 0h4m-4 0H6"/>`
	keyInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M8 9a7 7 0 1 1 5.562 6.852L12 17.414a2 2 0 0 1-1.414.586H10a2 2 0 0 1-2 2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-2.586A2 2 0 0 1 2.586 16l5.562-5.562A7.026 7.026 0 0 1 8 9zm7-3a1 1 0 1 0 0 2a1 1 0 0 1 1 1a1 1 0 1 0 2 0a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`
	keyLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15a6 6 0 1 0-5.743-4.257L9 11l-5.707 5.707a1 1 0 0 0-.293.707V20a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1a1 1 0 0 1 1-1a1 1 0 0 0 1-1a1 1 0 0 1 1-1h.586a1 1 0 0 0 .707-.293L13 15l.257-.257A5.999 5.999 0 0 0 15 15zm2-6a2 2 0 0 0-2-2"/>`
	keyboardInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M4 19a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H4zm6-6a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4zm3-3a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H14a1 1 0 0 1-1-1zm5-1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H18zm-9 1a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1zM6 9a1 1 0 0 0 0 2h.01a1 1 0 1 0 0-2H6zm-1 5a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1zm13-1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H18z" clip-rule="evenodd"/>`
	keyboardLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14h4m8 2V8a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2zm-8-6h.01M18 10h.01M10 10h.01M6 10h.01M6 14h.01M18 14h.01"/>`
	laptopInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v7.764l1.447 2.894l-.894.448l.894-.448c.998 1.995-.453 4.342-2.683 4.342H4.236c-2.23 0-3.68-2.347-2.683-4.342L3 13.764V6zm3-1a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H6z" clip-rule="evenodd"/>`
	laptopLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8H4V6zm16 8H4l-1.553 3.106A2 2 0 0 0 4.237 20h15.527a2 2 0 0 0 1.789-2.894L20 14z"/>`
	leafThreeAngledInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M20.156 6.473a2.639 2.639 0 0 0-2.63-2.629C15.049 3.817 9.909 4.193 6.5 7.6c-2.796 2.796-2.915 7.3-.157 10.057c2.758 2.758 7.26 2.639 10.057-.157c3.407-3.408 3.783-8.548 3.756-11.027zm-4.62 1.991a1 1 0 0 1 0 1.415l-.978.977l.587.195a1 1 0 0 1-.633 1.898l-1.535-.512l-1.247 1.247l.586.196a1 1 0 0 1-.632 1.897l-1.173-.39a3.01 3.01 0 0 1-.344-.14l-2.41 2.41l-3.121 3.121a1 1 0 0 1-1.414-1.414l3.121-3.121l2.41-2.41a3.009 3.009 0 0 1-.14-.344l-.39-1.173a1 1 0 0 1 1.897-.632l.196.586l1.247-1.247l-.512-1.535a1 1 0 0 1 1.898-.633l.195.587l.977-.978a1 1 0 0 1 1.414 0z" clip-rule="evenodd"/>`
	leafThreeAngledLineInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.05 16.95c2.343 2.343 6.213 2.273 8.643-.157c3.1-3.1 3.49-7.869 3.463-10.309a1.64 1.64 0 0 0-1.64-1.64c-2.44-.026-7.208.363-10.309 3.463c-2.43 2.43-2.5 6.3-.157 8.643zm0 0l-2.12 2.12m2.12-2.12l5.657-5.657m2.121-2.121l-2.12 2.12m0 0l2.12.708m-2.12-.707L12 9.172m0 5.656l-1.173-.39a2 2 0 0 1-1.265-1.265L9.172 12"/>`
	libraryInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M11.514 3.126a1 1 0 0 1 .972 0l9 5A1 1 0 0 1 21 10v9a1 1 0 1 1 0 2H3a1 1 0 1 1 0-2v-9a1 1 0 0 1-.486-1.874l9-5zM9 13a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0v-3zm4 0a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0v-3zm4 0a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0v-3z" clip-rule="evenodd"/>`
	libraryLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9v11M4 9h16M4 9H3l9-5l9 5h-1M4 20h16M4 20H3m17 0V9m0 11h1M8 13v3m4 0v-3m4 0v3"/>`
	lidquidDropWavesTwoInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M15 8.824C15 11.455 13 12 12 12s-3-.545-3-3.176C9 6.192 12 3 12 3s3 3.192 3 5.824z"/><path d="m3 15l.872.697a2.33 2.33 0 0 0 3.102-.171v0a2.33 2.33 0 0 1 3.164-.122l.18.154c.968.83 2.396.83 3.364 0l.18-.154a2.33 2.33 0 0 1 3.164.121v0a2.33 2.33 0 0 0 3.102.172L21 15m-.5 4l-.442.442c-.86.86-2.255.86-3.116 0v0a2.203 2.203 0 0 0-2.99-.114l-.27.23c-.968.83-2.396.83-3.364 0l-.27-.23a2.203 2.203 0 0 0-2.99.114v0c-.86.86-2.255.86-3.116 0L3.5 19"/></g>`
	lidquidDropWavesTwoLineInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8.824C15 11.455 13 12 12 12s-3-.545-3-3.176C9 6.192 12 3 12 3s3 3.192 3 5.824zM3 15l.872.697a2.33 2.33 0 0 0 3.102-.171v0a2.33 2.33 0 0 1 3.164-.122l.18.154c.968.83 2.396.83 3.364 0l.18-.154a2.33 2.33 0 0 1 3.164.121v0a2.33 2.33 0 0 0 3.102.172L21 15m-.5 4l-.442.442c-.86.86-2.255.86-3.116 0v0a2.203 2.203 0 0 0-2.99-.114l-.27.23c-.968.83-2.396.83-3.364 0l-.27-.23a2.203 2.203 0 0 0-2.99.114v0c-.86.86-2.255.86-3.116 0L3.5 19"/>`
	lifebuoyInnerSVG                  = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M18.33 19.742A9.959 9.959 0 0 1 12 22a9.959 9.959 0 0 1-6.33-2.258a.967.967 0 0 0 .037-.035l4.261-4.26c.595.351 1.29.553 2.032.553s1.437-.202 2.032-.554l4.26 4.261a.93.93 0 0 0 .038.035zm1.412-1.412A9.959 9.959 0 0 0 22 12a9.959 9.959 0 0 0-2.258-6.33a.967.967 0 0 1-.035.037l-4.26 4.261c.351.595.553 1.29.553 2.032s-.202 1.437-.554 2.032l4.261 4.26a.93.93 0 0 1 .035.038zM18.33 4.258a.977.977 0 0 0-.037.035l-4.261 4.26A3.982 3.982 0 0 0 12 8c-.742 0-1.437.202-2.032.554l-4.26-4.261a1.018 1.018 0 0 0-.038-.035A9.959 9.959 0 0 1 12 2a9.96 9.96 0 0 1 6.33 2.258zM4.258 5.67A9.959 9.959 0 0 0 2 12a9.96 9.96 0 0 0 2.258 6.33a.977.977 0 0 1 .035-.037l4.26-4.261A3.982 3.982 0 0 1 8 12c0-.742.202-1.437.554-2.032l-4.261-4.26a1.018 1.018 0 0 1-.035-.038z" clip-rule="evenodd"/><circle cx="12" cy="12" r="4" stroke="currentColor" stroke-width="2"/></g>`
	lifebuoyLineInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m18 6l-3.172 3.172M6 18l3.172-3.172M6 6l3.172 3.172M18 18l-3.172-3.172m0-5.656A3.987 3.987 0 0 0 12 8a3.987 3.987 0 0 0-2.828 1.172m5.656 0A3.984 3.984 0 0 1 16 12a3.987 3.987 0 0 1-1.172 2.828M9.172 9.172A3.987 3.987 0 0 0 8 12c0 1.105.448 2.105 1.172 2.828m0 0A3.987 3.987 0 0 0 12 16a3.987 3.987 0 0 0 2.828-1.172"/></g>`
	lightBulbInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M13 3a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0V3zM6.207 4.793a1 1 0 0 0-1.414 1.414l1 1a1 1 0 0 0 1.414-1.414l-1-1zm13 1.414a1 1 0 0 0-1.414-1.414l-1 1a1 1 0 0 0 1.414 1.414l1-1zM12 6a6 6 0 0 0-3.317 11h6.634A6 6 0 0 0 12 6zm3 12H9v1a3 3 0 1 0 6 0v-1zM3 11a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H3zm17 0a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2h-1z" fill="currentColor"/></g>`
	lightBulbLineInnerSVG             = `<g fill="currentColor"><path d="M12 2a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0V3a1 1 0 0 1 1-1zM4.793 4.793a1 1 0 0 1 1.414 0l1 1a1 1 0 0 1-1.414 1.414l-1-1a1 1 0 0 1 0-1.414zm14.414 0a1 1 0 0 1 0 1.414l-1 1a1 1 0 1 1-1.414-1.414l1-1a1 1 0 0 1 1.414 0zM12 8a4 4 0 1 0 0 8a4 4 0 0 0 0-8zm-6 4a6 6 0 1 1 9 5.197V19a3 3 0 1 1-6 0v-1.803A5.998 5.998 0 0 1 6 12zm5 5.917V19a1 1 0 1 0 2 0v-1.083a6.032 6.032 0 0 1-2 0zM2 12a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1zm17 0a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1z"/></g>`
	lightbulbShineInnerSVG            = `<g fill="none"><path fill="currentColor" d="M12 7a5 5 0 0 0-2 9.584V19h4v-2.416A5.001 5.001 0 0 0 12 7z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12h1m-3.5-6.5l1-1M12 3V2M5.5 5.5l-1-1M3 12H2m8 10h4m3-10a5 5 0 1 0-7 4.584V19h4v-2.416A5.001 5.001 0 0 0 17 12z"/></g>`
	lightbulbShineLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12h1m-3.5-6.5l1-1M12 3V2M5.5 5.5l-1-1M3 12H2m8 10h4m3-10a5 5 0 1 0-7 4.584V19h4v-2.416A5.001 5.001 0 0 0 17 12z"/>`
	lightningBoltInnerSVG             = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14L14 3v7h6L10 21v-7H4z"/>`
	lightningBoltLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14L14 3v7h6L10 21v-7H4z"/>`
	lineHeightInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10V5m0 0L4 7m2-2l2 2m-2 7v5m0 0l2-2m-2 2l-2-2m8-10h8m0 5h-8m0 5h8"/>`
	lineHeightLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10V5m0 0L4 7m2-2l2 2m-2 7v5m0 0l2-2m-2 2l-2-2m8-10h8m0 5h-8m0 5h8"/>`
	linkInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8h2c1.333 0 4 .8 4 4s-2.667 4-4 4h-2M9 8H7c-1.333 0-4 .8-4 4s2.667 4 4 4h2m-1-4h8"/>`
	linkCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm8-4c-.66 0-1.628.19-2.46.788C6.659 9.424 6 10.474 6 12c0 1.526.658 2.576 1.54 3.211A4.35 4.35 0 0 0 10 16a1 1 0 1 0 0-2c-.34 0-.872-.11-1.29-.412C8.341 13.325 8 12.874 8 12c0-.874.342-1.324.71-1.588C9.127 10.11 9.66 10 10 10a1 1 0 1 0 0-2zm4 0a1 1 0 1 0 0 2c.34 0 .872.11 1.29.412c.368.264.71.714.71 1.588c0 .874-.342 1.324-.71 1.588c-.418.302-.95.412-1.29.412a1 1 0 1 0 0 2c.66 0 1.628-.19 2.46-.789C17.341 14.576 18 13.526 18 12c0-1.526-.658-2.576-1.54-3.212A4.35 4.35 0 0 0 14 8zm-4 3a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4z" clip-rule="evenodd"/>`
	linkCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="9"/><path d="M10 9c-1 0-3 .6-3 3s2 3 3 3m4-6c1 0 3 .6 3 3s-2 3-3 3m-4-3h4"/></g>`
	linkLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8h2c1.333 0 4 .8 4 4s-2.667 4-4 4h-2M9 8H7c-1.333 0-4 .8-4 4s2.667 4 4 4h2m-1-4h8"/>`
	liraCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm11-6a1 1 0 1 0-2 0v1.28l-2.316.771a1 1 0 1 0 .632 1.898L10 9.387v.892l-2.316.772a1 1 0 0 0 .632 1.898L10 12.387V17a1 1 0 0 0 1 1c.993 0 2.461-.29 3.71-1.189C16.008 15.876 17 14.326 17 12a1 1 0 1 0-2 0c0 1.674-.675 2.624-1.46 3.188a4.402 4.402 0 0 1-1.54.687V11.72l2.316-.772a1 1 0 0 0-.632-1.898L12 9.613V8.72l2.316-.772a1 1 0 1 0-.632-1.898L12 6.613V6z" clip-rule="evenodd"/>`
	liraCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M11 6v11c1.667 0 5-1 5-5M8 9l6-2m-6 5l6-2"/></g>`
	listBoxInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6zm4 5a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2h-6a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2h-6a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2h-6a1 1 0 0 1-1-1zM7 7a1 1 0 0 0 0 2h.001a1 1 0 0 0 0-2H7zm-1 5a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1zm1 3a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	listBoxLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 8h5m0 4h-5m5 4h-5m-5 4h12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2zM8 8h.001M8 12h.001M8 16h.001"/>`
	locationMarkerInnerSVG            = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11.474 21.85L12 21l-.526.85zM12 21l.525.85c-.322.2-.73.2-1.051 0l-.003-.001l-.006-.004l-.02-.013a7.993 7.993 0 0 1-.311-.206a17.706 17.706 0 0 1-.841-.616a20.485 20.485 0 0 1-2.531-2.332C5.933 16.677 4 13.695 4 10c0-4.539 3.592-8 8-8c4.408 0 8 3.461 8 8c0 3.695-1.933 6.677-3.762 8.678a20.487 20.487 0 0 1-2.53 2.332a17.706 17.706 0 0 1-1.085.778a7.993 7.993 0 0 1-.068.044l-.02.013l-.006.004h-.002c0 .001-.002.002-.527-.849zM9 10a3 3 0 1 1 6 0a3 3 0 0 1-6 0z" fill="currentColor"/></g>`
	locationMarkerLineInnerSVG        = `<g fill="currentColor"><path d="M6 10c0-3.414 2.676-6 6-6s6 2.586 6 6c0 2.982-1.567 5.5-3.238 7.33A18.488 18.488 0 0 1 12 19.79a18.49 18.49 0 0 1-2.762-2.46C7.567 15.5 6 12.981 6 10zm5.474 11.85L12 21l-.525.85a1 1 0 0 0 1.05 0L12 21l.526.85h.001l.002-.001l.006-.004l.02-.013a7.993 7.993 0 0 0 .311-.206c.206-.141.496-.348.841-.616a20.487 20.487 0 0 0 2.531-2.332C18.067 16.677 20 13.695 20 10c0-4.539-3.592-8-8-8c-4.408 0-8 3.461-8 8c0 3.695 1.933 6.677 3.762 8.678a20.485 20.485 0 0 0 2.53 2.332a17.706 17.706 0 0 0 1.085.778c.029.02.052.034.068.044l.02.013l.006.004h.002v.001zM10 10a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm2-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8z"/></g>`
	lockInnerSVG                      = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M3 12a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-7zm10 2a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0v-3z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10V7a4 4 0 0 1 4-4v0a4 4 0 0 1 4 4v3"/></g>`
	lockClosedInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 4a3 3 0 0 0-3 3v3h6V7a3 3 0 0 0-3-3zM7 7v3H6a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3h-1V7A5 5 0 0 0 7 7zm6 8a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2z" fill="currentColor"/></g>`
	lockClosedLineInnerSVG            = `<g fill="currentColor"><path d="M3 13a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-6zm3-1a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H6z"/><path d="M7 7a5 5 0 0 1 10 0v4a1 1 0 1 1-2 0V7a3 3 0 1 0-6 0v4a1 1 0 1 1-2 0V7zm5 7a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1z"/></g>`
	lockLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10H6a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2h-2m-8 0V7a4 4 0 0 1 4-4v0a4 4 0 0 1 4 4v3m-8 0h8m-4 4v3"/>`
	lockOffInnerSVG                   = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M7 7.828V9H6a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h12c.872 0 1.657-.372 2.205-.966l-7.223-7.223A.97.97 0 0 1 13 14v3a1 1 0 1 1-2 0v-3a1 1 0 0 1 1.19-.982L7 7.828zm14 8.344V12a3 3 0 0 0-3-3h-1V7a5 5 0 0 0-8.62-3.449l1.415 1.415A3 3 0 0 1 15 7v2h-1.172L21 16.172z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18"/></g>`
	lockOffLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10H6a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h12a2 2 0 0 0 1.732-1M8 10V8m0 2h2m6 0h2a2 2 0 0 1 2 2v3m-4-5V7a4 4 0 0 0-6-3.465M16 10h-1m-3 4v3M3 3l18 18"/>`
	lockOpenInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M9 7c0-.507.16-1.289.612-1.916C10.026 4.508 10.726 4 12 4c1.274 0 1.974.508 2.389 1.084c.45.627.611 1.41.611 1.916a1 1 0 1 0 2 0c0-.827-.24-2.044-.988-3.084C15.226 2.825 13.926 2 12 2c-1.926 0-3.226.825-4.012 1.916C7.24 4.956 7 6.173 7 7v3H6a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3H9V7zm4 8a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2z" fill="currentColor"/></g>`
	lockOpenLineInnerSVG              = `<g fill="currentColor"><path d="M9.612 5.084C9.16 5.711 9 6.494 9 7v3h9a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h1V7c0-.827.24-2.044.988-3.084C8.774 2.825 10.074 2 12 2c1.926 0 3.226.825 4.012 1.916C16.76 4.956 17 6.173 17 7a1 1 0 1 1-2 0c0-.507-.16-1.289-.611-1.916C13.974 4.508 13.274 4 12 4c-1.274 0-1.974.508-2.388 1.084zM6 12a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H6zm6 2a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1z"/></g>`
	loginInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M11 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3h-6zm1.293 6.293a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414-1.414L13.586 13H5a1 1 0 1 1 0-2h8.586l-1.293-1.293a1 1 0 0 1 0-1.414z" clip-rule="evenodd"/>`
	loginHalfCircleInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M13.204 2.244C11.347 1.826 10 3.422 10 5v14c0 1.578 1.347 3.174 3.204 2.756C17.666 20.752 21 16.766 21 12s-3.334-8.752-7.796-9.756zm.089 6.049a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414-1.414L14.586 13H4a1 1 0 1 1 0-2h10.586l-1.293-1.293a1 1 0 0 1 0-1.414z" clip-rule="evenodd"/>`
	loginHalfCircleLineInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 12l-4-4m4 4l-4 4m4-4H5m5 9a9 9 0 1 0 0-18"/>`
	loginLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3h8a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H9m6-9l-4-4m4 4l-4 4m4-4H5"/>`
	logoutInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M6 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3H6zm10.293 5.293a1 1 0 0 1 1.414 0l4 4a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414-1.414L18.586 13H10a1 1 0 1 1 0-2h8.586l-2.293-2.293a1 1 0 0 1 0-1.414z" clip-rule="evenodd"/>`
	logoutHalfCircleInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M10.796 2.244C12.653 1.826 14 3.422 14 5v14c0 1.578-1.347 3.174-3.204 2.756C6.334 20.752 3 16.766 3 12s3.334-8.752 7.796-9.756zm5.497 6.049a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414-1.414L17.586 13H9a1 1 0 1 1 0-2h8.586l-1.293-1.293a1 1 0 0 1 0-1.414z" clip-rule="evenodd"/>`
	logoutHalfCircleLineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 12l-4-4m4 4l-4 4m4-4H9m5 9a9 9 0 1 1 0-18"/>`
	logoutLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m4-9l-4-4m4 4l-4 4m4-4H9"/>`
	mailInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M5 20a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H5zM7.625 8.22a1 1 0 1 0-1.25 1.56l3.75 3.001a3 3 0 0 0 3.75 0l3.75-3a1 1 0 1 0-1.25-1.562l-3.75 3a1 1 0 0 1-1.25 0l-3.75-3z" clip-rule="evenodd"/>`
	mailLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9l3.75 3a2 2 0 0 0 2.5 0L17 9m4 8V7a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2z"/>`
	mailOpenInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M4 10.165L2.033 8.554A3.004 3.004 0 0 1 4 6.17v3.994zm1 .82l5.732 4.696a2 2 0 0 0 2.536 0L19 10.985V5a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v5.985zm16.967-2.431L20 10.165V6.171a3.004 3.004 0 0 1 1.967 2.383zm-8.701 9.692L22 11.11V19a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-7.891l8.734 7.137a2 2 0 0 0 2.532 0zM10 9a1 1 0 0 0 0 2h4a1 1 0 1 0 0-2h-4zm0-4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4z" fill="currentColor"/></g>`
	mailOpenLineInnerSVG              = `<g fill="currentColor"><path d="M2 9a3 3 0 0 1 3-3h1a1 1 0 0 1 0 2H5a1 1 0 0 0-1 1v.51l7.386 5.746a1 1 0 0 0 1.228 0L20 9.51V9a1 1 0 0 0-1-1h-1a1 1 0 1 1 0-2h1a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V9zm18 3.045l-6.158 4.79a3 3 0 0 1-3.684 0L4 12.044V19a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-6.955z"/><path d="M8 4a1 1 0 0 0-1 1v7a1 1 0 1 1-2 0V5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v7a1 1 0 1 1-2 0V5a1 1 0 0 0-1-1H8z"/><path d="M9 7a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1zm0 3a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1z"/></g>`
	mapInnerSVG                       = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M21.472 3.118A1 1 0 0 1 22 4v12a1 1 0 0 1-.445.832L16 20.535V6.131l4.445-2.963a1 1 0 0 1 1.027-.05zM14 6.131l-4-2.666v14.404l4 2.666V6.131zM8 17.87l-4.445 2.963A1 1 0 0 1 2 20V8a1 1 0 0 1 .445-.832L8 3.465v14.404z" fill="currentColor"/></g>`
	mapLineInnerSVG                   = `<g fill="currentColor"><path d="M21.472 3.118A1 1 0 0 1 22 4v12a1 1 0 0 1-.445.832l-6 4a1 1 0 0 1-1.11 0L9 17.202l-5.445 3.63A1 1 0 0 1 2 20V8a1 1 0 0 1 .445-.832l6-4a1 1 0 0 1 1.11 0L15 6.798l5.445-3.63a1 1 0 0 1 1.027-.05zM14 8.535L10 5.87v9.596l4 2.666V8.535zm2 9.596l4-2.666V5.869l-4 2.666v9.596zm-8-2.666V5.869L4 8.535v9.596l4-2.666z"/></g>`
	mapMarkerInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M11.291 21.706L12 21l-.709.706zM12 21l.708.706a1 1 0 0 1-1.417 0l-.006-.007l-.017-.017l-.062-.063a47.708 47.708 0 0 1-1.04-1.106a49.562 49.562 0 0 1-2.456-2.908c-.892-1.15-1.804-2.45-2.497-3.734C4.535 12.612 4 11.248 4 10c0-4.539 3.592-8 8-8c4.408 0 8 3.461 8 8c0 1.248-.535 2.612-1.213 3.87c-.693 1.286-1.604 2.585-2.497 3.735a49.583 49.583 0 0 1-3.496 4.014l-.062.063l-.017.017l-.006.006L12 21zm0-8a3 3 0 1 0 0-6a3 3 0 0 0 0 6z" clip-rule="evenodd"/>`
	mapMarkerAreaInnerSVG             = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 16.016c1.245.529 2 1.223 2 1.984c0 1.657-3.582 3-8 3s-8-1.343-8-3c0-.76.755-1.456 2-1.984"/><path fill="currentColor" fill-rule="evenodd" d="M11.262 17.675L12 17l-.738.675zm1.476 0l.005-.005l.012-.014l.045-.05l.166-.186a38.19 38.19 0 0 0 2.348-2.957c.642-.9 1.3-1.92 1.801-2.933c.49-.99.885-2.079.885-3.086C18 4.871 15.382 2 12 2S6 4.87 6 8.444c0 1.007.395 2.096.885 3.086c.501 1.013 1.16 2.033 1.8 2.933a38.153 38.153 0 0 0 2.515 3.143l.045.05l.012.014l.005.005a1 1 0 0 0 1.476 0zM12 17l.738.674L12 17zm0-11a2 2 0 1 0 0 4a2 2 0 0 0 0-4z" clip-rule="evenodd"/></g>`
	mapMarkerAreaLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 16.016c1.245.529 2 1.223 2 1.984c0 1.657-3.582 3-8 3s-8-1.343-8-3c0-.76.755-1.456 2-1.984"/><path d="M17 8.444C17 11.537 12 17 12 17s-5-5.463-5-8.556C7 5.352 9.239 3 12 3s5 2.352 5 5.444z"/><circle cx="12" cy="8" r="1"/></g>`
	mapMarkerLineInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 10c0 3.976-7 11-7 11s-7-7.024-7-11s3.134-7 7-7s7 3.024 7 7z"/><circle cx="12" cy="10" r="3"/></g>`
	mapMarkerPathInnerSVG             = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm0 0v1m0 7v-1"/><path fill="currentColor" d="M7 19a2 2 0 1 0-4 0a2 2 0 0 0 4 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 19a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm0 0h1m9-2a2 2 0 1 1-2 2m2-2a2 2 0 0 0-2 2m2-2v-1m-2 3h-1"/><circle cx="5" cy="13" r="1" fill="currentColor"/><circle cx="11" cy="19" r="1" fill="currentColor"/><path fill="currentColor" fill-rule="evenodd" d="m17 13l-.647.763l-.001-.002l-.004-.002l-.01-.01l-.039-.033l-.137-.121a20.909 20.909 0 0 1-1.923-1.967c-.525-.617-1.072-1.343-1.491-2.109C12.335 8.766 12 7.892 12 7c0-1.547.538-2.825 1.49-3.711C14.431 2.413 15.69 2 17 2c1.31 0 2.569.413 3.51 1.289C21.462 4.175 22 5.453 22 7c0 .892-.335 1.766-.748 2.52c-.42.765-.965 1.491-1.49 2.108a20.909 20.909 0 0 1-2.061 2.088l-.038.033l-.011.01l-.004.002v.001h-.001L17 13zm0 0l.646.763a1 1 0 0 1-1.292 0L17 13zm0-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2z" clip-rule="evenodd"/></g>`
	mapMarkerPathLineInnerSVG         = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm0 0v1m0 7v-1m2 3a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm0 0h1m9-2a2 2 0 1 1-2 2m2-2a2 2 0 0 0-2 2m2-2v-1m-2 3h-1"/><circle cx="5" cy="13" r="1" fill="currentColor"/><circle cx="11" cy="19" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7c0 2.611-4 6-4 6s-4-3.389-4-6s1.79-4 4-4s4 1.389 4 4z"/><circle cx="17" cy="7" r="1" fill="currentColor"/></g>`
	mapMarkerPlusInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M11.291 21.706L12 21l-.709.706zM12 21l.708.706a1 1 0 0 1-1.417 0l-.006-.007l-.017-.017l-.062-.063a47.708 47.708 0 0 1-1.04-1.106a49.562 49.562 0 0 1-2.456-2.908c-.892-1.15-1.804-2.45-2.497-3.734C4.535 12.612 4 11.248 4 10c0-4.539 3.592-8 8-8c4.408 0 8 3.461 8 8c0 1.248-.535 2.612-1.213 3.87c-.693 1.286-1.604 2.585-2.497 3.735a49.583 49.583 0 0 1-3.496 4.014l-.062.063l-.017.017l-.006.006L12 21zm1-14a1 1 0 1 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V7z" clip-rule="evenodd"/>`
	mapMarkerPlusLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 10c0 3.976-7 11-7 11s-7-7.024-7-11s3.134-7 7-7s7 3.024 7 7zM9 10h3m3 0h-3m0 0V7m0 3v3"/>`
	mapSimpleInnerSVG                 = `<g fill="currentColor"><path d="M9 4L3 8v12l6-4l6 4l6-4V4l-6 4l-6-4z"/><path fill-rule="evenodd" d="M21.472 3.118A1 1 0 0 1 22 4v12a1 1 0 0 1-.445.832l-6 4a1 1 0 0 1-1.11 0L9 17.202l-5.445 3.63A1 1 0 0 1 2 20V8a1 1 0 0 1 .445-.832l6-4a1 1 0 0 1 1.11 0L15 6.798l5.445-3.63a1 1 0 0 1 1.027-.05zM4 8.535v9.596l4.445-2.963a1 1 0 0 1 1.11 0L15 18.798l5-3.333V5.869l-4.445 2.963a1 1 0 0 1-1.11 0L9 5.202L4 8.535z" clip-rule="evenodd"/></g>`
	mapSimpleDestinationInnerSVG      = `<path fill="currentColor" fill-rule="evenodd" d="M8.445 3.168a1 1 0 0 1 1.002-.062L15 5.882l5.553-2.776A1 1 0 0 1 22 4v12a1 1 0 0 1-.445.832l-6 4a1 1 0 0 1-1.002.062L9 18.118l-5.553 2.776A1 1 0 0 1 2 20V8a1 1 0 0 1 .445-.832l6-4zM5 12a1 1 0 1 0 2 0a1 1 0 1 0-2 0v.001zm5 1a1 1 0 0 1-1-1a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm4.707-3.708a1 1 0 1 0-1.414 1.414L14.586 12l-1.293 1.293a1 1 0 0 0 1.414 1.414L16 13.414l1.293 1.293a1 1 0 0 0 1.414-1.414L17.414 12l1.293-1.293a1 1 0 0 0-1.414-1.414L16 10.586l-1.293-1.293z" clip-rule="evenodd"/>`
	mapSimpleDestinationLineInnerSVG  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4L3 8v12l6-3l6 3l6-4V4l-6 3l-6-3zm-2 8.001V12m4 .001V12m3-2l2 2m2 2l-2-2m0 0l2-2m-2 2l-2 2"/>`
	mapSimpleLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4L3 8v12l6-4l6 4l6-4V4l-6 4l-6-4z"/>`
	mapSimpleMarkerInnerSVG           = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m17 21l-.647.763l-.001-.002l-.004-.002l-.01-.01l-.039-.033l-.137-.121a20.909 20.909 0 0 1-1.923-1.967c-.525-.617-1.072-1.343-1.491-2.109C12.335 16.766 12 15.892 12 15c0-1.547.538-2.825 1.49-3.711C14.431 10.413 15.69 10 17 10c1.31 0 2.569.413 3.51 1.289c.952.886 1.49 2.164 1.49 3.711c0 .892-.335 1.766-.748 2.52c-.42.765-.965 1.491-1.49 2.108a20.909 20.909 0 0 1-2.061 2.088l-.038.033l-.011.01l-.004.002v.001h-.001L17 21zm0 0l.646.763a1 1 0 0 1-1.292 0L17 21zm0-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/><path d="M22 4a1 1 0 0 0-1.555-.832L15 6.798l-5.445-3.63a1 1 0 0 0-1.11 0l-6 4A1 1 0 0 0 2 8v12a1 1 0 0 0 1.555.832L9 17.202l.75.5c-.254-.707-.417-1.45-.417-2.202c0-2.32.825-4.238 2.285-5.567C13.061 8.62 14.991 8 17 8c1.837 0 3.608.518 5 1.61V4z"/></g>`
	mapSimpleMarkerLineInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 9V4l-6 4l-6-4l-6 4v12l6-4l1 .667"/><path d="M21 15c0 2.611-4 6-4 6s-4-3.389-4-6s1.79-4 4-4s4 1.389 4 4zm-4 .001V15"/></g>`
	mapSimpleOffInnerSVG              = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18"/><path fill="currentColor" fill-rule="evenodd" d="M4.782 5.61L2.445 7.168A1 1 0 0 0 2 8v12a1 1 0 0 0 1.555.832L9 17.202l5.445 3.63a1 1 0 0 0 1.11 0l2.67-1.78L4.781 5.61zm16.834 11.178A1 1 0 0 0 22 16V4a1 1 0 0 0-1.555-.832L15 6.798l-5.445-3.63a1 1 0 0 0-1.11 0l-.269.18l13.44 13.44z" clip-rule="evenodd"/></g>`
	mapSimpleOffLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 4l6 4l6-4v12M6 6L3 8v12l6-4l6 4l3-2M3 3l18 18"/>`
	maximizeInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 20H4m0 0v-4m0 4l6-6m6-10h4m0 0v4m0-4l-6 6"/>`
	maximizeLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 20H4m0 0v-4m0 4l6-6m6-10h4m0 0v4m0-4l-6 6"/>`
	megaphoneInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M18.008 2.987C19.34 2.225 21 3.187 21 4.723v12.554c0 1.535-1.659 2.498-2.992 1.736L14 16.723V5.277l4.008-2.29zM12 6H7a5 5 0 0 0-1 9.9v3.6a2.5 2.5 0 0 0 5 0V16h1V6z" clip-rule="evenodd"/>`
	megaphoneLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 15V7m0 8l5.504 3.145A1 1 0 0 0 20 17.277V4.723a1 1 0 0 0-1.496-.868L13 7m0 8h-3m3-8H7a4 4 0 0 0-4 4v0a4 4 0 0 0 4 4v0m0 0v4.5A1.5 1.5 0 0 0 8.5 21v0a1.5 1.5 0 0 0 1.5-1.5V15m-3 0h3"/>`
	menuInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8h12M6 12h12M6 16h12"/>`
	menuAltInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 8a2 2 0 0 1 2-2h16a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2zm0 8a2 2 0 0 1 2-2h16a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2z" fill="currentColor"/></g>`
	menuAltLineInnerSVG               = `<g fill="currentColor"><path d="M3 8a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1zm0 8a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1z"/></g>`
	menuExpandLeftInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 8h-9m9 4H4m0 0l3-3m-3 3l3 3m13 1h-9"/>`
	menuExpandLeftLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 8h-9m9 4H4m0 0l3-3m-3 3l3 3m13 1h-9"/>`
	menuExpandRightInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h9m-9 4h16m0 0l-3-3m3 3l-3 3M4 16h9"/>`
	menuExpandRightLineInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h9m-9 4h16m0 0l-3-3m3 3l-3 3M4 16h9"/>`
	menuLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8h12M6 12h12M6 16h12"/>`
	messageInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 3a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3a1 1 0 0 1 1 1v.97c0 1.659 1.904 2.596 3.22 1.584l4.35-3.347a1 1 0 0 1 .61-.207H19a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H5z" fill="currentColor"/></g>`
	messageLineInnerSVG               = `<g fill="currentColor"><path d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-4.82a1 1 0 0 0-.61.207l-4.35 3.347C7.903 22.566 6 21.63 6 19.97V19a1 1 0 0 0-1-1a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1a3 3 0 0 1 3 3v.97l4.351-3.348a3 3 0 0 1 1.83-.622H19a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H5z"/></g>`
	messagesInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3h-3.586l-3.707 3.707A1 1 0 0 1 6 17v-3H5a3 3 0 0 1-3-3V5zm20 4v6c0 1-.6 3-3 3h-1v3c0 .333-.2 1-1 1c-.203 0-.368-.043-.5-.113L12.613 18H9l3-3h3c1.333 0 4-.8 4-4V6c1 0 3 .6 3 3z" fill="currentColor"/></g>`
	messagesLineInnerSVG              = `<g fill="currentColor"><path d="M2 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v1h1a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3h-1v3a1 1 0 0 1-1.707.707L12.586 18H7a1 1 0 0 1-1-1v-3H5a3 3 0 0 1-3-3V5zm16 3v3a3 3 0 0 1-3 3h-3.586l-2 2H13a1 1 0 0 1 .707.293L16 18.586V17a1 1 0 0 1 1-1h2a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-1zM8 14.586l2.293-2.293A1 1 0 0 1 11 12h4a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2a1 1 0 0 1 1 1v1.586z"/></g>`
	microphoneInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M9 6a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3v0a3 3 0 0 1-3-3V6z"/><path d="M12 18v0a7 7 0 0 1-7-7v0v-1m7 8v0a7 7 0 0 0 7-7v0v-1m-7 8v3"/></g>`
	microphoneLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3v0a3 3 0 0 1-3-3V6zm3 12v0a7 7 0 0 1-7-7v0v-1m7 8v0a7 7 0 0 0 7-7v0v-1m-7 8v3"/>`
	minimizeInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 14h4m0 0v4m0-4l-6 6m14-10h-4m0 0V6m0 4l6-6"/>`
	minimizeLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 14h4m0 0v4m0-4l-6 6m14-10h-4m0 0V6m0 4l6-6"/>`
	minusInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14"/>`
	minusCircleInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm7-1a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H9z" fill="currentColor"/></g>`
	minusCircleLineInnerSVG           = `<g fill="currentColor"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm6 0a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1z"/></g>`
	minusFiveCircleInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zm-9-4a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1.5a.5.5 0 0 1 0 1H14a1 1 0 1 0 0 2h1.5a2.5 2.5 0 0 0 0-5H15v-1h2a1 1 0 1 0 0-2h-3zm-7 3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	minusFiveCircleLineInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 12h3m7-3h-3v3h1.5a1.5 1.5 0 0 1 1.5 1.5v0a1.5 1.5 0 0 1-1.5 1.5H14"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`
	minusLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14"/>`
	minusTenCircleInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zM10 8a1 1 0 0 0 0 2h1v5a1 1 0 1 0 2 0V9a1 1 0 0 0-1-1h-2zm6.5 0a2.5 2.5 0 0 0-2.5 2.5v3a2.5 2.5 0 0 0 5 0v-3A2.5 2.5 0 0 0 16.5 8zm-.5 2.5a.5.5 0 0 1 1 0v3a.5.5 0 0 1-1 0v-3zM5 11a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2H5z" clip-rule="evenodd"/>`
	minusTenCircleLineInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12h3m2-3h2v6m6-1.5v-3A1.5 1.5 0 0 0 16.5 9v0a1.5 1.5 0 0 0-1.5 1.5v3a1.5 1.5 0 0 0 1.5 1.5v0a1.5 1.5 0 0 0 1.5-1.5z"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`
	moneyInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8zm9 4a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm1-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6z" clip-rule="evenodd"/>`
	moneyHandInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M16.948 9.95L14.998 8v6.587c0 .89-1.077 1.337-1.707.707L11.996 14c-.5-.5-1.701-.8-2.502 0c-.8.8-.5 2 0 2.5l3.603 4.4A3 3 0 0 0 15.42 22H18a1 1 0 0 0 1-1v-6.1a7 7 0 0 0-2.052-4.95z"/><path d="M11 2h2a2 2 0 0 1 2 2v2m-4-4c0 1.333.8 4 4 4m-4-4H9m6 4v6M5 12v2a2 2 0 0 0 2 2h2c0-1.333-.8-4-4-4zm0 0V6m4-4H7a2 2 0 0 0-2 2v2m4-4c0 1.333-.8 4-4 4"/><circle cx="10" cy="9" r="1" transform="rotate(90 10 9)"/></g>`
	moneyHandLineInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 22v-7.1a7 7 0 0 0-2.052-4.95L14.998 8v6.587c0 .89-1.077 1.337-1.707.707L11.996 14c-.5-.5-1.701-.8-2.502 0c-.8.8-.5 2 0 2.5l5.504 5.5"/><path d="M11 2h2a2 2 0 0 1 2 2v2m-4-4c0 1.333.8 4 4 4m-4-4H9m6 4v6M5 12v2a2 2 0 0 0 2 2h2c0-1.333-.8-4-4-4zm0 0V6m4-4H7a2 2 0 0 0-2 2v2m4-4c0 1.333-.8 4-4 4"/><circle cx="10" cy="9" r="1" transform="rotate(90 10 9)"/></g>`
	moneyLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 10V8a2 2 0 0 1 2-2h2m-4 4c1.333 0 4-.8 4-4m-4 4v4m18-4V8a2 2 0 0 0-2-2h-2m4 4c-1.333 0-4-.8-4-4m4 4v4M7 6h10m4 8v2a2 2 0 0 1-2 2h-2m4-4c-1.333 0-4 .8-4 4m0 0H7m-4-4v2a2 2 0 0 0 2 2h2m-4-4c1.333 0 4 .8 4 4"/><circle cx="12" cy="12" r="2"/></g>`
	moneyMinusInnerSVG                = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 18h6"/><path fill="currentColor" fill-rule="evenodd" d="M5 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h7.083A6.036 6.036 0 0 1 12 18c0-1.148.322-2.22.881-3.131A3.001 3.001 0 0 1 9 12a3 3 0 1 1 5.869.881A5.972 5.972 0 0 1 18 12c1.537 0 2.939.578 4 1.528V8a3 3 0 0 0-3-3H5zm7 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2z" clip-rule="evenodd"/></g>`
	moneyMinusLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 10V8a2 2 0 0 1 2-2h2m-4 4c1.333 0 4-.8 4-4m-4 4v4m18-4V8a2 2 0 0 0-2-2h-2m4 4c-1.333 0-4-.8-4-4m4 4v4M7 6h10M3 14v2a2 2 0 0 0 2 2h2m-4-4c1.333 0 4 .8 4 4m0 0h4"/><circle cx="12" cy="12" r="2"/><path d="M15 18h6"/></g>`
	moneyPlusInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 15v3m0 3v-3m0 0h-3m3 0h3"/><path fill="currentColor" fill-rule="evenodd" d="M5 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h7.083A6.036 6.036 0 0 1 12 18c0-1.148.322-2.22.881-3.131A3.001 3.001 0 0 1 9 12a3 3 0 1 1 5.869.881A5.972 5.972 0 0 1 18 12c1.537 0 2.939.578 4 1.528V8a3 3 0 0 0-3-3H5zm7 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2z" clip-rule="evenodd"/></g>`
	moneyPlusLineInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 10V8a2 2 0 0 1 2-2h2m-4 4c1.333 0 4-.8 4-4m-4 4v4m18-4V8a2 2 0 0 0-2-2h-2m4 4c-1.333 0-4-.8-4-4m4 4v2M7 6h10M3 14v2a2 2 0 0 0 2 2h2m-4-4c1.333 0 4 .8 4 4m0 0h4"/><circle cx="12" cy="12" r="2"/><path d="M18 15v3m0 3v-3m0 0h-3m3 0h3"/></g>`
	monitorInnerSVG                   = `<g fill="none"><path fill="currentColor" d="M19 4H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16h7a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h7zm0 0v4m0 0h4m-4 0H8"/></g>`
	monitorLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16h7a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h7zm0 0v4m0 0h4m-4 0H8"/>`
	moonInnerSVG                      = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.353 3C5.849 4.408 3 7.463 3 11.47A9.53 9.53 0 0 0 12.53 21c4.007 0 7.062-2.849 8.47-6.353C8.17 17.065 8.14 8.14 9.353 3z"/>`
	moonLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.353 3C5.849 4.408 3 7.463 3 11.47A9.53 9.53 0 0 0 12.53 21c4.007 0 7.062-2.849 8.47-6.353C8.17 17.065 8.14 8.14 9.353 3z"/>`
	moreMenuInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12h.01M8 12h.01M16 12h.01"/>`
	moreMenuLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12h.01M8 12h.01M16 12h.01"/>`
	moreMenuVerticalInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.005 11.995v.01m0-4.01v.01m0 7.99v.01"/>`
	moreMenuVerticalLineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.005 11.995v.01m0-4.01v.01m0 7.99v.01"/>`
	mouseInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M11 2.07A7.002 7.002 0 0 0 5 9v1h6V2.07zM5 12v3a7 7 0 1 0 14 0v-3H5zm14-2V9a7.002 7.002 0 0 0-6-6.93V10h6z" clip-rule="evenodd"/>`
	mouseLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 11V9a6 6 0 0 0-6-6v0m6 8v4a6 6 0 0 1-6 6v0a6 6 0 0 1-6-6v-4m12 0h-6m-6 0V9a6 6 0 0 1 6-6v0m-6 8h6m0 0V3"/>`
	multiplyInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12L6 6m6 6l6 6m-6-6l6-6m-6 6l-6 6"/>`
	multiplyLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12L6 6m6 6l6 6m-6-6l6-6m-6 6l-6 6"/>`
	musicInnerSVG                     = `<g fill="none"><circle cx="6" cy="18" r="3" fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="18" cy="17" r="3" fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" d="M21 3L9 6v4l12-3V3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 18v-8m12 7V7M9 10V6l12-3v4M9 10l12-3"/></g>`
	musicLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="18" r="3"/><circle cx="18" cy="17" r="3"/><path d="M9 18v-8m12 7V7M9 10V6l12-3v4M9 10l12-3"/></g>`
	musicNoteInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M13 4h4v4h-4v9c0 1-.6 3-3 3s-3-2-3-3s.6-3 3-3s3 2 3 3V4z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17V8m0 9c0 1-.6 3-3 3s-3-2-3-3s.6-3 3-3s3 2 3 3zm0-9V4h4v4h-4z"/></g>`
	musicNoteLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17V8m0 9c0 1-.6 3-3 3s-3-2-3-3s.6-3 3-3s3 2 3 3zm0-9V4h4v4h-4z"/>`
	newspaperInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h14c.493 0 1.211-.14 1.834-.588C21.51 19.925 22 19.125 22 18v-7.002A2.999 2.999 0 0 0 19 8v7.5a.5.5 0 0 1-1 0V6a3 3 0 0 0-3-3H5zm2 4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H7zm0 8a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H7z" fill="currentColor"/></g>`
	newspaperLineInnerSVG             = `<g fill="currentColor"><path d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H5z"/><path d="M6 8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V8zm2 1v2h2V9H8zm6-1a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1zm-8 4a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1z"/></g>`
	nextCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm6.498-4.865a1 1 0 0 1 .998-.003L15 10.848V8a1 1 0 1 1 2 0v8a1 1 0 1 1-2 0v-2.848l-6.504 3.716A1 1 0 0 1 7 16V8a1 1 0 0 1 .498-.865z" clip-rule="evenodd"/>`
	nextCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 8v8M8 8v8l7-4l-7-4z"/><circle cx="12" cy="12" r="10"/></g>`
	noteTextInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M5 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h9v-5a3 3 0 0 1 3-3h5V5a3 3 0 0 0-3-3H5zm12.293 19.121a3 3 0 0 1-1.293.762V17a1 1 0 0 1 1-1h4.883a3 3 0 0 1-.762 1.293l-3.828 3.828zM7 6a1 1 0 0 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	noteTextLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10m6-6v.172a2 2 0 0 1-.586 1.414l-3.828 3.828a2 2 0 0 1-1.414.586H15m6-6h-4a2 2 0 0 0-2 2v4M7 7h10M7 11h10M7 15h4"/>`
	noteTextMinusInnerSVG             = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6H3"/><path fill="currentColor" fill-rule="evenodd" d="M12 6h5a1 1 0 1 1 0 2h-5.341a5.997 5.997 0 0 1-1.187 2H17a1 1 0 1 1 0 2H7a.998.998 0 0 1-.287-.042A5.978 5.978 0 0 1 2 10.472V19a3 3 0 0 0 3 3h9v-5a3 3 0 0 1 3-3h5V5a3 3 0 0 0-3-3h-8.528A5.978 5.978 0 0 1 12 6zm4 15.883a3 3 0 0 0 1.293-.762l3.828-3.828A3 3 0 0 0 21.883 16H17a1 1 0 0 0-1 1v4.883zM6 15a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1z" clip-rule="evenodd"/></g>`
	noteTextMinusLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15V5a2 2 0 0 0-2-2h-6m8 12v.172a2 2 0 0 1-.586 1.414l-3.828 3.828a2 2 0 0 1-1.414.586H15m6-6h-4a2 2 0 0 0-2 2v4m0 0H5a2 2 0 0 1-2-2v-8m10-4h4M7 11h10M7 15h4M9 6H3"/>`
	noteTextPlusInnerSVG              = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3v3m0 3V6m0 0h3M6 6H3"/><path fill="currentColor" fill-rule="evenodd" d="M12 6h5a1 1 0 1 1 0 2h-5.341a5.997 5.997 0 0 1-1.187 2H17a1 1 0 1 1 0 2H7a.998.998 0 0 1-.287-.042A5.978 5.978 0 0 1 2 10.472V19a3 3 0 0 0 3 3h9v-5a3 3 0 0 1 3-3h5V5a3 3 0 0 0-3-3h-8.528A5.978 5.978 0 0 1 12 6zm4 15.883a3 3 0 0 0 1.293-.762l3.828-3.828A3 3 0 0 0 21.883 16H17a1 1 0 0 0-1 1v4.883zM6 15a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1z" clip-rule="evenodd"/></g>`
	noteTextPlusLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15V5a2 2 0 0 0-2-2h-7m9 12v.172a2 2 0 0 1-.586 1.414l-3.828 3.828a2 2 0 0 1-1.414.586H15m6-6h-4a2 2 0 0 0-2 2v4m0 0H5a2 2 0 0 1-2-2v-7m10-5h4m-7 4h7M7 15h4M6 3v3m0 3V6m0 0h3M6 6H3"/>`
	noteblockInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M8 2a1 1 0 0 1 1 1v1H7V3a1 1 0 0 1 1-1zm9 2h1a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h1v3a1 1 0 0 0 2 0V4h2v3a1 1 0 1 0 2 0V4h2v3a1 1 0 1 0 2 0V4zm0 0h-2V3a1 1 0 1 1 2 0v1zm-4 0V3a1 1 0 1 0-2 0v1h2z" clip-rule="evenodd"/>`
	noteblockLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3v4m4 0V3m4 0v4M6 21h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2z"/>`
	noteblockTextInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M9 3a1 1 0 0 0-2 0v1H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3h-1V3a1 1 0 1 0-2 0v1h-2V3a1 1 0 1 0-2 0v1H9V3zm2 1v3a1 1 0 1 0 2 0V4h-2zm4 0h2v3a1 1 0 1 1-2 0V4zM9 4v3a1 1 0 0 1-2 0V4h2zm-1 6a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`
	noteblockTextLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3v4m4 0V3m4 0v4m-8 4h8m-8 4h4m-6 6h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2z"/>`
	openInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-4m-8-2l8-8m0 0v5m0-5h-5"/>`
	openLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-4m-8-2l8-8m0 0v5m0-5h-5"/>`
	paperAirplaneInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2.245 21.655a1 1 0 0 1-.14-1.102l9-18a1 1 0 0 1 1.79 0l9 18a1 1 0 0 1-1.211 1.396L13 19.387V12a1 1 0 1 0-2 0v7.387L3.316 21.95a1 1 0 0 1-1.071-.294z" fill="currentColor"/></g>`
	paperAirplaneLineInnerSVG         = `<g fill="currentColor"><path d="M2.245 21.655a1 1 0 0 1-.14-1.102l9-18a1 1 0 0 1 1.79 0l9 18a1 1 0 0 1-1.211 1.396L12 19.054l-8.684 2.895a1 1 0 0 1-1.071-.294zM13 17.28l6.026 2.009L12 5.236L4.974 19.288L11 17.279V12a1 1 0 1 1 2 0v5.28z"/></g>`
	paperClipInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 2a4 4 0 0 0-4 4v12a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4H6zm13 7a1 1 0 1 0-2 0v4a5 5 0 0 1-10 0V9a3 3 0 0 1 6 0v4a1 1 0 1 1-2 0V9a1 1 0 1 0-2 0v4a3 3 0 1 0 6 0V9A5 5 0 0 0 5 9v4a7 7 0 1 0 14 0V9z" fill="currentColor"/></g>`
	paperClipLineInnerSVG             = `<g fill="currentColor"><path d="M18 6a1 1 0 0 1 1 1v8a7 7 0 1 1-14 0V7a5 5 0 0 1 10 0v8a3 3 0 1 1-6 0V7a1 1 0 1 1 2 0v8a1 1 0 1 0 2 0V7a3 3 0 1 0-6 0v8a5 5 0 0 0 10 0V7a1 1 0 0 1 1-1z"/></g>`
	paperFoldInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M6 21h12c-1 0-3-.6-3-3v-2H3v2c0 2.4 2 3 3 3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v13c0 1-.6 3-3 3m0 0H6c-1 0-3-.6-3-3v-2h12v2c0 2.4 2 3 3 3z"/></g>`
	paperFoldLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v13c0 1-.6 3-3 3m0 0H6c-1 0-3-.6-3-3v-2h12v2c0 2.4 2 3 3 3z"/>`
	paperFoldTextInnerSVG             = `<g fill="none"><path fill="currentColor" d="M6 21h12c-1 0-3-.6-3-3v-2H3v2c0 2.4 2 3 3 3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v13c0 1-.6 3-3 3m0 0H6c-1 0-3-.6-3-3v-2h12v2c0 2.4 2 3 3 3zM9 7h8m-8 4h4"/></g>`
	paperFoldTextLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v13c0 1-.6 3-3 3m0 0H6c-1 0-3-.6-3-3v-2h12v2c0 2.4 2 3 3 3zM9 7h8m-8 4h4"/>`
	paperRollTwoInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M5.29 4.188C6.54 3.29 8.008 3 9 3h7a6 6 0 1 1-4 10.472V20a1 1 0 0 1-1.555.832l-1.017-.678l-1.48.74a1 1 0 0 1-.895 0l-1.481-.74l-1.017.678A1 1 0 0 1 3 20V9c0-2.326.992-3.876 2.29-4.812zM11.529 5H9c-.673 0-1.706.21-2.54.812C5.674 6.376 5 7.326 5 9v9.134a1 1 0 0 1 .947-.028l1.553.776l1.553-.776a1 1 0 0 1 .947.028V9c0-1.537.578-2.938 1.528-4zM16 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4z" clip-rule="evenodd"/>`
	paperRollTwoLineInnerSVG          = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 9a5 5 0 1 0 5-5m-5 5a5 5 0 0 1 5-5m-5 5v11l-1.5-1l-2 1l-2-1L4 20V9c0-4 3.333-5 5-5h7"/><circle cx="16" cy="9" r="2" fill="currentColor"/></g>`
	paragraphInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 5h-3m-5 0h5m-5 0v10m0-10h-1c-1.667 0-5 1-5 5s3.333 5 5 5h1m0 4v-4m5 4V5"/>`
	paragraphLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 5h-3m-5 0h5m-5 0v10m0-10h-1c-1.667 0-5 1-5 5s3.333 5 5 5h1m0 4v-4m5 4V5"/>`
	pauseInnerSVG                     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm9-3a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0V9zm4 0a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0V9z" fill="currentColor"/></g>`
	pauseCircleInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1zM6 9a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V9zm2 5h1v-4H8v4zm6-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-3zm2 6h-1v-4h1v4z" clip-rule="evenodd"/>`
	pauseCircleLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 9h3v6H7zm7 0h3v6h-3z"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`
	pauseLineInnerSVG                 = `<g fill="currentColor"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm8-4a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1zm4 0a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1z"/></g>`
	pencilInnerSVG                    = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M15.586 3a2 2 0 0 1 2.828 0L21 5.586a2 2 0 0 1 0 2.828L19.414 10L14 4.586L15.586 3zm-3 3l-9 9A2 2 0 0 0 3 16.414V19a2 2 0 0 0 2 2h2.586A2 2 0 0 0 9 20.414l9-9L12.586 6z" fill="currentColor"/></g>`
	pencilAltInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M19.414 3a2 2 0 0 0-2.828 0l-8 8A2 2 0 0 0 8 12.414V14a2 2 0 0 0 2 2h1.586A2 2 0 0 0 13 15.414l8-8a2 2 0 0 0 0-2.828L19.414 3zM6 4a3 3 0 0 0-3 3v11a3 3 0 0 0 3 3h11a3 3 0 0 0 3-3v-3a1 1 0 1 0-2 0v3a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h3a1 1 0 1 0 0-2H6z" fill="currentColor"/></g>`
	pencilAltLineInnerSVG             = `<g fill="currentColor"><path d="M16.586 3a2 2 0 0 1 2.828 0L21 4.586a2 2 0 0 1 0 2.828l-8 8a2 2 0 0 1-1.414.586H10a2 2 0 0 1-2-2v-1.586A2 2 0 0 1 8.586 11l8-8zm-.172 3L18 7.586L19.586 6l.707.707L18 4.414L16.414 6zm.172 3L15 7.414l-5 5V14h1.586l5-5zM3 7a3 3 0 0 1 3-3h3a1 1 0 0 1 0 2H6a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-3a1 1 0 1 1 2 0v3a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V7z"/></g>`
	pencilLineInnerSVG                = `<g fill="currentColor"><path d="M15.586 3a2 2 0 0 1 2.828 0L21 5.586a2 2 0 0 1 0 2.828l-12 12A2 2 0 0 1 7.586 21H5a2 2 0 0 1-2-2v-2.586A2 2 0 0 1 3.586 15l12-12zm-.172 3L18 8.586L19.586 7L17 4.414L15.414 6zm1.172 4L14 7.414l-9 9V19h2.586l9-9z"/></g>`
	percentInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 18L18 6"/><circle cx="17" cy="17" r="2" fill="currentColor"/><circle cx="7" cy="7" r="2" fill="currentColor"/></g>`
	percentLineInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 18L18 6"/><circle cx="16.5" cy="16.5" r="2.5"/><circle cx="7.5" cy="7.5" r="2.5"/></g>`
	phoneInnerSVG                     = `<g fill="none"><path fill="currentColor" d="M20 16v4c-2.758 0-5.07-.495-7-1.325c-3.841-1.652-6.176-4.63-7.5-7.675C4.4 8.472 4 5.898 4 4h4l1 4l-3.5 3c1.324 3.045 3.659 6.023 7.5 7.675L16 15l4 1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 18.675c1.93.83 4.242 1.325 7 1.325v-4l-4-1l-3 3.675zm0 0C9.159 17.023 6.824 14.045 5.5 11m0 0C4.4 8.472 4 5.898 4 4h4l1 4l-3.5 3z"/></g>`
	phoneDialInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M19 17v4c-2.758 0-5.07-.495-7-1.325c-3.841-1.652-6.176-4.63-7.5-7.675C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3c1.324 3.045 3.659 6.023 7.5 7.675L15 16l4 1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19.675c1.93.83 4.242 1.325 7 1.325v-4l-4-1l-3 3.675zm0 0C8.159 18.023 5.824 15.045 4.5 12m0 0C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3zM14 4h.01M17 4h.01M20 4h.01M14 7h.01M17 7h.01M20 7h.01M14 10h.01M17 10h.01M20 10h.01"/></g>`
	phoneDialLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19.675c1.93.83 4.242 1.325 7 1.325v-4l-4-1l-3 3.675zm0 0C8.159 18.023 5.824 15.045 4.5 12m0 0C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3zM14 4h.01M17 4h.01M20 4h.01M14 7h.01M17 7h.01M20 7h.01M14 10h.01M17 10h.01M20 10h.01"/>`
	phoneHangupInnerSVG               = `<g fill="none"><path fill="currentColor" d="M23 12.5L20.5 15l-3-2V8.842C15.976 8.337 14.146 8 12 8c-2.145 0-3.976.337-5.5.842V13l-3 2L1 12.5c.665-.997 2.479-2.657 5.5-3.658C8.024 8.337 9.855 8 12 8c2.146 0 3.976.337 5.5.842c3.021 1 4.835 2.66 5.5 3.658z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 8.842C15.976 8.337 14.146 8 12 8c-2.145 0-3.976.337-5.5.842m11 0c3.021 1 4.835 2.66 5.5 3.658L20.5 15l-3-2V8.842zm-11 0c-3.021 1-4.835 2.66-5.5 3.658L3.5 15l3-2V8.842z"/></g>`
	phoneHangupLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 8.842C15.976 8.337 14.146 8 12 8c-2.145 0-3.976.337-5.5.842m11 0c3.021 1 4.835 2.66 5.5 3.658L20.5 15l-3-2V8.842zm-11 0c-3.021 1-4.835 2.66-5.5 3.658L3.5 15l3-2V8.842z"/>`
	phoneIncomingInnerSVG             = `<g fill="none"><path fill="currentColor" d="M19 17v4c-2.758 0-5.07-.495-7-1.325c-3.841-1.652-6.176-4.63-7.5-7.675C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3c1.324 3.045 3.659 6.023 7.5 7.675L15 16l4 1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19.675c1.93.83 4.242 1.325 7 1.325v-4l-4-1l-3 3.675zm0 0C8.159 18.023 5.824 15.045 4.5 12m0 0C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3zM20 4l-6 6m0 0V6m0 4h4"/></g>`
	phoneIncomingLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19.675c1.93.83 4.242 1.325 7 1.325v-4l-4-1l-3 3.675zm0 0C8.159 18.023 5.824 15.045 4.5 12m0 0C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3zM20 4l-6 6m0 0V6m0 4h4"/>`
	phoneLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 18.675c1.93.83 4.242 1.325 7 1.325v-4l-4-1l-3 3.675zm0 0C9.159 17.023 6.824 14.045 5.5 11m0 0C4.4 8.472 4 5.898 4 4h4l1 4l-3.5 3z"/>`
	phoneMissedCallInnerSVG           = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M14.293 3.293a1 1 0 0 1 1.414 0L17.5 5.086l1.793-1.793a1 1 0 1 1 1.414 1.414L18.914 6.5l1.793 1.793a1 1 0 0 1-1.414 1.414L17.5 7.914l-1.793 1.793a1 1 0 0 1-1.414-1.414L16.086 6.5l-1.793-1.793a1 1 0 0 1 0-1.414zM3 4a1 1 0 0 0-1 1c0 2.023.424 4.734 1.583 7.399c1.4 3.22 3.895 6.42 8.022 8.194C13.672 21.483 16.121 22 19 22a1 1 0 0 0 1-1v-4a1 1 0 0 0-.758-.97l-4-1a1 1 0 0 0-1.017.338l-2.509 3.073c-2.894-1.443-4.796-3.735-5.991-6.174L8.65 9.76a1 1 0 0 0 .32-1.002l-1-4A1 1 0 0 0 7 4H3z" fill="currentColor"/></g>`
	phoneMissedCallLineInnerSVG       = `<g fill="currentColor"><path d="M15.707 3.293a1 1 0 1 0-1.414 1.414L16.086 6.5l-1.793 1.793a1 1 0 0 0 1.414 1.414L17.5 7.914l1.793 1.793a1 1 0 0 0 1.414-1.414L18.914 6.5l1.793-1.793a1 1 0 0 0-1.414-1.414L17.5 5.086l-1.793-1.793zM3 4a1 1 0 0 0-1 1c0 2.023.424 4.734 1.583 7.399c1.4 3.22 3.895 6.42 8.022 8.194C13.672 21.483 16.121 22 19 22a1 1 0 0 0 1-1v-4a1 1 0 0 0-.758-.97l-4-1a1 1 0 0 0-1.017.338l-2.509 3.073c-2.894-1.443-4.796-3.735-5.991-6.174L8.65 9.76a1 1 0 0 0 .32-1.002l-1-4A1 1 0 0 0 7 4H3zm15 15.977c-1.63-.078-3.07-.343-4.343-.75l1.716-2.103l2.627.657v2.196zM6.88 8.643L4.922 10.32A17.47 17.47 0 0 1 4.037 6h2.182l.661 2.643z"/></g>`
	phoneOutgoingInnerSVG             = `<g fill="none"><path fill="currentColor" d="M19 17v4c-2.758 0-5.07-.495-7-1.325c-3.841-1.652-6.176-4.63-7.5-7.675C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3c1.324 3.045 3.659 6.023 7.5 7.675L15 16l4 1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19.675c1.93.83 4.242 1.325 7 1.325v-4l-4-1l-3 3.675zm0 0C8.159 18.023 5.824 15.045 4.5 12m0 0C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3zm9.5-2l6-6m0 0v4m0-4h-4"/></g>`
	phoneOutgoingLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19.675c1.93.83 4.242 1.325 7 1.325v-4l-4-1l-3 3.675zm0 0C8.159 18.023 5.824 15.045 4.5 12m0 0C3.4 9.472 3 6.898 3 5h4l1 4l-3.5 3zm9.5-2l6-6m0 0v4m0-4h-4"/>`
	phoneRetroInnerSVG                = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 4.842C15.976 4.337 14.146 4 12 4c-2.145 0-3.976.337-5.5.842m11 0c3.021 1 4.835 2.66 5.5 3.658L20.5 11l-3-2V4.842zm-11 0c-3.021 1-4.835 2.66-5.5 3.658L3.5 11l3-2V4.842z"/><path fill="currentColor" fill-rule="evenodd" d="M10 6a1 1 0 0 1 1 1v2h2V7a1 1 0 1 1 2 0v2.586l5.121 5.121A3 3 0 0 1 21 16.828V18a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-1.172a3 3 0 0 1 .879-2.12L9 9.585V7a1 1 0 0 1 1-1zm2 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4z" clip-rule="evenodd"/></g>`
	phoneRetroLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 4.842C15.976 4.337 14.146 4 12 4c-2.145 0-3.976.337-5.5.842m11 0c3.021 1 4.835 2.66 5.5 3.658L20.5 11l-3-2V4.842zm-11 0c-3.021 1-4.835 2.66-5.5 3.658L3.5 11l3-2V4.842zM10 7v3m0 0l-5.414 5.414A2 2 0 0 0 4 16.828V18a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-1.172a2 2 0 0 0-.586-1.414L14 10m-4 0h4m0 0V7"/><circle cx="12" cy="15" r="2"/></g>`
	phoneRingInnerSVG                 = `<g fill="none"><path fill="currentColor" d="M23 17.5L20.5 20l-3-2v-4.158C15.976 13.337 14.146 13 12 13c-2.145 0-3.976.337-5.5.842V18l-3 2L1 17.5c.665-.997 2.479-2.657 5.5-3.658C8.024 13.337 9.855 13 12 13c2.146 0 3.976.337 5.5.842c3.021 1 4.835 2.66 5.5 3.658z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 13.842C15.976 13.337 14.146 13 12 13c-2.145 0-3.976.337-5.5.842m11 0c3.021 1 4.835 2.66 5.5 3.658L20.5 20l-3-2v-4.158zm-11 0c-3.021 1-4.835 2.66-5.5 3.658L3.5 20l3-2v-4.158zM12 4v4M4.636 6.636l2.828 2.828m11.9-2.828l-2.828 2.828"/></g>`
	phoneRingLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 13.842C15.976 13.337 14.146 13 12 13c-2.145 0-3.976.337-5.5.842m11 0c3.021 1 4.835 2.66 5.5 3.658L20.5 20l-3-2v-4.158zm-11 0c-3.021 1-4.835 2.66-5.5 3.658L3.5 20l3-2v-4.158zM12 4v4M4.636 6.636l2.828 2.828m11.9-2.828l-2.828 2.828"/>`
	pillInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M4.929 19.071a5 5 0 0 1 0-7.071l2.828-2.828l7.071 7.07L12 19.073a5 5 0 0 1-7.071 0zm11.314-4.243L19.07 12A5 5 0 0 0 12 4.929L9.172 7.757l7.07 7.071z" clip-rule="evenodd"/>`
	pillLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.172 9.172l-3.536 3.535a4 4 0 0 0 0 5.657v0a4 4 0 0 0 5.657 0l3.535-3.536M9.172 9.172l3.535-3.536a4 4 0 0 1 5.657 0v0a4 4 0 0 1 0 5.657l-3.536 3.535M9.172 9.172l5.656 5.656"/>`
	pinInnerSVG                       = `<g fill="none"><path fill="currentColor" d="M12.956 18.956L9 15l-3.956-3.956a1 1 0 0 1 .314-1.626l5.261-2.255a1 1 0 0 0 .535-.548l1.283-3.207a1 1 0 0 1 1.635-.336l6.856 6.856a1 1 0 0 1-.336 1.635l-3.207 1.283a1 1 0 0 0-.548.535l-2.255 5.261a1 1 0 0 1-1.626.314z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 20l5-5m0 0l3.956 3.956a1 1 0 0 0 1.626-.314l2.255-5.261a1 1 0 0 1 .548-.535l3.207-1.283a1 1 0 0 0 .336-1.635l-6.856-6.856a1 1 0 0 0-1.635.336l-1.283 3.207a1 1 0 0 1-.535.548L5.358 9.418a1 1 0 0 0-.314 1.626L9 15z"/></g>`
	pinLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 20l5-5m0 0l3.956 3.956a1 1 0 0 0 1.626-.314l2.255-5.261a1 1 0 0 1 .548-.535l3.207-1.283a1 1 0 0 0 .336-1.635l-6.856-6.856a1 1 0 0 0-1.635.336l-1.283 3.207a1 1 0 0 1-.535.548L5.358 9.418a1 1 0 0 0-.314 1.626L9 15z"/>`
	pinwheelInnerSVG                  = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7a5.001 5.001 0 0 0-4.607 3.053C11.963 11.07 12.896 12 14 12h6c1.105 0 2.037-.93 1.607-1.947A5.001 5.001 0 0 0 17 7zm0 10a5.001 5.001 0 0 0-3.053-4.607C12.93 11.963 12 12.896 12 14v6c0 1.105.93 2.037 1.947 1.607A5.001 5.001 0 0 0 17 17zM7 17a5.001 5.001 0 0 0 4.607-3.053C12.037 12.93 11.104 12 10 12H4c-1.105 0-2.037.93-1.607 1.947A5.001 5.001 0 0 0 7 17zM7 7a5.001 5.001 0 0 0 3.053 4.607C11.07 12.037 12 11.104 12 10V4c0-1.105-.93-2.037-1.947-1.607A5.001 5.001 0 0 0 7 7z"/>`
	pinwheelLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7a5.001 5.001 0 0 0-4.607 3.053C11.963 11.07 12.896 12 14 12h6c1.105 0 2.037-.93 1.607-1.947A5.001 5.001 0 0 0 17 7zm0 10a5.001 5.001 0 0 0-3.053-4.607C12.93 11.963 12 12.896 12 14v6c0 1.105.93 2.037 1.947 1.607A5.001 5.001 0 0 0 17 17zM7 17a5.001 5.001 0 0 0 4.607-3.053C12.037 12.93 11.104 12 10 12H4c-1.105 0-2.037.93-1.607 1.947A5.001 5.001 0 0 0 7 17zM7 7a5.001 5.001 0 0 0 3.053 4.607C11.07 12.037 12 11.104 12 10V4c0-1.105-.93-2.037-1.947-1.607A5.001 5.001 0 0 0 7 7z"/>`
	planetInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4.859 5A9.973 9.973 0 0 1 12 2c4.479 0 8.267 2.943 9.542 7H15a1 1 0 1 0 0 2h6.95c.033.33.05.663.05 1c0 5.523-4.477 10-10 10a9.972 9.972 0 0 1-7.141-3H11a1 1 0 1 0 0-2H3.338a9.95 9.95 0 0 1-.88-2H13a1 1 0 1 0 0-2H2.05a9.957 9.957 0 0 1 1.289-6H12a1 1 0 1 0 0-2H4.859zM10 9a1 1 0 0 0 0 2h2a1 1 0 1 0 0-2h-2zm-2.292 1.708a1 1 0 0 1-1.414 0h-.001a1 1 0 1 1 1.414-1.415a1 1 0 0 1 0 1.415zm7.586 4a1 1 0 1 0 1.414-1.414v-.001a1 1 0 0 0-1.415 1.414z" clip-rule="evenodd"/>`
	planetLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.292 6A8.967 8.967 0 0 0 3 12c0 .687.077 1.357.223 2m2.069-8a9.003 9.003 0 0 1 15.485 4M5.292 6H12m8.777 4a9 9 0 0 1-15.485 8m15.485-8H15M3.223 14H13m-9.777 0a8.975 8.975 0 0 0 2.069 4m0 0H11m1-8h-2m-3 0h0m9 4h0"/>`
	planetRingTwoInnerSVG             = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 7.96c2.59-.125 4.379.274 4.625 1.193c.429 1.6-3.98 4.172-9.849 5.745c-5.868 1.572-10.972 1.55-11.401-.051c-.254-.948 1.188-2.236 3.625-3.455"/><path fill="currentColor" fill-rule="evenodd" d="M4 12a8 8 0 1 1 15.985.491c-1.653.879-3.904 1.754-6.467 2.44c-2.874.77-5.526 1.14-7.478 1.131a12.275 12.275 0 0 1-.956-.039A7.963 7.963 0 0 1 4 12zm2.766 6.05a8.003 8.003 0 0 0 12.658-3.065c-1.561.697-3.4 1.346-5.389 1.879c-2.668.715-5.208 1.115-7.269 1.186z" clip-rule="evenodd"/></g>`
	planetRingTwoLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="7"/><path d="M18 7.96c2.59-.125 4.379.274 4.625 1.193c.429 1.6-3.98 4.172-9.849 5.745c-5.868 1.572-10.972 1.55-11.401-.051c-.254-.948 1.188-2.236 3.625-3.455"/></g>`
	planetRocketInnerSVG              = `<g fill="none"><g clip-path="url(#majesticonsPlanetRocket0)"><path fill="currentColor" d="m21.048 8.868l1.402-.318l-.318 1.402a8 8 0 0 1-2.145 3.89L17.5 16.328l-.015.015c1.71 1.709-.702 4.935-1.414 5.628l-1.4-2.814l-2.828-2.829L9 14.9c.722-.703 4.001-3.1 5.686-1.415l2.472-2.472a8 8 0 0 1 3.89-2.145z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.672 19.157l-2.829-2.829m2.829 2.829l1.4 2.814c.711-.693 3.122-3.919 1.413-5.628m-2.813 2.814l2.813-2.814m-5.642-.015l2.829-2.828l.014-.015m-2.843 2.843L9 14.9c.722-.703 4.001-3.1 5.686-1.415m2.814 2.843l-.015.015m0 0l2.502-2.501a8 8 0 0 0 2.145-3.89l.318-1.402l-1.402.318a8 8 0 0 0-3.89 2.145l-2.472 2.472m-11.272-.172c-1.339 2.117-1.85 3.806-1.192 4.465c.586.586 1.987.246 3.778-.778m7.313-13.586c2.117-1.339 3.806-1.85 4.465-1.192c.886.885-.345 3.634-2.854 6.778m-10.67 5A7.002 7.002 0 0 1 14 4.254"/></g><defs><clipPath id="majesticonsPlanetRocket0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	planetRocketLineInnerSVG          = `<g fill="none"><g clip-path="url(#majesticonsPlanetRocketLine0)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.672 19.157l-2.829-2.829m2.829 2.829l1.4 2.814c.711-.693 3.122-3.919 1.413-5.628m-2.813 2.814l2.813-2.814m-5.642-.015l2.829-2.828l.014-.015m-2.843 2.843L9 14.9c.722-.703 4.001-3.1 5.686-1.415m2.814 2.843l-.015.015m0 0l2.502-2.501a8 8 0 0 0 2.145-3.89l.318-1.402l-1.402.318a8 8 0 0 0-3.89 2.145l-2.472 2.472m-11.272-.172c-1.339 2.117-1.85 3.806-1.192 4.465c.586.586 1.987.246 3.778-.778m7.313-13.586c2.117-1.339 3.806-1.85 4.465-1.192c.886.885-.345 3.634-2.854 6.778m-10.67 5A7.002 7.002 0 0 1 14 4.254"/></g><defs><clipPath id="majesticonsPlanetRocketLine0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	playCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm8.528-4.882a1 1 0 0 1 1.027.05l6 4a1 1 0 0 1 0 1.664l-6 4A1 1 0 0 1 9 16V8a1 1 0 0 1 .528-.882z" clip-rule="evenodd"/>`
	playCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 16V8l6 4l-6 4z"/><circle cx="12" cy="12" r="10"/></g>`
	playlistInnerSVG                  = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6h-6m6 4h-6m6 4h-8m8 4H4"/><path fill="currentColor" d="M8 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm0 0V7m0 0h2V6H8v1z"/></g>`
	playlistLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6h-6m6 4h-6m6 4h-8m8 4H4m4-6a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm0 0V7m0 0h2V6H8v1z"/>`
	plusInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h7m7 0h-7m0 0V5m0 7v7"/>`
	plusCircleInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm11-3a1 1 0 1 0-2 0v2H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V9z" fill="currentColor"/></g>`
	plusCircleLineInnerSVG            = `<g fill="currentColor"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm10-4a1 1 0 0 1 1 1v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9a1 1 0 1 1 0-2h2V9a1 1 0 0 1 1-1z"/></g>`
	plusFiveCircleInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zm-9-4a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1.5a.5.5 0 0 1 0 1H14a1 1 0 1 0 0 2h1.5a2.5 2.5 0 0 0 0-5H15v-1h2a1 1 0 1 0 0-2h-3zm-4 2a1 1 0 1 0-2 0v1H7a1 1 0 1 0 0 2h1v1a1 1 0 1 0 2 0v-1h1a1 1 0 1 0 0-2h-1v-1z" clip-rule="evenodd"/>`
	plusFiveCircleLineInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 12h2m2 0H9m0 0v-2m0 2v2m8-5h-3v3h1.5a1.5 1.5 0 0 1 1.5 1.5v0a1.5 1.5 0 0 1-1.5 1.5H14"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`
	plusLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h7m7 0h-7m0 0V5m0 7v7"/>`
	plusMinusInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h3m3 0H7m0 0V9m0 3v3m7-3h6"/>`
	plusMinusLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h3m3 0H7m0 0V9m0 3v3m7-3h6"/>`
	plusMinusTwoInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 4L4 20M4 7h3m3 0H7m0 0V4m0 3v3m7 7h6"/>`
	plusMinusTwoLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 4L4 20M4 7h3m3 0H7m0 0V4m0 3v3m7 7h6"/>`
	plusTenCircleInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zM10 8a1 1 0 0 0 0 2h1v5a1 1 0 1 0 2 0V9a1 1 0 0 0-1-1h-2zm6.5 0a2.5 2.5 0 0 0-2.5 2.5v3a2.5 2.5 0 0 0 5 0v-3A2.5 2.5 0 0 0 16.5 8zm-.5 2.5a.5.5 0 0 1 1 0v3a.5.5 0 0 1-1 0v-3zM8 10a1 1 0 1 0-2 0v1H5a1 1 0 1 0 0 2h1v1a1 1 0 1 0 2 0v-1h1a1 1 0 1 0 0-2H8v-1z" clip-rule="evenodd"/>`
	plusTenCircleLineInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12h2m2 0H7m0 0v-2m0 2v2m3-5h2v6m6-1.5v-3A1.5 1.5 0 0 0 16.5 9v0a1.5 1.5 0 0 0-1.5 1.5v3a1.5 1.5 0 0 0 1.5 1.5v0a1.5 1.5 0 0 0 1.5-1.5z"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`
	poundCircleInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm12-6c-.493 0-1.211.14-1.834.588C10.49 7.074 10 7.874 10 9v2H9a1 1 0 1 0 0 2h1v1c0 .173-.06.456-.212.666c-.114.159-.314.334-.788.334a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-3.166c.12-.367.166-.72.166-1v-1h1a1 1 0 1 0 0-2h-1V9c0-.474.175-.674.334-.788c.21-.152.493-.212.666-.212c.173 0 .456.06.666.212c.159.114.334.314.334.788a1 1 0 1 0 2 0c0-1.126-.492-1.926-1.166-2.412A3.233 3.233 0 0 0 13 6z" clip-rule="evenodd"/>`
	poundCircleLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M15 9c0-1.6-1.333-2-2-2s-2 .4-2 2v3m0 0v2c0 .667-.4 2-2 2h6m-4-4h2m-2 0H9"/></g>`
	presentationInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M3 2a1 1 0 0 0 0 2v11a1 1 0 0 0 1 1h4.613L7.05 20.684a1 1 0 0 0 1.898.632L10.72 16h2.558l1.772 5.316a1 1 0 0 0 1.898-.632L15.387 16H20a1 1 0 0 0 1-1V4a1 1 0 1 0 0-2H3z" clip-rule="evenodd"/>`
	presentationChartInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2v11a1 1 0 0 1-1 1h-4.613l1.562 4.684a1 1 0 0 1-1.898.632L13.28 16h-2.558L8.95 21.316a1 1 0 1 1-1.898-.632L8.613 16H4a1 1 0 0 1-1-1V4a1 1 0 0 1-1-1zm15 4a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0V7zm-4 2a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0V9zm-5 1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`
	presentationChartLineInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m4 0V7m-8 4h.01M3 3h1m17 0h-1m0 0v12h-6m6-12H4m0 0v12h6m0 0l-2 6m2-6h4m0 0l2 6"/>`
	presentationLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h1m17 0h-1m0 0v12h-6m6-12H4m0 0v12h6m0 0l-2 6m2-6h4m0 0l2 6"/>`
	presentationLineLineInnerSVG      = `<g fill="currentColor"><path d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-3.586l2.293 2.293a1 1 0 0 1-1.414 1.414L12 17.414l-3.293 3.293a1 1 0 0 1-1.414-1.414L9.586 17H6a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H6zm10.6 2.2a1 1 0 0 1 .2 1.4l-3 4a1 1 0 0 1-1.507.107l-2.138-2.137l-1.323 1.985a1 1 0 0 1-1.664-1.11l2-3a1 1 0 0 1 1.54-.152l2.184 2.185L15.2 7.4a1 1 0 0 1 1.4-.2z"/></g>`
	presentationPlayInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2v11a1 1 0 0 1-1 1h-4.613l1.562 4.684a1 1 0 0 1-1.898.632L13.28 16h-2.558L8.95 21.316a1 1 0 1 1-1.898-.632L8.613 16H4a1 1 0 0 1-1-1V4a1 1 0 0 1-1-1zm9.555 3.168A1 1 0 0 0 10 7v4a1 1 0 0 0 1.555.832l3-2a1 1 0 0 0 0-1.664l-3-2zM12.197 9L12 9.131V8.87l.197.131z" clip-rule="evenodd"/>`
	presentationPlayLineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h1m17 0h-1m0 0v12h-6m6-12H4m0 0v12h6m0 0l-2 6m2-6h4m0 0l2 6m-5-10V7l3 2l-3 2z"/>`
	presentationReportInnerSVG        = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-3.586l2.293 2.293a1 1 0 0 1-1.414 1.414L12 17.414l-3.293 3.293a1 1 0 0 1-1.414-1.414L9.586 17H6a3 3 0 0 1-3-3V6zm14 2a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0V8zm-4 2a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2zm-4 1a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0v-1z" fill="currentColor"/></g>`
	presentationReportLineInnerSVG    = `<g fill="currentColor"><path d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-3.586l2.293 2.293a1 1 0 0 1-1.414 1.414L12 17.414l-3.293 3.293a1 1 0 0 1-1.414-1.414L9.586 17H6a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H6zm10 2a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1zm-4 2a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1zm-5 3a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1z"/></g>`
	printerInnerSVG                   = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 17v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2M7 17v2a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-2M7 17H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1m0 0V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v2M6 7h12m0 0h1a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2"/><path fill="currentColor" fill-rule="evenodd" d="M2 9a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3h-2a1 1 0 0 1-1-1v-2a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v2a1 1 0 0 1-1 1H5a3 3 0 0 1-3-3V9zm5 0a1 1 0 0 0 0 2h1a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/></g>`
	printerLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 17v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2M7 17v2a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-2M7 17H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1m0 0V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v2M6 7h12m0 0h1a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2M7 10h1"/>`
	pulseInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h4l3 7l4-14l3 7h4"/>`
	pulseLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h4l3 7l4-14l3 7h4"/>`
	puzzleInnerSVG                    = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.75 6H20a1 1 0 0 1 1 1v3.25a.75.75 0 0 1-.75.75H20a2 2 0 1 0 0 4h.25a.75.75 0 0 1 .75.75V20a1 1 0 0 1-1 1h-3.25a.75.75 0 0 1-.75-.75V20a2 2 0 1 0-4 0v.25a.75.75 0 0 1-.75.75H7a1 1 0 0 1-1-1v-4.25a.75.75 0 0 0-.75-.75H5a2 2 0 1 1 0-4h.25a.75.75 0 0 0 .75-.75V7a1 1 0 0 1 1-1h4.25a.75.75 0 0 0 .75-.75V5a2 2 0 1 1 4 0v.25c0 .414.336.75.75.75z"/>`
	puzzleLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.75 6H20a1 1 0 0 1 1 1v3.25a.75.75 0 0 1-.75.75H20a2 2 0 1 0 0 4h.25a.75.75 0 0 1 .75.75V20a1 1 0 0 1-1 1h-3.25a.75.75 0 0 1-.75-.75V20a2 2 0 1 0-4 0v.25a.75.75 0 0 1-.75.75H7a1 1 0 0 1-1-1v-4.25a.75.75 0 0 0-.75-.75H5a2 2 0 1 1 0-4h.25a.75.75 0 0 0 .75-.75V7a1 1 0 0 1 1-1h4.25a.75.75 0 0 0 .75-.75V5a2 2 0 1 1 4 0v.25c0 .414.336.75.75.75z"/>`
	qrCodeInnerSVG                    = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.001"/><path fill="currentColor" fill-rule="evenodd" d="M4 3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4zm1 6V5h4v4H5zm9-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-6zM4 13a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H4zm10 0a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1h-1zm5 0a1 1 0 0 0-.707.293l-2 2a1 1 0 0 0 1.414 1.414L19.414 15H20a1 1 0 1 0 0-2h-1zm-5 4a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-1h1a1 1 0 1 0 0-2h-2zm7 1a1 1 0 1 0-2 0v1h-1a1 1 0 1 0 0 2h2a1 1 0 0 0 1-1v-2zM16 7a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H17a1 1 0 0 1-1-1zm-9 9a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/></g>`
	qrCodeLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 18v2h-2m2-6h-1l-2 2m-1 2h-2v2M4 4h6v6H4V4zm10 0h6v6h-6V4zM4 14h6v6H4v-6zm10 0v1h1v-1h-1zm3-7h.001M7 7h.001M7 17h.001"/>`
	qrcodeInnerSVG                    = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5zm8 1a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0V4zm4-1a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-2zm-4 6a1 1 0 1 0-2 0v2a2 2 0 0 0 2 2h1a1 1 0 1 0 0-2h-1V9zm-9 2a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H4zm14 0a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2zM5 15a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2H5zm8 0a2 2 0 0 0-2 2v3a1 1 0 1 0 2 0v-3h1a1 1 0 1 0 0-2h-1zm5 0a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2zm-2 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4z" fill="currentColor"/></g>`
	qrcodeLineInnerSVG                = `<g fill="currentColor"><path d="M3 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5zm4 0H5v2h2V5zm5-2a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0V4a1 1 0 0 1 1-1zm3 2a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2V5zm4 0h-2v2h2V5zm-7 3a1 1 0 0 1 1 1v2h1a1 1 0 1 1 0 2h-1a2 2 0 0 1-2-2V9a1 1 0 0 1 1-1zm-9 4a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1zm14 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1zM3 17a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2zm4 0H5v2h2v-2zm4 0a2 2 0 0 1 2-2h1a1 1 0 1 1 0 2h-1v3a1 1 0 1 1-2 0v-3zm6-1a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1zm-2 4a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1z"/></g>`
	questionCircleInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm0 7a1 1 0 0 0-1 1a1 1 0 1 1-2 0a3 3 0 1 1 4.44 2.633a1.404 1.404 0 0 0-.383.288a.303.303 0 0 0-.057.085v.494a1 1 0 1 1-2 0V13c0-.58.253-1.047.539-1.38c.281-.33.63-.572.94-.742A1 1 0 0 0 12 9zm.999 4.011v-.004v.005zM12 15a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H12z" clip-rule="evenodd"/>`
	questionCircleLineInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M12 12a2 2 0 1 0-2-2m2 5h.01"/></g>`
	questionMarkCircleInnerSVG        = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm10-5a2 2 0 0 0-2 2a1 1 0 0 1-2 0a4 4 0 1 1 5.31 3.78a.674.674 0 0 0-.273.169a.177.177 0 0 0-.037.054v.497a1 1 0 1 1-2 0V13c0-1.152.924-1.856 1.655-2.11A2.001 2.001 0 0 0 12 7zm1 6.007v-.004v.004zM13 17a1 1 0 1 1-2 0a1 1 0 0 1 2 0z" fill="currentColor"/></g>`
	questionMarkCircleLineInnerSVG    = `<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm0 6a2 2 0 0 0-2 2a1 1 0 1 1-2 0a4 4 0 1 1 5.31 3.78a.674.674 0 0 0-.273.169a.177.177 0 0 0-.037.054v.497a1 1 0 1 1-2 0V14c0-1.152.924-1.856 1.655-2.11A2.001 2.001 0 0 0 12 8zm1 6.007v-.004v.004zM11 17a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H12a1 1 0 0 1-1-1z"/></g>`
	radioListInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h10m-10 7h10m-10 7h10"/><circle cx="5" cy="5" r="2" fill="currentColor"/><circle cx="5" cy="5" r="2" fill="currentColor"/><circle cx="5" cy="12" r="2" fill="currentColor"/><circle cx="5" cy="19" r="2" fill="currentColor"/></g>`
	radioListLineInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 5h10m-10 7h10m-10 7h10"/><circle cx="5" cy="5" r="2"/><circle cx="5" cy="5" r="2"/><circle cx="5" cy="12" r="2"/><circle cx="5" cy="19" r="2"/></g>`
	receiptInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7 2a3 3 0 0 0-3 3v16a1 1 0 0 0 1.496.868L8.5 20.152l2.012 1.15a3 3 0 0 0 2.976 0l2.012-1.15l3.004 1.716A1 1 0 0 0 20 21V5a3 3 0 0 0-3-3H7z" fill="currentColor"/></g>`
	receiptLineInnerSVG               = `<g fill="currentColor"><path d="M4 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v16a1 1 0 0 1-1.496.868L15.5 20.152l-2.012 1.15a3 3 0 0 1-2.976 0L8.5 20.151l-3.004 1.716A1 1 0 0 1 4 21V5zm3-1a1 1 0 0 0-1 1v14.277l2.004-1.145a1 1 0 0 1 .992 0l2.508 1.433a1 1 0 0 0 .992 0l2.508-1.433a1 1 0 0 1 .992 0L18 19.277V5a1 1 0 0 0-1-1H7z"/></g>`
	receiptRefundInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M4 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v16a1 1 0 0 1-1.496.868L15.5 20.152l-2.012 1.15a3 3 0 0 1-2.976 0L8.5 20.151l-3.004 1.716A1 1 0 0 1 4 21V5zm6.293 1.293a1 1 0 1 1 1.414 1.414L10.414 9H12a4 4 0 0 1 4 4v1a1 1 0 1 1-2 0v-1a2 2 0 0 0-2-2h-1.586l1.293 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414l3-3z" fill="currentColor"/></g>`
	receiptRefundLineInnerSVG         = `<g fill="currentColor"><path d="M7 2a3 3 0 0 0-3 3v16a1 1 0 0 0 1.496.868L8.5 20.152l2.012 1.15a3 3 0 0 0 2.976 0l2.012-1.15l3.004 1.716A1 1 0 0 0 20 21V5a3 3 0 0 0-3-3H7zM6 5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v14.277l-2.004-1.145a1 1 0 0 0-.992 0l-2.508 1.433a1 1 0 0 1-.992 0l-2.508-1.433a1 1 0 0 0-.992 0L6 19.277V5zm4.293 1.293a1 1 0 1 1 1.414 1.414L10.414 9H12a4 4 0 0 1 4 4v1a1 1 0 1 1-2 0v-1a2 2 0 0 0-2-2h-1.586l1.293 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414l3-3z"/></g>`
	receiptTextInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M3 5a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v16a1 1 0 0 1-1.625.78l-1.929-1.542l-2.391 1.594a1 1 0 0 1-1.18-.051L12 20.28l-1.875 1.5a1 1 0 0 1-1.18.051l-2.391-1.594l-1.93 1.543A1 1 0 0 1 3 21V5zm5 1a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`
	receiptTextLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h8m-8 4h8m-8 4h4m8 6V5a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v16l2.5-2l3 2l2.5-2l2.5 2l3-2l2.5 2z"/>`
	redoInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 9v5h-5M4 16c.497-4.5 3.367-8 8-8c2.73 0 5.929 2.268 7.294 5.5"/>`
	redoLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 9v5h-5M4 16c.497-4.5 3.367-8 8-8c2.73 0 5.929 2.268 7.294 5.5"/>`
	refreshInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M17.217 5.333a1 1 0 1 1-1.49 1.334A5 5 0 0 0 7 10h2a1 1 0 0 1 .707 1.707l-3 3a1 1 0 0 1-1.414 0l-3-3A1 1 0 0 1 3 10h2a7 7 0 0 1 12.217-4.667zm4.49 6.96A1 1 0 0 1 21 14h-2v.002a7 7 0 0 1-12.217 4.665a1 1 0 1 1 1.49-1.334A5 5 0 0 0 17 14h-2a1 1 0 0 1-.707-1.707l3-3a1 1 0 0 1 1.414 0l3 3z" fill="currentColor"/></g>`
	refreshLineInnerSVG               = `<g fill="currentColor"><path d="M21.707 13.707a1 1 0 0 0 0-1.414l-3-3a1 1 0 0 0-1.414 0l-3 3a1 1 0 0 0 1.414 1.414L17 12.414V14a1 1 0 1 0 2 0v-1.586l1.293 1.293a1 1 0 0 0 1.414 0zm-12-3.414a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L5 11.586V10a1 1 0 0 1 2 0v1.586l1.293-1.293a1 1 0 0 1 1.414 0z"/><path d="M18 13a1 1 0 0 1 1 1a7 7 0 0 1-12.217 4.667a1 1 0 1 1 1.49-1.334A5 5 0 0 0 17 14a1 1 0 0 1 1-1zm-.86-6.255a1 1 0 0 0 .077-1.412A7 7 0 0 0 5 10a1 1 0 1 0 2 0a5 5 0 0 1 8.727-3.333a1 1 0 0 0 1.412.078z"/></g>`
	reloadInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 13.5A7.5 7.5 0 1 1 11.5 6H20m0 0l-3-3m3 3l-3 3"/>`
	reloadCircleInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm10.293-3.293a1 1 0 0 1 1.414-1.414l2 2a1 1 0 0 1 0 1.414l-2 2a1 1 0 0 1-1.414-1.414l.293-.293H12c-.34 0-.872.11-1.29.412c-.368.264-.71.714-.71 1.588c0 .874.342 1.324.71 1.588c.418.302.95.412 1.29.412c.368 0 .945-.128 1.37-.473a1 1 0 0 1 1.26 1.554c-.87.705-1.93.919-2.63.919a4.35 4.35 0 0 1-2.46-.788C8.659 15.576 8 14.526 8 13c0-1.526.658-2.576 1.54-3.212A4.35 4.35 0 0 1 12 9h.586l-.293-.293z" clip-rule="evenodd"/>`
	reloadCircleLineInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m13 8l2 2l-2 2m0-2h-1c-1 0-3 .6-3 3s2 3 3 3c.534 0 1.353-.171 2-.695"/></g>`
	reloadLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 13.5A7.5 7.5 0 1 1 11.5 6H20m0 0l-3-3m3 3l-3 3"/>`
	removeColumnInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h3M3 21h3m0 0h4a2 2 0 0 0 2-2V9M6 21V9m0-6h4a2 2 0 0 1 2 2v4M6 3v6M3 9h3m0 0h6m-9 6h9m3-6l3 3m0 0l3 3m-3-3l3-3m-3 3l-3 3"/>`
	removeColumnLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h3M3 21h3m0 0h4a2 2 0 0 0 2-2V9M6 21V9m0-6h4a2 2 0 0 1 2 2v4M6 3v6M3 9h3m0 0h6m-9 6h9m3-6l3 3m0 0l3 3m-3-3l3-3m-3 3l-3 3"/>`
	removeFormatInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6h5m5 0h-5m0 0l-2 6m-2 6l.667-2M5 5l14 14"/>`
	removeFormatLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6h5m5 0h-5m0 0l-2 6m-2 6l.667-2M5 5l14 14"/>`
	removeRowInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v3m18-3v3m0 0v4a2 2 0 0 1-2 2H9m12-6H9M3 6v4a2 2 0 0 0 2 2h4M3 6h6m0-3v3m0 0v6m6-9v9m-6 3l3 3m0 0l3 3m-3-3l3-3m-3 3l-3 3"/>`
	removeRowLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v3m18-3v3m0 0v4a2 2 0 0 1-2 2H9m12-6H9M3 6v4a2 2 0 0 0 2 2h4M3 6h6m0-3v3m0 0v6m6-9v9m-6 3l3 3m0 0l3 3m-3-3l3-3m-3 3l-3 3"/>`
	repeatCircleInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zM10.707 6.707a1 1 0 0 0-1.414-1.414l-2 2a1 1 0 0 0 0 1.414l2 2a1 1 0 0 0 1.414-1.414L10.414 9H14a1 1 0 0 1 1 1v1a1 1 0 1 0 2 0v-1a3 3 0 0 0-3-3h-3.586l.293-.293zM9 13a1 1 0 1 0-2 0v1a3 3 0 0 0 3 3h3.586l-.293.293a1 1 0 0 0 1.414 1.414l2-2a1 1 0 0 0 0-1.414l-2-2a1 1 0 0 0-1.414 1.414l.293.293H10a1 1 0 0 1-1-1v-1z" clip-rule="evenodd"/>`
	repeatCircleLineInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 8l2-2M8 8l2 2M8 8h6a2 2 0 0 1 2 2v1m0 5l-2 2m2-2l-2-2m2 2h-6a2 2 0 0 1-2-2v-1"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`
	replyInnerSVG                     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11 5a1 1 0 0 0-1.707-.707l-7 7a1 1 0 0 0 0 1.414l7 7A1 1 0 0 0 11 19v-3.025c1.691-.011 3.83.133 5.633.583c1.088.27 1.973.633 2.565 1.076c.567.424.802.864.802 1.366a1 1 0 1 0 2 0c0-1.925-.598-4.66-2.42-6.937c-1.719-2.15-4.462-3.805-8.58-4.036V5z" fill="currentColor"/></g>`
	replyAllInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M13.383 4.076A1 1 0 0 1 14 5v3.028c1.855.106 3.365.513 4.586 1.158a7.46 7.46 0 0 1 3.087 3.014C22.966 14.518 23 17.27 23 19a1 1 0 1 1-2 0c0-.622-.218-1.081-.597-1.46c-.404-.403-1.021-.741-1.832-1c-1.344-.428-3.02-.574-4.571-.565V19a1 1 0 0 1-1.707.707l-7-7a1 1 0 0 1 0-1.414l7-7a1 1 0 0 1 1.09-.217zM10 5v.586l-5.707 5.707a.997.997 0 0 0 0 1.414L10 18.414V19a1 1 0 0 1-1.707.707l-7-7a1 1 0 0 1 0-1.414l7-7A1 1 0 0 1 10 5zm9 15h-.008h.016H19z" fill="currentColor"/></g>`
	replyAllLineInnerSVG              = `<g fill="currentColor"><path d="M9.383 4.076A1 1 0 0 1 10 5v1.586l2.293-2.293A1 1 0 0 1 14 5v3.029c1.859.109 3.371.526 4.593 1.184a7.556 7.556 0 0 1 3.082 3.054C22.965 14.602 23 17.348 23 19a1 1 0 1 1-2 0c0-.621-.218-1.081-.597-1.46c-.404-.403-1.021-.741-1.832-1c-1.344-.428-3.02-.574-4.571-.565V19a1 1 0 0 1-1.707.707L10 17.414V19a1 1 0 0 1-1.707.707l-7-7a1 1 0 0 1 0-1.414l7-7a1 1 0 0 1 1.09-.217zM8 15.414l-2.707-2.707a1 1 0 0 1 0-1.414L8 8.586V7.414L3.414 12L8 16.586v-1.172zm12.677-.14a7.768 7.768 0 0 0-.753-2.04a5.556 5.556 0 0 0-2.28-2.26C16.566 10.393 15.075 10 13 10a1 1 0 0 1-1-1V7.414L7.414 12L12 16.586V15a1 1 0 0 1 .955-.999c1.879-.085 4.266.01 6.224.634a7.97 7.97 0 0 1 1.498.64z"/></g>`
	replyLineInnerSVG                 = `<g fill="currentColor"><path d="M10.383 4.076A1 1 0 0 1 11 5v3.028c4.117.23 6.861 1.885 8.58 4.035C21.403 14.339 22 17.075 22 19a1 1 0 1 1-2 0c0-.502-.235-.942-.802-1.366c-.592-.443-1.477-.805-2.565-1.076c-1.804-.45-3.941-.594-5.633-.583V19a1 1 0 0 1-1.707.707l-7-7a1 1 0 0 1 0-1.414l7-7a1 1 0 0 1 1.09-.217zm8.843 11.267a8.762 8.762 0 0 0-1.207-2.03C16.574 11.505 14.122 10 10 10a1 1 0 0 1-1-1V7.414L4.414 12L9 16.586V15a1 1 0 0 1 .955-.999c1.888-.086 4.743.014 7.162.616a11.6 11.6 0 0 1 2.11.726z"/></g>`
	restrictedInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 5.636a9 9 0 1 0 12.728 12.728M5.636 5.636a9 9 0 1 1 12.728 12.728M5.636 5.636L12 12l6.364 6.364"/>`
	restrictedLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 5.636a9 9 0 1 0 12.728 12.728M5.636 5.636a9 9 0 1 1 12.728 12.728M5.636 5.636L12 12l6.364 6.364"/>`
	rewindInnerSVG                    = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2.933 10.8l6.667-5c.989-.742 2.4-.036 2.4 1.2v3l5.6-4.2c.989-.742 2.4-.036 2.4 1.2v10c0 1.236-1.411 1.942-2.4 1.2L12 14v3c0 1.236-1.411 1.942-2.4 1.2l-6.667-5c-.8-.6-.8-1.8 0-2.4z" fill="currentColor"/></g>`
	rewindLineInnerSVG                = `<g fill="currentColor"><path d="M10 8v8l-5.333-4L10 8zm2-1c0-1.236-1.411-1.942-2.4-1.2l-6.667 5c-.8.6-.8 1.8 0 2.4l6.667 5c.989.742 2.4.036 2.4-1.2v-3l5.6 4.2c.989.742 2.4.036 2.4-1.2V7c0-1.236-1.411-1.942-2.4-1.2L12 10V7zm.667 5L18 8v8l-5.333-4z"/></g>`
	robotInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M12 3a2 2 0 0 0-1 3.732V8H8c-3.2 0-4 2.667-4 4v7c0 .667.4 2 2 2h1v-4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v4h1c1.6 0 2-1.333 2-2v-7c0-3.2-2.667-4-4-4h-3V6.732A2 2 0 0 0 12 3zm3 18v-3h-2v3h2zm-4 0v-3H9v3h2zm10-3v-5c.667 0 2 .4 2 2v1c0 .667-.4 2-2 2zM3 13v5c-1.6 0-2-1.333-2-2v-1c0-1.6 1.333-2 2-2zm6-1a1 1 0 1 0 0 2h.001a1 1 0 1 0 0-2H9zm5 1a1 1 0 0 1 1-1h.001a1 1 0 1 1 0 2H15a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`
	robotLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 20H5v-2m3 2v-4h4m-4 4h4m4 0h3v-2m-3 2v-4h-4m4 4h-4m0-4v4m0-11H9a4 4 0 0 0-4 4v0m7-4h3a4 4 0 0 1 4 4v0m-7-4V5m7 8h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1m0-5v5M5 14.5v2M5 13H4a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h1m0-5v5m4-5h.001M15 13h.001"/><circle cx="12" cy="5" r="1"/></g>`
	rocketThreeStartInnerSVG          = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.95 16.264s-1.703 2.54-.707 3.535c.995.996 3.535-.707 3.535-.707"/><path fill="currentColor" fill-rule="evenodd" d="M20.506 3.536a1 1 0 0 1 .268.928l-.317 1.402a9 9 0 0 1-2.414 4.375l-4.644 4.644c1.027 1.272 1.36 2.48 1.1 3.632c-.271 1.2-1.16 2.086-1.712 2.637l-.06.06a1 1 0 0 1-1.564-.193L9.17 17.696a1 1 0 0 0-.15-.192l-2.57-2.568l-.76-.456l3.459-3.843a.343.343 0 0 0 .007.005L13.8 6a9 9 0 0 1 4.376-2.414l1.402-.318a1 1 0 0 1 .928.269zM8.322 10.062c-.969-.565-1.9-.722-2.797-.52c-1.2.272-2.086 1.16-2.637 1.713l-.06.059a1 1 0 0 0 .193 1.564l1.796 1.078l3.505-3.894z" clip-rule="evenodd"/></g>`
	rocketThreeStartLineInnerSVG      = `<g fill="none"><g clip-path="url(#majesticonsRocket3StartLine0)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.9 16.97l7.436-7.436a8 8 0 0 0 2.145-3.89l.318-1.401l-1.402.317a8 8 0 0 0-3.89 2.146L9.192 12.02m.707 4.95l2.122 3.535c1.178-1.178 2.828-2.828 0-5.657L9.899 16.97zm0 0l-2.828-2.829m0 0L3.536 12.02c1.178-1.179 2.828-2.829 5.656 0m-2.12 2.121l2.12-2.121M4.95 16.263s-1.703 2.54-.707 3.536c.995.996 3.535-.707 3.535-.707"/></g><defs><clipPath id="majesticonsRocket3StartLine0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	rssInnerSVG                       = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5zm0 3.962A1 1 0 0 1 6.039 4c2.32.089 5.775.883 8.677 3.004C17.664 9.16 20 12.66 20 18a1 1 0 1 1-2 0c0-4.66-1.997-7.577-4.465-9.38C11.021 6.783 7.975 6.077 5.962 6A1 1 0 0 1 5 4.961zm.003 5.967a1 1 0 0 1 1.068-.926c1.314.093 3.257.59 4.899 1.806C12.654 13.055 14 15.049 14 18a1 1 0 1 1-2 0c0-2.248-.987-3.671-2.22-4.584c-1.274-.943-2.831-1.346-3.851-1.418a1 1 0 0 1-.926-1.07zM5 18a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1z" fill="currentColor"/></g>`
	rssLineInnerSVG                   = `<g fill="currentColor"><path d="M5 4.962A1 1 0 0 1 6.039 4c2.32.089 5.775.883 8.677 3.004C17.664 9.16 20 12.66 20 18a1 1 0 1 1-2 0c0-4.66-1.997-7.577-4.465-9.38C11.021 6.783 7.975 6.077 5.962 6A1 1 0 0 1 5 4.961zm.003 5.967a1 1 0 0 1 1.068-.926c1.314.093 3.257.59 4.899 1.806C12.654 13.055 14 15.049 14 18a1 1 0 1 1-2 0c0-2.248-.987-3.671-2.22-4.584c-1.274-.943-2.831-1.346-3.851-1.418a1 1 0 0 1-.926-1.07zM5 18a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1z"/></g>`
	rubelCircleInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm9-6a1 1 0 0 0-1 1v3H8a1 1 0 1 0 0 2h1v1H8a1 1 0 1 0 0 2h1v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-1h3a3 3 0 1 0 0-6h-4zm4 4h-3V8h3a1 1 0 1 1 0 2z" clip-rule="evenodd"/>`
	rubelCircleLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M10 11h4a2 2 0 0 0 2-2v0a2 2 0 0 0-2-2h-4v4zm0 0v3m0-3H8m2 6v-3m0 0H8m2 0h3"/></g>`
	rulerTwoInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M12.879 3.707a3 3 0 0 1 4.242 0l3.172 3.172a3 3 0 0 1 0 4.242l-9.172 9.172a3 3 0 0 1-4.242 0L3.707 17.12a3 3 0 0 1 0-4.242L4.586 12l1.707 1.707a1 1 0 0 0 1.414-1.414L6 10.586L7.586 9l2.707 2.707a1 1 0 0 0 1.414-1.414L9 7.586L10.586 6l1.707 1.707a1 1 0 1 0 1.414-1.414L12 4.586l.879-.879z" clip-rule="evenodd"/>`
	rulerTwoLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 12l-1.586 1.586a2 2 0 0 0 0 2.828l3.172 3.172a2 2 0 0 0 2.828 0l9.172-9.172a2 2 0 0 0 0-2.828l-3.172-3.172a2 2 0 0 0-2.828 0L12 6m-6 6l2 2m-2-2l3-3m3-3l2 2m-2-2L9 9m0 0l3 3"/>`
	rupeeCircleInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm7-6a1 1 0 0 0 0 2h3c.34 0 .872.11 1.29.412c.19.136.372.321.505.588H7.997a1 1 0 1 0 0 2h4.798a1.58 1.58 0 0 1-.504.588A2.352 2.352 0 0 1 11 12H7.997a1 1 0 0 0-.625 1.781l5.003 4a1 1 0 1 0 1.25-1.562L10.848 14h.15c.661 0 1.629-.19 2.46-.789A3.621 3.621 0 0 0 14.896 11H16a1 1 0 1 0 0-2h-1.104a3.81 3.81 0 0 0-.367-1H16a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`
	rupeeCircleLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M8 7h3m5 0h-5m5 3h-2m-6.003 0H14m-3-3c1 0 3 .6 3 3m-1 7l-5.003-4H11c1 0 3-.6 3-3"/></g>`
	saveInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M4 3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V8a1 1 0 0 0-.293-.707l-4-4A1 1 0 0 0 16 3H4zm6 11c0-1.6 1.333-2 2-2s2 .4 2 2s-1.333 2-2 2s-2-.4-2-2zm4-9H8v3h6V5z" clip-rule="evenodd"/>`
	saveAltInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 6a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H5zm7-3a1 1 0 0 1 1 1v7.586l1.293-1.293a1 1 0 1 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L11 11.586V4a1 1 0 0 1 1-1z" fill="currentColor"/></g>`
	saveAltLineInnerSVG               = `<g fill="currentColor"><path d="M12 3a1 1 0 0 1 1 1v7.586l1.293-1.293a1 1 0 1 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L11 11.586V4a1 1 0 0 1 1-1zM5 8a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-2a1 1 0 1 1 0-2h2a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h2a1 1 0 0 1 0 2H5z"/></g>`
	saveAsInnerSVG                    = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 8a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V8zm-2 .17A3.001 3.001 0 0 0 2 11v7a3 3 0 0 0 3 3h10a3.001 3.001 0 0 0 2.83-2H9a5 5 0 0 1-5-5V8.17zM14 3a1 1 0 0 1 1 1v5.586l1.293-1.293a1 1 0 1 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L13 9.586V4a1 1 0 0 1 1-1z" fill="currentColor"/></g>`
	saveAsLineInnerSVG                = `<g fill="currentColor"><path d="M14 3a1 1 0 0 1 1 1v5.586l1.293-1.293a1 1 0 1 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L13 9.586V4a1 1 0 0 1 1-1zM9 7a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1h-1a1 1 0 1 1 0-2h1a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3h-1v1a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h1V8a3 3 0 0 1 3-3h1a1 1 0 1 1 0 2H9zm-3 4H5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-1H9a3 3 0 0 1-3-3v-3z"/></g>`
	saveLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4H4v16h16V8l-4-4h-2M8 4v4h6V4M8 4h6m-2 8c-.667 0-2 .4-2 2s1.333 2 2 2s2-.4 2-2s-1.333-2-2-2z"/>`
	scaleLightInnerSVG                = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 6l2 5m0 0h3a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h3z"/><path fill="currentColor" fill-rule="evenodd" d="M2 7a3 3 0 0 1 3-3h2a1 1 0 0 1 1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V5a1 1 0 0 1 1-1h2a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V7zm8 9a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4z" clip-rule="evenodd"/></g>`
	scaleLightLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 6l2 5m0 0h3a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h3z"/><path d="M17 5h2a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h2m3 12h4"/></g>`
	scanFingerprintInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19v-8m-3 7v-7c0-1 .6-3 3-3s3 2 3 3v6m-9-3v-3c0-2 1.2-6 6-6s6 4 6 6m0 4v-1M6.001 17H6M7 3H5a2 2 0 0 0-2 2v2m0 10v2a2 2 0 0 0 2 2h2m10 0h2a2 2 0 0 0 2-2v-2m0-10V5a2 2 0 0 0-2-2h-2"/>`
	scanFingerprintLineInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19v-8m-3 7v-7c0-1 .6-3 3-3s3 2 3 3v6m-9-3v-3c0-2 1.2-6 6-6s6 4 6 6m0 4v-1M6.001 17H6M7 3H5a2 2 0 0 0-2 2v2m0 10v2a2 2 0 0 0 2 2h2m10 0h2a2 2 0 0 0 2-2v-2m0-10V5a2 2 0 0 0-2-2h-2"/>`
	scanUserInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3H5a2 2 0 0 0-2 2v2m0 10v2a2 2 0 0 0 2 2h2m10 0h2a2 2 0 0 0 2-2v-2m0-10V5a2 2 0 0 0-2-2h-2"/><circle cx="12" cy="9" r="3" fill="currentColor"/><path fill="currentColor" d="M12 12c-2.761 0-5 1.79-5 4h10c0-2.21-2.239-4-5-4z"/></g>`
	scanUserLineInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3H5a2 2 0 0 0-2 2v2m0 10v2a2 2 0 0 0 2 2h2m10 0h2a2 2 0 0 0 2-2v-2m0-10V5a2 2 0 0 0-2-2h-2"/><circle cx="12" cy="9" r="3"/><path d="M17 16c0-2.21-2.239-4-5-4s-5 1.79-5 4"/></g>`
	scannerInnerSVG                   = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M5 9a1 1 0 0 0-.962.725L2.246 16h19.508l-1.793-6.275A1 1 0 0 0 19 9H5zm16.78 9H2.22l.431 1.728A3 3 0 0 0 5.561 22H18.44a3 3 0 0 0 2.91-2.272L21.78 18zM6.968 11.83A1.5 1.5 0 0 1 8.31 11h7.382a1.5 1.5 0 0 1 0 3H8.309a1.5 1.5 0 0 1-1.342-2.17z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.728 5.55A2 2 0 0 1 5.651 3H18.35a2 2 0 0 1 1.923 2.55L19 10H5L3.728 5.55z"/></g>`
	scannerLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 17l-2-7H5l-2 7m18 0H3m18 0l-.621 2.485A2 2 0 0 1 18.439 21H5.561a2 2 0 0 1-1.94-1.515L3 17m5.309-3h7.382a.5.5 0 0 0 .447-.724v0a.5.5 0 0 0-.447-.276H8.309a.5.5 0 0 0-.447.276v0a.5.5 0 0 0 .447.724zM3.728 5.55A2 2 0 0 1 5.651 3H18.35a2 2 0 0 1 1.923 2.55L19 10H5L3.728 5.55z"/>`
	scissorsInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3 7a4 4 0 1 1 7.446 2.032L12 10.586l6.293-6.293a1 1 0 1 1 1.414 1.414l-9.26 9.261a4 4 0 1 1-1.414-1.414L10.585 12l-1.554-1.554A4 4 0 0 1 3 7zm11.707 6.293a1 1 0 0 0-1.414 1.414l5 5a1 1 0 0 0 1.414-1.414l-5-5z" fill="currentColor"/></g>`
	scissorsLineInnerSVG              = `<g fill="currentColor"><path d="M7 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4zM3 7a4 4 0 1 1 7.446 2.032L12 10.586l6.293-6.293a1 1 0 1 1 1.414 1.414l-7 7l-2.26 2.261a4 4 0 1 1-1.414-1.414L10.585 12l-1.554-1.554A4 4 0 0 1 3 7zm4 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm7.707-1.707a1 1 0 0 0-1.414 1.414l5 5a1 1 0 0 0 1.414-1.414l-5-5z"/></g>`
	scooterInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M8.126 4H6a1 1 0 0 0 0 2h2.126a4.007 4.007 0 0 1 0-2zm.32 2.836C5.725 8.314 5 11.355 5 13v5a3 3 0 0 0 3 3h.028A3.477 3.477 0 0 1 8 20.556V15a4 4 0 0 1 8 0v5.556c0 .15-.01.299-.028.444H16a3 3 0 0 0 3-3v-5c0-1.645-.726-4.686-3.445-6.164a4 4 0 0 1-7.11 0zM15.873 6H18a1 1 0 1 0 0-2h-2.126a4.01 4.01 0 0 1 0 2zM12 13a2 2 0 0 0-2 2v5.556c0 .797.647 1.444 1.444 1.444h1.112c.797 0 1.444-.647 1.444-1.444V15a2 2 0 0 0-2-2zm-2-8a2 2 0 1 1 4 0a2 2 0 0 1-4 0z" clip-rule="evenodd"/>`
	scooterLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5a3 3 0 1 0-6 0m6 0h3m-3 0c0 .903-.399 1.713-1.03 2.263M9 5H6m3 0c0 .903.399 1.713 1.03 2.263M14 20h2a2 2 0 0 0 2-2v-5c0-1.692-.859-4.816-4.03-5.737M14 20v0a2 2 0 0 1-2 2v0a2 2 0 0 1-2-2v0m4 0v-5a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v5m0 0H8a2 2 0 0 1-2-2v-5c0-1.692.859-4.816 4.03-5.737m3.94 0A2.988 2.988 0 0 1 12 8a2.988 2.988 0 0 1-1.97-.737"/>`
	scriptPrescriptionInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M6 3c1 0 3 .6 3 3S7 9 6 9h12M6 3h12c1 0 3 .6 3 3s-2 3-3 3M6 3c-1 0-3 .6-3 3v13.002C3 20.106 3.895 21 5 21h5.5M18 9v1M7 13h3m-3 4h3m4 2v-3m2 0h1.5a1.5 1.5 0 0 0 1.5-1.5v0a1.5 1.5 0 0 0-1.5-1.5H14v3m2 0l3 3m-3-3h-2m7 5l-2-2m0 0l2-2m-2 2l-2 2"/><path fill="currentColor" d="M9 6c0 2.4-2 3-3 3h12c1 0 3-.6 3-3s-2-3-3-3H6c1 0 3 .6 3 3z"/></g>`
	scriptPrescriptionLineInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3c1 0 3 .6 3 3S7 9 6 9h12M6 3h12c1 0 3 .6 3 3s-2 3-3 3M6 3c-1 0-3 .6-3 3v13.002C3 20.106 3.895 21 5 21h5.5M18 9v1M7 13h3m-3 4h3m4 2v-3m2 0h1.5a1.5 1.5 0 0 0 1.5-1.5v0a1.5 1.5 0 0 0-1.5-1.5H14v3m2 0l3 3m-3-3h-2m7 5l-2-2m0 0l2-2m-2 2l-2 2"/>`
	scrollInnerSVG                    = `<g fill="none"><path fill="currentColor" d="M3 16h8c0 1.333.8 4 4 4a3 3 0 0 0 3-3V4h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1v7a3 3 0 0 1-3 3H5a2 2 0 0 1-2-2v-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20c-3.2 0-4-2.667-4-4H3v2a2 2 0 0 0 2 2h10zm0 0a3 3 0 0 0 3-3v-7m0-6H7a2 2 0 0 0-2 2v9.5M18 4h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1m0-6v6"/></g>`
	scrollLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20c-3.2 0-4-2.667-4-4H3v2a2 2 0 0 0 2 2h10zm0 0a3 3 0 0 0 3-3v-7m0-6H7a2 2 0 0 0-2 2v10M18 4h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1m0-6v6"/>`
	scrollTextInnerSVG                = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 8H9m2 4H9"/><path fill="currentColor" d="M3 16h8c0 1.333.8 4 4 4a3 3 0 0 0 3-3V4h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1v7a3 3 0 0 1-3 3H5a2 2 0 0 1-2-2v-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20c-3.2 0-4-2.667-4-4H3v2a2 2 0 0 0 2 2h10zm0 0a3 3 0 0 0 3-3v-7m0-6H7a2 2 0 0 0-2 2v9.5M18 4h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1m0-6v6"/></g>`
	scrollTextLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20c-3.2 0-4-2.667-4-4H3v2a2 2 0 0 0 2 2h10zm0 0a3 3 0 0 0 3-3v-7m0-6H7a2 2 0 0 0-2 2v10M18 4h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1m0-6v6m-4-2H9m2 4H9"/>`
	searchInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M5 11a6 6 0 1 1 12 0a6 6 0 0 1-12 0zm6-8a8 8 0 1 0 4.906 14.32l3.387 3.387a1 1 0 0 0 1.414-1.414l-3.387-3.387A8 8 0 0 0 11 3zm0 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8z" clip-rule="evenodd"/>`
	searchCircleInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm-1 5a4 4 0 1 0 2.032 7.446l1.76 1.761a1 1 0 0 0 1.415-1.414l-1.761-1.761A4 4 0 0 0 11 7zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4z" fill="currentColor"/></g>`
	searchCircleLineInnerSVG          = `<g fill="currentColor"><path d="M11 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm-4 2a4 4 0 1 1 7.446 2.032l1.761 1.76a1 1 0 0 1-1.414 1.415l-1.761-1.761A4 4 0 0 1 7 11z"/><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12z"/></g>`
	searchLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 20l-4.05-4.05m0 0a7 7 0 1 0-9.9-9.9a7 7 0 0 0 9.9 9.9z"/>`
	searchMinusInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M3 11a8 8 0 1 1 14.32 4.906l3.387 3.387a1 1 0 0 1-1.414 1.414l-3.387-3.387A8 8 0 0 1 3 11zm5-1a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`
	searchMinusLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 20l-4.05-4.05m0 0a7 7 0 1 0-9.9-9.9a7 7 0 0 0 9.9 9.9zM8 11h6"/>`
	searchPlusInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M3 11a8 8 0 1 1 14.32 4.906l3.387 3.387a1 1 0 0 1-1.414 1.414l-3.387-3.387A8 8 0 0 1 3 11zm9-3a1 1 0 1 0-2 0v2H8a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V8z" clip-rule="evenodd"/>`
	searchPlusLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 20l-4.05-4.05m0 0a7 7 0 1 0-9.9-9.9a7 7 0 0 0 9.9 9.9zM8 11h3m3 0h-3m0 0V8m0 3v3"/>`
	searchTextInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M4 6a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm15-2.5a3.5 3.5 0 0 1-.42 1.665l1.127 1.128a1 1 0 0 1-1.414 1.414l-1.128-1.128A3.5 3.5 0 1 1 19 15.5z" fill="currentColor"/></g>`
	searchTextLineInnerSVG            = `<g fill="currentColor"><path d="M4 6a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm11.5-4a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zM12 15.5a3.5 3.5 0 1 1 6.58 1.665l1.127 1.128a1 1 0 0 1-1.414 1.414l-1.128-1.128A3.5 3.5 0 0 1 12 15.5z"/></g>`
	selectorInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12.707 4.293a1 1 0 0 0-1.414 0l-4 4A1 1 0 0 0 8 10h8a1 1 0 0 0 .707-1.707l-4-4zM8 14a1 1 0 0 0-.707 1.707l4 4a1 1 0 0 0 1.414 0l4-4A1 1 0 0 0 16 14H8z" fill="currentColor"/></g>`
	selectorLineInnerSVG              = `<g fill="currentColor"><path d="M11.293 4.293a1 1 0 0 1 1.414 0l4 4a1 1 0 0 1-1.414 1.414L12 6.414L8.707 9.707a1 1 0 0 1-1.414-1.414l4-4zm-4 10a1 1 0 0 1 1.414 0L12 17.586l3.293-3.293a1 1 0 0 1 1.414 1.414l-4 4a1 1 0 0 1-1.414 0l-4-4a1 1 0 0 1 0-1.414z"/></g>`
	sendInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M2.345 2.245a1 1 0 0 1 1.102-.14l18 9a1 1 0 0 1 0 1.79l-18 9a1 1 0 0 1-1.396-1.211L4.613 13H10a1 1 0 1 0 0-2H4.613L2.05 3.316a1 1 0 0 1 .294-1.071z" clip-rule="evenodd"/>`
	sendLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 12l-3 9l18-9L3 3l3 9zm0 0h6"/>`
	serverInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v3H2V8zm0 5v3a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-3H2zm4-6a1 1 0 0 0 0 2h.01a1 1 0 0 0 0-2H6zm-1 9a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1zm4-9a1 1 0 0 0 0 2h.01a1 1 0 0 0 0-2H9zm-1 9a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`
	serverLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12V8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v4M3 12h18M3 12v4a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4M6 9h.01M6 15h.01M9 9h.01M9 15h.01"/>`
	settingsCogInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M9.024 2.783A1 1 0 0 1 10 2h4a1 1 0 0 1 .976.783l.44 1.981c.4.19.781.41 1.14.66l1.938-.61a1 1 0 0 1 1.166.454l2 3.464a1 1 0 0 1-.19 1.237l-1.497 1.373a8.1 8.1 0 0 1 0 1.316l1.497 1.373a1 1 0 0 1 .19 1.237l-2 3.464a1 1 0 0 1-1.166.454l-1.937-.61c-.36.25-.741.47-1.14.66l-.44 1.98A1 1 0 0 1 14 22h-4a1 1 0 0 1-.976-.783l-.44-1.981c-.4-.19-.781-.41-1.14-.66l-1.938.61a1 1 0 0 1-1.166-.454l-2-3.464a1 1 0 0 1 .19-1.237l1.497-1.373a8.097 8.097 0 0 1 0-1.316L2.53 9.97a1 1 0 0 1-.19-1.237l2-3.464a1 1 0 0 1 1.166-.454l1.937.61c.36-.25.741-.47 1.14-.66l.44-1.98zM12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6z" clip-rule="evenodd"/>`
	settingsCogCheckInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M9.024 2.783A1 1 0 0 1 10 2h4a1 1 0 0 1 .976.783l.44 1.981c.4.19.781.41 1.14.66l1.938-.61a1 1 0 0 1 1.166.454l2 3.464a1 1 0 0 1-.19 1.237l-1.497 1.373a8.1 8.1 0 0 1 0 1.316l1.497 1.373a1 1 0 0 1 .19 1.237l-2 3.464a1 1 0 0 1-1.166.454l-1.937-.61c-.36.25-.741.47-1.14.66l-.44 1.98A1 1 0 0 1 14 22h-4a1 1 0 0 1-.976-.783l-.44-1.981c-.4-.19-.781-.41-1.14-.66l-1.938.61a1 1 0 0 1-1.166-.454l-2-3.464a1 1 0 0 1 .19-1.237l1.497-1.373a8.097 8.097 0 0 1 0-1.316L2.53 9.97a1 1 0 0 1-.19-1.237l2-3.464a1 1 0 0 1 1.166-.454l1.937.61c.36-.25.741-.47 1.14-.66l.44-1.98zm6.683 7.924a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4z" clip-rule="evenodd"/>`
	settingsCogCheckLineInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 21h-4l-.551-2.48a6.991 6.991 0 0 1-1.819-1.05l-2.424.763l-2-3.464l1.872-1.718a7.055 7.055 0 0 1 0-2.1L3.206 9.232l2-3.464l2.424.763A6.992 6.992 0 0 1 9.45 5.48L10 3h4l.551 2.48a6.994 6.994 0 0 1 1.819 1.05l2.424-.763l2 3.464l-1.872 1.718a7.05 7.05 0 0 1 0 2.1l1.872 1.718l-2 3.464l-2.424-.763a6.992 6.992 0 0 1-1.819 1.052L14 21z"/><path d="m9 12l2 2l4-4"/></g>`
	settingsCogLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M14 21h-4l-.551-2.48a6.991 6.991 0 0 1-1.819-1.05l-2.424.763l-2-3.464l1.872-1.718a7.055 7.055 0 0 1 0-2.1L3.206 9.232l2-3.464l2.424.763A6.992 6.992 0 0 1 9.45 5.48L10 3h4l.551 2.48a6.992 6.992 0 0 1 1.819 1.05l2.424-.763l2 3.464l-1.872 1.718a7.05 7.05 0 0 1 0 2.1l1.872 1.718l-2 3.464l-2.424-.763a6.99 6.99 0 0 1-1.819 1.052L14 21z"/><circle cx="12" cy="12" r="3"/></g>`
	settingsCogPlusInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M9.024 2.783A1 1 0 0 1 10 2h4a1 1 0 0 1 .976.783l.44 1.981c.4.19.781.41 1.14.66l1.938-.61a1 1 0 0 1 1.166.454l2 3.464a1 1 0 0 1-.19 1.237l-1.497 1.373a8.1 8.1 0 0 1 0 1.316l1.497 1.373a1 1 0 0 1 .19 1.237l-2 3.464a1 1 0 0 1-1.166.454l-1.937-.61c-.36.25-.741.47-1.14.66l-.44 1.98A1 1 0 0 1 14 22h-4a1 1 0 0 1-.976-.783l-.44-1.981c-.4-.19-.781-.41-1.14-.66l-1.938.61a1 1 0 0 1-1.166-.454l-2-3.464a1 1 0 0 1 .19-1.237l1.497-1.373a8.097 8.097 0 0 1 0-1.316L2.53 9.97a1 1 0 0 1-.19-1.237l2-3.464a1 1 0 0 1 1.166-.454l1.937.61c.36-.25.741-.47 1.14-.66l.44-1.98zM11 9a1 1 0 1 1 2 0v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9a1 1 0 1 1 0-2h2V9z" clip-rule="evenodd"/>`
	settingsCogPlusLineInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 21h-4l-.551-2.48a6.991 6.991 0 0 1-1.819-1.05l-2.424.763l-2-3.464l1.872-1.718a7.055 7.055 0 0 1 0-2.1L3.206 9.232l2-3.464l2.424.763A6.992 6.992 0 0 1 9.45 5.48L10 3h4l.551 2.48a6.994 6.994 0 0 1 1.819 1.05l2.424-.763l2 3.464l-1.872 1.718a7.05 7.05 0 0 1 0 2.1l1.872 1.718l-2 3.464l-2.424-.763a6.992 6.992 0 0 1-1.819 1.052L14 21zM12 9v3m0 3v-3m0 0H9m3 0h3"/>`
	shareInnerSVG                     = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 12l-6.4-7v3.5C10.4 8.5 4 10.6 4 19c0-1.167 1.92-3.5 9.6-3.5V19l6.4-7z"/>`
	shareCircleInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm11-5a1 1 0 0 1 1.707-.707l4 4a1 1 0 0 1 0 1.414l-4 4A1 1 0 0 1 12 15v-.96c-1.631.132-2.625.584-3.205 1.014a2.534 2.534 0 0 0-.675.71a1.02 1.02 0 0 0-.117.255a.03.03 0 0 1-.003.008v-.002A1 1 0 0 1 6 16c0-2.706 1.137-4.715 2.573-6.036C9.635 8.987 10.89 8.361 12 8.117V7zm1 5a1 1 0 0 1 .937.65L15.586 11l-1.65-1.65A1 1 0 0 1 13 10c-.704 0-1.971.422-3.073 1.436a5.91 5.91 0 0 0-1.062 1.296C9.903 12.29 11.257 12 13 12z" clip-rule="evenodd"/>`
	shareCircleLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 11l-4-4v2c-2 0-6 2.2-6 7c0-.667 1.2-3 6-3v2l4-4z"/><circle cx="12" cy="12" r="10"/></g>`
	shareLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 12l-6.4-7v3.5C10.4 8.5 4 10.6 4 19c0-1.167 1.92-3.5 9.6-3.5V19l6.4-7z"/>`
	shieldInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M12.707 2.293a1 1 0 0 0-1.414 0c-.81.81-1.973 1.29-3.183 1.54c-1.202.248-2.347.246-3 .173A1 1 0 0 0 4 5v9c0 1.3.568 2.449 1.304 3.395c.738.948 1.697 1.763 2.615 2.419a21.859 21.859 0 0 0 3.66 2.093l.018.008l.006.003h.003a1 1 0 0 0 .788.001L12 21c.394.92.395.919.395.919l.002-.001l.005-.003l.02-.008l.065-.029a21.482 21.482 0 0 0 1.07-.523a21.95 21.95 0 0 0 2.524-1.541c.918-.656 1.877-1.47 2.615-2.419C19.432 16.45 20 15.3 20 14V5a1 1 0 0 0-1.11-.994c-.653.073-1.798.075-3-.173c-1.21-.25-2.373-.73-3.183-1.54zM12 21l-.394.919L12 21z" clip-rule="evenodd"/>`
	shieldCheckInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 1.932l9 3.375V10c0 4.118-1.62 7.113-3.535 9.074a11.407 11.407 0 0 1-2.915 2.175c-.913.468-1.826.751-2.55.751c-.724 0-1.637-.283-2.55-.75a11.41 11.41 0 0 1-2.915-2.176C4.619 17.113 3 14.118 3 10V5.307l9-3.375zm3.707 8.775a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4z" fill="currentColor"/></g>`
	shieldCheckLineInnerSVG           = `<g fill="currentColor"><path d="M12 1.932l9 3.375V10c0 4.118-1.62 7.113-3.535 9.074a11.407 11.407 0 0 1-2.915 2.175c-.913.468-1.826.751-2.55.751c-.724 0-1.637-.283-2.55-.75a11.41 11.41 0 0 1-2.915-2.176C4.619 17.113 3 14.118 3 10V5.307l9-3.375zM5 6.693V10c0 3.549 1.38 6.054 2.965 7.676a9.411 9.411 0 0 0 2.397 1.793c.775.397 1.362.531 1.638.531c.276 0 .863-.134 1.638-.53a9.412 9.412 0 0 0 2.397-1.794C17.619 16.054 19 13.55 19 10V6.693l-7-2.625l-7 2.625z"/><path d="M15.707 9.293a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 1 1 1.414-1.414L11 12.586l3.293-3.293a1 1 0 0 1 1.414 0z"/></g>`
	shieldExclamationInnerSVG         = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 1.932l9 3.375V10c0 4.118-1.62 7.113-3.535 9.074a11.407 11.407 0 0 1-2.915 2.175c-.913.468-1.826.751-2.55.751c-.724 0-1.637-.283-2.55-.75a11.41 11.41 0 0 1-2.915-2.176C4.619 17.113 3 14.118 3 10V5.307l9-3.375zM12 7a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1zm1 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0z" fill="currentColor"/></g>`
	shieldExclamationLineInnerSVG     = `<g fill="currentColor"><path d="M12 1.932l9 3.375V10c0 4.118-1.62 7.113-3.535 9.074a11.407 11.407 0 0 1-2.915 2.175c-.913.468-1.826.751-2.55.751c-.724 0-1.637-.283-2.55-.75a11.41 11.41 0 0 1-2.915-2.176C4.619 17.113 3 14.118 3 10V5.307l9-3.375zM5 6.693V10c0 3.549 1.38 6.054 2.965 7.676a9.411 9.411 0 0 0 2.397 1.793c.775.397 1.362.531 1.638.531c.276 0 .863-.134 1.638-.53a9.412 9.412 0 0 0 2.397-1.794C17.619 16.054 19 13.55 19 10V6.693l-7-2.625l-7 2.625z"/><path d="M12 7a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1zm-1 8a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H12a1 1 0 0 1-1-1z"/></g>`
	shieldLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14c0 4-7 7-7 7s-7-3-7-7V5c1.5.167 5 0 7-2c2 2 5.5 2.167 7 2v9z"/>`
	shieldOffInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18"/><path fill="currentColor" fill-rule="evenodd" d="M4.013 4.841A1 1 0 0 0 4 5v9c0 1.3.568 2.449 1.304 3.395c.738.948 1.697 1.763 2.615 2.419a21.859 21.859 0 0 0 3.66 2.093l.018.008l.006.003h.003a1 1 0 0 0 .788.001L12 21c.394.92.395.919.395.919l.002-.001l.005-.003l.02-.008l.065-.029a21.482 21.482 0 0 0 1.07-.523a21.95 21.95 0 0 0 2.524-1.541c.536-.383 1.086-.82 1.598-1.306L4.013 4.84zm15.864 10.207A4.57 4.57 0 0 0 20 14V5a1 1 0 0 0-1.11-.994c-.653.073-1.798.075-3-.173c-1.21-.25-2.373-.73-3.183-1.54a1 1 0 0 0-1.414 0c-.71.709-1.69 1.166-2.735 1.436l11.319 11.32zm-8.271 6.87L12 21l-.394.919z" clip-rule="evenodd"/></g>`
	shieldOffLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14V5c-1.5.167-5 0-7-2c-.571.571-1.265.993-2 1.3M5 5v9c0 4 7 7 7 7s3.204-1.373 5.277-3.5M3 3l18 18"/>`
	shieldPlusInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M11.293 2.293a1 1 0 0 1 1.414 0c.81.81 1.973 1.29 3.183 1.54c1.202.248 2.347.246 3 .173A1 1 0 0 1 20 5v9c0 1.3-.568 2.449-1.304 3.395c-.738.948-1.697 1.763-2.615 2.419a21.853 21.853 0 0 1-3.66 2.093l-.018.008l-.006.003h-.002S12.394 21.92 12 21l-.394.919l-.003-.001l-.005-.003l-.02-.008a21.482 21.482 0 0 1-1.137-.553a21.723 21.723 0 0 1-2.522-1.54c-.918-.656-1.877-1.47-2.615-2.419C4.568 16.45 4 15.3 4 14V5a1 1 0 0 1 1.11-.994c.653.073 1.798.075 3-.173c1.21-.25 2.373-.73 3.183-1.54zm.313 19.626a1 1 0 0 0 .788 0L12 21l-.394.919zM13 11V9a1 1 0 1 0-2 0v2H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2z" clip-rule="evenodd"/>`
	shieldPlusLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14c0 4-7 7-7 7s-7-3-7-7V5c1.5.167 5 0 7-2c2 2 5.5 2.167 7 2v9zm-7-5v3m0 3v-3m0 0h3m-3 0H9"/>`
	shipInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M9 3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2h3a1 1 0 0 1 .981 1.192l-.437 2.238l-1.327-.295l-5-1.111a1 1 0 0 0-.434 0l-5 1.11l-1.327.296l-.437-2.238A1 1 0 0 1 6 5h3V3zm-6.092 7.996l-.125.028a1 1 0 0 0-.677 1.423l2 4a1 1 0 0 0 1.035.543L12 16.01l6.859.98a1 1 0 0 0 1.035-.543l2-4a1 1 0 0 0-.677-1.423l-.125-.028a1 1 0 0 1-.309-.02l-4-.889L12 9.024l-4.783 1.063l-4 .89a1 1 0 0 1-.309.019zm6.36 7.609a3.631 3.631 0 0 1 5.465 0l.035.04a1.57 1.57 0 0 0 2.053.273a3.57 3.57 0 0 1 3.305-.344l1.245.497a1 1 0 0 1-.742 1.857l-1.245-.498a1.57 1.57 0 0 0-1.454.152a3.57 3.57 0 0 1-4.667-.62l-.035-.04a1.631 1.631 0 0 0-2.456 0l-.035.04a3.57 3.57 0 0 1-4.667.62a1.57 1.57 0 0 0-1.454-.152l-1.245.498a1 1 0 1 1-.742-1.857l1.245-.497a3.57 3.57 0 0 1 3.305.344a1.57 1.57 0 0 0 2.053-.273l.035-.04z" clip-rule="evenodd"/>`
	shipLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 16l2-4l-4-.889M12 10v5m0-5l-5 1.111M12 10l5 1.111M5 16l-2-4l4-.889m0 0L6 6h4m7 5.111L18 6h-4m0 0V3h-4v3m4 0h-4M3 20l1.245-.498a2.57 2.57 0 0 1 2.38.248v0a2.57 2.57 0 0 0 3.36-.446l.035-.04a2.631 2.631 0 0 1 3.96 0l.036.04a2.57 2.57 0 0 0 3.36.446v0a2.57 2.57 0 0 1 2.38-.248L21 20"/>`
	shootingStarInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 8c-1.667.667-5.4 2.7-7 5.5m9.5-2.5C9.167 12.333 4 16.4 2 22m10.5-7.5c-1.167 1.167-3.8 4.1-5 6.5"/><path fill="currentColor" d="m14.674 6.45l.673-3.285l2.225 2.51l3.027-.294l-1.768 3.062l1.743 2.639l-3.286-.673l-2.51 2.225l.19-3.156l-3.062-1.768l2.768-1.26z"/></g>`
	shootingStarLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 8c-1.667.667-5.4 2.7-7 5.5m9.5-2.5C9.167 12.333 4 16.4 2 22m10.5-7.5c-1.167 1.167-3.8 4.1-5 6.5m7.174-14.55l.673-3.285l2.225 2.51l3.027-.294l-1.768 3.062l1.743 2.639l-3.286-.673l-2.51 2.225l.19-3.156l-3.062-1.768l2.768-1.26z"/>`
	shoppingBagInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 7a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3H6zm6-3a3 3 0 0 0-3 3v4a1 1 0 1 1-2 0V7a5 5 0 0 1 10 0v4a1 1 0 1 1-2 0V7a3 3 0 0 0-3-3z" fill="currentColor"/></g>`
	shoppingBagLineInnerSVG           = `<g fill="currentColor"><path d="M3 10a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-9zm3-1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1H6z"/><path d="M12 4a3 3 0 0 0-3 3v4a1 1 0 1 1-2 0V7a5 5 0 0 1 10 0v4a1 1 0 1 1-2 0V7a3 3 0 0 0-3-3z"/></g>`
	shoppingCartInnerSVG              = `<g fill="none"><path fill="currentColor" d="M18 15H7L5.5 6H21l-3 9z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.5 3m0 0L7 15h11l3-9H5.5z"/><circle cx="8" cy="20" r="1" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="17" cy="20" r="1" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g>`
	shoppingCartLineInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3h2l.5 3m0 0L7 15h11l3-9H5.5z"/><circle cx="8" cy="20" r="1"/><circle cx="17" cy="20" r="1"/></g>`
	simCardInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M8.707 2.879A3 3 0 0 1 10.828 2H17a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H6.998A2.998 2.998 0 0 1 4 19V8.828a3 3 0 0 1 .879-2.12l3.828-3.83zM14 18a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h4z" clip-rule="evenodd"/>`
	simCardLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 19V5a2 2 0 0 0-2-2h-6.172a2 2 0 0 0-1.414.586L5.586 7.414A2 2 0 0 0 5 8.828V19c0 1.105.893 2 1.997 2h10.004A2 2 0 0 0 19 19z"/>`
	sitemapInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M10 4h4v4h-4zm0 12h4v4h-4zm-8 0h4v4H2zm16 0h4v4h-4z"/><path d="M12 8v4m0 4v-4m0 0h6a2 2 0 0 1 2 2v2m-8-4H6a2 2 0 0 0-2 2v2"/></g>`
	sitemapLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 4h4v4h-4zm0 12h4v4h-4zm-7 0h4v4H3zm14 0h4v4h-4zm-5-8v4m0 4v-4m0 0h5a2 2 0 0 1 2 2v2m-7-4H7a2 2 0 0 0-2 2v2"/>`
	skullInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M22 11a9.96 9.96 0 0 1-2 6v2a3 3 0 0 1-3 3v-1a1 1 0 1 0-2 0v1h-2v-1a1 1 0 1 0-2 0v1H9v-1a1 1 0 1 0-2 0v1a3 3 0 0 1-3-3v-2a9.96 9.96 0 0 1-2-6c0-2.779 1.136-5.058 2.99-6.627C6.828 2.817 9.318 2 12 2c2.682 0 5.172.817 7.01 2.373C20.863 5.942 22 8.22 22 11zM8 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm4.894 5.415C12.725 14.16 12.38 14 12 14s-.725.16-.894.415l-1 1.5a.591.591 0 0 0 .043.73c.183.22.504.355.851.355h2c.347 0 .668-.135.85-.356a.591.591 0 0 0 .044-.73l-1-1.5zM14 11a2 2 0 1 1 4 0a2 2 0 0 1-4 0z" clip-rule="evenodd"/>`
	skullLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 16.657A8.962 8.962 0 0 1 3 11c0-4.97 4.03-8 9-8s9 3.03 9 8c0 2.143-.75 4.112-2 5.657m-14 0V19a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-2.343m-14 0h1.5m12.5 0h-1.5"/><circle cx="8" cy="12" r="1"/><path d="m11 16l1-2l1 2h-2zm-2 3v2m3-2v2m3-2v2"/><circle cx="16" cy="12" r="1"/></g>`
	smartphoneAppsInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M8 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3H8zm1 5.001a1 1 0 0 1-1-1V6a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm-1 2a1 1 0 0 0 2 0V9a1 1 0 1 0-2 0v.001zm4 1a1 1 0 0 1-1-1V9a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm-1-4a1 1 0 1 0 2 0V6a1 1 0 1 0-2 0v.001zm4 1a1 1 0 0 1-1-1V6a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm-1 2a1 1 0 1 0 2 0V9a1 1 0 1 0-2 0v.001zm-5 4a1 1 0 0 1-1-1V12a1 1 0 1 1 2 0v.001a1 1 0 0 1-1 1zm2-1a1 1 0 1 0 2 0V12a1 1 0 1 0-2 0v.001z" clip-rule="evenodd"/>`
	smartphoneAppsLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 19V5a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zM9 6.001V6m0 3.001V9m3 .001V9m0-2.999V6m3 .001V6m0 3.001V9m-6 3.001V12m3 .001V12"/>`
	sortHorizontalInnerSVG            = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M17.707 4.293A1 1 0 0 0 16 5v2h-6a1 1 0 1 0 0 2h6v2a1 1 0 0 0 1.707.707l3-3a1 1 0 0 0 0-1.414l-3-3zm-11.414 8A1 1 0 0 1 8 13v2h6a1 1 0 1 1 0 2H8v2a1 1 0 0 1-1.707.707l-3-3a1 1 0 0 1 0-1.414l3-3z" fill="currentColor"/></g>`
	sortHorizontalLineInnerSVG        = `<g fill="currentColor"><path d="M17.707 4.293a1 1 0 1 0-1.414 1.414L17.586 7H10a1 1 0 0 0 0 2h7.586l-1.293 1.293a1 1 0 0 0 1.414 1.414l3-3a1 1 0 0 0 0-1.414l-3-3zm-11.414 8a1 1 0 0 1 1.414 1.414L6.414 15H14a1 1 0 1 1 0 2H6.414l1.293 1.293a1 1 0 1 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414l3-3z"/></g>`
	sortVerticalInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M4.293 6.293A1 1 0 0 0 5 8h2v6a1 1 0 1 0 2 0V8h2a1 1 0 0 0 .707-1.707l-3-3a1 1 0 0 0-1.414 0l-3 3zm8 11.414A1 1 0 0 1 13 16h2v-6a1 1 0 1 1 2 0v6h2a1 1 0 0 1 .707 1.707l-3 3a1 1 0 0 1-1.414 0l-3-3z" fill="currentColor"/></g>`
	sortVerticalLineInnerSVG          = `<g fill="currentColor"><path d="M4.293 6.293a1 1 0 0 0 1.414 1.414L7 6.414V14a1 1 0 1 0 2 0V6.414l1.293 1.293a1 1 0 1 0 1.414-1.414l-3-3a1 1 0 0 0-1.414 0l-3 3zm8 11.414a1 1 0 0 1 1.414-1.414L15 17.586V10a1 1 0 1 1 2 0v7.586l1.293-1.293a1 1 0 0 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3z"/></g>`
	soundOffInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M13 4.703c0-1.857-2.31-2.711-3.519-1.301L5.84 7.65a1 1 0 0 1-.76.35H5a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h.08a1 1 0 0 1 .76.35l.712-.612l-.713.611l3.642 4.25c1.209 1.41 3.519.555 3.519-1.302V4.703zm3.293 4.59a1 1 0 0 1 1.414 0L19 10.586l1.293-1.293a1 1 0 1 1 1.414 1.414L20.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L19 13.414l-1.293 1.293a1 1 0 0 1-1.414-1.414L17.586 12l-1.293-1.293a1 1 0 0 1 0-1.414z" fill="currentColor"/></g>`
	soundOffLineInnerSVG              = `<g fill="currentColor"><path d="M13 4.703c0-1.857-2.31-2.711-3.519-1.301L5.84 7.65a1 1 0 0 1-.76.35H5a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h.08a1 1 0 0 1 .76.35l3.641 4.248c1.209 1.41 3.519.556 3.519-1.301V4.703zm-5.642 4.25L11 4.702v14.594l-3.642-4.25A3 3 0 0 0 5.08 14H5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h.08a3 3 0 0 0 2.278-1.048zm10.35.34a1 1 0 1 0-1.415 1.414L17.586 12l-1.293 1.293a1 1 0 0 0 1.414 1.414L19 13.414l1.293 1.293a1 1 0 0 0 1.414-1.414L20.414 12l1.293-1.293a1 1 0 0 0-1.414-1.414L19 10.586l-1.293-1.293z"/></g>`
	soundUpInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5.08 9H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h.08a2 2 0 0 1 1.519.698l3.642 4.25c.604.704 1.759.277 1.759-.651V4.703c0-.928-1.155-1.355-1.76-.65L6.6 8.301A2 2 0 0 1 5.08 9zm13.556-4.725a1 1 0 1 0-1.377 1.45c3.655 3.472 3.655 9.078 0 12.55a1 1 0 1 0 1.377 1.45c4.485-4.26 4.485-11.19 0-15.45zm-2.947 2.8a1 1 0 1 0-1.378 1.45c2.027 1.925 2.027 5.025 0 6.95a1 1 0 1 0 1.378 1.45c2.857-2.714 2.857-7.136 0-9.85z" fill="currentColor"/></g>`
	soundUpLineInnerSVG               = `<g fill="currentColor"><path d="M13 4.703c0-1.857-2.31-2.711-3.519-1.301L5.84 7.65a1 1 0 0 1-.76.35H5a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h.08a1 1 0 0 1 .76.35l3.641 4.248c1.209 1.41 3.519.556 3.519-1.301V4.703zm-5.642 4.25L11 4.702v14.594l-3.642-4.25A3 3 0 0 0 5.08 14H5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h.08a3 3 0 0 0 2.278-1.048zm11.278-4.678a1 1 0 1 0-1.377 1.45c3.655 3.472 3.655 9.078 0 12.55a1 1 0 0 0 1.377 1.45c4.485-4.26 4.485-11.19 0-15.45zm-2.947 2.8a1 1 0 1 0-1.378 1.45c2.027 1.925 2.027 5.025 0 6.95a1 1 0 0 0 1.378 1.45c2.857-2.714 2.857-7.136 0-9.85z"/></g>`
	sparklesInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 3a1 1 0 0 0-2 0v1H3a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0V6h1a1 1 0 0 0 0-2H6V3zm7-1a1 1 0 0 1 1 1c0 3.344 1.148 5.296 2.514 6.43C17.918 10.598 19.672 11 21 11a1 1 0 1 1 0 2c-3.26 0-4.924 1.324-5.838 2.881C14.2 17.524 14 19.556 14 21a1 1 0 1 1-2 0c0-3.344-1.148-5.296-2.514-6.43C8.082 13.402 6.328 13 5 13a1 1 0 1 1 0-2c3.26 0 4.924-1.324 5.838-2.881C11.8 6.476 12 4.444 12 3a1 1 0 0 1 1-1z" fill="currentColor"/></g>`
	sparklesLineInnerSVG              = `<g fill="currentColor"><path d="M6 3a1 1 0 0 0-2 0v1H3a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0V6h1a1 1 0 0 0 0-2H6V3zm8 0a1 1 0 1 0-2 0c0 1.444-.199 3.476-1.162 5.119C9.924 9.676 8.26 11 5 11a1 1 0 1 0 0 2c1.328 0 3.082.403 4.486 1.57C10.852 15.703 12 17.655 12 21a1 1 0 1 0 2 0c0-1.444.199-3.476 1.162-5.119C16.076 14.324 17.74 13 21 13a1 1 0 1 0 0-2c-1.328 0-3.082-.403-4.486-1.57C15.148 8.297 14 6.345 14 3zm-1.438 6.131c.187-.318.35-.645.493-.975a8.295 8.295 0 0 0 2.181 2.813c.444.369.911.678 1.388.934c-1.445.7-2.476 1.754-3.186 2.966a8.82 8.82 0 0 0-.493.975a8.296 8.296 0 0 0-2.181-2.813a8.388 8.388 0 0 0-1.388-.934c1.445-.7 2.476-1.754 3.187-2.966z"/></g>`
	speakerInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M7 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3H7zm4 5a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H12a1 1 0 0 1-1-1zm-1 7a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm2-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8z" clip-rule="evenodd"/>`
	speakerLineInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 3H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm-5 4h.01"/><circle cx="12" cy="14" r="3"/></g>`
	speakerphoneInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M21 4.723c0-1.535-1.659-2.498-2.992-1.736L12.734 6H7a5 5 0 0 0-1 9.9v3.6a2.5 2.5 0 0 0 5 0V16h1.734l5.274 3.013c1.333.762 2.992-.2 2.992-1.736V4.723z" fill="currentColor"/></g>`
	speakerphoneLineInnerSVG          = `<g fill="currentColor"><path d="M18.008 2.987C19.34 2.225 21 3.187 21 4.723v12.554c0 1.535-1.659 2.498-2.992 1.736L12.734 16H11v3.5a2.5 2.5 0 0 1-5 0v-3.6A5.002 5.002 0 0 1 7 6h5.734l5.274-3.013zM12 8H7a3 3 0 0 0 0 6h5V8zm2 6.42l5 2.857V4.723L14 7.58v6.84zM8 16v3.5a.5.5 0 0 0 1 0V16H8z"/></g>`
	starInnerSVG                      = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12.897 2.557a1 1 0 0 0-1.794 0L8.691 7.445l-5.394.784a1 1 0 0 0-.555 1.706l3.904 3.805l-.922 5.372a1 1 0 0 0 1.451 1.054L12 17.63l4.825 2.536a1 1 0 0 0 1.45-1.054l-.92-5.372l3.902-3.805a1 1 0 0 0-.554-1.706l-5.394-.784l-2.412-4.888z" fill="currentColor"/></g>`
	starLineInnerSVG                  = `<g fill="currentColor"><path d="M10.655 3.466c.55-1.115 2.14-1.115 2.69 0l1.964 3.98l4.392.638c1.23.178 1.721 1.69.831 2.558l-3.178 3.098l.75 4.374c.21 1.225-1.076 2.16-2.176 1.58L12 17.63l-3.928 2.065c-1.1.578-2.387-.356-2.176-1.581l.75-4.374l-3.178-3.098c-.89-.868-.399-2.38.831-2.558l4.392-.639l1.964-3.98zM12 5.26l-1.632 3.306a1.5 1.5 0 0 1-1.13.82l-3.649.531l2.641 2.574a1.5 1.5 0 0 1 .431 1.328l-.623 3.634l3.264-1.716a1.5 1.5 0 0 1 1.396 0l3.264 1.716l-.623-3.634a1.5 1.5 0 0 1 .431-1.328l2.64-2.574l-3.649-.53a1.5 1.5 0 0 1-1.129-.82L12 5.26z"/></g>`
	statusOfflineInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3.707 2.293a1 1 0 0 0-1.414 1.414L15.535 16.95l2.829 2.828l1.929 1.93a1 1 0 0 0 1.414-1.415l-1.253-1.254c3.607-4.321 3.382-10.76-.676-14.817a1 1 0 1 0-1.414 1.414a9.001 9.001 0 0 1 .668 11.982l-1.425-1.425a7.002 7.002 0 0 0-.657-9.143a1 1 0 1 0-1.414 1.414a5.002 5.002 0 0 1 .636 6.294l-1.465-1.465a3 3 0 0 0-4-4l-7-7zM3.75 8.4a1 1 0 0 0-1.834-.8C.161 11.624.928 16.485 4.222 19.778a1 1 0 0 0 1.414-1.414A9.004 9.004 0 0 1 3.749 8.4zm3.32 2.766a1 1 0 0 0-1.972-.332A6.992 6.992 0 0 0 7.05 16.95a1 1 0 1 0 1.414-1.414a4.993 4.993 0 0 1-1.394-4.37z" fill="currentColor"/></g>`
	statusOfflineLineInnerSVG         = `<g fill="currentColor"><path d="M3.707 2.293a1 1 0 0 0-1.414 1.414L15.535 16.95l2.829 2.828l1.929 1.93a1 1 0 0 0 1.414-1.415l-1.253-1.254c3.607-4.321 3.382-10.76-.676-14.817a1 1 0 0 0-1.414 1.414a9.001 9.001 0 0 1 .668 11.982l-1.425-1.425a7.002 7.002 0 0 0-.657-9.143a1 1 0 1 0-1.414 1.414a5.002 5.002 0 0 1 .636 6.294l-1.465-1.465a3 3 0 0 0-4-4l-7-7zM3.75 8.4a1 1 0 0 0-1.834-.8C.161 11.624.928 16.485 4.222 19.778a1 1 0 0 0 1.414-1.414A9.004 9.004 0 0 1 3.749 8.4zm3.32 2.766a1 1 0 0 0-1.972-.332A6.992 6.992 0 0 0 7.05 16.95a1 1 0 1 0 1.414-1.414a4.992 4.992 0 0 1-1.394-4.37z"/></g>`
	statusOnlineInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5.636 5.636a1 1 0 0 0-1.414-1.414c-4.296 4.296-4.296 11.26 0 15.556a1 1 0 0 0 1.414-1.414a9 9 0 0 1 0-12.728zm14.142-1.414a1 1 0 1 0-1.414 1.414a9 9 0 0 1 0 12.728a1 1 0 1 0 1.414 1.414c4.296-4.296 4.296-11.26 0-15.556zM8.464 8.464A1 1 0 0 0 7.05 7.05a7 7 0 0 0 0 9.9a1 1 0 1 0 1.414-1.414a5 5 0 0 1 0-7.072zM16.95 7.05a1 1 0 1 0-1.414 1.414a5 5 0 0 1 0 7.072a1 1 0 0 0 1.414 1.414a7 7 0 0 0 0-9.9zM9 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0z" fill="currentColor"/></g>`
	statusOnlineLineInnerSVG          = `<g fill="currentColor"><path d="M5.636 5.636a1 1 0 0 0-1.414-1.414c-4.296 4.296-4.296 11.26 0 15.556a1 1 0 0 0 1.414-1.414a9 9 0 0 1 0-12.728zm14.142-1.414a1 1 0 1 0-1.414 1.414a9 9 0 0 1 0 12.728a1 1 0 1 0 1.414 1.414c4.296-4.296 4.296-11.26 0-15.556zM8.464 8.464A1 1 0 0 0 7.05 7.05a7 7 0 0 0 0 9.9a1 1 0 0 0 1.414-1.414a5 5 0 0 1 0-7.072zM16.95 7.05a1 1 0 1 0-1.414 1.414a5 5 0 0 1 0 7.072a1 1 0 0 0 1.414 1.414a7 7 0 0 0 0-9.9zM11 12a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm1-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6z"/></g>`
	stopCircleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11zM9 8a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1H9zm1 6v-4h4v4h-4z" clip-rule="evenodd"/>`
	stopCircleLineInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 9h6v6H9z"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`
	strikeThroughInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 5h-7a3 3 0 0 0-3 3v1a3 3 0 0 0 3 3h7M7 19h7a3 3 0 0 0 3-3v-1M5 12h14"/>`
	strikeThroughLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 5h-7a3 3 0 0 0-3 3v1a3 3 0 0 0 3 3h7M7 19h7a3 3 0 0 0 3-3v-1M5 12h14"/>`
	suitcaseInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M11 3a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3h-3V6a3 3 0 0 0-3-3h-2zm3 4h-4V6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1z" clip-rule="evenodd"/>`
	suitcaseLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8V6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v2m6 0h4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4m6 0H9"/>`
	suitcaseThreeInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M8 5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3h2a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3a1 1 0 1 1-2 0H8a1 1 0 1 1-2 0a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h2zm2 0h4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1zm0 5a1 1 0 0 0-2 0v6a1 1 0 1 0 2 0v-6zm6 0a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0v-6z" clip-rule="evenodd"/>`
	suitcaseThreeLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6V5a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v1m6 0h3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1M15 6H9m0 0H6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h1m2-10v6m6 0v-6M7 20v1m0-1h10m0 0v1"/>`
	suitcaseTwoInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M8 6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v15H8V6zM6 7H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h1V7zm12 14h1a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3h-1v14zM10 7h4V6a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v1z" clip-rule="evenodd"/>`
	suitcaseTwoLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8V6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v2m6 0h4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H7m8-12H9m0 0H7m0 0H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h2M7 8v12M17 8v12"/>`
	sunInnerSVG                       = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M13 3a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0V3zM5.707 4.293a1 1 0 0 0-1.414 1.414l1 1a1 1 0 0 0 1.414-1.414l-1-1zm14 0a1 1 0 0 0-1.414 0l-1 1a1 1 0 0 0 1.414 1.414l1-1a1 1 0 0 0 0-1.414zM12 7a5 5 0 1 0 0 10a5 5 0 0 0 0-10zm-9 4a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H3zm17 0a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2h-1zM6.707 18.707a1 1 0 1 0-1.414-1.414l-1 1a1 1 0 1 0 1.414 1.414l1-1zm12-1.414a1 1 0 0 0-1.414 1.414l1 1a1 1 0 0 0 1.414-1.414l-1-1zM13 20a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0v-1z" fill="currentColor"/></g>`
	sunLineInnerSVG                   = `<g fill="currentColor"><path d="M12 2a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0V3a1 1 0 0 1 1-1zM4.293 4.293a1 1 0 0 1 1.414 0l1 1a1 1 0 0 1-1.414 1.414l-1-1a1 1 0 0 1 0-1.414zm14 0a1 1 0 1 1 1.414 1.414l-1 1a1 1 0 1 1-1.414-1.414l1-1zM12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0zm-5 0a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1zm17 0a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1zM6.707 17.293a1 1 0 0 1 0 1.414l-1 1a1 1 0 0 1-1.414-1.414l1-1a1 1 0 0 1 1.414 0zm10.586 0a1 1 0 0 1 1.414 0l1 1a1 1 0 0 1-1.414 1.414l-1-1a1 1 0 0 1 0-1.414zM12 19a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0v-1a1 1 0 0 1 1-1z"/></g>`
	supportInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5.67 4.258A9.959 9.959 0 0 1 12 2a9.96 9.96 0 0 1 6.33 2.258a.977.977 0 0 0-.037.035l-3.536 3.535A4.977 4.977 0 0 0 12 7c-1.02 0-1.967.305-2.757.828L5.707 4.293a1.018 1.018 0 0 0-.037-.035zM4.258 5.67A9.959 9.959 0 0 0 2 12a9.96 9.96 0 0 0 2.258 6.33a.977.977 0 0 1 .035-.037l3.535-3.536A4.977 4.977 0 0 1 7 12c0-1.02.305-1.967.828-2.757L4.293 5.707a1.018 1.018 0 0 1-.035-.037zM5.67 19.742A9.959 9.959 0 0 0 12 22a9.959 9.959 0 0 0 6.33-2.258a.93.93 0 0 1-.037-.035l-3.536-3.535A4.977 4.977 0 0 1 12 17a4.977 4.977 0 0 1-2.757-.828l-3.536 3.535a.967.967 0 0 1-.037.035zm14.072-1.412A9.959 9.959 0 0 0 22 12a9.959 9.959 0 0 0-2.258-6.33a.967.967 0 0 1-.035.037l-3.535 3.536c.523.79.828 1.738.828 2.757c0 1.02-.305 1.967-.828 2.757l3.535 3.536a.93.93 0 0 1 .035.037z" fill="currentColor"/><circle cx="12" cy="12" r="3" fill="currentColor"/></g>`
	supportLineInnerSVG               = `<g fill="currentColor"><path d="M5.68 7.094A7.965 7.965 0 0 0 4 12c0 1.849.627 3.551 1.68 4.906l2.148-2.149A4.977 4.977 0 0 1 7 12c0-1.02.305-1.967.828-2.757L5.68 7.094zM7.094 5.68l2.149 2.148A4.977 4.977 0 0 1 12 7c1.02 0 1.967.305 2.757.828l2.149-2.148A7.965 7.965 0 0 0 12 4a7.965 7.965 0 0 0-4.906 1.68zM18.32 7.094l-2.148 2.149c.523.79.828 1.738.828 2.757c0 1.02-.305 1.967-.828 2.757l2.148 2.149A7.965 7.965 0 0 0 20 12a7.966 7.966 0 0 0-1.68-4.906zM16.906 18.32l-2.149-2.148A4.977 4.977 0 0 1 12 17a4.977 4.977 0 0 1-2.757-.828L7.094 18.32A7.966 7.966 0 0 0 12 20a7.965 7.965 0 0 0 4.906-1.68zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm10-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6z"/></g>`
	tShirtInnerSVG                    = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10V5a1 1 0 0 1 1-1h4c0 1 .8 3 4 3s4-2 4-3h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-1a1 1 0 0 0-1 1v7a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 0-1-1H4a1 1 0 0 1-1-1z"/>`
	tShirtLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10V5a1 1 0 0 1 1-1h4c0 1 .8 3 4 3s4-2 4-3h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-1a1 1 0 0 0-1 1v7a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 0-1-1H4a1 1 0 0 1-1-1z"/>`
	tableInnerSVG                     = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v6m0-6h6m12 0v6m0-6H9m12 6v4a2 2 0 0 1-2 2H9m12-6H9m-6 0v4a2 2 0 0 0 2 2h4m-6-6h6m0-6v6m0 0v6m6-12v12"/><path fill="currentColor" d="M3 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2H3V6z"/></g>`
	tableHeartInnerSVG                = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20H9M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v6m0-6h6m12 0v3m0-3H9m-6 6v4a2 2 0 0 0 2 2h4m-6-6h6m0-6v6m0 0v6m0-6h1m5-6v3"/><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 16c0 1.2-2.667 3.833-4 5c-1.333-1.167-4-3.8-4-5c0-.667.4-2 2-2s2 1.333 2 2c0-.667.4-2 2-2s2 1.333 2 2z"/><path fill="currentColor" d="M3 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2H3V6z"/></g>`
	tableHeartLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20H9M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v6m0-6h6m12 0v3m0-3H9m-6 6v4a2 2 0 0 0 2 2h4m-6-6h6m0-6v6m0 0v6m0-6h1m5-6v3m6 5c0 1.2-2.667 3.833-4 5c-1.333-1.167-4-3.8-4-5c0-.667.4-2 2-2s2 1.333 2 2c0-.667.4-2 2-2s2 1.333 2 2z"/>`
	tableLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v6m0-6h6m12 0v6m0-6H9m12 6v4a2 2 0 0 1-2 2H9m12-6H9m-6 0v4a2 2 0 0 0 2 2h4m-6-6h6m0-6v6m0 0v6m6-12v12"/>`
	tablePlusInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v6m0-6h6m12 0v4m0-4H9m-6 6v4a2 2 0 0 0 2 2h4m-6-6h6m0-6v6m0 0h4a2 2 0 0 0 2-2V8m-6 6v6m0 0h2m7-5v3m0 0v3m0-3h3m-3 0h-3"/><path fill="currentColor" d="M3 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2H3V6z"/></g>`
	tablePlusLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v6m0-6h6m12 0v4m0-4H9m-6 6v4a2 2 0 0 0 2 2h4m-6-6h6m0-6v6m0 0h4a2 2 0 0 0 2-2V8m-6 6v6m0 0h2m7-5v3m0 0v3m0-3h3m-3 0h-3"/>`
	tagInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M2 5a3 3 0 0 1 3-3h6.172a3 3 0 0 1 2.12.879l8 8a3 3 0 0 1 0 4.242l-6.17 6.172a3 3 0 0 1-4.243 0l-8-8A3 3 0 0 1 2 11.172V5zm5 1a1 1 0 0 0 0 2h.001a1 1 0 0 0 0-2H7z" clip-rule="evenodd"/>`
	tagLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11.172V5a2 2 0 0 1 2-2h6.172a2 2 0 0 1 1.414.586l8 8a2 2 0 0 1 0 2.828l-6.172 6.172a2 2 0 0 1-2.828 0l-8-8A2 2 0 0 1 3 11.172zM7 7h.001"/>`
	tagOffInnerSVG                    = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 3L3 21"/><path fill="currentColor" fill-rule="evenodd" d="m13.293 2.879l1.5 1.5L4.379 14.793l-1.5-1.5A3 3 0 0 1 2 11.172V5a3 3 0 0 1 3-3h6.172a3 3 0 0 1 2.12.879zm4.328 4.328L7.207 17.621l3.672 3.672a3 3 0 0 0 4.242 0l6.172-6.172a3 3 0 0 0 0-4.242L17.62 7.207zM7 6a1 1 0 0 0 0 2h.001a1 1 0 0 0 0-2H7z" clip-rule="evenodd"/></g>`
	tagOffLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 5l-1.414-1.414A2 2 0 0 0 11.172 3H5a2 2 0 0 0-2 2v6.172a2 2 0 0 0 .586 1.414L5 14m14-4l1.586 1.586a2 2 0 0 1 0 2.828l-6.172 6.172a2 2 0 0 1-2.828 0L10 19M7 7h.001M21 3L3 21"/>`
	telescopeInnerSVG                 = `<g fill="none"><g clip-path="url(#majesticonsTelescope0)"><path fill="currentColor" d="m11.026 7.366l6.063-3.5l3 5.196l-6.063 3.5a1 1 0 0 1-1.366-.366l-3.03 1.75l-2.166 1.25a1 1 0 0 1-1.366-.366l-1.732 1a1 1 0 1 1-1-1.732l1.732-1a1 1 0 0 1 .366-1.366l5.196-3a1 1 0 0 1 .366-1.366z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16.588 3l.5.866m3.5 6.062l-.5-.866m-3-5.196l-6.062 3.5a1 1 0 0 0-.366 1.366v0m6.428-4.866l3 5.196m0 0l-6.062 3.5a1 1 0 0 1-1.366-.366v0m-2-3.464l-5.196 3a1 1 0 0 0-.366 1.366v0m5.562-4.366l2 3.464m0 0l-3.03 1.75m-4.532-.848l-1.732 1A1 1 0 0 0 3 15.464v0a1 1 0 0 0 1.366.366l1.732-1m-1-1.732l1 1.732m0 0v0a1 1 0 0 0 1.366.366l2.165-1.25m0 0L13 20m-3.37-6.054L7.097 20"/></g><defs><clipPath id="majesticonsTelescope0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	telescopeLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16.588 3l.5.866m3.5 6.062l-.5-.866m-3-5.196l-6.062 3.5a1 1 0 0 0-.366 1.366v0m6.428-4.866l3 5.196m0 0l-6.062 3.5a1 1 0 0 1-1.366-.366v0m-2-3.464l-5.196 3a1 1 0 0 0-.366 1.366v0m5.562-4.366l2 3.464m0 0l-3.03 1.75m-4.532-.848l-1.732 1A1 1 0 0 0 3 15.464v0a1 1 0 0 0 1.366.366l1.732-1m-1-1.732l1 1.732m0 0v0a1 1 0 0 0 1.366.366l2.165-1.25m0 0L13 20m-3.37-6.054L7.097 20"/>`
	terminalInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm4.617 1.076a1 1 0 0 1 1.09.217l4 4a1 1 0 0 1 0 1.414l-4 4A1 1 0 0 1 6 16V8a1 1 0 0 1 .617-.924zM12 16a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1z" fill="currentColor"/></g>`
	terminalLineInnerSVG              = `<g fill="currentColor"><path d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H5zm1.617 2.076a1 1 0 0 1 1.09.217l4 4a1 1 0 0 1 0 1.414l-4 4A1 1 0 0 1 6 16V8a1 1 0 0 1 .617-.924zM8 10.414v3.172L9.586 12L8 10.414zM12 16a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1z"/></g>`
	testTubeFilledInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M6 2a1 1 0 0 0 0 2h1v13a5 5 0 0 0 10 0V4h1a1 1 0 1 0 0-2H6zm3 5V4h6v3H9zm4.708 5.708a1 1 0 0 1-1.414 0h-.001a1 1 0 0 1 1.414-1.415a1 1 0 0 1 0 1.415zm-3.414 4a1 1 0 1 0 1.414-1.414v-.001a1 1 0 0 0-1.415 1.414z" clip-rule="evenodd"/>`
	testTubeFilledLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3h2m10 0h-2m0 0H8m8 0v5M8 3v5m0 0v9a4 4 0 0 0 4 4v0a4 4 0 0 0 4-4V8M8 8h8m-3 4h0m-2 4h0"/>`
	textInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 12h14M5 16h6"/>`
	textAlignCenterInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6h14M7 10h10M5 14h14M9 18h6"/>`
	textAlignCenterLineInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6h14M7 10h10M5 14h14M9 18h6"/>`
	textAlignJustifyInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6h14M5 10h14M5 14h14M5 18h14"/>`
	textAlignJustifyLineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6h14M5 10h14M5 14h14M5 18h14"/>`
	textAlignLeftInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6h14M5 10h10M5 14h14M5 18h6"/>`
	textAlignLeftLineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6h14M5 10h10M5 14h14M5 18h6"/>`
	textAlignRightInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6h14M7 10h10M5 14h14M9 18h6"/>`
	textAlignRightLineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6h14M7 10h10M5 14h14M9 18h6"/>`
	textLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 12h14M5 16h6"/>`
	textWrapInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h14M5 12h11c1 0 3 .5 3 2.5S17.333 17 16.5 17H12m-7 0h4m3 0l2-2m-2 2l2 2"/>`
	textWrapLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h14M5 12h11c1 0 3 .5 3 2.5S17.333 17 16.5 17H12m-7 0h4m3 0l2-2m-2 2l2 2"/>`
	textboxInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6zm5 1a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H8z" clip-rule="evenodd"/>`
	textboxLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8h8m0 4H8m0 4h4m-6 4h12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2z"/>`
	textboxMinusInnerSVG              = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 18h-6"/><path fill="currentColor" fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h6.803A5.972 5.972 0 0 1 12 18c0-.342.029-.677.084-1.003A1.048 1.048 0 0 1 12 17H8a1 1 0 1 1 0-2h4c.255 0 .488.095.665.253A6.029 6.029 0 0 1 14.681 13H8a1 1 0 1 1 0-2h8a1 1 0 0 1 .997 1.084A6.044 6.044 0 0 1 18 12c1.093 0 2.117.292 3 .803V6a3 3 0 0 0-3-3H6zm1 5a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1z" clip-rule="evenodd"/></g>`
	textboxMinusLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8h8m0 4H8m0 4h4m8-4V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6m9-2h-6"/>`
	textboxPlusInnerSVG               = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 15v3m0 3v-3m0 0h3m-3 0h-3"/><path fill="currentColor" fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h6.803A5.972 5.972 0 0 1 12 18c0-.342.029-.677.084-1.003A1.048 1.048 0 0 1 12 17H8a1 1 0 1 1 0-2h4c.255 0 .488.095.665.253A6.029 6.029 0 0 1 14.681 13H8a1 1 0 1 1 0-2h8a1 1 0 0 1 .997 1.084A6.044 6.044 0 0 1 18 12c1.093 0 2.117.292 3 .803V6a3 3 0 0 0-3-3H6zm1 5a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1z" clip-rule="evenodd"/></g>`
	textboxPlusLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8h8m0 4H8m0 4h4m8-4V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6m6-5v3m0 3v-3m0 0h3m-3 0h-3"/>`
	thumbDownInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M16 3H8a5 5 0 0 0-5 5v2a5 5 0 0 0 4.82 4.997l-.563 3.378a2.254 2.254 0 0 0 3.817 1.965L16 15.414V3zm2 12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3v12z" fill="currentColor"/></g>`
	thumbDownLineInnerSVG             = `<g fill="currentColor"><path d="M11.074 20.34a2.254 2.254 0 0 1-3.817-1.965l.563-3.378A5 5 0 0 1 3 10V8a5 5 0 0 1 5-5h10a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3h-1.586l-5.34 5.34zM17 13h1a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-1v8zm-2-8H8a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h3a1 1 0 1 1 0 2H9.847l-.617 3.704a.254.254 0 0 0 .43.222l5.34-5.34V5z"/></g>`
	thumbUpInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12.926 3.66a2.254 2.254 0 0 1 3.817 1.965l-.563 3.378A5 5 0 0 1 21 14v2a5 5 0 0 1-5 5H8V8.586l4.926-4.926zM6 9a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3V9z" fill="currentColor"/></g>`
	thumbUpLineInnerSVG               = `<g fill="currentColor"><path d="M12.926 3.66a2.254 2.254 0 0 1 3.817 1.965l-.563 3.378A5 5 0 0 1 21 14v2a5 5 0 0 1-5 5H6a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h1.586l5.34-5.34zM7 11H6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h1v-8zm2 8h7a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3h-3a1 1 0 1 1 0-2h1.153l.617-3.704a.254.254 0 0 0-.43-.222L9 10.414V19z"/></g>`
	ticketInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4 4h12a1 1 0 1 0 2 0h2a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-2a1 1 0 1 0-2 0H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2zm14 4.001a1 1 0 1 1-2 0V8a1 1 0 1 1 2 0v.001zm-1 5a1 1 0 0 0 1-1V12a1 1 0 1 0-2 0v.001a1 1 0 0 0 1 1zm1 3a1 1 0 1 1-2 0V16a1 1 0 1 1 2 0v.001z" clip-rule="evenodd"/>`
	ticketCheckInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M16 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a1 1 0 1 1 2 0h2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-2a1 1 0 1 1-2 0zm1 5.001a1 1 0 0 0 1-1V8a1 1 0 1 0-2 0v.001a1 1 0 0 0 1 1zm1 3a1 1 0 1 1-2 0V12a1 1 0 1 1 2 0v.001zm-1 5a1 1 0 0 0 1-1V16a1 1 0 1 0-2 0v.001a1 1 0 0 0 1 1zm-3.168-7.446a1 1 0 0 0-1.664-1.11L8.845 13.43l-1.138-1.137a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.54-.152l4-6z" clip-rule="evenodd"/>`
	ticketCheckLineInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 19H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h13m0 14h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-3m0 14v-1m0-13v1m0 3.001V9m0 3.001V12m0 3.001V15"/><path d="m7 13l2 2l4-6"/></g>`
	ticketLineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 19H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h13m0 14h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-3m0 14v-1m0-13v1m0 3.001V9m0 3.001V12m0 3.001V15"/>`
	ticketTextInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M16 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a1 1 0 1 1 2 0h2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-2a1 1 0 1 1-2 0zm1 5.001a1 1 0 0 0 1-1V8a1 1 0 1 0-2 0v.001a1 1 0 0 0 1 1zm1 3a1 1 0 1 1-2 0V12a1 1 0 1 1 2 0v.001zm-1 5a1 1 0 0 0 1-1V16a1 1 0 1 0-2 0v.001a1 1 0 0 0 1 1zM6 11a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H6zm0 4a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2H6z" clip-rule="evenodd"/>`
	ticketTextLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 19H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h13m0 14h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-3m0 14v-1m0-13v1m0 3.001V9m0 3.001V12m0 3.001V15M7 12h6m-6 3h3"/>`
	ticketsInnerSVG                   = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 20H4a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1h3m6 11h3a1 1 0 0 0 1-1v-2.5M13 20v-1"/><path fill="currentColor" fill-rule="evenodd" d="M22 6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V6zm-4 .001a1 1 0 1 1-2 0V6a1 1 0 1 1 2 0v.001zm-1 5a1 1 0 0 0 1-1V10a1 1 0 1 0-2 0v.001a1 1 0 0 0 1 1zm1 3a1 1 0 1 1-2 0V14a1 1 0 1 1 2 0v.001z" clip-rule="evenodd"/></g>`
	ticketsLineInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 16H8a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h9m0 11h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-3m0 11v-1m0-10v1"/><path d="M13 20H4a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1h3m6 11h3a1 1 0 0 0 1-1v-2.5M13 20v-1m4-9.999V9m0 3.001V12"/></g>`
	timerInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M8 3a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1zM3 14a9 9 0 0 1 14.618-7.032l.675-.675a1 1 0 1 1 1.414 1.414l-.675.675A9 9 0 1 1 3 14zm10-4a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4z" clip-rule="evenodd"/>`
	timerLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 7l-1.343 1.343m0 0A8 8 0 1 0 6.343 19.657A8 8 0 0 0 17.657 8.343zM12 10v4M9 3h6"/>`
	tooltipInnerSVG                   = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4H5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.172a2 2 0 0 1 1.414.586l1.707 1.707a1 1 0 0 0 1.414 0l1.707-1.707A2 2 0 0 1 15.828 18H19a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/>`
	tooltipLineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4H5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.172a2 2 0 0 1 1.414.586l1.707 1.707a1 1 0 0 0 1.414 0l1.707-1.707A2 2 0 0 1 15.828 18H19a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/>`
	tooltipTextInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3h-3.172a1 1 0 0 0-.707.293L13.414 21a2 2 0 0 1-2.828 0l-1.707-1.707A1 1 0 0 0 8.172 19H5a3 3 0 0 1-3-3V6zm5 0a1 1 0 0 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	tooltipTextLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h10M7 11h10M7 14h4m1.707 6.293l1.707-1.707A2 2 0 0 1 15.828 18H19a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.172a2 2 0 0 1 1.414.586l1.707 1.707a1 1 0 0 0 1.414 0z"/>`
	tooltipsInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M12 2H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89l.574.862a1 1 0 0 0 1.664 0l.574-.861A2 2 0 0 1 11.07 12H12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z"/><path d="M18 10h2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-.93a2 2 0 0 0-1.664.89l-.574.862a1 1 0 0 1-1.664 0l-.574-.861A2 2 0 0 0 12.93 20H12a2 2 0 0 1-2-2v-1"/></g>`
	tooltipsLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89l.574.862a1 1 0 0 0 1.664 0l.574-.861A2 2 0 0 1 11.07 12H12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm6 8h2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-.93a2 2 0 0 0-1.664.89l-.574.862a1 1 0 0 1-1.664 0l-.574-.861A2 2 0 0 0 12.93 20H12a2 2 0 0 1-2-2v-1"/>`
	tooltipsTwoInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M20 10h-8a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89l.574.862a1 1 0 0 0 1.664 0l.574-.861A2 2 0 0 1 19.07 20H20a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2z"/><path d="M14 7V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89L7 13.5"/></g>`
	tooltipsTwoLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 10h-8a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89l.574.862a1 1 0 0 0 1.664 0l.574-.861A2 2 0 0 1 19.07 20H20a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zm-6-3V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89L7 13.5"/>`
	translateInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5zm4 3a1 1 0 0 0-2 0v1H4a1 1 0 0 0 0 2h6.432a8.323 8.323 0 0 1-1.117 3.127L7.753 8.34a1 1 0 1 0-1.506 1.32l1.837 2.098C6.7 13.231 5.1 14 4 14a1 1 0 1 0 0 2c1.806 0 3.83-1.111 5.406-2.732l.341.39a1 1 0 0 0 1.506-1.316l-.567-.648c.908-1.341 1.573-2.941 1.76-4.694H14a1 1 0 1 0 0-2H9V4zm7.894 5.553a1 1 0 0 0-1.788 0l-3 6a.998.998 0 0 0-.055.13l-1 3a1 1 0 0 0 1.898.633L13.72 17h4.558l.772 2.316a1 1 0 0 0 1.898-.632l-1-3a.998.998 0 0 0-.055-.131l-3-6zM16 12.236L17.382 15h-2.764L16 12.236z" fill="currentColor"/></g>`
	translateLineInnerSVG             = `<g fill="currentColor"><path d="M9 2a1 1 0 0 1 1 1v1h5a1 1 0 1 1 0 2h-1.546c-.222 2.38-1.212 4.52-2.5 6.341l.827 1.034a1 1 0 1 1-1.562 1.25l-.541-.677c-2.068 2.338-4.539 4.056-6.212 4.937a1 1 0 1 1-.932-1.77c1.49-.784 3.765-2.363 5.654-4.502c.074-.083.147-.168.22-.253l-2.19-2.735a1 1 0 0 1 1.562-1.25l1.865 2.332C10.581 9.27 11.25 7.686 11.442 6H3a1 1 0 0 1 0-2h5V3a1 1 0 0 1 1-1zm7 8a1 1 0 0 1 .894.553l5 10a1 1 0 1 1-1.788.894L18.882 19h-5.764l-1.224 2.447a1 1 0 1 1-1.788-.894l5-10A1 1 0 0 1 16 10zm-1.882 7h3.764L16 13.236L14.118 17z"/></g>`
	trashInnerSVG                     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11 5a1 1 0 0 0-1 1h4a1 1 0 0 0-1-1h-2zm0-2a3 3 0 0 0-3 3H4a1 1 0 0 0 0 2h1v10a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8h1a1 1 0 1 0 0-2h-4a3 3 0 0 0-3-3h-2zm0 8a1 1 0 1 0-2 0v5a1 1 0 1 0 2 0v-5zm4 0a1 1 0 1 0-2 0v5a1 1 0 1 0 2 0v-5z" fill="currentColor"/></g>`
	trashLineInnerSVG                 = `<g fill="currentColor"><path d="M4 7a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2v10a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V8a1 1 0 0 1-1-1zm3 1v10a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V8H7z"/><path d="M11 5a1 1 0 0 0-1 1v1a1 1 0 0 1-2 0V6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v1a1 1 0 1 1-2 0V6a1 1 0 0 0-1-1h-2zm-1 5a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1zm4 0a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1z"/></g>`
	trendingDownInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3.707 6.293a1 1 0 0 0-1.414 1.414l6 6a1 1 0 0 0 1.414 0L13 10.414l3.086 3.086l-2.793 2.793A1 1 0 0 0 14 18h7a1 1 0 0 0 1-1v-7a1 1 0 0 0-1.707-.707L17.5 12.086l-3.793-3.793a1 1 0 0 0-1.414 0L9 11.586L3.707 6.293z" fill="currentColor"/></g>`
	trendingDownLineInnerSVG          = `<g fill="currentColor"><path d="M2.293 6.293a1 1 0 0 1 1.414 0L9 11.586l3.293-3.293a1 1 0 0 1 1.414 0L20 14.586V10a1 1 0 1 1 2 0v7a1 1 0 0 1-1 1h-7a1 1 0 1 1 0-2h4.586L13 10.414l-3.293 3.293a1 1 0 0 1-1.414 0l-6-6a1 1 0 0 1 0-1.414z"/></g>`
	trendingUpInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3.707 17.707a1 1 0 0 1-1.414-1.414l6-6a1 1 0 0 1 1.414 0L13 13.586l3.086-3.086l-2.793-2.793A1 1 0 0 1 14 6h7a1 1 0 0 1 1 1v7a1 1 0 0 1-1.707.707L17.5 11.914l-3.793 3.793a1 1 0 0 1-1.414 0L9 12.414l-5.293 5.293z" fill="currentColor"/></g>`
	trendingUpLineInnerSVG            = `<g fill="currentColor"><path d="M2.293 17.707a1 1 0 0 0 1.414 0L9 12.414l3.293 3.293a1 1 0 0 0 1.414 0L20 9.414V14a1 1 0 1 0 2 0V7a1 1 0 0 0-1-1h-7a1 1 0 1 0 0 2h4.586L13 13.586l-3.293-3.293a1 1 0 0 0-1.414 0l-6 6a1 1 0 0 0 0 1.414z"/></g>`
	truckInnerSVG                     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2h3a1 1 0 0 1 .707.293l4 4A1 1 0 0 1 22 12v5a1 1 0 0 1-1 1h-1.17a3 3 0 1 0-5.659 0H9.829a3 3 0 1 0-5.659 0H3a1 1 0 0 1-1-1V5zm7 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm8 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4z" fill="currentColor"/></g>`
	truckLineInnerSVG                 = `<g fill="currentColor"><path d="M2 8a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v8a1 1 0 0 1-1 1H9a1 1 0 1 1 0-2h3V8a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1v2a3 3 0 0 1-3-3V8z"/><path d="M12 9a1 1 0 0 1 1-1h4a1 1 0 0 1 .707.293l3.414 3.414A3 3 0 0 1 22 13.828V14a3 3 0 0 1-3 3v-2a1 1 0 0 0 1-1v-.172a1 1 0 0 0-.293-.707L16.586 10H14v5h1a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1V9zm-5 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0z"/><path d="M17 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0z"/></g>`
	tvOldInnerSVG                     = `<g fill="none"><path fill="currentColor" d="M3 17V9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7h7a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h7zm0 0L8 3m4 4l4-4"/></g>`
	tvOldLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7h7a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h7zm0 0L8 3m4 4l4-4"/>`
	umbrellaInnerSVG                  = `<g fill="none"><path fill="currentColor" d="M12 4c-7.2 0-9 6-9 9h18c0-3-1.8-9-9-9z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4c-7.2 0-9 6-9 9h9m0-9c7.2 0 9 6 9 9h-9m0-9V3m0 10v5c0 1 .6 3 3 3s3-2 3-3"/></g>`
	umbrellaLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4c-7.2 0-9 6-9 9h9m0-9c7.2 0 9 6 9 9h-9m0-9V3m0 10v5c0 1 .6 3 3 3s3-2 3-3"/>`
	underlineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5v6a4 4 0 0 0 4 4v0a4 4 0 0 0 4-4V5M5 19h14"/>`
	underlineLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5v6a4 4 0 0 0 4 4v0a4 4 0 0 0 4-4V5M5 19h14"/>`
	underlineTwoInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19h14M7 15l1.5-3m8.5 3l-1.5-3m0 0L12 5l-3.5 7m7 0h-7"/>`
	underlineTwoLineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19h14M7 15l1.5-3m8.5 3l-1.5-3m0 0L12 5l-3.5 7m7 0h-7"/>`
	undoInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9v5h5m11 2c-.497-4.5-3.367-8-8-8c-2.73 0-5.929 2.268-7.294 5.5"/>`
	undoLineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9v5h5m11 2c-.497-4.5-3.367-8-8-8c-2.73 0-5.929 2.268-7.294 5.5"/>`
	unlockOpenInnerSVG                = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.5 4.99C14.958 3.944 13.904 3 12 3C8.8 3 8 5.667 8 7v3"/><path fill="currentColor" fill-rule="evenodd" d="M3 12a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-7zm10 2a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0v-3z" clip-rule="evenodd"/></g>`
	unlockOpenLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10H6a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2H8zm0 0V7c0-1.333.8-4 4-4c1.904 0 2.958.944 3.5 1.99M12 14v3"/>`
	usbInnerSVG                       = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6v1m4-1v1m-6 4h8a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm2 8H9a4 4 0 0 1-4-4v-3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3a4 4 0 0 1-4 4h-1m-4 0v2m0-2h4m0 0v2"/><path fill="currentColor" d="M5 12a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3a4 4 0 0 1-4 4H9a4 4 0 0 1-4-4v-3z"/></g>`
	usbLineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6v1m4-1v1m-6 4h8a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1zm2 8H9a4 4 0 0 1-4-4v-3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3a4 4 0 0 1-4 4h-1m-4 0v2m0-2h4m0 0v2"/>`
	userInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="8" r="5" fill="currentColor"/><path d="M20 21a8 8 0 1 0-16 0"/><path fill="currentColor" d="M12 13a8 8 0 0 0-8 8h16a8 8 0 0 0-8-8z"/></g>`
	userAddInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 9a6 6 0 1 1 9.642 4.769C14.186 14.907 16 17.208 16 20a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1c0-2.792 1.814-5.093 4.358-6.231A5.99 5.99 0 0 1 2 9zm19 0a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V9z" fill="currentColor"/></g>`
	userAddLineInnerSVG               = `<g fill="currentColor"><path d="M21 9a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V9zM8 5a4 4 0 1 0 0 8a4 4 0 0 0 0-8zM2 9a6 6 0 1 1 12 0A6 6 0 0 1 2 9z"/><path d="M2.124 19h11.752c-.547-2.197-2.86-4-5.876-4c-3.016 0-5.329 1.803-5.876 4zM0 20c0-4.005 3.732-7 8-7s8 2.995 8 7a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1z"/></g>`
	userBoxInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="3" fill="currentColor"/><path fill="currentColor" d="M12 13c-2.761 0-5 1.79-5 4h10c0-2.21-2.239-4-5-4z"/><rect width="18" height="18" x="3" y="3" rx="3"/></g>`
	userBoxLineInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="3"/><path d="M17 17c0-2.21-2.239-4-5-4s-5 1.79-5 4"/><rect width="18" height="18" x="3" y="3" rx="3"/></g>`
	userCircleInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5.2 19.332A9.973 9.973 0 0 1 2 12C2 6.477 6.477 2 12 2s10 4.477 10 10a9.973 9.973 0 0 1-3.2 7.332a7.014 7.014 0 0 0-3.55-4.533a5 5 0 1 0-6.502 0A7.014 7.014 0 0 0 5.2 19.332zm1.81 1.336A9.954 9.954 0 0 0 12 22c1.817 0 3.52-.485 4.99-1.332a5 5 0 0 0-9.98 0zM12 8a3 3 0 1 0 0 6a3 3 0 0 0 0-6z" fill="currentColor"/></g>`
	userCircleLineInnerSVG            = `<g fill="currentColor"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12z"/><path d="M12 8a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0z"/><path d="M12 16a5.003 5.003 0 0 0-4.716 3.333a1 1 0 1 1-1.885-.666a7.002 7.002 0 0 1 13.202 0a1 1 0 1 1-1.885.666A5.002 5.002 0 0 0 12 16z"/></g>`
	userGroupInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 5a4 4 0 0 0-2.906 6.75a5.998 5.998 0 0 0-2.974 6.449A1 1 0 0 0 7.1 19h9.8a1 1 0 0 0 .98-.801a5.998 5.998 0 0 0-2.974-6.45A4 4 0 0 0 12 5zm4.584 5.779C16.85 10.243 17 9.639 17 9s-.15-1.243-.416-1.779A3 3 0 1 1 21.4 10.8A4 4 0 0 1 23 14v1a1 1 0 0 1-1 1h-3.083a6.006 6.006 0 0 0-3.012-4.25a4.01 4.01 0 0 0 .651-.917a3.56 3.56 0 0 1 .044-.033l-.016-.021zm-8.49.97a4.01 4.01 0 0 1-.65-.916A3.902 3.902 0 0 0 7.4 10.8l.016-.021A3.984 3.984 0 0 1 7 9c0-.639.15-1.243.416-1.779A3 3 0 1 0 2.6 10.8a3.994 3.994 0 0 0-1.372 4.534A1 1 0 0 0 2.17 16h2.912a6.006 6.006 0 0 1 3.011-4.25z" fill="currentColor"/></g>`
	userGroupLineInnerSVG             = `<g fill="currentColor"><path d="M12 7a2 2 0 1 0 0 4a2 2 0 0 0 0-4zM8 9a4 4 0 1 1 8 0a4 4 0 0 1-8 0zm11-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0zM5 8a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM2 9a3 3 0 1 1 6 0a3 3 0 0 1-6 0z"/><path d="M17.268 13h3.464a2 2 0 0 0-3.464 0zM15 14a4 4 0 0 1 8 0v1h-7l-1-1zM3.268 13h3.464a2 2 0 0 0-3.464 0zM1 14a4 4 0 0 1 8 0l-1 1H1v-1z"/><path d="M12 13a4 4 0 0 0-4 4h8a4 4 0 0 0-4-4zm-6 4a6 6 0 1 1 11.88 1.199l-.163.801H6.283l-.163-.801A6.022 6.022 0 0 1 6 17z"/></g>`
	userLineInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="8" r="5"/><path d="M20 21a8 8 0 1 0-16 0m16 0a8 8 0 1 0-16 0"/></g>`
	userRemoveInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 9a6 6 0 1 1 9.642 4.769C14.186 14.907 16 17.208 16 20a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1c0-2.792 1.814-5.093 4.358-6.231A5.99 5.99 0 0 1 2 9zm15 2a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6z" fill="currentColor"/></g>`
	userRemoveLineInnerSVG            = `<g fill="currentColor"><path d="M17 11a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6zM8 5a4 4 0 1 0 0 8a4 4 0 0 0 0-8zM2 9a6 6 0 1 1 12 0A6 6 0 0 1 2 9z"/><path d="M2.124 19h11.752c-.547-2.197-2.86-4-5.876-4c-3.016 0-5.329 1.803-5.876 4zM0 20c0-4.005 3.732-7 8-7s8 2.995 8 7a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1z"/></g>`
	usersInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="9" r="4" fill="currentColor"/><path fill="currentColor" d="M9 13c-3.866 0-7 2.686-7 6h14c0-3.314-3.134-6-7-6z"/><path d="M15 13a4 4 0 1 0-3-6.646m0 5.411c.897.942 2.193 1.235 3 1.235c3.866 0 7 2.686 7 6h-6"/></g>`
	usersLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="9" r="4"/><path d="M16 19c0-3.314-3.134-6-7-6s-7 2.686-7 6m13-6a4 4 0 1 0-3-6.646"/><path d="M22 19c0-3.314-3.134-6-7-6c-.807 0-2.103-.293-3-1.235"/></g>`
	uxCircleInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm13.8-2.6a1 1 0 0 0-1.6 1.2l1.05 1.4l-1.05 1.4a1 1 0 0 0 1.6 1.2l.7-.933l.7.933a1 1 0 0 0 1.6-1.2L16.75 12l1.05-1.4a1 1 0 0 0-1.6-1.2l-.7.933l-.7-.933zM8 10a1 1 0 1 0-2 0v2c0 .493.14 1.211.588 1.834C7.074 14.51 7.874 15 9 15s1.926-.492 2.412-1.166c.448-.623.588-1.34.588-1.834v-2a1 1 0 1 0-2 0v2c0 .173-.06.456-.212.666c-.114.159-.314.334-.788.334c-.474 0-.674-.175-.788-.334A1.239 1.239 0 0 1 8 12v-2z" clip-rule="evenodd"/>`
	uxCircleLineInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 10v2c0 .667.4 2 2 2s2-1.333 2-2v-2m3 0l1.5 2m1.5 2l-1.5-2m0 0l1.5-2m-1.5 2L14 14"/><circle cx="12" cy="12" r="10"/></g>`
	videoInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M5 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-1.586l2.293 2.293A1 1 0 0 0 22 16V8a1 1 0 0 0-1.707-.707L18 9.586V8a3 3 0 0 0-3-3H5z" clip-rule="evenodd"/>`
	videoCameraInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3v-1.586l2.44 2.44c.944.945 2.56.275 2.56-1.061V8.207c0-1.336-1.616-2.006-2.56-1.06L17 9.585V8a3 3 0 0 0-3-3H5z" fill="currentColor"/></g>`
	videoCameraLineInnerSVG           = `<g fill="currentColor"><path d="M2 8a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v1.586l2.44-2.44c.944-.944 2.56-.275 2.56 1.061v7.586c0 1.336-1.616 2.006-2.56 1.06L17 14.415V16a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8zm15.414 4L20 14.586V9.414L17.414 12zM5 7a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H5z"/></g>`
	videoLineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-4zm0 0l4-4v8l-4-4z"/>`
	videoMinusInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v1.586l2.293-2.293A1 1 0 0 1 22 8v8a1 1 0 0 1-1.707.707L18 14.414V16a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8zm5 3a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`
	videoMinusLineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-4zm0 0l4-4v8l-4-4zm-4 0H7"/>`
	videoPlusInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v1.586l2.293-2.293A1 1 0 0 1 22 8v8a1 1 0 0 1-1.707.707L18 14.414V16a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8zm9 1a1 1 0 1 0-2 0v2H7a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V9z" clip-rule="evenodd"/>`
	videoPlusLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-4zm0 0l4-4v8l-4-4zm-7-3v3m0 3v-3m0 0h3m-3 0H7"/>`
	viewBoardsInnerSVG                = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 7a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7zm14 0a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2V7zm-5-2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2z" fill="currentColor"/></g>`
	viewBoardsLineInnerSVG            = `<g fill="currentColor"><path d="M5 7a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H5zM2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8z"/><path d="M9 5a1 1 0 0 1 1 1v12a1 1 0 1 1-2 0V6a1 1 0 0 1 1-1zm6 0a1 1 0 0 1 1 1v12a1 1 0 1 1-2 0V6a1 1 0 0 1 1-1z"/></g>`
	viewColumnsInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M7 5H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h2V5zm2 14h6V5H9v14zm8-14v14h2a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3h-2z" clip-rule="evenodd"/>`
	viewColumnsLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6H5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h4M9 6v12M9 6h6M9 18h6m0 0h4a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-4m0 12V6"/>`
	viewGridInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 7a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2zm0 5a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2zm2 3a2 2 0 1 0 0 4h2a2 2 0 1 0 0-4H4zm5-8a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4h-2a2 2 0 0 1-2-2zm2 3a2 2 0 1 0 0 4h2a2 2 0 1 0 0-4h-2zm-2 7a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4h-2a2 2 0 0 1-2-2zm9-12a2 2 0 1 0 0 4h2a2 2 0 1 0 0-4h-2zm-2 7a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4h-2a2 2 0 0 1-2-2zm2 3a2 2 0 1 0 0 4h2a2 2 0 1 0 0-4h-2z" fill="currentColor"/></g>`
	viewGridLineInnerSVG              = `<g fill="currentColor"><path d="M5 7a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H5zM2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8z"/><path d="M9 5a1 1 0 0 1 1 1v3h4V6a1 1 0 1 1 2 0v3h5a1 1 0 1 1 0 2h-5v2h5a1 1 0 1 1 0 2h-5v3a1 1 0 1 1-2 0v-3h-4v3a1 1 0 1 1-2 0v-3H3a1 1 0 1 1 0-2h5v-2H3a1 1 0 1 1 0-2h5V6a1 1 0 0 1 1-1zm1 6v2h4v-2h-4z"/></g>`
	viewListInnerSVG                  = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M4 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H4zm1 2a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2H5zm5 0a1 1 0 0 0 0 2h9a1 1 0 1 0 0-2h-9zm-5 4a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H5zm5 0a1 1 0 1 0 0 2h9a1 1 0 1 0 0-2h-9zm-5 4a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H5zm5 0a1 1 0 1 0 0 2h9a1 1 0 1 0 0-2h-9z" fill="currentColor"/></g>`
	viewListLineInnerSVG              = `<g fill="currentColor"><path d="M2 6a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H3a1 1 0 0 1-1-1zm6 0a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1zm-6 4a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H3a1 1 0 0 1-1-1zm6 0a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1zm-6 4a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1zm6 0a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1zm-6 4a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1zm6 0a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1z"/></g>`
	viewRowsInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3H2zm0 2v4h20v-4H2zm20 6H2a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3z" clip-rule="evenodd"/>`
	viewRowsLineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6M3 10V8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 10h18M3 14h18"/>`
	watchInnerSVG                     = `<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M5 12a7 7 0 1 1 14 0a7 7 0 0 1-14 0zm8-3a1 1 0 1 0-2 0v3a1 1 0 0 0 .553.894l2 1a1 1 0 1 0 .894-1.788L13 11.382V9z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6.5V3a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v3.5m6 11V21a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-3.5"/></g>`
	watchLineInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 9v3l2 1m4-1a6 6 0 1 1-12 0a6 6 0 0 1 12 0z"/><path d="M15 6.5V3a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v3.5m6 11V21a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-3.5"/></g>`
	wifiInnerSVG                      = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5zm6.994 5c-3.202 0-6.122 1.348-8.29 3.569a1 1 0 1 1-1.431-1.397C4.787 5.597 8.21 4 11.994 4c3.783 0 7.207 1.597 9.721 4.172a1 1 0 1 1-1.43 1.397C18.115 7.349 15.196 6 11.995 6zm0 5c-2.023 0-3.883.78-5.335 2.095a1 1 0 1 1-1.342-1.483C7.11 9.99 9.44 9 11.994 9c2.554 0 4.885.99 6.677 2.612a1 1 0 1 1-1.342 1.483C15.877 11.78 14.017 11 11.994 11zm0 5a4.3 4.3 0 0 0-2.511.814a1 1 0 1 1-1.162-1.628A6.3 6.3 0 0 1 11.994 14c1.363 0 2.627.44 3.673 1.186a1 1 0 1 1-1.162 1.628a4.3 4.3 0 0 0-2.51-.814zm-1 3a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2h-.01a1 1 0 0 1-1-1z" fill="currentColor"/></g>`
	wifiLineInnerSVG                  = `<g fill="currentColor"><path d="M2.041 10.643A12.97 12.97 0 0 1 12 6a12.97 12.97 0 0 1 9.959 4.643a1 1 0 1 0 1.53-1.286A14.97 14.97 0 0 0 12 4A14.97 14.97 0 0 0 .51 9.357a1 1 0 0 0 1.531 1.286zM5.573 13.7A8.97 8.97 0 0 1 12 11a8.97 8.97 0 0 1 6.427 2.7a1 1 0 0 0 1.428-1.4A10.97 10.97 0 0 0 12 9a10.97 10.97 0 0 0-7.856 3.3a1 1 0 1 0 1.429 1.4zm3.663 3.133A4.972 4.972 0 0 1 12 16c1.024 0 1.973.307 2.764.833a1 1 0 1 0 1.107-1.666A6.972 6.972 0 0 0 12 14c-1.43 0-2.762.43-3.871 1.167a1 1 0 1 0 1.107 1.666zM12 18a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H12z"/></g>`
	yenCircleInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm8.32-5.573a1 1 0 1 0-1.64 1.146L10.08 11H9a1 1 0 1 0 0 2h2v1H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-1h2a1 1 0 1 0 0-2h-1.08l2.4-3.427a1 1 0 0 0-1.64-1.146L12 10.257l-2.68-3.83z" clip-rule="evenodd"/>`
	yenCircleLineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10zM8.5 7l3.5 5m0 0l3.5-5M12 12H9m3 0v3m0-3h3m-3 6v-3m0 0h3m-3 0H9"/>`
	zoomInInnerSVG                    = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 11a9 9 0 1 1 16.032 5.618l3.675 3.675a1 1 0 0 1-1.414 1.414l-3.675-3.675A9 9 0 0 1 2 11zm10-3a1 1 0 1 0-2 0v2H8a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V8z" fill="currentColor"/></g>`
	zoomInLineInnerSVG                = `<g fill="currentColor"><path d="M4 11a7 7 0 1 1 14 0a7 7 0 0 1-14 0zm7-9a9 9 0 1 0 5.618 16.032l3.675 3.675a1 1 0 0 0 1.414-1.414l-3.675-3.675A9 9 0 0 0 11 2zm1 6a1 1 0 1 0-2 0v2H8a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V8z"/></g>`
	zoomOutInnerSVG                   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 11a9 9 0 1 1 16.032 5.618l3.675 3.675a1 1 0 0 1-1.414 1.414l-3.675-3.675A9 9 0 0 1 2 11zm6-1a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H8z" fill="currentColor"/></g>`
	zoomOutLineInnerSVG               = `<g fill="currentColor"><path d="M4 11a7 7 0 1 1 14 0a7 7 0 0 1-14 0zm7-9a9 9 0 1 0 5.618 16.032l3.675 3.675a1 1 0 0 0 1.414-1.414l-3.675-3.675A9 9 0 0 0 11 2zm-3 8a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H8z"/></g>`
)

var sharedIconAttrs = engine.Attrs{"width": "1em", "height": "1em"}

func AcademicCap(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		academicCapInnerSVG,
		children,
	)
}

func AcademicCapLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		academicCapLineInnerSVG,
		children,
	)
}

func AddColumn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addColumnInnerSVG,
		children,
	)
}

func AddColumnLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addColumnLineInnerSVG,
		children,
	)
}

func AddRow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addRowInnerSVG,
		children,
	)
}

func AddRowLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addRowLineInnerSVG,
		children,
	)
}

func Adjustments(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		adjustmentsInnerSVG,
		children,
	)
}

func AdjustmentsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		adjustmentsLineInnerSVG,
		children,
	)
}

func Airplane(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		airplaneInnerSVG,
		children,
	)
}

func AirplaneFlightTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		airplaneFlightTwoInnerSVG,
		children,
	)
}

func AirplaneFlightTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		airplaneFlightTwoLineInnerSVG,
		children,
	)
}

func AirplaneLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		airplaneLineInnerSVG,
		children,
	)
}

func AlertCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertCircleInnerSVG,
		children,
	)
}

func AlertCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertCircleLineInnerSVG,
		children,
	)
}

func AlignBottom(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignBottomInnerSVG,
		children,
	)
}

func AlignBottomLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignBottomLineInnerSVG,
		children,
	)
}

func AlignHorizontalCenter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignHorizontalCenterInnerSVG,
		children,
	)
}

func AlignHorizontalCenterLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignHorizontalCenterLineInnerSVG,
		children,
	)
}

func AlignLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignLeftInnerSVG,
		children,
	)
}

func AlignLeftLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignLeftLineInnerSVG,
		children,
	)
}

func AlignRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignRightInnerSVG,
		children,
	)
}

func AlignRightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignRightLineInnerSVG,
		children,
	)
}

func AlignTop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignTopInnerSVG,
		children,
	)
}

func AlignTopLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignTopLineInnerSVG,
		children,
	)
}

func AlignVerticalCenter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignVerticalCenterInnerSVG,
		children,
	)
}

func AlignVerticalCenterLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignVerticalCenterLineInnerSVG,
		children,
	)
}

func Analytics(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsInnerSVG,
		children,
	)
}

func AnalyticsDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsDeleteInnerSVG,
		children,
	)
}

func AnalyticsDeleteLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsDeleteLineInnerSVG,
		children,
	)
}

func AnalyticsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsLineInnerSVG,
		children,
	)
}

func AnalyticsPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsPlusInnerSVG,
		children,
	)
}

func AnalyticsPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsPlusLineInnerSVG,
		children,
	)
}

func AnalyticsRestricted(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsRestrictedInnerSVG,
		children,
	)
}

func AnalyticsRestrictedLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsRestrictedLineInnerSVG,
		children,
	)
}

func Annotation(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		annotationInnerSVG,
		children,
	)
}

func AnnotationLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		annotationLineInnerSVG,
		children,
	)
}

func Applications(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		applicationsInnerSVG,
		children,
	)
}

func ApplicationsAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		applicationsAddInnerSVG,
		children,
	)
}

func ApplicationsAddLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		applicationsAddLineInnerSVG,
		children,
	)
}

func ApplicationsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		applicationsLineInnerSVG,
		children,
	)
}

func Archive(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		archiveInnerSVG,
		children,
	)
}

func ArchiveLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		archiveLineInnerSVG,
		children,
	)
}

func ArrowCircleDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleDownInnerSVG,
		children,
	)
}

func ArrowCircleDownLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleDownLineInnerSVG,
		children,
	)
}

func ArrowCircleLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleLeftInnerSVG,
		children,
	)
}

func ArrowCircleLeftLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleLeftLineInnerSVG,
		children,
	)
}

func ArrowCircleRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleRightInnerSVG,
		children,
	)
}

func ArrowCircleRightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleRightLineInnerSVG,
		children,
	)
}

func ArrowCircleUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleUpInnerSVG,
		children,
	)
}

func ArrowCircleUpLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleUpLineInnerSVG,
		children,
	)
}

func ArrowDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownInnerSVG,
		children,
	)
}

func ArrowDownCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownCircleInnerSVG,
		children,
	)
}

func ArrowDownCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownCircleLineInnerSVG,
		children,
	)
}

func ArrowDownLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownLineInnerSVG,
		children,
	)
}

func ArrowLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLeftInnerSVG,
		children,
	)
}

func ArrowLeftCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLeftCircleInnerSVG,
		children,
	)
}

func ArrowLeftCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLeftCircleLineInnerSVG,
		children,
	)
}

func ArrowLeftLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLeftLineInnerSVG,
		children,
	)
}

func ArrowRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowRightInnerSVG,
		children,
	)
}

func ArrowRightCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowRightCircleInnerSVG,
		children,
	)
}

func ArrowRightCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowRightCircleLineInnerSVG,
		children,
	)
}

func ArrowRightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowRightLineInnerSVG,
		children,
	)
}

func ArrowUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpInnerSVG,
		children,
	)
}

func ArrowUpCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpCircleInnerSVG,
		children,
	)
}

func ArrowUpCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpCircleLineInnerSVG,
		children,
	)
}

func ArrowUpLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpLineInnerSVG,
		children,
	)
}

func ArrowsCollapseFull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsCollapseFullInnerSVG,
		children,
	)
}

func ArrowsCollapseFullLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsCollapseFullLineInnerSVG,
		children,
	)
}

func ArrowsExpand(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsExpandInnerSVG,
		children,
	)
}

func ArrowsExpandFull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsExpandFullInnerSVG,
		children,
	)
}

func ArrowsExpandFullLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsExpandFullLineInnerSVG,
		children,
	)
}

func ArrowsExpandLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsExpandLineInnerSVG,
		children,
	)
}

func Article(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		articleInnerSVG,
		children,
	)
}

func ArticleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		articleLineInnerSVG,
		children,
	)
}

func ArticleSearch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		articleSearchInnerSVG,
		children,
	)
}

func ArticleSearchLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		articleSearchLineInnerSVG,
		children,
	)
}

func AtSymbol(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		atSymbolInnerSVG,
		children,
	)
}

func AtSymbolLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		atSymbolLineInnerSVG,
		children,
	)
}

func AtomTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		atomTwoInnerSVG,
		children,
	)
}

func AtomTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		atomTwoLineInnerSVG,
		children,
	)
}

func Attachment(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		attachmentInnerSVG,
		children,
	)
}

func AttachmentLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		attachmentLineInnerSVG,
		children,
	)
}

func Award(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		awardInnerSVG,
		children,
	)
}

func AwardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		awardLineInnerSVG,
		children,
	)
}

func BackCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backCircleInnerSVG,
		children,
	)
}

func BackCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backCircleLineInnerSVG,
		children,
	)
}

func Backspace(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backspaceInnerSVG,
		children,
	)
}

func BackspaceLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backspaceLineInnerSVG,
		children,
	)
}

func BackwardCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backwardCircleInnerSVG,
		children,
	)
}

func BackwardCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backwardCircleLineInnerSVG,
		children,
	)
}

func BackwardStartCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backwardStartCircleInnerSVG,
		children,
	)
}

func BackwardStartCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backwardStartCircleLineInnerSVG,
		children,
	)
}

func BadgeCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		badgeCheckInnerSVG,
		children,
	)
}

func BadgeCheckLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		badgeCheckLineInnerSVG,
		children,
	)
}

func Ban(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		banInnerSVG,
		children,
	)
}

func BanLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		banLineInnerSVG,
		children,
	)
}

func BandAids(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bandAidsInnerSVG,
		children,
	)
}

func BandAidsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bandAidsLineInnerSVG,
		children,
	)
}

func BarcodeTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		barcodeTwoInnerSVG,
		children,
	)
}

func BarcodeTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		barcodeTwoLineInnerSVG,
		children,
	)
}

func BasketTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		basketTwoInnerSVG,
		children,
	)
}

func BasketTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		basketTwoLineInnerSVG,
		children,
	)
}

func BathShower(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bathShowerInnerSVG,
		children,
	)
}

func BathShowerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bathShowerLineInnerSVG,
		children,
	)
}

func Battery(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryInnerSVG,
		children,
	)
}

func BatteryFull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryFullInnerSVG,
		children,
	)
}

func BatteryFullLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryFullLineInnerSVG,
		children,
	)
}

func BatteryHalf(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryHalfInnerSVG,
		children,
	)
}

func BatteryHalfLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryHalfLineInnerSVG,
		children,
	)
}

func BatteryLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryLineInnerSVG,
		children,
	)
}

func BatteryLow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryLowInnerSVG,
		children,
	)
}

func BatteryLowLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryLowLineInnerSVG,
		children,
	)
}

func Beach(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beachInnerSVG,
		children,
	)
}

func BeachLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beachLineInnerSVG,
		children,
	)
}

func Beaker(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beakerInnerSVG,
		children,
	)
}

func BeakerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beakerLineInnerSVG,
		children,
	)
}

func Bell(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bellInnerSVG,
		children,
	)
}

func BellLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bellLineInnerSVG,
		children,
	)
}

func BitcoinCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bitcoinCircleInnerSVG,
		children,
	)
}

func BitcoinCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bitcoinCircleLineInnerSVG,
		children,
	)
}

func Bluetooth(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bluetoothInnerSVG,
		children,
	)
}

func BluetoothLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bluetoothLineInnerSVG,
		children,
	)
}

func Bold(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		boldInnerSVG,
		children,
	)
}

func BoldLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		boldLineInnerSVG,
		children,
	)
}

func Book(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookInnerSVG,
		children,
	)
}

func BookLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookLineInnerSVG,
		children,
	)
}

func BookMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookMinusInnerSVG,
		children,
	)
}

func BookMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookMinusLineInnerSVG,
		children,
	)
}

func BookOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookOpenInnerSVG,
		children,
	)
}

func BookOpenLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookOpenLineInnerSVG,
		children,
	)
}

func BookPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookPlusInnerSVG,
		children,
	)
}

func BookPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookPlusLineInnerSVG,
		children,
	)
}

func Bookmark(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkInnerSVG,
		children,
	)
}

func BookmarkBook(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkBookInnerSVG,
		children,
	)
}

func BookmarkBookLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkBookLineInnerSVG,
		children,
	)
}

func BookmarkLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkLineInnerSVG,
		children,
	)
}

func BookmarkMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkMinusInnerSVG,
		children,
	)
}

func BookmarkMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkMinusLineInnerSVG,
		children,
	)
}

func BookmarkPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkPlusInnerSVG,
		children,
	)
}

func BookmarkPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkPlusLineInnerSVG,
		children,
	)
}

func Box(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		boxInnerSVG,
		children,
	)
}

func BoxLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		boxLineInnerSVG,
		children,
	)
}

func Briefcase(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseInnerSVG,
		children,
	)
}

func BriefcaseLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseLineInnerSVG,
		children,
	)
}

func Browser(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		browserInnerSVG,
		children,
	)
}

func BrowserCookie(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		browserCookieInnerSVG,
		children,
	)
}

func BrowserCookieLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		browserCookieLineInnerSVG,
		children,
	)
}

func BrowserLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		browserLineInnerSVG,
		children,
	)
}

func BugTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bugTwoInnerSVG,
		children,
	)
}

func BugTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bugTwoLineInnerSVG,
		children,
	)
}

func Burger(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		burgerInnerSVG,
		children,
	)
}

func BurgerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		burgerLineInnerSVG,
		children,
	)
}

func Bus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		busInnerSVG,
		children,
	)
}

func BusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		busLineInnerSVG,
		children,
	)
}

func Cake(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cakeInnerSVG,
		children,
	)
}

func CakeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cakeLineInnerSVG,
		children,
	)
}

func Calculator(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calculatorInnerSVG,
		children,
	)
}

func CalculatorLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calculatorLineInnerSVG,
		children,
	)
}

func Calendar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarInnerSVG,
		children,
	)
}

func CalendarLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarLineInnerSVG,
		children,
	)
}

func CalendarPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarPlusInnerSVG,
		children,
	)
}

func CalendarPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarPlusLineInnerSVG,
		children,
	)
}

func Camera(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraInnerSVG,
		children,
	)
}

func CameraLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraLineInnerSVG,
		children,
	)
}

func CameraOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraOffInnerSVG,
		children,
	)
}

func CameraOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraOffLineInnerSVG,
		children,
	)
}

func Car(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		carInnerSVG,
		children,
	)
}

func CarLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		carLineInnerSVG,
		children,
	)
}

func Cash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cashInnerSVG,
		children,
	)
}

func CashLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cashLineInnerSVG,
		children,
	)
}

func CentCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		centCircleInnerSVG,
		children,
	)
}

func CentCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		centCircleLineInnerSVG,
		children,
	)
}

func ChartBar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartBarInnerSVG,
		children,
	)
}

func ChartBarLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartBarLineInnerSVG,
		children,
	)
}

func ChartPie(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartPieInnerSVG,
		children,
	)
}

func ChartPieLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartPieLineInnerSVG,
		children,
	)
}

func Chat(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatInnerSVG,
		children,
	)
}

func ChatLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatLineInnerSVG,
		children,
	)
}

func ChatSignal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatSignalInnerSVG,
		children,
	)
}

func ChatSignalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatSignalLineInnerSVG,
		children,
	)
}

func ChatStatus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatStatusInnerSVG,
		children,
	)
}

func ChatStatusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatStatusLineInnerSVG,
		children,
	)
}

func ChatText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatTextInnerSVG,
		children,
	)
}

func ChatTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatTextLineInnerSVG,
		children,
	)
}

func ChatTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatTwoInnerSVG,
		children,
	)
}

func ChatTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatTwoLineInnerSVG,
		children,
	)
}

func ChatTwoText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatTwoTextInnerSVG,
		children,
	)
}

func ChatTwoTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatTwoTextLineInnerSVG,
		children,
	)
}

func Chats(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatsInnerSVG,
		children,
	)
}

func ChatsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatsLineInnerSVG,
		children,
	)
}

func ChatsTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatsTwoInnerSVG,
		children,
	)
}

func ChatsTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatsTwoLineInnerSVG,
		children,
	)
}

func Check(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkInnerSVG,
		children,
	)
}

func CheckCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkCircleInnerSVG,
		children,
	)
}

func CheckCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkCircleLineInnerSVG,
		children,
	)
}

func CheckLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkLineInnerSVG,
		children,
	)
}

func CheckboxList(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkboxListInnerSVG,
		children,
	)
}

func CheckboxListDetail(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkboxListDetailInnerSVG,
		children,
	)
}

func CheckboxListDetailLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkboxListDetailLineInnerSVG,
		children,
	)
}

func CheckboxListLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkboxListLineInnerSVG,
		children,
	)
}

func Cheese(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cheeseInnerSVG,
		children,
	)
}

func CheeseLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cheeseLineInnerSVG,
		children,
	)
}

func ChevronDoubleDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDoubleDownInnerSVG,
		children,
	)
}

func ChevronDoubleDownLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDoubleDownLineInnerSVG,
		children,
	)
}

func ChevronDoubleLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDoubleLeftInnerSVG,
		children,
	)
}

func ChevronDoubleLeftLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDoubleLeftLineInnerSVG,
		children,
	)
}

func ChevronDoubleRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDoubleRightInnerSVG,
		children,
	)
}

func ChevronDoubleRightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDoubleRightLineInnerSVG,
		children,
	)
}

func ChevronDoubleUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDoubleUpInnerSVG,
		children,
	)
}

func ChevronDoubleUpLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDoubleUpLineInnerSVG,
		children,
	)
}

func ChevronDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDownInnerSVG,
		children,
	)
}

func ChevronDownCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDownCircleInnerSVG,
		children,
	)
}

func ChevronDownCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDownCircleLineInnerSVG,
		children,
	)
}

func ChevronDownLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDownLineInnerSVG,
		children,
	)
}

func ChevronLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronLeftInnerSVG,
		children,
	)
}

func ChevronLeftCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronLeftCircleInnerSVG,
		children,
	)
}

func ChevronLeftCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronLeftCircleLineInnerSVG,
		children,
	)
}

func ChevronLeftLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronLeftLineInnerSVG,
		children,
	)
}

func ChevronRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronRightInnerSVG,
		children,
	)
}

func ChevronRightCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronRightCircleInnerSVG,
		children,
	)
}

func ChevronRightCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronRightCircleLineInnerSVG,
		children,
	)
}

func ChevronRightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronRightLineInnerSVG,
		children,
	)
}

func ChevronUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronUpInnerSVG,
		children,
	)
}

func ChevronUpCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronUpCircleInnerSVG,
		children,
	)
}

func ChevronUpCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronUpCircleLineInnerSVG,
		children,
	)
}

func ChevronUpLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronUpLineInnerSVG,
		children,
	)
}

func Chip(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chipInnerSVG,
		children,
	)
}

func ChipLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chipLineInnerSVG,
		children,
	)
}

func Chromecast(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chromecastInnerSVG,
		children,
	)
}

func ChromecastLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chromecastLineInnerSVG,
		children,
	)
}

func Church(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		churchInnerSVG,
		children,
	)
}

func ChurchLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		churchLineInnerSVG,
		children,
	)
}

func Clipboard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardInnerSVG,
		children,
	)
}

func ClipboardCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardCheckInnerSVG,
		children,
	)
}

func ClipboardCheckLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardCheckLineInnerSVG,
		children,
	)
}

func ClipboardCopy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardCopyInnerSVG,
		children,
	)
}

func ClipboardCopyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardCopyLineInnerSVG,
		children,
	)
}

func ClipboardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardLineInnerSVG,
		children,
	)
}

func ClipboardList(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardListInnerSVG,
		children,
	)
}

func ClipboardListLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardListLineInnerSVG,
		children,
	)
}

func ClipboardMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardMinusInnerSVG,
		children,
	)
}

func ClipboardMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardMinusLineInnerSVG,
		children,
	)
}

func ClipboardPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardPlusInnerSVG,
		children,
	)
}

func ClipboardPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardPlusLineInnerSVG,
		children,
	)
}

func Clock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clockInnerSVG,
		children,
	)
}

func ClockLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clockLineInnerSVG,
		children,
	)
}

func ClockPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clockPlusInnerSVG,
		children,
	)
}

func ClockPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clockPlusLineInnerSVG,
		children,
	)
}

func Close(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeInnerSVG,
		children,
	)
}

func CloseCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeCircleInnerSVG,
		children,
	)
}

func CloseCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeCircleLineInnerSVG,
		children,
	)
}

func CloseLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeLineInnerSVG,
		children,
	)
}

func Cloud(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudInnerSVG,
		children,
	)
}

func CloudDownload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownloadInnerSVG,
		children,
	)
}

func CloudDownloadLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownloadLineInnerSVG,
		children,
	)
}

func CloudLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudLineInnerSVG,
		children,
	)
}

func CloudUpload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUploadInnerSVG,
		children,
	)
}

func CloudUploadLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUploadLineInnerSVG,
		children,
	)
}

func Code(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeInnerSVG,
		children,
	)
}

func CodeBlock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeBlockInnerSVG,
		children,
	)
}

func CodeBlockLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeBlockLineInnerSVG,
		children,
	)
}

func CodeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeLineInnerSVG,
		children,
	)
}

func Cog(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cogInnerSVG,
		children,
	)
}

func CogLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cogLineInnerSVG,
		children,
	)
}

func Coins(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coinsInnerSVG,
		children,
	)
}

func CoinsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coinsLineInnerSVG,
		children,
	)
}

func Collection(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		collectionInnerSVG,
		children,
	)
}

func CollectionLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		collectionLineInnerSVG,
		children,
	)
}

func ColorSwatch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		colorSwatchInnerSVG,
		children,
	)
}

func ColorSwatchLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		colorSwatchLineInnerSVG,
		children,
	)
}

func Comet(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cometInnerSVG,
		children,
	)
}

func CometLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cometLineInnerSVG,
		children,
	)
}

func Comment(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentInnerSVG,
		children,
	)
}

func CommentLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentLineInnerSVG,
		children,
	)
}

func CommentText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentTextInnerSVG,
		children,
	)
}

func CommentTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentTextLineInnerSVG,
		children,
	)
}

func CommentTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentTwoInnerSVG,
		children,
	)
}

func CommentTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentTwoLineInnerSVG,
		children,
	)
}

func CommentTwoText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentTwoTextInnerSVG,
		children,
	)
}

func CommentTwoTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentTwoTextLineInnerSVG,
		children,
	)
}

func Comments(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentsInnerSVG,
		children,
	)
}

func CommentsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentsLineInnerSVG,
		children,
	)
}

func CommentsTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentsTwoInnerSVG,
		children,
	)
}

func CommentsTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentsTwoLineInnerSVG,
		children,
	)
}

func Community(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		communityInnerSVG,
		children,
	)
}

func CommunityLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		communityLineInnerSVG,
		children,
	)
}

func Compass(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		compassInnerSVG,
		children,
	)
}

func CompassLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		compassLineInnerSVG,
		children,
	)
}

func CompassTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		compassTwoInnerSVG,
		children,
	)
}

func CompassTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		compassTwoLineInnerSVG,
		children,
	)
}

func Contact(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		contactInnerSVG,
		children,
	)
}

func ContactLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		contactLineInnerSVG,
		children,
	)
}

func Cookie(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cookieInnerSVG,
		children,
	)
}

func CookieLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cookieLineInnerSVG,
		children,
	)
}

func Covid(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		covidInnerSVG,
		children,
	)
}

func CovidExclamation(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		covidExclamationInnerSVG,
		children,
	)
}

func CovidExclamationLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		covidExclamationLineInnerSVG,
		children,
	)
}

func CovidLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		covidLineInnerSVG,
		children,
	)
}

func CovidOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		covidOffInnerSVG,
		children,
	)
}

func CovidOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		covidOffLineInnerSVG,
		children,
	)
}

func Cpu(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cpuInnerSVG,
		children,
	)
}

func CpuLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cpuLineInnerSVG,
		children,
	)
}

func CreditCard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardInnerSVG,
		children,
	)
}

func CreditCardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardLineInnerSVG,
		children,
	)
}

func Creditcard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditcardInnerSVG,
		children,
	)
}

func CreditcardHand(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditcardHandInnerSVG,
		children,
	)
}

func CreditcardHandLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditcardHandLineInnerSVG,
		children,
	)
}

func CreditcardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditcardLineInnerSVG,
		children,
	)
}

func CreditcardPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditcardPlusInnerSVG,
		children,
	)
}

func CreditcardPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditcardPlusLineInnerSVG,
		children,
	)
}

func Crown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		crownInnerSVG,
		children,
	)
}

func CrownLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		crownLineInnerSVG,
		children,
	)
}

func Cube(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cubeInnerSVG,
		children,
	)
}

func CubeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cubeLineInnerSVG,
		children,
	)
}

func Cup(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cupInnerSVG,
		children,
	)
}

func CupLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cupLineInnerSVG,
		children,
	)
}

func CurlyBraces(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		curlyBracesInnerSVG,
		children,
	)
}

func CurlyBracesLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		curlyBracesLineInnerSVG,
		children,
	)
}

func CursorClick(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cursorClickInnerSVG,
		children,
	)
}

func CursorClickLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cursorClickLineInnerSVG,
		children,
	)
}

func Dashboard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dashboardInnerSVG,
		children,
	)
}

func DashboardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dashboardLineInnerSVG,
		children,
	)
}

func Data(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dataInnerSVG,
		children,
	)
}

func DataLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dataLineInnerSVG,
		children,
	)
}

func DataMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dataMinusInnerSVG,
		children,
	)
}

func DataMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dataMinusLineInnerSVG,
		children,
	)
}

func DataPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dataPlusInnerSVG,
		children,
	)
}

func DataPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dataPlusLineInnerSVG,
		children,
	)
}

func Database(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		databaseInnerSVG,
		children,
	)
}

func DatabaseLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		databaseLineInnerSVG,
		children,
	)
}

func DeleteBin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deleteBinInnerSVG,
		children,
	)
}

func DeleteBinLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deleteBinLineInnerSVG,
		children,
	)
}

func DesktopComputer(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		desktopComputerInnerSVG,
		children,
	)
}

func DesktopComputerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		desktopComputerLineInnerSVG,
		children,
	)
}

func DeviceMobile(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceMobileInnerSVG,
		children,
	)
}

func DeviceMobileLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceMobileLineInnerSVG,
		children,
	)
}

func DeviceTablet(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceTabletInnerSVG,
		children,
	)
}

func DeviceTabletLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceTabletLineInnerSVG,
		children,
	)
}

func DistributeHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		distributeHorizontalInnerSVG,
		children,
	)
}

func DistributeHorizontalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		distributeHorizontalLineInnerSVG,
		children,
	)
}

func DistributeVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		distributeVerticalInnerSVG,
		children,
	)
}

func DistributeVerticalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		distributeVerticalLineInnerSVG,
		children,
	)
}

func Divide(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		divideInnerSVG,
		children,
	)
}

func DivideLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		divideLineInnerSVG,
		children,
	)
}

func Document(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentInnerSVG,
		children,
	)
}

func DocumentAward(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentAwardInnerSVG,
		children,
	)
}

func DocumentAwardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentAwardLineInnerSVG,
		children,
	)
}

func DocumentLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentLineInnerSVG,
		children,
	)
}

func DollarCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dollarCircleInnerSVG,
		children,
	)
}

func DollarCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dollarCircleLineInnerSVG,
		children,
	)
}

func DoorEnter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		doorEnterInnerSVG,
		children,
	)
}

func DoorEnterLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		doorEnterLineInnerSVG,
		children,
	)
}

func DoorExit(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		doorExitInnerSVG,
		children,
	)
}

func DoorExitLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		doorExitLineInnerSVG,
		children,
	)
}

func DotsCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dotsCircleInnerSVG,
		children,
	)
}

func DotsCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dotsCircleLineInnerSVG,
		children,
	)
}

func DotsHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dotsHorizontalInnerSVG,
		children,
	)
}

func DotsHorizontalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dotsHorizontalLineInnerSVG,
		children,
	)
}

func DotsVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dotsVerticalInnerSVG,
		children,
	)
}

func DotsVerticalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dotsVerticalLineInnerSVG,
		children,
	)
}

func Duplicate(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		duplicateInnerSVG,
		children,
	)
}

func DuplicateLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		duplicateLineInnerSVG,
		children,
	)
}

func EarthSphere(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		earthSphereInnerSVG,
		children,
	)
}

func EarthSphereLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		earthSphereLineInnerSVG,
		children,
	)
}

func EditPenFour(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editPenFourInnerSVG,
		children,
	)
}

func EditPenFourLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editPenFourLineInnerSVG,
		children,
	)
}

func EditPenTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editPenTwoInnerSVG,
		children,
	)
}

func EditPenTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editPenTwoLineInnerSVG,
		children,
	)
}

func Eject(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ejectInnerSVG,
		children,
	)
}

func EjectLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ejectLineInnerSVG,
		children,
	)
}

func EmojiHappy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiHappyInnerSVG,
		children,
	)
}

func EmojiHappyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiHappyLineInnerSVG,
		children,
	)
}

func EmojiSad(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiSadInnerSVG,
		children,
	)
}

func EmojiSadLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiSadLineInnerSVG,
		children,
	)
}

func Eraser(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eraserInnerSVG,
		children,
	)
}

func EraserLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eraserLineInnerSVG,
		children,
	)
}

func EtheriumCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		etheriumCircleInnerSVG,
		children,
	)
}

func EtheriumCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		etheriumCircleLineInnerSVG,
		children,
	)
}

func EuroCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		euroCircleInnerSVG,
		children,
	)
}

func EuroCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		euroCircleLineInnerSVG,
		children,
	)
}

func Exclamation(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		exclamationInnerSVG,
		children,
	)
}

func ExclamationCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		exclamationCircleInnerSVG,
		children,
	)
}

func ExclamationCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		exclamationCircleLineInnerSVG,
		children,
	)
}

func ExclamationLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		exclamationLineInnerSVG,
		children,
	)
}

func ExternalLink(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		externalLinkInnerSVG,
		children,
	)
}

func ExternalLinkLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		externalLinkLineInnerSVG,
		children,
	)
}

func Eye(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeInnerSVG,
		children,
	)
}

func EyeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeLineInnerSVG,
		children,
	)
}

func EyeOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeOffInnerSVG,
		children,
	)
}

func EyeOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeOffLineInnerSVG,
		children,
	)
}

func FastForward(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fastForwardInnerSVG,
		children,
	)
}

func FastForwardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fastForwardLineInnerSVG,
		children,
	)
}

func FerrisWheel(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ferrisWheelInnerSVG,
		children,
	)
}

func FerrisWheelLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ferrisWheelLineInnerSVG,
		children,
	)
}

func File(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileInnerSVG,
		children,
	)
}

func FileAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileAddInnerSVG,
		children,
	)
}

func FileAddLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileAddLineInnerSVG,
		children,
	)
}

func FileDownload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileDownloadInnerSVG,
		children,
	)
}

func FileDownloadLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileDownloadLineInnerSVG,
		children,
	)
}

func FileDuplicate(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileDuplicateInnerSVG,
		children,
	)
}

func FileDuplicateLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileDuplicateLineInnerSVG,
		children,
	)
}

func FileLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileLineInnerSVG,
		children,
	)
}

func FileMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileMinusInnerSVG,
		children,
	)
}

func FileMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileMinusLineInnerSVG,
		children,
	)
}

func FilePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filePlusInnerSVG,
		children,
	)
}

func FilePlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filePlusLineInnerSVG,
		children,
	)
}

func FileRemove(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileRemoveInnerSVG,
		children,
	)
}

func FileRemoveLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileRemoveLineInnerSVG,
		children,
	)
}

func FileReport(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileReportInnerSVG,
		children,
	)
}

func FileReportLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileReportLineInnerSVG,
		children,
	)
}

func FileSearch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileSearchInnerSVG,
		children,
	)
}

func FileSearchLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileSearchLineInnerSVG,
		children,
	)
}

func FileText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileTextInnerSVG,
		children,
	)
}

func FileTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileTextLineInnerSVG,
		children,
	)
}

func Film(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filmInnerSVG,
		children,
	)
}

func FilmLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filmLineInnerSVG,
		children,
	)
}

func Filter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filterInnerSVG,
		children,
	)
}

func FilterLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filterLineInnerSVG,
		children,
	)
}

func Fire(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fireInnerSVG,
		children,
	)
}

func FireLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fireLineInnerSVG,
		children,
	)
}

func Fish(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fishInnerSVG,
		children,
	)
}

func FishLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fishLineInnerSVG,
		children,
	)
}

func Flag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flagInnerSVG,
		children,
	)
}

func FlagLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flagLineInnerSVG,
		children,
	)
}

func Flask(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flaskInnerSVG,
		children,
	)
}

func FlaskLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flaskLineInnerSVG,
		children,
	)
}

func FlowerTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flowerTwoInnerSVG,
		children,
	)
}

func FlowerTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flowerTwoLineInnerSVG,
		children,
	)
}

func Folder(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderInnerSVG,
		children,
	)
}

func FolderAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderAddInnerSVG,
		children,
	)
}

func FolderAddLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderAddLineInnerSVG,
		children,
	)
}

func FolderCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderCheckInnerSVG,
		children,
	)
}

func FolderCheckLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderCheckLineInnerSVG,
		children,
	)
}

func FolderDownload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderDownloadInnerSVG,
		children,
	)
}

func FolderDownloadLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderDownloadLineInnerSVG,
		children,
	)
}

func FolderLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderLineInnerSVG,
		children,
	)
}

func FolderMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderMinusInnerSVG,
		children,
	)
}

func FolderMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderMinusLineInnerSVG,
		children,
	)
}

func FolderOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderOpenInnerSVG,
		children,
	)
}

func FolderOpenLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderOpenLineInnerSVG,
		children,
	)
}

func FolderPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderPlusInnerSVG,
		children,
	)
}

func FolderPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderPlusLineInnerSVG,
		children,
	)
}

func FolderRemove(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderRemoveInnerSVG,
		children,
	)
}

func FolderRemoveLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderRemoveLineInnerSVG,
		children,
	)
}

func FontSize(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fontSizeInnerSVG,
		children,
	)
}

func FontSizeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fontSizeLineInnerSVG,
		children,
	)
}

func Forward(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		forwardInnerSVG,
		children,
	)
}

func ForwardCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		forwardCircleInnerSVG,
		children,
	)
}

func ForwardCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		forwardCircleLineInnerSVG,
		children,
	)
}

func ForwardEndCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		forwardEndCircleInnerSVG,
		children,
	)
}

func ForwardEndCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		forwardEndCircleLineInnerSVG,
		children,
	)
}

func ForwardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		forwardLineInnerSVG,
		children,
	)
}

func Ghost(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ghostInnerSVG,
		children,
	)
}

func GhostLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ghostLineInnerSVG,
		children,
	)
}

func GitBranch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitBranchInnerSVG,
		children,
	)
}

func GitBranchLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitBranchLineInnerSVG,
		children,
	)
}

func GitCommit(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitCommitInnerSVG,
		children,
	)
}

func GitCommitLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitCommitLineInnerSVG,
		children,
	)
}

func GitCompare(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitCompareInnerSVG,
		children,
	)
}

func GitCompareLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitCompareLineInnerSVG,
		children,
	)
}

func GitFork(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitForkInnerSVG,
		children,
	)
}

func GitForkLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitForkLineInnerSVG,
		children,
	)
}

func GitMerge(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitMergeInnerSVG,
		children,
	)
}

func GitMergeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitMergeLineInnerSVG,
		children,
	)
}

func GitPull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitPullInnerSVG,
		children,
	)
}

func GitPullLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitPullLineInnerSVG,
		children,
	)
}

func GlasWater(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		glasWaterInnerSVG,
		children,
	)
}

func GlasWaterLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		glasWaterLineInnerSVG,
		children,
	)
}

func Globe(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeInnerSVG,
		children,
	)
}

func GlobeEarth(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeEarthInnerSVG,
		children,
	)
}

func GlobeEarthLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeEarthLineInnerSVG,
		children,
	)
}

func GlobeEarthTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeEarthTwoInnerSVG,
		children,
	)
}

func GlobeEarthTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeEarthTwoLineInnerSVG,
		children,
	)
}

func GlobeGrid(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeGridInnerSVG,
		children,
	)
}

func GlobeGridLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeGridLineInnerSVG,
		children,
	)
}

func GlobeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeLineInnerSVG,
		children,
	)
}

func Hand(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		handInnerSVG,
		children,
	)
}

func HandLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		handLineInnerSVG,
		children,
	)
}

func HandPointer(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		handPointerInnerSVG,
		children,
	)
}

func HandPointerEvent(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		handPointerEventInnerSVG,
		children,
	)
}

func HandPointerEventLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		handPointerEventLineInnerSVG,
		children,
	)
}

func HandPointerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		handPointerLineInnerSVG,
		children,
	)
}

func HandPointerTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		handPointerTwoInnerSVG,
		children,
	)
}

func HandPointerTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		handPointerTwoLineInnerSVG,
		children,
	)
}

func HardDrive(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hardDriveInnerSVG,
		children,
	)
}

func HardDriveLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hardDriveLineInnerSVG,
		children,
	)
}

func Hashtag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hashtagInnerSVG,
		children,
	)
}

func HashtagLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hashtagLineInnerSVG,
		children,
	)
}

func Headset(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		headsetInnerSVG,
		children,
	)
}

func HeadsetLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		headsetLineInnerSVG,
		children,
	)
}

func Heart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartInnerSVG,
		children,
	)
}

func HeartLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartLineInnerSVG,
		children,
	)
}

func Home(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeInnerSVG,
		children,
	)
}

func HomeAnalytics(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeAnalyticsInnerSVG,
		children,
	)
}

func HomeAnalyticsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeAnalyticsLineInnerSVG,
		children,
	)
}

func HomeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeLineInnerSVG,
		children,
	)
}

func HomeSimple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeSimpleInnerSVG,
		children,
	)
}

func HomeSimpleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeSimpleLineInnerSVG,
		children,
	)
}

func Image(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageInnerSVG,
		children,
	)
}

func ImageCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageCircleInnerSVG,
		children,
	)
}

func ImageCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageCircleLineInnerSVG,
		children,
	)
}

func ImageCircleOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageCircleOffInnerSVG,
		children,
	)
}

func ImageCircleOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageCircleOffLineInnerSVG,
		children,
	)
}

func ImageCirclePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageCirclePlusInnerSVG,
		children,
	)
}

func ImageCirclePlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageCirclePlusLineInnerSVG,
		children,
	)
}

func ImageCircleStory(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageCircleStoryInnerSVG,
		children,
	)
}

func ImageCircleStoryLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageCircleStoryLineInnerSVG,
		children,
	)
}

func ImageFrame(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageFrameInnerSVG,
		children,
	)
}

func ImageFrameLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageFrameLineInnerSVG,
		children,
	)
}

func ImageInPicture(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageInPictureInnerSVG,
		children,
	)
}

func ImageInPictureLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageInPictureLineInnerSVG,
		children,
	)
}

func ImageLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageLineInnerSVG,
		children,
	)
}

func ImageMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageMultipleInnerSVG,
		children,
	)
}

func ImageMultipleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageMultipleLineInnerSVG,
		children,
	)
}

func ImageOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageOffInnerSVG,
		children,
	)
}

func ImageOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageOffLineInnerSVG,
		children,
	)
}

func ImagePhotography(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imagePhotographyInnerSVG,
		children,
	)
}

func ImagePhotographyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imagePhotographyLineInnerSVG,
		children,
	)
}

func ImagePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imagePlusInnerSVG,
		children,
	)
}

func ImagePlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imagePlusLineInnerSVG,
		children,
	)
}

func Inbox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxInnerSVG,
		children,
	)
}

func InboxIn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxInInnerSVG,
		children,
	)
}

func InboxInLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxInLineInnerSVG,
		children,
	)
}

func InboxLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxLineInnerSVG,
		children,
	)
}

func Incognito(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		incognitoInnerSVG,
		children,
	)
}

func IncognitoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		incognitoLineInnerSVG,
		children,
	)
}

func InfoCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		infoCircleInnerSVG,
		children,
	)
}

func InfoCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		infoCircleLineInnerSVG,
		children,
	)
}

func InformationCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		informationCircleInnerSVG,
		children,
	)
}

func InformationCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		informationCircleLineInnerSVG,
		children,
	)
}

func IphoneOldApps(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		iphoneOldAppsInnerSVG,
		children,
	)
}

func IphoneOldAppsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		iphoneOldAppsLineInnerSVG,
		children,
	)
}

func IphoneXApps(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		iphoneXAppsInnerSVG,
		children,
	)
}

func IphoneXAppsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		iphoneXAppsLineInnerSVG,
		children,
	)
}

func Italic(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		italicInnerSVG,
		children,
	)
}

func ItalicLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		italicLineInnerSVG,
		children,
	)
}

func Key(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		keyInnerSVG,
		children,
	)
}

func KeyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		keyLineInnerSVG,
		children,
	)
}

func Keyboard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		keyboardInnerSVG,
		children,
	)
}

func KeyboardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		keyboardLineInnerSVG,
		children,
	)
}

func Laptop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		laptopInnerSVG,
		children,
	)
}

func LaptopLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		laptopLineInnerSVG,
		children,
	)
}

func LeafThreeAngled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		leafThreeAngledInnerSVG,
		children,
	)
}

func LeafThreeAngledLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		leafThreeAngledLineInnerSVG,
		children,
	)
}

func Library(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		libraryInnerSVG,
		children,
	)
}

func LibraryLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		libraryLineInnerSVG,
		children,
	)
}

func LidquidDropWavesTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lidquidDropWavesTwoInnerSVG,
		children,
	)
}

func LidquidDropWavesTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lidquidDropWavesTwoLineInnerSVG,
		children,
	)
}

func Lifebuoy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lifebuoyInnerSVG,
		children,
	)
}

func LifebuoyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lifebuoyLineInnerSVG,
		children,
	)
}

func LightBulb(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightBulbInnerSVG,
		children,
	)
}

func LightBulbLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightBulbLineInnerSVG,
		children,
	)
}

func LightbulbShine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbShineInnerSVG,
		children,
	)
}

func LightbulbShineLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbShineLineInnerSVG,
		children,
	)
}

func LightningBolt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightningBoltInnerSVG,
		children,
	)
}

func LightningBoltLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightningBoltLineInnerSVG,
		children,
	)
}

func LineHeight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lineHeightInnerSVG,
		children,
	)
}

func LineHeightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lineHeightLineInnerSVG,
		children,
	)
}

func Link(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkInnerSVG,
		children,
	)
}

func LinkCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkCircleInnerSVG,
		children,
	)
}

func LinkCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkCircleLineInnerSVG,
		children,
	)
}

func LinkLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkLineInnerSVG,
		children,
	)
}

func LiraCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		liraCircleInnerSVG,
		children,
	)
}

func LiraCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		liraCircleLineInnerSVG,
		children,
	)
}

func ListBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listBoxInnerSVG,
		children,
	)
}

func ListBoxLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listBoxLineInnerSVG,
		children,
	)
}

func LocationMarker(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		locationMarkerInnerSVG,
		children,
	)
}

func LocationMarkerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		locationMarkerLineInnerSVG,
		children,
	)
}

func Lock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockInnerSVG,
		children,
	)
}

func LockClosed(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockClosedInnerSVG,
		children,
	)
}

func LockClosedLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockClosedLineInnerSVG,
		children,
	)
}

func LockLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockLineInnerSVG,
		children,
	)
}

func LockOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockOffInnerSVG,
		children,
	)
}

func LockOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockOffLineInnerSVG,
		children,
	)
}

func LockOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockOpenInnerSVG,
		children,
	)
}

func LockOpenLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockOpenLineInnerSVG,
		children,
	)
}

func Login(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loginInnerSVG,
		children,
	)
}

func LoginHalfCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loginHalfCircleInnerSVG,
		children,
	)
}

func LoginHalfCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loginHalfCircleLineInnerSVG,
		children,
	)
}

func LoginLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loginLineInnerSVG,
		children,
	)
}

func Logout(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logoutInnerSVG,
		children,
	)
}

func LogoutHalfCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logoutHalfCircleInnerSVG,
		children,
	)
}

func LogoutHalfCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logoutHalfCircleLineInnerSVG,
		children,
	)
}

func LogoutLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logoutLineInnerSVG,
		children,
	)
}

func Mail(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailInnerSVG,
		children,
	)
}

func MailLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailLineInnerSVG,
		children,
	)
}

func MailOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailOpenInnerSVG,
		children,
	)
}

func MailOpenLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailOpenLineInnerSVG,
		children,
	)
}

func Map(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapInnerSVG,
		children,
	)
}

func MapLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapLineInnerSVG,
		children,
	)
}

func MapMarker(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapMarkerInnerSVG,
		children,
	)
}

func MapMarkerArea(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapMarkerAreaInnerSVG,
		children,
	)
}

func MapMarkerAreaLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapMarkerAreaLineInnerSVG,
		children,
	)
}

func MapMarkerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapMarkerLineInnerSVG,
		children,
	)
}

func MapMarkerPath(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapMarkerPathInnerSVG,
		children,
	)
}

func MapMarkerPathLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapMarkerPathLineInnerSVG,
		children,
	)
}

func MapMarkerPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapMarkerPlusInnerSVG,
		children,
	)
}

func MapMarkerPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapMarkerPlusLineInnerSVG,
		children,
	)
}

func MapSimple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapSimpleInnerSVG,
		children,
	)
}

func MapSimpleDestination(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapSimpleDestinationInnerSVG,
		children,
	)
}

func MapSimpleDestinationLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapSimpleDestinationLineInnerSVG,
		children,
	)
}

func MapSimpleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapSimpleLineInnerSVG,
		children,
	)
}

func MapSimpleMarker(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapSimpleMarkerInnerSVG,
		children,
	)
}

func MapSimpleMarkerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapSimpleMarkerLineInnerSVG,
		children,
	)
}

func MapSimpleOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapSimpleOffInnerSVG,
		children,
	)
}

func MapSimpleOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapSimpleOffLineInnerSVG,
		children,
	)
}

func Maximize(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		maximizeInnerSVG,
		children,
	)
}

func MaximizeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		maximizeLineInnerSVG,
		children,
	)
}

func Megaphone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		megaphoneInnerSVG,
		children,
	)
}

func MegaphoneLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		megaphoneLineInnerSVG,
		children,
	)
}

func Menu(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuInnerSVG,
		children,
	)
}

func MenuAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuAltInnerSVG,
		children,
	)
}

func MenuAltLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuAltLineInnerSVG,
		children,
	)
}

func MenuExpandLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuExpandLeftInnerSVG,
		children,
	)
}

func MenuExpandLeftLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuExpandLeftLineInnerSVG,
		children,
	)
}

func MenuExpandRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuExpandRightInnerSVG,
		children,
	)
}

func MenuExpandRightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuExpandRightLineInnerSVG,
		children,
	)
}

func MenuLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuLineInnerSVG,
		children,
	)
}

func Message(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageInnerSVG,
		children,
	)
}

func MessageLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageLineInnerSVG,
		children,
	)
}

func Messages(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messagesInnerSVG,
		children,
	)
}

func MessagesLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messagesLineInnerSVG,
		children,
	)
}

func Microphone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		microphoneInnerSVG,
		children,
	)
}

func MicrophoneLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		microphoneLineInnerSVG,
		children,
	)
}

func Minimize(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minimizeInnerSVG,
		children,
	)
}

func MinimizeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minimizeLineInnerSVG,
		children,
	)
}

func Minus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusInnerSVG,
		children,
	)
}

func MinusCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusCircleInnerSVG,
		children,
	)
}

func MinusCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusCircleLineInnerSVG,
		children,
	)
}

func MinusFiveCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusFiveCircleInnerSVG,
		children,
	)
}

func MinusFiveCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusFiveCircleLineInnerSVG,
		children,
	)
}

func MinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusLineInnerSVG,
		children,
	)
}

func MinusTenCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusTenCircleInnerSVG,
		children,
	)
}

func MinusTenCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusTenCircleLineInnerSVG,
		children,
	)
}

func Money(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyInnerSVG,
		children,
	)
}

func MoneyHand(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyHandInnerSVG,
		children,
	)
}

func MoneyHandLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyHandLineInnerSVG,
		children,
	)
}

func MoneyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyLineInnerSVG,
		children,
	)
}

func MoneyMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyMinusInnerSVG,
		children,
	)
}

func MoneyMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyMinusLineInnerSVG,
		children,
	)
}

func MoneyPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyPlusInnerSVG,
		children,
	)
}

func MoneyPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyPlusLineInnerSVG,
		children,
	)
}

func Monitor(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		monitorInnerSVG,
		children,
	)
}

func MonitorLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		monitorLineInnerSVG,
		children,
	)
}

func Moon(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonInnerSVG,
		children,
	)
}

func MoonLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonLineInnerSVG,
		children,
	)
}

func MoreMenu(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreMenuInnerSVG,
		children,
	)
}

func MoreMenuLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreMenuLineInnerSVG,
		children,
	)
}

func MoreMenuVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreMenuVerticalInnerSVG,
		children,
	)
}

func MoreMenuVerticalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreMenuVerticalLineInnerSVG,
		children,
	)
}

func Mouse(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mouseInnerSVG,
		children,
	)
}

func MouseLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mouseLineInnerSVG,
		children,
	)
}

func Multiply(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		multiplyInnerSVG,
		children,
	)
}

func MultiplyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		multiplyLineInnerSVG,
		children,
	)
}

func Music(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		musicInnerSVG,
		children,
	)
}

func MusicLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		musicLineInnerSVG,
		children,
	)
}

func MusicNote(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		musicNoteInnerSVG,
		children,
	)
}

func MusicNoteLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		musicNoteLineInnerSVG,
		children,
	)
}

func Newspaper(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		newspaperInnerSVG,
		children,
	)
}

func NewspaperLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		newspaperLineInnerSVG,
		children,
	)
}

func NextCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		nextCircleInnerSVG,
		children,
	)
}

func NextCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		nextCircleLineInnerSVG,
		children,
	)
}

func NoteText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteTextInnerSVG,
		children,
	)
}

func NoteTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteTextLineInnerSVG,
		children,
	)
}

func NoteTextMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteTextMinusInnerSVG,
		children,
	)
}

func NoteTextMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteTextMinusLineInnerSVG,
		children,
	)
}

func NoteTextPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteTextPlusInnerSVG,
		children,
	)
}

func NoteTextPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteTextPlusLineInnerSVG,
		children,
	)
}

func Noteblock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteblockInnerSVG,
		children,
	)
}

func NoteblockLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteblockLineInnerSVG,
		children,
	)
}

func NoteblockText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteblockTextInnerSVG,
		children,
	)
}

func NoteblockTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteblockTextLineInnerSVG,
		children,
	)
}

func Open(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		openInnerSVG,
		children,
	)
}

func OpenLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		openLineInnerSVG,
		children,
	)
}

func PaperAirplane(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperAirplaneInnerSVG,
		children,
	)
}

func PaperAirplaneLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperAirplaneLineInnerSVG,
		children,
	)
}

func PaperClip(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperClipInnerSVG,
		children,
	)
}

func PaperClipLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperClipLineInnerSVG,
		children,
	)
}

func PaperFold(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperFoldInnerSVG,
		children,
	)
}

func PaperFoldLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperFoldLineInnerSVG,
		children,
	)
}

func PaperFoldText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperFoldTextInnerSVG,
		children,
	)
}

func PaperFoldTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperFoldTextLineInnerSVG,
		children,
	)
}

func PaperRollTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperRollTwoInnerSVG,
		children,
	)
}

func PaperRollTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperRollTwoLineInnerSVG,
		children,
	)
}

func Paragraph(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paragraphInnerSVG,
		children,
	)
}

func ParagraphLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paragraphLineInnerSVG,
		children,
	)
}

func Pause(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pauseInnerSVG,
		children,
	)
}

func PauseCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pauseCircleInnerSVG,
		children,
	)
}

func PauseCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pauseCircleLineInnerSVG,
		children,
	)
}

func PauseLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pauseLineInnerSVG,
		children,
	)
}

func Pencil(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pencilInnerSVG,
		children,
	)
}

func PencilAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pencilAltInnerSVG,
		children,
	)
}

func PencilAltLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pencilAltLineInnerSVG,
		children,
	)
}

func PencilLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pencilLineInnerSVG,
		children,
	)
}

func Percent(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		percentInnerSVG,
		children,
	)
}

func PercentLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		percentLineInnerSVG,
		children,
	)
}

func Phone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneInnerSVG,
		children,
	)
}

func PhoneDial(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneDialInnerSVG,
		children,
	)
}

func PhoneDialLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneDialLineInnerSVG,
		children,
	)
}

func PhoneHangup(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneHangupInnerSVG,
		children,
	)
}

func PhoneHangupLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneHangupLineInnerSVG,
		children,
	)
}

func PhoneIncoming(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneIncomingInnerSVG,
		children,
	)
}

func PhoneIncomingLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneIncomingLineInnerSVG,
		children,
	)
}

func PhoneLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneLineInnerSVG,
		children,
	)
}

func PhoneMissedCall(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneMissedCallInnerSVG,
		children,
	)
}

func PhoneMissedCallLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneMissedCallLineInnerSVG,
		children,
	)
}

func PhoneOutgoing(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneOutgoingInnerSVG,
		children,
	)
}

func PhoneOutgoingLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneOutgoingLineInnerSVG,
		children,
	)
}

func PhoneRetro(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneRetroInnerSVG,
		children,
	)
}

func PhoneRetroLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneRetroLineInnerSVG,
		children,
	)
}

func PhoneRing(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneRingInnerSVG,
		children,
	)
}

func PhoneRingLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneRingLineInnerSVG,
		children,
	)
}

func Pill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pillInnerSVG,
		children,
	)
}

func PillLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pillLineInnerSVG,
		children,
	)
}

func Pin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pinInnerSVG,
		children,
	)
}

func PinLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pinLineInnerSVG,
		children,
	)
}

func Pinwheel(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pinwheelInnerSVG,
		children,
	)
}

func PinwheelLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pinwheelLineInnerSVG,
		children,
	)
}

func Planet(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		planetInnerSVG,
		children,
	)
}

func PlanetLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		planetLineInnerSVG,
		children,
	)
}

func PlanetRingTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		planetRingTwoInnerSVG,
		children,
	)
}

func PlanetRingTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		planetRingTwoLineInnerSVG,
		children,
	)
}

func PlanetRocket(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		planetRocketInnerSVG,
		children,
	)
}

func PlanetRocketLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		planetRocketLineInnerSVG,
		children,
	)
}

func PlayCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playCircleInnerSVG,
		children,
	)
}

func PlayCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playCircleLineInnerSVG,
		children,
	)
}

func Playlist(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playlistInnerSVG,
		children,
	)
}

func PlaylistLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playlistLineInnerSVG,
		children,
	)
}

func Plus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusInnerSVG,
		children,
	)
}

func PlusCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusCircleInnerSVG,
		children,
	)
}

func PlusCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusCircleLineInnerSVG,
		children,
	)
}

func PlusFiveCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusFiveCircleInnerSVG,
		children,
	)
}

func PlusFiveCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusFiveCircleLineInnerSVG,
		children,
	)
}

func PlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusLineInnerSVG,
		children,
	)
}

func PlusMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusMinusInnerSVG,
		children,
	)
}

func PlusMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusMinusLineInnerSVG,
		children,
	)
}

func PlusMinusTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusMinusTwoInnerSVG,
		children,
	)
}

func PlusMinusTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusMinusTwoLineInnerSVG,
		children,
	)
}

func PlusTenCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusTenCircleInnerSVG,
		children,
	)
}

func PlusTenCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusTenCircleLineInnerSVG,
		children,
	)
}

func PoundCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		poundCircleInnerSVG,
		children,
	)
}

func PoundCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		poundCircleLineInnerSVG,
		children,
	)
}

func Presentation(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationInnerSVG,
		children,
	)
}

func PresentationChart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationChartInnerSVG,
		children,
	)
}

func PresentationChartLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationChartLineInnerSVG,
		children,
	)
}

func PresentationLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationLineInnerSVG,
		children,
	)
}

func PresentationLineLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationLineLineInnerSVG,
		children,
	)
}

func PresentationPlay(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationPlayInnerSVG,
		children,
	)
}

func PresentationPlayLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationPlayLineInnerSVG,
		children,
	)
}

func PresentationReport(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationReportInnerSVG,
		children,
	)
}

func PresentationReportLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		presentationReportLineInnerSVG,
		children,
	)
}

func Printer(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		printerInnerSVG,
		children,
	)
}

func PrinterLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		printerLineInnerSVG,
		children,
	)
}

func Pulse(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pulseInnerSVG,
		children,
	)
}

func PulseLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pulseLineInnerSVG,
		children,
	)
}

func Puzzle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		puzzleInnerSVG,
		children,
	)
}

func PuzzleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		puzzleLineInnerSVG,
		children,
	)
}

func QrCode(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		qrCodeInnerSVG,
		children,
	)
}

func QrCodeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		qrCodeLineInnerSVG,
		children,
	)
}

func Qrcode(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		qrcodeInnerSVG,
		children,
	)
}

func QrcodeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		qrcodeLineInnerSVG,
		children,
	)
}

func QuestionCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionCircleInnerSVG,
		children,
	)
}

func QuestionCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionCircleLineInnerSVG,
		children,
	)
}

func QuestionMarkCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionMarkCircleInnerSVG,
		children,
	)
}

func QuestionMarkCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionMarkCircleLineInnerSVG,
		children,
	)
}

func RadioList(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioListInnerSVG,
		children,
	)
}

func RadioListLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioListLineInnerSVG,
		children,
	)
}

func Receipt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		receiptInnerSVG,
		children,
	)
}

func ReceiptLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		receiptLineInnerSVG,
		children,
	)
}

func ReceiptRefund(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		receiptRefundInnerSVG,
		children,
	)
}

func ReceiptRefundLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		receiptRefundLineInnerSVG,
		children,
	)
}

func ReceiptText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		receiptTextInnerSVG,
		children,
	)
}

func ReceiptTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		receiptTextLineInnerSVG,
		children,
	)
}

func Redo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		redoInnerSVG,
		children,
	)
}

func RedoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		redoLineInnerSVG,
		children,
	)
}

func Refresh(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		refreshInnerSVG,
		children,
	)
}

func RefreshLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		refreshLineInnerSVG,
		children,
	)
}

func Reload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		reloadInnerSVG,
		children,
	)
}

func ReloadCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		reloadCircleInnerSVG,
		children,
	)
}

func ReloadCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		reloadCircleLineInnerSVG,
		children,
	)
}

func ReloadLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		reloadLineInnerSVG,
		children,
	)
}

func RemoveColumn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeColumnInnerSVG,
		children,
	)
}

func RemoveColumnLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeColumnLineInnerSVG,
		children,
	)
}

func RemoveFormat(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeFormatInnerSVG,
		children,
	)
}

func RemoveFormatLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeFormatLineInnerSVG,
		children,
	)
}

func RemoveRow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeRowInnerSVG,
		children,
	)
}

func RemoveRowLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeRowLineInnerSVG,
		children,
	)
}

func RepeatCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		repeatCircleInnerSVG,
		children,
	)
}

func RepeatCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		repeatCircleLineInnerSVG,
		children,
	)
}

func Reply(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		replyInnerSVG,
		children,
	)
}

func ReplyAll(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		replyAllInnerSVG,
		children,
	)
}

func ReplyAllLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		replyAllLineInnerSVG,
		children,
	)
}

func ReplyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		replyLineInnerSVG,
		children,
	)
}

func Restricted(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		restrictedInnerSVG,
		children,
	)
}

func RestrictedLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		restrictedLineInnerSVG,
		children,
	)
}

func Rewind(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rewindInnerSVG,
		children,
	)
}

func RewindLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rewindLineInnerSVG,
		children,
	)
}

func Robot(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		robotInnerSVG,
		children,
	)
}

func RobotLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		robotLineInnerSVG,
		children,
	)
}

func RocketThreeStart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rocketThreeStartInnerSVG,
		children,
	)
}

func RocketThreeStartLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rocketThreeStartLineInnerSVG,
		children,
	)
}

func Rss(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rssInnerSVG,
		children,
	)
}

func RssLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rssLineInnerSVG,
		children,
	)
}

func RubelCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rubelCircleInnerSVG,
		children,
	)
}

func RubelCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rubelCircleLineInnerSVG,
		children,
	)
}

func RulerTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rulerTwoInnerSVG,
		children,
	)
}

func RulerTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rulerTwoLineInnerSVG,
		children,
	)
}

func RupeeCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rupeeCircleInnerSVG,
		children,
	)
}

func RupeeCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rupeeCircleLineInnerSVG,
		children,
	)
}

func Save(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveInnerSVG,
		children,
	)
}

func SaveAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveAltInnerSVG,
		children,
	)
}

func SaveAltLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveAltLineInnerSVG,
		children,
	)
}

func SaveAs(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveAsInnerSVG,
		children,
	)
}

func SaveAsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveAsLineInnerSVG,
		children,
	)
}

func SaveLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveLineInnerSVG,
		children,
	)
}

func ScaleLight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scaleLightInnerSVG,
		children,
	)
}

func ScaleLightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scaleLightLineInnerSVG,
		children,
	)
}

func ScanFingerprint(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scanFingerprintInnerSVG,
		children,
	)
}

func ScanFingerprintLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scanFingerprintLineInnerSVG,
		children,
	)
}

func ScanUser(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scanUserInnerSVG,
		children,
	)
}

func ScanUserLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scanUserLineInnerSVG,
		children,
	)
}

func Scanner(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scannerInnerSVG,
		children,
	)
}

func ScannerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scannerLineInnerSVG,
		children,
	)
}

func Scissors(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scissorsInnerSVG,
		children,
	)
}

func ScissorsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scissorsLineInnerSVG,
		children,
	)
}

func Scooter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scooterInnerSVG,
		children,
	)
}

func ScooterLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scooterLineInnerSVG,
		children,
	)
}

func ScriptPrescription(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scriptPrescriptionInnerSVG,
		children,
	)
}

func ScriptPrescriptionLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scriptPrescriptionLineInnerSVG,
		children,
	)
}

func Scroll(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scrollInnerSVG,
		children,
	)
}

func ScrollLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scrollLineInnerSVG,
		children,
	)
}

func ScrollText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scrollTextInnerSVG,
		children,
	)
}

func ScrollTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scrollTextLineInnerSVG,
		children,
	)
}

func Search(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchInnerSVG,
		children,
	)
}

func SearchCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchCircleInnerSVG,
		children,
	)
}

func SearchCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchCircleLineInnerSVG,
		children,
	)
}

func SearchLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchLineInnerSVG,
		children,
	)
}

func SearchMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchMinusInnerSVG,
		children,
	)
}

func SearchMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchMinusLineInnerSVG,
		children,
	)
}

func SearchPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchPlusInnerSVG,
		children,
	)
}

func SearchPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchPlusLineInnerSVG,
		children,
	)
}

func SearchText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchTextInnerSVG,
		children,
	)
}

func SearchTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchTextLineInnerSVG,
		children,
	)
}

func Selector(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		selectorInnerSVG,
		children,
	)
}

func SelectorLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		selectorLineInnerSVG,
		children,
	)
}

func Send(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sendInnerSVG,
		children,
	)
}

func SendLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sendLineInnerSVG,
		children,
	)
}

func Server(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		serverInnerSVG,
		children,
	)
}

func ServerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		serverLineInnerSVG,
		children,
	)
}

func SettingsCog(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsCogInnerSVG,
		children,
	)
}

func SettingsCogCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsCogCheckInnerSVG,
		children,
	)
}

func SettingsCogCheckLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsCogCheckLineInnerSVG,
		children,
	)
}

func SettingsCogLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsCogLineInnerSVG,
		children,
	)
}

func SettingsCogPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsCogPlusInnerSVG,
		children,
	)
}

func SettingsCogPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsCogPlusLineInnerSVG,
		children,
	)
}

func Share(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shareInnerSVG,
		children,
	)
}

func ShareCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shareCircleInnerSVG,
		children,
	)
}

func ShareCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shareCircleLineInnerSVG,
		children,
	)
}

func ShareLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shareLineInnerSVG,
		children,
	)
}

func Shield(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldInnerSVG,
		children,
	)
}

func ShieldCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldCheckInnerSVG,
		children,
	)
}

func ShieldCheckLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldCheckLineInnerSVG,
		children,
	)
}

func ShieldExclamation(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldExclamationInnerSVG,
		children,
	)
}

func ShieldExclamationLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldExclamationLineInnerSVG,
		children,
	)
}

func ShieldLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldLineInnerSVG,
		children,
	)
}

func ShieldOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldOffInnerSVG,
		children,
	)
}

func ShieldOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldOffLineInnerSVG,
		children,
	)
}

func ShieldPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldPlusInnerSVG,
		children,
	)
}

func ShieldPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldPlusLineInnerSVG,
		children,
	)
}

func Ship(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shipInnerSVG,
		children,
	)
}

func ShipLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shipLineInnerSVG,
		children,
	)
}

func ShootingStar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shootingStarInnerSVG,
		children,
	)
}

func ShootingStarLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shootingStarLineInnerSVG,
		children,
	)
}

func ShoppingBag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingBagInnerSVG,
		children,
	)
}

func ShoppingBagLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingBagLineInnerSVG,
		children,
	)
}

func ShoppingCart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingCartInnerSVG,
		children,
	)
}

func ShoppingCartLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingCartLineInnerSVG,
		children,
	)
}

func SimCard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		simCardInnerSVG,
		children,
	)
}

func SimCardLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		simCardLineInnerSVG,
		children,
	)
}

func Sitemap(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sitemapInnerSVG,
		children,
	)
}

func SitemapLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sitemapLineInnerSVG,
		children,
	)
}

func Skull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		skullInnerSVG,
		children,
	)
}

func SkullLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		skullLineInnerSVG,
		children,
	)
}

func SmartphoneApps(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		smartphoneAppsInnerSVG,
		children,
	)
}

func SmartphoneAppsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		smartphoneAppsLineInnerSVG,
		children,
	)
}

func SortHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sortHorizontalInnerSVG,
		children,
	)
}

func SortHorizontalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sortHorizontalLineInnerSVG,
		children,
	)
}

func SortVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sortVerticalInnerSVG,
		children,
	)
}

func SortVerticalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sortVerticalLineInnerSVG,
		children,
	)
}

func SoundOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		soundOffInnerSVG,
		children,
	)
}

func SoundOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		soundOffLineInnerSVG,
		children,
	)
}

func SoundUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		soundUpInnerSVG,
		children,
	)
}

func SoundUpLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		soundUpLineInnerSVG,
		children,
	)
}

func Sparkles(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sparklesInnerSVG,
		children,
	)
}

func SparklesLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sparklesLineInnerSVG,
		children,
	)
}

func Speaker(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speakerInnerSVG,
		children,
	)
}

func SpeakerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speakerLineInnerSVG,
		children,
	)
}

func Speakerphone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speakerphoneInnerSVG,
		children,
	)
}

func SpeakerphoneLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speakerphoneLineInnerSVG,
		children,
	)
}

func Star(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		starInnerSVG,
		children,
	)
}

func StarLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		starLineInnerSVG,
		children,
	)
}

func StatusOffline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		statusOfflineInnerSVG,
		children,
	)
}

func StatusOfflineLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		statusOfflineLineInnerSVG,
		children,
	)
}

func StatusOnline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		statusOnlineInnerSVG,
		children,
	)
}

func StatusOnlineLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		statusOnlineLineInnerSVG,
		children,
	)
}

func StopCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		stopCircleInnerSVG,
		children,
	)
}

func StopCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		stopCircleLineInnerSVG,
		children,
	)
}

func StrikeThrough(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		strikeThroughInnerSVG,
		children,
	)
}

func StrikeThroughLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		strikeThroughLineInnerSVG,
		children,
	)
}

func Suitcase(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		suitcaseInnerSVG,
		children,
	)
}

func SuitcaseLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		suitcaseLineInnerSVG,
		children,
	)
}

func SuitcaseThree(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		suitcaseThreeInnerSVG,
		children,
	)
}

func SuitcaseThreeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		suitcaseThreeLineInnerSVG,
		children,
	)
}

func SuitcaseTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		suitcaseTwoInnerSVG,
		children,
	)
}

func SuitcaseTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		suitcaseTwoLineInnerSVG,
		children,
	)
}

func Sun(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunInnerSVG,
		children,
	)
}

func SunLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunLineInnerSVG,
		children,
	)
}

func Support(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		supportInnerSVG,
		children,
	)
}

func SupportLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		supportLineInnerSVG,
		children,
	)
}

func TShirt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tShirtInnerSVG,
		children,
	)
}

func TShirtLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tShirtLineInnerSVG,
		children,
	)
}

func Table(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tableInnerSVG,
		children,
	)
}

func TableHeart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tableHeartInnerSVG,
		children,
	)
}

func TableHeartLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tableHeartLineInnerSVG,
		children,
	)
}

func TableLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tableLineInnerSVG,
		children,
	)
}

func TablePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tablePlusInnerSVG,
		children,
	)
}

func TablePlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tablePlusLineInnerSVG,
		children,
	)
}

func Tag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tagInnerSVG,
		children,
	)
}

func TagLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tagLineInnerSVG,
		children,
	)
}

func TagOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tagOffInnerSVG,
		children,
	)
}

func TagOffLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tagOffLineInnerSVG,
		children,
	)
}

func Telescope(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		telescopeInnerSVG,
		children,
	)
}

func TelescopeLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		telescopeLineInnerSVG,
		children,
	)
}

func Terminal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		terminalInnerSVG,
		children,
	)
}

func TerminalLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		terminalLineInnerSVG,
		children,
	)
}

func TestTubeFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		testTubeFilledInnerSVG,
		children,
	)
}

func TestTubeFilledLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		testTubeFilledLineInnerSVG,
		children,
	)
}

func Text(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textInnerSVG,
		children,
	)
}

func TextAlignCenter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAlignCenterInnerSVG,
		children,
	)
}

func TextAlignCenterLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAlignCenterLineInnerSVG,
		children,
	)
}

func TextAlignJustify(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAlignJustifyInnerSVG,
		children,
	)
}

func TextAlignJustifyLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAlignJustifyLineInnerSVG,
		children,
	)
}

func TextAlignLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAlignLeftInnerSVG,
		children,
	)
}

func TextAlignLeftLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAlignLeftLineInnerSVG,
		children,
	)
}

func TextAlignRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAlignRightInnerSVG,
		children,
	)
}

func TextAlignRightLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAlignRightLineInnerSVG,
		children,
	)
}

func TextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textLineInnerSVG,
		children,
	)
}

func TextWrap(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textWrapInnerSVG,
		children,
	)
}

func TextWrapLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textWrapLineInnerSVG,
		children,
	)
}

func Textbox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textboxInnerSVG,
		children,
	)
}

func TextboxLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textboxLineInnerSVG,
		children,
	)
}

func TextboxMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textboxMinusInnerSVG,
		children,
	)
}

func TextboxMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textboxMinusLineInnerSVG,
		children,
	)
}

func TextboxPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textboxPlusInnerSVG,
		children,
	)
}

func TextboxPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textboxPlusLineInnerSVG,
		children,
	)
}

func ThumbDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thumbDownInnerSVG,
		children,
	)
}

func ThumbDownLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thumbDownLineInnerSVG,
		children,
	)
}

func ThumbUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thumbUpInnerSVG,
		children,
	)
}

func ThumbUpLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thumbUpLineInnerSVG,
		children,
	)
}

func Ticket(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ticketInnerSVG,
		children,
	)
}

func TicketCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ticketCheckInnerSVG,
		children,
	)
}

func TicketCheckLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ticketCheckLineInnerSVG,
		children,
	)
}

func TicketLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ticketLineInnerSVG,
		children,
	)
}

func TicketText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ticketTextInnerSVG,
		children,
	)
}

func TicketTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ticketTextLineInnerSVG,
		children,
	)
}

func Tickets(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ticketsInnerSVG,
		children,
	)
}

func TicketsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ticketsLineInnerSVG,
		children,
	)
}

func Timer(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		timerInnerSVG,
		children,
	)
}

func TimerLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		timerLineInnerSVG,
		children,
	)
}

func Tooltip(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tooltipInnerSVG,
		children,
	)
}

func TooltipLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tooltipLineInnerSVG,
		children,
	)
}

func TooltipText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tooltipTextInnerSVG,
		children,
	)
}

func TooltipTextLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tooltipTextLineInnerSVG,
		children,
	)
}

func Tooltips(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tooltipsInnerSVG,
		children,
	)
}

func TooltipsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tooltipsLineInnerSVG,
		children,
	)
}

func TooltipsTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tooltipsTwoInnerSVG,
		children,
	)
}

func TooltipsTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tooltipsTwoLineInnerSVG,
		children,
	)
}

func Translate(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		translateInnerSVG,
		children,
	)
}

func TranslateLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		translateLineInnerSVG,
		children,
	)
}

func Trash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trashInnerSVG,
		children,
	)
}

func TrashLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trashLineInnerSVG,
		children,
	)
}

func TrendingDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingDownInnerSVG,
		children,
	)
}

func TrendingDownLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingDownLineInnerSVG,
		children,
	)
}

func TrendingUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingUpInnerSVG,
		children,
	)
}

func TrendingUpLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingUpLineInnerSVG,
		children,
	)
}

func Truck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		truckInnerSVG,
		children,
	)
}

func TruckLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		truckLineInnerSVG,
		children,
	)
}

func TvOld(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tvOldInnerSVG,
		children,
	)
}

func TvOldLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tvOldLineInnerSVG,
		children,
	)
}

func Umbrella(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		umbrellaInnerSVG,
		children,
	)
}

func UmbrellaLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		umbrellaLineInnerSVG,
		children,
	)
}

func Underline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		underlineInnerSVG,
		children,
	)
}

func UnderlineLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		underlineLineInnerSVG,
		children,
	)
}

func UnderlineTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		underlineTwoInnerSVG,
		children,
	)
}

func UnderlineTwoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		underlineTwoLineInnerSVG,
		children,
	)
}

func Undo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		undoInnerSVG,
		children,
	)
}

func UndoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		undoLineInnerSVG,
		children,
	)
}

func UnlockOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		unlockOpenInnerSVG,
		children,
	)
}

func UnlockOpenLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		unlockOpenLineInnerSVG,
		children,
	)
}

func Usb(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		usbInnerSVG,
		children,
	)
}

func UsbLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		usbLineInnerSVG,
		children,
	)
}

func User(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userInnerSVG,
		children,
	)
}

func UserAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userAddInnerSVG,
		children,
	)
}

func UserAddLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userAddLineInnerSVG,
		children,
	)
}

func UserBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userBoxInnerSVG,
		children,
	)
}

func UserBoxLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userBoxLineInnerSVG,
		children,
	)
}

func UserCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userCircleInnerSVG,
		children,
	)
}

func UserCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userCircleLineInnerSVG,
		children,
	)
}

func UserGroup(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userGroupInnerSVG,
		children,
	)
}

func UserGroupLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userGroupLineInnerSVG,
		children,
	)
}

func UserLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userLineInnerSVG,
		children,
	)
}

func UserRemove(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userRemoveInnerSVG,
		children,
	)
}

func UserRemoveLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userRemoveLineInnerSVG,
		children,
	)
}

func Users(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		usersInnerSVG,
		children,
	)
}

func UsersLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		usersLineInnerSVG,
		children,
	)
}

func UxCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uxCircleInnerSVG,
		children,
	)
}

func UxCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uxCircleLineInnerSVG,
		children,
	)
}

func Video(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoInnerSVG,
		children,
	)
}

func VideoCamera(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoCameraInnerSVG,
		children,
	)
}

func VideoCameraLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoCameraLineInnerSVG,
		children,
	)
}

func VideoLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoLineInnerSVG,
		children,
	)
}

func VideoMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoMinusInnerSVG,
		children,
	)
}

func VideoMinusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoMinusLineInnerSVG,
		children,
	)
}

func VideoPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoPlusInnerSVG,
		children,
	)
}

func VideoPlusLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoPlusLineInnerSVG,
		children,
	)
}

func ViewBoards(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewBoardsInnerSVG,
		children,
	)
}

func ViewBoardsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewBoardsLineInnerSVG,
		children,
	)
}

func ViewColumns(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewColumnsInnerSVG,
		children,
	)
}

func ViewColumnsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewColumnsLineInnerSVG,
		children,
	)
}

func ViewGrid(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewGridInnerSVG,
		children,
	)
}

func ViewGridLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewGridLineInnerSVG,
		children,
	)
}

func ViewList(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewListInnerSVG,
		children,
	)
}

func ViewListLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewListLineInnerSVG,
		children,
	)
}

func ViewRows(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewRowsInnerSVG,
		children,
	)
}

func ViewRowsLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewRowsLineInnerSVG,
		children,
	)
}

func Watch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		watchInnerSVG,
		children,
	)
}

func WatchLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		watchLineInnerSVG,
		children,
	)
}

func Wifi(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		wifiInnerSVG,
		children,
	)
}

func WifiLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		wifiLineInnerSVG,
		children,
	)
}

func YenCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		yenCircleInnerSVG,
		children,
	)
}

func YenCircleLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		yenCircleLineInnerSVG,
		children,
	)
}

func ZoomIn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		zoomInInnerSVG,
		children,
	)
}

func ZoomInLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		zoomInLineInnerSVG,
		children,
	)
}

func ZoomOut(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		zoomOutInnerSVG,
		children,
	)
}

func ZoomOutLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		zoomOutLineInnerSVG,
		children,
	)
}

func ByName(name string) (*engine.HTMLElement, error) {
	switch name {
	case "academic-cap":
		return AcademicCap(), nil
	case "academic-cap-line":
		return AcademicCapLine(), nil
	case "add-column":
		return AddColumn(), nil
	case "add-column-line":
		return AddColumnLine(), nil
	case "add-row":
		return AddRow(), nil
	case "add-row-line":
		return AddRowLine(), nil
	case "adjustments":
		return Adjustments(), nil
	case "adjustments-line":
		return AdjustmentsLine(), nil
	case "airplane":
		return Airplane(), nil
	case "airplane-flight-2":
		return AirplaneFlightTwo(), nil
	case "airplane-flight-2-line":
		return AirplaneFlightTwoLine(), nil
	case "airplane-line":
		return AirplaneLine(), nil
	case "alert-circle":
		return AlertCircle(), nil
	case "alert-circle-line":
		return AlertCircleLine(), nil
	case "align-bottom":
		return AlignBottom(), nil
	case "align-bottom-line":
		return AlignBottomLine(), nil
	case "align-horizontal-center":
		return AlignHorizontalCenter(), nil
	case "align-horizontal-center-line":
		return AlignHorizontalCenterLine(), nil
	case "align-left":
		return AlignLeft(), nil
	case "align-left-line":
		return AlignLeftLine(), nil
	case "align-right":
		return AlignRight(), nil
	case "align-right-line":
		return AlignRightLine(), nil
	case "align-top":
		return AlignTop(), nil
	case "align-top-line":
		return AlignTopLine(), nil
	case "align-vertical-center":
		return AlignVerticalCenter(), nil
	case "align-vertical-center-line":
		return AlignVerticalCenterLine(), nil
	case "analytics":
		return Analytics(), nil
	case "analytics-delete":
		return AnalyticsDelete(), nil
	case "analytics-delete-line":
		return AnalyticsDeleteLine(), nil
	case "analytics-line":
		return AnalyticsLine(), nil
	case "analytics-plus":
		return AnalyticsPlus(), nil
	case "analytics-plus-line":
		return AnalyticsPlusLine(), nil
	case "analytics-restricted":
		return AnalyticsRestricted(), nil
	case "analytics-restricted-line":
		return AnalyticsRestrictedLine(), nil
	case "annotation":
		return Annotation(), nil
	case "annotation-line":
		return AnnotationLine(), nil
	case "applications":
		return Applications(), nil
	case "applications-add":
		return ApplicationsAdd(), nil
	case "applications-add-line":
		return ApplicationsAddLine(), nil
	case "applications-line":
		return ApplicationsLine(), nil
	case "archive":
		return Archive(), nil
	case "archive-line":
		return ArchiveLine(), nil
	case "arrow-circle-down":
		return ArrowCircleDown(), nil
	case "arrow-circle-down-line":
		return ArrowCircleDownLine(), nil
	case "arrow-circle-left":
		return ArrowCircleLeft(), nil
	case "arrow-circle-left-line":
		return ArrowCircleLeftLine(), nil
	case "arrow-circle-right":
		return ArrowCircleRight(), nil
	case "arrow-circle-right-line":
		return ArrowCircleRightLine(), nil
	case "arrow-circle-up":
		return ArrowCircleUp(), nil
	case "arrow-circle-up-line":
		return ArrowCircleUpLine(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-down-circle":
		return ArrowDownCircle(), nil
	case "arrow-down-circle-line":
		return ArrowDownCircleLine(), nil
	case "arrow-down-line":
		return ArrowDownLine(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-left-circle":
		return ArrowLeftCircle(), nil
	case "arrow-left-circle-line":
		return ArrowLeftCircleLine(), nil
	case "arrow-left-line":
		return ArrowLeftLine(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-right-circle":
		return ArrowRightCircle(), nil
	case "arrow-right-circle-line":
		return ArrowRightCircleLine(), nil
	case "arrow-right-line":
		return ArrowRightLine(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrow-up-circle":
		return ArrowUpCircle(), nil
	case "arrow-up-circle-line":
		return ArrowUpCircleLine(), nil
	case "arrow-up-line":
		return ArrowUpLine(), nil
	case "arrows-collapse-full":
		return ArrowsCollapseFull(), nil
	case "arrows-collapse-full-line":
		return ArrowsCollapseFullLine(), nil
	case "arrows-expand":
		return ArrowsExpand(), nil
	case "arrows-expand-full":
		return ArrowsExpandFull(), nil
	case "arrows-expand-full-line":
		return ArrowsExpandFullLine(), nil
	case "arrows-expand-line":
		return ArrowsExpandLine(), nil
	case "article":
		return Article(), nil
	case "article-line":
		return ArticleLine(), nil
	case "article-search":
		return ArticleSearch(), nil
	case "article-search-line":
		return ArticleSearchLine(), nil
	case "at-symbol":
		return AtSymbol(), nil
	case "at-symbol-line":
		return AtSymbolLine(), nil
	case "atom-2":
		return AtomTwo(), nil
	case "atom-2-line":
		return AtomTwoLine(), nil
	case "attachment":
		return Attachment(), nil
	case "attachment-line":
		return AttachmentLine(), nil
	case "award":
		return Award(), nil
	case "award-line":
		return AwardLine(), nil
	case "back-circle":
		return BackCircle(), nil
	case "back-circle-line":
		return BackCircleLine(), nil
	case "backspace":
		return Backspace(), nil
	case "backspace-line":
		return BackspaceLine(), nil
	case "backward-circle":
		return BackwardCircle(), nil
	case "backward-circle-line":
		return BackwardCircleLine(), nil
	case "backward-start-circle":
		return BackwardStartCircle(), nil
	case "backward-start-circle-line":
		return BackwardStartCircleLine(), nil
	case "badge-check":
		return BadgeCheck(), nil
	case "badge-check-line":
		return BadgeCheckLine(), nil
	case "ban":
		return Ban(), nil
	case "ban-line":
		return BanLine(), nil
	case "band-aids":
		return BandAids(), nil
	case "band-aids-line":
		return BandAidsLine(), nil
	case "barcode-2":
		return BarcodeTwo(), nil
	case "barcode-2-line":
		return BarcodeTwoLine(), nil
	case "basket-2":
		return BasketTwo(), nil
	case "basket-2-line":
		return BasketTwoLine(), nil
	case "bath-shower":
		return BathShower(), nil
	case "bath-shower-line":
		return BathShowerLine(), nil
	case "battery":
		return Battery(), nil
	case "battery-full":
		return BatteryFull(), nil
	case "battery-full-line":
		return BatteryFullLine(), nil
	case "battery-half":
		return BatteryHalf(), nil
	case "battery-half-line":
		return BatteryHalfLine(), nil
	case "battery-line":
		return BatteryLine(), nil
	case "battery-low":
		return BatteryLow(), nil
	case "battery-low-line":
		return BatteryLowLine(), nil
	case "beach":
		return Beach(), nil
	case "beach-line":
		return BeachLine(), nil
	case "beaker":
		return Beaker(), nil
	case "beaker-line":
		return BeakerLine(), nil
	case "bell":
		return Bell(), nil
	case "bell-line":
		return BellLine(), nil
	case "bitcoin-circle":
		return BitcoinCircle(), nil
	case "bitcoin-circle-line":
		return BitcoinCircleLine(), nil
	case "bluetooth":
		return Bluetooth(), nil
	case "bluetooth-line":
		return BluetoothLine(), nil
	case "bold":
		return Bold(), nil
	case "bold-line":
		return BoldLine(), nil
	case "book":
		return Book(), nil
	case "book-line":
		return BookLine(), nil
	case "book-minus":
		return BookMinus(), nil
	case "book-minus-line":
		return BookMinusLine(), nil
	case "book-open":
		return BookOpen(), nil
	case "book-open-line":
		return BookOpenLine(), nil
	case "book-plus":
		return BookPlus(), nil
	case "book-plus-line":
		return BookPlusLine(), nil
	case "bookmark":
		return Bookmark(), nil
	case "bookmark-book":
		return BookmarkBook(), nil
	case "bookmark-book-line":
		return BookmarkBookLine(), nil
	case "bookmark-line":
		return BookmarkLine(), nil
	case "bookmark-minus":
		return BookmarkMinus(), nil
	case "bookmark-minus-line":
		return BookmarkMinusLine(), nil
	case "bookmark-plus":
		return BookmarkPlus(), nil
	case "bookmark-plus-line":
		return BookmarkPlusLine(), nil
	case "box":
		return Box(), nil
	case "box-line":
		return BoxLine(), nil
	case "briefcase":
		return Briefcase(), nil
	case "briefcase-line":
		return BriefcaseLine(), nil
	case "browser":
		return Browser(), nil
	case "browser-cookie":
		return BrowserCookie(), nil
	case "browser-cookie-line":
		return BrowserCookieLine(), nil
	case "browser-line":
		return BrowserLine(), nil
	case "bug-2":
		return BugTwo(), nil
	case "bug-2-line":
		return BugTwoLine(), nil
	case "burger":
		return Burger(), nil
	case "burger-line":
		return BurgerLine(), nil
	case "bus":
		return Bus(), nil
	case "bus-line":
		return BusLine(), nil
	case "cake":
		return Cake(), nil
	case "cake-line":
		return CakeLine(), nil
	case "calculator":
		return Calculator(), nil
	case "calculator-line":
		return CalculatorLine(), nil
	case "calendar":
		return Calendar(), nil
	case "calendar-line":
		return CalendarLine(), nil
	case "calendar-plus":
		return CalendarPlus(), nil
	case "calendar-plus-line":
		return CalendarPlusLine(), nil
	case "camera":
		return Camera(), nil
	case "camera-line":
		return CameraLine(), nil
	case "camera-off":
		return CameraOff(), nil
	case "camera-off-line":
		return CameraOffLine(), nil
	case "car":
		return Car(), nil
	case "car-line":
		return CarLine(), nil
	case "cash":
		return Cash(), nil
	case "cash-line":
		return CashLine(), nil
	case "cent-circle":
		return CentCircle(), nil
	case "cent-circle-line":
		return CentCircleLine(), nil
	case "chart-bar":
		return ChartBar(), nil
	case "chart-bar-line":
		return ChartBarLine(), nil
	case "chart-pie":
		return ChartPie(), nil
	case "chart-pie-line":
		return ChartPieLine(), nil
	case "chat":
		return Chat(), nil
	case "chat-line":
		return ChatLine(), nil
	case "chat-signal":
		return ChatSignal(), nil
	case "chat-signal-line":
		return ChatSignalLine(), nil
	case "chat-status":
		return ChatStatus(), nil
	case "chat-status-line":
		return ChatStatusLine(), nil
	case "chat-text":
		return ChatText(), nil
	case "chat-text-line":
		return ChatTextLine(), nil
	case "chat-2":
		return ChatTwo(), nil
	case "chat-2-line":
		return ChatTwoLine(), nil
	case "chat-2-text":
		return ChatTwoText(), nil
	case "chat-2-text-line":
		return ChatTwoTextLine(), nil
	case "chats":
		return Chats(), nil
	case "chats-line":
		return ChatsLine(), nil
	case "chats-2":
		return ChatsTwo(), nil
	case "chats-2-line":
		return ChatsTwoLine(), nil
	case "check":
		return Check(), nil
	case "check-circle":
		return CheckCircle(), nil
	case "check-circle-line":
		return CheckCircleLine(), nil
	case "check-line":
		return CheckLine(), nil
	case "checkbox-list":
		return CheckboxList(), nil
	case "checkbox-list-detail":
		return CheckboxListDetail(), nil
	case "checkbox-list-detail-line":
		return CheckboxListDetailLine(), nil
	case "checkbox-list-line":
		return CheckboxListLine(), nil
	case "cheese":
		return Cheese(), nil
	case "cheese-line":
		return CheeseLine(), nil
	case "chevron-double-down":
		return ChevronDoubleDown(), nil
	case "chevron-double-down-line":
		return ChevronDoubleDownLine(), nil
	case "chevron-double-left":
		return ChevronDoubleLeft(), nil
	case "chevron-double-left-line":
		return ChevronDoubleLeftLine(), nil
	case "chevron-double-right":
		return ChevronDoubleRight(), nil
	case "chevron-double-right-line":
		return ChevronDoubleRightLine(), nil
	case "chevron-double-up":
		return ChevronDoubleUp(), nil
	case "chevron-double-up-line":
		return ChevronDoubleUpLine(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-down-circle":
		return ChevronDownCircle(), nil
	case "chevron-down-circle-line":
		return ChevronDownCircleLine(), nil
	case "chevron-down-line":
		return ChevronDownLine(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-left-circle":
		return ChevronLeftCircle(), nil
	case "chevron-left-circle-line":
		return ChevronLeftCircleLine(), nil
	case "chevron-left-line":
		return ChevronLeftLine(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-right-circle":
		return ChevronRightCircle(), nil
	case "chevron-right-circle-line":
		return ChevronRightCircleLine(), nil
	case "chevron-right-line":
		return ChevronRightLine(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "chevron-up-circle":
		return ChevronUpCircle(), nil
	case "chevron-up-circle-line":
		return ChevronUpCircleLine(), nil
	case "chevron-up-line":
		return ChevronUpLine(), nil
	case "chip":
		return Chip(), nil
	case "chip-line":
		return ChipLine(), nil
	case "chromecast":
		return Chromecast(), nil
	case "chromecast-line":
		return ChromecastLine(), nil
	case "church":
		return Church(), nil
	case "church-line":
		return ChurchLine(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clipboard-check":
		return ClipboardCheck(), nil
	case "clipboard-check-line":
		return ClipboardCheckLine(), nil
	case "clipboard-copy":
		return ClipboardCopy(), nil
	case "clipboard-copy-line":
		return ClipboardCopyLine(), nil
	case "clipboard-line":
		return ClipboardLine(), nil
	case "clipboard-list":
		return ClipboardList(), nil
	case "clipboard-list-line":
		return ClipboardListLine(), nil
	case "clipboard-minus":
		return ClipboardMinus(), nil
	case "clipboard-minus-line":
		return ClipboardMinusLine(), nil
	case "clipboard-plus":
		return ClipboardPlus(), nil
	case "clipboard-plus-line":
		return ClipboardPlusLine(), nil
	case "clock":
		return Clock(), nil
	case "clock-line":
		return ClockLine(), nil
	case "clock-plus":
		return ClockPlus(), nil
	case "clock-plus-line":
		return ClockPlusLine(), nil
	case "close":
		return Close(), nil
	case "close-circle":
		return CloseCircle(), nil
	case "close-circle-line":
		return CloseCircleLine(), nil
	case "close-line":
		return CloseLine(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-download":
		return CloudDownload(), nil
	case "cloud-download-line":
		return CloudDownloadLine(), nil
	case "cloud-line":
		return CloudLine(), nil
	case "cloud-upload":
		return CloudUpload(), nil
	case "cloud-upload-line":
		return CloudUploadLine(), nil
	case "code":
		return Code(), nil
	case "code-block":
		return CodeBlock(), nil
	case "code-block-line":
		return CodeBlockLine(), nil
	case "code-line":
		return CodeLine(), nil
	case "cog":
		return Cog(), nil
	case "cog-line":
		return CogLine(), nil
	case "coins":
		return Coins(), nil
	case "coins-line":
		return CoinsLine(), nil
	case "collection":
		return Collection(), nil
	case "collection-line":
		return CollectionLine(), nil
	case "color-swatch":
		return ColorSwatch(), nil
	case "color-swatch-line":
		return ColorSwatchLine(), nil
	case "comet":
		return Comet(), nil
	case "comet-line":
		return CometLine(), nil
	case "comment":
		return Comment(), nil
	case "comment-line":
		return CommentLine(), nil
	case "comment-text":
		return CommentText(), nil
	case "comment-text-line":
		return CommentTextLine(), nil
	case "comment-2":
		return CommentTwo(), nil
	case "comment-2-line":
		return CommentTwoLine(), nil
	case "comment-2-text":
		return CommentTwoText(), nil
	case "comment-2-text-line":
		return CommentTwoTextLine(), nil
	case "comments":
		return Comments(), nil
	case "comments-line":
		return CommentsLine(), nil
	case "comments-2":
		return CommentsTwo(), nil
	case "comments-2-line":
		return CommentsTwoLine(), nil
	case "community":
		return Community(), nil
	case "community-line":
		return CommunityLine(), nil
	case "compass":
		return Compass(), nil
	case "compass-line":
		return CompassLine(), nil
	case "compass-2":
		return CompassTwo(), nil
	case "compass-2-line":
		return CompassTwoLine(), nil
	case "contact":
		return Contact(), nil
	case "contact-line":
		return ContactLine(), nil
	case "cookie":
		return Cookie(), nil
	case "cookie-line":
		return CookieLine(), nil
	case "covid":
		return Covid(), nil
	case "covid-exclamation":
		return CovidExclamation(), nil
	case "covid-exclamation-line":
		return CovidExclamationLine(), nil
	case "covid-line":
		return CovidLine(), nil
	case "covid-off":
		return CovidOff(), nil
	case "covid-off-line":
		return CovidOffLine(), nil
	case "cpu":
		return Cpu(), nil
	case "cpu-line":
		return CpuLine(), nil
	case "credit-card":
		return CreditCard(), nil
	case "credit-card-line":
		return CreditCardLine(), nil
	case "creditcard":
		return Creditcard(), nil
	case "creditcard-hand":
		return CreditcardHand(), nil
	case "creditcard-hand-line":
		return CreditcardHandLine(), nil
	case "creditcard-line":
		return CreditcardLine(), nil
	case "creditcard-plus":
		return CreditcardPlus(), nil
	case "creditcard-plus-line":
		return CreditcardPlusLine(), nil
	case "crown":
		return Crown(), nil
	case "crown-line":
		return CrownLine(), nil
	case "cube":
		return Cube(), nil
	case "cube-line":
		return CubeLine(), nil
	case "cup":
		return Cup(), nil
	case "cup-line":
		return CupLine(), nil
	case "curly-braces":
		return CurlyBraces(), nil
	case "curly-braces-line":
		return CurlyBracesLine(), nil
	case "cursor-click":
		return CursorClick(), nil
	case "cursor-click-line":
		return CursorClickLine(), nil
	case "dashboard":
		return Dashboard(), nil
	case "dashboard-line":
		return DashboardLine(), nil
	case "data":
		return Data(), nil
	case "data-line":
		return DataLine(), nil
	case "data-minus":
		return DataMinus(), nil
	case "data-minus-line":
		return DataMinusLine(), nil
	case "data-plus":
		return DataPlus(), nil
	case "data-plus-line":
		return DataPlusLine(), nil
	case "database":
		return Database(), nil
	case "database-line":
		return DatabaseLine(), nil
	case "delete-bin":
		return DeleteBin(), nil
	case "delete-bin-line":
		return DeleteBinLine(), nil
	case "desktop-computer":
		return DesktopComputer(), nil
	case "desktop-computer-line":
		return DesktopComputerLine(), nil
	case "device-mobile":
		return DeviceMobile(), nil
	case "device-mobile-line":
		return DeviceMobileLine(), nil
	case "device-tablet":
		return DeviceTablet(), nil
	case "device-tablet-line":
		return DeviceTabletLine(), nil
	case "distribute-horizontal":
		return DistributeHorizontal(), nil
	case "distribute-horizontal-line":
		return DistributeHorizontalLine(), nil
	case "distribute-vertical":
		return DistributeVertical(), nil
	case "distribute-vertical-line":
		return DistributeVerticalLine(), nil
	case "divide":
		return Divide(), nil
	case "divide-line":
		return DivideLine(), nil
	case "document":
		return Document(), nil
	case "document-award":
		return DocumentAward(), nil
	case "document-award-line":
		return DocumentAwardLine(), nil
	case "document-line":
		return DocumentLine(), nil
	case "dollar-circle":
		return DollarCircle(), nil
	case "dollar-circle-line":
		return DollarCircleLine(), nil
	case "door-enter":
		return DoorEnter(), nil
	case "door-enter-line":
		return DoorEnterLine(), nil
	case "door-exit":
		return DoorExit(), nil
	case "door-exit-line":
		return DoorExitLine(), nil
	case "dots-circle":
		return DotsCircle(), nil
	case "dots-circle-line":
		return DotsCircleLine(), nil
	case "dots-horizontal":
		return DotsHorizontal(), nil
	case "dots-horizontal-line":
		return DotsHorizontalLine(), nil
	case "dots-vertical":
		return DotsVertical(), nil
	case "dots-vertical-line":
		return DotsVerticalLine(), nil
	case "duplicate":
		return Duplicate(), nil
	case "duplicate-line":
		return DuplicateLine(), nil
	case "earth-sphere":
		return EarthSphere(), nil
	case "earth-sphere-line":
		return EarthSphereLine(), nil
	case "edit-pen-4":
		return EditPenFour(), nil
	case "edit-pen-4-line":
		return EditPenFourLine(), nil
	case "edit-pen-2":
		return EditPenTwo(), nil
	case "edit-pen-2-line":
		return EditPenTwoLine(), nil
	case "eject":
		return Eject(), nil
	case "eject-line":
		return EjectLine(), nil
	case "emoji-happy":
		return EmojiHappy(), nil
	case "emoji-happy-line":
		return EmojiHappyLine(), nil
	case "emoji-sad":
		return EmojiSad(), nil
	case "emoji-sad-line":
		return EmojiSadLine(), nil
	case "eraser":
		return Eraser(), nil
	case "eraser-line":
		return EraserLine(), nil
	case "etherium-circle":
		return EtheriumCircle(), nil
	case "etherium-circle-line":
		return EtheriumCircleLine(), nil
	case "euro-circle":
		return EuroCircle(), nil
	case "euro-circle-line":
		return EuroCircleLine(), nil
	case "exclamation":
		return Exclamation(), nil
	case "exclamation-circle":
		return ExclamationCircle(), nil
	case "exclamation-circle-line":
		return ExclamationCircleLine(), nil
	case "exclamation-line":
		return ExclamationLine(), nil
	case "external-link":
		return ExternalLink(), nil
	case "external-link-line":
		return ExternalLinkLine(), nil
	case "eye":
		return Eye(), nil
	case "eye-line":
		return EyeLine(), nil
	case "eye-off":
		return EyeOff(), nil
	case "eye-off-line":
		return EyeOffLine(), nil
	case "fast-forward":
		return FastForward(), nil
	case "fast-forward-line":
		return FastForwardLine(), nil
	case "ferris-wheel":
		return FerrisWheel(), nil
	case "ferris-wheel-line":
		return FerrisWheelLine(), nil
	case "file":
		return File(), nil
	case "file-add":
		return FileAdd(), nil
	case "file-add-line":
		return FileAddLine(), nil
	case "file-download":
		return FileDownload(), nil
	case "file-download-line":
		return FileDownloadLine(), nil
	case "file-duplicate":
		return FileDuplicate(), nil
	case "file-duplicate-line":
		return FileDuplicateLine(), nil
	case "file-line":
		return FileLine(), nil
	case "file-minus":
		return FileMinus(), nil
	case "file-minus-line":
		return FileMinusLine(), nil
	case "file-plus":
		return FilePlus(), nil
	case "file-plus-line":
		return FilePlusLine(), nil
	case "file-remove":
		return FileRemove(), nil
	case "file-remove-line":
		return FileRemoveLine(), nil
	case "file-report":
		return FileReport(), nil
	case "file-report-line":
		return FileReportLine(), nil
	case "file-search":
		return FileSearch(), nil
	case "file-search-line":
		return FileSearchLine(), nil
	case "file-text":
		return FileText(), nil
	case "file-text-line":
		return FileTextLine(), nil
	case "film":
		return Film(), nil
	case "film-line":
		return FilmLine(), nil
	case "filter":
		return Filter(), nil
	case "filter-line":
		return FilterLine(), nil
	case "fire":
		return Fire(), nil
	case "fire-line":
		return FireLine(), nil
	case "fish":
		return Fish(), nil
	case "fish-line":
		return FishLine(), nil
	case "flag":
		return Flag(), nil
	case "flag-line":
		return FlagLine(), nil
	case "flask":
		return Flask(), nil
	case "flask-line":
		return FlaskLine(), nil
	case "flower-2":
		return FlowerTwo(), nil
	case "flower-2-line":
		return FlowerTwoLine(), nil
	case "folder":
		return Folder(), nil
	case "folder-add":
		return FolderAdd(), nil
	case "folder-add-line":
		return FolderAddLine(), nil
	case "folder-check":
		return FolderCheck(), nil
	case "folder-check-line":
		return FolderCheckLine(), nil
	case "folder-download":
		return FolderDownload(), nil
	case "folder-download-line":
		return FolderDownloadLine(), nil
	case "folder-line":
		return FolderLine(), nil
	case "folder-minus":
		return FolderMinus(), nil
	case "folder-minus-line":
		return FolderMinusLine(), nil
	case "folder-open":
		return FolderOpen(), nil
	case "folder-open-line":
		return FolderOpenLine(), nil
	case "folder-plus":
		return FolderPlus(), nil
	case "folder-plus-line":
		return FolderPlusLine(), nil
	case "folder-remove":
		return FolderRemove(), nil
	case "folder-remove-line":
		return FolderRemoveLine(), nil
	case "font-size":
		return FontSize(), nil
	case "font-size-line":
		return FontSizeLine(), nil
	case "forward":
		return Forward(), nil
	case "forward-circle":
		return ForwardCircle(), nil
	case "forward-circle-line":
		return ForwardCircleLine(), nil
	case "forward-end-circle":
		return ForwardEndCircle(), nil
	case "forward-end-circle-line":
		return ForwardEndCircleLine(), nil
	case "forward-line":
		return ForwardLine(), nil
	case "ghost":
		return Ghost(), nil
	case "ghost-line":
		return GhostLine(), nil
	case "git-branch":
		return GitBranch(), nil
	case "git-branch-line":
		return GitBranchLine(), nil
	case "git-commit":
		return GitCommit(), nil
	case "git-commit-line":
		return GitCommitLine(), nil
	case "git-compare":
		return GitCompare(), nil
	case "git-compare-line":
		return GitCompareLine(), nil
	case "git-fork":
		return GitFork(), nil
	case "git-fork-line":
		return GitForkLine(), nil
	case "git-merge":
		return GitMerge(), nil
	case "git-merge-line":
		return GitMergeLine(), nil
	case "git-pull":
		return GitPull(), nil
	case "git-pull-line":
		return GitPullLine(), nil
	case "glas-water":
		return GlasWater(), nil
	case "glas-water-line":
		return GlasWaterLine(), nil
	case "globe":
		return Globe(), nil
	case "globe-earth":
		return GlobeEarth(), nil
	case "globe-earth-line":
		return GlobeEarthLine(), nil
	case "globe-earth-2":
		return GlobeEarthTwo(), nil
	case "globe-earth-2-line":
		return GlobeEarthTwoLine(), nil
	case "globe-grid":
		return GlobeGrid(), nil
	case "globe-grid-line":
		return GlobeGridLine(), nil
	case "globe-line":
		return GlobeLine(), nil
	case "hand":
		return Hand(), nil
	case "hand-line":
		return HandLine(), nil
	case "hand-pointer":
		return HandPointer(), nil
	case "hand-pointer-event":
		return HandPointerEvent(), nil
	case "hand-pointer-event-line":
		return HandPointerEventLine(), nil
	case "hand-pointer-line":
		return HandPointerLine(), nil
	case "hand-pointer-2":
		return HandPointerTwo(), nil
	case "hand-pointer-2-line":
		return HandPointerTwoLine(), nil
	case "hard-drive":
		return HardDrive(), nil
	case "hard-drive-line":
		return HardDriveLine(), nil
	case "hashtag":
		return Hashtag(), nil
	case "hashtag-line":
		return HashtagLine(), nil
	case "headset":
		return Headset(), nil
	case "headset-line":
		return HeadsetLine(), nil
	case "heart":
		return Heart(), nil
	case "heart-line":
		return HeartLine(), nil
	case "home":
		return Home(), nil
	case "home-analytics":
		return HomeAnalytics(), nil
	case "home-analytics-line":
		return HomeAnalyticsLine(), nil
	case "home-line":
		return HomeLine(), nil
	case "home-simple":
		return HomeSimple(), nil
	case "home-simple-line":
		return HomeSimpleLine(), nil
	case "image":
		return Image(), nil
	case "image-circle":
		return ImageCircle(), nil
	case "image-circle-line":
		return ImageCircleLine(), nil
	case "image-circle-off":
		return ImageCircleOff(), nil
	case "image-circle-off-line":
		return ImageCircleOffLine(), nil
	case "image-circle-plus":
		return ImageCirclePlus(), nil
	case "image-circle-plus-line":
		return ImageCirclePlusLine(), nil
	case "image-circle-story":
		return ImageCircleStory(), nil
	case "image-circle-story-line":
		return ImageCircleStoryLine(), nil
	case "image-frame":
		return ImageFrame(), nil
	case "image-frame-line":
		return ImageFrameLine(), nil
	case "image-in-picture":
		return ImageInPicture(), nil
	case "image-in-picture-line":
		return ImageInPictureLine(), nil
	case "image-line":
		return ImageLine(), nil
	case "image-multiple":
		return ImageMultiple(), nil
	case "image-multiple-line":
		return ImageMultipleLine(), nil
	case "image-off":
		return ImageOff(), nil
	case "image-off-line":
		return ImageOffLine(), nil
	case "image-photography":
		return ImagePhotography(), nil
	case "image-photography-line":
		return ImagePhotographyLine(), nil
	case "image-plus":
		return ImagePlus(), nil
	case "image-plus-line":
		return ImagePlusLine(), nil
	case "inbox":
		return Inbox(), nil
	case "inbox-in":
		return InboxIn(), nil
	case "inbox-in-line":
		return InboxInLine(), nil
	case "inbox-line":
		return InboxLine(), nil
	case "incognito":
		return Incognito(), nil
	case "incognito-line":
		return IncognitoLine(), nil
	case "info-circle":
		return InfoCircle(), nil
	case "info-circle-line":
		return InfoCircleLine(), nil
	case "information-circle":
		return InformationCircle(), nil
	case "information-circle-line":
		return InformationCircleLine(), nil
	case "iphone-old-apps":
		return IphoneOldApps(), nil
	case "iphone-old-apps-line":
		return IphoneOldAppsLine(), nil
	case "iphone-x-apps":
		return IphoneXApps(), nil
	case "iphone-x-apps-line":
		return IphoneXAppsLine(), nil
	case "italic":
		return Italic(), nil
	case "italic-line":
		return ItalicLine(), nil
	case "key":
		return Key(), nil
	case "key-line":
		return KeyLine(), nil
	case "keyboard":
		return Keyboard(), nil
	case "keyboard-line":
		return KeyboardLine(), nil
	case "laptop":
		return Laptop(), nil
	case "laptop-line":
		return LaptopLine(), nil
	case "leaf-3-angled":
		return LeafThreeAngled(), nil
	case "leaf-3-angled-line":
		return LeafThreeAngledLine(), nil
	case "library":
		return Library(), nil
	case "library-line":
		return LibraryLine(), nil
	case "lidquid-drop-waves-2":
		return LidquidDropWavesTwo(), nil
	case "lidquid-drop-waves-2-line":
		return LidquidDropWavesTwoLine(), nil
	case "lifebuoy":
		return Lifebuoy(), nil
	case "lifebuoy-line":
		return LifebuoyLine(), nil
	case "light-bulb":
		return LightBulb(), nil
	case "light-bulb-line":
		return LightBulbLine(), nil
	case "lightbulb-shine":
		return LightbulbShine(), nil
	case "lightbulb-shine-line":
		return LightbulbShineLine(), nil
	case "lightning-bolt":
		return LightningBolt(), nil
	case "lightning-bolt-line":
		return LightningBoltLine(), nil
	case "line-height":
		return LineHeight(), nil
	case "line-height-line":
		return LineHeightLine(), nil
	case "link":
		return Link(), nil
	case "link-circle":
		return LinkCircle(), nil
	case "link-circle-line":
		return LinkCircleLine(), nil
	case "link-line":
		return LinkLine(), nil
	case "lira-circle":
		return LiraCircle(), nil
	case "lira-circle-line":
		return LiraCircleLine(), nil
	case "list-box":
		return ListBox(), nil
	case "list-box-line":
		return ListBoxLine(), nil
	case "location-marker":
		return LocationMarker(), nil
	case "location-marker-line":
		return LocationMarkerLine(), nil
	case "lock":
		return Lock(), nil
	case "lock-closed":
		return LockClosed(), nil
	case "lock-closed-line":
		return LockClosedLine(), nil
	case "lock-line":
		return LockLine(), nil
	case "lock-off":
		return LockOff(), nil
	case "lock-off-line":
		return LockOffLine(), nil
	case "lock-open":
		return LockOpen(), nil
	case "lock-open-line":
		return LockOpenLine(), nil
	case "login":
		return Login(), nil
	case "login-half-circle":
		return LoginHalfCircle(), nil
	case "login-half-circle-line":
		return LoginHalfCircleLine(), nil
	case "login-line":
		return LoginLine(), nil
	case "logout":
		return Logout(), nil
	case "logout-half-circle":
		return LogoutHalfCircle(), nil
	case "logout-half-circle-line":
		return LogoutHalfCircleLine(), nil
	case "logout-line":
		return LogoutLine(), nil
	case "mail":
		return Mail(), nil
	case "mail-line":
		return MailLine(), nil
	case "mail-open":
		return MailOpen(), nil
	case "mail-open-line":
		return MailOpenLine(), nil
	case "map":
		return Map(), nil
	case "map-line":
		return MapLine(), nil
	case "map-marker":
		return MapMarker(), nil
	case "map-marker-area":
		return MapMarkerArea(), nil
	case "map-marker-area-line":
		return MapMarkerAreaLine(), nil
	case "map-marker-line":
		return MapMarkerLine(), nil
	case "map-marker-path":
		return MapMarkerPath(), nil
	case "map-marker-path-line":
		return MapMarkerPathLine(), nil
	case "map-marker-plus":
		return MapMarkerPlus(), nil
	case "map-marker-plus-line":
		return MapMarkerPlusLine(), nil
	case "map-simple":
		return MapSimple(), nil
	case "map-simple-destination":
		return MapSimpleDestination(), nil
	case "map-simple-destination-line":
		return MapSimpleDestinationLine(), nil
	case "map-simple-line":
		return MapSimpleLine(), nil
	case "map-simple-marker":
		return MapSimpleMarker(), nil
	case "map-simple-marker-line":
		return MapSimpleMarkerLine(), nil
	case "map-simple-off":
		return MapSimpleOff(), nil
	case "map-simple-off-line":
		return MapSimpleOffLine(), nil
	case "maximize":
		return Maximize(), nil
	case "maximize-line":
		return MaximizeLine(), nil
	case "megaphone":
		return Megaphone(), nil
	case "megaphone-line":
		return MegaphoneLine(), nil
	case "menu":
		return Menu(), nil
	case "menu-alt":
		return MenuAlt(), nil
	case "menu-alt-line":
		return MenuAltLine(), nil
	case "menu-expand-left":
		return MenuExpandLeft(), nil
	case "menu-expand-left-line":
		return MenuExpandLeftLine(), nil
	case "menu-expand-right":
		return MenuExpandRight(), nil
	case "menu-expand-right-line":
		return MenuExpandRightLine(), nil
	case "menu-line":
		return MenuLine(), nil
	case "message":
		return Message(), nil
	case "message-line":
		return MessageLine(), nil
	case "messages":
		return Messages(), nil
	case "messages-line":
		return MessagesLine(), nil
	case "microphone":
		return Microphone(), nil
	case "microphone-line":
		return MicrophoneLine(), nil
	case "minimize":
		return Minimize(), nil
	case "minimize-line":
		return MinimizeLine(), nil
	case "minus":
		return Minus(), nil
	case "minus-circle":
		return MinusCircle(), nil
	case "minus-circle-line":
		return MinusCircleLine(), nil
	case "minus-five-circle":
		return MinusFiveCircle(), nil
	case "minus-five-circle-line":
		return MinusFiveCircleLine(), nil
	case "minus-line":
		return MinusLine(), nil
	case "minus-ten-circle":
		return MinusTenCircle(), nil
	case "minus-ten-circle-line":
		return MinusTenCircleLine(), nil
	case "money":
		return Money(), nil
	case "money-hand":
		return MoneyHand(), nil
	case "money-hand-line":
		return MoneyHandLine(), nil
	case "money-line":
		return MoneyLine(), nil
	case "money-minus":
		return MoneyMinus(), nil
	case "money-minus-line":
		return MoneyMinusLine(), nil
	case "money-plus":
		return MoneyPlus(), nil
	case "money-plus-line":
		return MoneyPlusLine(), nil
	case "monitor":
		return Monitor(), nil
	case "monitor-line":
		return MonitorLine(), nil
	case "moon":
		return Moon(), nil
	case "moon-line":
		return MoonLine(), nil
	case "more-menu":
		return MoreMenu(), nil
	case "more-menu-line":
		return MoreMenuLine(), nil
	case "more-menu-vertical":
		return MoreMenuVertical(), nil
	case "more-menu-vertical-line":
		return MoreMenuVerticalLine(), nil
	case "mouse":
		return Mouse(), nil
	case "mouse-line":
		return MouseLine(), nil
	case "multiply":
		return Multiply(), nil
	case "multiply-line":
		return MultiplyLine(), nil
	case "music":
		return Music(), nil
	case "music-line":
		return MusicLine(), nil
	case "music-note":
		return MusicNote(), nil
	case "music-note-line":
		return MusicNoteLine(), nil
	case "newspaper":
		return Newspaper(), nil
	case "newspaper-line":
		return NewspaperLine(), nil
	case "next-circle":
		return NextCircle(), nil
	case "next-circle-line":
		return NextCircleLine(), nil
	case "note-text":
		return NoteText(), nil
	case "note-text-line":
		return NoteTextLine(), nil
	case "note-text-minus":
		return NoteTextMinus(), nil
	case "note-text-minus-line":
		return NoteTextMinusLine(), nil
	case "note-text-plus":
		return NoteTextPlus(), nil
	case "note-text-plus-line":
		return NoteTextPlusLine(), nil
	case "noteblock":
		return Noteblock(), nil
	case "noteblock-line":
		return NoteblockLine(), nil
	case "noteblock-text":
		return NoteblockText(), nil
	case "noteblock-text-line":
		return NoteblockTextLine(), nil
	case "open":
		return Open(), nil
	case "open-line":
		return OpenLine(), nil
	case "paper-airplane":
		return PaperAirplane(), nil
	case "paper-airplane-line":
		return PaperAirplaneLine(), nil
	case "paper-clip":
		return PaperClip(), nil
	case "paper-clip-line":
		return PaperClipLine(), nil
	case "paper-fold":
		return PaperFold(), nil
	case "paper-fold-line":
		return PaperFoldLine(), nil
	case "paper-fold-text":
		return PaperFoldText(), nil
	case "paper-fold-text-line":
		return PaperFoldTextLine(), nil
	case "paper-roll-2":
		return PaperRollTwo(), nil
	case "paper-roll-2-line":
		return PaperRollTwoLine(), nil
	case "paragraph":
		return Paragraph(), nil
	case "paragraph-line":
		return ParagraphLine(), nil
	case "pause":
		return Pause(), nil
	case "pause-circle":
		return PauseCircle(), nil
	case "pause-circle-line":
		return PauseCircleLine(), nil
	case "pause-line":
		return PauseLine(), nil
	case "pencil":
		return Pencil(), nil
	case "pencil-alt":
		return PencilAlt(), nil
	case "pencil-alt-line":
		return PencilAltLine(), nil
	case "pencil-line":
		return PencilLine(), nil
	case "percent":
		return Percent(), nil
	case "percent-line":
		return PercentLine(), nil
	case "phone":
		return Phone(), nil
	case "phone-dial":
		return PhoneDial(), nil
	case "phone-dial-line":
		return PhoneDialLine(), nil
	case "phone-hangup":
		return PhoneHangup(), nil
	case "phone-hangup-line":
		return PhoneHangupLine(), nil
	case "phone-incoming":
		return PhoneIncoming(), nil
	case "phone-incoming-line":
		return PhoneIncomingLine(), nil
	case "phone-line":
		return PhoneLine(), nil
	case "phone-missed-call":
		return PhoneMissedCall(), nil
	case "phone-missed-call-line":
		return PhoneMissedCallLine(), nil
	case "phone-outgoing":
		return PhoneOutgoing(), nil
	case "phone-outgoing-line":
		return PhoneOutgoingLine(), nil
	case "phone-retro":
		return PhoneRetro(), nil
	case "phone-retro-line":
		return PhoneRetroLine(), nil
	case "phone-ring":
		return PhoneRing(), nil
	case "phone-ring-line":
		return PhoneRingLine(), nil
	case "pill":
		return Pill(), nil
	case "pill-line":
		return PillLine(), nil
	case "pin":
		return Pin(), nil
	case "pin-line":
		return PinLine(), nil
	case "pinwheel":
		return Pinwheel(), nil
	case "pinwheel-line":
		return PinwheelLine(), nil
	case "planet":
		return Planet(), nil
	case "planet-line":
		return PlanetLine(), nil
	case "planet-ring-2":
		return PlanetRingTwo(), nil
	case "planet-ring-2-line":
		return PlanetRingTwoLine(), nil
	case "planet-rocket":
		return PlanetRocket(), nil
	case "planet-rocket-line":
		return PlanetRocketLine(), nil
	case "play-circle":
		return PlayCircle(), nil
	case "play-circle-line":
		return PlayCircleLine(), nil
	case "playlist":
		return Playlist(), nil
	case "playlist-line":
		return PlaylistLine(), nil
	case "plus":
		return Plus(), nil
	case "plus-circle":
		return PlusCircle(), nil
	case "plus-circle-line":
		return PlusCircleLine(), nil
	case "plus-five-circle":
		return PlusFiveCircle(), nil
	case "plus-five-circle-line":
		return PlusFiveCircleLine(), nil
	case "plus-line":
		return PlusLine(), nil
	case "plus-minus":
		return PlusMinus(), nil
	case "plus-minus-line":
		return PlusMinusLine(), nil
	case "plus-minus-2":
		return PlusMinusTwo(), nil
	case "plus-minus-2-line":
		return PlusMinusTwoLine(), nil
	case "plus-ten-circle":
		return PlusTenCircle(), nil
	case "plus-ten-circle-line":
		return PlusTenCircleLine(), nil
	case "pound-circle":
		return PoundCircle(), nil
	case "pound-circle-line":
		return PoundCircleLine(), nil
	case "presentation":
		return Presentation(), nil
	case "presentation-chart":
		return PresentationChart(), nil
	case "presentation-chart-line":
		return PresentationChartLine(), nil
	case "presentation-line":
		return PresentationLine(), nil
	case "presentation-line-line":
		return PresentationLineLine(), nil
	case "presentation-play":
		return PresentationPlay(), nil
	case "presentation-play-line":
		return PresentationPlayLine(), nil
	case "presentation-report":
		return PresentationReport(), nil
	case "presentation-report-line":
		return PresentationReportLine(), nil
	case "printer":
		return Printer(), nil
	case "printer-line":
		return PrinterLine(), nil
	case "pulse":
		return Pulse(), nil
	case "pulse-line":
		return PulseLine(), nil
	case "puzzle":
		return Puzzle(), nil
	case "puzzle-line":
		return PuzzleLine(), nil
	case "qr-code":
		return QrCode(), nil
	case "qr-code-line":
		return QrCodeLine(), nil
	case "qrcode":
		return Qrcode(), nil
	case "qrcode-line":
		return QrcodeLine(), nil
	case "question-circle":
		return QuestionCircle(), nil
	case "question-circle-line":
		return QuestionCircleLine(), nil
	case "question-mark-circle":
		return QuestionMarkCircle(), nil
	case "question-mark-circle-line":
		return QuestionMarkCircleLine(), nil
	case "radio-list":
		return RadioList(), nil
	case "radio-list-line":
		return RadioListLine(), nil
	case "receipt":
		return Receipt(), nil
	case "receipt-line":
		return ReceiptLine(), nil
	case "receipt-refund":
		return ReceiptRefund(), nil
	case "receipt-refund-line":
		return ReceiptRefundLine(), nil
	case "receipt-text":
		return ReceiptText(), nil
	case "receipt-text-line":
		return ReceiptTextLine(), nil
	case "redo":
		return Redo(), nil
	case "redo-line":
		return RedoLine(), nil
	case "refresh":
		return Refresh(), nil
	case "refresh-line":
		return RefreshLine(), nil
	case "reload":
		return Reload(), nil
	case "reload-circle":
		return ReloadCircle(), nil
	case "reload-circle-line":
		return ReloadCircleLine(), nil
	case "reload-line":
		return ReloadLine(), nil
	case "remove-column":
		return RemoveColumn(), nil
	case "remove-column-line":
		return RemoveColumnLine(), nil
	case "remove-format":
		return RemoveFormat(), nil
	case "remove-format-line":
		return RemoveFormatLine(), nil
	case "remove-row":
		return RemoveRow(), nil
	case "remove-row-line":
		return RemoveRowLine(), nil
	case "repeat-circle":
		return RepeatCircle(), nil
	case "repeat-circle-line":
		return RepeatCircleLine(), nil
	case "reply":
		return Reply(), nil
	case "reply-all":
		return ReplyAll(), nil
	case "reply-all-line":
		return ReplyAllLine(), nil
	case "reply-line":
		return ReplyLine(), nil
	case "restricted":
		return Restricted(), nil
	case "restricted-line":
		return RestrictedLine(), nil
	case "rewind":
		return Rewind(), nil
	case "rewind-line":
		return RewindLine(), nil
	case "robot":
		return Robot(), nil
	case "robot-line":
		return RobotLine(), nil
	case "rocket-3-start":
		return RocketThreeStart(), nil
	case "rocket-3-start-line":
		return RocketThreeStartLine(), nil
	case "rss":
		return Rss(), nil
	case "rss-line":
		return RssLine(), nil
	case "rubel-circle":
		return RubelCircle(), nil
	case "rubel-circle-line":
		return RubelCircleLine(), nil
	case "ruler-2":
		return RulerTwo(), nil
	case "ruler-2-line":
		return RulerTwoLine(), nil
	case "rupee-circle":
		return RupeeCircle(), nil
	case "rupee-circle-line":
		return RupeeCircleLine(), nil
	case "save":
		return Save(), nil
	case "save-alt":
		return SaveAlt(), nil
	case "save-alt-line":
		return SaveAltLine(), nil
	case "save-as":
		return SaveAs(), nil
	case "save-as-line":
		return SaveAsLine(), nil
	case "save-line":
		return SaveLine(), nil
	case "scale-light":
		return ScaleLight(), nil
	case "scale-light-line":
		return ScaleLightLine(), nil
	case "scan-fingerprint":
		return ScanFingerprint(), nil
	case "scan-fingerprint-line":
		return ScanFingerprintLine(), nil
	case "scan-user":
		return ScanUser(), nil
	case "scan-user-line":
		return ScanUserLine(), nil
	case "scanner":
		return Scanner(), nil
	case "scanner-line":
		return ScannerLine(), nil
	case "scissors":
		return Scissors(), nil
	case "scissors-line":
		return ScissorsLine(), nil
	case "scooter":
		return Scooter(), nil
	case "scooter-line":
		return ScooterLine(), nil
	case "script-prescription":
		return ScriptPrescription(), nil
	case "script-prescription-line":
		return ScriptPrescriptionLine(), nil
	case "scroll":
		return Scroll(), nil
	case "scroll-line":
		return ScrollLine(), nil
	case "scroll-text":
		return ScrollText(), nil
	case "scroll-text-line":
		return ScrollTextLine(), nil
	case "search":
		return Search(), nil
	case "search-circle":
		return SearchCircle(), nil
	case "search-circle-line":
		return SearchCircleLine(), nil
	case "search-line":
		return SearchLine(), nil
	case "search-minus":
		return SearchMinus(), nil
	case "search-minus-line":
		return SearchMinusLine(), nil
	case "search-plus":
		return SearchPlus(), nil
	case "search-plus-line":
		return SearchPlusLine(), nil
	case "search-text":
		return SearchText(), nil
	case "search-text-line":
		return SearchTextLine(), nil
	case "selector":
		return Selector(), nil
	case "selector-line":
		return SelectorLine(), nil
	case "send":
		return Send(), nil
	case "send-line":
		return SendLine(), nil
	case "server":
		return Server(), nil
	case "server-line":
		return ServerLine(), nil
	case "settings-cog":
		return SettingsCog(), nil
	case "settings-cog-check":
		return SettingsCogCheck(), nil
	case "settings-cog-check-line":
		return SettingsCogCheckLine(), nil
	case "settings-cog-line":
		return SettingsCogLine(), nil
	case "settings-cog-plus":
		return SettingsCogPlus(), nil
	case "settings-cog-plus-line":
		return SettingsCogPlusLine(), nil
	case "share":
		return Share(), nil
	case "share-circle":
		return ShareCircle(), nil
	case "share-circle-line":
		return ShareCircleLine(), nil
	case "share-line":
		return ShareLine(), nil
	case "shield":
		return Shield(), nil
	case "shield-check":
		return ShieldCheck(), nil
	case "shield-check-line":
		return ShieldCheckLine(), nil
	case "shield-exclamation":
		return ShieldExclamation(), nil
	case "shield-exclamation-line":
		return ShieldExclamationLine(), nil
	case "shield-line":
		return ShieldLine(), nil
	case "shield-off":
		return ShieldOff(), nil
	case "shield-off-line":
		return ShieldOffLine(), nil
	case "shield-plus":
		return ShieldPlus(), nil
	case "shield-plus-line":
		return ShieldPlusLine(), nil
	case "ship":
		return Ship(), nil
	case "ship-line":
		return ShipLine(), nil
	case "shooting-star":
		return ShootingStar(), nil
	case "shooting-star-line":
		return ShootingStarLine(), nil
	case "shopping-bag":
		return ShoppingBag(), nil
	case "shopping-bag-line":
		return ShoppingBagLine(), nil
	case "shopping-cart":
		return ShoppingCart(), nil
	case "shopping-cart-line":
		return ShoppingCartLine(), nil
	case "sim-card":
		return SimCard(), nil
	case "sim-card-line":
		return SimCardLine(), nil
	case "sitemap":
		return Sitemap(), nil
	case "sitemap-line":
		return SitemapLine(), nil
	case "skull":
		return Skull(), nil
	case "skull-line":
		return SkullLine(), nil
	case "smartphone-apps":
		return SmartphoneApps(), nil
	case "smartphone-apps-line":
		return SmartphoneAppsLine(), nil
	case "sort-horizontal":
		return SortHorizontal(), nil
	case "sort-horizontal-line":
		return SortHorizontalLine(), nil
	case "sort-vertical":
		return SortVertical(), nil
	case "sort-vertical-line":
		return SortVerticalLine(), nil
	case "sound-off":
		return SoundOff(), nil
	case "sound-off-line":
		return SoundOffLine(), nil
	case "sound-up":
		return SoundUp(), nil
	case "sound-up-line":
		return SoundUpLine(), nil
	case "sparkles":
		return Sparkles(), nil
	case "sparkles-line":
		return SparklesLine(), nil
	case "speaker":
		return Speaker(), nil
	case "speaker-line":
		return SpeakerLine(), nil
	case "speakerphone":
		return Speakerphone(), nil
	case "speakerphone-line":
		return SpeakerphoneLine(), nil
	case "star":
		return Star(), nil
	case "star-line":
		return StarLine(), nil
	case "status-offline":
		return StatusOffline(), nil
	case "status-offline-line":
		return StatusOfflineLine(), nil
	case "status-online":
		return StatusOnline(), nil
	case "status-online-line":
		return StatusOnlineLine(), nil
	case "stop-circle":
		return StopCircle(), nil
	case "stop-circle-line":
		return StopCircleLine(), nil
	case "strike-through":
		return StrikeThrough(), nil
	case "strike-through-line":
		return StrikeThroughLine(), nil
	case "suitcase":
		return Suitcase(), nil
	case "suitcase-line":
		return SuitcaseLine(), nil
	case "suitcase-3":
		return SuitcaseThree(), nil
	case "suitcase-3-line":
		return SuitcaseThreeLine(), nil
	case "suitcase-2":
		return SuitcaseTwo(), nil
	case "suitcase-2-line":
		return SuitcaseTwoLine(), nil
	case "sun":
		return Sun(), nil
	case "sun-line":
		return SunLine(), nil
	case "support":
		return Support(), nil
	case "support-line":
		return SupportLine(), nil
	case "t-shirt":
		return TShirt(), nil
	case "t-shirt-line":
		return TShirtLine(), nil
	case "table":
		return Table(), nil
	case "table-heart":
		return TableHeart(), nil
	case "table-heart-line":
		return TableHeartLine(), nil
	case "table-line":
		return TableLine(), nil
	case "table-plus":
		return TablePlus(), nil
	case "table-plus-line":
		return TablePlusLine(), nil
	case "tag":
		return Tag(), nil
	case "tag-line":
		return TagLine(), nil
	case "tag-off":
		return TagOff(), nil
	case "tag-off-line":
		return TagOffLine(), nil
	case "telescope":
		return Telescope(), nil
	case "telescope-line":
		return TelescopeLine(), nil
	case "terminal":
		return Terminal(), nil
	case "terminal-line":
		return TerminalLine(), nil
	case "test-tube-filled":
		return TestTubeFilled(), nil
	case "test-tube-filled-line":
		return TestTubeFilledLine(), nil
	case "text":
		return Text(), nil
	case "text-align-center":
		return TextAlignCenter(), nil
	case "text-align-center-line":
		return TextAlignCenterLine(), nil
	case "text-align-justify":
		return TextAlignJustify(), nil
	case "text-align-justify-line":
		return TextAlignJustifyLine(), nil
	case "text-align-left":
		return TextAlignLeft(), nil
	case "text-align-left-line":
		return TextAlignLeftLine(), nil
	case "text-align-right":
		return TextAlignRight(), nil
	case "text-align-right-line":
		return TextAlignRightLine(), nil
	case "text-line":
		return TextLine(), nil
	case "text-wrap":
		return TextWrap(), nil
	case "text-wrap-line":
		return TextWrapLine(), nil
	case "textbox":
		return Textbox(), nil
	case "textbox-line":
		return TextboxLine(), nil
	case "textbox-minus":
		return TextboxMinus(), nil
	case "textbox-minus-line":
		return TextboxMinusLine(), nil
	case "textbox-plus":
		return TextboxPlus(), nil
	case "textbox-plus-line":
		return TextboxPlusLine(), nil
	case "thumb-down":
		return ThumbDown(), nil
	case "thumb-down-line":
		return ThumbDownLine(), nil
	case "thumb-up":
		return ThumbUp(), nil
	case "thumb-up-line":
		return ThumbUpLine(), nil
	case "ticket":
		return Ticket(), nil
	case "ticket-check":
		return TicketCheck(), nil
	case "ticket-check-line":
		return TicketCheckLine(), nil
	case "ticket-line":
		return TicketLine(), nil
	case "ticket-text":
		return TicketText(), nil
	case "ticket-text-line":
		return TicketTextLine(), nil
	case "tickets":
		return Tickets(), nil
	case "tickets-line":
		return TicketsLine(), nil
	case "timer":
		return Timer(), nil
	case "timer-line":
		return TimerLine(), nil
	case "tooltip":
		return Tooltip(), nil
	case "tooltip-line":
		return TooltipLine(), nil
	case "tooltip-text":
		return TooltipText(), nil
	case "tooltip-text-line":
		return TooltipTextLine(), nil
	case "tooltips":
		return Tooltips(), nil
	case "tooltips-line":
		return TooltipsLine(), nil
	case "tooltips-2":
		return TooltipsTwo(), nil
	case "tooltips-2-line":
		return TooltipsTwoLine(), nil
	case "translate":
		return Translate(), nil
	case "translate-line":
		return TranslateLine(), nil
	case "trash":
		return Trash(), nil
	case "trash-line":
		return TrashLine(), nil
	case "trending-down":
		return TrendingDown(), nil
	case "trending-down-line":
		return TrendingDownLine(), nil
	case "trending-up":
		return TrendingUp(), nil
	case "trending-up-line":
		return TrendingUpLine(), nil
	case "truck":
		return Truck(), nil
	case "truck-line":
		return TruckLine(), nil
	case "tv-old":
		return TvOld(), nil
	case "tv-old-line":
		return TvOldLine(), nil
	case "umbrella":
		return Umbrella(), nil
	case "umbrella-line":
		return UmbrellaLine(), nil
	case "underline":
		return Underline(), nil
	case "underline-line":
		return UnderlineLine(), nil
	case "underline-2":
		return UnderlineTwo(), nil
	case "underline-2-line":
		return UnderlineTwoLine(), nil
	case "undo":
		return Undo(), nil
	case "undo-line":
		return UndoLine(), nil
	case "unlock-open":
		return UnlockOpen(), nil
	case "unlock-open-line":
		return UnlockOpenLine(), nil
	case "usb":
		return Usb(), nil
	case "usb-line":
		return UsbLine(), nil
	case "user":
		return User(), nil
	case "user-add":
		return UserAdd(), nil
	case "user-add-line":
		return UserAddLine(), nil
	case "user-box":
		return UserBox(), nil
	case "user-box-line":
		return UserBoxLine(), nil
	case "user-circle":
		return UserCircle(), nil
	case "user-circle-line":
		return UserCircleLine(), nil
	case "user-group":
		return UserGroup(), nil
	case "user-group-line":
		return UserGroupLine(), nil
	case "user-line":
		return UserLine(), nil
	case "user-remove":
		return UserRemove(), nil
	case "user-remove-line":
		return UserRemoveLine(), nil
	case "users":
		return Users(), nil
	case "users-line":
		return UsersLine(), nil
	case "ux-circle":
		return UxCircle(), nil
	case "ux-circle-line":
		return UxCircleLine(), nil
	case "video":
		return Video(), nil
	case "video-camera":
		return VideoCamera(), nil
	case "video-camera-line":
		return VideoCameraLine(), nil
	case "video-line":
		return VideoLine(), nil
	case "video-minus":
		return VideoMinus(), nil
	case "video-minus-line":
		return VideoMinusLine(), nil
	case "video-plus":
		return VideoPlus(), nil
	case "video-plus-line":
		return VideoPlusLine(), nil
	case "view-boards":
		return ViewBoards(), nil
	case "view-boards-line":
		return ViewBoardsLine(), nil
	case "view-columns":
		return ViewColumns(), nil
	case "view-columns-line":
		return ViewColumnsLine(), nil
	case "view-grid":
		return ViewGrid(), nil
	case "view-grid-line":
		return ViewGridLine(), nil
	case "view-list":
		return ViewList(), nil
	case "view-list-line":
		return ViewListLine(), nil
	case "view-rows":
		return ViewRows(), nil
	case "view-rows-line":
		return ViewRowsLine(), nil
	case "watch":
		return Watch(), nil
	case "watch-line":
		return WatchLine(), nil
	case "wifi":
		return Wifi(), nil
	case "wifi-line":
		return WifiLine(), nil
	case "yen-circle":
		return YenCircle(), nil
	case "yen-circle-line":
		return YenCircleLine(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-in-line":
		return ZoomInLine(), nil
	case "zoom-out":
		return ZoomOut(), nil
	case "zoom-out-line":
		return ZoomOutLine(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in majesticons icon set", name)
	}
}
