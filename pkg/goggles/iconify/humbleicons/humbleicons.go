package humbleicons

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	activityInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h4l3-9l4 18l3-9h4"/>`
	adjustmentsInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v10m0 6v-.5M17.5 4v5m0 11v-5.56M6.5 4v2m0 14v-8.44M6.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm5.5 8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm5.5-5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`
	aidInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2h-2M8 8V6a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2M8 8h8m-4 4v2m0 0v2m0-2h-2m2 0h2"/>`
	alignObjectsBottomInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20H4m8-3V6m0 11l3-3m-3 3l-3-3"/>`
	alignObjectsCenterInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m4-8h5m-3-3l-3 3l3 3M8 12H3m3-3l3 3l-3 3"/>`
	alignObjectsLeftInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16m3-8h11M7 12l3-3m-3 3l3 3"/>`
	alignObjectsMiddleInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h16m-8-4V3M9 6l3 3l3-3m-3 10v5m-3-3l3-3l3 3"/>`
	alignObjectsRightInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 4v16m-3-8H6m11 0l-3-3m3 3l-3 3"/>`
	alignObjectsTopInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 4H4m8 3v11m0-11l3 3m-3-3l-3 3"/>`
	alignTextCenterInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 6h16M7 10h10M4 14h16M9 18h6"/>`
	alignTextJustifyInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"/>`
	alignTextLeftInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 6h16M4 10h13M4 14h16M4 18h5.5"/>`
	alignTextRightInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 6h16M7 10h13M4 14h16m-6 4h6"/>`
	archiveInnerSVG              = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M3 5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3H3V5z"/><path stroke-linecap="round" d="M9.5 13h5"/><path stroke-linejoin="round" d="M4 8h16v11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8z"/></g>`
	arrowDownInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 4v16m-6-6l6 6l6-6"/>`
	arrowGoBackInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 18h3.75a5.25 5.25 0 1 0 0-10.5H5M7.5 4L4 7.5L7.5 11"/>`
	arrowGoForwardInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 18H9.25a5.25 5.25 0 1 1 0-10.5H19M16.5 4L20 7.5L16.5 11"/>`
	arrowJoinInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 12l-3-3m3 3l-3 3m3-3h-5.929a5 5 0 0 0-3.535 1.464l-1.072 1.072A5 5 0 0 1 6.93 16H3m0-8h4.343a4 4 0 0 1 2.829 1.172l1.656 1.656A4 4 0 0 0 14.657 12H18"/>`
	arrowLeftInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4m0 0l6-6m-6 6l6 6"/>`
	arrowLeftDownInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M17.657 6.343L6.343 17.657m0-8.485v8.485h8.485"/>`
	arrowLeftUpInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 17.9L6.343 6.585m0 0v8.485m0-8.485h8.485"/>`
	arrowMainSplitSideInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 17H4.603M21 17l-3-3m3 3l-3 3M4.603 17H3m1.603 0a6 6 0 0 0 5.145-2.913l2.504-4.174A6 6 0 0 1 17.397 7H21m0 0l-3 3m3-3l-3-3"/>`
	arrowRightInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h16m0 0l-6 6m6-6l-6-6"/>`
	arrowRightDownInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.343 6.343l11.314 11.314m0 0H9.172m8.485 0V9.172"/>`
	arrowRightUpInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.343 17.657L17.657 6.343m0 0v8.485m0-8.485H9.172"/>`
	arrowSideJoinMainInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15h13.01m0 0a6 6 0 0 1-5.23-3.058l-1.06-1.884A6 6 0 0 0 4.49 7H3m13.01 8H21m0 0l-3 3m3-3l-3-3"/>`
	arrowSplitInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h4.597a5 5 0 0 1 3.904 1.877l.998 1.246A5 5 0 0 0 16.403 17H21m0 0l-3-3m3 3l-3 3m3-13h-5.078A4 4 0 0 0 12.8 8.501L11.201 10.5A4 4 0 0 1 8.078 12H6m15-5l-3-3m3 3l-3 3"/>`
	arrowUpInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20V4m0 0l6 6m-6-6l-6 6"/>`
	arrowsInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M12 4v6m0 4v6"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.5 6.5L12 4l2.5 2.5m-5 11L12 20l2.5-2.5m-8-8L4 12l2.5 2.5m11-5L20 12l-2.5 2.5M5.5 12h13"/></g>`
	arrowsHorizontalInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9l-3 3m0 0l3 3m-3-3h16m-3 3l3-3m0 0l-3-3"/>`
	arrowsVerticalInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 17l3 3m0 0l3-3m-3 3V4m3 3l-3-3m0 0L9 7"/>`
	asteriskSimpleInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 4l.005 7.993m0 0l7.603-2.465m-7.603 2.465l4.697 6.48m-4.697-6.48l-4.707 6.48m4.707-6.48L4.392 9.528"/>`
	atSymbolInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 20.064A9 9 0 1 1 21 12v1.5a2.5 2.5 0 0 1-5 0V8m0 4a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/>`
	banInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="m5.5 5.5l13 13M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`
	bandageInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v.001M14 12v.001M12 14v.001M10 12v.001M19 12l-7 7a4.95 4.95 0 1 1-7-7l7-7a4.95 4.95 0 0 1 7 7Z"/>`
	barsInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>`
	basketInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 10l1.36 8.164a1 1 0 0 0 .987.836h9.306a1 1 0 0 0 .986-.836L19 10M5 10h3m-3 0H4m15 0h-3m3 0h1M8 10l1-5m-1 5h8m0 0l-1-5m-6 9v1m6-1v1m-3-1v1"/>`
	batteryInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 10.5v3M6 17h10a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2Z"/>`
	batteryChargingInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 10.5v3M8.5 17H6a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h.5m9 10h.5a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-2.5M10 7l-1 5.5l4-1l-1 5.5"/>`
	batteryFullInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10.5v3M6 17h10a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2Z"/><path fill="currentColor" d="M15 10H7v4h8v-4Z"/></g>`
	batteryHalfInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10.5v3M6 17h10a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2Z"/><path fill="currentColor" d="M11 10H7v4h4v-4Z"/></g>`
	bellInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h2m0 0h10M7 17v-6a5 5 0 0 1 10 0v6m0 0h2M11.5 5.5V4a.5.5 0 0 1 1 0v1.5M13 20a1 1 0 1 1-2 0h2z"/>`
	bellOffInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l16 16M5 17h2m0 0h10M7 17v-6c0-1.126.372-2.164 1-3m3.5-2.5V4a.5.5 0 0 1 1 0v1.5M17 13v-2a5 5 0 0 0-6.5-4.771M13 20a1 1 0 1 1-2 0h2Z"/>`
	bikeInnerSVG                 = `<path fill="currentColor" d="m8.5 10.5l-.6-.8a1 1 0 0 0 .153 1.694L8.5 10.5Zm4 2l.99.141a1 1 0 0 0-.543-1.035l-.447.894Zm-1.49 3.359a1 1 0 1 0 1.98.282l-1.98-.282ZM12.5 7.5l.707-.707A1 1 0 0 0 11.9 6.7l.6.8ZM15 10l-.707.707A1 1 0 0 0 15 11v-1Zm2 1a1 1 0 1 0 0-2v2Zm-9 5a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4H8Zm-2 2a2 2 0 0 1-2-2H2a4 4 0 0 0 4 4v-2Zm-2-2a2 2 0 0 1 2-2v-2a4 4 0 0 0-4 4h2Zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4v2Zm2.053-2.606l4 2l.894-1.788l-4-2l-.894 1.788Zm3.457.965l-.5 3.5l1.98.282l.5-3.5l-1.98-.282ZM9.1 11.3l4-3l-1.2-1.6l-4 3l1.2 1.6Zm2.693-3.093l2.5 2.5l1.414-1.414l-2.5-2.5l-1.414 1.414ZM15 11h2V9h-2v2Zm5 5a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4h-2Zm-2 2a2 2 0 0 1-2-2h-2a4 4 0 0 0 4 4v-2Zm-2-2a2 2 0 0 1 2-2v-2a4 4 0 0 0-4 4h2Zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4v2Zm-3-9v2a2 2 0 0 0 2-2h-2Zm0 0h-2a2 2 0 0 0 2 2V5Zm0 0V3a2 2 0 0 0-2 2h2Zm0 0h2a2 2 0 0 0-2-2v2Z"/>`
	boldInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11.5V4h6.25A3.75 3.75 0 0 1 17 7.75v0a3.75 3.75 0 0 1-3.75 3.75H7Zm0 0V20h6.75A4.25 4.25 0 0 0 18 15.75v0a4.25 4.25 0 0 0-4.25-4.25H7Z"/>`
	bookInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M18 16V4H8a2 2 0 0 0-2 2v12"/><path d="M18 20H8a2 2 0 1 1 0-4h10c-.673 1.613-.66 2.488 0 4z"/></g>`
	bookOpenInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8.618V17.5m0-8.882a1 1 0 0 0-.553-.894l-.491-.246a7 7 0 0 0-4.12-.669l-.773.11A8 8 0 0 1 4.93 7H4v10h.38a8 8 0 0 0 2.197-.308l.553-.158a5 5 0 0 1 3.61.336l1.26.63m0-8.882a1 1 0 0 1 .553-.894l.491-.246a7 7 0 0 1 4.12-.669l.773.11c.375.054.753.081 1.131.081H20v10h-.38a8 8 0 0 1-2.197-.308l-.553-.158a5 5 0 0 0-3.61.336L12 17.5"/>`
	bookmarkInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6 20V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v15l-6-3l-6 3z"/>`
	boxInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 21l7.794-4.5v-9M12 21l-7.794-4.5v-9M12 21v-9m7.794-4.5L12 3L4.206 7.5m15.588 0L12 12M4.206 7.5L12 12"/>`
	brandGithubInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.5 21c2-2-.5-6 3.5-6m0 0c-3 0-7-1-7-5c0-1.445.116-2.89.963-4V3L9 4.283C9.821 4.101 10.81 4 12 4s2.178.1 3 .283L18 3v2.952c.88 1.116 1 2.582 1 4.048c0 4-4 5-7 5Zm0 0c4 0 1.5 4 3.5 6M3 15c3 0 1.5 4 6 3"/>`
	brandTwitterInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20.5 4.5l-2 1.5M21 7h-2M3.5 18c11.5 4.5 17-7 15-11s-7-1.5-5.5 3c-3.5 0-7-1-9-3.5c-1 3.5.5 8 4 9.5c-1.065 1.352-1.795 1.703-4.5 2Z"/>`
	briefcaseInnerSVG            = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8 8V6a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2M4 14l3.15.787a20 20 0 0 0 9.7 0L20 14"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 10a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8z"/><circle cx="12" cy="12" r="1" fill="currentColor"/></g>`
	brushBigInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 16l4.5 4.5l4-4a.707.707 0 0 1 1 0M6 16l-2.5-2.5l4-4a.707.707 0 0 0 0-1M6 16l2-2m7.5 2.5a.707.707 0 0 0 1 0l2.293-2.293a1 1 0 0 0 0-1.414l-2.086-2.086a1 1 0 0 1 0-1.414l3.586-3.586a1 1 0 0 0 0-1.414l-.586-.586a1 1 0 0 0-1.414 0l-3.586 3.586a1 1 0 0 1-1.414 0l-2.086-2.086a1 1 0 0 0-1.414 0L7.5 7.5a.707.707 0 0 0 0 1m8 8l-8-8"/>`
	bulbInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 18v-.107c0-.795-.496-1.488-1.117-1.984a5 5 0 1 1 6.235 0c-.622.497-1.118 1.189-1.118 1.984V18m-4 0v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2m-4 0h4m6-6h1M4 12H3m9-8V3m5.657 3.343l.707-.707m-12.02.707l-.708-.707M12 15v-2"/>`
	bulbOffInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 18v-.107c0-.795-.496-1.488-1.117-1.984a5 5 0 1 1 6.235 0c-.622.497-1.118 1.189-1.118 1.984V18m-4 0v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2m-4 0h4m-2-3v-2"/>`
	calendarInnerSVG             = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M4 6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v4H4V6z"/><path stroke-linecap="round" d="M8 6.5v-3m8 3v-3"/><path stroke-linejoin="round" d="M4 10h16v9a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9z"/></g>`
	cameraInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M20 17V8a1 1 0 0 0-1-1h-2.382a1 1 0 0 1-.894-.553l-.448-.894A1 1 0 0 0 14.382 5H9.618a1 1 0 0 0-.894.553l-.448.894A1 1 0 0 1 7.382 7H5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1Z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g>`
	cameraOffInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="m3 5l16 16m1-3V8a1 1 0 0 0-1-1h-2.382a1 1 0 0 1-.894-.553l-.448-.894A1 1 0 0 0 14.382 5H9.721a1 1 0 0 0-.949.684L8.5 6.5M16 18H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1v0"/><path d="M11 9.17A3 3 0 0 1 14.83 13"/></g>`
	cameraVideoInnerSVG          = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M16 16V8a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1z"/><path stroke-linejoin="round" d="m20 7l-4 3v4l4 3V7z"/></g>`
	cameraVideoOffInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 5l14 14M9 7h6a1 1 0 0 1 1 1v5.5a1 1 0 0 0 .4.8L20 17V7l-4 3m-1 7H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1"/>`
	carInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 11h2a2 2 0 0 1 2 2v2a1 1 0 0 1-1 1h-1.5M17 11h-6.5m6.5 0l-2.417-4.029A2 2 0 0 0 12.868 6H10.5m0 0v5m0-5H7.64a2 2 0 0 0-1.962 1.608L5 11m5.5 0H5m.5 5H4a1 1 0 0 1-1-1v-2a2 2 0 0 1 2-2v0m.5 5a2 2 0 1 0 4 0m-4 0a2 2 0 1 1 4 0m0 0h5m0 0a2 2 0 1 0 4 0m-4 0a2 2 0 1 1 4 0"/>`
	cartInnerSVG                 = `<g fill="none"><circle cx="7.5" cy="18.5" r="1.5" fill="currentColor"/><circle cx="16.5" cy="18.5" r="1.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5h2l.6 3m0 0L7 15h10l2-7H5.6z"/></g>`
	chartInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M17 5v15m-5-9v9m-5-6v6"/>`
	chatInnerSVG                 = `<g fill="currentColor"><path d="m4 19l-.93-.37a1 1 0 0 0 1.125 1.35L4 19zm4.706-.936l.474-.881l-.317-.17l-.352.07l.195.98zm-3.082-3.147l.93.37l.163-.414l-.196-.399l-.897.443zM19 12c0 3.246-2.853 6-6.53 6v2c4.641 0 8.53-3.514 8.53-8h-2zM5.941 12c0-3.246 2.854-6 6.53-6V4C7.83 4 3.94 7.514 3.94 12h2zm6.53-6C16.147 6 19 8.754 19 12h2c0-4.486-3.889-8-8.53-8v2zm0 12c-1.205 0-2.328-.3-3.291-.817l-.948 1.761A8.934 8.934 0 0 0 12.471 20v-2zm-8.276 1.98l4.706-.936l-.39-1.961l-4.706.936l.39 1.962zm2.326-5.506A5.564 5.564 0 0 1 5.94 12h-2c0 1.2.282 2.338.786 3.36l1.794-.886zm-1.826.073L3.07 18.631l1.858.738l1.624-4.083l-1.858-.739z"/><circle cx="9" cy="12" r="1"/><circle cx="12.5" cy="12" r="1"/><circle cx="16" cy="12" r="1"/></g>`
	chatsInnerSVG                = `<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9.882 15C13.261 15 16 12.538 16 9.5S13.261 4 9.882 4C6.504 4 3.765 6.462 3.765 9.5c0 .818.198 1.594.554 2.292L3 15l3.824-.736A6.62 6.62 0 0 0 9.882 15Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.804 18.124a6.593 6.593 0 0 0 3.314.876a6.623 6.623 0 0 0 3.059-.736L21 19l-1.32-3.208a5.02 5.02 0 0 0 .555-2.292c0-1.245-.46-2.393-1.235-3.315c-.749-.89-1.792-1.569-3-1.92"/><circle r="1" fill="currentColor" transform="matrix(-1 0 0 1 13 9.5)"/><circle r="1" fill="currentColor" transform="matrix(-1 0 0 1 10 9.5)"/><circle r="1" fill="currentColor" transform="matrix(-1 0 0 1 7 9.5)"/></g>`
	checkInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 14l4 4L19 8"/>`
	checkCircleInnerSVG          = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m8 13l2.5 2.5L16 10"/><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/></g>`
	chevronDownInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m5 10l7 7l7-7"/>`
	chevronLeftInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 5l-7 7l7 7"/>`
	chevronRightInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 5l7 7l-7 7"/>`
	chevronUpInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 14l7-7l7 7"/>`
	circleInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0z"/>`
	clipboardInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9 4H6a1 1 0 0 0-1 1v15a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-3M9 3h6v4H9V3z"/>`
	clockInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.5V12l3.5 2m5.5-2a9 9 0 1 1-18 0a9 9 0 0 1 18 0z"/>`
	cloudInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9.036a3.485 3.485 0 0 1 1.975.99M4 12.5A3.5 3.5 0 0 0 7.5 16h9.75a2.75 2.75 0 0 0 .734-5.4A5 5 0 0 0 8.37 9.108A3.5 3.5 0 0 0 4 12.5z"/>`
	cloudSunInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3.5V3m-6.5 7H3m11.596-4.596l.354-.354m-9.546.354L5.05 5.05M9 13.036a3.484 3.484 0 0 1 1.975.99M12.645 7a4 4 0 0 0-6.59 3.666M5 16.5A3.5 3.5 0 0 0 8.5 20h9.75a2.75 2.75 0 0 0 .734-5.4a5 5 0 0 0-9.614-1.491A3.504 3.504 0 0 0 5 16.5z"/>`
	codeInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m7 8l-4 4l4 4"/><path d="m10.5 18l3-12"/><path stroke-linejoin="round" d="m17 8l4 4l-4 4"/></g>`
	coffeeInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12H4v4a4 4 0 0 0 4 4h5a4 4 0 0 0 4-4v-4zm0 0h2a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-2m-4-8s1-1 .5-2l-1-2C12 4 13 3 13 3M8.64 9s1-1 .5-2l-1-2c-.5-1 .5-2 .5-2"/>`
	cogInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path d="M10.47 4.32c.602-1.306 2.458-1.306 3.06 0l.218.473a1.684 1.684 0 0 0 2.112.875l.49-.18c1.348-.498 2.66.814 2.162 2.163l-.18.489a1.684 1.684 0 0 0 .875 2.112l.474.218c1.305.602 1.305 2.458 0 3.06l-.474.218a1.684 1.684 0 0 0-.875 2.112l.18.49c.498 1.348-.814 2.66-2.163 2.162l-.489-.18a1.684 1.684 0 0 0-2.112.875l-.218.473c-.602 1.306-2.458 1.306-3.06 0l-.218-.473a1.684 1.684 0 0 0-2.112-.875l-.49.18c-1.348.498-2.66-.814-2.163-2.163l.181-.489a1.684 1.684 0 0 0-.875-2.112l-.474-.218c-1.305-.602-1.305-2.458 0-3.06l.474-.218a1.684 1.684 0 0 0 .875-2.112l-.18-.49c-.498-1.348.814-2.66 2.163-2.163l.489.181a1.684 1.684 0 0 0 2.112-.875l.218-.474Z"/></g>`
	coinsInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.5 9.09a5.502 5.502 0 0 1-1 10.91a5.5 5.5 0 0 1-4.9-3M7.5 8l1-.5v4m5.5-2a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Z"/>`
	columnsOneTwoThirdsInnerSVG  = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9 4v16m-5 0h16a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1z"/>`
	columnsThreeThirdsInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9 4v16m6-16v16M4 20h16a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1z"/>`
	columnsTwoHalfsInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 4v16m-8 0h16a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1z"/>`
	cornerDownLeftInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M20 5v6.5a3 3 0 0 1-3 3H5"/><path d="M7.5 18L4 14.5L7.5 11"/></g>`
	cornerDownRightInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5v6.5a3 3 0 0 0 3 3h12M16.5 18l3.5-3.5l-3.5-3.5"/>`
	cornerLeftDownInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M19 4h-6.5a3 3 0 0 0-3 3v12"/><path d="M6 16.5L9.5 20l3.5-3.5"/></g>`
	cornerLeftUpInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M19 20h-6.5a3 3 0 0 1-3-3V5"/><path d="M6 7.5L9.5 4L13 7.5"/></g>`
	cornerRightDownInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 4h6.5a3 3 0 0 1 3 3v12"/><path d="M18 16.5L14.5 20L11 16.5"/></g>`
	cornerRightUpInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 20h6.5a3 3 0 0 0 3-3V5"/><path d="M18 7.5L14.5 4L11 7.5"/></g>`
	cornerTopLeftInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 19v-6.5a3 3 0 0 0-3-3H5M7.5 6L4 9.5L7.5 13"/>`
	cornerUpRightInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19v-6.5a3 3 0 0 1 3-3h12M16.5 6L20 9.5L16.5 13"/>`
	cpuInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3v3m6-3v3m3 3h3m-3 6h3m-6 3v3m-6-3v3m-3-6H3m3-6H3m6.5 4.5v-3a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1ZM7 18h10a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1Z"/>`
	creativeCommonsInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10.333c-.313-.25-.703-.333-1.125-.333C7.839 10 7 10.895 7 12s.84 2 1.875 2c.422 0 .812-.082 1.125-.333m6-3.334c-.313-.25-.703-.333-1.125-.333c-1.036 0-1.875.895-1.875 2s.84 2 1.875 2c.422 0 .812-.082 1.125-.333M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`
	creativeCommonsByInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.5 10.5V18h-1v-7.5m1 0h-1m1 0H14v3m-2.5-3H10v3M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Zm-8-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`
	creativeCommonsNdInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 10.5h6m-6 3h6m6-1.5a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`
	creativeCommonsSaInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12a4 4 0 1 1 1.354 3M8 12l2.5-1M8 12l-1.5-2M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`
	creditCardInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2M4 10h16M4 10v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6M7 15h5"/>`
	cropInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7h12a1 1 0 0 1 1 1v12M7 10v6a1 1 0 0 0 1 1h6M7 4v3m10 10h3"/>`
	crownInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 18h14M5 14h14l1-9l-4 3l-4-5l-4 5l-4-3l1 9Z"/>`
	currencyDollarCircleInnerSVG = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.5 10l-.035-.139A2.457 2.457 0 0 0 12.082 8h-.522a1.841 1.841 0 0 0-.684 3.55l2.248.9A1.841 1.841 0 0 1 12.44 16h-.521a2.457 2.457 0 0 1-2.384-1.861L9.5 14M12 6v2m0 8v2m9-6a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`
	currencyEuroCircleInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12a4 4 0 0 0 6.4 3.2M8 12a4 4 0 0 1 6.4-3.2M8 12h3m-3 0H7m14 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`
	currencyPoundCircleInnerSVG  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 8.5l-.11-.164A3 3 0 0 0 12.395 7H12a3 3 0 0 0-3 3v3.757a3 3 0 0 1-.879 2.122L8 16h7m-7.5-4H12m9 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`
	dashboardInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5Zm10 0a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1V5ZM4 16a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-3Zm10-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-6Z"/>`
	databaseInnerSVG             = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M20 12c0 1.657-3.582 3-8 3s-8-1.343-8-3m16 6c0 1.657-3.582 3-8 3s-8-1.343-8-3"/><ellipse cx="12" cy="6" rx="8" ry="3"/><path d="M4 6v12M20 6v12"/></g>`
	desktopInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 20h8m-4 0v-4M4 5h16a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1z"/>`
	documentInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M5 20V4a1 1 0 0 1 1-1h6.172a2 2 0 0 1 1.414.586l4.828 4.828A2 2 0 0 1 19 9.828V20a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z"/><path d="M12 3v6a1 1 0 0 0 1 1h6"/></g>`
	documentAddInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M14.5 15.5h-5M12 18v-5"/><path d="M5 20V4a1 1 0 0 1 1-1h6.172a2 2 0 0 1 1.414.586l4.828 4.828A2 2 0 0 1 19 9.828V20a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z"/><path d="M12 3v6a1 1 0 0 0 1 1h6"/></g>`
	documentRemoveInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M14.5 15.5h-5"/><path d="M5 20V4a1 1 0 0 1 1-1h6.172a2 2 0 0 1 1.414.586l4.828 4.828A2 2 0 0 1 19 9.828V20a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z"/><path d="M12 3v6a1 1 0 0 0 1 1h6"/></g>`
	documentsInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v7h7M6 7H5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-1M8 4v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V9.389a1 1 0 0 0-.263-.676l-4.94-5.389A1 1 0 0 0 14.06 3H9a1 1 0 0 0-1 1Z"/>`
	dotsHorizontalInnerSVG       = `<g fill="currentColor"><rect width="4" height="4" x="3" y="10" rx="2"/><rect width="4" height="4" x="10" y="10" rx="2"/><rect width="4" height="4" x="17" y="10" rx="2"/></g>`
	dotsVerticalInnerSVG         = `<g fill="currentColor"><rect width="4" height="4" x="10" y="3" rx="2"/><rect width="4" height="4" x="10" y="10" rx="2"/><rect width="4" height="4" x="10" y="17" rx="2"/></g>`
	downloadInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11.5V20m0 0l3-3m-3 3l-3-3M8 7.036a3.484 3.484 0 0 1 1.975.99M17.5 14c1.519 0 2.5-1.231 2.5-2.75a2.75 2.75 0 0 0-2.016-2.65A5 5 0 0 0 8.37 7.108a3.5 3.5 0 0 0-1.87 6.746"/>`
	downloadAltInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v8.5m0 0l3-3m-3 3l-3-3M5 15v2a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-2"/>`
	dropletInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14v.5M12 4c-1.262 1.683-3.055 3.896-4.708 5.896c-2.288 2.767-1.796 6.907 1.115 9.009v0a6.137 6.137 0 0 0 7.186 0v0c2.91-2.102 3.403-6.242 1.116-9.009C15.055 7.896 13.262 5.683 12 4Z"/>`
	duplicateInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v1M9 20h10a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1z"/>`
	exchangeHorizontalInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 17h12M4 17l3.5-3.5M4 17l3.5 3.5M7 7h13m0 0l-3.5-3.5M20 7l-3.5 3.5"/>`
	exchangeVerticalInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20V8m0 12l-3.5-3.5M17 20l3.5-3.5M7 17V4m0 0L3.5 7.5M7 4l3.5 3.5"/>`
	exclamationInnerSVG          = `<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M13.253 5.98L12 13.5l-1.253-7.52a1.27 1.27 0 1 1 2.506 0Z"/><circle cx="12" cy="19" r="1" stroke-width="2"/></g>`
	exclamationTriangleInnerSVG  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v4m0 3v.01M5.313 20h13.374c1.505 0 2.471-1.6 1.77-2.931L13.77 4.363c-.75-1.425-2.79-1.425-3.54 0L3.543 17.068C2.842 18.4 3.808 20 5.313 20Z"/>`
	expandInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8.5V4m0 0h4.5M4 4l5.5 5.5m10.5-1V4m0 0h-4.5M20 4l-5.5 5.5M4 15.5V20m0 0h4.5M4 20l5.5-5.5m10.5 1V20m0 0h-4.5m4.5 0l-5.5-5.5"/>`
	externalLinkInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-5m-6 0l7.5-7.5M15 3h6v6"/>`
	eyeInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M3 12c5.4-8 12.6-8 18 0c-5.4 8-12.6 8-18 0z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/></g>`
	eyeCloseInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10a13.358 13.358 0 0 0 3 2.685M21 10a13.358 13.358 0 0 1-3 2.685m-8 1.624L9.5 16.5m.5-2.19a10.59 10.59 0 0 0 4 0m-4 0a11.275 11.275 0 0 1-4-1.625m8 1.624l.5 2.191m-.5-2.19a11.275 11.275 0 0 0 4-1.625m0 0l1.5 1.815M6 12.685L4.5 14.5"/>`
	eyeOffInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 5l16 16M11.148 9.123a3 3 0 0 1 3.722 3.752M8.41 6.878C12.674 4.762 17.267 6.47 21 12c-1.027 1.521-2.119 2.753-3.251 3.696m-2.509 1.59C11.076 19.142 6.631 17.38 3 12c1.01-1.496 2.083-2.713 3.196-3.65"/>`
	fastForwardInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15.196V8.804a1 1 0 0 1 1.53-.848l5.113 3.196a1 1 0 0 1 0 1.696L6.53 16.044A1 1 0 0 1 5 15.196Zm8 0V8.804a1 1 0 0 1 1.53-.848l5.113 3.196a1 1 0 0 1 0 1.696l-5.113 3.196a1 1 0 0 1-1.53-.848Z"/>`
	fingerprintInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 21l1.726-2.761a8 8 0 0 0 1.06-2.671l.059-.291a8.002 8.002 0 0 0 .155-1.57V12m-9 3.5c.5-1 1-2.043 1-3.5c0-1.74.556-3.352 1.5-4.665m0 11.165c2-2 2.5-5.237 2.5-6.5a4 4 0 0 1 7.954-.61M19.5 17.5c.5-2 .5-4 .5-5.5A8 8 0 0 0 8 5.07M15.954 15c-.174 1.393-.666 3.181-1.954 5.5"/>`
	flagInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M6 5v15"/><path stroke-linejoin="round" d="M13 5.5V4H7a1 1 0 0 0-1 1v8h7v1.5h7v-9h-7zm0 0v3"/></g>`
	flashInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 3l-3 8l6.5-1L12 21m0 0l3.5-2.5M12 21l-1.5-4"/>`
	flaskInnerSVG                = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M9.5 10V4h5v6l4.743 5.174A2.88 2.88 0 0 1 17.12 20H6.88a2.88 2.88 0 0 1-2.123-4.826L9.5 10z"/><path stroke-linecap="round" d="M8.5 4h7"/><path d="M6 14c3.5 2 4 0 6 0s2.5 2 6 0"/></g>`
	folderInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 18V6a2 2 0 0 1 2-2h4.539a2 2 0 0 1 1.562.75L12.2 6.126a1 1 0 0 0 .78.375H20a1 1 0 0 1 1 1V18a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>`
	folderAddInnerSVG            = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M3 18V6a2 2 0 0 1 2-2h4.539a2 2 0 0 1 1.562.75L12.2 6.126a1 1 0 0 0 .78.375H20a1 1 0 0 1 1 1V18a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/><path stroke-linecap="round" d="M9.5 12.5h5M12 15v-5"/></g>`
	folderOpenInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M17 8V7a1 1 0 0 0-1-1h-5.586a1 1 0 0 1-.707-.293L8.293 4.293A1 1 0 0 0 7.586 4H4a1 1 0 0 0-1 1v12a2 2 0 0 0 2 2h14a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v6a2 2 0 0 1-2 2"/>`
	folderRemoveInnerSVG         = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M3 18V6a2 2 0 0 1 2-2h4.539a2 2 0 0 1 1.562.75L12.2 6.126a1 1 0 0 0 .78.375H20a1 1 0 0 1 1 1V18a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/><path stroke-linecap="round" d="M9.5 12.5h5"/></g>`
	forkKnifeInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4v2a3 3 0 0 0 3 3m0 0V4m0 5v11M8 9a3 3 0 0 0 3-3V4m5 8V4c3 2 3 4 3 8h-3Zm0 0v8"/>`
	funnelInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M20 4H4v2l6 6v8.5l4-2.5v-6l6-6V4z"/>`
	giftInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 9V6a3 3 0 1 0-3 3h3zm0 0V7a2 2 0 1 1 2 2h-2zm-7 4v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7m-7-3v11m8-8v-3a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v3h16z"/>`
	globeInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 0 1-9 9m9-9a9 9 0 0 0-9-9m9 9H3m9 9a9 9 0 0 1-9-9m9 9c1.933 0 3.5-4.03 3.5-9S13.933 3 12 3m0 18c-1.933 0-3.5-4.03-3.5-9s1.567-9 3.5-9m-9 9a9 9 0 0 1 9-9"/>`
	headingInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4v8m0 8v-8m10-8v8m0 8v-8m0 0H7M5 4h4m6 0h4m0 16h-4m-6 0H5"/>`
	heartInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.574-1.635-4.46-2.135-6.035-.5c-1.573 1.635-1.34 3.836 0 5.752C7.306 15.168 9.41 16.89 12 19c2.59-2.11 4.694-3.832 6.035-5.748c1.34-1.916 1.573-4.117 0-5.752C16.46 5.865 13.574 6.365 12 8Z"/>`
	homeInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10v9a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-9M6 10l6-6l6 6M6 10l-2 2m14-2l2 2m-10 1h4v4h-4v-4z"/>`
	humbleiconInnerSVG           = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M8 4a2 2 0 0 1 2 2v11a2 2 0 1 1-4 0V6a2 2 0 0 1 2-2zm8 8a2 2 0 0 1 2 2v3a2 2 0 1 1-4 0v-3a2 2 0 0 1 2-2z"/><circle cx="16" cy="6" r="2"/></g>`
	imageInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 17l-3.293-3.293a1 1 0 0 0-1.414 0l-.586.586a1 1 0 0 1-1.414 0l-2.879-2.879a2 2 0 0 0-2.828 0L3 17M21 5v14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-5 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`
	imagesInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19v-9a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm0 0l4.293-4.293a1 1 0 0 1 1.414 0L14 20M7 6V5a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-1m-7-4v.01"/>`
	incognitoInnerSVG            = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M10 15.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm9 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m10 15l.211-.106a4 4 0 0 1 3.578 0L14 15m3-6l-1.65-4.125a1 1 0 0 0-1.245-.577l-1.473.491a2 2 0 0 1-1.264 0L9.894 4.3a1 1 0 0 0-1.245.576L7 9m-4 1c7.5-1 11-1 18 0"/></g>`
	incognitoTwoInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 7.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 0l.211-.106a4 4 0 0 1 3.578 0L14 7.5m0 0a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Zm-2 6.303c5-3 5 3.5 9 1.767c-1 4.233-6 4.233-9 1.233c-3 3-8 3-9-1.233c4 1.733 4-4.767 9-1.767Z"/>`
	infoCircleInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 11h1v5.5m0 0h1.5m-1.5 0h-1.5M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Zm-9.5-4v-.5h.5V8h-.5Z"/>`
	italicInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 20l4-16m2 0h-4m0 16H8"/>`
	keyInnerSVG                  = `<g fill="none"><circle cx="15.5" cy="8.5" r="1.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 16l5.18-5.652C10.033 9.875 10 9.523 10 9a5 5 0 1 1 5 5c-.523 0-.868-.01-1.342-.158L12 15.5h-2v2H8v2H5V16Z"/></g>`
	layersInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="m4 8l8-4l8 4l-8 4l-8-4z"/><path stroke-linecap="round" d="m4 12l8 4l8-4M4 16l8 4l8-4"/></g>`
	linkInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="m12 17l-1.5 1.5a3.536 3.536 0 0 1-5 0v0a3.536 3.536 0 0 1 0-5l3-3a3.536 3.536 0 0 1 5 0v0"/><path d="m12 7l1.5-1.5a3.536 3.536 0 0 1 5 0v0a3.536 3.536 0 0 1 0 5l-3 3a3.536 3.536 0 0 1-5 0v0"/></g>`
	locationInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M13 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M17.5 9.5c0 3.038-2 6.5-5.5 10.5c-3.5-4-5.5-7.462-5.5-10.5a5.5 5.5 0 1 1 11 0Z"/></g>`
	lockInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14v2m-4-6V8a4 4 0 1 1 8 0v2m-9 9h10a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1z"/>`
	lockOpenInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-4-6V7a4 4 0 1 1 8 0M7 20h10a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1z"/>`
	logoutInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12h-9.5m7.5 3l3-3l-3-3m-5-2V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-1"/>`
	mailInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M20 6H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1z"/><path stroke-linecap="round" d="m3 8l7.89 5.26a2 2 0 0 0 2.22 0L21 8"/></g>`
	mailOpenInnerSVG             = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M21 19V8.588a1 1 0 0 0-.514-.874l-7.03-3.905a3 3 0 0 0-2.913 0L3.514 7.714A1 1 0 0 0 3 8.588V19a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1z"/><path stroke-linecap="round" d="m3 10l7.89 5.26a2 2 0 0 0 2.22 0L21 10M9 14l-6 4m12-4l6 4"/></g>`
	mapInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.5 7v9.5m-7 .5V8.5M3.521 6.716l4.553-2.483a1 1 0 0 1 .872-.042l6.108 2.618a1 1 0 0 0 .872-.042l3.595-1.96A1 1 0 0 1 21 5.685v10.721a1 1 0 0 1-.521.878l-4.553 2.483a1 1 0 0 1-.872.042L8.946 17.19a1 1 0 0 0-.872.042l-3.595 1.96A1 1 0 0 1 3 18.315V7.595a1 1 0 0 1 .521-.878z"/>`
	maximizeInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8V6a2 2 0 0 1 2-2h2m8 0h2a2 2 0 0 1 2 2v2m0 8v2a2 2 0 0 1-2 2h-2m-8 0H6a2 2 0 0 1-2-2v-2"/>`
	microphoneInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 10.5A5.5 5.5 0 0 1 12 16m0 0a5.5 5.5 0 0 1-5.5-5.5M12 16v4m-4 0h4m0 0h4m-4-7a2.5 2.5 0 0 1-2.5-2.5v-4a2.5 2.5 0 0 1 5 0v4A2.5 2.5 0 0 1 12 13Z"/>`
	microphoneOffInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l16 16M6.5 10.5A5.5 5.5 0 0 0 12 16v4m-4 0h4m0 0h4M9.5 9V6.5a2.5 2.5 0 0 1 5 0v4a2.5 2.5 0 0 1-1.5 2.292m4.5-2.292c0 1.93-.803 3.523-2.309 4.504"/>`
	minusInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M18 12H6"/>`
	minusCircleInnerSVG          = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M16 12H8"/><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/></g>`
	mobileInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 17h4m3-12v14a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1z"/>`
	moneyInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8H3v10h18V8h-1.5m-4-1a3 3 0 1 1-4 4.258M13 6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/>`
	moonInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M20 14.12A7.78 7.78 0 0 1 9.88 4a7.782 7.782 0 0 0 2.9 15A7.782 7.782 0 0 0 20 14.12z"/>`
	moustacheInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12.57c-4 1.733-4-4.767-9-1.767c-5-3-5 3.5-9 1.767c1 4.233 6 4.233 9 1.233c3 3 8 3 9-1.233Z"/>`
	musicNoteInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9 18c0 1.105-1.12 2-2.5 2S4 19.105 4 18s1.12-2 2.5-2s2.5.895 2.5 2zm0 0V7l11-3v11m0 0c0 1.105-1.12 2-2.5 2s-2.5-.895-2.5-2s1.12-2 2.5-2s2.5.895 2.5 2z"/>`
	navigationInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 11l14-6l-6 14l-2-6l-6-2Z"/>`
	packageInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 8l8 4M4 8v8l8 4M4 8l4-2m4 6l4-2m-4 2v8m8-12l-8-4l-4 2m12 2v8l-8 4m8-12l-4 2m0 0L8 6"/>`
	pauseInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6 6h4v12H6V6Zm8 0h4v12h-4V6Z"/>`
	pencilInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.5 7.5l3 3M4 20v-3.5L15.293 5.207a1 1 0 0 1 1.414 0l2.086 2.086a1 1 0 0 1 0 1.414L7.5 20H4z"/>`
	phoneInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.515 4.621L9 4l2 3.5L9.5 9c1 2 3.5 4.5 5.5 5.5l1.5-1.5l3.5 2l-.621 2.485c-.223.89-1.029 1.534-1.928 1.351c-5.213-1.06-11.228-7.074-12.287-12.287c-.183-.9.46-1.705 1.35-1.928Z"/>`
	phoneCallInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.636 4A7.364 7.364 0 0 1 20 11.364M13 8a3 3 0 0 1 3 3M7 6l-2.485.621c-.89.223-1.534 1.029-1.352 1.928c1.06 5.213 7.075 11.228 12.288 12.287c.9.183 1.705-.46 1.928-1.35l.62-2.486l-3.5-2l-1.5 1.5c-2-1-4.5-3.5-5.5-5.5L9 9.5L7 6Z"/>`
	phoneForwardInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.36 6.75h7.779m0 0l-2.828 2.828m2.828-2.828l-2.828-2.828M8 5l-2.485.621c-.89.223-1.534 1.029-1.352 1.928c1.06 5.213 7.075 11.228 12.288 12.287c.9.183 1.705-.46 1.928-1.35l.62-2.486l-3.5-2l-1.5 1.5c-2-1-4.5-3.5-5.5-5.5L10 8.5L8 5Z"/>`
	phoneIncomingInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 4l-5.5 5.5m0 0v-4m0 4h4M8 5l-2.485.621c-.89.223-1.534 1.029-1.352 1.928c1.06 5.213 7.075 11.228 12.288 12.287c.9.183 1.705-.46 1.928-1.35l.62-2.486l-3.5-2l-1.5 1.5c-2-1-4.5-3.5-5.5-5.5L10 8.5L8 5Z"/>`
	phoneMissedInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 5l2 2m0 0l2 2m-2-2l2-2m-2 2l-2 2M8 5l-2.485.621c-.89.223-1.534 1.029-1.352 1.928c1.06 5.213 7.075 11.228 12.288 12.287c.9.183 1.705-.46 1.928-1.35l.62-2.486l-3.5-2l-1.5 1.5c-2-1-4.5-3.5-5.5-5.5L10 8.5L8 5Z"/>`
	phoneOffInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 13.906c.34.232.677.432 1 .594l1.5-1.5l3.5 2l-.621 2.485c-.223.89-1.03 1.534-1.93 1.351c-1.8-.366-3.695-1.323-5.449-2.634m-2.265-1.967C7.455 11.95 5.693 9.15 5.164 6.55c-.183-.9.46-1.706 1.35-1.929L9 4l2 3.5L9.5 9c.48.959 1.303 2.032 2.251 3M5 19L19 5"/>`
	phoneOutgoingInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 9.5L20 4m0 0v4m0-4h-4M8 5l-2.485.621c-.89.223-1.534 1.029-1.352 1.928c1.06 5.213 7.075 11.228 12.288 12.287c.9.183 1.705-.46 1.928-1.35l.62-2.486l-3.5-2l-1.5 1.5c-2-1-4.5-3.5-5.5-5.5L10 8.5L8 5Z"/>`
	pieChartInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 15A9.004 9.004 0 0 1 12 21A9 9 0 0 1 9 3.512M12 3a9 9 0 0 1 9 9h-9V3Z"/>`
	playInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 17.259V6.741a1 1 0 0 1 1.504-.864l9.015 5.26a1 1 0 0 1 0 1.727l-9.015 5.259A1 1 0 0 1 7 17.259Z"/>`
	plusInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 19V5m7 7H5"/>`
	plusCircleInnerSVG           = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M12 16V8m4 4H8"/><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/></g>`
	powerInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.5 7.638a7 7 0 1 0 9 0M12 4v7"/>`
	printInnerSVG                = `<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M8 17H5a1 1 0 0 1-1-1v-5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v5a1 1 0 0 1-1 1h-3M8 4h8v5H8V4zm0 11h8v4H8v-4z"/><circle cx="7" cy="12" r="1" fill="currentColor"/></g>`
	radioInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 18.364a9 9 0 0 1 0-12.728m12.728 0a9 9 0 0 1 0 12.728M8.111 15.889a5.5 5.5 0 0 1 0-7.778m7.778 0a5.5 5.5 0 0 1 0 7.778M14 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`
	rainInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.857 13L12 19m-2.5-7l-.857 6m7.857-6l-.857 6M8 7.036a3.484 3.484 0 0 1 1.975.99M6 13.662A3.5 3.5 0 0 1 8.37 7.11a5.002 5.002 0 0 1 9.614 1.49a2.751 2.751 0 0 1 1.59 4.122"/>`
	refreshInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20v-5h-5M4 4v5h5m10.938 2A8.001 8.001 0 0 0 5.07 8m-1.008 5a8.001 8.001 0 0 0 14.868 3"/>`
	rewindInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.113 15.196V8.804a1 1 0 0 0-1.53-.848l-5.113 3.196a1 1 0 0 0 0 1.696l5.113 3.196a1 1 0 0 0 1.53-.848Zm-8 0V8.804a1 1 0 0 0-1.53-.848L4.47 11.152a1 1 0 0 0 0 1.696l5.113 3.196a1 1 0 0 0 1.53-.848Z"/>`
	rocketInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 14l2.045-1.533C19.469 10.648 20.542 6.98 20 4c-2.981-.542-6.649.531-8.467 2.955L10 9m5 5l-3.5 2.5l-4-4L10 9m5 5v2.667a4 4 0 0 1-.8 2.4l-.7.933l-1-1M10 9H7.333a4 4 0 0 0-2.4.8L4 10.5l1 1M8.5 18L5 19l1.166-3.5m9.334-6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	rssInnerSVG                  = `<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M13 19a8 8 0 0 0-8-8m14 8c0-7.732-6.268-14-14-14"/><circle cx="6" cy="18" r="2" fill="currentColor"/></g>`
	saveInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.828a2 2 0 0 0-.586-1.414l-1.828-1.828A2 2 0 0 0 16.172 4H15M8 4v4a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V4M8 4h7M7 17v-3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3"/>`
	scissorsInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6c-3.573 2.225-5.943 3.854-8.55 6M20 18c-2.626-1.636-4.602-2.949-6.5-4.382M8.598 9.54a3 3 0 1 0-3.196-5.08a3 3 0 0 0 3.196 5.08Zm0 0A89.3 89.3 0 0 0 11.45 12m-2.852 2.46a3 3 0 1 0-3.196 5.079a3 3 0 0 0 3.196-5.078Zm0 0A89.287 89.287 0 0 1 11.45 12"/>`
	searchInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="m20 20l-6-6"/><path d="M15 9.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Z"/></g>`
	shareInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v8.5M15 7l-3-3l-3 3m-4 5v5a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-5"/>`
	shareAltInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.437 7.912a2.5 2.5 0 1 0 4.127-2.824a2.5 2.5 0 0 0-4.127 2.824Zm0 0l-4.96 3.394m0 0a3 3 0 1 0 .236 2.979m-.237-2.98c.33.483.524 1.066.524 1.695c0 .46-.103.895-.288 1.285m0 0l4.528 2.145m0 0a2.5 2.5 0 1 0 4.52 2.141a2.5 2.5 0 0 0-4.52-2.142Z"/>`
	shipInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 17l.756.378a3 3 0 0 0 2.523.074l1.04-.446a3 3 0 0 1 2.363 0l1.04.446a3 3 0 0 0 2.523-.074l.413-.207a3 3 0 0 1 2.684 0l.547.273a3 3 0 0 0 2.29.163L21 17M5 14.5L4 10h4m10 4.5l2.5-4.5h-8.245H13.5m0 0l-.721-3H8v3m5.5 0H8m3 3h.1M10 4.5l-.2-.2a2 2 0 0 0-1.899-.525l-.336.084a2 2 0 0 1-1.118-.043L5.5 3.5"/>`
	skipBackwardInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5.5v13m3.48-5.636l9.016 5.259A1 1 0 0 0 19 17.259V6.741a1 1 0 0 0-1.504-.864l-9.015 5.26a1 1 0 0 0 0 1.727Z"/>`
	skipForwardInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 5.5v13m-3.48-5.636l-9.016 5.259A1 1 0 0 1 5 17.259V6.741a1 1 0 0 1 1.504-.864l9.015 5.26a1 1 0 0 1 0 1.727Z"/>`
	snowInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 13h.01M16 13h.01M13 14.5h.01M10 16h.01M16 16h.01M13 17.5h.01M8 7.036a3.484 3.484 0 0 1 1.975.99M6 13.662A3.5 3.5 0 0 1 8.37 7.11a5.002 5.002 0 0 1 9.614 1.49a2.751 2.751 0 0 1 1.59 4.122"/>`
	squareInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5z"/>`
	starInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 2l2.939 5.955l6.572.955l-4.756 4.635l1.123 6.545L12 17l-5.878 3.09l1.123-6.545L2.489 8.91l6.572-.955L12 2Z"/>`
	stormInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.5 10l-3 5h4l-3 5.5M8 8.036a3.484 3.484 0 0 1 1.975.99M7.5 15a3.5 3.5 0 1 1 .87-6.891a5.002 5.002 0 0 1 9.614 1.49A2.751 2.751 0 0 1 17.25 15"/>`
	sunInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 4.5V3m0 18v-1.5m9-7.5h-1.5m-15 0H3m14.303-5.303l1.061-1.061M5.636 18.364l1.06-1.06m11.668 1.06l-1.06-1.06M6.696 6.696l-1.06-1.06M16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0z"/>`
	supportInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="m18 18l-3-3M9 9L6 6m9 3l3-3M6 18l3-3m12-3a9 9 0 1 1-18 0a9 9 0 0 1 18 0zm-5 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0z"/>`
	switchOffInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M12 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path d="M3 12a6 6 0 0 0 6 6h6a6 6 0 0 0 0-12H9a6 6 0 0 0-6 6Z"/></g>`
	switchOnInnerSVG             = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15 6H9a6 6 0 1 0 0 12h6a6 6 0 0 0 0-12Z"/><circle cx="15" cy="12" r="3" fill="currentColor"/></g>`
	tagInnerSVG                  = `<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 10.586V5a1 1 0 0 1 1-1h5.586a1 1 0 0 1 .707.293l8.5 8.5a1 1 0 0 1 0 1.414l-5.586 5.586a1 1 0 0 1-1.414 0l-8.5-8.5A1 1 0 0 1 4 10.586z"/><circle cx="8.5" cy="8.5" r="1.5" fill="currentColor"/></g>`
	tagsInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 10.95V6a1 1 0 0 1 1-1h4.95a1 1 0 0 1 .707.293l7.636 7.636a1 1 0 0 1 0 1.415l-4.95 4.949a1 1 0 0 1-1.414 0l-7.636-7.636A1 1 0 0 1 3 10.948z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m15.636 20l5.657-5.656a1 1 0 0 0 0-1.415L13.363 5"/><circle cx="6.5" cy="8.5" r="1.5" fill="currentColor"/></g>`
	textInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20V4m7 2V4H5v2m9 14h-4"/>`
	timesInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M6 18L18 6m0 12L6 6"/>`
	timesCircleInnerSVG          = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/><path stroke-linecap="round" d="m9 15l6-6m0 6L9 9"/></g>`
	trashInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 6l.934 13.071A1 1 0 0 0 7.93 20h8.138a1 1 0 0 0 .997-.929L18 6m-6 5v4m8-9H4m4.5 0l.544-1.632A2 2 0 0 1 10.941 3h2.117a2 2 0 0 1 1.898 1.368L15.5 6"/>`
	trendingDownInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.5 17L13 9.5l-4 4l-6-6m13 10h4.95v-5"/>`
	trendingUpInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.5 7.5L13 15l-4-4l-6 6M16 7h4.95v5"/>`
	triangleInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M11.125 2.584a1 1 0 0 1 1.75 0l8.805 15.932A1 1 0 0 1 20.805 20H3.195a1 1 0 0 1-.875-1.484l8.805-15.932z"/>`
	truckInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 11.5V16a1 1 0 0 1-1 1h-1.5m2.5-5.5h-7m7 0l-1.736-3.906A1 1 0 0 0 18.35 7H14M5.5 17H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v1M5.5 17a2 2 0 1 0 4 0m-4 0a2 2 0 1 1 4 0m0 0H14m0 0h.5m-.5 0v-5.5m.5 5.5a2 2 0 1 0 4 0m-4 0a2 2 0 1 1 4 0M14 11.5V7"/>`
	underlineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4v7a5 5 0 0 0 5 5v0a5 5 0 0 0 5-5V4M7 20h10"/>`
	uploadInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v9m0-9l3 3m-3-3l-3 3m8.5 2c1.519 0 2.5-1.231 2.5-2.75a2.75 2.75 0 0 0-2.016-2.65A5 5 0 0 0 8.37 8.108a3.5 3.5 0 0 0-1.87 6.746"/>`
	userInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 19v-1.25c0-2.071-1.919-3.75-4.286-3.75h-3.428C7.919 14 6 15.679 6 17.75V19m9-11a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/>`
	userAddInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11h3m0 0h3m-3 0v3m0-3V8m-3 11v-1.25c0-2.071-1.919-3.75-4.286-3.75H7.286C4.919 14 3 15.679 3 17.75V19m9-11a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/>`
	userAskingInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14H9.286C6.919 14 5 15.679 5 17.75V19M19 7v5a2 2 0 0 1-2 2h-2v5M14 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/>`
	userRemoveInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11h6m-6 8v-1.25c0-2.071-1.919-3.75-4.286-3.75H8.286C5.919 14 4 15.679 4 17.75V19m9-11a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/>`
	usersInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 19v-1.25C13 15.679 11.081 14 8.714 14H7.286C4.919 14 3 15.679 3 17.75V19m12.286-5h1.428C19.081 14 21 15.679 21 17.75V19M15 5.17a3 3 0 1 1 0 5.659M11 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/>`
	verifiedInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.5 12.5L11 15l4.5-4.5m-.595-5.512l-.48-.659a3 3 0 0 0-4.85 0l-.48.659l-.804-.127a3 3 0 0 0-3.43 3.43l.127.804l-.659.48a3 3 0 0 0 0 4.85l.659.48l-.127.804a3 3 0 0 0 3.43 3.43l.804-.127l.48.659a3 3 0 0 0 4.85 0l.48-.659l.804.127a3 3 0 0 0 3.43-3.43l-.127-.804l.659-.48a3 3 0 0 0 0-4.85l-.659-.48l.127-.804a3 3 0 0 0-3.43-3.43l-.804.127z"/>`
	viewGridInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 10.5V4h6v6.5H4Zm10 0V4h6v6.5h-6Zm-10 10V14h6v6.5H4Zm10 0V14h6v6.5h-6Z"/>`
	viewListInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"/>`
	volumeInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M10 9H7v6h3l5 4V5l-5 4z"/>`
	volumeOffInnerSVG            = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M7 9H4v6h3l5 4V5L7 9z"/><path stroke-linecap="round" d="m16 9.5l5 5m0-5l-5 5"/></g>`
	volumeOneInnerSVG            = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M8 9H5v6h3l5 4V5L8 9z"/><path stroke-linecap="round" d="M17 8a5.657 5.657 0 0 1 0 8"/></g>`
	volumeTwoInnerSVG            = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M6 9H3v6h3l5 4V5L6 9z"/><path stroke-linecap="round" d="M18.5 5.5a9.192 9.192 0 0 1 0 13M15 8a5.657 5.657 0 0 1 0 8"/></g>`
	wifiInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 8.437c-5.115-4.583-12.885-4.583-18 0m15.333 2.982a9.501 9.501 0 0 0-12.666 0m10 2.981a5.5 5.5 0 0 0-7.334 0M12 18.5l1-1.118a1.5 1.5 0 0 0-2 0l1 1.118Z"/>`
	wifiOffInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l16 16M9 5.336c4.143-.94 8.643.094 12 3.101m-8 .615a9.456 9.456 0 0 1 5.333 2.367m-10 2.981a5.491 5.491 0 0 1 4.345-1.358M3 8.437a13.445 13.445 0 0 1 3.206-2.134m-.54 5.116A9.446 9.446 0 0 1 9 9.484m3 9.016l1-1.118a1.5 1.5 0 0 0-2 0l1 1.118Z"/>`
	windInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h9.11A1.89 1.89 0 0 0 16 6.11v0c0-1.615-1.894-2.486-3.12-1.435L12.5 5M3 12h14.902C19.06 12 20 12.94 20 14.098v0c0 2.152-2.853 2.91-3.92 1.041L16 15M5 16h6.11A1.89 1.89 0 0 1 13 17.89v0c0 1.615-1.894 2.486-3.12 1.435L9.5 19"/>`
	zoomInInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M16 10a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z"/><path stroke-linecap="round" d="m20 20l-5-5m-7.5-5H10m0 0h2.5M10 10v2.5m0-2.5V7.5"/></g>`
	zoomOutInnerSVG              = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M16 10a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z"/><path stroke-linecap="round" d="m20 20l-5-5m-7.5-5h5"/></g>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Activity(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(activityInnerSVG).
		Element(children...)
}

