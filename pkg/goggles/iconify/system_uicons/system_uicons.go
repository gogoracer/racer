package system_uicons

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	airplayInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6.5 14.5l-1-.035c-1.102-.003-2-.932-2-2.034V6.465a2 2 0 0 1 2-2l10-.002a2 2 0 0 1 2 2v6.002a2 2 0 0 1-2 2l-1 .037"/><path d="m10.5 13.5l-3 3h6z"/></g>`
	alarmClockInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 4.565h-2a6 6 0 0 0-6 6V12.5a6 6 0 0 0 6 6h2a6 6 0 0 0 6-6v-1.935a6 6 0 0 0-6-6zm3.032-1.068c.884-.639 2.089-.71 2.968.003c.906.734 1.258 1.96.822 2.969M6.532 3.544C5.642 2.862 4.4 2.77 3.5 3.5c-.906.734-1.258 1.96-.822 2.97"/><path d="M10.5 7.5v4H14M5 17l-2 2m13-2l2 2"/></g>`
	alignHorizontalInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 5.5l3 3l3-3m-3-4v7m-3 7l3-3l3 3m-3-3v7m-7-9h14"/>`
	alignVerticalInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.5 7.5l-3 3l3 3m4-3h-7m-7-3l3 3l-3 3m3-3h-7m9-7v14"/>`
	angleInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 16.5a5 5 0 0 0-5-5"/><path d="M5.5 5.5v11h11"/></g>`
	archiveInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 7.5h14v7.998a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2zm0-3.978h14a1 1 0 0 1 1 1V6.5a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V4.522a1 1 0 0 1 1-1zm5 6.978h4"/>`
	arrowBottomLeftInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 7.5v7h7m1-8l-8 8"/>`
	arrowBottomRightInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 7.5v7h-7m-1-8l8 8"/>`
	arrowDownInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 13.499l4 4.001l4-4.001m-4 4.001v-13"/>`
	arrowDownCircleInnerSVG     = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m5.5 9.5l3 3l3-3m-3 3v-8"/></g>`
	arrowLeftInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.499 6.497L3.5 10.499l4 4.001m9-4h-13"/>`
	arrowLeftCircleInnerSVG     = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m7.5 11.5l-3-3l3-3m5 3h-8"/></g>`
	arrowRightInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5 6.497l4 4.002l-4 4.001m-9-4h13"/>`
	arrowRightCircleInnerSVG    = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m9.5 11.5l3-3l-3-3m3 3h-8"/></g>`
	arrowTopLeftInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 13.5v-7h7m-7 0l8 8"/>`
	arrowTopRightInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 13.5v-7h-7m7 0l-8 8"/>`
	arrowUpInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.5 7.5l-4-4l-4.029 4m4.029-4v13"/>`
	arrowUpCircleInnerSVG       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m11.5 7.5l-3-3l-3 3m3-3v8"/></g>`
	audioWaveInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 8.5v4m2-6v9m2-6v2m2-4v6.814m2-9.814v12"/>`
	backspaceInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.328 15.5H15.5a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2H8.328a2 2 0 0 0-1.414.586L3.207 9.793a1 1 0 0 0 0 1.414l3.707 3.707a2 2 0 0 0 1.414.586zm1.172-3l4-4m-4 0l4 4"/>`
	backwardInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 7.5c5.185-.47 8.52 1.53 10 6c-2.825-3.14-6.341-3.718-10-2v3l-5-5l5-5z"/>`
	bagInnerSVG                 = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.426 4.503L14.52 4.5a1 1 0 0 1 .997.92l.894 10.999a1 1 0 0 1-.916 1.078l-.08.003H5.58a1 1 0 0 1-1-1l.003-.077l.846-10.997a1 1 0 0 1 .997-.923z"/><path d="M13.5 8.5v.645c0 1.105-1.895 1.355-3 1.355s-3-.395-3-1.5v-.5"/></g>`
	batteryChargingInnerSVG     = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m12.5 6.5l2-.001c1.105-.002 2 .893 2.001 1.997l-.001 3.001a2 2 0 0 1-2 2l-2 .003m-5 0H4.487a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2H6.5"/><path d="M10.5 9.5H13L9.4 16l.1-5.5h-3l4-6zm8-1v3"/></g>`
	batteryEmptyInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h10a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2zm14 2v3"/>`
	batteryFullInnerSVG         = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h10a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M5 8h9a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 8.5v3"/></g>`
	batteryHalfInnerSVG         = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h10a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M5 8h4a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 8.5v3"/></g>`
	batteryLowInnerSVG          = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.497h10a2 2 0 0 1 2 2V11.5a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2V8.497a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M5 8a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 8.5v3"/></g>`
	batterySeventyFiveInnerSVG  = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h10a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M5 8h7a1 1 0 0 1 1 1v2.046a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 8.5v3"/></g>`
	bellInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.585 15.5H5.415A1.65 1.65 0 0 1 4 13a10.526 10.526 0 0 0 1.5-5.415V6.5a4 4 0 0 1 4-4h2a4 4 0 0 1 4 4v1.085c0 1.907.518 3.78 1.5 5.415a1.65 1.65 0 0 1-1.415 2.5zM13 17c-.667 1-1.5 1.5-2.5 1.5S8.667 18 8 17"/>`
	bellDisabledInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 15.5H5.415A1.65 1.65 0 0 1 4 13a10.526 10.526 0 0 0 1.5-5.415V6.5c0-.274-.053-.741 0-1m1.363-2.008A3.985 3.985 0 0 1 9.5 2.5h2a4 4 0 0 1 4 4v1.085c0 1.907.518 3.78 1.5 5.415c.238.397.29.854.181 1.269M17.5 17.5l-14-14M13 17c-.667 1-1.5 1.5-2.5 1.5S8.667 18 8 17"/>`
	bellRingingInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.585 15.5H5.415A1.65 1.65 0 0 1 4 13a10.526 10.526 0 0 0 1.5-5.415V6.5a4 4 0 0 1 4-4h2a4 4 0 0 1 4 4v1.085c0 1.907.518 3.78 1.5 5.415a1.65 1.65 0 0 1-1.415 2.5zm1.915-11c-.267-.934-.6-1.6-1-2s-1.066-.733-2-1m-10.912 3c.209-.934.512-1.6.912-2s1.096-.733 2.088-1M13 17c-.667 1-1.5 1.5-2.5 1.5S8.667 18 8 17"/>`
	bellSnoozeInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.5 7.585c0 1.907.518 3.78 1.5 5.415a1.65 1.65 0 0 1-1.416 2.5H5.415A1.65 1.65 0 0 1 4 13a10.526 10.526 0 0 0 1.5-5.415V6.5a4 4 0 0 1 4-4h2a3.98 3.98 0 0 1 2.178.645"/><path d="M10.5 5.5h2l-2 3h2m2-7h3l-3 4h3M13 17c-.667 1-1.5 1.5-2.5 1.5S8.667 18 8 17"/></g>`
	bluetoothInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7 6.75l7 7.5L10.5 18V3L14 6.75l-7 7.5"/>`
	bookInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 6.59c-1.333-.726-2.667-1.09-4-1.09s-2.667.364-4 1.09v9.91c1.333-.667 2.667-1 4-1s2.667.333 4 1zm-8 0c-1.333-.726-2.667-1.09-4-1.09s-2.667.364-4 1.09v9.91c1.333-.667 2.667-1 4-1s2.667.333 4 1z"/>`
	bookClosedInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 3.5h9a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-9a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path d="M6.5 13.5h10v3a1 1 0 0 1-1 1h-9a2 2 0 1 1 0-4z"/></g>`
	bookTextInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 6.59c-1.333-.726-2.667-1.09-4-1.09s-2.667.364-4 1.09v9.91c1.333-.667 2.667-1 4-1s2.667.333 4 1z"/><path d="M16.556 7.788c-.685-.192-1.37-.288-2.056-.288s-1.37.096-2.056.288m4.112 2c-.685-.192-1.37-.288-2.056-.288s-1.37.096-2.056.288m4.112 2c-.685-.192-1.37-.288-2.056-.288s-1.37.096-2.056.288m4.112 2c-.685-.192-1.37-.288-2.056-.288s-1.37.096-2.056.288m-3.888-6C7.87 7.596 7.186 7.5 6.5 7.5s-1.37.096-2.056.288m4.112 2C7.87 9.596 7.186 9.5 6.5 9.5s-1.37.096-2.056.288m4.112 2c-.685-.192-1.37-.288-2.056-.288s-1.37.096-2.056.288m4.112 2c-.685-.192-1.37-.288-2.056-.288s-1.37.096-2.056.288"/><path d="M10.5 6.59c-1.333-.726-2.667-1.09-4-1.09s-2.667.364-4 1.09v9.91c1.333-.667 2.667-1 4-1s2.667.333 4 1z"/></g>`
	bookmarkInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 4.5h6a1 1 0 0 1 1 1v12l-4-4l-4 4v-12a1 1 0 0 1 1-1z"/>`
	bookmarkBookInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 3.5h8a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path d="M7.5 3.5h4v5.012L9.5 6.5l-2 2.012z"/></g>`
	boxInnerSVG                 = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.492 4.067l5 2.857A2 2 0 0 1 17.5 8.661v4.678a2 2 0 0 1-1.008 1.737l-5 2.857a2 2 0 0 1-1.984 0l-5-2.857A2 2 0 0 1 3.5 13.339V8.661a2 2 0 0 1 1.008-1.737l5-2.857a2 2 0 0 1 1.984 0zM14 9.5l-7-4"/><path d="m4 8l5.552 2.99a2 2 0 0 0 1.896 0L17 8m-6.5 3.5V18"/></g>`
	boxAddInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.492 4.067l5 2.857A2 2 0 0 1 17.5 8.661v4.678a2 2 0 0 1-1.008 1.737l-5 2.857a2 2 0 0 1-1.984 0l-5-2.857A2 2 0 0 1 3.5 13.339V8.661a2 2 0 0 1 1.008-1.737l5-2.857a2 2 0 0 1 1.984 0zM17.5 1.5v4m2-2h-4m-1.5 6l-7-4"/><path d="m4 8l5.552 2.99a2 2 0 0 0 1.896 0L17 8m-6.5 3.5V18"/></g>`
	boxDownloadInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 13.5v-8a2 2 0 0 0-2-2h-3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-10a2 2 0 0 0-2-2h-3"/><path d="m7.5 10.5l3 3l3-3"/></g>`
	boxOpenInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 7.5l7-4l5.992 3.424A2 2 0 0 1 17.5 8.661v4.678a2 2 0 0 1-1.008 1.737l-5 2.857a2 2 0 0 1-1.984 0l-5-2.857A2 2 0 0 1 3.5 13.339v-2.802"/><path d="M9.552 10.99a2 2 0 0 0 1.896 0L17 8m-6.5 3.5V18"/><path d="m3.5 7.5l7 4l-3 1l-7-4zm7-4l7 4l2-2l-7-4z"/></g>`
	boxRemoveInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.492 4.067l5 2.857A2 2 0 0 1 17.5 8.661v4.678a2 2 0 0 1-1.008 1.737l-5 2.857a2 2 0 0 1-1.984 0l-5-2.857A2 2 0 0 1 3.5 13.339V8.661a2 2 0 0 1 1.008-1.737l5-2.857a2 2 0 0 1 1.984 0zM19.5 3.5h-4m-1.5 6l-7-4"/><path d="m4 8l5.552 2.99a2 2 0 0 0 1.896 0L17 8m-6.5 3.5V18"/></g>`
	boxesInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.5 15.429l3.548 1.837a1 1 0 0 0 .907.006l2.992-1.496a1 1 0 0 0 .553-.894v-2.764a1 1 0 0 0-.553-.894L14.5 9.5l-3.46 1.792a1 1 0 0 0-.54.888v3.249z"/><path d="m3.04 15.708l3.008 1.558a1 1 0 0 0 .907.006L10.5 15.5v-3.382a1 1 0 0 0-.553-.894L6.5 9.5l-3.46 1.792a1 1 0 0 0-.54.888v2.64a1 1 0 0 0 .54.888zM6.5 9.429l3.548 1.837a1 1 0 0 0 .907.006L14.5 9.5V6.118a1 1 0 0 0-.553-.894l-2.992-1.496a1 1 0 0 0-.907.006L7.04 5.292a1 1 0 0 0-.54.888v3.249z"/><path d="m6.846 5.673l3.207 1.603a1 1 0 0 0 .894 0L14.12 5.69h0M8.799 4.649L12.5 6.5m.299 4.149L16.5 12.5M4.799 10.649L8.5 12.5m2.346-.827l3.207 1.603a1 1 0 0 0 .894 0l3.172-1.586h0m-15.273-.017l3.207 1.603a1 1 0 0 0 .894 0l3.172-1.586h0M10.5 7.5v4m4 2V17m-8-3.5V17"/></g>`
	branchInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.5 8.5v-5h5"/><path d="m4.5 3.5l6 6v8m2-10l4-4"/></g>`
	briefcaseInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 7.5h10a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2zm4-3h2a2 2 0 0 1 2 2v1h-6v-1a2 2 0 0 1 2-2z"/>`
	browserInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 14.5v-9a2 2 0 0 0-2-2h-12a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2z"/><path d="M16.5 13.5v-3a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1zm-8-7.002h8M4.5 6.5h2"/></g>`
	browserAltInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 14.5v-9a2 2 0 0 0-2-2h-12a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2z"/><path d="M8.5 13.5v-7a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1zm8-6v-1a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1zm0 6v-2a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1z"/></g>`
	buttonAddInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 14.5v-8a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zm-6-7v6.056m3-3.056h-6"/>`
	buttonMinusInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 14.5v-8a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zm-3-4h-6"/>`
	calculatorInnerSVG          = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 4.5h6a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2zm-2 5h10"/><g fill="currentColor" transform="translate(5 4)"><circle cx="2.5" cy="7.5" r="1"/><circle cx="4.5" cy="7.5" r="1"/><circle cx="6.5" cy="7.5" r="1"/><circle cx="8.5" cy="7.5" r="1"/><circle cx="2.5" cy="9.5" r="1"/><circle cx="4.5" cy="9.5" r="1"/><circle cx="6.5" cy="9.5" r="1"/><circle cx="8.5" cy="9.5" r="1"/><circle cx="2.5" cy="11.5" r="1"/><circle cx="4.5" cy="11.5" r="1"/><circle cx="6.5" cy="11.5" r="1"/><circle cx="8.5" cy="11.5" r="1"/></g></g>`
	calendarInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2zm-2 4h16"/>`
	calendarAddInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2zm-2 4h16m-8 3v6.056m3-3.056h-6"/>`
	calendarDateInnerSVG        = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v11.99a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2V4.5a2 2 0 0 1 2-2zm-1.841 4H18.5"/><path fill="currentColor" d="M6.816 13.155v-1.079h.88c.668 0 1.122-.395 1.122-.972c0-.527-.415-.927-1.103-.927c-.713 0-1.152.366-1.201.996H5.15C5.201 9.874 6.201 9 7.788 9c1.563 0 2.432.864 2.427 1.895c-.005.854-.542 1.416-1.299 1.601v.093c.981.141 1.577.766 1.577 1.709c0 1.235-1.162 2.11-2.754 2.11S5.063 15.537 5 14.204h1.411c.044.596.552.977 1.309.977c.747 0 1.27-.406 1.27-1.016c0-.625-.489-1.01-1.28-1.01zm6.7 3.072v-5.611h-.087L11.7 11.808v-1.372l1.821-1.255h1.47v7.046z"/></g>`
	calendarDayInnerSVG         = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h12.027a2 2 0 0 1 2 2v11.99a2 2 0 0 1-1.85 1.995l-.16.006l-12.027-.058a2 2 0 0 1-1.99-2V2.5a2 2 0 0 1 2-2zm-2 4h16.027"/><circle cx="4.5" cy="8.5" r="1" fill="currentColor"/></g>`
	calendarDaysInnerSVG        = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2zm-2 4h16"/><g fill="currentColor" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="1"/><circle cx="4.5" cy="8.5" r="1"/><circle cx="4.5" cy="12.5" r="1"/></g></g>`
	calendarLastDayInnerSVG     = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h12.027a2 2 0 0 1 2 2v11.99a2 2 0 0 1-1.85 1.995l-.16.006l-12.027-.058a2 2 0 0 1-1.99-2V2.5a2 2 0 0 1 2-2zm-2 4h16.027"/><circle cx="12.5" cy="12.5" r="1" fill="currentColor"/></g>`
	calendarMonthInnerSVG       = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2zm-2 4h16"/><g fill="currentColor" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="1"/><circle cx="4.5" cy="8.5" r="1"/><circle cx="12.5" cy="8.5" r="1"/><circle cx="8.5" cy="12.5" r="1"/><circle cx="4.5" cy="12.5" r="1"/><circle cx="12.5" cy="12.5" r="1"/></g></g>`
	calendarMoveInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 10.462v-6.03a2 2 0 0 1 2-2h.01l12 .058a2 2 0 0 1 1.99 2v12a2 2 0 0 1-2 2h-.01l-12-.057a2 2 0 0 1-1.99-2V14.5m0-8h16"/><path d="m8.5 15.5l3-3l-3-3m3 3h-10"/></g>`
	calendarRemoveInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2zm-2 4h16m-5 6h-6"/>`
	calendarSplitInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2zm-2 4h16m-8 0v12"/>`
	calendarWeekInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 2.5h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2zm-2 4h16m-13 3v6m2-6v6m2-6v6m2-6v6m2-6v6m2-6v6"/>`
	cameraInnerSVG              = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 14.5v-6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M17 9a1 1 0 1 0-2 0a1 1 0 0 0 2 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 11.5a3 3 0 1 0-6 0a3 3 0 0 0 6 0zm-4-7h2a1 1 0 0 1 1 1v1h-4v-1a1 1 0 0 1 1-1z"/></g>`
	cameraAltInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 14.5v-6a2 2 0 0 1 2-2h2l2.079-2h3.92l1.92 2H16.5a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2z"/><path d="M13.5 11.5a3 3 0 1 0-6 0a3 3 0 0 0 6 0z"/></g>`
	cameraNoflashInnerSVG       = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 6.5h8a2 2 0 0 1 2 2v6c0 .559-.229 1.064-.598 1.427M15.5 16.5h-11a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h1m-2-2l14 14"/><path fill="currentColor" d="M17 9a1 1 0 1 0-2 0a1 1 0 0 0 2 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.215 9.557a3 3 0 0 0 4.27 4.193M13.5 11.5a3 3 0 0 0-3-3m-1-4h2a1 1 0 0 1 1 1v1h-4v-1a1 1 0 0 1 1-1z"/></g>`
	cameraNoflashAltInnerSVG    = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m7.54 5.5l1.039-1h3.92l1.92 2H16.5a2 2 0 0 1 2 2v6c0 .502-.185.961-.49 1.312m-2.51.688h-11a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h1m-2-2l14 14"/><path d="M8.306 9.454a3 3 0 0 0 4.202 4.275M13.5 11.5a3 3 0 0 0-3-3"/></g>`
	captureInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 7.5V4.494a2 2 0 0 0-1.994-2L13.5 2.485m5 11.015v3a2 2 0 0 1-2 2h-3m-6-16.015l-3.006.01a2 2 0 0 0-1.994 2V7.5m5 11h-3a2 2 0 0 1-2-2v-3"/>`
	cardTimelineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 5.5h-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1zm13 0h-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1zm-5-1h-4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-10a1 1 0 0 0-1-1z"/>`
	cardViewInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 3.5h6a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm8 2.5h1a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-1zm-10 0h-1a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h1z"/>`
	carouselInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 5.5h-8a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zm4 0v10m-16-10v10"/>`
	cartInnerSVG                = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 6.5h12.5l-1.586 5.55a2 2 0 0 1-1.923 1.45h-6.7a2 2 0 0 1-1.989-1.78L4.5 4.5h-2"/><g fill="currentColor" transform="translate(2 4)"><circle cx="5" cy="12" r="1"/><circle cx="13" cy="12" r="1"/></g></g>`
	castInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 6.5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2m-2 0a8 8 0 0 0-8-8m5 8a5 5 0 0 0-5-5m2 5a2 2 0 0 0-2-2"/>`
	chainInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 11.5c.97 1.367 3.011 1.127 4.011 0l1.989-2c1.124-1.228 1.164-2.814 0-4c-1.136-1.157-2.864-1.157-4 0l-2 2"/><path d="M11.5 10.57c-.97-1.367-3-1.197-4-.07l-2 1.975c-1.124 1.228-1.164 2.839 0 4.025c1.136 1.157 2.864 1.157 4 0l2-2"/></g>`
	chatAddInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.418 4.214A9.283 9.283 0 0 0 10.5 3.75c-4.418 0-8 3.026-8 6.759c0 1.457.546 2.807 1.475 3.91L3 19l3.916-2.447a9.181 9.181 0 0 0 3.584.714c4.418 0 8-3.026 8-6.758c0-.685-.12-1.346-.345-1.969M16.5 3.5v4m2-2h-4"/>`
	checkInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 11.5l3 3l8.028-8"/>`
	checkCircleInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m5.5 9.5l2 2l5-5"/></g>`
	checkCircleOutsideInnerSVG  = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.857 3.79a8 8 0 1 0 2.852 3.24"/><path d="m6.5 9.5l3 3l8-8"/></g>`
	checkboxCheckedInnerSVG     = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path d="m7.5 10.5l2 2l4-4"/></g>`
	checkboxEmptyInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/>`
	chevronCloseInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 15.5l3-3l3 3m-6-9l3 3l3-3"/>`
	chevronDownInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.5 8.5l-4 4l-4-4"/>`
	chevronDownCircleInnerSVG   = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m5.466 7.466l3 3.068l3-3.068"/></g>`
	chevronDownDoubleInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.5 6.5l-4 4l-4-4m8 4l-4 4l-4-4"/>`
	chevronLeftInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.5 14.5l-4-4l4-4"/>`
	chevronLeftCircleInnerSVG   = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m9.55 11.4l-3-2.9l3-3"/></g>`
	chevronLeftDoubleInnerSVG   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5 14.5l-4-4l4-4m-4 8l-4-4l4-4"/>`
	chevronOpenInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 8.5l3-3l3 3m-6 5l3 3l3-3"/>`
	chevronRightInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5 14.5l4-4l-4-4"/>`
	chevronRightCircleInnerSVG  = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m7.5 11.5l3-3l-3.068-3"/></g>`
	chevronRightDoubleInnerSVG  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 14.5l4-4l-4-4m4 8l4-4l-4-4"/>`
	chevronUpInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 12.5l4-4l4 4"/>`
	chevronUpCircleInnerSVG     = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m11.5 9.5l-3-3l-3 3"/></g>`
	chevronUpDoubleInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 14.5l4-4l4 4m-8-4l4-4l4 4"/>`
	circleInnerSVG              = `<circle cx="10.5" cy="10.5" r="8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`
	circleMenuInnerSVG          = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="currentColor" d="M8.5 9.5c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm-4 0c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm8 0c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1z"/></g>`
	circleSplitInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m8.527.5l-.027 16"/></g>`
	clipboardInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 4.5c-.441 0-1.039-.004-1.998-.005a1 1 0 0 0-.995.881l-.007.12v10.997a1 1 0 0 0 1 1l10 .006a1 1 0 0 0 .994-.882l.006-.117v-11a1 1 0 0 0-1-1h-2"/><path d="M8.5 3.5h4a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2zm-2 5h5m-5 2h7m-7 2h3m-3 2h6"/></g>`
	clipboardAddInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 4.5h-2a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1h-2"/><path d="M8.5 3.5h4a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2zm2 5v6.056m3-3.056h-6"/></g>`
	clipboardCheckInnerSVG      = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 4.5h-2a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1h-2"/><path d="M8.5 3.5h4a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2zm-1 8l2 2l5-5"/></g>`
	clipboardCopyInnerSVG       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.5 14.5l-3-3l3-3m-3 3h11"/><path d="M16.5 9.5V5.495a1 1 0 0 0-.883-.993l-.12-.007L13.5 4.5m-6 0l-1.998-.005a1 1 0 0 0-.995.881l-.007.12v10.997a1 1 0 0 0 1 1l10 .006a1 1 0 0 0 .994-.882l.006-.117V14"/><path d="M8.5 3.5h4a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2z"/></g>`
	clipboardCrossInnerSVG      = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 4.5h-2a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1h-2"/><path d="M8.5 3.5h4a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2zm-1 5l6 6m0-6l-6 6"/></g>`
	clipboardNotesInnerSVG      = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 4.5h-2a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1h-2"/><path d="M8.5 3.5h4a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2zm1 5h5m-5 3h5m-5 3h5m-8-6h1m-1 3h1m-1 3h1"/></g>`
	clipboardRemoveInnerSVG     = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.5 4.5h-2a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1h-2"/><path d="M8.5 3.5h4a1 1 0 1 1 0 2h-4a1 1 0 1 1 0-2zm5 8h-6"/></g>`
	clockInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="matrix(-1 0 0 1 19 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="M8.5 5.5v4H5"/></g>`
	closeInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 7.5l6 6m0-6l-6 6"/>`
	cloudInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 5.5a5 5 0 0 1 4.802 6.399A2 2 0 1 1 16.5 15.5h-11a3 3 0 1 1 .1-5.998A5.002 5.002 0 0 1 10.5 5.5z"/>`
	cloudDisconnectInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.739 5.819a5 5 0 0 1 6.563 6.08a2 2 0 1 1 2.234 3.312m-2.104.289H5.5a3 3 0 1 1 .1-5.998a4.988 4.988 0 0 1 1.286-2.457M4 4l13 13.071"/>`
	cloudDownloadInnerSVG       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 5.5a5 5 0 0 1 4.802 6.399A2 2 0 1 1 16.5 15.5h-11a3 3 0 1 1 .1-5.998A5.002 5.002 0 0 1 10.5 5.5z"/><path d="m12.5 11.5l-2 2l-2-2m2-4v6"/></g>`
	cloudDownloadAltInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 15.5h-4c-1.152-.419-2-1.703-2-3a3 3 0 0 1 3.1-2.998a5.002 5.002 0 1 1 9.702 2.397A2 2 0 1 1 16.5 15.5h-4m-4 2l2 2l2-2m-2-8v10"/>`
	cloudUploadInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 5.5a5 5 0 0 1 4.802 6.399A2 2 0 1 1 16.5 15.5h-11a3 3 0 1 1 .1-5.998A5.002 5.002 0 0 1 10.5 5.5z"/><path d="m8.5 9.5l2-2l2 2m-2-2v6"/></g>`
	cloudUploadAltInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.578 6.56A4.991 4.991 0 0 1 15.5 10.5c0 .485-.07.955-.198 1.399A2 2 0 1 1 16.5 15.5h-11a3 3 0 1 1 .1-5.998A5 5 0 0 1 7.447 6.54M8.5 4.5l2-2l2 2m-2-2v9"/>`
	codeInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.5 3.5l-4 14m-2-5l-4-4l4-4m8 12l4-4l-4-4"/>`
	coffeeInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 8.5h6a2 2 0 0 1 2 2V13a4.5 4.5 0 0 1-4.5 4.5H9A4.5 4.5 0 0 1 4.5 13v-2.5a2 2 0 0 1 2-2zm8 2h1a2 2 0 1 1 0 4h-1m-6-8v-4m2 4v-2"/>`
	coinInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 9.5c0-1.3 3.134-3 7-3s7 1.7 7 3v3c0 1.3-3.134 3-7 3s-7-1.7-7-3v-3z"/><path d="M10.5 12.484c3.866 0 7-1.606 7-2.986c0-1.381-3.134-2.998-7-2.998s-7 1.617-7 2.998c0 1.38 3.134 2.986 7 2.986z"/></g>`
	coinsInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 11.5v3c0 1.3-3.134 3-7 3s-7-1.7-7-3V12"/><path d="M4.794 12.259c.865 1.148 3.54 2.225 6.706 2.225c3.866 0 7-1.606 7-2.986c0-.775-.987-1.624-2.536-2.22"/><path d="M15.5 6.5v3c0 1.3-3.134 3-7 3s-7-1.7-7-3v-3"/><path d="M8.5 9.484c3.866 0 7-1.606 7-2.986c0-1.381-3.134-2.998-7-2.998s-7 1.617-7 2.998c0 1.38 3.134 2.986 7 2.986z"/></g>`
	compassInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m10.5 9.5l-4 3v-5l4-3z"/></g>`
	componentAddInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 5.5h-4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-4m0-8v6m3-3h-6"/>`
	contactsInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 3)"><circle cx="7.5" cy="5.5" r="2"/><path d="M.5 3.5h1v-1a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-1h-1m0-4h1m-1-2h1m-1 4h1"/><path d="M10.5 10.5v-1a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1z"/></g>`
	contractInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 16.5v-4.978l-5-.022m14-9l-7 7m5 0l-5 .023V4.5m-2 7l-7 7"/>`
	createInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10 4.5H5.5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V11"/><path d="M17.5 3.467a1.462 1.462 0 0 1-.017 2.05L10.5 12.5l-3 1l1-3l6.987-7.046a1.409 1.409 0 0 1 1.885-.104zm-2 2.033l.953 1"/></g>`
	creditCardInnerSVG          = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 5.5h12a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M2 9h17v2H2z"/></g>`
	cropInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 7.5h4v4m0 2v3m-6-9H4"/><path d="M7.5 4.5v9h9"/></g>`
	crossInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.5 15.5l-10-10zm0-10l-10 10"/>`
	crossCircleInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="m5.5 5.5l6 6m0-6l-6 6"/></g>`
	crosshairInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 16.5c3.329 0 6-2.645 6-5.973S13.829 4.5 10.5 4.5s-6 2.698-6 6.027s2.671 5.973 6 5.973zm-6-6h2m8 0h2m-6-6v2m0 8v2"/>`
	cubeInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.492 4.067l5 2.857A2 2 0 0 1 17.5 8.661v4.678a2 2 0 0 1-1.008 1.737l-5 2.857a2 2 0 0 1-1.984 0l-5-2.857A2 2 0 0 1 3.5 13.339V8.661a2 2 0 0 1 1.008-1.737l5-2.857a2 2 0 0 1 1.984 0zM10.5 11.5V18"/><path d="m4 8l5.552 2.99a2 2 0 0 0 1.896 0L17 8"/></g>`
	cubesInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.5 15.429l3.548 1.837a1 1 0 0 0 .907.006l2.992-1.496a1 1 0 0 0 .553-.894v-2.764a1 1 0 0 0-.553-.894L14.5 9.5l-3.46 1.792a1 1 0 0 0-.54.888v3.249z"/><path d="m3.04 15.708l3.008 1.558a1 1 0 0 0 .907.006L10.5 15.5v-3.382a1 1 0 0 0-.553-.894L6.5 9.5l-3.46 1.792a1 1 0 0 0-.54.888v2.64a1 1 0 0 0 .54.888zM6.5 9.429l3.548 1.837a1 1 0 0 0 .907.006L14.5 9.5V6.118a1 1 0 0 0-.553-.894l-2.992-1.496a1 1 0 0 0-.907.006L7.04 5.292a1 1 0 0 0-.54.888v3.249z"/><path d="m6.846 5.673l3.207 1.603a1 1 0 0 0 .894 0L14.12 5.69h0m-3.274 5.983l3.207 1.603a1 1 0 0 0 .894 0l3.172-1.586h0m-15.273-.017l3.207 1.603a1 1 0 0 0 .894 0l3.172-1.586h0M10.5 7.5v4m4 2V17m-8-3.5V17"/></g>`
	cylinderInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 5.353c0-1.3 2-2.853 5-2.853s5 1.553 5 2.853v10.294c0 1.3-2 2.853-5 2.853s-5-1.553-5-2.853V5.353z"/><path d="M5.5 5.5c0 1.38 2 3 5 3s5-1.62 5-3"/></g>`
	databaseInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.5 5.206c0-1.3 2.5-2.741 6-2.706s6 1.553 6 2.853v10.294c0 1.3-2.5 2.853-6 2.853s-6-1.7-6-3V5.206z"/><path d="M4.5 5.5c0 1.38 2 3 6 3s6-1.637 6-3.018M4.5 10.5c0 1.38 2 3 6 3s6-1.637 6-3.018"/></g>`
	diamondInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m15.5 4l3 4l-8 10l-8-10l3.009-4zm-13 4h16m-11 0l3 10m3-10l-3 10"/><path d="M5.509 4L7.5 8l3-4l3 4l2-4"/></g>`
	directionsInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5h11l2 2l-2 2h-11a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1zm12 7h-11l-2 2l2 2h11a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1zm-6-3v3m0 4v3"/>`
	discInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="10.5" cy="10.5" r="8"/><circle cx="10.5" cy="10.5" r="2"/></g>`
	displayInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 3.5h11a2 2 0 0 1 2 2v6.049a2 2 0 0 1-1.85 1.994l-.158.006l-11-.042a2 2 0 0 1-1.992-2V5.5a2 2 0 0 1 2-2zm.464 12H15.5m-8 2h6"/>`
	displayAltInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 7.5h11a2 2 0 0 1 2 2v6.049a2 2 0 0 1-1.85 1.994l-.158.006l-11-.042a2 2 0 0 1-1.992-2V9.5a2 2 0 0 1 2-2zm.464-2H15.5m-8-2h6"/>`
	documentInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M16.5 15.5v-7l-5-5h-5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zm-10-5h5m-5 2h7m-7 2h3"/><path d="M11.5 3.5v3a2 2 0 0 0 2 2h3"/></g>`
	documentJustifiedInnerSVG   = `<defs><path id="systemUiconsDocumentJustified0" d="M16.5 15.5v-10a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zm-9-8h6m-6 3h6m-6 3h6"/></defs><g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><use href="#systemUiconsDocumentJustified0"/><use href="#systemUiconsDocumentJustified0"/></g>`
	documentListInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 15.5v-10a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zm-7-8h5m-8 0h1m2 3h5m-8 0h1m2 3h5m-8 0h1"/>`
	documentStackInnerSVG       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 14.5v-10a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="m5.305 4.935l-2.004.73a2 2 0 0 0-1.195 2.563l3.42 9.397A2 2 0 0 0 8.09 18.82l5.568-2.198M8.5 7.5h5m-5 2h6m-6 2h3"/></g>`
	documentWordsInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 15.5v-10a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zm-9-7h5m-5 2h6m-6 2h3"/>`
	doorInnerSVG                = `<g fill="none" fill-rule="evenodd" transform="translate(4 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h8a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.202 14.5l-3.645-1.948A2 2 0 0 1 5.5 10.788V4.212a2 2 0 0 1 1.057-1.764L10.202.5"/><circle cx="7.5" cy="7.5" r="1" fill="currentColor"/></g>`
	doorAltInnerSVG             = `<g fill="none" fill-rule="evenodd" transform="translate(4 1)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 2.5h2v14h-2a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zM7.202.513l4 1.5A2 2 0 0 1 12.5 3.886v11.228a2 2 0 0 1-1.298 1.873l-4 1.5A2 2 0 0 1 4.5 16.614V2.386A2 2 0 0 1 7.202.513z"/><circle cx="6.5" cy="9.5" r="1" fill="currentColor"/></g>`
	downloadInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 10.5l4 4.232l4-4.191m-4-7.041v11m-6 3h12"/>`
	downloadAltInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m14.5 11.5l-3.978 4l-4.022-4m4.022-7.979V15.5"/><path d="M3.5 12v4.5a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V12"/></g>`
	downwardInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.23 10.23c3.264-3.79 6.687-5.033 10.27-3.73c-3.552.646-6.009 2.855-7.371 6.63L12.5 15.5h-8v-8z"/>`
	dragInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 7.5h12m-12.002 3h11.997M4.5 13.5h11.995"/>`
	dragCircleInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="10.5" cy="10.5" r="8"/><path d="M6.5 8.5h8m-8 2h8m-8 2h8"/></g>`
	dragVerticalInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M7 5h2v2H7zm5 0h2v2h-2zM7 9h2v2H7zm5 0h2v2h-2zm-5 4h2v2H7zm5 0h2v2h-2z"/>`
	duplicateInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 12.5v-8a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="M6.5 6.503h-2a2 2 0 0 0-2 2V16.5a2 2 0 0 0 2 2h.003l8-.014a2 2 0 0 0 1.997-2v-1.983m-2-9.003v6m3-3h-6"/></g>`
	duplicateAltInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 16.5v-8a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="M14.5 14.5h2a2 2 0 0 0 2-2V4.503a2 2 0 0 0-2-2h-.003l-8 .014a2 2 0 0 0-1.997 2V6.5m2 3v6m3-3h-6"/></g>`
	enterInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m9.5 13.5l3-3l-3-3m3 3h-9"/><path d="M4.5 8.5V5.492a2 2 0 0 1 1.992-2l7.952-.032a2 2 0 0 1 2.008 1.993l.04 10.029A2 2 0 0 1 14.5 17.49h-8a2 2 0 0 1-2-2V12.5"/></g>`
	enterAltInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 13.535l-3-3.035l3-3m7 3h-10"/><path d="M16.5 8.5V5.54a2 2 0 0 0-1.992-2l-8-.032A2 2 0 0 0 4.5 5.5v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-3"/></g>`
	episodesInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 16.5v-6a2 2 0 0 0-2-2h-9a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2z"/><path d="M16.5 16.5V9.505a3 3 0 0 0-3-3h-.005L4.5 6.521"/><path d="M18.5 14.5V8.507a4 4 0 0 0-4-4h-.007L6.5 4.52"/></g>`
	exitLeftInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.405 13.5l-2.905-3l2.905-3m-2.905 3h9m-6-7l8 .002c1.104.001 2 .896 2 2v9.995a2 2 0 0 1-2 2l-8 .003"/>`
	exitRightInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.595 13.5l2.905-3l-2.905-3m2.905 3h-9m6-7l-8 .002c-1.104.001-2 .896-2 2v9.995a2 2 0 0 0 2 2h8.095"/>`
	expandInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 7.5v-5h-5m5 0l-6 5.929M7.5 18.5l-5 .023V13.5m6-1l-6 6"/>`
	expandHeightInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.51 1.51H2.49m16.02 18H2.49m4.006-5.992l4 4l4.015-4m-8.015-6l4-4l4.015 4M10.51 17.51v-14"/>`
	expandWidthInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 2.5v16.021M19.5 2.5v16.021m-5.993-4.006l4-4l-4-4.015m-6 8.015l-4-4l4-4.015m9.993 4h-14"/>`
	externalInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 8.5v-5h-5m5 0l-7 7m-1-7h-5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2v-4"/>`
	eyeInnerSVG                 = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 16c3.13 0 5.963-1.833 8.5-5.5C16.463 6.833 13.63 5 10.5 5S4.537 6.833 2 10.5c2.537 3.667 5.37 5.5 8.5 5.5z"/><path d="M10.5 7c.185 0 .366.014.543.042a2.5 2.5 0 0 0 2.915 2.916A3.5 3.5 0 1 1 10.5 7z"/></g>`
	eyeClosedInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 10.5c2.537 3.667 5.37 5.5 8.5 5.5s5.963-1.833 8.5-5.5M4.5 13.423l-2 2.077m14-2.077l2 2.077m-6 .5l1 2.5m-5-2.5l-1 2.5"/>`
	eyeNoInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.211 6.26C4.727 7.173 3.323 8.587 2 10.5c2.537 3.667 5.37 5.5 8.5 5.5c1.423 0 2.785-.38 4.085-1.137m1.588-1.14c.98-.843 1.922-1.916 2.827-3.223C16.463 6.833 13.63 5 10.5 5c-.83 0-1.64.13-2.429.387M4 4l13 13.071"/>`
	faceDelightedInnerSVG       = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="6" cy="6" r="1" fill="currentColor"/><circle cx="11" cy="6" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 10.5c.666 2 2 3 4 3s3.334-1 4-3z"/></g>`
	faceHappyInnerSVG           = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="6" cy="6" r="1" fill="currentColor"/><circle cx="11" cy="6" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 9.5c.603 1.333 1.603 2 3 2s2.397-.667 3-2"/></g>`
	faceNeutralInnerSVG         = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="6" cy="6" r="1" fill="currentColor"/><circle cx="11" cy="6" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 10.5h6"/></g>`
	faceSadInnerSVG             = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 16.5a8 8 0 1 0 0-16a8 8 0 0 0 0 16z"/><circle cx="6" cy="6" r="1" fill="currentColor"/><circle cx="11" cy="6" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 11.5c.603-1.333 1.603-2 3-2s2.397.667 3 2"/></g>`
	fileDownloadInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 3.5H6.498a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2l.002-8l-4-4"/><path d="m13.5 10.586l-3 2.914l-3-2.914m3-8.086v11"/></g>`
	fileUploadInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 17.5h2a2 2 0 0 0 2-2v-8l-4-4h-6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h2"/><path d="m7.5 10.5l3-3l3 3m-3-3v11"/></g>`
	filesHistoryInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 14.5v-7l-5-5h-5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="M11.5 6.5v4h3m-9-6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2"/></g>`
	filesMultiInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 14.5v-7l-5-5h-5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="M3.5 4.5v10a4 4 0 0 0 4 4h8m-3-16v3a2 2 0 0 0 2 2h3"/></g>`
	filesStackInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 14.5v-7l-5-5h-5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="M12.5 2.5v3a2 2 0 0 0 2 2h3m-12-3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2"/></g>`
	filmInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 4.5h10a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2zm1 0v12m8-12v12m0-9h3m-3 6h3m-14-6h3m-3 3h14m-14 3h3"/>`
	filterInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 7.5h12m-10 3h8m-6 3h4"/>`
	filterCircleInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="10.5" cy="10.5" r="8"/><path d="M6.5 8.5h8m-6 2h4m-3 2h2"/></g>`
	filterSingleInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 8V2.5m0 16V13"/><circle cx="10.5" cy="10.5" r="2.5"/></g>`
	filteringInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 4a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0V5a1 1 0 0 1 1-1zm12 2h-11m-2 0h-3m4 8a1 1 0 0 1 1 1v2a1 1 0 0 1-2 0v-2a1 1 0 0 1 1-1zm12 2h-11m-2 0h-3m12-7a1 1 0 0 1 1 1v2a1 1 0 0 1-2 0v-2a1 1 0 0 1 1-1zm-1 2h-11m16 0h-3"/>`
	fingerprintInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M16.5 12.53c0 2.745-1.255 5.97-6 5.97a8.184 8.184 0 0 1-2.015-.232m-1.761-.781C5.034 16.317 4.5 14.32 4.5 12.53V8.5c0-1.566.655-2.98 1.705-3.981m1.672-1.354A5.475 5.475 0 0 1 10.5 2.5c1.753 0 3.493.723 4.5 2m1.206 1.22c.19.559.294 1.157.294 1.78v3"/><path d="M10.5 16.5c-1.333-.667-2-1.657-2-2.97V8.5a2 2 0 1 1 4 0v4.03c0 .667.333 1 1 1s1-.333 1-1V8.066a3 3 0 0 0-.304-1.316c-.732-1.5-1.964-2.25-3.696-2.25c-1.732 0-2.964.75-3.696 2.25A3 3 0 0 0 6.5 8.066V13.5c0 1 .167 1.667.5 2"/><path d="M10.5 8.5v4.03c0 1.98 1 2.97 3 2.97"/></g>`
	flagInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 17.5v-11m0 0c.667-1.333 1.667-2 3-2c2 0 2 2 4 2c1.333 0 2.333-.333 3-1v6c-.667.667-1.667 1-3 1c-2 0-2-2-4-2c-1.333 0-2.333.667-3 2z"/>`
	flameInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 19c3.038 0 5.5-2.429 5.5-6.714C16 9.429 14.167 6.333 10.5 3C6.833 6.333 5 9.429 5 12.286C5 16.57 7.462 19 10.5 19Z"/><path d="M10.5 19c1.519 0 2.75-1.214 2.75-3.357c0-1.429-.917-2.976-2.75-4.643c-1.833 1.667-2.75 3.214-2.75 4.643C7.75 17.786 8.981 19 10.5 19Z"/></g>`
	flameAltInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 3c3.667 3.333 5.5 6.429 5.5 9.286c0 3.078-1.27 5.198-3.111 6.148c.23-.491.361-1.092.361-1.791c0-1.429-.917-2.976-2.75-4.643c-1.833 1.667-2.75 3.214-2.75 4.643c0 .7.131 1.3.36 1.791c-1.84-.95-3.11-3.07-3.11-6.148C5 9.429 6.833 6.333 10.5 3z"/>`
	flipViewInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.69 16.743l.5-2a1 1 0 0 0-.97-1.243H3.78a1 1 0 0 0-.97 1.243l.5 2a1 1 0 0 0 .97.757h12.44a1 1 0 0 0 .97-.757zM17.5 11.5l.56-1.682a1 1 0 0 0-.95-1.316l-13.22.017a1 1 0 0 0-.947 1.318L3.5 11.5m14-5l.306-1.836A1 1 0 0 0 16.82 3.5H4.18a1 1 0 0 0-.986 1.164L3.5 6.5"/>`
	floppyInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 4.5h7l3 3v7a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2z"/><path d="M8.5 12.5h4a1 1 0 0 1 1 1v3h-6v-3a1 1 0 0 1 1-1zm-1-5h2v2h-2z"/></g>`
	folderAddInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 5.5v9a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V8.497a1.999 1.999 0 0 0-2-1.999l-5 .002l-2-2h-4a1 1 0 0 0-1 1zm5 6h4m-2 2.056V9.5"/>`
	folderClosedInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 5.5v9a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V8.497a1.999 1.999 0 0 0-1.85-1.994l-.15-.005l-5 .002l-2-2h-4a1 1 0 0 0-1 1zm0 1h7"/>`
	folderMinusInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 5.5v9a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V8.497a1.999 1.999 0 0 0-1.85-1.994l-.15-.005l-5 .002l-2-2h-4a1 1 0 0 0-1 1zm5 6h4"/>`
	folderOpenInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 8.5a2 2 0 0 0-2-2.003l-5 .003l-2-2h-4a1 1 0 0 0-1 1v3"/><path d="m2.81 9.742l1.311 5.243a2 2 0 0 0 1.94 1.515h8.877a2 2 0 0 0 1.94-1.515L18.19 9.74a1 1 0 0 0-.97-1.243L3.781 8.5a1 1 0 0 0-.97 1.243z"/></g>`
	forkGitInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m14.5 3.5l3 3l-3 3"/><path d="M17.5 6.5h-5l-4 5.086m6 .914l3 3l-3 3"/><path d="M17.5 15.5h-5l-4-4h-6"/></g>`
	forwardInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 7.5c-5.185-.47-8.52 1.53-10 6c2.825-3.14 6.341-3.718 10-2v3l5-5l-5-5z"/>`
	frameInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 4v14m8-14v14M3.5 7h14m-14 8h14"/>`
	fullscreenInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 7.5V2.522l-5.5.014m5.5-.014l-6 5.907m.5 10.092l5.5.002l-.013-5.5m.013 5.406l-6-5.907M2.5 7.5v-5H8m.5 5.929l-6-5.907M8 18.516l-5.5.007V13.5m6-1l-6 6"/>`
	funnelInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5h12l-4 7v3l-3 3v-6z"/>`
	gaugeInnerSVG               = `<g fill="none" fill-rule="evenodd" transform="translate(2 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14 14c1.448-1.448 2.5-3.29 2.5-5.5a8 8 0 1 0-16 0c0 2.21 1.052 4.052 2.5 5.5m5.5-5.5l-4-4"/><circle cx="8.5" cy="8.5" r="1.5" fill="currentColor"/></g>`
	giftInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.5 10.5h12v5a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2zm6-3v10"/><path d="M7.5 7.5h3v-2c0-1.5-2.219-1.781-3-1s-.781 2.219 0 3zm6 0h-3v-2c0-1.5 2.219-1.781 3-1c.781.781.781 2.219 0 3zm-9 0h12a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1z"/></g>`
	globeInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10 19c4.438 0 8-3.526 8-7.964C18 6.598 14.438 3 10 3c-4.438 0-8 3.598-8 8.036S5.562 19 10 19zM3 8h14M3 14h14"/><path d="M10 19c2.219 0 4-3.526 4-7.964C14 6.598 12.219 3 10 3c-2.219 0-4 3.598-4 8.036S7.781 19 10 19z"/></g>`
	gpsInnerSVG                 = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 14.5c3.329 0 6-2.645 6-5.973c0-3.329-2.671-6.027-6-6.027s-6 2.698-6 6.027c0 3.328 2.671 5.973 6 5.973z"/><circle cx="8.5" cy="8.5" r="3.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5h2m12 0h2m-8-8v2m0 12v2"/></g>`
	grabInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.35 9.505L7.5 9.5v-1a1 1 0 1 1 2 0v-1a1 1 0 1 1 2 0v1a1 1 0 1 1 2 0v1a1 1 0 1 1 2 0v4a5 5 0 0 1-5 5H10A4.5 4.5 0 0 1 5.5 14v-2.5a2 2 0 0 1 1.85-1.995zM7.5 8.5v3m2-4v2m2-2v2m2-1v2"/>`
	graphBarInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 3.5v12a2 2 0 0 0 2 2H17m-10.5-6v3m4-6v6m4-9v9"/>`
	graphBoxInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm2 10.935V6.5m3 7.985V10.5m3 4v-6"/>`
	graphIncreaseInnerSVG       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.5 3.5v11a2 2 0 0 0 2 2h11"/><path d="m6.5 12.5l3-3l2 2l5-5"/><path d="M16.5 9.5v-3h-3"/></g>`
	gridInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 3.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm-8 0h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm8 8h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm-8 0h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z"/>`
	gridCirclesInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7.5" cy="7.5" r="2"/><circle cx="13.5" cy="7.5" r="2"/><circle cx="7.5" cy="13.5" r="2"/><circle cx="13.5" cy="13.5" r="2"/></g>`
	gridCirclesAddInnerSVG      = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="13.5" cy="7.5" r="2"/><circle cx="7.5" cy="7.5" r="2"/><circle cx="7.5" cy="13.5" r="2"/><path d="M13.5 11.5v4m2-2h-4"/></g>`
	gridSmallInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M10 9h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1zm0-4h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1zm4 4h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1zm0-4h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1zm0 8h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1zm-4 0h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1zM6 9h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1zm0-4h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1zm0 8h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1z"/>`
	gridSquaresInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.499 5.5l-2 .003a1 1 0 0 0-1 1V8.5a1 1 0 0 0 .884.993l.118.007l2-.003a1 1 0 0 0 .999-1V6.501a1 1 0 0 0-.884-.994zm6 0l-2 .003a1 1 0 0 0-1 1V8.5a1 1 0 0 0 .884.993l.118.007l2-.003a1 1 0 0 0 .999-1V6.501a1 1 0 0 0-.884-.994zm-6 6l-2 .003a1 1 0 0 0-1 1V14.5a1 1 0 0 0 .884.993l.118.007l2-.003a1 1 0 0 0 .999-1v-1.996a1 1 0 0 0-.884-.994zm6 0l-2 .003a1 1 0 0 0-1 1V14.5a1 1 0 0 0 .884.993l.118.007l2-.003a1 1 0 0 0 .999-1v-1.996a1 1 0 0 0-.884-.994z"/>`
	gridSquaresAddInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.499 5.5l-2 .003a1 1 0 0 0-1 1V8.5a1 1 0 0 0 .884.993l.118.007l2-.003a1 1 0 0 0 .999-1V6.501a1 1 0 0 0-.884-.994zm6 0l-2 .003a1 1 0 0 0-1 1V8.5a1 1 0 0 0 .884.993l.118.007l2-.003a1 1 0 0 0 .999-1V6.501a1 1 0 0 0-.884-.994zm-6 6l-2 .003a1 1 0 0 0-1 1V14.5a1 1 0 0 0 .884.993l.118.007l2-.003a1 1 0 0 0 .999-1v-1.996a1 1 0 0 0-.884-.994zm5.001 0v4m2-2h-4"/>`
	handInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.25 9.25a1.532 1.532 0 0 1 1.723.934L7.5 11.5v-6a1 1 0 1 1 2 0V4a1 1 0 1 1 2 0v1a1 1 0 1 1 2 0v1.5a1 1 0 1 1 2 0v8a4 4 0 0 1-4 4c-2.35 0-4.4-1.6-4.97-3.88l-.03-.12l-1.93-3.86a.974.974 0 0 1 .68-1.39zM9.5 4.5v6m2-6v5m2-3v4"/>`
	harddriveInnerSVG           = `<g fill="none" fill-rule="evenodd" transform="rotate(-90 10 8)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h4l5.788 2.48A2 2 0 0 1 13.5 4.82v7.362a2 2 0 0 1-1.212 1.838L6.5 16.5h-4a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2z"/><circle cx="4" cy="3" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 1v15"/></g>`
	hashInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5 5.5l-2 10m-2-10l-2 10m-1-7h9m-10 4h9"/>`
	heartInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 6.5c.5-2.5 4.343-2.657 6-1c1.603 1.603 1.5 4.334 0 6l-6 6l-6-6a4.243 4.243 0 0 1 0-6c1.55-1.55 5.5-1.5 6 1z"/>`
	heartRateInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 10.5h2l2.5-6l2 12l3-9l2.095 6l1.405-3h2"/>`
	heartRemoveInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.929 14.07l-3.43 3.43l-6-6a4.243 4.243 0 0 1 0-6a2.96 2.96 0 0 1 .567-.438m2.453-.605c1.388.034 2.706.668 2.98 2.043c.5-2.5 4.344-2.657 6-1c1.604 1.603 1.5 4.334 0 6l-.937.937M4 4l13 13.071"/>`
	heightInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.519 5.497l-4-4l-4.015 4m8.015 10l-4 4l-4.015-4m4-13.993v18"/>`
	hierarchyInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 2.5h6v5h-6zm5 11h6v5h-6zm-10 0h6v5h-6zm2.998 0v-3h10v3m-4.998-3v-3"/>`
	homeInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.5 10.5l9-9l9 9"/><path d="M3.5 8.5v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-7"/></g>`
	homeAltInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.5 10.5l9-9l9 9m-16 1v4a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-4"/>`
	homeCheckInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.5 10.5l9-9l9 9"/><path d="M3.5 8.5v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-7"/><path d="m7.5 11.5l2 2l4-4"/></g>`
	homeDoorInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m1.5 10.5l9-9l9 9"/><path d="M3.5 8.5v8a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-4a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v4a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-8"/></g>`
	importInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 3.5h-4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-10"/><path d="m13.5 10.5l-3 3l-3-3"/><path d="M17.5 3.5h-4a3 3 0 0 0-3 3v7"/></g>`
	inboxInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.16 4.5h8.68a1 1 0 0 1 .92.606L18.5 11.5v4a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-4l2.74-6.394a1 1 0 0 1 .92-.606z"/><path d="M2.5 11.5h4a1 1 0 0 1 1 1v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h4"/></g>`
	inboxAltInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 4.5h2.34a1 1 0 0 1 .92.606L18.5 11.5v4a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-4l2.74-6.394a1 1 0 0 1 .92-.606H8.5"/><path d="m13.5 7.586l-3 2.914l-3-2.914m3-6.086v9m-8 1h4a1 1 0 0 1 1 1v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h4"/></g>`
	infoCircleInnerSVG          = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="8.5" cy="8.5" r="8"/><path d="M8.5 12.5v-4h-1m0 4h2"/></g><circle cx="8.5" cy="5.5" r="1" fill="currentColor"/></g>`
	iphoneLandscapeInnerSVG     = `<g fill="none" fill-rule="evenodd" transform="translate(3 5)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2z"/><circle cx="11.5" cy="5.5" r="1" fill="currentColor"/></g>`
	iphonePortraitInnerSVG      = `<g fill="none" fill-rule="evenodd" transform="translate(5 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h6a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><circle cx="5.5" cy="11.5" r="1" fill="currentColor"/></g>`
	jumpBackwardInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 14.5v-2a3 3 0 0 0-3-3h-8m0 3l-3.001-3l3.001-3"/><path d="m9.5 12.5l-3.001-3l3.001-3"/></g>`
	jumpForwardInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 14.5v-2a3 3 0 0 1 3-3h8"/><path d="m11.499 12.5l3.001-3l-3.001-3"/><path d="m14.499 12.5l3.001-3l-3.001-3"/></g>`
	jumpLeftInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.5 14.5v-2a3 3 0 0 0-3-3h-8"/><path d="m7.5 12.5l-3.001-3l3.001-3"/></g>`
	jumpRightInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 14.5v-2a3 3 0 0 1 3-3h8"/><path d="m13.499 12.5l3.001-3l-3.001-3"/></g>`
	keyboardInnerSVG            = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.5 13.5v-6a2 2 0 0 0-2-2h-14a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2z"/><g fill="currentColor" transform="translate(1 5)"><circle cx="3.5" cy="2.5" r="1"/><circle cx="6.5" cy="2.5" r="1"/><circle cx="9.5" cy="2.5" r="1"/><circle cx="12.5" cy="2.5" r="1"/><circle cx="15.5" cy="2.5" r="1"/><circle cx="3.5" cy="4.5" r="1"/><circle cx="6.5" cy="4.5" r="1"/><circle cx="9.5" cy="4.5" r="1"/><circle cx="12.5" cy="4.5" r="1"/><circle cx="15.5" cy="4.5" r="1"/><circle cx="3.5" cy="6.5" r="1"/><circle cx="6.5" cy="6.5" r="1"/><circle cx="9.5" cy="6.5" r="1"/><circle cx="12.5" cy="6.5" r="1"/><circle cx="15.5" cy="6.5" r="1"/><circle cx="3.5" cy="8.5" r="1"/><circle cx="15.5" cy="8.5" r="1"/></g><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 13.5h6"/></g>`
	laptopInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.485h11a1 1 0 0 1 1 1V13.5h-13V5.485a1 1 0 0 1 1-1zM3.118 17.5h13.764a1 1 0 0 0 .894-1.447L16.5 13.5h-13l-1.276 2.553a1 1 0 0 0 .894 1.447z"/>`
	lightbulbInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 18.5h4M10.5 5a4.5 4.5 0 0 1 2.001 8.532l-.001.968a2 2 0 1 1-4 0v-.968A4.5 4.5 0 0 1 10.5 5z"/>`
	lightbulbOnInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 18.5h4M10.5 5a4.5 4.5 0 0 1 2.001 8.532l-.001.968a2 2 0 1 1-4 0v-.968A4.5 4.5 0 0 1 10.5 5zm0-2.5v-1m5 3l1-1m-11 1l-1-1m11 10l1 1m-11-1l-1 1m-1-6h-1m16 0h-1"/>`
	lightningInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 9.5h4l-6 9v-6.997l-4-.003l6-9z"/>`
	lightningAltInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8.5h5l-6 8.997V11.5h-4l2-9h5z"/>`
	lineweightInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor"><rect width="14" height="2" x="3.5" y="6.5" fill="currentColor" rx="1"/><path fill="currentColor" d="M3.5 11.5h14v1h-14z"/><path stroke-linecap="round" stroke-linejoin="round" d="M3.5 15.5h13.981"/></g>`
	linkInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5 7.5l1-1a2.828 2.828 0 1 1 4 4l-1 1m-3 3l-1 1a2.828 2.828 0 1 1-4-4l1-1m1 3l5-5"/>`
	linkAltInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.388 11.69A3.043 3.043 0 0 1 6.543 6.5h4.914A3.043 3.043 0 0 1 14.5 9.543c0 1.68-1.362 2.957-3.043 2.957H9"/><path d="M16.612 10.31a3.043 3.043 0 0 1-2.155 5.19H9.543A3.043 3.043 0 0 1 6.5 12.457c0-1.68 1.362-2.957 3.043-2.957H12"/></g>`
	linkBrokenInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5 7.328l1-1a2.828 2.828 0 0 1 4 4l-1 1M10.328 14.5l-1 1a2.828 2.828 0 1 1-4-4l1-1m1.172-5v-3m-5 5h3m8 11v-3m2-2h3"/>`
	linkHorizontalInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 14.5h-1a4 4 0 1 1 0-8h1m4 0h1a4 4 0 1 1 0 8h-1m1-4h-6"/>`
	linkVerticalInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 12.5v1a4 4 0 1 1-8 0v-1m0-4v-1a4 4 0 1 1 8 0v1m-4-1v6"/>`
	listInnerSVG                = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 10.5h-7m7 4h-7m7-8h-7"/><path fill="currentColor" d="M5.499 7.5c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm0 4c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm0 4c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1z"/></g>`
	listAddInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h12m-12 4h12m-12 4h7m2 0h4zm2 2v-4z"/>`
	listNumberedInnerSVG        = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 10.5h-7m7 4h-7m7-8h-7"/><path fill="currentColor" d="M5.88 8V5.828h-.037l-.68.459V5.67l.717-.488h.717V8zm-.98 2.068c0-.572.45-.963 1.109-.963c.652 0 1.04.354 1.04.836c0 .334-.148.555-.597.961l-.555.502v.037h1.186V12H4.94v-.479l1.008-.912c.348-.318.406-.44.406-.605c0-.195-.136-.358-.382-.358c-.262 0-.416.178-.416.422zm.712 4.73v-.484h.362c.238 0 .392-.138.392-.341c0-.192-.146-.332-.388-.332c-.254 0-.409.134-.42.363h-.653c.01-.541.438-.899 1.108-.899c.66 0 1.021.346 1.02.766c0 .34-.22.565-.528.637v.037c.406.057.64.309.64.68c0 .504-.48.851-1.158.851c-.67 0-1.125-.361-1.15-.916h.684c.01.217.185.352.457.352c.261 0 .439-.143.439-.356c0-.222-.168-.357-.443-.357z"/></g>`
	loaderInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 3.5v2m5 0L14 7M5.5 5.5L7 7m3.5 10.5v-2m5 0L14 14m-8.5 1.5L7 14m-3.5-3.5h2m10 0h2"/>`
	locationInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(4 2)"><path d="m6.5 16.54l.631-.711c.716-.82 1.36-1.598 1.933-2.338l.473-.624c1.975-2.661 2.963-4.773 2.963-6.334C12.5 3.201 9.814.5 6.5.5s-6 2.701-6 6.033c0 1.561.988 3.673 2.963 6.334l.473.624a54.84 54.84 0 0 0 2.564 3.05z"/><circle cx="6.5" cy="6.5" r="2.5"/></g>`
	lockInnerSVG                = `<g fill="none" fill-rule="evenodd" transform="translate(4 1)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 8.5l-.006-1.995C2.487 2.502 3.822.5 6.5.5s4.011 2.002 4 6.005V8.5m-8 0h8.023a2 2 0 0 1 1.994 1.85l.006.156l-.017 6a2 2 0 0 1-2 1.994H2.5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2z"/><circle cx="6.5" cy="13.5" r="1.5" fill="currentColor"/></g>`
	lockOpenInnerSVG            = `<g fill="none" fill-rule="evenodd" transform="translate(4 1)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 8.5l-.006-1.995C2.487 2.502 3.822.5 6.5.5c2.191 0 3.61 1.32 4 4m-8 4h8.023a2 2 0 0 1 1.994 1.85l.006.156l-.017 6a2 2 0 0 1-2 1.994H2.5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2z"/><circle cx="6.5" cy="13.5" r="1.5" fill="currentColor"/></g>`
	mailInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 6.5v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2h-10a2 2 0 0 0-2 2z"/><path d="m5.5 7.5l5 3l5-3"/></g>`
	mailAddInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 5.5h-8a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-5"/><path d="m4.5 8.5l5 3l5-3m2-5v4m-2-2h4"/></g>`
	mailDeleteInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.508 4.067l-5 2.857A2 2 0 0 0 3.5 8.661V15.5a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V8.66a2 2 0 0 0-1.008-1.736l-5-2.857a2 2 0 0 0-1.984 0zM13.5 8.5l-6 6m0-6l6 6"/>`
	mailMinusInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 5.5h-8a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-7"/><path d="m4.5 8.5l5 3l5-3m0-3h4"/></g>`
	mailNewInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.508 4.067l-5 2.857A2 2 0 0 0 3.5 8.661V15.5a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V8.66a2 2 0 0 0-1.008-1.736l-5-2.857a2 2 0 0 0-1.984 0zM10.5 8.5v6m-3-3h6"/>`
	mailOpenInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 6.819V14.5a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6.819a2 2 0 0 0-1.212-1.838l-5-2.143a2 2 0 0 0-1.576 0l-5 2.143A2 2 0 0 0 3.5 6.819z"/><path d="m5.5 7.5l5 3l5-3"/></g>`
	mailRemoveInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.508 4.067l-5 2.857A2 2 0 0 0 3.5 8.661V15.5a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V8.66a2 2 0 0 0-1.008-1.736l-5-2.857a2 2 0 0 0-1.984 0zM7.5 11.5h6"/>`
	marqueeInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 5.5v-1a2 2 0 0 1 2-2h1m0 16h-1a2 2 0 0 1-2-2v-1m16-10v-1a2 2 0 0 0-2-2h-1m0 16h1a2 2 0 0 0 2-2v-1m-11-13h2m2 0h2m-6 16h2m2 0h2m5-11.002V9.5m0 1.998V13.5m-16-6.002V9.5m0 1.998V13.5"/>`
	maximiseInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 16.5v-12a2 2 0 0 0-2-2h-12a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2z"/><path d="M2.5 14.5h10a2 2 0 0 0 2-2v-10"/><path d="M2.5 10.5h7a1 1 0 0 0 1-1v-7"/></g>`
	menuHamburgerInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h12m-12.002 4h11.997M4.5 14.5h11.995"/>`
	menuHorizontalInnerSVG      = `<g fill="currentColor" fill-rule="evenodd"><circle cx="10.5" cy="10.5" r="1"/><circle cx="5.5" cy="10.5" r="1"/><circle cx="15.5" cy="10.5" r="1"/></g>`
	menuVerticalInnerSVG        = `<g fill="currentColor" fill-rule="evenodd"><circle cx="10.5" cy="10.5" r="1"/><circle cx="10.5" cy="5.5" r="1"/><circle cx="10.5" cy="15.5" r="1"/></g>`
	messageInnerSVG             = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 3.5a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2l-2.999-.001l-2.294 2.294a1 1 0 0 1-1.32.083l-.094-.083l-2.294-2.294L4.5 17.5a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm-1 5h-6"/><path fill="currentColor" d="M6.499 9.5c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm0 4c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 12.5h-6"/></g>`
	messageWritingInnerSVG      = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 3.5a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2l-2.999-.001l-2.294 2.294a1 1 0 0 1-1.32.083l-.094-.083l-2.294-2.294L4.5 17.5a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M10.499 11.5c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm-4 0c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm8 0c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1z"/></g>`
	microphoneInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.39 2.615l.11-.004a2.893 2.893 0 0 1 3 2.891V9.5a3 3 0 1 1-6 0V5.613a3 3 0 0 1 2.89-2.998z"/><path d="M15.5 9.5a5 5 0 0 1-9.995.217L5.5 9.5m5 5v4"/></g>`
	microphoneDisabledInnerSVG  = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.5 9.5c0 .918-.247 1.778-.68 2.518m-1.424 1.558A5 5 0 0 1 5.5 9.5m2.196-4.955a3.001 3.001 0 0 1 2.693-1.93l.111-.004a2.893 2.893 0 0 1 3 2.891V9.5c0 .388-.074.759-.208 1.099"/><path d="M11.957 12.123A3 3 0 0 1 7.5 9.5v-2m-4-4l14 14m-7-3v4"/></g>`
	microphoneMutedInnerSVG     = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.5 8l.001 1.5c-.214 2-1.215 3-3.001 3s-2.785-1-2.999-3L7.5 6c0-2 1.857-3.231 2.5-3.5m1.5 4l4-4m0 4l-4-4z"/><path d="M15.5 9.5a5 5 0 0 1-9.995.217L5.5 9.5m5.022 5v4"/></g>`
	midpointInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 7.5v-3a2 2 0 0 0-2-2h-3m-3 10v-4m-2 2h4m6 3v3a2 2 0 0 1-2 2h-3m-6-16h-3a2 2 0 0 0-2 2v3m5 11h-3a2 2 0 0 1-2-2v-3"/>`
	miniPlayerInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 4.5h10a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M12.5 10.5h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1z"/></g>`
	minimiseInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 17.5v-5h-5m14 0h-5v5m-9-9h5v-5m4 0v5h5"/>`
	minusInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 10.5h10"/>`
	minusCircleInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10 19c4.438 0 8-3.526 8-7.964C18 6.598 14.438 3 10 3c-4.438 0-8 3.598-8 8.036S5.562 19 10 19zm-4-8h8"/>`
	moonInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 3.5c1.328 0 2.57.37 3.628 1.012a6 6 0 0 0-.001 11.977A7 7 0 1 1 11.5 3.5z"/>`
	moveInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5h-18m9-9v18m6-6l3-3l-3-3m-12 6l-3-3l3-3m3-3l3-3l3 3m-6 12l3 3l3-3"/>`
	newspaperInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4.5 7.165h9m-8.019 3.038l1-.018a1 1 0 0 1 1.01.864l.009.135v.984a1 1 0 0 1-.981 1l-1 .018a1 1 0 0 1-1.01-.864l-.009-.136v-.983a1 1 0 0 1 .981-1z"/><path d="M3.5 4.15h11a2 2 0 0 1 2 2v10.015h-13a2 2 0 0 1-2-2V6.151a2 2 0 0 1 2-2zm6 6.015h4m-4 3h4"/><path d="M16 16.165a2.5 2.5 0 0 0 2.5-2.5v-6.5h-2"/></g>`
	noSignInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="M14 3L3 14"/></g>`
	notebookInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 3.5h8a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm1 14v-14"/>`
	notificationInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 2)"><path d="M14.5 6.5v7a2 2 0 0 1-2 2H2.543a2 2 0 0 1-2-1.991l-.043-10A2 2 0 0 1 2.49 1.5H9.5"/><circle cx="14" cy="2" r="2" fill="currentColor"/></g>`
	nutInnerSVG                 = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(4 3)"><path d="m6.5.5l6 4v6l-6 4l-6-4v-6z"/><circle cx="6.5" cy="7.5" r="3"/></g>`
	pagesInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.5 3.5h-7a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-9a2 2 0 0 0-2-2z"/><path d="M6.5 5.5a2 2 0 0 0-2 2v8a3 3 0 0 0 3 3h6a2 2 0 0 0 2-2"/></g>`
	panelBottomInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm1 12h8"/>`
	panelCenterInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path d="M8.5 7.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z"/></g>`
	panelLeftInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm0 11v-8"/>`
	panelRightInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm10 11v-8"/>`
	panelSectionedInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.498 15.498l-.01-10a2 2 0 0 0-2-1.998h-10a2 2 0 0 0-1.995 1.85l-.006.152l.01 10a2 2 0 0 0 2 1.998h10a2 2 0 0 0 1.995-1.85zM10.5 3.5v13.817m7-6.817h-14"/>`
	panelTopInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm1 2h8"/>`
	paperInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 15.5v-8l-4-4h-6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/>`
	paperFoldedInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M16.5 15.5v-7l-5-5h-5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="M11.5 3.5v3a2 2 0 0 0 2 2h3"/></g>`
	paperPlaneInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m1.5 9l16-6.535L14.7 17zm16-6.5l-11 10m0 0v5l3-3"/>`
	paperPlaneAltInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.5 2.465l-8 8.033m3 8.002l-3-8.002l-7-2.998l15-5z"/>`
	paperclipInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.5 11.5l-5 5c-2.521 2.763-4.854 3.096-7 1c-2.146-2.097-1.813-4.43 1-7l5-5c1.919-2.081 3.585-2.415 5-1s1.08 3.08-1 5l-5 5c-.945 1.055-1.78 1.22-2.5.5s-.555-1.555.5-2.5l5-5"/>`
	paragraphCenterInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h12m-9.002 4h5.997m-7.995 4h9.995"/>`
	paragraphEndInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 5.5h13.978M3.5 7.5h13.978M3.5 9.5h13.978M3.5 11.5h13.978M3.5 13.5h13.978M3.5 15.5h7"/>`
	paragraphLeftInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h12m-12.002 4h5.997m-5.995 4h9.995"/>`
	paragraphRightInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6.5h12m-6.002 4h5.997m-9.995 4h9.995"/>`
	paragraphStartInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 7.5h13.978M3.5 9.5h13.978M3.5 11.5h13.978M3.5 5.5h7m-7 8h13.978M3.5 15.5h13.978"/>`
	penInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 4a2.121 2.121 0 0 1 0 3l-9.5 9.5l-4 1l1-3.944l9.504-9.552a2.116 2.116 0 0 1 2.864-.125zm-1.5 2.5l1 1"/>`
	phoneLandscapeInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 14.475V8.5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6a2 2 0 0 1-1.85 1.994l-.155.006l-10-.025a2 2 0 0 1-1.995-2zm12-4.975v4"/>`
	phonePortraitInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 3.5h6a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2zm1 12h4"/>`
	pictureInnerSVG             = `<g fill="none" fill-rule="evenodd" transform="translate(3 3)"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path d="m14.5 10.5l-3-3l-3 2.985m4 4.015l-9-9l-3 3"/></g><circle cx="11" cy="4" r="1" fill="currentColor"/></g>`
	pieHalfInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.519 2.747a8.003 8.003 0 0 0-.045 15.494M18.5 10.5a8 8 0 0 1-8 8v-16a8 8 0 0 1 8 8z"/>`
	pieQuarterInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.519 2.747a8 8 0 1 0 9.705 9.845"/><path d="M18.5 10.5a8 8 0 0 0-8-8v8z"/></g>`
	pieThirdInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.519 2.747a8.013 8.013 0 0 0-5.791 5.849M10.5 2.5a8 8 0 1 1 0 16c-4.418 0-8-3.5-8-8h8v-8z"/>`
	pillInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 7.5v5.817m-7-2.817a3 3 0 0 0 3 3h8a3 3 0 0 0 0-6h-8a3 3 0 0 0-3 3z"/>`
	playButtonInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4.494 5.535l12-.038a2 2 0 0 1 2 1.845l.006.155V13.5a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2V7.535a2 2 0 0 1 1.994-2z"/><path fill="currentColor" d="m9.5 12.5l3-2l-3-2z"/></g>`
	plusInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 10.5h10m-5-5v10"/>`
	plusCircleInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10 19c4.438 0 8-3.526 8-7.964C18 6.598 14.438 3 10 3c-4.438 0-8 3.598-8 8.036S5.562 19 10 19zm-4-8h8m-4 4.056V7z"/>`
	postcardInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 4.5h10a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2z"/><path d="M13.5 6.5h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1zm-8 5h5m-5 2h5"/></g>`
	printerInnerSVG             = `<g fill="none" fill-rule="evenodd" transform="rotate(-90 10.5 8.5)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 4.384V2.486a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2V14.5a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 4.5h5.001v8H1.5a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1zm12 0h2a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-2"/><circle cx="9" cy="14" r="1" fill="currentColor"/></g>`
	projectorInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 14.5V4.485h-14V14.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1zm-9 1l-2 3.5m6-3.5l2 3m-13-14h18"/><path d="M10.499 2.498a2.005 2.005 0 0 1 1.995 1.853l.006.149l-4-.002c-.001-1.105.894-2 1.999-2z"/></g>`
	pullDownInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 13.5l4 4l4-4m-4-7v11m-7-14h14"/>`
	pullLeftInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 6.5l-4 4l4 4m7-4h-11m14-7v14"/>`
	pullRightInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.5 14.5l4-4l-4-4m4 4h-11m-3-7v14"/>`
	pullUpInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 7.753l4-4.253l4 4.212m-4-4.212v11m-7 3h14"/>`
	pushDownInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 10.5l4 4l4-4m-4-7v11m-7 3h14"/>`
	pushLeftInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.5 6.5l-4 4l4 4m7-4h-11m-3-7v14"/>`
	pushRightInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.5 6.5l4 4l-4 4m4-4h-11m14-7v14"/>`
	pushUpInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 10.5l4-4l4 4m-4-4v11m-7-14h14"/>`
	questionCircleInnerSVG      = `<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 9.5v-1l1.414-1.414a2 2 0 0 0 .586-1.414V5.5c0-.613-.346-1.173-.894-1.447l-.212-.106a2 2 0 0 0-1.788 0L7.5 4c-.613.306-1 .933-1 1.618V6.5"/><circle cx="8.5" cy="12.5" r="1" fill="currentColor"/></g>`
	radioOnInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="10.5" cy="10.5" r="8"/><circle cx="10.5" cy="10.5" r="5" fill="currentColor"/></g>`
	receiptInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.5 8.5h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-14a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2"/><path d="M5.5 4.5h10V16a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1z"/><path d="m8.5 11.5l2 2l2-2m-2 2v-6"/></g>`
	recordInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="10.5" cy="10.5" r="5"/><circle cx="10.5" cy="10.5" r="3" fill="currentColor"/></g>`
	redoInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 13.5c-3.17-4-6.17-6-9-6s-5.163 1-7 3"/><path d="M13.5 13.5h5v-5"/></g>`
	refreshInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 3.5c-2.412 1.378-4 4.024-4 7a8 8 0 0 0 8 8m4-1c2.287-1.408 4-4.118 4-7a8 8 0 0 0-8-8"/><path d="M6.5 7.5v-4h-4m12 10v4h4"/></g>`
	refreshAltInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 6.5c1.378-2.412 4.024-4 7-4a8 8 0 0 1 8 8m-1 4c-1.408 2.287-4.118 4-7 4a8 8 0 0 1-8-8"/><path d="M8.5 6.5h-5v-5m9 13h5v5"/></g>`
	replicateInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 12.5v-7a2 2 0 0 0-2-2h-7a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2z"/><path d="M6.5 14.5v1a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2h-1"/></g>`
	replicateAltInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 12.5v-7a2 2 0 0 0-2-2h-7a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2z"/><path d="M14.5 14.5v1a2 2 0 0 1-2 2h-7a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h1"/></g>`
	resetInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.578 6.487A8 8 0 1 1 2.5 10.5"/><path d="M7.5 6.5h-4v-4"/></g>`
	resetAltInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 3.5c2.414 1.377 4 4.022 4 7a8 8 0 1 1-8-8"/><path d="M14.5 7.5v-4h4"/></g>`
	resetForwardInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 3.5c-2.414 1.377-4 4.022-4 7a8 8 0 1 0 8-8"/><path d="M6.5 7.5v-4h-4"/></g>`
	resetHardInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 5.5l-3-3h6a8.003 8.003 0 0 1 7.427 5.02c.37.921.573 1.927.573 2.98a8 8 0 1 1-16 0c0-1.49.37-3.472 1.538-5.091M7.5 7.5l6 6m-6 0l6-6"/>`
	resetTemporaryInnerSVG      = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="matrix(0 1 1 0 2.5 2.5)"><path d="M3.987 1.078A8 8 0 1 0 8 0"/><circle cx="8" cy="8" r="2" fill="currentColor"/><path d="M4 5V1H0"/></g>`
	retweetInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.5 13.5l3 3l3-3"/><path d="M9.5 4.5h3a4 4 0 0 1 4 4v8m-9-9l-3-3l-3 3"/><path d="M11.5 16.5h-3a4 4 0 0 1-4-4v-8"/></g>`
	reuseInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 2.5h5v5"/><path d="M8.5 2.5c-3.333 2.837-5 5.67-5 8.5s1 5.33 3 7.5m11 0h-5v-5"/><path d="M12.5 18.5c3.333-2.837 5-5.67 5-8.5s-1-5.33-3-7.5"/></g>`
	reverseInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.5 10.5l-4 4l4 4m8-4h-12m8-12l4 4l-4 4m4-4h-12"/>`
	reverseAltInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5 9.5l-4 4l4 4m8-4h-12m6-10l4 4l-4 4m4-4h-12"/>`
	revertInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m7.5 5.5l-3-3h6a8.003 8.003 0 0 1 7.427 5.02c.37.921.573 1.927.573 2.98a8 8 0 1 1-16 0c0-1.49.37-3.472 1.538-5.091"/><path d="M10.5 5.5v5h3"/></g>`
	rocketInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(4 1)"><path d="M6.5 18.5c3-2.502 4-5.502 4-9s-1-6.498-4-9c-3 2.502-4 5.502-4 9s1 6.498 4 9z"/><path d="M10.062 13.362a5.65 5.65 0 0 1 1.171.902c1.12 1.119 1.69 2.574 1.714 4.365c-2.509-.11-3.882-.765-4.926-1.647m-5.115-3.62a5.646 5.646 0 0 0-1.172.902C.615 15.383.044 16.838.021 18.629c2.508-.11 3.882-.765 4.926-1.647"/><circle cx="6.5" cy="6.5" r="2"/></g>`
	rulerInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 8.5h14a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm2 .5v2.5m2-2.5v2.5m2-2.5v3.5m2-3.5v2.5m2-2.5v2.5m2-2.5v3.5"/>`
	scaleInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 7.5v-4h-4m-4 4v4h4m4-8l-8 8"/><path d="M11.5 3.5h-6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-6"/></g>`
	scaleContractInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.5 9.5l-4 .022V5.5m-2 10.023v-4l-4-.023"/>`
	scaleExtendInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 9.5V5.522l-4-.022m-2 10.023h-4V11.5"/>`
	scalpelInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9 15l7-7a1.414 1.414 0 0 0-2-2L3.5 16.5h7L7 13"/>`
	searchInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="8.5" cy="8.5" r="5"/><path d="M17.571 17.5L12 12"/></g>`
	serverInnerSVG              = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 14.5v-2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M6.5 13.5a1 1 0 1 0-2 0a1 1 0 0 0 2 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 8.494l.01-2a2 2 0 0 1 2-1.994H16.5a2 2 0 0 1 1.994 1.85l.006.156l-.01 2a2 2 0 0 1-2 1.994H4.5a2 2 0 0 1-1.995-1.85z"/><path fill="currentColor" d="M6.5 7.5a1 1 0 1 0-2 0a1 1 0 0 0 2 0z"/></g>`
	settingsInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 3)"><path stroke-width=".933" d="M7.5.5c.351 0 .697.026 1.034.076l.508 1.539c.445.127.868.308 1.26.536l1.46-.704c.517.397.977.865 1.365 1.389l-.73 1.447c.221.396.395.822.514 1.27l1.53.533a7.066 7.066 0 0 1-.017 1.948l-1.539.508a5.567 5.567 0 0 1-.536 1.26l.704 1.46a7.041 7.041 0 0 1-1.389 1.365l-1.447-.73a5.507 5.507 0 0 1-1.27.514l-.533 1.53a7.066 7.066 0 0 1-1.948-.017l-.508-1.539a5.567 5.567 0 0 1-1.26-.536l-1.46.704a7.041 7.041 0 0 1-1.365-1.389l.73-1.447a5.565 5.565 0 0 1-.514-1.27l-1.53-.534a7.066 7.066 0 0 1 .017-1.947l1.539-.508c.127-.445.308-.868.536-1.26l-.704-1.46a7.041 7.041 0 0 1 1.389-1.365l1.447.73a5.507 5.507 0 0 1 1.27-.514l.534-1.53A7.06 7.06 0 0 1 7.5.5z"/><circle cx="7.5" cy="7.5" r="3"/></g>`
	shareInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.5 7.5l-3.978-4l-4.022 4m4.022-3.979V15.5m-3.022-5h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2"/>`
	shareAltInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.5 4.5l-1.978-2l-2.022 2m2-2v9m-3-5h-1a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2h-1"/>`
	shuffleInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m8.501 11.5l-3.001 3l3.001 3"/><path d="M16.5 9.5v2a3 3 0 0 1-3 3h-8m6.999-5l3.001-3l-3.001-3"/><path d="M4.5 11.5v-2a3 3 0 0 1 3-3h8"/></g>`
	sideMenuInnerSVG            = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 6.5h8m-8 3.998h5m-5 4.002h8"/><path fill="currentColor" d="M4.499 7.5c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm0 4c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1zm0 4c.5 0 1-.5 1-1s-.5-1-1-1s-.999.5-.999 1s.499 1 .999 1z"/></g>`
	signalFullInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 16.5v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-2 0zm4 0v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-2 0zm4 0v-9a1 1 0 1 1 2 0v9a1 1 0 0 1-2 0z"/>`
	signalLowInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="currentColor" d="M5.5 16.5v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-2 0z"/><path d="M9.5 16.5v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-2 0zm4 0v-9a1 1 0 1 1 2 0v9a1 1 0 0 1-2 0z"/></g>`
	signalMediumInnerSVG        = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="currentColor" d="M5.5 16.5v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-2 0zm4 0v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-2 0z"/><path d="M13.5 16.5v-9a1 1 0 1 1 2 0v9a1 1 0 0 1-2 0z"/></g>`
	signalNoneInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 16.5v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-2 0zm4 0v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-2 0zm4 0v-9a1 1 0 1 1 2 0v9a1 1 0 0 1-2 0z"/>`
	slashBackwardInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.5 3.5l4 14"/>`
	slashForwardInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.5 3.5l-4 14"/>`
	slidersInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 9V2.5m0 16V14"/><circle cx="14.5" cy="11.5" r="2.5"/><path d="M6.5 5V2.5m0 16V10"/><circle cx="6.5" cy="7.5" r="2.5"/></g>`
	sortInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.5 12.5l4 4.107l4-4.107m-8-4l-4-4l-4 3.997m4-3.997v12m8-12v12"/>`
	sortAltInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.5 8.5l4 4l4-4m-4-6v10m-4 .044L6.5 8.5l-4 4.044m4-4.044v10"/>`
	speakerInnerSVG             = `<g fill="none" fill-rule="evenodd" transform="translate(5 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h6a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><circle cx="5.5" cy="9.5" r="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="5.5" cy="3.5" r="1" fill="currentColor"/></g>`
	speechBubbleInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 16.517c4.418 0 8-3.284 8-7.017C19 5.767 15.418 3 11 3S3 6.026 3 9.759c0 1.457.546 2.807 1.475 3.91L3.5 18.25l3.916-2.447a9.181 9.181 0 0 0 3.584.714z"/>`
	speechTypingInnerSVG        = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 16.517c4.418 0 8-3.026 8-6.758C19 6.026 15.418 3 11 3S3 6.026 3 9.759c0 1.457.546 2.807 1.475 3.91L3.5 18.25l3.916-2.447a9.181 9.181 0 0 0 3.584.714z"/><path fill="currentColor" d="M10.999 11c.5 0 1-.5 1-1s-.5-1-1-1S10 9.5 10 10s.499 1 .999 1zm-4 0c.5 0 1-.5 1-1s-.5-1-1-1S6 9.5 6 10s.499 1 .999 1zm8 0c.5 0 1.001-.5 1.001-1s-.5-1-1-1s-1 .5-1 1s.5 1 1 1z"/></g>`
	splitInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 5.5h-10a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zm-5 0v10"/>`
	splitThreeInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 5.5h-12a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zm-9 0v10m6-10v10"/>`
	starInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.5 14.5l-5 3l2-5.131l-4-3.869h5l2-5l2 5h5l-4 4l2 5z"/>`
	sunInnerSVG                 = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 14.5c2.219 0 4-1.763 4-3.982a4.003 4.003 0 0 0-4-4.018c-2.219 0-4 1.781-4 4c0 2.219 1.781 4 4 4zM4.136 4.136L5.55 5.55m9.9 9.9l1.414 1.414M1.5 10.5h2m14 0h2M4.135 16.863L5.55 15.45m9.899-9.9l1.414-1.415M10.5 19.5v-2m0-14v-2" opacity=".3"/><g transform="translate(-210 -1)"><path d="M220.5 2.5v2m6.5.5l-1.5 1.5"/><circle cx="220.5" cy="11.5" r="4"/><path d="m214 5l1.5 1.5m5 14v-2m6.5-.5l-1.5-1.5M214 18l1.5-1.5m-4-5h2m14 0h2"/></g></g>`
	supportInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="10.5" cy="10.5" r="8"/><circle cx="10.5" cy="10.5" r="4"/><path d="M13.5 7.5L16 5m-2.5 8.5L16 16m-8.5-2.5L5 16m2.5-8.5L5 5"/></g>`
	swapInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13 8h5V3"/><path d="M18 8c-2.837-3.333-5.67-5-8.5-5S4.17 4 2 6m4.5 5.5h-5v5"/><path d="M1.5 11.5c2.837 3.333 5.67 5 8.5 5s5.33-1 7.5-3"/></g>`
	switchInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(6 3)"><path d="M1.5.5h6a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1z"/><circle cx="4.5" cy="4" r="1.5"/><path d="M.5 7.5h8m-4 2v3"/></g>`
	tableInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.498 15.498l-.01-10a2 2 0 0 0-2-1.998h-10a2 2 0 0 0-1.995 1.85l-.006.152l.01 10a2 2 0 0 0 2 1.998h10a2 2 0 0 0 1.995-1.85zM7.5 3.5v13.817m10-9.817h-14"/>`
	tableHeaderInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.498 15.498l-.01-10a2 2 0 0 0-2-1.998h-10a2 2 0 0 0-1.995 1.85l-.006.152l.01 10a2 2 0 0 0 2 1.998h10a2 2 0 0 0 1.995-1.85zM7.5 7.5v9.817m10-9.817h-14"/>`
	tagInnerSVG                 = `<g fill="none" fill-rule="evenodd" transform="translate(3 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.914.5H12.5a2 2 0 0 1 2 2v3.586a1 1 0 0 1-.293.707l-6.793 6.793a2 2 0 0 1-2.828 0l-3.172-3.172a2 2 0 0 1 0-2.828L8.207.793A1 1 0 0 1 8.914.5z"/><circle cx="12" cy="3" r="1" fill="currentColor"/></g>`
	tagMilestoneInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 9.224V15.5a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9.224a2 2 0 0 0-.464-1.28l-3.768-4.522a1 1 0 0 0-1.536 0L5.964 7.944a2 2 0 0 0-.464 1.28z"/>`
	tagsInnerSVG                = `<g fill="none" fill-rule="evenodd" transform="translate(1 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.914.5H15.5a2 2 0 0 1 2 2v3.586a1 1 0 0 1-.293.707l-6.793 6.793a2 2 0 0 1-2.828 0l-3.172-3.172a2 2 0 0 1 0-2.828L11.207.793A1 1 0 0 1 11.914.5z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 13.5l-2.013 1.006A2 2 0 0 1 2.72 13.42L1.105 9.114a2 2 0 0 1 .901-2.45L9.5 2.5"/><rect width="2" height="2" x="14" y="2" fill="currentColor" rx="1"/></g>`
	targetInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="10.5" cy="10.5" r="8"/><circle cx="10.5" cy="10.5" r="2"/><circle cx="10.5" cy="10.5" r="5"/></g>`
	terminalInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 4.5h10a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2zm5 9h3"/><path d="m6.5 12.5l2-2l-2-2"/></g>`
	threadInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.5 5.5a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3l2.468-.001l1.715 2.43a1 1 0 0 0 .696.415l.121.008a1 1 0 0 0 .993-.884l.007-.116l.001-1.853l.999.001a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3z"/><path d="m6.5 13.5l-2 2v-4h-.906a2 2 0 0 1-2-1.977l-.07-6a2 2 0 0 1 2-2.023H12.5a2 2 0 0 1 2 2v2"/></g>`
	thumbsDownInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.643 4.243L10.499 5.5h-4v7h2L11.3 18c.58 0 1.075-.205 1.485-.615c.41-.41.615-.905.615-1.485l-.9-2.4l4.031-1.344a2 2 0 0 0 1.309-2.38l-.069-.22l-1.553-4.142a2 2 0 0 0-2.575-1.17z"/><path d="M3.5 13.5h2a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1z"/></g>`
	thumbsUpInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.643 16.757L10.499 15.5h-4v-7h2L11.3 3c.58 0 1.075.205 1.485.615c.41.41.615.905.615 1.485l-.9 2.4l4.031 1.344a2 2 0 0 1 1.309 2.38l-.069.22l-1.553 4.142a2 2 0 0 1-2.575 1.17z"/><path d="M3.5 7.5h2a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1z"/></g>`
	ticketInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 2.486V2.5a2 2 0 1 0 4 0l-.001-.015H14.5a1 1 0 0 1 1 1V16.5a1 1 0 0 1-1 1H12a2 2 0 1 0-4 0H5.5a1 1 0 0 1-1-1V3.485a1 1 0 0 1 1-1zM6.5 6.5h1m2 0h1m2 0h1m-7 7h1m2 0h1m2 0h1"/>`
	timelineInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 7.5h7m-3.002 4h6.669m-6.669-2H12.5m-3.002 4H17.5"/>`
	todoInnerSVG                = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><g fill="currentColor" transform="translate(3 3)"><circle cx="7.5" cy="7.5" r="1" transform="matrix(-1 0 0 1 15 0)"/><circle cx="3.5" cy="7.5" r="1"/><circle cx="11.5" cy="7.5" r="1"/></g></g>`
	toggleInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 7)"><circle cx="3.5" cy="3.5" r="3"/><path d="M6 1.5h6.5c.828 0 2 .325 2 2s-1.172 2-2 2H6"/></g>`
	togglesInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 4)"><circle cx="3.5" cy="3.5" r="3"/><path d="M6 1.5h6.5c.828 0 2 .325 2 2s-1.172 2-2 2H6m5.5 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/><path d="M9 8.5H2.5c-.828 0-2 .325-2 2s1.172 2 2 2H9"/></g>`
	translateInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 10.5v-6a2 2 0 0 0-2-2h-6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2z"/><path d="M6.5 8.503h-2a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.003l6-.01a2 2 0 0 0 1.997-2V14.5m-5-1.997h-3"/><path d="m9 14l-1 1c-.334.333-1.166.833-2.5 1.5"/><path d="M5.5 12.503c.334 1.166.834 2 1.5 2.499S8.5 16 9.5 16.5m4-12l-3 6m3-6l3 6m-1-2h-4"/></g>`
	trashInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 4.5h10v12a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2zm5-2a2 2 0 0 1 1.995 1.85l.005.15h-4a2 2 0 0 1 2-2zm-7 2h14m-9 3v8m4-8v8"/>`
	trashAltInnerSVG            = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.043 4.773c1.305-.502 2.79-.757 4.457-.763c1.667-.007 3.152.248 4.457.763a3 3 0 0 1 2.14 3.341l-1.075 6.994a4 4 0 0 1-3.954 3.392H8.932a4 4 0 0 1-3.954-3.392L3.902 8.114a3 3 0 0 1 2.141-3.34z"/><path fill="currentColor" d="M10.5 10c3.556 0 5-1.5 5-2.5s-1.444-2.25-5-2.25s-5 1.25-5 2.25s1.444 2.5 5 2.5z"/></g>`
	trophyInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 3.5h6a1 1 0 0 1 1 1v5a4 4 0 1 1-8 0v-5a1 1 0 0 1 1-1zm3 10v3m-3 0h6a1 1 0 0 1 0 2h-6a1 1 0 0 1 0-2zm7-11h2a1 1 0 0 1 1 1v1a2 2 0 0 1-2 2h-1zm-8 0h-2a1 1 0 0 0-1 1v1a2 2 0 0 0 2 2h1z"/>`
	tvModeInnerSVG              = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.493 5.534l10-.036a2 2 0 0 1 2.007 2V12.5a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2V7.534a2 2 0 0 1 1.993-2z"/><path fill="currentColor" fill-rule="nonzero" d="M12.467 9.6L9.8 7.6A.5.5 0 0 0 9 8v4a.5.5 0 0 0 .8.4l2.667-2a.5.5 0 0 0 0-.8z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.464 16.5H15.5"/></g>`
	unarchiveInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 7.5h14v8a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2zm0-3.978h14a1 1 0 0 1 1 1V6.5a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V4.522a1 1 0 0 1 1-1z"/><path d="m7.5 13.5l3-3l3 3"/></g>`
	undoInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 13.5c3.333-4 6.333-6 9-6s5 1 7 3"/><path d="M2.5 8.5v5h5"/></g>`
	undoHistoryInnerSVG         = `<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 10.55a8 8 0 1 1 1.073 3.952"/><path fill="currentColor" fill-rule="nonzero" d="m2.5 13.5l2.5-3H0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 6.5v5h3"/></g>`
	unlinkHorizontalInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 14.5h-1a4 4 0 1 1 0-8h1m4 0h1a4 4 0 1 1 0 8h-1"/>`
	unlinkVerticalInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 12.5v1a4 4 0 1 1-8 0v-1m0-4v-1a4 4 0 1 1 8 0v1"/>`
	uploadInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.5 7.753l4-4.232l4 4.191m-4-4.212v11m-6 3h12"/>`
	uploadAltInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.5 7.5l-3.978-4l-4.022 4m4.022-3.979V15.5M3.5 12v4.5a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V12"/>`
	upwardInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.21 10.211c-3.405 3.852-6.975 5.279-10.71 4.281c3.711-.99 6.282-3.418 7.711-7.28L8.5 4.5h8v8z"/>`
	userInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="m11.5 4.5l2 2v1a3 3 0 0 1-5.995.176L7.5 6.5z"/><path d="M5.5 12V7.5a5 5 0 1 1 10 0V12"/><path stroke-linecap="round" d="M17.5 16.5v-.728c0-3.187-3.686-5.272-7-5.272s-7 2.085-7 5.272v.728a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1z"/></g>`
	userAddInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 2.5a3 3 0 0 1 3 3v2a3 3 0 1 1-6 0v-2a3 3 0 0 1 3-3zm7 2v4m2-2h-4"/><path d="M17.5 16.5v-.728c0-3.187-3.686-5.272-7-5.272s-7 2.085-7 5.272v.728a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1z"/></g>`
	userCircleInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8" stroke-linecap="round"/><path stroke-linecap="round" d="m9.5 4.5l2 2v1a3 3 0 0 1-6 0v-1z"/><path d="M3.5 12V7.5a5 5 0 1 1 10 0V12"/><path stroke-linecap="round" d="M14.5 13.404c-.662-2.273-3.2-2.93-6-2.93c-2.727 0-5.27.774-6 2.93"/></g>`
	userMaleInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 2.5a3 3 0 0 1 3 3v2a3 3 0 1 1-6 0v-2a3 3 0 0 1 3-3zm7 14v-.728c0-3.187-3.686-5.272-7-5.272s-7 2.085-7 5.272v.728a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1z"/>`
	userMaleCircleInnerSVG      = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8"/><path d="M14.5 13.5c-.662-2.274-3.2-3.025-6-3.025c-2.727 0-5.27.869-6 3.025"/><path d="M8.5 2.5a3 3 0 0 1 3 3v2a3 3 0 0 1-6 0v-2a3 3 0 0 1 3-3z"/></g>`
	userRemoveInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 2.5a3 3 0 0 1 3 3v2a3 3 0 1 1-6 0v-2a3 3 0 0 1 3-3zm9 4h-4"/><path d="M17.5 16.5v-.728c0-3.187-3.686-5.272-7-5.272s-7 2.085-7 5.272v.728a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1z"/></g>`
	usersInnerSVG               = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 2.5a3 3 0 0 1 3 3v2a3 3 0 1 1-6 0v-2a3 3 0 0 1 3-3zm7 14v-.728c0-3.187-3.686-5.272-7-5.272s-7 2.085-7 5.272v.728a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1z"/><path fill="currentColor" d="M12.52 2.678A3.001 3.001 0 0 1 14.5 5.5v1c0 1.297-.848 2.581-2 3c.674-.919 1.01-2.086 1.01-3.5s-.331-2.523-.99-3.322zM17.5 17.5h1a1 1 0 0 0 1-1v-.728c0-2.17-1.71-3.83-3.847-4.667c0 0 2.847 2.395 1.847 6.395z"/></g>`
	vennInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="13.5" cy="10.5" r="5"/><circle cx="7.5" cy="10.5" r="5"/></g>`
	versionInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 12.5l8 4l8.017-4M2.5 8.657l8.008 3.843l8.009-3.843L10.508 4.5z"/>`
	versionsInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.5 10.5l8 4l8.017-4M2.5 14.5l8 4l8.017-4M2.5 6.657l8.008 3.843l8.009-3.843L10.508 2.5z"/>`
	videoInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 6.5h6a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2zm8 3l2.4-1.8a1 1 0 0 1 1.6.8v4a1 1 0 0 1-1.6.8l-2.4-1.8z"/>`
	volumeAddInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 7.5h3l5-5v16l-5-5h-3a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm10 3h4m-2 2v-4z"/>`
	volumeDisabledInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 9.5v9l-5-5h-3a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h3L8 6m1.521-1.521L11.5 2.5v5m-6-4l12 12m-4-7v1m2.22 4.208c-.337.475-1.077 1.073-2.22 1.792m0-4v1m3-.5v-1.5c0-1.828-.833-3.328-2.5-4.5"/>`
	volumeHighInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 7.5h3l5-5v16l-5-5h-3a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm10 8c1.333-1 2-2.667 2-5s-.667-4-2-5m0 3v4"/>`
	volumeLowInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 7.5h3l5-5v16l-5-5h-3a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm10 1v4"/>`
	volumeMinusInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 7.5h3l5-5v16l-5-5h-3a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm10 3h4"/>`
	volumeMutedInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 7.5h3l5-5v16l-5-5h-3a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1zm10 1l4 4m-4 0l4-4z"/>`
	volumeZeroInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 7.5h3l5-5v16l-5-5h-3a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z"/>`
	walletInnerSVG              = `<g fill="none" fill-rule="evenodd" transform="translate(3 4)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 2.5h12a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2zm1-2h9a1 1 0 0 1 1 1v1H.5v-1a1 1 0 0 1 1-1z"/><circle cx="11.5" cy="7.5" r="1" fill="currentColor"/></g>`
	warningCircleInnerSVG       = `<g fill="none" fill-rule="evenodd"><circle cx="10.5" cy="10.5" r="8" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 11.5v-5"/><circle cx="10.5" cy="14.5" r="1" fill="currentColor"/></g>`
	warningHexInnerSVG          = `<g fill="none" fill-rule="evenodd" transform="translate(-1 -1)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.517 3.5l4.983 5v6l-4.983 5H8.5l-5-5v-6l5-5zm-3.017 9v-5"/><circle cx="11.5" cy="15.5" r="1" fill="currentColor"/></g>`
	warningTriangleInnerSVG     = `<g fill="none" fill-rule="evenodd" transform="translate(1 1)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5.5l9 16H.5zm0 10v-5"/><circle cx="9.5" cy="13.5" r="1" fill="currentColor"/></g>`
	wavesInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 14.5a8 8 0 0 0-8-8m5 8a5 5 0 0 0-5-5m2 5a2 2 0 0 0-2-2"/>`
	widthInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.507 14.515l4-4l-4-4.015m-10 8.015l-4-4l4-4.015m13.993 4h-18"/>`
	wifiInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 6)"><path d="M2.727 5.033c2.781-2.264 6.82-2.264 9.6 0M.287 2.667c4.122-3.554 10.304-3.554 14.427 0m-9.58 4.74a4.167 4.167 0 0 1 4.739 0"/><circle cx="7.5" cy="9.5" r="1" fill="currentColor"/></g>`
	wifiErrorInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 4)"><path d="M2.727 7.033A7.539 7.539 0 0 1 5.492 5.61m4.05-.005a7.54 7.54 0 0 1 2.785 1.43M7.5 8.5l.027-8M.286 4.667A10.974 10.974 0 0 1 5.511 2.18m4.087.02a10.972 10.972 0 0 1 5.116 2.467m-9.58 4.74c.161-.112.328-.211.5-.298m3.706-.016c.183.09.361.195.533.314"/><circle cx="7.5" cy="11.5" r="1" fill="currentColor"/></g>`
	wifiNoneInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 4)"><path d="M2.727 7.033a7.555 7.555 0 0 1 3.46-1.58m2.495-.031a7.56 7.56 0 0 1 3.645 1.611M.287 4.667a10.936 10.936 0 0 1 3.331-1.969m2.09-.552c3.141-.51 6.465.33 9.006 2.52M1 0l13 13.071M5.134 9.407a4.167 4.167 0 0 1 4.739 0"/><circle cx="7.5" cy="11.5" r="1" fill="currentColor"/></g>`
	windowInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M5.5 5.5h10a2 2 0 0 1 2 2v-2c0-1-.895-2-2-2h-10c-1.105 0-2 1-2 2v2a2 2 0 0 1 2-2z"/></g>`
	windowCollapseLeftInnerSVG  = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 15.5v-10a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M5.5 15.5v-10a2 2 0 0 1 2-2h-2c-1 0-2 .895-2 2v10c0 1.105 1 2 2 2h2a2 2 0 0 1-2-2z"/><path d="m10.5 13.5l-3-3l3-3m5 3h-8"/></g>`
	windowCollapseRightInnerSVG = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 15.5v-10a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M15.5 15.5v-10a2 2 0 0 0-2-2h2c1 0 2 .895 2 2v10c0 1.105-1 2-2 2h-2a2 2 0 0 0 2-2z"/><path d="m10.5 13.5l3-3l-3-3m3 3h-8"/></g>`
	windowContentInnerSVG       = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M5.5 5.5h10a2 2 0 0 1 2 2v-2c0-1-.895-2-2-2h-10c-1.105 0-2 1-2 2v2a2 2 0 0 1 2-2z"/><path d="M7.498 10.5h1m-1-2h3.997m-3.997 4h5.997m-5.997 2h3.997"/></g>`
	wrapBackInnerSVG            = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.499 6.498L3.5 9.5l3 3"/><path d="M8.5 15.5h5c2 0 3-1 3-3s-1-3-3-3h-10"/></g>`
	wrapForwardInnerSVG         = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m14.5 12.5l2.998-3.002l-3-3"/><path d="M12.5 15.5h-5c-2 0-3-1-3-3s1-3 3-3h10"/></g>`
	writeInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 4a2.121 2.121 0 0 1 0 3l-9.5 9.5l-4 1l1-3.944l9.504-9.552a2.116 2.116 0 0 1 2.864-.125zM9.5 17.5h8m-2-11l1 1"/>`
	zoomCancelInnerSVG          = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 3)"><circle cx="5.5" cy="5.5" r="5"/><path d="m7.5 7.5l-4-4l4 4zm-4 0l4-4l-4 4zm11 7L9.076 9.076"/></g>`
	zoomInInnerSVG              = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 3)"><circle cx="5.5" cy="5.5" r="5"/><path d="M7.5 5.5h-4zm-2 2v-4zm9 7L9.133 9.133"/></g>`
	zoomOutInnerSVG             = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 3)"><circle cx="5.5" cy="5.5" r="5"/><path d="M7.5 5.5h-4zm7.071 9l-5.45-5.381"/></g>`
	zoomResetInnerSVG           = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 8.5a5 5 0 1 0 1.057-3.074"/><path d="M4.5 1.5v4h4m9 12l-5.379-5.379"/></g>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Airplay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(airplayInnerSVG).
		Element(children...)
}

