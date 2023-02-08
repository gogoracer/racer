package feather

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	activityInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12h-4l-3 9L9 3l-3 9H2"/>`
	airplayInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 17H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1"/><path d="m12 15l5 6H7l5-6z"/></g>`
	alertCircleInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 8v4m0 4h.01"/></g>`
	alertOctagonInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2zM12 8v4m0 4h.01"/>`
	alertTriangleInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0zM12 9v4m0 4h.01"/>`
	alignCenterInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 10H6m15-4H3m18 8H3m15 4H6"/>`
	alignJustifyInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 10H3m18-4H3m18 8H3m18 4H3"/>`
	alignLeftInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 10H3m18-4H3m18 8H3m14 4H3"/>`
	alignRightInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 10H7m14-4H3m18 8H3m18 4H7"/>`
	anchorInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="3"/><path d="M12 22V8m-7 4H2a10 10 0 0 0 20 0h-3"/></g>`
	apertureInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m14.31 8l5.74 9.94M9.69 8h11.48M7.38 12l5.74-9.94M9.69 16L3.95 6.06M14.31 16H2.83m13.79-4l-5.74 9.94"/></g>`
	archiveInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 8v13H3V8M1 3h22v5H1zm9 9h4"/>`
	arrowDownInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v14m7-7l-7 7l-7-7"/>`
	arrowDownCircleInnerSVG  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m8 12l4 4l4-4m-4-4v8"/></g>`
	arrowDownLeftInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7L7 17m10 0H7V7"/>`
	arrowDownRightInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10m0-10v10H7"/>`
	arrowLeftInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 12H5m7 7l-7-7l7-7"/>`
	arrowLeftCircleInnerSVG  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m12 8l-4 4l4 4m4-4H8"/></g>`
	arrowRightInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7-7l7 7l-7 7"/>`
	arrowRightCircleInnerSVG = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m12 16l4-4l-4-4m-4 4h8"/></g>`
	arrowUpInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19V5m-7 7l7-7l7 7"/>`
	arrowUpCircleInnerSVG    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m16 12l-4-4l-4 4m4 4V8"/></g>`
	arrowUpLeftInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17L7 7m0 10V7h10"/>`
	arrowUpRightInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 17L17 7M7 7h10v10"/>`
	atSignInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-3.92 7.94"/></g>`
	awardInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="8" r="7"/><path d="M8.21 13.89L7 23l5-3l5 3l-1.21-9.12"/></g>`
	barChartInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20V10m6 10V4M6 20v-4"/>`
	barChartTwoInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 20V10m-6 10V4M6 20v-6"/>`
	batteryInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="1" y="6" rx="2" ry="2"/><path d="M23 13v-2"/></g>`
	batteryChargingInnerSVG  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 18H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3.19M15 6h2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-3.19M23 13v-2M11 6l-4 6h6l-4 6"/>`
	bellInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9m-4.27 13a2 2 0 0 1-3.46 0"/>`
	bellOffInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.73 21a2 2 0 0 1-3.46 0m8.36-8A17.89 17.89 0 0 1 18 8M6.26 6.26A5.86 5.86 0 0 0 6 8c0 7-3 9-3 9h14m1-9a6 6 0 0 0-9.33-5M1 1l22 22"/>`
	bluetoothInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.5 6.5l11 11L12 23V1l5.5 5.5l-11 11"/>`
	boldInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4h8a4 4 0 0 1 4 4a4 4 0 0 1-4 4H6zm0 8h9a4 4 0 0 1 4 4a4 4 0 0 1-4 4H6z"/>`
	bookInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/><path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/></g>`
	bookOpenInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2zm20 0h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"/>`
	bookmarkInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 21l-7-5l-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/>`
	boxInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><path d="M3.27 6.96L12 12.01l8.73-5.05M12 22.08V12"/></g>`
	briefcaseInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="7" rx="2" ry="2"/><path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"/></g>`
	calendarInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18"/></g>`
	cameraInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/><circle cx="12" cy="13" r="4"/></g>`
	cameraOffInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1l22 22m-2-2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3m3-3h6l2 3h4a2 2 0 0 1 2 2v9.34m-7.72-2.06a4 4 0 1 1-5.56-5.56"/>`
	castInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 16.1A5 5 0 0 1 5.9 20M2 12.05A9 9 0 0 1 9.95 20M2 8V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-6M2 20h.01"/>`
	checkInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6L9 17l-5-5"/>`
	checkCircleInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><path d="M22 4L12 14.01l-3-3"/></g>`
	checkSquareInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 11l3 3L22 4"/><path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/></g>`
	chevronDownInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 9l6 6l6-6"/>`
	chevronLeftInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 18l-6-6l6-6"/>`
	chevronRightInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 18l6-6l-6-6"/>`
	chevronUpInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 15l-6-6l-6 6"/>`
	chevronsDownInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 13l5 5l5-5M7 6l5 5l5-5"/>`
	chevronsLeftInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 17l-5-5l5-5m7 10l-5-5l5-5"/>`
	chevronsRightInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 17l5-5l-5-5M6 17l5-5l-5-5"/>`
	chevronsUpInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 11l-5-5l-5 5m10 7l-5-5l-5 5"/>`
	chromeInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="4"/><path d="M21.17 8H12M3.95 6.06L8.54 14m2.34 7.94L15.46 14"/></g>`
	circleInnerSVG           = `<circle cx="12" cy="12" r="10" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>`
	clipboardInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/></g>`
	clockInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/></g>`
	cloudInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 10h-1.26A8 8 0 1 0 9 20h9a5 5 0 0 0 0-10z"/>`
	cloudDrizzleInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19v2m0-8v2m8 4v2m0-8v2m-4 6v2m0-8v2m8-.42A5 5 0 0 0 18 7h-1.26A8 8 0 1 0 4 15.25"/>`
	cloudLightningInnerSVG   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 16.9A5 5 0 0 0 18 7h-1.26a8 8 0 1 0-11.62 9"/><path d="m13 11l-4 6h6l-4 6"/></g>`
	cloudOffInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.61 16.95A5 5 0 0 0 18 10h-1.26a8 8 0 0 0-7.05-6M5 5a8 8 0 0 0 4 15h9a5 5 0 0 0 1.7-.3M1 1l22 22"/>`
	cloudRainInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 13v8m-8-8v8m4-6v8m8-6.42A5 5 0 0 0 18 7h-1.26A8 8 0 1 0 4 15.25"/>`
	cloudSnowInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 17.58A5 5 0 0 0 18 8h-1.26A8 8 0 1 0 4 16.25M8 16h.01M8 20h.01M12 18h.01M12 22h.01M16 16h.01M16 20h.01"/>`
	codeInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 18l6-6l-6-6M8 6l-6 6l6 6"/>`
	codepenInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 2l10 6.5v7L12 22L2 15.5v-7L12 2zm0 20v-6.5"/><path d="m22 8.5l-10 7l-10-7"/><path d="m2 15.5l10-7l10 7M12 2v6.5"/></g>`
	codesandboxInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><path d="m7.5 4.21l4.5 2.6l4.5-2.6m-9 15.58V14.6L3 12m18 0l-4.5 2.6v5.19M3.27 6.96L12 12.01l8.73-5.05M12 22.08V12"/></g>`
	coffeeInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 8h1a4 4 0 0 1 0 8h-1M2 8h16v9a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V8zm4-7v3m4-3v3m4-3v3"/>`
	columnsInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3h7a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-7m0-18H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7m0-18v18"/>`
	commandInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3a3 3 0 0 0 3-3a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3a3 3 0 0 0 3 3a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3a3 3 0 0 0-3 3a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3a3 3 0 0 0-3-3z"/>`
	compassInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m16.24 7.76l-2.12 6.36l-6.36 2.12l2.12-6.36l6.36-2.12z"/></g>`
	copyInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="13" height="13" x="9" y="9" rx="2" ry="2"/><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"/></g>`
	cornerDownLeftInnerSVG   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 10l-5 5l5 5"/><path d="M20 4v7a4 4 0 0 1-4 4H4"/></g>`
	cornerDownRightInnerSVG  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 10l5 5l-5 5"/><path d="M4 4v7a4 4 0 0 0 4 4h12"/></g>`
	cornerLeftDownInnerSVG   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 15l-5 5l-5-5"/><path d="M20 4h-7a4 4 0 0 0-4 4v12"/></g>`
	cornerLeftUpInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 9L9 4L4 9"/><path d="M20 20h-7a4 4 0 0 1-4-4V4"/></g>`
	cornerRightDownInnerSVG  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 15l5 5l5-5"/><path d="M4 4h7a4 4 0 0 1 4 4v12"/></g>`
	cornerRightUpInnerSVG    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 9l5-5l5 5"/><path d="M4 20h7a4 4 0 0 0 4-4V4"/></g>`
	cornerUpLeftInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 14L4 9l5-5"/><path d="M20 20v-7a4 4 0 0 0-4-4H4"/></g>`
	cornerUpRightInnerSVG    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 14l5-5l-5-5"/><path d="M4 20v-7a4 4 0 0 1 4-4h12"/></g>`
	cpuInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="4" rx="2" ry="2"/><path d="M9 9h6v6H9zm0-8v3m6-3v3M9 20v3m6-3v3m5-14h3m-3 5h3M1 9h3m-3 5h3"/></g>`
	creditCardInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="22" height="16" x="1" y="4" rx="2" ry="2"/><path d="M1 10h22"/></g>`
	cropInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.13 1L6 16a2 2 0 0 0 2 2h15"/><path d="M1 6.13L16 6a2 2 0 0 1 2 2v15"/></g>`
	crosshairInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M22 12h-4M6 12H2m10-6V2m0 20v-4"/></g>`
	databaseInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"/><path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"/></g>`
	deleteInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 4H8l-7 8l7 8h13a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zm-3 5l-6 6m0-6l6 6"/>`
	discInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="3"/></g>`
	divideInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="6" r="2"/><path d="M5 12h14"/><circle cx="12" cy="18" r="2"/></g>`
	divideCircleInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 12h8m-4 4h0m0-8h0"/><circle cx="12" cy="12" r="10"/></g>`
	divideSquareInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M8 12h8m-4 4h0m0-8h0"/></g>`
	dollarSignInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 1v22m5-18H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/>`
	downloadInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4m4-5l5 5l5-5m-5 5V3"/>`
	downloadCloudInnerSVG    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 17l4 4l4-4m-4-5v9"/><path d="M20.88 18.09A5 5 0 0 0 18 9h-1.26A8 8 0 1 0 3 16.29"/></g>`
	dribbbleInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8.56 2.75c4.37 6.03 6.02 9.42 8.03 17.72m2.54-15.38c-3.72 4.35-8.94 5.66-16.88 5.85m19.5 1.9c-3.5-.93-6.63-.82-8.94 0c-2.58.92-5.01 2.86-7.44 6.32"/></g>`
	dropletInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 2.69l5.66 5.66a8 8 0 1 1-11.31 0z"/>`
	editInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1l1-4l9.5-9.5z"/></g>`
	editThreeInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20h9M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1l1-4L16.5 3.5z"/>`
	editTwoInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3a2.828 2.828 0 1 1 4 4L7.5 20.5L2 22l1.5-5.5L17 3z"/>`
	externalLinkInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6m4-3h6v6m-11 5L21 3"/>`
	eyeInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M1 12s4-8 11-8s11 8 11 8s-4 8-11 8s-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></g>`
	eyeOffInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24M1 1l22 22"/>`
	facebookInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/>`
	fastForwardInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 19l9-7l-9-7v14zM2 19l9-7l-9-7v14z"/>`
	featherInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5zM16 8L2 22m15.5-7H9"/>`
	figmaInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 5.5A3.5 3.5 0 0 1 8.5 2H12v7H8.5A3.5 3.5 0 0 1 5 5.5zM12 2h3.5a3.5 3.5 0 1 1 0 7H12V2z"/><path d="M12 12.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 1 1-7 0zm-7 7A3.5 3.5 0 0 1 8.5 16H12v3.5a3.5 3.5 0 1 1-7 0zm0-7A3.5 3.5 0 0 1 8.5 9H12v7H8.5A3.5 3.5 0 0 1 5 12.5z"/></g>`
	fileInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9z"/><path d="M13 2v7h7"/></g>`
	fileMinusInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><path d="M14 2v6h6M9 15h6"/></g>`
	filePlusInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><path d="M14 2v6h6m-8 10v-6m-3 3h6"/></g>`
	fileTextInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><path d="M14 2v6h6m-4 5H8m8 4H8m2-8H8"/></g>`
	filmInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2.18" ry="2.18"/><path d="M7 2v20M17 2v20M2 12h20M2 7h5M2 17h5m10 0h5M17 7h5"/></g>`
	filterInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 3H2l8 9.46V19l4 2v-8.54L22 3z"/>`
	flagInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 15s1-1 4-1s5 2 8 2s4-1 4-1V3s-1 1-4 1s-5-2-8-2s-4 1-4 1zm0 7v-7"/>`
	folderInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"/>`
	folderMinusInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2zM9 14h6"/>`
	folderPlusInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2zm-10-8v6m-3-3h6"/>`
	framerInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 16V9h14V2H5l14 14h-7m-7 0l7 7v-7m-7 0h7"/>`
	frownInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 16s-1.5-2-4-2s-4 2-4 2m1-7h.01M15 9h.01"/></g>`
	giftInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12v10H4V12M2 7h20v5H2zm10 15V7m0 0H7.5a2.5 2.5 0 0 1 0-5C11 2 12 7 12 7zm0 0h4.5a2.5 2.5 0 0 0 0-5C13 2 12 7 12 7z"/>`
	gitBranchInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3v12"/><circle cx="18" cy="6" r="3"/><circle cx="6" cy="18" r="3"/><path d="M18 9a9 9 0 0 1-9 9"/></g>`
	gitCommitInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M1.05 12H7m10.01 0h5.95"/></g>`
	gitMergeInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M6 21V9a9 9 0 0 0 9 9"/></g>`
	gitPullRequestInnerSVG   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M13 6h3a2 2 0 0 1 2 2v7M6 9v12"/></g>`
	githubInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77A5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"/>`
	gitlabInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.65 14.39L12 22.13L1.35 14.39a.84.84 0 0 1-.3-.94l1.22-3.78l2.44-7.51A.42.42 0 0 1 4.82 2a.43.43 0 0 1 .58 0a.42.42 0 0 1 .11.18l2.44 7.49h8.1l2.44-7.51A.42.42 0 0 1 18.6 2a.43.43 0 0 1 .58 0a.42.42 0 0 1 .11.18l2.44 7.51L23 13.45a.84.84 0 0 1-.35.94z"/>`
	globeInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M2 12h20M12 2a15.3 15.3 0 0 1 4 10a15.3 15.3 0 0 1-4 10a15.3 15.3 0 0 1-4-10a15.3 15.3 0 0 1 4-10z"/></g>`
	gridInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h7v7H3zm11 0h7v7h-7zm0 11h7v7h-7zM3 14h7v7H3z"/>`
	hardDriveInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12H2m3.45-6.89L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11zM6 16h.01M10 16h.01"/>`
	hashInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9h16M4 15h16M10 3L8 21m8-18l-2 18"/>`
	headphonesInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 18v-6a9 9 0 0 1 18 0v6"/><path d="M21 19a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3zM3 19a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2v-3a2 2 0 0 0-2-2H3z"/></g>`
	heartInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78l1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>`
	helpCircleInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3m.08 4h.01"/></g>`
	hexagonInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>`
	homeInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 9l9-7l9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="M9 22V12h6v10"/></g>`
	imageInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><circle cx="8.5" cy="8.5" r="1.5"/><path d="m21 15l-5-5L5 21"/></g>`
	inboxInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12h-6l-2 3h-4l-2-3H2"/><path d="M5.45 5.11L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"/></g>`
	infoInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4m0-4h.01"/></g>`
	instagramInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="5" ry="5"/><path d="M16 11.37A4 4 0 1 1 12.63 8A4 4 0 0 1 16 11.37zm1.5-4.87h.01"/></g>`
	italicInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4h-9m4 16H5M15 4L9 20"/>`
	keyInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 2l-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778a5.5 5.5 0 0 1 7.777-7.777zm0 0L15.5 7.5m0 0l3 3L22 7l-3-3m-3.5 3.5L19 4"/>`
	layersInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2L2 7l10 5l10-5l-10-5zM2 17l10 5l10-5M2 12l10 5l10-5"/>`
	layoutInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M9 21V9"/></g>`
	lifeBuoyInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="4"/><path d="m4.93 4.93l4.24 4.24m5.66 5.66l4.24 4.24m-4.24-9.9l4.24-4.24m-4.24 4.24l3.53-3.53M4.93 19.07l4.24-4.24"/></g>`
	linkInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></g>`
	linkTwoInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7h3a5 5 0 0 1 5 5a5 5 0 0 1-5 5h-3m-6 0H6a5 5 0 0 1-5-5a5 5 0 0 1 5-5h3m-1 5h8"/>`
	linkedinInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2a2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6zM2 9h4v12H2z"/><circle cx="4" cy="4" r="2"/></g>`
	listInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6h13M8 12h13M8 18h13M3 6h.01M3 12h.01M3 18h.01"/>`
	loaderInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v4m0 12v4M4.93 4.93l2.83 2.83m8.48 8.48l2.83 2.83M2 12h4m12 0h4M4.93 19.07l2.83-2.83m8.48-8.48l2.83-2.83"/>`
	lockInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></g>`
	logInInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4m-5-4l5-5l-5-5m5 5H3"/>`
	logOutInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4m7 14l5-5l-5-5m5 5H9"/>`
	mailInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/><path d="m22 6l-10 7L2 6"/></g>`
	mapInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 6v16l7-4l8 4l7-4V2l-7 4l-8-4l-7 4zm7-4v16m8-12v16"/>`
	mapPinInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"/><circle cx="12" cy="10" r="3"/></g>`
	maximizeInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H5a2 2 0 0 0-2 2v3m18 0V5a2 2 0 0 0-2-2h-3m0 18h3a2 2 0 0 0 2-2v-3M3 16v3a2 2 0 0 0 2 2h3"/>`
	maximizeTwoInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h6v6M9 21H3v-6M21 3l-7 7M3 21l7-7"/>`
	mehInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 15h8M9 9h.01M15 9h.01"/></g>`
	menuInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18M3 6h18M3 18h18"/>`
	messageCircleInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 11.5a8.38 8.38 0 0 1-.9 3.8a8.5 8.5 0 0 1-7.6 4.7a8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8a8.5 8.5 0 0 1 4.7-7.6a8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/>`
	messageSquareInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/>`
	micInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"/><path d="M19 10v2a7 7 0 0 1-14 0v-2m7 9v4m-4 0h8"/></g>`
	micOffInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m1 1l22 22M9 9v3a3 3 0 0 0 5.12 2.12M15 9.34V4a3 3 0 0 0-5.94-.6"/><path d="M17 16.95A7 7 0 0 1 5 12v-2m14 0v2a7 7 0 0 1-.11 1.23M12 19v4m-4 0h8"/></g>`
	minimizeInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3v3a2 2 0 0 1-2 2H3m18 0h-3a2 2 0 0 1-2-2V3m0 18v-3a2 2 0 0 1 2-2h3M3 16h3a2 2 0 0 1 2 2v3"/>`
	minimizeTwoInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14h6v6m10-10h-6V4m0 6l7-7M3 21l7-7"/>`
	minusInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14"/>`
	minusCircleInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 12h8"/></g>`
	minusSquareInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M8 12h8"/></g>`
	monitorInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="3" rx="2" ry="2"/><path d="M8 21h8m-4-4v4"/></g>`
	moonInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12.79A9 9 0 1 1 11.21 3A7 7 0 0 0 21 12.79z"/>`
	moreHorizontalInnerSVG   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></g>`
	moreVerticalInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><circle cx="12" cy="5" r="1"/><circle cx="12" cy="19" r="1"/></g>`
	mousePointerInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l7.07 16.97l2.51-7.39l7.39-2.51L3 3zm10 10l6 6"/>`
	moveInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 9l-3 3l3 3M9 5l3-3l3 3m0 14l-3 3l-3-3M19 9l3 3l-3 3M2 12h20M12 2v20"/>`
	musicInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 18V5l12-2v13"/><circle cx="6" cy="18" r="3"/><circle cx="18" cy="16" r="3"/></g>`
	navigationInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 11l19-9l-9 19l-2-8l-8-2z"/>`
	navigationTwoInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 2l7 19l-7-4l-7 4l7-19z"/>`
	octagonInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2z"/>`
	packageInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16.5 9.4l-9-5.19M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><path d="M3.27 6.96L12 12.01l8.73-5.05M12 22.08V12"/></g>`
	paperclipInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21.44 11.05l-9.19 9.19a6 6 0 0 1-8.49-8.49l9.19-9.19a4 4 0 0 1 5.66 5.66l-9.2 9.19a2 2 0 0 1-2.83-2.83l8.49-8.48"/>`
	pauseInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4h4v16H6zm8 0h4v16h-4z"/>`
	pauseCircleInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M10 15V9m4 6V9"/></g>`
	penToolInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 19l7-7l3 3l-7 7l-3-3z"/><path d="m18 13l-1.5-7.5L2 2l3.5 14.5L13 18l5-5zM2 2l7.586 7.586"/><circle cx="11" cy="11" r="2"/></g>`
	percentInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 5L5 19"/><circle cx="6.5" cy="6.5" r="2.5"/><circle cx="17.5" cy="17.5" r="2.5"/></g>`
	phoneInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 16.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneCallInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.05 5A5 5 0 0 1 19 8.95M15.05 1A9 9 0 0 1 23 8.94m-1 7.98v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneForwardedInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 1l4 4l-4 4m-4-4h8m-1 11.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneIncomingInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 2v6h6m1-7l-7 7m6 8.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneMissedInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m23 1l-6 6m0-6l6 6m-1 9.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneOffInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.68 13.31a16 16 0 0 0 3.41 2.6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7a2 2 0 0 1 1.72 2v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.42 19.42 0 0 1-3.33-2.67m-2.67-3.34a19.79 19.79 0 0 1-3.07-8.63A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91M23 1L1 23"/>`
	phoneOutgoingInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 7V1h-6m-1 7l7-7m-1 15.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	pieChartInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.21 15.89A10 10 0 1 1 8 2.83"/><path d="M22 12A10 10 0 0 0 12 2v10z"/></g>`
	playInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 3l14 9l-14 9V3z"/>`
	playCircleInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m10 8l6 4l-6 4V8z"/></g>`
	plusInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v14m-7-7h14"/>`
	plusCircleInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 8v8m-4-4h8"/></g>`
	plusSquareInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M12 8v8m-4-4h8"/></g>`
	pocketInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3h16a2 2 0 0 1 2 2v6a10 10 0 0 1-10 10A10 10 0 0 1 2 11V5a2 2 0 0 1 2-2z"/><path d="m8 10l4 4l4-4"/></g>`
	powerInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.36 6.64a9 9 0 1 1-12.73 0M12 2v10"/>`
	printerInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 9V2h12v7M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"/><path d="M6 14h12v8H6z"/></g>`
	radioInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="2"/><path d="M16.24 7.76a6 6 0 0 1 0 8.49m-8.48-.01a6 6 0 0 1 0-8.49m11.31-2.82a10 10 0 0 1 0 14.14m-14.14 0a10 10 0 0 1 0-14.14"/></g>`
	refreshCcwInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M1 4v6h6m16 10v-6h-6"/><path d="M20.49 9A9 9 0 0 0 5.64 5.64L1 10m22 4l-4.64 4.36A9 9 0 0 1 3.51 15"/></g>`
	refreshCwInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M23 4v6h-6M1 20v-6h6"/><path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0 0 20.49 15"/></g>`
	repeatInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 1l4 4l-4 4"/><path d="M3 11V9a4 4 0 0 1 4-4h14M7 23l-4-4l4-4"/><path d="M21 13v2a4 4 0 0 1-4 4H3"/></g>`
	rewindInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 19l-9-7l9-7v14zm11 0l-9-7l9-7v14z"/>`
	rotateCcwInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M1 4v6h6"/><path d="M3.51 15a9 9 0 1 0 2.13-9.36L1 10"/></g>`
	rotateCwInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M23 4v6h-6"/><path d="M20.49 15a9 9 0 1 1-2.12-9.36L23 10"/></g>`
	rssInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 11a9 9 0 0 1 9 9M4 4a16 16 0 0 1 16 16"/><circle cx="5" cy="19" r="1"/></g>`
	saveInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/><path d="M17 21v-8H7v8M7 3v5h8"/></g>`
	scissorsInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="6" r="3"/><circle cx="6" cy="18" r="3"/><path d="M20 4L8.12 15.88m6.35-1.4L20 20M8.12 8.12L12 12"/></g>`
	searchInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.35-4.35"/></g>`
	sendInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 2L11 13M22 2l-7 20l-4-9l-9-4l20-7z"/>`
	serverInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="8" x="2" y="2" rx="2" ry="2"/><rect width="20" height="8" x="2" y="14" rx="2" ry="2"/><path d="M6 6h.01M6 18h.01"/></g>`
	settingsInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83a2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33a1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2a2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0a2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2a2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83a2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2a2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51a1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0a2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2a2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/></g>`
	shareInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8m-4-6l-4-4l-4 4m4-4v13"/>`
	shareTwoInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="5" r="3"/><circle cx="6" cy="12" r="3"/><circle cx="18" cy="19" r="3"/><path d="m8.59 13.51l6.83 3.98m-.01-10.98l-6.82 3.98"/></g>`
	shieldInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10z"/>`
	shieldOffInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.69 14a6.9 6.9 0 0 0 .31-2V5l-8-3l-3.16 1.18M4.73 4.73L4 5v7c0 6 8 10 8 10a20.29 20.29 0 0 0 5.62-4.38M1 1l22 22"/>`
	shoppingBagInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4zM3 6h18"/><path d="M16 10a4 4 0 0 1-8 0"/></g>`
	shoppingCartInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"/></g>`
	shuffleInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 3h5v5M4 20L21 3m0 13v5h-5m-1-6l6 6M4 4l5 5"/>`
	sidebarInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 3v18"/></g>`
	skipBackInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20L9 12l10-8v16zM5 19V5"/>`
	skipForwardInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 4l10 8l-10 8V4zm14 1v14"/>`
	slackInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 10c-.83 0-1.5-.67-1.5-1.5v-5c0-.83.67-1.5 1.5-1.5s1.5.67 1.5 1.5v5c0 .83-.67 1.5-1.5 1.5zm6 0H19V8.5c0-.83.67-1.5 1.5-1.5s1.5.67 1.5 1.5s-.67 1.5-1.5 1.5zm-11 4c.83 0 1.5.67 1.5 1.5v5c0 .83-.67 1.5-1.5 1.5S8 21.33 8 20.5v-5c0-.83.67-1.5 1.5-1.5zm-6 0H5v1.5c0 .83-.67 1.5-1.5 1.5S2 16.33 2 15.5S2.67 14 3.5 14zm10.5.5c0-.83.67-1.5 1.5-1.5h5c.83 0 1.5.67 1.5 1.5s-.67 1.5-1.5 1.5h-5c-.83 0-1.5-.67-1.5-1.5zm1.5 4.5H14v1.5c0 .83.67 1.5 1.5 1.5s1.5-.67 1.5-1.5s-.67-1.5-1.5-1.5zM10 9.5C10 8.67 9.33 8 8.5 8h-5C2.67 8 2 8.67 2 9.5S2.67 11 3.5 11h5c.83 0 1.5-.67 1.5-1.5zM8.5 5H10V3.5C10 2.67 9.33 2 8.5 2S7 2.67 7 3.5S7.67 5 8.5 5z"/>`
	slashInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m4.93 4.93l14.14 14.14"/></g>`
	slidersInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 21v-7m0-4V3m8 18v-9m0-4V3m8 18v-5m0-4V3M1 14h6m2-6h6m2 8h6"/>`
	smartphoneInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="2" ry="2"/><path d="M12 18h.01"/></g>`
	smileInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 14s1.5 2 4 2s4-2 4-2M9 9h.01M15 9h.01"/></g>`
	speakerInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><circle cx="12" cy="14" r="4"/><path d="M12 6h.01"/></g>`
	squareInnerSVG           = `<rect width="18" height="18" x="3" y="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2" ry="2"/>`
	starInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 2l3.09 6.26L22 9.27l-5 4.87l1.18 6.88L12 17.77l-6.18 3.25L7 14.14L2 9.27l6.91-1.01L12 2z"/>`
	stopCircleInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9 9h6v6H9z"/></g>`
	sunInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="5"/><path d="M12 1v2m0 18v2M4.22 4.22l1.42 1.42m12.72 12.72l1.42 1.42M1 12h2m18 0h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"/></g>`
	sunriseInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18a5 5 0 0 0-10 0m5-16v7m-7.78 1.22l1.42 1.42M1 18h2m18 0h2m-4.64-6.36l1.42-1.42M23 22H1M8 6l4-4l4 4"/>`
	sunsetInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18a5 5 0 0 0-10 0m5-9V2m-7.78 8.22l1.42 1.42M1 18h2m18 0h2m-4.64-6.36l1.42-1.42M23 22H1M16 5l-4 4l-4-4"/>`
	tabletInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><path d="M12 18h.01"/></g>`
	tagInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82zM7 7h.01"/>`
	targetInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="6"/><circle cx="12" cy="12" r="2"/></g>`
	terminalInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 17l6-6l-6-6m8 14h8"/>`
	thermometerInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 14.76V3.5a2.5 2.5 0 0 0-5 0v11.26a4.5 4.5 0 1 0 5 0z"/>`
	thumbsDownInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 15v4a3 3 0 0 0 3 3l4-9V2H5.72a2 2 0 0 0-2 1.7l-1.38 9a2 2 0 0 0 2 2.3zm7-13h2.67A2.31 2.31 0 0 1 22 4v7a2.31 2.31 0 0 1-2.33 2H17"/>`
	thumbsUpInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"/>`
	toggleLeftInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="22" height="14" x="1" y="5" rx="7" ry="7"/><circle cx="8" cy="12" r="3"/></g>`
	toggleRightInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="22" height="14" x="1" y="5" rx="7" ry="7"/><circle cx="16" cy="12" r="3"/></g>`
	toolInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>`
	trashInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-2 0v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>`
	trashTwoInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-2 0v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2m-6 5v6m4-6v6"/>`
	trelloInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M7 7h3v9H7zm7 0h3v5h-3z"/></g>`
	trendingDownInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m23 18l-9.5-9.5l-5 5L1 6"/><path d="M17 18h6v-6"/></g>`
	trendingUpInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m23 6l-9.5 9.5l-5-5L1 18"/><path d="M17 6h6v6"/></g>`
	triangleInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>`
	truckInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M1 3h15v13H1zm15 5h4l3 3v5h-7V8z"/><circle cx="5.5" cy="18.5" r="2.5"/><circle cx="18.5" cy="18.5" r="2.5"/></g>`
	tvInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="15" x="2" y="7" rx="2" ry="2"/><path d="m17 2l-5 5l-5-5"/></g>`
	twitchInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 2H3v16h5v4l4-4h5l4-4V2zm-10 9V7m5 4V7"/>`
	twitterInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 3a10.9 10.9 0 0 1-3.14 1.53a4.48 4.48 0 0 0-7.86 3v1A10.66 10.66 0 0 1 3 4s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5a4.5 4.5 0 0 0-.08-.83A7.72 7.72 0 0 0 23 3z"/>`
	typeInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7V4h16v3M9 20h6M12 4v16"/>`
	umbrellaInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 12a11.05 11.05 0 0 0-22 0zm-5 7a3 3 0 0 1-6 0v-7"/>`
	underlineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3v7a6 6 0 0 0 6 6a6 6 0 0 0 6-6V3M4 21h16"/>`
	unlockInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 9.9-1"/></g>`
	uploadInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4m14-7l-5-5l-5 5m5-5v12"/>`
	uploadCloudInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 16l-4-4l-4 4m4-4v9"/><path d="M20.39 18.39A5 5 0 0 0 18 9h-1.26A8 8 0 1 0 3 16.3"/><path d="m16 16l-4-4l-4 4"/></g>`
	userInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></g>`
	userCheckInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="8.5" cy="7" r="4"/><path d="m17 11l2 2l4-4"/></g>`
	userMinusInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="8.5" cy="7" r="4"/><path d="M23 11h-6"/></g>`
	userPlusInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="8.5" cy="7" r="4"/><path d="M20 8v6m3-3h-6"/></g>`
	userXInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="8.5" cy="7" r="4"/><path d="m18 8l5 5m0-5l-5 5"/></g>`
	usersInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 0 0-3-3.87m-4-12a4 4 0 0 1 0 7.75"/></g>`
	videoInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m23 7l-7 5l7 5V7z"/><rect width="15" height="14" x="1" y="5" rx="2" ry="2"/></g>`
	videoOffInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 16v1a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h2m5.66 0H14a2 2 0 0 1 2 2v3.34l1 1L23 7v10M1 1l22 22"/>`
	voicemailInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="5.5" cy="11.5" r="4.5"/><circle cx="18.5" cy="11.5" r="4.5"/><path d="M5.5 16h13"/></g>`
	volumeInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5z"/>`
	volumeOneInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm4.54 3.46a5 5 0 0 1 0 7.07"/>`
	volumeTwoInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm8.07-.07a10 10 0 0 1 0 14.14M15.54 8.46a5 5 0 0 1 0 7.07"/>`
	volumeXInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm12 4l-6 6m0-6l6 6"/>`
	watchInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="7"/><path d="M12 9v3l1.5 1.5m3.01 3.85l-.35 3.83a2 2 0 0 1-2 1.82H9.83a2 2 0 0 1-2-1.82l-.35-3.83m.01-10.7l.35-3.83A2 2 0 0 1 9.83 1h4.35a2 2 0 0 1 2 1.82l.35 3.83"/></g>`
	wifiInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12.55a11 11 0 0 1 14.08 0M1.42 9a16 16 0 0 1 21.16 0M8.53 16.11a6 6 0 0 1 6.95 0M12 20h.01"/>`
	wifiOffInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1l22 22m-6.28-11.94A10.94 10.94 0 0 1 19 12.55m-14 0a10.94 10.94 0 0 1 5.17-2.39m.54-5.11A16 16 0 0 1 22.58 9M1.42 9a15.91 15.91 0 0 1 4.7-2.88m2.41 9.99a6 6 0 0 1 6.95 0M12 20h.01"/>`
	windInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.59 4.59A2 2 0 1 1 11 8H2m10.59 11.41A2 2 0 1 0 14 16H2m15.73-8.27A2.5 2.5 0 1 1 19.5 12H2"/>`
	xInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6L6 18M6 6l12 12"/>`
	xCircleInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m15 9l-6 6m0-6l6 6"/></g>`
	xOctagonInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2zM15 9l-6 6m0-6l6 6"/>`
	xSquareInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="m9 9l6 6m0-6l-6 6"/></g>`
	youtubeInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22.54 6.42a2.78 2.78 0 0 0-1.94-2C18.88 4 12 4 12 4s-6.88 0-8.6.46a2.78 2.78 0 0 0-1.94 2A29 29 0 0 0 1 11.75a29 29 0 0 0 .46 5.33A2.78 2.78 0 0 0 3.4 19c1.72.46 8.6.46 8.6.46s6.88 0 8.6-.46a2.78 2.78 0 0 0 1.94-2a29 29 0 0 0 .46-5.25a29 29 0 0 0-.46-5.33z"/><path d="m9.75 15.02l5.75-3.27l-5.75-3.27v6.54z"/></g>`
	zapInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 2L3 14h9l-1 8l10-12h-9l1-8z"/>`
	zapOffInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.41 6.75L13 2l-2.43 2.92m8 7.99L21 10h-5.34M8 8l-5 6h9l-1 8l5-6M1 1l22 22"/>`
	zoomInInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.35-4.35M11 8v6m-3-3h6"/></g>`
	zoomOutInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.35-4.35M8 11h6"/></g>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Activity(children ...*engine.UberElement) *engine.UberElement {
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

