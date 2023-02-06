package line_md

import "github.com/gogoracer/racer/pkg/engine"

const (
	accountInnerSVG                                           = `<g fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2"><path d="M4 21V20C4 16.6863 6.68629 14 10 14H14C17.3137 14 20 16.6863 20 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path><path d="M12 11C9.79086 11 8 9.20914 8 7C8 4.79086 9.79086 3 12 3C14.2091 3 16 4.79086 16 7C16 9.20914 14.2091 11 12 11Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="28;0"/></path></g>`
	accountAddInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21V20C3 17.7909 4.79086 16 7 16H11C13.2091 16 15 17.7909 15 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 13C7.34315 13 6 11.6569 6 10C6 8.34315 7.34315 7 9 7C10.6569 7 12 8.34315 12 10C12 11.6569 10.6569 13 9 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M15 6H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M18 3V9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g>`
	accountAlertInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M5 21V20C5 17.7909 6.79086 16 9 16H13C15.2091 16 17 17.7909 17 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M11 13C9.34315 13 8 11.6569 8 10C8 8.34315 9.34315 7 11 7C12.6569 7 14 8.34315 14 10C14 11.6569 12.6569 13 11 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20 3V7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path></g><circle cx="20" cy="11" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle>`
	accountDeleteInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21V20C3 17.7909 4.79086 16 7 16H11C13.2091 16 15 17.7909 15 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 13C7.34315 13 6 11.6569 6 10C6 8.34315 7.34315 7 9 7C10.6569 7 12 8.34315 12 10C12 11.6569 10.6569 13 9 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M15 3L21 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 3L15 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="10;0"/></path></g>`
	accountRemoveInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21V20C3 17.7909 4.79086 16 7 16H11C13.2091 16 15 17.7909 15 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 13C7.34315 13 6 11.6569 6 10C6 8.34315 7.34315 7 9 7C10.6569 7 12 8.34315 12 10C12 11.6569 10.6569 13 9 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M15 6H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`
	accountSmallInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-dasharray="20" stroke-dashoffset="20" stroke-linecap="round" stroke-width="2"><path d="M6 19V18C6 15.7909 7.79086 14 10 14H14C16.2091 14 18 15.7909 18 18V19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path d="M12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5C13.6569 5 15 6.34315 15 8C15 9.65685 13.6569 11 12 11Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path></g>`
	alertInnerSVG                                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3L21 20H3L12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 10V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`
	alertCircleInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`
	alertCircleTwotoneInnerSVG                                = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`
	alertTwotoneInnerSVG                                      = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3L21 20H3L12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="6" stroke-dashoffset="6" d="M12 10V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`
	alignCenterInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="9" stroke-dashoffset="9" d="M12 5H19M12 5H5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="9;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M12 10H17M12 10H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="7;0"/></path><path stroke-dasharray="11" stroke-dashoffset="11" d="M12 15H21M12 15H3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="11;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M12 20H19M12 20H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="9;0"/></path></g>`
	alignJustifyInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width="2"><path d="M12 5H18M12 5H6"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="8;0"/></path><path d="M12 10H18M12 10H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="8;0"/></path><path d="M12 15H18M12 15H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></path><path d="M12 20H18M12 20H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`
	alignLeftInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4 5H17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M4 10H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="18;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4 20H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="15;0"/></path></g>`
	arrowAlignCenterInnerSVG                                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 3V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M22 12H15.5M2 12H8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M15 12L18 15M9 12L6 15M15 12L18 9M9 12L6 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`
	arrowAlignLeftInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 3V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M21 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 12L11 16M7 12L11 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`
	arrowCloseLeftInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 3V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M21 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 12L14 19M7 12L14 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`
	arrowLeftInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M21 12H3.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M3 12L10 19M3 12L10 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="12;0"/></path></g>`
	arrowLeftCircleInnerSVG                                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" stroke-dasharray="60" stroke-dashoffset="60"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></circle><path stroke-dasharray="12" stroke-dashoffset="12" d="M17 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 12L11 16M7 12L11 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path></g>`
	arrowLeftCircleTwotoneInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></circle><path stroke-dasharray="12" stroke-dashoffset="12" d="M17 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 12L11 16M7 12L11 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path></g>`
	arrowLongDiagonalInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M12 12L3.5 20.5M12 12L20.5 3.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H13M3 21V13M21 3V11M3 21H11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`
	arrowOpenLeftInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M21 3V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M17 12H3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M3 12L10 19M3 12L10 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`
	arrowSmallLeftInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M19 12H5.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M5 12L10 17M5 12L10 7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`
	arrowsDiagonalInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M10 14L3.5 20.5M14 10L20.5 3.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H15M3 21V15M21 3V9M3 21H9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`
	arrowsHorizontalInnerSVG                                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="12" stroke-dashoffset="12" d="M15 7H3.5M9 17H20.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M3 7L7 11M3 7L7 3M21 17L17 21M21 17L17 13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`
	arrowsHorizontalAltInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="12" stroke-dashoffset="12" d="M21 7H10.5M3 17H13.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 7L14 11M10 7L14 3M14 17L10 21M14 17L10 13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`
	beerInnerSVG                                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`
	beerAltFilledInnerSVG                                     = `<mask id="lineMdBeerAltFilled0"><g transform="translate(0 14)"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"/><path fill="#fff" d="M17 8L16 21H7L6 8z"/><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-14"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><rect width="11" height="14" x="6" y="6" fill="currentColor" mask="url(#lineMdBeerAltFilled0)"/>`
	beerAltFilledLoopInnerSVG                                 = `<mask id="lineMdBeerAltFilledLoop0"><g transform="translate(0 14)"><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9"><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/></path><path fill="#fff" d="M17 8L16 21H7L6 8z"/><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-14"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><rect width="11" height="14" x="6" y="6" fill="currentColor" mask="url(#lineMdBeerAltFilledLoop0)"/>`
	beerAltTwotoneInnerSVG                                    = `<mask id="lineMdBeerAltTwotone0"><g transform="translate(0 14)"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"/><path fill="#fff" fill-opacity=".3" d="M17 8L16 21H7L6 8z"/><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-14"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><rect width="11" height="14" x="6" y="6" fill="currentColor" mask="url(#lineMdBeerAltTwotone0)"/>`
	beerAltTwotoneLoopInnerSVG                                = `<mask id="lineMdBeerAltTwotoneLoop0"><g transform="translate(0 14)"><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9"><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/></path><path fill="#fff" fill-opacity=".3" d="M17 8L16 21H7L6 8z"/><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-14"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><rect width="11" height="14" x="6" y="6" fill="currentColor" mask="url(#lineMdBeerAltTwotoneLoop0)"/>`
	beerFilledInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g><path fill="currentColor" fill-opacity="0" d="M17 8L16 21H7L6 8z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path>`
	beerLoopInnerSVG                                          = `<mask id="lineMdBeerLoop0"><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9" opacity="0"><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/><animate fill="freeze" attributeName="opacity" begin="0.6s" dur="0.5s" values="0;1"/></path></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" d="M18 3L16 21H7L5 3z" mask="url(#lineMdBeerLoop0)"/>`
	beerTwotoneInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g><path fill="currentColor" fill-opacity="0" d="M17 8L16 21H7L6 8z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path>`
	beerTwotoneLoopInnerSVG                                   = `<mask id="lineMdBeerTwotoneLoop0"><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9" opacity="0"><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/><animate fill="freeze" attributeName="opacity" begin="0.6s" dur="0.5s" values="0;1"/></path></mask><path fill="currentColor" fill-opacity="0" d="M17 8L16 21H7L6 8z"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" d="M18 3L16 21H7L5 3z" mask="url(#lineMdBeerTwotoneLoop0)"/>`
	bellInnerSVG                                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`
	bellTwotoneInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`
	buyMeACoffeeInnerSVG                                      = `<mask id="lineMdBuyMeACoffee0"><path fill="#fff" d="M5 6C5 4 7 6 11.5 6C16 6 19 4 19 6L19 7C19 8.5 17 9 12.5 9C8 9 5 9 5 7L5 6Z"/></mask><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M7.5 10.5C7.5 10.5 8.33 17.43 8.5 19C8.67 20.57 10 21 11 21L13 21C14.5 21 15.875 19.86 16 19C16.125 18.14 17 7 17 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M16.5 6C16.5 3.5 14 3 12 3C10 3 9.1 3.43 8 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;24"/></path></g><rect width="16" height="5" x="20" y="4" fill="currentColor" mask="url(#lineMdBuyMeACoffee0)"><animate fill="freeze" attributeName="x" begin="0.4s" dur="0.4s" values="20;4"/></rect>`
	buyMeACoffeeFilledInnerSVG                                = `<mask id="lineMdBuyMeACoffeeFilled0"><path fill="#fff" d="M5 6C5 4 7 6 11.5 6C16 6 19 4 19 6L19 7C19 8.5 17 9 12.5 9C8 9 5 9 5 7L5 6Z"/></mask><mask id="lineMdBuyMeACoffeeFilled1"><path fill="#fff" d="M10.125 18.15C10.04 17.29 9.4 11.98 9.4 11.98C9.4 11.98 11.34 12.31 12.5 11.73C13.66 11.16 14.98 11 14.98 11C14.98 11 14.4 17.96 14.35 18.46C14.3 18.96 13.45 19.3 12.95 19.3L11.23 19.3C10.73 19.3 10.21 19 10.125 18.15Z"/></mask><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M7.5 10.5C7.5 10.5 8.33 17.43 8.5 19C8.67 20.57 10 21 11 21L13 21C14.5 21 15.875 19.86 16 19C16.125 18.14 17 7 17 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M16.5 6C16.5 3.5 14 3 12 3C10 3 9.1 3.43 8 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;24"/></path></g><rect width="16" height="5" x="20" y="4" fill="currentColor" mask="url(#lineMdBuyMeACoffeeFilled0)"><animate fill="freeze" attributeName="x" begin="0.4s" dur="0.4s" values="20;4"/></rect><rect width="8" height="10" x="8" y="20" fill="currentColor" mask="url(#lineMdBuyMeACoffeeFilled1)"><animate fill="freeze" attributeName="y" begin="1s" dur="0.4s" values="20;10"/></rect>`
	buyMeACoffeeTwotoneInnerSVG                               = `<mask id="lineMdBuyMeACoffeeTwotone0"><path fill="#fff" d="M5 6C5 4 7 6 11.5 6C16 6 19 4 19 6L19 7C19 8.5 17 9 12.5 9C8 9 5 9 5 7L5 6Z"/></mask><mask id="lineMdBuyMeACoffeeTwotone1"><path fill="#fff" d="M10.125 18.15C10.04 17.29 9.4 11.98 9.4 11.98C9.4 11.98 11.34 12.31 12.5 11.73C13.66 11.16 14.98 11 14.98 11C14.98 11 14.4 17.96 14.35 18.46C14.3 18.96 13.45 19.3 12.95 19.3L11.23 19.3C10.73 19.3 10.21 19 10.125 18.15Z"/></mask><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M7.5 10.5C7.5 10.5 8.33 17.43 8.5 19C8.67 20.57 10 21 11 21L13 21C14.5 21 15.875 19.86 16 19C16.125 18.14 17 7 17 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M16.5 6C16.5 3.5 14 3 12 3C10 3 9.1 3.43 8 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;24"/></path></g><rect width="16" height="5" x="20" y="4" fill="currentColor" mask="url(#lineMdBuyMeACoffeeTwotone0)"><animate fill="freeze" attributeName="x" begin="0.4s" dur="0.4s" values="20;4"/></rect><rect width="8" height="10" x="8" y="20" fill="currentColor" fill-opacity=".3" mask="url(#lineMdBuyMeACoffeeTwotone1)"><animate fill="freeze" attributeName="y" begin="1s" dur="0.4s" values="20;10"/></rect>`
	calendarInnerSVG                                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="4" stroke-dasharray="64" stroke-dashoffset="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/></rect><path stroke-dasharray="6" stroke-dashoffset="6" d="M7 4V2M17 4V2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 11H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M7 15H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="9;0"/></path></g><rect width="14" height="0" x="5" y="5" fill="currentColor"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;3"/></rect>`
	calendarOutInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="4" stroke-dasharray="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="0;64"/></rect><path stroke-dasharray="6" d="M7 4V2M17 4V2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="0;6"/></path><path stroke-dasharray="12" d="M7 11H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="0;12"/></path><path stroke-dasharray="9" d="M7 15H14"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="0;9"/></path></g><rect width="14" height="3" x="5" y="5" fill="currentColor"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="3;0"/></rect>`
	cancelInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M5.63604 5.63603C9.15076 2.12131 14.8492 2.12131 18.364 5.63603C21.8787 9.15075 21.8787 14.8492 18.364 18.364C14.8492 21.8787 9.15076 21.8787 5.63604 18.364C2.12132 14.8492 2.12132 9.15075 5.63604 5.63603Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M6 6L18 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="18;0"/></path></g>`
	cancelTwotoneInnerSVG                                     = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M5.63604 5.63603C9.15076 2.12131 14.8492 2.12131 18.364 5.63603C21.8787 9.15075 21.8787 14.8492 18.364 18.364C14.8492 21.8787 9.15076 21.8787 5.63604 18.364C2.12132 14.8492 2.12132 9.15075 5.63604 5.63603Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="18" stroke-dashoffset="18" d="M6 6L18 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="18;0"/></path></g>`
	checkListThreeInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-dasharray="10" stroke-dashoffset="10" stroke-width="2"><path d="M3 5L5 7L9 3"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path d="M3 12L5 14L9 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path d="M3 19L5 21L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g stroke-dasharray="22" stroke-dashoffset="22"><rect width="9" height="3" x="11.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="22;0"/></rect><rect width="9" height="3" x="11.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="22;0"/></rect><rect width="9" height="3" x="11.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.5s" values="22;0"/></rect></g></g>`
	checkListThreeFilledInnerSVG                              = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g fill="none" stroke-dasharray="10" stroke-dashoffset="10" stroke-width="2"><path d="M3 5L5 7L9 3"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path d="M3 12L5 14L9 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path d="M3 19L5 21L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0" stroke-dasharray="22" stroke-dashoffset="22"><rect width="9" height="3" x="11.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.5s" values="0;1"/></rect><rect width="9" height="3" x="11.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.9s" dur="0.5s" values="0;1"/></rect><rect width="9" height="3" x="11.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.5s" values="0;1"/></rect></g></g>`
	checkListThreeTwotoneInnerSVG                             = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g fill="none" stroke-dasharray="10" stroke-dashoffset="10" stroke-width="2"><path d="M3 5L5 7L9 3"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path d="M3 12L5 14L9 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path d="M3 19L5 21L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0" stroke-dasharray="22" stroke-dashoffset="22"><rect width="9" height="3" x="11.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.15s" values="0;0.3"/></rect><rect width="9" height="3" x="11.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.9s" dur="0.15s" values="0;0.3"/></rect><rect width="9" height="3" x="11.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.15s" values="0;0.3"/></rect></g></g>`
	chevronDoubleLeftInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12L12 5M5 12L12 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path d="M11 12L18 5M11 12L18 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="10;0"/></path></g>`
	chevronLeftInnerSVG                                       = `<path fill="none" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round" stroke-width="2" d="M8 12L15 5M8 12L15 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path>`
	chevronLeftCircleInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M10 12L13 9M10 12L13 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`
	chevronLeftCircleTwotoneInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M10 12L13 9M10 12L13 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`
	chevronSmallDoubleLeftInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12L17 7M12 12L17 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="8;0"/></path><path d="M6 12L11 7M6 12L11 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="8;0"/></path></g>`
	chevronSmallLeftInnerSVG                                  = `<path stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width="2" d="M9 12L14 7M9 12L14 17" fill="currentColor"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="8;0"/></path>`
	chevronSmallTripleLeftInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 12L19 7M14 12L19 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="8;0"/></path><path d="M9 12L14 7M9 12L14 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="8;0"/></path><path d="M4 12L9 7M4 12L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="8;0"/></path></g>`
	chevronTripleLeftInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12L9 5M2 12L9 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path d="M8 12L15 5M8 12L15 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="10;0"/></path><path d="M14 12L21 5M14 12L21 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="10;0"/></path></g>`
	circleInnerSVG                                            = `<path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path>`
	circleToConfirmCircleTransitionInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`
	circleToConfirmCircleTwotoneTransitionInnerSVG            = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.15s" values="0;0.3"/></circle><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`
	circleTwotoneInnerSVG                                     = `<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path>`
	circleTwotoneToConfirmCircleTwotoneTransitionInnerSVG     = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity=".3"/><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`
	clipboardInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="66" stroke-dashoffset="66" stroke-width="2" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardArrowInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M12 3H19V11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M19 17V21H5V3H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="44;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 14H12.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 14L15 17M12 14L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardArrowTwotoneInnerSVG                             = `<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V10H20V18H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M12 3H19V11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M19 17V21H5V3H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="44;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 14H12.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 14L15 17M12 14L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardCheckInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardCheckToClipboardTransitionInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;10"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`
	clipboardCheckTwotoneInnerSVG                             = `<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardCheckTwotoneToClipboardTwotoneTransitionInnerSVG = `<path fill="currentColor" fill-opacity=".3" d="M6 4H10V6H14V4H18V20H6V4Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;10"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`
	clipboardListInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 10H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardListTwotoneInnerSVG                              = `<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 10H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardMinusInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 13H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardMinusTwotoneInnerSVG                             = `<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 13H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardPlusInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 13H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 10V16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardPlusTwotoneInnerSVG                              = `<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 13H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 10V16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardToClipboardCheckTransitionInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`
	clipboardTwotoneInnerSVG                                  = `<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="66" stroke-dashoffset="66" stroke-width="2" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`
	clipboardTwotoneToClipboardTwotoneCheckTransitionInnerSVG = `<path fill="currentColor" fill-opacity=".3" d="M6 4H10V6H14V4H18V20H6V4Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`
	closeInnerSVG                                             = `<path fill="none" stroke="currentColor" stroke-dasharray="12" stroke-dashoffset="12" stroke-linecap="round" stroke-width="2" d="M12 12L19 19M12 12L5 5M12 12L5 19M12 12L19 5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="12;0"/></path>`
	closeCircleInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 12L16 16M12 12L8 8M12 12L8 16M12 12L16 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`
	closeCircleTwotoneInnerSVG                                = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="8" stroke-dashoffset="8" d="M12 12L16 16M12 12L8 8M12 12L8 16M12 12L16 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`
	closeToMenuAltTransitionInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 5L19 19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 5L19 19;M5 5L19 5"/></path><path d="M12 12H12" opacity="0"><animate fill="freeze" attributeName="d" begin="0.2s" dur="0.4s" values="M12 12H12;M5 12H19"/><set attributeName="opacity" begin="0.2s" to="1"/></path><path d="M5 19L19 5"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 19L19 5;M5 19L19 19"/></path></g>`
	closeToMenuTransitionInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 5L12 12L19 5"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 5L12 12L19 5;M5 5L12 5L19 5"/></path><path d="M12 12H12"><animate fill="freeze" attributeName="d" dur="0.4s" values="M12 12H12;M5 12H19"/></path><path d="M5 19L12 12L19 19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 19L12 12L19 19;M5 19L12 19L19 19"/></path></g>`
	cloudInnerSVG                                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 19C12 19 9.5 19 7 19C4.5 19 3 17 3 15C3 13 4.5 11 7 11C8 11 8.5 11.5 8.5 11.5M12 19H17C19.5 19 21 17 21 15C21 13 19.5 11 17 11C16 11 15.5 11.5 15.5 11.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 11C17 11 17 10.5 17 10C17 7.5 15 5 12 5M7 11V10C7 7.5 9 5 12 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path></g>`
	cloudBracesLoopInnerSVG                                   = `<mask id="lineMdCloudBracesLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;13;12"/></circle><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="10;11;10;9;10"/></rect></g><path d="M18.5 12H18a1 1 0 0 1-1-1v-1a2 2 0 0 0-2-2h-1.5v2H15v1a2 2 0 0 0 2 2 2 2 0 0 0-2 2v1h-1.5v2H15a2 2 0 0 0 2-2v-1a1 1 0 0 1 1-1h.5v-2Z"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h-1h2z" repeatCount="indefinite"/></path><path d="M5.5 12v2H6a1 1 0 0 1 1 1v1a2 2 0 0 0 2 2h1.5v-2H9v-1a2 2 0 0 0-2-2 2 2 0 0 0 2-2v-1h1.5V8H9a2 2 0 0 0-2 2v1a1 1 0 0 1-1 1h-.5Z"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h1h-2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudBracesLoop0)"/>`
	cloudDownInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="3" stroke-dashoffset="3" d="M12 22L14 20M12 22L10 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="3;0"/></path></g>`
	cloudDownTwotoneInnerSVG                                  = `<path fill="currentColor" fill-opacity="0" d="M9 6L12 4L15 6L21 14L17 19H7L3 14L9 6Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.3s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="3" stroke-dashoffset="3" d="M12 22L14 20M12 22L10 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="3;0"/></path></g>`
	cloudDownloadLoopInnerSVG                                 = `<mask id="lineMdCloudDownloadLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g><rect width="4" height="5" x="10" y="9"/><path d="M12 18L17 13H7L12 18Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v1v-2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudDownloadLoop0)"/>`
	cloudDownloadOutlineLoopInnerSVG                          = `<mask id="lineMdCloudDownloadOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"/><rect width="8" height="8" x="8" y="10"/><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><g fill="#fff"><rect width="3" height="4" x="10.5" y="10"/><path d="M12 17L16 13H8L12 17Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v1v-2z" repeatCount="indefinite"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudDownloadOutlineLoop0)"/>`
	cloudFilledInnerSVG                                       = `<path fill="currentColor" fill-opacity="0" d="M9 7L12 5L15 7L21 15L18 19H6L3 15L9 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.5s" values="0;1"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 19C12 19 9.5 19 7 19C4.5 19 3 17 3 15C3 13 4.5 11 7 11C8 11 8.5 11.5 8.5 11.5M12 19H17C19.5 19 21 17 21 15C21 13 19.5 11 17 11C16 11 15.5 11.5 15.5 11.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 11C17 11 17 10.5 17 10C17 7.5 15 5 12 5M7 11V10C7 7.5 9 5 12 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path></g>`
	cloudLoopInnerSVG                                         = `<mask id="lineMdCloudLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudLoop0)"/>`
	cloudOffOutlineLoopInnerSVG                               = `<mask id="lineMdCloudOffOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="8" x="8" y="10"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><g fill="none" stroke-linecap="round" stroke-width="2"><path stroke="#000" d="M1 11h24" transform="rotate(45 13 12)"/><path stroke="#fff" d="M1 13h22" transform="rotate(45 13 12)"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M1 13h22;M3 13h22;M1 13h22"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudOffOutlineLoop0)"/>`
	cloudOutlineLoopInnerSVG                                  = `<mask id="lineMdCloudOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="8" x="8" y="10"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="8;7;8;9;8"/></rect></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudOutlineLoop0)"/>`
	cloudPrintLoopInnerSVG                                    = `<mask id="lineMdCloudPrintLoop0"><g fill="#fff"><circle cx="12" cy="8" r="6"/><rect width="15" height="12" x="1" y="6" rx="6"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="8" rx="5"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g><rect width="12" height="11" x="6" y="11" fill="#fff"/><rect width="8" height="7" x="8" y="13"/><path fill="#fff" d="M9 12h6v1H9zM9 14h6v1H9zM9 16h6v1H9zM9 18h6v1H9z"><animateMotion dur="1.5s" path="M0 0v2" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudPrintLoop0)"/>`
	cloudPrintOutlineLoopInnerSVG                             = `<mask id="lineMdCloudPrintOutlineLoop0"><g fill="#fff"><circle cx="12" cy="8" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="17" height="12" x="1" y="6" rx="6"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="8" rx="5"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="8" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="8" x="8" y="8"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="11" height="8" x="3" y="8" rx="4"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="10" rx="3"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="12" height="11" x="6" y="11" fill="#fff"/><rect width="8" height="7" x="8" y="13"/><path fill="#fff" d="M9 12h6v1H9zM9 14h6v1H9zM9 16h6v1H9zM9 18h6v1H9z"><animateMotion dur="1.5s" path="M0 0v2" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudPrintOutlineLoop0)"/>`
	cloudTagsLoopInnerSVG                                     = `<mask id="lineMdCloudTagsLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;13;12"/></circle><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="10;11;10;9;10"/></rect></g><path d="M6 10H8V14H12V16H6V10Z" transform="rotate(45 9 13)"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h-1h2z" repeatCount="indefinite"/></path><path d="M12 10H14V14H18V16H12V10Z" transform="rotate(225 15 13)"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h1h-2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudTagsLoop0)"/>`
	cloudTwotoneInnerSVG                                      = `<path fill="currentColor" fill-opacity="0" d="M9 7L12 5L15 7L21 15L18 19H6L3 15L9 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 19C12 19 9.5 19 7 19C4.5 19 3 17 3 15C3 13 4.5 11 7 11C8 11 8.5 11.5 8.5 11.5M12 19H17C19.5 19 21 17 21 15C21 13 19.5 11 17 11C16 11 15.5 11.5 15.5 11.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 11C17 11 17 10.5 17 10C17 7.5 15 5 12 5M7 11V10C7 7.5 9 5 12 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path></g>`
	cloudUpInnerSVG                                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 20V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 13L14 15M12 13L10 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="4;0"/></path></g>`
	cloudUpTwotoneInnerSVG                                    = `<path fill="currentColor" fill-opacity="0" d="M9 6L12 4L15 6L21 14L17 19H7L3 14L9 6Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.3s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 20V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 13L14 15M12 13L10 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="4;0"/></path></g>`
	cloudUploadLoopInnerSVG                                   = `<mask id="lineMdCloudUploadLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g><rect width="4" height="5" x="10" y="12"/><path d="M12 8L17 13H7L12 8Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v-1v2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudUploadLoop0)"/>`
	cloudUploadOutlineLoopInnerSVG                            = `<mask id="lineMdCloudUploadOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"/><rect width="8" height="8" x="8" y="10"/><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><g fill="#fff"><rect width="3" height="4" x="10.5" y="12"/><path d="M12 9L16 13H8L12 9Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v-1v2z" repeatCount="indefinite"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudUploadOutlineLoop0)"/>`
	coffeeInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="48" stroke-dashoffset="48" d="M17 4v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V4z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="48;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 9H20C20.55 9 21 8.55 21 8V5C21 4.45 20.55 4 20 4H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M11 20h8M11 20h-8"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path></g>`
	coffeeArrowInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="40" stroke-dashoffset="40" d="M14 4V11C14 11.7956 13.6839 12.5587 13.1213 13.1213C12.5587 13.6839 11.7956 14 11 14H7C6.20435 14 5.44129 13.6839 4.87868 13.1213C4.31607 12.5587 4 11.7956 4 11V4z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="40;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M14 9H17C17.55 9 18 8.55 18 8V5C18 4.45 17.55 4 17 4H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 18H19.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.3s" values="18;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M19.5 18l-3 -3M19.5 18l-3 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g>`
	coffeeArrowFilledInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M14 4V11C14 11.7956 13.6839 12.5587 13.1213 13.1213C12.5587 13.6839 11.7956 14 11 14H7C6.20435 14 5.44129 13.6839 4.87868 13.1213C4.31607 12.5587 4 11.7956 4 11V4z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.5s" values="0;1"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M14 9H17C17.55 9 18 8.55 18 8V5C18 4.45 17.55 4 17 4H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 18H19.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.3s" values="18;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M19.5 18l-3 -3M19.5 18l-3 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g>`
	coffeeArrowTwotoneInnerSVG                                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M14 4V11C14 11.7956 13.6839 12.5587 13.1213 13.1213C12.5587 13.6839 11.7956 14 11 14H7C6.20435 14 5.44129 13.6839 4.87868 13.1213C4.31607 12.5587 4 11.7956 4 11V4z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M14 9H17C17.55 9 18 8.55 18 8V5C18 4.45 17.55 4 17 4H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 18H19.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.3s" values="18;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M19.5 18l-3 -3M19.5 18l-3 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g>`
	coffeeFilledInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="48" stroke-dashoffset="48" d="M17 4v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V4z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.5s" values="0;1"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 9H20C20.55 9 21 8.55 21 8V5C21 4.45 20.55 4 20 4H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M11 20h8M11 20h-8"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path></g>`
	coffeeHalfEmptyTwotoneLoopInnerSVG                        = `<path fill="currentColor" fill-opacity="0" d="M17 14V17a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V14z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="48" stroke-dashoffset="48" d="M17 9v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 14H20C20.55 14 21 13.55 21 13V10C21 9.45 20.55 9 20 9H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path></g><mask id="lineMdCoffeeHalfEmptyTwotoneLoop0"><path fill="none" stroke="#fff" stroke-width="2" d="M8 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M12 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M16 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4"><animateMotion calcMode="linear" dur="3s" path="M0 0v-8" repeatCount="indefinite"/></path></mask><rect width="24" height="0" y="7" fill="currentColor" mask="url(#lineMdCoffeeHalfEmptyTwotoneLoop0)"><animate fill="freeze" attributeName="y" begin="0.8s" dur="0.6s" values="7;2"/><animate fill="freeze" attributeName="height" begin="0.8s" dur="0.6s" values="0;5"/></rect>`
	coffeeLoopInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="48" stroke-dashoffset="48" d="M17 9v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 14H20C20.55 14 21 13.55 21 13V10C21 9.45 20.55 9 20 9H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path></g><mask id="lineMdCoffeeLoop0"><path fill="none" stroke="#fff" stroke-width="2" d="M8 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M12 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M16 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4"><animateMotion calcMode="linear" dur="3s" path="M0 0v-8" repeatCount="indefinite"/></path></mask><rect width="24" height="0" y="7" fill="currentColor" mask="url(#lineMdCoffeeLoop0)"><animate fill="freeze" attributeName="y" begin="0.8s" dur="0.6s" values="7;2"/><animate fill="freeze" attributeName="height" begin="0.8s" dur="0.6s" values="0;5"/></rect>`
	coffeeTwotoneInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="48" stroke-dashoffset="48" d="M17 4v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V4z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 9H20C20.55 9 21 8.55 21 8V5C21 4.45 20.55 4 20 4H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M11 20h8M11 20h-8"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path></g>`
	coffeeTwotoneLoopInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="48" stroke-dashoffset="48" d="M17 9v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 14H20C20.55 14 21 13.55 21 13V10C21 9.45 20.55 9 20 9H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path></g><mask id="lineMdCoffeeTwotoneLoop0"><path fill="none" stroke="#fff" stroke-width="2" d="M8 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M12 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M16 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4"><animateMotion calcMode="linear" dur="3s" path="M0 0v-8" repeatCount="indefinite"/></path></mask><rect width="24" height="0" y="7" fill="currentColor" mask="url(#lineMdCoffeeTwotoneLoop0)"><animate fill="freeze" attributeName="y" begin="0.8s" dur="0.6s" values="7;2"/><animate fill="freeze" attributeName="height" begin="0.8s" dur="0.6s" values="0;5"/></rect>`
	computerInnerSVG                                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 21H17M12 21H7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="6;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 21V17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="6;0"/></path><path stroke-dasharray="64" stroke-dashoffset="64" d="M12 17H3V5H21V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.6s" values="64;0"/></path></g>`
	computerTwotoneInnerSVG                                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 21H17M12 21H7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="6;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 21V17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="6;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M12 17H3V5H21V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path></g>`
	confirmInnerSVG                                           = `<path fill="none" stroke="currentColor" stroke-dasharray="24" stroke-dashoffset="24" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 11L11 17L21 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="24;0"/></path>`
	confirmCircleInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;0"/></path></g>`
	confirmCircleToCircleTransitionInnerSVG                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path stroke-dasharray="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="28;14"/></path></g>`
	confirmCircleTwotoneInnerSVG                              = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;0"/></path></g>`
	confirmCircleTwotoneToCircleTransitionInnerSVG            = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity=".3"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.15s" values="0.3;0"/></circle><path fill="none" stroke-dasharray="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="28;14"/></path></g>`
	confirmCircleTwotoneToCircleTwotoneTransitionInnerSVG     = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity=".3"/><path fill="none" stroke-dasharray="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="28;14"/></path></g>`
	constructionInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="44" stroke-dashoffset="44" d="M21 21H3V19C3 18 4 17 5 17H19C20 17 21 18 21 19V21Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="44;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M6 17L12 2M18 17L12 2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="18;0"/></path></g><path stroke-dasharray="8" stroke-dashoffset="8" d="M8 12L12.5 9.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M6 16L13.5 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9.5 17L14.5 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.3s" dur="0.2s" values="8;0"/></path></g>`
	constructionTwotoneInnerSVG                               = `<path fill="currentColor" fill-opacity="0" d="M21 21H3V19C3 18 4 17 5 17H19C20 17 21 18 21 19V21Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="44" stroke-dashoffset="44" d="M21 21H3V19C3 18 4 17 5 17H19C20 17 21 18 21 19V21Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="44;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M6 17L12 2M18 17L12 2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="18;0"/></path></g><path stroke-dasharray="8" stroke-dashoffset="8" d="M8 12L12.5 9.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M6 16L13.5 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9.5 17L14.5 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.3s" dur="0.2s" values="8;0"/></path></g>`
	discordInnerSVG                                           = `<g fill="currentColor" fill-opacity="0"><circle cx="9" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.4s" values="0;1"/></circle></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M15.5 17.5L16.5 19.5C16.5 19.5 20.671 18.172 22 16C22 15 22.53 7.853 19 5.5C17.5 4.5 15 4 15 4L14 6H12M8.52799 17.5L7.52799 19.5C7.52799 19.5 3.35699 18.172 2.02799 16C2.02799 15 1.49799 7.853 5.02799 5.5C6.52799 4.5 9.02799 4 9.02799 4L10.028 6H12.028"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;60"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5.5 16C10.5 18.5 13.5 18.5 18.5 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="16;0"/></path></g>`
	discordTwotoneInnerSVG                                    = `<g fill="currentColor" fill-opacity="0"><path d="M5 5L12 5.2L19 5L22 15L19 18.4H5L1.5 15L5 5Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.15s" values="0;0.3"/></path><circle cx="9" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.4s" values="0;1"/></circle></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M15.5 17.5L16.5 19.5C16.5 19.5 20.671 18.172 22 16C22 15 22.53 7.853 19 5.5C17.5 4.5 15 4 15 4L14 6H12M8.52799 17.5L7.52799 19.5C7.52799 19.5 3.35699 18.172 2.02799 16C2.02799 15 1.49799 7.853 5.02799 5.5C6.52799 4.5 9.02799 4 9.02799 4L10.028 6H12.028"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;60"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5.5 16C10.5 18.5 13.5 18.5 18.5 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="16;0"/></path></g>`
	documentInnerSVG                                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`
	documentAddInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2"><path d="M9 14H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path d="M12 11V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g></g>`
	documentAddTwotoneInnerSVG                                = `<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2"><path d="M9 14H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path d="M12 11V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g></g>`
	documentCodeInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2"><path d="M10 13L8 15L10 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path d="M14 13L16 15L14 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g></g>`
	documentCodeTwotoneInnerSVG                               = `<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2"><path d="M10 13L8 15L10 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path d="M14 13L16 15L14 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g></g>`
	documentListInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`
	documentListTwotoneInnerSVG                               = `<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`
	documentRemoveInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2" d="M9 14H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`
	documentRemoveTwotoneInnerSVG                             = `<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2" d="M9 14H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`
	documentReportInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-width="2"><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 17V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 17V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M15 17V12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g></g>`
	documentReportTwotoneInnerSVG                             = `<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-width="2"><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 17V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 17V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M15 17V12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g></g>`
	documentTwotoneInnerSVG                                   = `<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`
	doubleArrowHorizontalInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 12H3.5M12 12H20.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M3 12L7 16M21 12L17 16M3 12L7 8M21 12L17 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="6;0"/></path></g>`
	downloadLoopInnerSVG                                      = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="14;0"/></path><path fill="currentColor" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5;M12 4 h2 v3 h2.5 L12 11.5M12 4 h-2 v3 h-2.5 L12 11.5;M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"/></path></g>`
	downloadOffLoopInnerSVG                                   = `<mask id="lineMdDownloadOffLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="14;0"/></path><path fill="#fff" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5;M12 4 h2 v3 h2.5 L12 11.5M12 4 h-2 v3 h-2.5 L12 11.5;M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdDownloadOffLoop0)"/>`
	downloadOffOutlineInnerSVG                                = `<mask id="lineMdDownloadOffOutline0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdDownloadOffOutline0)"/>`
	downloadOffOutlineLoopInnerSVG                            = `<mask id="lineMdDownloadOffOutlineLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5;M12 4 h2 v3 h2.5 L12 11.5M12 4 h-2 v3 h-2.5 L12 11.5;M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdDownloadOffOutlineLoop0)"/>`
	downloadOutlineInnerSVG                                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path></g>`
	downloadOutlineLoopInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5;M12 4 h2 v3 h2.5 L12 11.5M12 4 h-2 v3 h-2.5 L12 11.5;M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"/></path></g>`
	downloadingLoopInnerSVG                                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="2 4" stroke-dashoffset="6" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21"><animate attributeName="stroke-dashoffset" dur="0.6s" repeatCount="indefinite" values="6;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.3s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 8v7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 15.5l3.5 -3.5M12 15.5l-3.5 -3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="6;0"/></path></g>`
	editInnerSVG                                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M7 17V13L17 3L21 7L11 17H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="44;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 6L18 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`
	editTwotoneInnerSVG                                       = `<path fill="currentColor" fill-opacity="0" d="M20 7L17 4L15 6L18 9L20 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M7 17V13L17 3L21 7L11 17H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="44;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 6L18 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`
	editTwotoneFullInnerSVG                                   = `<path fill="currentColor" fill-opacity="0" d="M20 7L17 4L8 13V16H11L20 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M7 17V13L17 3L21 7L11 17H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="44;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 6L18 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`
	emailInnerSVG                                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="14" x="3" y="5" stroke-dasharray="64" stroke-dashoffset="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></rect><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 6.5L12 12L21 6.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`
	emailOpenedInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M3 8.06083C3 7.71247 3.1813 7.38921 3.47855 7.20755L12 2L20.5214 7.20755C20.8187 7.38921 21 7.71247 21 8.06083V18C21 18.5523 20.5523 19 20 19H4C3.44772 19 3 18.5523 3 18V8.06083Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 8.5L12 14L21 8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`
	emailOpenedTwotoneInnerSVG                                = `<path fill="currentColor" fill-opacity="0" d="M12 13L4 8L12 3L20 8L12 13Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M3 8.06083C3 7.71247 3.1813 7.38921 3.47855 7.20755L12 2L20.5214 7.20755C20.8187 7.38921 21 7.71247 21 8.06083V18C21 18.5523 20.5523 19 20 19H4C3.44772 19 3 18.5523 3 18V8.06083Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 8.5L12 14L21 8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`
	emailOpenedTwotoneAltInnerSVG                             = `<path fill="currentColor" fill-opacity="0" d="M12 15L4 10V18H20V10L12 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M3 8.06083C3 7.71247 3.1813 7.38921 3.47855 7.20755L12 2L20.5214 7.20755C20.8187 7.38921 21 7.71247 21 8.06083V18C21 18.5523 20.5523 19 20 19H4C3.44772 19 3 18.5523 3 18V8.06083Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 8.5L12 14L21 8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`
	emailTwotoneInnerSVG                                      = `<path fill="currentColor" fill-opacity="0" d="M12 11L4 6H20L12 11Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="14" x="3" y="5" stroke-dasharray="64" stroke-dashoffset="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></rect><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 6.5L12 12L21 6.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`
	emailTwotoneAltInnerSVG                                   = `<path fill="currentColor" fill-opacity="0" d="M12 13L4 8V18H20V8L12 13Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="14" x="3" y="5" stroke-dasharray="64" stroke-dashoffset="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></rect><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 6.5L12 12L21 6.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`
	emojiAngryInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="5" stroke-dashoffset="5" stroke-width="1"><path d="M9.5 8.5L7.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="5;0"/></path><path d="M14.5 8.5L16.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="5;0"/></path></g></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiAngryTwotoneInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="5" stroke-dashoffset="5" stroke-width="1"><path d="M9.5 8.5L7.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="5;0"/></path><path d="M14.5 8.5L16.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="5;0"/></path></g></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiFrownInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiFrownOpenInnerSVG                                    = `<mask id="lineMdEmojiFrownOpen0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8.5 16C9 15 10 14 12 14C14 14 15 15 15.5 16M8.5 16C9.5 15.5 11 15.5 12 15.5C13 15.5 14.5 15.5 15.5 16"/></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><g fill="currentColor"><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiFrownOpen0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiFrownOpenTwotoneInnerSVG                             = `<mask id="lineMdEmojiFrownOpenTwotone0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8.5 16C9 15 10 14 12 14C14 14 15 15 15.5 16M8.5 16C9.5 15.5 11 15.5 12 15.5C13 15.5 14.5 15.5 15.5 16"/></mask><g fill="currentColor"><path fill-opacity="0" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiFrownOpenTwotone0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiFrownTwotoneInnerSVG                                 = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiGrinInnerSVG                                         = `<mask id="lineMdEmojiGrin0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14M8 14C9 14.5 10 15 12 15C14 15 15 14.5 16 14"/></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><g fill="currentColor"><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiGrin0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiGrinTwotoneInnerSVG                                  = `<mask id="lineMdEmojiGrinTwotone0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14M8 14C9 14.5 10 15 12 15C14 15 15 14.5 16 14"/></mask><g fill="currentColor"><path fill-opacity="0" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiGrinTwotone0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiNeutralInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M8 15H16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiNeutralTwotoneInnerSVG                               = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="10" stroke-dashoffset="10" d="M8 15H16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiSmileInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiSmileTwotoneInnerSVG                                 = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`
	emojiSmileWinkInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><path d="M15 8.5C15.8284 8.5 16.5 8.94772 16.5 9.5C16.5 10.0523 15.8284 10.5 15 10.5C14.1716 10.5 13.5 10.0523 13.5 9.5C13.5 8.94772 14.1716 8.5 15 8.5Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></path></g>`
	emojiSmileWinkTwotoneInnerSVG                             = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><path d="M15 8.5C15.8284 8.5 16.5 8.94772 16.5 9.5C16.5 10.0523 15.8284 10.5 15 10.5C14.1716 10.5 13.5 10.0523 13.5 9.5C13.5 8.94772 14.1716 8.5 15 8.5Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></path></g>`
	externalLinkInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="42" stroke-dashoffset="42" d="M11 5H5V19H19V13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="42;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M13 11L20 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H15M21 3V9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path></g>`
	externalLinkRoundedInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 5C8.13401 5 5 8.13401 5 12C5 15.866 8.13401 19 12 19C15.866 19 19 15.866 19 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="36;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M13 11L20 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H15M21 3V9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path></g>`
	facebookInnerSVG                                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-dasharray="24" stroke-dashoffset="24" d="M17 4L15 4C12.5 4 11 5.5 11 8V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="24;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M8 12H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="12;0"/></path></g>`
	githubInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 4C13.6683 4 14.6122 4.39991 15 4.5C15.5255 4.07463 16.9375 3 18.5 3C18.8438 4 18.7863 5.21921 18.5 6C19.25 7 19.5 8 19.5 9.5C19.5 11.6875 19.017 13.0822 18 14C16.983 14.9178 15.8887 15.3749 14.5 15.5C15.1506 16.038 15 17.3743 15 18C15 18.7256 15 21 15 21M12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98301 13.0822 6 14C7.01699 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18C9 18.7256 9 21 9 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 19C7.59375 19 6.15625 18.4375 5.3125 17.8125C4.46875 17.1875 4.21875 16.1562 3 15.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path></g>`
	githubLoopInnerSVG                                        = `<mask id="lineMdGithubLoop0" width="24" height="24" x="0" y="0"><g fill="#fff"><ellipse cx="9.5" cy="9" rx="1.5" ry="1"/><ellipse cx="14.5" cy="9" rx="1.5" ry="1"/></g></mask><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 4C13.6683 4 14.6122 4.39991 15 4.5C15.5255 4.07463 16.9375 3 18.5 3C18.8438 4 18.7863 5.21921 18.5 6C19.25 7 19.5 8 19.5 9.5C19.5 11.6875 19.017 13.0822 18 14C16.983 14.9178 15.8887 15.3749 14.5 15.5C15.1506 16.038 15 17.3743 15 18C15 18.7256 15 21 15 21M12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98301 13.0822 6 14C7.01699 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18C9 18.7256 9 21 9 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/><animate attributeName="d" dur="3s" repeatCount="indefinite" values="M9 19c-1.406 0-2.844-.563-3.688-1.188C4.47 17.188 4.22 16.157 3 15.5;M9 19c-1.406 0-3-.5-4-.5-.532 0-1 0-2-.5;M9 19c-1.406 0-2.844-.563-3.688-1.188C4.47 17.188 4.22 16.157 3 15.5"/></path></g><rect width="8" height="4" x="8" y="11" fill="currentColor" mask="url(#lineMdGithubLoop0)"><animate attributeName="y" dur="10s" keyTimes="0;0.45;0.46;0.54;0.55;1" repeatCount="indefinite" values="11;11;7;7;11;11"/></rect>`
	githubTwotoneInnerSVG                                     = `<path fill="currentColor" fill-opacity="0" d="M15 4.5C14.6122 4.39991 13.6683 4 12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98302 13.0822 6 14C7.01698 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18V22H15V18C15 17.3743 15.1506 16.038 14.5 15.5C15.8887 15.3749 16.983 14.9178 18 14C19.017 13.0822 19.5 11.6875 19.5 9.5C19.5 8 19.25 7 18.5 6C18.7863 5.21921 18.8438 4 18.5 3C16.9375 3 15.5255 4.07463 15 4.5Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 4C13.6683 4 14.6122 4.39991 15 4.5C15.5255 4.07463 16.9375 3 18.5 3C18.8438 4 18.7863 5.21921 18.5 6C19.25 7 19.5 8 19.5 9.5C19.5 11.6875 19.017 13.0822 18 14C16.983 14.9178 15.8887 15.3749 14.5 15.5C15.1506 16.038 15 17.3743 15 18C15 18.7256 15 21 15 21M12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98301 13.0822 6 14C7.01699 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18C9 18.7256 9 21 9 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 19C7.59375 19 6.15625 18.4375 5.3125 17.8125C4.46875 17.1875 4.21875 16.1562 3 15.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path></g>`
	gridThreeInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round"><g><circle cx="5" cy="5" r="1.5"/><circle cx="12" cy="5" r="1.5"/><circle cx="19" cy="5" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></g><g><circle cx="5" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="19" cy="12" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="10;0"/></g><g><circle cx="5" cy="19" r="1.5"/><circle cx="12" cy="19" r="1.5"/><circle cx="19" cy="19" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/></g></g>`
	gridThreeFilledInnerSVG                                   = `<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round"><g><circle cx="5" cy="5" r="1.5"/><circle cx="12" cy="5" r="1.5"/><circle cx="19" cy="5" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.5s" values="0;1"/></g><g><circle cx="5" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="19" cy="12" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.5s" values="0;1"/></g><g><circle cx="5" cy="19" r="1.5"/><circle cx="12" cy="19" r="1.5"/><circle cx="19" cy="19" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.5s" values="0;1"/></g></g>`
	gridThreeTwotoneInnerSVG                                  = `<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round"><g><circle cx="5" cy="5" r="1.5"/><circle cx="12" cy="5" r="1.5"/><circle cx="19" cy="5" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></g><g><circle cx="5" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="19" cy="12" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></g><g><circle cx="5" cy="19" r="1.5"/><circle cx="12" cy="19" r="1.5"/><circle cx="19" cy="19" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></g></g>`
	hashInnerSVG                                              = `<g fill="none" stroke="currentColor" stroke-dasharray="20" stroke-dashoffset="20" stroke-linecap="round" stroke-width="2"><path d="M4 9H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path d="M3 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="20;0"/></path><path d="M10 3L8 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="20;0"/></path><path d="M16 3L14 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="20;0"/></path></g>`
	hashSmallInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-dasharray="16" stroke-dashoffset="16" stroke-linecap="round" stroke-width="2"><path d="M6 9H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path d="M5 15H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="16;0"/></path><path d="M10 5L8 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="16;0"/></path><path d="M16 5L14 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="16;0"/></path></g>`
	heartInnerSVG                                             = `<path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`
	heartFilledInnerSVG                                       = `<path fill="currentColor" fill-opacity="0" d="M12 20L20.5 11V7L17 5.5L12 7L7 5.5L3.5 7V11L12 20Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`
	heartFilledHalfInnerSVG                                   = `<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`
	heartHalfInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 8V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="15;0"/></path></g>`
	heartHalfFilledInnerSVG                                   = `<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 8V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="15;0"/></path></g>`
	heartHalfTwotoneInnerSVG                                  = `<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 8V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="15;0"/></path></g>`
	heartTwotoneInnerSVG                                      = `<path fill="currentColor" fill-opacity="0" d="M12 20L20.5 11V7L17 5.5L12 7L7 5.5L3.5 7V11L12 20Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`
	heartTwotoneHalfInnerSVG                                  = `<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`
	heartTwotoneHalfFilledInnerSVG                            = `<g fill="currentColor" fill-opacity="0"><path d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path><path d="M12 20L20.5 11V7L17 5.5L12 7V20Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path></g><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`
	homeInnerSVG                                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5h15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5V8M19.5 21.5V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9.5 21.5V12.5H14.5V21.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" stroke-width="2" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="30;0"/></path></g>`
	homeMdInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9 21V13H15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`
	homeMdTwotoneInnerSVG                                     = `<path fill="currentColor" fill-opacity="0" d="M6 8L12 3L18 8V20H16V13L15 12H9L8 13V20H6V8Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9 21V13H15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`
	homeMdTwotoneAltInnerSVG                                  = `<rect width="4" height="6" x="10" y="14" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></rect><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9 21V13H15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`
	homeSimpleInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`
	homeSimpleFilledInnerSVG                                  = `<path fill="currentColor" fill-opacity="0" d="M6 8L12 3L18 8V20H6V8Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`
	homeSimpleTwotoneInnerSVG                                 = `<path fill="currentColor" fill-opacity="0" d="M6 8L12 3L18 8V20H6V8Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`
	homeTwotoneInnerSVG                                       = `<path fill="currentColor" fill-opacity="0" d="M5 8.5L12 3L19 8.5V21H15V13L14 12H10L9 13V21H5V8.5Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5h15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5V8M19.5 21.5V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9.5 21.5V12.5H14.5V21.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" stroke-width="2" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="30;0"/></path></g>`
	homeTwotoneAltInnerSVG                                    = `<rect width="4" height="8" x="10" y="13" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></rect><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5h15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5V8M19.5 21.5V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9.5 21.5V12.5H14.5V21.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" stroke-width="2" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="30;0"/></path></g>`
	iconifyOneInnerSVG                                        = `<path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><g fill="currentColor" fill-opacity="0"><path fill-rule="evenodd" d="M12 18C15.125 18 17.3257 15.122 17 14.5C16.6728 13.875 15.5 16 12 16C8.5 16 7.3125 13.875 7 14.5C6.6875 15.125 8.875 18 12 18Z" clip-rule="evenodd"><animate fill="freeze" attributeName="fill-opacity" begin="1.0s" dur="0.2s" values="0;1"/></path><path d="M9.5 9C9.5 8.48223 9.01777 8 8.5 8C7.98223 8 7.5 8.48223 7.5 9V10.4375C7.5 10.9553 7.98223 11.5 8.5 11.5C9.01777 11.5 9.5 11.0178 9.5 10.5V9Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></path><path d="M16.5 9C16.5 8.48223 16.0178 8 15.5 8C14.9822 8 14.5 8.48223 14.5 9V10.4375C14.5 10.9553 14.9822 11.5 15.5 11.5C16.0178 11.5 16.5 11.0178 16.5 10.5V9Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></path></g>`
	iconifyTwoInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M4 7V21" class="il-md-length-15 il-md-duration-2 il-md-delay-0"/><path d="M4 3V5" class="il-md-length-15 il-md-duration-2 il-md-delay-0"/><path stroke-linecap="round" d="M18 4.25204C17.3608 4.08751 16.6906 4 16 4C11.5817 4 8 7.58172 8 12C8 16.4183 11.5817 20 16 20C16.6906 20 17.3608 19.9125 18 19.748" class="il-md-length-40 il-md-duration-3 il-md-delay-2"/><path stroke-linecap="round" d="M16 8C13.7909 8 12 9.79086 12 12C12 14.2091 13.7909 16 16 16C18.2091 16 20 14.2091 20 12C20 9.79086 18.2091 8 16 8Z" class="il-md-length-40 il-md-duration-5 il-md-delay-5"/></g>`
	imageInnerSVG                                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="66" stroke-dashoffset="66" stroke-width="2" d="M3 14V5H21V19H3V14"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M3 16L7 13L10 15L16 10L21 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="26;0"/></path></g><circle cx="7.5" cy="9.5" r="1.5" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></circle>`
	imageTwotoneInnerSVG                                      = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="66" stroke-dashoffset="66" stroke-width="2" d="M3 14V5H21V19H3V14"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="52" stroke-dashoffset="52" d="M3 16L7 13L10 15L16 10L21 14V19H3Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.8s" values="52;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path></g><circle cx="7.5" cy="9.5" r="1.5" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></circle>`
	instagramInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H8C5.23858 3 3 5.23858 3 8V16C3 18.7614 5.23858 21 8 21H16C18.7614 21 21 18.7614 21 16V8C21 5.23858 18.7614 3 16 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;132"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M12 8C14.20914 8 16 9.79086 16 12C16 14.20914 14.20914 16 12 16C9.79086 16 8 14.2091 8 12C8 9.79086 9.79086 8 12 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="26;0"/></path></g><circle cx="17" cy="7" r="1.5" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></circle>`
	laptopInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="50" stroke-dashoffset="50" d="M12 17H5V7H19V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="50;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 19H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.3s" values="20;0"/></path></g>`
	laptopTwotoneInnerSVG                                     = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" d="M12 17H5V7H19V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.0s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M3 19H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.3s" values="20;0"/></path></g>`
	lightDarkInnerSVG                                         = `<defs><mask id="lineMdLightDark0"><circle cx="7.5" cy="7.5" r="5.5" fill="#fff"/><circle cx="7.5" cy="7.5" r="5.5"><animate fill="freeze" attributeName="cx" dur="0.4s" values="7.5;11"/><animate fill="freeze" attributeName="r" dur="0.4s" values="5.5;6.5"/></circle></mask><mask id="lineMdLightDark1"><g fill="#fff"><circle cx="12" cy="9" r="5.5"><animate fill="freeze" attributeName="cy" begin="1s" dur="0.5s" values="9;15"/></circle><g fill-opacity="0"><use href="#lineMdLightDark2" transform="rotate(-75 12 15)"/><use href="#lineMdLightDark2" transform="rotate(-25 12 15)"/><use href="#lineMdLightDark2" transform="rotate(25 12 15)"/><use href="#lineMdLightDark2" transform="rotate(75 12 15)"/><set attributeName="fill-opacity" begin="1.5s" to="1"/></g></g><path d="M0 10h26v5h-26z"/><path fill="none" stroke="#fff" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" d="M1 12h22"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;52"/></path></mask><symbol id="lineMdLightDark2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.4s" values="M11 18h2L12 20z;M10.5 21.5h3L12 24z"/></path></symbol></defs><g fill="currentColor"><rect width="13" height="13" x="1" y="1" mask="url(#lineMdLightDark0)"/><path d="M-2 11h28v13h-28z" mask="url(#lineMdLightDark1)" transform="rotate(-45 12 12)"/></g>`
	lightDarkLoopInnerSVG                                     = `<defs><mask id="lineMdLightDarkLoop0"><circle cx="7.5" cy="7.5" r="5.5" fill="#fff"/><circle cx="7.5" cy="7.5" r="5.5"><animate fill="freeze" attributeName="cx" dur="0.4s" values="7.5;11"/><animate fill="freeze" attributeName="r" dur="0.4s" values="5.5;6.5"/></circle></mask><mask id="lineMdLightDarkLoop1"><g fill="#fff"><circle cx="12" cy="9" r="5.5"><animate fill="freeze" attributeName="cy" begin="1s" dur="0.5s" values="9;15"/></circle><g><g fill-opacity="0"><use href="#lineMdLightDarkLoop2" transform="rotate(-125 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(-75 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(-25 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(25 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(75 12 15)"/><set attributeName="fill-opacity" begin="1.5s" to="1"/></g><animateTransform attributeName="transform" dur="5s" repeatCount="indefinite" type="rotate" values="0 12 15;50 12 15"/></g></g><path d="M0 10h26v5h-26z"/><path fill="none" stroke="#fff" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" d="M1 12h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 12h22;M2 12h22;M0 12h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;52"/></path></mask><symbol id="lineMdLightDarkLoop2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.4s" values="M11 18h2L12 20z;M10.5 21.5h3L12 24z"/></path></symbol></defs><g fill="currentColor"><rect width="13" height="13" x="1" y="1" mask="url(#lineMdLightDarkLoop0)"/><path d="M-2 11h28v13h-28z" mask="url(#lineMdLightDarkLoop1)" transform="rotate(-45 12 12)"/></g>`
	lightbulbInnerSVG                                         = `<path fill="none" stroke="currentColor" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/></path><rect width="6" height="0" x="9" y="20" fill="currentColor" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect>`
	lightbulbFilledInnerSVG                                   = `<g fill="currentColor"><path fill-opacity="0" stroke="currentColor" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g>`
	lightbulbOffInnerSVG                                      = `<mask id="lineMdLightbulbOff0"><path fill="none" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/></path><rect width="6" height="0" x="9" y="20" fill="#fff" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOff0)"/>`
	lightbulbOffFilledInnerSVG                                = `<mask id="lineMdLightbulbOffFilled0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffFilled0)"/>`
	lightbulbOffFilledLoopInnerSVG                            = `<mask id="lineMdLightbulbOffFilledLoop0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffFilledLoop0)"/>`
	lightbulbOffLoopInnerSVG                                  = `<mask id="lineMdLightbulbOffLoop0"><path fill="none" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/></path><rect width="6" height="0" x="9" y="20" fill="#fff" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffLoop0)"/>`
	lightbulbOffTwotoneInnerSVG                               = `<mask id="lineMdLightbulbOffTwotone0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffTwotone0)"/>`
	lightbulbOffTwotoneLoopInnerSVG                           = `<mask id="lineMdLightbulbOffTwotoneLoop0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffTwotoneLoop0)"/>`
	lightbulbTwotoneInnerSVG                                  = `<g fill="currentColor"><path fill-opacity="0" stroke="currentColor" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g>`
	linkedinInnerSVG                                          = `<circle cx="4" cy="4" r="2" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="0;1"/></circle><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-dasharray="12" stroke-dashoffset="12" d="M4 10V20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M10 10V20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M10 15C10 12.2386 12.2386 10 15 10C17.7614 10 20 12.2386 20 15V20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="24;0"/></path></g>`
	listInnerSVG                                              = `<g fill="none" stroke="currentColor" stroke-dasharray="14" stroke-dashoffset="14" stroke-linecap="round" stroke-width="2"><path d="M8 5H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.2s" values="14;0"/></path><path d="M8 10H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="14;0"/></path><path d="M8 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><path d="M8 20H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><circle cx="4" cy="5" r="1"><animate fill="freeze" attributeName="fill-opacity" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="10" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="15" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="20" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></g>`
	listIndentedInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 5H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.2s" values="14;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M10 10H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 20H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><g fill="currentColor" fill-opacity="0"><circle cx="4" cy="5" r="1"><animate fill="freeze" attributeName="fill-opacity" dur="0.2s" values="0;1"/></circle><circle cx="6" cy="10" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.2s" values="0;1"/></circle><circle cx="8" cy="15" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></circle><circle cx="10" cy="20" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></g>`
	listThreeInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-linecap="round"><g stroke-dasharray="10" stroke-dashoffset="10"><circle cx="5" cy="5" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></circle><circle cx="5" cy="12" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></circle><circle cx="5" cy="19" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="10;0"/></circle></g><g stroke-dasharray="28" stroke-dashoffset="28"><rect width="11" height="3" x="9.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.5s" values="28;0"/></rect><rect width="11" height="3" x="9.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.5s" values="28;0"/></rect><rect width="11" height="3" x="9.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.5s" values="28;0"/></rect></g></g>`
	listThreeFilledInnerSVG                                   = `<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round"><g stroke-dasharray="10" stroke-dashoffset="10"><circle cx="5" cy="5" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.0s" dur="0.5s" values="0;1"/></circle><circle cx="5" cy="12" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.2s" dur="0.5s" values="0;1"/></circle><circle cx="5" cy="19" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.4s" dur="0.5s" values="0;1"/></circle></g><g stroke-dasharray="28" stroke-dashoffset="28"><rect width="11" height="3" x="9.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.0s" dur="0.5s" values="0;1"/></rect><rect width="11" height="3" x="9.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.2s" dur="0.5s" values="0;1"/></rect><rect width="11" height="3" x="9.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.4s" dur="0.5s" values="0;1"/></rect></g></g>`
	listThreeTwotoneInnerSVG                                  = `<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round"><g stroke-dasharray="10" stroke-dashoffset="10"><circle cx="5" cy="5" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.0s" dur="0.15s" values="0;0.3"/></circle><circle cx="5" cy="12" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.2s" dur="0.15s" values="0;0.3"/></circle><circle cx="5" cy="19" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.4s" dur="0.15s" values="0;0.3"/></circle></g><g stroke-dasharray="28" stroke-dashoffset="28"><rect width="11" height="3" x="9.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.0s" dur="0.15s" values="0;0.3"/></rect><rect width="11" height="3" x="9.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.2s" dur="0.15s" values="0;0.3"/></rect><rect width="11" height="3" x="9.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.4s" dur="0.15s" values="0;0.3"/></rect></g></g>`
	loadingAltLoopInnerSVG                                    = `<circle cx="12" cy="3.5" r="1.5" fill="currentColor" opacity="0"><animateTransform attributeName="transform" calcMode="discrete" dur="2.4s" repeatCount="indefinite" type="rotate" values="0 12 12;90 12 12;180 12 12;270 12 12"/><animate attributeName="opacity" dur="0.6s" keyTimes="0;0.5;1" repeatCount="indefinite" values="1;1;0"/></circle><circle cx="12" cy="3.5" r="1.5" fill="currentColor" opacity="0"><animateTransform attributeName="transform" begin="0.2s" calcMode="discrete" dur="2.4s" repeatCount="indefinite" type="rotate" values="30 12 12;120 12 12;210 12 12;300 12 12"/><animate attributeName="opacity" begin="0.2s" dur="0.6s" keyTimes="0;0.5;1" repeatCount="indefinite" values="1;1;0"/></circle><circle cx="12" cy="3.5" r="1.5" fill="currentColor" opacity="0"><animateTransform attributeName="transform" begin="0.4s" calcMode="discrete" dur="2.4s" repeatCount="indefinite" type="rotate" values="60 12 12;150 12 12;240 12 12;330 12 12"/><animate attributeName="opacity" begin="0.4s" dur="0.6s" keyTimes="0;0.5;1" repeatCount="indefinite" values="1;1;0"/></circle>`
	loadingLoopInnerSVG                                       = `<path fill="none" stroke="currentColor" stroke-dasharray="15" stroke-dashoffset="15" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="15;0"/><animateTransform attributeName="transform" dur="1.5s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></path>`
	loadingTwotoneLoopInnerSVG                                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" stroke-opacity=".3" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="1.3s" values="60;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 3C16.9706 3 21 7.02944 21 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="15;0"/><animateTransform attributeName="transform" dur="1.5s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></path></g>`
	markerInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="44" stroke-dashoffset="44" stroke-width="2" d="M6 14L17 3L21 7L10 18L6 14Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="44;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 17L6.5 19.5H2.5L7 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="20;0"/></path></g>`
	markerFilledInnerSVG                                      = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="44" stroke-dashoffset="44" stroke-width="2" d="M6 14L17 3L21 7L10 18L6 14Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="44;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="20" stroke-dashoffset="20" d="M9 17L6.5 19.5H2.5L7 15Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.5s" values="0;1"/></path></g>`
	markerTwotoneInnerSVG                                     = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="44" stroke-dashoffset="44" stroke-width="2" d="M6 14L17 3L21 7L10 18L6 14Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="44;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="20" stroke-dashoffset="20" d="M9 17L6.5 19.5H2.5L7 15Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path></g>`
	mastodonInnerSVG                                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="80" stroke-dashoffset="80" d="M15.5 21.5C6 23.5 6.5 16.5 7.5 16.5C11 16.5 21 19 21 12.5V8.5C21 4.5 18.5 3 14 3H10C5.5 3 3 4 3 10V13C3 19 5 24 15.5 21.5Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="80;160"/></path><path stroke-dasharray="32" stroke-dashoffset="32" d="M7 13.5L7 8C7 8 7.5 6 9.5 6C11.5 6 12 8 12 8L12 10.5L12 8C12 8 12.5 6 14.5 6C16.5 6 17 8 17 8L17 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="32;0"/></path></g>`
	mastodonFilledInnerSVG                                    = `<mask id="lineMdMastodonFilled0"><path fill="#fff" d="M17.25 2.315c2.366.346 4.361 2.131 4.67 4.396c.167 1.683.022 4.421.02 4.87c0 .133-.019 1.34-.027 1.468c-.207 3.236-2.247 4.514-4.391 4.922c-.03.008-.063.014-.097.02c-1.36.263-2.815.333-4.197.372c-.33.008-.66.008-.99.008c-1.373 0-2.742-.16-4.077-.479a.046.046 0 0 0-.059.047c.038.43.133.853.281 1.259c.185.47.832 1.597 3.234 1.597a17.89 17.89 0 0 0 4.146-.479a.048.048 0 0 1 .053.025a.047.047 0 0 1 .005.02v1.589a.05.05 0 0 1-.02.038c-.444.318-1.048.5-1.562.661c-.228.071-.459.133-.692.187c-2.12.478-4.331.362-6.388-.333c-1.921-.667-3.882-2.302-4.367-4.266a22.953 22.953 0 0 1-.545-3.233c-.151-1.64-.164-3.282-.229-4.928c-.045-1.148-.019-2.4.226-3.528c.51-2.292 2.61-3.896 4.91-4.233c.4-.058 1.15-.27 4.655-.27h.026c3.503 0 5.016.212 5.415.27Z"/><g fill="none" stroke="#000" stroke-dasharray="18" stroke-dashoffset="18" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.2 12L12.2 9C12.2 8 13.1 6.5 14.85 6.5C16.6 6.5 17.5 8 17.5 9L17.5 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="18;0"/></path><path d="M6.9 14L6.9 9C6.9 8 7.8 6.45 9.54 6.5C11.3 6.47 12.2 8 12.2 9L12.2 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMastodonFilled0)"/>`
	mastodonTwotoneInnerSVG                                   = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="80" stroke-dashoffset="80" d="M15.5 21.5C6 23.5 6.5 16.5 7.5 16.5C11 16.5 21 19 21 12.5V8.5C21 4.5 18.5 3 14 3H10C5.5 3 3 4 3 10V13C3 19 5 24 15.5 21.5Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="80;160"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="32" stroke-dashoffset="32" d="M7 13.5L7 8C7 8 7.5 6 9.5 6C11.5 6 12 8 12 8L12 10.5L12 8C12 8 12.5 6 14.5 6C16.5 6 17 8 17 8L17 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="32;0"/></path></g>`
	menuInnerSVG                                              = `<g fill="none" stroke="currentColor" stroke-dasharray="24" stroke-dashoffset="24" stroke-linecap="round" stroke-width="2"><path d="M5 5H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="24;0"/></path><path d="M5 12H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="24;0"/></path><path d="M5 19H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="24;0"/></path></g>`
	menuFoldLeftInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M7 9L4 12L7 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 5H5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M19 12H10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 19H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path></g>`
	menuFoldRightInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M17 9L20 12L17 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5 5H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M5 12H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5 19H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path></g>`
	menuToCloseAltTransitionInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 5L19 5"><animate fill="freeze" attributeName="d" begin="0.2s" dur="0.4s" values="M5 5L19 5;M5 5L19 19"/></path><path d="M5 12H19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 12H19;M12 12H12"/><set attributeName="opacity" begin="0.4s" to="0"/></path><path d="M5 19L19 19"><animate fill="freeze" attributeName="d" begin="0.2s" dur="0.4s" values="M5 19L19 19;M5 19L19 5"/></path></g>`
	menuToCloseTransitionInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 5L12 5L19 5"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 5L12 5L19 5;M5 5L12 12L19 5"/></path><path d="M5 12H19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 12H19;M12 12H12"/></path><path d="M5 19L12 19L19 19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 19L12 19L19 19;M5 19L12 12L19 19"/></path></g>`
	menuUnfoldLeftInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 9L18 12L21 15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 5H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M14 12H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 19H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/></path></g>`
	menuUnfoldRightInnerSVG                                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M3 9L6 12L3 15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5 5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M10 12H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5 19H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/></path></g>`
	minusInnerSVG                                             = `<path stroke="currentColor" stroke-dasharray="18" stroke-dashoffset="18" stroke-linecap="round" stroke-width="2" d="M5 12H19" fill="currentColor"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="18;0"/></path>`
	minusCircleInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`
	minusCircleTwotoneInnerSVG                                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`
	moonInnerSVG                                              = `<g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></path></g><path fill="none" stroke="currentColor" stroke-dasharray="56" stroke-dashoffset="56" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path>`
	moonAltLoopInnerSVG                                       = `<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop0.begin+2s;lineMdMoonAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop0.begin+1.2s;lineMdMoonAltLoop0.begin+3.2s;lineMdMoonAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop1.begin+2s;lineMdMoonAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop1.begin+1.2s;lineMdMoonAltLoop1.begin+3.2s;lineMdMoonAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop2.begin+1.2s;lineMdMoonAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><path fill="none" stroke="currentColor" stroke-dasharray="56" stroke-dashoffset="56" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path>`
	moonAltToSunnyOutlineLoopTransitionInnerSVG               = `<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="1.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><mask id="lineMdMoonAltToSunnyOutlineLoopTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="8"><animate fill="freeze" attributeName="r" dur="0.4s" values="8;4"/></circle><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonAltToSunnyOutlineLoopTransition0)"><animate fill="freeze" attributeName="r" dur="0.4s" values="10;6"/></circle>`
	moonFilledInnerSVG                                        = `<g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.5s" values="0;1"/></path></g>`
	moonFilledAltLoopInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonFilledAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonFilledAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop0.begin+2s;lineMdMoonFilledAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop0.begin+1.2s;lineMdMoonFilledAltLoop0.begin+3.2s;lineMdMoonFilledAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonFilledAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonFilledAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonFilledAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop1.begin+2s;lineMdMoonFilledAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop1.begin+1.2s;lineMdMoonFilledAltLoop1.begin+3.2s;lineMdMoonFilledAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonFilledAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonFilledAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonFilledAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop2.begin+1.2s;lineMdMoonFilledAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonFilledAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.5s" values="0;1"/></path></g>`
	moonFilledAltToSunnyFilledLoopTransitionInnerSVG          = `<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="1.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><mask id="lineMdMoonFilledAltToSunnyFilledLoopTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonFilledAltToSunnyFilledLoopTransition0)"><animate fill="freeze" attributeName="r" dur="0.4s" values="10;6"/></circle>`
	moonFilledLoopInnerSVG                                    = `<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonFilledLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonFilledLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.5s" values="0;1"/></path></g>`
	moonFilledToSunnyFilledLoopTransitionInnerSVG             = `<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="1.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><g fill="currentColor"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.4s" values="1;0"/></path></g><g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"/><set attributeName="opacity" begin="0.6s" to="0"/></g><mask id="lineMdMoonFilledToSunnyFilledLoopTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonFilledToSunnyFilledLoopTransition0)" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;6"/></circle>`
	moonFilledToSunnyFilledTransitionInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="1.2s" values="2;0"/></path></g><g fill="currentColor"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.4s" values="1;0"/></path></g><g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"/><set attributeName="opacity" begin="0.6s" to="0"/></g><mask id="lineMdMoonFilledToSunnyFilledTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonFilledToSunnyFilledTransition0)" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;6"/></circle>`
	moonLoopInnerSVG                                          = `<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><path fill="none" stroke="currentColor" stroke-dasharray="56" stroke-dashoffset="56" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path>`
	moonRisingAltLoopInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonRisingAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonRisingAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop0.begin+2s;lineMdMoonRisingAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop0.begin+1.2s;lineMdMoonRisingAltLoop0.begin+3.2s;lineMdMoonRisingAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonRisingAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonRisingAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop1.begin+2s;lineMdMoonRisingAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop1.begin+1.2s;lineMdMoonRisingAltLoop1.begin+3.2s;lineMdMoonRisingAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonRisingAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonRisingAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop2.begin+1.2s;lineMdMoonRisingAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`
	moonRisingFilledAltLoopInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonRisingFilledAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonRisingFilledAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop0.begin+2s;lineMdMoonRisingFilledAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop0.begin+1.2s;lineMdMoonRisingFilledAltLoop0.begin+3.2s;lineMdMoonRisingFilledAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonRisingFilledAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonRisingFilledAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop1.begin+2s;lineMdMoonRisingFilledAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop1.begin+1.2s;lineMdMoonRisingFilledAltLoop1.begin+3.2s;lineMdMoonRisingFilledAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonRisingFilledAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonRisingFilledAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop2.begin+1.2s;lineMdMoonRisingFilledAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`
	moonRisingFilledLoopInnerSVG                              = `<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonRisingFilledLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonRisingFilledLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`
	moonRisingLoopInnerSVG                                    = `<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonRisingLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonRisingLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`
	moonRisingTwotoneAltLoopInnerSVG                          = `<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonRisingTwotoneAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonRisingTwotoneAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop0.begin+2s;lineMdMoonRisingTwotoneAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop0.begin+1.2s;lineMdMoonRisingTwotoneAltLoop0.begin+3.2s;lineMdMoonRisingTwotoneAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonRisingTwotoneAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonRisingTwotoneAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop1.begin+2s;lineMdMoonRisingTwotoneAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop1.begin+1.2s;lineMdMoonRisingTwotoneAltLoop1.begin+3.2s;lineMdMoonRisingTwotoneAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonRisingTwotoneAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonRisingTwotoneAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop2.begin+1.2s;lineMdMoonRisingTwotoneAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><path fill="currentColor" fill-opacity=".3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`
	moonRisingTwotoneLoopInnerSVG                             = `<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonRisingTwotoneLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonRisingTwotoneLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><path fill="currentColor" fill-opacity=".3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`
	moonToSunnyOutlineLoopTransitionInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="1.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><g fill="currentColor"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.4s" values="1;0"/></path></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"/><set attributeName="opacity" begin="0.6s" to="0"/></g><mask id="lineMdMoonToSunnyOutlineLoopTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="8"><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="8;4"/></circle><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonToSunnyOutlineLoopTransition0)" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;6"/></circle>`
	moonToSunnyOutlineTransitionInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="1.2s" values="2;0"/></path></g><g fill="currentColor"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.4s" values="1;0"/></path></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"/><set attributeName="opacity" begin="0.6s" to="0"/></g><mask id="lineMdMoonToSunnyOutlineTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="8"><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="8;4"/></circle><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonToSunnyOutlineTransition0)" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;6"/></circle>`
	moonTwotoneInnerSVG                                       = `<g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;0.3"/></path></g>`
	moonTwotoneAltLoopInnerSVG                                = `<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonTwotoneAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonTwotoneAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop0.begin+2s;lineMdMoonTwotoneAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop0.begin+1.2s;lineMdMoonTwotoneAltLoop0.begin+3.2s;lineMdMoonTwotoneAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonTwotoneAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonTwotoneAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop1.begin+2s;lineMdMoonTwotoneAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop1.begin+1.2s;lineMdMoonTwotoneAltLoop1.begin+3.2s;lineMdMoonTwotoneAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonTwotoneAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonTwotoneAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop2.begin+1.2s;lineMdMoonTwotoneAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;0.3"/></path></g>`
	moonTwotoneLoopInnerSVG                                   = `<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonTwotoneLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonTwotoneLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;0.3"/></path></g>`
	navigationLeftDownInnerSVG                                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M21 9H12C9.23858 9 7 11.2386 7 14V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8"><path d="M7 21L3 17"/><path d="M7 21L11 17"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></g></g>`
	navigationLeftUpInnerSVG                                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M21 15H12C9.23858 15 7 12.7614 7 10V4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 3L3 7M7 3L11 7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></path></g>`
	navigationRightDownInnerSVG                               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M3 9H12C14.76142 9 17 11.2386 17 14V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8"><path d="M17 21L21 17"/><path d="M17 21L13 17"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></g></g>`
	navigationRightUpInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M3 15H12C14.76142 15 17 12.7614 17 10V4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8"><path d="M17 3L21 7"/><path d="M17 3L13 7"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></g></g>`
	paintDropInnerSVG                                         = `<path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`
	paintDropFilledInnerSVG                                   = `<path fill="currentColor" fill-opacity="0" d="M12 4C12 4 18 9 18 15C18 19 15 20 12 20C9 20 6 19 6 15C6 9 12 4 12 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.5s" values="0;1"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`
	paintDropHalfFilledInnerSVG                               = `<path fill="currentColor" fill-opacity="0" d="M6 15C6 9 12 4 12 4C12 4 14.9522 6.46019 16.715 10L6.8347 18C6.31173 17.2671 6 16.2894 6 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.5s" values="0;1"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`
	paintDropHalfFilledTwotoneInnerSVG                        = `<path fill="currentColor" fill-opacity="0" d="M6 15C6 9 12 4 12 4C12 4 14.9522 6.46019 16.715 10L6.8347 18C6.31173 17.2671 6 16.2894 6 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.5s" values="0;1"/></path><path fill="currentColor" fill-opacity="0" d="M12 4C12 4 18 9 18 15C18 19 15 20 12 20C9 20 6 19 6 15C6 9 12 4 12 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`
	paintDropHalfTwotoneInnerSVG                              = `<path fill="currentColor" fill-opacity="0" d="M6 15C6 9 12 4 12 4C12 4 14.9522 6.46019 16.715 10L6.8347 18C6.31173 17.2671 6 16.2894 6 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`
	paintDropTwotoneInnerSVG                                  = `<path fill="currentColor" fill-opacity="0" d="M12 4C12 4 18 9 18 15C18 19 15 20 12 20C9 20 6 19 6 15C6 9 12 4 12 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`
	pauseInnerSVG                                             = `<g fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6h2v12h-2z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path d="M15 6h2v12h-2z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="32;0"/></path></g>`
	pauseToPlayFilledTransitionInnerSVG                       = `<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 6L9 18L7 18L7 6z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M9 18L7 18L7 6L9 6L9 18;M13 15L8 18L8 6L13 9L13 15"/><set attributeName="opacity" begin="0.4s" to="0"/></path><path d="M15 6L17 6L17 18L15 18L15 6"><animate fill="freeze" attributeName="d" dur="0.4s" values="M15 6L17 6L17 18L15 18L15 6;M13 9L18 12L18 12L13 15L13 9"/><set attributeName="opacity" begin="0.4s" to="0"/></path><path d="M8 6L18 12L8 18z" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/></path></g>`
	pauseToPlayTransitionInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 6L9 18L7 18L7 6z"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.66;1" values="M9 18L7 18L7 6L9 6L9 18;M13 15L8 18L8 6L13 9L13 15;M13 15L8 18L8 6L13 9L13 9"/><set attributeName="opacity" begin="0.6s" to="0"/></path><path d="M15 6L17 6L17 18L15 18L15 6"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.66;1" values="M15 6L17 6L17 18L15 18L15 6;M13 9L18 12L18 12L13 15L13 9;M13 9L18 12L18 12L13 15L13 15"/><set attributeName="opacity" begin="0.6s" to="0"/></path><path d="M8 6L18 12L8 18z" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/></path></g>`
	pencilInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M3 21L4.99998 15L16 4C17 3 19 3 20 4C21 5 21 7 20 8L8.99998 19L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path><g stroke-dasharray="6" stroke-dashoffset="6"><path d="M15 5L19 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path><path stroke-width="1" d="M6 15L9 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g></g>`
	pencilTwotoneInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M3 21L4.99998 15L16 4C17 3 19 3 20 4C21 5 21 7 20 8L8.99998 19L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path><g stroke-dasharray="6" stroke-dashoffset="6"><path d="M15 5L19 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path><path stroke-width="1" d="M6 15L9 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g></g><path fill="currentColor" fill-opacity="0" d="M17 4H20V7L9 18L6 15L17 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path>`
	pencilTwotoneAltInnerSVG                                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M3 21L4.99998 15L16 4C17 3 19 3 20 4C21 5 21 7 20 8L8.99998 19L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path><g stroke-dasharray="6" stroke-dashoffset="6"><path d="M15 5L19 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path><path stroke-width="1" d="M6 15L9 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g></g><path fill="currentColor" fill-opacity="0" d="M9 18L18 9L15 6L6 15L9 18Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path>`
	playInnerSVG                                              = `<path fill="none" stroke="currentColor" stroke-dasharray="36" stroke-dashoffset="36" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6L18 12L8 18z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/></path>`
	playFilledInnerSVG                                        = `<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="36" stroke-dashoffset="36" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6L18 12L8 18z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path>`
	playFilledToPauseTransitionInnerSVG                       = `<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 15L8 18L8 6L13 9L13 15"><animate fill="freeze" attributeName="d" dur="0.4s" values="M13 15L8 18L8 6L13 9L13 15;M9 18L7 18L7 6L9 6L9 18"/></path><path d="M13 9L18 12L18 12L13 15L13 9"><animate fill="freeze" attributeName="d" dur="0.4s" values="M13 9L18 12L18 12L13 15L13 9;M15 6L17 6L17 18L15 18L15 6"/></path></g>`
	playToPauseTransitionInnerSVG                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 15L8 18L8 6L13 9L13 9"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.33;1" values="M13 15L8 18L8 6L13 9L13 9;M13 15L8 18L8 6L13 9L13 15;M9 18L7 18L7 6L9 6L9 18"/></path><path d="M13 9L18 12L18 12L13 15L13 15"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.33;1" values="M13 9L18 12L18 12L13 15L13 15;M13 9L18 12L18 12L13 15L13 9;M15 6L17 6L17 18L15 18L15 6"/></path></g>`
	playTwotoneInnerSVG                                       = `<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="36" stroke-dashoffset="36" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6L18 12L8 18z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path>`
	plusInnerSVG                                              = `<g fill="none" stroke="currentColor" stroke-dasharray="18" stroke-dashoffset="18" stroke-linecap="round" stroke-width="2"><path d="M12 5V19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.3s" values="18;0"/></path><path d="M5 12H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="18;0"/></path></g>`
	plusCircleInnerSVG                                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g stroke-dasharray="12" stroke-dashoffset="12"><path d="M12 7V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g><path stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path></g>`
	plusCircleTwotoneInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g stroke-dasharray="12" stroke-dashoffset="12"><path d="M12 7V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path></g>`
	questionInnerSVG                                          = `<path fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-width="2" d="M7 8C7 5.23858 9.23857 3 12 3C14.7614 3 17 5.23858 17 8C17 9.6356 16.2147 11.0878 15.0005 12C14.1647 12.6279 12 14 12 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/></path><circle cx="12" cy="21" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.2s" values="0;1"/></circle>`
	questionCircleInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M8.99999 10C8.99999 8.34315 10.3431 7 12 7C13.6569 7 15 8.34315 15 10C15 10.9814 14.5288 11.8527 13.8003 12.4C13.0718 12.9473 12.5 13 12 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="20;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.2s" values="0;1"/></circle>`
	questionCircleTwotoneInnerSVG                             = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M8.99999 10C8.99999 8.34315 10.3431 7 12 7C13.6569 7 15 8.34315 15 10C15 10.9814 14.5288 11.8527 13.8003 12.4C13.0718 12.9473 12.5 13 12 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="20;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.2s" values="0;1"/></circle>`
	redditInnerSVG                                            = `<mask id="lineMdReddit0"><g fill="#fff" fill-opacity="0"><ellipse cx="12" cy="14.71" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="8" ry="5.29"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></ellipse><circle cx="7.24" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="7.24;3.94"/></circle><circle cx="16.76" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="16.76;20.06"/></circle><circle cx="18.45" cy="4.23" r="1.61"><set attributeName="fill-opacity" begin="2.6s" to="1"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="12" stroke-dashoffset="12" stroke-linecap="round" stroke-linejoin="round" stroke-width=".8" d="M12 8.75L13.18 3.11L18.21 4.18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.4s" dur="0.2s" values="12;0"/></path><g fill-opacity="0"><circle cx="8.45" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15.55" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width=".8" d="M8.47 17.52C8.47 17.52 9.41 18.58 12 18.58C14.58 18.58 15.53 17.52 15.53 17.52"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2s" dur="0.2s" values="8;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdReddit0)"/>`
	redditCircleInnerSVG                                      = `<mask id="lineMdRedditCircle0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path><g fill-opacity="0"><ellipse cx="12" cy="13.77" rx="5.83" ry="4.06"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></ellipse><circle cx="8.99" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="8.99;6.79"/></circle><circle cx="15.01" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="15.01;17.21"/></circle><circle cx="16.18" cy="7.01" r="1.04"><set attributeName="fill-opacity" begin="3.1s" to="1"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width=".54" d="M12 9.91L12.76 6.27L16 6.98"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.9s" dur="0.2s" values="8;0"/></path><g fill="#fff" fill-opacity="0"><circle cx="9.71" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.4s" values="0;1"/></circle><circle cx="14.29" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="6" stroke-dashoffset="6" stroke-linecap="round" stroke-width=".54" d="M9.72 15.6C9.72 15.6 10.33 16.29 12 16.291C13.67 16.29 14.28 15.6 14.28 15.6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.5s" dur="0.2s" values="6;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdRedditCircle0)"/>`
	redditCircleLoopInnerSVG                                  = `<mask id="lineMdRedditCircleLoop0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path><g fill-opacity="0"><ellipse cx="12" cy="13.77" rx="5.83" ry="4.06"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></ellipse><circle cx="8.99" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="8.99;6.79"/></circle><circle cx="15.01" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="15.01;17.21"/></circle><circle cx="16.18" cy="7.01" r="1.04"><set attributeName="fill-opacity" begin="3.1s" to="1"/><animate attributeName="cx" begin="2.9s" dur="6s" repeatCount="indefinite" values="16.18;7.82;16.18"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width=".54" d="M12 9.91L12.76 6.27L16 6.98"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.9s" dur="0.2s" values="8;0"/><animate attributeName="d" begin="2.9s" dur="6s" repeatCount="indefinite" values="M12 9.91L12.76 6.27L16 6.98;M12 9.91L12 5.2L12 6.98;M12 9.91L11.24 6.27L8 6.98;M12 9.91L12 5.2L12 6.98;M12 9.91L12.76 6.27L16 6.98"/></path><g fill="#fff" fill-opacity="0"><circle cx="9.71" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.4s" values="0;1"/></circle><circle cx="14.29" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="6" stroke-dashoffset="6" stroke-linecap="round" stroke-width=".54" d="M9.72 15.6C9.72 15.6 10.33 16.29 12 16.291C13.67 16.29 14.28 15.6 14.28 15.6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.5s" dur="0.2s" values="6;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdRedditCircleLoop0)"/>`
	redditLoopInnerSVG                                        = `<mask id="lineMdRedditLoop0"><g fill="#fff" fill-opacity="0"><ellipse cx="12" cy="14.71" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="8" ry="5.29"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></ellipse><circle cx="7.24" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="7.24;3.94"/></circle><circle cx="16.76" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="16.76;20.06"/></circle><circle cx="18.45" cy="4.23" r="1.61"><set attributeName="fill-opacity" begin="2.6s" to="1"/><animate attributeName="cx" begin="2.4s" dur="6s" repeatCount="indefinite" values="18.45;5.75;18.45"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="12" stroke-dashoffset="12" stroke-linecap="round" stroke-linejoin="round" stroke-width=".8" d="M12 8.75L13.18 3.11L18.21 4.18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.4s" dur="0.2s" values="12;0"/><animate attributeName="d" begin="2.4s" dur="6s" repeatCount="indefinite" values="M12 8.75L13.18 3.11L18.21 4.18;M12 8.75L12 2L12 4.18;M12 8.75L10.82 3.11L5.79 4.18;M12 8.75L12 2L12 4.18;M12 8.75L13.18 3.11L18.21 4.18"/></path><g fill-opacity="0"><circle cx="8.45" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15.55" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width=".8" d="M8.47 17.52C8.47 17.52 9.41 18.58 12 18.58C14.58 18.58 15.53 17.52 15.53 17.52"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2s" dur="0.2s" values="8;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdRedditLoop0)"/>`
	removeInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-dasharray="22" stroke-dashoffset="22" stroke-linecap="round" stroke-width="2"><path d="M19 5L5 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="22;0"/></path><path d="M5 5L19 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="22;0"/></path></g>`
	rotateNinetyInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M12 6C15.3137 6 18 8.68629 18 12V14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 15L21 12M18 15L15 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="6;0"/></path></g>`
	rotateOneHundredEightyInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="24" stroke-dashoffset="24" d="M12 6C15.3137 6 18 8.68629 18 12C18 15.3137 15.3137 18 12 18H9.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="24;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 18L12 21M9 18L12 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="6;0"/></path></g>`
	rotateTwoHundredSeventyInnerSVG                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M12 6C15.3137 6 18 8.68629 18 12C18 15.3137 15.3137 18 12 18C8.68629 18 6 15.3137 6 12V9.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="32;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M6 9L3 12M6 9L9 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`
	searchInnerSVG                                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M10.5 13.5L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path><path stroke-dasharray="40" stroke-dashoffset="40" d="M10.7574 13.2426C8.41421 10.8995 8.41421 7.10051 10.7574 4.75736C13.1005 2.41421 16.8995 2.41421 19.2426 4.75736C21.5858 7.10051 21.5858 10.8995 19.2426 13.2426C16.8995 15.5858 13.1005 15.5858 10.7574 13.2426Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="40;0"/></path></g>`
	searchFilledInnerSVG                                      = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="none" stroke-dasharray="16" stroke-dashoffset="16" d="M10.5 13.5L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M10.7574 13.2426C8.41421 10.8995 8.41421 7.10051 10.7574 4.75736C13.1005 2.41421 16.8995 2.41421 19.2426 4.75736C21.5858 7.10051 21.5858 10.8995 19.2426 13.2426C16.8995 15.5858 13.1005 15.5858 10.7574 13.2426Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path></g>`
	searchTwotoneInnerSVG                                     = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="none" stroke-dasharray="16" stroke-dashoffset="16" d="M10.5 13.5L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M10.7574 13.2426C8.41421 10.8995 8.41421 7.10051 10.7574 4.75736C13.1005 2.41421 16.8995 2.41421 19.2426 4.75736C21.5858 7.10051 21.5858 10.8995 19.2426 13.2426C16.8995 15.5858 13.1005 15.5858 10.7574 13.2426Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></g>`
	sunRisingFilledLoopInnerSVG                               = `<circle cx="12" cy="32" r="6" fill="currentColor"><animate fill="freeze" attributeName="cy" dur="0.6s" values="32;12"/></circle><g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g>`
	sunRisingLoopInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="32" r="5"><animate fill="freeze" attributeName="cy" dur="0.6s" values="32;12"/></circle><g stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`
	sunRisingTwotoneLoopInnerSVG                              = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="32" r="5" fill="currentColor" fill-opacity=".3"><animate fill="freeze" attributeName="cy" dur="0.6s" values="32;12"/></circle><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`
	sunnyFilledInnerSVG                                       = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path></g></g>`
	sunnyFilledLoopInnerSVG                                   = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`
	sunnyFilledLoopToMoonAltFilledLoopTransitionInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path fill="currentColor" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0" fill="freeze" attributeName="stroke-dashoffset" begin="0.6s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+1.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+3.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1" fill="freeze" attributeName="stroke-dashoffset" begin="1s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+1.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+3.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2" fill="freeze" attributeName="stroke-dashoffset" begin="2.8s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+1.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><mask id="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition3"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyFilledLoopToMoonAltFilledLoopTransition3)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`
	sunnyFilledLoopToMoonFilledLoopTransitionInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path fill="currentColor" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdSunnyFilledLoopToMoonFilledLoopTransition0" fill="freeze" attributeName="fill-opacity" begin="0.6s;lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><mask id="lineMdSunnyFilledLoopToMoonFilledLoopTransition1"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyFilledLoopToMoonFilledLoopTransition1)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`
	sunnyFilledLoopToMoonFilledTransitionInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path fill="currentColor" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></path></g><mask id="lineMdSunnyFilledLoopToMoonFilledTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyFilledLoopToMoonFilledTransition0)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`
	sunnyOutlineInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/></path><g stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path></g></g>`
	sunnyOutlineLoopInnerSVG                                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/></path><g stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`
	sunnyOutlineToMoonAltLoopTransitionInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdSunnyOutlineToMoonAltLoopTransition0" fill="freeze" attributeName="stroke-dashoffset" begin="0.6s;lineMdSunnyOutlineToMoonAltLoopTransition0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+2s;lineMdSunnyOutlineToMoonAltLoopTransition0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+1.2s;lineMdSunnyOutlineToMoonAltLoopTransition0.begin+3.2s;lineMdSunnyOutlineToMoonAltLoopTransition0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdSunnyOutlineToMoonAltLoopTransition1" fill="freeze" attributeName="stroke-dashoffset" begin="1s;lineMdSunnyOutlineToMoonAltLoopTransition1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+2s;lineMdSunnyOutlineToMoonAltLoopTransition1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+1.2s;lineMdSunnyOutlineToMoonAltLoopTransition1.begin+3.2s;lineMdSunnyOutlineToMoonAltLoopTransition1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdSunnyOutlineToMoonAltLoopTransition2" fill="freeze" attributeName="stroke-dashoffset" begin="2.8s;lineMdSunnyOutlineToMoonAltLoopTransition2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition2.begin+1.2s;lineMdSunnyOutlineToMoonAltLoopTransition2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><mask id="lineMdSunnyOutlineToMoonAltLoopTransition3"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="4"><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="4;8"/></circle><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyOutlineToMoonAltLoopTransition3)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`
	sunnyOutlineToMoonLoopTransitionInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdSunnyOutlineToMoonLoopTransition0" fill="freeze" attributeName="fill-opacity" begin="0.6s;lineMdSunnyOutlineToMoonLoopTransition0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><mask id="lineMdSunnyOutlineToMoonLoopTransition1"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="4"><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="4;8"/></circle><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyOutlineToMoonLoopTransition1)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`
	sunnyOutlineToMoonTransitionInnerSVG                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></path></g><mask id="lineMdSunnyOutlineToMoonTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="4"><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="4;8"/></circle><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyOutlineToMoonTransition0)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`
	sunnyOutlineTwotoneInnerSVG                               = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path></g></g>`
	sunnyOutlineTwotoneLoopInnerSVG                           = `<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`
	switchInnerSVG                                            = `<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-dasharray="54" stroke-dashoffset="54" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><circle cx="17" cy="12" r="3" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></circle>`
	switchFilledInnerSVG                                      = `<mask id="lineMdSwitchFilled0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="8" x="3" y="8" stroke-dasharray="54" stroke-dashoffset="54" rx="4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></g><circle cx="17" cy="12" r="3" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></mask><rect width="20" height="10" x="2" y="7" fill="currentColor" mask="url(#lineMdSwitchFilled0)"/>`
	switchFilledToSwitchOffFilledTransitionInnerSVG           = `<path fill="currentColor" d="M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm-10 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"><animate fill="freeze" attributeName="d" dur="0.2s" values="M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm-10 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z;M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm0 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/></path>`
	switchOffInnerSVG                                         = `<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-dasharray="54" stroke-dashoffset="54" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><circle cx="7" cy="12" r="3" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></circle>`
	switchOffFilledInnerSVG                                   = `<mask id="lineMdSwitchOffFilled0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="8" x="3" y="8" stroke-dasharray="54" stroke-dashoffset="54" rx="4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></g><circle cx="7" cy="12" r="3" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></mask><rect width="20" height="10" x="2" y="7" fill="currentColor" mask="url(#lineMdSwitchOffFilled0)"/>`
	switchOffFilledToSwitchFilledTransitionInnerSVG           = `<path fill="currentColor" d="M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm0 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"><animate fill="freeze" attributeName="d" dur="0.2s" values="M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm0 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z;M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm-10 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/></path>`
	switchOffToSwitchTransitionInnerSVG                       = `<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"/><circle cx="7" cy="12" r="3" fill="currentColor"><animate fill="freeze" attributeName="cx" dur="0.2s" values="7;17"/></circle>`
	switchToSwitchOffTransitionInnerSVG                       = `<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"/><circle cx="17" cy="12" r="3" fill="currentColor"><animate fill="freeze" attributeName="cx" dur="0.2s" values="17;7"/></circle>`
	telegramInnerSVG                                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M21 5L18.5 20M21 5L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="22" stroke-dashoffset="22" d="M21 5L2 12.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="22;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M18.5 20L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M2 12.5L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.3s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 16L9 19M9 13.5L9 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.3s" values="6;0"/></path></g>`
	textBoxInnerSVG                                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="62" stroke-dashoffset="62" d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M8 8h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M8 12h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M8 16h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path></g>`
	textBoxMultipleInnerSVG                                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="62" stroke-dashoffset="62" d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M10 6h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M10 10h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M10 14h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.4s" values="34;68"/></path></g>`
	textBoxMultipleToTextBoxTransitionInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17zM10 6h8M10 10h8M10 14h5"><animateMotion fill="freeze" begin="0.4s" calcMode="linear" dur="0.4s" path="M0 0l-2 2"/></path><path stroke-dasharray="34" stroke-dashoffset="68" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="68;34"/></path></g>`
	textBoxMultipleTwotoneInnerSVG                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="62" stroke-dashoffset="62" d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/><animate fill="freeze" attributeName="fill-opacity" begin="1.8s" dur="0.15s" values="0;0.3"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M10 6h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M10 10h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M10 14h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.4s" values="34;68"/></path></g>`
	textBoxMultipleTwotoneToTextBoxTwotoneTransitionInnerSVG  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity=".3" d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17zM10 6h8M10 10h8M10 14h5"><animateMotion fill="freeze" begin="0.4s" calcMode="linear" dur="0.4s" path="M0 0l-2 2"/></path><path stroke-dasharray="34" stroke-dashoffset="68" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="68;34"/></path></g>`
	textBoxToTextBoxMultipleTransitionInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19zM8 8h8M8 12h8M8 16h5"><animateMotion fill="freeze" calcMode="linear" dur="0.4s" path="M0 0l2-2"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="34;68"/></path></g>`
	textBoxTwotoneInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="62" stroke-dashoffset="62" d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/><animate fill="freeze" attributeName="fill-opacity" begin="1.3s" dur="0.15s" values="0;0.3"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M8 8h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M8 12h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M8 16h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path></g>`
	textBoxTwotoneToTextBoxMultipleTwotoneTransitionInnerSVG  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity=".3" d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19zM8 8h8M8 12h8M8 16h5"><animateMotion fill="freeze" calcMode="linear" dur="0.4s" path="M0 0l2-2"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="34;68"/></path></g>`
	thumbsDownInnerSVG                                        = `<path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4H18L21 11V14H14L15 20L12 21L7 13H3V4H7V13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`
	thumbsDownTwotoneInnerSVG                                 = `<path fill="currentColor" fill-opacity="0" d="M3 4H7V13H3V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4H18L21 11V14H14L15 20L12 21L7 13H3V4H7V13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`
	thumbsUpInnerSVG                                          = `<path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11L12 3L15 4L14 10H21V13L18 20H7H3V11H7V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`
	thumbsUpTwotoneInnerSVG                                   = `<path fill="currentColor" fill-opacity="0" d="M7 11V20H3V11H7Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11L12 3L15 4L14 10H21V13L18 20H7H3V11H7V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`
	twitterInnerSVG                                           = `<path fill="none" stroke="currentColor" stroke-dasharray="62" stroke-dashoffset="62" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.8906 7.34375C19.7969 7.67188 19.4001 8.50548 18.7219 9.29669C18.2698 17.9717 9.84907 20.7974 4.08456 17.8869C3.29335 16.8414 6.93856 17.2653 8.26666 15.259C3.23684 12.6876 3.63244 5.82103 4.64971 6.1036C7.02333 9.29669 10.8381 9.57926 11.4597 9.29669C11.4597 8.562 11.1489 6.97958 12.8726 5.65148C13.8616 4.94505 15.9297 4.3125 17.8047 6.34375C18.125 6.55469 18.5859 6.64844 19.2734 6.49219C19.6797 6.28125 20.2262 6.427 19.9453 7.15625"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;0"/></path>`
	twitterTwotoneInnerSVG                                    = `<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="62" stroke-dashoffset="62" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.8906 7.34375C19.7969 7.67188 19.4001 8.50548 18.7219 9.29669C18.2698 17.9717 9.84907 20.7974 4.08456 17.8869C3.29335 16.8414 6.93856 17.2653 8.26666 15.259C3.23684 12.6876 3.63244 5.82103 4.64971 6.1036C7.02333 9.29669 10.8381 9.57926 11.4597 9.29669C11.4597 8.562 11.1489 6.97958 12.8726 5.65148C13.8616 4.94505 15.9297 4.3125 17.8047 6.34375C18.125 6.55469 18.5859 6.64844 19.2734 6.49219C19.6797 6.28125 20.2262 6.427 19.9453 7.15625"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path>`
	uploadLoopInnerSVG                                        = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="14;0"/></path><path fill="currentColor" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5;M12 15 h2 v-3 h2.5 L12 7.5M12 15 h-2 v-3 h-2.5 L12 7.5;M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"/></path></g>`
	uploadOffLoopInnerSVG                                     = `<mask id="lineMdUploadOffLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="14;0"/></path><path fill="#fff" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5;M12 15 h2 v-3 h2.5 L12 7.5M12 15 h-2 v-3 h-2.5 L12 7.5;M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdUploadOffLoop0)"/>`
	uploadOffOutlineInnerSVG                                  = `<mask id="lineMdUploadOffOutline0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdUploadOffOutline0)"/>`
	uploadOffOutlineLoopInnerSVG                              = `<mask id="lineMdUploadOffOutlineLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5;M12 15 h2 v-3 h2.5 L12 7.5M12 15 h-2 v-3 h-2.5 L12 7.5;M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdUploadOffOutlineLoop0)"/>`
	uploadOutlineInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path></g>`
	uploadOutlineLoopInnerSVG                                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5;M12 15 h2 v-3 h2.5 L12 7.5M12 15 h-2 v-3 h-2.5 L12 7.5;M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"/></path></g>`
	uploadingLoopInnerSVG                                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="2 4" stroke-dashoffset="6" d="M12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3"><animate attributeName="stroke-dashoffset" dur="0.6s" repeatCount="indefinite" values="6;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.3s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 16v-7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 8.5l3.5 3.5M12 8.5l-3.5 3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="6;0"/></path></g>`
	valignBaselineInnerSVG                                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 18.5H21.5M2.5 11.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`
	valignBaselineTwotoneInnerSVG                             = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 18.5H21.5M2.5 11.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`
	valignBottomInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 9H18V21H6V9H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`
	valignBottomTwotoneInnerSVG                               = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 9H18V21H6V9H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`
	valignMiddleInnerSVG                                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`
	valignMiddleTwotoneInnerSVG                               = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`
	valignTopInnerSVG                                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 3H18V15H6V3H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`
	valignTopTwotoneInnerSVG                                  = `<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 3H18V15H6V3H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`
	weatherCloudyLoopInnerSVG                                 = `<mask id="lineMdWeatherCloudyLoop0"><g fill="#fff"><circle cx="12" cy="11" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="10" height="7" x="8" y="12"/><rect width="16" height="10" x="1" y="9" rx="5"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="8" x="6" y="11" rx="4"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="11" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="6" x="8" y="11"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="12" height="6" x="3" y="11" rx="3"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="4" x="8" y="13" rx="2"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="8;7;8;9;8"/></rect></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWeatherCloudyLoop0)"/>`
)