func Adjustments(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustmentsInnerSVG).
		Element(children...)
}

func Aid(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aidInnerSVG).
		Element(children...)
}

func AlignObjectsBottom(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignObjectsBottomInnerSVG).
		Element(children...)
}

func AlignObjectsCenter(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignObjectsCenterInnerSVG).
		Element(children...)
}

func AlignObjectsLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignObjectsLeftInnerSVG).
		Element(children...)
}

func AlignObjectsMiddle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignObjectsMiddleInnerSVG).
		Element(children...)
}

func AlignObjectsRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignObjectsRightInnerSVG).
		Element(children...)
}

func AlignObjectsTop(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignObjectsTopInnerSVG).
		Element(children...)
}

func AlignTextCenter(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextCenterInnerSVG).
		Element(children...)
}

func AlignTextJustify(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextJustifyInnerSVG).
		Element(children...)
}

func AlignTextLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextLeftInnerSVG).
		Element(children...)
}

func AlignTextRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextRightInnerSVG).
		Element(children...)
}

func Archive(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archiveInnerSVG).
		Element(children...)
}

func ArrowDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownInnerSVG).
		Element(children...)
}

func ArrowGoBack(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowGoBackInnerSVG).
		Element(children...)
}

func ArrowGoForward(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowGoForwardInnerSVG).
		Element(children...)
}

