package uim

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	adobeInnerSVG                = `<path fill="currentColor" d="M2 22.041a.998.998 0 0 1-1-1V2.959a1 1 0 0 1 1-1h7.425a1 1 0 0 1 .925 1.38L2.925 21.42a1 1 0 0 1-.925.62Zm14.244 0H13.63a1 1 0 0 1-.891-.546l-1.522-2.99H8.963a1 1 0 0 1-.928-1.372L11.094 9.5a1 1 0 0 1 .928-.628h.01a1 1 0 0 1 .926.646l4.221 11.168a1 1 0 0 1-.935 1.354ZM22 21.809a.999.999 0 0 1-.921-.613L13.56 3.346a1 1 0 0 1 .921-1.387H22a1 1 0 0 1 1 1v17.85a1 1 0 0 1-.802.98a1.049 1.049 0 0 1-.198.02Z"/>`
	adobeAltInnerSVG             = `<path fill="currentColor" d="M14.483 2.959L22 20.809V2.959ZM2 2.959V21.04L9.425 2.96Zm7.069 14.324h2.784l1.748 3.433h2.537l-4.1-10.843Z"/>`
	airplayInnerSVG              = `<rect width="20" height="15" x="2" y="3" fill="currentColor" opacity=".5" rx="3"/><path fill="currentColor" d="M16 21H8a1 1 0 0 1-.832-1.555l4-6a1.038 1.038 0 0 1 1.664 0l4 6A1 1 0 0 1 16 21Z"/>`
	alignInnerSVG                = `<path fill="currentColor" d="M21 7H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8 4H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm8 8H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8-4H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M19 14.666a1 1 0 0 1-1-1v-3.333a1 1 0 0 1 1.64-.768l2 1.667a1 1 0 0 1 0 1.536l-2 1.667a.999.999 0 0 1-.64.231Z"/>`
	alignAltInnerSVG             = `<path fill="currentColor" d="M10 5H7a1 1 0 0 1 0-2h3a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h7a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h7a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h7a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h7a1 1 0 0 1 0 2zM21 5h-7a1 1 0 0 1 0-2h7a1 1 0 0 1 0 2zm0 4h-7a1 1 0 0 1 0-2h7a1 1 0 0 1 0 2zm0 4h-7a1 1 0 0 1 0-2h7a1 1 0 0 1 0 2zm0 4h-7a1 1 0 0 1 0-2h7a1 1 0 0 1 0 2zm-4 4h-3a1 1 0 0 1 0-2h3a1 1 0 0 1 0 2z"/>`
	alignCenterInnerSVG          = `<path fill="currentColor" d="M21 7H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-4 4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm4 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-4 4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z"/>`
	alignCenterJustifyInnerSVG   = `<path fill="currentColor" d="M21 5H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-4 4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z"/>`
	alignJustifyInnerSVG         = `<path fill="currentColor" d="M21 7H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2z"/>`
	alignLeftInnerSVG            = `<path fill="currentColor" d="M21 7H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-4 4H3a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2zm4 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-4 4H3a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2z"/>`
	alignLeftJustifyInnerSVG     = `<path fill="currentColor" d="M21 5H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-6 4H3a1 1 0 0 1 0-2h12a1 1 0 0 1 0 2z"/>`
	alignLetterRightInnerSVG     = `<path fill="currentColor" d="M21 4H10a1 1 0 0 1 0-2h11a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 6H11a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2z"/>`
	alignRightInnerSVG           = `<path fill="currentColor" d="M21 7H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H7a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H7a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2z"/>`
	alignRightJustifyInnerSVG    = `<path fill="currentColor" d="M21 5H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H11a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z"/>`
	amazonInnerSVG               = `<path fill="currentColor" d="M1.04 17.52q.1-.16.32-.02a21.308 21.308 0 0 0 10.88 2.9a21.524 21.524 0 0 0 7.74-1.46q.1-.04.29-.12t.27-.12a.356.356 0 0 1 .47.12q.17.24-.11.44q-.36.26-.92.6a14.99 14.99 0 0 1-3.84 1.58A16.175 16.175 0 0 1 12 22a16.017 16.017 0 0 1-5.9-1.09a16.246 16.246 0 0 1-4.98-3.07a.273.273 0 0 1-.12-.2a.215.215 0 0 1 .04-.12Zm6.02-5.7a4.036 4.036 0 0 1 .68-2.36A4.197 4.197 0 0 1 9.6 7.98a10.063 10.063 0 0 1 2.66-.66q.54-.06 1.76-.16v-.34a3.562 3.562 0 0 0-.28-1.72a1.5 1.5 0 0 0-1.32-.6h-.16a2.189 2.189 0 0 0-1.14.42a1.64 1.64 0 0 0-.62 1a.508.508 0 0 1-.4.46L7.8 6.1q-.34-.08-.34-.36a.587.587 0 0 1 .02-.14a3.834 3.834 0 0 1 1.67-2.64A6.268 6.268 0 0 1 12.26 2h.5a5.054 5.054 0 0 1 3.56 1.18a3.81 3.81 0 0 1 .37.43a3.875 3.875 0 0 1 .27.41a2.098 2.098 0 0 1 .18.52q.08.34.12.47a2.856 2.856 0 0 1 .06.56q.02.43.02.51v4.84a2.868 2.868 0 0 0 .15.95a2.475 2.475 0 0 0 .29.62q.14.19.46.61a.599.599 0 0 1 .12.32a.346.346 0 0 1-.16.28q-1.66 1.44-1.8 1.56a.557.557 0 0 1-.58.04q-.28-.24-.49-.46t-.3-.32a4.466 4.466 0 0 1-.29-.39q-.2-.29-.28-.39a4.91 4.91 0 0 1-2.2 1.52a6.038 6.038 0 0 1-1.68.2a3.505 3.505 0 0 1-2.53-.95a3.553 3.553 0 0 1-.99-2.69Zm3.44-.4a1.895 1.895 0 0 0 .39 1.25a1.294 1.294 0 0 0 1.05.47a1.022 1.022 0 0 0 .17-.02a1.022 1.022 0 0 1 .15-.02a2.033 2.033 0 0 0 1.3-1.08a3.13 3.13 0 0 0 .33-.83a3.8 3.8 0 0 0 .12-.73q.01-.28.01-.92v-.5a7.287 7.287 0 0 0-1.76.16a2.144 2.144 0 0 0-1.76 2.22Zm8.4 6.44a.626.626 0 0 1 .12-.16a3.14 3.14 0 0 1 .96-.46a6.52 6.52 0 0 1 1.48-.22a1.195 1.195 0 0 1 .38.02q.9.08 1.08.3a.655.655 0 0 1 .08.36v.14a4.56 4.56 0 0 1-.38 1.65a3.84 3.84 0 0 1-1.06 1.53a.302.302 0 0 1-.18.08a.177.177 0 0 1-.08-.02q-.12-.06-.06-.22a7.632 7.632 0 0 0 .74-2.42a.513.513 0 0 0-.08-.32q-.2-.24-1.12-.24q-.34 0-.8.04q-.5.06-.92.12a.232.232 0 0 1-.16-.04a.065.065 0 0 1-.02-.08a.153.153 0 0 1 .02-.06Z"/>`
	analysisInnerSVG             = `<path fill="currentColor" d="M14 16a1 1 0 0 1-.707-1.707l7-7a1 1 0 0 1 1.414 1.414l-7 7A.997.997 0 0 1 14 16zM3 17a1 1 0 0 1-.707-1.707l6-6a1 1 0 0 1 1.414 1.414l-6 6A.997.997 0 0 1 3 17z"/><path fill="currentColor" d="M14 16a.997.997 0 0 1-.707-.293l-5-5a1 1 0 0 1 1.414-1.414l5 5A1 1 0 0 1 14 16Z"/>`
	analyticsInnerSVG            = `<path fill="currentColor" d="M5 22a1 1 0 0 1-1-1v-8a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1zm5 0a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v18a1 1 0 0 1-1 1zm5 0a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v12a1 1 0 0 1-1 1zm5 0a1 1 0 0 1-1-1v-4a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1z"/>`
	anchorInnerSVG               = `<path fill="currentColor" d="M12 22a8.01 8.01 0 0 1-8-8a1 1 0 0 1 2 0a6 6 0 0 0 12 0a1 1 0 0 1 2 0a8.01 8.01 0 0 1-8 8Z"/><path fill="currentColor" d="M12 22a1 1 0 0 1-1-1V7a1 1 0 0 1 2 0v14a1 1 0 0 1-1 1zm-5-7H5a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2zm12 0h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2z"/><path fill="currentColor" d="M14 11h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2zm-2-3a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1z"/>`
	androidInnerSVG              = `<path fill="currentColor" d="m14.975 3.019l.96-1.732a.193.193 0 0 0-.338-.187l-.97 1.75a6.541 6.541 0 0 0-5.253 0l-.97-1.75a.193.193 0 0 0-.34.187l.96 1.732a5.546 5.546 0 0 0-3.092 4.876h12.137a5.546 5.546 0 0 0-3.094-4.876ZM5.931 17.17A1.467 1.467 0 0 0 7.4 18.64h.973v3a1.36 1.36 0 1 0 2.721 0v-3h1.814v3a1.36 1.36 0 1 0 2.72 0v-3h.974a1.467 1.467 0 0 0 1.468-1.468V8.375H5.93ZM4.064 8.14a1.362 1.362 0 0 0-1.36 1.361v5.669a1.36 1.36 0 1 0 2.72 0V9.502a1.362 1.362 0 0 0-1.36-1.36Zm15.872 0a1.362 1.362 0 0 0-1.36 1.361v5.669a1.36 1.36 0 1 0 2.72 0V9.502a1.362 1.362 0 0 0-1.36-1.36Z"/><circle cx="9.199" cy="5.168" r=".507" fill="currentColor" opacity=".5"/><circle cx="14.801" cy="5.168" r=".507" fill="currentColor" opacity=".5"/>`
	androidAltInnerSVG           = `<path fill="currentColor" d="m16.2 4.7l.7-1.2c.2-.5.1-1.1-.4-1.4c-.5-.2-1.1-.1-1.4.4l-.6 1.1c.6.2 1.2.6 1.7 1.1zm-8.4 0c.5-.5 1-.8 1.7-1.1l-.6-1.1c-.3-.5-.9-.6-1.4-.4s-.6.9-.4 1.4l.7 1.2zM6 9h12v2H6z"/><path fill="currentColor" d="M12 3C8.7 3 6 5.7 6 9h12c0-3.3-2.7-6-6-6zm9 6c-.6 0-1 .4-1 1v4c0 .6.4 1 1 1s1-.4 1-1v-4c0-.6-.4-1-1-1zM3 9c-.6 0-1 .4-1 1v4c0 .6.4 1 1 1s1-.4 1-1v-4c0-.6-.4-1-1-1zm3 8c0 .6.4 1 1 1h2v3c0 .6.4 1 1 1s1-.4 1-1v-3h2v3c0 .6.4 1 1 1s1-.4 1-1v-3h2c.6 0 1-.4 1-1v-6H6v6z" opacity=".5"/>`
	angleDoubleDownInnerSVG      = `<path fill="currentColor" d="M12 11.75a.997.997 0 0 1-.707-.293l-3-3a1 1 0 0 1 1.414-1.414L12 9.336l2.293-2.293a1 1 0 0 1 1.414 1.414l-3 3a.997.997 0 0 1-.707.293zm0 5.5a.997.997 0 0 1-.707-.293l-3-3a1 1 0 0 1 1.414-1.414L12 14.836l2.293-2.293a1 1 0 0 1 1.414 1.414l-3 3a.997.997 0 0 1-.707.293z"/>`
	angleDoubleLeftInnerSVG      = `<path fill="currentColor" d="M16.25 16a.997.997 0 0 1-.707-.293l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 0 1 1.414 1.414L14.664 12l2.293 2.293A1 1 0 0 1 16.25 16zm-5.5 0a.997.997 0 0 1-.707-.293l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 0 1 1.414 1.414L9.164 12l2.293 2.293A1 1 0 0 1 10.75 16z"/>`
	angleDoubleRightInnerSVG     = `<path fill="currentColor" d="M7.75 16a1 1 0 0 1-.707-1.707L9.336 12L7.043 9.707a1 1 0 0 1 1.414-1.414l3 3a1 1 0 0 1 0 1.414l-3 3A.997.997 0 0 1 7.75 16zm5.5 0a1 1 0 0 1-.707-1.707L14.836 12l-2.293-2.293a1 1 0 0 1 1.414-1.414l3 3a1 1 0 0 1 0 1.414l-3 3a.997.997 0 0 1-.707.293z"/>`
	angleDoubleUpInnerSVG        = `<path fill="currentColor" d="M15 17.25a.997.997 0 0 1-.707-.293L12 14.664l-2.293 2.293a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.414 0l3 3A1 1 0 0 1 15 17.25zm0-5.5a.997.997 0 0 1-.707-.293L12 9.164l-2.293 2.293a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.414 0l3 3A1 1 0 0 1 15 11.75z"/>`
	angleDownInnerSVG            = `<path fill="currentColor" d="M12 15.121a.997.997 0 0 1-.707-.293L7.05 10.586a1 1 0 0 1 1.414-1.414L12 12.707l3.536-3.535a1 1 0 0 1 1.414 1.414l-4.243 4.242a.997.997 0 0 1-.707.293Z"/>`
	angleLeftInnerSVG            = `<path fill="currentColor" d="M14.121 17.243a.997.997 0 0 1-.707-.293l-4.242-4.243a1 1 0 0 1 0-1.414l4.242-4.243a1 1 0 0 1 1.414 1.414L11.293 12l3.535 3.536a1 1 0 0 1-.707 1.707Z"/>`
	angleRightInnerSVG           = `<path fill="currentColor" d="M9.879 17.243a1 1 0 0 1-.707-1.707L12.707 12L9.172 8.464a1 1 0 0 1 1.414-1.414l4.242 4.243a1 1 0 0 1 0 1.414l-4.242 4.243a.997.997 0 0 1-.707.293Z"/>`
	angleRightBInnerSVG          = `<path fill="currentColor" d="M9.172 18.657a1 1 0 0 1-.707-1.707l4.95-4.95l-4.95-4.95a1 1 0 0 1 1.414-1.414l5.656 5.657a1 1 0 0 1 0 1.414L9.88 18.364a.997.997 0 0 1-.707.293Z"/>`
	angleUpInnerSVG              = `<path fill="currentColor" d="M16.243 15.121a.997.997 0 0 1-.707-.293L12 11.293l-3.536 3.535a1 1 0 0 1-1.414-1.414l4.243-4.242a1 1 0 0 1 1.414 0l4.243 4.242a1 1 0 0 1-.707 1.707Z"/>`
	appleInnerSVG                = `<path fill="currentColor" d="M17.458 12.625A4.523 4.523 0 0 1 19.62 8.82a4.672 4.672 0 0 0-3.658-1.984c-1.558-.158-3.04.917-3.829.917c-.79 0-2.009-.894-3.3-.87a4.896 4.896 0 0 0-4.14 2.508c-1.762 3.06-.449 7.593 1.268 10.076c.84 1.214 1.843 2.581 3.158 2.532c1.268-.05 1.746-.82 3.277-.82c1.531 0 1.962.82 3.3.795c1.364-.025 2.229-1.239 3.062-2.457a10.946 10.946 0 0 0 1.385-2.845a4.42 4.42 0 0 1-2.685-4.047Zm-2.517-7.432A4.405 4.405 0 0 0 15.981 2a4.483 4.483 0 0 0-2.945 1.516a4.186 4.186 0 0 0-1.061 3.093a3.71 3.71 0 0 0 2.966-1.416Z"/>`
	appleAltInnerSVG             = `<path fill="currentColor" d="M17.594 12.625a4.523 4.523 0 0 1 2.162-3.805a4.672 4.672 0 0 0-3.658-1.984c-1.558-.158-3.04.917-3.83.917c-.789 0-2.008-.894-3.3-.87A4.896 4.896 0 0 0 4.83 9.392c-1.763 3.06-.45 7.593 1.267 10.076c.84 1.214 1.843 2.581 3.158 2.532c1.268-.05 1.746-.82 3.277-.82c1.53 0 1.961.82 3.3.795c1.364-.025 2.229-1.239 3.062-2.457a10.946 10.946 0 0 0 1.384-2.845a4.42 4.42 0 0 1-2.684-4.047Z"/><path fill="currentColor" d="M15.216 5.04A5.56 5.56 0 0 0 16.536 1a5.672 5.672 0 0 0-3.73 1.92l-.02.047a5.56 5.56 0 0 0-1.32 4.04a5.672 5.672 0 0 0 3.73-1.92Z"/>`
	appsInnerSVG                 = `<rect width="9" height="9" x="2" y="2" fill="currentColor" rx="1"/><rect width="9" height="9" x="2" y="13" fill="currentColor" opacity=".5" rx="1"/><rect width="9" height="9" x="13" y="2" fill="currentColor" opacity=".5" rx="1"/><rect width="9" height="9" x="13" y="13" fill="currentColor" opacity=".5" rx="1"/>`
	arrowCircleDownInnerSVG      = `<path fill="currentColor" d="m7.293 12.707l4 4a1.004 1.004 0 0 0 1.414 0l4-4a1 1 0 0 0-1.414-1.414L13 13.586V8a1 1 0 0 0-2 0v5.586l-2.293-2.293a1 1 0 0 0-1.414 1.414Z"/><path fill="currentColor" d="M12 22A10 10 0 1 0 2 12a10.011 10.011 0 0 0 10 10ZM7.293 11.293a1 1 0 0 1 1.414 0L11 13.586V8a1 1 0 0 1 2 0v5.586l2.293-2.293a1 1 0 0 1 1.414 1.414l-4 4a1.004 1.004 0 0 1-1.414 0l-4-4a1 1 0 0 1 0-1.414Z" opacity=".5"/>`
	arrowCircleLeftInnerSVG      = `<path fill="currentColor" d="m11.293 7.293l-4 4a1.004 1.004 0 0 0 0 1.414l4 4a1 1 0 0 0 1.414-1.414L10.414 13H16a1 1 0 0 0 0-2h-5.586l2.293-2.293a1 1 0 0 0-1.414-1.414Z"/><path fill="currentColor" d="M2 12A10 10 0 1 0 12 2A10.011 10.011 0 0 0 2 12Zm10.707-4.707a1 1 0 0 1 0 1.414L10.414 11H16a1 1 0 0 1 0 2h-5.586l2.293 2.293a1 1 0 0 1-1.414 1.414l-4-4a1.004 1.004 0 0 1 0-1.414l4-4a1 1 0 0 1 1.414 0Z" opacity=".5"/>`
	arrowCircleRightInnerSVG     = `<path fill="currentColor" d="m12.707 16.707l4-4a1.004 1.004 0 0 0 0-1.414l-4-4a1 1 0 0 0-1.414 1.414L13.586 11H8a1 1 0 0 0 0 2h5.586l-2.293 2.293a1 1 0 0 0 1.414 1.414Z"/><path fill="currentColor" d="M22 12a10 10 0 1 0-10 10a10.011 10.011 0 0 0 10-10Zm-10.707 4.707a1 1 0 0 1 0-1.414L13.586 13H8a1 1 0 0 1 0-2h5.586l-2.293-2.293a1 1 0 0 1 1.414-1.414l4 4a1.004 1.004 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414 0Z" opacity=".5"/>`
	arrowCircleUpInnerSVG        = `<path fill="currentColor" d="m16.707 11.293l-4-4a1.004 1.004 0 0 0-1.414 0l-4 4a1 1 0 0 0 1.414 1.414L11 10.414V16a1 1 0 0 0 2 0v-5.586l2.293 2.293a1 1 0 0 0 1.414-1.414Z"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2Zm4.707 10.707a1 1 0 0 1-1.414 0L13 10.414V16a1 1 0 0 1-2 0v-5.586l-2.293 2.293a1 1 0 0 1-1.414-1.414l4-4a1.004 1.004 0 0 1 1.414 0l4 4a1 1 0 0 1 0 1.414Z" opacity=".5"/>`
	arrowDownLeftInnerSVG        = `<path fill="currentColor" d="M17 18H7a1 1 0 0 1-1-1V7a1 1 0 0 1 2 0v9h9a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M7 18a1 1 0 0 1-.707-1.707l10-10a1 1 0 0 1 1.414 1.414l-10 10A.997.997 0 0 1 7 18Z"/>`
	arrowDownRightInnerSVG       = `<path fill="currentColor" d="M17 18H7a1 1 0 0 1 0-2h9V7a1 1 0 0 1 2 0v10a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M17 18a.997.997 0 0 1-.707-.293l-10-10a1 1 0 0 1 1.414-1.414l10 10A1 1 0 0 1 17 18Z"/>`
	arrowUpLeftInnerSVG          = `<path fill="currentColor" d="M7 18a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h10a1 1 0 0 1 0 2H8v9a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M17 18a.997.997 0 0 1-.707-.293l-10-10a1 1 0 0 1 1.414-1.414l10 10A1 1 0 0 1 17 18Z"/>`
	arrowUpRightInnerSVG         = `<path fill="currentColor" d="M17 18a1 1 0 0 1-1-1V8H7a1 1 0 0 1 0-2h10a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M7 18a1 1 0 0 1-.707-1.707l10-10a1 1 0 0 1 1.414 1.414l-10 10A.997.997 0 0 1 7 18Z"/>`
	atInnerSVG                   = `<path fill="currentColor" d="M12 16.5a4.5 4.5 0 1 1 4.5-4.5a4.505 4.505 0 0 1-4.5 4.5Zm0-7a2.5 2.5 0 1 0 2.5 2.5A2.503 2.503 0 0 0 12 9.5Z"/><path fill="currentColor" d="M12 22a10 10 0 1 1 10-10v.75a3.75 3.75 0 0 1-7.5 0V8.5a1 1 0 0 1 2 0v4.25a1.75 1.75 0 0 0 3.5 0V12a8 8 0 1 0-4 6.928a1 1 0 1 1 1 1.733A10.02 10.02 0 0 1 12 22Z"/>`
	bagInnerSVG                  = `<path fill="currentColor" d="M19 6H5a3 3 0 0 0-3 3v2.72L8.837 14h6.326L22 11.72V9a3 3 0 0 0-3-3z" opacity=".5"/><path fill="currentColor" d="M10 6V5h4v1h2V5a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2v1h2zm-1.163 8L2 11.72V18a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3v-6.28L15.163 14H8.837z"/>`
	barsInnerSVG                 = `<path fill="currentColor" d="M21 13H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 5H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0-10H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2z"/>`
	batteryBoltInnerSVG          = `<path fill="currentColor" d="M13.988 12.059A1.023 1.023 0 0 0 13 11H9.732l1.444-2.5a1 1 0 1 0-1.733-1l-2.31 4A1.022 1.022 0 0 0 8 13h3.268l-1.444 2.5a1 1 0 0 0 1.733 1l2.31-4a1.267 1.267 0 0 0 .121-.441Z"/><path fill="currentColor" d="M17 7H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h6.69a1 1 0 0 1-.866-1.5l1.444-2.5H8a1.022 1.022 0 0 1-.866-1.5l2.31-4a1 1 0 1 1 1.732 1L9.732 11H13a1.023 1.023 0 0 1 .988 1.059a1.268 1.268 0 0 1-.122.441l-2.31 4a1 1 0 0 1-.867.5H17a2.002 2.002 0 0 0 2-2V9a2.002 2.002 0 0 0-2-2zm4 7a1 1 0 0 1-1-1v-2a1 1 0 0 1 2 0v2a1 1 0 0 1-1 1z" opacity=".5"/>`
	batteryEmptyInnerSVG         = `<path fill="currentColor" d="M21 14a1 1 0 0 1-1-1v-2a1 1 0 0 1 2 0v2a1 1 0 0 1-1 1zm-4 3H4a2.002 2.002 0 0 1-2-2V9a2.002 2.002 0 0 1 2-2h13a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2z"/>`
	behanceInnerSVG              = `<path fill="currentColor" d="M10.396 11.519h-.108l.24-.132a2.333 2.333 0 0 0 1.095-1.203a3.464 3.464 0 0 0 .145-1.697a2.815 2.815 0 0 0-1.732-2.297a4.728 4.728 0 0 0-1.925-.385H2v12.354h5.582a9.448 9.448 0 0 0 1.587-.132a3.465 3.465 0 0 0 2.659-1.973a3.741 3.741 0 0 0 .3-2.406a2.694 2.694 0 0 0-1.732-2.13ZM4.9 7.959h2.406a4.67 4.67 0 0 1 1.203.156a1.035 1.035 0 0 1 .794.926a1.9 1.9 0 0 1 0 .746a.999.999 0 0 1-.517.65a2.478 2.478 0 0 1-1.203.264H4.875Zm4.655 6.904a1.395 1.395 0 0 1-1.202 1.13a4.363 4.363 0 0 1-.794.085H4.875v-3.272h2.863a2.72 2.72 0 0 1 .902.156a1.311 1.311 0 0 1 .914 1.203a2.549 2.549 0 0 1 0 .698Zm12.439-.806c.012-.024.012-.048-.024-.024v-1.335a4.379 4.379 0 0 0-.41-1.769a3.729 3.729 0 0 0-1.924-1.852a5.137 5.137 0 0 0-2.947-.289a3.886 3.886 0 0 0-3.212 2.815a6.7 6.7 0 0 0-.144 3.465a3.5 3.5 0 0 0 .866 1.732a4.523 4.523 0 0 0 2.069 1.203a5.45 5.45 0 0 0 2.321.12a4.006 4.006 0 0 0 2.515-1.323a3.165 3.165 0 0 0 .71-1.203a.523.523 0 0 0 .084-.349H19.54a.12.12 0 0 0-.12.072a1.708 1.708 0 0 1-1.203.902a2.681 2.681 0 0 1-1.083 0a1.9 1.9 0 0 1-1.432-1.058a2.406 2.406 0 0 1-.204-.722v-.385Zm-6.352-1.732a4.515 4.515 0 0 1 .193-.566a1.732 1.732 0 0 1 1.576-1.082a2.287 2.287 0 0 1 1.046.144a1.78 1.78 0 0 1 1.13 1.444v.252l-3.945.036a2.099 2.099 0 0 1 0-.228Zm4.415-6.015h-5.004v1.444l5.004-.012Z"/>`
	behanceAltInnerSVG           = `<path fill="currentColor" d="M7.5 20H2a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h5.5a4.5 4.5 0 0 1 0 9ZM3 18h4.5a2.5 2.5 0 0 0 0-5H3Z"/><path fill="currentColor" d="M7 13H2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h5a4 4 0 0 1 0 8zm-4-2h4a2 2 0 0 0 0-4H3zm17-4h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2zm2 7h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2z"/><path fill="currentColor" d="M18 20a5.006 5.006 0 0 1-5-5v-2a5 5 0 1 1 10 0a1 1 0 0 1-2 0a3 3 0 0 0-6 0v2a2.998 2.998 0 0 0 5.9.754a1 1 0 0 1 1.94.492A4.985 4.985 0 0 1 18 20Z"/>`
	bingInnerSVG                 = `<path fill="currentColor" d="m10.129 8.596l1.735 4.328l2.77 1.29L19 16.247V11.7l-8.871-3.104z" opacity=".7"/><path fill="currentColor" d="M14.634 14.214L9 17.457V3.4L5 2v17.76L9 22l10-5.753V11.7l-4.366 2.514z"/>`
	bitcoinInnerSVG              = `<path fill="currentColor" d="M16.358 10.575c.197-1.33-.805-2.046-2.175-2.523l.445-1.802l-1.085-.274l-.433 1.755c-.285-.072-.578-.14-.869-.207l.436-1.766l-1.085-.274l-.444 1.802a36.28 36.28 0 0 1-.693-.164l.001-.006l-1.496-.378l-.288 1.172s.804.187.788.198c.313.04.537.324.505.638l-1.217 4.938a.394.394 0 0 1-.498.26c.011.017-.788-.198-.788-.198l-.539 1.256l1.412.356c.263.066.52.136.773.201l-.449 1.824l1.084.273l.445-1.804c.296.082.583.157.864.227l-.443 1.795l1.085.274l.449-1.82c1.85.354 3.241.211 3.827-1.48a1.933 1.933 0 0 0-.997-2.662a1.743 1.743 0 0 0 1.385-1.61zm-2.479 3.516c-.335 1.362-2.603.626-3.339.44l.596-2.414c.735.185 3.094.553 2.743 1.974zm.336-3.535c-.306 1.239-2.194.61-2.806.455l.54-2.19c.612.154 2.584.442 2.266 1.735z" opacity=".5"/><path fill="currentColor" d="m11.949 8.82l-.54 2.191c.612.154 2.5.784 2.806-.455c.318-1.293-1.654-1.581-2.266-1.736zm-.813 3.297l-.596 2.415c.736.185 3.004.921 3.34-.441c.35-1.421-2.009-1.789-2.744-1.974z"/><path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm4.358 8.575a1.743 1.743 0 0 1-1.385 1.611a1.933 1.933 0 0 1 .997 2.66c-.586 1.693-1.977 1.836-3.827 1.482l-.449 1.82l-1.085-.274l.443-1.795a35.27 35.27 0 0 1-.864-.227l-.445 1.804l-1.084-.273l.45-1.824c-.254-.065-.511-.135-.774-.201l-1.412-.356l.539-1.256s.8.215.788.199l.005.001a.394.394 0 0 0 .493-.262l1.217-4.938a.583.583 0 0 0-.505-.638c.016-.011-.788-.198-.788-.198l.288-1.172l1.496.378l-.001.006c.225.056.457.11.693.164l.444-1.802l1.085.274l-.436 1.766c.291.068.584.135.87.207l.432-1.755l1.085.274l-.445 1.802c1.37.477 2.372 1.193 2.175 2.523z"/>`
	bitcoinAltInnerSVG           = `<path fill="currentColor" d="M16.313 11.24A3.998 3.998 0 0 0 13 5V4a1 1 0 0 0-2 0v1H9V4a1 1 0 0 0-2 0v1H6a1 1 0 0 0 0 2h1v10H6a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h2v1a1 1 0 0 0 2 0v-1h2a3.99 3.99 0 0 0 1.313-7.76ZM15 9a2.002 2.002 0 0 1-2 2H9V7h4a2.002 2.002 0 0 1 2 2Zm0 8H9v-4h6a2 2 0 0 1 0 4Z"/>`
	blackberryInnerSVG           = `<path fill="currentColor" d="M7.188 13.352H5.27l.523-2.415h1.838c1.195 0 1.48.586 1.48 1.05c0 .657-.428 1.365-1.923 1.365zm.661-3.637H5.936l.52-2.416h1.84c1.194 0 1.479.587 1.479 1.051c0 .657-.428 1.365-1.926 1.365zm3.754 7.425H9.685l.522-2.418h1.84c1.193 0 1.476.589 1.476 1.054c0 .658-.425 1.364-1.92 1.364zm.708-3.788h-1.914l.521-2.415h1.84c1.193 0 1.476.586 1.476 1.05c0 .657-.425 1.365-1.923 1.365zm.664-3.637h-1.914l.523-2.416h1.836c1.196 0 1.48.587 1.48 1.051c0 .657-.427 1.365-1.925 1.365zm3.747 5.605h-1.915l.522-2.418h1.838c1.194 0 1.478.59 1.478 1.053c0 .657-.425 1.365-1.923 1.365zm.713-3.443H15.52l.522-2.416h1.84c1.194 0 1.476.587 1.476 1.052c0 .657-.425 1.364-1.923 1.364z" opacity=".5"/><path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zM7.188 13.352H5.27l.523-2.415h1.838c1.195 0 1.48.586 1.48 1.05c0 .657-.428 1.365-1.923 1.365zm.661-3.637H5.936l.52-2.416h1.84c1.194 0 1.479.587 1.479 1.051c0 .657-.428 1.365-1.926 1.365zm3.754 7.425H9.685l.522-2.418h1.84c1.193 0 1.476.589 1.476 1.054c0 .658-.425 1.364-1.92 1.364zm.708-3.788h-1.914l.521-2.415h1.84c1.193 0 1.476.586 1.476 1.05c0 .657-.425 1.365-1.923 1.365zm.664-3.637h-1.914l.523-2.416h1.836c1.196 0 1.48.587 1.48 1.051c0 .657-.427 1.365-1.925 1.365zm3.747 5.605h-1.915l.522-2.418h1.838c1.194 0 1.478.59 1.478 1.053c0 .657-.425 1.365-1.923 1.365zm.713-3.443H15.52l.522-2.416h1.84c1.194 0 1.476.587 1.476 1.052c0 .657-.425 1.364-1.923 1.364z"/>`
	bloggerInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="m18.333 10.922l-.11-.222l-.178-.14c-.235-.184-1.423.011-1.742-.278c-.228-.208-.263-.584-.33-1.09a3.154 3.154 0 0 0-.362-1.367a4.46 4.46 0 0 0-3.12-2.2H9.667a4.053 4.053 0 0 0-4.042 4.03v4.695a4.05 4.05 0 0 0 4.042 4.025h4.639a4.052 4.052 0 0 0 4.038-4.022l.02-2.675l.005-.579l-.036-.177zM9.718 8.917h2.24a.77.77 0 0 1 0 1.541h-2.24a.77.77 0 0 1 0-1.541zm4.55 6.125h-4.55a.75.75 0 1 1 0-1.5h4.55a.75.75 0 1 1 0 1.5z" clip-rule="evenodd" opacity=".5"/><path fill="currentColor" d="M14.267 13.542H9.673a.75.75 0 0 0 .045 1.5h4.587a.75.75 0 0 0-.038-1.5zm-4.549-3.084h2.24a.77.77 0 1 0 0-1.541h-2.24a.77.77 0 1 0 0 1.541z"/><path fill="currentColor" d="M19.5 2h-15A2.5 2.5 0 0 0 2 4.5v15A2.5 2.5 0 0 0 4.5 22h15a2.5 2.5 0 0 0 2.5-2.5v-15A2.5 2.5 0 0 0 19.5 2zm-1.136 9.678l-.02 2.675a4.051 4.051 0 0 1-4.038 4.022h-4.64a4.05 4.05 0 0 1-4.041-4.025V9.656a4.053 4.053 0 0 1 4.042-4.031h2.824a4.46 4.46 0 0 1 3.12 2.2c.222.422.346.89.362 1.367c.067.506.102.882.33 1.09c.32.29 1.507.094 1.742.278l.179.14l.109.222l.036.177l-.005.579z"/>`
	bloggerAltInnerSVG           = `<path fill="currentColor" d="M20 23H4a3.003 3.003 0 0 1-3-3V4a3.003 3.003 0 0 1 3-3h16a3.003 3.003 0 0 1 3 3v16a3.003 3.003 0 0 1-3 3z" opacity=".5"/><path fill="currentColor" d="M16.003 10.002H16V9a4.004 4.004 0 0 0-4-4h-2a5.006 5.006 0 0 0-5 5v4a5.006 5.006 0 0 0 5 5h4a5.006 5.006 0 0 0 5-5v-1a3 3 0 0 0-2.997-2.998zM10 9h1a1 1 0 1 1 0 2h-1a1 1 0 1 1 0-2zm4 6h-4a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2z"/>`
	bookmarkInnerSVG             = `<path fill="currentColor" d="M18 22a1 1 0 0 1-.5-.134L12 18.694l-5.5 3.172A1 1 0 0 1 5 21V5a3.003 3.003 0 0 1 3-3h8a3.003 3.003 0 0 1 3 3v16a1 1 0 0 1-1 1Z"/>`
	borderAltInnerSVG            = `<path fill="currentColor" d="M3.5 20.5a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1h16a1 1 0 0 1 0 2h-15v15a1 1 0 0 1-1 1Z"/><circle cx="19.5" cy="11.5" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="7.5" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="15.5" r="1" fill="currentColor" opacity=".5"/><circle cx="7.5" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="11.5" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="15.5" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="19.5" r="1" fill="currentColor" opacity=".5"/>`
	borderBottomInnerSVG         = `<path fill="currentColor" d="M20 21.5H4a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2Z"/><circle cx="12" cy="16.5" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="12.5" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="8.5" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="4.5" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="16.5" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="12.5" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="8.5" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="4.5" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="4.5" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="4.5" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="12.5" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="12.5" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="16.5" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="12.5" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="8.5" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="4.5" r="1" fill="currentColor" opacity=".5"/>`
	borderClearInnerSVG          = `<circle cx="12" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="4" r="1" fill="currentColor" opacity=".5"/>`
	borderHorizontalInnerSVG     = `<path fill="currentColor" d="M20 13H4a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2Z"/><circle cx="12" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="4" r="1" fill="currentColor" opacity=".5"/>`
	borderInnerInnerSVG          = `<path fill="currentColor" d="M19.965 13h-16a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M11.965 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z"/><circle cx="3.964" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="3.964" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="3.964" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="3.964" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="7.964" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="15.964" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="7.964" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="15.964" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="19.964" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="19.964" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="19.964" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="19.964" cy="4" r="1" fill="currentColor" opacity=".5"/>`
	borderLeftInnerSVG           = `<path fill="currentColor" d="M3.5 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z"/><circle cx="7.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="11.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="15.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="7.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="11.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="15.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="11.5" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="11.5" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="7.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="11.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="15.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="20" r="1" fill="currentColor" opacity=".5"/>`
	borderOutInnerSVG            = `<path fill="currentColor" d="M20 21H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1ZM5 19h14V5H5Z"/><circle cx="12" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="12" r="1" fill="currentColor" opacity=".5"/>`
	borderRightInnerSVG          = `<path fill="currentColor" d="M20.5 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z"/><circle cx="16.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="8.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="16.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="8.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="16.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="8.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="4" r="1" fill="currentColor" opacity=".5"/>`
	borderTopInnerSVG            = `<path fill="currentColor" d="M20 4.5H4a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2Z"/><circle cx="12" cy="7.5" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="11.5" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="15.5" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="7.5" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="11.5" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="15.5" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="11.5" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="11.5" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="7.5" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="11.5" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="15.5" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="19.5" r="1" fill="currentColor" opacity=".5"/>`
	borderVerticalInnerSVG       = `<path fill="currentColor" d="M11 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z"/><circle cx="7" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="15" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="7" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="15" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="7" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="15" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="20" r="1" fill="currentColor" opacity=".5"/>`
	boxInnerSVG                  = `<path fill="currentColor" d="M20.23 7.24L12 12L3.77 7.24a1.98 1.98 0 0 1 .7-.71L11 2.76c.62-.35 1.38-.35 2 0l6.53 3.77c.29.173.531.418.7.71z" opacity=".25"/><path fill="currentColor" d="M12 12v9.5a2.09 2.09 0 0 1-.91-.21L4.5 17.48a2.003 2.003 0 0 1-1-1.73v-7.5a2.06 2.06 0 0 1 .27-1.01L12 12z" opacity=".5"/><path fill="currentColor" d="M20.5 8.25v7.5a2.003 2.003 0 0 1-1 1.73l-6.62 3.82c-.275.13-.576.198-.88.2V12l8.23-4.76c.175.308.268.656.27 1.01z"/>`
	briefcaseInnerSVG            = `<path fill="currentColor" d="M10 6V5h4v1h2V5a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2v1h2z"/><path fill="currentColor" d="M9 15a1 1 0 0 1-1-1v-2a1 1 0 1 1 2 0v2a1 1 0 0 1-1 1zm6 0a1 1 0 0 1-1-1v-2a1 1 0 1 1 2 0v2a1 1 0 0 1-1 1z" opacity=".25"/><path fill="currentColor" d="M20 6H4a2 2 0 0 0-2 2v3a2 2 0 0 0 2 2h4v-1a1 1 0 1 1 2 0v1h4v-1a1 1 0 1 1 2 0v1h4a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2z" opacity=".5"/><path fill="currentColor" d="M20 13h-4v1a1 1 0 1 1-2 0v-1h-4v1a1 1 0 1 1-2 0v-1H4a2 2 0 0 1-2-2v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8a2 2 0 0 1-2 2z"/>`
	calenderInnerSVG             = `<path fill="currentColor" d="M22 10H2v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-9zM7 8a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1zm10 0a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1z" opacity=".5"/><path fill="currentColor" d="M19 4h-1v3a1 1 0 0 1-2 0V4H8v3a1 1 0 0 1-2 0V4H5a3 3 0 0 0-3 3v3h20V7a3 3 0 0 0-3-3z"/>`
	chartInnerSVG                = `<path fill="currentColor" d="M7 18a1 1 0 0 1-1-1v-4a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1zm5 0a1 1 0 0 1-1-1V7a1 1 0 0 1 2 0v10a1 1 0 0 1-1 1zm5 0a1 1 0 0 1-1-1v-6a1 1 0 0 1 2 0v6a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3ZM8 17a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0Zm5 0a1 1 0 0 1-2 0V7a1 1 0 0 1 2 0Zm5 0a1 1 0 0 1-2 0v-6a1 1 0 0 1 2 0Z" opacity=".5"/>`
	chartPieInnerSVG             = `<path fill="currentColor" d="M12 12V2c5.523 0 10 4.477 10 10H12z" opacity=".25"/><path fill="currentColor" d="m12 12l5 8.66A10.01 10.01 0 0 0 22 12H12z" opacity=".5"/><path fill="currentColor" d="M17 20.66L12 12V2c-5.523.002-9.999 4.48-9.997 10.003c.002 5.523 4.48 9.999 10.004 9.997A10 10 0 0 0 17 20.662l.003-.005l-.004.003z"/>`
	checkInnerSVG                = `<path fill="currentColor" d="M9.84 17.08a.997.997 0 0 1-.707-.293l-3.84-3.84a1 1 0 1 1 1.414-1.414l3.133 3.133l7.453-7.453a1 1 0 0 1 1.414 1.414l-8.16 8.16a.997.997 0 0 1-.707.293Z"/>`
	checkCircleInnerSVG          = `<path fill="currentColor" d="M10.313 16.094a.997.997 0 0 1-.708-.293l-2.812-2.813a1 1 0 0 1 1.414-1.414l2.105 2.106l5.481-5.48a1 1 0 0 1 1.414 1.413l-6.188 6.188a.997.997 0 0 1-.707.293Z" opacity=".99"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2Zm5.207 7.613l-6.188 6.188a1 1 0 0 1-1.414 0l-2.812-2.813a1 1 0 0 1 1.414-1.414l2.105 2.106l5.481-5.48a1 1 0 0 1 1.414 1.413Z" opacity=".5"/>`
	checkSquareInnerSVG          = `<path fill="currentColor" d="M10.2 16.4a.997.997 0 0 1-.707-.293l-3.2-3.2a1 1 0 0 1 1.414-1.414l2.493 2.493l6.093-6.093a1 1 0 0 1 1.414 1.414l-6.8 6.8a.997.997 0 0 1-.707.293Z"/><path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm-3.293 7.307l-6.8 6.8a1 1 0 0 1-1.414 0l-3.2-3.2a1 1 0 0 1 1.414-1.414l2.493 2.493l6.093-6.093a1 1 0 0 1 1.414 1.414Z" opacity=".5"/>`
	circleInnerSVG               = `<circle cx="12" cy="12" r="10" fill="currentColor"/>`
	circleLayerInnerSVG          = `<path fill="currentColor" d="M15 2a6.998 6.998 0 0 0-6.88 5.737a6 6 0 0 1 8.143 8.143A6.997 6.997 0 0 0 15 2z" opacity=".25"/><circle cx="7" cy="17" r="5" fill="currentColor"/><path fill="currentColor" d="M11 7a6 6 0 0 0-5.97 5.406a4.997 4.997 0 0 1 6.564 6.564A6 6 0 0 0 11 7z" opacity=".5"/>`
	clinicMedicalInnerSVG        = `<path fill="currentColor" d="M21 12a.996.996 0 0 1-.664-.252L12 4.338l-8.336 7.41a1 1 0 0 1-1.328-1.496l9-8a.999.999 0 0 1 1.328 0l9 8A1 1 0 0 1 21 12Z" opacity=".5"/><path fill="currentColor" d="M14 13h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2Z"/><path fill="currentColor" d="m12 4.338l-8 7.111V21a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9.551ZM14 15h-1v1a1 1 0 0 1-2 0v-1h-1a1 1 0 0 1 0-2h1v-1a1 1 0 0 1 2 0v1h1a1 1 0 0 1 0 2Z" opacity=".25"/>`
	clockInnerSVG                = `<path fill="currentColor" d="M12 6a1 1 0 0 1 1 1v4.422l2.098 1.212a1 1 0 0 1-1 1.732l-2.598-1.5A1.005 1.005 0 0 1 11 12V7a1 1 0 0 1 1-1Z"/><path fill="currentColor" d="M2 12A10 10 0 1 0 12 2A10 10 0 0 0 2 12Zm9-5a1 1 0 0 1 2 0v4.422l2.098 1.212a1 1 0 0 1-1 1.732l-2.598-1.5A1.005 1.005 0 0 1 11 12Z" opacity=".5"/>`
	clockEightInnerSVG           = `<path fill="currentColor" d="M12 6a1 1 0 0 0-1 1v4.422l-2.098 1.212a1 1 0 0 0 1 1.732l2.598-1.5A1.005 1.005 0 0 0 13 12V7a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm1 10a1.005 1.005 0 0 1-.5.866l-2.598 1.5a1 1 0 0 1-1-1.732L11 11.422V7a1 1 0 0 1 2 0Z" opacity=".5"/>`
	clockFiveInnerSVG            = `<path fill="currentColor" d="M12 6a1.003 1.003 0 0 1 1 1v4.73l1.6 2.77a1 1 0 1 1-1.73 1l-1.69-2.93A.999.999 0 0 1 11 12V7a1.003 1.003 0 0 1 1-1Z"/><path fill="currentColor" d="M2 12A10 10 0 1 0 12 2A10 10 0 0 0 2 12Zm9-5a1 1 0 0 1 2 0v4.73l1.6 2.77a1 1 0 1 1-1.73 1l-1.69-2.93A.999.999 0 0 1 11 12Z" opacity=".5"/>`
	clockNineInnerSVG            = `<path fill="currentColor" d="M12 6a1 1 0 0 0-1 1v4H9a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm1 10a1 1 0 0 1-1 1H9a1 1 0 0 1 0-2h2V7a1 1 0 0 1 2 0Z" opacity=".5"/>`
	clockSevenInnerSVG           = `<path fill="currentColor" d="M12 6a1.003 1.003 0 0 0-1 1v4.73L9.4 14.5a1 1 0 1 0 1.73 1l1.69-2.93A.999.999 0 0 0 13 12V7a1.003 1.003 0 0 0-1-1Z"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm1 10a.999.999 0 0 1-.18.57l-1.69 2.93a1 1 0 1 1-1.73-1l1.6-2.77V7a1 1 0 0 1 2 0Z" opacity=".5"/>`
	clockTenInnerSVG             = `<circle cx="12" cy="12" r="10" fill="currentColor" opacity=".5"/><path fill="currentColor" d="M13 7a1 1 0 0 0-2 0v3.268l-1.098-.634a1 1 0 0 0-1 1.732l2.598 1.5A1.014 1.014 0 0 0 13 12Z"/>`
	clockThreeInnerSVG           = `<path fill="currentColor" d="M12 6a1 1 0 0 1 1 1v4h2a1 1 0 0 1 0 2h-3a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1Z"/><path fill="currentColor" d="M2 12A10 10 0 1 0 12 2A10 10 0 0 0 2 12Zm9-5a1 1 0 0 1 2 0v4h2a1 1 0 0 1 0 2h-3a1 1 0 0 1-1-1Z" opacity=".5"/>`
	clockTwoInnerSVG             = `<circle cx="12" cy="12" r="10" fill="currentColor" opacity=".5"/><path fill="currentColor" d="M11 7a1 1 0 0 1 2 0v3.268l1.098-.634a1 1 0 0 1 1 1.732l-2.598 1.5A1.014 1.014 0 0 1 11 12Z"/>`
	columnsInnerSVG              = `<path fill="currentColor" d="M11 2h2v20h-2z" opacity=".25"/><path fill="currentColor" d="M3 2h8v20H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M13 2h8a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1h-8V2z" opacity=".5"/>`
	commentInnerSVG              = `<path fill="currentColor" d="M12 22H3a1 1 0 0 1-.707-1.707l1.964-1.964A10 10 0 1 1 12 22Z" opacity=".5"/>`
	commentAltInnerSVG           = `<path fill="currentColor" d="M21 22a.999.999 0 0 1-.707-.293L16.586 18H5a3.003 3.003 0 0 1-3-3V5a3.003 3.003 0 0 1 3-3h14a3.003 3.003 0 0 1 3 3v16a1 1 0 0 1-1 1Z" opacity=".5"/>`
	commentAltDotsInnerSVG       = `<path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v10a3.003 3.003 0 0 0 3 3h11.586l3.707 3.707A1 1 0 0 0 22 21V5a3.003 3.003 0 0 0-3-3ZM8 11a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z" opacity=".5"/><circle cx="8" cy="10" r="1" fill="currentColor"/><circle cx="12" cy="10" r="1" fill="currentColor"/><circle cx="16" cy="10" r="1" fill="currentColor"/>`
	commentAltMessageInnerSVG    = `<path fill="currentColor" d="M17 9H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm0 4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z"/><path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v10a3.003 3.003 0 0 0 3 3h11.586l3.707 3.707A1 1 0 0 0 22 21V5a3.003 3.003 0 0 0-3-3Zm-2 11H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Zm0-4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z" opacity=".5"/>`
	commentAltPlusInnerSVG       = `<path fill="currentColor" d="M15 9h-2V7a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2Z"/><path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v10a3.003 3.003 0 0 0 3 3h11.586l3.707 3.707A1 1 0 0 0 22 21V5a3.003 3.003 0 0 0-3-3Zm-4 9h-2v2a1 1 0 0 1-2 0v-2H9a1 1 0 0 1 0-2h2V7a1 1 0 0 1 2 0v2h2a1 1 0 0 1 0 2Z" opacity=".5"/>`
	commentDotsInnerSVG          = `<circle cx="12" cy="12" r="1" fill="currentColor"/><path fill="currentColor" d="M12 2a10 10 0 0 0-7.743 16.33l-1.964 1.963A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20ZM8 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z" opacity=".5"/><circle cx="8" cy="12" r="1" fill="currentColor"/><circle cx="16" cy="12" r="1" fill="currentColor"/>`
	commentMessageInnerSVG       = `<path fill="currentColor" d="M17 13H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M12 2a10 10 0 0 0-7.743 16.33l-1.964 1.963A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20ZM9 7h6a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2Zm6 10H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Zm2-4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z" opacity=".5"/><path fill="currentColor" d="M15 17H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2zm0-8H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2z"/>`
	commentPlusInnerSVG          = `<path fill="currentColor" d="M15 11h-2V9a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2Z"/><path fill="currentColor" d="M12 2a10 10 0 0 0-7.743 16.33l-1.964 1.963A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20Zm3 11h-2v2a1 1 0 0 1-2 0v-2H9a1 1 0 0 1 0-2h2V9a1 1 0 0 1 2 0v2h2a1 1 0 0 1 0 2Z" opacity=".5"/>`
	compressInnerSVG             = `<path fill="currentColor" d="M8 22a1 1 0 0 1-1-1v-4H3a1 1 0 0 1 0-2h5a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1zM8 9H3a1 1 0 0 1 0-2h4V3a1 1 0 0 1 2 0v5a1 1 0 0 1-1 1zm13 0h-5a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v4h4a1 1 0 0 1 0 2zm-5 13a1 1 0 0 1-1-1v-5a1 1 0 0 1 1-1h5a1 1 0 0 1 0 2h-4v4a1 1 0 0 1-1 1z"/>`
	cornerDownLeftInnerSVG       = `<path fill="currentColor" d="M9.942 22a.997.997 0 0 1-.707-.293l-4.628-4.628a1 1 0 0 1 0-1.414l4.628-4.63a1 1 0 0 1 1.414 1.415l-3.92 3.922l3.92 3.921A1 1 0 0 1 9.942 22Z"/><path fill="currentColor" d="M15.686 17.372H5.314a1 1 0 0 1 0-2h10.372a2.002 2.002 0 0 0 2-2V3a1 1 0 0 1 2 0v10.372a4.004 4.004 0 0 1-4 4Z"/>`
	cornerDownRightInnerSVG      = `<path fill="currentColor" d="M14.058 22a1 1 0 0 1-.707-1.707l3.92-3.921l-3.92-3.922a1 1 0 1 1 1.414-1.414l4.628 4.629a1 1 0 0 1 0 1.414l-4.628 4.628a.997.997 0 0 1-.707.293Z"/><path fill="currentColor" d="M18.686 17.372H9.314a5.006 5.006 0 0 1-5-5V3a1 1 0 0 1 2 0v9.372a3.003 3.003 0 0 0 3 3h9.372a1 1 0 0 1 0 2Z"/>`
	cornerLeftDownInnerSVG       = `<path fill="currentColor" d="M21 4.314h-9.372a5.006 5.006 0 0 0-5 5v6.957l-2.921-2.92a1 1 0 0 0-1.414 1.413l4.628 4.628a1.003 1.003 0 0 0 1.415 0l4.628-4.628a1 1 0 0 0-1.414-1.414l-2.922 2.922V9.314a3.003 3.003 0 0 1 3-3H21a1 1 0 0 0 0-2Z"/>`
	cornerRightDownInnerSVG      = `<path fill="currentColor" d="M21.707 13.35a1 1 0 0 0-1.414 0l-2.921 2.921V8.314a4.005 4.005 0 0 0-4-4H3a1 1 0 0 0 0 2h10.372a2.002 2.002 0 0 1 2 2v7.958L12.45 13.35a1 1 0 0 0-1.414 1.414l4.628 4.628a1.004 1.004 0 0 0 1.415 0l4.628-4.628a1 1 0 0 0 0-1.414Z"/>`
	cornerUpLeftInnerSVG         = `<path fill="currentColor" d="M14.686 6.628H7.729l2.92-2.921a1 1 0 0 0-1.413-1.414L4.608 6.921a1.003 1.003 0 0 0 0 1.415l4.628 4.628a1 1 0 0 0 1.414-1.414L7.728 8.628h6.958a3.003 3.003 0 0 1 3 3V21a1 1 0 0 0 2 0v-9.372a5.006 5.006 0 0 0-5-5Z"/>`
	cornerUpRightInnerSVG        = `<path fill="currentColor" d="M19.608 8.01a1.004 1.004 0 0 0-.216-1.089l-4.628-4.628a1 1 0 0 0-1.414 1.414l2.921 2.921H8.314a4.005 4.005 0 0 0-4 4V21a1 1 0 0 0 2 0V10.628a2.002 2.002 0 0 1 2-2h7.958L13.35 11.55a1 1 0 1 0 1.414 1.414l4.628-4.628a1 1 0 0 0 .216-.325Z"/>`
	coronavirusInnerSVG          = `<circle cx="9.5" cy="10.5" r="1.5" fill="currentColor" opacity=".7"/><circle cx="9" cy="15" r="1" fill="currentColor"/><circle cx="14.5" cy="13.5" r="1.5" fill="currentColor" opacity=".7"/><circle cx="15" cy="9" r="1" fill="currentColor"/><path fill="currentColor" d="M12 8a1 1 0 0 1-1-1V2a1 1 0 0 1 2 0v5a1 1 0 0 1-1 1zm0 15a1 1 0 0 1-1-1v-5a1 1 0 0 1 2 0v5a1 1 0 0 1-1 1zm10-10h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2zM6 13H2a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2z" opacity=".25"/><path fill="currentColor" d="M18 13a1 1 0 0 1 0-2h2.941A9.013 9.013 0 0 0 13 3.059V7a1 1 0 0 1-2 0V3.059A9.013 9.013 0 0 0 3.059 11H6a1 1 0 0 1 0 2H3.059A9.013 9.013 0 0 0 11 20.941V17a1 1 0 0 1 2 0v3.941A9.013 9.013 0 0 0 20.941 13Zm-9 3a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm.5-4a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 9.5 12Zm5 3a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Zm.5-5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z" opacity=".5"/><path fill="currentColor" d="M4.979 6.393a9.063 9.063 0 0 1 1.414-1.414l-.757-.757a1 1 0 0 0-1.414 1.414zm0 11.214l-.757.757a1 1 0 1 0 1.414 1.414l.757-.757a9.063 9.063 0 0 1-1.414-1.414zm14.042 0a9.063 9.063 0 0 1-1.414 1.414l.757.757a1 1 0 0 0 1.414-1.414zm0-11.214l.757-.757a1 1 0 0 0-1.414-1.414l-.757.757a9.063 9.063 0 0 1 1.414 1.414z"/>`
	cssThreeInnerSVG             = `<path fill="currentColor" d="m9.883 21.095l9.098-3.02L22 3H4.958L4.31 6.364h13.657l-.425 2.148H3.864l-.669 3.363h13.658l-.75 3.85l-5.511 1.803l-4.762-1.802l.325-1.682H2.79L2 18.076l7.883 3.02"/>`
	cssThreeSimpleInnerSVG       = `<path fill="currentColor" d="M20.772 3.364A1 1 0 0 0 20 3H6a1 1 0 0 0 0 2h12.786l-.78 4H5a1 1 0 0 0 0 2h12.615l-1.163 5.955l-6.323 1.997l-5.41-1.7l.203-1.064a1 1 0 0 0-1.964-.375l-.37 1.94a1 1 0 0 0 .682 1.141l6.56 2.06a1.002 1.002 0 0 0 .601 0l7.19-2.27a1 1 0 0 0 .68-.763l2.68-13.73a1 1 0 0 0-.209-.827Z"/>`
	cubeInnerSVG                 = `<path fill="currentColor" d="M12 12.3L3.5 7.05L12 1.8l8.5 5.25z"/><path fill="currentColor" d="M12 22.2v-9.9l8.5-5.25v9.9z" opacity=".25"/><path fill="currentColor" d="m12 22.2l-8.5-5.25v-9.9L12 12.3z" opacity=".5"/>`
	dialpadInnerSVG              = `<circle cx="12" cy="9" r="2" fill="currentColor" opacity=".5"/><circle cx="12" cy="3" r="2" fill="currentColor" opacity=".5"/><circle cx="12" cy="15" r="2" fill="currentColor" opacity=".5"/><circle cx="6" cy="9" r="2" fill="currentColor" opacity=".5"/><circle cx="6" cy="3" r="2" fill="currentColor" opacity=".5"/><circle cx="6" cy="15" r="2" fill="currentColor" opacity=".5"/><circle cx="18" cy="9" r="2" fill="currentColor" opacity=".5"/><circle cx="18" cy="3" r="2" fill="currentColor" opacity=".5"/><circle cx="18" cy="15" r="2" fill="currentColor" opacity=".5"/><circle cx="12" cy="21" r="2" fill="currentColor" opacity=".5"/>`
	dialpadAltInnerSVG           = `<rect width="4" height="4" x="10" y="6.955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="10" y=".955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="10" y="13.045" fill="currentColor" rx=".545"/><rect width="4" height="4" x="4" y="6.955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="4" y=".955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="4" y="13.045" fill="currentColor" rx=".545"/><rect width="4" height="4" x="16" y="6.955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="16" y=".955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="16" y="13.045" fill="currentColor" rx=".545"/><rect width="4" height="4" x="10" y="19" fill="currentColor" rx=".545"/>`
	directionInnerSVG            = `<path fill="currentColor" d="M12 19.5a.997.997 0 0 1-.707-.293l-3-3a1 1 0 0 1 1.414-1.414L12 17.086l2.293-2.293a1 1 0 0 1 1.414 1.414l-3 3A.997.997 0 0 1 12 19.5zm3-9a.997.997 0 0 1-.707-.293L12 7.914l-2.293 2.293a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.414 0l3 3A1 1 0 0 1 15 10.5z"/>`
	discordInnerSVG              = `<path fill="currentColor" d="M18.891 2H5.11A2.114 2.114 0 0 0 3 4.119v13.906a2.114 2.114 0 0 0 2.109 2.119h11.664l-.546-1.903l1.317 1.224l1.245 1.152L21 22.571V4.12A2.114 2.114 0 0 0 18.891 2Zm-3.97 13.433s-.37-.442-.679-.833a3.246 3.246 0 0 0 1.862-1.224a5.88 5.88 0 0 1-1.183.607a6.769 6.769 0 0 1-1.491.442a7.205 7.205 0 0 1-2.664-.01a8.641 8.641 0 0 1-1.512-.442a6.035 6.035 0 0 1-.751-.35c-.031-.02-.062-.031-.093-.052a.142.142 0 0 1-.04-.03a3.612 3.612 0 0 1-.289-.175a3.199 3.199 0 0 0 1.8 1.213c-.308.391-.689.854-.689.854a3.73 3.73 0 0 1-3.137-1.563a13.774 13.774 0 0 1 1.481-5.997a5.086 5.086 0 0 1 2.89-1.08l.103.124a6.937 6.937 0 0 0-2.705 1.347s.226-.123.607-.298a7.724 7.724 0 0 1 2.335-.648a1.009 1.009 0 0 1 .175-.02a8.703 8.703 0 0 1 2.077-.021a8.384 8.384 0 0 1 3.096.987a6.845 6.845 0 0 0-2.56-1.306l.143-.165a5.086 5.086 0 0 1 2.89 1.08a13.774 13.774 0 0 1 1.482 5.997a3.76 3.76 0 0 1-3.148 1.563Zm-1.028-4.803a1.146 1.146 0 1 0 1.049 1.141a1.096 1.096 0 0 0-1.05-1.141"/><path fill="currentColor" d="M14.921 15.433s-.37-.442-.679-.833a3.246 3.246 0 0 0 1.862-1.224a5.88 5.88 0 0 1-1.183.607a6.769 6.769 0 0 1-1.491.442a7.205 7.205 0 0 1-2.664-.01a8.641 8.641 0 0 1-1.512-.442a6.035 6.035 0 0 1-.751-.35c-.031-.02-.062-.031-.093-.052a.142.142 0 0 1-.04-.03a3.612 3.612 0 0 1-.289-.175a3.199 3.199 0 0 0 1.8 1.213c-.308.391-.689.854-.689.854a3.73 3.73 0 0 1-3.137-1.563a13.774 13.774 0 0 1 1.481-5.997a5.086 5.086 0 0 1 2.89-1.08l.103.124a6.937 6.937 0 0 0-2.705 1.347s.226-.123.607-.298a7.724 7.724 0 0 1 2.335-.648a1.009 1.009 0 0 1 .175-.02a8.703 8.703 0 0 1 2.077-.021a8.384 8.384 0 0 1 3.096.987a6.845 6.845 0 0 0-2.56-1.306l.143-.165a5.086 5.086 0 0 1 2.89 1.08a13.774 13.774 0 0 1 1.482 5.997a3.76 3.76 0 0 1-3.148 1.563Zm-1.028-4.803a1.146 1.146 0 1 0 1.049 1.141a1.096 1.096 0 0 0-1.05-1.141" opacity=".25"/>`
	dockerInnerSVG               = `<path fill="currentColor" d="M21.805 10.077a2.627 2.627 0 0 0-1.632-.427a5.189 5.189 0 0 0-.844.074A3.18 3.18 0 0 0 17.9 7.581l-.287-.167l-.186.27a3.967 3.967 0 0 0-.51 1.187a2.819 2.819 0 0 0 .334 2.217a3.936 3.936 0 0 1-1.457.352H2.623a.622.622 0 0 0-.622.622a9.386 9.386 0 0 0 .575 3.385a5.078 5.078 0 0 0 2.004 2.607A8.868 8.868 0 0 0 8.977 19a13.486 13.486 0 0 0 2.44-.223a10.068 10.068 0 0 0 3.19-1.16a8.734 8.734 0 0 0 2.17-1.78a11.81 11.81 0 0 0 2.125-3.664h.185a3.052 3.052 0 0 0 2.236-.844a2.47 2.47 0 0 0 .594-.872l.083-.241Z"/><path fill="currentColor" d="M3.847 11.06H5.61a.156.156 0 0 0 .157-.158V9.325a.156.156 0 0 0-.157-.157H3.847a.156.156 0 0 0-.158.157v1.577a.162.162 0 0 0 .158.158Zm2.43 0H8.04a.156.156 0 0 0 .157-.158V9.325a.156.156 0 0 0-.157-.157H6.277a.156.156 0 0 0-.157.157v1.577a.162.162 0 0 0 .157.158m2.477 0h1.762a.156.156 0 0 0 .158-.158V9.325a.156.156 0 0 0-.158-.157H8.754a.156.156 0 0 0-.158.157v1.577a.151.151 0 0 0 .158.158Zm2.44 0h1.762a.156.156 0 0 0 .158-.158V9.325a.156.156 0 0 0-.158-.157h-1.762a.156.156 0 0 0-.158.157v1.577a.156.156 0 0 0 .158.158ZM6.277 8.806H8.04a.163.163 0 0 0 .157-.158V7.071a.156.156 0 0 0-.157-.157H6.277a.156.156 0 0 0-.157.157v1.577a.17.17 0 0 0 .157.158m2.477 0h1.762a.163.163 0 0 0 .158-.158V7.071a.156.156 0 0 0-.158-.157H8.754a.156.156 0 0 0-.158.157v1.577a.156.156 0 0 0 .158.158m2.44 0h1.762a.163.163 0 0 0 .158-.158V7.071a.163.163 0 0 0-.158-.157h-1.762a.156.156 0 0 0-.158.157v1.577a.163.163 0 0 0 .158.158m0-2.263h1.762a.156.156 0 0 0 .158-.158V4.808a.163.163 0 0 0-.158-.158h-1.762a.156.156 0 0 0-.158.158v1.577a.163.163 0 0 0 .158.158m2.458 4.517h1.762a.156.156 0 0 0 .158-.158V9.325a.156.156 0 0 0-.158-.157h-1.762a.156.156 0 0 0-.158.157v1.577a.162.162 0 0 0 .158.158" opacity=".25"/>`
	documentLayoutCenterInnerSVG = `<path fill="currentColor" d="M21 8h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2zm0 4h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2zM5 8H3a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2z" opacity=".5"/><rect width="8" height="8" x="8" y="4" fill="currentColor" rx="1"/><path fill="currentColor" d="M21 16H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8 4H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z" opacity=".5"/>`
	documentLayoutLeftInnerSVG   = `<path fill="currentColor" d="M21 8h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2zm0 4h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2z" opacity=".5"/><rect width="8" height="8" x="2" y="4" fill="currentColor" rx="1"/><path fill="currentColor" d="M21 16H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8 4H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z" opacity=".5"/>`
	documentLayoutRightInnerSVG  = `<path fill="currentColor" d="M11 8H3a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2z" opacity=".5"/><rect width="8" height="8" x="14" y="4" fill="currentColor" rx="1"/><path fill="currentColor" d="M21 16H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8 4H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z" opacity=".5"/>`
	downloadAltInnerSVG          = `<path fill="currentColor" d="M15.707 13.293a1 1 0 0 0-1.414 0L13 14.586V3a1 1 0 0 0-2 0v11.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0 0-1.414Z"/><path fill="currentColor" d="M18 9h-5v5.586l1.293-1.293a1 1 0 0 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 1.414-1.414L11 14.586V9H6a3.003 3.003 0 0 0-3 3v7a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3v-7a3.003 3.003 0 0 0-3-3Z" opacity=".5"/>`
	dribbbleInnerSVG             = `<circle cx="11.97" cy="11.97" r="9" fill="currentColor" opacity=".5"/><path fill="currentColor" d="M2 12a9.796 9.796 0 0 1 1.34-5.02a9.93 9.93 0 0 1 3.64-3.64A9.796 9.796 0 0 1 12 2a9.796 9.796 0 0 1 5.02 1.34a9.929 9.929 0 0 1 3.64 3.64A9.796 9.796 0 0 1 22 12a9.796 9.796 0 0 1-1.34 5.02a9.929 9.929 0 0 1-3.64 3.64A9.796 9.796 0 0 1 12 22a9.796 9.796 0 0 1-5.02-1.34a9.93 9.93 0 0 1-3.64-3.64A9.796 9.796 0 0 1 2 12zm1.66 0a8.064 8.064 0 0 0 2.1 5.5c.755-1.39 1.79-2.61 3.04-3.58a9.942 9.942 0 0 1 4.06-2.14c-.2-.467-.393-.887-.58-1.26a24.3 24.3 0 0 1-7.44 1.1c-.52 0-.907-.007-1.16-.02c0 .053-.003.12-.01.2s-.01.147-.01.2zm.26-2.06c.293.027.727.04 1.3.04a21.95 21.95 0 0 0 6.34-.9a20.251 20.251 0 0 0-3.34-4.5A8.12 8.12 0 0 0 5.51 6.8a8.539 8.539 0 0 0-1.59 3.14zm2.98 8.64a8.173 8.173 0 0 0 8.04 1.2a29.368 29.368 0 0 0-1.56-6.62a8.53 8.53 0 0 0-3.71 2.02a11.345 11.345 0 0 0-2.77 3.4zM9.96 3.94a21.254 21.254 0 0 1 3.26 4.54a9.96 9.96 0 0 0 4.1-2.9A8.107 8.107 0 0 0 12 3.66a7.677 7.677 0 0 0-2.04.28zm3.98 5.96c.2.427.427.967.68 1.62c.987-.093 2.06-.14 3.22-.14c.827 0 1.647.02 2.46.06a8.03 8.03 0 0 0-1.96-4.84a9.896 9.896 0 0 1-4.4 3.3zm1.18 3.02c.68 1.97 1.142 4.01 1.38 6.08a8.345 8.345 0 0 0 2.58-2.62a8.078 8.078 0 0 0 1.2-3.46c-.973-.067-1.86-.1-2.66-.1c-.733 0-1.567.033-2.5.1z"/>`
	dropboxInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M16.53 10.091L21 12.939l-4.502 2.868L12 12.941l-4.498 2.866L3 12.939l4.47-2.848L3 7.243l4.502-2.868L12 7.241l4.498-2.866L21 7.243z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M16.467 10.091L12 7.245l-4.467 2.846L12 12.936z" clip-rule="evenodd" opacity=".25"/><path fill="currentColor" fill-rule="evenodd" d="m7.531 16.757l-.029-.95L12 12.941l4.498 2.866l.036.95l-4.502 2.868z" clip-rule="evenodd" opacity=".5"/>`
	ellipsisHInnerSVG            = `<circle cx="12" cy="12" r="2" fill="currentColor"/><circle cx="5" cy="12" r="2" fill="currentColor"/><circle cx="19" cy="12" r="2" fill="currentColor"/>`
	ellipsisVInnerSVG            = `<circle cx="12" cy="12" r="2" fill="currentColor"/><circle cx="12" cy="5" r="2" fill="currentColor"/><circle cx="12" cy="19" r="2" fill="currentColor"/>`
	entryInnerSVG                = `<path fill="currentColor" d="M20 11h-8.586l2.293-2.293a1 1 0 0 0-1.414-1.414l-4 4a1 1 0 0 0 0 1.414l4 4a1 1 0 0 0 1.414-1.414L11.414 13H20Z"/><path fill="currentColor" d="M11.414 11H20V5a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h10a3.003 3.003 0 0 0 3-3v-6h-8.586l2.293 2.293a1 1 0 1 1-1.414 1.414l-4-4a1 1 0 0 1 0-1.414l4-4a1 1 0 0 1 1.414 1.414Z" opacity=".5"/>`
	exclamationCircleInnerSVG    = `<circle cx="12" cy="12" r="10" fill="currentColor" opacity=".5"/><circle cx="12" cy="16" r="1" fill="currentColor"/><path fill="currentColor" d="M12 13a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1Z"/>`
	exclamationOctagonInnerSVG   = `<path fill="currentColor" d="M15.728 22H8.272a1 1 0 0 1-.707-.293l-5.272-5.272A1 1 0 0 1 2 15.728V8.272a1 1 0 0 1 .293-.707l5.272-5.272A1 1 0 0 1 8.272 2h7.456a1 1 0 0 1 .707.293l5.272 5.272a1 1 0 0 1 .293.707v7.456a1 1 0 0 1-.293.707l-5.272 5.272a1 1 0 0 1-.707.293Z" opacity=".5"/><circle cx="12" cy="16" r="1" fill="currentColor"/><path fill="currentColor" d="M12 13a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1Z"/>`
	exclamationTriangleInnerSVG  = `<path fill="currentColor" d="M20.057 22H3.943a3.023 3.023 0 0 1-2.618-4.534L9.382 3.511a3.023 3.023 0 0 1 5.236 0l8.057 13.955A3.023 3.023 0 0 1 20.057 22Z" opacity=".5"/><circle cx="12" cy="17" r="1" fill="currentColor"/><path fill="currentColor" d="M12 14a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1Z"/>`
	exitInnerSVG                 = `<path fill="currentColor" d="m15.707 11.293l-4-4a1 1 0 0 0-1.414 1.414L12.586 11H4v2h8.586l-2.293 2.293a1 1 0 1 0 1.414 1.414l4-4a1 1 0 0 0 0-1.414Z"/><path fill="currentColor" d="M17 2H7a3.003 3.003 0 0 0-3 3v6h8.586l-2.293-2.293a1 1 0 0 1 1.414-1.414l4 4a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414-1.414L12.586 13H4v6a3.003 3.003 0 0 0 3 3h10a3.003 3.003 0 0 0 3-3V5a3.003 3.003 0 0 0-3-3Z" opacity=".5"/>`
	facebookInnerSVG             = `<path fill="currentColor" d="M12.683 22v-7.745h-2.607v-3.018h2.607V9.01a3.637 3.637 0 0 1 3.882-3.99a21.32 21.32 0 0 1 2.329.119v2.7h-1.599c-1.253 0-1.495.596-1.495 1.47v1.927h2.989l-.39 3.018h-2.6V22h-3.116z" opacity=".5"/><path fill="currentColor" d="M20.999 2H2.998a1 1 0 0 0-1 1v18.001a1 1 0 0 0 1 1h18.001a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zm-2.105 5.84h-1.599c-1.253 0-1.495.596-1.495 1.47v1.926h2.989l-.39 3.019h-2.6V22h-3.116v-7.745h-2.607v-3.019h2.607V9.01a3.637 3.637 0 0 1 3.882-3.99a21.37 21.37 0 0 1 2.329.12v2.7z"/>`
	facebookFInnerSVG            = `<path fill="currentColor" d="M13.355 22v-9.123h3.062l.459-3.555h-3.52v-2.27c0-1.03.285-1.731 1.761-1.731L17 5.32V2.14A25.233 25.233 0 0 0 14.257 2c-2.715 0-4.573 1.657-4.573 4.7v2.622h-3.07v3.555h3.07V22h3.671Z"/>`
	facebookMessengerInnerSVG    = `<path fill="currentColor" d="m18.004 9.465l-2.936 4.658a1.501 1.501 0 0 1-2.169.4l-2.336-1.752a.6.6 0 0 0-.722.002l-3.157 2.396a.475.475 0 0 1-.688-.632l2.938-4.66a1.501 1.501 0 0 1 2.169-.4l2.336 1.753a.6.6 0 0 0 .722-.002l3.155-2.395a.475.475 0 0 1 .688.632z" opacity=".5"/><path fill="currentColor" d="M12.59 2A9.725 9.725 0 0 0 12 2c-5.327-.193-9.801 3.969-9.994 9.295c-.005.135-.007.27-.006.405a9.49 9.49 0 0 0 3.14 7.175a.806.806 0 0 1 .27.57l.055 1.779a.801.801 0 0 0 1.122.708l1.984-.875a.798.798 0 0 1 .534-.04c.943.257 1.917.386 2.895.384c5.357.163 9.833-4.048 9.996-9.405c.163-5.358-4.048-9.833-9.405-9.996zm5.414 7.465l-2.936 4.658a1.501 1.501 0 0 1-2.169.4l-2.336-1.752a.6.6 0 0 0-.722.002l-3.157 2.396a.475.475 0 0 1-.688-.632l2.938-4.66a1.501 1.501 0 0 1 2.169-.4l2.336 1.752a.6.6 0 0 0 .722-.001l3.155-2.395a.475.475 0 0 1 .688.632z"/>`
	facebookMessengerAltInnerSVG = `<path fill="currentColor" d="M12 2a9.651 9.651 0 0 0-10 9.7a9.49 9.49 0 0 0 3.14 7.175a.806.806 0 0 1 .27.57l.055 1.779a.801.801 0 0 0 1.122.708l1.984-.875a.798.798 0 0 1 .534-.04A10.876 10.876 0 0 0 12 21.4A9.705 9.705 0 1 0 12 2Z" opacity=".5"/><path fill="currentColor" d="M6.499 14.772a1 1 0 0 1-.765-1.642l3.052-3.636a.996.996 0 0 1 1.29-.21l3.346 2.056l2.312-2.755a1 1 0 1 1 1.532 1.285l-2.867 3.416a1 1 0 0 1-1.289.21L9.764 11.44l-2.498 2.975a.994.994 0 0 1-.767.357Z"/>`
	favoriteInnerSVG             = `<path fill="currentColor" d="M17.562 21.56a1 1 0 0 1-.465-.116L12 18.764l-5.097 2.68a1 1 0 0 1-1.45-1.053l.973-5.676l-4.124-4.02a1 1 0 0 1 .554-1.705l5.699-.828l2.549-5.164a1.04 1.04 0 0 1 1.793 0l2.548 5.164l5.699.828a1 1 0 0 1 .554 1.705l-4.124 4.02l.974 5.676a1 1 0 0 1-.985 1.169Z"/>`
	fiveHundredPxInnerSVG        = `<path fill="currentColor" d="m19.705 19.266l-.067.067a8.826 8.826 0 0 1-2.89 1.953a9.136 9.136 0 0 1-3.539.714a9.044 9.044 0 0 1-6.428-2.667a8.972 8.972 0 0 1-2.556-4.99q-.045-.312.536-.401q.569-.09.625.223a.065.065 0 0 1 .01.045a9.896 9.896 0 0 0 .514 1.774a7.543 7.543 0 0 0 1.697 2.523a7.916 7.916 0 0 0 8.683 1.696a7.961 7.961 0 0 0 2.51-1.696l.068-.067a.27.27 0 0 1 .279-.067a1.023 1.023 0 0 1 .368.245q.402.413.19.648Zm-5.268-6.864l-.737.736l.704.704q.234.234-.079.546a.532.532 0 0 1-.357.19a.287.287 0 0 1-.212-.111l-.692-.681l-.736.736a.227.227 0 0 1-.168.056a.506.506 0 0 1-.346-.178l-.022-.023a.452.452 0 0 1-.201-.323a.33.33 0 0 1 .09-.19l.736-.726l-.737-.736q-.178-.179.157-.502a.55.55 0 0 1 .346-.201a.255.255 0 0 1 .145.055l.725.737l.726-.725q.2-.19.535.145q.302.301.123.49Zm5.257.636a6.48 6.48 0 0 1-1.92 4.62a6.615 6.615 0 0 1-2.087 1.407a6.563 6.563 0 0 1-5.09 0a6.615 6.615 0 0 1-2.086-1.407a6.34 6.34 0 0 1-1.395-2.075a1.631 1.631 0 0 1-.168-.447h-.01q-.101-.301.479-.49q.558-.18.67.133a5.91 5.91 0 0 0 1.082 1.864h.011v-3.806a3.592 3.592 0 0 1 1.139-2.59a3.886 3.886 0 0 1 2.823-1.149a3.836 3.836 0 0 1 2.802 1.15a3.768 3.768 0 0 1 1.16 2.779a3.973 3.973 0 0 1-3.962 3.962a4.377 4.377 0 0 1-1.25-.179q-.312-.123-.145-.68q.179-.57.491-.48l.157.033q.156.033.362.067a2.207 2.207 0 0 0 .34.033a2.682 2.682 0 0 0 1.965-.798a2.64 2.64 0 0 0 .804-1.947a2.56 2.56 0 0 0-.804-1.909a2.654 2.654 0 0 0-1.953-.792a2.545 2.545 0 0 0-1.987.893a2.625 2.625 0 0 0-.714 1.786v4.609a5.093 5.093 0 0 0 2.7.748a5.39 5.39 0 0 0 2.066-.408a5.328 5.328 0 0 0 3.303-4.927a5.354 5.354 0 0 0-5.357-5.346a5.166 5.166 0 0 0-3.795 1.563a8.098 8.098 0 0 0-.859.97l-.022.023a1.183 1.183 0 0 1-.145.173a.704.704 0 0 1-.24.106a.677.677 0 0 1-.43-.034a1.104 1.104 0 0 1-.407-.184a.366.366 0 0 1-.173-.296V2.424a.426.426 0 0 1 .117-.296A.396.396 0 0 1 7.473 2h9.788q.335 0 .335.614t-.335.614H8.209v5.39h.012a6.958 6.958 0 0 1 1.138-.937A7.006 7.006 0 0 1 10.564 7a6.563 6.563 0 0 1 2.578-.513A6.374 6.374 0 0 1 15.687 7a6.524 6.524 0 0 1 4.007 6.038Zm-.346-6.484a.264.264 0 0 1 .1.2a.346.346 0 0 1-.061.201q-.062.09-.184.235q-.29.29-.436.29a.246.246 0 0 1-.178-.078a8.79 8.79 0 0 0-2.31-1.485a7.6 7.6 0 0 0-3.08-.625a8.175 8.175 0 0 0-2.925.547q-.301.112-.502-.413a1.125 1.125 0 0 1-.09-.424a.26.26 0 0 1 .18-.223a8.24 8.24 0 0 1 3.336-.636a9 9 0 0 1 3.527.714a8.57 8.57 0 0 1 2.623 1.697Z"/>`
	flipHInnerSVG                = `<path fill="currentColor" d="M21 13H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2Z" opacity=".25"/><path fill="currentColor" d="m12 21l5-5H7l5 5z"/><path fill="currentColor" d="M12 22a.997.997 0 0 1-.707-.293l-5-5A1 1 0 0 1 7 15h10a1 1 0 0 1 .707 1.707l-5 5A.997.997 0 0 1 12 22Zm-2.586-5L12 19.586L14.586 17Z"/><path fill="currentColor" d="M12 4a.99.99 0 0 1-.92-1.38a1.03 1.03 0 0 1 .21-.33a.998.998 0 0 1 1.09-.21a1.034 1.034 0 0 1 .33.21A1.052 1.052 0 0 1 13 3a.838.838 0 0 1-.08.38a1.171 1.171 0 0 1-.21.33A.992.992 0 0 1 12 4zM8.79 6.21a1.008 1.008 0 0 1 0-1.42a1.007 1.007 0 0 1 1.42 0a1.008 1.008 0 0 1 0 1.42a1.027 1.027 0 0 1-.71.29a1.026 1.026 0 0 1-.71-.29zM7 9a.99.99 0 0 1-1-1a1 1 0 0 1 1.71-.71A1.052 1.052 0 0 1 8 8a.99.99 0 0 1-1 1zm5.67-1a.997.997 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a.997.997 0 0 1-1-1zM9.33 8a1.003 1.003 0 0 1 1-1a.997.997 0 0 1 1 1a.997.997 0 0 1-1 1a1.003 1.003 0 0 1-1-1zM17 9a.99.99 0 0 1-1-1a1 1 0 1 1 2 0a.99.99 0 0 1-1 1zm-3.21-2.79a1.008 1.008 0 0 1 0-1.42a1.007 1.007 0 0 1 1.42 0a1.008 1.008 0 0 1 0 1.42a1.027 1.027 0 0 1-.71.29a1.026 1.026 0 0 1-.71-.29z" opacity=".5"/>`
	flipHAltInnerSVG             = `<path fill="currentColor" d="M20 11a.99.99 0 0 1-.71-.29a1.16 1.16 0 0 1-.21-.33a.941.941 0 0 1 0-.76a1.029 1.029 0 0 1 .21-.33a1.047 1.047 0 0 1 1.42 0a1.147 1.147 0 0 1 .21.33a.941.941 0 0 1 0 .76a1.16 1.16 0 0 1-.21.33a.993.993 0 0 1-.71.29zm-1-4.5a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1zM20 4a.99.99 0 0 1-.92-1.38a1.029 1.029 0 0 1 .21-.33a1.047 1.047 0 0 1 1.42 0a1.029 1.029 0 0 1 .21.33a.99.99 0 0 1-.21 1.09a1.16 1.16 0 0 1-.33.21A1 1 0 0 1 20 4zM7.03 6.24a.994.994 0 0 1 .73-1.21a.99.99 0 0 1 1.21.73a.999.999 0 0 1-.73 1.21A.965.965 0 0 1 8 7a.991.991 0 0 1-.97-.76zm4-1a.994.994 0 0 1 .73-1.21a.99.99 0 0 1 1.21.73a.999.999 0 0 1-.73 1.21A.965.965 0 0 1 12 6a1.003 1.003 0 0 1-.97-.76zm4-1a.998.998 0 0 1 .73-1.21a1.003 1.003 0 0 1 1.21.73a.999.999 0 0 1-.73 1.21A.965.965 0 0 1 16 5a.991.991 0 0 1-.97-.76zM4 8a.99.99 0 0 1-1-1a.832.832 0 0 1 .08-.38a1.029 1.029 0 0 1 .21-.33a1.047 1.047 0 0 1 1.42 0a1.029 1.029 0 0 1 .21.33A.99.99 0 0 1 4 8zm0 3a.99.99 0 0 1-.71-.29a1.16 1.16 0 0 1-.21-.33a.941.941 0 0 1 0-.76a1.147 1.147 0 0 1 .21-.33a1.047 1.047 0 0 1 1.42 0a1.147 1.147 0 0 1 .21.33a.941.941 0 0 1 0 .76a1.16 1.16 0 0 1-.21.33A.993.993 0 0 1 4 11zm11-1a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1zm-4 0a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1zm-4 0a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1z" opacity=".5"/><path fill="currentColor" d="M20 14v7L4 17v-3h16z"/><path fill="currentColor" d="M20 22a.974.974 0 0 1-.242-.03l-16-4A1 1 0 0 1 3 17v-3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1ZM5 16.219l14 3.5V15H5Z"/>`
	flipVInnerSVG                = `<path fill="currentColor" d="M12 22a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v18a1 1 0 0 1-1 1Z" opacity=".25"/><path fill="currentColor" d="m21 12l-5-5v10l5-5z"/><path fill="currentColor" d="M16 18a1 1 0 0 1-1-1V7a1 1 0 0 1 1.707-.707l5 5a1 1 0 0 1 0 1.414l-5 5A.999.999 0 0 1 16 18Zm1-8.586v5.172L19.586 12Z"/><path fill="currentColor" d="M3 13a.99.99 0 0 1-.92-1.38a1.149 1.149 0 0 1 .21-.33a1.047 1.047 0 0 1 1.42 0a1.037 1.037 0 0 1 .21.33A.838.838 0 0 1 4 12a.99.99 0 0 1-1 1zm1.79 2.21a1.008 1.008 0 0 1 0-1.42a1.007 1.007 0 0 1 1.42 0a1.008 1.008 0 0 1 0 1.42a1.027 1.027 0 0 1-.71.29a1.026 1.026 0 0 1-.71-.29zM8 18a.992.992 0 0 1-.71-.29a1.162 1.162 0 0 1-.21-.33A1 1 0 0 1 7 17a1.05 1.05 0 0 1 .29-.71a1.039 1.039 0 0 1 1.09-.21a1.154 1.154 0 0 1 .33.21A1.052 1.052 0 0 1 9 17a1 1 0 0 1-.08.38a1.171 1.171 0 0 1-.21.33A.992.992 0 0 1 8 18zm-1-4.33a.996.996 0 0 1 1-1a.996.996 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1zm0-3.34a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a.997.997 0 0 1-1 1a.997.997 0 0 1-1-1zM8 8a.99.99 0 0 1-.92-1.38a1.03 1.03 0 0 1 .21-.33a.985.985 0 0 1 1.63.33A.99.99 0 0 1 8 8zm-3.21 2.21a1.008 1.008 0 0 1 0-1.42a1.007 1.007 0 0 1 1.42 0a1.008 1.008 0 0 1 0 1.42a1.027 1.027 0 0 1-.71.29a1.026 1.026 0 0 1-.71-.29z" opacity=".5"/>`
	flipVAltInnerSVG             = `<path fill="currentColor" d="M10 21a.99.99 0 0 1-.71-.29a1.16 1.16 0 0 1-.21-.33A.832.832 0 0 1 9 20a1 1 0 1 1 2 0a.99.99 0 0 1-1 1zm-4.5-1a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1zM3 21a.99.99 0 0 1-.71-.29a1.16 1.16 0 0 1-.21-.33A1 1 0 0 1 2 20a1.048 1.048 0 0 1 .29-.71a1.047 1.047 0 0 1 1.42 0A1.052 1.052 0 0 1 4 20a1 1 0 0 1-.08.38a1.16 1.16 0 0 1-.21.33A.993.993 0 0 1 3 21zm.76-4.03a.998.998 0 0 1-.73-1.21a.994.994 0 0 1 1.21-.73a.999.999 0 0 1 .73 1.21A.992.992 0 0 1 4 17a.965.965 0 0 1-.24-.03zm.99-4a.992.992 0 0 1-.72-1.21a.998.998 0 0 1 1.21-.73a.999.999 0 0 1 .73 1.21A1 1 0 0 1 5 13a1.104 1.104 0 0 1-.25-.03zm1.01-4a.998.998 0 0 1-.73-1.21a.994.994 0 0 1 1.21-.73a.999.999 0 0 1 .73 1.21A.992.992 0 0 1 6 9a.965.965 0 0 1-.24-.03zM7 5a.99.99 0 0 1-1-1a1.048 1.048 0 0 1 .29-.71a1.034 1.034 0 0 1 1.41 0A1.018 1.018 0 0 1 8 4a1.007 1.007 0 0 1-1 1zm3 0a.99.99 0 0 1-1-1a1.048 1.048 0 0 1 .29-.71a1.047 1.047 0 0 1 1.42 0A1.052 1.052 0 0 1 11 4a.99.99 0 0 1-1 1zM9 16a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1zm0-4a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1zm0-4a1.003 1.003 0 0 1 1-1a1.003 1.003 0 0 1 1 1a1.003 1.003 0 0 1-1 1a1.003 1.003 0 0 1-1-1z" opacity=".5"/><path fill="currentColor" d="M14 20h7L17 4h-3v16z"/><path fill="currentColor" d="M21 21h-7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h3a1 1 0 0 1 .97.758l4 16A1 1 0 0 1 21 21Zm-6-2h4.719l-3.5-14H15Z"/>`
	githubInnerSVG               = `<path fill="currentColor" d="M8.854 21.57a9.91 9.91 0 0 0 6.29.001a.493.493 0 0 1-.644-.475c0-.338.013-1.413.013-2.75a2.368 2.368 0 0 0-.675-1.85c2.225-.25 4.562-1.1 4.562-4.938a3.87 3.87 0 0 0-1.025-2.687a3.594 3.594 0 0 0-.1-2.65s-.838-.275-2.75 1.025a9.427 9.427 0 0 0-5 0C7.612 5.958 6.775 6.22 6.775 6.22a3.593 3.593 0 0 0-.1 2.65a3.892 3.892 0 0 0-1.025 2.687c0 3.825 2.325 4.688 4.55 4.938a2.106 2.106 0 0 0-.638 1.337a2.137 2.137 0 0 1-2.91-.82l-.002-.005a2.001 2.001 0 0 0-1.538-1.025c-.837.013-.337.475.013.663c.451.38.803.865 1.025 1.412c.2.563.85 1.638 3.362 1.175c0 .838.013 1.625.013 1.863c0 .259-.185.551-.67.475z" opacity=".5"/><path fill="currentColor" d="M12 2.083c-5.523 0-10 4.477-10 10c0 4.423 2.875 8.169 6.855 9.488c.485.075.67-.216.67-.475c0-.238-.012-1.025-.012-1.863c-2.513.463-3.163-.612-3.363-1.175a3.637 3.637 0 0 0-1.025-1.412c-.35-.188-.85-.65-.013-.663a2 2 0 0 1 1.538 1.025l.003.006a2.137 2.137 0 0 0 2.91.82c.043-.51.27-.984.637-1.338c-2.225-.25-4.55-1.113-4.55-4.938a3.892 3.892 0 0 1 1.025-2.687a3.594 3.594 0 0 1 .1-2.65s.837-.263 2.75 1.025a9.427 9.427 0 0 1 5 0c1.912-1.3 2.75-1.025 2.75-1.025c.37.838.406 1.786.1 2.65a3.87 3.87 0 0 1 1.025 2.687c0 3.838-2.338 4.688-4.562 4.938c.482.49.729 1.164.675 1.85c0 1.337-.013 2.412-.013 2.75a.493.493 0 0 0 .643.476C19.124 20.253 22 16.507 22 12.083c0-5.523-4.477-10-10-10z"/>`
	githubAltInnerSVG            = `<path fill="currentColor" d="M20.116 5.901a3.883 3.883 0 0 0-.26-.31a4.413 4.413 0 0 0 .21-.76a5.284 5.284 0 0 0-.35-2.8s-1.12-.35-3.69 1.38a12.477 12.477 0 0 0-3.35-.45a12.604 12.604 0 0 0-3.36.45c-2.57-1.75-3.69-1.38-3.69-1.38a5.263 5.263 0 0 0-.35 2.77a4.21 4.21 0 0 0 .22.79c-.09.1-.18.21-.26.31a5.14 5.14 0 0 0-1.12 3.3a7.686 7.686 0 0 0 .04.85c.32 4.43 3.27 5.46 6.08 5.78a2.558 2.558 0 0 0-.77 1.39a4.022 4.022 0 0 0-.13 1.09v1.31c-1.119.1-2.267-.063-2.623-1.061a5.695 5.695 0 0 0-1.834-2.413a1.179 1.179 0 0 1-.17-.112a1.001 1.001 0 0 0-.93-.645h-.005a1 1 0 0 0-1 .995c-.003.812.81 1.337 1.143 1.515a4.466 4.466 0 0 1 .923 1.359c.364 1.023 1.429 2.578 4.466 2.376l.002.098l.004.268a1 1 0 0 0 1 1h4.714a1 1 0 0 0 1-1s.008-3.16.008-3.69a4.024 4.024 0 0 0-.13-1.09l-.002-.006l.004.006c-.009-.035-.022-.063-.032-.097a2.532 2.532 0 0 0-.74-1.293l.012.021l-.02-.02c2.81-.32 5.74-1.37 6.06-5.78a7.687 7.687 0 0 0 .04-.85a5.23 5.23 0 0 0-1.11-3.3Z"/>`
	gitlabInnerSVG               = `<path fill="currentColor" d="m21.94 12.865l-1.066-3.28l.001.005v-.005L18.76 3.076a.833.833 0 0 0-.799-.57a.822.822 0 0 0-.788.575l-2.007 6.178H8.834L6.824 3.08a.822.822 0 0 0-.788-.575H6.03a.839.839 0 0 0-.796.575L3.127 9.584l-.002.006l.001-.005l-1.068 3.28a1.195 1.195 0 0 0 .434 1.34l9.229 6.705l.004.003l.012.008l-.011-.008l.002.001l.001.001a.475.475 0 0 0 .045.028l.006.004l.004.002l.003.001h.002l.006.003l.024.01l.023.01h.001l.004.002l.005.002h.002l.007.002h.002c.011.004.023.006.034.009l.013.003h.002l.005.002l.007.001h.007a.465.465 0 0 0 .066.006h.001a.467.467 0 0 0 .067-.005h.007l.007-.002l.004-.001h.002l.014-.004l.034-.008h.002l.007-.003h.002l.005-.002l.003-.001h.001l.025-.011l.023-.01l.005-.002h.002l.003-.002l.004-.002l.007-.004a.478.478 0 0 0 .044-.027l.004-.003l.005-.003l9.23-6.706a1.195 1.195 0 0 0 .434-1.339Zm-3.973-9.18l1.81 5.574h-3.62Zm1.493 6.516l-.738.947l-5.448 6.98l2.577-7.927Zm-7.91 10.474l.001.004Zm-.827-2.546L4.54 10.2h3.61ZM6.03 3.686l1.813 5.573h-3.62Zm-2.984 9.756a.255.255 0 0 1-.092-.285l.794-2.438l5.822 7.464Zm8.659 7.456l-.006-.005l-.011-.01l-.02-.018l.002.001l.002.002a.478.478 0 0 0 .042.037l.003.002Zm.293-1.894l-1.514-4.665l-1.343-4.138h5.719Zm.31 1.88l-.01.008l-.002.001l-.005.005l-.012.009l.002-.002a.46.46 0 0 0 .043-.036l.001-.002l.002-.002Zm8.643-7.442l-6.523 4.74l5.824-7.463l.791 2.437a.255.255 0 0 1-.092.286Z"/>`
	gitlabAltInnerSVG            = `<path fill="currentColor" d="m12 21.42l3.684-11.333H8.32L12 21.421Z"/><path fill="currentColor" d="m3.16 10.087l-1.123 3.444a.763.763 0 0 0 .277.852l9.685 7.038l-8.84-11.334Z" opacity=".25"/><path fill="currentColor" d="M3.16 10.087h5.16L6.1 3.262a.383.383 0 0 0-.728 0L3.16 10.087Z"/><path fill="currentColor" d="m20.845 10.087l1.118 3.444a.763.763 0 0 1-.276.852l-9.688 7.038l8.846-11.334Z" opacity=".25"/><path fill="currentColor" d="M20.845 10.087h-5.161L17.9 3.262a.383.383 0 0 1 .727 0l2.217 6.825Z"/><path fill="currentColor" d="m11.999 21.421l3.685-11.334h5.161l-8.846 11.334z" opacity=".5"/><path fill="currentColor" d="m11.999 21.421l-8.84-11.334H8.32l3.679 11.334z" opacity=".5"/>`
	googleInnerSVG               = `<path fill="currentColor" d="M12.222 5.977a5.402 5.402 0 0 1 3.823 1.496l2.868-2.868A9.61 9.61 0 0 0 12.222 2a9.996 9.996 0 0 0-8.937 5.51l3.341 2.59a5.96 5.96 0 0 1 5.596-4.123z" opacity=".7"/><path fill="currentColor" d="M3.285 7.51a10.013 10.013 0 0 0 0 8.98l3.341-2.59a5.913 5.913 0 0 1 0-3.8l-3.34-2.59z"/><path fill="currentColor" d="M15.608 17.068A6.033 6.033 0 0 1 6.626 13.9l-3.34 2.59A9.996 9.996 0 0 0 12.221 22a9.547 9.547 0 0 0 6.618-2.423l-3.232-2.509z" opacity=".5"/><path fill="currentColor" d="M21.64 10.182h-9.418v3.868h5.382a4.6 4.6 0 0 1-1.996 3.018l-.01.006l.01-.006l3.232 2.51a9.753 9.753 0 0 0 2.982-7.35c0-.687-.06-1.371-.182-2.046z" opacity=".25"/>`
	googleDriveInnerSVG          = `<path fill="currentColor" d="M15.334 14.887H22l-6.666-11.55H8.667l6.667 11.55z" opacity=".25"/><path fill="currentColor" d="M8.667 3.338L2 14.887l3.334 5.775L12 9.113z"/><path fill="currentColor" d="m8.667 14.887l-3.333 5.775h13.333L22 14.887z" opacity=".5"/>`
	googleDriveAltInnerSVG       = `<path fill="currentColor" d="m6 19.796l3-5.197h12l-3 5.197H6z"/><path fill="currentColor" d="M15 14.599h6L15 4.204H9L15 14.6z" opacity=".25"/><path fill="currentColor" d="m3 14.599l3 5.197L12 9.4L9 4.204L3 14.6z" opacity=".5"/>`
	googleHangoutsInnerSVG       = `<path fill="currentColor" d="M9.818 14.313a3.795 3.795 0 0 1-.447.085a.21.21 0 0 1-.235-.152v-.968c.006-.095-.032-.197.044-.28a.266.266 0 0 1 .138-.085c.173-.032.34-.088.495-.169c.273-.158.468-.423.538-.731l.035-.151l-2.041-.004a.65.65 0 0 1-.175-.018a.413.413 0 0 1-.306-.385c-.001-1.029 0-2.058.001-3.087a.467.467 0 0 1 .139-.436a.39.39 0 0 1 .256-.102h3.002a.416.416 0 0 1 .4.527c.015.05.02.102.016.153c0 1.122.006 2.243-.002 3.364a2.51 2.51 0 0 1-.84 1.878c-.029.026-.056.054-.084.08c-.277.22-.594.384-.934.481zm5.536-.071l-.038.01l-.022.013c-.193.056-.39.099-.589.129a.216.216 0 0 1-.27-.217l-.002-.322l.005-.578l-.002-.127c0-.147.057-.21.24-.252c.235-.04.454-.142.634-.297a1.26 1.26 0 0 0 .378-.74l-1.886-.003h-.205a.419.419 0 0 1-.433-.44c0-1.016.001-2.033.003-3.049l-.002-.03a.435.435 0 0 1 .204-.45a.39.39 0 0 1 .212-.06h2.982a.416.416 0 0 1 .4.528c.015.05.02.101.016.153v3.37a2.557 2.557 0 0 1-1.557 2.327c-.023.01-.045.023-.068.035z" opacity=".5"/><path fill="currentColor" d="M20.866 9.055a8.505 8.505 0 0 0-2.173-4.334a8.565 8.565 0 0 0-5.263-2.667A4.157 4.157 0 0 1 12.95 2h-1.135c-.013.029-.039.018-.06.02a8.818 8.818 0 0 0-.87.114a8.453 8.453 0 0 0-5.177 3.104a8.358 8.358 0 0 0-1.84 4.709a8.59 8.59 0 0 0 .185 2.529c.138.618.347 1.22.624 1.79l.074.146a8.601 8.601 0 0 0 1.783 2.396a8.53 8.53 0 0 0 5.763 2.333c.1.001.128.026.127.128c-.005.874-.003 1.747-.002 2.621c0 .034.002.068.004.11l.063-.026a17.99 17.99 0 0 0 4.49-2.966c.268-.243.524-.497.77-.763l.315-.341c.163-.175.304-.37.457-.553c.178-.212.337-.437.494-.663c.35-.499.657-1.025.919-1.575c.132-.279.257-.56.369-.847a.097.097 0 0 0 .01-.016c.3-.786.502-1.605.601-2.44a12.1 12.1 0 0 0 .075-.877a8.142 8.142 0 0 0-.123-1.878zm-9.19 2.82a2.51 2.51 0 0 1-.84 1.877c-.029.026-.056.054-.084.08c-.277.22-.594.384-.934.481a3.791 3.791 0 0 1-.447.085a.21.21 0 0 1-.235-.152v-.968c.006-.095-.032-.197.044-.28a.265.265 0 0 1 .138-.085c.173-.032.339-.088.495-.169c.273-.158.468-.423.538-.731l.035-.151c-.68-.001-1.36-.003-2.041-.003a.65.65 0 0 1-.175-.019a.413.413 0 0 1-.306-.385a870.3 870.3 0 0 1 .001-3.087a.467.467 0 0 1 .139-.436c.07-.064.161-.1.256-.102h3.002a.416.416 0 0 1 .4.527c.015.05.02.102.016.153c0 1.122.006 2.243-.002 3.364zm5.303.004a2.557 2.557 0 0 1-1.557 2.328c-.023.01-.045.023-.068.035l-.038.01l-.022.013c-.193.056-.39.099-.589.129a.216.216 0 0 1-.27-.217l-.002-.322l.005-.578l-.002-.127c0-.147.057-.21.24-.252c.235-.04.454-.142.634-.297a1.26 1.26 0 0 0 .378-.74l-1.886-.003h-.205a.419.419 0 0 1-.433-.44c0-1.016.001-2.033.003-3.049l-.002-.03a.435.435 0 0 1 .204-.45a.39.39 0 0 1 .212-.06h2.982a.416.416 0 0 1 .4.528c.015.05.02.101.016.153v3.37z"/>`
	googleHangoutsAltInnerSVG    = `<path fill="currentColor" d="M12 23a1 1 0 0 1-1-1v-1.477a9.842 9.842 0 1 1 10.705-8.527a12.531 12.531 0 0 1-9.466 10.975A.997.997 0 0 1 12 23Z" opacity=".5"/><path fill="currentColor" d="M9 11.689a2 2 0 0 1-2-2a2 2 0 0 1 2-2a2 2 0 0 1 2 2a2 2 0 0 1-2 2Z"/><path fill="currentColor" d="M8.515 14.688a1 1 0 0 1 0-2a.501.501 0 0 0 .5-.5v-2.5a1 1 0 0 1 2 0v2.5a2.502 2.502 0 0 1-2.5 2.5zm6.485-3a2 2 0 0 1-2-2a2 2 0 0 1 2-2a2 2 0 0 1 2 2a2 2 0 0 1-2 2z"/><path fill="currentColor" d="M14.515 14.688a1 1 0 0 1 0-2a.501.501 0 0 0 .5-.5v-2.5a1 1 0 0 1 2 0v2.5a2.502 2.502 0 0 1-2.5 2.5Z"/>`
	googlePlayInnerSVG           = `<path fill="currentColor" d="m14.556 12.895l2.573 2.554l3.785-2.186c.506-.253.83-.766.842-1.332a1.474 1.474 0 0 0-.837-1.275a511.68 511.68 0 0 1-4.023-2.323l-.002-.002L4.639 1.256a1.644 1.644 0 0 0-1.657-.06c-.05.034-.097.071-.142.111l11.716 11.588z" opacity=".5"/><path fill="currentColor" d="M20.919 10.656a505.13 505.13 0 0 1-3.791-2.188l-2.597 2.577L2.785 22.698c.048.044.1.084.153.12c.21.12.446.182.687.181c.292-.004.578-.084.83-.232c.389-.226 12.439-7.182 12.439-7.182l4.02-2.322c.506-.253.83-.766.842-1.332a1.474 1.474 0 0 0-.837-1.275z" opacity=".7"/><path fill="currentColor" d="M13.61 11.96L2.84 1.309a1.509 1.509 0 0 0-.597 1.252v18.985c-.013.452.197.871.543 1.15L13.61 11.96z" opacity=".25"/><path fill="currentColor" d="M20.919 10.656c-.744-.4-3.37-1.944-3.791-2.188l-2.597 2.577l-.921.914l.946.936l2.573 2.554l3.785-2.186c.506-.253.83-.766.842-1.332a1.474 1.474 0 0 0-.837-1.275z"/>`
	graphBarInnerSVG             = `<path fill="currentColor" d="M6 23H2a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1Z" opacity=".25"/><path fill="currentColor" d="M14 23h-4a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v20a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M22 23h-4a1 1 0 0 1-1-1V10a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1Z" opacity=".5"/>`
	gridInnerSVG                 = `<path fill="currentColor" d="M22 6H2V3a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v3z"/><path fill="currentColor" d="M2 8h9v6H2zm0 8h9v6H3a1 1 0 0 1-1-1v-5zm11-8h9v6h-9zm8 14h-8v-6h9v5a1 1 0 0 1-1 1z" opacity=".5"/><path fill="currentColor" d="M22 8V6H2v2h9v6H2v2h9v6h2v-6h9v-2h-9V8z" opacity=".25"/>`
	gridsInnerSVG                = `<path fill="currentColor" d="M23 21V3a1 1 0 0 0-1-1h-5v20h5a1 1 0 0 0 1-1Z" opacity=".5"/><path fill="currentColor" d="M1 3v18a1 1 0 0 0 1 1h5V2H2a1 1 0 0 0-1 1Z"/><path fill="currentColor" d="M9 2h6v20H9z" opacity=".5"/><path fill="currentColor" d="M7 2h2v20H7zm8 0h2v20h-2z" opacity=".25"/>`
	gripHorizontalLineInnerSVG   = `<path fill="currentColor" d="M21 11H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2z"/>`
	headSideInnerSVG             = `<path fill="currentColor" d="M18.5 22h-9a1 1 0 0 1-1-1v-2h-1a2.002 2.002 0 0 1-2-2v-2h-1a1 1 0 0 1-.904-1.426L5.5 9.53V9.5a7.44 7.44 0 0 1 2.277-5.383a7.365 7.365 0 0 1 5.453-2.114a7.698 7.698 0 0 1 7.27 7.77V10a1.031 1.031 0 0 1-.033.256l-1.93 7.266l.924 3.2A1 1 0 0 1 18.5 22Z" opacity=".5"/>`
	headSideCoughInnerSVG        = `<circle cx="6" cy="16.999" r="1" fill="currentColor"/><circle cx="2" cy="18" r="1" fill="currentColor"/><circle cx="5" cy="21" r="1" fill="currentColor"/><path fill="currentColor" d="M21.13 21h-8.463a1 1 0 0 1-1-1v-2H10.8a1.935 1.935 0 0 1-1.934-1.934v-1.8H8a1 1 0 0 1-.904-1.426l1.77-3.758v-.016a7.067 7.067 0 0 1 7.284-7.063A7.252 7.252 0 0 1 23 9.321v.212a1.031 1.031 0 0 1-.033.257l-1.796 6.767l.919 3.164A1 1 0 0 1 21.13 21Z" opacity=".5"/>`
	headSideMaskInnerSVG         = `<path fill="currentColor" d="M20.476 9.287L12.337 12H4.5a1 1 0 0 0-.938 1.348l1.448 3.895a.87.87 0 0 0 .043.099A2.984 2.984 0 0 0 7.736 19H12.5a1 1 0 0 0 .32-.052l5.864-1.978l.616-2.319l-5.8 1.956v-2.886l6.634-2.211l.333-1.254A1.031 1.031 0 0 0 20.5 10v-.228c0-.163-.014-.324-.024-.485Z"/><path fill="currentColor" d="M3.508 12.894v.03v-.03zm.17-.465A1.001 1.001 0 0 1 4.5 12h7.837l8.14-2.713a7.676 7.676 0 0 0-7.247-7.284a7.359 7.359 0 0 0-5.453 2.114A7.441 7.441 0 0 0 5.5 9.465l-1.852 3.011c-.005.009-.004.02-.01.03c.015-.026.024-.053.04-.077zm9.822 1.292v2.886l5.8-1.956l1.167-4.395l-.333 1.254l-6.634 2.211zm5.037 3.801l.2-.754l-.053.202l-5.865 1.978A1 1 0 0 1 12.5 19h-4v2a1 1 0 0 0 1 1h9a1 1 0 0 0 .96-1.277Z" opacity=".5"/>`
	hipchatInnerSVG              = `<path fill="currentColor" d="M17.544 19.985a.192.192 0 0 1 .048.116c0 .058-.055.098-.124.098a7.539 7.539 0 0 1-2.857-1.518a.538.538 0 0 0-.488-.08a9.98 9.98 0 0 1-2.123.227c-4.597 0-8.323-3.076-8.323-6.873c0-3.794 3.726-6.87 8.323-6.87s8.323 3.076 8.323 6.87a6.566 6.566 0 0 1-3.374 5.526a.546.546 0 0 0-.282.438a4.667 4.667 0 0 0 .877 2.066Z" opacity=".25"/><path fill="currentColor" d="M17.462 13.869a.4.4 0 0 0-.399-.398a.392.392 0 0 0-.263.101a7.32 7.32 0 0 1-4.796 1.711h-.008a7.623 7.623 0 0 1-4.796-1.71a.497.497 0 0 0-.272-.094a.405.405 0 0 0-.39.44a.932.932 0 0 0 .228.481a6.445 6.445 0 0 0 5.212 2.408h.044a6.445 6.445 0 0 0 5.212-2.408a1.09 1.09 0 0 0 .228-.531"/><path fill="currentColor" d="M17.544 19.985a.192.192 0 0 1 .048.116c0 .058-.055.098-.124.098a7.539 7.539 0 0 1-2.857-1.518a.538.538 0 0 0-.488-.08a9.98 9.98 0 0 1-2.123.227c-4.597 0-8.323-3.076-8.323-6.873c0-3.794 3.726-6.87 8.323-6.87s8.323 3.076 8.323 6.87a6.566 6.566 0 0 1-3.374 5.526a.546.546 0 0 0-.282.438a4.667 4.667 0 0 0 .877 2.066Zm2.676.839a3.672 3.672 0 0 1-1.684-2.015a.369.369 0 0 1 .134-.44A8.144 8.144 0 0 0 22 11.941c0-4.765-4.477-8.627-10-8.627S2 7.176 2 11.941c0 4.767 4.477 8.63 10 8.63a11.567 11.567 0 0 0 2.104-.192a.565.565 0 0 1 .412.077A10.758 10.758 0 0 0 19.71 22a.703.703 0 0 0 .798-.698a.54.54 0 0 0-.288-.478Z"/>`
	historyInnerSVG              = `<path fill="currentColor" d="M12 2a10.017 10.017 0 0 0-6.994 2.872V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2H6.218A7.98 7.98 0 1 1 4 12a1 1 0 0 0-2 0A10 10 0 1 0 12 2Z"/><path fill="currentColor" d="M14 13h-2a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v2h1a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M12 4a8.008 8.008 0 0 0-5.782 2.5h2.288a1 1 0 0 1 0 2h-4.5a.99.99 0 0 1-.978-.889A9.922 9.922 0 0 0 2 12a1 1 0 0 1 2 0a8 8 0 1 0 8-8Zm2 9h-2a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v2h1a1 1 0 0 1 0 2Z" opacity=".5"/>`
	historyAltInnerSVG           = `<path fill="currentColor" d="M12 2a10.017 10.017 0 0 0-7 2.877V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2H6.218A7.992 7.992 0 1 1 12 20a1 1 0 0 0 0 2a10 10 0 0 0 0-20Z"/><path fill="currentColor" d="M14 13h-2a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v2h1a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M12 4a8.008 8.008 0 0 0-5.782 2.5H8.5a1 1 0 0 1 0 2H4a.989.989 0 0 1-.976-.88A9.977 9.977 0 0 0 12 22a1 1 0 0 1 0-2a8 8 0 0 0 0-16Zm2 9h-2a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v2h1a1 1 0 0 1 0 2Z" opacity=".5"/>`
	horizontalAlignLeftInnerSVG  = `<path fill="currentColor" d="M16 10H4V6h11a1 1 0 0 1 1 1v3z" opacity=".5"/><path fill="currentColor" d="M21 18H4v-8h17a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M3 22a1 1 0 0 1-1-.999V3a1 1 0 0 1 2 0v18a1 1 0 0 1-.999 1H3z" opacity=".25"/>`
	hospitalInnerSVG             = `<path fill="currentColor" d="M12.5 14.5h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2zm0 4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2zm-5-4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2zm0 4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2zm10-4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2zm0 4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M13.5 7H13v-.5a1 1 0 0 0-2 0V7h-.5a1 1 0 0 0 0 2h.5v.5a1 1 0 0 0 2 0V9h.5a1 1 0 0 0 0-2Z"/><path fill="currentColor" d="M21.5 6.5h-3v-4a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1v4h-3a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h19a1 1 0 0 0 1-1v-14a1 1 0 0 0-1-1Zm-14 12h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2Zm0-4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2Zm5 4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2Zm0-4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2Zm1-5.5H13v.5a1 1 0 0 1-2 0V9h-.5a1 1 0 0 1 0-2h.5v-.5a1 1 0 0 1 2 0V7h.5a1 1 0 0 1 0 2Zm4 9.5h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2Zm0-4h-1a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2Z" opacity=".25"/>`
	hospitalSquareSignInnerSVG   = `<path fill="currentColor" d="M15 6a1 1 0 0 0-1 1v4h-4V7a1 1 0 0 0-2 0v10a1 1 0 0 0 2 0v-4h4v4a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V5a3.003 3.003 0 0 0-3-3Zm-3 15a1 1 0 0 1-2 0v-4h-4v4a1 1 0 0 1-2 0V7a1 1 0 0 1 2 0v4h4V7a1 1 0 0 1 2 0Z" opacity=".5"/>`
	hospitalSymbolInnerSVG       = `<path fill="currentColor" d="M15 7a1 1 0 0 0-1 1v3h-4V8a1 1 0 0 0-2 0v8a1 1 0 0 0 2 0v-3h4v3a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2Zm4 14a1 1 0 0 1-2 0v-3h-4v3a1 1 0 0 1-2 0V8a1 1 0 0 1 2 0v3h4V8a1 1 0 0 1 2 0Z" opacity=".5"/>`
	houseUserInnerSVG            = `<path fill="currentColor" d="M12 18a3.5 3.5 0 1 1 3.5-3.5A3.504 3.504 0 0 1 12 18Z" opacity=".5"/><path fill="currentColor" d="M14.64 16.772a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 7 21a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228Z"/><path fill="currentColor" d="M21 12a.996.996 0 0 1-.664-.252L12 4.338l-8.336 7.41a1 1 0 0 1-1.328-1.496l9-8a.999.999 0 0 1 1.328 0l9 8A1 1 0 0 1 21 12Z" opacity=".5"/><path fill="currentColor" d="m12 4.338l-8 7.111V21a1 1 0 0 0 1 1h3a1 1 0 0 1-1-1a4.988 4.988 0 0 1 2.36-4.228A3.469 3.469 0 0 1 8.5 14.5a3.5 3.5 0 0 1 7 0a3.469 3.469 0 0 1-.86 2.272A4.988 4.988 0 0 1 17 21a1 1 0 0 1-1 1h3a1 1 0 0 0 1-1v-9.551Z" opacity=".25"/>`
	htmlFiveInnerSVG             = `<path fill="currentColor" d="m3.182 2l1.605 18l7.202 2l7.222-2.002L20.818 2Zm14.143 5.887H8.877l.202 2.261h8.045l-.606 6.778L12 18.178l-.01.003l-4.523-1.255l-.309-3.466h2.216l.158 1.76l2.458.664h.002l2.463-.665l.256-2.863H7.06L6.464 5.68h11.059Z"/><path fill="currentColor" d="M17.325 7.887H8.877l.202 2.261h8.045l-.606 6.778L12 18.178l-.01.003l-4.523-1.255l-.309-3.466h2.216l.158 1.76l2.458.664h.002l2.463-.665l.256-2.863H7.06L6.464 5.68h11.059Z" opacity=".25"/>`
	htmlFiveAltInnerSVG          = `<path fill="currentColor" d="M21.468 2.325A1 1 0 0 0 20.73 2H3.27a1 1 0 0 0-.996 1.089l1.52 17a1 1 0 0 0 .728.874l7.2 2a.996.996 0 0 0 .268.037a1.012 1.012 0 0 0 .267-.036l7.22-2a1 1 0 0 0 .729-.875l1.52-17a1 1 0 0 0-.258-.764Z"/><path fill="currentColor" d="M7.82 13h6.895l-.327 3.271l-2.568.917l-3.164-1.13a1 1 0 1 0-.673 1.884l3.5 1.25a1.003 1.003 0 0 0 .673 0l3.5-1.25a1 1 0 0 0 .659-.842l.5-5a1 1 0 0 0-.995-1.1H8.725l-.3-3h7.895a1 1 0 0 0 0-2h-9a1 1 0 0 0-.995 1.1l.5 5a1 1 0 0 0 .995.9Z" opacity=".25"/>`
	htmlThreeInnerSVG            = `<path fill="currentColor" d="m17.476 6.123l-.534 5.994l.002.033l-.002.074v-.001l-.38 4.192l-.041.372L12 18.037v.001l-.004.003l-4.512-1.258l-.306-3.465h2.213l.157 1.762l2.453.665l-.001.001l2.461-.675l.261-2.869H9.576l-.044-.485l-.101-1.136l-.052-.611h5.538l.202-2.231H6.682l-.044-.485l-.1-1.137l-.053-.61h11.044z" opacity=".5"/><path fill="currentColor" d="m3.195 2l1.602 17.994L11.989 22l7.212-2.013L20.805 2H3.195zm14.281 4.123l-.534 5.994l.002.033l-.002.074v-.001l-.38 4.192l-.041.372L12 18.037v.001l-.004.003l-4.512-1.258l-.306-3.465h2.213l.157 1.762l2.453.665l-.001.001l2.461-.675l.261-2.869H9.576l-.044-.485l-.101-1.136l-.052-.611h5.538l.202-2.231H6.682l-.044-.485l-.1-1.137l-.053-.61h11.044l-.053.616z"/>`
	htmlThreeAltInnerSVG         = `<path fill="currentColor" d="M11.99 22a1.001 1.001 0 0 1-.268-.037l-6.473-1.805a1 1 0 0 1-.728-.874L3.08 3.09A1 1 0 0 1 4.075 2h15.85a1 1 0 0 1 .996 1.089l-1.443 16.188a.999.999 0 0 1-.728.874l-6.491 1.812a1.001 1.001 0 0 1-.269.037Z" opacity=".5"/><path fill="currentColor" d="M16.777 6.325A1 1 0 0 0 16.04 6H7.96a1 1 0 1 0 0 2h6.987l-.177 2H10a1 1 0 0 0 0 2h4.592l-.264 2.977l-2.328.528l-2.328-.53l-.096-1.064a1 1 0 0 0-1.992.178l.16 1.79a1.001 1.001 0 0 0 .775.887l3.26.74a1.019 1.019 0 0 0 .443 0l3.26-.74a1.001 1.001 0 0 0 .774-.887l.78-8.79a1.001 1.001 0 0 0-.259-.764Z"/>`
	imageVInnerSVG               = `<path fill="currentColor" d="M19 2H5a3.009 3.009 0 0 0-3 3v8.86l3.88-3.88a3.075 3.075 0 0 1 4.24 0l2.871 2.887l.888-.888a3.008 3.008 0 0 1 4.242 0L22 15.86V5a3.009 3.009 0 0 0-3-3z" opacity=".5"/><path fill="currentColor" d="M10.12 9.98a3.075 3.075 0 0 0-4.24 0L2 13.86V19a3.009 3.009 0 0 0 3 3h14a3 3 0 0 0 2.16-.92L10.12 9.98z"/><path fill="currentColor" d="m22 15.858l-3.879-3.879a3.008 3.008 0 0 0-4.242 0l-.888.888l8.165 8.209c.542-.555.845-1.3.844-2.076v-3.142z" opacity=".25"/>`
	instagramInnerSVG            = `<path fill="currentColor" d="M12 6.865A5.135 5.135 0 1 0 17.135 12A5.135 5.135 0 0 0 12 6.865Zm0 8.469A3.334 3.334 0 1 1 15.334 12A3.333 3.333 0 0 1 12 15.334Z"/><path fill="currentColor" d="M21.94 7.877a7.333 7.333 0 0 0-.465-2.427a4.918 4.918 0 0 0-1.153-1.772a4.894 4.894 0 0 0-1.77-1.153a7.323 7.323 0 0 0-2.428-.464C15.058 2.012 14.717 2 12.001 2s-3.057.011-4.123.06a7.333 7.333 0 0 0-2.428.465a4.905 4.905 0 0 0-1.77 1.153A4.886 4.886 0 0 0 2.525 5.45a7.333 7.333 0 0 0-.464 2.427c-.05 1.066-.06 1.407-.06 4.123s.01 3.057.06 4.123a7.334 7.334 0 0 0 .464 2.427a4.888 4.888 0 0 0 1.154 1.772a4.917 4.917 0 0 0 1.771 1.153a7.338 7.338 0 0 0 2.428.464C8.944 21.988 9.285 22 12 22s3.057-.011 4.123-.06a7.333 7.333 0 0 0 2.427-.465a5.113 5.113 0 0 0 2.925-2.925a7.316 7.316 0 0 0 .465-2.427c.048-1.067.06-1.407.06-4.123s-.012-3.057-.06-4.123Zm-1.8 8.164a5.549 5.549 0 0 1-.343 1.857a3.311 3.311 0 0 1-1.898 1.898a5.522 5.522 0 0 1-1.857.344c-1.055.048-1.371.058-4.042.058s-2.986-.01-4.04-.058a5.526 5.526 0 0 1-1.857-.344a3.108 3.108 0 0 1-1.15-.748a3.085 3.085 0 0 1-.748-1.15a5.521 5.521 0 0 1-.344-1.857c-.048-1.054-.058-1.37-.058-4.04s.01-2.987.058-4.042a5.563 5.563 0 0 1 .344-1.857a3.107 3.107 0 0 1 .748-1.15a3.082 3.082 0 0 1 1.15-.748A5.523 5.523 0 0 1 7.96 3.86C9.014 3.81 9.331 3.8 12 3.8s2.987.011 4.042.059a5.564 5.564 0 0 1 1.857.344a3.31 3.31 0 0 1 1.898 1.898a5.523 5.523 0 0 1 .344 1.857c.048 1.055.058 1.37.058 4.041s-.01 2.986-.058 4.041ZM17.339 5.462Z"/>`
	instagramAltInnerSVG         = `<path fill="currentColor" d="M20.936 7.564a6.06 6.06 0 0 0-.377-2.039a3.4 3.4 0 0 0-.821-1.263a3.398 3.398 0 0 0-1.263-.82a6.094 6.094 0 0 0-2.039-.378C15.278 3.012 14.931 3 12 3c-2.93 0-3.279.01-4.436.064a6.06 6.06 0 0 0-2.039.377a3.39 3.39 0 0 0-1.263.821a3.425 3.425 0 0 0-.82 1.264a6.096 6.096 0 0 0-.378 2.038C3.012 8.723 3 9.07 3 12.001s.01 3.279.064 4.436a6.06 6.06 0 0 0 .378 2.039c.175.476.456.908.821 1.261a3.41 3.41 0 0 0 1.263.822c.652.242 1.342.37 2.038.377C8.721 20.99 9.07 21 12 21s3.28-.01 4.436-.064a6.059 6.059 0 0 0 2.039-.377a3.635 3.635 0 0 0 2.084-2.083a6.09 6.09 0 0 0 .377-2.04C20.99 15.28 21 14.932 21 12s-.01-3.278-.064-4.436z" opacity=".5"/><path fill="currentColor" d="M17.793 7.618a2.459 2.459 0 0 0-1.41-1.41a4.132 4.132 0 0 0-1.38-.256c-.784-.036-1.02-.044-3.003-.044s-2.218.008-3 .044a4.1 4.1 0 0 0-1.38.255c-.649.25-1.16.762-1.41 1.41c-.164.443-.25.91-.256 1.38c-.036.784-.043 1.019-.043 3.003s.007 2.22.043 3.003c.005.471.092.938.255 1.38a2.46 2.46 0 0 0 1.41 1.41c.443.163.91.25 1.38.255c.783.036 1.018.043 3.002.043s2.22-.007 3.003-.043a4.104 4.104 0 0 0 1.38-.256a2.46 2.46 0 0 0 1.41-1.41c.164-.441.25-.908.256-1.38c.036-.783.043-1.018.043-3.002s-.007-2.219-.043-3.002a4.103 4.103 0 0 0-.256-1.38zM12 15.815a3.815 3.815 0 1 1 0-7.63a3.815 3.815 0 0 1 0 7.63zm3.966-6.89a.892.892 0 0 1 0-1.783h.001a.892.892 0 0 1 .005 1.783h-.006z"/><path fill="currentColor" d="M12 9.523A2.477 2.477 0 1 0 14.477 12a2.477 2.477 0 0 0-2.476-2.477z"/><path fill="currentColor" d="M21.93 7.07a6.734 6.734 0 0 0-.42-2.264a3.779 3.779 0 0 0-.913-1.403a3.776 3.776 0 0 0-1.403-.913a6.771 6.771 0 0 0-2.265-.42C15.642 2.014 15.257 2 12 2c-3.255 0-3.642.012-4.928.07a6.76 6.76 0 0 0-2.265.42a3.76 3.76 0 0 0-1.403.913a3.806 3.806 0 0 0-.913 1.403a6.773 6.773 0 0 0-.42 2.265C2.014 8.358 2 8.744 2 12.001s.012 3.643.071 4.929c.01.773.151 1.54.42 2.265c.195.53.507 1.01.912 1.402a3.79 3.79 0 0 0 1.403.913a6.732 6.732 0 0 0 2.265.42c1.286.059 1.673.07 4.929.07c3.257 0 3.643-.012 4.93-.07a6.732 6.732 0 0 0 2.264-.42a4.04 4.04 0 0 0 2.316-2.315c.268-.726.41-1.492.42-2.265c.058-1.287.07-1.674.07-4.93c0-3.257-.012-3.642-.07-4.93zm-2.545 7.993a5.435 5.435 0 0 1-.345 1.804a3.8 3.8 0 0 1-2.173 2.173a5.45 5.45 0 0 1-1.803.345c-.793.037-1.046.045-3.064.045s-2.27-.009-3.063-.045a5.453 5.453 0 0 1-1.803-.345a3.799 3.799 0 0 1-2.174-2.173a5.45 5.45 0 0 1-.345-1.804C4.58 14.271 4.57 14.018 4.57 12s.009-2.271.045-3.063a5.45 5.45 0 0 1 .345-1.804A3.798 3.798 0 0 1 7.134 4.96a5.45 5.45 0 0 1 1.803-.345c.793-.037 1.047-.045 3.064-.045s2.27.009 3.063.045a5.442 5.442 0 0 1 1.804.345a3.8 3.8 0 0 1 2.172 2.173c.217.578.333 1.187.345 1.804c.036.792.045 1.045.045 3.063s-.009 2.27-.045 3.063z"/>`
	intercomInnerSVG             = `<path fill="currentColor" d="M19.333 12.996a.667.667 0 0 1-1.333 0V7a.667.667 0 0 1 1.333 0v5.996zm-.232 4.173A11.328 11.328 0 0 1 12 19.329a11.328 11.328 0 0 1-7.1-2.16a.667.667 0 1 1 .866-1.014A10.23 10.23 0 0 0 12 17.995a10.096 10.096 0 0 0 6.233-1.839a.667.667 0 1 1 .868 1.013zM4.667 7A.667.667 0 0 1 6 7v5.996a.667.667 0 0 1-1.333 0V7zM8 5.667a.667.667 0 0 1 1.333 0v8.906a.667.667 0 0 1-1.333 0V5.667zm3.333-.338a.667.667 0 0 1 1.333 0v9.667a.667.667 0 0 1-1.333 0V5.329zm3.334.338a.667.667 0 0 1 1.333 0v8.906a.667.667 0 0 1-1.333 0V5.667z" opacity=".5"/><path fill="currentColor" d="M19.5 2h-15A2.5 2.5 0 0 0 2 4.5v15A2.5 2.5 0 0 0 4.5 22h15a2.5 2.5 0 0 0 2.5-2.5v-15A2.5 2.5 0 0 0 19.5 2zm-4.833 3.667a.667.667 0 0 1 1.333 0v8.906a.667.667 0 0 1-1.333 0V5.667zm-3.334-.338a.667.667 0 0 1 1.333 0v9.667a.667.667 0 0 1-1.333 0V5.329zM8 5.667a.667.667 0 0 1 1.333 0v8.906a.667.667 0 0 1-1.333 0V5.667zM4.667 7A.667.667 0 0 1 6 7v5.996a.667.667 0 0 1-1.333 0V7zM19.1 17.169a11.328 11.328 0 0 1-7.1 2.16a11.327 11.327 0 0 1-7.1-2.16a.667.667 0 1 1 .866-1.014A10.23 10.23 0 0 0 12 17.995a10.102 10.102 0 0 0 6.233-1.838a.667.667 0 1 1 .868 1.012zm.232-4.173a.667.667 0 0 1-1.333 0V7a.667.667 0 0 1 1.333 0v5.996z"/>`
	intercomAltInnerSVG          = `<path fill="currentColor" d="M20 23H4a3.003 3.003 0 0 1-3-3V4a3.003 3.003 0 0 1 3-3h16a3.003 3.003 0 0 1 3 3v16a3.003 3.003 0 0 1-3 3Z" opacity=".5"/><path fill="currentColor" d="M12 19a10.8 10.8 0 0 1-6.644-2.06a1 1 0 0 1 1.288-1.532A8.987 8.987 0 0 0 12 17a8.995 8.995 0 0 0 5.361-1.595a1 1 0 0 1 1.282 1.535A10.795 10.795 0 0 1 12 19zm-6-6a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1zm4 2a1 1 0 0 1-1-1V6a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1zm4 0a1 1 0 0 1-1-1V6a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1zm4-2a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1z"/>`
	javaScriptInnerSVG           = `<path fill="currentColor" d="M11.181 2.213a1.677 1.677 0 0 1 1.637 0l7.479 4.225a1.54 1.54 0 0 1 .778 1.325v8.475a1.556 1.556 0 0 1-.836 1.354l-7.452 4.204a1.63 1.63 0 0 1-1.655-.046l-2.236-1.292a1.316 1.316 0 0 1-.432-.311c.095-.128.265-.144.403-.2a5.22 5.22 0 0 0 .883-.412a.206.206 0 0 1 .23.014c.636.365 1.267.741 1.907 1.103c.136.079.274-.026.39-.09q3.658-2.068 7.317-4.13a.242.242 0 0 0 .133-.238q.003-4.193.001-8.387a.262.262 0 0 0-.156-.261q-3.714-2.092-7.426-4.186a.258.258 0 0 0-.292 0Q8.141 5.452 4.43 7.547a.259.259 0 0 0-.157.26v8.387a.237.237 0 0 0 .135.235q.99.562 1.983 1.12a1.532 1.532 0 0 0 1.24.166a.94.94 0 0 0 .609-.883c.003-2.78-.002-5.56.002-8.338a.21.21 0 0 1 .228-.214c.318-.002.635-.004.953.001a.22.22 0 0 1 .207.254c-.001 2.797.003 5.594-.002 8.39a2.127 2.127 0 0 1-.995 1.922a3.064 3.064 0 0 1-2.738-.075c-.727-.363-1.42-.79-2.133-1.18a1.553 1.553 0 0 1-.835-1.354V7.763a1.542 1.542 0 0 1 .804-1.342q3.726-2.103 7.451-4.208Z"/><path fill="currentColor" d="M13.348 8.083a5.88 5.88 0 0 1 3.218.493a2.434 2.434 0 0 1 1.187 2.106a.228.228 0 0 1-.247.168c-.315 0-.629.004-.943-.002a.243.243 0 0 1-.228-.236a1.436 1.436 0 0 0-.687-.992a4.071 4.071 0 0 0-1.884-.27a2.558 2.558 0 0 0-1.349.336a.855.855 0 0 0-.284.955c.101.24.378.317.605.388c1.306.342 2.69.308 3.97.757a1.87 1.87 0 0 1 1.23 1.095a2.338 2.338 0 0 1-.396 2.23a3.126 3.126 0 0 1-1.678.905a8.063 8.063 0 0 1-2.533.108a3.992 3.992 0 0 1-2.228-.839a2.338 2.338 0 0 1-.753-1.844a.211.211 0 0 1 .231-.186c.317-.003.633-.004.95 0a.222.222 0 0 1 .226.22a1.439 1.439 0 0 0 .536 1.01a4.06 4.06 0 0 0 2.187.398a2.97 2.97 0 0 0 1.794-.439a.99.99 0 0 0 .27-.946c-.078-.28-.37-.41-.621-.495c-1.29-.408-2.688-.26-3.965-.72a1.93 1.93 0 0 1-1.218-1.063a2.175 2.175 0 0 1 .434-2.262a3.506 3.506 0 0 1 2.176-.875Z"/>`
	keySkeletonInnerSVG          = `<path fill="currentColor" d="M7 12a5 5 0 1 0 0 10a5 5 0 0 0 0-10zm-1.415 7.413a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="m21.001 4.413l.706-.706a1 1 0 1 0-1.414-1.414L9.754 12.832a5.02 5.02 0 0 1 1.414 1.414l5.591-5.591l2.12 2.12a1 1 0 0 0 1.414-1.414l-2.12-2.12l1.414-1.414l.706.705a1 1 0 0 0 1.414-1.414l-.706-.705z"/>`
	keySkeletonAltInnerSVG       = `<path fill="currentColor" d="M7 12a5 5 0 1 0 5 5a5 5 0 0 0-5-5zm-1.415 7.412a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M21.708 6.534L20.293 5.12l1.413-1.413a1 1 0 1 0-1.414-1.414L9.753 12.831a5.018 5.018 0 0 1 1.415 1.414l4.883-4.883l1.414 1.414a1 1 0 0 0 1.414-1.414l-1.414-1.414l1.414-1.414l1.415 1.414a1 1 0 0 0 1.414-1.414z"/>`
	keyholeCircleInnerSVG        = `<path fill="currentColor" d="M14 10a2 2 0 1 0-3 1.723V15a1 1 0 0 0 2 0v-3.277A1.991 1.991 0 0 0 14 10Z"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm1 9.723V15a1 1 0 0 1-2 0v-3.277a2 2 0 1 1 2 0Z" opacity=".5"/>`
	keyholeSquareInnerSVG        = `<path fill="currentColor" d="M14 10a2 2 0 1 0-3 1.723V15a1 1 0 0 0 2 0v-3.277A1.991 1.991 0 0 0 14 10Z"/><path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V5a3.003 3.003 0 0 0-3-3Zm-6 9.723V15a1 1 0 0 1-2 0v-3.277a2 2 0 1 1 2 0Z" opacity=".5"/>`
	keyholeSquareFullInnerSVG    = `<path fill="currentColor" d="M14 10a2 2 0 1 0-3 1.723V15a1 1 0 0 0 2 0v-3.277A1.991 1.991 0 0 0 14 10Z"/><path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm-8 9.723V15a1 1 0 0 1-2 0v-3.277a2 2 0 1 1 2 0Z" opacity=".5"/>`
	layerGroupInnerSVG           = `<path fill="currentColor" d="M12 14.195c-.176 0-.348-.046-.5-.133l-9-5.198a1 1 0 0 1 0-1.732l9-5.194c.31-.177.69-.177 1 0l9 5.194a1 1 0 0 1 0 1.732l-9 5.198a1.002 1.002 0 0 1-.5.133z" opacity=".25"/><path fill="currentColor" d="m21.5 11.132l-1.964-1.134l-7.036 4.064c-.31.178-.69.178-1 0L4.464 9.998L2.5 11.132a1 1 0 0 0 0 1.732l9 5.198c.31.178.69.178 1 0l9-5.198a1 1 0 0 0 0-1.732z" opacity=".5"/><path fill="currentColor" d="m21.5 15.132l-1.964-1.134l-7.036 4.064c-.31.178-.69.178-1 0l-7.036-4.064L2.5 15.132a1 1 0 0 0 0 1.732l9 5.198c.31.178.69.178 1 0l9-5.198a1 1 0 0 0 0-1.732z"/>`
	layersAltInnerSVG            = `<path fill="currentColor" d="M21 2H9a1 1 0 0 0-1 .999V7h8a1 1 0 0 1 1 .999V16h4a1 1 0 0 0 1-.999V3a1 1 0 0 0-.999-1H21z" opacity=".25"/><path fill="currentColor" d="M3 12h8a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M16 7H6a1 1 0 0 0-1 .999V12h6a1 1 0 0 1 1 .999V19h4a1 1 0 0 0 1-.999V8a1 1 0 0 0-.999-1H16z" opacity=".5"/>`
	leftIndentInnerSVG           = `<path fill="currentColor" d="M21 7H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8 4H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm8 8H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8-4H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M21 14.666a.999.999 0 0 1-.64-.231l-2-1.667a1 1 0 0 1 0-1.536l2-1.667a1 1 0 0 1 1.64.769v3.333a1 1 0 0 1-1 1Z"/>`
	leftIndentAltInnerSVG        = `<path fill="currentColor" d="M21 19h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2zm0-4h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2zm0-8h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2zm0 4h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M9 19a1 1 0 0 1-1-1V6a1 1 0 0 1 2 0v12a1 1 0 0 1-1 1zm-4-4.333a.999.999 0 0 1-.64-.231l-2-1.667a1 1 0 0 1 0-1.538l2-1.667a1 1 0 0 1 1.28 1.538L4.562 12l1.078.898A1 1 0 0 1 5 14.667z"/>`
	lineInnerSVG                 = `<path fill="currentColor" d="M18.223 14.278c-.041.06-.087.116-.136.17l-.004.004a6.36 6.36 0 0 1-.798.796c-2.03 1.876-5.37 4.109-5.81 3.764c-.383-.299.63-1.764-.54-2.007a9.241 9.241 0 0 1-.243-.031h-.002c-3.437-.489-6.047-2.893-6.047-5.785c0-3.245 3.285-5.876 7.338-5.876s7.337 2.63 7.337 5.876a5.021 5.021 0 0 1-1.095 3.089z" opacity=".5"/><path fill="currentColor" d="M17.907 2H6.093A4.093 4.093 0 0 0 2 6.093v11.814A4.093 4.093 0 0 0 6.093 22h11.814A4.093 4.093 0 0 0 22 17.907V6.093A4.093 4.093 0 0 0 17.907 2zm.316 12.278a1.53 1.53 0 0 1-.136.17l-.004.004a6.364 6.364 0 0 1-.798.796c-2.03 1.876-5.37 4.109-5.81 3.764c-.383-.299.631-1.764-.54-2.007a9.284 9.284 0 0 1-.243-.031h-.002c-3.437-.489-6.047-2.893-6.047-5.785c0-3.245 3.285-5.876 7.338-5.876s7.337 2.63 7.337 5.876a5.022 5.022 0 0 1-1.095 3.089z"/><path fill="currentColor" d="M9.203 12.265H8.138V9.842a.384.384 0 0 0-.383-.383h-.032a.384.384 0 0 0-.382.383v2.838c0 .211.171.382.382.383h1.48a.384.384 0 0 0 .383-.383v-.032a.384.384 0 0 0-.383-.383zm4.572-2.806h-.032a.384.384 0 0 0-.382.383v1.666l-1.383-1.855a.383.383 0 0 0-.333-.194h-.032a.384.384 0 0 0-.382.383v2.838c0 .211.171.382.382.383h.032a.384.384 0 0 0 .383-.383v-1.697l1.393 1.903a.401.401 0 0 0 .028.038a.32.32 0 0 0 .15.11a.374.374 0 0 0 .144.029h.032a.379.379 0 0 0 .18-.045a.27.27 0 0 0 .092-.07a.382.382 0 0 0 .111-.268V9.842a.384.384 0 0 0-.383-.383zm-3.498 0h.032c.212 0 .383.171.383.383v2.838a.383.383 0 0 1-.383.383h-.032a.383.383 0 0 1-.382-.383V9.842c0-.212.17-.383.382-.383zm6.234.007h-1.48a.384.384 0 0 0-.383.383v2.838c0 .21.172.382.383.382h1.48a.384.384 0 0 0 .382-.382v-.032a.384.384 0 0 0-.382-.383h-1.066v-.606h1.066a.384.384 0 0 0 .382-.382v-.032a.384.384 0 0 0-.382-.383h-1.066v-.606h1.066a.384.384 0 0 0 .382-.383v-.03a.384.384 0 0 0-.382-.383z"/>`
	lineSpacingInnerSVG          = `<path fill="currentColor" d="M21 8H10a1 1 0 0 1 0-2h11a1 1 0 0 1 0 2zm0 5H10a1 1 0 0 1 0-2h11a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M4.667 18.5a1 1 0 0 1-1-1v-11a1 1 0 0 1 2 0v11a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M21 18H10a1 1 0 0 1 0-2h11a1 1 0 0 1 0 2Z" opacity=".5"/><path fill="currentColor" d="M6.334 9a.997.997 0 0 1-.77-.36l-.897-1.078l-.898 1.079A1 1 0 0 1 2.23 7.359l1.667-2A1.002 1.002 0 0 1 4.667 5a1 1 0 0 1 .769.36l1.666 2A1 1 0 0 1 6.334 9zM4.667 19a1.002 1.002 0 0 1-.769-.36l-1.667-2a1 1 0 1 1 1.538-1.28l.898 1.078l.897-1.078a1 1 0 0 1 1.538 1.28l-1.666 2a1 1 0 0 1-.769.36z"/>`
	linkHInnerSVG                = `<path fill="currentColor" d="M10 17H7A5 5 0 0 1 7 7h3a1 1 0 0 1 0 2H7a3 3 0 0 0 0 6h3a1 1 0 0 1 0 2zm7 0h-3a1 1 0 0 1 0-2h3a3 3 0 0 0 0-6h-3a1 1 0 0 1 0-2h3a5 5 0 0 1 0 10z"/><path fill="currentColor" d="M15 13H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Z" opacity=".5"/>`
	linkedinInnerSVG             = `<path fill="currentColor" d="M5.086 9.711h3.002v9.031H5.086zm1.501-1.233h-.02a1.565 1.565 0 1 1 .04-3.12a1.565 1.565 0 1 1-.02 3.12zm12.325 10.264H15.91v-4.83c0-1.215-.434-2.043-1.52-2.043a1.643 1.643 0 0 0-1.54 1.098a2.043 2.043 0 0 0-.1.732v5.043h-3c0-.003.04-8.184 0-9.03h3.002v1.28a2.978 2.978 0 0 1 2.705-1.493c1.975 0 3.456 1.291 3.456 4.065v5.178z" opacity=".5"/><path fill="currentColor" d="M20.468 2H3.532a1.452 1.452 0 0 0-1.47 1.433v17.135c.011.8.669 1.442 1.47 1.432h16.936a1.451 1.451 0 0 0 1.47-1.432V3.433A1.451 1.451 0 0 0 20.467 2zM8.088 18.742H5.086V9.711h3.002v9.031zM6.833 8.48a1.57 1.57 0 0 1-.246-.002h-.02a1.565 1.565 0 1 1 .04-3.12a1.565 1.565 0 0 1 .226 3.122zm12.079 10.262H15.91v-4.83c0-1.215-.434-2.043-1.52-2.043a1.643 1.643 0 0 0-1.54 1.098a2.06 2.06 0 0 0-.1.732v5.043h-3c0-.003.04-8.184 0-9.03h3.002v1.28a2.978 2.978 0 0 1 2.705-1.493c1.975 0 3.456 1.291 3.456 4.065v5.178z"/>`
	linkedinAltInnerSVG          = `<path fill="currentColor" d="M5.004 7h-.029a2.235 2.235 0 1 1 .057-4.457A2.235 2.235 0 1 1 5.004 7Zm-1.986 3h4v12h-4zm14.5 0a4.473 4.473 0 0 0-3.5 1.703V10h-4v12h4v-5.5a2 2 0 0 1 4 0V22h4v-7.5a4.5 4.5 0 0 0-4.5-4.5Z"/>`
	linuxInnerSVG                = `<path fill="currentColor" d="M12.6 6.5c.1.1.2.1.4.2c.1 0 .3.1.5.2h.1c.2-.3.3-.6.3-.9c.1-.6-.4-1.2-1-1.2c-.6.1-1 .7-1 1.3v.1c.2 0 .5.1.7.3zm4.5 8c-.3-.2-.6-.2-1-.3c-.2-.9-.5-1.7-.9-2.5c-.6-1.2-1.1-2.4-1.4-3.7c-.2.3-.5.5-.9.6c-.1.1-.3.1-.4.2c-.3.2-.7.4-1.1.4h-.1c-.4 0-.7-.2-.9-.4c-.1-.1-.2-.2-.3-.2c-.1 0-.1-.1-.2-.1c-.1 1.5-1 3.2-1.5 4.3c-.3.8-.5 1.6-.5 2.4c-.8-1.2-.2-2.7.1-3.3c.4-.8.4-.9.3-.8c-.7 1.1-2.1 3.5-.1 4.8c2.1 1.3 2.2 2.6 1 2.5c.1.2.3.4.4.5c1.2 1 2.9 1.1 4.2.3c0-.1.1-.3.1-.4c.1-.3.2-.7.2-1c.1-1.1.1-2.2.9-2.6c.5-.2 1.2-.1 1.6.2c.1.1.2.1.3.2c.2.1.4.1.5.1h.6c.3-.4-.1-.8-.9-1.2zm-6.9-7.7l.1-.1c.2-.3.5-.4.8-.5V6c0-.6-.4-1.1-.8-1.1c-.3.1-.6.6-.6 1.2c0 .3.1.6.3.8c.1 0 .2-.1.2-.1z" opacity=".25"/><path fill="currentColor" d="M8.5 17.4s0 .1 0 0c-.1-.1-.1-.3-.2-.4c.1.2.1.3.2.4z" opacity=".25"/><path fill="currentColor" d="M15.4 22h-.2c-.6-.1-1.1-.4-1.3-1c-.2-.7-.2-1.5.2-2.2c.1-.3.2-.7.2-1c.1-1.1.1-2.2.9-2.6c.5-.2 1.2-.1 1.6.2c.1.1.2.1.3.2c.2.1.4.1.5.1c.4-.1.8 0 1.1.3c.3.3.4.6.5 1c0 .2.1.5.2.6c.5.5.7 1 .6 1.3c-.1.5-.6.7-1.1 1c-.5.2-1 .5-1.4.9c-.5.7-1.3 1.1-2.1 1.2z" opacity=".5"/><path fill="currentColor" d="M17.9 15.9c-.4.5-1 .8-1.6.8c-.6-.1-.8-.9-.7-1.5c.1-.7.7-.7 1.5-.4c.8.3 1.1.7.8 1.1zm-5.3-9.3c.1.1.3.1.4.2c.2-.2.2-.4.2-.6c0-.4-.2-.7-.4-.7s-.5.3-.5.7v.3c.1.1.2.1.3.1zm-2.2.2l.3-.3v-.3c0-.3-.2-.6-.4-.5c-.2 0-.3.3-.3.6c.1.3.3.5.4.5z"/><path fill="currentColor" d="M17.3 10.8c-.8-1.3-2-2.1-2-3.7c0-1.9.2-5.4-3.3-5.1c-3.5.3-2.5 4-2.6 5.3c0 1.1-.5 2.2-1.3 3.1c-.1.1-.2.3-.3.4c-.9 1.3-1.9 3-1.8 4.6c.2-.1.4-.1.5-.1c.8.1 1.2.9 1.7 1.8c.1.1.1.3.2.4l.6.9l.1.1c1.2.1 1.2-1.2-1-2.5c-2-1.3-.6-3.7.1-4.8c.1-.1.1 0-.3.8c-.3.6-.9 2.1-.1 3.2c0-.8.2-1.6.5-2.4c.5-1 1.4-2.8 1.5-4.3c-.1-.2-.3-.5-.3-.8c0-.2.1-.4.3-.6c.1 0 .2-.1.3-.2c-.2-.2-.4-.5-.4-.8c0-.6.3-1.1.7-1.1c.4 0 .7.4.8 1.1v.2c.2-.1.5-.1.7 0v-.1c-.1-.6.3-1.2 1-1.3c.6.1 1.1.6 1 1.2c0 .3-.1.6-.3.9c.3.1.5.4.5.7c0 .2-.1.3-.2.5c.4 1.3.8 2.5 1.4 3.7c.4.8.7 1.6.9 2.5c.3 0 .7.1 1 .3c.3.2.6.3.8.5c.1.1.1.2.2.2c0 0 0 .1.1.1v.3c.1 0 .3.1.4.2c.5-2-.5-3.9-1.4-5.2z"/><path fill="currentColor" d="M11.4 8.5c-.5 0-1-.3-1.4-.7l-.1-.1c-.1 0-.1 0-.1-.1s.1-.1.1-.1c.1 0 .1.1.3.2c.3.4.7.6 1.2.6c.5-.1 1-.2 1.4-.5c.2-.1.4-.2.7-.3c.1 0 .1 0 .1.1s0 .1-.1.1c-.2.1-.4.1-.6.3c-.5.3-1 .5-1.5.5z" opacity=".7"/><path fill="currentColor" d="M13.5 6.9c-.2-.1-.3-.1-.5-.2c-.1 0-.3-.1-.4-.2c-.6-.6-1.7-.5-2.3.2l-.1.1s-.2.2-.4.3c-.2.2-.3.4-.3.6c.1.4.3.7.7.9c.1.1.2.2.3.2c.2.3.6.4.9.4h.1c.4 0 .8-.1 1.1-.4c.1-.1.2-.2.4-.2c.5-.1 1-.5 1.1-1.1c0-.2-.3-.5-.6-.6zm-.1.8c-.2.1-.4.1-.6.3c-.4.3-.9.5-1.5.6c-.5 0-1-.3-1.4-.7l-.1-.1c-.1 0-.1 0-.1-.1s.1-.1.1-.1c.1 0 .1.1.3.2c.3.4.7.6 1.2.6c.5-.1 1-.2 1.4-.5c.2-.1.4-.2.7-.3c.1-.1.1-.1 0 .1c.1-.1.1 0 0 0zm-5 14.1c-.6 0-1.2-.2-1.7-.5c-.5-.2-1-.4-1.6-.4c-.7-.1-1.2-.1-1.5-.6c-.2-.4-.2-1 0-1.4v-.7c-.1-.4-.1-.7.1-1.1c.2-.3.5-.6.9-.7c.2 0 .4-.1.5-.3c.1-.1.2-.2.2-.3c.2-.5.8-.8 1.3-.8c.9.1 1.4 1.1 1.9 2.2l.6.9c.5.7 1.1 1.3 1 2c0 .5-.3 1-.8 1.2c-.2.4-.6.5-.9.5z" opacity=".5"/><path fill="currentColor" d="M13.8 19.4c-1.3.7-3 .6-4.2-.4c.1.2.2.3.3.5v.1c.1.3.2.5.2.8c0 .2-.1.3-.1.5c.5 0 1.1-.2 2-.3c.5 0 1.1.1 1.8.2c-.1-.5-.1-1 0-1.4z"/>`
	listUiAltInnerSVG            = `<path fill="currentColor" d="M21.5 8h-14a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2zm0 5h-10a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm0 5h-6a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2z"/><circle cx="3.5" cy="7" r="1" fill="currentColor"/><circle cx="7.5" cy="12" r="1" fill="currentColor"/><circle cx="11.5" cy="17" r="1" fill="currentColor"/>`
	listUlInnerSVG               = `<path fill="currentColor" d="M21 8H7a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2zm0 5H7a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2zm0 5H7a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2zM3 8a1 1 0 0 1-.38-.08a1.151 1.151 0 0 1-.33-.21a1.162 1.162 0 0 1-.21-.33a.946.946 0 0 1 0-.76a1.149 1.149 0 0 1 .21-.33a.998.998 0 0 1 1.09-.21a1.034 1.034 0 0 1 .33.21a1.158 1.158 0 0 1 .21.33a.941.941 0 0 1 0 .76a1.171 1.171 0 0 1-.21.33A.992.992 0 0 1 3 8zm0 5a1 1 0 0 1-.38-.08a1.151 1.151 0 0 1-.33-.21a1.162 1.162 0 0 1-.21-.33a.946.946 0 0 1 0-.76a1.03 1.03 0 0 1 .21-.33a1.151 1.151 0 0 1 .33-.21a.999.999 0 0 1 1.09.21a1.037 1.037 0 0 1 .21.33a.941.941 0 0 1 0 .76a1.171 1.171 0 0 1-.21.33a1.154 1.154 0 0 1-.33.21A1 1 0 0 1 3 13zm0 5a1 1 0 0 1-.38-.08a1.151 1.151 0 0 1-.33-.21a.991.991 0 0 1-.21-1.09a1.03 1.03 0 0 1 .21-.33a1.027 1.027 0 0 1 .33-.21a.995.995 0 0 1 .76 0a1.034 1.034 0 0 1 .33.21a1.037 1.037 0 0 1 .21.33a.99.99 0 0 1-.21 1.09a1.154 1.154 0 0 1-.33.21A1 1 0 0 1 3 18z"/>`
	lockInnerSVG                 = `<path fill="currentColor" d="M16 11H8a1 1 0 0 1-1-1V7a5 5 0 0 1 10 0v3a1 1 0 0 1-1 1ZM9 9h6V7a3 3 0 0 0-6 0Z" opacity=".5"/><rect width="16" height="13" x="4" y="9" fill="currentColor" rx="3"/>`
	lockAccessInnerSVG           = `<path fill="currentColor" d="M21 10c-.6 0-1-.4-1-1V4h-5c-.6 0-1-.4-1-1s.4-1 1-1h6c.6 0 1 .4 1 1v6c0 .6-.4 1-1 1zM3 10c-.6 0-1-.4-1-1V3c0-.6.4-1 1-1h6c.6 0 1 .4 1 1s-.4 1-1 1H4v5c0 .6-.4 1-1 1zm6 12H3c-.6 0-1-.4-1-1v-6c0-.6.4-1 1-1s1 .4 1 1v5h5c.6 0 1 .4 1 1s-.4 1-1 1zm12 0h-6c-.6 0-1-.4-1-1s.4-1 1-1h5v-5c0-.6.4-1 1-1s1 .4 1 1v6c0 .6-.4 1-1 1z"/><path fill="currentColor" d="M9 10h6c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2H9c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2zm2-1c0-.6.4-1 1-1s1 .4 1 1v1h2V9c0-1.7-1.3-3-3-3S9 7.3 9 9v1h2V9z" opacity=".5"/>`
	lockAltInnerSVG              = `<path fill="currentColor" d="M9 7a3 3 0 1 1 6 0v2h2V7A5 5 0 0 0 7 7v2h2V7zm3 11a1 1 0 0 1-1-1v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-1 1z" opacity=".5"/><path fill="currentColor" d="M17 9H7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm-4 8a1 1 0 0 1-2 0v-3a1 1 0 1 1 2 0v3z"/>`
	lockOpenAltInnerSVG          = `<path fill="currentColor" d="M8 11a1 1 0 0 1-1-.999V7a5.002 5.002 0 0 1 8.532-3.542a5.078 5.078 0 0 1 1.307 2.293a1 1 0 1 1-1.937.501v-.003a3.057 3.057 0 0 0-.786-1.379A3.002 3.002 0 0 0 9 7v3a1 1 0 0 1-.999 1H8zm5.5 3.5a1.5 1.5 0 1 0-3 0c0 .443.195.836.5 1.11v1.392A1 1 0 0 0 12 18h.001A1 1 0 0 0 13 17v-1.39c.305-.274.5-.667.5-1.11z" opacity=".5"/><path fill="currentColor" d="M17 9H7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm-4 6.61V17a1 1 0 0 1-.999 1H12a1 1 0 0 1-1-.999V15.61a1.49 1.49 0 0 1-.5-1.11a1.5 1.5 0 1 1 3 0c0 .443-.195.836-.5 1.11z"/>`
	lottiefilesInnerSVG          = `<path fill="currentColor" d="M7 18a1 1 0 0 1 0-2c1.66 0 2.856-2.177 4.124-4.482C12.616 8.805 14.159 6 17 6a1 1 0 0 1 0 2c-1.66 0-2.856 2.177-4.124 4.482C11.384 15.195 9.841 18 7 18Z" opacity=".5"/><path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V5a3.003 3.003 0 0 0-3-3Zm-2 6c-1.66 0-2.856 2.177-4.124 4.482C11.384 15.195 9.841 18 7 18a1 1 0 0 1 0-2c1.66 0 2.856-2.177 4.124-4.482C12.616 8.805 14.159 6 17 6a1 1 0 0 1 0 2Z"/>`
	masterCardInnerSVG           = `<path fill="currentColor" d="M15.273 18.728A6.728 6.728 0 1 1 22 11.999V12a6.735 6.735 0 0 1-6.727 6.728z" opacity=".5"/><path fill="currentColor" d="M8.727 18.728A6.728 6.728 0 1 1 15.455 12a6.735 6.735 0 0 1-6.728 6.728z"/>`
	mediumMInnerSVG              = `<path fill="currentColor" d="M22 6.417h-.791a.898.898 0 0 0-.709.694v9.83a.84.84 0 0 0 .709.642H22v2.334h-7.167v-2.334h1.5V7.25h-.073l-3.503 12.667h-2.712L6.588 7.25H6.5v10.333H8v2.334H2v-2.334h.768a.841.841 0 0 0 .732-.641v-9.83a.896.896 0 0 0-.732-.695H2V4.083h7.503l2.463 9.167h.068l2.486-9.167H22v2.334"/>`
	microscopeInnerSVG           = `<path fill="currentColor" d="M11.441 13.604a1 1 0 0 1-.707-.293L7.198 9.772a1 1 0 0 1 0-1.414l6.363-6.363a.997.997 0 0 1 .39-.242l2.121-.707a1 1 0 0 1 1.024.242l2.122 2.121a1 1 0 0 1 .241 1.024l-.708 2.122a.991.991 0 0 1-.241.39l-6.362 6.366a1 1 0 0 1-.707.293Zm6.362-7.366Z" opacity=".25"/><path fill="currentColor" d="m7.198 9.772l-1.416 1.415a1 1 0 0 0 0 1.415l2.122 2.12a1 1 0 0 0 1.414 0l1.414-1.413Z" opacity=".5"/><path fill="currentColor" d="M8 18.005H4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M20 23.005H4a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2Z" opacity=".5"/><path fill="currentColor" d="M14.816 21.005a2.965 2.965 0 0 0 .184-1a3 3 0 0 0-6 0a2.965 2.965 0 0 0 .184 1Z" opacity=".25"/><path fill="currentColor" d="m17.873 7.583l-1.415 1.415A5.955 5.955 0 0 1 18 13.005a6.048 6.048 0 0 1-3.455 5.431a2.971 2.971 0 0 1 .455 1.57a2.645 2.645 0 0 1-.04.407A8.044 8.044 0 0 0 20 13.005a7.945 7.945 0 0 0-2.127-5.422zM9.42 18.499a7.036 7.036 0 0 1-1.095-.56a.983.983 0 0 1-.326.066H5.326a8.873 8.873 0 0 0 3.72 2.472A2.69 2.69 0 0 1 9 20.005a2.966 2.966 0 0 1 .42-1.506z" opacity=".5"/>`
	microsoftInnerSVG            = `<path fill="currentColor" d="M2 2h9.503v9.503H2zm10.493 0h9.503v9.503h-9.503zM2 12.497h9.503V22H2zm10.493 0h9.503V22h-9.503z"/>`
	minusSquareFullInnerSVG      = `<path fill="currentColor" d="M17 13H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm-4 11H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z" opacity=".5"/>`
	multiplyInnerSVG             = `<path fill="currentColor" d="M7 18a1 1 0 0 1-.707-1.707l10-10a1 1 0 0 1 1.414 1.414l-10 10A.997.997 0 0 1 7 18Z"/><path fill="currentColor" d="M17 18a.997.997 0 0 1-.707-.293l-10-10a1 1 0 0 1 1.414-1.414l10 10A1 1 0 0 1 17 18Z"/>`
	objectGroupInnerSVG          = `<path fill="currentColor" d="M11 10h5a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-5a1 1 0 0 1-1-1v-5a1 1 0 0 1 1-1z" opacity=".5"/><path fill="currentColor" d="M10 11a1 1 0 0 1 1-1h3V8a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h2v-3zM4 22a2 2 0 1 1 0-4a2 2 0 0 1 0 4zM4 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm16 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/><path fill="currentColor" d="M18.278 5a1.936 1.936 0 0 1 0-2H5.722a1.936 1.936 0 0 1 0 2h12.556zM20 18a1.98 1.98 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v12.556A1.98 1.98 0 0 1 20 18zM4 18a1.98 1.98 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v12.556A1.98 1.98 0 0 1 4 18zm14.278 1H5.722a1.936 1.936 0 0 1 0 2h12.556a1.936 1.936 0 0 1 0-2z" opacity=".25"/>`
	objectUngroupInnerSVG        = `<path fill="currentColor" d="M4 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM4 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM14 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM14 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002z" opacity=".5"/><path fill="currentColor" d="M10 22a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM10 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM20 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002zM20 22a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-2.002z"/><path fill="currentColor" d="M12.278 5a1.936 1.936 0 0 1 0-2H5.722a1.936 1.936 0 0 1 0 2h6.556zM4 12a1.98 1.98 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v6.556A1.98 1.98 0 0 1 4 12z" opacity=".25"/><path fill="currentColor" d="M20 18a1.98 1.98 0 0 1 1 .278v-6.556a1.936 1.936 0 0 1-2 0v6.556A1.98 1.98 0 0 1 20 18zm-10 0a1.98 1.98 0 0 1 1 .278v-6.556a1.936 1.936 0 0 1-2 0v6.556A1.98 1.98 0 0 1 10 18z" opacity=".5"/><path fill="currentColor" d="M12.278 13H11v2h1.278a1.936 1.936 0 0 1 0-2zM9 15v-2H5.722a1.936 1.936 0 0 1 0 2H9z" opacity=".25"/><path fill="currentColor" d="M18.278 19h-6.556a1.936 1.936 0 0 1 0 2h6.556a1.936 1.936 0 0 1 0-2zm0-8a1.936 1.936 0 0 1 0-2h-6.556a1.936 1.936 0 0 1 0 2h6.556z" opacity=".5"/><path fill="currentColor" d="M15 9V5.722a1.936 1.936 0 0 1-2 0V9h2zm-2 2v1.278a1.936 1.936 0 0 1 2 0V11h-2z" opacity=".25"/>`
	oktaInnerSVG                 = `<path fill="currentColor" d="M11.998 2a10 10 0 1 0 10 10a9.995 9.995 0 0 0-10-10Zm0 15.01a5.01 5.01 0 1 1 5.01-5.01a5.014 5.014 0 0 1-5.01 5.01Z"/>`
	operaInnerSVG                = `<path fill="currentColor" d="M11.996 2c-5.462 0-9.278 3.958-9.278 9.899C2.718 17.189 6.43 22 12.004 22c5.567 0 9.278-4.819 9.278-10.101c0-5.94-3.824-9.899-9.286-9.899Zm0 18.384c-3.397 0-3.77-5.013-3.77-8.71V11.6c0-3.996.598-8.23 3.748-8.23s3.786 4.361 3.786 8.357c0 3.696-.367 8.657-3.764 8.657Z"/>`
	operaAltInnerSVG             = `<path fill="currentColor" d="M11.996 2c-5.462 0-9.278 3.958-9.278 9.899C2.718 17.189 6.43 22 12.004 22c5.567 0 9.278-4.819 9.278-10.101c0-5.94-3.824-9.899-9.286-9.899Zm0 16c-2.397 0-2.66-3.536-2.66-6.143v-.052C9.336 8.987 9.758 6 11.979 6s2.67 3.076 2.67 5.894c0 2.607-.258 6.106-2.654 6.106Z"/>`
	padlockInnerSVG              = `<path fill="currentColor" d="M9 7c0-1.7 1.3-3 3-3s3 1.3 3 3v2h2V7c0-2.8-2.2-5-5-5S7 4.2 7 7v2h2V7zm4.5 7.5c0-.8-.7-1.5-1.5-1.5s-1.5.7-1.5 1.5c0 .4.2.8.5 1.1V17c0 .6.4 1 1 1s1-.4 1-1v-1.4c.3-.3.5-.7.5-1.1z" opacity=".5"/><path fill="currentColor" d="M17 9H7c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3zm-4 6.6V17c0 .6-.4 1-1 1s-1-.4-1-1v-1.4c-.3-.3-.5-.7-.5-1.1c0-.8.7-1.5 1.5-1.5s1.5.7 1.5 1.5c0 .4-.2.8-.5 1.1z"/>`
	pagelinesInnerSVG            = `<path fill="currentColor" d="M15.1 14.006a10.743 10.743 0 0 1-.895 2.275a3.538 3.538 0 0 1 1.36-.456c1.957-.21 3.75 1.543 3.75 1.543s-1.38 2.092-3.34 2.299a4.954 4.954 0 0 1-3.337-1.194A10.786 10.786 0 0 1 4.67 22a.672.672 0 0 1 0-1.344a9.419 9.419 0 0 0 5.845-2.032a3.745 3.745 0 0 1-1.735-.243c-1.816-.76-2.536-3.162-2.536-3.162s2.215-1.167 4.028-.404a4.394 4.394 0 0 1 2.012 1.98a9.432 9.432 0 0 0 1.458-2.9a4.72 4.72 0 0 1-3.258-.076c-1.773-.852-2.376-3.284-2.376-3.284s2.275-1.056 4.049-.203a4.396 4.396 0 0 1 1.9 2.058a9.323 9.323 0 0 0 .076-1.01a6.819 6.819 0 0 1-2.518-4.03C11.268 4.582 13.71 2 13.71 2s3.002 1.901 3.346 4.664a6.993 6.993 0 0 1-1.583 4.726a8.957 8.957 0 0 1-.06.99a4.206 4.206 0 0 1 1.7-1.822c1.75-.905 4.054.085 4.054.085s-.532 2.449-2.282 3.35a5.154 5.154 0 0 1-3.786.013Zm0 0"/>`
	pagerdutyInnerSVG            = `<path fill="currentColor" d="M6 16.673h2.93V22H6zM17.034 2.977C15.46 2.139 14.37 2 11.796 2H6v12.124h5.77c2.296 0 4.008-.14 5.517-1.141a5.769 5.769 0 0 0 2.499-4.997a5.487 5.487 0 0 0-2.752-5.01Zm-4.591 8.61H8.93V4.6l3.31-.026c3.018-.038 4.527 1.028 4.527 3.437c0 2.588-1.864 3.577-4.324 3.577Z"/>`
	paperclipInnerSVG            = `<path fill="currentColor" d="M8.892 21.854a6.25 6.25 0 0 1-4.42-10.67l7.955-7.955a4.5 4.5 0 0 1 6.364 6.364l-6.895 6.894a2.816 2.816 0 0 1-3.89 0a2.75 2.75 0 0 1 .002-3.888l5.126-5.127a1 1 0 1 1 1.414 1.414l-5.126 5.127a.75.75 0 0 0 0 1.06a.768.768 0 0 0 1.06 0l6.895-6.894a2.503 2.503 0 0 0 0-3.535a2.56 2.56 0 0 0-3.536 0l-7.955 7.955a4.25 4.25 0 1 0 6.01 6.01l6.188-6.187a1 1 0 1 1 1.414 1.414l-6.187 6.186a6.206 6.206 0 0 1-4.42 1.832Z"/>`
	paragraphInnerSVG            = `<path fill="currentColor" d="M13 15.5H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm8-5H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2z"/>`
	paypalInnerSVG               = `<path fill="currentColor" d="M8.882 19.94a1 1 0 0 1-.988.843H4.062a1.533 1.533 0 0 1-1.515-1.785l2.59-16.406A1.892 1.892 0 0 1 7 1h6.214c2.56 0 4.408.62 5.492 1.843a4.96 4.96 0 0 1 1.08 4.395c-.021.135-.043.27-.075.418c-.823 4.218-3.655 6.457-8.186 6.457H9.807l-.925 5.827zm-4.36-.628l-.001.006l.001-.006zM7.113 2.897v.002v-.002z" opacity=".5"/><path fill="currentColor" d="M20.437 7.104a4.098 4.098 0 0 0-.545-.508a4.956 4.956 0 0 1-.106.642c-.021.135-.043.27-.075.418c-.823 4.218-3.655 6.457-8.186 6.457H9.807l-.925 5.827a1 1 0 0 1-.988.843H6.727l-.082.52A1.467 1.467 0 0 0 8.093 23h3.234a1.76 1.76 0 0 0 1.751-1.469l.64-4.031l.012-.055h.298c4.033 0 6.551-1.993 7.286-5.762a5.15 5.15 0 0 0-.877-4.578z"/>`
	pentagonInnerSVG             = `<path fill="currentColor" d="M17.562 21.56H6.437a1 1 0 0 1-.95-.692l-3.438-10.58a.999.999 0 0 1 .363-1.117l9-6.54a.996.996 0 0 1 1.176 0l9 6.54a.999.999 0 0 1 .363 1.117l-3.437 10.58a1 1 0 0 1-.952.692Z"/>`
	plusSquareInnerSVG           = `<path fill="currentColor" d="M17 11h-4V7a1 1 0 0 0-2 0v4H7a1 1 0 0 0 0 2h4v4a1 1 0 0 0 2 0v-4h4a1 1 0 0 0 0-2Z"/><path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm-4 11h-4v4a1 1 0 0 1-2 0v-4H7a1 1 0 0 1 0-2h4V7a1 1 0 0 1 2 0v4h4a1 1 0 0 1 0 2Z" opacity=".5"/>`
	polygonInnerSVG              = `<path fill="currentColor" d="M16.5 20.794h-9a1 1 0 0 1-.866-.5l-4.5-7.794a1.002 1.002 0 0 1 0-1l4.5-7.794a1 1 0 0 1 .866-.5h9a1 1 0 0 1 .866.5l4.5 7.794a1.002 1.002 0 0 1 0 1l-4.5 7.794a1 1 0 0 1-.866.5Z"/>`
	previousInnerSVG             = `<path fill="currentColor" d="M16 17a.997.997 0 0 1-.707-.293l-4-4a1 1 0 0 1 0-1.414l4-4a1 1 0 0 1 1.414 1.414L13.414 12l3.293 3.293A1 1 0 0 1 16 17zm-8 0a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1z"/>`
	processInnerSVG              = `<path fill="currentColor" d="M8.625 8.5h-4.5a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v3.5h3.5a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M21 13a1 1 0 0 1-1-1A7.995 7.995 0 0 0 5.08 8.001a1 1 0 0 1-1.731-1.002A9.995 9.995 0 0 1 22 12a1 1 0 0 1-1 1zm-1.125 9a1 1 0 0 1-1-1v-3.5h-3.5a1 1 0 0 1 0-2h4.5a1 1 0 0 1 1 1V21a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M12 22A10.012 10.012 0 0 1 2 12a1 1 0 0 1 2 0a7.995 7.995 0 0 0 14.92 3.999a1 1 0 0 1 1.731 1.002A10.032 10.032 0 0 1 12 22Z"/>`
	reactInnerSVG                = `<path fill="currentColor" d="M19.108 12.376q-.176-.201-.371-.403q.136-.144.264-.287c1.605-1.804 2.283-3.614 1.655-4.701c-.602-1.043-2.393-1.354-4.636-.918q-.331.065-.659.146q-.063-.216-.133-.43C14.467 3.49 13.238 1.999 11.982 2c-1.204 0-2.368 1.397-3.111 3.558q-.11.32-.203.644q-.219-.054-.44-.1c-2.366-.485-4.271-.165-4.898.924c-.601 1.043.027 2.75 1.528 4.472q.224.255.46.5c-.186.19-.361.381-.525.571c-1.465 1.698-2.057 3.376-1.457 4.415c.62 1.074 2.498 1.425 4.785.975q.278-.055.553-.124q.1.351.221.697C9.635 20.649 10.792 22 11.992 22c1.24 0 2.482-1.453 3.235-3.659c.06-.174.115-.355.169-.541q.355.088.715.156c2.203.417 3.952.09 4.551-.95c.619-1.075-.02-2.877-1.554-4.63ZM4.07 7.452c.386-.67 1.943-.932 3.986-.512q.196.04.399.09a20.464 20.464 0 0 0-.422 2.678A20.887 20.887 0 0 0 5.93 11.4q-.219-.227-.427-.465C4.216 9.461 3.708 8.081 4.07 7.452Zm3.887 5.728c-.51-.387-.985-.783-1.416-1.181c.43-.396.905-.79 1.415-1.176q-.028.589-.027 1.179q0 .59.028 1.178Zm0 3.94a7.237 7.237 0 0 1-2.64.094a1.766 1.766 0 0 1-1.241-.657c-.365-.63.111-1.978 1.364-3.43q.236-.273.488-.532a20.49 20.49 0 0 0 2.107 1.7a20.802 20.802 0 0 0 .426 2.712q-.25.063-.505.114Zm7.1-8.039q-.503-.317-1.018-.613q-.508-.292-1.027-.563a18.7 18.7 0 0 1 1.739-.635a18.218 18.218 0 0 1 .306 1.811ZM9.68 5.835c.636-1.85 1.578-2.98 2.304-2.98c.773-.001 1.777 1.218 2.434 3.197q.064.194.12.39a20.478 20.478 0 0 0-2.526.97a20.061 20.061 0 0 0-2.52-.981q.087-.3.188-.596Zm-.4 1.424a18.307 18.307 0 0 1 1.73.642q-1.052.542-2.048 1.181c.08-.638.187-1.249.318-1.823Zm-.317 7.66q.497.319 1.009.613q.522.3 1.057.576a18.196 18.196 0 0 1-1.744.665a19.144 19.144 0 0 1-.322-1.853Zm5.456 3.146a7.236 7.236 0 0 1-1.238 2.333a1.766 1.766 0 0 1-1.188.748c-.729 0-1.658-1.085-2.29-2.896q-.112-.321-.206-.648a20.109 20.109 0 0 0 2.53-1.01a20.8 20.8 0 0 0 2.547.979q-.072.249-.155.494Zm.362-1.324a19.267 19.267 0 0 1-1.762-.646q.509-.267 1.025-.565q.53-.306 1.032-.627a18.152 18.152 0 0 1-.295 1.838Zm.447-4.743q0 .911-.057 1.82c-.493.334-1.013.66-1.554.972c-.54.311-1.073.597-1.597.856q-.827-.396-1.622-.854q-.79-.455-1.544-.969q-.07-.91-.07-1.822q0-.911.068-1.821a24.168 24.168 0 0 1 3.158-1.823q.816.397 1.603.851q.79.454 1.55.959q.065.914.065 1.831Zm.956-5.093c1.922-.373 3.37-.122 3.733.507c.387.67-.167 2.148-1.554 3.706q-.115.129-.238.259a20.061 20.061 0 0 0-2.144-1.688a20.04 20.04 0 0 0-.405-2.649q.31-.076.608-.135Zm-.13 3.885a18.164 18.164 0 0 1 1.462 1.188a18.12 18.12 0 0 1-1.457 1.208q.023-.594.023-1.188q0-.604-.028-1.208Zm3.869 5.789c-.364.631-1.768.894-3.653.538q-.324-.061-.664-.146a20.069 20.069 0 0 0 .387-2.682a19.94 19.94 0 0 0 2.137-1.715q.177.183.336.364a7.234 7.234 0 0 1 1.403 2.238a1.766 1.766 0 0 1 .054 1.403Zm-8.819-6.141a1.786 1.786 0 1 0 2.44.654a1.786 1.786 0 0 0-2.44-.654Z"/>`
	recordAudioInnerSVG          = `<circle cx="12" cy="12" r="6" fill="currentColor" opacity=".5"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2Zm0 16a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6Z"/>`
	redditAlienAltInnerSVG       = `<path fill="currentColor" d="M18.893 7a3.014 3.014 0 0 1-3-3.022a3 3 0 0 1 6-.023v.023a3.014 3.014 0 0 1-3 3.022zM11.8 23c-6.168 0-11-3.075-11-7s4.832-7 11-7s11 3.075 11 7s-4.832 7-11 7z" opacity=".5"/><path fill="currentColor" d="M20.499 8.292a3.772 3.772 0 0 0-4.497 1.223c2.952.756 5.224 2.295 6.228 4.234c.1-.161.194-.326.269-.5a3.779 3.779 0 0 0-2-4.957zm-8.7.708c.392 0 .777.013 1.157.037l.856-5.705l2.087.71c0-.022-.006-.042-.006-.064a2.972 2.972 0 0 1 .653-1.828l-3.224-1.096a.999.999 0 0 0-1.312.798l-1.076 7.175c.287-.014.573-.027.866-.027zm-4.207.516c-.088-.117-.172-.238-.273-.344a3.777 3.777 0 0 0-5.952 4.582c1.002-1.94 3.273-3.48 6.225-4.238z" opacity=".25"/><path fill="currentColor" d="M11.842 19.5a5.11 5.11 0 0 1-3.781-1.218a1 1 0 0 1 1.416-1.414c.68.5 1.525.726 2.365.632a3.375 3.375 0 0 0 2.368-.633a1 1 0 1 1 1.414 1.416a5.12 5.12 0 0 1-3.782 1.217zM10 14.002a1 1 0 1 0-1 1c.552-.001 1-.448 1-1zm6 0a1 1 0 1 0-1 1c.552-.001 1-.448 1-1z"/>`
	redoInnerSVG                 = `<path fill="currentColor" d="M19.875 8.5h-4.5a1 1 0 0 1 0-2h3.5V3a1 1 0 0 1 2 0v4.5a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M12 22a10 10 0 1 1 8.651-15.001a1 1 0 0 1-1.73 1.002A7.99 7.99 0 1 0 20 12a1 1 0 0 1 2 0a10.011 10.011 0 0 1-10 10Z"/>`
	refreshInnerSVG              = `<circle cx="12" cy="12" r="3" fill="currentColor"/><path fill="currentColor" d="M12 2a10.016 10.016 0 0 0-7 2.877V3a1 1 0 1 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2H6.218A7.98 7.98 0 0 1 20 12a1 1 0 0 0 2 0A10.012 10.012 0 0 0 12 2zm7.989 13.5h-4.5a1 1 0 0 0 0 2h2.293A7.98 7.98 0 0 1 4 12a1 1 0 0 0-2 0a9.986 9.986 0 0 0 16.989 7.133V21a1 1 0 0 0 2 0v-4.5a1 1 0 0 0-1-1z" opacity=".5"/>`
	repeatInnerSVG               = `<path fill="currentColor" d="M11.498 22a.997.997 0 0 1-.707-.293l-2.5-2.5a1 1 0 0 1 0-1.414l2.5-2.5a1 1 0 0 1 1.414 1.414L10.412 18.5l1.793 1.793A1 1 0 0 1 11.498 22z" opacity=".5"/><path fill="currentColor" d="M21 4.5h-2a1 1 0 0 0 0 2h1v11h-8.588l-1 1l1 1H21a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1z" opacity=".5"/><path fill="currentColor" d="M12.5 2c.265 0 .52.105.707.293l2.5 2.5a1 1 0 0 1 0 1.414l-2.5 2.5a1 1 0 0 1-1.414-1.414L13.586 5.5l-1.793-1.793A1 1 0 0 1 12.5 2z"/><path fill="currentColor" d="M5 17.5H4v-11h8.586l1-1l-1-1H3a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2z"/>`
	rightIndentAltInnerSVG       = `<path fill="currentColor" d="M21 19h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2zm0-4h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2zm0-8h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2zm0 4h-8a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M9 19a1 1 0 0 1-1-1V6a1 1 0 0 1 2 0v12a1 1 0 0 1-1 1zm-6-4.333a1 1 0 0 1-.64-1.769L3.438 12l-1.078-.898a1 1 0 0 1 1.28-1.538l2 1.667a1 1 0 0 1 0 1.538l-2 1.667a.999.999 0 0 1-.64.231z"/>`
	rocketInnerSVG               = `<path fill="currentColor" d="m17.737 14.622l-2.426 2.23a11.603 11.603 0 0 1-4.299 2.37l.644 3.004a1 1 0 0 0 1.469.661l3.905-2.202a3.035 3.035 0 0 0 1.375-3.304zM7.266 8.776l2.088-2.48l-2.604-.628a2.777 2.777 0 0 0-3.387 1.357l-2.2 3.9a1 1 0 0 0 .661 1.469l3.073.659a12.887 12.887 0 0 1 2.369-4.277zm9.468.04a1.5 1.5 0 1 0-1.5-1.5a1.5 1.5 0 0 0 1.5 1.5z"/><path fill="currentColor" d="M22.601 2.062a1 1 0 0 0-.713-.713A11.249 11.249 0 0 0 10.47 4.972L7.266 8.776a12.936 12.936 0 0 0-2.924 6.71a1 1 0 0 0 .284.837l3.1 3.1a1 1 0 0 0 .708.293a.838.838 0 0 0 .086-.004a11.847 11.847 0 0 0 6.79-2.86l3.664-3.368A11.204 11.204 0 0 0 22.6 2.062Zm-5.867 6.754a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Z" opacity=".5"/>`
	rulerInnerSVG                = `<path fill="currentColor" d="M10.586 20.485L7.05 16.95a1 1 0 0 1 0-1.414a1 1 0 0 1 1.415 0L12 19.07zm2.828-2.828l-2.121-2.121a1 1 0 0 1 0-1.415a1 1 0 0 1 1.414 0l2.122 2.122zm2.829-2.828l-3.536-3.536a1 1 0 0 1 0-1.414a1 1 0 0 1 1.414 0l3.536 3.535zM19.07 12l-2.12-2.12a1 1 0 0 1 0-1.414a1 1 0 0 1 1.414 0l2.121 2.12z"/><path fill="currentColor" d="M22.606 7.05L16.95 1.395a1 1 0 0 0-1.414 0L1.394 15.535a1 1 0 0 0 0 1.414l5.656 5.657a1 1 0 0 0 1.414 0l2.122-2.12L7.05 16.95a1 1 0 1 1 1.415-1.414L12 19.07l1.414-1.414l-2.121-2.121a1 1 0 0 1 1.414-1.415l2.121 2.122l1.415-1.415l-3.536-3.535a1 1 0 0 1 1.414-1.414l3.536 3.535L19.07 12l-2.12-2.12a1 1 0 0 1 1.414-1.414l2.121 2.12l2.121-2.12a1 1 0 0 0 0-1.414Z" opacity=".5"/>`
	rulerCombinedInnerSVG        = `<path fill="currentColor" d="M14 10h-2V7a1 1 0 0 1 1-1a1 1 0 0 1 1 1zM9 6a1 1 0 0 0-1 1v1H7a1 1 0 0 0 0 2h3V7a1 1 0 0 0-1-1zm1 6v2H7a1 1 0 0 1-1-1a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M10 12v2H7a1 1 0 0 1-1-1a1 1 0 0 1 1-1Z"/><path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-3H7a1 1 0 0 1 0-2h3v-2H7a1 1 0 0 1 0-2h3v-2H7a1 1 0 0 1 0-2h1V7a1 1 0 0 1 2 0v3h2V7a1 1 0 0 1 2 0v3h2V7a1 1 0 0 1 2 0v3h3a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Z" opacity=".5"/><path fill="currentColor" d="M10 16v2H7a1 1 0 0 1-1-1a1 1 0 0 1 1-1zm8-6h-2V7a1 1 0 0 1 1-1a1 1 0 0 1 1 1z"/>`
	sanitizerInnerSVG            = `<path fill="currentColor" d="M5 5a1 1 0 0 1-.707-1.707l.829-.829A4.967 4.967 0 0 1 8.657 1H17a1 1 0 0 1 0 2H8.657a3.022 3.022 0 0 0-2.121.878l-.829.829A.997.997 0 0 1 5 5Z"/><path fill="currentColor" d="M10 3v2.5l.4-.3A1 1 0 0 1 11 5h4a1 1 0 0 1 .6.2l.4.3V3Z" opacity=".5"/><circle cx="13" cy="15" r="2" fill="currentColor"/><path fill="currentColor" d="M13 18a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Z"/><path fill="currentColor" d="m18.8 7.6l-3.2-2.4A1 1 0 0 0 15 5h-4a1 1 0 0 0-.6.2L7.2 7.6A3.016 3.016 0 0 0 6 10v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V10a3.015 3.015 0 0 0-1.2-2.4ZM13 18a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Z" opacity=".25"/>`
	sanitizerAltInnerSVG         = `<path fill="currentColor" d="M12 3v1h2V3h1a1 1 0 0 0 0-2h-4.764a4.593 4.593 0 0 0-4.13 2.553a1 1 0 0 0 1.789.894A2.603 2.603 0 0 1 10.235 3Z" opacity=".5"/><path fill="currentColor" d="M16 5a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v3h6zm-2 12h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2z"/><path fill="currentColor" d="M16 8h-6a3.003 3.003 0 0 0-3 3v9a3.003 3.003 0 0 0 3 3h6a3.003 3.003 0 0 0 3-3v-9a3.003 3.003 0 0 0-3-3Zm-2 9h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2Z" opacity=".25"/>`
	sceneryInnerSVG              = `<path fill="currentColor" d="M13.5 9a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z" opacity=".25"/><path fill="currentColor" d="M19 2H5a3.009 3.009 0 0 0-3 3v8.86l3.88-3.88a3.075 3.075 0 0 1 4.24 0l2.871 2.887l.888-.888a3.008 3.008 0 0 1 4.242 0L22 15.86V5a3.009 3.009 0 0 0-3-3zm-5.5 7a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z" opacity=".5"/><path fill="currentColor" d="M10.12 9.98a3.075 3.075 0 0 0-4.24 0L2 13.86V19a3.009 3.009 0 0 0 3 3h14a3 3 0 0 0 2.16-.92L10.12 9.98z"/><path fill="currentColor" d="m22 15.858l-3.879-3.879a3.008 3.008 0 0 0-4.242 0l-.888.888l8.165 8.209c.542-.555.845-1.3.844-2.076v-3.142z" opacity=".25"/>`
	scheduleInnerSVG             = `<path fill="currentColor" d="M7 6a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v2a1 1 0 0 1-.999 1H7zm10 0a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v2a1 1 0 0 1-.999 1H17z" opacity=".5"/><path fill="currentColor" d="M19 4h-1v1a1 1 0 0 1-2 0V4H8v1a1 1 0 0 1-2 0V4H5a3 3 0 0 0-3 3v2h20V7a3 3 0 0 0-3-3z"/><circle cx="7" cy="13" r="1" fill="currentColor"/><circle cx="7" cy="17" r="1" fill="currentColor"/><circle cx="12" cy="13" r="1" fill="currentColor"/><circle cx="12" cy="17" r="1" fill="currentColor"/><circle cx="17" cy="13" r="1" fill="currentColor"/><circle cx="17" cy="17" r="1" fill="currentColor"/><path fill="currentColor" d="M2 9v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9H2zm5 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm5 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm5 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" opacity=".5"/>`
	shieldPlusInnerSVG           = `<path fill="currentColor" d="M14 11h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2Z"/><path fill="currentColor" d="M19.63 3.65a1.002 1.002 0 0 0-.835-.203a7.985 7.985 0 0 1-6.223-1.267a.999.999 0 0 0-1.144 0a7.98 7.98 0 0 1-6.223 1.267A1 1 0 0 0 4 4.427v7.456a9.019 9.019 0 0 0 3.769 7.324l3.65 2.607a1 1 0 0 0 1.162 0l3.65-2.607A9.017 9.017 0 0 0 20 11.883V4.426a1.001 1.001 0 0 0-.37-.776ZM14 13h-1v1a1 1 0 0 1-2 0v-1h-1a1 1 0 0 1 0-2h1v-1a1 1 0 0 1 2 0v1h1a1 1 0 0 1 0 2Z" opacity=".5"/>`
	signInInnerSVG               = `<path fill="currentColor" d="M20 11h-8.586l2.293-2.293a1 1 0 0 0-1.414-1.414l-4 4a1 1 0 0 0 0 1.414l4 4a1 1 0 0 0 1.414-1.414L11.414 13H20Z"/><path fill="currentColor" d="M11.414 11H20V5a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h10a3.003 3.003 0 0 0 3-3v-6h-8.586l2.293 2.293a1 1 0 1 1-1.414 1.414l-4-4a1 1 0 0 1 0-1.414l4-4a1 1 0 0 1 1.414 1.414Z" opacity=".5"/>`
	signInAltInnerSVG            = `<path fill="currentColor" d="M21 12c0-.34-.02-.67-.05-1H12.5V9.5a1 1 0 0 0-1.707-.707l-2.5 2.5a1 1 0 0 0 0 1.414l2.5 2.5A1 1 0 0 0 12.5 14.5V13h8.45c.03-.33.05-.66.05-1Z"/><path fill="currentColor" d="M12.5 13v1.5a1 1 0 0 1-1.707.707l-2.5-2.5a1 1 0 0 1 0-1.414l2.5-2.5A1 1 0 0 1 12.5 9.5V11h8.45a10 10 0 1 0 0 2Z" opacity=".5"/>`
	signOutAltInnerSVG           = `<path fill="currentColor" d="m15.707 11.293l-4-4a1 1 0 0 0-1.414 1.414L12.586 11H2.05c-.03.33-.05.66-.05 1s.02.67.05 1h10.536l-2.293 2.293a1 1 0 1 0 1.414 1.414l4-4a1 1 0 0 0 0-1.414Z"/><path fill="currentColor" d="M12 2a10 10 0 0 0-9.95 9h10.536l-2.293-2.293a1 1 0 0 1 1.414-1.414l4 4a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414-1.414L12.586 13H2.05A10 10 0 1 0 12 2Z" opacity=".5"/>`
	signalAltInnerSVG            = `<path fill="currentColor" d="M5 22a1 1 0 0 1-1-1v-2a1 1 0 0 1 2 0v2a1 1 0 0 1-1 1zm5 0a1 1 0 0 1-1-1v-6a1 1 0 0 1 2 0v6a1 1 0 0 1-1 1zm5 0a1 1 0 0 1-1-1V10a1 1 0 0 1 2 0v11a1 1 0 0 1-1 1zm5 0a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v18a1 1 0 0 1-1 1z"/>`
	signalAltThreeInnerSVG       = `<path fill="currentColor" d="M6 23H2a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1Z" opacity=".25"/><path fill="currentColor" d="M14 23h-4a1 1 0 0 1-1-1V10a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1Z" opacity=".5"/><path fill="currentColor" d="M22 23h-4a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v20a1 1 0 0 1-1 1Z"/>`
	signinInnerSVG               = `<path fill="currentColor" d="M19 11h-5.586l1.293-1.293a1 1 0 0 0-1.414-1.414l-3 3a1 1 0 0 0 0 1.414l3 3a1 1 0 0 0 1.414-1.414L13.414 13H19a1 1 0 0 0 0-2Z"/><path fill="currentColor" d="m13.414 13l1.293 1.293a1 1 0 1 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 0 1 1.414 1.414L13.414 11H18V5a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h8a3.003 3.003 0 0 0 3-3v-6Z" opacity=".5"/>`
	signoutInnerSVG              = `<path fill="currentColor" d="m21.207 11.293l-3-3a1 1 0 1 0-1.414 1.415L18.086 11H12.5a1 1 0 0 0 0 2h5.586l-1.293 1.293a1 1 0 1 0 1.414 1.414l3-3a1 1 0 0 0 0-1.415Z"/><path fill="currentColor" d="M12.5 13a1 1 0 0 1 0-2h4V5a3.003 3.003 0 0 0-3-3h-8a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h8a3.003 3.003 0 0 0 3-3v-6Z" opacity=".5"/>`
	skypeInnerSVG                = `<path fill="currentColor" d="M16.44 15.993c-.414.555-.978.98-1.625 1.225a6.34 6.34 0 0 1-2.52.447a6.217 6.217 0 0 1-2.898-.612a3.733 3.733 0 0 1-1.32-1.178a2.574 2.574 0 0 1-.494-1.413a.88.88 0 0 1 .307-.684a1.09 1.09 0 0 1 .776-.282a.944.944 0 0 1 .637.212c.197.184.35.409.447.659c.121.314.288.608.495.873c.19.248.441.443.73.564c.395.167.82.247 1.249.236a2.922 2.922 0 0 0 1.72-.447c.402-.236.652-.665.66-1.131a1.135 1.135 0 0 0-.354-.871a2.185 2.185 0 0 0-.92-.52c-.376-.117-.895-.235-1.53-.376a13.99 13.99 0 0 1-2.144-.636a3.348 3.348 0 0 1-1.366-1.013a2.474 2.474 0 0 1-.495-1.578c0-.579.19-1.142.542-1.602c.399-.497.93-.873 1.53-1.084a6.652 6.652 0 0 1 2.38-.376a6.403 6.403 0 0 1 1.885.258c.476.138.923.362 1.318.66c.316.235.58.532.778.872c.151.265.232.565.236.87c0 .269-.11.525-.307.708a.991.991 0 0 1-.753.306a.973.973 0 0 1-.636-.189a2.382 2.382 0 0 1-.471-.611a2.937 2.937 0 0 0-.778-.967a2.376 2.376 0 0 0-1.46-.353a2.703 2.703 0 0 0-1.508.377a1.076 1.076 0 0 0-.565.896a.958.958 0 0 0 .188.565c.146.175.332.312.542.4c.215.117.445.204.683.26c.236.07.613.164 1.154.282c.66.142 1.273.306 1.815.471c.49.145.958.36 1.389.636c.367.24.673.563.895.942c.227.428.34.906.33 1.39a2.89 2.89 0 0 1-.542 1.814z" opacity=".5"/><path fill="currentColor" d="M21.435 14.156a9.589 9.589 0 0 0 .211-2.027a9.477 9.477 0 0 0-9.53-9.422h-.01a9.117 9.117 0 0 0-1.625.141A5.536 5.536 0 0 0 2 7.466c0 .97.26 1.921.753 2.756c-.122.62-.185 1.252-.189 1.884a9.339 9.339 0 0 0 9.54 9.258a8.567 8.567 0 0 0 1.743-.166a5.58 5.58 0 0 0 2.616.802a5.433 5.433 0 0 0 4.97-7.844zm-4.995 1.837c-.414.555-.978.98-1.625 1.225a6.34 6.34 0 0 1-2.52.447a6.217 6.217 0 0 1-2.898-.612a3.732 3.732 0 0 1-1.32-1.178a2.574 2.574 0 0 1-.494-1.413a.88.88 0 0 1 .307-.684a1.09 1.09 0 0 1 .776-.282a.944.944 0 0 1 .637.212c.197.184.35.409.447.659c.121.314.288.608.495.873c.19.248.441.443.73.564c.395.167.82.247 1.249.236a2.922 2.922 0 0 0 1.72-.447c.402-.236.653-.665.66-1.131a1.135 1.135 0 0 0-.354-.871a2.184 2.184 0 0 0-.92-.52c-.376-.117-.895-.235-1.53-.376a13.992 13.992 0 0 1-2.144-.636a3.348 3.348 0 0 1-1.366-1.013a2.474 2.474 0 0 1-.495-1.578c0-.579.19-1.142.542-1.602c.399-.497.93-.873 1.53-1.084a6.65 6.65 0 0 1 2.38-.376a6.402 6.402 0 0 1 1.885.258c.476.138.923.362 1.319.66c.315.235.58.532.777.872c.151.265.232.565.236.871c0 .268-.11.524-.307.707a.991.991 0 0 1-.753.306a.974.974 0 0 1-.636-.189a2.383 2.383 0 0 1-.471-.611a2.937 2.937 0 0 0-.778-.967a2.376 2.376 0 0 0-1.46-.353a2.703 2.703 0 0 0-1.508.377a1.075 1.075 0 0 0-.565.896a.958.958 0 0 0 .188.565c.146.175.332.312.542.4c.215.117.445.204.683.26c.236.07.613.164 1.154.282c.66.142 1.273.306 1.815.471c.49.145.958.36 1.389.636c.367.24.673.563.895.942c.227.428.34.906.33 1.39a2.89 2.89 0 0 1-.542 1.814z"/>`
	skypeAltInnerSVG             = `<path fill="currentColor" d="M16.5 23a6.501 6.501 0 0 1-2.809-.639A10.491 10.491 0 0 1 1.64 10.31a6.499 6.499 0 0 1 8.67-8.67a10.491 10.491 0 0 1 12.05 12.05A6.499 6.499 0 0 1 16.5 23Z" opacity=".5"/><path fill="currentColor" d="M16.136 12.865a3.165 3.165 0 0 0-.997-1.052a5.076 5.076 0 0 0-1.29-.593c-.28-.085-.59-.168-.911-.247c-.28-.078-.612-.158-1.022-.248a9.315 9.315 0 0 1-1.436-.424a1.496 1.496 0 0 1-.616-.447a.843.843 0 0 1-.16-.566a.967.967 0 0 1 .205-.597a1.598 1.598 0 0 1 .7-.475A4.012 4.012 0 0 1 12.031 8a3.787 3.787 0 0 1 1.106.146a2.083 2.083 0 0 1 .663.322a1.235 1.235 0 0 1 .325.343a1 1 0 1 0 1.761-.948a3.147 3.147 0 0 0-.837-.958a4.006 4.006 0 0 0-1.319-.669A5.768 5.768 0 0 0 12.032 6a5.963 5.963 0 0 0-2.145.35A3.552 3.552 0 0 0 8.31 7.492a2.977 2.977 0 0 0-.604 1.797a2.839 2.839 0 0 0 .58 1.792a3.5 3.5 0 0 0 1.438 1.072a10.582 10.582 0 0 0 1.307.408l1.531.378c.248.064.487.129.706.196a3.023 3.023 0 0 1 .763.344a1.127 1.127 0 0 1 .363.368a1.201 1.201 0 0 1 .118.585a1.152 1.152 0 0 1-.214.732a1.763 1.763 0 0 1-.802.585a3.787 3.787 0 0 1-1.487.252a3.689 3.689 0 0 1-1.703-.344a1.756 1.756 0 0 1-.616-.547a1.016 1.016 0 0 1-.202-.503a1 1 0 0 0-2 0a2.94 2.94 0 0 0 .556 1.64a3.774 3.774 0 0 0 1.342 1.187a5.621 5.621 0 0 0 2.623.567a5.708 5.708 0 0 0 2.254-.405a3.71 3.71 0 0 0 1.665-1.273a3.146 3.146 0 0 0 .584-1.926a3.09 3.09 0 0 0-.375-1.53Z"/>`
	slackInnerSVG                = `<path fill="currentColor" d="M9.338 2a2 2 0 0 0 .001 4h1.996V4a2 2 0 0 0-1.997-2m0 5.333H4.016a2 2 0 0 0 0 4h5.322a2 2 0 0 0 0-4Z"/><path fill="currentColor" d="M21.98 9.333a1.996 1.996 0 1 0-3.993 0v2h1.997a1.998 1.998 0 0 0 1.996-2Zm-5.323 0V4a1.996 1.996 0 1 0-3.992 0v5.333a1.996 1.996 0 1 0 3.992 0Z" opacity=".25"/><path fill="currentColor" d="M14.661 22a2 2 0 0 0 0-4h-1.996v2a2 2 0 0 0 1.996 2Zm0-5.334h5.323a2 2 0 0 0 0-4h-5.322a2 2 0 0 0-.001 4Z" opacity=".7"/><path fill="currentColor" d="M2.02 14.666a1.996 1.996 0 1 0 3.993 0v-2H4.016a1.998 1.998 0 0 0-1.996 2Zm5.323 0V20a1.996 1.996 0 1 0 3.992 0v-5.332a1.996 1.996 0 1 0-3.992-.002" opacity=".5"/>`
	slackAltInnerSVG             = `<path fill="currentColor" d="M9.34 2a2 2 0 0 0 0 4h2V4a2 2 0 0 0-2-2m0 5.33H4a2 2 0 1 0 0 4h5.34a2 2 0 0 0 0-4" opacity=".25"/><path fill="currentColor" d="M22 9.33a2 2 0 1 0-4 0v2h2a2 2 0 0 0 2-2m-5.32 0V4a2 2 0 1 0-4 0v5.33a2 2 0 1 0 4 0" opacity=".5"/><path fill="currentColor" d="M14.66 22a2 2 0 0 0 0-4h-2v2a2 2 0 0 0 2 2m0-5.33H20a2 2 0 0 0 0-4h-5.34a2 2 0 0 0 0 4" opacity=".7"/><path fill="currentColor" d="M2 14.67a2 2 0 1 0 4 0v-2H4a2 2 0 0 0-2 2m5.32 0V20a2 2 0 1 0 4 0v-5.33a2 2 0 1 0-4 0"/>`
	snapchatAltInnerSVG          = `<path fill="currentColor" d="M3.808 11.651a1 1 0 0 0 1.06-.49a3.12 3.12 0 0 0 1.27.34a1.81 1.81 0 0 0 1.28-.5a1 1 0 0 0 .3-.77v-.62a9.82 9.82 0 0 1 .15-3.55a4.27 4.27 0 0 1 4-2.55h.4a4.26 4.26 0 0 1 4 2.55a10.18 10.18 0 0 1 .15 3.63v.54a1 1 0 0 0 .31.78a1.78 1.78 0 0 0 1.25.5a3.21 3.21 0 0 0 1.24-.37a1 1 0 0 0 .92.6a1 1 0 0 0 1-1a1.63 1.63 0 0 0-1.19-1.45a1.92 1.92 0 0 0-1.47 0a9.49 9.49 0 0 0-.36-4a6.23 6.23 0 0 0-5.93-3.79h-.4a6.22 6.22 0 0 0-5.79 3.73a9.43 9.43 0 0 0-.35 4.1l-.11-.05a1.9 1.9 0 0 0-2.54 1.2a1 1 0 0 0 .81 1.17zm18.19 4.46a4.67 4.67 0 0 1-3-2.17a1.001 1.001 0 1 0-1.66 1.12a7.66 7.66 0 0 0 2.4 2.33a4.121 4.121 0 0 1-.45.08a1.37 1.37 0 0 0-1.06 1.21a6 6 0 0 0-2 0a4.58 4.58 0 0 0-2 1a3.5 3.5 0 0 1-2.11.87h-.26a3.5 3.5 0 0 1-2.11-.87a4.49 4.49 0 0 0-1.93-.95a6.27 6.27 0 0 0-2 0a1.39 1.39 0 0 0-1.07-1.21l-.45-.12a6.78 6.78 0 0 0 2.43-2.4a1 1 0 1 0-1.73-1a5.54 5.54 0 0 1-.45.62a4.24 4.24 0 0 1-2.55 1.49a1.23 1.23 0 0 0-1 1.24c0 .183.037.363.11.53c.32.72 1.16 1.17 2.79 1.48v.12c0 .12.06.25.09.35a1.3 1.3 0 0 0 1.28 1a2.22 2.22 0 0 0 .6-.09a4.57 4.57 0 0 1 1.59 0a3 3 0 0 1 1.12.61a5.44 5.44 0 0 0 3.42 1.15h.16a5.43 5.43 0 0 0 3.28-1.25a3 3 0 0 1 1.12-.61a4.08 4.08 0 0 1 1.58.06c.205.036.412.056.62.06a1.25 1.25 0 0 0 1.24-.92c0-.12.07-.24.1-.36v-.12c1.63-.31 2.47-.75 2.76-1.42a1.31 1.31 0 0 0 .14-.51a1.25 1.25 0 0 0-1-1.32z"/>`
	snapchatGhostInnerSVG        = `<path fill="currentColor" d="M21.798 16.987c-2.867-.472-4.151-3.401-4.204-3.526l-.006-.011a1.07 1.07 0 0 1-.102-.898c.192-.454.83-.656 1.251-.79c.106-.033.205-.065.283-.096c.763-.3.918-.613.914-.822a.662.662 0 0 0-.5-.543l-.007-.002a.946.946 0 0 0-.356-.069a.755.755 0 0 0-.313.063a2.54 2.54 0 0 1-.955.266a.821.821 0 0 1-.53-.178l.032-.53l.004-.065a10.102 10.102 0 0 0-.24-4.035a5.248 5.248 0 0 0-4.874-3.14l-.402.005a5.24 5.24 0 0 0-4.864 3.137a10.09 10.09 0 0 0-.242 4.031q.02.299.036.598a.848.848 0 0 1-.584.178a2.453 2.453 0 0 1-1.014-.268a.575.575 0 0 0-.245-.049a.834.834 0 0 0-.81.533c-.082.43.532.743.906.89c.08.032.178.063.283.096c.422.134 1.06.336 1.252.79a1.072 1.072 0 0 1-.102.898l-.006.011a7.028 7.028 0 0 1-1.069 1.663A5.215 5.215 0 0 1 2.2 16.987a.24.24 0 0 0-.2.25a.38.38 0 0 0 .03.13c.177.411 1.059.75 2.553.981c.14.022.198.25.28.622c.032.15.066.305.113.465a.293.293 0 0 0 .32.228a2.485 2.485 0 0 0 .424-.06a5.53 5.53 0 0 1 1.12-.127a4.954 4.954 0 0 1 .809.068a3.877 3.877 0 0 1 1.535.784a4.443 4.443 0 0 0 2.69 1.06c.033 0 .067-.001.1-.004c.04.002.095.004.151.004a4.448 4.448 0 0 0 2.692-1.06a3.873 3.873 0 0 1 1.533-.784a4.973 4.973 0 0 1 .808-.068a5.593 5.593 0 0 1 1.12.119a2.391 2.391 0 0 0 .425.053h.024a.279.279 0 0 0 .295-.22a6.52 6.52 0 0 0 .114-.462c.081-.371.14-.598.28-.62c1.494-.23 2.377-.57 2.551-.978a.385.385 0 0 0 .032-.13a.24.24 0 0 0-.2-.25Z"/>`
	snapchatSquareInnerSVG       = `<path fill="currentColor" d="M19.174 15.652a4.522 4.522 0 0 1-3.079-2.582l-.004-.009a.784.784 0 0 1-.074-.657c.14-.332.607-.48.916-.578c.077-.024.15-.048.207-.07c.559-.22.672-.45.67-.602a.485.485 0 0 0-.368-.398l-.004-.002a.694.694 0 0 0-.26-.05a.552.552 0 0 0-.23.046a1.86 1.86 0 0 1-.7.195a.602.602 0 0 1-.387-.13l.023-.389l.003-.047a7.4 7.4 0 0 0-.176-2.955a3.843 3.843 0 0 0-3.568-2.298l-.295.002a3.837 3.837 0 0 0-3.562 2.298a7.388 7.388 0 0 0-.177 2.951l.027.438a.621.621 0 0 1-.428.13a1.796 1.796 0 0 1-.742-.195a.42.42 0 0 0-.18-.036a.61.61 0 0 0-.593.39c-.06.315.39.544.664.652c.057.023.13.046.207.07c.309.098.775.246.916.578c.073.22.046.46-.075.658l-.004.008c-.202.44-.465.85-.782 1.217a3.818 3.818 0 0 1-2.296 1.365a.176.176 0 0 0-.147.183c.002.033.01.065.023.095c.129.301.775.55 1.869.718c.102.016.145.183.205.456c.024.11.048.223.083.34c.023.107.124.18.234.167a1.82 1.82 0 0 0 .31-.044a3.566 3.566 0 0 1 1.413-.043c.413.105.797.3 1.124.574a3.253 3.253 0 0 0 1.97.776c.024 0 .048 0 .072-.003c.03.002.07.003.112.003a3.257 3.257 0 0 0 1.97-.776c.327-.273.71-.47 1.123-.574c.196-.033.393-.05.592-.05c.275.001.55.03.82.087c.102.022.207.035.311.04h.017a.204.204 0 0 0 .217-.163c.034-.115.059-.225.083-.337c.06-.272.103-.438.205-.454c1.094-.169 1.74-.417 1.868-.716a.28.28 0 0 0 .023-.096a.176.176 0 0 0-.146-.183z" opacity=".5"/><path fill="currentColor" d="M17.508 2h-11a4.5 4.5 0 0 0-4.5 4.5v11a4.5 4.5 0 0 0 4.5 4.5h11a4.5 4.5 0 0 0 4.5-4.5v-11a4.5 4.5 0 0 0-4.5-4.5zm1.79 13.93c-.129.3-.775.548-1.869.717c-.102.016-.146.182-.205.454c-.024.112-.05.222-.083.337a.204.204 0 0 1-.217.162h-.017a1.751 1.751 0 0 1-.31-.04a4.094 4.094 0 0 0-.821-.086c-.199 0-.396.017-.592.05c-.413.105-.796.3-1.123.574a3.257 3.257 0 0 1-1.97.776c-.042 0-.082-.001-.112-.003a.85.85 0 0 1-.073.003a3.253 3.253 0 0 1-1.97-.776a2.84 2.84 0 0 0-1.123-.574a3.63 3.63 0 0 0-.592-.05c-.276.003-.55.034-.82.093a1.822 1.822 0 0 1-.311.044a.214.214 0 0 1-.234-.167a4.753 4.753 0 0 1-.083-.34c-.06-.273-.103-.44-.205-.456c-1.094-.168-1.74-.417-1.869-.718a.279.279 0 0 1-.023-.095a.176.176 0 0 1 .147-.183a3.818 3.818 0 0 0 2.296-1.365c.317-.367.58-.776.783-1.217l.003-.008a.785.785 0 0 0 .075-.658c-.14-.332-.607-.48-.916-.578a2.904 2.904 0 0 1-.207-.07c-.274-.108-.724-.337-.664-.652a.61.61 0 0 1 .593-.39a.42.42 0 0 1 .18.036c.23.118.484.185.742.196a.622.622 0 0 0 .428-.131a47.507 47.507 0 0 0-.027-.438a7.4 7.4 0 0 1 .177-2.951a3.837 3.837 0 0 1 3.562-2.298l.295-.002a3.843 3.843 0 0 1 3.568 2.298a7.4 7.4 0 0 1 .176 2.955l-.003.047l-.023.389c.11.087.247.133.388.13a1.86 1.86 0 0 0 .7-.195a.552.552 0 0 1 .228-.046c.09 0 .178.017.261.05l.004.002a.485.485 0 0 1 .367.398c.003.153-.11.381-.669.602c-.057.022-.13.046-.207.07c-.31.098-.776.246-.916.578a.784.784 0 0 0 .074.657l.004.009a4.522 4.522 0 0 0 3.079 2.582a.176.176 0 0 1 .146.183a.281.281 0 0 1-.023.096z"/>`
	socialDistancingInnerSVG     = `<path fill="currentColor" d="M6 11a3.5 3.5 0 1 1 3.5-3.5A3.504 3.504 0 0 1 6 11Z" opacity=".25"/><path fill="currentColor" d="M8.64 9.772a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 1 14a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228zM18 11a3.5 3.5 0 1 1 3.5-3.5A3.504 3.504 0 0 1 18 11z" opacity=".5"/><path fill="currentColor" d="M20.64 9.772a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 13 14a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228Z" opacity=".7"/><path fill="currentColor" d="m21.207 18.293l-2-2a1 1 0 0 0-1.414 1.414l.293.293H15.5a1 1 0 0 0 0 2h2.586l-.293.293a1 1 0 1 0 1.414 1.414l2-2a1 1 0 0 0 0-1.414zM8.5 18H5.914l.293-.293a1 1 0 0 0-1.414-1.414l-2 2a1 1 0 0 0 0 1.414l2 2a1 1 0 0 0 1.414-1.414L5.914 20H8.5a1 1 0 0 0 0-2z"/>`
	sortingInnerSVG              = `<path fill="currentColor" d="M15 19.5a.997.997 0 0 1-.707-.293L12 16.914l-2.293 2.293a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.414 0l3 3A1 1 0 0 1 15 19.5zm-3-9a.997.997 0 0 1-.707-.293l-3-3a1 1 0 0 1 1.414-1.414L12 8.086l2.293-2.293a1 1 0 0 1 1.414 1.414l-3 3A.997.997 0 0 1 12 10.5z"/>`
	spaceKeyInnerSVG             = `<path fill="currentColor" d="M21 15H3a1 1 0 0 1-1-1v-4a1 1 0 0 1 2 0v3h16v-3a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1Z"/>`
	squareInnerSVG               = `<rect width="20" height="20" x="2" y="2" fill="currentColor" rx="1"/>`
	squareFullInnerSVG           = `<path fill="currentColor" d="M21 22H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1ZM4 20h16V4H4Z"/>`
	squareShapeInnerSVG          = `<rect width="20" height="20" x="2" y="2" fill="currentColor" rx="1"/>`
	squreShapeInnerSVG           = `<rect width="20" height="20" x="2" y="2" fill="currentColor" rx="5"/>`
	starInnerSVG                 = `<path fill="currentColor" d="M17.562 21.56a1.003 1.003 0 0 1-.465-.115L12 18.765l-5.097 2.68a1 1 0 0 1-1.451-1.054l.973-5.676l-4.123-4.02a1 1 0 0 1 .554-1.705l5.699-.828l2.548-5.164a1.042 1.042 0 0 1 1.794 0l2.548 5.164l5.699.828a1 1 0 0 1 .554 1.706l-4.123 4.019l.973 5.676a1 1 0 0 1-.986 1.169Z"/>`
	starHalfAltInnerSVG          = `<path fill="currentColor" d="M21.951 9.67a1 1 0 0 0-.807-.68l-5.699-.828l-2.548-5.164A.978.978 0 0 0 12 2.486v16.28l5.097 2.679a1 1 0 0 0 1.451-1.054l-.973-5.676l4.123-4.02a1 1 0 0 0 .253-1.025z" opacity=".5"/><path fill="currentColor" d="M11.103 2.998L8.555 8.162l-5.699.828a1 1 0 0 0-.554 1.706l4.123 4.019l-.973 5.676a1 1 0 0 0 1.45 1.054L12 18.765V2.503a1.028 1.028 0 0 0-.897.495z"/>`
	stepForwardInnerSVG          = `<path fill="currentColor" d="M8 17a1 1 0 0 1-.707-1.707L10.586 12L7.293 8.707a1 1 0 0 1 1.414-1.414l4 4a1 1 0 0 1 0 1.414l-4 4A.997.997 0 0 1 8 17zm8 0a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1z"/>`
	stethoscopeInnerSVG          = `<path fill="currentColor" d="M19 14a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zM8 15a6.007 6.007 0 0 1-6-6V3a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H4v5a4 4 0 0 0 8 0V4h-1a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v6a6.007 6.007 0 0 1-6 6z" opacity=".5"/><path fill="currentColor" d="M19 14a2.965 2.965 0 0 1-1-.184V15.5a4.5 4.5 0 0 1-9 0v-.59a5.58 5.58 0 0 1-2 0v.59a6.5 6.5 0 0 0 13 0v-1.684A2.965 2.965 0 0 1 19 14Z"/>`
	stethoscopeAltInnerSVG       = `<path fill="currentColor" d="M8 15a.998.998 0 0 1-.625-.22l-3.499-2.798A4.975 4.975 0 0 1 2 8.078V3a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H4v4.078a2.985 2.985 0 0 0 1.126 2.342L8 12.72l2.875-2.3A2.986 2.986 0 0 0 12 8.078V4h-1a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v5.078a4.976 4.976 0 0 1-1.876 3.904l-3.5 2.799A.998.998 0 0 1 8 15zm11-1a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3z" opacity=".5"/><path fill="currentColor" d="M19 14a2.965 2.965 0 0 1-1-.184V15.5a4.5 4.5 0 0 1-9 0v-1.02l-.375.3a1 1 0 0 1-1.25 0L7 14.48v1.02a6.5 6.5 0 0 0 13 0v-1.684A2.965 2.965 0 0 1 19 14Z"/>`
	storeSlashInnerSVG           = `<path fill="currentColor" d="M20 18.586v-7.143a3.947 3.947 0 0 1-5-.8a3.957 3.957 0 0 1-1.789 1.154Z" opacity=".25"/><path fill="currentColor" d="M22 23a.997.997 0 0 1-.707-.293l-20-20a1 1 0 0 1 1.414-1.414l20 20A1 1 0 0 1 22 23Z"/><path fill="currentColor" d="M12.586 14H10a1 1 0 0 0-1 1v7h6v-5.586Z" opacity=".7"/><path fill="currentColor" d="M10 14h2.586l-2.49-2.49A3.84 3.84 0 0 1 9 10.642a3.998 3.998 0 0 1-5 .82V21a1 1 0 0 0 1 1h4v-7a1 1 0 0 1 1-1zm5 2.414V22h4a.993.993 0 0 0 .93-.655z" opacity=".25"/><path fill="currentColor" d="M13.211 11.797A3.957 3.957 0 0 0 15 10.643A3.998 3.998 0 0 0 22 8a1.006 1.006 0 0 0-.071-.371l-2-5A1 1 0 0 0 19 2H5a1 1 0 0 0-.929.629l-.008.02zM3.255 4.669L2.071 7.63A1.006 1.006 0 0 0 2 8a3.998 3.998 0 0 0 7 2.643a3.84 3.84 0 0 0 1.095.866z" opacity=".5"/>`
	subjectInnerSVG              = `<path fill="currentColor" d="M21 8H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-8 10H3a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2zm8-5H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2z"/>`
	swiggyInnerSVG               = `<path fill="currentColor" d="M10.047 14.693c-.654 0-1.308-.012-1.961.004a1.342 1.342 0 0 1-1.33-.748A11.188 11.188 0 0 1 5.23 8.432a6.272 6.272 0 0 1 1.847-4.285a6.718 6.718 0 0 1 2.694-1.752a6.608 6.608 0 0 1 3.648-.243a6.831 6.831 0 0 1 5.343 5.711a.507.507 0 0 1-.344.603a4.95 4.95 0 0 1-1.573.374a14.831 14.831 0 0 1-3.148.045c-.42-.052-.496-.135-.5-.569c-.004-.592-.001-1.184-.002-1.777l-.003-.276a.297.297 0 0 0-.314-.335c-.235-.008-.339.125-.34.34a767.16 767.16 0 0 0-.001 3.034c0 .288.182.358.427.358c.909-.001 1.818-.004 2.727.002a6.945 6.945 0 0 1 1.67.162a1.203 1.203 0 0 1 1.012 1.635a14.17 14.17 0 0 1-1.494 3.871a34.228 34.228 0 0 1-3.577 5.26c-.345.43-.711.842-1.06 1.268c-.158.193-.277.189-.43-.01a43.866 43.866 0 0 1-3.427-4.956c-.127-.22-.226-.457-.333-.688c-.103-.222-.02-.368.193-.467a2.596 2.596 0 0 1 .8-.195a7.704 7.704 0 0 1 2.348.02c.45.082.524.188.523.639c0 .347-.005.694 0 1.042c.003.22.066.431.33.43c.262-.003.324-.215.326-.435c.004-.725-.002-1.45.003-2.176c.002-.312-.195-.371-.445-.372q-1.041-.002-2.083 0Z"/>`
	syncExclamationInnerSVG      = `<path fill="currentColor" d="M12 13a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v3a1 1 0 0 1-1 1zm0 3a.99.99 0 0 1-1-1a1.05 1.05 0 0 1 .29-.71a1.027 1.027 0 0 1 .33-.21a1.002 1.002 0 0 1 1.09.21A1.052 1.052 0 0 1 13 15a.99.99 0 0 1-1 1z"/><path fill="currentColor" d="M12 2a10.017 10.017 0 0 0-7 2.877V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2H6.218A7.98 7.98 0 0 1 20 12a1 1 0 0 0 2 0A10.011 10.011 0 0 0 12 2zm8 13.5h-4.5a1 1 0 0 0 0 2h2.282A7.98 7.98 0 0 1 4 12a1 1 0 0 0-2 0a9.987 9.987 0 0 0 17 7.123V21a1 1 0 0 0 2 0v-4.5a1 1 0 0 0-1-1z" opacity=".5"/>`
	syncSlashInnerSVG            = `<path fill="currentColor" d="M12 2C9.4 2 6.9 3 5 4.9V3c0-.6-.4-1-1-1s-1 .4-1 1v4.5c0 .6.4 1 1 1h4.5c.6 0 1-.4 1-1s-.4-1-1-1H6.2c3-3.2 8.1-3.3 11.3-.3C19.1 7.7 20 9.8 20 12c0 .6.4 1 1 1s1-.4 1-1c0-5.5-4.5-10-10-10zm8 13.5h-4.5c-.6 0-1 .4-1 1s.4 1 1 1h2.3c-3 3.2-8.1 3.3-11.3.3C4.9 16.3 4 14.2 4 12c0-.6-.4-1-1-1s-1 .4-1 1c0 5.5 4.5 10 10 10c2.6 0 5.2-1 7-2.9V21c0 .6.4 1 1 1s1-.4 1-1v-4.5c0-.6-.4-1-1-1z" opacity=".5"/><path fill="currentColor" d="M3 22c-.6 0-1-.4-1-1c0-.3.1-.5.3-.7l18-18c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4l-18 18c-.2.2-.4.3-.7.3z"/>`
	tableInnerSVG                = `<path fill="currentColor" d="M21 22H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1ZM4 20h16V4H4Z"/><path fill="currentColor" d="M9 22a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v18a1 1 0 0 1-1 1zm6 0a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v18a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M21 10H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 6H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2z"/>`
	telegramInnerSVG             = `<path fill="currentColor" d="M15.174 17.152a.705.705 0 0 1-1.002.352l-2.715-2.11l-1.742 1.608a.3.3 0 0 1-.285.039l.334-2.989l.01.009l.007-.059s4.885-4.448 5.084-4.637c.202-.189.135-.23.135-.23c.012-.23-.361 0-.361 0l-6.473 4.164l-2.695-.918s-.414-.148-.453-.475c-.041-.324.466-.5.466-.5l10.717-4.258s.881-.392.881.258l-1.908 9.746z" opacity=".5"/><path fill="currentColor" d="M11.994 2c-5.523 0-10 4.477-10 10s4.477 10 10 10s10-4.477 10-10s-4.477-10-10-10zm3.18 15.152a.705.705 0 0 1-1.002.352l-2.715-2.11l-1.742 1.608a.3.3 0 0 1-.285.039l.334-2.989l.01.009l.007-.059s4.885-4.448 5.084-4.637c.202-.189.135-.23.135-.23c.012-.23-.361 0-.361 0l-6.473 4.164l-2.695-.918s-.414-.149-.453-.475c-.041-.324.466-.5.466-.5l10.717-4.258s.881-.392.881.258l-1.908 9.746z"/>`
	telegramAltInnerSVG          = `<path fill="currentColor" d="M17.688 21.744a2.02 2.02 0 0 1-1.242-.427l-4.03-3.122l-2.702 2.983a1 1 0 0 1-1.698-.383l-2.02-6.682l-3.626-1.26a2.042 2.042 0 0 1-.103-3.818L20.187 1.8a2.042 2.042 0 0 1 2.771 2.295L19.695 20.11a2.054 2.054 0 0 1-2.008 1.633Z" opacity=".5"/><path fill="currentColor" d="M8.973 21.506a1 1 0 0 1-.957-.71l-2.168-7.16a.999.999 0 0 1 .495-1.176L16.91 6.958a1 1 0 0 1 1.17 1.594l-7.084 7.083l-1.044 5.072a1 1 0 0 1-.933.798h-.046Z"/>`
	thLargeInnerSVG              = `<path fill="currentColor" d="M20 3H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1Zm-9 16H5v-6h6Zm0-8H5V5h6Zm8 8h-6v-6h6Zm0-8h-6V5h6Z"/>`
	timesCircleInnerSVG          = `<path fill="currentColor" d="m13.414 12l3.293-3.293a1 1 0 0 0-1.414-1.414L12 10.586L8.707 7.293a1 1 0 0 0-1.414 1.414L10.586 12l-3.293 3.293a1 1 0 0 0 1.414 1.414L12 13.414l3.293 3.293a1 1 0 0 0 1.414-1.414Z"/><path fill="currentColor" d="M19.07 4.93A10 10 0 1 0 4.93 19.07A10 10 0 1 0 19.07 4.93Zm-2.363 10.363a1 1 0 1 1-1.414 1.414L12 13.414l-3.293 3.293a1 1 0 0 1-1.414-1.414L10.586 12L7.293 8.707a1 1 0 0 1 1.414-1.414L12 10.586l3.293-3.293a1 1 0 0 1 1.414 1.414L13.414 12Z" opacity=".5"/>`
	toggleOffInnerSVG            = `<path fill="currentColor" d="M16.5 17.5h-9a5.5 5.5 0 1 1 0-11h9a5.5 5.5 0 1 1 0 11z" opacity=".5"/><circle cx="7.5" cy="12" r="2.5" fill="currentColor"/>`
	toggleOnInnerSVG             = `<circle cx="16.5" cy="12" r="2.5" fill="currentColor" opacity=".5"/><path fill="currentColor" d="M16.5 6.5h-9a5.5 5.5 0 0 0 0 11h9a5.5 5.5 0 0 0 0-11zm0 8a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z"/>`
	toiletPaperInnerSVG          = `<ellipse cx="7" cy="7.993" fill="currentColor" rx="1" ry="1.25"/><path fill="currentColor" d="M7 2C4.243 2 2 4.691 2 8s2.243 6 5 6s5-2.691 5-6s-2.243-6-5-6Zm0 7.243a1.146 1.146 0 0 1-1-1.25a1.146 1.146 0 0 1 1-1.25a1.146 1.146 0 0 1 1 1.25a1.146 1.146 0 0 1-1 1.25Z" opacity=".5"/><path fill="currentColor" d="M22.76 20.35A7.504 7.504 0 0 1 21 15.459V8c0-3.309-2.243-6-5-6H7c2.757 0 5 2.691 5 6v7.459a9.507 9.507 0 0 0 2.24 6.191A1.001 1.001 0 0 0 15 22h7a1 1 0 0 0 .76-1.65Z" opacity=".25"/><path fill="currentColor" d="M12 8c0 3.309-2.243 6-5 6h5Z"/>`
	triangleInnerSVG             = `<path fill="currentColor" d="M21 20.794H3a1 1 0 0 1-.866-1.5l9-15.588a1.04 1.04 0 0 1 1.732 0l9 15.588a1 1 0 0 1-.866 1.5Z"/>`
	tumblrInnerSVG               = `<path fill="currentColor" d="M16.785 17.974a4.729 4.729 0 0 1-1.614.346a1.755 1.755 0 0 1-1.925-1.972v-6.226h4.017V7.095H13.26V2h-2.93a.157.157 0 0 0-.143.15a6.12 6.12 0 0 1-3.935 5.389v2.583h2.024v6.536c0 2.236 1.65 5.415 6.007 5.34a5.337 5.337 0 0 0 3.463-1.17l-.962-2.854"/>`
	tumblrAltInnerSVG            = `<path fill="currentColor" d="M16.785 17.974a4.729 4.729 0 0 1-1.614.346a1.755 1.755 0 0 1-1.925-1.972V11h4.017V7.095H13.26V2h-2.93a.157.157 0 0 0-.143.15a6.12 6.12 0 0 1-3.935 5.389v3.583h2.024v5.536c0 2.236 1.65 5.415 6.007 5.34a5.337 5.337 0 0 0 3.463-1.17l-.962-2.854"/>`
	tumblrSquareInnerSVG         = `<path fill="currentColor" d="M13.571 17.85a3.396 3.396 0 0 1-3.747-3.332v-4.076H8.562v-1.61a3.817 3.817 0 0 0 2.454-3.363a.099.099 0 0 1 .09-.093h1.827v3.178h2.496v1.888h-2.505v3.884a1.094 1.094 0 0 0 1.2 1.229a2.936 2.936 0 0 0 1.007-.215l.6 1.779c-.611.49-1.376.75-2.16.73z" opacity=".5"/><path fill="currentColor" d="M2.019 2v20h19.963V2H2.019zm11.552 15.85a3.396 3.396 0 0 1-3.747-3.332v-4.076H8.562v-1.61a3.817 3.817 0 0 0 2.454-3.363a.099.099 0 0 1 .09-.093h1.827v3.178h2.496v1.888h-2.505v3.884a1.094 1.094 0 0 0 1.2 1.229a2.936 2.936 0 0 0 1.007-.215l.6 1.779c-.611.49-1.376.75-2.16.73z"/>`
	twitterInnerSVG              = `<path fill="currentColor" d="M22 5.796a8.192 8.192 0 0 1-2.357.646a4.115 4.115 0 0 0 1.805-2.27a8.197 8.197 0 0 1-2.606.996a4.105 4.105 0 0 0-7.097 2.808a4.15 4.15 0 0 0 .105.935a11.65 11.65 0 0 1-8.456-4.287a4.107 4.107 0 0 0 1.27 5.478a4.084 4.084 0 0 1-1.86-.513v.052a4.105 4.105 0 0 0 3.292 4.023a4.082 4.082 0 0 1-1.081.143a4.165 4.165 0 0 1-.773-.072a4.108 4.108 0 0 0 3.832 2.85A8.261 8.261 0 0 1 2 18.282a11.611 11.611 0 0 0 6.29 1.844A11.594 11.594 0 0 0 19.962 8.453q0-.267-.013-.53A8.36 8.36 0 0 0 22 5.796Z"/>`
	twitterAltInnerSVG           = `<path fill="currentColor" d="M20.475 7.805c.01.185.01.37.01.545a11.885 11.885 0 0 1-.493 3.362A11.729 11.729 0 0 1 8.485 20.36a11.9 11.9 0 0 1-6.467-1.902a8.152 8.152 0 0 0 1.007.061a8.45 8.45 0 0 0 5.234-1.81a4.223 4.223 0 0 1-3.938-2.92a5.04 5.04 0 0 0 .792.072a4.04 4.04 0 0 0 1.12-.154a4.2 4.2 0 0 1-3.372-3.815a1.66 1.66 0 0 1-.02-.319v-.051a4.167 4.167 0 0 0 1.912.524A4.202 4.202 0 0 1 2.88 6.54a4.273 4.273 0 0 1 .566-2.129a11.968 11.968 0 0 0 8.7 4.412a4.219 4.219 0 0 1 7.187-3.846a8.443 8.443 0 0 0 2.684-1.028a4.978 4.978 0 0 1-1.543 3.856Z"/>`
	unlockInnerSVG               = `<path fill="currentColor" d="M7 9h10a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-7a3 3 0 0 1 3-3z"/><path fill="currentColor" d="M9 7a3.002 3.002 0 0 1 5.116-2.13c.378.383.65.858.786 1.379l.002.007a1 1 0 0 0 1.934-.505a5.09 5.09 0 0 0-1.306-2.293A5.002 5.002 0 0 0 7 7v2h2V7z" opacity=".5"/>`
	unlockAltInnerSVG            = `<path fill="currentColor" d="M8 11a1 1 0 0 1-1-1V7a5.002 5.002 0 0 1 8.532-3.542a5.09 5.09 0 0 1 1.306 2.293a1 1 0 0 1-1.934.505l-.002-.007a3.072 3.072 0 0 0-.786-1.379A3.002 3.002 0 0 0 9 7v3a1 1 0 0 1-1 1zm4 7a1 1 0 0 1-1-1v-3a1 1 0 1 1 2 0v3a1 1 0 0 1-1 1z" opacity=".5"/><path fill="currentColor" d="M17 9H7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm-4 8a1 1 0 0 1-2 0v-3a1 1 0 1 1 2 0v3z"/>`
	uploadAltInnerSVG            = `<path fill="currentColor" d="m15.707 5.293l-3-3a1 1 0 0 0-1.414 0l-3 3a1 1 0 0 0 1.414 1.414L11 5.414V17a1 1 0 0 0 2 0V5.414l1.293 1.293a1 1 0 0 0 1.414-1.414Z"/><path fill="currentColor" d="M18 9h-5v8a1 1 0 0 1-2 0V9H6a3.003 3.003 0 0 0-3 3v7a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3v-7a3.003 3.003 0 0 0-3-3Z" opacity=".5"/>`
	userArrowsInnerSVG           = `<path fill="currentColor" d="M6 16.5A3.5 3.5 0 1 1 9.5 13A3.504 3.504 0 0 1 6 16.5Z" opacity=".25"/><path fill="currentColor" d="M8.64 15.272a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 1 19.5a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228zM18 16.5a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5z" opacity=".5"/><path fill="currentColor" d="M20.64 15.272a3.452 3.452 0 0 1-5.28 0A4.988 4.988 0 0 0 13 19.5a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a4.988 4.988 0 0 0-2.36-4.228Z" opacity=".7"/><path fill="currentColor" d="m17.207 5.793l-2-2a1 1 0 0 0-1.414 1.414l.293.293H9.914l.293-.293a1 1 0 0 0-1.414-1.414l-2 2a1 1 0 0 0 0 1.414l2 2a1 1 0 0 0 1.414-1.414L9.914 7.5h4.172l-.293.293a1 1 0 1 0 1.414 1.414l2-2a1 1 0 0 0 0-1.414Z"/>`
	userMdInnerSVG               = `<path fill="currentColor" d="M17.998 8.064L6.003 8.11l-.277-3.325A3 3 0 0 1 8.17 1.482l.789-.143a17.031 17.031 0 0 1 6.086 0l.786.143a3 3 0 0 1 2.443 3.302Z"/><path fill="currentColor" d="M6.009 8.109a5.994 5.994 0 0 0 11.984-.045Z" opacity=".25"/><path fill="currentColor" d="m17.198 13.385l-4.49 4.49a1 1 0 0 1-1.415 0l-4.491-4.49a9.945 9.945 0 0 0-4.736 7.44a1 1 0 0 0 .994 1.108h17.88a1 1 0 0 0 .994-1.108a9.945 9.945 0 0 0-4.736-7.44Z"/><path fill="currentColor" d="M15.69 12.654a6.012 6.012 0 0 1-7.381 0a10.004 10.004 0 0 0-1.507.73l4.491 4.492a1 1 0 0 0 1.414 0l4.491-4.491a10.005 10.005 0 0 0-1.507-.731Z" opacity=".5"/>`
	userNurseInnerSVG            = `<path fill="currentColor" d="M20.94 22H3.06a1 1 0 0 1-.994-1.108a9.995 9.995 0 0 1 19.868 0A1 1 0 0 1 20.94 22Z" opacity=".5"/><path fill="currentColor" d="m12.708 18.307l4.706-4.715a10.001 10.001 0 0 0-10.833.003l4.712 4.712A1 1 0 0 0 12 18.6a1.002 1.002 0 0 0 .708-.293Z" opacity=".25"/><path fill="currentColor" d="M11.995 14a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6Z" opacity=".25"/><path fill="currentColor" d="M6.09 9a5.993 5.993 0 0 0 11.82 0Z"/>`
	vectorSquareInnerSVG         = `<path fill="currentColor" d="M5 8a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1zm14 4a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1zM5 22a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1zm14 4a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1z"/><path fill="currentColor" d="M16.184 20a2.805 2.805 0 0 1 0-2H7.816a2.806 2.806 0 0 1 0 2zM19 8a2.965 2.965 0 0 1-1-.184v8.368a2.806 2.806 0 0 1 2 0V7.816A2.965 2.965 0 0 1 19 8zM7.816 4A2.965 2.965 0 0 1 8 5a2.965 2.965 0 0 1-.184 1h8.368A2.965 2.965 0 0 1 16 5a2.965 2.965 0 0 1 .184-1zM5 16a2.965 2.965 0 0 1 1 .184V7.816A2.965 2.965 0 0 1 5 8a2.965 2.965 0 0 1-1-.184v8.368A2.965 2.965 0 0 1 5 16z" opacity=".5"/>`
	vectorSquareAltInnerSVG      = `<path fill="currentColor" d="M4 22a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm0-2.002zM4 6a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm0-2.002zM20 6a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm0-2.002zM20 22a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm0-2.002z"/><rect width="10" height="10" x="7" y="7" fill="currentColor" opacity=".5" rx="1"/><path fill="currentColor" d="M18.278 5a1.936 1.936 0 0 1 0-2H5.722a1.936 1.936 0 0 1 0 2zM20 18a1.976 1.976 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v12.556A1.976 1.976 0 0 1 20 18zM4 18a1.976 1.976 0 0 1 1 .278V5.722a1.936 1.936 0 0 1-2 0v12.556A1.976 1.976 0 0 1 4 18zm14.278 1H5.722a1.936 1.936 0 0 1 0 2h12.556a1.936 1.936 0 0 1 0-2z" opacity=".25"/>`
	virusSlashInnerSVG           = `<circle cx="9" cy="15" r="1" fill="currentColor"/><circle cx="15" cy="9" r="1" fill="currentColor"/><path fill="currentColor" d="M12 8a1 1 0 0 1-1-1V2a1 1 0 0 1 2 0v5a1 1 0 0 1-1 1zm0 15a1 1 0 0 1-1-1v-5a1 1 0 0 1 2 0v5a1 1 0 0 1-1 1zm10-10h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2z" opacity=".25"/><path fill="currentColor" d="m4.98 17.607l-.758.757a1 1 0 1 0 1.414 1.414l.758-.757a9.06 9.06 0 0 1-1.415-1.414zM19.02 6.394l.758-.758a1 1 0 0 0-1.414-1.414l-.757.757a9.06 9.06 0 0 1 1.414 1.415z"/><path fill="currentColor" d="M6 13H2a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Z" opacity=".25"/><path fill="currentColor" d="M22 23.002a.997.997 0 0 1-.707-.293l-20-20a1 1 0 0 1 1.414-1.414l20 20A1 1 0 0 1 22 23.002Z"/><path fill="currentColor" d="M15.746 14.334a1.5 1.5 0 0 0-2.08-2.08zM8.252 9.668a1.5 1.5 0 0 0 2.08 2.08z" opacity=".7"/><path fill="currentColor" d="m17.606 19.022l-7.274-7.274a1.5 1.5 0 0 1-2.08-2.08L4.978 6.394A8.943 8.943 0 0 0 3.06 11H6a1 1 0 0 1 0 2H3.059A9.012 9.012 0 0 0 11 20.94V17a1 1 0 0 1 2 0v3.941a8.949 8.949 0 0 0 4.606-1.919zM9 16a1 1 0 1 1 1-1a1 1 0 0 1-1 1zM6.392 4.98l7.274 7.274a1.5 1.5 0 0 1 2.08 2.08l3.275 3.274A8.949 8.949 0 0 0 20.94 13H18a1 1 0 0 1 0-2h2.94A9.012 9.012 0 0 0 13 3.059V7a1 1 0 0 1-2 0V3.06a8.953 8.953 0 0 0-4.608 1.92zM15 8a1 1 0 1 1-1 1a1 1 0 0 1 1-1z" opacity=".5"/>`
	visualStudioInnerSVG         = `<path fill="currentColor" d="M9.133 14.242L4.2 18.087l-2.152-1.072V6.994l2.144-1.08l4.9 3.853L16.849 2l5.104 2.033v15.9L16.875 22Zm7.563 1.352V8.406l-4.645 3.598l4.645 3.59ZM4.369 14.301l2.442-2.22l-2.442-2.433Z"/>`
	vkInnerSVG                   = `<path fill="currentColor" d="M18.146 16.27h-1.459c-.552 0-.718-.447-1.708-1.437c-.864-.833-1.229-.937-1.448-.937c-.302 0-.385.083-.385.5v1.312c0 .355-.115.563-1.042.563a5.692 5.692 0 0 1-4.448-2.667a11.626 11.626 0 0 1-2.302-4.833c0-.219.083-.417.5-.417h1.458c.375 0 .51.167.657.552c.708 2.084 1.916 3.896 2.406 3.896c.187 0 .27-.083.27-.552v-2.146c-.062-.979-.583-1.062-.583-1.416a.36.36 0 0 1 .375-.334h2.292c.313 0 .417.156.417.531v2.896c0 .313.135.417.229.417c.187 0 .333-.104.677-.448c.739-.9 1.342-1.904 1.792-2.98a.628.628 0 0 1 .635-.416h1.458c.438 0 .532.219.438.531a18.205 18.205 0 0 1-1.958 3.365c-.157.24-.22.365 0 .646c.145.219.656.646 1 1.052c.5.499.915 1.076 1.229 1.708c.125.406-.084.615-.5.615z" opacity=".5"/><path fill="currentColor" d="M15.073 2H8.937C3.333 2 2 3.333 2 8.927v6.136C2 20.666 3.323 22 8.927 22h6.136C20.666 22 22 20.677 22 15.073V8.937C22 3.333 20.677 2 15.073 2zm3.073 14.27h-1.459c-.552 0-.718-.447-1.708-1.437c-.864-.833-1.229-.937-1.448-.937c-.302 0-.385.083-.385.5v1.312c0 .355-.115.563-1.042.563a5.692 5.692 0 0 1-4.448-2.667a11.626 11.626 0 0 1-2.302-4.833c0-.219.083-.417.5-.417h1.458c.375 0 .51.167.657.552c.708 2.084 1.916 3.896 2.406 3.896c.188 0 .27-.083.27-.552v-2.146c-.062-.979-.582-1.062-.582-1.416a.36.36 0 0 1 .374-.334h2.292c.313 0 .417.156.417.531v2.896c0 .313.135.417.229.417c.188 0 .333-.104.677-.448a12 12 0 0 0 1.792-2.98a.628.628 0 0 1 .635-.416h1.459c.437 0 .53.219.437.531a18.205 18.205 0 0 1-1.958 3.365c-.157.24-.22.365 0 .646c.146.219.656.646 1 1.052c.5.499.915 1.076 1.229 1.708c.125.406-.084.615-.5.615z"/>`
	vkAltInnerSVG                = `<path fill="currentColor" d="M15.063 23H8.927C2.777 23 1 21.22 1 15.062V8.928C1 2.778 2.78 1 8.938 1h6.135C21.223 1 23 2.78 23 8.938v6.135C23 21.222 21.22 23 15.062 23z" opacity=".5"/><path fill="currentColor" d="M11.654 17.002a.998.998 0 0 1-.586-.19a13.334 13.334 0 0 1-5.4-8.643a1 1 0 0 1 1.972-.334a11.307 11.307 0 0 0 3.014 5.99V8.001a1 1 0 1 1 2 0v8a1 1 0 0 1-.999 1z"/><path fill="currentColor" d="M16.655 17.002a.998.998 0 0 1-.738-.324L14.27 14.88a3.552 3.552 0 0 0-1.615-1.005v2.126a1 1 0 1 1-2 0v-3.27a1 1 0 0 1 .999-1a5.562 5.562 0 0 1 4.09 1.797l1.647 1.797a1 1 0 0 1-.736 1.676z"/><path fill="currentColor" d="M11.65 14.96a1 1 0 0 1-1-1V8.04a1 1 0 1 1 2 0v5.92a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M12.57 13.83a1 1 0 0 1-.38-1.926a5.86 5.86 0 0 0 3.485-4.126a1 1 0 1 1 1.95.444a7.865 7.865 0 0 1-4.676 5.534a.996.996 0 0 1-.379.074zm-.916-4.828h-.003L10.997 9A1 1 0 0 1 10 7.997A1.03 1.03 0 0 1 11.003 7l.654.002a1 1 0 0 1-.003 2z"/>`
	vuejsInnerSVG                = `<path fill="currentColor" d="m6.976 3.433l3.646.002l1.384 1.956l1.374-1.956l3.643-.001l-5 8.406l-5.047-8.407z"/><path fill="currentColor" d="M14.6 2.43a1 1 0 0 0-.819.425L12 5.39l-1.791-2.537a1 1 0 0 0-.817-.423H6.38l3.55 5.92l2.1 3.5l2.07-3.5l3.52-5.92H14.6z"/><path fill="currentColor" d="m22.001 2.438l-4.384-.003L14.1 8.35l-2.07 3.5l-2.1-3.5l-3.546-5.913l-4.39.024a1 1 0 0 0-.857 1.506l10.02 17.105a1 1 0 0 0 1.727-.002l9.98-17.128a1 1 0 0 0-.863-1.504z" opacity=".5"/>`
	vuejsAltInnerSVG             = `<path fill="currentColor" d="M12.018 19.151L4.315 6h2.823l4.886 8.407L16.871 6h2.809z" opacity=".5"/><path fill="currentColor" d="m14.38 4.001l-2.374 3.956l-2.384-3.956H.825L12.02 23.115L23.161 4H14.38zm-2.362 15.15L4.315 6h2.823l4.886 8.407L16.871 6h2.809l-7.662 13.151z"/>`
	webGridInnerSVG              = `<path fill="currentColor" d="M2 13h12v9h2V2h-2v9H2z" opacity=".25"/><path fill="currentColor" d="M21 22h-5V2h5a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M14 22H3a1 1 0 0 1-1-1v-8h12v9zm0-11H2V3a1 1 0 0 1 1-1h11v9z" opacity=".5"/>`
	webGridAltInnerSVG           = `<path fill="currentColor" d="M22 8H2v2h9v12h2V10h9z" opacity=".25"/><path fill="currentColor" d="M3 2h18a1 1 0 0 1 1 1v5H2V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M2 10h9v12H3a1 1 0 0 1-1-1V10zm11 0h9v11a1 1 0 0 1-1 1h-8V10z" opacity=".5"/>`
	webSectionInnerSVG           = `<path fill="currentColor" d="M14 2h2v20h-2z" opacity=".25"/><path fill="currentColor" d="M3 2h11v20H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1z" opacity=".5"/><path fill="currentColor" d="M16 2h5a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1h-5V2z"/>`
	webSectionAltInnerSVG        = `<path fill="currentColor" d="M9.9 2H8.1A2.584 2.584 0 0 0 8 3v18c-.032.337.002.676.1 1h1.8c.098-.324.132-.663.1-1V3a2.584 2.584 0 0 0-.1-1z" opacity=".25"/><path fill="currentColor" d="M3 2h5v20H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M10 2h11a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H10V2z" opacity=".5"/>`
	whatsappInnerSVG             = `<path fill="currentColor" d="M21.99 6.547a10.59 10.59 0 0 0-.103-1.282a4.312 4.312 0 0 0-.363-1.09A3.853 3.853 0 0 0 19.83 2.48a4.299 4.299 0 0 0-1.083-.362a10.523 10.523 0 0 0-1.292-.105c-.183-.007-.42-.01-.53-.01L7.077 2c-.11 0-.347.003-.53.01a10.565 10.565 0 0 0-1.282.103a4.312 4.312 0 0 0-1.09.363A3.854 3.854 0 0 0 2.48 4.17a4.303 4.303 0 0 0-.362 1.083a10.545 10.545 0 0 0-.106 1.292c-.006.183-.01.42-.01.53L2 16.923c0 .11.003.347.01.53a10.565 10.565 0 0 0 .103 1.282a4.313 4.313 0 0 0 .363 1.09A3.854 3.854 0 0 0 4.17 21.52a4.305 4.305 0 0 0 1.083.362a10.52 10.52 0 0 0 1.292.105c.183.007.42.01.53.01l9.848.002c.11 0 .347-.003.53-.01a10.578 10.578 0 0 0 1.282-.103a4.316 4.316 0 0 0 1.09-.363a3.854 3.854 0 0 0 1.696-1.694a4.301 4.301 0 0 0 .362-1.083a10.533 10.533 0 0 0 .106-1.292c.006-.183.01-.42.01-.53L22 7.077c0-.11-.003-.347-.01-.53Zm-9.773 12.41h-.003a7.126 7.126 0 0 1-3.407-.868l-3.78.991l1.012-3.693a7.13 7.13 0 1 1 6.178 3.57Z"/><path fill="currentColor" d="M12.22 5.901a5.927 5.927 0 0 0-5.023 9.076l.141.224l-.599 2.186l2.243-.588l.216.128a5.918 5.918 0 0 0 3.016.826h.003A5.926 5.926 0 0 0 12.219 5.9Zm3.484 8.47a1.834 1.834 0 0 1-1.202.847a2.443 2.443 0 0 1-1.122-.07a10.276 10.276 0 0 1-1.015-.376a7.94 7.94 0 0 1-3.043-2.689a3.463 3.463 0 0 1-.728-1.842a1.997 1.997 0 0 1 .624-1.485a.655.655 0 0 1 .475-.223c.118 0 .237 0 .341.006c.11.005.256-.042.4.306c.15.356.506 1.233.55 1.322a.328.328 0 0 1 .015.312a1.216 1.216 0 0 1-.178.297c-.09.104-.187.232-.267.312c-.09.089-.182.185-.079.363a5.366 5.366 0 0 0 .991 1.234a4.863 4.863 0 0 0 1.433.884c.178.09.282.074.386-.045s.445-.52.564-.698s.237-.148.4-.089s1.04.49 1.218.58s.297.133.341.207a1.488 1.488 0 0 1-.104.847Z"/>`
	windowGridInnerSVG           = `<path fill="currentColor" d="M22 11H10V2H8v20h2v-9h12z" opacity=".25"/><path fill="currentColor" d="M3 2h5v20H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M10 2h11a1 1 0 0 1 1 1v8H10V2zm0 11h12v8a1 1 0 0 1-1 1H10v-9z" opacity=".5"/>`
	windowMaximizeInnerSVG       = `<path fill="currentColor" d="M2 8h20v2H2z" opacity=".25"/><path fill="currentColor" d="M3 2h18a1 1 0 0 1 1 1v5H2V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M2 10h20v11a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V10z" opacity=".5"/>`
	windowSectionInnerSVG        = `<path fill="currentColor" d="M9 10h6v12H9zm-7 0v11a1 1 0 0 0 1 1h4V10H2z" opacity=".5"/><path fill="currentColor" d="M22 8H2v2h5v12h2V10h6v12h2V10h5z" opacity=".25"/><path fill="currentColor" d="M17 10v12h4a1 1 0 0 0 1-1V10h-5z" opacity=".5"/><path fill="currentColor" d="M3 2h18a1 1 0 0 1 1 1v5H2V3a1 1 0 0 1 1-1z"/>`
	windowsInnerSVG              = `<path fill="currentColor" d="m2.03 4.832l8.147-1.11l.004 7.86l-8.144.046l-.008-6.796Zm8.144 7.655l.006 7.867l-8.144-1.12l-.001-6.8l8.138.053Zm.987-8.91L21.965 2v9.482l-10.804.085v-7.99Zm10.807 8.984L21.965 22l-10.804-1.525l-.015-7.932Z"/>`
	wordpressInnerSVG            = `<path fill="currentColor" d="M12 2.6a9.4 9.4 0 1 0 9.4 9.4A9.4 9.4 0 0 0 12 2.6Z" opacity=".25"/><path fill="currentColor" d="m12.146 12.729l-2.5 7.265a8.337 8.337 0 0 0 5.121-.133a.744.744 0 0 1-.06-.115zm-8.48-.73a8.334 8.334 0 0 0 4.698 7.5L4.388 8.61A8.3 8.3 0 0 0 3.667 12z"/><path fill="currentColor" d="M17.626 11.58a4.389 4.389 0 0 0-.687-2.299a3.903 3.903 0 0 1-.819-1.954a1.443 1.443 0 0 1 1.4-1.48c.037 0 .072.005.107.007a8.331 8.331 0 0 0-12.59 1.569c.196.006.38.01.537.01c.871 0 2.22-.106 2.22-.106a.344.344 0 0 1 .054.686s-.452.053-.954.08l3.035 9.026l1.824-5.47l-1.299-3.556c-.449-.027-.874-.08-.874-.08a.344.344 0 0 1 .053-.686s1.376.106 2.195.106c.871 0 2.221-.106 2.221-.106a.344.344 0 0 1 .053.686s-.452.053-.953.08l3.011 8.957l.86-2.724a9.604 9.604 0 0 0 .606-2.747Z"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm7.795 15.255a9.427 9.427 0 0 1-4.136 3.406a9.388 9.388 0 0 1-8.914-.866a9.432 9.432 0 0 1-3.407-4.136a9.386 9.386 0 0 1 .867-8.914a9.427 9.427 0 0 1 4.136-3.407a9.388 9.388 0 0 1 8.914.867a9.432 9.432 0 0 1 3.407 4.136a9.386 9.386 0 0 1-.867 8.914Z"/><path fill="currentColor" d="M19.369 8.859a7.865 7.865 0 0 1-.634 2.985l-2.545 7.358a8.334 8.334 0 0 0 3.123-11.2a6.381 6.381 0 0 1 .056.857Z"/>`
	wordpressSimpleInnerSVG      = `<circle cx="12" cy="12" r="10" fill="currentColor" opacity=".25"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2ZM3.01 12a8.955 8.955 0 0 1 .778-3.66l4.289 11.751A8.991 8.991 0 0 1 3.009 12ZM12 20.99a8.988 8.988 0 0 1-2.54-.366l2.698-7.839l2.763 7.571a.843.843 0 0 0 .065.124a8.971 8.971 0 0 1-2.986.51Zm1.239-13.207c.541-.028 1.03-.085 1.03-.085a.372.372 0 0 0-.058-.741s-1.457.114-2.397.114c-.883 0-2.368-.114-2.368-.114a.372.372 0 0 0-.057.74s.459.058.943.086l1.401 3.838l-1.968 5.901l-3.274-9.739a18.252 18.252 0 0 0 1.03-.085a.372.372 0 0 0-.058-.741s-1.456.114-2.396.114c-.17 0-.368-.004-.579-.01a8.988 8.988 0 0 1 13.583-1.693c-.039-.002-.076-.007-.116-.007a1.557 1.557 0 0 0-1.51 1.596a4.21 4.21 0 0 0 .883 2.109a4.736 4.736 0 0 1 .741 2.48a10.883 10.883 0 0 1-.684 2.906l-.897 2.996ZM16.52 19.77l2.746-7.94a8.489 8.489 0 0 0 .684-3.22a6.91 6.91 0 0 0-.06-.925a8.992 8.992 0 0 1-3.37 12.085Z"/>`
	wrapTextInnerSVG             = `<path fill="currentColor" d="M21 7.167H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-12 10H3a1 1 0 1 1 0-2h6a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M18.5 17.167H15a1 1 0 0 1 0-2h3.5a1.5 1.5 0 0 0 0-3H3a1 1 0 0 1 0-2h15.5a3.5 3.5 0 1 1 0 7Z"/><path fill="currentColor" d="M14.999 18.833a.995.995 0 0 1-.639-.231l-2-1.666a1 1 0 0 1 0-1.538l2-1.667a1 1 0 0 1 1.28 1.538l-1.078.898l1.078.897a1 1 0 0 1-.641 1.769Z"/>`
	youtubeInnerSVG              = `<path fill="currentColor" d="M15.663 11.775c-2.017-1.078-3.948-2.078-5.922-3.112v6.19c2.077-1.13 4.267-2.164 5.931-3.087l-.009.009z" opacity=".5"/><path fill="currentColor" d="M22.974 9.715a8.539 8.539 0 0 0-.914-4.13a2.868 2.868 0 0 0-1.715-1.017a81.07 81.07 0 0 0-8.337-.293a80.928 80.928 0 0 0-8.336.293c-.55.102-1.058.36-1.466.741c-.895.83-.96 2.251-1.095 3.449a47.52 47.52 0 0 0 0 6.482c.025.676.126 1.347.302 2c.12.506.362.974.707 1.363c.408.402.928.672 1.491.775a44.69 44.69 0 0 0 6.5.328c3.495.056 6.573-.003 10.199-.276a2.89 2.89 0 0 0 1.534-.784c.273-.272.48-.602.604-.966c.362-1.095.54-2.243.526-3.396c.034-.595.034-3.974 0-4.57zm-7.31 2.06l-.004-.002c-.83.46-1.788.946-2.798 1.459a125.297 125.297 0 0 0-3.121 1.62v-6.19c.987.518 1.963 1.027 2.945 1.54c.98.514 1.967 1.033 2.974 1.57l.012-.006l-.008.009z"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Adobe(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adobeInnerSVG).
		Element(children...)
}

func AdobeAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adobeAltInnerSVG).
		Element(children...)
}

func Airplay(children ...engine.UberChild) *engine.UberElement {
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

func Align(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignInnerSVG).
		Element(children...)
}

func AlignAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignAltInnerSVG).
		Element(children...)
}

func AlignCenter(children ...engine.UberChild) *engine.UberElement {
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

func AlignCenterJustify(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterJustifyInnerSVG).
		Element(children...)
}

func AlignJustify(children ...engine.UberChild) *engine.UberElement {
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

func AlignLeft(children ...engine.UberChild) *engine.UberElement {
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

func AlignLeftJustify(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignLeftJustifyInnerSVG).
		Element(children...)
}

func AlignLetterRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignLetterRightInnerSVG).
		Element(children...)
}

func AlignRight(children ...engine.UberChild) *engine.UberElement {
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

func AlignRightJustify(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignRightJustifyInnerSVG).
		Element(children...)
}

func Amazon(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(amazonInnerSVG).
		Element(children...)
}

func Analysis(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(analysisInnerSVG).
		Element(children...)
}

func Analytics(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(analyticsInnerSVG).
		Element(children...)
}

func Anchor(children ...engine.UberChild) *engine.UberElement {
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

func Android(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(androidInnerSVG).
		Element(children...)
}

func AndroidAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(androidAltInnerSVG).
		Element(children...)
}

func AngleDoubleDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDoubleDownInnerSVG).
		Element(children...)
}

func AngleDoubleLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDoubleLeftInnerSVG).
		Element(children...)
}

func AngleDoubleRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDoubleRightInnerSVG).
		Element(children...)
}

func AngleDoubleUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDoubleUpInnerSVG).
		Element(children...)
}

func AngleDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleDownInnerSVG).
		Element(children...)
}

func AngleLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleLeftInnerSVG).
		Element(children...)
}

func AngleRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleRightInnerSVG).
		Element(children...)
}

func AngleRightB(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleRightBInnerSVG).
		Element(children...)
}

func AngleUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(angleUpInnerSVG).
		Element(children...)
}

func Apple(children ...engine.UberChild) *engine.UberElement {
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

func AppleAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleAltInnerSVG).
		Element(children...)
}

func Apps(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appsInnerSVG).
		Element(children...)
}

func ArrowCircleDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleDownInnerSVG).
		Element(children...)
}

func ArrowCircleLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleLeftInnerSVG).
		Element(children...)
}

func ArrowCircleRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleRightInnerSVG).
		Element(children...)
}

func ArrowCircleUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowCircleUpInnerSVG).
		Element(children...)
}

func ArrowDownLeft(children ...engine.UberChild) *engine.UberElement {
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

func ArrowDownRight(children ...engine.UberChild) *engine.UberElement {
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

func ArrowUpLeft(children ...engine.UberChild) *engine.UberElement {
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

func ArrowUpRight(children ...engine.UberChild) *engine.UberElement {
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

func At(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atInnerSVG).
		Element(children...)
}

func Bag(children ...engine.UberChild) *engine.UberElement {
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

func BatteryBolt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryBoltInnerSVG).
		Element(children...)
}

func BatteryEmpty(children ...engine.UberChild) *engine.UberElement {
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

func Behance(children ...engine.UberChild) *engine.UberElement {
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

func BehanceAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(behanceAltInnerSVG).
		Element(children...)
}

func Bing(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bingInnerSVG).
		Element(children...)
}

func Bitcoin(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bitcoinInnerSVG).
		Element(children...)
}

func BitcoinAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bitcoinAltInnerSVG).
		Element(children...)
}

func Blackberry(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(blackberryInnerSVG).
		Element(children...)
}

func Blogger(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bloggerInnerSVG).
		Element(children...)
}

func BloggerAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bloggerAltInnerSVG).
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

func BorderAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderAltInnerSVG).
		Element(children...)
}

func BorderBottom(children ...engine.UberChild) *engine.UberElement {
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

func BorderClear(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderClearInnerSVG).
		Element(children...)
}

func BorderHorizontal(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderHorizontalInnerSVG).
		Element(children...)
}

func BorderInner(children ...engine.UberChild) *engine.UberElement {
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

func BorderLeft(children ...engine.UberChild) *engine.UberElement {
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

func BorderOut(children ...engine.UberChild) *engine.UberElement {
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

func BorderRight(children ...engine.UberChild) *engine.UberElement {
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

func BorderTop(children ...engine.UberChild) *engine.UberElement {
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

func BorderVertical(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderVerticalInnerSVG).
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

func Calender(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calenderInnerSVG).
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

func ChartPie(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartPieInnerSVG).
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

func CheckSquare(children ...engine.UberChild) *engine.UberElement {
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

func CircleLayer(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleLayerInnerSVG).
		Element(children...)
}

func ClinicMedical(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clinicMedicalInnerSVG).
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

func ClockEight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockEightInnerSVG).
		Element(children...)
}

func ClockFive(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockFiveInnerSVG).
		Element(children...)
}

func ClockNine(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockNineInnerSVG).
		Element(children...)
}

func ClockSeven(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockSevenInnerSVG).
		Element(children...)
}

func ClockTen(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockTenInnerSVG).
		Element(children...)
}

func ClockThree(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockThreeInnerSVG).
		Element(children...)
}

func ClockTwo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockTwoInnerSVG).
		Element(children...)
}

func Columns(children ...engine.UberChild) *engine.UberElement {
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

func Comment(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentInnerSVG).
		Element(children...)
}

func CommentAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentAltInnerSVG).
		Element(children...)
}

func CommentAltDots(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentAltDotsInnerSVG).
		Element(children...)
}

func CommentAltMessage(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentAltMessageInnerSVG).
		Element(children...)
}

func CommentAltPlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentAltPlusInnerSVG).
		Element(children...)
}

func CommentDots(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentDotsInnerSVG).
		Element(children...)
}

func CommentMessage(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentMessageInnerSVG).
		Element(children...)
}

func CommentPlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentPlusInnerSVG).
		Element(children...)
}

func Compress(children ...engine.UberChild) *engine.UberElement {
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

func CornerUpLeft(children ...engine.UberChild) *engine.UberElement {
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

func Coronavirus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coronavirusInnerSVG).
		Element(children...)
}

func CssThree(children ...engine.UberChild) *engine.UberElement {
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

func CssThreeSimple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cssThreeSimpleInnerSVG).
		Element(children...)
}

func Cube(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cubeInnerSVG).
		Element(children...)
}

func Dialpad(children ...engine.UberChild) *engine.UberElement {
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

func DialpadAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dialpadAltInnerSVG).
		Element(children...)
}

func Direction(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(directionInnerSVG).
		Element(children...)
}

func Discord(children ...engine.UberChild) *engine.UberElement {
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

func Docker(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dockerInnerSVG).
		Element(children...)
}

func DocumentLayoutCenter(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentLayoutCenterInnerSVG).
		Element(children...)
}

func DocumentLayoutLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentLayoutLeftInnerSVG).
		Element(children...)
}

func DocumentLayoutRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentLayoutRightInnerSVG).
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

func Dribbble(children ...engine.UberChild) *engine.UberElement {
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

func Dropbox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropboxInnerSVG).
		Element(children...)
}

func EllipsisH(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsisHInnerSVG).
		Element(children...)
}

func EllipsisV(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ellipsisVInnerSVG).
		Element(children...)
}

func Entry(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(entryInnerSVG).
		Element(children...)
}

func ExclamationCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationCircleInnerSVG).
		Element(children...)
}

func ExclamationOctagon(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationOctagonInnerSVG).
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

func Exit(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exitInnerSVG).
		Element(children...)
}

func Facebook(children ...engine.UberChild) *engine.UberElement {
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

func FacebookF(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookFInnerSVG).
		Element(children...)
}

func FacebookMessenger(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookMessengerInnerSVG).
		Element(children...)
}

func FacebookMessengerAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookMessengerAltInnerSVG).
		Element(children...)
}

func Favorite(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(favoriteInnerSVG).
		Element(children...)
}

func FiveHundredPx(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fiveHundredPxInnerSVG).
		Element(children...)
}

func FlipH(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipHInnerSVG).
		Element(children...)
}

func FlipHAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipHAltInnerSVG).
		Element(children...)
}

func FlipV(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipVInnerSVG).
		Element(children...)
}

func FlipVAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipVAltInnerSVG).
		Element(children...)
}

func Github(children ...engine.UberChild) *engine.UberElement {
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

func GithubAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(githubAltInnerSVG).
		Element(children...)
}

func Gitlab(children ...engine.UberChild) *engine.UberElement {
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

func GitlabAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitlabAltInnerSVG).
		Element(children...)
}

func Google(children ...engine.UberChild) *engine.UberElement {
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

func GoogleDrive(children ...engine.UberChild) *engine.UberElement {
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

func GoogleDriveAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleDriveAltInnerSVG).
		Element(children...)
}

func GoogleHangouts(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleHangoutsInnerSVG).
		Element(children...)
}

func GoogleHangoutsAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleHangoutsAltInnerSVG).
		Element(children...)
}

func GooglePlay(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googlePlayInnerSVG).
		Element(children...)
}

func GraphBar(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphBarInnerSVG).
		Element(children...)
}

func Grid(children ...engine.UberChild) *engine.UberElement {
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

func Grids(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridsInnerSVG).
		Element(children...)
}

func GripHorizontalLine(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gripHorizontalLineInnerSVG).
		Element(children...)
}

func HeadSide(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headSideInnerSVG).
		Element(children...)
}

func HeadSideCough(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headSideCoughInnerSVG).
		Element(children...)
}

func HeadSideMask(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headSideMaskInnerSVG).
		Element(children...)
}

func Hipchat(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hipchatInnerSVG).
		Element(children...)
}

func History(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(historyInnerSVG).
		Element(children...)
}

func HistoryAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(historyAltInnerSVG).
		Element(children...)
}

func HorizontalAlignLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(horizontalAlignLeftInnerSVG).
		Element(children...)
}

func Hospital(children ...engine.UberChild) *engine.UberElement {
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

func HospitalSquareSign(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hospitalSquareSignInnerSVG).
		Element(children...)
}

func HospitalSymbol(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hospitalSymbolInnerSVG).
		Element(children...)
}

func HouseUser(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(houseUserInnerSVG).
		Element(children...)
}

func HtmlFive(children ...engine.UberChild) *engine.UberElement {
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

func HtmlFiveAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(htmlFiveAltInnerSVG).
		Element(children...)
}

func HtmlThree(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(htmlThreeInnerSVG).
		Element(children...)
}

func HtmlThreeAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(htmlThreeAltInnerSVG).
		Element(children...)
}

func ImageV(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageVInnerSVG).
		Element(children...)
}

func Instagram(children ...engine.UberChild) *engine.UberElement {
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

func InstagramAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(instagramAltInnerSVG).
		Element(children...)
}

func Intercom(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(intercomInnerSVG).
		Element(children...)
}

func IntercomAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(intercomAltInnerSVG).
		Element(children...)
}

func JavaScript(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(javaScriptInnerSVG).
		Element(children...)
}

func KeySkeleton(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keySkeletonInnerSVG).
		Element(children...)
}

func KeySkeletonAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keySkeletonAltInnerSVG).
		Element(children...)
}

func KeyholeCircle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyholeCircleInnerSVG).
		Element(children...)
}

func KeyholeSquare(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyholeSquareInnerSVG).
		Element(children...)
}

func KeyholeSquareFull(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyholeSquareFullInnerSVG).
		Element(children...)
}

func LayerGroup(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layerGroupInnerSVG).
		Element(children...)
}

func LayersAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layersAltInnerSVG).
		Element(children...)
}

func LeftIndent(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftIndentInnerSVG).
		Element(children...)
}

func LeftIndentAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(leftIndentAltInnerSVG).
		Element(children...)
}

func Line(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineInnerSVG).
		Element(children...)
}

func LineSpacing(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineSpacingInnerSVG).
		Element(children...)
}

func LinkH(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkHInnerSVG).
		Element(children...)
}

func Linkedin(children ...engine.UberChild) *engine.UberElement {
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

func LinkedinAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkedinAltInnerSVG).
		Element(children...)
}

func Linux(children ...engine.UberChild) *engine.UberElement {
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

func ListUiAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listUiAltInnerSVG).
		Element(children...)
}

func ListUl(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listUlInnerSVG).
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

func LockAccess(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockAccessInnerSVG).
		Element(children...)
}

func LockAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockAltInnerSVG).
		Element(children...)
}

func LockOpenAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockOpenAltInnerSVG).
		Element(children...)
}

func Lottiefiles(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lottiefilesInnerSVG).
		Element(children...)
}

func MasterCard(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(masterCardInnerSVG).
		Element(children...)
}

func MediumM(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediumMInnerSVG).
		Element(children...)
}

func Microscope(children ...engine.UberChild) *engine.UberElement {
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

func Microsoft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microsoftInnerSVG).
		Element(children...)
}

func MinusSquareFull(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusSquareFullInnerSVG).
		Element(children...)
}

func Multiply(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiplyInnerSVG).
		Element(children...)
}

func ObjectGroup(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(objectGroupInnerSVG).
		Element(children...)
}

func ObjectUngroup(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(objectUngroupInnerSVG).
		Element(children...)
}

func Okta(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(oktaInnerSVG).
		Element(children...)
}

func Opera(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(operaInnerSVG).
		Element(children...)
}

func OperaAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(operaAltInnerSVG).
		Element(children...)
}

func Padlock(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(padlockInnerSVG).
		Element(children...)
}

func Pagelines(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pagelinesInnerSVG).
		Element(children...)
}

func Pagerduty(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pagerdutyInnerSVG).
		Element(children...)
}

func Paperclip(children ...engine.UberChild) *engine.UberElement {
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

func Paragraph(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paragraphInnerSVG).
		Element(children...)
}

func Paypal(children ...engine.UberChild) *engine.UberElement {
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

func Pentagon(children ...engine.UberChild) *engine.UberElement {
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

func PlusSquare(children ...engine.UberChild) *engine.UberElement {
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

func Polygon(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(polygonInnerSVG).
		Element(children...)
}

func Previous(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(previousInnerSVG).
		Element(children...)
}

func Process(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(processInnerSVG).
		Element(children...)
}

func React(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reactInnerSVG).
		Element(children...)
}

func RecordAudio(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(recordAudioInnerSVG).
		Element(children...)
}

func RedditAlienAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redditAlienAltInnerSVG).
		Element(children...)
}

func Redo(children ...engine.UberChild) *engine.UberElement {
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

func Repeat(children ...engine.UberChild) *engine.UberElement {
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

func RightIndentAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rightIndentAltInnerSVG).
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

func Ruler(children ...engine.UberChild) *engine.UberElement {
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

func RulerCombined(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rulerCombinedInnerSVG).
		Element(children...)
}

func Sanitizer(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sanitizerInnerSVG).
		Element(children...)
}

func SanitizerAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sanitizerAltInnerSVG).
		Element(children...)
}

func Scenery(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sceneryInnerSVG).
		Element(children...)
}

func Schedule(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scheduleInnerSVG).
		Element(children...)
}

func ShieldPlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldPlusInnerSVG).
		Element(children...)
}

func SignIn(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signInInnerSVG).
		Element(children...)
}

func SignInAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signInAltInnerSVG).
		Element(children...)
}

func SignOutAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signOutAltInnerSVG).
		Element(children...)
}

func SignalAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signalAltInnerSVG).
		Element(children...)
}

func SignalAltThree(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signalAltThreeInnerSVG).
		Element(children...)
}

func Signin(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signinInnerSVG).
		Element(children...)
}

func Signout(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signoutInnerSVG).
		Element(children...)
}

func Skype(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skypeInnerSVG).
		Element(children...)
}

func SkypeAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skypeAltInnerSVG).
		Element(children...)
}

func Slack(children ...engine.UberChild) *engine.UberElement {
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

func SlackAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slackAltInnerSVG).
		Element(children...)
}

func SnapchatAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snapchatAltInnerSVG).
		Element(children...)
}

func SnapchatGhost(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snapchatGhostInnerSVG).
		Element(children...)
}

func SnapchatSquare(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snapchatSquareInnerSVG).
		Element(children...)
}

func SocialDistancing(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(socialDistancingInnerSVG).
		Element(children...)
}

func Sorting(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortingInnerSVG).
		Element(children...)
}

func SpaceKey(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spaceKeyInnerSVG).
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

func SquareFull(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareFullInnerSVG).
		Element(children...)
}

func SquareShape(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareShapeInnerSVG).
		Element(children...)
}

func SqureShape(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squreShapeInnerSVG).
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

func StarHalfAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starHalfAltInnerSVG).
		Element(children...)
}

func StepForward(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stepForwardInnerSVG).
		Element(children...)
}

func Stethoscope(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stethoscopeInnerSVG).
		Element(children...)
}

func StethoscopeAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stethoscopeAltInnerSVG).
		Element(children...)
}

func StoreSlash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(storeSlashInnerSVG).
		Element(children...)
}

func Subject(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subjectInnerSVG).
		Element(children...)
}

func Swiggy(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swiggyInnerSVG).
		Element(children...)
}

func SyncExclamation(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(syncExclamationInnerSVG).
		Element(children...)
}

func SyncSlash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(syncSlashInnerSVG).
		Element(children...)
}

func Table(children ...engine.UberChild) *engine.UberElement {
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

func Telegram(children ...engine.UberChild) *engine.UberElement {
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

func TelegramAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(telegramAltInnerSVG).
		Element(children...)
}

func ThLarge(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thLargeInnerSVG).
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

func ToggleOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleOffInnerSVG).
		Element(children...)
}

func ToggleOn(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleOnInnerSVG).
		Element(children...)
}

func ToiletPaper(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toiletPaperInnerSVG).
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

func Tumblr(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tumblrInnerSVG).
		Element(children...)
}

func TumblrAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tumblrAltInnerSVG).
		Element(children...)
}

func TumblrSquare(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tumblrSquareInnerSVG).
		Element(children...)
}

func Twitter(children ...engine.UberChild) *engine.UberElement {
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

func TwitterAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitterAltInnerSVG).
		Element(children...)
}

func Unlock(children ...engine.UberChild) *engine.UberElement {
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

func UnlockAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockAltInnerSVG).
		Element(children...)
}

func UploadAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadAltInnerSVG).
		Element(children...)
}

func UserArrows(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userArrowsInnerSVG).
		Element(children...)
}

func UserMd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userMdInnerSVG).
		Element(children...)
}

func UserNurse(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userNurseInnerSVG).
		Element(children...)
}

func VectorSquare(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vectorSquareInnerSVG).
		Element(children...)
}

func VectorSquareAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vectorSquareAltInnerSVG).
		Element(children...)
}

func VirusSlash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(virusSlashInnerSVG).
		Element(children...)
}

func VisualStudio(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(visualStudioInnerSVG).
		Element(children...)
}

func Vk(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vkInnerSVG).
		Element(children...)
}

func VkAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vkAltInnerSVG).
		Element(children...)
}

func Vuejs(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vuejsInnerSVG).
		Element(children...)
}

func VuejsAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vuejsAltInnerSVG).
		Element(children...)
}

func WebGrid(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webGridInnerSVG).
		Element(children...)
}

func WebGridAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webGridAltInnerSVG).
		Element(children...)
}

func WebSection(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webSectionInnerSVG).
		Element(children...)
}

func WebSectionAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webSectionAltInnerSVG).
		Element(children...)
}

func Whatsapp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(whatsappInnerSVG).
		Element(children...)
}

func WindowGrid(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowGridInnerSVG).
		Element(children...)
}

func WindowMaximize(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowMaximizeInnerSVG).
		Element(children...)
}

func WindowSection(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowSectionInnerSVG).
		Element(children...)
}

func Windows(children ...engine.UberChild) *engine.UberElement {
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

func Wordpress(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wordpressInnerSVG).
		Element(children...)
}

func WordpressSimple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wordpressSimpleInnerSVG).
		Element(children...)
}

func WrapText(children ...engine.UberChild) *engine.UberElement {
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

func Youtube(children ...engine.UberChild) *engine.UberElement {
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

func ByName(name string) (*engine.UberElement, error) {
	switch name {
	case "adobe":
		return Adobe(), nil
	case "adobe-alt":
		return AdobeAlt(), nil
	case "airplay":
		return Airplay(), nil
	case "align":
		return Align(), nil
	case "align-alt":
		return AlignAlt(), nil
	case "align-center":
		return AlignCenter(), nil
	case "align-center-justify":
		return AlignCenterJustify(), nil
	case "align-justify":
		return AlignJustify(), nil
	case "align-left":
		return AlignLeft(), nil
	case "align-left-justify":
		return AlignLeftJustify(), nil
	case "align-letter-right":
		return AlignLetterRight(), nil
	case "align-right":
		return AlignRight(), nil
	case "align-right-justify":
		return AlignRightJustify(), nil
	case "amazon":
		return Amazon(), nil
	case "analysis":
		return Analysis(), nil
	case "analytics":
		return Analytics(), nil
	case "anchor":
		return Anchor(), nil
	case "android":
		return Android(), nil
	case "android-alt":
		return AndroidAlt(), nil
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
	case "angle-right-b":
		return AngleRightB(), nil
	case "angle-up":
		return AngleUp(), nil
	case "apple":
		return Apple(), nil
	case "apple-alt":
		return AppleAlt(), nil
	case "apps":
		return Apps(), nil
	case "arrow-circle-down":
		return ArrowCircleDown(), nil
	case "arrow-circle-left":
		return ArrowCircleLeft(), nil
	case "arrow-circle-right":
		return ArrowCircleRight(), nil
	case "arrow-circle-up":
		return ArrowCircleUp(), nil
	case "arrow-down-left":
		return ArrowDownLeft(), nil
	case "arrow-down-right":
		return ArrowDownRight(), nil
	case "arrow-up-left":
		return ArrowUpLeft(), nil
	case "arrow-up-right":
		return ArrowUpRight(), nil
	case "at":
		return At(), nil
	case "bag":
		return Bag(), nil
	case "bars":
		return Bars(), nil
	case "battery-bolt":
		return BatteryBolt(), nil
	case "battery-empty":
		return BatteryEmpty(), nil
	case "behance":
		return Behance(), nil
	case "behance-alt":
		return BehanceAlt(), nil
	case "bing":
		return Bing(), nil
	case "bitcoin":
		return Bitcoin(), nil
	case "bitcoin-alt":
		return BitcoinAlt(), nil
	case "blackberry":
		return Blackberry(), nil
	case "blogger":
		return Blogger(), nil
	case "blogger-alt":
		return BloggerAlt(), nil
	case "bookmark":
		return Bookmark(), nil
	case "border-alt":
		return BorderAlt(), nil
	case "border-bottom":
		return BorderBottom(), nil
	case "border-clear":
		return BorderClear(), nil
	case "border-horizontal":
		return BorderHorizontal(), nil
	case "border-inner":
		return BorderInner(), nil
	case "border-left":
		return BorderLeft(), nil
	case "border-out":
		return BorderOut(), nil
	case "border-right":
		return BorderRight(), nil
	case "border-top":
		return BorderTop(), nil
	case "border-vertical":
		return BorderVertical(), nil
	case "box":
		return Box(), nil
	case "briefcase":
		return Briefcase(), nil
	case "calender":
		return Calender(), nil
	case "chart":
		return Chart(), nil
	case "chart-pie":
		return ChartPie(), nil
	case "check":
		return Check(), nil
	case "check-circle":
		return CheckCircle(), nil
	case "check-square":
		return CheckSquare(), nil
	case "circle":
		return Circle(), nil
	case "circle-layer":
		return CircleLayer(), nil
	case "clinic-medical":
		return ClinicMedical(), nil
	case "clock":
		return Clock(), nil
	case "clock-eight":
		return ClockEight(), nil
	case "clock-five":
		return ClockFive(), nil
	case "clock-nine":
		return ClockNine(), nil
	case "clock-seven":
		return ClockSeven(), nil
	case "clock-ten":
		return ClockTen(), nil
	case "clock-three":
		return ClockThree(), nil
	case "clock-two":
		return ClockTwo(), nil
	case "columns":
		return Columns(), nil
	case "comment":
		return Comment(), nil
	case "comment-alt":
		return CommentAlt(), nil
	case "comment-alt-dots":
		return CommentAltDots(), nil
	case "comment-alt-message":
		return CommentAltMessage(), nil
	case "comment-alt-plus":
		return CommentAltPlus(), nil
	case "comment-dots":
		return CommentDots(), nil
	case "comment-message":
		return CommentMessage(), nil
	case "comment-plus":
		return CommentPlus(), nil
	case "compress":
		return Compress(), nil
	case "corner-down-left":
		return CornerDownLeft(), nil
	case "corner-down-right":
		return CornerDownRight(), nil
	case "corner-left-down":
		return CornerLeftDown(), nil
	case "corner-right-down":
		return CornerRightDown(), nil
	case "corner-up-left":
		return CornerUpLeft(), nil
	case "corner-up-right":
		return CornerUpRight(), nil
	case "coronavirus":
		return Coronavirus(), nil
	case "css3":
		return CssThree(), nil
	case "css3-simple":
		return CssThreeSimple(), nil
	case "cube":
		return Cube(), nil
	case "dialpad":
		return Dialpad(), nil
	case "dialpad-alt":
		return DialpadAlt(), nil
	case "direction":
		return Direction(), nil
	case "discord":
		return Discord(), nil
	case "docker":
		return Docker(), nil
	case "document-layout-center":
		return DocumentLayoutCenter(), nil
	case "document-layout-left":
		return DocumentLayoutLeft(), nil
	case "document-layout-right":
		return DocumentLayoutRight(), nil
	case "download-alt":
		return DownloadAlt(), nil
	case "dribbble":
		return Dribbble(), nil
	case "dropbox":
		return Dropbox(), nil
	case "ellipsis-h":
		return EllipsisH(), nil
	case "ellipsis-v":
		return EllipsisV(), nil
	case "entry":
		return Entry(), nil
	case "exclamation-circle":
		return ExclamationCircle(), nil
	case "exclamation-octagon":
		return ExclamationOctagon(), nil
	case "exclamation-triangle":
		return ExclamationTriangle(), nil
	case "exit":
		return Exit(), nil
	case "facebook":
		return Facebook(), nil
	case "facebook-f":
		return FacebookF(), nil
	case "facebook-messenger":
		return FacebookMessenger(), nil
	case "facebook-messenger-alt":
		return FacebookMessengerAlt(), nil
	case "favorite":
		return Favorite(), nil
	case "500px":
		return FiveHundredPx(), nil
	case "flip-h":
		return FlipH(), nil
	case "flip-h-alt":
		return FlipHAlt(), nil
	case "flip-v":
		return FlipV(), nil
	case "flip-v-alt":
		return FlipVAlt(), nil
	case "github":
		return Github(), nil
	case "github-alt":
		return GithubAlt(), nil
	case "gitlab":
		return Gitlab(), nil
	case "gitlab-alt":
		return GitlabAlt(), nil
	case "google":
		return Google(), nil
	case "google-drive":
		return GoogleDrive(), nil
	case "google-drive-alt":
		return GoogleDriveAlt(), nil
	case "google-hangouts":
		return GoogleHangouts(), nil
	case "google-hangouts-alt":
		return GoogleHangoutsAlt(), nil
	case "google-play":
		return GooglePlay(), nil
	case "graph-bar":
		return GraphBar(), nil
	case "grid":
		return Grid(), nil
	case "grids":
		return Grids(), nil
	case "grip-horizontal-line":
		return GripHorizontalLine(), nil
	case "head-side":
		return HeadSide(), nil
	case "head-side-cough":
		return HeadSideCough(), nil
	case "head-side-mask":
		return HeadSideMask(), nil
	case "hipchat":
		return Hipchat(), nil
	case "history":
		return History(), nil
	case "history-alt":
		return HistoryAlt(), nil
	case "horizontal-align-left":
		return HorizontalAlignLeft(), nil
	case "hospital":
		return Hospital(), nil
	case "hospital-square-sign":
		return HospitalSquareSign(), nil
	case "hospital-symbol":
		return HospitalSymbol(), nil
	case "house-user":
		return HouseUser(), nil
	case "html5":
		return HtmlFive(), nil
	case "html5-alt":
		return HtmlFiveAlt(), nil
	case "html3":
		return HtmlThree(), nil
	case "html3-alt":
		return HtmlThreeAlt(), nil
	case "image-v":
		return ImageV(), nil
	case "instagram":
		return Instagram(), nil
	case "instagram-alt":
		return InstagramAlt(), nil
	case "intercom":
		return Intercom(), nil
	case "intercom-alt":
		return IntercomAlt(), nil
	case "java-script":
		return JavaScript(), nil
	case "key-skeleton":
		return KeySkeleton(), nil
	case "key-skeleton-alt":
		return KeySkeletonAlt(), nil
	case "keyhole-circle":
		return KeyholeCircle(), nil
	case "keyhole-square":
		return KeyholeSquare(), nil
	case "keyhole-square-full":
		return KeyholeSquareFull(), nil
	case "layer-group":
		return LayerGroup(), nil
	case "layers-alt":
		return LayersAlt(), nil
	case "left-indent":
		return LeftIndent(), nil
	case "left-indent-alt":
		return LeftIndentAlt(), nil
	case "line":
		return Line(), nil
	case "line-spacing":
		return LineSpacing(), nil
	case "link-h":
		return LinkH(), nil
	case "linkedin":
		return Linkedin(), nil
	case "linkedin-alt":
		return LinkedinAlt(), nil
	case "linux":
		return Linux(), nil
	case "list-ui-alt":
		return ListUiAlt(), nil
	case "list-ul":
		return ListUl(), nil
	case "lock":
		return Lock(), nil
	case "lock-access":
		return LockAccess(), nil
	case "lock-alt":
		return LockAlt(), nil
	case "lock-open-alt":
		return LockOpenAlt(), nil
	case "lottiefiles":
		return Lottiefiles(), nil
	case "master-card":
		return MasterCard(), nil
	case "medium-m":
		return MediumM(), nil
	case "microscope":
		return Microscope(), nil
	case "microsoft":
		return Microsoft(), nil
	case "minus-square-full":
		return MinusSquareFull(), nil
	case "multiply":
		return Multiply(), nil
	case "object-group":
		return ObjectGroup(), nil
	case "object-ungroup":
		return ObjectUngroup(), nil
	case "okta":
		return Okta(), nil
	case "opera":
		return Opera(), nil
	case "opera-alt":
		return OperaAlt(), nil
	case "padlock":
		return Padlock(), nil
	case "pagelines":
		return Pagelines(), nil
	case "pagerduty":
		return Pagerduty(), nil
	case "paperclip":
		return Paperclip(), nil
	case "paragraph":
		return Paragraph(), nil
	case "paypal":
		return Paypal(), nil
	case "pentagon":
		return Pentagon(), nil
	case "plus-square":
		return PlusSquare(), nil
	case "polygon":
		return Polygon(), nil
	case "previous":
		return Previous(), nil
	case "process":
		return Process(), nil
	case "react":
		return React(), nil
	case "record-audio":
		return RecordAudio(), nil
	case "reddit-alien-alt":
		return RedditAlienAlt(), nil
	case "redo":
		return Redo(), nil
	case "refresh":
		return Refresh(), nil
	case "repeat":
		return Repeat(), nil
	case "right-indent-alt":
		return RightIndentAlt(), nil
	case "rocket":
		return Rocket(), nil
	case "ruler":
		return Ruler(), nil
	case "ruler-combined":
		return RulerCombined(), nil
	case "sanitizer":
		return Sanitizer(), nil
	case "sanitizer-alt":
		return SanitizerAlt(), nil
	case "scenery":
		return Scenery(), nil
	case "schedule":
		return Schedule(), nil
	case "shield-plus":
		return ShieldPlus(), nil
	case "sign-in":
		return SignIn(), nil
	case "sign-in-alt":
		return SignInAlt(), nil
	case "sign-out-alt":
		return SignOutAlt(), nil
	case "signal-alt":
		return SignalAlt(), nil
	case "signal-alt-3":
		return SignalAltThree(), nil
	case "signin":
		return Signin(), nil
	case "signout":
		return Signout(), nil
	case "skype":
		return Skype(), nil
	case "skype-alt":
		return SkypeAlt(), nil
	case "slack":
		return Slack(), nil
	case "slack-alt":
		return SlackAlt(), nil
	case "snapchat-alt":
		return SnapchatAlt(), nil
	case "snapchat-ghost":
		return SnapchatGhost(), nil
	case "snapchat-square":
		return SnapchatSquare(), nil
	case "social-distancing":
		return SocialDistancing(), nil
	case "sorting":
		return Sorting(), nil
	case "space-key":
		return SpaceKey(), nil
	case "square":
		return Square(), nil
	case "square-full":
		return SquareFull(), nil
	case "square-shape":
		return SquareShape(), nil
	case "squre-shape":
		return SqureShape(), nil
	case "star":
		return Star(), nil
	case "star-half-alt":
		return StarHalfAlt(), nil
	case "step-forward":
		return StepForward(), nil
	case "stethoscope":
		return Stethoscope(), nil
	case "stethoscope-alt":
		return StethoscopeAlt(), nil
	case "store-slash":
		return StoreSlash(), nil
	case "subject":
		return Subject(), nil
	case "swiggy":
		return Swiggy(), nil
	case "sync-exclamation":
		return SyncExclamation(), nil
	case "sync-slash":
		return SyncSlash(), nil
	case "table":
		return Table(), nil
	case "telegram":
		return Telegram(), nil
	case "telegram-alt":
		return TelegramAlt(), nil
	case "th-large":
		return ThLarge(), nil
	case "times-circle":
		return TimesCircle(), nil
	case "toggle-off":
		return ToggleOff(), nil
	case "toggle-on":
		return ToggleOn(), nil
	case "toilet-paper":
		return ToiletPaper(), nil
	case "triangle":
		return Triangle(), nil
	case "tumblr":
		return Tumblr(), nil
	case "tumblr-alt":
		return TumblrAlt(), nil
	case "tumblr-square":
		return TumblrSquare(), nil
	case "twitter":
		return Twitter(), nil
	case "twitter-alt":
		return TwitterAlt(), nil
	case "unlock":
		return Unlock(), nil
	case "unlock-alt":
		return UnlockAlt(), nil
	case "upload-alt":
		return UploadAlt(), nil
	case "user-arrows":
		return UserArrows(), nil
	case "user-md":
		return UserMd(), nil
	case "user-nurse":
		return UserNurse(), nil
	case "vector-square":
		return VectorSquare(), nil
	case "vector-square-alt":
		return VectorSquareAlt(), nil
	case "virus-slash":
		return VirusSlash(), nil
	case "visual-studio":
		return VisualStudio(), nil
	case "vk":
		return Vk(), nil
	case "vk-alt":
		return VkAlt(), nil
	case "vuejs":
		return Vuejs(), nil
	case "vuejs-alt":
		return VuejsAlt(), nil
	case "web-grid":
		return WebGrid(), nil
	case "web-grid-alt":
		return WebGridAlt(), nil
	case "web-section":
		return WebSection(), nil
	case "web-section-alt":
		return WebSectionAlt(), nil
	case "whatsapp":
		return Whatsapp(), nil
	case "window-grid":
		return WindowGrid(), nil
	case "window-maximize":
		return WindowMaximize(), nil
	case "window-section":
		return WindowSection(), nil
	case "windows":
		return Windows(), nil
	case "wordpress":
		return Wordpress(), nil
	case "wordpress-simple":
		return WordpressSimple(), nil
	case "wrap-text":
		return WrapText(), nil
	case "youtube":
		return Youtube(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in uim icon set", name)
	}
}