func AlarmClock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alarmClockInnerSVG).
		Element(children...)
}

func AlignHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignHorizontalInnerSVG).
		Element(children...)
}

func AlignVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignVerticalInnerSVG).
		Element(children...)
}

func Angle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleInnerSVG).
		Element(children...)
}

func Archive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archiveInnerSVG).
		Element(children...)
}

func ArrowBottomLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBottomLeftInnerSVG).
		Element(children...)
}

func ArrowBottomRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBottomRightInnerSVG).
		Element(children...)
}

func ArrowDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownCircleInnerSVG).
		Element(children...)
}

func ArrowLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightCircleInnerSVG).
		Element(children...)
}

func ArrowTopLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTopLeftInnerSVG).
		Element(children...)
}

func ArrowTopRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTopRightInnerSVG).
		Element(children...)
}

func ArrowUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpCircleInnerSVG).
		Element(children...)
}

func AudioWave(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(audioWaveInnerSVG).
		Element(children...)
}

func Backspace(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backspaceInnerSVG).
		Element(children...)
}

func Backward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backwardInnerSVG).
		Element(children...)
}

func Bag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagInnerSVG).
		Element(children...)
}

func BatteryCharging(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFullInnerSVG).
		Element(children...)
}

func BatteryHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryHalfInnerSVG).
		Element(children...)
}