func Airplay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplayInnerSVG).
		Element(children...)
}

func AlertCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertCircleInnerSVG).
		Element(children...)
}

func AlertOctagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertOctagonInnerSVG).
		Element(children...)
}

func AlertTriangle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertTriangleInnerSVG).
		Element(children...)
}

func AlignCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterInnerSVG).
		Element(children...)
}

func AlignJustify(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignJustifyInnerSVG).
		Element(children...)
}

func AlignLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignLeftInnerSVG).
		Element(children...)
}

func AlignRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignRightInnerSVG).
		Element(children...)
}

func Anchor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(anchorInnerSVG).
		Element(children...)
}

func Aperture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(apertureInnerSVG).
		Element(children...)
}

func Archive(children ...*engine.UberElement) *engine.UberElement {
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

func ArrowDown(children ...*engine.UberElement) *engine.UberElement {
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

func ArrowDownCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownCircleInnerSVG).
		Element(children...)
}

func ArrowDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownLeftInnerSVG).
		Element(children...)
}

func ArrowDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownRightInnerSVG).
		Element(children...)
}

func ArrowLeft(children ...*engine.UberElement) *engine.UberElement {
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

func ArrowLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftCircleInnerSVG).
		Element(children...)
}

func ArrowRight(children ...*engine.UberElement) *engine.UberElement {
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

func ArrowRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightCircleInnerSVG).
		Element(children...)
}