func ArrowJoin(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowJoinInnerSVG).
		Element(children...)
}

func ArrowLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftInnerSVG).
		Element(children...)
}

func ArrowLeftDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftDownInnerSVG).
		Element(children...)
}

func ArrowLeftUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftUpInnerSVG).
		Element(children...)
}

func ArrowMainSplitSide(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowMainSplitSideInnerSVG).
		Element(children...)
}

func ArrowRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightInnerSVG).
		Element(children...)
}

func ArrowRightDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightDownInnerSVG).
		Element(children...)
}

func ArrowRightUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightUpInnerSVG).
		Element(children...)
}

func ArrowSideJoinMain(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowSideJoinMainInnerSVG).
		Element(children...)
}

func ArrowSplit(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowSplitInnerSVG).
		Element(children...)
}

func ArrowUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpInnerSVG).
		Element(children...)
}

func Arrows(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsInnerSVG).
		Element(children...)
}

func ArrowsHorizontal(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsHorizontalInnerSVG).
		Element(children...)
}

func ArrowsVertical(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsVerticalInnerSVG).
		Element(children...)
}

func AsteriskSimple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(asteriskSimpleInnerSVG).
		Element(children...)
}

func AtSymbol(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atSymbolInnerSVG).
		Element(children...)
}

