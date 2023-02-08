package oi

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	accountLoginInnerSVG          = `<path fill="currentColor" d="M3 0v1h4v5H3v1h5V0H3zm1 2v1H0v1h4v1l2-1.5L4 2z"/>`
	accountLogoutInnerSVG         = `<path fill="currentColor" d="M3 0v1h4v5H3v1h5V0H3zM2 2L0 3.5L2 5V4h4V3H2V2z"/>`
	actionRedoInnerSVG            = `<path fill="currentColor" d="M3.5 1C1.57 1 0 2.57 0 4.5a2.5 2.5 0 0 1 5 0V5H4l2 2l2-2H7v-.5C7 2.57 5.43 1 3.5 1z"/>`
	actionUndoInnerSVG            = `<path fill="currentColor" d="M4.5 1C2.57 1 1 2.57 1 4.5V5H0l2 2l2-2H3v-.5a2.5 2.5 0 0 1 5 0C8 2.57 6.43 1 4.5 1z"/>`
	alignCenterInnerSVG           = `<path fill="currentColor" d="M0 0v1h8V0H0zm1 2v1h6V2H1zM0 4v1h8V4H0zm1 2v1h6V6H1z"/>`
	alignLeftInnerSVG             = `<path fill="currentColor" d="M0 0v1h8V0H0zm0 2v1h6V2H0zm0 2v1h8V4H0zm0 2v1h6V6H0z"/>`
	alignRightInnerSVG            = `<path fill="currentColor" d="M0 0v1h8V0H0zm2 2v1h6V2H2zM0 4v1h8V4H0zm2 2v1h6V6H2z"/>`
	apertureInnerSVG              = `<path fill="currentColor" d="M4 0c-.69 0-1.34.19-1.91.5l3.22 2.34l.75-2.25C5.46.23 4.75 0 4 0zM1.25 1.13C.49 1.86 0 2.87 0 4.01c0 .25.02.48.06.72l3.09-2.22l-1.91-1.38zm5.63.13L5.66 5.01h2.19c.08-.32.16-.65.16-1c0-1.07-.44-2.03-1.13-2.75zM2.16 4.48L.41 5.73c.55 1.13 1.6 1.99 2.88 2.22L2.16 4.48zm1.56 1.53l.63 1.97c1.33-.12 2.46-.88 3.09-1.97H3.72z"/>`
	arrowBottomInnerSVG           = `<path fill="currentColor" d="M3 0v5H1l2.53 3L6 5H4V0H3z"/>`
	arrowCircleBottomInnerSVG     = `<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zM3 1h2v3h2L4 7L1 4h2V1z"/>`
	arrowCircleLeftInnerSVG       = `<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zm0 1v2h3v2H4v2L1 4l3-3z"/>`
	arrowCircleRightInnerSVG      = `<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zm0 1l3 3l-3 3V5H1V3h3V1z"/>`
	arrowCircleTopInnerSVG        = `<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zm0 1l3 3H5v3H3V4H1l3-3z"/>`
	arrowLeftInnerSVG             = `<path fill="currentColor" d="M3 1L0 3.53L3 6V4h5V3H3V1z"/>`
	arrowRightInnerSVG            = `<path fill="currentColor" d="M5 1v2H0v1h5v2l3-2.53L5 1z"/>`
	arrowThickBottomInnerSVG      = `<path fill="currentColor" d="M3 0v5H1l3.03 3L7 5H5V0H3z"/>`
	arrowThickLeftInnerSVG        = `<path fill="currentColor" d="M3 1L0 4.03L3 7V5h5V3H3V1z"/>`
	arrowThickRightInnerSVG       = `<path fill="currentColor" d="M5 1v2H0v2h5v2l3-3.03L5 1z"/>`
	arrowThickTopInnerSVG         = `<path fill="currentColor" d="M3.97 0L1 3h2v5h2V3h2L3.97 0z"/>`
	arrowTopInnerSVG              = `<path fill="currentColor" d="M3.47 0L1 3h2v5h1V3h2L3.47 0z"/>`
	audioInnerSVG                 = `<path fill="currentColor" d="M1.16 1C.44 1.72 0 2.71 0 3.81s.43 2.12 1.16 2.84l.72-.72C1.34 5.39 1 4.64 1 3.8c0-.83.33-1.55.88-2.09L1.16.99zm5.69 0l-.72.72c.54.54.88 1.26.88 2.09c0 .83-.33 1.58-.88 2.13l.72.72c.72-.72 1.16-1.74 1.16-2.84c0-1.1-.43-2.09-1.16-2.81zM2.6 2.41c-.36.36-.59.86-.59 1.41c0 .55.23 1.08.59 1.44l.69-.72c-.18-.18-.28-.44-.28-.72c0-.28.1-.5.28-.69l-.69-.72zm2.81 0l-.69.72c.18.18.28.41.28.69c0 .28-.1.54-.28.72l.69.72c.36-.36.59-.89.59-1.44c0-.55-.23-1.05-.59-1.41z"/>`
	audioSpectrumInnerSVG         = `<path fill="currentColor" d="M4 0v8h1V0H4zM2 1v6h1V1H2zm4 1v4h1V2H6zM0 3v2h1V3H0z"/>`
	badgeInnerSVG                 = `<path fill="currentColor" d="M4 0C2.9 0 2 .9 2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zM3 4.81V8l1-1l1 1V4.81c-.31.11-.65.19-1 .19s-.69-.08-1-.19z"/>`
	banInnerSVG                   = `<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c.66 0 1.26.21 1.75.56L1.56 5.75C1.21 5.26 1 4.66 1 4c0-1.66 1.34-3 3-3zm2.44 1.25C6.79 2.74 7 3.34 7 4c0 1.66-1.34 3-3 3c-.66 0-1.26-.21-1.75-.56l4.19-4.19z"/>`
	barChartInnerSVG              = `<path fill="currentColor" d="M0 0v7h8V6H1V0H0zm5 0v5h2V0H5zM2 2v3h2V2H2z"/>`
	basketInnerSVG                = `<path fill="currentColor" d="M3.97 0c-.13.01-.26.08-.34.19L1.29 3H.01v1h1v3.66c0 .18.16.34.34.34h5.31c.18 0 .34-.16.34-.34V4h1V3H6.72C6.45 2.67 4.33.14 4.31.12c-.11-.09-.22-.14-.34-.13zM4 1.28L5.44 3H2.56L4 1.28zM2.5 5c.28 0 .5.22.5.5v1c0 .28-.22.5-.5.5S2 6.78 2 6.5v-1c0-.28.22-.5.5-.5zm3 0c.28 0 .5.22.5.5v1c0 .28-.22.5-.5.5S5 6.78 5 6.5v-1c0-.28.22-.5.5-.5z"/>`
	batteryEmptyInnerSVG          = `<path fill="currentColor" d="M.09 1C.03 1 0 1.04 0 1.09V6.9c0 .05.04.09.09.09H6.9c.05 0 .09-.04.09-.09V4.99h1v-2h-1V1.08c0-.06-.04-.09-.09-.09H.09zM1 2h5v4H1V2z"/>`
	batteryFullInnerSVG           = `<path fill="currentColor" d="M.09 1C.03 1 0 1.04 0 1.09V6.9c0 .05.04.09.09.09H6.9c.05 0 .09-.04.09-.09V4.99h1v-2h-1V1.08c0-.06-.04-.09-.09-.09H.09z"/>`
	beakerInnerSVG                = `<path fill="currentColor" d="M1.34 0a.5.5 0 0 0 .16 1H2v1.41C1.91 2.58.8 4.72.34 5.5C.18 5.76 0 6.11 0 6.56c0 .39.15.77.41 1.03S1.05 8 1.44 8h5.13c.38 0 .74-.16 1-.41h.03c.26-.26.41-.64.41-1.03c0-.45-.19-.8-.34-1.06c-.46-.78-1.57-2.92-1.66-3.09V1h.5a.5.5 0 1 0 0-1h-5a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0zM3 1h2v1.63l.06.09S5.69 3.95 6.25 5h-4.5c.56-1.05 1.19-2.28 1.19-2.28L3 2.63V1z"/>`
	bellInnerSVG                  = `<path fill="currentColor" d="M4 0C2.9 0 2 .9 2 2C2 3.04 1.48 3.98.66 4.66C.25 5 0 5.48 0 6h8c0-.52-.24-1-.66-1.34C6.52 3.98 6 3.04 6 2a2 2 0 0 0-2-2zM3 7c0 .55.45 1 1 1s1-.45 1-1H3z"/>`
	bluetoothInnerSVG             = `<path fill="currentColor" d="M2.5 0v2.5l-.75-.75L1 2.5L2.5 4L1 5.5l.75.75l.75-.75V8H3l3.5-2.5l-2.25-1.53L6.5 2.5L3 0h-.5zm1 1.5l1.5 1l-1.5 1v-2zm0 3l1.5 1l-1.5 1v-2z"/>`
	boldInnerSVG                  = `<path fill="currentColor" d="M0 0v1c.55 0 1 .45 1 1v4c0 .55-.45 1-1 1v1h5.5A2.5 2.5 0 0 0 8 5.5c0-1-.59-1.85-1.44-2.25A2 2 0 0 0 5 0H0zm3 1h1c.55 0 1 .45 1 1s-.45 1-1 1H3V1zm0 3h1.5C5.33 4 6 4.67 6 5.5S5.33 7 4.5 7H3V4z"/>`
	boltInnerSVG                  = `<path fill="currentColor" d="M4 0L1 5h2v3l3-5H4V0z"/>`
	bookInnerSVG                  = `<path fill="currentColor" d="M1 0C.93 0 .87.01.81.03C.42.11.11.42.03.81C0 .87 0 .93 0 1v5.5C0 7.33.67 8 1.5 8H7V7H1.5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H7V.5c0-.28-.22-.5-.5-.5H6v3L5 2L4 3V0H1z"/>`
	bookmarkInnerSVG              = `<path fill="currentColor" d="M2 0v8l2-2l2 2V0H2z"/>`
	boxInnerSVG                   = `<path fill="currentColor" d="M0 0v1h8V0H0zm0 2v5.91c0 .05.04.09.09.09H7.9c.05 0 .09-.04.09-.09V2H5.02v1.03H2.99V2h-3z"/>`
	briefcaseInnerSVG             = `<path fill="currentColor" d="M3 0c-.55 0-1 .45-1 1v1H.09C.03 2 0 2.04 0 2.09V4.5c0 .28.22.5.5.5h7c.28 0 .5-.22.5-.5V2.09C8 2.03 7.96 2 7.91 2H6V1c0-.55-.45-1-1-1H3zm0 1h2v1H3V1zM0 5.91v2c0 .05.04.09.09.09H7.9c.05 0 .09-.04.09-.09v-2c-.16.06-.32.09-.5.09h-7c-.18 0-.34-.04-.5-.09z"/>`
	britishPoundInnerSVG          = `<path fill="currentColor" d="M4 0c-.62 0-1.16.26-1.5.69c-.34.43-.5.99-.5 1.56c0 .69.16 1.25.25 1.75H1v1h1.22c-.11.45-.37.96-1.06 1.66L1 6.79v1.22h6v-1H2.09c.64-.73.98-1.4 1.13-2H5v-1H3.28C3.2 3.33 3 2.77 3 2.26c0-.39.11-.73.28-.94c.17-.21.37-.31.72-.31c.39 0 .61.11.75.25s.25.36.25.75h1c0-.58-.17-1.1-.53-1.47C5.1.17 4.58.01 4 .01z"/>`
	browserInnerSVG               = `<path fill="currentColor" d="M.34 0A.5.5 0 0 0 0 .5v7a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0zM1.5 1c.28 0 .5.22.5.5s-.22.5-.5.5s-.5-.22-.5-.5s.22-.5.5-.5zm2 0h3c.28 0 .5.22.5.5s-.22.5-.5.5h-3c-.28 0-.5-.22-.5-.5s.22-.5.5-.5zM1 3h6v4H1V3z"/>`
	brushInnerSVG                 = `<path fill="currentColor" d="M7.44.03c-.03 0-.04.02-.06.03L3.63 2.72c-.04.03-.1.11-.13.16l-.13.25c.72.23 1.27.78 1.5 1.5l.25-.13c.05-.03.12-.08.16-.13L7.94.62c.03-.05.04-.09 0-.13L7.5.05C7.48.03 7.46.02 7.44.02zM2.66 4c-.74 0-1.31.61-1.31 1.34c0 .99-.55 1.85-1.34 2.31c.39.22.86.34 1.34.34c1.47 0 2.66-1.18 2.66-2.66c0-.74-.61-1.34-1.34-1.34z"/>`
	bugInnerSVG                   = `<path fill="currentColor" d="M3.5 0C2.31 0 1.52 1.69 2.31 2.5c-.09.07-.2.14-.28.22L.72 2.06A.5.5 0 0 0 .38 2a.5.5 0 0 0-.09.94l1.16.56c-.09.16-.19.33-.25.5H.51a.5.5 0 0 0-.09 0a.5.5 0 1 0 .09 1h.5c0 .23.02.45.06.66l-.78.41a.5.5 0 1 0 .44.88l.66-.34c.25.46.62.85 1.03 1.09c.35-.19.59-.44.59-.72V5.54a.5.5 0 0 0 0-.09v-.81a.5.5 0 0 0 0-.22c.05-.23.26-.41.5-.41c.28 0 .5.22.5.5v.88a.5.5 0 0 0 0 .09v.06a.5.5 0 0 0 0 .09v1.34c0 .27.24.53.59.72c.41-.25.79-.63 1.03-1.09l.66.34a.5.5 0 1 0 .44-.88l-.78-.41c.04-.21.06-.43.06-.66h.5a.5.5 0 1 0 0-1h-.69c-.06-.17-.16-.34-.25-.5l1.16-.56a.5.5 0 0 0-.31-.94a.5.5 0 0 0-.13.06l-1.31.66c-.09-.08-.19-.15-.28-.22c.78-.83 0-2.5-1.19-2.5z"/>`
	bullhornInnerSVG              = `<path fill="currentColor" d="M6 0v6c.03.01.07 0 .09 0h.81c.05 0 .09-.04.09-.09V.1c0-.06-.04-.09-.09-.09h-.91zM5 .5L2.09 1.97c-.05.02-.13.03-.19.03H.09C.03 2 0 2.04 0 2.09V3.9c0 .06.04.09.09.09H1l1.03 2.72c.11.25.44.36.69.25c.25-.11.36-.44.25-.69l-.75-1.78c.03-.14.13-.22.28-.22v-.03L5 5.49v-5z"/>`
	calculatorInnerSVG            = `<path fill="currentColor" d="M.09 0C.03 0 0 .04 0 .09V7.9c0 .05.04.09.09.09H6.9c.05 0 .09-.04.09-.09V.09C6.99.03 6.95 0 6.9 0H.09zM1 1h5v2H1V1zm0 3h1v1H1V4zm2 0h1v1H3V4zm2 0h1v3H5V4zM1 6h1v1H1V6zm2 0h1v1H3V6z"/>`
	calendarInnerSVG              = `<path fill="currentColor" d="M0 0v2h7V0H0zm0 3v4.91c0 .05.04.09.09.09H6.9c.05 0 .09-.04.09-.09V3h-7zm1 1h1v1H1V4zm2 0h1v1H3V4zm2 0h1v1H5V4zM1 6h1v1H1V6zm2 0h1v1H3V6z"/>`
	cameraSlrInnerSVG             = `<path fill="currentColor" d="M4.09 0c-.05 0-.1.04-.13.09L3.02 1.9c-.02.05-.07.09-.13.09H1.48c-.83 0-1.5.67-1.5 1.5V7.9c0 .05.04.09.09.09h7.81c.05 0 .09-.04.09-.09V2.09c0-.06-.04-.09-.09-.09h-.81c-.05 0-.1-.04-.13-.09L6 .1C5.97.05 5.93.01 5.87.01H4.06zM1.5 3c.28 0 .5.22.5.5s-.22.5-.5.5s-.5-.22-.5-.5s.22-.5.5-.5zM5 3c1.1 0 2 .9 2 2s-.9 2-2 2s-2-.9-2-2s.9-2 2-2zm0 1c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1z"/>`
	caretBottomInnerSVG           = `<path fill="currentColor" d="m0 2l4 4l4-4H0z"/>`
	caretLeftInnerSVG             = `<path fill="currentColor" d="M6 0L2 4l4 4V0z"/>`
	caretRightInnerSVG            = `<path fill="currentColor" d="M2 0v8l4-4l-4-4z"/>`
	caretTopInnerSVG              = `<path fill="currentColor" d="M4 2L0 6h8L4 2z"/>`
	cartInnerSVG                  = `<path fill="currentColor" d="M.34 1A.506.506 0 0 0 .5 2H2l.09.25l.41 1.25l.41 1.25c.04.13.21.25.34.25h3.5c.14 0 .3-.12.34-.25l.81-2.5c.04-.13-.02-.25-.16-.25H3.3l-.38-.72A.5.5 0 0 0 2.48 1h-2a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0zM3.5 6c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5zm3 0c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5z"/>`
	chatInnerSVG                  = `<path fill="currentColor" d="M0 0v5l1-1h1V1h3V0H0zm3 2v4h4l1 1V2H3z"/>`
	checkInnerSVG                 = `<path fill="currentColor" d="m6.41 1l-.69.72L2.94 4.5l-.81-.78L1.41 3L0 4.41l.72.72l1.5 1.5l.69.72l.72-.72l3.5-3.5l.72-.72L6.41 1z"/>`
	chevronBottomInnerSVG         = `<path fill="currentColor" d="M1.5 1L0 2.5l4 4l4-4L6.5 1L4 3.5L1.5 1z"/>`
	chevronLeftInnerSVG           = `<path fill="currentColor" d="M5 0L1 4l4 4l1.5-1.5L4 4l2.5-2.5L5 0z"/>`
	chevronRightInnerSVG          = `<path fill="currentColor" d="M2.5 0L1 1.5L3.5 4L1 6.5L2.5 8l4-4l-4-4z"/>`
	chevronTopInnerSVG            = `<path fill="currentColor" d="M4 1L0 5l1.5 1.5L4 4l2.5 2.5L8 5L4 1z"/>`
	circleCheckInnerSVG           = `<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zm2 1.78l.72.72L3.5 5.72L1.78 4l.72-.72l1 1L6 1.78z"/>`
	circleXInnerSVG               = `<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zM2.5 1.78L4 3.28l1.5-1.5l.72.72L4.72 4l1.5 1.5l-.72.72L4 4.72l-1.5 1.5l-.72-.72L3.28 4l-1.5-1.5l.72-.72z"/>`
	clipboardInnerSVG             = `<path fill="currentColor" d="M3.5 0c-.28 0-.5.22-.5.5V1h-.75c-.14 0-.25.11-.25.25V2h3v-.75C5 1.11 4.89 1 4.75 1H4V.5c0-.28-.22-.5-.5-.5zM.25 1C.11 1 0 1.11 0 1.25v6.5c0 .14.11.25.25.25h6.5c.14 0 .25-.11.25-.25v-6.5C7 1.11 6.89 1 6.75 1H6v2H1V1H.25z"/>`
	clockInnerSVG                 = `<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c1.66 0 3 1.34 3 3S5.66 7 4 7S1 5.66 1 4s1.34-3 3-3zm-.5 1v2.22l.16.13l.5.5l.34.38l.72-.72l-.38-.34l-.34-.34V2.02h-1z"/>`
	cloudInnerSVG                 = `<path fill="currentColor" d="M4.5 1C3.29 1 2.23 1.86 2 3C.9 3 0 3.9 0 5s.9 2 2 2h4.5C7.33 7 8 6.33 8 5.5c0-.65-.42-1.29-1-1.5v-.5A2.5 2.5 0 0 0 4.5 1z"/>`
	cloudDownloadInnerSVG         = `<path fill="currentColor" d="M4.5 0C3.29 0 2.23.86 2 2C.9 2 0 2.9 0 4c0 .37.11.71.28 1H3v-.5C3 3.67 3.67 3 4.5 3S6 3.67 6 4.5V5h1.91c.06-.16.09-.32.09-.5c0-.65-.42-1.29-1-1.5v-.5A2.5 2.5 0 0 0 4.5 0zm-.16 4a.5.5 0 0 0-.34.5V6H2.5l2 2l2-2H5V4.5a.5.5 0 0 0-.59-.5a.5.5 0 0 0-.06 0z"/>`
	cloudUploadInnerSVG           = `<path fill="currentColor" d="M4.5 0C3.29 0 2.23.86 2 2C.9 2 0 2.9 0 4c0 .37.11.71.28 1H2.5l2-2l2 2h1.41c.06-.16.09-.32.09-.5c0-.65-.42-1.29-1-1.5v-.5A2.5 2.5 0 0 0 4.5 0zm0 4.5L2 7h2v.5a.5.5 0 1 0 1 0V7h2L4.5 4.5z"/>`
	cloudyInnerSVG                = `<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 0 2.5c0 .39.09.74.25 1.06c.3-.21.64-.37 1-.47C1.8 1.84 3.07 1 4.5 1c-.46-.6-1.18-1-2-1zm2 2C3.29 2 2.23 2.86 2 4C.9 4 0 4.9 0 6s.9 2 2 2h4.5C7.33 8 8 7.33 8 6.5c0-.65-.42-1.29-1-1.5v-.5A2.5 2.5 0 0 0 4.5 2z"/>`
	codeInnerSVG                  = `<path fill="currentColor" d="M5 1L2 7h1l3-6H5zM1 2L0 4l1 2h1L1 4l1-2H1zm5 0l1 2l-1 2h1l1-2l-1-2H6z"/>`
	cogInnerSVG                   = `<path fill="currentColor" d="M3.5 0L3 1.19c-.1.03-.19.08-.28.13L1.53.82l-.72.72l.5 1.19c-.05.1-.09.18-.13.28l-1.19.5v1l1.19.5c.04.1.08.18.13.28l-.5 1.19l.72.72l1.19-.5c.09.04.18.09.28.13l.5 1.19h1L5 6.83c.09-.04.19-.08.28-.13l1.19.5l.72-.72l-.5-1.19c.04-.09.09-.19.13-.28l1.19-.5v-1l-1.19-.5c-.03-.09-.08-.19-.13-.28l.5-1.19l-.72-.72l-1.19.5c-.09-.04-.19-.09-.28-.13L4.5 0h-1zM4 2.5c.83 0 1.5.67 1.5 1.5S4.83 5.5 4 5.5S2.5 4.83 2.5 4S3.17 2.5 4 2.5z"/>`
	collapseDownInnerSVG          = `<path fill="currentColor" d="M0 0v2h8V0H0zm2 3l2 2l2-2H2zM0 7v1h8V7H0z"/>`
	collapseLeftInnerSVG          = `<path fill="currentColor" d="M0 0v8h1V0H0zm6 0v8h2V0H6zM5 2L3 4l2 2V2z"/>`
	collapseRightInnerSVG         = `<path fill="currentColor" d="M0 0v8h2V0H0zm7 0v8h1V0H7zM3 2v4l2-2l-2-2z"/>`
	collapseUpInnerSVG            = `<path fill="currentColor" d="M0 0v1h8V0H0zm4 3L2 5h4L4 3zM0 6v2h8V6H0z"/>`
	commandInnerSVG               = `<path fill="currentColor" d="M1.5 0C.67 0 0 .67 0 1.5S.67 3 1.5 3H2v1h-.5C.67 4 0 4.67 0 5.5S.67 7 1.5 7S3 6.33 3 5.5V5h1v.5C4 6.33 4.67 7 5.5 7S7 6.33 7 5.5S6.33 4 5.5 4H5V3h.5C6.33 3 7 2.33 7 1.5S6.33 0 5.5 0S4 .67 4 1.5V2H3v-.5C3 .67 2.33 0 1.5 0zm0 1c.28 0 .5.22.5.5V2h-.5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5zm4 0c.28 0 .5.22.5.5s-.22.5-.5.5H5v-.5c0-.28.22-.5.5-.5zM3 3h1v1H3V3zM1.5 5H2v.5c0 .28-.22.5-.5.5S1 5.78 1 5.5s.22-.5.5-.5zM5 5h.5c.28 0 .5.22.5.5s-.22.5-.5.5s-.5-.22-.5-.5V5z"/>`
	commentSquareInnerSVG         = `<path fill="currentColor" d="M.09 0C.03 0 0 .04 0 .09V5.9c0 .05.04.09.09.09H6l2 2V.08c0-.06-.04-.09-.09-.09H.1z"/>`
	compassInnerSVG               = `<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c1.66 0 3 1.34 3 3S5.66 7 4 7S1 5.66 1 4s1.34-3 3-3zm2 1L3 3L2 6l3-1l1-3zM4 3.5c.28 0 .5.22.5.5s-.22.5-.5.5s-.5-.22-.5-.5s.22-.5.5-.5z"/>`
	contrastInnerSVG              = `<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c1.66 0 3 1.34 3 3S5.66 7 4 7V1z"/>`
	copywritingInnerSVG           = `<path fill="currentColor" d="M0 0v1h8V0H0zm0 2v1h5V2H0zm0 3v1h8V5H0zm0 2v1h6V7H0zm7.5 0c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5z"/>`
	creditCardInnerSVG            = `<path fill="currentColor" d="M.25 1C.11 1 0 1.11 0 1.25V2h8v-.75C8 1.11 7.89 1 7.75 1H.25zM0 3v3.75c0 .14.11.25.25.25h7.5c.14 0 .25-.11.25-.25V3H0zm1 2h1v1H1V5zm2 0h1v1H3V5z"/>`
	cropInnerSVG                  = `<path fill="currentColor" d="M1 0v1H0v1h1v5h5v1h1V7h1V6H7V1.5l1-1l-.5-.5l-1 1H2V0H1zm1 2h3.5L2 5.5V2zm4 .5V6H2.5L6 2.5z"/>`
	dashboardInnerSVG             = `<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c1.66 0 3 1.34 3 3S5.66 7 4 7S1 5.66 1 4s1.34-3 3-3zm0 1c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5S4.28 2 4 2zM2.34 3a.5.5 0 0 0-.19.84l.91.91C3.04 4.83 3 4.91 3 5c0 .55.45 1 1 1s1-.45 1-1s-.45-1-1-1c-.09 0-.17.04-.25.06l-.91-.91a.5.5 0 0 0-.44-.16a.5.5 0 0 0-.06 0zM5.5 3c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5z"/>`
	dataTransferDownloadInnerSVG  = `<path fill="currentColor" d="M3 0v3H1l3 3l3-3H5V0H3zM0 7v1h8V7H0z"/>`
	dataTransferUploadInnerSVG    = `<path fill="currentColor" d="M0 0v1h8V0H0zm4 2L1 5h2v3h2V5h2L4 2z"/>`
	deleteInnerSVG                = `<path fill="currentColor" d="M2 1L0 4l2 3h6V1H2zm1.5.78L5 3.28l1.5-1.5l.72.72L5.72 4l1.5 1.5l-.72.72L5 4.72l-1.5 1.5l-.72-.72L4.28 4l-1.5-1.5l.72-.72z"/>`
	dialInnerSVG                  = `<path fill="currentColor" d="M4 1C1.8 1 0 2.8 0 5h1c0-1.66 1.34-3 3-3s3 1.34 3 3h1c0-2.2-1.8-4-4-4zm-.59 2.09A2 2 0 0 0 4 7c1.11 0 2-.89 2-2c0-.9-.59-1.65-1.41-1.91L4 3.97l-.59-.88z"/>`
	documentInnerSVG              = `<path fill="currentColor" d="M0 0v8h7V4H3V0H0zm4 0v3h3L4 0zM1 2h1v1H1V2zm0 2h1v1H1V4zm0 2h4v1H1V6z"/>`
	dollarInnerSVG                = `<path fill="currentColor" d="M3 0v1h-.75C1.57 1 1 1.57 1 2.25v.5c0 .68.44 1.24 1.09 1.41l2.56.66c.14.04.34.29.34.44v.5c0 .14-.11.25-.25.25h-2.5a.56.56 0 0 1-.25-.06v-.94h-1v1c0 .34.2.63.44.78c.23.16.52.22.81.22h.75v1h1v-1h.75c.69 0 1.25-.56 1.25-1.25v-.5c0-.68-.44-1.24-1.09-1.41l-2.56-.66C2.2 3.15 2 2.9 2 2.75v-.5c0-.14.11-.25.25-.25h2.5c.11 0 .21.04.25.06V3h1V2c0-.34-.2-.63-.44-.78c-.23-.16-.52-.22-.81-.22H4V0H3z"/>`
	doubleQuoteSansLeftInnerSVG   = `<path fill="currentColor" d="M0 1v6l3-3V1H0zm5 0v6l3-3V1H5z"/>`
	doubleQuoteSansRightInnerSVG  = `<path fill="currentColor" d="M3 1L0 4v3h3V1zm5 0L5 4v3h3V1z"/>`
	doubleQuoteSerifLeftInnerSVG  = `<path fill="currentColor" d="M3 1C1.35 1 0 2.35 0 4v3h3V4H1c0-1.11.89-2 2-2V1zm5 0C6.35 1 5 2.35 5 4v3h3V4H6c0-1.11.89-2 2-2V1z"/>`
	doubleQuoteSerifRightInnerSVG = `<path fill="currentColor" d="M0 1v3h2c0 1.11-.89 2-2 2v1c1.65 0 3-1.35 3-3V1H0zm5 0v3h2c0 1.11-.89 2-2 2v1c1.65 0 3-1.35 3-3V1H5z"/>`
	dropletInnerSVG               = `<path fill="currentColor" d="m4 0l-.34.34C3.55.45 1 3.03 1 5.22c0 1.65 1.35 3 3 3s3-1.35 3-3C7 3.04 4.45.45 4.34.34L4 0zM2.5 4.72c.28 0 .5.22.5.5c0 .55.45 1 1 1c.28 0 .5.22.5.5s-.22.5-.5.5c-1.1 0-2-.9-2-2c0-.28.22-.5.5-.5z"/>`
	ejectInnerSVG                 = `<path fill="currentColor" d="M4 0L0 5h8L4 0zM0 6v2h8V6H0z"/>`
	elevatorInnerSVG              = `<path fill="currentColor" d="M4 0L1 3h6L4 0zM1 5l3 3l3-3H1z"/>`
	ellipsesInnerSVG              = `<path fill="currentColor" d="M0 3v2h2V3H0zm3 0v2h2V3H3zm3 0v2h2V3H6z"/>`
	envelopeClosedInnerSVG        = `<path fill="currentColor" d="M0 1v1l4 2l4-2V1H0zm0 2v4h8V3L4 5L0 3z"/>`
	envelopeOpenInnerSVG          = `<path fill="currentColor" d="M4 0L0 2v6h8V2L4 0zm0 1.13l3 1.5v1.88l-3 1.5l-3-1.5V2.63l3-1.5zM2 3.01v1l2 1l2-1v-1H2z"/>`
	euroInnerSVG                  = `<path fill="currentColor" d="M5 0C3.14 0 1.6 1.28 1.16 3H-.75L-1 4h2.01c0 .35.07.68.16 1H-.8l-.19 1h2.56c.7 1.19 1.97 2 3.44 2c.73 0 1.41-.21 2-.56V6.22c-.53.48-1.22.78-2 .78c-.89 0-1.67-.39-2.22-1h2.22l.16-1H2.2a3.01 3.01 0 0 1-.19-1h3.34l.16-1H2.2a2.988 2.988 0 0 1 4.56-1.44L6.92.5C6.35.19 5.71 0 5.01 0z"/>`
	excerptInnerSVG               = `<path fill="currentColor" d="M0 0v1h7V0H0zm0 2v1h5V2H0zm0 2v1h8V4H0zm0 2v1h1V6H0zm2 0v1h1V6H2zm2 0v1h1V6H4z"/>`
	expandDownInnerSVG            = `<path fill="currentColor" d="M0 0v1h8V0H0zm2 2l2 2l2-2H2zM0 6v2h8V6H0z"/>`
	expandLeftInnerSVG            = `<path fill="currentColor" d="M0 0v8h1V0H0zm6 0v8h2V0H6zM2 2v4l2-2l-2-2z"/>`
	expandRightInnerSVG           = `<path fill="currentColor" d="M0 0v8h2V0H0zm7 0v8h1V0H7zM6 2L4 4l2 2V2z"/>`
	expandUpInnerSVG              = `<path fill="currentColor" d="M0 0v2h8V0H0zm4 4L2 6h4L4 4zM0 7v1h8V7H0z"/>`
	externalLinkInnerSVG          = `<path fill="currentColor" d="M0 0v8h8V6H7v1H1V1h1V0H0zm4 0l1.5 1.5L3 4l1 1l2.5-2.5L8 4V0H4z"/>`
	eyeInnerSVG                   = `<path fill="currentColor" d="M4.03 1C1.5 1 0 4 0 4s1.5 3 4.03 3C6.5 7 8 4 8 4S6.5 1 4.03 1zM4 2a2 2 0 0 1 2 2c0 1.11-.89 2-2 2a2 2 0 1 1 0-4zm0 1c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1c0-.1-.04-.19-.06-.28a.495.495 0 1 1-.66-.66A.875.875 0 0 0 4 3z"/>`
	eyedropperInnerSVG            = `<path fill="currentColor" d="M3.31 0a.5.5 0 0 0-.19.84l.63.63L.16 5.13L0 5.29v2.72h2.69l.16-.16l3.66-3.66l.63.66a.5.5 0 1 0 .72-.69l-.94-.94l.66-.66c.59-.58.59-1.54 0-2.13c-.57-.57-1.56-.57-2.13 0l-.66.66l-.94-.94a.5.5 0 0 0-.47-.16a.5.5 0 0 0-.06 0zm1.16 2.19L5.78 3.5L2.62 6.66L1.34 5.35l3.13-3.16z"/>`
	fileInnerSVG                  = `<path fill="currentColor" d="M0 0v8h7V4H3V0H0zm4 0v3h3L4 0z"/>`
	fireInnerSVG                  = `<path fill="currentColor" d="M2 0c1 2-2 3-2 5s2 3 2 3c-.98-1.98 2-3 2-5S2 0 2 0zm3 3c1 2-2 3-2 5h3c.4 0 1-.5 1-2c0-2-2-3-2-3z"/>`
	flagInnerSVG                  = `<path fill="currentColor" d="M0 0v8h1V0H0zm2 0v4h2v1h4L6 3.03L8 1H5V0H2z"/>`
	flashInnerSVG                 = `<path fill="currentColor" d="M3.5 0L2 3h2l-.66 2H2l1 3l3-3H4.5L6 2H4l1-2H3.5z"/>`
	folderInnerSVG                = `<path fill="currentColor" d="M0 0v2h8V1H3V0H0zm0 3v4.5c0 .28.22.5.5.5h7c.28 0 .5-.22.5-.5V3H0z"/>`
	forkInnerSVG                  = `<path fill="currentColor" d="M1.5 0C.67 0 0 .67 0 1.5c0 .66.41 1.2 1 1.41V5.1c-.59.2-1 .75-1 1.41c0 .83.67 1.5 1.5 1.5S3 7.34 3 6.51c0-.6-.34-1.1-.84-1.34c.09-.09.21-.16.34-.16h2c.82 0 1.5-.68 1.5-1.5v-.59c.59-.2 1-.75 1-1.41C7 .68 6.33.01 5.5.01S4 .68 4 1.51c0 .66.41 1.2 1 1.41v.59c0 .28-.22.5-.5.5h-2c-.17 0-.35.04-.5.09V2.91c.59-.2 1-.75 1-1.41C3 .67 2.33 0 1.5 0z"/>`
	fullscreenEnterInnerSVG       = `<path fill="currentColor" d="M0 0v4l1.5-1.5L3 4l1-1l-1.5-1.5L4 0H0zm5 4L4 5l1.5 1.5L4 8h4V4L6.5 5.5L5 4z"/>`
	fullscreenExitInnerSVG        = `<path fill="currentColor" d="M1 0L0 1l1.5 1.5L0 4h4V0L2.5 1.5L1 0zm3 4v4l1.5-1.5L7 8l1-1l-1.5-1.5L8 4H4z"/>`
	globeInnerSVG                 = `<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zm0 1c.33 0 .64.09.94.19c-.21.2-.45.38-.41.56c.04.18.69.13.69.5c0 .27-.42.35-.13.66c.35.35-.64.98-.66 1.44c-.03.83.84.97 1.53.97c.42 0 .53.2.5.44c-.54.77-1.46 1.25-2.47 1.25c-.38 0-.73-.09-1.06-.22c.22-.44-.28-1.31-.75-1.59c-.23-.23-.72-.14-1-.25c-.09-.27-.18-.54-.19-.84c.03-.05.08-.09.16-.09c.19 0 .45.38.59.34c.18-.04-.74-1.31-.31-1.56c.2-.12.6.39.47-.16c-.12-.51.36-.28.66-.41c.26-.11.45-.41.13-.59c-.06-.03-.13-.1-.22-.19c.45-.27.97-.44 1.53-.44zm2.31 1.09c.18.22.32.46.44.72v.03c-.04.07-.11.11-.22.22c-.28.28-.32-.21-.44-.31c-.13-.12-.6.02-.66-.13c-.07-.18.5-.42.88-.53z"/>`
	graphInnerSVG                 = `<path fill="currentColor" d="M7.03 0L4 3L3 2L0 5.03l1 1L3 4l1 1l4-4l-.97-1zM0 7v1h8V7H0z"/>`
	gridFourUpInnerSVG            = `<path fill="currentColor" d="M0 0v1h1V0H0zm2 0v1h1V0H2zm2 0v1h1V0H4zm2 0v1h1V0H6zM0 2v1h1V2H0zm2 0v1h1V2H2zm2 0v1h1V2H4zm2 0v1h1V2H6zM0 4v1h1V4H0zm2 0v1h1V4H2zm2 0v1h1V4H4zm2 0v1h1V4H6zM0 6v1h1V6H0zm2 0v1h1V6H2zm2 0v1h1V6H4zm2 0v1h1V6H6z"/>`
	gridThreeUpInnerSVG           = `<path fill="currentColor" d="M0 0v2h2V0H0zm3 0v2h2V0H3zm3 0v2h2V0H6zM0 3v2h2V3H0zm3 0v2h2V3H3zm3 0v2h2V3H6zM0 6v2h2V6H0zm3 0v2h2V6H3zm3 0v2h2V6H6z"/>`
	gridTwoUpInnerSVG             = `<path fill="currentColor" d="M0 0v3h3V0H0zm5 0v3h3V0H5zM0 5v3h3V5H0zm5 0v3h3V5H5z"/>`
	hardDriveInnerSVG             = `<path fill="currentColor" d="M.19 0C.08 0 0 .08 0 .19V3.5c0 .28.22.5.5.5h6c.28 0 .5-.22.5-.5V.19C7 .08 6.92 0 6.81 0H.18zM0 4.91v2.91c0 .11.08.19.19.19h6.63c.11 0 .19-.08.19-.19V4.91c-.16.06-.32.09-.5.09h-6c-.18 0-.34-.04-.5-.09zM5.5 6c.28 0 .5.22.5.5s-.22.5-.5.5s-.5-.22-.5-.5s.22-.5.5-.5z"/>`
	headerInnerSVG                = `<path fill="currentColor" d="M0 0v1h.5c.28 0 .5.22.5.5v4c0 .28-.22.5-.5.5H0v1h3V6h-.5c-.28 0-.5-.22-.5-.5V4h3v1.5c0 .28-.22.5-.5.5H4v1h3V6h-.5c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5H7V0H4v1h.5c.28 0 .5.22.5.5V3H2V1.5c0-.28.22-.5.5-.5H3V0H0z"/>`
	headphonesInnerSVG            = `<path fill="currentColor" d="M4 0C2.35 0 1 1.35 1 3v1H.5a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5V3c0-1.11.89-2 2-2c1.11 0 2 .89 2 2v3.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5H7V3c0-1.65-1.35-3-3-3z"/>`
	heartInnerSVG                 = `<path fill="currentColor" d="M2 1c-.55 0-1.04.23-1.41.59C.23 1.95 0 2.44 0 3c0 .55.23 1.04.59 1.41L4 7.82l3.41-3.41C7.77 4.05 8 3.56 8 3c0-.55-.23-1.04-.59-1.41C7.05 1.23 6.56 1 6 1c-.55 0-1.04.23-1.41.59C4.23 1.95 4 2.44 4 3c0-.55-.23-1.04-.59-1.41C3.05 1.23 2.56 1 2 1z"/>`
	homeInnerSVG                  = `<path fill="currentColor" d="M4 0L0 3h1v4h2V5h2v2h2V2.97L8 3L4 0z"/>`
	imageInnerSVG                 = `<path fill="currentColor" d="M0 0v8h8V0H0zm1 1h6v3L6 3L5 4l2 2v1H6L2 3L1 4V1z"/>`
	inboxInnerSVG                 = `<path fill="currentColor" d="M.19 0C.08 0 0 .08 0 .19v7.63c0 .11.08.19.19.19h7.63c.11 0 .19-.08.19-.19V.19c0-.11-.08-.19-.19-.19H.19zM1 2h6v3H6L5 6H3L2 5H1V2z"/>`
	infinityInnerSVG              = `<path fill="currentColor" d="M2 2C.69 2 0 3.01 0 4s.69 2 2 2c.79 0 1.42-.56 2-1.22C4.58 5.44 5.19 6 6 6c1.31 0 2-1.01 2-2s-.69-2-2-2c-.81 0-1.42.56-2 1.22C3.42 2.56 2.79 2 2 2zm0 1c.42 0 .88.47 1.34 1c-.46.53-.92 1-1.34 1c-.74 0-1-.54-1-1c0-.46.26-1 1-1zm4 0c.74 0 1 .54 1 1c0 .46-.26 1-1 1c-.43 0-.89-.47-1.34-1c.46-.53.91-1 1.34-1z"/>`
	infoInnerSVG                  = `<path fill="currentColor" d="M5 0c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1zM3.5 2.5C2.67 2.5 2 3.17 2 4h1c0-.28.22-.5.5-.5s.5.22.5.5s-1 1.64-1 2.5C3 7.36 3.67 8 4.5 8S6 7.33 6 6.5H5c0 .28-.22.5-.5.5S4 6.78 4 6.5C4 6.14 5 4.66 5 4c0-.81-.67-1.5-1.5-1.5z"/>`
	italicInnerSVG                = `<path fill="currentColor" d="M2 0v1h1.63l-.06.13l-2 5l-.34.88H.01v1h5v-1H3.38l.06-.13l2-5L5.78 1H7V0H2z"/>`
	justifyCenterInnerSVG         = `<path fill="currentColor" d="M0 0v1h8V0H0zm0 2v1h8V2H0zm0 2v1h8V4H0zm1 2v1h6V6H1z"/>`
	justifyLeftInnerSVG           = `<path fill="currentColor" d="M0 0v1h8V0H0zm0 2v1h8V2H0zm0 2v1h8V4H0zm0 2v1h6V6H0z"/>`
	justifyRightInnerSVG          = `<path fill="currentColor" d="M0 0v1h8V0H0zm0 2v1h8V2H0zm0 2v1h8V4H0zm2 2v1h6V6H2z"/>`
	keyInnerSVG                   = `<path fill="currentColor" d="M5.5 0A2.5 2.5 0 0 0 3 2.5c0 .16 0 .32.03.47L0 6v2h3V6h2V5l.03-.03c.15.03.31.03.47.03a2.5 2.5 0 0 0 0-5zM6 1c.55 0 1 .45 1 1s-.45 1-1 1s-1-.45-1-1s.45-1 1-1z"/>`
	laptopInnerSVG                = `<path fill="currentColor" d="M1.34 1a.5.5 0 0 0-.34.5V5H0v1.5c0 .28.22.5.5.5h7.01c.28 0 .5-.22.5-.5V5h-1V1.5a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0zM2 2h4v3H5v1H3V5H2V2z"/>`
	layersInnerSVG                = `<path fill="currentColor" d="M0 0v4h4V0H0zm5 2v3H2v1h4V2H5zm2 2v3H4v1h4V4H7z"/>`
	lightbulbInnerSVG             = `<path fill="currentColor" d="M4.41 0a.5.5 0 0 0-.13.06l-3 1.5a.5.5 0 1 0 .44.88l3-1.5A.5.5 0 0 0 4.41 0zm1 1.5a.5.5 0 0 0-.13.06l-4 2a.5.5 0 1 0 .44.88l4-2a.5.5 0 0 0-.31-.94zm0 2a.5.5 0 0 0-.13.06l-3 1.5A.5.5 0 0 0 2.5 6h2a.506.506 0 0 0 .16-1l1.06-.56a.5.5 0 0 0-.31-.94zM2.85 7a.506.506 0 0 0 .16 1h1a.5.5 0 1 0 0-1h-1a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0z"/>`
	linkBrokenInnerSVG            = `<path fill="currentColor" d="M2 0v1H1v1h2V0H2zm3.88.03a1.9 1.9 0 0 0-.53.09c-.27.1-.53.25-.75.47l-.44.44a.5.5 0 1 0 .69.69l.44-.44c.11-.11.24-.17.38-.22c.35-.12.78-.07 1.06.22c.39.39.39 1.04 0 1.44l-1.5 1.5a.5.5 0 1 0 .69.69l1.5-1.5A1.98 1.98 0 0 0 6.45.07C6.27.03 6.07.03 5.89.04zM2.29 2.94a.5.5 0 0 0-.19.16L.6 4.6a1.98 1.98 0 0 0 0 2.81c.56.56 1.36.72 2.06.47c.27-.1.53-.25.75-.47l.44-.44a.5.5 0 1 0-.69-.69l-.44.44c-.11.11-.24.17-.38.22c-.35.12-.78.07-1.06-.22c-.39-.39-.39-1.04 0-1.44l1.5-1.5a.5.5 0 0 0-.44-.84a.5.5 0 0 0-.06 0zM5.01 6v2h1V7h1V6h-2z"/>`
	linkIntactInnerSVG            = `<path fill="currentColor" d="M5.88.03a1.9 1.9 0 0 0-.53.09c-.27.1-.53.25-.75.47a.5.5 0 1 0 .69.69c.11-.11.24-.17.38-.22c.35-.12.78-.07 1.06.22c.39.39.39 1.04 0 1.44l-1.5 1.5c-.44.44-.8.48-1.06.47c-.26-.01-.41-.13-.41-.13a.5.5 0 1 0-.5.88s.34.22.84.25c.5.03 1.2-.16 1.81-.78l1.5-1.5A1.98 1.98 0 0 0 6.44.07C6.26.03 6.06.03 5.88.04zm-2 2.31c-.5-.02-1.19.15-1.78.75L.6 4.59a1.98 1.98 0 0 0 0 2.81c.56.56 1.36.72 2.06.47c.27-.1.53-.25.75-.47a.5.5 0 1 0-.69-.69c-.11.11-.24.17-.38.22c-.35.12-.78.07-1.06-.22c-.39-.39-.39-1.04 0-1.44l1.5-1.5c.4-.4.75-.45 1.03-.44c.28.01.47.09.47.09a.5.5 0 1 0 .44-.88s-.34-.2-.84-.22z"/>`
	listInnerSVG                  = `<path fill="currentColor" d="M.5 0C.22 0 0 .22 0 .5s.22.5.5.5s.5-.22.5-.5S.78 0 .5 0zM2 0v1h6V0H2zM.5 2c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5S.78 2 .5 2zM2 2v1h6V2H2zM.5 4c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5S.78 4 .5 4zM2 4v1h6V4H2zM.5 6c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5S.78 6 .5 6zM2 6v1h6V6H2z"/>`
	listRichInnerSVG              = `<path fill="currentColor" d="M0 0v3h3V0H0zm4 0v1h4V0H4zm0 2v1h3V2H4zM0 4v3h3V4H0zm4 0v1h4V4H4zm0 2v1h3V6H4z"/>`
	locationInnerSVG              = `<path fill="currentColor" d="M8 0L0 4l3 1l1 3l4-8z"/>`
	lockLockedInnerSVG            = `<path fill="currentColor" d="M4 0C2.9 0 2 .9 2 2v1H1v4h6V3H6V2c0-1.1-.9-2-2-2zm0 1c.56 0 1 .44 1 1v1H3V2c0-.56.44-1 1-1z"/>`
	lockUnlockedInnerSVG          = `<path fill="currentColor" d="M4 0C2.9 0 2 .9 2 2h1c0-.56.44-1 1-1s1 .44 1 1v2H1v4h6V4H6V2c0-1.1-.9-2-2-2z"/>`
	loopInnerSVG                  = `<path fill="currentColor" d="M6 0v1H1c-.55 0-1 .45-1 1v1h1V2h5v1l2-1.5L6 0zM2 4L0 5.5L2 7V6h5c.55 0 1-.45 1-1V4H7v1H2V4z"/>`
	loopCircularInnerSVG          = `<path fill="currentColor" d="M4 1C2.35 1 1 2.35 1 4H0l1.5 2L3 4H2c0-1.11.89-2 2-2V1zm2.5 1L5 4h1c0 1.11-.89 2-2 2v1c1.65 0 3-1.35 3-3h1L6.5 2z"/>`
	loopSquareInnerSVG            = `<path fill="currentColor" d="M1 0v2h1V1h4v2H5l1.5 2.5L8 3H7V0H1zm.5 2.5L0 5h1v3h6V6H6v1H2V5h1L1.5 2.5z"/>`
	magnifyingGlassInnerSVG       = `<path fill="currentColor" d="M3.5 0C1.57 0 0 1.57 0 3.5S1.57 7 3.5 7c.59 0 1.17-.14 1.66-.41a1 1 0 0 0 .13.13l1 1a1.02 1.02 0 1 0 1.44-1.44l-1-1a1 1 0 0 0-.16-.13c.27-.49.44-1.06.44-1.66c0-1.93-1.57-3.5-3.5-3.5zm0 1C4.89 1 6 2.11 6 3.5c0 .66-.24 1.27-.66 1.72l-.03.03a1 1 0 0 0-.13.13c-.44.4-1.04.63-1.69.63c-1.39 0-2.5-1.11-2.5-2.5s1.11-2.5 2.5-2.5z"/>`
	mapInnerSVG                   = `<path fill="currentColor" d="M0 0v8h8V5.62a.5.5 0 0 0 0-.22V-.01H0zm1 1h6v4H5.5a.5.5 0 0 0-.09 0a.5.5 0 1 0 .09 1H7v1H1V1zm2.5 1C2.67 2 2 2.67 2 3.5C2 4.5 3.5 6 3.5 6S5 4.5 5 3.5C5 2.67 4.33 2 3.5 2zm0 1c.28 0 .5.22.5.5s-.22.5-.5.5s-.5-.22-.5-.5s.22-.5.5-.5z"/>`
	mapMarkerInnerSVG             = `<path fill="currentColor" d="M4 0C2.34 0 1 1.34 1 3c0 2 3 5 3 5s3-3 3-5c0-1.66-1.34-3-3-3zm0 1a2 2 0 0 1 2 2c0 1.11-.89 2-2 2a2 2 0 1 1 0-4z"/>`
	mediaPauseInnerSVG            = `<path fill="currentColor" d="M1 1v6h2V1H1zm4 0v6h2V1H5z"/>`
	mediaPlayInnerSVG             = `<path fill="currentColor" d="M1 1v6l6-3l-6-3z"/>`
	mediaRecordInnerSVG           = `<path fill="currentColor" d="M4 1C2.34 1 1 2.34 1 4s1.34 3 3 3s3-1.34 3-3s-1.34-3-3-3z"/>`
	mediaSkipBackwardInnerSVG     = `<path fill="currentColor" d="M4 1L0 4l4 3V1zm0 3l4 3V1L4 4z"/>`
	mediaSkipForwardInnerSVG      = `<path fill="currentColor" d="M0 1v6l4-3l-4-3zm4 3v3l4-3l-4-3v3z"/>`
	mediaStepBackwardInnerSVG     = `<path fill="currentColor" d="M0 1v6h2V1H0zm2 3l5 3V1L2 4z"/>`
	mediaStepForwardInnerSVG      = `<path fill="currentColor" d="M0 1v6l5-3l-5-3zm5 3v3h2V1H5v3z"/>`
	mediaStopInnerSVG             = `<path fill="currentColor" d="M1 1v6h6V1H1z"/>`
	medicalCrossInnerSVG          = `<path fill="currentColor" d="M2 0v2H0v4h2v2h4V6h2V2H6V0H2z"/>`
	menuInnerSVG                  = `<path fill="currentColor" d="M0 1v1h8V1H0zm0 2.97v1h8v-1H0zm0 3v1h8v-1H0z"/>`
	microphoneInnerSVG            = `<path fill="currentColor" d="M3.91-.03a1 1 0 0 0-.13.03A1 1 0 0 0 3 1v2a1 1 0 1 0 2 0V1A1 1 0 0 0 3.91-.03zM1.35 2a.5.5 0 0 0-.34.5V3c0 1.48 1.09 2.69 2.5 2.94V7h-.5c-.55 0-1 .45-1 1h4.01c0-.55-.45-1-1-1h-.5V5.94C5.93 5.7 7.02 4.48 7.02 3v-.5a.5.5 0 1 0-1 0V3c0 1.11-.89 2-2 2c-1.11 0-2-.89-2-2v-.5a.5.5 0 0 0-.59-.5a.5.5 0 0 0-.06 0z"/>`
	minusInnerSVG                 = `<path fill="currentColor" d="M0 3v2h8V3H0z"/>`
	monitorInnerSVG               = `<path fill="currentColor" d="M.34 0A.5.5 0 0 0 0 .5v5a.5.5 0 0 0 .5.5H3v1H2c-.55 0-1 .45-1 1h6c0-.55-.45-1-1-1H5V6h2.5a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0zM1 1h6v4H1V1z"/>`
	moonInnerSVG                  = `<path fill="currentColor" d="M2.72 0A3.988 3.988 0 0 0 0 3.78c0 2.21 1.79 4 4 4c1.76 0 3.25-1.14 3.78-2.72c-.4.13-.83.22-1.28.22c-2.21 0-4-1.79-4-4c0-.45.08-.88.22-1.28z"/>`
	moveInnerSVG                  = `<path fill="currentColor" d="M3.5 0L2 1.5h1V3H1.5V2L0 3.5L1.5 5V4H3v1.5H2L3.5 7L5 5.5H4V4h1.5v1L7 3.5L5.5 2v1H4V1.5h1L3.5 0z"/>`
	musicalNoteInnerSVG           = `<path fill="currentColor" d="M8 0C3 0 2 1 2 1v4.09A1.64 1.64 0 0 0 1.5 5C.67 5 0 5.67 0 6.5S.67 8 1.5 8S3 7.33 3 6.5V2.53c.73-.23 1.99-.44 4-.5v2.06A1.64 1.64 0 0 0 6.5 4C5.67 4 5 4.67 5 5.5S5.67 7 6.5 7S8 6.33 8 5.5V0z"/>`
	paperclipInnerSVG             = `<path fill="currentColor" d="M5 0c-.51 0-1.02.21-1.41.59L.81 3.31a2.746 2.746 0 0 0 0 3.88a2.746 2.746 0 0 0 3.88 0l1.25-1.25l-.69-.69l-1.16 1.13l-.09.13c-.69.69-1.81.69-2.5 0c-.68-.68-.66-1.78 0-2.47l2.78-2.75c.39-.39 1.04-.39 1.44 0c.39.39.37 1.01 0 1.41l-2.5 2.47c-.1.1-.27.1-.38 0c-.1-.1-.1-.27 0-.38l.06-.03l.91-.94l-.69-.69l-.97.97c-.48.48-.48 1.27 0 1.75s1.27.49 1.75 0l2.5-2.44c.78-.78.78-2.04 0-2.81C6.01.21 5.51.01 4.99.01z"/>`
	pencilInnerSVG                = `<path fill="currentColor" d="M6 0L5 1l2 2l1-1l-2-2zM4 2L0 6v2h2l4-4l-2-2z"/>`
	peopleInnerSVG                = `<path fill="currentColor" d="M5.5 0c-.51 0-.95.35-1.22.88c.45.54.72 1.28.72 2.13c0 .29-.03.55-.09.81c.19.11.38.19.59.19c.83 0 1.5-.9 1.5-2s-.67-2-1.5-2zm-3 1C1.67 1 1 1.9 1 3s.67 2 1.5 2S4 4.1 4 3s-.67-2-1.5-2zm4.75 3.16c-.43.51-1.02.82-1.69.84c.27.38.44.84.44 1.34V7h2V5.34c0-.52-.31-.97-.75-1.19zm-6.5 1C.31 5.38 0 5.83 0 6.35v1.66h5V6.35c0-.52-.31-.97-.75-1.19C3.8 5.69 3.19 6 2.5 6S1.2 5.68.75 5.16z"/>`
	personInnerSVG                = `<path fill="currentColor" d="M4 0C2.9 0 2 1.12 2 2.5S2.9 5 4 5s2-1.12 2-2.5S5.1 0 4 0zM1.91 5C.85 5.05 0 5.92 0 7v1h8V7c0-1.08-.84-1.95-1.91-2c-.54.61-1.28 1-2.09 1c-.81 0-1.55-.39-2.09-1z"/>`
	phoneInnerSVG                 = `<path fill="currentColor" d="M1.19 0C1.08 0 1 .08 1 .19v7.63c0 .11.08.19.19.19h4.63c.11 0 .19-.08.19-.19V.19c0-.11-.08-.19-.19-.19H1.19zM2 1h3v5H2V1zm1.5 5.5c.28 0 .5.22.5.5s-.22.5-.5.5S3 7.28 3 7s.22-.5.5-.5z"/>`
	pieChartInnerSVG              = `<path fill="currentColor" d="M3.5 0c-.97 0-1.84.4-2.47 1.03L4 4V.03A4.07 4.07 0 0 0 3.5 0zM5 1.06v3.41L2.28 7.19c.61.5 1.37.81 2.22.81C6.43 8 8 6.43 8 4.5c0-1.76-1.31-3.19-3-3.44zM.91 2.37C.35 2.91 0 3.66 0 4.5c0 .96.46 1.79 1.16 2.34l2.13-2.13L.91 2.37z"/>`
	pinInnerSVG                   = `<path fill="currentColor" d="M1.34 0a.5.5 0 0 0 .16 1H2v2H1c-.55 0-1 .45-1 1h3v3l.44 1L4 7V4h3c0-.55-.45-1-1-1H5V1h.5a.5.5 0 1 0 0-1h-4a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0z"/>`
	playCircleInnerSVG            = `<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zM3 2l3 2l-3 2V2z"/>`
	plusInnerSVG                  = `<path fill="currentColor" d="M3 0v3H0v2h3v3h2V5h3V3H5V0H3z"/>`
	powerStandbyInnerSVG          = `<path fill="currentColor" d="M3 0v4h1V0H3zM1.72 1.44l-.38.31C.53 2.39 0 3.39 0 4.5C0 6.43 1.57 8 3.5 8S7 6.43 7 4.5c0-1.11-.53-2.11-1.34-2.75l-.38-.31l-.63.78l.38.31C5.61 2.99 6 3.7 6 4.5C6 5.89 4.89 7 3.5 7S1 5.89 1 4.5c0-.8.36-1.51.94-1.97l.41-.31l-.63-.78z"/>`
	printInnerSVG                 = `<path fill="currentColor" d="M2 0v2h4V0H2zM.09 3C.03 3 0 3.04 0 3.09V5.9c0 .05.04.09.09.09H1v-2h6v2h.91c.05 0 .09-.04.09-.09V3.09C8 3.03 7.96 3 7.91 3H.1zM2 5v3h4V5H2z"/>`
	projectInnerSVG               = `<path fill="currentColor" d="M0 0v7h1V0H0zm7 0v7h1V0H7zM2 1v1h2V1H2zm1 2v1h2V3H3zm1 2v1h2V5H4z"/>`
	pulseInnerSVG                 = `<path fill="currentColor" d="m3.25 0l-.47 1.53L2 4.09l-.03-.06l-.09-.34H0v1h1.16l.38 1.16l.47 1.47l.47-1.5l.78-2.5l.78 2.5l.41 1.34l.53-1.28l.59-1.47l.13.28h2.31v-1H6.32l-.38-.75l-.5-.97L5.03 3l-.47 1.19l-.84-2.66L3.25 0z"/>`
	puzzlePieceInnerSVG           = `<path fill="currentColor" d="M3 0c-.28 0-.54.1-.72.28C2.1.46 2 .72 2 1c0 .28.18.48.28.72c.03.06.03.16.03.28H0v6h2.31c0-.12-.01-.22-.03-.28C2.18 7.48 2 7.28 2 7c0-.28.1-.54.28-.72C2.46 6.1 2.72 6 3 6c.28 0 .54.1.72.28c.18.18.28.44.28.72c0 .28-.18.48-.28.72c-.03.06-.03.16-.03.28H6V5.69c.12 0 .22.01.28.03c.24.1.44.28.72.28c.28 0 .54-.1.72-.28C7.9 5.54 8 5.28 8 5c0-.28-.1-.54-.28-.72C7.54 4.1 7.28 4 7 4c-.28 0-.48.18-.72.28c-.06.03-.16.03-.28.03V2H3.69c0-.12.01-.22.03-.28c.1-.24.28-.44.28-.72c0-.28-.1-.54-.28-.72C3.54.1 3.28 0 3 0z"/>`
	questionMarkInnerSVG          = `<path fill="currentColor" d="M4.47 0c-.85 0-1.48.26-1.88.66c-.4.4-.54.9-.59 1.28l1 .13c.04-.25.12-.5.31-.69C3.5 1.19 3.8 1 4.47 1c.66 0 1.02.16 1.22.34c.2.18.28.4.28.66c0 .83-.34 1.06-.84 1.5c-.5.44-1.16 1.08-1.16 2.25V6h1v-.25c0-.83.31-1.06.81-1.5c.5-.44 1.19-1.08 1.19-2.25c0-.48-.17-1.02-.59-1.41C5.95.2 5.31 0 4.47 0zm-.5 7v1h1V7h-1z"/>`
	rainInnerSVG                  = `<path fill="currentColor" d="M4.5 0C3.29 0 2.23.86 2 2C.9 2 0 2.9 0 4c0 .53.2.99.53 1.34c.26-.22.6-.34.97-.34c.2 0 .39.05.56.13C2.23 4.49 2.8 4 3.5 4c.69 0 1.27.49 1.44 1.13c.17-.07.36-.13.56-.13c.63 0 1.15.39 1.38.94c.64-.17 1.13-.75 1.13-1.44c0-.65-.42-1.29-1-1.5v-.5A2.5 2.5 0 0 0 4.51 0zM3.34 5a.5.5 0 0 0-.34.5v2a.5.5 0 1 0 1 0v-2a.5.5 0 0 0-.59-.5a.5.5 0 0 0-.06 0zm-2 1a.5.5 0 0 0-.34.5v1a.5.5 0 1 0 1 0v-1a.5.5 0 0 0-.59-.5a.5.5 0 0 0-.06 0zm4 0a.5.5 0 0 0-.34.5v1a.5.5 0 1 0 1 0v-1a.5.5 0 0 0-.59-.5a.5.5 0 0 0-.06 0z"/>`
	randomInnerSVG                = `<path fill="currentColor" d="M6 0v1h-.5c-.35 0-.56.1-.78.38L3.31 3.16L1.78 1.38C1.56 1.12 1.34 1 1 1H0v1h1c-.05 0 .01.04.03.03l1.63 1.91L1 6H0v1h1c.35 0 .56-.1.78-.38l1.53-1.91l1.66 1.91c.22.26.44.38.78.38H6v1l2-1.5L6 5v1h-.22c-.01-.01-.05-.04-.06-.03L3.97 3.91L5.5 2H6v1l2-1.5L6 0z"/>`
	reloadInnerSVG                = `<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4c1.1 0 2.12-.43 2.84-1.16l-.72-.72c-.54.54-1.29.88-2.13.88c-1.66 0-3-1.34-3-3s1.34-3 3-3c.83 0 1.55.36 2.09.91L4.99 3h3V0L6.8 1.19C6.08.47 5.09 0 3.99 0z"/>`
	resizeBothInnerSVG            = `<path fill="currentColor" d="m4 0l1.66 1.66l-4 4L0 4v4h4L2.34 6.34l4-4L8 4V0H4z"/>`
	resizeHeightInnerSVG          = `<path fill="currentColor" d="M3.5 0L1 3h2v2H1l2.5 3L6 5H4V3h2L3.5 0z"/>`
	resizeWidthInnerSVG           = `<path fill="currentColor" d="M3 1L0 3.5L3 6V4h2v2l3-2.5L5 1v2H3V1z"/>`
	rssInnerSVG                   = `<path fill="currentColor" d="M1 0v1c3.32 0 6 2.68 6 6h1c0-3.86-3.14-7-7-7zm0 2v1c2.21 0 4 1.79 4 4h1c0-2.76-2.24-5-5-5zm0 2v1c1.11 0 2 .89 2 2h1c0-1.65-1.35-3-3-3zm0 2c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1z"/>`
	rssAltInnerSVG                = `<path fill="currentColor" d="M0 0v2c3.33 0 6 2.67 6 6h2c0-4.41-3.59-8-8-8zm0 3v2c1.67 0 3 1.33 3 3h2c0-2.75-2.25-5-5-5zm0 3v2h2a2 2 0 0 0-2-2z"/>`
	scriptInnerSVG                = `<path fill="currentColor" d="M3 0c-.55 0-1 .45-1 1v5.5c0 .28-.22.5-.5.5S1 6.78 1 6.5V5H0v2c0 .55.45 1 1 1h5c.55 0 1-.45 1-1V4H3V1.5c0-.28.22-.5.5-.5s.5.22.5.5V3h4V1c0-.55-.45-1-1-1H3z"/>`
	shareInnerSVG                 = `<path fill="currentColor" d="M5 0v2C1 2 0 4.05 0 7c.52-1.98 2-3 4-3h1v2l3-3.16L5 0z"/>`
	shareBoxedInnerSVG            = `<path fill="currentColor" d="M.75 0C.34 0 0 .34 0 .75v5.5c0 .41.34.75.75.75h4.5c.41 0 .75-.34.75-.75V5H5v1H1V1h2V0H.75zM6 0v1C3.95 1 2.3 2.54 2.06 4.53C2.27 3.65 3.05 3 4 3h2v1l2-2l-2-2z"/>`
	shieldInnerSVG                = `<path fill="currentColor" d="m4 0l-.19.09l-3.5 1.47l-.31.13V2c0 1.66.67 3.12 1.47 4.19c.4.53.83.97 1.25 1.28c.42.31.83.53 1.28.53c.46 0 .86-.22 1.28-.53c.42-.31.85-.75 1.25-1.28C7.33 5.12 8 3.66 8 2v-.31l-.31-.13L4.19.09L4 0zm0 1.09V7c-.04 0-.33-.07-.66-.31s-.71-.63-1.06-1.09a6.26 6.26 0 0 1-1.22-3.28L4 1.1z"/>`
	signalInnerSVG                = `<path fill="currentColor" d="M6 0v8h1V0H6zM4 1v7h1V1H4zM2 3v5h1V3H2zM0 5v3h1V5H0z"/>`
	signpostInnerSVG              = `<path fill="currentColor" d="M3 0v1H1L0 2l1 1h2v5h1V4h2l1-1l-1-1H4V0H3z"/>`
	sortAscendingInnerSVG         = `<path fill="currentColor" d="M2 0v6H0l2.5 2L5 6H3V0H2zm2 0v1h2V0H4zm0 2v1h3V2H4zm0 2v1h4V4H4z"/>`
	sortDescendingInnerSVG        = `<path fill="currentColor" d="M2 0v6H0l2.5 2L5 6H3V0H2zm2 0v1h4V0H4zm0 2v1h3V2H4zm0 2v1h2V4H4z"/>`
	spreadsheetInnerSVG           = `<path fill="currentColor" d="M.75 0C.34 0 0 .34 0 .75v5.5c0 .41.34.75.75.75h6.5c.41 0 .75-.34.75-.75V.75C8 .34 7.66 0 7.25 0H.75zM1 1h1v1H1V1zm2 0h4v1H3V1zM1 3h1v1H1V3zm2 0h4v1H3V3zM1 5h1v1H1V5zm2 0h4v1H3V5z"/>`
	starInnerSVG                  = `<path fill="currentColor" d="M4 0L3 3H0l2.5 2l-1 3L4 6l2.5 2l-1-3L8 3H5L4 0z"/>`
	sunInnerSVG                   = `<path fill="currentColor" d="M4 0c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5S4.28 0 4 0zM1.5 1c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5zm5 0c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5zM4 2c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zM.5 3.5c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5zm7 0c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5zM1.5 6c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5zm5 0c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5zM4 7c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5S4.28 7 4 7z"/>`
	tabletInnerSVG                = `<path fill="currentColor" d="M.34 0A.35.35 0 0 0 0 .34v7.31c0 .18.16.34.34.34h6.31c.18 0 .34-.16.34-.34V.34A.35.35 0 0 0 6.65 0H.34zM1 1h5v5H1V1zm2.5 5.5c.38 0 .63.42.44.75s-.68.33-.88 0c-.19-.33.05-.75.44-.75z"/>`
	tagInnerSVG                   = `<path fill="currentColor" d="M0 0v3l5 5l3-3l-5-5H0zm2 1c.55 0 1 .45 1 1s-.45 1-1 1s-1-.45-1-1s.45-1 1-1z"/>`
	tagsInnerSVG                  = `<path fill="currentColor" d="M0 1v2l3 3l1.5-1.5L5 4L3 2L2 1H0zm3.41 0l3 3l-1.19 1.22L6 6l2-2l-3-3H3.41zM1.5 2c.28 0 .5.22.5.5s-.22.5-.5.5s-.5-.22-.5-.5s.22-.5.5-.5z"/>`
	targetInnerSVG                = `<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c1.66 0 3 1.34 3 3S5.66 7 4 7S1 5.66 1 4s1.34-3 3-3zm0 1c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm0 1c.56 0 1 .44 1 1s-.44 1-1 1s-1-.44-1-1s.44-1 1-1z"/>`
	taskInnerSVG                  = `<path fill="currentColor" d="M0 0v7h7V3.41l-1 1V6H1V1h3.59l1-1H0zm7 0L4 3L3 2L2 3l2 2l4-4l-1-1z"/>`
	terminalInnerSVG              = `<path fill="currentColor" d="M.09 0C.03 0 0 .04 0 .09V7.9c0 .05.04.09.09.09H7.9c.05 0 .09-.04.09-.09V.09C7.99.03 7.95 0 7.9 0H.09zM1.5.78L3.22 2.5L1.5 4.22L.78 3.5l1-1l-1-1L1.5.78zM4 3h3v1H4V3z"/>`
	textInnerSVG                  = `<path fill="currentColor" d="M0 0v2h.5c0-.55.45-1 1-1H3v5.5c0 .28-.22.5-.5.5H2v1h4V7h-.5c-.28 0-.5-.22-.5-.5V1h1.5c.55 0 1 .45 1 1H8V0H0z"/>`
	thumbDownInnerSVG             = `<path fill="currentColor" d="M0 0v4h1V0H0zm2 0v4c.28 0 .53.09.72.28c.19.19 1.15 2.12 1.28 2.38c.13.26.39.39.66.31c.26-.08.4-.36.31-.63c-.08-.26-.47-1.59-.47-1.84S4.72 4 5 4h1.5c.28 0 .5-.22.5-.5S5.97.31 5.97.31A.518.518 0 0 0 5.5 0H2z"/>`
	thumbUpInnerSVG               = `<path fill="currentColor" d="M4.47 0C4.28.02 4.1.15 4 .34C3.87.6 2.91 2.53 2.72 2.72C2.53 2.91 2.28 3 2 3v4h3.5c.21 0 .39-.13.47-.31C5.97 6.69 7 3.78 7 3.5c0-.28-.22-.5-.5-.5H5c-.28 0-.5-.25-.5-.5S4.89.92 4.97.66a.504.504 0 0 0-.31-.63C4.59.01 4.54-.01 4.47 0zM0 3v4h1V3H0z"/>`
	timerInnerSVG                 = `<path fill="currentColor" d="M2 0v1h1v.03A3.503 3.503 0 0 0 3.5 8C5.43 8 7 6.43 7 4.5c0-.45-.1-.87-.25-1.25l-.91.38c.11.29.16.57.16.88c0 1.39-1.11 2.5-2.5 2.5S1 5.9 1 4.51s1.11-2.5 2.5-2.5c.3 0 .59.05.88.16l.34-.94c-.23-.08-.47-.12-.72-.16v-.06h1v-1H2zm5 1.16s-3.65 2.81-3.84 3c-.19.2-.19.49 0 .69c.19.2.49.2.69 0c.2-.2 3.16-3.69 3.16-3.69z"/>`
	transferInnerSVG              = `<path fill="currentColor" d="M6 0v1H0v1h6v1l2-1.5L6 0zM2 4L0 5.5L2 7V6h6V5H2V4z"/>`
	trashInnerSVG                 = `<path fill="currentColor" d="M3 0c-.55 0-1 .45-1 1H1c-.55 0-1 .45-1 1h7c0-.55-.45-1-1-1H5c0-.55-.45-1-1-1H3zM1 3v4.81c0 .11.08.19.19.19h4.63c.11 0 .19-.08.19-.19V3h-1v3.5c0 .28-.22.5-.5.5s-.5-.22-.5-.5V3h-1v3.5c0 .28-.22.5-.5.5s-.5-.22-.5-.5V3h-1z"/>`
	underlineInnerSVG             = `<path fill="currentColor" d="M1 0v4c0 1.1 1.12 2 2.5 2H4c1.1 0 2-.9 2-2V0H5v4c0 .55-.45 1-1 1s-1-.45-1-1V0H1zM0 7v1h7V7H0z"/>`
	verticalAlignBottomInnerSVG   = `<path fill="currentColor" d="M.09 0C.03 0 0 .04 0 .09V4.9c0 .05.04.09.09.09H1.9c.05 0 .09-.04.09-.09V.09C1.99.03 1.95 0 1.9 0H.09zm6 0A.09.09 0 0 0 6 .09V4.9c0 .05.04.09.09.09H7.9c.05 0 .09-.04.09-.09V.09C7.99.03 7.95 0 7.9 0H6.09zm-3 2c-.06 0-.09.04-.09.09V4.9c0 .05.04.09.09.09H4.9c.05 0 .09-.04.09-.09V2.09c0-.06-.04-.09-.09-.09H3.09zM0 6v1h8V6H0z"/>`
	verticalAlignCenterInnerSVG   = `<path fill="currentColor" d="M.09 0C.03 0 0 .04 0 .09V2h2V.09C2 .03 1.96 0 1.91 0H.1zm6 0A.09.09 0 0 0 6 .09V2h2V.09C8 .03 7.96 0 7.91 0H6.1zm-3 1c-.06 0-.09.04-.09.09V2h2v-.91A.09.09 0 0 0 4.91 1H3.1zM0 3v1h8V3H0zm0 2v1.91c0 .05.04.09.09.09H1.9c.05 0 .09-.04.09-.09V5h-2zm3 0v.91c0 .05.04.09.09.09H4.9c.05 0 .09-.04.09-.09V5h-2zm3 0v1.91c0 .05.04.09.09.09H7.9c.05 0 .09-.04.09-.09V5h-2z"/>`
	verticalAlignTopInnerSVG      = `<path fill="currentColor" d="M0 0v1h8V0H0zm.09 2C.03 2 0 2.04 0 2.09V6.9c0 .05.04.09.09.09H1.9c.05 0 .09-.04.09-.09V2.09c0-.06-.04-.09-.09-.09H.09zm3 0c-.06 0-.09.04-.09.09V4.9c0 .05.04.09.09.09H4.9c.05 0 .09-.04.09-.09V2.09c0-.06-.04-.09-.09-.09H3.09zm3 0a.09.09 0 0 0-.09.09V6.9c0 .05.04.09.09.09H7.9c.05 0 .09-.04.09-.09V2.09c0-.06-.04-.09-.09-.09H6.09z"/>`
	videoInnerSVG                 = `<path fill="currentColor" d="M.5 1c-.28 0-.5.23-.5.5v4c0 .28.23.5.5.5h5c.28 0 .5-.22.5-.5V4l1 1h1V2H7L6 3V1.5c0-.28-.22-.5-.5-.5h-5z"/>`
	volumeHighInnerSVG            = `<path fill="currentColor" d="M3.34 0L2 2H0v4h2l1.34 2H4V0h-.66zM5 1v1c.17 0 .34.02.5.06c.86.22 1.5 1 1.5 1.94a1.987 1.987 0 0 1-2 2v1c.25 0 .48-.04.72-.09h.03C7.05 6.58 8 5.4 8 4c0-1.4-.95-2.58-2.25-2.91C5.52 1.03 5.26 1 5 1zm0 2v2c.09 0 .18-.01.25-.03c.43-.11.75-.51.75-.97a.997.997 0 0 0-1-1z"/>`
	volumeLowInnerSVG             = `<path fill="currentColor" d="M4.34 0L3 2H1v4h2l1.34 2H5V0h-.66zM6 3v2c.09 0 .18-.01.25-.03c.43-.11.75-.51.75-.97a.997.997 0 0 0-1-1z"/>`
	volumeOffInnerSVG             = `<path fill="currentColor" d="M5.34 0L4 2H2v4h2l1.34 2H6V0h-.66z"/>`
	warningInnerSVG               = `<path fill="currentColor" d="M3.09 0c-.06 0-.1.04-.13.09L.02 6.9c-.02.05-.03.13-.03.19v.81c0 .05.04.09.09.09h6.81c.05 0 .09-.04.09-.09v-.81c0-.05-.01-.14-.03-.19L4.01.09A.142.142 0 0 0 3.88 0h-.81zM3 3h1v2H3V3zm0 3h1v1H3V6z"/>`
	wifiInnerSVG                  = `<path fill="currentColor" d="M3.75 0C2.37 0 1.09.4 0 1.09l.53.81c.93-.59 2.03-.91 3.22-.91c1.2 0 2.32.31 3.25.91l.53-.81C6.44.39 5.13 0 3.75 0zm0 3c-.79 0-1.5.23-2.13.63l.53.84c.47-.3 1-.47 1.59-.47c.59 0 1.16.17 1.63.47l.53-.84C5.28 3.24 4.53 3 3.74 3zm0 3c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1z"/>`
	wrenchInnerSVG                = `<path fill="currentColor" d="M5.5 0A2.5 2.5 0 0 0 3 2.5c0 .32.08.62.19.91L.28 6.29c-.39.39-.39 1.05 0 1.44c.2.2.46.28.72.28c.26 0 .52-.09.72-.28L4.6 4.82c.28.11.58.19.91.19a2.5 2.5 0 0 0 2.5-2.5c0-.16 0-.32-.03-.47l-.97.97h-2v-2l.97-.97C5.83.01 5.67.01 5.51.01zM1 6.5c.28 0 .5.22.5.5s-.22.5-.5.5S.5 7.28.5 7s.22-.5.5-.5z"/>`
	xInnerSVG                     = `<path fill="currentColor" d="M1.41 0L0 1.41l.72.72L2.5 3.94L.72 5.72L0 6.41l1.41 1.44l.72-.72l1.81-1.81l1.78 1.81l.69.72l1.44-1.44l-.72-.69l-1.81-1.78l1.81-1.81l.72-.72L6.41 0l-.69.72L3.94 2.5L2.13.72L1.41 0z"/>`
	yenInnerSVG                   = `<path fill="currentColor" d="m0 0l2.25 3H0v1h3v1H0v1h3v2h1V6h3V5H4V4h3V3H4.75L7 0H6L3.69 3h-.38L1 0H0z"/>`
	zoomInInnerSVG                = `<path fill="currentColor" d="M3.5 0C1.57 0 0 1.57 0 3.5S1.57 7 3.5 7c.61 0 1.19-.16 1.69-.44a1 1 0 0 0 .09.13l1 1.03a1.02 1.02 0 1 0 1.44-1.44l-1.03-1a1 1 0 0 0-.13-.09c.27-.5.44-1.08.44-1.69C7 1.57 5.43 0 3.5 0zm0 1C4.89 1 6 2.11 6 3.5c0 .59-.2 1.14-.53 1.56l-.03.03a1 1 0 0 0-.06.03a1 1 0 0 0-.25.28c-.44.37-1.01.59-1.63.59c-1.39 0-2.5-1.11-2.5-2.5S2.11.99 3.5.99zM3 2v1H2v1h1v1h1V4h1V3H4V2H3z"/>`
	zoomOutInnerSVG               = `<path fill="currentColor" d="M3.5 0C1.57 0 0 1.57 0 3.5S1.57 7 3.5 7c.61 0 1.19-.16 1.69-.44a1 1 0 0 0 .09.13l1 1.03a1.02 1.02 0 1 0 1.44-1.44l-1.03-1a1 1 0 0 0-.13-.09c.27-.5.44-1.08.44-1.69C7 1.57 5.43 0 3.5 0zm0 1C4.89 1 6 2.11 6 3.5c0 .59-.2 1.14-.53 1.56l-.03.03a1 1 0 0 0-.06.03a1 1 0 0 0-.25.28c-.44.37-1.01.59-1.63.59c-1.39 0-2.5-1.11-2.5-2.5S2.11.99 3.5.99zM2 3v1h3V3H2z"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func AccountLogin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accountLoginInnerSVG).
		Element(children...)
}

func AccountLogout(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accountLogoutInnerSVG).
		Element(children...)
}

func ActionRedo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(actionRedoInnerSVG).
		Element(children...)
}

func ActionUndo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(actionUndoInnerSVG).
		Element(children...)
}

func AlignCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterInnerSVG).
		Element(children...)
}

func AlignLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignRightInnerSVG).
		Element(children...)
}

func Aperture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(apertureInnerSVG).
		Element(children...)
}

func ArrowBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBottomInnerSVG).
		Element(children...)
}

func ArrowCircleBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleBottomInnerSVG).
		Element(children...)
}

func ArrowCircleLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleLeftInnerSVG).
		Element(children...)
}

func ArrowCircleRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleRightInnerSVG).
		Element(children...)
}

func ArrowCircleTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleTopInnerSVG).
		Element(children...)
}

func ArrowLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftInnerSVG).
		Element(children...)
}

func ArrowRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightInnerSVG).
		Element(children...)
}

func ArrowThickBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowThickBottomInnerSVG).
		Element(children...)
}

func ArrowThickLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowThickLeftInnerSVG).
		Element(children...)
}

func ArrowThickRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowThickRightInnerSVG).
		Element(children...)
}

func ArrowThickTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowThickTopInnerSVG).
		Element(children...)
}

func ArrowTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTopInnerSVG).
		Element(children...)
}

func Audio(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(audioInnerSVG).
		Element(children...)
}

func AudioSpectrum(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(audioSpectrumInnerSVG).
		Element(children...)
}

func Badge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(badgeInnerSVG).
		Element(children...)
}

func Ban(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(banInnerSVG).
		Element(children...)
}

func BarChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartInnerSVG).
		Element(children...)
}

func Basket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketInnerSVG).
		Element(children...)
}

func BatteryEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryEmptyInnerSVG).
		Element(children...)
}

func BatteryFull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFullInnerSVG).
		Element(children...)
}

func Beaker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(beakerInnerSVG).
		Element(children...)
}

func Bell(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellInnerSVG).
		Element(children...)
}

func Bluetooth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boldInnerSVG).
		Element(children...)
}

func Bolt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boltInnerSVG).
		Element(children...)
}

func Book(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookInnerSVG).
		Element(children...)
}

func Bookmark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseInnerSVG).
		Element(children...)
}

func BritishPound(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(britishPoundInnerSVG).
		Element(children...)
}

func Browser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(browserInnerSVG).
		Element(children...)
}

func Brush(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brushInnerSVG).
		Element(children...)
}

func Bug(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bugInnerSVG).
		Element(children...)
}

func Bullhorn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bullhornInnerSVG).
		Element(children...)
}

func Calculator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calculatorInnerSVG).
		Element(children...)
}

func Calendar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarInnerSVG).
		Element(children...)
}

func CameraSlr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraSlrInnerSVG).
		Element(children...)
}

func CaretBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretBottomInnerSVG).
		Element(children...)
}

func CaretLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretLeftInnerSVG).
		Element(children...)
}

func CaretRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretRightInnerSVG).
		Element(children...)
}

func CaretTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretTopInnerSVG).
		Element(children...)
}

func Cart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartInnerSVG).
		Element(children...)
}

func Chat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatInnerSVG).
		Element(children...)
}

func Check(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkInnerSVG).
		Element(children...)
}

func ChevronBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronBottomInnerSVG).
		Element(children...)
}

func ChevronLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightInnerSVG).
		Element(children...)
}

func ChevronTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronTopInnerSVG).
		Element(children...)
}

func CircleCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleCheckInnerSVG).
		Element(children...)
}

func CircleX(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleXInnerSVG).
		Element(children...)
}

func Clipboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudInnerSVG).
		Element(children...)
}

func CloudDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDownloadInnerSVG).
		Element(children...)
}

func CloudUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudUploadInnerSVG).
		Element(children...)
}

func Cloudy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudyInnerSVG).
		Element(children...)
}

func Code(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeInnerSVG).
		Element(children...)
}

func Cog(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cogInnerSVG).
		Element(children...)
}

func CollapseDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(collapseDownInnerSVG).
		Element(children...)
}

func CollapseLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(collapseLeftInnerSVG).
		Element(children...)
}

func CollapseRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(collapseRightInnerSVG).
		Element(children...)
}

func CollapseUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(collapseUpInnerSVG).
		Element(children...)
}

func Command(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commandInnerSVG).
		Element(children...)
}

func CommentSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentSquareInnerSVG).
		Element(children...)
}

func Compass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassInnerSVG).
		Element(children...)
}

func Contrast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contrastInnerSVG).
		Element(children...)
}

func Copywriting(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(copywritingInnerSVG).
		Element(children...)
}

func CreditCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropInnerSVG).
		Element(children...)
}

func Dashboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashboardInnerSVG).
		Element(children...)
}

func DataTransferDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dataTransferDownloadInnerSVG).
		Element(children...)
}

func DataTransferUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dataTransferUploadInnerSVG).
		Element(children...)
}

func Delete(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deleteInnerSVG).
		Element(children...)
}

func Dial(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dialInnerSVG).
		Element(children...)
}

func Document(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentInnerSVG).
		Element(children...)
}

func Dollar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dollarInnerSVG).
		Element(children...)
}

func DoubleQuoteSansLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleQuoteSansLeftInnerSVG).
		Element(children...)
}

func DoubleQuoteSansRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleQuoteSansRightInnerSVG).
		Element(children...)
}

func DoubleQuoteSerifLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleQuoteSerifLeftInnerSVG).
		Element(children...)
}

func DoubleQuoteSerifRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleQuoteSerifRightInnerSVG).
		Element(children...)
}

func Droplet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropletInnerSVG).
		Element(children...)
}

func Eject(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ejectInnerSVG).
		Element(children...)
}

func Elevator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(elevatorInnerSVG).
		Element(children...)
}

func Ellipses(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsesInnerSVG).
		Element(children...)
}

func EnvelopeClosed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeClosedInnerSVG).
		Element(children...)
}

func EnvelopeOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeOpenInnerSVG).
		Element(children...)
}

func Euro(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(euroInnerSVG).
		Element(children...)
}

func Excerpt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(excerptInnerSVG).
		Element(children...)
}

func ExpandDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandDownInnerSVG).
		Element(children...)
}

func ExpandLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandLeftInnerSVG).
		Element(children...)
}

func ExpandRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandRightInnerSVG).
		Element(children...)
}

func ExpandUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandUpInnerSVG).
		Element(children...)
}

func ExternalLink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeInnerSVG).
		Element(children...)
}

func Eyedropper(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyedropperInnerSVG).
		Element(children...)
}

func File(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileInnerSVG).
		Element(children...)
}

func Fire(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fireInnerSVG).
		Element(children...)
}

func Flag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagInnerSVG).
		Element(children...)
}

func Flash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flashInnerSVG).
		Element(children...)
}

func Folder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderInnerSVG).
		Element(children...)
}

func Fork(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forkInnerSVG).
		Element(children...)
}

func FullscreenEnter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fullscreenEnterInnerSVG).
		Element(children...)
}

func FullscreenExit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fullscreenExitInnerSVG).
		Element(children...)
}

func Globe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeInnerSVG).
		Element(children...)
}

func Graph(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphInnerSVG).
		Element(children...)
}

func GridFourUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridFourUpInnerSVG).
		Element(children...)
}

func GridThreeUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridThreeUpInnerSVG).
		Element(children...)
}

func GridTwoUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridTwoUpInnerSVG).
		Element(children...)
}

func HardDrive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hardDriveInnerSVG).
		Element(children...)
}

func Header(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headerInnerSVG).
		Element(children...)
}

func Headphones(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartInnerSVG).
		Element(children...)
}

func Home(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxInnerSVG).
		Element(children...)
}

func Infinity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infinityInnerSVG).
		Element(children...)
}

func Info(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoInnerSVG).
		Element(children...)
}

func Italic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(italicInnerSVG).
		Element(children...)
}

func JustifyCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(justifyCenterInnerSVG).
		Element(children...)
}

func JustifyLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(justifyLeftInnerSVG).
		Element(children...)
}

func JustifyRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(justifyRightInnerSVG).
		Element(children...)
}

func Key(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyInnerSVG).
		Element(children...)
}

func Laptop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopInnerSVG).
		Element(children...)
}

func Layers(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersInnerSVG).
		Element(children...)
}

func Lightbulb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightbulbInnerSVG).
		Element(children...)
}

func LinkBroken(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkBrokenInnerSVG).
		Element(children...)
}

func LinkIntact(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkIntactInnerSVG).
		Element(children...)
}

func List(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listInnerSVG).
		Element(children...)
}

func ListRich(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listRichInnerSVG).
		Element(children...)
}

func Location(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationInnerSVG).
		Element(children...)
}

func LockLocked(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockLockedInnerSVG).
		Element(children...)
}

func LockUnlocked(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockUnlockedInnerSVG).
		Element(children...)
}

func Loop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loopInnerSVG).
		Element(children...)
}

func LoopCircular(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loopCircularInnerSVG).
		Element(children...)
}

func LoopSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loopSquareInnerSVG).
		Element(children...)
}

func MagnifyingGlass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magnifyingGlassInnerSVG).
		Element(children...)
}

func Map(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapInnerSVG).
		Element(children...)
}

func MapMarker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapMarkerInnerSVG).
		Element(children...)
}

func MediaPause(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaPauseInnerSVG).
		Element(children...)
}

func MediaPlay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaPlayInnerSVG).
		Element(children...)
}

func MediaRecord(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaRecordInnerSVG).
		Element(children...)
}

func MediaSkipBackward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaSkipBackwardInnerSVG).
		Element(children...)
}

func MediaSkipForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaSkipForwardInnerSVG).
		Element(children...)
}

func MediaStepBackward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaStepBackwardInnerSVG).
		Element(children...)
}

func MediaStepForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaStepForwardInnerSVG).
		Element(children...)
}

func MediaStop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaStopInnerSVG).
		Element(children...)
}

func MedicalCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(medicalCrossInnerSVG).
		Element(children...)
}

func Menu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuInnerSVG).
		Element(children...)
}

func Microphone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneInnerSVG).
		Element(children...)
}

func Minus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusInnerSVG).
		Element(children...)
}

func Monitor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonInnerSVG).
		Element(children...)
}

func Move(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveInnerSVG).
		Element(children...)
}

func MusicalNote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(musicalNoteInnerSVG).
		Element(children...)
}

func Paperclip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperclipInnerSVG).
		Element(children...)
}

func Pencil(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pencilInnerSVG).
		Element(children...)
}

func People(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(peopleInnerSVG).
		Element(children...)
}

func Person(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(personInnerSVG).
		Element(children...)
}

func Phone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneInnerSVG).
		Element(children...)
}

func PieChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartInnerSVG).
		Element(children...)
}

func Pin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinInnerSVG).
		Element(children...)
}

func PlayCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusInnerSVG).
		Element(children...)
}

func PowerStandby(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerStandbyInnerSVG).
		Element(children...)
}

func Print(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printInnerSVG).
		Element(children...)
}

func Project(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(projectInnerSVG).
		Element(children...)
}

func Pulse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pulseInnerSVG).
		Element(children...)
}

func PuzzlePiece(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(puzzlePieceInnerSVG).
		Element(children...)
}

func QuestionMark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionMarkInnerSVG).
		Element(children...)
}

func Rain(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rainInnerSVG).
		Element(children...)
}

func Random(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(randomInnerSVG).
		Element(children...)
}

func Reload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reloadInnerSVG).
		Element(children...)
}

func ResizeBoth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resizeBothInnerSVG).
		Element(children...)
}

func ResizeHeight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resizeHeightInnerSVG).
		Element(children...)
}

func ResizeWidth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resizeWidthInnerSVG).
		Element(children...)
}

func Rss(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssInnerSVG).
		Element(children...)
}

func RssAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssAltInnerSVG).
		Element(children...)
}

func Script(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scriptInnerSVG).
		Element(children...)
}

func Share(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareInnerSVG).
		Element(children...)
}

func ShareBoxed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareBoxedInnerSVG).
		Element(children...)
}

func Shield(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldInnerSVG).
		Element(children...)
}

func Signal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signalInnerSVG).
		Element(children...)
}

func Signpost(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signpostInnerSVG).
		Element(children...)
}

func SortAscending(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortAscendingInnerSVG).
		Element(children...)
}

func SortDescending(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortDescendingInnerSVG).
		Element(children...)
}

func Spreadsheet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spreadsheetInnerSVG).
		Element(children...)
}

func Star(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starInnerSVG).
		Element(children...)
}

func Sun(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunInnerSVG).
		Element(children...)
}

func Tablet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagInnerSVG).
		Element(children...)
}

func Tags(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagsInnerSVG).
		Element(children...)
}

func Target(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(targetInnerSVG).
		Element(children...)
}

func Task(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(taskInnerSVG).
		Element(children...)
}

func Terminal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalInnerSVG).
		Element(children...)
}

func Text(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textInnerSVG).
		Element(children...)
}

func ThumbDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbDownInnerSVG).
		Element(children...)
}

func ThumbUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbUpInnerSVG).
		Element(children...)
}

func Timer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timerInnerSVG).
		Element(children...)
}

func Transfer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transferInnerSVG).
		Element(children...)
}

func Trash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashInnerSVG).
		Element(children...)
}

func Underline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(underlineInnerSVG).
		Element(children...)
}

func VerticalAlignBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(verticalAlignBottomInnerSVG).
		Element(children...)
}

func VerticalAlignCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(verticalAlignCenterInnerSVG).
		Element(children...)
}

func VerticalAlignTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(verticalAlignTopInnerSVG).
		Element(children...)
}

func Video(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(videoInnerSVG).
		Element(children...)
}

func VolumeHigh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeHighInnerSVG).
		Element(children...)
}

func VolumeLow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeLowInnerSVG).
		Element(children...)
}

func VolumeOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeOffInnerSVG).
		Element(children...)
}

func Warning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningInnerSVG).
		Element(children...)
}

func Wifi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiInnerSVG).
		Element(children...)
}

func Wrench(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wrenchInnerSVG).
		Element(children...)
}

func X(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xInnerSVG).
		Element(children...)
}

func Yen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yenInnerSVG).
		Element(children...)
}

func ZoomIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
			engine.NewAttribute("viewBox", "0 0 8 8"),
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
	case "account-login":
		return AccountLogin(), nil
	case "account-logout":
		return AccountLogout(), nil
	case "action-redo":
		return ActionRedo(), nil
	case "action-undo":
		return ActionUndo(), nil
	case "align-center":
		return AlignCenter(), nil
	case "align-left":
		return AlignLeft(), nil
	case "align-right":
		return AlignRight(), nil
	case "aperture":
		return Aperture(), nil
	case "arrow-bottom":
		return ArrowBottom(), nil
	case "arrow-circle-bottom":
		return ArrowCircleBottom(), nil
	case "arrow-circle-left":
		return ArrowCircleLeft(), nil
	case "arrow-circle-right":
		return ArrowCircleRight(), nil
	case "arrow-circle-top":
		return ArrowCircleTop(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-thick-bottom":
		return ArrowThickBottom(), nil
	case "arrow-thick-left":
		return ArrowThickLeft(), nil
	case "arrow-thick-right":
		return ArrowThickRight(), nil
	case "arrow-thick-top":
		return ArrowThickTop(), nil
	case "arrow-top":
		return ArrowTop(), nil
	case "audio":
		return Audio(), nil
	case "audio-spectrum":
		return AudioSpectrum(), nil
	case "badge":
		return Badge(), nil
	case "ban":
		return Ban(), nil
	case "bar-chart":
		return BarChart(), nil
	case "basket":
		return Basket(), nil
	case "battery-empty":
		return BatteryEmpty(), nil
	case "battery-full":
		return BatteryFull(), nil
	case "beaker":
		return Beaker(), nil
	case "bell":
		return Bell(), nil
	case "bluetooth":
		return Bluetooth(), nil
	case "bold":
		return Bold(), nil
	case "bolt":
		return Bolt(), nil
	case "book":
		return Book(), nil
	case "bookmark":
		return Bookmark(), nil
	case "box":
		return Box(), nil
	case "briefcase":
		return Briefcase(), nil
	case "british-pound":
		return BritishPound(), nil
	case "browser":
		return Browser(), nil
	case "brush":
		return Brush(), nil
	case "bug":
		return Bug(), nil
	case "bullhorn":
		return Bullhorn(), nil
	case "calculator":
		return Calculator(), nil
	case "calendar":
		return Calendar(), nil
	case "camera-slr":
		return CameraSlr(), nil
	case "caret-bottom":
		return CaretBottom(), nil
	case "caret-left":
		return CaretLeft(), nil
	case "caret-right":
		return CaretRight(), nil
	case "caret-top":
		return CaretTop(), nil
	case "cart":
		return Cart(), nil
	case "chat":
		return Chat(), nil
	case "check":
		return Check(), nil
	case "chevron-bottom":
		return ChevronBottom(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-top":
		return ChevronTop(), nil
	case "circle-check":
		return CircleCheck(), nil
	case "circle-x":
		return CircleX(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clock":
		return Clock(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-download":
		return CloudDownload(), nil
	case "cloud-upload":
		return CloudUpload(), nil
	case "cloudy":
		return Cloudy(), nil
	case "code":
		return Code(), nil
	case "cog":
		return Cog(), nil
	case "collapse-down":
		return CollapseDown(), nil
	case "collapse-left":
		return CollapseLeft(), nil
	case "collapse-right":
		return CollapseRight(), nil
	case "collapse-up":
		return CollapseUp(), nil
	case "command":
		return Command(), nil
	case "comment-square":
		return CommentSquare(), nil
	case "compass":
		return Compass(), nil
	case "contrast":
		return Contrast(), nil
	case "copywriting":
		return Copywriting(), nil
	case "credit-card":
		return CreditCard(), nil
	case "crop":
		return Crop(), nil
	case "dashboard":
		return Dashboard(), nil
	case "data-transfer-download":
		return DataTransferDownload(), nil
	case "data-transfer-upload":
		return DataTransferUpload(), nil
	case "delete":
		return Delete(), nil
	case "dial":
		return Dial(), nil
	case "document":
		return Document(), nil
	case "dollar":
		return Dollar(), nil
	case "double-quote-sans-left":
		return DoubleQuoteSansLeft(), nil
	case "double-quote-sans-right":
		return DoubleQuoteSansRight(), nil
	case "double-quote-serif-left":
		return DoubleQuoteSerifLeft(), nil
	case "double-quote-serif-right":
		return DoubleQuoteSerifRight(), nil
	case "droplet":
		return Droplet(), nil
	case "eject":
		return Eject(), nil
	case "elevator":
		return Elevator(), nil
	case "ellipses":
		return Ellipses(), nil
	case "envelope-closed":
		return EnvelopeClosed(), nil
	case "envelope-open":
		return EnvelopeOpen(), nil
	case "euro":
		return Euro(), nil
	case "excerpt":
		return Excerpt(), nil
	case "expand-down":
		return ExpandDown(), nil
	case "expand-left":
		return ExpandLeft(), nil
	case "expand-right":
		return ExpandRight(), nil
	case "expand-up":
		return ExpandUp(), nil
	case "external-link":
		return ExternalLink(), nil
	case "eye":
		return Eye(), nil
	case "eyedropper":
		return Eyedropper(), nil
	case "file":
		return File(), nil
	case "fire":
		return Fire(), nil
	case "flag":
		return Flag(), nil
	case "flash":
		return Flash(), nil
	case "folder":
		return Folder(), nil
	case "fork":
		return Fork(), nil
	case "fullscreen-enter":
		return FullscreenEnter(), nil
	case "fullscreen-exit":
		return FullscreenExit(), nil
	case "globe":
		return Globe(), nil
	case "graph":
		return Graph(), nil
	case "grid-four-up":
		return GridFourUp(), nil
	case "grid-three-up":
		return GridThreeUp(), nil
	case "grid-two-up":
		return GridTwoUp(), nil
	case "hard-drive":
		return HardDrive(), nil
	case "header":
		return Header(), nil
	case "headphones":
		return Headphones(), nil
	case "heart":
		return Heart(), nil
	case "home":
		return Home(), nil
	case "image":
		return Image(), nil
	case "inbox":
		return Inbox(), nil
	case "infinity":
		return Infinity(), nil
	case "info":
		return Info(), nil
	case "italic":
		return Italic(), nil
	case "justify-center":
		return JustifyCenter(), nil
	case "justify-left":
		return JustifyLeft(), nil
	case "justify-right":
		return JustifyRight(), nil
	case "key":
		return Key(), nil
	case "laptop":
		return Laptop(), nil
	case "layers":
		return Layers(), nil
	case "lightbulb":
		return Lightbulb(), nil
	case "link-broken":
		return LinkBroken(), nil
	case "link-intact":
		return LinkIntact(), nil
	case "list":
		return List(), nil
	case "list-rich":
		return ListRich(), nil
	case "location":
		return Location(), nil
	case "lock-locked":
		return LockLocked(), nil
	case "lock-unlocked":
		return LockUnlocked(), nil
	case "loop":
		return Loop(), nil
	case "loop-circular":
		return LoopCircular(), nil
	case "loop-square":
		return LoopSquare(), nil
	case "magnifying-glass":
		return MagnifyingGlass(), nil
	case "map":
		return Map(), nil
	case "map-marker":
		return MapMarker(), nil
	case "media-pause":
		return MediaPause(), nil
	case "media-play":
		return MediaPlay(), nil
	case "media-record":
		return MediaRecord(), nil
	case "media-skip-backward":
		return MediaSkipBackward(), nil
	case "media-skip-forward":
		return MediaSkipForward(), nil
	case "media-step-backward":
		return MediaStepBackward(), nil
	case "media-step-forward":
		return MediaStepForward(), nil
	case "media-stop":
		return MediaStop(), nil
	case "medical-cross":
		return MedicalCross(), nil
	case "menu":
		return Menu(), nil
	case "microphone":
		return Microphone(), nil
	case "minus":
		return Minus(), nil
	case "monitor":
		return Monitor(), nil
	case "moon":
		return Moon(), nil
	case "move":
		return Move(), nil
	case "musical-note":
		return MusicalNote(), nil
	case "paperclip":
		return Paperclip(), nil
	case "pencil":
		return Pencil(), nil
	case "people":
		return People(), nil
	case "person":
		return Person(), nil
	case "phone":
		return Phone(), nil
	case "pie-chart":
		return PieChart(), nil
	case "pin":
		return Pin(), nil
	case "play-circle":
		return PlayCircle(), nil
	case "plus":
		return Plus(), nil
	case "power-standby":
		return PowerStandby(), nil
	case "print":
		return Print(), nil
	case "project":
		return Project(), nil
	case "pulse":
		return Pulse(), nil
	case "puzzle-piece":
		return PuzzlePiece(), nil
	case "question-mark":
		return QuestionMark(), nil
	case "rain":
		return Rain(), nil
	case "random":
		return Random(), nil
	case "reload":
		return Reload(), nil
	case "resize-both":
		return ResizeBoth(), nil
	case "resize-height":
		return ResizeHeight(), nil
	case "resize-width":
		return ResizeWidth(), nil
	case "rss":
		return Rss(), nil
	case "rss-alt":
		return RssAlt(), nil
	case "script":
		return Script(), nil
	case "share":
		return Share(), nil
	case "share-boxed":
		return ShareBoxed(), nil
	case "shield":
		return Shield(), nil
	case "signal":
		return Signal(), nil
	case "signpost":
		return Signpost(), nil
	case "sort-ascending":
		return SortAscending(), nil
	case "sort-descending":
		return SortDescending(), nil
	case "spreadsheet":
		return Spreadsheet(), nil
	case "star":
		return Star(), nil
	case "sun":
		return Sun(), nil
	case "tablet":
		return Tablet(), nil
	case "tag":
		return Tag(), nil
	case "tags":
		return Tags(), nil
	case "target":
		return Target(), nil
	case "task":
		return Task(), nil
	case "terminal":
		return Terminal(), nil
	case "text":
		return Text(), nil
	case "thumb-down":
		return ThumbDown(), nil
	case "thumb-up":
		return ThumbUp(), nil
	case "timer":
		return Timer(), nil
	case "transfer":
		return Transfer(), nil
	case "trash":
		return Trash(), nil
	case "underline":
		return Underline(), nil
	case "vertical-align-bottom":
		return VerticalAlignBottom(), nil
	case "vertical-align-center":
		return VerticalAlignCenter(), nil
	case "vertical-align-top":
		return VerticalAlignTop(), nil
	case "video":
		return Video(), nil
	case "volume-high":
		return VolumeHigh(), nil
	case "volume-low":
		return VolumeLow(), nil
	case "volume-off":
		return VolumeOff(), nil
	case "warning":
		return Warning(), nil
	case "wifi":
		return Wifi(), nil
	case "wrench":
		return Wrench(), nil
	case "x":
		return X(), nil
	case "yen":
		return Yen(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in oi icon set", name)
	}
}