func ArrowUp(children ...*engine.UberElement) *engine.UberElement {
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

func ArrowUpCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpCircleInnerSVG).
		Element(children...)
}

func ArrowUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpLeftInnerSVG).
		Element(children...)
}

func ArrowUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpRightInnerSVG).
		Element(children...)
}

func AtSign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atSignInnerSVG).
		Element(children...)
}

func Award(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(awardInnerSVG).
		Element(children...)
}

func BarChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartInnerSVG).
		Element(children...)
}

func BarChartTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartTwoInnerSVG).
		Element(children...)
}

func Battery(children ...*engine.UberElement) *engine.UberElement {
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

func BatteryCharging(children ...*engine.UberElement) *engine.UberElement {
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

func Bell(children ...*engine.UberElement) *engine.UberElement {
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

func BellOff(children ...*engine.UberElement) *engine.UberElement {
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

func Bluetooth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothInnerSVG).
		Element(children...)
}

func Bold(children ...*engine.UberElement) *engine.UberElement {
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

func Book(children ...*engine.UberElement) *engine.UberElement {
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

func BookOpen(children ...*engine.UberElement) *engine.UberElement {
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

func Bookmark(children ...*engine.UberElement) *engine.UberElement {
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

func Box(children ...*engine.UberElement) *engine.UberElement {
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

func Briefcase(children ...*engine.UberElement) *engine.UberElement {
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

func Calendar(children ...*engine.UberElement) *engine.UberElement {
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

func Camera(children ...*engine.UberElement) *engine.UberElement {
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

func CameraOff(children ...*engine.UberElement) *engine.UberElement {
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

func Cast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(castInnerSVG).
		Element(children...)
}

func Check(children ...*engine.UberElement) *engine.UberElement {
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

func CheckCircle(children ...*engine.UberElement) *engine.UberElement {
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

func CheckSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkSquareInnerSVG).
		Element(children...)
}

func ChevronDown(children ...*engine.UberElement) *engine.UberElement {
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

func ChevronLeft(children ...*engine.UberElement) *engine.UberElement {
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

func ChevronRight(children ...*engine.UberElement) *engine.UberElement {
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

func ChevronUp(children ...*engine.UberElement) *engine.UberElement {
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

func ChevronsDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsDownInnerSVG).
		Element(children...)
}

func ChevronsLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsLeftInnerSVG).
		Element(children...)
}

func ChevronsRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsRightInnerSVG).
		Element(children...)
}

func ChevronsUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsUpInnerSVG).
		Element(children...)
}

func Chrome(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chromeInnerSVG).
		Element(children...)
}

func Circle(children ...*engine.UberElement) *engine.UberElement {
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

func Clipboard(children ...*engine.UberElement) *engine.UberElement {
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

func Clock(children ...*engine.UberElement) *engine.UberElement {
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

func Cloud(children ...*engine.UberElement) *engine.UberElement {
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

func CloudDrizzle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDrizzleInnerSVG).
		Element(children...)
}

func CloudLightning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudLightningInnerSVG).
		Element(children...)
}

func CloudOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudOffInnerSVG).
		Element(children...)
}

func CloudRain(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudRainInnerSVG).
		Element(children...)
}

func CloudSnow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudSnowInnerSVG).
		Element(children...)
}