func BatteryLow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryLowInnerSVG).
		Element(children...)
}

func BatterySeventyFive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batterySeventyFiveInnerSVG).
		Element(children...)
}

func Bell(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellInnerSVG).
		Element(children...)
}

func BellDisabled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellDisabledInnerSVG).
		Element(children...)
}

func BellRinging(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellRingingInnerSVG).
		Element(children...)
}

func BellSnooze(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellSnoozeInnerSVG).
		Element(children...)
}

func Bluetooth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothInnerSVG).
		Element(children...)
}

func Book(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookInnerSVG).
		Element(children...)
}

func BookClosed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookClosedInnerSVG).
		Element(children...)
}

func BookText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookTextInnerSVG).
		Element(children...)
}

func Bookmark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkInnerSVG).
		Element(children...)
}

func BookmarkBook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkBookInnerSVG).
		Element(children...)
}

func Box(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxInnerSVG).
		Element(children...)
}

func BoxAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxAddInnerSVG).
		Element(children...)
}

func BoxDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxDownloadInnerSVG).
		Element(children...)
}

func BoxOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOpenInnerSVG).
		Element(children...)
}

func BoxRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxRemoveInnerSVG).
		Element(children...)
}

func Boxes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxesInnerSVG).
		Element(children...)
}