func Ban(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(banInnerSVG).
		Element(children...)
}

func Bandage(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bandageInnerSVG).
		Element(children...)
}

func Bars(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barsInnerSVG).
		Element(children...)
}

func Basket(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketInnerSVG).
		Element(children...)
}

func Battery(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryInnerSVG).
		Element(children...)
}

func BatteryCharging(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryChargingInnerSVG).
		Element(children...)
}

func BatteryFull(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFullInnerSVG).
		Element(children...)
}

func BatteryHalf(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryHalfInnerSVG).
		Element(children...)
}

func Bell(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellInnerSVG).
		Element(children...)
}

func BellOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellOffInnerSVG).
		Element(children...)
}

func Bike(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bikeInnerSVG).
		Element(children...)
}

func Bold(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boldInnerSVG).
		Element(children...)
}

func Book(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookInnerSVG).
		Element(children...)
}

func BookOpen(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookOpenInnerSVG).
		Element(children...)
}

func Bookmark(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkInnerSVG).
		Element(children...)
}

func Box(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxInnerSVG).
		Element(children...)
}

func BrandGithub(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brandGithubInnerSVG).
		Element(children...)
}

func BrandTwitter(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brandTwitterInnerSVG).
		Element(children...)
}

func Briefcase(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseInnerSVG).
		Element(children...)
}