func Code(children ...*engine.UberElement) *engine.UberElement {
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

func Codepen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codepenInnerSVG).
		Element(children...)
}

func Codesandbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codesandboxInnerSVG).
		Element(children...)
}

func Coffee(children ...*engine.UberElement) *engine.UberElement {
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

func Columns(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(columnsInnerSVG).
		Element(children...)
}

func Command(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commandInnerSVG).
		Element(children...)
}

func Compass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassInnerSVG).
		Element(children...)
}

func Copy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(copyInnerSVG).
		Element(children...)
}

func CornerDownLeft(children ...*engine.UberElement) *engine.UberElement {
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

func CornerDownRight(children ...*engine.UberElement) *engine.UberElement {
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

func CornerLeftDown(children ...*engine.UberElement) *engine.UberElement {
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

func CornerLeftUp(children ...*engine.UberElement) *engine.UberElement {
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

func CornerRightDown(children ...*engine.UberElement) *engine.UberElement {
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

func CornerRightUp(children ...*engine.UberElement) *engine.UberElement {
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

func CornerUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerUpLeftInnerSVG).
		Element(children...)
}

func CornerUpRight(children ...*engine.UberElement) *engine.UberElement {
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

func Cpu(children ...*engine.UberElement) *engine.UberElement {
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

func CreditCard(children ...*engine.UberElement) *engine.UberElement {
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

func Crop(children ...*engine.UberElement) *engine.UberElement {
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

func Crosshair(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crosshairInnerSVG).
		Element(children...)
}

func Database(children ...*engine.UberElement) *engine.UberElement {
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

func Delete(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deleteInnerSVG).
		Element(children...)
}

func Disc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discInnerSVG).
		Element(children...)
}

func Divide(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(divideInnerSVG).
		Element(children...)
}

func DivideCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(divideCircleInnerSVG).
		Element(children...)
}

func DivideSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(divideSquareInnerSVG).
		Element(children...)
}

func DollarSign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dollarSignInnerSVG).
		Element(children...)
}

func Download(children ...*engine.UberElement) *engine.UberElement {
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

func DownloadCloud(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadCloudInnerSVG).
		Element(children...)
}

func Dribbble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dribbbleInnerSVG).
		Element(children...)
}