func Branch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(branchInnerSVG).
		Element(children...)
}

func Briefcase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(browserInnerSVG).
		Element(children...)
}

func BrowserAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(browserAltInnerSVG).
		Element(children...)
}

func ButtonAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buttonAddInnerSVG).
		Element(children...)
}

func ButtonMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buttonMinusInnerSVG).
		Element(children...)
}

func Calculator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarInnerSVG).
		Element(children...)
}

func CalendarAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarAddInnerSVG).
		Element(children...)
}

func CalendarDate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarDateInnerSVG).
		Element(children...)
}

func CalendarDay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarDayInnerSVG).
		Element(children...)
}

func CalendarDays(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarDaysInnerSVG).
		Element(children...)
}

func CalendarLastDay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarLastDayInnerSVG).
		Element(children...)
}

func CalendarMonth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMonthInnerSVG).
		Element(children...)
}

func CalendarMove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMoveInnerSVG).
		Element(children...)
}

func CalendarRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarRemoveInnerSVG).
		Element(children...)
}

func CalendarSplit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarSplitInnerSVG).
		Element(children...)
}

func CalendarWeek(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarWeekInnerSVG).
		Element(children...)
}

func Camera(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraInnerSVG).
		Element(children...)
}

func CameraAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraAltInnerSVG).
		Element(children...)
}

