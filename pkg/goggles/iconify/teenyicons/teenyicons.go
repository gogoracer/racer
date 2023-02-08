package teenyicons

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	abTestingOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M1.5 11V6.5a2 2 0 1 1 4 0V11m-4-2.5h4m6.5-1H9.5m2.5 0a1.5 1.5 0 0 0 0-3H9.5v3m2.5 0a1.5 1.5 0 0 1 0 3H9.5v-3M7.5 1v13"/>`
	abTestingSolidInnerSVG                   = `<path fill="currentColor" d="M3.5 5A1.5 1.5 0 0 0 2 6.5V8h3V6.5A1.5 1.5 0 0 0 3.5 5Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4A1.5 1.5 0 0 1 7 2.5v10A1.5 1.5 0 0 1 5.5 14h-4A1.5 1.5 0 0 1 0 12.5v-10ZM2 11V9h3v2h1V6.5a2.5 2.5 0 0 0-5 0V11h1Z" clip-rule="evenodd"/><path fill="currentColor" d="M12 7h-2V5h2a1 1 0 1 1 0 2Zm0 3h-2V8h2a1 1 0 1 1 0 2Z"/><path fill="currentColor" fill-rule="evenodd" d="M8 2.5A1.5 1.5 0 0 1 9.5 1h4A1.5 1.5 0 0 1 15 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-4A1.5 1.5 0 0 1 8 12.5v-10ZM12 4H9v7h3a2 2 0 0 0 1.323-3.5A2 2 0 0 0 12 4Z" clip-rule="evenodd"/>`
	addOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M7.5 1v13M1 7.5h13"/>`
	addSmallOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M7.5 4v7M4 7.5h7"/>`
	addSmallSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7 7V4h1v3h3v1H8v3H7V8H4V7h3Z" clip-rule="evenodd"/>`
	addSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M7 7V1h1v6h6v1H8v6H7V8H1V7h6Z" clip-rule="evenodd"/>`
	addressBookOutlineInnerSVG               = `<path fill="currentColor" d="M5.5 11.5H5v.5h.5v-.5Zm5 0v.5h.5v-.5h-.5Zm-4.5 0V11H5v.5h1Zm4-.5v.5h1V11h-1Zm.5 0h-5v1h5v-1ZM8 9a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3v1Zm-2 2a2 2 0 0 1 2-2V8a3 3 0 0 0-3 3h1Zm2-8a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1V3Zm2 2a2 2 0 0 0-2-2v1a1 1 0 0 1 1 1h1ZM8 7a2 2 0 0 0 2-2H9a1 1 0 0 1-1 1v1Zm0-1a1 1 0 0 1-1-1H6a2 2 0 0 0 2 2V6ZM3.5 1h9V0h-9v1Zm9.5.5v12h1v-12h-1ZM12.5 14h-9v1h9v-1ZM3 13.5v-12H2v12h1Zm.5.5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 15v-1Zm9.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM12.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 12.5 0v1Zm-9-1A1.5 1.5 0 0 0 2 1.5h1a.5.5 0 0 1 .5-.5V0ZM4 4H1v1h3V4Zm0 6H1v1h3v-1Z"/>`
	addressBookSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 0A1.5 1.5 0 0 0 2 1.5V4H1v1h1v5H1v1h1v2.5A1.5 1.5 0 0 0 3.5 15h9a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 12.5 0h-9ZM6 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm-1 6a3 3 0 1 1 6 0v1H5v-1Z" clip-rule="evenodd"/>`
	adjustHorizontalAltOutlineInnerSVG       = `<path fill="none" stroke="currentColor" d="M15 3.5H6.5m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H0m15 8h-2.5m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H0"/>`
	adjustHorizontalAltSolidInnerSVG         = `<path fill="currentColor" d="M4.5 1a2.5 2.5 0 0 0-2.45 2H0v1h2.05a2.5 2.5 0 0 0 4.9 0H15V3H6.95A2.5 2.5 0 0 0 4.5 1Zm6 8a2.5 2.5 0 0 0-2.45 2H0v1h8.05a2.5 2.5 0 0 0 4.9 0H15v-1h-2.05a2.5 2.5 0 0 0-2.45-2Z"/>`
	adjustHorizontalOutlineInnerSVG          = `<path fill="none" stroke="currentColor" d="M2.5 7.5H0m15 5h-2.5m2.5-10H8.5m-2 0H0m4.5 5H15m-4.5 5H0m10.5-2v4h2v-4h-2Zm-8-5v4h2v-4h-2Zm4-5v4h2v-4h-2Z"/>`
	adjustHorizontalSolidInnerSVG            = `<path fill="currentColor" d="M9 0H6v2H0v1h6v2h3V3h6V2H9V0ZM5 5H2v2H0v1h2v2h3V8h10V7H5V5Zm8 5h-3v2H0v1h10v2h3v-2h2v-1h-2v-2Z"/>`
	adjustVerticalAltOutlineInnerSVG         = `<path fill="none" stroke="currentColor" d="M3.5 0v8.5m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0V15m8-15v2.5m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0V15"/>`
	adjustVerticalAltSolidInnerSVG           = `<path fill="currentColor" d="M3 0v8.05a2.5 2.5 0 0 0 0 4.9V15h1v-2.05a2.5 2.5 0 0 0 0-4.9V0H3Zm8 0v2.05a2.5 2.5 0 0 0 0 4.9V15h1V6.95a2.5 2.5 0 0 0 0-4.9V0h-1Z"/>`
	adjustVerticalOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M7.5 12.5V15m5-15v2.5M2.5 0v6.5m0 2V15m5-4.5V0m5 4.5V15m-2-10.5h4v-2h-4v2Zm-5 8h4v-2h-4v2Zm-5-4h4v-2h-4v2Z"/>`
	adjustVerticalSolidInnerSVG              = `<path fill="currentColor" d="M2 0v6H0v3h2v6h1V9h2V6H3V0H2Zm3 10h2V0h1v10h2v3H8v2H7v-2H5v-3Zm7-10v2h-2v3h2v10h1V5h2V2h-2V0h-1Z"/>`
	airplayOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M3 11.5H.5v-10h14v10H12m-4.5-2l-4 4h8l-4-4Z"/>`
	airplaySolidInnerSVG                     = `<path fill="currentColor" d="M.5 1a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5H3v-1H1V2h13v9h-2v1h2.5a.5.5 0 0 0 .5-.5v-10a.5.5 0 0 0-.5-.5H.5Z"/><path fill="currentColor" d="M7.854 9.146a.5.5 0 0 0-.708 0l-4 4A.5.5 0 0 0 3.5 14h8a.5.5 0 0 0 .354-.854l-4-4Z"/>`
	airpodsOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M6.5 3.5a2.648 2.648 0 0 1-2.977 2.628l-.32-.04a2.667 2.667 0 0 1-1.27-.513L.5 4.5v-2l1.433-1.075a2.667 2.667 0 0 1 1.27-.513l.32-.04A2.648 2.648 0 0 1 6.5 3.5Zm0 0v11h-2V6m4-2.5a2.648 2.648 0 0 0 2.977 2.628l.32-.04c.46-.058.898-.234 1.27-.513L14.5 4.5v-2l-1.433-1.075a2.667 2.667 0 0 0-1.27-.513l-.32-.04A2.648 2.648 0 0 0 8.5 3.5Zm0 0v11h2V6M2 3.5h2m7 0h2"/>`
	airpodsSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7 3.5A3.148 3.148 0 0 0 3.461.376l-.32.04a3.167 3.167 0 0 0-1.508.609L0 2.25v2.5l1.633 1.225c.441.33.96.54 1.508.609l.32.04c.182.023.362.03.539.022V15h3V3.5ZM4 4H2V3h2v1Zm4-.5A3.148 3.148 0 0 1 11.539.376l.32.04a3.167 3.167 0 0 1 1.508.609L15 2.25v2.5l-1.633 1.225c-.441.33-.96.54-1.508.609l-.32.04a3.18 3.18 0 0 1-.539.022V15H8V3.5Zm3 .5h2V3h-2v1Z" clip-rule="evenodd"/>`
	alarmOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="m.5 3.5l3-3m8 0l3 3M7.5 5v3.5H10m-2.5-6a6 6 0 1 0 0 12a6 6 0 0 0 0-12Z"/>`
	alarmSolidInnerSVG                       = `<path fill="currentColor" d="m3.854.854l-3 3l-.708-.708l3-3l.708.708Zm10.292 3l-3-3l.708-.708l3 3l-.707.708Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 8.5a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0ZM8 8V5H7v4h3V8H8Z" clip-rule="evenodd"/>`
	alienOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" d="M6.52 11.435c.24.107.719.497.981.497c.263 0 .741-.39.982-.497m-3.926-4.97l1.472 1.49m4.417-1.49l-1.472 1.49M7.5.5C3.94.5 1.967 2.45 1.612 4.974c-.358 2.33.136 4.53 1.472 6.461c.643.953 1.486 1.876 2.454 2.486c1.271.772 2.654.772 3.925 0c.967-.61 1.81-1.533 2.454-2.486c1.33-1.934 1.824-4.13 1.472-6.461C13.034 2.449 11.062.5 7.501.5Z"/>`
	alienSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M3.039 1.403C4.127.513 5.629 0 7.501 0c1.873 0 3.374.514 4.463 1.403c1.089.89 1.727 2.125 1.921 3.498c.37 2.453-.151 4.776-1.554 6.816c-.668.99-1.558 1.97-2.6 2.627l-.007.004c-1.431.87-3.014.87-4.445 0l-.007-.004c-1.042-.657-1.931-1.637-2.6-2.627C1.264 9.68.742 7.354 1.118 4.901c.194-1.373.832-2.608 1.921-3.498ZM6.736 7.96l-.711.703L3.85 6.46l.712-.702L6.736 7.96Zm1.53 0l.712.703l2.175-2.203l-.712-.702L8.267 7.96Zm-2.001 2.816l.457.202c.097.043.2.105.28.156a15.37 15.37 0 0 1 .26.163a2.036 2.036 0 0 0 .24.127a2.036 2.036 0 0 0 .24-.128A8.542 8.542 0 0 0 8 11.134c.082-.051.184-.113.28-.156l.458-.202l.404.915l-.457.202c-.024.01-.071.037-.156.09l-.086.054l-.183.114c-.097.059-.21.124-.324.175a1.084 1.084 0 0 1-.435.106c-.173 0-.33-.06-.434-.106a3.003 3.003 0 0 1-.324-.175c-.06-.036-.125-.078-.183-.114l-.087-.055a1.543 1.543 0 0 0-.155-.09l-.457-.201l.404-.915Z" clip-rule="evenodd"/>`
	alignBottomOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M15 14.5H0m11.5-3h-3v-7h3v7Zm-5 0h-3V.5h3v11Z"/>`
	alignBottomSolidInnerSVG                 = `<path fill="currentColor" d="M7 0H3v12h4V0Zm5 4H8v8h4V4Zm3 10H0v1h15v-1Z"/>`
	alignCenterHorizontalOutlineInnerSVG     = `<path fill="none" stroke="currentColor" d="M7.5 0v3.5m0 8V15m0-8.5v2M4 3.5v3h7v-3H4Zm-2.5 5v3h12v-3h-12Z"/>`
	alignCenterHorizontalSolidInnerSVG       = `<path fill="currentColor" d="M7 0v3H3.5v4H7v1H1v4h6v3h1v-3h6V8H8V7h3.5V3H8V0H7Z"/>`
	alignCenterVerticalOutlineInnerSVG       = `<path fill="none" stroke="currentColor" d="M15 7.5h-3.5m-8 0H0m8.5 0h-2m5 3.5h-3V4h3v7Zm-5 2.5h-3v-12h3v12Z"/>`
	alignCenterVerticalSolidInnerSVG         = `<path fill="currentColor" d="M7 1H3v6H0v1h3v6h4V8h1v3.5h4V8h3V7h-3V3.5H8V7H7V1Z"/>`
	alignLeftOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M.5 0v15m3-11.5v3h7v-3h-7Zm0 5v3h11v-3h-11Z"/>`
	alignLeftSolidInnerSVG                   = `<path fill="currentColor" d="M0 0v15h1V0H0Zm11 3H3v4h8V3Zm4 5H3v4h12V8Z"/>`
	alignRightOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M14.5 0v15m-3-11.5v3h-7v-3h7Zm0 5v3H.5v-3h11Z"/>`
	alignRightSolidInnerSVG                  = `<path fill="currentColor" d="M14 0v15h1V0h-1Zm-2 3H4v4h8V3Zm0 5H0v4h12V8Z"/>`
	alignTextCenterOutlineInnerSVG           = `<path fill="none" stroke="currentColor" d="M15 3.5H0m10 4H5m7 4H3"/>`
	alignTextCenterSolidInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M0 3h15v1H0V3Zm5 4h5v1H5V7Zm-2 4h9v1H3v-1Z" clip-rule="evenodd"/>`
	alignTextJustifyOutlineInnerSVG          = `<path fill="none" stroke="currentColor" d="M0 3.5h15m-15 8h15m-15-4h15"/>`
	alignTextJustifySolidInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M15 4H0V3h15v1Zm0 4H0V7h15v1Zm0 4H0v-1h15v1Z" clip-rule="evenodd"/>`
	alignTextLeftOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M0 3.5h15m-15 8h9m-9-4h6"/>`
	alignTextLeftSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M15 4H0V3h15v1ZM6 8H0V7h6v1Zm3 4H0v-1h9v1Z" clip-rule="evenodd"/>`
	alignTextRightOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M15 3.5H0m15 8H6m9-4H9"/>`
	alignTextRightSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M0 3h15v1H0V3Zm9 4h6v1H9V7Zm-3 4h9v1H6v-1Z" clip-rule="evenodd"/>`
	alignTopOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M15 .5H0m11.5 3h-3v7h3v-7Zm-5 0h-3v11h3v-11Z"/>`
	alignTopSolidInnerSVG                    = `<path fill="currentColor" d="M15 0H0v1h15V0ZM7 3H3v12h4V3Zm5 0H8v8h4V3Z"/>`
	anchorOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5 4.497a2 2 0 0 0 2-1.998a2 2 0 0 0-4 0a2 2 0 0 0 2 1.998Zm0 0v9.994m0 0c-3.866 0-7-3.132-7-6.996h2m5 6.996c3.866 0 7-3.132 7-6.996h-2"/>`
	anchorSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 1A1.5 1.5 0 0 0 6 2.499a1.5 1.5 0 0 0 3 0A1.5 1.5 0 0 0 7.5 1ZM5 2.499a2.5 2.5 0 1 1 3 2.448v9.025a6.499 6.499 0 0 0 5.981-5.977H12v-1h3v.5a7.498 7.498 0 0 1-7.5 7.496A7.498 7.498 0 0 1 0 7.495v-.5h3v1H1.019A6.499 6.499 0 0 0 7 13.972V4.947A2.5 2.5 0 0 1 5 2.5Z" clip-rule="evenodd"/>`
	androidOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="m2.5 2.5l2 3m8-3l-2 3M4 9.5h1m5 0h1m-9.5 3v-2a6 6 0 1 1 12 0v2h-12Z"/>`
	androidSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 4a6.473 6.473 0 0 0-2.934.698l-1.65-2.475l-.832.554l1.627 2.44A6.492 6.492 0 0 0 1 10.5V13h13v-2.5a6.492 6.492 0 0 0-2.711-5.282l1.627-2.44l-.832-.555l-1.65 2.475A6.473 6.473 0 0 0 7.5 4ZM5 10H4V9h1v1Zm5 0h1V9h-1v1Z" clip-rule="evenodd"/>`
	angularOutlineInnerSVG                   = `<path fill="currentColor" d="m7.5.5l.137-.48a.5.5 0 0 0-.274 0L7.5.5Zm-7 2l-.137-.48a.5.5 0 0 0-.36.535L.5 2.5Zm1 9l-.497.055a.5.5 0 0 0 .273.392L1.5 11.5Zm6 3l-.224.447a.5.5 0 0 0 .448 0L7.5 14.5Zm6-3l.224.447a.5.5 0 0 0 .273-.392L13.5 11.5Zm1-9l.497.055a.5.5 0 0 0-.36-.536L14.5 2.5Zm-7 .5l.458-.2L7.5 1.753L7.042 2.8L7.5 3ZM7.363.02l-7 2l.274.96l7-2l-.274-.96ZM.003 2.554l1 9l.994-.11l-1-9l-.994.11Zm1.273 9.392l6 3l.448-.894l-6-3l-.448.894Zm6.448 3l6-3l-.448-.894l-6 3l.448.894Zm6.273-3.392l1-9l-.994-.11l-1 9l.994.11Zm.64-9.536l-7-2l-.274.962l7 2l.274-.962ZM4.458 11.2l3.5-8l-.916-.4l-3.5 8l.916.4Zm2.584-8l3.5 8l.916-.4l-3.5-8l-.916.4ZM5 9h5V8H5v1Z"/>`
	angularSolidInnerSVG                     = `<path fill="currentColor" d="M7.5 4.247L9.142 8H5.858L7.5 4.247Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.363.02a.5.5 0 0 1 .274 0l7 2a.5.5 0 0 1 .36.535l-1 9a.5.5 0 0 1-.273.392l-6 3a.5.5 0 0 1-.448 0l-6-3a.5.5 0 0 1-.273-.392l-1-9a.5.5 0 0 1 .36-.536l7-2ZM7.5 1.752l3.958 9.047l-.916.4L9.579 9H5.421l-.963 2.2l-.916-.4L7.5 1.753Z" clip-rule="evenodd"/>`
	anjaOutlineInnerSVG                      = `<path fill="currentColor" d="m9.007 4.4l.034-.498l-.034.498ZM2.955 8.139l-.461-.194l.46.194ZM7.5 14c-1.627 0-3.1-.958-3.762-2.444l-.913.406A5.116 5.116 0 0 0 7.5 15v-1Zm3.761-2.444A4.116 4.116 0 0 1 7.5 14v1a5.116 5.116 0 0 0 4.675-3.038l-.914-.406ZM2 15V6H1v9h1ZM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0v1ZM2 6.5A5.5 5.5 0 0 1 7.5 1V0A6.5 6.5 0 0 0 1 6.5h1Zm11 0V15h1V6.5h-1ZM8.974 4.9l1.493.099l.066-.998l-1.492-.1l-.067.998ZM11 4.5V3h-1v1.5h1ZM3.416 8.331A5.624 5.624 0 0 1 8.974 4.9l.067-.997a6.624 6.624 0 0 0-6.547 4.042l.922.387Zm6.73-3.477c.47.47 1.33 1.49 1.69 3.342l.98-.19c-.406-2.1-1.394-3.291-1.962-3.86l-.708.708ZM6.5 13H8v-1H6.5v1ZM8 13a2 2 0 0 0 2-2H9a1 1 0 0 1-1 1v1Zm4.5-4h-2v1h2V9Zm-4-1H12V7H8.5v1Zm-5 0h3V7h-3v1Zm3 0h2V7h-2v1Zm-2 1h-2v1h2V9ZM6 7.5A1.5 1.5 0 0 1 4.5 9v1A2.5 2.5 0 0 0 7 7.5H6ZM10.5 9A1.5 1.5 0 0 1 9 7.5H8a2.5 2.5 0 0 0 2.5 2.5V9Zm-6.762 2.556c-.495-1.116-.73-2.255-.322-3.225l-.922-.387c-.57 1.36-.2 2.826.33 4.018l.914-.406Zm8.437.406c.53-1.19.91-2.574.642-3.957l-.982.19c.212 1.09-.08 2.25-.574 3.36l.914.407Z"/>`
	anjaSolidInnerSVG                        = `<path fill="currentColor" d="M8 13a2 2 0 0 0 2-2H9a1 1 0 0 1-1 1H6.5v1H8Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 15a5.14 5.14 0 0 1-.543-.029L7 15H1V6h.019A6.5 6.5 0 0 1 14 6.5V15H7.5ZM8.974 4.9l1.3.086A5.951 5.951 0 0 1 11.503 7H4.207a5.616 5.616 0 0 1 4.767-2.1ZM9.085 8h2.71l.04.196a4 4 0 0 1 .07.804H10.5a1.5 1.5 0 0 1-1.415-1Zm1.415 2h1.286a8.016 8.016 0 0 1-.524 1.556a4.116 4.116 0 0 1-7.524 0A6.23 6.23 0 0 1 3.254 10H4.5a2.5 2.5 0 0 0 2.45-2h1.1a2.5 2.5 0 0 0 2.45 2Zm-6-1H3.235c.032-.23.09-.453.18-.669c.048-.112.099-.223.153-.331h2.347A1.5 1.5 0 0 1 4.5 9Z" clip-rule="evenodd"/>`
	antiClockwiseOutlineInnerSVG             = `<path fill="currentColor" d="m6.5 2.499l-.354-.354l-.353.354l.353.353L6.5 2.5Zm1-.5H7v1h.5v-1ZM2 8.495v-.5H1v.5h1ZM8.145.146l-1.999 2l.708.706L8.853.854L8.145.146ZM6.146 2.852l2 1.999l.707-.707l-2-1.999l-.707.707ZM7.5 3C10.537 3 13 5.461 13 8.496h1A6.499 6.499 0 0 0 7.5 2v1ZM13 8.495a5.499 5.499 0 0 1-5.5 5.496v1c3.589 0 6.5-2.909 6.5-6.496h-1ZM7.5 13.99A5.499 5.499 0 0 1 2 8.495H1a6.499 6.499 0 0 0 6.5 6.496v-1Z"/>`
	antiClockwiseSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="m8.145.146l.708.708l-1.149 1.148A6.499 6.499 0 0 1 14 8.495a6.499 6.499 0 0 1-6.5 6.496A6.499 6.499 0 0 1 1 8.495v-.5h1v.5a5.499 5.499 0 0 0 5.5 5.496A5.5 5.5 0 0 0 13 8.495a5.499 5.499 0 0 0-5.289-5.492l1.142 1.14l-.708.708l-2.352-2.352L8.145.146Z" clip-rule="evenodd"/>`
	appleOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M7.825 3.241a.343.343 0 0 1-.343-.342A2.399 2.399 0 0 1 9.881.5c.19 0 .342.154.342.343A2.399 2.399 0 0 1 7.825 3.24Zm5.003 7.216c.132.099.175.28.1.427c-1.205 2.414-2.168 3.616-3.047 3.616c-.409 0-.811-.132-1.203-.39a1.782 1.782 0 0 0-1.895-.041c-.474.284-.927.431-1.356.431C4.133 14.5 2 10.518 2 8.332C2 6 3.223 4.22 5.084 4.22c.875 0 1.631.13 2.266.39c.269.112.573.104.836-.022c.515-.248 1.194-.368 2.038-.368c1.03 0 1.926.513 2.672 1.508a.343.343 0 0 1-.068.48c-.833.624-1.234 1.326-1.234 2.124c0 .799.401 1.5 1.234 2.125Z"/>`
	appleSolidInnerSVG                       = `<path fill="currentColor" d="M7.875 2.937A.371.371 0 0 1 7.5 2.57C7.5 1.15 8.676 0 10.124 0a.37.37 0 0 1 .375.367c0 1.42-1.175 2.57-2.624 2.57Z"/><path fill="currentColor" d="M7.875 2.937A.371.371 0 0 1 7.5 2.57C7.5 1.15 8.676 0 10.124 0a.37.37 0 0 1 .375.367c0 1.42-1.175 2.57-2.624 2.57Zm5.475 7.731c.145.106.192.3.11.458C12.14 13.712 11.087 15 10.125 15c-.448 0-.888-.142-1.317-.418a1.985 1.985 0 0 0-2.073-.044c-.52.305-1.015.462-1.485.462c-1.415 0-3.75-4.267-3.75-6.608c0-2.5 1.339-4.406 3.375-4.406c.958 0 1.785.138 2.48.419c.294.118.627.11.914-.025c.564-.266 1.308-.394 2.23-.394c1.127 0 2.11.55 2.926 1.615a.362.362 0 0 1-.075.514c-.911.67-1.35 1.421-1.35 2.277c0 .855.439 1.607 1.35 2.276Z"/><path fill="currentColor" d="M13.35 10.668c.145.106.192.3.11.458C12.14 13.712 11.087 15 10.125 15c-.448 0-.888-.142-1.317-.418a1.985 1.985 0 0 0-2.073-.044c-.52.305-1.015.462-1.485.462c-1.415 0-3.75-4.267-3.75-6.608c0-2.5 1.339-4.406 3.375-4.406c.958 0 1.785.138 2.48.419c.294.118.627.11.914-.025c.564-.266 1.308-.394 2.23-.394c1.127 0 2.11.55 2.926 1.615a.362.362 0 0 1-.075.514c-.911.67-1.35 1.421-1.35 2.277c0 .855.439 1.607 1.35 2.276Z"/>`
	appointmentsOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5M3 7.5h3m6 0H9m-6 3h3m3 0h3m-10.5-8h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`
	appointmentsSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M12 2h1.5A1.5 1.5 0 0 1 15 3.5v10a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-10A1.5 1.5 0 0 1 1.5 2H3V0h1v2h7V0h1v2ZM6 8H3V7h3v1Zm6-1H9v1h3V7Zm-6 4H3v-1h3v1Zm3 0h3v-1H9v1Z" clip-rule="evenodd"/>`
	archiveOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M5 8.5h5M.5.5h14v4H.5v-4Zm1 4v9a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-9h-12Z"/>`
	archiveSolidInnerSVG                     = `<path fill="currentColor" d="M0 0h15v5H0V0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6v7.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V6H1Zm9 3H5V8h5v1Z" clip-rule="evenodd"/>`
	areaChartAltOutlineInnerSVG              = `<path fill="currentColor" d="M.5 14.5H0v.5h.5v-.5Zm1.5-3v.5h1v-.5H2Zm13-9V2h-1v.5h1ZM0 0v14.5h1V0H0Zm.5 15H15v-1H.5v1ZM3 11.5c0-1.454.244-2.88.707-3.922C4.177 6.52 4.798 6 5.5 6V5c-1.298 0-2.178.98-2.707 2.172C2.256 8.38 2 9.954 2 11.5h1ZM5.5 6c.32 0 .642.158 1.005.492c.366.338.713.798 1.095 1.308c.368.49.77 1.03 1.217 1.442c.45.416 1.004.758 1.683.758V9c-.32 0-.642-.158-1.005-.492C9.13 8.17 8.782 7.71 8.4 7.2c-.368-.49-.77-1.03-1.217-1.442C6.733 5.342 6.179 5 5.5 5v1Zm5 4c1.223 0 2.363-.763 3.173-2.045C14.485 6.668 15 4.819 15 2.5h-1c0 2.18-.485 3.832-1.173 4.92C12.137 8.514 11.277 9 10.5 9v1Z"/>`
	areaChartAltSolidInnerSVG                = `<path fill="currentColor" d="M1 0H0v15h15V2.5a.5.5 0 1 0-1 0c0 2.18-.485 3.832-1.173 4.92C12.137 8.514 11.277 9 10.5 9c-.32 0-.642-.158-1.005-.492C9.13 8.17 8.782 7.71 8.4 7.2l-.016-.021c-.363-.485-.761-1.015-1.201-1.421C6.733 5.342 6.179 5 5.5 5c-1.298 0-2.178.98-2.707 2.172C2.256 8.38 2 9.954 2 11.5v.197c0 .842-.37 1.636-1 2.177V0Z"/>`
	areaChartOutlineInnerSVG                 = `<path fill="currentColor" d="M.5 14.5H0v.5h.5v-.5Zm6-9l.4-.3a.5.5 0 0 0-.816.023L6.5 5.5Zm3 4l-.4.3a.5.5 0 0 0 .807-.01L9.5 9.5ZM0 0v14.5h1V0H0Zm.5 15H15v-1H.5v1Zm2.416-3.223l4-6l-.832-.554l-4 6l.832.554ZM6.1 5.8l3 4l.8-.6l-3-4l-.8.6Zm3.807 3.99l5-7l-.814-.58l-5 7l.814.58Z"/>`
	areaChartSolidInnerSVG                   = `<path fill="currentColor" d="M1 0H0v15h15V2.5a.5.5 0 0 0-.907-.29L9.49 8.653L6.9 5.2a.5.5 0 0 0-.816.023L1 12.849V0Z"/>`
	arrowDownCircleOutlineInnerSVG           = `<path fill="currentColor" d="M5.854 8.146L5.5 7.793l-.707.707l.353.354l.708-.708ZM7.5 10.5l-.354.354l.354.353l.354-.353L7.5 10.5Zm2.354-1.646l.353-.354l-.707-.707l-.354.353l.708.708ZM.5 7.5H0h.5Zm7-7V0v.5Zm0 14V14v.5Zm7-7H14h.5ZM5.146 8.854l2 2l.708-.708l-2-2l-.708.708Zm2.708 2l2-2l-.708-.708l-2 2l.708.708ZM8 10.5V4H7v6.5h1Zm-7-3A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1Zm1 0A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1Z"/>`
	arrowDownCircleSolidInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15Zm2.707 8.5L7.5 11.207L4.793 8.5l.707-.707l1.5 1.5V4h1v5.293l1.5-1.5l.707.707Z" clip-rule="evenodd"/>`
	arrowDownOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="m7.5 13.5l4-4m-4 4l-4-4m4 4V1"/>`
	arrowDownSmallOutlineInnerSVG            = `<path fill="currentColor" d="m9.854 8.854l.353-.354l-.707-.707l-.354.353l.708.708ZM7.5 10.5l-.354.354l.354.353l.354-.353L7.5 10.5ZM5.854 8.146L5.5 7.793l-.707.707l.353.354l.708-.708Zm3.292 0l-2 2l.708.708l2-2l-.708-.708Zm-1.292 2l-2-2l-.708.708l2 2l.708-.708ZM8 10.5V4H7v6.5h1Z"/>`
	arrowDownSmallSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M8 4v5.293l1.5-1.5l.707.707L7.5 11.207L4.793 8.5l.707-.707l1.5 1.5V4h1Z" clip-rule="evenodd"/>`
	arrowDownSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M8 1v11.293l3.146-3.147l.708.708L7.5 14.207L3.146 9.854l.708-.708L7 12.293V1h1Z" clip-rule="evenodd"/>`
	arrowLeftCircleOutlineInnerSVG           = `<path fill="currentColor" d="m6.854 5.854l.353-.354l-.707-.707l-.354.353l.708.708ZM4.5 7.5l-.354-.354l-.353.354l.353.354L4.5 7.5Zm1.646 2.354l.354.353l.707-.707l-.353-.354l-.708.708ZM7.5.5V0v.5Zm7 7h.5h-.5Zm-14 0H1H.5Zm7 7V14v.5ZM6.146 5.146l-2 2l.708.708l2-2l-.708-.708Zm-2 2.708l2 2l.708-.708l-2-2l-.708.708ZM4.5 8H11V7H4.5v1Zm3-7A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1Zm0 1A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1Z"/>`
	arrowLeftCircleSolidInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-8.5 2.707L3.793 7.5L6.5 4.793l.707.707l-1.5 1.5H11v1H5.707l1.5 1.5l-.707.707Z" clip-rule="evenodd"/>`
	arrowLeftOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="m1.5 7.5l4-4m-4 4l4 4m-4-4H14"/>`
	arrowLeftSmallOutlineInnerSVG            = `<path fill="currentColor" d="m6.146 9.854l.354.353l.707-.707l-.353-.354l-.708.708ZM4.5 7.5l-.354-.354l-.353.354l.353.354L4.5 7.5Zm2.354-1.646l.353-.354l-.707-.707l-.354.353l.708.708Zm0 3.292l-2-2l-.708.708l2 2l.708-.708Zm-2-1.292l2-2l-.708-.708l-2 2l.708.708ZM4.5 8H11V7H4.5v1Z"/>`
	arrowLeftSmallSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M7.207 5.5L5.707 7H11v1H5.707l1.5 1.5l-.707.707L3.793 7.5L6.5 4.793l.707.707Z" clip-rule="evenodd"/>`
	arrowLeftSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="m2.707 8l3.147 3.146l-.708.707L.793 7.5l4.353-4.354l.708.708L2.707 7H14v1H2.707Z" clip-rule="evenodd"/>`
	arrowOutlineInnerSVG                     = `<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.5.5h.5Zm0 4H0a.5.5 0 0 0 .854.354L.5 4.5Zm4-4l.354.354A.5.5 0 0 0 4.5 0v.5ZM2.146 2.854l12 12l.708-.708l-12-12l-.708.708ZM0 .5v4h1v-4H0Zm.854 4.354l4-4l-.708-.708l-4 4l.708.708ZM4.5 0h-4v1h4V0Z"/>`
	arrowRightCircleOutlineInnerSVG          = `<path fill="currentColor" d="m8.146 9.146l-.353.354l.707.707l.354-.353l-.708-.708ZM10.5 7.5l.354.354l.353-.354l-.353-.354l-.354.354ZM8.854 5.146L8.5 4.793l-.707.707l.353.354l.708-.708Zm0 4.708l2-2l-.708-.708l-2 2l.708.708Zm2-2.708l-2-2l-.708.708l2 2l.708-.708ZM10.5 7H4v1h6.5V7Zm-3 7A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	arrowRightCircleSolidInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm8.5-2.707L11.207 7.5L8.5 10.207L7.793 9.5l1.5-1.5H4V7h5.293l-1.5-1.5l.707-.707Z" clip-rule="evenodd"/>`
	arrowRightOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="m13.5 7.5l-4-4m4 4l-4 4m4-4H1"/>`
	arrowRightSmallOutlineInnerSVG           = `<path fill="currentColor" d="m8.146 9.146l-.353.354l.707.707l.354-.353l-.708-.708ZM10.5 7.5l.354.354l.353-.354l-.353-.354l-.354.354ZM8.854 5.146L8.5 4.793l-.707.707l.353.354l.708-.708Zm0 4.708l2-2l-.708-.708l-2 2l.708.708Zm2-2.708l-2-2l-.708.708l2 2l.708-.708ZM10.5 7H4v1h6.5V7Z"/>`
	arrowRightSmallSolidInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M8.5 4.793L11.207 7.5L8.5 10.207L7.793 9.5l1.5-1.5H4V7h5.293l-1.5-1.5l.707-.707Z" clip-rule="evenodd"/>`
	arrowRightSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M9.854 3.146L14.207 7.5l-4.353 4.354l-.708-.708L12.293 8H1V7h11.293L9.146 3.854l.708-.708Z" clip-rule="evenodd"/>`
	arrowSolidInnerSVG                       = `<path fill="currentColor" d="M4.5 0h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .854.354L2.5 3.207l11.646 11.647l.708-.708L3.207 2.5L4.854.854A.5.5 0 0 0 4.5 0Z"/>`
	arrowUpCircleOutlineInnerSVG             = `<path fill="currentColor" d="m9.146 6.854l.354.353l.707-.707l-.353-.354l-.708.708ZM7.5 4.5l.354-.354l-.354-.353l-.354.353l.354.354ZM5.146 6.146l-.353.354l.707.707l.354-.353l-.708-.708ZM14.5 7.5H14h.5Zm-7 7V14v.5Zm0-14V0v.5Zm-7 7H0h.5Zm9.354-1.354l-2-2l-.708.708l2 2l.708-.708Zm-2.708-2l-2 2l.708.708l2-2l-.708-.708ZM7 4.5V11h1V4.5H7Zm7 3A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1Zm-1 0A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0Z"/>`
	arrowUpCircleSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 15a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15ZM4.793 6.5L7.5 3.793L10.207 6.5l-.707.707l-1.5-1.5V11H7V5.707l-1.5 1.5l-.707-.707Z" clip-rule="evenodd"/>`
	arrowUpOutlineInnerSVG                   = `<path fill="currentColor" d="m7.5 1.5l.354-.354L7.5.793l-.354.353l.354.354Zm-.354.354l4 4l.708-.708l-4-4l-.708.708Zm0-.708l-4 4l.708.708l4-4l-.708-.708ZM7 1.5V14h1V1.5H7Z"/>`
	arrowUpSmallOutlineInnerSVG              = `<path fill="currentColor" d="m5.146 6.146l-.353.354l.707.707l.354-.353l-.708-.708ZM7.5 4.5l.354-.354l-.354-.353l-.354.353l.354.354Zm1.646 2.354l.354.353l.707-.707l-.353-.354l-.708.708Zm-3.292 0l2-2l-.708-.708l-2 2l.708.708Zm1.292-2l2 2l.708-.708l-2-2l-.708.708ZM7 4.5V11h1V4.5H7Z"/>`
	arrowUpSmallSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.793L10.207 6.5l-.707.707l-1.5-1.5V11H7V5.707l-1.5 1.5l-.707-.707L7.5 3.793Z" clip-rule="evenodd"/>`
	arrowUpSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="m7.5.793l4.354 4.353l-.707.708L8 2.707V14H7V2.707L3.854 5.854l-.708-.708L7.5.793Z" clip-rule="evenodd"/>`
	artboardOutlineInnerSVG                  = `<path fill="currentColor" d="M4.5 4.5V4a.5.5 0 0 0-.5.5h.5Zm6 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 6v.5a.5.5 0 0 0 .5-.5h-.5Zm-6 0H4a.5.5 0 0 0 .5.5v-.5ZM4 0v2h1V0H4Zm6 0v2h1V0h-1ZM0 5h2V4H0v1Zm0 6h2v-1H0v1Zm13-6h2V4h-2v1Zm0 6h2v-1h-2v1Zm-9 2v2h1v-2H4Zm6 0v2h1v-2h-1ZM4.5 5h6V4h-6v1Zm5.5-.5v6h1v-6h-1Zm.5 5.5h-6v1h6v-1Zm-5.5.5v-6H4v6h1Z"/>`
	artboardSolidInnerSVG                    = `<path fill="currentColor" d="M4 0v2h1V0H4Zm6 0v2h1V0h-1ZM2 5H0V4h2v1Zm-2 6h2v-1H0v1Zm15-6h-2V4h2v1Zm-2 6h2v-1h-2v1Zm-9 4v-2h1v2H4Zm6-2v2h1v-2h-1ZM4.5 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-6Z"/>`
	atOutlineInnerSVG                        = `<path fill="none" stroke="currentColor" d="M10.5 7.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm0 0v1a2 2 0 1 0 4 0v-1a7 7 0 1 0-3 5.745"/>`
	atSolidInnerSVG                          = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0v1a2.499 2.499 0 0 1-4.727 1.136A3.5 3.5 0 1 1 11 7.5v1a1.5 1.5 0 0 0 3 0v-1a6.5 6.5 0 1 0-2.786 5.335l.572.82A7.5 7.5 0 0 1 0 7.5Zm10 0a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0Z" clip-rule="evenodd"/>`
	attachOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m13.5 7.5l-5.757 5.757a4.243 4.243 0 0 1-6-6l5.929-5.929a2.828 2.828 0 0 1 4 4l-5.758 5.758a1.414 1.414 0 0 1-2-2L9.5 3.5"/>`
	attachSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7.318.975a3.328 3.328 0 1 1 4.707 4.707l-5.757 5.757A1.914 1.914 0 1 1 3.56 8.732l5.585-5.586l.708.708l-5.586 5.585a.914.914 0 1 0 1.293 1.293l5.757-5.757a2.328 2.328 0 1 0-3.293-3.293L2.096 7.611a3.743 3.743 0 0 0 5.293 5.293l5.757-5.758l.708.708l-5.758 5.757A4.743 4.743 0 0 1 1.39 6.904L7.318.974Z" clip-rule="evenodd"/>`
	attachmentOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M.5 0v4.5a2 2 0 1 0 4 0v-3a1 1 0 0 0-2 0V5M6 .5h6.5a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1V8M11 4.5H7m4 3H7m4 3H4"/>`
	attachmentSolidInnerSVG                  = `<path fill="currentColor" d="M0 4.5V0h1v4.5a1.5 1.5 0 1 0 3 0v-3a.5.5 0 0 0-1 0V5H2V1.5a1.5 1.5 0 1 1 3 0v3a2.5 2.5 0 0 1-5 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M12.5 0H6v4.5A3.5 3.5 0 0 1 2.5 8H1v5.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 12.5 0ZM11 4H7v1h4V4Zm0 3H7v1h4V7Zm-7 3h7v1H4v-1Z" clip-rule="evenodd"/>`
	audioCableOutlineInnerSVG                = `<path fill="currentColor" d="m6.5 1.5l-.354-.354A.5.5 0 0 0 6 1.5h.5Zm1-1l.354-.354a.5.5 0 0 0-.708 0L7.5.5Zm1 1H9a.5.5 0 0 0-.146-.354L8.5 1.5ZM6.5 9h2V8h-2v1Zm2.5.5v3h1v-3H9ZM8.5 13h-2v1h2v-1ZM6 12.5v-3H5v3h1Zm.5.5a.5.5 0 0 1-.5-.5H5A1.5 1.5 0 0 0 6.5 14v-1Zm2.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5H9ZM8.5 9a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 8.5 8v1Zm-2-1A1.5 1.5 0 0 0 5 9.5h1a.5.5 0 0 1 .5-.5V8Zm.5.5v-5H6v5h1ZM6.5 4h2V3h-2v1ZM8 3.5v5h1v-5H8ZM7 13v2h1v-2H7Zm0-9.5v-2H6v2h1Zm-.146-1.646l1-1l-.708-.708l-1 1l.708.708Zm.292-1l1 1l.708-.708l-1-1l-.708.708ZM8 1.5v2h1v-2H8Z"/>`
	audioCableSolidInnerSVG                  = `<path fill="currentColor" d="M7.854.146a.5.5 0 0 0-.708 0l-1 1A.5.5 0 0 0 6 1.5V3h3V1.5a.5.5 0 0 0-.146-.354l-1-1ZM9 4v4H6V4h3Zm1 5v3.5A1.5 1.5 0 0 1 8.5 14H8v1H7v-1h-.5A1.5 1.5 0 0 1 5 12.5V9h5Z"/>`
	audioDocumentOutlineInnerSVG             = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-3 4l.4-.3a.5.5 0 0 0-.9.3h.5Zm.3.4l.4-.3l-.4.3Zm4.7 9.1h-10v1h10v-1ZM2 13.5v-12H1v12h1Zm11-10v10h1v-10h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM6 11a1 1 0 0 1-1-1H4a2 2 0 0 0 2 2v-1Zm1-1a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2H7ZM6 9a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2v1Zm0-1a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1V8Zm1-1.5V10h1V6.5H7ZM8 7V4.5H7V7h1Zm-.9-2.2l.3.4l.8-.6l-.3-.4l-.8.6Zm.3.4A4.5 4.5 0 0 0 11 7V6a3.5 3.5 0 0 1-2.8-1.4l-.8.6Z"/>`
	audioDocumentSolidInnerSVG               = `<path fill="currentColor" d="M7 10a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm6.342 2.526A.5.5 0 0 1 7.9 4.2l.3.4A3.5 3.5 0 0 0 11 6v1a4.5 4.5 0 0 1-3-1.146V10a2 2 0 1 1-1-1.732V4.5a.5.5 0 0 1 .342-.474Z" clip-rule="evenodd"/>`
	azureOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m3.5 4.5l-3 7h3l4-11l-4 4Zm11 9l-5-11l-2 5l3 4l-6 2h10Z"/>`
	azureSolidInnerSVG                       = `<path fill="currentColor" d="M7.97.67a.5.5 0 0 0-.824-.524l-4 4a.5.5 0 0 0-.106.157l-3 7A.5.5 0 0 0 .5 12h3a.5.5 0 0 0 .47-.33l4-11Zm1.985 1.623a.5.5 0 0 0-.92.021l-2 5A.5.5 0 0 0 7.1 7.8l2.584 3.445l-5.342 1.78A.5.5 0 0 0 4.5 14h10a.5.5 0 0 0 .455-.707l-5-11Z"/>`
	backspaceOutlineInnerSVG                 = `<path fill="currentColor" d="m4.5 12.5l-.39.312A.5.5 0 0 0 4.5 13v-.5Zm-4-5l-.39-.312a.5.5 0 0 0 0 .624L.5 7.5Zm4-5V2a.5.5 0 0 0-.39.188l.39.312Zm9.5 1v8h1v-8h-1Zm-.5 8.5h-9v1h9v-1Zm-8.61.188l-4-5l-.78.624l4 5l.78-.624Zm-4-4.376l4-5l-.78-.624l-4 5l.78.624ZM4.5 3h9V2h-9v1Zm9.5 8.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1Zm1-8A1.5 1.5 0 0 0 13.5 2v1a.5.5 0 0 1 .5.5h1ZM6.146 5.854l4 4l.708-.708l-4-4l-.708.708Zm4-.708l-4 4l.708.708l4-4l-.708-.708Z"/>`
	backspaceSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M4.11 2.188A.5.5 0 0 1 4.5 2h9A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-9a.5.5 0 0 1-.39-.188l-4-5a.5.5 0 0 1 0-.624l4-5Zm6.036 7.666L8.5 8.207L6.854 9.854l-.708-.708L7.793 7.5L6.146 5.854l.708-.708L8.5 6.793l1.646-1.647l.708.708L9.207 7.5l1.647 1.646l-.707.708Z" clip-rule="evenodd"/>`
	bagAltOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M10.5 2v2.5a3 3 0 0 1-6 0V2m-3-.5v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1Z"/>`
	bagAltSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h10A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm9 3a2.5 2.5 0 0 1-5 0V2H4v2.5a3.5 3.5 0 1 0 7 0V2h-1v2.5Z" clip-rule="evenodd"/>`
	bagMinusOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M4.5 4v-.5a3 3 0 0 1 6 0V4M5 9.5h5M2.401 6.39l-.778 7a1 1 0 0 0 .994 1.11h9.766a1 1 0 0 0 .994-1.11l-.778-7a1 1 0 0 0-.994-.89h-8.21a1 1 0 0 0-.994.89Z"/>`
	bagMinusSolidInnerSVG                    = `<path fill="currentColor" d="M7.5 1A2.5 2.5 0 0 0 5 3.5V4H4v-.5a3.5 3.5 0 1 1 7 0V4h-1v-.5A2.5 2.5 0 0 0 7.5 1Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.395 5a1.5 1.5 0 0 0-1.49 1.334l-.778 7A1.5 1.5 0 0 0 2.617 15h9.766a1.5 1.5 0 0 0 1.49-1.666l-.777-7A1.5 1.5 0 0 0 11.606 5h-8.21ZM5 9v1h5V9H5Z" clip-rule="evenodd"/>`
	bagOutlineInnerSVG                       = `<path fill="currentColor" d="m2.401 6.39l-.497-.056l.497.056Zm-.778 7l.497.055l-.497-.055Zm11.754 0l-.497.055l.497-.055Zm-.778-7l.497-.056l-.497.056ZM1.904 6.334l-.778 7l.994.11l.778-7l-.994-.11ZM2.617 15h9.766v-1H2.617v1Zm11.257-1.666l-.778-7l-.994.11l.778 7l.993-.11ZM11.604 5H3.396v1h8.21V5Zm1.492 1.334A1.5 1.5 0 0 0 11.605 5v1a.5.5 0 0 1 .497.445l.994-.11ZM12.383 15a1.5 1.5 0 0 0 1.49-1.666l-.993.11a.5.5 0 0 1-.497.556v1ZM1.126 13.334A1.5 1.5 0 0 0 2.617 15v-1a.5.5 0 0 1-.497-.555l-.994-.11Zm1.772-6.89A.5.5 0 0 1 3.395 6V5a1.5 1.5 0 0 0-1.49 1.334l.993.11ZM5 4v-.5H4V4h1Zm5-.5V4h1v-.5h-1ZM7.5 1A2.5 2.5 0 0 1 10 3.5h1A3.5 3.5 0 0 0 7.5 0v1ZM5 3.5A2.5 2.5 0 0 1 7.5 1V0A3.5 3.5 0 0 0 4 3.5h1Z"/>`
	bagPlusOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M4.5 4v-.5a3 3 0 0 1 6 0V4m-3 3v5M5 9.5h5M2.401 6.39l-.778 7a1 1 0 0 0 .994 1.11h9.766a1 1 0 0 0 .994-1.11l-.778-7a1 1 0 0 0-.994-.89h-8.21a1 1 0 0 0-.994.89Z"/>`
	bagPlusSolidInnerSVG                     = `<path fill="currentColor" d="M5 3.5a2.5 2.5 0 0 1 5 0V4h1v-.5a3.5 3.5 0 1 0-7 0V4h1v-.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M1.904 6.334A1.5 1.5 0 0 1 3.395 5h8.21a1.5 1.5 0 0 1 1.49 1.334l.779 7A1.5 1.5 0 0 1 12.383 15H2.617a1.5 1.5 0 0 1-1.49-1.666l.777-7ZM7 9V7h1v2h2v1H8v2H7v-2H5V9h2Z" clip-rule="evenodd"/>`
	bagSolidInnerSVG                         = `<path fill="currentColor" d="M5 3.5a2.5 2.5 0 0 1 5 0V4h1v-.5a3.5 3.5 0 1 0-7 0V4h1v-.5ZM1.904 6.334A1.5 1.5 0 0 1 3.395 5h8.21a1.5 1.5 0 0 1 1.49 1.334l.779 7A1.5 1.5 0 0 1 12.383 15H2.617a1.5 1.5 0 0 1-1.49-1.666l.777-7Z"/>`
	bankOutlineInnerSVG                      = `<path fill="currentColor" d="m7.5.5l.224-.447a.5.5 0 0 0-.448 0L7.5.5ZM0 15h15v-1H0v1ZM7.276.053l-6 3l.448.894l6-3l-.448-.894ZM0 6h15V5H0v1Zm13.724-2.947l-6-3l-.448.894l6 3l.448-.894ZM5 8v4h1V8H5Zm4 0v4h1V8H9ZM1 5.5v9h1v-9H1Zm12 0v9h1v-9h-1Z"/>`
	bankSolidInnerSVG                        = `<path fill="currentColor" d="M7.724.053a.5.5 0 0 0-.448 0l-6 3l.448.894L7.5 1.06l5.776 2.888l.448-.894l-6-3Z"/><path fill="currentColor" fill-rule="evenodd" d="M14 6h1V5H0v1h1v8H0v1h15v-1h-1V6Zm-9 6V8h1v4H5Zm4 0V8h1v4H9Z" clip-rule="evenodd"/>`
	barChartOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M0 14.5h15M5.5 12V6m4 6V3m4 9V0m-12 9v3"/>`
	barChartSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M13 0h1v12h-1V0Zm-3 3v9H9V3h1ZM6 6v6H5V6h1Zm-5 6V9h1v3H1Zm14 3H0v-1h15v1Z" clip-rule="evenodd"/>`
	barcodeOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M.5 2v11m5-11v11m2-11v11m7-11v11m-4-11v11"/>`
	barcodeSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 13V2h1v11H0Zm5 0V2h1v11H5Zm2 0V2h1v11H7Zm3 0V2h1v11h-1Zm4 0V2h1v11h-1Z" clip-rule="evenodd"/>`
	basketMinusOutlineInnerSVG               = `<path fill="currentColor" d="m2.449 14.398l.447-.224l-.447.224Zm10.102 0l.447.224l-.447-.224ZM.703 6h13.594V5H.703v1ZM14 5.703v.439h1v-.439h-1ZM12.386 14H2.614v1h9.772v-1ZM1 6.142v-.439H0v.439h1Zm1.896 8.032A17.96 17.96 0 0 1 1 6.142H0c0 2.944.685 5.847 2.002 8.48l.894-.448ZM2.614 14c.12 0 .229.068.282.174l-.894.448a.685.685 0 0 0 .612.378v-1Zm9.49.174a.315.315 0 0 1 .282-.174v1c.26 0 .496-.146.612-.378l-.894-.448ZM14 6.142c0 2.788-.65 5.538-1.896 8.032l.894.448A18.96 18.96 0 0 0 15 6.142h-1ZM14.297 6A.297.297 0 0 1 14 5.703h1A.703.703 0 0 0 14.297 5v1ZM.703 5A.703.703 0 0 0 0 5.703h1A.297.297 0 0 1 .703 6V5Zm3.226.757l3-5L6.07.243l-3 5l.858.514Zm4.142-5l3 5l.858-.514l-3-5l-.858.514ZM5 10h5V9H5v1Z"/>`
	basketMinusSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M4.383 5L6.93.757L6.07.243L3.217 5H.703A.703.703 0 0 0 0 5.703v.439c0 2.944.685 5.847 2.002 8.48a.685.685 0 0 0 .612.378h9.772c.26 0 .496-.146.612-.379A18.96 18.96 0 0 0 15 6.141v-.438A.703.703 0 0 0 14.297 5h-2.514L8.93.243l-.86.514L10.617 5H4.383ZM5 10h5V9H5v1Z" clip-rule="evenodd"/>`
	basketOutlineInnerSVG                    = `<path fill="currentColor" d="m2.449 14.398l.447-.224l-.447.224Zm10.102 0l.447.224l-.447-.224ZM.703 6h13.594V5H.703v1ZM14 5.703v.439h1v-.439h-1ZM12.386 14H2.614v1h9.772v-1ZM1 6.142v-.439H0v.439h1Zm1.896 8.032A17.96 17.96 0 0 1 1 6.142H0c0 2.944.685 5.847 2.002 8.48l.894-.448ZM2.614 14c.12 0 .229.068.282.174l-.894.448a.685.685 0 0 0 .612.378v-1Zm9.49.174a.315.315 0 0 1 .282-.174v1c.26 0 .496-.146.612-.378l-.894-.448ZM14 6.142c0 2.788-.65 5.538-1.896 8.032l.894.448A18.96 18.96 0 0 0 15 6.142h-1ZM14.297 6A.297.297 0 0 1 14 5.703h1A.703.703 0 0 0 14.297 5v1ZM.703 5A.703.703 0 0 0 0 5.703h1A.297.297 0 0 1 .703 6V5Zm3.226.757l3-5L6.07.243l-3 5l.858.514Zm4.142-5l3 5l.858-.514l-3-5l-.858.514Z"/>`
	basketPlusOutlineInnerSVG                = `<path fill="currentColor" d="m2.449 14.398l.447-.224l-.447.224Zm10.102 0l.447.224l-.447-.224ZM.703 6h13.594V5H.703v1ZM14 5.703v.439h1v-.439h-1ZM12.386 14H2.614v1h9.772v-1ZM1 6.142v-.439H0v.439h1Zm1.896 8.032A17.96 17.96 0 0 1 1 6.142H0c0 2.944.685 5.847 2.002 8.48l.894-.448ZM2.614 14c.12 0 .229.068.282.174l-.894.448a.685.685 0 0 0 .612.378v-1Zm9.49.174a.315.315 0 0 1 .282-.174v1c.26 0 .496-.146.612-.378l-.894-.448ZM14 6.142c0 2.788-.65 5.538-1.896 8.032l.894.448A18.96 18.96 0 0 0 15 6.142h-1ZM14.297 6A.297.297 0 0 1 14 5.703h1A.703.703 0 0 0 14.297 5v1ZM.703 5A.703.703 0 0 0 0 5.703h1A.297.297 0 0 1 .703 6V5Zm3.226.757l3-5L6.07.243l-3 5l.858.514Zm4.142-5l3 5l.858-.514l-3-5l-.858.514ZM7 7v5h1V7H7Zm-2 3h5V9H5v1Z"/>`
	basketPlusSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M4.383 5L6.93.757L6.07.243L3.217 5H.703A.703.703 0 0 0 0 5.703v.439c0 2.944.685 5.847 2.002 8.48a.685.685 0 0 0 .612.378h9.772c.26 0 .496-.146.612-.379A18.96 18.96 0 0 0 15 6.141v-.438A.703.703 0 0 0 14.297 5h-2.514L8.93.243l-.86.514L10.617 5H4.383ZM7 12v-2H5V9h2V7h1v2h2v1H8v2H7Z" clip-rule="evenodd"/>`
	basketSolidInnerSVG                      = `<path fill="currentColor" d="M6.929.757L4.383 5h6.234L8.07.757l.86-.514L11.783 5h2.514c.388 0 .703.315.703.703v.439a18.96 18.96 0 0 1-2.002 8.48a.684.684 0 0 1-.612.378H2.614a.685.685 0 0 1-.612-.379A18.96 18.96 0 0 1 0 6.141v-.438C0 5.315.315 5 .703 5h2.514L6.07.243l.858.514Z"/>`
	bathOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M0 7.5h15m-10.5 5h6m-6 0a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h2V2m-2 10.5V15m6-2.5a3 3 0 0 0 3-3v-2m-3 5V15M5 3.5h3"/>`
	bathSolidInnerSVG                        = `<path fill="currentColor" d="M2 3.5A2.5 2.5 0 0 1 4.5 1H6v1h1V0H4.5A3.5 3.5 0 0 0 1 3.5V7H0v1h1v1.5a3.5 3.5 0 0 0 3 3.465V15h1v-2h5v2h1v-2.035A3.501 3.501 0 0 0 14 9.5V8h1V7H2V3.5Z"/><path fill="currentColor" d="M8 4H5V3h3v1Z"/>`
	batteryChargeOutlineInnerSVG             = `<path fill="currentColor" d="M12.5 11.5H12h.5Zm-1 1v.5v-.5Zm0-10V2v.5Zm1 1h.5h-.5Zm-12 0H0h.5Zm1-1V3v-.5Zm-1 9H1H.5Zm1 1V12v.5Zm5-3l-.224.447A.5.5 0 0 0 7 9.5h-.5Zm0-4l.224-.447A.5.5 0 0 0 6 5.5h.5Zm-5.5 6v-8H0v8h1ZM1.5 3h10V2h-10v1Zm10.5.5v8h1v-8h-1Zm-.5 8.5h-10v1h10v-1Zm.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM11.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 2v1ZM1 3.5a.5.5 0 0 1 .5-.5V2A1.5 1.5 0 0 0 0 3.5h1Zm-1 8A1.5 1.5 0 0 0 1.5 13v-1a.5.5 0 0 1-.5-.5H0ZM15 10V5h-1v5h1ZM2.276 7.947l4 2l.448-.894l-4-2l-.448.894ZM7 9.5v-4H6v4h1Zm-.724-3.553l4 2l.448-.894l-4-2l-.448.894Z"/>`
	batteryChargeSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5v8Zm6.724-6.447A.5.5 0 0 0 6 5.5v3.191L2.724 7.053l-.448.894l4 2A.5.5 0 0 0 7 9.5V6.309l3.276 1.638l.448-.894l-4-2Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5h1Z"/>`
	batteryFiveOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m2 7V4m2 7V4m2 7V4m2 7V4m2 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`
	batteryFiveSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5v8Zm3-.5V4H2v7h1Zm2 0V4H4v7h1Zm2-7v7H6V4h1Zm2 7V4H8v7h1Zm2-7v7h-1V4h1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5h1Z"/>`
	batteryFourOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m2 7V4m2 7V4m2 7V4m4 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`
	batteryFourSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5v8ZM3 4v7H2V4h1Zm2 0v7H4V4h1Zm2 7V4H6v7h1Zm2-7v7H8V4h1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5h1Z"/>`
	batteryOneOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m10 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`
	batteryOneSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5v8Zm3-.5V4H2v7h1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5h1Z"/>`
	batteryThreeOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m2 7V4m2 7V4m6 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`
	batteryThreeSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5v8Zm3-.5V4H2v7h1Zm2 0V4H4v7h1Zm2-7v7H6V4h1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5h1Z"/>`
	batteryTwoOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m2 7V4m8 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`
	batteryTwoSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5v8ZM3 4v7H2V4h1Zm2 0v7H4V4h1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5h1Z"/>`
	batteryZeroOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M14.5 10V5m-2 6.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`
	batteryZeroSolidInnerSVG                 = `<path fill="currentColor" d="M1.5 13A1.5 1.5 0 0 1 0 11.5v-8A1.5 1.5 0 0 1 1.5 2h10A1.5 1.5 0 0 1 13 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-10ZM15 10V5h-1v5h1Z"/>`
	bedDoubleOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M14.5 7v8M.5 7v8M0 10.5h15m-15-3h15m-13-2h4m3 0h4M.5 6V.5h14V6"/>`
	bedDoubleSolidInnerSVG                   = `<path fill="currentColor" d="M0 0h15v6h-1V1H1v5H0V0Z"/><path fill="currentColor" d="M6 6H2V5h4v1Zm-6 9h1v-4h13v4h1V7H0v8Zm9-9h4V5H9v1Z"/>`
	bedSingleOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M14.5 4v11M.5 0v15M0 10.5h15m-15-3h15m-13-2h4"/>`
	bedSingleSolidInnerSVG                   = `<path fill="currentColor" d="M0 15h1v-4h13v4h1V4h-1v3H1V0H0v15Z"/><path fill="currentColor" d="M6 6H2V5h4v1Z"/>`
	behanceOutlineInnerSVG                   = `<path fill="currentColor" d="M14.5 9.5v.5h.5v-.5h-.5Zm-14-7V2H0v.5h.5Zm0 10H0v.5h.5v-.5Zm7.5-3a3.5 3.5 0 0 0 3.5 3.5v-1A2.5 2.5 0 0 1 9 9.5H8ZM11.5 6A3.5 3.5 0 0 0 8 9.5h1A2.5 2.5 0 0 1 11.5 7V6ZM15 9.5A3.5 3.5 0 0 0 11.5 6v1A2.5 2.5 0 0 1 14 9.5h1Zm-1.473 1.464A2.496 2.496 0 0 1 11.5 12v1a3.496 3.496 0 0 0 2.837-1.45l-.81-.586ZM.5 3H3V2H.5v1ZM3 7H.5v1H3V7Zm-2 .5v-5H0v5h1ZM5 5a2 2 0 0 1-2 2v1a3 3 0 0 0 3-3H5ZM3 3a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3v1ZM.5 8H4V7H.5v1ZM4 12H.5v1H4v-1Zm-3 .5v-5H0v5h1ZM6 10a2 2 0 0 1-2 2v1a3 3 0 0 0 3-3H6ZM4 8a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3v1Zm4.5 2h6V9h-6v1ZM9 4h5V3H9v1Z"/>`
	behanceSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 2h3a3 3 0 0 1 2.051 5.19A3.001 3.001 0 0 1 4 13H0V2Zm1 6v4h3a2 2 0 1 0 0-4H1Zm0-1h2a2 2 0 1 0 0-4H1v4Zm13-3H9V3h5v1ZM8 9.5a3.5 3.5 0 1 1 7 0v.5H9.05a2.5 2.5 0 0 0 4.477.964l.81.586A3.5 3.5 0 0 1 8 9.5ZM9.05 9h4.9a2.5 2.5 0 0 0-4.9 0Z" clip-rule="evenodd"/>`
	bellOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M1 10.5h13m-11.5 0v-5a5 5 0 0 1 10 0v5m-7 1.5v.5a2 2 0 1 0 4 0V12"/>`
	bellSolidInnerSVG                        = `<path fill="currentColor" d="M7.5 0A5.5 5.5 0 0 0 2 5.5V10H1v1h13v-1h-1V5.5A5.5 5.5 0 0 0 7.5 0ZM5 12.5V12h5v.5a2.5 2.5 0 0 1-5 0Z"/>`
	binOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M4.5 3V1.5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1V3M0 3.5h15m-13.5 0v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-10M7.5 7v5m-3-3v3m6-3v3"/>`
	binSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M11 3V1.5A1.5 1.5 0 0 0 9.5 0h-4A1.5 1.5 0 0 0 4 1.5V3H0v1h1v9.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4h1V3h-4ZM5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3H5V1.5ZM7 7v5h1V7H7Zm-3 5V9h1v3H4Zm6-3v3h1V9h-1Z" clip-rule="evenodd"/>`
	bitbucketOutlineInnerSVG                 = `<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.495.57L.5.5Zm14 0l.495.07A.5.5 0 0 0 14.5 0v.5Zm-2 14v.5a.5.5 0 0 0 .495-.43l-.495-.07Zm-10 0l-.495.07A.5.5 0 0 0 2.5 15v-.5ZM5 4.5V4a.5.5 0 0 0-.498.542L5 4.5Zm4.5 6v.5a.5.5 0 0 0 .498-.459L9.5 10.5Zm-4 0l-.498.041A.5.5 0 0 0 5.5 11v-.5ZM.5 1h14V0H.5v1ZM14.005.43l-2 14l.99.14l2-14l-.99-.14ZM12.5 14h-10v1h10v-1Zm-9.505.43l-2-14l-.99.14l2 14l.99-.14ZM5 5h5V4H5v1Zm4.502-.542l-.5 6l.996.083l.5-6l-.996-.083ZM9.5 10h-4v1h4v-1Zm-3.502.459l-.5-6l-.996.083l.5 6l.996-.083ZM10 5h4V4h-4v1Z"/>`
	bitbucketSolidInnerSVG                   = `<path fill="currentColor" d="M.5 0a.5.5 0 0 0-.495.57l2 14A.5.5 0 0 0 2.5 15h10a.5.5 0 0 0 .495-.43L14.219 6H9.373l-.333 4H5.96l-.417-5h8.82l.632-4.43A.5.5 0 0 0 14.5 0H.5Z"/>`
	bitcoinOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M3.5 1.5h5a3 3 0 1 1 0 6h-5m0-6v6m0-6H2m1.5 0V0m0 7.5h6a3 3 0 1 1 0 6h-6m0-6v6m0-6H2m1.5 6H2m1.5 0V15m4-15v1.5m0 12V15"/>`
	bitcoinSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M3 1V0h1v1h3V0h1v1h.5a3.5 3.5 0 0 1 2.21 6.215A3.501 3.501 0 0 1 9.5 14H8v1H7v-1H4v1H3v-1H2v-1h1V8H2V7h1V2H2V1h1Zm1 1v5h4.5a2.5 2.5 0 0 0 0-5H4Zm0 6h5.5a2.5 2.5 0 0 1 0 5H4V8Z" clip-rule="evenodd"/>`
	bluetoothOutlineInnerSVG                 = `<path fill="currentColor" d="M7.5 14.5H7a.5.5 0 0 0 .787.41L7.5 14.5Zm0-14l.287-.41A.5.5 0 0 0 7 .5h.5Zm5 3.5l.282.413a.5.5 0 0 0 .005-.823L12.5 4Zm0 7l.287.41a.5.5 0 0 0-.005-.823L12.5 11ZM8 14.5V7.41H7v7.09h1Zm0-7.09V.5H7v6.91h1ZM7.213.91l5 3.5l.574-.82l-5-3.5l-.574.82Zm5.005 2.677l-5 3.409l.564.826l5-3.409l-.564-.826Zm-5 3.409l-6 4.09l.564.827l6-4.09l-.564-.827Zm.569 7.914l5-3.5l-.574-.82l-5 3.5l.574.82Zm4.995-4.323l-11-7.5l-.564.826l11 7.5l.564-.826Z"/>`
	bluetoothSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M7.27.057a.5.5 0 0 1 .517.033l5 3.5a.5.5 0 0 1-.005.823L8.254 7.5l4.528 3.087a.5.5 0 0 1 .005.823l-5 3.5A.5.5 0 0 1 7 14.5V8.355l-5.218 3.558l-.564-.826L6.48 7.5L1.22 3.913l.563-.826L7 6.645V.5a.5.5 0 0 1 .27-.443ZM8 8.537l3.62 2.468L8 13.54V8.537Zm0-2.074V1.46l3.62 2.535L8 6.463Z" clip-rule="evenodd"/>`
	boldOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M3.5 7.5h5a3 3 0 1 0 0-6H3.58a.08.08 0 0 0-.08.08V7.5Zm0 0h6a3 3 0 1 1 0 6H3.59a.09.09 0 0 1-.09-.09V7.5Z"/>`
	boldSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M3 1.58c0-.32.26-.58.58-.58H8.5a3.5 3.5 0 0 1 2.21 6.215A3.501 3.501 0 0 1 9.5 14H3.59a.59.59 0 0 1-.59-.59V1.58ZM4 8v5h5.5a2.5 2.5 0 0 0 0-5H4Zm0-1h4.5a2.5 2.5 0 0 0 0-5H4v5Z" clip-rule="evenodd"/>`
	bookOutlineInnerSVG                      = `<path fill="currentColor" d="M1.5.5V0a.5.5 0 0 0-.5.5h.5Zm0 13H1a.5.5 0 0 0 .5.5v-.5ZM4 0v15h1V0H4ZM1.5 1h10V0h-10v1ZM13 2.5v9h1v-9h-1ZM11.5 13h-10v1h10v-1Zm-9.5.5V.5H1v13h1Zm11-2a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5h-1ZM11.5 1A1.5 1.5 0 0 1 13 2.5h1A2.5 2.5 0 0 0 11.5 0v1ZM7 5h4V4H7v1Z"/>`
	bookSolidInnerSVG                        = `<path fill="currentColor" d="M1.5 0a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 .5.5H3V0H1.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M4 15h1v-1h6.5a2.5 2.5 0 0 0 2.5-2.5v-9A2.5 2.5 0 0 0 11.5 0H4v15Zm7-10H7V4h4v1Z" clip-rule="evenodd"/>`
	bookmarkOutlineInnerSVG                  = `<path fill="currentColor" d="m12.5 14.5l-.312.39A.5.5 0 0 0 13 14.5h-.5Zm0-14h.5V0h-.5v.5Zm-10 0V0H2v.5h.5Zm0 14H2a.5.5 0 0 0 .812.39L2.5 14.5Zm5-4l.312-.39a.5.5 0 0 0-.624 0l.312.39Zm5.5 4V.5h-1v14h1ZM2 .5v14h1V.5H2Zm.812 14.39l5-4l-.624-.78l-5 4l.624.78Zm4.376-4l5 4l.624-.78l-5-4l-.624.78ZM12.5 0h-10v1h10V0Z"/>`
	bookmarkSolidInnerSVG                    = `<path fill="currentColor" d="M13 0H2v14.5a.5.5 0 0 0 .812.39L7.5 11.14l4.688 3.75A.5.5 0 0 0 13 14.5V0Z"/>`
	borderAllOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M7.5 1.5v12m-6-6h12m-12-6h12v12h-12v-12Z"/>`
	borderAllSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v13H1V1Zm1 1v5h5V2H2Zm6 0v5h5V2H8Zm5 6H8v5h5V8Zm-6 5V8H2v5h5Z" clip-rule="evenodd"/>`
	borderBottomOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M1 1.5h1m5 0h1m2 0h1m2 0h1m-1 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m5 3h1m-7 0h1m-7 0h1m-1-3h1m2 0h1m-4-3h1m2-3h1m-4 12h13"/>`
	borderBottomSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1v1Zm3 0H4V1h1v1Zm3 0H7V1h1v1Zm3 0h-1V1h1v1Zm3 0h-1V1h1v1ZM2 5H1V4h1v1Zm6 0H7V4h1v1Zm6 0h-1V4h1v1ZM2 8H1V7h1v1Zm3 0H4V7h1v1Zm3 0H7V7h1v1Zm3 0h-1V7h1v1Zm3 0h-1V7h1v1ZM2 11H1v-1h1v1Zm6 0H7v-1h1v1Zm6 0h-1v-1h1v1Zm0 3H1v-1h13v1Z" clip-rule="evenodd"/>`
	borderHorizontalOutlineInnerSVG          = `<path fill="none" stroke="currentColor" d="M1 1.5h1m5 0h1m2 0h1m2 0h1m-1 3h1m-7 0h1m5 6h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1m-1-6h1m2-3h1m-4 6h13"/>`
	borderHorizontalSolidInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1v1Zm3 0H4V1h1v1Zm3 0H7V1h1v1Zm3 0h-1V1h1v1Zm3 0h-1V1h1v1ZM2 5H1V4h1v1Zm6 0H7V4h1v1Zm6 0h-1V4h1v1Zm0 3H1V7h13v1ZM2 11H1v-1h1v1Zm6 0H7v-1h1v1Zm6 0h-1v-1h1v1ZM2 14H1v-1h1v1Zm3 0H4v-1h1v1Zm3 0H7v-1h1v1Zm3 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	borderInnerOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M1 1.5h1m8 0h1m2 0h1m-1 3h1m-1 6h1m-1 3h1m-4 0h1m-7 0h1m-4 0h1m-1-3h1m-1-6h1m2-3h1m-4 6h13M7.5 1v13"/>`
	borderInnerSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1v1Zm3 0H4V1h1v1Zm2 5H1v1h6v6h1V8h6V7H8V1H7v6Zm4-5h-1V1h1v1Zm3 0h-1V1h1v1ZM2 5H1V4h1v1Zm12 0h-1V4h1v1ZM2 11H1v-1h1v1Zm12 0h-1v-1h1v1ZM2 14H1v-1h1v1Zm3 0H4v-1h1v1Zm6 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	borderLeftOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M7 1.5h1m2 0h1m2 0h1m-1 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m5 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m-4 0h1m-1-6h1m-1-6h1M1.5 1v13"/>`
	borderLeftSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M1 14V1h1v13H1ZM5 2H4V1h1v1Zm3 0H7V1h1v1Zm3 0h-1V1h1v1Zm3 0h-1V1h1v1ZM8 5H7V4h1v1Zm6 0h-1V4h1v1ZM5 8H4V7h1v1Zm3 0H7V7h1v1Zm3 0h-1V7h1v1Zm3 0h-1V7h1v1Zm-6 3H7v-1h1v1Zm6 0h-1v-1h1v1Zm-9 3H4v-1h1v1Zm3 0H7v-1h1v1Zm3 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	borderNoneOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M1 1.5h1m5 0h1m2 0h1m2 0h1m-1 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m5 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m2-3h1"/>`
	borderNoneSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1v1Zm3 0H4V1h1v1Zm3 0H7V1h1v1Zm3 0h-1V1h1v1Zm3 0h-1V1h1v1ZM2 5H1V4h1v1Zm6 0H7V4h1v1Zm6 0h-1V4h1v1ZM2 8H1V7h1v1Zm3 0H4V7h1v1Zm3 0H7V7h1v1Zm3 0h-1V7h1v1Zm3 0h-1V7h1v1ZM2 11H1v-1h1v1Zm6 0H7v-1h1v1Zm6 0h-1v-1h1v1ZM2 14H1v-1h1v1Zm3 0H4v-1h1v1Zm3 0H7v-1h1v1Zm3 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	borderOuterOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M7 4.5h1m2 3h1m-4 0h1m-1 3h1m-4-3h1m-3.5-6h12v12h-12v-12Z"/>`
	borderOuterSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v13H1V1Zm1 1v11h11V2H2Zm6 3H7V4h1v1ZM5 8H4V7h1v1Zm3 0H7V7h1v1Zm3 0h-1V7h1v1Zm-3 3H7v-1h1v1Z" clip-rule="evenodd"/>`
	borderRadiusOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M10 1.5h1m2 0h1m-1 3h1m-1 3h1m-1 3h1m-1 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1M1.5 8V5.5a4 4 0 0 1 4-4H8"/>`
	borderRadiusSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 2A3.5 3.5 0 0 0 2 5.5V8H1V5.5A4.5 4.5 0 0 1 5.5 1H8v1H5.5ZM11 2h-1V1h1v1Zm3 0h-1V1h1v1Zm0 3h-1V4h1v1Zm0 3h-1V7h1v1ZM2 11H1v-1h1v1Zm12 0h-1v-1h1v1ZM2 14H1v-1h1v1Zm3 0H4v-1h1v1Zm3 0H7v-1h1v1Zm3 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	borderRightOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M1 1.5h1m5 0h1m2 0h1m-4 3h1m2 3h1m-4 0h1m-1 3h1m2 3h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m2-3h1m8.5-.5v13"/>`
	borderRightSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1v1Zm3 0H4V1h1v1Zm3 0H7V1h1v1Zm3 0h-1V1h1v1Zm2 12V1h1v13h-1ZM2 5H1V4h1v1Zm6 0H7V4h1v1ZM2 8H1V7h1v1Zm3 0H4V7h1v1Zm3 0H7V7h1v1Zm3 0h-1V7h1v1Zm-9 3H1v-1h1v1Zm6 0H7v-1h1v1Zm-6 3H1v-1h1v1Zm3 0H4v-1h1v1Zm3 0H7v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	borderTopOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M13 4.5h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m5 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m-1-3h13"/>`
	borderTopSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M14 2H1V1h13v1ZM2 5H1V4h1v1Zm6 0H7V4h1v1Zm6 0h-1V4h1v1ZM2 8H1V7h1v1Zm3 0H4V7h1v1Zm3 0H7V7h1v1Zm3 0h-1V7h1v1Zm3 0h-1V7h1v1ZM2 11H1v-1h1v1Zm6 0H7v-1h1v1Zm6 0h-1v-1h1v1ZM2 14H1v-1h1v1Zm3 0H4v-1h1v1Zm3 0H7v-1h1v1Zm3 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	borderVerticalOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M1 1.5h1m8 0h1m2 0h1m-1 3h1m-1 3h1m-4 0h1m2 3h1m-1 3h1m-4 0h1m-7 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m2-3h1M7.5 1v13"/>`
	borderVerticalSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1v1Zm3 0H4V1h1v1Zm2 12V1h1v13H7Zm4-12h-1V1h1v1Zm3 0h-1V1h1v1ZM2 5H1V4h1v1Zm12 0h-1V4h1v1ZM2 8H1V7h1v1Zm3 0H4V7h1v1Zm6 0h-1V7h1v1Zm3 0h-1V7h1v1ZM2 11H1v-1h1v1Zm12 0h-1v-1h1v1ZM2 14H1v-1h1v1Zm3 0H4v-1h1v1Zm6 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	bottomLeftOutlineInnerSVG                = `<path fill="currentColor" d="M1.5 13.5H1a.5.5 0 0 0 .5.5v-.5Zm0 .5H7v-1H1.5v1Zm.5-.5V8H1v5.5h1Zm-.146.354l12-12l-.708-.708l-12 12l.708.708Z"/>`
	bottomLeftSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L2.707 13H7v1H1V8h1v4.293L13.146 1.146l.708.708Z" clip-rule="evenodd"/>`
	bottomRightOutlineInnerSVG               = `<path fill="currentColor" d="M13.5 13.5v.5a.5.5 0 0 0 .5-.5h-.5Zm0-.5H8v1h5.5v-1Zm.5.5V8h-1v5.5h1Zm-.146-.354l-12-12l-.708.708l12 12l.707-.708Z"/>`
	bottomRightSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M1.854 1.146L13 12.293V8h1v6H8v-1h4.293L1.146 1.854l.708-.708Z" clip-rule="evenodd"/>`
	boxOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M.5 3.498L7.5.5l7 2.998m-14 0l7 2.998m-7-2.998V3.5m14-.002l-7 2.998m7-2.998V11.5l-7 3m7-11.002L7.5 6.5v8m0-8.004V14.5m0-8.004L.5 3.5m7 11l-7-3v-8"/>`
	boxSolidInnerSVG                         = `<path fill="currentColor" d="M7.303.04a.5.5 0 0 1 .394 0L14.5 2.956l-7 3l-7-3L7.303.04ZM0 3.83v7.67c0 .2.12.38.303.46L7 14.83v-8l-7-3Zm8 3l7-3v7.67a.5.5 0 0 1-.303.46L8 14.83v-8Z"/>`
	bracketOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M0 7.5h2m13 0h-2m-8 7H3.5a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2H5m5 14h1.5a2 2 0 0 0 2-2v-10a2 2 0 0 0-2-2H10"/>`
	bracketSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 1A1.5 1.5 0 0 0 2 2.5v10A1.5 1.5 0 0 0 3.5 14H5v1H3.5A2.5 2.5 0 0 1 1 12.5V8H0V7h1V2.5A2.5 2.5 0 0 1 3.5 0H5v1H3.5ZM10 0h1.5A2.5 2.5 0 0 1 14 2.5V7h1v1h-1v4.5a2.5 2.5 0 0 1-2.5 2.5H10v-1h1.5a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 11.5 1H10V0Z" clip-rule="evenodd"/>`
	briefcaseAltOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M5.5 3v-.5a2 2 0 1 1 4 0V3m-9 3.5c3.706 4.235 10.294 4.235 14 0m-13-3h12a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1Z"/>`
	briefcaseAltSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5v9A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-9A1.5 1.5 0 0 0 13.5 3H10v-.5a2.5 2.5 0 0 0-5 0ZM7.5 1A1.5 1.5 0 0 0 6 2.5V3h3v-.5A1.5 1.5 0 0 0 7.5 1ZM.66 7.367a10.083 10.083 0 0 0 13.68 0l-.68-.734a9.083 9.083 0 0 1-12.32 0l-.68.734Z" clip-rule="evenodd"/>`
	briefcaseOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M5.5 3v-.5a2 2 0 1 1 4 0V3m-9 8.5h14m-13-8h12a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1Z"/>`
	briefcaseSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5V11h15V4.5A1.5 1.5 0 0 0 13.5 3H10v-.5a2.5 2.5 0 0 0-5 0ZM7.5 1A1.5 1.5 0 0 0 6 2.5V3h3v-.5A1.5 1.5 0 0 0 7.5 1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 12H0v1.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5V12Z"/>`
	brushOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M13.5 7.5v-5a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v5m12 0h-12m12 0v2a2 2 0 0 1-2 2h-3v2a1 1 0 1 1-2 0v-2h-3a2 2 0 0 1-2-2v-2m4-7V5m2-4.5V3"/>`
	brushSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M1 2.5A2.5 2.5 0 0 1 3.5 0H5v5h1V0h1v3h1V0h3.5A2.5 2.5 0 0 1 14 2.5v7a2.5 2.5 0 0 1-2.5 2.5H9v1.5a1.5 1.5 0 0 1-3 0V12H3.5A2.5 2.5 0 0 1 1 9.5v-7ZM2 8v1.5A1.5 1.5 0 0 0 3.5 11H7v2.5a.5.5 0 0 0 1 0V11h3.5A1.5 1.5 0 0 0 13 9.5V8H2Z" clip-rule="evenodd"/>`
	bugOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M4.5 4.5h6m-6 0l-.367.733A6 6 0 0 0 3.5 7.916V10.5a4 4 0 0 0 8 0V7.916a6 6 0 0 0-.633-2.683L10.5 4.5m-6 0v-1a3 3 0 0 1 6 0v1M0 8.5h3.5m11.5 0h-3.5M1 14l3-1.5M14 14l-3-1.5M1 3l3 1.5M14 3l-3 1.5"/>`
	bugSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M4 3.941V3.5a3.5 3.5 0 1 1 7 0v.441l2.776-1.388l.448.894l-2.953 1.477l.043.086A6.5 6.5 0 0 1 12 7.916V8h3v1h-3v1.5c0 .625-.127 1.22-.358 1.762l2.582 1.29l-.448.895l-2.627-1.313A4.494 4.494 0 0 1 7.5 15a4.494 4.494 0 0 1-3.65-1.866l-2.626 1.313l-.448-.894l2.582-1.291A4.486 4.486 0 0 1 3 10.5V9H0V8h3v-.084a6.5 6.5 0 0 1 .686-2.906l.043-.086L.776 3.447l.448-.894L4 3.94ZM5 3.5a2.5 2.5 0 0 1 5 0V4H5v-.5Z" clip-rule="evenodd"/>`
	buildingOutlineInnerSVG                  = `<path fill="currentColor" d="m7.5.5l.224-.447a.5.5 0 0 0-.448 0L7.5.5Zm-3 8V8H4v.5h.5Zm4 0H9V8h-.5v.5ZM0 15h15v-1H0v1ZM7.276.053l-6 3l.448.894l6-3l-.448-.894ZM0 6h15V5H0v1Zm13.724-2.947l-6-3l-.448.894l6 3l.448-.894ZM1 5.5v9h1v-9H1Zm12 0v9h1v-9h-1Zm-8 9v-6H4v6h1ZM4.5 9h4V8h-4v1ZM8 8.5v6h1v-6H8Z"/>`
	buildingSolidInnerSVG                    = `<path fill="currentColor" d="M7.724.053a.5.5 0 0 0-.448 0l-6 3l.448.894L7.5 1.06l5.776 2.888l.448-.894l-6-3ZM14 6h1V5H0v1h1v8H0v1h4V8h5v7h6v-1h-1V6Z"/><path fill="currentColor" d="M8 15V9H5v6h3Z"/>`
	bulbOffOutlineInnerSVG                   = `<path fill="currentColor" d="M4.076 6.47L3.58 6.4l.495.07Zm-.01.07l.495.07l-.495-.07Zm6.858-.07l-.495.07l.495-.07Zm.01.07l.495-.07l-.495.07ZM9.5 12.5v.5a.5.5 0 0 0 .5-.5h-.5Zm-4 0H5a.5.5 0 0 0 .5.5v-.5Zm-.745-3.347l.396-.306l-.396.306Zm5.49 0l-.396-.306l.396.306ZM6 15h3v-1H6v1ZM3.58 6.4l-.01.07l.99.14l.01-.07l-.99-.14ZM7.5 3a3.959 3.959 0 0 0-3.92 3.4l.99.14A2.959 2.959 0 0 1 7.5 4V3Zm3.92 3.4A3.959 3.959 0 0 0 7.5 3v1a2.96 2.96 0 0 1 2.93 2.54l.99-.14Zm.01.07l-.01-.07l-.99.14l.01.07l.99-.14Zm-.79 2.989c.63-.814.948-1.875.79-2.99l-.99.142a2.951 2.951 0 0 1-.59 2.236l.79.612ZM9 10.9v1.6h1v-1.599H9Zm.5 1.1h-4v1h4v-1Zm-3.5.5v-1.599H5V12.5h1ZM3.57 6.47a3.951 3.951 0 0 0 .79 2.989l.79-.612a2.951 2.951 0 0 1-.59-2.236l-.99-.142ZM6 10.9c0-.823-.438-1.523-.85-2.054l-.79.612c.383.495.64.968.64 1.442h1Zm3.85-2.054C9.437 9.378 9 10.077 9 10.9h1c0-.474.257-.947.64-1.442l-.79-.612Z"/>`
	bulbOffSolidInnerSVG                     = `<path fill="currentColor" d="M7.5 3a3.959 3.959 0 0 0-3.92 3.4l-.01.07a3.951 3.951 0 0 0 .79 2.989c.383.495.64.968.64 1.442V12.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-1.599c0-.474.256-.947.64-1.442c.63-.814.948-1.875.79-2.99l-.01-.07A3.959 3.959 0 0 0 7.5 3ZM6 15h3v-1H6v1Z"/>`
	bulbOnOutlineInnerSVG                    = `<path fill="currentColor" d="m4.076 6.47l.495.07l-.495-.07Zm-.01.07l-.495-.07l.495.07Zm6.858-.07l.495-.07l-.495.07Zm.01.07l-.495.07l.495-.07ZM9.5 12.5v.5a.5.5 0 0 0 .5-.5h-.5Zm-4 0H5a.5.5 0 0 0 .5.5v-.5Zm-.745-3.347l.396-.306l-.396.306Zm5.49 0l-.396-.306l.396.306ZM6 15h3v-1H6v1ZM3.58 6.4l-.01.07l.99.14l.01-.07l-.99-.14ZM7.5 3a3.959 3.959 0 0 0-3.92 3.4l.99.14A2.959 2.959 0 0 1 7.5 4V3Zm3.92 3.4A3.959 3.959 0 0 0 7.5 3v1a2.96 2.96 0 0 1 2.93 2.54l.99-.14Zm.01.07l-.01-.07l-.99.14l.01.07l.99-.14Zm-.79 2.989c.63-.814.948-1.875.79-2.99l-.99.142a2.951 2.951 0 0 1-.59 2.236l.79.612ZM9 10.9v1.6h1v-1.599H9Zm.5 1.1h-4v1h4v-1Zm-3.5.5v-1.599H5V12.5h1ZM3.57 6.47a3.951 3.951 0 0 0 .79 2.989l.79-.612a2.951 2.951 0 0 1-.59-2.236l-.99-.142ZM6 10.9c0-.823-.438-1.523-.85-2.054l-.79.612c.383.495.64.968.64 1.442h1Zm3.85-2.054C9.437 9.378 9 10.077 9 10.9h1c0-.474.257-.947.64-1.442l-.79-.612ZM7 0v2h1V0H7ZM0 8h2V7H0v1Zm13 0h2V7h-2v1ZM3.354 3.646l-1.5-1.5l-.708.708l1.5 1.5l.708-.708Zm9 .708l1.5-1.5l-.708-.708l-1.5 1.5l.708.708Z"/>`
	bulbOnSolidInnerSVG                      = `<path fill="currentColor" d="M7 0v2h1V0H7ZM3.354 3.646l-1.5-1.5l-.708.708l1.5 1.5l.708-.708Zm9 .708l1.5-1.5l-.708-.708l-1.5 1.5l.708.708ZM7.5 3a3.959 3.959 0 0 0-3.92 3.4l-.01.07a3.951 3.951 0 0 0 .79 2.989c.383.495.64.968.64 1.442V12.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-1.599c0-.474.257-.947.64-1.442c.63-.814.948-1.875.79-2.99l-.01-.07A3.959 3.959 0 0 0 7.5 3ZM0 8h2V7H0v1Zm13 0h2V7h-2v1Zm-7 7h3v-1H6v1Z"/>`
	buttonOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M5.5 10V8.5m0 0v-5a1 1 0 0 1 2 0v4h3.353c.91 0 1.647.737 1.647 1.647V10A4.5 4.5 0 0 1 8 14.5h-.5a4 4 0 0 1-4-4a2 2 0 0 1 2-2Zm3.5-3h2a2.5 2.5 0 0 0 0-5H4a2.5 2.5 0 0 0 0 5"/>`
	buttonSolidInnerSVG                      = `<path fill="currentColor" d="M4 6a3 3 0 0 1 0-6h7a3 3 0 1 1 0 6H9V3.5a2.5 2.5 0 0 0-5 0V6Z"/><path fill="currentColor" d="M6.5 2A1.5 1.5 0 0 0 5 3.5v4.55a2.5 2.5 0 0 0-2 2.45A4.5 4.5 0 0 0 7.5 15H8a5 5 0 0 0 5-5v-.853A2.147 2.147 0 0 0 10.853 7H8V3.5A1.5 1.5 0 0 0 6.5 2Z"/>`
	cOutlineInnerSVG                         = `<path fill="none" stroke="currentColor" d="m10 5.5l-.068-.068a3.182 3.182 0 0 0-2.25-.932H7.5a3 3 0 0 0 0 6h.182c.844 0 1.653-.335 2.25-.932L10 9.5m-8.5 1v-6l6-3.5l6 3.5v6l-6 3.5l-6-3.5Z"/>`
	cSharpOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M9.5 5v5m2-5v5M8 6.5h5m-5 2h5M8 10l-.402.201A2.831 2.831 0 0 1 3.5 7.668v-.336a2.832 2.832 0 0 1 4.098-2.533L8 5m-6.5 5.5v-6l6-3.5l6 3.5v6l-6 3.5l-6-3.5Z"/>`
	cSharpSolidInnerSVG                      = `<path fill="currentColor" d="M10 8V7h1v1h-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5.421L14 4.213v6.574L7.5 14.58L1 10.787V4.213L7.5.42ZM6.332 4A3.332 3.332 0 0 0 3 7.332v.336a3.332 3.332 0 0 0 4.821 2.98l.403-.2l-.448-.895l-.402.2A2.331 2.331 0 0 1 4 7.669v-.336a2.332 2.332 0 0 1 3.374-2.086l.402.201l.448-.894l-.403-.201A3.332 3.332 0 0 0 6.331 4ZM9 5v1H8v1h1v1H8v1h1v1h1V9h1v1h1V9h1V8h-1V7h1V6h-1V5h-1v1h-1V5H9Z" clip-rule="evenodd"/>`
	cSolidInnerSVG                           = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.421L14 4.213v6.574L7.5 14.58L1 10.787V4.213L7.5.42ZM7.5 4a3.5 3.5 0 1 0 0 7h.182c.976 0 1.913-.388 2.604-1.078l.068-.068l-.708-.708l-.068.068A2.682 2.682 0 0 1 7.682 10H7.5a2.5 2.5 0 0 1 0-5h.182c.711 0 1.393.283 1.896.786l.068.068l.708-.708l-.068-.068A3.682 3.682 0 0 0 7.682 4H7.5Z" clip-rule="evenodd"/>`
	calculatorOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M4 4.5h7m-7 4h1m2 0h1m2 0h1m-7 3h1m2 0h1m2 0h1m-8.5 3h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`
	calculatorSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h10A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM4 5h7V4H4v1Zm0 4h1V8H4v1Zm4 0H7V8h1v1Zm2 0h1V8h-1v1Zm-5 3H4v-1h1v1Zm2 0h1v-1H7v1Zm4 0h-1v-1h1v1Z" clip-rule="evenodd"/>`
	calendarMinusOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5M5 8.5h5m-8.5-6h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`
	calendarMinusSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2ZM5 9h5V8H5v1Z" clip-rule="evenodd"/>`
	calendarNoAccessOutlineInnerSVG          = `<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-2 1.5l-4 4m-4-8h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Zm6 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`
	calendarNoAccessSolidInnerSVG            = `<path fill="currentColor" d="M7.5 6a2.5 2.5 0 0 0-2.086 3.879L8.88 6.414A2.488 2.488 0 0 0 7.5 6Zm0 5c-.51 0-.983-.152-1.379-.414L9.586 7.12A2.5 2.5 0 0 1 7.5 11Z"/><path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2ZM4 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/>`
	calendarOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-10-2.5h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`
	calendarPlusOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-4 1v5M5 8.5h5m-8.5-6h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`
	calendarPlusSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2ZM7 11V9H5V8h2V6h1v2h2v1H8v2H7Z" clip-rule="evenodd"/>`
	calendarSolidInnerSVG                    = `<path fill="currentColor" d="M12 2h1.5A1.5 1.5 0 0 1 15 3.5v10a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-10A1.5 1.5 0 0 1 1.5 2H3V0h1v2h7V0h1v2Z"/>`
	calendarTickOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5M5 8.5l2 2l3.5-4m-9-4h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`
	calendarTickSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2Zm-6.476 9.232l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708l2.378 2.378Z" clip-rule="evenodd"/>`
	calendarXOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-6 1.5l4 4m-4 0l4-4m-8-4h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`
	calendarXSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2Zm-4.354 8.854L7.5 9.207l-1.646 1.646l-.708-.707L6.793 8.5L5.146 6.854l.708-.708L7.5 7.793l1.646-1.647l.708.708L8.207 8.5l1.647 1.646l-.708.707Z" clip-rule="evenodd"/>`
	cameraOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M7 1.5H2m12.5 11v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1Zm-5-2a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`
	cameraSolidInnerSVG                      = `<path fill="currentColor" d="M2 1h5v1H2V1Zm6 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 12.5A1.5 1.5 0 0 0 1.5 14h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 3h-12A1.5 1.5 0 0 0 0 4.5v8ZM9.5 6a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z" clip-rule="evenodd"/>`
	candleChartOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M.5 0v14.5H15M8.5 0v3.5m-5 6V12m0-8v1.5m10-1.5v2.5m0 4V13m-11-7.5h2v4h-2v-4Zm5-2h2v4h-2v-4Zm5 3h2v4h-2v-4Z"/>`
	candleChartSolidInnerSVG                 = `<path fill="currentColor" d="M1 0H0v15h15v-1H1V0Z"/><path fill="currentColor" d="M8 0v3H7v5h3V3H9V0H8ZM3 4v1H2v5h1v2h1v-2h1V5H4V4H3Zm9 2h1V4h1v2h1v5h-1v2h-1v-2h-1V6Z"/>`
	carOutlineInnerSVG                       = `<path fill="currentColor" d="M14.5 6.497h.5v-.139l-.071-.119l-.429.258Zm-14 0l-.429-.258L0 6.36v.138h.5Zm2.126-3.541l-.429-.258l.429.258Zm9.748 0l.429-.258l-.429.258ZM3.5 11.5V11H3v.5h.5Zm8 0h.5V11h-.5v.5ZM14 6.497V12.5h1V6.497h-1ZM.929 6.754l2.126-3.54l-.858-.516L.071 6.24l.858.515ZM5.198 2h4.604V1H5.198v1Zm6.747 1.213l2.126 3.541l.858-.515l-2.126-3.54l-.858.514ZM2.5 13h-1v1h1v-1Zm.5-1.5v1h1v-1H3ZM13.5 13h-1v1h1v-1Zm-1.5-.5v-1h-1v1h1Zm-.5-1.5h-8v1h8v-1ZM1 12.5V6.497H0V12.5h1Zm11.5.5a.5.5 0 0 1-.5-.5h-1a1.5 1.5 0 0 0 1.5 1.5v-1Zm-10 1A1.5 1.5 0 0 0 4 12.5H3a.5.5 0 0 1-.5.5v1Zm-1-1a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 14v-1ZM9.802 2a2.5 2.5 0 0 1 2.143 1.213l.858-.515A3.5 3.5 0 0 0 9.802 1v1ZM3.055 3.213A2.5 2.5 0 0 1 5.198 2V1a3.5 3.5 0 0 0-3 1.698l.857.515ZM14 12.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM2 10h3V9H2v1Zm11-1h-3v1h3V9ZM3 7h9V6H3v1Z"/>`
	carSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M2.197 2.698A3.5 3.5 0 0 1 5.198 1h4.604a3.5 3.5 0 0 1 3 1.698L15 6.358V12.5a1.5 1.5 0 0 1-1.5 1.5h-1a1.5 1.5 0 0 1-1.5-1.5V12H4v.5A1.5 1.5 0 0 1 2.5 14h-1A1.5 1.5 0 0 1 0 12.5V6.358l2.197-3.66ZM12 7H3V6h9v1ZM2 10h3V9H2v1Zm11-1h-3v1h3V9Z" clip-rule="evenodd"/>`
	caretVerticalCircleOutlineInnerSVG       = `<path fill="currentColor" d="m10.312 9.39l.39-.312l-.624-.78l-.39.312l.624.78ZM7.5 11l-.312.39l.312.25l.312-.25L7.5 11ZM5.312 8.61l-.39-.313l-.625.781l.39.312l.625-.78Zm-.624-3l-.39.312l.624.78l.39-.312l-.624-.78ZM7.5 4l.312-.39l-.312-.25l-.312.25L7.5 4Zm2.188 2.39l.39.313l.625-.781l-.39-.312l-.625.78Zm0 2.22l-2.5 2l.624.78l2.5-2l-.624-.78Zm-1.876 2l-2.5-2l-.624.78l2.5 2l.624-.78Zm-2.5-4.22l2.5-2l-.624-.78l-2.5 2l.624.78Zm1.876-2l2.5 2l.624-.78l-2.5-2l-.624.78ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	caretVerticalCircleSolidInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM11 6L7.5 3.375L4 6h7ZM4 9l3.5 2.625L11 9H4Z" clip-rule="evenodd"/>`
	caretVerticalOutlineInnerSVG             = `<path fill="currentColor" d="m7.5 13l-.332.374l.332.295l.332-.295L7.5 13Zm0-11l.34-.367l-.333-.308l-.34.301L7.5 2Zm.332 11.374l4.5-4l-.664-.748l-4.5 4l.664.748Zm0-.748l-4.5-4l-.664.748l4.5 4l.664-.748Zm-.664-11l-4.5 4l.664.748l4.5-4l-.664-.748Zm-.008.74l4.313 4l.68-.733l-4.313-4l-.68.734Z"/>`
	caretVerticalSmallOutlineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m10 9l-2.5 2L5 9m0-3l2.5-2L10 6"/>`
	caretVerticalSmallSolidInnerSVG          = `<path fill="currentColor" d="M7.5 3.375L11 6H4l3.5-2.625Zm0 8.25L4 9h7l-3.5 2.625Z"/>`
	caretVerticalSolidInnerSVG               = `<path fill="currentColor" d="M7.5 1.336L2.17 6h10.66L7.5 1.336Zm0 12.328L12.83 9H2.17l5.33 4.664Z"/>`
	cartMinusOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="m.5.5l.6 2m0 0l2.4 8h11v-6a2 2 0 0 0-2-2H1.1Zm4.9 4h5m1.5 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-8-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`
	cartMinusSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M.021.644L.979.356L1.472 2H12.5A2.5 2.5 0 0 1 15 4.5V11H3.128L.02.644ZM6 7h5V6H6v1Zm-2 6.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm5.5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`
	cartOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="m.5.5l.6 2m0 0l2.4 8h11v-6a2 2 0 0 0-2-2H1.1Zm11.4 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-8-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`
	cartPlusOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="m.5.5l.6 2m0 0l2.4 8h11v-6a2 2 0 0 0-2-2H1.1ZM8.5 4v5M6 6.5h5m1.5 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-8-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`
	cartPlusSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M.021.644L.979.356L1.472 2H12.5A2.5 2.5 0 0 1 15 4.5V11H3.128L.02.644ZM8 9V7H6V6h2V4h1v2h2v1H9v2H8Zm-4 4.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm5.5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`
	cartSolidInnerSVG                        = `<path fill="currentColor" d="M.979.356L.02.644L3.128 11H15V4.5A2.5 2.5 0 0 0 12.5 2H1.472L.979.356Z"/><path fill="currentColor" fill-rule="evenodd" d="M5.5 12a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM5 13.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm7.5-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-.5 1.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z" clip-rule="evenodd"/>`
	certificateOutlineInnerSVG               = `<path fill="currentColor" d="M9.5 14.5H9a.5.5 0 0 0 .8.4l-.3-.4Zm2-1.5l.3-.4a.5.5 0 0 0-.6 0l.3.4Zm2 1.5l-.3.4a.5.5 0 0 0 .8-.4h-.5Zm-2-3.5A2.5 2.5 0 0 1 9 8.5H8a3.5 3.5 0 0 0 3.5 3.5v-1ZM14 8.5a2.5 2.5 0 0 1-2.5 2.5v1A3.5 3.5 0 0 0 15 8.5h-1ZM11.5 6A2.5 2.5 0 0 1 14 8.5h1A3.5 3.5 0 0 0 11.5 5v1Zm0-1A3.5 3.5 0 0 0 8 8.5h1A2.5 2.5 0 0 1 11.5 6V5ZM9 10.5v4h1v-4H9Zm.8 4.4l2-1.5l-.6-.8l-2 1.5l.6.8Zm1.4-1.5l2 1.5l.6-.8l-2-1.5l-.6.8Zm2.8 1.1v-4h-1v4h1ZM15 5V1.5h-1V5h1Zm-1.5-5h-12v1h12V0ZM0 1.5v12h1v-12H0ZM1.5 15H8v-1H1.5v1ZM0 13.5A1.5 1.5 0 0 0 1.5 15v-1a.5.5 0 0 1-.5-.5H0ZM1.5 0A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5V0ZM15 1.5A1.5 1.5 0 0 0 13.5 0v1a.5.5 0 0 1 .5.5h1ZM3 5h5V4H3v1Zm0 3h3V7H3v1Z"/>`
	certificateSolidInnerSVG                 = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v4l-.8-.601a4.5 4.5 0 0 0-6.3 6.3l.1.134V15H1.5A1.5 1.5 0 0 1 0 13.5v-12ZM8 5H3V4h5v1ZM3 8h3V7H3v1Z"/><path d="M11.5 5A3.5 3.5 0 0 0 9 10.95v3.55a.5.5 0 0 0 .8.4l1.7-1.275l1.7 1.275a.5.5 0 0 0 .8-.4v-3.55A3.5 3.5 0 0 0 11.5 5ZM10 13.5v-1.837c.455.216.963.337 1.5.337s1.045-.12 1.5-.337V13.5l-1.2-.9a.5.5 0 0 0-.6 0l-1.2.9Z"/></g>`
	chatOutlineInnerSVG                      = `<path fill="currentColor" d="m11.5 13.5l.157-.475l-.218-.072l-.197.119l.258.428Zm2-2l-.421-.27l-.129.202l.076.226l.474-.158Zm1 2.99l-.157.476a.5.5 0 0 0 .631-.634l-.474.159Zm-3.258-1.418c-.956.575-2.485.919-3.742.919v1c1.385 0 3.106-.37 4.258-1.063l-.516-.856ZM7.5 13.99c-3.59 0-6.5-2.909-6.5-6.496H0a7.498 7.498 0 0 0 7.5 7.496v-1ZM1 7.495A6.498 6.498 0 0 1 7.5 1V0A7.498 7.498 0 0 0 0 7.495h1ZM7.5 1C11.09 1 14 3.908 14 7.495h1A7.498 7.498 0 0 0 7.5 0v1ZM14 7.495c0 1.331-.296 2.758-.921 3.735l.842.54C14.686 10.575 15 8.937 15 7.495h-1Zm-2.657 6.48l3 .99l.314-.949l-3-.99l-.314.949Zm3.631.357l-1-2.99l-.948.316l1 2.991l.948-.317Z"/>`
	chatSolidInnerSVG                        = `<path fill="currentColor" d="M7.5 0A7.498 7.498 0 0 0 0 7.495a7.498 7.498 0 0 0 7.5 7.496c1.306 0 2.91-.328 4.054-.947l2.79.922a.5.5 0 0 0 .63-.634l-.926-2.771c.672-1.173.952-2.706.952-4.066A7.498 7.498 0 0 0 7.5 0Z"/>`
	chatTypingAltOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M7 7.5h1m-4 0h1m5 0h1m3.5 7h-7a7 7 0 1 1 7-7v7Z"/>`
	chatTypingAltSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0V15H7.5A7.5 7.5 0 0 1 0 7.5ZM4 8h1V7H4v1Zm7 0h-1V7h1v1ZM7 8h1V7H7v1Z" clip-rule="evenodd"/>`
	chatTypingOutlineInnerSVG                = `<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222v.5Zm2 2.998l-.416.277a.5.5 0 0 0 .832 0l-.416-.277Zm2-2.998v-.5a.5.5 0 0 0-.416.222l.416.278Zm-4.416.277l2 2.998l.832-.555l-2-2.998l-.832.555Zm2.832 2.998l2-2.998l-.832-.555l-2 2.998l.832.555ZM9.5 11.993h4v-1h-4v1Zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5v1Zm1.5-1.5V1.5h-1v8.994h1ZM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5h1ZM13.5 0h-12v1h12V0Zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5V0ZM0 1.5v8.993h1V1.5H0Zm0 8.993c0 .83.671 1.5 1.5 1.5v-1a.499.499 0 0 1-.5-.5H0Zm1.5 1.5h4v-1h-4v1ZM7 7h1V6H7v1ZM4 7h1V6H4v1Zm6 0h1V6h-1v1Z"/>`
	chatTypingSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.993c0 .83-.671 1.5-1.5 1.5H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5V1.5ZM4 7h1V6H4v1Zm3 0h1V6H7v1Zm4 0h-1V6h1v1Z" clip-rule="evenodd"/>`
	chatbotOutlineInnerSVG                   = `<path fill="currentColor" d="M9 2.5V2v.5Zm-3 0V3v-.5Zm6.856 9.422l-.35-.356l-.205.2l.07.277l.485-.12ZM13.5 14.5l-.121.485a.5.5 0 0 0 .606-.606l-.485.12Zm-4-1l-.354-.354l-.624.625l.857.214l.121-.485Zm.025-.025l.353.354a.5.5 0 0 0-.4-.852l.047.498ZM.5 8H0h.5ZM7 0v2.5h1V0H7Zm2 2H6v1h3V2Zm6 6a6 6 0 0 0-6-6v1a5 5 0 0 1 5 5h1Zm-1.794 4.279A5.983 5.983 0 0 0 15 7.999h-1a4.983 4.983 0 0 1-1.495 3.567l.701.713Zm.78 2.1L13.34 11.8l-.97.242l.644 2.578l.97-.242Zm-4.607-.394l4 1l.242-.97l-4-1l-.242.97Zm-.208-.863l-.025.024l.708.707l.024-.024l-.707-.707ZM9 14c.193 0 .384-.01.572-.027l-.094-.996A5.058 5.058 0 0 1 9 13v1Zm-3 0h3v-1H6v1ZM0 8a6 6 0 0 0 6 6v-1a5 5 0 0 1-5-5H0Zm6-6a6 6 0 0 0-6 6h1a5 5 0 0 1 5-5V2Zm1.5 6A1.5 1.5 0 0 1 6 6.5H5A2.5 2.5 0 0 0 7.5 9V8ZM9 6.5A1.5 1.5 0 0 1 7.5 8v1A2.5 2.5 0 0 0 10 6.5H9ZM7.5 5A1.5 1.5 0 0 1 9 6.5h1A2.5 2.5 0 0 0 7.5 4v1Zm0-1A2.5 2.5 0 0 0 5 6.5h1A1.5 1.5 0 0 1 7.5 5V4Zm0 8c1.064 0 2.042-.37 2.813-.987l-.626-.78c-.6.48-1.359.767-2.187.767v1Zm-2.813-.987c.77.617 1.75.987 2.813.987v-1a3.483 3.483 0 0 1-2.187-.767l-.626.78Z"/>`
	chatbotSolidInnerSVG                     = `<path fill="currentColor" d="M7.5 5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/><path fill="currentColor" fill-rule="evenodd" d="M9 2H8V0H7v2H6a6 6 0 0 0 0 12h3c.13 0 .26-.004.389-.013l3.99.998a.5.5 0 0 0 .606-.606l-.577-2.309A6 6 0 0 0 9 2ZM5 6.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0ZM7.5 12a4.483 4.483 0 0 1-2.813-.987l.626-.78c.599.48 1.359.767 2.187.767c.828 0 1.588-.287 2.187-.767l.626.78A4.483 4.483 0 0 1 7.5 12Z" clip-rule="evenodd"/>`
	chromeOutlineInnerSVG                    = `<path fill="currentColor" d="M13.5 5a.5.5 0 0 0 0-1v1ZM2.964 2.814a.5.5 0 1 0-.928.372l.928-.372ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0 4h6V4h-6v1ZM4.964 7.814l-2-5l-.928.372l2 5l.928-.372ZM9.25 9.301l-3.65 4.9l.802.598l3.65-4.9l-.802-.598ZM7.5 10A2.5 2.5 0 0 1 5 7.5H4A3.5 3.5 0 0 0 7.5 11v-1ZM10 7.5A2.5 2.5 0 0 1 7.5 10v1A3.5 3.5 0 0 0 11 7.5h-1ZM7.5 5A2.5 2.5 0 0 1 10 7.5h1A3.5 3.5 0 0 0 7.5 4v1Zm0-1A3.5 3.5 0 0 0 4 7.5h1A2.5 2.5 0 0 1 7.5 5V4Z"/>`
	chromeSolidInnerSVG                      = `<path fill="currentColor" d="M2.503 1.907A7.472 7.472 0 0 1 7.5 0a7.498 7.498 0 0 1 6.635 4H7.5a3.501 3.501 0 0 0-3.23 2.149L2.503 1.907Z"/><path fill="currentColor" d="M1.745 2.69a7.503 7.503 0 0 0 3.41 11.937l2.812-3.658a3.501 3.501 0 0 1-3.88-2.685a.502.502 0 0 1-.049-.092L1.745 2.69Z"/><path fill="currentColor" d="M6.215 14.89a7.5 7.5 0 0 0 8.357-9.895A.503.503 0 0 1 14.5 5H9.95A3.49 3.49 0 0 1 11 7.5a3.487 3.487 0 0 1-.954 2.405l-3.83 4.985Z"/><path fill="currentColor" d="M5 7.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Z"/>`
	churchOutlineInnerSVG                    = `<path fill="currentColor" d="M3 8.5a.5.5 0 0 0-1 0h1Zm10 0a.5.5 0 0 0-1 0h1Zm-7.5 2V10H5v.5h.5Zm4 0h.5V10h-.5v.5ZM0 15h15v-1H0v1Zm7.252-9.934l-7 4l.496.868l7-4l-.496-.868Zm7.496 4l-7-4l-.496.868l7 4l.496-.868ZM7 0v2.5h1V0H7Zm0 2.5v3h1v-3H7ZM5 3h2.5V2H5v1Zm2.5 0H10V2H7.5v1ZM2 8.5v6h1v-6H2Zm10 0v6h1v-6h-1Zm-6 6v-4H5v4h1ZM5.5 11h4v-1h-4v1Zm3.5-.5v4h1v-4H9Z"/>`
	churchSolidInnerSVG                      = `<path fill="currentColor" d="M7 2V0h1v2h2v1H8v2.21l6.748 3.856l-.496.868L13 9.22V14h2v1h-5v-5H5v5H0v-1h2V9.219l-1.252.715l-.496-.868L7 5.21V3H5V2h2Z"/><path fill="currentColor" d="M6 15h3v-4H6v4Z"/>`
	circleOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M.5 7.5a7 7 0 1 0 14 0a7 7 0 0 0-14 0Z"/>`
	circleSolidInnerSVG                      = `<path fill="currentColor" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15Z"/>`
	clipboardMinusOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m1 7h5M4.5.5h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2Z"/>`
	clipboardMinusSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3V0Zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V1Zm0 8h5V8H5v1Z" clip-rule="evenodd"/>`
	clipboardNoAccessOutlineInnerSVG         = `<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m5.5 5l-4 4m-1-10h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2Zm3 11a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`
	clipboardNoAccessSolidInnerSVG           = `<path fill="currentColor" d="M7.5 6a2.5 2.5 0 0 0-2.086 3.879L8.88 6.414A2.488 2.488 0 0 0 7.5 6Zm0 5c-.51 0-.983-.152-1.379-.414L9.586 7.12A2.5 2.5 0 0 1 7.5 11Z"/><path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3V0Zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V1ZM4 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/>`
	clipboardOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m.5-1h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2Z"/>`
	clipboardPlusOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4M7.5 6v5M5 8.5h5M4.5.5h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2Z"/>`
	clipboardPlusSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3V0Zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V1Zm2 10V9H5V8h2V6h1v2h2v1H8v2H7Z" clip-rule="evenodd"/>`
	clipboardSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M11 0H4v1H1v12.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V1h-3V0Zm-1 1H5v1.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V1Z" clip-rule="evenodd"/>`
	clipboardTickOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m1 7l2 2l3.5-4m-6-6h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2Z"/>`
	clipboardTickSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3V0Zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V1Zm2.024 10.232l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708l2.378 2.378Z" clip-rule="evenodd"/>`
	clipboardXOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m1.5 5l4 4m-4 0l4-4m-5-6h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2Z"/>`
	clipboardXSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3V0Zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V1Zm4.146 9.854L7.5 9.207l-1.646 1.646l-.708-.707L6.793 8.5L5.146 6.854l.708-.708L7.5 7.793l1.646-1.647l.708.708L8.207 8.5l1.647 1.646l-.708.707Z" clip-rule="evenodd"/>`
	clockOutlineInnerSVG                     = `<path fill="currentColor" d="M7.5 7.5H7a.5.5 0 0 0 .146.354L7.5 7.5Zm0 6.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM7 3v4.5h1V3H7Zm.146 4.854l3 3l.708-.708l-3-3l-.708.708Z"/>`
	clockSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm7 0V3h1v4.293l2.854 2.853l-.708.708l-3-3A.499.499 0 0 1 7 7.5Z" clip-rule="evenodd"/>`
	clockwiseOutlineInnerSVG                 = `<path fill="currentColor" d="M14 8.495v-.5h-1v.5h1ZM7.5 2.999H8v-1h-.5v1Zm1-.5l.353.353l.354-.353l-.354-.354l-.353.354ZM13 8.495a5.499 5.499 0 0 1-5.5 5.496v1c3.589 0 6.5-2.909 6.5-6.496h-1ZM7.5 13.99A5.499 5.499 0 0 1 2 8.495H1a6.499 6.499 0 0 0 6.5 6.496v-1ZM2 8.495a5.499 5.499 0 0 1 5.5-5.496v-1A6.499 6.499 0 0 0 1 8.495h1ZM6.147.854l2 1.998l.706-.707l-2-1.999l-.706.708Zm2 1.291l-2 1.999l.706.707l2-1.999l-.706-.707Z"/>`
	clockwiseSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M7.295 2.002L6.147.854l.706-.708L9.207 2.5L6.853 4.85l-.706-.707l1.141-1.141A5.499 5.499 0 0 0 2 8.495a5.499 5.499 0 0 0 5.5 5.496A5.5 5.5 0 0 0 13 8.495v-.5h1v.5a6.499 6.499 0 0 1-6.5 6.496A6.499 6.499 0 0 1 1 8.495a6.499 6.499 0 0 1 6.295-6.493Z" clip-rule="evenodd"/>`
	codeOutlineInnerSVG                      = `<path fill="currentColor" d="m10.146 10.146l-.353.354l.707.707l.354-.353l-.708-.708ZM13.5 7.5l.354.354l.353-.354l-.353-.354l-.354.354Zm-2.646-3.354l-.354-.353l-.707.707l.353.354l.708-.708Zm-6.708 6.708l.354.353l.707-.707l-.353-.354l-.708.708ZM1.5 7.5l-.354-.354l-.353.354l.353.354L1.5 7.5Zm3.354-2.646l.353-.354l-.707-.707l-.354.353l.708.708Zm6 6l3-3l-.708-.708l-3 3l.708.708Zm3-3.708l-3-3l-.708.708l3 3l.708-.708Zm-9 3l-3-3l-.708.708l3 3l.708-.708Zm-3-2.292l3-3l-.708-.708l-3 3l.708.708Zm6.153-6.436l-2 12l.986.164l2-12l-.986-.164Z"/>`
	codeSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="m6.007 13.418l2-12l.986.164l-2 12l-.986-.164Zm-.8-8.918l-3 3l3 3l-.707.707L.793 7.5L4.5 3.793l.707.707Zm5.293-.707L14.207 7.5L10.5 11.207l-.707-.707l3-3l-3-3l.707-.707Z" clip-rule="evenodd"/>`
	codepenOutlineInnerSVG                   = `<path fill="currentColor" d="m7.5.5l.29-.407a.5.5 0 0 0-.58 0L7.5.5Zm7 5h.5a.5.5 0 0 0-.21-.407l-.29.407Zm0 4l.29.407A.5.5 0 0 0 15 9.5h-.5Zm-7 5l-.29.407a.5.5 0 0 0 .58 0L7.5 14.5Zm-7-5H0a.5.5 0 0 0 .21.407L.5 9.5Zm0-4l-.29-.407A.5.5 0 0 0 0 5.5h.5ZM7.21.907l7 5l.58-.814l-7-5l-.58.814ZM14 5.5v4h1v-4h-1Zm.21 3.593l-7 5l.58.814l7-5l-.58-.814Zm-6.42 5l-7-5l-.58.814l7 5l.58-.814ZM1 9.5v-4H0v4h1ZM.79 5.907l7-5l-.58-.814l-7 5l.58.814Zm0 4l7-5l-.58-.814l-7 5l.58.814Zm6.42-5l7 5l.58-.814l-7-5l-.58.814Zm-7 1l7 5l.58-.814l-7-5l-.58.814Zm7.58 5l7-5l-.58-.814l-7 5l.58.814ZM7 .5v4h1v-4H7Zm0 10v4h1v-4H7Z"/>`
	codepenSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.21.093a.5.5 0 0 1 .58 0l7 5A.5.5 0 0 1 15 5.5v4a.5.5 0 0 1-.21.407l-7 5a.5.5 0 0 1-.58 0l-7-5A.5.5 0 0 1 0 9.5v-4a.5.5 0 0 1 .21-.407l7-5ZM1 6.472L2.44 7.5L1 8.528V6.472ZM1.36 9.5L7 13.528v-2.77L3.3 8.113L1.36 9.5Zm2.8-2L7.5 9.886L10.84 7.5L7.5 5.114L4.16 7.5ZM8 4.243l3.7 2.643L13.64 5.5L8 1.472v2.77ZM7 1.472v2.77L3.3 6.887L1.36 5.5L7 1.472Zm7 5L12.56 7.5L14 8.528V6.472ZM13.64 9.5L11.7 8.114L8 10.757v2.771L13.64 9.5Z" clip-rule="evenodd"/>`
	cogOutlineInnerSVG                       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" clip-rule="evenodd"><path d="m5.944.5l-.086.437l-.329 1.598a5.52 5.52 0 0 0-1.434.823L2.487 2.82l-.432-.133l-.224.385L.724 4.923L.5 5.31l.328.287l1.244 1.058c-.045.277-.103.55-.103.841c0 .291.058.565.103.842L.828 9.395L.5 9.682l.224.386l1.107 1.85l.224.387l.432-.135l1.608-.537c.431.338.908.622 1.434.823l.329 1.598l.086.437h3.111l.087-.437l.328-1.598a5.524 5.524 0 0 0 1.434-.823l1.608.537l.432.135l.225-.386l1.106-1.851l.225-.386l-.329-.287l-1.244-1.058c.046-.277.103-.55.103-.842c0-.29-.057-.564-.103-.841l1.244-1.058l.329-.287l-.225-.386l-1.106-1.85l-.225-.386l-.432.134l-1.608.537a5.52 5.52 0 0 0-1.434-.823L9.142.937L9.055.5H5.944Z"/><path d="M9.5 7.495a2 2 0 0 1-4 0a2 2 0 0 1 4 0Z"/></g>`
	cogSolidInnerSVG                         = `<path fill="currentColor" d="M6.026 7.5c0-.827.66-1.5 1.473-1.5c.813 0 1.473.673 1.473 1.5S8.312 9 7.499 9c-.813 0-1.473-.673-1.473-1.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M5.568 0h3.86l.164.837V.84l.27 1.335c.383.172.74.386 1.068.627l1.345-.458l.796-.251l.417.727l1.087 1.854l.425.743l-.633.563l-1.009.874c.032.197.061.418.061.646c0 .228-.03.45-.06.646l1.012.878l.629.559l-.428.748l-1.084 1.849l-.417.73l-.807-.258l-1.334-.454a5.877 5.877 0 0 1-1.068.627l-.27 1.335v.003L9.43 15H5.57l-.163-.839l-.27-1.336a5.877 5.877 0 0 1-1.068-.627l-1.343.457l-.799.255l-.415-.73l-1.088-1.855L0 9.583l.632-.563l1.008-.875a4.075 4.075 0 0 1-.06-.645c0-.227.03-.45.061-.645l-1.014-.88L0 5.418l.426-.748l1.085-1.85l.415-.728l.808.256l1.334.454a5.875 5.875 0 0 1 1.068-.627l.27-1.337L5.568 0ZM7.5 5C6.144 5 5.045 6.12 5.045 7.5s1.1 2.5 2.454 2.5c1.355 0 2.454-1.12 2.454-2.5S8.853 5 7.5 5Z" clip-rule="evenodd"/>`
	compassOutlineInnerSVG                   = `<path fill="currentColor" d="m4.5 10.5l-.447-.224a.5.5 0 0 0 .67.671L4.5 10.5Zm2-4l-.224-.447a.5.5 0 0 0-.223.223L6.5 6.5Zm4-2l.447.224a.5.5 0 0 0-.67-.671l.223.447Zm-2 4l.224.447a.5.5 0 0 0 .223-.223L8.5 8.5Zm-1 5.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM4.947 10.724l2-4l-.894-.448l-2 4l.894.448Zm1.777-3.777l4-2l-.448-.894l-4 2l.448.894Zm3.329-2.67l-2 4l.894.447l2-4l-.894-.448ZM8.276 8.052l-4 2l.448.894l4-2l-.448-.894Z"/>`
	compassSolidInnerSVG                     = `<path fill="currentColor" d="m5.618 9.382l1.255-2.51l2.509-1.254l-1.255 2.51l-2.509 1.254Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm10.947-2.776a.5.5 0 0 0-.67-.671l-4 2a.5.5 0 0 0-.224.223l-2 4a.5.5 0 0 0 .67.671l4-2a.5.5 0 0 0 .224-.223l2-4Z" clip-rule="evenodd"/>`
	computerOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M9.5 14.5h4a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Zm0 0H4m2.5-3v3m2-9h6m-4.5 6h3m-11.5-8h7v8h-7a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`
	computerSolidInnerSVG                    = `<path fill="currentColor" d="M10 12h3v-1h-3v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M9.5 0A1.5 1.5 0 0 0 8 1.5V3H1.5A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12H6v2H4v1h9.5a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0h-4ZM8.085 14H7v-2h1v1.5c0 .175.03.344.085.5ZM9.5 14h4a.5.5 0 0 0 .5-.5V6H9v7.5a.5.5 0 0 0 .5.5ZM9 5h5V1.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5V5Z" clip-rule="evenodd"/>`
	contactOutlineInnerSVG                   = `<path fill="currentColor" d="M2 12.5v.5h1v-.5H2Zm5 0v.5h1v-.5H7Zm-4 0V12H2v.5h1Zm4-.5v.5h1V12H7Zm-2-2a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3v1Zm-2 2a2 2 0 0 1 2-2V9a3 3 0 0 0-3 3h1Zm2-8a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1V4Zm2 2a2 2 0 0 0-2-2v1a1 1 0 0 1 1 1h1ZM5 8a2 2 0 0 0 2-2H6a1 1 0 0 1-1 1v1Zm0-1a1 1 0 0 1-1-1H3a2 2 0 0 0 2 2V7ZM1.5 3h12V2h-12v1Zm12.5.5v8h1v-8h-1Zm-.5 8.5h-12v1h12v-1ZM1 11.5v-8H0v8h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 13v-1Zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 2v1Zm-12-1A1.5 1.5 0 0 0 0 3.5h1a.5.5 0 0 1 .5-.5V2ZM9 6h3V5H9v1Zm0 3h3V8H9v1Zm-9 6h15v-1H0v1ZM3 0v2.5h1V0H3Zm8 0v2.5h1V0h-1Z"/>`
	contactSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M11 2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v8A1.5 1.5 0 0 0 1.5 13h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 2H12V0h-1v2ZM3 6a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm-.618 4.618a2.927 2.927 0 0 1 5.236 0l.33.658A.5.5 0 0 1 7.5 12h-5a.5.5 0 0 1-.447-.724l.329-.658ZM9 6h3V5H9v1Zm0 3h3V8H9v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 14v1H0v-1h15Z"/>`
	contractOutlineInnerSVG                  = `<path fill="currentColor" d="M4.5 7H4v1h.5V7Zm6 1h.5V7h-.5v1Zm-6-4H4v1h.5V4Zm2 1H7V4h-.5v1Zm4-4.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354ZM8 11l-.354.354L8 11Zm-.5.5l.224.447l.04-.02l.036-.027l-.3-.4ZM4.5 8h6V7h-6v1Zm0-3h2V4h-2v1Zm8 9h-10v1h10v-1ZM2 13.5v-12H1v12h1ZM2.5 1h8V0h-8v1ZM13 3.5v10h1v-10h-1ZM10.146.854l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1Zm3.474 10.158c.111-.333.427-.642.83-.75c.377-.101.862-.035 1.342.446l.708-.708c-.72-.72-1.569-.903-2.309-.704c-.713.192-1.297.733-1.52 1.4l.95.316Zm2.172-.304a.933.933 0 0 1 .079.087l.79-.614a1.945 1.945 0 0 0-.161-.18l-.708.707Zm.079.087c.078.1.06.132.063.11c.002-.014.006.009-.054.063a1 1 0 0 1-.29.169a1.781 1.781 0 0 1-.394.108a.848.848 0 0 1-.25.01c-.017-.004.018 0 .07.037a.388.388 0 0 1 .131.43a.209.209 0 0 1-.023.047c-.002.002.015-.02.072-.067c.114-.092.324-.226.674-.4l-.448-.895c-.377.188-.66.36-.854.517a1.375 1.375 0 0 0-.26.267a.705.705 0 0 0-.14.438a.61.61 0 0 0 .255.468c.113.084.238.12.33.139c.187.037.4.027.593-.002c.38-.058.872-.222 1.207-.526c.174-.159.339-.387.374-.686c.036-.306-.074-.593-.267-.84l-.79.613Zm.075.459a2.56 2.56 0 0 1 .518-.307l-.397-.918c-.24.104-.48.245-.721.425l.6.8Zm.518-.307c.65-.281 1.231-.133 1.826.15c.15.072.296.15.444.23c.144.078.296.161.44.235c.276.139.618.292.972.292v-1c-.094 0-.248-.047-.52-.184c-.128-.066-.262-.14-.416-.223c-.15-.081-.316-.17-.49-.252c-.698-.333-1.611-.616-2.653-.166l.397.918Z"/>`
	contractSolidInnerSVG                    = `<path fill="currentColor" d="M6.796 11.9H6.8h-.003Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM7 4H4v1h3V4Zm4 3H4v1h7V7Zm-4.695 3.908c-.404.108-.72.417-.83.75l-.95-.316c.223-.667.807-1.208 1.52-1.4c.707-.19 1.514-.03 2.212.611a2.75 2.75 0 0 1 .622-.107c.54-.029 1.023.107 1.438.28c.305.127.6.287.85.422c.078.044.153.084.222.12c.323.17.5.232.611.232v1c-.39 0-.774-.188-1.076-.346a21.802 21.802 0 0 1-.272-.146a7.689 7.689 0 0 0-.72-.359c-.334-.14-.663-.222-.999-.204a1.686 1.686 0 0 0-.15.014l.001.014c.027.324-.107.591-.28.783c-.318.354-.837.54-1.227.61a1.962 1.962 0 0 1-.614.025a.9.9 0 0 1-.33-.11a.623.623 0 0 1-.303-.433a.677.677 0 0 1 .111-.48a1.28 1.28 0 0 1 .262-.282c.19-.157.465-.327.834-.513l.027-.02a1.23 1.23 0 0 0-.96-.145Z" clip-rule="evenodd"/>`
	costEstimateOutlineInnerSVG              = `<path fill="currentColor" d="M4.5 7H4v1h.5V7Zm6 1h.5V7h-.5v1Zm-2 2H8v1h.5v-1Zm2 1h.5v-1h-.5v1Zm-6-7H4v1h.5V4Zm2 1H7V4h-.5v1Zm4-4.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354ZM4.5 8h6V7h-6v1Zm4 3h2v-1h-2v1Zm-4-6h2V4h-2v1Zm8 9h-10v1h10v-1ZM2 13.5v-12H1v12h1ZM2.5 1h8V0h-8v1ZM13 3.5v10h1v-10h-1ZM10.146.854l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1Z"/>`
	costEstimateSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM4 4h3v1H4V4Zm7 3H4v1h7V7Zm0 3H8v1h3v-1Z" clip-rule="evenodd"/>`
	cplusplusOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M9.5 9.5c-.64.64-1.509 1-2.414 1H6.5a3 3 0 0 1 0-6h.586c.905 0 1.774.36 2.414 1m-2 .5v3M6 7.5h6M10.5 6v3m-9 1.5v-6l6-3.5l6 3.5v6l-6 3.5l-6-3.5Z"/>`
	cplusplusSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.421L14 4.213v6.574L7.5 14.58L1 10.787V4.213L7.5.42ZM6.5 4a3.5 3.5 0 1 0 0 7h.586a3.914 3.914 0 0 0 2.768-1.146l-.708-.708a2.914 2.914 0 0 1-2.06.854H6.5a2.5 2.5 0 0 1 0-5h.586a2.91 2.91 0 0 1 2.06.854l.708-.708A3.914 3.914 0 0 0 7.086 4H6.5ZM7 7V6h1v1h2V6h1v1h1v1h-1v1h-1V8H8v1H7V8H6V7h1Z" clip-rule="evenodd"/>`
	creditCardOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M.5 5.5h14M2 9.5h6m2 0h3M.5 3.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`
	creditCardSolidInnerSVG                  = `<path fill="currentColor" d="M13.5 2A1.5 1.5 0 0 1 15 3.5V5H0V3.5A1.5 1.5 0 0 1 1.5 2h12Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6v5.5A1.5 1.5 0 0 0 1.5 13h12a1.5 1.5 0 0 0 1.5-1.5V6H0Zm2 4h6V9H2v1Zm11 0h-3V9h3v1Z" clip-rule="evenodd"/>`
	cropOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M3.5 3.5h8v8m-8-8V0m0 3.5H0m11.5 8h-8V6m8 5.5V15m0-3.5H15"/>`
	cropSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M3 3V0h1v3h8v8h3v1h-3v3h-1v-3H3V6h1v5h7V4H0V3h3Z" clip-rule="evenodd"/>`
	cssThreeOutlineInnerSVG                  = `<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.498.542L.5.5Zm14 0l.498.042A.5.5 0 0 0 14.5 0v.5Zm-1 12l.158.474a.5.5 0 0 0 .34-.433L13.5 12.5Zm-6 2l-.158.474a.499.499 0 0 0 .316 0L7.5 14.5Zm-6-2l-.498.041a.5.5 0 0 0 .34.433L1.5 12.5Zm9-9h.5V3h-.5v.5Zm0 6l.158.474L11 9.86V9.5h-.5Zm-3 1l-.158.474l.158.053l.158-.053L7.5 10.5Zm-3-1H4v.36l.342.114L4.5 9.5ZM.5 1h14V0H.5v1ZM14.002.458l-1 12l.996.083l1-12l-.996-.083Zm-.66 11.568l-6 2l.316.948l6-2l-.316-.948Zm-5.684 2l-6-2l-.316.948l6 2l.316-.948Zm-5.66-1.567l-1-12l-.996.083l1 12l.996-.083ZM10.5 3H4v1h6.5V3ZM6 7h4.5V6H6v1Zm4-.5v3h1v-3h-1Zm.342 2.526l-3 1l.316.948l3-1l-.316-.948Zm-2.684 1l-3-1l-.316.948l3 1l.316-.948ZM5 9.5V8H4v1.5h1Zm5-6v3h1v-3h-1Z"/>`
	cssThreeSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M.132.162A.5.5 0 0 1 .5 0h14a.5.5 0 0 1 .498.542l-1 11.916a.5.5 0 0 1-.34.432l-6 2a.5.5 0 0 1-.316 0l-6-2a.5.5 0 0 1-.34-.432L.002.542a.5.5 0 0 1 .13-.38ZM11 3H4v1h6v2H6v1h4v2.14l-2.5.833L5 9.14V8H4v1.86l3.5 1.167L11 9.86V3Z" clip-rule="evenodd"/>`
	csvOutlineInnerSVG                       = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-4 6V6H6v.5h.5Zm0 2H6V9h.5v-.5Zm2 0H9V8h-.5v.5Zm0 2v.5H9v-.5h-.5Zm2-1H10v.207l.146.147l.354-.354Zm1 1l-.354.354l.354.353l.354-.353l-.354-.354Zm1-1l.354.354l.146-.147V9.5h-.5Zm-10-3V6H2v.5h.5Zm0 4H2v.5h.5v-.5ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM9 6H6.5v1H9V6Zm-3 .5v2h1v-2H6ZM6.5 9h2V8h-2v1ZM8 8.5v2h1v-2H8Zm.5 1.5H6v1h2.5v-1ZM10 6v3.5h1V6h-1Zm.146 3.854l1 1l.708-.708l-1-1l-.708.708Zm1.708 1l1-1l-.708-.708l-1 1l.708.708ZM13 9.5V6h-1v3.5h1ZM5 6H2.5v1H5V6Zm-3 .5v4h1v-4H2Zm.5 4.5H5v-1H2.5v1Z"/>`
	csvSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM2 6h3v1H3v3h2v1H2V6Zm7 0H6v3h2v1H6v1h3V8H7V7h2V6Zm2 0h-1v3.707l1.5 1.5l1.5-1.5V6h-1v3.293l-.5.5l-.5-.5V6Z" clip-rule="evenodd"/>`
	cupOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M11.5 6.5v5a3 3 0 0 1-3 3h-5a3 3 0 0 1-3-3v-5a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1Zm0 0h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2M4.5 4V2m3 2V0"/>`
	cupSolidInnerSVG                         = `<path fill="currentColor" d="M7 4h1V0H7v4ZM5 2v2H4V2h1Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6.5A1.5 1.5 0 0 1 1.5 5h9a1.5 1.5 0 0 1 1.415 1H13.5A1.5 1.5 0 0 1 15 7.5v2a1.5 1.5 0 0 1-1.5 1.5H12v.5A3.5 3.5 0 0 1 8.5 15h-5A3.5 3.5 0 0 1 0 11.5v-5ZM12 10h1.5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5H12v3Z" clip-rule="evenodd"/>`
	curvedConnectorOutlineInnerSVG           = `<path fill="none" stroke="currentColor" d="M2.5 1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 0h2a3 3 0 0 1 3 3v6a3 3 0 0 0 3 3h2m0 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`
	curvedConnectorSolidInnerSVG             = `<path fill="currentColor" d="M1.5 0a1.5 1.5 0 1 0 1.415 2H4.5A2.5 2.5 0 0 1 7 4.5v6a3.5 3.5 0 0 0 3.5 3.5h1.585a1.5 1.5 0 1 0 0-1H10.5A2.5 2.5 0 0 1 8 10.5v-6A3.5 3.5 0 0 0 4.5 1H2.915A1.5 1.5 0 0 0 1.5 0Z"/>`
	dThreeOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M0 1.5h1.5a6 6 0 1 1 0 12H0m7-12h4.5a3 3 0 1 1 0 6m0 0H9m2.5 0h-2m2 0a3 3 0 1 1 0 6H7"/>`
	dThreeSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 2H0V1h1.5a6.5 6.5 0 0 1 0 13H0v-1h1.5a5.5 5.5 0 1 0 0-11Zm10 0H7V1h4.5a3.5 3.5 0 0 1 1.804 6.5A3.5 3.5 0 0 1 11.5 14H7v-1h4.5a2.5 2.5 0 0 0 0-5H9V7h2.5a2.5 2.5 0 0 0 0-5Z" clip-rule="evenodd"/>`
	databaseOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M14.5 2.499c0 1.103-3.134 1.998-7 1.998S.5 3.602.5 2.5m14 0c0-1.105-3.134-2-7-2s-7 .895-7 1.999m14 0v9.993c0 1.103-3.134 1.999-7 1.999s-7-.895-7-1.999V2.5m14 4.996c0 1.104-3.134 2-7 2s-7-.896-7-2"/>`
	databaseSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 0C5.534 0 3.736.227 2.413.605c-.658.188-1.227.42-1.643.701C.372 1.576 0 1.97 0 2.5v9.993c0 .53.372.924.77 1.192c.416.281.985.514 1.643.702c1.323.378 3.121.605 5.087.605c1.966 0 3.764-.227 5.087-.605c.658-.188 1.227-.421 1.643-.702c.398-.268.77-.662.77-1.192V2.5c0-.53-.372-.924-.77-1.193c-.416-.28-.985-.513-1.643-.701C11.264.227 9.466 0 7.5 0ZM1.262 2.864l.452.214c1.127.534 3.27.92 5.786.92c2.517 0 4.659-.386 5.786-.92l.452-.214l.428.904l-.452.214c-1.323.627-3.638 1.017-6.214 1.017c-2.576 0-4.891-.39-6.214-1.017l-.452-.215l.428-.903Zm.452 5.184l-.452-.214l-.428.904l.452.214c1.323.627 3.638 1.017 6.214 1.017c2.576 0 4.891-.39 6.214-1.017l.452-.214l-.428-.904l-.452.214c-1.127.534-3.27.92-5.786.92c-2.517 0-4.659-.386-5.786-.92Z" clip-rule="evenodd"/>`
	deniedOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m2.5 2.5l10 10m-5 2a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	deniedSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm2.564-4.23a6.5 6.5 0 0 0 9.165 9.165L2.564 3.272Zm.707-.706l9.165 9.165a6.5 6.5 0 0 0-9.165-9.165Z" clip-rule="evenodd"/>`
	denoOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M7.5 14.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm0 0v-5H6a2.5 2.5 0 0 1 0-5h1.694a3.495 3.495 0 0 1 3.49 3.301L11.5 13.5M7 9.5h.382c.685 0 1.312-.387 1.618-1m-5-2h1m1 0h1"/>`
	denoSolidInnerSVG                        = `<path fill="currentColor" d="M7 7H6V6h1v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM8 13.981a6.462 6.462 0 0 0 2.971-.985l-.287-5.167A2.995 2.995 0 0 0 7.694 5H6a2 2 0 0 0-1.732 1H5v1H4a2 2 0 0 0 2 2h.882c.496 0 .95-.28 1.17-.724l.895.448A2.308 2.308 0 0 1 8 9.71v4.27Z" clip-rule="evenodd"/>`
	depthChartOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M.5 0v14.5h14V0M.5 4.5h2v1h2v3h2v3h1v3v-2h2v-2h2v-3h1v-2h2"/>`
	depthChartSolidInnerSVG                  = `<path fill="currentColor" d="M1 0H0v15h15V0h-1v5h-2v2h-1v3H9v2H8v-1H7V8H5V5H3V4H1V0Z"/>`
	desklampOutlineInnerSVG                  = `<path fill="currentColor" d="m1.5 8.5l-.395-.307a.5.5 0 0 0 .041.66L1.5 8.5Zm5-7l-.354.354L6.5 1.5Zm2 2l.354.354l.353-.354l-.353-.354L8.5 3.5Zm-2 2l-.354.354l.354.353l.354-.353L6.5 5.5Zm-2-2l-.354.354L4.5 3.5Zm3.136.864l.354.353l-.354-.353Zm-.272.272l-.354-.353l.354.353Zm.272 5.728l-.353-.354l.353.354Zm5.728-5.728l-.354-.353l.354.353Zm0-.272l-.354.353l.354-.353Zm-6 6l.353-.354l-.353.354ZM2 15h11v-1H2v1Zm5.854-.854l-6-6l-.708.708l6 6l.708-.707Zm-5.96-5.339l3.5-4.5l-.789-.614l-3.5 4.5l.79.614Zm4.252-6.953l2 2l.708-.708l-2-2l-.708.708Zm2 1.292l-2 2l.708.708l2-2l-.708-.708Zm-1.292 2l-2-2l-.708.708l2 2l.708-.708Zm-2-3.292a.914.914 0 0 1 1.292 0l.708-.708a1.914 1.914 0 0 0-2.708 0l.708.708Zm-.708-.708a1.914 1.914 0 0 0 0 2.708l.708-.708a.914.914 0 0 1 0-1.292l-.708-.708ZM7.283 4.01l-.273.273l.707.707l.273-.273l-.707-.707Zm.707 6.707l5.727-5.727l-.707-.707l-5.727 5.727l.707.707Zm5.727-5.727c.27-.27.27-.71 0-.98l-.707.707a.307.307 0 0 1 0-.434l.707.707ZM7.01 10.717c.27.27.71.27.98 0l-.707-.707c.12-.12.314-.12.434 0l-.707.707Zm0-6.434a4.55 4.55 0 0 0 0 6.434l.707-.707a3.55 3.55 0 0 1 0-5.02l-.707-.707Zm.98.434a3.55 3.55 0 0 1 5.02 0l.707-.707a4.55 4.55 0 0 0-6.434 0l.707.707Z"/>`
	desklampSolidInnerSVG                    = `<path fill="currentColor" d="m6.854 1.146l1.884 1.885a4.551 4.551 0 0 1 4.98.98c.27.27.27.708 0 .979L7.99 10.717c-.27.27-.71.27-.98 0a4.551 4.551 0 0 1-.979-4.979l-.984-.984L2.166 8.46L7.707 14H13v1H2v-1h4.293L1.146 8.854a.5.5 0 0 1-.04-.66L4.333 4.04l-.188-.187a1.914 1.914 0 1 1 2.708-2.708Z"/>`
	diamondOutlineInnerSVG                   = `<path fill="currentColor" d="m7.5 14.5l-.395.307a.5.5 0 0 0 .79 0L7.5 14.5Zm-7-9l-.429-.257a.5.5 0 0 0 .034.564L.5 5.5Zm3-5V0h-.283L3.07.243L3.5.5Zm8 0l.429-.257L11.783 0H11.5v.5Zm3 5l.395.307a.5.5 0 0 0 .034-.564L14.5 5.5Zm-6.605 8.693l-7-9l-.79.614l7 9l.79-.614ZM.929 5.757l3-5L3.07.243l-3 5l.858.514ZM3.5 1h8V0h-8v1Zm7.571-.243l3 5l.858-.514l-3-5l-.858.514Zm3.034 4.436l-7 9l.79.614l7-9l-.79-.614Z"/>`
	diamondSolidInnerSVG                     = `<path fill="currentColor" d="M11.783 0H3.217L.07 5.243a.5.5 0 0 0 .034.564l7 9a.5.5 0 0 0 .79 0l7-9a.5.5 0 0 0 .034-.564L11.783 0Z"/>`
	directionOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m.5.5l6 14l2-6l6-2l-14-6Z"/>`
	directionSolidInnerSVG                   = `<path fill="currentColor" d="M.697.04A.5.5 0 0 0 .04.697l6 14a.5.5 0 0 0 .934-.039l1.921-5.763l5.763-1.92a.5.5 0 0 0 .039-.935l-14-6Z"/>`
	discordOutlineInnerSVG                   = `<path fill="currentColor" d="m11.5 13.5l-.326.379a.5.5 0 0 0 .342.12L11.5 13.5Zm-1.066-1.712a.5.5 0 0 0-.785.62l.785-.62Zm.398-.41l-.174-.468a.672.672 0 0 0-.02.007l.194.461Zm-1.738.516L9.01 11.4l-.008.001l.092.492Zm-3.104-.012l-.095.49l.003.001l.092-.491Zm-1.762-.515l-.182.465l.182-.466Zm-.875-.408l-.278.415a.46.46 0 0 0 .033.021l.245-.436Zm-.108-.06l.277-.416a.491.491 0 0 0-.054-.031l-.223.447Zm-.048-.036l.353-.354a.502.502 0 0 0-.11-.083l-.243.437Zm2.154 1.52a.5.5 0 0 0-.785-.62l.785.62ZM3.5 13.5l-.016.5a.5.5 0 0 0 .347-.125L3.5 13.5Zm-3-2.253H0a.5.5 0 0 0 .006.08l.494-.08Zm1.726-8.488l-.3-.4a.5.5 0 0 0-.168.225l.468.175ZM5.594 1.5l.498-.047A.5.5 0 0 0 5.605 1l-.01.5Zm-.378 1.306a.5.5 0 0 0 .996-.095l-.996.095Zm3.526-.063a.5.5 0 0 0 .992.127l-.992-.127ZM9.406 1.5L9.395 1a.5.5 0 0 0-.485.436l.496.064Zm3.368 1.259l.468-.175a.5.5 0 0 0-.168-.225l-.3.4Zm1.726 8.488l.494.08a.497.497 0 0 0 .006-.08h-.5ZM6.481 8.8l-.5-.008V8.8h.5Zm5.019 4.7l.326-.379l-.002-.002a.794.794 0 0 1-.044-.038a21.355 21.355 0 0 1-.536-.48c-.325-.298-.66-.622-.81-.813l-.785.62c.208.264.603.64.918.93a29.109 29.109 0 0 0 .593.53l.01.008l.003.002l.327-.378Zm.436-3.246c-.46.303-.894.513-1.278.656l.348.937a7.352 7.352 0 0 0 1.48-.758l-.55-.835Zm-1.297.663a7.387 7.387 0 0 1-1.629.484l.168.986a8.39 8.39 0 0 0 1.848-.548l-.387-.922Zm-1.637.485a7.895 7.895 0 0 1-2.92-.012l-.184.983a8.896 8.896 0 0 0 3.288.012l-.184-.983Zm-2.917-.011a9.57 9.57 0 0 1-1.675-.49l-.364.931c.512.2 1.13.402 1.849.54l.19-.981Zm-1.675-.49a6.536 6.536 0 0 1-.813-.378l-.489.872c.326.183.648.324.938.437l.364-.931Zm-.78-.358a.802.802 0 0 0-.108-.061c-.02-.01-.011-.007 0 .001l-.555.832a.87.87 0 0 0 .108.061c.021.01.012.007 0-.002l.556-.83Zm-.162-.091a.332.332 0 0 1 .082.058l-.707.707c.023.023.081.08.178.13l.447-.895Zm-.028-.026a4.697 4.697 0 0 1-.28-.168l-.011-.008a.025.025 0 0 0-.001 0l-.287.41l-.286.409l.001.001l.002.002l.007.004l.021.014l.075.049c.064.04.156.096.273.161l.486-.874Zm1.126 1.338c-.152.193-.489.525-.813.829a30.38 30.38 0 0 1-.538.491l-.034.031l-.01.008l-.001.002h-.001l.331.375l.331.375l.001-.001l.003-.002l.01-.009l.036-.032a38.039 38.039 0 0 0 .555-.508c.315-.296.708-.677.915-.94l-.785-.62ZM3.516 13c-1.166-.037-1.778-.521-2.11-.96a2.394 2.394 0 0 1-.4-.82a1.1 1.1 0 0 1-.013-.056v.002l-.493.08c-.494.08-.494.08-.493.081v.006a1.367 1.367 0 0 0 .028.127a3.394 3.394 0 0 0 .573 1.183c.505.667 1.393 1.31 2.876 1.357l.032-1ZM1 11.247c0-1.867.42-3.94.847-5.564a35.45 35.45 0 0 1 .776-2.552a16.43 16.43 0 0 1 .067-.186l.004-.01v-.001l-.468-.175l-.469-.175v.001l-.001.003l-.004.011a9.393 9.393 0 0 0-.072.2a36.445 36.445 0 0 0-.8 2.629C.443 7.083 0 9.253 0 11.247h1Zm1.526-8.088c.8-.6 1.577-.89 2.15-1.03a4.764 4.764 0 0 1 .86-.128A1.48 1.48 0 0 1 5.585 2h-.001l.01-.5l.01-.5H5.6a.848.848 0 0 0-.028 0h-.068a3.973 3.973 0 0 0-.24.016a5.763 5.763 0 0 0-.825.141a6.938 6.938 0 0 0-2.513 1.2l.6.8Zm2.57-1.612l.12 1.259l.996-.095l-.12-1.258l-.996.094ZM9.734 2.87l.168-1.306l-.992-.128l-.168 1.307l.992.127ZM9.406 1.5l.01.5h-.001a.497.497 0 0 1 .049 0c.038.002.1.005.179.013c.16.014.394.047.681.117a5.94 5.94 0 0 1 2.15 1.029l.6-.8a6.937 6.937 0 0 0-2.513-1.2a5.76 5.76 0 0 0-.825-.142A3.98 3.98 0 0 0 9.399 1h-.003l.01.5Zm3.368 1.259l-.469.174l.001.003l.004.009l.013.037l.053.149a35.482 35.482 0 0 1 .777 2.552c.428 1.624.847 3.697.847 5.564h1c0-1.994-.444-4.164-.88-5.819a36.512 36.512 0 0 0-.8-2.629a15.246 15.246 0 0 0-.057-.158l-.015-.042l-.004-.01l-.001-.004l-.47.174Zm1.726 8.488l-.493-.08v-.003l-.002.008l-.01.047c-.012.045-.03.113-.061.197c-.062.17-.167.396-.34.624c-.332.439-.944.923-2.11.96l.032 1c1.483-.047 2.37-.69 2.876-1.356a3.395 3.395 0 0 0 .573-1.184a2.05 2.05 0 0 0 .026-.118l.002-.01v-.004c0-.001 0-.002-.493-.081ZM5.259 6.97c-1.002 0-1.723.867-1.723 1.83h1c0-.498.357-.83.723-.83v-1ZM3.536 8.8c0 .967.736 1.83 1.723 1.83v-1c-.357 0-.723-.334-.723-.83h-1Zm1.723 1.83c1 0 1.722-.866 1.722-1.83h-1c0 .5-.357.83-.722.83v1ZM6.98 8.81c.016-.978-.728-1.84-1.722-1.84v1.001c.372 0 .73.338.722.822l1 .017Zm2.653-1.84c-1.002.001-1.723.868-1.723 1.831h1c0-.498.357-.83.723-.83v-1ZM7.91 8.802c0 .967.736 1.83 1.723 1.83v-1c-.357 0-.723-.334-.723-.83h-1Zm1.723 1.83c1 0 1.722-.866 1.722-1.83h-1c0 .5-.357.83-.722.83v1Zm1.722-1.83c0-.963-.721-1.83-1.722-1.83v1c.365 0 .722.332.722.83h1ZM3.74 4.44c1.443-.787 2.619-1.154 3.763-1.155c1.145 0 2.318.365 3.758 1.154l.48-.876c-1.522-.835-2.865-1.279-4.238-1.278c-1.373 0-2.717.445-4.241 1.277l.478.878Z"/>`
	discordSolidInnerSVG                     = `<path fill="currentColor" d="M4.42 7.727c0-.534.382-.89.774-.89c.399 0 .783.363.774.882v.008c0 .535-.383.89-.774.89c-.382 0-.774-.359-.774-.89Zm4.687 0c0-.282.086-.48.194-.599a.622.622 0 0 1 .485-.195c.391 0 .774.355.774.89c0 .534-.383.89-.774.89c-.16 0-.316-.078-.45-.255a1.233 1.233 0 0 1-.229-.73Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.8 12.142c.277.108.583.217.918.317c-.59.612-1.504 1.47-1.504 1.47C.376 13.839 0 11.514 0 11.514C0 7.38 1.85 2.42 1.85 2.42c1.848-1.387 3.607-1.35 3.607-1.35l.163 1.096c-.852.227-1.726.601-2.663 1.113l.513.94c1.546-.843 2.805-1.236 4.032-1.237c1.226 0 2.483.391 4.025 1.237l.515-.94c-.983-.539-1.896-.925-2.788-1.147l.287-1.062s1.76-.038 3.609 1.349c0 0 1.849 4.959 1.849 9.094c0 0-.376 2.325-3.214 2.415c0 0-.854-.842-1.418-1.449c.296-.09.596-.2.9-.326a7.88 7.88 0 0 0 1.574-.809l-.59-.895c-.492.325-.957.55-1.368.703l-.007.003l-.005.002l-.01.004a7.916 7.916 0 0 1-1.744.518l-.009.001a8.462 8.462 0 0 1-3.127-.012a10.255 10.255 0 0 1-1.793-.525c-.504-.197-.935-.406-1.352-.698l-.614.878a7.3 7.3 0 0 0 1.576.818Zm1.394-6.376c-1.073 0-1.846.93-1.846 1.961c0 1.036.79 1.962 1.846 1.962c1.071 0 1.843-.927 1.845-1.957c.015-1.046-.782-1.966-1.845-1.966Zm4.592.096c-.507 0-.959.193-1.28.547c-.315.35-.47.818-.47 1.318c0 .5.156.995.446 1.378c.293.388.744.68 1.304.68c1.073 0 1.845-.93 1.845-1.962s-.772-1.961-1.845-1.961Z" clip-rule="evenodd"/>`
	discountOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M5 5.5h1m3 4h1M10 5l-5 5M6.801.79L5.672 1.917a.988.988 0 0 1-.698.29H3.196a.988.988 0 0 0-.988.988v1.778a.988.988 0 0 1-.29.698L.79 6.802a.988.988 0 0 0 0 1.397l1.13 1.129a.987.987 0 0 1 .289.698v1.778c0 .546.442.988.988.988h1.778c.262 0 .513.104.698.29l1.13 1.129a.988.988 0 0 0 1.397 0l1.129-1.13a.988.988 0 0 1 .698-.289h1.778a.988.988 0 0 0 .988-.988v-1.778c0-.262.104-.513.29-.698l1.129-1.13a.988.988 0 0 0 0-1.397l-1.13-1.129a.988.988 0 0 1-.289-.698V3.196a.988.988 0 0 0-.988-.988h-1.778a.988.988 0 0 1-.698-.29L8.198.79a.988.988 0 0 0-1.397 0Z"/>`
	discountSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="m6.448.436l-1.13 1.129a.488.488 0 0 1-.344.143H3.196c-.822 0-1.488.666-1.488 1.488v1.778a.49.49 0 0 1-.143.345L.435 6.448a1.488 1.488 0 0 0 0 2.104l1.13 1.13a.488.488 0 0 1 .143.344v1.778c0 .822.666 1.488 1.488 1.488h1.778a.49.49 0 0 1 .345.143l1.129 1.13a1.488 1.488 0 0 0 2.104 0l1.13-1.13a.488.488 0 0 1 .344-.143h1.778c.822 0 1.488-.666 1.488-1.488v-1.778a.49.49 0 0 1 .143-.345l1.13-1.129a1.488 1.488 0 0 0 0-2.104l-1.13-1.13a.488.488 0 0 1-.143-.344V3.196c0-.822-.666-1.488-1.488-1.488h-1.778a.488.488 0 0 1-.345-.143L8.552.435a1.488 1.488 0 0 0-2.104 0Zm-1.802 9.21l5-5l.708.708l-5 5l-.708-.708ZM5 5v1h1V5H5Zm4 5h1V9H9v1Z" clip-rule="evenodd"/>`
	distributeHorizontalOutlineInnerSVG      = `<path fill="none" stroke="currentColor" d="M.5 15V0m14 15V0m-9 13.5v-12h4v12h-4Z"/>`
	distributeHorizontalSolidInnerSVG        = `<path fill="currentColor" d="M1 15V0H0v15h1Zm14 0V0h-1v15h1ZM10 1H5v13h5V1Z"/>`
	distributeVerticalOutlineInnerSVG        = `<path fill="none" stroke="currentColor" d="M15 14.5H0m15-14H0m13.5 9h-12v-4h12v4Z"/>`
	distributeVerticalSolidInnerSVG          = `<path fill="currentColor" d="M15 0H0v1h15V0Zm-1 5H1v5h13V5Zm1 9H0v1h15v-1Z"/>`
	dividerLineOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M0 7.5h15m-12-3h7m-7-3h9m-9 9h9m-9 3h7"/>`
	dividerLineSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M12 2H3V1h9v1Zm-2 3H3V4h7v1Zm5 3H0V7h15v1Zm-3 3H3v-1h9v1Zm-2 3H3v-1h7v1Z" clip-rule="evenodd"/>`
	docOutlineInnerSVG                       = `<path fill="currentColor" d="M2.5 6.5V6H2v.5h.5Zm0 4H2v.5h.5v-.5Zm10-4h.5V6h-.5v.5Zm0 4v.5h.5v-.5h-.5Zm1-7h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5ZM2 6.5v4h1v-4H2Zm.5 4.5h1v-1h-1v1ZM5 9.5v-2H4v2h1ZM3.5 6h-1v1h1V6ZM5 7.5A1.5 1.5 0 0 0 3.5 6v1a.5.5 0 0 1 .5.5h1ZM3.5 11A1.5 1.5 0 0 0 5 9.5H4a.5.5 0 0 1-.5.5v1ZM6 7.5v2h1v-2H6Zm3 2v-2H8v2h1Zm0-2A1.5 1.5 0 0 0 7.5 6v1a.5.5 0 0 1 .5.5h1ZM7.5 11A1.5 1.5 0 0 0 9 9.5H8a.5.5 0 0 1-.5.5v1ZM6 9.5A1.5 1.5 0 0 0 7.5 11v-1a.5.5 0 0 1-.5-.5H6Zm1-2a.5.5 0 0 1 .5-.5V6A1.5 1.5 0 0 0 6 7.5h1ZM10 6v5h1V6h-1Zm.5 1h2V6h-2v1Zm1.5-.5V8h1V6.5h-1ZM10.5 11h2v-1h-2v1Zm2.5-.5V9h-1v1.5h1ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1Z"/>`
	docSolidInnerSVG                         = `<path fill="currentColor" d="M3 10V7h.5a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H3Zm4-2.5a.5.5 0 0 1 1 0v2a.5.5 0 0 1-1 0v-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM3.5 6H2v5h1.5A1.5 1.5 0 0 0 5 9.5v-2A1.5 1.5 0 0 0 3.5 6Zm4 0A1.5 1.5 0 0 0 6 7.5v2a1.5 1.5 0 0 0 3 0v-2A1.5 1.5 0 0 0 7.5 6Zm2.5 5V6h3v2h-1V7h-1v3h1V9h1v2h-3Z" clip-rule="evenodd"/>`
	dockerOutlineInnerSVG                    = `<path fill="currentColor" d="M.5 5.5V5H0v.5h.5Zm2-2V3H2v.5h.5Zm4-2V1H6v.5h.5Zm2 0H9V1h-.5v.5Zm4 6H12V8h.5v-.5ZM1 7.5v-2H0v2h1Zm2 0v-4H2v4h1ZM2.5 4h6V3h-6v1ZM8 3.5v4h1v-4H8Zm-3 4v-4H4v4h1Zm2 0v-6H6v6h1ZM6.5 2h2V1h-2v1ZM8 1.5v2h1v-2H8Zm5.736 8.5H15V9h-1.264v1ZM10 5v.5h1V5h-1Zm2 1.5v1h1v-1h-1Zm.5 1.5h1V7h-1v1Zm1.5.5v1h1v-1h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 7v1Zm-2-2a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 5v1ZM3 10h1V9H3v1Zm5.5-3h-8v1h8V7ZM0 7.5v1h1v-1H0ZM5.5 14h.528v-1H5.5v1Zm.528 0a7.736 7.736 0 0 0 6.23-3.15l-.805-.593A6.737 6.737 0 0 1 6.028 13v1ZM0 8.5A5.5 5.5 0 0 0 5.5 14v-1A4.5 4.5 0 0 1 1 8.5H0ZM.5 6h11V5H.5v1Zm9.5-.5A1.5 1.5 0 0 1 8.5 7v1A2.5 2.5 0 0 0 11 5.5h-1ZM13.736 9c-.96 0-1.769.558-2.283 1.257l.806.593c.383-.522.922-.85 1.477-.85V9Z"/>`
	dockerSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M9 1H6v2H2v2H0v3.5A5.5 5.5 0 0 0 5.5 14h.528a7.736 7.736 0 0 0 6.774-4H15V8.5A1.5 1.5 0 0 0 13.5 7H13v-.5A1.5 1.5 0 0 0 11.5 5H9V1ZM1 7h1V6H1v1Zm2 0h1V6H3v1Zm2 0h1V6H5v1Zm2 0h1V6H7v1Zm2 0h1V6H9v1ZM8 3V2H7v1h1ZM6 4H5v1h1V4Zm1 1V4h1v1H7ZM4 5V4H3v1h1Zm-1 5h1V9H3v1Z" clip-rule="evenodd"/>`
	documentsOutlineInnerSVG                 = `<path fill="currentColor" d="m10.5.5l.277-.416L10.651 0H10.5v.5Zm3 2h.5v-.268l-.223-.148l-.277.416Zm-1 9.5h-8v1h8v-1ZM4 11.5v-10H3v10h1ZM4.5 1h6V0h-6v1ZM13 2.5v9h1v-9h-1ZM10.223.916l3 2l.554-.832l-3-2l-.554.832ZM4.5 12a.5.5 0 0 1-.5-.5H3A1.5 1.5 0 0 0 4.5 13v-1Zm8 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM4 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 3 1.5h1Zm-3 2v10h1v-10H1ZM2.5 15h8v-1h-8v1Zm0-12h1V2h-1v1ZM12 13.5v-1h-1v1h1ZM10.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1Zm1-10a.5.5 0 0 1 .5-.5V2A1.5 1.5 0 0 0 1 3.5h1Z"/>`
	documentsSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h6.151L14 2.232V11.5a1.5 1.5 0 0 1-1.5 1.5H12v.5a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 1 13.5v-10A1.5 1.5 0 0 1 2.5 2H3v-.5ZM3 3h-.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V13H4.5A1.5 1.5 0 0 1 3 11.5V3Z" clip-rule="evenodd"/>`
	dollarOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M2.5 10.5a3 3 0 0 0 3 3h4a3 3 0 1 0 0-6h-4a3 3 0 0 1 0-6h4a3 3 0 0 1 3 3M7.5 0v1.5m0 13.5v-1.5"/>`
	dollarSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7 1V0h1v1h1.5A3.5 3.5 0 0 1 13 4.5h-1A2.5 2.5 0 0 0 9.5 2h-4a2.5 2.5 0 0 0 0 5h4a3.5 3.5 0 1 1 0 7H8v1H7v-1H5.5A3.5 3.5 0 0 1 2 10.5h1A2.5 2.5 0 0 0 5.5 13h4a2.5 2.5 0 0 0 0-5h-4a3.5 3.5 0 1 1 0-7H7Z" clip-rule="evenodd"/>`
	donutChartOutlineInnerSVG                = `<path fill="currentColor" d="M7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Zm2.697 6.46l3.5-1.5l-.394-.92l-3.5 1.5l.394.92ZM7 0v4.5h1V0H7Zm2.146 9.854l3 3l.708-.708l-3-3l-.708.708ZM7.5 10A2.5 2.5 0 0 1 5 7.5H4A3.5 3.5 0 0 0 7.5 11v-1ZM10 7.5A2.5 2.5 0 0 1 7.5 10v1A3.5 3.5 0 0 0 11 7.5h-1ZM7.5 5A2.5 2.5 0 0 1 10 7.5h1A3.5 3.5 0 0 0 7.5 4v1Zm0-1A3.5 3.5 0 0 0 4 7.5h1A2.5 2.5 0 0 1 7.5 5V4Z"/>`
	donutChartSolidInnerSVG                  = `<path fill="currentColor" d="M0 7.5A7.5 7.5 0 0 1 7 .016v4.02a3.5 3.5 0 1 0 2.596 6.267l2.842 2.842A7.5 7.5 0 0 1 0 7.5Z"/><path fill="currentColor" d="M13.145 12.438A7.471 7.471 0 0 0 15 7.5c0-1.034-.21-2.018-.587-2.914L10.755 6.21a3.498 3.498 0 0 1-.452 3.385l2.842 2.842ZM8 4.035V.016a7.499 7.499 0 0 1 5.963 3.676L10.254 5.34A3.497 3.497 0 0 0 8 4.035Z"/>`
	doubleCaretDownCircleOutlineInnerSVG     = `<path fill="currentColor" d="m9.854 8.854l.353-.354l-.707-.707l-.354.353l.708.708ZM7.5 10.5l-.354.354l.354.353l.354-.353L7.5 10.5ZM5.854 8.146L5.5 7.793l-.707.707l.353.354l.708-.708Zm4-2.292l.353-.354l-.707-.707l-.354.353l.708.708ZM7.5 7.5l-.354.354l.354.353l.354-.353L7.5 7.5ZM5.854 5.146L5.5 4.793l-.707.707l.353.354l.708-.708ZM14.5 7.5H14h.5Zm-7-7V1V.5Zm0 14v.5v-.5Zm-7-7H0h.5Zm8.646.646l-2 2l.708.708l2-2l-.708-.708Zm-1.292 2l-2-2l-.708.708l2 2l.708-.708Zm1.292-5l-2 2l.708.708l2-2l-.708-.708Zm-1.292 2l-2-2l-.708.708l2 2l.708-.708ZM15 7.5A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1ZM7.5 15A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1ZM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0Zm1 0A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1Z"/>`
	doubleCaretDownCircleSolidInnerSVG       = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM4.793 5.5L7.5 8.207L10.207 5.5L9.5 4.793l-2 2l-2-2l-.707.707Zm0 3L7.5 11.207L10.207 8.5L9.5 7.793l-2 2l-2-2l-.707.707Z" clip-rule="evenodd"/>`
	doubleCaretDownOutlineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m1 7l6.5 7L14 7M1 1.5l6.5 7l6.5-7"/>`
	doubleCaretDownSmallOutlineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m9.5 8.5l-2 2l-2-2m4-3l-2 2l-2-2"/>`
	doubleCaretDownSmallSolidInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="m5.5 4.793l2 2l2-2l.707.707L7.5 8.207L4.793 5.5l.707-.707Zm0 3l2 2l2-2l.707.707L7.5 11.207L4.793 8.5l.707-.707Z" clip-rule="evenodd"/>`
	doubleCaretDownSolidInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M14.707 1.474L7.5 9.234L.293 1.475l.733-.68L7.5 7.764L13.974.793l.733.68Zm-13.68 4.82l6.473 6.97l6.474-6.972l.733.68L7.5 14.736L.293 6.974l.733-.68Z" clip-rule="evenodd"/>`
	doubleCaretLeftCircleOutlineInnerSVG     = `<path fill="currentColor" d="m6.146 9.854l.354.353l.707-.707l-.353-.354l-.708.708ZM4.5 7.5l-.354-.354l-.353.354l.353.354L4.5 7.5Zm2.354-1.646l.353-.354l-.707-.707l-.354.353l.708.708Zm2.292 4l.354.353l.707-.707l-.353-.354l-.708.708ZM7.5 7.5l-.354-.354l-.353.354l.353.354L7.5 7.5Zm2.354-1.646l.353-.354l-.707-.707l-.354.353l.708.708Zm-3 3.292l-2-2l-.708.708l2 2l.708-.708Zm-2-1.292l2-2l-.708-.708l-2 2l.708.708Zm5 1.292l-2-2l-.708.708l2 2l.708-.708Zm-2-1.292l2-2l-.708-.708l-2 2l.708.708ZM7.5 15A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1ZM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Zm0 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Z"/>`
	doubleCaretLeftCircleSolidInnerSVG       = `<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 0-15 0a7.5 7.5 0 0 0 15 0ZM9.5 4.793L6.793 7.5L9.5 10.207l.707-.707l-2-2l2-2l-.707-.707Zm-3 0L3.793 7.5L6.5 10.207l.707-.707l-2-2l2-2l-.707-.707Z" clip-rule="evenodd"/>`
	doubleCaretLeftOutlineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M8 1L1 7.5L8 14m5.5-13l-7 6.5l7 6.5"/>`
	doubleCaretLeftSmallOutlineInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m6.5 9.5l-2-2l2-2m3 4l-2-2l2-2"/>`
	doubleCaretLeftSmallSolidInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="m7.207 5.5l-2 2l2 2l-.707.707L3.793 7.5L6.5 4.793l.707.707Zm3 0l-2 2l2 2l-.707.707L6.793 7.5L9.5 4.793l.707.707Z" clip-rule="evenodd"/>`
	doubleCaretLeftSolidInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M8.707 1.026L1.735 7.5l6.972 6.474l-.68.733L.264 7.5L8.026.293l.68.733Zm5.5 0L7.235 7.5l6.972 6.474l-.68.733L5.764 7.5L13.526.293l.68.733Z" clip-rule="evenodd"/>`
	doubleCaretRightCircleOutlineInnerSVG    = `<path fill="currentColor" d="m8.146 9.146l-.353.354l.707.707l.354-.353l-.708-.708ZM10.5 7.5l.354.354l.353-.354l-.353-.354l-.354.354ZM8.854 5.146L8.5 4.793l-.707.707l.353.354l.708-.708Zm-3.708 4l-.353.354l.707.707l.354-.353l-.708-.708ZM7.5 7.5l.354.354l.353-.354l-.353-.354L7.5 7.5ZM5.854 5.146L5.5 4.793l-.707.707l.353.354l.708-.708Zm3 4.708l2-2l-.708-.708l-2 2l.708.708Zm2-2.708l-2-2l-.708.708l2 2l.708-.708Zm-5 2.708l2-2l-.708-.708l-2 2l.708.708Zm2-2.708l-2-2l-.708.708l2 2l.708-.708ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	doubleCaretRightCircleSolidInnerSVG      = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm5.5-2.707L8.207 7.5L5.5 10.207L4.793 9.5l2-2l-2-2l.707-.707Zm3 0L11.207 7.5L8.5 10.207L7.793 9.5l2-2l-2-2l.707-.707Z" clip-rule="evenodd"/>`
	doubleCaretRightOutlineInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m7 14l7-6.5L7 1M1.5 14l7-6.5l-7-6.5"/>`
	doubleCaretRightSmallOutlineInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m8.5 9.5l2-2l-2-2m-3 4l2-2l-2-2"/>`
	doubleCaretRightSmallSolidInnerSVG       = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 4.793L8.207 7.5L5.5 10.207L4.793 9.5l2-2l-2-2l.707-.707Zm3 0L11.207 7.5L8.5 10.207L7.793 9.5l2-2l-2-2l.707-.707Z" clip-rule="evenodd"/>`
	doubleCaretRightSolidInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M1.474.293L9.234 7.5l-7.76 7.207l-.68-.733L7.764 7.5L.793 1.026l.68-.733Zm5.5 0l7.76 7.207l-7.76 7.207l-.68-.733l6.97-6.474l-6.971-6.474l.68-.733Z" clip-rule="evenodd"/>`
	doubleCaretUpCircleOutlineInnerSVG       = `<path fill="currentColor" d="m5.146 6.146l-.353.354l.707.707l.354-.353l-.708-.708ZM7.5 4.5l.354-.354l-.354-.353l-.354.353l.354.354Zm1.646 2.354l.354.353l.707-.707l-.353-.354l-.708.708Zm-4 2.292l-.353.354l.707.707l.354-.353l-.708-.708ZM7.5 7.5l.354-.354l-.354-.353l-.354.353l.354.354Zm1.646 2.354l.354.353l.707-.707l-.353-.354l-.708.708ZM.5 7.5H0h.5Zm7 7v.5v-.5Zm0-14V1V.5Zm7 7H14h.5Zm-8.646-.646l2-2l-.708-.708l-2 2l.708.708Zm1.292-2l2 2l.708-.708l-2-2l-.708.708Zm-1.292 5l2-2l-.708-.708l-2 2l.708.708Zm1.292-2l2 2l.708-.708l-2-2l-.708.708ZM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM15 7.5A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1Zm-1 0A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1Z"/>`
	doubleCaretUpCircleSolidInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 15a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm2.707-5.5L7.5 6.793L4.793 9.5l.707.707l2-2l2 2l.707-.707Zm0-3L7.5 3.793L4.793 6.5l.707.707l2-2l2 2l.707-.707Z" clip-rule="evenodd"/>`
	doubleCaretUpOutlineInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M14 8L7.5 1L1 8m13 5.5l-6.5-7l-6.5 7"/>`
	doubleCaretUpSmallOutlineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m5.5 6.5l2-2l2 2m-4 3l2-2l2 2"/>`
	doubleCaretUpSmallSolidInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.793L10.207 6.5l-.707.707l-2-2l-2 2l-.707-.707L7.5 3.793Zm0 3L10.207 9.5l-.707.707l-2-2l-2 2l-.707-.707L7.5 6.793Z" clip-rule="evenodd"/>`
	doubleCaretUpSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="m7.5.265l7.207 7.761l-.733.68L7.5 1.736L1.026 8.707l-.733-.68L7.5.264Zm0 5.5l7.207 7.761l-.733.68L7.5 7.236l-6.474 6.972l-.733-.68L7.5 5.764Z" clip-rule="evenodd"/>`
	downCircleOutlineInnerSVG                = `<path fill="currentColor" d="M4.854 6.146L4.5 5.793l-.707.707l.353.354l.708-.708ZM7.5 9.5l-.354.354l.354.353l.354-.353L7.5 9.5Zm3.354-2.646l.353-.354l-.707-.707l-.354.353l.708.708Zm-6.708 0l3 3l.708-.708l-3-3l-.708.708Zm3.708 3l3-3l-.708-.708l-3 3l.708.708ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1Zm-1 0A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0Z"/>`
	downCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 15a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15Zm4.207-9L7.5 10.207L3.293 6h8.414Z" clip-rule="evenodd"/>`
	downOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m14 5l-6.5 7L1 5"/>`
	downSmallOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m4.5 6.5l3 3l3-3"/>`
	downSmallSolidInnerSVG                   = `<path fill="currentColor" d="M7.5 10.207L11.707 6H3.293L7.5 10.207Z"/>`
	downSolidInnerSVG                        = `<path fill="currentColor" d="M7.5 12L0 4h15l-7.5 8Z"/>`
	downloadOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="m7.5 10.5l-3.25-3m3.25 3l3-3m-3 3V1m6 6v6.5h-12V7"/>`
	downloadSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7 9.358V1h1v8.293l2.146-2.147l.708.708l-3.34 3.34L3.91 7.866l.678-.734L7 9.358ZM2 13V7H1v7h13V7h-1v6H2Z" clip-rule="evenodd"/>`
	dragHorizontalOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M3 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-10 4a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/>`
	dragHorizontalSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 5.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm-10 4a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/>`
	dragOutlineInnerSVG                      = `<path fill="currentColor" d="m7.5 7.5l.137-.48a.5.5 0 0 0-.618.617L7.5 7.5Zm2 7l-.48.137a.5.5 0 0 0 .94.06L9.5 14.5Zm5-5l.197.46a.5.5 0 0 0-.06-.94l-.137.48ZM11 11l-.197-.46a.5.5 0 0 0-.263.263L11 11ZM3.5 3.5V3H3v.5h.5Zm10 0h.5V3h-.5v.5Zm-10 10H3v.5h.5v-.5ZM0 1h1V0H0v1Zm4 0h1V0H4v1Zm4 0h1V0H8v1ZM0 5h1V4H0v1Zm0 4h1V8H0v1Zm7.02-1.363l2 7l.96-.274l-2-7l-.96.274Zm7.617 1.382l-7-2l-.274.962l7 2l.274-.962ZM9.96 14.697l1.5-3.5l-.92-.394l-1.5 3.5l.92.394Zm1.237-3.237l3.5-1.5l-.394-.92l-3.5 1.5l.394.92ZM3.5 4h10V3h-10v1Zm9.5-.5V7h1V3.5h-1Zm-10 0v10h1v-10H3ZM3.5 14H7v-1H3.5v1Z"/>`
	dragSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M1 1H0V0h1v1Zm4 0H4V0h1v1Zm4 0H8V0h1v1ZM3 3h11v4h-1V4H4v9h3v1H3V3ZM1 5H0V4h1v1Zm6.146 2.146a.5.5 0 0 1 .491-.127l7 2a.5.5 0 0 1 .06.94l-3.316 1.422l-1.421 3.316a.5.5 0 0 1-.94-.06l-2-7a.5.5 0 0 1 .126-.49Zm1.082 1.082l1.366 4.782l.946-2.207a.5.5 0 0 1 .263-.263l2.207-.946l-4.782-1.366ZM1 9H0V8h1v1Z" clip-rule="evenodd"/>`
	dragVerticalOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M9.5 3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm-4-10a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/>`
	dragVerticalSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 2.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm4 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm-4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm4 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm-4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm4 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/>`
	dribbbleOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.839 1.024c3.346 4.041 5.096 7.922 5.704 12.782M.533 6.82c5.985-.138 9.402-1.083 11.97-4.216M2.7 12.594c3.221-4.902 7.171-5.65 11.755-4.293M14.5 7.5a7 7 0 1 0-14 0a7 7 0 0 0 14 0Z"/>`
	dribbbleSolidInnerSVG                    = `<path fill="currentColor" d="M4.044.842A7.508 7.508 0 0 0 .092 6.32h.435c2.805-.065 5.004-.308 6.8-.858c-.78-1.383-1.732-2.74-2.874-4.12l-.001-.001l-.408-.499ZM.002 7.32L0 7.5c0 2.017.796 3.848 2.091 5.196l.161-.324a.498.498 0 0 1 .03-.052C3.94 9.798 5.816 8.298 7.914 7.625c.142-.046.286-.088.43-.126a21.803 21.803 0 0 0-.537-1.14c-1.965.633-4.327.893-7.263.96H.003Z"/><path fill="currentColor" d="M2.86 13.393A7.468 7.468 0 0 0 7.5 15c.893 0 1.75-.156 2.543-.442v-.72c-.244-1.935-.673-3.71-1.318-5.404c-.17.042-.339.09-.506.143c-1.822.585-3.525 1.903-5.085 4.268l-.273.548Zm8.183.719a7.512 7.512 0 0 0 3.802-5.086l-.565-.255c-1.626-.478-3.141-.674-4.553-.515c.638 1.72 1.067 3.526 1.312 5.488c.003.02.004.04.004.062v.306Zm3.941-6.12a7.471 7.471 0 0 0-1.805-5.39l-.297.329c-1.17 1.423-2.506 2.41-4.13 3.082c.212.424.41.851.593 1.284c1.672-.24 3.43-.014 5.251.525a.497.497 0 0 1 .065.024l.323.146Zm-2.508-6.104A7.472 7.472 0 0 0 7.5 0c-.878 0-1.72.15-2.503.428l.228.278c1.22 1.473 2.232 2.929 3.058 4.418c1.543-.623 2.766-1.534 3.834-2.837a.48.48 0 0 1 .015-.018l.344-.38Z"/>`
	dropOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5.5l4.286 4.142a5.737 5.737 0 0 1 1.607 2.971a5.62 5.62 0 0 1-.362 3.334a5.85 5.85 0 0 1-2.21 2.584A6.15 6.15 0 0 1 7.5 14.5a6.15 6.15 0 0 1-3.32-.97a5.85 5.85 0 0 1-2.211-2.583a5.62 5.62 0 0 1-.363-3.334a5.737 5.737 0 0 1 1.608-2.97L7.5.5Z"/>`
	dropSolidInnerSVG                        = `<path fill="currentColor" d="M7.847.14a.5.5 0 0 0-.694 0L2.867 4.283l-.004.003a6.237 6.237 0 0 0-1.747 3.23a6.12 6.12 0 0 0 .394 3.63a6.35 6.35 0 0 0 2.4 2.806A6.65 6.65 0 0 0 7.5 15a6.65 6.65 0 0 0 3.59-1.048a6.348 6.348 0 0 0 2.4-2.805a6.12 6.12 0 0 0 .394-3.63a6.238 6.238 0 0 0-1.747-3.23L7.847.14Z"/>`
	dropperOutlineInnerSVG                   = `<path fill="currentColor" d="m.5 12.5l-.354-.354l-.146.147v.207h.5ZM8 5l.354-.354a.5.5 0 0 0-.708 0L8 5Zm2 2l.354.354a.5.5 0 0 0 0-.708L10 7Zm-7.5 7.5v.5h.207l.147-.146L2.5 14.5Zm-2 0H0a.5.5 0 0 0 .5.5v-.5Zm7-10l-.354-.354l-.353.354l.353.354L7.5 4.5Zm3.086-3.086l-.354-.353l.354.353Zm2.828 0l-.353.354l.353-.354Zm.172.172l.353-.354l-.353.354Zm0 2.828l-.354-.353l.354.353ZM10.5 7.5l-.354.354l.354.353l.354-.353L10.5 7.5ZM.854 12.854l7.5-7.5l-.708-.708l-7.5 7.5l.708.708Zm6.792-7.5l2 2l.708-.708l-2-2l-.708.708Zm2 1.292l-7.5 7.5l.708.708l7.5-7.5l-.708-.708ZM2.5 14h-2v1h2v-1Zm-1.5.5v-2H0v2h1ZM6.146 3.854l5 5l.708-.708l-5-5l-.708.708Zm1.708 1l3.085-3.086l-.707-.707l-3.086 3.085l.708.708Zm5.207-3.086l.171.171l.707-.707l-.171-.171l-.707.707Zm.171 2.293l-3.086 3.085l.708.708l3.085-3.086l-.707-.707Zm-2.378 3.085l-3-3l-.708.708l3 3l.708-.708Zm2.378-5.207a1.5 1.5 0 0 1 0 2.122l.707.707a2.5 2.5 0 0 0 0-3.536l-.707.707Zm-2.293-.171a1.5 1.5 0 0 1 2.122 0l.707-.707a2.5 2.5 0 0 0-3.536 0l.707.707Z"/>`
	dropperSolidInnerSVG                     = `<path fill="currentColor" d="M13.768 1.06a2.5 2.5 0 0 0-3.536 0L7.5 3.794l-.646-.647l-.708.708l5 5l.708-.708l-.647-.646l2.732-2.732a2.5 2.5 0 0 0 0-3.536l-.171-.171ZM6.146 6.146a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708L2.707 15H.5a.5.5 0 0 1-.5-.5v-2.207l6.146-6.147Z"/>`
	edgeOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M4.875 3.79C2.512 3.79.526 5.402.5 7.4a7 7 0 0 1 7-6.9c2.72 0 5.22 1.304 6.38 3.57c.182.361.634 1.176.62 2.401a3.357 3.357 0 0 1-1.664 2.892a3.26 3.26 0 0 1-1.656.457c-.003 0-1.486.072-2.272-.475c-.165-.115-.258-.252-.258-.395c0-.173.136-.252.18-.31c.277-.36.427-.795.427-1.088c0-.294-.06-1.031-.399-1.615M4.875 3.79c.191 0 1.283.019 2.296.55c.894.47 1.362 1.037 1.687 1.598M4.875 3.79C2.69 3.79.827 5.17.538 6.958a3.228 3.228 0 0 0-.032.451a7 7 0 0 0 9.214 6.732c-.57.18-1.172.234-1.765.159m.903-8.363c.292.503.377 1.12.395 1.47a1.769 1.769 0 0 0-1.75-1.652c-.373-.008-.657.127-.83.208l-.02.01a4.414 4.414 0 0 0-1.49 1.194M7.956 14.3a4.131 4.131 0 0 1-1.668-.596l-.005-.002c-.488-.303-.91-.7-1.244-1.167a4.414 4.414 0 0 1 .126-5.368m2.79 7.133a4.13 4.13 0 0 0 1.275-.037c.179-.038.347-.08.483-.124l.072-.024a7.018 7.018 0 0 0 3.642-2.887a.219.219 0 0 0-.291-.309a5.13 5.13 0 0 1-.577.258a5.571 5.571 0 0 1-1.963.353c-2.587 0-4.841-1.78-4.841-4.064a1.721 1.721 0 0 1 .895-1.492a4.414 4.414 0 0 0-1.486 1.193m4.091.3v.018m0 0Z"/>`
	edgeSolidInnerSVG                        = `<path fill="currentColor" d="M7.5 0A7.5 7.5 0 0 0 0 7.382v.013h.006v.012c-.01 1.13.232 2.25.709 3.275v.002a7.5 7.5 0 0 0 9.163 3.932l-.001-.004l.067-.023l.004-.001a7.518 7.518 0 0 0 3.902-3.093a.718.718 0 0 0-.95-1.017c-.167.088-.34.164-.516.23a5.07 5.07 0 0 1-1.787.322c-2.405 0-4.34-1.64-4.342-3.561a1.221 1.221 0 0 1 .621-1.048l.011-.006c.168-.079.36-.165.607-.16h.005a1.269 1.269 0 0 1 1.254 1.17a2.4 2.4 0 0 1 .004.127c0 .157-.095.479-.31.767l-.01.009c-.02.02-.06.058-.102.108a.803.803 0 0 0-.185.514c0 .375.241.644.471.805h.001c.502.35 1.176.477 1.662.53c.504.054.915.035.896.035h.002a3.757 3.757 0 0 0 1.907-.526A3.856 3.856 0 0 0 15 6.475c.015-1.286-.438-2.17-.642-2.569l-.032-.063C13.065 1.377 10.369 0 7.5 0Z"/>`
	editCircleOutlineInnerSVG                = `<path fill="currentColor" d="m4.5 8.5l-.354-.354L4 8.293V8.5h.5Zm4-4l.354-.354a.5.5 0 0 0-.708 0L8.5 4.5Zm2 2l.354.354a.5.5 0 0 0 0-.708L10.5 6.5Zm-4 4v.5h.207l.147-.146L6.5 10.5Zm-2 0H4a.5.5 0 0 0 .5.5v-.5Zm3 3.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM4.854 8.854l4-4l-.708-.708l-4 4l.708.708Zm3.292-4l2 2l.708-.708l-2-2l-.708.708Zm2 1.292l-4 4l.708.708l4-4l-.708-.708ZM6.5 10h-2v1h2v-1Zm-1.5.5v-2H4v2h1Z"/>`
	editCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm8.146-3.354a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708L6.707 11H4.5a.5.5 0 0 1-.5-.5V8.293l4.146-4.147Z" clip-rule="evenodd"/>`
	editOneOutlineInnerSVG                   = `<path fill="currentColor" d="m.5 9.5l-.354-.354L0 9.293V9.5h.5Zm9-9l.354-.354a.5.5 0 0 0-.708 0L9.5.5Zm5 5l.354.354a.5.5 0 0 0 0-.708L14.5 5.5Zm-9 9v.5h.207l.147-.146L5.5 14.5Zm-5 0H0a.5.5 0 0 0 .5.5v-.5Zm.354-4.646l9-9l-.708-.708l-9 9l.708.708Zm8.292-9l5 5l.708-.708l-5-5l-.708.708Zm5 4.292l-9 9l.708.708l9-9l-.708-.708ZM5.5 14h-5v1h5v-1Zm-4.5.5v-5H0v5h1ZM6.146 3.854l5 5l.708-.708l-5-5l-.708.708ZM8 15h7v-1H8v1Z"/>`
	editOneSolidInnerSVG                     = `<path fill="currentColor" d="M9.854.146a.5.5 0 0 0-.708 0L6.5 2.793L12.207 8.5l2.647-2.646a.5.5 0 0 0 0-.708l-5-5ZM0 9.293L5.793 3.5L11.5 9.207L5.707 15H.5a.5.5 0 0 1-.5-.5V9.293ZM8 15h7v-1H8v1Z"/>`
	editOutlineInnerSVG                      = `<path fill="currentColor" d="m.5 10.5l-.354-.354l-.146.147v.207h.5Zm10-10l.354-.354a.5.5 0 0 0-.708 0L10.5.5Zm4 4l.354.354a.5.5 0 0 0 0-.708L14.5 4.5Zm-10 10v.5h.207l.147-.146L4.5 14.5Zm-4 0H0a.5.5 0 0 0 .5.5v-.5Zm.354-3.646l10-10l-.708-.708l-10 10l.708.708Zm9.292-10l4 4l.708-.708l-4-4l-.708.708Zm4 3.292l-10 10l.708.708l10-10l-.708-.708ZM4.5 14h-4v1h4v-1Zm-3.5.5v-4H0v4h1Z"/>`
	editSmallOutlineInnerSVG                 = `<path fill="currentColor" d="m4.5 8.5l-.354-.354L4 8.293V8.5h.5Zm4-4l.354-.354a.5.5 0 0 0-.708 0L8.5 4.5Zm2 2l.354.354a.5.5 0 0 0 0-.708L10.5 6.5Zm-4 4v.5h.207l.147-.146L6.5 10.5Zm-2 0H4a.5.5 0 0 0 .5.5v-.5Zm.354-1.646l4-4l-.708-.708l-4 4l.708.708Zm3.292-4l2 2l.708-.708l-2-2l-.708.708Zm2 1.292l-4 4l.708.708l4-4l-.708-.708ZM6.5 10h-2v1h2v-1Zm-1.5.5v-2H4v2h1Z"/>`
	editSmallSolidInnerSVG                   = `<path fill="currentColor" d="M8.854 4.146a.5.5 0 0 0-.708 0L4 8.293V10.5a.5.5 0 0 0 .5.5h2.207l4.147-4.146a.5.5 0 0 0 0-.708l-2-2Z"/>`
	editSolidInnerSVG                        = `<path fill="currentColor" d="M10.854.146a.5.5 0 0 0-.708 0L0 10.293V14.5a.5.5 0 0 0 .5.5h4.207L14.854 4.854a.5.5 0 0 0 0-.708l-4-4Z"/>`
	elbowConnectorOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M2.5 1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 0h5v12h5m0 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`
	elbowConnectorSolidInnerSVG              = `<path fill="currentColor" d="M1.5 0a1.5 1.5 0 1 0 1.415 2H7v12h5.085a1.5 1.5 0 1 0 0-1H8V1H2.915A1.5 1.5 0 0 0 1.5 0Z"/>`
	envelopeOpenOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="m.5 5l7 3.5l7-3.5m0 .08v8.42a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1V5.08a1 1 0 0 1 .504-.868l6-3.428a1 1 0 0 1 .992 0l6 3.428a1 1 0 0 1 .504.868Z"/>`
	envelopeOpenSolidInnerSVG                = `<path fill="currentColor" d="M6.756.35a1.5 1.5 0 0 1 1.488 0l6 3.428a1.5 1.5 0 0 1 .408.341L7.5 7.933L.348 4.12c.113-.135.25-.251.408-.341l6-3.429Z"/><path fill="currentColor" d="M0 5.067V13.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5V5.067l-7.5 4l-7.5-4Z"/>`
	envelopeOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="m.5 4.5l7 4l7-4m-13-3h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`
	envelopeSolidInnerSVG                    = `<path fill="currentColor" d="M.5 2.5A1.5 1.5 0 0 1 2 1h12a1.5 1.5 0 0 1 1.5 1.5v1.208L8 7.926L.5 3.708V2.5Z"/><path fill="currentColor" d="M.5 4.855V12.5A1.5 1.5 0 0 0 2 14h12a1.5 1.5 0 0 0 1.5-1.5V4.855L8 9.074L.5 4.854Z"/>`
	epsOutlineInnerSVG                       = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm0 6V6H10v.5h.5Zm0 2H10V9h.5v-.5Zm2 0h.5V8h-.5v.5Zm0 2v.5h.5v-.5h-.5Zm-6-4V6H6v.5h.5Zm2 0H9V6h-.5v.5Zm0 2V9H9v-.5h-.5Zm-6-2V6H2v.5h.5Zm0 4H2v.5h.5v-.5ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM13 6h-2.5v1H13V6Zm-3 .5v2h1v-2h-1Zm.5 2.5h2V8h-2v1Zm1.5-.5v2h1v-2h-1Zm.5 1.5H10v1h2.5v-1Zm-6-3h2V6h-2v1ZM8 6.5v2h1v-2H8ZM7 11V8.5H6V11h1Zm0-2.5v-2H6v2h1ZM8.5 8h-2v1h2V8ZM5 6H2.5v1H5V6Zm-3 .5v4h1v-4H2Zm.5 4.5H5v-1H2.5v1Zm0-2H5V8H2.5v1Z"/>`
	epsSolidInnerSVG                         = `<path fill="currentColor" d="M7 8h1V7H7v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM5 6H2v5h3v-1H3V9h2V8H3V7h2V6Zm1 0h3v3H7v2H6V6Zm4 0h3v1h-2v1h2v3h-3v-1h2V9h-2V6Z" clip-rule="evenodd"/>`
	eslintOutlineInnerSVG                    = `<path fill="currentColor" d="m.5 7.5l-.447-.224a.5.5 0 0 0 0 .448L.5 7.5Zm3-6V1a.5.5 0 0 0-.447.276L3.5 1.5Zm8 0l.447-.224A.5.5 0 0 0 11.5 1v.5Zm3 6l.447.224a.5.5 0 0 0 0-.448L14.5 7.5Zm-3 6v.5a.5.5 0 0 0 .447-.276L11.5 13.5Zm-8 0l-.447.224A.5.5 0 0 0 3.5 14v-.5Zm4-9.5l.277-.416l-.277-.185l-.277.185L7.5 4Zm-3 2l-.277-.416L4 5.732V6h.5Zm0 3H4v.268l.223.148L4.5 9Zm3 2l-.277.416l.277.185l.277-.185L7.5 11Zm3-2l.277.416l.223-.148V9h-.5Zm0-3h.5v-.268l-.223-.148L10.5 6ZM.947 7.724l3-6l-.894-.448l-3 6l.894.448ZM3.5 2h8V1h-8v1Zm7.553-.276l3 6l.894-.448l-3-6l-.894.448Zm3 5.552l-3 6l.894.448l3-6l-.894-.448ZM11.5 13h-8v1h8v-1Zm-7.553.276l-3-6l-.894.448l3 6l.894-.448Zm3.276-9.692l-3 2l.554.832l3-2l-.554-.832ZM4 6v3h1V6H4Zm.223 3.416l3 2l.554-.832l-3-2l-.554.832Zm3.554 2l3-2l-.554-.832l-3 2l.554.832ZM11 9V6h-1v3h1Zm-.223-3.416l-3-2l-.554.832l3 2l.554-.832Z"/>`
	eslintSolidInnerSVG                      = `<path fill="currentColor" d="M5 8.732V6.268L7.5 4.6L10 6.268v2.464L7.5 10.4L5 8.732Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.053 1.276A.5.5 0 0 1 3.5 1h8a.5.5 0 0 1 .447.276l3 6a.5.5 0 0 1 0 .448l-3 6A.5.5 0 0 1 11.5 14h-8a.5.5 0 0 1-.447-.276l-3-6a.5.5 0 0 1 0-.448l3-6ZM11 5.732L7.5 3.4L4 5.732v3.536L7.5 11.6L11 9.268V5.732Z" clip-rule="evenodd"/>`
	ethereumOutlineInnerSVG                  = `<path fill="currentColor" d="m7.5.5l.384-.32a.5.5 0 0 0-.768 0L7.5.5Zm-5 6l-.384-.32a.5.5 0 0 0-.04.585L2.5 6.5Zm5 8l-.424.265a.5.5 0 0 0 .848 0L7.5 14.5Zm5-8l.424.265a.5.5 0 0 0-.04-.585l-.384.32Zm-5-2l.186-.464L7.5 3.96l-.186.075l.186.464ZM7.116.18l-5 6l.768.64l5-6l-.768-.64Zm-5.04 6.585l5 8l.848-.53l-5-8l-.848.53Zm5.848 8l5-8l-.848-.53l-5 8l.848.53Zm4.96-8.585l-5-6l-.768.64l5 6l.768-.64Zm-10.198.784l5-2l-.372-.928l-5 2l.372.928Zm4.628-2l5 2l.372-.928l-5-2l-.372.928Z"/>`
	ethereumSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a.5.5 0 0 1 .384.18l5 6a.5.5 0 0 1 .04.585l-5 8a.5.5 0 0 1-.848 0l-5-8a.5.5 0 0 1 .04-.585l5-6A.5.5 0 0 1 7.5 0ZM3.241 6.742L7.5 5.04l4.259 1.703L7.5 13.557L3.241 6.742Zm7.61-1.44L7.5 3.962l-3.35 1.34L7.5 1.28l3.35 4.02Z" clip-rule="evenodd"/>`
	euroOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M13.374 3A6 6 0 0 0 2.5 6.5v2A6 6 0 0 0 13.374 12M0 5.5h7m-7 4h7"/>`
	euroSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M2.174 5A6.503 6.503 0 0 1 13.78 2.708l-.812.584A5.502 5.502 0 0 0 3.207 5H7v1H3.022A5.57 5.57 0 0 0 3 6.5v2c0 .169.008.335.022.5H7v1H3.207a5.502 5.502 0 0 0 9.761 1.708l.812.584A6.503 6.503 0 0 1 2.174 10H0V9h2.019A6.593 6.593 0 0 1 2 8.5v-2c0-.168.006-.335.019-.5H0V5h2.174Z" clip-rule="evenodd"/>`
	exclamationCircleOutlineInnerSVG         = `<path fill="currentColor" d="M8 10.5V10H7v.5h1Zm-1 .01v.5h1v-.5H7ZM7 4v4h1V4H7Zm0 6.5v.01h1v-.01H7Zm.5 3.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	exclamationCircleSolidInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM7 8V4h1v4H7Zm1 2v1.01H7V10h1Z" clip-rule="evenodd"/>`
	exclamationOutlineInnerSVG               = `<path fill="currentColor" d="M8 13.5V13H7v.5h1Zm-1 .01v.5h1v-.5H7ZM7 1v10h1V1H7Zm0 12.5v.01h1v-.01H7Z"/>`
	exclamationSmallOutlineInnerSVG          = `<path fill="currentColor" d="M8 10.5V10H7v.5h1Zm-1 .01v.5h1v-.5H7ZM7 4v4h1V4H7Zm0 6.5v.01h1v-.01H7Z"/>`
	exclamationSmallSolidInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M7 8V4h1v4H7Zm1 2v1.01H7V10h1Z" clip-rule="evenodd"/>`
	exclamationSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M7 11V1h1v10H7Zm1 2v1.01H7V13h1Z" clip-rule="evenodd"/>`
	expandAltOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M10 1.5h3.5m0 0V5m0-3.5l-4 4m-8 4.5v3.5m0 0H5m-3.5 0l4-4"/>`
	expandAltSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M12.293 2H10V1h4v4h-1V2.707L9.854 5.854l-.708-.708L12.293 2Zm-6.44 7.854L2.708 13H5v1H1v-4h1v2.293l3.146-3.147l.708.708Z" clip-rule="evenodd"/>`
	expandOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M13.5 13.5H10m3.5 0V10m0 3.5l-4-4m.5-8h3.5m0 0V5m0-3.5l-4 4M5 1.5H1.5m0 0V5m0-3.5l4 4m-4 4.5v3.5m0 0H5m-3.5 0l4-4"/>`
	expandSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h4v1H2.707l3.147 3.146l-.708.708L2 2.707V5H1V1Zm11.293 1H10V1h4v4h-1V2.707L9.854 5.854l-.708-.708L12.293 2Zm-6.44 7.854L2.708 13H5v1H1v-4h1v2.293l3.146-3.147l.708.708Zm4-.708L13 12.293V10h1v4h-4v-1h2.293L9.146 9.854l.708-.708Z" clip-rule="evenodd"/>`
	eyeClosedOutlineInnerSVG                 = `<path fill="currentColor" d="M7.5 9C5.186 9 3.561 7.848 2.497 6.666a9.368 9.368 0 0 1-1.449-2.164a5.065 5.065 0 0 1-.08-.18l-.004-.007v-.001L.5 4.5l-.464.186v.002l.003.004a2.107 2.107 0 0 0 .026.063l.078.173a10.367 10.367 0 0 0 1.61 2.406C2.94 8.652 4.814 10 7.5 10V9Zm7-4.5a68.887 68.887 0 0 1-.464-.186l-.003.008l-.015.035l-.066.145a9.37 9.37 0 0 1-1.449 2.164C11.44 7.848 9.814 9 7.5 9v1c2.686 0 4.561-1.348 5.747-2.666a10.365 10.365 0 0 0 1.61-2.406a6.164 6.164 0 0 0 .104-.236l.002-.004v-.001h.001L14.5 4.5ZM8 12V9.5H7V12h1Zm-6.646-1.646l2-2l-.708-.708l-2 2l.708.708Zm10.292-2l2 2l.708-.708l-2-2l-.708.708Z"/>`
	eyeClosedSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M2.497 6.666C3.56 7.848 5.186 9 7.5 9s3.939-1.152 5.003-2.334a9.37 9.37 0 0 0 1.449-2.164a4.967 4.967 0 0 0 .08-.18l.004-.007v-.001l.464.186l.464.186v.002l-.003.004l-.005.014a3.334 3.334 0 0 1-.1.222a10.37 10.37 0 0 1-1.61 2.406a9.204 9.204 0 0 1-.598.607l1.706 1.705l-.708.708l-1.774-1.775A7.304 7.304 0 0 1 8 9.984V12H7V9.984A7.304 7.304 0 0 1 3.128 8.58l-1.774 1.775l-.708-.708l1.706-1.705a9.237 9.237 0 0 1-.599-.607a10.367 10.367 0 0 1-1.61-2.406a6.064 6.064 0 0 1-.099-.222L.04 4.692l-.002-.004v-.001H.035L.5 4.5l.464-.186l.004.008a2.633 2.633 0 0 0 .08.18a9.368 9.368 0 0 0 1.449 2.164ZM.964 4.314Z" clip-rule="evenodd"/>`
	eyeOutlineInnerSVG                       = `<path fill="currentColor" d="m.5 7.5l-.464-.186a.5.5 0 0 0 0 .372L.5 7.5Zm14 0l.464.186a.5.5 0 0 0 0-.372L14.5 7.5Zm-7 4.5c-2.314 0-3.939-1.152-5.003-2.334a9.368 9.368 0 0 1-1.449-2.164a5.065 5.065 0 0 1-.08-.18l-.004-.007v-.001L.5 7.5l-.464.186v.002l.003.004a2.107 2.107 0 0 0 .026.063l.078.173a10.368 10.368 0 0 0 1.61 2.406C2.94 11.652 4.814 13 7.5 13v-1Zm-7-4.5l.464.186l.004-.008a2.62 2.62 0 0 1 .08-.18a9.368 9.368 0 0 1 1.449-2.164C3.56 4.152 5.186 3 7.5 3V2C4.814 2 2.939 3.348 1.753 4.666a10.367 10.367 0 0 0-1.61 2.406a6.05 6.05 0 0 0-.104.236l-.002.004v.001H.035L.5 7.5Zm7-4.5c2.314 0 3.939 1.152 5.003 2.334a9.37 9.37 0 0 1 1.449 2.164a4.705 4.705 0 0 1 .08.18l.004.007v.001L14.5 7.5l.464-.186v-.002l-.003-.004a.656.656 0 0 0-.026-.063a9.094 9.094 0 0 0-.39-.773a10.365 10.365 0 0 0-1.298-1.806C12.06 3.348 10.186 2 7.5 2v1Zm7 4.5a68.887 68.887 0 0 1-.464-.186l-.003.008l-.015.035l-.066.145a9.37 9.37 0 0 1-1.449 2.164C11.44 10.848 9.814 12 7.5 12v1c2.686 0 4.561-1.348 5.747-2.665a10.366 10.366 0 0 0 1.61-2.407a6.164 6.164 0 0 0 .104-.236l.002-.004v-.001h.001L14.5 7.5ZM7.5 9A1.5 1.5 0 0 1 6 7.5H5A2.5 2.5 0 0 0 7.5 10V9ZM9 7.5A1.5 1.5 0 0 1 7.5 9v1A2.5 2.5 0 0 0 10 7.5H9ZM7.5 6A1.5 1.5 0 0 1 9 7.5h1A2.5 2.5 0 0 0 7.5 5v1Zm0-1A2.5 2.5 0 0 0 5 7.5h1A1.5 1.5 0 0 1 7.5 6V5Z"/>`
	eyeSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 7.686V9a1.5 1.5 0 0 1 0-3v1.686ZM7.5 5a2.5 2.5 0 0 0 0 5v3c-2.686 0-4.561-1.348-5.747-2.665a10.368 10.368 0 0 1-1.61-2.407a6.05 6.05 0 0 1-.099-.222l-.006-.014l-.001-.004l-.001-.002L.5 7.5l-.464.186a.5.5 0 0 1 0-.372l.066.027a1.304 1.304 0 0 1-.066-.028v-.001l.002-.004l.006-.014a3.62 3.62 0 0 1 .1-.222a10.367 10.367 0 0 1 1.61-2.406C2.938 3.348 4.813 2 7.5 2v3Zm0 1v3a1.5 1.5 0 1 0 0-3Zm0 4a2.5 2.5 0 0 0 0-5V2c2.686 0 4.561 1.348 5.747 2.666a10.365 10.365 0 0 1 1.61 2.406a6.164 6.164 0 0 1 .099.222l.005.014l.002.004l.001.002l-.364.146l.364-.146a.5.5 0 0 1 0 .372L14.5 7.5l.464.187v.001l-.003.004l-.005.014a3.334 3.334 0 0 1-.1.222a10.366 10.366 0 0 1-1.61 2.406C12.062 11.652 10.187 13 7.5 13v-3Z" clip-rule="evenodd"/>`
	faceIdOutlineInnerSVG                    = `<path fill="currentColor" d="M4 6h1V5H4v1Zm6 0h1V5h-1v1Zm.1 2.7a3.25 3.25 0 0 1-5.2 0l-.8.6c1.7 2.267 5.1 2.267 6.8 0l-.8-.6ZM1 5V2.5H0V5h1Zm1.5-4H5V0H2.5v1ZM1 2.5A1.5 1.5 0 0 1 2.5 1V0A2.5 2.5 0 0 0 0 2.5h1ZM0 10v2.5h1V10H0Zm2.5 5H5v-1H2.5v1ZM0 12.5A2.5 2.5 0 0 0 2.5 15v-1A1.5 1.5 0 0 1 1 12.5H0ZM10 1h2.5V0H10v1Zm4 1.5V5h1V2.5h-1ZM12.5 1A1.5 1.5 0 0 1 14 2.5h1A2.5 2.5 0 0 0 12.5 0v1ZM10 15h2.5v-1H10v1Zm5-2.5V10h-1v2.5h1ZM12.5 15a2.5 2.5 0 0 0 2.5-2.5h-1a1.5 1.5 0 0 1-1.5 1.5v1Z"/>`
	faceIdSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M2.5 1A1.5 1.5 0 0 0 1 2.5V5H0V2.5A2.5 2.5 0 0 1 2.5 0H5v1H2.5Zm10 0H10V0h2.5A2.5 2.5 0 0 1 15 2.5V5h-1V2.5A1.5 1.5 0 0 0 12.5 1ZM5 6H4V5h1v1Zm6 0h-1V5h1v1ZM4.9 8.7a3.25 3.25 0 0 0 5.2 0l.8.6c-1.7 2.267-5.1 2.267-6.8 0l.8-.6ZM0 12.5V10h1v2.5A1.5 1.5 0 0 0 2.5 14H5v1H2.5A2.5 2.5 0 0 1 0 12.5ZM15 10v2.5a2.5 2.5 0 0 1-2.5 2.5H10v-1h2.5a1.5 1.5 0 0 0 1.5-1.5V10h1Z" clip-rule="evenodd"/>`
	facebookOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M7.5 14.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm0 0v-8a2 2 0 0 1 2-2h.5m-5 4h5"/>`
	facebookSolidInnerSVG                    = `<path fill="currentColor" d="M0 7.5a7.5 7.5 0 1 1 8 7.484V9h2V8H8V6.5A1.5 1.5 0 0 1 9.5 5h.5V4h-.5A2.5 2.5 0 0 0 7 6.5V8H5v1h2v5.984A7.5 7.5 0 0 1 0 7.5Z"/>`
	figmaOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M7.5 1v8.5m0 0v2a2 2 0 1 1-2-2m2 0h-2m0 0a2 2 0 1 1 0-4m0 0h2m-2 0h4m-4 0a2 2 0 1 1 0-4h4a2 2 0 1 1 0 4m0 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`
	figmaSolidInnerSVG                       = `<path fill="currentColor" d="M4 9.5a2.496 2.496 0 0 1-1-2c0-.818.393-1.544 1-2A2.5 2.5 0 0 1 5.5 1h4A2.5 2.5 0 0 1 11 5.5a2.5 2.5 0 0 1-3 4v2a2.5 2.5 0 1 1-4-2Z"/>`
	fileMinusOutlineInnerSVG                 = `<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354Zm-1 10.5h-10v1h10v-1ZM2 13.5v-12H1v12h1ZM2.5 1h8V0h-8v1ZM13 3.5v10h1v-10h-1ZM10.146.854l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM5 8h5V7H5v1Z"/>`
	fileMinusSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM5 8h5V7H5v1Z" clip-rule="evenodd"/>`
	fileNoAccessOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="m9.5 5.5l-4 4m5-9h-8a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-10l-3-3Zm-3 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`
	fileNoAccessSolidInnerSVG                = `<path fill="currentColor" d="M7.5 5a2.5 2.5 0 0 0-2.086 3.879L8.88 5.414A2.488 2.488 0 0 0 7.5 5Zm0 5c-.51 0-.983-.152-1.379-.414L9.586 6.12A2.5 2.5 0 0 1 7.5 10Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm3 6a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/>`
	fileOutlineInnerSVG                      = `<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354Zm-1 10.5h-10v1h10v-1ZM2 13.5v-12H1v12h1ZM2.5 1h8V0h-8v1ZM13 3.5v10h1v-10h-1ZM10.146.854l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1Z"/>`
	filePlusOutlineInnerSVG                  = `<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354Zm-1 10.5h-10v1h10v-1ZM2 13.5v-12H1v12h1ZM2.5 1h8V0h-8v1ZM13 3.5v10h1v-10h-1ZM10.146.854l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM7 5v5h1V5H7ZM5 8h5V7H5v1Z"/>`
	filePlusSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM7 10V8H5V7h2V5h1v2h2v1H8v2H7Z" clip-rule="evenodd"/>`
	fileSolidInnerSVG                        = `<path fill="currentColor" d="M2.5 0A1.5 1.5 0 0 0 1 1.5v12A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V3.293L10.707 0H2.5Z"/>`
	fileTickOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="m5 7.5l2 2l3.5-4m0-5h-8a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-10l-3-3Z"/>`
	fileTickSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm6.024 8.732l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708l2.378 2.378Z" clip-rule="evenodd"/>`
	fileXOutlineInnerSVG                     = `<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354Zm-1 10.5h-10v1h10v-1ZM2 13.5v-12H1v12h1ZM2.5 1h8V0h-8v1ZM13 3.5v10h1v-10h-1ZM10.146.854l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1Zm3.146 4.354l4 4l.708-.708l-4-4l-.708.708Zm.708 4l4-4l-.708-.708l-4 4l.708.708Z"/>`
	fileXSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm8.146 8.354L7.5 8.207L5.854 9.854l-.708-.708L6.793 7.5L5.146 5.854l.708-.708L7.5 6.793l1.646-1.647l.708.708L8.207 7.5l1.647 1.646l-.708.708Z" clip-rule="evenodd"/>`
	filterOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M0 2.5h15m-12 5h9m-7 5h5"/>`
	filterSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M15 3H0V2h15v1Zm-3 5H3V7h9v1Zm-2 5H5v-1h5v1Z" clip-rule="evenodd"/>`
	fingerprintOutlineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" d="M12.587 3.513a6.03 6.03 0 0 1 .818 3.745v.75c0 .788.205 1.563.595 2.247M4.483 6.508c0-.795.313-1.557.871-2.119a2.963 2.963 0 0 1 2.103-.877c.789 0 1.545.315 2.103.877c.558.562.871 1.324.871 2.12v.748c0 1.621.522 3.198 1.487 4.495m-4.46-5.244v1.498A10.542 10.542 0 0 0 9.315 14M4.483 9.505A13.559 13.559 0 0 0 5.821 14m-3.643-1.498a16.63 16.63 0 0 1-.669-5.244V6.51a6.028 6.028 0 0 1 .79-3.002a5.97 5.97 0 0 1 2.177-2.2a5.914 5.914 0 0 1 5.955-.004"/>`
	fingerprintSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M10.18 1.735a5.414 5.414 0 0 0-5.452.004a5.471 5.471 0 0 0-1.994 2.016a5.528 5.528 0 0 0-.725 2.753v.767a16.07 16.07 0 0 0 .649 5.085l.141.48l-.959.283l-.141-.48a17.096 17.096 0 0 1-.69-5.393v-.74a6.528 6.528 0 0 1 .857-3.25A6.47 6.47 0 0 1 4.224.874A6.414 6.414 0 0 1 10.683.87l.432.251l-.503.864l-.432-.25Zm2.58 1.092l.257.43a6.53 6.53 0 0 1 .888 4.028v.723c0 .701.182 1.39.53 1.999l.247.434l-.868.496l-.248-.434a5.02 5.02 0 0 1-.66-2.496v-.749c0-.018 0-.036.002-.054a5.53 5.53 0 0 0-.75-3.435l-.256-.429l.858-.513ZM7.457 4.013c-.655 0-1.284.262-1.748.73a2.508 2.508 0 0 0-.726 1.766v.5h-1v-.5c0-.926.365-1.815 1.016-2.47a3.463 3.463 0 0 1 2.458-1.026c.923 0 1.807.369 2.458 1.025A3.508 3.508 0 0 1 10.93 6.51v.75c0 1.513.487 2.985 1.388 4.195l.299.401l-.802.597l-.299-.4A8.028 8.028 0 0 1 9.93 7.259V6.51c0-.663-.261-1.299-.726-1.766a2.463 2.463 0 0 0-1.748-.73Zm.5 3.995a10.042 10.042 0 0 0 1.77 5.708l.284.412l-.823.567l-.284-.411a11.042 11.042 0 0 1-1.947-6.277m-2.035.944l.058.497a13.058 13.058 0 0 0 1.289 4.329l.223.447l-.894.447l-.224-.447a14.06 14.06 0 0 1-1.388-4.66l-.057-.497l.993-.116Z" clip-rule="evenodd"/>`
	firebaseOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m2.5 11.5l9-8l1 9l-5 2l-5-3Zm0 0l5-9l2 3m-7 6l1-11l3 3"/>`
	firebaseSolidInnerSVG                    = `<path fill="currentColor" d="M3.33.03a.5.5 0 0 1 .524.116l2.078 2.08a.505.505 0 0 0-.032.056L2.175 9.988L3.33.03ZM2.262 11.94l4.98 2.989a.5.5 0 0 0 .444.035l5-2a.5.5 0 0 0 .31-.52l-1-9a.5.5 0 0 0-.828-.318L9.513 4.597L2.262 11.94Zm6.676-8.184L7.916 2.223a.5.5 0 0 0-.853.034l-.31.558l-2.986 6.177l5.171-5.236Z"/>`
	flagAltOutlineInnerSVG                   = `<path fill="currentColor" d="m12.5 6.5l.224.447a.5.5 0 0 0 .033-.876L12.5 6.5Zm-10-6l.257-.429A.5.5 0 0 0 2 .5h.5Zm10.257 5.571l-10-6l-.514.858l10 6l.514-.858ZM2 .5v11h1V.5H2Zm.724 11.447l10-5l-.448-.894l-10 5l.448.894ZM3 15v-3.5H2V15h1Z"/>`
	flagAltSolidInnerSVG                     = `<path fill="currentColor" d="M2.254.065a.5.5 0 0 1 .503.006l10 6a.5.5 0 0 1-.033.876L3 11.81V15H2V.5a.5.5 0 0 1 .254-.435Z"/>`
	flagOutlineInnerSVG                      = `<path fill="currentColor" d="m14.5.5l.457.203A.5.5 0 0 0 14.5 0v.5ZM.5.5V0a.5.5 0 0 0-.5.5h.5Zm14 9v.5a.5.5 0 0 0 .457-.703L14.5 9.5Zm-2-4.5l-.457-.203a.5.5 0 0 0 0 .406L12.5 5Zm2-5H.5v1h14V0ZM0 .5v9h1v-9H0ZM.5 10h14V9H.5v1Zm14.457-.703l-2-4.5l-.914.406l2 4.5l.914-.406Zm-2-4.094l2-4.5l-.914-.406l-2 4.5l.914.406ZM1 15V9.5H0V15h1Z"/>`
	flagSolidInnerSVG                        = `<path fill="currentColor" d="M0 .5A.5.5 0 0 1 .5 0h14a.5.5 0 0 1 .457.703L13.047 5l1.91 4.297A.5.5 0 0 1 14.5 10H1v5H0V.5Z"/>`
	flipHorizontalOutlineInnerSVG            = `<path fill="currentColor" d="m3.5.5l.224-.447A.5.5 0 0 0 3 .5h.5Zm8 4V5a.5.5 0 0 0 .224-.947L11.5 4.5Zm-8 0H3a.5.5 0 0 0 .5.5v-.5Zm0 6V10a.5.5 0 0 0-.5.5h.5Zm8 0l.224.447A.5.5 0 0 0 11.5 10v.5Zm-8 4H3a.5.5 0 0 0 .724.447L3.5 14.5ZM3.276.947l8 4l.448-.894l-8-4l-.448.894ZM11.5 4h-8v1h8V4ZM4 4.5v-4H3v4h1ZM0 8h15V7H0v1Zm3.5 3h8v-1h-8v1Zm7.776-.947l-8 4l.448.894l8-4l-.448-.894ZM4 14.5v-4H3v4h1Z"/>`
	flipHorizontalSolidInnerSVG              = `<path fill="currentColor" d="M3.237.075a.5.5 0 0 1 .487-.022l8 4A.5.5 0 0 1 11.5 5h-8a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .237-.425ZM0 8h15V7H0v1Zm3.5 2a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .724.447l8-4A.5.5 0 0 0 11.5 10h-8Z"/>`
	flipVerticalOutlineInnerSVG              = `<path fill="currentColor" d="M4.5 3.5H5a.5.5 0 0 0-.947-.224L4.5 3.5Zm0 8v.5a.5.5 0 0 0 .5-.5h-.5Zm-4 0l-.447-.224A.5.5 0 0 0 .5 12v-.5Zm10-8l.447-.224A.5.5 0 0 0 10 3.5h.5Zm0 8H10a.5.5 0 0 0 .5.5v-.5Zm4 0v.5a.5.5 0 0 0 .447-.724l-.447.224ZM4 3.5v8h1v-8H4Zm.5 7.5h-4v1h4v-1Zm-3.553.724l4-8l-.894-.448l-4 8l.894.448ZM10 3.5v8h1v-8h-1Zm.5 8.5h4v-1h-4v1Zm4.447-.724l-4-8l-.894.448l4 8l.894-.448ZM7 0v15h1V0H7Z"/>`
	flipVerticalSolidInnerSVG                = `<path fill="currentColor" d="M7 0v15h1V0H7ZM4.615 3.013A.5.5 0 0 1 5 3.5v8a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.447-.724l4-8a.5.5 0 0 1 .562-.263Zm5.77 0a.5.5 0 0 1 .562.263l4 8A.5.5 0 0 1 14.5 12h-4a.5.5 0 0 1-.5-.5v-8a.5.5 0 0 1 .385-.487Z"/>`
	floatCenterOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M0 5.5h2m11 0h2m-15-4h2m11 0h2m-15 8h15m-15 4h15M5.5.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Z"/>`
	floatCenterSolidInnerSVG                 = `<path fill="currentColor" d="M5.5 0A1.5 1.5 0 0 0 4 1.5v4A1.5 1.5 0 0 0 5.5 7h4A1.5 1.5 0 0 0 11 5.5v-4A1.5 1.5 0 0 0 9.5 0h-4ZM0 2h2V1H0v1Zm13 0h2V1h-2v1ZM0 6h2V5H0v1Zm13 0h2V5h-2v1ZM0 10h15V9H0v1Zm0 4h15v-1H0v1Z"/>`
	floatLeftOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M9 5.5h6m-6-4h6m-15 8h15m-15 4h15M1.5.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Z"/>`
	floatLeftSolidInnerSVG                   = `<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v4A1.5 1.5 0 0 0 1.5 7h4A1.5 1.5 0 0 0 7 5.5v-4A1.5 1.5 0 0 0 5.5 0h-4ZM9 2h6V1H9v1Zm0 4h6V5H9v1Zm-9 4h15V9H0v1Zm0 4h15v-1H0v1Z"/>`
	floatRightOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M0 5.5h6m-6-4h6m-6 8h15m-15 4h15M9.5.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Z"/>`
	floatRightSolidInnerSVG                  = `<path fill="currentColor" d="M9.5 0A1.5 1.5 0 0 0 8 1.5v4A1.5 1.5 0 0 0 9.5 7h4A1.5 1.5 0 0 0 15 5.5v-4A1.5 1.5 0 0 0 13.5 0h-4ZM0 2h6V1H0v1Zm0 4h6V5H0v1Zm0 4h15V9H0v1Zm0 4h15v-1H0v1Z"/>`
	floorplanOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M10 .5h4.5v14H.5V.5h4l3 2m-1 12v-7M4 7.5h5m3 0h2.5"/>`
	floorplanSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M0 0h4.651l3.126 2.084l-.554.832L4.349 1H1v13h5V8H4V7h5v1H7v6h7V8h-2V7h2V1h-4V0h5v15H0V0Z" clip-rule="evenodd"/>`
	folderMinusOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M5 8.5h5m-9.5-6v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Z"/>`
	folderMinusSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10ZM5 9h5V8H5v1Z" clip-rule="evenodd"/>`
	folderNoAccessOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="m9.5 6.5l-4 4m-5-8v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Zm7 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`
	folderNoAccessSolidInnerSVG              = `<path fill="currentColor" d="M7.5 6a2.5 2.5 0 0 0-2.086 3.879L8.88 6.414A2.488 2.488 0 0 0 7.5 6Zm0 5c-.51 0-.983-.152-1.379-.414L9.586 7.12A2.5 2.5 0 0 1 7.5 11Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10Zm4 6a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/>`
	folderOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M.5 12.5v-10a1 1 0 0 1 1-1h4l2 2h6a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1Z"/>`
	folderPlusOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M7.5 6v5M5 8.5h5m-9.5-6v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Z"/>`
	folderPlusSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10ZM7 11V9H5V8h2V6h1v2h2v1H8v2H7Z" clip-rule="evenodd"/>`
	folderSolidInnerSVG                      = `<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v10A1.5 1.5 0 0 0 1.5 14h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 3H7.707l-2-2H1.5Z"/>`
	folderTickOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="m5 8.5l2 2l3.5-4m-10-4v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Z"/>`
	folderTickSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10Zm7.024 8.732l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708l2.378 2.378Z" clip-rule="evenodd"/>`
	folderXOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="m5.5 6.5l4 4m-4 0l4-4m-9-4v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Z"/>`
	folderXSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10Zm9.146 8.354L7.5 9.207l-1.646 1.647l-.708-.707L6.793 8.5L5.146 6.854l.708-.708L7.5 7.793l1.646-1.647l.708.708L8.207 8.5l1.647 1.646l-.708.708Z" clip-rule="evenodd"/>`
	foldersOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M3.5 8.5v-7a1 1 0 0 1 1-1h3l2 2h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-9m-1-1a1 1 0 0 0 1 1m-1-1v-3h-2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-4h-7"/>`
	foldersSolidInnerSVG                     = `<path fill="currentColor" d="M4.5 0A1.5 1.5 0 0 0 3 1.5v7A1.5 1.5 0 0 0 4.5 10h9A1.5 1.5 0 0 0 15 8.5v-5A1.5 1.5 0 0 0 13.5 2H9.707l-2-2H4.5Z"/><path fill="currentColor" d="M12 11H4.5A2.5 2.5 0 0 1 2 8.5V5h-.5A1.5 1.5 0 0 0 0 6.5v7A1.5 1.5 0 0 0 1.5 15h9a1.5 1.5 0 0 0 1.5-1.5V11Z"/>`
	forwardCircleOutlineInnerSVG             = `<path fill="currentColor" d="m8.5 5.5l.248-.434A.5.5 0 0 0 8 5.5h.5Zm0 4H8a.5.5 0 0 0 .748.434L8.5 9.5Zm3.5-2l.248.434a.5.5 0 0 0 0-.868L12 7.5Zm-7.5-2l.248-.434A.5.5 0 0 0 4 5.5h.5Zm0 4H4a.5.5 0 0 0 .748.434L4.5 9.5Zm3.5-2l.248.434a.5.5 0 0 0 0-.868L8 7.5ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM8 5.5v4h1v-4H8Zm.748 4.434l3.5-2l-.496-.868l-3.5 2l.496.868Zm3.5-2.868l-3.5-2l-.496.868l3.5 2l.496-.868ZM4 5.5v4h1v-4H4Zm.748 4.434l3.5-2l-.496-.868l-3.5 2l.496.868Zm3.5-2.868l-3.5-2l-.496.868l3.5 2l.496-.868Z"/>`
	forwardCircleSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm4.249-2.432a.5.5 0 0 1 .5-.002L8 6.924V5.5a.5.5 0 0 1 .748-.434l3.5 2a.5.5 0 0 1 0 .868l-3.5 2A.5.5 0 0 1 8 9.5V8.076L4.748 9.934A.5.5 0 0 1 4 9.5v-4a.5.5 0 0 1 .249-.432Z" clip-rule="evenodd"/>`
	forwardOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M.5 12.5v-10l7 5l-7 5Zm7 0v-10l7 5l-7 5Z"/>`
	forwardSmallOutlineInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M8.5 9.5v-4l3.5 2l-3.5 2Zm-4 0v-4l3.5 2l-3.5 2Z"/>`
	forwardSmallSolidInnerSVG                = `<path fill="currentColor" d="M4.748 5.066A.5.5 0 0 0 4 5.5v4a.5.5 0 0 0 .748.434L8 8.076V9.5a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868l-3.5-2A.5.5 0 0 0 8 5.5v1.424L4.748 5.066Z"/>`
	forwardSolidInnerSVG                     = `<path fill="currentColor" d="M.79 2.093A.5.5 0 0 0 0 2.5v10a.5.5 0 0 0 .79.407L7 8.472V12.5a.5.5 0 0 0 .79.407l7-5a.5.5 0 0 0 0-.814l-7-5A.5.5 0 0 0 7 2.5v4.028L.79 2.093Z"/>`
	frameOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M3.5 0v15m8-15v15M0 3.5h15m-15 8h15"/>`
	frameSolidInnerSVG                       = `<path fill="none" stroke="currentColor" d="M3.5 0v15m8-15v15M0 3.5h15m-15 8h15"/>`
	framerOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M7.5 5.5h5v-5h-10l5 5Zm0 0h-5v4l5 5v-4h5l-5-5Z"/>`
	framerSolidInnerSVG                      = `<path fill="currentColor" d="M2.038.309A.5.5 0 0 1 2.5 0h10a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5H8.707l4.147 4.146A.5.5 0 0 1 12.5 11H8v3.5a.5.5 0 0 1-.854.354l-5-5A.5.5 0 0 1 2 9.5v-4a.5.5 0 0 1 .5-.5h3.793L2.146.854a.5.5 0 0 1-.108-.545Z"/>`
	gameControllerOutlineInnerSVG            = `<path fill="currentColor" d="m9.817 11.133l-.447.224l.447-.224ZM9.5 10.5l.447-.224A.5.5 0 0 0 9.5 10v.5Zm-4 0V10a.5.5 0 0 0-.447.276l.447.224Zm8.5-5v4.528h1V5.5h-1Zm-3.736 5.41l-.317-.634l-.894.448l.316.633l.895-.447ZM9.5 10h-4v1h4v-1Zm-4.447.276l-.317.634l.894.447l.317-.633l-.894-.448ZM1 10.028V5.5H0v4.528h1ZM3.5 3h8V2h-8v1Zm-.528 9A1.972 1.972 0 0 1 1 10.028H0A2.972 2.972 0 0 0 2.972 13v-1Zm9.056 0c-.747 0-1.43-.422-1.764-1.09l-.894.447A2.972 2.972 0 0 0 12.027 13v-1ZM14 10.028A1.972 1.972 0 0 1 12.028 12v1A2.972 2.972 0 0 0 15 10.028h-1Zm-9.264.882A1.972 1.972 0 0 1 2.972 12v1a2.972 2.972 0 0 0 2.658-1.643l-.894-.447ZM15 5.5A3.5 3.5 0 0 0 11.5 2v1A2.5 2.5 0 0 1 14 5.5h1Zm-14 0A2.5 2.5 0 0 1 3.5 3V2A3.5 3.5 0 0 0 0 5.5h1ZM3 7h3V6H3v1Zm1-2v3h1V5H4Zm7 1h1V5h-1v1ZM9 8h1V7H9v1Z"/>`
	gameControllerRetroOutlineInnerSVG       = `<path fill="none" stroke="currentColor" d="M8 5.5h2M4.5 6v3M3 7.5h3m4 2h2M.5 3.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`
	gameControllerRetroSolidInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M13.5 2A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5v-8A1.5 1.5 0 0 1 1.5 2h12ZM10 6H8V5h2v1ZM4 7V6h1v1h1v1H5v1H4V8H3V7h1Zm8 3h-2V9h2v1Z" clip-rule="evenodd"/>`
	gameControllerSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M0 5.5A3.5 3.5 0 0 1 3.5 2h8A3.5 3.5 0 0 1 15 5.5v4.528a2.972 2.972 0 0 1-5.63 1.329L9.19 11H5.809l-.179.357A2.972 2.972 0 0 1 0 10.027V5.5ZM4 8V7H3V6h1V5h1v1h1v1H5v1H4Zm6 0H9V7h1v1Zm1-2h1V5h-1v1Z" clip-rule="evenodd"/>`
	ganttChartOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M.5 0v14.5H15M5 2.5H2m6 3H3m5 3H5m10 3H8"/>`
	ganttChartSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 0h1v14h14v1H0V0Zm2 2h3v1H2V2Zm1 3h5v1H3V5Zm2 3h3v1H5V8Zm3 3h7v1H8v-1Z" clip-rule="evenodd"/>`
	garageOutlineInnerSVG                    = `<path fill="currentColor" d="m.5 5.5l-.29-.407A.5.5 0 0 0 0 5.5h.5Zm7-5l.29-.407a.5.5 0 0 0-.58 0L7.5.5Zm7 5h.5a.5.5 0 0 0-.21-.407l-.29.407Zm-12 2V7H2v.5h.5Zm10 0h.5V7h-.5v.5ZM1 15V5.5H0V15h1ZM.79 5.907l7-5l-.58-.814l-7 5l.58.814Zm6.42-5l7 5l.58-.814l-7-5l-.58.814ZM14 5.5V15h1V5.5h-1ZM3 15V7.5H2V15h1Zm-.5-7h10V7h-10v1Zm9.5-.5V15h1V7.5h-1ZM2.5 11h10v-1h-10v1ZM6 13h3v-1H6v1Z"/>`
	garageSolidInnerSVG                      = `<path fill="currentColor" d="M7.21.093a.5.5 0 0 1 .58 0l7 5A.5.5 0 0 1 15 5.5v9a.5.5 0 0 1-.5.5H13V7H2v8H.5a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .21-.407l7-5Z"/><path fill="currentColor" fill-rule="evenodd" d="M3 15h9v-4H3v4Zm6-2H6v-1h3v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M12 10V8H3v2h9Z"/>`
	gatsbyjsOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M11.07 4a5.002 5.002 0 0 0-8.342 2L9 12.271A5.004 5.004 0 0 0 12.475 8H9m-6.5.5l4 4m1 2a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	gatsbyjsSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm3.305-1.631a4.502 4.502 0 0 1 7.409-1.519l.714-.7a5.502 5.502 0 0 0-9.176 2.2l-.09.29l6.699 6.699l.289-.09a5.504 5.504 0 0 0 3.823-4.7l.054-.549H9v1h2.889a4.51 4.51 0 0 1-2.758 3.195L3.305 5.87Zm2.841 6.985l-4-4l.708-.708l4 4l-.708.708Z" clip-rule="evenodd"/>`
	gbaOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M2 6.5h3m7 0h1m-2 2h1M3.5 5v3m1.606-5.5h4.788a2 2 0 0 1 1.11.336L12 3.5h.138a2 2 0 0 1 1.995 1.858l.32 4.475a1 1 0 0 1-.55.966l-2.091 1.045a3.175 3.175 0 0 1-.65.24a15.097 15.097 0 0 1-7.324 0a3.176 3.176 0 0 1-.65-.24l-2.09-1.045a1 1 0 0 1-.55-.966l.32-4.476A2 2 0 0 1 2.861 3.5H3l.996-.664a2 2 0 0 1 1.11-.336Z"/>`
	gbaSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M3.719 2.42A2.5 2.5 0 0 1 5.106 2h4.789a2.5 2.5 0 0 1 1.386.42l.87.58a2.5 2.5 0 0 1 2.48 2.322l.32 4.476a1.5 1.5 0 0 1-.825 1.448l-2.09 1.045a3.68 3.68 0 0 1-.753.279a15.62 15.62 0 0 1-7.566 0a3.68 3.68 0 0 1-.753-.279l-2.09-1.045A1.5 1.5 0 0 1 .05 9.798l.32-4.476A2.5 2.5 0 0 1 2.849 3l.87-.58ZM3 8V7H2V6h1V5h1v1h1v1H4v1H3Zm10-1h-1V6h1v1Zm-2 2h1V8h-1v1Z" clip-rule="evenodd"/>`
	gbcOutlineInnerSVG                       = `<path fill="currentColor" d="M4.5 2.5V2a.5.5 0 0 0-.5.5h.5Zm6 0h.5a.5.5 0 0 0-.5-.5v.5Zm-6 4H4a.5.5 0 0 0 .5.5v-.5ZM3.5 1h8V0h-8v1Zm8.5.5v10h1v-10h-1ZM9.5 14h-6v1h6v-1ZM3 13.5v-12H2v12h1Zm.5.5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 15v-1Zm8.5-2.5A2.5 2.5 0 0 1 9.5 14v1a3.5 3.5 0 0 0 3.5-3.5h-1ZM11.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 0v1Zm-8-1A1.5 1.5 0 0 0 2 1.5h1a.5.5 0 0 1 .5-.5V0Zm1 3h6V2h-6v1Zm5.5-.5v3h1v-3h-1ZM9.5 6h-5v1h5V6ZM5 6.5v-4H4v4h1Zm5-1a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 11 5.5h-1ZM5 8v3h1V8H5Zm-1 2h3V9H4v1Zm6-1h1V8h-1v1Zm-1 2h1v-1H9v1Z"/>`
	gbcSolidInnerSVG                         = `<path fill="currentColor" d="M5 6V3h5v2.5a.5.5 0 0 1-.5.5H5Z"/><path fill="currentColor" fill-rule="evenodd" d="M2 1.5A1.5 1.5 0 0 1 3.5 0h8A1.5 1.5 0 0 1 13 1.5v10A3.5 3.5 0 0 1 9.5 15h-6A1.5 1.5 0 0 1 2 13.5v-12Zm2.5.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h5A1.5 1.5 0 0 0 11 5.5v-3a.5.5 0 0 0-.5-.5h-6ZM5 8v1H4v1h1v1h1v-1h1V9H6V8H5Zm5 0v1h1V8h-1Zm-1 3v-1h1v1H9Z" clip-rule="evenodd"/>`
	ghostOutlineInnerSVG                     = `<path fill="currentColor" d="m3.914 14.086l.354.353l-.354-.353ZM4 14l-.354-.354L4 14Zm2.5 0l.354-.354L6.5 14Zm.146.146l-.353.354l.353-.354Zm1.708 0L8 13.793l.354.353ZM8.5 14l.354.354L8.5 14Zm2.5 0l-.354.354L11 14Zm.086.086l-.354.353l.354-.353ZM5 6h1V5H5v1Zm4 0h1V5H9v1Zm1.1 1.7a3.25 3.25 0 0 1-5.2 0l-.8.6c1.7 2.267 5.1 2.267 6.8 0l-.8-.6Zm-5.832 6.74l.086-.086l-.708-.708l-.085.086l.707.707Zm.94-.44h.085v-1h-.086v1Zm.938.354l.147.146l.707-.707l-.146-.147l-.708.708Zm2.561.146l.147-.146l-.708-.708l-.146.147l.707.707Zm1-.5h.086v-1h-.086v1Zm.94.354l.085.085l.707-.707l-.085-.086l-.708.708Zm1.439.646A1.914 1.914 0 0 0 14 13.086h-1c0 .505-.41.914-.914.914v1Zm-1.354-.56c.36.358.846.56 1.354.56v-1a.914.914 0 0 1-.647-.268l-.707.707Zm-.94-.44c.321 0 .628.127.854.354l.708-.708A2.207 2.207 0 0 0 9.793 13v1Zm-.938.354c.226-.227.533-.354.853-.354v-1c-.585 0-1.147.232-1.56.646l.707.708ZM7.5 15c.453 0 .887-.18 1.207-.5L8 13.793a.707.707 0 0 1-.5.207v1Zm-1.207-.5c.32.32.754.5 1.207.5v-1a.707.707 0 0 1-.5-.207l-.707.707Zm-1-.5c.32 0 .627.127.853.354l.708-.708A2.207 2.207 0 0 0 5.293 13v1Zm-.94.354c.227-.227.534-.354.854-.354v-1c-.585 0-1.147.232-1.56.646l.707.708ZM2.915 15c.508 0 .995-.202 1.354-.56l-.707-.708a.914.914 0 0 1-.647.268v1ZM1 13.086C1 14.143 1.857 15 2.914 15v-1A.914.914 0 0 1 2 13.086H1ZM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0v1ZM2 6.5A5.5 5.5 0 0 1 7.5 1V0A6.5 6.5 0 0 0 1 6.5h1Zm-1 0v6.586h1V6.5H1Zm13 6.586V6.5h-1v6.586h1Z"/>`
	ghostSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 0 1 13 0v6.586a1.914 1.914 0 0 1-3.268 1.353l-.086-.085A1.207 1.207 0 0 0 9.793 14h-.086c-.32 0-.627.127-.853.354l-.147.146a1.707 1.707 0 0 1-2.414 0l-.147-.146A1.207 1.207 0 0 0 5.293 14h-.086c-.32 0-.627.127-.853.354l-.086.085A1.914 1.914 0 0 1 1 13.086V6.5ZM5 6h1V5H5v1Zm4 0h1V5H9v1ZM4.9 7.7a3.25 3.25 0 0 0 5.2 0l.8.6c-1.7 2.267-5.1 2.267-6.8 0l.8-.6Z" clip-rule="evenodd"/>`
	gifOutlineInnerSVG                       = `<path fill="currentColor" d="M2.5 10.5H2v.5h.5v-.5Zm2 0v.5H5v-.5h-.5Zm9-7h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5ZM2 6v4.5h1V6H2Zm.5 5h2v-1h-2v1Zm2.5-.5v-2H4v2h1ZM3 7h2V6H3v1ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM6 7h3V6H6v1Zm0 4h3v-1H6v1Zm1-4.5v4h1v-4H7Zm3.5.5H13V6h-2.5v1ZM10 6v5h1V6h-1Zm.5 3H12V8h-1.5v1Z"/>`
	gifSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM2 11V6h3v1H3v3h1V8.5h1V11H2Zm6-4h1V6H6v1h1v3H6v1h3v-1H8V7Zm5-1v1h-2v1h1v1h-1v2h-1V6h3Z" clip-rule="evenodd"/>`
	giftOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M13.5 7.5h-12m12 0a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1m12 0v5a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-5m6-3v-1m0 1H4.214A1.714 1.714 0 0 1 2.5 2.786V2.5a2 2 0 0 1 2-2a3 3 0 0 1 3 3m0 1h3.286c.947 0 1.714-.768 1.714-1.714V2.5a2 2 0 0 0-2-2a3 3 0 0 0-3 3m0 1v10"/>`
	giftSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 0A2.5 2.5 0 0 0 2 2.5v.286c0 .448.133.865.362 1.214H1.5A1.5 1.5 0 0 0 0 5.5v1A1.5 1.5 0 0 0 1.5 8H7V4h1v4h5.5A1.5 1.5 0 0 0 15 6.5v-1A1.5 1.5 0 0 0 13.5 4h-.862c.229-.349.362-.766.362-1.214V2.5A2.5 2.5 0 0 0 10.5 0c-1.273 0-2.388.68-3 1.696A3.498 3.498 0 0 0 4.5 0ZM8 4h2.786C11.456 4 12 3.456 12 2.786V2.5A1.5 1.5 0 0 0 10.5 1A2.5 2.5 0 0 0 8 3.5V4ZM7 4H4.214C3.544 4 3 3.456 3 2.786V2.5A1.5 1.5 0 0 1 4.5 1A2.5 2.5 0 0 1 7 3.5V4Z" clip-rule="evenodd"/><path fill="currentColor" d="M7 9H1v3.5A2.5 2.5 0 0 0 3.5 15H7V9Zm1 6h3.5a2.5 2.5 0 0 0 2.5-2.5V9H8v6Z"/>`
	gitBranchOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M2.5 4.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 0v6m2 2a2 2 0 1 1-2-2m2 2a2 2 0 0 0-2-2m2 2h5a3 3 0 0 0 3-3v-2m0 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`
	gitBranchSolidInnerSVG                   = `<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1A2.5 2.5 0 1 0 4.95 13H9.5A3.5 3.5 0 0 0 13 9.5V7.95a2.5 2.5 0 1 0-1 0V9.5A2.5 2.5 0 0 1 9.5 12H4.95A2.503 2.503 0 0 0 3 10.05v-5.1A2.5 2.5 0 0 0 2.5 0Z"/>`
	gitCommitOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M7.5 10.5a3 3 0 0 1 0-6m0 6a3 3 0 0 0 0-6m0 6V15m0-10.5V0"/>`
	gitCommitSolidInnerSVG                   = `<path fill="currentColor" d="M4 7.5a3.5 3.5 0 0 1 3-3.465V0h1v4.035a3.5 3.5 0 0 1 0 6.93V15H7v-4.035A3.5 3.5 0 0 1 4 7.5Z"/>`
	gitCompareOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="m8.5.5l-2 2m0 0l2 2m-2-2h3a3 3 0 0 1 3 3v5m-10-6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 0v5a3 3 0 0 0 3 3H8m-1.5 2l2-2l-2-2m6 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`
	gitCompareSolidInnerSVG                  = `<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95V9.5A3.5 3.5 0 0 0 5.5 13h1.793l-1.147 1.146l.708.708L9.207 12.5l-2.353-2.354l-.708.708L7.293 12H5.5A2.5 2.5 0 0 1 3 9.5V4.95A2.5 2.5 0 0 0 2.5 0Zm6.354.854L8.146.146L5.793 2.5l2.353 2.354l.708-.708L7.707 3H9.5A2.5 2.5 0 0 1 12 5.5v4.55a2.5 2.5 0 1 0 1 0V5.5A3.5 3.5 0 0 0 9.5 2H7.707L8.854.854Z"/>`
	gitForkOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M2.5 4.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 0v6m0 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm10-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 0v1a3 3 0 0 1-3 3h-7"/>`
	gitForkSolidInnerSVG                     = `<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1a2.5 2.5 0 1 0 1 0V9h6.5A3.5 3.5 0 0 0 13 5.5v-.55a2.5 2.5 0 1 0-1 0v.55A2.5 2.5 0 0 1 9.5 8H3V4.95A2.5 2.5 0 0 0 2.5 0Z"/>`
	gitMergeOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M2.5 10.5a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0v-6m2-2a2 2 0 1 0-2 2m2-2a2 2 0 0 1-2 2m2-2h5a3 3 0 0 1 3 3v2m0 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`
	gitMergeSolidInnerSVG                    = `<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1a2.5 2.5 0 1 0 1 0v-5.1A2.503 2.503 0 0 0 4.95 3H9.5A2.5 2.5 0 0 1 12 5.5v1.55a2.5 2.5 0 1 0 1 0V5.5A3.5 3.5 0 0 0 9.5 2H4.95A2.5 2.5 0 0 0 2.5 0Z"/>`
	gitOutlineInnerSVG                       = `<path fill="currentColor" d="m6.793 1.207l.353.354l-.353-.354ZM1.207 6.793l-.353-.354l.353.354Zm0 1.414l.354-.353l-.354.353Zm5.586 5.586l-.354.353l.354-.353Zm1.414 0l-.353-.354l.353.354Zm5.586-5.586l.353.354l-.353-.354Zm0-1.414l-.354.353l.354-.353ZM8.207 1.207l.354-.353l-.354.353ZM6.44.854L.854 6.439l.707.707l5.585-5.585L6.44.854ZM.854 8.56l5.585 5.585l.707-.707l-5.585-5.585l-.707.707Zm7.707 5.585l5.585-5.585l-.707-.707l-5.585 5.585l.707.707Zm5.585-7.707L8.561.854l-.707.707l5.585 5.585l.707-.707Zm0 2.122a1.5 1.5 0 0 0 0-2.122l-.707.707a.5.5 0 0 1 0 .708l.707.707ZM6.44 14.146a1.5 1.5 0 0 0 2.122 0l-.707-.707a.5.5 0 0 1-.708 0l-.707.707ZM.854 6.44a1.5 1.5 0 0 0 0 2.122l.707-.707a.5.5 0 0 1 0-.708L.854 6.44Zm6.292-4.878a.5.5 0 0 1 .708 0L8.56.854a1.5 1.5 0 0 0-2.122 0l.707.707Zm-2 1.293l1 1l.708-.708l-1-1l-.708.708ZM7.5 5a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 6V5Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 4.5H8ZM7.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 3v1Zm0-1A1.5 1.5 0 0 0 6 4.5h1a.5.5 0 0 1 .5-.5V3Zm.646 2.854l1.5 1.5l.707-.708l-1.5-1.5l-.707.708ZM10.5 8a.5.5 0 0 1-.5-.5H9A1.5 1.5 0 0 0 10.5 9V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 12 7.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 6v1Zm0-1A1.5 1.5 0 0 0 9 7.5h1a.5.5 0 0 1 .5-.5V6ZM7 5.5v4h1v-4H7Zm.5 5.5a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 12v-1Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 10.5H8Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 9v1Zm0-1A1.5 1.5 0 0 0 6 10.5h1a.5.5 0 0 1 .5-.5V9Z"/>`
	gitPullOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="m8.5.5l-2 2m0 0l2 2m-2-2h3a3 3 0 0 1 3 3v2m-10 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0v-6m0 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`
	gitPullSolidInnerSVG                     = `<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1a2.5 2.5 0 1 0 1 0v-5.1A2.5 2.5 0 0 0 2.5 0Zm6.354.854L8.146.146L5.793 2.5l2.353 2.354l.708-.708L7.707 3H9.5A2.5 2.5 0 0 1 12 5.5v1.55a2.5 2.5 0 1 0 1 0V5.5A3.5 3.5 0 0 0 9.5 2H7.707L8.854.854Z"/>`
	gitSolidInnerSVG                         = `<path fill="currentColor" d="M6.44.854a1.5 1.5 0 0 1 2.12 0l5.586 5.585a1.5 1.5 0 0 1 0 2.122l-5.585 5.585a1.5 1.5 0 0 1-2.122 0L.854 8.561a1.5 1.5 0 0 1 0-2.122L4.793 2.5l1.353 1.353A1.5 1.5 0 0 0 7 5.914v3.171a1.5 1.5 0 1 0 1 0v-3.17c.05-.018.1-.038.147-.061l1 1a1.5 1.5 0 1 0 .707-.707l-1-1a1.5 1.5 0 0 0-2-2L5.5 1.792l.94-.94Z"/>`
	githubOutlineInnerSVG                    = `<path fill="currentColor" d="M5.65 12.477a.5.5 0 1 0-.3-.954l.3.954Zm-3.648-2.96l-.484-.128l-.254.968l.484.127l.254-.968ZM9 14.5v.5h1v-.5H9Zm.063-4.813l-.054-.497a.5.5 0 0 0-.299.852l.352-.354ZM12.5 5.913h.5V5.91l-.5.002Zm-.833-2.007l-.466-.18a.5.5 0 0 0 .112.533l.354-.353Zm-.05-2.017l.456-.204a.5.5 0 0 0-.319-.276l-.137.48Zm-2.173.792l-.126.484a.5.5 0 0 0 .398-.064l-.272-.42Zm-3.888 0l-.272.42a.5.5 0 0 0 .398.064l-.126-.484ZM3.383 1.89l-.137-.48a.5.5 0 0 0-.32.276l.457.204Zm-.05 2.017l.354.353a.5.5 0 0 0 .112-.534l-.466.181ZM2.5 5.93H3v-.002l-.5.002Zm3.438 3.758l.352.355a.5.5 0 0 0-.293-.851l-.06.496ZM5.5 11H6l-.001-.037L5.5 11ZM5 14.5v.5h1v-.5H5Zm.35-2.977c-.603.19-.986.169-1.24.085c-.251-.083-.444-.25-.629-.49a4.8 4.8 0 0 1-.27-.402c-.085-.139-.182-.302-.28-.447c-.191-.281-.473-.633-.929-.753l-.254.968c.08.02.184.095.355.346c.082.122.16.252.258.412c.094.152.202.32.327.484c.253.33.598.663 1.11.832c.51.168 1.116.15 1.852-.081l-.3-.954Zm4.65-.585c0-.318-.014-.608-.104-.878c-.096-.288-.262-.51-.481-.727l-.705.71c.155.153.208.245.237.333c.035.105.053.254.053.562h1Zm-.884-.753c.903-.097 1.888-.325 2.647-.982c.78-.675 1.237-1.729 1.237-3.29h-1c0 1.359-.39 2.1-.892 2.534c-.524.454-1.258.653-2.099.743l.107.995ZM13 5.91a3.354 3.354 0 0 0-.98-2.358l-.707.706c.438.44.685 1.034.687 1.655l1-.003Zm-.867-1.824c.15-.384.22-.794.21-1.207l-1 .025a2.12 2.12 0 0 1-.142.82l.932.362Zm.21-1.207a3.119 3.119 0 0 0-.27-1.195l-.913.408c.115.256.177.532.184.812l1-.025Zm-.726-.99c.137-.481.137-.482.136-.482h-.003l-.004-.002a.462.462 0 0 0-.03-.007a1.261 1.261 0 0 0-.212-.024a2.172 2.172 0 0 0-.51.054c-.425.091-1.024.317-1.82.832l.542.84c.719-.464 1.206-.634 1.488-.694a1.2 1.2 0 0 1 .306-.03l-.008-.001a.278.278 0 0 1-.01-.002l-.006-.002h-.003l-.002-.001c-.001 0-.002 0 .136-.482Zm-2.047.307a8.209 8.209 0 0 0-4.14 0l.252.968a7.209 7.209 0 0 1 3.636 0l.252-.968Zm-3.743.064c-.797-.514-1.397-.74-1.822-.83a2.17 2.17 0 0 0-.51-.053a1.259 1.259 0 0 0-.241.03l-.004.002h-.003l.136.481l.137.481h-.001l-.002.001l-.003.001a.327.327 0 0 1-.016.004l-.008.001h.008a1.19 1.19 0 0 1 .298.03c.282.06.769.23 1.488.694l.543-.84Zm-2.9-.576a3.12 3.12 0 0 0-.27 1.195l1 .025a2.09 2.09 0 0 1 .183-.812l-.913-.408Zm-.27 1.195c-.01.413.06.823.21 1.207l.932-.362a2.12 2.12 0 0 1-.143-.82l-1-.025Zm.322.673a3.354 3.354 0 0 0-.726 1.091l.924.38c.118-.285.292-.545.51-.765l-.708-.706Zm-.726 1.091A3.354 3.354 0 0 0 2 5.93l1-.003c0-.31.06-.616.177-.902l-.924-.38ZM2 5.93c0 1.553.458 2.597 1.239 3.268c.757.65 1.74.88 2.64.987l.118-.993C5.15 9.09 4.416 8.89 3.89 8.438C3.388 8.007 3 7.276 3 5.928H2Zm3.585 3.404c-.5.498-.629 1.09-.584 1.704L6 10.963c-.03-.408.052-.683.291-.921l-.705-.709ZM5 11v3.5h1V11H5Zm5 3.5V13H9v1.5h1Zm0-1.5v-2.063H9V13h1Z"/>`
	githubSolidInnerSVG                      = `<path fill="currentColor" d="M9.358 2.145a8.209 8.209 0 0 0-3.716 0c-.706-.433-1.245-.632-1.637-.716a2.17 2.17 0 0 0-.51-.053a1.258 1.258 0 0 0-.232.028l-.01.002l-.004.002h-.003l.137.481l-.137-.48a.5.5 0 0 0-.32.276a3.12 3.12 0 0 0-.159 2.101A3.354 3.354 0 0 0 2 5.93c0 1.553.458 2.597 1.239 3.268c.547.47 1.211.72 1.877.863a2.34 2.34 0 0 0-.116.958v.598c-.407.085-.689.058-.89-.008c-.251-.083-.444-.25-.629-.49a4.798 4.798 0 0 1-.27-.402l-.057-.093a9.216 9.216 0 0 0-.224-.354c-.19-.281-.472-.633-.928-.753l-.484-.127l-.254.968l.484.127c.08.02.184.095.355.346a7.2 7.2 0 0 1 .19.302l.068.11c.094.152.202.32.327.484c.253.33.598.663 1.11.832c.35.116.748.144 1.202.074V14.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-3.563c0-.315-.014-.604-.103-.873c.663-.14 1.322-.39 1.866-.86c.78-.676 1.237-1.73 1.237-3.292v-.001a3.354 3.354 0 0 0-.768-2.125a3.12 3.12 0 0 0-.159-2.1a.5.5 0 0 0-.319-.277l-.137.48c.137-.48.136-.48.135-.48l-.002-.001l-.004-.002l-.009-.002a.671.671 0 0 0-.075-.015a1.261 1.261 0 0 0-.158-.013a2.172 2.172 0 0 0-.51.053c-.391.084-.93.283-1.636.716Z"/>`
	gitlabOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m.5 8.5l7 6l7-6l-2-8l-2 6h-6l-2-6l-2 8Z"/>`
	gitlabSolidInnerSVG                      = `<path fill="currentColor" d="M2.974.342a.5.5 0 0 0-.96.037l-2 8a.5.5 0 0 0 .16.5l7 6a.5.5 0 0 0 .651 0l7-6a.5.5 0 0 0 .16-.5l-2-8a.5.5 0 0 0-.96-.037L10.14 6H4.86L2.974.342Z"/>`
	globeAfricaOutlineInnerSVG               = `<path fill="currentColor" d="M7.5 1H8a.5.5 0 0 0-.276-.447L7.5 1Zm0 1l.224.447A.5.5 0 0 0 8 2h-.5Zm-1 .5l-.224-.447a.5.5 0 0 0-.053.863L6.5 2.5Zm1.5 1l-.277.416A.5.5 0 0 0 8 4v-.5Zm.5 0V4a.5.5 0 0 0 .5-.5h-.5Zm0-1V2a.5.5 0 0 0-.5.5h.5Zm1.5 0h.5A.5.5 0 0 0 10 2v.5Zm0 1h-.5a.5.5 0 0 0 .146.354L10 3.5Zm.5.5h.5a.5.5 0 0 0-.146-.354L10.5 4Zm0 1l.354.354A.5.5 0 0 0 11 5h-.5Zm-.5.5V6a.5.5 0 0 0 .354-.146L10 5.5Zm-1 0l-.224.447A.5.5 0 0 0 9 6v-.5Zm-2-1l.224-.447a.5.5 0 0 0-.448 0L7 4.5ZM6 5v.5a.5.5 0 0 0 .224-.053L6 5ZM4.5 5v-.5a.5.5 0 0 0-.485.379L4.5 5ZM4 7l-.485-.121a.5.5 0 0 0 .131.475L4 7Zm1.5 1.5l-.354.354a.5.5 0 0 0 .13.093L5.5 8.5Zm1 .5H7a.5.5 0 0 0-.276-.447L6.5 9Zm0 1.5H6a.5.5 0 0 0 .146.354L6.5 10.5Zm.5.5h.5a.5.5 0 0 0-.146-.354L7 11Zm0 1h-.5a.5.5 0 0 0 .053.224L7 12Zm.5 1l-.447.224a.5.5 0 0 0 .447.276V13Zm1 0v.5a.5.5 0 0 0 .416-.223L8.5 13Zm1-1.5l.416.277a.5.5 0 0 0 .031-.053L9.5 11.5Zm.5-1l.447.224a.5.5 0 0 0 .053-.224H10Zm0-1l-.4-.3a.5.5 0 0 0-.1.3h.5Zm1.5-2l.4.3a.5.5 0 0 0 .047-.524L11.5 7.5Zm-.5-1V6a.5.5 0 0 0-.447.724L11 6.5Zm1.5 0V7a.5.5 0 0 0 .5-.5h-.5Zm0-1V5a.5.5 0 0 0-.5.5h.5Zm-5 8.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM6.276.947l1 .5l.448-.894l-1-.5l-.448.894ZM7 1v1h1V1H7Zm.276.553l-1 .5l.448.894l1-.5l-.448-.894ZM6.223 2.916l1.5 1l.554-.832l-1.5-1l-.554.832ZM8 4h.5V3H8v1Zm1-.5v-1H8v1h1ZM8.5 3H10V2H8.5v1Zm1-.5v1h1v-1h-1Zm.146 1.354l.5.5l.708-.708l-.5-.5l-.708.708ZM10 4v1h1V4h-1Zm.146.646l-.5.5l.708.708l.5-.5l-.708-.708ZM10 5H9v1h1V5Zm-.776.053l-2-1l-.448.894l2 1l.448-.894Zm-2.448-1l-1 .5l.448.894l1-.5l-.448-.894ZM6 4.5H4.5v1H6v-1Zm-1.985.379l-.5 2l.97.242l.5-2l-.97-.242Zm-.369 2.475l1.5 1.5l.708-.708l-1.5-1.5l-.708.708Zm1.63 1.593l1 .5l.448-.894l-1-.5l-.448.894ZM6 9v1.5h1V9H6Zm.146 1.854l.5.5l.708-.708l-.5-.5l-.708.708ZM6.5 11v1h1v-1h-1Zm.053 1.224l.5 1l.894-.448l-.5-1l-.894.448ZM7.5 13.5h1v-1h-1v1Zm1.416-.223l1-1.5l-.832-.554l-1 1.5l.832.554Zm1.031-1.553l.5-1l-.894-.448l-.5 1l.894.448ZM10.5 10.5v-1h-1v1h1Zm-.1-.7l1.5-2l-.8-.6l-1.5 2l.8.6Zm1.547-2.524l-.5-1l-.894.448l.5 1l.894-.448ZM11 7h1.5V6H11v1Zm2-.5v-1h-1v1h1Zm-.5-.5h2V5h-2v1Z"/>`
	globeAfricaSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm1 0a6.502 6.502 0 0 1 5.527-6.428L7 1.31v.382l-.724.362a.5.5 0 0 0-.053.863l1.5 1A.5.5 0 0 0 8 4h.5a.5.5 0 0 0 .5-.5V3h.5v.5a.5.5 0 0 0 .146.354l.354.353v.586L9.793 5h-.675l-1.894-.947a.5.5 0 0 0-.448 0l-.894.447H4.5a.5.5 0 0 0-.485.379l-.5 2a.5.5 0 0 0 .131.475l1.5 1.5a.5.5 0 0 0 .13.093L6 9.31v1.19a.5.5 0 0 0 .146.354l.354.353V12a.5.5 0 0 0 .053.224l.5 1a.5.5 0 0 0 .447.276h1a.5.5 0 0 0 .416-.223l1-1.5a.5.5 0 0 0 .031-.053l.5-1a.5.5 0 0 0 .053-.224v-.833L11.9 7.8a.5.5 0 0 0 .047-.524L11.81 7h.691a.5.5 0 0 0 .5-.5V6h.826A6.5 6.5 0 1 1 1 7.5Z" clip-rule="evenodd"/>`
	globeAmericasOutlineInnerSVG             = `<path fill="currentColor" d="M4.5 5.5V5a.5.5 0 0 0-.5.5h.5Zm0 1l.354.354L5 6.707V6.5h-.5Zm-1.707.293l-.354.353l.354-.353ZM6.5 13H7v-.207l-.146-.147L6.5 13Zm-1-1H5v.207l.146.147L5.5 12Zm0-1.5H6v-.207l-.146-.147l-.354.354Zm-1-1H4v.207l.146.147L4.5 9.5Zm0-1V8a.5.5 0 0 0-.5.5h.5ZM9 .5v2h1v-2H9ZM8.5 3h-1v1h1V3Zm-3 2h-1v1h1V5ZM4 5.5v1h1v-1H4Zm.146.646l-.292.293l.707.707l.293-.292l-.708-.708Zm-1 .293L1.354 4.646l-.708.708L2.44 7.146l.707-.707ZM6 4.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 7 4.5H6ZM7.5 3A1.5 1.5 0 0 0 6 4.5h1a.5.5 0 0 1 .5-.5V3ZM3.854 6.44a.5.5 0 0 1-.708 0l-.707.706a1.5 1.5 0 0 0 2.122 0l-.707-.707ZM9 2.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 10 2.5H9Zm-2 12V13H6v1.5h1Zm-.146-1.854l-1-1l-.708.708l1 1l.708-.708ZM6 12v-1.5H5V12h1Zm-.146-1.854l-1-1l-.708.708l1 1l.708-.708ZM5 9.5v-1H4v1h1ZM4.5 9h4V8h-4v1Zm4.5.5v.667h1V9.5H9Zm1.833 2.5H13.5v-1h-2.667v1ZM10 11.167c0 .46.373.833.833.833v-1c.092 0 .167.075.167.167h-1ZM9.833 11c.092 0 .167.075.167.167h1C11 10.522 10.478 10 9.833 10v1ZM9 10.167c0 .46.373.833.833.833v-1c.092 0 .167.075.167.167H9ZM8.5 9a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 8.5 8v1Zm-1 5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	globeAmericasSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM1.197 5.904A6.503 6.503 0 0 0 6 13.826v-.619l-1-1v-1.5l-1-1V8.5a.5.5 0 0 1 .5-.5h4A1.5 1.5 0 0 1 10 9.5v.512c.51.073.915.477.988.988h1.99A6.502 6.502 0 0 0 10 1.498V2.5A1.5 1.5 0 0 1 8.5 4h-1a.5.5 0 0 0-.5.5A1.5 1.5 0 0 1 5.5 6H5v.707l-.44.44a1.5 1.5 0 0 1-2.12 0L1.196 5.903Z" clip-rule="evenodd"/>`
	globeOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M7.5 15v-3M4 14.5h7M11.469 1A6 6 0 1 1 3.5 9.972m4-7.472a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/>`
	globeSolidInnerSVG                       = `<path fill="currentColor" d="M14 5.5A6.486 6.486 0 0 0 11.8.625l-.662.75A5.5 5.5 0 1 1 3.834 9.6l-.667.745A6.476 6.476 0 0 0 7 11.98V14H4v1h7v-1H8v-2.019A6.5 6.5 0 0 0 14 5.5Z"/><path fill="currentColor" d="M7.5 2a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Z"/>`
	googleAdOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M3.5 10V7a1.5 1.5 0 1 1 3 0v3m5-5v5m0-2.5h-2a1 1 0 0 0 0 2h2m-8-1h3m-4-6h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2Z"/>`
	googleAdSolidInnerSVG                    = `<path fill="currentColor" d="M5 6a1 1 0 0 0-1 1v1h2V7a1 1 0 0 0-1-1Zm6 2H9.5a.5.5 0 0 0 0 1H11V8Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 4.5A2.5 2.5 0 0 1 2.5 2h10A2.5 2.5 0 0 1 15 4.5v6a2.5 2.5 0 0 1-2.5 2.5h-10A2.5 2.5 0 0 1 0 10.5v-6ZM4 10V9h2v1h1V7a2 2 0 1 0-4 0v3h1Zm7-3H9.5a1.5 1.5 0 1 0 0 3H12V5h-1v2Z" clip-rule="evenodd"/>`
	googleDriveOutlineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m5 1.5l-4.5 8l2 4M5 1.5l2.5 4l-5 8M5 1.5h5l4.5 8M5 1.5l5 8h4.5m-12 4l2.5-4h9.5m-12 4h10l2-4"/>`
	googleDriveSolidInnerSVG                 = `<path fill="currentColor" d="M4.563 1.258A.5.5 0 0 1 5 1h5a.5.5 0 0 1 .436.255L14.23 8H8.473L4.576 1.765a.5.5 0 0 1-.013-.507ZM3.91 14h8.59a.5.5 0 0 0 .447-.276l2-4A.5.5 0 0 0 14.5 9H7.092l-3.181 5ZM.064 9.255l3.792-6.742l3.05 4.88l-3.982 6.372a.5.5 0 0 1-.871-.041l-2-4a.5.5 0 0 1 .011-.47Z"/>`
	googleOutlineInnerSVG                    = `<path fill="currentColor" d="M14.5 7.5h.5V7h-.5v.5Zm-7 6.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Zm0 1c1.819 0 3.463.746 4.643 1.951l.714-.7A7.479 7.479 0 0 0 7.5 0v1ZM8 8h6.5V7H8v1Z"/>`
	googlePlayStoreOutlineInnerSVG           = `<path fill="currentColor" d="m1.5.5l.252-.432A.5.5 0 0 0 1 .5h.5Zm0 14H1a.5.5 0 0 0 .752.432L1.5 14.5Zm12-7l.252.432a.5.5 0 0 0 0-.864L13.5 7.5ZM1 .5v14h1V.5H1Zm.752 14.432l12-7l-.504-.864l-12 7l.504.864Zm12-7.864l-12-7l-.504.864l12 7l.504-.864ZM1.829 12.876l8-7l-.658-.752l-8 7l.658.752Zm-.658-10l8 7l.658-.752l-8-7l-.658.752Z"/>`
	googlePlayStoreSolidInnerSVG             = `<path fill="currentColor" d="M1.251.066a.5.5 0 0 1 .5.002l7.843 4.575l-2.427 2.184L1 1.277V.5a.5.5 0 0 1 .251-.434ZM1 2.623v9.754L6.42 7.5L1 2.623Zm0 11.1v.777a.5.5 0 0 0 .752.432l7.842-4.575l-2.427-2.184L1 13.723Zm9.501-3.895l3.25-1.896a.5.5 0 0 0 0-.864l-3.25-1.896L7.914 7.5l2.587 2.328Z"/>`
	googleSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 0 1 12.857-5.249l-.714.7A6.5 6.5 0 1 0 13.98 8H8V7h7v.5a7.5 7.5 0 0 1-15 0Z" clip-rule="evenodd"/>`
	googleStreetviewOutlineInnerSVG          = `<path fill="currentColor" d="M4.5 11.5H4a.5.5 0 0 0 .5.5v-.5Zm6 0v.5a.5.5 0 0 0 .5-.5h-.5Zm-6 .5h6v-1h-6v1Zm6.5-.5v-2h-1v2h1Zm-7-2v2h1v-2H4ZM7.5 6A3.5 3.5 0 0 0 4 9.5h1A2.5 2.5 0 0 1 7.5 7V6ZM11 9.5A3.5 3.5 0 0 0 7.5 6v1A2.5 2.5 0 0 1 10 9.5h1Zm3 2c0 .245-.114.52-.406.816c-.294.299-.745.59-1.341.846c-1.191.51-2.871.838-4.753.838v1c1.984 0 3.804-.344 5.147-.92c.67-.287 1.245-.642 1.659-1.061c.416-.422.694-.936.694-1.519h-1ZM7.5 14c-1.882 0-3.562-.328-4.753-.838c-.596-.256-1.047-.547-1.341-.846C1.114 12.02 1 11.746 1 11.5H0c0 .583.278 1.097.694 1.519c.414.42.989.774 1.66 1.062C3.695 14.656 5.515 15 7.5 15v-1ZM1 11.5c0-.242.11-.513.394-.805c.286-.294.725-.582 1.306-.837l-.4-.916c-.656.287-1.218.64-1.622 1.056C.27 10.416 0 10.925 0 11.5h1Zm11.3-1.642c.581.255 1.02.543 1.305.837c.284.292.395.563.395.805h1c0-.575-.27-1.084-.678-1.502c-.404-.416-.966-.769-1.622-1.056l-.4.916ZM7.5 4A1.5 1.5 0 0 1 6 2.5H5A2.5 2.5 0 0 0 7.5 5V4ZM9 2.5A1.5 1.5 0 0 1 7.5 4v1A2.5 2.5 0 0 0 10 2.5H9ZM7.5 1A1.5 1.5 0 0 1 9 2.5h1A2.5 2.5 0 0 0 7.5 0v1Zm0-1A2.5 2.5 0 0 0 5 2.5h1A1.5 1.5 0 0 1 7.5 1V0Z"/>`
	googleStreetviewSolidInnerSVG            = `<path fill="currentColor" d="M5 2.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0ZM7.5 6A3.5 3.5 0 0 0 4 9.5v2a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-2A3.5 3.5 0 0 0 7.5 6Z"/><path fill="currentColor" d="M1.394 10.695c-.283.292-.394.563-.394.805c0 .245.114.52.406.816c.294.299.745.59 1.341.846c1.191.51 2.871.838 4.753.838c1.882 0 3.562-.328 4.753-.838c.596-.256 1.047-.547 1.341-.846c.292-.296.406-.57.406-.816c0-.242-.11-.513-.395-.805c-.285-.294-.724-.582-1.305-.837l.4-.916c.656.287 1.218.64 1.622 1.056c.407.418.678.927.678 1.502c0 .583-.278 1.097-.694 1.519c-.414.42-.989.774-1.66 1.062c-1.342.575-3.162.919-5.146.919s-3.804-.344-5.147-.92c-.67-.287-1.245-.642-1.659-1.061c-.416-.422-.694-.936-.694-1.52c0-.575.27-1.083.678-1.501c.404-.416.966-.769 1.622-1.056l.4.916c-.581.255-1.02.543-1.306.837Z"/>`
	graphqlOutlineInnerSVG                   = `<path fill="currentColor" d="m2.61 4.432l4.142-2.417l-.504-.864l-4.143 2.417l.504.864ZM2 9.5v-4H1v4h1Zm6.248-7.485l4.143 2.417l.504-.864l-4.143-2.417l-.504.864ZM13 5.5v4h1v-4h-1Zm-.252 4.86l-4.5 2.625l.504.864l4.5-2.625l-.504-.864Zm-5.996 2.625l-4.143-2.417l-.504.864l4.143 2.417l.504-.864ZM6.584 1.973l-5 7.5l.832.554l5-7.5l-.832-.554Zm6.832 7.5l-5-7.5l-.832.554l5 7.5l.832-.554ZM2.5 11h10v-1h-10v1Zm5-9a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 3V2Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 1.5H8ZM7.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 0v1Zm0-1A1.5 1.5 0 0 0 6 1.5h1a.5.5 0 0 1 .5-.5V0Zm6 5a.5.5 0 0 1-.5-.5h-1A1.5 1.5 0 0 0 13.5 6V5Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 15 4.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 3v1Zm0-1A1.5 1.5 0 0 0 12 4.5h1a.5.5 0 0 1 .5-.5V3Zm0 8a.5.5 0 0 1-.5-.5h-1a1.5 1.5 0 0 0 1.5 1.5v-1Zm.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 9v1Zm0-1a1.5 1.5 0 0 0-1.5 1.5h1a.5.5 0 0 1 .5-.5V9Zm-6 5a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 15v-1Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 13.5H8Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 12v1Zm0-1A1.5 1.5 0 0 0 6 13.5h1a.5.5 0 0 1 .5-.5v-1Zm-6-1a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 12v-1Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 3 10.5H2Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 1.5 9v1Zm0-1A1.5 1.5 0 0 0 0 10.5h1a.5.5 0 0 1 .5-.5V9Zm0-4a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 6V5Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 3 4.5H2ZM1.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 1.5 3v1Zm0-1A1.5 1.5 0 0 0 0 4.5h1a.5.5 0 0 1 .5-.5V3Z"/>`
	graphqlSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M6.015 1.287a1.5 1.5 0 0 1 2.97 0l3.545 2.069A1.5 1.5 0 1 1 14 5.916v3.17a1.5 1.5 0 1 1-1.47 2.559l-3.545 2.068a1.5 1.5 0 0 1-2.97 0L2.47 11.645A1.5 1.5 0 1 1 1 9.085v-3.17a1.5 1.5 0 1 1 1.47-2.56l3.545-2.068Zm.225 1.027L2.974 4.219A1.5 1.5 0 0 1 2 5.914V8.85L6.3 2.4a1.5 1.5 0 0 1-.06-.085Zm.891.64l-4.43 6.647c.09.12.163.254.214.399h9.17a1.51 1.51 0 0 1 .214-.4L7.87 2.955a1.503 1.503 0 0 1-.738 0Zm1.57-.555L13 8.85V5.915a1.5 1.5 0 0 1-.974-1.696L8.76 2.314a1.5 1.5 0 0 1-.06.085ZM11.65 11H3.349l2.89 1.686a1.499 1.499 0 0 1 2.521 0L11.65 11ZM7.5 1a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-6 3a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm12 0a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-12 6a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm12 0a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-6 3a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`
	gridLayoutOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M5.5.5h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Zm8 0h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Zm0 8h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Zm-8 0h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Z"/>`
	gridLayoutSolidInnerSVG                  = `<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v4A1.5 1.5 0 0 0 1.5 7h4A1.5 1.5 0 0 0 7 5.5v-4A1.5 1.5 0 0 0 5.5 0h-4Zm8 0A1.5 1.5 0 0 0 8 1.5v4A1.5 1.5 0 0 0 9.5 7h4A1.5 1.5 0 0 0 15 5.5v-4A1.5 1.5 0 0 0 13.5 0h-4Zm-8 8A1.5 1.5 0 0 0 0 9.5v4A1.5 1.5 0 0 0 1.5 15h4A1.5 1.5 0 0 0 7 13.5v-4A1.5 1.5 0 0 0 5.5 8h-4Zm8 0A1.5 1.5 0 0 0 8 9.5v4A1.5 1.5 0 0 0 9.5 15h4a1.5 1.5 0 0 0 1.5-1.5v-4A1.5 1.5 0 0 0 13.5 8h-4Z"/>`
	hashtagOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M2 5.5h11m-11 4h11m-6.5-8l-2 12m6-12l-2 12"/>`
	hashtagSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="m5.41 5l.597-3.582l.986.164L6.423 5H9.41l.597-3.582l.986.164L10.423 5H13v1h-2.743l-.5 3H13v1H9.59l-.597 3.582l-.986-.164l.57-3.418H5.59l-.597 3.582l-.986-.164l.57-3.418H2V9h2.743l.5-3H2V5h3.41Zm.847 1l-.5 3h2.986l.5-3H6.257Z" clip-rule="evenodd"/>`
	hdScreenOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M2 13.5h11m-5.5-3V14m6.69-2.589a24.35 24.35 0 0 0-13.38 0a.243.243 0 0 1-.31-.234V1.823c0-.162.155-.279.31-.234a24.35 24.35 0 0 0 13.38 0a.243.243 0 0 1 .31.234v9.354a.243.243 0 0 1-.31.234Z"/>`
	hdScreenSolidInnerSVG                    = `<path fill="currentColor" d="M.948 1.108A.744.744 0 0 0 0 1.823v9.354c0 .494.473.85.948.715A23.85 23.85 0 0 1 7 10.98V13H2v1h11v-1H8v-2.02c2.039.042 4.073.346 6.052.912a.744.744 0 0 0 .948-.715V1.823a.744.744 0 0 0-.948-.715a23.85 23.85 0 0 1-13.104 0Z"/>`
	hdmiCableOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M3.5 4V.5h8V4M5 2.5h1m3 0h1M5.5 13v2m4-2v2m-7-10.5h10v5l-2 1v2h-6v-2l-2-1v-5Z"/>`
	hdmiCableSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M3 0h9v5H3V0Zm3 3H5V2h1v1Zm4 0H9V2h1v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M2 6h11v3.809l-2 1V13h-1v2H9v-2H6v2H5v-2H4v-2.191l-2-1V6Z"/>`
	headphonesOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M1.5 9.5h1a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1Zm0 0v-3a6 6 0 1 1 12 0v3m0 0h-1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1Z"/>`
	headphonesSolidInnerSVG                  = `<path fill="currentColor" d="M2 6.5a5.5 5.5 0 1 1 11 0V9h-.5a1.5 1.5 0 0 0-1.5 1.5v3a1.5 1.5 0 0 0 1.5 1.5h1a1.5 1.5 0 0 0 1.5-1.5v-3a1.5 1.5 0 0 0-1-1.415V6.5a6.5 6.5 0 1 0-13 0v2.585A1.5 1.5 0 0 0 0 10.5v3A1.5 1.5 0 0 0 1.5 15h1A1.5 1.5 0 0 0 4 13.5v-3A1.5 1.5 0 0 0 2.5 9H2V6.5Z"/>`
	headsetOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M12.5 12.5a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1Zm0 0a2 2 0 0 1-2 2H8m6.5-4.5V7.5a7 7 0 1 0-14 0V10m2 2.5a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2Z"/>`
	headsetSolidInnerSVG                     = `<path fill="currentColor" d="M2.5 6a2.49 2.49 0 0 0-1.414.438a6.502 6.502 0 0 1 12.828 0A2.488 2.488 0 0 0 12.5 6A1.5 1.5 0 0 0 11 7.5v4a1.5 1.5 0 0 0 .947 1.395A1.5 1.5 0 0 1 10.5 14H8v1h2.5a2.5 2.5 0 0 0 2.458-2.042A2.5 2.5 0 0 0 15 10.5v-3a7.5 7.5 0 0 0-15 0v3A2.5 2.5 0 0 0 2.5 13A1.5 1.5 0 0 0 4 11.5v-4A1.5 1.5 0 0 0 2.5 6Z"/>`
	heartCircleOutlineInnerSVG               = `<g fill="none" stroke="currentColor"><path d="m6.5 5.5l1 1l1-1a1.414 1.414 0 0 1 2 2l-3 3l-3-3a1.414 1.414 0 0 1 2-2Z"/><path d="M.5 7.5a7 7 0 1 0 14 0a7 7 0 0 0-14 0Z"/></g>`
	heartCircleSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm4.146-2.354a1.914 1.914 0 0 1 2.707 0l.647.647l.646-.647a1.914 1.914 0 0 1 2.707 2.708L7.5 11.207L4.146 7.854a1.914 1.914 0 0 1 0-2.708Z" clip-rule="evenodd"/>`
	heartOutlineInnerSVG                     = `<path fill="currentColor" d="m7.5 13.5l-.354.354a.5.5 0 0 0 .708 0L7.5 13.5ZM1.536 7.536l-.354.353l.354-.353Zm5-5l-.354.353l.354-.353ZM7.5 3.5l-.354.354a.5.5 0 0 0 .708 0L7.5 3.5Zm.964-.964l-.353-.354l.353.354Zm-.61 10.61L1.889 7.182l-.707.707l5.964 5.965l.708-.708Zm5.257-5.964l-5.965 5.964l.708.708l5.964-5.965l-.707-.707ZM6.182 2.889l.964.965l.708-.708l-.965-.964l-.707.707Zm1.672.965l.964-.965l-.707-.707l-.965.964l.708.708ZM10.964 1c-1.07 0-2.096.425-2.853 1.182l.707.707A3.037 3.037 0 0 1 10.964 2V1ZM14 5.036c0 .805-.32 1.577-.89 2.146l.708.707A4.036 4.036 0 0 0 15 5.036h-1Zm1 0A4.036 4.036 0 0 0 10.964 1v1A3.036 3.036 0 0 1 14 5.036h1ZM4.036 2c.805 0 1.577.32 2.146.89l.707-.708A4.036 4.036 0 0 0 4.036 1v1ZM1 5.036A3.036 3.036 0 0 1 4.036 2V1A4.036 4.036 0 0 0 0 5.036h1Zm.89 2.146A3.035 3.035 0 0 1 1 5.036H0c0 1.07.425 2.096 1.182 2.853l.707-.707Z"/>`
	heartSmallOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="m6.5 5.5l1 1l1-1a1.414 1.414 0 1 1 2 2l-3 3l-3-3a1.414 1.414 0 1 1 2-2Z"/>`
	heartSmallSolidInnerSVG                  = `<path fill="currentColor" d="M6.765 5.235a1.79 1.79 0 1 0-2.53 2.53L7.5 11.03l3.265-3.265a1.79 1.79 0 0 0-2.53-2.53L7.5 5.97l-.735-.735Z"/>`
	heartSolidInnerSVG                       = `<path fill="currentColor" d="M4.036 1a4.036 4.036 0 0 0-2.854 6.89l5.964 5.964a.5.5 0 0 0 .708 0l5.964-5.965a4.036 4.036 0 0 0-5.707-5.707l-.611.61l-.61-.61A4.036 4.036 0 0 0 4.035 1Z"/>`
	hexagonOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M1.5 4.5v6l6 3.5l6-3.5v-6L7.5 1l-6 3.5Z"/>`
	hexagonSolidInnerSVG                     = `<path fill="currentColor" d="M14 4.213L7.5.42L1 4.213v6.574l6.5 3.792l6.5-3.792V4.213Z"/>`
	historyOutlineInnerSVG                   = `<path fill="currentColor" d="m2.5 12.399l.37-.336l-.006-.007l-.007-.007l-.357.35Zm1 1.101v.5H4v-.5h-.5Zm3.5.982l.018-.5l-.053 1l.035-.5ZM7.5 7.5H7a.5.5 0 0 0 .146.354L7.5 7.5Zm6.5 0A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM2.857 12.049A6.477 6.477 0 0 1 1 7.5H0c0 2.043.818 3.897 2.143 5.249l.714-.7Zm-.727.686l1 1.101l.74-.672l-1-1.101l-.74.672ZM7.5 14a6.62 6.62 0 0 1-.465-.016l-.07.997c.177.013.355.019.535.019v-1Zm.018 0l-.5-.017l-.036 1l.5.017l.036-1ZM7 3v4.5h1V3H7Zm.146 4.854l3 3l.708-.708l-3-3l-.708.708ZM0 14h3.5v-1H0v1Zm4-.5V10H3v3.5h1Z"/>`
	historySolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 6.965 7.481l.053-.998l.49.017a6.5 6.5 0 1 0-4.65-1.951l.006.007l.136.15V10h1v4H0v-1h2.37l-.234-.258A7.477 7.477 0 0 1 0 7.5Zm7 0V3h1v4.293l2.854 2.853l-.708.708l-3-3A.5.5 0 0 1 7 7.5Z" clip-rule="evenodd"/>`
	homeAltOutlineInnerSVG                   = `<path fill="currentColor" d="m7.5.5l.354-.354a.5.5 0 0 0-.708 0L7.5.5Zm-6 6l-.354-.354L1 6.293V6.5h.5Zm12 0h.5v-.207l-.146-.147l-.354.354Zm.354-.354l-6-6l-.708.708l6 6l.708-.708Zm-6.708-6l-6 6l.708.708l6-6l-.708-.708ZM14 13.5v-7h-1v7h1Zm-13-7v7h1v-7H1ZM2.5 15h10v-1h-10v1ZM13 13.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1Zm-12 0A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1Z"/>`
	homeAltSolidInnerSVG                     = `<path fill="currentColor" d="M7.854.146a.5.5 0 0 0-.708 0L1 6.293V13.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V6.293L7.854.146Z"/>`
	homeOutlineInnerSVG                      = `<path fill="currentColor" d="m7.5.5l.325-.38a.5.5 0 0 0-.65 0L7.5.5Zm-7 6l-.325-.38L0 6.27v.23h.5Zm5 8v.5a.5.5 0 0 0 .5-.5h-.5Zm4 0H9a.5.5 0 0 0 .5.5v-.5Zm5-8h.5v-.23l-.175-.15l-.325.38ZM1.5 15h4v-1h-4v1Zm13.325-8.88l-7-6l-.65.76l7 6l.65-.76Zm-7.65-6l-7 6l.65.76l7-6l-.65-.76ZM6 14.5v-3H5v3h1Zm3-3v3h1v-3H9Zm.5 3.5h4v-1h-4v1Zm5.5-1.5v-7h-1v7h1Zm-15-7v7h1v-7H0ZM7.5 10A1.5 1.5 0 0 1 9 11.5h1A2.5 2.5 0 0 0 7.5 9v1Zm0-1A2.5 2.5 0 0 0 5 11.5h1A1.5 1.5 0 0 1 7.5 10V9Zm6 6a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1Zm-12-1a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15v-1Z"/>`
	homeSolidInnerSVG                        = `<path fill="currentColor" d="M7.825.12a.5.5 0 0 0-.65 0L0 6.27v7.23A1.5 1.5 0 0 0 1.5 15h4a.5.5 0 0 0 .5-.5v-3a1.5 1.5 0 0 1 3 0v3a.5.5 0 0 0 .5.5h4a1.5 1.5 0 0 0 1.5-1.5V6.27L7.825.12Z"/>`
	hospitalOutlineInnerSVG                  = `<path fill="currentColor" d="m7.5.5l.224-.447a.5.5 0 0 0-.448 0L7.5.5Zm-6 3l-.224-.447A.5.5 0 0 0 1 3.5h.5Zm12 0h.5a.5.5 0 0 0-.276-.447L13.5 3.5Zm-8 7V10H5v.5h.5Zm4 0h.5V10h-.5v.5ZM0 15h15v-1H0v1ZM7.276.053l-6 3l.448.894l6-3l-.448-.894Zm6.448 3l-6-3l-.448.894l6 3l.448-.894ZM7 3v2.5h1V3H7Zm0 2.5V8h1V5.5H7ZM5 6h2.5V5H5v1Zm2.5 0H10V5H7.5v1ZM1 3.5v11h1v-11H1Zm12 0v11h1v-11h-1Zm-7 11v-4H5v4h1ZM5.5 11h4v-1h-4v1Zm3.5-.5v4h1v-4H9Z"/>`
	hospitalSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.724.053a.5.5 0 0 0-.448 0l-5.99 2.995A.5.5 0 0 0 1 3.5V14H0v1h5v-5h5v5h5v-1h-1V3.5a.5.5 0 0 0-.286-.452L7.724.053ZM7 5V3h1v2h2v1H8v2H7V6H5V5h2Z" clip-rule="evenodd"/><path fill="currentColor" d="M9 15v-4H6v4h3Z"/>`
	hourglassOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M3.5 0v4.672a2 2 0 0 0 .586 1.414l.707.707a1 1 0 0 1 0 1.414l-1 1a1 1 0 0 0-.293.707V15m8-15v5.086a1 1 0 0 1-.293.707l-1 1a1 1 0 0 0 0 1.414l1 1a1 1 0 0 1 .293.707V15M1 .5h13m-13 14h13"/>`
	hourglassSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M12 1h2V0H1v1h2v3.672a2.5 2.5 0 0 0 .732 1.767l.707.707a.5.5 0 0 1 0 .708l-1 1A1.5 1.5 0 0 0 3 9.914V14H1v1h13v-1h-2V9.914a1.5 1.5 0 0 0-.44-1.06l-1-1a.5.5 0 0 1 0-.708l1-1a1.5 1.5 0 0 0 .44-1.06V1ZM4.25 5.5h6.543l.06-.06A.5.5 0 0 0 11 5.085V1H4v3.672c0 .296.088.584.25.828Z" clip-rule="evenodd"/>`
	houseOutlineInnerSVG                     = `<path fill="currentColor" d="m.5 5.5l-.29-.407l-.21.15V5.5h.5Zm7-5l.29-.407a.5.5 0 0 0-.58 0L7.5.5Zm7 5h.5v-.257l-.21-.15l-.29.407Zm-12 3V8H2v.5h.5Zm4 0H7V8h-.5v.5ZM1 15V5.5H0V15h1ZM.79 5.907l7-5l-.58-.814l-7 5l.58.814Zm6.42-5l7 5l.58-.814l-7-5l-.58.814ZM14 5.5V15h1V5.5h-1ZM3 15V8.5H2V15h1Zm-.5-6h4V8h-4v1ZM6 8.5V15h1V8.5H6ZM5 12h1.5v-1H5v1Zm6-4v4h1V8h-1Zm2-6v2.5h1V2h-1Z"/>`
	houseSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M7.79.093a.5.5 0 0 0-.58 0l-7 5A.5.5 0 0 0 0 5.5v9a.5.5 0 0 0 .5.5H2V8h5v7h7.5a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.21-.407L14 4.528V2h-1v1.814L7.79.094ZM11 12V8h1v4h-1Z" clip-rule="evenodd"/><path fill="currentColor" d="M6 15v-3H5v-1h1V9H3v6h3Z"/>`
	htmlFiveOutlineInnerSVG                  = `<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.498.542L.5.5Zm14 0l.498.042A.5.5 0 0 0 14.5 0v.5Zm-1 12l.158.474a.5.5 0 0 0 .34-.433L13.5 12.5Zm-6 2l-.158.474a.499.499 0 0 0 .316 0L7.5 14.5Zm-6-2l-.498.041a.5.5 0 0 0 .34.433L1.5 12.5Zm3-9V3H4v.5h.5Zm0 3H4V7h.5v-.5Zm6 0h.5V6h-.5v.5Zm0 3l.158.474L11 9.86V9.5h-.5Zm-3 1l-.158.474l.158.053l.158-.053L7.5 10.5Zm-3-1H4v.36l.342.114L4.5 9.5ZM.5 1h14V0H.5v1ZM14.002.458l-1 12l.996.083l1-12l-.996-.083Zm-.66 11.568l-6 2l.316.948l6-2l-.316-.948Zm-5.684 2l-6-2l-.316.948l6 2l.316-.948Zm-5.66-1.567l-1-12l-.996.083l1 12l.996-.083ZM11 3H4.5v1H11V3Zm-7 .5v3h1v-3H4ZM4.5 7h6V6h-6v1Zm5.5-.5v3h1v-3h-1Zm.342 2.526l-3 1l.316.948l3-1l-.316-.948Zm-2.684 1l-3-1l-.316.948l3 1l.316-.948ZM5 9.5V8H4v1.5h1Z"/>`
	htmlFiveSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M.132.161A.5.5 0 0 1 .5 0h14a.5.5 0 0 1 .498.542l-1 12a.5.5 0 0 1-.34.432l-6 2a.499.499 0 0 1-.316 0l-6-2a.5.5 0 0 1-.34-.433l-1-12a.5.5 0 0 1 .13-.38ZM11 3H4v4h6v2.14l-2.5.833L5 9.14V8H4v1.86l3.5 1.167L11 9.86V6H5V4h6V3Z" clip-rule="evenodd"/>`
	idOutlineInnerSVG                        = `<path fill="currentColor" d="M2 12.5v.5h1v-.5H2Zm5 0v.5h1v-.5H7Zm-4 0V12H2v.5h1Zm4-.5v.5h1V12H7Zm-2-2a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3v1Zm-2 2a2 2 0 0 1 2-2V9a3 3 0 0 0-3 3h1Zm2-8a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1V4Zm2 2a2 2 0 0 0-2-2v1a1 1 0 0 1 1 1h1ZM5 8a2 2 0 0 0 2-2H6a1 1 0 0 1-1 1v1Zm0-1a1 1 0 0 1-1-1H3a2 2 0 0 0 2 2V7ZM1.5 3h12V2h-12v1Zm12.5.5v8h1v-8h-1Zm-.5 8.5h-12v1h12v-1ZM1 11.5v-8H0v8h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 13v-1Zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 2v1Zm-12-1A1.5 1.5 0 0 0 0 3.5h1a.5.5 0 0 1 .5-.5V2ZM9 6h3V5H9v1Zm0 3h3V8H9v1Z"/>`
	idSolidInnerSVG                          = `<path fill="currentColor" fill-rule="evenodd" d="M0 3.5A1.5 1.5 0 0 1 1.5 2h12A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5v-8ZM3 6a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm9 0H9V5h3v1Zm0 3H9V8h3v1ZM5 9a2.927 2.927 0 0 0-2.618 1.618l-.33.658A.5.5 0 0 0 2.5 12h5a.5.5 0 0 0 .447-.724l-.329-.658A2.927 2.927 0 0 0 5 9Z" clip-rule="evenodd"/>`
	imacOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M2 14.5h11m-5.5-4v4M0 7.5h15M.5 1.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`
	imacSolidInnerSVG                        = `<path fill="currentColor" d="M13.5 0A1.5 1.5 0 0 1 15 1.5V7H0V1.5A1.5 1.5 0 0 1 1.5 0h12ZM0 8v1.5A1.5 1.5 0 0 0 1.5 11H7v3H2v1h11v-1H8v-3h5.5A1.5 1.5 0 0 0 15 9.5V8H0Z"/>`
	imageAltOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M10 3.5h1m3.5 5.993l-3-2.998l-3 2.998l-4-4.996L.5 9.5m1-9h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`
	imageAltSolidInnerSVG                    = `<path fill="currentColor" d="M11 4h-1V3h1v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12.01A1.5 1.5 0 0 1 13.5 15h-12A1.5 1.5 0 0 1 0 13.5v-12Zm14 6.787l-2.5-2.498l-2.959 2.956L4.5 3.696L1 8.074V1.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5v6.787Z" clip-rule="evenodd"/>`
	imageDocumentOutlineInnerSVG             = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm1 7.995l.354-.353l-.353-.354l-.354.354l.353.353Zm-3 2.998l-.39.313l.349.436l.394-.395l-.353-.354ZM4.5 6.5l.39-.313l-.403-.503L4.1 6.2l.4.3Zm8 7.5h-10v1h10v-1ZM2 13.5v-12H1v12h1Zm11-10v10h1v-10h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM10 6h1V5h-1v1Zm3.854 4.147l-2-2.005l-.708.707l2 2.004l.708-.706Zm-2.707-2.005l-3 2.998l.706.707l3-2.998l-.706-.707ZM8.89 11.18l-4-4.994l-.78.626l4 4.993l.78-.625ZM4.1 6.2l-3 4l.8.6l3-4l-.8-.6Z"/>`
	imageDocumentSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293v6.999l-2.5-2.504l-2.959 2.957L4.5 5.7L1 10.075V1.5ZM11 6h-1V5h1v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M1 11.675V13.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-1.793l-2.5-2.504l-3.041 3.039L4.5 7.3L1 11.675Z"/>`
	imageOutlineInnerSVG                     = `<path fill="currentColor" d="m4.5 3.5l.354-.354a.5.5 0 0 0-.708 0L4.5 3.5ZM1.5 1h12V0h-12v1Zm12.5.5v12h1v-12h-1ZM13.5 14h-12v1h12v-1ZM1 13.5v-12H0v12h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15v-1Zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0v1Zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5V0Zm-1 11h14v-1H.5v1Zm.354-3.146l4-4l-.708-.708l-4 4l.708.708Zm3.292-4l7 7l.708-.708l-7-7l-.708.708ZM10.5 5a.5.5 0 0 1-.5-.5H9A1.5 1.5 0 0 0 10.5 6V5Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 12 4.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 3v1Zm0-1A1.5 1.5 0 0 0 9 4.5h1a.5.5 0 0 1 .5-.5V3Z"/>`
	imageSolidInnerSVG                       = `<path fill="currentColor" d="M10.5 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.505 1.505 0 0 1-.178.71A1.5 1.5 0 0 1 13.5 15h-12A1.5 1.5 0 0 1 0 13.5v-12Zm4.85 1.642l-.35-.35l-3.5 3.5V1.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5V10h-2.293L4.854 3.146l-.005-.004Z" clip-rule="evenodd"/>`
	inEarHeadphonesOutlineInnerSVG           = `<path fill="currentColor" d="M14.5 12.5v.5h.5v-.5h-.5Zm-2 0H12v.5h.5v-.5Zm-12 0H0v.5h.5v-.5Zm2 0v.5H3v-.5h-.5ZM10 5V1.654H9V5h1Zm4-1.5v9h1v-9h-1Zm.5 8.5h-2v1h2v-1Zm-2-6H11v1h1.5V6Zm-1.846-5h.846V0h-.846v1ZM11 6a1 1 0 0 1-1-1H9a2 2 0 0 0 2 2V6Zm4-2.5A3.5 3.5 0 0 0 11.5 0v1A2.5 2.5 0 0 1 14 3.5h1Zm-5-1.846c0-.361.293-.654.654-.654V0C9.74 0 9 .74 9 1.654h1ZM13 12.5V15h1v-2.5h-1ZM6 5V1.654H5V5h1ZM0 3.5v9h1v-9H0ZM.5 13h2v-1h-2v1Zm2-6H4V6H2.5v1Zm1.846-7H3.5v1h.846V0ZM4 7a2 2 0 0 0 2-2H5a1 1 0 0 1-1 1v1ZM1 3.5A2.5 2.5 0 0 1 3.5 1V0A3.5 3.5 0 0 0 0 3.5h1Zm5-1.846C6 .74 5.26 0 4.346 0v1c.361 0 .654.293.654.654h1ZM1 12.5V15h1v-2.5H1ZM12 4v2.5h1V4h-1ZM2 4v2.5h1V4H2Zm11 8.5v-6h-1v6h1Zm-10 0v-6H2v6h1Z"/>`
	inEarHeadphonesSolidInnerSVG             = `<path fill="currentColor" d="M6 1.654C6 .74 5.26 0 4.346 0H3.5A3.5 3.5 0 0 0 0 3.5V13h1v2h1v-2h1V7h1a2 2 0 0 0 2-2V1.654ZM10.654 0C9.74 0 9 .74 9 1.654V5a2 2 0 0 0 2 2h1v6h1v2h1v-2h1V3.5A3.5 3.5 0 0 0 11.5 0h-.846Z"/>`
	inboxOutlineInnerSVG                     = `<path fill="currentColor" d="m7.713 11.493l-.035-.5l.035.5ZM1.5 1h12V0h-12v1Zm12.5.5v12h1v-12h-1ZM13.5 14h-12v1h12v-1ZM1 13.5v-12H0v12h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15v-1Zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0v1Zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5V0Zm6 12c.083 0 .166-.003.248-.009l-.07-.997A2.546 2.546 0 0 1 7.5 11v1Zm-.823-.098c.264.064.54.098.823.098v-1c-.203 0-.4-.024-.589-.07l-.234.973Zm.234-.972c-.969-.233-1.9-.895-2.97-1.586C2.924 8.687 1.771 8 .5 8v1c.938 0 1.858.512 2.899 1.184c.987.638 2.099 1.434 3.278 1.719l.234-.973Zm.837 1.061c1.386-.097 2.7-.927 3.865-1.632C12.843 9.616 13.922 9 15 9V8c-1.407 0-2.732.794-3.905 1.503c-1.237.749-2.324 1.414-3.417 1.49l.07.998Z"/>`
	inboxSolidInnerSVG                       = `<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5V8h-4.5a.5.5 0 0 0-.5.5a2.5 2.5 0 0 1-5 0a.5.5 0 0 0-.5-.5H0V1.5Z"/><path fill="currentColor" d="M0 9v4.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5V9h-4.035a3.5 3.5 0 0 1-6.93 0H0Z"/>`
	indentDecreaseOutlineInnerSVG            = `<path fill="currentColor" d="m.5 7.5l-.354-.354a.5.5 0 0 0 0 .708L.5 7.5ZM3 4h12V3H3v1Zm4 4h8V7H7v1Zm-4 4h12v-1H3v1Zm-.146-2.854l-2-2l-.708.708l2 2l.708-.708Zm-2-1.292l2-2l-.708-.708l-2 2l.708.708Z"/>`
	indentDecreaseSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M15 4H3V3h12v1ZM1.207 7.5l1.647-1.646l-.708-.708l-2 2a.5.5 0 0 0 0 .708l2 2l.708-.708L1.207 7.5ZM15 8H7V7h8v1Zm0 4H3v-1h12v1Z" clip-rule="evenodd"/>`
	indentIncreaseOutlineInnerSVG            = `<path fill="currentColor" d="m2.5 7.5l.354.354a.5.5 0 0 0 0-.708L2.5 7.5ZM3 4h12V3H3v1Zm4 4h8V7H7v1Zm-4 4h12v-1H3v1ZM.854 9.854l2-2l-.708-.708l-2 2l.708.708Zm2-2.708l-2-2l-.708.708l2 2l.708-.708Z"/>`
	indentIncreaseSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M15 4H3V3h12v1ZM.854 5.146l2 2a.5.5 0 0 1 0 .708l-2 2l-.708-.708L1.793 7.5L.146 5.854l.708-.708ZM15 8H7V7h8v1Zm0 4H3v-1h12v1Z" clip-rule="evenodd"/>`
	infoCircleOutlineInnerSVG                = `<path fill="currentColor" d="M7 4.5V5h1v-.5H7Zm1-.01v-.5H7v.5h1ZM8 11V7H7v4h1Zm0-6.5v-.01H7v.01h1ZM6 8h1.5V7H6v1Zm0 3h3v-1H6v1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1Zm0 1A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1Z"/>`
	infoCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0ZM7 5V3.99h1V5H7Zm1 2v3h1v1H6v-1h1V8H6V7h2Z" clip-rule="evenodd"/>`
	infoOutlineInnerSVG                      = `<path fill="currentColor" d="M7 1.5V2h1v-.5H7Zm1-.01v-.5H7v.5h1ZM8 13.5V4H7v9.5h1Zm0-12v-.01H7v.01h1ZM4 5h3.5V4H4v1Zm-2 9h11v-1H2v1Z"/>`
	infoSmallOutlineInnerSVG                 = `<path fill="currentColor" d="M7 4.5V5h1v-.5H7Zm1-.01v-.5H7v.5h1ZM8 11V7H7v4h1Zm0-6.5v-.01H7v.01h1ZM6 8h1.5V7H6v1Zm0 3h3v-1H6v1Z"/>`
	infoSmallSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M8 3.99V5H7V3.99h1ZM6 11v-1h1V8H6V7h2v3h1v1H6Z" clip-rule="evenodd"/>`
	infoSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M8 .99V2H7V.99h1ZM7 13H2v1h11v-1H8V4H4v1h3v8Z" clip-rule="evenodd"/>`
	instagramOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M11 3.5h1M4.5.5h6a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4h-6a4 4 0 0 1-4-4v-6a4 4 0 0 1 4-4Zm3 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`
	instagramSolidInnerSVG                   = `<path fill="currentColor" d="M7.5 5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z"/><path fill="currentColor" fill-rule="evenodd" d="M4.5 0A4.5 4.5 0 0 0 0 4.5v6A4.5 4.5 0 0 0 4.5 15h6a4.5 4.5 0 0 0 4.5-4.5v-6A4.5 4.5 0 0 0 10.5 0h-6ZM4 7.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0ZM11 4h1V3h-1v1Z" clip-rule="evenodd"/>`
	invoiceOutlineInnerSVG                   = `<path fill="currentColor" d="M4.5 6.995H4v1h.5v-1Zm6 1h.5v-1h-.5v1ZM4.5 9.5H4v.5h.5v-.5Zm6 0v.5h.5v-.5h-.5Zm-6-4V5H4v.5h.5Zm6 0h.5V5h-.5v.5Zm3-2h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-6 7.495h6v-1h-6v1ZM4.5 10h6V9h-6v1Zm0-4h6V5h-6v1Zm8 8h-10v1h10v-1ZM2 13.5v-12H1v12h1Zm11-10v10h1v-10h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1Zm2 4v4h1v-4H4Zm3 0v4h1v-4H7Zm3 0v4h1v-4h-1ZM4 4h3V3H4v1Zm4 8h3v-1H8v1Z"/>`
	invoiceSolidInnerSVG                     = `<path fill="currentColor" d="M10 7.995V9H8V7.995h2ZM10 6v.995H8V6h2ZM7 6H5v.995h2V6Zm0 1.995H5V9h2V7.995Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM4 4h3V3H4v1Zm7 1H4v5h7V5Zm0 7H8v-1h3v1Z" clip-rule="evenodd"/>`
	italicOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M4 1.5h9m-11 12h9m-2.5-12l-2 12"/>`
	italicSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7.91 2H4V1h9v1H8.924L7.09 13H11v1H2v-1h4.076L7.91 2Z" clip-rule="evenodd"/>`
	javascriptOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M12.5 8v-.167c0-.736-.597-1.333-1.333-1.333H10a1.5 1.5 0 1 0 0 3h1a1.5 1.5 0 0 1 0 3h-1A1.5 1.5 0 0 1 8.5 11m-2-5v5a1.5 1.5 0 0 1-3 0M.5.5h14v14H.5V.5Z"/>`
	javascriptSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M15 0H0v15h15V0ZM8 8a2 2 0 0 1 2-2h1.167C12.179 6 13 6.82 13 7.833V8h-1v-.167A.833.833 0 0 0 11.167 7H10a1 1 0 0 0 0 2h1a2 2 0 1 1 0 4h-1a2 2 0 0 1-2-2h1a1 1 0 0 0 1 1h1a1 1 0 1 0 0-2h-1a2 2 0 0 1-2-2ZM6 6v5a1 1 0 1 1-2 0H3a2 2 0 1 0 4 0V6H6Z" clip-rule="evenodd"/>`
	joystickOutlineInnerSVG                  = `<path fill="currentColor" d="M1.5 10.5V10a.5.5 0 0 0-.5.5h.5Zm12 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 4v.5a.5.5 0 0 0 .5-.5h-.5Zm-12 0H1a.5.5 0 0 0 .5.5v-.5Zm0-3.5h12v-1h-12v1Zm11.5-.5v4h1v-4h-1Zm.5 3.5h-12v1h12v-1ZM2 14.5v-4H1v4h1Zm6-4v-4H7v4h1ZM7.5 0A3.5 3.5 0 0 0 4 3.5h1A2.5 2.5 0 0 1 7.5 1V0ZM11 3.5A3.5 3.5 0 0 0 7.5 0v1A2.5 2.5 0 0 1 10 3.5h1ZM7.5 7A3.5 3.5 0 0 0 11 3.5h-1A2.5 2.5 0 0 1 7.5 6v1Zm0-1A2.5 2.5 0 0 1 5 3.5H4A3.5 3.5 0 0 0 7.5 7V6ZM3 9h2V8H3v1Z"/>`
	joystickSolidInnerSVG                    = `<path fill="currentColor" d="M4 3.5a3.5 3.5 0 1 1 4 3.465V10h5.5a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-12a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5H7V6.965A3.5 3.5 0 0 1 4 3.5Z"/><path fill="currentColor" d="M3 8v1h2V8H3Z"/>`
	jpgOutlineInnerSVG                       = `<path fill="currentColor" d="M6.5 6.5V6H6v.5h.5Zm4 4H10v.5h.5v-.5Zm2 0v.5h.5v-.5h-.5Zm1-7h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-6 6H5V6h-.5v.5Zm0 4v.5H5v-.5h-.5Zm-2 0H2v.5h.5v-.5Zm4-3.5h1V6h-1v1Zm.5 4V8.5H6V11h1Zm0-2.5v-2H6v2h1Zm.5-.5h-1v1h1V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5H8ZM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6v1ZM10 6v4.5h1V6h-1Zm.5 5h2v-1h-2v1Zm2.5-.5v-2h-1v2h1ZM10.5 7H13V6h-2.5v1ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM2 7h2.5V6H2v1Zm2-.5v4h1v-4H4Zm.5 3.5h-2v1h2v-1Zm-1.5.5V9H2v1.5h1Z"/>`
	jpgSolidInnerSVG                         = `<path fill="currentColor" d="M7 8h.5a.5.5 0 0 0 0-1H7v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM4 7H2V6h3v5H2V9h1v1h1V7Zm2-1h1.5a1.5 1.5 0 1 1 0 3H7v2H6V6Zm4 0h3v1h-2v3h1V8.5h1V11h-3V6Z" clip-rule="evenodd"/>`
	kanbanOutlineInnerSVG                    = `<path fill="currentColor" d="M.5 3.5V3a.5.5 0 0 0-.5.5h.5Zm6 0H7a.5.5 0 0 0-.5-.5v.5Zm0 11v.5a.5.5 0 0 0 .5-.5h-.5Zm-6 0H0a.5.5 0 0 0 .5.5v-.5Zm8-11V3a.5.5 0 0 0-.5.5h.5Zm6 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 6v.5a.5.5 0 0 0 .5-.5h-.5Zm-6 0H8a.5.5 0 0 0 .5.5v-.5ZM0 1h7V0H0v1Zm8 0h7V0H8v1ZM.5 4h6V3h-6v1ZM6 3.5v11h1v-11H6ZM6.5 14h-6v1h6v-1Zm-5.5.5v-11H0v11h1ZM8.5 4h6V3h-6v1Zm5.5-.5v6h1v-6h-1Zm.5 5.5h-6v1h6V9ZM9 9.5v-6H8v6h1Z"/>`
	kanbanSolidInnerSVG                      = `<path fill="currentColor" d="M0 1h7V0H0v1Zm8 0h7V0H8v1ZM.5 3a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-11a.5.5 0 0 0-.5-.5h-6Zm8 0a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-6Z"/>`
	keyOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="m.5 14.5l8-8m-6 6l2 2m0-4l2 2m4.5-5a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Z"/>`
	keySolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M7 4a4 4 0 1 1 1.547 3.16l-3.34 3.34l1.647 1.646l-.708.708L4.5 11.207L3.207 12.5l1.647 1.646l-.708.708L2.5 13.207L.854 14.854l-.708-.708L7.84 6.453A3.983 3.983 0 0 1 7 4Zm4-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z" clip-rule="evenodd"/>`
	keyboardOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M11 11.5H4m7-3h-1m-2 0H7m-2 0H4m7-2h-1m-2 0H7m-2 0H4m3.5-2V0m6 4.5h-12a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1Z"/>`
	keyboardSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M8 0v4h5.5A1.5 1.5 0 0 1 15 5.5v7a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-7A1.5 1.5 0 0 1 1.5 4H7V0h1Zm2 6h1v1h-1V6Zm1 2h-1v1h1V8Zm0 3H4v1h7v-1ZM7 8h1v1H7V8ZM5 8H4v1h1V8Zm3-2H7v1h1V6ZM4 6h1v1H4V6Z" clip-rule="evenodd"/>`
	lanCableOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M6.5 11.5V15m2-3.5V15M6 9.5h3M6.5.5v2h2v-2m-4 0h6v4h1v3l-2 4h-4l-2-4v-3h1v-4Zm2 4v2a1 1 0 0 0 2 0v-2h-2Z"/>`
	lanCableSolidInnerSVG                    = `<path fill="currentColor" d="M7 6.5V5h1v1.5a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M9 0h2v4h1v3.618L9.809 12H9v3H8v-3H7v3H6v-3h-.809L3 7.618V4h1V0h2v3h3V0Zm0 4H6v2.5a1.5 1.5 0 1 0 3 0V4ZM6 9v1h3V9H6Z" clip-rule="evenodd"/><path fill="currentColor" d="M8 0H7v2h1V0Z"/>`
	laptopOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M0 13.5h15M1.5 2.5v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1Z"/>`
	laptopSolidInnerSVG                      = `<path fill="currentColor" d="M2.5 1A1.5 1.5 0 0 0 1 2.5v8A1.5 1.5 0 0 0 2.5 12h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 12.5 1h-10ZM0 14h15v-1H0v1Z"/>`
	laravelOutlineInnerSVG                   = `<path fill="currentColor" d="m2 3.571l-.158-.474a.5.5 0 0 0-.327.596L2 3.57Zm2.357-.785l.447-.224a.5.5 0 0 0-.605-.25l.158.474Zm4.714 9.428l-.447.224a.5.5 0 0 0 .705.205l-.258-.429ZM13 9.857l.257.429a.5.5 0 0 0 .165-.697L13 9.857Zm-2.75-4.321l-.121-.485a.5.5 0 0 0-.3.753l.421-.268Zm1.571-.393l.4-.3a.5.5 0 0 0-.52-.185l.12.485ZM13 6.714l.158.475a.5.5 0 0 0 .242-.775l-.4.3ZM3.571 9.857l-.485.121a.5.5 0 0 0 .644.354l-.159-.475ZM2.158 4.046l2.357-.786l-.316-.949l-2.357.786l.316.949ZM3.91 3.009l2.918 5.837l.895-.447l-2.919-5.837l-.894.447Zm2.918 5.837l1.796 3.592l.895-.447l-1.796-3.592l-.895.447Zm2.5 3.797l3.93-2.357l-.515-.858l-3.929 2.358l.515.857Zm4.094-3.054l-1.65-2.593l-.844.537l1.65 2.593l.844-.537Zm-1.65-2.593l-1.1-1.729l-.844.537l1.1 1.729l.844-.537Zm-1.4-.975l1.57-.393l-.242-.97l-1.571.393l.242.97Zm1.05-.578L12.6 7.014l.8-.6l-1.179-1.571l-.8.6Zm1.42.797l-1.65.55l.316.949l1.65-.55l-.316-.949Zm-1.65.55L7.117 8.148l.317.949l4.074-1.358l-.316-.949ZM7.117 8.148L3.413 9.383l.317.949l3.704-1.235l-.317-.949Zm-3.06 1.588L2.484 3.45l-.97.243l1.571 6.285l.97-.242ZM2.5 1h10V0h-10v1ZM14 2.5v10h1v-10h-1ZM12.5 14h-10v1h10v-1ZM1 12.5v-10H0v10h1ZM2.5 14A1.5 1.5 0 0 1 1 12.5H0A2.5 2.5 0 0 0 2.5 15v-1ZM14 12.5a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5h-1ZM12.5 1A1.5 1.5 0 0 1 14 2.5h1A2.5 2.5 0 0 0 12.5 0v1Zm-10-1A2.5 2.5 0 0 0 0 2.5h1A1.5 1.5 0 0 1 2.5 1V0Z"/>`
	laravelSolidInnerSVG                     = `<path fill="currentColor" d="M4.104 3.397L6.57 8.33l-2.645.882L2.597 3.9l1.507-.502Zm5.173 8.111L7.981 8.915l3.157-1.053l1.165 1.83l-3.026 1.816Zm2.907-5.048l-.622.207l-.518-.814l.577-.145l.563.751Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 2.5A2.5 2.5 0 0 1 2.5 0h10A2.5 2.5 0 0 1 15 2.5v10a2.5 2.5 0 0 1-2.5 2.5h-10A2.5 2.5 0 0 1 0 12.5v-10Zm4.804.062a.5.5 0 0 0-.605-.25l-2.357.785a.5.5 0 0 0-.327.596l1.571 6.285a.5.5 0 0 0 .644.354l3.292-1.098l1.602 3.204a.5.5 0 0 0 .705.205l3.928-2.357a.5.5 0 0 0 .165-.697l-1.306-2.053l1.042-.347a.5.5 0 0 0 .242-.775l-1.178-1.571a.5.5 0 0 0-.522-.185l-1.571.393a.5.5 0 0 0-.3.753l.755 1.188L7.53 8.011L4.804 2.562Z" clip-rule="evenodd"/>`
	layersDifferenceOutlineInnerSVG          = `<path fill="none" stroke="currentColor" d="M4.5 4.5v-3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-3m-6-6h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-3m-6-6V7m0-2.5H7m3.5 6H8m2.5 0V8M8 4.5h2.5V7m-6 1v2.5H7"/>`
	layersDifferenceSolidInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M4 4V1.5A1.5 1.5 0 0 1 5.5 0h8A1.5 1.5 0 0 1 15 1.5v8a1.5 1.5 0 0 1-1.5 1.5H11v2.5A1.5 1.5 0 0 1 9.5 15h-8A1.5 1.5 0 0 1 0 13.5v-8A1.5 1.5 0 0 1 1.5 4H4Zm6 1v5H5V5h5Z" clip-rule="evenodd"/>`
	layersIntersectOutlineInnerSVG           = `<path fill="none" stroke="currentColor" d="M4.5 3V1.5a1 1 0 0 1 1-1H7m5 0h1.5a1 1 0 0 1 1 1V3M8 .5h3m1 10h1.5a1 1 0 0 0 1-1V8m0-4v3M3 4.5H1.5a1 1 0 0 0-1 1V7m0 5v1.5a1 1 0 0 0 1 1H3M.5 8v3M8 14.5h1.5a1 1 0 0 0 1-1V12M4 14.5h3m-2.5-10v6h6v-6h-6Z"/>`
	layersIntersectSolidInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M4 1.5V4H1.5A1.5 1.5 0 0 0 0 5.5v8A1.5 1.5 0 0 0 1.5 15h8a1.5 1.5 0 0 0 1.5-1.5V11h2.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-8A1.5 1.5 0 0 0 4 1.5ZM5.5 1a.5.5 0 0 0-.5.5V4h4.5A1.5 1.5 0 0 1 11 5.5V10h2.5a.5.5 0 0 0 .5-.5v-8a.5.5 0 0 0-.5-.5h-8Zm0 10A1.5 1.5 0 0 1 4 9.5V5H1.5a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V11H5.5Z" clip-rule="evenodd"/>`
	layersOutlineInnerSVG                    = `<path fill="currentColor" d="m7.5 1.5l.197-.46a.5.5 0 0 0-.394 0l.197.46Zm-7 3l-.197-.46a.5.5 0 0 0 0 .92L.5 4.5Zm7 3l-.197.46a.5.5 0 0 0 .394 0L7.5 7.5Zm7-3l.197.46a.5.5 0 0 0 0-.92l-.197.46Zm-7 6l-.197.46l.197.084l.197-.084l-.197-.46Zm0 3l-.197.46l.197.084l.197-.084l-.197-.46ZM7.303 1.04l-7 3l.394.92l7-3l-.394-.92Zm-7 3.92l7 3l.394-.92l-7-3l-.394.92Zm7.394 3l7-3l-.394-.92l-7 3l.394.92Zm7-3.92l-7-3l-.394.92l7 3l.394-.92ZM.303 7.96l7 3l.394-.92l-7-3l-.394.92Zm7.394 3l7-3l-.394-.92l-7 3l.394.92Zm-7.394 0l7 3l.394-.92l-7-3l-.394.92Zm7.394 3l7-3l-.394-.92l-7 3l.394.92Z"/>`
	layersSolidInnerSVG                      = `<path fill="currentColor" d="M7.697 1.04a.5.5 0 0 0-.394 0l-7 3a.5.5 0 0 0 0 .92l7 3a.5.5 0 0 0 .394 0l7-3a.5.5 0 0 0 0-.92l-7-3Z"/><path fill="currentColor" d="M7.5 9.956L.697 7.04l-.394.92L7.5 11.044l7.197-3.084l-.394-.92L7.5 9.956Z"/><path fill="currentColor" d="m.697 10.04l-.394.92L7.5 14.044l7.197-3.084l-.394-.92L7.5 12.956L.697 10.04Z"/>`
	layersSubtractOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M10.5 10.5v3a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h3m0-3v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1Z"/>`
	layersSubtractSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M4 4H1.5A1.5 1.5 0 0 0 0 5.5v8A1.5 1.5 0 0 0 1.5 15h8a1.5 1.5 0 0 0 1.5-1.5V11h2.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-8A1.5 1.5 0 0 0 4 1.5V4Zm1-2.5a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-8Z" clip-rule="evenodd"/>`
	layersUnionOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M4.5 1.5v3h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-3h3a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1Z"/>`
	layersUnionSolidInnerSVG                 = `<path fill="currentColor" d="M5.5 0A1.5 1.5 0 0 0 4 1.5V4H1.5A1.5 1.5 0 0 0 0 5.5v8A1.5 1.5 0 0 0 1.5 15h8a1.5 1.5 0 0 0 1.5-1.5V11h2.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-8Z"/>`
	leftCircleOutlineInnerSVG                = `<path fill="currentColor" d="m8.854 4.854l.353-.354l-.707-.707l-.354.353l.708.708ZM5.5 7.5l-.354-.354l-.353.354l.353.354L5.5 7.5Zm2.646 3.354l.354.353l.707-.707l-.353-.354l-.708.708Zm0-6.708l-3 3l.708.708l3-3l-.708-.708Zm-3 3.708l3 3l.708-.708l-3-3l-.708.708ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	leftCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm9 4.207L4.793 7.5L9 3.293v8.414Z" clip-rule="evenodd"/>`
	leftOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M10 14L3 7.5L10 1"/>`
	leftSmallOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m8.5 4.5l-3 3l3 3"/>`
	leftSmallSolidInnerSVG                   = `<path fill="currentColor" d="M4.793 7.5L9 11.707V3.293L4.793 7.5Z"/>`
	leftSolidInnerSVG                        = `<path fill="currentColor" d="M3 7.5L11 0v15L3 7.5Z"/>`
	legoOutlineInnerSVG                      = `<path fill="currentColor" d="M3.5 3h8V2h-8v1ZM13 4.5v8h1v-8h-1ZM11.5 14h-8v1h8v-1ZM2 12.5v-8H1v8h1ZM3.5 14A1.5 1.5 0 0 1 2 12.5H1A2.5 2.5 0 0 0 3.5 15v-1Zm9.5-1.5a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5h-1ZM11.5 3A1.5 1.5 0 0 1 13 4.5h1A2.5 2.5 0 0 0 11.5 2v1Zm-8-1A2.5 2.5 0 0 0 1 4.5h1A1.5 1.5 0 0 1 3.5 3V2ZM5 8h1V7H5v1Zm4 0h1V7H9v1Zm1.1 1.7a3.25 3.25 0 0 1-5.2 0l-.8.6c1.7 2.267 5.1 2.267 6.8 0l-.8-.6ZM5 2.5v-1H4v1h1ZM5.5 1h4V0h-4v1Zm4.5.5v1h1v-1h-1ZM9.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 0v1ZM5 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 4 1.5h1Z"/>`
	legoSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M4 1.5A1.5 1.5 0 0 1 5.5 0h4A1.5 1.5 0 0 1 11 1.5V2h.5A2.5 2.5 0 0 1 14 4.5v8a2.5 2.5 0 0 1-2.5 2.5h-8A2.5 2.5 0 0 1 1 12.5v-8A2.5 2.5 0 0 1 3.5 2H4v-.5ZM5 7v1h1V7H5Zm4 0v1h1V7H9ZM4.9 9.7a3.25 3.25 0 0 0 5.2 0l.8.6c-1.7 2.267-5.1 2.267-6.8 0l.8-.6Z" clip-rule="evenodd"/>`
	lifebuoyOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M10.329 10.328a4 4 0 0 1-5.657 0m5.657 0a4 4 0 0 0 0-5.656m0 5.656l2.12 2.122m-7.777-2.122a4 4 0 0 1 0-5.656m0 5.656L2.55 12.45m7.779-7.778a4 4 0 0 0-5.657 0m5.657 0l2.12-2.122M4.673 4.672L2.55 2.55m9.9 9.9a7 7 0 0 1-9.9 0m9.9 0a7 7 0 0 0 0-9.9m-9.9 9.9a7 7 0 0 1 0-9.9m9.9 0a7 7 0 0 0-9.9 0"/>`
	lifebuoySolidInnerSVG                    = `<path fill="currentColor" d="M5.404 4.697L2.562 1.855a7.501 7.501 0 0 1 9.876 0L9.596 4.697a3.502 3.502 0 0 0-4.192 0Zm4.899.707a3.502 3.502 0 0 1 0 4.192l2.842 2.842a7.501 7.501 0 0 0 0-9.876l-2.842 2.842Zm-.707 4.899a3.502 3.502 0 0 1-4.192 0l-2.842 2.842a7.501 7.501 0 0 0 9.876 0l-2.842-2.842ZM4.697 5.404a3.502 3.502 0 0 0 0 4.192l-2.842 2.842a7.501 7.501 0 0 1 0-9.876l2.842 2.842Z"/>`
	lightningCableOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M4.5 5.5h6m-6 0a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1m-6 0v-4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4M9 2.5H6M5.5 13v2m4-2v2"/>`
	lightningCableSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 0A1.5 1.5 0 0 0 4 1.5V5h7V1.5A1.5 1.5 0 0 0 9.5 0h-4ZM6 2h3v1H6V2Z" clip-rule="evenodd"/><path fill="currentColor" d="M3 6h9v5.5a1.5 1.5 0 0 1-1.5 1.5H10v2H9v-2H6v2H5v-2h-.5A1.5 1.5 0 0 1 3 11.5V6Z"/>`
	lineOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="m2 2l11 11M1.5 2.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm12 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`
	lineSolidInnerSVG                        = `<path fill="currentColor" d="M1.5 0a1.5 1.5 0 1 0 .647 2.854l10 10a1.5 1.5 0 1 0 .707-.707l-10-10A1.5 1.5 0 0 0 1.5 0Z"/>`
	linkOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M4.5 6.5L1.328 9.672a2.828 2.828 0 1 0 4 4L8.5 10.5m2-2l3.172-3.172a2.829 2.829 0 0 0-4-4L6.5 4.5m-2 6l6-6"/>`
	linkRemoveOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M4.5 6.5L1.328 9.672a2.828 2.828 0 1 0 4 4L8.5 10.5m2-2l3.172-3.172a2.829 2.829 0 0 0-4-4L6.5 4.5m-2 6l6-6M3 4.5H0M4.5 0v3m6 9v3m1.5-4.5h3"/>`
	linkRemoveSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M4 3V0h1v3H4ZM9.318.975a3.328 3.328 0 1 1 4.707 4.707l-3.171 3.172l-.708-.708l3.172-3.171a2.328 2.328 0 1 0-3.293-3.293L6.854 4.854l-.708-.708L9.318.975ZM0 4h3v1H0V4Zm10.854.854l-6 6l-.708-.708l6-6l.708.708Zm-6 2l-3.172 3.171a2.329 2.329 0 0 0 3.293 3.293l3.171-3.172l.708.708l-3.172 3.171A3.328 3.328 0 1 1 .975 9.318l3.171-3.172l.708.708ZM15 11h-3v-1h3v1Zm-5 4v-3h1v3h-1Z" clip-rule="evenodd"/>`
	linkSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M9.318.975a3.328 3.328 0 1 1 4.707 4.707l-3.171 3.172l-.708-.708l3.172-3.171a2.328 2.328 0 1 0-3.293-3.293L6.854 4.854l-.708-.708L9.318.975Zm1.536 3.879l-6 6l-.708-.708l6-6l.708.708Zm-6 2l-3.172 3.171a2.329 2.329 0 0 0 3.293 3.293l3.171-3.172l.708.708l-3.172 3.171A3.328 3.328 0 1 1 .975 9.318l3.171-3.172l.708.708Z" clip-rule="evenodd"/>`
	linkedinOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M4.5 6v5m6 0V8.5a2 2 0 1 0-4 0V11V6M4 4.5h1M1.5.5h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`
	linkedinSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-12ZM5 5H4V4h1v1Zm-1 6V6h1v5H4Zm4.5-4A1.5 1.5 0 0 0 7 8.5V11H6V6h1v.5a2.5 2.5 0 0 1 4 2V11h-1V8.5A1.5 1.5 0 0 0 8.5 7Z" clip-rule="evenodd"/>`
	linuxAltOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M2.5 9.662c0-.758.224-1.498.645-2.129l.565-.847a7.203 7.203 0 0 0 1.07-2.583l.328-1.642a2.44 2.44 0 0 1 4.784 0l.329 1.642a7.18 7.18 0 0 0 1.07 2.583l.564.847c.42.63.645 1.371.645 2.129m-7.392 3.637c.386.13.8.201 1.23.201h2.324c.43 0 .844-.07 1.23-.201M6.5 5.5L4.947 8.606a2 2 0 0 0 1.79 2.894h1.527a2 2 0 0 0 1.789-2.894L8.5 5.5M6.5 3v1.5m2-1.5v1.5m-1.894-.053L5.5 5l.586.586a2 2 0 0 0 2.828 0L9.5 5l-1.106-.553a2 2 0 0 0-1.788 0ZM.77 11.325l.479-1.196a1 1 0 0 1 .928-.629h1.164a1 1 0 0 1 .919.606l.93 2.172a1 1 0 0 1-.319 1.194l-.738.553a1 1 0 0 1-1.24-.031l-1.835-1.529a1 1 0 0 1-.288-1.14Zm13.46 0l-.479-1.196a1 1 0 0 0-.928-.629h-1.164a1 1 0 0 0-.919.606l-.93 2.172a1 1 0 0 0 .319 1.194l.738.553a1 1 0 0 0 1.24-.031l1.835-1.529a1 1 0 0 0 .288-1.14Z"/>`
	linuxAltSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4.54 2.407A3.013 3.013 0 0 1 7.5 0a3.013 3.013 0 0 1 2.96 2.407l.338 1.672a6.796 6.796 0 0 0 1.022 2.449l.58.863c.363.539.6 1.148.698 1.782a1.54 1.54 0 0 1 1.3.955l.492 1.22a1.52 1.52 0 0 1-.445 1.74l-1.884 1.558a1.55 1.55 0 0 1-1.91.048l-.76-.564a1.256 1.256 0 0 1-.029-.023c-.372.1-.764.154-1.168.154H6.306c-.404 0-.796-.053-1.168-.154l-.03.023l-.759.564a1.55 1.55 0 0 1-1.91-.048L.554 13.089a1.52 1.52 0 0 1-.444-1.742l.492-1.219a1.54 1.54 0 0 1 1.3-.955a4.397 4.397 0 0 1 .697-1.782l.58-.863a6.797 6.797 0 0 0 1.023-2.449l.337-1.672Zm6.304 7.069a1.53 1.53 0 0 0-.488.618l-.768 1.777a2.581 2.581 0 0 1-1.303.353h-1.57c-.467 0-.915-.126-1.303-.353l-.768-1.777a1.528 1.528 0 0 0-.489-.618c.026-.322.114-.641.264-.938l1.258-2.495l.007.007a2.583 2.583 0 0 0 3.632 0l.007-.007l1.258 2.495c.15.297.238.616.264.938ZM9.041 4.269V3.056H8.014v.801c.218.044.431.117.634.218l.393.194Zm-2.055-.412v-.801H5.96v1.213l.393-.194c.203-.101.417-.174.634-.218Zm-4.953 6.33a.514.514 0 0 0-.477.32l-.492 1.219a.507.507 0 0 0 .148.58l1.884 1.557a.517.517 0 0 0 .637.017l.759-.565a.507.507 0 0 0 .164-.608L3.7 10.495a.514.514 0 0 0-.472-.309H2.033Zm11.411.32a.514.514 0 0 0-.477-.32h-1.195a.514.514 0 0 0-.472.308l-.956 2.212a.507.507 0 0 0 .164.608l.759.565c.19.141.454.134.637-.017l1.884-1.557a.507.507 0 0 0 .148-.58l-.492-1.22ZM6.811 4.986a1.552 1.552 0 0 1 1.378 0l.498.247l-.097.097a1.55 1.55 0 0 1-2.18 0l-.097-.097l.498-.247Z" clip-rule="evenodd"/>`
	linuxOutlineInnerSVG                     = `<path fill="currentColor" d="m.5 14.5l-.354-.354A.5.5 0 0 0 .5 15v-.5Zm.656-.656l-.353-.354l.353.354ZM14.5 14.5v.5a.5.5 0 0 0 .354-.854l-.354.354Zm-11-4l-.224-.447a.5.5 0 0 0-.13.8L3.5 10.5Zm.87-.435l.223.447l-.223-.447Zm6.26 0l-.223.447l.223-.447Zm.87.435l.354.354a.5.5 0 0 0-.13-.801l-.224.447Zm-7.414.586l.353-.354l-.353.354ZM.854 14.854l.656-.657l-.707-.707l-.657.656l.708.708ZM2 13.014V6.5H1v6.514h1ZM13 6.5v6.514h1V6.5h-1Zm.49 7.697l.656.657l.708-.708l-.657-.656l-.707.707ZM14.5 14H.5v1h14v-1Zm-1.5-.986c0 .444.176.87.49 1.183l.707-.707a.673.673 0 0 1-.197-.476h-1ZM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0v1ZM2 6.5A5.5 5.5 0 0 1 7.5 1V0A6.5 6.5 0 0 0 1 6.5h1Zm-.49 7.697c.314-.314.49-.74.49-1.183H1c0 .179-.07.35-.197.476l.707.707Zm2.214-3.25l.87-.434l-.448-.895l-.87.435l.448.894Zm6.683-.434l.87.434l.447-.894l-.87-.435l-.447.894Zm.74-.367l-.586.586l.707.707l.586-.585l-.708-.708Zm-6.708.586l-.585-.586l-.708.708l.586.585l.707-.707ZM7.5 12a4.328 4.328 0 0 1-3.06-1.268l-.708.707c1 1 2.355 1.561 3.768 1.561v-1Zm3.06-1.268A4.328 4.328 0 0 1 7.5 12v1a5.328 5.328 0 0 0 3.768-1.56l-.707-.708Zm-5.967-.22a6.5 6.5 0 0 1 5.814 0l.447-.894a7.5 7.5 0 0 0-6.708 0l.447.894ZM7 6.5v1h1v-1H7Zm-3 1v-1H3v1h1ZM5.5 9A1.5 1.5 0 0 1 4 7.5H3A2.5 2.5 0 0 0 5.5 10V9ZM7 7.5A1.5 1.5 0 0 1 5.5 9v1A2.5 2.5 0 0 0 8 7.5H7ZM5.5 5A1.5 1.5 0 0 1 7 6.5h1A2.5 2.5 0 0 0 5.5 4v1Zm0-1A2.5 2.5 0 0 0 3 6.5h1A1.5 1.5 0 0 1 5.5 5V4ZM11 6.5v1h1v-1h-1ZM9.5 9A1.5 1.5 0 0 1 8 7.5H7A2.5 2.5 0 0 0 9.5 10V9ZM11 7.5A1.5 1.5 0 0 1 9.5 9v1A2.5 2.5 0 0 0 12 7.5h-1ZM9.5 5A1.5 1.5 0 0 1 11 6.5h1A2.5 2.5 0 0 0 9.5 4v1Zm0-1A2.5 2.5 0 0 0 7 6.5h1A1.5 1.5 0 0 1 9.5 5V4ZM5 7v1h1V7H5Zm4 0v1h1V7H9Z"/>`
	linuxSolidInnerSVG                       = `<path fill="currentColor" d="M5 8V7h1v1H5Zm4 0V7h1v1H9Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 0 1 13 0v6.514c0 .179.07.35.197.476l.657.656A.5.5 0 0 1 14.5 15H.5a.5.5 0 0 1-.354-.854l.657-.656A.673.673 0 0 0 1 13.014V6.5Zm3 0a1.5 1.5 0 1 1 3 0v1a1.5 1.5 0 1 1-3 0v-1Zm4 0a1.5 1.5 0 1 1 3 0v1a1.5 1.5 0 0 1-3 0v-1Zm-3.407 4.012a6.5 6.5 0 0 1 5.814 0l.249.125l-.095.095a4.329 4.329 0 0 1-6.122 0l-.095-.095l.249-.125Z" clip-rule="evenodd"/>`
	listLayoutOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M7 3.5h8m-8 8h8M1.5 1.5h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1Zm0 8h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1Z"/>`
	listLayoutSolidInnerSVG                  = `<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v2A1.5 1.5 0 0 0 1.5 6h2A1.5 1.5 0 0 0 5 4.5v-2A1.5 1.5 0 0 0 3.5 1h-2ZM7 4h8V3H7v1ZM1.5 9A1.5 1.5 0 0 0 0 10.5v2A1.5 1.5 0 0 0 1.5 14h2A1.5 1.5 0 0 0 5 12.5v-2A1.5 1.5 0 0 0 3.5 9h-2ZM7 12h8v-1H7v1Z"/>`
	listOrderedOutlineInnerSVG               = `<path fill="currentColor" d="m.5 13.5l-.354-.354A.5.5 0 0 0 .5 14v-.5Zm1-11H2V2h-.5v.5ZM5 8h10V7H5v1Zm0-4h10V3H5v1Zm0 8h10v-1H5v1Zm-2 1H.5v1H3v-1Zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792l.708.708ZM1.793 10H1.5v1h.293v-1ZM1.5 10A1.5 1.5 0 0 0 0 11.5h1a.5.5 0 0 1 .5-.5v-1ZM3 11.207C3 10.54 2.46 10 1.793 10v1c.114 0 .207.093.207.207h1Zm-.354.854c.227-.227.354-.534.354-.854H2a.207.207 0 0 1-.06.147l.706.707ZM0 6h3V5H0v1Zm2-.5v-3H1v3h1ZM1.5 2H0v1h1.5V2Z"/>`
	listOrderedSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 2h2v3h1v1H0V5h1V3H0V2Zm15 2H5V3h10v1Zm0 4H5V7h10v1ZM0 11.5A1.5 1.5 0 0 1 1.5 10h.293a1.207 1.207 0 0 1 .853 2.06l-.939.94H3v1H.5a.5.5 0 0 1-.354-.854l1.793-1.792A.207.207 0 0 0 1.793 11H1.5a.5.5 0 0 0-.5.5H0Zm15 .5H5v-1h10v1Z" clip-rule="evenodd"/>`
	listUnorderedOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M4 7.5h11m-15 0h2m2-4h11m-15 0h2m2 8h11m-15 0h2"/>`
	listUnorderedSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M2 4H0V3h2v1Zm13 0H4V3h11v1ZM2 8H0V7h2v1Zm13 0H4V7h11v1ZM2 12H0v-1h2v1Zm13 0H4v-1h11v1Z" clip-rule="evenodd"/>`
	litecoinOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="m5.5 1.5l-3 12H13m-12-5l6-3"/>`
	litecoinSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="m3.714 6.584l1.3-5.205l.971.242l-1.093 4.374l1.884-.942l.448.894l-2.652 1.326L3.14 13H13v1H1.86l1.534-6.138l-2.17 1.085l-.448-.894l2.938-1.469Z" clip-rule="evenodd"/>`
	loaderOutlineInnerSVG                    = `<path fill="currentColor" d="M8 1V.5H7V1h1ZM7 4.5V5h1v-.5H7Zm1 6V10H7v.5h1ZM7 14v.5h1V14H7ZM4.5 7.995H5v-1h-.5v1Zm-3.5-1H.5v1H1v-1ZM14 8h.5V7H14v1Zm-3.5-1.005H10v1h.5v-1ZM7 1v3.5h1V1H7Zm0 9.5V14h1v-3.5H7ZM4.5 6.995H1v1h3.5v-1ZM14 7l-3.5-.005v1L14 8V7ZM2.147 2.854l3 3l.708-.708l-3-3l-.708.708Zm10-.708l-3 3l.708.708l3-3l-.708-.708ZM2.854 12.854l3-3l-.708-.708l-3 3l.708.708Zm6.292-3l3 3l.708-.708l-3-3l-.708.708Z"/>`
	loaderSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M8 .5V5H7V.5h1ZM5.146 5.854l-3-3l.708-.708l3 3l-.708.708Zm4-.708l3-3l.708.708l-3 3l-.708-.708Zm.855 1.849L14.5 7l-.002 1l-4.5-.006l.002-1Zm-9.501 0H5v1H.5v-1Zm5.354 2.859l-3 3l-.708-.708l3-3l.708.708Zm6.292 3l-3-3l.708-.708l3 3l-.708.708ZM8 10v4.5H7V10h1Z" clip-rule="evenodd"/>`
	locationOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5.5v14m7-7.005H.5m13 0a6.006 6.006 0 0 1-6 6.005c-3.313 0-6-2.694-6-6.005a5.999 5.999 0 0 1 6-5.996a6 6 0 0 1 6 5.996Z"/>`
	locationSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M8 1.018V0H7v1.018a6.5 6.5 0 0 0-5.981 5.977H0v1h1.019A6.508 6.508 0 0 0 7 13.981V15h1v-1.019a6.508 6.508 0 0 0 5.981-5.986H15v-1h-1.019A6.5 6.5 0 0 0 8 1.018ZM8 3v3.995h4v1H8V12H7V7.995H3v-1h4V3h1Z" clip-rule="evenodd"/>`
	lockCircleOutlineInnerSVG                = `<path fill="currentColor" d="M5.5 7h4V6h-4v1Zm4.5.5v3h1v-3h-1ZM9.5 11h-4v1h4v-1ZM5 10.5v-3H4v3h1Zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12v-1Zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6v1Zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5V6Zm.5.5v-1H5v1h1Zm3-1v1h1v-1H9ZM7.5 4A1.5 1.5 0 0 1 9 5.5h1A2.5 2.5 0 0 0 7.5 3v1ZM6 5.5A1.5 1.5 0 0 1 7.5 4V3A2.5 2.5 0 0 0 5 5.5h1Zm-5 2A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1Zm1 0A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1Z"/>`
	lockCircleSolidInnerSVG                  = `<path fill="currentColor" d="M7.5 4A1.5 1.5 0 0 0 6 5.5V6h3v-.5A1.5 1.5 0 0 0 7.5 4Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15ZM5 5.5v.585A1.5 1.5 0 0 0 4 7.5v3A1.5 1.5 0 0 0 5.5 12h4a1.5 1.5 0 0 0 1.5-1.5v-3a1.5 1.5 0 0 0-1-1.415V5.5a2.5 2.5 0 0 0-5 0Z" clip-rule="evenodd"/>`
	lockOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M4.5 6.5v-3a3 3 0 0 1 6 0v3m-8 0h10a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`
	lockSmallOutlineInnerSVG                 = `<path fill="currentColor" d="M5.5 7h4V6h-4v1Zm4.5.5v3h1v-3h-1ZM9.5 11h-4v1h4v-1ZM5 10.5v-3H4v3h1Zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12v-1Zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6v1Zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5V6Zm.5.5v-1H5v1h1Zm3-1v1h1v-1H9ZM7.5 4A1.5 1.5 0 0 1 9 5.5h1A2.5 2.5 0 0 0 7.5 3v1ZM6 5.5A1.5 1.5 0 0 1 7.5 4V3A2.5 2.5 0 0 0 5 5.5h1Z"/>`
	lockSmallSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M10 5.5v.585A1.5 1.5 0 0 1 11 7.5v3A1.5 1.5 0 0 1 9.5 12h-4A1.5 1.5 0 0 1 4 10.5v-3a1.5 1.5 0 0 1 1-1.415V5.5a2.5 2.5 0 0 1 5 0Zm-4 0a1.5 1.5 0 1 1 3 0V6H6v-.5Z" clip-rule="evenodd"/>`
	lockSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M11 3.5V6h1.5A1.5 1.5 0 0 1 14 7.5v6a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-6A1.5 1.5 0 0 1 2.5 6H4V3.5a3.5 3.5 0 1 1 7 0Zm-6 0a2.5 2.5 0 0 1 5 0V6H5V3.5Z" clip-rule="evenodd"/>`
	logoutOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m13.5 7.5l-3 3.25m3-3.25l-3-3m3 3H4m4 6H1.5v-12H8"/>`
	logoutSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h7v1H2v11h6v1H1V1Zm9.854 3.146l3.34 3.34l-3.327 3.603l-.734-.678L12.358 8H4V7h8.293l-2.147-2.146l.708-.708Z" clip-rule="evenodd"/>`
	loopOutlineInnerSVG                      = `<path fill="currentColor" d="m14.5 3.5l.354.354a.5.5 0 0 0 0-.708L14.5 3.5Zm-14 8l-.354-.354a.5.5 0 0 0 0 .708L.5 11.5ZM11.146.854l3 3l.708-.708l-3-3l-.708.708Zm3 2.292l-3 3l.708.708l3-3l-.708-.708Zm-10.292 11l-3-3l-.708.708l3 3l.708-.708Zm-3-2.292l3-3l-.708-.708l-3 3l.708.708ZM.5 12h11v-1H.5v1ZM15 8.5V7h-1v1.5h1ZM11.5 12A3.5 3.5 0 0 0 15 8.5h-1a2.5 2.5 0 0 1-2.5 2.5v1Zm3-9h-11v1h11V3ZM0 6.5V8h1V6.5H0ZM3.5 3A3.5 3.5 0 0 0 0 6.5h1A2.5 2.5 0 0 1 3.5 4V3Z"/>`
	loopSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M13.293 3L11.146.854l.708-.708l3 3a.5.5 0 0 1 0 .708l-3 3l-.708-.708L13.293 4H3.5A2.5 2.5 0 0 0 1 6.5V8H0V6.5A3.5 3.5 0 0 1 3.5 3h9.793ZM15 7v1.5a3.5 3.5 0 0 1-3.5 3.5H1.707l2.147 2.146l-.708.708l-3-3a.5.5 0 0 1 0-.707l3-3l.708.707L1.707 11H11.5A2.5 2.5 0 0 0 14 8.5V7h1Z" clip-rule="evenodd"/>`
	magsafeOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M3.5 2.5v-2h8v2m-6 9V15m4-3.5V15m-5-5.5v2h6v-2m-9-7h12v7h-12v-7Z"/>`
	magsafeSolidInnerSVG                     = `<path fill="currentColor" d="M12 0H3v3h9V0Zm2 4H1v6h3v2h1v3h1v-3h3v3h1v-3h1v-2h3V4Z"/>`
	markdownOutlineInnerSVG                  = `<path fill="currentColor" d="m2.5 5.5l.354-.354A.5.5 0 0 0 2 5.5h.5Zm2 2l-.354.354l.354.353l.354-.353L4.5 7.5Zm2-2H7a.5.5 0 0 0-.854-.354L6.5 5.5Zm4 4l-.354.354l.354.353l.354-.353L10.5 9.5ZM1.5 3h12V2h-12v1Zm12.5.5v8h1v-8h-1Zm-.5 8.5h-12v1h12v-1ZM1 11.5v-8H0v8h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 13v-1Zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 2v1Zm-12-1A1.5 1.5 0 0 0 0 3.5h1a.5.5 0 0 1 .5-.5V2ZM3 10V5.5H2V10h1Zm-.854-4.146l2 2l.708-.708l-2-2l-.708.708Zm2.708 2l2-2l-.708-.708l-2 2l.708.708ZM6 5.5V10h1V5.5H6Zm4-.5v4.5h1V5h-1ZM8.146 7.854l2 2l.708-.708l-2-2l-.708.708Zm2.708 2l2-2l-.708-.708l-2 2l.708.708Z"/>`
	markdownSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M0 3.5A1.5 1.5 0 0 1 1.5 2h12A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5v-8ZM10 5v3.293L8.854 7.146l-.708.708l2 2a.5.5 0 0 0 .708 0l2-2l-.707-.708L11 8.293V5h-1Zm-7.146.146A.5.5 0 0 0 2 5.5V10h1V6.707l1.5 1.5l1.5-1.5V10h1V5.5a.5.5 0 0 0-.854-.354L4.5 6.793L2.854 5.146Z" clip-rule="evenodd"/>`
	mediumOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m4.5 4.5l3 4l3-4m-6 0H3m1.5 0V11m6-6.5H12m-1.5 0V11M3 10.5h3m3 0h3M1.5.5h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`
	mediumSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-12ZM4 5H3V4h1.5a.5.5 0 0 1 .4.2l2.6 3.467l2.593-3.458A.5.5 0 0 1 10.5 4H12v1h-1v5h1v1H9v-1h1V6L7.5 9.333L5 6v4h1v1H3v-1h1V5Z" clip-rule="evenodd"/>`
	menuOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M0 5.5h15m-15-4h15m-15 8h15m-15 4h15"/>`
	menuSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M15 2H0V1h15v1Zm0 4H0V5h15v1Zm0 4H0V9h15v1Zm0 4H0v-1h15v1Z" clip-rule="evenodd"/>`
	messageMinusOutlineInnerSVG              = `<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222v.5Zm2 2.998l-.416.277a.5.5 0 0 0 .832 0l-.416-.277Zm2-2.998v-.5a.5.5 0 0 0-.416.222l.416.278Zm-4.416.277l2 2.998l.832-.555l-2-2.998l-.832.555Zm2.832 2.998l2-2.998l-.832-.555l-2 2.998l.832.555ZM9.5 11.993h4v-1h-4v1Zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5v1Zm1.5-1.5V1.5h-1v8.994h1ZM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5h1ZM13.5 0h-12v1h12V0Zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5V0ZM0 1.5v8.993h1V1.5H0Zm0 8.993c0 .83.671 1.5 1.5 1.5v-1a.499.499 0 0 1-.5-.5H0Zm1.5 1.5h4v-1h-4v1ZM5 7h5V6H5v1Z"/>`
	messageMinusSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.993c0 .83-.671 1.5-1.5 1.5H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5V1.5ZM5 7h5V6H5v1Z" clip-rule="evenodd"/>`
	messageNoAccessOutlineInnerSVG           = `<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222v.5Zm2 2.998l-.416.277a.5.5 0 0 0 .832 0l-.416-.277Zm2-2.998v-.5a.5.5 0 0 0-.416.222l.416.278Zm-4.416.277l2 2.998l.832-.555l-2-2.998l-.832.555Zm2.832 2.998l2-2.998l-.832-.555l-2 2.998l.832.555ZM9.5 11.993h4v-1h-4v1Zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5v1Zm1.5-1.5V1.5h-1v8.994h1ZM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5h1ZM13.5 0h-12v1h12V0Zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5V0ZM0 1.5v8.993h1V1.5H0Zm0 8.993c0 .83.671 1.5 1.5 1.5v-1a.499.499 0 0 1-.5-.5H0Zm1.5 1.5h4v-1h-4v1ZM7.5 9A2.5 2.5 0 0 1 5 6.5H4A3.5 3.5 0 0 0 7.5 10V9ZM10 6.5A2.5 2.5 0 0 1 7.5 9v1A3.5 3.5 0 0 0 11 6.5h-1ZM7.5 4A2.5 2.5 0 0 1 10 6.5h1A3.5 3.5 0 0 0 7.5 3v1Zm0-1A3.5 3.5 0 0 0 4 6.5h1A2.5 2.5 0 0 1 7.5 4V3Zm1.646 1.146l-4 4l.708.708l4-4l-.708-.708Z"/>`
	messageNoAccessSolidInnerSVG             = `<path fill="currentColor" d="M7.5 4a2.5 2.5 0 0 0-2.086 3.879L8.88 4.414A2.488 2.488 0 0 0 7.5 4Zm0 5c-.51 0-.983-.152-1.379-.414L9.586 5.12A2.5 2.5 0 0 1 7.5 9Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.993c0 .83-.671 1.5-1.5 1.5H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5V1.5Zm4 5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/>`
	messageOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" d="m5.5 11.493l2 2.998l2-2.998h4a1 1 0 0 0 1-1V1.5a.999.999 0 0 0-1-.999h-12a1 1 0 0 0-1 1v8.993a1 1 0 0 0 1 1h4Z" clip-rule="evenodd"/>`
	messagePlusOutlineInnerSVG               = `<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222v.5Zm2 2.998l-.416.277a.5.5 0 0 0 .832 0l-.416-.277Zm2-2.998v-.5a.5.5 0 0 0-.416.222l.416.278Zm-4.416.277l2 2.998l.832-.555l-2-2.998l-.832.555Zm2.832 2.998l2-2.998l-.832-.555l-2 2.998l.832.555ZM9.5 11.993h4v-1h-4v1Zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5v1Zm1.5-1.5V1.5h-1v8.994h1ZM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5h1ZM13.5 0h-12v1h12V0Zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5V0ZM0 1.5v8.993h1V1.5H0Zm0 8.993c0 .83.671 1.5 1.5 1.5v-1a.499.499 0 0 1-.5-.5H0Zm1.5 1.5h4v-1h-4v1ZM7 4v5h1V4H7ZM5 7h5V6H5v1Z"/>`
	messagePlusSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.993c0 .83-.671 1.5-1.5 1.5H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5V1.5ZM7 9V7H5V6h2V4h1v2h2v1H8v2H7Z" clip-rule="evenodd"/>`
	messageSolidInnerSVG                     = `<path fill="currentColor" d="M1.5 0C.671 0 0 .67 0 1.5v8.993c0 .83.671 1.5 1.5 1.5h3.732l1.852 2.775a.5.5 0 0 0 .832 0l1.851-2.775H13.5c.829 0 1.5-.67 1.5-1.5V1.5c0-.83-.671-1.5-1.5-1.5h-12Z"/>`
	messageTextAltOutlineInnerSVG            = `<path fill="currentColor" d="M3.5 11.493H4v-.5h-.5v.5Zm0 2.998H3a.5.5 0 0 0 .8.4l-.3-.4Zm4-2.998v-.5h-.167l-.133.1l.3.4Zm-3-7.496H4v1h.5v-1Zm6 1h.5v-1h-.5v1Zm-6 1.998H4v1h.5v-1Zm4 1H9v-1h-.5v1ZM3 11.493v2.998h1v-2.998H3Zm.8 3.398l4-2.998l-.6-.8l-4 2.998l.6.8Zm3.7-2.898h6v-1h-6v1Zm6 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5v1Zm1.5-1.5V1.5h-1v8.994h1ZM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5h1ZM13.5 0h-12v1h12V0Zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5V0ZM0 1.5v8.993h1V1.5H0Zm0 8.993c0 .83.671 1.5 1.5 1.5v-1a.499.499 0 0 1-.5-.5H0Zm1.5 1.5h2v-1h-2v1Zm3-6.996h6v-1h-6v1Zm0 2.998h4v-1h-4v1Z"/>`
	messageTextAltSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.993c0 .83-.671 1.5-1.5 1.5H7.667L3.8 14.89a.5.5 0 0 1-.8-.4v-2.498H1.5c-.829 0-1.5-.67-1.5-1.5V1.5Zm4 2.497h7v1H4v-1Zm0 2.998h5v1H4v-1Z" clip-rule="evenodd"/>`
	messageTextOutlineInnerSVG               = `<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222v.5Zm2 2.998l-.416.277a.5.5 0 0 0 .832 0l-.416-.277Zm2-2.998v-.5a.5.5 0 0 0-.416.222l.416.278Zm-4.416.277l2 2.998l.832-.555l-2-2.998l-.832.555Zm2.832 2.998l2-2.998l-.832-.555l-2 2.998l.832.555ZM9.5 11.993h4v-1h-4v1Zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5v1Zm1.5-1.5V1.5h-1v8.994h1ZM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5h1ZM13.5 0h-12v1h12V0Zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5V0ZM0 1.5v8.993h1V1.5H0Zm0 8.993c0 .83.671 1.5 1.5 1.5v-1a.499.499 0 0 1-.5-.5H0Zm1.5 1.5h4v-1h-4v1ZM5 8h5V7H5v1ZM4 5h7V4H4v1Z"/>`
	messageTextSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.993c0 .83-.671 1.5-1.5 1.5H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5V1.5ZM11 5H4V4h7v1Zm-1 3H5V7h5v1Z" clip-rule="evenodd"/>`
	messageTickOutlineInnerSVG               = `<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222v.5Zm2 2.998l-.416.277a.5.5 0 0 0 .832 0l-.416-.277Zm2-2.998v-.5a.5.5 0 0 0-.416.222l.416.278ZM7 8l-.354.354l.378.377l.352-.402L7 8Zm-1.916 3.77l2 2.998l.832-.555l-2-2.998l-.832.555Zm2.832 2.998l2-2.998l-.832-.555l-2 2.998l.832.555ZM9.5 11.993h4v-1h-4v1Zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5v1Zm1.5-1.5V1.5h-1v8.994h1ZM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5h1ZM13.5 0h-12v1h12V0Zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5V0ZM0 1.5v8.993h1V1.5H0Zm0 8.993c0 .83.671 1.5 1.5 1.5v-1a.499.499 0 0 1-.5-.5H0Zm1.5 1.5h4v-1h-4v1Zm3.146-5.64l2 2l.708-.707l-2-2l-.708.708Zm2.73 1.976l3.5-4l-.752-.658l-3.5 4l.752.658Z"/>`
	messageTickSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.993c0 .83-.671 1.5-1.5 1.5H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5V1.5Zm7.024 7.23l3.852-4.402l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708L7.024 8.73Z" clip-rule="evenodd"/>`
	messageXOutlineInnerSVG                  = `<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222v.5Zm2 2.998l-.416.277a.5.5 0 0 0 .832 0l-.416-.277Zm2-2.998v-.5a.5.5 0 0 0-.416.222l.416.278Zm-4.416.277l2 2.998l.832-.555l-2-2.998l-.832.555Zm2.832 2.998l2-2.998l-.832-.555l-2 2.998l.832.555ZM9.5 11.993h4v-1h-4v1Zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5v1Zm1.5-1.5V1.5h-1v8.994h1ZM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5h1ZM13.5 0h-12v1h12V0Zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5V0ZM0 1.5v8.993h1V1.5H0Zm0 8.993c0 .83.671 1.5 1.5 1.5v-1a.499.499 0 0 1-.5-.5H0Zm1.5 1.5h4v-1h-4v1Zm3.646-7.14l4 4l.708-.707l-4-4l-.708.708Zm.708 4l4-4l-.708-.707l-4 4l.708.708Z"/>`
	messageXSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.993c0 .83-.671 1.5-1.5 1.5H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5V1.5Zm9.146 7.354L7.5 7.207L5.854 8.854l-.708-.708L6.793 6.5L5.146 4.854l.708-.708L7.5 5.793l1.646-1.647l.708.708L8.207 6.5l1.647 1.646l-.708.708Z" clip-rule="evenodd"/>`
	messengerOutlineInnerSVG                 = `<path fill="currentColor" d="m2.935 12.264l.482.132a.5.5 0 0 0-.164-.517l-.318.385ZM2.326 14.5l-.482-.131a.5.5 0 0 0 .694.584l-.212-.453Zm2.435-1.141l.188-.463a.5.5 0 0 0-.4.01l.212.453ZM6.5 6.5l.3-.4l-.293-.22l-.298.213l.291.407Zm2 1.5l-.3.4l.316.237l.304-.253L8.5 8Zm-1-8C3.379 0 0 3.201 0 7.196h1C1 3.795 3.889 1 7.5 1V0ZM0 7.196c0 2.188 1.023 4.139 2.617 5.454l.636-.771C1.87 10.739 1 9.062 1 7.196H0Zm2.452 4.937l-.608 2.236l.965.262l.608-2.235l-.965-.263Zm.086 2.82l2.435-1.142l-.424-.905l-2.435 1.141l.424.906Zm2.035-1.131c.9.366 1.89.57 2.927.57v-1a6.764 6.764 0 0 1-2.55-.496l-.377.926Zm2.927.57c4.121 0 7.5-3.202 7.5-7.196h-1c0 3.4-2.889 6.195-6.5 6.195v1ZM15 7.195C15 3.2 11.621 0 7.5 0v1C11.111 1 14 3.795 14 7.196h1ZM3.29 9.406l3.5-2.5l-.58-.813l-3.5 2.5l.58.814ZM6.2 6.9l2 1.5l.6-.8l-2-1.5l-.6.8Zm2.62 1.484l3-2.5l-.64-.768l-3 2.5l.64.768Z"/>`
	messengerSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.196C0 3.2 3.379 0 7.5 0S15 3.201 15 7.196c0 3.994-3.379 7.195-7.5 7.195a7.77 7.77 0 0 1-2.72-.489l-2.242 1.05a.5.5 0 0 1-.694-.583l.526-1.932C.918 11.129 0 9.269 0 7.196Zm8.516 1.441l3.304-2.753l-.64-.768l-2.696 2.247L6.507 5.88L2.71 8.593l.582.814L6.493 7.12l2.023 1.517Z" clip-rule="evenodd"/>`
	microSdCardOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M6.5 3v3m2-3v3m2-3v3m-8 8.5h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1v5l-2 2v5a1 1 0 0 0 1 1Z"/>`
	microSdCardSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h8A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V8.293l2-2V1.5ZM6 3v3h1V3H6Zm2 0v3h1V3H8Zm2 3V3h1v3h-1Z" clip-rule="evenodd"/>`
	microphoneOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M2.5 4v2.5a5 5 0 0 0 5 5m5-7.5v2.5a5 5 0 0 1-5 5m0 0V15M5 14.5h5m-.5-12v4a2 2 0 1 1-4 0v-4a2 2 0 1 1 4 0Z"/>`
	microphoneSolidInnerSVG                  = `<path fill="currentColor" d="M5 2.5a2.5 2.5 0 0 1 5 0v4a2.5 2.5 0 0 1-5 0v-4Z"/><path fill="currentColor" d="M2 4v2.5a5.5 5.5 0 0 0 5 5.478V14H5v1h5v-1H8v-2.022A5.5 5.5 0 0 0 13 6.5V4h-1v2.5a4.5 4.5 0 0 1-9 0V4H2Z"/>`
	minimiseAltOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M13 5.5H9.5m0 0V2m0 3.5l4-4M5.5 13V9.5m0 0H2m3.5 0l-4 4"/>`
	minimiseAltSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L10.707 5H13v1H9V2h1v2.293l3.146-3.147l.708.708ZM2 9h4v4H5v-2.293l-3.146 3.147l-.708-.707L4.293 10H2V9Z" clip-rule="evenodd"/>`
	minimiseOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M9.5 9.5H13m-3.5 0V13m0-3.5l4 4m-.5-8H9.5m0 0V2m0 3.5l4-4M2 5.5h3.5m0 0V2m0 3.5l-4-4m4 11.5V9.5m0 0H2m3.5 0l-4 4"/>`
	minimiseSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L10.707 5H13v1H9V2h1v2.293l3.146-3.147l.708.708ZM4.293 5L1.146 1.854l.708-.708L5 4.293V2h1v4H2V5h2.293ZM2 9h4v4H5v-2.293l-3.146 3.147l-.708-.707L4.293 10H2V9Zm7 0h4v1h-2.293l3.147 3.146l-.708.708L10 10.707V13H9V9Z" clip-rule="evenodd"/>`
	minusCircleOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M4 7.5h7m-3.5 7a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	minusCircleSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 8h7V7H4v1Z" clip-rule="evenodd"/>`
	minusOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M1 7.5h13"/>`
	minusSmallOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M4 7.5h7"/>`
	minusSmallSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M11 8H4V7h7v1Z" clip-rule="evenodd"/>`
	minusSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M14 8H1V7h13v1Z" clip-rule="evenodd"/>`
	mobileOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M6 11.5h3m-5.5 3h8a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`
	mobileSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M2 1.5A1.5 1.5 0 0 1 3.5 0h8A1.5 1.5 0 0 1 13 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 13.5v-12ZM6 12h3v-1H6v1Z" clip-rule="evenodd"/>`
	moneyOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M11 10.5h1.5V9M11 4.5h1.5V6M4 4.5H2.5V6m0 3v1.5H4m3.5-1a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6-7h12a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1Z"/>`
	moneySolidInnerSVG                       = `<path fill="currentColor" d="M6 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 3.5A1.5 1.5 0 0 1 1.5 2h12A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5v-8ZM4 4H2v2h1V5h1V4Zm8 1h-1V4h2v2h-1V5ZM7.5 5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Zm3.5 5v1h2V9h-1v1h-1ZM2 9h1v1h1v1H2V9Z" clip-rule="evenodd"/>`
	moneyStackOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M0 12.5h15m-15 2h15M2.5 4V2.5H4m7 0h1.5V4m-10 3v1.5H4m7 0h1.5V7m-5 .5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6-7h12a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1Z"/>`
	moneyStackSolidInnerSVG                  = `<path fill="currentColor" d="M6 5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v8A1.5 1.5 0 0 0 1.5 11h12A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-12ZM4 2H2v2h1V3h1V2Zm3.5 1a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM12 3h-1V2h2v2h-1V3ZM3 7H2v2h2V8H3V7Zm8 2V8h1V7h1v2h-2Z" clip-rule="evenodd"/><path fill="currentColor" d="M0 12v1h15v-1H0Zm0 2v1h15v-1H0Z"/>`
	mongodbOutlineInnerSVG                   = `<path fill="currentColor" d="m7.5.5l.369-.338a.5.5 0 0 0-.738 0L7.5.5Zm0 13l-.393.309a.5.5 0 0 0 .786 0L7.5 13.5ZM4.623 9.838l-.393.31l.393-.31Zm.246-6.467L4.5 3.032l.369.337Zm5.262 0l.369-.338l-.369.337Zm.246 6.467l.393.31l-.393-.31ZM8 15V.5H7V15h1Zm-.107-1.809L5.016 9.53l-.786.618l2.877 3.662l.786-.618ZM5.237 3.708L7.87.838L7.13.162L4.5 3.032l.736.676ZM7.131.838l2.632 2.87l.737-.675L7.869.163L7.13.837Zm2.853 8.691l-2.877 3.662l.786.618l2.877-3.662l-.786-.618Zm-.221-5.82a4.5 4.5 0 0 1 .22 5.82l.787.618a5.5 5.5 0 0 0-.27-7.114l-.737.675Zm-4.747 5.82a4.5 4.5 0 0 1 .221-5.82L4.5 3.032a5.5 5.5 0 0 0-.27 7.114l.786-.618Z"/>`
	mongodbSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.869.162a.5.5 0 0 0-.738 0l-2.63 2.87a5.5 5.5 0 0 0-.271 7.115L7 13.673V15h1v-1.327l2.77-3.526a5.5 5.5 0 0 0-.27-7.114L7.869.163ZM8 3a.5.5 0 0 0-1 0v7.5a.5.5 0 0 0 1 0V3Z" clip-rule="evenodd"/>`
	moodFlatOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M4 5.5h1m5 0h1m-7 4h7m-3.5 5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	moodFlatSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1Zm1 3v1H4V9h7Z" clip-rule="evenodd"/>`
	moodFrownOutlineInnerSVG                 = `<path fill="currentColor" d="m5.29 10.187l-.424.266l.531.847l.424-.265l-.53-.848Zm4.897-.244l.5.035l.069-.998l-.5-.034l-.069.997Zm-4.366 1.092a7.278 7.278 0 0 1 4.367-1.092l.069-.997a8.278 8.278 0 0 0-4.967 1.241l.53.848ZM4 6h1V5H4v1Zm6 0h1V5h-1v1Zm-2.5 8A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	moodFrownSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm1.82 5.035a7.278 7.278 0 0 1 4.368-1.092l.498.035l.07-.998l-.5-.034a8.278 8.278 0 0 0-4.966 1.241l-.424.266l.531.847l.424-.265ZM11 6h-1V5h1v1Z" clip-rule="evenodd"/>`
	moodLaughOutlineInnerSVG                 = `<path fill="currentColor" d="M4.5 7.5V7a.5.5 0 0 0-.5.5h.5Zm6 0h.5a.5.5 0 0 0-.5-.5v.5Zm0-.5h-6v1h6V7Zm-3 4C9.47 11 11 9.47 11 7.5h-1C10 8.918 8.918 10 7.5 10v1ZM4 7.5C4 9.47 5.53 11 7.5 11v-1C6.082 10 5 8.918 5 7.5H4ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1Z"/>`
	moodLaughSolidInnerSVG                   = `<path fill="currentColor" d="M7.5 10c-1.246 0-2.233-.835-2.454-2h4.908c-.221 1.165-1.208 2-2.454 2Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm.5 1a.5.5 0 0 0-.5.5C4 9.47 5.53 11 7.5 11S11 9.47 11 7.5a.5.5 0 0 0-.5-.5h-6ZM11 6h-1V5h1v1Z" clip-rule="evenodd"/>`
	moodSadOutlineInnerSVG                   = `<path fill="currentColor" d="m4.1 9.7l-.3.4l.8.6l.3-.4l-.8-.6Zm6 .6l.3.4l.8-.6l-.3-.4l-.8.6ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1Zm.9 3.7c-1.7-2.267-5.1-2.267-6.8 0l.8.6a3.25 3.25 0 0 1 5.2 0l.8-.6Z"/>`
	moodSadSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1Zm-5.1 4.3a3.25 3.25 0 0 1 5.2 0l.8-.6c-1.7-2.267-5.1-2.267-6.8 0l.8.6Z" clip-rule="evenodd"/>`
	moodSmileOutlineInnerSVG                 = `<path fill="currentColor" d="m4.9 8.7l-.3-.4l-.8.6l.3.4l.8-.6Zm6 .6l.3-.4l-.8-.6l-.3.4l.8.6ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1Zm.1 2.7a3.25 3.25 0 0 1-5.2 0l-.8.6c1.7 2.267 5.1 2.267 6.8 0l-.8-.6Z"/>`
	moodSmileSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1ZM4.9 8.7a3.25 3.25 0 0 0 5.2 0l.8.6c-1.7 2.267-5.1 2.267-6.8 0l.8-.6Z" clip-rule="evenodd"/>`
	moodSurprisedOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M4 5.5h1m5 0h1m-3.5 9a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm-.5-7h1a1.5 1.5 0 1 1 0 3H7a1.5 1.5 0 1 1 0-3Z"/>`
	moodSurprisedSolidInnerSVG               = `<path fill="currentColor" d="M7 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2H7Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1ZM5 9a2 2 0 0 1 2-2h1a2 2 0 1 1 0 4H7a2 2 0 0 1-2-2Z" clip-rule="evenodd"/>`
	moodTongueOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M4 5.5h1m5 0h1m-7 3h7m-5.5 0v2a2 2 0 1 0 4 0v-2m-2 6a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	moodTongueSolidInnerSVG                  = `<path fill="currentColor" d="M6 10.5V9h3v1.5a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM5 6H4V5h1v1Zm6 0h-1V5h1v1ZM4 9h1v1.5a2.5 2.5 0 0 0 5 0V9h1V8H4v1Z" clip-rule="evenodd"/>`
	moonOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M1.66 11.362A6.5 6.5 0 0 0 7.693.502a7 7 0 1 1-6.031 10.86Z"/>`
	moonSolidInnerSVG                        = `<path fill="currentColor" d="M7.707.003a.5.5 0 0 0-.375.846a6 6 0 0 1-5.569 10.024a.5.5 0 0 0-.519.765A7.5 7.5 0 1 0 7.707.003Z"/>`
	moreHorizontalOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M3 7.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/>`
	moreHorizontalSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 7.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/>`
	moreVerticalOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M7.5 3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/>`
	moreVerticalSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M6.5 2.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm0 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm0 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/>`
	mouseOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M7.5 4v3m0 7.5a5 5 0 0 1-5-5v-4a5 5 0 0 1 10 0v4a5 5 0 0 1-5 5Z"/>`
	mouseSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M2 5.5a5.5 5.5 0 1 1 11 0v4a5.5 5.5 0 1 1-11 0v-4ZM7 4v3h1V4H7Z" clip-rule="evenodd"/>`
	movOutlineInnerSVG                       = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-8 6l.354-.354A.5.5 0 0 0 2 6.5h.5Zm1 1l-.354.354l.354.353l.354-.353L3.5 7.5Zm1-1H5a.5.5 0 0 0-.854-.354L4.5 6.5Zm6 3H10v.207l.146.147l.354-.354Zm1 1l-.354.354a.5.5 0 0 0 .708 0L11.5 10.5Zm1-1l.354.354l.146-.147V9.5h-.5ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM3 11V6.5H2V11h1Zm-.854-4.146l1 1l.708-.708l-1-1l-.708.708Zm1.708 1l1-1l-.708-.708l-1 1l.708.708ZM4 6.5V11h1V6.5H4Zm4 1v2h1v-2H8Zm-1 2v-2H6v2h1Zm.5.5a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 11v-1Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 9.5H8ZM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6v1Zm0-1A1.5 1.5 0 0 0 6 7.5h1a.5.5 0 0 1 .5-.5V6ZM10 6v3.5h1V6h-1Zm.146 3.854l1 1l.708-.708l-1-1l-.708.708Zm1.708 1l1-1l-.708-.708l-1 1l.708.708ZM13 9.5V6h-1v3.5h1Z"/>`
	movSolidInnerSVG                         = `<path fill="currentColor" d="M7 7.5a.5.5 0 0 1 1 0v2a.5.5 0 0 1-1 0v-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM7.5 6A1.5 1.5 0 0 0 6 7.5v2a1.5 1.5 0 0 0 3 0v-2A1.5 1.5 0 0 0 7.5 6Zm-4.646.146A.5.5 0 0 0 2 6.5V11h1V7.707l.5.5l.5-.5V11h1V6.5a.5.5 0 0 0-.854-.354l-.646.647l-.646-.647ZM10 6h1v3.293l.5.5l.5-.5V6h1v3.707l-1.146 1.147a.5.5 0 0 1-.708 0L10 9.707V6Z" clip-rule="evenodd"/>`
	mpFourOutlineInnerSVG                    = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-8 6l.354-.354A.5.5 0 0 0 2 6.5h.5Zm1 1l-.354.354l.354.353l.354-.353L3.5 7.5Zm1-1H5a.5.5 0 0 0-.854-.354L4.5 6.5Zm2 0V6H6v.5h.5Zm4 2H10V9h.5v-.5ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM3 11V6.5H2V11h1Zm-.854-4.146l1 1l.708-.708l-1-1l-.708.708Zm1.708 1l1-1l-.708-.708l-1 1l.708.708ZM4 6.5V11h1V6.5H4Zm2.5.5h1V6h-1v1Zm.5 4V8.5H6V11h1Zm0-2.5v-2H6v2h1Zm.5-.5h-1v1h1V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5H8ZM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6v1ZM10 6v2.5h1V6h-1Zm.5 3h2V8h-2v1ZM12 6v5h1V6h-1Z"/>`
	mpFourSolidInnerSVG                      = `<path fill="currentColor" d="M7 8h.5a.5.5 0 0 0 0-1H7v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM6 6h1.5a1.5 1.5 0 1 1 0 3H7v2H6V6Zm-3.691.038a.5.5 0 0 1 .545.108l.646.647l.646-.647A.5.5 0 0 1 5 6.5V11H4V7.707l-.5.5l-.5-.5V11H2V6.5a.5.5 0 0 1 .309-.462ZM11 6h-1v3h2v2h1V6h-1v2h-1V6Z" clip-rule="evenodd"/>`
	mpThreeOutlineInnerSVG                   = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-8 6l.354-.354A.5.5 0 0 0 2 6.5h.5Zm1 1l-.354.354l.354.353l.354-.353L3.5 7.5Zm1-1H5a.5.5 0 0 0-.854-.354L4.5 6.5Zm2 0V6H6v.5h.5Zm6 0l.4.3a.5.5 0 0 0-.4-.8v.5Zm-1.5 2l-.4-.3a.5.5 0 0 0 .4.8v-.5ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM3 11V6.5H2V11h1Zm-.854-4.146l1 1l.708-.708l-1-1l-.708.708Zm1.708 1l1-1l-.708-.708l-1 1l.708.708ZM4 6.5V11h1V6.5H4Zm2.5.5h1V6h-1v1Zm.5 4V8.5H6V11h1Zm0-2.5v-2H6v2h1Zm.5-.5h-1v1h1V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5H8ZM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6v1ZM10 7h2.5V6H10v1Zm2.1-.8l-1.5 2l.8.6l1.5-2l-.8-.6ZM11 9h.5V8H11v1Zm.5 1H10v1h1.5v-1Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 13 9.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 8v1Z"/>`
	mpThreeSolidInnerSVG                     = `<path fill="currentColor" d="M7.5 8H7V7h.5a.5.5 0 0 1 0 1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM7.5 6H6v5h1V9h.5a1.5 1.5 0 1 0 0-3Zm-4.646.146A.5.5 0 0 0 2 6.5V11h1V7.707l.5.5l.5-.5V11h1V6.5a.5.5 0 0 0-.854-.354l-.646.647l-.646-.647ZM11.5 7H10V6h2.5a.5.5 0 0 1 .4.8l-.951 1.268A1.5 1.5 0 0 1 11.5 11H10v-1h1.5a.5.5 0 0 0 0-1H11a.5.5 0 0 1-.4-.8l.9-1.2Z" clip-rule="evenodd"/>`
	msExcelOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M2.5 3.5v-2a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-2m0-6l4 4m-4 0l4-4m-5-2h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`
	msExcelSolidInnerSVG                     = `<path fill="currentColor" d="M3.793 7.5L2.146 5.854l.708-.708L4.5 6.793l1.646-1.647l.708.708L5.207 7.5l1.647 1.646l-.708.708L4.5 8.207L2.854 9.854l-.708-.708L3.793 7.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.5 0A1.5 1.5 0 0 0 2 1.5V3h-.5A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12H2v1.5A1.5 1.5 0 0 0 3.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0h-10Zm-2 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-6Z" clip-rule="evenodd"/>`
	msPowerpointOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M1.755 3.5a7 7 0 1 1 0 8M2.5 10V8.5m0 0v-3H5a1.5 1.5 0 1 1 0 3H2.5Zm-1-5h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`
	msPowerpointSolidInnerSVG                = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 5H2v5h1V9h2a2 2 0 1 0 0-4Zm0 3H3V6h2a1 1 0 0 1 0 2Z"/><path d="M7.5 0a7.49 7.49 0 0 0-6 3A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12a7.5 7.5 0 1 0 6-12ZM1 4.5a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.487v6.026a.5.5 0 0 1-.5.487h-6a.5.5 0 0 1-.5-.5v-6Z"/></g>`
	msWordOutlineInnerSVG                    = `<path fill="currentColor" d="m3.5 9.5l-.485.121a.5.5 0 0 0 .901.156L3.5 9.5Zm1-1.5l.416-.277a.5.5 0 0 0-.832 0L4.5 8Zm1 1.5l-.416.277a.5.5 0 0 0 .901-.156L5.5 9.5ZM1.5 4h6V3h-6v1Zm6.5.5v6h1v-6H8ZM7.5 11h-6v1h6v-1ZM1 10.5v-6H0v6h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 12v-1Zm6.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 10.5H8ZM7.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 3v1Zm-6-1A1.5 1.5 0 0 0 0 4.5h1a.5.5 0 0 1 .5-.5V3Zm.515 2.621l1 4l.97-.242l-1-4l-.97.242Zm1.901 4.156l1-1.5l-.832-.554l-1 1.5l.832.554Zm.168-1.5l1 1.5l.832-.554l-1-1.5l-.832.554Zm1.901 1.344l1-4l-.97-.242l-1 4l.97.242ZM3 3.5v-2H2v2h1ZM3.5 1h10V0h-10v1Zm10.5.5v12h1v-12h-1ZM13.5 14h-10v1h10v-1ZM3 13.5v-2H2v2h1Zm.5.5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 15v-1Zm10.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0v1ZM3 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 2 1.5h1Z"/>`
	msWordSolidInnerSVG                      = `<path fill="currentColor" d="m2.015 5.621l1 4a.5.5 0 0 0 .901.156l.584-.876l.584.876a.5.5 0 0 0 .901-.156l1-4l-.97-.242l-.726 2.903l-.373-.56a.5.5 0 0 0-.832 0l-.373.56l-.726-2.903l-.97.242Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.5 0A1.5 1.5 0 0 0 2 1.5V3h-.5A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12H2v1.5A1.5 1.5 0 0 0 3.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0h-10Zm-2 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-6Z" clip-rule="evenodd"/>`
	nSixtyFourOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M4.5 4v3M3 5.5h3m3 0h1m1 1h1m-7.5-4h-1a3 3 0 0 0-3 3v4.838a1.162 1.162 0 0 0 2.265.367L3.5 8.5h1.095a1 1 0 0 1 .995.9l.26 2.607a1.657 1.657 0 0 0 3.3 0L9.41 9.4a1 1 0 0 1 .995-.9H11.5l.735 2.205a1.162 1.162 0 0 0 2.265-.367V5.5a3 3 0 0 0-3-3h-1l-1-1h-4l-1 1Z"/>`
	nSixtyFourSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M5.293 1h4.414l1 1h.793A3.5 3.5 0 0 1 15 5.5v4.838a1.662 1.662 0 0 1-3.24.525L11.14 9h-.735a.5.5 0 0 0-.498.45l-.26 2.607a2.157 2.157 0 0 1-4.294 0l-.26-2.607A.5.5 0 0 0 4.595 9H3.86l-.62 1.863A1.662 1.662 0 0 1 0 10.338V5.5A3.5 3.5 0 0 1 3.5 2h.793l1-1ZM4 7V6H3V5h1V4h1v1h1v1H5v1H4Zm5-1h1V5H9v1Zm3 0v1h-1V6h1Z" clip-rule="evenodd"/>`
	nesOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M3.5 7v3M2 8.5h3m6 1h1m-3 0h1m-8.5-6h12a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`
	nesSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M0 4.5A1.5 1.5 0 0 1 1.5 3h12A1.5 1.5 0 0 1 15 4.5v6a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 10.5v-6ZM3 10V9H2V8h1V7h1v1h1v1H4v1H3Zm8 0h1V9h-1v1Zm-1 0H9V9h1v1Z" clip-rule="evenodd"/>`
	netlifyOutlineInnerSVG                   = `<path fill="currentColor" d="m.5 7.5l-.354-.354a.5.5 0 0 0 0 .708L.5 7.5Zm7-7l.354-.354a.5.5 0 0 0-.708 0L7.5.5Zm7 7l.354.354a.5.5 0 0 0 0-.708L14.5 7.5Zm-7 7l-.354.354a.5.5 0 0 0 .708 0L7.5 14.5ZM.854 7.854l7-7l-.708-.708l-7 7l.708.708Zm6.292-7l7 7l.708-.708l-7-7l-.708.708Zm7 6.292l-7 7l.708.708l7-7l-.708-.708Zm-6.292 7l-7-7l-.708.708l7 7l.708-.708ZM4.314 3.964l10 4l.372-.928l-10-4l-.372.928ZM8.58 1.728l-5.5 8.5l.84.544l5.5-8.5l-.84-.544ZM2.1 5.8l6 8l.8-.6l-6-8l-.8.6ZM.394 7.989l11.5 2.5l.212-.978l-11.5-2.5l-.212.978Zm2.32 1.963l9.5-4.5l-.428-.904l-9.5 4.5l.428.904Zm7.292-6.53l-1.5 9.5l.988.156l1.5-9.5l-.988-.156Z"/>`
	netlifySolidInnerSVG                     = `<path fill="currentColor" d="M7.146.146a.5.5 0 0 1 .708 0l.98.98l-1.82 2.94l-2.706-1.081L7.146.146ZM3.539 3.754L2.426 4.867L4.72 7.772L6.479 4.93l-2.94-1.176ZM1.714 5.579L.146 7.146a.502.502 0 0 0-.061.075l3.522.755l-1.893-2.397ZM.66 8.367l1.007 1.007l1.175-.54L.66 8.368Zm1.76 1.761l.52.52l.654-1.058l-1.173.538Zm1.246 1.245l3.48 3.48a.5.5 0 0 0 .708 0l.323-.323l.119-.615L4.819 9.51l-1.153 1.863Zm5.772 1.896l2.488-2.488l-1.93-.413l-.558 2.9Zm3.33-3.33l1.781-1.78l-3.833-1.534l-.531 2.76l2.583.554ZM14.92 7.23a.51.51 0 0 0-.066-.084L12.99 5.283l-1.368.628L14.92 7.23Zm-2.684-2.7l-.937-.938l-.288 1.499l1.225-.562Zm-1.791-1.792l-.885-.885l-1.604 2.59l2.007.804l.482-2.509Zm-4.447 5.75L9.662 6.81l-.455 2.367l-3.209-.688Zm.114 1.047l2.906.623l-.473 2.46l-2.433-3.083Zm.075-2.233l2.9-1.329l-1.665-.666l-1.235 1.995Z"/>`
	nextCircleOutlineInnerSVG                = `<path fill="currentColor" d="m4.5 5.5l.248-.434A.5.5 0 0 0 4 5.5h.5Zm0 4H4a.5.5 0 0 0 .748.434L4.5 9.5Zm3.5-2l.248.434a.5.5 0 0 0 0-.868L8 7.5ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM4 5.5v4h1v-4H4Zm.748 4.434l3.5-2l-.496-.868l-3.5 2l.496.868Zm3.5-2.868l-3.5-2l-.496.868l3.5 2l.496-.868ZM10 5v5h1V5h-1Z"/>`
	nextCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM10 5h1v5h-1V5Zm-5.252.066A.5.5 0 0 0 4 5.5v4a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868l-3.5-2Z" clip-rule="evenodd"/>`
	nextOutlineInnerSVG                      = `<path fill="currentColor" d="m1.5 2.5l.29-.407A.5.5 0 0 0 1 2.5h.5Zm0 10H1a.5.5 0 0 0 .79.407L1.5 12.5Zm7-5l.29.407a.5.5 0 0 0 0-.814L8.5 7.5ZM1 2.5v10h1v-10H1Zm.79 10.407l7-5l-.58-.814l-7 5l.58.814Zm7-5.814l-7-5l-.58.814l7 5l.58-.814ZM13 2v11h1V2h-1Z"/>`
	nextSmallOutlineInnerSVG                 = `<path fill="currentColor" d="m4.5 5.5l.248-.434A.5.5 0 0 0 4 5.5h.5Zm0 4H4a.5.5 0 0 0 .748.434L4.5 9.5Zm3.5-2l.248.434a.5.5 0 0 0 0-.868L8 7.5Zm-4-2v4h1v-4H4Zm.748 4.434l3.5-2l-.496-.868l-3.5 2l.496.868Zm3.5-2.868l-3.5-2l-.496.868l3.5 2l.496-.868ZM10 5v5h1V5h-1Z"/>`
	nextSmallSolidInnerSVG                   = `<path fill="currentColor" d="M4.748 5.066A.5.5 0 0 0 4 5.5v4a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868l-3.5-2ZM10 10h1V5h-1v5Z"/>`
	nextSolidInnerSVG                        = `<path fill="currentColor" d="M1.79 2.093A.5.5 0 0 0 1 2.5v10a.5.5 0 0 0 .79.407l7-5a.5.5 0 0 0 0-.814l-7-5ZM13 13h1V2h-1v11Z"/>`
	nextjsOutlineInnerSVG                    = `<path fill="currentColor" d="m4.5 4.5l.405-.293A.5.5 0 0 0 4 4.5h.5Zm3 9.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM5 12V4.5H4V12h1Zm-.905-7.207l6.5 9l.81-.586l-6.5-9l-.81.586ZM10 4v6h1V4h-1Z"/>`
	nextjsSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 11.697 6.216L4.907 4.21A.5.5 0 0 0 4 4.5V12h1V6.06l5.83 8.162A7.5 7.5 0 0 1 0 7.5ZM10 10V4h1v6h-1Z" clip-rule="evenodd"/>`
	ngcOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M3 3.5a11.08 11.08 0 0 1 9 0M.5 6.5v5.764a1.236 1.236 0 0 0 2.342.553L3.5 11.5m11-5v5.764a1.236 1.236 0 0 1-2.342.553L11.5 11.5M6 7.5h3m-5.7 3.499L1.148 8.31A2.961 2.961 0 0 1 .5 6.461v-.396a2.565 2.565 0 0 1 5.032-.705l1.117 3.91a1.922 1.922 0 0 1-3.35 1.729Zm8.4 0l2.151-2.688a2.96 2.96 0 0 0 .649-1.85v-.396a2.565 2.565 0 0 0-5.032-.705L8.351 9.27a1.922 1.922 0 0 0 3.35 1.729Z"/>`
	ngcSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M12.117 3.005a11.58 11.58 0 0 0-9.234 0A3.065 3.065 0 0 0 0 6.065v6.199a1.736 1.736 0 0 0 3.289.776l.516-1.033A2.422 2.422 0 0 0 7.13 9.133L6.806 8h1.388l-.323 1.133a2.422 2.422 0 0 0 3.324 2.874l.516 1.033A1.736 1.736 0 0 0 15 12.264V6.065a3.065 3.065 0 0 0-2.883-3.06Zm-7.473.433c.65.39 1.15 1.018 1.368 1.785L6.52 7h1.96l.508-1.777a3.063 3.063 0 0 1 1.368-1.785a10.582 10.582 0 0 0-5.712 0ZM14 8.925l-1.909 2.386a2.416 2.416 0 0 1-.08.095l.595 1.187a.736.736 0 0 0 1.394-.33V8.926Zm-11.012 2.48a2.516 2.516 0 0 1-.08-.094L1 8.925v3.339a.736.736 0 0 0 1.394.33l.594-1.188Z" clip-rule="evenodd"/>`
	nintendoSwitchOutlineInnerSVG            = `<path fill="currentColor" d="M8.5 14.5H8a.5.5 0 0 0 .5.5v-.5Zm0-12V2a.5.5 0 0 0-.5.5h.5Zm6 9h.5h-.5Zm-3 3V14v.5Zm3-9H14h.5Zm-3-3V3v-.5Zm-5-2H7a.5.5 0 0 0-.5-.5v.5Zm0 12v.5a.5.5 0 0 0 .5-.5h-.5Zm-6-9H1H.5Zm3-3V1V.5Zm0 12v.5v-.5Zm-3-3H0h.5Zm13.5-4v6h1v-6h-1ZM11.5 14h-3v1h3v-1Zm-2.5.5v-12H8v12h1ZM8.5 3h3V2h-3v1Zm5.5 8.5a2.5 2.5 0 0 1-2.5 2.5v1a3.5 3.5 0 0 0 3.5-3.5h-1Zm1-6A3.5 3.5 0 0 0 11.5 2v1A2.5 2.5 0 0 1 14 5.5h1Zm-9-5v12h1V.5H6ZM6.5 12h-3v1h3v-1ZM1 9.5v-6H0v6h1ZM3.5 1h3V0h-3v1ZM1 3.5A2.5 2.5 0 0 1 3.5 1V0A3.5 3.5 0 0 0 0 3.5h1ZM3.5 12A2.5 2.5 0 0 1 1 9.5H0A3.5 3.5 0 0 0 3.5 13v-1Zm9.5-1.5A1.5 1.5 0 0 0 11.5 9v1a.5.5 0 0 1 .5.5h1ZM11.5 12a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM10 10.5a1.5 1.5 0 0 0 1.5 1.5v-1a.5.5 0 0 1-.5-.5h-1Zm1 0a.5.5 0 0 1 .5-.5V9a1.5 1.5 0 0 0-1.5 1.5h1ZM3.5 5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 6V5Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 4.5H4ZM3.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 3v1Zm0-1A1.5 1.5 0 0 0 2 4.5h1a.5.5 0 0 1 .5-.5V3Z"/>`
	nintendoSwitchSolidInnerSVG              = `<path fill="currentColor" d="M3.5 4a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M7 .5v12a.5.5 0 0 1-.5.5h-3A3.5 3.5 0 0 1 0 9.5v-6A3.5 3.5 0 0 1 3.5 0h3a.5.5 0 0 1 .5.5Zm-5 4a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/><path fill="currentColor" d="M11.5 10a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M15 5.5v6a3.5 3.5 0 0 1-3.5 3.5h-3a.5.5 0 0 1-.5-.5v-12a.5.5 0 0 1 .5-.5h3A3.5 3.5 0 0 1 15 5.5Zm-5 5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/>`
	nodejsOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M11.5 6v-.167c0-.736-.597-1.333-1.333-1.333H9a1.5 1.5 0 1 0 0 3h1a1.5 1.5 0 0 1 0 3H9A1.5 1.5 0 0 1 7.5 9m-2-5v5.264a2 2 0 0 1-1.106 1.789L3.5 11.5m-2-1v-6l6-3.5l6 3.5v6l-6 3.5l-6-3.5Z"/>`
	nodejsSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M14 4.213L7.5.42L1 4.213v6.574l1.006.587l2.057-.832A1.5 1.5 0 0 0 5 9.152V4h1v5.152a2.5 2.5 0 0 1-1.562 2.317l-1.34.542L7.5 14.58l6.5-3.792V4.213ZM7 6a2 2 0 0 1 2-2h1.167C11.179 4 12 4.82 12 5.833V6h-1v-.167A.833.833 0 0 0 10.167 5H9a1 1 0 0 0 0 2h1a2 2 0 1 1 0 4H9a2 2 0 0 1-2-2h1a1 1 0 0 0 1 1h1a1 1 0 1 0 0-2H9a2 2 0 0 1-2-2Z" clip-rule="evenodd"/>`
	noteOutlineInnerSVG                      = `<path fill="currentColor" d="M10.5 14.5H10a.5.5 0 0 0 .854.354L10.5 14.5Zm0-4V10a.5.5 0 0 0-.5.5h.5Zm4 0l.354.354A.5.5 0 0 0 14.5 10v.5ZM1.5 1h12V0h-12v1ZM1 13.5v-12H0v12h1Zm13-12v8.586h1V1.5h-1ZM10.086 14H1.5v1h8.586v-1Zm3.768-3.56l-3.415 3.414l.707.707l3.415-3.415l-.707-.707ZM10.086 15a1.5 1.5 0 0 0 1.06-.44l-.707-.706a.5.5 0 0 1-.353.146v1ZM14 10.086a.5.5 0 0 1-.146.353l.707.707a1.5 1.5 0 0 0 .439-1.06h-1ZM0 13.5A1.5 1.5 0 0 0 1.5 15v-1a.5.5 0 0 1-.5-.5H0ZM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0v1Zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5V0ZM11 14.5v-4h-1v4h1Zm-.5-3.5h4v-1h-4v1Zm3.646-.854l-4 4l.708.708l4-4l-.708-.708ZM3 4h9V3H3v1Z"/>`
	noteSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5V9H9.5a.5.5 0 0 0-.5.5V15H1.5A1.5 1.5 0 0 1 0 13.5v-12ZM12 4H3V3h9v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M10 15h.086a1.5 1.5 0 0 0 1.06-.44l3.415-3.414a1.5 1.5 0 0 0 .439-1.06V10h-5v5Z"/>`
	npmOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M4.5 10.5v2h2v-2h8v-6H.5v6h4Zm0 0v-6m4 0v6M6.5 6v3m-4-3v4.5m8-4.5v4.5m2-4.5v4.5"/>`
	npmSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M0 4h15v7H7v2H4v-2H0V4Zm4 6V5H1v5h1V6h1v4h1Zm1-5v7h1v-2h2V5H5Zm4 0v5h1V6h1v4h1V6h1v4h1V5H9ZM6 9V6h1v3H6Z" clip-rule="evenodd"/>`
	nuxtjsOutlineInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="m.5 12.5l6-10l6 10H.5Z"/><path d="m4.5 12.5l5-8.5l5 8.5h-10Z"/></g>`
	nuxtjsSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M6.5 2a.5.5 0 0 1 .429.243l1.527 2.545l.613-1.042a.5.5 0 0 1 .862 0l5 8.5A.5.5 0 0 1 14.5 13H.5a.5.5 0 0 1-.429-.757l6-10A.5.5 0 0 1 6.5 2ZM5.374 12h6.243L8.465 6.746L5.375 12ZM7.88 5.77L4.214 12h-2.83L6.5 3.472L7.879 5.77Zm1.163-.005L12.783 12h.843L9.5 4.986l-.458.779Z" clip-rule="evenodd"/>`
	omegaOutlineInnerSVG                     = `<path fill="currentColor" d="M9.333 11.687a.5.5 0 1 0 .334.943l-.334-.943Zm-4 .943a.5.5 0 1 0 .334-.943l-.334.943ZM5.5 14.5v.5H6v-.5h-.5Zm4 0H9v.5h.5v-.5ZM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0v1Zm0-1A6.5 6.5 0 0 0 1 6.5h1A5.5 5.5 0 0 1 7.5 1V0ZM13 6.5a5.503 5.503 0 0 1-3.667 5.187l.334.943A6.503 6.503 0 0 0 14 6.5h-1Zm-7.333 5.187A5.503 5.503 0 0 1 2 6.5H1c0 2.83 1.81 5.238 4.333 6.13l.334-.943ZM0 15h5.5v-1H0v1Zm6-.5V12H5v2.5h1Zm9-.5H9.5v1H15v-1Zm-5 .5V12H9v2.5h1Z"/>`
	omegaSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 1 1 9 6.002V14h5v1H9v-3h.026a.499.499 0 0 1 .307-.313a5.5 5.5 0 1 0-3.667 0a.496.496 0 0 1 .308.313H6v3H0v-1h5v-1.498A6.502 6.502 0 0 1 1 6.5Z" clip-rule="evenodd"/>`
	operaOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M7.5 12.5a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3m0 10a3 3 0 0 0 3-3v-4a3 3 0 0 0-3-3m0 10h.585c1.907 0 3.78-.518 5.415-1.5m-6-8.5h.585c1.907 0 3.78.518 5.415 1.5m-6 10.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	operaSolidInnerSVG                       = `<path fill="currentColor" d="M0 7.5a7.5 7.5 0 0 1 13.146-4.936A17.5 17.5 0 0 0 8.741 2H7.5A3.5 3.5 0 0 0 4 5.5v4A3.5 3.5 0 0 0 7.5 13h1.241c1.488 0 2.969-.19 4.405-.563A7.5 7.5 0 0 1 0 7.5Z"/><path fill="currentColor" d="M14.073 11.115A7.466 7.466 0 0 0 15 7.5c0-1.31-.336-2.543-.927-3.615l-.114-.038a16.5 16.5 0 0 0-3.962-.8A3.489 3.489 0 0 1 11 5.5v4a3.49 3.49 0 0 1-1.003 2.452a16.499 16.499 0 0 0 3.962-.799l.114-.038Z"/>`
	otpOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M6 5.5h3m-1.5 0V10m3 0V7.5m0 0v-2h1a1 1 0 1 1 0 2h-1Zm-6-1v2a1 1 0 0 1-2 0v-2a1 1 0 0 1 2 0Zm-3-6h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`
	otpSolidInnerSVG                         = `<path fill="currentColor" d="M3.5 6a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.5-.5ZM11 7h.5a.5.5 0 0 0 0-1H11v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-12Zm2 5a1.5 1.5 0 1 1 3 0v2a1.5 1.5 0 1 1-3 0v-2ZM7 6H6V5h3v1H8v4H7V6Zm3-1h1.5a1.5 1.5 0 0 1 0 3H11v2h-1V5Z" clip-rule="evenodd"/>`
	pageBreakOutlineInnerSVG                 = `<path fill="currentColor" d="M.5 8.993H0v1h.5v-1Zm2 1H3v-1h-.5v1Zm2-1H4v1h.5v-1Zm2 1H7v-1h-.5v1Zm2-1H8v1h.5v-1Zm2 1h.5v-1h-.5v1Zm2-1H12v1h.5v-1Zm2 1h.5v-1h-.5v1ZM10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354ZM.5 9.993h2v-1h-2v1Zm4 0h2v-1h-2v1Zm4 0h2v-1h-2v1Zm4 0h2v-1h-2v1Zm0 4.007h-10v1h10v-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM13 3.5V8h1V3.5h-1Zm0 8.25v1.75h1v-1.75h-1ZM2 8V1.5H1V8h1Zm0 5.5V11H1v2.5h1Zm.5.5a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2.5 0A1.5 1.5 0 0 0 1 1.5h1a.5.5 0 0 1 .5-.5V0Z"/>`
	pageBreakSolidInnerSVG                   = `<path fill="currentColor" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V8H1V1.5ZM1 11h13v2.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V11ZM0 8.993h3v1H0v-1Zm4 0h3v1H4v-1Zm7 0H8v1h3v-1Zm1 0h3v1h-3v-1Z"/>`
	pageNumberOutlineInnerSVG                = `<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354Zm-4 9l-.354-.354A.5.5 0 0 0 9.5 13v-.5Zm3 1.5h-10v1h10v-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2.5 0A1.5 1.5 0 0 0 1 1.5h1a.5.5 0 0 1 .5-.5V0ZM12 12H9.5v1H12v-1Zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792l.708.708ZM10.793 9H10.5v1h.293V9ZM10.5 9A1.5 1.5 0 0 0 9 10.5h1a.5.5 0 0 1 .5-.5V9Zm1.5 1.207C12 9.54 11.46 9 10.793 9v1c.114 0 .207.093.207.207h1Zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147l.706.707ZM13 3.5v10h1v-10h-1Zm-11 10v-12H1v12h1Z"/>`
	pageNumberSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM10.5 9A1.5 1.5 0 0 0 9 10.5h1a.5.5 0 0 1 .5-.5h.293a.207.207 0 0 1 .146.354l-1.793 1.792A.5.5 0 0 0 9.5 13H12v-1h-1.293l.94-.94A1.207 1.207 0 0 0 10.793 9H10.5Z" clip-rule="evenodd"/>`
	paintbrushOutlineInnerSVG                = `<path fill="currentColor" d="m.5 14l-.481-.136a.5.5 0 0 0 .758.552L.5 14Zm.971-3.422l.481.136l-.48-.136Zm-.46 3.081l-.277-.416l.277.416Zm2.062-.148l-.215.451l.215-.451ZM5 6l-.25-.433a.5.5 0 0 0-.104.787L5 6Zm4 4l-.354.354a.5.5 0 0 0 .787-.103L9 10ZM14.5.5l.433.25a.5.5 0 0 0-.684-.683L14.5.5ZM.981 14.136l.971-3.422l-.962-.273l-.971 3.422l.962.274Zm2.956-4.922H4v-1h-.063v1Zm2.063 2v.053h1v-.053H6Zm-1.947 2h-.08v1h.08v-1Zm-3.32.03l-.51.34l.554.832l.512-.34l-.555-.833Zm2.555-.185a2.594 2.594 0 0 0-2.554.184l.555.832a1.594 1.594 0 0 1 1.569-.113l.43-.903Zm.685.155c-.237 0-.471-.053-.685-.155l-.43.903c.348.166.73.252 1.115.252v-1ZM6 11.267a1.947 1.947 0 0 1-1.947 1.947v1A2.947 2.947 0 0 0 7 11.267H6ZM4 9.214a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3v1Zm-2.048 1.5a2.063 2.063 0 0 1 1.985-1.5v-1c-1.37 0-2.573.91-2.947 2.227l.962.273Zm2.694-4.36l4 4l.708-.708l-4-4l-.708.708ZM14.25.067l-9.5 5.5l.5.866l9.5-5.5l-.5-.866ZM9.432 10.251l5.5-9.5l-.866-.502l-5.5 9.5l.866.502ZM7.146 4.854l3 3l.708-.708l-3-3l-.708.708Z"/>`
	paintbrushSolidInnerSVG                  = `<path fill="currentColor" d="M14.854.146a.5.5 0 0 1 .079.605l-3.841 6.634l-3.477-3.477L14.25.068a.5.5 0 0 1 .605.078ZM6.72 4.427l-1.97 1.14a.5.5 0 0 0-.104.787l4 4a.5.5 0 0 0 .787-.103l1.14-1.97L6.72 4.426ZM.99 10.441a3.063 3.063 0 0 1 2.947-2.227H4a3 3 0 0 1 3 3v.053a2.947 2.947 0 0 1-2.947 2.947h-.08a2.59 2.59 0 0 1-1.115-.252a1.594 1.594 0 0 0-1.57.113l-.51.341a.5.5 0 0 1-.759-.553l.971-3.422Z"/>`
	paintbucketOutlineInnerSVG               = `<path fill="currentColor" d="m6.086 2.914l-.354-.353l.354.353ZM1.914 7.086l-.353-.354l.353.354Zm0 2.828l.354-.353l-.354.353Zm3.172 3.172l.353-.354l-.353.354Zm2.828 0l.354.353l-.354-.353Zm4.172-4.172l-.354-.353l.354.353Zm0-2.828l.353-.354l-.353.354ZM8.914 2.914l-.353.354l.353-.354Zm3.767 7.586l.353-.354l-.353-.353l-.354.353l.354.354Zm1.18 1.18l-.353.354l.354-.353ZM11.639 14l.312-.39l-.312.39Zm2.086 0l-.312-.39l.312.39ZM5.732 2.56L1.561 6.733l.707.707l4.171-4.171l-.707-.707Zm-4.171 7.708l3.171 3.171l.707-.707l-3.171-3.171l-.707.707Zm6.707 3.171l4.171-4.171l-.707-.707l-4.171 4.171l.707.707Zm4.171-7.707L9.268 2.561l-.707.707l3.171 3.171l.707-.707Zm0 3.536a2.5 2.5 0 0 0 0-3.536l-.707.707a1.5 1.5 0 0 1 0 2.122l.707.707Zm-7.707 4.171a2.5 2.5 0 0 0 3.536 0l-.707-.707a1.5 1.5 0 0 1-2.122 0l-.707.707ZM1.561 6.732a2.5 2.5 0 0 0 0 3.536l.707-.707a1.5 1.5 0 0 1 0-2.122l-.707-.707Zm4.878-3.464a1.5 1.5 0 0 1 2.122 0l.707-.707a2.5 2.5 0 0 0-3.536 0l.707.707ZM7 7V2.5H6V7h1ZM2 2.5v4h1v-4H2ZM4.5 0A2.5 2.5 0 0 0 2 2.5h1A1.5 1.5 0 0 1 4.5 1V0ZM7 2.5A2.5 2.5 0 0 0 4.5 0v1A1.5 1.5 0 0 1 6 2.5h1ZM1.5 10h10V9h-10v1Zm10.827.854l1.181 1.18l.707-.707l-1.18-1.18l-.708.707Zm-.473 1.18l1.18-1.18l-.707-.708l-1.18 1.181l.707.707Zm.096 1.576c-.29-.232-.422-.51-.438-.77c-.015-.257.081-.545.342-.806l-.707-.707c-.444.444-.667 1.004-.633 1.575c.035.569.324 1.099.811 1.488l.625-.78Zm1.462 0a1.17 1.17 0 0 1-1.462 0l-.625.78a2.17 2.17 0 0 0 2.711 0l-.624-.78Zm.096-1.576c.26.26.357.549.341.807c-.016.259-.147.537-.437.769l.624.78c.487-.39.777-.919.811-1.489c.035-.57-.188-1.13-.632-1.574l-.707.707Z"/>`
	paintbucketSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 0A2.5 2.5 0 0 0 2 2.5v3.793l-.44.44a2.5 2.5 0 0 0 0 3.535l3.172 3.171a2.5 2.5 0 0 0 3.536 0l4.171-4.171a2.5 2.5 0 0 0 0-3.536L9.268 2.561a2.498 2.498 0 0 0-2.342-.666A2.501 2.501 0 0 0 4.5 0ZM6 3.707V7h1V2.914a1.5 1.5 0 0 1 1.56.354l3.172 3.171a1.5 1.5 0 0 1 0 2.122l-.44.439H1.915a1.5 1.5 0 0 1 .354-1.56L6 3.706Zm-.009-1.372A1.5 1.5 0 0 0 3 2.5v2.793L5.732 2.56c.082-.083.169-.158.259-.226Z" clip-rule="evenodd"/><path fill="currentColor" d="m12.645 9.737l1.534 1.534a2.17 2.17 0 1 1-3.069 0l1.535-1.534Z"/>`
	paragraphOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M13 1.5H6.5a4 4 0 1 0 0 8h1m3 4.5V1.5M7.5 14V1.5"/>`
	paragraphSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M2 5.5A4.5 4.5 0 0 1 6.5 1H13v1h-2v12h-1V2H8v12H7v-4h-.5A4.5 4.5 0 0 1 2 5.5ZM7 9V2h-.5a3.5 3.5 0 1 0 0 7H7Z" clip-rule="evenodd"/>`
	passwordOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M12.5 8.5v-1a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-1m0-4h-4a2 2 0 1 0 0 4h4m0-4a2 2 0 1 1 0 4m-9-6v-3a3 3 0 0 1 6 0v3m2.5 4h1m-3 0h1m-3 0h1"/>`
	passwordSolidInnerSVG                    = `<path fill="currentColor" d="M11 11h-1v-1h1v1Zm-3 0h1v-1H8v1Zm5 0h-1v-1h1v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M3 6V3.5a3.5 3.5 0 1 1 7 0V6h1.5A1.5 1.5 0 0 1 13 7.5v.55a2.5 2.5 0 0 1 0 4.9v.55a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 0 13.5v-6A1.5 1.5 0 0 1 1.5 6H3Zm1-2.5a2.5 2.5 0 0 1 5 0V6H4V3.5ZM8.5 9a1.5 1.5 0 1 0 0 3h4a1.5 1.5 0 0 0 0-3h-4Z" clip-rule="evenodd"/>`
	patreonOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M2.5.5h-2v14h2V.5Zm2 5a5 5 0 1 0 10 0a5 5 0 0 0-10 0Z"/>`
	patreonSolidInnerSVG                     = `<path fill="currentColor" d="M3 0H0v15h3V0Zm6.5 0a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11Z"/>`
	pauseCircleOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="M6.5 5v5m2-5v5m-1 4.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	pauseCircleSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM7 10H6V5h1v5Zm2 0H8V5h1v5Z" clip-rule="evenodd"/>`
	pauseOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M5.5 3v9m4-9v9"/>`
	pauseSmallOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M6.5 5v5m2-5v5"/>`
	pauseSmallSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M6 10V5h1v5H6Zm2 0V5h1v5H8Z" clip-rule="evenodd"/>`
	pauseSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M5 12V3h1v9H5Zm4 0V3h1v9H9Z" clip-rule="evenodd"/>`
	pawOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M6.5 3V2a1.5 1.5 0 1 0-3 0v1a1.5 1.5 0 1 0 3 0Zm5 0V2a1.5 1.5 0 0 0-3 0v1a1.5 1.5 0 1 0 3 0Zm3 4.5V7a1.5 1.5 0 0 0-3 0v.5a1.5 1.5 0 0 0 3 0Zm-11 0V7a1.5 1.5 0 1 0-3 0v.5a1.5 1.5 0 1 0 3 0Zm-.645 4.14l2.918-3.543a2.237 2.237 0 0 1 3.454 0l2.918 3.543c.939 1.14.128 2.86-1.35 2.86c-.194 0-.385-.045-.559-.132l-.36-.18a5.315 5.315 0 0 0-4.753 0l-.36.18a1.248 1.248 0 0 1-.558.132c-1.478 0-2.289-1.72-1.35-2.86Z"/>`
	pawSolidInnerSVG                         = `<path fill="currentColor" d="M5 0a2 2 0 0 0-2 2v1a2 2 0 1 0 4 0V2a2 2 0 0 0-2-2Zm5 0a2 2 0 0 0-2 2v1a2 2 0 1 0 4 0V2a2 2 0 0 0-2-2ZM2 5a2 2 0 0 0-2 2v.5a2 2 0 1 0 4 0V7a2 2 0 0 0-2-2Zm11 0a2 2 0 0 0-2 2v.5a2 2 0 1 0 4 0V7a2 2 0 0 0-2-2ZM9.613 7.779a2.737 2.737 0 0 0-4.226 0L2.47 11.322C1.261 12.789 2.305 15 4.205 15c.272 0 .54-.063.782-.185l.36-.18a4.814 4.814 0 0 1 4.306 0l.36.18c.242.122.51.185.782.185c1.9 0 2.944-2.211 1.736-3.678L9.613 7.779Z"/>`
	pawsOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-miterlimit="10" d="m1.425 2.118l.307.647M.75 4.647l.307.647m5.277-3.176l.307.706M3.88 1l.306.647m6.934 5.118l-.306.647m-2.148.47l-.307.706m5.891 1.824l-.307.647m-.368-3.177l-.307.647M3.941 4.235c.43-.176.92-.176 1.35.06l1.657.823c.552.294.736 1 .307 1.47c-.185.236-.491.353-.798.353H5.72c-.306 0-.613.118-.797.353l-.43.53c-.184.235-.49.352-.798.352c-.613 0-1.104-.588-.981-1.176l.368-1.765c.123-.411.43-.823.859-1ZM11.059 10c.43.177.737.588.86 1.059l.367 1.764A.976.976 0 0 1 11.305 14c-.307 0-.614-.118-.798-.353l-.43-.53c-.184-.235-.49-.352-.797-.352h-.737c-.307 0-.613-.118-.798-.353a.978.978 0 0 1 .307-1.47l1.657-.824a1.25 1.25 0 0 1 1.35-.118Z"/>`
	pawsSolidInnerSVG                        = `<path fill="currentColor" d="M4.852 1.885L4.117.334l-.903.428l.735 1.551l.903-.428ZM2.398 3.002l-.735-1.55l-.904.428l.735 1.55l.904-.428Zm4.901.081L6.593 1.46l-.917.398l.706 1.623l.917-.398Zm-3.549.69a2.057 2.057 0 0 1 1.772.077l1.66.826c.786.419 1.124 1.484.454 2.236c-.3.373-.764.53-1.18.53H5.72c-.192 0-.336.074-.403.16l-.006.007l-.427.527c-.301.381-.769.54-1.189.54c-.917 0-1.661-.866-1.47-1.778l.367-1.765a.5.5 0 0 1 .01-.04c.157-.524.55-1.074 1.149-1.32ZM1.723 5.532L.988 3.98l-.904.428l.735 1.551l.904-.428Zm9.328 2.546l.735-1.55l-.903-.43l-.735 1.551l.903.429Zm2.455 1.117l.735-1.55l-.904-.429l-.735 1.551l.904.428Zm-4.888.051l.706-1.623l-.917-.399L7.7 8.847l.917.4Zm.825.445a1.751 1.751 0 0 1 1.832-.143c.591.254.976.806 1.127 1.385a.4.4 0 0 1 .006.024l.368 1.764c.186.89-.47 1.779-1.471 1.779c-.42 0-.888-.16-1.189-.54l-.433-.534c-.067-.086-.21-.161-.403-.161h-.737c-.42 0-.89-.16-1.191-.545a1.479 1.479 0 0 1 .465-2.22l.013-.007l1.613-.802Zm4.737 2.034l.736-1.55l-.904-.43l-.735 1.551l.904.429Z"/>`
	paypalOutlineInnerSVG                    = `<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M.5 12.5h4l1-4h1.795a4.625 4.625 0 0 0 4.33-3.001C12.532 3.08 10.745.5 8.161.5H3.5l-3 12Z"/><path d="M4 14.5h4L9 11h1.577c1.477 0 2.82-.859 3.438-2.2c.927-2.008-.54-4.3-2.75-4.3H6.5L4 14.5Z"/></g>`
	paypalSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M3.015.379A.5.5 0 0 1 3.5 0h4.661c2.397 0 4.191 1.957 4.204 4.172c1.928.626 3.021 2.851 2.104 4.837a4.287 4.287 0 0 1-3.892 2.491h-1.2l-.896 3.137A.5.5 0 0 1 8 15H4a.5.5 0 0 1-.485-.621L3.86 13H.5a.5.5 0 0 1-.485-.621l3-12ZM8.16 1c1.762 0 3.097 1.388 3.197 3.001L11.264 4H6.5a.5.5 0 0 0-.485.379L4.11 12H1.14L3.89 1h4.271Zm-.866 8H5.89l-1.25 5h2.983l.896-3.137A.5.5 0 0 1 9 10.5h1.577a3.287 3.287 0 0 0 2.985-1.91c.626-1.357-.057-2.87-1.32-3.396c-.039.16-.089.32-.149.48A5.125 5.125 0 0 1 7.296 9Z" clip-rule="evenodd"/>`
	pdfOutlineInnerSVG                       = `<path fill="currentColor" d="M2.5 6.5V6H2v.5h.5Zm4 0V6H6v.5h.5Zm0 4H6v.5h.5v-.5Zm7-7h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5ZM2.5 7h1V6h-1v1Zm.5 4V8.5H2V11h1Zm0-2.5v-2H2v2h1Zm.5-.5h-1v1h1V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 7.5H4ZM3.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 6v1ZM6 6.5v4h1v-4H6Zm.5 4.5h1v-1h-1v1ZM9 9.5v-2H8v2h1ZM7.5 6h-1v1h1V6ZM9 7.5A1.5 1.5 0 0 0 7.5 6v1a.5.5 0 0 1 .5.5h1ZM7.5 11A1.5 1.5 0 0 0 9 9.5H8a.5.5 0 0 1-.5.5v1ZM10 6v5h1V6h-1Zm.5 1H13V6h-2.5v1Zm0 2H12V8h-1.5v1ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1Z"/>`
	pdfSolidInnerSVG                         = `<path fill="currentColor" d="M3.5 8H3V7h.5a.5.5 0 0 1 0 1ZM7 10V7h.5a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H7Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM3.5 6H2v5h1V9h.5a1.5 1.5 0 1 0 0-3Zm4 0H6v5h1.5A1.5 1.5 0 0 0 9 9.5v-2A1.5 1.5 0 0 0 7.5 6Zm2.5 5V6h3v1h-2v1h1v1h-1v2h-1Z" clip-rule="evenodd"/>`
	penOutlineInnerSVG                       = `<path fill="currentColor" d="M2.5.5V0H2v.5h.5Zm10 0h.5V0h-.5v.5ZM4.947 4.724a.5.5 0 0 0-.894-.448l.894.448ZM2.5 8.494l-.447-.223l-.146.293l.21.251l.383-.32Zm5 5.997l-.384.32a.5.5 0 0 0 .769 0l-.385-.32Zm5-5.996l.384.32l.21-.251l-.146-.293l-.447.224Zm-1.553-4.219a.5.5 0 0 0-.894.448l.894-.448ZM8 9.494v-.5H7v.5h1Zm-.5-4.497A4.498 4.498 0 0 1 3 .5H2a5.498 5.498 0 0 0 5.5 5.497v-1ZM2.5 1h10V0h-10v1ZM12 .5a4.498 4.498 0 0 1-4.5 4.497v1c3.038 0 5.5-2.46 5.5-5.497h-1ZM4.053 4.276l-2 3.995l.895.448l2-3.995l-.895-.448ZM2.116 8.815l5 5.996l.769-.64l-5-5.996l-.769.64Zm5.768 5.996l5-5.996l-.768-.64l-5 5.996l.769.64Zm5.064-6.54l-2-3.995l-.895.448l2 3.995l.895-.448ZM8 14.49V9.494H7v4.997h1Z"/>`
	penSolidInnerSVG                         = `<path fill="currentColor" d="M2 0h11v.5a5.482 5.482 0 0 1-1.874 4.134l1.968 3.93L8 14.672V8.994H7v5.678L1.907 8.564l1.967-3.93A5.482 5.482 0 0 1 2 .5V0Z"/>`
	phoneOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M4.72.5H2.5a2 2 0 0 0-2 2v2c0 5.523 4.477 10 10 10h2a2 2 0 0 0 2-2v-1.382a1 1 0 0 0-.553-.894l-2.416-1.208a1 1 0 0 0-1.396.578l-.297.893a1.21 1.21 0 0 1-1.385.804A6.047 6.047 0 0 1 3.71 6.547a1.21 1.21 0 0 1 .804-1.385l1.108-.37a1 1 0 0 0 .654-1.19L5.69 1.257A1 1 0 0 0 4.72.5Z"/>`
	phoneSolidInnerSVG                       = `<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 0 2.5v2C0 10.299 4.701 15 10.5 15h2a2.5 2.5 0 0 0 2.5-2.5v-1.382a1.5 1.5 0 0 0-.83-1.342l-2.415-1.208a1.5 1.5 0 0 0-2.094.868l-.298.893a.71.71 0 0 1-.812.471A5.547 5.547 0 0 1 4.2 6.45a.71.71 0 0 1 .471-.812l1.109-.37a1.5 1.5 0 0 0 .98-1.787l-.586-2.344A1.5 1.5 0 0 0 4.72 0H2.5Z"/>`
	phonecallBlockedOutlineInnerSVG          = `<path fill="none" stroke="currentColor" d="m13.5 1.5l-4 4m-7-5h2.22a1 1 0 0 1 .97.757l.585 2.345a1 1 0 0 1-.654 1.19l-1.108.37a1.21 1.21 0 0 0-.804 1.385a6.047 6.047 0 0 0 4.744 4.744a1.21 1.21 0 0 0 1.385-.804l.297-.893a1 1 0 0 1 1.396-.578l2.416 1.208a1 1 0 0 1 .553.894V12.5a2 2 0 0 1-2 2h-2c-5.523 0-10-4.477-10-10v-2a2 2 0 0 1 2-2Zm9 6a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`
	phonecallBlockedSolidInnerSVG            = `<path fill="currentColor" d="M0 2.5A2.5 2.5 0 0 1 2.5 0h2.22a1.5 1.5 0 0 1 1.454 1.136L6.76 3.48a1.5 1.5 0 0 1-.98 1.787l-1.109.37a.71.71 0 0 0-.471.812A5.547 5.547 0 0 0 8.55 10.8a.71.71 0 0 0 .812-.471l.298-.893a1.5 1.5 0 0 1 2.094-.868l2.416 1.208a1.5 1.5 0 0 1 .83 1.342V12.5a2.5 2.5 0 0 1-2.5 2.5h-2C4.701 15 0 10.299 0 4.5v-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M8 3.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0ZM11.5 1a2.5 2.5 0 0 0-2.086 3.879l3.465-3.465A2.488 2.488 0 0 0 11.5 1Zm0 5c-.51 0-.983-.152-1.379-.414l3.465-3.465A2.5 2.5 0 0 1 11.5 6Z" clip-rule="evenodd"/>`
	phonecallOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="m14.5.5l-6 6m6-6V4m0-3.5H11M2.5.5h2.22a1 1 0 0 1 .97.757l.585 2.345a1 1 0 0 1-.654 1.19l-1.108.37a1.21 1.21 0 0 0-.804 1.385a6.047 6.047 0 0 0 4.744 4.744a1.21 1.21 0 0 0 1.385-.804l.297-.893a1 1 0 0 1 1.396-.578l2.416 1.208a1 1 0 0 1 .553.894V12.5a2 2 0 0 1-2 2h-2c-5.523 0-10-4.477-10-10v-2a2 2 0 0 1 2-2Z"/>`
	phonecallReceiveOutlineInnerSVG          = `<path fill="none" stroke="currentColor" d="m14.5.5l-6 6m0 0V3m0 3.5H12M2.5.5h2.22a1 1 0 0 1 .97.757l.585 2.345a1 1 0 0 1-.654 1.19l-1.108.37a1.21 1.21 0 0 0-.804 1.385a6.047 6.047 0 0 0 4.744 4.744a1.21 1.21 0 0 0 1.385-.804l.297-.893a1 1 0 0 1 1.396-.578l2.416 1.208a1 1 0 0 1 .553.894V12.5a2 2 0 0 1-2 2h-2c-5.523 0-10-4.477-10-10v-2a2 2 0 0 1 2-2Z"/>`
	phonecallReceiveSolidInnerSVG            = `<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 0 2.5v2C0 10.299 4.701 15 10.5 15h2a2.5 2.5 0 0 0 2.5-2.5v-1.382a1.5 1.5 0 0 0-.83-1.342l-2.415-1.208a1.5 1.5 0 0 0-2.094.868l-.298.893a.71.71 0 0 1-.812.471A5.547 5.547 0 0 1 4.2 6.45a.71.71 0 0 1 .471-.812l1.109-.37a1.5 1.5 0 0 0 .98-1.787l-.586-2.344A1.5 1.5 0 0 0 4.72 0H2.5Z"/><path fill="currentColor" d="M14.146.146L9 5.293V3H8v4h4V6H9.707L14.854.854l-.708-.708Z"/>`
	phonecallSolidInnerSVG                   = `<path fill="currentColor" d="M15 0h-4v1h2.293L8.146 6.146l.708.708L14 1.707V4h1V0Z"/><path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 0 2.5v2C0 10.299 4.701 15 10.5 15h2a2.5 2.5 0 0 0 2.5-2.5v-1.382a1.5 1.5 0 0 0-.83-1.342l-2.415-1.208a1.5 1.5 0 0 0-2.094.868l-.298.893a.71.71 0 0 1-.812.471A5.547 5.547 0 0 1 4.2 6.45a.71.71 0 0 1 .471-.812l1.109-.37a1.5 1.5 0 0 0 .98-1.787l-.586-2.344A1.5 1.5 0 0 0 4.72 0H2.5Z"/>`
	pieChartAltOutlineInnerSVG               = `<path fill="currentColor" d="M6.5.5V0a.5.5 0 0 0-.5.5h.5Zm8 8V9a.5.5 0 0 0 .5-.5h-.5Zm-8 0H6V9h.5v-.5Zm0 6.5A6.5 6.5 0 0 0 13 8.5h-1A5.5 5.5 0 0 1 6.5 14v1ZM0 8.5A6.5 6.5 0 0 0 6.5 15v-1A5.5 5.5 0 0 1 1 8.5H0Zm1 0A5.5 5.5 0 0 1 6.5 3V2A6.5 6.5 0 0 0 0 8.5h1ZM6.5 1A7.5 7.5 0 0 1 14 8.5h1A8.5 8.5 0 0 0 6.5 0v1ZM6 .5v8h1v-8H6ZM6.5 9h8V8h-8v1Z"/>`
	pieChartAltSolidInnerSVG                 = `<path fill="currentColor" d="M6.5 0H6v9h9v-.5A8.5 8.5 0 0 0 6.5 0Z"/><path fill="currentColor" d="M12.826 10H5V2.174A6.5 6.5 0 1 0 12.826 10Z"/>`
	pieChartOutlineInnerSVG                  = `<path fill="currentColor" d="M7.5 7.5H7a.5.5 0 0 0 .146.354L7.5 7.5Zm0 6.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM7 0v7.5h1V0H7Zm.724 7.947l6-3l-.448-.894l-6 3l.448.894Zm-.578-.093l5 5l.708-.708l-5-5l-.708.708Z"/>`
	pieChartSolidInnerSVG                    = `<path fill="currentColor" d="M7 .016a7.5 7.5 0 1 0 5.438 13.13L7.15 7.857A.5.5 0 0 1 7 7.5V.016Z"/><path fill="currentColor" d="M13.145 12.438A7.47 7.47 0 0 0 15 7.5a7.476 7.476 0 0 0-.581-2.9L8.344 7.637l4.801 4.801Zm.825-8.732A7.499 7.499 0 0 0 8 .016v6.675l5.97-2.985Z"/>`
	pinAltOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M7.5 15V8.5m0 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`
	pinAltSolidInnerSVG                      = `<path fill="currentColor" d="M7.5 0A4.5 4.5 0 0 0 7 8.973V15h1V8.973A4.5 4.5 0 0 0 7.5 0Z"/>`
	pinOutlineInnerSVG                       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="square" clip-rule="evenodd"><path d="M7.5 8.495a2 2 0 0 0 2-1.999a2 2 0 0 0-4 0a2 2 0 0 0 2 1.999Z"/><path d="M13.5 6.496c0 4.997-5 7.995-6 7.995s-6-2.998-6-7.995A5.999 5.999 0 0 1 7.5.5c3.313 0 6 2.685 6 5.996Z"/></g>`
	pinSolidInnerSVG                         = `<path fill="currentColor" d="M6 6.496a1.5 1.5 0 0 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.496A6.499 6.499 0 0 1 7.5 0C11.089 0 14 2.909 14 6.496c0 2.674-1.338 4.793-2.772 6.225a10.865 10.865 0 0 1-2.115 1.654c-.322.19-.623.34-.885.442c-.247.098-.506.174-.728.174c-.222 0-.481-.076-.728-.174a6.453 6.453 0 0 1-.885-.442a10.868 10.868 0 0 1-2.115-1.654C2.338 11.289 1 9.17 1 6.496Zm6.5-2.499a2.5 2.5 0 0 0-2.5 2.5a2.5 2.5 0 0 0 5 0a2.5 2.5 0 0 0-2.5-2.5Z" clip-rule="evenodd"/>`
	pinterestOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="m4.5 13.5l3-7m-3.236 3a2.989 2.989 0 0 1-.764-2V7A3.5 3.5 0 0 1 7 3.5h1A3.5 3.5 0 0 1 11.5 7v.5a3 3 0 0 1-3 3a2.081 2.081 0 0 1-1.974-1.423L6.5 9m1 5.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	pinterestSolidInnerSVG                   = `<path fill="currentColor" d="M0 7.5a7.5 7.5 0 1 1 4.584 6.912l1.911-4.367A2.58 2.58 0 0 0 8.5 11A3.5 3.5 0 0 0 12 7.5V7a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v.5c0 .896.337 1.715.891 2.333l.745-.666A2.489 2.489 0 0 1 4 7.5V7a3 3 0 0 1 3-3h1a3 3 0 0 1 3 3v.5A2.5 2.5 0 0 1 8.5 10A1.58 1.58 0 0 1 7 8.919l-.005-.016l.963-2.203l-.916-.4l-3.352 7.66A7.496 7.496 0 0 1 0 7.5Z"/>`
	plantOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M7.5 15V7m0 .5v3m0-3a4 4 0 0 0-4-4h-3v3a4 4 0 0 0 4 4h3m0-3h3a4 4 0 0 0 4-4v-3h-3a4 4 0 0 0-4 4v3Zm0 0l4-4m-4 7l-4-4"/>`
	plantSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M7 4.5A4.5 4.5 0 0 1 11.5 0H15v3.5A4.5 4.5 0 0 1 10.5 8H8v7H7v-4H4.5A4.5 4.5 0 0 1 0 6.5V3h3.5c1.414 0 2.675.652 3.5 1.671V4.5Zm1.146 1.646l3-3l.708.708l-3 3l-.708-.708Zm-2 3.708l-3-3l.708-.708l3 3l-.708.708Z" clip-rule="evenodd"/>`
	playCircleOutlineInnerSVG                = `<path fill="currentColor" d="m6.5 5.5l.248-.434A.5.5 0 0 0 6 5.5h.5Zm0 4H6a.5.5 0 0 0 .748.434L6.5 9.5Zm3.5-2l.248.434a.5.5 0 0 0 0-.868L10 7.5ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM6 5.5v4h1v-4H6Zm.748 4.434l3.5-2l-.496-.868l-3.5 2l.496.868Zm3.5-2.868l-3.5-2l-.496.868l3.5 2l.496-.868Z"/>`
	playCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm6.249-2.432a.5.5 0 0 1 .5-.002l3.5 2a.5.5 0 0 1 0 .868l-3.5 2A.5.5 0 0 1 6 9.5v-4a.5.5 0 0 1 .249-.432Z" clip-rule="evenodd"/>`
	playOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M4.5 12.5v-10l7 5l-7 5Z"/>`
	playSmallOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M6.5 9.5v-4l3.5 2l-3.5 2Z"/>`
	playSmallSolidInnerSVG                   = `<path fill="currentColor" d="M6.748 5.066A.5.5 0 0 0 6 5.5v4a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868l-3.5-2Z"/>`
	playSolidInnerSVG                        = `<path fill="currentColor" d="M4.79 2.093A.5.5 0 0 0 4 2.5v10a.5.5 0 0 0 .79.407l7-5a.5.5 0 0 0 0-.814l-7-5Z"/>`
	plugOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M6.5 11.5V15m2-3.5V15m-4-15v4.5m6-4.5v4.5m-8 0h10v3h-1v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-2h-1v-3Z"/>`
	plugSolidInnerSVG                        = `<path fill="currentColor" d="M10 4H5V0H4v4H2v4h1v1.5A2.5 2.5 0 0 0 5.5 12H6v3h1v-3h1v3h1v-3h.5A2.5 2.5 0 0 0 12 9.5V8h1V4h-2V0h-1v4Z"/>`
	plusCircleOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M7.5 4v7M4 7.5h7m-3.5 7a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	plusCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM7 11V8H4V7h3V4h1v3h3v1H8v3H7Z" clip-rule="evenodd"/>`
	pngOutlineInnerSVG                       = `<path fill="currentColor" d="M2.5 6.5V6H2v.5h.5Zm8 4H10v.5h.5v-.5Zm2 0v.5h.5v-.5h-.5Zm1-7h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-4 6l.447-.224L6 6.5h.5Zm-.5 4v.5h1v-.5H6Zm2.5 0l-.447.224A.5.5 0 0 0 9 10.5h-.5Zm.5-4V6H8v.5h1ZM2.5 7h1V6h-1v1Zm.5 4V8.5H2V11h1Zm0-2.5v-2H2v2h1Zm.5-.5h-1v1h1V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 7.5H4ZM3.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 6v1ZM10 6v4.5h1V6h-1Zm.5 5h2v-1h-2v1Zm2.5-.5v-2h-1v2h1ZM10.5 7H13V6h-2.5v1ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1Zm5-7v4h1v-4H6Zm.053.224l2 4l.894-.448l-2-4l-.894.448ZM8 6.5v4h1v-4H8Z"/>`
	pngSolidInnerSVG                         = `<path fill="currentColor" d="M3 8h.5a.5.5 0 0 0 0-1H3v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM2 6h1.5a1.5 1.5 0 1 1 0 3H3v2H2V6Zm8 0h3v1h-2v3h1V8.5h1V11h-3V6ZM7 8.618V11H6V6h1v.382l1 2V6h1v5H8v-.382l-1-2Z" clip-rule="evenodd"/>`
	poolOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="10" d="M1 12.6c.65.733 1.655 1.4 2.955 1.4c2.954 0 3.545-2 6.5-2c1.359 0 2.6.733 3.545 1.467M2.5 12V3.727C2.5 1.945 3.855.5 5.636.5M9.5 10V3.636C9.5 1.855 10.9.5 12.682.5M2.5 4.5h7m-7 4h7"/>`
	poolSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M5.636 1C4.15 1 3 2.203 3 3.727V4h6v-.364C9 1.57 10.633 0 12.682 0v1C11.167 1 10 2.14 10 3.636V10H9V9H3v3H2V3.727C2 1.688 3.56 0 5.636 0v1ZM3 8h6V5H3v3Z" clip-rule="evenodd"/><path fill="currentColor" d="M7.44 13.442c-.895.504-1.877 1.058-3.485 1.058c-1.483 0-2.614-.762-3.33-1.568l.75-.664c.584.66 1.462 1.232 2.58 1.232c1.339 0 2.128-.442 3.004-.935l.01-.007c.895-.504 1.877-1.058 3.485-1.058c1.531 0 2.884.82 3.852 1.572l-.612.79c-.923-.716-2.052-1.362-3.24-1.362c-1.339 0-2.128.442-3.004.935l-.01.007Z"/>`
	poundOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="m.5 14.5l1.103-.367A2.775 2.775 0 0 0 3.5 11.5V4.442A3.942 3.942 0 0 1 7.442.5h.865C9.934.5 11.396 1.49 12 3M3 13h1.084a6 6 0 0 1 2.683.633l.05.025a6 6 0 0 0 5.366 0L13.5 13M1 7.5h8"/>`
	poundSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M3 4.442A4.442 4.442 0 0 1 7.442 0h.865a4.477 4.477 0 0 1 4.157 2.814l-.928.372A3.477 3.477 0 0 0 8.307 1h-.865A3.44 3.44 0 0 0 4 4.442V7h5v1H4v3.5c0 .346-.054.683-.156 1h.24a6.5 6.5 0 0 1 2.906.686l.05.025l-.223.447l.223-.447a5.5 5.5 0 0 0 4.92 0l1.316-.658l.448.894l-1.317.659a6.5 6.5 0 0 1-5.814 0l-.05-.025l.224-.448l-.224.448a5.5 5.5 0 0 0-2.46-.581h-.765a3.27 3.27 0 0 1-1.557 1.107l-1.103.367l-.316-.948l1.103-.368A2.275 2.275 0 0 0 3 11.5V8H1V7h2V4.442Z" clip-rule="evenodd"/>`
	powerOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5 8.5v-8M2.618 2.499A6.963 6.963 0 0 0 .5 7.495c0 3.864 3.135 7.005 7 7.005c3.867 0 7-3.141 7-7.005A6.968 6.968 0 0 0 12.395 2.5"/>`
	powerSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M8 0v9H7V0h1Zm4.387 1.792l.358.35A7.468 7.468 0 0 1 15 7.494C15 11.635 11.644 15 7.5 15C3.358 15 0 11.635 0 7.495a7.46 7.46 0 0 1 2.269-5.354l.357-.35l.7.716l-.359.35A6.463 6.463 0 0 0 1 7.494A6.506 6.506 0 0 0 7.5 14c3.59 0 6.5-2.917 6.5-6.505a6.464 6.464 0 0 0-1.955-4.639l-.357-.35l.7-.714Z" clip-rule="evenodd"/>`
	pptOutlineInnerSVG                       = `<path fill="currentColor" d="M2.5 6.5V6H2v.5h.5Zm4 0V6H6v.5h.5Zm7-3h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5ZM2.5 7h1V6h-1v1Zm.5 4V8.5H2V11h1Zm0-2.5v-2H2v2h1Zm.5-.5h-1v1h1V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 7.5H4ZM3.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 6v1Zm3 0h1V6h-1v1Zm.5 4V8.5H6V11h1Zm0-2.5v-2H6v2h1Zm.5-.5h-1v1h1V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5H8ZM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6v1ZM11 6v5h1V6h-1Zm-1 1h3V6h-3v1ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1Z"/>`
	pptSolidInnerSVG                         = `<path fill="currentColor" d="M3 8h.5a.5.5 0 0 0 0-1H3v1Zm4 0h.5a.5.5 0 0 0 0-1H7v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM2 6h1.5a1.5 1.5 0 1 1 0 3H3v2H2V6Zm4 0h1.5a1.5 1.5 0 1 1 0 3H7v2H6V6Zm5 5h1V7h1V6h-3v1h1v4Z" clip-rule="evenodd"/>`
	printOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M3.5 12.5h-2a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-2m-8-6v-5a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v5m-8 4h8v4h-8v-4Z"/>`
	printSolidInnerSVG                       = `<path fill="currentColor" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h6A1.5 1.5 0 0 1 12 1.5V5H3V1.5ZM1.5 6A1.5 1.5 0 0 0 0 7.5v4A1.5 1.5 0 0 0 1.5 13H2V9h11v4h.5a1.5 1.5 0 0 0 1.5-1.5v-4A1.5 1.5 0 0 0 13.5 6h-12Z"/><path fill="currentColor" d="M3 10h9v5H3v-5Z"/>`
	pythonOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M6 2.5h1M4.5 4V1.5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-4a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V11M8 4.5H1.5a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h3m2.5-1h6.5a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1h-3m-2.5 9h1"/>`
	pythonSolidInnerSVG                      = `<path fill="currentColor" d="M3 12H1.5A1.5 1.5 0 0 1 0 10.5v-5A1.5 1.5 0 0 1 1.5 4H8V3H7V2H6v1H4V1.5A1.5 1.5 0 0 1 5.5 0h4A1.5 1.5 0 0 1 11 1.5v5a.5.5 0 0 1-.5.5h-6A1.5 1.5 0 0 0 3 8.5V12Z"/><path fill="currentColor" d="M12 3v3.5A1.5 1.5 0 0 1 10.5 8h-6a.5.5 0 0 0-.5.5v5A1.5 1.5 0 0 0 5.5 15h4a1.5 1.5 0 0 0 1.5-1.5V12H9v1H8v-1H7v-1h6.5A1.5 1.5 0 0 0 15 9.5v-5A1.5 1.5 0 0 0 13.5 3H12Z"/>`
	qrCodeOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M12 8.5H8.5V12M14 8.5h1m-3 6H8m3-3h3.5V15M3 3.5h1m7 0h1m-9 8h1M1.5.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Zm8 0h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Zm-8 8h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Z"/>`
	qrCodeSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h4A1.5 1.5 0 0 1 7 1.5v4A1.5 1.5 0 0 1 5.5 7h-4A1.5 1.5 0 0 1 0 5.5v-4ZM1.5 1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4Zm6.5.5A1.5 1.5 0 0 1 9.5 0h4A1.5 1.5 0 0 1 15 1.5v4A1.5 1.5 0 0 1 13.5 7h-4A1.5 1.5 0 0 1 8 5.5v-4ZM9.5 1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4ZM4 4H3V3h1v1Zm8 0h-1V3h1v1ZM0 9.5A1.5 1.5 0 0 1 1.5 8h4A1.5 1.5 0 0 1 7 9.5v4A1.5 1.5 0 0 1 5.5 15h-4A1.5 1.5 0 0 1 0 13.5v-4ZM1.5 9a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4ZM8 8h4v1H9v3H8V8Zm7 1h-1V8h1v1ZM4 12H3v-1h1v1Zm10 0h-3v-1h4v4h-1v-3Zm-6 2h4v1H8v-1Z" clip-rule="evenodd"/>`
	questionCircleOutlineInnerSVG            = `<path fill="none" stroke="currentColor" d="M7.5 9V7.5H8A1.5 1.5 0 0 0 9.5 6v-.1a1.4 1.4 0 0 0-1.4-1.4h-.6A1.5 1.5 0 0 0 6 6m1 4.5h1m-.5 4a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	questionCircleSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM5.5 6a2 2 0 0 1 2-2h.6c1.05 0 1.9.85 1.9 1.9V6a2 2 0 0 1-2 2v1H7V7h1a1 1 0 0 0 1-1v-.1a.9.9 0 0 0-.9-.9h-.6a1 1 0 0 0-1 1h-1ZM7 11v-1h1v1H7Z" clip-rule="evenodd"/>`
	questionOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M7.5 12v-1.264c0-1.37.774-2.623 2-3.236a3.65 3.65 0 0 0 2-3.257C11.5 2.195 9.84.5 7.792.5h-.207c-1.335 0-2.615.53-3.56 1.474L3.5 2.5m3.5 12h1"/>`
	questionSmallOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M7.5 9V7.5H8A1.5 1.5 0 0 0 9.5 6v-.1a1.4 1.4 0 0 0-1.4-1.4h-.6A1.5 1.5 0 0 0 6 6m1 4.5h1"/>`
	questionSmallSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 6a2 2 0 0 1 2-2h.6c1.05 0 1.9.85 1.9 1.9V6a2 2 0 0 1-2 2v1H7V7h1a1 1 0 0 0 1-1v-.1a.9.9 0 0 0-.9-.9h-.6a1 1 0 0 0-1 1h-1ZM8 10v1H7v-1h1Z" clip-rule="evenodd"/>`
	questionSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M3.672 1.62A5.534 5.534 0 0 1 7.585 0h.207C10.122 0 12 1.925 12 4.243a4.15 4.15 0 0 1-2.276 3.704A3.118 3.118 0 0 0 8 10.737V12H7v-1.264c0-1.56.881-2.986 2.276-3.683A3.15 3.15 0 0 0 11 4.243C11 2.465 9.558 1 7.792 1h-.207a4.534 4.534 0 0 0-3.206 1.328l-.525.526l-.708-.708l.526-.525ZM8 15H7v-1h1v1Z" clip-rule="evenodd"/>`
	quoteOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M1.5 6.5h4a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v4Zm0 0V10A3.5 3.5 0 0 0 5 13.5m3.5-7h4a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v4Zm0 0V10a3.5 3.5 0 0 0 3.5 3.5"/>`
	quoteSolidInnerSVG                       = `<path fill="currentColor" d="M2.5 1A1.5 1.5 0 0 0 1 2.5V10a4 4 0 0 0 4 4v-1a3 3 0 0 1-3-3V7h3.5A1.5 1.5 0 0 0 7 5.5v-3A1.5 1.5 0 0 0 5.5 1h-3Zm7 0A1.5 1.5 0 0 0 8 2.5V10a4 4 0 0 0 4 4v-1a3 3 0 0 1-3-3V7h3.5A1.5 1.5 0 0 0 14 5.5v-3A1.5 1.5 0 0 0 12.5 1h-3Z"/>`
	randOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M3.5 14V8.5m0 0v-7H8a3.5 3.5 0 1 1 0 7H3.5Zm0 0h5a3 3 0 0 1 3 3V14"/>`
	randSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M3 1h5a4 4 0 0 1 2.117 7.395A3.5 3.5 0 0 1 12 11.5V14h-1v-2.5A2.5 2.5 0 0 0 8.5 9H4v5H3V1Zm1 7h4a3 3 0 0 0 0-6H4v6Z" clip-rule="evenodd"/>`
	reactOutlineInnerSVG                     = `<g fill="none" stroke="currentColor"><path d="M14.5 7.584c0 1.657-3.134 3-7 3s-7-1.343-7-3s3.134-3 7-3s7 1.343 7 3Z"/><path d="M4.166 13.739c1.457.79 4.13-1.327 5.972-4.726c1.841-3.4 2.153-6.795.696-7.584c-1.457-.79-4.13 1.327-5.972 4.726c-1.841 3.4-2.153 6.795-.696 7.584Z"/><path d="M10.834 13.739c-1.457.79-4.13-1.327-5.972-4.726c-1.841-3.4-2.153-6.795-.696-7.584c1.457-.79 4.13 1.327 5.972 4.726c1.841 3.4 2.153 6.795.696 7.584Z"/><path d="M6.5 7.584a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/></g>`
	reactSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M5.315 1.837c-.4-.116-.695-.085-.91.032c-.216.116-.404.347-.526.745c-.122.401-.163.936-.104 1.582c.01.105.022.212.037.321a13.654 13.654 0 0 1 1.676-.311a13.28 13.28 0 0 1 1.275-1.54a7.111 7.111 0 0 0-.066-.053c-.508-.402-.98-.66-1.382-.776Zm2.185.14c-.06-.05-.121-.1-.182-.148c-.572-.452-1.158-.789-1.724-.953C5.024.711 4.441.711 3.928.99c-.513.278-.833.767-1.005 1.334c-.172.564-.21 1.238-.144 1.965c.016.17.037.344.065.523c-.17.06-.334.125-.49.192c-.671.287-1.246.642-1.66 1.062C.278 6.487 0 7 0 7.584c0 .583.278 1.097.694 1.519c.414.42.989.774 1.66 1.062c.156.067.32.131.49.192a8.672 8.672 0 0 0-.065.523c-.066.726-.028 1.4.144 1.965c.172.567.492 1.056 1.005 1.333c.513.278 1.097.279 1.666.114c.566-.165 1.152-.5 1.724-.953l.182-.149c.06.051.121.1.182.149c.572.452 1.158.788 1.724.953c.569.165 1.153.164 1.666-.114c.513-.277.833-.766 1.005-1.333c.172-.564.21-1.239.144-1.965a8.616 8.616 0 0 0-.065-.523a8.2 8.2 0 0 0 .49-.192c.671-.288 1.246-.643 1.66-1.062c.416-.422.694-.936.694-1.52c0-.582-.278-1.096-.694-1.518c-.414-.42-.989-.775-1.66-1.062a8.706 8.706 0 0 0-.49-.192c.027-.179.049-.353.065-.523c.066-.727.028-1.4-.144-1.965c-.172-.567-.492-1.056-1.005-1.334C10.56.711 9.975.711 9.406.876c-.566.164-1.152.5-1.724.953l-.182.149Zm0 1.365c-.225.23-.45.483-.672.755a16.99 16.99 0 0 1 1.344 0a11.385 11.385 0 0 0-.672-.755Zm2.012.864c-.41-.574-.84-1.092-1.275-1.54l.065-.053c.51-.402.98-.66 1.383-.776c.399-.116.695-.085.91.032c.216.116.404.347.525.745c.122.401.164.936.105 1.582c-.01.105-.022.212-.037.32a13.655 13.655 0 0 0-1.676-.31Zm-.563.944a15.628 15.628 0 0 0-2.898 0A15.627 15.627 0 0 0 4.72 7.584a15.693 15.693 0 0 0 1.33 2.433a15.634 15.634 0 0 0 2.9 0a15.63 15.63 0 0 0 1.33-2.433A15.691 15.691 0 0 0 8.95 5.15Zm1.824 1.138a17.244 17.244 0 0 0-.527-.956c.26.05.511.106.752.168c-.063.256-.138.52-.225.788Zm0 2.591a16.837 16.837 0 0 1-.527.957c.26-.05.511-.106.752-.169a11.69 11.69 0 0 0-.225-.788Zm1.18.487a13.805 13.805 0 0 0-.588-1.782c.246-.61.443-1.209.588-1.782c.103.038.203.079.3.12c.596.256 1.047.547 1.341.845c.292.296.406.572.406.817s-.114.52-.406.816c-.294.299-.745.59-1.341.846a7.467 7.467 0 0 1-.3.12Zm-.765 1.285a13.57 13.57 0 0 1-1.676.311c-.41.574-.84 1.091-1.275 1.54l.066.052c.508.403.98.66 1.382.777c.399.116.695.085.91-.032c.216-.117.404-.348.525-.746c.123-.4.164-.936.105-1.582a7.422 7.422 0 0 0-.037-.32ZM7.5 11.826c.225-.23.45-.483.672-.755a16.945 16.945 0 0 1-1.344 0c.222.272.447.524.672.755Zm-2.746-1.99a16.972 16.972 0 0 1-.527-.957c-.087.27-.162.532-.225.788c.24.063.492.119.752.169Zm-.942.815a13.57 13.57 0 0 0 1.676.311c.41.574.839 1.091 1.275 1.54a6.761 6.761 0 0 1-.066.052c-.508.403-.98.66-1.382.777c-.4.116-.695.085-.911-.032c-.216-.117-.403-.348-.525-.746c-.122-.4-.163-.936-.104-1.582a7.47 7.47 0 0 1 .037-.32Zm-.765-1.285c.145-.574.341-1.172.588-1.782a13.806 13.806 0 0 1-.588-1.782a8.518 8.518 0 0 0-.3.12c-.596.256-1.047.547-1.341.845c-.292.296-.406.572-.406.817s.114.52.406.816c.294.299.745.59 1.341.846c.097.041.197.081.3.12Zm.955-3.865c.063.255.138.519.225.787a17.116 17.116 0 0 1 .527-.956c-.26.05-.511.106-.752.169ZM6 7.584a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`
	receiptOutlineInnerSVG                   = `<path fill="currentColor" d="M1.5.5V0a.5.5 0 0 0-.5.5h.5Zm12 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 14l-.224.447A.5.5 0 0 0 14 14.5h-.5Zm-2-1l.224-.447a.5.5 0 0 0-.448 0l.224.447Zm-2 1l-.224.447a.5.5 0 0 0 .448 0L9.5 14.5Zm-2-1l.224-.447a.5.5 0 0 0-.448 0l.224.447Zm-2 1l-.224.447a.5.5 0 0 0 .448 0L5.5 14.5Zm-4 0H1a.5.5 0 0 0 .724.447L1.5 14.5Zm2-1l.224-.447a.5.5 0 0 0-.448 0l.224.447ZM1.5 1h12V0h-12v1ZM13 .5v14h1V.5h-1Zm.724 13.553l-2-1l-.448.894l2 1l.448-.894Zm-2.448-1l-2 1l.448.894l2-1l-.448-.894Zm-1.552 1l-2-1l-.448.894l2 1l.448-.894Zm-2.448-1l-2 1l.448.894l2-1l-.448-.894ZM2 14.5V.5H1v14h1Zm3.724-.447l-2-1l-.448.894l2 1l.448-.894Zm-2.448-1l-2 1l.448.894l2-1l-.448-.894ZM4 5h2V4H4v1Zm4 0h3V4H8v1ZM4 8h2V7H4v1Zm4 0h3V7H8v1Zm0 3h3v-1H8v1Z"/>`
	receiptSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M1 .5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-.724.447L11.5 14.06l-1.776.888a.5.5 0 0 1-.448 0L7.5 14.06l-1.776.888a.5.5 0 0 1-.448 0L3.5 14.06l-1.776.888A.5.5 0 0 1 1 14.5V.5ZM4 5h2V4H4v1Zm4 0h3V4H8v1ZM6 8H4V7h2v1Zm2 0h3V7H8v1Zm3 3H8v-1h3v1Z" clip-rule="evenodd"/>`
	redditOutlineInnerSVG                    = `<path fill="currentColor" d="m7.5 1.5l.121-.485A.5.5 0 0 0 7 1.5h.5Zm5.5 8c0 .774-.55 1.641-1.583 2.343C10.4 12.533 8.998 13 7.5 13v1c1.696 0 3.294-.525 4.479-1.33C13.148 11.876 14 10.743 14 9.5h-1ZM7.5 13c-1.498 0-2.9-.466-3.917-1.157C2.551 11.14 2 10.273 2 9.5H1c0 1.243.852 2.376 2.021 3.17C4.206 13.475 5.804 14 7.5 14v-1ZM2 9.5c0-.774.55-1.641 1.583-2.343C4.6 6.467 6.002 6 7.5 6V5c-1.696 0-3.294.525-4.479 1.33C1.852 7.124 1 8.257 1 9.5h1ZM7.5 6c1.498 0 2.9.467 3.917 1.157C12.449 7.86 13 8.727 13 9.5h1c0-1.243-.852-2.376-2.021-3.17C10.794 5.525 9.196 5 7.5 5v1Zm2.306 4.54c-.69.29-1.32.46-2.306.46v1c1.136 0 1.898-.204 2.694-.54l-.388-.92ZM7.5 11c-.987 0-1.617-.17-2.306-.46l-.388.92c.796.336 1.558.54 2.694.54v-1ZM8 5.5v-4H7v4h1Zm-.621-3.515l4 1l.242-.97l-4-1l-.242.97ZM3.974 6.841c-.286-.855-1.12-1.297-1.952-1.297v1c.51 0 .886.261 1.004.615l.948-.318ZM2.022 5.544A2.022 2.022 0 0 0 0 7.566h1a1.02 1.02 0 0 1 1.022-1.022v-1ZM0 7.566C0 8.589.76 9.424 1.74 9.56l.139-.99A1.016 1.016 0 0 1 1 7.565H0Zm11.974-.407c.118-.354.493-.615 1.004-.615v-1c-.832 0-1.666.442-1.952 1.297l.948.318Zm1.004-.615A1.02 1.02 0 0 1 14 7.566h1a2.022 2.022 0 0 0-2.022-2.022v1ZM14 7.566c0 .511-.38.934-.879 1.004l.139.99A2.016 2.016 0 0 0 15 7.567h-1ZM12.5 3a.5.5 0 0 1-.5-.5h-1A1.5 1.5 0 0 0 12.5 4V3Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 14 2.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 12.5 1v1Zm0-1A1.5 1.5 0 0 0 11 2.5h1a.5.5 0 0 1 .5-.5V1ZM5 9h1V8H5v1Zm4 0h1V8H9v1Z"/>`
	redditSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7.621 1.015A.5.5 0 0 0 7 1.5v3.515A8.16 8.16 0 0 0 3.455 6.06c-.388-.341-.911-.515-1.433-.515A2.022 2.022 0 0 0 0 7.566c0 .747.406 1.394 1.007 1.742A2.725 2.725 0 0 0 1 9.5c0 1.243.852 2.376 2.021 3.17C4.206 13.475 5.804 14 7.5 14c1.696 0 3.294-.525 4.479-1.33C13.148 11.876 14 10.743 14 9.5c0-.064-.002-.128-.007-.192A2.008 2.008 0 0 0 15 7.566a2.022 2.022 0 0 0-2.022-2.022c-.522 0-1.045.174-1.433.515A8.16 8.16 0 0 0 8 5.015V2.14l3.055.764a1.5 1.5 0 1 0 .074-1.012L7.62 1.015ZM12.5 3a.5.5 0 0 1-.5-.492v-.016a.5.5 0 1 1 .5.508ZM5 9h1V8H5v1Zm2.5 2c-.987 0-1.617-.17-2.306-.46l-.388.92c.796.336 1.558.54 2.694.54s1.898-.204 2.694-.54l-.388-.92c-.69.29-1.32.46-2.306.46ZM10 9H9V8h1v1Z" clip-rule="evenodd"/>`
	redwoodjsOutlineInnerSVG                 = `<path fill="currentColor" d="m7.5.5l.248-.434a.5.5 0 0 0-.496 0L7.5.5ZM4 2.5l-.248-.434a.5.5 0 0 0-.142.122L4 2.5ZM2 5l-.39-.312a.5.5 0 0 0-.09.175L2 5ZM1 8.5l-.48-.137a.5.5 0 0 0 .033.36L1 8.5Zm2 4l-.447.224a.5.5 0 0 0 .244.233L3 12.5Zm4.5 2l-.203.457a.5.5 0 0 0 .406 0L7.5 14.5Zm4.5-2l.203.457a.5.5 0 0 0 .244-.233L12 12.5Zm2-4l.447.224a.5.5 0 0 0 .034-.361L14 8.5ZM13 5l.48-.137a.5.5 0 0 0-.09-.175L13 5Zm-2-2.5l.39-.312a.499.499 0 0 0-.142-.122L11 2.5ZM7.252.066l-3.5 2l.496.868l3.5-2l-.496-.868ZM3.61 2.188l-2 2.5l.78.624l2-2.5l-.78-.624ZM1.52 4.863l-1 3.5l.96.274l1-3.5l-.96-.274Zm-.967 3.86l2 4l.894-.447l-2-4l-.894.448Zm2.244 4.234l4.5 2l.406-.914l-4.5-2l-.406.914Zm4.906 2l4.5-2l-.406-.914l-4.5 2l.406.914Zm4.744-2.233l2-4l-.894-.448l-2 4l.894.448Zm2.034-4.361l-1-3.5l-.962.274l1 3.5l.962-.274Zm-1.09-3.675l-2-2.5l-.781.624l2 2.5l.78-.624Zm-2.143-2.622l-3.5-2l-.496.868l3.5 2l.496-.868Zm-8.011.86l10.5 6.5l.526-.851l-10.5-6.5l-.526.85Zm8-.851l-10.5 6.5l.526.85l10.5-6.5l-.526-.85Zm-9.59 3.279l5.5 5.5l.707-.708l-5.5-5.5l-.708.708Zm5.63 5.593l4 2l.447-.894l-4-2l-.448.894Zm5.37-6.3l-5.5 5.5l.707.707l5.5-5.5l-.708-.708Zm-5.37 5.406l-4 2l.447.894l4-2l-.448-.894Z"/>`
	redwoodjsSolidInnerSVG                   = `<path fill="currentColor" d="M7.252.066a.5.5 0 0 1 .496 0l3.36 1.92L7.5 4.337L3.893 1.985L7.252.065ZM3.185 2.718L1.96 4.25L4.2 6.49l2.386-1.556l-3.4-2.217ZM1.438 5.146L.52 8.363a.5.5 0 0 0 .034.36l.054.11l2.735-1.784l-1.904-1.903Zm-.38 4.587l1.371 2.743l4.227-2.113l-2.591-2.591l-3.007 1.961Zm2.219 3.437l4.02 1.787a.5.5 0 0 0 .406 0l4.02-1.787L7.5 11.06l-4.223 2.11Zm9.293-.694l1.372-2.743l-3.007-1.961l-2.59 2.591l4.226 2.113Zm1.823-3.643l.054-.11a.5.5 0 0 0 .034-.36l-.92-3.217l-1.903 1.903l2.735 1.784Zm-1.352-4.582l-1.226-1.533l-3.4 2.217l2.387 1.556l2.24-2.24ZM7.5 5.532l2.58 1.682L7.5 9.793l-2.58-2.58l2.58-1.68Z"/>`
	refreshAltOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M.5 7.5A7 7 0 0 1 13 3.17m1.5 4.33A7 7 0 0 1 2 11.83m3-.33H1.5V15m12-15v3.5H10"/>`
	refreshAltSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 1A6.5 6.5 0 0 0 1 7.5H0a7.5 7.5 0 0 1 13-5.1V0h1v4h-4V3h2.19A6.48 6.48 0 0 0 7.5 1Zm0 13A6.5 6.5 0 0 0 14 7.5h1a7.5 7.5 0 0 1-13 5.099V15H1v-4h4v1H2.81a6.48 6.48 0 0 0 4.69 2Z" clip-rule="evenodd"/>`
	refreshOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M7.5 14.5A7 7 0 0 1 3.17 2M7.5.5A7 7 0 0 1 11.83 13m-.33-3v3.5H15M0 1.5h3.5V5"/>`
	refreshSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M14 7.5A6.5 6.5 0 0 0 7.5 1V0a7.5 7.5 0 0 1 5.099 13H15v1h-4v-4h1v2.19a6.48 6.48 0 0 0 2-4.69ZM2.4 2H0V1h4v4H3V2.81A6.5 6.5 0 0 0 7.5 14v1A7.5 7.5 0 0 1 2.4 2Z" clip-rule="evenodd"/>`
	rewindCircleOutlineInnerSVG              = `<path fill="currentColor" d="M6.5 5.5H7a.5.5 0 0 0-.748-.434L6.5 5.5Zm0 4l-.248.434A.5.5 0 0 0 7 9.5h-.5ZM3 7.5l-.248-.434a.5.5 0 0 0 0 .868L3 7.5Zm7.5-2h.5a.5.5 0 0 0-.748-.434l.248.434Zm0 4l-.248.434A.5.5 0 0 0 11 9.5h-.5ZM7 7.5l-.248-.434a.5.5 0 0 0 0 .868L7 7.5Zm.5 7.5A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1ZM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Zm0 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM6 5.5v4h1v-4H6Zm.748 3.566l-3.5-2l-.496.868l3.5 2l.496-.868Zm-3.5-1.132l3.5-2l-.496-.868l-3.5 2l.496.868ZM10 5.5v4h1v-4h-1Zm.748 3.566l-3.5-2l-.496.868l3.5 2l.496-.868Zm-3.5-1.132l3.5-2l-.496-.868l-3.5 2l.496.868Z"/>`
	rewindCircleSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 0-15 0a7.5 7.5 0 0 0 15 0Zm-4.249-2.432a.5.5 0 0 0-.5-.002L7 6.924V5.5a.5.5 0 0 0-.748-.434l-3.5 2a.5.5 0 0 0 0 .868l3.5 2A.5.5 0 0 0 7 9.5V8.076l3.252 1.858A.5.5 0 0 0 11 9.5v-4a.5.5 0 0 0-.249-.432Z" clip-rule="evenodd"/>`
	rewindOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M14.5 12.5v-10l-7 5l7 5Zm-7 0v-10l-7 5l7 5Z"/>`
	rewindSmallOutlineInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M6.5 9.5v-4L3 7.5l3.5 2Zm4 0v-4L7 7.5l3.5 2Z"/>`
	rewindSmallSolidInnerSVG                 = `<path fill="currentColor" d="M7 5.5a.5.5 0 0 0-.748-.434l-3.5 2a.5.5 0 0 0 0 .868l3.5 2A.5.5 0 0 0 7 9.5V8.076l3.252 1.858A.5.5 0 0 0 11 9.5v-4a.5.5 0 0 0-.748-.434L7 6.924V5.5Z"/>`
	rewindSolidInnerSVG                      = `<path fill="currentColor" d="M8 2.5a.5.5 0 0 0-.79-.407l-7 5a.5.5 0 0 0 0 .814l7 5A.5.5 0 0 0 8 12.5V8.472l6.21 4.435A.5.5 0 0 0 15 12.5v-10a.5.5 0 0 0-.79-.407L8 6.528V2.5Z"/>`
	rightCircleOutlineInnerSVG               = `<path fill="currentColor" d="m6.146 10.146l-.353.354l.707.707l.354-.353l-.708-.708ZM9.5 7.5l.354.354l.353-.354l-.353-.354L9.5 7.5ZM6.854 4.146L6.5 3.793l-.707.707l.353.354l.708-.708Zm0 6.708l3-3l-.708-.708l-3 3l.708.708Zm3-3.708l-3-3l-.708.708l3 3l.708-.708ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1Zm0 1A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1Z"/>`
	rightCircleSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0ZM6 3.293L10.207 7.5L6 11.707V3.293Z" clip-rule="evenodd"/>`
	rightOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m5 14l7-6.5L5 1"/>`
	rightSmallOutlineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m6.5 10.5l3-3l-3-3"/>`
	rightSmallSolidInnerSVG                  = `<path fill="currentColor" d="M10.207 7.5L6 3.293v8.414L10.207 7.5Z"/>`
	rightSolidInnerSVG                       = `<path fill="currentColor" d="M12 7.5L4 0v15l8-7.5Z"/>`
	rippleOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m1.5 1.5l3.06 3.316a4 4 0 0 0 5.88 0L13.5 1.5m-12 12l3.06-3.316a4 4 0 0 1 5.88 0L13.5 13.5"/>`
	rippleSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="m4.193 5.155l-3.06-3.316l.734-.678l3.061 3.316a3.5 3.5 0 0 0 5.144 0l3.06-3.316l.735.678l-3.06 3.316a4.5 4.5 0 0 1-6.614 0Zm5.879 5.368a3.5 3.5 0 0 0-5.144 0l-3.06 3.316l-.735-.678l3.06-3.316a4.5 4.5 0 0 1 6.614 0l3.06 3.316l-.734.678l-3.061-3.316Z" clip-rule="evenodd"/>`
	robotOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M7.5 2.5a5 5 0 0 1 5 5v6a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1v-6a5 5 0 0 1 5-5Zm0 0V0M4 11.5h7M.5 8v4m14-4v4m-9-2.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm4 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`
	robotSolidInnerSVG                       = `<path fill="currentColor" d="M5 8.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm4 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M8 2.022A5.5 5.5 0 0 1 13 7.5v6a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 13.5v-6a5.5 5.5 0 0 1 5-5.478V0h1v2.022ZM5.5 7a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm4 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm1.5 5H4v-1h7v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M0 8v4h1V8H0Zm15 0h-1v4h1V8Z"/>`
	rollerOutlineInnerSVG                    = `<path fill="currentColor" d="m13.05 7.914l-.138-.48l.137.48ZM8.224 9.293l.138.48l-.138-.48ZM5.5 11.5V11a.5.5 0 0 0-.5.5h.5Zm4 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 3v.5a.5.5 0 0 0 .5-.5h-.5Zm-4 0H5a.5.5 0 0 0 .5.5v-.5ZM1.5 1h10V0h-10v1Zm10.5.5v2h1v-2h-1ZM11.5 4h-10v1h10V4ZM1 3.5v-2H0v2h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 5V4ZM12 3.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 13 3.5h-1ZM11.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 0v1Zm-10-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5V0ZM14 4.5v1.491h1V4.5h-1Zm-1.088 2.934L8.088 8.812l.275.962l4.824-1.379l-.275-.961ZM7 10.254V11.5h1v-1.246H7Zm1.088-1.442A1.5 1.5 0 0 0 7 10.254h1a.5.5 0 0 1 .363-.48l-.275-.962ZM14 5.992a1.5 1.5 0 0 1-1.088 1.442l.275.961A2.5 2.5 0 0 0 15 5.991h-1ZM12.5 3A1.5 1.5 0 0 1 14 4.5h1A2.5 2.5 0 0 0 12.5 2v1Zm-7 9h4v-1h-4v1Zm3.5-.5v3h1v-3H9Zm.5 2.5h-4v1h4v-1Zm-3.5.5v-3H5v3h1Z"/>`
	rollerSolidInnerSVG                      = `<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h10A1.5 1.5 0 0 1 13 1.5v.55a2.5 2.5 0 0 1 2 2.45v1.491a2.5 2.5 0 0 1-1.813 2.404L8.363 9.774a.5.5 0 0 0-.363.48V11h1.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5H7v-.746a1.5 1.5 0 0 1 1.088-1.442l4.824-1.378A1.5 1.5 0 0 0 14 5.99V4.5a1.5 1.5 0 0 0-1-1.415V3.5A1.5 1.5 0 0 1 11.5 5h-10A1.5 1.5 0 0 1 0 3.5v-2Z"/>`
	rollupjsOutlineInnerSVG                  = `<path fill="currentColor" d="M1.5.5V0a.5.5 0 0 0-.5.5h.5Zm12 14v.5a.5.5 0 0 0 .455-.707l-.455.207Zm-2.488-5.474l-.224-.447a.5.5 0 0 0-.231.654l.455-.207ZM7.36 2.78l.447.224l-.447-.224ZM4.765 14.058l-.428-.258l.428.258ZM8.772 3.317l.475-.158l-.475.158Zm.457 1.371l-.474.158l.474-.158Zm2.235-.295l-.474.158l.474-.158Zm-2.487-2.61l.417.278l-.417-.277Zm1.362 3.885l.098-.49l-.098.49ZM1 .5v12h1V.5H1ZM3.5 15h10v-1h-10v1Zm10.455-.707L11.467 8.82l-.91.414l2.488 5.474l.91-.414ZM1.5 1H9V0H1.5v1ZM13 5a4 4 0 0 1-2.212 3.58l.448.893A5 5 0 0 0 14 5h-1ZM9 1a4 4 0 0 1 4 4h1a5 5 0 0 0-5-5v1ZM1 12.5A2.5 2.5 0 0 0 3.5 15v-1A1.5 1.5 0 0 1 2 12.5H1Zm1.447 1.224l5.36-10.72l-.894-.447l-5.36 10.72l.894.447Zm5.36-10.72a2.838 2.838 0 0 1 2.624-1.568l.03-1a3.838 3.838 0 0 0-3.548 2.12l.894.448ZM4.929 14.757l.265-.442l-.857-.515l-.266.443l.858.514Zm.265-.442a58.922 58.922 0 0 1 6.182-8.486l-.752-.658a59.924 59.924 0 0 0-6.287 8.63l.857.514Zm3.104-10.84l.457 1.371l.949-.316l-.457-1.371l-.95.316Zm3.64.76l-.464-1.393l-.948.316l.464 1.393l.949-.316ZM10.443.5a2.26 2.26 0 0 0-1.88 1.006l.832.555a1.26 1.26 0 0 1 1.048-.561v-1Zm-.201 5.658a1.483 1.483 0 0 0 1.698-1.923l-.949.316c.117.352-.19.7-.553.627l-.196.98ZM8.755 4.846a1.974 1.974 0 0 0 1.486 1.312l.196-.98a.974.974 0 0 1-.733-.648l-.949.316Zm.492-1.687a1.26 1.26 0 0 1 .147-1.098l-.833-.555a2.26 2.26 0 0 0-.263 1.969l.949-.316Zm1.184-1.723a.438.438 0 0 1-.25-.09a.475.475 0 0 1-.186-.41c.026-.385.382-.436.447-.436v1c.132 0 .523-.082.551-.497a.525.525 0 0 0-.207-.454a.562.562 0 0 0-.325-.112l-.03 1Z"/>`
	rollupjsSolidInnerSVG                    = `<path fill="currentColor" d="M1 .5a.5.5 0 0 1 .5-.5H9c.782 0 1.523.18 2.182.5h-.135a3.838 3.838 0 0 0-4.134 2.057L1.325 13.733A2.489 2.489 0 0 1 1 12.5V.5Z"/><path fill="currentColor" d="M4.87 15h8.63a.5.5 0 0 0 .455-.707l-2.298-5.057A4.997 4.997 0 0 0 14 5a4.984 4.984 0 0 0-1.43-3.5h-2.128a1.26 1.26 0 0 0-1.195 1.659l.457 1.371c.11.332.39.579.733.648a.485.485 0 0 0 .178.002l.009-.01l.007.007a.484.484 0 0 0 .359-.626l-.464-1.393l.948-.316l.465 1.393a1.483 1.483 0 0 1-.736 1.793a55.058 55.058 0 0 0-5.95 8.315L4.872 15Z"/>`
	routerOutlineInnerSVG                    = `<path fill="currentColor" d="M1.5 9.5V9H1v.5h.5Zm12 0h.5V9h-.5v.5Zm0 5v.5h.5v-.5h-.5Zm-12 0H1v.5h.5v-.5Zm1.72-8.339C4.373 4.761 5.916 4 7.5 4V3c-1.917 0-3.732.924-5.052 2.525l.771.636ZM7.5 4c1.583 0 3.126.762 4.281 2.161l.771-.636C11.232 3.924 9.417 3 7.5 3v1Zm-6.614.318C2.658 2.17 5.04 1 7.5 1V0C4.71 0 2.055 1.33.114 3.682l.772.636ZM7.5 1c2.46 0 4.842 1.17 6.614 3.318l.772-.636C12.945 1.329 10.29 0 7.5 0v1ZM7 6v3h1V6H7Zm-5.5 4h12V9h-12v1ZM13 9.5v5h1v-5h-1Zm.5 4.5h-12v1h12v-1ZM2 14.5v-5H1v5h1Z"/>`
	routerSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 1C5.04 1 2.658 2.17.886 4.318l-.772-.636C2.055 1.329 4.71 0 7.5 0s5.445 1.33 7.386 3.682l-.772.636C12.342 2.17 9.96 1 7.5 1Zm0 3c-1.583 0-3.126.762-4.28 2.161l-.772-.636C3.768 3.924 5.583 3 7.5 3c1.917 0 3.732.924 5.052 2.525l-.771.636C10.626 4.761 9.083 4 7.5 4ZM7 9V6h1v3h6v6H1V9h6Z" clip-rule="evenodd"/>`
	rssOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M.5 13.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm14 1.5C14.5 6.992 8.008.5 0 .5m0 6A8.5 8.5 0 0 1 8.5 15"/>`
	rssSolidInnerSVG                         = `<path fill="currentColor" d="M14 15C14 7.268 7.732 1 0 1V0c8.284 0 15 6.716 15 15h-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 13.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/><path fill="currentColor" d="M9 15a9 9 0 0 0-9-9v1a8 8 0 0 1 8 8h1Z"/>`
	rubyOutlineInnerSVG                      = `<path fill="currentColor" d="M14.5 14.5v.5h.5v-.5h-.5Zm0-14h.5V0h-.5v.5Zm-6 0V0h-.207l-.147.146L8.5.5Zm-8 8l-.354-.354L0 8.293V8.5h.5Zm0 6H0v.5h.5v-.5Zm4-4l-.224.447l.019.01l.02.007l.185-.464Zm0-6V4a.5.5 0 0 0-.5.5h.5Zm6 0l.464-.186l-.008-.019l-.009-.019l-.447.224Zm4.5 10V.5h-1v14h1ZM14.5 0h-6v1h6V0ZM8.146.146l-8 8l.708.708l8-8l-.708-.708ZM0 8.5v6h1v-6H0ZM.5 15h14v-1H.5v1ZM14.146.146l-14 14l.708.708l14-14l-.708-.708ZM5 10.5v-6H4v6h1ZM4.5 5h6V4h-6v1Zm-.186 5.964l10 4l.372-.928l-10-4l-.372.928Zm5.722-6.278l4 10l.928-.372l-4-10l-.928.372ZM8.053.724l2 4l.894-.448l-2-4l-.894.448ZM.276 8.947l4 2l.448-.894l-4-2l-.448.894Z"/>`
	rubySolidInnerSVG                        = `<path fill="currentColor" d="M4.293 4L8.058.236L9.73 4H4.293Zm10-4l-3.632 3.632L9.047 0h5.246ZM.236 8.058L4 9.73V4.293L.236 8.058Zm3.396 2.603L0 9.047v5.246l3.632-3.632ZM5 9.293L9.293 5H5v4.293Zm10 4.438l-3.907-9.117L15 .707v13.024Zm-.952.317l-3.717-8.672l-4.955 4.955l8.672 3.717Zm-9.434-2.955L13.731 15H.707l3.907-3.907Z"/>`
	rupeeOutlineInnerSVG                     = `<path fill="currentColor" d="M2.5 8.5V8a.5.5 0 0 0-.325.88L2.5 8.5ZM2 1h11V0H2v1Zm.5 8h3V8h-3v1Zm3-9h-3v1h3V0ZM2.175 8.88l7 6l.65-.76l-7-6l-.65.76ZM10 4.5A4.5 4.5 0 0 0 5.5 0v1A3.5 3.5 0 0 1 9 4.5h1ZM5.5 9A4.5 4.5 0 0 0 10 4.5H9A3.5 3.5 0 0 1 5.5 8v1ZM2 5h11V4H2v1Z"/>`
	rupeeSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M2 1V0h11v1H8.329a4.494 4.494 0 0 1 1.644 3H13v1H9.973A4.5 4.5 0 0 1 5.5 9H3.852l5.973 5.12l-.65.76l-7-6A.5.5 0 0 1 2.5 8h3a3.5 3.5 0 0 0 3.465-3H2V4h6.965A3.5 3.5 0 0 0 5.5 1H2Z" clip-rule="evenodd"/>`
	rustOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M2 10.5h5m-4.5-6H9a1.5 1.5 0 1 1 0 3H4.5m0-3v6m3-3h1a2 2 0 0 1 2 2a1 1 0 0 0 1 1H13M7.5.5l1.344 1.11l1.693-.417l.73 1.584l1.706.359l-.03 1.743l1.382 1.063L13.54 7.5l.784 1.558l-1.382 1.063l.03 1.743l-1.707.359l-.729 1.584l-1.693-.418L7.5 14.5l-1.344-1.11l-1.693.417l-.73-1.584l-1.706-.359l.03-1.743L.675 9.058L1.46 7.5L.675 5.942L2.057 4.88l-.03-1.743l1.706-.359l.73-1.584l1.693.417L7.5.5Z"/>`
	rustSolidInnerSVG                        = `<path fill="currentColor" d="M7.174.115a.521.521 0 0 1 .652 0l1.178.95l1.483-.357a.516.516 0 0 1 .588.276l.64 1.355l1.494.307c.24.05.411.258.407.498l-.027 1.492l1.211.91a.492.492 0 0 1 .145.621L14.26 7.5l.686 1.333a.49.49 0 0 1-.145.62l-.726.547h-2.478a.506.506 0 0 1-.512-.5a2.47 2.47 0 0 0-.861-1.87c.521-.363.86-.958.86-1.63c0-1.105-.916-2-2.047-2H1.399l-.015-.856a.503.503 0 0 1 .407-.498l1.495-.307l.639-1.355a.516.516 0 0 1 .588-.276l1.483.357l1.178-.95Z"/><path fill="currentColor" d="M.926 5L.2 5.546a.492.492 0 0 0-.145.621L.74 7.5L.055 8.833a.492.492 0 0 0 .145.62L.926 10h2.99V5H.926Z"/><path fill="currentColor" d="m1.4 11l-.016.856a.503.503 0 0 0 .407.498l1.495.307l.639 1.355c.103.218.35.334.588.276l1.483-.357l1.178.95a.52.52 0 0 0 .652 0l1.178-.95l1.483.357a.516.516 0 0 0 .588-.276l.64-1.355l1.494-.307a.504.504 0 0 0 .407-.498L13.6 11h-2.005c-.848 0-1.536-.672-1.536-1.5S9.372 8 8.524 8H4.94v2h2.048v1H1.399Zm7.636-4H4.94V5h4.096c.565 0 1.024.448 1.024 1S9.6 7 9.036 7Z"/>`
	safariOutlineInnerSVG                    = `<path fill="currentColor" d="m3.5 11.5l-.429-.257a.5.5 0 0 0 .686.686L3.5 11.5Zm3-5l-.257-.429l-.107.065l-.065.107l.429.257Zm5-3l.429.257a.5.5 0 0 0-.686-.686l.257.429Zm-3 5l.257.429l.107-.065l.065-.107L8.5 8.5Zm5.5-1A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM3.929 11.757l3-5l-.858-.514l-3 5l.858.514ZM6.757 6.93l5-3l-.514-.858l-5 3l.514.858Zm4.314-3.686l-3 5l.858.514l3-5l-.858-.514ZM8.243 8.07l-5 3l.514.858l5-3l-.514-.858Z"/>`
	safariSolidInnerSVG                      = `<path fill="currentColor" d="m4.958 10.042l1.906-3.178l3.178-1.906l-1.906 3.178l-3.178 1.906Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5A7.5 7.5 0 0 1 7.5 0A7.5 7.5 0 0 1 15 7.5A7.5 7.5 0 0 1 7.5 15A7.5 7.5 0 0 1 0 7.5Zm11.929-3.743a.5.5 0 0 0-.686-.686L6.136 6.136L3.07 11.243a.5.5 0 0 0 .686.686l5.107-3.065l3.065-5.107Z" clip-rule="evenodd"/>`
	safeOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M2.5 3v3m0 2v3m0 2.5V15m10-1.5V15m-3-5.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-8-9h12a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1Z"/>`
	safeSolidInnerSVG                        = `<path fill="currentColor" d="M8 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v11A1.5 1.5 0 0 0 1.5 14H2v1h1v-1h9v1h1v-1h.5a1.5 1.5 0 0 0 1.5-1.5v-11A1.5 1.5 0 0 0 13.5 0h-12ZM2 3v3h1V3H2Zm7.5 2a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM2 11V8h1v3H2Z" clip-rule="evenodd"/>`
	saveOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M4.5 14.5v-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v3m3 0h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1h8.586a1 1 0 0 1 .707.293l3.414 3.414a1 1 0 0 1 .293.707V13.5a1 1 0 0 1-1 1Z"/>`
	saveSolidInnerSVG                        = `<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h8.586a1.5 1.5 0 0 1 1.06.44l3.415 3.414A1.5 1.5 0 0 1 15 4.914V13.5a1.5 1.5 0 0 1-1.5 1.5H11v-3.5A1.5 1.5 0 0 0 9.5 10h-4A1.5 1.5 0 0 0 4 11.5V15H1.5A1.5 1.5 0 0 1 0 13.5v-12Z"/><path fill="currentColor" d="M5 15h5v-3.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5V15Z"/>`
	scanOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M.5 5V2.5a2 2 0 0 1 2-2H5m5 0h2.5a2 2 0 0 1 2 2V5m-14 5v2.5a2 2 0 0 0 2 2H5m9.5-4.5v2.5a2 2 0 0 1-2 2H10m-8-7h11"/>`
	scanSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M2.5 1A1.5 1.5 0 0 0 1 2.5V5H0V2.5A2.5 2.5 0 0 1 2.5 0H5v1H2.5Zm10 0H10V0h2.5A2.5 2.5 0 0 1 15 2.5V5h-1V2.5A1.5 1.5 0 0 0 12.5 1Zm.5 7H2V7h11v1ZM0 12.5V10h1v2.5A1.5 1.5 0 0 0 2.5 14H5v1H2.5A2.5 2.5 0 0 1 0 12.5Zm14 0V10h1v2.5a2.5 2.5 0 0 1-2.5 2.5H10v-1h2.5a1.5 1.5 0 0 0 1.5-1.5Z" clip-rule="evenodd"/>`
	schoolOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m7.5 4.5l4 2v8h-8v-8l4-2Zm0 0V0M0 14.5h15m-13.5 0v-6h2m10 6v-6h-2m-5 6v-3h2v3m-1-14h3v2h-3m0 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`
	schoolSolidInnerSVG                      = `<path fill="currentColor" d="M7.5 8a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z"/><path fill="currentColor" fill-rule="evenodd" d="m12 6.191l-4-2V3h3V0H7v4.191l-4 2V8H1v6H0v1h6v-4h3v4h6v-1h-1V8h-2V6.191ZM13 14V9h-1v5h1ZM3 14H2V9h1v5Zm3-5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/><path fill="currentColor" d="M8 15v-3H7v3h1Z"/>`
	screenAltOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M2 14.5h11m-5.5-4v4m-7-13v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`
	screenAltSolidInnerSVG                   = `<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v8A1.5 1.5 0 0 0 1.5 11H7v3H2v1h11v-1H8v-3h5.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-12Z"/>`
	screenAltTwoOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="M2 14.5h11m-7.5-4v4m4-4v4m-9-13v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`
	screenAltTwoSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v8A1.5 1.5 0 0 0 1.5 11H5v3H2v1h11v-1h-3v-3h3.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-12ZM6 14v-3h3v3H6Z" clip-rule="evenodd"/>`
	screenOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M4 14.5h7M.5 2.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`
	screenSolidInnerSVG                      = `<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v8A1.5 1.5 0 0 0 1.5 12h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 1h-12ZM4 15h7v-1H4v1Z"/>`
	scribbleOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M1.5 5C3 2 7.3.5 6.5 2.5C5.5 5-.5 9.5 3 11c1.343.576 3.055.45 4.654-.05m0 0C10.222 10.145 12.5 8.377 12.5 7C12.5 4.5 9 5.5 8 9c-.206.722-.328 1.381-.346 1.95Zm0 0C7.584 13.133 9.032 13.983 13 12"/>`
	scribbleSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M5.06 2.334c-1.077.463-2.426 1.515-3.113 2.89l-.894-.448c.813-1.625 2.364-2.823 3.612-3.36c.316-.136.626-.236.912-.287c.278-.05.571-.061.833.017a.923.923 0 0 1 .645.619c.091.3.028.623-.09.92c-.284.708-.897 1.514-1.538 2.302c-.196.241-.396.482-.596.723a35.682 35.682 0 0 0-1.365 1.71c-.593.811-.972 1.501-1.04 2.034c-.032.247.007.434.103.588c.1.16.293.338.668.498c1.104.474 2.543.426 3.98.028a9.65 9.65 0 0 1 .342-1.705c.536-1.876 1.757-3.141 2.93-3.581c.583-.219 1.223-.254 1.743.053c.542.32.808.924.808 1.665c0 .48-.196.947-.483 1.367c-.29.424-.692.834-1.164 1.213c-.86.69-1.99 1.308-3.195 1.73c.02.326.086.581.186.77c.123.234.31.394.603.476c.313.088.77.092 1.418-.062c.643-.154 1.44-.456 2.411-.941l.448.894c-1.013.507-1.885.842-2.627 1.02c-.738.175-1.382.202-1.92.052a1.917 1.917 0 0 1-1.218-.972a2.711 2.711 0 0 1-.279-.946c-1.483.37-3.064.421-4.377-.142c-.5-.214-.885-.505-1.123-.888c-.242-.389-.3-.819-.246-1.244c.103-.81.63-1.683 1.225-2.497c.43-.59.939-1.201 1.425-1.786c.195-.235.386-.465.567-.688c.656-.805 1.168-1.499 1.385-2.042a.832.832 0 0 0 .06-.216a1.008 1.008 0 0 0-.343.015a3.27 3.27 0 0 0-.693.221Zm3.172 7.88c.952-.375 1.825-.876 2.495-1.414c.419-.336.745-.676.964-.996C11.91 7.48 12 7.209 12 7c0-.509-.171-.718-.317-.804c-.168-.099-.465-.134-.882.022c-.827.31-1.856 1.295-2.32 2.92a9.81 9.81 0 0 0-.249 1.077Z" clip-rule="evenodd"/>`
	sdCardOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M6.5 3v3m2-3v3m2-3v3m-8 8.5h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-7l-4 4v9a1 1 0 0 0 1 1Z"/>`
	sdCardSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M5.293 0H12.5A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V4.293L5.293 0ZM6 3v3h1V3H6Zm2 0v3h1V3H8Zm2 3V3h1v3h-1Z" clip-rule="evenodd"/>`
	searchCircleOutlineInnerSVG              = `<path fill="none" stroke="currentColor" d="m8.5 8.5l2 2M7 9.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm.5 5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	searchCircleSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM7 4a3 3 0 1 0 1.738 5.445l1.409 1.409l.707-.707l-1.409-1.409A3 3 0 0 0 7 4Z" clip-rule="evenodd"/>`
	searchOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m14.5 14.5l-4-4m-4 2a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`
	searchPropertyOutlineInnerSVG            = `<path fill="currentColor" d="M4.5 8.5H4a.5.5 0 0 0 .5.5v-.5Zm4 0V9a.5.5 0 0 0 .5-.5h-.5Zm0-2.5H9a.5.5 0 0 0-.146-.354L8.5 6Zm-2-2l.354-.354a.5.5 0 0 0-.708 0L6.5 4Zm-2 2l-.354-.354A.5.5 0 0 0 4 6h.5Zm10.354 8.146l-4-4l-.708.708l4 4l.708-.708ZM6.5 12A5.5 5.5 0 0 1 1 6.5H0A6.5 6.5 0 0 0 6.5 13v-1ZM12 6.5A5.5 5.5 0 0 1 6.5 12v1A6.5 6.5 0 0 0 13 6.5h-1ZM6.5 1A5.5 5.5 0 0 1 12 6.5h1A6.5 6.5 0 0 0 6.5 0v1Zm0-1A6.5 6.5 0 0 0 0 6.5h1A5.5 5.5 0 0 1 6.5 1V0Zm-2 9h4V8h-4v1ZM9 8.5V6H8v2.5h1Zm-.146-2.854l-2-2l-.708.708l2 2l.708-.708Zm-2.708-2l-2 2l.708.708l2-2l-.708-.708ZM4 6v2.5h1V6H4Z"/>`
	searchPropertySolidInnerSVG              = `<path fill="currentColor" d="M5 8V6.207l1.5-1.5l1.5 1.5V8H5Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6.5a6.5 6.5 0 1 1 11.436 4.23l3.418 3.416l-.707.707l-3.418-3.417A6.5 6.5 0 0 1 0 6.5Zm8.854-.854l-2-2a.5.5 0 0 0-.708 0l-2 2A.5.5 0 0 0 4 6v2.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.146-.354Z" clip-rule="evenodd"/>`
	searchSmallOutlineInnerSVG               = `<path fill="none" stroke="currentColor" d="m8.5 8.5l2 2M7 9.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`
	searchSmallSolidInnerSVG                 = `<path fill="currentColor" d="M7 4a3 3 0 1 0 1.738 5.445l1.408 1.409l.708-.708l-1.409-1.408A3 3 0 0 0 7 4Z"/>`
	searchSolidInnerSVG                      = `<path fill="currentColor" d="M6.5 0a6.5 6.5 0 1 0 4.23 11.436l3.416 3.418l.708-.708l-3.418-3.417A6.5 6.5 0 0 0 6.5 0Z"/>`
	sectionAddOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M4 .5H1.5a1 1 0 0 0-1 1V4M6 .5h3m2 0h2.5a1 1 0 0 1 1 1V4M.5 6v3m14-3v3m-14 2v2.5a1 1 0 0 0 1 1H4M14.5 11v2.5a1 1 0 0 1-1 1H11M7.5 4v7M4 7.5h7m-5 7h3"/>`
	sectionAddSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0H4v1H1.5a.5.5 0 0 0-.5.5V4H0V1.5ZM9 1H6V0h3v1Zm4.5 0H11V0h2.5A1.5 1.5 0 0 1 15 1.5V4h-1V1.5a.5.5 0 0 0-.5-.5ZM7 7V4h1v3h3v1H8v3H7V8H4V7h3ZM0 9V6h1v3H0Zm14 0V6h1v3h-1ZM0 13.5V11h1v2.5a.5.5 0 0 0 .5.5H4v1H1.5A1.5 1.5 0 0 1 0 13.5Zm14 0V11h1v2.5a1.5 1.5 0 0 1-1.5 1.5H11v-1h2.5a.5.5 0 0 0 .5-.5ZM9 15H6v-1h3v1Z" clip-rule="evenodd"/>`
	sectionRemoveOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M4 .5H1.5a1 1 0 0 0-1 1V4M6 .5h3m2 0h2.5a1 1 0 0 1 1 1V4M.5 6v3m14-3v3m-14 2v2.5a1 1 0 0 0 1 1H4M14.5 11v2.5a1 1 0 0 1-1 1H11m-7-7h7m-5 7h3"/>`
	sectionRemoveSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0H4v1H1.5a.5.5 0 0 0-.5.5V4H0V1.5ZM9 1H6V0h3v1Zm4.5 0H11V0h2.5A1.5 1.5 0 0 1 15 1.5V4h-1V1.5a.5.5 0 0 0-.5-.5ZM0 9V6h1v3H0Zm14 0V6h1v3h-1Zm-3-1H4V7h7v1ZM0 13.5V11h1v2.5a.5.5 0 0 0 .5.5H4v1H1.5A1.5 1.5 0 0 1 0 13.5Zm14 0V11h1v2.5a1.5 1.5 0 0 1-1.5 1.5H11v-1h2.5a.5.5 0 0 0 .5-.5ZM9 15H6v-1h3v1Z" clip-rule="evenodd"/>`
	sendDownOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="m7.5 13.5l-4-4m4 4l4-4m-4 4V3M14 1.5H1"/>`
	sendDownSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v1H1V1Zm7 2v9.293l3.146-3.147l.708.708L7.5 14.207L3.146 9.854l.708-.708L7 12.293V3h1Z" clip-rule="evenodd"/>`
	sendLeftOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="m1.5 7.5l4-4m-4 4l4 4m-4-4H12m1.5 6.5V1"/>`
	sendLeftSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M13 1h1v13h-1V1ZM2.707 8l3.147 3.146l-.708.708L.793 7.5l4.353-4.354l.708.708L2.707 7H12v1H2.707Z" clip-rule="evenodd"/>`
	sendOutlineInnerSVG                      = `<path fill="currentColor" d="m14.5.5l.46.197a.5.5 0 0 0-.657-.657L14.5.5Zm-14 6l-.197-.46a.5.5 0 0 0-.06.889L.5 6.5Zm8 8l-.429.257a.5.5 0 0 0 .889-.06L8.5 14.5ZM14.303.04l-14 6l.394.92l14-6l-.394-.92ZM.243 6.93l5 3l.514-.858l-5-3l-.514.858ZM5.07 9.757l3 5l.858-.514l-3-5l-.858.514Zm3.889 4.94l6-14l-.92-.394l-6 14l.92.394ZM14.146.147l-9 9l.708.707l9-9l-.708-.708Z"/>`
	sendRightOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="m13.5 7.5l-4 4m4-4l-4-4m4 4H3M1.5 1v13"/>`
	sendRightSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1 14V1h1v13H1ZM9.854 3.146L14.207 7.5l-4.353 4.354l-.708-.708L12.293 8H3V7h9.293L9.146 3.854l.708-.708Z" clip-rule="evenodd"/>`
	sendSolidInnerSVG                        = `<path fill="currentColor" d="M14.954.71a.5.5 0 0 1-.1.144L5.4 10.306l2.67 4.451a.5.5 0 0 0 .889-.06L14.954.71ZM4.694 9.6L.243 6.928a.5.5 0 0 1 .06-.889L14.293.045a.5.5 0 0 0-.146.101L4.694 9.6Z"/>`
	sendUpOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m7.5 1.5l4 4m-4-4l-4 4m4-4V12M1 13.5h13"/>`
	sendUpSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="m7.5.793l4.354 4.353l-.708.708L8 2.707V12H7V2.707L3.854 5.854l-.708-.708L7.5.793ZM14 13v1H1v-1h13Z" clip-rule="evenodd"/>`
	serversOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M13.5 5.5h-12m12 0a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1m12 0a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1m-12-4a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1m12 0h-12m12 0a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1m.5-6h3m-3 4h3m-3 4h3"/>`
	serversSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h12A1.5 1.5 0 0 1 15 2.5v2c0 .384-.144.735-.382 1c.238.265.382.616.382 1v2c0 .384-.144.735-.382 1c.238.265.382.616.382 1v2a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-2c0-.384.144-.735.382-1A1.494 1.494 0 0 1 0 8.5v-2c0-.384.144-.735.382-1A1.494 1.494 0 0 1 0 4.5v-2ZM2 4h3V3H2v1Zm3 4H2V7h3v1Zm-3 4h3v-1H2v1Z" clip-rule="evenodd"/>`
	shareOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M4.5 7.5h4M11 4L8.5 7.495L11 11m3.5-8.501a2 2 0 0 1-4 0a2 2 0 0 1 4 0Zm0 9.993a2 2 0 0 1-4 0a2 2 0 0 1 4 0Zm-10-4.997a2 2 0 0 1-4 0a2 2 0 0 1 4 0Z"/>`
	shareSolidInnerSVG                       = `<path fill="currentColor" d="M10 2.499a2.5 2.5 0 0 1 5 0a2.5 2.5 0 0 1-3.566 2.26L9.131 7.52l2.038 2.858A2.5 2.5 0 0 1 15 12.493a2.5 2.5 0 1 1-4.559-1.417L8.246 8H4.949A2.501 2.501 0 0 1 0 7.495A2.5 2.5 0 0 1 4.95 7h3.312l2.37-2.84A2.488 2.488 0 0 1 10 2.499Z"/>`
	shieldOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 4.5l7-4l7 4v.72a9.651 9.651 0 0 1-7 9.28a9.651 9.651 0 0 1-7-9.28V4.5Z"/>`
	shieldSolidInnerSVG                      = `<path fill="currentColor" d="M7.748.066a.5.5 0 0 0-.496 0l-7 4A.5.5 0 0 0 0 4.5v.72a10.15 10.15 0 0 0 7.363 9.76a.5.5 0 0 0 .274 0A10.152 10.152 0 0 0 15 5.22V4.5a.5.5 0 0 0-.252-.434l-7-4Z"/>`
	shieldTickOutlineInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M4 7.5L7 10l4-5M7.5.5l-7 4v.72a9.651 9.651 0 0 0 7 9.28a9.651 9.651 0 0 0 7-9.28V4.5l-7-4Z"/>`
	shieldTickSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7.252.066a.5.5 0 0 1 .496 0l7 4A.5.5 0 0 1 15 4.5v.72a10.15 10.15 0 0 1-7.363 9.76a.5.5 0 0 1-.274 0A10.152 10.152 0 0 1 0 5.22V4.5a.5.5 0 0 1 .252-.434l7-4Zm-.18 10.645l4.318-5.399l-.78-.624l-3.682 4.601L4.32 7.116l-.64.768l3.392 2.827Z" clip-rule="evenodd"/>`
	shieldXOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m4.5 4.5l6 6m-6 0l6-6m-3-4l-7 4v.72a9.651 9.651 0 0 0 7 9.28a9.651 9.651 0 0 0 7-9.28V4.5l-7-4Z"/>`
	shieldXSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.252.066a.5.5 0 0 1 .496 0l7 4A.5.5 0 0 1 15 4.5v.72a10.15 10.15 0 0 1-7.363 9.76a.5.5 0 0 1-.274 0A10.152 10.152 0 0 1 0 5.22V4.5a.5.5 0 0 1 .252-.434l7-4Zm2.894 10.788L7.5 8.207l-2.646 2.647l-.708-.707L6.793 7.5L4.146 4.854l.708-.708L7.5 6.793l2.646-2.647l.708.708L8.207 7.5l2.647 2.646l-.707.708Z" clip-rule="evenodd"/>`
	shopOutlineInnerSVG                      = `<path fill="currentColor" d="M1.5 2.5V2a.5.5 0 0 0-.49.402l.49.098Zm12 0l.49-.098A.5.5 0 0 0 13.5 2v.5Zm1 5V8a.5.5 0 0 0 .49-.598l-.49.098Zm-14 0l-.49-.098A.5.5 0 0 0 .5 8v-.5Zm3 3H3v.5h.5v-.5Zm8 0v.5h.5v-.5h-.5ZM0 15h15v-1H0v1Zm1-7.5v7h1v-7H1Zm12 0v7h1v-7h-1ZM1.5 3h12V2h-12v1Zm11.51-.402l1 5l.98-.196l-1-5l-.98.196ZM14.5 7H.5v1h14V7ZM.99 7.598l1-5l-.98-.196l-1 5l.98.196ZM1 1h13V0H1v1Zm2 6.5v3h1v-3H3Zm.5 3.5h8v-1h-8v1Zm8.5-.5v-3h-1v3h1Z"/>`
	shopSolidInnerSVG                        = `<path fill="currentColor" d="M14 1H1V0h13v1ZM1.01 2.402A.5.5 0 0 1 1.5 2h12a.5.5 0 0 1 .49.402l1 5A.5.5 0 0 1 14.5 8H.5a.5.5 0 0 1-.49-.598l1-5ZM1 9v5H0v1h15v-1h-1V9h-2v2H3V9H1Z"/><path fill="currentColor" d="M4 9h7v1H4V9Z"/>`
	signOutlineInnerSVG                      = `<path fill="currentColor" d="M4.5 6.5V6a.5.5 0 0 0-.5.5h.5Zm10 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 6v.5a.5.5 0 0 0 .5-.5h-.5Zm-10 0H4a.5.5 0 0 0 .5.5v-.5ZM1 1v14h1V1H1ZM0 4h15V3H0v1Zm4.5 3h10V6h-10v1Zm9.5-.5v6h1v-6h-1Zm.5 5.5h-10v1h10v-1Zm-9.5.5v-6H4v6h1Zm1-9v3h1v-3H6Zm6 0v3h1v-3h-1Z"/>`
	signSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M1 1v2H0v1h1v11h1V4h4v2H4.5a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5H13V4h2V3H2V1H1Zm6 5V4h5v2H7Z" clip-rule="evenodd"/>`
	signinOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m10.5 7.5l-3 3.25m3-3.25l-3-3m3 3H1m6-6h6.5v12H7"/>`
	signinSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7 1h7v13H7v-1h6V2H7V1Zm.854 3.146l3.34 3.34l-3.327 3.603l-.734-.678L9.358 8H1V7h8.293L7.146 4.854l.708-.708Z" clip-rule="evenodd"/>`
	simOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M4 5.5h3.5V12M4 8.5h2m-2 3h2m3-6h2m-2 3h2m-2 3h2m1.5 3h-10a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1h7l4 4v9a1 1 0 0 1-1 1Z"/>`
	simSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M9.707 0H2.5A1.5 1.5 0 0 0 1 1.5v12A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4.293L9.707 0ZM7 6H4V5h4v7H7V6ZM6 9H4V8h2v1Zm-2 3h2v-1H4v1Zm7-6H9V5h2v1ZM9 9h2V8H9v1Zm2 3H9v-1h2v1Z" clip-rule="evenodd"/>`
	simohamedOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M12.5 5.5v3a6 6 0 0 1-1.2 3.6l-.6.8a4 4 0 0 1-6.4 0l-.6-.8a6 6 0 0 1-1.2-3.6v-3m10 0a5 5 0 0 0-10 0m10 0l-1.121-1.121A3 3 0 0 0 9.257 3.5H5.743a3 3 0 0 0-2.122.879L2.5 5.5m2.5 3h1m3 0h1m-4-2H2.5A1.562 1.562 0 0 0 .985 8.44l.151.605A1.921 1.921 0 0 0 3 10.5m6-4h3.5c1.016 0 1.761.955 1.515 1.94l-.151.605A1.921 1.921 0 0 1 12 10.5M3 11l.837.502a7 7 0 0 0 3.602.998h.122a7 7 0 0 0 3.602-.998L12 11m-6.5 1l1-1h2l1 1"/>`
	simohamedSolidInnerSVG                   = `<path fill="currentColor" d="M6 9H5V8h1v1Zm4 0H9V8h1v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M2 5.5a5.5 5.5 0 1 1 11 0v.56a2.063 2.063 0 0 1 1.5 2.502l-.151.604a2.422 2.422 0 0 1-1.825 1.777A6.5 6.5 0 0 1 11.7 12.4l-.6.8a4.5 4.5 0 0 1-7.2 0l-.6-.8a6.5 6.5 0 0 1-.823-1.457A2.422 2.422 0 0 1 .65 9.166L.5 8.562A2.063 2.063 0 0 1 2 6.06V5.5ZM3 7h3V6H3v-.293l.975-.975A2.5 2.5 0 0 1 5.743 4h3.514a2.5 2.5 0 0 1 1.768.732l.975.975V6H9v1h3v1.5a5.5 5.5 0 0 1-.455 2.19l-.64.384c-.35.21-.718.386-1.098.526l-1.1-1.1H6.293l-1.1 1.1a6.5 6.5 0 0 1-1.098-.526l-.64-.384A5.5 5.5 0 0 1 3 8.5V7Zm-1 .124c-.411.22-.653.702-.53 1.195l.151.605c.078.309.253.573.488.762A6.499 6.499 0 0 1 2 8.5V7.124Zm4.307 4.777c.372.066.75.099 1.132.099h.122c.381 0 .76-.034 1.133-.1l-.401-.4H6.707l-.4.4Zm6.584-2.215c.235-.19.41-.453.488-.762l.15-.605A1.062 1.062 0 0 0 13 7.124V8.5c0 .4-.037.797-.11 1.186Z" clip-rule="evenodd"/>`
	skullOutlineInnerSVG                     = `<path fill="currentColor" d="M3.5 11.5H4v-.309l-.276-.138l-.224.447Zm-1.447-.724l.223-.447l-.223.447Zm10.894 0l-.223-.447l.223.447ZM11.5 11.5l-.224-.447l-.276.138v.309h.5Zm-4-2l.354-.354a.5.5 0 0 0-.708 0L7.5 9.5ZM2 9.882V6.5H1v3.382h1ZM13 6.5v3.382h1V6.5h-1Zm-9.276 4.553l-1.448-.724l-.447.895l1.447.723l.448-.894Zm9-.724l-1.448.724l.448.894l1.447-.723l-.447-.895ZM4 12.5v-1H3v1h1Zm7-1v1h1v-1h-1ZM9.5 14h-4v1h4v-1Zm1.5-1.5A1.5 1.5 0 0 1 9.5 14v1a2.5 2.5 0 0 0 2.5-2.5h-1Zm2-2.618a.5.5 0 0 1-.276.447l.447.895A1.5 1.5 0 0 0 14 9.882h-1ZM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0v1ZM3 12.5A2.5 2.5 0 0 0 5.5 15v-1A1.5 1.5 0 0 1 4 12.5H3Zm-1-6A5.5 5.5 0 0 1 7.5 1V0A6.5 6.5 0 0 0 1 6.5h1ZM1 9.882a1.5 1.5 0 0 0 .83 1.342l.446-.895A.5.5 0 0 1 2 9.882H1ZM4.5 8a.5.5 0 0 1-.5-.5H3A1.5 1.5 0 0 0 4.5 9V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 6 7.5H5ZM4.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 4.5 6v1Zm0-1A1.5 1.5 0 0 0 3 7.5h1a.5.5 0 0 1 .5-.5V6Zm6 2a.5.5 0 0 1-.5-.5H9A1.5 1.5 0 0 0 10.5 9V8Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 12 7.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 6v1Zm0-1A1.5 1.5 0 0 0 9 7.5h1a.5.5 0 0 1 .5-.5V6Zm-3.646 4.854l1-1l-.708-.708l-1 1l.708.708Zm.292-1l1 1l.708-.708l-1-1l-.708.708Z"/>`
	skullSolidInnerSVG                       = `<path fill="currentColor" d="M4 7.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm6 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 0 1 13 0v3.382a1.5 1.5 0 0 1-.83 1.342l-1.17.585v.691A2.5 2.5 0 0 1 9.5 15h-4A2.5 2.5 0 0 1 3 12.5v-.691l-1.17-.585A1.5 1.5 0 0 1 1 9.882V6.5ZM4.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm6 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM7.146 9.146a.5.5 0 0 1 .708 0l1 1l-.708.708l-.646-.647l-.646.647l-.708-.708l1-1Z" clip-rule="evenodd"/>`
	skypeOutlineInnerSVG                     = `<path fill="currentColor" d="m13.83 8.34l-.496-.066a.5.5 0 0 0 .08.346l.415-.28ZM6.66 1.17l-.28.415a.5.5 0 0 0 .344.081l-.065-.496ZM1.17 6.66l.496.065a.5.5 0 0 0-.08-.345l-.416.28Zm7.17 7.17l.279-.415a.5.5 0 0 0-.345-.08l.066.495Zm5.985-5.423c.039-.29.066-.593.066-.907h-1c0 .257-.022.513-.057.774l.99.133Zm.066-.907A6.892 6.892 0 0 0 7.5.609v1A5.892 5.892 0 0 1 13.391 7.5h1ZM7.5.609a7.07 7.07 0 0 0-.905.065l.129.992c.264-.034.52-.057.776-.057v-1Zm-.562.146A4.437 4.437 0 0 0 4.457 0v1c.712 0 1.374.216 1.923.585l.558-.83ZM4.457 0A4.456 4.456 0 0 0 0 4.457h1A3.456 3.456 0 0 1 4.457 1V0ZM0 4.457c0 .918.279 1.772.755 2.481l.83-.558A3.436 3.436 0 0 1 1 4.457H0Zm.675 2.137a6.88 6.88 0 0 0-.066.906h1c0-.257.022-.513.057-.775l-.991-.131ZM.609 7.5A6.891 6.891 0 0 0 7.5 14.391v-1A5.891 5.891 0 0 1 1.609 7.5h-1ZM7.5 14.391c.314 0 .616-.027.906-.066l-.132-.99a5.888 5.888 0 0 1-.774.056v1Zm.561-.146c.71.477 1.564.755 2.482.755v-1a3.439 3.439 0 0 1-1.924-.585l-.558.83Zm2.482.755A4.457 4.457 0 0 0 15 10.543h-1A3.457 3.457 0 0 1 10.543 14v1ZM15 10.543c0-.918-.28-1.772-.756-2.481l-.83.558c.37.55.586 1.21.586 1.924h1ZM7 8h1V7H7v1Zm2.947-2.724A2.309 2.309 0 0 0 7.882 4v1c.496 0 .949.28 1.17.724l.895-.448ZM9 9a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2H9ZM8 8a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2v1Zm-.882 2c-.496 0-.95-.28-1.17-.724l-.895.448A2.309 2.309 0 0 0 7.118 11v-1ZM5 6a2 2 0 0 0 2 2V7a1 1 0 0 1-1-1H5Zm1 0a1 1 0 0 1 1-1V4a2 2 0 0 0-2 2h1Zm2 4h-.882v1H8v-1Zm-.118-6H7v1h.882V4Z"/>`
	skypeSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M0 4.457A4.456 4.456 0 0 1 6.778.652A6.46 6.46 0 0 1 7.5.61a6.892 6.892 0 0 1 6.891 6.89c0 .248-.017.49-.043.723a4.457 4.457 0 0 1-6.126 6.125a6.477 6.477 0 0 1-.722.043A6.891 6.891 0 0 1 .609 7.5c0-.248.017-.49.043-.723A4.435 4.435 0 0 1 0 4.457ZM6 6a1 1 0 0 1 1-1h.882c.496 0 .949.28 1.17.724l.895-.448A2.309 2.309 0 0 0 7.882 4H7a2 2 0 1 0 0 4h1a1 1 0 0 1 0 2h-.882c-.496 0-.95-.28-1.17-.724l-.895.448A2.309 2.309 0 0 0 7.118 11H8a2 2 0 1 0 0-4H7a1 1 0 0 1-1-1Z" clip-rule="evenodd"/>`
	slackOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M10.5 7.5V6A1.5 1.5 0 1 1 12 7.5h-1.5Zm0 0h-3m3 0V2a1.5 1.5 0 1 0-3 0v5.5m0 0v-3m0 3H2a1.5 1.5 0 1 1 0-3h5.5m0 3H13a1.5 1.5 0 0 1 0 3H7.5m0-3v3m0-3h-3m3 0V13a1.5 1.5 0 0 1-3 0V7.5m3-3V3A1.5 1.5 0 1 0 6 4.5h1.5Zm0 6H9A1.5 1.5 0 1 1 7.5 12v-1.5Zm-3-3V9A1.5 1.5 0 1 1 3 7.5h1.5Z"/>`
	slackSolidInnerSVG                       = `<path fill="currentColor" d="M10.385 6.923H8.077v-5.77a1.154 1.154 0 0 1 2.308 0v5.77Zm2.307 0H11.54V5.77a1.154 1.154 0 1 1 1.153 1.154Zm1.154 1.154h-5.77v2.308h5.77a1.154 1.154 0 0 0 0-2.308Zm-5.769 4.615V11.54H9.23a1.154 1.154 0 1 1-1.154 1.153ZM1.154 4.615h5.77v2.308h-5.77a1.154 1.154 0 0 1 0-2.308Zm5.769-2.307v1.154H5.77a1.154 1.154 0 1 1 1.154-1.154ZM1.154 9.23c0-.636.516-1.153 1.154-1.153h1.154V9.23a1.154 1.154 0 0 1-2.308 0Zm3.461 4.616v-5.77h2.308v5.77a1.154 1.154 0 0 1-2.308 0Z"/>`
	snapchatOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M7.5 1c-1.155 0-2.174.412-2.894 1.281c-.642.775-1.006 2.35-1.066 3.666l-.073.01l-.022.004a8.68 8.68 0 0 0-.368.059c-.465.089-1.346.326-1.543 1.277c-.093.445.011.833.247 1.134c.211.269.497.429.708.53c.09.041.181.08.274.117c-.21.584-.579 1.184-.987 1.728c-.382.508-.28 1.153-.083 1.573c.197.421.57.402 1.192.43c.352.015.722.051 1.09.12c.166.03.362.098.606.2c.142.06.283.123.423.187c.113.052.235.106.374.167c.573.25 1.276.517 2.056.517s1.483-.267 2.055-.517c.14-.06.26-.115.375-.167l.025-.012c.135-.06.26-.117.398-.174c.243-.103.44-.17.606-.201a7.951 7.951 0 0 1 1.09-.12c.622-.028 1.104-.009 1.303-.43c.197-.42.298-1.065-.084-1.573c-.406-.54-.772-1.136-.983-1.716a5.24 5.24 0 0 0 .305-.127c.216-.098.518-.261.73-.543c.245-.326.315-.739.175-1.184c-.28-.886-1.092-1.122-1.568-1.216a6.857 6.857 0 0 0-.355-.058l-.012-.002l-.056-.009c-.065-1.234-.41-2.795-1.036-3.581C9.695 1.485 8.682 1 7.5 1Z"/>`
	snapchatSolidInnerSVG                    = `<path fill="currentColor" d="M7.5 0C6.249 0 5.144.476 4.365 1.479c-.696.893-1.09 2.712-1.155 4.23l-.08.012l-.022.003c-.117.018-.256.04-.4.068c-.504.103-1.458.376-1.672 1.474c-.1.513.013.961.268 1.308c.23.31.54.495.768.61c.097.05.196.095.296.137c-.228.673-.627 1.366-1.069 1.993c-.414.587-.304 1.33-.09 1.816c.213.486.617.464 1.291.495c.382.018.783.06 1.181.139c.18.035.393.114.657.232c.153.07.306.14.458.215c.123.06.254.123.405.193c.62.288 1.383.596 2.227.596c.845 0 1.607-.308 2.227-.596c.151-.07.282-.133.406-.193l.027-.014a16.8 16.8 0 0 1 .43-.201c.265-.118.477-.196.657-.232a8.09 8.09 0 0 1 1.18-.139c.675-.031 1.198-.01 1.413-.495c.213-.485.323-1.229-.09-1.816c-.44-.623-.837-1.31-1.066-1.98a4.72 4.72 0 0 0 .33-.146c.234-.113.562-.301.79-.626c.267-.377.342-.853.19-1.366c-.302-1.023-1.182-1.295-1.697-1.403a6.404 6.404 0 0 0-.385-.068l-.014-.002l-.06-.01c-.07-1.424-.443-3.225-1.123-4.132C9.878.56 8.78 0 7.5 0Z"/>`
	snesOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M3.5 7v3M2 8.5h3m6 0h1m-1-2h1m-3 2h1m-1-2h1m-8.5-3h12a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`
	snesSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M0 4.5A1.5 1.5 0 0 1 1.5 3h12A1.5 1.5 0 0 1 15 4.5v6a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 10.5v-6ZM12 7h-1V6h1v1Zm-9 3V9H2V8h1V7h1v1h1v1H4v1H3Zm8-1h1V8h-1v1ZM9 9h1V8H9v1Zm1-2H9V6h1v1Z" clip-rule="evenodd"/>`
	sortAlphabeticallyOutlineInnerSVG        = `<path fill="currentColor" d="m3.5 14.5l-.354.354a.5.5 0 0 0 .708 0L3.5 14.5Zm6-6V8H9v.5h.5Zm0 6H9v.5h.5v-.5Zm-5.646.354l3-3l-.708-.708l-3 3l.708.708Zm0-.708l-3-3l-.708.708l3 3l.708-.708ZM3 0v14.5h1V0H3Zm6.5 9H12V8H9.5v1Zm2.5 2H9.5v1H12v-1Zm-2 .5v-3H9v3h1Zm3-1.5a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2h-1Zm-1-1a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2v1Zm0 5H9.5v1H12v-1Zm-2 .5v-3H9v3h1Zm3-1.5a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2h-1Zm-1-1a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2v1Zm-2-5V2.5H9V7h1Zm3-4.5V7h1V2.5h-1ZM11.5 1A1.5 1.5 0 0 1 13 2.5h1A2.5 2.5 0 0 0 11.5 0v1ZM10 2.5A1.5 1.5 0 0 1 11.5 1V0A2.5 2.5 0 0 0 9 2.5h1ZM9.5 5h4V4h-4v1Z"/>`
	sortAlphabeticallySolidInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M3 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708L3 13.293ZM11.5 1A1.5 1.5 0 0 0 10 2.5V4h3V2.5A1.5 1.5 0 0 0 11.5 1ZM13 5v2h1V2.5a2.5 2.5 0 0 0-5 0V7h1V5h3ZM9 8h3a2 2 0 0 1 1.323 3.5A2 2 0 0 1 12 15H9V8Zm3 3a1 1 0 1 0 0-2h-2v2h2Zm-2 1h2a1 1 0 1 1 0 2h-2v-2Z" clip-rule="evenodd"/>`
	sortDownOutlineInnerSVG                  = `<path fill="currentColor" d="m3.5 14.5l-.354.354a.5.5 0 0 0 .708 0L3.5 14.5Zm.354.354l3-3l-.708-.708l-3 3l.708.708Zm0-.708l-3-3l-.708.708l3 3l.708-.708ZM3 0v14.5h1V0H3Zm6 4h6V3H9v1Zm0 4h4V7H9v1Zm0 4h2v-1H9v1Z"/>`
	sortDownSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M3 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708L3 13.293ZM15 4H9V3h6v1Zm-2 4H9V7h4v1Zm-2 4H9v-1h2v1Z" clip-rule="evenodd"/>`
	sortHighToLowOutlineInnerSVG             = `<path fill="currentColor" d="M11.5.5h.5a.5.5 0 0 0-.5-.5v.5Zm1 11l.474.158l-.474-.158Zm.482-1.446l-.474-.158l.474.158ZM4.5 14.5l-.354.354a.5.5 0 0 0 .708 0L4.5 14.5ZM10 1h1.5V0H10v1Zm1-.5v6h1v-6h-1ZM10 7h3V6h-3v1Zm1.862 1H11v1h.862V8Zm1.112 3.658l.482-1.446l-.948-.316l-.482 1.446l.948.316ZM11 12h1.5v-1H11v1Zm.974 2.658l1-3l-.948-.316l-1 3l.948.316ZM9 10a2 2 0 0 0 2 2v-1a1 1 0 0 1-1-1H9Zm2-2a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1V8Zm.862 1a.68.68 0 0 1 .646.896l.948.316A1.68 1.68 0 0 0 11.862 8v1Zm-7.008 5.854l3-3l-.708-.708l-3 3l.708.708Zm0-.708l-3-3l-.708.708l3 3l.708-.708ZM4 0v14.5h1V0H4Z"/>`
	sortHighToLowSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M4 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708L4 13.293ZM11 1h-1V0h1.5a.5.5 0 0 1 .5.5V6h1v1h-3V6h1V1Zm-2 9a2 2 0 0 1 2-2h.862a1.68 1.68 0 0 1 1.594 2.212l-1.482 4.446l-.948-.316l.78-2.342H11a2 2 0 0 1-2-2Zm3.14 1l.368-1.104A.68.68 0 0 0 11.862 9H11a1 1 0 1 0 0 2h1.14Z" clip-rule="evenodd"/>`
	sortLowToHighOutlineInnerSVG             = `<path fill="currentColor" d="M11.5.5h.5a.5.5 0 0 0-.5-.5v.5Zm-7 0l.354-.354a.5.5 0 0 0-.708 0L4.5.5ZM10 1h1.5V0H10v1Zm1-.5v6h1v-6h-1ZM10 7h3V6h-3v1Zm1.578 1H11v1h.578V8Zm1.396 3.658l.393-1.176l-.95-.317l-.391 1.177l.948.316ZM11 12h1.5v-1H11v1Zm.974 2.658l1-3l-.948-.316l-1 3l.948.316ZM9 10a2 2 0 0 0 2 2v-1a1 1 0 0 1-1-1H9Zm2-2a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1V8Zm.578 1a.885.885 0 0 1 .84 1.165l.949.317A1.885 1.885 0 0 0 11.578 8v1ZM4.146.146l-3 3l.708.708l3-3l-.708-.708Zm0 .708l3 3l.708-.708l-3-3l-.708.708ZM4 .5V15h1V.5H4Z"/>`
	sortLowToHighSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M4.146.146a.5.5 0 0 1 .708 0l3 3l-.708.708L5 1.707V15H4V1.707L1.854 3.854l-.708-.708l3-3ZM11 1h-1V0h1.5a.5.5 0 0 1 .5.5V6h1v1h-3V6h1V1Zm-2 9a2 2 0 0 1 2-2h.578a1.885 1.885 0 0 1 1.789 2.482l-1.393 4.176l-.948-.316l.78-2.342H11a2 2 0 0 1-2-2Zm3.14 1l.278-.835A.885.885 0 0 0 11.578 9H11a1 1 0 1 0 0 2h1.14Z" clip-rule="evenodd"/>`
	sortReverseAlphabeticallyOutlineInnerSVG = `<path fill="currentColor" d="m3.5.5l.354-.354a.5.5 0 0 0-.708 0L3.5.5Zm6 8V8H9v.5h.5Zm0 6H9v.5h.5v-.5ZM3.146.146l-3 3l.708.708l3-3l-.708-.708Zm0 .708l3 3l.708-.708l-3-3l-.708.708ZM3 .5V15h1V.5H3ZM9.5 9H12V8H9.5v1Zm2.5 2H9.5v1H12v-1Zm-2 .5v-3H9v3h1Zm3-1.5a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2h-1Zm-1-1a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2v1Zm0 5H9.5v1H12v-1Zm-2 .5v-3H9v3h1Zm3-1.5a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2h-1Zm-1-1a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2v1Zm-2-5V2.5H9V7h1Zm3-4.5V7h1V2.5h-1ZM11.5 1A1.5 1.5 0 0 1 13 2.5h1A2.5 2.5 0 0 0 11.5 0v1ZM10 2.5A1.5 1.5 0 0 1 11.5 1V0A2.5 2.5 0 0 0 9 2.5h1ZM9.5 5h4V4h-4v1Z"/>`
	sortReverseAlphabeticallySolidInnerSVG   = `<path fill="currentColor" fill-rule="evenodd" d="M3.146.146a.5.5 0 0 1 .708 0l3 3l-.708.708L4 1.707V15H3V1.707L.854 3.854l-.708-.708l3-3ZM11.5 1A1.5 1.5 0 0 0 10 2.5V4h3V2.5A1.5 1.5 0 0 0 11.5 1ZM13 5v2h1V2.5a2.5 2.5 0 0 0-5 0V7h1V5h3ZM9 8h3a2 2 0 0 1 1.323 3.5A2 2 0 0 1 12 15H9V8Zm3 3a1 1 0 1 0 0-2h-2v2h2Zm-2 1h2a1 1 0 1 1 0 2h-2v-2Z" clip-rule="evenodd"/>`
	sortUpOutlineInnerSVG                    = `<path fill="currentColor" d="m3.5.5l.354-.354a.5.5 0 0 0-.708 0L3.5.5ZM3.146.146l-3 3l.708.708l3-3l-.708-.708Zm0 .708l3 3l.708-.708l-3-3l-.708.708ZM3 .5V15h1V.5H3ZM9 4h6V3H9v1Zm0 4h4V7H9v1Zm0 4h2v-1H9v1Z"/>`
	sortUpSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M3.146.146a.5.5 0 0 1 .708 0l3 3l-.708.708L4 1.707V15H3V1.707L.854 3.854l-.708-.708l3-3ZM15 4H9V3h6v1Zm-2 4H9V7h4v1Zm-2 4H9v-1h2v1Z" clip-rule="evenodd"/>`
	soundOffOutlineInnerSVG                  = `<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5v.5Zm0-5.996v.5h.138l.12-.071l-.258-.429Zm5-2.998H9a.5.5 0 0 0-.757-.429L8.5 1.5Zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492h-.5Zm2.94-7.763l-.354-.353l-.707.707l.354.354l.706-.708Zm2.12 3.534l.354.353l.707-.707l-.354-.353l-.707.707Zm.708-2.826l.353-.354l-.707-.707l-.353.353l.707.708Zm-3.535 2.119l-.354.353l.707.707l.354-.353l-.707-.707ZM3.5 9.994h-2v1h2v-1Zm-2 0a.499.499 0 0 1-.5-.5H0c0 .83.671 1.5 1.5 1.5v-1Zm-.5-.5V5.498H0v3.998h1Zm0-3.997c0-.276.223-.499.5-.499v-1c-.829 0-1.5.67-1.5 1.5h1Zm.5-.499h2v-1h-2v1Zm2.257-.071l5-2.998l-.514-.858l-5 2.998l.514.858ZM8 1.5v11.992h1V1.5H8Zm.757 11.563l-5-2.998l-.514.858l5 2.998l.514-.858Zm1.976-6.626l2.827 2.826l.707-.707l-2.828-2.827l-.707.708Zm2.828-.708l-2.828 2.827l.707.707l2.828-2.826l-.707-.708Z"/>`
	soundOffSolidInnerSVG                    = `<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5c-.829 0-1.5.67-1.5 1.5V9.5c0 .829.67 1.5 1.5 1.5h1.862l4.88 2.929A.5.5 0 0 0 9 13.5v-12Zm4.207 5.996l1.414 1.413l-.707.707L12.5 8.203l-1.414 1.413l-.707-.707l1.414-1.413l-1.414-1.413l.707-.707L12.5 6.789l1.415-1.413l.706.707l-1.414 1.413Z"/>`
	soundOnOutlineInnerSVG                   = `<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5v.5Zm0-5.996v.5h.138l.12-.071l-.258-.429Zm5-2.998H9a.5.5 0 0 0-.757-.429L8.5 1.5Zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492h-.5Zm5.353-9.348l-.353-.353l-.707.707l.354.354l.706-.708Zm-.706 5.996l-.354.354l.707.707l.353-.353l-.706-.708ZM3.5 9.994h-2v1h2v-1Zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5v-1Zm-.5-.5V5.498H0v3.998h1Zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5h1Zm.5-.499h2v-1h-2v1Zm2.257-.071l5-2.998l-.514-.858l-5 2.998l.514.858ZM8 1.5v11.992h1V1.5H8Zm.757 11.563l-5-2.998l-.514.858l5 2.998l.514-.858ZM13.5 4.498c-.353.354-.354.354-.354.353a.01.01 0 0 1-.002-.002l.003.003l.02.022a3.186 3.186 0 0 1 .386.597c.22.439.447 1.112.447 2.025h1c0-1.086-.272-1.911-.553-2.472a4.198 4.198 0 0 0-.39-.639a2.932 2.932 0 0 0-.181-.217l-.014-.015l-.005-.005l-.002-.002l-.001-.001l-.354.353Zm.5 2.998c0 .913-.228 1.586-.447 2.025a3.184 3.184 0 0 1-.386.597a.83.83 0 0 1-.02.022l-.003.003l.001-.001l.001-.001l.354.353c.353.354.354.354.354.353h.001l.002-.003l.005-.005l.014-.014l.043-.048c.035-.04.082-.097.137-.17c.11-.146.251-.36.391-.639c.28-.56.553-1.386.553-2.472h-1Z"/>`
	soundOnSolidInnerSVG                     = `<path fill="currentColor" d="M8.746 1.065A.5.5 0 0 1 9 1.5v12a.5.5 0 0 1-.757.429L3.362 11H1.5A1.5 1.5 0 0 1 0 9.5V5.497a1.5 1.5 0 0 1 1.5-1.499h1.862l4.88-2.927a.5.5 0 0 1 .504-.006Zm5.108 3.079l-.354-.353l-.707.707l.351.352l.003.002l.02.022a3.194 3.194 0 0 1 .386.597c.22.439.447 1.112.447 2.025c0 .913-.228 1.586-.447 2.025a3.19 3.19 0 0 1-.297.486a1.988 1.988 0 0 1-.11.133l-.002.003l-.351.35l.707.708l.354-.353l-.354-.354l.354.354l.001-.002l.002-.002l.005-.005l.014-.014l.043-.048c.035-.04.082-.097.137-.17c.11-.146.251-.36.391-.639c.28-.56.553-1.386.553-2.472s-.272-1.911-.553-2.472a4.19 4.19 0 0 0-.39-.639a2.89 2.89 0 0 0-.181-.217l-.014-.015l-.005-.005l-.002-.002c0-.001-.002-.002-.355.352l.354-.354Z"/>`
	spotifyOutlineInnerSVG                   = `<path fill="currentColor" d="M14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm-3.838 9.116c.986-.24 2.2-.345 3.392-.228c1.196.117 2.334.454 3.202 1.065l.575-.819c-1.053-.74-2.375-1.113-3.679-1.241a11.446 11.446 0 0 0-3.727.252l.237.971Zm-.336-2.124c3.446-.61 5.848-.297 7.839 1.132l.583-.813C9.45 6.661 6.732 6.374 3.152 7.008l.174.984Zm-.225-2.151c1.353-.478 3.003-.676 4.624-.544c1.623.133 3.18.595 4.364 1.4l.561-.828c-1.364-.927-3.099-1.426-4.843-1.568c-1.746-.143-3.539.067-5.039.597l.333.943Z"/>`
	spotifySolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm7.726-2.203c-1.621-.132-3.272.066-4.625.544l-.333-.943c1.5-.53 3.293-.74 5.04-.597c1.743.142 3.478.641 4.843 1.568l-.562.827c-1.185-.804-2.74-1.266-4.363-1.399Zm-4.4 2.695c3.446-.61 5.848-.297 7.839 1.132l.583-.813C9.45 6.661 6.732 6.374 3.152 7.008l.174.984Zm.336 2.124c.986-.24 2.2-.345 3.392-.228c1.196.117 2.334.454 3.202 1.065l.575-.819c-1.053-.74-2.375-1.113-3.68-1.241a11.446 11.446 0 0 0-3.726.252l.237.971Z" clip-rule="evenodd"/>`
	spreadsheetOutlineInnerSVG               = `<path fill="currentColor" d="M4.5 6.995H4v1h.5v-1Zm6 1h.5v-1h-.5v1Zm-6 2.505H4v.5h.5v-.5Zm6 0v.5h.5v-.5h-.5Zm-6-6.503H4v1h.5v-1Zm6 1h.5v-1h-.5v1Zm3-1.497h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-6 7.495h6v-1h-6v1ZM4.5 11h6v-1h-6v1Zm0-6.003h6v-1h-6v1Zm8 9.003h-10v1h10v-1ZM2 13.5v-12H1v12h1Zm11-10v10h1v-10h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1Zm2 3v6h1v-6H4Zm3 0v6h1v-6H7Zm3 0v6h1v-6h-1Z"/>`
	spreadsheetSolidInnerSVG                 = `<path fill="currentColor" d="M10 7.995V10H8V7.995h2Zm0-2.998v1.998H8V4.997h2Zm-3 0H5v1.998h2V4.997Zm0 2.998H5V10h2V7.995Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm10 2.497H4V11h7V3.997Z" clip-rule="evenodd"/>`
	squareOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M13.5.5h-12a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1Z"/>`
	squareSolidInnerSVG                      = `<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v12A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0h-12Z"/>`
	stackoverflowOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M2.5 9v5.5h10V9M4 12.5h7M4 10l7 .7M4.3 7.5L11 9M5.3 4.5l6.1 3.1M7 2l5 4.5M9.5.5l3.5 5"/>`
	stackoverflowSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="m12.59 5.787l-3.5-5l.82-.574l3.5 5l-.82.574Zm-.925 1.085l-5-4.5l.67-.744l5 4.5l-.67.744Zm-.491 1.174l-6.1-3.1l.453-.892l6.1 3.1l-.454.892Zm-.283 1.442l-6.7-1.5l.218-.976l6.7 1.5l-.218.976ZM2 9h1v5h9V9h1v6H2V9Zm8.95 2.197l-7-.7l.1-.995l7 .7l-.1.995ZM11 13H4v-1h7v1Z" clip-rule="evenodd"/>`
	stampOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M0 14.5h15m-8.5-8v3m2-3v3m-1-3a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-7 6v-1a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v1H.5Z"/>`
	stampSolidInnerSVG                       = `<path fill="currentColor" d="M4 3.5a3.5 3.5 0 1 1 5 3.163V9h3.5a2.5 2.5 0 0 1 2.5 2.5V13H0v-1.5A2.5 2.5 0 0 1 2.5 9H6V6.663A3.5 3.5 0 0 1 4 3.5ZM0 14v1h15v-1H0Z"/>`
	starCircleOutlineInnerSVG                = `<path fill="currentColor" d="m7.5 9.804l.242-.437a.5.5 0 0 0-.484 0l.242.437ZM5.337 11l-.494-.08a.5.5 0 0 0 .736.518L5.337 11Zm.413-2.533l.493.08a.5.5 0 0 0-.135-.429l-.358.35ZM4 6.674l-.075-.495a.5.5 0 0 0-.283.844L4 6.673Zm2.418-.37l.076.495a.5.5 0 0 0 .377-.282l-.453-.213ZM7.5 4l.453-.212a.5.5 0 0 0-.906 0L7.5 4Zm1.082 2.304l-.453.213a.5.5 0 0 0 .377.282l.076-.495Zm2.418.37l.358.349a.5.5 0 0 0-.283-.844L11 6.674ZM9.25 8.467l-.358-.349a.5.5 0 0 0-.135.43l.493-.08ZM9.663 11l-.242.438a.5.5 0 0 0 .736-.519L9.663 11ZM7.258 9.367l-2.163 1.195l.484.876l2.163-1.196l-.484-.875ZM5.83 11.08l.413-2.532l-.986-.161l-.414 2.532l.987.162Zm.278-2.962l-1.75-1.794l-.716.699l1.75 1.793l.716-.698Zm-2.033-.95l2.419-.37l-.151-.988l-2.418.37l.15.988Zm2.796-.651l1.082-2.305l-.906-.424l-1.081 2.304l.905.425Zm.176-2.305L8.13 6.517l.905-.425l-1.081-2.304l-.906.424ZM8.507 6.8l2.418.369l.15-.989l-2.418-.369l-.15.989Zm2.135-.475l-1.75 1.794l.716.698l1.75-1.793l-.716-.699ZM8.757 8.548l.413 2.533l.987-.162l-.414-2.532l-.986.16Zm1.148 2.014L7.742 9.367l-.484.875l2.163 1.196l.484-.876ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	starCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm7.5-4a.5.5 0 0 1 .453.288L8.92 5.85l2.155.33a.5.5 0 0 1 .282.843L9.784 8.636l.373 2.284a.5.5 0 0 1-.736.518L7.5 10.375l-1.921 1.063a.5.5 0 0 1-.736-.519l.373-2.283l-1.574-1.613a.5.5 0 0 1 .283-.844l2.154-.329l.968-2.062A.5.5 0 0 1 7.5 3.5Z" clip-rule="evenodd"/>`
	starOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 12.04l-4.326 2.275L4 9.497L.5 6.086l4.837-.703L7.5 1l2.163 4.383l4.837.703L11 9.497l.826 4.818L7.5 12.041Z"/>`
	starSmallOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 9.804L5.337 11l.413-2.533L4 6.674l2.418-.37L7.5 4l1.082 2.304l2.418.37l-1.75 1.793L9.663 11L7.5 9.804Z"/>`
	starSmallSolidInnerSVG                   = `<path fill="currentColor" d="M7.953 3.788a.5.5 0 0 0-.906 0L6.08 5.85l-2.154.33a.5.5 0 0 0-.283.843l1.574 1.613l-.373 2.284a.5.5 0 0 0 .736.518l1.92-1.063l1.921 1.063a.5.5 0 0 0 .736-.519l-.373-2.283l1.574-1.613a.5.5 0 0 0-.283-.844L8.921 5.85l-.968-2.062Z"/>`
	starSolidInnerSVG                        = `<path fill="currentColor" d="M7.948.779a.5.5 0 0 0-.896 0L5.005 4.926l-4.577.665a.5.5 0 0 0-.277.853l3.312 3.228l-.782 4.559a.5.5 0 0 0 .725.527L7.5 12.605l4.094 2.153a.5.5 0 0 0 .725-.527l-.782-4.56l3.312-3.227a.5.5 0 0 0-.277-.853l-4.577-.665L7.948.78Z"/>`
	stopCircleOutlineInnerSVG                = `<g fill="none" stroke="currentColor"><path d="M.5 7.5a7 7 0 1 1 14 0a7 7 0 0 1-14 0Z"/><path d="M9.5 5.5h-4v4h4v-4Z"/></g>`
	stopCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM5 5h5v5H5V5Z" clip-rule="evenodd"/>`
	stopOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M11.5 3.5h-8v8h8v-8Z"/>`
	stopSmallOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M9.5 5.5h-4v4h4v-4Z"/>`
	stopSmallSolidInnerSVG                   = `<path fill="currentColor" d="M10 5H5v5h5V5Z"/>`
	stopSolidInnerSVG                        = `<path fill="currentColor" d="M12 3H3v9h9V3Z"/>`
	stopwatchOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M7.5 5v3.5H10m-4-8h3m-1.5 2a6 6 0 1 0 0 12a6 6 0 0 0 0-12Z"/>`
	stopwatchSolidInnerSVG                   = `<path fill="currentColor" d="M9 1H6V0h3v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 2a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13ZM8 8V5H7v4h3V8H8Z" clip-rule="evenodd"/>`
	strikethroughOutlineInnerSVG             = `<path fill="none" stroke="currentColor" d="M3.5 10A3.5 3.5 0 0 0 7 13.5h1.487a3.013 3.013 0 0 0 3.013-3.013c0-1.205-.892-2.512-2-2.987M6.08 6.177A2.607 2.607 0 0 1 4.5 3.781A2.281 2.281 0 0 1 6.781 1.5H8A2.5 2.5 0 0 1 10.5 4M2 7.5h11"/>`
	strikethroughSolidInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M4 3.781A2.781 2.781 0 0 1 6.781 1H8a3 3 0 0 1 3 3h-1a2 2 0 0 0-2-2H6.781A1.78 1.78 0 0 0 5 3.781c0 .843.502 1.605 1.277 1.937l-.394.919A3.107 3.107 0 0 1 4 3.78ZM9.392 8H2V7h11v1h-2.01c.123.14.237.287.34.44c.405.602.67 1.326.67 2.047A3.513 3.513 0 0 1 8.487 14H7a4 4 0 0 1-4-4h1a3 3 0 0 0 3 3h1.487A2.513 2.513 0 0 0 11 10.487c0-.484-.182-1.017-.5-1.488c-.296-.44-.69-.797-1.108-.999Z" clip-rule="evenodd"/>`
	subscriptOutlineInnerSVG                 = `<path fill="currentColor" d="m12.5 14.5l-.354-.354A.5.5 0 0 0 12.5 15v-.5ZM15 14h-2.5v1H15v-1Zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792l.708.708ZM13.793 11H13.5v1h.293v-1Zm-.293 0a1.5 1.5 0 0 0-1.5 1.5h1a.5.5 0 0 1 .5-.5v-1Zm1.5 1.207C15 11.54 14.46 11 13.793 11v1c.114 0 .207.093.207.207h1Zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147l.706.707ZM1.9 13.8l9-12l-.8-.6l-9 12l.8.6Zm-.8-12l9 12l.8-.6l-9-12l-.8.6Z"/>`
	subscriptSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M6 8.333L1.9 13.8l-.8-.6l4.275-5.7L1.1 1.8l.8-.6L6 6.667L10.1 1.2l.8.6l-4.275 5.7l4.275 5.7l-.8.6L6 8.333Zm6 4.167a1.5 1.5 0 0 1 1.5-1.5h.293a1.207 1.207 0 0 1 .854 2.06l-.94.94H15v1h-2.5a.5.5 0 0 1-.354-.854l1.793-1.792a.207.207 0 0 0-.146-.354H13.5a.5.5 0 0 0-.5.5h-1Z" clip-rule="evenodd"/>`
	sunOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5 1.5v-1m0 13.99v-.998m6-5.997h1m-13 0h-1m2-4.996l-1-1m12 0l-1 1m-10 9.993l-1 1m12 0l-1-1m-2-4.997a2.999 2.999 0 0 1-3 2.998a2.999 2.999 0 1 1 3-2.998Z"/>`
	sunSolidInnerSVG                         = `<path fill="currentColor" d="M8 2V0H7v2h1Zm-4.793.498L1.5.792L.793 1.5L2.5 3.206l.707-.708Zm9.293.708L14.207 1.5L13.5.792l-1.707 1.706l.707.708Zm-5 .791a3.499 3.499 0 1 0 0 6.996a3.499 3.499 0 1 0 0-6.996ZM2 6.995H0v1h2v-1Zm13 0h-2v1h2v-1ZM1.5 14.199l1.707-1.707l-.707-.707l-1.707 1.706l.707.708Zm12.707-.708L12.5 11.785l-.707.707L13.5 14.2l.707-.708ZM8 14.99v-1.998H7v1.999h1Z"/>`
	superscriptOutlineInnerSVG               = `<path fill="currentColor" d="m12.5 3.5l-.354-.354A.5.5 0 0 0 12.5 4v-.5Zm1.793-1.793l-.354-.353l.354.353ZM15 3h-2.5v1H15V3Zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792l.708.708ZM13.793 0H13.5v1h.293V0ZM13.5 0A1.5 1.5 0 0 0 12 1.5h1a.5.5 0 0 1 .5-.5V0ZM15 1.207C15 .54 14.46 0 13.793 0v1c.114 0 .207.093.207.207h1Zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147l.706.707ZM1.1 1.8l9 12l.8-.6l-9-12l-.8.6Zm9-.6l-9 12l.8.6l9-12l-.8-.6Z"/>`
	superscriptSolidInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M12 1.5A1.5 1.5 0 0 1 13.5 0h.293a1.207 1.207 0 0 1 .854 2.06l-.94.94H15v1h-2.5a.5.5 0 0 1-.354-.854l1.793-1.792A.207.207 0 0 0 13.793 1H13.5a.5.5 0 0 0-.5.5h-1Zm-6.625 6L1.1 1.8l.8-.6L6 6.667L10.1 1.2l.8.6l-4.275 5.7l4.275 5.7l-.8.6L6 8.333L1.9 13.8l-.8-.6l4.275-5.7Z" clip-rule="evenodd"/>`
	svelteOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m9.625 8.357l-5.088 3.18m2.968-1.855a3.5 3.5 0 0 1-3.71-5.937l4.241-2.65A3.5 3.5 0 0 1 12.405 6.5M7.536 5.296a3.5 3.5 0 0 1 3.71 5.936l-4.24 2.65A3.5 3.5 0 0 1 2.614 8.5m2.8-1.88l5.09-3.179"/>`
	svelteSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M13.283 1.944a4.001 4.001 0 0 1-.27 4.622a4 4 0 0 1-1.503 5.09l-4.24 2.65A4 4 0 0 1 2.028 8.41a4 4 0 0 1 1.503-5.09L7.77.671a4 4 0 0 1 5.512 1.273ZM8.3 1.52a3 3 0 0 1 4.13 4.143a3.993 3.993 0 0 0-2.386-1.345l.724-.453l-.53-.848l-5.088 3.18l.53.847L7.8 5.72a3 3 0 1 1 3.18 5.088l-4.24 2.65a3 3 0 0 1-4.13-4.143a3.993 3.993 0 0 0 2.386 1.345l-.725.452l.53.848L9.89 8.78l-.53-.847l-2.12 1.325a3 3 0 0 1-3.18-5.089L8.3 1.52Z" clip-rule="evenodd"/>`
	svgOutlineInnerSVG                       = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-8 6V6H2v.5h.5Zm0 2H2V9h.5v-.5Zm2 0H5V8h-.5v.5Zm0 2v.5H5v-.5h-.5Zm2-1H6v.207l.146.147L6.5 9.5Zm1 1l-.354.354l.354.353l.354-.353L7.5 10.5Zm1-1l.354.354L9 9.707V9.5h-.5Zm2-3V6H10v.5h.5Zm0 4H10v.5h.5v-.5Zm2 0v.5h.5v-.5h-.5ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM5 6H2.5v1H5V6Zm-3 .5v2h1v-2H2ZM2.5 9h2V8h-2v1ZM4 8.5v2h1v-2H4Zm.5 1.5H2v1h2.5v-1ZM6 6v3.5h1V6H6Zm.146 3.854l1 1l.708-.708l-1-1l-.708.708Zm1.708 1l1-1l-.708-.708l-1 1l.708.708ZM9 9.5V6H8v3.5h1ZM13 6h-2.5v1H13V6Zm-3 .5v4h1v-4h-1Zm.5 4.5h2v-1h-2v1Zm2.5-.5v-2h-1v2h1Z"/>`
	svgSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM5 6H2v3h2v1H2v1h3V8H3V7h2V6Zm2 0H6v3.707l1.5 1.5l1.5-1.5V6H8v3.293l-.5.5l-.5-.5V6Zm3 0h3v1h-2v3h1V8.5h1V11h-3V6Z" clip-rule="evenodd"/>`
	tableOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M.5 4.5h14m-10-4v14m-3-14h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`
	tableSolidInnerSVG                       = `<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0H4v4H0V1.5ZM0 5v8.5A1.5 1.5 0 0 0 1.5 15H4V5H0Zm5 10h8.5a1.5 1.5 0 0 0 1.5-1.5V5H5v10ZM15 4V1.5A1.5 1.5 0 0 0 13.5 0H5v4h10Z"/>`
	tabletOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M6 11.5h3m-6.5 3h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`
	tabletSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h10A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM6 12h3v-1H6v1Z" clip-rule="evenodd"/>`
	tagOutlineInnerSVG                       = `<path fill="currentColor" d="m.5 7.5l-.354-.354a.5.5 0 0 0 0 .708L.5 7.5Zm7 7l-.354.354a.5.5 0 0 0 .708 0L7.5 14.5Zm7-7l.354.354A.5.5 0 0 0 15 7.5h-.5Zm-7-7V0a.5.5 0 0 0-.354.146L7.5.5ZM.146 7.854l7 7l.708-.708l-7-7l-.708.708Zm7.708 7l7-7l-.708-.708l-7 7l.708.708ZM15 7.5v-6h-1v6h1ZM13.5 0h-6v1h6V0ZM7.146.146l-7 7l.708.708l7-7l-.708-.708ZM15 1.5A1.5 1.5 0 0 0 13.5 0v1a.5.5 0 0 1 .5.5h1ZM10.5 5a.5.5 0 0 1-.5-.5H9A1.5 1.5 0 0 0 10.5 6V5Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 12 4.5h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 3v1Zm0-1A1.5 1.5 0 0 0 9 4.5h1a.5.5 0 0 1 .5-.5V3Z"/>`
	tagSolidInnerSVG                         = `<path fill="currentColor" d="M10 4.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.146.146A.5.5 0 0 1 7.5 0h6A1.5 1.5 0 0 1 15 1.5v6a.5.5 0 0 1-.146.354l-7 7a.5.5 0 0 1-.708 0l-7-7a.5.5 0 0 1 0-.708l7-7ZM10.5 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z" clip-rule="evenodd"/>`
	tailwindOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M7.5 3C5.633 3 4.467 4 4 6c.7-1 1.517-1.375 2.45-1.125c.533.143.913.557 1.334 1.015C8.471 6.636 9.265 7.5 11 7.5c1.867 0 3.033-1 3.5-3c-.7 1-1.517 1.375-2.45 1.125c-.533-.143-.913-.557-1.334-1.015C10.029 3.864 9.235 3 7.5 3ZM4 7.5c-1.867 0-3.033 1-3.5 3c.7-1 1.517-1.375 2.45-1.125c.533.143.913.557 1.334 1.015C4.971 11.136 5.765 12 7.5 12c1.867 0 3.033-1 3.5-3c-.7 1-1.517 1.375-2.45 1.125c-.533-.143-.913-.557-1.334-1.015C6.529 8.364 5.735 7.5 4 7.5Z" clip-rule="evenodd"/>`
	tailwindSolidInnerSVG                    = `<path fill="currentColor" d="M7.5 2.5c-1.026 0-1.908.277-2.6.87c-.688.59-1.137 1.447-1.387 2.516a.5.5 0 0 0 .897.4c.316-.452.632-.723.936-.863c.294-.135.611-.162.975-.065c.366.098.65.386 1.095.87l.015.016c.336.365.745.81 1.305 1.156c.582.359 1.305.6 2.264.6c1.026 0 1.908-.277 2.6-.87c.688-.59 1.138-1.447 1.387-2.516a.5.5 0 0 0-.897-.4c-.316.452-.632.723-.936.863c-.294.135-.611.162-.975.065c-.366-.098-.65-.386-1.095-.87l-.015-.016c-.336-.365-.745-.81-1.305-1.156c-.582-.359-1.305-.6-2.264-.6ZM4 7c-1.026 0-1.908.277-2.6.87C.712 8.46.263 9.317.013 10.386a.5.5 0 0 0 .897.4c.316-.452.632-.723.936-.863c.294-.135.611-.162.975-.065c.366.098.65.386 1.095.87l.015.016c.336.365.745.81 1.305 1.156c.582.359 1.305.6 2.264.6c1.026 0 1.908-.277 2.6-.87c.688-.59 1.138-1.447 1.387-2.516a.5.5 0 0 0-.897-.4c-.316.452-.632.723-.936.863c-.294.135-.611.162-.975.065c-.366-.098-.65-.386-1.095-.87l-.015-.016c-.335-.365-.745-.81-1.305-1.156C5.682 7.24 4.959 7 4 7Z"/>`
	targetOutlineInnerSVG                    = `<g fill="none" stroke="currentColor"><path d="M.5 7.5a7 7 0 1 0 14 0a7 7 0 0 0-14 0Z"/><path d="M3.5 7.5a4 4 0 1 0 8 0a4 4 0 0 0-8 0Z"/><path d="M6.5 7.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/></g>`
	targetSolidInnerSVG                      = `<path fill="currentColor" d="M7 7.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM3 7.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Z" clip-rule="evenodd"/>`
	telegramOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m14.5 1.5l-14 5l4 2l6-4l-4 5l6 4l2-12Z"/>`
	telegramSolidInnerSVG                    = `<path fill="currentColor" d="M14.993 1.582a.5.5 0 0 0-.661-.553l-14 5a.5.5 0 0 0-.056.918l4 2a.5.5 0 0 0 .501-.031l3.32-2.214L6.11 9.188a.5.5 0 0 0 .113.728l6 4a.5.5 0 0 0 .77-.334l2-12Z"/>`
	terminalOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="m3.5 4.5l3 3l-3 3m4.5 0h4m-10.5-9h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`
	terminalSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h12A1.5 1.5 0 0 1 15 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10Zm5.793 5L3.146 4.854l.708-.708L7.207 7.5l-3.353 3.354l-.708-.707L5.793 7.5ZM12 11H8v-1h4v1Z" clip-rule="evenodd"/>`
	textDocumentAltOutlineInnerSVG           = `<path fill="currentColor" d="M4.5 6.995H4v1h.5v-1Zm6 1h.5v-1h-.5v1ZM4.5 10H4v1h.5v-1Zm6 1h.5v-1h-.5v1Zm-6-7.003H4v1h.5v-1Zm6 1h.5v-1h-.5v1Zm3-1.497h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-6 7.495h6v-1h-6v1ZM4.5 11h6v-1h-6v1Zm0-6.003h6v-1h-6v1Zm8 9.003h-10v1h10v-1ZM2 13.5v-12H1v12h1Zm11-10v10h1v-10h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1Z"/>`
	textDocumentAltSolidInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm3 2.497h7v1H4v-1Zm7 2.998H4v1h7v-1ZM11 10H4v1h7v-1Z" clip-rule="evenodd"/>`
	textDocumentOutlineInnerSVG              = `<path fill="currentColor" d="M4.5 6.995H4v1h.5v-1Zm6 1h.5v-1h-.5v1Zm-6 1.998H4v1h.5v-1Zm6 1.007h.5v-1h-.5v1Zm-6-7.003H4v1h.5v-1ZM8.5 5H9V4h-.5v1Zm2-4.5l.354-.354L10.707 0H10.5v.5Zm3 3h.5v-.207l-.146-.147l-.354.354Zm-9 4.495h6v-1h-6v1Zm0 2.998l6 .007v-1l-6-.007v1Zm0-5.996L8.5 5V4l-4-.003v1Zm8 9.003h-10v1h10v-1ZM2 13.5v-12H1v12h1ZM2.5 1h8V0h-8v1ZM13 3.5v10h1v-10h-1ZM10.146.854l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1Z"/>`
	textDocumentSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm3 2.497L9 4v1l-5-.003v-1Zm7 2.998H4v1h7v-1Zm0 3.006l-7-.008v1L11 11v-1Z" clip-rule="evenodd"/>`
	textOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" d="M7.5 14V1.5M1.5 5V1.5h12V5m-10 8.5H11"/>`
	textSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v4h-1V2H8v11h3v1H3.5v-1H7V2H2v3H1V1Z" clip-rule="evenodd"/>`
	threeHundredSixtyOutlineInnerSVG         = `<path fill="currentColor" d="m4.5 5.5l.4.3a.5.5 0 0 0-.4-.8v.5ZM3 7.5l-.4-.3A.5.5 0 0 0 3 8v-.5Zm9.736-4.646l.382.323l-.382-.323ZM2 6h2.5V5H2v1Zm2.1-.8l-1.5 2l.8.6l1.5-2l-.8-.6ZM3 8h.5V7H3v1Zm.5 1H2v1h1.5V9Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 8.5H4ZM3.5 8a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 7v1ZM8 5h-.5v1H8V5ZM6 6.5v1h1v-1H6Zm0 1v1h1v-1H6ZM7.5 7h-1v1h1V7ZM9 8.5A1.5 1.5 0 0 0 7.5 7v1a.5.5 0 0 1 .5.5h1ZM7.5 10A1.5 1.5 0 0 0 9 8.5H8a.5.5 0 0 1-.5.5v1Zm0-1a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 10V9Zm0-4A1.5 1.5 0 0 0 6 6.5h1a.5.5 0 0 1 .5-.5V5ZM12 6.5v2h1v-2h-1Zm-1 2v-2h-1v2h1Zm.5.5a.5.5 0 0 1-.5-.5h-1a1.5 1.5 0 0 0 1.5 1.5V9Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 13 8.5h-1ZM11.5 6a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 5v1Zm0-1A1.5 1.5 0 0 0 10 6.5h1a.5.5 0 0 1 .5-.5V5Zm-4 9A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Zm6 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 2v1Zm0 1a.5.5 0 0 1-.5-.5h-1A1.5 1.5 0 0 0 13.5 5V4Zm-.5-.5c0-.123.044-.235.118-.323l-.763-.646c-.221.261-.355.6-.355.969h1Zm.118-.323A.498.498 0 0 1 13.5 3V2c-.46 0-.871.207-1.145.531l.763.646ZM7.5 1c1.934 0 3.671.844 4.862 2.186l.748-.664A7.483 7.483 0 0 0 7.5 0v1Zm5.854 3.67c.414.856.646 1.815.646 2.83h1c0-1.17-.268-2.277-.746-3.265l-.9.436ZM14 3.5a.5.5 0 0 1-.348.477l.304.952A1.5 1.5 0 0 0 15 3.5h-1Zm-.348.477A.5.5 0 0 1 13.5 4v1a1.5 1.5 0 0 0 .456-.07l-.304-.953Z"/>`
	threeHundredSixtySolidInnerSVG           = `<path fill="currentColor" d="M7 8.5V8h.5a.5.5 0 1 1-.5.5ZM11.5 6a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.5-.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 0 1 12.787-5.32a1.5 1.5 0 0 1 1.659 2.484A7.52 7.52 0 0 1 15 7.5a7.5 7.5 0 0 1-15 0ZM13.5 3a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM2 6h1.5l-.9 1.2A.5.5 0 0 0 3 8h.5a.5.5 0 0 1 0 1H2v1h1.5a1.5 1.5 0 0 0 .449-2.932L4.9 5.8a.5.5 0 0 0-.4-.8H2v1Zm5.5-1A1.5 1.5 0 0 0 6 6.5v2A1.5 1.5 0 1 0 7.5 7H7v-.5a.5.5 0 0 1 .5-.5H8V5h-.5ZM10 6.5a1.5 1.5 0 0 1 3 0v2a1.5 1.5 0 0 1-3 0v-2Z" clip-rule="evenodd"/>`
	thumbDownOutlineInnerSVG                 = `<path fill="currentColor" d="M3.5 9.5H3v.146l.078.122L3.5 9.5Zm2.698 4.24l.421-.27l-.421.27Zm2.667-1.51l-.448.223l.448-.224ZM7.5 9.5V9a.5.5 0 0 0-.447.724L7.5 9.5Zm7-5h.5v-.167l-.1-.133l-.4.3Zm-2.4-3.2l.4-.3l-.4.3ZM8.282 14.231l.257.429l-.257-.429ZM1 10V0H0v10h1Zm2.078-.232l2.698 4.24l.843-.537l-2.697-4.24l-.844.537Zm6.234 2.237L7.947 9.276l-.894.448l1.364 2.729l.895-.447ZM7.5 10h5V9h-5v1ZM15 7.5v-3h-1v3h1Zm-.1-3.3L12.5 1l-.8.6l2.4 3.2l.8-.6ZM10.5 0h-5v1h5V0ZM3 2.5v7h1v-7H3Zm9.5 7.5A2.5 2.5 0 0 0 15 7.5h-1A1.5 1.5 0 0 1 12.5 9v1Zm-7-10A2.5 2.5 0 0 0 3 2.5h1A1.5 1.5 0 0 1 5.5 1V0Zm3.039 14.66a2.034 2.034 0 0 0 .773-2.655l-.895.448c.242.483.07 1.071-.393 1.35l.515.857ZM12.5 1a2.5 2.5 0 0 0-2-1v1a1.5 1.5 0 0 1 1.2.6l.8-.6ZM5.776 14.008a2.034 2.034 0 0 0 2.763.652l-.515-.858a1.034 1.034 0 0 1-1.405-.331l-.843.537Z"/>`
	thumbDownSolidInnerSVG                   = `<path fill="currentColor" d="M1 10V0H0v10h1ZM5.5 0A2.5 2.5 0 0 0 3 2.5v7.146l2.776 4.361a2.034 2.034 0 0 0 3.536-2.002L8.309 10H12.5A2.5 2.5 0 0 0 15 7.5V4.333L12.5 1a2.5 2.5 0 0 0-2-1h-5Z"/>`
	thumbUpOutlineInnerSVG                   = `<path fill="currentColor" d="m3.5 5.5l-.422-.268L3 5.354V5.5h.5Zm2.698-4.24l.421.27l-.421-.27Zm2.667 1.51l-.448-.223l.448.224ZM7.5 5.5l-.447-.224A.5.5 0 0 0 7.5 6v-.5Zm7 5l.4.3l.1-.133V10.5h-.5Zm-2.4 3.2l.4.3l-.4-.3ZM8.282.769L8.539.34l-.257.43ZM0 5v10h1V5H0Zm3.922.768L6.619 1.53L5.776.992l-2.698 4.24l.844.536Zm4.495-3.22L7.053 5.275l.894.448l1.365-2.73l-.895-.447ZM7.5 6h5V5h-5v1ZM14 7.5v3h1v-3h-1Zm.1 2.7l-2.4 3.2l.8.6l2.4-3.2l-.8-.6ZM10.5 14h-5v1h5v-1ZM4 12.5v-7H3v7h1ZM12.5 6A1.5 1.5 0 0 1 14 7.5h1A2.5 2.5 0 0 0 12.5 5v1Zm-7 8A1.5 1.5 0 0 1 4 12.5H3A2.5 2.5 0 0 0 5.5 15v-1ZM8.024 1.198c.464.278.635.866.393 1.35l.895.446A2.034 2.034 0 0 0 8.539.34l-.515.858ZM11.7 13.4a1.5 1.5 0 0 1-1.2.6v1a2.5 2.5 0 0 0 2-1l-.8-.6ZM6.62 1.53c.3-.474.924-.62 1.404-.332L8.54.34a2.034 2.034 0 0 0-2.763.652l.843.537Z"/>`
	thumbUpSolidInnerSVG                     = `<path fill="currentColor" d="M9.312 2.995A2.034 2.034 0 0 0 5.776.992L3 5.354V12.5A2.5 2.5 0 0 0 5.5 15h5a2.5 2.5 0 0 0 2-1l2.5-3.333V7.5A2.5 2.5 0 0 0 12.5 5H8.309l1.003-2.005ZM0 5v10h1V5H0Z"/>`
	thumbtackOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M.5 14.5L5 10M.5 5.5l9 9m-1-14l6 6m-13 0l8-5m-1 12l5-8"/>`
	thumbtackSolidInnerSVG                   = `<path fill="currentColor" d="M8.702 1.41L8.146.853l.708-.708l6 6l-.707.708l-.556-.556l-4.456 7.13l.719.718l-.708.708L5 10.707L.854 14.854l-.708-.707L4.293 10L.146 5.854l.708-.708l.718.72l7.13-4.457Z"/>`
	tickCircleOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M4 7.5L7 10l4-5m-3.5 9.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	tickCircleSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm7.072 3.21l4.318-5.398l-.78-.624l-3.682 4.601L4.32 7.116l-.64.768l3.392 2.827Z" clip-rule="evenodd"/>`
	tickOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m1 7l4.5 4.5L14 3"/>`
	tickSmallOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M4 7.5L7 10l4-5"/>`
	tickSmallSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="m11.39 5.312l-4.318 5.399L3.68 7.884l.64-.768l2.608 2.173l3.682-4.601l.78.624Z" clip-rule="evenodd"/>`
	tickSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M14.707 3L5.5 12.207L.293 7L1 6.293l4.5 4.5l8.5-8.5l.707.707Z" clip-rule="evenodd"/>`
	tiktokOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M9.5 0v11A3.5 3.5 0 1 1 6 7.5m8-2A4.5 4.5 0 0 1 9.5 1"/>`
	tiktokSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M9 0h1v1a4 4 0 0 0 4 4v1a4.992 4.992 0 0 1-4-2v7a4 4 0 1 1-4-4v1a3 3 0 1 0 3 3V0Z" clip-rule="evenodd"/>`
	toggleOutlineInnerSVG                    = `<g fill="none" stroke="currentColor"><path d="M3.5 2.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/><path d="M11.5.5h-8a3 3 0 0 0 0 6h8a3 3 0 1 0 0-6Zm0 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/><path d="M3.5 14.5h8a3 3 0 1 0 0-6h-8a3 3 0 0 0 0 6Z"/></g>`
	toggleSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M0 3.5A3.5 3.5 0 0 1 3.5 0h8a3.5 3.5 0 1 1 0 7h-8A3.5 3.5 0 0 1 0 3.5ZM3.5 2a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM15 11.5a3.5 3.5 0 0 1-3.5 3.5h-8a3.5 3.5 0 1 1 0-7h8a3.5 3.5 0 0 1 3.5 3.5ZM11.5 13a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z" clip-rule="evenodd"/>`
	topLeftOutlineInnerSVG                   = `<path fill="currentColor" d="M1.5 1.5V1a.5.5 0 0 0-.5.5h.5Zm0 .5H7V1H1.5v1ZM1 1.5V7h1V1.5H1Zm.146.354l12 12l.707-.708l-12-12l-.707.708Z"/>`
	topLeftSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h6v1H2.707l11.147 11.146l-.708.708L2 2.707V7H1V1Z" clip-rule="evenodd"/>`
	topRightOutlineInnerSVG                  = `<path fill="currentColor" d="M13.5 1.5h.5a.5.5 0 0 0-.5-.5v.5Zm0-.5H8v1h5.5V1Zm-.5.5V7h1V1.5h-1Zm.146-.354l-12 12l.708.708l12-12l-.708-.708Z"/>`
	topRightSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M8 1h6v6h-1V2.707L1.854 13.854l-.708-.708L12.293 2H8V1Z" clip-rule="evenodd"/>`
	trendDownOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M.5 3.5L5 8l3-3l6.5 6.5m0 0V5m0 6.5H8"/>`
	trendDownSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="m.146 3.854l.708-.708L5 7.293l3-3l6 6V5h1v7H8v-1h5.293L8 5.707l-3 3L.146 3.854Z" clip-rule="evenodd"/>`
	trendUpOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M.5 11.5L5 7l3 3l6.5-6.5m0 0V10m0-6.5H8"/>`
	trendUpSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M8 3h7v7h-1V4.707l-6 6l-3-3l-4.146 4.147l-.708-.708L5 6.293l3 3L13.293 4H8V3Z" clip-rule="evenodd"/>`
	triangleOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m7.5 1.5l-7 12h14l-7-12Z"/>`
	triangleSolidInnerSVG                    = `<path fill="currentColor" d="M7.932 1.248a.5.5 0 0 0-.864 0l-7 12A.5.5 0 0 0 .5 14h14a.5.5 0 0 0 .432-.752l-7-12Z"/>`
	trophyOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M4 14.5h7m-3.5 0v-5m0 0a4 4 0 0 0 4-4v-4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v4a4 4 0 0 0 4 4Zm-4-7h-1a2 2 0 1 0 0 4h1m8-4h1a2 2 0 1 1 0 4h-1"/>`
	trophySolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 0A1.5 1.5 0 0 0 3 1.5V2h-.5a2.5 2.5 0 0 0 0 5h.756A4.504 4.504 0 0 0 7 9.973V14H4v1h7v-1H8V9.973A4.504 4.504 0 0 0 11.744 7h.756a2.5 2.5 0 0 0 0-5H12v-.5A1.5 1.5 0 0 0 10.5 0h-6ZM12 3v2.5c0 .169-.01.336-.027.5h.527a1.5 1.5 0 0 0 0-3H12ZM2.5 3H3v2.5c0 .169.01.336.027.5H2.5a1.5 1.5 0 1 1 0-3Z" clip-rule="evenodd"/>`
	tvOutlineInnerSVG                        = `<path fill="none" stroke="currentColor" d="M2.5 11.5v2m10-2v2M1 13.5h3m7 0h3M.5 2.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`
	tvSolidInnerSVG                          = `<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v8A1.5 1.5 0 0 0 1.5 12H2v1H1v1h3v-1H3v-1h9v1h-1v1h3v-1h-1v-1h.5a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 1h-12Z"/>`
	twitchOutlineInnerSVG                    = `<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.5.5h.5Zm14 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 8l.354.354A.5.5 0 0 0 15 8.5h-.5Zm-3 3v.5a.5.5 0 0 0 .354-.146L11.5 11.5Zm-5 0V11a.5.5 0 0 0-.325.12l.325.38Zm-3.5 3h-.5a.5.5 0 0 0 .825.38L3 14.5Zm0-3h.5A.5.5 0 0 0 3 11v.5Zm-2.5 0H0a.5.5 0 0 0 .5.5v-.5ZM.5 1h14V0H.5v1ZM14 .5v8h1v-8h-1Zm.146 7.646l-3 3l.708.708l3-3l-.708-.708ZM11.5 11h-5v1h5v-1Zm-5.325.12l-3.5 3l.65.76l3.5-3l-.65-.76ZM3.5 14.5v-3h-1v3h1ZM3 11H.5v1H3v-1Zm-2 .5V.5H0v11h1ZM10 3v5h1V3h-1ZM7 3v5h1V3H7Z"/>`
	twitchSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M.5 0a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h2v2.5a.5.5 0 0 0 .825.38L6.685 12H11.5a.5.5 0 0 0 .354-.146l3-3A.5.5 0 0 0 15 8.5v-8a.5.5 0 0 0-.5-.5H.5ZM10 8V3h1v5h-1ZM7 3v5h1V3H7Z" clip-rule="evenodd"/>`
	twitterOutlineInnerSVG                   = `<path fill="currentColor" d="m14.478 1.5l.5-.033a.5.5 0 0 0-.871-.301l.371.334Zm-.498 2.959a.5.5 0 1 0-1 0h1Zm-6.49.082h-.5h.5Zm0 .959h.5h-.5Zm-6.99 7V12a.5.5 0 0 0-.278.916L.5 12.5Zm.998-11l.469-.175a.5.5 0 0 0-.916-.048l.447.223Zm3.994 9l.354.353a.5.5 0 0 0-.195-.827l-.159.474Zm7.224-8.027l-.37.336l.18.199l.265-.04l-.075-.495Zm1.264-.94c.051.778.003 1.25-.123 1.606c-.122.345-.336.629-.723 1l.692.722c.438-.42.776-.832.974-1.388c.193-.546.232-1.178.177-2.006l-.998.066Zm0 3.654V4.46h-1v.728h1Zm-6.99-.646V5.5h1v-.959h-1Zm0 .959V6h1v-.5h-1ZM10.525 1a3.539 3.539 0 0 0-3.537 3.541h1A2.539 2.539 0 0 1 10.526 2V1Zm2.454 4.187C12.98 9.503 9.487 13 5.18 13v1c4.86 0 8.8-3.946 8.8-8.813h-1ZM1.03 1.675C1.574 3.127 3.614 6 7.49 6V5C4.174 5 2.421 2.54 1.966 1.325l-.937.35Zm.021-.398C.004 3.373-.157 5.407.604 7.139c.759 1.727 2.392 3.055 4.73 3.835l.317-.948c-2.155-.72-3.518-1.892-4.132-3.29c-.612-1.393-.523-3.11.427-5.013l-.895-.446Zm4.087 8.87C4.536 10.75 2.726 12 .5 12v1c2.566 0 4.617-1.416 5.346-2.147l-.708-.706Zm7.949-8.009A3.445 3.445 0 0 0 10.526 1v1c.721 0 1.37.311 1.82.809l.74-.671Zm-.296.83a3.513 3.513 0 0 0 2.06-1.134l-.744-.668a2.514 2.514 0 0 1-1.466.813l.15.989ZM.222 12.916C1.863 14.01 3.583 14 5.18 14v-1c-1.63 0-3.048-.011-4.402-.916l-.556.832Z"/>`
	twitterSolidInnerSVG                     = `<path fill="currentColor" d="M14.977 1.467a.5.5 0 0 0-.87-.301a2.559 2.559 0 0 1-1.226.763A3.441 3.441 0 0 0 10.526 1a3.539 3.539 0 0 0-3.537 3.541v.44C3.998 4.75 2.4 2.477 1.967 1.325a.5.5 0 0 0-.916-.048C.004 3.373-.157 5.407.604 7.139C1.27 8.656 2.61 9.864 4.51 10.665C3.647 11.276 2.194 12 .5 12a.5.5 0 0 0-.278.916C1.847 14 3.55 14 5.132 14h.048c4.861 0 8.8-3.946 8.8-8.812v-.479c.363-.37.646-.747.82-1.236c.193-.546.232-1.178.177-2.006Z"/>`
	typescriptOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M12.5 8v-.167c0-.736-.597-1.333-1.333-1.333H10a1.5 1.5 0 1 0 0 3h1a1.5 1.5 0 0 1 0 3h-1A1.5 1.5 0 0 1 8.5 11M8 6.5H3m2.5 0V13M.5.5h14v14H.5V.5Z"/>`
	typescriptSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M0 0h15v15H0V0Zm10 6a2 2 0 1 0 0 4h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1H8a2 2 0 0 0 2 2h1a2 2 0 1 0 0-4h-1a1 1 0 0 1 0-2h1.167c.46 0 .833.373.833.833V8h1v-.167C13 6.821 12.18 6 11.167 6H10ZM3 6h5v1H6v6H5V7H3V6Z" clip-rule="evenodd"/>`
	underlineOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M2 13.5h11M3.5 1v6.5a4 4 0 1 0 8 0V1"/>`
	underlineSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M3 7.5V1h1v6.5a3.5 3.5 0 1 0 7 0V1h1v6.5a4.5 4.5 0 0 1-9 0ZM13 13v1H2v-1h11Z" clip-rule="evenodd"/>`
	unlockCircleOutlineInnerSVG              = `<path fill="currentColor" d="m9.121 4.121l.354-.353l-.354.353ZM5.5 7h4V6h-4v1Zm4.5.5v3h1v-3h-1ZM9.5 11h-4v1h4v-1ZM5 10.5v-3H4v3h1Zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12v-1Zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6v1Zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5V6Zm.5.5v-.879H5V6.5h1Zm2.768-2.025l.378.379l.708-.708l-.38-.378l-.706.707ZM7.62 4c.43 0 .843.17 1.147.475l.707-.707A2.621 2.621 0 0 0 7.62 3v1ZM6 5.621C6 4.726 6.726 4 7.621 4V3A2.621 2.621 0 0 0 5 5.621h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	unlockCircleSolidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM7.621 4C6.726 4 6 4.726 6 5.621V6h3.5A1.5 1.5 0 0 1 11 7.5v3A1.5 1.5 0 0 1 9.5 12h-4A1.5 1.5 0 0 1 4 10.5v-3a1.5 1.5 0 0 1 1-1.415v-.464a2.621 2.621 0 0 1 4.475-1.853l.379.378l-.708.708l-.378-.38A1.621 1.621 0 0 0 7.62 4Z" clip-rule="evenodd"/>`
	unlockOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M4.5 6.5v-3a3 3 0 0 1 6 0V4m-8 2.5h10a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`
	unlockSmallOutlineInnerSVG               = `<path fill="currentColor" d="m9.121 4.121l.354-.353l-.354.353ZM5.5 7h4V6h-4v1Zm4.5.5v3h1v-3h-1ZM9.5 11h-4v1h4v-1ZM5 10.5v-3H4v3h1Zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12v-1Zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6v1Zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5V6Zm.5.5v-.879H5V6.5h1Zm2.768-2.025l.378.379l.708-.708l-.38-.378l-.706.707ZM7.62 4c.43 0 .843.17 1.147.475l.707-.707A2.621 2.621 0 0 0 7.62 3v1ZM6 5.621C6 4.726 6.726 4 7.621 4V3A2.621 2.621 0 0 0 5 5.621h1Z"/>`
	unlockSmallSolidInnerSVG                 = `<path fill="currentColor" d="M6 5.621a1.621 1.621 0 0 1 2.768-1.146l.378.379l.708-.708l-.38-.378A2.621 2.621 0 0 0 5 5.62v.464A1.5 1.5 0 0 0 4 7.5v3A1.5 1.5 0 0 0 5.5 12h4a1.5 1.5 0 0 0 1.5-1.5v-3A1.5 1.5 0 0 0 9.5 6H6v-.379Z"/>`
	unlockSolidInnerSVG                      = `<path fill="currentColor" d="M5 3.5a2.5 2.5 0 0 1 5 0V4h1v-.5a3.5 3.5 0 1 0-7 0V6H2.5A1.5 1.5 0 0 0 1 7.5v6A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-6A1.5 1.5 0 0 0 12.5 6H5V3.5Z"/>`
	upCircleOutlineInnerSVG                  = `<path fill="currentColor" d="m10.146 8.854l.354.353l.707-.707l-.353-.354l-.708.708ZM7.5 5.5l.354-.354l-.354-.353l-.354.353l.354.354ZM4.146 8.146l-.353.354l.707.707l.354-.353l-.708-.708Zm6.708 0l-3-3l-.708.708l3 3l.708-.708Zm-3.708-3l-3 3l.708.708l3-3l-.708-.708ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1Zm1 0A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1Z"/>`
	upCircleSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15ZM3.293 9L7.5 4.793L11.707 9H3.293Z" clip-rule="evenodd"/>`
	upOutlineInnerSVG                        = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m1 10l6.5-7l6.5 7"/>`
	upSmallOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="m10.5 8.5l-3-3l-3 3"/>`
	upSmallSolidInnerSVG                     = `<path fill="currentColor" d="M7.5 4.793L3.293 9h8.414L7.5 4.793Z"/>`
	upSolidInnerSVG                          = `<path fill="currentColor" d="m7.5 3l7.5 8H0l7.5-8Z"/>`
	uploadOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m7.5 1.5l3.25 3m-3.25-3l-3 3m3-3V11m6-4v6.5h-12V7"/>`
	uploadSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="m7.486.807l3.603 3.326l-.678.734L8 2.642V11H7V2.707L4.854 4.854l-.708-.708l3.34-3.34ZM2 13V7H1v7h13V7h-1v6H2Z" clip-rule="evenodd"/>`
	usbCableOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M4.5 5.5h6m-6 0a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1m-6 0v-5h6v5M6.5 2v2m2-2v2m-3 7.5v2h4v-2m-3 2V15m2-1.5V15"/>`
	usbCableSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M11 0H4v6h7V0ZM6 4V2h1v2H6Zm2 0V2h1v2H8Z" clip-rule="evenodd"/><path fill="currentColor" d="M12 7H3v3.5A1.5 1.5 0 0 0 4.5 12H5v2h1v1h1v-1h1v1h1v-1h1v-2h.5a1.5 1.5 0 0 0 1.5-1.5V7Z"/>`
	userCircleOutlineInnerSVG                = `<path fill="currentColor" d="M3 13v.5h1V13H3Zm8 0v.5h1V13h-1Zm-7 0v-.5H3v.5h1Zm2.5-3h2V9h-2v1Zm4.5 2.5v.5h1v-.5h-1ZM8.5 10a2.5 2.5 0 0 1 2.5 2.5h1A3.5 3.5 0 0 0 8.5 9v1ZM4 12.5A2.5 2.5 0 0 1 6.5 10V9A3.5 3.5 0 0 0 3 12.5h1ZM7.5 3A2.5 2.5 0 0 0 5 5.5h1A1.5 1.5 0 0 1 7.5 4V3ZM10 5.5A2.5 2.5 0 0 0 7.5 3v1A1.5 1.5 0 0 1 9 5.5h1ZM7.5 8A2.5 2.5 0 0 0 10 5.5H9A1.5 1.5 0 0 1 7.5 7v1Zm0-1A1.5 1.5 0 0 1 6 5.5H5A2.5 2.5 0 0 0 7.5 8V7Zm0 7A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`
	userCircleSolidInnerSVG                  = `<path fill="currentColor" d="M5 5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM1 7.5a6.5 6.5 0 1 1 10.988 4.702A3.5 3.5 0 0 0 8.5 9h-2a3.5 3.5 0 0 0-3.488 3.202A6.482 6.482 0 0 1 1 7.5Z" clip-rule="evenodd"/>`
	userMinusOutlineInnerSVG                 = `<path fill="currentColor" d="M10.5 14.49v.5h.5v-.5h-.5Zm-10 0H0v.5h.5v-.5Zm7-4.996v.5v-.5Zm-4 0v-.5v.5ZM8 3.498a2.499 2.499 0 0 1-2.5 2.498v1C7.433 6.996 9 5.43 9 3.498H8ZM5.5 5.996A2.499 2.499 0 0 1 3 3.498H2a3.499 3.499 0 0 0 3.5 3.498v-1ZM3 3.498A2.499 2.499 0 0 1 5.5 1V0A3.499 3.499 0 0 0 2 3.498h1ZM5.5 1A2.5 2.5 0 0 1 8 3.498h1A3.499 3.499 0 0 0 5.5 0v1Zm5 12.99H.5v1h10v-1Zm-9.5.5v-1.995H0v1.995h1Zm2.5-4.496h4v-1h-4v1Zm6.5 2.5v1.996h1v-1.996h-1Zm-2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5v1Zm-6.5 2.5a2.5 2.5 0 0 1 2.5-2.5v-1a3.5 3.5 0 0 0-3.5 3.5h1ZM10 8h5V7h-5v1Z"/>`
	userMinusSolidInnerSVG                   = `<path fill="currentColor" d="M5.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 5.5 0ZM10 8h5V7h-5v1Zm-2.5.994h-4a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5Z"/>`
	userOutlineInnerSVG                      = `<path fill="none" stroke="currentColor" stroke-linecap="square" d="M10.5 3.498a2.999 2.999 0 0 1-3 2.998a2.999 2.999 0 1 1 3-2.998Zm2 10.992h-10v-1.996a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v1.997Z" clip-rule="evenodd"/>`
	userPlusOutlineInnerSVG                  = `<path fill="currentColor" d="M10.5 14.49v.5h.5v-.5h-.5Zm-10 0H0v.5h.5v-.5Zm7-4.996v.5v-.5Zm-4 0v-.5v.5ZM8 3.498a2.499 2.499 0 0 1-2.5 2.498v1C7.433 6.996 9 5.43 9 3.498H8ZM5.5 5.996A2.499 2.499 0 0 1 3 3.498H2a3.499 3.499 0 0 0 3.5 3.498v-1ZM3 3.498A2.499 2.499 0 0 1 5.5 1V0A3.499 3.499 0 0 0 2 3.498h1ZM5.5 1A2.5 2.5 0 0 1 8 3.498h1A3.499 3.499 0 0 0 5.5 0v1Zm5 12.99H.5v1h10v-1Zm-9.5.5v-1.995H0v1.995h1Zm2.5-4.496h4v-1h-4v1Zm6.5 2.5v1.996h1v-1.996h-1Zm-2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5v1Zm-6.5 2.5a2.5 2.5 0 0 1 2.5-2.5v-1a3.5 3.5 0 0 0-3.5 3.5h1ZM12 5v5h1V5h-1Zm-2 3h5V7h-5v1Z"/>`
	userPlusSolidInnerSVG                    = `<path fill="currentColor" d="M5.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 5.5 0ZM12 5v2h-2v1h2v2h1V8h2V7h-2V5h-1ZM7.5 8.994h-4a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5Z"/>`
	userSolidInnerSVG                        = `<path fill="currentColor" d="M7.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 7.5 0Zm-2 8.994a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5h-4Z"/>`
	userSquareOutlineInnerSVG                = `<path fill="currentColor" d="M3 14.5v.5h1v-.5H3Zm8 0v.5h1v-.5h-1Zm-7 0v-2H3v2h1ZM6.5 10h2V9h-2v1Zm4.5 2.5v2h1v-2h-1ZM8.5 10a2.5 2.5 0 0 1 2.5 2.5h1A3.5 3.5 0 0 0 8.5 9v1ZM4 12.5A2.5 2.5 0 0 1 6.5 10V9A3.5 3.5 0 0 0 3 12.5h1ZM7.5 3A2.5 2.5 0 0 0 5 5.5h1A1.5 1.5 0 0 1 7.5 4V3ZM10 5.5A2.5 2.5 0 0 0 7.5 3v1A1.5 1.5 0 0 1 9 5.5h1ZM7.5 8A2.5 2.5 0 0 0 10 5.5H9A1.5 1.5 0 0 1 7.5 7v1Zm0-1A1.5 1.5 0 0 1 6 5.5H5A2.5 2.5 0 0 0 7.5 8V7Zm-6-6h12V0h-12v1Zm12.5.5v12h1v-12h-1ZM13.5 14h-12v1h12v-1ZM1 13.5v-12H0v12h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15v-1Zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0v1Zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5V0Z"/>`
	userSquareSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v12A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0h-12Zm5 9A3.5 3.5 0 0 0 3 12.5v1a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-1A3.5 3.5 0 0 0 8.5 9h-2ZM5 5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Z" clip-rule="evenodd"/>`
	usersOutlineInnerSVG                     = `<path fill="currentColor" d="M10.5 14.49v.5h.5v-.5h-.5Zm-10 0H0v.5h.5v-.5Zm14 .01v.5h.5v-.5h-.5ZM8 3.498a2.499 2.499 0 0 1-2.5 2.498v1C7.433 6.996 9 5.43 9 3.498H8ZM5.5 5.996A2.499 2.499 0 0 1 3 3.498H2a3.499 3.499 0 0 0 3.5 3.498v-1ZM3 3.498A2.499 2.499 0 0 1 5.5 1V0A3.499 3.499 0 0 0 2 3.498h1ZM5.5 1A2.5 2.5 0 0 1 8 3.498h1A3.499 3.499 0 0 0 5.5 0v1Zm5 12.99H.5v1h10v-1Zm-9.5.5v-1.995H0v1.995h1Zm2.5-4.496h4v-1h-4v1Zm6.5 2.5v1.996h1v-1.996h-1Zm-2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5v1Zm-6.5 2.5a2.5 2.5 0 0 1 2.5-2.5v-1a3.5 3.5 0 0 0-3.5 3.5h1ZM14 13v1.5h1V13h-1Zm.5 1H12v1h2.5v-1ZM12 11a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3v1Zm-.5-3A1.5 1.5 0 0 1 10 6.5H9A2.5 2.5 0 0 0 11.5 9V8ZM13 6.5A1.5 1.5 0 0 1 11.5 8v1A2.5 2.5 0 0 0 14 6.5h-1ZM11.5 5A1.5 1.5 0 0 1 13 6.5h1A2.5 2.5 0 0 0 11.5 4v1Zm0-1A2.5 2.5 0 0 0 9 6.5h1A1.5 1.5 0 0 1 11.5 5V4Z"/>`
	usersSolidInnerSVG                       = `<path fill="currentColor" d="M5.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 5.5 0Zm-2 8.994a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5h-4Zm9 1.006H12v5h3v-2.5a2.5 2.5 0 0 0-2.5-2.5Z"/><path fill="currentColor" d="M11.5 4a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z"/>`
	vectorDocumentOutlineInnerSVG            = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-7 4V4H3v.5h.5Zm2 0H6V4h-.5v.5Zm0 2V7H6v-.5h-.5Zm-2 0H3V7h.5v-.5Zm8-2h.5V4h-.5v.5Zm0 2V7h.5v-.5h-.5Zm-2 0H9V7h.5v-.5Zm0-2V4H9v.5h.5Zm-6 6V10H3v.5h.5Zm2 0H6V10h-.5v.5Zm0 2v.5H6v-.5h-.5Zm-2 0H3v.5h.5v-.5Zm6-2V10H9v.5h.5Zm0 2H9v.5h.5v-.5Zm2 0v.5h.5v-.5h-.5Zm0-2h.5V10h-.5v.5Zm1 3.5h-10v1h10v-1ZM2 13.5v-12H1v12h1Zm11-10v10h1v-10h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM3.5 5h2V4h-2v1ZM5 4.5v2h1v-2H5ZM5.5 6h-2v1h2V6ZM4 6.5v-2H3v2h1Zm7-2v2h1v-2h-1Zm.5 1.5h-2v1h2V6Zm-1.5.5v-2H9v2h1ZM9.5 5h2V4h-2v1Zm-6 6h2v-1h-2v1Zm1.5-.5v2h1v-2H5Zm.5 1.5h-2v1h2v-1Zm-1.5.5v-2H3v2h1Zm5-2v2h1v-2H9Zm.5 2.5h2v-1h-2v1Zm2.5-.5v-2h-1v2h1Zm-.5-2.5h-2v1h2v-1Zm-6-4h4V5h-4v1ZM4 6.5v4h1v-4H4Zm6 0v4h1v-4h-1ZM5.5 12h4v-1h-4v1Z"/>`
	vectorDocumentSolidInnerSVG              = `<path fill="currentColor" d="M9 10v1H6v-1H5V7h1V6h3v1h1v3H9ZM4 5v1h1V5H4Zm6 0v1h1V5h-1Zm-6 7v-1h1v1H4Zm6-1v1h1v-1h-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM3 4h3v1h3V4h3v3h-1v3h1v3H9v-1H6v1H3v-3h1V7H3V4Z" clip-rule="evenodd"/>`
	vennDiagramOutlineInnerSVG               = `<g fill="none" stroke="currentColor"><path d="M2.5 9.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path d="M.5 5.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path d="M4.5 5.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/></g>`
	vennDiagramSolidInnerSVG                 = `<path fill="currentColor" d="M7.5 15a5.502 5.502 0 0 1-5.255-3.873A6.47 6.47 0 0 0 5.5 12c.698 0 1.37-.11 2-.313a6.51 6.51 0 0 0 2 .313a6.47 6.47 0 0 0 3.255-.873A5.503 5.503 0 0 1 7.5 15Zm6.454-6.273A5.5 5.5 0 0 0 9 .023a6.5 6.5 0 0 1 2.96 4.748a6.484 6.484 0 0 1 1.994 3.956ZM6 .022a5.5 5.5 0 0 0-4.954 8.704A6.484 6.484 0 0 1 3.04 4.772A6.5 6.5 0 0 1 6 .022ZM2.005 9.747A5.477 5.477 0 0 0 6 10.977a6.5 6.5 0 0 1-2.953-4.704a5.475 5.475 0 0 0-1.04 3.474ZM13 9.5c0-1.205-.388-2.32-1.046-3.227a6.5 6.5 0 0 1-2.953 4.705a5.477 5.477 0 0 0 3.994-1.23c.003-.083.005-.165.005-.248ZM7.5.375a5.515 5.515 0 0 1 3.255 3.498A6.47 6.47 0 0 0 7.5 3a6.47 6.47 0 0 0-3.255.873A5.515 5.515 0 0 1 7.5.375ZM7.5 4c-1.327 0-2.544.47-3.495 1.253A5.502 5.502 0 0 0 7.5 10.625a5.502 5.502 0 0 0 3.495-5.372A5.477 5.477 0 0 0 7.5 4Z"/>`
	viewColumnOutlineInnerSVG                = `<path fill="none" stroke="currentColor" d="M9.5 0v15m4-15v15m-8-15v15m-4-15v15"/>`
	viewColumnSolidInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M1 15V0h1v15H1Zm4 0V0h1v15H5Zm4 0V0h1v15H9Zm4 0V0h1v15h-1Z" clip-rule="evenodd"/>`
	viewGridOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M0 5.5h15m-15-4h15m-15 8h15m-15 4h15M9.5 0v15m4-15v15m-8-15v15m-4-15v15"/>`
	viewGridSolidInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M1 1V0h1v1h3V0h1v1h3V0h1v1h3V0h1v1h1v1h-1v3h1v1h-1v3h1v1h-1v3h1v1h-1v1h-1v-1h-3v1H9v-1H6v1H5v-1H2v1H1v-1H0v-1h1v-3H0V9h1V6H0V5h1V2H0V1h1Zm1 1v3h3V2H2Zm4 0v3h3V2H6Zm4 0v3h3V2h-3Zm3 4h-3v3h3V6Zm0 4h-3v3h3v-3Zm-4 3v-3H6v3h3Zm-4 0v-3H2v3h3ZM2 9h3V6H2v3Zm4-3v3h3V6H6Z" clip-rule="evenodd"/>`
	vimOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M1.5 3.5h1v10h3l8-10v-2h-5v2h2l-5 6v-6h1v-2h-5v2Z"/>`
	vimSolidInnerSVG                         = `<path fill="currentColor" d="M7 1H1v3h1v10h3.74L14 3.675V1H8v3h1.432L6 8.119V4h1V1Z"/>`
	volumeOneOutlineInnerSVG                 = `<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5v.5Zm0-5.996v.5h.138l.12-.071l-.258-.429Zm5-2.998H9a.5.5 0 0 0-.757-.429L8.5 1.5Zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492h-.5Zm-5-3.498h-2v1h2v-1Zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5v-1Zm-.5-.5V5.498H0v3.998h1Zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5h1Zm.5-.499h2v-1h-2v1Zm2.257-.071l5-2.998l-.514-.858l-5 2.998l.514.858ZM8 1.5v11.992h1V1.5H8Zm.757 11.563l-5-2.998l-.514.858l5 2.998l.514-.858ZM10 6v3h1V6h-1Z"/>`
	volumeOneSolidInnerSVG                   = `<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5a1.5 1.5 0 0 0-1.5 1.5v3.997c0 .83.672 1.5 1.5 1.5h1.862l4.88 2.926A.5.5 0 0 0 9 13.492V1.5ZM10 6v3h1V6h-1Z"/>`
	volumeThreeOutlineInnerSVG               = `<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5v.5Zm0-5.996v.5h.138l.12-.071l-.258-.429Zm5-2.998H9a.5.5 0 0 0-.757-.429L8.5 1.5Zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492h-.5Zm-5-3.498h-2v1h2v-1Zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5v-1Zm-.5-.5V5.498H0v3.998h1Zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5h1Zm.5-.499h2v-1h-2v1Zm2.257-.071l5-2.998l-.514-.858l-5 2.998l.514.858ZM8 1.5v11.992h1V1.5H8Zm.757 11.563l-5-2.998l-.514.858l5 2.998l.514-.858ZM10 6v3h1V6h-1Zm2-2v7h1V4h-1Zm2-2v11h1V2h-1Z"/>`
	volumeThreeSolidInnerSVG                 = `<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5a1.5 1.5 0 0 0-1.5 1.5v3.997c0 .83.672 1.5 1.5 1.5h1.862l4.88 2.926A.5.5 0 0 0 9 13.492V1.5Zm5 .5v11h1V2h-1Zm-2 2v7h1V4h-1Zm-2 2v3h1V6h-1Z"/>`
	volumeTwoOutlineInnerSVG                 = `<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5v.5Zm0-5.996v.5h.138l.12-.071l-.258-.429Zm5-2.998H9a.5.5 0 0 0-.757-.429L8.5 1.5Zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492h-.5Zm-5-3.498h-2v1h2v-1Zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5v-1Zm-.5-.5V5.498H0v3.998h1Zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5h1Zm.5-.499h2v-1h-2v1Zm2.257-.071l5-2.998l-.514-.858l-5 2.998l.514.858ZM8 1.5v11.992h1V1.5H8Zm.757 11.563l-5-2.998l-.514.858l5 2.998l.514-.858ZM10 6v3h1V6h-1Zm2-2v7h1V4h-1Z"/>`
	volumeTwoSolidInnerSVG                   = `<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5a1.5 1.5 0 0 0-1.5 1.5v3.997c0 .83.672 1.5 1.5 1.5h1.862l4.88 2.926A.5.5 0 0 0 9 13.492V1.5ZM12 4v7h1V4h-1Zm-2 2v3h1V6h-1Z"/>`
	vrHeadsetOutlineInnerSVG                 = `<path fill="currentColor" d="m.851 5.4l.137.48l-.137-.48Zm13.298 0l.137-.481l-.137.48ZM4.58 12.352l.44.24l-.44-.24Zm5.84 0l.438-.239l-.439.24ZM2.996 3.757l-.464-.185l.464.185Zm-.961 1.057a.5.5 0 0 0 .928.372l-.928-.372Zm9.967-1.057l.464-.185l-.464.185Zm.033 1.429a.5.5 0 1 0 .928-.372l-.928.372Zm1.964.68V9.21h1V5.865h-1ZM11.21 12h-.542v1h.542v-1Zm-6.878 0H3.79v1h.542v-1ZM1 9.21V5.865H0V9.21h1ZM.988 5.88a23.703 23.703 0 0 1 13.024 0l.274-.961a24.703 24.703 0 0 0-13.572 0l.274.961ZM3.79 12A2.79 2.79 0 0 1 1 9.21H0A3.79 3.79 0 0 0 3.79 13v-1Zm.352.113a.217.217 0 0 1 .19-.113v1a.78.78 0 0 0 .687-.408l-.877-.479Zm.877.479c1.071-1.963 3.89-1.963 4.962 0l.877-.479c-1.45-2.658-5.267-2.658-6.716 0l.877.479ZM10.668 12c.08 0 .152.043.19.113l-.877.479a.78.78 0 0 0 .687.408v-1ZM14 9.21A2.79 2.79 0 0 1 11.21 12v1A3.79 3.79 0 0 0 15 9.21h-1Zm1-3.345a.984.984 0 0 0-.714-.946l-.274.961A.016.016 0 0 1 14 5.865h1Zm-14 0a.016.016 0 0 1-.012.015l-.274-.96A.984.984 0 0 0 0 5.865h1Zm1.533-2.293l-.497 1.242l.928.372l.497-1.243l-.928-.371Zm9.006.37l.497 1.244l.928-.372l-.497-1.242l-.928.37ZM4.854 3h5.292V2H4.854v1Zm7.613.572A2.5 2.5 0 0 0 10.146 2v1a1.5 1.5 0 0 1 1.393.943l.928-.371Zm-9.006.37A1.5 1.5 0 0 1 4.854 3V2a2.5 2.5 0 0 0-2.321 1.572l.928.37Z"/>`
	vrHeadsetSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="m12.467 3.572l.394.985c.478.106.953.227 1.425.362c.423.12.714.507.714.946V9.21A3.79 3.79 0 0 1 11.21 13h-.542a.783.783 0 0 1-.687-.408c-1.071-1.963-3.89-1.963-4.962 0a.783.783 0 0 1-.687.408H3.79A3.79 3.79 0 0 1 0 9.21V5.865c0-.44.291-.825.714-.946c.472-.135.947-.256 1.425-.362l.394-.985A2.5 2.5 0 0 1 4.854 2h5.292a2.5 2.5 0 0 1 2.321 1.572Zm-9.006.37A1.5 1.5 0 0 1 4.854 3h5.292a1.5 1.5 0 0 1 1.393.943l.153.384a24.703 24.703 0 0 0-8.384 0l.153-.384Z" clip-rule="evenodd"/>`
	vueOutlineInnerSVG                       = `<path fill="currentColor" d="m7.5 13.5l-.432.252a.5.5 0 0 0 .864 0L7.5 13.5Zm4-12l.434.248A.5.5 0 0 0 11.5 1v.5Zm-4 7l-.434.248a.5.5 0 0 0 .868 0L7.5 8.5Zm-4-7V1a.5.5 0 0 0-.434.748L3.5 1.5Zm3 0l.447-.224L6.81 1H6.5v.5Zm1 2l-.447.224a.5.5 0 0 0 .894 0L7.5 3.5Zm1-2V1h-.309l-.138.276l.447.224Zm-8.432.252l7 12l.864-.504l-7-12l-.864.504Zm7.864 12l7-12l-.864-.504l-7 12l.864.504Zm3.134-12.5l-4 7l.868.496l4-7l-.868-.496Zm-3.132 7l-4-7l-.868.496l4 7l.868-.496ZM3.5 2h3V1h-3v1Zm2.553-.276l1 2l.894-.448l-1-2l-.894.448Zm1.894 2l1-2l-.894-.448l-1 2l.894.448ZM8.5 2h3V1h-3v1Z"/>`
	vueSolidInnerSVG                         = `<path fill="currentColor" d="M2.717 1H.5a.5.5 0 0 0-.432.752l7 12a.5.5 0 0 0 .864 0l7-12A.5.5 0 0 0 14.5 1h-2.217L7.5 8.972L2.717 1Z"/><path fill="currentColor" d="M11.117 1H8.19L7.5 2.382L6.809 1H3.883L7.5 7.028L11.117 1Z"/>`
	walletAltOutlineInnerSVG                 = `<path fill="none" stroke="currentColor" d="M13.5 6V2.5a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V9m.93-3.5H9.5a2 2 0 1 0 0 4h4.93a.07.07 0 0 0 .07-.07V5.57a.07.07 0 0 0-.07-.07Z"/>`
	walletAltSolidInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M14 12.5V10h.43a.57.57 0 0 0 .57-.57V5.57a.57.57 0 0 0-.57-.57H14V2.5A1.5 1.5 0 0 0 12.5 1h-11A1.5 1.5 0 0 0 0 2.5v10A1.5 1.5 0 0 0 1.5 14h11a1.5 1.5 0 0 0 1.5-1.5ZM14 6v3H9.5a1.5 1.5 0 1 1 0-3H14Z" clip-rule="evenodd"/>`
	walletOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="M.5 3.5v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1H3m-2.5 0v-1a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v1H3m-2.5 0H3m6 6h3"/>`
	walletSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h8A1.5 1.5 0 0 1 11 2.5V3h2.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10ZM9 10h3V9H9v1Z" clip-rule="evenodd"/>`
	wanOutlineInnerSVG                       = `<path fill="currentColor" d="m3.5 14.5l-.489.105a.5.5 0 0 0 .967.042L3.5 14.5Zm4-13l.478-.147a.5.5 0 0 0-.956 0L7.5 1.5Zm4 13l-.478.147a.5.5 0 0 0 .967-.042L11.5 14.5ZM.011.605l3 14l.978-.21l-3-14l-.978.21Zm3.967 14.042l4-13l-.956-.294l-4 13l.956.294Zm3.044-13l4 13l.956-.294l-4-13l-.956.294Zm4.967 12.958l3-14l-.978-.21l-3 14l.978.21ZM0 6h15V5H0v1Zm0 4h15V9H0v1Z"/>`
	wanSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M.953 5L.01.605l.98-.21L1.976 5H5.9l1.122-3.647a.5.5 0 0 1 .956 0L9.1 5h3.924l.987-4.605l.978.21L14.047 5H15v1h-1.167l-.643 3H15v1h-2.024l-.987 4.605a.5.5 0 0 1-.967.042L9.592 10H5.408l-1.43 4.647a.5.5 0 0 1-.967-.042L2.024 10H0V9h1.81l-.643-3H0V5h.953ZM2.19 6l.643 3h1.836l.923-3H2.19Zm4.449 0l-.924 3h3.57L8.36 6H6.64Zm1.415-1L7.5 3.2L6.946 5h1.108Zm1.354 1l.923 3h1.836l.643-3H9.408Zm2.545 4h-1.314l.774 2.518l.54-2.518ZM4.36 10H3.047l.54 2.518L4.36 10Z" clip-rule="evenodd"/>`
	wandOutlineInnerSVG                      = `<path fill="currentColor" d="m6.853 8.848l.354-.353l-.707-.707l-.353.353l.706.707ZM9 10.493v.5h1v-.5H9Zm1-.999v-.5H9v.5h1ZM9 1.5V2h1v-.5H9Zm1-1V0H9v.5h1ZM4.5 4.997H4v1h.5v-1Zm1 1H6v-1h-.5v1Zm8-1H13v1h.5v-1Zm1 1h.5v-1h-.5v1Zm-2.353-3.852l-.354.354l.707.707l.353-.354l-.706-.707Zm1.706-.292l.354-.353L13.5.792l-.353.354l.706.707Zm-8-.707L5.5.792l-.707.708l.354.353l.706-.707Zm.294 1.706l.353.354l.707-.707l-.354-.354l-.706.707Zm6.706 5.29l-.353-.354l-.707.707l.354.353l.706-.707Zm.294 1.706l.353.353l.707-.707l-.354-.354l-.706.708ZM.853 14.844l6-5.996l-.706-.707l-6 5.996l.706.707ZM10 10.495v-1H9v1h1ZM10 1.5v-1H9v1h1ZM4.5 5.997h1v-1h-1v1Zm9 0h1v-1h-1v1Zm-.647-3.145l1-.999l-.706-.707l-1 1l.706.706Zm-7.706-.999l1 1l.706-.708l-1-1l-.706.708Zm7 6.995l1 1l.706-.708l-1-.999l-.706.707Z"/>`
	wandSolidInnerSVG                        = `<path fill="currentColor" fill-rule="evenodd" d="M10 0v2H9V0h1ZM5.5.792L7.207 2.5l-.707.707L4.793 1.5L5.5.792Zm8.707.708L12.5 3.206l-.707-.707L13.5.792l.707.708ZM4 4.997h2v1H4v-1Zm9 0h2v1h-2v-1ZM7.207 8.495l-6.354 6.35l-.706-.708L6.5 7.787l.707.708Zm5.293-.707l1.707 1.706l-.707.707l-1.707-1.706l.707-.707ZM10 8.994v2H9v-2h1Z" clip-rule="evenodd"/>`
	watchOutlineInnerSVG                     = `<path fill="none" stroke="currentColor" d="M4.5 3.5h6m-6 0a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m0-8v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2m0 0a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1m0 0h-6m6 0v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2m9-5.5v3"/>`
	watchSolidInnerSVG                       = `<path fill="currentColor" fill-rule="evenodd" d="M4 3.085V1.5A1.5 1.5 0 0 1 5.5 0h4A1.5 1.5 0 0 1 11 1.5v1.585A1.5 1.5 0 0 1 12 4.5v6a1.5 1.5 0 0 1-1 1.415V13.5A1.5 1.5 0 0 1 9.5 15h-4A1.5 1.5 0 0 1 4 13.5v-1.585A1.5 1.5 0 0 1 3 10.5v-6a1.5 1.5 0 0 1 1-1.415ZM5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3H5V1.5ZM5 12h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V12Z" clip-rule="evenodd"/><path fill="currentColor" d="M13 6v3h1V6h-1Z"/>`
	webpackOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="M1.5 10.5v-6m0 6l6 3.5m-6-3.5L5 9M1.5 4.5l6-3.5m-6 3.5l6 3m0-6.5l6 3.5M7.5 1v3.5m6 0v6m0-6l-6 3m6 3l-6 3.5m6-3.5L10 9m-2.5 5V7.5m0-3L5 6v3m2.5-4.5L10 6v3M5 9l2.5 1.5L10 9"/>`
	webpackSolidInnerSVG                     = `<path fill="currentColor" d="m8 4.217l2.25 1.35l3.268-1.635L8 .712v3.505ZM7 .713L1.482 3.932L4.75 5.566L7 4.216V.714ZM1 4.809v5.422l3.5-1.556V6.56L1 4.809Zm.523 6.283L7 14.287v-3.504l-2.034-1.22l-3.443 1.53ZM8 14.287l5.477-3.195l-3.443-1.53L8 10.783v3.504Zm6-4.057V4.81l-3.5 1.75v2.116l3.5 1.556Zm-6-.613l1.5-.9V7.059l-1.5.75v1.808ZM7 7.809v1.808l-1.5-.9V7.059l1.5.75ZM5.811 6.096l1.689.845l1.689-.845L7.5 5.083L5.811 6.096Z"/>`
	whatsappOutlineInnerSVG                  = `<path fill="currentColor" d="m1.9 11.7l.447.224l.138-.277L2.3 11.4l-.4.3Zm1.4 1.4l.3-.4l-.247-.185l-.277.138l.224.447ZM.5 14.5l-.447-.224a.5.5 0 0 0 .67.671L.5 14.5Zm4-10l-.277-.416A.5.5 0 0 0 4 4.5h.5Zm6 6v.5a.5.5 0 0 0 .416-.223L10.5 10.5ZM6.254 5.026l.493-.083l-.493.083Zm.14.836l-.493.082l.493-.082Zm-.432.997l.277.416l-.277-.416Zm4.68 3.428l.416.277l-.416-.277Zm-.668-1.541l.083-.493l-.083.493Zm-.836-.14l-.082.493l.082-.493Zm-.997.432l-.416-.277l.416.277ZM0 7.5c0 1.688.558 3.247 1.5 4.5l.8-.6A6.47 6.47 0 0 1 1 7.5H0ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM15 7.5A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1ZM7.5 15A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1ZM3 13.5A7.47 7.47 0 0 0 7.5 15v-1a6.469 6.469 0 0 1-3.9-1.3l-.6.8ZM.723 14.947l2.8-1.4l-.448-.894l-2.8 1.4l.448.894Zm.729-3.47l-1.4 2.8l.894.447l1.4-2.8l-.894-.447ZM4 4.5v1h1v-1H4ZM9.5 11h1v-1h-1v1ZM4 5.5A5.5 5.5 0 0 0 9.5 11v-1A4.5 4.5 0 0 1 5 5.5H4Zm.777-.584l.214-.142l-.555-.832l-.213.142l.554.832Zm.984.192l.14.836l.986-.164l-.14-.837l-.986.165Zm-.076 1.335l-.962.641l.554.832l.962-.641l-.554-.832Zm.216-.499a.5.5 0 0 1-.216.499l.554.832a1.5 1.5 0 0 0 .648-1.495l-.986.164Zm-.91-1.17a.5.5 0 0 1 .77.334l.986-.165a1.5 1.5 0 0 0-2.311-1.001l.555.832Zm5.925 6.003l.142-.213l-.832-.555l-.142.214l.832.554Zm-.86-2.524l-.836-.14l-.164.987l.836.139l.165-.986Zm-2.33.508l-.642.962l.832.554l.641-.962l-.832-.554Zm1.494-.648a1.5 1.5 0 0 0-1.495.648l.832.554a.5.5 0 0 1 .499-.216l.164-.986Zm1.838 2.451a1.5 1.5 0 0 0-1.001-2.311l-.165.986a.5.5 0 0 1 .334.77l.832.555Z"/>`
	whatsappSolidInnerSVG                    = `<path fill="currentColor" d="M5 4.768a.5.5 0 0 1 .761.34l.14.836a.5.5 0 0 1-.216.499l-.501.334A4.501 4.501 0 0 1 5 5.5v-.732ZM9.5 10a4.5 4.5 0 0 1-1.277-.184l.334-.5a.5.5 0 0 1 .499-.217l.836.14a.5.5 0 0 1 .34.761H9.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 3.253 6.182l-2.53 1.265a.5.5 0 0 1-.67-.67l1.265-2.53A7.467 7.467 0 0 1 0 7.5Zm4.23-3.42l.206-.138a1.5 1.5 0 0 1 2.311 1.001l.14.837a1.5 1.5 0 0 1-.648 1.495l-.658.438A4.522 4.522 0 0 0 7.286 9.42l.44-.658a1.5 1.5 0 0 1 1.494-.648l.837.14a1.5 1.5 0 0 1 1.001 2.311l-.138.207a.5.5 0 0 1-.42.229h-1A5.5 5.5 0 0 1 4 5.5v-1a.5.5 0 0 1 .23-.42Z" clip-rule="evenodd"/>`
	wifiFullOutlineInnerSVG                  = `<path fill="currentColor" d="M3.219 9.318c1.155-1.4 2.698-2.161 4.281-2.161v-1c-1.917 0-3.732.924-5.052 2.525l.771.636ZM7.5 7.157c1.583 0 3.126.762 4.281 2.161l.771-.636C11.232 7.08 9.417 6.157 7.5 6.157v1ZM.886 6.318C2.659 4.168 5.042 2.985 7.5 2.985v-1c-2.793 0-5.446 1.346-7.386 3.697l.772.636ZM7.5 2.985c2.458 0 4.84 1.183 6.614 3.333l.772-.636C12.946 3.33 10.293 1.985 7.5 1.985v1ZM7.5 12a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 13v-1Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 11.5H8Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 10v1Zm0-1A1.5 1.5 0 0 0 6 11.5h1a.5.5 0 0 1 .5-.5v-1Z"/>`
	wifiFullSolidInnerSVG                    = `<path fill="currentColor" d="M7.5 2.985c-2.458 0-4.84 1.183-6.614 3.333l-.772-.636C2.054 3.33 4.707 1.985 7.5 1.985s5.446 1.346 7.386 3.697l-.772.636C12.341 4.168 9.958 2.985 7.5 2.985Z"/><path fill="currentColor" d="M7.5 7.157c-1.583 0-3.126.762-4.28 2.161l-.772-.636C3.768 7.08 5.583 6.157 7.5 6.157c1.918 0 3.732.924 5.053 2.525l-.772.636c-1.155-1.4-2.698-2.161-4.28-2.161ZM6 11.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`
	wifiLowOutlineInnerSVG                   = `<path fill="currentColor" d="M3.22 9.318c1.154-1.4 2.697-2.161 4.28-2.161v-1c-1.917 0-3.732.924-5.052 2.525l.771.636ZM7.5 7.157c1.583 0 3.126.762 4.281 2.161l.771-.636C11.232 7.08 9.417 6.157 7.5 6.157v1ZM7.5 12a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 13v-1Zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 11.5H8Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 10v1Zm0-1A1.5 1.5 0 0 0 6 11.5h1a.5.5 0 0 1 .5-.5v-1Z"/>`
	wifiLowSolidInnerSVG                     = `<path fill="currentColor" d="M3.219 9.318c1.155-1.4 2.698-2.161 4.28-2.161c1.584 0 3.127.762 4.282 2.161l.771-.636C11.232 7.08 9.417 6.157 7.5 6.157c-1.918 0-3.732.924-5.052 2.525l.77.636ZM6 11.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`
	wifiNoneOutlineInnerSVG                  = `<path fill="none" stroke="currentColor" d="M6.5 11.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`
	wifiNoneSolidInnerSVG                    = `<path fill="currentColor" d="M7.5 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`
	windowsOutlineInnerSVG                   = `<path fill="currentColor" d="m.5 3.5l-.105-.489A.5.5 0 0 0 0 3.5h.5Zm14-3h.5a.5.5 0 0 0-.605-.489L14.5.5Zm0 14l-.07.495A.5.5 0 0 0 15 14.5h-.5Zm-14-2H0a.5.5 0 0 0 .43.495L.5 12.5Zm.105-8.511l14-3l-.21-.978l-14 3l.21.978ZM14 .5v14h1V.5h-1Zm.57 13.505l-14-2l-.14.99l14 2l.14-.99ZM1 12.5v-9H0v9h1ZM.5 8h14V7H.5v1ZM6 2v11h1V2H6Z"/>`
	windowsSolidInnerSVG                     = `<path fill="currentColor" d="M14.814.111A.5.5 0 0 1 15 .5V7H7V1.596L14.395.01a.5.5 0 0 1 .42.1ZM6 1.81L.395 3.011A.5.5 0 0 0 0 3.5V7h6V1.81ZM0 8v4.5a.5.5 0 0 0 .43.495l5.57.796V8H0Zm7 5.934l7.43 1.061A.5.5 0 0 0 15 14.5V8H7v5.934Z"/>`
	wordpressOutlineInnerSVG                 = `<path fill="currentColor" d="m5.5 13.5l-.479.144a.5.5 0 0 0 .957.003L5.5 13.5Zm4 0l-.479.144a.5.5 0 0 0 .943.042L9.5 13.5Zm3.53-8.827l.465.186l-.464-.186Zm-.54-1.931l-.3.4l.3-.4Zm-.133-.1l.3-.4l-.3.4Zm-2 2l.4-.3l-.4.3Zm.286-2l-.3-.4l.3.4ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM2.021 3.644l3 10l.958-.288l-3-10l-.958.288Zm3.957 10.003l2-6.5l-.956-.294l-2 6.5l.956.294ZM6.02 3.644l3 10l.958-.288l-3-10l-.958.288Zm6.768-1.302l-.132-.1l-.6.8l.132.1l.6-.8Zm-2.832 2.6l.643.858l.8-.6l-.643-.857l-.8.6Zm.386-2.7a1.929 1.929 0 0 0-.386 2.7l.8-.6a.929.929 0 0 1 .186-1.3l-.6-.8Zm2.314 0a1.928 1.928 0 0 0-2.314 0l.6.8a.928.928 0 0 1 1.114 0l.6-.8Zm.838 2.617a2.149 2.149 0 0 0-.706-2.517l-.6.8c.416.312.57.863.377 1.345l.929.372ZM2 4h2V3H2v1Zm3 0h3V3H5v1Zm4.964 9.686l3.531-8.827l-.928-.372l-3.531 8.827l.928.372Z"/>`
	wordpressSolidInnerSVG                   = `<path fill="currentColor" d="M1.5 3a7.5 7.5 0 1 1-.634 1h1.262l2.893 9.644a.5.5 0 0 0 .957.003l1.541-5.01l1.502 5.007a.5.5 0 0 0 .943.042l3.53-8.827a2.149 2.149 0 0 0-.705-2.517l-.132-.1a1.929 1.929 0 0 0-2.7 2.7l.643.858l.8-.6l-.643-.857a.929.929 0 0 1 1.3-1.3l.132.099c.416.312.57.863.377 1.345l-2.999 7.498L7.172 4H8V3H5v1h1.128l.875 2.916l-1.497 4.864L3.172 4H4V3H1.5Z"/>`
	xCircleOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="m4.5 4.5l6 6m-6 0l6-6m-3 10a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`
	xCircleSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm10.146 3.354L7.5 8.207l-2.646 2.647l-.708-.707L6.793 7.5L4.146 4.854l.708-.708L7.5 6.793l2.646-2.647l.708.708L8.207 7.5l2.647 2.646l-.707.708Z" clip-rule="evenodd"/>`
	xOutlineInnerSVG                         = `<path fill="none" stroke="currentColor" d="m1.5 1.5l12 12m-12 0l12-12"/>`
	xSmallOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m4.5 4.5l6 6m-6 0l6-6"/>`
	xSmallSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M6.793 7.5L4.146 4.854l.708-.708L7.5 6.793l2.646-2.647l.708.708L8.207 7.5l2.647 2.646l-.708.707L7.5 8.207l-2.646 2.646l-.708-.707L6.793 7.5Z" clip-rule="evenodd"/>`
	xSolidInnerSVG                           = `<path fill="currentColor" fill-rule="evenodd" d="M6.793 7.5L1.146 1.854l.708-.708L7.5 6.793l5.646-5.647l.708.708L8.207 7.5l5.647 5.646l-.707.707L7.5 8.207l-5.646 5.646l-.708-.707L6.793 7.5Z" clip-rule="evenodd"/>`
	xlsOutlineInnerSVG                       = `<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm-4 10H6v.5h.5v-.5Zm-2-1H5v-.207l-.146-.147L4.5 9.5Zm-2-2H2v.207l.146.147L2.5 7.5Zm8-1V6H10v.5h.5Zm0 2H10V9h.5v-.5Zm2 0h.5V8h-.5v.5Zm0 2v.5h.5v-.5h-.5Zm-10-1l-.354-.354L2 9.293V9.5h.5Zm2-2l.354.354L5 7.707V7.5h-.5ZM2 5V1.5H1V5h1Zm11-1.5V5h1V3.5h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM1 12v1.5h1V12H1Zm1.5 3h10v-1h-10v1ZM14 13.5V12h-1v1.5h1ZM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5H1ZM6 6v4.5h1V6H6Zm.5 5H9v-1H6.5v1ZM4 9.5V11h1V9.5H4Zm.854-.354l-2-2l-.708.708l2 2l.708-.708ZM3 7.5V6H2v1.5h1ZM13 6h-2.5v1H13V6Zm-3 .5v2h1v-2h-1Zm.5 2.5h2V8h-2v1Zm1.5-.5v2h1v-2h-1Zm.5 1.5H10v1h2.5v-1ZM3 11V9.5H2V11h1Zm-.146-1.146l2-2l-.708-.708l-2 2l.708.708ZM5 7.5V6H4v1.5h1Z"/>`
	xlsSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm2 5.793V6H2v1.707l.793.793L2 9.293V11h1V9.707l.5-.5l.5.5V11h1V9.293L4.207 8.5L5 7.707V6H4v1.293l-.5.5l-.5-.5ZM6 6h1v4h2v1H6V6Zm7 0h-3v3h2v1h-2v1h3V8h-2V7h2V6Z" clip-rule="evenodd"/>`
	yenOutlineInnerSVG                       = `<path fill="none" stroke="currentColor" d="M7.5 15V7.5m0 0l-5-7m5 7l5-7M3 7.5h9m-9 4h9"/>`
	yenSolidInnerSVG                         = `<path fill="currentColor" fill-rule="evenodd" d="M6.528 7H3v1h4v3H3v1h4v3h1v-3h4v-1H8V8h4V7H8.472L12.907.79l-.814-.58L7.5 6.64L2.907.21l-.814.58L6.528 7Z" clip-rule="evenodd"/>`
	youtubeOutlineInnerSVG                   = `<path fill="currentColor" d="m1.61 12.738l-.104.489l.105-.489Zm11.78 0l.104.489l-.105-.489Zm0-10.476l.104-.489l-.105.489Zm-11.78 0l.106.489l-.105-.489ZM6.5 5.5l.277-.416A.5.5 0 0 0 6 5.5h.5Zm0 4H6a.5.5 0 0 0 .777.416L6.5 9.5Zm3-2l.277.416a.5.5 0 0 0 0-.832L9.5 7.5ZM0 3.636v7.728h1V3.636H0Zm15 7.728V3.636h-1v7.728h1ZM1.506 13.227c3.951.847 8.037.847 11.988 0l-.21-.978a27.605 27.605 0 0 1-11.568 0l-.21.978ZM13.494 1.773a28.606 28.606 0 0 0-11.988 0l.21.978a27.607 27.607 0 0 1 11.568 0l.21-.978ZM15 3.636c0-.898-.628-1.675-1.506-1.863l-.21.978c.418.09.716.458.716.885h1Zm-1 7.728a.905.905 0 0 1-.716.885l.21.978A1.905 1.905 0 0 0 15 11.364h-1Zm-14 0c0 .898.628 1.675 1.506 1.863l.21-.978A.905.905 0 0 1 1 11.364H0Zm1-7.728c0-.427.298-.796.716-.885l-.21-.978A1.905 1.905 0 0 0 0 3.636h1ZM6 5.5v4h1v-4H6Zm.777 4.416l3-2l-.554-.832l-3 2l.554.832Zm3-2.832l-3-2l-.554.832l3 2l.554-.832Z"/>`
	youtubeSolidInnerSVG                     = `<path fill="currentColor" d="M8.599 7.5L7 8.566V6.434L8.599 7.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M1.506 1.773a28.606 28.606 0 0 1 11.988 0A1.905 1.905 0 0 1 15 3.636v7.728c0 .898-.628 1.675-1.506 1.863a28.605 28.605 0 0 1-11.988 0A1.905 1.905 0 0 1 0 11.364V3.636c0-.898.628-1.675 1.506-1.863Zm5.271 3.311A.5.5 0 0 0 6 5.5v4a.5.5 0 0 0 .777.416l3-2a.5.5 0 0 0 0-.832l-3-2Z" clip-rule="evenodd"/>`
	zipOutlineInnerSVG                       = `<path fill="currentColor" d="M3.5 8.5V8h-.39l-.095.379l.485.121Zm2 0l.485-.121L5.89 8H5.5v.5Zm.69 2.758l.484-.122l-.485.122Zm-3.38 0l.486.12l-.485-.12ZM9.5 10.5l.277-.416A.5.5 0 0 0 9 10.5h.5Zm3 2l-.277.416A.5.5 0 0 0 13 12.5h-.5Zm-3-8H9V5h.5v-.5ZM1.5 1h12V0h-12v1Zm12.5.5v12h1v-12h-1ZM13.5 14h-12v1h12v-1ZM1 13.5v-12H0v12h1Zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15v-1Zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0v1Zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5V0ZM3 3h3V2H3v1Zm0 2h3V4H3v1Zm0 2h3V6H3v1Zm.5 2h2V8h-2v1Zm1.515-.379l.69 2.758l.97-.243l-.69-2.757l-.97.242ZM5.219 12H3.781v1h1.438v-1Zm-1.923-.621l.69-2.758l-.971-.242l-.69 2.757l.97.243ZM3.78 12a.5.5 0 0 1-.485-.621l-.97-.243A1.5 1.5 0 0 0 3.78 13v-1Zm1.923-.621A.5.5 0 0 1 5.22 12v1a1.5 1.5 0 0 0 1.455-1.864l-.97.243ZM10 13v-2.5H9V13h1Zm-.777-2.084l3 2l.554-.832l-3-2l-.554.832ZM13 12.5V10h-1v2.5h1ZM9 6v3h1V6H9Zm3 0v3h1V6h-1ZM9.5 8h3V7h-3v1Zm.5-3.5v-1H9v1h1Zm3-.5h-1.5v1H13V4Zm-1.5 0h-2v1h2V4Zm-.5-.5v1h1v-1h-1Zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 2v1Zm-.5.5a.5.5 0 0 1 .5-.5V2A1.5 1.5 0 0 0 9 3.5h1Z"/>`
	zipSolidInnerSVG                         = `<path fill="currentColor" d="M3.296 11.379L3.89 9h1.22l.594 2.379A.5.5 0 0 1 5.22 12H3.781a.5.5 0 0 1-.485-.621ZM10.5 3a.5.5 0 0 0-.5.5V4h1v-.5a.5.5 0 0 0-.5-.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-12ZM3 3h3V2H3v1Zm0 2h3V4H3v1Zm3 2H3V6h3v1Zm-.11 1H3.11l-.784 3.136A1.5 1.5 0 0 0 3.78 13h1.438a1.5 1.5 0 0 0 1.455-1.864L5.89 8Zm3.374 2.06a.5.5 0 0 1 .513.024L12 11.566V10h1v2.5a.5.5 0 0 1-.777.416L10 11.434V13H9v-2.5a.5.5 0 0 1 .264-.44ZM9 6v3h1V8h2v1h1V6h-1v1h-2V6H9Zm3-2v-.5a1.5 1.5 0 0 0-3 0V5h4V4h-1Z" clip-rule="evenodd"/>`
	zoomInOutlineInnerSVG                    = `<path fill="none" stroke="currentColor" d="m14.5 14.5l-4-4M6.5 4v5M4 6.5h5m-2.5 6a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`
	zoomInSolidInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M0 6.5a6.5 6.5 0 1 1 11.436 4.23l3.418 3.416l-.707.708l-3.418-3.418A6.5 6.5 0 0 1 0 6.5ZM6 9V7H4V6h2V4h1v2h2v1H7v2H6Z" clip-rule="evenodd"/>`
	zoomOutOutlineInnerSVG                   = `<path fill="none" stroke="currentColor" d="m14.5 14.5l-4-4M4 6.5h5m-2.5 6a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`
	zoomOutSolidInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M0 6.5a6.5 6.5 0 1 1 11.436 4.23l3.418 3.416l-.707.708l-3.418-3.418A6.5 6.5 0 0 1 0 6.5ZM4 7h5V6H4v1Z" clip-rule="evenodd"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func AbTestingOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(abTestingOutlineInnerSVG).
		Element(children...)
}

func AbTestingSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(abTestingSolidInnerSVG).
		Element(children...)
}

func AddOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addOutlineInnerSVG).
		Element(children...)
}

func AddSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addSmallOutlineInnerSVG).
		Element(children...)
}

func AddSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addSmallSolidInnerSVG).
		Element(children...)
}

func AddSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addSolidInnerSVG).
		Element(children...)
}

func AddressBookOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addressBookOutlineInnerSVG).
		Element(children...)
}

func AddressBookSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addressBookSolidInnerSVG).
		Element(children...)
}

func AdjustHorizontalAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustHorizontalAltOutlineInnerSVG).
		Element(children...)
}

func AdjustHorizontalAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustHorizontalAltSolidInnerSVG).
		Element(children...)
}

func AdjustHorizontalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustHorizontalOutlineInnerSVG).
		Element(children...)
}

func AdjustHorizontalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustHorizontalSolidInnerSVG).
		Element(children...)
}

func AdjustVerticalAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustVerticalAltOutlineInnerSVG).
		Element(children...)
}

func AdjustVerticalAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustVerticalAltSolidInnerSVG).
		Element(children...)
}

func AdjustVerticalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustVerticalOutlineInnerSVG).
		Element(children...)
}

func AdjustVerticalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adjustVerticalSolidInnerSVG).
		Element(children...)
}

func AirplayOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplayOutlineInnerSVG).
		Element(children...)
}

func AirplaySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplaySolidInnerSVG).
		Element(children...)
}

func AirpodsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airpodsOutlineInnerSVG).
		Element(children...)
}

func AirpodsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airpodsSolidInnerSVG).
		Element(children...)
}

func AlarmOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alarmOutlineInnerSVG).
		Element(children...)
}

func AlarmSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alarmSolidInnerSVG).
		Element(children...)
}

func AlienOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alienOutlineInnerSVG).
		Element(children...)
}

func AlienSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alienSolidInnerSVG).
		Element(children...)
}

func AlignBottomOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignBottomOutlineInnerSVG).
		Element(children...)
}

func AlignBottomSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignBottomSolidInnerSVG).
		Element(children...)
}

func AlignCenterHorizontalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterHorizontalOutlineInnerSVG).
		Element(children...)
}

func AlignCenterHorizontalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterHorizontalSolidInnerSVG).
		Element(children...)
}

func AlignCenterVerticalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterVerticalOutlineInnerSVG).
		Element(children...)
}

func AlignCenterVerticalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterVerticalSolidInnerSVG).
		Element(children...)
}

func AlignLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignLeftOutlineInnerSVG).
		Element(children...)
}

func AlignLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignLeftSolidInnerSVG).
		Element(children...)
}

func AlignRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignRightOutlineInnerSVG).
		Element(children...)
}

func AlignRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignRightSolidInnerSVG).
		Element(children...)
}

func AlignTextCenterOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextCenterOutlineInnerSVG).
		Element(children...)
}

func AlignTextCenterSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextCenterSolidInnerSVG).
		Element(children...)
}

func AlignTextJustifyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextJustifyOutlineInnerSVG).
		Element(children...)
}

func AlignTextJustifySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextJustifySolidInnerSVG).
		Element(children...)
}

func AlignTextLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextLeftOutlineInnerSVG).
		Element(children...)
}

func AlignTextLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextLeftSolidInnerSVG).
		Element(children...)
}

func AlignTextRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextRightOutlineInnerSVG).
		Element(children...)
}

func AlignTextRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTextRightSolidInnerSVG).
		Element(children...)
}

func AlignTopOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTopOutlineInnerSVG).
		Element(children...)
}

func AlignTopSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTopSolidInnerSVG).
		Element(children...)
}

func AnchorOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(anchorOutlineInnerSVG).
		Element(children...)
}

func AnchorSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(anchorSolidInnerSVG).
		Element(children...)
}

func AndroidOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(androidOutlineInnerSVG).
		Element(children...)
}

func AndroidSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(androidSolidInnerSVG).
		Element(children...)
}

func AngularOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angularOutlineInnerSVG).
		Element(children...)
}

func AngularSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angularSolidInnerSVG).
		Element(children...)
}

func AnjaOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(anjaOutlineInnerSVG).
		Element(children...)
}

func AnjaSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(anjaSolidInnerSVG).
		Element(children...)
}

func AntiClockwiseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(antiClockwiseOutlineInnerSVG).
		Element(children...)
}

func AntiClockwiseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(antiClockwiseSolidInnerSVG).
		Element(children...)
}

func AppleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleOutlineInnerSVG).
		Element(children...)
}

func AppleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleSolidInnerSVG).
		Element(children...)
}

func AppointmentsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appointmentsOutlineInnerSVG).
		Element(children...)
}

func AppointmentsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appointmentsSolidInnerSVG).
		Element(children...)
}

func ArchiveOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archiveOutlineInnerSVG).
		Element(children...)
}

func ArchiveSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archiveSolidInnerSVG).
		Element(children...)
}

func AreaChartAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(areaChartAltOutlineInnerSVG).
		Element(children...)
}

func AreaChartAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(areaChartAltSolidInnerSVG).
		Element(children...)
}

func AreaChartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(areaChartOutlineInnerSVG).
		Element(children...)
}

func AreaChartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(areaChartSolidInnerSVG).
		Element(children...)
}

func ArrowDownCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownCircleOutlineInnerSVG).
		Element(children...)
}

func ArrowDownCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownCircleSolidInnerSVG).
		Element(children...)
}

func ArrowDownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownOutlineInnerSVG).
		Element(children...)
}

func ArrowDownSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownSmallOutlineInnerSVG).
		Element(children...)
}

func ArrowDownSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownSmallSolidInnerSVG).
		Element(children...)
}

func ArrowDownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownSolidInnerSVG).
		Element(children...)
}

func ArrowLeftCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftCircleOutlineInnerSVG).
		Element(children...)
}

func ArrowLeftCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftCircleSolidInnerSVG).
		Element(children...)
}

func ArrowLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftOutlineInnerSVG).
		Element(children...)
}

func ArrowLeftSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftSmallOutlineInnerSVG).
		Element(children...)
}

func ArrowLeftSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftSmallSolidInnerSVG).
		Element(children...)
}

func ArrowLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftSolidInnerSVG).
		Element(children...)
}

func ArrowOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowOutlineInnerSVG).
		Element(children...)
}

func ArrowRightCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightCircleOutlineInnerSVG).
		Element(children...)
}

func ArrowRightCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightCircleSolidInnerSVG).
		Element(children...)
}

func ArrowRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightOutlineInnerSVG).
		Element(children...)
}

func ArrowRightSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightSmallOutlineInnerSVG).
		Element(children...)
}

func ArrowRightSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightSmallSolidInnerSVG).
		Element(children...)
}

func ArrowRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightSolidInnerSVG).
		Element(children...)
}

func ArrowSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowSolidInnerSVG).
		Element(children...)
}

func ArrowUpCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpCircleOutlineInnerSVG).
		Element(children...)
}

func ArrowUpCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpCircleSolidInnerSVG).
		Element(children...)
}

func ArrowUpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpOutlineInnerSVG).
		Element(children...)
}

func ArrowUpSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpSmallOutlineInnerSVG).
		Element(children...)
}

func ArrowUpSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpSmallSolidInnerSVG).
		Element(children...)
}

func ArrowUpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpSolidInnerSVG).
		Element(children...)
}

func ArtboardOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(artboardOutlineInnerSVG).
		Element(children...)
}

func ArtboardSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(artboardSolidInnerSVG).
		Element(children...)
}

func AtOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atOutlineInnerSVG).
		Element(children...)
}

func AtSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atSolidInnerSVG).
		Element(children...)
}

func AttachOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(attachOutlineInnerSVG).
		Element(children...)
}

func AttachSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(attachSolidInnerSVG).
		Element(children...)
}

func AttachmentOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(attachmentOutlineInnerSVG).
		Element(children...)
}

func AttachmentSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(attachmentSolidInnerSVG).
		Element(children...)
}

func AudioCableOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(audioCableOutlineInnerSVG).
		Element(children...)
}

func AudioCableSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(audioCableSolidInnerSVG).
		Element(children...)
}

func AudioDocumentOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(audioDocumentOutlineInnerSVG).
		Element(children...)
}

func AudioDocumentSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(audioDocumentSolidInnerSVG).
		Element(children...)
}

func AzureOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(azureOutlineInnerSVG).
		Element(children...)
}

func AzureSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(azureSolidInnerSVG).
		Element(children...)
}

func BackspaceOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backspaceOutlineInnerSVG).
		Element(children...)
}

func BackspaceSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backspaceSolidInnerSVG).
		Element(children...)
}

func BagAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagAltOutlineInnerSVG).
		Element(children...)
}

func BagAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagAltSolidInnerSVG).
		Element(children...)
}

func BagMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagMinusOutlineInnerSVG).
		Element(children...)
}

func BagMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagMinusSolidInnerSVG).
		Element(children...)
}

func BagOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagOutlineInnerSVG).
		Element(children...)
}

func BagPlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagPlusOutlineInnerSVG).
		Element(children...)
}

func BagPlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagPlusSolidInnerSVG).
		Element(children...)
}

func BagSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagSolidInnerSVG).
		Element(children...)
}

func BankOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bankOutlineInnerSVG).
		Element(children...)
}

func BankSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bankSolidInnerSVG).
		Element(children...)
}

func BarChartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartOutlineInnerSVG).
		Element(children...)
}

func BarChartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartSolidInnerSVG).
		Element(children...)
}

func BarcodeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barcodeOutlineInnerSVG).
		Element(children...)
}

func BarcodeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barcodeSolidInnerSVG).
		Element(children...)
}

func BasketMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketMinusOutlineInnerSVG).
		Element(children...)
}

func BasketMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketMinusSolidInnerSVG).
		Element(children...)
}

func BasketOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketOutlineInnerSVG).
		Element(children...)
}

func BasketPlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketPlusOutlineInnerSVG).
		Element(children...)
}

func BasketPlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketPlusSolidInnerSVG).
		Element(children...)
}

func BasketSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketSolidInnerSVG).
		Element(children...)
}

func BathOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bathOutlineInnerSVG).
		Element(children...)
}

func BathSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bathSolidInnerSVG).
		Element(children...)
}

func BatteryChargeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryChargeOutlineInnerSVG).
		Element(children...)
}

func BatteryChargeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryChargeSolidInnerSVG).
		Element(children...)
}

func BatteryFiveOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFiveOutlineInnerSVG).
		Element(children...)
}

func BatteryFiveSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFiveSolidInnerSVG).
		Element(children...)
}

func BatteryFourOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFourOutlineInnerSVG).
		Element(children...)
}

func BatteryFourSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFourSolidInnerSVG).
		Element(children...)
}

func BatteryOneOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryOneOutlineInnerSVG).
		Element(children...)
}

func BatteryOneSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryOneSolidInnerSVG).
		Element(children...)
}

func BatteryThreeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryThreeOutlineInnerSVG).
		Element(children...)
}

func BatteryThreeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryThreeSolidInnerSVG).
		Element(children...)
}

func BatteryTwoOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryTwoOutlineInnerSVG).
		Element(children...)
}

func BatteryTwoSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryTwoSolidInnerSVG).
		Element(children...)
}

func BatteryZeroOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryZeroOutlineInnerSVG).
		Element(children...)
}

func BatteryZeroSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryZeroSolidInnerSVG).
		Element(children...)
}

func BedDoubleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bedDoubleOutlineInnerSVG).
		Element(children...)
}

func BedDoubleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bedDoubleSolidInnerSVG).
		Element(children...)
}

func BedSingleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bedSingleOutlineInnerSVG).
		Element(children...)
}

func BedSingleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bedSingleSolidInnerSVG).
		Element(children...)
}

func BehanceOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(behanceOutlineInnerSVG).
		Element(children...)
}

func BehanceSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(behanceSolidInnerSVG).
		Element(children...)
}

func BellOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellOutlineInnerSVG).
		Element(children...)
}

func BellSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellSolidInnerSVG).
		Element(children...)
}

func BinOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binOutlineInnerSVG).
		Element(children...)
}

func BinSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binSolidInnerSVG).
		Element(children...)
}

func BitbucketOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bitbucketOutlineInnerSVG).
		Element(children...)
}

func BitbucketSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bitbucketSolidInnerSVG).
		Element(children...)
}

func BitcoinOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bitcoinOutlineInnerSVG).
		Element(children...)
}

func BitcoinSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bitcoinSolidInnerSVG).
		Element(children...)
}

func BluetoothOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothOutlineInnerSVG).
		Element(children...)
}

func BluetoothSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothSolidInnerSVG).
		Element(children...)
}

func BoldOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boldOutlineInnerSVG).
		Element(children...)
}

func BoldSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boldSolidInnerSVG).
		Element(children...)
}

func BookOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookOutlineInnerSVG).
		Element(children...)
}

func BookSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookSolidInnerSVG).
		Element(children...)
}

func BookmarkOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkOutlineInnerSVG).
		Element(children...)
}

func BookmarkSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkSolidInnerSVG).
		Element(children...)
}

func BorderAllOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderAllOutlineInnerSVG).
		Element(children...)
}

func BorderAllSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderAllSolidInnerSVG).
		Element(children...)
}

func BorderBottomOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBottomOutlineInnerSVG).
		Element(children...)
}

func BorderBottomSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBottomSolidInnerSVG).
		Element(children...)
}

func BorderHorizontalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderHorizontalOutlineInnerSVG).
		Element(children...)
}

func BorderHorizontalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderHorizontalSolidInnerSVG).
		Element(children...)
}

func BorderInnerOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderInnerOutlineInnerSVG).
		Element(children...)
}

func BorderInnerSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderInnerSolidInnerSVG).
		Element(children...)
}

func BorderLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderLeftOutlineInnerSVG).
		Element(children...)
}

func BorderLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderLeftSolidInnerSVG).
		Element(children...)
}

func BorderNoneOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderNoneOutlineInnerSVG).
		Element(children...)
}

func BorderNoneSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderNoneSolidInnerSVG).
		Element(children...)
}

func BorderOuterOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderOuterOutlineInnerSVG).
		Element(children...)
}

func BorderOuterSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderOuterSolidInnerSVG).
		Element(children...)
}

func BorderRadiusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderRadiusOutlineInnerSVG).
		Element(children...)
}

func BorderRadiusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderRadiusSolidInnerSVG).
		Element(children...)
}

func BorderRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderRightOutlineInnerSVG).
		Element(children...)
}

func BorderRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderRightSolidInnerSVG).
		Element(children...)
}

func BorderTopOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopOutlineInnerSVG).
		Element(children...)
}

func BorderTopSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopSolidInnerSVG).
		Element(children...)
}

func BorderVerticalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderVerticalOutlineInnerSVG).
		Element(children...)
}

func BorderVerticalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderVerticalSolidInnerSVG).
		Element(children...)
}

func BottomLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bottomLeftOutlineInnerSVG).
		Element(children...)
}

func BottomLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bottomLeftSolidInnerSVG).
		Element(children...)
}

func BottomRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bottomRightOutlineInnerSVG).
		Element(children...)
}

func BottomRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bottomRightSolidInnerSVG).
		Element(children...)
}

func BoxOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOutlineInnerSVG).
		Element(children...)
}

func BoxSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxSolidInnerSVG).
		Element(children...)
}

func BracketOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bracketOutlineInnerSVG).
		Element(children...)
}

func BracketSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bracketSolidInnerSVG).
		Element(children...)
}

func BriefcaseAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseAltOutlineInnerSVG).
		Element(children...)
}

func BriefcaseAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseAltSolidInnerSVG).
		Element(children...)
}

func BriefcaseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseOutlineInnerSVG).
		Element(children...)
}

func BriefcaseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseSolidInnerSVG).
		Element(children...)
}

func BrushOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brushOutlineInnerSVG).
		Element(children...)
}

func BrushSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brushSolidInnerSVG).
		Element(children...)
}

func BugOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bugOutlineInnerSVG).
		Element(children...)
}

func BugSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bugSolidInnerSVG).
		Element(children...)
}

func BuildingOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buildingOutlineInnerSVG).
		Element(children...)
}

func BuildingSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buildingSolidInnerSVG).
		Element(children...)
}

func BulbOffOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bulbOffOutlineInnerSVG).
		Element(children...)
}

func BulbOffSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bulbOffSolidInnerSVG).
		Element(children...)
}

func BulbOnOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bulbOnOutlineInnerSVG).
		Element(children...)
}

func BulbOnSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bulbOnSolidInnerSVG).
		Element(children...)
}

func ButtonOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buttonOutlineInnerSVG).
		Element(children...)
}

func ButtonSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buttonSolidInnerSVG).
		Element(children...)
}

func COutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cOutlineInnerSVG).
		Element(children...)
}

func CSharpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cSharpOutlineInnerSVG).
		Element(children...)
}

func CSharpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cSharpSolidInnerSVG).
		Element(children...)
}

func CSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cSolidInnerSVG).
		Element(children...)
}

func CalculatorOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calculatorOutlineInnerSVG).
		Element(children...)
}

func CalculatorSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calculatorSolidInnerSVG).
		Element(children...)
}

func CalendarMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMinusOutlineInnerSVG).
		Element(children...)
}

func CalendarMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMinusSolidInnerSVG).
		Element(children...)
}

func CalendarNoAccessOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarNoAccessOutlineInnerSVG).
		Element(children...)
}

func CalendarNoAccessSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarNoAccessSolidInnerSVG).
		Element(children...)
}

func CalendarOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarOutlineInnerSVG).
		Element(children...)
}

func CalendarPlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarPlusOutlineInnerSVG).
		Element(children...)
}

func CalendarPlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarPlusSolidInnerSVG).
		Element(children...)
}

func CalendarSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarSolidInnerSVG).
		Element(children...)
}

func CalendarTickOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarTickOutlineInnerSVG).
		Element(children...)
}

func CalendarTickSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarTickSolidInnerSVG).
		Element(children...)
}

func CalendarXOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarXOutlineInnerSVG).
		Element(children...)
}

func CalendarXSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarXSolidInnerSVG).
		Element(children...)
}

func CameraOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraOutlineInnerSVG).
		Element(children...)
}

func CameraSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraSolidInnerSVG).
		Element(children...)
}

func CandleChartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(candleChartOutlineInnerSVG).
		Element(children...)
}

func CandleChartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(candleChartSolidInnerSVG).
		Element(children...)
}

func CarOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(carOutlineInnerSVG).
		Element(children...)
}

func CarSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(carSolidInnerSVG).
		Element(children...)
}

func CaretVerticalCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretVerticalCircleOutlineInnerSVG).
		Element(children...)
}

func CaretVerticalCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretVerticalCircleSolidInnerSVG).
		Element(children...)
}

func CaretVerticalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretVerticalOutlineInnerSVG).
		Element(children...)
}

func CaretVerticalSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretVerticalSmallOutlineInnerSVG).
		Element(children...)
}

func CaretVerticalSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretVerticalSmallSolidInnerSVG).
		Element(children...)
}

func CaretVerticalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretVerticalSolidInnerSVG).
		Element(children...)
}

func CartMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartMinusOutlineInnerSVG).
		Element(children...)
}

func CartMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartMinusSolidInnerSVG).
		Element(children...)
}

func CartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartOutlineInnerSVG).
		Element(children...)
}

func CartPlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartPlusOutlineInnerSVG).
		Element(children...)
}

func CartPlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartPlusSolidInnerSVG).
		Element(children...)
}

func CartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartSolidInnerSVG).
		Element(children...)
}

func CertificateOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(certificateOutlineInnerSVG).
		Element(children...)
}

func CertificateSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(certificateSolidInnerSVG).
		Element(children...)
}

func ChatOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatOutlineInnerSVG).
		Element(children...)
}

func ChatSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatSolidInnerSVG).
		Element(children...)
}

func ChatTypingAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatTypingAltOutlineInnerSVG).
		Element(children...)
}

func ChatTypingAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatTypingAltSolidInnerSVG).
		Element(children...)
}

func ChatTypingOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatTypingOutlineInnerSVG).
		Element(children...)
}

func ChatTypingSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatTypingSolidInnerSVG).
		Element(children...)
}

func ChatbotOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatbotOutlineInnerSVG).
		Element(children...)
}

func ChatbotSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatbotSolidInnerSVG).
		Element(children...)
}

func ChromeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chromeOutlineInnerSVG).
		Element(children...)
}

func ChromeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chromeSolidInnerSVG).
		Element(children...)
}

func ChurchOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(churchOutlineInnerSVG).
		Element(children...)
}

func ChurchSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(churchSolidInnerSVG).
		Element(children...)
}

func CircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleOutlineInnerSVG).
		Element(children...)
}

func CircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleSolidInnerSVG).
		Element(children...)
}

func ClipboardMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardMinusOutlineInnerSVG).
		Element(children...)
}

func ClipboardMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardMinusSolidInnerSVG).
		Element(children...)
}

func ClipboardNoAccessOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardNoAccessOutlineInnerSVG).
		Element(children...)
}

func ClipboardNoAccessSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardNoAccessSolidInnerSVG).
		Element(children...)
}

func ClipboardOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardOutlineInnerSVG).
		Element(children...)
}

func ClipboardPlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardPlusOutlineInnerSVG).
		Element(children...)
}

func ClipboardPlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardPlusSolidInnerSVG).
		Element(children...)
}

func ClipboardSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardSolidInnerSVG).
		Element(children...)
}

func ClipboardTickOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardTickOutlineInnerSVG).
		Element(children...)
}

func ClipboardTickSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardTickSolidInnerSVG).
		Element(children...)
}

func ClipboardXOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardXOutlineInnerSVG).
		Element(children...)
}

func ClipboardXSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardXSolidInnerSVG).
		Element(children...)
}

func ClockOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockOutlineInnerSVG).
		Element(children...)
}

func ClockSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockSolidInnerSVG).
		Element(children...)
}

func ClockwiseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockwiseOutlineInnerSVG).
		Element(children...)
}

func ClockwiseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockwiseSolidInnerSVG).
		Element(children...)
}

func CodeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeOutlineInnerSVG).
		Element(children...)
}

func CodeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeSolidInnerSVG).
		Element(children...)
}

func CodepenOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codepenOutlineInnerSVG).
		Element(children...)
}

func CodepenSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codepenSolidInnerSVG).
		Element(children...)
}

func CogOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cogOutlineInnerSVG).
		Element(children...)
}

func CogSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cogSolidInnerSVG).
		Element(children...)
}

func CompassOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassOutlineInnerSVG).
		Element(children...)
}

func CompassSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassSolidInnerSVG).
		Element(children...)
}

func ComputerOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(computerOutlineInnerSVG).
		Element(children...)
}

func ComputerSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(computerSolidInnerSVG).
		Element(children...)
}

func ContactOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contactOutlineInnerSVG).
		Element(children...)
}

func ContactSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contactSolidInnerSVG).
		Element(children...)
}

func ContractOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contractOutlineInnerSVG).
		Element(children...)
}

func ContractSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contractSolidInnerSVG).
		Element(children...)
}

func CostEstimateOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(costEstimateOutlineInnerSVG).
		Element(children...)
}

func CostEstimateSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(costEstimateSolidInnerSVG).
		Element(children...)
}

func CplusplusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cplusplusOutlineInnerSVG).
		Element(children...)
}

func CplusplusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cplusplusSolidInnerSVG).
		Element(children...)
}

func CreditCardOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardOutlineInnerSVG).
		Element(children...)
}

func CreditCardSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardSolidInnerSVG).
		Element(children...)
}

func CropOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropOutlineInnerSVG).
		Element(children...)
}

func CropSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropSolidInnerSVG).
		Element(children...)
}

func CssThreeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cssThreeOutlineInnerSVG).
		Element(children...)
}

func CssThreeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cssThreeSolidInnerSVG).
		Element(children...)
}

func CsvOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(csvOutlineInnerSVG).
		Element(children...)
}

func CsvSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(csvSolidInnerSVG).
		Element(children...)
}

func CupOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cupOutlineInnerSVG).
		Element(children...)
}

func CupSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cupSolidInnerSVG).
		Element(children...)
}

func CurvedConnectorOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(curvedConnectorOutlineInnerSVG).
		Element(children...)
}

func CurvedConnectorSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(curvedConnectorSolidInnerSVG).
		Element(children...)
}

func DThreeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dThreeOutlineInnerSVG).
		Element(children...)
}

func DThreeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dThreeSolidInnerSVG).
		Element(children...)
}

func DatabaseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseOutlineInnerSVG).
		Element(children...)
}

func DatabaseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseSolidInnerSVG).
		Element(children...)
}

func DeniedOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deniedOutlineInnerSVG).
		Element(children...)
}

func DeniedSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deniedSolidInnerSVG).
		Element(children...)
}

func DenoOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(denoOutlineInnerSVG).
		Element(children...)
}

func DenoSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(denoSolidInnerSVG).
		Element(children...)
}

func DepthChartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(depthChartOutlineInnerSVG).
		Element(children...)
}

func DepthChartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(depthChartSolidInnerSVG).
		Element(children...)
}

func DesklampOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(desklampOutlineInnerSVG).
		Element(children...)
}

func DesklampSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(desklampSolidInnerSVG).
		Element(children...)
}

func DiamondOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diamondOutlineInnerSVG).
		Element(children...)
}

func DiamondSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diamondSolidInnerSVG).
		Element(children...)
}

func DirectionOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(directionOutlineInnerSVG).
		Element(children...)
}

func DirectionSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(directionSolidInnerSVG).
		Element(children...)
}

func DiscordOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discordOutlineInnerSVG).
		Element(children...)
}

func DiscordSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discordSolidInnerSVG).
		Element(children...)
}

func DiscountOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discountOutlineInnerSVG).
		Element(children...)
}

func DiscountSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discountSolidInnerSVG).
		Element(children...)
}

func DistributeHorizontalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(distributeHorizontalOutlineInnerSVG).
		Element(children...)
}

func DistributeHorizontalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(distributeHorizontalSolidInnerSVG).
		Element(children...)
}

func DistributeVerticalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(distributeVerticalOutlineInnerSVG).
		Element(children...)
}

func DistributeVerticalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(distributeVerticalSolidInnerSVG).
		Element(children...)
}

func DividerLineOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dividerLineOutlineInnerSVG).
		Element(children...)
}

func DividerLineSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dividerLineSolidInnerSVG).
		Element(children...)
}

func DocOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(docOutlineInnerSVG).
		Element(children...)
}

func DocSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(docSolidInnerSVG).
		Element(children...)
}

func DockerOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dockerOutlineInnerSVG).
		Element(children...)
}

func DockerSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dockerSolidInnerSVG).
		Element(children...)
}

func DocumentsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentsOutlineInnerSVG).
		Element(children...)
}

func DocumentsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentsSolidInnerSVG).
		Element(children...)
}

func DollarOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dollarOutlineInnerSVG).
		Element(children...)
}

func DollarSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dollarSolidInnerSVG).
		Element(children...)
}

func DonutChartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(donutChartOutlineInnerSVG).
		Element(children...)
}

func DonutChartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(donutChartSolidInnerSVG).
		Element(children...)
}

func DoubleCaretDownCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretDownCircleOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretDownCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretDownCircleSolidInnerSVG).
		Element(children...)
}

func DoubleCaretDownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretDownOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretDownSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretDownSmallOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretDownSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretDownSmallSolidInnerSVG).
		Element(children...)
}

func DoubleCaretDownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretDownSolidInnerSVG).
		Element(children...)
}

func DoubleCaretLeftCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretLeftCircleOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretLeftCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretLeftCircleSolidInnerSVG).
		Element(children...)
}

func DoubleCaretLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretLeftOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretLeftSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretLeftSmallOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretLeftSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretLeftSmallSolidInnerSVG).
		Element(children...)
}

func DoubleCaretLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretLeftSolidInnerSVG).
		Element(children...)
}

func DoubleCaretRightCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretRightCircleOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretRightCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretRightCircleSolidInnerSVG).
		Element(children...)
}

func DoubleCaretRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretRightOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretRightSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretRightSmallOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretRightSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretRightSmallSolidInnerSVG).
		Element(children...)
}

func DoubleCaretRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretRightSolidInnerSVG).
		Element(children...)
}

func DoubleCaretUpCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretUpCircleOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretUpCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretUpCircleSolidInnerSVG).
		Element(children...)
}

func DoubleCaretUpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretUpOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretUpSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretUpSmallOutlineInnerSVG).
		Element(children...)
}

func DoubleCaretUpSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretUpSmallSolidInnerSVG).
		Element(children...)
}

func DoubleCaretUpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleCaretUpSolidInnerSVG).
		Element(children...)
}

func DownCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downCircleOutlineInnerSVG).
		Element(children...)
}

func DownCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downCircleSolidInnerSVG).
		Element(children...)
}

func DownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downOutlineInnerSVG).
		Element(children...)
}

func DownSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downSmallOutlineInnerSVG).
		Element(children...)
}

func DownSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downSmallSolidInnerSVG).
		Element(children...)
}

func DownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downSolidInnerSVG).
		Element(children...)
}

func DownloadOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadOutlineInnerSVG).
		Element(children...)
}

func DownloadSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadSolidInnerSVG).
		Element(children...)
}

func DragHorizontalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragHorizontalOutlineInnerSVG).
		Element(children...)
}

func DragHorizontalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragHorizontalSolidInnerSVG).
		Element(children...)
}

func DragOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragOutlineInnerSVG).
		Element(children...)
}

func DragSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragSolidInnerSVG).
		Element(children...)
}

func DragVerticalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragVerticalOutlineInnerSVG).
		Element(children...)
}

func DragVerticalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragVerticalSolidInnerSVG).
		Element(children...)
}

func DribbbleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dribbbleOutlineInnerSVG).
		Element(children...)
}

func DribbbleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dribbbleSolidInnerSVG).
		Element(children...)
}

func DropOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropOutlineInnerSVG).
		Element(children...)
}

func DropSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropSolidInnerSVG).
		Element(children...)
}

func DropperOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropperOutlineInnerSVG).
		Element(children...)
}

func DropperSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropperSolidInnerSVG).
		Element(children...)
}

func EdgeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(edgeOutlineInnerSVG).
		Element(children...)
}

func EdgeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(edgeSolidInnerSVG).
		Element(children...)
}

func EditCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editCircleOutlineInnerSVG).
		Element(children...)
}

func EditCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editCircleSolidInnerSVG).
		Element(children...)
}

func EditOneOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editOneOutlineInnerSVG).
		Element(children...)
}

func EditOneSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editOneSolidInnerSVG).
		Element(children...)
}

func EditOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editOutlineInnerSVG).
		Element(children...)
}

func EditSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editSmallOutlineInnerSVG).
		Element(children...)
}

func EditSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editSmallSolidInnerSVG).
		Element(children...)
}

func EditSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editSolidInnerSVG).
		Element(children...)
}

func ElbowConnectorOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(elbowConnectorOutlineInnerSVG).
		Element(children...)
}

func ElbowConnectorSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(elbowConnectorSolidInnerSVG).
		Element(children...)
}

func EnvelopeOpenOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeOpenOutlineInnerSVG).
		Element(children...)
}

func EnvelopeOpenSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeOpenSolidInnerSVG).
		Element(children...)
}

func EnvelopeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeOutlineInnerSVG).
		Element(children...)
}

func EnvelopeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeSolidInnerSVG).
		Element(children...)
}

func EpsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(epsOutlineInnerSVG).
		Element(children...)
}

func EpsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(epsSolidInnerSVG).
		Element(children...)
}

func EslintOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eslintOutlineInnerSVG).
		Element(children...)
}

func EslintSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eslintSolidInnerSVG).
		Element(children...)
}

func EthereumOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ethereumOutlineInnerSVG).
		Element(children...)
}

func EthereumSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ethereumSolidInnerSVG).
		Element(children...)
}

func EuroOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(euroOutlineInnerSVG).
		Element(children...)
}

func EuroSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(euroSolidInnerSVG).
		Element(children...)
}

func ExclamationCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationCircleOutlineInnerSVG).
		Element(children...)
}

func ExclamationCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationCircleSolidInnerSVG).
		Element(children...)
}

func ExclamationOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationOutlineInnerSVG).
		Element(children...)
}

func ExclamationSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationSmallOutlineInnerSVG).
		Element(children...)
}

func ExclamationSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationSmallSolidInnerSVG).
		Element(children...)
}

func ExclamationSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationSolidInnerSVG).
		Element(children...)
}

func ExpandAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandAltOutlineInnerSVG).
		Element(children...)
}

func ExpandAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandAltSolidInnerSVG).
		Element(children...)
}

func ExpandOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandOutlineInnerSVG).
		Element(children...)
}

func ExpandSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandSolidInnerSVG).
		Element(children...)
}

func EyeClosedOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeClosedOutlineInnerSVG).
		Element(children...)
}

func EyeClosedSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeClosedSolidInnerSVG).
		Element(children...)
}

func EyeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeOutlineInnerSVG).
		Element(children...)
}

func EyeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeSolidInnerSVG).
		Element(children...)
}

func FaceIdOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceIdOutlineInnerSVG).
		Element(children...)
}

func FaceIdSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceIdSolidInnerSVG).
		Element(children...)
}

func FacebookOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookOutlineInnerSVG).
		Element(children...)
}

func FacebookSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookSolidInnerSVG).
		Element(children...)
}

func FigmaOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(figmaOutlineInnerSVG).
		Element(children...)
}

func FigmaSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(figmaSolidInnerSVG).
		Element(children...)
}

func FileMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileMinusOutlineInnerSVG).
		Element(children...)
}

func FileMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileMinusSolidInnerSVG).
		Element(children...)
}

func FileNoAccessOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileNoAccessOutlineInnerSVG).
		Element(children...)
}

func FileNoAccessSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileNoAccessSolidInnerSVG).
		Element(children...)
}

func FileOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileOutlineInnerSVG).
		Element(children...)
}

func FilePlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filePlusOutlineInnerSVG).
		Element(children...)
}

func FilePlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filePlusSolidInnerSVG).
		Element(children...)
}

func FileSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileSolidInnerSVG).
		Element(children...)
}

func FileTickOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTickOutlineInnerSVG).
		Element(children...)
}

func FileTickSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTickSolidInnerSVG).
		Element(children...)
}

func FileXOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileXOutlineInnerSVG).
		Element(children...)
}

func FileXSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileXSolidInnerSVG).
		Element(children...)
}

func FilterOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterOutlineInnerSVG).
		Element(children...)
}

func FilterSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterSolidInnerSVG).
		Element(children...)
}

func FingerprintOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintOutlineInnerSVG).
		Element(children...)
}

func FingerprintSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintSolidInnerSVG).
		Element(children...)
}

func FirebaseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(firebaseOutlineInnerSVG).
		Element(children...)
}

func FirebaseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(firebaseSolidInnerSVG).
		Element(children...)
}

func FlagAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagAltOutlineInnerSVG).
		Element(children...)
}

func FlagAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagAltSolidInnerSVG).
		Element(children...)
}

func FlagOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagOutlineInnerSVG).
		Element(children...)
}

func FlagSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagSolidInnerSVG).
		Element(children...)
}

func FlipHorizontalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipHorizontalOutlineInnerSVG).
		Element(children...)
}

func FlipHorizontalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipHorizontalSolidInnerSVG).
		Element(children...)
}

func FlipVerticalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipVerticalOutlineInnerSVG).
		Element(children...)
}

func FlipVerticalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipVerticalSolidInnerSVG).
		Element(children...)
}

func FloatCenterOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatCenterOutlineInnerSVG).
		Element(children...)
}

func FloatCenterSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatCenterSolidInnerSVG).
		Element(children...)
}

func FloatLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatLeftOutlineInnerSVG).
		Element(children...)
}

func FloatLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatLeftSolidInnerSVG).
		Element(children...)
}

func FloatRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatRightOutlineInnerSVG).
		Element(children...)
}

func FloatRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatRightSolidInnerSVG).
		Element(children...)
}

func FloorplanOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floorplanOutlineInnerSVG).
		Element(children...)
}

func FloorplanSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floorplanSolidInnerSVG).
		Element(children...)
}

func FolderMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderMinusOutlineInnerSVG).
		Element(children...)
}

func FolderMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderMinusSolidInnerSVG).
		Element(children...)
}

func FolderNoAccessOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderNoAccessOutlineInnerSVG).
		Element(children...)
}

func FolderNoAccessSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderNoAccessSolidInnerSVG).
		Element(children...)
}

func FolderOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderOutlineInnerSVG).
		Element(children...)
}

func FolderPlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderPlusOutlineInnerSVG).
		Element(children...)
}

func FolderPlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderPlusSolidInnerSVG).
		Element(children...)
}

func FolderSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderSolidInnerSVG).
		Element(children...)
}

func FolderTickOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderTickOutlineInnerSVG).
		Element(children...)
}

func FolderTickSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderTickSolidInnerSVG).
		Element(children...)
}

func FolderXOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderXOutlineInnerSVG).
		Element(children...)
}

func FolderXSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderXSolidInnerSVG).
		Element(children...)
}

func FoldersOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(foldersOutlineInnerSVG).
		Element(children...)
}

func FoldersSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(foldersSolidInnerSVG).
		Element(children...)
}

func ForwardCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardCircleOutlineInnerSVG).
		Element(children...)
}

func ForwardCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardCircleSolidInnerSVG).
		Element(children...)
}

func ForwardOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardOutlineInnerSVG).
		Element(children...)
}

func ForwardSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardSmallOutlineInnerSVG).
		Element(children...)
}

func ForwardSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardSmallSolidInnerSVG).
		Element(children...)
}

func ForwardSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardSolidInnerSVG).
		Element(children...)
}

func FrameOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameOutlineInnerSVG).
		Element(children...)
}

func FrameSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameSolidInnerSVG).
		Element(children...)
}

func FramerOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(framerOutlineInnerSVG).
		Element(children...)
}

func FramerSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(framerSolidInnerSVG).
		Element(children...)
}

func GameControllerOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gameControllerOutlineInnerSVG).
		Element(children...)
}

func GameControllerRetroOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gameControllerRetroOutlineInnerSVG).
		Element(children...)
}

func GameControllerRetroSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gameControllerRetroSolidInnerSVG).
		Element(children...)
}

func GameControllerSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gameControllerSolidInnerSVG).
		Element(children...)
}

func GanttChartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ganttChartOutlineInnerSVG).
		Element(children...)
}

func GanttChartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ganttChartSolidInnerSVG).
		Element(children...)
}

func GarageOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(garageOutlineInnerSVG).
		Element(children...)
}

func GarageSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(garageSolidInnerSVG).
		Element(children...)
}

func GatsbyjsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gatsbyjsOutlineInnerSVG).
		Element(children...)
}

func GatsbyjsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gatsbyjsSolidInnerSVG).
		Element(children...)
}

func GbaOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbaOutlineInnerSVG).
		Element(children...)
}

func GbaSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbaSolidInnerSVG).
		Element(children...)
}

func GbcOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbcOutlineInnerSVG).
		Element(children...)
}

func GbcSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbcSolidInnerSVG).
		Element(children...)
}

func GhostOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ghostOutlineInnerSVG).
		Element(children...)
}

func GhostSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ghostSolidInnerSVG).
		Element(children...)
}

func GifOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gifOutlineInnerSVG).
		Element(children...)
}

func GifSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gifSolidInnerSVG).
		Element(children...)
}

func GiftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(giftOutlineInnerSVG).
		Element(children...)
}

func GiftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(giftSolidInnerSVG).
		Element(children...)
}

func GitBranchOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitBranchOutlineInnerSVG).
		Element(children...)
}

func GitBranchSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitBranchSolidInnerSVG).
		Element(children...)
}

func GitCommitOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCommitOutlineInnerSVG).
		Element(children...)
}

func GitCommitSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCommitSolidInnerSVG).
		Element(children...)
}

func GitCompareOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCompareOutlineInnerSVG).
		Element(children...)
}

func GitCompareSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCompareSolidInnerSVG).
		Element(children...)
}

func GitForkOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitForkOutlineInnerSVG).
		Element(children...)
}

func GitForkSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitForkSolidInnerSVG).
		Element(children...)
}

func GitMergeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitMergeOutlineInnerSVG).
		Element(children...)
}

func GitMergeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitMergeSolidInnerSVG).
		Element(children...)
}

func GitOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitOutlineInnerSVG).
		Element(children...)
}

func GitPullOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitPullOutlineInnerSVG).
		Element(children...)
}

func GitPullSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitPullSolidInnerSVG).
		Element(children...)
}

func GitSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitSolidInnerSVG).
		Element(children...)
}

func GithubOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(githubOutlineInnerSVG).
		Element(children...)
}

func GithubSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(githubSolidInnerSVG).
		Element(children...)
}

func GitlabOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitlabOutlineInnerSVG).
		Element(children...)
}

func GitlabSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitlabSolidInnerSVG).
		Element(children...)
}

func GlobeAfricaOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeAfricaOutlineInnerSVG).
		Element(children...)
}

func GlobeAfricaSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeAfricaSolidInnerSVG).
		Element(children...)
}

func GlobeAmericasOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeAmericasOutlineInnerSVG).
		Element(children...)
}

func GlobeAmericasSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeAmericasSolidInnerSVG).
		Element(children...)
}

func GlobeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeOutlineInnerSVG).
		Element(children...)
}

func GlobeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeSolidInnerSVG).
		Element(children...)
}

func GoogleAdOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleAdOutlineInnerSVG).
		Element(children...)
}

func GoogleAdSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleAdSolidInnerSVG).
		Element(children...)
}

func GoogleDriveOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleDriveOutlineInnerSVG).
		Element(children...)
}

func GoogleDriveSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleDriveSolidInnerSVG).
		Element(children...)
}

func GoogleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleOutlineInnerSVG).
		Element(children...)
}

func GooglePlayStoreOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googlePlayStoreOutlineInnerSVG).
		Element(children...)
}

func GooglePlayStoreSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googlePlayStoreSolidInnerSVG).
		Element(children...)
}

func GoogleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleSolidInnerSVG).
		Element(children...)
}

func GoogleStreetviewOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleStreetviewOutlineInnerSVG).
		Element(children...)
}

func GoogleStreetviewSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleStreetviewSolidInnerSVG).
		Element(children...)
}

func GraphqlOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphqlOutlineInnerSVG).
		Element(children...)
}

func GraphqlSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphqlSolidInnerSVG).
		Element(children...)
}

func GridLayoutOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridLayoutOutlineInnerSVG).
		Element(children...)
}

func GridLayoutSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridLayoutSolidInnerSVG).
		Element(children...)
}

func HashtagOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hashtagOutlineInnerSVG).
		Element(children...)
}

func HashtagSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hashtagSolidInnerSVG).
		Element(children...)
}

func HdScreenOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hdScreenOutlineInnerSVG).
		Element(children...)
}

func HdScreenSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hdScreenSolidInnerSVG).
		Element(children...)
}

func HdmiCableOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hdmiCableOutlineInnerSVG).
		Element(children...)
}

func HdmiCableSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hdmiCableSolidInnerSVG).
		Element(children...)
}

func HeadphonesOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headphonesOutlineInnerSVG).
		Element(children...)
}

func HeadphonesSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headphonesSolidInnerSVG).
		Element(children...)
}

func HeadsetOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headsetOutlineInnerSVG).
		Element(children...)
}

func HeadsetSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headsetSolidInnerSVG).
		Element(children...)
}

func HeartCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartCircleOutlineInnerSVG).
		Element(children...)
}

func HeartCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartCircleSolidInnerSVG).
		Element(children...)
}

func HeartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartOutlineInnerSVG).
		Element(children...)
}

func HeartSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartSmallOutlineInnerSVG).
		Element(children...)
}

func HeartSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartSmallSolidInnerSVG).
		Element(children...)
}

func HeartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartSolidInnerSVG).
		Element(children...)
}

func HexagonOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hexagonOutlineInnerSVG).
		Element(children...)
}

func HexagonSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hexagonSolidInnerSVG).
		Element(children...)
}

func HistoryOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(historyOutlineInnerSVG).
		Element(children...)
}

func HistorySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(historySolidInnerSVG).
		Element(children...)
}

func HomeAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeAltOutlineInnerSVG).
		Element(children...)
}

func HomeAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeAltSolidInnerSVG).
		Element(children...)
}

func HomeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeOutlineInnerSVG).
		Element(children...)
}

func HomeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeSolidInnerSVG).
		Element(children...)
}

func HospitalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hospitalOutlineInnerSVG).
		Element(children...)
}

func HospitalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hospitalSolidInnerSVG).
		Element(children...)
}

func HourglassOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hourglassOutlineInnerSVG).
		Element(children...)
}

func HourglassSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hourglassSolidInnerSVG).
		Element(children...)
}

func HouseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(houseOutlineInnerSVG).
		Element(children...)
}

func HouseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(houseSolidInnerSVG).
		Element(children...)
}

func HtmlFiveOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(htmlFiveOutlineInnerSVG).
		Element(children...)
}

func HtmlFiveSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(htmlFiveSolidInnerSVG).
		Element(children...)
}

func IdOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(idOutlineInnerSVG).
		Element(children...)
}

func IdSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(idSolidInnerSVG).
		Element(children...)
}

func ImacOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imacOutlineInnerSVG).
		Element(children...)
}

func ImacSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imacSolidInnerSVG).
		Element(children...)
}

func ImageAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageAltOutlineInnerSVG).
		Element(children...)
}

func ImageAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageAltSolidInnerSVG).
		Element(children...)
}

func ImageDocumentOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageDocumentOutlineInnerSVG).
		Element(children...)
}

func ImageDocumentSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageDocumentSolidInnerSVG).
		Element(children...)
}

func ImageOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageOutlineInnerSVG).
		Element(children...)
}

func ImageSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageSolidInnerSVG).
		Element(children...)
}

func InEarHeadphonesOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inEarHeadphonesOutlineInnerSVG).
		Element(children...)
}

func InEarHeadphonesSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inEarHeadphonesSolidInnerSVG).
		Element(children...)
}

func InboxOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxOutlineInnerSVG).
		Element(children...)
}

func InboxSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxSolidInnerSVG).
		Element(children...)
}

func IndentDecreaseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(indentDecreaseOutlineInnerSVG).
		Element(children...)
}

func IndentDecreaseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(indentDecreaseSolidInnerSVG).
		Element(children...)
}

func IndentIncreaseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(indentIncreaseOutlineInnerSVG).
		Element(children...)
}

func IndentIncreaseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(indentIncreaseSolidInnerSVG).
		Element(children...)
}

func InfoCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoCircleOutlineInnerSVG).
		Element(children...)
}

func InfoCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoCircleSolidInnerSVG).
		Element(children...)
}

func InfoOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoOutlineInnerSVG).
		Element(children...)
}

func InfoSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoSmallOutlineInnerSVG).
		Element(children...)
}

func InfoSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoSmallSolidInnerSVG).
		Element(children...)
}

func InfoSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoSolidInnerSVG).
		Element(children...)
}

func InstagramOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(instagramOutlineInnerSVG).
		Element(children...)
}

func InstagramSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(instagramSolidInnerSVG).
		Element(children...)
}

func InvoiceOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(invoiceOutlineInnerSVG).
		Element(children...)
}

func InvoiceSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(invoiceSolidInnerSVG).
		Element(children...)
}

func ItalicOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(italicOutlineInnerSVG).
		Element(children...)
}

func ItalicSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(italicSolidInnerSVG).
		Element(children...)
}

func JavascriptOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(javascriptOutlineInnerSVG).
		Element(children...)
}

func JavascriptSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(javascriptSolidInnerSVG).
		Element(children...)
}

func JoystickOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(joystickOutlineInnerSVG).
		Element(children...)
}

func JoystickSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(joystickSolidInnerSVG).
		Element(children...)
}

func JpgOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jpgOutlineInnerSVG).
		Element(children...)
}

func JpgSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jpgSolidInnerSVG).
		Element(children...)
}

func KanbanOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kanbanOutlineInnerSVG).
		Element(children...)
}

func KanbanSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kanbanSolidInnerSVG).
		Element(children...)
}

func KeyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyOutlineInnerSVG).
		Element(children...)
}

func KeySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keySolidInnerSVG).
		Element(children...)
}

func KeyboardOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyboardOutlineInnerSVG).
		Element(children...)
}

func KeyboardSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyboardSolidInnerSVG).
		Element(children...)
}

func LanCableOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lanCableOutlineInnerSVG).
		Element(children...)
}

func LanCableSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lanCableSolidInnerSVG).
		Element(children...)
}

func LaptopOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopOutlineInnerSVG).
		Element(children...)
}

func LaptopSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopSolidInnerSVG).
		Element(children...)
}

func LaravelOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laravelOutlineInnerSVG).
		Element(children...)
}

func LaravelSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laravelSolidInnerSVG).
		Element(children...)
}

func LayersDifferenceOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersDifferenceOutlineInnerSVG).
		Element(children...)
}

func LayersDifferenceSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersDifferenceSolidInnerSVG).
		Element(children...)
}

func LayersIntersectOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersIntersectOutlineInnerSVG).
		Element(children...)
}

func LayersIntersectSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersIntersectSolidInnerSVG).
		Element(children...)
}

func LayersOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersOutlineInnerSVG).
		Element(children...)
}

func LayersSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersSolidInnerSVG).
		Element(children...)
}

func LayersSubtractOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersSubtractOutlineInnerSVG).
		Element(children...)
}

func LayersSubtractSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersSubtractSolidInnerSVG).
		Element(children...)
}

func LayersUnionOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersUnionOutlineInnerSVG).
		Element(children...)
}

func LayersUnionSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersUnionSolidInnerSVG).
		Element(children...)
}

func LeftCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftCircleOutlineInnerSVG).
		Element(children...)
}

func LeftCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftCircleSolidInnerSVG).
		Element(children...)
}

func LeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftOutlineInnerSVG).
		Element(children...)
}

func LeftSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftSmallOutlineInnerSVG).
		Element(children...)
}

func LeftSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftSmallSolidInnerSVG).
		Element(children...)
}

func LeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftSolidInnerSVG).
		Element(children...)
}

func LegoOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(legoOutlineInnerSVG).
		Element(children...)
}

func LegoSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(legoSolidInnerSVG).
		Element(children...)
}

func LifebuoyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lifebuoyOutlineInnerSVG).
		Element(children...)
}

func LifebuoySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lifebuoySolidInnerSVG).
		Element(children...)
}

func LightningCableOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightningCableOutlineInnerSVG).
		Element(children...)
}

func LightningCableSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightningCableSolidInnerSVG).
		Element(children...)
}

func LineOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineOutlineInnerSVG).
		Element(children...)
}

func LineSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineSolidInnerSVG).
		Element(children...)
}

func LinkOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkOutlineInnerSVG).
		Element(children...)
}

func LinkRemoveOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkRemoveOutlineInnerSVG).
		Element(children...)
}

func LinkRemoveSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkRemoveSolidInnerSVG).
		Element(children...)
}

func LinkSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkSolidInnerSVG).
		Element(children...)
}

func LinkedinOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkedinOutlineInnerSVG).
		Element(children...)
}

func LinkedinSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkedinSolidInnerSVG).
		Element(children...)
}

func LinuxAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linuxAltOutlineInnerSVG).
		Element(children...)
}

func LinuxAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linuxAltSolidInnerSVG).
		Element(children...)
}

func LinuxOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linuxOutlineInnerSVG).
		Element(children...)
}

func LinuxSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linuxSolidInnerSVG).
		Element(children...)
}

func ListLayoutOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listLayoutOutlineInnerSVG).
		Element(children...)
}

func ListLayoutSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listLayoutSolidInnerSVG).
		Element(children...)
}

func ListOrderedOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listOrderedOutlineInnerSVG).
		Element(children...)
}

func ListOrderedSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listOrderedSolidInnerSVG).
		Element(children...)
}

func ListUnorderedOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listUnorderedOutlineInnerSVG).
		Element(children...)
}

func ListUnorderedSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listUnorderedSolidInnerSVG).
		Element(children...)
}

func LitecoinOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(litecoinOutlineInnerSVG).
		Element(children...)
}

func LitecoinSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(litecoinSolidInnerSVG).
		Element(children...)
}

func LoaderOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loaderOutlineInnerSVG).
		Element(children...)
}

func LoaderSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loaderSolidInnerSVG).
		Element(children...)
}

func LocationOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationOutlineInnerSVG).
		Element(children...)
}

func LocationSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationSolidInnerSVG).
		Element(children...)
}

func LockCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockCircleOutlineInnerSVG).
		Element(children...)
}

func LockCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockCircleSolidInnerSVG).
		Element(children...)
}

func LockOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockOutlineInnerSVG).
		Element(children...)
}

func LockSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockSmallOutlineInnerSVG).
		Element(children...)
}

func LockSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockSmallSolidInnerSVG).
		Element(children...)
}

func LockSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockSolidInnerSVG).
		Element(children...)
}

func LogoutOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(logoutOutlineInnerSVG).
		Element(children...)
}

func LogoutSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(logoutSolidInnerSVG).
		Element(children...)
}

func LoopOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loopOutlineInnerSVG).
		Element(children...)
}

func LoopSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loopSolidInnerSVG).
		Element(children...)
}

func MagsafeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magsafeOutlineInnerSVG).
		Element(children...)
}

func MagsafeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magsafeSolidInnerSVG).
		Element(children...)
}

func MarkdownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(markdownOutlineInnerSVG).
		Element(children...)
}

func MarkdownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(markdownSolidInnerSVG).
		Element(children...)
}

func MediumOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediumOutlineInnerSVG).
		Element(children...)
}

func MediumSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediumSolidInnerSVG).
		Element(children...)
}

func MenuOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuOutlineInnerSVG).
		Element(children...)
}

func MenuSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuSolidInnerSVG).
		Element(children...)
}

func MessageMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageMinusOutlineInnerSVG).
		Element(children...)
}

func MessageMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageMinusSolidInnerSVG).
		Element(children...)
}

func MessageNoAccessOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageNoAccessOutlineInnerSVG).
		Element(children...)
}

func MessageNoAccessSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageNoAccessSolidInnerSVG).
		Element(children...)
}

func MessageOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageOutlineInnerSVG).
		Element(children...)
}

func MessagePlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messagePlusOutlineInnerSVG).
		Element(children...)
}

func MessagePlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messagePlusSolidInnerSVG).
		Element(children...)
}

func MessageSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageSolidInnerSVG).
		Element(children...)
}

func MessageTextAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageTextAltOutlineInnerSVG).
		Element(children...)
}

func MessageTextAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageTextAltSolidInnerSVG).
		Element(children...)
}

func MessageTextOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageTextOutlineInnerSVG).
		Element(children...)
}

func MessageTextSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageTextSolidInnerSVG).
		Element(children...)
}

func MessageTickOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageTickOutlineInnerSVG).
		Element(children...)
}

func MessageTickSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageTickSolidInnerSVG).
		Element(children...)
}

func MessageXOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageXOutlineInnerSVG).
		Element(children...)
}

func MessageXSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageXSolidInnerSVG).
		Element(children...)
}

func MessengerOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messengerOutlineInnerSVG).
		Element(children...)
}

func MessengerSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messengerSolidInnerSVG).
		Element(children...)
}

func MicroSdCardOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microSdCardOutlineInnerSVG).
		Element(children...)
}

func MicroSdCardSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microSdCardSolidInnerSVG).
		Element(children...)
}

func MicrophoneOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneOutlineInnerSVG).
		Element(children...)
}

func MicrophoneSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneSolidInnerSVG).
		Element(children...)
}

func MinimiseAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minimiseAltOutlineInnerSVG).
		Element(children...)
}

func MinimiseAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minimiseAltSolidInnerSVG).
		Element(children...)
}

func MinimiseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minimiseOutlineInnerSVG).
		Element(children...)
}

func MinimiseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minimiseSolidInnerSVG).
		Element(children...)
}

func MinusCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircleOutlineInnerSVG).
		Element(children...)
}

func MinusCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircleSolidInnerSVG).
		Element(children...)
}

func MinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusOutlineInnerSVG).
		Element(children...)
}

func MinusSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusSmallOutlineInnerSVG).
		Element(children...)
}

func MinusSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusSmallSolidInnerSVG).
		Element(children...)
}

func MinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusSolidInnerSVG).
		Element(children...)
}

func MobileOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mobileOutlineInnerSVG).
		Element(children...)
}

func MobileSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mobileSolidInnerSVG).
		Element(children...)
}

func MoneyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneyOutlineInnerSVG).
		Element(children...)
}

func MoneySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneySolidInnerSVG).
		Element(children...)
}

func MoneyStackOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneyStackOutlineInnerSVG).
		Element(children...)
}

func MoneyStackSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moneyStackSolidInnerSVG).
		Element(children...)
}

func MongodbOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mongodbOutlineInnerSVG).
		Element(children...)
}

func MongodbSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mongodbSolidInnerSVG).
		Element(children...)
}

func MoodFlatOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodFlatOutlineInnerSVG).
		Element(children...)
}

func MoodFlatSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodFlatSolidInnerSVG).
		Element(children...)
}

func MoodFrownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodFrownOutlineInnerSVG).
		Element(children...)
}

func MoodFrownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodFrownSolidInnerSVG).
		Element(children...)
}

func MoodLaughOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodLaughOutlineInnerSVG).
		Element(children...)
}

func MoodLaughSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodLaughSolidInnerSVG).
		Element(children...)
}

func MoodSadOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodSadOutlineInnerSVG).
		Element(children...)
}

func MoodSadSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodSadSolidInnerSVG).
		Element(children...)
}

func MoodSmileOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodSmileOutlineInnerSVG).
		Element(children...)
}

func MoodSmileSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodSmileSolidInnerSVG).
		Element(children...)
}

func MoodSurprisedOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodSurprisedOutlineInnerSVG).
		Element(children...)
}

func MoodSurprisedSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodSurprisedSolidInnerSVG).
		Element(children...)
}

func MoodTongueOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodTongueOutlineInnerSVG).
		Element(children...)
}

func MoodTongueSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodTongueSolidInnerSVG).
		Element(children...)
}

func MoonOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonOutlineInnerSVG).
		Element(children...)
}

func MoonSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonSolidInnerSVG).
		Element(children...)
}

func MoreHorizontalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreHorizontalOutlineInnerSVG).
		Element(children...)
}

func MoreHorizontalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreHorizontalSolidInnerSVG).
		Element(children...)
}

func MoreVerticalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreVerticalOutlineInnerSVG).
		Element(children...)
}

func MoreVerticalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreVerticalSolidInnerSVG).
		Element(children...)
}

func MouseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mouseOutlineInnerSVG).
		Element(children...)
}

func MouseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mouseSolidInnerSVG).
		Element(children...)
}

func MovOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(movOutlineInnerSVG).
		Element(children...)
}

func MovSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(movSolidInnerSVG).
		Element(children...)
}

func MpFourOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mpFourOutlineInnerSVG).
		Element(children...)
}

func MpFourSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mpFourSolidInnerSVG).
		Element(children...)
}

func MpThreeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mpThreeOutlineInnerSVG).
		Element(children...)
}

func MpThreeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mpThreeSolidInnerSVG).
		Element(children...)
}

func MsExcelOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(msExcelOutlineInnerSVG).
		Element(children...)
}

func MsExcelSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(msExcelSolidInnerSVG).
		Element(children...)
}

func MsPowerpointOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(msPowerpointOutlineInnerSVG).
		Element(children...)
}

func MsPowerpointSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(msPowerpointSolidInnerSVG).
		Element(children...)
}

func MsWordOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(msWordOutlineInnerSVG).
		Element(children...)
}

func MsWordSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(msWordSolidInnerSVG).
		Element(children...)
}

func NSixtyFourOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nSixtyFourOutlineInnerSVG).
		Element(children...)
}

func NSixtyFourSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nSixtyFourSolidInnerSVG).
		Element(children...)
}

func NesOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nesOutlineInnerSVG).
		Element(children...)
}

func NesSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nesSolidInnerSVG).
		Element(children...)
}

func NetlifyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(netlifyOutlineInnerSVG).
		Element(children...)
}

func NetlifySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(netlifySolidInnerSVG).
		Element(children...)
}

func NextCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextCircleOutlineInnerSVG).
		Element(children...)
}

func NextCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextCircleSolidInnerSVG).
		Element(children...)
}

func NextOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextOutlineInnerSVG).
		Element(children...)
}

func NextSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextSmallOutlineInnerSVG).
		Element(children...)
}

func NextSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextSmallSolidInnerSVG).
		Element(children...)
}

func NextSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextSolidInnerSVG).
		Element(children...)
}

func NextjsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextjsOutlineInnerSVG).
		Element(children...)
}

func NextjsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextjsSolidInnerSVG).
		Element(children...)
}

func NgcOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ngcOutlineInnerSVG).
		Element(children...)
}

func NgcSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ngcSolidInnerSVG).
		Element(children...)
}

func NintendoSwitchOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nintendoSwitchOutlineInnerSVG).
		Element(children...)
}

func NintendoSwitchSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nintendoSwitchSolidInnerSVG).
		Element(children...)
}

func NodejsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nodejsOutlineInnerSVG).
		Element(children...)
}

func NodejsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nodejsSolidInnerSVG).
		Element(children...)
}

func NoteOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noteOutlineInnerSVG).
		Element(children...)
}

func NoteSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noteSolidInnerSVG).
		Element(children...)
}

func NpmOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(npmOutlineInnerSVG).
		Element(children...)
}

func NpmSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(npmSolidInnerSVG).
		Element(children...)
}

func NuxtjsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nuxtjsOutlineInnerSVG).
		Element(children...)
}

func NuxtjsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nuxtjsSolidInnerSVG).
		Element(children...)
}

func OmegaOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(omegaOutlineInnerSVG).
		Element(children...)
}

func OmegaSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(omegaSolidInnerSVG).
		Element(children...)
}

func OperaOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(operaOutlineInnerSVG).
		Element(children...)
}

func OperaSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(operaSolidInnerSVG).
		Element(children...)
}

func OtpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(otpOutlineInnerSVG).
		Element(children...)
}

func OtpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(otpSolidInnerSVG).
		Element(children...)
}

func PageBreakOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageBreakOutlineInnerSVG).
		Element(children...)
}

func PageBreakSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageBreakSolidInnerSVG).
		Element(children...)
}

func PageNumberOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageNumberOutlineInnerSVG).
		Element(children...)
}

func PageNumberSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pageNumberSolidInnerSVG).
		Element(children...)
}

func PaintbrushOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paintbrushOutlineInnerSVG).
		Element(children...)
}

func PaintbrushSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paintbrushSolidInnerSVG).
		Element(children...)
}

func PaintbucketOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paintbucketOutlineInnerSVG).
		Element(children...)
}

func PaintbucketSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paintbucketSolidInnerSVG).
		Element(children...)
}

func ParagraphOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphOutlineInnerSVG).
		Element(children...)
}

func ParagraphSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphSolidInnerSVG).
		Element(children...)
}

func PasswordOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(passwordOutlineInnerSVG).
		Element(children...)
}

func PasswordSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(passwordSolidInnerSVG).
		Element(children...)
}

func PatreonOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(patreonOutlineInnerSVG).
		Element(children...)
}

func PatreonSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(patreonSolidInnerSVG).
		Element(children...)
}

func PauseCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseCircleOutlineInnerSVG).
		Element(children...)
}

func PauseCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseCircleSolidInnerSVG).
		Element(children...)
}

func PauseOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseOutlineInnerSVG).
		Element(children...)
}

func PauseSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseSmallOutlineInnerSVG).
		Element(children...)
}

func PauseSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseSmallSolidInnerSVG).
		Element(children...)
}

func PauseSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseSolidInnerSVG).
		Element(children...)
}

func PawOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pawOutlineInnerSVG).
		Element(children...)
}

func PawSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pawSolidInnerSVG).
		Element(children...)
}

func PawsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pawsOutlineInnerSVG).
		Element(children...)
}

func PawsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pawsSolidInnerSVG).
		Element(children...)
}

func PaypalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paypalOutlineInnerSVG).
		Element(children...)
}

func PaypalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paypalSolidInnerSVG).
		Element(children...)
}

func PdfOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pdfOutlineInnerSVG).
		Element(children...)
}

func PdfSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pdfSolidInnerSVG).
		Element(children...)
}

func PenOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penOutlineInnerSVG).
		Element(children...)
}

func PenSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penSolidInnerSVG).
		Element(children...)
}

func PhoneOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneOutlineInnerSVG).
		Element(children...)
}

func PhoneSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneSolidInnerSVG).
		Element(children...)
}

func PhonecallBlockedOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phonecallBlockedOutlineInnerSVG).
		Element(children...)
}

func PhonecallBlockedSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phonecallBlockedSolidInnerSVG).
		Element(children...)
}

func PhonecallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phonecallOutlineInnerSVG).
		Element(children...)
}

func PhonecallReceiveOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phonecallReceiveOutlineInnerSVG).
		Element(children...)
}

func PhonecallReceiveSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phonecallReceiveSolidInnerSVG).
		Element(children...)
}

func PhonecallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phonecallSolidInnerSVG).
		Element(children...)
}

func PieChartAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartAltOutlineInnerSVG).
		Element(children...)
}

func PieChartAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartAltSolidInnerSVG).
		Element(children...)
}

func PieChartOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartOutlineInnerSVG).
		Element(children...)
}

func PieChartSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartSolidInnerSVG).
		Element(children...)
}

func PinAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinAltOutlineInnerSVG).
		Element(children...)
}

func PinAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinAltSolidInnerSVG).
		Element(children...)
}

func PinOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinOutlineInnerSVG).
		Element(children...)
}

func PinSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinSolidInnerSVG).
		Element(children...)
}

func PinterestOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinterestOutlineInnerSVG).
		Element(children...)
}

func PinterestSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinterestSolidInnerSVG).
		Element(children...)
}

func PlantOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plantOutlineInnerSVG).
		Element(children...)
}

func PlantSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plantSolidInnerSVG).
		Element(children...)
}

func PlayCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playCircleOutlineInnerSVG).
		Element(children...)
}

func PlayCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playCircleSolidInnerSVG).
		Element(children...)
}

func PlayOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playOutlineInnerSVG).
		Element(children...)
}

func PlaySmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playSmallOutlineInnerSVG).
		Element(children...)
}

func PlaySmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playSmallSolidInnerSVG).
		Element(children...)
}

func PlaySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playSolidInnerSVG).
		Element(children...)
}

func PlugOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plugOutlineInnerSVG).
		Element(children...)
}

func PlugSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plugSolidInnerSVG).
		Element(children...)
}

func PlusCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircleOutlineInnerSVG).
		Element(children...)
}

func PlusCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircleSolidInnerSVG).
		Element(children...)
}

func PngOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pngOutlineInnerSVG).
		Element(children...)
}

func PngSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pngSolidInnerSVG).
		Element(children...)
}

func PoolOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(poolOutlineInnerSVG).
		Element(children...)
}

func PoolSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(poolSolidInnerSVG).
		Element(children...)
}

func PoundOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(poundOutlineInnerSVG).
		Element(children...)
}

func PoundSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(poundSolidInnerSVG).
		Element(children...)
}

func PowerOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerOutlineInnerSVG).
		Element(children...)
}

func PowerSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerSolidInnerSVG).
		Element(children...)
}

func PptOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pptOutlineInnerSVG).
		Element(children...)
}

func PptSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pptSolidInnerSVG).
		Element(children...)
}

func PrintOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printOutlineInnerSVG).
		Element(children...)
}

func PrintSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printSolidInnerSVG).
		Element(children...)
}

func PythonOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pythonOutlineInnerSVG).
		Element(children...)
}

func PythonSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pythonSolidInnerSVG).
		Element(children...)
}

func QrCodeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(qrCodeOutlineInnerSVG).
		Element(children...)
}

func QrCodeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(qrCodeSolidInnerSVG).
		Element(children...)
}

func QuestionCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionCircleOutlineInnerSVG).
		Element(children...)
}

func QuestionCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionCircleSolidInnerSVG).
		Element(children...)
}

func QuestionOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionOutlineInnerSVG).
		Element(children...)
}

func QuestionSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionSmallOutlineInnerSVG).
		Element(children...)
}

func QuestionSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionSmallSolidInnerSVG).
		Element(children...)
}

func QuestionSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionSolidInnerSVG).
		Element(children...)
}

func QuoteOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quoteOutlineInnerSVG).
		Element(children...)
}

func QuoteSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quoteSolidInnerSVG).
		Element(children...)
}

func RandOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(randOutlineInnerSVG).
		Element(children...)
}

func RandSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(randSolidInnerSVG).
		Element(children...)
}

func ReactOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reactOutlineInnerSVG).
		Element(children...)
}

func ReactSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reactSolidInnerSVG).
		Element(children...)
}

func ReceiptOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(receiptOutlineInnerSVG).
		Element(children...)
}

func ReceiptSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(receiptSolidInnerSVG).
		Element(children...)
}

func RedditOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redditOutlineInnerSVG).
		Element(children...)
}

func RedditSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redditSolidInnerSVG).
		Element(children...)
}

func RedwoodjsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redwoodjsOutlineInnerSVG).
		Element(children...)
}

func RedwoodjsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redwoodjsSolidInnerSVG).
		Element(children...)
}

func RefreshAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshAltOutlineInnerSVG).
		Element(children...)
}

func RefreshAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshAltSolidInnerSVG).
		Element(children...)
}

func RefreshOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshOutlineInnerSVG).
		Element(children...)
}

func RefreshSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshSolidInnerSVG).
		Element(children...)
}

func RewindCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rewindCircleOutlineInnerSVG).
		Element(children...)
}

func RewindCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rewindCircleSolidInnerSVG).
		Element(children...)
}

func RewindOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rewindOutlineInnerSVG).
		Element(children...)
}

func RewindSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rewindSmallOutlineInnerSVG).
		Element(children...)
}

func RewindSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rewindSmallSolidInnerSVG).
		Element(children...)
}

func RewindSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rewindSolidInnerSVG).
		Element(children...)
}

func RightCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rightCircleOutlineInnerSVG).
		Element(children...)
}

func RightCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rightCircleSolidInnerSVG).
		Element(children...)
}

func RightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rightOutlineInnerSVG).
		Element(children...)
}

func RightSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rightSmallOutlineInnerSVG).
		Element(children...)
}

func RightSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rightSmallSolidInnerSVG).
		Element(children...)
}

func RightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rightSolidInnerSVG).
		Element(children...)
}

func RippleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rippleOutlineInnerSVG).
		Element(children...)
}

func RippleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rippleSolidInnerSVG).
		Element(children...)
}

func RobotOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(robotOutlineInnerSVG).
		Element(children...)
}

func RobotSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(robotSolidInnerSVG).
		Element(children...)
}

func RollerOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rollerOutlineInnerSVG).
		Element(children...)
}

func RollerSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rollerSolidInnerSVG).
		Element(children...)
}

func RollupjsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rollupjsOutlineInnerSVG).
		Element(children...)
}

func RollupjsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rollupjsSolidInnerSVG).
		Element(children...)
}

func RouterOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(routerOutlineInnerSVG).
		Element(children...)
}

func RouterSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(routerSolidInnerSVG).
		Element(children...)
}

func RssOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssOutlineInnerSVG).
		Element(children...)
}

func RssSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssSolidInnerSVG).
		Element(children...)
}

func RubyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rubyOutlineInnerSVG).
		Element(children...)
}

func RubySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rubySolidInnerSVG).
		Element(children...)
}

func RupeeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rupeeOutlineInnerSVG).
		Element(children...)
}

func RupeeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rupeeSolidInnerSVG).
		Element(children...)
}

func RustOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rustOutlineInnerSVG).
		Element(children...)
}

func RustSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rustSolidInnerSVG).
		Element(children...)
}

func SafariOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(safariOutlineInnerSVG).
		Element(children...)
}

func SafariSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(safariSolidInnerSVG).
		Element(children...)
}

func SafeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(safeOutlineInnerSVG).
		Element(children...)
}

func SafeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(safeSolidInnerSVG).
		Element(children...)
}

func SaveOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(saveOutlineInnerSVG).
		Element(children...)
}

func SaveSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(saveSolidInnerSVG).
		Element(children...)
}

func ScanOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scanOutlineInnerSVG).
		Element(children...)
}

func ScanSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scanSolidInnerSVG).
		Element(children...)
}

func SchoolOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(schoolOutlineInnerSVG).
		Element(children...)
}

func SchoolSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(schoolSolidInnerSVG).
		Element(children...)
}

func ScreenAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenAltOutlineInnerSVG).
		Element(children...)
}

func ScreenAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenAltSolidInnerSVG).
		Element(children...)
}

func ScreenAltTwoOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenAltTwoOutlineInnerSVG).
		Element(children...)
}

func ScreenAltTwoSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenAltTwoSolidInnerSVG).
		Element(children...)
}

func ScreenOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenOutlineInnerSVG).
		Element(children...)
}

func ScreenSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenSolidInnerSVG).
		Element(children...)
}

func ScribbleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scribbleOutlineInnerSVG).
		Element(children...)
}

func ScribbleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scribbleSolidInnerSVG).
		Element(children...)
}

func SdCardOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sdCardOutlineInnerSVG).
		Element(children...)
}

func SdCardSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sdCardSolidInnerSVG).
		Element(children...)
}

func SearchCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchCircleOutlineInnerSVG).
		Element(children...)
}

func SearchCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchCircleSolidInnerSVG).
		Element(children...)
}

func SearchOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchOutlineInnerSVG).
		Element(children...)
}

func SearchPropertyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchPropertyOutlineInnerSVG).
		Element(children...)
}

func SearchPropertySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchPropertySolidInnerSVG).
		Element(children...)
}

func SearchSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchSmallOutlineInnerSVG).
		Element(children...)
}

func SearchSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchSmallSolidInnerSVG).
		Element(children...)
}

func SearchSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchSolidInnerSVG).
		Element(children...)
}

func SectionAddOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionAddOutlineInnerSVG).
		Element(children...)
}

func SectionAddSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionAddSolidInnerSVG).
		Element(children...)
}

func SectionRemoveOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionRemoveOutlineInnerSVG).
		Element(children...)
}

func SectionRemoveSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionRemoveSolidInnerSVG).
		Element(children...)
}

func SendDownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendDownOutlineInnerSVG).
		Element(children...)
}

func SendDownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendDownSolidInnerSVG).
		Element(children...)
}

func SendLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendLeftOutlineInnerSVG).
		Element(children...)
}

func SendLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendLeftSolidInnerSVG).
		Element(children...)
}

func SendOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendOutlineInnerSVG).
		Element(children...)
}

func SendRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendRightOutlineInnerSVG).
		Element(children...)
}

func SendRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendRightSolidInnerSVG).
		Element(children...)
}

func SendSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendSolidInnerSVG).
		Element(children...)
}

func SendUpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendUpOutlineInnerSVG).
		Element(children...)
}

func SendUpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendUpSolidInnerSVG).
		Element(children...)
}

func ServersOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(serversOutlineInnerSVG).
		Element(children...)
}

func ServersSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(serversSolidInnerSVG).
		Element(children...)
}

func ShareOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareOutlineInnerSVG).
		Element(children...)
}

func ShareSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareSolidInnerSVG).
		Element(children...)
}

func ShieldOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldOutlineInnerSVG).
		Element(children...)
}

func ShieldSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldSolidInnerSVG).
		Element(children...)
}

func ShieldTickOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldTickOutlineInnerSVG).
		Element(children...)
}

func ShieldTickSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldTickSolidInnerSVG).
		Element(children...)
}

func ShieldXOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldXOutlineInnerSVG).
		Element(children...)
}

func ShieldXSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldXSolidInnerSVG).
		Element(children...)
}

func ShopOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shopOutlineInnerSVG).
		Element(children...)
}

func ShopSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shopSolidInnerSVG).
		Element(children...)
}

func SignOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signOutlineInnerSVG).
		Element(children...)
}

func SignSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signSolidInnerSVG).
		Element(children...)
}

func SigninOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signinOutlineInnerSVG).
		Element(children...)
}

func SigninSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signinSolidInnerSVG).
		Element(children...)
}

func SimOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(simOutlineInnerSVG).
		Element(children...)
}

func SimSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(simSolidInnerSVG).
		Element(children...)
}

func SimohamedOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(simohamedOutlineInnerSVG).
		Element(children...)
}

func SimohamedSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(simohamedSolidInnerSVG).
		Element(children...)
}

func SkullOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skullOutlineInnerSVG).
		Element(children...)
}

func SkullSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skullSolidInnerSVG).
		Element(children...)
}

func SkypeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skypeOutlineInnerSVG).
		Element(children...)
}

func SkypeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skypeSolidInnerSVG).
		Element(children...)
}

func SlackOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slackOutlineInnerSVG).
		Element(children...)
}

func SlackSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slackSolidInnerSVG).
		Element(children...)
}

func SnapchatOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snapchatOutlineInnerSVG).
		Element(children...)
}

func SnapchatSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snapchatSolidInnerSVG).
		Element(children...)
}

func SnesOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snesOutlineInnerSVG).
		Element(children...)
}

func SnesSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snesSolidInnerSVG).
		Element(children...)
}

func SortAlphabeticallyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortAlphabeticallyOutlineInnerSVG).
		Element(children...)
}

func SortAlphabeticallySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortAlphabeticallySolidInnerSVG).
		Element(children...)
}

func SortDownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortDownOutlineInnerSVG).
		Element(children...)
}

func SortDownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortDownSolidInnerSVG).
		Element(children...)
}

func SortHighToLowOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortHighToLowOutlineInnerSVG).
		Element(children...)
}

func SortHighToLowSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortHighToLowSolidInnerSVG).
		Element(children...)
}

func SortLowToHighOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortLowToHighOutlineInnerSVG).
		Element(children...)
}

func SortLowToHighSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortLowToHighSolidInnerSVG).
		Element(children...)
}

func SortReverseAlphabeticallyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortReverseAlphabeticallyOutlineInnerSVG).
		Element(children...)
}

func SortReverseAlphabeticallySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortReverseAlphabeticallySolidInnerSVG).
		Element(children...)
}

func SortUpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortUpOutlineInnerSVG).
		Element(children...)
}

func SortUpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortUpSolidInnerSVG).
		Element(children...)
}

func SoundOffOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundOffOutlineInnerSVG).
		Element(children...)
}

func SoundOffSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundOffSolidInnerSVG).
		Element(children...)
}

func SoundOnOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundOnOutlineInnerSVG).
		Element(children...)
}

func SoundOnSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundOnSolidInnerSVG).
		Element(children...)
}

func SpotifyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spotifyOutlineInnerSVG).
		Element(children...)
}

func SpotifySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spotifySolidInnerSVG).
		Element(children...)
}

func SpreadsheetOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spreadsheetOutlineInnerSVG).
		Element(children...)
}

func SpreadsheetSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spreadsheetSolidInnerSVG).
		Element(children...)
}

func SquareOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareOutlineInnerSVG).
		Element(children...)
}

func SquareSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareSolidInnerSVG).
		Element(children...)
}

func StackoverflowOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stackoverflowOutlineInnerSVG).
		Element(children...)
}

func StackoverflowSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stackoverflowSolidInnerSVG).
		Element(children...)
}

func StampOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stampOutlineInnerSVG).
		Element(children...)
}

func StampSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stampSolidInnerSVG).
		Element(children...)
}

func StarCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starCircleOutlineInnerSVG).
		Element(children...)
}

func StarCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starCircleSolidInnerSVG).
		Element(children...)
}

func StarOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starOutlineInnerSVG).
		Element(children...)
}

func StarSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starSmallOutlineInnerSVG).
		Element(children...)
}

func StarSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starSmallSolidInnerSVG).
		Element(children...)
}

func StarSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starSolidInnerSVG).
		Element(children...)
}

func StopCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopCircleOutlineInnerSVG).
		Element(children...)
}

func StopCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopCircleSolidInnerSVG).
		Element(children...)
}

func StopOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopOutlineInnerSVG).
		Element(children...)
}

func StopSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopSmallOutlineInnerSVG).
		Element(children...)
}

func StopSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopSmallSolidInnerSVG).
		Element(children...)
}

func StopSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopSolidInnerSVG).
		Element(children...)
}

func StopwatchOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopwatchOutlineInnerSVG).
		Element(children...)
}

func StopwatchSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopwatchSolidInnerSVG).
		Element(children...)
}

func StrikethroughOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(strikethroughOutlineInnerSVG).
		Element(children...)
}

func StrikethroughSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(strikethroughSolidInnerSVG).
		Element(children...)
}

func SubscriptOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subscriptOutlineInnerSVG).
		Element(children...)
}

func SubscriptSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subscriptSolidInnerSVG).
		Element(children...)
}

func SunOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunOutlineInnerSVG).
		Element(children...)
}

func SunSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunSolidInnerSVG).
		Element(children...)
}

func SuperscriptOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(superscriptOutlineInnerSVG).
		Element(children...)
}

func SuperscriptSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(superscriptSolidInnerSVG).
		Element(children...)
}

func SvelteOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(svelteOutlineInnerSVG).
		Element(children...)
}

func SvelteSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(svelteSolidInnerSVG).
		Element(children...)
}

func SvgOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(svgOutlineInnerSVG).
		Element(children...)
}

func SvgSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(svgSolidInnerSVG).
		Element(children...)
}

func TableOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableOutlineInnerSVG).
		Element(children...)
}

func TableSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableSolidInnerSVG).
		Element(children...)
}

func TabletOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tabletOutlineInnerSVG).
		Element(children...)
}

func TabletSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tabletSolidInnerSVG).
		Element(children...)
}

func TagOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagOutlineInnerSVG).
		Element(children...)
}

func TagSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagSolidInnerSVG).
		Element(children...)
}

func TailwindOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tailwindOutlineInnerSVG).
		Element(children...)
}

func TailwindSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tailwindSolidInnerSVG).
		Element(children...)
}

func TargetOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(targetOutlineInnerSVG).
		Element(children...)
}

func TargetSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(targetSolidInnerSVG).
		Element(children...)
}

func TelegramOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(telegramOutlineInnerSVG).
		Element(children...)
}

func TelegramSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(telegramSolidInnerSVG).
		Element(children...)
}

func TerminalOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalOutlineInnerSVG).
		Element(children...)
}

func TerminalSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalSolidInnerSVG).
		Element(children...)
}

func TextDocumentAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textDocumentAltOutlineInnerSVG).
		Element(children...)
}

func TextDocumentAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textDocumentAltSolidInnerSVG).
		Element(children...)
}

func TextDocumentOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textDocumentOutlineInnerSVG).
		Element(children...)
}

func TextDocumentSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textDocumentSolidInnerSVG).
		Element(children...)
}

func TextOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textOutlineInnerSVG).
		Element(children...)
}

func TextSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textSolidInnerSVG).
		Element(children...)
}

func ThreeHundredSixtyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeHundredSixtyOutlineInnerSVG).
		Element(children...)
}

func ThreeHundredSixtySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeHundredSixtySolidInnerSVG).
		Element(children...)
}

func ThumbDownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbDownOutlineInnerSVG).
		Element(children...)
}

func ThumbDownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbDownSolidInnerSVG).
		Element(children...)
}

func ThumbUpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbUpOutlineInnerSVG).
		Element(children...)
}

func ThumbUpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbUpSolidInnerSVG).
		Element(children...)
}

func ThumbtackOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbtackOutlineInnerSVG).
		Element(children...)
}

func ThumbtackSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbtackSolidInnerSVG).
		Element(children...)
}

func TickCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tickCircleOutlineInnerSVG).
		Element(children...)
}

func TickCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tickCircleSolidInnerSVG).
		Element(children...)
}

func TickOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tickOutlineInnerSVG).
		Element(children...)
}

func TickSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tickSmallOutlineInnerSVG).
		Element(children...)
}

func TickSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tickSmallSolidInnerSVG).
		Element(children...)
}

func TickSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tickSolidInnerSVG).
		Element(children...)
}

func TiktokOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tiktokOutlineInnerSVG).
		Element(children...)
}

func TiktokSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tiktokSolidInnerSVG).
		Element(children...)
}

func ToggleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleOutlineInnerSVG).
		Element(children...)
}

func ToggleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleSolidInnerSVG).
		Element(children...)
}

func TopLeftOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(topLeftOutlineInnerSVG).
		Element(children...)
}

func TopLeftSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(topLeftSolidInnerSVG).
		Element(children...)
}

func TopRightOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(topRightOutlineInnerSVG).
		Element(children...)
}

func TopRightSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(topRightSolidInnerSVG).
		Element(children...)
}

func TrendDownOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendDownOutlineInnerSVG).
		Element(children...)
}

func TrendDownSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendDownSolidInnerSVG).
		Element(children...)
}

func TrendUpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendUpOutlineInnerSVG).
		Element(children...)
}

func TrendUpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendUpSolidInnerSVG).
		Element(children...)
}

func TriangleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleOutlineInnerSVG).
		Element(children...)
}

func TriangleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleSolidInnerSVG).
		Element(children...)
}

func TrophyOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trophyOutlineInnerSVG).
		Element(children...)
}

func TrophySolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trophySolidInnerSVG).
		Element(children...)
}

func TvOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tvOutlineInnerSVG).
		Element(children...)
}

func TvSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tvSolidInnerSVG).
		Element(children...)
}

func TwitchOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitchOutlineInnerSVG).
		Element(children...)
}

func TwitchSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitchSolidInnerSVG).
		Element(children...)
}

func TwitterOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitterOutlineInnerSVG).
		Element(children...)
}

func TwitterSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitterSolidInnerSVG).
		Element(children...)
}

func TypescriptOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(typescriptOutlineInnerSVG).
		Element(children...)
}

func TypescriptSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(typescriptSolidInnerSVG).
		Element(children...)
}

func UnderlineOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(underlineOutlineInnerSVG).
		Element(children...)
}

func UnderlineSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(underlineSolidInnerSVG).
		Element(children...)
}

func UnlockCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockCircleOutlineInnerSVG).
		Element(children...)
}

func UnlockCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockCircleSolidInnerSVG).
		Element(children...)
}

func UnlockOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockOutlineInnerSVG).
		Element(children...)
}

func UnlockSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockSmallOutlineInnerSVG).
		Element(children...)
}

func UnlockSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockSmallSolidInnerSVG).
		Element(children...)
}

func UnlockSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockSolidInnerSVG).
		Element(children...)
}

func UpCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upCircleOutlineInnerSVG).
		Element(children...)
}

func UpCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upCircleSolidInnerSVG).
		Element(children...)
}

func UpOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upOutlineInnerSVG).
		Element(children...)
}

func UpSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upSmallOutlineInnerSVG).
		Element(children...)
}

func UpSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upSmallSolidInnerSVG).
		Element(children...)
}

func UpSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upSolidInnerSVG).
		Element(children...)
}

func UploadOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadOutlineInnerSVG).
		Element(children...)
}

func UploadSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadSolidInnerSVG).
		Element(children...)
}

func UsbCableOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usbCableOutlineInnerSVG).
		Element(children...)
}

func UsbCableSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usbCableSolidInnerSVG).
		Element(children...)
}

func UserCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCircleOutlineInnerSVG).
		Element(children...)
}

func UserCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCircleSolidInnerSVG).
		Element(children...)
}

func UserMinusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userMinusOutlineInnerSVG).
		Element(children...)
}

func UserMinusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userMinusSolidInnerSVG).
		Element(children...)
}

func UserOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userOutlineInnerSVG).
		Element(children...)
}

func UserPlusOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userPlusOutlineInnerSVG).
		Element(children...)
}

func UserPlusSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userPlusSolidInnerSVG).
		Element(children...)
}

func UserSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userSolidInnerSVG).
		Element(children...)
}

func UserSquareOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userSquareOutlineInnerSVG).
		Element(children...)
}

func UserSquareSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userSquareSolidInnerSVG).
		Element(children...)
}

func UsersOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usersOutlineInnerSVG).
		Element(children...)
}

func UsersSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usersSolidInnerSVG).
		Element(children...)
}

func VectorDocumentOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vectorDocumentOutlineInnerSVG).
		Element(children...)
}

func VectorDocumentSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vectorDocumentSolidInnerSVG).
		Element(children...)
}

func VennDiagramOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vennDiagramOutlineInnerSVG).
		Element(children...)
}

func VennDiagramSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vennDiagramSolidInnerSVG).
		Element(children...)
}

func ViewColumnOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewColumnOutlineInnerSVG).
		Element(children...)
}

func ViewColumnSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewColumnSolidInnerSVG).
		Element(children...)
}

func ViewGridOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewGridOutlineInnerSVG).
		Element(children...)
}

func ViewGridSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewGridSolidInnerSVG).
		Element(children...)
}

func VimOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vimOutlineInnerSVG).
		Element(children...)
}

func VimSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vimSolidInnerSVG).
		Element(children...)
}

func VolumeOneOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeOneOutlineInnerSVG).
		Element(children...)
}

func VolumeOneSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeOneSolidInnerSVG).
		Element(children...)
}

func VolumeThreeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeThreeOutlineInnerSVG).
		Element(children...)
}

func VolumeThreeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeThreeSolidInnerSVG).
		Element(children...)
}

func VolumeTwoOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeTwoOutlineInnerSVG).
		Element(children...)
}

func VolumeTwoSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeTwoSolidInnerSVG).
		Element(children...)
}

func VrHeadsetOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vrHeadsetOutlineInnerSVG).
		Element(children...)
}

func VrHeadsetSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vrHeadsetSolidInnerSVG).
		Element(children...)
}

func VueOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vueOutlineInnerSVG).
		Element(children...)
}

func VueSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vueSolidInnerSVG).
		Element(children...)
}

func WalletAltOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(walletAltOutlineInnerSVG).
		Element(children...)
}

func WalletAltSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(walletAltSolidInnerSVG).
		Element(children...)
}

func WalletOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(walletOutlineInnerSVG).
		Element(children...)
}

func WalletSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(walletSolidInnerSVG).
		Element(children...)
}

func WanOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wanOutlineInnerSVG).
		Element(children...)
}

func WanSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wanSolidInnerSVG).
		Element(children...)
}

func WandOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wandOutlineInnerSVG).
		Element(children...)
}

func WandSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wandSolidInnerSVG).
		Element(children...)
}

func WatchOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(watchOutlineInnerSVG).
		Element(children...)
}

func WatchSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(watchSolidInnerSVG).
		Element(children...)
}

func WebpackOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webpackOutlineInnerSVG).
		Element(children...)
}

func WebpackSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webpackSolidInnerSVG).
		Element(children...)
}

func WhatsappOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(whatsappOutlineInnerSVG).
		Element(children...)
}

func WhatsappSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(whatsappSolidInnerSVG).
		Element(children...)
}

func WifiFullOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiFullOutlineInnerSVG).
		Element(children...)
}

func WifiFullSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiFullSolidInnerSVG).
		Element(children...)
}

func WifiLowOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiLowOutlineInnerSVG).
		Element(children...)
}

func WifiLowSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiLowSolidInnerSVG).
		Element(children...)
}

func WifiNoneOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiNoneOutlineInnerSVG).
		Element(children...)
}

func WifiNoneSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiNoneSolidInnerSVG).
		Element(children...)
}

func WindowsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowsOutlineInnerSVG).
		Element(children...)
}

func WindowsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowsSolidInnerSVG).
		Element(children...)
}

func WordpressOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wordpressOutlineInnerSVG).
		Element(children...)
}

func WordpressSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wordpressSolidInnerSVG).
		Element(children...)
}

func XCircleOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xCircleOutlineInnerSVG).
		Element(children...)
}

func XCircleSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xCircleSolidInnerSVG).
		Element(children...)
}

func XOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xOutlineInnerSVG).
		Element(children...)
}

func XSmallOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xSmallOutlineInnerSVG).
		Element(children...)
}

func XSmallSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xSmallSolidInnerSVG).
		Element(children...)
}

func XSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xSolidInnerSVG).
		Element(children...)
}

func XlsOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xlsOutlineInnerSVG).
		Element(children...)
}

func XlsSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xlsSolidInnerSVG).
		Element(children...)
}

func YenOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yenOutlineInnerSVG).
		Element(children...)
}

func YenSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yenSolidInnerSVG).
		Element(children...)
}

func YoutubeOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(youtubeOutlineInnerSVG).
		Element(children...)
}

func YoutubeSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(youtubeSolidInnerSVG).
		Element(children...)
}

func ZipOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zipOutlineInnerSVG).
		Element(children...)
}

func ZipSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zipSolidInnerSVG).
		Element(children...)
}

func ZoomInOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomInOutlineInnerSVG).
		Element(children...)
}

func ZoomInSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomInSolidInnerSVG).
		Element(children...)
}

func ZoomOutOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomOutOutlineInnerSVG).
		Element(children...)
}

func ZoomOutSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomOutSolidInnerSVG).
		Element(children...)
}

func ByName(name string) (*engine.UberElement, error) {
	switch name {
	case "ab-testing-outline":
		return AbTestingOutline(), nil
	case "ab-testing-solid":
		return AbTestingSolid(), nil
	case "add-outline":
		return AddOutline(), nil
	case "add-small-outline":
		return AddSmallOutline(), nil
	case "add-small-solid":
		return AddSmallSolid(), nil
	case "add-solid":
		return AddSolid(), nil
	case "address-book-outline":
		return AddressBookOutline(), nil
	case "address-book-solid":
		return AddressBookSolid(), nil
	case "adjust-horizontal-alt-outline":
		return AdjustHorizontalAltOutline(), nil
	case "adjust-horizontal-alt-solid":
		return AdjustHorizontalAltSolid(), nil
	case "adjust-horizontal-outline":
		return AdjustHorizontalOutline(), nil
	case "adjust-horizontal-solid":
		return AdjustHorizontalSolid(), nil
	case "adjust-vertical-alt-outline":
		return AdjustVerticalAltOutline(), nil
	case "adjust-vertical-alt-solid":
		return AdjustVerticalAltSolid(), nil
	case "adjust-vertical-outline":
		return AdjustVerticalOutline(), nil
	case "adjust-vertical-solid":
		return AdjustVerticalSolid(), nil
	case "airplay-outline":
		return AirplayOutline(), nil
	case "airplay-solid":
		return AirplaySolid(), nil
	case "airpods-outline":
		return AirpodsOutline(), nil
	case "airpods-solid":
		return AirpodsSolid(), nil
	case "alarm-outline":
		return AlarmOutline(), nil
	case "alarm-solid":
		return AlarmSolid(), nil
	case "alien-outline":
		return AlienOutline(), nil
	case "alien-solid":
		return AlienSolid(), nil
	case "align-bottom-outline":
		return AlignBottomOutline(), nil
	case "align-bottom-solid":
		return AlignBottomSolid(), nil
	case "align-center-horizontal-outline":
		return AlignCenterHorizontalOutline(), nil
	case "align-center-horizontal-solid":
		return AlignCenterHorizontalSolid(), nil
	case "align-center-vertical-outline":
		return AlignCenterVerticalOutline(), nil
	case "align-center-vertical-solid":
		return AlignCenterVerticalSolid(), nil
	case "align-left-outline":
		return AlignLeftOutline(), nil
	case "align-left-solid":
		return AlignLeftSolid(), nil
	case "align-right-outline":
		return AlignRightOutline(), nil
	case "align-right-solid":
		return AlignRightSolid(), nil
	case "align-text-center-outline":
		return AlignTextCenterOutline(), nil
	case "align-text-center-solid":
		return AlignTextCenterSolid(), nil
	case "align-text-justify-outline":
		return AlignTextJustifyOutline(), nil
	case "align-text-justify-solid":
		return AlignTextJustifySolid(), nil
	case "align-text-left-outline":
		return AlignTextLeftOutline(), nil
	case "align-text-left-solid":
		return AlignTextLeftSolid(), nil
	case "align-text-right-outline":
		return AlignTextRightOutline(), nil
	case "align-text-right-solid":
		return AlignTextRightSolid(), nil
	case "align-top-outline":
		return AlignTopOutline(), nil
	case "align-top-solid":
		return AlignTopSolid(), nil
	case "anchor-outline":
		return AnchorOutline(), nil
	case "anchor-solid":
		return AnchorSolid(), nil
	case "android-outline":
		return AndroidOutline(), nil
	case "android-solid":
		return AndroidSolid(), nil
	case "angular-outline":
		return AngularOutline(), nil
	case "angular-solid":
		return AngularSolid(), nil
	case "anja-outline":
		return AnjaOutline(), nil
	case "anja-solid":
		return AnjaSolid(), nil
	case "anti-clockwise-outline":
		return AntiClockwiseOutline(), nil
	case "anti-clockwise-solid":
		return AntiClockwiseSolid(), nil
	case "apple-outline":
		return AppleOutline(), nil
	case "apple-solid":
		return AppleSolid(), nil
	case "appointments-outline":
		return AppointmentsOutline(), nil
	case "appointments-solid":
		return AppointmentsSolid(), nil
	case "archive-outline":
		return ArchiveOutline(), nil
	case "archive-solid":
		return ArchiveSolid(), nil
	case "area-chart-alt-outline":
		return AreaChartAltOutline(), nil
	case "area-chart-alt-solid":
		return AreaChartAltSolid(), nil
	case "area-chart-outline":
		return AreaChartOutline(), nil
	case "area-chart-solid":
		return AreaChartSolid(), nil
	case "arrow-down-circle-outline":
		return ArrowDownCircleOutline(), nil
	case "arrow-down-circle-solid":
		return ArrowDownCircleSolid(), nil
	case "arrow-down-outline":
		return ArrowDownOutline(), nil
	case "arrow-down-small-outline":
		return ArrowDownSmallOutline(), nil
	case "arrow-down-small-solid":
		return ArrowDownSmallSolid(), nil
	case "arrow-down-solid":
		return ArrowDownSolid(), nil
	case "arrow-left-circle-outline":
		return ArrowLeftCircleOutline(), nil
	case "arrow-left-circle-solid":
		return ArrowLeftCircleSolid(), nil
	case "arrow-left-outline":
		return ArrowLeftOutline(), nil
	case "arrow-left-small-outline":
		return ArrowLeftSmallOutline(), nil
	case "arrow-left-small-solid":
		return ArrowLeftSmallSolid(), nil
	case "arrow-left-solid":
		return ArrowLeftSolid(), nil
	case "arrow-outline":
		return ArrowOutline(), nil
	case "arrow-right-circle-outline":
		return ArrowRightCircleOutline(), nil
	case "arrow-right-circle-solid":
		return ArrowRightCircleSolid(), nil
	case "arrow-right-outline":
		return ArrowRightOutline(), nil
	case "arrow-right-small-outline":
		return ArrowRightSmallOutline(), nil
	case "arrow-right-small-solid":
		return ArrowRightSmallSolid(), nil
	case "arrow-right-solid":
		return ArrowRightSolid(), nil
	case "arrow-solid":
		return ArrowSolid(), nil
	case "arrow-up-circle-outline":
		return ArrowUpCircleOutline(), nil
	case "arrow-up-circle-solid":
		return ArrowUpCircleSolid(), nil
	case "arrow-up-outline":
		return ArrowUpOutline(), nil
	case "arrow-up-small-outline":
		return ArrowUpSmallOutline(), nil
	case "arrow-up-small-solid":
		return ArrowUpSmallSolid(), nil
	case "arrow-up-solid":
		return ArrowUpSolid(), nil
	case "artboard-outline":
		return ArtboardOutline(), nil
	case "artboard-solid":
		return ArtboardSolid(), nil
	case "at-outline":
		return AtOutline(), nil
	case "at-solid":
		return AtSolid(), nil
	case "attach-outline":
		return AttachOutline(), nil
	case "attach-solid":
		return AttachSolid(), nil
	case "attachment-outline":
		return AttachmentOutline(), nil
	case "attachment-solid":
		return AttachmentSolid(), nil
	case "audio-cable-outline":
		return AudioCableOutline(), nil
	case "audio-cable-solid":
		return AudioCableSolid(), nil
	case "audio-document-outline":
		return AudioDocumentOutline(), nil
	case "audio-document-solid":
		return AudioDocumentSolid(), nil
	case "azure-outline":
		return AzureOutline(), nil
	case "azure-solid":
		return AzureSolid(), nil
	case "backspace-outline":
		return BackspaceOutline(), nil
	case "backspace-solid":
		return BackspaceSolid(), nil
	case "bag-alt-outline":
		return BagAltOutline(), nil
	case "bag-alt-solid":
		return BagAltSolid(), nil
	case "bag-minus-outline":
		return BagMinusOutline(), nil
	case "bag-minus-solid":
		return BagMinusSolid(), nil
	case "bag-outline":
		return BagOutline(), nil
	case "bag-plus-outline":
		return BagPlusOutline(), nil
	case "bag-plus-solid":
		return BagPlusSolid(), nil
	case "bag-solid":
		return BagSolid(), nil
	case "bank-outline":
		return BankOutline(), nil
	case "bank-solid":
		return BankSolid(), nil
	case "bar-chart-outline":
		return BarChartOutline(), nil
	case "bar-chart-solid":
		return BarChartSolid(), nil
	case "barcode-outline":
		return BarcodeOutline(), nil
	case "barcode-solid":
		return BarcodeSolid(), nil
	case "basket-minus-outline":
		return BasketMinusOutline(), nil
	case "basket-minus-solid":
		return BasketMinusSolid(), nil
	case "basket-outline":
		return BasketOutline(), nil
	case "basket-plus-outline":
		return BasketPlusOutline(), nil
	case "basket-plus-solid":
		return BasketPlusSolid(), nil
	case "basket-solid":
		return BasketSolid(), nil
	case "bath-outline":
		return BathOutline(), nil
	case "bath-solid":
		return BathSolid(), nil
	case "battery-charge-outline":
		return BatteryChargeOutline(), nil
	case "battery-charge-solid":
		return BatteryChargeSolid(), nil
	case "battery-5-outline":
		return BatteryFiveOutline(), nil
	case "battery-5-solid":
		return BatteryFiveSolid(), nil
	case "battery-4-outline":
		return BatteryFourOutline(), nil
	case "battery-4-solid":
		return BatteryFourSolid(), nil
	case "battery-1-outline":
		return BatteryOneOutline(), nil
	case "battery-1-solid":
		return BatteryOneSolid(), nil
	case "battery-3-outline":
		return BatteryThreeOutline(), nil
	case "battery-3-solid":
		return BatteryThreeSolid(), nil
	case "battery-2-outline":
		return BatteryTwoOutline(), nil
	case "battery-2-solid":
		return BatteryTwoSolid(), nil
	case "battery-0-outline":
		return BatteryZeroOutline(), nil
	case "battery-0-solid":
		return BatteryZeroSolid(), nil
	case "bed-double-outline":
		return BedDoubleOutline(), nil
	case "bed-double-solid":
		return BedDoubleSolid(), nil
	case "bed-single-outline":
		return BedSingleOutline(), nil
	case "bed-single-solid":
		return BedSingleSolid(), nil
	case "behance-outline":
		return BehanceOutline(), nil
	case "behance-solid":
		return BehanceSolid(), nil
	case "bell-outline":
		return BellOutline(), nil
	case "bell-solid":
		return BellSolid(), nil
	case "bin-outline":
		return BinOutline(), nil
	case "bin-solid":
		return BinSolid(), nil
	case "bitbucket-outline":
		return BitbucketOutline(), nil
	case "bitbucket-solid":
		return BitbucketSolid(), nil
	case "bitcoin-outline":
		return BitcoinOutline(), nil
	case "bitcoin-solid":
		return BitcoinSolid(), nil
	case "bluetooth-outline":
		return BluetoothOutline(), nil
	case "bluetooth-solid":
		return BluetoothSolid(), nil
	case "bold-outline":
		return BoldOutline(), nil
	case "bold-solid":
		return BoldSolid(), nil
	case "book-outline":
		return BookOutline(), nil
	case "book-solid":
		return BookSolid(), nil
	case "bookmark-outline":
		return BookmarkOutline(), nil
	case "bookmark-solid":
		return BookmarkSolid(), nil
	case "border-all-outline":
		return BorderAllOutline(), nil
	case "border-all-solid":
		return BorderAllSolid(), nil
	case "border-bottom-outline":
		return BorderBottomOutline(), nil
	case "border-bottom-solid":
		return BorderBottomSolid(), nil
	case "border-horizontal-outline":
		return BorderHorizontalOutline(), nil
	case "border-horizontal-solid":
		return BorderHorizontalSolid(), nil
	case "border-inner-outline":
		return BorderInnerOutline(), nil
	case "border-inner-solid":
		return BorderInnerSolid(), nil
	case "border-left-outline":
		return BorderLeftOutline(), nil
	case "border-left-solid":
		return BorderLeftSolid(), nil
	case "border-none-outline":
		return BorderNoneOutline(), nil
	case "border-none-solid":
		return BorderNoneSolid(), nil
	case "border-outer-outline":
		return BorderOuterOutline(), nil
	case "border-outer-solid":
		return BorderOuterSolid(), nil
	case "border-radius-outline":
		return BorderRadiusOutline(), nil
	case "border-radius-solid":
		return BorderRadiusSolid(), nil
	case "border-right-outline":
		return BorderRightOutline(), nil
	case "border-right-solid":
		return BorderRightSolid(), nil
	case "border-top-outline":
		return BorderTopOutline(), nil
	case "border-top-solid":
		return BorderTopSolid(), nil
	case "border-vertical-outline":
		return BorderVerticalOutline(), nil
	case "border-vertical-solid":
		return BorderVerticalSolid(), nil
	case "bottom-left-outline":
		return BottomLeftOutline(), nil
	case "bottom-left-solid":
		return BottomLeftSolid(), nil
	case "bottom-right-outline":
		return BottomRightOutline(), nil
	case "bottom-right-solid":
		return BottomRightSolid(), nil
	case "box-outline":
		return BoxOutline(), nil
	case "box-solid":
		return BoxSolid(), nil
	case "bracket-outline":
		return BracketOutline(), nil
	case "bracket-solid":
		return BracketSolid(), nil
	case "briefcase-alt-outline":
		return BriefcaseAltOutline(), nil
	case "briefcase-alt-solid":
		return BriefcaseAltSolid(), nil
	case "briefcase-outline":
		return BriefcaseOutline(), nil
	case "briefcase-solid":
		return BriefcaseSolid(), nil
	case "brush-outline":
		return BrushOutline(), nil
	case "brush-solid":
		return BrushSolid(), nil
	case "bug-outline":
		return BugOutline(), nil
	case "bug-solid":
		return BugSolid(), nil
	case "building-outline":
		return BuildingOutline(), nil
	case "building-solid":
		return BuildingSolid(), nil
	case "bulb-off-outline":
		return BulbOffOutline(), nil
	case "bulb-off-solid":
		return BulbOffSolid(), nil
	case "bulb-on-outline":
		return BulbOnOutline(), nil
	case "bulb-on-solid":
		return BulbOnSolid(), nil
	case "button-outline":
		return ButtonOutline(), nil
	case "button-solid":
		return ButtonSolid(), nil
	case "c-outline":
		return COutline(), nil
	case "c-sharp-outline":
		return CSharpOutline(), nil
	case "c-sharp-solid":
		return CSharpSolid(), nil
	case "c-solid":
		return CSolid(), nil
	case "calculator-outline":
		return CalculatorOutline(), nil
	case "calculator-solid":
		return CalculatorSolid(), nil
	case "calendar-minus-outline":
		return CalendarMinusOutline(), nil
	case "calendar-minus-solid":
		return CalendarMinusSolid(), nil
	case "calendar-no-access-outline":
		return CalendarNoAccessOutline(), nil
	case "calendar-no-access-solid":
		return CalendarNoAccessSolid(), nil
	case "calendar-outline":
		return CalendarOutline(), nil
	case "calendar-plus-outline":
		return CalendarPlusOutline(), nil
	case "calendar-plus-solid":
		return CalendarPlusSolid(), nil
	case "calendar-solid":
		return CalendarSolid(), nil
	case "calendar-tick-outline":
		return CalendarTickOutline(), nil
	case "calendar-tick-solid":
		return CalendarTickSolid(), nil
	case "calendar-x-outline":
		return CalendarXOutline(), nil
	case "calendar-x-solid":
		return CalendarXSolid(), nil
	case "camera-outline":
		return CameraOutline(), nil
	case "camera-solid":
		return CameraSolid(), nil
	case "candle-chart-outline":
		return CandleChartOutline(), nil
	case "candle-chart-solid":
		return CandleChartSolid(), nil
	case "car-outline":
		return CarOutline(), nil
	case "car-solid":
		return CarSolid(), nil
	case "caret-vertical-circle-outline":
		return CaretVerticalCircleOutline(), nil
	case "caret-vertical-circle-solid":
		return CaretVerticalCircleSolid(), nil
	case "caret-vertical-outline":
		return CaretVerticalOutline(), nil
	case "caret-vertical-small-outline":
		return CaretVerticalSmallOutline(), nil
	case "caret-vertical-small-solid":
		return CaretVerticalSmallSolid(), nil
	case "caret-vertical-solid":
		return CaretVerticalSolid(), nil
	case "cart-minus-outline":
		return CartMinusOutline(), nil
	case "cart-minus-solid":
		return CartMinusSolid(), nil
	case "cart-outline":
		return CartOutline(), nil
	case "cart-plus-outline":
		return CartPlusOutline(), nil
	case "cart-plus-solid":
		return CartPlusSolid(), nil
	case "cart-solid":
		return CartSolid(), nil
	case "certificate-outline":
		return CertificateOutline(), nil
	case "certificate-solid":
		return CertificateSolid(), nil
	case "chat-outline":
		return ChatOutline(), nil
	case "chat-solid":
		return ChatSolid(), nil
	case "chat-typing-alt-outline":
		return ChatTypingAltOutline(), nil
	case "chat-typing-alt-solid":
		return ChatTypingAltSolid(), nil
	case "chat-typing-outline":
		return ChatTypingOutline(), nil
	case "chat-typing-solid":
		return ChatTypingSolid(), nil
	case "chatbot-outline":
		return ChatbotOutline(), nil
	case "chatbot-solid":
		return ChatbotSolid(), nil
	case "chrome-outline":
		return ChromeOutline(), nil
	case "chrome-solid":
		return ChromeSolid(), nil
	case "church-outline":
		return ChurchOutline(), nil
	case "church-solid":
		return ChurchSolid(), nil
	case "circle-outline":
		return CircleOutline(), nil
	case "circle-solid":
		return CircleSolid(), nil
	case "clipboard-minus-outline":
		return ClipboardMinusOutline(), nil
	case "clipboard-minus-solid":
		return ClipboardMinusSolid(), nil
	case "clipboard-no-access-outline":
		return ClipboardNoAccessOutline(), nil
	case "clipboard-no-access-solid":
		return ClipboardNoAccessSolid(), nil
	case "clipboard-outline":
		return ClipboardOutline(), nil
	case "clipboard-plus-outline":
		return ClipboardPlusOutline(), nil
	case "clipboard-plus-solid":
		return ClipboardPlusSolid(), nil
	case "clipboard-solid":
		return ClipboardSolid(), nil
	case "clipboard-tick-outline":
		return ClipboardTickOutline(), nil
	case "clipboard-tick-solid":
		return ClipboardTickSolid(), nil
	case "clipboard-x-outline":
		return ClipboardXOutline(), nil
	case "clipboard-x-solid":
		return ClipboardXSolid(), nil
	case "clock-outline":
		return ClockOutline(), nil
	case "clock-solid":
		return ClockSolid(), nil
	case "clockwise-outline":
		return ClockwiseOutline(), nil
	case "clockwise-solid":
		return ClockwiseSolid(), nil
	case "code-outline":
		return CodeOutline(), nil
	case "code-solid":
		return CodeSolid(), nil
	case "codepen-outline":
		return CodepenOutline(), nil
	case "codepen-solid":
		return CodepenSolid(), nil
	case "cog-outline":
		return CogOutline(), nil
	case "cog-solid":
		return CogSolid(), nil
	case "compass-outline":
		return CompassOutline(), nil
	case "compass-solid":
		return CompassSolid(), nil
	case "computer-outline":
		return ComputerOutline(), nil
	case "computer-solid":
		return ComputerSolid(), nil
	case "contact-outline":
		return ContactOutline(), nil
	case "contact-solid":
		return ContactSolid(), nil
	case "contract-outline":
		return ContractOutline(), nil
	case "contract-solid":
		return ContractSolid(), nil
	case "cost-estimate-outline":
		return CostEstimateOutline(), nil
	case "cost-estimate-solid":
		return CostEstimateSolid(), nil
	case "cplusplus-outline":
		return CplusplusOutline(), nil
	case "cplusplus-solid":
		return CplusplusSolid(), nil
	case "credit-card-outline":
		return CreditCardOutline(), nil
	case "credit-card-solid":
		return CreditCardSolid(), nil
	case "crop-outline":
		return CropOutline(), nil
	case "crop-solid":
		return CropSolid(), nil
	case "css3-outline":
		return CssThreeOutline(), nil
	case "css3-solid":
		return CssThreeSolid(), nil
	case "csv-outline":
		return CsvOutline(), nil
	case "csv-solid":
		return CsvSolid(), nil
	case "cup-outline":
		return CupOutline(), nil
	case "cup-solid":
		return CupSolid(), nil
	case "curved-connector-outline":
		return CurvedConnectorOutline(), nil
	case "curved-connector-solid":
		return CurvedConnectorSolid(), nil
	case "d3-outline":
		return DThreeOutline(), nil
	case "d3-solid":
		return DThreeSolid(), nil
	case "database-outline":
		return DatabaseOutline(), nil
	case "database-solid":
		return DatabaseSolid(), nil
	case "denied-outline":
		return DeniedOutline(), nil
	case "denied-solid":
		return DeniedSolid(), nil
	case "deno-outline":
		return DenoOutline(), nil
	case "deno-solid":
		return DenoSolid(), nil
	case "depth-chart-outline":
		return DepthChartOutline(), nil
	case "depth-chart-solid":
		return DepthChartSolid(), nil
	case "desklamp-outline":
		return DesklampOutline(), nil
	case "desklamp-solid":
		return DesklampSolid(), nil
	case "diamond-outline":
		return DiamondOutline(), nil
	case "diamond-solid":
		return DiamondSolid(), nil
	case "direction-outline":
		return DirectionOutline(), nil
	case "direction-solid":
		return DirectionSolid(), nil
	case "discord-outline":
		return DiscordOutline(), nil
	case "discord-solid":
		return DiscordSolid(), nil
	case "discount-outline":
		return DiscountOutline(), nil
	case "discount-solid":
		return DiscountSolid(), nil
	case "distribute-horizontal-outline":
		return DistributeHorizontalOutline(), nil
	case "distribute-horizontal-solid":
		return DistributeHorizontalSolid(), nil
	case "distribute-vertical-outline":
		return DistributeVerticalOutline(), nil
	case "distribute-vertical-solid":
		return DistributeVerticalSolid(), nil
	case "divider-line-outline":
		return DividerLineOutline(), nil
	case "divider-line-solid":
		return DividerLineSolid(), nil
	case "doc-outline":
		return DocOutline(), nil
	case "doc-solid":
		return DocSolid(), nil
	case "docker-outline":
		return DockerOutline(), nil
	case "docker-solid":
		return DockerSolid(), nil
	case "documents-outline":
		return DocumentsOutline(), nil
	case "documents-solid":
		return DocumentsSolid(), nil
	case "dollar-outline":
		return DollarOutline(), nil
	case "dollar-solid":
		return DollarSolid(), nil
	case "donut-chart-outline":
		return DonutChartOutline(), nil
	case "donut-chart-solid":
		return DonutChartSolid(), nil
	case "double-caret-down-circle-outline":
		return DoubleCaretDownCircleOutline(), nil
	case "double-caret-down-circle-solid":
		return DoubleCaretDownCircleSolid(), nil
	case "double-caret-down-outline":
		return DoubleCaretDownOutline(), nil
	case "double-caret-down-small-outline":
		return DoubleCaretDownSmallOutline(), nil
	case "double-caret-down-small-solid":
		return DoubleCaretDownSmallSolid(), nil
	case "double-caret-down-solid":
		return DoubleCaretDownSolid(), nil
	case "double-caret-left-circle-outline":
		return DoubleCaretLeftCircleOutline(), nil
	case "double-caret-left-circle-solid":
		return DoubleCaretLeftCircleSolid(), nil
	case "double-caret-left-outline":
		return DoubleCaretLeftOutline(), nil
	case "double-caret-left-small-outline":
		return DoubleCaretLeftSmallOutline(), nil
	case "double-caret-left-small-solid":
		return DoubleCaretLeftSmallSolid(), nil
	case "double-caret-left-solid":
		return DoubleCaretLeftSolid(), nil
	case "double-caret-right-circle-outline":
		return DoubleCaretRightCircleOutline(), nil
	case "double-caret-right-circle-solid":
		return DoubleCaretRightCircleSolid(), nil
	case "double-caret-right-outline":
		return DoubleCaretRightOutline(), nil
	case "double-caret-right-small-outline":
		return DoubleCaretRightSmallOutline(), nil
	case "double-caret-right-small-solid":
		return DoubleCaretRightSmallSolid(), nil
	case "double-caret-right-solid":
		return DoubleCaretRightSolid(), nil
	case "double-caret-up-circle-outline":
		return DoubleCaretUpCircleOutline(), nil
	case "double-caret-up-circle-solid":
		return DoubleCaretUpCircleSolid(), nil
	case "double-caret-up-outline":
		return DoubleCaretUpOutline(), nil
	case "double-caret-up-small-outline":
		return DoubleCaretUpSmallOutline(), nil
	case "double-caret-up-small-solid":
		return DoubleCaretUpSmallSolid(), nil
	case "double-caret-up-solid":
		return DoubleCaretUpSolid(), nil
	case "down-circle-outline":
		return DownCircleOutline(), nil
	case "down-circle-solid":
		return DownCircleSolid(), nil
	case "down-outline":
		return DownOutline(), nil
	case "down-small-outline":
		return DownSmallOutline(), nil
	case "down-small-solid":
		return DownSmallSolid(), nil
	case "down-solid":
		return DownSolid(), nil
	case "download-outline":
		return DownloadOutline(), nil
	case "download-solid":
		return DownloadSolid(), nil
	case "drag-horizontal-outline":
		return DragHorizontalOutline(), nil
	case "drag-horizontal-solid":
		return DragHorizontalSolid(), nil
	case "drag-outline":
		return DragOutline(), nil
	case "drag-solid":
		return DragSolid(), nil
	case "drag-vertical-outline":
		return DragVerticalOutline(), nil
	case "drag-vertical-solid":
		return DragVerticalSolid(), nil
	case "dribbble-outline":
		return DribbbleOutline(), nil
	case "dribbble-solid":
		return DribbbleSolid(), nil
	case "drop-outline":
		return DropOutline(), nil
	case "drop-solid":
		return DropSolid(), nil
	case "dropper-outline":
		return DropperOutline(), nil
	case "dropper-solid":
		return DropperSolid(), nil
	case "edge-outline":
		return EdgeOutline(), nil
	case "edge-solid":
		return EdgeSolid(), nil
	case "edit-circle-outline":
		return EditCircleOutline(), nil
	case "edit-circle-solid":
		return EditCircleSolid(), nil
	case "edit-1-outline":
		return EditOneOutline(), nil
	case "edit-1-solid":
		return EditOneSolid(), nil
	case "edit-outline":
		return EditOutline(), nil
	case "edit-small-outline":
		return EditSmallOutline(), nil
	case "edit-small-solid":
		return EditSmallSolid(), nil
	case "edit-solid":
		return EditSolid(), nil
	case "elbow-connector-outline":
		return ElbowConnectorOutline(), nil
	case "elbow-connector-solid":
		return ElbowConnectorSolid(), nil
	case "envelope-open-outline":
		return EnvelopeOpenOutline(), nil
	case "envelope-open-solid":
		return EnvelopeOpenSolid(), nil
	case "envelope-outline":
		return EnvelopeOutline(), nil
	case "envelope-solid":
		return EnvelopeSolid(), nil
	case "eps-outline":
		return EpsOutline(), nil
	case "eps-solid":
		return EpsSolid(), nil
	case "eslint-outline":
		return EslintOutline(), nil
	case "eslint-solid":
		return EslintSolid(), nil
	case "ethereum-outline":
		return EthereumOutline(), nil
	case "ethereum-solid":
		return EthereumSolid(), nil
	case "euro-outline":
		return EuroOutline(), nil
	case "euro-solid":
		return EuroSolid(), nil
	case "exclamation-circle-outline":
		return ExclamationCircleOutline(), nil
	case "exclamation-circle-solid":
		return ExclamationCircleSolid(), nil
	case "exclamation-outline":
		return ExclamationOutline(), nil
	case "exclamation-small-outline":
		return ExclamationSmallOutline(), nil
	case "exclamation-small-solid":
		return ExclamationSmallSolid(), nil
	case "exclamation-solid":
		return ExclamationSolid(), nil
	case "expand-alt-outline":
		return ExpandAltOutline(), nil
	case "expand-alt-solid":
		return ExpandAltSolid(), nil
	case "expand-outline":
		return ExpandOutline(), nil
	case "expand-solid":
		return ExpandSolid(), nil
	case "eye-closed-outline":
		return EyeClosedOutline(), nil
	case "eye-closed-solid":
		return EyeClosedSolid(), nil
	case "eye-outline":
		return EyeOutline(), nil
	case "eye-solid":
		return EyeSolid(), nil
	case "face-id-outline":
		return FaceIdOutline(), nil
	case "face-id-solid":
		return FaceIdSolid(), nil
	case "facebook-outline":
		return FacebookOutline(), nil
	case "facebook-solid":
		return FacebookSolid(), nil
	case "figma-outline":
		return FigmaOutline(), nil
	case "figma-solid":
		return FigmaSolid(), nil
	case "file-minus-outline":
		return FileMinusOutline(), nil
	case "file-minus-solid":
		return FileMinusSolid(), nil
	case "file-no-access-outline":
		return FileNoAccessOutline(), nil
	case "file-no-access-solid":
		return FileNoAccessSolid(), nil
	case "file-outline":
		return FileOutline(), nil
	case "file-plus-outline":
		return FilePlusOutline(), nil
	case "file-plus-solid":
		return FilePlusSolid(), nil
	case "file-solid":
		return FileSolid(), nil
	case "file-tick-outline":
		return FileTickOutline(), nil
	case "file-tick-solid":
		return FileTickSolid(), nil
	case "file-x-outline":
		return FileXOutline(), nil
	case "file-x-solid":
		return FileXSolid(), nil
	case "filter-outline":
		return FilterOutline(), nil
	case "filter-solid":
		return FilterSolid(), nil
	case "fingerprint-outline":
		return FingerprintOutline(), nil
	case "fingerprint-solid":
		return FingerprintSolid(), nil
	case "firebase-outline":
		return FirebaseOutline(), nil
	case "firebase-solid":
		return FirebaseSolid(), nil
	case "flag-alt-outline":
		return FlagAltOutline(), nil
	case "flag-alt-solid":
		return FlagAltSolid(), nil
	case "flag-outline":
		return FlagOutline(), nil
	case "flag-solid":
		return FlagSolid(), nil
	case "flip-horizontal-outline":
		return FlipHorizontalOutline(), nil
	case "flip-horizontal-solid":
		return FlipHorizontalSolid(), nil
	case "flip-vertical-outline":
		return FlipVerticalOutline(), nil
	case "flip-vertical-solid":
		return FlipVerticalSolid(), nil
	case "float-center-outline":
		return FloatCenterOutline(), nil
	case "float-center-solid":
		return FloatCenterSolid(), nil
	case "float-left-outline":
		return FloatLeftOutline(), nil
	case "float-left-solid":
		return FloatLeftSolid(), nil
	case "float-right-outline":
		return FloatRightOutline(), nil
	case "float-right-solid":
		return FloatRightSolid(), nil
	case "floorplan-outline":
		return FloorplanOutline(), nil
	case "floorplan-solid":
		return FloorplanSolid(), nil
	case "folder-minus-outline":
		return FolderMinusOutline(), nil
	case "folder-minus-solid":
		return FolderMinusSolid(), nil
	case "folder-no-access-outline":
		return FolderNoAccessOutline(), nil
	case "folder-no-access-solid":
		return FolderNoAccessSolid(), nil
	case "folder-outline":
		return FolderOutline(), nil
	case "folder-plus-outline":
		return FolderPlusOutline(), nil
	case "folder-plus-solid":
		return FolderPlusSolid(), nil
	case "folder-solid":
		return FolderSolid(), nil
	case "folder-tick-outline":
		return FolderTickOutline(), nil
	case "folder-tick-solid":
		return FolderTickSolid(), nil
	case "folder-x-outline":
		return FolderXOutline(), nil
	case "folder-x-solid":
		return FolderXSolid(), nil
	case "folders-outline":
		return FoldersOutline(), nil
	case "folders-solid":
		return FoldersSolid(), nil
	case "forward-circle-outline":
		return ForwardCircleOutline(), nil
	case "forward-circle-solid":
		return ForwardCircleSolid(), nil
	case "forward-outline":
		return ForwardOutline(), nil
	case "forward-small-outline":
		return ForwardSmallOutline(), nil
	case "forward-small-solid":
		return ForwardSmallSolid(), nil
	case "forward-solid":
		return ForwardSolid(), nil
	case "frame-outline":
		return FrameOutline(), nil
	case "frame-solid":
		return FrameSolid(), nil
	case "framer-outline":
		return FramerOutline(), nil
	case "framer-solid":
		return FramerSolid(), nil
	case "game-controller-outline":
		return GameControllerOutline(), nil
	case "game-controller-retro-outline":
		return GameControllerRetroOutline(), nil
	case "game-controller-retro-solid":
		return GameControllerRetroSolid(), nil
	case "game-controller-solid":
		return GameControllerSolid(), nil
	case "gantt-chart-outline":
		return GanttChartOutline(), nil
	case "gantt-chart-solid":
		return GanttChartSolid(), nil
	case "garage-outline":
		return GarageOutline(), nil
	case "garage-solid":
		return GarageSolid(), nil
	case "gatsbyjs-outline":
		return GatsbyjsOutline(), nil
	case "gatsbyjs-solid":
		return GatsbyjsSolid(), nil
	case "gba-outline":
		return GbaOutline(), nil
	case "gba-solid":
		return GbaSolid(), nil
	case "gbc-outline":
		return GbcOutline(), nil
	case "gbc-solid":
		return GbcSolid(), nil
	case "ghost-outline":
		return GhostOutline(), nil
	case "ghost-solid":
		return GhostSolid(), nil
	case "gif-outline":
		return GifOutline(), nil
	case "gif-solid":
		return GifSolid(), nil
	case "gift-outline":
		return GiftOutline(), nil
	case "gift-solid":
		return GiftSolid(), nil
	case "git-branch-outline":
		return GitBranchOutline(), nil
	case "git-branch-solid":
		return GitBranchSolid(), nil
	case "git-commit-outline":
		return GitCommitOutline(), nil
	case "git-commit-solid":
		return GitCommitSolid(), nil
	case "git-compare-outline":
		return GitCompareOutline(), nil
	case "git-compare-solid":
		return GitCompareSolid(), nil
	case "git-fork-outline":
		return GitForkOutline(), nil
	case "git-fork-solid":
		return GitForkSolid(), nil
	case "git-merge-outline":
		return GitMergeOutline(), nil
	case "git-merge-solid":
		return GitMergeSolid(), nil
	case "git-outline":
		return GitOutline(), nil
	case "git-pull-outline":
		return GitPullOutline(), nil
	case "git-pull-solid":
		return GitPullSolid(), nil
	case "git-solid":
		return GitSolid(), nil
	case "github-outline":
		return GithubOutline(), nil
	case "github-solid":
		return GithubSolid(), nil
	case "gitlab-outline":
		return GitlabOutline(), nil
	case "gitlab-solid":
		return GitlabSolid(), nil
	case "globe-africa-outline":
		return GlobeAfricaOutline(), nil
	case "globe-africa-solid":
		return GlobeAfricaSolid(), nil
	case "globe-americas-outline":
		return GlobeAmericasOutline(), nil
	case "globe-americas-solid":
		return GlobeAmericasSolid(), nil
	case "globe-outline":
		return GlobeOutline(), nil
	case "globe-solid":
		return GlobeSolid(), nil
	case "google-ad-outline":
		return GoogleAdOutline(), nil
	case "google-ad-solid":
		return GoogleAdSolid(), nil
	case "google-drive-outline":
		return GoogleDriveOutline(), nil
	case "google-drive-solid":
		return GoogleDriveSolid(), nil
	case "google-outline":
		return GoogleOutline(), nil
	case "google-play-store-outline":
		return GooglePlayStoreOutline(), nil
	case "google-play-store-solid":
		return GooglePlayStoreSolid(), nil
	case "google-solid":
		return GoogleSolid(), nil
	case "google-streetview-outline":
		return GoogleStreetviewOutline(), nil
	case "google-streetview-solid":
		return GoogleStreetviewSolid(), nil
	case "graphql-outline":
		return GraphqlOutline(), nil
	case "graphql-solid":
		return GraphqlSolid(), nil
	case "grid-layout-outline":
		return GridLayoutOutline(), nil
	case "grid-layout-solid":
		return GridLayoutSolid(), nil
	case "hashtag-outline":
		return HashtagOutline(), nil
	case "hashtag-solid":
		return HashtagSolid(), nil
	case "hd-screen-outline":
		return HdScreenOutline(), nil
	case "hd-screen-solid":
		return HdScreenSolid(), nil
	case "hdmi-cable-outline":
		return HdmiCableOutline(), nil
	case "hdmi-cable-solid":
		return HdmiCableSolid(), nil
	case "headphones-outline":
		return HeadphonesOutline(), nil
	case "headphones-solid":
		return HeadphonesSolid(), nil
	case "headset-outline":
		return HeadsetOutline(), nil
	case "headset-solid":
		return HeadsetSolid(), nil
	case "heart-circle-outline":
		return HeartCircleOutline(), nil
	case "heart-circle-solid":
		return HeartCircleSolid(), nil
	case "heart-outline":
		return HeartOutline(), nil
	case "heart-small-outline":
		return HeartSmallOutline(), nil
	case "heart-small-solid":
		return HeartSmallSolid(), nil
	case "heart-solid":
		return HeartSolid(), nil
	case "hexagon-outline":
		return HexagonOutline(), nil
	case "hexagon-solid":
		return HexagonSolid(), nil
	case "history-outline":
		return HistoryOutline(), nil
	case "history-solid":
		return HistorySolid(), nil
	case "home-alt-outline":
		return HomeAltOutline(), nil
	case "home-alt-solid":
		return HomeAltSolid(), nil
	case "home-outline":
		return HomeOutline(), nil
	case "home-solid":
		return HomeSolid(), nil
	case "hospital-outline":
		return HospitalOutline(), nil
	case "hospital-solid":
		return HospitalSolid(), nil
	case "hourglass-outline":
		return HourglassOutline(), nil
	case "hourglass-solid":
		return HourglassSolid(), nil
	case "house-outline":
		return HouseOutline(), nil
	case "house-solid":
		return HouseSolid(), nil
	case "html5-outline":
		return HtmlFiveOutline(), nil
	case "html5-solid":
		return HtmlFiveSolid(), nil
	case "id-outline":
		return IdOutline(), nil
	case "id-solid":
		return IdSolid(), nil
	case "imac-outline":
		return ImacOutline(), nil
	case "imac-solid":
		return ImacSolid(), nil
	case "image-alt-outline":
		return ImageAltOutline(), nil
	case "image-alt-solid":
		return ImageAltSolid(), nil
	case "image-document-outline":
		return ImageDocumentOutline(), nil
	case "image-document-solid":
		return ImageDocumentSolid(), nil
	case "image-outline":
		return ImageOutline(), nil
	case "image-solid":
		return ImageSolid(), nil
	case "in-ear-headphones-outline":
		return InEarHeadphonesOutline(), nil
	case "in-ear-headphones-solid":
		return InEarHeadphonesSolid(), nil
	case "inbox-outline":
		return InboxOutline(), nil
	case "inbox-solid":
		return InboxSolid(), nil
	case "indent-decrease-outline":
		return IndentDecreaseOutline(), nil
	case "indent-decrease-solid":
		return IndentDecreaseSolid(), nil
	case "indent-increase-outline":
		return IndentIncreaseOutline(), nil
	case "indent-increase-solid":
		return IndentIncreaseSolid(), nil
	case "info-circle-outline":
		return InfoCircleOutline(), nil
	case "info-circle-solid":
		return InfoCircleSolid(), nil
	case "info-outline":
		return InfoOutline(), nil
	case "info-small-outline":
		return InfoSmallOutline(), nil
	case "info-small-solid":
		return InfoSmallSolid(), nil
	case "info-solid":
		return InfoSolid(), nil
	case "instagram-outline":
		return InstagramOutline(), nil
	case "instagram-solid":
		return InstagramSolid(), nil
	case "invoice-outline":
		return InvoiceOutline(), nil
	case "invoice-solid":
		return InvoiceSolid(), nil
	case "italic-outline":
		return ItalicOutline(), nil
	case "italic-solid":
		return ItalicSolid(), nil
	case "javascript-outline":
		return JavascriptOutline(), nil
	case "javascript-solid":
		return JavascriptSolid(), nil
	case "joystick-outline":
		return JoystickOutline(), nil
	case "joystick-solid":
		return JoystickSolid(), nil
	case "jpg-outline":
		return JpgOutline(), nil
	case "jpg-solid":
		return JpgSolid(), nil
	case "kanban-outline":
		return KanbanOutline(), nil
	case "kanban-solid":
		return KanbanSolid(), nil
	case "key-outline":
		return KeyOutline(), nil
	case "key-solid":
		return KeySolid(), nil
	case "keyboard-outline":
		return KeyboardOutline(), nil
	case "keyboard-solid":
		return KeyboardSolid(), nil
	case "lan-cable-outline":
		return LanCableOutline(), nil
	case "lan-cable-solid":
		return LanCableSolid(), nil
	case "laptop-outline":
		return LaptopOutline(), nil
	case "laptop-solid":
		return LaptopSolid(), nil
	case "laravel-outline":
		return LaravelOutline(), nil
	case "laravel-solid":
		return LaravelSolid(), nil
	case "layers-difference-outline":
		return LayersDifferenceOutline(), nil
	case "layers-difference-solid":
		return LayersDifferenceSolid(), nil
	case "layers-intersect-outline":
		return LayersIntersectOutline(), nil
	case "layers-intersect-solid":
		return LayersIntersectSolid(), nil
	case "layers-outline":
		return LayersOutline(), nil
	case "layers-solid":
		return LayersSolid(), nil
	case "layers-subtract-outline":
		return LayersSubtractOutline(), nil
	case "layers-subtract-solid":
		return LayersSubtractSolid(), nil
	case "layers-union-outline":
		return LayersUnionOutline(), nil
	case "layers-union-solid":
		return LayersUnionSolid(), nil
	case "left-circle-outline":
		return LeftCircleOutline(), nil
	case "left-circle-solid":
		return LeftCircleSolid(), nil
	case "left-outline":
		return LeftOutline(), nil
	case "left-small-outline":
		return LeftSmallOutline(), nil
	case "left-small-solid":
		return LeftSmallSolid(), nil
	case "left-solid":
		return LeftSolid(), nil
	case "lego-outline":
		return LegoOutline(), nil
	case "lego-solid":
		return LegoSolid(), nil
	case "lifebuoy-outline":
		return LifebuoyOutline(), nil
	case "lifebuoy-solid":
		return LifebuoySolid(), nil
	case "lightning-cable-outline":
		return LightningCableOutline(), nil
	case "lightning-cable-solid":
		return LightningCableSolid(), nil
	case "line-outline":
		return LineOutline(), nil
	case "line-solid":
		return LineSolid(), nil
	case "link-outline":
		return LinkOutline(), nil
	case "link-remove-outline":
		return LinkRemoveOutline(), nil
	case "link-remove-solid":
		return LinkRemoveSolid(), nil
	case "link-solid":
		return LinkSolid(), nil
	case "linkedin-outline":
		return LinkedinOutline(), nil
	case "linkedin-solid":
		return LinkedinSolid(), nil
	case "linux-alt-outline":
		return LinuxAltOutline(), nil
	case "linux-alt-solid":
		return LinuxAltSolid(), nil
	case "linux-outline":
		return LinuxOutline(), nil
	case "linux-solid":
		return LinuxSolid(), nil
	case "list-layout-outline":
		return ListLayoutOutline(), nil
	case "list-layout-solid":
		return ListLayoutSolid(), nil
	case "list-ordered-outline":
		return ListOrderedOutline(), nil
	case "list-ordered-solid":
		return ListOrderedSolid(), nil
	case "list-unordered-outline":
		return ListUnorderedOutline(), nil
	case "list-unordered-solid":
		return ListUnorderedSolid(), nil
	case "litecoin-outline":
		return LitecoinOutline(), nil
	case "litecoin-solid":
		return LitecoinSolid(), nil
	case "loader-outline":
		return LoaderOutline(), nil
	case "loader-solid":
		return LoaderSolid(), nil
	case "location-outline":
		return LocationOutline(), nil
	case "location-solid":
		return LocationSolid(), nil
	case "lock-circle-outline":
		return LockCircleOutline(), nil
	case "lock-circle-solid":
		return LockCircleSolid(), nil
	case "lock-outline":
		return LockOutline(), nil
	case "lock-small-outline":
		return LockSmallOutline(), nil
	case "lock-small-solid":
		return LockSmallSolid(), nil
	case "lock-solid":
		return LockSolid(), nil
	case "logout-outline":
		return LogoutOutline(), nil
	case "logout-solid":
		return LogoutSolid(), nil
	case "loop-outline":
		return LoopOutline(), nil
	case "loop-solid":
		return LoopSolid(), nil
	case "magsafe-outline":
		return MagsafeOutline(), nil
	case "magsafe-solid":
		return MagsafeSolid(), nil
	case "markdown-outline":
		return MarkdownOutline(), nil
	case "markdown-solid":
		return MarkdownSolid(), nil
	case "medium-outline":
		return MediumOutline(), nil
	case "medium-solid":
		return MediumSolid(), nil
	case "menu-outline":
		return MenuOutline(), nil
	case "menu-solid":
		return MenuSolid(), nil
	case "message-minus-outline":
		return MessageMinusOutline(), nil
	case "message-minus-solid":
		return MessageMinusSolid(), nil
	case "message-no-access-outline":
		return MessageNoAccessOutline(), nil
	case "message-no-access-solid":
		return MessageNoAccessSolid(), nil
	case "message-outline":
		return MessageOutline(), nil
	case "message-plus-outline":
		return MessagePlusOutline(), nil
	case "message-plus-solid":
		return MessagePlusSolid(), nil
	case "message-solid":
		return MessageSolid(), nil
	case "message-text-alt-outline":
		return MessageTextAltOutline(), nil
	case "message-text-alt-solid":
		return MessageTextAltSolid(), nil
	case "message-text-outline":
		return MessageTextOutline(), nil
	case "message-text-solid":
		return MessageTextSolid(), nil
	case "message-tick-outline":
		return MessageTickOutline(), nil
	case "message-tick-solid":
		return MessageTickSolid(), nil
	case "message-x-outline":
		return MessageXOutline(), nil
	case "message-x-solid":
		return MessageXSolid(), nil
	case "messenger-outline":
		return MessengerOutline(), nil
	case "messenger-solid":
		return MessengerSolid(), nil
	case "micro-sd-card-outline":
		return MicroSdCardOutline(), nil
	case "micro-sd-card-solid":
		return MicroSdCardSolid(), nil
	case "microphone-outline":
		return MicrophoneOutline(), nil
	case "microphone-solid":
		return MicrophoneSolid(), nil
	case "minimise-alt-outline":
		return MinimiseAltOutline(), nil
	case "minimise-alt-solid":
		return MinimiseAltSolid(), nil
	case "minimise-outline":
		return MinimiseOutline(), nil
	case "minimise-solid":
		return MinimiseSolid(), nil
	case "minus-circle-outline":
		return MinusCircleOutline(), nil
	case "minus-circle-solid":
		return MinusCircleSolid(), nil
	case "minus-outline":
		return MinusOutline(), nil
	case "minus-small-outline":
		return MinusSmallOutline(), nil
	case "minus-small-solid":
		return MinusSmallSolid(), nil
	case "minus-solid":
		return MinusSolid(), nil
	case "mobile-outline":
		return MobileOutline(), nil
	case "mobile-solid":
		return MobileSolid(), nil
	case "money-outline":
		return MoneyOutline(), nil
	case "money-solid":
		return MoneySolid(), nil
	case "money-stack-outline":
		return MoneyStackOutline(), nil
	case "money-stack-solid":
		return MoneyStackSolid(), nil
	case "mongodb-outline":
		return MongodbOutline(), nil
	case "mongodb-solid":
		return MongodbSolid(), nil
	case "mood-flat-outline":
		return MoodFlatOutline(), nil
	case "mood-flat-solid":
		return MoodFlatSolid(), nil
	case "mood-frown-outline":
		return MoodFrownOutline(), nil
	case "mood-frown-solid":
		return MoodFrownSolid(), nil
	case "mood-laugh-outline":
		return MoodLaughOutline(), nil
	case "mood-laugh-solid":
		return MoodLaughSolid(), nil
	case "mood-sad-outline":
		return MoodSadOutline(), nil
	case "mood-sad-solid":
		return MoodSadSolid(), nil
	case "mood-smile-outline":
		return MoodSmileOutline(), nil
	case "mood-smile-solid":
		return MoodSmileSolid(), nil
	case "mood-surprised-outline":
		return MoodSurprisedOutline(), nil
	case "mood-surprised-solid":
		return MoodSurprisedSolid(), nil
	case "mood-tongue-outline":
		return MoodTongueOutline(), nil
	case "mood-tongue-solid":
		return MoodTongueSolid(), nil
	case "moon-outline":
		return MoonOutline(), nil
	case "moon-solid":
		return MoonSolid(), nil
	case "more-horizontal-outline":
		return MoreHorizontalOutline(), nil
	case "more-horizontal-solid":
		return MoreHorizontalSolid(), nil
	case "more-vertical-outline":
		return MoreVerticalOutline(), nil
	case "more-vertical-solid":
		return MoreVerticalSolid(), nil
	case "mouse-outline":
		return MouseOutline(), nil
	case "mouse-solid":
		return MouseSolid(), nil
	case "mov-outline":
		return MovOutline(), nil
	case "mov-solid":
		return MovSolid(), nil
	case "mp4-outline":
		return MpFourOutline(), nil
	case "mp4-solid":
		return MpFourSolid(), nil
	case "mp3-outline":
		return MpThreeOutline(), nil
	case "mp3-solid":
		return MpThreeSolid(), nil
	case "ms-excel-outline":
		return MsExcelOutline(), nil
	case "ms-excel-solid":
		return MsExcelSolid(), nil
	case "ms-powerpoint-outline":
		return MsPowerpointOutline(), nil
	case "ms-powerpoint-solid":
		return MsPowerpointSolid(), nil
	case "ms-word-outline":
		return MsWordOutline(), nil
	case "ms-word-solid":
		return MsWordSolid(), nil
	case "n64-outline":
		return NSixtyFourOutline(), nil
	case "n64-solid":
		return NSixtyFourSolid(), nil
	case "nes-outline":
		return NesOutline(), nil
	case "nes-solid":
		return NesSolid(), nil
	case "netlify-outline":
		return NetlifyOutline(), nil
	case "netlify-solid":
		return NetlifySolid(), nil
	case "next-circle-outline":
		return NextCircleOutline(), nil
	case "next-circle-solid":
		return NextCircleSolid(), nil
	case "next-outline":
		return NextOutline(), nil
	case "next-small-outline":
		return NextSmallOutline(), nil
	case "next-small-solid":
		return NextSmallSolid(), nil
	case "next-solid":
		return NextSolid(), nil
	case "nextjs-outline":
		return NextjsOutline(), nil
	case "nextjs-solid":
		return NextjsSolid(), nil
	case "ngc-outline":
		return NgcOutline(), nil
	case "ngc-solid":
		return NgcSolid(), nil
	case "nintendo-switch-outline":
		return NintendoSwitchOutline(), nil
	case "nintendo-switch-solid":
		return NintendoSwitchSolid(), nil
	case "nodejs-outline":
		return NodejsOutline(), nil
	case "nodejs-solid":
		return NodejsSolid(), nil
	case "note-outline":
		return NoteOutline(), nil
	case "note-solid":
		return NoteSolid(), nil
	case "npm-outline":
		return NpmOutline(), nil
	case "npm-solid":
		return NpmSolid(), nil
	case "nuxtjs-outline":
		return NuxtjsOutline(), nil
	case "nuxtjs-solid":
		return NuxtjsSolid(), nil
	case "omega-outline":
		return OmegaOutline(), nil
	case "omega-solid":
		return OmegaSolid(), nil
	case "opera-outline":
		return OperaOutline(), nil
	case "opera-solid":
		return OperaSolid(), nil
	case "otp-outline":
		return OtpOutline(), nil
	case "otp-solid":
		return OtpSolid(), nil
	case "page-break-outline":
		return PageBreakOutline(), nil
	case "page-break-solid":
		return PageBreakSolid(), nil
	case "page-number-outline":
		return PageNumberOutline(), nil
	case "page-number-solid":
		return PageNumberSolid(), nil
	case "paintbrush-outline":
		return PaintbrushOutline(), nil
	case "paintbrush-solid":
		return PaintbrushSolid(), nil
	case "paintbucket-outline":
		return PaintbucketOutline(), nil
	case "paintbucket-solid":
		return PaintbucketSolid(), nil
	case "paragraph-outline":
		return ParagraphOutline(), nil
	case "paragraph-solid":
		return ParagraphSolid(), nil
	case "password-outline":
		return PasswordOutline(), nil
	case "password-solid":
		return PasswordSolid(), nil
	case "patreon-outline":
		return PatreonOutline(), nil
	case "patreon-solid":
		return PatreonSolid(), nil
	case "pause-circle-outline":
		return PauseCircleOutline(), nil
	case "pause-circle-solid":
		return PauseCircleSolid(), nil
	case "pause-outline":
		return PauseOutline(), nil
	case "pause-small-outline":
		return PauseSmallOutline(), nil
	case "pause-small-solid":
		return PauseSmallSolid(), nil
	case "pause-solid":
		return PauseSolid(), nil
	case "paw-outline":
		return PawOutline(), nil
	case "paw-solid":
		return PawSolid(), nil
	case "paws-outline":
		return PawsOutline(), nil
	case "paws-solid":
		return PawsSolid(), nil
	case "paypal-outline":
		return PaypalOutline(), nil
	case "paypal-solid":
		return PaypalSolid(), nil
	case "pdf-outline":
		return PdfOutline(), nil
	case "pdf-solid":
		return PdfSolid(), nil
	case "pen-outline":
		return PenOutline(), nil
	case "pen-solid":
		return PenSolid(), nil
	case "phone-outline":
		return PhoneOutline(), nil
	case "phone-solid":
		return PhoneSolid(), nil
	case "phonecall-blocked-outline":
		return PhonecallBlockedOutline(), nil
	case "phonecall-blocked-solid":
		return PhonecallBlockedSolid(), nil
	case "phonecall-outline":
		return PhonecallOutline(), nil
	case "phonecall-receive-outline":
		return PhonecallReceiveOutline(), nil
	case "phonecall-receive-solid":
		return PhonecallReceiveSolid(), nil
	case "phonecall-solid":
		return PhonecallSolid(), nil
	case "pie-chart-alt-outline":
		return PieChartAltOutline(), nil
	case "pie-chart-alt-solid":
		return PieChartAltSolid(), nil
	case "pie-chart-outline":
		return PieChartOutline(), nil
	case "pie-chart-solid":
		return PieChartSolid(), nil
	case "pin-alt-outline":
		return PinAltOutline(), nil
	case "pin-alt-solid":
		return PinAltSolid(), nil
	case "pin-outline":
		return PinOutline(), nil
	case "pin-solid":
		return PinSolid(), nil
	case "pinterest-outline":
		return PinterestOutline(), nil
	case "pinterest-solid":
		return PinterestSolid(), nil
	case "plant-outline":
		return PlantOutline(), nil
	case "plant-solid":
		return PlantSolid(), nil
	case "play-circle-outline":
		return PlayCircleOutline(), nil
	case "play-circle-solid":
		return PlayCircleSolid(), nil
	case "play-outline":
		return PlayOutline(), nil
	case "play-small-outline":
		return PlaySmallOutline(), nil
	case "play-small-solid":
		return PlaySmallSolid(), nil
	case "play-solid":
		return PlaySolid(), nil
	case "plug-outline":
		return PlugOutline(), nil
	case "plug-solid":
		return PlugSolid(), nil
	case "plus-circle-outline":
		return PlusCircleOutline(), nil
	case "plus-circle-solid":
		return PlusCircleSolid(), nil
	case "png-outline":
		return PngOutline(), nil
	case "png-solid":
		return PngSolid(), nil
	case "pool-outline":
		return PoolOutline(), nil
	case "pool-solid":
		return PoolSolid(), nil
	case "pound-outline":
		return PoundOutline(), nil
	case "pound-solid":
		return PoundSolid(), nil
	case "power-outline":
		return PowerOutline(), nil
	case "power-solid":
		return PowerSolid(), nil
	case "ppt-outline":
		return PptOutline(), nil
	case "ppt-solid":
		return PptSolid(), nil
	case "print-outline":
		return PrintOutline(), nil
	case "print-solid":
		return PrintSolid(), nil
	case "python-outline":
		return PythonOutline(), nil
	case "python-solid":
		return PythonSolid(), nil
	case "qr-code-outline":
		return QrCodeOutline(), nil
	case "qr-code-solid":
		return QrCodeSolid(), nil
	case "question-circle-outline":
		return QuestionCircleOutline(), nil
	case "question-circle-solid":
		return QuestionCircleSolid(), nil
	case "question-outline":
		return QuestionOutline(), nil
	case "question-small-outline":
		return QuestionSmallOutline(), nil
	case "question-small-solid":
		return QuestionSmallSolid(), nil
	case "question-solid":
		return QuestionSolid(), nil
	case "quote-outline":
		return QuoteOutline(), nil
	case "quote-solid":
		return QuoteSolid(), nil
	case "rand-outline":
		return RandOutline(), nil
	case "rand-solid":
		return RandSolid(), nil
	case "react-outline":
		return ReactOutline(), nil
	case "react-solid":
		return ReactSolid(), nil
	case "receipt-outline":
		return ReceiptOutline(), nil
	case "receipt-solid":
		return ReceiptSolid(), nil
	case "reddit-outline":
		return RedditOutline(), nil
	case "reddit-solid":
		return RedditSolid(), nil
	case "redwoodjs-outline":
		return RedwoodjsOutline(), nil
	case "redwoodjs-solid":
		return RedwoodjsSolid(), nil
	case "refresh-alt-outline":
		return RefreshAltOutline(), nil
	case "refresh-alt-solid":
		return RefreshAltSolid(), nil
	case "refresh-outline":
		return RefreshOutline(), nil
	case "refresh-solid":
		return RefreshSolid(), nil
	case "rewind-circle-outline":
		return RewindCircleOutline(), nil
	case "rewind-circle-solid":
		return RewindCircleSolid(), nil
	case "rewind-outline":
		return RewindOutline(), nil
	case "rewind-small-outline":
		return RewindSmallOutline(), nil
	case "rewind-small-solid":
		return RewindSmallSolid(), nil
	case "rewind-solid":
		return RewindSolid(), nil
	case "right-circle-outline":
		return RightCircleOutline(), nil
	case "right-circle-solid":
		return RightCircleSolid(), nil
	case "right-outline":
		return RightOutline(), nil
	case "right-small-outline":
		return RightSmallOutline(), nil
	case "right-small-solid":
		return RightSmallSolid(), nil
	case "right-solid":
		return RightSolid(), nil
	case "ripple-outline":
		return RippleOutline(), nil
	case "ripple-solid":
		return RippleSolid(), nil
	case "robot-outline":
		return RobotOutline(), nil
	case "robot-solid":
		return RobotSolid(), nil
	case "roller-outline":
		return RollerOutline(), nil
	case "roller-solid":
		return RollerSolid(), nil
	case "rollupjs-outline":
		return RollupjsOutline(), nil
	case "rollupjs-solid":
		return RollupjsSolid(), nil
	case "router-outline":
		return RouterOutline(), nil
	case "router-solid":
		return RouterSolid(), nil
	case "rss-outline":
		return RssOutline(), nil
	case "rss-solid":
		return RssSolid(), nil
	case "ruby-outline":
		return RubyOutline(), nil
	case "ruby-solid":
		return RubySolid(), nil
	case "rupee-outline":
		return RupeeOutline(), nil
	case "rupee-solid":
		return RupeeSolid(), nil
	case "rust-outline":
		return RustOutline(), nil
	case "rust-solid":
		return RustSolid(), nil
	case "safari-outline":
		return SafariOutline(), nil
	case "safari-solid":
		return SafariSolid(), nil
	case "safe-outline":
		return SafeOutline(), nil
	case "safe-solid":
		return SafeSolid(), nil
	case "save-outline":
		return SaveOutline(), nil
	case "save-solid":
		return SaveSolid(), nil
	case "scan-outline":
		return ScanOutline(), nil
	case "scan-solid":
		return ScanSolid(), nil
	case "school-outline":
		return SchoolOutline(), nil
	case "school-solid":
		return SchoolSolid(), nil
	case "screen-alt-outline":
		return ScreenAltOutline(), nil
	case "screen-alt-solid":
		return ScreenAltSolid(), nil
	case "screen-alt-2-outline":
		return ScreenAltTwoOutline(), nil
	case "screen-alt-2-solid":
		return ScreenAltTwoSolid(), nil
	case "screen-outline":
		return ScreenOutline(), nil
	case "screen-solid":
		return ScreenSolid(), nil
	case "scribble-outline":
		return ScribbleOutline(), nil
	case "scribble-solid":
		return ScribbleSolid(), nil
	case "sd-card-outline":
		return SdCardOutline(), nil
	case "sd-card-solid":
		return SdCardSolid(), nil
	case "search-circle-outline":
		return SearchCircleOutline(), nil
	case "search-circle-solid":
		return SearchCircleSolid(), nil
	case "search-outline":
		return SearchOutline(), nil
	case "search-property-outline":
		return SearchPropertyOutline(), nil
	case "search-property-solid":
		return SearchPropertySolid(), nil
	case "search-small-outline":
		return SearchSmallOutline(), nil
	case "search-small-solid":
		return SearchSmallSolid(), nil
	case "search-solid":
		return SearchSolid(), nil
	case "section-add-outline":
		return SectionAddOutline(), nil
	case "section-add-solid":
		return SectionAddSolid(), nil
	case "section-remove-outline":
		return SectionRemoveOutline(), nil
	case "section-remove-solid":
		return SectionRemoveSolid(), nil
	case "send-down-outline":
		return SendDownOutline(), nil
	case "send-down-solid":
		return SendDownSolid(), nil
	case "send-left-outline":
		return SendLeftOutline(), nil
	case "send-left-solid":
		return SendLeftSolid(), nil
	case "send-outline":
		return SendOutline(), nil
	case "send-right-outline":
		return SendRightOutline(), nil
	case "send-right-solid":
		return SendRightSolid(), nil
	case "send-solid":
		return SendSolid(), nil
	case "send-up-outline":
		return SendUpOutline(), nil
	case "send-up-solid":
		return SendUpSolid(), nil
	case "servers-outline":
		return ServersOutline(), nil
	case "servers-solid":
		return ServersSolid(), nil
	case "share-outline":
		return ShareOutline(), nil
	case "share-solid":
		return ShareSolid(), nil
	case "shield-outline":
		return ShieldOutline(), nil
	case "shield-solid":
		return ShieldSolid(), nil
	case "shield-tick-outline":
		return ShieldTickOutline(), nil
	case "shield-tick-solid":
		return ShieldTickSolid(), nil
	case "shield-x-outline":
		return ShieldXOutline(), nil
	case "shield-x-solid":
		return ShieldXSolid(), nil
	case "shop-outline":
		return ShopOutline(), nil
	case "shop-solid":
		return ShopSolid(), nil
	case "sign-outline":
		return SignOutline(), nil
	case "sign-solid":
		return SignSolid(), nil
	case "signin-outline":
		return SigninOutline(), nil
	case "signin-solid":
		return SigninSolid(), nil
	case "sim-outline":
		return SimOutline(), nil
	case "sim-solid":
		return SimSolid(), nil
	case "simohamed-outline":
		return SimohamedOutline(), nil
	case "simohamed-solid":
		return SimohamedSolid(), nil
	case "skull-outline":
		return SkullOutline(), nil
	case "skull-solid":
		return SkullSolid(), nil
	case "skype-outline":
		return SkypeOutline(), nil
	case "skype-solid":
		return SkypeSolid(), nil
	case "slack-outline":
		return SlackOutline(), nil
	case "slack-solid":
		return SlackSolid(), nil
	case "snapchat-outline":
		return SnapchatOutline(), nil
	case "snapchat-solid":
		return SnapchatSolid(), nil
	case "snes-outline":
		return SnesOutline(), nil
	case "snes-solid":
		return SnesSolid(), nil
	case "sort-alphabetically-outline":
		return SortAlphabeticallyOutline(), nil
	case "sort-alphabetically-solid":
		return SortAlphabeticallySolid(), nil
	case "sort-down-outline":
		return SortDownOutline(), nil
	case "sort-down-solid":
		return SortDownSolid(), nil
	case "sort-high-to-low-outline":
		return SortHighToLowOutline(), nil
	case "sort-high-to-low-solid":
		return SortHighToLowSolid(), nil
	case "sort-low-to-high-outline":
		return SortLowToHighOutline(), nil
	case "sort-low-to-high-solid":
		return SortLowToHighSolid(), nil
	case "sort-reverse-alphabetically-outline":
		return SortReverseAlphabeticallyOutline(), nil
	case "sort-reverse-alphabetically-solid":
		return SortReverseAlphabeticallySolid(), nil
	case "sort-up-outline":
		return SortUpOutline(), nil
	case "sort-up-solid":
		return SortUpSolid(), nil
	case "sound-off-outline":
		return SoundOffOutline(), nil
	case "sound-off-solid":
		return SoundOffSolid(), nil
	case "sound-on-outline":
		return SoundOnOutline(), nil
	case "sound-on-solid":
		return SoundOnSolid(), nil
	case "spotify-outline":
		return SpotifyOutline(), nil
	case "spotify-solid":
		return SpotifySolid(), nil
	case "spreadsheet-outline":
		return SpreadsheetOutline(), nil
	case "spreadsheet-solid":
		return SpreadsheetSolid(), nil
	case "square-outline":
		return SquareOutline(), nil
	case "square-solid":
		return SquareSolid(), nil
	case "stackoverflow-outline":
		return StackoverflowOutline(), nil
	case "stackoverflow-solid":
		return StackoverflowSolid(), nil
	case "stamp-outline":
		return StampOutline(), nil
	case "stamp-solid":
		return StampSolid(), nil
	case "star-circle-outline":
		return StarCircleOutline(), nil
	case "star-circle-solid":
		return StarCircleSolid(), nil
	case "star-outline":
		return StarOutline(), nil
	case "star-small-outline":
		return StarSmallOutline(), nil
	case "star-small-solid":
		return StarSmallSolid(), nil
	case "star-solid":
		return StarSolid(), nil
	case "stop-circle-outline":
		return StopCircleOutline(), nil
	case "stop-circle-solid":
		return StopCircleSolid(), nil
	case "stop-outline":
		return StopOutline(), nil
	case "stop-small-outline":
		return StopSmallOutline(), nil
	case "stop-small-solid":
		return StopSmallSolid(), nil
	case "stop-solid":
		return StopSolid(), nil
	case "stopwatch-outline":
		return StopwatchOutline(), nil
	case "stopwatch-solid":
		return StopwatchSolid(), nil
	case "strikethrough-outline":
		return StrikethroughOutline(), nil
	case "strikethrough-solid":
		return StrikethroughSolid(), nil
	case "subscript-outline":
		return SubscriptOutline(), nil
	case "subscript-solid":
		return SubscriptSolid(), nil
	case "sun-outline":
		return SunOutline(), nil
	case "sun-solid":
		return SunSolid(), nil
	case "superscript-outline":
		return SuperscriptOutline(), nil
	case "superscript-solid":
		return SuperscriptSolid(), nil
	case "svelte-outline":
		return SvelteOutline(), nil
	case "svelte-solid":
		return SvelteSolid(), nil
	case "svg-outline":
		return SvgOutline(), nil
	case "svg-solid":
		return SvgSolid(), nil
	case "table-outline":
		return TableOutline(), nil
	case "table-solid":
		return TableSolid(), nil
	case "tablet-outline":
		return TabletOutline(), nil
	case "tablet-solid":
		return TabletSolid(), nil
	case "tag-outline":
		return TagOutline(), nil
	case "tag-solid":
		return TagSolid(), nil
	case "tailwind-outline":
		return TailwindOutline(), nil
	case "tailwind-solid":
		return TailwindSolid(), nil
	case "target-outline":
		return TargetOutline(), nil
	case "target-solid":
		return TargetSolid(), nil
	case "telegram-outline":
		return TelegramOutline(), nil
	case "telegram-solid":
		return TelegramSolid(), nil
	case "terminal-outline":
		return TerminalOutline(), nil
	case "terminal-solid":
		return TerminalSolid(), nil
	case "text-document-alt-outline":
		return TextDocumentAltOutline(), nil
	case "text-document-alt-solid":
		return TextDocumentAltSolid(), nil
	case "text-document-outline":
		return TextDocumentOutline(), nil
	case "text-document-solid":
		return TextDocumentSolid(), nil
	case "text-outline":
		return TextOutline(), nil
	case "text-solid":
		return TextSolid(), nil
	case "360-outline":
		return ThreeHundredSixtyOutline(), nil
	case "360-solid":
		return ThreeHundredSixtySolid(), nil
	case "thumb-down-outline":
		return ThumbDownOutline(), nil
	case "thumb-down-solid":
		return ThumbDownSolid(), nil
	case "thumb-up-outline":
		return ThumbUpOutline(), nil
	case "thumb-up-solid":
		return ThumbUpSolid(), nil
	case "thumbtack-outline":
		return ThumbtackOutline(), nil
	case "thumbtack-solid":
		return ThumbtackSolid(), nil
	case "tick-circle-outline":
		return TickCircleOutline(), nil
	case "tick-circle-solid":
		return TickCircleSolid(), nil
	case "tick-outline":
		return TickOutline(), nil
	case "tick-small-outline":
		return TickSmallOutline(), nil
	case "tick-small-solid":
		return TickSmallSolid(), nil
	case "tick-solid":
		return TickSolid(), nil
	case "tiktok-outline":
		return TiktokOutline(), nil
	case "tiktok-solid":
		return TiktokSolid(), nil
	case "toggle-outline":
		return ToggleOutline(), nil
	case "toggle-solid":
		return ToggleSolid(), nil
	case "top-left-outline":
		return TopLeftOutline(), nil
	case "top-left-solid":
		return TopLeftSolid(), nil
	case "top-right-outline":
		return TopRightOutline(), nil
	case "top-right-solid":
		return TopRightSolid(), nil
	case "trend-down-outline":
		return TrendDownOutline(), nil
	case "trend-down-solid":
		return TrendDownSolid(), nil
	case "trend-up-outline":
		return TrendUpOutline(), nil
	case "trend-up-solid":
		return TrendUpSolid(), nil
	case "triangle-outline":
		return TriangleOutline(), nil
	case "triangle-solid":
		return TriangleSolid(), nil
	case "trophy-outline":
		return TrophyOutline(), nil
	case "trophy-solid":
		return TrophySolid(), nil
	case "tv-outline":
		return TvOutline(), nil
	case "tv-solid":
		return TvSolid(), nil
	case "twitch-outline":
		return TwitchOutline(), nil
	case "twitch-solid":
		return TwitchSolid(), nil
	case "twitter-outline":
		return TwitterOutline(), nil
	case "twitter-solid":
		return TwitterSolid(), nil
	case "typescript-outline":
		return TypescriptOutline(), nil
	case "typescript-solid":
		return TypescriptSolid(), nil
	case "underline-outline":
		return UnderlineOutline(), nil
	case "underline-solid":
		return UnderlineSolid(), nil
	case "unlock-circle-outline":
		return UnlockCircleOutline(), nil
	case "unlock-circle-solid":
		return UnlockCircleSolid(), nil
	case "unlock-outline":
		return UnlockOutline(), nil
	case "unlock-small-outline":
		return UnlockSmallOutline(), nil
	case "unlock-small-solid":
		return UnlockSmallSolid(), nil
	case "unlock-solid":
		return UnlockSolid(), nil
	case "up-circle-outline":
		return UpCircleOutline(), nil
	case "up-circle-solid":
		return UpCircleSolid(), nil
	case "up-outline":
		return UpOutline(), nil
	case "up-small-outline":
		return UpSmallOutline(), nil
	case "up-small-solid":
		return UpSmallSolid(), nil
	case "up-solid":
		return UpSolid(), nil
	case "upload-outline":
		return UploadOutline(), nil
	case "upload-solid":
		return UploadSolid(), nil
	case "usb-cable-outline":
		return UsbCableOutline(), nil
	case "usb-cable-solid":
		return UsbCableSolid(), nil
	case "user-circle-outline":
		return UserCircleOutline(), nil
	case "user-circle-solid":
		return UserCircleSolid(), nil
	case "user-minus-outline":
		return UserMinusOutline(), nil
	case "user-minus-solid":
		return UserMinusSolid(), nil
	case "user-outline":
		return UserOutline(), nil
	case "user-plus-outline":
		return UserPlusOutline(), nil
	case "user-plus-solid":
		return UserPlusSolid(), nil
	case "user-solid":
		return UserSolid(), nil
	case "user-square-outline":
		return UserSquareOutline(), nil
	case "user-square-solid":
		return UserSquareSolid(), nil
	case "users-outline":
		return UsersOutline(), nil
	case "users-solid":
		return UsersSolid(), nil
	case "vector-document-outline":
		return VectorDocumentOutline(), nil
	case "vector-document-solid":
		return VectorDocumentSolid(), nil
	case "venn-diagram-outline":
		return VennDiagramOutline(), nil
	case "venn-diagram-solid":
		return VennDiagramSolid(), nil
	case "view-column-outline":
		return ViewColumnOutline(), nil
	case "view-column-solid":
		return ViewColumnSolid(), nil
	case "view-grid-outline":
		return ViewGridOutline(), nil
	case "view-grid-solid":
		return ViewGridSolid(), nil
	case "vim-outline":
		return VimOutline(), nil
	case "vim-solid":
		return VimSolid(), nil
	case "volume-1-outline":
		return VolumeOneOutline(), nil
	case "volume-1-solid":
		return VolumeOneSolid(), nil
	case "volume-3-outline":
		return VolumeThreeOutline(), nil
	case "volume-3-solid":
		return VolumeThreeSolid(), nil
	case "volume-2-outline":
		return VolumeTwoOutline(), nil
	case "volume-2-solid":
		return VolumeTwoSolid(), nil
	case "vr-headset-outline":
		return VrHeadsetOutline(), nil
	case "vr-headset-solid":
		return VrHeadsetSolid(), nil
	case "vue-outline":
		return VueOutline(), nil
	case "vue-solid":
		return VueSolid(), nil
	case "wallet-alt-outline":
		return WalletAltOutline(), nil
	case "wallet-alt-solid":
		return WalletAltSolid(), nil
	case "wallet-outline":
		return WalletOutline(), nil
	case "wallet-solid":
		return WalletSolid(), nil
	case "wan-outline":
		return WanOutline(), nil
	case "wan-solid":
		return WanSolid(), nil
	case "wand-outline":
		return WandOutline(), nil
	case "wand-solid":
		return WandSolid(), nil
	case "watch-outline":
		return WatchOutline(), nil
	case "watch-solid":
		return WatchSolid(), nil
	case "webpack-outline":
		return WebpackOutline(), nil
	case "webpack-solid":
		return WebpackSolid(), nil
	case "whatsapp-outline":
		return WhatsappOutline(), nil
	case "whatsapp-solid":
		return WhatsappSolid(), nil
	case "wifi-full-outline":
		return WifiFullOutline(), nil
	case "wifi-full-solid":
		return WifiFullSolid(), nil
	case "wifi-low-outline":
		return WifiLowOutline(), nil
	case "wifi-low-solid":
		return WifiLowSolid(), nil
	case "wifi-none-outline":
		return WifiNoneOutline(), nil
	case "wifi-none-solid":
		return WifiNoneSolid(), nil
	case "windows-outline":
		return WindowsOutline(), nil
	case "windows-solid":
		return WindowsSolid(), nil
	case "wordpress-outline":
		return WordpressOutline(), nil
	case "wordpress-solid":
		return WordpressSolid(), nil
	case "x-circle-outline":
		return XCircleOutline(), nil
	case "x-circle-solid":
		return XCircleSolid(), nil
	case "x-outline":
		return XOutline(), nil
	case "x-small-outline":
		return XSmallOutline(), nil
	case "x-small-solid":
		return XSmallSolid(), nil
	case "x-solid":
		return XSolid(), nil
	case "xls-outline":
		return XlsOutline(), nil
	case "xls-solid":
		return XlsSolid(), nil
	case "yen-outline":
		return YenOutline(), nil
	case "yen-solid":
		return YenSolid(), nil
	case "youtube-outline":
		return YoutubeOutline(), nil
	case "youtube-solid":
		return YoutubeSolid(), nil
	case "zip-outline":
		return ZipOutline(), nil
	case "zip-solid":
		return ZipSolid(), nil
	case "zoom-in-outline":
		return ZoomInOutline(), nil
	case "zoom-in-solid":
		return ZoomInSolid(), nil
	case "zoom-out-outline":
		return ZoomOutOutline(), nil
	case "zoom-out-solid":
		return ZoomOutSolid(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in teenyicons icon set", name)
	}
}