func Droplet(children ...*engine.UberElement) *engine.UberElement {
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

func Edit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editInnerSVG).
		Element(children...)
}

func EditThree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editThreeInnerSVG).
		Element(children...)
}

func EditTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editTwoInnerSVG).
		Element(children...)
}

func ExternalLink(children ...*engine.UberElement) *engine.UberElement {
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

func Eye(children ...*engine.UberElement) *engine.UberElement {
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

func EyeOff(children ...*engine.UberElement) *engine.UberElement {
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

func Facebook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookInnerSVG).
		Element(children...)
}

func FastForward(children ...*engine.UberElement) *engine.UberElement {
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

func Feather(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(featherInnerSVG).
		Element(children...)
}

func Figma(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(figmaInnerSVG).
		Element(children...)
}

func File(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileInnerSVG).
		Element(children...)
}

func FileMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileMinusInnerSVG).
		Element(children...)
}

func FilePlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filePlusInnerSVG).
		Element(children...)
}

func FileText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTextInnerSVG).
		Element(children...)
}

func Film(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filmInnerSVG).
		Element(children...)
}

func Filter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterInnerSVG).
		Element(children...)
}

func Flag(children ...*engine.UberElement) *engine.UberElement {
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

func Folder(children ...*engine.UberElement) *engine.UberElement {
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

func FolderMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderMinusInnerSVG).
		Element(children...)
}

func FolderPlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderPlusInnerSVG).
		Element(children...)
}

func Framer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(framerInnerSVG).
		Element(children...)
}

func Frown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frownInnerSVG).
		Element(children...)
}

func Gift(children ...*engine.UberElement) *engine.UberElement {
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

func GitBranch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitBranchInnerSVG).
		Element(children...)
}

func GitCommit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCommitInnerSVG).
		Element(children...)
}

func GitMerge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitMergeInnerSVG).
		Element(children...)
}

func GitPullRequest(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitPullRequestInnerSVG).
		Element(children...)
}

func Github(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(githubInnerSVG).
		Element(children...)
}

func Gitlab(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitlabInnerSVG).
		Element(children...)
}

func Globe(children ...*engine.UberElement) *engine.UberElement {
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

func Grid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridInnerSVG).
		Element(children...)
}

func HardDrive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hardDriveInnerSVG).
		Element(children...)
}

func Hash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hashInnerSVG).
		Element(children...)
}

func Headphones(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headphonesInnerSVG).
		Element(children...)
}

func Heart(children ...*engine.UberElement) *engine.UberElement {
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

func HelpCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(helpCircleInnerSVG).
		Element(children...)
}

func Hexagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hexagonInnerSVG).
		Element(children...)
}

func Home(children ...*engine.UberElement) *engine.UberElement {
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

func Image(children ...*engine.UberElement) *engine.UberElement {
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

func Inbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxInnerSVG).
		Element(children...)
}

func Info(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoInnerSVG).
		Element(children...)
}

func Instagram(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(instagramInnerSVG).
		Element(children...)
}

func Italic(children ...*engine.UberElement) *engine.UberElement {
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

func Key(children ...*engine.UberElement) *engine.UberElement {
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

func Layers(children ...*engine.UberElement) *engine.UberElement {
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

func Layout(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutInnerSVG).
		Element(children...)
}

func LifeBuoy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lifeBuoyInnerSVG).
		Element(children...)
}

func Link(children ...*engine.UberElement) *engine.UberElement {
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

func LinkTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkTwoInnerSVG).
		Element(children...)
}

func Linkedin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkedinInnerSVG).
		Element(children...)
}

func List(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listInnerSVG).
		Element(children...)
}

func Loader(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loaderInnerSVG).
		Element(children...)
}

func Lock(children ...*engine.UberElement) *engine.UberElement {
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

func LogIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(logInInnerSVG).
		Element(children...)
}

func LogOut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(logOutInnerSVG).
		Element(children...)
}

func Mail(children ...*engine.UberElement) *engine.UberElement {
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

func Map(children ...*engine.UberElement) *engine.UberElement {
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

func MapPin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapPinInnerSVG).
		Element(children...)
}

func Maximize(children ...*engine.UberElement) *engine.UberElement {
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

func MaximizeTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maximizeTwoInnerSVG).
		Element(children...)
}

func Meh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mehInnerSVG).
		Element(children...)
}

func Menu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuInnerSVG).
		Element(children...)
}

func MessageCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageCircleInnerSVG).
		Element(children...)
}

func MessageSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageSquareInnerSVG).
		Element(children...)
}

func Mic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(micInnerSVG).
		Element(children...)
}

func MicOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(micOffInnerSVG).
		Element(children...)
}

func Minimize(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minimizeInnerSVG).
		Element(children...)
}

func MinimizeTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minimizeTwoInnerSVG).
		Element(children...)
}

func Minus(children ...*engine.UberElement) *engine.UberElement {
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

func MinusCircle(children ...*engine.UberElement) *engine.UberElement {
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

func MinusSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusSquareInnerSVG).
		Element(children...)
}

func Monitor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(monitorInnerSVG).
		Element(children...)
}

func Moon(children ...*engine.UberElement) *engine.UberElement {
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

func MoreHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreHorizontalInnerSVG).
		Element(children...)
}

func MoreVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreVerticalInnerSVG).
		Element(children...)
}

func MousePointer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mousePointerInnerSVG).
		Element(children...)
}

func Move(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveInnerSVG).
		Element(children...)
}

func Music(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(musicInnerSVG).
		Element(children...)
}

func Navigation(children ...*engine.UberElement) *engine.UberElement {
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

func NavigationTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navigationTwoInnerSVG).
		Element(children...)
}

func Octagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(octagonInnerSVG).
		Element(children...)
}

func Package(children ...*engine.UberElement) *engine.UberElement {
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

func Paperclip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperclipInnerSVG).
		Element(children...)
}

func Pause(children ...*engine.UberElement) *engine.UberElement {
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

func PauseCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseCircleInnerSVG).
		Element(children...)
}

func PenTool(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penToolInnerSVG).
		Element(children...)
}

func Percent(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(percentInnerSVG).
		Element(children...)
}

func Phone(children ...*engine.UberElement) *engine.UberElement {
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

func PhoneCall(children ...*engine.UberElement) *engine.UberElement {
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

func PhoneForwarded(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneForwardedInnerSVG).
		Element(children...)
}

func PhoneIncoming(children ...*engine.UberElement) *engine.UberElement {
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

func PhoneMissed(children ...*engine.UberElement) *engine.UberElement {
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

func PhoneOff(children ...*engine.UberElement) *engine.UberElement {
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

func PhoneOutgoing(children ...*engine.UberElement) *engine.UberElement {
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

func PieChart(children ...*engine.UberElement) *engine.UberElement {
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

func Play(children ...*engine.UberElement) *engine.UberElement {
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

func PlayCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playCircleInnerSVG).
		Element(children...)
}

func Plus(children ...*engine.UberElement) *engine.UberElement {
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

func PlusCircle(children ...*engine.UberElement) *engine.UberElement {
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

func PlusSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusSquareInnerSVG).
		Element(children...)
}

func Pocket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pocketInnerSVG).
		Element(children...)
}

func Power(children ...*engine.UberElement) *engine.UberElement {
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

func Printer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printerInnerSVG).
		Element(children...)
}

func Radio(children ...*engine.UberElement) *engine.UberElement {
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

func RefreshCcw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshCcwInnerSVG).
		Element(children...)
}

func RefreshCw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshCwInnerSVG).
		Element(children...)
}

func Repeat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(repeatInnerSVG).
		Element(children...)
}

func Rewind(children ...*engine.UberElement) *engine.UberElement {
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

func RotateCcw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateCcwInnerSVG).
		Element(children...)
}

func RotateCw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateCwInnerSVG).
		Element(children...)
}

func Rss(children ...*engine.UberElement) *engine.UberElement {
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

func Save(children ...*engine.UberElement) *engine.UberElement {
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

func Scissors(children ...*engine.UberElement) *engine.UberElement {
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

func Search(children ...*engine.UberElement) *engine.UberElement {
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

func Send(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendInnerSVG).
		Element(children...)
}

func Server(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(serverInnerSVG).
		Element(children...)
}

func Settings(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(settingsInnerSVG).
		Element(children...)
}

func Share(children ...*engine.UberElement) *engine.UberElement {
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

func ShareTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareTwoInnerSVG).
		Element(children...)
}

func Shield(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldInnerSVG).
		Element(children...)
}

func ShieldOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldOffInnerSVG).
		Element(children...)
}