func BrushBig(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brushBigInnerSVG).
		Element(children...)
}

func Bulb(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bulbInnerSVG).
		Element(children...)
}

func BulbOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bulbOffInnerSVG).
		Element(children...)
}

func Calendar(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarInnerSVG).
		Element(children...)
}

func Camera(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraInnerSVG).
		Element(children...)
}

func CameraOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraOffInnerSVG).
		Element(children...)
}

func CameraVideo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraVideoInnerSVG).
		Element(children...)
}

func CameraVideoOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraVideoOffInnerSVG).
		Element(children...)
}

func Car(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(carInnerSVG).
		Element(children...)
}

func Cart(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartInnerSVG).
		Element(children...)
}

func Chart(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartInnerSVG).
		Element(children...)
}

func Chat(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatInnerSVG).
		Element(children...)
}

func Chats(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatsInnerSVG).
		Element(children...)
}

func Check(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkInnerSVG).
		Element(children...)
}

func CheckCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkCircleInnerSVG).
		Element(children...)
}

func ChevronDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronDownInnerSVG).
		Element(children...)
}

func ChevronLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronLeftInnerSVG).
		Element(children...)
}

func ChevronRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightInnerSVG).
		Element(children...)
}

func ChevronUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpInnerSVG).
		Element(children...)
}