func CameraNoflash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraNoflashInnerSVG).
		Element(children...)
}

func CameraNoflashAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraNoflashAltInnerSVG).
		Element(children...)
}

func Capture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(captureInnerSVG).
		Element(children...)
}

func CardTimeline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardTimelineInnerSVG).
		Element(children...)
}

func CardView(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardViewInnerSVG).
		Element(children...)
}

func Carousel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(carouselInnerSVG).
		Element(children...)
}

func Cart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartInnerSVG).
		Element(children...)
}

func Cast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(castInnerSVG).
		Element(children...)
}

func Chain(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chainInnerSVG).
		Element(children...)
}

func ChatAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatAddInnerSVG).
		Element(children...)
}

func Check(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkCircleInnerSVG).
		Element(children...)
}

func CheckCircleOutside(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkCircleOutsideInnerSVG).
		Element(children...)
}

func CheckboxChecked(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxCheckedInnerSVG).
		Element(children...)
}

func CheckboxEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxEmptyInnerSVG).
		Element(children...)
}

func ChevronClose(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronCloseInnerSVG).
		Element(children...)
}

func ChevronDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronDownInnerSVG).
		Element(children...)
}

func ChevronDownCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronDownCircleInnerSVG).
		Element(children...)
}

func ChevronDownDouble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronDownDoubleInnerSVG).
		Element(children...)
}

func ChevronLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronLeftInnerSVG).
		Element(children...)
}

func ChevronLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronLeftCircleInnerSVG).
		Element(children...)
}

func ChevronLeftDouble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronLeftDoubleInnerSVG).
		Element(children...)
}

func ChevronOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronOpenInnerSVG).
		Element(children...)
}

func ChevronRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightInnerSVG).
		Element(children...)
}

func ChevronRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightCircleInnerSVG).
		Element(children...)
}

func ChevronRightDouble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightDoubleInnerSVG).
		Element(children...)
}

func ChevronUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpInnerSVG).
		Element(children...)
}

func ChevronUpCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpCircleInnerSVG).
		Element(children...)
}

func ChevronUpDouble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpDoubleInnerSVG).
		Element(children...)
}

func Circle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleInnerSVG).
		Element(children...)
}

func CircleMenu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleMenuInnerSVG).
		Element(children...)
}

func CircleSplit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleSplitInnerSVG).
		Element(children...)
}

func Clipboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardInnerSVG).
		Element(children...)
}

func ClipboardAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardAddInnerSVG).
		Element(children...)
}

func ClipboardCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardCheckInnerSVG).
		Element(children...)
}

func ClipboardCopy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardCopyInnerSVG).
		Element(children...)
}

func ClipboardCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardCrossInnerSVG).
		Element(children...)
}

func ClipboardNotes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardNotesInnerSVG).
		Element(children...)
}

func ClipboardRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardRemoveInnerSVG).
		Element(children...)
}

func Clock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeInnerSVG).
		Element(children...)
}

func Cloud(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudInnerSVG).
		Element(children...)
}

func CloudDisconnect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDisconnectInnerSVG).
		Element(children...)
}

func CloudDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDownloadInnerSVG).
		Element(children...)
}

func CloudDownloadAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDownloadAltInnerSVG).
		Element(children...)
}

func CloudUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudUploadInnerSVG).
		Element(children...)
}

func CloudUploadAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudUploadAltInnerSVG).
		Element(children...)
}

func Code(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coffeeInnerSVG).
		Element(children...)
}

func Coin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coinInnerSVG).
		Element(children...)
}

func Coins(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coinsInnerSVG).
		Element(children...)
}

func Compass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassInnerSVG).
		Element(children...)
}

func ComponentAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(componentAddInnerSVG).
		Element(children...)
}

func Contacts(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contactsInnerSVG).
		Element(children...)
}

func Contract(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contractInnerSVG).
		Element(children...)
}

func Create(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(createInnerSVG).
		Element(children...)
}

func CreditCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropInnerSVG).
		Element(children...)
}

func Cross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crossInnerSVG).
		Element(children...)
}

func CrossCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crossCircleInnerSVG).
		Element(children...)
}

func Crosshair(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crosshairInnerSVG).
		Element(children...)
}

func Cube(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cubesInnerSVG).
		Element(children...)
}

func Cylinder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cylinderInnerSVG).
		Element(children...)
}

func Database(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseInnerSVG).
		Element(children...)
}

func Diamond(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diamondInnerSVG).
		Element(children...)
}

func Directions(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(directionsInnerSVG).
		Element(children...)
}

func Disc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discInnerSVG).
		Element(children...)
}

func Display(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(displayInnerSVG).
		Element(children...)
}

func DisplayAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(displayAltInnerSVG).
		Element(children...)
}

func Document(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentInnerSVG).
		Element(children...)
}

func DocumentJustified(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentJustifiedInnerSVG).
		Element(children...)
}

func DocumentList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentListInnerSVG).
		Element(children...)
}

func DocumentStack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentStackInnerSVG).
		Element(children...)
}

func DocumentWords(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentWordsInnerSVG).
		Element(children...)
}

func Door(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doorInnerSVG).
		Element(children...)
}

func DoorAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doorAltInnerSVG).
		Element(children...)
}

func Download(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadAltInnerSVG).
		Element(children...)
}

func Downward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downwardInnerSVG).
		Element(children...)
}

func Drag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragInnerSVG).
		Element(children...)
}

func DragCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragCircleInnerSVG).
		Element(children...)
}

func DragVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragVerticalInnerSVG).
		Element(children...)
}

func Duplicate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(duplicateInnerSVG).
		Element(children...)
}

func DuplicateAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(duplicateAltInnerSVG).
		Element(children...)
}

func Enter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enterInnerSVG).
		Element(children...)
}

func EnterAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enterAltInnerSVG).
		Element(children...)
}

func Episodes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(episodesInnerSVG).
		Element(children...)
}

func ExitLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exitLeftInnerSVG).
		Element(children...)
}

func ExitRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exitRightInnerSVG).
		Element(children...)
}

func Expand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandInnerSVG).
		Element(children...)
}

func ExpandHeight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandHeightInnerSVG).
		Element(children...)
}

func ExpandWidth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(expandWidthInnerSVG).
		Element(children...)
}

func External(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(externalInnerSVG).
		Element(children...)
}

func Eye(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeInnerSVG).
		Element(children...)
}

func EyeClosed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeClosedInnerSVG).
		Element(children...)
}

func EyeNo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeNoInnerSVG).
		Element(children...)
}

func FaceDelighted(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceDelightedInnerSVG).
		Element(children...)
}

func FaceHappy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceHappyInnerSVG).
		Element(children...)
}

func FaceNeutral(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceNeutralInnerSVG).
		Element(children...)
}

func FaceSad(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceSadInnerSVG).
		Element(children...)
}

func FileDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileDownloadInnerSVG).
		Element(children...)
}

func FileUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileUploadInnerSVG).
		Element(children...)
}

func FilesHistory(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filesHistoryInnerSVG).
		Element(children...)
}

func FilesMulti(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filesMultiInnerSVG).
		Element(children...)
}

func FilesStack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filesStackInnerSVG).
		Element(children...)
}

func Film(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterInnerSVG).
		Element(children...)
}

func FilterCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterCircleInnerSVG).
		Element(children...)
}

func FilterSingle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterSingleInnerSVG).
		Element(children...)
}

func Filtering(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filteringInnerSVG).
		Element(children...)
}

func Fingerprint(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerprintInnerSVG).
		Element(children...)
}

func Flag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagInnerSVG).
		Element(children...)
}

func Flame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flameInnerSVG).
		Element(children...)
}

func FlameAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flameAltInnerSVG).
		Element(children...)
}

func FlipView(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipViewInnerSVG).
		Element(children...)
}

func Floppy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floppyInnerSVG).
		Element(children...)
}

func FolderAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderAddInnerSVG).
		Element(children...)
}

func FolderClosed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderClosedInnerSVG).
		Element(children...)
}

func FolderMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderMinusInnerSVG).
		Element(children...)
}

func FolderOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderOpenInnerSVG).
		Element(children...)
}

func ForkGit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forkGitInnerSVG).
		Element(children...)
}

func Forward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardInnerSVG).
		Element(children...)
}

func Frame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameInnerSVG).
		Element(children...)
}

func Fullscreen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fullscreenInnerSVG).
		Element(children...)
}

func Funnel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(funnelInnerSVG).
		Element(children...)
}

func Gauge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gaugeInnerSVG).
		Element(children...)
}

func Gift(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(giftInnerSVG).
		Element(children...)
}

func Globe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeInnerSVG).
		Element(children...)
}

func Gps(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gpsInnerSVG).
		Element(children...)
}

func Grab(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(grabInnerSVG).
		Element(children...)
}

func GraphBar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphBarInnerSVG).
		Element(children...)
}

func GraphBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphBoxInnerSVG).
		Element(children...)
}

func GraphIncrease(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphIncreaseInnerSVG).
		Element(children...)
}

func Grid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridInnerSVG).
		Element(children...)
}

func GridCircles(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridCirclesInnerSVG).
		Element(children...)
}

func GridCirclesAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridCirclesAddInnerSVG).
		Element(children...)
}

func GridSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridSmallInnerSVG).
		Element(children...)
}

func GridSquares(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridSquaresInnerSVG).
		Element(children...)
}

func GridSquaresAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridSquaresAddInnerSVG).
		Element(children...)
}

func Hand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(handInnerSVG).
		Element(children...)
}

func Harddrive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(harddriveInnerSVG).
		Element(children...)
}

func Hash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hashInnerSVG).
		Element(children...)
}

func Heart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartInnerSVG).
		Element(children...)
}

func HeartRate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartRateInnerSVG).
		Element(children...)
}

func HeartRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartRemoveInnerSVG).
		Element(children...)
}

func Height(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heightInnerSVG).
		Element(children...)
}

func Hierarchy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hierarchyInnerSVG).
		Element(children...)
}

func Home(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeAltInnerSVG).
		Element(children...)
}

func HomeCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeCheckInnerSVG).
		Element(children...)
}

func HomeDoor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeDoorInnerSVG).
		Element(children...)
}

func Import(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(importInnerSVG).
		Element(children...)
}

func Inbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxInnerSVG).
		Element(children...)
}

func InboxAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxAltInnerSVG).
		Element(children...)
}

func InfoCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoCircleInnerSVG).
		Element(children...)
}

func IphoneLandscape(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iphoneLandscapeInnerSVG).
		Element(children...)
}

func IphonePortrait(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iphonePortraitInnerSVG).
		Element(children...)
}

func JumpBackward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jumpBackwardInnerSVG).
		Element(children...)
}

func JumpForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jumpForwardInnerSVG).
		Element(children...)
}

func JumpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jumpLeftInnerSVG).
		Element(children...)
}

func JumpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jumpRightInnerSVG).
		Element(children...)
}

func Keyboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyboardInnerSVG).
		Element(children...)
}

func Laptop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laptopInnerSVG).
		Element(children...)
}

func Lightbulb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightbulbInnerSVG).
		Element(children...)
}

func LightbulbOn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightbulbOnInnerSVG).
		Element(children...)
}

func Lightning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightningInnerSVG).
		Element(children...)
}

func LightningAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightningAltInnerSVG).
		Element(children...)
}

func Lineweight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineweightInnerSVG).
		Element(children...)
}

func Link(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkInnerSVG).
		Element(children...)
}

func LinkAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkAltInnerSVG).
		Element(children...)
}

func LinkBroken(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkBrokenInnerSVG).
		Element(children...)
}

func LinkHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkHorizontalInnerSVG).
		Element(children...)
}

func LinkVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkVerticalInnerSVG).
		Element(children...)
}

func List(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listInnerSVG).
		Element(children...)
}

func ListAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listAddInnerSVG).
		Element(children...)
}

func ListNumbered(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listNumberedInnerSVG).
		Element(children...)
}

func Loader(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loaderInnerSVG).
		Element(children...)
}

func Location(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationInnerSVG).
		Element(children...)
}

func Lock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockInnerSVG).
		Element(children...)
}

func LockOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockOpenInnerSVG).
		Element(children...)
}

func Mail(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailInnerSVG).
		Element(children...)
}

func MailAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailAddInnerSVG).
		Element(children...)
}

func MailDelete(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailDeleteInnerSVG).
		Element(children...)
}

func MailMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailMinusInnerSVG).
		Element(children...)
}

func MailNew(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailNewInnerSVG).
		Element(children...)
}

func MailOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailOpenInnerSVG).
		Element(children...)
}

func MailRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailRemoveInnerSVG).
		Element(children...)
}

func Marquee(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(marqueeInnerSVG).
		Element(children...)
}

func Maximise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maximiseInnerSVG).
		Element(children...)
}

func MenuHamburger(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuHamburgerInnerSVG).
		Element(children...)
}

func MenuHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuHorizontalInnerSVG).
		Element(children...)
}

func MenuVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuVerticalInnerSVG).
		Element(children...)
}

func Message(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageInnerSVG).
		Element(children...)
}

func MessageWriting(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageWritingInnerSVG).
		Element(children...)
}

func Microphone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneInnerSVG).
		Element(children...)
}

func MicrophoneDisabled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneDisabledInnerSVG).
		Element(children...)
}

func MicrophoneMuted(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneMutedInnerSVG).
		Element(children...)
}

func Midpoint(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(midpointInnerSVG).
		Element(children...)
}

func MiniPlayer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(miniPlayerInnerSVG).
		Element(children...)
}

func Minimise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minimiseInnerSVG).
		Element(children...)
}

func Minus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircleInnerSVG).
		Element(children...)
}

func Moon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveInnerSVG).
		Element(children...)
}

func Newspaper(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(newspaperInnerSVG).
		Element(children...)
}

func NoSign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noSignInnerSVG).
		Element(children...)
}

func Notebook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notebookInnerSVG).
		Element(children...)
}

func Notification(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notificationInnerSVG).
		Element(children...)
}

func Nut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nutInnerSVG).
		Element(children...)
}

func Pages(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pagesInnerSVG).
		Element(children...)
}

func PanelBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panelBottomInnerSVG).
		Element(children...)
}

func PanelCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panelCenterInnerSVG).
		Element(children...)
}

func PanelLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panelLeftInnerSVG).
		Element(children...)
}

func PanelRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panelRightInnerSVG).
		Element(children...)
}

func PanelSectioned(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panelSectionedInnerSVG).
		Element(children...)
}

func PanelTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panelTopInnerSVG).
		Element(children...)
}

func Paper(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperInnerSVG).
		Element(children...)
}

func PaperFolded(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperFoldedInnerSVG).
		Element(children...)
}

func PaperPlane(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperPlaneInnerSVG).
		Element(children...)
}

func PaperPlaneAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperPlaneAltInnerSVG).
		Element(children...)
}

func Paperclip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperclipInnerSVG).
		Element(children...)
}

func ParagraphCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphCenterInnerSVG).
		Element(children...)
}

func ParagraphEnd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphEndInnerSVG).
		Element(children...)
}

func ParagraphLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphLeftInnerSVG).
		Element(children...)
}

func ParagraphRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphRightInnerSVG).
		Element(children...)
}

func ParagraphStart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphStartInnerSVG).
		Element(children...)
}

func Pen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(penInnerSVG).
		Element(children...)
}

func PhoneLandscape(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneLandscapeInnerSVG).
		Element(children...)
}

func PhonePortrait(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phonePortraitInnerSVG).
		Element(children...)
}

func Picture(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pictureInnerSVG).
		Element(children...)
}

func PieHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieHalfInnerSVG).
		Element(children...)
}

func PieQuarter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieQuarterInnerSVG).
		Element(children...)
}

func PieThird(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieThirdInnerSVG).
		Element(children...)
}

func Pill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pillInnerSVG).
		Element(children...)
}

func PlayButton(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playButtonInnerSVG).
		Element(children...)
}

func Plus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircleInnerSVG).
		Element(children...)
}

func Postcard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(postcardInnerSVG).
		Element(children...)
}

func Printer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printerInnerSVG).
		Element(children...)
}

func Projector(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(projectorInnerSVG).
		Element(children...)
}

func PullDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pullDownInnerSVG).
		Element(children...)
}

func PullLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pullLeftInnerSVG).
		Element(children...)
}

func PullRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pullRightInnerSVG).
		Element(children...)
}

func PullUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pullUpInnerSVG).
		Element(children...)
}

func PushDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pushDownInnerSVG).
		Element(children...)
}

func PushLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pushLeftInnerSVG).
		Element(children...)
}

func PushRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pushRightInnerSVG).
		Element(children...)
}

func PushUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pushUpInnerSVG).
		Element(children...)
}

func QuestionCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionCircleInnerSVG).
		Element(children...)
}

func RadioOn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radioOnInnerSVG).
		Element(children...)
}

func Receipt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(receiptInnerSVG).
		Element(children...)
}

func Record(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(recordInnerSVG).
		Element(children...)
}

func Redo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redoInnerSVG).
		Element(children...)
}

func Refresh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshInnerSVG).
		Element(children...)
}

func RefreshAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshAltInnerSVG).
		Element(children...)
}

func Replicate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(replicateInnerSVG).
		Element(children...)
}

func ReplicateAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(replicateAltInnerSVG).
		Element(children...)
}

func Reset(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resetInnerSVG).
		Element(children...)
}

func ResetAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resetAltInnerSVG).
		Element(children...)
}

func ResetForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resetForwardInnerSVG).
		Element(children...)
}

func ResetHard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resetHardInnerSVG).
		Element(children...)
}

func ResetTemporary(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resetTemporaryInnerSVG).
		Element(children...)
}

func Retweet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(retweetInnerSVG).
		Element(children...)
}

func Reuse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reuseInnerSVG).
		Element(children...)
}

func Reverse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reverseInnerSVG).
		Element(children...)
}

func ReverseAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reverseAltInnerSVG).
		Element(children...)
}

func Revert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(revertInnerSVG).
		Element(children...)
}

func Rocket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rocketInnerSVG).
		Element(children...)
}

func Ruler(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rulerInnerSVG).
		Element(children...)
}

func Scale(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scaleInnerSVG).
		Element(children...)
}

func ScaleContract(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scaleContractInnerSVG).
		Element(children...)
}

func ScaleExtend(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scaleExtendInnerSVG).
		Element(children...)
}

func Scalpel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scalpelInnerSVG).
		Element(children...)
}

func Search(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchInnerSVG).
		Element(children...)
}

func Server(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareInnerSVG).
		Element(children...)
}

func ShareAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareAltInnerSVG).
		Element(children...)
}

func Shuffle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shuffleInnerSVG).
		Element(children...)
}

func SideMenu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sideMenuInnerSVG).
		Element(children...)
}

func SignalFull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signalFullInnerSVG).
		Element(children...)
}

func SignalLow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signalLowInnerSVG).
		Element(children...)
}

func SignalMedium(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signalMediumInnerSVG).
		Element(children...)
}

func SignalNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signalNoneInnerSVG).
		Element(children...)
}

func SlashBackward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slashBackwardInnerSVG).
		Element(children...)
}

func SlashForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slashForwardInnerSVG).
		Element(children...)
}

func Sliders(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slidersInnerSVG).
		Element(children...)
}

func Sort(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortInnerSVG).
		Element(children...)
}

func SortAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortAltInnerSVG).
		Element(children...)
}

func Speaker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speakerInnerSVG).
		Element(children...)
}

func SpeechBubble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speechBubbleInnerSVG).
		Element(children...)
}

func SpeechTyping(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speechTypingInnerSVG).
		Element(children...)
}

func Split(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splitInnerSVG).
		Element(children...)
}

func SplitThree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splitThreeInnerSVG).
		Element(children...)
}

func Star(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunInnerSVG).
		Element(children...)
}

func Support(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(supportInnerSVG).
		Element(children...)
}

func Swap(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swapInnerSVG).
		Element(children...)
}

func Switch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(switchInnerSVG).
		Element(children...)
}

func Table(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableInnerSVG).
		Element(children...)
}

func TableHeader(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableHeaderInnerSVG).
		Element(children...)
}

func Tag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagInnerSVG).
		Element(children...)
}

func TagMilestone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagMilestoneInnerSVG).
		Element(children...)
}

func Tags(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalInnerSVG).
		Element(children...)
}

func Thread(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threadInnerSVG).
		Element(children...)
}

func ThumbsDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbsUpInnerSVG).
		Element(children...)
}

func Ticket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ticketInnerSVG).
		Element(children...)
}

func Timeline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timelineInnerSVG).
		Element(children...)
}

func Todo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(todoInnerSVG).
		Element(children...)
}

func Toggle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleInnerSVG).
		Element(children...)
}

func Toggles(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(togglesInnerSVG).
		Element(children...)
}

func Translate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashInnerSVG).
		Element(children...)
}

func TrashAlt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashAltInnerSVG).
		Element(children...)
}

func Trophy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trophyInnerSVG).
		Element(children...)
}

func TvMode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tvModeInnerSVG).
		Element(children...)
}

func Unarchive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unarchiveInnerSVG).
		Element(children...)
}

func Undo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(undoInnerSVG).
		Element(children...)
}

func UndoHistory(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(undoHistoryInnerSVG).
		Element(children...)
}

func UnlinkHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlinkHorizontalInnerSVG).
		Element(children...)
}

func UnlinkVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlinkVerticalInnerSVG).
		Element(children...)
}

func Upload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadAltInnerSVG).
		Element(children...)
}

func Upward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upwardInnerSVG).
		Element(children...)
}

func User(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userInnerSVG).
		Element(children...)
}

func UserAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userAddInnerSVG).
		Element(children...)
}

func UserCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userCircleInnerSVG).
		Element(children...)
}

func UserMale(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userMaleInnerSVG).
		Element(children...)
}

func UserMaleCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userMaleCircleInnerSVG).
		Element(children...)
}

func UserRemove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userRemoveInnerSVG).
		Element(children...)
}

func Users(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usersInnerSVG).
		Element(children...)
}

func Venn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vennInnerSVG).
		Element(children...)
}

func Version(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(versionInnerSVG).
		Element(children...)
}

func Versions(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(versionsInnerSVG).
		Element(children...)
}

func Video(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(videoInnerSVG).
		Element(children...)
}

func VolumeAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeAddInnerSVG).
		Element(children...)
}

func VolumeDisabled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeDisabledInnerSVG).
		Element(children...)
}

func VolumeHigh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeLowInnerSVG).
		Element(children...)
}

func VolumeMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeMinusInnerSVG).
		Element(children...)
}

func VolumeMuted(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeMutedInnerSVG).
		Element(children...)
}

func VolumeZero(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeZeroInnerSVG).
		Element(children...)
}

func Wallet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningCircleInnerSVG).
		Element(children...)
}

func WarningHex(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningHexInnerSVG).
		Element(children...)
}

func WarningTriangle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningTriangleInnerSVG).
		Element(children...)
}

func Waves(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wavesInnerSVG).
		Element(children...)
}

func Width(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(widthInnerSVG).
		Element(children...)
}

func Wifi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiErrorInnerSVG).
		Element(children...)
}

func WifiNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiNoneInnerSVG).
		Element(children...)
}

func Window(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowInnerSVG).
		Element(children...)
}

func WindowCollapseLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowCollapseLeftInnerSVG).
		Element(children...)
}

func WindowCollapseRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowCollapseRightInnerSVG).
		Element(children...)
}

func WindowContent(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowContentInnerSVG).
		Element(children...)
}

func WrapBack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wrapBackInnerSVG).
		Element(children...)
}

func WrapForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wrapForwardInnerSVG).
		Element(children...)
}

func Write(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(writeInnerSVG).
		Element(children...)
}

func ZoomCancel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomCancelInnerSVG).
		Element(children...)
}

func ZoomIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
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
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomOutInnerSVG).
		Element(children...)
}

func ZoomReset(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 21 21"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomResetInnerSVG).
		Element(children...)
}