func ShoppingBag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagInnerSVG).
		Element(children...)
}

func ShoppingCart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingCartInnerSVG).
		Element(children...)
}

func Shuffle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shuffleInnerSVG).
		Element(children...)
}

func Sidebar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sidebarInnerSVG).
		Element(children...)
}

func SkipBack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skipBackInnerSVG).
		Element(children...)
}

func SkipForward(children ...*engine.UberElement) *engine.UberElement {
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

func Slack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slackInnerSVG).
		Element(children...)
}

func Slash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slashInnerSVG).
		Element(children...)
}

func Sliders(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slidersInnerSVG).
		Element(children...)
}

func Smartphone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smartphoneInnerSVG).
		Element(children...)
}

func Smile(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smileInnerSVG).
		Element(children...)
}

func Speaker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speakerInnerSVG).
		Element(children...)
}

func Square(children ...*engine.UberElement) *engine.UberElement {
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

func Star(children ...*engine.UberElement) *engine.UberElement {
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

func StopCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopCircleInnerSVG).
		Element(children...)
}

func Sun(children ...*engine.UberElement) *engine.UberElement {
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

func Sunrise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunriseInnerSVG).
		Element(children...)
}

func Sunset(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunsetInnerSVG).
		Element(children...)
}

func Tablet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tabletInnerSVG).
		Element(children...)
}

func Tag(children ...*engine.UberElement) *engine.UberElement {
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

func Target(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(targetInnerSVG).
		Element(children...)
}

func Terminal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalInnerSVG).
		Element(children...)
}

func Thermometer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thermometerInnerSVG).
		Element(children...)
}

func ThumbsDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbsDownInnerSVG).
		Element(children...)
}

func ThumbsUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbsUpInnerSVG).
		Element(children...)
}

func ToggleLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleLeftInnerSVG).
		Element(children...)
}

func ToggleRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleRightInnerSVG).
		Element(children...)
}

func Tool(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toolInnerSVG).
		Element(children...)
}

func Trash(children ...*engine.UberElement) *engine.UberElement {
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

func TrashTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashTwoInnerSVG).
		Element(children...)
}

func Trello(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trelloInnerSVG).
		Element(children...)
}

func TrendingDown(children ...*engine.UberElement) *engine.UberElement {
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

func TrendingUp(children ...*engine.UberElement) *engine.UberElement {
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

func Triangle(children ...*engine.UberElement) *engine.UberElement {
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

func Truck(children ...*engine.UberElement) *engine.UberElement {
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

func Tv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tvInnerSVG).
		Element(children...)
}

func Twitch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitchInnerSVG).
		Element(children...)
}

func Twitter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitterInnerSVG).
		Element(children...)
}

func Type(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(typeInnerSVG).
		Element(children...)
}

func Umbrella(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(umbrellaInnerSVG).
		Element(children...)
}

func Underline(children ...*engine.UberElement) *engine.UberElement {
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

func Unlock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockInnerSVG).
		Element(children...)
}

func Upload(children ...*engine.UberElement) *engine.UberElement {
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

func UploadCloud(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadCloudInnerSVG).
		Element(children...)
}

func User(children ...*engine.UberElement) *engine.UberElement {
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

func UserCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCheckInnerSVG).
		Element(children...)
}

func UserMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userMinusInnerSVG).
		Element(children...)
}

func UserPlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userPlusInnerSVG).
		Element(children...)
}

func UserX(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userXInnerSVG).
		Element(children...)
}

func Users(children ...*engine.UberElement) *engine.UberElement {
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

func Video(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(videoInnerSVG).
		Element(children...)
}

func VideoOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(videoOffInnerSVG).
		Element(children...)
}

func Voicemail(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voicemailInnerSVG).
		Element(children...)
}

func Volume(children ...*engine.UberElement) *engine.UberElement {
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

func VolumeOne(children ...*engine.UberElement) *engine.UberElement {
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

func VolumeTwo(children ...*engine.UberElement) *engine.UberElement {
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

func VolumeX(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeXInnerSVG).
		Element(children...)
}

func Watch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(watchInnerSVG).
		Element(children...)
}

func Wifi(children ...*engine.UberElement) *engine.UberElement {
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

func WifiOff(children ...*engine.UberElement) *engine.UberElement {
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

func Wind(children ...*engine.UberElement) *engine.UberElement {
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

func X(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xInnerSVG).
		Element(children...)
}

func XCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xCircleInnerSVG).
		Element(children...)
}

func XOctagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xOctagonInnerSVG).
		Element(children...)
}

func XSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xSquareInnerSVG).
		Element(children...)
}

func Youtube(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(youtubeInnerSVG).
		Element(children...)
}

func Zap(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zapInnerSVG).
		Element(children...)
}

func ZapOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zapOffInnerSVG).
		Element(children...)
}