func Circle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleInnerSVG).
		Element(children...)
}

func Clipboard(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardInnerSVG).
		Element(children...)
}

func Clock(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockInnerSVG).
		Element(children...)
}

func Cloud(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudInnerSVG).
		Element(children...)
}

func CloudSun(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudSunInnerSVG).
		Element(children...)
}

func Code(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeInnerSVG).
		Element(children...)
}

func Coffee(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coffeeInnerSVG).
		Element(children...)
}

func Cog(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cogInnerSVG).
		Element(children...)
}

func Coins(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coinsInnerSVG).
		Element(children...)
}

func ColumnsOneTwoThirds(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(columnsOneTwoThirdsInnerSVG).
		Element(children...)
}

func ColumnsThreeThirds(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(columnsThreeThirdsInnerSVG).
		Element(children...)
}

func ColumnsTwoHalfs(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(columnsTwoHalfsInnerSVG).
		Element(children...)
}

func CornerDownLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerDownLeftInnerSVG).
		Element(children...)
}

func CornerDownRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerDownRightInnerSVG).
		Element(children...)
}

func CornerLeftDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerLeftDownInnerSVG).
		Element(children...)
}

func CornerLeftUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerLeftUpInnerSVG).
		Element(children...)
}

func CornerRightDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerRightDownInnerSVG).
		Element(children...)
}

func CornerRightUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerRightUpInnerSVG).
		Element(children...)
}

func CornerTopLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerTopLeftInnerSVG).
		Element(children...)
}

func CornerUpRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerUpRightInnerSVG).
		Element(children...)
}

func Cpu(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cpuInnerSVG).
		Element(children...)
}

func CreativeCommons(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creativeCommonsInnerSVG).
		Element(children...)
}

func CreativeCommonsBy(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creativeCommonsByInnerSVG).
		Element(children...)
}

func CreativeCommonsNd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creativeCommonsNdInnerSVG).
		Element(children...)
}

func CreativeCommonsSa(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creativeCommonsSaInnerSVG).
		Element(children...)
}

func CreditCard(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardInnerSVG).
		Element(children...)
}

func Crop(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropInnerSVG).
		Element(children...)
}

func Crown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crownInnerSVG).
		Element(children...)
}

func CurrencyDollarCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(currencyDollarCircleInnerSVG).
		Element(children...)
}

func CurrencyEuroCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(currencyEuroCircleInnerSVG).
		Element(children...)
}

func CurrencyPoundCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(currencyPoundCircleInnerSVG).
		Element(children...)
}

func Dashboard(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashboardInnerSVG).
		Element(children...)
}

func Database(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseInnerSVG).
		Element(children...)
}

func Desktop(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(desktopInnerSVG).
		Element(children...)
}

func Document(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentInnerSVG).
		Element(children...)
}

func DocumentAdd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentAddInnerSVG).
		Element(children...)
}

func DocumentRemove(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentRemoveInnerSVG).
		Element(children...)
}

func Documents(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentsInnerSVG).
		Element(children...)
}

func DotsHorizontal(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dotsHorizontalInnerSVG).
		Element(children...)
}

func DotsVertical(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dotsVerticalInnerSVG).
		Element(children...)
}

func Download(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadInnerSVG).
		Element(children...)
}

func DownloadAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadAltInnerSVG).
		Element(children...)
}

func Droplet(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropletInnerSVG).
		Element(children...)
}

func Duplicate(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(duplicateInnerSVG).
		Element(children...)
}

func ExchangeHorizontal(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exchangeHorizontalInnerSVG).
		Element(children...)
}

func ExchangeVertical(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exchangeVerticalInnerSVG).
		Element(children...)
}

func Exclamation(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationInnerSVG).
		Element(children...)
}

func ExclamationTriangle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationTriangleInnerSVG).
		Element(children...)
}

func Expand(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandInnerSVG).
		Element(children...)
}

func ExternalLink(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(externalLinkInnerSVG).
		Element(children...)
}

func Eye(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeInnerSVG).
		Element(children...)
}

func EyeClose(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeCloseInnerSVG).
		Element(children...)
}

func EyeOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeOffInnerSVG).
		Element(children...)
}

func FastForward(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastForwardInnerSVG).
		Element(children...)
}

func Fingerprint(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintInnerSVG).
		Element(children...)
}

func Flag(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagInnerSVG).
		Element(children...)
}

func Flash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flashInnerSVG).
		Element(children...)
}

func Flask(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flaskInnerSVG).
		Element(children...)
}

func Folder(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderInnerSVG).
		Element(children...)
}

func FolderAdd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderAddInnerSVG).
		Element(children...)
}

func FolderOpen(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderOpenInnerSVG).
		Element(children...)
}

func FolderRemove(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderRemoveInnerSVG).
		Element(children...)
}

func ForkKnife(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forkKnifeInnerSVG).
		Element(children...)
}

func Funnel(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(funnelInnerSVG).
		Element(children...)
}

func Gift(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(giftInnerSVG).
		Element(children...)
}

func Globe(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeInnerSVG).
		Element(children...)
}

func Heading(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headingInnerSVG).
		Element(children...)
}

func Heart(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartInnerSVG).
		Element(children...)
}

func Home(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeInnerSVG).
		Element(children...)
}

func Humbleicon(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(humbleiconInnerSVG).
		Element(children...)
}

func Image(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageInnerSVG).
		Element(children...)
}

func Images(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imagesInnerSVG).
		Element(children...)
}

func Incognito(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(incognitoInnerSVG).
		Element(children...)
}

func IncognitoTwo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(incognitoTwoInnerSVG).
		Element(children...)
}

func InfoCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoCircleInnerSVG).
		Element(children...)
}

func Italic(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(italicInnerSVG).
		Element(children...)
}

func Key(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyInnerSVG).
		Element(children...)
}

func Layers(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersInnerSVG).
		Element(children...)
}

func Link(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkInnerSVG).
		Element(children...)
}

func Location(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationInnerSVG).
		Element(children...)
}

func Lock(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockInnerSVG).
		Element(children...)
}

func LockOpen(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockOpenInnerSVG).
		Element(children...)
}

func Logout(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(logoutInnerSVG).
		Element(children...)
}

func Mail(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailInnerSVG).
		Element(children...)
}

func MailOpen(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailOpenInnerSVG).
		Element(children...)
}

func Map(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapInnerSVG).
		Element(children...)
}

func Maximize(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maximizeInnerSVG).
		Element(children...)
}

func Microphone(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneInnerSVG).
		Element(children...)
}

func MicrophoneOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneOffInnerSVG).
		Element(children...)
}

func Minus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusInnerSVG).
		Element(children...)
}

func MinusCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircleInnerSVG).
		Element(children...)
}

func Mobile(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mobileInnerSVG).
		Element(children...)
}

func Money(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneyInnerSVG).
		Element(children...)
}

func Moon(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonInnerSVG).
		Element(children...)
}

func Moustache(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moustacheInnerSVG).
		Element(children...)
}

func MusicNote(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(musicNoteInnerSVG).
		Element(children...)
}

func Navigation(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navigationInnerSVG).
		Element(children...)
}

func Package(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(packageInnerSVG).
		Element(children...)
}

func Pause(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseInnerSVG).
		Element(children...)
}

func Pencil(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pencilInnerSVG).
		Element(children...)
}

func Phone(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneInnerSVG).
		Element(children...)
}

func PhoneCall(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneCallInnerSVG).
		Element(children...)
}

func PhoneForward(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneForwardInnerSVG).
		Element(children...)
}

func PhoneIncoming(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneIncomingInnerSVG).
		Element(children...)
}

func PhoneMissed(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneMissedInnerSVG).
		Element(children...)
}

func PhoneOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneOffInnerSVG).
		Element(children...)
}

func PhoneOutgoing(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneOutgoingInnerSVG).
		Element(children...)
}

func PieChart(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartInnerSVG).
		Element(children...)
}

func Play(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playInnerSVG).
		Element(children...)
}

func Plus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusInnerSVG).
		Element(children...)
}

func PlusCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircleInnerSVG).
		Element(children...)
}

func Power(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerInnerSVG).
		Element(children...)
}

func Print(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printInnerSVG).
		Element(children...)
}

func Radio(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radioInnerSVG).
		Element(children...)
}

func Rain(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rainInnerSVG).
		Element(children...)
}

func Refresh(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshInnerSVG).
		Element(children...)
}

func Rewind(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rewindInnerSVG).
		Element(children...)
}

func Rocket(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rocketInnerSVG).
		Element(children...)
}

func Rss(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssInnerSVG).
		Element(children...)
}

func Save(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(saveInnerSVG).
		Element(children...)
}

func Scissors(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scissorsInnerSVG).
		Element(children...)
}

func Search(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchInnerSVG).
		Element(children...)
}

func Share(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareInnerSVG).
		Element(children...)
}

func ShareAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareAltInnerSVG).
		Element(children...)
}

func Ship(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shipInnerSVG).
		Element(children...)
}

func SkipBackward(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skipBackwardInnerSVG).
		Element(children...)
}

func SkipForward(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skipForwardInnerSVG).
		Element(children...)
}

func Snow(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snowInnerSVG).
		Element(children...)
}

func Square(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareInnerSVG).
		Element(children...)
}

func Star(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starInnerSVG).
		Element(children...)
}

func Storm(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stormInnerSVG).
		Element(children...)
}

func Sun(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunInnerSVG).
		Element(children...)
}

func Support(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(supportInnerSVG).
		Element(children...)
}

func SwitchOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(switchOffInnerSVG).
		Element(children...)
}

func SwitchOn(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(switchOnInnerSVG).
		Element(children...)
}

func Tag(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagInnerSVG).
		Element(children...)
}

func Tags(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagsInnerSVG).
		Element(children...)
}

func Text(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textInnerSVG).
		Element(children...)
}

func Times(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timesInnerSVG).
		Element(children...)
}

func TimesCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timesCircleInnerSVG).
		Element(children...)
}

func Trash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashInnerSVG).
		Element(children...)
}

func TrendingDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendingDownInnerSVG).
		Element(children...)
}

func TrendingUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendingUpInnerSVG).
		Element(children...)
}

func Triangle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleInnerSVG).
		Element(children...)
}

func Truck(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(truckInnerSVG).
		Element(children...)
}

func Underline(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(underlineInnerSVG).
		Element(children...)
}

func Upload(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadInnerSVG).
		Element(children...)
}

func User(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userInnerSVG).
		Element(children...)
}

func UserAdd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userAddInnerSVG).
		Element(children...)
}

func UserAsking(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userAskingInnerSVG).
		Element(children...)
}

func UserRemove(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userRemoveInnerSVG).
		Element(children...)
}

func Users(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usersInnerSVG).
		Element(children...)
}

func Verified(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(verifiedInnerSVG).
		Element(children...)
}

func ViewGrid(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewGridInnerSVG).
		Element(children...)
}

func ViewList(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewListInnerSVG).
		Element(children...)
}

func Volume(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeInnerSVG).
		Element(children...)
}

func VolumeOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeOffInnerSVG).
		Element(children...)
}

func VolumeOne(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeOneInnerSVG).
		Element(children...)
}

func VolumeTwo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeTwoInnerSVG).
		Element(children...)
}

func Wifi(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiInnerSVG).
		Element(children...)
}

func WifiOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiOffInnerSVG).
		Element(children...)
}

func Wind(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windInnerSVG).
		Element(children...)
}

func ZoomIn(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomInInnerSVG).
		Element(children...)
}

func ZoomOut(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomOutInnerSVG).
		Element(children...)
}