var sharedIconAttrs = engine.Attrs{"width": "1em", "height": "1em"}

func Account(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		accountInnerSVG,
		children,
	)
}

func AccountAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		accountAddInnerSVG,
		children,
	)
}

func AccountAlert(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		accountAlertInnerSVG,
		children,
	)
}

func AccountDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		accountDeleteInnerSVG,
		children,
	)
}

func AccountRemove(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		accountRemoveInnerSVG,
		children,
	)
}

func AccountSmall(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		accountSmallInnerSVG,
		children,
	)
}

func Alert(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertInnerSVG,
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

func AlertCircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertCircleTwotoneInnerSVG,
		children,
	)
}

func AlertTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertTwotoneInnerSVG,
		children,
	)
}

func AlignCenter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignCenterInnerSVG,
		children,
	)
}

func AlignJustify(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignJustifyInnerSVG,
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

func ArrowAlignCenter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowAlignCenterInnerSVG,
		children,
	)
}

func ArrowAlignLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowAlignLeftInnerSVG,
		children,
	)
}

func ArrowCloseLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCloseLeftInnerSVG,
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

func ArrowLeftCircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLeftCircleTwotoneInnerSVG,
		children,
	)
}

func ArrowLongDiagonal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLongDiagonalInnerSVG,
		children,
	)
}

func ArrowOpenLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowOpenLeftInnerSVG,
		children,
	)
}

func ArrowSmallLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowSmallLeftInnerSVG,
		children,
	)
}

func ArrowsDiagonal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsDiagonalInnerSVG,
		children,
	)
}

func ArrowsHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsHorizontalInnerSVG,
		children,
	)
}

func ArrowsHorizontalAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsHorizontalAltInnerSVG,
		children,
	)
}

func Beer(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerInnerSVG,
		children,
	)
}

func BeerAltFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerAltFilledInnerSVG,
		children,
	)
}

func BeerAltFilledLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerAltFilledLoopInnerSVG,
		children,
	)
}

func BeerAltTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerAltTwotoneInnerSVG,
		children,
	)
}

func BeerAltTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerAltTwotoneLoopInnerSVG,
		children,
	)
}

func BeerFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerFilledInnerSVG,
		children,
	)
}

func BeerLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerLoopInnerSVG,
		children,
	)
}

func BeerTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerTwotoneInnerSVG,
		children,
	)
}

func BeerTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		beerTwotoneLoopInnerSVG,
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

func BellTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bellTwotoneInnerSVG,
		children,
	)
}

func BuyMeACoffee(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		buyMeACoffeeInnerSVG,
		children,
	)
}

func BuyMeACoffeeFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		buyMeACoffeeFilledInnerSVG,
		children,
	)
}

func BuyMeACoffeeTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		buyMeACoffeeTwotoneInnerSVG,
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

func CalendarOut(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarOutInnerSVG,
		children,
	)
}

func Cancel(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cancelInnerSVG,
		children,
	)
}

func CancelTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cancelTwotoneInnerSVG,
		children,
	)
}

func CheckListThree(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkListThreeInnerSVG,
		children,
	)
}

func CheckListThreeFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkListThreeFilledInnerSVG,
		children,
	)
}

func CheckListThreeTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkListThreeTwotoneInnerSVG,
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

func ChevronLeftCircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronLeftCircleTwotoneInnerSVG,
		children,
	)
}

func ChevronSmallDoubleLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronSmallDoubleLeftInnerSVG,
		children,
	)
}

func ChevronSmallLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronSmallLeftInnerSVG,
		children,
	)
}

func ChevronSmallTripleLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronSmallTripleLeftInnerSVG,
		children,
	)
}

func ChevronTripleLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronTripleLeftInnerSVG,
		children,
	)
}

func Circle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		circleInnerSVG,
		children,
	)
}

func CircleToConfirmCircleTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		circleToConfirmCircleTransitionInnerSVG,
		children,
	)
}

func CircleToConfirmCircleTwotoneTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		circleToConfirmCircleTwotoneTransitionInnerSVG,
		children,
	)
}

func CircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		circleTwotoneInnerSVG,
		children,
	)
}

func CircleTwotoneToConfirmCircleTwotoneTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		circleTwotoneToConfirmCircleTwotoneTransitionInnerSVG,
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

func ClipboardArrow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardArrowInnerSVG,
		children,
	)
}

func ClipboardArrowTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardArrowTwotoneInnerSVG,
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

func ClipboardCheckToClipboardTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardCheckToClipboardTransitionInnerSVG,
		children,
	)
}

func ClipboardCheckTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardCheckTwotoneInnerSVG,
		children,
	)
}

func ClipboardCheckTwotoneToClipboardTwotoneTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardCheckTwotoneToClipboardTwotoneTransitionInnerSVG,
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

func ClipboardListTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardListTwotoneInnerSVG,
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

func ClipboardMinusTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardMinusTwotoneInnerSVG,
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

func ClipboardPlusTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardPlusTwotoneInnerSVG,
		children,
	)
}

func ClipboardToClipboardCheckTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardToClipboardCheckTransitionInnerSVG,
		children,
	)
}

func ClipboardTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardTwotoneInnerSVG,
		children,
	)
}

func ClipboardTwotoneToClipboardTwotoneCheckTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardTwotoneToClipboardTwotoneCheckTransitionInnerSVG,
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

func CloseCircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeCircleTwotoneInnerSVG,
		children,
	)
}

func CloseToMenuAltTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeToMenuAltTransitionInnerSVG,
		children,
	)
}

func CloseToMenuTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeToMenuTransitionInnerSVG,
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

func CloudBracesLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudBracesLoopInnerSVG,
		children,
	)
}

func CloudDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownInnerSVG,
		children,
	)
}

func CloudDownTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownTwotoneInnerSVG,
		children,
	)
}

func CloudDownloadLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownloadLoopInnerSVG,
		children,
	)
}

func CloudDownloadOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownloadOutlineLoopInnerSVG,
		children,
	)
}

func CloudFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudFilledInnerSVG,
		children,
	)
}

func CloudLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudLoopInnerSVG,
		children,
	)
}

func CloudOffOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudOffOutlineLoopInnerSVG,
		children,
	)
}

func CloudOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudOutlineLoopInnerSVG,
		children,
	)
}

func CloudPrintLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudPrintLoopInnerSVG,
		children,
	)
}

func CloudPrintOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudPrintOutlineLoopInnerSVG,
		children,
	)
}

func CloudTagsLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudTagsLoopInnerSVG,
		children,
	)
}

func CloudTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudTwotoneInnerSVG,
		children,
	)
}

func CloudUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUpInnerSVG,
		children,
	)
}

func CloudUpTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUpTwotoneInnerSVG,
		children,
	)
}

func CloudUploadLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUploadLoopInnerSVG,
		children,
	)
}

func CloudUploadOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUploadOutlineLoopInnerSVG,
		children,
	)
}

func Coffee(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeInnerSVG,
		children,
	)
}

func CoffeeArrow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeArrowInnerSVG,
		children,
	)
}

func CoffeeArrowFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeArrowFilledInnerSVG,
		children,
	)
}

func CoffeeArrowTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeArrowTwotoneInnerSVG,
		children,
	)
}

func CoffeeFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeFilledInnerSVG,
		children,
	)
}

func CoffeeHalfEmptyTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeHalfEmptyTwotoneLoopInnerSVG,
		children,
	)
}

func CoffeeLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeLoopInnerSVG,
		children,
	)
}

func CoffeeTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeTwotoneInnerSVG,
		children,
	)
}

func CoffeeTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeTwotoneLoopInnerSVG,
		children,
	)
}

func Computer(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		computerInnerSVG,
		children,
	)
}

func ComputerTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		computerTwotoneInnerSVG,
		children,
	)
}

func Confirm(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		confirmInnerSVG,
		children,
	)
}

func ConfirmCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		confirmCircleInnerSVG,
		children,
	)
}

func ConfirmCircleToCircleTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		confirmCircleToCircleTransitionInnerSVG,
		children,
	)
}

func ConfirmCircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		confirmCircleTwotoneInnerSVG,
		children,
	)
}

func ConfirmCircleTwotoneToCircleTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		confirmCircleTwotoneToCircleTransitionInnerSVG,
		children,
	)
}

func ConfirmCircleTwotoneToCircleTwotoneTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		confirmCircleTwotoneToCircleTwotoneTransitionInnerSVG,
		children,
	)
}

func Construction(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		constructionInnerSVG,
		children,
	)
}

func ConstructionTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		constructionTwotoneInnerSVG,
		children,
	)
}

func Discord(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		discordInnerSVG,
		children,
	)
}

func DiscordTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		discordTwotoneInnerSVG,
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

func DocumentAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentAddInnerSVG,
		children,
	)
}

func DocumentAddTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentAddTwotoneInnerSVG,
		children,
	)
}

func DocumentCode(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentCodeInnerSVG,
		children,
	)
}

func DocumentCodeTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentCodeTwotoneInnerSVG,
		children,
	)
}

func DocumentList(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentListInnerSVG,
		children,
	)
}

func DocumentListTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentListTwotoneInnerSVG,
		children,
	)
}

func DocumentRemove(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentRemoveInnerSVG,
		children,
	)
}

func DocumentRemoveTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentRemoveTwotoneInnerSVG,
		children,
	)
}

func DocumentReport(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentReportInnerSVG,
		children,
	)
}

func DocumentReportTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentReportTwotoneInnerSVG,
		children,
	)
}

func DocumentTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		documentTwotoneInnerSVG,
		children,
	)
}

func DoubleArrowHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		doubleArrowHorizontalInnerSVG,
		children,
	)
}

func DownloadLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadLoopInnerSVG,
		children,
	)
}

func DownloadOffLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadOffLoopInnerSVG,
		children,
	)
}

func DownloadOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadOffOutlineInnerSVG,
		children,
	)
}

func DownloadOffOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadOffOutlineLoopInnerSVG,
		children,
	)
}

func DownloadOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadOutlineInnerSVG,
		children,
	)
}

func DownloadOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadOutlineLoopInnerSVG,
		children,
	)
}

func DownloadingLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadingLoopInnerSVG,
		children,
	)
}

func Edit(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editInnerSVG,
		children,
	)
}

func EditTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editTwotoneInnerSVG,
		children,
	)
}

func EditTwotoneFull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editTwotoneFullInnerSVG,
		children,
	)
}

func Email(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emailInnerSVG,
		children,
	)
}

func EmailOpened(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emailOpenedInnerSVG,
		children,
	)
}

func EmailOpenedTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emailOpenedTwotoneInnerSVG,
		children,
	)
}

func EmailOpenedTwotoneAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emailOpenedTwotoneAltInnerSVG,
		children,
	)
}

func EmailTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emailTwotoneInnerSVG,
		children,
	)
}

func EmailTwotoneAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emailTwotoneAltInnerSVG,
		children,
	)
}

func EmojiAngry(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiAngryInnerSVG,
		children,
	)
}

func EmojiAngryTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiAngryTwotoneInnerSVG,
		children,
	)
}

func EmojiFrown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiFrownInnerSVG,
		children,
	)
}

func EmojiFrownOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiFrownOpenInnerSVG,
		children,
	)
}

func EmojiFrownOpenTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiFrownOpenTwotoneInnerSVG,
		children,
	)
}

func EmojiFrownTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiFrownTwotoneInnerSVG,
		children,
	)
}

func EmojiGrin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiGrinInnerSVG,
		children,
	)
}

func EmojiGrinTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiGrinTwotoneInnerSVG,
		children,
	)
}

func EmojiNeutral(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiNeutralInnerSVG,
		children,
	)
}

func EmojiNeutralTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiNeutralTwotoneInnerSVG,
		children,
	)
}

func EmojiSmile(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiSmileInnerSVG,
		children,
	)
}

func EmojiSmileTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiSmileTwotoneInnerSVG,
		children,
	)
}

func EmojiSmileWink(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiSmileWinkInnerSVG,
		children,
	)
}

func EmojiSmileWinkTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emojiSmileWinkTwotoneInnerSVG,
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

func ExternalLinkRounded(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		externalLinkRoundedInnerSVG,
		children,
	)
}

func Facebook(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		facebookInnerSVG,
		children,
	)
}

func Github(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		githubInnerSVG,
		children,
	)
}

func GithubLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		githubLoopInnerSVG,
		children,
	)
}

func GithubTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		githubTwotoneInnerSVG,
		children,
	)
}

func GridThree(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gridThreeInnerSVG,
		children,
	)
}

func GridThreeFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gridThreeFilledInnerSVG,
		children,
	)
}

func GridThreeTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gridThreeTwotoneInnerSVG,
		children,
	)
}

func Hash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hashInnerSVG,
		children,
	)
}

func HashSmall(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hashSmallInnerSVG,
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

func HeartFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartFilledInnerSVG,
		children,
	)
}

func HeartFilledHalf(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartFilledHalfInnerSVG,
		children,
	)
}

func HeartHalf(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartHalfInnerSVG,
		children,
	)
}

func HeartHalfFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartHalfFilledInnerSVG,
		children,
	)
}

func HeartHalfTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartHalfTwotoneInnerSVG,
		children,
	)
}

func HeartTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartTwotoneInnerSVG,
		children,
	)
}

func HeartTwotoneHalf(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartTwotoneHalfInnerSVG,
		children,
	)
}

func HeartTwotoneHalfFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartTwotoneHalfFilledInnerSVG,
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

func HomeMd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeMdInnerSVG,
		children,
	)
}

func HomeMdTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeMdTwotoneInnerSVG,
		children,
	)
}

func HomeMdTwotoneAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeMdTwotoneAltInnerSVG,
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

func HomeSimpleFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeSimpleFilledInnerSVG,
		children,
	)
}

func HomeSimpleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeSimpleTwotoneInnerSVG,
		children,
	)
}

func HomeTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeTwotoneInnerSVG,
		children,
	)
}

func HomeTwotoneAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeTwotoneAltInnerSVG,
		children,
	)
}

func IconifyOne(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		iconifyOneInnerSVG,
		children,
	)
}

func IconifyTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		iconifyTwoInnerSVG,
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

func ImageTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageTwotoneInnerSVG,
		children,
	)
}

func Instagram(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		instagramInnerSVG,
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

func LaptopTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		laptopTwotoneInnerSVG,
		children,
	)
}

func LightDark(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightDarkInnerSVG,
		children,
	)
}

func LightDarkLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightDarkLoopInnerSVG,
		children,
	)
}

func Lightbulb(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbInnerSVG,
		children,
	)
}

func LightbulbFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbFilledInnerSVG,
		children,
	)
}

func LightbulbOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbOffInnerSVG,
		children,
	)
}

func LightbulbOffFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbOffFilledInnerSVG,
		children,
	)
}

func LightbulbOffFilledLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbOffFilledLoopInnerSVG,
		children,
	)
}

func LightbulbOffLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbOffLoopInnerSVG,
		children,
	)
}

func LightbulbOffTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbOffTwotoneInnerSVG,
		children,
	)
}

func LightbulbOffTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbOffTwotoneLoopInnerSVG,
		children,
	)
}

func LightbulbTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lightbulbTwotoneInnerSVG,
		children,
	)
}

func Linkedin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkedinInnerSVG,
		children,
	)
}

func List(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listInnerSVG,
		children,
	)
}

func ListIndented(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listIndentedInnerSVG,
		children,
	)
}

func ListThree(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listThreeInnerSVG,
		children,
	)
}

func ListThreeFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listThreeFilledInnerSVG,
		children,
	)
}

func ListThreeTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listThreeTwotoneInnerSVG,
		children,
	)
}

func LoadingAltLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loadingAltLoopInnerSVG,
		children,
	)
}

func LoadingLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loadingLoopInnerSVG,
		children,
	)
}

func LoadingTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loadingTwotoneLoopInnerSVG,
		children,
	)
}

func Marker(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		markerInnerSVG,
		children,
	)
}

func MarkerFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		markerFilledInnerSVG,
		children,
	)
}

func MarkerTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		markerTwotoneInnerSVG,
		children,
	)
}

func Mastodon(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mastodonInnerSVG,
		children,
	)
}

func MastodonFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mastodonFilledInnerSVG,
		children,
	)
}

func MastodonTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mastodonTwotoneInnerSVG,
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

func MenuFoldLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuFoldLeftInnerSVG,
		children,
	)
}

func MenuFoldRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuFoldRightInnerSVG,
		children,
	)
}

func MenuToCloseAltTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuToCloseAltTransitionInnerSVG,
		children,
	)
}

func MenuToCloseTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuToCloseTransitionInnerSVG,
		children,
	)
}

func MenuUnfoldLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuUnfoldLeftInnerSVG,
		children,
	)
}

func MenuUnfoldRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuUnfoldRightInnerSVG,
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

func MinusCircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusCircleTwotoneInnerSVG,
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

func MoonAltLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonAltLoopInnerSVG,
		children,
	)
}

func MoonAltToSunnyOutlineLoopTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonAltToSunnyOutlineLoopTransitionInnerSVG,
		children,
	)
}

func MoonFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonFilledInnerSVG,
		children,
	)
}

func MoonFilledAltLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonFilledAltLoopInnerSVG,
		children,
	)
}

func MoonFilledAltToSunnyFilledLoopTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonFilledAltToSunnyFilledLoopTransitionInnerSVG,
		children,
	)
}

func MoonFilledLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonFilledLoopInnerSVG,
		children,
	)
}

func MoonFilledToSunnyFilledLoopTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonFilledToSunnyFilledLoopTransitionInnerSVG,
		children,
	)
}

func MoonFilledToSunnyFilledTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonFilledToSunnyFilledTransitionInnerSVG,
		children,
	)
}

func MoonLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonLoopInnerSVG,
		children,
	)
}

func MoonRisingAltLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonRisingAltLoopInnerSVG,
		children,
	)
}

func MoonRisingFilledAltLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonRisingFilledAltLoopInnerSVG,
		children,
	)
}

func MoonRisingFilledLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonRisingFilledLoopInnerSVG,
		children,
	)
}

func MoonRisingLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonRisingLoopInnerSVG,
		children,
	)
}

func MoonRisingTwotoneAltLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonRisingTwotoneAltLoopInnerSVG,
		children,
	)
}

func MoonRisingTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonRisingTwotoneLoopInnerSVG,
		children,
	)
}

func MoonToSunnyOutlineLoopTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonToSunnyOutlineLoopTransitionInnerSVG,
		children,
	)
}

func MoonToSunnyOutlineTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonToSunnyOutlineTransitionInnerSVG,
		children,
	)
}

func MoonTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonTwotoneInnerSVG,
		children,
	)
}

func MoonTwotoneAltLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonTwotoneAltLoopInnerSVG,
		children,
	)
}

func MoonTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonTwotoneLoopInnerSVG,
		children,
	)
}

func NavigationLeftDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		navigationLeftDownInnerSVG,
		children,
	)
}

func NavigationLeftUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		navigationLeftUpInnerSVG,
		children,
	)
}

func NavigationRightDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		navigationRightDownInnerSVG,
		children,
	)
}

func NavigationRightUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		navigationRightUpInnerSVG,
		children,
	)
}

func PaintDrop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paintDropInnerSVG,
		children,
	)
}

func PaintDropFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paintDropFilledInnerSVG,
		children,
	)
}

func PaintDropHalfFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paintDropHalfFilledInnerSVG,
		children,
	)
}

func PaintDropHalfFilledTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paintDropHalfFilledTwotoneInnerSVG,
		children,
	)
}

func PaintDropHalfTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paintDropHalfTwotoneInnerSVG,
		children,
	)
}

func PaintDropTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paintDropTwotoneInnerSVG,
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

func PauseToPlayFilledTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pauseToPlayFilledTransitionInnerSVG,
		children,
	)
}

func PauseToPlayTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pauseToPlayTransitionInnerSVG,
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

func PencilTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pencilTwotoneInnerSVG,
		children,
	)
}

func PencilTwotoneAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pencilTwotoneAltInnerSVG,
		children,
	)
}

func Play(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playInnerSVG,
		children,
	)
}

func PlayFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playFilledInnerSVG,
		children,
	)
}

func PlayFilledToPauseTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playFilledToPauseTransitionInnerSVG,
		children,
	)
}

func PlayToPauseTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playToPauseTransitionInnerSVG,
		children,
	)
}

func PlayTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playTwotoneInnerSVG,
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

func PlusCircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusCircleTwotoneInnerSVG,
		children,
	)
}

func Question(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionInnerSVG,
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

func QuestionCircleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionCircleTwotoneInnerSVG,
		children,
	)
}

func Reddit(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		redditInnerSVG,
		children,
	)
}

func RedditCircle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		redditCircleInnerSVG,
		children,
	)
}

func RedditCircleLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		redditCircleLoopInnerSVG,
		children,
	)
}

func RedditLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		redditLoopInnerSVG,
		children,
	)
}

func Remove(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeInnerSVG,
		children,
	)
}

func RotateNinety(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rotateNinetyInnerSVG,
		children,
	)
}

func RotateOneHundredEighty(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rotateOneHundredEightyInnerSVG,
		children,
	)
}

func RotateTwoHundredSeventy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rotateTwoHundredSeventyInnerSVG,
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

func SearchFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchFilledInnerSVG,
		children,
	)
}

func SearchTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchTwotoneInnerSVG,
		children,
	)
}

func SunRisingFilledLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunRisingFilledLoopInnerSVG,
		children,
	)
}

func SunRisingLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunRisingLoopInnerSVG,
		children,
	)
}

func SunRisingTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunRisingTwotoneLoopInnerSVG,
		children,
	)
}

func SunnyFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyFilledInnerSVG,
		children,
	)
}

func SunnyFilledLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyFilledLoopInnerSVG,
		children,
	)
}

func SunnyFilledLoopToMoonAltFilledLoopTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyFilledLoopToMoonAltFilledLoopTransitionInnerSVG,
		children,
	)
}

func SunnyFilledLoopToMoonFilledLoopTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyFilledLoopToMoonFilledLoopTransitionInnerSVG,
		children,
	)
}

func SunnyFilledLoopToMoonFilledTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyFilledLoopToMoonFilledTransitionInnerSVG,
		children,
	)
}

func SunnyOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyOutlineInnerSVG,
		children,
	)
}

func SunnyOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyOutlineLoopInnerSVG,
		children,
	)
}

func SunnyOutlineToMoonAltLoopTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyOutlineToMoonAltLoopTransitionInnerSVG,
		children,
	)
}

func SunnyOutlineToMoonLoopTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyOutlineToMoonLoopTransitionInnerSVG,
		children,
	)
}

func SunnyOutlineToMoonTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyOutlineToMoonTransitionInnerSVG,
		children,
	)
}

func SunnyOutlineTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyOutlineTwotoneInnerSVG,
		children,
	)
}

func SunnyOutlineTwotoneLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunnyOutlineTwotoneLoopInnerSVG,
		children,
	)
}

func Switch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		switchInnerSVG,
		children,
	)
}

func SwitchFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		switchFilledInnerSVG,
		children,
	)
}

func SwitchFilledToSwitchOffFilledTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		switchFilledToSwitchOffFilledTransitionInnerSVG,
		children,
	)
}

func SwitchOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		switchOffInnerSVG,
		children,
	)
}

func SwitchOffFilled(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		switchOffFilledInnerSVG,
		children,
	)
}

func SwitchOffFilledToSwitchFilledTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		switchOffFilledToSwitchFilledTransitionInnerSVG,
		children,
	)
}

func SwitchOffToSwitchTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		switchOffToSwitchTransitionInnerSVG,
		children,
	)
}

func SwitchToSwitchOffTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		switchToSwitchOffTransitionInnerSVG,
		children,
	)
}

func Telegram(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		telegramInnerSVG,
		children,
	)
}

func TextBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textBoxInnerSVG,
		children,
	)
}

func TextBoxMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textBoxMultipleInnerSVG,
		children,
	)
}

func TextBoxMultipleToTextBoxTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textBoxMultipleToTextBoxTransitionInnerSVG,
		children,
	)
}

func TextBoxMultipleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textBoxMultipleTwotoneInnerSVG,
		children,
	)
}

func TextBoxMultipleTwotoneToTextBoxTwotoneTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textBoxMultipleTwotoneToTextBoxTwotoneTransitionInnerSVG,
		children,
	)
}

func TextBoxToTextBoxMultipleTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textBoxToTextBoxMultipleTransitionInnerSVG,
		children,
	)
}

func TextBoxTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textBoxTwotoneInnerSVG,
		children,
	)
}

func TextBoxTwotoneToTextBoxMultipleTwotoneTransition(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textBoxTwotoneToTextBoxMultipleTwotoneTransitionInnerSVG,
		children,
	)
}

func ThumbsDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thumbsDownInnerSVG,
		children,
	)
}

func ThumbsDownTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thumbsDownTwotoneInnerSVG,
		children,
	)
}

func ThumbsUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thumbsUpInnerSVG,
		children,
	)
}

func ThumbsUpTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thumbsUpTwotoneInnerSVG,
		children,
	)
}

func Twitter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		twitterInnerSVG,
		children,
	)
}

func TwitterTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		twitterTwotoneInnerSVG,
		children,
	)
}

func UploadLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadLoopInnerSVG,
		children,
	)
}

func UploadOffLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadOffLoopInnerSVG,
		children,
	)
}

func UploadOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadOffOutlineInnerSVG,
		children,
	)
}

func UploadOffOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadOffOutlineLoopInnerSVG,
		children,
	)
}

func UploadOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadOutlineInnerSVG,
		children,
	)
}

func UploadOutlineLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadOutlineLoopInnerSVG,
		children,
	)
}

func UploadingLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadingLoopInnerSVG,
		children,
	)
}

func ValignBaseline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		valignBaselineInnerSVG,
		children,
	)
}

func ValignBaselineTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		valignBaselineTwotoneInnerSVG,
		children,
	)
}

func ValignBottom(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		valignBottomInnerSVG,
		children,
	)
}

func ValignBottomTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		valignBottomTwotoneInnerSVG,
		children,
	)
}

func ValignMiddle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		valignMiddleInnerSVG,
		children,
	)
}

func ValignMiddleTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		valignMiddleTwotoneInnerSVG,
		children,
	)
}

func ValignTop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		valignTopInnerSVG,
		children,
	)
}

func ValignTopTwotone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		valignTopTwotoneInnerSVG,
		children,
	)
}

func WeatherCloudyLoop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		weatherCloudyLoopInnerSVG,
		children,
	)
}