func ByName(name string) (*engine.UberElement, error) {
	switch name {
	case "airplay":
		return Airplay(), nil
	case "alarm-clock":
		return AlarmClock(), nil
	case "align-horizontal":
		return AlignHorizontal(), nil
	case "align-vertical":
		return AlignVertical(), nil
	case "angle":
		return Angle(), nil
	case "archive":
		return Archive(), nil
	case "arrow-bottom-left":
		return ArrowBottomLeft(), nil
	case "arrow-bottom-right":
		return ArrowBottomRight(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-down-circle":
		return ArrowDownCircle(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-left-circle":
		return ArrowLeftCircle(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-right-circle":
		return ArrowRightCircle(), nil
	case "arrow-top-left":
		return ArrowTopLeft(), nil
	case "arrow-top-right":
		return ArrowTopRight(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrow-up-circle":
		return ArrowUpCircle(), nil
	case "audio-wave":
		return AudioWave(), nil
	case "backspace":
		return Backspace(), nil
	case "backward":
		return Backward(), nil
	case "bag":
		return Bag(), nil
	case "battery-charging":
		return BatteryCharging(), nil
	case "battery-empty":
		return BatteryEmpty(), nil
	case "battery-full":
		return BatteryFull(), nil
	case "battery-half":
		return BatteryHalf(), nil
	case "battery-low":
		return BatteryLow(), nil
	case "battery-75":
		return BatterySeventyFive(), nil
	case "bell":
		return Bell(), nil
	case "bell-disabled":
		return BellDisabled(), nil
	case "bell-ringing":
		return BellRinging(), nil
	case "bell-snooze":
		return BellSnooze(), nil
	case "bluetooth":
		return Bluetooth(), nil
	case "book":
		return Book(), nil
	case "book-closed":
		return BookClosed(), nil
	case "book-text":
		return BookText(), nil
	case "bookmark":
		return Bookmark(), nil
	case "bookmark-book":
		return BookmarkBook(), nil
	case "box":
		return Box(), nil
	case "box-add":
		return BoxAdd(), nil
	case "box-download":
		return BoxDownload(), nil
	case "box-open":
		return BoxOpen(), nil
	case "box-remove":
		return BoxRemove(), nil
	case "boxes":
		return Boxes(), nil
	case "branch":
		return Branch(), nil
	case "briefcase":
		return Briefcase(), nil
	case "browser":
		return Browser(), nil
	case "browser-alt":
		return BrowserAlt(), nil
	case "button-add":
		return ButtonAdd(), nil
	case "button-minus":
		return ButtonMinus(), nil
	case "calculator":
		return Calculator(), nil
	case "calendar":
		return Calendar(), nil
	case "calendar-add":
		return CalendarAdd(), nil
	case "calendar-date":
		return CalendarDate(), nil
	case "calendar-day":
		return CalendarDay(), nil
	case "calendar-days":
		return CalendarDays(), nil
	case "calendar-last-day":
		return CalendarLastDay(), nil
	case "calendar-month":
		return CalendarMonth(), nil
	case "calendar-move":
		return CalendarMove(), nil
	case "calendar-remove":
		return CalendarRemove(), nil
	case "calendar-split":
		return CalendarSplit(), nil
	case "calendar-week":
		return CalendarWeek(), nil
	case "camera":
		return Camera(), nil
	case "camera-alt":
		return CameraAlt(), nil
	case "camera-noflash":
		return CameraNoflash(), nil
	case "camera-noflash-alt":
		return CameraNoflashAlt(), nil
	case "capture":
		return Capture(), nil
	case "card-timeline":
		return CardTimeline(), nil
	case "card-view":
		return CardView(), nil
	case "carousel":
		return Carousel(), nil
	case "cart":
		return Cart(), nil
	case "cast":
		return Cast(), nil
	case "chain":
		return Chain(), nil
	case "chat-add":
		return ChatAdd(), nil
	case "check":
		return Check(), nil
	case "check-circle":
		return CheckCircle(), nil
	case "check-circle-outside":
		return CheckCircleOutside(), nil
	case "checkbox-checked":
		return CheckboxChecked(), nil
	case "checkbox-empty":
		return CheckboxEmpty(), nil
	case "chevron-close":
		return ChevronClose(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-down-circle":
		return ChevronDownCircle(), nil
	case "chevron-down-double":
		return ChevronDownDouble(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-left-circle":
		return ChevronLeftCircle(), nil
	case "chevron-left-double":
		return ChevronLeftDouble(), nil
	case "chevron-open":
		return ChevronOpen(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-right-circle":
		return ChevronRightCircle(), nil
	case "chevron-right-double":
		return ChevronRightDouble(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "chevron-up-circle":
		return ChevronUpCircle(), nil
	case "chevron-up-double":
		return ChevronUpDouble(), nil
	case "circle":
		return Circle(), nil
	case "circle-menu":
		return CircleMenu(), nil
	case "circle-split":
		return CircleSplit(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clipboard-add":
		return ClipboardAdd(), nil
	case "clipboard-check":
		return ClipboardCheck(), nil
	case "clipboard-copy":
		return ClipboardCopy(), nil
	case "clipboard-cross":
		return ClipboardCross(), nil
	case "clipboard-notes":
		return ClipboardNotes(), nil
	case "clipboard-remove":
		return ClipboardRemove(), nil
	case "clock":
		return Clock(), nil
	case "close":
		return Close(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-disconnect":
		return CloudDisconnect(), nil
	case "cloud-download":
		return CloudDownload(), nil
	case "cloud-download-alt":
		return CloudDownloadAlt(), nil
	case "cloud-upload":
		return CloudUpload(), nil
	case "cloud-upload-alt":
		return CloudUploadAlt(), nil
	case "code":
		return Code(), nil
	case "coffee":
		return Coffee(), nil
	case "coin":
		return Coin(), nil
	case "coins":
		return Coins(), nil
	case "compass":
		return Compass(), nil
	case "component-add":
		return ComponentAdd(), nil
	case "contacts":
		return Contacts(), nil
	case "contract":
		return Contract(), nil
	case "create":
		return Create(), nil
	case "credit-card":
		return CreditCard(), nil
	case "crop":
		return Crop(), nil
	case "cross":
		return Cross(), nil
	case "cross-circle":
		return CrossCircle(), nil
	case "crosshair":
		return Crosshair(), nil
	case "cube":
		return Cube(), nil
	case "cubes":
		return Cubes(), nil
	case "cylinder":
		return Cylinder(), nil
	case "database":
		return Database(), nil
	case "diamond":
		return Diamond(), nil
	case "directions":
		return Directions(), nil
	case "disc":
		return Disc(), nil
	case "display":
		return Display(), nil
	case "display-alt":
		return DisplayAlt(), nil
	case "document":
		return Document(), nil
	case "document-justified":
		return DocumentJustified(), nil
	case "document-list":
		return DocumentList(), nil
	case "document-stack":
		return DocumentStack(), nil
	case "document-words":
		return DocumentWords(), nil
	case "door":
		return Door(), nil
	case "door-alt":
		return DoorAlt(), nil
	case "download":
		return Download(), nil
	case "download-alt":
		return DownloadAlt(), nil
	case "downward":
		return Downward(), nil
	case "drag":
		return Drag(), nil
	case "drag-circle":
		return DragCircle(), nil
	case "drag-vertical":
		return DragVertical(), nil
	case "duplicate":
		return Duplicate(), nil
	case "duplicate-alt":
		return DuplicateAlt(), nil
	case "enter":
		return Enter(), nil
	case "enter-alt":
		return EnterAlt(), nil
	case "episodes":
		return Episodes(), nil
	case "exit-left":
		return ExitLeft(), nil
	case "exit-right":
		return ExitRight(), nil
	case "expand":
		return Expand(), nil
	case "expand-height":
		return ExpandHeight(), nil
	case "expand-width":
		return ExpandWidth(), nil
	case "external":
		return External(), nil
	case "eye":
		return Eye(), nil
	case "eye-closed":
		return EyeClosed(), nil
	case "eye-no":
		return EyeNo(), nil
	case "face-delighted":
		return FaceDelighted(), nil
	case "face-happy":
		return FaceHappy(), nil
	case "face-neutral":
		return FaceNeutral(), nil
	case "face-sad":
		return FaceSad(), nil
	case "file-download":
		return FileDownload(), nil
	case "file-upload":
		return FileUpload(), nil
	case "files-history":
		return FilesHistory(), nil
	case "files-multi":
		return FilesMulti(), nil
	case "files-stack":
		return FilesStack(), nil
	case "film":
		return Film(), nil
	case "filter":
		return Filter(), nil
	case "filter-circle":
		return FilterCircle(), nil
	case "filter-single":
		return FilterSingle(), nil
	case "filtering":
		return Filtering(), nil
	case "fingerprint":
		return Fingerprint(), nil
	case "flag":
		return Flag(), nil
	case "flame":
		return Flame(), nil
	case "flame-alt":
		return FlameAlt(), nil
	case "flip-view":
		return FlipView(), nil
	case "floppy":
		return Floppy(), nil
	case "folder-add":
		return FolderAdd(), nil
	case "folder-closed":
		return FolderClosed(), nil
	case "folder-minus":
		return FolderMinus(), nil
	case "folder-open":
		return FolderOpen(), nil
	case "fork-git":
		return ForkGit(), nil
	case "forward":
		return Forward(), nil
	case "frame":
		return Frame(), nil
	case "fullscreen":
		return Fullscreen(), nil
	case "funnel":
		return Funnel(), nil
	case "gauge":
		return Gauge(), nil
	case "gift":
		return Gift(), nil
	case "globe":
		return Globe(), nil
	case "gps":
		return Gps(), nil
	case "grab":
		return Grab(), nil
	case "graph-bar":
		return GraphBar(), nil
	case "graph-box":
		return GraphBox(), nil
	case "graph-increase":
		return GraphIncrease(), nil
	case "grid":
		return Grid(), nil
	case "grid-circles":
		return GridCircles(), nil
	case "grid-circles-add":
		return GridCirclesAdd(), nil
	case "grid-small":
		return GridSmall(), nil
	case "grid-squares":
		return GridSquares(), nil
	case "grid-squares-add":
		return GridSquaresAdd(), nil
	case "hand":
		return Hand(), nil
	case "harddrive":
		return Harddrive(), nil
	case "hash":
		return Hash(), nil
	case "heart":
		return Heart(), nil
	case "heart-rate":
		return HeartRate(), nil
	case "heart-remove":
		return HeartRemove(), nil
	case "height":
		return Height(), nil
	case "hierarchy":
		return Hierarchy(), nil
	case "home":
		return Home(), nil
	case "home-alt":
		return HomeAlt(), nil
	case "home-check":
		return HomeCheck(), nil
	case "home-door":
		return HomeDoor(), nil
	case "import":
		return Import(), nil
	case "inbox":
		return Inbox(), nil
	case "inbox-alt":
		return InboxAlt(), nil
	case "info-circle":
		return InfoCircle(), nil
	case "iphone-landscape":
		return IphoneLandscape(), nil
	case "iphone-portrait":
		return IphonePortrait(), nil
	case "jump-backward":
		return JumpBackward(), nil
	case "jump-forward":
		return JumpForward(), nil
	case "jump-left":
		return JumpLeft(), nil
	case "jump-right":
		return JumpRight(), nil
	case "keyboard":
		return Keyboard(), nil
	case "laptop":
		return Laptop(), nil
	case "lightbulb":
		return Lightbulb(), nil
	case "lightbulb-on":
		return LightbulbOn(), nil
	case "lightning":
		return Lightning(), nil
	case "lightning-alt":
		return LightningAlt(), nil
	case "lineweight":
		return Lineweight(), nil
	case "link":
		return Link(), nil
	case "link-alt":
		return LinkAlt(), nil
	case "link-broken":
		return LinkBroken(), nil
	case "link-horizontal":
		return LinkHorizontal(), nil
	case "link-vertical":
		return LinkVertical(), nil
	case "list":
		return List(), nil
	case "list-add":
		return ListAdd(), nil
	case "list-numbered":
		return ListNumbered(), nil
	case "loader":
		return Loader(), nil
	case "location":
		return Location(), nil
	case "lock":
		return Lock(), nil
	case "lock-open":
		return LockOpen(), nil
	case "mail":
		return Mail(), nil
	case "mail-add":
		return MailAdd(), nil
	case "mail-delete":
		return MailDelete(), nil
	case "mail-minus":
		return MailMinus(), nil
	case "mail-new":
		return MailNew(), nil
	case "mail-open":
		return MailOpen(), nil
	case "mail-remove":
		return MailRemove(), nil
	case "marquee":
		return Marquee(), nil
	case "maximise":
		return Maximise(), nil
	case "menu-hamburger":
		return MenuHamburger(), nil
	case "menu-horizontal":
		return MenuHorizontal(), nil
	case "menu-vertical":
		return MenuVertical(), nil
	case "message":
		return Message(), nil
	case "message-writing":
		return MessageWriting(), nil
	case "microphone":
		return Microphone(), nil
	case "microphone-disabled":
		return MicrophoneDisabled(), nil
	case "microphone-muted":
		return MicrophoneMuted(), nil
	case "midpoint":
		return Midpoint(), nil
	case "mini-player":
		return MiniPlayer(), nil
	case "minimise":
		return Minimise(), nil
	case "minus":
		return Minus(), nil
	case "minus-circle":
		return MinusCircle(), nil
	case "moon":
		return Moon(), nil
	case "move":
		return Move(), nil
	case "newspaper":
		return Newspaper(), nil
	case "no-sign":
		return NoSign(), nil
	case "notebook":
		return Notebook(), nil
	case "notification":
		return Notification(), nil
	case "nut":
		return Nut(), nil
	case "pages":
		return Pages(), nil
	case "panel-bottom":
		return PanelBottom(), nil
	case "panel-center":
		return PanelCenter(), nil
	case "panel-left":
		return PanelLeft(), nil
	case "panel-right":
		return PanelRight(), nil
	case "panel-sectioned":
		return PanelSectioned(), nil
	case "panel-top":
		return PanelTop(), nil
	case "paper":
		return Paper(), nil
	case "paper-folded":
		return PaperFolded(), nil
	case "paper-plane":
		return PaperPlane(), nil
	case "paper-plane-alt":
		return PaperPlaneAlt(), nil
	case "paperclip":
		return Paperclip(), nil
	case "paragraph-center":
		return ParagraphCenter(), nil
	case "paragraph-end":
		return ParagraphEnd(), nil
	case "paragraph-left":
		return ParagraphLeft(), nil
	case "paragraph-right":
		return ParagraphRight(), nil
	case "paragraph-start":
		return ParagraphStart(), nil
	case "pen":
		return Pen(), nil
	case "phone-landscape":
		return PhoneLandscape(), nil
	case "phone-portrait":
		return PhonePortrait(), nil
	case "picture":
		return Picture(), nil
	case "pie-half":
		return PieHalf(), nil
	case "pie-quarter":
		return PieQuarter(), nil
	case "pie-third":
		return PieThird(), nil
	case "pill":
		return Pill(), nil
	case "play-button":
		return PlayButton(), nil
	case "plus":
		return Plus(), nil
	case "plus-circle":
		return PlusCircle(), nil
	case "postcard":
		return Postcard(), nil
	case "printer":
		return Printer(), nil
	case "projector":
		return Projector(), nil
	case "pull-down":
		return PullDown(), nil
	case "pull-left":
		return PullLeft(), nil
	case "pull-right":
		return PullRight(), nil
	case "pull-up":
		return PullUp(), nil
	case "push-down":
		return PushDown(), nil
	case "push-left":
		return PushLeft(), nil
	case "push-right":
		return PushRight(), nil
	case "push-up":
		return PushUp(), nil
	case "question-circle":
		return QuestionCircle(), nil
	case "radio-on":
		return RadioOn(), nil
	case "receipt":
		return Receipt(), nil
	case "record":
		return Record(), nil
	case "redo":
		return Redo(), nil
	case "refresh":
		return Refresh(), nil
	case "refresh-alt":
		return RefreshAlt(), nil
	case "replicate":
		return Replicate(), nil
	case "replicate-alt":
		return ReplicateAlt(), nil
	case "reset":
		return Reset(), nil
	case "reset-alt":
		return ResetAlt(), nil
	case "reset-forward":
		return ResetForward(), nil
	case "reset-hard":
		return ResetHard(), nil
	case "reset-temporary":
		return ResetTemporary(), nil
	case "retweet":
		return Retweet(), nil
	case "reuse":
		return Reuse(), nil
	case "reverse":
		return Reverse(), nil
	case "reverse-alt":
		return ReverseAlt(), nil
	case "revert":
		return Revert(), nil
	case "rocket":
		return Rocket(), nil
	case "ruler":
		return Ruler(), nil
	case "scale":
		return Scale(), nil
	case "scale-contract":
		return ScaleContract(), nil
	case "scale-extend":
		return ScaleExtend(), nil
	case "scalpel":
		return Scalpel(), nil
	case "search":
		return Search(), nil
	case "server":
		return Server(), nil
	case "settings":
		return Settings(), nil
	case "share":
		return Share(), nil
	case "share-alt":
		return ShareAlt(), nil
	case "shuffle":
		return Shuffle(), nil
	case "side-menu":
		return SideMenu(), nil
	case "signal-full":
		return SignalFull(), nil
	case "signal-low":
		return SignalLow(), nil
	case "signal-medium":
		return SignalMedium(), nil
	case "signal-none":
		return SignalNone(), nil
	case "slash-backward":
		return SlashBackward(), nil
	case "slash-forward":
		return SlashForward(), nil
	case "sliders":
		return Sliders(), nil
	case "sort":
		return Sort(), nil
	case "sort-alt":
		return SortAlt(), nil
	case "speaker":
		return Speaker(), nil
	case "speech-bubble":
		return SpeechBubble(), nil
	case "speech-typing":
		return SpeechTyping(), nil
	case "split":
		return Split(), nil
	case "split-three":
		return SplitThree(), nil
	case "star":
		return Star(), nil
	case "sun":
		return Sun(), nil
	case "support":
		return Support(), nil
	case "swap":
		return Swap(), nil
	case "switch":
		return Switch(), nil
	case "table":
		return Table(), nil
	case "table-header":
		return TableHeader(), nil
	case "tag":
		return Tag(), nil
	case "tag-milestone":
		return TagMilestone(), nil
	case "tags":
		return Tags(), nil
	case "target":
		return Target(), nil
	case "terminal":
		return Terminal(), nil
	case "thread":
		return Thread(), nil
	case "thumbs-down":
		return ThumbsDown(), nil
	case "thumbs-up":
		return ThumbsUp(), nil
	case "ticket":
		return Ticket(), nil
	case "timeline":
		return Timeline(), nil
	case "todo":
		return Todo(), nil
	case "toggle":
		return Toggle(), nil
	case "toggles":
		return Toggles(), nil
	case "translate":
		return Translate(), nil
	case "trash":
		return Trash(), nil
	case "trash-alt":
		return TrashAlt(), nil
	case "trophy":
		return Trophy(), nil
	case "tv-mode":
		return TvMode(), nil
	case "unarchive":
		return Unarchive(), nil
	case "undo":
		return Undo(), nil
	case "undo-history":
		return UndoHistory(), nil
	case "unlink-horizontal":
		return UnlinkHorizontal(), nil
	case "unlink-vertical":
		return UnlinkVertical(), nil
	case "upload":
		return Upload(), nil
	case "upload-alt":
		return UploadAlt(), nil
	case "upward":
		return Upward(), nil
	case "user":
		return User(), nil
	case "user-add":
		return UserAdd(), nil
	case "user-circle":
		return UserCircle(), nil
	case "user-male":
		return UserMale(), nil
	case "user-male-circle":
		return UserMaleCircle(), nil
	case "user-remove":
		return UserRemove(), nil
	case "users":
		return Users(), nil
	case "venn":
		return Venn(), nil
	case "version":
		return Version(), nil
	case "versions":
		return Versions(), nil
	case "video":
		return Video(), nil
	case "volume-add":
		return VolumeAdd(), nil
	case "volume-disabled":
		return VolumeDisabled(), nil
	case "volume-high":
		return VolumeHigh(), nil
	case "volume-low":
		return VolumeLow(), nil
	case "volume-minus":
		return VolumeMinus(), nil
	case "volume-muted":
		return VolumeMuted(), nil
	case "volume-0":
		return VolumeZero(), nil
	case "wallet":
		return Wallet(), nil
	case "warning-circle":
		return WarningCircle(), nil
	case "warning-hex":
		return WarningHex(), nil
	case "warning-triangle":
		return WarningTriangle(), nil
	case "waves":
		return Waves(), nil
	case "width":
		return Width(), nil
	case "wifi":
		return Wifi(), nil
	case "wifi-error":
		return WifiError(), nil
	case "wifi-none":
		return WifiNone(), nil
	case "window":
		return Window(), nil
	case "window-collapse-left":
		return WindowCollapseLeft(), nil
	case "window-collapse-right":
		return WindowCollapseRight(), nil
	case "window-content":
		return WindowContent(), nil
	case "wrap-back":
		return WrapBack(), nil
	case "wrap-forward":
		return WrapForward(), nil
	case "write":
		return Write(), nil
	case "zoom-cancel":
		return ZoomCancel(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	case "zoom-reset":
		return ZoomReset(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in system_uicons icon set", name)
	}
}