func ByName(name string) (*engine.UberElement, error) {
	switch name {
	case "activity":
		return Activity(), nil
	case "adjustments":
		return Adjustments(), nil
	case "aid":
		return Aid(), nil
	case "align-objects-bottom":
		return AlignObjectsBottom(), nil
	case "align-objects-center":
		return AlignObjectsCenter(), nil
	case "align-objects-left":
		return AlignObjectsLeft(), nil
	case "align-objects-middle":
		return AlignObjectsMiddle(), nil
	case "align-objects-right":
		return AlignObjectsRight(), nil
	case "align-objects-top":
		return AlignObjectsTop(), nil
	case "align-text-center":
		return AlignTextCenter(), nil
	case "align-text-justify":
		return AlignTextJustify(), nil
	case "align-text-left":
		return AlignTextLeft(), nil
	case "align-text-right":
		return AlignTextRight(), nil
	case "archive":
		return Archive(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-go-back":
		return ArrowGoBack(), nil
	case "arrow-go-forward":
		return ArrowGoForward(), nil
	case "arrow-join":
		return ArrowJoin(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-left-down":
		return ArrowLeftDown(), nil
	case "arrow-left-up":
		return ArrowLeftUp(), nil
	case "arrow-main-split-side":
		return ArrowMainSplitSide(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-right-down":
		return ArrowRightDown(), nil
	case "arrow-right-up":
		return ArrowRightUp(), nil
	case "arrow-side-join-main":
		return ArrowSideJoinMain(), nil
	case "arrow-split":
		return ArrowSplit(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrows":
		return Arrows(), nil
	case "arrows-horizontal":
		return ArrowsHorizontal(), nil
	case "arrows-vertical":
		return ArrowsVertical(), nil
	case "asterisk-simple":
		return AsteriskSimple(), nil
	case "at-symbol":
		return AtSymbol(), nil
	case "ban":
		return Ban(), nil
	case "bandage":
		return Bandage(), nil
	case "bars":
		return Bars(), nil
	case "basket":
		return Basket(), nil
	case "battery":
		return Battery(), nil
	case "battery-charging":
		return BatteryCharging(), nil
	case "battery-full":
		return BatteryFull(), nil
	case "battery-half":
		return BatteryHalf(), nil
	case "bell":
		return Bell(), nil
	case "bell-off":
		return BellOff(), nil
	case "bike":
		return Bike(), nil
	case "bold":
		return Bold(), nil
	case "book":
		return Book(), nil
	case "book-open":
		return BookOpen(), nil
	case "bookmark":
		return Bookmark(), nil
	case "box":
		return Box(), nil
	case "brand-github":
		return BrandGithub(), nil
	case "brand-twitter":
		return BrandTwitter(), nil
	case "briefcase":
		return Briefcase(), nil
	case "brush-big":
		return BrushBig(), nil
	case "bulb":
		return Bulb(), nil
	case "bulb-off":
		return BulbOff(), nil
	case "calendar":
		return Calendar(), nil
	case "camera":
		return Camera(), nil
	case "camera-off":
		return CameraOff(), nil
	case "camera-video":
		return CameraVideo(), nil
	case "camera-video-off":
		return CameraVideoOff(), nil
	case "car":
		return Car(), nil
	case "cart":
		return Cart(), nil
	case "chart":
		return Chart(), nil
	case "chat":
		return Chat(), nil
	case "chats":
		return Chats(), nil
	case "check":
		return Check(), nil
	case "check-circle":
		return CheckCircle(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "circle":
		return Circle(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clock":
		return Clock(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-sun":
		return CloudSun(), nil
	case "code":
		return Code(), nil
	case "coffee":
		return Coffee(), nil
	case "cog":
		return Cog(), nil
	case "coins":
		return Coins(), nil
	case "columns-one-two-thirds":
		return ColumnsOneTwoThirds(), nil
	case "columns-three-thirds":
		return ColumnsThreeThirds(), nil
	case "columns-two-halfs":
		return ColumnsTwoHalfs(), nil
	case "corner-down-left":
		return CornerDownLeft(), nil
	case "corner-down-right":
		return CornerDownRight(), nil
	case "corner-left-down":
		return CornerLeftDown(), nil
	case "corner-left-up":
		return CornerLeftUp(), nil
	case "corner-right-down":
		return CornerRightDown(), nil
	case "corner-right-up":
		return CornerRightUp(), nil
	case "corner-top-left":
		return CornerTopLeft(), nil
	case "corner-up-right":
		return CornerUpRight(), nil
	case "cpu":
		return Cpu(), nil
	case "creative-commons":
		return CreativeCommons(), nil
	case "creative-commons-by":
		return CreativeCommonsBy(), nil
	case "creative-commons-nd":
		return CreativeCommonsNd(), nil
	case "creative-commons-sa":
		return CreativeCommonsSa(), nil
	case "credit-card":
		return CreditCard(), nil
	case "crop":
		return Crop(), nil
	case "crown":
		return Crown(), nil
	case "currency-dollar-circle":
		return CurrencyDollarCircle(), nil
	case "currency-euro-circle":
		return CurrencyEuroCircle(), nil
	case "currency-pound-circle":
		return CurrencyPoundCircle(), nil
	case "dashboard":
		return Dashboard(), nil
	case "database":
		return Database(), nil
	case "desktop":
		return Desktop(), nil
	case "document":
		return Document(), nil
	case "document-add":
		return DocumentAdd(), nil
	case "document-remove":
		return DocumentRemove(), nil
	case "documents":
		return Documents(), nil
	case "dots-horizontal":
		return DotsHorizontal(), nil
	case "dots-vertical":
		return DotsVertical(), nil
	case "download":
		return Download(), nil
	case "download-alt":
		return DownloadAlt(), nil
	case "droplet":
		return Droplet(), nil
	case "duplicate":
		return Duplicate(), nil
	case "exchange-horizontal":
		return ExchangeHorizontal(), nil
	case "exchange-vertical":
		return ExchangeVertical(), nil
	case "exclamation":
		return Exclamation(), nil
	case "exclamation-triangle":
		return ExclamationTriangle(), nil
	case "expand":
		return Expand(), nil
	case "external-link":
		return ExternalLink(), nil
	case "eye":
		return Eye(), nil
	case "eye-close":
		return EyeClose(), nil
	case "eye-off":
		return EyeOff(), nil
	case "fast-forward":
		return FastForward(), nil
	case "fingerprint":
		return Fingerprint(), nil
	case "flag":
		return Flag(), nil
	case "flash":
		return Flash(), nil
	case "flask":
		return Flask(), nil
	case "folder":
		return Folder(), nil
	case "folder-add":
		return FolderAdd(), nil
	case "folder-open":
		return FolderOpen(), nil
	case "folder-remove":
		return FolderRemove(), nil
	case "fork-knife":
		return ForkKnife(), nil
	case "funnel":
		return Funnel(), nil
	case "gift":
		return Gift(), nil
	case "globe":
		return Globe(), nil
	case "heading":
		return Heading(), nil
	case "heart":
		return Heart(), nil
	case "home":
		return Home(), nil
	case "humbleicon":
		return Humbleicon(), nil
	case "image":
		return Image(), nil
	case "images":
		return Images(), nil
	case "incognito":
		return Incognito(), nil
	case "incognito-2":
		return IncognitoTwo(), nil
	case "info-circle":
		return InfoCircle(), nil
	case "italic":
		return Italic(), nil
	case "key":
		return Key(), nil
	case "layers":
		return Layers(), nil
	case "link":
		return Link(), nil
	case "location":
		return Location(), nil
	case "lock":
		return Lock(), nil
	case "lock-open":
		return LockOpen(), nil
	case "logout":
		return Logout(), nil
	case "mail":
		return Mail(), nil
	case "mail-open":
		return MailOpen(), nil
	case "map":
		return Map(), nil
	case "maximize":
		return Maximize(), nil
	case "microphone":
		return Microphone(), nil
	case "microphone-off":
		return MicrophoneOff(), nil
	case "minus":
		return Minus(), nil
	case "minus-circle":
		return MinusCircle(), nil
	case "mobile":
		return Mobile(), nil
	case "money":
		return Money(), nil
	case "moon":
		return Moon(), nil
	case "moustache":
		return Moustache(), nil
	case "music-note":
		return MusicNote(), nil
	case "navigation":
		return Navigation(), nil
	case "package":
		return Package(), nil
	case "pause":
		return Pause(), nil
	case "pencil":
		return Pencil(), nil
	case "phone":
		return Phone(), nil
	case "phone-call":
		return PhoneCall(), nil
	case "phone-forward":
		return PhoneForward(), nil
	case "phone-incoming":
		return PhoneIncoming(), nil
	case "phone-missed":
		return PhoneMissed(), nil
	case "phone-off":
		return PhoneOff(), nil
	case "phone-outgoing":
		return PhoneOutgoing(), nil
	case "pie-chart":
		return PieChart(), nil
	case "play":
		return Play(), nil
	case "plus":
		return Plus(), nil
	case "plus-circle":
		return PlusCircle(), nil
	case "power":
		return Power(), nil
	case "print":
		return Print(), nil
	case "radio":
		return Radio(), nil
	case "rain":
		return Rain(), nil
	case "refresh":
		return Refresh(), nil
	case "rewind":
		return Rewind(), nil
	case "rocket":
		return Rocket(), nil
	case "rss":
		return Rss(), nil
	case "save":
		return Save(), nil
	case "scissors":
		return Scissors(), nil
	case "search":
		return Search(), nil
	case "share":
		return Share(), nil
	case "share-alt":
		return ShareAlt(), nil
	case "ship":
		return Ship(), nil
	case "skip-backward":
		return SkipBackward(), nil
	case "skip-forward":
		return SkipForward(), nil
	case "snow":
		return Snow(), nil
	case "square":
		return Square(), nil
	case "star":
		return Star(), nil
	case "storm":
		return Storm(), nil
	case "sun":
		return Sun(), nil
	case "support":
		return Support(), nil
	case "switch-off":
		return SwitchOff(), nil
	case "switch-on":
		return SwitchOn(), nil
	case "tag":
		return Tag(), nil
	case "tags":
		return Tags(), nil
	case "text":
		return Text(), nil
	case "times":
		return Times(), nil
	case "times-circle":
		return TimesCircle(), nil
	case "trash":
		return Trash(), nil
	case "trending-down":
		return TrendingDown(), nil
	case "trending-up":
		return TrendingUp(), nil
	case "triangle":
		return Triangle(), nil
	case "truck":
		return Truck(), nil
	case "underline":
		return Underline(), nil
	case "upload":
		return Upload(), nil
	case "user":
		return User(), nil
	case "user-add":
		return UserAdd(), nil
	case "user-asking":
		return UserAsking(), nil
	case "user-remove":
		return UserRemove(), nil
	case "users":
		return Users(), nil
	case "verified":
		return Verified(), nil
	case "view-grid":
		return ViewGrid(), nil
	case "view-list":
		return ViewList(), nil
	case "volume":
		return Volume(), nil
	case "volume-off":
		return VolumeOff(), nil
	case "volume-1":
		return VolumeOne(), nil
	case "volume-2":
		return VolumeTwo(), nil
	case "wifi":
		return Wifi(), nil
	case "wifi-off":
		return WifiOff(), nil
	case "wind":
		return Wind(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in humbleicons icon set", name)
	}
}
