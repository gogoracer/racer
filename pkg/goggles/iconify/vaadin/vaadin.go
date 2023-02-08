package vaadin

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	abacusInnerSVG               = `<path fill="currentColor" d="M0 0v16h16V0H0zm14 2v3h-.1c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1H7.9c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1h-.2c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1H2V2h12zm-.1 8c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1h-.2c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1H4.9c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1H2V6h.1c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h3.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.1l-.1 4zM2 14v-3h.1c.2.6.8 1 1.4 1s1.2-.4 1.4-1h3.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.1v3H2z"/>`
	absolutePositionInnerSVG     = `<path fill="currentColor" d="M0 0v16h16V0H0zm15 15H1V9h3v1l3-2l-3-2v1H1V1h6v3H6l2 3l2-3H9V1h6v14z"/>`
	academyCapInnerSVG           = `<path fill="currentColor" d="M15.09 12.79a1 1 0 0 0-.086-1.638L15 5.33L14 6v5.15a1.001 1.001 0 0 0-.092 1.629l-.378.502a2.48 2.48 0 0 0-.53 1.498v1.222h.815a.88.88 0 0 0 .853-.664l.331-1.336v2h1v-1.21a2.486 2.486 0 0 0-.534-1.505zM8 0L0 4l8 5l8-5l-8-4z"/><path fill="currentColor" d="M8 10L3 6.67v1.71C3 9.29 5.94 12 8 12s5-2.71 5-3.62V6.67z"/>`
	accessibilityInnerSVG        = `<path fill="currentColor" d="M10.4 10h-.5c.1.3.1.7.1 1c0 2.2-1.8 4-4 4s-4-1.8-4-4c0-2.1 1.6-3.8 3.7-4l-.2-1C2.9 6.4 1 8.4 1 11c0 2.8 2.2 5 5 5c2.4 0 4.4-1.7 4.9-3.9l-.5-2.1z"/><path fill="currentColor" d="M13.1 13L12 8H7.9l-.2-1H11V6H7.5l-.6-2.5c.9-.1 1.6-.8 1.6-1.7C8.5.8 7.7 0 6.7 0S5 .8 5 1.8c0 .6.3 1.2.8 1.5L7.1 9h4.1l1.2 5H15v-1h-1.9z"/>`
	accordionMenuInnerSVG        = `<path fill="currentColor" d="M0 4v8h16V4H0zm15 7H1V7h14v4zM0 0h16v3H0V0zm0 13h16v3H0v-3z"/>`
	addDockInnerSVG              = `<path fill="currentColor" d="M0 11v5h16v-5H0zm12 4H9v-3h3v3zm0-8V5c0-5-8-5-8-5s5 0 5 5v2H7l3.5 3L14 7h-2z"/>`
	adjustInnerSVG               = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zM2 8c0-3.3 2.7-6 6-6v12c-3.3 0-6-2.7-6-6z"/>`
	adobeFlashInnerSVG           = `<path fill="currentColor" d="M0 0v16h16V0H0zm13 4.4C10 4.4 9.7 7 9.7 7H11v2H8.6C6.8 14.8 3 14 3 14v-2.5s2.5.6 3.9-4C8.7 1.4 13 2 13 2v2.4z"/>`
	airplaneInnerSVG             = `<path fill="currentColor" d="M12.3 6.5c.5-.5.9-.8 1.2-1.1c1.6-1.6 3.2-4.1 2.2-5.1s-3.4.6-5 2.2c-.3.3-.6.7-1.1 1.2L2.6.5C1.9.2 1.1.3.6.8l-.6.5L6.6 7c-1.3 1.6-2.7 3.1-3.4 4l-1.1-.6c-.5-.3-1.2-.3-1.6.2l-.3.3L3 13l2 2.8l.3-.3c.4-.4.5-1.1.2-1.6L5 12.8c.9-.7 2.4-2.1 4-3.4l5.7 6.6l.5-.5c.5-.5.6-1.3.3-2l-3.2-7z"/>`
	alarmInnerSVG                = `<path fill="currentColor" d="M8 5H7v5h4V9l-2.93.07L8 5zM5.46.87A2.099 2.099 0 0 0 2.651.335L1.66 1.1a2.095 2.095 0 0 0-.207 2.844zm8.88.23l-1-.77a2.097 2.097 0 0 0-2.796.534L14.54 3.94c.287-.356.46-.813.46-1.312c0-.602-.253-1.145-.659-1.528z"/><path fill="currentColor" d="M12.87 14A6.979 6.979 0 0 0 15 9.002A7.052 7.052 0 0 0 8.003 2A7.051 7.051 0 0 0 1 8.997a6.98 6.98 0 0 0 2.128 5.001l-.938.942a.63.63 0 0 0 .882.878l.998-.999c1.092.758 2.446 1.211 3.905 1.211s2.813-.453 3.928-1.226l.977 1.015a.63.63 0 0 0 .878-.882zm-10-5a5.18 5.18 0 0 1 5.127-5.13a5.181 5.181 0 0 1 5.133 5.127a5.181 5.181 0 0 1-5.127 5.133A5.181 5.181 0 0 1 2.87 9.003z"/>`
	alignCenterInnerSVG          = `<path fill="currentColor" d="M5 0h6v3H5V0zM1 4h14v3H1V4zm2 4h10v3H3V8zm-3 4h16v3H0v-3z"/>`
	alignJustifyInnerSVG         = `<path fill="currentColor" d="M0 0h16v3H0V0zm0 4h16v3H0V4zm0 8h16v3H0v-3zm0-4h16v3H0V8z"/>`
	alignLeftInnerSVG            = `<path fill="currentColor" d="M0 0h11v3H0V0zm0 4h15v3H0V4zm0 4h13v3H0V8zm0 4h16v3H0v-3z"/>`
	alignRightInnerSVG           = `<path fill="currentColor" d="M5 0h11v3H5V0zM1 4h15v3H1V4zm2 4h13v3H3V8zm-3 4h16v3H0v-3z"/>`
	altInnerSVG                  = `<path fill="currentColor" d="M3.89 9h2.22L5 5.1L3.89 9z"/><path fill="currentColor" d="M0 0v16h16V0H0zm7 12l-.61-2H3.61L3 12H2l2.27-8h1.46L8 12H7zm3 0H9V3h1v9zm4-5h-1v3.5s0 .5 1 .5v1c-1 0-2-.44-2-1.44V7h-.5V6h.5V5h1v1h1v1z"/>`
	altAInnerSVG                 = `<path fill="currentColor" d="M14 7V6h-1V5h-1v1h-.5v1h.5v3.56c0 1 .56 1.44 2 1.44v-1a.899.899 0 0 1-.998-.495L13 7h1zM9 3h1v9H9V3zm-6 9l.57-2h2.82L7 12h1L5.73 4H4.27L2 12h1zm2-6.9L6.11 9H3.89z"/>`
	ambulanceInnerSVG            = `<path fill="currentColor" d="M6.18 14a2 2 0 1 1-3.999.001A2 2 0 0 1 6.18 14zM14 14a2 2 0 1 1-3.999.001A2 2 0 0 1 14 14zM5 6H4v1H3v1h1v1h1V8h1V7H5V6z"/><path fill="currentColor" d="m15.76 8.64l-3-4.53A2.501 2.501 0 0 0 10.682 3H8V2a1 1 0 0 0-2 0v1H1.5A1.5 1.5 0 0 0 0 4.5V13h1.37a3.238 3.238 0 0 1 2.812-2a3.236 3.236 0 0 1 2.81 1.978l2.188.021a3.238 3.238 0 0 1 2.812-2a3.003 3.003 0 0 1 2.822 1.979l1.187.021v-3.57a1.427 1.427 0 0 0-.243-.795zm-8.84-.52a2.5 2.5 0 1 1-3.017-2.997a2.48 2.48 0 0 1 3.014 3.014zM10 8V5h.859a2.25 2.25 0 0 1 1.866.992L14.05 8z"/>`
	anchorInnerSVG               = `<path fill="currentColor" d="M13 9v2s-.8 1.7-4 1.9V6h2.2c.2.3.5.5.8.5c.6 0 1-.4 1-1s-.4-1-1-1c-.4 0-.7.2-.8.5H9V3.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2S6 .9 6 2c0 .7.4 1.4 1 1.7V5H4.8c-.1-.3-.4-.5-.8-.5c-.6 0-1 .4-1 1s.4 1 1 1c.4 0 .7-.2.8-.5H7v7c-3.3-.3-4-2-4-2V9H0s2.8 7 8 7c5 0 8-7 8-7h-3zM8 1c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1z"/>`
	angleDoubleDownInnerSVG      = `<path fill="currentColor" d="M3 2v2l5 5l5-5V2L8 7z"/><path fill="currentColor" d="M3 7v2l5 5l5-5V7l-5 5z"/>`
	angleDoubleLeftInnerSVG      = `<path fill="currentColor" d="M14 3h-2L7 8l5 5h2L9 8z"/><path fill="currentColor" d="M9 3H7L2 8l5 5h2L4 8z"/>`
	angleDoubleRightInnerSVG     = `<path fill="currentColor" d="M2 13h2l5-5l-5-5H2l5 5z"/><path fill="currentColor" d="M7 13h2l5-5l-5-5H7l5 5z"/>`
	angleDoubleUpInnerSVG        = `<path fill="currentColor" d="M13 14v-2L8 7l-5 5v2l5-5z"/><path fill="currentColor" d="M13 9V7L8 2L3 7v2l5-5z"/>`
	angleDownInnerSVG            = `<path fill="currentColor" d="M13 4v2l-5 5l-5-5V4l5 5z"/>`
	angleLeftInnerSVG            = `<path fill="currentColor" d="M12 13h-2L5 8l5-5h2L7 8z"/>`
	angleRightInnerSVG           = `<path fill="currentColor" d="M4 13h2l5-5l-5-5H4l5 5z"/>`
	angleUpInnerSVG              = `<path fill="currentColor" d="M3 12v-2l5-5l5 5v2L8 7z"/>`
	archiveInnerSVG              = `<path fill="currentColor" d="M0 1h16v3H0V1zm1 4v11h14V5H1zm10 4H5V7h6v2z"/>`
	archivesInnerSVG             = `<path fill="currentColor" d="M11 2H5v4h6V2zM9 4H7V3h2v1z"/><path fill="currentColor" d="M3 0v16h2v-1h6v1h2V0H3zm9 14H4V8h8v6zm0-7H4V1h8v6z"/><path fill="currentColor" d="M11 9H5v4h6V9zm-2 2H7v-1h2v1z"/>`
	areaSelectInnerSVG           = `<path fill="currentColor" d="m7.9 7.9l2.1 7.5l1.7-2.6l3.2 3.2l1.1-1.1l-3.3-3.2l2.7-1.6z"/><path fill="currentColor" d="M8 12H1V3h12v5.4l1 .2V2H0v11h8.2z"/>`
	arrowBackwardInnerSVG        = `<path fill="currentColor" d="M0 7.9L6 3v3h2c8 0 8 8 8 8s-1-4-7.8-4H6v2.9l-6-5z"/>`
	arrowCircleDownInnerSVG      = `<path fill="currentColor" d="M0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8zm9 1.6l1.8-1.8l1.4 1.4L8 13.4L3.8 9.2l1.4-1.4L7 9.6V3h2v6.6z"/>`
	arrowCircleDownOInnerSVG     = `<path fill="currentColor" d="M1 8c0-3.9 3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7zM0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8z"/><path fill="currentColor" d="m9 9.6l1.8-1.8l1.4 1.4L8 13.4L3.8 9.2l1.4-1.4L7 9.6V3h2v6.6z"/>`
	arrowCircleLeftInnerSVG      = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zM6.4 9l1.8 1.8l-1.4 1.4L2.6 8l4.2-4.2l1.4 1.4L6.4 7H13v2H6.4z"/>`
	arrowCircleLeftOInnerSVG     = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="m6.4 9l1.8 1.8l-1.4 1.4L2.6 8l4.2-4.2l1.4 1.4L6.4 7H13v2H6.4z"/>`
	arrowCircleRightInnerSVG     = `<path fill="currentColor" d="M8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8zm1.6-9L7.8 5.2l1.4-1.4L13.4 8l-4.2 4.2l-1.4-1.4L9.6 9H3V7h6.6z"/>`
	arrowCircleRightOInnerSVG    = `<path fill="currentColor" d="M8 15c-3.9 0-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7zm0 1c4.4 0 8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8z"/><path fill="currentColor" d="M9.6 7L7.8 5.2l1.4-1.4L13.4 8l-4.2 4.2l-1.4-1.4L9.6 9H3V7h6.6z"/>`
	arrowCircleUpInnerSVG        = `<path fill="currentColor" d="M16 8c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8zM7 6.4L5.2 8.2L3.8 6.8L8 2.6l4.2 4.2l-1.4 1.4L9 6.4V13H7V6.4z"/>`
	arrowCircleUpOInnerSVG       = `<path fill="currentColor" d="M15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7zm1 0c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8z"/><path fill="currentColor" d="M7 6.4L5.2 8.2L3.8 6.8L8 2.6l4.2 4.2l-1.4 1.4L9 6.4V13H7V6.4z"/>`
	arrowDownInnerSVG            = `<path fill="currentColor" d="M12.5 8.6L9 12.2V0H7v12.2L3.5 8.6l-1.4 1.5L8 16l5.9-5.9z"/>`
	arrowForwardInnerSVG         = `<path fill="currentColor" d="M16 7.9L10 3v3H8c-8 0-8 8-8 8s1-4 7.8-4H10v2.9l6-5z"/>`
	arrowLeftInnerSVG            = `<path fill="currentColor" d="M7.4 12.5L3.8 9H16V7H3.8l3.6-3.5l-1.5-1.4L0 8l5.9 5.9z"/>`
	arrowLongDownInnerSVG        = `<path fill="currentColor" d="M7 1h2v11h2l-3 3l-3-3h2z"/>`
	arrowLongLeftInnerSVG        = `<path fill="currentColor" d="M15 7v2H4v2L1 8l3-3v2z"/>`
	arrowRightInnerSVG           = `<path fill="currentColor" d="M8.6 3.5L12.1 7H0v2h12.1l-3.5 3.5l1.4 1.4L16 8l-6-5.9z"/>`
	arrowUpInnerSVG              = `<path fill="currentColor" d="M3.4 7.4L7 3.8V16h2V3.8l3.5 3.6l1.4-1.5L8 0L2 5.9z"/>`
	arrowsInnerSVG               = `<path fill="currentColor" d="m16 8l-3-3v2H9V3h2L8 0L5 3h2v4H3V5L0 8l3 3V9h4v4H5l3 3l3-3H9V9h4v2z"/>`
	arrowsCrossInnerSVG          = `<path fill="currentColor" d="M15 5V1h-4l1.3 1.3L8 6.6L3.7 2.3L5 1H1v4l1.3-1.3L6.6 8l-4.3 4.3L1 11v4h4l-1.3-1.3L8 9.4l4.3 4.3L11 15h4v-4l-1.3 1.3L9.4 8l4.3-4.3z"/>`
	arrowsLongHInnerSVG          = `<path fill="currentColor" d="m16 8l-3-3v2H3V5L0 8l3 3V9h10v2z"/>`
	arrowsLongRightInnerSVG      = `<path fill="currentColor" d="M1 9V7h11V5l3 3l-3 3V9z"/>`
	arrowsLongUpInnerSVG         = `<path fill="currentColor" d="M9 15H7V4H5l3-3l3 3H9z"/>`
	arrowsLongVInnerSVG          = `<path fill="currentColor" d="M9 3h2L8 0L5 3h2v10H5l3 3l3-3H9z"/>`
	asteriskInnerSVG             = `<path fill="currentColor" d="m15.9 5.7l-2-3.4L10 4.5V0H6v4.5L2 2.3L0 5.7L3.9 8L0 10.3l2 3.4l4-2.2V16h4v-4.5l3.9 2.2l2-3.4l-4-2.3z"/>`
	atInnerSVG                   = `<path fill="currentColor" d="M7.5 12.2c-2.3 0-4.2-1.9-4.2-4.2s1.9-4.2 4.2-4.2s4.2 1.9 4.2 4.2c.1 2.3-1.9 4.2-4.2 4.2zm0-7C6 5.2 4.8 6.5 4.8 8s1.2 2.8 2.8 2.8s2.8-1.2 2.8-2.8S9 5.2 7.5 5.2z"/><path fill="currentColor" d="M8 16c-4.4 0-8-3.6-8-8s3.6-8 8-8s8 3.6 8 8c0 1.5-.4 3-1.2 4.2c-.3.5-1.1 1.2-2.3 1.2c-.8 0-1.3-.3-1.6-.6c-.7-.7-.6-1.8-.6-1.9V4h1.5v7c0 .2 0 .6.2.8c0 0 .2.2.5.2c.7 0 1.1-.5 1.1-.5c.6-1 1-2.2 1-3.4c0-3.6-2.9-6.5-6.5-6.5S1.5 4.4 1.5 8s2.9 6.5 6.5 6.5c.7 0 1.3-.1 1.9-.3l.4 1.4c-.7.3-1.5.4-2.3.4z"/>`
	automationInnerSVG           = `<path fill="currentColor" d="M14 12a2 2 0 1 1-3.999.001A2 2 0 0 1 14 12z"/><path fill="currentColor" d="M11.7 16c-.8 0-1.6-.2-2.3-.7L3.2 12c-.5-.4-.9-.6-1.3-1C.7 9.8 0 8.1 0 6.4s.7-3.3 1.9-4.5C3.1.7 4.7 0 6.4 0S9.7.7 11 1.9c.4.4.6.7 1 1.2l3.5 6.4c1 1.7.7 3.8-.7 5.2c-.9.9-1.9 1.3-3.1 1.3zM6.4 1C5 1 3.6 1.6 2.6 2.6S1 5 1 6.4c0 1.5.6 2.8 1.6 3.8c.3.3.6.5 1.1.8l6.3 3.4c.6.4 1.2.5 1.8.5c.9 0 1.7-.3 2.3-1c1.1-1.1 1.3-2.7.5-4l-3.5-6.4c-.3-.4-.5-.7-.8-1C9.2 1.6 7.9 1 6.4 1z"/><path fill="currentColor" d="M11 7V6l-1.4-.5c-.1-.2-.1-.3-.2-.5l.6-1.3l-.7-.7l-1.3.6c-.2-.1-.3-.1-.5-.2L7 2H6l-.5 1.4c-.2.1-.3.1-.5.2L3.7 3l-.7.7l.6 1.3c-.1.2-.1.3-.2.5L2 6v1l1.4.5c.1.2.1.3.2.5L3 9.3l.7.7L5 9.4c.2.1.3.2.5.2L6 11h1l.5-1.4c.2-.1.3-.1.5-.2l1.3.6l.7-.7L9.4 8c.1-.2.2-.3.2-.5L11 7zM6.5 8C5.7 8 5 7.3 5 6.5S5.7 5 6.5 5S8 5.7 8 6.5S7.3 8 6.5 8z"/>`
	backspaceInnerSVG            = `<path fill="currentColor" d="M0 2v12h16V2H0zm13 7H6v2L3 8l3-3v2h7v2z"/>`
	backspaceAInnerSVG           = `<path fill="currentColor" d="M5 12L0 8l5-4v2h11v4H5v2z"/>`
	backwardsInnerSVG            = `<path fill="currentColor" d="M16 15V1L8 8zm-8 0V1L0 8z"/>`
	banInnerSVG                  = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm0 2c1.3 0 2.5.4 3.5 1.1l-8.4 8.4C2.4 10.5 2 9.3 2 8c0-3.3 2.7-6 6-6zm0 12c-1.3 0-2.5-.4-3.5-1.1l8.4-8.4c.7 1 1.1 2.2 1.1 3.5c0 3.3-2.7 6-6 6z"/>`
	barChartInnerSVG             = `<path fill="currentColor" d="M0 15h15v1H0v-1zm0-4h3v3H0v-3zm4-2h3v5H4V9zm4-4h3v9H8V5zm4-5h3v14h-3V0z"/>`
	barChartHInnerSVG            = `<path fill="currentColor" d="M1 15V0H0v16h16v-1H1z"/><path fill="currentColor" d="M2 8h4v6H2V8zm5-6h4v12H7V2zm5 4h4v8h-4V6z"/>`
	barChartVInnerSVG            = `<path fill="currentColor" d="M1 15V0H0v16h16v-1H1z"/><path fill="currentColor" d="M8 0v4H2V0h6zm6 5v4H2V5h12zm-4 5v4H2v-4h8z"/>`
	barcodeInnerSVG              = `<path fill="currentColor" d="M0 3h1v10H0V3zm8 0h2v10H8V3zm3 0h1v10h-1V3zm2 0h1v10h-1V3zm2 0h1v10h-1V3zM2 3h3v10H2V3zm4 0h1v10H6V3z"/>`
	bedInnerSVG                  = `<path fill="currentColor" d="M4.28 7H7L5.85 5.32a3.315 3.315 0 0 0-2.295-1.319L3 4v1.54a1.248 1.248 0 0 0 1.232 1.461L4.282 7zM13 7v-.29A1.71 1.71 0 0 0 11.322 5H6.63C7.13 5.62 8 7 8 7h5z"/><path fill="currentColor" d="M15 5.1a1 1 0 0 0-1 1V8H2V4a1 1 0 0 0-2 0v9h2v-2h12v2h2V6.1a1 1 0 0 0-1-1z"/>`
	bellInnerSVG                 = `<path fill="currentColor" d="M6 14h4s.1 2-2 2s-2-2-2-2zm6.7-2.6c-.5-.2-.7-.7-.7-1.2V5s-.2-2.4-3-2.9V1s.1-1-1-1s-1 1-1 1v1.1C4.2 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1h12v-1l-1.3-.6zM6 4.8V12H4c.8 0 1-1 1-1V5s0-.8.7-1.4C6.4 2.9 7 3 7 3s-1 .7-1 1.8z"/>`
	bellOInnerSVG                = `<path fill="currentColor" d="M12.7 11.4c-.5-.2-.7-.7-.7-1.2V5s0-2.4-3-2.9V1s.1-1-1-1s-1 1-1 1v1.1C4 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v2h4s-.1 2 2 2s2-2 2-2h4v-2l-1.3-.6zM13 13H3v-.4l.7-.4c.8-.3 1.3-1.1 1.3-2V5c0-.1 0-1.6 2.2-1.9l.8-.2l.8.1c2 .4 2.2 1.7 2.2 2v5.2c0 .9.5 1.7 1.3 2.1l.7.4v.3z"/>`
	bellSlashInnerSVG            = `<path fill="currentColor" d="m15.2 0l-3.6 3.6C11.1 3 10.4 2.3 9 2.1V1s.1-1-1-1s-1 1-1 1v1.1C4.2 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1h.3L0 15.3v.7h.7L16 .6V0h-.8zM6 4.8v4.5l-1 1V5s0-.8.7-1.4C6.4 2.9 7 3 7 3s-1 .7-1 1.8zM8 16c2.1 0 2-2 2-2H6s-.1 2 2 2zm4-5.8V5.6l-6 6l-.3.4l-1 1H14v-1l-1.3-.6c-.4-.3-.7-.7-.7-1.2z"/>`
	bellSlashOInnerSVG           = `<path fill="currentColor" d="m15.2 0l-3.6 3.6C11.2 3 10.4 2.3 9 2.1V1s.1-1-1-1s-1 1-1 1v1.1C4 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1.3l-2 2v.7h.7L16 .6V0h-.8zM5 10.3c0-.1 0-.1 0 0V5c0-.1.1-1.6 2.2-1.9l.8-.2l.8.1c1.2.2 1.8.8 2 1.3l-5.8 6zm7-.1V5.6l-1 1v3.5c0 .9.5 1.7 1.3 2.1l.7.4v.4H4.7l-1 1h2.4s-.1 2 2 2s2-2 2-2H14v-2l-1.3-.6c-.4-.3-.7-.7-.7-1.2z"/>`
	boatInnerSVG                 = `<path fill="currentColor" d="M1.5 9.6c1.1.7 2.5 1.9 2.5 3.3V14h.1s.9 0 2-1c1 1 2 1 2 1s1 0 2-1c1 1 1.9 1 1.9 1h.1v-1.1c0-1.4 1.4-2.6 2.5-3.3c.6-.4.5-1.2-.2-1.4L13 7.8V4h-1V3H9V1H7v2H4v1H3v3.8l-1.3.4c-.8.2-.8 1-.2 1.4zM4 5h1V4h6v1h1v2.5l-3.3-1c-.5-.1-1-.1-1.5 0L4 7.5V5zm10 9c-1 1-2 1-2 1s-1 0-2-1c-1 1-2 1-2 1s-1 0-2-1c-1 1-2 1-2 1s-1 0-2-1c-1 1-2 1-2 1v1h16v-1s-1 0-2-1z"/>`
	boldInnerSVG                 = `<path fill="currentColor" d="M11 7.5s2-.8 2-3.6C13-.2 7.9 0 6 0H2v16h4c3.7 0 8 0 8-4.4c0-3.8-3-4.1-3-4.1zM9 4.4C9 6.2 7.5 6 6 6V3c1.8 0 3 .1 3 1.4zM6 13V9c1.8 0 4-.3 4 2.2c0 1.9-2.5 1.8-4 1.8z"/>`
	boltInnerSVG                 = `<path fill="currentColor" d="M7.99 0L.98 9.38L7 8.96L2.04 16L15 6l-7.01.47L15 0H7.99z"/>`
	bombInnerSVG                 = `<path fill="currentColor" d="M12 1h1v1h-1V1zm0 4h1v1h-1V5zm2-2h1v1h-1V3zm-4 0h1v1h-1V3zm4.6-.9l.7-.7l-.7-.7l-1.4 1.4l.7.7zm-.7 2.1l-.7.7l1.4 1.4l.7-.7l-.7-.7zm-2.8-1.4l.7-.7L10.4.7l-.7.7l.7.7z"/><path fill="currentColor" d="m10.4 6.4l2-2l-.7-.7l-2 2L9 5l-.7.8C7.5 5.3 6.5 5 5.5 5C2.5 5 0 7.5 0 10.5S2.5 16 5.5 16s5.5-2.5 5.5-5.5c0-1-.3-1.9-.7-2.8L11 7l-.6-.6zM6 7.2C4 7.2 2.6 9 2.6 10h-1C1.6 8 4 6.2 6 6.2v1z"/>`
	bookInnerSVG                 = `<path fill="currentColor" d="M12.6 2.5C11 1.3 11 0 11 0H2v12.5C2 14.4 3.6 16 5.5 16H14V3s-1-.2-1.4-.5zM4 2h5v2H4V2zm9 13H5.5c-1 0-1.8-.6-2-1.3c-.1-.4 0-.7.4-.7H11V2.7c.4.6 1.2 1.1 2 1.3v11z"/>`
	bookDollarInnerSVG           = `<path fill="currentColor" d="M12.9 2.5C11.3 1.3 11.5 0 11.5 0H2v12.5C2 14.4 4.1 16 6 16h8V3s-.8-.2-1.1-.5zM7 6.3c-.9-.3-2.3-.8-2.3-1.9C4.8 3.6 6 3 6 2.8V2h1v.7c1 .1 1.8.4 1.9.4l-.3.9s-.7-.3-1.5-.3c-.7 0-1.1.3-1.2.8c0 .3.5.6 1.3.9c1.5.5 1.9 1.1 1.9 1.9C9.1 8 9 8.9 7 9.1v.9H6v-.8c0-.1-1.4-.5-1.5-.5l.5-.9s1.1.5 2 .4s1.3-.6 1.3-1c.1-.3-.4-.6-1.3-.9zm6 8.7H6c-1 0-1.8-.6-2-1.3c-.1-.3 0-.7.4-.7H11V2.7c1 .6 2 1.1 2 1.3v11z"/>`
	bookPercentInnerSVG          = `<path fill="currentColor" d="M12.6 2.5C11 1.3 11 0 11 0H2v12.5C2 14.4 3.6 16 5.5 16H14V3s-1-.2-1.4-.5zm-7.1.7c.8 0 1.5.7 1.5 1.6s-.7 1.4-1.5 1.4S4 5.6 4 4.8s.7-1.6 1.5-1.6zM9 3h1l-5 7H4l5-7zm1 5.5c0 .8-.7 1.5-1.5 1.5S7 9.3 7 8.5S7.7 7 8.5 7s1.5.7 1.5 1.5zm3 6.5H5.5c-1 0-1.8-.6-2-1.3c-.1-.4 0-.7.4-.7H11V2.7c0 .6 1 1.1 2 1.3v11z"/><path fill="currentColor" d="M9 8.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0zM6 4.8a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0z"/>`
	bookmarkInnerSVG             = `<path fill="currentColor" d="M3 0v1h10l.1-1zm0 2h10v14l-5-5l-5 5z"/>`
	bookmarkOInnerSVG            = `<path fill="currentColor" d="M3 0v16l5-5l5 5V0H3zm9 13.7L8 9.8l-4 3.9V3h8v10.7zM12 2H4V1h8v1z"/>`
	briefcaseInnerSVG            = `<path fill="currentColor" d="M11 3V1H5v2H0v12h16V3h-5zm-1 0H6V2h4v1z"/>`
	browserInnerSVG              = `<path fill="currentColor" d="M15 1V0H0v15h1v1h15V1h-1zM3 1h9v1H3V1zM1 1h1v1H1V1zm0 2h13v11H1V3z"/>`
	bugInnerSVG                  = `<path fill="currentColor" d="M8 6.2a7.596 7.596 0 0 0 3.903-1.129a.31.31 0 0 0 .098-.229L12 4.819a2.914 2.914 0 0 0-1.781-2.522c-.137-.05-.219-.16-.219-.29V.499a.5.5 0 0 0-1 0v1.2a.3.3 0 0 1-.3.3H7.3a.3.3 0 0 1-.3-.3v-1.2a.5.5 0 0 0-1 0v1.5a.3.3 0 0 1-.198.269A2.92 2.92 0 0 0 4 4.812l-.001.029c0 .102.051.193.13.247a5.847 5.847 0 0 0 3.89 1.11zM10 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zM6 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2z"/><path fill="currentColor" d="M13 8V7a6.196 6.196 0 0 0 2-4.497a.5.5 0 1 0-1-.003a4.544 4.544 0 0 1-1.105 2.906A4.777 4.777 0 0 1 9.326 7H6.73l-.075.001a4.777 4.777 0 0 1-3.561-1.586A4.54 4.54 0 0 1 2 2.514a.5.5 0 1 0-1-.003a6.192 6.192 0 0 0 1.996 4.486L3 8.001c-3 .06-3 1.42-3 3.47a.5.5 0 0 0 1 0c0-1.72 0-2.4 2-2.47a9.633 9.633 0 0 0 .612 3.136c-.383.006-.696.176-.942.414a3.857 3.857 0 0 0-.711 2.242c0 .2.015.397.044.589L2 15.5a.5.5 0 0 0 1 0v-.14a3.272 3.272 0 0 1 .389-2.096c.165-.152.401-.257.66-.264a3.802 3.802 0 0 0 2.934 1.998L7 14h2v1a3.798 3.798 0 0 0 2.94-1.98c.012-.02.015-.02.018-.02a1 1 0 0 1 .663.251a3.26 3.26 0 0 1 .377 2.127l.002.121a.5.5 0 0 0 1 0v-.14a3.875 3.875 0 0 0-.678-2.802a1.841 1.841 0 0 0-.9-.466c.336-.917.55-1.975.578-3.08c2-.012 2 .708 2 2.458a.5.5 0 0 0 1 0c0-2.03 0-3.39-3-3.47zm-4 5H7v-1h2v1zm0-2H7v-1h2v1zm0-2H7V8h2v1z"/>`
	bugOInnerSVG                 = `<path fill="currentColor" d="M13 8V7a6.196 6.196 0 0 0 2-4.497a.5.5 0 1 0-1-.003a4.544 4.544 0 0 1-1.105 2.906A4.777 4.777 0 0 1 9.326 7H6.73l-.075.001a4.777 4.777 0 0 1-3.561-1.586A4.54 4.54 0 0 1 2 2.514a.5.5 0 1 0-1-.003a6.192 6.192 0 0 0 1.996 4.486L3 8.001c-3 .06-3 1.42-3 3.47a.5.5 0 0 0 1 0c0-1.72 0-2.4 2-2.47a9.633 9.633 0 0 0 .612 3.136c-.383.006-.696.176-.942.414a3.857 3.857 0 0 0-.711 2.242c0 .2.015.397.044.589L2 15.5a.5.5 0 0 0 1 0v-.14a3.272 3.272 0 0 1 .389-2.096c.165-.152.401-.257.66-.264a4.748 4.748 0 0 0 2.92 1.994L7 14h2v1a4.745 4.745 0 0 0 2.939-1.983c.013-.017.016-.017.019-.017a1 1 0 0 1 .663.251a3.26 3.26 0 0 1 .377 2.127l.002.121a.5.5 0 0 0 1 0v-.14a3.875 3.875 0 0 0-.678-2.802a1.841 1.841 0 0 0-.9-.466c.336-.917.55-1.975.578-3.08c2-.012 2 .708 2 2.458a.5.5 0 0 0 1 0c0-2.03 0-3.39-3-3.47zm-7 5.5a3.333 3.333 0 0 1-1.083-.989L4.67 12.1l-.15-.39A8.478 8.478 0 0 1 4 9.013V7.35a5.425 5.425 0 0 0 1.973.647L6 13.57zm3-.5H7v-1h2v1zm0-2H7v-1h2v1zm0-2H7V8h2v1zm3 0a8.642 8.642 0 0 1-.54 2.77l-.13.33l-.24.4c-.285.411-.65.747-1.074.992L10 8a5.515 5.515 0 0 0 2.029-.624L12 9z"/><path fill="currentColor" d="M8 6.2a7.596 7.596 0 0 0 3.903-1.129a.31.31 0 0 0 .098-.229L12 4.819a2.914 2.914 0 0 0-1.781-2.522c-.137-.05-.219-.16-.219-.29V.499a.5.5 0 0 0-1 0v1.2a.3.3 0 0 1-.3.3H7.3a.3.3 0 0 1-.3-.3v-1.2a.5.5 0 0 0-1 0v1.5a.3.3 0 0 1-.198.269A2.92 2.92 0 0 0 4 4.812l-.001.029c0 .102.051.193.13.247a5.847 5.847 0 0 0 3.89 1.11zM10 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zM6 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2z"/>`
	buildingInnerSVG             = `<path fill="currentColor" d="M3 0v16h4v-3h2v3h4V0H3zm3 12H4v-2h2v2zm0-3H4V7h2v2zm0-3H4V4h2v2zm0-3H4V1h2v2zm3 9H7v-2h2v2zm0-3H7V7h2v2zm0-3H7V4h2v2zm0-3H7V1h2v2zm3 9h-2v-2h2v2zm0-3h-2V7h2v2zm0-3h-2V4h2v2zm0-3h-2V1h2v2z"/>`
	buildingOInnerSVG            = `<path fill="currentColor" d="M2 0v16h12V0H2zm11 15H9v-3H7v3H3V1h10v14z"/><path fill="currentColor" d="M4 9h2v2H4V9zm3 0h2v2H7V9zm3 0h2v2h-2V9zM4 6h2v2H4V6zm3 0h2v2H7V6zm3 0h2v2h-2V6zM4 3h2v2H4V3zm3 0h2v2H7V3zm3 0h2v2h-2V3z"/>`
	bulletsInnerSVG              = `<path fill="currentColor" d="M0 2.5C0 3.3.7 4 1.5 4S3 3.3 3 2.5S2.3 1 1.5 1S0 1.7 0 2.5zm0 5C0 8.3.7 9 1.5 9S3 8.3 3 7.5S2.3 6 1.5 6S0 6.7 0 7.5zm0 5c0 .8.7 1.5 1.5 1.5S3 13.3 3 12.5S2.3 11 1.5 11S0 11.7 0 12.5zM5 1h11v3H5V1zm0 5h11v3H5V6zm0 5h11v3H5v-3z"/>`
	bullseyeInnerSVG             = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm0 14.9c-3.8 0-6.9-3.1-6.9-6.9S4.2 1.1 8 1.1s6.9 3.1 6.9 6.9s-3.1 6.9-6.9 6.9z"/><path fill="currentColor" d="M8 2.3C4.8 2.3 2.3 4.8 2.3 8s2.6 5.7 5.7 5.7s5.7-2.6 5.7-5.7S11.2 2.3 8 2.3zm0 10.3c-2.5 0-4.6-2.1-4.6-4.6S5.5 3.4 8 3.4s4.6 2.1 4.6 4.6c0 2.5-2.1 4.6-4.6 4.6z"/><path fill="currentColor" d="M8 4.6C6.1 4.6 4.6 6.1 4.6 8s1.5 3.4 3.4 3.4s3.4-1.5 3.4-3.4S9.9 4.6 8 4.6z"/>`
	bussInnerSVG                 = `<path fill="currentColor" d="M14.67 4H14V2a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v2h-.68a.32.32 0 0 0-.32.32v2.36c0 .177.143.32.32.32H2v6c0 .55 0 1 1 1v1.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V14h4v1.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V14c1 0 1-.45 1-1V7h.67a.33.33 0 1 0 0-.66a.33.33 0 0 0 0 .66a.33.33 0 0 0 .33-.33V4.33a.33.33 0 0 0-.33-.33zM6 1h4v1H6V1zM4 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2zM3 8V3h10v5H3zm9 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/>`
	buttonInnerSVG               = `<path fill="currentColor" d="m15.7 5.3l-1-1c-.2-.2-.4-.3-.7-.3H1c-.6 0-1 .4-1 1v5c0 .3.1.6.3.7l1 1c.2.2.4.3.7.3h13c.6 0 1-.4 1-1V6c0-.3-.1-.5-.3-.7zM14 10H1V5h13v5z"/>`
	calcInnerSVG                 = `<path fill="currentColor" d="M9 3h6v2H9V3zm0 8h6v2H9v-2zM5 1H3v2H1v2h2v2h2V5h2V3H5zm2 9.4L5.6 9L4 10.6L2.4 9L1 10.4L2.6 12L1 13.6L2.4 15L4 13.4L5.6 15L7 13.6L5.4 12zm6 4.1a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm0-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`
	calcBookInnerSVG             = `<path fill="currentColor" d="M11.9 0c-1.3 0-2 .4-2.4.8C9.1.4 8.4 0 7 0C3.6 0 3 2 3 2v4H0v10h7v-4.6l1.5-.2s.2-.3.3.7h1.3c.1-1 .4-.7.4-.7l5.5.7V2.1S15.4 0 11.9 0zM1 7h5v2H1V7zm5 3v1H5v-1h1zm-2 0v1H3v-1h1zm-2 5H1v-1h1v1zm0-2H1v-1h1v1zm0-2H1v-1h1v1zm2 4H3v-1h1v1zm0-2H3v-1h1v1zm2 2H5v-1h1v1zm0-2H5v-1h1v1zm3-3.5c-.9-.1-1.3-.3-2-.3V6H4V2.1c0-.4.8-1.5 3-1.5c1.8 0 1.9.8 1.9 1v7.9zm6 .4c-1-.4-1.1-.7-2.5-.7h-.2c-1 0-1.3.2-2.3.4V1.8c0-.2.2-1.1 1.9-1.1c2.3 0 3.1.9 3.1 1.4v7.8z"/>`
	calendarInnerSVG             = `<path fill="currentColor" d="M14 1v3h-3V1H5v3H2V1H0v15h16V1h-2zM3 15H1v-2h2v2zm0-3H1v-2h2v2zm0-3H1V7h2v2zm3 6H4v-2h2v2zm0-3H4v-2h2v2zm0-3H4V7h2v2zm3 6H7v-2h2v2zm0-3H7v-2h2v2zm0-3H7V7h2v2zm3 6h-2v-2h2v2zm0-3h-2v-2h2v2zm0-3h-2V7h2v2zm3 6h-2v-2h2v2zm0-3h-2v-2h2v2zm0-3h-2V7h2v2z"/><path fill="currentColor" d="M3 0h1v3H3V0zm9 0h1v3h-1V0z"/>`
	calendarBriefcaseInnerSVG    = `<path fill="currentColor" d="M3 0h1v3H3V0zm8 0h1v3h-1V0z"/><path fill="currentColor" d="M13 1v3h-3V1H5v3H2V1H0v14h5v-1H1V6h13v3h1V1z"/><path fill="currentColor" d="M13 10V8H9v2H6v6h10v-6h-3zm-3-1h2v1h-2V9z"/>`
	calendarClockInnerSVG        = `<path fill="currentColor" d="M3 0h1v3H3V0zm8 0h1v3h-1V0z"/><path fill="currentColor" d="M6.6 14H1V6h13v.6c.4.2.7.4 1 .7V1h-2v3h-3V1H5v3H2V1H0v14h7.3c-.3-.3-.5-.6-.7-1z"/><path fill="currentColor" d="M14 12h-3V9h1v2h2z"/><path fill="currentColor" d="M11.5 8c1.9 0 3.5 1.6 3.5 3.5S13.4 15 11.5 15S8 13.4 8 11.5S9.6 8 11.5 8zm0-1C9 7 7 9 7 11.5S9 16 11.5 16s4.5-2 4.5-4.5S14 7 11.5 7z"/>`
	calendarEnvelopeInnerSVG     = `<path fill="currentColor" d="M3 0h1v2H3V0zm6 0h1v2H9V0z"/><path fill="currentColor" d="M13 7V1h-2v2H8V1H5v2H2V1H0v12h4v3h12V7h-3zm-9 5H1V5h11v2H4v5zm1-1.8l2.6 1.5L5 14.3v-4.1zm.7 4.8l2.8-2.8l1.5.9l1.5-.8l2.8 2.8H5.7zm9.3-.7l-2.6-2.6l2.6-1.4v4zm0-5.1l-5 2.7L5 9V8h10v1.2zm.4.4z"/>`
	calendarOInnerSVG            = `<path fill="currentColor" d="M14 1v3h-3V1H5v3H2V1H0v15h16V1h-2zm1 14H1V6h14v9z"/><path fill="currentColor" d="M3 0h1v3H3V0zm9 0h1v3h-1V0z"/>`
	calendarUserInnerSVG         = `<path fill="currentColor" d="M3 0h1v3H3V0zm8 0h1v3h-1V0z"/><path fill="currentColor" d="M9 14.1c0-.1 0-.1 0 0L1 14V6h13v1.2c.4.1.7.3 1 .6V1h-2v3h-3V1H5v3H2V1H0v14h9v-.9z"/><path fill="currentColor" d="M15 10a2 2 0 1 1-3.999.001A2 2 0 0 1 15 10z"/><path fill="currentColor" d="M13.9 12h-1.8c-1.1 0-2.1.9-2.1 2.1V16h6v-1.9c0-1.2-.9-2.1-2.1-2.1z"/>`
	cameraInnerSVG               = `<path fill="currentColor" d="M11 9a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/><path fill="currentColor" d="M11 4V1H5v3H0v9h5c.8.6 1.9 1 3 1s2.2-.4 3-1h5V4h-5zM6 2h4v2H6V2zm2 11c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4zm7-7h-2V5h2v1z"/>`
	carInnerSVG                  = `<path fill="currentColor" d="m15 6.1l-1.4-2.9c-.4-.7-1.1-1.2-1.9-1.2H4.3c-.8 0-1.5.5-1.9 1.2L1 6.1c-.6.1-1 .6-1 1.1v3.5c0 .6.4 1.1 1 1.2v2c0 .6.5 1.1 1.1 1.1H3c.5 0 1-.5 1-1.1V12h8v1.9c0 .6.5 1.1 1.1 1.1h.9c.6 0 1.1-.5 1.1-1.1v-2c.6-.1 1-.6 1-1.2V7.2c-.1-.5-.5-1-1.1-1.1zM4 8.4c0 .3-.3.6-.6.6H1.6c-.3 0-.6-.3-.6-.6v-.8c0-.3.3-.6.6-.6h1.8c.3 0 .6.3.6.6v.8zm6 2.6H6v-1h4v1zM2.1 6l1.2-2.4c.2-.4.6-.6 1-.6h7.4c.4 0 .8.2 1 .6L13.9 6H2.1zM15 8.4c0 .3-.3.6-.6.6h-1.8c-.3 0-.6-.3-.6-.6v-.8c0-.3.3-.6.6-.6h1.8c.3 0 .6.3.6.6v.8z"/>`
	caretDownInnerSVG            = `<path fill="currentColor" d="M3 4h10l-5 7z"/>`
	caretLeftInnerSVG            = `<path fill="currentColor" d="M11 3v10L4 8z"/>`
	caretRightInnerSVG           = `<path fill="currentColor" d="M5 13V3l7 5z"/>`
	caretSquareDownOInnerSVG     = `<path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/><path fill="currentColor" d="M4 6h8l-4 5z"/>`
	caretSquareLeftOInnerSVG     = `<path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/><path fill="currentColor" d="M10 4v8L5 8z"/>`
	caretSquareRightOInnerSVG    = `<path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/><path fill="currentColor" d="M5.9 12V4l5 4z"/>`
	caretSquareUpOInnerSVG       = `<path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/><path fill="currentColor" d="M12 10H4l4-5z"/>`
	caretUpInnerSVG              = `<path fill="currentColor" d="M13 12H3l5-7z"/>`
	cartInnerSVG                 = `<path fill="currentColor" d="M14 13.1V12H4.6l.6-1.1l9.2-.9L16 4H3.7L3 1H0v1h2.2l2.1 8.4L3 13v1.5c0 .8.7 1.5 1.5 1.5S6 15.3 6 14.5S5.3 13 4.5 13H12v1.5c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5c0-.7-.4-1.2-1-1.4z"/>`
	cartOInnerSVG                = `<path fill="currentColor" d="M14 13.1V12H4.6l.6-1.1l9.2-.9L16 4H3.7L3 1H0v1h2.2l2.1 8.4L3 13v1.5c0 .8.7 1.5 1.5 1.5S6 15.3 6 14.5S5.3 13 4.5 13H12v1.5c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5c0-.7-.4-1.2-1-1.4zM4 5h10.7l-1.1 4l-8.4.9L4 5z"/>`
	cashInnerSVG                 = `<path fill="currentColor" d="M16 14H2v-1h13V6h1v8z"/><path fill="currentColor" d="M13 4v7H1V4h12zm1-1H0v9h14V3z"/><path fill="currentColor" d="M3 6H2v3h1v1h4a2.5 2.5 0 1 1 0-5H3v1zm8 0V5H7a2.5 2.5 0 1 1 0 5h4V9h1V6h-1z"/>`
	chartInnerSVG                = `<path fill="currentColor" d="M0 15h16v1H0v-1z"/><path fill="currentColor" d="M0 0h1v16H0V0zm9 8L6.1 5L2 9v5h14V.9z"/>`
	chartGridInnerSVG            = `<path fill="currentColor" d="M0 9v7h16V9H0zm5 6H1v-1h4v1zm0-2H1v-1h4v1zm0-2H1v-1h4v1zm5 4H6v-1h4v1zm0-2H6v-1h4v1zm0-2H6v-1h4v1zm5 4h-4v-1h4v1zm0-2h-4v-1h4v1zm0-2h-4v-1h4v1zm1-3H0V0h1v7h15v1z"/><path fill="currentColor" d="M15 1.57L9.98 4.43L6.02 2.45L2 4.06v1.08l3.98-1.59l4.04 2.02L15 2.72V1.57z"/>`
	chartLineInnerSVG            = `<path fill="currentColor" d="M0 16h16V0h-1v2.6L11 6V0h-1v6.4l-4-.9V0H5v5.7L1 8.6V0H0zm5-2H1v-1.7l4-2.9V14zm5 0H6V8.7l.1-.1l3.9.9V14zm5 0h-4V9.7h.1L15 6.5V14z"/>`
	chartThreeDInnerSVG          = `<path fill="currentColor" d="M12 4V2L8 0L4 2v1L0 5v5l12 6l4-2V6zm-8 6.88l-3-1.5v-3.3l3 1.53v3.27zm0-4.39l-2.34-1.2L4 4.12v2.37zm4 6.39l-3-1.5V3.07l3 1.54v8.27zM5.66 2.29L8 1.12l2.34 1.17L8 3.49zM12 14.88l-3-1.5V7.07l3 1.54v6.27zm0-7.39l-2.34-1.2L12 5.12l2.34 1.17z"/>`
	chartTimelineInnerSVG        = `<path fill="currentColor" d="M16 13v-1H1V0H0v13h5v2H0v1h16v-1h-5v-2h5z"/><path fill="currentColor" d="M9 7L6 4L2 8v3h14V0L9 7z"/>`
	chatInnerSVG                 = `<path fill="currentColor" d="M14 14.2c0-.6 2-1.8 2-3.1c0-1.5-1.4-2.7-3.1-3.2c.7-.8 1.1-1.7 1.1-2.8C14 2.3 11.1 0 7.4 0C3.9 0 0 2.1 0 5.1c0 2.1 1.6 3.6 2.3 4.2c-.1 1.2-.6 1.7-.6 1.7L.5 12H2c1.6 0 2.9-.5 3.7-1.1v.2c0 2 2.2 3.6 5 3.6h.6c.4.5 1.7 1.4 3.4 1.4c.1-.1-.7-.5-.7-1.9zM7.4 1C10.5 1 13 2.9 13 5.1s-2.6 4.1-5.8 4.1H6.1l-.1.2c-.3.4-1.5 1.2-3.1 1.5c.1-.4.1-1 .1-1.8v-.3C2 8 .9 6.6.9 5.2C.9 3 4.1 1 7.4 1z"/>`
	checkInnerSVG                = `<path fill="currentColor" d="M7.3 14.2L.2 9l1.7-2.4l4.8 3.5l6.6-8.5l2.3 1.8z"/>`
	checkCircleInnerSVG          = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm-.9 11.7L2.9 7.6l1.4-1.4L7 8.9L12 4l1.4 1.4l-6.3 6.3z"/>`
	checkCircleOInnerSVG         = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M7.1 11.7L2.9 7.6l1.4-1.4l2.8 2.7L12 4l1.4 1.4z"/>`
	checkSquareInnerSVG          = `<path fill="currentColor" d="M13 .9L12 2H0v14h14V5.5l1.7-2L13 .9zM6.5 11.7L2.3 7.5l1.4-1.4l2.7 2.7L13 2.2l1.4 1.4l-7.9 8.1z"/>`
	checkSquareOInnerSVG         = `<path fill="currentColor" d="M14 6.2V14H2V2h10.5l1-1H1v14h14V5.2z"/><path fill="currentColor" d="M7.9 10.9L3.7 6.7l1.5-1.4l2.7 2.8l6.7-6.7L16 2.8z"/>`
	chevronCircleDownInnerSVG    = `<path fill="currentColor" d="M0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8zm11.6-2.8L13 6.6l-5 5l-5-5l1.4-1.4L8 8.8l3.6-3.6z"/>`
	chevronCircleDownOInnerSVG   = `<path fill="currentColor" d="m13 6.6l-5 5l-5-5l1.4-1.4L8 8.8l3.6-3.6z"/><path fill="currentColor" d="M1 8c0-3.9 3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7zM0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8z"/>`
	chevronCircleLeftInnerSVG    = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm2.8 11.6L9.4 13l-5-5l5-5l1.4 1.4L7.2 8l3.6 3.6z"/>`
	chevronCircleLeftOInnerSVG   = `<path fill="currentColor" d="m9.4 13l-5-5l5-5l1.4 1.4L7.2 8l3.6 3.6z"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/>`
	chevronCircleRightInnerSVG   = `<path fill="currentColor" d="M8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8zM5.2 4.4L6.6 3l5 5l-5 5l-1.4-1.4L8.8 8L5.2 4.4z"/>`
	chevronCircleRightOInnerSVG  = `<path fill="currentColor" d="m6.6 13l5-5l-5-5l-1.4 1.4L8.8 8l-3.6 3.6z"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/>`
	chevronCircleUpInnerSVG      = `<path fill="currentColor" d="M16 8c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8zM4.4 10.8L3 9.4l5-5l5 5l-1.4 1.4L8 7.2l-3.6 3.6z"/>`
	chevronCircleUpOInnerSVG     = `<path fill="currentColor" d="m3 9.4l5-5l5 5l-1.4 1.4L8 7.2l-3.6 3.6z"/><path fill="currentColor" d="M15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7zm1 0c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8z"/>`
	chevronDownInnerSVG          = `<path fill="currentColor" d="m8 13.1l-8-8l2.1-2.2L8 8.8l5.9-5.9L16 5.1z"/>`
	chevronDownSmallInnerSVG     = `<path fill="currentColor" d="M8 12L1.68 5.68L3.35 4L8 8.65L12.65 4l1.67 1.68L8 12z"/>`
	chevronLeftInnerSVG          = `<path fill="currentColor" d="m2.9 8l8-8l2.2 2.1L7.2 8l5.9 5.9l-2.2 2.1z"/>`
	chevronLeftSmallInnerSVG     = `<path fill="currentColor" d="m4 8l6.32-6.32L12 3.35L7.35 8L12 12.65l-1.68 1.67L4 8z"/>`
	chevronRightInnerSVG         = `<path fill="currentColor" d="m13.1 8l-8 8l-2.2-2.1L8.8 8L2.9 2.1L5.1 0z"/>`
	chevronRightSmallInnerSVG    = `<path fill="currentColor" d="M12 8L5.68 1.68L4 3.35L8.65 8L4 12.65l1.68 1.67L12 8z"/>`
	chevronUpInnerSVG            = `<path fill="currentColor" d="m8 2.9l8 8l-2.1 2.2L8 7.2l-5.9 5.9L0 10.9z"/>`
	chevronUpSmallInnerSVG       = `<path fill="currentColor" d="m8 4l-6.32 6.32L3.35 12L8 7.35L12.65 12l1.67-1.68L8 4z"/>`
	childInnerSVG                = `<path fill="currentColor" d="M10 5a2 2 0 1 1-3.999.001A2 2 0 0 1 10 5z"/><path fill="currentColor" d="m12.79 10.32l-2.6-2.63A2.311 2.311 0 0 0 8.54 7H7.469c-.648 0-1.235.264-1.659.69l-2.6 2.63a.73.73 0 1 0 .998 1.003L6 9.53V16h1.5v-4h1v4H10V9.53l1.75 1.8a.73.73 0 1 0 1.041-1.009z"/>`
	circleInnerSVG               = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/>`
	circleThinInnerSVG           = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/>`
	clipboardInnerSVG            = `<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/>`
	clipboardCheckInnerSVG       = `<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/><path fill="currentColor" d="m7.39 12.47l-3-2.73l1.35-1.48L7.32 9.7l2.87-2.9l1.42 1.4l-4.22 4.27z"/>`
	clipboardCrossInnerSVG       = `<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/><path fill="currentColor" d="M11 8H9V6H7v2H5v2h2v2h2v-2h2z"/>`
	clipboardHeartInnerSVG       = `<path fill="currentColor" d="M9.5 7c-.6 0-1.1.6-1.5 1c-.4-.4-.9-1-1.5-1c-1.5 0-2.1 1.9-1 2.9L8 12l2.5-2.1c1.1-1 .5-2.9-1-2.9z"/><path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/>`
	clipboardPulseInnerSVG       = `<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/><path fill="currentColor" d="M9.3 13c-.2 0-.3-.1-.4-.3l-.8-4.8l-.7 3.1c0 .1-.1.2-.3.3c-.1 0-.3 0-.4-.1l-1-1.3H4.4c-.2 0-.4-.2-.4-.4s.2-.4.4-.4H6c.1 0 .2.1.3.1l.6.8l.9-4.3c0-.2.2-.3.4-.3s.3.2.3.4l.9 5.3l.6-1.7c.1-.1.2-.2.3-.2h1.3c.2 0 .4.2.4.4s-.2.4-.4.4h-1l-1 2.9s-.2.1-.3.1z"/>`
	clipboardTextInnerSVG        = `<path fill="currentColor" d="M4 6h8v1H4V6zm0 2h8v1H4V8zm0 2h5v1H4v-1z"/><path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/>`
	clipboardUserInnerSVG        = `<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/><path fill="currentColor" d="M8 6C5.5 6 6.7 9.2 6.7 9.2c.3.4.7.4.7.6c0 .3-.3.3-.6.4c-.5.1-.9-.1-1.4.8c-.3.4-.4 2-.4 2h6s-.1-1.6-.4-2c-.4-.8-.9-.7-1.4-.8c-.3 0-.6-.1-.6-.4s.3-.2.6-.6C9.3 9.2 10.5 6 8 6z"/>`
	clockInnerSVG                = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm0 14c-3.3 0-6-2.7-6-6s2.7-6 6-6s6 2.7 6 6s-2.7 6-6 6z"/><path fill="currentColor" d="M8 3H7v6h5V8H8z"/>`
	closeInnerSVG                = `<path fill="currentColor" d="M15.1 3.1L12.9.9L8 5.9L3.1.9L.9 3.1l5 4.9l-5 4.9l2.2 2.2l4.9-5l4.9 5l2.2-2.2l-5-4.9z"/>`
	closeBigInnerSVG             = `<path fill="currentColor" d="m16 0l-1 .01L8 7L1 .01L0 0v1l7 7l-7 7v1h1l7-7l7 7h1v-1L9 8l7-7V0z"/>`
	closeCircleInnerSVG          = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm4.2 10.8l-1.4 1.4L8 9.4l-2.8 2.8l-1.4-1.4L6.6 8L3.8 5.2l1.4-1.4L8 6.6l2.8-2.8l1.4 1.4L9.4 8l2.8 2.8z"/>`
	closeCircleOInnerSVG         = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M12.2 10.8L9.4 8l2.8-2.8l-1.4-1.4L8 6.6L5.2 3.8L3.8 5.2L6.6 8l-2.8 2.8l1.4 1.4L8 9.4l2.8 2.8z"/>`
	closeSmallInnerSVG           = `<path fill="currentColor" d="m12.96 4.46l-1.42-1.42L8 6.59L4.46 3.04L3.04 4.46L6.59 8l-3.55 3.54l1.42 1.42L8 9.41l3.54 3.55l1.42-1.42L9.41 8l3.55-3.54z"/>`
	cloudInnerSVG                = `<path fill="currentColor" d="M14 13c1.1 0 2-.9 2-2s-.9-2-2-2h-.1c0-.3.1-.6.1-1c0-2.2-1.8-4-4-4c-.8 0-1.5.2-2.2.6C7.5 3.7 6.6 3 5.5 3C4.1 3 3 4.1 3 5.5c0 .6.2 1.1.6 1.6C3.4 7 3.2 7 3 7c-1.7 0-3 1.3-3 3s1.3 3 3 3h11z"/>`
	cloudDownloadInnerSVG        = `<path fill="currentColor" d="M14 10h-.1c0-.3.1-.6.1-1c0-2.2-1.8-4-4-4V1H6v3.1C5.8 4 5.7 4 5.5 4C4.1 4 3 5.1 3 6.5c0 .6.2 1.1.6 1.6C3.4 8 3.2 8 3 8c-1.7 0-3 1.3-3 3s1.3 3 3 3h11c1.1 0 2-.9 2-2s-.9-2-2-2zm-6 1.4L5.1 8H7V2h2v6h1.9L8 11.4z"/>`
	cloudDownloadOInnerSVG       = `<path fill="currentColor" d="M14.1 9.8v-.6c0-2.4-1.9-4.3-4.2-4.3c-.3.1-.6.1-.9.1V2H7v2.4c-.4-.3-.9-.4-1.3-.4c-1.6 0-2.9 1.3-2.9 2.9c0 .3.1.6.2.9c-1.6.2-3 1.8-3 3.6C0 13.3 1.5 15 3.3 15h10.3c1.4 0 2.4-1.5 2.4-2.7s-.8-2.3-1.9-2.5zm-.5 4.2H3.3C2.1 14 1 12.7 1 11.4s1.1-2.6 2.3-2.6h.4l1.4.2l-.9-1c-.2-.3-.4-.7-.4-1.2c0-1 .8-1.8 1.8-1.8c.5 0 1 .2 1.3.6V8H5l3 4l3-4H9V6.1c.3-.1.6-.1.9-.1c1.8 0 3.2 1.5 3.2 3.3c0 .3 0 .6-.1.9l-.2.6l.8.1c.7 0 1.4.7 1.4 1.5c0 .7-.6 1.6-1.4 1.6z"/>`
	cloudOInnerSVG               = `<path fill="currentColor" d="M14.1 8.9v-.6c0-2.4-1.9-4.3-4.2-4.3c-.6 0-1.2.1-1.8.4c-.5-.7-1.5-1.2-2.4-1.2c-1.6 0-2.9 1.2-2.9 2.8c0 .3.1.6.2.9c-1.6.2-3 1.8-3 3.5C0 12.3 1.5 14 3.3 14h10.3c1.4 0 2.4-1.4 2.4-2.6s-.8-2.2-1.9-2.5zm-.5 4.1H3.3C2.1 13 1 11.8 1 10.5S2.1 8 3.3 8h.4l1.3.3l-.8-1.2c-.2-.3-.4-.7-.4-1.1c0-1 .8-1.8 1.8-1.8c.8 0 1.5.5 1.7 1.2l.3.6l.5-.3c.5-.3 1.1-.5 1.8-.5c1.8 0 3.2 1.5 3.2 3.3c0 .3 0 .6-.1.9l-.2.6h.8c.7 0 1.4.7 1.4 1.5c0 .6-.6 1.5-1.4 1.5z"/>`
	cloudUploadInnerSVG          = `<path fill="currentColor" d="M14 10h-.1c0-.3.1-.6.1-1c0-1.6-1-3-2.4-3.6L8 1L5.5 4C4.1 4 3 5.1 3 6.5c0 .6.2 1.1.6 1.6C3.4 8 3.2 8 3 8c-1.7 0-3 1.3-3 3s1.3 3 3 3h11c1.1 0 2-.9 2-2s-.9-2-2-2zM9 6v6H7V6H5.1L8 2.6L10.9 6H9z"/>`
	cloudUploadOInnerSVG         = `<path fill="currentColor" d="M14.1 10.9v-.6c0-2.4-1.9-4.3-4.2-4.3c-.3 0-.6 0-.9.1V4h2L8 0L5 4h2v1.5c-.4-.2-.9-.3-1.3-.3c-1.6 0-2.9 1.2-2.9 2.8c0 .3.1.6.2.9c-1.6.2-3 1.8-3 3.5C0 14.3 1.5 16 3.3 16h10.3c1.4 0 2.4-1.4 2.4-2.6s-.8-2.2-1.9-2.5zm-.5 4.1H3.3C2.1 15 1 13.8 1 12.5S2.1 10 3.3 10h.4l1.3.3l-.8-1.2c-.2-.3-.4-.7-.4-1.1c0-1 .8-1.8 1.8-1.8c.5 0 1 .2 1.3.6V10h2V7.2c.3-.1.6-.1.9-.1c1.8 0 3.2 1.5 3.2 3.3c0 .3 0 .6-.1.9l-.2.6h.8c.7 0 1.4.7 1.4 1.5c.1.7-.5 1.6-1.3 1.6z"/>`
	clusterInnerSVG              = `<path fill="currentColor" d="M14 12a1.993 1.993 0 0 0-1.008.305L10.78 10.15a3.439 3.439 0 0 0 .74-1.993L13.09 8a1.49 1.49 0 1 0-.089-.768l-1.591.128a3.512 3.512 0 0 0-1.978-2.521L9.74 4H10a2 2 0 1 0-1.01-.265l-.27.855a3.31 3.31 0 0 0-.754-.084c-.83 0-1.59.296-2.181.789L2.791 2.291a1.5 1.5 0 1 0-1.291.71c.281-.001.544-.079.767-.214L5.26 5.791a3.446 3.446 0 0 0-.76 2.168v.203l-.66.11a2 2 0 1 0 .161.786L4 8.999l.63-.097a3.522 3.522 0 0 0 1.466 1.992l-.556 1.188a1.947 1.947 0 0 0-.539-.08h-.017a2 2 0 1 0 1.231.423l.566-1.153c.364.146.787.231 1.229.231c.847 0 1.621-.311 2.216-.824l2.176 2.124a2 2 0 1 0 1.6-.8zm-9 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm3-4.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z"/>`
	codeInnerSVG                 = `<path fill="currentColor" d="M5.2 14L9.7 2h1.1L6.3 14zm5.9-1h1.2L16 8l-3.7-5H11l3.8 5zm-6.2 0H3.7L0 8l3.7-5H5L1.2 8z"/>`
	coffeeInnerSVG               = `<path fill="currentColor" d="m14 13l-4 1H4l-4-1v-1h14zm.7-10H13V2H1v5c0 1.5.8 2.8 2 3.4v.6h8v-.6c.9-.5 1.6-1.4 1.9-2.4h.1c2.3 0 2.9-2 3-3.5c.1-.8-.5-1.5-1.3-1.5zM13 7V4h1.7c.1 0 .2.1.2.1s.1.1.1.3C14.8 7 13.4 7 13 7z"/>`
	cogInnerSVG                  = `<path fill="currentColor" d="M16 9V7l-1.7-.6c-.2-.6-.4-1.2-.7-1.8l.8-1.6L13 1.6l-1.6.8c-.5-.3-1.1-.6-1.8-.7L9 0H7l-.6 1.7c-.6.2-1.2.4-1.7.7l-1.6-.8l-1.5 1.5l.8 1.6c-.3.5-.5 1.1-.7 1.7L0 7v2l1.7.6c.2.6.4 1.2.7 1.8L1.6 13L3 14.4l1.6-.8c.5.3 1.1.6 1.8.7L7 16h2l.6-1.7c.6-.2 1.2-.4 1.8-.7l1.6.8l1.4-1.4l-.8-1.6c.3-.5.6-1.1.7-1.8L16 9zm-8 3c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4z"/><path fill="currentColor" d="M10.6 7.9a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0z"/>`
	cogOInnerSVG                 = `<path fill="currentColor" d="m15.2 6l-1.1-.2c-.1-.2-.1-.4-.2-.6l.6-.9l.5-.7L12.4 1l-.7.5l-.9.6c-.2-.1-.4-.1-.6-.2L10 .8L9.8 0H6.2L6 .8l-.2 1.1c-.2.1-.4.1-.6.2l-.9-.6l-.7-.4l-2.5 2.5l.5.7l.6.9c-.2.2-.2.4-.3.6L.8 6l-.8.2v3.6l.8.2l1.1.2c.1.2.1.4.2.6l-.6.9l-.5.7L3.6 15l.7-.5l.9-.6c.2.1.4.1.6.2l.2 1.1l.2.8h3.6l.2-.8l.2-1.1c.2-.1.4-.1.6-.2l.9.6l.7.5l2.6-2.6l-.5-.7l-.6-.9c.1-.2.2-.4.2-.6l1.1-.2l.8-.2V6.2l-.8-.2zM15 9l-1.7.3c-.1.5-.3 1-.6 1.5l.9 1.4l-1.4 1.4l-1.4-.9c-.5.3-1 .5-1.5.6L9 15H7l-.3-1.7c-.5-.1-1-.3-1.5-.6l-1.4.9l-1.4-1.4l.9-1.4c-.3-.5-.5-1-.6-1.5L1 9V7l1.7-.3c.1-.5.3-1 .6-1.5l-1-1.4l1.4-1.4l1.4.9c.5-.3 1-.5 1.5-.6L7 1h2l.3 1.7c.5.1 1 .3 1.5.6l1.4-.9l1.4 1.4l-.9 1.4c.3.5.5 1 .6 1.5L15 7v2z"/><path fill="currentColor" d="M8 4.5C6.1 4.5 4.5 6.1 4.5 8s1.6 3.5 3.5 3.5s3.5-1.6 3.5-3.5S9.9 4.5 8 4.5zm0 6c-1.4 0-2.5-1.1-2.5-2.5S6.6 5.5 8 5.5s2.5 1.1 2.5 2.5s-1.1 2.5-2.5 2.5z"/>`
	cogsInnerSVG                 = `<path fill="currentColor" d="M12 7V5l-1.2-.4c-.1-.3-.2-.7-.4-1l.6-1.2l-1.5-1.3l-1.1.5c-.3-.2-.6-.3-1-.4L7 0H5l-.4 1.2c-.3.1-.7.2-1 .4l-1.1-.5l-1.4 1.4l.6 1.2c-.2.3-.3.6-.4 1L0 5v2l1.2.4c.1.3.2.7.4 1l-.5 1.1l1.4 1.4l1.2-.6c.3.2.6.3 1 .4L5 12h2l.4-1.2c.3-.1.7-.2 1-.4l1.2.6L11 9.6l-.6-1.2c.2-.3.3-.6.4-1L12 7zM3 6c0-1.7 1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3s-3-1.3-3-3z"/><path fill="currentColor" d="M7.5 6a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 7.5 6zM16 3V2h-.6c0-.2-.1-.4-.2-.5l.4-.4l-.7-.7l-.4.4c-.2-.1-.3-.2-.5-.2V0h-1v.6c-.2 0-.4.1-.5.2l-.4-.4l-.7.7l.4.4c-.1.2-.2.3-.2.5H11v1h.6c0 .2.1.4.2.5l-.4.4l.7.7l.4-.4c.2.1.3.2.5.2V5h1v-.6c.2 0 .4-.1.5-.2l.4.4l.7-.7l-.4-.4c.1-.2.2-.3.2-.5h.6zm-2.5.5c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1zm1.9 8.3c-.1-.3-.2-.6-.4-.9l.3-.6l-.7-.7l-.5.4c-.3-.2-.6-.3-.9-.4L13 9h-1l-.2.6c-.3.1-.6.2-.9.4l-.6-.3l-.7.7l.3.6c-.2.3-.3.6-.4.9L9 12v1l.6.2c.1.3.2.6.4.9l-.3.6l.7.7l.6-.3c.3.2.6.3.9.4l.1.5h1l.2-.6c.3-.1.6-.2.9-.4l.6.3l.7-.7l-.4-.5c.2-.3.3-.6.4-.9l.6-.2v-1l-.6-.2zM12.5 14c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5z"/>`
	coinPilesInnerSVG            = `<path fill="currentColor" d="M10.5 0C7.46 0 5 .88 5 2v2c-3 .1-5 .94-5 2v6c0 1.09 2.46 2 5.5 2h.067c.732 0 1.45-.055 2.153-.16c.698 1.305 2.094 2.158 3.69 2.158a4.382 4.382 0 0 0 4.224-3.217c.209-.199.344-.442.367-.717V2c0-1.12-2.46-2-5.5-2zm-5 5C8 5 10 5.45 10 6S8 7 5.5 7S1 6.55 1 6s2-1 4.5-1zm0 8c-2.71 0-4.25-.71-4.5-1v-.8a10.405 10.405 0 0 0 4.522.799c.508-.001 1.03-.03 1.544-.085c-.043.371.022.712.123 1.037c-.452.021-.967.051-1.488.051L5.49 13zm1.57-2.09c-.467.057-1.008.09-1.556.09H5.5c-2.709 0-4.249-.71-4.499-1v-.84a10.41 10.41 0 0 0 4.518.84a14.496 14.496 0 0 0 1.897-.128c-.197.306-.291.654-.342 1.015zM5.5 9C2.79 9 1.25 8.29 1 8v-.9a10.41 10.41 0 0 0 4.518.84a10.548 10.548 0 0 0 4.551-.866l-.068.366a4.397 4.397 0 0 0-1.935 1.304C7.314 8.909 6.455 9 5.575 9h-.077zm5.91 6a3.41 3.41 0 1 1 0-6.82a3.41 3.41 0 0 1 0 6.82zM15 8c-.175.167-.385.3-.617.386c-.288-.244-.6-.46-.938-.634a7.615 7.615 0 0 0 1.593-.61L15 8zm0-2c-.24.31-1.61.94-4 1V6h.011a9.963 9.963 0 0 0 4.053-.855L15 6zm0-2c-.25.33-1.79 1-4.5 1h-.23a9.073 9.073 0 0 0-4.169-1H6v-.9a10.41 10.41 0 0 0 4.518.84a10.548 10.548 0 0 0 4.551-.866L15.001 4zm-4.5-1C8 3 6 2.55 6 2s2-1 4.5-1s4.5.45 4.5 1s-2 1-4.5 1z"/><path fill="currentColor" d="M10.5 11h.5v3h1V9h-.5l-1 2z"/>`
	coinsInnerSVG                = `<path fill="currentColor" d="M11.5 0A4.5 4.5 0 0 0 7 4.5c.004.261.029.513.074.758a4.647 4.647 0 0 0-1.591-.261a5.51 5.51 0 1 0 5.266 3.884c.23.077.484.099.742.099A4.49 4.49 0 0 0 11.5 0zM10 10.5A4.5 4.5 0 1 1 5.5 6a4.51 4.51 0 0 1 4.5 4.499zM12.5 7h-2v-.5h.5v-3h-.5l1-1.5h.5v4.5h.5V7z"/><path fill="currentColor" d="M5.63 8a1.258 1.258 0 0 1 1.371 1.255L7 9.302C7 11 5.14 12 5.14 12h1.37v-.5H7V13H4v-1s2-1.27 2-2.33C6 9.3 6 9 5.58 9c-.69 0-.65 1-.65 1H4s-.23-2 1.63-2z"/>`
	comboboxInnerSVG             = `<path fill="currentColor" d="M15 4H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1zm-5 7H1V5h9v6zm3-2.6L11 7h4l-2 1.4z"/><path fill="currentColor" d="M2 6h1v4H2V6z"/>`
	commentInnerSVG              = `<path fill="currentColor" d="M8 1C3.6 1 0 3.5 0 6.5c0 2 2 3.8 4 4.8c0 2.1-2 2.8-2 2.8c2.8 0 4.4-1.3 5.1-2.1H8c4.4 0 8-2.5 8-5.5S12.4 1 8 1z"/>`
	commentEllipsisInnerSVG      = `<path fill="currentColor" d="M8 1C3.6 1 0 3.5 0 6.5c0 2 2 3.8 4 4.8c0 2.1-2 2.8-2 2.8c2.8 0 4.4-1.3 5.1-2.1H8c4.4 0 8-2.5 8-5.5S12.4 1 8 1zM5 8c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1zm3 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1zm3 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1z"/>`
	commentEllipsisOInnerSVG     = `<path fill="currentColor" d="M3 11.2c0 .1 0 .1 0 0c0 .1 0 .1 0 0zM8.3 1C3.9 1 0 3.6 0 6.6c0 2 1.1 3.7 3 4.7s0 .1 0 .1c-.1 1.3-.9 1.7-.9 1.7L.3 14h2c2.5 0 4.3-1.1 5.1-1.9h.8c4.3 0 7.8-2.5 7.8-5.6S12.6 1 8.3 1zm-.1 10.1H7.1l-.2.2c-.5.5-1.6 1.4-3.3 1.7c.3-.5.5-1.1.5-2v-.3l-.3-.1C2 9.7 1 8.3 1 6.6C1 4.2 4.5 2 8.3 2C12 2 15 4 15 6.6c0 2.4-3.1 4.5-6.8 4.5z"/><path fill="currentColor" d="M6 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`
	commentOInnerSVG             = `<path fill="currentColor" d="M3 11.2c0 .1 0 .1 0 0c0 .1 0 .1 0 0zM8.3 1C3.9 1 0 3.6 0 6.6c0 2 1.1 3.7 3 4.7s0 .1 0 .1c-.1 1.3-.9 1.7-.9 1.7L.3 14h2c2.5 0 4.3-1.1 5.1-1.9h.8c4.3 0 7.8-2.5 7.8-5.6S12.6 1 8.3 1zm-.1 10.1H7l-.2.2c-.5.5-1.6 1.4-3.3 1.7c.3-.5.5-1.1.5-2v-.3l-.3-.1C1.9 9.7 1 8.3 1 6.6C1 4.2 4.5 2 8.3 2C12 2 15 4 15 6.6c0 2.4-3.1 4.5-6.8 4.5z"/>`
	commentsInnerSVG             = `<path fill="currentColor" d="M16 11.1c0-1.5-1.5-2.8-3.2-3.3c-1.3 1.5-3.9 2.4-6.4 2.4h-.5c-.1.3-.1.5-.1.8c0 2 2.2 3.6 5 3.6h.6c.4.5 1.7 1.4 3.4 1.4c0 0-.8-.4-.8-1.8c0-.6 2-1.8 2-3.1z"/><path fill="currentColor" d="M13 4.6C13 2.1 10.2 0 6.6 0S0 2.1 0 4.6c0 1.7 2 3.2 3 4C3 10.4 1.6 11 1.6 11c2.3 0 3.6-1.1 4.2-1.8h.8c3.5.1 6.4-2 6.4-4.6z"/>`
	commentsOInnerSVG            = `<path fill="currentColor" d="M14.2 14c.6-.5 1.8-1.6 1.8-3.2c0-1.4-1.2-2.6-2.8-3.3c.5-.6.8-1.5.8-2.4C14 2.3 11.1 0 7.4 0C3.9 0 0 2.1 0 5.1c0 2.1 1.6 3.6 2.3 4.2c-.1 1.2-.6 1.7-.6 1.7L.5 12H2c1.2 0 2.2-.3 3-.7c.3 1.9 2.5 3.4 5.3 3.4h.5c.6.5 1.8 1.3 3.5 1.3h1.4l-1.1-.9s-.3-.3-.4-1.1zm-3.9-.3C8 13.7 6 12.4 6 10.9v-.2c.2-.2.4-.3.5-.5h.7c2.1 0 4-.7 5.2-1.9c1.5.5 2.6 1.5 2.6 2.5s-.9 2-1.7 2.5l-.3.2v.3c0 .5.2.8.3 1.1c-1-.2-1.7-.7-1.9-1l-.1-.2h-1zM7.4 1C10.5 1 13 2.9 13 5.1s-2.6 4.1-5.8 4.1H6.1l-.1.2c-.3.4-1.5 1.2-3.1 1.5c.1-.4.1-1 .1-1.8v-.3C2 8 .9 6.6.9 5.2C.9 3 4.1 1 7.4 1z"/>`
	compileInnerSVG              = `<path fill="currentColor" d="M1 12h4v4H1v-4zm5 0h4v4H6v-4zm5 0h4v4h-4v-4zM1 7h4v4H1V7zm0-5h4v4H1V2zm5 5h4v4H6V7zm1-6h4v4H7V1zm4 6h4v4h-4V7zm2-7h3v3h-3V0z"/>`
	compressInnerSVG             = `<path fill="currentColor" d="m5.3 9.3l-5 5l1.4 1.4l5-5L8 12V8H4zm10.4-7.6L14.3.3l-4 4L9 3v4h4l-1.3-1.3z"/>`
	compressSquareInnerSVG       = `<path fill="currentColor" d="M12 0H0v12l1-1V1h10zM4 16h12V4l-1 1v10H5z"/><path fill="currentColor" d="M7 9H2l1.8 1.8L0 14.6L1.4 16l3.8-3.8L7 14zm9-7.6L14.6 0l-3.8 3.8L9 2v5h5l-1.8-1.8z"/>`
	connectInnerSVG              = `<path fill="currentColor" d="M12 10c-.8 0-1.4.3-2 .8L6.8 9c.1-.3.2-.7.2-1s-.1-.7-.2-1L10 5.2c.6.5 1.2.8 2 .8c1.7 0 3-1.3 3-3s-1.3-3-3-3s-3 1.3-3 3v.5L5.5 5.4C5.1 5.2 4.6 5 4 5C2.4 5 1 6.3 1 8c0 1.6 1.4 3 3 3c.6 0 1.1-.2 1.5-.4L9 12.5v.5c0 1.7 1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3z"/>`
	connectOInnerSVG             = `<path fill="currentColor" d="M12.5 9c-1 0-1.8.4-2.4 1L6.9 8.3c.1-.3.1-.5.1-.8v-.4l2.9-1.3c.6.7 1.5 1.2 2.6 1.2C14.4 7 16 5.4 16 3.5S14.4 0 12.5 0S9 1.6 9 3.5v.4L6.1 5.2C5.5 4.5 4.6 4 3.5 4C1.6 4 0 5.6 0 7.5S1.6 11 3.5 11c1 0 1.8-.4 2.4-1L9 11.7v.8c0 1.9 1.6 3.5 3.5 3.5s3.5-1.6 3.5-3.5S14.4 9 12.5 9zm0-8C13.9 1 15 2.1 15 3.5S13.9 6 12.5 6S10 4.9 10 3.5S11.1 1 12.5 1zm-9 9C2.1 10 1 8.9 1 7.5S2.1 5 3.5 5S6 6.1 6 7.5S4.9 10 3.5 10zm9 5c-1.4 0-2.5-1.1-2.5-2.5s1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5s-1.1 2.5-2.5 2.5z"/>`
	controllerInnerSVG           = `<path fill="currentColor" d="m5.951.249l.981-.195l.195.981l-.981.195l-.195-.981zm2.926 14.717l.981-.195l.195.981l-.981.195l-.195-.981zM.055 9.071l.981-.195l.195.981l-.981.195l-.195-.981zm14.718-2.926l.981-.195l.195.981l-.981.195l-.195-.981zm-3.302-4.248l.556-.831l.831.556l-.556.831l-.831-.556zM3.139 14.441l.56-.83l.83.56l-.56.83l-.83-.56zM1.069 3.989l.56-.83l.83.56l-.56.83l-.83-.56zm12.478 8.31l.556-.831l.831.556l-.556.831l-.831-.556zM8.875 1.039L9.07.058l.981.195l-.195.981l-.981-.195zM5.953 15.745l.195-.981l.981.195l-.195.981l-.981-.195zM.061 6.931l.195-.981l.981.195l-.195.981l-.981-.195zm14.706 2.923l.195-.981l.981.195l-.195.981l-.981-.195zM3.139 1.628l.831-.556l.556.831l-.831.556l-.556-.831zm8.338 12.473l.831-.556l.556.831l-.831.556l-.556-.831zM1.071 12.033l.831-.556l.556.831l-.831.556l-.556-.831zM13.539 3.63l.83-.56l.56.83l-.83.56l-.56-.83z"/><path fill="currentColor" d="M14 8a5.99 5.99 0 0 0-2.258-4.681L8.42 8.31l-.84-.59l3.32-5A5.93 5.93 0 0 0 8 1.973a6 6 0 1 0 6 6.029z"/>`
	copyInnerSVG                 = `<path fill="currentColor" d="M6 0v3h3z"/><path fill="currentColor" d="M9 4H5V0H0v12h9zm4 0v3h3z"/><path fill="currentColor" d="M12 4h-2v9H7v3h9V8h-4z"/>`
	copyOInnerSVG                = `<path fill="currentColor" d="M13 3h-3L7 0H0v13h6v3h10V6l-3-3zM7 1l2 2H7V1zM1 12V1h5v3h3v8H1zm14 3H7v-2h3V4h2v3h3v8zm-2-9V4l2 2h-2z"/>`
	copyrightInnerSVG            = `<path fill="currentColor" d="M8 1.5c3.6 0 6.5 2.9 6.5 6.5s-2.9 6.5-6.5 6.5S1.5 11.6 1.5 8S4.4 1.5 8 1.5zM8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M9.9 10.3c-.5.4-1.2.7-1.9.7c-1.7 0-3-1.3-3-3s1.3-3 3-3c.8 0 1.6.3 2.1.9l1.1-1.1c-.8-.8-2-1.3-3.2-1.3c-2.5 0-4.5 2-4.5 4.5s2 4.5 4.5 4.5c1.1 0 2-.4 2.8-1l-.9-1.2z"/>`
	cornerLowerLeftInnerSVG      = `<path fill="currentColor" d="M16 16L0 0v16z"/>`
	cornerLowerRightInnerSVG     = `<path fill="currentColor" d="M16 16H0L16 0z"/>`
	cornerUpperLeftInnerSVG      = `<path fill="currentColor" d="M0 16L16 0H0z"/>`
	cornerUpperRightInnerSVG     = `<path fill="currentColor" d="M16 16L0 0h16z"/>`
	creditCardInnerSVG           = `<path fill="currentColor" d="M0 2v12h16V2H0zm15 11H1V8h14v5zm0-8H1V3h14v2z"/><path fill="currentColor" d="M10 11h3v1h-3v-1zm-8 0h6v1H2v-1z"/>`
	cropInnerSVG                 = `<path fill="currentColor" d="M16 .7V0h-.7l-3 3H5V0H3v3H0v2h3v8h8v3h2v-3h3v-2h-3V3.7l3-3zM5 5h5.3L5 10.3V5zm6 6H5.7L11 5.7V11z"/>`
	crossCutleryInnerSVG         = `<path fill="currentColor" d="M10.9 8.6c.6-.1 1.2-.4 1.6-.9l3.1-3.1c.4-.4.4-1 0-1.4l-.1-.2l-3 3c-.2.2-.6.2-.9 0s-.2-.6 0-.9l2.6-2.6c.2-.2.2-.6 0-.9c-.2-.2-.6-.2-.9 0l-2.6 2.6c-.2.2-.6.2-.9 0c-.2-.2-.2-.6 0-.9l3-3l-.1-.1c-.4-.4-1-.4-1.4 0L8.2 3.5c-.4.4-.7 1-.8 1.6L2.5.3c-.4-.4-1-.3-1.3 0L1 .5C-.4 1.9.1 4.7 2.5 7.1l.8.8c.4.4.9.7 1.5.8c-.5.4-.8.8-.8.8L.6 12.9c-.7.7-.7 1.9 0 2.6s1.9.7 2.6 0L6.5 12c.2-.2.7-.8 1.3-1.5c.3.4.5.6.5.6l4.3 4.3c.7.7 1.9.7 2.6 0s.7-1.9 0-2.6l-4.3-4.2z"/>`
	crosshairsInnerSVG           = `<path fill="currentColor" d="M7.5 0h1v4L8 6l-.5-2V0zm1 16h-1v-4l.5-2l.5 2v4zM16 7.5v1h-4L10 8l2-.5h4zm-16 1v-1h4L6 8l-2 .5H0z"/><path fill="currentColor" d="M8 2.5a5.5 5.5 0 1 1 0 11A5.5 5.5 0 0 1 2.5 8a5.51 5.51 0 0 1 5.499-5.5zM8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1z"/>`
	cssInnerSVG                  = `<path fill="currentColor" d="M4.1 11C5.5 11 6 10 6 10l-.8-.5s-.3.5-1 .5S3 9.1 3 7.8C3 6.6 3.6 6 4.2 6c.5 0 .9.4.9.4l.8-.6S5.2 5 4.2 5C3.1 5 2 5.9 2 7.8S2.9 11 4.1 11zm4.6-1.1c-.3.1-.7 0-1-.4l-.8.5c.4.6 1 1 1.6 1c.1 0 .3 0 .4-.1c.7-.2 1.1-.8 1.1-1.6c0-1.2-.8-1.6-1.3-1.8c-.5-.3-.7-.4-.7-.8s.1-.7.6-.7c.3 0 .6.4.6.4l.8-.6c-.2-.3-.7-.8-1.4-.8C7.7 5 7 5.6 7 6.6c0 1.1.7 1.5 1.2 1.8c.6.2.8.4.8.9c0 .3 0 .6-.3.6zm4 0c-.3.1-.7 0-1-.4l-.8.5c.4.6 1 1 1.6 1c.1 0 .3 0 .4-.1c.7-.2 1.1-.8 1.1-1.6c0-1.2-.8-1.6-1.3-1.8c-.5-.3-.7-.4-.7-.8s.1-.7.6-.7c.3 0 .6.4.6.4l.8-.6c-.2-.3-.7-.8-1.4-.8c-.9 0-1.6.6-1.6 1.6c0 1.1.7 1.5 1.2 1.8c.6.2.8.4.8.9c0 .3 0 .6-.3.6zM0 0v16h16V0H0zm15 15H1V1h14v14z"/>`
	ctrlInnerSVG                 = `<path fill="currentColor" d="M0 0v16h16V0H0zm4.19 12C2.16 12 1 10.54 1 8s1.16-4 3.19-4h.029c.539 0 1.052.114 1.515.32l-.424.901a2.719 2.719 0 0 0-1.08-.22h-.042C2.38 5.001 2 6.631 2 8.001s.38 3 2.19 3c.497-.013.96-.145 1.366-.368l.444.898a3.943 3.943 0 0 1-1.806.47zM9 7H8v3.5a.902.902 0 0 0 1.005.499L8.999 12a1.822 1.822 0 0 1-1.998-1.428L6.999 7h-.51V6h.51V5h1v1h1v1zm2 2v3h-1V6h1v.92a2.447 2.447 0 0 1 2.005-.919l-.004 1a1.88 1.88 0 0 0-2 2.006zm4 3h-1V3h1v9z"/>`
	ctrlAInnerSVG                = `<path fill="currentColor" d="M9 7V6H8V5H7v1h-.5v1H7v3.56a1.823 1.823 0 0 0 2.009 1.439L9 11a.899.899 0 0 1-.998-.495L8 7h1zm5-4h1v9h-1V3zm-1 3l-.085-.001c-.773 0-1.462.358-1.911.917L11 6.001h-1v6h1v-3a1.88 1.88 0 0 1 2.006-2l-.006-1zm-8.81 6C2.16 12 1 10.54 1 8s1.16-4 3.19-4h.029c.539 0 1.052.114 1.515.32l-.424.901a2.719 2.719 0 0 0-1.08-.22h-.042C2.38 5.001 2 6.631 2 8.001s.38 3 2.19 3c.497-.013.96-.145 1.366-.368l.444.898a3.943 3.943 0 0 1-1.806.47z"/>`
	cubeInnerSVG                 = `<path fill="currentColor" d="M8 0L0 2v10l8 4l8-4V2L8 0zm6.4 2.6L8.5 4.8L1.9 2.6L8 1l6.4 1.6zM1 11.4V3.3l7 2.4v9.2l-7-3.5z"/>`
	cubesInnerSVG                = `<path fill="currentColor" d="M12 6V2L8 0L4 2v4L0 8v5l4 2l4-2l4 2l4-2V8zM8.09 1.12L11 2.56l-2.6 1.3l-2.91-1.44zM5 2.78l3 1.5v3.6l-3-1.5v-3.6zm-1 11.1l-3-1.5v-3.6l3 1.5v3.6zm.28-4L1.4 8.42L4 7.12l2.88 1.44zm7.72 4l-3-1.5v-3.6l3 1.5v3.6zm.28-4L9.4 8.42l2.6-1.3l2.88 1.44z"/>`
	curlyBracketsInnerSVG        = `<path fill="currentColor" d="M2.1 3.1c.2 1.3.4 1.6.4 2.9C2.5 6.8 1 7.5 1 7.5v1s1.5.7 1.5 1.5c0 1.3-.2 1.6-.4 2.9c-.3 2.1.8 3.1 1.8 3.1H6v-2s-1.8.2-1.8-1c0-.9.2-.9.4-2.9c.1-.9-.5-1.6-1.1-2.1c.6-.5 1.2-1.1 1.1-2c-.3-2-.4-2-.4-2.9C4.2 1.9 6 2 6 2V0H3.9C2.8 0 1.8 1 2.1 3.1zm11.8 0c-.2 1.3-.4 1.6-.4 2.9c0 .8 1.5 1.5 1.5 1.5v1s-1.5.7-1.5 1.5c0 1.3.2 1.6.4 2.9c.3 2.1-.8 3.1-1.8 3.1H10v-2s1.8.2 1.8-1c0-.9-.2-.9-.4-2.9c-.1-.9.5-1.6 1.1-2.1c-.6-.5-1.2-1.1-1.1-2c.2-2 .4-2 .4-2.9C11.8 1.9 10 2 10 2V0h2.1c1.1 0 2.1 1 1.8 3.1z"/>`
	cursorInnerSVG               = `<path fill="currentColor" d="M4 0v13l3.31-3.47L10 16l1.37-.63L8.65 9H13L4 0z"/>`
	cursorOInnerSVG              = `<path fill="currentColor" d="M5 2.6L10.75 9H8.29l.63 1.41l1.8 4l-.91.34l-1.88-4.3l-.5-1.11l-1 .71L5 11.07V2.6zM4 0v13l3-2.14L9.26 16l2.8-1l-2.23-5H13L4 0z"/>`
	cutleryInnerSVG              = `<path fill="currentColor" d="M13 .8c0-.5-.4-.8-.8-.8H12c-1.7 0-3 1.9-3 4.7v.9c0 1 .5 1.9 1.4 2.4c-.3 1.2-.4 2.5-.4 2.5v4c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-4c0-.4-.1-1.4-.3-2.3c.2-.2.3-.4.3-.7V.8zM7.2 0H7v3.5c0 .3-.2.5-.5.5S6 3.8 6 3.5v-3c0-.3-.2-.5-.5-.5S5 .2 5 .5v3c0 .3-.2.5-.5.5S4 3.8 4 3.5V0h-.2c-.4 0-.8.4-.8.8v3.7c0 1 .6 1.9 1.5 2.3c-.4 1.6-.5 3.7-.5 3.7v4c0 .8.7 1.5 1.5 1.5S7 15.3 7 14.5v-4c0-.5-.1-2.3-.4-3.7C7.4 6.4 8 5.5 8 4.5V.8c0-.4-.4-.8-.8-.8z"/>`
	dashboardInnerSVG            = `<path fill="currentColor" d="M16 10.1C16 5.7 12.4 2 8 2s-8 3.7-8 8.1c0 1.4.3 2.9.9 3.9h4.9c.5.6 1.3 1 2.2 1s1.7-.4 2.2-1h4.9c.6-1 .9-2.5.9-3.9zM14 7v1l-4.1 3.5c0 .1.1.3.1.5c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2c.3 0 .6.1.8.2L13 7h1zm-4-3h1v1h-1V4zM5 4h1v1H5V4zm-3 8H1v-1h1v1zm1-4H2V7h1v1zm12 4h-1v-1h1v1z"/><path fill="currentColor" d="M9 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`
	databaseInnerSVG             = `<path fill="currentColor" d="M14 2.5C14 3.328 11.314 4 8 4s-6-.672-6-1.5S4.686 1 8 1s6 .672 6 1.5z"/><path fill="currentColor" d="M8 5c-3.3 0-6-.7-6-1.5v3C2 7.3 4.7 8 8 8s6-.7 6-1.5v-3C14 4.3 11.3 5 8 5z"/><path fill="currentColor" d="M8 9c-3.3 0-6-.7-6-1.5v3c0 .8 2.7 1.5 6 1.5s6-.7 6-1.5v-3C14 8.3 11.3 9 8 9z"/><path fill="currentColor" d="M8 13c-3.3 0-6-.7-6-1.5v3c0 .8 2.7 1.5 6 1.5s6-.7 6-1.5v-3c0 .8-2.7 1.5-6 1.5z"/>`
	dateInputInnerSVG            = `<path fill="currentColor" d="M14 1v3h-3V1H5v3H2V1H0v15h16V1h-2zm1 14H1V6h14v9z"/><path fill="currentColor" d="M3 0h1v3H3V0zm9 0h1v3h-1V0zM3 8h1v5H3V8z"/>`
	deindentInnerSVG             = `<path fill="currentColor" d="M4 10.5v-6l-4 3zM0 0h16v3H0V0zm6 4h10v3H6V4zm0 4h10v3H6V8zm-6 4h16v3H0v-3z"/>`
	delInnerSVG                  = `<path fill="currentColor" d="M0 0v16h16V0H0zm3 12H1V3h2a4.111 4.111 0 0 1 3.999 4.517c.013.1.02.236.02.374A4.11 4.11 0 0 1 3.005 12zm10-3H9v.012c0 .607.211 1.164.564 1.603c.252.244.601.397.986.397l.074-.002a2.4 2.4 0 0 0 1.518-.631l.708.712a3.4 3.4 0 0 1-2.225.92l-.09.002a2.393 2.393 0 0 1-1.696-.702a3.522 3.522 0 0 1-.84-2.289v-.041c0-.968.344-1.855.915-2.547a2.144 2.144 0 0 1 1.578-.623l.086-.002a2.33 2.33 0 0 1 1.641.672c.47.532.762 1.23.78 1.996v.524zm2 3h-1V3h1v9z"/><path fill="currentColor" d="M3 4H2v7h1c.31 0 3-.12 3-3.5S3.12 4 3 4zm7.49 2.8a1.432 1.432 0 0 0-1.339 1.192L11.93 8a1.734 1.734 0 0 0-.431-.831a1.355 1.355 0 0 0-.934-.371l-.079.002z"/>`
	delAInnerSVG                 = `<path fill="currentColor" d="M14 3h1v9h-1V3zM3 12H1V3h2a4.111 4.111 0 0 1 3.999 4.517c.013.1.02.236.02.374A4.11 4.11 0 0 1 3.005 12zm-1-1h1c.31 0 3-.12 3-3.5S3.12 4 3 4H2v7zm11-2v-.5a3.116 3.116 0 0 0-.783-2.003a2.332 2.332 0 0 0-1.732-.666l-.054-.001c-.594 0-1.132.241-1.521.631A3.978 3.978 0 0 0 8 9.001v.009c0 .881.322 1.686.854 2.306c.43.429 1.03.697 1.692.697l.089-.002a3.398 3.398 0 0 0 2.228-.922l-.712-.708a2.402 2.402 0 0 1-1.515.63l-.076.002c-.385 0-.734-.153-.99-.402A2.54 2.54 0 0 1 9 9.001h4zm-2.5-2.2l.066-.002c.362 0 .691.141.935.372c.209.224.361.505.427.818l-2.778.011a1.433 1.433 0 0 1 1.337-1.2z"/>`
	dentalChairInnerSVG          = `<path fill="currentColor" d="M11.5 8.2c-.3-.1-.6-.2-.8-.2H8V7h3c0-.6-.4-1-1-1H6c0 .6.4 1 1 1v1c-.5 0-1-.2-1.2-.6L4.7 5.6C4.4 5.2 4 5 3.6 5H3v-.7c0-.3-.1-.5-.2-.8l-.3-.7C2.2 2.3 1.6 2 1 2H0l5 7c.4.6 1.1 1 1.8 1H8v1H7v2h-.6c-.9 0-1.8.4-2.4 1H3v1h11v-1h-1c-.6-.6-1.5-1-2.4-1H10v-2H9v-1h1.6c.2 0 .5.1.7.2l1.7.9c.9.5 2 .5 2.9 0h.1l-4.5-2.9z"/>`
	desktopInnerSVG              = `<path fill="currentColor" d="M16 0H0v13h6v2H4v1h8v-1h-2v-2h6V0zM9 12H7v-1h2v1zm6-2H1V1h14v9z"/>`
	diamondInnerSVG              = `<path fill="currentColor" d="M0 6h4l3 8.6L0 6zm16 0h-4l-3 8.6L16 6zm-8 9L5 6h6l-3 9zM4 5H0l2-3l2 3zm12 0h-4l2-3l2 3zm-6 0H6l2-3l2 3zM3.34 2H7L5 5L3.34 2zM9 2h4l-2 3l-2-3z"/>`
	diamondOInnerSVG             = `<path fill="currentColor" d="M13 2H3L0 5.5L8 15l8-9.5zM4.64 5H1.75l1.52-1.78zm1.78 0L8 3.16L9.58 5H6.42zM10 6l-2 6.68L6 6h4zM5.26 6l1.89 6.44L1.73 6h3.53zm5.49 0h3.53l-5.43 6.44zm.62-1l1.37-1.78L14.25 5h-2.9zM12 3l-1.44 1.81L9.1 3H12zM5.43 4.83L4 3h2.9z"/>`
	diplomaInnerSVG              = `<path fill="currentColor" d="M14 10.58a.371.371 0 0 0-.001-.332l-.479-.698c-.009-.013-.014-.028-.014-.045s.005-.032.014-.045l.48-.7a.371.371 0 0 0-.001-.332a.4.4 0 0 0-.236-.237l-.823-.301a.091.091 0 0 1-.06-.069V6.98a.386.386 0 0 0-.169-.299a.407.407 0 0 0-.231-.071h-.17l-.85.22a.095.095 0 0 1-.071 0l-.549-.65a.428.428 0 0 0-.63 0l-.55.65a.095.095 0 0 1-.071 0h.001l-.9-.23h-.108a.417.417 0 0 0-.234.071a.388.388 0 0 0-.168.298v.841a.092.092 0 0 1-.059.07l-.821.3a.403.403 0 0 0-.338.395c0 .06.014.117.039.167l.479.698c.009.013.014.028.014.045s-.005.032-.014.045l-.48.7a.371.371 0 0 0 .001.332a.4.4 0 0 0 .236.237l.823.301a.091.091 0 0 1 .06.069v.841a.386.386 0 0 0 .169.299a.417.417 0 0 0 .234.071h.168l.31-.07V16l1.53-2l1.47 2v-3.69l.31.08h.118a.417.417 0 0 0 .234-.071a.388.388 0 0 0 .168-.298v-.841a.092.092 0 0 1 .059-.07l.821-.3a.405.405 0 0 0 .289-.227z"/><path fill="currentColor" d="M0 1v12h8l-.11-.05a1.131 1.131 0 0 1-.49-.867V12H1V2h14v10h-1.43v.08a1.134 1.134 0 0 1-.486.868L13 13h3V1H0z"/><path fill="currentColor" d="M7.43 6.91a1.13 1.13 0 0 1 .486-.908A.184.184 0 0 1 8.001 6H3v1h4.43v-.09zM6.42 8H3v1h3.36a.986.986 0 0 1-.047-.837A.292.292 0 0 1 6.42 8zM3 4h10v1H3V4z"/>`
	diplomaScrollInnerSVG        = `<path fill="currentColor" d="M12.61 8.41a5.46 5.46 0 0 1-1.454-.424A60.006 60.006 0 0 1 15.61 4.43l.1-.07l.06-.11a2.924 2.924 0 0 0-.765-3.496a2.916 2.916 0 0 0-3.459-.376l-.126.133A54.96 54.96 0 0 1 6.512 6.41a50.219 50.219 0 0 1-6.018 4.592L.1 11.25v.23a4.53 4.53 0 0 0 .7 3.655A2.83 2.83 0 0 0 3.007 16a1.997 1.997 0 0 0 1.778-.902C5.03 14.79 6.85 12.49 8.79 10.39c.268.464.476 1.003.594 1.575a6.29 6.29 0 0 1-.399 3.078L10.67 13L12 14a12.122 12.122 0 0 0-.584-3.336a5.341 5.341 0 0 0-.915-1.214c.406.346.871.643 1.372.874c.94.338 1.989.572 3.076.672L14 9.73L16 8a8.982 8.982 0 0 1-2.777.431c-.216 0-.43-.007-.642-.022zm-.45-7.23a1.89 1.89 0 0 1 .842-.194c.506 0 .966.196 1.309.516a1.926 1.926 0 0 1 .595 2.192c-.486.307-2.346 1.717-4.146 3.307a2.003 2.003 0 0 0-.668-1.298a1.596 1.596 0 0 0-1.24-.372A58.169 58.169 0 0 0 12.16 1.18zM2.7 11.81a.458.458 0 0 1 .262-.082l.04.002h.068c.179.052.334.142.461.261l-.871.719a1.28 1.28 0 0 1-.119-.716a.334.334 0 0 1 .158-.183zM4 14.5a1 1 0 0 1-1.005.499a1.852 1.852 0 0 1-1.485-.54a3.432 3.432 0 0 1-.583-1.922c0-.251.027-.495.077-.73l.706-.457c-.094.14-.164.304-.199.481a2.985 2.985 0 0 0 .535 1.958l.354.44l1.7-1.4a2.396 2.396 0 0 1-.106 1.685zm.86-2.45a2.825 2.825 0 0 0-1.54-1.274c-.071-.012-.13-.016-.19-.016s-.119.004-.177.01a1.56 1.56 0 0 0-.35 0a44.978 44.978 0 0 0 3.988-3.052c.398.071.812.25 1.131.533c.297.313.48.739.48 1.209l-.002.094a54.377 54.377 0 0 0-3 3.506a2.75 2.75 0 0 0-.357-1.023z"/>`
	discInnerSVG                 = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm7 8c0 1.1-.2 2.1-.7 3l-2.7-1.2c.2-.6.4-1.2.4-1.8c0-2.2-1.8-4-4-4c-.5 0-.9.1-1.4.3L5.4 1.5c.6-.2 1.2-.4 1.8-.5l.3 3H8V1c3.9 0 7 3.1 7 7zM8 5c1.7 0 3 1.3 3 3s-1.3 3-3 3s-3-1.3-3-3s1.3-3 3-3zM1 8c0-1.1.2-2.1.7-3l2.7 1.2C4.2 6.8 4 7.4 4 8c0 2.2 1.8 4 4 4c.5 0 .9-.1 1.4-.3l1.2 2.8c-.6.2-1.2.4-1.8.5l-.3-3H8v3c-3.9 0-7-3.1-7-7z"/><path fill="currentColor" d="M10 8a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8z"/>`
	doctorInnerSVG               = `<path fill="currentColor" d="M14 11.3c-1-1.9-2-1.6-3.1-1.7c.1.3.1.6.1 1c1.6.4 2 2.3 2 3.4v1h-2v-1h1s0-2.5-1.5-2.5S9 13.9 9 14h1v1H8v-1c0-1.1.4-3.1 2-3.4c0-.6-.1-1.1-.2-1.3c-.2-.1-.4-.3-.4-.6c0-.6.8-.4 1.4-1.5c0 0 .9-2.3.6-4.3h-1c0-.2.1-.3.1-.5s0-.3-.1-.5h.8C10.9.9 9.9 0 8 0C6.1 0 5.1.9 4.7 2h.8c0 .2-.1.3-.1.5s0 .3.1.5h-1c-.2 2 .6 4.3.6 4.3c.6 1 1.4.8 1.4 1.5c0 .5-.5.7-1.1.8c-.2.2-.4.6-.4 1.4v1.2c.6.2 1 .8 1 1.4c0 .7-.7 1.4-1.5 1.4S3 14.3 3 13.5c0-.7.4-1.2 1-1.4v-1.2c0-.5.1-.9.2-1.3c-.7.1-1.5.4-2.2 1.7c-.6 1.1-.9 4.7-.9 4.7h13.7c.1 0-.2-3.6-.8-4.7zM6.5 2.5C6.5 1.7 7.2 1 8 1s1.5.7 1.5 1.5S8.8 4 8 4s-1.5-.7-1.5-1.5z"/><path fill="currentColor" d="M5 13.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0z"/>`
	doctorBriefcaseInnerSVG      = `<path fill="currentColor" d="m16 12l-1.4-6.7c-.2-.7-.9-1.3-1.7-1.3H11V2.8c0-1-.8-1.8-1.8-1.8H6.8C5.8 1 5 1.8 5 2.8V4H3.1c-.8 0-1.5.6-1.7 1.3L0 12c-.2 1 .6 2 1.7 2h12.5c1.2 0 2-1 1.8-2zM6 2.8c0-.4.4-.8.8-.8h2.4c.4 0 .8.4.8.8V4H6V2.8zm5 7.2H9v2H7v-2H5V8h2V6h2v2h2v2z"/>`
	dollarInnerSVG               = `<path fill="currentColor" d="M8.2 6.8c-.1 0-.1-.1-.2-.1V3.6c1.2.1 2.2.6 2.2.6l.9-1.8c-.1 0-1.5-.8-3.1-.8V0H7v1.6c-.8.2-1.4.5-2 .9c-.6.6-1 1.4-1 2.3c0 .7.2 2.3 3 3.6v3.9c-.9-.2-2-.7-2.4-.9l-1 1.7c.2.1 1.8 1 3.4 1.2V16h1v-1.7c2.3-.3 3.6-2.1 3.6-3.8c0-1.5-1-2.7-3.4-3.7zM7 6.2c-.8-.5-1-1-1-1.3c0-.4.1-.7.4-.9l.6-.3v2.5zm1 6.1V8.9c1.1.5 1.6 1.1 1.6 1.6c0 .6-.3 1.6-1.6 1.8z"/>`
	dotCircleInnerSVG            = `<path fill="currentColor" d="M8 4C5.8 4 4 5.8 4 8s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4z"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/>`
	downloadInnerSVG             = `<path fill="currentColor" d="M16 10h-5.5L8 12.5L5.5 10H0v6h16v-6zM4 14H2v-2h2v2z"/><path fill="currentColor" d="M10 6V0H6v6H3l5 5l5-5z"/>`
	downloadAltInnerSVG          = `<path fill="currentColor" d="M0 14h16v2H0v-2zm8-1l5-5h-3V0H6v8H3z"/>`
	dropInnerSVG                 = `<path fill="currentColor" d="M8 0S3 8.2 3 11s2.2 5 5 5s5-2.2 5-5S8 0 8 0zm.9 14.9l-.2-1c1.4-.3 2.4-1.7 2.4-3.2c0-.3-.1-1.1-.8-2.6l.9-.4c.6 1.4.8 2.4.8 3c0 2-1.3 3.8-3.1 4.2z"/>`
	editInnerSVG                 = `<path fill="currentColor" d="M16 4s0-1-1-2s-1.9-1-1.9-1L12 2.1V0H0v16h12V8l4-4zm-9.7 7.4l-.6-.6l.3-1.1l1.5 1.5l-1.2.2zm.9-1.9l-.6-.6l5.2-5.2c.2.1.4.3.6.5zm6.9-7l-.9 1c-.2-.2-.4-.3-.6-.5l.9-.9c.1.1.3.2.6.4zM11 15H1V1h10v2.1L5.1 9L4 13.1L8.1 12L11 9v6z"/>`
	ejectInnerSVG                = `<path fill="currentColor" d="M1 11h14L8 1zm0 1h14v3H1v-3z"/>`
	elasticInnerSVG              = `<path fill="currentColor" d="M4.7 16c-1.7 0-3.1-.8-4-2.1c-1.1-1.7-.9-4 .4-5.8C2 6.8 3.2 6 4.7 5.7c1.2-.3 2.2-1.1 2.5-2.2c.2-.8.7-1.5 1.3-2C9.4.5 10.7 0 12 0c1.1 0 2.2.4 2.9 1.2c1.5 1.6 1.5 4.2-.1 6c-.5.6-1.2 1.1-2 1.4c-1.2.5-2.2 1.6-2.6 3c-.3 1-.8 1.9-1.5 2.6c-1.1 1.2-2.6 1.8-4 1.8zM12 1c-1 0-2 .4-2.8 1.2c-.5.5-.8 1-1 1.6c-.5 1.5-1.8 2.5-3.3 2.9c-1.2.2-2.2.9-3 2C.8 10.2.7 12 1.6 13.4c.6 1 1.8 1.6 3.1 1.6c1.2 0 2.4-.5 3.3-1.4c.6-.6 1.1-1.4 1.3-2.2c.4-1.7 1.6-3 3.2-3.6c.6-.2 1.2-.7 1.6-1.2c1.2-1.4 1.3-3.5.1-4.7c-.6-.6-1.4-.9-2.2-.9z"/>`
	ellipsisCircleInnerSVG       = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zM6 9H4V7h2v2zm3 0H7V7h2v2zm3 0h-2V7h2v2z"/>`
	ellipsisCircleOInnerSVG      = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M4 7h2v2H4V7zm3 0h2v2H7V7zm3 0h2v2h-2V7z"/>`
	ellipsisDotsHInnerSVG        = `<path fill="currentColor" d="M4 8a2 2 0 1 1-3.999.001A2 2 0 0 1 4 8zm6 0a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8zm6 0a2 2 0 1 1-3.999.001A2 2 0 0 1 16 8z"/>`
	ellipsisDotsVInnerSVG        = `<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2zm0 6a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8zm0 6a2 2 0 1 1-3.999.001A2 2 0 0 1 10 14z"/>`
	ellipsisHInnerSVG            = `<path fill="currentColor" d="M0 6h4v4H0V6zm6 0h4v4H6V6zm6 0h4v4h-4V6z"/>`
	ellipsisVInnerSVG            = `<path fill="currentColor" d="M6 0h4v4H6V0zm0 6h4v4H6V6zm0 6h4v4H6v-4z"/>`
	enterInnerSVG                = `<path fill="currentColor" d="M4 0v6H1v10h14V0H4zm8 11H7v2l-3-2.5L7 8v2h4V7h1v4z"/>`
	enterArrowInnerSVG           = `<path fill="currentColor" d="m0 9l7 4v-3h9V3l-3 2v2H7V4L0 9z"/>`
	envelopeInnerSVG             = `<path fill="currentColor" d="M0 3h16v2.4l-8 4l-8-4z"/><path fill="currentColor" d="m0 14l5.5-4.8L8 10.6l2.5-1.4L16 14zm4.6-5.2L0 6.5V13zm6.8 0L16 6.5V13z"/>`
	envelopeOInnerSVG            = `<path fill="currentColor" d="M0 3v11h16V3H0zm1 4.1l3.9 2L1 12.5V7.1zm.9 5.9l4-3.5L8 10.6l2.1-1.1l4 3.5H1.9zm13.1-.5L11.1 9L15 7v5.5zm0-6.6L8 9.4L1 5.9V4h14v1.9z"/>`
	envelopeOpenInnerSVG         = `<path fill="currentColor" d="M14 3.7v3.7l2-1V5zM2 3.8L0 5v1.5l2 1zM11.2 2L8 0L4.8 2zM13 3H3v4.9l3.4 1.7L8 8.4l1.6 1.2L13 7.9zm3 4.6l-5.5 2.7l5.5 4.4zm-8 2L0 16h16zm-2.5.7L0 7.6v7.1z"/>`
	envelopeOpenOInnerSVG        = `<path fill="currentColor" d="M14 3.7V3h-1.5L8 0L3.4 3H2v.7L0 5v11h16V5.1l-2-1.4zM8 1.2L10.7 3H5.2L8 1.2zM3 4h10v3.7L9.5 9.4L8 8.1L6.5 9.5L3 7.8V4zM1 5.5l1-.7v2.4l-1-.4V5.5zm0 2.4l4.6 2.3l-4.6 4V7.9zm.9 7.1L8 9.7l6.1 5.3H1.9zm13.1-.8l-4.7-4.1L15 7.8v6.4zm0-7.5l-1 .5V4.9l1 .7v1.1z"/>`
	envelopesInnerSVG            = `<path fill="currentColor" d="M16 14H2v-1h13V4h1v10z"/><path fill="currentColor" d="M14 10.77V5.29L9.32 7.47l4.68 3.3zM8.28 7.96L7 8.55l-1.28-.59L0 11.99V12l14-.01l-5.72-4.03zM7 7.45l7-3.27V2H0v2.18l7 3.27zm-2.32.02L0 5.29v5.48l4.68-3.3z"/>`
	envelopesOInnerSVG           = `<path fill="currentColor" d="M14 2H0v10h14V2zM5.71 8L7 8.55L8.29 8L13 11H1zM1 9.83v-4l3.64 1.63zm8.36-2.37L13 5.78v4zM13 3v1.68L7 7.45L1 4.68V3h12z"/><path fill="currentColor" d="M15 4v9H2v1h14V4h-1z"/>`
	eraserInnerSVG               = `<path fill="currentColor" d="m8.1 14l6.4-7.2c.6-.7.6-1.8-.1-2.5l-2.7-2.7c-.3-.4-.8-.6-1.3-.6H8.6c-.5 0-1 .2-1.4.6L.5 9.2c-.6.7-.6 1.9.1 2.5l2.7 2.7c.3.4.8.6 1.3.6H16v-1H8.1zm-1.3-.1s0-.1 0 0l-2.7-2.7c-.4-.4-.4-.9 0-1.3L7.5 6h-1l-3 3.3c-.6.7-.6 1.7.1 2.4L5.9 14H4.6c-.2 0-.4-.1-.6-.2L1.2 11c-.3-.3-.3-.8 0-1.1L4.7 6h1.8L10 2h1L7.5 6l3.1 3.7l-3.5 4c-.1.1-.2.1-.3.2z"/>`
	escInnerSVG                  = `<path fill="currentColor" d="M0 0v16h16V0H0zm5 4H2v3h3v1H2v3h3v1H1V3h4v1zm5 6.54c-.067.511-.364.94-.782 1.186a2.426 2.426 0 0 1-1.129.276L7.996 12a6.254 6.254 0 0 1-2.038-.425l.403-.915c.435.202.945.319 1.482.319c.326 0 .643-.043.943-.125a.662.662 0 0 0 .215-.484c.07-.43-.22-.62-1.17-1c-.83-.29-2.04-.76-1.83-2.08c.072-.594.46-1.082.989-1.296a3.252 3.252 0 0 1 2.649.552L9.07 7.3a2.32 2.32 0 0 0-1.663-.368a.617.617 0 0 0-.387.547c-.08.401.14.581 1.15 1.001c.83.3 2.02.75 1.83 2.06zm2.67.18c.345.176.752.279 1.183.279c.292 0 .573-.047.835-.134l.311.945c-.383.121-.823.19-1.279.19h-.001l-.089.001c-.583 0-1.124-.18-1.57-.488a2.998 2.998 0 0 1-1.061-2.524a2.866 2.866 0 0 1 1.044-2.516a3.502 3.502 0 0 1 1.72-.446c.443 0 .868.081 1.259.23l-.374.922a2.548 2.548 0 0 0-2.016.066a2.013 2.013 0 0 0-.633 1.764a2.056 2.056 0 0 0 .637 1.708z"/>`
	escAInnerSVG                 = `<path fill="currentColor" d="M8 12a6.268 6.268 0 0 1-2.043-.425l.403-.915c.435.202.945.319 1.482.319c.326 0 .643-.043.943-.125A.662.662 0 0 0 9 10.37c.07-.43-.22-.62-1.17-1C7 9.08 5.79 8.61 6 7.29c.072-.594.46-1.082.989-1.296a3.252 3.252 0 0 1 2.649.552l-.569.754a2.32 2.32 0 0 0-1.663-.368a.617.617 0 0 0-.387.547c-.08.401.14.581 1.15 1.001c.85.33 2 .77 1.8 2.08c-.067.511-.364.94-.782 1.186A2.42 2.42 0 0 1 7.994 12zm5.71 0l-.089.001c-.583 0-1.124-.18-1.57-.488a2.995 2.995 0 0 1-1.05-2.524a2.866 2.866 0 0 1 1.044-2.516a3.502 3.502 0 0 1 1.72-.446c.443 0 .868.081 1.259.23l-.374.922a2.548 2.548 0 0 0-2.016.066a2.013 2.013 0 0 0-.633 1.764a2.052 2.052 0 0 0 .647 1.748c.346.177.754.28 1.185.28c.292 0 .573-.047.835-.134l.331.905c-.383.121-.823.19-1.279.19h-.012zM5 4V3H1v9h4v-1H2V8h3V7H2V4h3z"/>`
	euroInnerSVG                 = `<path fill="currentColor" d="M10.89 3a5.45 5.45 0 0 1 3.127 1.011L14 1.69a7.369 7.369 0 0 0-3.129-.686a7.482 7.482 0 0 0-7.323 5.947L2 7v1h1.41v.5a3.848 3.848 0 0 0 0 .512L1.999 9v1h1.54c.882 3.353 3.805 5.818 7.331 5.999a7.42 7.42 0 0 0 3.175-.708L14 13a5.426 5.426 0 0 1-3.108 1a5.909 5.909 0 0 1-5.28-3.959L12 10V9H5.41a3.224 3.224 0 0 1 .001-.511L5.41 8H12V7H5.6c.678-2.325 2.788-3.996 5.29-4z"/>`
	exchangeInnerSVG             = `<path fill="currentColor" d="M16 5v2H3v2L0 6l3-3v2zM0 12v-2h13V8l3 3l-3 3v-2z"/>`
	exclamationInnerSVG          = `<path fill="currentColor" d="M6 0h4v4l-1 7H7L6 4zm4 14a2 2 0 1 1-3.999.001A2 2 0 0 1 10 14z"/>`
	exclamationCircleInnerSVG    = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm1 13H7v-2h2v2zm0-3H7V3h2v7z"/>`
	exclamationCircleOInnerSVG   = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M7 3h2v7H7V3zm0 8h2v2H7v-2z"/>`
	exitInnerSVG                 = `<path fill="currentColor" d="M14 6h-1.7c-.2 0-.4-.1-.6-.2l-1.3-1.3c-.2-.3-.6-.5-1.1-.5H9c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2c0 .7.4 1.4 1 1.7l-.2.3h-2c-1.1 0-2.3.5-3 1.5l-.6.8c-.4.4-.2 1 .2 1.3c.4.2.9.1 1.2-.3l.5-.7c.3-.4.7-.6 1.2-.6h.8l-.7 1.6c-.3.6-.4 1.2-.4 1.9v2c0 .3-.2.5-.5.5H2c-.6 0-1 .4-1 1s.4 1 1 1h3.5c.8 0 1.5-.7 1.5-1.5V10l3.8 4.5c.6.9 1.7 1.5 2.8 1.5h.9L9.1 9.3c-.3-.4-.2-.8 0-1.3l.6-1.5l.7.8c.4.4 1 .7 1.6.7h2c.6 0 1-.4 1-1s-.4-1-1-1z"/>`
	exitOInnerSVG                = `<path fill="currentColor" d="M10 0c1.1 0 2 .9 2 2c0 .9-.6 1.7-1.5 1.9V4c.4 0 .7.2 1 .5l1.3 1.3c.1.1.3.2.5.2H15V0h-5z"/><path fill="currentColor" d="M11.8 14.5L8 10v2.5c0 .8-.7 1.5-1.5 1.5H3c-.6 0-1-.4-1-1s.4-1 1-1h2.5c.3 0 .5-.2.5-.5v-2c0-.7.1-1.3.4-2L7.1 6h-.8c-.5 0-.9.2-1.2.6l-.5.7c-.2.4-.7.5-1.2.3c-.4-.3-.6-.9-.2-1.3l.6-.8C4.5 4.5 5.7 4 6.9 4h2l.1-.3c-.6-.3-1-1-1-1.7c0-1.1.9-2 2-2H3v4.9l-.6.8c-.3.4-.5.9-.4 1.5c.1.5.4 1 .9 1.3V11c-1.1 0-2 .9-2 2s.9 2 2 2v1h11.6c-1.1 0-2.1-.6-2.7-1.5z"/><path fill="currentColor" d="m11.4 7.3l-.7-.8l-.6 1.5c-.2.5-.3.9 0 1.3l4.9 6.1V8h-2.1c-.6 0-1.1-.2-1.5-.7z"/>`
	expandInnerSVG               = `<path fill="currentColor" d="M15 1h-4l1.3 1.3l-4.5 4.5l1.4 1.4l4.5-4.5L15 5zM6.8 7.8l-4.5 4.5L1 11v4h4l-1.3-1.3l4.5-4.5z"/>`
	expandFullInnerSVG           = `<path fill="currentColor" d="m5.3 6.7l1.4-1.4l-3-3L5 1H1v4l1.3-1.3zm1.4 4L5.3 9.3l-3 3L1 11v4h4l-1.3-1.3zm4-1.4l-1.4 1.4l3 3L11 15h4v-4l-1.3 1.3zM11 1l1.3 1.3l-3 3l1.4 1.4l3-3L15 5V1z"/>`
	expandSquareInnerSVG         = `<path fill="currentColor" d="M11 2H2v9l1-1V3h7zM5 14h9V5l-1 1v7H6z"/><path fill="currentColor" d="M16 0h-5l1.8 1.8l-4.5 4.5l1.4 1.4l4.5-4.5L16 5zM7.7 9.7L6.3 8.3l-4.5 4.5L0 11v5h5l-1.8-1.8z"/>`
	externalBrowserInnerSVG      = `<path fill="currentColor" d="M11 10L8.1 6.8L4.8 10H7v1.8c0 1.7-.9 4.2-4 4.2c4.8 0 6-1.4 6-4.3V10h2z"/><path fill="currentColor" d="M0 0v13h6v-1H1V3h14v9h-5v1h6V0H0zm2 2H1V1h1v1zm11 0H3V1h10v1z"/>`
	externalLinkInnerSVG         = `<path fill="currentColor" d="M14 16V5l-1 1v9H1V3h9l1-1H0v14z"/><path fill="currentColor" d="M16 0h-5l1.8 1.8L6 8.6L7.4 10l6.8-6.8L16 5z"/>`
	eyeInnerSVG                  = `<path fill="currentColor" d="M8 3.9C1.3 3.9 0 9 0 9s2.2 4.1 7.9 4.1s8.1-4 8.1-4s-1.3-5.2-8-5.2zM5.3 5.4c.5-.3 1.3-.3 1.3-.3s-.5.9-.5 1.6c0 .7.2 1.1.2 1.1L5.2 8s-.3-.5-.3-1.2c0-.8.4-1.4.4-1.4zm2.6 6.7c-4.1 0-6.2-2.3-6.8-3.2c.3-.7 1.1-2.2 3.1-3.2c-.1.4-.2.8-.2 1.3c0 2.2 1.8 4 4 4s4-1.8 4-4c0-.5-.1-.9-.2-1.3c2 .9 2.8 2.5 3.1 3.2c-.7.9-2.8 3.2-7 3.2z"/>`
	eyeSlashInnerSVG             = `<path fill="currentColor" d="m12.9 5.2l-.8.8c1.7.9 2.5 2.3 2.8 3c-.7.9-2.8 3.1-7 3.1c-.7 0-1.2-.1-1.8-.2l-.8.8c.8.3 1.7.4 2.6.4c5.7 0 8.1-4 8.1-4s-.6-2.4-3.1-3.9z"/><path fill="currentColor" d="M12 7.1c0-.3 0-.6-.1-.8L7.1 11c.3 0 .6.1.9.1c2.2 0 4-1.8 4-4zM15.3 0l-4.4 4.4C10.1 4.2 9.1 4 8 4C1.3 4 0 9.1 0 9.1s1 1.8 3.3 3L0 15.3v.7h.7L16 .7V0h-.7zM4 11.3C2.4 10.6 1.5 9.5 1.1 9c.3-.7 1.1-2.2 3.1-3.2c-.1.4-.2.8-.2 1.3c0 1.1.5 2.2 1.3 2.9L4 11.3zm2.2-3.4l-1 .2s-.3-.5-.3-1.2c0-.8.4-1.5.4-1.5c.5-.3 1.3-.3 1.3-.3s-.5.9-.5 1.7c-.1.7.1 1.1.1 1.1z"/>`
	eyedropperInnerSVG           = `<path fill="currentColor" d="M15 1c-1.8-1.8-3.7-.7-4.6.1c-.4.4-.7.9-.7 1.5c0 1.1-1.1 1.8-2.1 1.5L7.5 4l-.7.8l.7.7l-6 6l-.8 2.3l-.7.7L1.5 16l.8-.8l2.3-.8l6-6l.7.7l.7-.6l-.1-.2c-.3-1 .4-2.1 1.5-2.1c.6 0 1.1-.2 1.4-.6c.9-.9 2-2.8.2-4.6zM3.9 13.6l-2 .7l-.2.1l.1-.2l.7-2l5.8-5.8l1.5 1.5l-5.9 5.7z"/>`
	facebookInnerSVG             = `<path fill="currentColor" d="M7.2 16V8.5h-2V5.8h2V3.5C7.2 1.7 8.4 0 11.1 0c1.1 0 1.9.1 1.9.1l-.1 2.5h-1.7c-1 0-1.1.4-1.1 1.2v2H13l-.1 2.7h-2.8V16H7.2z"/>`
	facebookSquareInnerSVG       = `<path fill="currentColor" d="M0 0v16h16V0H0zm12.9 8.4h-2.1V14H8.7V8.4H7.2v-2h1.5V4.7c0-1.5.9-2.7 2.9-2.7c.8 0 1.4.1 1.4.1V4h-1.3c-.7 0-.8.3-.8.9v1.5H13l-.1 2z"/>`
	factoryInnerSVG              = `<path fill="currentColor" d="M4.4 1.3c-.6.3-.8 1.1-.4 1.5c.5-.9 1.3-.6 2.5.4c.8.7 1.9.1 1.9.1s.2 1.2 1.7 1.4c1.7.2 2.3-.8 2.3-.8s.4 1 1.9.4c1.1-.4.7-1.1.7-1.1s1 0 1-.7c0-.9-1.1-.8-1.1-.8s.2-1-.9-1.1c-1-.1-1.3.5-1.3.5S12.4 0 10.9 0C9.5 0 9 1.3 9 1.3S8.6.7 7.4.7c-.9 0-1.3.7-1.3.7S5 .9 4.4 1.3z"/><path fill="currentColor" d="M12 12.1V10l-4 2.1V10H5.6L5 3H3l-.6 7H0v6h16v-6l-4 2.1zM6 14H2v-2h4v2z"/>`
	familyInnerSVG               = `<path fill="currentColor" d="M9.5 7.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 9.5 7.5zM14.27 4h-2.54A1.73 1.73 0 0 0 10 5.73V9a1 1 0 0 0 1 1v6h4v-6a1 1 0 0 0 1-1V5.73A1.73 1.73 0 0 0 14.27 4z"/><path fill="currentColor" d="M15 2a2 2 0 1 1-3.999.001A2 2 0 0 1 15 2zM4.27 5H1.73a1.73 1.73 0 1 0 .001 3.461A1.73 1.73 0 0 0 1.73 5A1.73 1.73 0 0 0 0 6.73V9a1 1 0 0 0 1 1l-1 3h1v3h4v-3h1l-1-3a1 1 0 0 0 1-1V6.73A1.73 1.73 0 0 0 4.27 5z"/><path fill="currentColor" d="M5 3a2 2 0 1 1-3.999.001A2 2 0 0 1 5 3zm2 10v3h2v-3a1 1 0 0 0 1-1v-1.54A1.46 1.46 0 0 0 8.54 9H7.46A1.46 1.46 0 0 0 6 10.46V12a1 1 0 0 0 1 1z"/>`
	fastBackwardInnerSVG         = `<path fill="currentColor" d="M16 15V1L9 8zm-7 0V1L2 8zM0 1h2v14H0V1z"/>`
	fastForwardInnerSVG          = `<path fill="currentColor" d="M0 1v14l7-7zm7 0v14l7-7zm7 0h2v14h-2V1z"/>`
	femaleInnerSVG               = `<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2z"/><path fill="currentColor" d="M10 8V6.5l1.8 1.8c.3.3.7.3 1 0s.3-.8 0-1l-2.6-2.6c-.4-.5-1-.7-1.7-.7h-1c-.7 0-1.3.2-1.7.7L3.2 7.3c-.3.3-.3.8 0 1c.3.3.7.3 1 0L6 6.5V8l-4 5h4v3h4v-3h4l-4-5z"/>`
	fileInnerSVG                 = `<path fill="currentColor" d="M9 5h5v11H2V0h7v5zm1-1V0l4 4h-4z"/>`
	fileAddInnerSVG              = `<path fill="currentColor" d="M12 15H2V1h6v4h4v1h1V4L9 0H1v16h12v-2h-1v1zM9 1l3 3H9V1z"/><path fill="currentColor" d="M13 7h-2v2H9v2h2v2h2v-2h2V9h-2V7z"/>`
	fileCodeInnerSVG             = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M6.2 13h-.7l-2-2.5l2-2.5h.7l-2 2.5zm3.6 0h.7l2-2.5l-2-2.5h-.7l2 2.5zm-3.1 1h.6l2.1-7h-.8z"/>`
	fileFontInnerSVG             = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M5 7v2h2v5h2V9h2V7z"/>`
	fileMovieInnerSVG            = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M10 10V8H4v5h6v-2l2 2V8z"/>`
	fileOInnerSVG                = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/>`
	filePictureInnerSVG          = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M4 11.5V14h8v-1.7s.1-1.3-1.3-1.5c-1.3-.2-1.5.4-2.5.5c-.8 0-.6-1.3-2.2-1.3c-1.2 0-2 1.5-2 1.5zm8-3a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 12 8.5z"/>`
	filePresentationInnerSVG     = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zm3 15H3V1h6v4h4v10zM10 4V1l3 3h-3z"/><path fill="currentColor" d="M9 6H7v1H4v6h2v1h1v-1h2v1h1v-1h2V7H9V6zm2 2v4H5V8h6z"/><path fill="currentColor" d="M7 9v2l2-1z"/>`
	fileProcessInnerSVG          = `<path fill="currentColor" d="M12 0H5v6h.7l.2.7l.1.1V1h5v4h4v9H9l.3.5l-.5.5H16V4l-4-4zm0 4V1l3 3h-3zm-6.5 7.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/><path fill="currentColor" d="M7.9 12.4L9 12v-1l-1.1-.4c-.1-.3-.2-.6-.4-.9l.5-1l-.7-.7l-1 .5c-.3-.2-.6-.3-.9-.4L5 7H4l-.4 1.1c-.3.1-.6.2-.9.4l-1-.5l-.7.7l.5 1.1c-.2.3-.3.6-.4.9L0 11v1l1.1.4c.1.3.2.6.4.9l-.5 1l.7.7l1.1-.5c.3.2.6.3.9.4L4 16h1l.4-1.1c.3-.1.6-.2.9-.4l1 .5l.7-.7l-.5-1.1c.2-.2.3-.5.4-.8zm-3.4 1.1c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2z"/>`
	fileRefreshInnerSVG          = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zm3 15H3V1h6v4h4v10zM10 4V1l3 3h-3z"/><path fill="currentColor" d="M4.7 7.7L4 7v3h3L5.8 8.8C6.2 8 7.1 7.5 8 7.5c1.4 0 2.5 1.1 2.5 2.5H12c0-2.2-1.8-4-4-4c-1.3 0-2.5.7-3.3 1.7zm5.1 4.1c-.5.5-1.1.8-1.8.7c-1 0-1.9-.6-2.3-1.5H4.1c.4 1.7 2 3 3.8 3c1.1 0 2.1-.5 2.8-1.2L12 14v-3H9l.8.8z"/>`
	fileRemoveInnerSVG           = `<path fill="currentColor" d="M12 15H2V1h6v4h4v2.59l1-1V4L9 0H1v16h12v-2.59l-1-1V15zM9 1l3 3H9V1z"/><path fill="currentColor" d="m15 8l-1-1l-2 2l-2-2l-1 1l2 2l-2 2l1 1l2-2l2 2l1-1l-2-2l2-2z"/>`
	fileSearchInnerSVG           = `<path fill="currentColor" d="M12 13.47V15H2V1h6v4h4v.56c.386.229.716.504.996.825L13 4L9 0H1v16h12v-1.53zM9 1l3 3H9V1z"/><path fill="currentColor" d="m14.78 12.72l-1.92-1.92a.727.727 0 0 0-.325-.179a3.014 3.014 0 0 0 .468-1.618a3 3 0 1 0-1.371 2.52c.02.136.083.248.169.337l1.92 1.92a.75.75 0 0 0 1.059-1.061zM10 11a2 2 0 1 1-.001-3.999A2 2 0 0 1 10 11z"/>`
	fileSoundInnerSVG            = `<path fill="currentColor" d="M11.4 10.5c0 1.2-.4 2.2-1 3l.4.5c.7-.9 1.2-2.1 1.2-3.5s-.5-2.6-1.2-3.5l-.4.5c.6.8 1 1.9 1 3z"/><path fill="currentColor" d="m9.9 8l-.4.5c.4.5.7 1.2.7 2s-.3 1.5-.7 2l.4.5c.5-.6.8-1.5.8-2.5s-.3-1.8-.8-2.5z"/><path fill="currentColor" d="m9.1 9l-.4.5c.2.3.3.6.3 1s-.1.7-.3 1l.4.5c.3-.4.5-.9.5-1.5S9.4 9.4 9.1 9z"/><path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M6 9H4v3h2l2 2V7z"/>`
	fileStartInnerSVG            = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zm3 15H3V1h6v4h4v10zM10 4V1l3 3h-3z"/><path fill="currentColor" d="M5 6v6l6-3z"/>`
	fileTableInnerSVG            = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M4 7v6h8V7H4zm2 5H5v-1h1v1zm0-2H5V9h1v1zm3 2H7v-1h2v1zm0-2H7V9h2v1zm2 2h-1v-1h1v1zm0-2h-1V9h1v1z"/>`
	fileTextInnerSVG             = `<path fill="currentColor" d="M10 0v4h4z"/><path fill="currentColor" d="M9 0H2v16h12V5H9V0zm3 12H4v-1h8v1zm0-2H4V9h8v1zm0-3v1H4V7h8z"/>`
	fileTextOInnerSVG            = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M4 7h8v1H4V7zm0 2h8v1H4V9zm0 2h8v1H4v-1z"/>`
	fileTreeInnerSVG             = `<path fill="currentColor" d="M16 10V6H5v1H3V4h9V0H0v4h2v10h3v2h11v-4H5v1H3V8h2v2z"/>`
	fileTreeSmallInnerSVG        = `<path fill="currentColor" d="M5 12v2h11V9H5v2H3V7h9V2H0v5h2v5z"/>`
	fileTreeSubInnerSVG          = `<path fill="currentColor" d="M8 11v1H7v-2h5V6H4v1H3V5h6V1H0v4h2v3h2v2h2v3h2v2h8v-4z"/>`
	fileZipInnerSVG              = `<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 15H5v-2.8l.7-2.2h2.4l.9 2.2V15zm4 0h-3v-3L9 9H7V8H5v1l-1 3v3H3V1h4v1h2v1H7v1h2v1h4v10zM10 4V1l3 3h-3z"/><path fill="currentColor" d="M5 6h2v1H5V6zm0-4h2v1H5V2zm0 2h2v1H5V4zm2 1h2v1H7V5zm0 2h2v1H7V7zm-1 5h2v2H6v-2z"/>`
	fillInnerSVG                 = `<path fill="currentColor" d="M13 14.5c.468-2.207.985-4.05 1.604-5.846c.411 1.796.928 3.638 1.337 5.521C16 15.328 15.329 16 14.5 16s-1.5-.672-1.5-1.5zM8 1L6.56 2.44l-2-2a1.539 1.539 0 0 0-2.121 0a1.496 1.496 0 0 0 .001 2.119l2 2L0 8.999l7 7l8-8zm0 1.41L13.59 8H2.41l2.75-2.75a.49.49 0 0 0 .669-.672l.721-.718l1.54 1.53a.502.502 0 0 0 .71-.71L7.27 3.15zm-4.85-.56a.5.5 0 0 1 .355-.854c.138 0 .263.055.355.144l2 2l-.71.71z"/>`
	filmInnerSVG                 = `<path fill="currentColor" d="M0 0v16h1v-1h1v1h12v-1h1v1h1V0H0zm2 14H1v-1h1v1zm0-2H1v-1h1v1zm0-2H1V9h1v1zm0-2H1V7h1v1zm0-2H1V5h1v1zm0-2H1V3h1v1zm0-2H1V1h1v1zm11 13H3V9h10v6zm0-8H3V1h10v6zm2 7h-1v-1h1v1zm0-2h-1v-1h1v1zm0-2h-1V9h1v1zm0-2h-1V7h1v1zm0-2h-1V5h1v1zm0-2h-1V3h1v1zm0-2h-1V1h1v1z"/>`
	filterInnerSVG               = `<path fill="currentColor" d="M1 2h14v2L9 9v7l-2-2V9L1 4V2zm0-2h14v1H1V0z"/>`
	fireInnerSVG                 = `<path fill="currentColor" d="M4.9 15.8S1 15.4 1 10.1C1 6 4.1 3.6 4.1 3.6S5.4 5 6.4 5.5C7.4 6.1 7.8 0 7.8 0S15 3.9 15 9.8c0 6.1-4 5.9-4 5.9s1.8-2.4 1.8-5.2c0-3-3.9-6.7-3.9-6.7s-.5 4.4-2.1 5c-1.6-.9-2.5-2.3-2.5-2.3s-3.7 5.8.6 9.3z"/><path fill="currentColor" d="M8.2 16.1c-2-.1-3.7-1.4-3.7-3.2s.7-2.6.7-2.6s.5 1 1.1 1.5s1.8.8 2.4.1c.6-.6.8-2.3.8-2.3s1.4 1.1 1.2 3c-.1 2-.9 3.5-2.5 3.5z"/>`
	flagInnerSVG                 = `<path fill="currentColor" d="M4 2c0-1.1-.9-2-2-2S0 .9 0 2c0 .7.4 1.4 1 1.7V16h2V3.7c.6-.3 1-1 1-1.7zm0 2s1-3 3.6-3c2.7 0 2.3 1 4.4 1c2.4 0 4-1 4-1v8s-.7 2-4 2c-2.2 0-2.3-1-5-1c-2.3 0-3 2-3 2V4z"/>`
	flagCheckeredInnerSVG        = `<path fill="currentColor" d="M2 0C.9 0 0 .9 0 2c0 .7.4 1.4 1 1.7V16h2V3.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2zm10 2c-2.1 0-1.8-1-4.4-1S4 4 4 4v8s.7-2 3-2c2.7 0 2.8 1 5 1c3.3 0 4-2 4-2V1s-1.6 1-4 1zm3 2.5c-.2.2-.8.4-2 .6V2.9c.8-.1 1.5-.2 2-.4v2zM5 7.9V5.3c.4-.6 1.1-1.1 2-1.1V2.1c.2-.1.4-.1.6-.1c1.2 0 1.6.2 2.1.4c.1.1.2.2.3.2v2.2c.5.2 1.1.4 2 .4c.4 0 .7 0 1-.1v2.6c-.3 0-.6.1-1 .1c-1.1 0-1.5-.2-2-.5v2.3C9.3 9.3 8.5 9 7 9V6.8c-.9.2-1.5.6-2 1.1zm8 2V7.7c1.1-.2 1.7-.6 2-.8v1.8c-.2.3-.7 1-2 1.2z"/><path fill="currentColor" d="M10 7.2V4.8s-1.2-.6-3-.6v2.6c1.7-.4 3 .4 3 .4z"/>`
	flagOInnerSVG                = `<path fill="currentColor" d="M4 2c0-1.1-.9-2-2-2S0 .9 0 2c0 .7.4 1.4 1 1.7V16h2V3.7c.6-.3 1-1 1-1.7zm3.6 0c1.2 0 1.6.2 2.1.4c.5.3 1.1.6 2.3.6s2.2-.2 3-.5v6.3c-.2.3-.9 1.2-3 1.2c-.9 0-1.3-.2-1.9-.4C9.4 9.3 8.6 9 7 9c-.8 0-1.5.2-2 .5V4.2C5.2 3.7 6 2 7.6 2zM16 1s-1.6 1-4 1c-2.1 0-1.8-1-4.4-1S4 4 4 4v8s.7-2 3-2c2.7 0 2.8 1 5 1c3.3 0 4-2 4-2V1z"/>`
	flashInnerSVG                = `<path fill="currentColor" d="m16 8l-2.2-1.6L14.9 4l-2.7-.2l-.2-2.7l-2.4 1.1L8 0L6.4 2.2L4 1.1l-.2 2.7l-2.7.2l1.1 2.4L0 8l2.2 1.6L1.1 12l2.7.2l.2 2.7l2.4-1.1L8 16l1.6-2.2l2.4 1.1l.2-2.7l2.7-.2l-1.1-2.4L16 8z"/>`
	flaskInnerSVG                = `<path fill="currentColor" d="M2 16h12l-4-8V1h1V0H5v1h1v7l-4 8zM9 1v7.2l1.9 3.8H5.1L7 8.2V1h2z"/>`
	flightLandingInnerSVG        = `<path fill="currentColor" d="M13.64 7c-.71-.2-1.89-.43-3.23-.67L6.59 2.09a2.268 2.268 0 0 0-.746-.544L4.65 1c-.09 0-.15 0-.1.11S6 4 6.84 5.7c-1.84-.29-3.5-.53-4.23-.63a.917.917 0 0 1-.608-.406L1.28 3.59a.925.925 0 0 0-.474-.358L0 3l.61 3.26c.067.34.318.609.644.699C2.58 7.34 6.07 8.3 8.78 8.88c6 1.28 6.8 1.28 7.12.91S15.23 7.41 13.64 7zM0 13h16v1H0v-1z"/>`
	flightTakeoffInnerSVG        = `<path fill="currentColor" d="M12.57 2.26c-.65.29-1.66.85-2.8 1.5L4.31 3a2.172 2.172 0 0 0-.916.064L2.209 3.4c-.1 0-.1.1 0 .14l4.56 2c-1.54.92-2.91 1.76-3.51 2.14a.858.858 0 0 1-.726.088L1.339 7.39a.864.864 0 0 0-.586.002l-.754.308l2.52 2.1a.879.879 0 0 0 .926.128C4.649 9.39 7.819 7.93 10.179 6.7c5.24-2.78 5.82-3.26 5.82-3.7c0-.69-2-1.4-3.43-.74zM0 13h16v1H0v-1z"/>`
	flipHInnerSVG                = `<path fill="currentColor" d="m0 15l6-5l-6-4.9zm9-4.9l6 4.9V5l-6 5.1zm5 2.8l-3.4-2.8l3.4-3v5.8zM7 5h1v1H7V5zm0-2h1v1H7V3zm0 4h1v1H7V7zm0 2h1v1H7V9zm0 2h1v1H7v-1zm0 2h1v1H7v-1zm0 2h1v1H7v-1z"/><path fill="currentColor" d="M7.5 1c1.3 0 2.6.7 3.6 1.9L10 4h3V1l-1.2 1.2C10.6.8 9.1 0 7.5 0C5.6 0 3.9 1 2.6 2.9l.8.6C4.5 1.9 5.9 1 7.5 1z"/>`
	flipVInnerSVG                = `<path fill="currentColor" d="m1 1l5 6l4.94-6H1zm4.94 9L1 16h10zm-2.82 5l2.83-3.44l3 3.44H3.12zM10 8h1v1h-1V8zm2 0h1v1h-1V8zM8 8h1v1H8V8zM6 8h1v1H6V8zM4 8h1v1H4V8zM2 8h1v1H2V8zM0 8h1v1H0V8z"/><path fill="currentColor" d="M15 8.47a4.807 4.807 0 0 1-1.879 3.632L12 11v3h3l-1.18-1.18A5.757 5.757 0 0 0 16 8.478V8.47a6.062 6.062 0 0 0-2.884-4.905l-.596.805a5.095 5.095 0 0 1 2.479 4.087z"/>`
	folderInnerSVG               = `<path fill="currentColor" d="M16 15H0V4h1l1-2h4l1 2h9z"/>`
	folderAddInnerSVG            = `<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v11h14v-1H1V5h.62l1-2h2.57l1.19 2H13v1h1z"/><path fill="currentColor" d="M14 7h-2v2h-2v2h2v2h2v-2h2V9h-2V7z"/>`
	folderOInnerSVG              = `<path fill="currentColor" d="M7 4L6 2H2L1 4H0v11h16V4H7zm8 10H1V5h.6l1-2h2.6l1.2 2H15v9z"/>`
	folderOpenInnerSVG           = `<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v9.5L3 6z"/><path fill="currentColor" d="M3.7 7L.5 15h12.8l2.5-8z"/>`
	folderOpenOInnerSVG          = `<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v11h14l2-9h-2zm.9 1l-1.6 7l-11.9-.1L3.7 7h11.2zM1 5h.6l1-2h2.6l1.2 2H13v1H3l-2 5.9V5z"/>`
	folderRemoveInnerSVG         = `<path fill="currentColor" d="M13 12.41V14H1V5h.62l1-2h2.57l1.19 2H13v2.59l1-1V4H7L6 2H2L1 4H0v11h14v-1.59l-1-1z"/><path fill="currentColor" d="m16 8l-1-1l-2 2l-2-2l-1 1l2 2l-2 2l1 1l2-2l2 2l1-1l-2-2l2-2z"/>`
	folderSearchInnerSVG         = `<path fill="currentColor" d="M13 13.47V14H1V5h.62l1-2h2.57l1.19 2H13v.91c.385.179.716.407 1.001.681L14 4H7L6 2H2L1 4H0v11h14v-.53z"/><path fill="currentColor" d="m15.78 12.72l-1.92-1.92a.727.727 0 0 0-.325-.179a3.014 3.014 0 0 0 .468-1.618a3 3 0 1 0-1.371 2.52c.02.136.083.248.169.337l1.92 1.92a.75.75 0 0 0 1.059-1.061zM11 11a2 2 0 1 1-.001-3.999A2 2 0 0 1 11 11z"/>`
	fontInnerSVG                 = `<path fill="currentColor" d="M12 16h3L9 0H7L1 16h3l1.9-5h4.2l1.9 5zM6.7 9L8 5.4L9.3 9H6.7z"/>`
	formInnerSVG                 = `<path fill="currentColor" d="M15 2v2H6V2h9zm1-1H5v4h11V1zM0 1h4v4H0V1zm15 6v2H6V7h9zm1-1H5v4h11V6zM0 6h4v4H0V6zm15 6v2H6v-2h9zm1-1H5v4h11v-4zM0 11h4v4H0v-4z"/>`
	forwardInnerSVG              = `<path fill="currentColor" d="M0 1v14l8-7zm8 0v14l8-7z"/>`
	frownOInnerSVG               = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm.3 6.3c-.7-1.1-2-1.8-3.3-1.8s-2.6.7-3.3 1.8l-.8-.6c.9-1.4 2.4-2.2 4.1-2.2s3.2.8 4.1 2.2l-.8.6z"/>`
	funcionInnerSVG              = `<path fill="currentColor" d="M10 0S7.9 0 7.3 3l-.4 2H5l-.5 1h2.2l-1.4 7c-.4 2-1.9 2-1.9 2h-1L2 16h3s2.1 0 2.7-3l1.4-7h2.4l.5-1H9.3l.4-2c.4-2 1.8-2 1.8-2h1l.5-1h-3z"/>`
	funnelInnerSVG               = `<path fill="currentColor" d="M6 11h4v4H6v-4zm7.6-6L16 1H0l2.4 4h11.2zM3 6l2.4 4h5.2L13 6H3z"/>`
	gamepadInnerSVG              = `<path fill="currentColor" d="M12.16 2a9.844 9.844 0 0 1-4.149 1a9.968 9.968 0 0 1-4.229-1.026C1.171 2 .001 3.17.001 5.84v6A2.19 2.19 0 0 0 2.191 14h.232a2.19 2.19 0 0 0 2.074-1.485C4.802 11.6 5.642 10 6.582 10h2.84c.94 0 1.78 1.6 2.08 2.5A2.194 2.194 0 0 0 13.58 14h.232c1.21 0 2.19-.98 2.19-2.19v-6c0-2.64-1.17-3.81-3.84-3.81zM5 7H4v1H3V7H2V6h1V5h1v1h1v1zm5.06 1.11a1.06 1.06 0 1 1 .001-2.121a1.06 1.06 0 0 1-.001 2.121zM13 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/>`
	gavelInnerSVG                = `<path fill="currentColor" d="M6.4 4.1c-.4-.4-.4-.9-.1-1.2L8.9.3c.3-.3.8-.3 1.2 0l.1.1c.3.3.3.8 0 1.2L7.6 4.1c-.3.3-.9.3-1.2 0zM12 9.7c-.4-.4-.4-.9-.1-1.3l2.6-2.6c.3-.3.8-.3 1.2 0l.1.1c.3.3.3.8 0 1.2l-2.6 2.6c-.4.3-.9.3-1.2 0zm-2-2L8.3 6c-.4-.4-.4-1 0-1.4l2.3-2.3c.4-.4 1-.4 1.4 0L13.7 4c.4.4.4 1 0 1.4l-2.3 2.3c-.4.4-1 .4-1.4 0zm-6 6.5c.6-.6 4-5.6 4.5-5.3c.4.2 1-.5 1-.5L7.6 6.5s-.7.6-.5 1c.3.5-4.7 3.9-5.3 4.5c0 0-2.8 2.2-1.4 3.6S4 14.2 4 14.2z"/>`
	giftInnerSVG                 = `<path fill="currentColor" d="M10.1 5c2-.3 3.9-1.1 2.2-3.6c-.7-1-1.4-1.4-2-1.4c-1 0-1.7 1.1-2.3 2.2C7.4 1.1 6.7 0 5.7 0c-.6 0-1.3.4-2 1.4c-1.8 2.5.2 3.3 2.2 3.6H0v3h16V5h-5.9zm.2-4c.1 0 .5.1 1.2 1c.5.7.6 1.1.5 1.3c-.2.3-1.3.7-3.3.8c0-.2-.1-.4-.2-.6C9.1 2.1 9.8 1 10.3 1zM4 3.3c-.1-.2 0-.6.5-1.3c.7-.9 1.1-1 1.2-1c.5 0 1.2 1.1 1.8 2.5c-.1.2-.2.4-.2.6c-2-.1-3.1-.5-3.3-.8zM7 7V5h2v2H7zm2 8H7V9H1v7h14V9H9z"/>`
	glassInnerSVG                = `<path fill="currentColor" d="M11 15H9V7l6-7H0l6 7v8H4c-2 0-2 1-2 1h11s0-1-2-1zm1.9-14l-1.8 2H3.9L2.2 1h10.7zM7 15V7h1v8H7z"/>`
	glassesInnerSVG              = `<path fill="currentColor" d="M15.5 7h-.7c-.4-1.2-1.5-2-2.8-2s-2.4.9-2.8 2.1c-.3-.4-.7-.6-1.2-.6s-.9.2-1.2.6C6.4 5.9 5.3 5 4 5s-2.4.9-2.8 2H.5c-.3 0-.5.2-.5.5s.2.5.5.5H1c0 1.7 1.3 3 3 3c1.5 0 2.7-1.1 3-2.5c.3 0 .5-.2.5-.5s.2-.5.5-.5s.5.2.5.5s.2.5.5.5c.2 1.4 1.5 2.5 3 2.5c1.7 0 3-1.3 3-3h.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5zM4 10c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2zm8 0c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2z"/>`
	globeInnerSVG                = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm5.2 5.3c.4 0 .7.3 1.1.3c-.3.4-1.6.4-2-.1c.3-.1.5-.2.9-.2zM1 8c0-.4 0-.8.1-1.3c.1 0 .2.1.3.1c0 0 .1.1.1.2c0 .3.3.5.5.5c.8.1 1.1.8 1.8 1c.2.1.1.3 0 .5c-.6.8-.1 1.4.4 1.9c.5.4.5.8.6 1.4c0 .7.1 1.5.4 2.2C2.7 13.3 1 10.9 1 8zm7 7c-.7 0-1.5-.1-2.1-.3c-.1-.2-.1-.4 0-.6c.4-.8.8-1.5 1.3-2.2c.2-.2.4-.4.4-.7c0-.2.1-.5.2-.7c.3-.5.2-.8-.2-.9c-.8-.2-1.2-.9-1.8-1.2s-1.2-.5-1.7-.2c-.2.1-.5.2-.5-.1c0-.4-.5-.7-.4-1.1c-.1 0-.2 0-.3.1s-.2.2-.4.1c-.2-.2-.1-.4-.1-.6c.1-.2.2-.3.4-.4c.4-.1.8-.1 1 .4c.3-.9.9-1.4 1.5-1.8c0 0 .8-.7.9-.7s.2.2.4.3c.2 0 .3 0 .3-.2c.1-.5-.2-1.1-.6-1.2c0-.1.1-.1.1-.1c.3-.1.7-.3.6-.6c0-.4-.4-.6-.8-.6c-.2 0-.4 0-.6.1c-.4.2-.9.4-1.5.4C5.2 1.4 6.6 1 8 1h.8c-.6.1-1.2.3-1.6.5c.6.1.7.4.5.9c-.1.2 0 .4.2.5s.4.1.5-.1c.2-.3.6-.4.9-.5c.4-.1.7-.3 1-.7c0-.1.1-.1.2-.2c.6.2 1.2.6 1.8 1c-.1 0-.1.1-.2.1c-.2.2-.5.3-.2.7c.1.2 0 .3-.1.4c-.2.1-.3 0-.4-.1s-.1-.3-.4-.3c-.1.2-.4.3-.4.6c.5 0 .4.4.5.7c-.6.1-.8.4-.5.9c.1.2-.1.3-.2.4c-.4.6-.8 1-.8 1.7s.5 1.4 1.3 1.3c.9-.1.9-.1 1.2.7c0 .1.1.2.1.3c.1.2.2.4.1.6c-.3.8.1 1.4.4 2c.1.2.2.3.3.4c-1.3 1.4-3 2.2-5 2.2z"/>`
	globeWireInnerSVG            = `<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm6.8 9.5c0 .5-.7.66-2 1a11.68 11.68 0 0 0 .229-1.98l2.001-.02c0 .36-.08.5-.16 1zm-13.6 0c-.1-.5-.15-.64-.2-1h2c.024.723.106 1.411.244 2.079C1.9 10.16 1.2 10 1.2 9.5zm0-3c0-.5.7-.66 2-1A11.835 11.835 0 0 0 3 7.489L1 7.5c0-.36.08-.5.16-1zM8.5 5c1.13.013 2.226.107 3.298.277c.047.643.165 1.41.201 2.199L8.5 7.501v-2.5zm0-1V1.06c1.17.27 2.2 1.47 2.84 3.15A24.21 24.21 0 0 0 8.523 4zm-1-2.94V4a25.617 25.617 0 0 0-2.968.214C5.3 2.53 6.33 1.33 7.5 1.06zM7.5 5v2.5H4c.031-.806.142-1.571.326-2.307c.932-.08 2.035-.177 3.158-.193zM4 8.5h3.5V11a22.663 22.663 0 0 1-3.298-.277c-.047-.643-.165-1.41-.201-2.199zM7.5 12v2.94c-1.17-.27-2.2-1.47-2.84-3.15a24.21 24.21 0 0 0 2.817.21zm1 2.94V12a25.617 25.617 0 0 0 2.968-.214C10.7 13.47 9.67 14.67 8.5 14.94zm0-3.94V8.5H12a11.247 11.247 0 0 1-.326 2.307c-.932.08-2.035.177-3.158.193zM15 7.5h-2a12.229 12.229 0 0 0-.244-2.079c1.354.399 2.014.559 2.014 1.079c.13.5.18.64.23 1zm-.7-2.59a9.585 9.585 0 0 0-1.726-.5c-.361-1.019-.809-1.898-1.389-2.672c1.355.726 2.413 1.811 3.067 3.131zM4.84 1.76a8.24 8.24 0 0 0-1.305 2.581c-.699.189-1.299.365-1.874.593c.751-1.39 1.823-2.475 3.139-3.156zm-3.11 9.33c.506.204 1.106.38 1.726.5c.361 1.019.809 1.898 1.389 2.672c-1.367-.722-2.436-1.807-3.097-3.131zm9.44 3.15a8.25 8.25 0 0 0 1.295-2.581c.699-.189 1.299-.365 1.874-.593c-.751 1.39-1.823 2.475-3.139 3.156z"/>`
	golfInnerSVG                 = `<path fill="currentColor" d="M7 2a2 2 0 1 1-3.999.001A2 2 0 0 1 7 2z"/><path fill="currentColor" d="M9.8 1.8c-.2-.5-1.7-.1-2 .5c-.2.3-.2 1.2-1.2 1.9c-.8.5-1.6.5-1.6.5c-.3.6-.1 1.1.2 1.6c.5.9.6 1.8.7 2.8c.1 1.3-.5 2.4-2.3 3.2c-.8.3-1.3.9-1 1.9c0 0 2-.3 3.1-1.2c1.5-1.2 1.8-2.3 1.8-2.3s.1.7 0 1.9c-.1 1-.2 1.5-.4 2.2S7.4 16 8 16s1-.4 1-1l.3-1.9c.3-2.1 0-4.3-.8-6.3c0-.1-.1-.1-.1-.2c-.6-1.6.2-2.6.6-3c.3-.4 1.2-1.2.8-1.8zM12 0v10h1V4l3-2zm4 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0zM1 8.4l3.7-3.7l-.7-.3L.2 8s-.4.7.1 1.7s1.6.3 1.6.3c.4-.2.2-.4 0-.6s-.9-1-.9-1z"/>`
	googlePlusInnerSVG           = `<path fill="currentColor" d="M16 3.9h-2.8V1.3h-.6v2.6H9.9v.8h2.7v2.6h.6V4.7H16zM6.9 9c-.4-.2-1.1-.9-1.1-1.3s.1-.7.8-1.2c.7-.5 1.2-1.2 1.2-2.1c0-1.1-.5-2.1-1.3-2.4h1.3l.9-.7H4.5C2.6 1.3.9 2.7.9 4.4s1.3 3 3.2 3h.4c-.2.2-.2.4-.2.7c0 .5.3.8.6 1.2h-.7c-2.3 0-4.1 1.5-4.1 3s2 2.5 4.3 2.5c2.6 0 4.1-1.5 4.1-3c-.1-1.3-.5-2-1.6-2.8zM4.7 6.9c-1.1 0-2.1-1.2-2.3-2.6S2.9 1.8 4 1.8S6.1 3 6.3 4.4S5.8 7 4.7 6.9zm-.4 7.2c-1.6 0-2.8-1-2.8-2.2s1.4-2.2 3-2.2c.4 0 .7.1 1 .2c.9.6 1.5.9 1.7 1.6c0 .1.1.3.1.4c0 1.2-.8 2.2-3 2.2z"/>`
	googlePlusSquareInnerSVG     = `<path fill="currentColor" d="M5 3.4c-.8 0-1.3.8-1.2 1.8c.1 1.1.9 1.9 1.7 2c.8 0 1.3-.8 1.2-1.9c-.1-1-.9-1.9-1.7-1.9zm.4 5.9c-1.2 0-2.3.7-2.3 1.6s.9 1.7 2.1 1.7c1.7 0 2.3-.7 2.3-1.6v-.3c-.1-.5-.6-.8-1.3-1.2c-.2-.2-.5-.2-.8-.2z"/><path fill="currentColor" d="M0 0v16h16V0H0zm7.9 5.3c0 .7-.4 1.2-.9 1.6s-.6.6-.6.9c0 .3.5.8.8 1c.8.6 1.1 1.1 1.1 2c0 1.1-1.1 2.3-3.1 2.3c-1.7 0-3.2-.7-3.2-1.8C2 10.1 3.3 9 5.1 9h.5c-.2-.3-.4-.6-.4-.9c0-.2.1-.4.2-.6h-.3c-1.4 0-2.4-1-2.4-2.3C2.7 4 4 2.9 5.4 2.9h3.1l-.7.6h-1c.7.2 1.1 1 1.1 1.8zm6.1.2h-2.1v2h-.5v-2h-2V5h2V3h.5v2H14v.5z"/>`
	grabInnerSVG                 = `<path fill="currentColor" d="M12.6 4H12c0-.2-.2-.6-.4-.8s-.5-.4-1.1-.4c-.2 0-.4 0-.6.1c-.1-.2-.2-.3-.3-.5c-.2-.2-.5-.4-1.1-.4c-.8 0-1.2.5-1.4 1c-.1 0-.3-.1-.5-.1c-.5 0-.8.2-1.1.4C5 3.9 5 4.7 5 4.8v.4c-.6 0-1.1.2-1.4.5C3 6.4 3 7.3 3 8.5v.7c0 1.4.7 2.1 1.4 2.8l.3.4C6 13.6 7.2 14 9.8 14c2.9 0 4.2-1.6 4.2-5.1V6.4c0-.7-.2-2.1-1.4-2.4zm-2.1-.2c.4 0 .5.4.5.6v.8c0 .3.2.5.4.5c.3 0 .5-.1.5-.4c0 0 0-.4.4-.3c.6.2.7 1.1.7 1.3v2.6c0 3.4-1.3 4.1-3.2 4.1c-2.4 0-3.3-.3-4.3-1.3l-.4-.4C4.4 10.6 4 10.2 4 9.2v-.6c0-1 0-1.8.3-2.1c.1-.2.4-.3.7-.3V7l-.3 1.2c0 .1 0 .1.1.1c.1.1.2 0 .2 0l1-1.2V5c0-.1 0-.6.2-.8c.1-.1.2-.2.4-.2c.3 0 .4.2.4.4v.4c0 .2.2.5.5.5S8 5 8 4.8V3.5c0-.1 0-.5.5-.5c.3 0 .5.2.5.5v1.2c0 .3.2.6.5.6s.5-.3.5-.5v-.5c0-.3.2-.5.5-.5z"/>`
	gridInnerSVG                 = `<path fill="currentColor" d="M0 0v16h16V0H0zm5 15H1v-4h4v4zm0-5H1V6h4v4zm0-5H1V1h4v4zm5 10H6v-4h4v4zm0-5H6V6h4v4zm0-5H6V1h4v4zm5 10h-4v-4h4v4zm0-5h-4V6h4v4zm0-5h-4V1h4v4z"/>`
	gridBevelInnerSVG            = `<path fill="currentColor" d="M14 2V1H1v13h1v1h13V2h-1zM5 13H2v-3h3v3zm0-4H2V6h3v3zm0-4H2V2h3v3zm4 8H6v-3h3v3zm0-4H6V6h3v3zm0-4H6V2h3v3zm4 8h-3v-3h3v3zm0-4h-3V6h3v3zm0-4h-3V2h3v3z"/>`
	gridBigInnerSVG              = `<path fill="currentColor" d="M0 0h7v7H0V0zm9 0h7v7H9V0zM0 9h7v7H0V9zm9 0h7v7H9V9z"/>`
	gridBigOInnerSVG             = `<path fill="currentColor" d="M0 7h7V0H0v7zm1-6h5v5H1V1zm8-1v7h7V0H9zm6 6h-5V1h5v5zM0 16h7V9H0v7zm1-6h5v5H1v-5zm8 6h7V9H9v7zm1-6h5v5h-5v-5z"/>`
	gridHInnerSVG                = `<path fill="currentColor" d="M0 0v16h16V0H0zm5 15H1V1h4v14zm5 0H6V1h4v14zm5 0h-4V1h4v14z"/>`
	gridSmallInnerSVG            = `<path fill="currentColor" d="M0 0h4v4H0V0zm0 6h4v4H0V6zm0 6h4v4H0v-4zM6 0h4v4H6V0zm0 6h4v4H6V6zm0 6h4v4H6v-4zm6-12h4v4h-4V0zm0 6h4v4h-4V6zm0 6h4v4h-4v-4z"/>`
	gridSmallOInnerSVG           = `<path fill="currentColor" d="M0 4h4V0H0v4zm1-3h2v2H1V1zm-1 9h4V6H0v4zm1-3h2v2H1V7zm-1 9h4v-4H0v4zm1-3h2v2H1v-2zm5-9h4V0H6v4zm1-3h2v2H7V1zm-1 9h4V6H6v4zm1-3h2v2H7V7zm-1 9h4v-4H6v4zm1-3h2v2H7v-2zm5-13v4h4V0h-4zm3 3h-2V1h2v2zm-3 7h4V6h-4v4zm1-3h2v2h-2V7zm-1 9h4v-4h-4v4zm1-3h2v2h-2v-2z"/>`
	gridVInnerSVG                = `<path fill="currentColor" d="M16 0H0v16h16V0zM1 5V1h14v4H1zm0 5V6h14v4H1zm0 5v-4h14v4H1z"/>`
	groupInnerSVG                = `<path fill="currentColor" d="M5 16v-5.3c-.6-.3-1-1-1-1.7V5c0-.7.4-1.3 1-1.7V3c0-1.1-.9-2-2-2s-2 .9-2 2s.9 2 2 2H1c-.5 0-1 .5-1 1v4c0 .5.5 1 1 1v5h4zM15 5h-2c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2v.3c.6.4 1 1 1 1.7v4c0 .7-.4 1.4-1 1.7V16h4v-5c.5 0 1-.5 1-1V6c0-.5-.5-1-1-1zm-5-3a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2z"/><path fill="currentColor" d="M10 4H6c-.5 0-1 .5-1 1v4c0 .5.5 1 1 1v6h4v-6c.5 0 1-.5 1-1V5c0-.5-.5-1-1-1z"/>`
	hammerInnerSVG               = `<path fill="currentColor" d="m6 2l7 7l3-3l-4.48-4.48S8.55 2.55 7 1zm2.8 3.79L.27 14.31a.998.998 0 0 0 0 1.361a.998.998 0 0 0 1.371.049l8.569-8.519z"/>`
	handInnerSVG                 = `<path fill="currentColor" d="M13.5 2.4c-.4-.4-1-.5-1.5-.3c0-.3-.1-.6-.4-.9c-.2-.2-.6-.4-1.1-.4c-.3 0-.5.1-.7.1c0-.2-.1-.3-.2-.5c-.5-.6-1.5-.6-2 0c-.2.2-.4.4-.4.6C7 1 6.8.9 6.6.9c-.5 0-.8.2-1.1.5C5 1.9 5 2.7 5 2.7v3.8c-.3-.3-.8-.8-1.5-.8c-.2 0-.5.1-.7.2c-.4.2-.6.5-.7.9c-.3 1 .6 2.4.6 2.5c.1.1 1.2 2.7 2.2 3.8C5.9 14.3 7 15 9.8 15c2.9 0 4.2-1.6 4.2-5.1V4.4c0-.1.1-1.3-.5-2zM8 2c0-.3-.1-1 .5-1c.5 0 .5.5.5 1v4c0 .3.2.5.5.5s.5-.2.5-.5V2.2s0-.4.5-.4c.6 0 .5.9.5.9V6c0 .3.2.5.5.5s.5-.2.5-.5V3.6c0-.1 0-.6.5-.6s.5 1 .5 1v5.9c0 3.4-1.3 4.1-3.2 4.1c-2.4 0-3.3-.5-4.1-1.6c-.9-1-2.1-3.6-2.1-3.7c-.3-.3-.7-1.2-.6-1.6c0-.1.1-.2.2-.3c.1 0 .2-.1.2-.1c.4 0 .8.5.9.7l.6.9c.1.2.4.3.6.2c.4 0 .5-.2.5-.4V2.9c0-.4 0-1 .5-1c.4 0 .5.3.5.8V6c0 .3.2.5.5.5S8 6.3 8 6z"/>`
	handleCornerInnerSVG         = `<path fill="currentColor" d="M6.7 16L16 6.7V5.3L5.3 16zm3 0L16 9.7V8.3L8.3 16zm3 0l3.3-3.3v-1.4L11.3 16zm3 0l.3-.3v-1.4L14.3 16z"/>`
	handsUpInnerSVG              = `<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2z"/><path fill="currentColor" d="M6 16h1.5v-5h1v5H10V7l-.001-.052c0-.521.194-.997.513-1.36L13.79 2.27a.73.73 0 1 0-.998-1.003L10.43 3.65c-.212.216-.508.35-.835.35H6.404c-.327 0-.622-.134-.834-.35L3.25 1.26a.73.73 0 1 0-1.003.998L5.49 5.59c.317.361.511.836.511 1.358L6 7.003V16z"/>`
	handshakeInnerSVG            = `<path fill="currentColor" d="M13 3a5.393 5.393 0 0 1-1.902 1.178c-.748.132-2.818-.828-3.838.152c-.17.17-.38.34-.6.51c-.48-.21-1.22-.53-1.76-.84S3 3 3 3L0 6.5s.74 1 1.2 1.66c.3.44.67 1.11.91 1.56l-.34.4a.876.876 0 0 0 .15 1a.833.833 0 0 0 1.002-.002a.62.62 0 0 0 .077.881a.994.994 0 0 0 1.006-.002a.806.806 0 0 0-.003 1.005a1.012 1.012 0 0 0 .892-.114a.822.822 0 0 0 .187.912a1.093 1.093 0 0 0 1.054-.092l.516-.467c.472.47 1.123.761 1.842.761l.061-.001a1.311 1.311 0 0 0 1.094-.791c.146.056.312.094.488.094c.236 0 .455-.068.64-.185c.585-.387.445-.687.445-.687a1.07 1.07 0 0 0 1.229-.279a.996.996 0 0 0 .138-1.215a.036.036 0 0 0 .021.005c.421 0 .787-.232.978-.574a1.564 1.564 0 0 0-.191-1.48l.003.005c.82-.16.79-.57 1.19-1.17a4.725 4.725 0 0 1 1.387-1.208zm-.05 7.06c-.44.44-.78.25-1.53-.32S9.18 8.1 9.18 8.1c.061.305.202.57.401.781c.319.359 1.269 1.179 1.719 1.599c.28.26 1 .78.58 1.18s-.75 0-1.44-.56s-2.23-1.94-2.23-1.94a.937.937 0 0 0 .27.72c.17.2 1.12 1.12 1.52 1.54s.75.67.41 1s-1.03-.19-1.41-.58c-.59-.57-1.76-1.63-1.76-1.63l-.001.053c0 .284.098.544.263.75c.288.378.848.868 1.188 1.248s.54.7 0 1s-1.34-.44-1.69-.8v-.002a.411.411 0 0 0-.1-.269a.896.896 0 0 0-.906-.188A.609.609 0 0 0 6 11.1a.754.754 0 0 0-.912.001a.61.61 0 0 0-.085-.95a1 1 0 0 0-1.174.08a.66.66 0 0 0-.068-.911a.996.996 0 0 0-1.186-.128L1.91 8.069c-.46-.73-1-1.49-1-1.49l2.28-2.77s.81.5 1.48.88c.33.19.9.44 1.33.64c-.68.51-1.25 1-1.08 1.34a1.834 1.834 0 0 0 2.087.036a2.41 2.41 0 0 1 1.343-.403c.347 0 .677.072.976.203c.554.374 1.574 1.294 2.504 1.874c1.17.85 1.4 1.4 1.12 1.68z"/>`
	harddriveInnerSVG            = `<path fill="currentColor" d="M13 1H3L.3 9h15.4zM0 10v5h16v-5H0zm3 3H2v-1h1v1zm4 0H4v-1h3v1z"/>`
	harddriveOInnerSVG           = `<path fill="currentColor" d="M2 12h1v1H2v-1zm2 0h3v1H4v-1z"/><path fill="currentColor" d="M13 1H3l-3 9v5h16v-5l-3-9zM3.7 2h8.6l2.7 8H1.1l2.6-8zM1 14v-3h14v3H1z"/>`
	hashInnerSVG                 = `<path fill="currentColor" d="M15 6V4h-2.6l.6-2.8l-2-.4l-.7 3.2h-3L8 1.2L6 .8L5.3 4H2v2h2.9L4 10H1v2h2.6L3 14.8l2 .4l.7-3.2h3L8 14.8l2 .4l.7-3.2H14v-2h-2.9l.9-4h3zm-6 4H6l1-4h3l-1 4z"/>`
	headerInnerSVG               = `<path fill="currentColor" d="M11 0v7H5V0H2v16h3V9h6v7h3V0z"/>`
	headphonesInnerSVG           = `<path fill="currentColor" d="M14 8.3V6c0-3.3-2.7-6-6-6S2 2.7 2 6v2.3c-1.2.5-2 1.7-2 3.1v1.2c0 1.8 1.3 3.2 3 3.4h2V8H4V6c0-2.2 1.8-4 4-4s4 1.8 4 4v2h-1v8h2c1.7-.2 3-1.7 3-3.4v-1.2c0-1.4-.8-2.6-2-3.1zM4 15H3V9h1v6zm9 0h-1V9h1v6z"/>`
	headsetInnerSVG              = `<path fill="currentColor" d="M14.82 8a3.016 3.016 0 0 0-1.799-1.813L13 4.5C13 2 10.53 0 7.5 0S2 2 2 4.5v1.68A3.006 3.006 0 0 0 0 9v1a3 3 0 0 0 3 3h1V6H3V4.5C3 2.57 5 1 7.5 1S12 2.57 12 4.5V6h-1v7h1a3 3 0 0 0 3-3v1.73A3.27 3.27 0 0 1 11.73 15H10a1 1 0 0 0-1-1H8a1 1 0 0 0 0 2h3.73A4.27 4.27 0 0 0 16 11.73V8h-1.18z"/>`
	healthCardInnerSVG           = `<path fill="currentColor" d="M15 3v10H1V3h14zm1-1H0v12h16V2z"/><path fill="currentColor" d="M9 5h5v1H9V5zm0 2h5v1H9V7zm0 2h2v1H9V9zM6.5 5c-.6 0-1.1.6-1.5 1c-.4-.4-.9-1-1.5-1c-1.5 0-2.1 1.9-1 2.9L5 10l2.5-2.1C8.6 6.9 8 5 6.5 5z"/>`
	heartInnerSVG                = `<path fill="currentColor" d="M12 2S9 2 8 5C7 2 4 2 4 2C1.8 2 0 3.8 0 6c0 4.1 8 9 8 9s8-5 8-9c0-2.2-1.8-4-4-4z"/>`
	heartOInnerSVG               = `<path fill="currentColor" d="M11.7 2C10.8 2 9 2.5 8 4.1C7 2.5 5.2 2 4.2 2C1.9 2 0 3.9 0 6.2c0 4 7.4 8.5 7.7 8.7l.3.2l.3-.2c.3-.2 7.7-4.8 7.7-8.7C16 3.9 14.1 2 11.7 2zM8 13.9c-2.2-1.4-7-5-7-7.7C1 4.4 2.5 3 4.2 3c.1 0 2.5.1 3.3 2.4L8 6.8l.5-1.4C9.3 3.1 11.7 3 11.8 3C13.5 3 15 4.4 15 6.2c0 2.7-4.8 6.3-7 7.7z"/>`
	homeInnerSVG                 = `<path fill="currentColor" d="M8 1.4L6 2.7V1H4v3L0 6.6l.6.8L8 2.6l7.4 4.8l.6-.8z"/><path fill="currentColor" d="M8 4L2 8v7h5v-3h2v3h5V8z"/>`
	homeOInnerSVG                = `<path fill="currentColor" d="M16 6.6L8 1.4L6 2.7V1H4v3L0 6.6l1.9 2.7l.1-.1V15h5v-4h2v4h5V9.2l.1.1L16 6.6zm-14.6.3L8 2.6l6.6 4.3l-.7 1L8 4L2.1 7.9l-.7-1zM13 14h-3v-4H6v4H3V8.6l5-3.3l5 3.3V14z"/>`
	hospitalInnerSVG             = `<path fill="currentColor" d="M15 4V0H8v4H0v12h6v-3h4v3h6V4h-1zM4 11H2V9h2v2zm0-3H2V6h2v2zm3 3H5V9h2v2zm0-3H5V6h2v2zm3-5V2h1V1h1v1h1v1h-1v1h-1V3h-1zm1 8H9V9h2v2zm0-3H9V6h2v2zm3 3h-2V9h2v2zm0-3h-2V6h2v2z"/>`
	hourglassInnerSVG            = `<path fill="currentColor" d="M6.16 4.6A4.054 4.054 0 0 1 8 7.994c0-1.415.726-2.66 1.825-3.384c.23-.199.426-.395.609-.602L5.56 4.001c.19.214.386.41.593.594z"/><path fill="currentColor" d="M11.18 6.06A4.399 4.399 0 0 0 13 2.5V2h1V0H2v2h1v.5a4.391 4.391 0 0 0 1.808 3.551A2.564 2.564 0 0 1 6 7.99a2.755 2.755 0 0 1-1.209 2.003a4.441 4.441 0 0 0-1.79 3.503v.503h-1v2h12v-2h-1v-.5a4.435 4.435 0 0 0-1.769-3.492a2.762 2.762 0 0 1-1.23-1.996a2.551 2.551 0 0 1 1.169-1.946zM9 8a3.693 3.693 0 0 0 1.519 2.763A3.477 3.477 0 0 1 12 13.495V14h-1.77c-.7-.87-1.71-2-2.23-2s-1.53 1.13-2.23 2H4v-.5a3.472 3.472 0 0 1 1.459-2.723a3.698 3.698 0 0 0 1.54-2.766a3.482 3.482 0 0 0-1.498-2.683a3.438 3.438 0 0 1-1.502-2.827v-.5h8v.5a3.426 3.426 0 0 1-1.479 2.813a3.487 3.487 0 0 0-1.521 2.678z"/>`
	hourglassEmptyInnerSVG       = `<path fill="currentColor" d="M11.18 6.06A4.399 4.399 0 0 0 13 2.5V2h1V0H2v2h1v.5a4.391 4.391 0 0 0 1.808 3.551A2.564 2.564 0 0 1 6 7.99a2.755 2.755 0 0 1-1.209 2.003a4.441 4.441 0 0 0-1.79 3.503v.503h-1v2h12v-2h-1v-.5a4.435 4.435 0 0 0-1.769-3.492a2.762 2.762 0 0 1-1.23-1.996a2.551 2.551 0 0 1 1.169-1.946zM9 8a3.693 3.693 0 0 0 1.519 2.763A3.477 3.477 0 0 1 12 13.495V14H4v-.5a3.472 3.472 0 0 1 1.459-2.723a3.698 3.698 0 0 0 1.54-2.766a3.482 3.482 0 0 0-1.498-2.683a3.438 3.438 0 0 1-1.502-2.827v-.5h8v.5a3.426 3.426 0 0 1-1.479 2.813a3.487 3.487 0 0 0-1.521 2.678z"/>`
	hourglassEndInnerSVG         = `<path fill="currentColor" d="M11.18 6.06A4.399 4.399 0 0 0 13 2.5V2h1V0H2v2h1v.5a4.391 4.391 0 0 0 1.808 3.551A2.564 2.564 0 0 1 6 7.99a2.755 2.755 0 0 1-1.209 2.003a4.441 4.441 0 0 0-1.79 3.503v.503h-1v2h12v-2h-1v-.5a4.435 4.435 0 0 0-1.769-3.492a2.762 2.762 0 0 1-1.23-1.996a2.551 2.551 0 0 1 1.169-1.946zM9 8a3.693 3.693 0 0 0 1.519 2.763A3.477 3.477 0 0 1 12 13.495V14h-1s-1.62-3.5-3-3.5S5 14 5 14H4v-.5a3.472 3.472 0 0 1 1.459-2.723a3.698 3.698 0 0 0 1.54-2.766a3.482 3.482 0 0 0-1.498-2.683a3.438 3.438 0 0 1-1.502-2.827v-.5h8v.5a3.426 3.426 0 0 1-1.479 2.813a3.487 3.487 0 0 0-1.521 2.678z"/>`
	hourglassStartInnerSVG       = `<path fill="currentColor" d="M6.16 4.6A4.054 4.054 0 0 1 8 7.994c0-1.415.726-2.66 1.825-3.384a2.857 2.857 0 0 0 1.17-1.589L5 3.001c.191.67.603 1.225 1.15 1.594z"/><path fill="currentColor" d="M11.18 6.06A4.399 4.399 0 0 0 13 2.5V2h1V0H2v2h1v.5a4.391 4.391 0 0 0 1.808 3.551A2.564 2.564 0 0 1 6 7.99a2.755 2.755 0 0 1-1.209 2.003a4.441 4.441 0 0 0-1.79 3.503v.503h-1v2h12v-2h-1v-.5a4.435 4.435 0 0 0-1.769-3.492a2.762 2.762 0 0 1-1.23-1.996a2.551 2.551 0 0 1 1.169-1.946zM9 8a3.693 3.693 0 0 0 1.519 2.763A3.477 3.477 0 0 1 12 13.495V14H4v-.5a3.472 3.472 0 0 1 1.459-2.723a3.698 3.698 0 0 0 1.54-2.766a3.482 3.482 0 0 0-1.498-2.683a3.438 3.438 0 0 1-1.502-2.827v-.5h8v.5a3.426 3.426 0 0 1-1.479 2.813a3.487 3.487 0 0 0-1.521 2.678z"/>`
	inboxInnerSVG                = `<path fill="currentColor" d="M10 6V0H6v6H4l4 5l4-5z"/><path fill="currentColor" d="M13 1h-2v1h1.3l2.6 8H11v2H5v-2H1.1l2.6-8H5V1H3l-3 9v5h16v-5z"/>`
	indentInnerSVG               = `<path fill="currentColor" d="M0 0h16v3H0V0zm6 4h10v3H6V4zm0 4h10v3H6V8zm-6 4h16v3H0v-3zm0-7.5v6l4-3z"/>`
	infoInnerSVG                 = `<path fill="currentColor" d="M6 5h4v11H6V5zm4-3a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2z"/>`
	infoCircleInnerSVG           = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm1 13H7V6h2v7zm0-8H7V3h2v2z"/>`
	infoCircleOInnerSVG          = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M7 6h2v7H7V6zm0-3h2v2H7V3z"/>`
	inputInnerSVG                = `<path fill="currentColor" d="M16 5c0-.6-.4-1-1-1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V5zm-1 6H1V5h14v6z"/><path fill="currentColor" d="M2 6h1v4H2V6z"/>`
	insertInnerSVG               = `<path fill="currentColor" d="M14 16V5l-1 1v9H1V3h9l1-1H0v14z"/><path fill="currentColor" d="M16 1.4L14.6 0L7.8 6.8L6 5v5h5L9.2 8.2z"/>`
	institutionInnerSVG          = `<path fill="currentColor" d="M8 0L0 3v2h16V3zM0 14h16v2H0v-2zm16-7V6H0v1h1v5H0v1h16v-1h-1V7h1zM4 12H3V7h1v5zm3 0H6V7h1v5zm3 0H9V7h1v5zm3 0h-1V7h1v5z"/>`
	invoiceInnerSVG              = `<path fill="currentColor" d="M4.4 10.2c-.6.1-1.4-.3-1.7-.4l-.5.9s.9.4 1.7.5v.8h1v-.9c.9-.3 1.4-1.1 1.5-1.8c0-.8-.6-1.4-1.9-1.9c-.4-.2-1.1-.5-1.1-.9c0-.5.4-.8 1-.8c.7 0 1.4.3 1.4.3l.4-.9s-.5-.2-1.2-.4V4H4v.7c-.9.2-1.5.8-1.6 1.7c0 1.2 1.3 1.7 1.8 1.9c.6.2 1.3.6 1.3.9c0 .4-.4.9-1.1 1z"/><path fill="currentColor" d="M0 2v12h16V2H0zm15 11H1V3h14v10z"/><path fill="currentColor" d="M8 5h6v1H8V5zm0 2h6v1H8V7zm0 2h3v1H8V9z"/>`
	italicInnerSVG               = `<path fill="currentColor" d="M8 0h3L8 16H5z"/>`
	keyInnerSVG                  = `<path fill="currentColor" d="m8.1 7l-.6-.3L15 0h-2L6 6.1C5.7 6 5.4 6 5 6c-2.8 0-5 2.2-5 5s2.3 5 5 5s5-2.2 5-5c0-.6-.1-1.2-.3-1.7L11 8V6h2V4h2l1-1V0L8.1 7zM4 13.2c-.7 0-1.2-.6-1.2-1.2s.6-1.2 1.2-1.2s1.2.6 1.2 1.2s-.5 1.2-1.2 1.2z"/>`
	keyOInnerSVG                 = `<path fill="currentColor" d="M13 0L6 6.1C5.7 6 5.4 6 5 6c-2.8 0-5 2.2-5 5s2.3 5 5 5s5-2.2 5-5c0-.3 0-.6-.1-.9L11 9V7h2V5h2l1-1V0h-3zm-1 6h-1.7L12 4.6V6zm3-2.4l-.4.4h-1.9L15 2v1.6zm-7.7 4L8 8l2-1.7v2.3l-.8.8l-.3.4l.1.5c0 .2.1.5.1.7c0 2.2-1.8 4-4 4s-4-1.8-4-4s1.8-4 4-4c.3 0 .5 0 .8.1l.5.1l.4-.3L13.4 1H15L7.3 7.6z"/><path fill="currentColor" d="M6 11.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 6 11.5z"/>`
	keyboardInnerSVG             = `<path fill="currentColor" d="M0 4v9h16V4H0zm10 2h1v1h-1V6zM8 6h1v1H8V6zm2 2v1H9V8h1zM6 6h1v1H6V6zm2 2v1H7V8h1zM4 6h1v1H4V6zm2 2v1H5V8h1zM2 6h1v1H2V6zm1 5H2v-1h1v1zm0-3h1v1H3V8zm9 3H4v-1h8v1zm0-2h-1V8h1v1zm2 2h-1v-1h1v1zm0-2h-1V7h-1V6h2v3z"/>`
	keyboardOInnerSVG            = `<path fill="currentColor" d="M15 5v7H1V5h14zm1-1H0v9h16V4z"/><path fill="currentColor" d="M4 10h8v1H4v-1zm-2 0h1v1H2v-1zm11 0h1v1h-1v-1zm-2-2h1v1h-1V8zM9 8h1v1H9V8zM7 8h1v1H7V8zM5 8h1v1H5V8zM3 8h1v1H3V8zm7-2h1v1h-1V6zm2 0v1h1v2h1V6zM8 6h1v1H8V6zM6 6h1v1H6V6zM4 6h1v1H4V6zM2 6h1v1H2V6z"/>`
	laptopInnerSVG               = `<path fill="currentColor" d="M14 11V2H2v9H0v2h16v-2h-2zm-4 1H6v-1h4v1zm3-2H3V3h10v7z"/>`
	layoutInnerSVG               = `<path fill="currentColor" d="M0 0v16h16V0H0zm1 3h4v12H1V3zm14 12H6V3h9v12z"/>`
	levelDownInnerSVG            = `<path fill="currentColor" d="M5 1h6v11h2l-3 3l-3-3h2V3H3z"/>`
	levelDownBoldInnerSVG        = `<path fill="currentColor" d="m9 16l4-7h-3V0H3l2 3h2v6H4z"/>`
	levelLeftInnerSVG            = `<path fill="currentColor" d="M15 12V6H4V4L1 7l3 3V8h9v6z"/>`
	levelLeftBoldInnerSVG        = `<path fill="currentColor" d="m0 7l7-4v3h9v7l-3-2V9H7v3z"/>`
	levelRightInnerSVG           = `<path fill="currentColor" d="M1 12V6h11V4l3 3l-3 3V8H3v6z"/>`
	levelRightBoldInnerSVG       = `<path fill="currentColor" d="M16 7L9 3v3H0v7l3-2V9h6v3z"/>`
	levelUpInnerSVG              = `<path fill="currentColor" d="M11 15H5V4H3l3-3l3 3H7v9h6z"/>`
	levelUpBoldInnerSVG          = `<path fill="currentColor" d="m9 0l4 7h-3v9H3l2-3h2V7H4z"/>`
	lifebuoyInnerSVG             = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zM4 8c0-2.2 1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4s-4-1.8-4-4zm8.6 1.8c.3-.5.4-1.2.4-1.8s-.1-1.3-.4-1.8l1.5-1.5c.6 1 .9 2.1.9 3.3s-.3 2.3-.8 3.3l-1.6-1.5zm-1.3-8L9.8 3.4C9.3 3.1 8.6 3 8 3s-1.3.1-1.8.4L4.7 1.8C5.7 1.3 6.8 1 8 1s2.3.3 3.3.8zM1.8 4.7l1.5 1.5C3.1 6.7 3 7.4 3 8s.1 1.3.4 1.8l-1.5 1.5C1.3 10.3 1 9.2 1 8s.3-2.3.8-3.3zm2.9 9.5l1.5-1.5c.5.2 1.2.3 1.8.3s1.3-.1 1.8-.4l1.5 1.5c-1 .6-2.1.9-3.3.9s-2.3-.3-3.3-.8z"/>`
	lightbulbInnerSVG            = `<path fill="currentColor" d="M8 0a5 5 0 0 0-5 5a4.8 4.8 0 0 0 2.182 3.989L6 11a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1h.41c.342.55.915.929 1.581.999a2.122 2.122 0 0 0 1.594-.99L10 15a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1l.8-2A4.805 4.805 0 0 0 13 5.002A5.001 5.001 0 0 0 8 0zm2.25 8.21l-.25.17l-.11.29L9 10.81a.292.292 0 0 1-.27.19H7.22a.29.29 0 0 1-.219-.188L6.13 8.67L6 8.38l-.25-.18A3.914 3.914 0 0 1 4 5.003A4 4 0 1 1 12 5a3.905 3.905 0 0 1-1.736 3.201z"/><path fill="currentColor" d="M10.29 3a3.153 3.153 0 0 0-2.289-1L8 3a2.133 2.133 0 0 1 1.5.62c.278.386.459.858.499 1.37L11 4.999a3.785 3.785 0 0 0-.718-2.011z"/>`
	lineBarChartInnerSVG         = `<path fill="currentColor" d="M5 11h3v5H5v-5zm-4 3h3v2H1v-2zm12-2h3v4h-3v-4zM9 9h3v7H9V9zm7-8.93l-5.68 4.97l-5.47-1.7L0 7.1V9l5.15-4l5.53 1.72L16 2.06V.07z"/>`
	lineChartInnerSVG            = `<path fill="currentColor" d="M1 15V0H0v16h16v-1H1z"/><path fill="currentColor" d="M9 8L6 5L2 9v2l4-4l3 3l7-7V1z"/>`
	lineHInnerSVG                = `<path fill="currentColor" d="M0 7h16v1H0V7z"/>`
	lineVInnerSVG                = `<path fill="currentColor" d="M8 0h1v16H8V0z"/>`
	linesInnerSVG                = `<path fill="currentColor" d="M0 1h16v2H0V1zm0 4h16v2H0V5zm0 4h16v2H0V9zm0 4h16v2H0v-2z"/>`
	linesListInnerSVG            = `<path fill="currentColor" d="M0 1h3v2H0V1zm0 4h3v2H0V5zm0 4h3v2H0V9zm0 4h3v2H0v-2zM4 1h12v2H4V1zm0 4h12v2H4V5zm0 4h12v2H4V9zm0 4h12v2H4v-2z"/>`
	linkInnerSVG                 = `<path fill="currentColor" d="M14.9 1.1c-1.4-1.4-3.7-1.4-5.1 0L5.4 5.4C4 6.9 4 9.1 5.4 10.6c.1.1.3.2.4.3l1.5-1.5c-.1-.1-.3-.2-.4-.3c-.6-.6-.6-1.6 0-2.2l4.4-4.4c.6-.6 1.6-.6 2.2 0s.6 1.6 0 2.2L12.2 6c.4.8.5 1.7.4 2.5l2.3-2.3c1.5-1.4 1.5-3.7 0-5.1z"/><path fill="currentColor" d="M10.2 5.1L8.7 6.6s.3.2.4.3c.6.6.6 1.6 0 2.2l-4.4 4.4c-.6.6-1.6.6-2.2 0s-.6-1.6 0-2.2L3.8 10c-.4-.8-.1-1.3-.4-2.5L1.1 9.8c-1.4 1.4-1.4 3.7 0 5.1s3.7 1.4 5.1 0l4.4-4.4c1.4-1.4 1.4-3.7 0-5.1c-.2-.1-.4-.3-.4-.3z"/>`
	listInnerSVG                 = `<path fill="currentColor" d="M0 0h4v3H0V0zm0 4h4v3H0V4zm0 8h4v3H0v-3zm0-4h4v3H0V8zm5-8h11v3H5V0zm0 4h11v3H5V4zm0 8h11v3H5v-3zm0-4h11v3H5V8z"/>`
	listOlInnerSVG               = `<path fill="currentColor" d="M4 0h12v4H4V0zm0 6h12v4H4V6zm0 6h12v4H4v-4zM1 0L.1.5l.2.7l.7-.3V4h1V0zm1.2 13.9c.3-.2.5-.5.5-.8c0-.5-.4-1-1.3-1c-.5 0-1 .1-1.2.3H.1l.2.8l.1-.1c.1-.1.4-.2.7-.2s.4.1.4.3c0 .4-.5.4-.6.4H.5v.7h.4c.3 0 .6.1.6.4c0 .2-.2.4-.6.4s-.7-.2-.8-.2l-.1-.1v.9h.1c.2.2.6.3 1.1.3c1 0 1.6-.5 1.6-1.2c0-.4-.2-.8-.6-.9zM.1 6.4l.3 1s.7-.6 1.2-.3C2.7 7.9 0 9.5 0 9.5v.5h3V9H1.8c.6-.5 1.2-1.2 1-1.9C2.3 5.2.1 6.4.1 6.4z"/>`
	listSelectInnerSVG           = `<path fill="currentColor" d="M1 0h12v2H1V0zm0 8h13v2H1V8zm0 3h11v2H1v-2zm0 3h14v2H1v-2zM0 3v4h16V3H0zm11 3H1V4h10v2z"/>`
	listUlInnerSVG               = `<path fill="currentColor" d="M0 1h3v3H0V1zm0 5h3v3H0V6zm0 5h3v3H0v-3zM5 1h11v3H5V1zm0 5h11v3H5V6zm0 5h11v3H5v-3z"/>`
	locationArrowInnerSVG        = `<path fill="currentColor" d="m0 9l16-9l-9 16V9z"/>`
	locationArrowCircleInnerSVG  = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zM7 14V9H2l10-5l-5 10z"/>`
	locationArrowCircleOInnerSVG = `<path fill="currentColor" d="M1 8c0-3.9 3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7zM0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8z"/><path fill="currentColor" d="m2 9l10-5l-5 10V9z"/>`
	lockInnerSVG                 = `<path fill="currentColor" d="M12 8V4.9C12 2.7 10.4 1 8.2 1h-.3C5.8 1 4 2.7 4 4.9V8H3l.1 5S3 16 8 16s5-3 5-3V8h-1zm-3 6H8v-2c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1v3zm1-6H6V4.9C6 3.8 6.9 3 7.9 3h.3c1 0 1.8.8 1.8 1.9V8z"/>`
	magicInnerSVG                = `<path fill="currentColor" d="M0 5h3v1H0V5zm5-5h1v3H5V0zm1 11H5V8.5l1 1zm5-5H9.5l-1-1H11zM3.131 7.161l.707.707l-2.97 2.97l-.707-.707l2.97-2.97zm7-7l.707.707l-2.97 2.97l-.707-.707l2.97-2.97zM.836.199l3.465 3.465l-.707.707L.129.906L.836.199zM6.1 4.1L4 6.1l9.8 9.9l2.2-2.1l-9.9-9.8zm0 1.4L8.5 8l-.6.6l-2.5-2.5l.7-.6z"/>`
	magnetInnerSVG               = `<path fill="currentColor" d="M11 0h5v4h-5V0zm0 5v3c0 1.6-1.4 3-3 3S5 9.6 5 8V5H0v3c0 4.4 3.6 8 8 8s8-3.6 8-8V5h-5zM0 0h5v4H0V0z"/>`
	mailboxInnerSVG              = `<path fill="currentColor" d="M13 1H3l-3 9v5h16v-5l-3-9zm-2 9v2H5v-2H1.1l2.7-8h8.6l2.7 8H11z"/>`
	maleInnerSVG                 = `<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2z"/><path fill="currentColor" d="m12.79 7.32l-2.6-2.63A2.311 2.311 0 0 0 8.54 4H7.469c-.648 0-1.235.264-1.659.69l-2.6 2.63a.73.73 0 1 0 .998 1.003L6 6.53V16h1.5v-5h1v5H10V6.53l1.75 1.8a.73.73 0 1 0 1.041-1.009z"/>`
	mapMarkerInnerSVG            = `<path fill="currentColor" d="M8 0C5.2 0 3 2.2 3 5s4 11 5 11s5-8.2 5-11s-2.2-5-5-5zm0 8C6.3 8 5 6.7 5 5s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3z"/>`
	marginInnerSVG               = `<path fill="currentColor" d="M0 0h1v1H0V0zm2 0h1v1H2V0zM1 1h1v1H1V1zM0 2h1v1H0V2zm2 0h1v1H2V2zM1 3h1v1H1V3zM0 4h1v1H0V4zm1 1h1v1H1V5zM0 6h1v1H0V6zm1 1h1v1H1V7zM0 8h1v1H0V8zm1 1h1v1H1V9zm-1 1h1v1H0v-1zm1 1h1v1H1v-1zm-1 1h1v1H0v-1zm1 1h1v1H1v-1zm-1 1h1v1H0v-1zm2 0h1v1H2v-1zm-1 1h1v1H1v-1zm2 0h1v1H3v-1zm2 0h1v1H5v-1zM4 0h1v1H4V0zM3 1h1v1H3V1zm2 0h1v1H5V1zM4 14h1v1H4v-1zM6 0h1v1H6V0zm2 0h1v1H8V0zM7 1h1v1H7V1zM6 14h1v1H6v-1zm2 0h1v1H8v-1zm-1 1h1v1H7v-1zm2 0h1v1H9v-1zm2 0h1v1h-1v-1zM10 0h1v1h-1V0zM9 1h1v1H9V1zm2 0h1v1h-1V1zm-1 13h1v1h-1v-1zm2-14h1v1h-1V0zm2 0h1v1h-1V0zm-1 1h1v1h-1V1z"/><path fill="currentColor" d="M13 2h-1v1h-1V2h-1v1H9V2H8v1H7V2H6v1H5V2H4v1H3v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v-1h1v-1h-1v-1h1V9h-1V8h1V7h-1V6h1V5h-1V4h1V3h-1V2zm-1 10H4V4h8v8zm2-10h1v1h-1V2zm0 2h1v1h-1V4zm0 2h1v1h-1V6zm0 2h1v1h-1V8zm0 2h1v1h-1v-1zm0 2h1v1h-1v-1z"/><path fill="currentColor" d="M13 13h1v1h-1v-1zm-1 1h1v1h-1v-1zm2 0h1v1h-1v-1zm-1 1h1v1h-1v-1zm2 0h1v1h-1v-1zm0-14h1v1h-1V1zm0 2h1v1h-1V3zm0 2h1v1h-1V5zm0 2h1v1h-1V7zm0 2h1v1h-1V9zm0 2h1v1h-1v-1zm0 2h1v1h-1v-1z"/>`
	marginBottomInnerSVG         = `<path fill="currentColor" d="M0 0v14h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1V0H0zm15 12H1V1h14v11zM0 15h1v1H0v-1zm1-1h1v1H1v-1zm1 1h1v1H2v-1zm1-1h1v1H3v-1zm1 1h1v1H4v-1zm1-1h1v1H5v-1zm1 1h1v1H6v-1zm1-1h1v1H7v-1zm1 1h1v1H8v-1zm1-1h1v1H9v-1zm1 1h1v1h-1v-1zm1-1h1v1h-1v-1zm1 1h1v1h-1v-1zm1-1h1v1h-1v-1zm1 1h1v1h-1v-1zm1-1h1v1h-1v-1z"/>`
	marginLeftInnerSVG           = `<path fill="currentColor" d="M2 0v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1h13V0H2zm13 15H4V1h11v14zM0 0h1v1H0V0zm1 1h1v1H1V1zM0 2h1v1H0V2zm1 1h1v1H1V3zM0 4h1v1H0V4zm1 1h1v1H1V5zM0 6h1v1H0V6zm1 1h1v1H1V7zM0 8h1v1H0V8zm1 1h1v1H1V9zm-1 1h1v1H0v-1zm1 1h1v1H1v-1zm-1 1h1v1H0v-1zm1 1h1v1H1v-1zm-1 1h1v1H0v-1zm1 1h1v1H1v-1z"/>`
	marginRightInnerSVG          = `<path fill="currentColor" d="M14 2V1h-1V0H0v16h14v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1V9h-1V8h1V7h-1V6h1V5h-1V4h1V3h-1V2h1zm-2 13H1V1h11v14zm3 0h1v1h-1v-1zm-1-1h1v1h-1v-1zm1-1h1v1h-1v-1zm-1-1h1v1h-1v-1zm1-1h1v1h-1v-1zm-1-1h1v1h-1v-1zm1-1h1v1h-1V9zm-1-1h1v1h-1V8zm1-1h1v1h-1V7zm-1-1h1v1h-1V6zm1-1h1v1h-1V5zm-1-1h1v1h-1V4zm1-1h1v1h-1V3zm-1-1h1v1h-1V2zm1-1h1v1h-1V1zm-1-1h1v1h-1V0z"/>`
	marginTopInnerSVG            = `<path fill="currentColor" d="M15 2v1h-1V2h-1v1h-1V2h-1v1h-1V2H9v1H8V2H7v1H6V2H5v1H4V2H3v1H2V2H1v1H0v13h16V2h-1zm0 13H1V4h14v11zm0-15h1v1h-1V0zm-1 1h1v1h-1V1zm-1-1h1v1h-1V0zm-1 1h1v1h-1V1zm-1-1h1v1h-1V0zm-1 1h1v1h-1V1zM9 0h1v1H9V0zM8 1h1v1H8V1zM7 0h1v1H7V0zM6 1h1v1H6V1zM5 0h1v1H5V0zM4 1h1v1H4V1zM3 0h1v1H3V0zM2 1h1v1H2V1zM1 0h1v1H1V0zM0 1h1v1H0V1z"/>`
	medalInnerSVG                = `<path fill="currentColor" d="M10 12.2c-.3 0-.5-.1-.8-.2L8 11.5l-1.2.5c-.2.1-.5.2-.8.2c-.2 0-.3 0-.5-.1L5 16l3-2l3 2l-.6-3.9c-.1.1-.3.1-.4.1zm2.9-6.3c-.1-.2-.1-.5 0-.7l.6-1.2c.2-.4 0-.9-.5-1.1l-1.3-.5c-.2-.1-.4-.3-.5-.5L10.7.6c-.1-.4-.4-.6-.7-.6c-.1 0-.3 0-.4.1L8.3.7H8c-.1 0-.2 0-.3-.1L6.4.1C6.3 0 6.1 0 6 0c-.3 0-.6.2-.8.5l-.5 1.4c0 .2-.2.4-.4.5l-1.4.5c-.4.1-.6.6-.4 1.1l.6 1.3c.1.2.1.5 0 .7l-.6 1.2c-.2.4 0 .9.5 1.1l1.3.5c.2.1.4.3.5.5l.5 1.3c.1.4.4.6.7.6c.1 0 .2 0 .3-.1l1.3-.6c.1 0 .2-.1.3-.1s.2 0 .3.1l1.3.6c.1.1.2.1.3.1c.3 0 .6-.2.8-.5l.5-1.3c.1-.2.3-.4.5-.5l1.3-.5c.4-.2.7-.7.5-1.1l-.5-1.4zM8 9.6c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4z"/><path fill="currentColor" d="M11 5.6a3 3 0 1 1-6 0a3 3 0 0 1 6 0z"/>`
	megafoneInnerSVG             = `<path fill="currentColor" d="M15.5 5.4L15 5V1c0-.6-.4-1-1-1s-1 .4-1 1v.5C11 2.4 8 4 5 4H2.5C1.1 4 0 5.2 0 6.5c0 .9.5 1.7 1.2 2.1l1.1 5.9c0 .3.3.5.7.5h.2l3.6-.7c.4-.1.6-.4.5-.7c-.3-.6-.8-1.5-1.2-1.8c-.2-.1-.5-.9-.7-1.8H6v-.9c2.7.3 6 1.6 7 2.4v.5c0 .6.4 1 1 1s1-.4 1-1V8l.4-.3c.4-.3.6-.7.6-1.1v-.2c0-.4-.2-.7-.5-1zM2 5h3v1H2V5zm3.6 7.6c.1 0 .3.3.5.7l-2.8.7l-1-5h1.9c.2 1.3.6 3.2 1.4 3.6zm7.4-2.3c-1.6-.8-4.4-2-7-2.3V5c2.6-.3 5.4-1.4 7-2.3v7.6z"/>`
	mehOInnerSVG                 = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm-7 4h8v1H4v-1z"/>`
	menuInnerSVG                 = `<path fill="currentColor" d="M0 1h16v3H0V1zm0 5h16v3H0V6zm0 5h16v3H0v-3z"/>`
	microphoneInnerSVG           = `<path fill="currentColor" d="M8 10c-1.7 0-3-1.3-3-3V3c0-1.6 1.3-3 3-3c1.6 0 3 1.3 3 3v4c0 1.6-1.4 3-3 3z"/><path fill="currentColor" d="M12 5v2.5c0 1.9-1.8 3.5-3.8 3.5h-.4C5.8 11 4 9.4 4 7.5V5c-.6 0-1 .4-1 1v1.5c0 2.2 1.8 4.1 4 4.4V14c-3 0-2.5 2-2.5 2h7s.5-2-2.5-2v-2.1c2.2-.4 4-2.2 4-4.4V6c0-.6-.4-1-1-1z"/>`
	minusInnerSVG                = `<path fill="currentColor" d="M2 7h12v2H2V7z"/>`
	minusCircleInnerSVG          = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm5 9H3V7h10v2z"/>`
	minusCircleOInnerSVG         = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M3 7h10v2H3V7z"/>`
	minusSquareOInnerSVG         = `<path fill="currentColor" d="M4 7h8v2H4V7z"/><path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/>`
	mobileInnerSVG               = `<path fill="currentColor" d="M4 1v14h8V1H4zm5 13H7v-1h2v1zm2-2H5V3h6v9z"/>`
	mobileBrowserInnerSVG        = `<path fill="currentColor" d="M16 0H3v5H0v11h7v-3h9V0zM6 1h9v1H6V1zM4 1h1v1H4V1zm0 14H3v-1h1v1zm2-2H1V6h5v7zm9-1H7V5H4V3h11v9z"/>`
	mobileRetroInnerSVG          = `<path fill="currentColor" d="M11 0h-1v2H4v14h7V0zM6 14H5v-1h1v1zm0-2H5v-1h1v1zm0-2H5V9h1v1zm2 4H7v-1h1v1zm0-2H7v-1h1v1zm0-2H7V9h1v1zm2 4H9v-1h1v1zm0-2H9v-1h1v1zm0-2H9V9h1v1zm0-2H5V4h5v4z"/>`
	modalInnerSVG                = `<path fill="currentColor" d="M0 1v14h16V1H0zm15 13H1V4h14v10zm0-11h-1V2h1v1z"/>`
	modalListInnerSVG            = `<path fill="currentColor" d="M3 6h2v1H3V6zm3 0h7v1H6V6zM3 8h2v1H3V8zm3 0h7v1H6V8zm-3 2h2v1H3v-1zm3 0h7v1H6v-1z"/><path fill="currentColor" d="M0 1v14h16V1H0zm15 13H1V4h14v10zm0-11h-1V2h1v1z"/>`
	moneyInnerSVG                = `<path fill="currentColor" d="M15 4v8H1V4h14zm1-1H0v10h16V3z"/><path fill="currentColor" d="M8 5c1.7 0 3 1.3 3 3s-1.3 3-3 3h5v-1h1V6h-1V5H8zM5 8c0-1.7 1.3-3 3-3H3v1H2v4h1v1h5c-1.7 0-3-1.3-3-3z"/>`
	moneyDepositInnerSVG         = `<path fill="currentColor" d="m8 16l-2-3h1v-2h2v2h1l-2 3zm7-15v8H1V1h14zm1-1H0v10h16V0z"/><path fill="currentColor" d="M8 2a3 3 0 1 1 0 6h5V7h1V3h-1V2H8zM5 5a3 3 0 0 1 3-3H3v1H2v4h1v1h5a3 3 0 0 1-3-3z"/>`
	moneyExchangeInnerSVG        = `<path fill="currentColor" d="m16 14l-3 2v-1H8.25l2-2H13v-1l3 2zM0 2l3-2v1h4.75l-2 2H3v1L0 2zm9.74-2L0 9.74L6.26 16L16 6.26zM1.39 9.74l8.35-8.35l4.87 4.87l-8.35 8.35z"/><path fill="currentColor" d="m4.17 9.74l-.7.7l2.09 2.09l.7-.7l.74.69l2.74-2.78a2.461 2.461 0 0 1-3.48-3.48L3.48 9z"/><path fill="currentColor" d="m12.52 5.57l-2.09-2.09l-.7.7l-.73-.7l-2.74 2.78a2.461 2.461 0 0 1 3.48 3.48L12.52 7l-.7-.7z"/>`
	moneyWithdrawInnerSVG        = `<path fill="currentColor" d="m8 0l2 3H9v2H7V3H6l2-3zm7 7v8H1V7h14zm1-1H0v10h16V6z"/><path fill="currentColor" d="M8 8a3 3 0 1 1 0 6h5v-1h1V9h-1V8H8zm-3 3a3 3 0 0 1 3-3H3v1H2v4h1v1h5a3 3 0 0 1-3-3z"/>`
	moonInnerSVG                 = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm0 15c-3.9 0-7-3.1-7-7c0-2.4 1.2-4.6 3.2-5.9C4.1 2.7 4 3.4 4 4c0 4.9 4 8.9 8.9 9c-1.3 1.3-3 2-4.9 2z"/>`
	moonOInnerSVG                = `<path fill="currentColor" d="M13.2 11.9c-4.5 0-8.1-3.6-8.1-8.1c0-1.4.3-2.7.9-3.8c-3.4.9-6 4.1-6 7.9C0 12.4 3.6 16 8.1 16c3.1 0 5.8-1.8 7.2-4.4c-.6.2-1.3.3-2.1.3zM8.1 15C4.2 15 1 11.8 1 7.9c0-2.5 1.3-4.7 3.3-6c-.2.6-.2 1.2-.2 1.9c0 5 4.1 9.1 9.1 9.2c-1.4 1.2-3.2 2-5.1 2z"/>`
	morningInnerSVG              = `<path fill="currentColor" d="m14 10l-1.58-1.18L13.2 7l-2-.23L11 4.8l-1.82.78L8 4L6.82 5.58L5 4.8l-.23 2L2.8 7l.78 1.82L2 10H0v1h16v-1h-2zM4 10a4.143 4.143 0 0 1 3.993-4A4.143 4.143 0 0 1 12 9.993L4 10z"/>`
	movieInnerSVG                = `<path fill="currentColor" d="M12 7V4H0v9h12v-3l4 2V5l-4 2zm-3 4H2V6h7v5z"/><path fill="currentColor" d="M5 8.4a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`
	musicInnerSVG                = `<path fill="currentColor" d="M4 3v9.4c-.4-.2-.9-.4-1.5-.4c-1.4 0-2.5.9-2.5 2s1.1 2 2.5 2S5 15.1 5 14V6.7l7-2.3v5.1c-.4-.3-.9-.5-1.5-.5C9.1 9 8 9.9 8 11s1.1 2 2.5 2s2.5-.9 2.5-2V0L4 3z"/>`
	muteInnerSVG                 = `<path fill="currentColor" d="M15.2 0L11 4.2V3c0-1.7-1.3-3-3-3S5 1.3 5 3v4c0 .9.4 1.7 1 2.2l-.8.8C4.5 9.4 4 8.5 4 7.5V5c-.6 0-1 .4-1 1v1.5c0 1.3.6 2.4 1.5 3.2L0 15.3v.7h.7L16 .6V0h-.8zm-2.7 5.1l-.5.5v1.9c0 1.9-1.8 3.5-3.8 3.5h-.4c-.3 0-.6-.1-.9-.1l-.9.7c.3.1.6.2 1 .3V14c-3 0-2.5 2-2.5 2h7s.5-2-2.5-2v-2.1c2.2-.4 4-2.2 4-4.4V6c0-.4-.2-.7-.5-.9z"/><path fill="currentColor" d="M11 7v-.4L7.7 10H8c1.7 0 3-1.4 3-3z"/>`
	nativeButtonInnerSVG         = `<path fill="currentColor" d="M15 12H1c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h14c.6 0 1 .4 1 1v6c0 .6-.4 1-1 1z"/>`
	newspaperInnerSVG            = `<path fill="currentColor" d="M2 4h11v4H2V4zm0-2h11v1H2V2zm6 11h3v1H8v-1zm0-2h5v1H8v-1zm0-2h5v1H8V9zm-6 4h5v1H2v-1zm0-2h5v1H2v-1zm0-2h5v1H2V9z"/><path fill="currentColor" d="M15 2V0H0v14.5A1.5 1.5 0 0 0 1.5 16h13a1.5 1.5 0 0 0 1.5-1.5V2h-1zM1.5 15a.5.5 0 0 1-.5-.5V1h13v12.5c0 1.5 1 1.5 1 1.5H1.5z"/>`
	notebookInnerSVG             = `<path fill="currentColor" d="M2 0v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0 0 .96H2v1h-.52a.48.48 0 0 0 0 .96H2v1h-.52a.48.48 0 0 0 0 .96H2v2h12V0H2zm1.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM12 6H6V3h6v3z"/>`
	nurseInnerSVG                = `<path fill="currentColor" d="M15.2 12a4.076 4.076 0 0 0-3.931-2.37L8 13.53l-3.28-3.9a4.16 4.16 0 0 0-3.909 2.345a9.072 9.072 0 0 0-.808 2.988L2 15v1h12v-1h2a9.199 9.199 0 0 0-.824-3.057z"/><path fill="currentColor" d="M6.57 8.73a.82.82 0 0 1-.685.729L8 12l2.12-2.52a.823.823 0 0 1-.69-.727c0-.613.8-.413 1.43-1.453C10.86 7.27 13.74 0 8 0S5.14 7.27 5.14 7.27c.63 1.05 1.44.85 1.43 1.46z"/>`
	officeInnerSVG               = `<path fill="currentColor" d="M14 15V4h-3V1H2v14H0v1h7v-3h2v3h7v-1h-2zm-8-4H4V9h2v2zm0-3H4V6h2v2zm0-3H4V3h2v2zm3 6H7V9h2v2zm0-3H7V6h2v2zm0-3H7V3h2v2zm4 6h-2V9h2v2zm0-3h-2V6h2v2z"/>`
	openBookInnerSVG             = `<path fill="currentColor" d="M15 4.7V4a6.804 6.804 0 0 0-4.484-1.999a2.844 2.844 0 0 0-2.513.995a3.02 3.02 0 0 0-2.515-.995A6.804 6.804 0 0 0 1 4v.7L0 5v10l6.7-1.4l.3.4h2l.3-.4L16 15V5zm-9.52 6.61a8.206 8.206 0 0 0-3.526.902L2 4.42A5.22 5.22 0 0 1 5.369 3a4.553 4.553 0 0 1 2.159.701l-.019 7.869a6.548 6.548 0 0 0-2.039-.259zm8.52.88a8.122 8.122 0 0 0-3.468-.88l-.161-.002c-.66 0-1.297.096-1.899.274l.047-7.902a4.484 4.484 0 0 1 2.096-.679a5.216 5.216 0 0 1 3.386 1.422l-.003 7.768z"/>`
	optionInnerSVG               = `<path fill="currentColor" d="M4 11a1 1 0 0 0 2 0v-1H5a1 1 0 0 0-1 1z"/><path fill="currentColor" d="M0 0v16h16V0H0zm11 9a2 2 0 1 1-2 2v-1H7v1a2 2 0 1 1-2-2h1V7H5a2 2 0 1 1 2-2v1h2V5a2 2 0 1 1 2 2h-1v2h1z"/><path fill="currentColor" d="M12 5a1 1 0 0 0-2 0v1h1a1 1 0 0 0 1-1zM5 4a1 1 0 0 0 0 2h1V5a1 1 0 0 0-1-1zm2 3h2v2H7V7zm3 4a1 1 0 1 0 1-1h-1v1z"/>`
	optionAInnerSVG              = `<path fill="currentColor" d="M12.5 10H11V6h1.5A2.5 2.5 0 1 0 10 3.5V5H6V3.5A2.5 2.5 0 1 0 3.5 6H5v4H3.5A2.5 2.5 0 1 0 6 12.5V11h4v1.5a2.5 2.5 0 1 0 2.5-2.5zM11 3.5A1.5 1.5 0 1 1 12.5 5H11V3.5zm-6 9A1.5 1.5 0 1 1 3.5 11H5v1.5zM5 5H3.5A1.5 1.5 0 1 1 5 3.5V5zm5 5H6V6h4v4zm2.5 4a1.5 1.5 0 0 1-1.5-1.5V11h1.5a1.5 1.5 0 0 1 0 3z"/>`
	optionsInnerSVG              = `<path fill="currentColor" d="M5 3.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 5 3.5z"/><path fill="currentColor" d="M3.5 0C1.6 0 0 1.6 0 3.5S1.6 7 3.5 7S7 5.4 7 3.5S5.4 0 3.5 0zm0 6C2.1 6 1 4.9 1 3.5S2.1 1 3.5 1S6 2.1 6 3.5S4.9 6 3.5 6zm0 2C1.6 8 0 9.6 0 11.5S1.6 15 3.5 15S7 13.4 7 11.5S5.4 8 3.5 8zm0 6C2.1 14 1 12.9 1 11.5S2.1 9 3.5 9S6 10.1 6 11.5S4.9 14 3.5 14zM8 2h8v3H8V2zm0 8h8v3H8v-3z"/>`
	orientationInnerSVG          = `<path fill="currentColor" d="M11 2.1c2 0 3 1.3 3 2.9h-1l1.5 2L16 5h-1c0-2.2-2-3.9-4-3.9V0L9 1.5L11 3v-.9z"/><path fill="currentColor" d="M9 9h6v6H8V0H0v16h16V8H9v1zM7 8H6v1h1v6H1V1h6v7z"/><path fill="currentColor" d="M2 8h1v1H2V8zm2 0h1v1H4V8z"/>`
	outInnerSVG                  = `<path fill="currentColor" d="M3.5 8c.3 0 .5.2.5.5v2c0 .3-.2.5-.5.5s-.5-.2-.5-.5v-2c0-.3.2-.5.5-.5zm0-1C2.7 7 2 7.7 2 8.5v2c0 .8.7 1.5 1.5 1.5S5 11.3 5 10.5v-2C5 7.7 4.3 7 3.5 7zM8 7v3.5c0 .3-.2.5-.5.5s-.5-.2-.5-.5V7H6v3.5c0 .8.7 1.5 1.5 1.5S9 11.3 9 10.5V7H8zm5 0h-3v1h1v4h1V8h1z"/><path fill="currentColor" d="M15 6V5h-2.4L8.9 2c.1-.2.1-.3.1-.5C9 .7 8.3 0 7.5 0S6 .7 6 1.5c0 .2 0 .3.1.5L2.4 5H0v9h1v1h15V6h-1zM6.7 2.8c.3.1.5.2.8.2s.5-.1.8-.2L11 5H4l2.7-2.2zM14 13H1V6h13v7z"/>`
	outboxInnerSVG               = `<path fill="currentColor" d="M6 5v6h4V5h2L8 0L4 5z"/><path fill="currentColor" d="M13 2h-2l.9 1h.4l2.6 8H11v2H5v-2H1.1l2.6-8h.4L5 2H3l-3 9v5h16v-5z"/>`
	packageInnerSVG              = `<path fill="currentColor" d="M8 0L0 2v10l8 4l8-4V2L8 0zm0 1l2.1.5l-5.9 1.9l-2.3-.8L8 1zm0 13.9l-7-3.5V3.3l3 1v3.4L5 8V4.7l3 1v9.2zm.5-10.1l-2.7-.9L12 2l2.4.6l-5.9 2.2z"/>`
	paddingInnerSVG              = `<path fill="currentColor" d="M0 0v16h16V0H0zm15 3h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v-1h1v-1H1v-1h1v-1H1V9h1V8H1V7h1V6H1V5h1V4H1V3h1V2H1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1v1z"/><path fill="currentColor" d="M3 2h1v1H3V2zm1 1h1v1H4V3zm2 0h1v1H6V3zM5 2h1v1H5V2zm2 0h1v1H7V2zm2 0h1v1H9V2zM8 3h1v1H8V3zm2 0h1v1h-1V3zm2 0h1v1h-1V3zm-1-1h1v1h-1V2zm2 0h1v1h-1V2zm-1 3h1v1h-1V5zm1-1h1v1h-1V4zm-1 3h1v1h-1V7zm1-1h1v1h-1V6zm-1 3h1v1h-1V9zm1-1h1v1h-1V8zm-1 3h1v1h-1v-1zm1-1h1v1h-1v-1zm-1 3h1v1h-1v-1zm1-1h1v1h-1v-1zM2 3h1v1H2V3zm1 1h1v1H3V4zM2 5h1v1H2V5zm1 1h1v1H3V6zM2 7h1v1H2V7zm1 1h1v1H3V8zM2 9h1v1H2V9zm1 1h1v1H3v-1zm-1 1h1v1H2v-1zm0 2h1v1H2v-1zm1-1h1v1H3v-1zm1-1h1v1H4v-1zm0 2h1v1H4v-1zm1-1h1v1H5v-1zm1 1h1v1H6v-1zm1-1h1v1H7v-1zm2 0h1v1H9v-1zm-1 1h1v1H8v-1zm3-1h1v1h-1v-1zm-1 1h1v1h-1v-1z"/>`
	paddingBottomInnerSVG        = `<path fill="currentColor" d="M16 16V0H0v16h16zM1 13h1v-1H1V1h14v12h-1v1h1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v-1z"/><path fill="currentColor" d="M12 13h1v1h-1v-1zm1-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm1 1h1v1h-1v-1zm-2 0h1v1H8v-1zm-2 0h1v1H6v-1zm1-1h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm1 1h1v1H4v-1zm-2 0h1v1H2v-1z"/>`
	paddingLeftInnerSVG          = `<path fill="currentColor" d="M0 16h16V0H0v16zM3 1v1h1V1h11v14H3v-1H2v1H1v-1h1v-1H1v-1h1v-1H1v-1h1V9H1V8h1V7H1V6h1V5H1V4h1V3H1V2h1V1h1z"/><path fill="currentColor" d="M2 12h1v1H2v-1zm1 1h1v1H3v-1zm0-2h1v1H3v-1zm0-2h1v1H3V9zm-1 1h1v1H2v-1zm0-2h1v1H2V8zm0-2h1v1H2V6zm1 1h1v1H3V7zm0-2h1v1H3V5zm0-2h1v1H3V3zM2 4h1v1H2V4zm0-2h1v1H2V2z"/>`
	paddingRightInnerSVG         = `<path fill="currentColor" d="M16 0H0v16h16V0zm-3 15v-1h-1v1H1V1h12v1h1V1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h-1z"/><path fill="currentColor" d="M13 3h1v1h-1V3zm-1-1h1v1h-1V2zm0 2h1v1h-1V4zm0 2h1v1h-1V6zm1-1h1v1h-1V5zm0 2h1v1h-1V7zm0 2h1v1h-1V9zm-1-1h1v1h-1V8zm0 2h1v1h-1v-1zm0 2h1v1h-1v-1zm1-1h1v1h-1v-1zm0 2h1v1h-1v-1z"/>`
	paddingTopInnerSVG           = `<path fill="currentColor" d="M0 0v16h16V0H0zm15 3h-1v1h1v11H1V3h1V2H1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1v1z"/><path fill="currentColor" d="M3 2h1v1H3V2zM2 3h1v1H2V3zm2 0h1v1H4V3zm2 0h1v1H6V3zM5 2h1v1H5V2zm2 0h1v1H7V2zm2 0h1v1H9V2zM8 3h1v1H8V3zm2 0h1v1h-1V3zm2 0h1v1h-1V3zm-1-1h1v1h-1V2zm2 0h1v1h-1V2z"/>`
	paintRollInnerSVG            = `<path fill="currentColor" d="M16 6.9V2h-2V0H1v1H0v3h1v1h13V3h1v3.1l-8 1V9H6v.9s.5 0 .5.9s-.5.6-.5 1.5v2.8s0 .9 1.5.9s1.5-.9 1.5-.9v-2.8c0-.9-.5-.7-.5-1.5s.5-.9.5-.9V9H8V7.9l8-1z"/>`
	paintbrushInnerSVG           = `<path fill="currentColor" d="m5.6 11.6l-1.2-1.2c-.8-.2-2-.1-2.7 1c-.8 1.1-.3 2.8-1.7 4.6c0 0 3.5 0 4.8-1.3c1.2-1.2 1.2-2.2 1-3l-.2-.1zm.2-3.5c-.2.3-.5.7-.7 1c0 .2-.1.3-.2.4L6.4 11c.1-.1.3-.2.4-.3c.3-.2.7-.4 1-.7c.4 0 .6-.2.8-.4L6.4 7.4c-.2.2-.4.4-.6.7zm10-7.9c-.3-.3-.7-.3-1-.1c0 0-3 2.5-5.9 5.1c-.4.4-.7.7-1.1 1c-.2.2-.4.4-.6.5l2.1 2.1c.2-.2.4-.4.5-.7c.3-.4.6-.7.9-1.1c2.5-3 5.1-5.9 5.1-5.9c.3-.2.3-.6 0-.9z"/>`
	paleteInnerSVG               = `<path fill="currentColor" d="M8.25 0C1.87 0-.86 7.38.24 9.92c.82 1.89 2.62.08 3.34 1c1.88 2.46-2.11 3.81.09 4.68C6.26 16.66 16 16 16 7.07C16 4.38 14.66 0 8.25 0zM4.47 9A1.5 1.5 0 1 1 6 7.5A1.5 1.5 0 0 1 4.5 9h-.032zM6 3.5A1.5 1.5 0 1 1 7.5 5h-.032A1.5 1.5 0 0 1 6 3.5zM8.47 14A1.5 1.5 0 1 1 10 12.5A1.5 1.5 0 0 1 8.5 14h-.032zm4-3A1.5 1.5 0 1 1 14 9.5a1.5 1.5 0 0 1-1.5 1.5h-.032zm0-5A1.5 1.5 0 1 1 14 4.5A1.5 1.5 0 0 1 12.5 6h-.032z"/>`
	panelInnerSVG                = `<path fill="currentColor" d="M0 0v16h16V0H0zm13 15H1V3h12v12zm2 0h-1v-1h1v1zm0-2h-1V5h1v8zm0-9h-1V3h1v1z"/>`
	paperclipInnerSVG            = `<path fill="currentColor" d="M2.7 15.3c-.7 0-1.4-.3-1.9-.8c-.9-.9-1.2-2.5 0-3.7l8.9-8.9c1.4-1.4 3.8-1.4 5.2 0s1.4 3.8 0 5.2l-7.4 7.4c-.2.2-.5.2-.7 0s-.2-.5 0-.7l7.4-7.4c1-1 1-2.7 0-3.7s-2.7-1-3.7 0l-8.9 8.9c-.8.8-.6 1.7 0 2.2c.6.6 1.5.8 2.2 0l8.9-8.9c.2-.2.2-.5 0-.7s-.5-.2-.7 0l-7.4 7.4c-.2.2-.5.2-.7 0s-.2-.5 0-.7l7.4-7.4c.6-.6 1.6-.6 2.2 0s.6 1.6 0 2.2l-8.9 8.9c-.6.4-1.3.7-1.9.7z"/>`
	paperplaneInnerSVG           = `<path fill="currentColor" d="m0 8l4.9 1.4H5v-.1L12.1 4L11 5.2l-6.2 6.6L5 15l2.9-3.2L10 16l6-16z"/>`
	paperplaneOInnerSVG          = `<path fill="currentColor" d="M16 0L0 8l4.7 1.6L5 15l2.5-2.8L10 16l6-16zM7.5 10.4l4.3-5.9l-6.2 4.3l-3-1L14.2 2L9.7 13.8l-2.2-3.4z"/>`
	paragraphInnerSVG            = `<path fill="currentColor" d="M5.5 0C3 0 1 2 1 4.5S3 9 5.5 9H8v7h2V2h1v14h2V2h2V0H5.5z"/>`
	passwordInnerSVG             = `<path fill="currentColor" d="M16 5c0-.6-.4-1-1-1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V5zm-1 6H1V5h14v6z"/><path fill="currentColor" d="M6 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`
	pasteInnerSVG                = `<path fill="currentColor" d="M13 4h-3V0H0v14h6v2h10V7l-3-3zM3 1h4v1H3V1zm12 14H7V5h5v3h3v7zm-2-8V5l2 2h-2z"/>`
	pauseInnerSVG                = `<path fill="currentColor" d="M0 1h7v14H0V1zm9 0h7v14H9V1z"/>`
	pencilInnerSVG               = `<path fill="currentColor" d="M1 11.9L0 16l4.1-1l9.2-9.2l-3.1-3.1L1 11.9zm.5 3.1l-.4-.5l.4-2l2 2l-2 .5zm9.4-10.6l-8.1 8l-.6-.6l8.1-8l.6.6zM15.3.7C14.2-.4 12.7.2 12.7.2l-1.5 1.5l3.1 3.1l1.5-1.5c0-.1.6-1.5-.5-2.6zm-1.9.9l-.5-.5C13.5.5 14 1 14 1l-.6.6z"/>`
	phoneInnerSVG                = `<path fill="currentColor" d="M12.2 10c-1.1-.1-1.7 1.4-2.5 1.8C8.4 12.5 6 10 6 10S3.5 7.6 4.1 6.3c.5-.8 2-1.4 1.9-2.5c-.1-1-2.3-4.6-3.4-3.6C.2 2.4 0 3.3 0 5.1c-.1 3.1 3.9 7 3.9 7c.4.4 3.9 4 7 3.9c1.8 0 2.7-.2 4.9-2.6c1-1.1-2.5-3.3-3.6-3.4z"/>`
	phoneLandlineInnerSVG        = `<path fill="currentColor" d="m15.88 3.86l-.61-1.31a1.214 1.214 0 0 0-.792-.658c-1.938-.528-4.161-.851-6.453-.891a27.46 27.46 0 0 0-6.687.934c-.165.048-.453.29-.605.609L.12 3.861a1.195 1.195 0 0 0-.12.52v.87l-.001.041c0 .392.318.71.71.71l.033-.001H3.26a.76.76 0 0 0 .742-.76L4 5.188v-.85a.65.65 0 0 1 .298-.546a6.913 6.913 0 0 1 3.724-.791a6.965 6.965 0 0 1 3.717.788c.143.099.262.3.262.529v.872a.76.76 0 0 0 .739.81h2.521l.031.001a.71.71 0 0 0 .71-.71l-.001-.043V4.38a1.208 1.208 0 0 0-.123-.527z"/><path fill="currentColor" d="M12 8.3a4.73 4.73 0 0 1-1.001-2.92L11 5.296V5h-1v1H6V5H5v.33l.001.08a4.74 4.74 0 0 1-1.009 2.93L1 12v3h14v-3zM8 13a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/><path fill="currentColor" d="M10 10a2 2 0 1 1-3.999.001A2 2 0 0 1 10 10z"/>`
	pictureInnerSVG              = `<path fill="currentColor" d="M16 14H0V2h16v12zM1 13h14V3H1v10z"/><path fill="currentColor" d="M2 10v2h12v-1s.2-1.7-2-2c-1.9-.3-2.2.6-3.8.6C7.1 9.6 7.3 8 5 8c-1.7 0-3 2-3 2zm11-4a2 2 0 1 1-3.999.001A2 2 0 0 1 13 6z"/>`
	pieBarChartInnerSVG          = `<path fill="currentColor" d="M5 11h3v5H5v-5zm-4 3h3v2H1v-2zm12-2h3v4h-3v-4zM9 9h3v7H9V9zM5 0a5 5 0 1 0 .001 10.001A5 5 0 0 0 5 0zm0 9a4 4 0 0 1 0-8v4h4a4 4 0 0 1-4 4z"/>`
	pieChartInnerSVG             = `<path fill="currentColor" d="M9 1c3.2.2 5.7 2.8 6 6H9V1zm-.5-1H8v8h8v-.5C16 3.4 12.6 0 8.5 0z"/><path fill="currentColor" d="M7 9V1c-3.9.3-7 3.5-7 7.5C0 12.6 3.4 16 7.5 16c4 0 7.2-3.1 7.5-7H7z"/>`
	piggyBankInnerSVG            = `<path fill="currentColor" d="M15.93 5.75a1.25 1.25 0 0 0-.3-.51a.833.833 0 0 0-.394-.238c.074.117.141.252.191.396c.056.147.092.304.103.467a1.784 1.784 0 0 1 0 .424a1.005 1.005 0 0 0-.142-.292a1.193 1.193 0 0 0-.48-.383a.938.938 0 0 0-1.195.382c-.05.082-.09.171-.12.266c-1.182-1.968-3.309-3.271-5.741-3.271c-.124 0-.247.003-.369.01a8.217 8.217 0 0 0-2.231.313C5.19 2.88 4.62 2 2 2l.8 2.51a5.207 5.207 0 0 0-1.247 1.465L0 6s-.17 4 1 4h.54a5.276 5.276 0 0 0 1.445 1.589L3 13.999h1.08c1.31 0 1.92 0 1.92-.75v-.39a8.256 8.256 0 0 0 3.051-.008L9 13.249c0 .75.62.75 1.94.75H12v-2.39a4.79 4.79 0 0 0 1.903-2.717c.057-.027.114-.024.172-.024s.115-.003.172-.01c.251-.046.48-.144.679-.283a1.65 1.65 0 0 0 .591-.765c.028-.093.049-.191.063-.292l.001-.01c.221-.262.372-.59.419-.951a1.776 1.776 0 0 0-.072-.822zm-12.42 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zM5.88 5a.502.502 0 0 1-.599-.247a.39.39 0 0 1 .296-.503a8.024 8.024 0 0 1 2.009-.22l.101-.001c.672 0 1.324.08 1.949.232c.126.024.262.182.262.372a.385.385 0 0 1-.019.119a.483.483 0 0 1-.346.247H9.38a7.198 7.198 0 0 0-1.706-.2h-.089c-.605 0-1.193.073-1.756.211zm8.7 2.93a1.16 1.16 0 0 1-.453.199L14 8.14v-.45c.165.125.374.2.6.2h.021zm.08-.68a.44.44 0 0 1-.459-.248a.607.607 0 0 1 .001-.566a.332.332 0 0 1 .43-.096a.48.48 0 0 1 .308.448v.001a1.457 1.457 0 0 1-.001.418a1.26 1.26 0 0 1-.282.022z"/>`
	piggyBankCoinInnerSVG        = `<path fill="currentColor" d="M15.93 7.75a1.25 1.25 0 0 0-.3-.51a.833.833 0 0 0-.394-.238c.074.117.141.252.191.396c.056.147.092.304.103.467a1.784 1.784 0 0 1 0 .424a1.005 1.005 0 0 0-.142-.292a1.193 1.193 0 0 0-.48-.383a.938.938 0 0 0-1.195.382c-.05.082-.09.171-.12.266c-1.182-1.968-3.309-3.271-5.741-3.271c-.124 0-.247.003-.369.01a8.217 8.217 0 0 0-2.231.313C5.19 4.88 4.62 4 2 4l.8 2.51a5.207 5.207 0 0 0-1.247 1.465L0 8s-.17 4 1 4h.54a5.276 5.276 0 0 0 1.445 1.589L3 16h1.08C5.39 16 6 16 6 15.25v-.39a8.256 8.256 0 0 0 3.051-.008L9 15.25c0 .75.62.75 1.94.75H12v-2.39a4.79 4.79 0 0 0 1.903-2.717c.057-.027.114-.024.172-.024s.115-.003.172-.01c.251-.046.48-.144.679-.283a1.65 1.65 0 0 0 .591-.765c.028-.093.049-.191.063-.292l.001-.01c.221-.262.372-.59.419-.951a1.776 1.776 0 0 0-.072-.822zm-12.42 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zM5.88 7a.502.502 0 0 1-.599-.247a.39.39 0 0 1 .296-.503a8.024 8.024 0 0 1 2.009-.22l.101-.001c.672 0 1.324.08 1.949.232c.126.024.262.182.262.372a.385.385 0 0 1-.019.119a.483.483 0 0 1-.346.247H9.38a7.198 7.198 0 0 0-1.706-.2h-.089c-.605 0-1.193.073-1.756.211zm8.7 2.93a1.16 1.16 0 0 1-.453.199L14 10.13v-.44c.165.125.374.2.6.2h.021zm.08-.68a.44.44 0 0 1-.459-.248a.607.607 0 0 1 .001-.566a.332.332 0 0 1 .43-.096a.48.48 0 0 1 .308.448v.001a1.457 1.457 0 0 1-.001.418a1.26 1.26 0 0 1-.282.022zM8 3H7v-.17h.25V1.74H7l.55-.55h.2v1.64H8V3z"/><path fill="currentColor" d="M7.5.75a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 7.5.75zm0-.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5z"/>`
	pillInnerSVG                 = `<path fill="currentColor" d="m14.8 1.4l-.2-.2C13.9.4 12.8 0 11.8 0S9.7.4 8.9 1.2L1.2 8.9c-1.6 1.6-1.6 4.1 0 5.7l.2.2c.7.8 1.8 1.2 2.8 1.2s2.1-.4 2.9-1.2L14.9 7c1.5-1.5 1.5-4.1-.1-5.6zm-.7 5l-3.9 3.9l-3.5-3.6l-3.8 3.8c-1.1 1.1-1.1 2.5-1 3.5c-1.2-1.2-1.2-3.1 0-4.3l7.8-7.8c.5-.6 1.3-.9 2.1-.9s1.6.3 2.2.9l.2.2c.5.5.8 1.3.8 2.1s-.3 1.6-.9 2.2z"/>`
	pillsInnerSVG                = `<path fill="currentColor" d="m3.5 8l6.3-6.3c.4-.4 1-.7 1.7-.7s1.3.3 1.8.7c1 1 1 2.6 0 3.5L10.5 8h1.4l2-2c1.4-1.4 1.4-3.6 0-4.9c-.7-.7-1.6-1-2.5-1S9.7.3 9 1L2.7 7.4c-.3.2-.5.5-.7.9c.5-.2 1-.3 1.5-.3z"/><path fill="currentColor" d="M7.3 5.6L4.9 8h4.7zM12.5 9h-9C1.6 9 0 10.6 0 12.5S1.6 16 3.5 16h9c1.9 0 3.5-1.6 3.5-3.5S14.4 9 12.5 9zm0 6H8v-4H3.5c-1.1 0-2 .6-2.5 1.2C1.2 11 2.2 10 3.5 10h9c1.4 0 2.5 1.1 2.5 2.5S13.9 15 12.5 15z"/>`
	pinInnerSVG                  = `<path fill="currentColor" d="M11 6.5V1h1V0H4v1h1v5.5S3 8 3 10c0 .5 1.9.7 4 .7v2.2c0 .7.2 1.4.5 2.1l.5 1l.5-1c.3-.6.5-1.3.5-2.1v-2.2c2.1 0 4-.3 4-.7c0-2-2-3.5-2-3.5zm-4 .1s-.5.3-1.6 1.4c-1 1-1.5 1.9-1.5 1.9s.1-1 .8-1.9C5.6 6.9 6 6.6 6 6.6V1h1v5.6z"/>`
	pinPostInnerSVG              = `<path fill="currentColor" d="M15 4V3H9c0-1.69 1-2 1-2V0H5v1s1 .31 1 2H0v12h2v1h14V4h-1zm-1 10H1V4h4v1h2v2h1V5h2V4h4v10z"/>`
	playInnerSVG                 = `<path fill="currentColor" d="M2 1v14l12-7z"/>`
	playCircleInnerSVG           = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zM6 12V4l6 4l-6 4z"/>`
	playCircleOInnerSVG          = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M6 4v8l6-4z"/>`
	plugInnerSVG                 = `<path fill="currentColor" d="M14.7 3.1c-.4-.4-1-.4-1.4 0l-2.8 2.8L9 4.5l2.8-2.8c.4-.4.4-1 0-1.4s-1-.4-1.4 0L7.6 3.1L6.2 1.7L4.8 3.1l.7.7l-1.4 1.4c-1.4 1.4-1.5 3.5-.5 5.1C1.9 11.8 1 14.1 1 16h2c0-1.3.4-3.2 2.1-4.4c1.5.8 3.4.5 4.6-.7l1.4-1.4l.7.7l1.4-1.4l-1.4-1.4l2.8-2.8c.5-.5.5-1.1.1-1.5z"/>`
	plusInnerSVG                 = `<path fill="currentColor" d="M14 7H9V2H7v5H2v2h5v5h2V9h5V7z"/>`
	plusCircleInnerSVG           = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm5 9H9v4H7V9H3V7h4V3h2v4h4v2z"/>`
	plusCircleOInnerSVG          = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M13 7H9V3H7v4H3v2h4v4h2V9h4z"/>`
	plusMinusInnerSVG            = `<path fill="currentColor" d="M10 7h6v2h-6V7zM4 5H2v2H0v2h2v2h2V9h2V7H4zm2-3l3 12h1L7 2z"/>`
	plusSquareOInnerSVG          = `<path fill="currentColor" d="M12 7H9V4H7v3H4v2h3v3h2V9h3z"/><path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/>`
	pointerInnerSVG              = `<path fill="currentColor" d="M12.6 5H12c0-.2-.2-.6-.4-.8s-.6-.4-1.1-.4c-.2 0-.4 0-.6.1c-.1-.2-.2-.3-.3-.5c-.2-.2-.5-.4-1.1-.4c-.2 0-.4 0-.5.1V1.4C8 .8 7.6 0 6.6 0c-.4 0-.8.2-1.1.4C5 1 5 1.8 5 1.8v4.3c-.6.1-1.1.3-1.4.6C3 7.4 3 8.3 3 9.5v.7c0 1.4.7 2.1 1.4 2.8l.3.4C6 14.6 7.1 15 9.8 15c2.9 0 4.2-1.6 4.2-5.1V7.4c0-.7-.2-2.1-1.4-2.4zm.4 2.4V10c0 3.4-1.3 4.1-3.2 4.1c-2.4 0-3.3-.3-4.3-1.3l-.4-.4c-.7-.8-1.1-1.2-1.1-2.2v-.7c0-1 0-1.7.3-2.1c.1-.1.4-.2.7-.2v.5l-.3 1.5c0 .1 0 .1.1.2s.2 0 .2 0l1-1.2V1.8c0-.1 0-.5.2-.7c.1 0 .2-.1.4-.1c.3 0 .4.3.4.4v4.3c0 .3.2.6.5.6S8 6 8 5.8V4.5c0-.1.1-.5.5-.5c.3 0 .5.1.5.4v1.3c0 .3.2.6.5.6s.5-.3.5-.5v-.7c0-.1.1-.3.5-.3c.2 0 .3.1.3.1c.2.1.2.4.2.4v.8c0 .3.2.5.4.5c.3 0 .5-.1.5-.4c0-.1.1-.2.2-.3h.2c.6.2.7 1.2.7 1.5c0-.1 0-.1 0 0z"/>`
	powerOffInnerSVG             = `<path fill="currentColor" d="M10 2.3v3.3c1.2.7 2 2 2 3.4c0 2.2-1.8 4-4 4s-4-1.8-4-4c0-1.5.8-2.8 2-3.4V2.3C3.1 3.2 1 5.8 1 9c0 3.9 3.1 7 7 7s7-3.1 7-7c0-3.2-2.1-5.8-5-6.7z"/><path fill="currentColor" d="M7 1h2v7H7V1z"/>`
	presentationInnerSVG         = `<path fill="currentColor" d="M16 1H9V0H7v1H0v11h5l-2 4h2.2l2-4h1.5l2 4H13l-2-4h5V1zm-1 10H1V2h14v9z"/><path fill="currentColor" d="M6 4v5l4-2.5z"/>`
	printInnerSVG                = `<path fill="currentColor" d="M0 10v4h2v2h12v-2h2v-4H0zm13 5H3v-3h10v3zm-1-9V2L9.3 0H4v6H0v3h16V6h-4zM9 1l1.3 1H9V1zm2 6H5V1h3v2h3v4zm4 1h-1V7h1v1z"/>`
	progressbarInnerSVG          = `<path fill="currentColor" d="M0 5v6h16V5H0zm15 5H1V6h14v4z"/><path fill="currentColor" d="M2 7h7v2H2V7z"/>`
	puzzlePieceInnerSVG          = `<path fill="currentColor" d="M14.9.9c-1.1-1-2.5-1.3-3.1-.4c-.7 1.1.5 1.7-.3 2.5c-.5.6-2-.8-2-.8l-.8-.8l-1.4 1.4c-.6.7-2.1 1.5-2.6 1.1c-.7-.6.1-1.8-.5-2.6c-.7-1-2.1-.8-3 .3c-1 1.1-1.4 2.4-.5 3c1.1.7 1.9-.3 2.7.5c.4.4-.2 1.7-.5 2.1L.6 9.5L7.1 16l1.7-1.7c.7-.7 1.5-2 1.1-2.4c-.6-.7-1.7.1-2.5-.4c-1-.7-.8-2 .3-3s2.5-1.3 3.1-.4c.7 1.1-.4 1.8.4 2.6c.4.4 1.6-.2 2-.6L15.3 8l-1.1-1.1c-.6-.6-1.9-2-1.4-2.5c.6-.7 1.7.2 2.5-.4c.9-.8.6-2.1-.4-3.1z"/>`
	pyramidChartInnerSVG         = `<path fill="currentColor" d="M10.29 5L8 1L5.71 5h4.58zm-8 6L0 15h16l-2.29-4H2.29zm10.85-1l-2.28-4H5.14l-2.28 4h10.28z"/>`
	qrcodeInnerSVG               = `<path fill="currentColor" d="M6 0H0v6h6V0zM5 5H1V1h4v4z"/><path fill="currentColor" d="M2 2h2v2H2V2zM0 16h6v-6H0v6zm1-5h4v4H1v-4z"/><path fill="currentColor" d="M2 12h2v2H2v-2zm8-12v6h6V0h-6zm5 5h-4V1h4v4z"/><path fill="currentColor" d="M12 2h2v2h-2V2zM2 7H0v2h3V8H2zm5 2h2v2H7V9zM3 7h2v1H3V7zm6 5H7v1h1v1h1v-1zM6 7v1H5v1h2V7zm2-3h1v2H8V4zm1 4v1h2V7H8v1zM7 6h1v1H7V6zm2 8h2v2H9v-2zm-2 0h1v2H7v-2zm2-3h1v1H9v-1zm0-8V1H8V0H7v4h1V3zm3 11h1v2h-1v-2zm0-2h2v1h-2v-1zm-1 1h1v1h-1v-1zm-1-1h1v1h-1v-1zm4-2v1h1v1h1v-2h-1zm1 3h-1v3h2v-2h-1zm-5-3v1h3V9h-2v1zm2-3v1h2v1h2V7h-2z"/>`
	questionInnerSVG             = `<path fill="currentColor" d="M9 11H6c0-3 1.6-4 2.7-4.6c.4-.2.7-.4.9-.6c.5-.5.3-1.2.2-1.4c-.3-.7-1-1.4-2.3-1.4C5.4 3 5 4.9 5 5.3l-3-.4C2.2 3.2 3.7 0 7.5 0c2.3 0 4.3 1.3 5.1 3.2c.7 1.7.4 3.5-.8 4.7c-.5.5-1.1.8-1.6 1.1c-.9.5-1.2 1-1.2 2zm.5 3a2 2 0 1 1-3.999.001A2 2 0 0 1 9.5 14z"/>`
	questionCircleInnerSVG       = `<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm.9 13h-2v-2h2v2zM11 8.1c-.4.4-.8.6-1.2.7c-.6.4-.8.2-.8 1.2H7c0-2 1.2-2.6 2-3c.3-.1.5-.2.7-.4c.1-.1.3-.3.1-.7c-.2-.5-.8-1-1.7-1c-1.4 0-1.6 1.2-1.7 1.5l-2-.3C4.5 5 5.4 2.9 8 2.9c1.6 0 3 .9 3.6 2.2c.4 1.1.2 2.2-.6 3z"/>`
	questionCircleOInnerSVG      = `<path fill="currentColor" d="M9 10H7c0-2 1.2-2.6 2-3c.3-.1.5-.2.7-.4c.1-.1.3-.3.1-.7c-.2-.5-.8-1-1.7-1c-1.4 0-1.6 1.2-1.7 1.5l-2-.3C4.5 5 5.4 2.9 8 2.9c1.6 0 3 .9 3.6 2.2c.4 1.1.2 2.2-.6 3c-.4.4-.8.6-1.2.7c-.6.4-.8.2-.8 1.2z"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M6.9 11h2v2h-2v-2z"/>`
	quoteLeftInnerSVG            = `<path fill="currentColor" d="M7 7v7H0V6.9c0-4.8 4.5-5.4 4.5-5.4l.6 1.4s-2 .3-2.4 1.9C2.3 6 3.1 7 3.1 7H7zm9 0v7H9V6.9c0-4.8 4.5-5.4 4.5-5.4l.6 1.4s-2 .3-2.4 1.9c-.4 1.2.4 2.2.4 2.2H16z"/>`
	quoteRightInnerSVG           = `<path fill="currentColor" d="M9 9V2h7v7.1c0 4.8-4.5 5.4-4.5 5.4l-.6-1.4s2-.3 2.4-1.9c.4-1.2-.4-2.2-.4-2.2H9zM0 9V2h7v7.1c0 4.8-4.5 5.4-4.5 5.4l-.6-1.4s2-.3 2.4-1.9C4.7 10 3.9 9 3.9 9H0z"/>`
	randomInnerSVG               = `<path fill="currentColor" d="M13 12h-2c-1 0-1.7-1.2-2.4-2.7c-.3.7-.6 1.5-1 2.3C8.4 13 9.4 14 11 14h2v2l3-3l-3-3v2zM5.4 6.6c.3-.7.6-1.5 1-2.2C5.6 3 4.5 2 3 2H0v2h3c1 0 1.7 1.2 2.4 2.6z"/><path fill="currentColor" d="m16 3l-3-3v2h-2C8.3 2 7.1 5 6 7.7C5.2 9.8 4.3 12 3 12H0v2h3c2.6 0 3.8-2.8 4.9-5.6C8.8 6.2 9.7 4 11 4h2v2l3-3z"/>`
	rasterInnerSVG               = `<path fill="currentColor" d="M7 7h1v1H7V7zM5 7h1v1H5V7zM3 7h1v1H3V7zM1 7h1v1H1V7zm5-1h1v1H6V6zM4 6h1v1H4V6zM2 6h1v1H2V6zM0 6h1v1H0V6zm7-1h1v1H7V5zM5 5h1v1H5V5zM3 5h1v1H3V5zM1 5h1v1H1V5zm5-1h1v1H6V4zM4 4h1v1H4V4zM2 4h1v1H2V4zM0 4h1v1H0V4zm7-1h1v1H7V3zM5 3h1v1H5V3zM3 3h1v1H3V3zM1 3h1v1H1V3zm5-1h1v1H6V2zM4 2h1v1H4V2zM2 2h1v1H2V2zM0 2h1v1H0V2zm7-1h1v1H7V1zM5 1h1v1H5V1zM3 1h1v1H3V1zM1 1h1v1H1V1zm5-1h1v1H6V0zM4 0h1v1H4V0zM2 0h1v1H2V0zM0 0h1v1H0V0zm15 7h1v1h-1V7zm-2 0h1v1h-1V7zm-2 0h1v1h-1V7zM9 7h1v1H9V7zm5-1h1v1h-1V6zm-2 0h1v1h-1V6zm-2 0h1v1h-1V6zM8 6h1v1H8V6zm7-1h1v1h-1V5zm-2 0h1v1h-1V5zm-2 0h1v1h-1V5zM9 5h1v1H9V5zm5-1h1v1h-1V4zm-2 0h1v1h-1V4zm-2 0h1v1h-1V4zM8 4h1v1H8V4zm7-1h1v1h-1V3zm-2 0h1v1h-1V3zm-2 0h1v1h-1V3zM9 3h1v1H9V3zm5-1h1v1h-1V2zm-2 0h1v1h-1V2zm-2 0h1v1h-1V2zM8 2h1v1H8V2zm7-1h1v1h-1V1zm-2 0h1v1h-1V1zm-2 0h1v1h-1V1zM9 1h1v1H9V1zm5-1h1v1h-1V0zm-2 0h1v1h-1V0zm-2 0h1v1h-1V0zM8 0h1v1H8V0zM7 15h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm-2 0h1v1H1v-1zm5-1h1v1H6v-1zm-2 0h1v1H4v-1zm-2 0h1v1H2v-1zm-2 0h1v1H0v-1zm7-1h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm-2 0h1v1H1v-1zm5-1h1v1H6v-1zm-2 0h1v1H4v-1zm-2 0h1v1H2v-1zm-2 0h1v1H0v-1zm7-1h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm-2 0h1v1H1v-1zm5-1h1v1H6v-1zm-2 0h1v1H4v-1zm-2 0h1v1H2v-1zm-2 0h1v1H0v-1zm7-1h1v1H7V9zM5 9h1v1H5V9zM3 9h1v1H3V9zM1 9h1v1H1V9zm5-1h1v1H6V8zM4 8h1v1H4V8zM2 8h1v1H2V8zM0 8h1v1H0V8zm15 7h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1V9zm-2 0h1v1h-1V9zm-2 0h1v1h-1V9zM9 9h1v1H9V9zm5-1h1v1h-1V8zm-2 0h1v1h-1V8zm-2 0h1v1h-1V8zM8 8h1v1H8V8z"/>`
	rasterLowerLeftInnerSVG      = `<path fill="currentColor" d="M15 7h1v1h-1V7zm-2 0h1v1h-1V7zm-2 0h1v1h-1V7zM9 7h1v1H9V7zm5-1h1v1h-1V6zm-2 0h1v1h-1V6zm-2 0h1v1h-1V6zm5-1h1v1h-1V5zm-2 0h1v1h-1V5zm-2 0h1v1h-1V5zm3-1h1v1h-1V4zm-2 0h1v1h-1V4zm3-1h1v1h-1V3zm-2 0h1v1h-1V3zm1-1h1v1h-1V2zm1-1h1v1h-1V1zM7 15h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm-2 0h1v1H1v-1zm5-1h1v1H6v-1zm-2 0h1v1H4v-1zm-2 0h1v1H2v-1zm5-1h1v1H7v-1zm-2 0h1v1H5v-1zm-2 0h1v1H3v-1zm3-1h1v1H6v-1zm-2 0h1v1H4v-1zm3-1h1v1H7v-1zm-2 0h1v1H5v-1zm1-1h1v1H6v-1zm1-1h1v1H7V9zm8 6h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H9v-1zm5-1h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1h-1v-1zm-2 0h1v1H8v-1zm7-1h1v1h-1V9zm-2 0h1v1h-1V9zm-2 0h1v1h-1V9zM9 9h1v1H9V9zm5-1h1v1h-1V8zm-2 0h1v1h-1V8zm-2 0h1v1h-1V8zM8 8h1v1H8V8z"/>`
	recordsInnerSVG              = `<path fill="currentColor" d="M4 9h4v2H4V9z"/><path fill="currentColor" d="M16 2h-1V0H5v2H3v1.25L2.4 4H1v1.75L0 7v9h12l4-5V2zM2 5h8v2H2V5zm9 10H1V8h10v7zm1-8h-1V4H4V3h8v4zm2-2.5l-1 1.25V2H6V1h8v3.5z"/>`
	recycleInnerSVG              = `<path fill="currentColor" d="m8 3.1l1.4 2.2l-1.6 1.1l1.3.3l2.8.6l.6-2.7l.4-1.4l-1.8 1.1l-2-3.3H6.9L4.3 5.3l1.7 1zm8 8.9l-2.7-4.3l-1.7 1l2 3.3H11v-2l-3 3l3 3v-2h3.7zM2.4 12l1.4-2.3l1.7 1.1l-.9-4.2l-2.8.7l-1.3.3l1.6 1L0 12l1.3 2H7v-2z"/>`
	refreshInnerSVG              = `<path fill="currentColor" d="M2.6 5.6C3.5 3.5 5.6 2 8 2c3 0 5.4 2.2 5.9 5h2c-.5-3.9-3.8-7-7.9-7c-3 0-5.6 1.6-6.9 4.1L0 3v4h4L2.6 5.6zM16 9h-4.1l1.5 1.4c-.9 2.1-3 3.6-5.5 3.6C5 14 2.5 11.8 2 9H0c.5 3.9 3.9 7 7.9 7c3 0 5.6-1.7 7-4.1L16 13V9z"/>`
	replyInnerSVG                = `<path fill="currentColor" d="M16 8c0-5-4.9-5-4.9-5H6V0L0 6l6 6V9h5.2c3.5 0 1.8 7 1.8 7s3-4.1 3-8z"/>`
	replyAllInnerSVG             = `<path fill="currentColor" d="M16 8c0-5-4.9-5-4.9-5H9V0L3 6l6 6V9h2.2c3.5 0 1.8 7 1.8 7s3-4.1 3-8z"/><path fill="currentColor" d="m0 6l6 6v-1.5L1.5 6L6 1.5V0z"/>`
	resizeHInnerSVG              = `<path fill="currentColor" d="M0 7h16v2H0V7zm7-1h2V3h2L8 0L5 3h2zm2 4H7v3H5l3 3l3-3H9z"/>`
	resizeVInnerSVG              = `<path fill="currentColor" d="M7 0h2v16H7V0zM3 5L0 8l3 3V9h3V7H3zm13 3l-3-3v2h-3v2h3v2z"/>`
	retweetInnerSVG              = `<path fill="currentColor" d="M2 1h12v5h2l-3 3l-3-3h2V3H4v2H2zm12 13H2V9H0l3-3l3 3H4v3h8v-2h2z"/>`
	rhombusInnerSVG              = `<path fill="currentColor" d="M8 0L0 8l8 8l8-8l-8-8zM2 8l6-6l6 6l-6 6l-6-6z"/>`
	roadInnerSVG                 = `<path fill="currentColor" d="M9 11v4h7L12 1H9v3H7V1H4L0 15h7v-4h2zM7 6h2v3H7V6z"/>`
	roadBranchInnerSVG           = `<path fill="currentColor" d="M16 4H0v3h3.2L7 10.6c1.6 1.5 3.6 2.4 5.8 2.4H16v-3h-3.2c-1.4 0-2.7-.5-3.7-1.5L7.5 7H16V4z"/>`
	roadBranchesInnerSVG         = `<path fill="currentColor" d="M16 4V1H0v3h1.7l7.7 9.5c1.3 1.6 3.1 2.5 5 2.5H16v-3h-1.5c-1 0-1.9-.5-2.7-1.4L10.5 10H16V7H8L5.6 4H16z"/>`
	roadSplitInnerSVG            = `<path fill="currentColor" d="M14 13v-1c0-.2 0-4.1-2.8-5.4C9 5.6 9 3.1 9 3V0H7v3c0 .1 0 2.6-2.2 3.6C2 7.9 2 11.8 2 12v1H0l3 3l3-3H4v-1s0-2.8 1.7-3.6c1.1-.5 1.8-1.3 2.3-2c.5.8 1.2 1.5 2.3 2C12 9.2 12 12 12 12v1h-2l3 3l3-3h-2z"/>`
	rocketInnerSVG               = `<path fill="currentColor" d="M16 0s-3.5-.4-6.7 2.8C7.7 4.3 6.4 6.3 5.4 8.1l-2.5-.6l-1.6 1.6l2.8 1.4c-.3.6-.4 1-.4 1l.8.8s.4-.2 1-.4l1.4 2.8l1.6-1.6l-.5-2.5c1.7-1 3.8-2.3 5.3-3.8C16.4 3.6 16 0 16 0zm-3.2 4.8c-.4.4-1.1.4-1.6 0c-.4-.4-.4-1.1 0-1.6c.4-.4 1.1-.4 1.6 0c.4.4.4 1.1 0 1.6z"/><path fill="currentColor" d="M4 14.2c-.8.8-2.6.4-2.6.4s-.4-1.8.4-2.6s1.5-.9 1.5-.9s-1.3-.3-2.1.6c-1.6 1.6-1 4.2-1 4.2s2.6.6 4.2-1c.9-.9.6-2.2.6-2.2s-.2.7-1 1.5z"/>`
	rotateLeftInnerSVG           = `<path fill="currentColor" d="M8 0C5 0 2.4 1.6 1.1 4.1L0 3v4h4L2.5 5.5C3.5 3.5 5.6 2 8 2c3.3 0 6 2.7 6 6s-2.7 6-6 6c-1.8 0-3.4-.8-4.5-2.1L2 13.2C3.4 14.9 5.6 16 8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8z"/>`
	rotateRightInnerSVG          = `<path fill="currentColor" d="M16 7V3l-1.1 1.1C13.6 1.6 11 0 8 0C3.6 0 0 3.6 0 8s3.6 8 8 8c2.4 0 4.6-1.1 6-2.8l-1.5-1.3C11.4 13.2 9.8 14 8 14c-3.3 0-6-2.7-6-6s2.7-6 6-6c2.4 0 4.5 1.5 5.5 3.5L12 7h4z"/>`
	rssInnerSVG                  = `<path fill="currentColor" d="M4.4 13.8a2.2 2.2 0 1 1-4.4 0a2.2 2.2 0 0 1 4.4 0z"/><path fill="currentColor" d="M10.6 16H7.5c0-4.1-3.4-7.5-7.5-7.5V5.4c5.9 0 10.6 4.7 10.6 10.6z"/><path fill="currentColor" d="M12.8 16C12.8 8.9 7.1 3.2 0 3.2V0c8.8 0 16 7.2 16 16h-3.2z"/>`
	rssSquareInnerSVG            = `<path fill="currentColor" d="M0 0v16h16V0H0zm3.6 14c-.9 0-1.6-.7-1.6-1.6s.7-1.6 1.6-1.6s1.6.7 1.6 1.6S4.6 14 3.6 14zm4 0c0-3.1-2.5-5.6-5.6-5.6V6c4.4 0 8 3.6 8 8H7.6zm4 0c0-5.3-4.3-9.6-9.6-9.6V2c6.6 0 12 5.4 12 12h-2.4z"/>`
	safeInnerSVG                 = `<path fill="currentColor" d="M1 0v16h3v-1h8v1h3V0H1zm13 10h-1V5h1v5zm0-7h-1V2H3v11h10v-1h1v2H2V1h12v2zM8.5 7.5c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2z"/><path fill="currentColor" d="M7.5 7.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`
	safeLockInnerSVG             = `<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3.13 14.25l-.37-.9l-.92.38l.37.9c-.659.23-1.419.363-2.21.363s-1.551-.133-2.259-.378l.419-.885l-.92-.38l-.37.9a7.061 7.061 0 0 1-3.102-3.08l.882-.41l-.38-.93l-.9.37c-.23-.659-.363-1.419-.363-2.21s.133-1.551.378-2.259l.885.419l.38-.92l-.9-.37a7.054 7.054 0 0 1 3.08-3.092l.41.882l.92-.38l-.37-.9c.659-.23 1.419-.363 2.21-.363s1.551.133 2.259.378l-.419.885l.92.38l.37-.9a7.061 7.061 0 0 1 3.102 3.08l-.882.41l.38.92l.9-.37c.23.659.363 1.419.363 2.21s-.133 1.551-.378 2.259l-.885-.419l-.38.92l.9.37a7.061 7.061 0 0 1-3.08 3.102z"/><path fill="currentColor" d="M10.36 3.62L9.2 6.41A1.986 1.986 0 0 0 8.001 6h.279l1.15-2.77A4.836 4.836 0 0 0 8.003 3h-.071a5.06 5.06 0 1 0 0 10.12a5.06 5.06 0 0 0 2.454-9.486z"/>`
	scaleInnerSVG                = `<path fill="currentColor" d="m15.81 10l-2.5-5H14a.5.5 0 0 0 0-1h-.79a6.04 6.04 0 0 0-4.198-1.95L9 2a1 1 0 0 0-2 0v.05a6.168 6.168 0 0 0-4.247 1.947L2 4a.5.5 0 0 0 0 1h.69l-2.5 5H0c0 1.1 1.34 2 3 2s3-.9 3-2h-.19L3.26 4.91a.525.525 0 0 0 .159-.148A4.842 4.842 0 0 1 6.994 3.06L7 14H6v1H4v1h8v-1h-2v-1H9V3.06a4.71 4.71 0 0 1 3.524 1.693a.519.519 0 0 0 .193.186L10.19 10H10c0 1.1 1.34 2 3 2s3-.9 3-2h-.19zM5 10H1l2-3.94zm6 0l2-3.94L15 10h-4z"/>`
	scaleUnbalanceInnerSVG       = `<path fill="currentColor" d="m15.81 9l-2.47-4.93l.83-.15a.509.509 0 1 0-.183-.999l-.777.14a5.96 5.96 0 0 0-4.236-1.178a.958.958 0 0 0-.967-.882h-.008a1 1 0 0 0-1 1v.2a6.332 6.332 0 0 0-4.066 2.697l-.754.153a.503.503 0 1 0 .092 1h.088l.35-.05l-2.52 5h-.19c0 1.1 1.34 2 3 2s3-.9 3-2h-.19l-2.56-5.12h.1a.512.512 0 0 0 .379-.297c.021-.093.701-1.583 3.271-2.363v10.78h-1v1h-2v1h8v-1h-2v-1h-1V2.881a4.617 4.617 0 0 1 3.686 1.065l-.006-.005l-2.49 5.06h-.19c0 1.1 1.34 2 3 2s3-.9 3-2h-.19zM5 11H1l2-3.94zm6-2l2-3.94L15 9h-4z"/>`
	scatterChartInnerSVG         = `<path fill="currentColor" d="M1 15V0H0v16h16v-1H1z"/><path fill="currentColor" d="M5 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm6-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm-3 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`
	scissorsInnerSVG             = `<path fill="currentColor" d="M16 3.1s-2.1-1.1-3.5-1c-.3 0-.5.1-.7.2L7.5 5.7L5.7 4.2c.1-.3.2-.6.3-1C6.1 1.4 4.6-.2 2.7 0C1.5.1.4 1 .1 2.2c-.3 1.3.2 2.5 1.2 3.2L4.6 8l-3.3 2.6c-1 .7-1.5 1.9-1.2 3.2c.3 1.2 1.4 2 2.6 2.2c1.9.2 3.4-1.4 3.2-3.2c0-.3-.1-.7-.3-1l1.8-1.5l4.3 3.4c.2.1.4.2.7.2c1.4.1 3.5-1 3.5-1L10.2 8L16 3.1zM2.8 4.6c-.9-.1-1.6-.9-1.5-1.8s.9-1.6 1.8-1.5c.9.1 1.6.9 1.5 1.8c0 .9-.9 1.6-1.8 1.5zm.3 10.1c-.9.1-1.7-.6-1.8-1.5s.6-1.7 1.5-1.8c.9-.1 1.7.6 1.8 1.5s-.6 1.7-1.5 1.8zm9.3-11.5h.2c.4 0 .9.1 1.4.2L7.2 9.1L6.3 8l6.1-4.8zm1.6 9.4c-.5.2-1 .3-1.4.2h-.2l-4-3.2l1-.9l4.6 3.9z"/>`
	screwdriverInnerSVG          = `<path fill="currentColor" d="m8 10.8l.9-.8l-.9-.9l5.7-5.7l1.2-.4L16 .8l-.7-.7l-2.3 1l-.5 1.2L6.9 8L6 7.1l-.8.9s.8.6-.1 1.5c-.5.5-1.3-.1-2.8 1.4L.2 13s-.6 1 .6 2.2s2.2.6 2.2.6l2.1-2.1c1.4-1.4.9-2.3 1.3-2.7c.9-.9 1.6-.2 1.6-.2zm-3.1-.4l.7.7l-3.8 3.8l-.7-.7z"/>`
	searchInnerSVG               = `<path fill="currentColor" d="m15.7 14.3l-4.2-4.2c-.2-.2-.5-.3-.8-.3c.8-1 1.3-2.4 1.3-3.8c0-3.3-2.7-6-6-6S0 2.7 0 6s2.7 6 6 6c1.4 0 2.8-.5 3.8-1.4c0 .3 0 .6.3.8l4.2 4.2c.2.2.5.3.7.3s.5-.1.7-.3c.4-.3.4-.9 0-1.3zM6 10.5c-2.5 0-4.5-2-4.5-4.5s2-4.5 4.5-4.5s4.5 2 4.5 4.5s-2 4.5-4.5 4.5z"/>`
	searchMinusInnerSVG          = `<path fill="currentColor" d="m15.7 14.3l-4.2-4.2c-.2-.2-.5-.3-.8-.3c.8-1 1.3-2.4 1.3-3.8c0-3.3-2.7-6-6-6S0 2.7 0 6s2.7 6 6 6c1.4 0 2.8-.5 3.8-1.4c0 .3 0 .6.3.8l4.2 4.2c.2.2.5.3.7.3s.5-.1.7-.3c.4-.3.4-.9 0-1.3zM6 10.5c-2.5 0-4.5-2-4.5-4.5s2-4.5 4.5-4.5s4.5 2 4.5 4.5s-2 4.5-4.5 4.5z"/><path fill="currentColor" d="M3 5h6v2H3V5z"/>`
	searchPlusInnerSVG           = `<path fill="currentColor" d="m15.7 14.3l-4.2-4.2c-.2-.2-.5-.3-.8-.3c.8-1 1.3-2.4 1.3-3.8c0-3.3-2.7-6-6-6S0 2.7 0 6s2.7 6 6 6c1.4 0 2.8-.5 3.8-1.4c0 .3 0 .6.3.8l4.2 4.2c.2.2.5.3.7.3s.5-.1.7-.3c.4-.3.4-.9 0-1.3zM6 10.5c-2.5 0-4.5-2-4.5-4.5s2-4.5 4.5-4.5s4.5 2 4.5 4.5s-2 4.5-4.5 4.5z"/><path fill="currentColor" d="M7 3H5v2H3v2h2v2h2V7h2V5H7z"/>`
	selectInnerSVG               = `<path fill="currentColor" d="M15 4H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1zm-3 5l-2-2h4l-2 2z"/>`
	serverInnerSVG               = `<path fill="currentColor" d="M3 5v3h10V5H3zm4 2H4V6h3v1zM3 4h10l-2-4H5zm0 8h10V9H3v3zm8-2h1v1h-1v-1zm-2 0h1v1H9v-1zm-6 6h10v-3H3v3zm1-2h3v1H4v-1z"/>`
	shareInnerSVG                = `<path fill="currentColor" d="M10 3H4.9S0 3 0 8c0 3.9 3 8 3 8S1.3 9 4.8 9H10v3l6-6l-6-6v3z"/>`
	shareSquareInnerSVG          = `<path fill="currentColor" d="M11 3H7.4S3 2.8 3 7.3C3 10.8 5 14 5 14s-.4-7 2.3-7H11v3l5-5l-5-5v3z"/><path fill="currentColor" d="M14 9v6H1V2h9V1H0v15h15V8z"/>`
	shieldInnerSVG               = `<path fill="currentColor" d="M1 0v7c0 5.6 7 9 7 9s7-3.4 7-9V0H1zm13 7c0 4.2-4.6 7.1-6 7.9V1h6v6z"/>`
	shiftInnerSVG                = `<path fill="currentColor" d="M0 2v12h16V2H0zm6 6v3H4V8H2l3-3l3 3H6z"/>`
	shiftArrowInnerSVG           = `<path fill="currentColor" d="M8 2L1 9h4v5h6V9h4zm2 6v5H6V8H3.5L8 3.42L12.5 8H10z"/>`
	shopInnerSVG                 = `<path fill="currentColor" d="M0 15h16v1H0v-1zM0 0v6c.005.732.401 1.37.991 1.715L1 14h9V9h3v5h2V7.72c.599-.35.995-.988 1-1.719V0H0zm4 2h2v4a1 1 0 0 1-2 0V2zM2 7a1 1 0 0 1-1-1V2h2v4a1 1 0 0 1-1 1zm6 5H3V9h5v3zm1-6a1 1 0 0 1-2 0V2h2v4zm3 0a1 1 0 0 1-2 0V2h2v4zm3 0a1 1 0 0 1-2 0V2h2v4z"/>`
	signInInnerSVG               = `<path fill="currentColor" d="M7 1v2l1 1V2h7v12H8v-2l-1 1v2h9V1z"/><path fill="currentColor" d="M10 8L5 4v2H0v4h5v2z"/>`
	signInAltInnerSVG            = `<path fill="currentColor" d="M0 0h2v16H0V0zm3 10h8v3l5-5l-5-5v3H3z"/>`
	signOutInnerSVG              = `<path fill="currentColor" d="M9 4V1H0v14h9v-3H8v2H1V2h7v2z"/><path fill="currentColor" d="m16 8l-5-4v2H6v4h5v2z"/>`
	signOutAltInnerSVG           = `<path fill="currentColor" d="M14 0h2v16h-2V0zM8 6H0v4h8v3l5-5l-5-5z"/>`
	signalInnerSVG               = `<path fill="currentColor" d="M6.9 13.2L8 14.3l1.1-1.1c-.3-.3-.7-.5-1.1-.5s-.9.2-1.1.5zM8 4.6c2.7 0 5.1 1.1 6.9 2.8L16 6.3C14 4.3 11.1 3 8 3S2 4.3 0 6.3l1.1 1.1C2.9 5.7 5.3 4.6 8 4.6z"/><path fill="currentColor" d="m2.3 8.6l1.1 1.1C4.6 8.6 6.2 7.9 8 7.9s3.4.7 4.6 1.9l1.1-1.1C12.3 7.1 10.2 6.2 8 6.2s-4.3.9-5.7 2.4z"/><path fill="currentColor" d="M4.6 10.9L5.7 12c.6-.6 1.4-.9 2.3-.9s1.7.4 2.3.9l1.1-1.1C10.6 10 9.3 9.5 8 9.5s-2.6.5-3.4 1.4z"/>`
	sitemapInnerSVG              = `<path fill="currentColor" d="M14.5 12V7.5h-6V4H10V0H6v4h1.5v3.5h-6V12H0v4h4v-4H2.5V8.5h5V12H6v4h4v-4H8.5V8.5h5V12H12v4h4v-4z"/>`
	sliderInnerSVG               = `<path fill="currentColor" d="M16 6h-3.6c-.7-1.2-2-2-3.4-2s-2.8.8-3.4 2H0v4h5.6c.7 1.2 2 2 3.4 2s2.8-.8 3.4-2H16V6zM1 9V7h4.1c0 .3-.1.7-.1 1s.1.7.1 1H1zm8 2c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3z"/>`
	slidersInnerSVG              = `<path fill="currentColor" d="M7 0h2v3H7V0zM6 4v3h1v9h2V7h1V4zM2 0h2v8H2V0zM1 9v3h1v4h2v-4h1V9zm11-9h2v10h-2V0zm-1 11v3h1v2h2v-2h1v-3z"/>`
	smileyOInnerSVG              = `<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M8 13.2c-2 0-3.8-1.2-4.6-3.1l.9-.4c.6 1.5 2.1 2.4 3.7 2.4s3.1-1 3.7-2.4l.9.4c-.8 2-2.6 3.1-4.6 3.1zM7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`
	sortInnerSVG                 = `<path fill="currentColor" d="M11 7H5l3-4zM5 9h6l-3 4z"/>`
	soundDisableInnerSVG         = `<path fill="currentColor" d="M4 5H0v6h4l5 4V1zm11.9.6l-.8-.7l-2.3 2.4l-2.4-2.4l-.8.7L12 8l-2.4 2.4l.8.7l2.4-2.4l2.3 2.4l.8-.7L13.5 8z"/>`
	sparkLineInnerSVG            = `<path fill="currentColor" d="M14 6a2 2 0 0 0-2 2v.16l-.81.25l-2.3-3.48l-1.73 4.32L6 3.44L3.7 8.22L2.06 6.91L0 8v1.08l1.94-1l2.11 1.7l1.56-3.22l1.23 6.19l2.27-5.68l1.68 2.52l1.55-.48A2 2 0 1 0 14.004 6H14zm0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/>`
	specialistInnerSVG           = `<path fill="currentColor" d="M4.1 8c.2.6.3 1.1.3 1.1c.8 1.3 1.8 1.1 1.8 1.8c0 .3-.2.6-.5.7L8 13.4l2.3-1.7c-.3-.2-.5-.4-.5-.7c0-.8 1-.5 1.8-1.8c0 0 .2-.4.3-1.1c.3-1.1.6-3.1.5-4.1h-1.5c0-.3.1-.6.1-1h1.1c-.3-1.4-1-2-2.2-2.3C9.4.3 8.7 0 8 0S6.6.3 6.1.7C4.9 1 4.3 1.6 3.9 3H5c0 .4.1.7.2 1H3.6c-.1 1 .2 3 .5 4zm7.1.5l-.3.3l-.5.6c-.4.5-.8.8-1.4.9l-.4.1c-.4.1-.9.1-1.4 0l-.4-.1c-.6-.2-1.1-.5-1.5-1.1l-.2-.4l-.3-.3l-.7-.5l3.1-.9c.5-.1 1-.2 1.5 0l3.2.9l-.7.5zM6 3c0-1.1.9-2 2-2s2 .9 2 2s-.9 2-2 2s-2-.9-2-2z"/><path fill="currentColor" d="M15.5 14.2c-1.3-2.4-2.6-2-3.9-2.2h-.1L8 14.6L4.5 12h-.1c-1.4.1-2.6-.2-3.9 2.2c-.2.4-.4 1.1-.5 1.8h16c-.1-.7-.3-1.4-.5-1.8z"/>`
	spinnerInnerSVG              = `<path fill="currentColor" d="m9.9.2l-.2 1C12.7 2 15 4.7 15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7c0-3.3 2.3-6 5.3-6.8l-.2-1C2.6 1.1 0 4.3 0 8c0 4.4 3.6 8 8 8s8-3.6 8-8c0-3.7-2.6-6.9-6.1-7.8z"/>`
	spinnerArcInnerSVG           = `<path fill="currentColor" d="M15 8c0 3.9-3.1 7-7 7s-7-3-7-7H0c0 4 3.6 8 8 8s8-3.6 8-8h-1z"/>`
	spinnerThirdInnerSVG         = `<path fill="currentColor" d="M12.9 3.1C14.2 4.3 15 6.1 15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7c0-1.9.8-3.7 2.1-4.9l-.8-.8C.9 3.8 0 5.8 0 8c0 4.4 3.6 8 8 8s8-3.6 8-8c0-2.2-.9-4.2-2.3-5.7l-.8.8z"/>`
	splineAreaChartInnerSVG      = `<path fill="currentColor" d="M1 15V0H0v16h16v-1H1z"/><path fill="currentColor" d="M10 7C8 7 7.92 6 6 6C3.66 6 2 9 2 9v5h14V2c-2 0-3.86 5-6 5z"/>`
	splineChartInnerSVG          = `<path fill="currentColor" d="M1 15V0H0v16h16v-1H1z"/><path fill="currentColor" d="M12 5c-.69 1-1.41 2-2 2l-.087.001c-.601 0-1.164-.16-1.65-.44a4.519 4.519 0 0 0-2.2-.562h-.067a5.83 5.83 0 0 0-3.991 1.993l-.006 2.347c.77-1.12 2.32-2.84 4-2.84h.048c.579 0 1.121.156 1.587.428a4.682 4.682 0 0 0 2.264.573l.106-.001c1.395 0 2.335-1.32 3.245-2.6s1.75-2.4 2.75-2.4v-1.5c-1.81 0-3 1.61-4 3z"/>`
	splitInnerSVG                = `<path fill="currentColor" d="M0 11h6v5H0v-5zm11-1V8l-.64.64a4.427 4.427 0 0 1-1.358-3.658L11.001 5V0h-6v5h2a4.43 4.43 0 0 1-1.358 3.639l-.642-.638v2h2l-.65-.65A5.426 5.426 0 0 0 8 4.982c-.01.149-.016.299-.016.45c0 1.539.639 2.928 1.665 3.917l-.648.652h2zm-1 1h6v5h-6v-5z"/>`
	splitHInnerSVG               = `<path fill="currentColor" d="M0 1v14h16V1H0zm1 3h6.5v10H1V4zm14 10H8.5V4H15v10zm0-11h-1V2h1v1z"/>`
	splitVInnerSVG               = `<path fill="currentColor" d="M0 1v14h16V1H0zm14 1h1v1h-1V2zm1 2v4.5H1V4h14zM1 14V9.5h14V14H1z"/>`
	spoonInnerSVG                = `<path fill="currentColor" d="M10.5 4.8c0-1.8-.9-4.8-3-4.8s-3 3-3 4.8c0 1.5.8 2.8 2.2 3.1c-.5 1.6-.7 4.6-.7 4.6v2c0 .8.7 1.5 1.5 1.5S9 15.3 9 14.5v-2c0-.6-.2-3.2-.7-4.6c1.4-.3 2.2-1.6 2.2-3.1z"/>`
	squareShadowInnerSVG         = `<path fill="currentColor" d="M14 2V0H0v14h2v2h14V2h-2zm-1 11H1V1h12v12z"/>`
	starInnerSVG                 = `<path fill="currentColor" d="M12.9 15.4L8 12.8l-4.9 2.6L4 10L0 6.1l5.5-.8l2.4-5l2.4 5l5.5.8L12 10l.9 5.4z"/>`
	starHalfLeftInnerSVG         = `<path fill="currentColor" d="m5.6 5.4l-5.5.8L4 10l-.9 5.5L8 12.9V.4z"/>`
	starHalfLeftOInnerSVG        = `<path fill="currentColor" d="m15.9 6.2l-5.5-.8L8 .4l-2.4 5l-5.5.8L4 10l-.9 5.4L8 12.9l4.9 2.6L12 10l3.9-3.8zM8 11.8V2.7l1.8 3.6l4 .6l-2.9 2.8l.7 4L8 11.8z"/>`
	starHalfRightInnerSVG        = `<path fill="currentColor" d="m10.5 5.4l5.5.8l-4 3.8l.9 5.5L8 12.9V.4z"/>`
	starHalfRightOInnerSVG       = `<path fill="currentColor" d="m15.9 6.2l-5.5-.8L8 .4l-2.4 5l-5.5.8L4 10l-.9 5.4L8 12.9l4.9 2.6L12 10l3.9-3.8zM4.4 13.7l.7-4l-2.9-2.8l4-.6L8 2.7v9.1l-3.6 1.9z"/>`
	starOInnerSVG                = `<path fill="currentColor" d="m15.9 6.2l-5.5-.8L8 .4l-2.4 5l-5.5.8L4 10l-.9 5.4L8 12.9l4.9 2.6L12 10l3.9-3.8zM8 11.8l-3.6 1.9l.7-4l-2.9-2.8l4-.6L8 2.7l1.8 3.6l4 .6l-2.9 2.8l.7 4L8 11.8z"/>`
	startCogInnerSVG             = `<path fill="currentColor" d="M4 0v6h1.7l.2.7l.2.6h.1l1.2-.6l1.8 1.8l-.6 1.2v.1l.6.2l.7.2v.2L16 7L4 0zm.5 10.5c-.2 0-.4.1-.5.2c-.3.2-.5.5-.5.8s.2.7.5.8c.1.1.3.2.5.2c.6 0 1-.4 1-1s-.4-1-1-1z"/><path fill="currentColor" d="M9 12v-1l-1.1-.4c-.1-.3-.2-.6-.4-.9l.5-1l-.7-.7l-1 .5c-.3-.2-.6-.3-.9-.4L5 7H4l-.4 1.1c-.3.1-.6.2-.9.4l-1-.5l-.7.7l.5 1.1c-.2.3-.3.6-.4.9L0 11v1l1.1.4c.1.3.2.6.4.9l-.5 1l.7.7l1.1-.5c.3.2.6.3.9.4L4 16h1l.4-1.1c.3-.1.6-.2.9-.4l1 .5l.7-.7l-.5-1.1c.2-.3.3-.6.4-.9L9 12zm-4.5 1.5c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2z"/>`
	stepBackwardInnerSVG         = `<path fill="currentColor" d="M14 15V1L4 8zM2 1h2v14H2V1z"/>`
	stepForwardInnerSVG          = `<path fill="currentColor" d="M2 1v14l10-7zm10 0h2v14h-2V1z"/>`
	stethoscopeInnerSVG          = `<path fill="currentColor" d="M5.7 15.2c.3.3 1 .8 1.8.8c2.7 0 3.3-2 3.4-3.6c.2-2.3.8-2.2 1.1-2.2c.7 0 .9.4.9 1.1c-.6.4-1 1-1 1.7c0 1.1.9 2 2 2s2-.9 2-2s-.9-2-2-2h-.2c-.2-.9-.7-1.8-1.8-1.8c-1.6 0-2 1.4-2.1 2.9C9.7 14.2 9 15 7.5 15c-.4 0-.8-.2-1-.4c-.6-.5-.5-2.3-.5-2.3c2 0 4-1.8 4.7-4.8l-.2-.1c.3-1.2.5-2.6.5-3.6c0-1.1-.3-1.9-1-2.5S8.5.5 7.9.5C7.7.2 7.4 0 7 0c-.5 0-1 .4-1 1s.4 1 1 1c.4 0 .7-.2.8-.5c.5 0 1 .2 1.5.6s.7.9.7 1.7c0 .9-.2 2.2-.5 3.5l-.2-.1C9 8.3 8 10.8 6 10.8H5c-2 0-3-2.5-3.3-3.6l-.2.1C1.2 6 1 4.7 1 3.8c0-.8.2-1.3.7-1.7c.4-.4 1-.5 1.5-.6c.1.3.4.5.8.5c.6 0 1-.4 1-1s-.4-1-1-1c-.4 0-.7.2-.9.5c-.6 0-1.4.2-2.1.8S0 2.7 0 3.8c0 1 .2 2.4.5 3.7l-.2.1C1 10.5 3 12.3 5 12.3c0 0-.1 2.2.7 2.9zM14 14c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.5 1-1 1z"/>`
	stockInnerSVG                = `<path fill="currentColor" d="M12 6V0H4v6H0v7h16V6h-4zm-5 6H1V7h2v1h2V7h2v5zM5 6V1h2v1h2V1h2v5H5zm10 6H9V7h2v1h2V7h2v5zM0 16h3v-1h10v1h3v-2H0v2z"/>`
	stopInnerSVG                 = `<path fill="currentColor" d="M1 1h14v14H1V1z"/>`
	stopCogInnerSVG              = `<path fill="currentColor" d="M1 0v7.2l.5-.5l1.2.6h.1l.2-.6l.3-.7h2.4l.2.7l.2.6h.1l1.2-.6l1.8 1.8l-.6 1.2v.1l.6.2l.7.2v2.4l-.7.2l-.6.2v.1l.6 1.2l-.4.7H16V0H1z"/><path fill="currentColor" d="M5.5 11.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/><path fill="currentColor" d="M7.9 12.4L9 12v-1l-1.1-.4c-.1-.3-.2-.6-.4-.9l.5-1l-.7-.7l-1 .5c-.3-.2-.6-.3-.9-.4L5 7H4l-.4 1.1c-.3.1-.6.2-.9.4l-1-.5l-.7.7l.5 1.1c-.2.3-.3.6-.4.9L0 11v1l1.1.4c.1.3.2.6.4.9l-.5 1l.7.7l1.1-.5c.3.2.6.3.9.4L4 16h1l.4-1.1c.3-.1.6-.2.9-.4l1 .5l.7-.7l-.5-1.1c.2-.2.3-.5.4-.8zm-3.4 1.1c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2z"/>`
	stopwatchInnerSVG            = `<path fill="currentColor" d="M8.5 8.14V4.5h-1v3.64a1.001 1.001 0 0 0 .5 1.866a1 1 0 0 0 .505-1.863z"/><path fill="currentColor" d="M8 2a7 7 0 1 0 0 14A7 7 0 0 0 8 2zm0 12.5a5.5 5.5 0 1 1 0-11A5.5 5.5 0 0 1 13.5 9a5.51 5.51 0 0 1-5.499 5.5zM6 0h4v1.5H6V0z"/><path fill="currentColor" d="m.005 4.438l2.713-2.939L3.82 2.516L1.107 5.455L.005 4.438zm12.181-1.919l1.102-1.017l2.713 2.939l-1.102 1.017l-2.713-2.939z"/>`
	storageInnerSVG              = `<path fill="currentColor" d="M16 4L7.94 0L0 4v1h1v11h2V7h10v9h2V5h1V4zM4 6V5h2v1H4zm3 0V5h2v1H7zm3 0V5h2v1h-2z"/><path fill="currentColor" d="M6 9H5V8H4v3h3V8H6v1zm0 4H5v-1H4v3h3v-3H6v1zm4 0H9v-1H8v3h3v-3h-1v1z"/>`
	strikethroughInnerSVG        = `<path fill="currentColor" d="M10.5 7c-.5-.3-1-.5-1.4-.7c-2-.9-2.1-1.1-2-1.9s.4-1 .6-1.2c.9-.5 2.8-.1 3.5.2L12.3.6C11.9.4 8.6-.8 6.2.6c-.8.5-1.9 1.5-2.1 3.4c-.2 1.3.1 2.3.7 3H0v1h16V7h-5.5zM7.7 9s.1 0 .1.1c2 .9 2.4 1.2 2.2 2.5c-.2.9-.5 1.1-.8 1.3c-1.1.6-3.3 0-4.4-.5L3.6 15c.3.1 2.3 1 4.5 1c.9 0 1.8-.2 2.6-.6c.9-.5 2-1.4 2.4-3.4c.2-1.3 0-2.3-.4-3.1h-5z"/>`
	subscriptInnerSVG            = `<path fill="currentColor" d="M16 15v1h-4v-1s3.3-1.6 2.6-3.2c-.5-1.1-2-.2-2-.2l-.5-.9s1.9-1.4 3.1-.2c2.4 2.3-1.4 4.5-1.4 4.5H16zM12 3H8.6L6 6L3.4 3H0l4.3 5L0 13h3.4L6 10l2.6 3H12L7.7 8z"/>`
	suitcaseInnerSVG             = `<path fill="currentColor" d="M11 3V1H5v2H0v12h16V3h-5zM4 14H3V4h1v10zm6-11H6V2h4v1zm3 11h-1V4h1v10z"/>`
	sunDownInnerSVG              = `<path fill="currentColor" d="M10 3H9V1H7v2H6l2 3l2-3zm4 10l-1.58-1.18l.78-1.82l-2-.23L11 7.8l-1.82.78L8 7L6.82 8.58L5 7.8l-.23 2l-1.97.2l.78 1.82L2 13H0v1h16v-1h-2zM4 13a4.143 4.143 0 0 1 3.993-4A4.143 4.143 0 0 1 12 12.993L4 13z"/>`
	sunOInnerSVG                 = `<path fill="currentColor" d="m16 8l-2.2-1.6L14.9 4l-2.7-.2l-.2-2.7l-2.4 1.1L8 0L6.4 2.2L4 1.1l-.2 2.7l-2.7.2l1.1 2.4L0 8l2.2 1.6L1.1 12l2.7.2l.2 2.7l2.4-1.1L8 16l1.6-2.2l2.4 1.1l.2-2.7l2.7-.2l-1.1-2.4L16 8zm-8 5c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5z"/>`
	sunRiseInnerSVG              = `<path fill="currentColor" d="M6 4h1v2h2V4h1L8 1L6 4zm6.42 7.82L13.2 10l-2-.23L11 7.8l-1.82.78L8 7L6.82 8.58L5 7.8l-.23 2l-1.97.2l.78 1.82L2 13H0v1h16v-1h-2zM4 13a4.143 4.143 0 0 1 3.993-4A4.143 4.143 0 0 1 12 12.993L4 13z"/>`
	superscriptInnerSVG          = `<path fill="currentColor" d="M16 5v1h-4V5s3.3-1.6 2.6-3.2c-.5-1.1-2-.2-2-.2l-.5-.9S14-.7 15.2.5C17.6 2.8 13.8 5 13.8 5H16zm-4-2H8.6L6 6L3.4 3H0l4.3 5L0 13h3.4L6 10l2.6 3H12L7.7 8z"/>`
	swordInnerSVG                = `<path fill="currentColor" d="m15.8.5l-.1-.2l-.2-.1c-.1 0-2.5-.8-4.2.9L4.6 7.7c-.9-.6-1.7-1.2-1.8-1l-.4.3c-.2.2.9 1.7 1.8 2.7l-2.5 3.4c-.3-.3-.8-.3-1.1 0l-.3.3c-.3.3-.3.8 0 1.1l1 1c.3.3.8.3 1.1 0l.3-.3c.3-.3.3-.8 0-1.1l3.5-2.5c1 .9 2.5 2 2.7 1.8l.4-.4c.1-.1-.4-1-1.1-1.8l6.7-6.7c1.7-1.5.9-3.9.9-4zm-8.1 10l-.8-.8l6.2-6.9L6.2 9l-.7-.7L12 1.8c1-1 2.3-.8 2.9-.7c.1.6.3 1.9-.7 2.8l-6.5 6.6z"/>`
	tabInnerSVG                  = `<path fill="currentColor" d="M0 2v12h16V2H0zm13 9h-1V8l-3 3V9H3V7h6V5l3 3V5h1v6z"/>`
	tabAInnerSVG                 = `<path fill="currentColor" d="M9 10H0V6h9V4l5 4l-5 4v-2zm5-6h2v8h-2V4z"/>`
	tableInnerSVG                = `<path fill="currentColor" d="M0 1v15h16V1H0zm5 14H1v-2h4v2zm0-3H1v-2h4v2zm0-3H1V7h4v2zm0-3H1V4h4v2zm5 9H6v-2h4v2zm0-3H6v-2h4v2zm0-3H6V7h4v2zm0-3H6V4h4v2zm5 9h-4v-2h4v2zm0-3h-4v-2h4v2zm0-3h-4V7h4v2zm0-3h-4V4h4v2z"/>`
	tabletInnerSVG               = `<path fill="currentColor" d="M0 2v12h16V2H0zm13 11H2V3h11v10zm2-4h-1V7h1v2z"/>`
	tabsInnerSVG                 = `<path fill="currentColor" d="M14 4V2H0v12h16V4h-2zm-4-1h3v1h-3V3zM6 3h3v1H6V3zm9 10H1V3h4v2h10v8z"/>`
	tagInnerSVG                  = `<path fill="currentColor" d="M8 1H1v7l7 7l7-7zM3.75 5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5z"/>`
	tagsInnerSVG                 = `<path fill="currentColor" d="M9 2H7.5l7 7l-5.3 5.2l.8.8l6-6z"/><path fill="currentColor" d="M6 2H0v6l7 7l6-6l-7-7zM2.8 6c-.7 0-1.3-.6-1.3-1.2s.6-1.2 1.2-1.2S4 4.1 4 4.8S3.4 6 2.8 6z"/>`
	tasksInnerSVG                = `<path fill="currentColor" d="M6 0h10v4H6V0zm0 6h10v4H6V6zm0 6h10v4H6v-4zM3 1v2H1V1h2zm1-1H0v4h4V0zM3 13v2H1v-2h2zm1-1H0v4h4v-4zm1.3-6.1l-.6-.8l-.9.9H0v4h4V7.2l1.3-1.3zM2.7 7l-.7.7l-.8-.7h1.5zM1 8.2l.9.8H1v-.8zM3 9h-.9l.9-.9V9z"/>`
	taxiInnerSVG                 = `<path fill="currentColor" d="m15 6.1l-1.4-2.9c-.4-.7-1.1-1.2-2-1.2H11V.7c0-.4-.3-.7-.7-.7H5.7c-.4 0-.7.3-.7.7V2h-.7c-.8 0-1.6.5-1.9 1.2L1 6.1c-.6.1-1 .6-1 1.1v3.5c0 .6 0 1.1 1 1.2v2c0 .6.4 1.1 1 1.1h.9c.6 0 1.1-.5 1.1-1.1V12h8v1.9c0 .6.4 1.1 1 1.1h.9c.6 0 1.1-.5 1.1-1.1v-2c1-.1 1-.6 1-1.2V7.2c0-.5-.4-1-1-1.1zM4 8.4c0 .3-.3.6-.6.6H1.6c-.3 0-.6-.3-.6-.6v-.8c0-.3.3-.6.6-.6h1.8c.3 0 .6.3.6.6v.8zm6 2.6H6v-1h4v1zM2.1 6l1.2-2.4c.2-.4.6-.6 1-.6h7.4c.4 0 .8.2 1 .6L13.9 6H2.1zM15 8.4c0 .3-.3.6-.6.6h-1.8c-.3 0-.6-.3-.6-.6v-.8c0-.3.3-.6.6-.6h1.8c.3 0 .6.3.6.6v.8z"/>`
	teethInnerSVG                = `<path fill="currentColor" d="M4.6 7.6c-.1.1-.5.4-1.6.4c1.1 0 1.5.3 1.6.4c.2-.2.6-.4 1.5-.4c-.9 0-1.3-.2-1.5-.4z"/><path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm5.1 11.6c-1 0-1.4-.8-1.6-1.6c-.2.9-.6 2-1.8 2c-1.1 0-1.5-1.1-1.7-2c-.2 1-.6 2-1.7 2s-1.6-1.1-1.8-2c-.2.8-.6 1.6-1.6 1.6c-2 0-1.9-3-1.9-3S1.2 8 2.7 8C1.2 8 1 7.5 1 7.5s-.1-3 1.9-3c1 0 1.4.8 1.6 1.6c.2-.9.6-2 1.8-2C7.4 4 7.8 5.1 8 6c.2-1 .6-2 1.7-2s1.6 1.1 1.8 2c.2-.8.6-1.6 1.6-1.6c2 0 1.9 3 1.9 3s-.3.6-1.8.6c-1.2 0-1.6-.3-1.8-.4c-.2.2-.7.4-1.6.4c-1.2 0-1.6-.2-1.8-.4c-.1.1-.6.4-1.6.4c1 0 1.4.3 1.6.4c.2-.2.6-.4 1.8-.4c1 0 1.4.2 1.7.4c0-.1.5-.4 1.7-.4c1.5 0 1.8.6 1.8.6s.1 3-1.9 3z"/>`
	terminalInnerSVG             = `<path fill="currentColor" d="M6 12h9v1H6v-1zm-4.9 1h1.2L6 8L2.3 3H1l3.8 5z"/>`
	textHeightInnerSVG           = `<path fill="currentColor" d="M15 3h1l-1.5-3L13 3h1v10h-1l1.5 3l1.5-3h-1zM1 0v3h4v13h3V3h4V0z"/>`
	textInputInnerSVG            = `<path fill="currentColor" d="M2 2h1v4H2V2z"/><path fill="currentColor" d="M1 0C.4 0 0 .4 0 1v14c0 .6.4 1 1 1h15V0H1zm12 15H1V1h12v14zm2 0h-1v-1h1v1zm0-2h-1V3h1v10zm0-11h-1V1h1v1z"/>`
	textLabelInnerSVG            = `<path fill="currentColor" d="M12.5 4.9c-1.4 0-2.5.8-2.6.9l1.2 1.6s.7-.5 1.4-.5c1.4 0 1.5 1.2 1.5 1.6c-.4-.1-1.1-.3-2-.1c-1.4.3-2.8 2-2.1 3.9c.7 1.8 3.1 2.1 4.1.6v1h2V8.6c0-2.7-1.9-3.7-3.5-3.7zm-1 6.5C11.4 9.5 13 9.5 14 9.6v1c0 1.2-2.3 2.3-2.5.8zM6.9 14H9L5.8 2H3.1L0 14h2.1l1-4h2.7l1.1 4zM3.6 8l.8-3.2l.9 3.2H3.6z"/>`
	textWidthInnerSVG            = `<path fill="currentColor" d="M15 14.5L12 13v1H3v-1l-3 1.5L3 16v-1h9v1zM0 0v3h6v9h3V3h6V0z"/>`
	thinSquareInnerSVG           = `<path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/>`
	thumbsDownInnerSVG           = `<path fill="currentColor" d="M15.6 7.8s.5.5.4 1.6c0 1.5-1.6 1.6-1.6 1.6H12c-.2 0-.3.2-.3.4c.3.7.8 2.1.6 3.1c-.3 1.4-1.5 1.5-1.9 1.5c-.1 0-.2-.1-.2-.2l-1-2.8s0-.1-.1-.1l-2.6-2.8c-.1-.1-.2-.1-.3-.1H6V3h.2c.7 0 3.2-2 5.4-2s2.7.3 3.1 1c.4.7.1 1.3.1 1.3s.5.3.6 1c.1.7-.1 1.1-.1 1.1s.5.4.5 1.2c.1.9-.2 1.2-.2 1.2zM0 11h5V3H0v8zm2.5-3.5c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1z"/>`
	thumbsDownOInnerSVG          = `<path fill="currentColor" d="M15.6 7.3c.1-.3.3-.7.2-1.2c0-.6-.3-1.1-.5-1.3c.1-.3.1-.6 0-1.1s-.4-.8-.6-1c.1-.3.1-.8-.3-1.4C14 .3 13.2 0 10.8 0C9.1 0 7.5.8 6.2 1.5c-.4.2-1 .5-1.2.5H0v9h5v-.9l2.7 2.7l1 2.8c.2.2.4.4.8.4h.1c.5 0 2-.1 2.4-1.9c.2-.9-.1-2.2-.5-3.1h2.3c.7-.1 2.1-.6 2.2-2.1c0-.7-.2-1.3-.4-1.6zm-13.1.2c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zM13.8 10h-2.5c-.3 0-.5.1-.7.4c-.2.2-.2.5-.1.8c.5 1.2.7 2.2.6 2.8c-.2.9-.9 1.1-1.4 1.1l-1-2.7c0-.1-.1-.2-.2-.3L5.6 9.2c-.1-.1-.3-.2-.5-.2H5V3c.4 0 .8-.2 1.7-.6C7.8 1.8 9.4 1 10.8 1c2.5 0 2.7.4 2.9.7c.3.5.1.9.1.9l-.2.4l.4.3s.4.2.5.7c.1.4 0 .7 0 .7l-.3.3l.3.3s.4.3.4.9c0 .5-.2.7-.2.7l-.4.3l.4.4s.4.4.3 1.2c0 1.1-1.1 1.2-1.2 1.2z"/>`
	thumbsUpInnerSVG             = `<path fill="currentColor" d="M15.6 8.2s.5-.5.4-1.6C16 5.1 14.4 5 14.4 5H12c-.2 0-.3-.2-.3-.4c.3-.7.8-2.1.6-3.1C12 .1 10.8 0 10.4 0c-.1 0-.2.1-.2.2L9.2 3s0 .1-.1.1L6.5 5.9c-.1.1-.2.1-.3.1H6v7h.2c.7 0 3.2 2 5.4 2s2.7-.3 3.1-1c.4-.7.1-1.3.1-1.3s.5-.3.6-1c.1-.7-.1-1.1-.1-1.1s.5-.4.5-1.2c.1-.9-.2-1.2-.2-1.2zM0 14h5V6H0v8zm2.5-3.5c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1z"/>`
	thumbsUpOInnerSVG            = `<path fill="currentColor" d="M16 7.1C16 5.6 14.6 5 13.8 5h-2.2c.4-1 .7-2.2.5-3.1C11.6.1 10.1 0 9.6 0h-.1c-.4 0-.6.2-.8.5l-1 2.8L5 6H0v9h5v-1c.2 0 .7.3 1.2.6c1.2.6 2.9 1.5 4.5 1.5c2.4 0 3.2-.3 3.8-1.3c.3-.6.3-1.1.3-1.4c.2-.2.5-.5.6-1s.1-.8 0-1.1c.2-.3.4-.7.5-1.3c0-.5-.1-.9-.2-1.2c.1-.4.3-.9.3-1.7zM2.5 13.5c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1zm12.2-4.4s.2.2.2.7c0 .6-.4.9-.4.9l-.3.3l.2.3s.2.3 0 .7c-.1.4-.5.7-.5.7l-.3.3l.2.4s.2.4-.1.9c-.2.4-.4.7-2.9.7c-1.4 0-3-.8-4.1-1.4c-.8-.4-1.3-.6-1.7-.6V7h.1c.2 0 .4-.1.6-.2L8.5 4c.1-.1.1-.2.2-.3l1-2.7c.5 0 1.2.2 1.4 1.1c.1.6-.1 1.6-.6 2.8c-.1.3-.1.5.1.8c.1.2.4.3.7.3h2.5c.1 0 1.2.2 1.2 1.1c0 .8-.3 1.2-.3 1.2l-.3.4l.3.4z"/>`
	ticketInnerSVG               = `<path fill="currentColor" d="M14 3H2c0 1.1-.9 2-2 2v6c1.1 0 2 .9 2 2h12c0-1.1.9-2 2-2V5c-1.1 0-2-.9-2-2zm-1 9H3V4h10v8z"/><path fill="currentColor" d="M4 5h8v6H4V5z"/>`
	timeBackwardInnerSVG         = `<path fill="currentColor" d="M8 4H7v5h4V8H8z"/><path fill="currentColor" d="M8 0C5 0 2.4 1.6 1.1 4.1L0 3v4h4L2.5 5.5C3.5 3.5 5.6 2 8 2c3.3 0 6 2.7 6 6s-2.7 6-6 6c-1.8 0-3.4-.8-4.5-2.1L2 13.2C3.4 14.9 5.6 16 8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8z"/>`
	timeForwardInnerSVG          = `<path fill="currentColor" d="M8 4H7v5h4V8H8z"/><path fill="currentColor" d="M16 7V3l-1.1 1.1C13.6 1.6 11 0 8 0C3.6 0 0 3.6 0 8s3.6 8 8 8c2.4 0 4.6-1.1 6-2.8l-1.5-1.3C11.4 13.2 9.8 14 8 14c-3.3 0-6-2.7-6-6s2.7-6 6-6c2.4 0 4.5 1.5 5.5 3.5L12 7h4z"/>`
	timerInnerSVG                = `<path fill="currentColor" d="M9.06 9.06c.271-.271.439-.646.439-1.06s-.168-.789-.439-1.06c-.59-.59-6.72-4.6-6.72-4.6s4 6.13 4.59 6.72a1.497 1.497 0 0 0 2.13 0z"/><path fill="currentColor" d="M8 0v3h1V1.59c3.153.495 5.536 3.192 5.536 6.445a6.52 6.52 0 1 1-12.07-3.423L1.55 3.29A7.94 7.94 0 0 0 .017 8a8 8 0 1 0 8-8H8z"/>`
	toolboxInnerSVG              = `<path fill="currentColor" d="M0 8h6v2h4V8h6v6H0z"/><path fill="currentColor" d="M7 7h2v2H7V7zm4-3V2H5v2H0v3h6V6h4v1h6V4h-5zM6 4V3h4v1H6z"/>`
	toolsInnerSVG                = `<path fill="currentColor" d="m10.3 8.2l-.9.9l.9.9l-1.2 1.2l4.3 4.3c.6.6 1.5.6 2.1 0s.6-1.5 0-2.1l-5.2-5.2zm3.9 6.8c-.4 0-.8-.3-.8-.8c0-.4.3-.8.8-.8s.8.3.8.8s-.3.8-.8.8zM3.6 8l.9-.6L6 5.7l.9.9l.9-.9l-.1-.1c.2-.5.3-1 .3-1.6c0-2.2-1.8-4-4-4c-.6 0-1.1.1-1.6.3l2.9 2.9l-2.1 2.1L.3 2.4C.1 2.9 0 3.4 0 4c0 2.1 1.6 3.7 3.6 4z"/><path fill="currentColor" d="m8 10.8l.9-.8l-.9-.9l5.7-5.7l1.2-.4L16 .8l-.7-.7l-2.3 1l-.5 1.2L6.9 8L6 7.1l-.8.9s.8.6-.1 1.5c-.5.5-1.3-.1-2.8 1.4L.2 13s-.6 1 .6 2.2s2.2.6 2.2.6l2.1-2.1c1.4-1.4.9-2.3 1.3-2.7c.9-.9 1.6-.2 1.6-.2zm-3.1-.4l.7.7l-3.8 3.8l-.7-.7z"/>`
	toothInnerSVG                = `<path fill="currentColor" d="M11.3 16c-1.2 0-1.7-3.9-1.7-4.1c-.1-1.3-1-2.1-1.6-2.2c-.6 0-1.4.9-1.6 2.2c0 .2-.5 4.1-1.7 4.1s-1.8-4.4-1.9-4.4c-.2-1.4.1-3.4.2-4c-.4-1.2-1.8-5.6-.5-7C3 .2 3.6 0 4.4 0c.6 0 1.3.1 2 .3c.6.1 1.1.2 1.6.2S9 .4 9.6.3c.7-.2 1.4-.3 2-.3c.8 0 1.4.2 1.8.7c1.3 1.4-.1 5.8-.5 7c.1.5.4 2.5.2 3.9c.1 0-.5 4.4-1.8 4.4zM8 8.7c1.3.1 2.4 1.4 2.6 3.1c.1 1.2.5 2.4.7 2.9c.3-.6.7-2.1.9-3.3c.2-1.4-.2-3.7-.2-3.7v-.2c.7-2.1 1.4-5.3.8-6.1c-.3-.3-.7-.4-1.2-.4s-1.2.1-1.8.3c-.6.1-1.2.2-1.8.2s-1.2-.1-1.8-.2C5.6 1.1 4.9 1 4.4 1s-.9.1-1.1.4c-.7.7 0 4 .8 6.1v.2s-.4 2.3-.2 3.7c.2 1.2.6 2.7.9 3.3c.2-.6.6-1.7.7-2.9c.1-1.6 1.2-3 2.5-3.1z"/>`
	touchInnerSVG                = `<path fill="currentColor" d="M12.62 6a1.312 1.312 0 0 0-.629.002a1.54 1.54 0 0 0-.402-.843a1.417 1.417 0 0 0-1.008-.42l-.053.001h-.007c-.22 0-.43.044-.621.124a1.2 1.2 0 0 0-.29-.464a1.453 1.453 0 0 0-1.115-.399H8.5a2.5 2.5 0 1 0-3.506.486L5 7.151a2.194 2.194 0 0 0-1.432.581C3 8.351 3 9.311 3 10.511v.72a3.617 3.617 0 0 0 1.402 2.764l.358.356c1.24 1.27 2.38 1.65 5.02 1.65c2.88 0 4.22-1.61 4.22-5.06v-2.51c0-.77-.22-2.12-1.38-2.43zM13 8.35v2.59C13 14.31 11.71 15 9.78 15c-2.6 0-3.4-.39-4.3-1.33l-.36-.37A2.652 2.652 0 0 1 4 11.235V10.5a3.347 3.347 0 0 1 .298-2.09c.186-.132.431-.228.698-.24l.003.7v-.22l-.34 1.5a.186.186 0 1 0 .34.151l1-1.211a.214.214 0 0 0 0-.091V3.39l-.001-.039c0-.256.083-.492.223-.684a.51.51 0 0 1 .399-.157h-.001c.21 0 .38.17.38.38v3.88a.5.5 0 0 0 1 0V5.45A.47.47 0 0 1 8.501 5A.42.42 0 0 1 9 5.412l-.001.029V6.77a.5.5 0 0 0 1 0v-.64a.45.45 0 0 1 .502-.379a.431.431 0 0 1 .338.11c.1.129.16.294.16.473v.836a.5.5 0 0 0 .504.494a.51.51 0 0 0 .496-.39a.513.513 0 0 1 .16-.273a.27.27 0 0 1 .223 0A1.38 1.38 0 0 1 13 8.327l.001-.007z"/>`
	trainInnerSVG                = `<path fill="currentColor" d="M13 11.2V3.8c0-1-.8-1.8-1.8-1.8H9V1h2V0H5v1h2v1H4.8C3.8 2 3 2.8 3 3.8v7.4c0 1 .8 1.8 1.8 1.8H5l-.7 1H3v1h.7L3 16h2l.6-1h4.9l.6 1h2l-.7-1h.6v-1h-1.3l-.7-1h.2c1 0 1.8-.8 1.8-1.8zM4 3.9c0-.5.4-.9.9-.9H11c.6 0 1 .4 1 .9V6c0 .6-.4 1-.9 1H4.9c-.5 0-.9-.4-.9-.9V3.9zM4 11c0-.6.4-1 1-1s1 .4 1 1s-.4 1-1 1s-1-.4-1-1zm5.9 3H6.1l.6-1h2.6l.6 1zm.1-3c0-.6.4-1 1-1s1 .4 1 1s-.4 1-1 1s-1-.4-1-1z"/>`
	trashInnerSVG                = `<path fill="currentColor" d="M13 3s0-.51-2-.8v-.7A1.53 1.53 0 0 0 9.47 0h-3A1.5 1.5 0 0 0 5 1.5v.7a3.706 3.706 0 0 0-2.007.806L2 3v1h12V3h-1zM6 1.5a.51.51 0 0 1 .499-.5h3.002a.53.53 0 0 1 .529.499v.561a10.224 10.224 0 0 0-1.527-.059c-.553 0-2.063 0-2.503.07v-.57zM2 5v1h1v9c1.234.631 2.692 1 4.236 1h1.529a9.418 9.418 0 0 0 4.289-1.025L13 6h1V5H2zm4 8.92q-.51-.06-1-.17V7h1v6.92zM9 14H7V7h2v7zm2-.28c-.267.07-.606.136-.95.184L10 7h1v6.72z"/>`
	treeTableInnerSVG            = `<path fill="currentColor" d="M6 10V8H4V7h1V5H2v2h1v6h3v-2H4v-1z"/><path fill="currentColor" d="M0 0v16h16V0H0zm7 15H1V3h6v12zm4 0H8V3h3v12zm4 0h-3V3h3v12z"/>`
	trendindDownInnerSVG         = `<path fill="currentColor" d="M16 14h-4l1.29-1.29L9 8.41l-3 3l-6-6V2.59l6 6l3-3l5.71 5.7L15.99 10l.01 4z"/>`
	trendingUpInnerSVG           = `<path fill="currentColor" d="M16 2h-4l1.29 1.29L9 7.59l-3-3l-6 6v2.82l6-6l3 3l5.71-5.7L15.99 6L16 2z"/>`
	trophyInnerSVG               = `<path fill="currentColor" d="M11.7 8c4.2-.3 4.3-2.7 4.3-5h-3V0H3v3H0c0 2.3.1 4.7 4.3 5c.9 1.4 2.1 2 2.7 2v4c-3 0-3 2-3 2h8s0-2-3-2v-4c.6 0 1.8-.6 2.7-2zM13 4h2c-.1 1.6-.4 2.7-2.7 2.9c.3-.8.6-1.7.7-2.9zM1 4h2c.1 1.2.4 2.1.7 2.9C1.5 6.7 1.1 5.6 1 4zm3.5 2.1C4 4.4 4 3 4 3V1h1v2s0 1.7.4 3.1C5.9 7.8 7 9 7 9s-1.8-.2-2.5-2.9z"/>`
	truckInnerSVG                = `<path fill="currentColor" d="M6 3h10v7H6V3zm9 11a2 2 0 1 1-3.999.001A2 2 0 0 1 15 14zm-2-3c1.3 0 2.4.8 2.8 2h.2v-2h-3z"/><path fill="currentColor" d="M5 5H1L0 9v4h1.2c.4-1.2 1.5-2 2.8-2s2.4.8 2.8 2h3.4c.4-1.2 1.5-2 2.8-2H5V5zM4 9H1l.8-3H4v3z"/><path fill="currentColor" d="M6 14a2 2 0 1 1-3.999.001A2 2 0 0 1 6 14z"/>`
	twinColSelectInnerSVG        = `<path fill="currentColor" d="M0 2v12h16V2H0zm7 11H1V3h6v10zm8 0H9V3h6v10z"/><path fill="currentColor" d="M10 4h4v1h-4V4zM2 4h4v1H2V4zm0 2h4v1H2V6zm0 2h4v1H2V8z"/>`
	twitterInnerSVG              = `<path fill="currentColor" d="M16 3c-.6.3-1.2.4-1.9.5c.7-.4 1.2-1 1.4-1.8c-.6.4-1.3.6-2.1.8c-.6-.6-1.5-1-2.4-1c-1.7 0-3.2 1.5-3.2 3.3c0 .3 0 .5.1.7c-2.7-.1-5.2-1.4-6.8-3.4c-.3.5-.4 1-.4 1.7c0 1.1.6 2.1 1.5 2.7c-.5 0-1-.2-1.5-.4C.7 7.7 1.8 9 3.3 9.3c-.3.1-.6.1-.9.1c-.2 0-.4 0-.6-.1c.4 1.3 1.6 2.3 3.1 2.3c-1.1.9-2.5 1.4-4.1 1.4H0c1.5.9 3.2 1.5 5 1.5c6 0 9.3-5 9.3-9.3v-.4C15 4.3 15.6 3.7 16 3z"/>`
	twitterSquareInnerSVG        = `<path fill="currentColor" d="M0 0v16h16V0H0zm12.8 5.6v.3c0 3.3-2.5 7-7 7c-1.4 0-2.7-.4-3.8-1.1h.6c1.2 0 2.2-.4 3.1-1.1c-1.1 0-2-.7-2.3-1.7h.5c.2 0 .4 0 .6-.1c-1.1-.2-2-1.2-2-2.4c.3.2.7.3 1.1.3c-.7-.4-1.1-1.2-1.1-2c0-.5.1-.9.3-1.2C4 5.1 5.9 6 7.9 6.1c0-.2-.1-.4-.1-.6C7.8 4.1 8.9 3 10.3 3c.7 0 1.3.3 1.8.8c.6-.1 1.1-.3 1.6-.6c-.2.6-.6 1.1-1.1 1.4c.5-.1 1-.2 1.4-.4c-.3.6-.7 1-1.2 1.4z"/>`
	umbrellaInnerSVG             = `<path fill="currentColor" d="M5.36.9L5.09.33a.54.54 0 0 0-.773-.259a.55.55 0 0 0-.316.762l.319.577C-1.88 4.9.42 12 .42 12h.06c.25-.12.8-1.64 2.05-2.25s2.78-.09 3-.21l.12-.06a11.75 11.75 0 0 1 1.58-1.97l3.37 7.07a2.364 2.364 0 0 0 1.334 1.335a1.764 1.764 0 0 0 1.237-.069l.359-.176c.263-.145.462-.38.558-.662a2.34 2.34 0 0 0 .183-.913c0-.401-.1-.778-.277-1.108a.628.628 0 0 0-.768-.286a.55.55 0 0 0-.244.752c.041.077.529 1.067-.101 1.337s-1.19-.73-1.19-.73L8.271 7a11.482 11.482 0 0 1 2.532.005l.068-.065c.25-.12.8-1.64 2.05-2.25s2.78-.09 3-.21h.06S12.001-1.93 5.361.9zm2 5.46a3.81 3.81 0 0 0-2.211 2.224c-.55-1.082-.909-2.375-1.007-3.74C4 2.6 4.75 1.92 4.75 1.92l.77-.32a4.989 4.989 0 0 1 2.816 1.195a8.689 8.689 0 0 1 2.233 3.265c-.339-.021-.752-.067-1.175-.067c-.724 0-1.417.134-2.054.379z"/>`
	underlineInnerSVG            = `<path fill="currentColor" d="M2 15h12v1H2v-1zm9-15v8.4C11 9.9 9.9 11 8.4 11h-.8C6.1 11 5 9.9 5 8.4V0H2v8.4C2 11.5 4.5 14 7.6 14h.9c3.1 0 5.6-2.5 5.6-5.6V0H11z"/>`
	unlinkInnerSVG               = `<path fill="currentColor" d="M8 0h1v4H8V0zm0 12h1v4H8v-4zM7 9H3a1 1 0 0 1 0-2h4V5H3a3 3 0 1 0 0 6h4V9zm6-4H9v2h4a1 1 0 0 1 0 2H9v2h4a3 3 0 1 0 0-6zM4.51 15.44L7 12H5.77l-2.08 2.88l.82.56zm7.98 0L10 12h1.23l2.08 2.88l-.82.56zm0-14.45L10 4h1.23l2.08-2.66l-.82-.35zm-7.98 0L7 4H5.77L3.69 1.34l.82-.35z"/>`
	unlockInnerSVG               = `<path fill="currentColor" d="M8 8V4.9C8 2.7 6.2 1 4.1 1h-.3C1.6 1 0 2.7 0 4.9V7h2V4.9C2 3.8 2.7 3 3.8 3h.3c1 0 1.9.8 1.9 1.9V8H5l.1 5S5 16 10 16s5-3 5-3V8H8zm3 6h-1v-1.8c-.6 0-1-.6-1-1.1c0-.6.4-1.1 1-1.1s1 .4 1 .9V14z"/>`
	uploadInnerSVG               = `<path fill="currentColor" d="M11 10v2H5v-2H0v6h16v-6h-5zm-7 4H2v-2h2v2z"/><path fill="currentColor" d="M13 5L8 0L3 5h3v6h4V5z"/>`
	uploadAltInnerSVG            = `<path fill="currentColor" d="M0 14h16v2H0v-2zM8 0L3 5h3v8h4V5h3z"/>`
	userInnerSVG                 = `<path fill="currentColor" d="M8 0C2.4 0 5.1 7.3 5.1 7.3c.6 1 1.4.8 1.4 1.5c0 .6-.7.8-1.4.9C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h13.7s-.3-3.6-.8-4.7c-1-1.9-2-1.6-3.1-1.7c-.7-.1-1.4-.3-1.4-.9s.8-.4 1.4-1.5C10.9 7.3 13.6 0 8 0z"/>`
	userCardInnerSVG             = `<path fill="currentColor" d="M15 3v10H1V3h14zm1-1H0v12h16V2z"/><path fill="currentColor" d="M8 5h6v1H8V5zm0 2h6v1H8V7zm0 2h3v1H8V9zM5.4 7H5v-.1c.6-.2 1-.8 1-1.4C6 4.7 5.3 4 4.5 4S3 4.7 3 5.5c0 .7.4 1.2 1 1.4V7h-.4C2.7 7 2 7.7 2 8.6V11h5V8.6C7 7.7 6.3 7 5.4 7z"/>`
	userCheckInnerSVG            = `<path fill="currentColor" d="M7.5 14.4c-.8-.8-.8-2 0-2.8s2-.8 2.8 0l.6.6l1.9-2.1c-.7-.4-1.3-.4-2-.4c-.7-.1-1.4-.3-1.4-.9s.8-.4 1.4-1.5c0 0 2.7-7.3-2.9-7.3c-5.5 0-2.8 7.3-2.8 7.3c.6 1 1.4.8 1.4 1.5s-.7.7-1.4.8C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h8l-1.6-1.6zm5.3 1.6h2.1s-.1-.9-.2-2l-1.9 2z"/><path fill="currentColor" d="M11 16c-.3 0-.5-.1-.7-.3l-2-2c-.4-.4-.4-1 0-1.4s1-.4 1.4 0l1.3 1.3l3.3-3.6c.4-.4 1-.4 1.4-.1c.4.4.4 1 .1 1.4l-4 4.3c-.3.3-.5.4-.8.4z"/>`
	userClockInnerSVG            = `<path fill="currentColor" d="M14 13h-3v-3h1v2h2z"/><path fill="currentColor" d="M16 12.5C16 10 14 8 11.5 8c-.7 0-1.4.2-2 .5c.2-.3.8-.3 1.4-1.2c0 0 2.7-7.3-2.9-7.3S5.1 7.3 5.1 7.3c.6 1 1.4.8 1.4 1.5s-.7.7-1.4.8C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h10.4C9.6 16 8 14.4 8 12.5S9.6 9 11.5 9s3.5 1.6 3.5 3.5s-1.6 3.5-3.5 3.5h3.4v-.5c.6-.8 1.1-1.8 1.1-3z"/>`
	userHeartInnerSVG            = `<path fill="currentColor" d="M14.2 16h.6v-.6l-.6.6zm-5.6-2.1c-.7-.7-1-1.8-.8-2.8S8.6 9.3 9.5 9c0-.1-.1-.2-.1-.2c0-.6.8-.4 1.4-1.5c0 0 2.7-7.3-2.9-7.3c-5.5 0-2.8 7.3-2.8 7.3c.6 1 1.4.8 1.4 1.5s-.7.7-1.4.8C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h9.6l-2.1-2.1z"/><path fill="currentColor" d="M14.9 10.1c-.2-.1-.5-.1-.7-.1c-.7 0-1.3.6-1.7 1.1c-.4-.5-1-1.1-1.7-1.1c-.3 0-.5 0-.7.1c-1.2.4-1.4 2-.5 2.9l3 2.9l3-2.9c.8-.9.5-2.5-.7-2.9z"/>`
	userStarInnerSVG             = `<path fill="currentColor" d="m8.92 13.67l-1.61-1.53l-1.5-1.42l2-.29l2.25-.32l.29-.57h-.02a1 1 0 0 1-.979-.794c-.001-.617.799-.417 1.429-1.457c.08-.02 2.82-7.29-2.78-7.29s-2.86 7.27-2.86 7.27c.63 1 1.44.85 1.43 1.45s-.74.8-1.43.87C4 9.719 3 9.459 2 11.349c-.6 1.09-.85 4.65-.85 4.65h7.36v-.17zm2.8 2.33h.56l-.28-.14l-.28.14z"/><path fill="currentColor" d="M12 14.73L14.47 16L14 13.31l2-1.9l-2.76-.39L12 8.57l-1.24 2.45l-2.76.39l2 1.9L9.53 16L12 14.73z"/>`
	usersInnerSVG                = `<path fill="currentColor" d="M5.3 9.7c-.4 0-.9-.2-.9-.6s.5-.3.9-1c0 0 1.8-4.9-1.8-4.9S1.7 8.1 1.7 8.1c.4.7.9.6.9 1s-.5.6-.9.6c-.6.1-1.1 0-1.7.6V16h5c.2-1.7.7-5.2 1.1-6.1l.1-.1c-.2-.1-.5-.1-.9-.1zM16 9.5c-.7-.8-1.3-.7-2-.8c-.5-.1-1.1-.2-1.1-.7s.6-.3 1.1-1.2c0 0 2.1-5.9-2.2-5.9c-4.4.1-2.3 6-2.3 6c.5.8 1.1.7 1.1 1.1c0 .5-.6.6-1.1.7c-.9.1-1.7 0-2.5 1.5c-.4.9-1 5.8-1 5.8h10V9.5z"/>`
	vaadinHInnerSVG              = `<path fill="currentColor" d="M15.21.35a.79.79 0 0 0-.79.79v.46c0 .5-.32.85-1.07.85H9.8c-1.61 0-1.73 1.19-1.8 1.83c-.06-.64-.18-1.83-1.79-1.83H2.64c-.75 0-1.09-.37-1.09-.86v-.47A.77.77 0 0 0 .78.35a.78.78 0 0 0-.78.78v.011v-.001v1.32C0 4 .7 4.77 2.34 4.77H6c1.09 0 1.19.46 1.19.9v.13a.851.851 0 0 0 1.69.004V5.67c0-.44.1-.9 1.19-.9h3.61C15.29 4.77 16 4 16 2.46V1.14a.79.79 0 0 0-.79-.79zm-4 7.03l-.04-.001a1 1 0 0 0-.958.714l-.002.007L8 12.31l-2.3-4.2a1.003 1.003 0 0 0-.963-.731l-.039.001H4.7l-.039-.001a1.05 1.05 0 0 0-.879 1.625L3.78 9l3.29 6.1a.942.942 0 0 0 1.718.006l.002-.006L12.13 9a1.05 1.05 0 0 0-.906-1.58h-.014h.001z"/>`
	vaadinVInnerSVG              = `<path fill="currentColor" d="M5.8 7.16h-.13c-.44 0-.9-.1-.9-1.19V2.35C4.77.71 4 0 2.46 0H1.14a.79.79 0 0 0 0 1.58h.46c.5 0 .85.32.85 1.07V6.2c0 1.61 1.19 1.73 1.83 1.8c-.64.06-1.83.18-1.83 1.79v3.55c0 .75-.37 1.09-.86 1.09h-.47a.77.77 0 0 0-.77.77c0 .431.349.78.78.78h1.33c1.54 0 2.31-.7 2.31-2.34v-3.59c0-1.09.46-1.19.9-1.19h.13a.851.851 0 0 0 .004-1.69H5.8zm9.3.03L9 3.87a1.05 1.05 0 0 0-1.58.906v.014v-.001l-.001.04a1 1 0 0 0 .714.958l.007.002l4.21 2.26l-4.24 2.25a1.003 1.003 0 0 0-.73 1.002v-.002l-.001.039a1.05 1.05 0 0 0 1.625.879L9 12.219l6.1-3.29a.942.942 0 0 0 .006-1.718l-.006-.002z"/>`
	viewportInnerSVG             = `<path fill="currentColor" d="M1 4H0V0h4v1H1zm11-3V0h4v4h-1V1zm3 11h1v4h-4v-1h3zM4 15v1H0v-4h1v3z"/><path fill="currentColor" d="M13 3v10H3V3h10zm1-1H2v12h12V2z"/>`
	vimeoInnerSVG                = `<path fill="currentColor" d="M15.9 4.4c-.9 5-5.9 9.3-7.4 10.3s-2.9-.4-3.4-1.4C4.6 12 2.9 5.7 2.4 5.1C2 4.6.6 5.7.6 5.7L0 4.8s2.7-3.3 4.8-3.7C7 .7 7 4.5 7.5 6.6c.5 2 .9 3.2 1.3 3.2s1.3-1.1 2.2-2.9c.9-1.7 0-3.3-1.9-2.2c.8-4.3 7.7-5.4 6.8-.3z"/>`
	vimeoSquareInnerSVG          = `<path fill="currentColor" d="M0 0v16h16V0H0zm13.9 5.3c-.7 3.8-4.4 7-5.5 7.7s-2.2-.3-2.5-1.1c-.4-.9-1.7-5.7-2-6.1c-.4-.3-1.4.5-1.4.5L2 5.6s2-2.4 3.6-2.7s1.6 2.5 2 4.1c.4 1.5.6 2.4 1 2.4c.3 0 1-.9 1.7-2.2s0-2.5-1.4-1.6c.5-3.3 5.7-4.1 5-.3z"/>`
	volumeInnerSVG               = `<path fill="currentColor" d="m11.8 2.4l-.5 1C12.4 4.8 13 6.6 13 8.5c0 1.7-.5 3.2-1.3 4.6l.7.8c1.1-1.5 1.7-3.4 1.7-5.4c-.1-2.3-.9-4.4-2.3-6.1z"/><path fill="currentColor" d="m10.8 4.4l-.5 1.1c.5.9.8 1.9.8 3c0 1-.3 2-.7 2.9l.7.9c.6-1.1 1-2.4 1-3.7c-.1-1.6-.5-3-1.3-4.2zM4 5H0v6h4l5 4V1z"/>`
	volumeDownInnerSVG           = `<path fill="currentColor" d="m10.8 4.4l-.5 1.1c.5.9.8 1.9.8 3c0 1-.3 2-.7 2.9l.7.9c.6-1.1 1-2.4 1-3.7c-.1-1.6-.5-3-1.3-4.2zM4 5H0v6h4l5 4V1z"/>`
	volumeOffInnerSVG            = `<path fill="currentColor" d="M4 5H0v6h4l5 4V1z"/>`
	volumeUpInnerSVG             = `<path fill="currentColor" d="M15 8.5c0 2.3-.8 4.5-2 6.2l.7.8c1.5-1.9 2.4-4.4 2.4-7c0-3.1-1.2-5.9-3.2-8l-.5 1C14 3.3 15 5.8 15 8.5z"/><path fill="currentColor" d="m11.8 2.4l-.5 1C12.4 4.8 13 6.6 13 8.5c0 1.7-.5 3.2-1.3 4.6l.7.8c1.1-1.5 1.7-3.4 1.7-5.4c-.1-2.3-.9-4.4-2.3-6.1z"/><path fill="currentColor" d="m10.8 4.4l-.5 1.1c.5.9.8 1.9.8 3c0 1-.3 2-.7 2.9l.7.9c.6-1.1 1-2.4 1-3.7c-.1-1.6-.5-3-1.3-4.2zM4 5H0v6h4l5 4V1z"/>`
	walletInnerSVG               = `<path fill="currentColor" d="M14.5 4H2.38a1 1 0 0 1-1.19-.982v-.019L14 2.5V1.31A1.18 1.18 0 0 0 12.684.001L1.31 1.85A2.004 2.004 0 0 0 0 3.727v10.772a1.5 1.5 0 0 0 1.5 1.5h13a1.5 1.5 0 0 0 1.5-1.5v-9.01l.001-.041a1.45 1.45 0 0 0-1.45-1.45l-.053.001zM13 11a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 13 11z"/>`
	warningInnerSVG              = `<path fill="currentColor" d="M8 1L0 15h16L8 1zm0 12c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1zm-1-3V6h2v4H7z"/>`
	workplaceInnerSVG            = `<path fill="currentColor" d="M11 3V0H2v14H0v1h7v-5h2V8h5V3h-3zm-5 7H4V8h2v2zm0-3H4V5h2v2zm0-3H4V2h2v2zm3 3H7V5h2v2zm0-3H7V2h2v2zm4 3h-2V5h2v2zm1 4h2v5H8v-5h2V9h4v2z"/>`
	wrenchInnerSVG               = `<path fill="currentColor" d="M15.5 13.4L7.7 5.6c.2-.5.3-1 .3-1.6c0-2.2-1.8-4-4-4c-.6 0-1.1.1-1.6.3l2.9 2.9l-2.1 2.1L.3 2.4C.1 2.9 0 3.4 0 4c0 2.2 1.8 4 4 4c.6 0 1.1-.1 1.6-.3l7.8 7.8c.6.6 1.5.6 2.1 0s.6-1.5 0-2.1zM6.8 7.6L5.4 6.2l.9-.9l1.4 1.4l-.9.9zm7.4 7.4c-.4 0-.8-.3-.8-.8c0-.4.3-.8.8-.8s.8.3.8.8s-.3.8-.8.8z"/>`
	youtubeInnerSVG              = `<path fill="currentColor" d="M6.6 0h-.9l-.6 2.3L4.5 0h-1c.2.6.4 1.1.6 1.7c.3.8.5 1.5.5 1.9V6h.9V3.6L6.6 0zM9 4.5V3c0-.5-.1-.8-.3-1.1s-.5-.4-.9-.4s-.7.2-.9.5c-.2.2-.3.5-.3 1v1.6c0 .5.1.8.3 1c.2.3.5.4.9.4s.7-.2.9-.5c.2-.1.3-.5.3-1zm-.8.2c0 .4-.1.6-.4.6s-.4-.2-.4-.6V2.8c0-.4.1-.6.4-.6s.4.2.4.6v1.9zM12 6V1.5h-.8v3.4c-.2.3-.3.4-.5.4c-.1 0-.2-.1-.2-.2V1.5h-.8V5c0 .3 0 .5.1.7c0 .2.2.3.5.3s.6-.2.9-.5V6h.8zm.4 4.5c-.3 0-.4.2-.4.6v.4h.8v-.4c0-.4-.1-.6-.4-.6zm-2.9 0c-.1 0-.3.1-.4.2v2.7c.1.1.3.2.4.2c.2 0 .3-.2.3-.6v-1.9c0-.4-.1-.6-.3-.6z"/><path fill="currentColor" d="M14.4 8.3C14.2 7.6 13.6 7 13 7c-1.6-.2-3.3-.2-5-.2s-3.3 0-5 .2c-.6 0-1.2.6-1.4 1.3c-.2 1-.2 2.1-.2 3.1s0 2.1.2 3.1c.2.7.7 1.2 1.4 1.3c1.7.2 3.3.2 5 .2s3.3 0 5-.2c.7-.1 1.3-.6 1.4-1.3c.2-1 .2-2.1.2-3.1s0-2.1-.2-3.1zm-9.2.9h-1v5.1h-.9V9.2h-.9v-.9h2.8v.9zm2.4 5.1h-.8v-.5c-.3.4-.6.5-.9.5s-.4-.1-.5-.3c0-.1-.1-.3-.1-.7V9.8h.8v3.5c0 .1.1.2.2.2c.2 0 .3-.1.5-.4V9.8h.8v4.5zm3-1.4c0 .4 0 .7-.1.9c-.1.3-.3.5-.6.5s-.6-.2-.8-.5v.4h-.8V8.3h.8v1.9c.3-.3.5-.5.8-.5s.5.2.6.5c.1.2.1.5.1.9v1.8zm3-.7H12v.8c0 .4.1.6.4.6c.2 0 .3-.1.4-.3v-.5h.8v.6c0 .2-.1.3-.2.5c-.2.3-.5.5-1 .5c-.4 0-.7-.2-1-.5c-.2-.2-.3-.6-.3-1v-1.5c0-.5.1-.8.2-1c.2-.3.5-.5 1-.5c.4 0 .7.2.9.5c.2.2.2.6.2 1v.8z"/>`
	youtubeSquareInnerSVG        = `<path fill="currentColor" d="M7.9 6c.2 0 .3-.2.3-.5V4.1c0-.3-.1-.5-.3-.5s-.3.2-.3.5v1.4c0 .3.1.5.3.5zm-.8 5.9c-.1.2-.3.3-.4.3s-.1 0-.1-.1V9.4H6V12c0 .2 0 .4.1.5c.1.2.2.2.4.2s.4-.1.7-.4v.4h.6V9.4h-.7v2.5zm-3.3-3h.7v3.8h.7V8.9h.7v-.7H3.8zm5.6.4c-.2 0-.4.2-.6.4V8.2h-.6v4.4h.6v-.3c.2.2.4.4.6.4s.4-.1.5-.4c0-.1.1-.4.1-.7v-1.3c0-.3 0-.5-.1-.7c-.1-.1-.2-.3-.5-.3zm0 2.4c0 .3-.1.4-.3.4c-.1 0-.2 0-.3-.1v-2c.1-.1.2-.1.3-.1c.2 0 .3.2.3.5v1.3zm1.9-2.4c-.3 0-.5.1-.7.3c-.1.2-.2.4-.2.8v1.2c0 .4.1.6.2.8c.2.2.4.3.7.3s.6-.1.7-.4c.1-.1.1-.2.1-.4v-.5h-.6v.4c0 .2-.1.2-.3.2s-.3-.2-.3-.5v-.6h1.2v-.7c0-.4-.1-.6-.2-.8c0 .1-.3-.1-.6-.1zm.3 1.3H11v-.3c0-.3.1-.5.3-.5s.3.2.3.5v.3z"/><path fill="currentColor" d="M0 0v16h16V0H0zm9.3 3.1h.6v2.7c0 .1 0 .2.1.2s.2-.1.4-.3V3.1h.6v3.3h-.6v-.3c-.2.3-.5.4-.7.4s-.3-.1-.4-.2c0-.1-.1-.3-.1-.5V3.1zM7 4.2c0-.3 0-.6.2-.8s.4-.3.7-.3c.3 0 .5.1.7.3c.1.2.2.4.2.8v1.2c0 .4-.1.6-.2.8c-.2.2-.4.3-.7.3s-.5-.1-.7-.3C7 6 7 5.8 7 5.4V4.2zM5.3 2l.5 1.8l.5-1.8H7l-.8 2.7v1.8h-.7V4.7c-.1-.4-.2-.8-.4-1.5c-.2-.4-.3-.8-.5-1.2h.7zm7.5 10.9c-.1.5-.6.9-1.1 1c-1.2.1-2.5.1-3.7.1s-2.5 0-3.7-.1c-.5-.1-1-.4-1.1-1c-.2-.8-.2-1.6-.2-2.4c0-.7 0-1.5.2-2.3c.1-.5.6-.9 1.1-1c1.2-.1 2.5-.1 3.7-.1s2.5 0 3.7.1c.5.1 1 .4 1.1 1c.2.8.2 1.6.2 2.3c0 .8 0 1.6-.2 2.4z"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Abacus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(abacusInnerSVG).
		Element(children...)
}

func AbsolutePosition(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(absolutePositionInnerSVG).
		Element(children...)
}

func AcademyCap(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(academyCapInnerSVG).
		Element(children...)
}

func Accessibility(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accessibilityInnerSVG).
		Element(children...)
}

func AccordionMenu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accordionMenuInnerSVG).
		Element(children...)
}

func AddDock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addDockInnerSVG).
		Element(children...)
}

func Adjust(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustInnerSVG).
		Element(children...)
}

func AdobeFlash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adobeFlashInnerSVG).
		Element(children...)
}

func Airplane(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplaneInnerSVG).
		Element(children...)
}

func Alarm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alarmInnerSVG).
		Element(children...)
}

func AlignCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignRightInnerSVG).
		Element(children...)
}

func Alt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(altInnerSVG).
		Element(children...)
}

func AltA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(altAInnerSVG).
		Element(children...)
}

func Ambulance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ambulanceInnerSVG).
		Element(children...)
}

func Anchor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(anchorInnerSVG).
		Element(children...)
}

func AngleDoubleDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDoubleDownInnerSVG).
		Element(children...)
}

func AngleDoubleLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDoubleLeftInnerSVG).
		Element(children...)
}

func AngleDoubleRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDoubleRightInnerSVG).
		Element(children...)
}

func AngleDoubleUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDoubleUpInnerSVG).
		Element(children...)
}

func AngleDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDownInnerSVG).
		Element(children...)
}

func AngleLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleLeftInnerSVG).
		Element(children...)
}

func AngleRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleRightInnerSVG).
		Element(children...)
}

func AngleUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleUpInnerSVG).
		Element(children...)
}

func Archive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archiveInnerSVG).
		Element(children...)
}

func Archives(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archivesInnerSVG).
		Element(children...)
}

func AreaSelect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(areaSelectInnerSVG).
		Element(children...)
}

func ArrowBackward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBackwardInnerSVG).
		Element(children...)
}

func ArrowCircleDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleDownInnerSVG).
		Element(children...)
}

func ArrowCircleDownO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleDownOInnerSVG).
		Element(children...)
}

func ArrowCircleLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleLeftInnerSVG).
		Element(children...)
}

func ArrowCircleLeftO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleLeftOInnerSVG).
		Element(children...)
}

func ArrowCircleRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleRightInnerSVG).
		Element(children...)
}

func ArrowCircleRightO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleRightOInnerSVG).
		Element(children...)
}

func ArrowCircleUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleUpInnerSVG).
		Element(children...)
}

func ArrowCircleUpO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleUpOInnerSVG).
		Element(children...)
}

func ArrowDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownInnerSVG).
		Element(children...)
}

func ArrowForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowForwardInnerSVG).
		Element(children...)
}

func ArrowLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftInnerSVG).
		Element(children...)
}

func ArrowLongDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLongDownInnerSVG).
		Element(children...)
}

func ArrowLongLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLongLeftInnerSVG).
		Element(children...)
}

func ArrowRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightInnerSVG).
		Element(children...)
}

func ArrowUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpInnerSVG).
		Element(children...)
}

func Arrows(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsInnerSVG).
		Element(children...)
}

func ArrowsCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsCrossInnerSVG).
		Element(children...)
}

func ArrowsLongH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsLongHInnerSVG).
		Element(children...)
}

func ArrowsLongRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsLongRightInnerSVG).
		Element(children...)
}

func ArrowsLongUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsLongUpInnerSVG).
		Element(children...)
}

func ArrowsLongV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowsLongVInnerSVG).
		Element(children...)
}

func Asterisk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(asteriskInnerSVG).
		Element(children...)
}

func At(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atInnerSVG).
		Element(children...)
}

func Automation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(automationInnerSVG).
		Element(children...)
}

func Backspace(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backspaceInnerSVG).
		Element(children...)
}

func BackspaceA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backspaceAInnerSVG).
		Element(children...)
}

func Backwards(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backwardsInnerSVG).
		Element(children...)
}

func Ban(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartInnerSVG).
		Element(children...)
}

func BarChartH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartHInnerSVG).
		Element(children...)
}

func BarChartV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartVInnerSVG).
		Element(children...)
}

func Barcode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barcodeInnerSVG).
		Element(children...)
}

func Bed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bedInnerSVG).
		Element(children...)
}

func Bell(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellInnerSVG).
		Element(children...)
}

func BellO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellOInnerSVG).
		Element(children...)
}

func BellSlash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellSlashInnerSVG).
		Element(children...)
}

func BellSlashO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellSlashOInnerSVG).
		Element(children...)
}

func Boat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boatInnerSVG).
		Element(children...)
}

func Bold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boltInnerSVG).
		Element(children...)
}

func Bomb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bombInnerSVG).
		Element(children...)
}

func Book(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookInnerSVG).
		Element(children...)
}

func BookDollar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookDollarInnerSVG).
		Element(children...)
}

func BookPercent(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookPercentInnerSVG).
		Element(children...)
}

func Bookmark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkInnerSVG).
		Element(children...)
}

func BookmarkO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkOInnerSVG).
		Element(children...)
}

func Briefcase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseInnerSVG).
		Element(children...)
}

func Browser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(browserInnerSVG).
		Element(children...)
}

func Bug(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bugInnerSVG).
		Element(children...)
}

func BugO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bugOInnerSVG).
		Element(children...)
}

func Building(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buildingInnerSVG).
		Element(children...)
}

func BuildingO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buildingOInnerSVG).
		Element(children...)
}

func Bullets(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bulletsInnerSVG).
		Element(children...)
}

func Bullseye(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bullseyeInnerSVG).
		Element(children...)
}

func Buss(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bussInnerSVG).
		Element(children...)
}

func Button(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buttonInnerSVG).
		Element(children...)
}

func Calc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calcInnerSVG).
		Element(children...)
}

func CalcBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calcBookInnerSVG).
		Element(children...)
}

func Calendar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarInnerSVG).
		Element(children...)
}

func CalendarBriefcase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarBriefcaseInnerSVG).
		Element(children...)
}

func CalendarClock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarClockInnerSVG).
		Element(children...)
}

func CalendarEnvelope(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarEnvelopeInnerSVG).
		Element(children...)
}

func CalendarO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarOInnerSVG).
		Element(children...)
}

func CalendarUser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarUserInnerSVG).
		Element(children...)
}

func Camera(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraInnerSVG).
		Element(children...)
}

func Car(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(carInnerSVG).
		Element(children...)
}

func CaretDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretDownInnerSVG).
		Element(children...)
}

func CaretLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretRightInnerSVG).
		Element(children...)
}

func CaretSquareDownO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretSquareDownOInnerSVG).
		Element(children...)
}

func CaretSquareLeftO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretSquareLeftOInnerSVG).
		Element(children...)
}

func CaretSquareRightO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretSquareRightOInnerSVG).
		Element(children...)
}

func CaretSquareUpO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretSquareUpOInnerSVG).
		Element(children...)
}

func CaretUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretUpInnerSVG).
		Element(children...)
}

func Cart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartInnerSVG).
		Element(children...)
}

func CartO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartOInnerSVG).
		Element(children...)
}

func Cash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cashInnerSVG).
		Element(children...)
}

func Chart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartInnerSVG).
		Element(children...)
}

func ChartGrid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartGridInnerSVG).
		Element(children...)
}

func ChartLine(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartLineInnerSVG).
		Element(children...)
}

func ChartThreeD(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartThreeDInnerSVG).
		Element(children...)
}

func ChartTimeline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartTimelineInnerSVG).
		Element(children...)
}

func Chat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkCircleInnerSVG).
		Element(children...)
}

func CheckCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkCircleOInnerSVG).
		Element(children...)
}

func CheckSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkSquareInnerSVG).
		Element(children...)
}

func CheckSquareO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkSquareOInnerSVG).
		Element(children...)
}

func ChevronCircleDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCircleDownInnerSVG).
		Element(children...)
}

func ChevronCircleDownO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCircleDownOInnerSVG).
		Element(children...)
}

func ChevronCircleLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCircleLeftInnerSVG).
		Element(children...)
}

func ChevronCircleLeftO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCircleLeftOInnerSVG).
		Element(children...)
}

func ChevronCircleRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCircleRightInnerSVG).
		Element(children...)
}

func ChevronCircleRightO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCircleRightOInnerSVG).
		Element(children...)
}

func ChevronCircleUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCircleUpInnerSVG).
		Element(children...)
}

func ChevronCircleUpO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCircleUpOInnerSVG).
		Element(children...)
}

func ChevronDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronDownInnerSVG).
		Element(children...)
}

func ChevronDownSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronDownSmallInnerSVG).
		Element(children...)
}

func ChevronLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronLeftInnerSVG).
		Element(children...)
}

func ChevronLeftSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronLeftSmallInnerSVG).
		Element(children...)
}

func ChevronRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightInnerSVG).
		Element(children...)
}

func ChevronRightSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightSmallInnerSVG).
		Element(children...)
}

func ChevronUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpInnerSVG).
		Element(children...)
}

func ChevronUpSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpSmallInnerSVG).
		Element(children...)
}

func Child(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(childInnerSVG).
		Element(children...)
}

func Circle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleInnerSVG).
		Element(children...)
}

func CircleThin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleThinInnerSVG).
		Element(children...)
}

func Clipboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardInnerSVG).
		Element(children...)
}

func ClipboardCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardCheckInnerSVG).
		Element(children...)
}

func ClipboardCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardCrossInnerSVG).
		Element(children...)
}

func ClipboardHeart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardHeartInnerSVG).
		Element(children...)
}

func ClipboardPulse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardPulseInnerSVG).
		Element(children...)
}

func ClipboardText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardTextInnerSVG).
		Element(children...)
}

func ClipboardUser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardUserInnerSVG).
		Element(children...)
}

func Clock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockInnerSVG).
		Element(children...)
}

func Close(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeInnerSVG).
		Element(children...)
}

func CloseBig(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeBigInnerSVG).
		Element(children...)
}

func CloseCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeCircleInnerSVG).
		Element(children...)
}

func CloseCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeCircleOInnerSVG).
		Element(children...)
}

func CloseSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeSmallInnerSVG).
		Element(children...)
}

func Cloud(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDownloadInnerSVG).
		Element(children...)
}

func CloudDownloadO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDownloadOInnerSVG).
		Element(children...)
}

func CloudO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudOInnerSVG).
		Element(children...)
}

func CloudUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudUploadInnerSVG).
		Element(children...)
}

func CloudUploadO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudUploadOInnerSVG).
		Element(children...)
}

func Cluster(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clusterInnerSVG).
		Element(children...)
}

func Code(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeInnerSVG).
		Element(children...)
}

func Coffee(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coffeeInnerSVG).
		Element(children...)
}

func Cog(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cogInnerSVG).
		Element(children...)
}

func CogO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cogOInnerSVG).
		Element(children...)
}

func Cogs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cogsInnerSVG).
		Element(children...)
}

func CoinPiles(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coinPilesInnerSVG).
		Element(children...)
}

func Coins(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coinsInnerSVG).
		Element(children...)
}

func Combobox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(comboboxInnerSVG).
		Element(children...)
}

func Comment(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentInnerSVG).
		Element(children...)
}

func CommentEllipsis(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentEllipsisInnerSVG).
		Element(children...)
}

func CommentEllipsisO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentEllipsisOInnerSVG).
		Element(children...)
}

func CommentO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentOInnerSVG).
		Element(children...)
}

func Comments(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentsInnerSVG).
		Element(children...)
}

func CommentsO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentsOInnerSVG).
		Element(children...)
}

func Compile(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compileInnerSVG).
		Element(children...)
}

func Compress(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compressInnerSVG).
		Element(children...)
}

func CompressSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compressSquareInnerSVG).
		Element(children...)
}

func Connect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(connectInnerSVG).
		Element(children...)
}

func ConnectO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(connectOInnerSVG).
		Element(children...)
}

func Controller(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(controllerInnerSVG).
		Element(children...)
}

func Copy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(copyInnerSVG).
		Element(children...)
}

func CopyO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(copyOInnerSVG).
		Element(children...)
}

func Copyright(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(copyrightInnerSVG).
		Element(children...)
}

func CornerLowerLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerLowerLeftInnerSVG).
		Element(children...)
}

func CornerLowerRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerLowerRightInnerSVG).
		Element(children...)
}

func CornerUpperLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerUpperLeftInnerSVG).
		Element(children...)
}

func CornerUpperRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerUpperRightInnerSVG).
		Element(children...)
}

func CreditCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropInnerSVG).
		Element(children...)
}

func CrossCutlery(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crossCutleryInnerSVG).
		Element(children...)
}

func Crosshairs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crosshairsInnerSVG).
		Element(children...)
}

func Css(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cssInnerSVG).
		Element(children...)
}

func Ctrl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ctrlInnerSVG).
		Element(children...)
}

func CtrlA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ctrlAInnerSVG).
		Element(children...)
}

func Cube(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cubeInnerSVG).
		Element(children...)
}

func Cubes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cubesInnerSVG).
		Element(children...)
}

func CurlyBrackets(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(curlyBracketsInnerSVG).
		Element(children...)
}

func Cursor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cursorInnerSVG).
		Element(children...)
}

func CursorO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cursorOInnerSVG).
		Element(children...)
}

func Cutlery(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cutleryInnerSVG).
		Element(children...)
}

func Dashboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashboardInnerSVG).
		Element(children...)
}

func Database(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseInnerSVG).
		Element(children...)
}

func DateInput(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dateInputInnerSVG).
		Element(children...)
}

func Deindent(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deindentInnerSVG).
		Element(children...)
}

func Del(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(delInnerSVG).
		Element(children...)
}

func DelA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(delAInnerSVG).
		Element(children...)
}

func DentalChair(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dentalChairInnerSVG).
		Element(children...)
}

func Desktop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(desktopInnerSVG).
		Element(children...)
}

func Diamond(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diamondInnerSVG).
		Element(children...)
}

func DiamondO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diamondOInnerSVG).
		Element(children...)
}

func Diploma(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diplomaInnerSVG).
		Element(children...)
}

func DiplomaScroll(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diplomaScrollInnerSVG).
		Element(children...)
}

func Disc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discInnerSVG).
		Element(children...)
}

func Doctor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doctorInnerSVG).
		Element(children...)
}

func DoctorBriefcase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doctorBriefcaseInnerSVG).
		Element(children...)
}

func Dollar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dollarInnerSVG).
		Element(children...)
}

func DotCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dotCircleInnerSVG).
		Element(children...)
}

func Download(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadInnerSVG).
		Element(children...)
}

func DownloadAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadAltInnerSVG).
		Element(children...)
}

func Drop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropInnerSVG).
		Element(children...)
}

func Edit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editInnerSVG).
		Element(children...)
}

func Eject(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ejectInnerSVG).
		Element(children...)
}

func Elastic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(elasticInnerSVG).
		Element(children...)
}

func EllipsisCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsisCircleInnerSVG).
		Element(children...)
}

func EllipsisCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsisCircleOInnerSVG).
		Element(children...)
}

func EllipsisDotsH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsisDotsHInnerSVG).
		Element(children...)
}

func EllipsisDotsV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsisDotsVInnerSVG).
		Element(children...)
}

func EllipsisH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsisHInnerSVG).
		Element(children...)
}

func EllipsisV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsisVInnerSVG).
		Element(children...)
}

func Enter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enterInnerSVG).
		Element(children...)
}

func EnterArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enterArrowInnerSVG).
		Element(children...)
}

func Envelope(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeInnerSVG).
		Element(children...)
}

func EnvelopeO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeOInnerSVG).
		Element(children...)
}

func EnvelopeOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeOpenInnerSVG).
		Element(children...)
}

func EnvelopeOpenO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeOpenOInnerSVG).
		Element(children...)
}

func Envelopes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopesInnerSVG).
		Element(children...)
}

func EnvelopesO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopesOInnerSVG).
		Element(children...)
}

func Eraser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eraserInnerSVG).
		Element(children...)
}

func Esc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(escInnerSVG).
		Element(children...)
}

func EscA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(escAInnerSVG).
		Element(children...)
}

func Euro(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(euroInnerSVG).
		Element(children...)
}

func Exchange(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exchangeInnerSVG).
		Element(children...)
}

func Exclamation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationInnerSVG).
		Element(children...)
}

func ExclamationCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationCircleInnerSVG).
		Element(children...)
}

func ExclamationCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationCircleOInnerSVG).
		Element(children...)
}

func Exit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exitInnerSVG).
		Element(children...)
}

func ExitO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exitOInnerSVG).
		Element(children...)
}

func Expand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandInnerSVG).
		Element(children...)
}

func ExpandFull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandFullInnerSVG).
		Element(children...)
}

func ExpandSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandSquareInnerSVG).
		Element(children...)
}

func ExternalBrowser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(externalBrowserInnerSVG).
		Element(children...)
}

func ExternalLink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeInnerSVG).
		Element(children...)
}

func EyeSlash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeSlashInnerSVG).
		Element(children...)
}

func Eyedropper(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyedropperInnerSVG).
		Element(children...)
}

func Facebook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookInnerSVG).
		Element(children...)
}

func FacebookSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookSquareInnerSVG).
		Element(children...)
}

func Factory(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(factoryInnerSVG).
		Element(children...)
}

func Family(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(familyInnerSVG).
		Element(children...)
}

func FastBackward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastBackwardInnerSVG).
		Element(children...)
}

func FastForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastForwardInnerSVG).
		Element(children...)
}

func Female(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(femaleInnerSVG).
		Element(children...)
}

func File(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileInnerSVG).
		Element(children...)
}

func FileAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileAddInnerSVG).
		Element(children...)
}

func FileCode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileCodeInnerSVG).
		Element(children...)
}

func FileFont(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileFontInnerSVG).
		Element(children...)
}

func FileMovie(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileMovieInnerSVG).
		Element(children...)
}

func FileO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileOInnerSVG).
		Element(children...)
}

func FilePicture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filePictureInnerSVG).
		Element(children...)
}

func FilePresentation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filePresentationInnerSVG).
		Element(children...)
}

func FileProcess(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileProcessInnerSVG).
		Element(children...)
}

func FileRefresh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileRefreshInnerSVG).
		Element(children...)
}

func FileRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileRemoveInnerSVG).
		Element(children...)
}

func FileSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileSearchInnerSVG).
		Element(children...)
}

func FileSound(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileSoundInnerSVG).
		Element(children...)
}

func FileStart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileStartInnerSVG).
		Element(children...)
}

func FileTable(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTableInnerSVG).
		Element(children...)
}

func FileText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTextInnerSVG).
		Element(children...)
}

func FileTextO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTextOInnerSVG).
		Element(children...)
}

func FileTree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTreeInnerSVG).
		Element(children...)
}

func FileTreeSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTreeSmallInnerSVG).
		Element(children...)
}

func FileTreeSub(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTreeSubInnerSVG).
		Element(children...)
}

func FileZip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileZipInnerSVG).
		Element(children...)
}

func Fill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fillInnerSVG).
		Element(children...)
}

func Film(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterInnerSVG).
		Element(children...)
}

func Fire(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagInnerSVG).
		Element(children...)
}

func FlagCheckered(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagCheckeredInnerSVG).
		Element(children...)
}

func FlagO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagOInnerSVG).
		Element(children...)
}

func Flash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flashInnerSVG).
		Element(children...)
}

func Flask(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flaskInnerSVG).
		Element(children...)
}

func FlightLanding(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flightLandingInnerSVG).
		Element(children...)
}

func FlightTakeoff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flightTakeoffInnerSVG).
		Element(children...)
}

func FlipH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipHInnerSVG).
		Element(children...)
}

func FlipV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipVInnerSVG).
		Element(children...)
}

func Folder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderInnerSVG).
		Element(children...)
}

func FolderAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderAddInnerSVG).
		Element(children...)
}

func FolderO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderOInnerSVG).
		Element(children...)
}

func FolderOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderOpenInnerSVG).
		Element(children...)
}

func FolderOpenO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderOpenOInnerSVG).
		Element(children...)
}

func FolderRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderRemoveInnerSVG).
		Element(children...)
}

func FolderSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderSearchInnerSVG).
		Element(children...)
}

func Font(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fontInnerSVG).
		Element(children...)
}

func Form(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formInnerSVG).
		Element(children...)
}

func Forward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardInnerSVG).
		Element(children...)
}

func FrownO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frownOInnerSVG).
		Element(children...)
}

func Funcion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(funcionInnerSVG).
		Element(children...)
}

func Funnel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(funnelInnerSVG).
		Element(children...)
}

func Gamepad(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadInnerSVG).
		Element(children...)
}

func Gavel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gavelInnerSVG).
		Element(children...)
}

func Gift(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(giftInnerSVG).
		Element(children...)
}

func Glass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(glassInnerSVG).
		Element(children...)
}

func Glasses(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(glassesInnerSVG).
		Element(children...)
}

func Globe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeInnerSVG).
		Element(children...)
}

func GlobeWire(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeWireInnerSVG).
		Element(children...)
}

func Golf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(golfInnerSVG).
		Element(children...)
}

func GooglePlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googlePlusInnerSVG).
		Element(children...)
}

func GooglePlusSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googlePlusSquareInnerSVG).
		Element(children...)
}

func Grab(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(grabInnerSVG).
		Element(children...)
}

func Grid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridInnerSVG).
		Element(children...)
}

func GridBevel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridBevelInnerSVG).
		Element(children...)
}

func GridBig(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridBigInnerSVG).
		Element(children...)
}

func GridBigO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridBigOInnerSVG).
		Element(children...)
}

func GridH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridHInnerSVG).
		Element(children...)
}

func GridSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridSmallInnerSVG).
		Element(children...)
}

func GridSmallO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridSmallOInnerSVG).
		Element(children...)
}

func GridV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridVInnerSVG).
		Element(children...)
}

func Group(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(groupInnerSVG).
		Element(children...)
}

func Hammer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hammerInnerSVG).
		Element(children...)
}

func Hand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(handInnerSVG).
		Element(children...)
}

func HandleCorner(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(handleCornerInnerSVG).
		Element(children...)
}

func HandsUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(handsUpInnerSVG).
		Element(children...)
}

func Handshake(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(handshakeInnerSVG).
		Element(children...)
}

func Harddrive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(harddriveInnerSVG).
		Element(children...)
}

func HarddriveO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(harddriveOInnerSVG).
		Element(children...)
}

func Hash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hashInnerSVG).
		Element(children...)
}

func Header(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headphonesInnerSVG).
		Element(children...)
}

func Headset(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headsetInnerSVG).
		Element(children...)
}

func HealthCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(healthCardInnerSVG).
		Element(children...)
}

func Heart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartInnerSVG).
		Element(children...)
}

func HeartO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartOInnerSVG).
		Element(children...)
}

func Home(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeInnerSVG).
		Element(children...)
}

func HomeO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeOInnerSVG).
		Element(children...)
}

func Hospital(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hospitalInnerSVG).
		Element(children...)
}

func Hourglass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hourglassInnerSVG).
		Element(children...)
}

func HourglassEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hourglassEmptyInnerSVG).
		Element(children...)
}

func HourglassEnd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hourglassEndInnerSVG).
		Element(children...)
}

func HourglassStart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hourglassStartInnerSVG).
		Element(children...)
}

func Inbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxInnerSVG).
		Element(children...)
}

func Indent(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(indentInnerSVG).
		Element(children...)
}

func Info(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoInnerSVG).
		Element(children...)
}

func InfoCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoCircleInnerSVG).
		Element(children...)
}

func InfoCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoCircleOInnerSVG).
		Element(children...)
}

func Input(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inputInnerSVG).
		Element(children...)
}

func Insert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(insertInnerSVG).
		Element(children...)
}

func Institution(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(institutionInnerSVG).
		Element(children...)
}

func Invoice(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(invoiceInnerSVG).
		Element(children...)
}

func Italic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyInnerSVG).
		Element(children...)
}

func KeyO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyOInnerSVG).
		Element(children...)
}

func Keyboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyboardInnerSVG).
		Element(children...)
}

func KeyboardO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyboardOInnerSVG).
		Element(children...)
}

func Laptop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopInnerSVG).
		Element(children...)
}

func Layout(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutInnerSVG).
		Element(children...)
}

func LevelDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(levelDownInnerSVG).
		Element(children...)
}

func LevelDownBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(levelDownBoldInnerSVG).
		Element(children...)
}

func LevelLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(levelLeftInnerSVG).
		Element(children...)
}

func LevelLeftBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(levelLeftBoldInnerSVG).
		Element(children...)
}

func LevelRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(levelRightInnerSVG).
		Element(children...)
}

func LevelRightBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(levelRightBoldInnerSVG).
		Element(children...)
}

func LevelUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(levelUpInnerSVG).
		Element(children...)
}

func LevelUpBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(levelUpBoldInnerSVG).
		Element(children...)
}

func Lifebuoy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lifebuoyInnerSVG).
		Element(children...)
}

func Lightbulb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightbulbInnerSVG).
		Element(children...)
}

func LineBarChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineBarChartInnerSVG).
		Element(children...)
}

func LineChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineChartInnerSVG).
		Element(children...)
}

func LineH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineHInnerSVG).
		Element(children...)
}

func LineV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineVInnerSVG).
		Element(children...)
}

func Lines(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linesInnerSVG).
		Element(children...)
}

func LinesList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linesListInnerSVG).
		Element(children...)
}

func Link(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkInnerSVG).
		Element(children...)
}

func List(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listInnerSVG).
		Element(children...)
}

func ListOl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listOlInnerSVG).
		Element(children...)
}

func ListSelect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listSelectInnerSVG).
		Element(children...)
}

func ListUl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listUlInnerSVG).
		Element(children...)
}

func LocationArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationArrowInnerSVG).
		Element(children...)
}

func LocationArrowCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationArrowCircleInnerSVG).
		Element(children...)
}

func LocationArrowCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationArrowCircleOInnerSVG).
		Element(children...)
}

func Lock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockInnerSVG).
		Element(children...)
}

func Magic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magicInnerSVG).
		Element(children...)
}

func Magnet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magnetInnerSVG).
		Element(children...)
}

func Mailbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailboxInnerSVG).
		Element(children...)
}

func Male(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maleInnerSVG).
		Element(children...)
}

func MapMarker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapMarkerInnerSVG).
		Element(children...)
}

func Margin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(marginInnerSVG).
		Element(children...)
}

func MarginBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(marginBottomInnerSVG).
		Element(children...)
}

func MarginLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(marginLeftInnerSVG).
		Element(children...)
}

func MarginRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(marginRightInnerSVG).
		Element(children...)
}

func MarginTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(marginTopInnerSVG).
		Element(children...)
}

func Medal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(medalInnerSVG).
		Element(children...)
}

func Megafone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(megafoneInnerSVG).
		Element(children...)
}

func MehO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mehOInnerSVG).
		Element(children...)
}

func Menu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircleInnerSVG).
		Element(children...)
}

func MinusCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircleOInnerSVG).
		Element(children...)
}

func MinusSquareO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusSquareOInnerSVG).
		Element(children...)
}

func Mobile(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mobileInnerSVG).
		Element(children...)
}

func MobileBrowser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mobileBrowserInnerSVG).
		Element(children...)
}

func MobileRetro(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mobileRetroInnerSVG).
		Element(children...)
}

func Modal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(modalInnerSVG).
		Element(children...)
}

func ModalList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(modalListInnerSVG).
		Element(children...)
}

func Money(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneyInnerSVG).
		Element(children...)
}

func MoneyDeposit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneyDepositInnerSVG).
		Element(children...)
}

func MoneyExchange(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneyExchangeInnerSVG).
		Element(children...)
}

func MoneyWithdraw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneyWithdrawInnerSVG).
		Element(children...)
}

func Moon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonInnerSVG).
		Element(children...)
}

func MoonO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonOInnerSVG).
		Element(children...)
}

func Morning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(morningInnerSVG).
		Element(children...)
}

func Movie(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(movieInnerSVG).
		Element(children...)
}

func Music(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(musicInnerSVG).
		Element(children...)
}

func Mute(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(muteInnerSVG).
		Element(children...)
}

func NativeButton(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nativeButtonInnerSVG).
		Element(children...)
}

func Newspaper(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(newspaperInnerSVG).
		Element(children...)
}

func Notebook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notebookInnerSVG).
		Element(children...)
}

func Nurse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nurseInnerSVG).
		Element(children...)
}

func Office(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(officeInnerSVG).
		Element(children...)
}

func OpenBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openBookInnerSVG).
		Element(children...)
}

func Option(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(optionInnerSVG).
		Element(children...)
}

func OptionA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(optionAInnerSVG).
		Element(children...)
}

func Options(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(optionsInnerSVG).
		Element(children...)
}

func Orientation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(orientationInnerSVG).
		Element(children...)
}

func Out(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(outInnerSVG).
		Element(children...)
}

func Outbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(outboxInnerSVG).
		Element(children...)
}

func Package(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(packageInnerSVG).
		Element(children...)
}

func Padding(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paddingInnerSVG).
		Element(children...)
}

func PaddingBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paddingBottomInnerSVG).
		Element(children...)
}

func PaddingLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paddingLeftInnerSVG).
		Element(children...)
}

func PaddingRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paddingRightInnerSVG).
		Element(children...)
}

func PaddingTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paddingTopInnerSVG).
		Element(children...)
}

func PaintRoll(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paintRollInnerSVG).
		Element(children...)
}

func Paintbrush(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paintbrushInnerSVG).
		Element(children...)
}

func Palete(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paleteInnerSVG).
		Element(children...)
}

func Panel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panelInnerSVG).
		Element(children...)
}

func Paperclip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperclipInnerSVG).
		Element(children...)
}

func Paperplane(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperplaneInnerSVG).
		Element(children...)
}

func PaperplaneO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperplaneOInnerSVG).
		Element(children...)
}

func Paragraph(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphInnerSVG).
		Element(children...)
}

func Password(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(passwordInnerSVG).
		Element(children...)
}

func Paste(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pasteInnerSVG).
		Element(children...)
}

func Pause(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseInnerSVG).
		Element(children...)
}

func Pencil(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pencilInnerSVG).
		Element(children...)
}

func Phone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneInnerSVG).
		Element(children...)
}

func PhoneLandline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneLandlineInnerSVG).
		Element(children...)
}

func Picture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pictureInnerSVG).
		Element(children...)
}

func PieBarChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieBarChartInnerSVG).
		Element(children...)
}

func PieChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartInnerSVG).
		Element(children...)
}

func PiggyBank(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(piggyBankInnerSVG).
		Element(children...)
}

func PiggyBankCoin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(piggyBankCoinInnerSVG).
		Element(children...)
}

func Pill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pillInnerSVG).
		Element(children...)
}

func Pills(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pillsInnerSVG).
		Element(children...)
}

func Pin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinInnerSVG).
		Element(children...)
}

func PinPost(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinPostInnerSVG).
		Element(children...)
}

func Play(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playCircleInnerSVG).
		Element(children...)
}

func PlayCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playCircleOInnerSVG).
		Element(children...)
}

func Plug(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plugInnerSVG).
		Element(children...)
}

func Plus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircleInnerSVG).
		Element(children...)
}

func PlusCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircleOInnerSVG).
		Element(children...)
}

func PlusMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusMinusInnerSVG).
		Element(children...)
}

func PlusSquareO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusSquareOInnerSVG).
		Element(children...)
}

func Pointer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pointerInnerSVG).
		Element(children...)
}

func PowerOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerOffInnerSVG).
		Element(children...)
}

func Presentation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(presentationInnerSVG).
		Element(children...)
}

func Print(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printInnerSVG).
		Element(children...)
}

func Progressbar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(progressbarInnerSVG).
		Element(children...)
}

func PuzzlePiece(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(puzzlePieceInnerSVG).
		Element(children...)
}

func PyramidChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pyramidChartInnerSVG).
		Element(children...)
}

func Qrcode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(qrcodeInnerSVG).
		Element(children...)
}

func Question(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionInnerSVG).
		Element(children...)
}

func QuestionCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionCircleInnerSVG).
		Element(children...)
}

func QuestionCircleO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionCircleOInnerSVG).
		Element(children...)
}

func QuoteLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quoteLeftInnerSVG).
		Element(children...)
}

func QuoteRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quoteRightInnerSVG).
		Element(children...)
}

func Random(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(randomInnerSVG).
		Element(children...)
}

func Raster(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rasterInnerSVG).
		Element(children...)
}

func RasterLowerLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rasterLowerLeftInnerSVG).
		Element(children...)
}

func Records(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(recordsInnerSVG).
		Element(children...)
}

func Recycle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(recycleInnerSVG).
		Element(children...)
}

func Refresh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshInnerSVG).
		Element(children...)
}

func Reply(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(replyInnerSVG).
		Element(children...)
}

func ReplyAll(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(replyAllInnerSVG).
		Element(children...)
}

func ResizeH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resizeHInnerSVG).
		Element(children...)
}

func ResizeV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resizeVInnerSVG).
		Element(children...)
}

func Retweet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(retweetInnerSVG).
		Element(children...)
}

func Rhombus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rhombusInnerSVG).
		Element(children...)
}

func Road(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(roadInnerSVG).
		Element(children...)
}

func RoadBranch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(roadBranchInnerSVG).
		Element(children...)
}

func RoadBranches(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(roadBranchesInnerSVG).
		Element(children...)
}

func RoadSplit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(roadSplitInnerSVG).
		Element(children...)
}

func Rocket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rocketInnerSVG).
		Element(children...)
}

func RotateLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateLeftInnerSVG).
		Element(children...)
}

func RotateRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateRightInnerSVG).
		Element(children...)
}

func Rss(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssInnerSVG).
		Element(children...)
}

func RssSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssSquareInnerSVG).
		Element(children...)
}

func Safe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(safeInnerSVG).
		Element(children...)
}

func SafeLock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(safeLockInnerSVG).
		Element(children...)
}

func Scale(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scaleInnerSVG).
		Element(children...)
}

func ScaleUnbalance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scaleUnbalanceInnerSVG).
		Element(children...)
}

func ScatterChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scatterChartInnerSVG).
		Element(children...)
}

func Scissors(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scissorsInnerSVG).
		Element(children...)
}

func Screwdriver(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screwdriverInnerSVG).
		Element(children...)
}

func Search(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchInnerSVG).
		Element(children...)
}

func SearchMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchMinusInnerSVG).
		Element(children...)
}

func SearchPlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchPlusInnerSVG).
		Element(children...)
}

func Select(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(selectInnerSVG).
		Element(children...)
}

func Server(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(serverInnerSVG).
		Element(children...)
}

func Share(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareInnerSVG).
		Element(children...)
}

func ShareSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareSquareInnerSVG).
		Element(children...)
}

func Shield(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldInnerSVG).
		Element(children...)
}

func Shift(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shiftInnerSVG).
		Element(children...)
}

func ShiftArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shiftArrowInnerSVG).
		Element(children...)
}

func Shop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shopInnerSVG).
		Element(children...)
}

func SignIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signInInnerSVG).
		Element(children...)
}

func SignInAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signInAltInnerSVG).
		Element(children...)
}

func SignOut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signOutInnerSVG).
		Element(children...)
}

func SignOutAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signOutAltInnerSVG).
		Element(children...)
}

func Signal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signalInnerSVG).
		Element(children...)
}

func Sitemap(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sitemapInnerSVG).
		Element(children...)
}

func Slider(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sliderInnerSVG).
		Element(children...)
}

func Sliders(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slidersInnerSVG).
		Element(children...)
}

func SmileyO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smileyOInnerSVG).
		Element(children...)
}

func Sort(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortInnerSVG).
		Element(children...)
}

func SoundDisable(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundDisableInnerSVG).
		Element(children...)
}

func SparkLine(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sparkLineInnerSVG).
		Element(children...)
}

func Specialist(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(specialistInnerSVG).
		Element(children...)
}

func Spinner(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spinnerInnerSVG).
		Element(children...)
}

func SpinnerArc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spinnerArcInnerSVG).
		Element(children...)
}

func SpinnerThird(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spinnerThirdInnerSVG).
		Element(children...)
}

func SplineAreaChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splineAreaChartInnerSVG).
		Element(children...)
}

func SplineChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splineChartInnerSVG).
		Element(children...)
}

func Split(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splitInnerSVG).
		Element(children...)
}

func SplitH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splitHInnerSVG).
		Element(children...)
}

func SplitV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splitVInnerSVG).
		Element(children...)
}

func Spoon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spoonInnerSVG).
		Element(children...)
}

func SquareShadow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareShadowInnerSVG).
		Element(children...)
}

func Star(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starInnerSVG).
		Element(children...)
}

func StarHalfLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starHalfLeftInnerSVG).
		Element(children...)
}

func StarHalfLeftO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starHalfLeftOInnerSVG).
		Element(children...)
}

func StarHalfRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starHalfRightInnerSVG).
		Element(children...)
}

func StarHalfRightO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starHalfRightOInnerSVG).
		Element(children...)
}

func StarO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starOInnerSVG).
		Element(children...)
}

func StartCog(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(startCogInnerSVG).
		Element(children...)
}

func StepBackward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stepBackwardInnerSVG).
		Element(children...)
}

func StepForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stepForwardInnerSVG).
		Element(children...)
}

func Stethoscope(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stethoscopeInnerSVG).
		Element(children...)
}

func Stock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stockInnerSVG).
		Element(children...)
}

func Stop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopInnerSVG).
		Element(children...)
}

func StopCog(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopCogInnerSVG).
		Element(children...)
}

func Stopwatch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopwatchInnerSVG).
		Element(children...)
}

func Storage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(storageInnerSVG).
		Element(children...)
}

func Strikethrough(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(strikethroughInnerSVG).
		Element(children...)
}

func Subscript(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subscriptInnerSVG).
		Element(children...)
}

func Suitcase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(suitcaseInnerSVG).
		Element(children...)
}

func SunDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunDownInnerSVG).
		Element(children...)
}

func SunO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunOInnerSVG).
		Element(children...)
}

func SunRise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunRiseInnerSVG).
		Element(children...)
}

func Superscript(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(superscriptInnerSVG).
		Element(children...)
}

func Sword(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swordInnerSVG).
		Element(children...)
}

func Tab(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tabInnerSVG).
		Element(children...)
}

func TabA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tabAInnerSVG).
		Element(children...)
}

func Table(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableInnerSVG).
		Element(children...)
}

func Tablet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tabletInnerSVG).
		Element(children...)
}

func Tabs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tabsInnerSVG).
		Element(children...)
}

func Tag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagsInnerSVG).
		Element(children...)
}

func Tasks(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tasksInnerSVG).
		Element(children...)
}

func Taxi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(taxiInnerSVG).
		Element(children...)
}

func Teeth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(teethInnerSVG).
		Element(children...)
}

func Terminal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalInnerSVG).
		Element(children...)
}

func TextHeight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textHeightInnerSVG).
		Element(children...)
}

func TextInput(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textInputInnerSVG).
		Element(children...)
}

func TextLabel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textLabelInnerSVG).
		Element(children...)
}

func TextWidth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textWidthInnerSVG).
		Element(children...)
}

func ThinSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thinSquareInnerSVG).
		Element(children...)
}

func ThumbsDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbsDownInnerSVG).
		Element(children...)
}

func ThumbsDownO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbsDownOInnerSVG).
		Element(children...)
}

func ThumbsUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbsUpInnerSVG).
		Element(children...)
}

func ThumbsUpO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbsUpOInnerSVG).
		Element(children...)
}

func Ticket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ticketInnerSVG).
		Element(children...)
}

func TimeBackward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timeBackwardInnerSVG).
		Element(children...)
}

func TimeForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timeForwardInnerSVG).
		Element(children...)
}

func Timer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timerInnerSVG).
		Element(children...)
}

func Toolbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toolboxInnerSVG).
		Element(children...)
}

func Tools(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toolsInnerSVG).
		Element(children...)
}

func Tooth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toothInnerSVG).
		Element(children...)
}

func Touch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(touchInnerSVG).
		Element(children...)
}

func Train(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trainInnerSVG).
		Element(children...)
}

func Trash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashInnerSVG).
		Element(children...)
}

func TreeTable(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(treeTableInnerSVG).
		Element(children...)
}

func TrendindDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendindDownInnerSVG).
		Element(children...)
}

func TrendingUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendingUpInnerSVG).
		Element(children...)
}

func Trophy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trophyInnerSVG).
		Element(children...)
}

func Truck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(truckInnerSVG).
		Element(children...)
}

func TwinColSelect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twinColSelectInnerSVG).
		Element(children...)
}

func Twitter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitterInnerSVG).
		Element(children...)
}

func TwitterSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitterSquareInnerSVG).
		Element(children...)
}

func Umbrella(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(underlineInnerSVG).
		Element(children...)
}

func Unlink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlinkInnerSVG).
		Element(children...)
}

func Unlock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadInnerSVG).
		Element(children...)
}

func UploadAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadAltInnerSVG).
		Element(children...)
}

func User(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userInnerSVG).
		Element(children...)
}

func UserCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCardInnerSVG).
		Element(children...)
}

func UserCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCheckInnerSVG).
		Element(children...)
}

func UserClock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userClockInnerSVG).
		Element(children...)
}

func UserHeart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userHeartInnerSVG).
		Element(children...)
}

func UserStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userStarInnerSVG).
		Element(children...)
}

func Users(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usersInnerSVG).
		Element(children...)
}

func VaadinH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vaadinHInnerSVG).
		Element(children...)
}

func VaadinV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vaadinVInnerSVG).
		Element(children...)
}

func Viewport(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewportInnerSVG).
		Element(children...)
}

func Vimeo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vimeoInnerSVG).
		Element(children...)
}

func VimeoSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vimeoSquareInnerSVG).
		Element(children...)
}

func Volume(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeInnerSVG).
		Element(children...)
}

func VolumeDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeDownInnerSVG).
		Element(children...)
}

func VolumeOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeOffInnerSVG).
		Element(children...)
}

func VolumeUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeUpInnerSVG).
		Element(children...)
}

func Wallet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(walletInnerSVG).
		Element(children...)
}

func Warning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningInnerSVG).
		Element(children...)
}

func Workplace(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(workplaceInnerSVG).
		Element(children...)
}

func Wrench(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wrenchInnerSVG).
		Element(children...)
}

func Youtube(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(youtubeInnerSVG).
		Element(children...)
}

func YoutubeSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(youtubeSquareInnerSVG).
		Element(children...)
}

func ByName(name string) (*engine.UberElement, error) {
	switch name {
	case "abacus":
		return Abacus(), nil
	case "absolute-position":
		return AbsolutePosition(), nil
	case "academy-cap":
		return AcademyCap(), nil
	case "accessibility":
		return Accessibility(), nil
	case "accordion-menu":
		return AccordionMenu(), nil
	case "add-dock":
		return AddDock(), nil
	case "adjust":
		return Adjust(), nil
	case "adobe-flash":
		return AdobeFlash(), nil
	case "airplane":
		return Airplane(), nil
	case "alarm":
		return Alarm(), nil
	case "align-center":
		return AlignCenter(), nil
	case "align-justify":
		return AlignJustify(), nil
	case "align-left":
		return AlignLeft(), nil
	case "align-right":
		return AlignRight(), nil
	case "alt":
		return Alt(), nil
	case "alt-a":
		return AltA(), nil
	case "ambulance":
		return Ambulance(), nil
	case "anchor":
		return Anchor(), nil
	case "angle-double-down":
		return AngleDoubleDown(), nil
	case "angle-double-left":
		return AngleDoubleLeft(), nil
	case "angle-double-right":
		return AngleDoubleRight(), nil
	case "angle-double-up":
		return AngleDoubleUp(), nil
	case "angle-down":
		return AngleDown(), nil
	case "angle-left":
		return AngleLeft(), nil
	case "angle-right":
		return AngleRight(), nil
	case "angle-up":
		return AngleUp(), nil
	case "archive":
		return Archive(), nil
	case "archives":
		return Archives(), nil
	case "area-select":
		return AreaSelect(), nil
	case "arrow-backward":
		return ArrowBackward(), nil
	case "arrow-circle-down":
		return ArrowCircleDown(), nil
	case "arrow-circle-down-o":
		return ArrowCircleDownO(), nil
	case "arrow-circle-left":
		return ArrowCircleLeft(), nil
	case "arrow-circle-left-o":
		return ArrowCircleLeftO(), nil
	case "arrow-circle-right":
		return ArrowCircleRight(), nil
	case "arrow-circle-right-o":
		return ArrowCircleRightO(), nil
	case "arrow-circle-up":
		return ArrowCircleUp(), nil
	case "arrow-circle-up-o":
		return ArrowCircleUpO(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-forward":
		return ArrowForward(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-long-down":
		return ArrowLongDown(), nil
	case "arrow-long-left":
		return ArrowLongLeft(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrows":
		return Arrows(), nil
	case "arrows-cross":
		return ArrowsCross(), nil
	case "arrows-long-h":
		return ArrowsLongH(), nil
	case "arrows-long-right":
		return ArrowsLongRight(), nil
	case "arrows-long-up":
		return ArrowsLongUp(), nil
	case "arrows-long-v":
		return ArrowsLongV(), nil
	case "asterisk":
		return Asterisk(), nil
	case "at":
		return At(), nil
	case "automation":
		return Automation(), nil
	case "backspace":
		return Backspace(), nil
	case "backspace-a":
		return BackspaceA(), nil
	case "backwards":
		return Backwards(), nil
	case "ban":
		return Ban(), nil
	case "bar-chart":
		return BarChart(), nil
	case "bar-chart-h":
		return BarChartH(), nil
	case "bar-chart-v":
		return BarChartV(), nil
	case "barcode":
		return Barcode(), nil
	case "bed":
		return Bed(), nil
	case "bell":
		return Bell(), nil
	case "bell-o":
		return BellO(), nil
	case "bell-slash":
		return BellSlash(), nil
	case "bell-slash-o":
		return BellSlashO(), nil
	case "boat":
		return Boat(), nil
	case "bold":
		return Bold(), nil
	case "bolt":
		return Bolt(), nil
	case "bomb":
		return Bomb(), nil
	case "book":
		return Book(), nil
	case "book-dollar":
		return BookDollar(), nil
	case "book-percent":
		return BookPercent(), nil
	case "bookmark":
		return Bookmark(), nil
	case "bookmark-o":
		return BookmarkO(), nil
	case "briefcase":
		return Briefcase(), nil
	case "browser":
		return Browser(), nil
	case "bug":
		return Bug(), nil
	case "bug-o":
		return BugO(), nil
	case "building":
		return Building(), nil
	case "building-o":
		return BuildingO(), nil
	case "bullets":
		return Bullets(), nil
	case "bullseye":
		return Bullseye(), nil
	case "buss":
		return Buss(), nil
	case "button":
		return Button(), nil
	case "calc":
		return Calc(), nil
	case "calc-book":
		return CalcBook(), nil
	case "calendar":
		return Calendar(), nil
	case "calendar-briefcase":
		return CalendarBriefcase(), nil
	case "calendar-clock":
		return CalendarClock(), nil
	case "calendar-envelope":
		return CalendarEnvelope(), nil
	case "calendar-o":
		return CalendarO(), nil
	case "calendar-user":
		return CalendarUser(), nil
	case "camera":
		return Camera(), nil
	case "car":
		return Car(), nil
	case "caret-down":
		return CaretDown(), nil
	case "caret-left":
		return CaretLeft(), nil
	case "caret-right":
		return CaretRight(), nil
	case "caret-square-down-o":
		return CaretSquareDownO(), nil
	case "caret-square-left-o":
		return CaretSquareLeftO(), nil
	case "caret-square-right-o":
		return CaretSquareRightO(), nil
	case "caret-square-up-o":
		return CaretSquareUpO(), nil
	case "caret-up":
		return CaretUp(), nil
	case "cart":
		return Cart(), nil
	case "cart-o":
		return CartO(), nil
	case "cash":
		return Cash(), nil
	case "chart":
		return Chart(), nil
	case "chart-grid":
		return ChartGrid(), nil
	case "chart-line":
		return ChartLine(), nil
	case "chart-3d":
		return ChartThreeD(), nil
	case "chart-timeline":
		return ChartTimeline(), nil
	case "chat":
		return Chat(), nil
	case "check":
		return Check(), nil
	case "check-circle":
		return CheckCircle(), nil
	case "check-circle-o":
		return CheckCircleO(), nil
	case "check-square":
		return CheckSquare(), nil
	case "check-square-o":
		return CheckSquareO(), nil
	case "chevron-circle-down":
		return ChevronCircleDown(), nil
	case "chevron-circle-down-o":
		return ChevronCircleDownO(), nil
	case "chevron-circle-left":
		return ChevronCircleLeft(), nil
	case "chevron-circle-left-o":
		return ChevronCircleLeftO(), nil
	case "chevron-circle-right":
		return ChevronCircleRight(), nil
	case "chevron-circle-right-o":
		return ChevronCircleRightO(), nil
	case "chevron-circle-up":
		return ChevronCircleUp(), nil
	case "chevron-circle-up-o":
		return ChevronCircleUpO(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-down-small":
		return ChevronDownSmall(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-left-small":
		return ChevronLeftSmall(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-right-small":
		return ChevronRightSmall(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "chevron-up-small":
		return ChevronUpSmall(), nil
	case "child":
		return Child(), nil
	case "circle":
		return Circle(), nil
	case "circle-thin":
		return CircleThin(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clipboard-check":
		return ClipboardCheck(), nil
	case "clipboard-cross":
		return ClipboardCross(), nil
	case "clipboard-heart":
		return ClipboardHeart(), nil
	case "clipboard-pulse":
		return ClipboardPulse(), nil
	case "clipboard-text":
		return ClipboardText(), nil
	case "clipboard-user":
		return ClipboardUser(), nil
	case "clock":
		return Clock(), nil
	case "close":
		return Close(), nil
	case "close-big":
		return CloseBig(), nil
	case "close-circle":
		return CloseCircle(), nil
	case "close-circle-o":
		return CloseCircleO(), nil
	case "close-small":
		return CloseSmall(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-download":
		return CloudDownload(), nil
	case "cloud-download-o":
		return CloudDownloadO(), nil
	case "cloud-o":
		return CloudO(), nil
	case "cloud-upload":
		return CloudUpload(), nil
	case "cloud-upload-o":
		return CloudUploadO(), nil
	case "cluster":
		return Cluster(), nil
	case "code":
		return Code(), nil
	case "coffee":
		return Coffee(), nil
	case "cog":
		return Cog(), nil
	case "cog-o":
		return CogO(), nil
	case "cogs":
		return Cogs(), nil
	case "coin-piles":
		return CoinPiles(), nil
	case "coins":
		return Coins(), nil
	case "combobox":
		return Combobox(), nil
	case "comment":
		return Comment(), nil
	case "comment-ellipsis":
		return CommentEllipsis(), nil
	case "comment-ellipsis-o":
		return CommentEllipsisO(), nil
	case "comment-o":
		return CommentO(), nil
	case "comments":
		return Comments(), nil
	case "comments-o":
		return CommentsO(), nil
	case "compile":
		return Compile(), nil
	case "compress":
		return Compress(), nil
	case "compress-square":
		return CompressSquare(), nil
	case "connect":
		return Connect(), nil
	case "connect-o":
		return ConnectO(), nil
	case "controller":
		return Controller(), nil
	case "copy":
		return Copy(), nil
	case "copy-o":
		return CopyO(), nil
	case "copyright":
		return Copyright(), nil
	case "corner-lower-left":
		return CornerLowerLeft(), nil
	case "corner-lower-right":
		return CornerLowerRight(), nil
	case "corner-upper-left":
		return CornerUpperLeft(), nil
	case "corner-upper-right":
		return CornerUpperRight(), nil
	case "credit-card":
		return CreditCard(), nil
	case "crop":
		return Crop(), nil
	case "cross-cutlery":
		return CrossCutlery(), nil
	case "crosshairs":
		return Crosshairs(), nil
	case "css":
		return Css(), nil
	case "ctrl":
		return Ctrl(), nil
	case "ctrl-a":
		return CtrlA(), nil
	case "cube":
		return Cube(), nil
	case "cubes":
		return Cubes(), nil
	case "curly-brackets":
		return CurlyBrackets(), nil
	case "cursor":
		return Cursor(), nil
	case "cursor-o":
		return CursorO(), nil
	case "cutlery":
		return Cutlery(), nil
	case "dashboard":
		return Dashboard(), nil
	case "database":
		return Database(), nil
	case "date-input":
		return DateInput(), nil
	case "deindent":
		return Deindent(), nil
	case "del":
		return Del(), nil
	case "del-a":
		return DelA(), nil
	case "dental-chair":
		return DentalChair(), nil
	case "desktop":
		return Desktop(), nil
	case "diamond":
		return Diamond(), nil
	case "diamond-o":
		return DiamondO(), nil
	case "diploma":
		return Diploma(), nil
	case "diploma-scroll":
		return DiplomaScroll(), nil
	case "disc":
		return Disc(), nil
	case "doctor":
		return Doctor(), nil
	case "doctor-briefcase":
		return DoctorBriefcase(), nil
	case "dollar":
		return Dollar(), nil
	case "dot-circle":
		return DotCircle(), nil
	case "download":
		return Download(), nil
	case "download-alt":
		return DownloadAlt(), nil
	case "drop":
		return Drop(), nil
	case "edit":
		return Edit(), nil
	case "eject":
		return Eject(), nil
	case "elastic":
		return Elastic(), nil
	case "ellipsis-circle":
		return EllipsisCircle(), nil
	case "ellipsis-circle-o":
		return EllipsisCircleO(), nil
	case "ellipsis-dots-h":
		return EllipsisDotsH(), nil
	case "ellipsis-dots-v":
		return EllipsisDotsV(), nil
	case "ellipsis-h":
		return EllipsisH(), nil
	case "ellipsis-v":
		return EllipsisV(), nil
	case "enter":
		return Enter(), nil
	case "enter-arrow":
		return EnterArrow(), nil
	case "envelope":
		return Envelope(), nil
	case "envelope-o":
		return EnvelopeO(), nil
	case "envelope-open":
		return EnvelopeOpen(), nil
	case "envelope-open-o":
		return EnvelopeOpenO(), nil
	case "envelopes":
		return Envelopes(), nil
	case "envelopes-o":
		return EnvelopesO(), nil
	case "eraser":
		return Eraser(), nil
	case "esc":
		return Esc(), nil
	case "esc-a":
		return EscA(), nil
	case "euro":
		return Euro(), nil
	case "exchange":
		return Exchange(), nil
	case "exclamation":
		return Exclamation(), nil
	case "exclamation-circle":
		return ExclamationCircle(), nil
	case "exclamation-circle-o":
		return ExclamationCircleO(), nil
	case "exit":
		return Exit(), nil
	case "exit-o":
		return ExitO(), nil
	case "expand":
		return Expand(), nil
	case "expand-full":
		return ExpandFull(), nil
	case "expand-square":
		return ExpandSquare(), nil
	case "external-browser":
		return ExternalBrowser(), nil
	case "external-link":
		return ExternalLink(), nil
	case "eye":
		return Eye(), nil
	case "eye-slash":
		return EyeSlash(), nil
	case "eyedropper":
		return Eyedropper(), nil
	case "facebook":
		return Facebook(), nil
	case "facebook-square":
		return FacebookSquare(), nil
	case "factory":
		return Factory(), nil
	case "family":
		return Family(), nil
	case "fast-backward":
		return FastBackward(), nil
	case "fast-forward":
		return FastForward(), nil
	case "female":
		return Female(), nil
	case "file":
		return File(), nil
	case "file-add":
		return FileAdd(), nil
	case "file-code":
		return FileCode(), nil
	case "file-font":
		return FileFont(), nil
	case "file-movie":
		return FileMovie(), nil
	case "file-o":
		return FileO(), nil
	case "file-picture":
		return FilePicture(), nil
	case "file-presentation":
		return FilePresentation(), nil
	case "file-process":
		return FileProcess(), nil
	case "file-refresh":
		return FileRefresh(), nil
	case "file-remove":
		return FileRemove(), nil
	case "file-search":
		return FileSearch(), nil
	case "file-sound":
		return FileSound(), nil
	case "file-start":
		return FileStart(), nil
	case "file-table":
		return FileTable(), nil
	case "file-text":
		return FileText(), nil
	case "file-text-o":
		return FileTextO(), nil
	case "file-tree":
		return FileTree(), nil
	case "file-tree-small":
		return FileTreeSmall(), nil
	case "file-tree-sub":
		return FileTreeSub(), nil
	case "file-zip":
		return FileZip(), nil
	case "fill":
		return Fill(), nil
	case "film":
		return Film(), nil
	case "filter":
		return Filter(), nil
	case "fire":
		return Fire(), nil
	case "flag":
		return Flag(), nil
	case "flag-checkered":
		return FlagCheckered(), nil
	case "flag-o":
		return FlagO(), nil
	case "flash":
		return Flash(), nil
	case "flask":
		return Flask(), nil
	case "flight-landing":
		return FlightLanding(), nil
	case "flight-takeoff":
		return FlightTakeoff(), nil
	case "flip-h":
		return FlipH(), nil
	case "flip-v":
		return FlipV(), nil
	case "folder":
		return Folder(), nil
	case "folder-add":
		return FolderAdd(), nil
	case "folder-o":
		return FolderO(), nil
	case "folder-open":
		return FolderOpen(), nil
	case "folder-open-o":
		return FolderOpenO(), nil
	case "folder-remove":
		return FolderRemove(), nil
	case "folder-search":
		return FolderSearch(), nil
	case "font":
		return Font(), nil
	case "form":
		return Form(), nil
	case "forward":
		return Forward(), nil
	case "frown-o":
		return FrownO(), nil
	case "funcion":
		return Funcion(), nil
	case "funnel":
		return Funnel(), nil
	case "gamepad":
		return Gamepad(), nil
	case "gavel":
		return Gavel(), nil
	case "gift":
		return Gift(), nil
	case "glass":
		return Glass(), nil
	case "glasses":
		return Glasses(), nil
	case "globe":
		return Globe(), nil
	case "globe-wire":
		return GlobeWire(), nil
	case "golf":
		return Golf(), nil
	case "google-plus":
		return GooglePlus(), nil
	case "google-plus-square":
		return GooglePlusSquare(), nil
	case "grab":
		return Grab(), nil
	case "grid":
		return Grid(), nil
	case "grid-bevel":
		return GridBevel(), nil
	case "grid-big":
		return GridBig(), nil
	case "grid-big-o":
		return GridBigO(), nil
	case "grid-h":
		return GridH(), nil
	case "grid-small":
		return GridSmall(), nil
	case "grid-small-o":
		return GridSmallO(), nil
	case "grid-v":
		return GridV(), nil
	case "group":
		return Group(), nil
	case "hammer":
		return Hammer(), nil
	case "hand":
		return Hand(), nil
	case "handle-corner":
		return HandleCorner(), nil
	case "hands-up":
		return HandsUp(), nil
	case "handshake":
		return Handshake(), nil
	case "harddrive":
		return Harddrive(), nil
	case "harddrive-o":
		return HarddriveO(), nil
	case "hash":
		return Hash(), nil
	case "header":
		return Header(), nil
	case "headphones":
		return Headphones(), nil
	case "headset":
		return Headset(), nil
	case "health-card":
		return HealthCard(), nil
	case "heart":
		return Heart(), nil
	case "heart-o":
		return HeartO(), nil
	case "home":
		return Home(), nil
	case "home-o":
		return HomeO(), nil
	case "hospital":
		return Hospital(), nil
	case "hourglass":
		return Hourglass(), nil
	case "hourglass-empty":
		return HourglassEmpty(), nil
	case "hourglass-end":
		return HourglassEnd(), nil
	case "hourglass-start":
		return HourglassStart(), nil
	case "inbox":
		return Inbox(), nil
	case "indent":
		return Indent(), nil
	case "info":
		return Info(), nil
	case "info-circle":
		return InfoCircle(), nil
	case "info-circle-o":
		return InfoCircleO(), nil
	case "input":
		return Input(), nil
	case "insert":
		return Insert(), nil
	case "institution":
		return Institution(), nil
	case "invoice":
		return Invoice(), nil
	case "italic":
		return Italic(), nil
	case "key":
		return Key(), nil
	case "key-o":
		return KeyO(), nil
	case "keyboard":
		return Keyboard(), nil
	case "keyboard-o":
		return KeyboardO(), nil
	case "laptop":
		return Laptop(), nil
	case "layout":
		return Layout(), nil
	case "level-down":
		return LevelDown(), nil
	case "level-down-bold":
		return LevelDownBold(), nil
	case "level-left":
		return LevelLeft(), nil
	case "level-left-bold":
		return LevelLeftBold(), nil
	case "level-right":
		return LevelRight(), nil
	case "level-right-bold":
		return LevelRightBold(), nil
	case "level-up":
		return LevelUp(), nil
	case "level-up-bold":
		return LevelUpBold(), nil
	case "lifebuoy":
		return Lifebuoy(), nil
	case "lightbulb":
		return Lightbulb(), nil
	case "line-bar-chart":
		return LineBarChart(), nil
	case "line-chart":
		return LineChart(), nil
	case "line-h":
		return LineH(), nil
	case "line-v":
		return LineV(), nil
	case "lines":
		return Lines(), nil
	case "lines-list":
		return LinesList(), nil
	case "link":
		return Link(), nil
	case "list":
		return List(), nil
	case "list-ol":
		return ListOl(), nil
	case "list-select":
		return ListSelect(), nil
	case "list-ul":
		return ListUl(), nil
	case "location-arrow":
		return LocationArrow(), nil
	case "location-arrow-circle":
		return LocationArrowCircle(), nil
	case "location-arrow-circle-o":
		return LocationArrowCircleO(), nil
	case "lock":
		return Lock(), nil
	case "magic":
		return Magic(), nil
	case "magnet":
		return Magnet(), nil
	case "mailbox":
		return Mailbox(), nil
	case "male":
		return Male(), nil
	case "map-marker":
		return MapMarker(), nil
	case "margin":
		return Margin(), nil
	case "margin-bottom":
		return MarginBottom(), nil
	case "margin-left":
		return MarginLeft(), nil
	case "margin-right":
		return MarginRight(), nil
	case "margin-top":
		return MarginTop(), nil
	case "medal":
		return Medal(), nil
	case "megafone":
		return Megafone(), nil
	case "meh-o":
		return MehO(), nil
	case "menu":
		return Menu(), nil
	case "microphone":
		return Microphone(), nil
	case "minus":
		return Minus(), nil
	case "minus-circle":
		return MinusCircle(), nil
	case "minus-circle-o":
		return MinusCircleO(), nil
	case "minus-square-o":
		return MinusSquareO(), nil
	case "mobile":
		return Mobile(), nil
	case "mobile-browser":
		return MobileBrowser(), nil
	case "mobile-retro":
		return MobileRetro(), nil
	case "modal":
		return Modal(), nil
	case "modal-list":
		return ModalList(), nil
	case "money":
		return Money(), nil
	case "money-deposit":
		return MoneyDeposit(), nil
	case "money-exchange":
		return MoneyExchange(), nil
	case "money-withdraw":
		return MoneyWithdraw(), nil
	case "moon":
		return Moon(), nil
	case "moon-o":
		return MoonO(), nil
	case "morning":
		return Morning(), nil
	case "movie":
		return Movie(), nil
	case "music":
		return Music(), nil
	case "mute":
		return Mute(), nil
	case "native-button":
		return NativeButton(), nil
	case "newspaper":
		return Newspaper(), nil
	case "notebook":
		return Notebook(), nil
	case "nurse":
		return Nurse(), nil
	case "office":
		return Office(), nil
	case "open-book":
		return OpenBook(), nil
	case "option":
		return Option(), nil
	case "option-a":
		return OptionA(), nil
	case "options":
		return Options(), nil
	case "orientation":
		return Orientation(), nil
	case "out":
		return Out(), nil
	case "outbox":
		return Outbox(), nil
	case "package":
		return Package(), nil
	case "padding":
		return Padding(), nil
	case "padding-bottom":
		return PaddingBottom(), nil
	case "padding-left":
		return PaddingLeft(), nil
	case "padding-right":
		return PaddingRight(), nil
	case "padding-top":
		return PaddingTop(), nil
	case "paint-roll":
		return PaintRoll(), nil
	case "paintbrush":
		return Paintbrush(), nil
	case "palete":
		return Palete(), nil
	case "panel":
		return Panel(), nil
	case "paperclip":
		return Paperclip(), nil
	case "paperplane":
		return Paperplane(), nil
	case "paperplane-o":
		return PaperplaneO(), nil
	case "paragraph":
		return Paragraph(), nil
	case "password":
		return Password(), nil
	case "paste":
		return Paste(), nil
	case "pause":
		return Pause(), nil
	case "pencil":
		return Pencil(), nil
	case "phone":
		return Phone(), nil
	case "phone-landline":
		return PhoneLandline(), nil
	case "picture":
		return Picture(), nil
	case "pie-bar-chart":
		return PieBarChart(), nil
	case "pie-chart":
		return PieChart(), nil
	case "piggy-bank":
		return PiggyBank(), nil
	case "piggy-bank-coin":
		return PiggyBankCoin(), nil
	case "pill":
		return Pill(), nil
	case "pills":
		return Pills(), nil
	case "pin":
		return Pin(), nil
	case "pin-post":
		return PinPost(), nil
	case "play":
		return Play(), nil
	case "play-circle":
		return PlayCircle(), nil
	case "play-circle-o":
		return PlayCircleO(), nil
	case "plug":
		return Plug(), nil
	case "plus":
		return Plus(), nil
	case "plus-circle":
		return PlusCircle(), nil
	case "plus-circle-o":
		return PlusCircleO(), nil
	case "plus-minus":
		return PlusMinus(), nil
	case "plus-square-o":
		return PlusSquareO(), nil
	case "pointer":
		return Pointer(), nil
	case "power-off":
		return PowerOff(), nil
	case "presentation":
		return Presentation(), nil
	case "print":
		return Print(), nil
	case "progressbar":
		return Progressbar(), nil
	case "puzzle-piece":
		return PuzzlePiece(), nil
	case "pyramid-chart":
		return PyramidChart(), nil
	case "qrcode":
		return Qrcode(), nil
	case "question":
		return Question(), nil
	case "question-circle":
		return QuestionCircle(), nil
	case "question-circle-o":
		return QuestionCircleO(), nil
	case "quote-left":
		return QuoteLeft(), nil
	case "quote-right":
		return QuoteRight(), nil
	case "random":
		return Random(), nil
	case "raster":
		return Raster(), nil
	case "raster-lower-left":
		return RasterLowerLeft(), nil
	case "records":
		return Records(), nil
	case "recycle":
		return Recycle(), nil
	case "refresh":
		return Refresh(), nil
	case "reply":
		return Reply(), nil
	case "reply-all":
		return ReplyAll(), nil
	case "resize-h":
		return ResizeH(), nil
	case "resize-v":
		return ResizeV(), nil
	case "retweet":
		return Retweet(), nil
	case "rhombus":
		return Rhombus(), nil
	case "road":
		return Road(), nil
	case "road-branch":
		return RoadBranch(), nil
	case "road-branches":
		return RoadBranches(), nil
	case "road-split":
		return RoadSplit(), nil
	case "rocket":
		return Rocket(), nil
	case "rotate-left":
		return RotateLeft(), nil
	case "rotate-right":
		return RotateRight(), nil
	case "rss":
		return Rss(), nil
	case "rss-square":
		return RssSquare(), nil
	case "safe":
		return Safe(), nil
	case "safe-lock":
		return SafeLock(), nil
	case "scale":
		return Scale(), nil
	case "scale-unbalance":
		return ScaleUnbalance(), nil
	case "scatter-chart":
		return ScatterChart(), nil
	case "scissors":
		return Scissors(), nil
	case "screwdriver":
		return Screwdriver(), nil
	case "search":
		return Search(), nil
	case "search-minus":
		return SearchMinus(), nil
	case "search-plus":
		return SearchPlus(), nil
	case "select":
		return Select(), nil
	case "server":
		return Server(), nil
	case "share":
		return Share(), nil
	case "share-square":
		return ShareSquare(), nil
	case "shield":
		return Shield(), nil
	case "shift":
		return Shift(), nil
	case "shift-arrow":
		return ShiftArrow(), nil
	case "shop":
		return Shop(), nil
	case "sign-in":
		return SignIn(), nil
	case "sign-in-alt":
		return SignInAlt(), nil
	case "sign-out":
		return SignOut(), nil
	case "sign-out-alt":
		return SignOutAlt(), nil
	case "signal":
		return Signal(), nil
	case "sitemap":
		return Sitemap(), nil
	case "slider":
		return Slider(), nil
	case "sliders":
		return Sliders(), nil
	case "smiley-o":
		return SmileyO(), nil
	case "sort":
		return Sort(), nil
	case "sound-disable":
		return SoundDisable(), nil
	case "spark-line":
		return SparkLine(), nil
	case "specialist":
		return Specialist(), nil
	case "spinner":
		return Spinner(), nil
	case "spinner-arc":
		return SpinnerArc(), nil
	case "spinner-third":
		return SpinnerThird(), nil
	case "spline-area-chart":
		return SplineAreaChart(), nil
	case "spline-chart":
		return SplineChart(), nil
	case "split":
		return Split(), nil
	case "split-h":
		return SplitH(), nil
	case "split-v":
		return SplitV(), nil
	case "spoon":
		return Spoon(), nil
	case "square-shadow":
		return SquareShadow(), nil
	case "star":
		return Star(), nil
	case "star-half-left":
		return StarHalfLeft(), nil
	case "star-half-left-o":
		return StarHalfLeftO(), nil
	case "star-half-right":
		return StarHalfRight(), nil
	case "star-half-right-o":
		return StarHalfRightO(), nil
	case "star-o":
		return StarO(), nil
	case "start-cog":
		return StartCog(), nil
	case "step-backward":
		return StepBackward(), nil
	case "step-forward":
		return StepForward(), nil
	case "stethoscope":
		return Stethoscope(), nil
	case "stock":
		return Stock(), nil
	case "stop":
		return Stop(), nil
	case "stop-cog":
		return StopCog(), nil
	case "stopwatch":
		return Stopwatch(), nil
	case "storage":
		return Storage(), nil
	case "strikethrough":
		return Strikethrough(), nil
	case "subscript":
		return Subscript(), nil
	case "suitcase":
		return Suitcase(), nil
	case "sun-down":
		return SunDown(), nil
	case "sun-o":
		return SunO(), nil
	case "sun-rise":
		return SunRise(), nil
	case "superscript":
		return Superscript(), nil
	case "sword":
		return Sword(), nil
	case "tab":
		return Tab(), nil
	case "tab-a":
		return TabA(), nil
	case "table":
		return Table(), nil
	case "tablet":
		return Tablet(), nil
	case "tabs":
		return Tabs(), nil
	case "tag":
		return Tag(), nil
	case "tags":
		return Tags(), nil
	case "tasks":
		return Tasks(), nil
	case "taxi":
		return Taxi(), nil
	case "teeth":
		return Teeth(), nil
	case "terminal":
		return Terminal(), nil
	case "text-height":
		return TextHeight(), nil
	case "text-input":
		return TextInput(), nil
	case "text-label":
		return TextLabel(), nil
	case "text-width":
		return TextWidth(), nil
	case "thin-square":
		return ThinSquare(), nil
	case "thumbs-down":
		return ThumbsDown(), nil
	case "thumbs-down-o":
		return ThumbsDownO(), nil
	case "thumbs-up":
		return ThumbsUp(), nil
	case "thumbs-up-o":
		return ThumbsUpO(), nil
	case "ticket":
		return Ticket(), nil
	case "time-backward":
		return TimeBackward(), nil
	case "time-forward":
		return TimeForward(), nil
	case "timer":
		return Timer(), nil
	case "toolbox":
		return Toolbox(), nil
	case "tools":
		return Tools(), nil
	case "tooth":
		return Tooth(), nil
	case "touch":
		return Touch(), nil
	case "train":
		return Train(), nil
	case "trash":
		return Trash(), nil
	case "tree-table":
		return TreeTable(), nil
	case "trendind-down":
		return TrendindDown(), nil
	case "trending-up":
		return TrendingUp(), nil
	case "trophy":
		return Trophy(), nil
	case "truck":
		return Truck(), nil
	case "twin-col-select":
		return TwinColSelect(), nil
	case "twitter":
		return Twitter(), nil
	case "twitter-square":
		return TwitterSquare(), nil
	case "umbrella":
		return Umbrella(), nil
	case "underline":
		return Underline(), nil
	case "unlink":
		return Unlink(), nil
	case "unlock":
		return Unlock(), nil
	case "upload":
		return Upload(), nil
	case "upload-alt":
		return UploadAlt(), nil
	case "user":
		return User(), nil
	case "user-card":
		return UserCard(), nil
	case "user-check":
		return UserCheck(), nil
	case "user-clock":
		return UserClock(), nil
	case "user-heart":
		return UserHeart(), nil
	case "user-star":
		return UserStar(), nil
	case "users":
		return Users(), nil
	case "vaadin-h":
		return VaadinH(), nil
	case "vaadin-v":
		return VaadinV(), nil
	case "viewport":
		return Viewport(), nil
	case "vimeo":
		return Vimeo(), nil
	case "vimeo-square":
		return VimeoSquare(), nil
	case "volume":
		return Volume(), nil
	case "volume-down":
		return VolumeDown(), nil
	case "volume-off":
		return VolumeOff(), nil
	case "volume-up":
		return VolumeUp(), nil
	case "wallet":
		return Wallet(), nil
	case "warning":
		return Warning(), nil
	case "workplace":
		return Workplace(), nil
	case "wrench":
		return Wrench(), nil
	case "youtube":
		return Youtube(), nil
	case "youtube-square":
		return YoutubeSquare(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in vaadin icon set", name)
	}
}
