package iconoir

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	accessibilityInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10ZM7 9l5 1m5-1l-5 1m0 0v3m0 0l-2 5m2-5l2 5"/><path fill="currentColor" d="M12 7a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	accessibilitySignInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16h3.889l1.555 2.5H19m-7-10V11m0 5v-5m0 0h3.889M12 6.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/><path d="M14.882 19.516A5.5 5.5 0 1 1 8.678 11"/></g>`
	accessibilityTechInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.72 11.786a3 3 0 1 0 3.576 4.646M12 14h2.5l1 1.5h1m-4.5-6V11m0 3v-3m0 0h2.5"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 7.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	activityInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h3l3-9l6 18l3-9h3"/>`
	addCircleInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h4m4 0h-4m0 0V8m0 4v4m0 6c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	addDatabaseScriptInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 14V8.5M6 13V6a3 3 0 0 1 3-3h5m2.992 1h3m3 0h-3m0 0V1m0 3v3M12 21H6a4 4 0 0 1 0-8h12a4 4 0 1 0 4 4v-3"/>`
	addFolderInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 6h2m2 0h-2m0 0V4m0 2v2m1.4 12H2.6a.6.6 0 0 1-.6-.6V11h19.4a.6.6 0 0 1 .6.6v7.8a.6.6 0 0 1-.6.6ZM2 11V4.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H14"/>`
	addFrameInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-miterlimit="1.5" d="M4.998 2H2v2.998h2.998V2Zm.001 1.5h14M3.5 4.998V19M20.498 5v14.002M4.999 20.5h14M4.998 19H2v2.998h2.998V19ZM21.997 2.001H19v2.998h2.998V2.001Zm0 17H19v2.998h2.998v-2.998Z"/><path d="M9 12h3m3 0h-3m0 0V9m0 3v3"/></g>`
	addHexagonInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3m3 0h-3m0 0V9m0 3v3m-.3-13.827a.6.6 0 0 1 .6 0l8.926 5.154a.6.6 0 0 1 .3.52v10.307a.6.6 0 0 1-.3.52L12.3 22.826a.6.6 0 0 1-.6 0l-8.926-5.154a.6.6 0 0 1-.3-.52V6.847a.6.6 0 0 1 .3-.52L11.7 1.174Z"/>`
	addKeyframeInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 5h3m3 0h-3m0 0V2m0 3v3m-3.152 5.317l-4.343 4.963a2 2 0 0 1-3.01 0l-4.343-4.963a2 2 0 0 1 0-2.634L7.495 5.72a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634Z"/>`
	addKeyframeAltInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18.819 13.329l-5.324 5.99a2 2 0 0 1-2.99 0l-5.324-5.99a2 2 0 0 1 0-2.658l5.324-5.99a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.658ZM9 12h3m3 0h-3m0 0V9m0 3v3"/>`
	addKeyframesInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12h3m3 0H5m0 0V9m0 3v3m1.25-9l.245-.28a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634L9.505 18.28a2 2 0 0 1-3.01 0L6.25 18"/><path d="m13 19l4.884-5.698a2 2 0 0 0 0-2.604L13 5"/><path d="m17 19l4.884-5.698a2 2 0 0 0 0-2.604L17 5"/></g>`
	addLensInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.992 6h3m3 0h-3m0 0V3m0 3v3m-3.88 4.5C2.835 18.311 6.987 22 12 22c5.523 0 10-4.477 10-10c0-5.013-3.689-9.165-8.5-9.888"/><path d="M17.197 9c-.1-.172-.207-.34-.323-.5m.937 5a6.01 6.01 0 0 1-4.311 4.311"/></g>`
	addMediaImageInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13"/><path d="m3 16l7-3l5.5 2.5M16 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 9h3m3 0h-3m0 0v-3m0 3v3"/></g>`
	addMediaVideoInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13m-5 6h3m3 0h-3m0 0v-3m0 3v3"/><path d="M9.898 8.513a.6.6 0 0 0-.898.52v5.933a.6.6 0 0 0 .898.521l5.19-2.966a.6.6 0 0 0 0-1.042l-5.19-2.966Z"/></g>`
	addPageInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 12h3m3 0h-3m0 0V9m0 3v3m-8 6.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	addPageAltInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 12V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H11"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20M1.992 19h3m3 0h-3m0 0v-3m0 3v3"/></g>`
	addPinAltInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 9.2C16 13.177 9 20 9 20S2 13.177 2 9.2C2 5.224 5.134 2 9 2s7 3.224 7 7.2Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 19h3m3 0h-3m0 0v-3m0 3v3"/></g>`
	addSelectionInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h4m4 0h-4m0 0V8m0 4v4M7 4H4v3m0 4v2m7-9h2m-2 16h2m7-9v2m-3-9h3v3M7 20H4v-3m13 3h3v-3"/>`
	addSquareInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3m3 0h-3m0 0V9m0 3v3m9-11.4v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	addToCartInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h19l-3 10H6L3 6Zm0 0l-.75-2.5M9.992 11h2m2 0h-2m0 0V9m0 2v2M11 19.5a1.5 1.5 0 0 1-3 0m9 0a1.5 1.5 0 0 1-3 0"/>`
	addUserInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 10h3m3 0h-3m0 0V7m0 3v3M1 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1m-7-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`
	africanTreeInnerSVG                       = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirAfricanTree0)"><path d="M12 22V12m0-4v4m0 0l3-3m-2.576 9.576l6.169-6.169a5.502 5.502 0 0 0-.513-8.234a9.904 9.904 0 0 0-12.16 0a5.502 5.502 0 0 0-.513 8.234l6.169 6.169a.6.6 0 0 0 .848 0Z"/></g><defs><clipPath id="iconoirAfricanTree0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	agileInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 19H22m0 0l-2.5-2.5M22 19l-2.5 2.5M12 2L9.5 4.5L12 7"/><path d="M10.5 4.5a7.5 7.5 0 0 1 0 15H2"/><path d="M6.756 5.5A7.497 7.497 0 0 0 3 12c0 1.688.558 3.246 1.5 4.5"/></g>`
	airConditionerInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 3.6V11H2V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6ZM18 7h1M2 11l.79 2.584A2 2 0 0 0 4.702 15H6m16-4l-.79 2.584A2 2 0 0 1 19.298 15H18m-8.5-.5s0 7-3.5 7m8.5-7s0 7 3.5 7m-6-7v7"/>`
	airplaneInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 4.5v4.667a.6.6 0 0 1-.282.51l-7.436 4.647a.6.6 0 0 0-.282.508v.9a.6.6 0 0 0 .746.582l6.508-1.628a.6.6 0 0 1 .746.582v2.96a.6.6 0 0 1-.205.451l-2.16 1.89c-.458.402-.097 1.151.502 1.042l3.256-.591a.6.6 0 0 1 .214 0l3.256.591c.599.11.96-.64.502-1.041l-2.16-1.89a.6.6 0 0 1-.205-.452v-2.96a.6.6 0 0 1 .745-.582l6.51 1.628a.6.6 0 0 0 .745-.582v-.9a.6.6 0 0 0-.282-.508l-7.436-4.648a.6.6 0 0 1-.282-.509V4.5a1.5 1.5 0 0 0-3 0Z"/>`
	airplaneHelixInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M12 9s-1.988-1.975-2-4c.001-1.993-.05-4.001 2-4c1.948.001 1.997 1.976 2 4c.003 1.985-2 4-2 4Zm3 3s1.975-1.988 4-2c1.993.001 4.001-.05 4 2c-.001 1.948-1.976 1.997-4 2c-1.985.003-4-2-4-2Zm-6 0s-1.975 1.988-4 2c-1.993-.001-4.001.05-4-2c.001-1.948 1.976-1.997 4-2c1.985-.003 4 2 4 2Zm3 3s1.988 1.975 2 4c-.001 1.993.05 4.001-2 4c-1.948-.001-1.997-1.976-2-4c-.003-1.985 2-4 2-4Z" clip-rule="evenodd"/></g>`
	airplaneHelixFortyFiveDegInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M14.12 14.121A3 3 0 1 0 9.879 9.88a3 3 0 0 0 4.243 4.242Z"/><path d="M9.879 9.879s-2.803.009-4.243-1.415c-1.409-1.41-2.864-2.793-1.414-4.242c1.378-1.377 2.81-.015 4.242 1.414C9.87 7.037 9.88 9.879 9.88 9.879Zm4.242 0s-.009-2.803 1.415-4.243c1.41-1.409 2.793-2.864 4.242-1.414c1.377 1.378.015 2.81-1.414 4.242c-1.402 1.406-4.243 1.415-4.243 1.415Zm-4.242 4.242s.009 2.803-1.415 4.243c-1.41 1.409-2.793 2.864-4.242 1.414c-1.377-1.378-.015-2.81 1.414-4.242c1.401-1.406 4.243-1.415 4.243-1.415Zm4.242 0s2.803-.009 4.243 1.415c1.409 1.41 2.864 2.793 1.414 4.242c-1.378 1.377-2.81.015-4.242-1.414c-1.406-1.402-1.415-4.243-1.415-4.243Z" clip-rule="evenodd"/></g>`
	airplaneOffInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.881 9.887l-7.099 4.437a.6.6 0 0 0-.282.508v.9a.6.6 0 0 0 .746.582l6.508-1.628a.6.6 0 0 1 .746.582v2.96a.6.6 0 0 1-.205.451l-2.16 1.89c-.458.402-.097 1.151.502 1.042l3.256-.591a.6.6 0 0 1 .214 0l3.256.591c.599.11.96-.64.502-1.041l-2.16-1.89a.6.6 0 0 1-.205-.452v-2.96a.6.6 0 0 1 .745-.582l.458.115M10.5 7.5v-3A1.5 1.5 0 0 1 12 3v0a1.5 1.5 0 0 1 1.5 1.5v4.667a.6.6 0 0 0 .282.51l7.436 4.647a.6.6 0 0 1 .282.508v.9a.6.6 0 0 1-.745.582l-2.006-.502M3 3l18 18"/>`
	airplaneRotationInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M9.879 14.122a3 3 0 1 0 4.242-4.243a3 3 0 0 0-4.242 4.243Z"/><path stroke-width="1.497" d="M4.37 16.773A8.956 8.956 0 0 1 3.002 12c0-4.236 2.934-7.792 6.878-8.747A8.998 8.998 0 0 1 12 3.002m7.715 4.365A8.953 8.953 0 0 1 20.999 12c0 3.806-2.368 7.063-5.709 8.378c-1.02.4-2.13.621-3.29.621"/><path d="M14.121 9.88s-.009-2.803 1.415-4.243c1.41-1.409 2.793-2.865 4.242-1.415c1.377 1.378.015 2.81-1.414 4.243c-1.402 1.406-4.243 1.414-4.243 1.414Zm-4.242 4.24s.009 2.803-1.415 4.243c-1.41 1.409-2.793 2.865-4.242 1.415c-1.377-1.378-.015-2.81 1.414-4.243c1.402-1.406 4.243-1.414 4.243-1.414Z" clip-rule="evenodd"/></g>`
	airplayInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 17H3V4h18v13h-3"/><path d="M8.622 19.067L11.5 14.75a.6.6 0 0 1 .998 0l2.88 4.318a.6.6 0 0 1-.5.933H9.12a.6.6 0 0 1-.5-.933Z"/></g>`
	alarmInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 13h-5V8M5 3.5L7 2m12 1.5L17 2"/><path d="M12 22a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z"/></g>`
	albumInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" d="M12 15.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V7.6a.6.6 0 0 1 .6-.6H15"/></g>`
	albumCarouselInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 19.4V4.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" d="M22 6v12m-11-3.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V8.6a.6.6 0 0 1 .6-.6H13"/></g>`
	albumListInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 17.4V2.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" d="M8 22h13.4a.6.6 0 0 0 .6-.6V8m-11 4.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V6.6a.6.6 0 0 1 .6-.6H13"/></g>`
	albumOpenInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M15 2.2c4.564.927 8 4.962 8 9.8c0 4.838-3.436 8.873-8 9.8"/><path stroke-linejoin="round" d="M15 9c1.141.284 2 1.519 2 3s-.859 2.716-2 3M1 2h10v20H1"/><path d="M4 15.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V7.6a.6.6 0 0 1 .6-.6H7"/></g>`
	alignBottomBoxInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 8l.01.011M4 4l.01.011M8 4l.01.011M12 4l.01.011M16 4l.01.011M20 4l.01.011M20 8l.01.011M4 12v8h16v-8H4Z"/>`
	alignCenterInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h18M3 14h18M6 10h12M6 18h12"/>`
	alignJustifyInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h18M3 10h18M3 14h18M3 18h18"/>`
	alignLeftInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 10h14M3 6h18M3 18h14M3 14h18"/>`
	alignLeftBoxInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.004 3.995l-.011.01m4.011-.01l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m-3.989-.01l-.011.01m-3.987-16.01h-8v16h8v-16Z"/>`
	alignRightInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 10h14M3 6h18M7 18h14M3 14h18"/>`
	alignRightBoxInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.006 20.005l.01-.01m-4.01.01l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m3.99.01l.01-.01m3.99 16.01h8v-16h-8v16Z"/>`
	alignTopBoxInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 16l.01-.011M4 20l.01-.011M8 20l.01-.011M12 20l.01-.011M16 20l.01-.011M20 20l.01-.011M20 16l.01-.011M4 12V4h16v8H4Z"/>`
	angleToolInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M3 21V3h6v12h12v6H3Z"/><path d="M13 19v2m-4-2v2M3 7h2m-2 4h2m-2 4h2m12 4v2"/></g>`
	antennaInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M16 1s1.5 1 1.5 3S16 7 16 7M8 1S6.5 2 6.5 4S8 7 8 7M7 23l1.111-4M17 23l-1.111-4M14.5 14L12 5l-2.5 9m5 0h-5m5 0l1.389 5M9.5 14l-1.389 5m0 0h7.778"/></g>`
	antennaOffInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="m7 23l1.111-4M17 23l-1.111-4M9.5 14l-1.389 5M9.5 14h4m-4 0l.8-2.88M8.11 19h7.778m0 0l-1.184-4.264M11.444 7L12 5l1.047 3.768M3 3l18 18"/></g>`
	antennaSignalInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 8S19 9.5 19 12s-1.5 4-1.5 4m3-11S23 7.5 23 12s-2.5 7-2.5 7M6.5 8S5 9.5 5 12s1.5 4 1.5 4m-3-11S1 7.5 1 12s2.5 7 2.5 7"/><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	antennaSignalTagInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 9s1 1.125 1 3s-1 3-1 3m-3-2.99l.01-.011M17 7s2 1.786 2 5s-2 5-2 5M9 9s-1 1.125-1 3s1 3 1 3M7 7s-2 1.786-2 5s2 5 2 5"/></g>`
	appNotificationInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm2 4v3a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6V9a6 6 0 0 1 6-6h3"/>`
	appWindowInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 19V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/></g>`
	appleInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m12.147 21.265l-.147-.03l-.147.03c-2.377.475-4.62.21-6.26-1.1C3.964 18.86 2.75 16.373 2.75 12c0-4.473 1.008-6.29 2.335-6.954c.695-.347 1.593-.448 2.735-.317c1.141.132 2.458.488 3.943.983l.26.086l.255-.102c2.482-.992 4.713-1.373 6.28-.641c1.47.685 2.692 2.538 2.692 6.945c0 4.374-1.213 6.86-2.843 8.164c-1.64 1.312-3.883 1.576-6.26 1.1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5.5C12 3 11 2 9 2"/></g>`
	appleHalfInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m12.147 21.265l-.147-.03l-.147.03c-2.377.475-4.62.21-6.26-1.1C3.964 18.86 2.75 16.373 2.75 12c0-4.473 1.008-6.29 2.335-6.954c.695-.347 1.593-.448 2.735-.317c1.141.132 2.458.488 3.943.983l.26.086l.255-.102c2.482-.992 4.713-1.373 6.28-.641c1.47.685 2.692 2.538 2.692 6.945c0 4.374-1.213 6.86-2.843 8.164c-1.64 1.312-3.883 1.576-6.26 1.1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5.5C12 3 11 2 9 2"/><path d="M12 6v15"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 12v2"/></g>`
	appleHalfAltInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m12.147 21.265l-.147-.03l-.147.03c-2.377.475-4.62.21-6.26-1.1C3.964 18.86 2.75 16.373 2.75 12c0-4.473 1.008-6.29 2.335-6.954c.695-.347 1.593-.448 2.735-.317c1.141.132 2.458.488 3.943.983l.26.086l.255-.102c2.482-.992 4.713-1.373 6.28-.641c1.47.685 2.692 2.538 2.692 6.945c0 4.374-1.213 6.86-2.843 8.164c-1.64 1.312-3.883 1.576-6.26 1.1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5.5C12 3 11 2 9 2"/><path d="M12 6v15"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 12v2"/></g>`
	appleImacTwoThousandTwentyOneInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 15.5V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.9m-20 0v1.9a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6v-1.9m-20 0h20M9 22h1.5m0 0v-4m0 4h3m0 0H15m-1.5 0v-4"/>`
	appleImacTwoThousandTwentyOneSideInnerSVG = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 22h2m6 0H8m0 0l2-8.5m0 0L7 2m3 11.5l1.5 5.5m5.5 3h1"/>`
	appleMacInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 2c.363 2.18-1.912 3.83-3.184 4.571c-.375.219-.799-.06-.734-.489C12.299 4.64 13.094 2 16 2Z"/><path d="M9 6.5c.897 0 1.69.2 2.294.42a3.58 3.58 0 0 0 2.412 0A6.73 6.73 0 0 1 16 6.5c1.085 0 2.465.589 3.5 1.767C16 11 17 15.5 20.269 16.692c-1.044 2.867-3.028 4.808-4.77 4.808c-1.5 0-1.499-.7-2.999-.7s-1.5.7-3 .7c-2.5 0-5.5-4-5.5-9c0-4 3-6 5-6Z"/></g>`
	appleSwiftInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.457 14.59c.446-1.437 1.451-6.75-5.93-11.49a.636.636 0 0 0-.808.1a.593.593 0 0 0-.022.79c.03.036 2.75 3.35 1.783 7.135c-1.673-1.151-8.324-6.423-8.324-6.423L11 11L3.862 6.4s5.046 6.195 8.134 8.525c-1.495.537-4.743 1.105-9.033-1.561a.637.637 0 0 0-.771.074a.593.593 0 0 0-.106.743C2.229 14.42 5.668 20 12.939 20c1.995 0 3.16-.568 4.098-1.024c.576-.279 1.031-.501 1.528-.501c1.236 0 2.047 1.227 2.054 1.238a.632.632 0 0 0 .583.285a.62.62 0 0 0 .526-.37c.893-2.074-.645-4.269-1.271-5.039Z"/>`
	appleWalletInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2Z"/><path d="M3 15h6.4c.331 0 .605.278.75.576c.206.423.694.924 1.85.924c1.156 0 1.644-.5 1.85-.924c.145-.298.419-.576.75-.576H21M3 7h18M3 11h18"/></g>`
	arSymbolInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 15.5v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H13v4.2m2.857 0H13m2.857 0L18 15.5m-7 0L9.929 13M5 15.5L6.071 13m0 0L8 8.5L9.929 13M6.07 13h3.86"/></g>`
	arcadeInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11 8.5L9.8 9l-7.448 3.386a.6.6 0 0 0-.352.546v.136a.6.6 0 0 0 .352.546l8.82 4.01a2 2 0 0 0 1.656 0l8.82-4.01a.6.6 0 0 0 .352-.546v-.136a.6.6 0 0 0-.352-.546L14.2 9L13 8.5"/><path d="M22 13v4.112a.6.6 0 0 1-.354.547l-8.825 3.972a2 2 0 0 1-1.642 0l-8.825-3.972A.6.6 0 0 1 2 17.112V13"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/><path d="M11 8v5a1 1 0 1 0 2 0V8"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 13h1"/></g>`
	archeryInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 12h11M7 12l-2-2H1l2 2l-2 2h4l2-2Zm11 0l-2-2m2 2l-2 2m1.5 8c3.038 0 5.5-4.477 5.5-10S20.538 2 17.5 2S12 6.477 12 12s2.462 10 5.5 10Z"/>`
	archeryMatchInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.611 15.89l12.02-12.022M8.612 15.89H5.783l-2.829 2.829h2.829v2.828l2.828-2.828V15.89Zm12.02-12.02h-2.828m2.829 0v2.828M15.39 15.89L3.367 3.867M15.39 15.89h2.829l2.828 2.829h-2.828v2.828l-2.829-2.828V15.89ZM3.37 3.87h2.828m-2.829 0v2.828"/>`
	archiveInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 6h10M7 9h10m-8 8h6"/><path d="M3 12h-.4a.6.6 0 0 0-.6.6v8.8a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6v-8.8a.6.6 0 0 0-.6-.6H21M3 12V2.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V12M3 12h18"/></g>`
	areaSearchInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.124 20.119a3 3 0 1 0-4.248-4.237a3 3 0 0 0 4.248 4.237Zm0 0L22 22M7 2H4v3m0 6v2m7-11h2m-2 20h2m7-11v2M17 2h3v3M7 22H4v-3"/>`
	arrowArcheryInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.611 15.89l12.02-12.022M8.612 15.89H5.783l-2.829 2.829h2.829v2.828l2.828-2.828V15.89Zm12.02-12.02h-2.828m2.829 0v2.828"/>`
	arrowBlCircleInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.828 9.172l-5.656 5.656m0 0h4.95m-4.95 0v-4.95M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	arrowBlSquareInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.828 9.172l-5.656 5.656m0 0h4.95m-4.95 0v-4.95M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	arrowBrCircleInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.171 9.172l5.657 5.656m0 0h-4.95m4.95 0v-4.95M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	arrowBrSquareInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.171 9.172l5.657 5.656m0 0h-4.95m4.95 0v-4.95M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	arrowDownInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 5.5V18m0 0l-6-6m6 6l6-6"/>`
	arrowDownCircleInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v8m0 0l3.5-3.5M12 16l-3.5-3.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	arrowEmailForwardInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 10H8c-8 0-8 11 0 11m14-11l-7-7m7 7l-7 7"/>`
	arrowLeftInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.5 12H6m0 0l6-6m-6 6l6 6"/>`
	arrowLeftCircleInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 12H8m0 0l3.5 3.5M8 12l3.5-3.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	arrowRightInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h12.5m0 0l-6-6m6 6l-6 6"/>`
	arrowRightCircleInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h8m0 0l-3.5-3.5M16 12l-3.5 3.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	arrowSeparateInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8L6 11.5L9.5 15M14 8l3.5 3.5L14 15"/>`
	arrowSeparateVerticalInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 9.5L12 6L8.5 9.5m7 4.5L12 17.5L8.5 14"/>`
	arrowTlCircleInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.828 14.828L9.172 9.172m0 0h4.95m-4.95 0v4.95M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	arrowTlSquareInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.828 14.828L9.172 9.172m0 0h4.95m-4.95 0v4.95M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	arrowTrCircleInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.171 14.828l5.657-5.656m0 0h-4.95m4.95 0v4.95M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	arrowTrSquareInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.171 14.828l5.657-5.656m0 0h-4.95m4.95 0v4.95M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	arrowUnionInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.5 8L14 11.5l3.5 3.5M6 8l3.5 3.5L6 15"/>`
	arrowUnionVerticalInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 6L12 9.5L8.5 6m7 11.5L12 14l-3.5 3.5"/>`
	arrowUpInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 18.5V6m0 0l6 6m-6-6l-6 6"/>`
	arrowUpCircleInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16V8m0 0l3.5 3.5M12 8l-3.5 3.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	asanaInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm-5 9a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm10 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`
	atSignInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 19.125A9 9 0 1 1 21 12c0 5.5-6 5.5-6 2V8"/><path d="M15 12v-1.5C15 9.12 13.657 8 12 8s-3 1.12-3 2.5V12m6 0v1.5c0 1.38-1.343 2.5-3 2.5s-3-1.12-3-2.5V12"/></g>`
	atSignCircleInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.278 17.541A7 7 0 1 1 19 12c0 4.278-5 3.722-5 1V8.5"/><path d="M14 12v-.5a2.5 2.5 0 0 0-5 0v.5m5 0v.5a2.5 2.5 0 0 1-5 0V12"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Z"/></g>`
	atomInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.404 13.61C3.515 13.145 3 12.592 3 12c0-1.657 4.03-3 9-3s9 1.343 9 3c0 .714-.75 1.37-2 1.886m-7-2.876l.01-.011"/><path d="M16.883 6c-.005-1.023-.263-1.747-.797-2.02c-1.477-.751-4.503 2.23-6.76 6.658c-2.256 4.429-2.889 8.629-1.412 9.381c.527.269 1.252.061 2.07-.519"/><path d="M9.6 4.252c-.66-.386-1.243-.497-1.686-.271c-1.477.752-.844 4.952 1.413 9.38c2.256 4.43 5.282 7.41 6.758 6.658c1.313-.669.959-4.061-.72-7.917"/></g>`
	attachmentInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21.438 11.662l-9.19 9.19a6.003 6.003 0 1 1-8.49-8.49l9.19-9.19a4.002 4.002 0 0 1 5.66 5.66l-9.2 9.19a2.001 2.001 0 1 1-2.83-2.83l8.49-8.48"/>`
	augmentedRealityInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m5.5 15.5l.614-1.718M10.5 15.5l-.614-1.718m-3.772 0L8 8.5l1.886 5.282m-3.772 0h3.772M13 15.5v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H13v4.2m2.857 0H13m2.857 0L18 15.5"/><path d="M2 18.4V5.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`
	autoFlashInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16 9.5l.692-1.5M22 9.5L21.308 8m0 0L19 3l-2.308 5m4.616 0h-4.616M13 10h-3V3L2 14h6v7l6-8.25"/>`
	aviFormatInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m12 9l1.5 6L15 9m3 6V9"/><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M6 15v-3m0 0v-1.5A1.5 1.5 0 0 1 7.5 9v0A1.5 1.5 0 0 1 9 10.5V12m-3 0h3m0 0v3"/></g>`
	axesInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 19.452l-9-6.61m0 0V3m0 9.843l-9 6.609m17.438-2.742L21 19.452L18.187 20M9.75 5.194L12 3l2.25 2.194M5.813 20L3 19.452l.563-2.742"/>`
	backwardFifteenSecondsInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 13a9 9 0 1 0 9-9M9 9v7"/><path d="M15 9h-2a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h1a1 1 0 0 1 1 1V15a1 1 0 0 1-1 1h-2m0-12H4.5m0 0l2-2m-2 2l2 2"/></g>`
	bagInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M4.508 20h14.984a.6.6 0 0 0 .592-.501l1.8-10.8A.6.6 0 0 0 21.292 8H2.708a.6.6 0 0 0-.592.699l1.8 10.8a.6.6 0 0 0 .592.501Z"/><path d="M7 8V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"/></g>`
	bankInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9.5L12 4l9 5.5M5 20h14M10 9h4m-8 8v-5m4 5v-5m4 5v-5m4 5v-5"/>`
	barcodeInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 19V5h1m6 14V5h1M9 5v14m7-14v14m3-14v14M6 5v14H5m8-14v14h-1"/>`
	basketballInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0 0V2"/><path d="M21.95 11c-6.47 2.667-12.254 2.667-19.9 0"/><path d="M18.572 4.462c-2.667 4.53-2.667 9.723 0 15.076M5.428 4.462c2.667 4.53 2.667 9.723 0 15.076"/></g>`
	basketballAltInnerSVG                     = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirBasketballAlt0)"><path d="M17.736 20.192c4.524-3.168 5.623-9.404 2.455-13.928C17.024 1.74 10.788.641 6.264 3.81C1.74 6.976.641 13.212 3.808 17.736c3.168 4.524 9.404 5.623 13.928 2.456Zm0 0L6.264 3.809"/><path d="M19.577 5.473c-3.77 5.896-8.508 9.214-16.302 11.415"/><path d="M13.06 2.056c.413 5.24 3.392 9.494 8.646 12.35M2.293 9.595c4.783 2.18 7.761 6.434 8.647 12.349"/></g><defs><clipPath id="iconoirBasketballAlt0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	basketballFieldInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 5h9.4a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H12m0-14H2.6a.6.6 0 0 0-.6.6v12.8a.6.6 0 0 0 .6.6H12m0-14v14"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6ZM2 17A5 5 0 0 0 2 7m20 10a5 5 0 0 1 0-10"/></g>`
	batteryChargingInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.167 9L8.5 12h4l-1.667 3"/></g>`
	batteryEmptyInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/></g>`
	batteryFiftyInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path d="M4 14.4V9.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/></g>`
	batteryFullInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path d="M4 14.4V9.6a.6.6 0 0 1 .6-.6h11.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/></g>`
	batteryIndicatorInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 13h4M6 13h2m2 0H8m0 0v-2m0 2v2M6 7H2.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H18M6 7V5h2v2M6 7h2m0 0h8m0 0V5h2v2m-2 0h2"/>`
	batterySeventyFiveInnerSVG                = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path d="M4 14.4V9.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/></g>`
	batteryTwentyFiveInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path d="M4 14.4V9.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/></g>`
	batteryWarningInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.5 9v2m0 4.01l.01-.011"/></g>`
	bbqInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.5 14.5L5 22M8 6s1-1.061 1-2c0-1.333-1-2-1-2m4 4s1-1.061 1-2c0-1.333-1-2-1-2m4 4s1-1.061 1-2c0-1.333-1-2-1-2"/><path stroke-linejoin="round" d="M16.5 17.5h-9"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.5 14.5l2.1 4.5m.9 3a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/><path d="M12 15a7 7 0 0 0 6.975-6.4a.563.563 0 0 0-.575-.6H5.6a.563.563 0 0 0-.575.6A7 7 0 0 0 12 15Z"/></g>`
	beachBagInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m2.77 13l-.633-3.287A.6.6 0 0 1 2.727 9h18.547a.6.6 0 0 1 .589.713L21.23 13M2.769 13h18.462M2.769 13l.616 4m17.846-4l-.616 4m0 0l-.537 3.491a.6.6 0 0 1-.593.509H4.515a.6.6 0 0 1-.593-.509L3.385 17m17.23 0H3.385"/><path d="M8 9V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4"/></g>`
	beachBagBigInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="m2.77 12l-.633-3.287A.6.6 0 0 1 2.727 8h18.547a.6.6 0 0 1 .589.713L21.23 12M2.769 12h18.462M2.769 12l.616 4m17.846-4l-.616 4m0 0l-.537 3.491a.6.6 0 0 1-.593.509H4.515a.6.6 0 0 1-.593-.509L3.385 16m17.23 0H3.385M5 8V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2"/>`
	bedInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 4v16a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2Z"/><path d="M3 8h8V6m10 2h-8V6"/></g>`
	bedReadyInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 16V4a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h9"/><path d="M3 8h8V6m10 2h-8V6m3 14l2 2l4-4"/></g>`
	behanceInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.197 11.217c5.07 0 5.07 6.783 0 6.783H2v-6.783m6.197 0H2m6.197 0c5.07 0 5.07-6.217 0-6.217H2v6.217M18 9c-2.21 0-4 2.015-4 4.5h8c0-2.485-1.79-4.5-4-4.5Zm-4 4.5c0 2.485 1.79 4.5 4 4.5c2.755 0 3.5-2 3.5-2m-1-10h-5"/>`
	behanceTagInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5Z"/><path d="M9.099 11.826c2.535 0 2.535 4.174 0 4.174H6v-4.174m3.099 0H6m3.099 0c2.535 0 2.535-3.826 0-3.826H6v3.826M15.5 11a2.5 2.5 0 0 0-2.5 2.5h5a2.5 2.5 0 0 0-2.5-2.5ZM13 13.5a2.5 2.5 0 0 0 2.5 2.5c.928 0 1.49-.322 1.813-.62M17 8.5h-3"/></g>`
	bellInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 8.4c0-1.697-.632-3.325-1.757-4.525C15.117 2.675 13.59 2 12 2c-1.591 0-3.117.674-4.243 1.875C6.632 5.075 6 6.703 6 8.4C6 15.867 3 18 3 18h18s-3-2.133-3-9.6ZM13.73 21a1.999 1.999 0 0 1-3.46 0"/>`
	bellNotificationInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.134 11C18.715 16.375 21 18 21 18H3s3-2.133 3-9.6c0-1.697.632-3.325 1.757-4.525C8.883 2.675 10.41 2 12 2c.337 0 .672.03 1 .09M19 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm-5.27 13a1.999 1.999 0 0 1-3.46 0"/>`
	bellOffInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.27 6.5C6.093 7.11 6 7.75 6 8.4C6 15.867 3 18 3 18h15M7.757 3.875C8.883 2.675 10.41 2 12 2c1.591 0 3.117.674 4.243 1.875C17.368 5.075 18 6.703 18 8.4c0 7.467 3 9.6 3 9.6m-7.27 3a1.999 1.999 0 0 1-3.46 0M3 3l18 18"/>`
	bicycleInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM8.5 7.5h6M19 15l-4-7.5h-.5m0 0l2-3m0 0H14m2.5 0h2"/><path d="m5 15l3.5-7.5L12 14h3M8.5 7.5c-.333-1-1.5-3-3.5-3"/><path d="M19 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`
	binInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944l-1.7-12.744Z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`
	binAddInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.992 14h3m3 0h-3m0 0v-3m0 3v3"/><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944l-1.7-12.744Z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`
	binFullInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m19.262 17.038l1.676-12.575a.6.6 0 0 0-.372-.636L16 2h-5.5l-.682 1.5L5 2L3.21 3.79a.6.6 0 0 0-.17.504l1.698 12.744a4 4 0 0 0 1.98 2.944l.32.183a10 10 0 0 0 9.923 0l.32-.183a4 4 0 0 0 1.98-2.944ZM16 2l-2 5m-5-.5l.818-3"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`
	binHalfInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944l-1.7-12.744Z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0M11 18l3-3.5m0 0l5 2.5m-5-2.5l6-3M4.5 16l3.236-.462a.6.6 0 0 1 .469.133L11 18l3 3m-6-5.5l2.615-3.05a.6.6 0 0 1 .84-.071L14 14.5"/></g>`
	binMinusInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.992 13h6"/><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944l-1.7-12.744Z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`
	bishopInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 17h10m-8-5h6m-3-9V2m-.446 1.582c-.921 1.251-2.916 4.243-2.497 6.168C9.451 11.558 11.02 12 12 12c.981 0 2.549-.442 2.943-2.25c.42-1.925-1.576-4.917-2.497-6.168a.548.548 0 0 0-.892 0ZM17.8 22H6.2a.617.617 0 0 1-.5-.97c1.316-1.866 4.063-5.986 4.493-8.434c.057-.326.326-.596.657-.596h2.3c.331 0 .6.27.657.596c.43 2.448 3.177 6.568 4.492 8.434a.617.617 0 0 1-.499.97Z"/>`
	bitbucketInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20.916 4.674l-1.85 14.8a.6.6 0 0 1-.596.526H5.53a.6.6 0 0 1-.596-.526l-1.85-14.8A.6.6 0 0 1 3.68 4h16.64a.6.6 0 0 1 .596.674Z"/><path d="m16.75 7.75l-.938 7.97a.6.6 0 0 1-.595.53H8.784a.6.6 0 0 1-.596-.53l-.859-7.3a.6.6 0 0 1 .596-.67h8.825Zm0 0h3.75"/></g>`
	bluetoothInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.75 8l10.5 8.5l-5.5 5.5V2l5.5 5.5L6.75 16"/>`
	bluetoothTagInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m9 9.6l6 5.1l-3.143 3.3V6L15 9.3l-6 5.1"/><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/></g>`
	boldInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M12 11.667H8m4 0s3.333 0 3.333-3.334C15.333 5 12 5 12 5H8.6a.6.6 0 0 0-.6.6v6.067m4 0s4 0 4 3.666C16 19 12 19 12 19H8.6a.6.6 0 0 1-.6-.6v-6.733"/>`
	boldSquareInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path d="M12 12H9m3 0s2.5 0 2.5-2.5S12 7 12 7H9.6a.6.6 0 0 0-.6.6V12m3 0s3 0 3 2.75s-3 2.75-3 2.75H9.6a.6.6 0 0 1-.6-.6V12"/></g>`
	bonfireInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M9 14c0 1.61 1.377 2 3.076 2c2.89 0 3.845-1.667 1.922-5c-2.691 3-3.076-1.667-2.691-3C10.153 10 9 11.879 9 14Z"/><path stroke-linejoin="round" d="M12 16c3.156 0 5-2.098 5-5.688S12 3 12 3s-5 3.723-5 7.313S8.844 16 12 16Z"/><path d="m4.273 21.07l15.454-4.14m-15.454 0L12 19m7.727 2.07l-3.863-1.035"/></g>`
	bookInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/><path d="M9 7h6"/></g>`
	bookStackInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M5 19.5V5a2 2 0 0 1 2-2h11.4a.6.6 0 0 1 .6.6V21M9 7h6m-8.5 8H19M6.5 18H19M6.5 21H19"/><path stroke-linejoin="round" d="M6.5 18c-1 0-1.5-.672-1.5-1.5S5.5 15 6.5 15m0 6c-1 0-1.5-.672-1.5-1.5S5.5 18 6.5 18"/></g>`
	bookmarkBookInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M8 3v8l2.5-1.6L13 11V3"/><path d="M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/></g>`
	bookmarkCircleInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16v-6a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v6l-1.89-1.26a2 2 0 0 0-2.22 0L9 16Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	bookmarkEmptyInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 21V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16l-5.918-3.805a2 2 0 0 0-2.164 0L5 21Z"/>`
	borderBlInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4v16h16m.01-4l-.011.01M20.01 12l-.011.01M20.01 8l-.011.01M20.01 4l-.011.01M8.01 4l-.011.01M12.01 4l-.011.01M16.01 4l-.011.01"/>`
	borderBottomInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 4.01l.01-.011M16 4.01l.01-.011M12 4.01l.01-.011M8 4.01l.01-.011M4 4.01l.01-.011M4 8.01l.01-.011M4 12.01l.01-.011m7.99.011l.01-.011M4 16.01l.01-.011M20 8.01l.01-.011M20 12.01l.01-.011M20 16.01l.01-.011M4 20h16"/>`
	borderBrInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.01 4v16h-16M4 16l.011.01M4 12l.011.01M4 8l.011.01M4 4l.011.01M16 4l.011.01M12 4l.011.01M8 4l.011.01"/>`
	borderInnerInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 20.01l.01-.011M4 16.01l.01-.011M4 8.01l.01-.011M4 4.01l.01-.011M8 4.01l.01-.011M16 4.01l.01-.011M20 4.01l.01-.011M20 8.01l.01-.011M8 20.01l.01-.011m7.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M4 12h8m8 0h-8m0 0V4m0 8v8"/>`
	borderLeftInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.01 20l-.011.01M20.01 16l-.011.01M20.01 12l-.011.01M20.01 8l-.011.01M20.01 4l-.011.01M8.01 4l-.011.01M12.01 4l-.011.01M12.01 12l-.011.01M16.01 4l-.011.01M8.01 20l-.011.01M12.01 20l-.011.01M16.01 20l-.011.01M4 4v16"/>`
	borderOutInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.01 16l-.01.011M12.01 12l-.01.011M12.01 8l-.01.011M8.01 12l-.01.011M16.01 12l-.01.011M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	borderRightInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 20l.01.01M4 16l.01.01M4 12l.01.01M4 8l.01.01M4 4l.01.01M16 4l.01.01M12 4l.01.01M12 12l.01.01M8 4l.01.01M16 20l.01.01M12 20l.01.01M8 20l.01.01M20.01 4v16"/>`
	borderTlInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M4 20V4h16"/>`
	borderTopInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 20.01l.01-.011M16 20.01l.01-.011M12 20.01l.01-.011M8 20.01l.01-.011M4 20.01l.01-.011M4 8.01l.01-.011M4 12.01l.01-.011m7.99.011l.01-.011M4 16.01l.01-.011M20 8.01l.01-.011M20 12.01l.01-.011M20 16.01l.01-.011M4 4h16"/>`
	borderTrInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.01 20V4h-16M4 8l.01-.01M4 12l.01-.01M4 16l.01-.01M4 20l.01-.01M16 20l.01-.01M12 20l.01-.01M8 20l.01-.01"/>`
	bounceLeftInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm15 8.5c-3-1-5.5-.5-8 4.5c-.5-3-2-7.5-3.5-10"/>`
	bounceRightInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM4 15.5c3-1 5.5-.5 8 4.5c.5-3 2-7.5 3.5-10"/>`
	bowlingBallInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M11.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-4 3a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm4 2a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/></g>`
	boxInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12h4M3 3h18m0 4v13.4a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V7"/>`
	boxIsoInnerSVG                            = `<g fill="none" stroke-width="1.5"><path fill="currentColor" d="m2.695 7.185l9 4l.61-1.37l-9-4l-.61 1.37ZM12.75 21.5v-11h-1.5v11h1.5Zm-.445-10.315l9-4l-.61-1.37l-9 4l.61 1.37Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 17.11V6.89a.6.6 0 0 1 .356-.548l8.4-3.734a.6.6 0 0 1 .488 0l8.4 3.734A.6.6 0 0 1 21 6.89v10.22a.6.6 0 0 1-.356.548l-8.4 3.734a.6.6 0 0 1-.488 0l-8.4-3.734A.6.6 0 0 1 3 17.11Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 4.5l8.644 3.842a.6.6 0 0 1 .356.548v3.61"/></g>`
	boxingGloveInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.489 17.727h9.867m-9.867 0V21h9.867v-3.273m-9.867 0C5.2 15.546 3.556 10.091 4.104 8.455c.438-1.31 2.375-.91 3.289-.546C7.393 4.091 9.037 3 13.423 3C17.806 3 20 4.09 20 9.545c0 4.364-1.096 7.273-1.644 8.182"/><path d="M7.393 7.91C7.758 8.272 8.818 9 10.133 9h4.934M7.393 7.91c0 3.817 1.644 4.363 2.74 4.363"/></g>`
	brainInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a3 3 0 1 0 1 5.83"/><path d="M4.264 15.605a4 4 0 0 1-.874-6.636m.03-.081A2.5 2.5 0 0 1 7 5.5m.238.065A2.5 2.5 0 1 1 12 4.5V20m-4 0a2 2 0 1 0 4 0m0-13a3 3 0 0 0 3 3m2 4a3 3 0 1 1-1 5.83"/><path d="M19.736 15.605a4 4 0 0 0 .874-6.636m-.03-.081A2.5 2.5 0 0 0 17 5.5m-5-1a2.5 2.5 0 1 1 4.762 1.065M16 20a2 2 0 1 1-4 0"/></g>`
	brainElectricityInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a3 3 0 1 0 1 5.83"/><path d="M4.264 15.605a4 4 0 0 1-.874-6.636m.03-.081A2.5 2.5 0 0 1 7 5.5m.238.065A2.5 2.5 0 1 1 12 4.5V20m-4 0a2 2 0 1 0 4 0m0-13a3 3 0 0 0 3 3m5.61-1.031A3.99 3.99 0 0 1 22 12c0 .703-.181 1.364-.5 1.938m-.92-5.05A2.5 2.5 0 0 0 17 5.5m-5-1a2.5 2.5 0 1 1 4.762 1.065M14 22a2 2 0 0 1-2-2m6.667-4L17 19h4l-1.667 3"/></g>`
	brainResearchInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a3 3 0 1 0 1 5.83"/><path d="M4.264 15.605a4 4 0 0 1-.874-6.636m.03-.081A2.5 2.5 0 0 1 7 5.5m.238.065A2.5 2.5 0 1 1 12 4.5V20m-4 0a2 2 0 1 0 4 0m0-13a3 3 0 0 0 3 3m5.61-1.031A3.99 3.99 0 0 1 22 12c0 .703-.181 1.364-.5 1.938m-.92-5.05A2.5 2.5 0 0 0 17 5.5m-5-1a2.5 2.5 0 1 1 4.762 1.065M14 22a2 2 0 0 1-2-2m8.5.5L22 22m-6-3.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Z"/></g>`
	brainWarningInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a3 3 0 1 0 1 5.83"/><path d="M4.264 15.605a4 4 0 0 1-.874-6.636m.03-.081A2.5 2.5 0 0 1 7 5.5m5-1a2.5 2.5 0 1 0-4.762 1.065M8 20a2 2 0 1 0 4 0m5-6a3 3 0 1 1-1 5.83"/><path d="M19.736 15.605a4 4 0 0 0 .874-6.636m-.03-.081A2.5 2.5 0 0 0 17 5.5m-5-1a2.5 2.5 0 1 1 4.762 1.065M16 20a2 2 0 1 1-4 0m0-12v4m0 4.01l.01-.011"/></g>`
	breadSliceInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 20V9S3 4 9.5 4H17c7 0 3 5 3 5v9a2 2 0 0 1-2 2H7Z"/><path d="M7 20H6a2 2 0 0 1-2-2V9S0 4 6.5 4H10"/></g>`
	brightCrownInnerSVG                       = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirBrightCrown0)"><path d="M22 12h1M12 2V1m0 22v-1m8-2l-1-1m1-15l-1 1M4 20l1-1M4 4l1 1m-4 7h1m14.8 3.5l1.2-7l-4.2 2.1L12 8.5l-1.8 2.1L6 8.5l1.2 7h9.6Z"/></g><defs><clipPath id="iconoirBrightCrown0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	brightStarInnerSVG                        = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirBrightStar0)"><path d="m9.952 9.623l1.559-3.305a.535.535 0 0 1 .978 0l1.559 3.305l3.485.533c.447.068.625.644.302.974l-2.522 2.57l.595 3.631c.077.467-.391.822-.791.602L12 16.218l-3.117 1.715c-.4.22-.868-.135-.791-.602l.595-3.63l-2.522-2.571c-.323-.33-.145-.906.302-.974l3.485-.533ZM22 12h1M12 2V1m0 22v-1m8-2l-1-1m1-15l-1 1M4 20l1-1M4 4l1 1m-4 7h1"/></g><defs><clipPath id="iconoirBrightStar0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	brightnessInnerSVG                        = `<g fill="none" stroke-width="1.5"><path stroke="currentColor" d="m12 7l1.53 1.304l2.006.16l.16 2.005L17 12l-1.305 1.53l-.16 2.006l-2.004.16L12 17l-1.53-1.305l-2.006-.16l-.16-2.004L7 12l1.304-1.53l.16-2.006l2.005-.16L12 7Z"/><path fill="currentColor" d="M10.47 15.695L12 17V7l-1.53 1.304l-2.006.16l-.16 2.005L7 12l1.304 1.53l.16 2.006l2.005.16Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/></g>`
	brightnessWindowInnerSVG                  = `<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" d="M12 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v7"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/><path stroke="currentColor" d="m18 14l1.225 1.044l1.603.128l.128 1.603L22 18l-1.044 1.225l-.128 1.603l-1.603.128L18 22l-1.225-1.044l-1.603-.128l-.128-1.603L14 18l1.044-1.225l.128-1.603l1.603-.128L18 14Z"/><path fill="currentColor" d="M16.775 20.956L18 22v-8l-1.225 1.044l-1.603.128l-.128 1.603L14 18l1.044 1.225l.128 1.603l1.603.128Z"/></g>`
	bubbleDownloadInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 2v6m0 0l3-3m-3 3l-3-3m-3-2.95c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`
	bubbleErrorInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 2.05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1m-4.829-3.636l2.121-2.121m0 0l2.122-2.122m-2.122 2.122l-2.12-2.122m2.12 2.122l2.122 2.121"/>`
	bubbleIncomeInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 5h-6m0 0l3-3m-3 3l3 3m-6-5.95c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`
	bubbleOutcomeInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 5h6m0 0l-3 3m3-3l-3-3m-6 .05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`
	bubbleSearchInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 6.5L22 8m-6-3.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Z"/><path d="M13 2.05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/></g>`
	bubbleStarInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.306 4.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.233.128-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568l2.033-.31Z"/><path d="M13 2.05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/></g>`
	bubbleUploadInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 8V2m0 0l3 3m-3-3l-3 3m-3-2.95c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`
	bubbleWarningInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2v3m0 4.01l.01-.011M13 2.05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`
	buildingInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10 9.01l.01-.011M14 9.01l.01-.011M10 13.01l.01-.011m3.99.011l.01-.011M10 17.01l.01-.011m3.99.011l.01-.011M6 20.4V5.6a.6.6 0 0 1 .6-.6H12V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6Z"/>`
	busInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 16.01l.01-.011m9.99.011l.01-.011M3 12h18v7a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-7Zm18-4V6a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2m4 0h10"/><path d="M4.5 20v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20m7 0v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20"/></g>`
	busStopInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m16 16.01l.01-.011M6 16.01l.01-.011M20 22V8m0 0h-2V2h4v6h-2Zm-4 12H2.6a.6.6 0 0 1-.6-.6v-6.8a.6.6 0 0 1 .6-.6H16m-2-4H6m8-6H6a4 4 0 0 0-4 4v2"/><path d="M3.5 20v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20m7 0v1.9a.6.6 0 0 0 .6.6h.9"/></g>`
	cableTagInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.667 8L10 12h4l-1.667 4"/></g>`
	calculatorInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 21V3a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 7h4m-4 8.5h4m-4 3h4M5 7h2m2 0H7m0 0V5m0 2v2m-1.414 9.414L7 17m1.415-1.414L7 17m0 0l-1.414-1.414M7 17l1.415 1.414"/></g>`
	calendarInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 4V2m0 2v2m0-2h-4.5M3 10v9a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-9H3Zm0 0V6a2 2 0 0 1 2-2h2m0-2v4m14 4V6a2 2 0 0 0-2-2h-.5"/>`
	calendarMinusInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 21H5a2 2 0 0 1-2-2v-9h18v3m-6-9V2m0 2v2m0-2h-4.5M3 10V6a2 2 0 0 1 2-2h2m0-2v4m14 4V6a2 2 0 0 0-2-2h-.5m-3.508 14H21"/>`
	calendarPlusInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 21H5a2 2 0 0 1-2-2v-9h18v3m-6-9V2m0 2v2m0-2h-4.5M3 10V6a2 2 0 0 1 2-2h2m0-2v4m14 4V6a2 2 0 0 0-2-2h-.5m-3.508 14h3M21 18h-3.008m0 0v-3m0 3v3"/>`
	cameraInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 19V9a2 2 0 0 1 2-2h.5a2 2 0 0 0 1.6-.8l2.22-2.96A.6.6 0 0 1 8.8 3h6.4a.6.6 0 0 1 .48.24L17.9 6.2a2 2 0 0 0 1.6.8h.5a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path d="M12 17a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`
	cancelInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.758 17.243L12.001 12m5.243-5.243L12 12m0 0L6.758 6.757M12.001 12l5.243 5.243"/>`
	candlestickChartInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 16v-2m7 7v-2m7-6v-2M5 8V6m7 7v-2m7-6V3M7 8.6v4.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6Zm7 5v4.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6v-4.8a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6Zm7-8v4.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6Z"/>`
	carInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10h8m-9 4h1m8 0h1"/><path d="M3 18v-6.59a2 2 0 0 1 .162-.787l2.319-5.41A2 2 0 0 1 7.319 4h9.362a2 2 0 0 1 1.838 1.212l2.32 5.41a2 2 0 0 1 .161.789V18M3 18v2.4a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V18m-4 0h4m14 0v2.4a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6V18m4 0h-4M7 18h10"/></g>`
	carbonInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M14 10v-.2A1.8 1.8 0 0 0 12.2 8h-.4A1.8 1.8 0 0 0 10 9.8v4.4a1.8 1.8 0 0 0 1.8 1.8h.4a1.8 1.8 0 0 0 1.8-1.8V14"/></g>`
	cardIssueInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 9V5.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v12.8a.6.6 0 0 0 .6.6H12M22 9H6m16 0v4m-1 3.05a3.5 3.5 0 1 0-5 4.9m4.998-4.9A3.5 3.5 0 1 1 16 20.95m5-4.9l-5 4.9"/>`
	cardLockedInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 9V5.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v12.8a.6.6 0 0 0 .6.6H14m8-10H6m16 0v4m-.833 5.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/>`
	cardSecurityInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 9V5.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v12.8a.6.6 0 0 0 .6.6H12M22 9H6m16 0v2"/><path d="m18.992 14.125l2.556.649c.266.068.453.31.445.584C21.821 21.116 18.5 22 18.5 22s-3.321-.884-3.493-6.642a.588.588 0 0 1 .445-.584l2.556-.649c.323-.082.661-.082.984 0Z"/></g>`
	cardWalletInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 20H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2ZM7 7V3.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6V7m-7-4v4m2-4v4"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	cartInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M19.5 22a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-10 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/><path d="M5 4h17l-2 11H7L5 4Zm0 0c-.167-.667-1-2-3-2m18 13H5.23c-1.784 0-2.73.781-2.73 2c0 1.219.946 2 2.73 2H19.5"/></g>`
	cartAltInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M19.5 22a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-10 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/><path d="M16.5 4H22l-2 11h-4.5m1-11l-1 11m1-11h-5.75m4.75 11h-4m-.75-11H5l2 11h4.5m-.75-11l.75 11M5 4c-.167-.667-1-2-3-2m18 13H5.23c-1.784 0-2.73.781-2.73 2c0 1.219.946 2 2.73 2H19.5"/></g>`
	cashInnerSVG                              = `<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14 5h7.4a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H14m0-14v14m0-14h-4m4 14h-4m0 0H2.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6H10m0 14V5"/><path fill="currentColor" d="M7 9.849c0-.414-.413-.699-.75-.457A2.996 2.996 0 0 0 5 11.829c0 1.004.493 1.893 1.25 2.438c.337.242.75-.043.75-.457V9.85Zm10 0c0-.414.413-.699.75-.457A2.996 2.996 0 0 1 19 11.829a2.996 2.996 0 0 1-1.25 2.438c-.337.242-.75-.043-.75-.457V9.85Z"/></g>`
	centerAlignInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 16.01l.01-.011M4 20.01l.01-.011M4 8.01l.01-.011M4 4.01l.01-.011M4 12.01l.01-.011M8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M20 4.01l.01-.011M16 4.01l.01-.011M12 4.01l.01-.011M8 4.01l.01-.011M8 16V8h8v8H8Z"/>`
	chatAddInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3m3 0h-3m0 0V9m0 3v3m0 7c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/>`
	chatBubbleInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M17 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/></g>`
	chatBubbleCheckInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 12l3 3l5-5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/></g>`
	chatBubbleCheckOneInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 12l3 3l5-5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/></g>`
	chatBubbleEmptyInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/>`
	chatBubbleErrorInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.5 14.5l2.493-2.5M14.5 9.5L11.993 12m0 0L9.5 9.5m2.493 2.5l2.507 2.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/>`
	chatBubbleQuestionInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 9c0-3.5 5.5-3.5 5.5 0c0 2.5-2.5 2-2.5 5m0 4.01l.01-.011"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/></g>`
	chatBubbleTranslateInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/><path d="M7 8.517h5m5 0h-1.786m-3.214 0h3.214m-3.214 0V7m3.214 1.517c-.586 2.075-1.813 4.037-3.214 5.76M8.429 18C9.56 16.97 10.84 15.705 12 14.277m0 0c-.714-.829-1.714-2.17-2-2.777m2 2.777l2.143 2.206"/></g>`
	chatBubbleWarningInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v4m0 4.01l.01-.011M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/>`
	chatLinesInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 10h8m-8 4h4m0 8c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22Z"/>`
	chatRemoveInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22ZM9 12h6"/>`
	checkInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 13l4 4L19 7"/>`
	checkCircleInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7 12.5l3 3l7-7"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	checkWindowInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M16 20l2 2l4-4"/></g>`
	chocolateInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 6.5c-3 0-4.5-.5-4.5-3.5H5v18h14V6.5Zm0 8.5H5m0-6h14m-7 12V3"/>`
	chromecastInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2 20.01l.01-.011M15 20h5a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v1m0 9c2 .5 3.5 2 4 4m-4-8c4 .5 7.5 4 8 8"/>`
	chromecastActiveInnerSVG                  = `<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2 20.01l.01-.011M15 20h5a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v1m0 9c2 .5 3.5 2 4 4m-4-8c4 .5 7.5 4 8 8"/><path fill="currentColor" fill-rule="evenodd" d="M5.002 7.63a.6.6 0 0 1 .6-.6h12.804a.6.6 0 0 1 .6.6v8.832a.6.6 0 0 1-.6.6H13.44a.617.617 0 0 1-.556-.355c-.422-.892-1.622-3.26-3.07-4.707c-1.42-1.419-3.572-2.444-4.435-2.82a.624.624 0 0 1-.378-.569v-.98Z" clip-rule="evenodd"/></g>`
	churchInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 6l-7.718 4.824a.6.6 0 0 0-.282.508V21.4a.6.6 0 0 0 .6.6H12m0-16l7.718 4.824a.6.6 0 0 1 .282.508V21.4a.6.6 0 0 1-.6.6H12m0-16V4m0-2v2m-2 0h2m0 0h2m-2 18v-5m4 .01l.01-.011M16 13.01l.01-.011M12 13.01l.01-.011M8 13.01l.01-.011M8 17.01l.01-.011"/>`
	churchAltInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.576 7.424a.6.6 0 0 1 .848 0l3.4 3.4a.6.6 0 0 1 .176.425V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V11.249a.6.6 0 0 1 .176-.425l3.4-3.4ZM8 7V4m0-2v2m0 0H6m2 0h2"/><path d="M12 22h7.4a.6.6 0 0 0 .6-.6V10.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 7H8m0 15v-5m0-3.99l.01-.011"/></g>`
	cinemaOldInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-5-5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Zm0 0v10"/></g>`
	circleInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	cityInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 9.01l.01-.011M11 9.01l.01-.011M7 13.01l.01-.011m3.99.011l.01-.011M7 17.01l.01-.011m3.99.011l.01-.011M15 21H3.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6H9V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6V9m0 12h5.4a.6.6 0 0 0 .6-.6V9.6a.6.6 0 0 0-.6-.6H15m0 12v-4m0-8v4m0 0h2m-2 0v4m0 0h2"/>`
	cleanWaterInnerSVG                        = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirCleanWater0)"><path d="M13 21.57A8.132 8.132 0 0 1 6.25 7.75l5.326-5.326a.6.6 0 0 1 .848 0L17.75 7.75A8.131 8.131 0 0 1 19.74 16M16 20l2 2l4-4"/></g><defs><clipPath id="iconoirCleanWater0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	clipboardCheckInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M8.5 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6m3.5-18H18a2 2 0 0 1 2 2v9"/><path d="M8 6.4V4.5a.5.5 0 0 1 .5-.5c.276 0 .504-.224.552-.496C9.2 2.652 9.774 1 12 1s2.8 1.652 2.948 2.504c.048.272.276.496.552.496a.5.5 0 0 1 .5.5v1.9a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linejoin="round" d="m15.5 20.5l2 2l5-5"/></g>`
	clockInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6h6"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	closedCaptionsInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 15V9a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" d="m10.5 10l-.172-.172a2.828 2.828 0 0 0-2-.828v0A2.828 2.828 0 0 0 5.5 11.828v.344A2.828 2.828 0 0 0 8.328 15v0c.75 0 1.47-.298 2-.828L10.5 14m8-4l-.172-.172a2.828 2.828 0 0 0-2-.828v0a2.828 2.828 0 0 0-2.828 2.828v.344A2.828 2.828 0 0 0 16.328 15v0c.75 0 1.47-.298 2-.828L18.5 14"/></g>`
	closetInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M9 14H8m8 0h-1"/><path d="M12 2h8.4a.6.6 0 0 1 .6.6v18.8a.6.6 0 0 1-.6.6H12m0-20H3.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H12m0-20v20"/></g>`
	cloudInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M12 4c-6 0-6 4-6 6c-1.667 0-5 1-5 5s3.333 5 5 5h12c1.667 0 5-1 5-5s-3.333-5-5-5c0-2 0-6-6-6Z"/>`
	cloudBookAltInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 12h7v10L12 20l-3.5 2V12Z"/><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/></g>`
	cloudCheckInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 18l3 3l5-5"/><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/></g>`
	cloudDesyncInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607m12.42 1.88l-1.768 1.768a4 4 0 0 1-5.656 0l-.354-.353"/><path d="m16.067 21.962l.353-2.475l-2.475.354l2.122 2.121Zm-8.487-5.06l1.768-1.768a4 4 0 0 1 5.657 0l.354.353"/><path d="m7.934 14.427l-.353 2.475l2.474-.354l-2.12-2.121Z"/></g>`
	cloudDownloadInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 13v9m0 0l3.5-3.5M12 22l-3.5-3.5m11.5-.893c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`
	cloudErrorInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 22l3-3m3-3l-3 3m0 0l-3-3m3 3l3 3m5-4.393c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`
	cloudSunnyInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 13c-1.667 0-5 1-5 5s3.333 5 5 5h12c1.667 0 5-1 5-5s-3.333-5-5-5m-6-1a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm7-3h1m-8-7V1m6.5 2.5l-1 1m-12-1l1 1M4 9h1"/>`
	cloudSyncInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607m3.58 1.88l1.768 1.768a4 4 0 0 0 5.657 0l.354-.353"/><path d="m7.934 21.962l-.353-2.475l2.474.354l-2.12 2.121Zm8.364-5.06l-1.768-1.768a4 4 0 0 0-5.657 0l-.353.353"/><path d="m15.944 14.427l.354 2.475l-2.475-.354l2.121-2.121Z"/></g>`
	cloudUploadInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22v-9m0 0l3.5 3.5M12 13l-3.5 3.5M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`
	cluteryInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20h3m3 0H9m0 0v-5m8 5v-8s2.5-1 2.5-3V4.5m-2.5 4v-4M4.5 11c1 2.128 4.5 4 4.5 4s3.5-1.872 4.5-4c1.08-2.297 0-6.5 0-6.5h-9s-1.08 4.203 0 6.5Z"/>`
	codeInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.5 6L10 18.5m-3.5-10L3 12l3.5 3.5m11-7L21 12l-3.5 3.5"/>`
	codeBracketsInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 21H8c-1.105 0-2-.894-2-1.999V14c0-1-1.5-2-1.5-2S6 11 6 10V5a2 2 0 0 1 2-2h1m6 18h1c1.105 0 2-.894 2-1.999V14c0-1 1.5-2 1.5-2S18 11 18 10V5a2 2 0 0 0-2-2h-1"/>`
	codeBracketsSquareInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 17h-.667a2 2 0 0 1-2-2v-1.889C7.333 12.556 6 12 6 12s1.333-.556 1.333-1.111V9a2 2 0 0 1 2-2H10m4 10h.667a2 2 0 0 0 2-2v-1.889C16.667 12.556 18 12 18 12s-1.333-.556-1.333-1.111V9a2 2 0 0 0-2-2H14"/><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/></g>`
	codepenInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 9v6M3 15V9m9 12v-6m0-12v6m0 6L3 9l9-6l9 6l-9 6Z"/><path d="m12 21l-9-6l9-6l9 6l-9 6Z"/></g>`
	coffeeCupInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 11.6V15a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6v-3.4a.6.6 0 0 1 .6-.6h12.8a.6.6 0 0 1 .6.6ZM12 9c0-1 .714-2 2.143-2v0A2.857 2.857 0 0 0 17 4.143V3.5M8 9v-.5a3 3 0 0 1 3-3v0a2 2 0 0 0 2-2V3"/><path d="M16 11h2.5a2.5 2.5 0 0 1 0 5H17"/></g>`
	coinInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M15 8.5c-.685-.685-1.891-1.161-3-1.191M9 15c.644.86 1.843 1.35 3 1.391m0-9.082c-1.32-.036-2.5.561-2.5 2.191c0 3 5.5 1.5 5.5 4.5c0 1.711-1.464 2.446-3 2.391m0-9.082V5.5m0 10.891V18.5"/></g>`
	collageFrameInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M19.4 20H4.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6ZM11 12V4m-7 8h16"/>`
	collapseInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 20l-5-5m0 0v4m0-4h4M4 20l5-5m0 0v4m0-4H5M20 4l-5 5m0 0V5m0 4h4M4 4l5 5m0 0V5m0 4H5"/>`
	colorFilterInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 14.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M16 21.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M8 21.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/></g>`
	colorPickerInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 13.161l5.464-5.464a1 1 0 0 1 1.415 0l2.12 2.12a1 1 0 0 1 0 1.415l-1.928 1.929m-7.071 0l-2.172 2.172a.999.999 0 0 0-.218.327l-1.028 2.496c-.508 1.233.725 2.466 1.958 1.959l2.497-1.028a.998.998 0 0 0 .326-.218l5.708-5.708m-7.071 0h7.071m-.193-9.707l2.121 2.121m4.243 4.243l-2.121-2.121m-2.122-2.122l1.414-1.414a1 1 0 0 1 1.415 0l.707.707a1 1 0 0 1 0 1.414L18.12 7.697m-2.122-2.122l2.122 2.122"/>`
	colorPickerEmptyInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.879 7.697L16 9.817a1 1 0 0 1 0 1.415L8.363 18.87a1.001 1.001 0 0 1-.326.218L5.54 20.114c-1.233.508-2.466-.725-1.958-1.958L4.61 15.66a.999.999 0 0 1 .218-.327l7.636-7.636a1 1 0 0 1 1.415 0Zm0-4.243L16 5.575m4.243 4.243L18.12 7.697M16 5.575l1.413-1.414a1 1 0 0 1 1.414 0l.708.707a1 1 0 0 1 0 1.414L18.12 7.697M16 5.575l2.12 2.122"/>`
	combineInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 9.6v10.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6V9.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6Z"/><path d="M15 3.6v10.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6Z"/></g>`
	communityInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 18v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1M1 18v-1a3 3 0 0 1 3-3v0m19 4v-1a3 3 0 0 0-3-3v0m-8-2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm-8 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm16 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`
	compactDiscInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	compassInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.586 10.586L16.95 7.05l-3.536 6.364L7.05 16.95l3.536-6.364Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	compressInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 12H6m6 10v-6m0 0l3 3m-3-3l-3 3m3-17v6m0 0l3-3m-3 3L9 5"/>`
	compressLinesInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2H6m12 20H6m6-17v5m0 0l3-3m-3 3L9 7m3 12v-5m0 0l3 3m-3-3l-3 3"/>`
	computerInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2 21h15m4 0h1"/><path d="M2 16.4V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`
	consumableInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22.003 3v4.497A.503.503 0 0 1 21.5 8v0a.52.52 0 0 1-.466-.3A10 10 0 0 0 12.003 2c-5.185 0-9.449 3.947-9.95 9"/><path d="M17 10v5a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2Zm-5 1V8"/><path d="M2.05 21v-4.497c0-.278.226-.503.504-.503v0c.2 0 .38.119.466.3a10.001 10.001 0 0 0 9.03 5.7c5.186 0 9.45-3.947 9.951-9"/></g>`
	controlSliderInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m6.755 17.283l-1.429-10A2 2 0 0 1 7.306 5h3.388a2 2 0 0 1 1.98 2.283l-1.429 10A2 2 0 0 1 9.265 19h-.53a2 2 0 0 1-1.98-1.717Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 12h4m16 0H12"/></g>`
	cookieInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M7.5 11a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm5.5 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM11 7.01l.01-.011M8 16.01l.01-.011M16 9.01l.01-.011M17 14.01l.01-.011M13 12.01l.01-.011"/></g>`
	coolingInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6ZM12 7v5m0 5v-5m0 0L7.5 9.5M12 12l4.5 2.5M12 12l4.5-2.5M12 12l-4.5 2.5"/>`
	copyInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.4 20H9.6a.6.6 0 0 1-.6-.6V9.6a.6.6 0 0 1 .6-.6h9.8a.6.6 0 0 1 .6.6v9.8a.6.6 0 0 1-.6.6Z"/><path d="M15 9V4.6a.6.6 0 0 0-.6-.6H4.6a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6H9"/></g>`
	copyrightInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M13.5 9.17a3 3 0 1 0 0 5.659"/></g>`
	cornerBottomLeftInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 8l.01.011M4 4l.01.011M8 4l.01.011M12 4l.01.011M16 4l.01.011M20 4l.01.011M20 8l.01.011M20 12l.01.011M20 16l.01.011M20 20l.01.011M16 20l.01.011M4 12.01v8h8v-8H4Z"/>`
	cornerBottomRightInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.01 8l-.01.011M20.01 4l-.01.011M16.01 4l-.01.011M12.01 4l-.01.011M8.01 4L8 4.011M4.01 4L4 4.011M4.01 8L4 8.011M4.01 12l-.01.011M4.01 16l-.01.011M4.01 20l-.01.011M8.01 20l-.01.011m12.01-8.001v8h-8v-8h8Z"/>`
	cornerTopLeftInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 16.01l.01-.011M4 20.01l.01-.011M8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M20 4.01l.01-.011M16 4.01l.01-.011M4 12V4h8v8H4Z"/>`
	cornerTopRightInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.01 16.01l-.01-.011m.01 4.011l-.01-.011m-3.99.011l-.01-.011m-3.99.011l-.01-.011m-3.99.011L8 19.999m-3.99.011L4 19.999m.01-3.989L4 15.999m.01-3.989L4 11.999m.01-3.989L4 7.999m.01-3.989L4 3.999m4.01.011L8 3.999M20.01 12V4h-8v8h8Z"/>`
	cpuInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 15.4V8.6a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6Z"/><path d="M20 4.6v14.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6ZM17 4V2m-5 2V2M7 4V2m0 18v2m5-2v2m5-2v2m3-5h2m-2-5h2m-2-5h2M4 17H2m2-5H2m2-5H2"/></g>`
	cpuWarningInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 2v4m0 4.01l.01-.011M16 14v1.4a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6H14"/><path d="M20 14v5.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6H14m6 13h2m-5 3v2m-5-2v2m-5-2v2m-3-5H2m2-5H2m2-5H2m10-3V2M7 4V2"/></g>`
	crackedEggInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22a8 8 0 0 0 8-8c0-4.418-3.582-12-8-12S4 9.582 4 14a8 8 0 0 0 8 8Z"/><path d="M9.5 3.5L12 8l-2.5 3l2.5 3.5"/></g>`
	creativeCommonsInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M10.5 9.17a3 3 0 1 0 0 5.659m6.25-5.659a3 3 0 1 0 0 5.659"/></g>`
	creditCardInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 9v9.4a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6V9Zm0 0H6"/>`
	creditCardTwoInnerSVG                     = `<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 9V5.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V9Zm0 0h14"/><rect width="4" height="4" x="15" y="12" fill="currentColor" rx=".6"/></g>`
	creditCardsInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 11.429V19.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6v-2.9m17-5.071V8.6a.6.6 0 0 0-.6-.6H19m3 3.429h-3"/><path d="M19 8v7.9a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V5.1a.6.6 0 0 1 .6-.6h15.8a.6.6 0 0 1 .6.6V8Zm0 0H5.5"/></g>`
	cribInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M3 5v16"/><path d="M3 16h18M3 7h18m-3 9V7m-4 9V7m-4 9V7m-4 9V7M3 19h18"/><path stroke-linecap="round" d="M21 5v16m0-16a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM3 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	cropInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 18H6V3"/><path d="M3 6h15v15"/></g>`
	cropRotateBlInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 14v3a4 4 0 0 0 4 4h4"/><path d="M1.5 16.5L4 14l2.5 2.5M20 11V5a1 1 0 0 0-1-1h-6M8 4h2m10 12v-2M10 2v11a1 1 0 0 0 1 1h11"/></g>`
	cropRotateBrInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 20h3a4 4 0 0 0 4-4v-4"/><path d="M16.5 22.5L14 20l2.5-2.5M14 11V5a1 1 0 0 0-1-1H7M2 4h2m10 12v-2M4 2v11a1 1 0 0 0 1 1h11"/></g>`
	cropRotateTlInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 4H7a4 4 0 0 0-4 4v4"/><path d="M7.5 1.5L10 4L7.5 6.5M20 17v-6a1 1 0 0 0-1-1h-6m-5 0h2m10 12v-2M10 8v11a1 1 0 0 0 1 1h11"/></g>`
	cropRotateTrInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 10V7a4 4 0 0 0-4-4h-4"/><path d="M22.5 7.5L20 10l-2.5-2.5M14 17v-6a1 1 0 0 0-1-1H7m-5 0h2m10 12v-2M4 8v11a1 1 0 0 0 1 1h11"/></g>`
	crownInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.2 17L21 7l-6.3 3L12 7l-2.7 3L3 7l1.8 10h14.4Z"/>`
	crownCircleInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Z"/><path d="m16.8 15.5l1.2-7l-4.2 2.1L12 8.5l-1.8 2.1L6 8.5l1.2 7h9.6Z"/></g>`
	cssThreeInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4 3l1.778 17.09L12 22l6.222-1.91L20 3H4Z"/><path d="M7 7h9.5l-1 10l-3.5 1l-3.5-1l-.25-2.5m7.75-3H7.5"/></g>`
	cursorPointerInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M19.503 9.97c1.204.489 1.112 2.224-.137 2.583l-6.306 1.813l-2.88 5.895c-.57 1.168-2.295.957-2.568-.314L4.677 6.257A1.369 1.369 0 0 1 6.53 4.7l12.973 5.27Z" clip-rule="evenodd"/>`
	cutInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 12h1m4 0h1M6.236 7a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4Zm0 0L19 18M6.236 17a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4Zm0 0L19 6"/>`
	cutAltInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.236 8a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4Zm0 0L16 16m1-4h1m4 0h1M6.236 16a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4Zm0 0L16 8"/>`
	cyclingInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm4 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM6 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm5.5-3l1.5-4l-4.882-2l3-3.5l3 2.5h3.5"/>`
	cylinderInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 2c8 0 8 3 8 3s0 3-8 3s-8-3-8-3s0-3 8-3Zm0 14c8 0 8 3 8 3s0 3-8 3s-8-3-8-3s0-3 8-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 5v14M4 5v14"/></g>`
	dashFlagInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 15l.95-10.454A.6.6 0 0 1 6.548 4h13.795a.6.6 0 0 1 .598.654l-.891 9.8a.6.6 0 0 1-.598.546H5Zm0 0l-.6 6M9 7.5l7 4"/>`
	dashboardInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M15 15.8c0-1.767-3-4.8-3-4.8s-3 3.033-3 4.8s1.343 3.2 3 3.2s3-1.433 3-3.2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 4v4m-8.5-.5l3 3m11 0l3-3M2 17h4m12 0h4"/></g>`
	dashboardDotsInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12 7.01l.01-.011M16 9.01l.01-.011M8 9.01l.01-.011M18 13.01l.01-.011M6 13.01l.01-.011M17 17.01l.01-.011M7 17.01l.01-.011M12 17l1-6m-4.5 9.001H4A9.956 9.956 0 0 1 2 14C2 8.477 6.477 4 12 4s10 4.477 10 10c0 2.252-.744 4.33-2 6.001L15.5 20"/><path d="M12 23a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	dashboardSpeedInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 4v4M4 8l2.5 2.5m11 0L20 8M3 17h3m6 0l1-6m5 6h3M8.5 20.001H4A9.956 9.956 0 0 1 2 14C2 8.477 6.477 4 12 4s10 4.477 10 10c0 2.252-.744 4.33-2 6.001L15.5 20"/><path d="M12 23a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	dataTransferBothInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20V4m0 0l3 3m-3-3l-3 3M7 4v16m0 0l3-3m-3 3l-3-3"/>`
	dataTransferCheckInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 19l3 3l5-5m-5-3V4m0 0l3 3m-3-3l-3 3M7 4v16m0 0l3-3m-3 3l-3-3"/>`
	dataTransferDownInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20v-1m0-4v-1M7 4v16m0 0l-3-3m3 3l3-3m7-7V4m0 0l-3 3m3-3l3 3"/>`
	dataTransferUpInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4v1m0 4v1m10 10V4m0 0l3 3m-3-3l-3 3m-7 7v6m0 0l3-3m-3 3l-3-3"/>`
	dataTransferWarningInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4v1m0 4v1m10 2V4m0 0l3 3m-3-3l-3 3m6 9v2m0 4.01l.01-.011M7 14v6m0 0l3-3m-3 3l-3-3"/>`
	databaseBackupInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3c.592 0 1.135-.021 1.631-.06M18 6v6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6m18.666 5.667C22.048 16.097 20.634 15 18.99 15c-1.758 0-3.252 1.255-3.793 3"/><path d="M20.995 17.667h1.671v0c.185 0 .334-.15.334-.334v-1.888m-7.666 4.888C15.952 21.903 17.366 23 19.01 23c1.758 0 3.252-1.255 3.793-3"/><path d="M17.005 20.333h-1.671v0a.334.334 0 0 0-.334.334v1.888"/></g>`
	databaseExportInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 16v6m0 0l3-3m-3 3l-3-3M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6"/></g>`
	databaseMonitorInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6m14 10h3m-1.5-2.571h2.333V16h-4.666v3.429H19.5Z"/></g>`
	databaseRestoreInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6m15 10v-6m0 0l3 3m-3-3l-3 3"/></g>`
	databaseScriptInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 14V6a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v7"/><path d="M12 21H6a4 4 0 0 1 0-8h12a4 4 0 1 0 4 4v-3"/></g>`
	databaseSettingsInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6m15 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-dasharray=".3 2" d="M19 22a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	databaseStarInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6m13.306 5.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.233.129-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568l2.033-.31Z"/></g>`
	databaseStatsInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path stroke-linejoin="round" d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6"/><path d="M15 21v-2m3 2v-4m3 4v-6"/></g>`
	databaseTagInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.357 12c.714 0 2.143 0 2.143-2s-1.429-2-2.143-2H13.5v4m2.857 0H13.5m2.857 0c.714 0 2.143 0 2.143 2s-1.429 2-2.143 2H13.5v-4M8.357 8H5.5v8h2.857c.714 0 2.143 0 2.143-2v-4c0-2-1.429-2-2.143-2Z"/></g>`
	dbInnerSVG                                = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5 12v6s0 3 7 3s7-3 7-3v-6"/><path d="M5 6v6s0 3 7 3s7-3 7-3V6"/><path d="M12 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Z"/></g>`
	dbCheckInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14 19l3 3l5-5M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6"/></g>`
	dbErrorInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.121 21.364l2.122-2.121m2.121-2.122l-2.121 2.122m0 0L17.12 17.12m2.122 2.122l2.121 2.121M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6"/></g>`
	dbSearchInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 20.5L22 22m-6-3.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0ZM4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6"/></g>`
	dbStarInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6s0-3 7-3s7 3 7 3M4 6s0 3 7 3s7-3 7-3M4 6v6m14-6v6s0 3-7 3s-7-3-7-3m7 9c-7 0-7-3-7-3v-6m13.306 5.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.233.129-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568l2.033-.31Z"/>`
	dbWarningInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 16v2m0 4.01l.01-.011M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Zm0 18c-7 0-7-3-7-3v-6"/></g>`
	deCompressInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 12H6m6 4v6m0 0l3-3m-3 3l-3-3m3-11V2m0 0l3 3m-3-3L9 5"/>`
	deleteCircleInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 14.828L12.001 12m2.828-2.828L12.001 12m0 0L9.172 9.172M12.001 12l2.828 2.828M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	deliveryInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6m8-3V4M8 8H3"/>`
	deliveryTruckInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M8 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M10.05 17H15V6.6a.6.6 0 0 0-.6-.6H1m4.65 11H3.6a.6.6 0 0 1-.6-.6v-4.9"/><path stroke-linejoin="round" d="M2 9h4"/><path d="M15 9h5.61a.6.6 0 0 1 .548.356l1.79 4.028a.6.6 0 0 1 .052.243V16.4a.6.6 0 0 1-.6.6h-1.9M15 17h1"/></g>`
	depthInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 20h20M5 4h14M3 16.01l.01-.011m18 .011l-.01-.011M4 12.01l.01-.011m16 .011l-.01-.011M5 8.01l.01-.011m14 .011L19 7.999M12 7v10m0-10l-1.5 1.5M12 7l1.5 1.5M12 17l-3-3m3 3l3-3"/>`
	designNibInnerSVG                         = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirDesignNib0)"><path d="m17.674 11.408l-1.905 5.715a.6.6 0 0 1-.398.386L3.693 20.98a.6.6 0 0 1-.74-.765L6.745 8.841a.6.6 0 0 1 .34-.365l5.387-2.218a.6.6 0 0 1 .653.13l4.404 4.405a.6.6 0 0 1 .145.615ZM3.296 20.602l6.364-6.364"/><path d="m17.792 11.056l2.828-2.829a2 2 0 0 0 0-2.828L18.5 3.277a2 2 0 0 0-2.829 0l-2.828 2.829m-1.062 6.01a1.5 1.5 0 1 0-2.121 2.122a1.5 1.5 0 0 0 2.121-2.122Z"/></g><defs><clipPath id="iconoirDesignNib0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	designPencilInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Z"/><path d="M8 21.168V14l4-7l4 7v7.168"/><path d="M8 14s1.127 1 2 1s2-1 2-1s1.127 1 2 1s2-1 2-1"/></g>`
	deskInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 7v10M1 7h22M4 10h16m-6 4h6m0-7v10M14 7v10m3-7v1m0 3v1"/>`
	dialpadInnerSVG                           = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.5 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1ZM12 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm6.5-15a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/>`
	diameterInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm7-10l-3-3m3 3l-3 3m3-3H5m0 0l3-3m-3 3l3 3"/>`
	diceFiveInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1ZM12 12.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1ZM7.5 17a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	diceFourInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm-9 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	diceOneInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 12.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	diceSixInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm-9 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm-9 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	diceThreeInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm4.5 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm4.5 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	diceTwoInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm9 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	dimmerSwitchInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm3.5-10.5L12 12"/></g>`
	directorChairInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M19 12L5 21M5 3v9m14-9v9M5 12l14 9M4 12h16"/><path d="M5 4h14M5 7h14"/></g>`
	discordInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.5 16c5 2.5 8 2.5 13 0"/><path d="m15.5 17.5l1 2s4.171-1.328 5.5-3.5c0-1 .53-8.147-3-10.5c-1.5-1-4-1.5-4-1.5l-1 2h-2"/><path d="m8.528 17.5l-1 2s-4.171-1.328-5.5-3.5c0-1-.53-8.147 3-10.5c1.5-1 4-1.5 4-1.5l1 2h2"/><path d="M8.5 14c-.828 0-1.5-.895-1.5-2s.672-2 1.5-2s1.5.895 1.5 2s-.672 2-1.5 2Zm7 0c-.828 0-1.5-.895-1.5-2s.672-2 1.5-2s1.5.895 1.5 2s-.672 2-1.5 2Z"/></g>`
	dishwasherInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-2.5M21 7H3m9 9v5m0 0h-2m2 0h2"/><path d="M12 16c1.657 0 3-1.492 3-3.333V10H9v2.667C9 14.507 10.343 16 12 16Zm6-10.99l.01-.011M15 5.01l.01-.011"/></g>`
	divideInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21v-9a5 5 0 0 0-5-5H3m9 14v-9a5 5 0 0 1 5-5h4"/><path d="M7 3L3 7l4 4m10-8l4 4l-4 4"/></g>`
	divideSelectionOneInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 20H4v-4h16v4ZM2 12h20M7 4H4v3m7-3h2m4 0h3v3"/>`
	divideSelectionTwoInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 12h20M7 4H4v3m7-3h2m4 0h3v3m-9 13h2m-6 0H4v-3m13 3h3v-3"/>`
	divideThreeInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21v-4a5 5 0 0 0-5-5H3m9 9v-4a5 5 0 0 1 5-5h4M12 2v20"/><path d="m6 8l-4 4l4 4M16 6l-4-4l-4 4m10 2l4 4l-4 4"/></g>`
	dnaInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 3c0 5.625 8 9 8 9s8 3.375 8 9"/><path d="M20 3c0 5.625-8 9-8 9s-8 3.375-8 9M8 6h11M8 18h11m-8-9h5.5M11 15h5.5"/></g>`
	docSearchInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 20.5L22 22m-7-4a3 3 0 1 0 6 0a3 3 0 0 0-6 0Z"/><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	docSearchAltInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14 15l1.5 1.5m-7-4a3 3 0 1 0 6 0a3 3 0 0 0-6 0Z"/><path d="M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	docStarInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m16.306 17.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.234.129-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568l2.033-.31Z"/><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	docStarAltInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/><path d="m10.635 10.415l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L12 14.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.217-.22-.098-.604.2-.65l2.324-.355ZM16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	dollarInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.154 7.154c-.949-.949-2.619-1.608-4.154-1.65m-4.154 10.65c.893 1.19 2.552 1.868 4.154 1.926m0-12.576c-1.826-.049-3.461.778-3.461 3.034c0 4.154 7.615 2.077 7.615 6.231c0 2.37-2.027 3.387-4.154 3.31m0-12.575V3m0 15.08V21"/>`
	domoticIssueInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v3m0 4.01l.01-.011M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13"/>`
	donateInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M16 6.28a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.822-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.232 0L12 4.78l.106-.107A2.276 2.276 0 0 1 16 6.28Z"/><path stroke-linecap="round" d="m18 20l3.824-3.824a.6.6 0 0 0 .176-.424V10.5A1.5 1.5 0 0 0 20.5 9v0a1.5 1.5 0 0 0-1.5 1.5V15"/><path stroke-linecap="round" d="m18 16l.858-.858a.484.484 0 0 0 .142-.343v0a.485.485 0 0 0-.268-.433l-.443-.221a2 2 0 0 0-2.308.374l-.895.895a2 2 0 0 0-.586 1.414V20M6 20l-3.824-3.824A.6.6 0 0 1 2 15.752V10.5A1.5 1.5 0 0 1 3.5 9v0A1.5 1.5 0 0 1 5 10.5V15"/><path stroke-linecap="round" d="m6 16l-.858-.858A.485.485 0 0 1 5 14.799v0c0-.183.104-.35.268-.433l.443-.221a2 2 0 0 1 2.308.374l.895.895a2 2 0 0 1 .586 1.414V20"/></g>`
	doubleCheckInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m1.5 12.5l4.076 4.076a.6.6 0 0 0 .848 0L9 14m7-7l-4 4"/><path d="m7 12l4.576 4.576a.6.6 0 0 0 .848 0L22 7"/></g>`
	downRoundArrowInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 5h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V9a4 4 0 0 1 4-4Z"/><path d="m14.5 10.75l-2.5 2.5l-2.5-2.5"/></g>`
	downloadInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20h12M12 4v12m0 0l3.5-3.5M12 16l-3.5-3.5"/>`
	downloadCircleInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 17h6M12 6v7m0 0l3.5-3.5M12 13L8.5 9.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	downloadDataWindowInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M19.5 16v6m0 0L17 19.5m2.5 2.5l2.5-2.5"/></g>`
	downloadSquareInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18h12M12 6v8m0 0l3.5-3.5M12 14l-3.5-3.5"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	dragInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12L4 4m0 0v4m0-4h4m4 8l8-8m0 0v4m0-4h-4m-4 8l-8 8m0 0v-4m0 4h4m4-8l8 8m0 0v-4m0 4h-4"/>`
	dragHandGestureInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7 10.5l-2.004 2.672a2 2 0 0 0 .126 2.552l3.784 4.128c.378.413.912.648 1.473.648H15c2.4 0 4-1.5 4-4c0 0 0 0 0 0V7.929M16 8.5v-.571c0-2.286 3-2.286 3 0"/><path d="M13 8.5V7.027m0-.527v.527M16 8.5V7.027c0-2.286-3-2.286-3 0"/><path d="M13 8.5V7.027c0-2.286 3-2.286 3 0V8.5m-6 0v-2c0-2.286 3-2.286 3 0c0 0 0 0 0 0v2m-6 5v-7A1.5 1.5 0 0 1 8.5 5v0c.828 0 1.5.555 1.5 1.384V8.5"/></g>`
	drawerInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 14H3m0-6h18m-10 9h2m-2-6h2m-2-6h2m8-2.4v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V2.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6ZM17.5 20v2m-11-2v2"/>`
	dribbbleInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Z"/><path d="M16.673 20.843C15.5 14 12.5 8 8.5 2.63"/><path d="M2.067 10.84C6 11 15.283 10.5 19.142 5m2.826 7.81C15.344 10.84 7.5 14 5.23 19.361"/></g>`
	droneInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m4.5 11.5l-5-4"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/></g>`
	droneChargeFullInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="m14.25 14.75l.25-2.25l.426-3.834A.6.6 0 0 0 14.33 8H9.67a.6.6 0 0 0-.596.666l.867 7.8a.6.6 0 0 0 .596.534H11"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 19v2m-8-2v2m2-2v2m2-2v2"/><path d="M13 22.4v-4.8a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-6.8a.6.6 0 0 1-.6-.6Z"/></g>`
	droneChargeHalfInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="m14.25 14.75l.25-2.25l.426-3.834A.6.6 0 0 0 14.33 8H9.67a.6.6 0 0 0-.596.666l.867 7.8a.6.6 0 0 0 .596.534H11"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 19v2m-8-2v2m2-2v2"/><path d="M13 22.4v-4.8a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-6.8a.6.6 0 0 1-.6-.6Z"/></g>`
	droneChargeLowInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="m14.25 14.75l.25-2.25l.426-3.834A.6.6 0 0 0 14.33 8H9.67a.6.6 0 0 0-.596.666l.867 7.8a.6.6 0 0 0 .596.534H11"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 19v2m-8-2v2"/><path d="M13 22.4v-4.8a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-6.8a.6.6 0 0 1-.6-.6Z"/></g>`
	droneCheckInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="m16 20l2 2l4-4"/></g>`
	droneErrorInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="m18 22.243l2.121-2.122m0 0L22.243 18m-2.122 2.121L18 18m2.121 2.121l2.122 2.122"/></g>`
	droneLandingInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M19.5 16v6m0 0L17 19.5m2.5 2.5l2.5-2.5"/></g>`
	droneRefreshInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m14.5 12.5l.426-3.834A.6.6 0 0 0 14.33 8H9.67a.6.6 0 0 0-.596.666l.867 7.8a.6.6 0 0 0 .596.534H11"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M21.666 16.667C21.048 15.097 19.634 14 17.99 14c-1.758 0-3.252 1.255-3.793 3"/><path stroke-linejoin="round" d="M19.995 16.772H21.4a.6.6 0 0 0 .6-.6V14.55m-7.666 4.783C14.952 20.903 16.366 22 18.01 22c1.758 0 3.252-1.255 3.793-3"/><path stroke-linejoin="round" d="M16.005 19.228H14.6a.6.6 0 0 0-.6.6v1.622"/></g>`
	droneTakeOffInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke-linejoin="round" d="M19.5 22v-6m0 0L17 18.5m2.5-2.5l2.5 2.5"/></g>`
	dropletInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M19 13.8C19 9.824 12 3 12 3S5 9.824 5 13.8c0 3.976 3.134 7.2 7 7.2s7-3.224 7-7.2Z"/>`
	dropletHalfInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M5.333 16A7.384 7.384 0 0 1 5 13.8C5 9.824 12 3 12 3s3.748 3.653 5.759 7.171M5.333 16c.904 2.9 3.547 5 6.667 5c3.866 0 7-3.224 7-7.2c0-1.067-.504-2.339-1.241-3.629M5.333 16l12.426-5.829"/>`
	easeCurveControlPointsInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm0 0h-2M7 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 0h2m5 0h-2m0 16h-2m-7 0c8 0 10-16 18-16"/>`
	easeInInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20c8 0 18-16 18-16"/>`
	easeInControlPointInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20c8 0 18-16 18-16m-4 16a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm0 0h-2m-3 0h-2"/>`
	easeInOutInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20c8 0 10-16 18-16"/>`
	easeOutInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20S13 4 21 4"/>`
	easeOutControlPointInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20S13 4 21 4M7 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 0h2m5 0h-2"/>`
	ecologyBookInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M10 14s.9-3.118 3-5"/><path stroke-linejoin="round" d="m12.802 12.425l-.134.012a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.498Z"/><path d="M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/></g>`
	editInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21h18M12.222 5.828L15.05 3L20 7.95l-2.828 2.828m-4.95-4.95l-5.607 5.607a1 1 0 0 0-.293.707v4.536h4.536a1 1 0 0 0 .707-.293l5.607-5.607m-4.95-4.95l4.95 4.95"/>`
	editPencilInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.02 5.828L15.85 3l4.949 4.95l-2.829 2.828m-4.95-4.95l-9.606 9.607a1 1 0 0 0-.293.707v4.536h4.536a1 1 0 0 0 .707-.293l9.606-9.607m-4.95-4.95l4.95 4.95"/>`
	eggInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22a8 8 0 0 0 8-8c0-4.418-3.582-12-8-12S4 9.582 4 14a8 8 0 0 0 8 8Z"/>`
	ejectInnerSVG                             = `<path fill="currentColor" d="m5 14l-.592-.46A.75.75 0 0 0 5 14.75V14Zm14 0v.75a.75.75 0 0 0 .592-1.21L19 14Zm-14 .75h14v-1.5H5v1.5Zm5.619-9.196L4.408 13.54l1.184.92l6.21-7.985l-1.183-.92Zm8.973 7.986l-6.21-7.986l-1.185.921l6.211 7.986l1.184-.921Zm-7.79-7.065a.25.25 0 0 1 .395 0l1.184-.92a1.749 1.749 0 0 0-2.762 0l1.184.92ZM5 17.25a.75.75 0 0 0 0 1.5v-1.5Zm14 1.5a.75.75 0 0 0 0-1.5v1.5Zm-14 0h14v-1.5H5v1.5Z"/>`
	electronicsChipInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 19.4V4.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6Zm7 .6v2.5M10 20v2.5M14 4V1.5M10 4V1.5M7 12H4.5m15 0H17M7 6.5H4.5m15 0H17m-10 11H4.5m15 0H17"/>`
	electronicsTransistorInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16V3.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6V16M7 16h2m-2 0H5m12 0h-2m2 0h2m-7 0v6m0-6H9m3 0h3m-6 0v6m6-6v6"/>`
	emojiInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Z"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2"/><path fill="currentColor" d="M15.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm-7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiBallInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9H8m8 0h-2m4 6H6m-4-3c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Z"/>`
	emojiBlinkLeftInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 9H8m-6 3c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Z"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2"/><path fill="currentColor" d="M15.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiBlinkRightInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path d="M14 9h2m6 3c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z"/><path d="M7.5 14.5s1.5 2 4.5 2s4.5-2 4.5-2"/></g>`
	emojiLookDownInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path d="M10 18h4m8-6c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z"/></g>`
	emojiLookLeftInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path d="M2.458 15A9.996 9.996 0 0 1 2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10c-4.478 0-8.268-2.943-9.542-7Zm0 0H7"/></g>`
	emojiLookRightInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M15.5 9a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/><path d="M21.542 15A9.997 9.997 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10c4.478 0 8.268-2.943 9.542-7Zm0 0H17"/></g>`
	emojiLookUpInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 7a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path d="M10 11h4m8 1c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z"/></g>`
	emojiPuzzledInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2"/><path d="M11.5 15.5s1.5-2 4.5-2s4.5 2 4.5 2M3 4c0-2.754 4-2.754 4 0c0 1.967-2 1.64-2 4m0 3.01l.01-.011"/><path fill="currentColor" d="M17.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm-7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiQuiteInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15h6m7-3c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiReallyInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9H8m8 0h-2m1 6H9m-7-3c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Z"/>`
	emojiSadInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M7.5 15.5s1.5-2 4.5-2s4.5 2 4.5 2"/></g>`
	emojiSatisfiedInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 9H8m8 0h-2M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Z"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2"/></g>`
	emojiSingLeftInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiSingLeftNoteInnerSVG                 = `<g fill="none" stroke-width="1.5"><path fill="currentColor" d="M2.8 8.1a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Z"/><path stroke="currentColor" stroke-linecap="round" d="M2.8 8.1a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm0 0V3.6a.6.6 0 0 1 .6-.6H5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.05 13c.501 5.053 4.765 9 9.95 9c5.523 0 10-4.477 10-10S17.523 2 12 2a9.966 9.966 0 0 0-4 .832"/><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 9a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-7 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/></g>`
	emojiSingRightInnerSVG                    = `<defs><path id="iconoirEmojiSingRight0" fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></defs><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><use href="#iconoirEmojiSingRight0"/><use href="#iconoirEmojiSingRight0"/></g>`
	emojiSingRightNoteInnerSVG                = `<g fill="none" stroke-width="1.5"><path fill="currentColor" d="M20.8 8.1a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Z"/><path stroke="currentColor" stroke-linecap="round" d="M20.8 8.1a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm0 0V3.6a.6.6 0 0 1 .6-.6H23"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.95 13c-.501 5.053-4.765 9-9.95 9c-5.523 0-10-4.477-10-10S6.477 2 12 2a9.97 9.97 0 0 1 4 .832"/><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiSurpriseInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm3.5 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3.5-8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiSurpriseAltInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiTalkingAngryInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9H8m8 0h-2M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Zm12 6h-4v-3c0-.667.4-2 2-2s2 1.333 2 2v3Z"/>`
	emojiTalkingHappyInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9H8m8 0h-2M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Zm12 1h-4v3c0 .667.4 2 2 2s2-1.333 2-2v-3Z"/>`
	emojiThinkLeftInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 15H7m-5-3c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Z"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emojiThinkRightInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 15h3m5-3c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	emptyPageInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	energyUsageWindowInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 19V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M11.667 11L10 14h4l-1.667 3"/></g>`
	enlargeInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 9l5-5m0 0v4m0-4h-4M9 15l-5 5m0 0v-4m0 4h4"/>`
	enlargeRoundArrowInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.5 9.5L6 12l2.5 2.5m7-5L18 12l-2.5 2.5"/><path d="M2 15V9a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4Z"/></g>`
	eraseInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21H9m6.889-6.11L8.464 7.463m-5.571 5.144l9.193-9.193a2 2 0 0 1 2.828 0l4.95 4.95a2 2 0 0 1 0 2.828l-9.243 9.243a1.929 1.929 0 0 1-2.728 0l-5-5a2 2 0 0 1 0-2.828Z"/>`
	errorWindowInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M15 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M18 22.243l2.121-2.122m0 0L22.243 18m-2.122 2.121L18 18m2.121 2.121l2.122 2.122"/></g>`
	euroInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.5 4.804a8 8 0 1 0 0 14.392M5 10h11M5 14h11"/>`
	euroSquareInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 7.503A4.746 4.746 0 0 0 13.87 7C11.18 7 9 9.239 9 12s2.18 5 4.87 5a4.73 4.73 0 0 0 2.13-.503M8 11h6m-6 2h6"/></g>`
	evChargeInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M6 9v10a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-5M9 5.6V7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6ZM4 5V3m4 2V3"/><path stroke-linejoin="round" d="M18.167 4L16.5 7h4l-1.667 3"/></g>`
	evChargeAltInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m22 5l-2 4l-2-4m-2 0h-2v4h2m-2-2h1.5"/><path d="M6 9v10a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-5M9 5.6V7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6ZM4 5V3m4 2V3"/></g>`
	evPlugInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 13.154V21m5-12.615v2.769a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-2.77a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2Zm-1.667-2V3M8.667 6.385V3"/>`
	evPlugChargingInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M10 13.154V21m5-12.615v2.769a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-2.77a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2Zm-1.667-2V3M6.667 6.385V3"/><path stroke-linejoin="round" d="M16.667 16L15 19h4l-1.667 3"/></g>`
	evPlugErrorInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M10 13.154V21m5-12.615v2.769a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-2.77a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2Zm-1.667-2V3M6.667 6.385V3"/><path stroke-linejoin="round" d="m15.121 21.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L15.12 17.12m2.122 2.122l2.121 2.121"/></g>`
	evStationInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 5v4"/><path d="M5 19V9a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2Z"/><path d="M5 10V5a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v5"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.167 11L9.5 14h4l-1.667 3"/></g>`
	evTagInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m18.5 9l-3 6l-3-6M10 9H6v6h4m-4-3h3"/><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/></g>`
	excludeInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 15h-.9a.6.6 0 0 0-.6.6v4.8a.6.6 0 0 0 .6.6h10.8a.6.6 0 0 0 .6-.6V9.6a.6.6 0 0 0-.6-.6h-4.8a.6.6 0 0 0-.6.6v.9"/><path d="M13.5 15h.9a.6.6 0 0 0 .6-.6v-.9m-6 0v.9a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-.9"/><path d="M9 10.5v-.9a.6.6 0 0 1 .6-.6h.9"/></g>`
	expandInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 9L4 4m0 0v4m0-4h4m7 5l5-5m0 0v4m0-4h-4M9 15l-5 5m0 0v-4m0 4h4m7-5l5 5m0 0v-4m0 4h-4"/>`
	expandLinesInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2H6m12 20H6m6-8v5m0 0l3-3m-3 3l-3-3m3-6V5m0 0l3 3m-3-3L9 8"/>`
	eyeAltInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 12.5c3-6.5 12-6.5 15 0"/><path d="M12 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g>`
	eyeCloseInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 8c3 6.5 12 6.5 15 0m-2.684 3.318L19.5 15M12 12.875V16.5m-4.816-5.182L4.5 15"/>`
	eyeEmptyInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M21 12c-1.889 2.991-5.282 6-9 6s-7.111-3.009-9-6c2.299-2.842 4.992-6 9-6s6.701 3.158 9 6Z"/></g>`
	eyeOffInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 3l18 18M10.5 10.677a2 2 0 0 0 2.823 2.823"/><path d="M7.362 7.561C5.68 8.74 4.279 10.42 3 12c1.889 2.991 5.282 6 9 6c1.55 0 3.043-.523 4.395-1.35M12 6c4.008 0 6.701 3.158 9 6a15.66 15.66 0 0 1-1.078 1.5"/></g>`
	faceIdInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 3H5a2 2 0 0 0-2 2v2m14-4h2a2 2 0 0 1 2 2v2m-5 1v2M8 8v2m1 6s1 1 3 1s3-1 3-1m-3-8v5h-1m-4 8H5a2 2 0 0 1-2-2v-2m14 4h2a2 2 0 0 0 2-2v-2"/>`
	facebookInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 2h-3a5 5 0 0 0-5 5v3H6v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3V2Z"/>`
	facebookTagInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5Z"/><path d="M11 21v-9c0-2.188.5-4 4-4m-6 5h6"/></g>`
	facetimeInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 16V8a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v8a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path d="M6 13v-2a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17.04 9.22l-3.067 2.3a.6.6 0 0 0 0 .96l3.067 2.3a.6.6 0 0 0 .96-.48V9.7a.6.6 0 0 0-.96-.48Z"/></g>`
	farmInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 20H2V8l6-3l6 3v12h-3m-6 0v-7h6v7m-6 0h6m7-6v6m-4-3h8m-8-3h8"/>`
	fastArrowDownInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 13l6 6l6-6M6 5l6 6l6-6"/>`
	fastArrowDownBoxInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.5 7.5L12 11L8.5 7.5m7 6L12 17l-3.5-3.5"/></g>`
	fastArrowLeftInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11 6l-6 6l6 6m8-12l-6 6l6 6"/>`
	fastArrowLeftBoxInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.5 8.5L13 12l3.5 3.5m-6-7L7 12l3.5 3.5"/></g>`
	fastArrowRightInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13 6l6 6l-6 6M5 6l6 6l-6 6"/>`
	fastArrowRightBoxInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m8 8.5l3.5 3.5L8 15.5m6-7l3.5 3.5l-3.5 3.5"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	fastArrowUpInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 11l6-6l6 6M6 19l6-6l6 6"/>`
	fastArrowUpBoxInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15.5 16.5L12 13l-3.5 3.5m7-6L12 7l-3.5 3.5"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	fastDownCircleInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 7.5L12 11l3.5-3.5m-7 6L12 17l3.5-3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	fastLeftCircleInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M16.5 8.5L13 12l3.5 3.5m-6-7L7 12l3.5 3.5"/></g>`
	fastRightCircleInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 8.5l3.5 3.5L8 15.5m6-7l3.5 3.5l-3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	fastUpCircleInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 16.5L12 13l3.5 3.5m-7-6L12 7l3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	favouriteBookInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M16 8.78a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.822-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.232 0L12 7.28l.106-.107A2.276 2.276 0 0 1 16 8.78Z"/><path stroke-linecap="round" d="M6 17h14M6 21h14"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/></g>`
	favouriteWindowInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M13 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path stroke-linejoin="round" d="M22 17.28a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.823-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.233 0l.106.107l.106-.107A2.277 2.277 0 0 1 22 17.28Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/></g>`
	femaleInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0 0v4m0 2v-2m0 0h-2m2 0h2"/>`
	figmaInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6a3 3 0 0 1 3-3h3v6H9a3 3 0 0 1-3-3Zm6-3h3a3 3 0 0 1 0 6h-3V3Z"/><path d="M12 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0v0Zm-6 6a3 3 0 0 1 3-3h3v3a3 3 0 0 1-6 0Zm0-6a3 3 0 0 1 3-3h3v6H9a3 3 0 0 1-3-3Z"/></g>`
	fileNotFoundInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v4m0 4.01l.01-.011M9 3H4v3m0 5v2m16-2v2M15 3h5v3M9 21H4v-3m11 3h5v-3"/>`
	filterInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 3h16a1 1 0 0 1 1 1v1.586a1 1 0 0 1-.293.707l-6.415 6.414a1 1 0 0 0-.292.707v6.305a1 1 0 0 1-1.243.97l-2-.5a1 1 0 0 1-.757-.97v-5.805a1 1 0 0 0-.293-.707L3.292 6.293A1 1 0 0 1 3 5.586V4a1 1 0 0 1 1-1Z"/>`
	filterAltInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3M3 7l6.65 5.7a1 1 0 0 1 .35.76v6.26a1 1 0 0 0 1.242.97l2-.5a1 1 0 0 0 .758-.97v-5.76a1 1 0 0 1 .35-.76L21 7M3 7h18"/>`
	finderInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 15V9a6 6 0 0 1 6-6h6a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6Z"/><path d="M15 3s-4.5 0-4.5 9H13c0 9 2 9 2 9"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2M7 9v2m10-2v2"/></g>`
	fingerprintInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 3.516A9.004 9.004 0 0 1 20.648 8.5M21 22v-8M3 22V11c0-1.052.18-2.062.512-3"/><path d="M18 22V11.3C18 7.82 15.314 5 12 5s-6 2.82-6 6.3V14m0 8v-4"/><path d="M9 22V11.15C9 9.41 10.343 8 12 8c.865 0 1.645.385 2.193 1M15 22v-8m-3 8v-3.5m0-7.5v3"/></g>`
	fingerprintCheckCircleInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10a9.98 9.98 0 0 1-.458 3M15.5 20.5l2 2l5-5"/></g>`
	fingerprintCircleInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	fingerprintErrorCircleInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 14v-1.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10a9.98 9.98 0 0 1-.458 3m-4.421 7.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L17.12 18.12m2.122 2.122l2.121 2.121"/></g>`
	fingerprintLockCircleInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 13.5v-.685m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M14 21.8c-.646.131-1.315.2-2 .2c-5.523 0-10-4.477-10-10S6.477 2 12 2s10 4.477 10 10c0 .254-.01.506-.028.755"/><path d="M21.167 18.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/></g>`
	fingerprintPhoneInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11V6.362c0-.51.1-1 .284-1.454M22 11V7.815m-7.778-5.08A5.507 5.507 0 0 1 17 2c2.28 0 4.203 1.33 4.805 3.15M15 12V9.824M19 12V6.853C19 5.83 18.105 5 17 5s-2 .83-2 1.853v.794M8 17.01l.01-.011M8 5H3.6a.6.6 0 0 0-.6.6v14.8a.6.6 0 0 0 .6.6h8.8a.6.6 0 0 0 .6-.6V16"/>`
	fingerprintScanInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794M6 3H3v3m15-3h3v3M6 21H3v-3m15 3h3v-3"/>`
	fingerprintSquareInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/></g>`
	fingerprintWindowInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M9 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v6"/><path stroke-linejoin="round" d="M12 21v-4.639c0-.51.1-.999.284-1.453M22 21v-3.185m-7.778-5.08A5.506 5.506 0 0 1 17 12c2.28 0 4.203 1.33 4.805 3.15M15 22v-2.177M19 22v-5.147C19 15.83 18.105 15 17 15s-2 .83-2 1.853v.794M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/></g>`
	fireFlameInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 18c0 2.415 1.79 3 4 3c3.759 0 5-2.5 2.5-7.5C11 18 10.5 11 11 9c-1.5 3-3 5.818-3 9Z"/><path d="M12 21c5.05 0 8-2.904 8-7.875C20 8.155 12 3 12 3S4 8.154 4 13.125C4 18.095 6.95 21 12 21Z"/></g>`
	fishingInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 0v10c0 6-10 6-10 0v-4l2 2"/>`
	flareInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.456 2.665a.6.6 0 0 1 1.088 0l2.864 6.137a.6.6 0 0 0 .29.29l6.137 2.864a.6.6 0 0 1 0 1.088l-6.137 2.864a.6.6 0 0 0-.29.29l-2.864 6.137a.6.6 0 0 1-1.088 0l-2.864-6.137a.6.6 0 0 0-.29-.29l-6.137-2.864a.6.6 0 0 1 0-1.088l6.137-2.864a.6.6 0 0 0 .29-.29l2.864-6.137Z"/>`
	flashInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 10V3L5 14h6v7l8-11h-6Z"/>`
	flashOffInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.795 8.782L5 14h6v7l4-5.5m2.182-3L19 10h-6V3l-2.182 3M4 4l16 16"/>`
	flaskInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M18.5 15h-13"/><path stroke-linecap="round" d="M16 4H8m1 .5v5.76a2 2 0 0 1-.481 1.302L3.48 17.438A2 2 0 0 0 3 18.74V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-.26a2 2 0 0 0-.482-1.302l-5.037-5.876A2 2 0 0 1 15 10.26V4.5m-3 4.51l.01-.011M11 2.01l.01-.011"/></g>`
	flipInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 20H2L9.5 4v16Zm10.625 0H22l-.938-2m-4.687 2H14.5v-2m0-6v2m3.75-2l.938 2m-2.813-6L14.5 4v4"/>`
	flipReverseInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 12v2m-3.75-2l-.938 2m2.813 6H9.5v-2m-5.625 2H2l.938-2M7.625 8L9.5 4v4m5 12H22L14.5 4v16Z"/>`
	flowerInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm1-6s1-2 1-4s-2-4-2-4s-2 2-2 4s1 4 1 4"/><path d="M9 11s-2-1-4-1s-4 2-4 2s2 2 4 2s4-1 4-1m4 2s1 2 1 4s-2 4-2 4s-2-2-2-4s1-4 1-4m4-4s2-1 4-1s4 2 4 2s-2 2-4 2s-4-1-4-1m-4.414-3.828S9.879 7.05 8.464 5.636C7.05 4.222 4.222 4.222 4.222 4.222s0 2.828 1.414 4.243c1.414 1.414 3.536 2.121 3.536 2.121m0 2.828s-2.122.707-3.536 2.122c-1.414 1.414-1.414 4.242-1.414 4.242s2.828 0 4.242-1.414c1.415-1.414 2.122-3.536 2.122-3.536m4.243-1.414s2.12.707 3.535 2.122c1.414 1.414 1.414 4.242 1.414 4.242s-2.828 0-4.242-1.414c-1.415-1.414-2.122-3.536-2.122-3.536m0-5.656s.707-2.122 2.122-3.536c1.414-1.414 4.242-1.414 4.242-1.414s0 2.828-1.414 4.243c-1.414 1.414-3.536 2.121-3.536 2.121"/></g>`
	fluorineInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M10 16V8h4m-4 4h4"/></g>`
	fogInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 14h6m-6 8h6m-8-4h10m-13.5-.618C2.188 16.707 1 15.388 1 13c0-4 3.333-5 5-5c0-2 0-6 6-6s6 4 6 6c1.667 0 5 1 5 5c0 2.388-1.188 3.707-2.5 4.382"/>`
	folderInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 11V4.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6V11M2 11v8.4a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V11M2 11h20"/>`
	folderAlertInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 3v4m0 4.01l.01-.011M22 7v12.4a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V11"/><path d="M14 7h-1.278a.6.6 0 0 1-.39-.144L9.168 4.144A.6.6 0 0 0 8.778 4H2.6a.6.6 0 0 0-.6.6V11h12"/></g>`
	folderSettingsInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.6 4h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6ZM22 10v4"/><path d="M2 10v9.4a.6.6 0 0 0 .6.6H13m6 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-dasharray=".3 2" d="M19 22a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	fontSizeInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 21V11m0 10l-2-2.5m2 2.5l2-2.5M18 11l-2 2m2-2l2 2M9 5v12m0 0H7m2 0h2m4-10V5H3v2"/>`
	footballInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 13.828V21M5 3v10.828h14V3"/><path d="M13 6.732c1.071-.618 1.434-2.114 1.549-2.833a.505.505 0 0 0-.321-.556c-.68-.26-2.157-.693-3.228-.075C9.93 3.886 9.567 5.38 9.452 6.1a.505.505 0 0 0 .32.556c.681.26 2.158.693 3.228.075Z"/></g>`
	footballBallInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.115 14.015a22.314 22.314 0 0 0-.103 3.665a2.413 2.413 0 0 0 2.309 2.308c1.007.052 2.294.055 3.664-.103m-5.87-5.87C4.394 11.604 5.17 8.93 7.05 7.05c1.88-1.88 4.554-2.656 6.965-2.935m-9.9 9.9l5.87 5.87m0 0c2.411-.279 5.084-1.055 6.965-2.935c1.88-1.88 2.656-4.554 2.935-6.965m-5.87-5.87a22.314 22.314 0 0 1 3.665-.103a2.413 2.413 0 0 1 2.308 2.309a22.312 22.312 0 0 1-.103 3.664m-5.87-5.87l5.87 5.87M9.172 14.828l1.414-1.414m0 0L9.172 12m1.414 1.414L12 14.828m-1.414-1.414l2.828-2.828m0 0l1.414-1.414m-1.414 1.414L12 9.172m1.414 1.414L14.828 12"/>`
	forwardInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.956 5.704A.6.6 0 0 0 2 6.187v11.626a.6.6 0 0 0 .956.483l7.889-5.813a.6.6 0 0 0 0-.966l-7.89-5.813Zm11 0a.6.6 0 0 0-.956.483v11.626a.6.6 0 0 0 .956.483l7.889-5.813a.6.6 0 0 0 0-.966l-7.89-5.813Z"/>`
	forwardFifteenSecondsInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 13a9 9 0 1 1-9-9m0 0h7.5m0 0l-2-2m2 2l-2 2M9 9v7"/><path d="M15 9h-2a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h1a1 1 0 0 1 1 1V15a1 1 0 0 1-1 1h-2"/></g>`
	forwardMessageInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 8l5 3l5-3"/><path d="M10 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v6.857"/><path stroke-linejoin="round" d="M22 17.111h-6.3c-3.6 0-3.6 4.889 0 4.889m6.3-4.889L18.85 14M22 17.111l-3.15 3.111"/></g>`
	fourKDisplayInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M13.5 9v4m0 2v-2m0 0l1.37-1.566M17 9l-2.13 2.434m0 0L17 15M9.5 9l-3 4.5H10V15"/><path d="M2 18.4V5.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`
	frameInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M4.998 2.001H2v2.998h2.998V2.001Zm0 8.501H2V13.5h2.998v-2.998ZM20.498 5v5.503M3.5 5v5.503m16.998 2.999v5.502M3.5 13.502v5.502m1.499 1.498h5.5"/><path stroke-width="1.22" d="M4.999 3.503h5.5"/><path d="M13.498 20.499h5.5"/><path stroke-width="1.22" d="M13.498 3.501h5.5"/><path d="M4.998 19.001H2v2.998h2.998v-2.998ZM21.997 2.002H19V5h2.998V2.002ZM13.497 2H10.5v2.998h2.998V2Zm8.5 8.503H19V13.5h2.998v-2.998Zm0 8.499H19V22h2.998v-2.998Zm-8.5-.002H10.5v2.998h2.998V19Z"/></g>`
	frameAltInnerSVG                          = `<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 3v18M18 3v18M3 6h18"/><path fill="currentColor" fill-rule="evenodd" d="M9.6 9h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6V9.6a.6.6 0 0 1 .6-.6Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 18h18"/></g>`
	frameAltEmptyInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3v18M18 3v18M3 6h18M3 18h18"/>`
	frameSelectInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M4.998 2H2v2.998h2.998V2Zm.001 1.501h14M3.5 4.999V19M20.498 5v14.002M4.999 20.501h14M4.998 19H2v2.998h2.998V19ZM21.997 2.002H19V5h2.998V2.002Zm0 17H19V22h2.998v-2.998Z"/><path d="m10.997 15.002l-3-7l7 3l-2.998.999l-1.002 3.001Z" clip-rule="evenodd"/><path d="m11.999 12.002l2.998 3l-2.998-3Z" clip-rule="evenodd"/></g>`
	frameSimpleInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5" d="M4.998 2H2v2.998h2.998V2Zm0 1.501h14M3.499 4.998V19M20.497 5v14.002M4.998 20.501h14M4.998 19H2v2.998h2.998V19ZM21.996 2.002h-2.998V5h2.998V2.002Zm0 17h-2.998V22h2.998v-2.998Z"/>`
	frameToolInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 7h1M2 17h1M21 7h1m-1 10h1M17 3V2M7 3V2m10 20v-1M7 22v-1M18 6.6v10.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6V6.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6Z"/>`
	fridgeInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 14H9m1-8H9"/><path d="M5 10V2.6a.6.6 0 0 1 .6-.6h12.8a.6.6 0 0 1 .6.6V10M5 10v11.4a.6.6 0 0 0 .6.6h12.8a.6.6 0 0 0 .6-.6V10M5 10h14"/></g>`
	fxInnerSVG                                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 17V7h7m-7 5h5m5 5l4-5m0 0l4-5m-4 5l-4-5m4 5l4 5"/>`
	fxTagInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 15V9h5m-5 3h3.571M13 15l2.5-3m0 0L18 9m-2.5 3L13 9m2.5 3l2.5 3"/></g>`
	gamepadInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 17.5c2.5 3.5 6.449.915 5.5-2.5c-1.425-5.129-2.2-7.984-2.603-9.492A2.032 2.032 0 0 0 18.438 4H5.562c-.918 0-1.718.625-1.941 1.515C2.78 8.863 2.033 11.802 1.144 15c-.948 3.415 3 6 5.5 2.5M18 8.5l.011.01M16.49 7l.011.01M16.49 10l.011.01M15 8.5l.011.01M7 7v3M5.5 8.5h3"/><path d="M8 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm8 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	garageInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20H3V6l9-2l9 2v14h-3M6 20h12M6 20v-4m12 4v-4M6 12V8h12v4M6 12h12M6 12v4m12-4v4M6 16h12"/>`
	gasInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 8a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v13.4a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6V8Zm0 3h6m-3-6V2m0 0h-1m1 0h1"/>`
	gasTankInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path stroke-width="1.493" d="M3 7.562A2.562 2.562 0 0 1 5.563 5H7V3h5v2h2.002A6.998 6.998 0 0 1 21 11.998v6.442a2.562 2.562 0 0 1-2.563 2.562H5.563A2.565 2.565 0 0 1 3 18.44V7.562Z" clip-rule="evenodd"/><path stroke-width="1.502" d="m8 8.878l8 8.238l-4-4.121l-4 4.12l4-4.12l4-4.121"/></g>`
	gasTankDropInnerSVG                       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5" clip-rule="evenodd"><path stroke-width="1.493" d="M3 7.562A2.562 2.562 0 0 1 5.563 5H7V3h5v2h2.002A6.998 6.998 0 0 1 21 11.998v6.442a2.562 2.562 0 0 1-2.563 2.562H5.563A2.565 2.565 0 0 1 3 18.44V7.562Z"/><path d="M12 9s3 2.993 3 4.886c0 1.656-1.345 3-3 3c-1.656 0-2.988-1.344-3-3C9.01 11.992 12 9 12 9Z"/></g>`
	gifFormatInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M15.5 15V9h3m-3 3h2M12 15V9M8.5 9h-3v6h3v-2.4"/></g>`
	giftInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 12v9.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V12m17.4-5H2.6a.6.6 0 0 0-.6.6v3.8a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6ZM12 22V7m0 0H7.5a2.5 2.5 0 1 1 0-5C11 2 12 7 12 7Zm0 0h4.5a2.5 2.5 0 0 0 0-5C13 2 12 7 12 7Z"/>`
	gitBranchInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-4V3"/><path d="M8 18h1c3.5 0 9-2.1 9-8.5V8"/></g>`
	gitCherryPickCommitInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 11v-1a2 2 0 0 0-2-2h-3m-5 3v-1a2 2 0 0 1 2-2h3m0 0V4m0 16a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm-3-3H3m12 0h6"/>`
	gitCommandInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6ZM10 16l4-8"/>`
	gitCommitInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm-3-3H3m12 0h6"/>`
	gitCompareInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm12 10V7s0-2-2-2h-3M6 7v10s0 2 2 2h3"/><path d="M15 7.5L12.5 5L15 2.5m-6.5 14L11 19l-2.5 2.5"/></g>`
	gitForkInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM7 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM7 7v10M17 7v1c0 2.5-2 3-2 3l-6 2s-2 .5-2 3v1"/>`
	gitMergeInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 21V7"/><path d="M6 7v2c0 3.5 2.5 9 8.5 9H16M6 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	gitPullRequestInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 7v10m12 0V7s0-2-2-2h-3"/><path d="M15 7.5L12.5 5L15 2.5"/></g>`
	gitPullRequestClosedInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-10v6m12 0V7s0-2-2-2h-4M4 7.243L6.121 5.12m0 0L8.243 3M6.12 5.121L4 3m2.121 2.121l2.122 2.122"/>`
	githubInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 22.027v-2.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7a5.44 5.44 0 0 0-1.5-3.75a5.07 5.07 0 0 0-.09-3.77s-1.18-.35-3.91 1.48a13.38 13.38 0 0 0-7 0c-2.73-1.83-3.91-1.48-3.91-1.48A5.07 5.07 0 0 0 5 5.797a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7a3.37 3.37 0 0 0-.94 2.58v2.87"/><path d="M9 20.027c-3 .973-5.5 0-7-3"/></g>`
	githubCircleInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M14.333 19v-1.863c.025-.31-.018-.62-.126-.913a2.18 2.18 0 0 0-.5-.781c2.093-.227 4.293-1 4.293-4.544a3.48 3.48 0 0 0-1-2.434a3.211 3.211 0 0 0-.06-2.448s-.787-.227-2.607.961a9.152 9.152 0 0 0-4.666 0c-1.82-1.188-2.607-.96-2.607-.96A3.211 3.211 0 0 0 7 8.464a3.482 3.482 0 0 0-1 2.453c0 3.519 2.2 4.291 4.293 4.544a2.18 2.18 0 0 0-.496.773a2.134 2.134 0 0 0-.13.902V19"/><path d="M9.667 17.702c-2 .631-3.667 0-4.667-1.948"/></g>`
	gitlabFullInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M17.057 2.544a.2.2 0 0 1 .378-.008l3.114 8.31l1.398 3.73a.2.2 0 0 1-.07.232l-9.76 7.106a.2.2 0 0 1-.235 0l-9.76-7.106a.2.2 0 0 1-.069-.231l1.398-3.73l.167-.45l2.944-7.861a.2.2 0 0 1 .378.008l2.47 7.6a.2.2 0 0 0 .19.137h4.8a.2.2 0 0 0 .19-.138l2.467-7.599Z" clip-rule="evenodd"/>`
	glassEmptyInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m7.5 11l1 5"/><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944l-1.7-12.744Z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`
	glassHalfInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944l-1.7-12.744Z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0M4 13c1.032 1.203 3.925 1.864 7 1.981a25.406 25.406 0 0 0 4-.158c2.266-.279 4.197-.886 5-1.823M4 13c2.286-2.667 13.714-2.667 16 0"/></g>`
	glassHalfAltInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944l-1.7-12.744Z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0M4 13c1.032 1.203 3.925 1.864 7 1.981c3.739.143 7.746-.518 9-1.981m-5 1.823V20.5M4 13c2.286-2.667 13.714-2.667 16 0"/></g>`
	glassesInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 14a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm0 0V6m20 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 0V6m-8 8h-4"/>`
	globeInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="m2.5 12.5l5.5 2L7 18l1 3m9-.5l-.5-2.5l-2.5-1v-3.5l3-1l4.5.5M19 5.5L18.5 7l-3.5.5v3l2.5-1h2l2 1m-19 0l2.5-2L7.5 8l2-3l-1-2"/></g>`
	golfInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 18v-6m0 0V3.41a.6.6 0 0 1 .836-.552l8.444 3.62a.6.6 0 0 1 .022 1.093L12 12Zm0 10c3.866 0 7-1.567 7-3.5S15.866 15 12 15s-7 1.567-7 3.5S8.134 22 12 22Z"/>`
	googleInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M15.547 8.303A5.148 5.148 0 0 0 12.11 7C9.287 7 7 9.239 7 12s2.287 5 5.109 5c3.47 0 4.751-2.57 4.891-4.583h-4.159"/><path stroke-linecap="round" stroke-linejoin="round" d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5Z"/></g>`
	googleCircleInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M15.547 8.303A5.148 5.148 0 0 0 12.11 7C9.287 7 7 9.239 7 12s2.287 5 5.109 5c3.47 0 4.751-2.57 4.891-4.583h-4.159"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	googleDocsInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2ZM7 7h10M7 12h10M7 17h6"/>`
	googleDriveInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.143 3.004L14.857 3m-5.714.004L2 15.004m7.143-12L18.433 21M14.856 3L22 15.004M14.857 3L5.575 21m12.857 0H5.575m12.857 0L22 15.004M5.575 21L2 15.004m20 0H2"/>`
	googleDriveCheckInnerSVG                  = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirGoogleDriveCheck0)"><path d="M9.143 3.004L14.857 3m-5.714.004L2 15.004m7.143-12l4.902 9.496m.812-9.5L5.575 21m9.282-18L21.5 14M5.575 21L2 15.004M5.575 21h6.429M2 15.004h10.5M15 19l3 3l5-5"/></g><defs><clipPath id="iconoirGoogleDriveCheck0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	googleDriveSyncInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.143 3.004L14.857 3m-5.714.004L2 15.004m7.143-12l4.902 9.496m.812-9.5L5.575 21m9.282-18l5.356 9M5.575 21L2 15.004M5.575 21h6.429M2 15.004h10.5m10.166 2.663C22.048 16.097 20.634 15 18.99 15c-1.758 0-3.252 1.255-3.793 3"/><path d="M20.995 17.772H22.4a.6.6 0 0 0 .6-.6V15.55m-7.666 4.783C15.952 21.903 17.366 23 19.01 23c1.758 0 3.252-1.255 3.793-3"/><path d="M17.005 20.228H15.6a.6.6 0 0 0-.6.6v1.622"/></g>`
	googleDriveWarningInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.143 3.004L14.857 3m-5.714.004L2 15.004m7.143-12l4.902 9.496m.812-9.5L5.575 21m9.282-18L21.5 14M5.575 21L2 15.004M5.575 21h6.429M2 15.004h10.5M18 16v2m0 4.01l.01-.011"/>`
	googleHomeInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M17.708 17A9 9 0 1 0 4.291 5a9 9 0 0 0 13.417 12Zm0 0H19.5a2.5 2.5 0 0 1 2.5 2.5v0a2.5 2.5 0 0 1-2.5 2.5H17"/><path stroke-linejoin="round" d="m11 11.01l.01-.011M8 11.01l.01-.011m5.99.011l.01-.011"/></g>`
	googleOneInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11 5v14a2 2 0 1 0 4 0V5a2 2 0 1 0-4 0Z"/><path d="M11.64 3.53L6.747 8.171a2 2 0 0 0 2.754 2.901l4.892-4.642a2 2 0 0 0-2.753-2.902Z"/></g>`
	gpsInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 8.5h-2.25A1.75 1.75 0 0 0 18 10.25v0c0 .966.784 1.75 1.75 1.75h1.5c.966 0 1.75.784 1.75 1.75v0a1.75 1.75 0 0 1-1.75 1.75H18m-7.5 0v-2.8m0 0h2.857c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H10.5v4.2Zm-4-3.573a3.5 3.5 0 1 0-2 6.373C6.433 15.5 8 14 8 12H5"/>`
	graduationCapInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2.573 8.463l8.659-4.329a.6.6 0 0 1 .536 0l8.659 4.33a.6.6 0 0 1 0 1.073l-8.659 4.329a.6.6 0 0 1-.536 0l-8.659-4.33a.6.6 0 0 1 0-1.073Z"/><path d="M22.5 13V9.5l-2-1m-16 2v5.412a2 2 0 0 0 1.142 1.807l5 2.374a2 2 0 0 0 1.716 0l5-2.374a2 2 0 0 0 1.142-1.807V10.5"/></g>`
	graphDownInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 20H4V4"/><path d="m4 7l8 8l3-3l4.5 4.5"/></g>`
	graphUpInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 20H4V4"/><path d="M4 16.5L12 9l3 3l4.5-4.5"/></g>`
	greenBusInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 16.01l.01-.011M21 12H3v7a1 1 0 0 0 1 1h9m4 3s.9-3.118 3-5"/><path stroke-linejoin="round" d="m19.802 21.424l-.134.013a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.497ZM21 8V6a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2m4 0h10"/><path d="M4.5 20v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20"/></g>`
	greenTruckInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M7 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M14 15V4.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6h2.05M14 15H9.05M14 7h5.61a.6.6 0 0 1 .548.356l1.29 2.903a.6.6 0 0 1 .052.243V12"/><path stroke-linejoin="round" d="M17 23s.9-3.118 3-5"/><path stroke-linejoin="round" d="m19.802 21.424l-.134.013a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.497Z"/></g>`
	greenVehicleInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M7 10h8m-9 4h1m8 0h1"/><path d="M6 18H2v2.4a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V18Zm0 0h7M2 18v-6.59a2 2 0 0 1 .162-.787l2.319-5.41A2 2 0 0 1 6.319 4h9.362a2 2 0 0 1 1.839 1.212l2.318 5.41a2 2 0 0 1 .162.789V12.5"/><path stroke-linejoin="round" d="M17 23s.9-3.118 3-5"/><path stroke-linejoin="round" d="m19.802 21.424l-.134.013a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.497Z"/></g>`
	gridAddInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M13.992 17h3m3 0h-3m0 0v-3m0 3v3"/><path d="M4 9.4V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm0 10v-4.8a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm10-10V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-4.8a.6.6 0 0 1-.6-.6Z"/></g>`
	gridMinusInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M13.992 17h6"/><path d="M4 9.4V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm0 10v-4.8a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm10-10V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-4.8a.6.6 0 0 1-.6-.6Z"/></g>`
	gridRemoveInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M14.871 19.121L16.993 17m2.121-2.121L16.993 17m0 0l-2.122-2.121M16.993 17l2.121 2.121"/><path d="M4 9.4V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm0 10v-4.8a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm10-10V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-4.8a.6.6 0 0 1-.6-.6Z"/></g>`
	groupInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M1 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1"/><path d="M13 14v0a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v.5"/><path stroke-linejoin="round" d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm10-3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	gymInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.4 7H4.6a.6.6 0 0 0-.6.6v8.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6Zm12 0h-2.8a.6.6 0 0 0-.6.6v8.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6Z"/><path d="M1 14.4V9.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H1.6a.6.6 0 0 1-.6-.6Zm22 0V9.6a.6.6 0 0 0-.6-.6h-1.8a.6.6 0 0 0-.6.6v4.8a.6.6 0 0 0 .6.6h1.8a.6.6 0 0 0 .6-.6ZM8 12h8"/></g>`
	halfCookieInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.8 14c-.927 4.564-4.962 8-9.8 8c-5.523 0-10-4.477-10-10c0-5.185 3.947-9.449 9-9.95"/><path d="M6.5 10a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm14-6a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1ZM12 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-5-3.99l.01-.011m9.99.011l.01-.011M11 12.01l.01-.011M21 9.01l.01-.011M17 6.01l.01-.011M11 2c-.5 1.5.5 3 2.085 3C11 8.5 13 12 18 11.5c0 2.5 2.5 3 3.7 2.514"/></g>`
	halfMoonInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 11.507a9.493 9.493 0 0 0 18 4.219c-8.507 0-12.726-4.22-12.726-12.726A9.494 9.494 0 0 0 3 11.507Z"/>`
	hammerInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.634 11.056L2.148 19.54l2.122 2.121l8.485-8.485"/><path d="m10.634 11.056l1.414-1.415s.354-3.182-3.182-6.717l1.06-1.06l8.486 5.656l-1.06 1.06l1.413 1.415l1.061-1.06l2.475 2.474l-4.95 4.95l-2.475-2.475l1.061-1.06l-1.414-1.415l-1.768 1.768l-2.121-2.121Z"/></g>`
	handBrakeInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 16v-4m0-3V8"/><circle cx="12" cy="12" r="8"/><path stroke-linecap="round" stroke-linejoin="round" d="M3.953 4.5A10.961 10.961 0 0 0 1 12c0 2.899 1.121 5.535 2.953 7.5m16.094-15A10.962 10.962 0 0 1 23 12c0 2.899-1.121 5.535-2.953 7.5"/></g>`
	handbagInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-9a2 2 0 0 0-2-2h-5M9 8V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6V8M9 8h6M9 8v6m6-6v6"/>`
	hardDriveInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m10 17.01l.01-.011M6 17.01l.01-.011"/><path d="M2 13v7.4a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V13M2 13h20M2 13l2.872-9.572A.6.6 0 0 1 5.446 3h13.108a.6.6 0 0 1 .574.428L22 13"/></g>`
	hatInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 17v-2a7 7 0 1 1 14 0v2H7Zm0 0H2M14 6.01l.01-.011"/>`
	hdInnerSVG                                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7v5m0 5v-5m0 0h7m0 0V7m0 5v5m3-5V7c4 0 8 0 8 5s-4 5-8 5v-5Z"/>`
	hdDisplayInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 8.5V12m0 3.5V12m0 0h4.5m0 0V8.5m0 3.5v3.5M14 12V8.5c2.5 0 5 0 5 3.5s-2.5 3.5-5 3.5V12Z"/><path d="M1 15V9a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6Z"/></g>`
	hdrInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.5 8.5V12m0 3.5V12m0 0H6m0 0V8.5M6 12v3.5m11.5 0v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H17.5v4.2m2.857 0H17.5m2.857 0l2.143 2.8M9.5 12V8.5c2.5 0 5 0 5 3.5s-2.5 3.5-5 3.5V12Z"/>`
	headsetInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 13.5V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5"/><path d="M2 17.439v-1.877a2 2 0 0 1 1.515-1.94L4 13.5l1.254-.314a.6.6 0 0 1 .746.582v5.464a.6.6 0 0 1-.746.582l-1.74-.435A2 2 0 0 1 2 17.439Zm20 0v-1.877a2 2 0 0 0-1.515-1.94L20 13.5l-1.255-.314a.6.6 0 0 0-.745.582v5.464a.6.6 0 0 0 .745.582l1.74-.435A2 2 0 0 0 22 17.439Z"/></g>`
	headsetChargeInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.5 13L10 17h4l-2.5 4"/><path d="M4 13.5V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5"/><path d="M2 17.439v-1.877a2 2 0 0 1 1.515-1.94L4 13.5l1.254-.314a.6.6 0 0 1 .746.582v5.464a.6.6 0 0 1-.746.582l-1.74-.435A2 2 0 0 1 2 17.439Zm20 0v-1.877a2 2 0 0 0-1.515-1.94L20 13.5l-1.255-.314a.6.6 0 0 0-.745.582v5.464a.6.6 0 0 0 .745.582l1.74-.435A2 2 0 0 0 22 17.439Z"/></g>`
	headsetHelpInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M20 11a8 8 0 1 0-16 0"/><path d="M2 15.438v-1.876a2 2 0 0 1 1.515-1.94l1.74-.436a.6.6 0 0 1 .745.582v5.463a.6.6 0 0 1-.746.583l-1.74-.435A2 2 0 0 1 2 15.439Zm20 0v-1.876a2 2 0 0 0-1.515-1.94l-1.74-.436a.6.6 0 0 0-.745.582v5.463a.6.6 0 0 0 .745.583l1.74-.435A2 2 0 0 0 22 15.439ZM20 18v.5a2 2 0 0 1-2 2h-3.5"/><path d="M13.5 22h-3a1.5 1.5 0 0 1 0-3h3a1.5 1.5 0 0 1 0 3Z"/></g>`
	headsetIssueInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12v5m0 4.01l.01-.011M4 13.5V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5"/><path d="M2 17.439v-1.877a2 2 0 0 1 1.515-1.94L4 13.5l1.254-.314a.6.6 0 0 1 .746.582v5.464a.6.6 0 0 1-.746.582l-1.74-.435A2 2 0 0 1 2 17.439Zm20 0v-1.877a2 2 0 0 0-1.515-1.94L20 13.5l-1.255-.314a.6.6 0 0 0-.745.582v5.464a.6.6 0 0 0 .745.582l1.74-.435A2 2 0 0 0 22 17.439Z"/></g>`
	healthShieldInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.667 16h-3.334v-2.333H8v-3.334h2.333V8h3.334v2.333H16v3.334h-2.333V16Z"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/></g>`
	healthcareInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m18 20l3.824-3.824a.6.6 0 0 0 .176-.424V10.5A1.5 1.5 0 0 0 20.5 9v0a1.5 1.5 0 0 0-1.5 1.5V15"/><path d="m18 16l.858-.858a.484.484 0 0 0 .142-.343v0a.485.485 0 0 0-.268-.433l-.443-.221a2 2 0 0 0-2.308.374l-.895.895a2 2 0 0 0-.586 1.414V20M6 20l-3.824-3.824A.6.6 0 0 1 2 15.752V10.5A1.5 1.5 0 0 1 3.5 9v0A1.5 1.5 0 0 1 5 10.5V15"/><path d="m6 16l-.858-.858A.485.485 0 0 1 5 14.799v0c0-.183.104-.35.268-.433l.443-.221a2 2 0 0 1 2.308.374l.895.895a2 2 0 0 1 .586 1.414V20m4.167-8h-3.334V9.667H8V6.333h2.333V4h3.334v2.333H16v3.334h-2.333V12Z"/></g>`
	heartInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M22 8.862a5.95 5.95 0 0 1-1.654 4.13c-2.441 2.531-4.809 5.17-7.34 7.608c-.581.55-1.502.53-2.057-.045l-7.295-7.562c-2.205-2.286-2.205-5.976 0-8.261a5.58 5.58 0 0 1 8.08 0l.266.274l.265-.274A5.612 5.612 0 0 1 16.305 3c1.52 0 2.973.624 4.04 1.732A5.95 5.95 0 0 1 22 8.862Z"/>`
	heatingInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M8 6s-2.5 3 0 6s0 6 0 6m4-12s-2.5 3 0 6s0 6 0 6m4-12s-2.5 3 0 6s0 6 0 6"/></g>`
	heavyRainInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 13v7m8-7v7m-4-5v7m8-4.393c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`
	helpCircleInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M9 9c0-3.5 5.5-3.5 5.5 0c0 2.5-2.5 2-2.5 5m0 4.01l.01-.011"/></g>`
	helpSquareInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M9 9c0-3.5 5.5-3.5 5.5 0c0 2.5-2.5 2-2.5 5m0 4.01l.01-.011"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	heptagonInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.74 1.625a.6.6 0 0 1 .52 0l8.08 3.891a.6.6 0 0 1 .324.407l1.996 8.743a.6.6 0 0 1-.116.508l-5.591 7.01a.6.6 0 0 1-.47.227H7.517a.6.6 0 0 1-.469-.226l-5.591-7.011a.6.6 0 0 1-.116-.508l1.996-8.743a.6.6 0 0 1 .324-.407l8.08-3.89Z"/>`
	herSlipsInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M1 4.6a.6.6 0 0 1 .6-.6h20.8a.6.6 0 0 1 .6.6v3.912c0 .284-.199.53-.476.595c-1.052.247-3.635.914-5.524 1.893c-3.444 1.786-3.93 6.655-3.993 8.382a.637.637 0 0 1-.627.618h-.76a.637.637 0 0 1-.627-.618C10.93 17.655 10.443 12.786 7 11c-1.889-.98-4.472-1.646-5.524-1.893A.614.614 0 0 1 1 8.512V4.6Z"/>`
	hexagonInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.7 1.173a.6.6 0 0 1 .6 0l8.926 5.154a.6.6 0 0 1 .3.52v10.307a.6.6 0 0 1-.3.52L12.3 22.826a.6.6 0 0 1-.6 0l-8.926-5.154a.6.6 0 0 1-.3-.52V6.847a.6.6 0 0 1 .3-.52L11.7 1.174Z"/>`
	hexagonAltInnerSVG                        = `<g fill="none" stroke-width="1.5"><g clip-path="url(#iconoirHexagonAlt0)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.327 2.774a.6.6 0 0 1 .52-.3h10.307a.6.6 0 0 1 .52.3l5.153 8.926a.6.6 0 0 1 0 .6l-5.154 8.926a.6.6 0 0 1-.52.3H6.847a.6.6 0 0 1-.52-.3L1.174 12.3a.6.6 0 0 1 0-.6l5.154-8.926Z"/></g><defs><clipPath id="iconoirHexagonAlt0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	hexagonDiceInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M11.7 1.173a.6.6 0 0 1 .6 0l8.926 5.154a.6.6 0 0 1 .3.52v10.307a.6.6 0 0 1-.3.52L12.3 22.826a.6.6 0 0 1-.6 0l-8.926-5.154a.6.6 0 0 1-.3-.52V6.847a.6.6 0 0 1 .3-.52L11.7 1.174Z"/><path stroke-linecap="round" d="M17 15H7l5-8l5 8Z"/><path d="M2.5 6.5L12 7m-9.5-.5L7 15m14.5-8.5L17 15m4.5-8.5L12 7V1m9.5 16.5L17 15M2.5 17.5L7 15m0 0l5 8l5-8"/></g>`
	highPriorityInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848L11.576 1.424ZM12 8v4m0 4.01l.01-.011"/>`
	historicShieldInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 15.528V2.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v12.928a4 4 0 0 1-2.211 3.578l-5.52 2.76a.6.6 0 0 1-.537 0l-5.52-2.76A4 4 0 0 1 4 15.528Z"/>`
	historicShieldAltInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.732 21.866l-5.52-2.76A4 4 0 0 1 4 15.528V2.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v12.928a4 4 0 0 1-2.211 3.578l-5.52 2.76a.6.6 0 0 1-.537 0ZM12 10V2m-8 8h16"/>`
	homeInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13"/>`
	homeAltInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12.4 17h-.8a.6.6 0 0 1-.6-.6v-1.8a.6.6 0 0 1 .6-.6h.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 1-.6.6Z"/><path d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13"/></g>`
	homeAltSlimInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13m7 4v-4"/>`
	homeAltSlimHorizInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13m5 3h4"/>`
	homeHospitalInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13"/><path d="M13.667 17h-3.334v-2.333H8v-3.334h2.333V9h3.334v2.333H16v3.334h-2.333V17Z"/></g>`
	homeSaleInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13"/><path d="M14 9.846c-1-.923-3.667-1.23-3.667.616S14 11.385 14 13.23s-3 1.846-4 .615m2 .857V16m0-6.887V8"/></g>`
	homeSecureInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13"/><path d="M14 12h.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6v-2.8a.6.6 0 0 1 .6-.6h.4m4 0v-2c0-.667-.4-2-2-2s-2 1.333-2 2v2m4 0h-4"/></g>`
	homeShieldInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13"/><path d="m12.502 9.13l2.049.531c.264.069.45.309.441.582C14.826 15.232 12 16 12 16s-2.826-.768-2.992-5.757a.584.584 0 0 1 .441-.582l2.05-.53a2 2 0 0 1 1.003 0Z"/></g>`
	homeSimpleInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 21H7a4 4 0 0 1-4-4v-6.292a4 4 0 0 1 1.927-3.421l5-3.03a4 4 0 0 1 4.146 0l5 3.03A4 4 0 0 1 21 10.707V17a4 4 0 0 1-4 4Zm-8-4h6"/>`
	homeSimpleDoorInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 21H7a4 4 0 0 1-4-4v-6.292a4 4 0 0 1 1.927-3.421l5-3.03a4 4 0 0 1 4.146 0l5 3.03A4 4 0 0 1 21 10.707V17a4 4 0 0 1-4 4h-2m-6 0v-4a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v4m-6 0h6"/>`
	homeTableInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 7v10M1 7h22M4 10h16m0-3v10"/>`
	homeUserInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.5 9.5L12 4l9.5 5.5"/><path d="M7 21v-1a5 5 0 0 1 10 0v1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	horizDistributionLeftInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 17V7m0 10h-5.4a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6H19m0 10v3m0-13V4M9 17V7m0 10H5.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6H9m0 10v3M9 7V4"/>`
	horizDistributionRightInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 17V7m0 10h5.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H5m0 10v3M5 7V4m10 13V7m0 10h3.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H15m0 10v3m0-13V4"/>`
	horizontalMergeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 12h-8m0 0l3.5-3.5M14 12l3.5 3.5M2 12h8m0 0L6.5 8.5M10 12l-3.5 3.5M10 21V3m4 18V3"/>`
	horizontalSplitInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12H2m0 0l3.5-3.5M2 12l3.5 3.5M14 12h8m0 0l-3.5-3.5M22 12l-3.5 3.5M10 21V3m4 18V3"/>`
	hospitalInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M6.4 8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 0 .6.6h1.8a.6.6 0 0 1 .6.6v11.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6h1.8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.992 8h2m2 0h-2m0 0V6m0 2v2M16 17.01l.01-.011M16 13.01l.01-.011M12 13.01l.01-.011M8 13.01l.01-.011M8 17.01l.01-.011m3.99.011l.01-.011"/></g>`
	hospitalSignInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10ZM8 12h8m-8 0V7m0 5v5m8-5v5m0-5V7"/>`
	hotAirBalloonInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4 9.5c0 4.571 5.714 8 5.714 8h4.572S20 14.071 20 9.5s-3.582-8-8-8s-8 3.429-8 8Z"/><path stroke-linejoin="round" d="M9 2c-3 6 1 15.5 1 15.5M14.884 2c3 6-1 15.5-1 15.5"/><path stroke-linecap="round" d="M13.4 23h-2.8a.6.6 0 0 1-.6-.6v-1.8a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 1-.6.6Z"/></g>`
	hourglassInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12a7 7 0 0 0 7-7H5a7 7 0 0 0 7 7Zm0 0a7 7 0 0 1 7 7H5a7 7 0 0 1 7-7ZM5 2h14M5 22h14"/>`
	htmlFiveInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4 3l1.778 17.09L12 22l6.222-1.91L20 3H4Z"/><path d="M17 7H7.5l.5 4.5h8l-.5 5.5l-3.5 1l-3.5-1l-.25-2.5"/></g>`
	hydrogenInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6ZM10 8v4m0 4v-4m0 0h4m0 0V8m0 4v4"/>`
	iconoirInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Z"/></g>`
	importInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 13v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6M12 3v12m0 0l-3.5-3.5M12 15l3.5-3.5"/>`
	inclinationInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 19H3.41a.6.6 0 0 1-.431-1.016L16.444 4"/><path d="M20 16c-.5-3.5-1-5-3-8"/></g>`
	industryInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 10c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V2m7 19h3v-9h-3v4.5m0 4.5v-4.5m0 4.5H3v-4l3.5-3l4 2.5l4-2.5l3.5 2.5m3-6.5c0-6-4-6-4-6s4 .5 4-2"/>`
	infiniteInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 9l-.25.375M10 9a5 5 0 1 0 0 6l.334-.5M10 9l4 6a5 5 0 1 0 0-6"/>`
	infoEmptyInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11.5v5m0-8.99l.01-.011M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	inputFieldInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2Zm1 2.5h1.5m1.5 0H6.5m0 0v7m0 0H5m1.5 0H8"/>`
	inputOutputInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-miterlimit="1.5" d="M14 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path d="M3 19V5"/></g>`
	inputSearchInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12v-2a5 5 0 0 0-5-5H8a5 5 0 0 0-5 5v0a5 5 0 0 0 5 5h4m8.124 4.119a3 3 0 1 0-4.248-4.237a3 3 0 0 0 4.248 4.237Zm0 0L22 21"/>`
	instagramInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M3 16V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17.5 6.51l.01-.011"/></g>`
	internetInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 2.05S16 6 16 12m-5 9.95S8 18 8 12c0-6 3-9.95 3-9.95M2.63 15.5H12m-9.37-7h18.74"/><path d="M21.879 17.917c.494.304.463 1.043-.045 1.101l-2.567.291l-1.151 2.312c-.228.459-.933.234-1.05-.334l-1.255-6.116c-.099-.48.333-.782.75-.525l5.318 3.271Z" clip-rule="evenodd"/></g>`
	intersectInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 13.5v3M13.5 21h3m0-12H9.6a.6.6 0 0 0-.6.6v6.9m1.5 4.5h-.9a.6.6 0 0 1-.6-.6v-.9m12 0v.9a.6.6 0 0 1-.6.6h-.9m0-12h.9a.6.6 0 0 1 .6.6v.9m-18 0v-3M7.5 3h3"/><path d="M7.5 15h6.9a.6.6 0 0 0 .6-.6V7.5M4.5 15h-.9a.6.6 0 0 1-.6-.6v-.9m0-9v-.9a.6.6 0 0 1 .6-.6h.9m9 0h.9a.6.6 0 0 1 .6.6v.9"/></g>`
	intersectAltInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.01 3l-.01.011M11.01 3l-.01.011M7.01 3L7 3.011M3.01 3L3 3.011M3.01 7L3 7.011M3.01 11l-.01.011M3.01 15l-.01.011m6 5.999l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M21 17.01l.01-.011M21 13.01l.01-.011M21 9.01l.01-.011M9 17v-7a1 1 0 0 1 1-1h7"/><path d="M15 7v7a1 1 0 0 1-1 1H7"/></g>`
	iosSettingsInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm6-6h-6M9 6.803L12 12m0 0l-3 5.197"/><path stroke-dasharray="1 3" d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	ipAddressInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v6M9 9v6m3-3h2.5a1.5 1.5 0 0 0 1.5-1.5v0A1.5 1.5 0 0 0 14.5 9H12"/></g>`
	irisScanInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 3H3v3m9 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M21 12c-1.889 2.991-5.282 6-9 6s-7.111-3.009-9-6c2.299-2.842 4.992-6 9-6s6.701 3.158 9 6Zm-3-9h3v3M6 21H3v-3m15 3h3v-3"/></g>`
	italicInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 5h3m3 0h-3m0 0l-4 14m0 0H7m3 0h3"/>`
	italicSquareInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 6h2m2 0h-2m0 0l-4 12m0 0H8m2 0h2"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	journalInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6h12M6 10h12m-5 4h5m-5 4h5M2 21.4V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v18.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path d="M6 18v-4h3v4H6Z"/></g>`
	journalPageInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6h8m-8 4h12m-5 4h5m-5 4h5M2 21.4V2.6a.6.6 0 0 1 .6-.6h15.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 22 5.75V21.4a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path d="M6 18v-4h3v4H6ZM18 2v3.4a.6.6 0 0 0 .6.6H22"/></g>`
	jpegFormatInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M7 15v-3m0 0V9h3v3H7Zm9-3h-3v6h3m6-6h-3v6h3v-2.4M4 9v4.2C4 15 2 15 2 15m11-3h2"/><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/></g>`
	jpgFormatInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M10 15v-3m0 0V9h3v3h-3Zm9-3h-3v6h3v-2.4M7 9v4.2C7 15 5 15 5 15"/></g>`
	kanbanBoardInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3.6v16.8a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6ZM6 6v10m4-10v3m4-3v7m4-7v5"/>`
	keyAltInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 0h12v3m-4-3v3"/>`
	keyAltBackInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm0 0H2v3m4-3v3"/>`
	keyAltMinusInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.992 18h6m-8.58-7.657a4 4 0 1 0 5.657-5.657a4 4 0 0 0-5.657 5.657Zm0 0l-8.485 8.485l2.121 2.122M6.755 16l2.122 2.121"/>`
	keyAltPlusInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.992 18h3m3 0h-3m0 0v-3m0 3v3m-5.58-10.657a4 4 0 1 0 5.657-5.657a4 4 0 0 0-5.657 5.657Zm0 0l-8.485 8.485l2.121 2.122M6.755 16l2.122 2.121"/>`
	keyAltRemoveInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.87 20.121L17.993 18m2.121-2.121L17.993 18m0 0l-2.122-2.121M17.992 18l2.121 2.121m-7.701-9.778a4 4 0 1 0 5.657-5.657a4 4 0 0 0-5.657 5.657Zm0 0l-8.485 8.485l2.121 2.122M6.755 16l2.122 2.121"/>`
	keyCommandInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6v12m6-12v12M9 6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3M9 18a3 3 0 1 1-3-3h12a3 3 0 1 1-3 3"/>`
	keyframeInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18.819 13.329l-5.324 5.99a2 2 0 0 1-2.99 0l-5.324-5.99a2 2 0 0 1 0-2.658l5.324-5.99a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.658Z"/>`
	keyframeAlignCenterInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.7 12.375l-3.231 4.04a.6.6 0 0 1-.938 0L8.3 12.374a.6.6 0 0 1 0-.75l3.231-4.04a.6.6 0 0 1 .938 0l3.231 4.04a.6.6 0 0 1 0 .75ZM12 22v-2m0-16V2M4 12H2m20 0h-2"/>`
	keyframeAlignHorizontalInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.7 12.375l-3.231 4.04a.6.6 0 0 1-.938 0L8.3 12.374a.6.6 0 0 1 0-.75l3.231-4.04a.6.6 0 0 1 .938 0l3.231 4.04a.6.6 0 0 1 0 .75ZM4 12H2m20 0h-2"/>`
	keyframeAlignVerticalInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.7 12.375l-3.231 4.04a.6.6 0 0 1-.938 0L8.3 12.374a.6.6 0 0 1 0-.75l3.231-4.04a.6.6 0 0 1 .938 0l3.231 4.04a.6.6 0 0 1 0 .75ZM12 22v-2m0-16V2"/>`
	keyframePositionInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.73 8.36l-2.25 3a.6.6 0 0 1-.96 0l-2.25-3a.6.6 0 0 1 0-.72l2.25-3a.6.6 0 0 1 .96 0l2.25 3a.6.6 0 0 1 0 .72ZM3 20h18m-9-3v-2"/>`
	keyframesInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.848 13.317L9.505 18.28a2 2 0 0 1-3.01 0l-4.343-4.963a2 2 0 0 1 0-2.634L6.495 5.72a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634Z"/><path d="m13 19l4.884-5.698a2 2 0 0 0 0-2.604L13 5"/><path d="m17 19l4.884-5.698a2 2 0 0 0 0-2.604L17 5"/></g>`
	keyframesCoupleInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.819 13.329l-5.324 5.99a2 2 0 0 1-2.99 0l-5.324-5.99a2 2 0 0 1 0-2.658l5.324-5.99a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.658Z"/><path d="m12 6.375l1.505-1.693a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.657l-5.324 5.99a2 2 0 0 1-2.99 0L12 17.624"/></g>`
	labelInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 17.4V6.6a.6.6 0 0 1 .6-.6h13.079c.2 0 .388.1.5.267l3.6 5.4a.6.6 0 0 1 0 .666l-3.6 5.4a.6.6 0 0 1-.5.267H3.6a.6.6 0 0 1-.6-.6Z"/>`
	lampInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 21h3m3 0h-3m0 0V11m0-4v4m0 0H6l3-8h6l3 8h-6Z"/>`
	languageInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Z"/><path d="M13 2.05S16 6 16 12c0 6-3 9.95-3 9.95m-2 0S8 18 8 12c0-6 3-9.95 3-9.95M2.63 15.5h18.74m-18.74-7h18.74"/></g>`
	laptopInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.2 14.222V4a2 2 0 0 1 2-2h13.6a2 2 0 0 1 2 2v10.222m-17.6 0h17.6m-17.6 0l-1.48 5.234A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234"/><path stroke-linecap="round" stroke-linejoin="round" d="M11 19h2"/></g>`
	laptopChargingInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.2 14.222V4a2 2 0 0 1 2-2h13.6a2 2 0 0 1 2 2v10.222m-17.6 0h17.6m-17.6 0l-1.48 5.234A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.667 5L10 8h4l-1.667 3M11 19h2"/></g>`
	laptopFixInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.8 14.222H3.654a.6.6 0 0 0-.578.437L1.72 19.456A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234Zm0 0v-6.11m-17.6 6.11V4a2 2 0 0 1 2-2H12m-1 17h2m4.657-14.172l-2.829 2.829m5.657-2.829A2 2 0 1 1 17.657 2m-2.829 8.485A2 2 0 0 0 12 7.657"/>`
	laptopIssueInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.2 14.222V4a2 2 0 0 1 2-2h13.6a2 2 0 0 1 2 2v10.222m-17.6 0h17.6m-17.6 0l-1.48 5.234A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5v3m0 3.01l.01-.011M11 19h2"/></g>`
	largeSuitcaseInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M8 7H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-4M8 7V3.6a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6V7M8 7h8"/>`
	layoutLeftInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3.6 3h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6Zm6.15 6.75V21M3 9.75h18"/>`
	layoutRightInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M20.4 3H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6Zm-6.15 6.75V21M21 9.75H3"/>`
	leaderboardInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 19H9V8.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6V19Zm0-14H9m11.4 14H15v-3.9a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v3.3a.6.6 0 0 1-.6.6ZM9 19v-5.9a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v5.3a.6.6 0 0 0 .6.6H9Z"/>`
	leaderboardStarInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 21H9v-8.4a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6V21Zm5.4 0H15v-2.9a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6ZM9 21v-4.9a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v4.3a.6.6 0 0 0 .6.6H9Zm1.806-15.887l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.176.568l-1.47 1.5l.347 2.118c.044.272-.228.48-.462.351l-1.818-1l-1.818 1c-.233.128-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.175-.568l2.034-.31Z"/>`
	leafInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 21s.5-4.5 4-8.5"/><path d="m19.13 4.242l.594 6.175c.374 3.886-2.54 7.346-6.425 7.72c-3.813.367-7.267-2.42-7.634-6.233a6.936 6.936 0 0 1 6.239-7.569l6.571-.632a.6.6 0 0 1 .655.54Z"/></g>`
	learningInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.818 22v-2.857C6.52 16.166 3 14.572 3 10c0-4.57 2.727-8.056 8.182-8c3.927.042 7.636 2.286 7.636 6.858L21 12.286c0 2.286-2.182 2.286-2.182 2.286s.546 5.714-4.364 5.714V22"/><path d="M11 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-dasharray=".3 2" d="M11 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	leftRoundArrowInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.75 12h-10m0 0l2.75 2.75M6.75 12L9.5 9.25"/><path d="M2 15V9a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4Z"/></g>`
	lensInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M17.197 9c-.1-.172-.207-.34-.323-.5m.937 5a6.01 6.01 0 0 1-4.311 4.311"/></g>`
	lifebeltInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M8 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm1.235 2.89L5 19m9.765-4.11L19 19m-4.235-9.89L19 5M9.235 9.11L5 5"/></g>`
	lightBulbInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 18h6m-5 3h4m-5-6c.001-2-.499-2.5-1.5-3.5c-1-1-1.476-2.013-1.5-3.5c-.047-3.05 2-5 6-5c4.001 0 6.049 1.95 6 5c-.023 1.487-.5 2.5-1.5 3.5c-.999 1-1.499 1.5-1.5 3.5"/>`
	lightBulbOffInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 18h6m-5 3h4m2.5-9.5c1-1 1.477-2.013 1.5-3.5c.048-3.05-2-5-6-5c-1.168 0-2.169.166-3 .477M9 15c0-2-.5-2.5-1.5-3.5S6.023 9.487 6 8a5.618 5.618 0 0 1 .168-1.5M3 3l18 18"/>`
	lightBulbOnInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 2l-1 1M3 2l1 1m17 13l-1-1M3 16l1-1m5 3h6m-5 3h4M12 3C8 3 5.952 4.95 6 8c.023 1.487.5 2.5 1.5 3.5S9 13 9 15h6c0-2 .5-2.5 1.5-3.5h0c1-1 1.477-2.013 1.5-3.5c.048-3.05-2-5-6-5Z"/>`
	lineSpaceInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 7h9m-9 5h9m-9 5h9M6 17V7m0 10l-2-2.5M6 17l2-2.5M6 7L4 9m2-2l2 2"/>`
	linearInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20L21 4"/>`
	linkInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 11.998C14 9.506 11.683 7 8.857 7H7.143C4.303 7 2 9.238 2 11.998c0 2.378 1.71 4.368 4 4.873a5.3 5.3 0 0 0 1.143.124"/><path d="M10 11.998c0 2.491 2.317 4.997 5.143 4.997h1.714c2.84 0 5.143-2.237 5.143-4.997c0-2.379-1.71-4.37-4-4.874A5.304 5.304 0 0 0 16.857 7"/></g>`
	linkedinInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5ZM7 17v-7"/><path d="M11 17v-3.25M11 10v3.75m0 0c0-3.75 6-3.75 6 0V17M7 7.01l.01-.011"/></g>`
	linuxInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M2.5 20c1 0 2-.8 2-2v-7c0-3.5 3.1-7 7.5-7m9.5 16c-1 0-2-.8-2-2v-7c0-3.5-3.1-7-7.5-7"/><path stroke-linejoin="round" d="M12 19c2.761 0 5-1.12 5-2.5S14.761 14 12 14s-5 1.12-5 2.5S9.239 19 12 19Z"/><path stroke-linejoin="round" d="M7.75 15c-.463-.635-.75-1.52-.75-2.5C7 10.567 8.12 9 9.5 9s2.5 1.567 2.5 3.5c0 .455-.062.89-.175 1.29M16.25 15c.463-.635.75-1.52.75-2.5c0-1.933-1.12-3.5-2.5-3.5S12 10.567 12 12.5c0 .455.062.89.175 1.29M9.5 12v2m5-2v2"/></g>`
	listInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 6h12M4 6.01l.01-.011M4 12.01l.01-.011M4 18.01l.01-.011M8 12h12M8 18h12"/>`
	listSelectInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6h11M5 6.01l.01-.011M5 12.01l.01-.011M3.8 17.8l.8.8l2-2M9 12h11M9 18h11"/>`
	loadActionFloppyInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 6.5V5a2 2 0 0 1 2-2h11.172a2 2 0 0 1 1.414.586l2.828 2.828A2 2 0 0 1 21 7.828V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-1.5"/><path d="M8 3h8v5.4a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6V3Zm10 18v-7.4a.6.6 0 0 0-.6-.6H15m-9 8v-3.5m6-5.5H1m0 0l3-3m-3 3l3 3"/></g>`
	lockInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 12h1.4a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6v-6.8a.6.6 0 0 1 .6-.6H8m8 0V8c0-1.333-.8-4-4-4S8 6.667 8 8v4m8 0H8"/>`
	lockKeyInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M14.667 12h.733a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6v-3.8a.6.6 0 0 1 .6-.6h.733m5.334 0V9.5c0-.833-.534-2.5-2.667-2.5S9.333 8.667 9.333 9.5V12m5.334 0H9.333"/><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/></g>`
	lockedBookInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M14 10h.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6v-2.8a.6.6 0 0 1 .6-.6h.4m4 0V8c0-.667-.4-2-2-2s-2 1.333-2 2v2m4 0h-4"/><path d="M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/></g>`
	lockedWindowInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M21.167 18.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/></g>`
	logDeniedInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.857 9.2a4 4 0 0 0-5.713 5.6m5.713-5.6a4 4 0 0 1-5.713 5.6m5.713-5.6l-5.714 5.6"/><path d="M19 6V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-1"/></g>`
	logInInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 12h-7m0 0l3 3m-3-3l3-3m4-3V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-1"/>`
	logOutInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12h7m0 0l-3 3m3-3l-3-3m3-3V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-1"/>`
	longArrowDownLeftInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.25 19.25l-3.5-3.5l3.5-3.5"/><path d="M6.75 15.75h6a4 4 0 0 0 4-4v-7"/></g>`
	longArrowDownRightInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.25 19.25l3.5-3.5l-3.5-3.5"/><path d="M16.75 15.75h-6a4 4 0 0 1-4-4v-7"/></g>`
	longArrowLeftDownInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 13.5L8 17l3.5-3.5"/><path d="M8 17v-6a4 4 0 0 1 4-4h7"/></g>`
	longArrowLeftUpInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 10.5L8 7l3.5 3.5"/><path d="M8 7v6a4 4 0 0 0 4 4h7"/></g>`
	longArrowRightDownInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 13.5L15.5 17L12 13.5"/><path d="M15.5 17v-6a4 4 0 0 0-4-4h-7"/></g>`
	longArrowRightUpInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 10.5L15.5 7L12 10.5"/><path d="M15.5 7v6a4 4 0 0 1-4 4h-7"/></g>`
	longArrowRightUpOneInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 7v6a4 4 0 0 1-4 4h-7m11-10l3.5 3.5M15.5 7L12 10.5"/>`
	longArrowUpLeftInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.25 4.75l-3.5 3.5l3.5 3.5"/><path d="M6.75 8.25h6a4 4 0 0 1 4 4v7"/></g>`
	longArrowUpRightInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.25 4.75l3.5 3.5l-3.5 3.5"/><path d="M16.75 8.25h-6a4 4 0 0 0-4 4v7"/></g>`
	lotOfCashInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1 16V4h18"/><path d="M16 8h6.4a.6.6 0 0 1 .6.6v10.8a.6.6 0 0 1-.6.6H16m0-12v12m0-12h-4m4 12h-4m0 0H5.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6H12m0 12V8"/></g>`
	macControlKeyInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 14l4-4l4 4"/></g>`
	macDockInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8 17a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm4 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm4 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path d="M21 21H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1ZM2 17.5l2-1m18 1l-2-1"/></g>`
	macOptionKeyInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 10h3m0 4h-5l-1.667-4H7"/></g>`
	macOsWindowInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 17V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m6 8.01l.01-.011M8 8.01l.01-.011M10 8.01l.01-.011"/></g>`
	magnetInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4v8.296C4 16.551 7.582 20 12 20s8-3.45 8-7.704V4"/><path d="M4 4h5.63v6.818C9.63 12.023 10.69 13 12 13s2.37-.977 2.37-2.182V4H20M9 8H4m16 0h-5"/></g>`
	magnetEnergyInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M5 9v6.74C5 19.199 8.134 22 12 22s7-2.802 7-6.26V9M5 9h3m8 0h3"/><path stroke-linecap="round" d="M14.074 11.5v3.56c0 1.072-.928 1.94-2.074 1.94s-2.074-.868-2.074-1.94V11.5"/><path d="M10 13H5m14 0h-5"/><path stroke-linecap="round" d="M11.667 2L10 5h4l-1.667 3"/></g>`
	mailInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m7 9l5 3.5L17 9"/><path d="M2 17V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/></g>`
	mailInInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m5 9l4.5 3L14 9"/><path d="M17 19H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v2"/><path stroke-linejoin="round" d="M23 14h-6m0 0l3-3m-3 3l3 3"/></g>`
	mailOpenedInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m7 12l5 3.5l5-3.5"/><path d="M2 20V9.132a2 2 0 0 1 .971-1.715l8-4.8a2 2 0 0 1 2.058 0l8 4.8A2 2 0 0 1 22 9.132V20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/></g>`
	mailOutInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m5 9l4.5 3L14 9"/><path d="M17 19H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v2"/><path stroke-linejoin="round" d="M17 14h6m0 0l-3-3m3 3l-3 3"/></g>`
	maleInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.232 9.747a6 6 0 1 0-8.465 8.506a6 6 0 0 0 8.465-8.506Zm0 0L20 4m0 0h-4m4 0v4"/>`
	mapInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 19l-5.21 1.737a.6.6 0 0 1-.79-.57V5.433a.6.6 0 0 1 .41-.569L9 3m0 16l6 2m-6-2V3m6 18l5.59-1.863a.6.6 0 0 0 .41-.57V3.832a.6.6 0 0 0-.79-.569L15 5m0 16V5m0 0L9 3"/>`
	mapIssueInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 19l-5.21 1.737a.6.6 0 0 1-.79-.57V5.433a.6.6 0 0 1 .41-.569L9 3m0 16l5.21 1.737a.6.6 0 0 0 .79-.57V5.433a.6.6 0 0 0-.41-.569L9 3m0 16V3m6 2l5.21-1.737a.6.6 0 0 1 .79.57V15m-3.879 7.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L17.12 18.12m2.122 2.122l2.121 2.121"/>`
	mapsArrowInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.685 18.783l7.88-14.008a.5.5 0 0 1 .87 0l7.88 14.008a.5.5 0 0 1-.617.71l-7.517-2.922a.5.5 0 0 0-.362 0l-7.517 2.923a.5.5 0 0 1-.617-.711Z"/>`
	mapsArrowDiagonalInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.031 8.917l15.477-4.334a.5.5 0 0 1 .616.617l-4.333 15.476a.5.5 0 0 1-.94.067l-3.248-7.382a.5.5 0 0 0-.256-.257L3.965 9.856a.5.5 0 0 1 .066-.94Z"/>`
	mapsArrowIssueInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 17.278l-1.819-.707a.5.5 0 0 0-.362 0l-7.517 2.923a.5.5 0 0 1-.617-.711l7.88-14.008a.5.5 0 0 1 .87 0l6.065 10.78m-1.379 6.809l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L17.12 18.12m2.122 2.122l2.121 2.121"/>`
	mapsGoStraightInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865ZM12 10.5V4m0 0L8 6.5M12 4l4 2.5"/>`
	mapsTurnBackInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865ZM5.5 11V6v0s0-3.5 3-3.5C12 2.5 12 6 12 6v4.5"/><path d="M9 7.5L5.5 11L2 7.5"/></g>`
	mapsTurnLeftInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865ZM5 6.5h3s0 0 0 0s4 0 4 4"/><path d="M8.5 9L5 6.5L8.5 4"/></g>`
	mapsTurnRightInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865ZM19 6.5h-3s0 0 0 0s-4 0-4 4"/><path d="M15.5 9L19 6.5L15.5 4"/></g>`
	maskSquareInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M10 17.659a6 6 0 0 0 4-11.317m-4 11.317a6 6 0 1 1 4-11.317m-4 11.317L14 6.34"/></g>`
	mastercardCardInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 9v9.4a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6V9Zm0 0H6"/><path d="M16.5 13.382a1.5 1.5 0 1 1 0 2.236m0-2.236a1.5 1.5 0 1 0 0 2.236"/></g>`
	mathBookInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/><path d="M10 10h4"/><path stroke-linejoin="round" d="m12 13.01l.01-.011M12 7.01l.01-.011"/></g>`
	maximizeInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4H4v3m13-3h3v3M7 20H4v-3m13 3h3v-3"/>`
	medalInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.272 10.445L18 2m-8.684 8.632L5 2m7.761 8.048L8.835 2m5.525 0l-1.04 2.5M6 16a6 6 0 1 0 12 0a6 6 0 0 0-12 0Z"/>`
	mediaImageInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="m3 16l7-3l11 5m-5-8a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g>`
	mediaImageFolderInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 12.6v7.8a.6.6 0 0 1-.6.6h-7.8a.6.6 0 0 1-.6-.6v-7.8a.6.6 0 0 1 .6-.6h7.8a.6.6 0 0 1 .6.6Zm-2.5 1.91l.01-.011"/><path d="m13 18.2l3.5-1.2l5.5 2M2 10V3.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6V9M2 10v8.4a.6.6 0 0 0 .6.6H10m-8-9h8"/></g>`
	mediaImageListInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.6v12.8a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6h12.8a.6.6 0 0 1 .6.6Z"/><path d="M18 4H4.6a.6.6 0 0 0-.6.6V18m3-1.2l5.444-1.8L21 18m-4.5-5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/></g>`
	mediaVideoInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M9.898 8.513a.6.6 0 0 0-.898.52v5.933a.6.6 0 0 0 .898.521l5.19-2.966a.6.6 0 0 0 0-1.042l-5.19-2.966Z"/></g>`
	mediaVideoFolderInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 12.6v7.8a.6.6 0 0 1-.6.6h-7.8a.6.6 0 0 1-.6-.6v-7.8a.6.6 0 0 1 .6-.6h7.8a.6.6 0 0 1 .6.6Z"/><path d="M16.918 14.574a.6.6 0 0 0-.918.508v2.835a.6.6 0 0 0 .918.51l2.268-1.418a.6.6 0 0 0 0-1.018l-2.268-1.417ZM2 10V3.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6V9M2 10v8.4a.6.6 0 0 0 .6.6H10m-8-9h8"/></g>`
	mediaVideoListInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.6v12.8a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6h12.8a.6.6 0 0 1 .6.6Z"/><path d="M18 4H4.6a.6.6 0 0 0-.6.6V18m8.909-6.455a.6.6 0 0 0-.909.515v3.88a.6.6 0 0 0 .909.515l3.233-1.94a.6.6 0 0 0 0-1.03l-3.233-1.94Z"/></g>`
	mediumInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm8 0c1.105 0 2-1.79 2-4s-.895-4-2-4s-2 1.79-2 4s.895 4 2 4Zm5 0c.552 0 1-1.79 1-4s-.448-4-1-4s-1 1.79-1 4s.448 4 1 4Z"/>`
	mediumPriorityInnerSVG                    = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirMediumPriority0)"><path d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848L11.576 1.424ZM6 12h4m4 0h4"/></g><defs><clipPath id="iconoirMediumPriority0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	megaphoneInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M14 14V6m0 8l6.102 3.487a.6.6 0 0 0 .898-.52V3.033a.6.6 0 0 0-.898-.521L14 6m0 8H7a4 4 0 1 1 0-8h7M7.757 19.3L7 14h4l.677 4.74a1.98 1.98 0 0 1-3.92.56Z"/>`
	menuInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5h18M3 12h18M3 19h18"/>`
	menuScaleInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5h8m-8 7h13M3 19h18"/>`
	messageInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 20.29V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.961a2 2 0 0 0-1.561.75l-2.331 2.914A.6.6 0 0 1 3 20.29Z"/>`
	messageAlertInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 7v2m0 4.01l.01-.011"/><path d="M3 20.29V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.961a2 2 0 0 0-1.561.75l-2.331 2.914A.6.6 0 0 1 3 20.29Z"/></g>`
	messageTextInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 12h10M7 8h6"/><path d="M3 20.29V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.961a2 2 0 0 0-1.561.75l-2.331 2.914A.6.6 0 0 1 3 20.29Z"/></g>`
	metroInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m14 16.01l.01-.011M10 16.01l.01-.011M22 12v3a5 5 0 0 1-5 5H7a5 5 0 0 1-5-5v-3C2 6.477 6.477 2 12 2s10 4.477 10 10Z"/><path stroke-linejoin="round" d="M18 12v3a5 5 0 0 1-5 5h-2a5 5 0 0 1-5-5v-3a5 5 0 0 1 5-5h2a5 5 0 0 1 5 5Z"/><path d="m10.5 20l-2 2.5m5-2.5l2 2.5m1-2.5l2 2.5M7.5 20l-2 2.5"/><path stroke-linejoin="round" d="M11.786 10h.428C13.2 10 14 10.8 14 11.786a.214.214 0 0 1-.214.214h-3.572a.214.214 0 0 1-.214-.214C10 10.8 10.8 10 11.786 10Z"/></g>`
	micInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="6" height="12" x="9" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H9m3 0h3"/></g>`
	micAddInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.992 19h3m3 0h-3m0 0v-3m0 3v3"/><rect width="6" height="12" x="5" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`
	micCheckInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m15.5 20.5l2 2l5-5"/><rect width="6" height="12" x="5" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`
	micMuteInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 3l18 18M9 9v0a5 5 0 0 0 5 5v0m1-3.5V5a3 3 0 0 0-3-3v0a3 3 0 0 0-3 3v.5"/><path d="M5 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H9m3 0h3"/></g>`
	micRemoveInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.992 19h6"/><rect width="6" height="12" x="5" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`
	micSpeakingInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="6" height="12" x="9" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 3v2M1 2v4m18-3v2m4-3v4M5 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H9m3 0h3"/></g>`
	micWarningInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 14v4m0 4.01l.01-.011"/><rect width="6" height="12" x="7" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H7m3 0h3"/></g>`
	microscopeInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 22H7m-2 0h2m0 0v-3m12-3h-9m6-14h-4m0 5c-3 0-5 1-5 4v2m9-8.4v6.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6Z"/><path d="M7 19a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/></g>`
	minusInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h12"/>`
	minusCircleInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h8m-4 10c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	minusHexagonInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6M11.7 1.173a.6.6 0 0 1 .6 0l8.926 5.154a.6.6 0 0 1 .3.52v10.307a.6.6 0 0 1-.3.52L12.3 22.826a.6.6 0 0 1-.6 0l-8.926-5.154a.6.6 0 0 1-.3-.52V6.847a.6.6 0 0 1 .3-.52L11.7 1.174Z"/>`
	minusOneInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h12"/>`
	minusPinAltInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 9.2C16 13.177 9 20 9 20S2 13.177 2 9.2C2 5.224 5.134 2 9 2s7 3.224 7 7.2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 19h6"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	minusSquareInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m6-8.4v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	mirrorInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 4v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2Zm0 1l-6 5m6-1l-7.5 6"/>`
	missingFontInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.469 18.374l1.064-2.341m9.58 2.341l-1.064-2.341m0 0L8.79 6.664l-4.258 9.367m8.516 0H4.533m10.645-7.237c0-3.725 5.854-3.725 5.854 0c0 2.661-2.66 2.13-2.66 5.322m-.001 4.269l.01-.012"/>`
	modernTvInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10"/><path d="M2 16.4V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`
	modernTvFourKInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10M13.5 7v4m0 2v-2m0 0l1.37-1.566M17 7l-2.13 2.434m0 0L17 13M9.5 7l-3 4.5H10V13"/><path d="M2 16.4V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`
	moneySquareInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 8.5c-.685-.685-1.891-1.161-3-1.191M9 15c.644.86 1.843 1.35 3 1.391m0-9.082c-1.32-.036-2.5.561-2.5 2.191c0 3 5.5 1.5 5.5 4.5c0 1.711-1.464 2.446-3 2.391m0-9.082V5.5m0 10.891V18.5"/></g>`
	moonSatInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path d="M7.633 3.067A3.001 3.001 0 1 1 4.017 6.32M22 13.05a3.5 3.5 0 1 0-3 5.914"/><path stroke-linecap="round" stroke-linejoin="round" d="m14.5 8.51l.01-.011M10 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	moreHorizInnerSVG                         = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-6 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-6 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`
	moreHorizCircleInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M7 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	moreVertInnerSVG                          = `<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm0 6a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm0-12a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`
	moreVertCircleInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 7.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm0 10a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm0-5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	motorcycleInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14-4l-3-9l1-1"/><path d="M16 8.5h-4.5l-4.5 3m-1.5 4H12l1-2.5l3.5-4.5m-8 1.5c-2-1.5-5-1.5-7 0M19 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`
	mouseButtonLeftInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20 10v4a8 8 0 1 1-16 0V9a7 7 0 0 1 7-7h1a8 8 0 0 1 8 8Z"/><path d="M12 2v6.4a.6.6 0 0 1-.6.6H4"/></g>`
	mouseButtonRightInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 10v4a8 8 0 1 0 16 0V9a7 7 0 0 0-7-7h-1a8 8 0 0 0-8 8Z"/><path d="M12 2v6.4a.6.6 0 0 0 .6.6H20"/></g>`
	mouseScrollWheelInnerSVG                  = `<path fill="currentColor" d="m12 5l.53-.53a.75.75 0 0 0-1.06 0L12 5Zm0 8l-.53.53a.75.75 0 0 0 1.06 0L12 13ZM9.47 6.47a.75.75 0 0 0 1.06 1.06L9.47 6.47Zm4 1.06a.75.75 0 1 0 1.06-1.06l-1.06 1.06Zm-2.94 2.94a.75.75 0 1 0-1.06 1.06l1.06-1.06Zm4 1.06a.75.75 0 1 0-1.06-1.06l1.06 1.06ZM3.25 10v4h1.5v-4h-1.5Zm17.5 4v-4h-1.5v4h1.5Zm-9.5-9v8h1.5V5h-1.5Zm.22-.53l-2 2l1.06 1.06l2-2l-1.06-1.06Zm0 1.06l2 2l1.06-1.06l-2-2l-1.06 1.06Zm1.06 6.94l-2-2l-1.06 1.06l2 2l1.06-1.06Zm0 1.06l2-2l-1.06-1.06l-2 2l1.06 1.06ZM20.75 10A8.75 8.75 0 0 0 12 1.25v1.5A7.25 7.25 0 0 1 19.25 10h1.5ZM12 22.75A8.75 8.75 0 0 0 20.75 14h-1.5A7.25 7.25 0 0 1 12 21.25v1.5ZM3.25 14A8.75 8.75 0 0 0 12 22.75v-1.5A7.25 7.25 0 0 1 4.75 14h-1.5Zm1.5-4A7.25 7.25 0 0 1 12 2.75v-1.5A8.75 8.75 0 0 0 3.25 10h1.5Z"/>`
	moveDownInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 3v13m0 0l3-3m-3 3l-3-3"/>`
	moveLeftInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-5-2H2m0 0l3-3m-3 3l3 3"/>`
	moveRightInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm5-2h13m0 0l-3-3m3 3l-3 3"/>`
	moveRulerInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.4 22H8.6a.6.6 0 0 1-.6-.6V2.6a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v18.8a.6.6 0 0 1-.6.6Zm.6-5h-3m3-10h-3m0 5h10m0 0l-2 2m2-2l-2-2M1 12l2-2m-2 2l2 2m-2-2h7"/>`
	moveUpInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0-7V2m0 0l3 3m-3-3L9 5"/>`
	movieInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 8.01l.01-.011M17 8.01l.01-.011M7 12.01l.01-.011m9.99.011l.01-.011M7 16.01l.01-.011m9.99.011l.01-.011M7 2H3.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H7M7 2v2m0-2h10m0 0h3.4a.6.6 0 0 1 .6.6v18.8a.6.6 0 0 1-.6.6H17m0-20v2m0 18v-2m0 2H7m0 0v-2"/>`
	mpegFormatInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M7.5 15v-3m0 0V9h3v3h-3Zm-6 3V9L3 12l1.5-3v6m12-6h-3v6h3m6-6h-3v6h3v-2.4m-9-.6h2"/></g>`
	multiBubbleInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 22a5.5 5.5 0 1 0-4.764-2.75l-.461 2.475l2.475-.46A5.474 5.474 0 0 0 7.5 22Z"/><path d="M15.282 17.898A7.946 7.946 0 0 0 18 16.93l3.6.67l-.67-3.6A8 8 0 1 0 6.083 8.849"/></g>`
	multiMacOsWindowInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 19v-8a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m10 12.01l.01-.011m1.99.011l.01-.011m1.99.011l.01-.011"/><path d="M6.5 16H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v3"/><path stroke-linecap="round" stroke-linejoin="round" d="m5 7.01l.01-.011M7 7.01l.01-.011M9 7.01l.01-.011"/></g>`
	multiWindowInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 19v-8a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2Z"/><path d="M6.5 16H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v3"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 12h1M5 7h1"/></g>`
	multiplePagesInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 18h7m-7-4h1m-1-4h3M7 2h9.5L21 6.5V19"/><path d="M3 20.5v-14A1.5 1.5 0 0 1 4.5 5h9.752a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 18 8.75V20.5a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 3 20.5Z"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`
	multiplePagesAddInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.992 19h3m3 0h-3m0 0v-3m0 3v3M7 2h9.5L21 6.5V19"/><path d="M11 22h5.5a1.5 1.5 0 0 0 1.5-1.5V8.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 14.25 5H4.5A1.5 1.5 0 0 0 3 6.5V13"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`
	multiplePagesDeleteInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.87 21.121L4.993 19m2.121-2.121L4.993 19m0 0L2.87 16.879M4.992 19l2.121 2.121M7 2h9.5L21 6.5V19"/><path d="M11 22h5.5a1.5 1.5 0 0 0 1.5-1.5V8.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 14.25 5H4.5A1.5 1.5 0 0 0 3 6.5V13"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`
	multiplePagesEmptyInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 2h9.5L21 6.5V19"/><path d="M3 20.5v-14A1.5 1.5 0 0 1 4.5 5h9.752a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 18 8.75V20.5a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 3 20.5Z"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`
	multiplePagesRemoveInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 2h9.5L21 6.5V19"/><path d="M11 22h5.5a1.5 1.5 0 0 0 1.5-1.5V8.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 14.25 5H4.5A1.5 1.5 0 0 0 3 6.5V13m-1.008 6h6"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`
	musicDoubleNoteInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 14V3L9 5v11"/><path d="M17 19h1a2 2 0 0 0 2-2v-3h-3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2ZM6 21h1a2 2 0 0 0 2-2v-3H6a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2Z"/></g>`
	musicDoubleNoteAddInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 6.5h3m3 0h-3m0 0v-3m0 3v3M6 16V5l8-1m1 10v-4m-3 9h1a2 2 0 0 0 2-2v-3h-3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2Zm-9 2h1a2 2 0 0 0 2-2v-3H3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2Z"/>`
	musicNoteInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16v3a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h3Zm0 0V8m0 0V4l5-1v4l-5 1Z"/>`
	musicNoteAddInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 10h3m3 0h-3m0 0V7m0 3v3M7 16v3a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h3Zm0 0V8m0 0V4l5-1v4L7 8Z"/>`
	navArrowDownInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 9l6 6l6-6"/>`
	navArrowLeftInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 6l-6 6l6 6"/>`
	navArrowRightInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 6l6 6l-6 6"/>`
	navArrowUpInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 15l6-6l6 6"/>`
	navigatorInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path d="M17.873 15.475c.46.87-.437 1.831-1.336 1.432l-4.538-2.017l-4.537 2.017c-.9.4-1.797-.562-1.337-1.432l4.959-9.365a1.036 1.036 0 0 1 1.831 0l4.958 9.365Z" clip-rule="evenodd"/></g>`
	navigatorAltInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path d="M13.93 17.869c-.322.93-1.637.929-1.958-.001l-1.62-4.694l-4.57-1.943c-.905-.385-.814-1.698.136-1.954L16.15 6.516a1.036 1.036 0 0 1 1.249 1.34L13.93 17.868Z" clip-rule="evenodd"/></g>`
	networkInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" x="3" y="2" rx=".6"/><rect width="7" height="5" x="8.5" y="17" rx=".6"/><rect width="7" height="5" x="14" y="2" rx=".6"/><path d="M6.5 7v3.5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V7M12 12.5V17"/></g>`
	networkAltInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" rx=".6" transform="matrix(1 0 0 -1 3 22)"/><rect width="7" height="5" rx=".6" transform="matrix(1 0 0 -1 8.5 7)"/><rect width="7" height="5" rx=".6" transform="matrix(1 0 0 -1 14 22)"/><path d="M6.5 17v-3.5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2V17M12 11.5V7"/></g>`
	networkLeftInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" x="2" y="21" rx=".6" transform="rotate(-90 2 21)"/><rect width="7" height="5" x="17" y="15.5" rx=".6" transform="rotate(-90 17 15.5)"/><rect width="7" height="5" x="2" y="10" rx=".6" transform="rotate(-90 2 10)"/><path d="M7 17.5h3.5a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2H7m5.5 5.5H17"/></g>`
	networkRightInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 22 21)"/><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 7 15.5)"/><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 22 10)"/><path d="M17 17.5h-3.5a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2H17M11.5 12H7"/></g>`
	newTabInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 19V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M9 14h3m3 0h-3m0 0v-3m0 3v3"/></g>`
	nintendoSwitchInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 17V7a4 4 0 0 1 4-4h3.9a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H6a4 4 0 0 1-4-4Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm11 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M22 17V7a4 4 0 0 0-4-4h-3.9a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H18a4 4 0 0 0 4-4Z"/></g>`
	nitrogenInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M10 16V8l4 8V8"/></g>`
	noAccessWindowInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M21 16.05a3.5 3.5 0 1 0-5 4.9m4.998-4.9A3.5 3.5 0 1 1 16 20.95m5-4.9l-5 4.9"/></g>`
	noBatteryInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 3l18 18m2-11v4M5.5 6H3a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h14.5m2.5-3.5V8a2 2 0 0 0-2-2h-6.5"/>`
	noCoinInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.623 5.248A9.964 9.964 0 0 0 2 12c0 5.523 4.477 10 10 10a9.962 9.962 0 0 0 6.615-2.5m2.687-3.822A9.974 9.974 0 0 0 22 12c0-5.523-4.477-10-10-10c-1.231 0-2.41.223-3.5.63"/><path d="M9 15c.644.86 1.843 1.35 3 1.391c1.114.04 2.19-.336 2.697-1.198M12 16.391V18.5m-2.5-9c0 1.181.852 1.665 1.886 2M15 8.5c-.685-.685-1.891-1.161-3-1.191V5.5M3 3l18 18"/></g>`
	noCreditCardInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 5H2.6a.6.6 0 0 0-.6.6v12.8a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V5.6a.6.6 0 0 0-.6-.6H10m12 4h-8M6 9h3M3 3l18 18"/>`
	noLinkInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 11.998C14 9.506 11.683 7 8.857 7H7.143C4.303 7 2 9.238 2 11.998c0 2.378 1.71 4.368 4 4.873a5.3 5.3 0 0 0 1.143.124M16.857 7c.393 0 .775.043 1.143.124c2.29.505 4 2.495 4 4.874a4.92 4.92 0 0 1-1.634 3.653"/><path d="M10 11.998c0 2.491 2.317 4.997 5.143 4.997M18 22.243l2.121-2.122m0 0L22.243 18m-2.122 2.121L18 18m2.121 2.121l2.122 2.122"/></g>`
	noLockInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.5 12H6.6a.6.6 0 0 0-.6.6v6.8a.6.6 0 0 0 .6.6h10.8a.6.6 0 0 0 .6-.6v-.9M16 12V8c0-1.333-.8-4-4-4c-.747 0-1.363.145-1.869.385M16 12h1.4a.6.6 0 0 1 .6.6v.4M8 8v4M3 3l18 18"/>`
	noSmokingInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 15v3m0-7c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V2m11 9c0-6-4-6-4-6s4 1 4-3m0 13v3"/><path d="M2.6 18H18l-3-3H2.6a.6.6 0 0 0-.6.6v1.8a.6.6 0 0 0 .6.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m3 3l18 18"/></g>`
	noSmokingCircleInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15 12v3m0-6c0-1-.714-2-2.143-2v0A2.857 2.857 0 0 1 10 4.143V3m8 6V4m0 8v3"/><path d="M15 15H6.6a.6.6 0 0 1-.6-.6v-1.8a.6.6 0 0 1 .6-.6H12"/><path stroke-linecap="round" stroke-linejoin="round" d="m5 5l14 14m-7 3c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	notesInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 14h8m-8-4h2m-2 8h4M10 3H6a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-3.5M10 3V1m0 2v2"/>`
	npmInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M1 8h22v7H11v2H7.5v-2H1V8Zm6.5 0v7m6-7v7"/><path d="M18 11v4M5 11v4m6-4v1m9.5-1v4"/></g>`
	npmSquareInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M8 16h8V8H8v8Z"/><path d="M13 11v5"/><path stroke-linejoin="round" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/></g>`
	numberedListLeftInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6h11M5 8V4m1 10H4.6a.6.6 0 0 1-.6-.6v-.8a.6.6 0 0 1 .6-.6h.8a.6.6 0 0 0 .6-.6v-.8a.6.6 0 0 0-.6-.6H4m0 6h1.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H4m2-2H4m5-6h11M9 18h11"/>`
	numberedListRightInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h11m4 2V4m1 10h-1.4a.6.6 0 0 1-.6-.6v-.8a.6.6 0 0 1 .6-.6h.8a.6.6 0 0 0 .6-.6v-.8a.6.6 0 0 0-.6-.6H18m0 6h1.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H18m2-2h-2M4 12h11M4 18h11"/>`
	octagonInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.77 1.095a.6.6 0 0 1 .46 0l7.319 3.032a.6.6 0 0 1 .324.324l3.032 7.32a.6.6 0 0 1 0 .459l-3.032 7.319a.6.6 0 0 1-.324.324l-7.32 3.032a.6.6 0 0 1-.459 0l-7.319-3.032a.6.6 0 0 1-.324-.324l-3.032-7.32a.6.6 0 0 1 0-.459l3.032-7.319a.6.6 0 0 1 .324-.324l7.32-3.032Z"/>`
	offTagInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 15V9a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6Z"/><path d="M7 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15V9h3m2 6V9h3m-8 3h2.572M17 12h2.572"/></g>`
	oilIndustryInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 10c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V2"/><path d="M9 10.8C9 9.033 6 6 6 6s-3 3.033-3 4.8S4.343 14 6 14s3-1.433 3-3.2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 21h3v-9h-3v4.5m0 4.5v-4.5m0 4.5h-7.5v-4.5l4-2.5l3.5 2.5m3-6.5c0-6-4-6-4-6s4 .5 4-2"/></g>`
	okrsInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-7a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM3 5h10M3 12h10M3 19h10m3 2.243l2.121-2.122m0 0L20.243 17m-2.122 2.121L16 17m2.121 2.121l2.122 2.122"/>`
	onTagInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 15V9a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6Z"/><path d="M9 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 15V9l4 6V9"/></g>`
	oneFingerSelectHandGestureInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7.5 12l-2.004 2.672a2 2 0 0 0 .126 2.552l3.784 4.128c.378.413.912.648 1.473.648H15.5c2.4 0 4-2 4-4c0 0 0 0 0 0V9.429m-3 .571v-.571c0-2.286 3-2.286 3 0"/><path d="M13.5 10V8.286c0-2.286 3-2.286 3 0V10m-6 0V7.5c0-2.286 3-2.286 3 0c0 0 0 0 0 0V10m-3 0V3.499A1.5 1.5 0 0 0 9 2v0a1.5 1.5 0 0 0-1.5 1.5V15"/></g>`
	onePointCircleInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="m19 19l-1.5-1.5m-2-2l-1-1"/></g>`
	oneStMedalInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.272 10.445L18 2m-8.684 8.632L5 2m7.761 8.048L8.835 2m5.525 0l-1.04 2.5M6 16a6 6 0 1 0 12 0a6 6 0 0 0-12 0Z"/><path d="m10.5 15l2-1.5v5"/></g>`
	openBookInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 21V7a2 2 0 0 1 2-2h7.4a.6.6 0 0 1 .6.6v13.114M12 21V7a2 2 0 0 0-2-2H2.6a.6.6 0 0 0-.6.6v13.114M14 19h8m-12 0H2"/><path stroke-linejoin="round" d="M12 21a2 2 0 0 1 2-2m-2 2a2 2 0 0 0-2-2"/></g>`
	openInBrowserInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 21h12.4a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6V16m7-10h8M6 6h1M3.5 20.5L12 12m0 0v4m0-4H8"/>`
	openInWindowInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 21h12.4a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6V16m.5 4.5L12 12m0 0v4m0-4H8"/>`
	openNewWindowInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21 3h-6m6 0l-9 9m9-9v6"/><path d="M21 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6"/></g>`
	openSelectHandGestureInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 14.571l-1.823-1.736a1.558 1.558 0 0 0-2.247.103v0a1.558 1.558 0 0 0 .035 2.092l5.942 6.338c.379.403.906.632 1.459.632H16c2.4 0 4-2 4-4c0 0 0 0 0 0V9.429"/><path d="M17 10v-.571c0-2.286 3-2.286 3 0M14 10V8.286C14 6 17 6 17 8.286V10m-6 0V7.5c0-2.286 3-2.286 3 0c0 0 0 0 0 0V10m-6 4.571V3.5A1.5 1.5 0 0 1 9.5 2v0c.828 0 1.5.67 1.5 1.499V10"/></g>`
	openVpnInnerSVG                           = `<path fill="currentColor" d="m10.835 15.29l.738.136l-.738-.137Zm-.358-.708l.381-.646l-.38.646Zm-.275 7.247l.138-.738l-.138.738Zm-.452-.678l.738.136l-.738-.136Zm7.099-1.337l.737-.139l-.737.14Zm.872.378l-.43-.615l.43.615Zm-9.85-4.208l-.736-.139l.737.14Zm-.139-.52l-.581.474l.581-.474Zm5.791-.882l.382.646l-.382-.646Zm-.358.707l.738-.136l-.738.136Zm3.103.175l-.581-.473l.581.473Zm-.14.52l.737-.139l-.737.14Zm-1.878 5.167l-.737.137l.737-.137Zm-.453.679l.138.737l-.138-.737ZM6.28 20.192l-.43.614l.43-.614ZM2.75 12A9.25 9.25 0 0 1 12 2.75v-1.5C6.062 1.25 1.25 6.062 1.25 12h1.5Zm3.959 7.577C4.315 17.902 2.75 15.137 2.75 12h-1.5c0 3.65 1.824 6.865 4.599 8.806l.86-1.229Zm.426-3.732l-.721 3.83l1.474.278l.72-3.83l-1.473-.278ZM5.75 12c0 1.494.526 2.865 1.4 3.938l1.163-.947A4.713 4.713 0 0 1 7.25 12h-1.5ZM12 5.75A6.25 6.25 0 0 0 5.75 12h1.5A4.75 4.75 0 0 1 12 7.25v-1.5ZM18.25 12A6.25 6.25 0 0 0 12 5.75v1.5A4.75 4.75 0 0 1 16.75 12h1.5Zm-1.4 3.938A6.213 6.213 0 0 0 18.25 12h-1.5a4.713 4.713 0 0 1-1.063 2.99l1.162.948Zm.736 3.737l-.72-3.83l-1.475.278l.72 3.83l1.475-.277ZM21.25 12c0 3.137-1.565 5.902-3.959 7.577l.86 1.23C20.926 18.864 22.75 15.65 22.75 12h-1.5ZM12 2.75A9.25 9.25 0 0 1 21.25 12h1.5c0-5.938-4.812-10.75-10.75-10.75v1.5ZM15.75 12A3.75 3.75 0 0 0 12 8.25v1.5A2.25 2.25 0 0 1 14.25 12h1.5Zm-1.845 3.228A3.745 3.745 0 0 0 15.75 12h-1.5c0 .823-.443 1.544-1.108 1.936l.763 1.292Zm1.083 5.787l-1.085-5.862l-1.475.273l1.085 5.862l1.475-.273ZM12 22.75c.665 0 1.31-.067 1.935-.183l-.275-1.474a9.036 9.036 0 0 1-1.66.157v1.5Zm-1.937-.184c.625.117 1.271.184 1.937.184v-1.5a8.958 8.958 0 0 1-1.66-.159l-.277 1.475Zm.035-7.413l-1.085 5.861l1.475.273l1.085-5.861l-1.475-.273ZM8.25 12c0 1.377.744 2.578 1.846 3.228l.762-1.292A2.245 2.245 0 0 1 9.75 12h-1.5ZM12 8.25A3.75 3.75 0 0 0 8.25 12h1.5A2.25 2.25 0 0 1 12 9.75v-1.5Zm-.427 7.176c.122-.662-.259-1.22-.715-1.49l-.762 1.292a.053.053 0 0 1 .01.008c.003.003.002.003 0-.001a.1.1 0 0 1-.009-.03a.153.153 0 0 1 0-.052l1.476.273Zm-1.233 5.665c.119.023.16.129.148.196l-1.475-.273c-.129.694.305 1.412 1.05 1.552l.277-1.474Zm5.772-1.138c.168.892 1.212 1.432 2.04.853l-.86-1.229a.21.21 0 0 1 .197-.019c.052.023.088.07.097.117l-1.474.278Zm-7.503-3.83a1.387 1.387 0 0 0-.296-1.133l-1.162.948c-.007-.008-.026-.04-.016-.093l1.474.278Zm4.533-2.187c-.456.27-.837.828-.714 1.49l1.475-.273a.153.153 0 0 1 0 .053a.1.1 0 0 1-.009.029c-.002.004-.003.004 0 .001a.05.05 0 0 1 .01-.008l-.762-1.292Zm2.545 1.055c-.245.3-.375.709-.296 1.132l1.474-.278c.01.054-.009.085-.016.093l-1.162-.947Zm-2.174 6.297a.174.174 0 0 1 .147-.195l.275 1.474c.745-.139 1.181-.857 1.053-1.552l-1.475.273Zm-7.664-.482c.827.579 1.871.038 2.04-.853l-1.475-.277a.166.166 0 0 1 .097-.118a.21.21 0 0 1 .198.02l-.86 1.228Z"/>`
	orangeHalfInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.5 0 10-4.5 10-10S17.5 2 12 2m0 20C6.5 22 2 17.5 2 12S6.5 2 12 2m0 20V12m0-10v10m0 0l5 5.5M12 12l5-5m-5 5h7"/>`
	orangeSliceInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.613 10.11l7.778-7.777c4.295 4.296 4.295 11.26 0 15.556c-4.296 4.296-11.261 4.296-15.557 0l7.779-7.778Zm0 0l-.354 8.133m.354-8.132h7.778m-7.778 0l5.303 5.303"/>`
	orangeSliceAltInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.39 10.11L5.61 2.334c-4.295 4.296-4.295 11.26 0 15.556c4.296 4.296 11.26 4.296 15.557 0l-7.778-7.778Zm0 0l.353 8.133m-.354-8.132H5.612m7.779 0l-5.304 5.303"/>`
	organicFoodInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 18s.9-3.741 3-6"/><path d="m16.186 7.241l.374 3.89c.243 2.523-1.649 4.77-4.172 5.012c-2.475.238-4.718-1.571-4.956-4.047a4.503 4.503 0 0 1 4.05-4.914l4.147-.4a.51.51 0 0 1 .557.46Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	organicFoodSquareInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 18s.9-3.741 3-6"/><path d="m16.186 7.241l.374 3.89c.243 2.523-1.649 4.77-4.172 5.012c-2.475.238-4.718-1.571-4.956-4.047a4.503 4.503 0 0 1 4.05-4.914l4.147-.4a.51.51 0 0 1 .557.46Z"/><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/></g>`
	orthogonalViewInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21V3h18v18H3Zm0-4.5h18M3 12h18M3 7.5h18M16.5 3v18M12 3v18M7.5 3v18"/>`
	oxygenInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M12.2 8h-.4A1.8 1.8 0 0 0 10 9.8v4.4a1.8 1.8 0 0 0 1.8 1.8h.4a1.8 1.8 0 0 0 1.8-1.8V9.8A1.8 1.8 0 0 0 12.2 8Z"/></g>`
	packageInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 6v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2Zm-8 3V4"/>`
	packageLockInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 20H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v6m-8-3V4m9.167 14.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/>`
	packagesInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 15v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2Zm6-10v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2Zm6 10v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2ZM6 16v-3m6-7V3m6 13v-3"/>`
	pacmanInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16 12l.011.01M19 12l.011.01M22 12l.011.01M2 12c0 5.523 4.477 10 10 10a9.985 9.985 0 0 0 8-3.999L12 12l8-6.001A9.985 9.985 0 0 0 12 2C6.477 2 2 6.477 2 12Z"/>`
	pageInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6ZM8 10h8m-8 8h8m-8-4h4"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	pageDownInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.5 11l3.5 3.5l3.5-3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	pageEditInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11M8 10h8M8 6h4m-4 8h3m6.954 2.94l1-1a1.121 1.121 0 0 1 1.586 0v0a1.121 1.121 0 0 1 0 1.585l-1 1m-1.586-1.586l-2.991 2.991a1 1 0 0 0-.28.553l-.244 1.557l1.557-.243a1 1 0 0 0 .553-.28l2.99-2.992m-1.585-1.586l1.586 1.586"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	pageFlipInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 11h5m-5-4h5m-9 8V3.6a.6.6 0 0 1 .6-.6h11.8a.6.6 0 0 1 .6.6V17a4 4 0 0 1-4 4v0"/><path d="M5 15h7.4c.331 0 .603.267.63.597C13.153 17.115 13.78 21 17 21H6a3 3 0 0 1-3-3v-1a2 2 0 0 1 2-2Z"/></g>`
	pageLeftInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 8.5L9.5 12l3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	pageRightInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11 8.5l3.5 3.5l-3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	pageSearchInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11M8 10h8M8 6h4m-4 8h3m9.5 6.5L22 22"/><path d="M15 18a3 3 0 1 0 6 0a3 3 0 0 0-6 0Zm1-16v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	pageStarInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11M8 10h8M8 6h4m-4 8h3"/><path d="m16.306 17.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.234.129-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568l2.033-.31ZM16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	pageUpInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 13L12 9.5l3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	paletteInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20.51 9.54a1.899 1.899 0 0 1-1 1.09A7 7 0 0 0 15.37 17c.001.47.048.939.14 1.4a2.16 2.16 0 0 1-.31 1.65a1.79 1.79 0 0 1-1.21.8a9 9 0 0 1-10.62-9.13A9.05 9.05 0 0 1 11.85 3h.51a9 9 0 0 1 8.06 5a2 2 0 0 1 .09 1.52v.02Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 16.01l.01-.011M6 12.01l.01-.011M8 8.01l.01-.011M12 6.01l.01-.011M16 8.01l.01-.011"/></g>`
	panoramaEnlargeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 5c2.995 0 7.235.692 8.576.925a.581.581 0 0 1 .48.503c.13 1.028.444 3.691.444 5.572c0 1.88-.313 4.544-.444 5.572a.581.581 0 0 1-.48.503c-1.34.233-5.58.925-8.576.925c-2.995 0-7.235-.692-8.576-.925a.582.582 0 0 1-.48-.503C2.814 16.544 2.5 13.881 2.5 12c0-1.88.313-4.544.444-5.572a.582.582 0 0 1 .48-.503C4.764 5.692 9.004 5 12 5Z"/>`
	panoramaReduceInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 6.862v10.276a.615.615 0 0 1-.811.58C18.546 17.165 14.749 16 12 16c-2.749 0-6.546 1.166-8.189 1.717a.615.615 0 0 1-.811-.58V6.863c0-.418.415-.712.811-.58C5.454 6.835 9.251 8 12 8c2.749 0 6.546-1.166 8.189-1.717a.615.615 0 0 1 .811.58Z"/>`
	pantsInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5.035 3.633a.6.6 0 0 1 .6-.633h12.73a.6.6 0 0 1 .6.633l-.933 16.8a.6.6 0 0 1-.6.567h-2.898a.6.6 0 0 1-.596-.53L12.596 9.065c-.083-.706-1.109-.706-1.192 0L10.062 20.47a.6.6 0 0 1-.596.53H6.568a.6.6 0 0 1-.6-.567l-.933-16.8Z"/><path d="M5 7.5h1.5a2 2 0 0 0 2-2V3m10 4.5h-1a2 2 0 0 1-2-2V3"/></g>`
	pantsAltInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19h4.436a.6.6 0 0 0 .6-.563l.924-14.8A.6.6 0 0 0 17.361 3H6.634a.6.6 0 0 0-.599.633l.934 16.8a.6.6 0 0 0 .599.567H11.4a.6.6 0 0 0 .6-.6V8"/>`
	parkingInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 15.5v-2.8m0 0h2.857c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H10v4.2Z"/><circle cx="12" cy="12" r="10"/></g>`
	passwordCursorInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 13V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h7"/><path d="M20.879 16.917c.494.304.463 1.043-.045 1.101l-2.567.291l-1.151 2.312c-.228.459-.933.234-1.05-.334l-1.255-6.116c-.099-.48.333-.782.75-.525l5.318 3.271Z" clip-rule="evenodd"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 11.01l.01-.011m3.99.011l.01-.011M8 11.01l.01-.011"/></g>`
	passwordErrorInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.121 20.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L15.12 16.12m2.122 2.122l2.121 2.121M21 13V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6m1-4.99l.01-.011m3.99.011l.01-.011M8 11.01l.01-.011"/>`
	passwordPassInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 13V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h7m2.5 2.5l2 2l4-4M12 11.01l.01-.011m3.99.011l.01-.011M8 11.01l.01-.011"/>`
	pasteClipboardInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M8.5 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-2.5"/><path d="M8 6.4V4.5a.5.5 0 0 1 .5-.5c.276 0 .504-.224.552-.496C9.2 2.652 9.774 1 12 1s2.8 1.652 2.948 2.504c.048.272.276.496.552.496a.5.5 0 0 1 .5.5v1.9a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6Z"/></g>`
	pathArrowInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 16.5V3m0 0l3.5 3.5M18 3l-3.5 3.5m3.5 10a3.5 3.5 0 1 1-7 0v-9m0 0a3.5 3.5 0 1 0-7 0v12"/>`
	pauseInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M6 18.4V5.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6Zm8 0V5.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6Z"/>`
	pauseWindowInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M16 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9"/><path stroke-linejoin="round" d="M19 17v5m3-5v5M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/></g>`
	paypalInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 17.5L6 3h7c6 0 6 9 0 9H8.7l-1.2 5.5H3Z"/><path d="m6.8 21l3-14.5h7c6 0 6 9 0 9h-4.3L11.3 21H6.8Z"/></g>`
	pcCheckInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 22h10"/><path d="M2 17V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 10.5l2 2l4-4"/></g>`
	pcFirewallInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 22h10"/><path d="M2 17V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m12.485 6.121l3.06.765a.59.59 0 0 1 .449.586C15.818 14 12 15 12 15s-3.818-1-3.994-7.528a.59.59 0 0 1 .448-.586l3.06-.765a2 2 0 0 1 .971 0Z"/></g>`
	pcMouseInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 2v0a8 8 0 0 1 8 8v4a8 8 0 0 1-8 8v0a8 8 0 0 1-8-8v-4a8 8 0 0 1 8-8v0Zm0 0v7"/>`
	pcNoEntryInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 22h10"/><path d="M2 17V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14.857 7.7a4 4 0 1 0-5.713 5.6m5.713-5.6a4 4 0 0 1-5.713 5.6m5.713-5.6l-5.714 5.6"/></g>`
	pcWarningInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 22h10"/><path d="M2 17V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 7v3m0 4.01l.01-.011"/></g>`
	peaceHandInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14.149 9.472v-5.86c0-.89-.722-1.612-1.612-1.612v0c-.89 0-1.611.722-1.611 1.612v4.834"/><path d="m16.346 12.841l2.176-7.252a1.584 1.584 0 0 0-1.083-1.98v0a1.585 1.585 0 0 0-1.961 1.098l-1.33 4.764M7.62 9.25l1.055 2.341a1.612 1.612 0 0 1-2.938 1.325L4.68 10.575A1.612 1.612 0 0 1 7.62 9.25Z"/><path d="M11.72 12.261v0a2.322 2.322 0 0 0-.068-1.742l-1.073-2.38a1.584 1.584 0 0 0-2.101-.79v0a1.584 1.584 0 0 0-.764 2.14l.135.276"/><path d="m13.857 17.677l.492-.984a.176.176 0 0 0-.108-.248l-3.55-1.044a1.537 1.537 0 0 1-1.095-1.635v0a1.537 1.537 0 0 1 1.67-1.37l4.788.446s3.81.586 2.49 4.395c-1.318 3.81-1.757 5.128-4.687 5.128H8.876a4.249 4.249 0 0 1-4.249-4.249v0L4.48 9.912"/></g>`
	penConnectBluetoothInnerSVG               = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirPenConnectBluetooth0)"><path d="m6.5 17.5l-1 4l3.731-.933a1 1 0 0 0 .465-.263L21.5 8.5a2.121 2.121 0 0 0 0-3v0a2.121 2.121 0 0 0-3 0l-4 4m3-3l3 3M5 6.6l7 5.1L8.333 15V3L12 6.3l-7 5.1"/></g><defs><clipPath id="iconoirPenConnectBluetooth0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	penConnectWifiInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 9.76l.01-.011M3 6.25c2.5-3 7.5-3 10 0m-8 2c1.5-2 4.5-2 6 0m6.5-1.75l1-1a2.121 2.121 0 0 1 3 0v0a2.121 2.121 0 0 1 0 3l-1 1m-3-3L6.696 17.304a1 1 0 0 0-.263.465L5.5 21.5l3.731-.933a1 1 0 0 0 .465-.263L20.5 9.5m-3-3l3 3"/>`
	penTabletInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M22 5v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2ZM2 12h4m0-9v18"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.5 11.5l-3.5 3m5-4.49l.01-.011"/></g>`
	penTabletConnectUsbInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 7V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2"/><path d="M2 12h4m0-9v18"/><path stroke-linecap="round" stroke-linejoin="round" d="M19.25 12H11m7.7 0l-.825 3h-1.65m1.375-3l-1.1-3h-1.925M22 12a1.37 1.37 0 0 0-1.375-1.364c-.76 0-1.375.61-1.375 1.364a1.37 1.37 0 0 0 1.375 1.364c.76 0 1.375-.61 1.375-1.364Z"/></g>`
	penTabletConnectWifiInnerSVG              = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m17 15.51l.01-.011M12 12c2.5-3 7.5-3 10 0m-8 2c1.5-2 4.5-2 6 0"/><path stroke-linecap="round" stroke-linejoin="round" d="M22 7V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2"/><path d="M2 12h4m0-9v18"/></g>`
	pentagonInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.647 2.256a.6.6 0 0 1 .706 0l9.756 7.089a.6.6 0 0 1 .218.67L18.6 21.485a.6.6 0 0 1-.57.414H5.97a.6.6 0 0 1-.57-.414l-3.727-11.47a.6.6 0 0 1 .218-.67l9.756-7.089Z"/>`
	peopleTagInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 16V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 10a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`
	percentageInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 19a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM7 9a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm12-4L5 19"/>`
	percentageCircleInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M15.5 16a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-7-7a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/><path d="m16 8l-8 8"/></g>`
	percentageSquareInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 16a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-7-7a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m16 8l-8 8"/></g>`
	perspectiveViewInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1 21L4.143 3h15.714L23 21H1Zm1-4.5h20M3 12h18M4 7.5h16M12 3v18M8 3.5l-1.5 17m9.5-17l1.5 17"/>`
	pharmacyCrossCircleInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M13.9 18h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 1-.6-.6v-3.8a.6.6 0 0 1 .6-.6h2.3a.6.6 0 0 0 .6-.6V6.6a.6.6 0 0 1 .6-.6h3.8a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 0 .6.6h2.3a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6h-2.3a.6.6 0 0 0-.6.6v2.3a.6.6 0 0 1-.6.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	pharmacyCrossSquareInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5Z"/><path d="M13.9 18h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 1-.6-.6v-3.8a.6.6 0 0 1 .6-.6h2.3a.6.6 0 0 0 .6-.6V6.6a.6.6 0 0 1 .6-.6h3.8a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 0 .6.6h2.3a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6h-2.3a.6.6 0 0 0-.6.6v2.3a.6.6 0 0 1-.6.6Z"/></g>`
	phoneInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.118 14.702L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584l-3.882-1.479Z"/>`
	phoneAddInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.243 5.243h3m3 0h-3m0 0v-3m0 3v3m-1.125 6.459L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584l-3.882-1.479Z"/>`
	phoneDeleteInnerSVG                       = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirPhoneDelete0)"><path d="m17.121 7.364l2.122-2.121m2.121-2.122l-2.121 2.122m0 0L17.12 3.12m2.122 2.122l2.121 2.121m-3.245 7.339L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584l-3.882-1.479Z"/></g><defs><clipPath id="iconoirPhoneDelete0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	phoneDisabledInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.78 8.5l.49-2.63L7.815 2H4.064c-1.128 0-2.016.93-1.848 2.046c.288 1.902.957 4.861 2.51 7.7M10.94 13.5c.837.744 1.847 1.392 3.059 2l4.118-.798L22 16.182v3.584c0 1.192-1.032 2.1-2.197 1.847c-2.83-.616-7.83-2.107-11.58-5.432M21 3L3 21"/>`
	phoneIncomeInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 5h-6m0 0l3-3m-3 3l3 3m-.882 6.702L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584l-3.882-1.479Z"/>`
	phoneOutcomeInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 5h6m0 0l-3-3m3 3l-3 3m-.882 6.702L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584l-3.882-1.479Z"/>`
	phonePausedInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2v5m4-5v5m-3.882 7.702L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584l-3.882-1.479Z"/>`
	phoneRemoveInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.242 5.243h6m-4.124 9.459L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584l-3.882-1.479Z"/>`
	piggyBankInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M14.5 8.5c-.78-.202-1.866-.5-2.735-.5C7.476 8 4 10.668 4 13.958c0 1.891 1.148 3.577 2.938 4.668l-.485 1.6a.6.6 0 0 0 .574.774h1.764a.6.6 0 0 0 .36-.12l1.395-1.047h2.437l1.395 1.047a.6.6 0 0 0 .36.12h1.764a.6.6 0 0 0 .574-.774l-.485-1.6c1.067-.65 1.905-1.511 2.409-2.501M14.5 8.5L19 7l-.084 3.628L21 11.5V15l-1.926 1"/><path fill="currentColor" stroke-linecap="round" d="M15.5 13a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path stroke-linecap="round" d="M2 10s0 2.4 2 3"/><path d="M12.8 7.753c.13-.372.2-.772.2-1.188C13 4.596 11.433 3 9.5 3S6 4.596 6 6.565c0 .941.358 1.798.944 2.435"/></g>`
	pillowInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21.04 12.283l.599 4.19a2 2 0 0 1-2.179 2.273l-7.26-.726a2.005 2.005 0 0 0-.398 0l-7.261.726a2 2 0 0 1-2.179-2.273l.599-4.19a2 2 0 0 0 0-.566l-.599-4.19A2 2 0 0 1 4.54 5.254l7.261.726a2 2 0 0 0 .398 0l7.261-.726a2 2 0 0 1 2.179 2.273l-.599 4.19a2 2 0 0 0 0 .566ZM21 6l-4 3M7 15l-4 3"/>`
	pinInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 14.5L3 21M5 9.485l9.193 9.193l1.697-1.697l-.393-3.787l5.51-4.673l-5.85-5.85l-4.674 5.51l-3.786-.393L5 9.485Z"/>`
	pinAltInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 10c0 4.418-8 12-8 12s-8-7.582-8-12a8 8 0 1 1 16 0Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	pineTreeInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 2L7 6.643S10.042 7 12 7c1.958 0 5-.357 5-.357L12 2ZM8.5 7L5 10.94S7.625 12 12 12s7-1.06 7-1.06L15.5 7"/><path d="M6.5 11.5L3 15.523S5.7 18 12 18s9-2.477 9-2.477L17.5 11.5M12 22v-3"/></g>`
	pinterestInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 14.5c-3-4.5 1.462-8 4.5-8c3.038 0 5.5 1.654 5.5 5.5c0 3.038-2 5-4 5s-3-2-2.5-5m.5-2L9 21.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	pizzaSliceInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m14 9.01l.01-.011M8 8.01l.01-.011M8 14.01l.01-.011"/><path d="M6 19L2.236 3.004a.6.6 0 0 1 .754-.713L19 7"/><path stroke-linecap="round" d="M22.198 8.425a1.75 1.75 0 0 0-3.396-.85c-.391 1.568-1.9 4.05-4.227 6.375c-2.3 2.301-5.148 4.194-7.968 4.845a1.75 1.75 0 1 0 .787 3.41c3.68-.849 7.082-3.206 9.656-5.78c2.549-2.549 4.54-5.568 5.148-8Z"/></g>`
	planetInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="8"/><path d="M17.5 6.348c2.297-.538 3.945-.476 4.338.312c.73 1.466-3.158 4.89-8.687 7.645c-5.528 2.757-10.602 3.802-11.333 2.336c-.392-.786.544-2.134 2.349-3.64"/></g>`
	planetAltInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="8"/><path d="M19.812 12.99c1.813 1.51 2.755 2.864 2.362 3.651c-.731 1.467-5.805.42-11.333-2.336C5.312 11.55 1.423 8.126 2.154 6.66c.392-.786 2.033-.85 4.322-.315"/></g>`
	planetSatInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="8"/><path d="M17.5 6.348c2.297-.538 3.945-.476 4.338.312c.73 1.466-3.158 4.89-8.687 7.645c-5.528 2.757-10.602 3.802-11.333 2.336c-.392-.786.544-2.134 2.349-3.64"/><path stroke-linecap="round" stroke-linejoin="round" d="m9.5 10.51l.01-.011"/></g>`
	playInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.906 4.537A.6.6 0 0 0 6 5.053v13.894a.6.6 0 0 0 .906.516l11.723-6.947a.6.6 0 0 0 0-1.032L6.906 4.537Z"/>`
	playlistInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2 11h14M2 17h11M2 5h18"/><path d="M20 18.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0v-7.9a.6.6 0 0 1 .6-.6H22"/></g>`
	playlistAddInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 18h2m2 0h-2m0 0v-2m0 2v2M2 11h18M2 17h12M2 5h18"/>`
	playlistPlayInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 17.5L18.5 20v-5l3.5 2.5ZM2 5h18M2 11h18M2 17h12"/>`
	playstationGamepadInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 17.5c2.5 3.5 6.449.915 5.5-2.5c-1.425-5.129-2.2-7.984-2.603-9.492A2.032 2.032 0 0 0 18.438 4H5.562c-.918 0-1.718.625-1.941 1.515C2.78 8.863 2.033 11.802 1.144 15c-.948 3.415 3 6 5.5 2.5"/><path d="M16 4v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V4m0 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm8 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	plugTypeAInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10ZM9 10v4m6-4v4"/>`
	plugTypeCInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M8 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm8 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	plugTypeGInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0-15v3m2 4h3M7 14h3"/>`
	plugTypeLInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 3H5.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H10m1-14h2m-2 5h2m-2 5h2m1-14h4.4a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H14"/>`
	plusInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h6m6 0h-6m0 0V6m0 6v6"/>`
	pngFormatInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.5 15v-3m0 0V9h3v3h-3Zm6 3V9l3 6V9m6 0h-3v6h3v-2.4"/><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/></g>`
	pocketInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 6v5a9 9 0 1 1-18 0V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2Z"/><path d="m8 10l4 4l4-4"/></g>`
	podcastInnerSVG                           = `<defs><path id="iconoirPodcast0" d="M6 19a9.985 9.985 0 0 1-4-8C2 5.477 6.477 1 12 1s10 4.477 10 10a9.985 9.985 0 0 1-4 8"/></defs><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><use href="#iconoirPodcast0"/><use href="#iconoirPodcast0"/><path d="M7.528 15a6 6 0 1 1 8.944 0"/><path d="M12 13a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-1.924 3.283l.815-.543a2 2 0 0 1 2.218 0l.815.543a2 2 0 0 1 .863 1.993l-.509 3.053A2 2 0 0 1 12.306 23h-.612a2 2 0 0 1-1.973-1.671l-.508-3.053a2 2 0 0 1 .863-1.993Z"/></g>`
	pokeballInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M2 12h7m6 0h7"/></g>`
	positionInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0 0v2m-7-9H3m9-7V3m7 9h2"/>`
	positionAlignInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 16.01l.01-.011M4 20.01l.01-.011M4 8.01l.01-.011M4 4.01l.01-.011M4 12.01l.01-.011m7.99.011l.01-.011M8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M20 4.01l.01-.011M16 4.01l.01-.011M12 4.01l.01-.011M8 4.01l.01-.011"/>`
	potionInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 4h4v2.568c0 .258.17.487.412.579C22.938 10.37 20.907 22 15 22H9c-5.907 0-7.937-11.63.588-14.853a.629.629 0 0 0 .412-.58V4Z"/><path d="M6 10h12"/><path stroke-linecap="round" d="M9 2h6"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.667 13L10 16h4l-1.667 3"/></g>`
	poundInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.667 13.6c-1.111 2.667-2.778 5.333-5 6.4h10.555S17.89 20 19 18.933M15.111 13.6H4m13.333-4.8c0-2.651-2.238-4.8-5-4.8c-2.761 0-5 2.149-5 4.8s2.239 4.8 5 4.8"/>`
	precisionToolInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v2m0 8v2m-4-6H6m12 0h-2m-4 10c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	presentationInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 4.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6ZM8.5 21.5L12 18l3.5 3.5M12 2v2m-3 8v2m3-4v4m3-6v6"/>`
	printerInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 17h18M6 10V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6V10m3 10.4V14a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v6.4a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17 13.01l.01-.011"/></g>`
	printerAltInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 13.01l.01-.011M7 17h10M6 10V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6V10m3 10.4V14a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v6.4a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6Z"/>`
	printingPageInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17.571 18H20.4a.6.6 0 0 0 .6-.6V11a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v6.4a.6.6 0 0 0 .6.6h2.829M8 7V3.6a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6V7"/><path d="M6.098 20.315L6.428 18l.498-3.485A.6.6 0 0 1 7.52 14h8.96a.6.6 0 0 1 .594.515L17.57 18l.331 2.315a.6.6 0 0 1-.594.685H6.692a.6.6 0 0 1-.594-.685Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17 10.01l.01-.011"/></g>`
	priorityDownInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848L11.576 1.424ZM12 16l4-4m-4 4l-4-4.167M12 16V7"/>`
	priorityUpInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848L11.576 1.424ZM12 7l4 4m-4-4l-4 4.167M12 7v9"/>`
	privateWifiInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 18.51l.01-.011M2 7c6-4.5 14-4.5 20 0M5 11c4-3 10-3 14 0M8.5 14.5c2.25-1.4 4.75-1.4 7 0m5.667 4h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/>`
	profileCircleInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Z"/><path d="M4.271 18.346S6.5 15.5 12 15.5s7.73 2.846 7.73 2.846M12 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`
	prohibitionInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.141 5A9.97 9.97 0 0 0 12 2C6.477 2 2 6.477 2 12a9.968 9.968 0 0 0 2.859 7M19.14 5A9.967 9.967 0 0 1 22 12c0 5.523-4.477 10-10 10a9.97 9.97 0 0 1-7.141-3M19.14 5L4.86 19"/>`
	puzzleInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 14v4.4a.6.6 0 0 0 .6.6H10m9-5v4.4a.6.6 0 0 1-.6.6H14m0-14h4.4a.6.6 0 0 1 .6.6V10M4 10V5.6a.6.6 0 0 1 .6-.6H10m4 14v1a2 2 0 1 1-4 0v-1m-6-9h1a2 2 0 1 1 0 4H4m15-4h1a2 2 0 1 1 0 4h-1m-5-9V4a2 2 0 1 0-4 0v1"/>`
	qrCodeInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 12v3M12 3v3m6 6v3m-6 3h9m-3 3h3M6 12h3M6 6.011L6.01 6M12 12.011l.01-.011M3 12.011L3.01 12M12 9.011L12.01 9M12 15.011l.01-.011M15 21.011l.01-.011m-3.01.011l.01-.011M21 12.011l.01-.011M21 15.011l.01-.011M18 6.011L18.01 6M9 3.6v4.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6Zm12 0v4.8a.6.6 0 0 1-.6.6h-4.8a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6ZM6 18.011L6.01 18M9 15.6v4.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6v-4.8a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6Z"/>`
	questionMarkInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.9 8.08c0-4.773 7.5-4.773 7.5 0c0 3.409-3.409 2.727-3.409 6.818M12 19.01l.01-.011"/>`
	quoteInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M10 12H5a1 1 0 0 1-1-1V7.5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1V12Zm0 0c0 2.5-1 4-4 5.5M20 12h-5a1 1 0 0 1-1-1V7.5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1V12Zm0 0c0 2.5-1 4-4 5.5"/>`
	quoteMessageInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.29V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.961a2 2 0 0 0-1.561.75l-2.331 2.914A.6.6 0 0 1 3 20.29Z"/><path stroke-linecap="round" d="M10.5 10h-2a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v2Zm0 0c0 1-1 2-2 3m8-3h-2a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v2Zm0 0c0 1-1 2-2 3"/></g>`
	radiationInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20.662a9.955 9.955 0 0 1-5 1.337a9.954 9.954 0 0 1-5-1.337L10 16s1 .4 2 .4s2-.4 2-.4l3 4.662Zm-.002-17.323A9.954 9.954 0 0 1 20.656 7a9.954 9.954 0 0 1 1.342 5l-5.537-.268s-.154-1.066-.654-1.932c-.5-.866-1.346-1.532-1.346-1.532l2.537-4.93ZM1.998 12A9.954 9.954 0 0 1 3.34 7a9.954 9.954 0 0 1 3.658-3.66l2.537 4.928S8.69 8.934 8.19 9.8s-.654 1.932-.654 1.932L1.998 12ZM12 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`
	radiusInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm7-10l-3-3m3 3l-3 3m3-3h-7"/><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	rainInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14v2m0 4v2m-4-4v2m8-2v2m4-2.393c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`
	rawFormatInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M16.5 9v6l1.5-3l1.5 3V9m-9 6v-3m0 0v-1.5A1.5 1.5 0 0 1 12 9v0a1.5 1.5 0 0 1 1.5 1.5V12m-3 0h3m0 0v3m-9 0V9h2.4a.6.6 0 0 1 .6.6v.9A1.5 1.5 0 0 1 6 12v0"/><path stroke-linejoin="round" d="M4.5 12H6v0a1.5 1.5 0 0 1 1.5 1.5V15"/></g>`
	receiveDollarsInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 8.23c-.8-.737-2.207-1.25-3.5-1.282M3 15.23c.752.925 2.15 1.453 3.5 1.498m0-9.781c-1.539-.038-2.917.604-2.917 2.36c0 3.23 6.417 1.615 6.417 4.846c0 1.842-1.708 2.634-3.5 2.575m0-9.781V5m0 11.729V19M21 12h-8m0 0l3.84-4M13 12l3.84 4"/>`
	receiveEurosInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12h-8m0 0l3.84-4M13 12l3.84 4M11 7.503A4.746 4.746 0 0 0 8.87 7C6.18 7 4 9.239 4 12s2.18 5 4.87 5a4.73 4.73 0 0 0 2.13-.503M3 11h6m-6 2h6"/>`
	receivePoundsInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12h-8m0 0l3.84-4M13 12l3.84 4M7 13c-.667 1.667-1.667 3.333-3 4h6.333s1 0 1.667-.667M9.667 13H3m8-3a3 3 0 1 0-3 3"/>`
	receiveYensInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12h-8m0 0l3.84-4M13 12l3.84 4M3 13h8M3 7l4 5.5M11 7l-4 5.5m0 0V18m-4-3h8"/>`
	redoInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 9.5H9c-.162 0-4 0-4 4C5 18 8.702 18 9 18h8"/><path d="M15.5 13L19 9.5L15.5 6"/></g>`
	redoActionInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 7v5m-3-2.5H9C3.5 9.5 3.5 18 9 18h10"/><path d="M12.5 13L16 9.5L12.5 6"/></g>`
	redoCircleInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 10.625H9.8s0 0 0 0s-2.8 0-2.8 3C7 17 9.8 17 9.8 17h.8"/><path d="m13.5 14l3.5-3.375L13.5 7"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	reduceInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 20l5-5m0 0v4m0-4H5M20 4l-5 5m0 0V5m0 4h4"/>`
	reduceRoundArrowInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 9.5L9.5 12L7 14.5m9.5-5L14 12l2.5 2.5"/><path d="M6 5h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V9a4 4 0 0 1 4-4Z"/></g>`
	refreshInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.888 13.5C21.164 18.311 17.013 22 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2c4.1 0 7.625 2.468 9.168 6"/><path d="M17 8h4.4a.6.6 0 0 0 .6-.6V3"/></g>`
	refreshCircularInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.583 9.667C15.81 8.097 14.043 7 11.988 7C9.388 7 7.25 8.754 7 11"/><path stroke-linecap="round" stroke-linejoin="round" d="M14.494 9.722H16.4a.6.6 0 0 0 .6-.6V7.5m-9.583 6.167C8.191 15.629 9.957 17 12.012 17c2.6 0 4.736-2.193 4.988-5"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.506 13.622H7.6a.6.6 0 0 0-.6.6V16.4"/></g>`
	refreshDoubleInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.168 8A10.003 10.003 0 0 0 12 2c-5.185 0-9.45 3.947-9.95 9"/><path d="M17 8h4.4a.6.6 0 0 0 .6-.6V3M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path d="M7.05 16h-4.4a.6.6 0 0 0-.6.6V21"/></g>`
	reloadWindowInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M11 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v7"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011m10.656 11.668C21.049 15.097 19.636 14 17.99 14c-1.758 0-3.252 1.255-3.793 3"/><path stroke-linejoin="round" d="M19.995 16.772H21.4a.6.6 0 0 0 .6-.6V14.55m-7.666 4.783C14.953 20.903 16.366 22 18.01 22c1.758 0 3.252-1.255 3.793-3"/><path stroke-linejoin="round" d="M16.005 19.228H14.6a.6.6 0 0 0-.6.6v1.622"/></g>`
	reminderHandGestureInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m17.5 12l2.004 2.672a2 2 0 0 1-.126 2.552l-3.784 4.127a1.998 1.998 0 0 1-1.473.649H9.5c-2.358 0-3.622-2.575-3.982-3.93a.55.55 0 0 1-.018-.143V9.43c0-2.286 3-2.286 3 0V10"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.5 10V8.286c0-2.286-3-2.286-3 0V10m6 0V7.5c0-2.286-3-2.286-3 0c0 0 0 0 0 0V10m3 0V3.499A1.5 1.5 0 0 1 16 2v0a1.5 1.5 0 0 1 1.5 1.5V15"/><path d="M17.563 6.5h2.062C20.5 6.5 21 6.078 21 5.25C21 4.422 20.5 4 19.625 4H12.25C11.56 4 11 4.56 11 5.25v.25a1 1 0 0 0 1 1"/></g>`
	removeDatabaseScriptInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 14V8.5M6 13V6a3 3 0 0 1 3-3h5m4 1h4M12 21H6a4 4 0 0 1 0-8h12a4 4 0 1 0 4 4v-3"/>`
	removeFolderInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 6h4m-.6 14H2.6a.6.6 0 0 1-.6-.6V11h19.4a.6.6 0 0 1 .6.6v7.8a.6.6 0 0 1-.6.6ZM2 11V4.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H14"/>`
	removeFrameInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-miterlimit="1.5" d="M4.998 2H2v2.998h2.998V2Zm.001 1.5h14M3.5 4.998V19M20.498 5v14.002M4.999 20.5h14M4.998 19H2v2.998h2.998V19ZM21.997 2.001H19v2.998h2.998V2.001Zm0 17H19v2.998h2.998v-2.998Z"/><path d="M9 12h6"/></g>`
	removeFromCartInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h19l-3 10H6L3 6Zm0 0l-.75-2.5M9.992 11h4M11 19.5a1.5 1.5 0 0 1-3 0m9 0a1.5 1.5 0 0 1-3 0"/>`
	removeKeyframeInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 5h6m-6.152 8.317l-4.343 4.963a2 2 0 0 1-3.01 0l-4.343-4.963a2 2 0 0 1 0-2.634L7.495 5.72a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634Z"/>`
	removeKeyframeAltInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18.819 13.329l-5.324 5.99a2 2 0 0 1-2.99 0l-5.324-5.99a2 2 0 0 1 0-2.658l5.324-5.99a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.658ZM9 12h6"/>`
	removeKeyframesInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12h6M6.25 6l.245-.28a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634L9.505 18.28a2 2 0 0 1-3.01 0L6.25 18"/><path d="m13 19l4.884-5.698a2 2 0 0 0 0-2.604L13 5"/><path d="m17 19l4.884-5.698a2 2 0 0 0 0-2.604L17 5"/></g>`
	removeLinkInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.143 16.995c-.393 0-.775-.043-1.143-.123c-2.29-.506-4-2.496-4-4.874c0-2.714 2.226-4.923 5-4.996m6.318 2.632A5.517 5.517 0 0 0 11 7.5"/><path d="M16.857 7c.393 0 .775.043 1.143.124c2.29.505 4 2.495 4 4.874c0 2.76-2.302 4.997-5.143 4.997h-1.714c-2.826 0-5.143-2.506-5.143-4.997c0 0 0-.998.5-1.498M3 3l18 18"/></g>`
	removeMediaImageInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 16l7-3l4 1.818M16 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm.879 11.121L19 19m2.121-2.121L19 19m0 0l-2.121-2.121M19 19l2.121 2.121"/><path d="M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13"/></g>`
	removeMediaVideoInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.879 21.121L19 19m2.121-2.121L19 19m0 0l-2.121-2.121M19 19l2.121 2.121M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13"/><path d="M9.898 8.513a.6.6 0 0 0-.898.52v5.933a.6.6 0 0 0 .898.521l5.19-2.966a.6.6 0 0 0 0-1.042l-5.19-2.966Z"/></g>`
	removePageInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 12h6M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`
	removePageAltInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 12V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H11"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20M1.992 19h6"/></g>`
	removePinInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 14.5L3 21M7.676 7.89l-.979-.102L5 9.485l9.193 9.193l1.697-1.697l-.102-.981m-4.303-9l3.672-4.329l5.85 5.85l-4.308 3.654M3 3l18 18"/>`
	removePinAltInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 9.2C16 13.177 9 20 9 20S2 13.177 2 9.2C2 5.224 5.134 2 9 2s7 3.224 7 7.2Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.88 21.121L19 19m2.122-2.121L19 19m0 0l-2.12-2.121M19 19l2.122 2.121"/></g>`
	removeSelectionInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4H4v3m4 5h8M4 11v2m7-9h2m-2 16h2m7-9v2m-3-9h3v3M7 20H4v-3m13 3h3v-3"/>`
	removeSquareInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.879 14.121L12 12m2.121-2.121L12 12m0 0L9.879 9.879M12 12l2.121 2.121M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	removeUserInnerSVG                        = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirRemoveUser0)"><path d="M18.621 12.121L20.743 10m2.121-2.121L20.743 10m0 0L18.62 7.879M20.743 10l2.121 2.121M1 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1m-7-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g><defs><clipPath id="iconoirRemoveUser0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	repeatInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 17H8c-1.667 0-5-1-5-5s3.333-5 5-5h8c1.667 0 5 1 5 5c0 1.494-.465 2.57-1.135 3.331"/><path d="M14.5 14.5L17 17l-2.5 2.5"/></g>`
	repeatOnceInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 17H8c-1.667 0-5-1-5-5m5-5h8c1.667 0 5 1 5 5c0 1.494-.465 2.57-1.135 3.331"/><path d="M14.5 14.5L17 17l-2.5 2.5M4 8V3L2 4"/></g>`
	replyInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 10h14c8 0 8 11 0 11M2 10l7-7m-7 7l7 7"/>`
	replyToMessageInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 8l5 3l5-3"/><path d="M10 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v6.857"/><path stroke-linejoin="round" d="M13 17.111h6.3c3.6 0 3.6 4.889 0 4.889M13 17.111L16.15 14M13 17.111l3.15 3.111"/></g>`
	reportColumnsInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 7.4V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm11 13v-3.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm0-8V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v8.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 8v-8.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v8.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/>`
	reportsInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M9 21h6m-6 0v-5m0 5H3.6a.6.6 0 0 1-.6-.6v-3.8a.6.6 0 0 1 .6-.6H9m6 5V9m0 12h5.4a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6h-4.8a.6.6 0 0 0-.6.6V9m0 0H9.6a.6.6 0 0 0-.6.6V16"/>`
	repositoryInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M15 17v5l2.5-1.6L20 22v-5"/><path d="M6 17h14"/><path stroke-linejoin="round" d="M6 17a2 2 0 1 0 0 4h5.5"/></g>`
	restartInnerSVG                           = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirRestart0)"><path d="M6.677 20.567C2.531 18.021.758 12.758 2.717 8.144C4.875 3.06 10.745.688 15.829 2.846c5.084 2.158 7.456 8.029 5.298 13.113a9.954 9.954 0 0 1-3.962 4.608"/><path d="M17 16v4.4a.6.6 0 0 0 .6.6H22m-10 1.01l.01-.011"/></g><defs><clipPath id="iconoirRestart0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	rewindInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.044 5.704a.6.6 0 0 1 .956.483v11.626a.6.6 0 0 1-.956.483l-7.889-5.813a.6.6 0 0 1 0-.966l7.89-5.813Zm-11 0a.6.6 0 0 1 .956.483v11.626a.6.6 0 0 1-.956.483l-7.888-5.813a.6.6 0 0 1 0-.966l7.888-5.813Z"/>`
	rhombusInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848L11.576 1.424Z"/>`
	rightRoundArrowInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6.75 12h10m0 0L14 14.75M16.75 12L14 9.25"/><path d="M2 15V9a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4Z"/></g>`
	ringsInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 8a6 6 0 1 0 0 12A6 6 0 0 0 8 8Zm0 0V3"/><path d="M16 8a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm0 0V3"/></g>`
	rocketInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.061 10.404L14 17h-4l-2.061-6.596a6 6 0 0 1 .998-5.484l2.59-3.315a.6.6 0 0 1 .946 0l2.59 3.315a6 6 0 0 1 .998 5.484ZM10 20c0 2 2 3 2 3s2-1 2-3m-5.5-7.5C5 15 7 19 7 19l3-2m5.931-4.5c3.5 2.5 1.5 6.5 1.5 6.5l-3-2"/><path d="M12 11a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g>`
	rookInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 16h10m-8-5h6m-5-7v2m4-2v2m3.4 3H6.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6Zm.501 12H6.099a.615.615 0 0 1-.521-.932C6.792 18.06 9.5 13.328 9.5 11V9.6a.6.6 0 0 1 .6-.6h3.8a.6.6 0 0 1 .6.6V11c0 2.327 2.708 7.061 3.922 9.068a.615.615 0 0 1-.521.932Z"/>`
	rotateCameraLeftInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.05 3v4.497c0 .278.226.503.504.503v0c.2 0 .38-.119.466-.3A10.001 10.001 0 0 1 12.05 2c5.186 0 9.45 3.947 9.951 9m0 10v-4.497a.503.503 0 0 0-.503-.503v0a.52.52 0 0 0-.465.3A10.001 10.001 0 0 1 12 22c-5.185 0-9.448-3.947-9.95-9"/><path d="M6 16.4V9.394a.6.6 0 0 1 .6-.6h1.173a.6.6 0 0 0 .504-.275l1.446-2.244A.6.6 0 0 1 10.227 6h3.546a.6.6 0 0 1 .504.275l1.446 2.244a.6.6 0 0 0 .504.275H17.4a.6.6 0 0 1 .6.6V16.4a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6Z"/><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	rotateCameraRightInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22.003 3v4.497A.503.503 0 0 1 21.5 8v0c-.2 0-.38-.119-.466-.3A10.001 10.001 0 0 0 12.003 2c-5.186 0-9.45 3.947-9.95 9"/><path d="M6 16.4V9.394a.6.6 0 0 1 .6-.6h1.173a.6.6 0 0 0 .504-.275l1.446-2.244A.6.6 0 0 1 10.227 6h3.546a.6.6 0 0 1 .504.275l1.446 2.244a.6.6 0 0 0 .504.275H17.4a.6.6 0 0 1 .6.6V16.4a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6Z"/><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-9.95 7v-4.497c0-.278.226-.503.504-.503v0c.2 0 .38.119.466.3a10.001 10.001 0 0 0 9.03 5.7c5.186 0 9.45-3.947 9.951-9"/></g>`
	roundFlaskInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M19 15H5"/><path stroke-linecap="round" d="M16 4H8m7 .5v4.253c0 .763.445 1.445 1.078 1.871C18.287 12.11 20 14.617 20 17.462c0 .812-.114 1.596-.325 2.338c-.215.75-.945 1.2-1.726 1.2H6.051c-.78 0-1.511-.45-1.726-1.2A8.505 8.505 0 0 1 4 17.462c0-2.845 1.713-5.353 3.922-6.838C8.555 10.198 9 9.516 9 8.754V4.5m4 2.51l.01-.011M11 2.01l.01-.011"/></g>`
	roundedMirrorInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 10v4a8 8 0 1 1-16 0v-4a8 8 0 1 1 16 0Zm-2.5-5.5L13 8m6-1l-7.5 6"/>`
	rssFeedInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19c0-4.2-2.8-7-7-7m14 7c0-8.4-5.6-14-14-14m0 14.01l.01-.011"/>`
	rssFeedTagInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17c0-3-2-5-5-5m10 5c0-6-4-10-10-10m0 10.01l.01-.011"/><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5Z"/></g>`
	rubikCubeInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6ZM9 3v18M3 9h18M3 15h18M15 3v18"/>`
	rulerInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7V2.6a.6.6 0 0 0-.6-.6H8.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6h6.8a.6.6 0 0 0 .6-.6V17m0-10h-3m3 0v5m0 0h-3m3 0v5m0 0h-3"/>`
	rulerAddInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 5h3m3 0h-3m0 0V2m0 3v3m-7-1V2.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6h6.8a.6.6 0 0 0 .6-.6V17m0-10H8m3 0v5m0 0H8m3 0v5m0 0H8"/>`
	rulerCombineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 21.4V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H10.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6ZM16 10V7m-6 3V7m0 9H7m3-6H7"/>`
	rulerRemoveInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 5h6M11 7V2.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6h6.8a.6.6 0 0 0 .6-.6V17m0-10H8m3 0v5m0 0H8m3 0v5m0 0H8"/>`
	runningInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-2.387 1.267l-3.308 4.135l4.135 4.135l-2.067 4.55"/><path d="m6.41 9.508l3.387-3.309l2.816 2.068l2.895 3.308h3.722M8.892 15.71l-1.241.827H4.343"/></g>`
	safariInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.586 10.586L16.95 7.05l-3.536 6.364m-2.828-2.828L7.05 16.95l6.364-3.536m-2.828-2.828l2.828 2.828"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm7-10h-1M6 12H5m7-7v1m0 12v1M7.05 7.05l.707.707m8.486 8.486l.707.707"/></g>`
	sandalsInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 7s.5-4-4-4s-4 4-4 4m8 0h-8m8 0l-.214 3M14 7l.214 3m7.572 0l-.587 8.214A3 3 0 0 1 18.207 21h-.414a3 3 0 0 1-2.992-2.786L14.214 10m7.572 0h-7.572M10 7s.5-4-4-4s-4 4-4 4m8 0H2m8 0l-.214 3M2 7l.214 3m7.572 0l-.587 8.214A3 3 0 0 1 6.207 21h-.414a3 3 0 0 1-2.992-2.786L2.214 10m7.572 0H2.214"/>`
	saveActionFloppyInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 7.5V5a2 2 0 0 1 2-2h11.172a2 2 0 0 1 1.414.586l2.828 2.828A2 2 0 0 1 21 7.828V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2.5M6 21v-4"/><path d="M18 21v-7.4a.6.6 0 0 0-.6-.6H15m1-10v5.4a.6.6 0 0 1-.6.6h-1.9M8 3v3m-7 6h11m0 0L9 9m3 3l-3 3"/></g>`
	saveFloppyDiskInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h11.172a2 2 0 0 1 1.414.586l2.828 2.828A2 2 0 0 1 21 7.828V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path d="M8.6 9h6.8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6H8.6a.6.6 0 0 0-.6.6v4.8a.6.6 0 0 0 .6.6ZM6 13.6V21h12v-7.4a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 0-.6.6Z"/></g>`
	scaleFrameEnlargeInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 13.6V21H3.6a.6.6 0 0 1-.6-.6V13h7.4a.6.6 0 0 1 .6.6Zm0 7.4h3M3 13v-3m3-7H3.6a.6.6 0 0 0-.6.6V6m11-3h-4m11 7v4M18 3h2.4a.6.6 0 0 1 .6.6V6m-3 15h2.4a.6.6 0 0 0 .6-.6V18m-10-8h3v3"/>`
	scaleFrameReduceInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11 15v-1.4a.6.6 0 0 0-.6-.6H9m-3 0H3m8 5v3"/><path stroke-miterlimit="1.5" stroke-width="1.499" d="M20.4 3H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6Z"/><path d="M16 11h-3V8"/></g>`
	scanBarcodeInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12V6h1m-1 6h1V6m-1 12v-3h1m0 0v3h-1M7 6v6m0 3v3m7-12v6m0 3v3m3-12v6m0 3v3M6 3H3v3m-1 6h20m-4-9h3v3M6 21H3v-3m15 3h3v-3"/>`
	scanQrCodeInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6.6v1.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6V6.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6ZM6 12h3m6 0v3m-3 3h3m-3-5.989l.01-.011m5.99.011l.01-.011M12 15.011l.01-.011m5.99.011l.01-.011M18 18.011l.01-.011M12 9.011L12.01 9M12 6.011L12.01 6M9 15.6v1.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6v-1.8a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6Zm9-9v1.8a.6.6 0 0 1-.6.6h-1.8a.6.6 0 0 1-.6-.6V6.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6ZM18 3h3v3m-3 15h3v-3M6 3H3v3m3 15H3v-3"/>`
	scanningInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3H3v3m-1 6h20M9 19v-4m3 1v-1m3 2v-2m-3 6v-3m6-15h3v3M6 21H3v-3m15 3h3v-3"/>`
	scarfInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 11H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v16m-3 0v-2M15 3v18m0-14H3"/>`
	scissorInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.236 7a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4Zm0 0L20 18M7.236 17a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4Zm0 0L20 6"/>`
	scissorAltInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.236 8a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4Zm0 0L20 16m-9.764 0a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4Zm0 0L20 8"/>`
	screenshotInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 21.4v-7.006a.6.6 0 0 1 .6-.6h1.173a.6.6 0 0 0 .504-.275l1.446-2.244a.6.6 0 0 1 .504-.275h3.546a.6.6 0 0 1 .504.275l1.446 2.244a.6.6 0 0 0 .504.275H21.4a.6.6 0 0 1 .6.6V21.4a.6.6 0 0 1-.6.6H10.6a.6.6 0 0 1-.6-.6Z"/><path d="M16 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM3 18v3h2.5M3 9.5v5M3 6V3h3m3.5 0h5M18 3h3v2.5m0 4.5V8.5"/></g>`
	seaAndSunInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 15c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3M3 20c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3m-2-10a7 7 0 1 0-14 0"/>`
	seaWavesInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 10c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3M3 17c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3"/>`
	searchInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 17l4 4M3 11a8 8 0 1 0 16 0a8 8 0 0 0-16 0Z"/>`
	searchEngineInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.856 13.85a3.429 3.429 0 1 0-4.855-4.842a3.429 3.429 0 0 0 4.855 4.842Zm0 0L16 16"/></g>`
	searchFontInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 19.5L21 21m-7-4a3 3 0 1 0 6 0a3 3 0 0 0-6 0ZM9 5v12m0 0H7m2 0h2m4-10V5H3v2"/>`
	searchWindowInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011m9.114 15.12a3 3 0 1 0-4.248-4.237a3 3 0 0 0 4.248 4.237Zm0 0L22 22"/></g>`
	secureWindowInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v7"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011m7.982 9.126l2.556.649c.266.068.453.31.445.584C21.821 21.116 18.5 22 18.5 22s-3.321-.884-3.493-6.642a.588.588 0 0 1 .445-.584l2.556-.649c.323-.082.661-.082.984 0Z"/></g>`
	securityPassInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 11l3 3l8-8"/><path d="M20 12a8 8 0 1 1-5.3-7.533"/></g>`
	selectWindowInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M14 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/><path d="M22.082 18.365c.494.304.464 1.043-.045 1.1l-2.566.292l-1.152 2.312c-.228.458-.933.234-1.05-.334l-1.255-6.116c-.098-.48.333-.782.75-.525l5.318 3.271Z" clip-rule="evenodd"/></g>`
	selectionInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4H4v3m0 4v2m7-9h2m-2 16h2m7-9v2m-3-9h3v3M7 20H4v-3m13 3h3v-3"/>`
	selectiveToolInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	sendInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 12L3 20l3.563-8L3 4l19 8ZM6.5 12H22"/>`
	sendDiagonalInnerSVG                      = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirSendDiagonal0)"><path d="M22.152 3.553L11.178 21.004l-1.67-8.596L2 7.898l20.152-4.345ZM9.456 12.444l12.696-8.89"/></g><defs><clipPath id="iconoirSendDiagonal0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	sendDollarsInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 8.23c-.8-.737-2.207-1.25-3.5-1.282M3 15.23c.752.925 2.15 1.453 3.5 1.498m0-9.781c-1.539-.038-2.917.604-2.917 2.36c0 3.23 6.417 1.615 6.417 4.846c0 1.842-1.708 2.634-3.5 2.575m0-9.781V5m0 11.729V19m6.5-7h8m0 0l-3.84-4M21 12l-3.84 4"/>`
	sendEurosInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 12h8m0 0l-3.84-4M21 12l-3.84 4M11 7.503A4.746 4.746 0 0 0 8.87 7C6.18 7 4 9.239 4 12s2.18 5 4.87 5a4.73 4.73 0 0 0 2.13-.503M3 11h6m-6 2h6"/>`
	sendMailInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m9 9l4.5 3L18 9M3 13.5h2m-4-3h4"/><path d="M5 7.5V7a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-.5"/></g>`
	sendPoundsInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 12h8m0 0l-3.84-4M21 12l-3.84 4M7 13c-.667 1.667-1.667 3.333-3 4h6.333s1 0 1.667-.667M9.667 13H3m8-3a3 3 0 1 0-3 3"/>`
	sendYensInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 12h8m0 0l-3.84-4M21 12l-3.84 4M3 13h8M3 7l4 5.5M11 7l-4 5.5m0 0V18m-4-3h8"/>`
	serverInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m6 18.01l.01-.011M6 6.01l.01-.011"/><path d="M2 9.4V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Zm0 12v-6.8a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`
	serverConnectionInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 19h9m9 0h-9m0 0v-6m0 0h6V5H6v8h6ZM9 9.01l.01-.011M12 9.01l.01-.011"/>`
	settingsInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="m19.622 10.395l-1.097-2.65L20 6l-2-2l-1.735 1.483l-2.707-1.113L12.935 2h-1.954l-.632 2.401l-2.645 1.115L6 4L4 6l1.453 1.789l-1.08 2.657L2 11v2l2.401.655L5.516 16.3L4 18l2 2l1.791-1.46l2.606 1.072L11 22h2l.604-2.387l2.651-1.098C16.697 18.831 18 20 18 20l2-2l-1.484-1.75l1.098-2.652l2.386-.62V11l-2.378-.605Z"/></g>`
	settingsCloudInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linejoin="round" d="M12 8c-3.273 0-3.273 2-3.273 3C7.818 11 6 11.5 6 13.5S7.818 16 8.727 16h6.546c.909 0 2.727-.5 2.727-2.5S16.182 11 15.273 11c0-1 0-3-3.273-3Z"/></g>`
	settingsProfilesInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11.607 2.342a.6.6 0 0 1 .787 0l1.948 1.692a.6.6 0 0 0 .445.145l2.572-.224a.6.6 0 0 1 .636.463l.582 2.514a.6.6 0 0 0 .275.38l2.212 1.33a.6.6 0 0 1 .243.748l-1.008 2.376a.6.6 0 0 0 0 .468l1.008 2.376a.6.6 0 0 1-.243.749l-2.212 1.33a.6.6 0 0 0-.275.379l-.582 2.514a.6.6 0 0 1-.636.463l-2.572-.224a.6.6 0 0 0-.445.144l-1.949 1.693a.6.6 0 0 1-.787 0l-1.948-1.693a.6.6 0 0 0-.445-.144l-2.572.224a.6.6 0 0 1-.636-.463l-.582-2.514a.6.6 0 0 0-.275-.38l-2.212-1.33a.6.6 0 0 1-.243-.748l1.008-2.376a.6.6 0 0 0 0-.468L2.693 9.39a.6.6 0 0 1 .243-.749l2.212-1.33a.6.6 0 0 0 .275-.379l.582-2.514a.6.6 0 0 1 .636-.463l2.572.224a.6.6 0 0 0 .445-.145l1.949-1.692Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 13l2 2l5-5"/></g>`
	shareAndroidInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 22a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-14a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM6 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="m15.5 6.5l-7 4m0 3l7 4"/></g>`
	shareIosInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6m8 2V3m0 0L8.5 6.5M12 3l3.5 3.5"/>`
	shieldInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/>`
	shieldAddInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3m3 0h-3m0 0V9m0 3v3m-7 3L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/>`
	shieldAlertInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7v5m0 4.01l.01-.011M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/>`
	shieldAltInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.571 8l-.44-3.084A1 1 0 0 1 3.904 3.8l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8a1 1 0 0 1 .773 1.117L20.43 8M3.57 8h16.86M3.57 8c.309 2.16.69 4.822 1 7m15.86-7c-.309 2.16-.69 4.822-1 7m0 0L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18l-.429-3m14.858 0H4.57"/>`
	shieldBrokenInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.5 7L9 12h6l-2.5 5"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/></g>`
	shieldCheckInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.5 11.5l3 3l5-5"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/></g>`
	shieldCrossInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.871 14.121L11.993 12m2.121-2.121L11.993 12m0 0L9.87 9.879M11.993 12l2.121 2.121M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/>`
	shieldDownloadInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v7m0 0l3-3m-3 3l-3-3m-4 6L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/>`
	shieldEyeInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 9s1-1 4-1s4 1 4 1"/><path fill="currentColor" d="M12 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/></g>`
	shieldLoadingInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 10.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/>`
	shieldMinusInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/>`
	shieldQuestionInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 9c0-3.5 5.5-3.5 5.5 0c0 2.5-2.5 2-2.5 5m0 4.01l.01-.011"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/></g>`
	shieldSearchInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.5 13l1.5 1.5M9 11a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Z"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/></g>`
	shieldUploadInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15V8m0 0l3 3m-3-3l-3 3m-4 7L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18Z"/>`
	shopInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 9v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9"/><path d="M20.485 3h-3.992l.5 5s1 1 2.5 1a3.23 3.23 0 0 0 2.139-.806a.503.503 0 0 0 .15-.465L21.076 3.5a.6.6 0 0 0-.591-.5Z"/><path d="m16.493 3l.5 5s-1 1-2.5 1s-2.5-1-2.5-1V3h4.5Z"/><path d="M11.993 3v5s-1 1-2.5 1s-2.5-1-2.5-1l.5-5h4.5Z"/><path d="M7.493 3H3.502a.6.6 0 0 0-.592.501L2.205 7.73a.504.504 0 0 0 .15.465c.328.29 1.061.806 2.138.806c1.5 0 2.5-1 2.5-1l.5-5Z"/></g>`
	shopAltInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20.485 3h-3.992l.5 5s1 1 2.5 1a3.23 3.23 0 0 0 2.139-.806a.503.503 0 0 0 .15-.465L21.076 3.5a.6.6 0 0 0-.591-.5Z"/><path d="m16.493 3l.5 5s-1 1-2.5 1s-2.5-1-2.5-1V3h4.5Z"/><path d="M11.993 3v5s-1 1-2.5 1s-2.5-1-2.5-1l.5-5h4.5Z"/><path d="M7.493 3H3.502a.6.6 0 0 0-.592.501L2.205 7.73a.504.504 0 0 0 .15.465c.328.29 1.061.806 2.138.806c1.5 0 2.5-1 2.5-1l.5-5Z"/><path d="M3 9v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9"/><path stroke-miterlimit="16" d="M14.833 21v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v6"/></g>`
	shoppingBagInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.26 9.696l1.385 9A2 2 0 0 1 18.67 21H5.33a2 2 0 0 1-1.977-2.304l1.385-9A2 2 0 0 1 6.716 8h10.568a2 2 0 0 1 1.977 1.696ZM14 5a2 2 0 1 0-4 0"/>`
	shoppingBagAddInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.26 9.696l1.385 9A2 2 0 0 1 18.67 21H5.33a2 2 0 0 1-1.977-2.304l1.385-9A2 2 0 0 1 6.716 8h10.568a2 2 0 0 1 1.977 1.696ZM14 5a2 2 0 1 0-4 0M8.992 15h3m3 0h-3m0 0v-3m0 3v3"/>`
	shoppingBagAltInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.26 9.696l1.385 9A2 2 0 0 1 18.67 21H5.33a2 2 0 0 1-1.977-2.304l1.385-9A2 2 0 0 1 6.716 8h10.568a2 2 0 0 1 1.977 1.696ZM9 11v7m6-7v7M14 5a2 2 0 1 0-4 0"/>`
	shoppingBagArrowDownInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 21h2.169a2 2 0 0 0 1.976-2.304l-1.384-9A2 2 0 0 0 17.284 8H6.716a2 2 0 0 0-1.977 1.696l-1.385 9A2 2 0 0 0 5.331 21H7.5m4.5-9v7m0 0l3-3m-3 3l-3-3m5-11a2 2 0 1 0-4 0"/>`
	shoppingBagArrowUpInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 21h2.169a2 2 0 0 0 1.976-2.304l-1.384-9A2 2 0 0 0 17.284 8H6.716a2 2 0 0 0-1.977 1.696l-1.385 9A2 2 0 0 0 5.331 21H7.5m4.5-2v-7m0 0l3 3m-3-3l-3 3m5-10a2 2 0 1 0-4 0"/>`
	shoppingBagCheckInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20 14.5l-.74-4.804A2 2 0 0 0 17.285 8H6.716a2 2 0 0 0-1.977 1.696l-1.385 9A2 2 0 0 0 5.331 21H12"/><path d="m14 19l3 3l5-5M14 5a2 2 0 1 0-4 0"/></g>`
	shoppingBagIssueInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 14.5l-.74-4.804A2 2 0 0 0 17.285 8H6.716a2 2 0 0 0-1.977 1.696l-1.385 9A2 2 0 0 0 5.331 21H12m5.5-4v2m0 3.01l.01-.011M14 5a2 2 0 1 0-4 0"/>`
	shoppingBagRemoveInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.26 9.696l1.385 9A2 2 0 0 1 18.67 21H5.33a2 2 0 0 1-1.977-2.304l1.385-9A2 2 0 0 1 6.716 8h10.568a2 2 0 0 1 1.977 1.696ZM14 5a2 2 0 1 0-4 0M8.992 15h6"/>`
	shoppingCodeInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 5v2m4-2v6m8-6v1M6 10v6m0 2.5v.5m4-.5v.5m4-.5v.5m4-.5v.5m-8-5v2m4-3v3m0-11v5m4-1v7"/>`
	shoppingCodeCheckInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 5v2m4-2v6m8-6v1M6 10v6m0 2.5v.5m4-.5v.5m0-5v2m4-3v2m0-10v5m4-1v4m-3 6l2 2l4-4"/>`
	shoppingCodeErrorInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 5v2m4-2v6m8-6v1M6 10v6m0 2.5v.5m4-.5v.5m0-5v2m4-3v2m0-10v5m4-1v4m-1.879 8.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L16.12 17.12m2.122 2.122l2.121 2.121"/>`
	shortPantsInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.06 5.655A.6.6 0 0 1 3.658 5h16.684a.6.6 0 0 1 .598.655l-1.176 12.8a.6.6 0 0 1-.597.545h-4.152a.6.6 0 0 1-.574-.424l-1.867-6.102c-.174-.566-.974-.566-1.148 0l-1.868 6.102a.6.6 0 0 1-.573.424H4.833a.6.6 0 0 1-.597-.545L3.643 12L3.06 5.655Z"/><path d="M4 9.5h1.5a2 2 0 0 0 2-2V5m13 4.5h-2a2 2 0 0 1-2-2V5"/></g>`
	shortPantsAltInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16.8h6.966a.6.6 0 0 0 .596-.53l1.36-11.6a.6.6 0 0 0-.596-.67H3.659a.6.6 0 0 0-.597.656l1.387 14.8a.6.6 0 0 0 .597.544H11.4a.6.6 0 0 0 .6-.6V12"/>`
	shortcutInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M15.025 8.025h-4.95m4.95 0v4.95m0-4.95l-3.535 3.536c-2.475 2.475 0 4.95 0 4.95"/></g>`
	shuffleInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 7c-3 0-8.5 0-10.5 5.5S5 18 2 18"/><path d="m20 5l2 2l-2 2m2 9c-3 0-8.5 0-10.5-5.5S5 7 2 7"/><path d="m20 20l2-2l-2-2"/></g>`
	sidebarCollapseInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2Z"/><path d="M7.25 10L5.5 12l1.75 2m2.25 7V3"/></g>`
	sidebarExpandInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2Zm-9.5 0V3"/><path d="m5.5 10l1.75 2l-1.75 2"/></g>`
	sigmaFunctionInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v3M4 20h16v-3M4 20l8-8l-8-8"/>`
	simpleCartInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h19l-3 10H6L3 6Zm0 0l-.75-2.5m8.75 16a1.5 1.5 0 0 1-3 0m9 0a1.5 1.5 0 0 1-3 0"/>`
	sineWaveInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-3.857 1.286-9 3.857-9c3.857 0 6.429 18 10.286 18C19.714 21 21 15.857 21 12M3 12h2m14 0h2m-5.5 0h1m-9 0h1"/>`
	singleTapGestureInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 20.5a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path d="M4 7.29C5.496 5.039 8.517 3.5 12 3.5c3.483 0 6.504 1.539 8 3.79"/></g>`
	skateboardInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 16a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM2 9l3.333 1h13.334L22 9m-4.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`
	skateboardingInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 19l2.333 1h9.334L19 19M8 22.01l.01-.011m7.99.011l.01-.011M7 7.833l3-1.5c2-1 4.27.568 4.27.568l-4.308 3.135L14 13.334v4m-4.451-3.989l-1.24.827H5M15.165 9.21h2.722M17 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`
	skipNextInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7v10M6.972 5.267A.6.6 0 0 0 6 5.738v12.524a.6.6 0 0 0 .972.47l7.931-6.261a.6.6 0 0 0 0-.942L6.972 5.267Z"/>`
	skipPrevInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 7v10M17.028 5.267a.6.6 0 0 1 .972.471v12.524a.6.6 0 0 1-.972.47l-7.931-6.261a.6.6 0 0 1 0-.942l7.931-6.262Z"/>`
	sleeperChairInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 18v3m1-11V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5"/><path d="M19.5 10a2.5 2.5 0 0 0-2.5 2.5V14H7v-1.5a2.5 2.5 0 1 0-3 2.45V18h16v-3.05a2.5 2.5 0 0 0-.5-4.95Zm.5 8v3"/></g>`
	smallLampInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.872 3.428l-2.64 8.8a.6.6 0 0 0 .574.772h14.388a.6.6 0 0 0 .574-.772l-2.64-8.8A.6.6 0 0 0 16.554 3H7.446a.6.6 0 0 0-.574.428ZM12 17v-2m-3.4 6h6.8c.331 0 .595-.268.542-.596C15.763 19.29 15.026 17 12 17c-3.026 0-3.763 2.29-3.942 3.404c-.053.328.21.596.542.596Z"/>`
	smallLampAltInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.872 3.428l-2.64 8.8a.6.6 0 0 0 .574.772h14.388a.6.6 0 0 0 .574-.772l-2.64-8.8A.6.6 0 0 0 16.554 3H7.446a.6.6 0 0 0-.574.428ZM8 15v-2m0 8h8m-4-6v6"/>`
	smallShopInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m21.818 9.364l-1.694-5.929A.6.6 0 0 0 19.547 3H15.5l.475 5.704a.578.578 0 0 0 .278.45c.39.233 1.152.663 1.747.846c1.016.313 2.5.2 3.346.096a.57.57 0 0 0 .472-.732Z"/><path d="M14 10c.568-.175 1.288-.574 1.69-.812a.578.578 0 0 0 .28-.549L15.5 3h-7l-.47 5.639a.578.578 0 0 0 .28.55c.402.237 1.122.636 1.69.811c1.493.46 2.507.46 4 0Z"/><path d="m3.876 3.435l-1.694 5.93a.57.57 0 0 0 .472.73c.845.105 2.33.217 3.346-.095c.595-.183 1.358-.613 1.747-.845a.578.578 0 0 0 .278-.451L8.5 3H4.453a.6.6 0 0 0-.577.435Z"/><path d="M3 10v9a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-9"/></g>`
	smallShopAltInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 10v9a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-9"/><path stroke-miterlimit="16" d="M14.833 21v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v6"/><path d="m21.818 9.364l-1.694-5.929A.6.6 0 0 0 19.547 3H15.5l.475 5.704a.578.578 0 0 0 .278.45c.39.233 1.152.663 1.747.846c1.016.313 2.5.2 3.346.096a.57.57 0 0 0 .472-.732Z"/><path d="M14 10c.568-.175 1.288-.574 1.69-.812a.578.578 0 0 0 .28-.549L15.5 3h-7l-.47 5.639a.578.578 0 0 0 .28.55c.402.237 1.122.636 1.69.811c1.493.46 2.507.46 4 0Z"/><path d="m3.876 3.435l-1.694 5.93a.57.57 0 0 0 .472.73c.845.105 2.33.217 3.346-.095c.595-.183 1.358-.613 1.747-.845a.578.578 0 0 0 .278-.451L8.5 3H4.453a.6.6 0 0 0-.577.435Z"/></g>`
	smartphoneDeviceInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m12 16.01l.01-.011"/><path d="M7 19.4V4.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6Z"/></g>`
	smokingInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 19v3m0-6c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V8.5A2.5 2.5 0 0 1 13.5 6v0h.5m8 10c0-4.5-2-5.5-4-6c2-.5 4-1 4-4s-2.5-4-4-4m4 17v3"/><rect width="12" height="3" x="2" y="19" rx=".6"/></g>`
	snapchatInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 8.75c0-9-12-9-12 0v1.5H4.372c-.583 0-.823.749-.348 1.088L6 12.75c-.333 1.167-1.7 3.7-4.5 4.5c.333.5 1.3 1.5 2.5 1.5l1 1.5l2.5-.5c.833.667 2.9 2 4.5 2s3.667-1.333 4.5-2l2.5.5l1-1.5c1.2 0 2.167-1 2.5-1.5c-2.8-.8-4.167-3.333-4.5-4.5l1.976-1.412c.475-.339.235-1.088-.348-1.088H18v-1.5Z"/>`
	snowInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12v5m0 5v-5m0 0l-4.5-2.5M12 17l4.5 2.5M12 17l4.5-2.5M12 17l-4.5 2.5M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`
	snowFlakeInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 7l3.5 2M21 17l-3.5-2M12 12L6.5 9m5.5 3l-5.5 3m5.5-3V5m0 7v6.5m0-6.5l5.5 3M12 12l5.5-3M12 2v3m0 17v-3.5M21 7l-3.5 2M3 17l3.5-2m0-6L3 10m3.5-1L6 5.5m.5 9.5L3 14m3.5 1L6 18.5M12 5L9.5 4M12 5l2.5-1M12 18.5l2.5 1.5M12 18.5L9.5 20m8-5l.5 3.5m-.5-3.5l3.5-1m-3.5-5l3.5 1m-3.5-1l.5-3.5"/>`
	soapInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 11a4 4 0 0 1 4-4h2a4 4 0 0 1 4 4v9.4a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6V11Zm0 2h10m-5-6V3m0 0H9m3 0h1"/>`
	soccerBallInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 8l3.804 2.764M12 8l-3.804 2.764M12 8V5m3.804 5.764l-1.453 4.472m1.453-4.472L18.5 9.5m-4.149 5.736H9.65m4.702 0L16 17.5m-6.351-2.264l-1.453-4.472m1.453 4.472L8 17.5m.196-6.736L5.5 9.5m0 0L2.05 13M5.5 9.5l-1-4.115m14 4.115l3.45 3.5M18.5 9.5l1-4.115M12 5L8.624 2.584M12 5l3.376-2.416M8 17.5L3.338 17M8 17.5l2.5 4.388M16 17.5l4.662-.5M16 17.5l-2.5 4.388M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z"/>`
	sofaInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 16v3M4 9V7a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2"/><path d="M20 9a2 2 0 0 0-2 2v2H6v-2a2 2 0 1 0-4 0v6h20v-6a2 2 0 0 0-2-2Zm2 7v3"/></g>`
	soilInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 4h20M3 8.01l.01-.011M3 16.01l.01-.011M6 12.01l.01-.011M6 20.01l.01-.011M9 8.01l.01-.011M9 16.01l.01-.011M12 12.01l.01-.011M12 20.01l.01-.011M15 8.01l.01-.011M15 16.01l.01-.011M18 12.01l.01-.011M18 20.01l.01-.011M21 8.01l.01-.011M21 16.01l.01-.011"/>`
	soilAltInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12h4m11 0h5M3 20.01l.01-.011M6 16.01l.01-.011M9 20.01l.01-.011M12 16.01l.01-.011M15 20.01l.01-.011M18 16.01l.01-.011M21 20.01l.01-.011M9 13s.9-3.741 3-6"/><path d="m16.186 2.241l.374 3.89c.243 2.523-1.649 4.77-4.172 5.012c-2.475.238-4.718-1.571-4.956-4.047a4.503 4.503 0 0 1 4.05-4.914l4.147-.4a.51.51 0 0 1 .557.46Z"/></g>`
	sortInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 14H2m6-4H2m4-4H2m10 12H2m17 2V4m0 16l3-3m-3 3l-3-3m3-13l3 3m-3-3l-3 3"/>`
	sortDownInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 10H2m8 4H2m4 4H2M18 6H2m17 4v10m0 0l3-3m-3 3l-3-3"/>`
	sortUpInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 14H2m8-4H2m4-4H2m16 12H2m17-4V4m0 0l3 3m-3-3l-3 3"/>`
	soundHighInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 14v-4a1 1 0 0 1 1-1h2.697a1 1 0 0 0 .555-.168l4.193-2.796A1 1 0 0 1 12 6.87v10.26a1 1 0 0 1-1.555.832l-4.193-2.795A1 1 0 0 0 5.697 15H3a1 1 0 0 1-1-1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.5 7.5S18 9 18 11.5s-1.5 4-1.5 4m3-11S22 7 22 11.5s-2.5 7-2.5 7"/></g>`
	soundLowInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18.5 7.5S20 9 20 11.5s-1.5 4-1.5 4"/><path d="M4 14v-4a1 1 0 0 1 1-1h2.697a1 1 0 0 0 .555-.168l4.193-2.796A1 1 0 0 1 14 6.87v10.26a1 1 0 0 1-1.555.832l-4.193-2.795A1 1 0 0 0 7.697 15H5a1 1 0 0 1-1-1Z"/></g>`
	soundMinInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M7 14v-4a1 1 0 0 1 1-1h2.697a1 1 0 0 0 .555-.168l4.193-2.796A1 1 0 0 1 17 6.87v10.26a1 1 0 0 1-1.555.832l-4.193-2.795a1 1 0 0 0-.555-.168H8a1 1 0 0 1-1-1Z"/>`
	soundOffInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m17 14l2-2m2-2l-2 2m0 0l-2-2m2 2l2 2"/><path d="M3 14v-4a1 1 0 0 1 1-1h2.697a1 1 0 0 0 .555-.168l4.193-2.796A1 1 0 0 1 13 6.87v10.26a1 1 0 0 1-1.555.832l-4.193-2.795A1 1 0 0 0 6.697 15H4a1 1 0 0 1-1-1Z"/></g>`
	spadesInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 14.5c3 4.5 9 4.47 9-.5c0-4-4-7-9-12c-5 5-9 8-9 12c0 4.97 6 5 9 .5Z"/><path d="m11.47 15.493l-3 5.625A.6.6 0 0 0 9 22h6a.6.6 0 0 0 .53-.882l-3-5.625a.6.6 0 0 0-1.06 0Z"/></g>`
	sphereInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M12 22c-3.314 0-6-4.477-6-10S8.686 2 12 2"/></g>`
	spiralInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.5 6.004C3.5 7.808 6.357 9 11.5 9c7 0 8-2.996 8-2.996S18.5 3 11.5 3c-5.143 0-8 1.2-8 3.004Zm0 6c0 1.804 2.857 2.996 8 2.996c7 0 8-2.996 8-2.996S18.5 9 11.5 9c-5.143 0-8 1.2-8 3.004Zm0 6c0 1.804 2.857 2.996 8 2.996c7 0 8-2.996 8-2.996S18.5 15 11.5 15c-5.143 0-8 1.2-8 3.004ZM19.5 12s1-.975 1-3s-1-3-1-3m1-2c0 1.35-1 2-1 2m0 12s1-.975 1-3s-1-3-1-3m1 8c0-1.35-1-2-1-2"/>`
	spockHandGestureInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m18 7.5l.919.153a2 2 0 0 1 1.623 2.407l-.528 2.376a.602.602 0 0 0-.014.13V17.5s0 0 0 0c0 2-1.6 4-4 4H9.42a2 2 0 0 1-1.519-.698l-4.548-5.307a1.582 1.582 0 0 1-.034-2.018v0a1.582 1.582 0 0 1 2.426-.054L8 16v-3.5"/><path d="m9 5l-.79.132a2 2 0 0 0-1.595 2.522L8 12.5m3 0L8.923 4.606a1.514 1.514 0 0 1 1.215-1.879v0a1.514 1.514 0 0 1 1.713 1.108L14 12m3 .5l1-5l.247-1.485a1.536 1.536 0 0 0-1.262-1.768v0a1.536 1.536 0 0 0-1.762 1.233L14 12"/></g>`
	squareInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/>`
	squareWaveInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h3V4h6v16h6v-8h3m-6.5 0h1m-7 0h1"/>`
	stackoverflowInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 15v6H5v-6m11 2H8m7.913-2.337L8.087 13m8.626-.62L9.463 9m8.71 1.642L12.044 5.5m7.99 3.304L15.109 2.5"/>`
	starInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.587 8.236l2.598-5.232a.911.911 0 0 1 1.63 0l2.598 5.232l5.808.844a.902.902 0 0 1 .503 1.542l-4.202 4.07l.992 5.75c.127.738-.653 1.3-1.32.952L12 18.678l-5.195 2.716c-.666.349-1.446-.214-1.319-.953l.992-5.75l-4.202-4.07a.902.902 0 0 1 .503-1.54l5.808-.845Z"/>`
	starDashedInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.806 5l-.99-1.996a.911.911 0 0 0-1.631 0l-.496.998m4.322 3.425l.402.809l1.452.211m2.905.423l1.451.21a.902.902 0 0 1 .503 1.542l-1.05 1.017m-2.102 2.035l-1.05 1.018l.248 1.437m.496 2.875l.248 1.437c.127.739-.653 1.302-1.32.953l-1.298-.679M10.428 19.5L12 18.678l1.299.679m-7.628.012l-.185 1.072c-.127.739.653 1.302 1.32.953l.847-.443M6.253 16l.225-1.308l-.695-.673M3.699 12l-1.423-1.378a.902.902 0 0 1 .503-1.542l1.11-.161M7 8.467l1.587-.231l.804-1.618"/>`
	starHalfDashedInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.815 3.004a.911.911 0 0 0-1.63 0l-.496.998M12 18.678l-1.572.822m-4.757-.131l-.185 1.072c-.127.739.653 1.302 1.32.953l.847-.443M6.253 16l.225-1.308l-.695-.673M3.699 12l-1.423-1.378a.902.902 0 0 1 .503-1.542l1.11-.161M7 8.467l1.587-.231l.804-1.618"/><path d="m15.413 8.236l-2.598-5.232A.899.899 0 0 0 12 2.5v16.178l5.195 2.716c.666.349 1.446-.214 1.319-.953l-.992-5.75l4.202-4.07a.902.902 0 0 0-.503-1.54l-5.808-.845Z"/></g>`
	statDownInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 10l8 8l3-3l5 5M16 4v8m0 0l3-3m-3 3l-3-3"/>`
	statUpInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 20v-8m0 0l3 3m-3-3l-3 3m-9-1l8-8l3 3l5-5"/>`
	statsDownSquareInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8 16V8m4 8v-5m4 5v-3"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	statsReportInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 9H6m9.5 2a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5ZM6 6h3m9 12l-4.5-3l-2.5 2l-5-4"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	statsUpSquareInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16 16V8m-4 8v-5m-4 5v-3"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	strategyInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 20.5C7 11 11.5 8 20 6"/><path d="m15.909 3.81l4.486 2.09l-2.092 4.486M5 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm11 13.243l2.121-2.122m0 0L20.243 16m-2.122 2.121L16 16m2.121 2.121l2.122 2.122"/></g>`
	stretchingInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM5 20l4.91-.524l2.726-5.238L13.727 9l-4.909 1.048l1.636 2.095m4.364 3.143H17V20"/>`
	strollerInnerSVG                          = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirStroller0)"><path d="M11.5 3a8.5 8.5 0 0 0-7.212 13m14.424 0A8.46 8.46 0 0 0 20 11.5v-2h2.5M8 21a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm7 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM11.5 3v9m-8 0h16"/></g><defs><clipPath id="iconoirStroller0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	styleBorderInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.499"><path stroke-dasharray="2 2" d="M16 2H8a6 6 0 0 0-6 6v8a6 6 0 0 0 6 6h8a6 6 0 0 0 6-6V8a6 6 0 0 0-6-6Z"/><path d="M16 5H8a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3Z"/></g>`
	submitDocumentInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 13V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H14"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20m-4 13h6m0 0l-3-3m3 3l-3 3"/></g>`
	substractInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 3.6v10.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6ZM13.5 21h3m4.5-7.5v3m0 3v.9a.6.6 0 0 1-.6.6h-.9m-9 0h-.9a.6.6 0 0 1-.6-.6v-.9M19.5 9h.9a.6.6 0 0 1 .6.6v.9"/><path d="M16.5 9H9.6a.6.6 0 0 0-.6.6v6.9"/></g>`
	suggestionInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v13.8a.6.6 0 0 1-.6.6h-4.14a.6.6 0 0 0-.438.189l-3.385 3.597a.6.6 0 0 1-.874 0l-3.385-3.597A.6.6 0 0 0 7.74 18H3.6a.6.6 0 0 1-.6-.6V3.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 7l1.425 2.575L16 11l-2.575 1.425L12 15l-1.425-2.575L8 11l2.575-1.425L12 7Z"/></g>`
	sunLightInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm10-6h1M12 2V1m0 22v-1m8-2l-1-1m1-15l-1 1M4 20l1-1M4 4l1 1m-4 7h1"/>`
	svgFormatInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.5 15h2A1.5 1.5 0 0 0 8 13.5v0A1.5 1.5 0 0 0 6.5 12H6a1.5 1.5 0 0 1-1.5-1.5v0A1.5 1.5 0 0 1 6 9h1.5m3 0l1.5 6l1.5-6m6 0h-3v6h3v-2.4"/><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/></g>`
	swimmingInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 15c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3M3 20c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3M5 10.5L9 8L7.813 6.516a1.262 1.262 0 0 1 .228-1.797v0a1.261 1.261 0 0 1 1.726.202L14 10m2.5-2a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`
	swipeDownGestureInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0 0v8m0 0l-3-3m3 3l3-3"/>`
	swipeLeftGestureInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12a6 6 0 1 0 12 0a6 6 0 0 0-12 0Zm0 0H2m0 0l3-3m-3 3l3 3"/>`
	swipeRightGestureInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 12a6 6 0 1 1-12 0a6 6 0 0 1 12 0Zm0 0h8m0 0l-3-3m3 3l-3 3"/>`
	swipeTwoFingersDownGestureInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 12a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0 0v7m0 0L9 16.6M6.5 19L4 16.6M17.5 12a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0 0v7m0 0l2.5-2.4M17.5 19L15 16.6"/>`
	swipeTwoFingersLeftGestureInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Zm0 0H5m0 0L7.4 15M5 17.5L7.4 20M12 6.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Zm0 0H5m0 0L7.4 4M5 6.5L7.4 9"/>`
	swipeTwoFingersRightGestureInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm0 0h7m0 0L16.6 15m2.4 2.5L16.6 20M12 6.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm0 0h7m0 0L16.6 4M19 6.5L16.6 9"/>`
	swipeTwoFingersUpGestureInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 12a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm0 0V5m0 0L9 7.4M6.5 5L4 7.4M17.5 12a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm0 0V5m0 0L20 7.4M17.5 5L15 7.4"/>`
	swipeUpGestureInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 10a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm0 0V2m0 0l3 3m-3-3L9 5"/>`
	switchOffInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M17 17H7A5 5 0 0 1 7 7h10a5 5 0 0 1 0 10Z"/></g>`
	switchOnInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M17 17H7A5 5 0 0 1 7 7h10a5 5 0 0 1 0 10Z"/></g>`
	systemRestartInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2v4m0 12v4m10-10h-4M6 12H2m2.929-7.071l2.828 2.828m8.486 8.486l2.828 2.828m0-14.142l-2.828 2.828m-8.486 8.486L4.93 19.07"/>`
	systemShutInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7v10m0 5c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	tShirtInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 4h3s0 3 3 3s3-3 3-3h3m0 7v8.4a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6V11m12-7l4.443 1.777a.6.6 0 0 1 .334.78l-1.626 4.066a.6.6 0 0 1-.557.377H18M6 4L1.557 5.777a.6.6 0 0 0-.334.78l1.626 4.066a.6.6 0 0 0 .557.377H6"/>`
	tableInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3v18H3V3h18ZM3 16.5h18M3 12h18M3 7.5h18M16.5 3v18M12 3v18M7.5 3v18"/>`
	tableRowsInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 12h18M3 12v4.5M3 12V7.5M21 12v4.5m0-4.5V7.5m-18 9v3.9a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6v-3.9m-18 0h18m0-9V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v3.9m18 0H3"/>`
	tableTwoColumnsInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm0-3.9h18M3 12h18m0-4.5H3M12 21V3"/>`
	taskListInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6h11M3.8 5.8l.8.8l2-2m-2.8 7.2l.8.8l2-2m-2.8 7.2l.8.8l2-2M9 12h11M9 18h11"/>`
	telegramInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 5L2 12.5l7 1M21 5l-2.5 15L9 13.5M21 5L9 13.5m0 0V19l3.249-3.277"/>`
	telegramCircleInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 8L5 12.5L9.5 14M18 8l-8.5 6M18 8l-4 10.5L9.5 14"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	tennisBallInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M18.572 4.462c-2.667 4.53-2.667 9.723 0 15.076M5.428 4.462c2.667 4.53 2.667 9.723 0 15.076"/></g>`
	tennisBallAltInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.66 7c2.762 4.783 1.123 10.899-3.66 13.66C12.217 23.422 6.101 21.783 3.34 17C.578 12.217 2.217 6.1 7 3.34C11.783.578 17.899 2.217 20.66 7Z"/><path d="M21.46 15.242c-4.986-3.303-7.582-7.8-7.538-13.056m-3.844 19.628C9.71 15.844 7.114 11.347 2.54 8.758"/></g>`
	terminalInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 17h7M5 7l5 5l-5 5"/>`
	terminalTagInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 16h5M6 8l4 4l-4 4"/><path d="M2 18V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/></g>`
	testTubeInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.141 19.995c2.458 1.72 4.281-.012 5.318-1.492l7.3-10.426l1.966-1.065l-6.553-4.588l-8.447 12.064c-1.037 1.48-2.041 3.786.416 5.507Z"/><path d="M16.091 11.02c-2.876-.853-4.403.781-7.28-.071"/></g>`
	textInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7V5H5v2m7-2v14m0 0h-2m2 0h2"/>`
	textAltInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M7 9V7h10v2m-5-2v10m0 0h-2m2 0h2"/></g>`
	textBoxInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M12 8v8m0-8H8m4 0h4"/><path d="M2 20.4V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path d="M1 13v-2h2v2H1Zm20 0v-2h2v2h-2Z"/></g>`
	textSizeInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 8V6h12v2m-6-2v12m0 0h2m-2 0H8m6-4.5V12h6v1.5M17 12v6m0 0h-1.5m1.5 0h1.5"/>`
	threeDAddHoleInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524ZM3.528 7.294L8.4 10m12.1-2.722L15.6 10M12 21v-5"/></g>`
	threeDArcInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 16c0-5.523-4.477-10-10-10S2 10.477 2 16"/><path fill="currentColor" d="M2 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm20 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	threeDArcCenterPtInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-dasharray="3 3" d="M22 16c0-5.523-4.477-10-10-10c-4.1 0-7.625 2.468-9.168 6"/><path fill="currentColor" d="M2 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M2 16h10"/><path fill="currentColor" d="M12 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	threeDBridgeInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 4h3"/><path fill="currentColor" d="M10 21a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-16a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M10 20s6.5-2.5 2-8s2-8 2-8M3 20h3"/></g>`
	threeDCenterBoxInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524Z"/><path d="m20.5 16.722l-8.209-4.56a.6.6 0 0 0-.582 0L3.5 16.722m.028-9.428l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 3v9m0 7.5V22"/></g>`
	threeDEllipseInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M12 22c4.418 0 8-4.477 8-10S16.418 2 12 2S4 6.477 4 12s3.582 10 8 10Z"/><path fill="currentColor" d="M12 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	threeDEllipseThreePtsInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M5 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M5 22h8m-8 0V2"/><path fill="currentColor" d="M5 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path stroke-dasharray="3 3" d="M8 4.193C9.37 2.821 11.108 2 13 2c4.418 0 8 4.477 8 10c0 3.271-1.256 6.175-3.2 8"/><path d="M8.2 20A9.098 9.098 0 0 1 7 18.615"/><path fill="currentColor" d="M13 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	threeDPtBoxInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M3 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524Z"/><path d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21v-9"/></g>`
	threeDRectCornerToCornerInnerSVG          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm18 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	threeDRectFromCenterInnerSVG              = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	threeDRectThreePtsInnerSVG                = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 21V3.6a.6.6 0 0 1 .6-.6H21"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 21h3.4a.6.6 0 0 0 .6-.6V17m0-10v2m0 3v2M7 21h2m3 0h2"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM21 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`
	threeDSelectEdgeInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524ZM12 21v-9"/><path fill="currentColor" d="M12.5 11v10a.5.5 0 0 1-1 0V11a.5.5 0 0 1 1 0Z"/></g>`
	threeDSelectFaceInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524Z"/><path stroke-linecap="round" d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21v-9"/><path fill="currentColor" d="m11.691 11.829l-7.8-4.334A.6.6 0 0 0 3 8.02v8.627a.6.6 0 0 0 .309.525l7.8 4.333A.6.6 0 0 0 12 20.98v-8.627a.6.6 0 0 0-.309-.524Z"/></g>`
	threeDSelectPointInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524Z"/></g>`
	threeDSelectSolidInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524Z"/><path d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21v-9"/></g>`
	threeDThreePtsBoxInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM3 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524Z"/><path d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21v-9"/></g>`
	threeHundredSixtyViewInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 8.5h1.75m0 0a1.75 1.75 0 1 1 0 3.5H3m2.75-3.5a1.75 1.75 0 1 0 0-3.5H3m18 10c0 3.314-4.03 6-9 6s-9-2.686-9-6M14 5h-1a3 3 0 0 0-3 3v2m4.5-.5v.5a2 2 0 0 1-2 2H12a2 2 0 0 1-2-2v-.5a2 2 0 0 1 2-2h.5a2 2 0 0 1 2 2Zm2.5-1V7a2 2 0 0 1 2-2h.5a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H19a2 2 0 0 1-2-2V8.5Z"/>`
	threePointsCircleInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M5 10.5V9m0 6v-1.5"/><path fill="currentColor" d="M5 20a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm14 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M10.5 19H9m6 0h-1.5"/></g>`
	threeStarsInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.635 14.415l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L6 18.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.217-.22-.098-.604.2-.65l2.324-.355Zm12 0l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L18 18.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.216-.22-.098-.604.2-.65l2.324-.355Zm-6-9l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L12 9.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.217-.22-.098-.604.2-.65l2.324-.355Z"/>`
	thumbsDownInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M16.472 3.5H4.1a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6h2.768a2 2 0 0 1 1.715.971l2.71 4.517a1.631 1.631 0 0 0 2.961-1.308l-1.022-3.408a.6.6 0 0 1 .574-.772h4.575a2 2 0 0 0 1.93-2.526l-1.91-7A2 2 0 0 0 16.473 3.5Z"/><path stroke-linejoin="round" d="M7 14.5v-11"/></g>`
	thumbsUpInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M16.472 20H4.1a.6.6 0 0 1-.6-.6V9.6a.6.6 0 0 1 .6-.6h2.768a2 2 0 0 0 1.715-.971l2.71-4.517a1.631 1.631 0 0 1 2.961 1.308l-1.022 3.408a.6.6 0 0 0 .574.772h4.575a2 2 0 0 1 1.93 2.526l-1.91 7A2 2 0 0 1 16.473 20Z"/><path stroke-linejoin="round" d="M7 20V9"/></g>`
	thunderstormInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.5 12L9 17h6l-2.5 5"/><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/></g>`
	tifFormatInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6"/><path stroke-linejoin="round" d="M15 15V9h3M6.5 9H8m1.5 0H8m0 0v6m7-3h2.5M12 15V9"/><path d="M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/></g>`
	tiffFormatInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M12 15V9h3m2.5 6V9h3m-17 0H5m1.5 0H5m0 0v6m7-3h2.5m3 0H20M9 15V9"/></g>`
	tiktokInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5Z"/><path d="M10 12a3 3 0 1 0 3 3V6c.333 1 1.6 3 4 3"/></g>`
	timerInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 2h6m-3 8v4m0 8a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/>`
	timerOffInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 2h6M5 7l14 14.5M12 10v4M6.19 8.5a8 8 0 0 0 11.05 11.544M19.42 17A8 8 0 0 0 9.21 6.5"/>`
	toolsInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.05 10.607l-7.07 7.07a2 2 0 0 0 0 2.83v0a2 2 0 0 0 2.828 0l7.07-7.072m4.315.365l3.878 3.878a2 2 0 0 1 0 2.828v0a2 2 0 0 1-2.828 0l-6.209-6.208M6.733 5.904L4.61 6.61L2.49 3.075l1.414-1.414L7.44 3.782l-.707 2.122Zm0 0l2.83 2.83"/><path d="M10.05 10.607c-.844-2.153-.679-4.978 1.061-6.718c1.74-1.74 4.95-2.121 6.717-1.06l-3.04 3.04l-.283 3.111l3.111-.282l3.04-3.041c1.062 1.768.68 4.978-1.06 6.717c-1.74 1.74-4.564 1.905-6.717 1.061"/></g>`
	tournamentInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h5v6H3m5-3h7v12H8m7-6h7M3 15h5v6H3"/>`
	towerInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M17 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333l-1.8 2.698a.6.6 0 0 0-.1.333V20a2 2 0 0 1-2 2Z"/>`
	towerCheckInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13 19l3 3l5-5M9 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333l-1.8 2.698a.6.6 0 0 0-.1.333V13.5"/>`
	towerNoAccessInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.857 15.2a4 4 0 0 0-5.713 5.6m5.713-5.6a4 4 0 0 1-5.713 5.6m5.713-5.6l-5.714 5.6"/><path d="M9 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333L20 9.5"/></g>`
	towerWarningInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 10v3m0 4.01l.01-.011"/><path d="M17 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333l-1.8 2.698a.6.6 0 0 0-.1.333V20a2 2 0 0 1-2 2Z"/></g>`
	trademarkInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 15.5v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H9.5v4.2m2.857 0H9.5m2.857 0l2.143 2.8"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	trainInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M9.609 7h4.782A2.609 2.609 0 0 1 17 9.609a.391.391 0 0 1-.391.391H7.39A.391.391 0 0 1 7 9.609A2.609 2.609 0 0 1 9.609 7Z"/><path stroke-linejoin="round" d="M9 3h6a6 6 0 0 1 6 6v4a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6V9a6 6 0 0 1 6-6Zm7 12.01l.01-.011M8 15.01l.01-.011"/><path d="m10.5 19l-2 2.5m5-2.5l2 2.5m1-2.5l2 2.5M7.5 19l-2 2.5"/></g>`
	tramInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m15 16.01l.01-.011M9 16.01l.01-.011M13 6h2a5 5 0 0 1 5 5v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-7a5 5 0 0 1 5-5h4Zm0 0l1-4m0 0h3m-3 0H7"/><path d="m10.5 20l-2 2.5m5-2.5l2 2.5m1-2.5l2 2.5M7.5 20l-2 2.5"/><path stroke-linejoin="round" d="M9.609 9h4.782A2.609 2.609 0 0 1 17 11.609a.391.391 0 0 1-.391.391H7.39a.391.391 0 0 1-.39-.391A2.609 2.609 0 0 1 9.609 9Z"/></g>`
	transitionDownInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M18 2H6a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 16v2a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-2m-9-6v8m0 0l-3-3m3 3l3-3"/></g>`
	transitionLeftInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M22 18V6a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 3H6a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h2m6-9H6m0 0l3-3m-3 3l3 3"/></g>`
	transitionRightInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 18V6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 3h2a4 4 0 0 1 4 4v10a4 4 0 0 1-4 4h-2m-6-9h8m0 0l-3-3m3 3l-3 3"/></g>`
	transitionUpInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M18 22H6a3 3 0 0 1-3-3v-2a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 8V6a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v2m-9 6V6m0 0L9 9m3-3l3 3"/></g>`
	translateInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 5h7m7 0h-2.5M9 5h4.5M9 5V3m4.5 2c-.82 2.735-2.539 5.32-4.5 7.593M4 17.5c1.585-1.359 3.376-3.026 5-4.907m0 0C8 11.5 6.4 9.3 6 8.5m3 4.093l3 2.907m1.5 5.5l1.143-3m6.857 3l-1.143-3m-5.714 0l2.857-7.5l2.857 7.5m-5.714 0h5.714"/>`
	trashInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 9l-1.995 11.346A2 2 0 0 1 16.035 22h-8.07a2 2 0 0 1-1.97-1.654L4 9m17-3h-5.625M3 6h5.625m0 0V4a2 2 0 0 1 2-2h2.75a2 2 0 0 1 2 2v2m-6.75 0h6.75"/>`
	treadmillInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-2.387 1.267l-3.308 4.135l4.135 4.135l-2.067 4.55"/><path d="m4.41 8.508l3.387-3.309l2.816 2.068l2.895 3.308h1.722M6.892 14.71l-1.241.827H2.343m1 6l15.308-2V8"/><path d="M20.892 6L18.65 8L17 9.5m3.891 12.21l-2.24-2.173"/></g>`
	treeInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22v-8m0-4v4m0 0l4-2m1-5A5 5 0 0 0 7 7m5 11H7.5a5.5 5.5 0 1 1 0-11H9m3 11h4.5A5.5 5.5 0 0 0 17 7.022"/>`
	trekkingInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18 10l-3 1.5l-4-3l-1 5.5l3.5 3l.5 4.5m4-13v13M10 17l-2 4.5m.5-13C7 9.5 6 12 6 12l2 1m4-6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`
	trelloInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Z"/><path d="M10.4 6H6.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 0 .6.6h3.8a.6.6 0 0 0 .6-.6V6.6a.6.6 0 0 0-.6-.6Zm7 0h-3.8a.6.6 0 0 0-.6.6v6.8a.6.6 0 0 0 .6.6h3.8a.6.6 0 0 0 .6-.6V6.6a.6.6 0 0 0-.6-.6Z"/></g>`
	triangleInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.475 2.947a.6.6 0 0 1 1.05 0l9.373 16.912a.6.6 0 0 1-.524.891H2.626a.6.6 0 0 1-.525-.89l9.374-16.913Z"/>`
	triangleFlagInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 21v-5m0 0V3.577a.6.6 0 0 1 .916-.51l8.79 5.442a.6.6 0 0 1 .017 1.009L8 16Z"/>`
	triangleFlagCircleInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 21.5v-6m0 0V6.997a.6.6 0 0 1 .88-.53l6.67 3.53a.6.6 0 0 1 .024 1.048L9 15.5ZM22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z"/>`
	triangleFlagTwoStripesInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 21v-5m0 0l9.723-6.482a.6.6 0 0 0-.017-1.01l-8.79-5.441a.6.6 0 0 0-.916.51V16Zm0-5l6.5-4.476"/>`
	trophyInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.745 4h10.568s-.88 13.257-5.284 13.257c-2.15 0-3.461-3.164-4.239-6.4C6.976 7.468 6.745 4 6.745 4Z"/><path d="M17.313 4s.921-.983 1.687-1c1.5-.034 1.777 1 1.777 1c.294.61.529 2.194-.88 3.657c-1.41 1.463-2.987 2.743-3.629 3.2M6.745 4S5.785 3.006 5 3c-1.5-.012-1.777 1-1.777 1c-.294.61-.529 2.194.88 3.657a29.896 29.896 0 0 0 3.687 3.2M8.507 20c0-1.829 3.522-2.743 3.522-2.743s3.523.914 3.523 2.743H8.507Z"/></g>`
	truckInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M7 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M14 17V6.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6h2.05M14 17H9.05M14 9h5.61a.6.6 0 0 1 .548.356l1.79 4.028a.6.6 0 0 1 .052.243V16.4a.6.6 0 0 1-.6.6h-1.9M14 17h1"/></g>`
	truckLengthInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M7 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M14 14V3.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6h2.05M14 14H9.05M14 6h5.61a.6.6 0 0 1 .548.356l1.79 4.028a.6.6 0 0 1 .052.243V13.4a.6.6 0 0 1-.6.6h-1.9M14 14h1"/><path stroke-linejoin="round" d="M3 20h17.75M3 20l1.75 1.75M3 20l1.75-1.75m16 1.75L19 21.75M20.75 20L19 18.25"/></g>`
	tunnelInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21 20L3 14"/><path d="M16 10v1m-4-2v1M8 8v1"/><path stroke-linejoin="round" d="M3 21h18v-9a9 9 0 1 0-18 0v9Z"/></g>`
	tvInnerSVG                                = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 20V9a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.5 2.5L12 6l3.5-3.5"/></g>`
	tvFixInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 20V9a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m13.657 12.828l-2.829 2.829m5.657-2.829A2 2 0 1 1 13.657 10m-2.829 8.485A2 2 0 0 0 8 15.657M8.5 2.5L12 6l3.5-3.5"/></g>`
	tvIssueInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 11v3m0 4.01l.01-.011"/><path d="M2 20V9a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.5 2.5L12 6l3.5-3.5"/></g>`
	twitterInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M23 3.01s-2.018 1.192-3.14 1.53a4.48 4.48 0 0 0-7.86 3v1a10.66 10.66 0 0 1-9-4.53s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5c0-.278-.028-.556-.08-.83C21.94 5.674 23 3.01 23 3.01Z"/>`
	twitterVerifiedBadgeInnerSVG              = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M10.521 2.624a2 2 0 0 1 2.958 0l1.02 1.12a2 2 0 0 0 1.572.651l1.513-.07a2 2 0 0 1 2.092 2.09l-.071 1.514a2 2 0 0 0 .651 1.572l1.12 1.02a2 2 0 0 1 0 2.958l-1.12 1.02a2 2 0 0 0-.651 1.572l.07 1.513a2 2 0 0 1-2.09 2.092l-1.514-.071a2 2 0 0 0-1.572.651l-1.02 1.12a2 2 0 0 1-2.958 0l-1.02-1.12a2 2 0 0 0-1.572-.651l-1.513.07a2 2 0 0 1-2.092-2.09l.071-1.514a2 2 0 0 0-.651-1.572l-1.12-1.02a2 2 0 0 1 0-2.958l1.12-1.02a2 2 0 0 0 .651-1.572l-.07-1.513a2 2 0 0 1 2.09-2.092l1.514.071a2 2 0 0 0 1.572-.651l1.02-1.12Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 12l2 2l4-4"/></g>`
	twoPointsCircleInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm14 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="m19 19l-1.5-1.5m-2-2l-1-1m-2-2l-1-1m-2-2l-1-1m-2-2L5 5"/></g>`
	twoSeaterSofaInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 16v3m10-6V7a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2m-8 4V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v2"/><path d="M20 9a2 2 0 0 0-2 2v2H6v-2a2 2 0 1 0-4 0v6h20v-6a2 2 0 0 0-2-2Zm2 7v3"/></g>`
	twoXTwoCellInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M21 3.6V12h-9V3h8.4a.6.6 0 0 1 .6.6Zm0 16.8V12h-9v9h8.4a.6.6 0 0 0 .6-.6ZM3 12V3.6a.6.6 0 0 1 .6-.6H12v9H3Zm0 0v8.4a.6.6 0 0 0 .6.6H12v-9H3Z"/>`
	typeInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 16.248a.6.6 0 0 0-.176-.424l-3.648-3.648A.6.6 0 0 1 3 11.75V4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v7.752a.6.6 0 0 1-.176.424l-3.648 3.648a.6.6 0 0 0-.176.425V20a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-3.752Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9.5 11.5l.5-1.1m4.5 1.1l-.5-1.1m0 0L12 6l-2 4.4m4 0h-4"/></g>`
	umbrellaFullInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19.778 4.043C17.701 2.081 14.938 1 12 1C9.062 1 6.3 2.08 4.222 4.043C2.144 6.006 1 8.616 1 11.391c0 .336.289.609.644.609c.356 0 .645-.273.645-.609c0-1.013.872-1.837 1.944-1.837C6.126 9.554 5.431 12 6.823 12c1.39 0 .696-2.446 2.588-2.446C11.304 9.554 12 12 12 12s.697-2.446 2.589-2.446S15.988 12 17.178 12s.696-2.446 2.589-2.446c1.072 0 1.944.824 1.944 1.837c0 .336.289.609.645.609c.355 0 .644-.273.644-.609c0-2.775-1.144-5.385-3.222-7.348Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12v8c0 4-6 4-6 0"/></g>`
	underlineInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 5v6a4 4 0 0 1-4 4v0a4 4 0 0 1-4-4V5M6 19h12"/>`
	underlineSquareInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 6v4a4 4 0 0 1-4 4v0a4 4 0 0 1-4-4V6M6 18h12"/></g>`
	undoInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 9.5h10c.162 0 4 0 4 4c0 4.5-3.702 4.5-4 4.5H7"/><path d="M8.5 13L5 9.5L8.5 6"/></g>`
	undoActionInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 7v5m2.875-2.5h7c5.5 0 5.5 8.5 0 8.5h-10"/><path d="m11.375 13l-3.5-3.5l3.5-3.5"/></g>`
	undoCircleInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 10.625h7.2s0 0 0 0s2.8 0 2.8 3C17 17 14.2 17 14.2 17h-.8"/><path d="M10.5 14L7 10.625L10.5 7"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	unionInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 9h5.4a.6.6 0 0 1 .6.6v10.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6V15"/><path d="M15 9V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 0 .6.6H9"/></g>`
	unionAltInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 22A7 7 0 1 0 9 8a7 7 0 0 0 0 14Z"/><path d="M15 16a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/></g>`
	unionHorizAltInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 19A7 7 0 1 0 8 5a7 7 0 0 0 0 14Z"/><path d="M16 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/></g>`
	unityInnerSVG                             = `<g fill="none" stroke-width="1.5"><g clip-path="url(#iconoirUnity0)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 13h9.5M4 13l4 4.5M4 13l4-4.5m5.5 4.5l5-9m-5 9l5 7m0-16l-6 1m6-1L20 9.5M18.5 20l1.5-5.5M18.5 20l-6-.5"/></g><defs><clipPath id="iconoirUnity0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	unityFiveInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.672 20.786a.6.6 0 0 0 .656 0l9.284-6.062a.6.6 0 0 0 .24-.694L18.285 3.41a.6.6 0 0 0-.569-.41H6.221a.6.6 0 0 0-.57.412l-3.506 10.62a.6.6 0 0 0 .241.69l9.286 6.064Z"/><path d="M14.5 6H10l-.5 5a3 3 0 1 1 0 3"/></g>`
	upRoundArrowInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.5 13.25l-2.5-2.5l-2.5 2.5"/><path d="M6 5h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V9a4 4 0 0 1 4-4Z"/></g>`
	uploadInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20h12m-6-4V4m0 0l3.5 3.5M12 4L8.5 7.5"/>`
	uploadDataWindowInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M19.5 22v-6m0 0L17 18.5m2.5-2.5l2.5 2.5"/></g>`
	uploadSquareInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 18h12m-6-4V6m0 0l3.5 3.5M12 6L8.5 9.5"/></g>`
	usbInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.5 2v15m0-3l5.5-2V8.5M12.5 16L7 14.5v-3M12.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm4-16.5v3h3v-3h-3Zm-6-1.5l2-2l2 2M7 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`
	userInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1m-7-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`
	userBagInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M2 18a7 7 0 0 1 11.33-5.5m8.034 4.207l.296 2A2 2 0 0 1 19.682 21h-3.364a2 2 0 0 1-1.978-2.293l.296-2A2 2 0 0 1 16.614 15h2.772a2 2 0 0 1 1.978 1.707ZM17 13h2"/></g>`
	userCartInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m22 12.5l-.833 2.5m0 0L20 18.5h-4.5l-1-3.5h6.667ZM16.5 20.51l.01-.011m2.99.011l.01-.011M9 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M2 18a7 7 0 0 1 11.33-5.5"/></g>`
	userCircleInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M7 18v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><circle cx="12" cy="12" r="10"/></g>`
	userCrownInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M5 20v-1a7 7 0 0 1 10-6.326M21 22l1-6l-3.5 1.8L17 16l-1.5 1.8L12 16l1 6h8Z"/></g>`
	userLoveInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M12 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M22 17.28a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.823-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.233 0l.106.107l.106-.107A2.277 2.277 0 0 1 22 17.28Z"/><path stroke-linecap="round" d="M5 20v-1a7 7 0 0 1 10-6.326"/></g>`
	userScanInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3H3v3m15-3h3v3M6 21H3v-3m4 0v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1m-5-6a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm6 9h3v-3"/>`
	userSquareInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M7 18v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1"/><path stroke-linejoin="round" d="M12 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path stroke-linejoin="round" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/></g>`
	userStarInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M5 20v-1a7 7 0 0 1 10-6.326m1.635 3.741l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L18 20.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.216-.22-.098-.604.2-.65l2.324-.355Z"/></g>`
	veganInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 11.063C12.53 13.65 10.059 20 10.059 20S6.529 11.062 3 9"/><path d="m20.496 5.577l.426 4.424c.276 2.87-1.875 5.425-4.745 5.702c-2.816.27-5.367-1.788-5.638-4.604a5.122 5.122 0 0 1 4.608-5.59l4.716-.454a.58.58 0 0 1 .633.522Z"/></g>`
	veganCircleInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 11.5C12.75 13.382 11 18 11 18s-2.5-6.5-5-8"/><path d="m18.015 7.73l.297 3.08c.192 1.998-1.306 3.777-3.304 3.97c-1.96.188-3.736-1.245-3.925-3.205a3.566 3.566 0 0 1 3.208-3.892l3.284-.316a.404.404 0 0 1 .44.363Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/></g>`
	veganSquareInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M14.5 10.5C12.75 12.382 11 17 11 17s-2.5-6.5-5-8"/><path d="m18.015 6.73l.297 3.08c.192 1.998-1.306 3.777-3.304 3.97c-1.96.188-3.736-1.245-3.925-3.205a3.566 3.566 0 0 1 3.208-3.892l3.284-.316a.404.404 0 0 1 .44.363Z"/></g>`
	verifiedBadgeInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11.528 1.6a.6.6 0 0 1 .944 0l1.809 2.3a.6.6 0 0 0 .635.207l2.815-.798a.6.6 0 0 1 .764.554l.11 2.925a.6.6 0 0 0 .393.54l2.747 1.01a.6.6 0 0 1 .292.897l-1.63 2.431a.6.6 0 0 0 0 .668l1.63 2.431a.6.6 0 0 1-.292.897l-2.747 1.01a.6.6 0 0 0-.392.54l-.111 2.925a.6.6 0 0 1-.764.554l-2.815-.798a.6.6 0 0 0-.636.206L12.473 22.4a.6.6 0 0 1-.944 0L9.72 20.1a.6.6 0 0 0-.635-.207l-2.815.798a.6.6 0 0 1-.764-.554l-.11-2.925a.6.6 0 0 0-.393-.54l-2.747-1.01a.6.6 0 0 1-.292-.897l1.63-2.431a.6.6 0 0 0 0-.668l-1.63-2.431a.6.6 0 0 1 .292-.897l2.747-1.01a.6.6 0 0 0 .392-.54l.111-2.925a.6.6 0 0 1 .764-.554l2.815.798A.6.6 0 0 0 9.72 3.9l1.81-2.3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 12l2 2l4-4"/></g>`
	verifiedUserInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2 20v-1a7 7 0 0 1 7-7v0"/><path d="M15.804 12.313a1.618 1.618 0 0 1 2.392 0c.325.357.79.55 1.272.527a1.618 1.618 0 0 1 1.692 1.692c-.023.481.17.947.526 1.272c.705.642.705 1.75 0 2.392c-.356.325-.549.79-.526 1.272a1.618 1.618 0 0 1-1.692 1.692a1.618 1.618 0 0 0-1.272.526a1.618 1.618 0 0 1-2.392 0a1.618 1.618 0 0 0-1.272-.526a1.618 1.618 0 0 1-1.692-1.692a1.618 1.618 0 0 0-.527-1.272a1.618 1.618 0 0 1 0-2.392c.357-.325.55-.79.527-1.272a1.618 1.618 0 0 1 1.692-1.692c.481.023.947-.17 1.272-.527Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.364 17l1.09 1.09l2.182-2.18M9 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`
	verticalMergeInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2v8m0 0l3.5-3.5M12 10L8.5 6.5M12 22v-8m0 0l3.5 3.5M12 14l-3.5 3.5M3 14h18M3 10h18"/>`
	verticalSplitInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14v8m0 0l3.5-3.5M12 22l-3.5-3.5M12 10V2m0 0l3.5 3.5M12 2L8.5 5.5M3 14h18M3 10h18"/>`
	vialsInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21H3m6-9H5m14 0h-4m-8 6a2 2 0 0 1-2-2V3h4v13a2 2 0 0 1-2 2Zm10 0a2 2 0 0 1-2-2V3h4v13a2 2 0 0 1-2 2Z"/>`
	videoCameraInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 12v4.4a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6V12Zm0 0l5.016-4.18a.6.6 0 0 1 .984.461v7.438a.6.6 0 0 1-.984.46L15 12Z"/>`
	videoCameraOffInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 7H3.6a.6.6 0 0 0-.6.6v8.8a.6.6 0 0 0 .6.6h10.8a.6.6 0 0 0 .6-.6V15m-3.5-8h2.9a.6.6 0 0 1 .6.6v3.119a.6.6 0 0 0 .984.46l4.032-3.359a.6.6 0 0 1 .984.461V15.5M3 3l18 18"/>`
	videoProjectorInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M4 19h2m12 0h2"/><path d="M2 16.4V7.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v8.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m5 10.01l.01-.011M8 10.01l.01-.011m2.99.011l.01-.011M5 14.01l.01-.011M8 14.01l.01-.011m2.99.011l.01-.011M17 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	viewColumnsThreeInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M9 3H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H9M9 3v18M9 3h6M9 21h6m0-18h5.4a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H15m0-18v18"/>`
	viewColumnsTwoInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M12 3h8.4a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H12m0-18H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H12m0-18v18"/>`
	viewGridInnerSVG                          = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M14 20.4v-5.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 0v-5.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm11-11V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 0V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/>`
	viewStructureDownInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 20.4v-5.8a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm11-11V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 0V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/>`
	viewStructureUpInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 9.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm11 11v-5.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 0v-5.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/>`
	voiceInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4v16M8 9v6m12-5v4M4 10v4m12-7v10"/>`
	voiceCircleInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v12M9 9v6m9-4v2M6 11v2m9-6v10m-3 5c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	voiceErrorInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v16M8 8v6m12-5v4M4 9v4m12-7v8m.121 7.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L16.12 17.12m2.122 2.122l2.121 2.121"/>`
	voiceLockCircleInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 21.8c-.646.131-1.315.2-2 .2c-5.523 0-10-4.477-10-10S6.477 2 12 2s10 4.477 10 10c0 .254-.01.506-.028.755M12 6v12M9 9v6m9-4v2M6 11v2m9-6v10"/><path d="M21.167 18.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/></g>`
	voiceOkInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v16M8 8v6m12-5v4M4 9v4m12-7v9m-.5 4.5l2 2l5-5"/>`
	voicePhoneInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 17.01l.01-.011M8 5H3.6a.6.6 0 0 0-.6.6v14.8a.6.6 0 0 0 .6.6h8.8a.6.6 0 0 0 .6-.6V16m3-13v10m-3-8v6m9-4v2M10 7v2m9-5v8"/>`
	voiceScanInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v12M9 9v6m9-4v2M6 11v2m9-6v10M6 3H3v3m15-3h3v3M6 21H3v-3m15 3h3v-3"/>`
	voiceSquareInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6ZM12 6v12M9 9v6m9-4v2M6 11v2m9-6v10"/>`
	vrSymbolInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 15.5v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H13v4.2m2.857 0H13m2.857 0L18 15.5m-7-7l-3 7l-3-7"/></g>`
	waistInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.4 4s-1.6 3.75-1.6 6.857c0 .995.34 1.827.8 2.656c.528.954 1.214 1.903 1.717 3.09A8.49 8.49 0 0 1 20 20M5.6 4s1.6 3.75 1.6 6.857c0 .995-.34 1.827-.8 2.656c-.528.954-1.214 1.903-1.717 3.09A8.483 8.483 0 0 0 4 20m2.4-6.487h11.2"/><path d="M4.683 16.604S10.4 17.713 12 20c1.6-2.286 7.317-3.396 7.317-3.396"/></g>`
	walkingInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12.44 9.127l-1.408 5.635l4.93 6.339m-5.634-2.817L8.215 21.1"/><path d="M8.215 13.353c0-3.944 2.817-4.226 4.226-4.226h1.408c.235 1.174 1.268 3.663 3.522 4.226M13 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	walletInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 20H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path d="M18 7V5.603a2 2 0 0 0-2.515-1.932l-11 2.933A2 2 0 0 0 3 8.537V9"/></g>`
	warningCircleInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7v6m0 4.01l.01-.011M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/>`
	warningHexagonInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.173 12.3a.6.6 0 0 1 0-.6l5.154-8.926a.6.6 0 0 1 .52-.3h10.307a.6.6 0 0 1 .52.3l5.153 8.926a.6.6 0 0 1 0 .6l-5.154 8.926a.6.6 0 0 1-.52.3H6.847a.6.6 0 0 1-.52-.3L1.174 12.3ZM12 8v4m0 4.01l.01-.011"/>`
	warningSquareInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 7v6m0 4.01l.01-.011"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`
	warningTriangleInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20.043 21H3.957c-1.538 0-2.5-1.664-1.734-2.997l8.043-13.988c.77-1.337 2.699-1.337 3.468 0l8.043 13.988C22.543 19.336 21.58 21 20.043 21ZM12 9v4"/><path stroke-linejoin="round" d="m12 17.01l.01-.011"/></g>`
	warningWindowInnerSVG                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M18 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M21 16v2m0 4.01l.01-.011"/></g>`
	washInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m22 5l-1.954 12.314A2 2 0 0 1 18.07 19H5.93a2 2 0 0 1-1.975-1.686L2 5"/><path d="M21 11c-2 0-4.5-3-4.5-3s-2.149 3-4.5 3s-4.5-3-4.5-3S5 11 3 11"/></g>`
	washingMachineInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 4v16a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2Zm-3 1.01l.01-.011"/><path d="M12 19a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M12 16a3 3 0 0 1-3-3"/></g>`
	wateringSoilInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 12h2m16 0h2M3 20.01l.01-.011M6 16.01l.01-.011M9 20.01l.01-.011m5.99.011l.01-.011M18 16.01l.01-.011M21 20.01l.01-.011M12.396 3.396L15.5 6.5a4.95 4.95 0 1 1-7 0l3.104-3.104a.56.56 0 0 1 .792 0Z"/>`
	webWindowInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 17V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 8h1"/></g>`
	webWindowCloseInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 17V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m10 14.243l2.121-2.122m0 0L14.243 10m-2.122 2.121L10 10m2.121 2.121l2.122 2.122M6 8h1"/></g>`
	webWindowEnergyConsumptionInnerSVG        = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 17V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.667 9L10 12h4l-1.667 3M6 8h1"/></g>`
	webpFormatInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M13.5 15V9h2.4a.6.6 0 0 1 .6.6v.9A1.5 1.5 0 0 1 15 12v0"/><path stroke-linejoin="round" d="M13.5 15h2.4a.6.6 0 0 0 .6-.6v-.9A1.5 1.5 0 0 0 15 12v0h-1.5m6 3v-3m0 0V9h3v3h-3Zm-18-3v6L3 12l1.5 3V9m6 0h-3v6h3m-3-3h2"/></g>`
	weightInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 5h3.9a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h3.9"/><path d="m16.279 6.329l.205-1.23a.605.605 0 0 0 0-.198l-.206-1.23A2 2 0 0 0 14.307 2H9.694a2 2 0 0 0-1.973 1.671l-.205 1.23a.6.6 0 0 0 0 .198l.205 1.23A2 2 0 0 0 9.694 8h4.612a2 2 0 0 0 1.973-1.671ZM12 8l-1-2.5"/></g>`
	weightAltInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 5h3.9a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h3.9"/><path d="m16.279 6.329l.205-1.23a.605.605 0 0 0 0-.198l-.206-1.23A2 2 0 0 0 14.307 2H9.694a2 2 0 0 0-1.973 1.671l-.205 1.23a.6.6 0 0 0 0 .198l.205 1.23A2 2 0 0 0 9.694 8h4.612a2 2 0 0 0 1.973-1.671ZM12 8l-1-2.5M7 17h10"/></g>`
	whiteFlagInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 15l.95-10.454A.6.6 0 0 1 6.548 4h13.795a.6.6 0 0 1 .598.654l-.891 9.8a.6.6 0 0 1-.598.546H5Zm0 0l-.6 6"/>`
	wifiInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 19.51l.01-.011M2 8c6-4.5 14-4.5 20 0M5 12c4-3 10-3 14 0M8.5 15.5c2.25-1.4 4.75-1.4 7 0"/>`
	wifiErrorInnerSVG                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 18.51l.01-.011M2 7c6-4.5 14-4.5 20 0M5 11c4-3 10-3 14 0M8.5 14.5c2.25-1.4 4.75-1.4 7 0m1.621 6.864l2.122-2.121m2.121-2.122l-2.121 2.122m0 0L17.12 17.12m2.122 2.122l2.121 2.121"/>`
	wifiIssueInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2.126 8.324c-.2-.262-.155-.605.086-.79C5.29 5.179 8.552 4 11.999 4c3.447 0 6.71 1.178 9.788 3.535c.252.212.28.558.085.789l-9.455 11.173a.548.548 0 0 1-.836 0L2.126 8.324Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 8v2m0 4.01l.01-.011"/></g>`
	wifiOffInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 19.51l.01-.011M3 3l18 18M2 8a17.053 17.053 0 0 1 3.757-2.14M22 8c-3.572-2.68-7.854-3.763-12-3.252M5 12c1.333-1 2.889-1.667 4.518-2M19 12a11.274 11.274 0 0 0-4.282-1.95M8.5 15.5c2.25-1.4 4.75-1.4 7 0"/>`
	wifiSignalNoneInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M2.126 8.324c-.2-.262-.155-.605.086-.79C5.29 5.179 8.552 4 11.999 4c3.447 0 6.71 1.178 9.788 3.535c.252.212.28.558.085.789l-9.455 11.173a.548.548 0 0 1-.836 0L2.126 8.324Z"/>`
	wifiTagInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 14.76l.01-.011M7 11.25c2.5-3 7.5-3 10 0m-8 2c1.5-2 4.5-2 6 0"/></g>`
	windInnerSVG                              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.279 7C19.782 7 21 8.12 21 9.5S19.782 12 18.279 12H3m14.938 8c1.139 0 2.562-.5 2.562-2.5S19.077 15 17.937 15H3m7.412-11C11.842 4 13 5.12 13 6.5S11.841 9 10.412 9H3"/>`
	windowsInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-width="1.5" d="M4 16.986V7.014a.6.6 0 0 1 .507-.593l14.8-2.313a.6.6 0 0 1 .693.593v14.598a.6.6 0 0 1-.693.593l-14.8-2.313A.6.6 0 0 1 4 16.986ZM4 12h16m-9.5-6.5v13"/>`
	womenTShirtInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 21H6s1.66-4.825 1.5-8c-.1-1.989-1.524-3.079-1-5c.23-.842 1-2 1-2S9 7 12 7s4.5-1 4.5-1s.77 1.158 1 2c.524 1.921-.9 3.011-1 5c-.16 3.175 1.5 8 1.5 8ZM7.5 6V3m9 3V3"/>`
	wrapTextInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 7h16M4 17h5m-5-5h13.5a2.5 2.5 0 0 1 2.5 2.5v0a2.5 2.5 0 0 1-2.5 2.5h-5"/><path d="M15 15.5L12.5 17l2.5 1.5v-3Z"/></g>`
	wrenchInnerSVG                            = `<g fill="none" stroke-width="1.5"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" clip-path="url(#iconoirWrench0)"><path d="m10.05 10.607l-7.07 7.07a2 2 0 0 0 0 2.83v0a2 2 0 0 0 2.828 0l7.07-7.072m-2.828-2.828c-.844-2.153-.679-4.978 1.06-6.718c1.74-1.74 4.95-2.121 6.718-1.06l-3.04 3.04l-.283 3.111l3.111-.282l3.04-3.041c1.062 1.768.68 4.978-1.06 6.717c-1.74 1.74-4.564 1.905-6.717 1.061"/></g><defs><clipPath id="iconoirWrench0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`
	wristwatchInnerSVG                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 16.472V20a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-3.528m0-8.944V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3.528"/><path d="M18 12a6 6 0 1 0-12 0a6 6 0 0 0 12 0Z"/><path d="M14 12h-2v-2"/></g>`
	wwwInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.338 17A9.996 9.996 0 0 0 12 22a9.996 9.996 0 0 0 8.662-5M3.338 7A9.996 9.996 0 0 1 12 2a9.996 9.996 0 0 1 8.662 5"/><path d="M13 21.95s1.408-1.853 2.295-4.95M13 2.05S14.408 3.902 15.295 7M11 21.95S9.592 20.098 8.705 17M11 2.05S9.592 3.902 8.705 7M9 10l1.5 5l1.5-5l1.5 5l1.5-5M1 10l1.5 5L4 10l1.5 5L7 10m10 0l1.5 5l1.5-5l1.5 5l1.5-5"/></g>`
	xCoordinateInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6ZM10 8l4 8m0-8l-4 8"/>`
	xboxAInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="m15 16l-3-8l-3 8m5-2h-4"/></g>`
	xboxBInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M12.599 11.826c2.535 0 2.535 4.174 0 4.174H9.5v-4.174m3.099 0H9.5m3.099 0c2.535 0 2.535-3.826 0-3.826H9.5v3.826"/></g>`
	xboxXInnerSVG                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm3-6L9 8m0 8l6-8"/>`
	xboxYInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10ZM9 8l3 5"/><path d="M12 16v-3l3-5"/></g>`
	xrayViewInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524Z"/><path d="m20.5 16.722l-8.209-4.56a.6.6 0 0 0-.582 0L3.5 16.722m.028-9.428l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21V3"/></g>`
	yCoordinateInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6ZM10 8l2 4"/><path d="m14 8l-2 4v4"/></g>`
	yenInnerSVG                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h12M6 4l6 8m6-8l-6 8m0 0v8m-6-4h12"/>`
	yenSquareInnerSVG                         = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 13h8M8 7l4 5.5M16 7l-4 5.5m0 0V18m-4-3h8"/></g>`
	yogaInnerSVG                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.571 15.004l.858 1.845s3.857.819 3.857 2.767C19.286 21 17.57 21 17.57 21H13l-2.25-1.25"/><path d="m9.429 15.004l-.857 1.845s-3.858.819-3.858 2.767C4.714 21 6.43 21 6.43 21H8.5l2.25-1.25L13.5 18"/><path d="M3 15.926s2.143-.461 3.429-.922C7.714 8.546 11.57 9.007 12 9.007c.429 0 4.286-.461 5.571 5.997c1.286.46 3.429.922 3.429.922M12 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`
	youtubeInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-width="1.5"><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14 12l-3.5 2v-4l3.5 2Z"/><path d="M2 12.707v-1.415c0-2.895 0-4.343.905-5.274c.906-.932 2.332-.972 5.183-1.053C9.438 4.927 10.818 4.9 12 4.9c1.181 0 2.561.027 3.912.065c2.851.081 4.277.121 5.182 1.053c.906.931.906 2.38.906 5.274v1.415c0 2.896 0 4.343-.905 5.275c-.906.931-2.331.972-5.183 1.052c-1.35.039-2.73.066-3.912.066a141.1 141.1 0 0 1-3.912-.066c-2.851-.08-4.277-.12-5.183-1.052C2 17.05 2 15.602 2 12.708Z"/></g>`
	zCoordinateInnerSVG                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6Z"/><path d="M10 8h4l-4 8h4"/></g>`
	zoomInInnerSVG                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 11h3m3 0h-3m0 0V8m0 3v3m6 3l4 4M3 11a8 8 0 1 0 16 0a8 8 0 0 0-16 0Z"/>`
	zoomOutInnerSVG                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 17l4 4M3 11a8 8 0 1 0 16 0a8 8 0 0 0-16 0Zm5 0h6"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Accessibility(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accessibilityInnerSVG).
		Element(children...)
}

func AccessibilitySign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accessibilitySignInnerSVG).
		Element(children...)
}

func AccessibilityTech(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accessibilityTechInnerSVG).
		Element(children...)
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

func AddCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addCircleInnerSVG).
		Element(children...)
}

func AddDatabaseScript(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addDatabaseScriptInnerSVG).
		Element(children...)
}

func AddFolder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addFolderInnerSVG).
		Element(children...)
}

func AddFrame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addFrameInnerSVG).
		Element(children...)
}

func AddHexagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addHexagonInnerSVG).
		Element(children...)
}

func AddKeyframe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addKeyframeInnerSVG).
		Element(children...)
}

func AddKeyframeAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addKeyframeAltInnerSVG).
		Element(children...)
}

func AddKeyframes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addKeyframesInnerSVG).
		Element(children...)
}

func AddLens(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addLensInnerSVG).
		Element(children...)
}

func AddMediaImage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addMediaImageInnerSVG).
		Element(children...)
}

func AddMediaVideo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addMediaVideoInnerSVG).
		Element(children...)
}

func AddPage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addPageInnerSVG).
		Element(children...)
}

func AddPageAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addPageAltInnerSVG).
		Element(children...)
}

func AddPinAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addPinAltInnerSVG).
		Element(children...)
}

func AddSelection(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addSelectionInnerSVG).
		Element(children...)
}

func AddSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addSquareInnerSVG).
		Element(children...)
}

func AddToCart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addToCartInnerSVG).
		Element(children...)
}

func AddUser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addUserInnerSVG).
		Element(children...)
}

func AfricanTree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(africanTreeInnerSVG).
		Element(children...)
}

func Agile(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(agileInnerSVG).
		Element(children...)
}

func AirConditioner(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airConditionerInnerSVG).
		Element(children...)
}

func Airplane(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplaneInnerSVG).
		Element(children...)
}

func AirplaneHelix(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplaneHelixInnerSVG).
		Element(children...)
}

func AirplaneHelixFortyFiveDeg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplaneHelixFortyFiveDegInnerSVG).
		Element(children...)
}

func AirplaneOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplaneOffInnerSVG).
		Element(children...)
}

func AirplaneRotation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplaneRotationInnerSVG).
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

func Alarm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alarmInnerSVG).
		Element(children...)
}

func Album(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(albumInnerSVG).
		Element(children...)
}

func AlbumCarousel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(albumCarouselInnerSVG).
		Element(children...)
}

func AlbumList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(albumListInnerSVG).
		Element(children...)
}

func AlbumOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(albumOpenInnerSVG).
		Element(children...)
}

func AlignBottomBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignBottomBoxInnerSVG).
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

func AlignLeftBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignLeftBoxInnerSVG).
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

func AlignRightBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignRightBoxInnerSVG).
		Element(children...)
}

func AlignTopBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTopBoxInnerSVG).
		Element(children...)
}

func AngleTool(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleToolInnerSVG).
		Element(children...)
}

func Antenna(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(antennaInnerSVG).
		Element(children...)
}

func AntennaOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(antennaOffInnerSVG).
		Element(children...)
}

func AntennaSignal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(antennaSignalInnerSVG).
		Element(children...)
}

func AntennaSignalTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(antennaSignalTagInnerSVG).
		Element(children...)
}

func AppNotification(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appNotificationInnerSVG).
		Element(children...)
}

func AppWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appWindowInnerSVG).
		Element(children...)
}

func Apple(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleInnerSVG).
		Element(children...)
}

func AppleHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleHalfInnerSVG).
		Element(children...)
}

func AppleHalfAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleHalfAltInnerSVG).
		Element(children...)
}

func AppleImacTwoThousandTwentyOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleImacTwoThousandTwentyOneInnerSVG).
		Element(children...)
}

func AppleImacTwoThousandTwentyOneSide(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleImacTwoThousandTwentyOneSideInnerSVG).
		Element(children...)
}

func AppleMac(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleMacInnerSVG).
		Element(children...)
}

func AppleSwift(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleSwiftInnerSVG).
		Element(children...)
}

func AppleWallet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleWalletInnerSVG).
		Element(children...)
}

func ArSymbol(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arSymbolInnerSVG).
		Element(children...)
}

func Arcade(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arcadeInnerSVG).
		Element(children...)
}

func Archery(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archeryInnerSVG).
		Element(children...)
}

func ArcheryMatch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archeryMatchInnerSVG).
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

func AreaSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(areaSearchInnerSVG).
		Element(children...)
}

func ArrowArchery(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowArcheryInnerSVG).
		Element(children...)
}

func ArrowBlCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBlCircleInnerSVG).
		Element(children...)
}

func ArrowBlSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBlSquareInnerSVG).
		Element(children...)
}

func ArrowBrCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBrCircleInnerSVG).
		Element(children...)
}

func ArrowBrSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBrSquareInnerSVG).
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

func ArrowEmailForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowEmailForwardInnerSVG).
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

func ArrowSeparate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowSeparateInnerSVG).
		Element(children...)
}

func ArrowSeparateVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowSeparateVerticalInnerSVG).
		Element(children...)
}

func ArrowTlCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTlCircleInnerSVG).
		Element(children...)
}

func ArrowTlSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTlSquareInnerSVG).
		Element(children...)
}

func ArrowTrCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTrCircleInnerSVG).
		Element(children...)
}

func ArrowTrSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTrSquareInnerSVG).
		Element(children...)
}

func ArrowUnion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUnionInnerSVG).
		Element(children...)
}

func ArrowUnionVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUnionVerticalInnerSVG).
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

func Asana(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(asanaInnerSVG).
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

func AtSignCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atSignCircleInnerSVG).
		Element(children...)
}

func Atom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atomInnerSVG).
		Element(children...)
}

func Attachment(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(attachmentInnerSVG).
		Element(children...)
}

func AugmentedReality(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(augmentedRealityInnerSVG).
		Element(children...)
}

func AutoFlash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(autoFlashInnerSVG).
		Element(children...)
}

func AviFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aviFormatInnerSVG).
		Element(children...)
}

func Axes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(axesInnerSVG).
		Element(children...)
}

func BackwardFifteenSeconds(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backwardFifteenSecondsInnerSVG).
		Element(children...)
}

func Bag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagInnerSVG).
		Element(children...)
}

func Bank(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bankInnerSVG).
		Element(children...)
}

func Barcode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barcodeInnerSVG).
		Element(children...)
}

func Basketball(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketballInnerSVG).
		Element(children...)
}

func BasketballAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketballAltInnerSVG).
		Element(children...)
}

func BasketballField(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketballFieldInnerSVG).
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

func BatteryEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryEmptyInnerSVG).
		Element(children...)
}

func BatteryFifty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFiftyInnerSVG).
		Element(children...)
}

func BatteryFull(children ...*engine.UberElement) *engine.UberElement {
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

func BatteryIndicator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryIndicatorInnerSVG).
		Element(children...)
}

func BatterySeventyFive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batterySeventyFiveInnerSVG).
		Element(children...)
}

func BatteryTwentyFive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryTwentyFiveInnerSVG).
		Element(children...)
}

func BatteryWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryWarningInnerSVG).
		Element(children...)
}

func Bbq(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bbqInnerSVG).
		Element(children...)
}

func BeachBag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(beachBagInnerSVG).
		Element(children...)
}

func BeachBagBig(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(beachBagBigInnerSVG).
		Element(children...)
}

func Bed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bedInnerSVG).
		Element(children...)
}

func BedReady(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bedReadyInnerSVG).
		Element(children...)
}

func Behance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(behanceInnerSVG).
		Element(children...)
}

func BehanceTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(behanceTagInnerSVG).
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

func BellNotification(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellNotificationInnerSVG).
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

func Bicycle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bicycleInnerSVG).
		Element(children...)
}

func Bin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binInnerSVG).
		Element(children...)
}

func BinAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binAddInnerSVG).
		Element(children...)
}

func BinFull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binFullInnerSVG).
		Element(children...)
}

func BinHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binHalfInnerSVG).
		Element(children...)
}

func BinMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binMinusInnerSVG).
		Element(children...)
}

func Bishop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bishopInnerSVG).
		Element(children...)
}

func Bitbucket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bitbucketInnerSVG).
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

func BluetoothTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothTagInnerSVG).
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

func BoldSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boldSquareInnerSVG).
		Element(children...)
}

func Bonfire(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bonfireInnerSVG).
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

func BookStack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookStackInnerSVG).
		Element(children...)
}

func BookmarkBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkBookInnerSVG).
		Element(children...)
}

func BookmarkCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkCircleInnerSVG).
		Element(children...)
}

func BookmarkEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkEmptyInnerSVG).
		Element(children...)
}

func BorderBl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBlInnerSVG).
		Element(children...)
}

func BorderBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBottomInnerSVG).
		Element(children...)
}

func BorderBr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBrInnerSVG).
		Element(children...)
}

func BorderInner(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderInnerInnerSVG).
		Element(children...)
}

func BorderLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderLeftInnerSVG).
		Element(children...)
}

func BorderOut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderOutInnerSVG).
		Element(children...)
}

func BorderRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderRightInnerSVG).
		Element(children...)
}

func BorderTl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTlInnerSVG).
		Element(children...)
}

func BorderTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopInnerSVG).
		Element(children...)
}

func BorderTr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTrInnerSVG).
		Element(children...)
}

func BounceLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bounceLeftInnerSVG).
		Element(children...)
}

func BounceRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bounceRightInnerSVG).
		Element(children...)
}

func BowlingBall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bowlingBallInnerSVG).
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

func BoxIso(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxIsoInnerSVG).
		Element(children...)
}

func BoxingGlove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxingGloveInnerSVG).
		Element(children...)
}

func Brain(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brainInnerSVG).
		Element(children...)
}

func BrainElectricity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brainElectricityInnerSVG).
		Element(children...)
}

func BrainResearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brainResearchInnerSVG).
		Element(children...)
}

func BrainWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brainWarningInnerSVG).
		Element(children...)
}

func BreadSlice(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(breadSliceInnerSVG).
		Element(children...)
}

func BrightCrown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brightCrownInnerSVG).
		Element(children...)
}

func BrightStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brightStarInnerSVG).
		Element(children...)
}

func Brightness(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brightnessInnerSVG).
		Element(children...)
}

func BrightnessWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brightnessWindowInnerSVG).
		Element(children...)
}

func BubbleDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bubbleDownloadInnerSVG).
		Element(children...)
}

func BubbleError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bubbleErrorInnerSVG).
		Element(children...)
}

func BubbleIncome(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bubbleIncomeInnerSVG).
		Element(children...)
}

func BubbleOutcome(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bubbleOutcomeInnerSVG).
		Element(children...)
}

func BubbleSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bubbleSearchInnerSVG).
		Element(children...)
}

func BubbleStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bubbleStarInnerSVG).
		Element(children...)
}

func BubbleUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bubbleUploadInnerSVG).
		Element(children...)
}

func BubbleWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bubbleWarningInnerSVG).
		Element(children...)
}

func Building(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buildingInnerSVG).
		Element(children...)
}

func Bus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(busInnerSVG).
		Element(children...)
}

func BusStop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(busStopInnerSVG).
		Element(children...)
}

func CableTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cableTagInnerSVG).
		Element(children...)
}

func Calculator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
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
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarInnerSVG).
		Element(children...)
}

func CalendarMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMinusInnerSVG).
		Element(children...)
}

func CalendarPlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarPlusInnerSVG).
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

func Cancel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cancelInnerSVG).
		Element(children...)
}

func CandlestickChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(candlestickChartInnerSVG).
		Element(children...)
}

func Car(children ...*engine.UberElement) *engine.UberElement {
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

func Carbon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(carbonInnerSVG).
		Element(children...)
}

func CardIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardIssueInnerSVG).
		Element(children...)
}

func CardLocked(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardLockedInnerSVG).
		Element(children...)
}

func CardSecurity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardSecurityInnerSVG).
		Element(children...)
}

func CardWallet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardWalletInnerSVG).
		Element(children...)
}

func Cart(children ...*engine.UberElement) *engine.UberElement {
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

func CartAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartAltInnerSVG).
		Element(children...)
}

func Cash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cashInnerSVG).
		Element(children...)
}

func CenterAlign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(centerAlignInnerSVG).
		Element(children...)
}

func ChatAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatAddInnerSVG).
		Element(children...)
}

func ChatBubble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleInnerSVG).
		Element(children...)
}

func ChatBubbleCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleCheckInnerSVG).
		Element(children...)
}

func ChatBubbleCheckOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleCheckOneInnerSVG).
		Element(children...)
}

func ChatBubbleEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleEmptyInnerSVG).
		Element(children...)
}

func ChatBubbleError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleErrorInnerSVG).
		Element(children...)
}

func ChatBubbleQuestion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleQuestionInnerSVG).
		Element(children...)
}

func ChatBubbleTranslate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleTranslateInnerSVG).
		Element(children...)
}

func ChatBubbleWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleWarningInnerSVG).
		Element(children...)
}

func ChatLines(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatLinesInnerSVG).
		Element(children...)
}

func ChatRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatRemoveInnerSVG).
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

func CheckWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkWindowInnerSVG).
		Element(children...)
}

func Chocolate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chocolateInnerSVG).
		Element(children...)
}

func Chromecast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chromecastInnerSVG).
		Element(children...)
}

func ChromecastActive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chromecastActiveInnerSVG).
		Element(children...)
}

func Church(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(churchInnerSVG).
		Element(children...)
}

func ChurchAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(churchAltInnerSVG).
		Element(children...)
}

func CinemaOld(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cinemaOldInnerSVG).
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

func City(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cityInnerSVG).
		Element(children...)
}

func CleanWater(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cleanWaterInnerSVG).
		Element(children...)
}

func ClipboardCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardCheckInnerSVG).
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

func ClosedCaptions(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closedCaptionsInnerSVG).
		Element(children...)
}

func Closet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closetInnerSVG).
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

func CloudBookAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudBookAltInnerSVG).
		Element(children...)
}

func CloudCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudCheckInnerSVG).
		Element(children...)
}

func CloudDesync(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDesyncInnerSVG).
		Element(children...)
}

func CloudDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDownloadInnerSVG).
		Element(children...)
}

func CloudError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudErrorInnerSVG).
		Element(children...)
}

func CloudSunny(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudSunnyInnerSVG).
		Element(children...)
}

func CloudSync(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudSyncInnerSVG).
		Element(children...)
}

func CloudUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudUploadInnerSVG).
		Element(children...)
}

func Clutery(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cluteryInnerSVG).
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

func CodeBrackets(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeBracketsInnerSVG).
		Element(children...)
}

func CodeBracketsSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeBracketsSquareInnerSVG).
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

func CoffeeCup(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coffeeCupInnerSVG).
		Element(children...)
}

func Coin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coinInnerSVG).
		Element(children...)
}

func CollageFrame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(collageFrameInnerSVG).
		Element(children...)
}

func Collapse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(collapseInnerSVG).
		Element(children...)
}

func ColorFilter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(colorFilterInnerSVG).
		Element(children...)
}

func ColorPicker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(colorPickerInnerSVG).
		Element(children...)
}

func ColorPickerEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(colorPickerEmptyInnerSVG).
		Element(children...)
}

func Combine(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(combineInnerSVG).
		Element(children...)
}

func Community(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(communityInnerSVG).
		Element(children...)
}

func CompactDisc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compactDiscInnerSVG).
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

func Compress(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compressInnerSVG).
		Element(children...)
}

func CompressLines(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compressLinesInnerSVG).
		Element(children...)
}

func Computer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(computerInnerSVG).
		Element(children...)
}

func Consumable(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(consumableInnerSVG).
		Element(children...)
}

func ControlSlider(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(controlSliderInnerSVG).
		Element(children...)
}

func Cookie(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cookieInnerSVG).
		Element(children...)
}

func Cooling(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coolingInnerSVG).
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

func Copyright(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(copyrightInnerSVG).
		Element(children...)
}

func CornerBottomLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerBottomLeftInnerSVG).
		Element(children...)
}

func CornerBottomRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerBottomRightInnerSVG).
		Element(children...)
}

func CornerTopLeft(children ...*engine.UberElement) *engine.UberElement {
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

func CornerTopRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerTopRightInnerSVG).
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

func CpuWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cpuWarningInnerSVG).
		Element(children...)
}

func CrackedEgg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crackedEggInnerSVG).
		Element(children...)
}

func CreativeCommons(children ...*engine.UberElement) *engine.UberElement {
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

func CreditCardTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardTwoInnerSVG).
		Element(children...)
}

func CreditCards(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardsInnerSVG).
		Element(children...)
}

func Crib(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cribInnerSVG).
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

func CropRotateBl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropRotateBlInnerSVG).
		Element(children...)
}

func CropRotateBr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropRotateBrInnerSVG).
		Element(children...)
}

func CropRotateTl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropRotateTlInnerSVG).
		Element(children...)
}

func CropRotateTr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropRotateTrInnerSVG).
		Element(children...)
}

func Crown(children ...*engine.UberElement) *engine.UberElement {
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

func CrownCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crownCircleInnerSVG).
		Element(children...)
}

func CssThree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cssThreeInnerSVG).
		Element(children...)
}

func CursorPointer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cursorPointerInnerSVG).
		Element(children...)
}

func Cut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cutInnerSVG).
		Element(children...)
}

func CutAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cutAltInnerSVG).
		Element(children...)
}

func Cycling(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cyclingInnerSVG).
		Element(children...)
}

func Cylinder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cylinderInnerSVG).
		Element(children...)
}

func DashFlag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashFlagInnerSVG).
		Element(children...)
}

func Dashboard(children ...*engine.UberElement) *engine.UberElement {
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

func DashboardDots(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashboardDotsInnerSVG).
		Element(children...)
}

func DashboardSpeed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashboardSpeedInnerSVG).
		Element(children...)
}

func DataTransferBoth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dataTransferBothInnerSVG).
		Element(children...)
}

func DataTransferCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dataTransferCheckInnerSVG).
		Element(children...)
}

func DataTransferDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dataTransferDownInnerSVG).
		Element(children...)
}

func DataTransferUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dataTransferUpInnerSVG).
		Element(children...)
}

func DataTransferWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dataTransferWarningInnerSVG).
		Element(children...)
}

func DatabaseBackup(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseBackupInnerSVG).
		Element(children...)
}

func DatabaseExport(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseExportInnerSVG).
		Element(children...)
}

func DatabaseMonitor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseMonitorInnerSVG).
		Element(children...)
}

func DatabaseRestore(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseRestoreInnerSVG).
		Element(children...)
}

func DatabaseScript(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseScriptInnerSVG).
		Element(children...)
}

func DatabaseSettings(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseSettingsInnerSVG).
		Element(children...)
}

func DatabaseStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseStarInnerSVG).
		Element(children...)
}

func DatabaseStats(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseStatsInnerSVG).
		Element(children...)
}

func DatabaseTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseTagInnerSVG).
		Element(children...)
}

func Db(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dbInnerSVG).
		Element(children...)
}

func DbCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dbCheckInnerSVG).
		Element(children...)
}

func DbError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dbErrorInnerSVG).
		Element(children...)
}

func DbSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dbSearchInnerSVG).
		Element(children...)
}

func DbStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dbStarInnerSVG).
		Element(children...)
}

func DbWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dbWarningInnerSVG).
		Element(children...)
}

func DeCompress(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deCompressInnerSVG).
		Element(children...)
}

func DeleteCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deleteCircleInnerSVG).
		Element(children...)
}

func Delivery(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deliveryInnerSVG).
		Element(children...)
}

func DeliveryTruck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deliveryTruckInnerSVG).
		Element(children...)
}

func Depth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(depthInnerSVG).
		Element(children...)
}

func DesignNib(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(designNibInnerSVG).
		Element(children...)
}

func DesignPencil(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(designPencilInnerSVG).
		Element(children...)
}

func Desk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deskInnerSVG).
		Element(children...)
}

func Dialpad(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dialpadInnerSVG).
		Element(children...)
}

func Diameter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diameterInnerSVG).
		Element(children...)
}

func DiceFive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diceFiveInnerSVG).
		Element(children...)
}

func DiceFour(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diceFourInnerSVG).
		Element(children...)
}

func DiceOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diceOneInnerSVG).
		Element(children...)
}

func DiceSix(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diceSixInnerSVG).
		Element(children...)
}

func DiceThree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diceThreeInnerSVG).
		Element(children...)
}

func DiceTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diceTwoInnerSVG).
		Element(children...)
}

func DimmerSwitch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dimmerSwitchInnerSVG).
		Element(children...)
}

func DirectorChair(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(directorChairInnerSVG).
		Element(children...)
}

func Discord(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discordInnerSVG).
		Element(children...)
}

func Dishwasher(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dishwasherInnerSVG).
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

func DivideSelectionOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(divideSelectionOneInnerSVG).
		Element(children...)
}

func DivideSelectionTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(divideSelectionTwoInnerSVG).
		Element(children...)
}

func DivideThree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(divideThreeInnerSVG).
		Element(children...)
}

func Dna(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dnaInnerSVG).
		Element(children...)
}

func DocSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(docSearchInnerSVG).
		Element(children...)
}

func DocSearchAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(docSearchAltInnerSVG).
		Element(children...)
}

func DocStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(docStarInnerSVG).
		Element(children...)
}

func DocStarAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(docStarAltInnerSVG).
		Element(children...)
}

func Dollar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dollarInnerSVG).
		Element(children...)
}

func DomoticIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(domoticIssueInnerSVG).
		Element(children...)
}

func Donate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(donateInnerSVG).
		Element(children...)
}

func DoubleCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCheckInnerSVG).
		Element(children...)
}

func DownRoundArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downRoundArrowInnerSVG).
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

func DownloadCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadCircleInnerSVG).
		Element(children...)
}

func DownloadDataWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadDataWindowInnerSVG).
		Element(children...)
}

func DownloadSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadSquareInnerSVG).
		Element(children...)
}

func Drag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragInnerSVG).
		Element(children...)
}

func DragHandGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragHandGestureInnerSVG).
		Element(children...)
}

func Drawer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(drawerInnerSVG).
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

func Drone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneInnerSVG).
		Element(children...)
}

func DroneChargeFull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneChargeFullInnerSVG).
		Element(children...)
}

func DroneChargeHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneChargeHalfInnerSVG).
		Element(children...)
}

func DroneChargeLow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneChargeLowInnerSVG).
		Element(children...)
}

func DroneCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneCheckInnerSVG).
		Element(children...)
}

func DroneError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneErrorInnerSVG).
		Element(children...)
}

func DroneLanding(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneLandingInnerSVG).
		Element(children...)
}

func DroneRefresh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneRefreshInnerSVG).
		Element(children...)
}

func DroneTakeOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(droneTakeOffInnerSVG).
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

func DropletHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropletHalfInnerSVG).
		Element(children...)
}

func EaseCurveControlPoints(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(easeCurveControlPointsInnerSVG).
		Element(children...)
}

func EaseIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(easeInInnerSVG).
		Element(children...)
}

func EaseInControlPoint(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(easeInControlPointInnerSVG).
		Element(children...)
}

func EaseInOut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(easeInOutInnerSVG).
		Element(children...)
}

func EaseOut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(easeOutInnerSVG).
		Element(children...)
}

func EaseOutControlPoint(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(easeOutControlPointInnerSVG).
		Element(children...)
}

func EcologyBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ecologyBookInnerSVG).
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

func EditPencil(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editPencilInnerSVG).
		Element(children...)
}

func Egg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eggInnerSVG).
		Element(children...)
}

func Eject(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ejectInnerSVG).
		Element(children...)
}

func ElectronicsChip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(electronicsChipInnerSVG).
		Element(children...)
}

func ElectronicsTransistor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(electronicsTransistorInnerSVG).
		Element(children...)
}

func Emoji(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiInnerSVG).
		Element(children...)
}

func EmojiBall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiBallInnerSVG).
		Element(children...)
}

func EmojiBlinkLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiBlinkLeftInnerSVG).
		Element(children...)
}

func EmojiBlinkRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiBlinkRightInnerSVG).
		Element(children...)
}

func EmojiLookDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiLookDownInnerSVG).
		Element(children...)
}

func EmojiLookLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiLookLeftInnerSVG).
		Element(children...)
}

func EmojiLookRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiLookRightInnerSVG).
		Element(children...)
}

func EmojiLookUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiLookUpInnerSVG).
		Element(children...)
}

func EmojiPuzzled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiPuzzledInnerSVG).
		Element(children...)
}

func EmojiQuite(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiQuiteInnerSVG).
		Element(children...)
}

func EmojiReally(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiReallyInnerSVG).
		Element(children...)
}

func EmojiSad(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiSadInnerSVG).
		Element(children...)
}

func EmojiSatisfied(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiSatisfiedInnerSVG).
		Element(children...)
}

func EmojiSingLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiSingLeftInnerSVG).
		Element(children...)
}

func EmojiSingLeftNote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiSingLeftNoteInnerSVG).
		Element(children...)
}

func EmojiSingRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiSingRightInnerSVG).
		Element(children...)
}

func EmojiSingRightNote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiSingRightNoteInnerSVG).
		Element(children...)
}

func EmojiSurprise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiSurpriseInnerSVG).
		Element(children...)
}

func EmojiSurpriseAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiSurpriseAltInnerSVG).
		Element(children...)
}

func EmojiTalkingAngry(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiTalkingAngryInnerSVG).
		Element(children...)
}

func EmojiTalkingHappy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiTalkingHappyInnerSVG).
		Element(children...)
}

func EmojiThinkLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiThinkLeftInnerSVG).
		Element(children...)
}

func EmojiThinkRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiThinkRightInnerSVG).
		Element(children...)
}

func EmptyPage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emptyPageInnerSVG).
		Element(children...)
}

func EnergyUsageWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(energyUsageWindowInnerSVG).
		Element(children...)
}

func Enlarge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enlargeInnerSVG).
		Element(children...)
}

func EnlargeRoundArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enlargeRoundArrowInnerSVG).
		Element(children...)
}

func Erase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eraseInnerSVG).
		Element(children...)
}

func ErrorWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(errorWindowInnerSVG).
		Element(children...)
}

func Euro(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(euroInnerSVG).
		Element(children...)
}

func EuroSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(euroSquareInnerSVG).
		Element(children...)
}

func EvCharge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(evChargeInnerSVG).
		Element(children...)
}

func EvChargeAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(evChargeAltInnerSVG).
		Element(children...)
}

func EvPlug(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(evPlugInnerSVG).
		Element(children...)
}

func EvPlugCharging(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(evPlugChargingInnerSVG).
		Element(children...)
}

func EvPlugError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(evPlugErrorInnerSVG).
		Element(children...)
}

func EvStation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(evStationInnerSVG).
		Element(children...)
}

func EvTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(evTagInnerSVG).
		Element(children...)
}

func Exclude(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(excludeInnerSVG).
		Element(children...)
}

func Expand(children ...*engine.UberElement) *engine.UberElement {
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

func ExpandLines(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandLinesInnerSVG).
		Element(children...)
}

func EyeAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeAltInnerSVG).
		Element(children...)
}

func EyeClose(children ...*engine.UberElement) *engine.UberElement {
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

func EyeEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeEmptyInnerSVG).
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

func FaceId(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceIdInnerSVG).
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

func FacebookTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookTagInnerSVG).
		Element(children...)
}

func Facetime(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facetimeInnerSVG).
		Element(children...)
}

func Farm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(farmInnerSVG).
		Element(children...)
}

func FastArrowDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastArrowDownInnerSVG).
		Element(children...)
}

func FastArrowDownBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastArrowDownBoxInnerSVG).
		Element(children...)
}

func FastArrowLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastArrowLeftInnerSVG).
		Element(children...)
}

func FastArrowLeftBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastArrowLeftBoxInnerSVG).
		Element(children...)
}

func FastArrowRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastArrowRightInnerSVG).
		Element(children...)
}

func FastArrowRightBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastArrowRightBoxInnerSVG).
		Element(children...)
}

func FastArrowUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastArrowUpInnerSVG).
		Element(children...)
}

func FastArrowUpBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastArrowUpBoxInnerSVG).
		Element(children...)
}

func FastDownCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastDownCircleInnerSVG).
		Element(children...)
}

func FastLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastLeftCircleInnerSVG).
		Element(children...)
}

func FastRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastRightCircleInnerSVG).
		Element(children...)
}

func FastUpCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastUpCircleInnerSVG).
		Element(children...)
}

func FavouriteBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(favouriteBookInnerSVG).
		Element(children...)
}

func FavouriteWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(favouriteWindowInnerSVG).
		Element(children...)
}

func Female(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(femaleInnerSVG).
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

func FileNotFound(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileNotFoundInnerSVG).
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

func FilterAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterAltInnerSVG).
		Element(children...)
}

func Finder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(finderInnerSVG).
		Element(children...)
}

func Fingerprint(children ...*engine.UberElement) *engine.UberElement {
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

func FingerprintCheckCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintCheckCircleInnerSVG).
		Element(children...)
}

func FingerprintCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintCircleInnerSVG).
		Element(children...)
}

func FingerprintErrorCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintErrorCircleInnerSVG).
		Element(children...)
}

func FingerprintLockCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintLockCircleInnerSVG).
		Element(children...)
}

func FingerprintPhone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintPhoneInnerSVG).
		Element(children...)
}

func FingerprintScan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintScanInnerSVG).
		Element(children...)
}

func FingerprintSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintSquareInnerSVG).
		Element(children...)
}

func FingerprintWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintWindowInnerSVG).
		Element(children...)
}

func FireFlame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fireFlameInnerSVG).
		Element(children...)
}

func Fishing(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fishingInnerSVG).
		Element(children...)
}

func Flare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flareInnerSVG).
		Element(children...)
}

func Flash(children ...*engine.UberElement) *engine.UberElement {
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

func FlashOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flashOffInnerSVG).
		Element(children...)
}

func Flask(children ...*engine.UberElement) *engine.UberElement {
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

func Flip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipInnerSVG).
		Element(children...)
}

func FlipReverse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipReverseInnerSVG).
		Element(children...)
}

func Flower(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flowerInnerSVG).
		Element(children...)
}

func Fluorine(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fluorineInnerSVG).
		Element(children...)
}

func Fog(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fogInnerSVG).
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

func FolderAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderAlertInnerSVG).
		Element(children...)
}

func FolderSettings(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderSettingsInnerSVG).
		Element(children...)
}

func FontSize(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fontSizeInnerSVG).
		Element(children...)
}

func Football(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(footballInnerSVG).
		Element(children...)
}

func FootballBall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(footballBallInnerSVG).
		Element(children...)
}

func Forward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardInnerSVG).
		Element(children...)
}

func ForwardFifteenSeconds(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardFifteenSecondsInnerSVG).
		Element(children...)
}

func ForwardMessage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardMessageInnerSVG).
		Element(children...)
}

func FourKDisplay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fourKDisplayInnerSVG).
		Element(children...)
}

func Frame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameInnerSVG).
		Element(children...)
}

func FrameAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameAltInnerSVG).
		Element(children...)
}

func FrameAltEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameAltEmptyInnerSVG).
		Element(children...)
}

func FrameSelect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameSelectInnerSVG).
		Element(children...)
}

func FrameSimple(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameSimpleInnerSVG).
		Element(children...)
}

func FrameTool(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameToolInnerSVG).
		Element(children...)
}

func Fridge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fridgeInnerSVG).
		Element(children...)
}

func Fx(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fxInnerSVG).
		Element(children...)
}

func FxTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fxTagInnerSVG).
		Element(children...)
}

func Gamepad(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadInnerSVG).
		Element(children...)
}

func Garage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(garageInnerSVG).
		Element(children...)
}

func Gas(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gasInnerSVG).
		Element(children...)
}

func GasTank(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gasTankInnerSVG).
		Element(children...)
}

func GasTankDrop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gasTankDropInnerSVG).
		Element(children...)
}

func GifFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gifFormatInnerSVG).
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

func GitCherryPickCommit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCherryPickCommitInnerSVG).
		Element(children...)
}

func GitCommand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCommandInnerSVG).
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

func GitCompare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCompareInnerSVG).
		Element(children...)
}

func GitFork(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitForkInnerSVG).
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

func GitPullRequestClosed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitPullRequestClosedInnerSVG).
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

func GithubCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(githubCircleInnerSVG).
		Element(children...)
}

func GitlabFull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitlabFullInnerSVG).
		Element(children...)
}

func GlassEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(glassEmptyInnerSVG).
		Element(children...)
}

func GlassHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(glassHalfInnerSVG).
		Element(children...)
}

func GlassHalfAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(glassHalfAltInnerSVG).
		Element(children...)
}

func Glasses(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
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
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeInnerSVG).
		Element(children...)
}

func Golf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(golfInnerSVG).
		Element(children...)
}

func Google(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleInnerSVG).
		Element(children...)
}

func GoogleCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleCircleInnerSVG).
		Element(children...)
}

func GoogleDocs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleDocsInnerSVG).
		Element(children...)
}

func GoogleDrive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleDriveInnerSVG).
		Element(children...)
}

func GoogleDriveCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleDriveCheckInnerSVG).
		Element(children...)
}

func GoogleDriveSync(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleDriveSyncInnerSVG).
		Element(children...)
}

func GoogleDriveWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleDriveWarningInnerSVG).
		Element(children...)
}

func GoogleHome(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleHomeInnerSVG).
		Element(children...)
}

func GoogleOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleOneInnerSVG).
		Element(children...)
}

func Gps(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gpsInnerSVG).
		Element(children...)
}

func GraduationCap(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graduationCapInnerSVG).
		Element(children...)
}

func GraphDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphDownInnerSVG).
		Element(children...)
}

func GraphUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphUpInnerSVG).
		Element(children...)
}

func GreenBus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(greenBusInnerSVG).
		Element(children...)
}

func GreenTruck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(greenTruckInnerSVG).
		Element(children...)
}

func GreenVehicle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(greenVehicleInnerSVG).
		Element(children...)
}

func GridAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridAddInnerSVG).
		Element(children...)
}

func GridMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridMinusInnerSVG).
		Element(children...)
}

func GridRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridRemoveInnerSVG).
		Element(children...)
}

func Group(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(groupInnerSVG).
		Element(children...)
}

func Gym(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gymInnerSVG).
		Element(children...)
}

func HalfCookie(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(halfCookieInnerSVG).
		Element(children...)
}

func HalfMoon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(halfMoonInnerSVG).
		Element(children...)
}

func Hammer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hammerInnerSVG).
		Element(children...)
}

func HandBrake(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(handBrakeInnerSVG).
		Element(children...)
}

func Handbag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(handbagInnerSVG).
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

func Hat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hatInnerSVG).
		Element(children...)
}

func Hd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hdInnerSVG).
		Element(children...)
}

func HdDisplay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hdDisplayInnerSVG).
		Element(children...)
}

func Hdr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hdrInnerSVG).
		Element(children...)
}

func Headset(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headsetInnerSVG).
		Element(children...)
}

func HeadsetCharge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headsetChargeInnerSVG).
		Element(children...)
}

func HeadsetHelp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headsetHelpInnerSVG).
		Element(children...)
}

func HeadsetIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headsetIssueInnerSVG).
		Element(children...)
}

func HealthShield(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(healthShieldInnerSVG).
		Element(children...)
}

func Healthcare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(healthcareInnerSVG).
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

func Heating(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heatingInnerSVG).
		Element(children...)
}

func HeavyRain(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heavyRainInnerSVG).
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

func HelpSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(helpSquareInnerSVG).
		Element(children...)
}

func Heptagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heptagonInnerSVG).
		Element(children...)
}

func HerSlips(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(herSlipsInnerSVG).
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

func HexagonAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hexagonAltInnerSVG).
		Element(children...)
}

func HexagonDice(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hexagonDiceInnerSVG).
		Element(children...)
}

func HighPriority(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(highPriorityInnerSVG).
		Element(children...)
}

func HistoricShield(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(historicShieldInnerSVG).
		Element(children...)
}

func HistoricShieldAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(historicShieldAltInnerSVG).
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

func HomeAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeAltInnerSVG).
		Element(children...)
}

func HomeAltSlim(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeAltSlimInnerSVG).
		Element(children...)
}

func HomeAltSlimHoriz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeAltSlimHorizInnerSVG).
		Element(children...)
}

func HomeHospital(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeHospitalInnerSVG).
		Element(children...)
}

func HomeSale(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeSaleInnerSVG).
		Element(children...)
}

func HomeSecure(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeSecureInnerSVG).
		Element(children...)
}

func HomeShield(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeShieldInnerSVG).
		Element(children...)
}

func HomeSimple(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeSimpleInnerSVG).
		Element(children...)
}

func HomeSimpleDoor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeSimpleDoorInnerSVG).
		Element(children...)
}

func HomeTable(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeTableInnerSVG).
		Element(children...)
}

func HomeUser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeUserInnerSVG).
		Element(children...)
}

func HorizDistributionLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(horizDistributionLeftInnerSVG).
		Element(children...)
}

func HorizDistributionRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(horizDistributionRightInnerSVG).
		Element(children...)
}

func HorizontalMerge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(horizontalMergeInnerSVG).
		Element(children...)
}

func HorizontalSplit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(horizontalSplitInnerSVG).
		Element(children...)
}

func Hospital(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hospitalInnerSVG).
		Element(children...)
}

func HospitalSign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hospitalSignInnerSVG).
		Element(children...)
}

func HotAirBalloon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hotAirBalloonInnerSVG).
		Element(children...)
}

func Hourglass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hourglassInnerSVG).
		Element(children...)
}

func HtmlFive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(htmlFiveInnerSVG).
		Element(children...)
}

func Hydrogen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hydrogenInnerSVG).
		Element(children...)
}

func Iconoir(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iconoirInnerSVG).
		Element(children...)
}

func Import(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(importInnerSVG).
		Element(children...)
}

func Inclination(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inclinationInnerSVG).
		Element(children...)
}

func Industry(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(industryInnerSVG).
		Element(children...)
}

func Infinite(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infiniteInnerSVG).
		Element(children...)
}

func InfoEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoEmptyInnerSVG).
		Element(children...)
}

func InputField(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inputFieldInnerSVG).
		Element(children...)
}

func InputOutput(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inputOutputInnerSVG).
		Element(children...)
}

func InputSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inputSearchInnerSVG).
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

func Internet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(internetInnerSVG).
		Element(children...)
}

func Intersect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(intersectInnerSVG).
		Element(children...)
}

func IntersectAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(intersectAltInnerSVG).
		Element(children...)
}

func IosSettings(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iosSettingsInnerSVG).
		Element(children...)
}

func IpAddress(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ipAddressInnerSVG).
		Element(children...)
}

func IrisScan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(irisScanInnerSVG).
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

func ItalicSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(italicSquareInnerSVG).
		Element(children...)
}

func Journal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(journalInnerSVG).
		Element(children...)
}

func JournalPage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(journalPageInnerSVG).
		Element(children...)
}

func JpegFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jpegFormatInnerSVG).
		Element(children...)
}

func JpgFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jpgFormatInnerSVG).
		Element(children...)
}

func KanbanBoard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kanbanBoardInnerSVG).
		Element(children...)
}

func KeyAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyAltInnerSVG).
		Element(children...)
}

func KeyAltBack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyAltBackInnerSVG).
		Element(children...)
}

func KeyAltMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyAltMinusInnerSVG).
		Element(children...)
}

func KeyAltPlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyAltPlusInnerSVG).
		Element(children...)
}

func KeyAltRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyAltRemoveInnerSVG).
		Element(children...)
}

func KeyCommand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyCommandInnerSVG).
		Element(children...)
}

func Keyframe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyframeInnerSVG).
		Element(children...)
}

func KeyframeAlignCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyframeAlignCenterInnerSVG).
		Element(children...)
}

func KeyframeAlignHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyframeAlignHorizontalInnerSVG).
		Element(children...)
}

func KeyframeAlignVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyframeAlignVerticalInnerSVG).
		Element(children...)
}

func KeyframePosition(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyframePositionInnerSVG).
		Element(children...)
}

func Keyframes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyframesInnerSVG).
		Element(children...)
}

func KeyframesCouple(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyframesCoupleInnerSVG).
		Element(children...)
}

func Label(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(labelInnerSVG).
		Element(children...)
}

func Lamp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lampInnerSVG).
		Element(children...)
}

func Language(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(languageInnerSVG).
		Element(children...)
}

func Laptop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopInnerSVG).
		Element(children...)
}

func LaptopCharging(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopChargingInnerSVG).
		Element(children...)
}

func LaptopFix(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopFixInnerSVG).
		Element(children...)
}

func LaptopIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopIssueInnerSVG).
		Element(children...)
}

func LargeSuitcase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(largeSuitcaseInnerSVG).
		Element(children...)
}

func LayoutLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutLeftInnerSVG).
		Element(children...)
}

func LayoutRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutRightInnerSVG).
		Element(children...)
}

func Leaderboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leaderboardInnerSVG).
		Element(children...)
}

func LeaderboardStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leaderboardStarInnerSVG).
		Element(children...)
}

func Leaf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leafInnerSVG).
		Element(children...)
}

func Learning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(learningInnerSVG).
		Element(children...)
}

func LeftRoundArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftRoundArrowInnerSVG).
		Element(children...)
}

func Lens(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lensInnerSVG).
		Element(children...)
}

func Lifebelt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lifebeltInnerSVG).
		Element(children...)
}

func LightBulb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightBulbInnerSVG).
		Element(children...)
}

func LightBulbOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightBulbOffInnerSVG).
		Element(children...)
}

func LightBulbOn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightBulbOnInnerSVG).
		Element(children...)
}

func LineSpace(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineSpaceInnerSVG).
		Element(children...)
}

func Linear(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linearInnerSVG).
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

func Linux(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linuxInnerSVG).
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

func ListSelect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listSelectInnerSVG).
		Element(children...)
}

func LoadActionFloppy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loadActionFloppyInnerSVG).
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

func LockKey(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockKeyInnerSVG).
		Element(children...)
}

func LockedBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockedBookInnerSVG).
		Element(children...)
}

func LockedWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockedWindowInnerSVG).
		Element(children...)
}

func LogDenied(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(logDeniedInnerSVG).
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

func LongArrowDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowDownLeftInnerSVG).
		Element(children...)
}

func LongArrowDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowDownRightInnerSVG).
		Element(children...)
}

func LongArrowLeftDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowLeftDownInnerSVG).
		Element(children...)
}

func LongArrowLeftUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowLeftUpInnerSVG).
		Element(children...)
}

func LongArrowRightDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowRightDownInnerSVG).
		Element(children...)
}

func LongArrowRightUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowRightUpInnerSVG).
		Element(children...)
}

func LongArrowRightUpOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowRightUpOneInnerSVG).
		Element(children...)
}

func LongArrowUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowUpLeftInnerSVG).
		Element(children...)
}

func LongArrowUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(longArrowUpRightInnerSVG).
		Element(children...)
}

func LotOfCash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lotOfCashInnerSVG).
		Element(children...)
}

func MacControlKey(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(macControlKeyInnerSVG).
		Element(children...)
}

func MacDock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(macDockInnerSVG).
		Element(children...)
}

func MacOptionKey(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(macOptionKeyInnerSVG).
		Element(children...)
}

func MacOsWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(macOsWindowInnerSVG).
		Element(children...)
}

func Magnet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magnetInnerSVG).
		Element(children...)
}

func MagnetEnergy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magnetEnergyInnerSVG).
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

func MailIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailInInnerSVG).
		Element(children...)
}

func MailOpened(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailOpenedInnerSVG).
		Element(children...)
}

func MailOut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailOutInnerSVG).
		Element(children...)
}

func Male(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maleInnerSVG).
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

func MapIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapIssueInnerSVG).
		Element(children...)
}

func MapsArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapsArrowInnerSVG).
		Element(children...)
}

func MapsArrowDiagonal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapsArrowDiagonalInnerSVG).
		Element(children...)
}

func MapsArrowIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapsArrowIssueInnerSVG).
		Element(children...)
}

func MapsGoStraight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapsGoStraightInnerSVG).
		Element(children...)
}

func MapsTurnBack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapsTurnBackInnerSVG).
		Element(children...)
}

func MapsTurnLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapsTurnLeftInnerSVG).
		Element(children...)
}

func MapsTurnRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapsTurnRightInnerSVG).
		Element(children...)
}

func MaskSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maskSquareInnerSVG).
		Element(children...)
}

func MastercardCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mastercardCardInnerSVG).
		Element(children...)
}

func MathBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mathBookInnerSVG).
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

func Medal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(medalInnerSVG).
		Element(children...)
}

func MediaImage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaImageInnerSVG).
		Element(children...)
}

func MediaImageFolder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaImageFolderInnerSVG).
		Element(children...)
}

func MediaImageList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaImageListInnerSVG).
		Element(children...)
}

func MediaVideo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaVideoInnerSVG).
		Element(children...)
}

func MediaVideoFolder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaVideoFolderInnerSVG).
		Element(children...)
}

func MediaVideoList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaVideoListInnerSVG).
		Element(children...)
}

func Medium(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediumInnerSVG).
		Element(children...)
}

func MediumPriority(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediumPriorityInnerSVG).
		Element(children...)
}

func Megaphone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(megaphoneInnerSVG).
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

func MenuScale(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuScaleInnerSVG).
		Element(children...)
}

func Message(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageInnerSVG).
		Element(children...)
}

func MessageAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageAlertInnerSVG).
		Element(children...)
}

func MessageText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageTextInnerSVG).
		Element(children...)
}

func Metro(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(metroInnerSVG).
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

func MicAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(micAddInnerSVG).
		Element(children...)
}

func MicCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(micCheckInnerSVG).
		Element(children...)
}

func MicMute(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(micMuteInnerSVG).
		Element(children...)
}

func MicRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(micRemoveInnerSVG).
		Element(children...)
}

func MicSpeaking(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(micSpeakingInnerSVG).
		Element(children...)
}

func MicWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(micWarningInnerSVG).
		Element(children...)
}

func Microscope(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microscopeInnerSVG).
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

func MinusHexagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusHexagonInnerSVG).
		Element(children...)
}

func MinusOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusOneInnerSVG).
		Element(children...)
}

func MinusPinAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusPinAltInnerSVG).
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

func Mirror(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mirrorInnerSVG).
		Element(children...)
}

func MissingFont(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(missingFontInnerSVG).
		Element(children...)
}

func ModernTv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(modernTvInnerSVG).
		Element(children...)
}

func ModernTvFourK(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(modernTvFourKInnerSVG).
		Element(children...)
}

func MoneySquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneySquareInnerSVG).
		Element(children...)
}

func MoonSat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonSatInnerSVG).
		Element(children...)
}

func MoreHoriz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreHorizInnerSVG).
		Element(children...)
}

func MoreHorizCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreHorizCircleInnerSVG).
		Element(children...)
}

func MoreVert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreVertInnerSVG).
		Element(children...)
}

func MoreVertCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreVertCircleInnerSVG).
		Element(children...)
}

func Motorcycle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(motorcycleInnerSVG).
		Element(children...)
}

func MouseButtonLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mouseButtonLeftInnerSVG).
		Element(children...)
}

func MouseButtonRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mouseButtonRightInnerSVG).
		Element(children...)
}

func MouseScrollWheel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mouseScrollWheelInnerSVG).
		Element(children...)
}

func MoveDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveDownInnerSVG).
		Element(children...)
}

func MoveLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveLeftInnerSVG).
		Element(children...)
}

func MoveRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveRightInnerSVG).
		Element(children...)
}

func MoveRuler(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveRulerInnerSVG).
		Element(children...)
}

func MoveUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveUpInnerSVG).
		Element(children...)
}

func Movie(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(movieInnerSVG).
		Element(children...)
}

func MpegFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mpegFormatInnerSVG).
		Element(children...)
}

func MultiBubble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiBubbleInnerSVG).
		Element(children...)
}

func MultiMacOsWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiMacOsWindowInnerSVG).
		Element(children...)
}

func MultiWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiWindowInnerSVG).
		Element(children...)
}

func MultiplePages(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiplePagesInnerSVG).
		Element(children...)
}

func MultiplePagesAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiplePagesAddInnerSVG).
		Element(children...)
}

func MultiplePagesDelete(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiplePagesDeleteInnerSVG).
		Element(children...)
}

func MultiplePagesEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiplePagesEmptyInnerSVG).
		Element(children...)
}

func MultiplePagesRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiplePagesRemoveInnerSVG).
		Element(children...)
}

func MusicDoubleNote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(musicDoubleNoteInnerSVG).
		Element(children...)
}

func MusicDoubleNoteAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(musicDoubleNoteAddInnerSVG).
		Element(children...)
}

func MusicNote(children ...*engine.UberElement) *engine.UberElement {
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

func MusicNoteAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(musicNoteAddInnerSVG).
		Element(children...)
}

func NavArrowDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navArrowDownInnerSVG).
		Element(children...)
}

func NavArrowLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navArrowLeftInnerSVG).
		Element(children...)
}

func NavArrowRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navArrowRightInnerSVG).
		Element(children...)
}

func NavArrowUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navArrowUpInnerSVG).
		Element(children...)
}

func Navigator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navigatorInnerSVG).
		Element(children...)
}

func NavigatorAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navigatorAltInnerSVG).
		Element(children...)
}

func Network(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(networkInnerSVG).
		Element(children...)
}

func NetworkAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(networkAltInnerSVG).
		Element(children...)
}

func NetworkLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(networkLeftInnerSVG).
		Element(children...)
}

func NetworkRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(networkRightInnerSVG).
		Element(children...)
}

func NewTab(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(newTabInnerSVG).
		Element(children...)
}

func NintendoSwitch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nintendoSwitchInnerSVG).
		Element(children...)
}

func Nitrogen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nitrogenInnerSVG).
		Element(children...)
}

func NoAccessWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noAccessWindowInnerSVG).
		Element(children...)
}

func NoBattery(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noBatteryInnerSVG).
		Element(children...)
}

func NoCoin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noCoinInnerSVG).
		Element(children...)
}

func NoCreditCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noCreditCardInnerSVG).
		Element(children...)
}

func NoLink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noLinkInnerSVG).
		Element(children...)
}

func NoLock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noLockInnerSVG).
		Element(children...)
}

func NoSmoking(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noSmokingInnerSVG).
		Element(children...)
}

func NoSmokingCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noSmokingCircleInnerSVG).
		Element(children...)
}

func Notes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notesInnerSVG).
		Element(children...)
}

func Npm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(npmInnerSVG).
		Element(children...)
}

func NpmSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(npmSquareInnerSVG).
		Element(children...)
}

func NumberedListLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(numberedListLeftInnerSVG).
		Element(children...)
}

func NumberedListRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(numberedListRightInnerSVG).
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

func OffTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(offTagInnerSVG).
		Element(children...)
}

func OilIndustry(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(oilIndustryInnerSVG).
		Element(children...)
}

func Okrs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(okrsInnerSVG).
		Element(children...)
}

func OnTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(onTagInnerSVG).
		Element(children...)
}

func OneFingerSelectHandGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(oneFingerSelectHandGestureInnerSVG).
		Element(children...)
}

func OnePointCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(onePointCircleInnerSVG).
		Element(children...)
}

func OneStMedal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(oneStMedalInnerSVG).
		Element(children...)
}

func OpenBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openBookInnerSVG).
		Element(children...)
}

func OpenInBrowser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openInBrowserInnerSVG).
		Element(children...)
}

func OpenInWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openInWindowInnerSVG).
		Element(children...)
}

func OpenNewWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openNewWindowInnerSVG).
		Element(children...)
}

func OpenSelectHandGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openSelectHandGestureInnerSVG).
		Element(children...)
}

func OpenVpn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openVpnInnerSVG).
		Element(children...)
}

func OrangeHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(orangeHalfInnerSVG).
		Element(children...)
}

func OrangeSlice(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(orangeSliceInnerSVG).
		Element(children...)
}

func OrangeSliceAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(orangeSliceAltInnerSVG).
		Element(children...)
}

func OrganicFood(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(organicFoodInnerSVG).
		Element(children...)
}

func OrganicFoodSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(organicFoodSquareInnerSVG).
		Element(children...)
}

func OrthogonalView(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(orthogonalViewInnerSVG).
		Element(children...)
}

func Oxygen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(oxygenInnerSVG).
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

func PackageLock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(packageLockInnerSVG).
		Element(children...)
}

func Packages(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(packagesInnerSVG).
		Element(children...)
}

func Pacman(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pacmanInnerSVG).
		Element(children...)
}

func Page(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageInnerSVG).
		Element(children...)
}

func PageDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageDownInnerSVG).
		Element(children...)
}

func PageEdit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageEditInnerSVG).
		Element(children...)
}

func PageFlip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageFlipInnerSVG).
		Element(children...)
}

func PageLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageLeftInnerSVG).
		Element(children...)
}

func PageRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageRightInnerSVG).
		Element(children...)
}

func PageSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageSearchInnerSVG).
		Element(children...)
}

func PageStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageStarInnerSVG).
		Element(children...)
}

func PageUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageUpInnerSVG).
		Element(children...)
}

func Palette(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paletteInnerSVG).
		Element(children...)
}

func PanoramaEnlarge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panoramaEnlargeInnerSVG).
		Element(children...)
}

func PanoramaReduce(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panoramaReduceInnerSVG).
		Element(children...)
}

func Pants(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pantsInnerSVG).
		Element(children...)
}

func PantsAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pantsAltInnerSVG).
		Element(children...)
}

func Parking(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(parkingInnerSVG).
		Element(children...)
}

func PasswordCursor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(passwordCursorInnerSVG).
		Element(children...)
}

func PasswordError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(passwordErrorInnerSVG).
		Element(children...)
}

func PasswordPass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(passwordPassInnerSVG).
		Element(children...)
}

func PasteClipboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pasteClipboardInnerSVG).
		Element(children...)
}

func PathArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pathArrowInnerSVG).
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

func PauseWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseWindowInnerSVG).
		Element(children...)
}

func Paypal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paypalInnerSVG).
		Element(children...)
}

func PcCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pcCheckInnerSVG).
		Element(children...)
}

func PcFirewall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pcFirewallInnerSVG).
		Element(children...)
}

func PcMouse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pcMouseInnerSVG).
		Element(children...)
}

func PcNoEntry(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pcNoEntryInnerSVG).
		Element(children...)
}

func PcWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pcWarningInnerSVG).
		Element(children...)
}

func PeaceHand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(peaceHandInnerSVG).
		Element(children...)
}

func PenConnectBluetooth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penConnectBluetoothInnerSVG).
		Element(children...)
}

func PenConnectWifi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penConnectWifiInnerSVG).
		Element(children...)
}

func PenTablet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penTabletInnerSVG).
		Element(children...)
}

func PenTabletConnectUsb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penTabletConnectUsbInnerSVG).
		Element(children...)
}

func PenTabletConnectWifi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penTabletConnectWifiInnerSVG).
		Element(children...)
}

func Pentagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pentagonInnerSVG).
		Element(children...)
}

func PeopleTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(peopleTagInnerSVG).
		Element(children...)
}

func Percentage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(percentageInnerSVG).
		Element(children...)
}

func PercentageCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(percentageCircleInnerSVG).
		Element(children...)
}

func PercentageSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(percentageSquareInnerSVG).
		Element(children...)
}

func PerspectiveView(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(perspectiveViewInnerSVG).
		Element(children...)
}

func PharmacyCrossCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pharmacyCrossCircleInnerSVG).
		Element(children...)
}

func PharmacyCrossSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pharmacyCrossSquareInnerSVG).
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

func PhoneAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneAddInnerSVG).
		Element(children...)
}

func PhoneDelete(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneDeleteInnerSVG).
		Element(children...)
}

func PhoneDisabled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneDisabledInnerSVG).
		Element(children...)
}

func PhoneIncome(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneIncomeInnerSVG).
		Element(children...)
}

func PhoneOutcome(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneOutcomeInnerSVG).
		Element(children...)
}

func PhonePaused(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phonePausedInnerSVG).
		Element(children...)
}

func PhoneRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneRemoveInnerSVG).
		Element(children...)
}

func PiggyBank(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(piggyBankInnerSVG).
		Element(children...)
}

func Pillow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pillowInnerSVG).
		Element(children...)
}

func Pin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinInnerSVG).
		Element(children...)
}

func PinAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinAltInnerSVG).
		Element(children...)
}

func PineTree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pineTreeInnerSVG).
		Element(children...)
}

func Pinterest(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinterestInnerSVG).
		Element(children...)
}

func PizzaSlice(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pizzaSliceInnerSVG).
		Element(children...)
}

func Planet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(planetInnerSVG).
		Element(children...)
}

func PlanetAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(planetAltInnerSVG).
		Element(children...)
}

func PlanetSat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(planetSatInnerSVG).
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

func Playlist(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playlistInnerSVG).
		Element(children...)
}

func PlaylistAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playlistAddInnerSVG).
		Element(children...)
}

func PlaylistPlay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playlistPlayInnerSVG).
		Element(children...)
}

func PlaystationGamepad(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playstationGamepadInnerSVG).
		Element(children...)
}

func PlugTypeA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plugTypeAInnerSVG).
		Element(children...)
}

func PlugTypeC(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plugTypeCInnerSVG).
		Element(children...)
}

func PlugTypeG(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plugTypeGInnerSVG).
		Element(children...)
}

func PlugTypeL(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plugTypeLInnerSVG).
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

func PngFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pngFormatInnerSVG).
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

func Podcast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(podcastInnerSVG).
		Element(children...)
}

func Pokeball(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pokeballInnerSVG).
		Element(children...)
}

func Position(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(positionInnerSVG).
		Element(children...)
}

func PositionAlign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(positionAlignInnerSVG).
		Element(children...)
}

func Potion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(potionInnerSVG).
		Element(children...)
}

func Pound(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(poundInnerSVG).
		Element(children...)
}

func PrecisionTool(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(precisionToolInnerSVG).
		Element(children...)
}

func Presentation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(presentationInnerSVG).
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

func PrinterAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printerAltInnerSVG).
		Element(children...)
}

func PrintingPage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printingPageInnerSVG).
		Element(children...)
}

func PriorityDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(priorityDownInnerSVG).
		Element(children...)
}

func PriorityUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(priorityUpInnerSVG).
		Element(children...)
}

func PrivateWifi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(privateWifiInnerSVG).
		Element(children...)
}

func ProfileCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(profileCircleInnerSVG).
		Element(children...)
}

func Prohibition(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(prohibitionInnerSVG).
		Element(children...)
}

func Puzzle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(puzzleInnerSVG).
		Element(children...)
}

func QrCode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(qrCodeInnerSVG).
		Element(children...)
}

func QuestionMark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionMarkInnerSVG).
		Element(children...)
}

func Quote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quoteInnerSVG).
		Element(children...)
}

func QuoteMessage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quoteMessageInnerSVG).
		Element(children...)
}

func Radiation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radiationInnerSVG).
		Element(children...)
}

func Radius(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radiusInnerSVG).
		Element(children...)
}

func Rain(children ...*engine.UberElement) *engine.UberElement {
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

func RawFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rawFormatInnerSVG).
		Element(children...)
}

func ReceiveDollars(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(receiveDollarsInnerSVG).
		Element(children...)
}

func ReceiveEuros(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(receiveEurosInnerSVG).
		Element(children...)
}

func ReceivePounds(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(receivePoundsInnerSVG).
		Element(children...)
}

func ReceiveYens(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(receiveYensInnerSVG).
		Element(children...)
}

func Redo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redoInnerSVG).
		Element(children...)
}

func RedoAction(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redoActionInnerSVG).
		Element(children...)
}

func RedoCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redoCircleInnerSVG).
		Element(children...)
}

func Reduce(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reduceInnerSVG).
		Element(children...)
}

func ReduceRoundArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reduceRoundArrowInnerSVG).
		Element(children...)
}

func Refresh(children ...*engine.UberElement) *engine.UberElement {
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

func RefreshCircular(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshCircularInnerSVG).
		Element(children...)
}

func RefreshDouble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshDoubleInnerSVG).
		Element(children...)
}

func ReloadWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reloadWindowInnerSVG).
		Element(children...)
}

func ReminderHandGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reminderHandGestureInnerSVG).
		Element(children...)
}

func RemoveDatabaseScript(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeDatabaseScriptInnerSVG).
		Element(children...)
}

func RemoveFolder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeFolderInnerSVG).
		Element(children...)
}

func RemoveFrame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeFrameInnerSVG).
		Element(children...)
}

func RemoveFromCart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeFromCartInnerSVG).
		Element(children...)
}

func RemoveKeyframe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeKeyframeInnerSVG).
		Element(children...)
}

func RemoveKeyframeAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeKeyframeAltInnerSVG).
		Element(children...)
}

func RemoveKeyframes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeKeyframesInnerSVG).
		Element(children...)
}

func RemoveLink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeLinkInnerSVG).
		Element(children...)
}

func RemoveMediaImage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeMediaImageInnerSVG).
		Element(children...)
}

func RemoveMediaVideo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeMediaVideoInnerSVG).
		Element(children...)
}

func RemovePage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removePageInnerSVG).
		Element(children...)
}

func RemovePageAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removePageAltInnerSVG).
		Element(children...)
}

func RemovePin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removePinInnerSVG).
		Element(children...)
}

func RemovePinAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removePinAltInnerSVG).
		Element(children...)
}

func RemoveSelection(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeSelectionInnerSVG).
		Element(children...)
}

func RemoveSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeSquareInnerSVG).
		Element(children...)
}

func RemoveUser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeUserInnerSVG).
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

func RepeatOnce(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(repeatOnceInnerSVG).
		Element(children...)
}

func Reply(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(replyInnerSVG).
		Element(children...)
}

func ReplyToMessage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(replyToMessageInnerSVG).
		Element(children...)
}

func ReportColumns(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reportColumnsInnerSVG).
		Element(children...)
}

func Reports(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reportsInnerSVG).
		Element(children...)
}

func Repository(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(repositoryInnerSVG).
		Element(children...)
}

func Restart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(restartInnerSVG).
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

func Rhombus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rhombusInnerSVG).
		Element(children...)
}

func RightRoundArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rightRoundArrowInnerSVG).
		Element(children...)
}

func Rings(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ringsInnerSVG).
		Element(children...)
}

func Rocket(children ...*engine.UberElement) *engine.UberElement {
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

func Rook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rookInnerSVG).
		Element(children...)
}

func RotateCameraLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateCameraLeftInnerSVG).
		Element(children...)
}

func RotateCameraRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateCameraRightInnerSVG).
		Element(children...)
}

func RoundFlask(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(roundFlaskInnerSVG).
		Element(children...)
}

func RoundedMirror(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(roundedMirrorInnerSVG).
		Element(children...)
}

func RssFeed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssFeedInnerSVG).
		Element(children...)
}

func RssFeedTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssFeedTagInnerSVG).
		Element(children...)
}

func RubikCube(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rubikCubeInnerSVG).
		Element(children...)
}

func Ruler(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rulerInnerSVG).
		Element(children...)
}

func RulerAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rulerAddInnerSVG).
		Element(children...)
}

func RulerCombine(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rulerCombineInnerSVG).
		Element(children...)
}

func RulerRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rulerRemoveInnerSVG).
		Element(children...)
}

func Running(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(runningInnerSVG).
		Element(children...)
}

func Safari(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(safariInnerSVG).
		Element(children...)
}

func Sandals(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sandalsInnerSVG).
		Element(children...)
}

func SaveActionFloppy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(saveActionFloppyInnerSVG).
		Element(children...)
}

func SaveFloppyDisk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(saveFloppyDiskInnerSVG).
		Element(children...)
}

func ScaleFrameEnlarge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scaleFrameEnlargeInnerSVG).
		Element(children...)
}

func ScaleFrameReduce(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scaleFrameReduceInnerSVG).
		Element(children...)
}

func ScanBarcode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scanBarcodeInnerSVG).
		Element(children...)
}

func ScanQrCode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scanQrCodeInnerSVG).
		Element(children...)
}

func Scanning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scanningInnerSVG).
		Element(children...)
}

func Scarf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scarfInnerSVG).
		Element(children...)
}

func Scissor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scissorInnerSVG).
		Element(children...)
}

func ScissorAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scissorAltInnerSVG).
		Element(children...)
}

func Screenshot(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenshotInnerSVG).
		Element(children...)
}

func SeaAndSun(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(seaAndSunInnerSVG).
		Element(children...)
}

func SeaWaves(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(seaWavesInnerSVG).
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

func SearchEngine(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchEngineInnerSVG).
		Element(children...)
}

func SearchFont(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchFontInnerSVG).
		Element(children...)
}

func SearchWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchWindowInnerSVG).
		Element(children...)
}

func SecureWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(secureWindowInnerSVG).
		Element(children...)
}

func SecurityPass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(securityPassInnerSVG).
		Element(children...)
}

func SelectWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(selectWindowInnerSVG).
		Element(children...)
}

func Selection(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(selectionInnerSVG).
		Element(children...)
}

func SelectiveTool(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(selectiveToolInnerSVG).
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

func SendDiagonal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendDiagonalInnerSVG).
		Element(children...)
}

func SendDollars(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendDollarsInnerSVG).
		Element(children...)
}

func SendEuros(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendEurosInnerSVG).
		Element(children...)
}

func SendMail(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendMailInnerSVG).
		Element(children...)
}

func SendPounds(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendPoundsInnerSVG).
		Element(children...)
}

func SendYens(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendYensInnerSVG).
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

func ServerConnection(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(serverConnectionInnerSVG).
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

func SettingsCloud(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(settingsCloudInnerSVG).
		Element(children...)
}

func SettingsProfiles(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(settingsProfilesInnerSVG).
		Element(children...)
}

func ShareAndroid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareAndroidInnerSVG).
		Element(children...)
}

func ShareIos(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareIosInnerSVG).
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

func ShieldAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldAddInnerSVG).
		Element(children...)
}

func ShieldAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldAlertInnerSVG).
		Element(children...)
}

func ShieldAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldAltInnerSVG).
		Element(children...)
}

func ShieldBroken(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldBrokenInnerSVG).
		Element(children...)
}

func ShieldCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldCheckInnerSVG).
		Element(children...)
}

func ShieldCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldCrossInnerSVG).
		Element(children...)
}

func ShieldDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldDownloadInnerSVG).
		Element(children...)
}

func ShieldEye(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldEyeInnerSVG).
		Element(children...)
}

func ShieldLoading(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldLoadingInnerSVG).
		Element(children...)
}

func ShieldMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldMinusInnerSVG).
		Element(children...)
}

func ShieldQuestion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldQuestionInnerSVG).
		Element(children...)
}

func ShieldSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldSearchInnerSVG).
		Element(children...)
}

func ShieldUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldUploadInnerSVG).
		Element(children...)
}

func Shop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shopInnerSVG).
		Element(children...)
}

func ShopAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shopAltInnerSVG).
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

func ShoppingBagAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagAddInnerSVG).
		Element(children...)
}

func ShoppingBagAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagAltInnerSVG).
		Element(children...)
}

func ShoppingBagArrowDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagArrowDownInnerSVG).
		Element(children...)
}

func ShoppingBagArrowUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagArrowUpInnerSVG).
		Element(children...)
}

func ShoppingBagCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagCheckInnerSVG).
		Element(children...)
}

func ShoppingBagIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagIssueInnerSVG).
		Element(children...)
}

func ShoppingBagRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagRemoveInnerSVG).
		Element(children...)
}

func ShoppingCode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingCodeInnerSVG).
		Element(children...)
}

func ShoppingCodeCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingCodeCheckInnerSVG).
		Element(children...)
}

func ShoppingCodeError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingCodeErrorInnerSVG).
		Element(children...)
}

func ShortPants(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shortPantsInnerSVG).
		Element(children...)
}

func ShortPantsAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shortPantsAltInnerSVG).
		Element(children...)
}

func Shortcut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shortcutInnerSVG).
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

func SidebarCollapse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sidebarCollapseInnerSVG).
		Element(children...)
}

func SidebarExpand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sidebarExpandInnerSVG).
		Element(children...)
}

func SigmaFunction(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sigmaFunctionInnerSVG).
		Element(children...)
}

func SimpleCart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(simpleCartInnerSVG).
		Element(children...)
}

func SineWave(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sineWaveInnerSVG).
		Element(children...)
}

func SingleTapGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(singleTapGestureInnerSVG).
		Element(children...)
}

func Skateboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skateboardInnerSVG).
		Element(children...)
}

func Skateboarding(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skateboardingInnerSVG).
		Element(children...)
}

func SkipNext(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skipNextInnerSVG).
		Element(children...)
}

func SkipPrev(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skipPrevInnerSVG).
		Element(children...)
}

func SleeperChair(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sleeperChairInnerSVG).
		Element(children...)
}

func SmallLamp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smallLampInnerSVG).
		Element(children...)
}

func SmallLampAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smallLampAltInnerSVG).
		Element(children...)
}

func SmallShop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smallShopInnerSVG).
		Element(children...)
}

func SmallShopAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smallShopAltInnerSVG).
		Element(children...)
}

func SmartphoneDevice(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smartphoneDeviceInnerSVG).
		Element(children...)
}

func Smoking(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smokingInnerSVG).
		Element(children...)
}

func Snapchat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snapchatInnerSVG).
		Element(children...)
}

func Snow(children ...*engine.UberElement) *engine.UberElement {
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

func SnowFlake(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snowFlakeInnerSVG).
		Element(children...)
}

func Soap(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soapInnerSVG).
		Element(children...)
}

func SoccerBall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soccerBallInnerSVG).
		Element(children...)
}

func Sofa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sofaInnerSVG).
		Element(children...)
}

func Soil(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soilInnerSVG).
		Element(children...)
}

func SoilAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soilAltInnerSVG).
		Element(children...)
}

func Sort(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortInnerSVG).
		Element(children...)
}

func SortDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortDownInnerSVG).
		Element(children...)
}

func SortUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortUpInnerSVG).
		Element(children...)
}

func SoundHigh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundHighInnerSVG).
		Element(children...)
}

func SoundLow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundLowInnerSVG).
		Element(children...)
}

func SoundMin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundMinInnerSVG).
		Element(children...)
}

func SoundOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundOffInnerSVG).
		Element(children...)
}

func Spades(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spadesInnerSVG).
		Element(children...)
}

func Sphere(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sphereInnerSVG).
		Element(children...)
}

func Spiral(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spiralInnerSVG).
		Element(children...)
}

func SpockHandGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spockHandGestureInnerSVG).
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

func SquareWave(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareWaveInnerSVG).
		Element(children...)
}

func Stackoverflow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stackoverflowInnerSVG).
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

func StarDashed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starDashedInnerSVG).
		Element(children...)
}

func StarHalfDashed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starHalfDashedInnerSVG).
		Element(children...)
}

func StatDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statDownInnerSVG).
		Element(children...)
}

func StatUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statUpInnerSVG).
		Element(children...)
}

func StatsDownSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statsDownSquareInnerSVG).
		Element(children...)
}

func StatsReport(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statsReportInnerSVG).
		Element(children...)
}

func StatsUpSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statsUpSquareInnerSVG).
		Element(children...)
}

func Strategy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(strategyInnerSVG).
		Element(children...)
}

func Stretching(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stretchingInnerSVG).
		Element(children...)
}

func Stroller(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(strollerInnerSVG).
		Element(children...)
}

func StyleBorder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(styleBorderInnerSVG).
		Element(children...)
}

func SubmitDocument(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(submitDocumentInnerSVG).
		Element(children...)
}

func Substract(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(substractInnerSVG).
		Element(children...)
}

func Suggestion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(suggestionInnerSVG).
		Element(children...)
}

func SunLight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunLightInnerSVG).
		Element(children...)
}

func SvgFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(svgFormatInnerSVG).
		Element(children...)
}

func Swimming(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swimmingInnerSVG).
		Element(children...)
}

func SwipeDownGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swipeDownGestureInnerSVG).
		Element(children...)
}

func SwipeLeftGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swipeLeftGestureInnerSVG).
		Element(children...)
}

func SwipeRightGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swipeRightGestureInnerSVG).
		Element(children...)
}

func SwipeTwoFingersDownGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swipeTwoFingersDownGestureInnerSVG).
		Element(children...)
}

func SwipeTwoFingersLeftGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swipeTwoFingersLeftGestureInnerSVG).
		Element(children...)
}

func SwipeTwoFingersRightGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swipeTwoFingersRightGestureInnerSVG).
		Element(children...)
}

func SwipeTwoFingersUpGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swipeTwoFingersUpGestureInnerSVG).
		Element(children...)
}

func SwipeUpGesture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swipeUpGestureInnerSVG).
		Element(children...)
}

func SwitchOff(children ...*engine.UberElement) *engine.UberElement {
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

func SwitchOn(children ...*engine.UberElement) *engine.UberElement {
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

func SystemRestart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(systemRestartInnerSVG).
		Element(children...)
}

func SystemShut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(systemShutInnerSVG).
		Element(children...)
}

func TShirt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tShirtInnerSVG).
		Element(children...)
}

func Table(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableInnerSVG).
		Element(children...)
}

func TableRows(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableRowsInnerSVG).
		Element(children...)
}

func TableTwoColumns(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableTwoColumnsInnerSVG).
		Element(children...)
}

func TaskList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(taskListInnerSVG).
		Element(children...)
}

func Telegram(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(telegramInnerSVG).
		Element(children...)
}

func TelegramCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(telegramCircleInnerSVG).
		Element(children...)
}

func TennisBall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tennisBallInnerSVG).
		Element(children...)
}

func TennisBallAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tennisBallAltInnerSVG).
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

func TerminalTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalTagInnerSVG).
		Element(children...)
}

func TestTube(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(testTubeInnerSVG).
		Element(children...)
}

func Text(children ...*engine.UberElement) *engine.UberElement {
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

func TextAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAltInnerSVG).
		Element(children...)
}

func TextBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textBoxInnerSVG).
		Element(children...)
}

func TextSize(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textSizeInnerSVG).
		Element(children...)
}

func ThreeDAddHole(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDAddHoleInnerSVG).
		Element(children...)
}

func ThreeDArc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDArcInnerSVG).
		Element(children...)
}

func ThreeDArcCenterPt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDArcCenterPtInnerSVG).
		Element(children...)
}

func ThreeDBridge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDBridgeInnerSVG).
		Element(children...)
}

func ThreeDCenterBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDCenterBoxInnerSVG).
		Element(children...)
}

func ThreeDEllipse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDEllipseInnerSVG).
		Element(children...)
}

func ThreeDEllipseThreePts(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDEllipseThreePtsInnerSVG).
		Element(children...)
}

func ThreeDPtBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDPtBoxInnerSVG).
		Element(children...)
}

func ThreeDRectCornerToCorner(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDRectCornerToCornerInnerSVG).
		Element(children...)
}

func ThreeDRectFromCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDRectFromCenterInnerSVG).
		Element(children...)
}

func ThreeDRectThreePts(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDRectThreePtsInnerSVG).
		Element(children...)
}

func ThreeDSelectEdge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDSelectEdgeInnerSVG).
		Element(children...)
}

func ThreeDSelectFace(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDSelectFaceInnerSVG).
		Element(children...)
}

func ThreeDSelectPoint(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDSelectPointInnerSVG).
		Element(children...)
}

func ThreeDSelectSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDSelectSolidInnerSVG).
		Element(children...)
}

func ThreeDThreePtsBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDThreePtsBoxInnerSVG).
		Element(children...)
}

func ThreeHundredSixtyView(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeHundredSixtyViewInnerSVG).
		Element(children...)
}

func ThreePointsCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threePointsCircleInnerSVG).
		Element(children...)
}

func ThreeStars(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeStarsInnerSVG).
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

func Thunderstorm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thunderstormInnerSVG).
		Element(children...)
}

func TifFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tifFormatInnerSVG).
		Element(children...)
}

func TiffFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tiffFormatInnerSVG).
		Element(children...)
}

func Tiktok(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tiktokInnerSVG).
		Element(children...)
}

func Timer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timerInnerSVG).
		Element(children...)
}

func TimerOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timerOffInnerSVG).
		Element(children...)
}

func Tools(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toolsInnerSVG).
		Element(children...)
}

func Tournament(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tournamentInnerSVG).
		Element(children...)
}

func Tower(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(towerInnerSVG).
		Element(children...)
}

func TowerCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(towerCheckInnerSVG).
		Element(children...)
}

func TowerNoAccess(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(towerNoAccessInnerSVG).
		Element(children...)
}

func TowerWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(towerWarningInnerSVG).
		Element(children...)
}

func Trademark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trademarkInnerSVG).
		Element(children...)
}

func Train(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trainInnerSVG).
		Element(children...)
}

func Tram(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tramInnerSVG).
		Element(children...)
}

func TransitionDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transitionDownInnerSVG).
		Element(children...)
}

func TransitionLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transitionLeftInnerSVG).
		Element(children...)
}

func TransitionRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transitionRightInnerSVG).
		Element(children...)
}

func TransitionUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transitionUpInnerSVG).
		Element(children...)
}

func Translate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(translateInnerSVG).
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

func Treadmill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(treadmillInnerSVG).
		Element(children...)
}

func Tree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(treeInnerSVG).
		Element(children...)
}

func Trekking(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trekkingInnerSVG).
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

func TriangleFlag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleFlagInnerSVG).
		Element(children...)
}

func TriangleFlagCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleFlagCircleInnerSVG).
		Element(children...)
}

func TriangleFlagTwoStripes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleFlagTwoStripesInnerSVG).
		Element(children...)
}

func Trophy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
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
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(truckInnerSVG).
		Element(children...)
}

func TruckLength(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(truckLengthInnerSVG).
		Element(children...)
}

func Tunnel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tunnelInnerSVG).
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

func TvFix(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tvFixInnerSVG).
		Element(children...)
}

func TvIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tvIssueInnerSVG).
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

func TwitterVerifiedBadge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitterVerifiedBadgeInnerSVG).
		Element(children...)
}

func TwoPointsCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twoPointsCircleInnerSVG).
		Element(children...)
}

func TwoSeaterSofa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twoSeaterSofaInnerSVG).
		Element(children...)
}

func TwoXTwoCell(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twoXTwoCellInnerSVG).
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

func UmbrellaFull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(umbrellaFullInnerSVG).
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

func UnderlineSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(underlineSquareInnerSVG).
		Element(children...)
}

func Undo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(undoInnerSVG).
		Element(children...)
}

func UndoAction(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(undoActionInnerSVG).
		Element(children...)
}

func UndoCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(undoCircleInnerSVG).
		Element(children...)
}

func Union(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unionInnerSVG).
		Element(children...)
}

func UnionAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unionAltInnerSVG).
		Element(children...)
}

func UnionHorizAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unionHorizAltInnerSVG).
		Element(children...)
}

func Unity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unityInnerSVG).
		Element(children...)
}

func UnityFive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unityFiveInnerSVG).
		Element(children...)
}

func UpRoundArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upRoundArrowInnerSVG).
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

func UploadDataWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadDataWindowInnerSVG).
		Element(children...)
}

func UploadSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadSquareInnerSVG).
		Element(children...)
}

func Usb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usbInnerSVG).
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

func UserBag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userBagInnerSVG).
		Element(children...)
}

func UserCart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCartInnerSVG).
		Element(children...)
}

func UserCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCircleInnerSVG).
		Element(children...)
}

func UserCrown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCrownInnerSVG).
		Element(children...)
}

func UserLove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userLoveInnerSVG).
		Element(children...)
}

func UserScan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userScanInnerSVG).
		Element(children...)
}

func UserSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userSquareInnerSVG).
		Element(children...)
}

func UserStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userStarInnerSVG).
		Element(children...)
}

func Vegan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(veganInnerSVG).
		Element(children...)
}

func VeganCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(veganCircleInnerSVG).
		Element(children...)
}

func VeganSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(veganSquareInnerSVG).
		Element(children...)
}

func VerifiedBadge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(verifiedBadgeInnerSVG).
		Element(children...)
}

func VerifiedUser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(verifiedUserInnerSVG).
		Element(children...)
}

func VerticalMerge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(verticalMergeInnerSVG).
		Element(children...)
}

func VerticalSplit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(verticalSplitInnerSVG).
		Element(children...)
}

func Vials(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vialsInnerSVG).
		Element(children...)
}

func VideoCamera(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(videoCameraInnerSVG).
		Element(children...)
}

func VideoCameraOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(videoCameraOffInnerSVG).
		Element(children...)
}

func VideoProjector(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(videoProjectorInnerSVG).
		Element(children...)
}

func ViewColumnsThree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewColumnsThreeInnerSVG).
		Element(children...)
}

func ViewColumnsTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewColumnsTwoInnerSVG).
		Element(children...)
}

func ViewGrid(children ...*engine.UberElement) *engine.UberElement {
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

func ViewStructureDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewStructureDownInnerSVG).
		Element(children...)
}

func ViewStructureUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewStructureUpInnerSVG).
		Element(children...)
}

func Voice(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voiceInnerSVG).
		Element(children...)
}

func VoiceCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voiceCircleInnerSVG).
		Element(children...)
}

func VoiceError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voiceErrorInnerSVG).
		Element(children...)
}

func VoiceLockCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voiceLockCircleInnerSVG).
		Element(children...)
}

func VoiceOk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voiceOkInnerSVG).
		Element(children...)
}

func VoicePhone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voicePhoneInnerSVG).
		Element(children...)
}

func VoiceScan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voiceScanInnerSVG).
		Element(children...)
}

func VoiceSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voiceSquareInnerSVG).
		Element(children...)
}

func VrSymbol(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vrSymbolInnerSVG).
		Element(children...)
}

func Waist(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(waistInnerSVG).
		Element(children...)
}

func Walking(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(walkingInnerSVG).
		Element(children...)
}

func Wallet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(walletInnerSVG).
		Element(children...)
}

func WarningCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningCircleInnerSVG).
		Element(children...)
}

func WarningHexagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningHexagonInnerSVG).
		Element(children...)
}

func WarningSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningSquareInnerSVG).
		Element(children...)
}

func WarningTriangle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningTriangleInnerSVG).
		Element(children...)
}

func WarningWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningWindowInnerSVG).
		Element(children...)
}

func Wash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(washInnerSVG).
		Element(children...)
}

func WashingMachine(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(washingMachineInnerSVG).
		Element(children...)
}

func WateringSoil(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wateringSoilInnerSVG).
		Element(children...)
}

func WebWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webWindowInnerSVG).
		Element(children...)
}

func WebWindowClose(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webWindowCloseInnerSVG).
		Element(children...)
}

func WebWindowEnergyConsumption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webWindowEnergyConsumptionInnerSVG).
		Element(children...)
}

func WebpFormat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webpFormatInnerSVG).
		Element(children...)
}

func Weight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(weightInnerSVG).
		Element(children...)
}

func WeightAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(weightAltInnerSVG).
		Element(children...)
}

func WhiteFlag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(whiteFlagInnerSVG).
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

func WifiError(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiErrorInnerSVG).
		Element(children...)
}

func WifiIssue(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiIssueInnerSVG).
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

func WifiSignalNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiSignalNoneInnerSVG).
		Element(children...)
}

func WifiTag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiTagInnerSVG).
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

func Windows(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowsInnerSVG).
		Element(children...)
}

func WomenTShirt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(womenTShirtInnerSVG).
		Element(children...)
}

func WrapText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wrapTextInnerSVG).
		Element(children...)
}

func Wrench(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wrenchInnerSVG).
		Element(children...)
}

func Wristwatch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wristwatchInnerSVG).
		Element(children...)
}

func Www(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wwwInnerSVG).
		Element(children...)
}

func XCoordinate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xCoordinateInnerSVG).
		Element(children...)
}

func XboxA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xboxAInnerSVG).
		Element(children...)
}

func XboxB(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xboxBInnerSVG).
		Element(children...)
}

func XboxX(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xboxXInnerSVG).
		Element(children...)
}

func XboxY(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xboxYInnerSVG).
		Element(children...)
}

func XrayView(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xrayViewInnerSVG).
		Element(children...)
}

func YCoordinate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yCoordinateInnerSVG).
		Element(children...)
}

func Yen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yenInnerSVG).
		Element(children...)
}

func YenSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yenSquareInnerSVG).
		Element(children...)
}

func Yoga(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yogaInnerSVG).
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

func ZCoordinate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zCoordinateInnerSVG).
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
	case "accessibility":
		return Accessibility(), nil
	case "accessibility-sign":
		return AccessibilitySign(), nil
	case "accessibility-tech":
		return AccessibilityTech(), nil
	case "activity":
		return Activity(), nil
	case "add-circle":
		return AddCircle(), nil
	case "add-database-script":
		return AddDatabaseScript(), nil
	case "add-folder":
		return AddFolder(), nil
	case "add-frame":
		return AddFrame(), nil
	case "add-hexagon":
		return AddHexagon(), nil
	case "add-keyframe":
		return AddKeyframe(), nil
	case "add-keyframe-alt":
		return AddKeyframeAlt(), nil
	case "add-keyframes":
		return AddKeyframes(), nil
	case "add-lens":
		return AddLens(), nil
	case "add-media-image":
		return AddMediaImage(), nil
	case "add-media-video":
		return AddMediaVideo(), nil
	case "add-page":
		return AddPage(), nil
	case "add-page-alt":
		return AddPageAlt(), nil
	case "add-pin-alt":
		return AddPinAlt(), nil
	case "add-selection":
		return AddSelection(), nil
	case "add-square":
		return AddSquare(), nil
	case "add-to-cart":
		return AddToCart(), nil
	case "add-user":
		return AddUser(), nil
	case "african-tree":
		return AfricanTree(), nil
	case "agile":
		return Agile(), nil
	case "air-conditioner":
		return AirConditioner(), nil
	case "airplane":
		return Airplane(), nil
	case "airplane-helix":
		return AirplaneHelix(), nil
	case "airplane-helix-45deg":
		return AirplaneHelixFortyFiveDeg(), nil
	case "airplane-off":
		return AirplaneOff(), nil
	case "airplane-rotation":
		return AirplaneRotation(), nil
	case "airplay":
		return Airplay(), nil
	case "alarm":
		return Alarm(), nil
	case "album":
		return Album(), nil
	case "album-carousel":
		return AlbumCarousel(), nil
	case "album-list":
		return AlbumList(), nil
	case "album-open":
		return AlbumOpen(), nil
	case "align-bottom-box":
		return AlignBottomBox(), nil
	case "align-center":
		return AlignCenter(), nil
	case "align-justify":
		return AlignJustify(), nil
	case "align-left":
		return AlignLeft(), nil
	case "align-left-box":
		return AlignLeftBox(), nil
	case "align-right":
		return AlignRight(), nil
	case "align-right-box":
		return AlignRightBox(), nil
	case "align-top-box":
		return AlignTopBox(), nil
	case "angle-tool":
		return AngleTool(), nil
	case "antenna":
		return Antenna(), nil
	case "antenna-off":
		return AntennaOff(), nil
	case "antenna-signal":
		return AntennaSignal(), nil
	case "antenna-signal-tag":
		return AntennaSignalTag(), nil
	case "app-notification":
		return AppNotification(), nil
	case "app-window":
		return AppWindow(), nil
	case "apple":
		return Apple(), nil
	case "apple-half":
		return AppleHalf(), nil
	case "apple-half-alt":
		return AppleHalfAlt(), nil
	case "apple-imac-2021":
		return AppleImacTwoThousandTwentyOne(), nil
	case "apple-imac-2021-side":
		return AppleImacTwoThousandTwentyOneSide(), nil
	case "apple-mac":
		return AppleMac(), nil
	case "apple-swift":
		return AppleSwift(), nil
	case "apple-wallet":
		return AppleWallet(), nil
	case "ar-symbol":
		return ArSymbol(), nil
	case "arcade":
		return Arcade(), nil
	case "archery":
		return Archery(), nil
	case "archery-match":
		return ArcheryMatch(), nil
	case "archive":
		return Archive(), nil
	case "area-search":
		return AreaSearch(), nil
	case "arrow-archery":
		return ArrowArchery(), nil
	case "arrow-bl-circle":
		return ArrowBlCircle(), nil
	case "arrow-bl-square":
		return ArrowBlSquare(), nil
	case "arrow-br-circle":
		return ArrowBrCircle(), nil
	case "arrow-br-square":
		return ArrowBrSquare(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-down-circle":
		return ArrowDownCircle(), nil
	case "arrow-email-forward":
		return ArrowEmailForward(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-left-circle":
		return ArrowLeftCircle(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-right-circle":
		return ArrowRightCircle(), nil
	case "arrow-separate":
		return ArrowSeparate(), nil
	case "arrow-separate-vertical":
		return ArrowSeparateVertical(), nil
	case "arrow-tl-circle":
		return ArrowTlCircle(), nil
	case "arrow-tl-square":
		return ArrowTlSquare(), nil
	case "arrow-tr-circle":
		return ArrowTrCircle(), nil
	case "arrow-tr-square":
		return ArrowTrSquare(), nil
	case "arrow-union":
		return ArrowUnion(), nil
	case "arrow-union-vertical":
		return ArrowUnionVertical(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrow-up-circle":
		return ArrowUpCircle(), nil
	case "asana":
		return Asana(), nil
	case "at-sign":
		return AtSign(), nil
	case "at-sign-circle":
		return AtSignCircle(), nil
	case "atom":
		return Atom(), nil
	case "attachment":
		return Attachment(), nil
	case "augmented-reality":
		return AugmentedReality(), nil
	case "auto-flash":
		return AutoFlash(), nil
	case "avi-format":
		return AviFormat(), nil
	case "axes":
		return Axes(), nil
	case "backward-15-seconds":
		return BackwardFifteenSeconds(), nil
	case "bag":
		return Bag(), nil
	case "bank":
		return Bank(), nil
	case "barcode":
		return Barcode(), nil
	case "basketball":
		return Basketball(), nil
	case "basketball-alt":
		return BasketballAlt(), nil
	case "basketball-field":
		return BasketballField(), nil
	case "battery-charging":
		return BatteryCharging(), nil
	case "battery-empty":
		return BatteryEmpty(), nil
	case "battery-50":
		return BatteryFifty(), nil
	case "battery-full":
		return BatteryFull(), nil
	case "battery-indicator":
		return BatteryIndicator(), nil
	case "battery-75":
		return BatterySeventyFive(), nil
	case "battery-25":
		return BatteryTwentyFive(), nil
	case "battery-warning":
		return BatteryWarning(), nil
	case "bbq":
		return Bbq(), nil
	case "beach-bag":
		return BeachBag(), nil
	case "beach-bag-big":
		return BeachBagBig(), nil
	case "bed":
		return Bed(), nil
	case "bed-ready":
		return BedReady(), nil
	case "behance":
		return Behance(), nil
	case "behance-tag":
		return BehanceTag(), nil
	case "bell":
		return Bell(), nil
	case "bell-notification":
		return BellNotification(), nil
	case "bell-off":
		return BellOff(), nil
	case "bicycle":
		return Bicycle(), nil
	case "bin":
		return Bin(), nil
	case "bin-add":
		return BinAdd(), nil
	case "bin-full":
		return BinFull(), nil
	case "bin-half":
		return BinHalf(), nil
	case "bin-minus":
		return BinMinus(), nil
	case "bishop":
		return Bishop(), nil
	case "bitbucket":
		return Bitbucket(), nil
	case "bluetooth":
		return Bluetooth(), nil
	case "bluetooth-tag":
		return BluetoothTag(), nil
	case "bold":
		return Bold(), nil
	case "bold-square":
		return BoldSquare(), nil
	case "bonfire":
		return Bonfire(), nil
	case "book":
		return Book(), nil
	case "book-stack":
		return BookStack(), nil
	case "bookmark-book":
		return BookmarkBook(), nil
	case "bookmark-circle":
		return BookmarkCircle(), nil
	case "bookmark-empty":
		return BookmarkEmpty(), nil
	case "border-bl":
		return BorderBl(), nil
	case "border-bottom":
		return BorderBottom(), nil
	case "border-br":
		return BorderBr(), nil
	case "border-inner":
		return BorderInner(), nil
	case "border-left":
		return BorderLeft(), nil
	case "border-out":
		return BorderOut(), nil
	case "border-right":
		return BorderRight(), nil
	case "border-tl":
		return BorderTl(), nil
	case "border-top":
		return BorderTop(), nil
	case "border-tr":
		return BorderTr(), nil
	case "bounce-left":
		return BounceLeft(), nil
	case "bounce-right":
		return BounceRight(), nil
	case "bowling-ball":
		return BowlingBall(), nil
	case "box":
		return Box(), nil
	case "box-iso":
		return BoxIso(), nil
	case "boxing-glove":
		return BoxingGlove(), nil
	case "brain":
		return Brain(), nil
	case "brain-electricity":
		return BrainElectricity(), nil
	case "brain-research":
		return BrainResearch(), nil
	case "brain-warning":
		return BrainWarning(), nil
	case "bread-slice":
		return BreadSlice(), nil
	case "bright-crown":
		return BrightCrown(), nil
	case "bright-star":
		return BrightStar(), nil
	case "brightness":
		return Brightness(), nil
	case "brightness-window":
		return BrightnessWindow(), nil
	case "bubble-download":
		return BubbleDownload(), nil
	case "bubble-error":
		return BubbleError(), nil
	case "bubble-income":
		return BubbleIncome(), nil
	case "bubble-outcome":
		return BubbleOutcome(), nil
	case "bubble-search":
		return BubbleSearch(), nil
	case "bubble-star":
		return BubbleStar(), nil
	case "bubble-upload":
		return BubbleUpload(), nil
	case "bubble-warning":
		return BubbleWarning(), nil
	case "building":
		return Building(), nil
	case "bus":
		return Bus(), nil
	case "bus-stop":
		return BusStop(), nil
	case "cable-tag":
		return CableTag(), nil
	case "calculator":
		return Calculator(), nil
	case "calendar":
		return Calendar(), nil
	case "calendar-minus":
		return CalendarMinus(), nil
	case "calendar-plus":
		return CalendarPlus(), nil
	case "camera":
		return Camera(), nil
	case "cancel":
		return Cancel(), nil
	case "candlestick-chart":
		return CandlestickChart(), nil
	case "car":
		return Car(), nil
	case "carbon":
		return Carbon(), nil
	case "card-issue":
		return CardIssue(), nil
	case "card-locked":
		return CardLocked(), nil
	case "card-security":
		return CardSecurity(), nil
	case "card-wallet":
		return CardWallet(), nil
	case "cart":
		return Cart(), nil
	case "cart-alt":
		return CartAlt(), nil
	case "cash":
		return Cash(), nil
	case "center-align":
		return CenterAlign(), nil
	case "chat-add":
		return ChatAdd(), nil
	case "chat-bubble":
		return ChatBubble(), nil
	case "chat-bubble-check":
		return ChatBubbleCheck(), nil
	case "chat-bubble-check-1":
		return ChatBubbleCheckOne(), nil
	case "chat-bubble-empty":
		return ChatBubbleEmpty(), nil
	case "chat-bubble-error":
		return ChatBubbleError(), nil
	case "chat-bubble-question":
		return ChatBubbleQuestion(), nil
	case "chat-bubble-translate":
		return ChatBubbleTranslate(), nil
	case "chat-bubble-warning":
		return ChatBubbleWarning(), nil
	case "chat-lines":
		return ChatLines(), nil
	case "chat-remove":
		return ChatRemove(), nil
	case "check":
		return Check(), nil
	case "check-circle":
		return CheckCircle(), nil
	case "check-window":
		return CheckWindow(), nil
	case "chocolate":
		return Chocolate(), nil
	case "chromecast":
		return Chromecast(), nil
	case "chromecast-active":
		return ChromecastActive(), nil
	case "church":
		return Church(), nil
	case "church-alt":
		return ChurchAlt(), nil
	case "cinema-old":
		return CinemaOld(), nil
	case "circle":
		return Circle(), nil
	case "city":
		return City(), nil
	case "clean-water":
		return CleanWater(), nil
	case "clipboard-check":
		return ClipboardCheck(), nil
	case "clock":
		return Clock(), nil
	case "closed-captions":
		return ClosedCaptions(), nil
	case "closet":
		return Closet(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-book-alt":
		return CloudBookAlt(), nil
	case "cloud-check":
		return CloudCheck(), nil
	case "cloud-desync":
		return CloudDesync(), nil
	case "cloud-download":
		return CloudDownload(), nil
	case "cloud-error":
		return CloudError(), nil
	case "cloud-sunny":
		return CloudSunny(), nil
	case "cloud-sync":
		return CloudSync(), nil
	case "cloud-upload":
		return CloudUpload(), nil
	case "clutery":
		return Clutery(), nil
	case "code":
		return Code(), nil
	case "code-brackets":
		return CodeBrackets(), nil
	case "code-brackets-square":
		return CodeBracketsSquare(), nil
	case "codepen":
		return Codepen(), nil
	case "coffee-cup":
		return CoffeeCup(), nil
	case "coin":
		return Coin(), nil
	case "collage-frame":
		return CollageFrame(), nil
	case "collapse":
		return Collapse(), nil
	case "color-filter":
		return ColorFilter(), nil
	case "color-picker":
		return ColorPicker(), nil
	case "color-picker-empty":
		return ColorPickerEmpty(), nil
	case "combine":
		return Combine(), nil
	case "community":
		return Community(), nil
	case "compact-disc":
		return CompactDisc(), nil
	case "compass":
		return Compass(), nil
	case "compress":
		return Compress(), nil
	case "compress-lines":
		return CompressLines(), nil
	case "computer":
		return Computer(), nil
	case "consumable":
		return Consumable(), nil
	case "control-slider":
		return ControlSlider(), nil
	case "cookie":
		return Cookie(), nil
	case "cooling":
		return Cooling(), nil
	case "copy":
		return Copy(), nil
	case "copyright":
		return Copyright(), nil
	case "corner-bottom-left":
		return CornerBottomLeft(), nil
	case "corner-bottom-right":
		return CornerBottomRight(), nil
	case "corner-top-left":
		return CornerTopLeft(), nil
	case "corner-top-right":
		return CornerTopRight(), nil
	case "cpu":
		return Cpu(), nil
	case "cpu-warning":
		return CpuWarning(), nil
	case "cracked-egg":
		return CrackedEgg(), nil
	case "creative-commons":
		return CreativeCommons(), nil
	case "credit-card":
		return CreditCard(), nil
	case "credit-card-2":
		return CreditCardTwo(), nil
	case "credit-cards":
		return CreditCards(), nil
	case "crib":
		return Crib(), nil
	case "crop":
		return Crop(), nil
	case "crop-rotate-bl":
		return CropRotateBl(), nil
	case "crop-rotate-br":
		return CropRotateBr(), nil
	case "crop-rotate-tl":
		return CropRotateTl(), nil
	case "crop-rotate-tr":
		return CropRotateTr(), nil
	case "crown":
		return Crown(), nil
	case "crown-circle":
		return CrownCircle(), nil
	case "css3":
		return CssThree(), nil
	case "cursor-pointer":
		return CursorPointer(), nil
	case "cut":
		return Cut(), nil
	case "cut-alt":
		return CutAlt(), nil
	case "cycling":
		return Cycling(), nil
	case "cylinder":
		return Cylinder(), nil
	case "dash-flag":
		return DashFlag(), nil
	case "dashboard":
		return Dashboard(), nil
	case "dashboard-dots":
		return DashboardDots(), nil
	case "dashboard-speed":
		return DashboardSpeed(), nil
	case "data-transfer-both":
		return DataTransferBoth(), nil
	case "data-transfer-check":
		return DataTransferCheck(), nil
	case "data-transfer-down":
		return DataTransferDown(), nil
	case "data-transfer-up":
		return DataTransferUp(), nil
	case "data-transfer-warning":
		return DataTransferWarning(), nil
	case "database-backup":
		return DatabaseBackup(), nil
	case "database-export":
		return DatabaseExport(), nil
	case "database-monitor":
		return DatabaseMonitor(), nil
	case "database-restore":
		return DatabaseRestore(), nil
	case "database-script":
		return DatabaseScript(), nil
	case "database-settings":
		return DatabaseSettings(), nil
	case "database-star":
		return DatabaseStar(), nil
	case "database-stats":
		return DatabaseStats(), nil
	case "database-tag":
		return DatabaseTag(), nil
	case "db":
		return Db(), nil
	case "db-check":
		return DbCheck(), nil
	case "db-error":
		return DbError(), nil
	case "db-search":
		return DbSearch(), nil
	case "db-star":
		return DbStar(), nil
	case "db-warning":
		return DbWarning(), nil
	case "de-compress":
		return DeCompress(), nil
	case "delete-circle":
		return DeleteCircle(), nil
	case "delivery":
		return Delivery(), nil
	case "delivery-truck":
		return DeliveryTruck(), nil
	case "depth":
		return Depth(), nil
	case "design-nib":
		return DesignNib(), nil
	case "design-pencil":
		return DesignPencil(), nil
	case "desk":
		return Desk(), nil
	case "dialpad":
		return Dialpad(), nil
	case "diameter":
		return Diameter(), nil
	case "dice-five":
		return DiceFive(), nil
	case "dice-four":
		return DiceFour(), nil
	case "dice-one":
		return DiceOne(), nil
	case "dice-six":
		return DiceSix(), nil
	case "dice-three":
		return DiceThree(), nil
	case "dice-two":
		return DiceTwo(), nil
	case "dimmer-switch":
		return DimmerSwitch(), nil
	case "director-chair":
		return DirectorChair(), nil
	case "discord":
		return Discord(), nil
	case "dishwasher":
		return Dishwasher(), nil
	case "divide":
		return Divide(), nil
	case "divide-selection-1":
		return DivideSelectionOne(), nil
	case "divide-selection-2":
		return DivideSelectionTwo(), nil
	case "divide-three":
		return DivideThree(), nil
	case "dna":
		return Dna(), nil
	case "doc-search":
		return DocSearch(), nil
	case "doc-search-alt":
		return DocSearchAlt(), nil
	case "doc-star":
		return DocStar(), nil
	case "doc-star-alt":
		return DocStarAlt(), nil
	case "dollar":
		return Dollar(), nil
	case "domotic-issue":
		return DomoticIssue(), nil
	case "donate":
		return Donate(), nil
	case "double-check":
		return DoubleCheck(), nil
	case "down-round-arrow":
		return DownRoundArrow(), nil
	case "download":
		return Download(), nil
	case "download-circle":
		return DownloadCircle(), nil
	case "download-data-window":
		return DownloadDataWindow(), nil
	case "download-square":
		return DownloadSquare(), nil
	case "drag":
		return Drag(), nil
	case "drag-hand-gesture":
		return DragHandGesture(), nil
	case "drawer":
		return Drawer(), nil
	case "dribbble":
		return Dribbble(), nil
	case "drone":
		return Drone(), nil
	case "drone-charge-full":
		return DroneChargeFull(), nil
	case "drone-charge-half":
		return DroneChargeHalf(), nil
	case "drone-charge-low":
		return DroneChargeLow(), nil
	case "drone-check":
		return DroneCheck(), nil
	case "drone-error":
		return DroneError(), nil
	case "drone-landing":
		return DroneLanding(), nil
	case "drone-refresh":
		return DroneRefresh(), nil
	case "drone-take-off":
		return DroneTakeOff(), nil
	case "droplet":
		return Droplet(), nil
	case "droplet-half":
		return DropletHalf(), nil
	case "ease-curve-control-points":
		return EaseCurveControlPoints(), nil
	case "ease-in":
		return EaseIn(), nil
	case "ease-in-control-point":
		return EaseInControlPoint(), nil
	case "ease-in-out":
		return EaseInOut(), nil
	case "ease-out":
		return EaseOut(), nil
	case "ease-out-control-point":
		return EaseOutControlPoint(), nil
	case "ecology-book":
		return EcologyBook(), nil
	case "edit":
		return Edit(), nil
	case "edit-pencil":
		return EditPencil(), nil
	case "egg":
		return Egg(), nil
	case "eject":
		return Eject(), nil
	case "electronics-chip":
		return ElectronicsChip(), nil
	case "electronics-transistor":
		return ElectronicsTransistor(), nil
	case "emoji":
		return Emoji(), nil
	case "emoji-ball":
		return EmojiBall(), nil
	case "emoji-blink-left":
		return EmojiBlinkLeft(), nil
	case "emoji-blink-right":
		return EmojiBlinkRight(), nil
	case "emoji-look-down":
		return EmojiLookDown(), nil
	case "emoji-look-left":
		return EmojiLookLeft(), nil
	case "emoji-look-right":
		return EmojiLookRight(), nil
	case "emoji-look-up":
		return EmojiLookUp(), nil
	case "emoji-puzzled":
		return EmojiPuzzled(), nil
	case "emoji-quite":
		return EmojiQuite(), nil
	case "emoji-really":
		return EmojiReally(), nil
	case "emoji-sad":
		return EmojiSad(), nil
	case "emoji-satisfied":
		return EmojiSatisfied(), nil
	case "emoji-sing-left":
		return EmojiSingLeft(), nil
	case "emoji-sing-left-note":
		return EmojiSingLeftNote(), nil
	case "emoji-sing-right":
		return EmojiSingRight(), nil
	case "emoji-sing-right-note":
		return EmojiSingRightNote(), nil
	case "emoji-surprise":
		return EmojiSurprise(), nil
	case "emoji-surprise-alt":
		return EmojiSurpriseAlt(), nil
	case "emoji-talking-angry":
		return EmojiTalkingAngry(), nil
	case "emoji-talking-happy":
		return EmojiTalkingHappy(), nil
	case "emoji-think-left":
		return EmojiThinkLeft(), nil
	case "emoji-think-right":
		return EmojiThinkRight(), nil
	case "empty-page":
		return EmptyPage(), nil
	case "energy-usage-window":
		return EnergyUsageWindow(), nil
	case "enlarge":
		return Enlarge(), nil
	case "enlarge-round-arrow":
		return EnlargeRoundArrow(), nil
	case "erase":
		return Erase(), nil
	case "error-window":
		return ErrorWindow(), nil
	case "euro":
		return Euro(), nil
	case "euro-square":
		return EuroSquare(), nil
	case "ev-charge":
		return EvCharge(), nil
	case "ev-charge-alt":
		return EvChargeAlt(), nil
	case "ev-plug":
		return EvPlug(), nil
	case "ev-plug-charging":
		return EvPlugCharging(), nil
	case "ev-plug-error":
		return EvPlugError(), nil
	case "ev-station":
		return EvStation(), nil
	case "ev-tag":
		return EvTag(), nil
	case "exclude":
		return Exclude(), nil
	case "expand":
		return Expand(), nil
	case "expand-lines":
		return ExpandLines(), nil
	case "eye-alt":
		return EyeAlt(), nil
	case "eye-close":
		return EyeClose(), nil
	case "eye-empty":
		return EyeEmpty(), nil
	case "eye-off":
		return EyeOff(), nil
	case "face-id":
		return FaceId(), nil
	case "facebook":
		return Facebook(), nil
	case "facebook-tag":
		return FacebookTag(), nil
	case "facetime":
		return Facetime(), nil
	case "farm":
		return Farm(), nil
	case "fast-arrow-down":
		return FastArrowDown(), nil
	case "fast-arrow-down-box":
		return FastArrowDownBox(), nil
	case "fast-arrow-left":
		return FastArrowLeft(), nil
	case "fast-arrow-left-box":
		return FastArrowLeftBox(), nil
	case "fast-arrow-right":
		return FastArrowRight(), nil
	case "fast-arrow-right-box":
		return FastArrowRightBox(), nil
	case "fast-arrow-up":
		return FastArrowUp(), nil
	case "fast-arrow-up-box":
		return FastArrowUpBox(), nil
	case "fast-down-circle":
		return FastDownCircle(), nil
	case "fast-left-circle":
		return FastLeftCircle(), nil
	case "fast-right-circle":
		return FastRightCircle(), nil
	case "fast-up-circle":
		return FastUpCircle(), nil
	case "favourite-book":
		return FavouriteBook(), nil
	case "favourite-window":
		return FavouriteWindow(), nil
	case "female":
		return Female(), nil
	case "figma":
		return Figma(), nil
	case "file-not-found":
		return FileNotFound(), nil
	case "filter":
		return Filter(), nil
	case "filter-alt":
		return FilterAlt(), nil
	case "finder":
		return Finder(), nil
	case "fingerprint":
		return Fingerprint(), nil
	case "fingerprint-check-circle":
		return FingerprintCheckCircle(), nil
	case "fingerprint-circle":
		return FingerprintCircle(), nil
	case "fingerprint-error-circle":
		return FingerprintErrorCircle(), nil
	case "fingerprint-lock-circle":
		return FingerprintLockCircle(), nil
	case "fingerprint-phone":
		return FingerprintPhone(), nil
	case "fingerprint-scan":
		return FingerprintScan(), nil
	case "fingerprint-square":
		return FingerprintSquare(), nil
	case "fingerprint-window":
		return FingerprintWindow(), nil
	case "fire-flame":
		return FireFlame(), nil
	case "fishing":
		return Fishing(), nil
	case "flare":
		return Flare(), nil
	case "flash":
		return Flash(), nil
	case "flash-off":
		return FlashOff(), nil
	case "flask":
		return Flask(), nil
	case "flip":
		return Flip(), nil
	case "flip-reverse":
		return FlipReverse(), nil
	case "flower":
		return Flower(), nil
	case "fluorine":
		return Fluorine(), nil
	case "fog":
		return Fog(), nil
	case "folder":
		return Folder(), nil
	case "folder-alert":
		return FolderAlert(), nil
	case "folder-settings":
		return FolderSettings(), nil
	case "font-size":
		return FontSize(), nil
	case "football":
		return Football(), nil
	case "football-ball":
		return FootballBall(), nil
	case "forward":
		return Forward(), nil
	case "forward-15-seconds":
		return ForwardFifteenSeconds(), nil
	case "forward-message":
		return ForwardMessage(), nil
	case "4k-display":
		return FourKDisplay(), nil
	case "frame":
		return Frame(), nil
	case "frame-alt":
		return FrameAlt(), nil
	case "frame-alt-empty":
		return FrameAltEmpty(), nil
	case "frame-select":
		return FrameSelect(), nil
	case "frame-simple":
		return FrameSimple(), nil
	case "frame-tool":
		return FrameTool(), nil
	case "fridge":
		return Fridge(), nil
	case "fx":
		return Fx(), nil
	case "fx-tag":
		return FxTag(), nil
	case "gamepad":
		return Gamepad(), nil
	case "garage":
		return Garage(), nil
	case "gas":
		return Gas(), nil
	case "gas-tank":
		return GasTank(), nil
	case "gas-tank-drop":
		return GasTankDrop(), nil
	case "gif-format":
		return GifFormat(), nil
	case "gift":
		return Gift(), nil
	case "git-branch":
		return GitBranch(), nil
	case "git-cherry-pick-commit":
		return GitCherryPickCommit(), nil
	case "git-command":
		return GitCommand(), nil
	case "git-commit":
		return GitCommit(), nil
	case "git-compare":
		return GitCompare(), nil
	case "git-fork":
		return GitFork(), nil
	case "git-merge":
		return GitMerge(), nil
	case "git-pull-request":
		return GitPullRequest(), nil
	case "git-pull-request-closed":
		return GitPullRequestClosed(), nil
	case "github":
		return Github(), nil
	case "github-circle":
		return GithubCircle(), nil
	case "gitlab-full":
		return GitlabFull(), nil
	case "glass-empty":
		return GlassEmpty(), nil
	case "glass-half":
		return GlassHalf(), nil
	case "glass-half-alt":
		return GlassHalfAlt(), nil
	case "glasses":
		return Glasses(), nil
	case "globe":
		return Globe(), nil
	case "golf":
		return Golf(), nil
	case "google":
		return Google(), nil
	case "google-circle":
		return GoogleCircle(), nil
	case "google-docs":
		return GoogleDocs(), nil
	case "google-drive":
		return GoogleDrive(), nil
	case "google-drive-check":
		return GoogleDriveCheck(), nil
	case "google-drive-sync":
		return GoogleDriveSync(), nil
	case "google-drive-warning":
		return GoogleDriveWarning(), nil
	case "google-home":
		return GoogleHome(), nil
	case "google-one":
		return GoogleOne(), nil
	case "gps":
		return Gps(), nil
	case "graduation-cap":
		return GraduationCap(), nil
	case "graph-down":
		return GraphDown(), nil
	case "graph-up":
		return GraphUp(), nil
	case "green-bus":
		return GreenBus(), nil
	case "green-truck":
		return GreenTruck(), nil
	case "green-vehicle":
		return GreenVehicle(), nil
	case "grid-add":
		return GridAdd(), nil
	case "grid-minus":
		return GridMinus(), nil
	case "grid-remove":
		return GridRemove(), nil
	case "group":
		return Group(), nil
	case "gym":
		return Gym(), nil
	case "half-cookie":
		return HalfCookie(), nil
	case "half-moon":
		return HalfMoon(), nil
	case "hammer":
		return Hammer(), nil
	case "hand-brake":
		return HandBrake(), nil
	case "handbag":
		return Handbag(), nil
	case "hard-drive":
		return HardDrive(), nil
	case "hat":
		return Hat(), nil
	case "hd":
		return Hd(), nil
	case "hd-display":
		return HdDisplay(), nil
	case "hdr":
		return Hdr(), nil
	case "headset":
		return Headset(), nil
	case "headset-charge":
		return HeadsetCharge(), nil
	case "headset-help":
		return HeadsetHelp(), nil
	case "headset-issue":
		return HeadsetIssue(), nil
	case "health-shield":
		return HealthShield(), nil
	case "healthcare":
		return Healthcare(), nil
	case "heart":
		return Heart(), nil
	case "heating":
		return Heating(), nil
	case "heavy-rain":
		return HeavyRain(), nil
	case "help-circle":
		return HelpCircle(), nil
	case "help-square":
		return HelpSquare(), nil
	case "heptagon":
		return Heptagon(), nil
	case "her-slips":
		return HerSlips(), nil
	case "hexagon":
		return Hexagon(), nil
	case "hexagon-alt":
		return HexagonAlt(), nil
	case "hexagon-dice":
		return HexagonDice(), nil
	case "high-priority":
		return HighPriority(), nil
	case "historic-shield":
		return HistoricShield(), nil
	case "historic-shield-alt":
		return HistoricShieldAlt(), nil
	case "home":
		return Home(), nil
	case "home-alt":
		return HomeAlt(), nil
	case "home-alt-slim":
		return HomeAltSlim(), nil
	case "home-alt-slim-horiz":
		return HomeAltSlimHoriz(), nil
	case "home-hospital":
		return HomeHospital(), nil
	case "home-sale":
		return HomeSale(), nil
	case "home-secure":
		return HomeSecure(), nil
	case "home-shield":
		return HomeShield(), nil
	case "home-simple":
		return HomeSimple(), nil
	case "home-simple-door":
		return HomeSimpleDoor(), nil
	case "home-table":
		return HomeTable(), nil
	case "home-user":
		return HomeUser(), nil
	case "horiz-distribution-left":
		return HorizDistributionLeft(), nil
	case "horiz-distribution-right":
		return HorizDistributionRight(), nil
	case "horizontal-merge":
		return HorizontalMerge(), nil
	case "horizontal-split":
		return HorizontalSplit(), nil
	case "hospital":
		return Hospital(), nil
	case "hospital-sign":
		return HospitalSign(), nil
	case "hot-air-balloon":
		return HotAirBalloon(), nil
	case "hourglass":
		return Hourglass(), nil
	case "html5":
		return HtmlFive(), nil
	case "hydrogen":
		return Hydrogen(), nil
	case "iconoir":
		return Iconoir(), nil
	case "import":
		return Import(), nil
	case "inclination":
		return Inclination(), nil
	case "industry":
		return Industry(), nil
	case "infinite":
		return Infinite(), nil
	case "info-empty":
		return InfoEmpty(), nil
	case "input-field":
		return InputField(), nil
	case "input-output":
		return InputOutput(), nil
	case "input-search":
		return InputSearch(), nil
	case "instagram":
		return Instagram(), nil
	case "internet":
		return Internet(), nil
	case "intersect":
		return Intersect(), nil
	case "intersect-alt":
		return IntersectAlt(), nil
	case "ios-settings":
		return IosSettings(), nil
	case "ip-address":
		return IpAddress(), nil
	case "iris-scan":
		return IrisScan(), nil
	case "italic":
		return Italic(), nil
	case "italic-square":
		return ItalicSquare(), nil
	case "journal":
		return Journal(), nil
	case "journal-page":
		return JournalPage(), nil
	case "jpeg-format":
		return JpegFormat(), nil
	case "jpg-format":
		return JpgFormat(), nil
	case "kanban-board":
		return KanbanBoard(), nil
	case "key-alt":
		return KeyAlt(), nil
	case "key-alt-back":
		return KeyAltBack(), nil
	case "key-alt-minus":
		return KeyAltMinus(), nil
	case "key-alt-plus":
		return KeyAltPlus(), nil
	case "key-alt-remove":
		return KeyAltRemove(), nil
	case "key-command":
		return KeyCommand(), nil
	case "keyframe":
		return Keyframe(), nil
	case "keyframe-align-center":
		return KeyframeAlignCenter(), nil
	case "keyframe-align-horizontal":
		return KeyframeAlignHorizontal(), nil
	case "keyframe-align-vertical":
		return KeyframeAlignVertical(), nil
	case "keyframe-position":
		return KeyframePosition(), nil
	case "keyframes":
		return Keyframes(), nil
	case "keyframes-couple":
		return KeyframesCouple(), nil
	case "label":
		return Label(), nil
	case "lamp":
		return Lamp(), nil
	case "language":
		return Language(), nil
	case "laptop":
		return Laptop(), nil
	case "laptop-charging":
		return LaptopCharging(), nil
	case "laptop-fix":
		return LaptopFix(), nil
	case "laptop-issue":
		return LaptopIssue(), nil
	case "large-suitcase":
		return LargeSuitcase(), nil
	case "layout-left":
		return LayoutLeft(), nil
	case "layout-right":
		return LayoutRight(), nil
	case "leaderboard":
		return Leaderboard(), nil
	case "leaderboard-star":
		return LeaderboardStar(), nil
	case "leaf":
		return Leaf(), nil
	case "learning":
		return Learning(), nil
	case "left-round-arrow":
		return LeftRoundArrow(), nil
	case "lens":
		return Lens(), nil
	case "lifebelt":
		return Lifebelt(), nil
	case "light-bulb":
		return LightBulb(), nil
	case "light-bulb-off":
		return LightBulbOff(), nil
	case "light-bulb-on":
		return LightBulbOn(), nil
	case "line-space":
		return LineSpace(), nil
	case "linear":
		return Linear(), nil
	case "link":
		return Link(), nil
	case "linkedin":
		return Linkedin(), nil
	case "linux":
		return Linux(), nil
	case "list":
		return List(), nil
	case "list-select":
		return ListSelect(), nil
	case "load-action-floppy":
		return LoadActionFloppy(), nil
	case "lock":
		return Lock(), nil
	case "lock-key":
		return LockKey(), nil
	case "locked-book":
		return LockedBook(), nil
	case "locked-window":
		return LockedWindow(), nil
	case "log-denied":
		return LogDenied(), nil
	case "log-in":
		return LogIn(), nil
	case "log-out":
		return LogOut(), nil
	case "long-arrow-down-left":
		return LongArrowDownLeft(), nil
	case "long-arrow-down-right":
		return LongArrowDownRight(), nil
	case "long-arrow-left-down":
		return LongArrowLeftDown(), nil
	case "long-arrow-left-up":
		return LongArrowLeftUp(), nil
	case "long-arrow-right-down":
		return LongArrowRightDown(), nil
	case "long-arrow-right-up":
		return LongArrowRightUp(), nil
	case "long-arrow-right-up-1":
		return LongArrowRightUpOne(), nil
	case "long-arrow-up-left":
		return LongArrowUpLeft(), nil
	case "long-arrow-up-right":
		return LongArrowUpRight(), nil
	case "lot-of-cash":
		return LotOfCash(), nil
	case "mac-control-key":
		return MacControlKey(), nil
	case "mac-dock":
		return MacDock(), nil
	case "mac-option-key":
		return MacOptionKey(), nil
	case "mac-os-window":
		return MacOsWindow(), nil
	case "magnet":
		return Magnet(), nil
	case "magnet-energy":
		return MagnetEnergy(), nil
	case "mail":
		return Mail(), nil
	case "mail-in":
		return MailIn(), nil
	case "mail-opened":
		return MailOpened(), nil
	case "mail-out":
		return MailOut(), nil
	case "male":
		return Male(), nil
	case "map":
		return Map(), nil
	case "map-issue":
		return MapIssue(), nil
	case "maps-arrow":
		return MapsArrow(), nil
	case "maps-arrow-diagonal":
		return MapsArrowDiagonal(), nil
	case "maps-arrow-issue":
		return MapsArrowIssue(), nil
	case "maps-go-straight":
		return MapsGoStraight(), nil
	case "maps-turn-back":
		return MapsTurnBack(), nil
	case "maps-turn-left":
		return MapsTurnLeft(), nil
	case "maps-turn-right":
		return MapsTurnRight(), nil
	case "mask-square":
		return MaskSquare(), nil
	case "mastercard-card":
		return MastercardCard(), nil
	case "math-book":
		return MathBook(), nil
	case "maximize":
		return Maximize(), nil
	case "medal":
		return Medal(), nil
	case "media-image":
		return MediaImage(), nil
	case "media-image-folder":
		return MediaImageFolder(), nil
	case "media-image-list":
		return MediaImageList(), nil
	case "media-video":
		return MediaVideo(), nil
	case "media-video-folder":
		return MediaVideoFolder(), nil
	case "media-video-list":
		return MediaVideoList(), nil
	case "medium":
		return Medium(), nil
	case "medium-priority":
		return MediumPriority(), nil
	case "megaphone":
		return Megaphone(), nil
	case "menu":
		return Menu(), nil
	case "menu-scale":
		return MenuScale(), nil
	case "message":
		return Message(), nil
	case "message-alert":
		return MessageAlert(), nil
	case "message-text":
		return MessageText(), nil
	case "metro":
		return Metro(), nil
	case "mic":
		return Mic(), nil
	case "mic-add":
		return MicAdd(), nil
	case "mic-check":
		return MicCheck(), nil
	case "mic-mute":
		return MicMute(), nil
	case "mic-remove":
		return MicRemove(), nil
	case "mic-speaking":
		return MicSpeaking(), nil
	case "mic-warning":
		return MicWarning(), nil
	case "microscope":
		return Microscope(), nil
	case "minus":
		return Minus(), nil
	case "minus-circle":
		return MinusCircle(), nil
	case "minus-hexagon":
		return MinusHexagon(), nil
	case "minus-1":
		return MinusOne(), nil
	case "minus-pin-alt":
		return MinusPinAlt(), nil
	case "minus-square":
		return MinusSquare(), nil
	case "mirror":
		return Mirror(), nil
	case "missing-font":
		return MissingFont(), nil
	case "modern-tv":
		return ModernTv(), nil
	case "modern-tv-4k":
		return ModernTvFourK(), nil
	case "money-square":
		return MoneySquare(), nil
	case "moon-sat":
		return MoonSat(), nil
	case "more-horiz":
		return MoreHoriz(), nil
	case "more-horiz-circle":
		return MoreHorizCircle(), nil
	case "more-vert":
		return MoreVert(), nil
	case "more-vert-circle":
		return MoreVertCircle(), nil
	case "motorcycle":
		return Motorcycle(), nil
	case "mouse-button-left":
		return MouseButtonLeft(), nil
	case "mouse-button-right":
		return MouseButtonRight(), nil
	case "mouse-scroll-wheel":
		return MouseScrollWheel(), nil
	case "move-down":
		return MoveDown(), nil
	case "move-left":
		return MoveLeft(), nil
	case "move-right":
		return MoveRight(), nil
	case "move-ruler":
		return MoveRuler(), nil
	case "move-up":
		return MoveUp(), nil
	case "movie":
		return Movie(), nil
	case "mpeg-format":
		return MpegFormat(), nil
	case "multi-bubble":
		return MultiBubble(), nil
	case "multi-mac-os-window":
		return MultiMacOsWindow(), nil
	case "multi-window":
		return MultiWindow(), nil
	case "multiple-pages":
		return MultiplePages(), nil
	case "multiple-pages-add":
		return MultiplePagesAdd(), nil
	case "multiple-pages-delete":
		return MultiplePagesDelete(), nil
	case "multiple-pages-empty":
		return MultiplePagesEmpty(), nil
	case "multiple-pages-remove":
		return MultiplePagesRemove(), nil
	case "music-double-note":
		return MusicDoubleNote(), nil
	case "music-double-note-add":
		return MusicDoubleNoteAdd(), nil
	case "music-note":
		return MusicNote(), nil
	case "music-note-add":
		return MusicNoteAdd(), nil
	case "nav-arrow-down":
		return NavArrowDown(), nil
	case "nav-arrow-left":
		return NavArrowLeft(), nil
	case "nav-arrow-right":
		return NavArrowRight(), nil
	case "nav-arrow-up":
		return NavArrowUp(), nil
	case "navigator":
		return Navigator(), nil
	case "navigator-alt":
		return NavigatorAlt(), nil
	case "network":
		return Network(), nil
	case "network-alt":
		return NetworkAlt(), nil
	case "network-left":
		return NetworkLeft(), nil
	case "network-right":
		return NetworkRight(), nil
	case "new-tab":
		return NewTab(), nil
	case "nintendo-switch":
		return NintendoSwitch(), nil
	case "nitrogen":
		return Nitrogen(), nil
	case "no-access-window":
		return NoAccessWindow(), nil
	case "no-battery":
		return NoBattery(), nil
	case "no-coin":
		return NoCoin(), nil
	case "no-credit-card":
		return NoCreditCard(), nil
	case "no-link":
		return NoLink(), nil
	case "no-lock":
		return NoLock(), nil
	case "no-smoking":
		return NoSmoking(), nil
	case "no-smoking-circle":
		return NoSmokingCircle(), nil
	case "notes":
		return Notes(), nil
	case "npm":
		return Npm(), nil
	case "npm-square":
		return NpmSquare(), nil
	case "numbered-list-left":
		return NumberedListLeft(), nil
	case "numbered-list-right":
		return NumberedListRight(), nil
	case "octagon":
		return Octagon(), nil
	case "off-tag":
		return OffTag(), nil
	case "oil-industry":
		return OilIndustry(), nil
	case "okrs":
		return Okrs(), nil
	case "on-tag":
		return OnTag(), nil
	case "one-finger-select-hand-gesture":
		return OneFingerSelectHandGesture(), nil
	case "one-point-circle":
		return OnePointCircle(), nil
	case "1st-medal":
		return OneStMedal(), nil
	case "open-book":
		return OpenBook(), nil
	case "open-in-browser":
		return OpenInBrowser(), nil
	case "open-in-window":
		return OpenInWindow(), nil
	case "open-new-window":
		return OpenNewWindow(), nil
	case "open-select-hand-gesture":
		return OpenSelectHandGesture(), nil
	case "open-vpn":
		return OpenVpn(), nil
	case "orange-half":
		return OrangeHalf(), nil
	case "orange-slice":
		return OrangeSlice(), nil
	case "orange-slice-alt":
		return OrangeSliceAlt(), nil
	case "organic-food":
		return OrganicFood(), nil
	case "organic-food-square":
		return OrganicFoodSquare(), nil
	case "orthogonal-view":
		return OrthogonalView(), nil
	case "oxygen":
		return Oxygen(), nil
	case "package":
		return Package(), nil
	case "package-lock":
		return PackageLock(), nil
	case "packages":
		return Packages(), nil
	case "pacman":
		return Pacman(), nil
	case "page":
		return Page(), nil
	case "page-down":
		return PageDown(), nil
	case "page-edit":
		return PageEdit(), nil
	case "page-flip":
		return PageFlip(), nil
	case "page-left":
		return PageLeft(), nil
	case "page-right":
		return PageRight(), nil
	case "page-search":
		return PageSearch(), nil
	case "page-star":
		return PageStar(), nil
	case "page-up":
		return PageUp(), nil
	case "palette":
		return Palette(), nil
	case "panorama-enlarge":
		return PanoramaEnlarge(), nil
	case "panorama-reduce":
		return PanoramaReduce(), nil
	case "pants":
		return Pants(), nil
	case "pants-alt":
		return PantsAlt(), nil
	case "parking":
		return Parking(), nil
	case "password-cursor":
		return PasswordCursor(), nil
	case "password-error":
		return PasswordError(), nil
	case "password-pass":
		return PasswordPass(), nil
	case "paste-clipboard":
		return PasteClipboard(), nil
	case "path-arrow":
		return PathArrow(), nil
	case "pause":
		return Pause(), nil
	case "pause-window":
		return PauseWindow(), nil
	case "paypal":
		return Paypal(), nil
	case "pc-check":
		return PcCheck(), nil
	case "pc-firewall":
		return PcFirewall(), nil
	case "pc-mouse":
		return PcMouse(), nil
	case "pc-no-entry":
		return PcNoEntry(), nil
	case "pc-warning":
		return PcWarning(), nil
	case "peace-hand":
		return PeaceHand(), nil
	case "pen-connect-bluetooth":
		return PenConnectBluetooth(), nil
	case "pen-connect-wifi":
		return PenConnectWifi(), nil
	case "pen-tablet":
		return PenTablet(), nil
	case "pen-tablet-connect-usb":
		return PenTabletConnectUsb(), nil
	case "pen-tablet-connect-wifi":
		return PenTabletConnectWifi(), nil
	case "pentagon":
		return Pentagon(), nil
	case "people-tag":
		return PeopleTag(), nil
	case "percentage":
		return Percentage(), nil
	case "percentage-circle":
		return PercentageCircle(), nil
	case "percentage-square":
		return PercentageSquare(), nil
	case "perspective-view":
		return PerspectiveView(), nil
	case "pharmacy-cross-circle":
		return PharmacyCrossCircle(), nil
	case "pharmacy-cross-square":
		return PharmacyCrossSquare(), nil
	case "phone":
		return Phone(), nil
	case "phone-add":
		return PhoneAdd(), nil
	case "phone-delete":
		return PhoneDelete(), nil
	case "phone-disabled":
		return PhoneDisabled(), nil
	case "phone-income":
		return PhoneIncome(), nil
	case "phone-outcome":
		return PhoneOutcome(), nil
	case "phone-paused":
		return PhonePaused(), nil
	case "phone-remove":
		return PhoneRemove(), nil
	case "piggy-bank":
		return PiggyBank(), nil
	case "pillow":
		return Pillow(), nil
	case "pin":
		return Pin(), nil
	case "pin-alt":
		return PinAlt(), nil
	case "pine-tree":
		return PineTree(), nil
	case "pinterest":
		return Pinterest(), nil
	case "pizza-slice":
		return PizzaSlice(), nil
	case "planet":
		return Planet(), nil
	case "planet-alt":
		return PlanetAlt(), nil
	case "planet-sat":
		return PlanetSat(), nil
	case "play":
		return Play(), nil
	case "playlist":
		return Playlist(), nil
	case "playlist-add":
		return PlaylistAdd(), nil
	case "playlist-play":
		return PlaylistPlay(), nil
	case "playstation-gamepad":
		return PlaystationGamepad(), nil
	case "plug-type-a":
		return PlugTypeA(), nil
	case "plug-type-c":
		return PlugTypeC(), nil
	case "plug-type-g":
		return PlugTypeG(), nil
	case "plug-type-l":
		return PlugTypeL(), nil
	case "plus":
		return Plus(), nil
	case "png-format":
		return PngFormat(), nil
	case "pocket":
		return Pocket(), nil
	case "podcast":
		return Podcast(), nil
	case "pokeball":
		return Pokeball(), nil
	case "position":
		return Position(), nil
	case "position-align":
		return PositionAlign(), nil
	case "potion":
		return Potion(), nil
	case "pound":
		return Pound(), nil
	case "precision-tool":
		return PrecisionTool(), nil
	case "presentation":
		return Presentation(), nil
	case "printer":
		return Printer(), nil
	case "printer-alt":
		return PrinterAlt(), nil
	case "printing-page":
		return PrintingPage(), nil
	case "priority-down":
		return PriorityDown(), nil
	case "priority-up":
		return PriorityUp(), nil
	case "private-wifi":
		return PrivateWifi(), nil
	case "profile-circle":
		return ProfileCircle(), nil
	case "prohibition":
		return Prohibition(), nil
	case "puzzle":
		return Puzzle(), nil
	case "qr-code":
		return QrCode(), nil
	case "question-mark":
		return QuestionMark(), nil
	case "quote":
		return Quote(), nil
	case "quote-message":
		return QuoteMessage(), nil
	case "radiation":
		return Radiation(), nil
	case "radius":
		return Radius(), nil
	case "rain":
		return Rain(), nil
	case "raw-format":
		return RawFormat(), nil
	case "receive-dollars":
		return ReceiveDollars(), nil
	case "receive-euros":
		return ReceiveEuros(), nil
	case "receive-pounds":
		return ReceivePounds(), nil
	case "receive-yens":
		return ReceiveYens(), nil
	case "redo":
		return Redo(), nil
	case "redo-action":
		return RedoAction(), nil
	case "redo-circle":
		return RedoCircle(), nil
	case "reduce":
		return Reduce(), nil
	case "reduce-round-arrow":
		return ReduceRoundArrow(), nil
	case "refresh":
		return Refresh(), nil
	case "refresh-circular":
		return RefreshCircular(), nil
	case "refresh-double":
		return RefreshDouble(), nil
	case "reload-window":
		return ReloadWindow(), nil
	case "reminder-hand-gesture":
		return ReminderHandGesture(), nil
	case "remove-database-script":
		return RemoveDatabaseScript(), nil
	case "remove-folder":
		return RemoveFolder(), nil
	case "remove-frame":
		return RemoveFrame(), nil
	case "remove-from-cart":
		return RemoveFromCart(), nil
	case "remove-keyframe":
		return RemoveKeyframe(), nil
	case "remove-keyframe-alt":
		return RemoveKeyframeAlt(), nil
	case "remove-keyframes":
		return RemoveKeyframes(), nil
	case "remove-link":
		return RemoveLink(), nil
	case "remove-media-image":
		return RemoveMediaImage(), nil
	case "remove-media-video":
		return RemoveMediaVideo(), nil
	case "remove-page":
		return RemovePage(), nil
	case "remove-page-alt":
		return RemovePageAlt(), nil
	case "remove-pin":
		return RemovePin(), nil
	case "remove-pin-alt":
		return RemovePinAlt(), nil
	case "remove-selection":
		return RemoveSelection(), nil
	case "remove-square":
		return RemoveSquare(), nil
	case "remove-user":
		return RemoveUser(), nil
	case "repeat":
		return Repeat(), nil
	case "repeat-once":
		return RepeatOnce(), nil
	case "reply":
		return Reply(), nil
	case "reply-to-message":
		return ReplyToMessage(), nil
	case "report-columns":
		return ReportColumns(), nil
	case "reports":
		return Reports(), nil
	case "repository":
		return Repository(), nil
	case "restart":
		return Restart(), nil
	case "rewind":
		return Rewind(), nil
	case "rhombus":
		return Rhombus(), nil
	case "right-round-arrow":
		return RightRoundArrow(), nil
	case "rings":
		return Rings(), nil
	case "rocket":
		return Rocket(), nil
	case "rook":
		return Rook(), nil
	case "rotate-camera-left":
		return RotateCameraLeft(), nil
	case "rotate-camera-right":
		return RotateCameraRight(), nil
	case "round-flask":
		return RoundFlask(), nil
	case "rounded-mirror":
		return RoundedMirror(), nil
	case "rss-feed":
		return RssFeed(), nil
	case "rss-feed-tag":
		return RssFeedTag(), nil
	case "rubik-cube":
		return RubikCube(), nil
	case "ruler":
		return Ruler(), nil
	case "ruler-add":
		return RulerAdd(), nil
	case "ruler-combine":
		return RulerCombine(), nil
	case "ruler-remove":
		return RulerRemove(), nil
	case "running":
		return Running(), nil
	case "safari":
		return Safari(), nil
	case "sandals":
		return Sandals(), nil
	case "save-action-floppy":
		return SaveActionFloppy(), nil
	case "save-floppy-disk":
		return SaveFloppyDisk(), nil
	case "scale-frame-enlarge":
		return ScaleFrameEnlarge(), nil
	case "scale-frame-reduce":
		return ScaleFrameReduce(), nil
	case "scan-barcode":
		return ScanBarcode(), nil
	case "scan-qr-code":
		return ScanQrCode(), nil
	case "scanning":
		return Scanning(), nil
	case "scarf":
		return Scarf(), nil
	case "scissor":
		return Scissor(), nil
	case "scissor-alt":
		return ScissorAlt(), nil
	case "screenshot":
		return Screenshot(), nil
	case "sea-and-sun":
		return SeaAndSun(), nil
	case "sea-waves":
		return SeaWaves(), nil
	case "search":
		return Search(), nil
	case "search-engine":
		return SearchEngine(), nil
	case "search-font":
		return SearchFont(), nil
	case "search-window":
		return SearchWindow(), nil
	case "secure-window":
		return SecureWindow(), nil
	case "security-pass":
		return SecurityPass(), nil
	case "select-window":
		return SelectWindow(), nil
	case "selection":
		return Selection(), nil
	case "selective-tool":
		return SelectiveTool(), nil
	case "send":
		return Send(), nil
	case "send-diagonal":
		return SendDiagonal(), nil
	case "send-dollars":
		return SendDollars(), nil
	case "send-euros":
		return SendEuros(), nil
	case "send-mail":
		return SendMail(), nil
	case "send-pounds":
		return SendPounds(), nil
	case "send-yens":
		return SendYens(), nil
	case "server":
		return Server(), nil
	case "server-connection":
		return ServerConnection(), nil
	case "settings":
		return Settings(), nil
	case "settings-cloud":
		return SettingsCloud(), nil
	case "settings-profiles":
		return SettingsProfiles(), nil
	case "share-android":
		return ShareAndroid(), nil
	case "share-ios":
		return ShareIos(), nil
	case "shield":
		return Shield(), nil
	case "shield-add":
		return ShieldAdd(), nil
	case "shield-alert":
		return ShieldAlert(), nil
	case "shield-alt":
		return ShieldAlt(), nil
	case "shield-broken":
		return ShieldBroken(), nil
	case "shield-check":
		return ShieldCheck(), nil
	case "shield-cross":
		return ShieldCross(), nil
	case "shield-download":
		return ShieldDownload(), nil
	case "shield-eye":
		return ShieldEye(), nil
	case "shield-loading":
		return ShieldLoading(), nil
	case "shield-minus":
		return ShieldMinus(), nil
	case "shield-question":
		return ShieldQuestion(), nil
	case "shield-search":
		return ShieldSearch(), nil
	case "shield-upload":
		return ShieldUpload(), nil
	case "shop":
		return Shop(), nil
	case "shop-alt":
		return ShopAlt(), nil
	case "shopping-bag":
		return ShoppingBag(), nil
	case "shopping-bag-add":
		return ShoppingBagAdd(), nil
	case "shopping-bag-alt":
		return ShoppingBagAlt(), nil
	case "shopping-bag-arrow-down":
		return ShoppingBagArrowDown(), nil
	case "shopping-bag-arrow-up":
		return ShoppingBagArrowUp(), nil
	case "shopping-bag-check":
		return ShoppingBagCheck(), nil
	case "shopping-bag-issue":
		return ShoppingBagIssue(), nil
	case "shopping-bag-remove":
		return ShoppingBagRemove(), nil
	case "shopping-code":
		return ShoppingCode(), nil
	case "shopping-code-check":
		return ShoppingCodeCheck(), nil
	case "shopping-code-error":
		return ShoppingCodeError(), nil
	case "short-pants":
		return ShortPants(), nil
	case "short-pants-alt":
		return ShortPantsAlt(), nil
	case "shortcut":
		return Shortcut(), nil
	case "shuffle":
		return Shuffle(), nil
	case "sidebar-collapse":
		return SidebarCollapse(), nil
	case "sidebar-expand":
		return SidebarExpand(), nil
	case "sigma-function":
		return SigmaFunction(), nil
	case "simple-cart":
		return SimpleCart(), nil
	case "sine-wave":
		return SineWave(), nil
	case "single-tap-gesture":
		return SingleTapGesture(), nil
	case "skateboard":
		return Skateboard(), nil
	case "skateboarding":
		return Skateboarding(), nil
	case "skip-next":
		return SkipNext(), nil
	case "skip-prev":
		return SkipPrev(), nil
	case "sleeper-chair":
		return SleeperChair(), nil
	case "small-lamp":
		return SmallLamp(), nil
	case "small-lamp-alt":
		return SmallLampAlt(), nil
	case "small-shop":
		return SmallShop(), nil
	case "small-shop-alt":
		return SmallShopAlt(), nil
	case "smartphone-device":
		return SmartphoneDevice(), nil
	case "smoking":
		return Smoking(), nil
	case "snapchat":
		return Snapchat(), nil
	case "snow":
		return Snow(), nil
	case "snow-flake":
		return SnowFlake(), nil
	case "soap":
		return Soap(), nil
	case "soccer-ball":
		return SoccerBall(), nil
	case "sofa":
		return Sofa(), nil
	case "soil":
		return Soil(), nil
	case "soil-alt":
		return SoilAlt(), nil
	case "sort":
		return Sort(), nil
	case "sort-down":
		return SortDown(), nil
	case "sort-up":
		return SortUp(), nil
	case "sound-high":
		return SoundHigh(), nil
	case "sound-low":
		return SoundLow(), nil
	case "sound-min":
		return SoundMin(), nil
	case "sound-off":
		return SoundOff(), nil
	case "spades":
		return Spades(), nil
	case "sphere":
		return Sphere(), nil
	case "spiral":
		return Spiral(), nil
	case "spock-hand-gesture":
		return SpockHandGesture(), nil
	case "square":
		return Square(), nil
	case "square-wave":
		return SquareWave(), nil
	case "stackoverflow":
		return Stackoverflow(), nil
	case "star":
		return Star(), nil
	case "star-dashed":
		return StarDashed(), nil
	case "star-half-dashed":
		return StarHalfDashed(), nil
	case "stat-down":
		return StatDown(), nil
	case "stat-up":
		return StatUp(), nil
	case "stats-down-square":
		return StatsDownSquare(), nil
	case "stats-report":
		return StatsReport(), nil
	case "stats-up-square":
		return StatsUpSquare(), nil
	case "strategy":
		return Strategy(), nil
	case "stretching":
		return Stretching(), nil
	case "stroller":
		return Stroller(), nil
	case "style-border":
		return StyleBorder(), nil
	case "submit-document":
		return SubmitDocument(), nil
	case "substract":
		return Substract(), nil
	case "suggestion":
		return Suggestion(), nil
	case "sun-light":
		return SunLight(), nil
	case "svg-format":
		return SvgFormat(), nil
	case "swimming":
		return Swimming(), nil
	case "swipe-down-gesture":
		return SwipeDownGesture(), nil
	case "swipe-left-gesture":
		return SwipeLeftGesture(), nil
	case "swipe-right-gesture":
		return SwipeRightGesture(), nil
	case "swipe-two-fingers-down-gesture":
		return SwipeTwoFingersDownGesture(), nil
	case "swipe-two-fingers-left-gesture":
		return SwipeTwoFingersLeftGesture(), nil
	case "swipe-two-fingers-right-gesture":
		return SwipeTwoFingersRightGesture(), nil
	case "swipe-two-fingers-up-gesture":
		return SwipeTwoFingersUpGesture(), nil
	case "swipe-up-gesture":
		return SwipeUpGesture(), nil
	case "switch-off":
		return SwitchOff(), nil
	case "switch-on":
		return SwitchOn(), nil
	case "system-restart":
		return SystemRestart(), nil
	case "system-shut":
		return SystemShut(), nil
	case "t-shirt":
		return TShirt(), nil
	case "table":
		return Table(), nil
	case "table-rows":
		return TableRows(), nil
	case "table-2-columns":
		return TableTwoColumns(), nil
	case "task-list":
		return TaskList(), nil
	case "telegram":
		return Telegram(), nil
	case "telegram-circle":
		return TelegramCircle(), nil
	case "tennis-ball":
		return TennisBall(), nil
	case "tennis-ball-alt":
		return TennisBallAlt(), nil
	case "terminal":
		return Terminal(), nil
	case "terminal-tag":
		return TerminalTag(), nil
	case "test-tube":
		return TestTube(), nil
	case "text":
		return Text(), nil
	case "text-alt":
		return TextAlt(), nil
	case "text-box":
		return TextBox(), nil
	case "text-size":
		return TextSize(), nil
	case "3d-add-hole":
		return ThreeDAddHole(), nil
	case "3d-arc":
		return ThreeDArc(), nil
	case "3d-arc-center-pt":
		return ThreeDArcCenterPt(), nil
	case "3d-bridge":
		return ThreeDBridge(), nil
	case "3d-center-box":
		return ThreeDCenterBox(), nil
	case "3d-ellipse":
		return ThreeDEllipse(), nil
	case "3d-ellipse-three-pts":
		return ThreeDEllipseThreePts(), nil
	case "3d-pt-box":
		return ThreeDPtBox(), nil
	case "3d-rect-corner-to-corner":
		return ThreeDRectCornerToCorner(), nil
	case "3d-rect-from-center":
		return ThreeDRectFromCenter(), nil
	case "3d-rect-three-pts":
		return ThreeDRectThreePts(), nil
	case "3d-select-edge":
		return ThreeDSelectEdge(), nil
	case "3d-select-face":
		return ThreeDSelectFace(), nil
	case "3d-select-point":
		return ThreeDSelectPoint(), nil
	case "3d-select-solid":
		return ThreeDSelectSolid(), nil
	case "3d-three-pts-box":
		return ThreeDThreePtsBox(), nil
	case "360-view":
		return ThreeHundredSixtyView(), nil
	case "three-points-circle":
		return ThreePointsCircle(), nil
	case "three-stars":
		return ThreeStars(), nil
	case "thumbs-down":
		return ThumbsDown(), nil
	case "thumbs-up":
		return ThumbsUp(), nil
	case "thunderstorm":
		return Thunderstorm(), nil
	case "tif-format":
		return TifFormat(), nil
	case "tiff-format":
		return TiffFormat(), nil
	case "tiktok":
		return Tiktok(), nil
	case "timer":
		return Timer(), nil
	case "timer-off":
		return TimerOff(), nil
	case "tools":
		return Tools(), nil
	case "tournament":
		return Tournament(), nil
	case "tower":
		return Tower(), nil
	case "tower-check":
		return TowerCheck(), nil
	case "tower-no-access":
		return TowerNoAccess(), nil
	case "tower-warning":
		return TowerWarning(), nil
	case "trademark":
		return Trademark(), nil
	case "train":
		return Train(), nil
	case "tram":
		return Tram(), nil
	case "transition-down":
		return TransitionDown(), nil
	case "transition-left":
		return TransitionLeft(), nil
	case "transition-right":
		return TransitionRight(), nil
	case "transition-up":
		return TransitionUp(), nil
	case "translate":
		return Translate(), nil
	case "trash":
		return Trash(), nil
	case "treadmill":
		return Treadmill(), nil
	case "tree":
		return Tree(), nil
	case "trekking":
		return Trekking(), nil
	case "trello":
		return Trello(), nil
	case "triangle":
		return Triangle(), nil
	case "triangle-flag":
		return TriangleFlag(), nil
	case "triangle-flag-circle":
		return TriangleFlagCircle(), nil
	case "triangle-flag-two-stripes":
		return TriangleFlagTwoStripes(), nil
	case "trophy":
		return Trophy(), nil
	case "truck":
		return Truck(), nil
	case "truck-length":
		return TruckLength(), nil
	case "tunnel":
		return Tunnel(), nil
	case "tv":
		return Tv(), nil
	case "tv-fix":
		return TvFix(), nil
	case "tv-issue":
		return TvIssue(), nil
	case "twitter":
		return Twitter(), nil
	case "twitter-verified-badge":
		return TwitterVerifiedBadge(), nil
	case "two-points-circle":
		return TwoPointsCircle(), nil
	case "two-seater-sofa":
		return TwoSeaterSofa(), nil
	case "2x2-cell":
		return TwoXTwoCell(), nil
	case "type":
		return Type(), nil
	case "umbrella-full":
		return UmbrellaFull(), nil
	case "underline":
		return Underline(), nil
	case "underline-square":
		return UnderlineSquare(), nil
	case "undo":
		return Undo(), nil
	case "undo-action":
		return UndoAction(), nil
	case "undo-circle":
		return UndoCircle(), nil
	case "union":
		return Union(), nil
	case "union-alt":
		return UnionAlt(), nil
	case "union-horiz-alt":
		return UnionHorizAlt(), nil
	case "unity":
		return Unity(), nil
	case "unity-5":
		return UnityFive(), nil
	case "up-round-arrow":
		return UpRoundArrow(), nil
	case "upload":
		return Upload(), nil
	case "upload-data-window":
		return UploadDataWindow(), nil
	case "upload-square":
		return UploadSquare(), nil
	case "usb":
		return Usb(), nil
	case "user":
		return User(), nil
	case "user-bag":
		return UserBag(), nil
	case "user-cart":
		return UserCart(), nil
	case "user-circle":
		return UserCircle(), nil
	case "user-crown":
		return UserCrown(), nil
	case "user-love":
		return UserLove(), nil
	case "user-scan":
		return UserScan(), nil
	case "user-square":
		return UserSquare(), nil
	case "user-star":
		return UserStar(), nil
	case "vegan":
		return Vegan(), nil
	case "vegan-circle":
		return VeganCircle(), nil
	case "vegan-square":
		return VeganSquare(), nil
	case "verified-badge":
		return VerifiedBadge(), nil
	case "verified-user":
		return VerifiedUser(), nil
	case "vertical-merge":
		return VerticalMerge(), nil
	case "vertical-split":
		return VerticalSplit(), nil
	case "vials":
		return Vials(), nil
	case "video-camera":
		return VideoCamera(), nil
	case "video-camera-off":
		return VideoCameraOff(), nil
	case "video-projector":
		return VideoProjector(), nil
	case "view-columns-3":
		return ViewColumnsThree(), nil
	case "view-columns-2":
		return ViewColumnsTwo(), nil
	case "view-grid":
		return ViewGrid(), nil
	case "view-structure-down":
		return ViewStructureDown(), nil
	case "view-structure-up":
		return ViewStructureUp(), nil
	case "voice":
		return Voice(), nil
	case "voice-circle":
		return VoiceCircle(), nil
	case "voice-error":
		return VoiceError(), nil
	case "voice-lock-circle":
		return VoiceLockCircle(), nil
	case "voice-ok":
		return VoiceOk(), nil
	case "voice-phone":
		return VoicePhone(), nil
	case "voice-scan":
		return VoiceScan(), nil
	case "voice-square":
		return VoiceSquare(), nil
	case "vr-symbol":
		return VrSymbol(), nil
	case "waist":
		return Waist(), nil
	case "walking":
		return Walking(), nil
	case "wallet":
		return Wallet(), nil
	case "warning-circle":
		return WarningCircle(), nil
	case "warning-hexagon":
		return WarningHexagon(), nil
	case "warning-square":
		return WarningSquare(), nil
	case "warning-triangle":
		return WarningTriangle(), nil
	case "warning-window":
		return WarningWindow(), nil
	case "wash":
		return Wash(), nil
	case "washing-machine":
		return WashingMachine(), nil
	case "watering-soil":
		return WateringSoil(), nil
	case "web-window":
		return WebWindow(), nil
	case "web-window-close":
		return WebWindowClose(), nil
	case "web-window-energy-consumption":
		return WebWindowEnergyConsumption(), nil
	case "webp-format":
		return WebpFormat(), nil
	case "weight":
		return Weight(), nil
	case "weight-alt":
		return WeightAlt(), nil
	case "white-flag":
		return WhiteFlag(), nil
	case "wifi":
		return Wifi(), nil
	case "wifi-error":
		return WifiError(), nil
	case "wifi-issue":
		return WifiIssue(), nil
	case "wifi-off":
		return WifiOff(), nil
	case "wifi-signal-none":
		return WifiSignalNone(), nil
	case "wifi-tag":
		return WifiTag(), nil
	case "wind":
		return Wind(), nil
	case "windows":
		return Windows(), nil
	case "women-t-shirt":
		return WomenTShirt(), nil
	case "wrap-text":
		return WrapText(), nil
	case "wrench":
		return Wrench(), nil
	case "wristwatch":
		return Wristwatch(), nil
	case "www":
		return Www(), nil
	case "x-coordinate":
		return XCoordinate(), nil
	case "xbox-a":
		return XboxA(), nil
	case "xbox-b":
		return XboxB(), nil
	case "xbox-x":
		return XboxX(), nil
	case "xbox-y":
		return XboxY(), nil
	case "xray-view":
		return XrayView(), nil
	case "y-coordinate":
		return YCoordinate(), nil
	case "yen":
		return Yen(), nil
	case "yen-square":
		return YenSquare(), nil
	case "yoga":
		return Yoga(), nil
	case "youtube":
		return Youtube(), nil
	case "z-coordinate":
		return ZCoordinate(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in iconoir icon set", name)
	}
}