func ZoomIn(children ...*engine.UberElement) *engine.UberElement {
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

func ZoomOut(children ...*engine.UberElement) *engine.UberElement {
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
	case "airplay":
		return Airplay(), nil
	case "alert-circle":
		return AlertCircle(), nil
	case "alert-octagon":
		return AlertOctagon(), nil
	case "alert-triangle":
		return AlertTriangle(), nil
	case "align-center":
		return AlignCenter(), nil
	case "align-justify":
		return AlignJustify(), nil
	case "align-left":
		return AlignLeft(), nil
	case "align-right":
		return AlignRight(), nil
	case "anchor":
		return Anchor(), nil
	case "aperture":
		return Aperture(), nil
	case "archive":
		return Archive(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-down-circle":
		return ArrowDownCircle(), nil
	case "arrow-down-left":
		return ArrowDownLeft(), nil
	case "arrow-down-right":
		return ArrowDownRight(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-left-circle":
		return ArrowLeftCircle(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-right-circle":
		return ArrowRightCircle(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrow-up-circle":
		return ArrowUpCircle(), nil
	case "arrow-up-left":
		return ArrowUpLeft(), nil
	case "arrow-up-right":
		return ArrowUpRight(), nil
	case "at-sign":
		return AtSign(), nil
	case "award":
		return Award(), nil
	case "bar-chart":
		return BarChart(), nil
	case "bar-chart-2":
		return BarChartTwo(), nil
	case "battery":
		return Battery(), nil
	case "battery-charging":
		return BatteryCharging(), nil
	case "bell":
		return Bell(), nil
	case "bell-off":
		return BellOff(), nil
	case "bluetooth":
		return Bluetooth(), nil
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
	case "briefcase":
		return Briefcase(), nil
	case "calendar":
		return Calendar(), nil
	case "camera":
		return Camera(), nil
	case "camera-off":
		return CameraOff(), nil
	case "cast":
		return Cast(), nil
	case "check":
		return Check(), nil
	case "check-circle":
		return CheckCircle(), nil
	case "check-square":
		return CheckSquare(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "chevrons-down":
		return ChevronsDown(), nil
	case "chevrons-left":
		return ChevronsLeft(), nil
	case "chevrons-right":
		return ChevronsRight(), nil
	case "chevrons-up":
		return ChevronsUp(), nil
	case "chrome":
		return Chrome(), nil
	case "circle":
		return Circle(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clock":
		return Clock(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-drizzle":
		return CloudDrizzle(), nil
	case "cloud-lightning":
		return CloudLightning(), nil
	case "cloud-off":
		return CloudOff(), nil
	case "cloud-rain":
		return CloudRain(), nil
	case "cloud-snow":
		return CloudSnow(), nil
	case "code":
		return Code(), nil
	case "codepen":
		return Codepen(), nil
	case "codesandbox":
		return Codesandbox(), nil
	case "coffee":
		return Coffee(), nil
	case "columns":
		return Columns(), nil
	case "command":
		return Command(), nil
	case "compass":
		return Compass(), nil
	case "copy":
		return Copy(), nil
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
	case "corner-up-left":
		return CornerUpLeft(), nil
	case "corner-up-right":
		return CornerUpRight(), nil
	case "cpu":
		return Cpu(), nil
	case "credit-card":
		return CreditCard(), nil
	case "crop":
		return Crop(), nil
	case "crosshair":
		return Crosshair(), nil
	case "database":
		return Database(), nil
	case "delete":
		return Delete(), nil
	case "disc":
		return Disc(), nil
	case "divide":
		return Divide(), nil
	case "divide-circle":
		return DivideCircle(), nil
	case "divide-square":
		return DivideSquare(), nil
	case "dollar-sign":
		return DollarSign(), nil
	case "download":
		return Download(), nil
	case "download-cloud":
		return DownloadCloud(), nil
	case "dribbble":
		return Dribbble(), nil
	case "droplet":
		return Droplet(), nil
	case "edit":
		return Edit(), nil
	case "edit-3":
		return EditThree(), nil
	case "edit-2":
		return EditTwo(), nil
	case "external-link":
		return ExternalLink(), nil
	case "eye":
		return Eye(), nil
	case "eye-off":
		return EyeOff(), nil
	case "facebook":
		return Facebook(), nil
	case "fast-forward":
		return FastForward(), nil
	case "feather":
		return Feather(), nil
	case "figma":
		return Figma(), nil
	case "file":
		return File(), nil
	case "file-minus":
		return FileMinus(), nil
	case "file-plus":
		return FilePlus(), nil
	case "file-text":
		return FileText(), nil
	case "film":
		return Film(), nil
	case "filter":
		return Filter(), nil
	case "flag":
		return Flag(), nil
	case "folder":
		return Folder(), nil
	case "folder-minus":
		return FolderMinus(), nil
	case "folder-plus":
		return FolderPlus(), nil
	case "framer":
		return Framer(), nil
	case "frown":
		return Frown(), nil
	case "gift":
		return Gift(), nil
	case "git-branch":
		return GitBranch(), nil
	case "git-commit":
		return GitCommit(), nil
	case "git-merge":
		return GitMerge(), nil
	case "git-pull-request":
		return GitPullRequest(), nil
	case "github":
		return Github(), nil
	case "gitlab":
		return Gitlab(), nil
	case "globe":
		return Globe(), nil
	case "grid":
		return Grid(), nil
	case "hard-drive":
		return HardDrive(), nil
	case "hash":
		return Hash(), nil
	case "headphones":
		return Headphones(), nil
	case "heart":
		return Heart(), nil
	case "help-circle":
		return HelpCircle(), nil
	case "hexagon":
		return Hexagon(), nil
	case "home":
		return Home(), nil
	case "image":
		return Image(), nil
	case "inbox":
		return Inbox(), nil
	case "info":
		return Info(), nil
	case "instagram":
		return Instagram(), nil
	case "italic":
		return Italic(), nil
	case "key":
		return Key(), nil
	case "layers":
		return Layers(), nil
	case "layout":
		return Layout(), nil
	case "life-buoy":
		return LifeBuoy(), nil
	case "link":
		return Link(), nil
	case "link-2":
		return LinkTwo(), nil
	case "linkedin":
		return Linkedin(), nil
	case "list":
		return List(), nil
	case "loader":
		return Loader(), nil
	case "lock":
		return Lock(), nil
	case "log-in":
		return LogIn(), nil
	case "log-out":
		return LogOut(), nil
	case "mail":
		return Mail(), nil
	case "map":
		return Map(), nil
	case "map-pin":
		return MapPin(), nil
	case "maximize":
		return Maximize(), nil
	case "maximize-2":
		return MaximizeTwo(), nil
	case "meh":
		return Meh(), nil
	case "menu":
		return Menu(), nil
	case "message-circle":
		return MessageCircle(), nil
	case "message-square":
		return MessageSquare(), nil
	case "mic":
		return Mic(), nil
	case "mic-off":
		return MicOff(), nil
	case "minimize":
		return Minimize(), nil
	case "minimize-2":
		return MinimizeTwo(), nil
	case "minus":
		return Minus(), nil
	case "minus-circle":
		return MinusCircle(), nil
	case "minus-square":
		return MinusSquare(), nil
	case "monitor":
		return Monitor(), nil
	case "moon":
		return Moon(), nil
	case "more-horizontal":
		return MoreHorizontal(), nil
	case "more-vertical":
		return MoreVertical(), nil
	case "mouse-pointer":
		return MousePointer(), nil
	case "move":
		return Move(), nil
	case "music":
		return Music(), nil
	case "navigation":
		return Navigation(), nil
	case "navigation-2":
		return NavigationTwo(), nil
	case "octagon":
		return Octagon(), nil
	case "package":
		return Package(), nil
	case "paperclip":
		return Paperclip(), nil
	case "pause":
		return Pause(), nil
	case "pause-circle":
		return PauseCircle(), nil
	case "pen-tool":
		return PenTool(), nil
	case "percent":
		return Percent(), nil
	case "phone":
		return Phone(), nil
	case "phone-call":
		return PhoneCall(), nil
	case "phone-forwarded":
		return PhoneForwarded(), nil
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
	case "play-circle":
		return PlayCircle(), nil
	case "plus":
		return Plus(), nil
	case "plus-circle":
		return PlusCircle(), nil
	case "plus-square":
		return PlusSquare(), nil
	case "pocket":
		return Pocket(), nil
	case "power":
		return Power(), nil
	case "printer":
		return Printer(), nil
	case "radio":
		return Radio(), nil
	case "refresh-ccw":
		return RefreshCcw(), nil
	case "refresh-cw":
		return RefreshCw(), nil
	case "repeat":
		return Repeat(), nil
	case "rewind":
		return Rewind(), nil
	case "rotate-ccw":
		return RotateCcw(), nil
	case "rotate-cw":
		return RotateCw(), nil
	case "rss":
		return Rss(), nil
	case "save":
		return Save(), nil
	case "scissors":
		return Scissors(), nil
	case "search":
		return Search(), nil
	case "send":
		return Send(), nil
	case "server":
		return Server(), nil
	case "settings":
		return Settings(), nil
	case "share":
		return Share(), nil
	case "share-2":
		return ShareTwo(), nil
	case "shield":
		return Shield(), nil
	case "shield-off":
		return ShieldOff(), nil
	case "shopping-bag":
		return ShoppingBag(), nil
	case "shopping-cart":
		return ShoppingCart(), nil
	case "shuffle":
		return Shuffle(), nil
	case "sidebar":
		return Sidebar(), nil
	case "skip-back":
		return SkipBack(), nil
	case "skip-forward":
		return SkipForward(), nil
	case "slack":
		return Slack(), nil
	case "slash":
		return Slash(), nil
	case "sliders":
		return Sliders(), nil
	case "smartphone":
		return Smartphone(), nil
	case "smile":
		return Smile(), nil
	case "speaker":
		return Speaker(), nil
	case "square":
		return Square(), nil
	case "star":
		return Star(), nil
	case "stop-circle":
		return StopCircle(), nil
	case "sun":
		return Sun(), nil
	case "sunrise":
		return Sunrise(), nil
	case "sunset":
		return Sunset(), nil
	case "tablet":
		return Tablet(), nil
	case "tag":
		return Tag(), nil
	case "target":
		return Target(), nil
	case "terminal":
		return Terminal(), nil
	case "thermometer":
		return Thermometer(), nil
	case "thumbs-down":
		return ThumbsDown(), nil
	case "thumbs-up":
		return ThumbsUp(), nil
	case "toggle-left":
		return ToggleLeft(), nil
	case "toggle-right":
		return ToggleRight(), nil
	case "tool":
		return Tool(), nil
	case "trash":
		return Trash(), nil
	case "trash-2":
		return TrashTwo(), nil
	case "trello":
		return Trello(), nil
	case "trending-down":
		return TrendingDown(), nil
	case "trending-up":
		return TrendingUp(), nil
	case "triangle":
		return Triangle(), nil
	case "truck":
		return Truck(), nil
	case "tv":
		return Tv(), nil
	case "twitch":
		return Twitch(), nil
	case "twitter":
		return Twitter(), nil
	case "type":
		return Type(), nil
	case "umbrella":
		return Umbrella(), nil
	case "underline":
		return Underline(), nil
	case "unlock":
		return Unlock(), nil
	case "upload":
		return Upload(), nil
	case "upload-cloud":
		return UploadCloud(), nil
	case "user":
		return User(), nil
	case "user-check":
		return UserCheck(), nil
	case "user-minus":
		return UserMinus(), nil
	case "user-plus":
		return UserPlus(), nil
	case "user-x":
		return UserX(), nil
	case "users":
		return Users(), nil
	case "video":
		return Video(), nil
	case "video-off":
		return VideoOff(), nil
	case "voicemail":
		return Voicemail(), nil
	case "volume":
		return Volume(), nil
	case "volume-1":
		return VolumeOne(), nil
	case "volume-2":
		return VolumeTwo(), nil
	case "volume-x":
		return VolumeX(), nil
	case "watch":
		return Watch(), nil
	case "wifi":
		return Wifi(), nil
	case "wifi-off":
		return WifiOff(), nil
	case "wind":
		return Wind(), nil
	case "x":
		return X(), nil
	case "x-circle":
		return XCircle(), nil
	case "x-octagon":
		return XOctagon(), nil
	case "x-square":
		return XSquare(), nil
	case "youtube":
		return Youtube(), nil
	case "zap":
		return Zap(), nil
	case "zap-off":
		return ZapOff(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in feather icon set", name)
	}
}
