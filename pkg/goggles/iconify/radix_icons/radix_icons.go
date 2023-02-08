package radix_icons

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	accessibilityInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 1 0 0 11.346a5.673 5.673 0 0 0 0-11.346ZM7.125 9c-.055.127-.793 2.96-.793 2.96a.5.5 0 1 1-.966-.26s.88-2.827.88-3.43V6.801l-1.958-.525a.5.5 0 1 1 .258-.966s1.654.563 2.3.563h1.309c.645 0 2.298-.563 2.298-.563a.5.5 0 1 1 .26.966l-1.966.527V8.27c0 .603.88 3.427.88 3.427a.5.5 0 0 1-.966.259S7.92 9.127 7.869 9c-.05-.127-.219-.127-.219-.127h-.307s-.173 0-.218.127ZM7.5 5.12a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25Z" clip-rule="evenodd"/>`
	activityLogInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A.5.5 0 0 1 .5 1h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Zm4 0a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm-4 3A.5.5 0 0 1 .5 7h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Zm4 0a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm-4 3a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Zm4 0a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	alignBaselineInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M10.5 1.75a.5.5 0 0 1 .47.33l3 8.32a.5.5 0 1 1-.94.34l-.982-2.724H8.952L7.97 10.74a.5.5 0 0 1-.94-.339l3-8.32a.5.5 0 0 1 .47-.33Zm0 1.974l1.241 3.442H9.26l1.24-3.442ZM2.5 2.1c.22 0 .4.18.4.4v7.034l1.317-1.317a.4.4 0 0 1 .565.566l-2 2a.4.4 0 0 1-.565 0l-2-2a.4.4 0 1 1 .565-.566L2.1 9.534V2.5c0-.22.18-.4.4-.4ZM.1 13.5a.4.4 0 0 1 .4-.4h14a.4.4 0 1 1 0 .8H.5a.4.4 0 0 1-.4-.4Z" clip-rule="evenodd"/>`
	alignBottomInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M9 3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v11H1.5a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1H9V3Z" clip-rule="evenodd"/>`
	alignCenterInnerSVG              = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 7.05V4h3v3.05H6zm-1 0H1.5a.45.45 0 0 0 0 .9H5v3.3c0 .414.336.75.75.75h3.5a.75.75 0 0 0 .75-.75v-3.3h3.5a.45.45 0 0 0 0-.9H10v-3.3A.75.75 0 0 0 9.25 3h-3.5a.75.75 0 0 0-.75.75v3.3zm4 .9V11H6V7.95h3z" fill="currentColor"/></g>`
	alignCenterHorizontallyInnerSVG  = `<path fill="currentColor" fill-rule="evenodd" d="M2 6a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h5v4.5a.5.5 0 0 0 1 0V9h5a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H8V1.5a.5.5 0 0 0-1 0V6H2Z" clip-rule="evenodd"/>`
	alignCenterVerticallyInnerSVG    = `<path fill="currentColor" fill-rule="evenodd" d="M7 1a1 1 0 0 0-1 1v5H1.5a.5.5 0 0 0 0 1H6v5a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V8h4.5a.5.5 0 0 0 0-1H9V2a1 1 0 0 0-1-1H7Z" clip-rule="evenodd"/>`
	alignEndInnerSVG                 = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 11V4h3v7H6zM5 3.75A.75.75 0 0 1 5.75 3h3.5a.75.75 0 0 1 .75.75v7.5a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1-.75-.75v-7.5zm-3.5 9.3a.45.45 0 0 0 0 .9h12a.45.45 0 1 0 0-.9h-12z" fill="currentColor"/></g>`
	alignHorizontalCentersInnerSVG   = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3.25 2C2.56 2 2 2.56 2 3.25v8.5c0 .69.56 1.25 1.25 1.25h2.5C6.44 13 7 12.44 7 11.75v-8.5C7 2.56 6.44 2 5.75 2h-2.5zM3 3.25A.25.25 0 0 1 3.25 3h2.5a.25.25 0 0 1 .25.25v8.5a.25.25 0 0 1-.25.25h-2.5a.25.25 0 0 1-.25-.25v-8.5zM9.25 4C8.56 4 8 4.56 8 5.25v4.5c0 .69.56 1.25 1.25 1.25h2.5c.69 0 1.25-.56 1.25-1.25v-4.5C13 4.56 12.44 4 11.75 4h-2.5z" fill="currentColor"/></g>`
	alignLeftInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M.5 1a.5.5 0 0 0-.5.5v12a.5.5 0 0 0 1 0V9h11a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H1V1.5A.5.5 0 0 0 .5 1Z" clip-rule="evenodd"/>`
	alignRightInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M14.5 1a.5.5 0 0 0-.5.5V6H3a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h11v4.5a.5.5 0 1 0 1 0v-12a.5.5 0 0 0-.5-.5Z" clip-rule="evenodd"/>`
	alignStartInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M1.5 1.05a.45.45 0 1 0 0 .9h12a.45.45 0 0 0 0-.9h-12zM6 11V4h3v7H6zM5 3.75A.75.75 0 0 1 5.75 3h3.5a.75.75 0 0 1 .75.75v7.5a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1-.75-.75v-7.5z" fill="currentColor"/></g>`
	alignStretchInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M1.05 1.5a.45.45 0 0 1 .45-.45h12a.45.45 0 0 1 0 .9h-12a.45.45 0 0 1-.45-.45zm0 12a.45.45 0 0 1 .45-.45h12a.45.45 0 0 1 0 .9h-12a.45.45 0 0 1-.45-.45zM6 11V4h3v7H6zM5 3.75A.75.75 0 0 1 5.75 3h3.5a.75.75 0 0 1 .75.75v7.5a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1-.75-.75v-7.5z" fill="currentColor"/></g>`
	alignTopInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 0a.5.5 0 0 0 0 1H6v11a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V1h4.5a.5.5 0 0 0 0-1h-12Z" clip-rule="evenodd"/>`
	alignVerticalCentersInnerSVG     = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 3.25C2 2.56 2.56 2 3.25 2h8.5c.69 0 1.25.56 1.25 1.25v2.5C13 6.44 12.44 7 11.75 7h-8.5C2.56 7 2 6.44 2 5.75v-2.5zM3.25 3a.25.25 0 0 0-.25.25v2.5c0 .138.112.25.25.25h8.5a.25.25 0 0 0 .25-.25v-2.5a.25.25 0 0 0-.25-.25h-8.5zM4 9.25C4 8.56 4.56 8 5.25 8h4.5c.69 0 1.25.56 1.25 1.25v2.5c0 .69-.56 1.25-1.25 1.25h-4.5C4.56 13 4 12.44 4 11.75v-2.5z" fill="currentColor"/></g>`
	allSidesInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.75L9.75 3h-4.5L7.5.75Zm0 13.5L9.75 12h-4.5l2.25 2.25Zm-4.5-9L.75 7.5L3 9.75v-4.5ZM14.25 7.5L12 5.25v4.5l2.25-2.25Z" clip-rule="evenodd"/>`
	angleInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M8.891 2.194a.5.5 0 0 1 .115.697L2.474 12H13.5a.5.5 0 0 1 0 1h-12a.5.5 0 0 1-.406-.791l7.1-9.9a.5.5 0 0 1 .697-.115ZM11.1 6.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM10.4 4a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.7 4.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm1.3 1.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`
	archiveInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M3.309 1a1 1 0 0 0-.894.553L1.053 4.276A.5.5 0 0 0 1 4.5V13a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V4.5a.5.5 0 0 0-.053-.224l-1.362-2.723A1 1 0 0 0 11.691 1H3.309Zm0 1H7v2H2.309l1-2ZM8 4V2h3.691l1 2H8Zm-.5 1H13v8H2V5h5.5Zm-2 2a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4Z" clip-rule="evenodd"/>`
	arrowBottomLeftInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M11.354 3.646a.5.5 0 0 1 0 .708L4.707 11H9a.5.5 0 0 1 0 1H3.5a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 1 0v4.293l6.646-6.647a.5.5 0 0 1 .708 0Z" clip-rule="evenodd"/>`
	arrowBottomRightInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M3.646 3.646a.5.5 0 0 0 0 .708L10.293 11H6a.5.5 0 0 0 0 1h5.5a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-1 0v4.293L4.354 3.646a.5.5 0 0 0-.708 0Z" clip-rule="evenodd"/>`
	arrowDownInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 2a.5.5 0 0 1 .5.5v8.793l3.146-3.147a.5.5 0 0 1 .708.708l-4 4a.5.5 0 0 1-.708 0l-4-4a.5.5 0 1 1 .708-.708L7 11.293V2.5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`
	arrowLeftInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M6.854 3.146a.5.5 0 0 1 0 .708L3.707 7H12.5a.5.5 0 0 1 0 1H3.707l3.147 3.146a.5.5 0 0 1-.708.708l-4-4a.5.5 0 0 1 0-.708l4-4a.5.5 0 0 1 .708 0Z" clip-rule="evenodd"/>`
	arrowRightInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M8.146 3.146a.5.5 0 0 1 .708 0l4 4a.5.5 0 0 1 0 .708l-4 4a.5.5 0 0 1-.708-.708L11.293 8H2.5a.5.5 0 0 1 0-1h8.793L8.146 3.854a.5.5 0 0 1 0-.708Z" clip-rule="evenodd"/>`
	arrowTopLeftInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M11.354 11.354a.5.5 0 0 0 0-.707L4.707 4H9a.5.5 0 0 0 0-1H3.5a.5.5 0 0 0-.5.5V9a.5.5 0 0 0 1 0V4.707l6.646 6.647a.5.5 0 0 0 .708 0Z" clip-rule="evenodd"/>`
	arrowTopRightInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M3.646 11.354a.5.5 0 0 1 0-.707L10.293 4H6a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 .5.5V9a.5.5 0 0 1-1 0V4.707l-6.646 6.647a.5.5 0 0 1-.708 0Z" clip-rule="evenodd"/>`
	arrowUpInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7.146 2.146a.5.5 0 0 1 .708 0l4 4a.5.5 0 0 1-.708.708L8 3.707V12.5a.5.5 0 0 1-1 0V3.707L3.854 6.854a.5.5 0 1 1-.708-.708l4-4Z" clip-rule="evenodd"/>`
	aspectRatioInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M2.5 2h10a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1-.5-.5v-10a.5.5 0 0 1 .5-.5ZM1 2.5A1.5 1.5 0 0 1 2.5 1h10A1.5 1.5 0 0 1 14 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 12.5v-10ZM7.5 4a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM8 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM7.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm1.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/>`
	avatarInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 0 0-4.193 9.494A4.971 4.971 0 0 1 7.5 9.025a4.97 4.97 0 0 1 4.193 2.296A5.673 5.673 0 0 0 7.5 1.827Zm3.482 10.152A4.023 4.023 0 0 0 7.5 9.975a4.023 4.023 0 0 0-3.482 2.004A5.648 5.648 0 0 0 7.5 13.173c1.312 0 2.52-.446 3.482-1.194ZM5.15 6.505a2.35 2.35 0 1 1 4.7 0a2.35 2.35 0 0 1-4.7 0Zm2.35-1.4a1.4 1.4 0 1 0 0 2.8a1.4 1.4 0 0 0 0-2.8Z" clip-rule="evenodd"/>`
	backpackInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M5 1a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v1h4a1 1 0 0 1 1 1v3c0 .889-.387 1.687-1 2.236V11.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 11.5V8.236A2.993 2.993 0 0 1 0 6V3a1 1 0 0 1 1-1h4V1Zm4 0v1H6V1h3ZM1 3h13v3a1.996 1.996 0 0 1-2 2H8v-.5a.5.5 0 0 0-1 0V8H3a1.996 1.996 0 0 1-2-2V3Zm6 6H3c-.35 0-.687-.06-1-.17v2.67a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V8.83c-.313.11-.65.17-1 .17H8v.5a.5.5 0 0 1-1 0V9Z" clip-rule="evenodd"/>`
	badgeInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 6h8a1.5 1.5 0 0 1 0 3h-8a1.5 1.5 0 1 1 0-3ZM1 7.5A2.5 2.5 0 0 1 3.5 5h8a2.5 2.5 0 0 1 0 5h-8A2.5 2.5 0 0 1 1 7.5ZM4.5 7a.5.5 0 0 0 0 1h6a.5.5 0 0 0 0-1h-6Z" clip-rule="evenodd"/>`
	barChartInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M11.5 1a.5.5 0 0 1 .5.5v12a.5.5 0 0 1-1 0v-12a.5.5 0 0 1 .5-.5Zm-2 2a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-1 0v-10a.5.5 0 0 1 .5-.5Zm4 0a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-1 0v-10a.5.5 0 0 1 .5-.5Zm-8 1a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-1 0v-9a.5.5 0 0 1 .5-.5Zm-4 1a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-1 0v-8a.5.5 0 0 1 .5-.5Zm6 0a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-1 0v-8a.5.5 0 0 1 .5-.5Zm-4 2a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0v-6a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`
	bellInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M8.601 1.25a1.1 1.1 0 0 1-.8 1.06A4.5 4.5 0 0 1 12 6.8v3.45c0 .806.033 1.457.724 1.803A.5.5 0 0 1 12.5 13H8.161a1 1 0 1 1-1.323 0H2.5a.5.5 0 0 1-.224-.947c.691-.346.724-.997.724-1.803V6.8a4.5 4.5 0 0 1 4.2-4.49a1.1 1.1 0 1 1 1.401-1.06ZM7.5 3.3A3.5 3.5 0 0 0 4 6.8v3.5c0 .446.001 1.108-.3 1.7h7.6c-.301-.592-.3-1.254-.3-1.7V6.8a3.5 3.5 0 0 0-3.5-3.5Z" clip-rule="evenodd"/>`
	blendingModeInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M3 9c0-2.5 1.5-4.75 4.5-7.5c3 2.75 4.5 5 4.5 7.5a4.5 4.5 0 1 1-9 0Zm7.952-.697c-1.279-.482-2.664.16-3.962.76c-1.057.488-2.056.95-2.893.759A3.51 3.51 0 0 1 4 9c0-1.888 1.027-3.728 3.5-6.126c2.168 2.102 3.225 3.776 3.452 5.43Z" clip-rule="evenodd"/>`
	bookmarkInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3 2.5a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v11a.5.5 0 0 1-.765.424L7.5 11.59l-3.735 2.334A.5.5 0 0 1 3 13.5v-11ZM4 3v9.598l2.97-1.856a1 1 0 0 1 1.06 0L11 12.598V3H4Z" clip-rule="evenodd"/>`
	bookmarkFilledInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 2a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .765.424L7.5 11.59l3.735 2.334A.5.5 0 0 0 12 13.5v-11a.5.5 0 0 0-.5-.5h-8Z" clip-rule="evenodd"/>`
	borderAllInnerSVG                = `<g fill="currentColor"><path fill-rule="evenodd" d="M.25 1A.75.75 0 0 1 1 .25h13a.75.75 0 0 1 .75.75v13a.75.75 0 0 1-.75.75H1A.75.75 0 0 1 .25 14V1Zm1.5.75v11.5h11.5V1.75H1.75Z" clip-rule="evenodd"/><rect width="1" height="1" x="7" y="5" rx=".5"/><rect width="1" height="1" x="7" y="3" rx=".5"/><rect width="1" height="1" x="7" y="7" rx=".5"/><rect width="1" height="1" x="5" y="7" rx=".5"/><rect width="1" height="1" x="3" y="7" rx=".5"/><rect width="1" height="1" x="9" y="7" rx=".5"/><rect width="1" height="1" x="11" y="7" rx=".5"/><rect width="1" height="1" x="7" y="9" rx=".5"/><rect width="1" height="1" x="7" y="11" rx=".5"/></g>`
	borderBottomInnerSVG             = `<g fill="currentColor"><path fill-rule="evenodd" d="M1 13.25h13v1.5H1v-1.5Z" clip-rule="evenodd"/><rect width="1" height="1" x="7" y="5" rx=".5"/><rect width="1" height="1" x="13" y="5" rx=".5"/><rect width="1" height="1" x="7" y="3" rx=".5"/><rect width="1" height="1" x="13" y="3" rx=".5"/><rect width="1" height="1" x="7" y="7" rx=".5"/><rect width="1" height="1" x="7" y="1" rx=".5"/><rect width="1" height="1" x="13" y="7" rx=".5"/><rect width="1" height="1" x="13" y="1" rx=".5"/><rect width="1" height="1" x="5" y="7" rx=".5"/><rect width="1" height="1" x="5" y="1" rx=".5"/><rect width="1" height="1" x="3" y="7" rx=".5"/><rect width="1" height="1" x="3" y="1" rx=".5"/><rect width="1" height="1" x="9" y="7" rx=".5"/><rect width="1" height="1" x="9" y="1" rx=".5"/><rect width="1" height="1" x="11" y="7" rx=".5"/><rect width="1" height="1" x="11" y="1" rx=".5"/><rect width="1" height="1" x="7" y="9" rx=".5"/><rect width="1" height="1" x="13" y="9" rx=".5"/><rect width="1" height="1" x="7" y="11" rx=".5"/><rect width="1" height="1" x="13" y="11" rx=".5"/><rect width="1" height="1" x="1" y="5" rx=".5"/><rect width="1" height="1" x="1" y="3" rx=".5"/><rect width="1" height="1" x="1" y="7" rx=".5"/><rect width="1" height="1" x="1" y="1" rx=".5"/><rect width="1" height="1" x="1" y="9" rx=".5"/><rect width="1" height="1" x="1" y="11" rx=".5"/></g>`
	borderDashedInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M0 7.5A.5.5 0 0 1 .5 7H3a.5.5 0 0 1 0 1H.5a.5.5 0 0 1-.5-.5Zm5.75 0a.5.5 0 0 1 .5-.5h2.5a.5.5 0 0 1 0 1h-2.5a.5.5 0 0 1-.5-.5ZM12 7a.5.5 0 0 0 0 1h2.5a.5.5 0 0 0 0-1H12Z" clip-rule="evenodd"/>`
	borderDottedInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 6.625a.875.875 0 1 0 0 1.75a.875.875 0 0 0 0-1.75Zm4 0a.875.875 0 1 0 0 1.75a.875.875 0 0 0 0-1.75Zm4 0a.875.875 0 1 0 0 1.75a.875.875 0 0 0 0-1.75Zm3.125.875a.875.875 0 1 1 1.75 0a.875.875 0 0 1-1.75 0Z" clip-rule="evenodd"/>`
	borderLeftInnerSVG               = `<g fill="currentColor"><path fill-rule="evenodd" d="M1.75 1v13H.25V1h1.5Z" clip-rule="evenodd"/><rect width="1" height="1" x="10" y="7" rx=".5" transform="rotate(90 10 7)"/><rect width="1" height="1" x="10" y="13" rx=".5" transform="rotate(90 10 13)"/><rect width="1" height="1" x="12" y="7" rx=".5" transform="rotate(90 12 7)"/><rect width="1" height="1" x="12" y="13" rx=".5" transform="rotate(90 12 13)"/><rect width="1" height="1" x="8" y="7" rx=".5" transform="rotate(90 8 7)"/><rect width="1" height="1" x="14" y="7" rx=".5" transform="rotate(90 14 7)"/><rect width="1" height="1" x="8" y="13" rx=".5" transform="rotate(90 8 13)"/><rect width="1" height="1" x="14" y="13" rx=".5" transform="rotate(90 14 13)"/><rect width="1" height="1" x="8" y="5" rx=".5" transform="rotate(90 8 5)"/><rect width="1" height="1" x="14" y="5" rx=".5" transform="rotate(90 14 5)"/><rect width="1" height="1" x="8" y="3" rx=".5" transform="rotate(90 8 3)"/><rect width="1" height="1" x="14" y="3" rx=".5" transform="rotate(90 14 3)"/><rect width="1" height="1" x="8" y="9" rx=".5" transform="rotate(90 8 9)"/><rect width="1" height="1" x="14" y="9" rx=".5" transform="rotate(90 14 9)"/><rect width="1" height="1" x="8" y="11" rx=".5" transform="rotate(90 8 11)"/><rect width="1" height="1" x="14" y="11" rx=".5" transform="rotate(90 14 11)"/><rect width="1" height="1" x="6" y="7" rx=".5" transform="rotate(90 6 7)"/><rect width="1" height="1" x="6" y="13" rx=".5" transform="rotate(90 6 13)"/><rect width="1" height="1" x="4" y="7" rx=".5" transform="rotate(90 4 7)"/><rect width="1" height="1" x="4" y="13" rx=".5" transform="rotate(90 4 13)"/><rect width="1" height="1" x="10" y="1" rx=".5" transform="rotate(90 10 1)"/><rect width="1" height="1" x="12" y="1" rx=".5" transform="rotate(90 12 1)"/><rect width="1" height="1" x="8" y="1" rx=".5" transform="rotate(90 8 1)"/><rect width="1" height="1" x="14" y="1" rx=".5" transform="rotate(90 14 1)"/><rect width="1" height="1" x="6" y="1" rx=".5" transform="rotate(90 6 1)"/><rect width="1" height="1" x="4" y="1" rx=".5" transform="rotate(90 4 1)"/></g>`
	borderNoneInnerSVG               = `<g fill="currentColor"><rect width="1" height="1" x="7" y="5.025" rx=".5"/><rect width="1" height="1" x="13" y="5.025" rx=".5"/><rect width="1" height="1" x="7" y="3.025" rx=".5"/><rect width="1" height="1" x="13" y="3.025" rx=".5"/><rect width="1" height="1" x="7" y="7.025" rx=".5"/><rect width="1" height="1" x="7" y="13.025" rx=".5"/><rect width="1" height="1" x="7" y="1.025" rx=".5"/><rect width="1" height="1" x="13" y="7.025" rx=".5"/><rect width="1" height="1" x="13" y="13.025" rx=".5"/><rect width="1" height="1" x="13" y="1.025" rx=".5"/><rect width="1" height="1" x="5" y="7.025" rx=".5"/><rect width="1" height="1" x="5" y="13.025" rx=".5"/><rect width="1" height="1" x="5" y="1.025" rx=".5"/><rect width="1" height="1" x="3" y="7.025" rx=".5"/><rect width="1" height="1" x="3" y="13.025" rx=".5"/><rect width="1" height="1" x="3" y="1.025" rx=".5"/><rect width="1" height="1" x="9" y="7.025" rx=".5"/><rect width="1" height="1" x="9" y="13.025" rx=".5"/><rect width="1" height="1" x="9" y="1.025" rx=".5"/><rect width="1" height="1" x="11" y="7.025" rx=".5"/><rect width="1" height="1" x="11" y="13.025" rx=".5"/><rect width="1" height="1" x="11" y="1.025" rx=".5"/><rect width="1" height="1" x="7" y="9.025" rx=".5"/><rect width="1" height="1" x="13" y="9.025" rx=".5"/><rect width="1" height="1" x="7" y="11.025" rx=".5"/><rect width="1" height="1" x="13" y="11.025" rx=".5"/><rect width="1" height="1" x="1" y="5.025" rx=".5"/><rect width="1" height="1" x="1" y="3.025" rx=".5"/><rect width="1" height="1" x="1" y="7.025" rx=".5"/><rect width="1" height="1" x="1" y="13.025" rx=".5"/><rect width="1" height="1" x="1" y="1.025" rx=".5"/><rect width="1" height="1" x="1" y="9.025" rx=".5"/><rect width="1" height="1" x="1" y="11.025" rx=".5"/></g>`
	borderRightInnerSVG              = `<g fill="currentColor"><path fill-rule="evenodd" d="M13.25 1v13h1.5V1h-1.5Z" clip-rule="evenodd"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 5 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 5 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 3 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 3 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 5)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 5)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 3)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 3)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 9)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 9)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 11)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 11)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 9 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 9 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 11 7)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 11 13)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 5 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 3 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 7 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 1 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 9 1)"/><rect width="1" height="1" rx=".5" transform="matrix(0 1 1 0 11 1)"/></g>`
	borderSolidInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M1.25 7.5a.5.5 0 0 1 .5-.5h11.5a.5.5 0 0 1 0 1H1.75a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	borderSplitInnerSVG              = `<g fill="currentColor"><rect width="1" height="1" x="7" y="5.025" rx=".5"/><rect width="1" height="1" x="7" y="3.025" rx=".5"/><rect width="1" height="1" x="7" y="7.025" rx=".5"/><rect width="1" height="1" x="7" y="13.025" rx=".5"/><rect width="1" height="1" x="7" y="1.025" rx=".5"/><rect width="1" height="1" x="13" y="7.025" rx=".5"/><rect width="1" height="1" x="5" y="7.025" rx=".5"/><rect width="1" height="1" x="3" y="7.025" rx=".5"/><rect width="1" height="1" x="9" y="7.025" rx=".5"/><rect width="1" height="1" x="11" y="7.025" rx=".5"/><rect width="1" height="1" x="7" y="9.025" rx=".5"/><rect width="1" height="1" x="7" y="11.025" rx=".5"/><rect width="1" height="1" x="1" y="7.025" rx=".5"/><path fill-rule="evenodd" d="M1 1.5a.5.5 0 0 1 .5-.5H6v1H2v4H1V1.5ZM13 2H9V1h4.5a.5.5 0 0 1 .5.5V6h-1V2ZM1 13.5V9h1v4h4v1H1.5a.5.5 0 0 1-.5-.5Zm12-.5V9h1v4.5a.5.5 0 0 1-.5.5h-4v-1H13Z" clip-rule="evenodd"/></g>`
	borderStyleInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 3a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1h-12ZM1 7.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Zm0 4a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.5.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.5.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-7-4a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2Zm4.5.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	borderTopInnerSVG                = `<g fill="currentColor"><path fill-rule="evenodd" d="M14 1.75H1V.25h13v1.5Z" clip-rule="evenodd"/><rect width="1" height="1" x="8" y="10" rx=".5" transform="rotate(-180 8 10)"/><rect width="1" height="1" x="2" y="10" rx=".5" transform="rotate(-180 2 10)"/><rect width="1" height="1" x="8" y="12" rx=".5" transform="rotate(-180 8 12)"/><rect width="1" height="1" x="2" y="12" rx=".5" transform="rotate(-180 2 12)"/><rect width="1" height="1" x="8" y="8" rx=".5" transform="rotate(-180 8 8)"/><rect width="1" height="1" x="8" y="14" rx=".5" transform="rotate(-180 8 14)"/><rect width="1" height="1" x="2" y="8" rx=".5" transform="rotate(-180 2 8)"/><rect width="1" height="1" x="2" y="14" rx=".5" transform="rotate(-180 2 14)"/><rect width="1" height="1" x="10" y="8" rx=".5" transform="rotate(-180 10 8)"/><rect width="1" height="1" x="10" y="14" rx=".5" transform="rotate(-180 10 14)"/><rect width="1" height="1" x="12" y="8" rx=".5" transform="rotate(-180 12 8)"/><rect width="1" height="1" x="12" y="14" rx=".5" transform="rotate(-180 12 14)"/><rect width="1" height="1" x="6" y="8" rx=".5" transform="rotate(-180 6 8)"/><rect width="1" height="1" x="6" y="14" rx=".5" transform="rotate(-180 6 14)"/><rect width="1" height="1" x="4" y="8" rx=".5" transform="rotate(-180 4 8)"/><rect width="1" height="1" x="4" y="14" rx=".5" transform="rotate(-180 4 14)"/><rect width="1" height="1" x="8" y="6" rx=".5" transform="rotate(-180 8 6)"/><rect width="1" height="1" x="2" y="6" rx=".5" transform="rotate(-180 2 6)"/><rect width="1" height="1" x="8" y="4" rx=".5" transform="rotate(-180 8 4)"/><rect width="1" height="1" x="2" y="4" rx=".5" transform="rotate(-180 2 4)"/><rect width="1" height="1" x="14" y="10" rx=".5" transform="rotate(-180 14 10)"/><rect width="1" height="1" x="14" y="12" rx=".5" transform="rotate(-180 14 12)"/><rect width="1" height="1" x="14" y="8" rx=".5" transform="rotate(-180 14 8)"/><rect width="1" height="1" x="14" y="14" rx=".5" transform="rotate(-180 14 14)"/><rect width="1" height="1" x="14" y="6" rx=".5" transform="rotate(-180 14 6)"/><rect width="1" height="1" x="14" y="4" rx=".5" transform="rotate(-180 14 4)"/></g>`
	borderWidthInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M1 3h13v1H1V3Zm0 3h13v2H1V6Zm13 4.25H1v2.5h13v-2.5Z" clip-rule="evenodd"/>`
	boxInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M12.5 2h-10a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5v-10a.5.5 0 0 0-.5-.5Zm-10-1A1.5 1.5 0 0 0 1 2.5v10A1.5 1.5 0 0 0 2.5 14h10a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 12.5 1h-10Z" clip-rule="evenodd"/>`
	boxModelInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M2 1a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H2Zm0 1h11v11H2V2Zm2.5 2a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-6Zm.5 6V5h5v5H5Z" clip-rule="evenodd"/>`
	buttonInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M2 5h11a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1ZM0 6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V6Zm4.5.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Zm2.25.75a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0Zm3.75-.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z" clip-rule="evenodd"/>`
	calendarInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 1a.5.5 0 0 1 .5.5V2h5v-.5a.5.5 0 0 1 1 0V2h1.5A1.5 1.5 0 0 1 14 3.5v9a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 12.5v-9A1.5 1.5 0 0 1 2.5 2H4v-.5a.5.5 0 0 1 .5-.5ZM10 3v.5a.5.5 0 0 0 1 0V3h1.5a.5.5 0 0 1 .5.5V5H2V3.5a.5.5 0 0 1 .5-.5H4v.5a.5.5 0 0 0 1 0V3h5ZM2 6v6.5a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V6H2Zm5 1.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM9.5 7a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.5.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm.5 1.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM9 9.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM7.5 9a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM5 9.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM3.5 9a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM3 11.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.5.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`
	cameraInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H2ZM0 4a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V4Zm2 .25A.25.25 0 0 1 2.25 4h2.5a.25.25 0 0 1 .25.25v1.505a.25.25 0 0 1-.25.25h-2.5a.25.25 0 0 1-.25-.25V4.25Zm10.101 3.334a2.601 2.601 0 1 1-5.202 0a2.601 2.601 0 0 1 5.202 0Zm1 0a3.601 3.601 0 1 1-7.202 0a3.601 3.601 0 0 1 7.202 0Z" clip-rule="evenodd"/>`
	cardStackInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2 3.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1-.5-.5v-6Zm0 7.415A1.5 1.5 0 0 1 1 9.5v-6A1.5 1.5 0 0 1 2.5 2h10A1.5 1.5 0 0 1 14 3.5v6a1.5 1.5 0 0 1-1 1.415v.585a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 11.5v-.585ZM12 11v.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5V11h9Z" clip-rule="evenodd"/>`
	cardStackMinusInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M2.5 3a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-10ZM1 9.5a1.5 1.5 0 0 0 1 1.415v.585A1.5 1.5 0 0 0 3.5 13h8a1.5 1.5 0 0 0 1.5-1.5v-.585A1.5 1.5 0 0 0 14 9.5v-6A1.5 1.5 0 0 0 12.5 2h-10A1.5 1.5 0 0 0 1 3.5v6Zm11 2V11H3v.5a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5ZM5.5 6a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4Z" clip-rule="evenodd"/>`
	cardStackPlusInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M2 3.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1-.5-.5v-6Zm0 7.415A1.5 1.5 0 0 1 1 9.5v-6A1.5 1.5 0 0 1 2.5 2h10A1.5 1.5 0 0 1 14 3.5v6a1.5 1.5 0 0 1-1 1.415v.585a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 11.5v-.585ZM12 11v.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5V11h9ZM5 6.5a.5.5 0 0 1 .5-.5H7V4.5a.5.5 0 0 1 1 0V6h1.5a.5.5 0 0 1 0 1H8v1.5a.5.5 0 0 1-1 0V7H5.5a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	caretDownInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M4.182 6.182a.45.45 0 0 1 .636 0L7.5 8.864l2.682-2.682a.45.45 0 0 1 .636.636l-3 3a.45.45 0 0 1-.636 0l-3-3a.45.45 0 0 1 0-.636Z" clip-rule="evenodd"/>`
	caretLeftInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M8.818 4.182a.45.45 0 0 1 0 .636L6.136 7.5l2.682 2.682a.45.45 0 1 1-.636.636l-3-3a.45.45 0 0 1 0-.636l3-3a.45.45 0 0 1 .636 0Z" clip-rule="evenodd"/>`
	caretRightInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M6.182 4.182a.45.45 0 0 1 .636 0l3 3a.45.45 0 0 1 0 .636l-3 3a.45.45 0 1 1-.636-.636L8.864 7.5L6.182 4.818a.45.45 0 0 1 0-.636Z" clip-rule="evenodd"/>`
	caretSortInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M4.932 5.432a.45.45 0 1 0 .636.636L7.5 4.136l1.932 1.932a.45.45 0 0 0 .636-.636l-2.25-2.25a.45.45 0 0 0-.636 0l-2.25 2.25Zm5.136 4.136a.45.45 0 0 0-.636-.636L7.5 10.864L5.568 8.932a.45.45 0 0 0-.636.636l2.25 2.25a.45.45 0 0 0 .636 0l2.25-2.25Z" clip-rule="evenodd"/>`
	caretUpInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M4.182 8.818a.45.45 0 0 1 0-.636l3-3a.45.45 0 0 1 .636 0l3 3a.45.45 0 0 1-.636.636L7.5 6.136L4.818 8.818a.45.45 0 0 1-.636 0Z" clip-rule="evenodd"/>`
	chatBubbleInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M12.5 3h-10A1.5 1.5 0 0 0 1 4.5v5A1.5 1.5 0 0 0 2.5 11h5a.5.5 0 0 1 .354.146L10 13.293V11.5a.5.5 0 0 1 .5-.5h2A1.5 1.5 0 0 0 14 9.5v-5A1.5 1.5 0 0 0 12.5 3Zm-10-1h10A2.5 2.5 0 0 1 15 4.5v5a2.5 2.5 0 0 1-2.5 2.5H11v2.5a.5.5 0 0 1-.854.354L7.293 12H2.5A2.5 2.5 0 0 1 0 9.5v-5A2.5 2.5 0 0 1 2.5 2Z" clip-rule="evenodd"/>`
	checkInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M11.467 3.727c.289.189.37.576.181.865l-4.25 6.5a.625.625 0 0 1-.944.12l-2.75-2.5a.625.625 0 0 1 .841-.925l2.208 2.007l3.849-5.886a.625.625 0 0 1 .865-.181Z" clip-rule="evenodd"/>`
	checkCircledInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0Zm8.332-1.962a.5.5 0 0 0-.818-.576L6.52 8.972L5.357 7.787a.5.5 0 0 0-.714.7L6.227 10.1a.5.5 0 0 0 .765-.062l3.167-4.5Z" clip-rule="evenodd"/>`
	checkboxInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3 3h9v9H3V3ZM2 3a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3Zm8.35 2.511a.5.5 0 0 0-.825-.566L6.64 9.15L5.197 7.41a.5.5 0 0 0-.77.638l1.866 2.25a.5.5 0 0 0 .797-.037l3.26-4.749Z" clip-rule="evenodd"/>`
	chevronDownInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M3.135 6.158a.5.5 0 0 1 .707-.023L7.5 9.565l3.658-3.43a.5.5 0 0 1 .684.73l-4 3.75a.5.5 0 0 1-.684 0l-4-3.75a.5.5 0 0 1-.023-.707Z" clip-rule="evenodd"/>`
	chevronLeftInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M8.842 3.135a.5.5 0 0 1 .023.707L5.435 7.5l3.43 3.658a.5.5 0 0 1-.73.684l-3.75-4a.5.5 0 0 1 0-.684l3.75-4a.5.5 0 0 1 .707-.023Z" clip-rule="evenodd"/>`
	chevronRightInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M6.158 3.135a.5.5 0 0 1 .707.023l3.75 4a.5.5 0 0 1 0 .684l-3.75 4a.5.5 0 1 1-.73-.684L9.566 7.5l-3.43-3.658a.5.5 0 0 1 .023-.707Z" clip-rule="evenodd"/>`
	chevronUpInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M3.135 8.842a.5.5 0 0 0 .707.023L7.5 5.435l3.658 3.43a.5.5 0 0 0 .684-.73l-4-3.75a.5.5 0 0 0-.684 0l-4 3.75a.5.5 0 0 0-.023.707Z" clip-rule="evenodd"/>`
	circleInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 1 0 0 11.346a5.673 5.673 0 0 0 0-11.346Z" clip-rule="evenodd"/>`
	circleBackslashInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM3.858 3.151a5.673 5.673 0 0 1 7.992 7.992L3.857 3.15Zm-.707.707a5.673 5.673 0 0 0 7.992 7.992L3.15 3.857Z" clip-rule="evenodd"/>`
	clipboardInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M5 2V1h5v1H5Zm-.25-2A.75.75 0 0 0 4 .75V1h-.5A1.5 1.5 0 0 0 2 2.5v10A1.5 1.5 0 0 0 3.5 14h8a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 11.5 1H11V.75a.75.75 0 0 0-.75-.75h-5.5ZM11 2v.25a.75.75 0 0 1-.75.75h-5.5A.75.75 0 0 1 4 2.25V2h-.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-10a.5.5 0 0 0-.5-.5H11Z" clip-rule="evenodd"/>`
	clipboardCopyInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M5 2V1h5v1H5Zm-.25-2A.75.75 0 0 0 4 .75V1h-.5A1.5 1.5 0 0 0 2 2.5v10A1.5 1.5 0 0 0 3.5 14H7v-1H3.5a.5.5 0 0 1-.5-.5v-10a.5.5 0 0 1 .5-.5H4v.25c0 .414.336.75.75.75h5.5a.75.75 0 0 0 .75-.75V2h.5a.5.5 0 0 1 .5.5V7h1V2.5A1.5 1.5 0 0 0 11.5 1H11V.75a.75.75 0 0 0-.75-.75h-5.5ZM9 8.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm1.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm1.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm.5 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-.5 2.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm0 2a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-6-4a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm.5 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM8.5 15a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm2.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm1.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/>`
	clockInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0ZM8 4.5a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .146.354l2 2a.5.5 0 0 0 .708-.708L8 7.293V4.5Z" clip-rule="evenodd"/>`
	codeInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M9.964 2.686a.5.5 0 1 0-.928-.372l-4 10a.5.5 0 1 0 .928.372l4-10Zm-6.11 2.46a.5.5 0 0 1 0 .708L2.207 7.5l1.647 1.646a.5.5 0 1 1-.708.708l-2-2a.5.5 0 0 1 0-.708l2-2a.5.5 0 0 1 .708 0Zm7.292 0a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708-.708L12.793 7.5l-1.647-1.646a.5.5 0 0 1 0-.708Z" clip-rule="evenodd"/>`
	codesandboxLogoInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M7.711.797a.5.5 0 0 0-.422 0l-6 2.8A.5.5 0 0 0 1 4.05v6.9a.5.5 0 0 0 .289.453l6 2.8a.5.5 0 0 0 .422 0l6-2.8A.5.5 0 0 0 14 10.95v-6.9a.5.5 0 0 0-.289-.453l-6-2.8ZM7.5 3.157L5.98 2.51L7.5 1.8l1.52.71l-1.52.646Zm.196 1.003l2.542-1.08l2.034.949L7.5 6.057L2.728 4.029l2.034-.95L7.304 4.16a.5.5 0 0 0 .392 0ZM8 6.93l5-2.124V7.93l-1.918.882a1 1 0 0 0-.582.908v2.078L8 12.965V6.93Zm3.5 4.402l1.5-.7V9.03l-1.5.69v1.612ZM7 6.93v6.034l-2.498-1.166V9.72a1 1 0 0 0-.582-.908L2 7.929V4.806L7 6.93Zm-5 3.7l1.502.702V9.72L2 9.03v1.602Z" clip-rule="evenodd"/>`
	colorWheelInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0Zm2.904-4.284A5.649 5.649 0 0 1 7.1 1.84v4.693L3.781 3.216Zm-.565.565A5.649 5.649 0 0 0 1.84 7.1h4.693L3.216 3.781ZM6.534 7.9H1.841a5.649 5.649 0 0 0 1.375 3.319L6.534 7.9Zm-2.753 3.884A5.65 5.65 0 0 0 7.1 13.16V8.466l-3.319 3.318ZM7.9 8.466v4.693a5.65 5.65 0 0 0 3.318-1.375L7.9 8.466Zm3.884 2.752A5.648 5.648 0 0 0 13.16 7.9H8.466l3.318 3.318ZM8.466 7.1h4.693a5.65 5.65 0 0 0-1.375-3.319L8.466 7.1Zm2.753-3.884A5.649 5.649 0 0 0 7.9 1.84v4.693l3.319-3.318Z" clip-rule="evenodd"/>`
	columnSpacingInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a.5.5 0 0 0-1 0v12a.5.5 0 0 0 1 0v-12ZM3.318 5.818a.45.45 0 1 0-.636-.636l-2 2a.45.45 0 0 0 0 .636l2 2a.45.45 0 1 0 .636-.636L2.086 7.95H5.5a.45.45 0 1 0 0-.9H2.086l1.232-1.232Zm9-.636a.45.45 0 1 0-.636.636l1.232 1.232H9.5a.45.45 0 0 0 0 .9h3.414l-1.232 1.232a.45.45 0 0 0 .636.636l2-2a.45.45 0 0 0 0-.636l-2-2Z" clip-rule="evenodd"/>`
	columnsInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M2.15 14V1H.85v13h1.3Zm4 0V1h-1.3v13h1.3Zm4-13v13h-1.3V1h1.3Zm4 13V1h-1.3v13h1.3Z" clip-rule="evenodd"/>`
	commitInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M9.95 7.5a2.45 2.45 0 1 1-4.9 0a2.45 2.45 0 0 1 4.9 0Zm.913.5a3.4 3.4 0 0 1-6.726 0H.5a.5.5 0 0 1 0-1h3.637a3.4 3.4 0 0 1 6.726 0H14.5a.5.5 0 0 1 0 1h-3.637Z" clip-rule="evenodd"/>`
	componentBooleanInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M7.854 1.49a.5.5 0 0 0-.707 0L1.49 7.146a.5.5 0 0 0 0 .708l5.657 5.656a.5.5 0 0 0 .707 0l5.657-5.656a.5.5 0 0 0 0-.708L7.854 1.49ZM7.5 2.55L2.55 7.5l4.95 4.95v-9.9Z" clip-rule="evenodd"/>`
	componentInstanceInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M7.146 1.49a.5.5 0 0 1 .708 0l5.657 5.656a.5.5 0 0 1 0 .708L7.854 13.51a.5.5 0 0 1-.708 0L1.49 7.854a.5.5 0 0 1 0-.708L7.146 1.49ZM7.5 2.55L2.55 7.5l4.95 4.95l4.95-4.95L7.5 2.55Z" clip-rule="evenodd"/>`
	componentNoneInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M7.854 1.49a.5.5 0 0 0-.708 0L1.49 7.146a.5.5 0 0 0 0 .708l2.474 2.474l-2.318 2.318a.5.5 0 0 0 .708.708l2.318-2.318l2.474 2.474a.5.5 0 0 0 .708 0l5.657-5.656a.5.5 0 0 0 0-.708l-2.475-2.474l2.318-2.318a.5.5 0 1 0-.708-.708l-2.318 2.318L7.854 1.49ZM9.62 4.672L7.5 2.55L2.55 7.5l2.122 2.121l4.95-4.95Zm-4.24 5.656L7.5 12.45l4.95-4.95l-2.121-2.121l-4.95 4.95Z" clip-rule="evenodd"/>`
	componentOneInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M4.671 3.146a.5.5 0 0 0 0 .707l2.475 2.475a.5.5 0 0 0 .707 0l2.475-2.475a.5.5 0 0 0 0-.707L7.853.671a.5.5 0 0 0-.707 0L4.671 3.146ZM7.5 5.268L5.732 3.5L7.5 1.732L9.267 3.5L7.5 5.268Zm1.17 1.878a.5.5 0 0 0 0 .707l2.475 2.475a.5.5 0 0 0 .707 0l2.475-2.475a.5.5 0 0 0 0-.707l-2.475-2.475a.5.5 0 0 0-.707 0L8.671 7.146Zm2.83 2.122L9.732 7.5L11.5 5.732L13.268 7.5L11.5 9.268Zm-6.83 2.585a.5.5 0 0 1 0-.707l2.475-2.475a.5.5 0 0 1 .707 0l2.475 2.475a.5.5 0 0 1 0 .707l-2.475 2.475a.5.5 0 0 1-.707 0L4.67 11.854Zm1.061-.353L7.5 13.268L9.267 11.5L7.5 9.732L5.732 11.5ZM.672 7.146a.5.5 0 0 0 0 .708l2.474 2.474a.5.5 0 0 0 .707 0l2.475-2.474a.5.5 0 0 0 0-.708L3.853 4.672a.5.5 0 0 0-.707 0L.671 7.146ZM3.5 9.268L1.732 7.5L3.5 5.732L5.267 7.5L3.5 9.268Z" clip-rule="evenodd"/>`
	componentPlaceholderInnerSVG     = `<g fill="currentColor"><path fill-rule="evenodd" d="M12.803 7.146a.5.5 0 1 1 .707.708a.5.5 0 0 1-.707-.708Zm-.707 1.415a.5.5 0 1 1-.707.707a.5.5 0 0 1 .707-.707Zm-1.414 1.414a.5.5 0 1 1-.707.707a.5.5 0 0 1 .707-.707Zm-1.414 1.414a.5.5 0 1 1-.707.707a.5.5 0 0 1 .707-.707ZM7.146 13.51a.5.5 0 1 1 .708-.707a.5.5 0 0 1-.708.707Z" clip-rule="evenodd"/><path d="M2.904 8.56a.5.5 0 1 0 .707.708a.5.5 0 0 0-.707-.707zm1.414 1.415a.5.5 0 1 0 .707.707a.5.5 0 0 0-.707-.707zm1.414 1.414a.5.5 0 1 0 .707.707a.5.5 0 0 0-.707-.707zM2.197 7.854a.5.5 0 0 0 .064-.63a.506.506 0 0 0-.771-.078a.5.5 0 0 0 .707.708zM3.61 5.732a.5.5 0 1 0-.707.707a.5.5 0 0 0 .707-.707zm1.414-1.414a.5.5 0 1 0-.707.707a.5.5 0 0 0 .707-.707zM6.44 2.904a.5.5 0 1 0-.707.707a.5.5 0 0 0 .707-.707zM7.854 1.49a.5.5 0 1 0-.708.707a.5.5 0 0 0 .708-.707zm1.414 1.414a.5.5 0 1 0-.707.707a.5.5 0 0 0 .707-.707zm1.414 1.414a.5.5 0 1 0-.707.707a.5.5 0 0 0 .707-.707zm1.414 1.414a.5.5 0 1 0-.707.707a.5.5 0 0 0 .707-.707z"/></g>`
	componentTwoInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M3.88 1h7.24c.403 0 .735 0 1.006.022c.281.023.54.072.782.196a2 2 0 0 1 .874.874c.124.243.173.501.196.782c.022.27.022.603.022 1.005v7.241c0 .403 0 .735-.022 1.006c-.023.281-.072.54-.196.782a2 2 0 0 1-.874.874c-.243.124-.501.173-.782.196c-.27.022-.603.022-1.005.022H3.88c-.403 0-.735 0-1.006-.022c-.281-.023-.54-.072-.782-.196a2 2 0 0 1-.874-.874c-.124-.243-.173-.501-.196-.782C1 11.856 1 11.523 1 11.12V3.88c0-.403 0-.735.022-1.006c.023-.281.072-.54.196-.782a2 2 0 0 1 .874-.874c.243-.124.501-.173.782-.196C3.144 1 3.477 1 3.88 1Zm-.924 1.019c-.22.018-.332.05-.41.09a1 1 0 0 0-.437.437c-.04.078-.072.19-.09.41C2 3.18 2 3.472 2 3.9V7h5V2H3.9c-.428 0-.72 0-.944.019ZM7 8H2v3.1c0 .428 0 .72.019.944c.018.22.05.332.09.41a1 1 0 0 0 .437.437c.078.04.19.072.41.09c.225.019.516.019.944.019H7V8Zm1 0h5v3.1c0 .428 0 .72-.019.944c-.018.22-.05.332-.09.41a1 1 0 0 1-.437.437c-.078.04-.19.072-.41.09c-.225.019-.516.019-.944.019H8V8Zm5-1H8V2h3.1c.428 0 .72 0 .944.019c.22.018.332.05.41.09a1 1 0 0 1 .437.437c.04.078.072.19.09.41c.019.225.019.516.019.944V7Z" clip-rule="evenodd"/>`
	containerInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM5 13h5V2H5v11Zm-1 0a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v11Zm9.5-11a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 7.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 9.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 11.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 13.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/>`
	cookieInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M12.07 1.767c.41-.548-.61-1.313-1.02-.765c-.412.548.609 1.313 1.02.765Zm-3.677-.689a7.53 7.53 0 0 0-3.809.631c-1.198.547-2.257 1.444-2.796 2.71c-.412.968-.77 1.862-.787 3.025c-.016 1.154.304 2.525 1.13 4.484c.48 1.137 1.909 1.617 3.008 1.88c1.27.304 2.67.383 3.525.266c.701-.096 1.48-.246 2.135-.777c1.099-.893 2.102-1.902 2.7-2.958c.605-1.067.827-2.247.223-3.386c-.16-.3-.406-.376-.68-.46c-.105-.033-.213-.066-.321-.113c-1.1-.478-1.756-1.03-1.807-1.351c-.048-.3.008-.514.067-.737c.032-.12.064-.243.082-.383c.02-.172.018-.576-.363-.781c-.295-.16-.484-.47-.66-.759c-.043-.071-.085-.141-.128-.207c-.286-.438-.704-1-1.52-1.084ZM4.999 2.62c-1.034.472-1.873 1.212-2.29 2.191c-.41.964-.695 1.699-.708 2.648c-.014.956.25 2.179 1.052 4.081c.339.803 1.569 1.117 2.319 1.296c1.177.282 2.45.345 3.157.248c.684-.093 1.217-.219 1.64-.562c1.06-.862 1.952-1.778 2.46-2.675c.306-.539.885-2.143 0-2.459c-1.875-.668-2.367-1.32-2.639-1.974c-.18-.436-.161-.89-.033-1.334c.02-.066.038-.131.076-.19a2.391 2.391 0 0 1-.518-.517c-.086-.114-.168-.251-.253-.393c-.243-.408-.512-.859-.972-.906A6.531 6.531 0 0 0 5 2.619Zm8.21 1.156c.232-.548-.754-1.139-1.127-.673c-.5.625.82 1.39 1.126.673Zm1.158-1.515c.6.015.788-1.144.205-1.31c-.764-.219-.98 1.291-.205 1.31Zm-.081 2.831c.578.009.76-.709.197-.812c-.738-.135-.946.8-.197.812Zm-7.378-.804c-.439.38-.856-.135-.928-.55c-.112-.64.357-1.168.968-.825c.485.272.319 1.063-.04 1.375ZM4.944 5.882c.464-.464-.35-1.312-.83-.83c-.282.28-.106.602.122.83a.5.5 0 0 0 .708 0Zm2.032.838c.335.67 1.511.072 1.017-.572a.542.542 0 0 0-.373-.206c-.409-.046-.858.35-.644.778Zm.011 2.005c.153.394.523.594.931.438c.355-.135.413-.766.253-1.089c-.149-.299-.532-.396-.835-.28c-.352.136-.48.593-.351.925l.002.006Zm3.4.116c.37-.184 1.074-.013 1.074.471c0 .4-.503.905-.899.97c-.382.065-.858-.306-.797-.705c.048-.312.356-.603.623-.736Zm-2.821 2.002a.523.523 0 0 0-.453.034h-.001c-.205.118-.4.367-.438.607c-.085.52.457.64.852.615c.68-.044.538-1.056.04-1.256Zm-2.687.05a.5.5 0 0 0 .203-.578c-.153-.518-.758-.602-1.117-.185a.5.5 0 0 0 .026.68l.003.004l.003.004c.222.276.638.368.882.075Zm-1.833-3.08c.284.285.554.338.937.246c.624-.15.154-1.816-.937-1.186c-.338.195-.18.76 0 .94Z" clip-rule="evenodd"/>`
	copyInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M1 9.5A1.5 1.5 0 0 0 2.5 11H4v-1H2.5a.5.5 0 0 1-.5-.5v-7a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5V4H5.5A1.5 1.5 0 0 0 4 5.5v7A1.5 1.5 0 0 0 5.5 14h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 12.5 4H11V2.5A1.5 1.5 0 0 0 9.5 1h-7A1.5 1.5 0 0 0 1 2.5v7Zm4-4a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1-.5-.5v-7Z" clip-rule="evenodd"/>`
	cornerBottomLeftInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M9.877 12H11.5a.5.5 0 0 0 0-1H9.9c-1.128 0-1.945 0-2.586-.053c-.637-.052-1.057-.152-1.403-.329a3.5 3.5 0 0 1-1.53-1.529c-.176-.346-.276-.766-.328-1.403C4 7.045 4 6.228 4 5.1V3.5a.5.5 0 0 0-1 0v1.623c0 1.1 0 1.958.056 2.645c.057.698.175 1.265.434 1.775a4.5 4.5 0 0 0 1.967 1.966c.51.26 1.077.378 1.775.435C7.92 12 8.776 12 9.877 12Z" clip-rule="evenodd"/>`
	cornerBottomRightInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M5.123 12H3.5a.5.5 0 0 1 0-1h1.6c1.128 0 1.945 0 2.586-.053c.637-.052 1.057-.152 1.403-.329a3.5 3.5 0 0 0 1.53-1.529c.176-.346.276-.766.328-1.403C11 7.045 11 6.228 11 5.1V3.5a.5.5 0 0 1 1 0v1.623c0 1.1 0 1.958-.056 2.645c-.057.698-.175 1.265-.435 1.775a4.5 4.5 0 0 1-1.966 1.966c-.51.26-1.077.378-1.775.435C7.08 12 6.224 12 5.123 12Z" clip-rule="evenodd"/>`
	cornerTopLeftInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M9.877 3H11.5a.5.5 0 0 1 0 1H9.9c-1.128 0-1.945 0-2.586.053c-.637.052-1.057.152-1.403.328a3.5 3.5 0 0 0-1.53 1.53c-.176.346-.276.766-.328 1.403C4 7.955 4 8.772 4 9.9v1.6a.5.5 0 0 1-1 0V9.877c0-1.1 0-1.958.056-2.645c.057-.698.175-1.265.434-1.775A4.5 4.5 0 0 1 5.457 3.49c.51-.26 1.077-.377 1.775-.434C7.92 3 8.776 3 9.877 3Z" clip-rule="evenodd"/>`
	cornerTopRightInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M5.123 3H3.5a.5.5 0 0 0 0 1h1.6c1.128 0 1.945 0 2.586.053c.637.052 1.057.152 1.403.328a3.5 3.5 0 0 1 1.53 1.53c.176.346.276.766.328 1.403C11 7.955 11 8.772 11 9.9v1.6a.5.5 0 0 0 1 0V9.877c0-1.1 0-1.958-.056-2.645c-.057-.698-.175-1.265-.435-1.775A4.5 4.5 0 0 0 9.543 3.49c-.51-.26-1.077-.377-1.775-.434C7.08 3 6.224 3 5.123 3Z" clip-rule="evenodd"/>`
	cornersInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 2h-.62c-.403 0-.735 0-1.006.022c-.281.023-.54.072-.782.196a2 2 0 0 0-.874.874c-.124.243-.173.501-.196.782C2 4.144 2 4.477 2 4.88v.62a.5.5 0 0 0 1 0v-.6c0-.428 0-.72.019-.944c.018-.22.05-.332.09-.41a1 1 0 0 1 .437-.437c.078-.04.19-.072.41-.09C4.18 3 4.472 3 4.9 3h.6a.5.5 0 0 0 0-1ZM13 9.5a.5.5 0 0 0-1 0v.6c0 .428 0 .72-.019.944c-.018.22-.05.332-.09.41a1 1 0 0 1-.437.437c-.078.04-.19.072-.41.09c-.225.019-.516.019-.944.019h-.6a.5.5 0 0 0 0 1h.62c.403 0 .735 0 1.006-.022c.281-.023.54-.072.782-.196a2 2 0 0 0 .874-.874c.124-.243.173-.501.196-.782c.022-.27.022-.603.022-1.005V9.5ZM2.5 9a.5.5 0 0 1 .5.5v.6c0 .428 0 .72.019.944c.018.22.05.332.09.41a1 1 0 0 0 .437.437c.078.04.19.072.41.09c.225.019.516.019.944.019h.6a.5.5 0 0 1 0 1h-.62c-.403 0-.735 0-1.006-.022c-.281-.023-.54-.072-.782-.196a2 2 0 0 1-.874-.874c-.124-.243-.173-.501-.196-.782C2 10.856 2 10.523 2 10.12V9.5a.5.5 0 0 1 .5-.5Zm7.6-6c.428 0 .72 0 .944.019c.22.018.332.05.41.09a1 1 0 0 1 .437.437c.04.078.072.19.09.41c.019.225.019.516.019.944v.6a.5.5 0 0 0 1 0v-.62c0-.403 0-.735-.022-1.006c-.023-.281-.072-.54-.196-.782a2 2 0 0 0-.874-.874c-.243-.124-.501-.173-.782-.196A13.35 13.35 0 0 0 10.12 2H9.5a.5.5 0 0 0 0 1h.6Z" clip-rule="evenodd"/>`
	countdownTimerInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M13.15 7.5c0-2.835-2.21-5.65-5.65-5.65c-2.778 0-4.151 2.056-4.737 3.15H4.5a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 1 0v1.813C2.705 3.071 4.334.85 7.5.85c4.063 0 6.65 3.335 6.65 6.65c0 3.315-2.587 6.65-6.65 6.65c-1.944 0-3.562-.77-4.715-1.942a6.772 6.772 0 0 1-1.427-2.167a.5.5 0 1 1 .925-.38c.28.681.692 1.314 1.216 1.846c.972.99 2.336 1.643 4.001 1.643c3.44 0 5.65-2.815 5.65-5.65ZM7 10V5h1v5H7Z" clip-rule="evenodd"/>`
	counterClockwiseClockInnerSVG    = `<path fill="currentColor" fill-rule="evenodd" d="M13.15 7.5c0-2.835-2.21-5.65-5.65-5.65c-2.778 0-4.152 2.056-4.737 3.15H4.5a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 1 0v1.813C2.705 3.071 4.334.85 7.5.85c4.063 0 6.65 3.335 6.65 6.65c0 3.315-2.587 6.65-6.65 6.65c-1.944 0-3.562-.77-4.715-1.942a6.772 6.772 0 0 1-1.427-2.167a.5.5 0 1 1 .925-.38c.28.681.692 1.314 1.216 1.846c.972.99 2.336 1.643 4.001 1.643c3.44 0 5.65-2.815 5.65-5.65ZM7.5 4a.5.5 0 0 1 .5.5v2.793l1.854 1.853a.5.5 0 0 1-.708.708l-2-2A.5.5 0 0 1 7 7.5v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`
	cropInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 0a.5.5 0 0 1 .5.5V3h7.5a.5.5 0 0 1 .5.5V11h2.5a.5.5 0 1 1 0 1H12v2.5a.5.5 0 0 1-1 0V12H3.5a.5.5 0 0 1-.5-.5V4H.5a.5.5 0 1 1 0-1H3V.5a.5.5 0 0 1 .5-.5ZM4 4v7h7V4H4Z" clip-rule="evenodd"/>`
	crossCircledInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 1 0 0 11.346a5.673 5.673 0 0 0 0-11.346Zm2.354 3.32a.5.5 0 0 1 0 .707L8.207 7.5l1.647 1.646a.5.5 0 0 1-.708.708L7.5 8.207L5.854 9.854a.5.5 0 0 1-.708-.708L6.793 7.5L5.146 5.854a.5.5 0 0 1 .708-.708L7.5 6.793l1.646-1.647a.5.5 0 0 1 .708 0Z" clip-rule="evenodd"/>`
	crossOneInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M12.854 2.854a.5.5 0 0 0-.708-.708L7.5 6.793L2.854 2.146a.5.5 0 1 0-.708.708L6.793 7.5l-4.647 4.646a.5.5 0 0 0 .708.708L7.5 8.207l4.646 4.647a.5.5 0 0 0 .708-.708L8.207 7.5l4.647-4.646Z" clip-rule="evenodd"/>`
	crossTwoInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M11.782 4.032a.575.575 0 1 0-.813-.814L7.5 6.687L4.032 3.218a.575.575 0 0 0-.814.814L6.687 7.5l-3.469 3.468a.575.575 0 0 0 .814.814L7.5 8.313l3.469 3.469a.575.575 0 0 0 .813-.814L8.313 7.5l3.469-3.468Z" clip-rule="evenodd"/>`
	crosshairOneInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M.877 7.502a6.625 6.625 0 1 1 13.25 0a6.625 6.625 0 0 1-13.25 0ZM1.85 7A5.676 5.676 0 0 1 7 1.849V4.5a.5.5 0 1 0 1 0V1.849A5.677 5.677 0 0 1 13.155 7H10.5a.5.5 0 0 0 0 1h2.655A5.676 5.676 0 0 1 8 13.155V10.5a.5.5 0 0 0-1 0v2.655A5.677 5.677 0 0 1 1.849 8H4.5a.5.5 0 0 0 0-1H1.849Z" clip-rule="evenodd"/>`
	crosshairTwoInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a.5.5 0 0 1 .5.5v1.307A5.624 5.624 0 0 1 13.007 7H14.5a.5.5 0 0 1 0 1h-1.511A5.625 5.625 0 0 1 8 12.989V14.5a.5.5 0 0 1-1 0v-1.493A5.624 5.624 0 0 1 1.807 8H.5a.5.5 0 0 1 0-1h1.289A5.624 5.624 0 0 1 7 1.789V.5a.5.5 0 0 1 .5-.5ZM8 12.032V9.5a.5.5 0 0 0-1 0v2.554A4.675 4.675 0 0 1 2.763 8H5.5a.5.5 0 0 0 0-1H2.742A4.674 4.674 0 0 1 7 2.742V5.5a.5.5 0 0 0 1 0V2.763A4.675 4.675 0 0 1 12.054 7H9.5a.5.5 0 0 0 0 1h2.532A4.675 4.675 0 0 1 8 12.032Z" clip-rule="evenodd"/>`
	crumpledPaperInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M4.61.063a.45.45 0 0 0-.528.27l-.928 2.32l-2.321.93a.45.45 0 0 0-.193.687l1.371 1.829l-.948 3.792a.45.45 0 0 0 .264.524L4.2 11.613l1.425 2.137a.45.45 0 0 0 .498.183l3.5-1a.41.41 0 0 0 .018-.006l3-1a.45.45 0 0 0 .297-.33l1-4.5a.45.45 0 0 0-.169-.457l-1.82-1.365V3a.45.45 0 0 0-.249-.402l-3-1.5a.45.45 0 0 0-.092-.035l-4-1Zm6.822 5.949l1.343 1.007l-2.563 2.135a.45.45 0 0 0-.139.204l-.82 2.457l-1.515-.947a.45.45 0 1 0-.477.764l1.142.713l-2.21.632L4.873 11a.45.45 0 0 0-.2-.165L2.03 9.734l.821-3.284H5a.45.45 0 1 0 0-.9H2.725L1.712 4.2l1.802-.721l2.844.948a.45.45 0 0 0 .493-.146l1.773-2.216l2.426 1.213v1.917l-2.217.887a.45.45 0 0 0-.26.276l-.5 1.5a.45.45 0 1 0 .854.284L9.36 6.84l2.071-.828ZM7.711 1.766L6.345 3.474l-2.25-.75l.677-1.692l2.939.734Zm2.5 10.023l1.905-.635l.659-2.964l-1.889 1.574l-.675 2.025Z" clip-rule="evenodd"/>`
	cubeInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.289.797a.5.5 0 0 1 .422 0l6 2.8A.5.5 0 0 1 14 4.05v6.9a.5.5 0 0 1-.289.453l-6 2.8a.5.5 0 0 1-.422 0l-6-2.8A.5.5 0 0 1 1 10.95v-6.9a.5.5 0 0 1 .289-.453l6-2.8ZM2 4.806L7 6.93v6.034l-5-2.333V4.806Zm6 8.159l5-2.333V4.806L8 6.93v6.034Zm-.5-6.908l4.772-2.028L7.5 1.802L2.728 4.029L7.5 6.057Z" clip-rule="evenodd"/>`
	cursorArrowInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M3.292.049a.5.5 0 0 1 .535.077L12.86 7.95a.5.5 0 0 1-.306.878l-3.334.147l1.931 4.244a.5.5 0 0 1-.247.662l-2.153.98a.5.5 0 0 1-.662-.247L6.153 10.37l-2.29 2.416A.5.5 0 0 1 3 12.44V.504a.5.5 0 0 1 .292-.455ZM4 1.599v9.589l1.938-2.044a.5.5 0 0 1 .818.137l2.035 4.463l1.242-.566l-2.031-4.463a.5.5 0 0 1 .433-.707l2.82-.124L4 1.599Z" clip-rule="evenodd"/>`
	cursorTextInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 1a.5.5 0 0 0 0 1c.922 0 1.54.23 1.92.564c.373.325.58.802.58 1.436v3H5.75a.5.5 0 0 0 0 1H7v3c0 .634-.207 1.11-.58 1.436c-.38.334-.998.564-1.92.564a.5.5 0 0 0 0 1c1.078 0 1.96-.27 2.58-.811c.162-.142.302-.3.42-.47c.118.17.258.328.42.47c.62.541 1.502.811 2.58.811a.5.5 0 0 0 0-1c-.922 0-1.54-.23-1.92-.564C8.206 12.111 8 11.634 8 11V8h1.25a.5.5 0 0 0 0-1H8V4c0-.634.207-1.11.58-1.436C8.96 2.23 9.577 2 10.5 2a.5.5 0 0 0 0-1c-1.078 0-1.96.27-2.58.811c-.162.142-.302.3-.42.47a2.588 2.588 0 0 0-.42-.47C6.46 1.27 5.577 1 4.5 1Z" clip-rule="evenodd"/>`
	dashInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M5 7.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	dashboardInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M2.8 1h-.05c-.229 0-.426 0-.6.041A1.5 1.5 0 0 0 1.04 2.15c-.04.174-.04.37-.04.6v2.5c0 .229 0 .426.041.6A1.5 1.5 0 0 0 2.15 6.96c.174.04.37.04.6.04h2.5c.229 0 .426 0 .6-.041A1.5 1.5 0 0 0 6.96 5.85c.04-.174.04-.37.04-.6v-2.5c0-.229 0-.426-.041-.6A1.5 1.5 0 0 0 5.85 1.04C5.676 1 5.48 1 5.25 1H2.8Zm-.417 1.014c.043-.01.11-.014.417-.014h2.4c.308 0 .374.003.417.014a.5.5 0 0 1 .37.37c.01.042.013.108.013.416v2.4c0 .308-.003.374-.014.417a.5.5 0 0 1-.37.37C5.575 5.996 5.509 6 5.2 6H2.8c-.308 0-.374-.003-.417-.014a.5.5 0 0 1-.37-.37C2.004 5.575 2 5.509 2 5.2V2.8c0-.308.003-.374.014-.417a.5.5 0 0 1 .37-.37ZM9.8 1h-.05c-.229 0-.426 0-.6.041A1.5 1.5 0 0 0 8.04 2.15c-.04.174-.04.37-.04.6v2.5c0 .229 0 .426.041.6A1.5 1.5 0 0 0 9.15 6.96c.174.04.37.04.6.04h2.5c.229 0 .426 0 .6-.041a1.5 1.5 0 0 0 1.11-1.109c.04-.174.04-.37.04-.6v-2.5c0-.229 0-.426-.041-.6a1.5 1.5 0 0 0-1.109-1.11c-.174-.04-.37-.04-.6-.04H9.8Zm-.417 1.014c.043-.01.11-.014.417-.014h2.4c.308 0 .374.003.417.014a.5.5 0 0 1 .37.37c.01.042.013.108.013.416v2.4c0 .308-.004.374-.014.417a.5.5 0 0 1-.37.37c-.042.01-.108.013-.416.013H9.8c-.308 0-.374-.003-.417-.014a.5.5 0 0 1-.37-.37C9.004 5.575 9 5.509 9 5.2V2.8c0-.308.003-.374.014-.417a.5.5 0 0 1 .37-.37ZM2.75 8h2.5c.229 0 .426 0 .6.041A1.5 1.5 0 0 1 6.96 9.15c.04.174.04.37.04.6v2.5c0 .229 0 .426-.041.6a1.5 1.5 0 0 1-1.109 1.11c-.174.04-.37.04-.6.04h-2.5c-.229 0-.426 0-.6-.041a1.5 1.5 0 0 1-1.11-1.109c-.04-.174-.04-.37-.04-.6v-2.5c0-.229 0-.426.041-.6A1.5 1.5 0 0 1 2.15 8.04c.174-.04.37-.04.6-.04Zm.05 1c-.308 0-.374.003-.417.014a.5.5 0 0 0-.37.37C2.004 9.425 2 9.491 2 9.8v2.4c0 .308.003.374.014.417a.5.5 0 0 0 .37.37c.042.01.108.013.416.013h2.4c.308 0 .374-.004.417-.014a.5.5 0 0 0 .37-.37c.01-.042.013-.108.013-.416V9.8c0-.308-.003-.374-.014-.417a.5.5 0 0 0-.37-.37C5.575 9.004 5.509 9 5.2 9H2.8Zm7-1h-.05c-.229 0-.426 0-.6.041A1.5 1.5 0 0 0 8.04 9.15c-.04.174-.04.37-.04.6v2.5c0 .229 0 .426.041.6a1.5 1.5 0 0 0 1.109 1.11c.174.041.371.041.6.041h2.5c.229 0 .426 0 .6-.041a1.5 1.5 0 0 0 1.109-1.109c.041-.174.041-.371.041-.6V9.75c0-.229 0-.426-.041-.6a1.5 1.5 0 0 0-1.109-1.11c-.174-.04-.37-.04-.6-.04H9.8Zm-.417 1.014c.043-.01.11-.014.417-.014h2.4c.308 0 .374.003.417.014a.5.5 0 0 1 .37.37c.01.042.013.108.013.416v2.4c0 .308-.004.374-.014.417a.5.5 0 0 1-.37.37c-.042.01-.108.013-.416.013H9.8c-.308 0-.374-.004-.417-.014a.5.5 0 0 1-.37-.37C9.004 12.575 9 12.509 9 12.2V9.8c0-.308.003-.374.014-.417a.5.5 0 0 1 .37-.37Z" clip-rule="evenodd"/>`
	desktopInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M1 3.25A.25.25 0 0 1 1.25 3h12.5a.25.25 0 0 1 .25.25v7.5a.25.25 0 0 1-.25.25H1.25a.25.25 0 0 1-.25-.25v-7.5ZM1.25 2C.56 2 0 2.56 0 3.25v7.5C0 11.44.56 12 1.25 12h3.823l-.243 1.299a.55.55 0 0 0 .54.651h4.26a.55.55 0 0 0 .54-.651L9.927 12h3.823c.69 0 1.25-.56 1.25-1.25v-7.5C15 2.56 14.44 2 13.75 2H1.25Zm7.76 10H5.99l-.198 1.05h3.416L9.01 12Z" clip-rule="evenodd"/>`
	dimensionsInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M3 2.739a.25.25 0 0 1-.403.197L1.004 1.697a.25.25 0 0 1 0-.394L2.597.063A.25.25 0 0 1 3 .262v.74h6V.26a.25.25 0 0 1 .404-.197l1.592 1.239a.25.25 0 0 1 0 .394l-1.592 1.24A.25.25 0 0 1 9 2.738V2H3v.739ZM9.5 5h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5Zm-7-1A1.5 1.5 0 0 0 1 5.5v7A1.5 1.5 0 0 0 2.5 14h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 9.5 4h-7Zm12.239 2H14v6h.739a.25.25 0 0 1 .197.403l-1.239 1.593a.25.25 0 0 1-.394 0l-1.24-1.593a.25.25 0 0 1 .198-.403H13V6h-.739a.25.25 0 0 1-.197-.403l1.239-1.593a.25.25 0 0 1 .394 0l1.24 1.593a.25.25 0 0 1-.198.403Z" clip-rule="evenodd"/>`
	discInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0Zm6.546 0a.873.873 0 1 1-1.745 0a.873.873 0 0 1 1.745 0Zm.95 0a1.823 1.823 0 1 1-3.645 0a1.823 1.823 0 0 1 3.645 0Z" clip-rule="evenodd"/>`
	discordLogoInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M5.075 1.826a.48.48 0 0 0-.127-.003c-.841.091-2.121.545-2.877.955a.48.48 0 0 0-.132.106c-.314.359-.599.944-.822 1.498C.887 4.95.697 5.55.59 5.984C.236 7.394.043 9.087.017 10.693a.48.48 0 0 0 .056.23c.3.573.947 1.104 1.595 1.492c.655.393 1.42.703 2.036.763a.48.48 0 0 0 .399-.153c.154-.167.416-.557.614-.86c.09-.138.175-.27.241-.375c.662.12 1.492.19 2.542.19c1.048 0 1.878-.07 2.54-.19c.066.106.15.237.24.374c.198.304.46.694.615.861a.48.48 0 0 0 .399.153c.616-.06 1.38-.37 2.035-.763c.648-.388 1.295-.919 1.596-1.492a.48.48 0 0 0 .055-.23c-.025-1.606-.219-3.3-.571-4.71a12.98 12.98 0 0 0-.529-1.601c-.223-.554-.508-1.14-.821-1.498a.48.48 0 0 0-.133-.106c-.755-.41-2.035-.864-2.877-.955a.48.48 0 0 0-.126.003a1.18 1.18 0 0 0-.515.238a2.905 2.905 0 0 0-.794.999A14.046 14.046 0 0 0 7.5 3.02c-.402 0-.774.015-1.117.042a2.905 2.905 0 0 0-.794-.998a1.18 1.18 0 0 0-.514-.238Zm5.943 9.712a23.136 23.136 0 0 0 .433.643c.396-.09.901-.3 1.385-.59c.543-.325.974-.7 1.182-1.017c-.033-1.506-.219-3.07-.54-4.358a12.046 12.046 0 0 0-.488-1.475c-.2-.498-.415-.92-.602-1.162c-.65-.337-1.675-.693-2.343-.79a.603.603 0 0 0-.058.04a1.5 1.5 0 0 0-.226.22a2.52 2.52 0 0 0-.113.145c.305.056.577.123.818.197c.684.21 1.177.5 1.418.821a.48.48 0 1 1-.768.576c-.059-.078-.316-.29-.932-.48c-.595-.182-1.47-.328-2.684-.328c-1.214 0-2.09.146-2.684.329c-.616.19-.873.4-.932.479a.48.48 0 1 1-.768-.576c.241-.322.734-.61 1.418-.82c.24-.075.512-.141.816-.197a2.213 2.213 0 0 0-.114-.146a1.5 1.5 0 0 0-.225-.22a.604.604 0 0 0-.059-.04c-.667.097-1.692.453-2.342.79c-.188.243-.402.664-.603 1.162c-.213.53-.39 1.087-.487 1.475c-.322 1.288-.508 2.852-.54 4.358c.208.318.638.692 1.181 1.018c.485.29.989.5 1.386.589a16.32 16.32 0 0 0 .433-.643c-.785-.279-1.206-.662-1.48-1.072a.48.48 0 0 1 .8-.532c.26.392.944 1.086 4.2 1.086c3.257 0 3.94-.694 4.2-1.086a.48.48 0 0 1 .8.532c-.274.41-.696.794-1.482 1.072ZM4.08 7.012c.244-.262.575-.41.92-.412c.345.002.676.15.92.412c.243.263.38.618.38.988s-.137.725-.38.988c-.244.262-.575.41-.92.412a1.263 1.263 0 0 1-.92-.412A1.453 1.453 0 0 1 3.7 8c0-.37.137-.725.38-.988ZM10 6.6c-.345.002-.676.15-.92.412c-.243.263-.38.618-.38.988s.137.725.38.988c.244.262.575.41.92.412c.345-.002.676-.15.92-.412c.243-.263.38-.618.38-.988s-.137-.725-.38-.988A1.263 1.263 0 0 0 10 6.6Z" clip-rule="evenodd"/>`
	dividerHorizontalInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M2 7.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	dividerVerticalInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 2a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-1 0v-10a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`
	dotInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 9.125a1.625 1.625 0 1 0 0-3.25a1.625 1.625 0 0 0 0 3.25Zm0 1a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Z" clip-rule="evenodd"/>`
	dotFilledInnerSVG                = `<path fill="currentColor" d="M9.875 7.5a2.375 2.375 0 1 1-4.75 0a2.375 2.375 0 0 1 4.75 0Z"/>`
	dotsHorizontalInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M3.625 7.5a1.125 1.125 0 1 1-2.25 0a1.125 1.125 0 0 1 2.25 0Zm5 0a1.125 1.125 0 1 1-2.25 0a1.125 1.125 0 0 1 2.25 0ZM12.5 8.625a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25Z" clip-rule="evenodd"/>`
	dotsVerticalInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M8.625 2.5a1.125 1.125 0 1 1-2.25 0a1.125 1.125 0 0 1 2.25 0Zm0 5a1.125 1.125 0 1 1-2.25 0a1.125 1.125 0 0 1 2.25 0ZM7.5 13.625a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25Z" clip-rule="evenodd"/>`
	doubleArrowDownInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M3.854 2.146a.5.5 0 1 0-.708.708l4 4a.5.5 0 0 0 .708 0l4-4a.5.5 0 0 0-.708-.708L7.5 5.793L3.854 2.146Zm0 6a.5.5 0 1 0-.708.708l4 4a.5.5 0 0 0 .708 0l4-4a.5.5 0 0 0-.708-.708L7.5 11.793L3.854 8.146Z" clip-rule="evenodd"/>`
	doubleArrowLeftInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M6.854 3.854a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L3.207 7.5l3.647-3.646Zm6 0a.5.5 0 0 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L9.207 7.5l3.647-3.646Z" clip-rule="evenodd"/>`
	doubleArrowRightInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M2.146 11.146a.5.5 0 0 0 .708.708l4-4a.5.5 0 0 0 0-.708l-4-4a.5.5 0 1 0-.708.708L5.793 7.5l-3.647 3.646Zm6 0a.5.5 0 0 0 .708.708l4-4a.5.5 0 0 0 0-.708l-4-4a.5.5 0 1 0-.708.708L11.793 7.5l-3.647 3.646Z" clip-rule="evenodd"/>`
	doubleArrowUpInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M11.146 6.854a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 1 0 .708.708L7.5 3.207l3.646 3.647Zm0 6a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 0 0 .708.708L7.5 9.207l3.646 3.647Z" clip-rule="evenodd"/>`
	downloadInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 1.05a.45.45 0 0 1 .45.45v6.914l2.232-2.232a.45.45 0 1 1 .636.636l-3 3a.45.45 0 0 1-.636 0l-3-3a.45.45 0 1 1 .636-.636L7.05 8.414V1.5a.45.45 0 0 1 .45-.45ZM2.5 10a.5.5 0 0 1 .5.5V12c0 .554.446 1 .996 1h7.005A.999.999 0 0 0 12 12v-1.5a.5.5 0 0 1 1 0V12a2 2 0 0 1-1.999 2H3.996A1.997 1.997 0 0 1 2 12v-1.5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`
	dragHandleDotsOneInnerSVG        = `<g fill="currentColor"><circle cx="4.5" cy="2.5" r=".6"/><circle cx="4.5" cy="4.5" r=".6"/><circle cx="4.5" cy="6.499" r=".6"/><circle cx="4.5" cy="8.499" r=".6"/><circle cx="4.5" cy="10.498" r=".6"/><circle cx="4.5" cy="12.498" r=".6"/><circle cx="6.5" cy="2.5" r=".6"/><circle cx="6.5" cy="4.5" r=".6"/><circle cx="6.5" cy="6.499" r=".6"/><circle cx="6.5" cy="8.499" r=".6"/><circle cx="6.5" cy="10.498" r=".6"/><circle cx="6.5" cy="12.498" r=".6"/><circle cx="8.499" cy="2.5" r=".6"/><circle cx="8.499" cy="4.5" r=".6"/><circle cx="8.499" cy="6.499" r=".6"/><circle cx="8.499" cy="8.499" r=".6"/><circle cx="8.499" cy="10.498" r=".6"/><circle cx="8.499" cy="12.498" r=".6"/><circle cx="10.499" cy="2.5" r=".6"/><circle cx="10.499" cy="4.5" r=".6"/><circle cx="10.499" cy="6.499" r=".6"/><circle cx="10.499" cy="8.499" r=".6"/><circle cx="10.499" cy="10.498" r=".6"/><circle cx="10.499" cy="12.498" r=".6"/></g>`
	dragHandleDotsTwoInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 4.625a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25Zm4 0a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25ZM10.625 7.5a1.125 1.125 0 1 1-2.25 0a1.125 1.125 0 0 1 2.25 0ZM5.5 8.625a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25Zm5.125 2.875a1.125 1.125 0 1 1-2.25 0a1.125 1.125 0 0 1 2.25 0ZM5.5 12.625a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25Z" clip-rule="evenodd"/>`
	dragHandleHorizontalInnerSVG     = `<path fill="currentColor" fill-rule="evenodd" d="M2.5 4.1a.4.4 0 1 0 0 .8h10a.4.4 0 0 0 0-.8h-10Zm0 2a.4.4 0 1 0 0 .8h10a.4.4 0 0 0 0-.8h-10Zm-.4 2.4c0-.22.18-.4.4-.4h10a.4.4 0 0 1 0 .8h-10a.4.4 0 0 1-.4-.4Zm.4 1.6a.4.4 0 0 0 0 .8h10a.4.4 0 0 0 0-.8h-10Z" clip-rule="evenodd"/>`
	dragHandleVerticalInnerSVG       = `<path fill="currentColor" fill-rule="evenodd" d="M4.1 12.5a.4.4 0 0 0 .8 0v-10a.4.4 0 0 0-.8 0v10Zm2 0a.4.4 0 0 0 .8 0v-10a.4.4 0 0 0-.8 0v10Zm2.4.4a.4.4 0 0 1-.4-.4v-10a.4.4 0 1 1 .8 0v10a.4.4 0 0 1-.4.4Zm1.6-.4a.4.4 0 0 0 .8 0v-10a.4.4 0 0 0-.8 0v10Z" clip-rule="evenodd"/>`
	drawingPinInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M10.329 1.136a.5.5 0 0 0-.708.707l.653.653l-4.848 3.637l-1.108-1.108a.5.5 0 0 0-.707.707l1.414 1.414l1.06 1.061l-3.27 3.27a.5.5 0 1 0 .708.708l3.27-3.27l1.06 1.06l1.415 1.414a.5.5 0 1 0 .707-.707L8.867 9.574l3.637-4.848l.653.653a.5.5 0 1 0 .707-.707l-1.06-1.061l-1.415-1.414l-1.06-1.06Zm-4.19 5.711l4.85-3.637l.8.801l-3.636 4.85L6.14 6.846Z" clip-rule="evenodd"/>`
	drawingPinFilledInnerSVG         = `<defs><path id="radixIconsDrawingPinFilled0" d="M9.621 1.136a.5.5 0 0 1 .707 0l1.061 1.06l1.414 1.415l1.06 1.06a.5.5 0 1 1-.706.708l-.653-.653l-3.637 4.848l1.108 1.108a.5.5 0 0 1-.707.707L7.854 9.975l-1.061-1.06l-3.27 3.27a.5.5 0 1 1-.708-.708l3.27-3.27l-1.06-1.06l-1.414-1.415a.5.5 0 1 1 .707-.707l1.108 1.108l4.848-3.637l-.653-.653a.5.5 0 0 1 0-.707Z"/></defs><g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><use href="#radixIconsDrawingPinFilled0"/><use href="#radixIconsDrawingPinFilled0"/></g>`
	dropdownMenuInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.1a.4.4 0 1 0 0 .8h7a.4.4 0 0 0 0-.8h-7Zm0 2a.4.4 0 1 0 0 .8h7a.4.4 0 0 0 0-.8h-7Zm-.4 2.4c0-.22.18-.4.4-.4h7a.4.4 0 0 1 0 .8h-7a.4.4 0 0 1-.4-.4Zm.4 1.6a.4.4 0 1 0 0 .8h7a.4.4 0 0 0 0-.8h-7Zm-.4 2.4c0-.22.18-.4.4-.4h7a.4.4 0 0 1 0 .8h-7a.4.4 0 0 1-.4-.4ZM2.5 9.25L5 6H0l2.5 3.25Z" clip-rule="evenodd"/>`
	enterInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 1a.5.5 0 0 0 0 1H12v11H4.5a.5.5 0 0 0 0 1H12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H4.5Zm2.104 3.896a.5.5 0 1 0-.708.708L7.293 7H.5a.5.5 0 0 0 0 1h6.793L5.896 9.396a.5.5 0 0 0 .708.708l2.25-2.25a.5.5 0 0 0 0-.708l-2.25-2.25Z" clip-rule="evenodd"/>`
	enterFullScreenInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M2 2.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1H3v2.5a.5.5 0 0 1-1 0v-3Zm7 0a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0V3H9.5a.5.5 0 0 1-.5-.5ZM2.5 9a.5.5 0 0 1 .5.5V12h2.5a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5Zm10 0a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1H12V9.5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`
	envelopeClosedInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M1 2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H1Zm0 1h13v.925a.448.448 0 0 0-.241.07L7.5 7.967L1.241 3.995A.448.448 0 0 0 1 3.925V3Zm0 1.908V12h13V4.908L7.741 8.88a.45.45 0 0 1-.482 0L1 4.908Z" clip-rule="evenodd"/>`
	envelopeOpenInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M7.947.165a1 1 0 0 0-.894 0l-6.5 3.25A1 1 0 0 0 0 4.309V12a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V4.309a1 1 0 0 0-.553-.894l-6.5-3.25Zm5.622 3.928L7.5 1.06L1.431 4.093L7.5 7.291l6.069-3.198ZM1 4.883V12h13V4.884L7.71 8.198a.45.45 0 0 1-.42 0L1 4.884Z" clip-rule="evenodd"/>`
	eraserInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M8.36.73a.5.5 0 0 1 .708 0l5.203 5.202a.5.5 0 0 1 0 .707l-5.316 5.316l-1.608 1.609a1.5 1.5 0 0 1-2.122 0l-3.789-3.79a1.5 1.5 0 0 1 0-2.12l1.609-1.61L8.36.73Zm.354 1.06L4.106 6.398l4.496 4.496l4.608-4.608L8.714 1.79Zm-.82 9.811L3.398 7.106L2.143 8.36a.5.5 0 0 0 0 .708l3.79 3.789a.5.5 0 0 0 .706 0l1.255-1.255Z" clip-rule="evenodd"/>`
	exclamationTriangleInnerSVG      = `<path fill="currentColor" fill-rule="evenodd" d="M8.445.609a1.1 1.1 0 0 0-1.89 0L.161 11.337A1.1 1.1 0 0 0 1.106 13h12.788a1.1 1.1 0 0 0 .945-1.663L8.445.609Zm-1.03.512a.1.1 0 0 1 .17 0l6.395 10.728a.1.1 0 0 1-.086.151H1.106a.1.1 0 0 1-.086-.151L7.414 1.12Zm-.588 3.365a.674.674 0 1 1 1.346 0L8.02 8.487a.52.52 0 0 1-1.038 0l-.154-4Zm1.423 5.99a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z" clip-rule="evenodd"/>`
	exitInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M3 1a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h7.5a.5.5 0 0 0 0-1H3V2h7.5a.5.5 0 0 0 0-1H3Zm9.604 3.896a.5.5 0 0 0-.708.708L13.293 7H6.5a.5.5 0 0 0 0 1h6.793l-1.397 1.396a.5.5 0 0 0 .708.708l2.25-2.25a.5.5 0 0 0 0-.708l-2.25-2.25Z" clip-rule="evenodd"/>`
	exitFullScreenInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 2a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1H5V2.5a.5.5 0 0 1 .5-.5Zm4 0a.5.5 0 0 1 .5.5V5h2.5a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5ZM2 9.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0V10H2.5a.5.5 0 0 1-.5-.5Zm7 0a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1H10v2.5a.5.5 0 0 1-1 0v-3Z" clip-rule="evenodd"/>`
	externalLinkInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M3 2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V8.5a.5.5 0 0 0-1 0V12H3V3h3.5a.5.5 0 0 0 0-1H3Zm9.854.146a.5.5 0 0 1 .146.351V5.5a.5.5 0 0 1-1 0V3.707L6.854 8.854a.5.5 0 1 1-.708-.708L11.293 3H9.5a.5.5 0 0 1 0-1h3a.499.499 0 0 1 .354.146Z" clip-rule="evenodd"/>`
	eyeClosedInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M14.765 6.076a.5.5 0 0 1 .159.689a9.519 9.519 0 0 1-1.554 1.898l1.201 1.201a.5.5 0 0 1-.707.707l-1.263-1.263a8.472 8.472 0 0 1-2.667 1.343l.449 1.677a.5.5 0 0 1-.966.258l-.458-1.709a8.666 8.666 0 0 1-2.918 0l-.458 1.71a.5.5 0 1 1-.966-.26l.45-1.676a8.473 8.473 0 0 1-2.668-1.343l-1.263 1.263a.5.5 0 0 1-.707-.707l1.2-1.201A9.521 9.521 0 0 1 .077 6.765a.5.5 0 1 1 .848-.53a8.425 8.425 0 0 0 1.77 2.034A7.462 7.462 0 0 0 7.5 9.999c2.808 0 5.156-1.493 6.576-3.764a.5.5 0 0 1 .689-.159Z" clip-rule="evenodd"/>`
	eyeNoneInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M13.354 2.354a.5.5 0 0 0-.708-.708L10.683 3.61A8.475 8.475 0 0 0 7.5 3C4.308 3 1.656 4.706.076 7.235a.5.5 0 0 0 0 .53c.827 1.323 1.947 2.421 3.285 3.167l-1.715 1.714a.5.5 0 0 0 .708.708l1.963-1.964c.976.393 2.045.61 3.183.61c3.192 0 5.844-1.706 7.424-4.235a.5.5 0 0 0 0-.53c-.827-1.323-1.947-2.421-3.285-3.167l1.715-1.714Zm-3.45 2.035A7.517 7.517 0 0 0 7.5 4C4.803 4 2.53 5.378 1.096 7.5c.777 1.15 1.8 2.081 3.004 2.693L9.904 4.39ZM5.096 10.61L10.9 4.807c1.204.612 2.227 1.543 3.004 2.693C12.47 9.622 10.197 11 7.5 11a7.518 7.518 0 0 1-2.404-.389Z" clip-rule="evenodd"/>`
	eyeOpenInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 11c-2.697 0-4.97-1.378-6.404-3.5C2.53 5.378 4.803 4 7.5 4s4.97 1.378 6.404 3.5C12.47 9.622 10.197 11 7.5 11Zm0-8C4.308 3 1.656 4.706.076 7.235a.5.5 0 0 0 0 .53C1.656 10.294 4.308 12 7.5 12s5.844-1.706 7.424-4.235a.5.5 0 0 0 0-.53C13.344 4.706 10.692 3 7.5 3Zm0 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/>`
	faceInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0Zm3.21 1.714a.5.5 0 1 0-.82.572A3.996 3.996 0 0 0 7.5 11.5c1.36 0 2.56-.679 3.283-1.714a.5.5 0 0 0-.82-.572A2.996 2.996 0 0 1 7.5 10.5a2.996 2.996 0 0 1-2.463-1.286Zm.338-2.364a.875.875 0 1 0 0-1.75a.875.875 0 0 0 0 1.75Zm5.125-.875a.875.875 0 1 1-1.75 0a.875.875 0 0 1 1.75 0Z" clip-rule="evenodd"/>`
	figmaLogoInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M7 2.05H5.525a1.475 1.475 0 0 0 0 2.95H7V2.05Zm0-1h2.475a2.475 2.475 0 0 1 1.492 4.45A2.475 2.475 0 0 1 8 9.463v1.962A2.475 2.475 0 1 1 4.033 9.45a2.471 2.471 0 0 1-.983-1.975c0-.807.386-1.523.983-1.975a2.475 2.475 0 0 1 1.492-4.45H7Zm1 1V5h1.475a1.475 1.475 0 1 0 0-2.95H8Zm-2.475 6.9H7V6H5.525a1.475 1.475 0 0 0-.006 2.95h.006ZM4.05 11.425c0-.813.657-1.472 1.47-1.475H7v1.475a1.475 1.475 0 0 1-2.95 0ZM8 7.472a1.475 1.475 0 1 1 0 .006v-.006Z" clip-rule="evenodd"/>`
	fileInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 2a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V6H8.5a.5.5 0 0 1-.5-.5V2H3.5Zm5.5.707L11.293 5H9V2.707ZM2 2.5A1.5 1.5 0 0 1 3.5 1h5a.5.5 0 0 1 .354.146l4 4A.5.5 0 0 1 13 5.5v7a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 12.5v-10Z" clip-rule="evenodd"/>`
	fileMinusInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M3 2.5a.5.5 0 0 1 .5-.5h5.793L12 4.707V12.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-10ZM3.5 1A1.5 1.5 0 0 0 2 2.5v10A1.5 1.5 0 0 0 3.5 14h8a1.5 1.5 0 0 0 1.5-1.5V4.604a.75.75 0 0 0-.22-.53L9.854 1.145A.5.5 0 0 0 9.5 1h-6Zm1.75 6a.5.5 0 0 0 0 1h4.5a.5.5 0 0 0 0-1h-4.5Z" clip-rule="evenodd"/>`
	filePlusInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 2a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V4.707L9.293 2H3.5ZM2 2.5A1.5 1.5 0 0 1 3.5 1h6a.5.5 0 0 1 .354.146l2.926 2.927c.141.14.22.332.22.53V12.5a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 12.5v-10Zm2.75 5a.5.5 0 0 1 .5-.5H7V5.25a.5.5 0 0 1 1 0V7h1.75a.5.5 0 0 1 0 1H8v1.75a.5.5 0 0 1-1 0V8H5.25a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	fileTextInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3 2.5a.5.5 0 0 1 .5-.5h5.586a.5.5 0 0 1 .353.146l2.415 2.415a.5.5 0 0 1 .146.353V12.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-10ZM3.5 1A1.5 1.5 0 0 0 2 2.5v10A1.5 1.5 0 0 0 3.5 14h8a1.5 1.5 0 0 0 1.5-1.5V4.914a1.5 1.5 0 0 0-.44-1.06l-2.414-2.415A1.5 1.5 0 0 0 9.086 1H3.5Zm1 3a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Zm0 3a.5.5 0 0 0 0 1h6a.5.5 0 0 0 0-1h-6Zm0 3a.5.5 0 0 0 0 1h6a.5.5 0 0 0 0-1h-6Z" clip-rule="evenodd"/>`
	fontBoldInnerSVG                 = `<path fill="currentColor" d="M5.105 12c-.397 0-.681-.088-.853-.264c-.168-.18-.252-.465-.252-.853V4.117c0-.397.086-.681.258-.853c.176-.176.458-.264.847-.264H9.03c1.108 0 2.025.982 2.025 2.185c0 .9-.45 1.634-1.35 2.051c1.162.213 1.814 1.392 1.814 2.245c0 1.031-.528 2.519-2.24 2.519H5.104Zm3.274-3.997H5.8v2.628h2.579c.521 0 1.25-.51 1.25-1.332c0-.823-.729-1.296-1.25-1.296ZM5.8 4.37v2.327h2.38c.36 0 1.097-.337 1.097-1.196c0-.86-.797-1.131-1.097-1.131H5.8Z"/>`
	fontFamilyInnerSVG               = `<path fill="currentColor" d="M2.5 4.5C2.5 3.099 3.599 2 5 2h7.499a.5.5 0 0 1 .001 1H8.692l-.287.855A1887.39 1887.39 0 0 1 7.343 7H8.5a.5.5 0 0 1 0 1H7.004a286.12 286.12 0 0 1-1.046 3.039c-.322.9-.75 1.447-1.29 1.739c-.505.273-1.026.272-1.384.272H3.25a.55.55 0 0 1 0-1.1c.392 0 .654-.01.894-.14c.22-.119.511-.395.778-1.142c.185-.517.532-1.527.92-2.668H4.5a.5.5 0 0 1 0-1h1.682a1350.118 1350.118 0 0 0 1.18-3.496L7.533 3H5c-.849 0-1.5.651-1.5 1.5a.5.5 0 0 1-1 0Z"/>`
	fontItalicInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M5.675 3.5a.45.45 0 0 1 .45-.45h4.5a.45.45 0 1 1 0 .9h-1.62l-1.774 7.1h1.644a.45.45 0 0 1 0 .9h-4.5a.45.45 0 1 1 0-.9h1.619l1.775-7.1H6.125a.45.45 0 0 1-.45-.45Z" clip-rule="evenodd"/>`
	fontRomanInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M4.8 3.5a.45.45 0 0 1 .45-.45h4.5a.45.45 0 0 1 0 .9H8.1v7.1h1.65a.45.45 0 0 1 0 .9h-4.5a.45.45 0 1 1 0-.9H6.9v-7.1H5.25a.45.45 0 0 1-.45-.45Z" clip-rule="evenodd"/>`
	fontSizeInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M2.782 2.217a.4.4 0 0 0-.565 0l-2 2a.4.4 0 0 0 .565.566L2.1 3.466v8.068L.782 10.217a.4.4 0 0 0-.565.566l2 2a.4.4 0 0 0 .565 0l2-2a.4.4 0 0 0-.565-.566l-1.318 1.317V3.466l1.318 1.317a.4.4 0 0 0 .565-.566l-2-2Zm7.718.533a.5.5 0 0 1 .47.33l3 8.32a.5.5 0 0 1-.94.34l-.982-2.724H8.952l-.982 2.723a.5.5 0 0 1-.94-.34l3-8.319a.5.5 0 0 1 .47-.33Zm0 1.974l1.241 3.442H9.26l1.24-3.442Z" clip-rule="evenodd"/>`
	fontStyleInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M1.038 9.98c-.103-.322.078-.603.362-.731c.284-.129.695-.115.87.205c.584 1.06 1.376 1.274 2.217 1.274c1.04 0 1.81-.466 1.81-1.274c0-.65-.385-.99-1.426-1.311l-.712-.221c-1.514-.473-2.27-1.23-2.27-2.466c0-1.659 1.387-2.769 3.354-2.769c1.674 0 2.731.648 3.195 1.624c.133.278.138.602-.21.88s-.704.157-.995-.153c-.76-.811-1.238-.988-1.977-.988c-1.116 0-1.709.586-1.709 1.23c0 .586.416.952 1.4 1.254l.732.227c1.55.473 2.295 1.199 2.295 2.41c0 1.601-1.28 2.92-3.513 2.92c-1.595 0-3.061-.978-3.423-2.11Zm10.811-1.2c-1.188-.385-1.684-.919-1.684-1.792c0-1.12.999-1.942 2.448-1.942c1.242 0 2.05.587 2.365 1.589c.066.211-.019.345-.23.414c-.209.068-.43.05-.51-.153c-.302-.773-.886-1.133-1.638-1.133c-.953 0-1.586.489-1.586 1.153c0 .535.332.834 1.233 1.128l.588.189c1.227.397 1.717.905 1.717 1.785c0 1.18-1.071 2.026-2.56 2.026c-1.348 0-2.336-.763-2.572-1.708c-.055-.217-.008-.307.28-.374c.289-.067.371-.063.472.175c.284.674.981 1.19 1.86 1.19c.96 0 1.651-.547 1.651-1.264c0-.527-.287-.775-1.246-1.088l-.588-.195Z" clip-rule="evenodd"/>`
	frameInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M11 1.5a.5.5 0 0 0-1 0V4H5V1.5a.5.5 0 0 0-1 0V4H1.5a.5.5 0 0 0 0 1H4v5H1.5a.5.5 0 0 0 0 1H4v2.5a.5.5 0 0 0 1 0V11h5v2.5a.5.5 0 0 0 1 0V11h2.5a.5.5 0 0 0 0-1H11V5h2.5a.5.5 0 0 0 0-1H11V1.5ZM10 10V5H5v5h5Z" clip-rule="evenodd"/>`
	framerLogoInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M3.382 1.296A.5.5 0 0 1 3.84.997h7.66a.5.5 0 0 1 .5.5V5.5a.5.5 0 0 1-.5.5H8.635l2.894 3.162a.5.5 0 0 1-.369.838H8v3.5a.5.5 0 0 1-.854.354l-4-4A.5.5 0 0 1 3 9.5v-4a.5.5 0 0 1 .5-.5h2.865L3.471 1.835a.5.5 0 0 1-.089-.54ZM7.72 5L4.975 1.997H11V5H7.72Zm-.44 1H4v3h6.025L7.28 6Zm-2.573 4L7 12.293V10H4.707Z" clip-rule="evenodd"/>`
	gearInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.07.65a.85.85 0 0 0-.828.662l-.238 1.05c-.38.11-.74.262-1.08.448l-.91-.574a.85.85 0 0 0-1.055.118l-.606.606a.85.85 0 0 0-.118 1.054l.574.912c-.186.338-.337.7-.447 1.079l-1.05.238a.85.85 0 0 0-.662.829v.857a.85.85 0 0 0 .662.829l1.05.238c.11.379.261.74.448 1.08l-.575.91a.85.85 0 0 0 .118 1.055l.607.606a.85.85 0 0 0 1.054.118l.911-.574c.339.186.7.337 1.079.447l.238 1.05a.85.85 0 0 0 .829.662h.857a.85.85 0 0 0 .829-.662l.238-1.05c.38-.11.74-.26 1.08-.447l.911.574a.85.85 0 0 0 1.054-.118l.606-.606a.85.85 0 0 0 .118-1.054l-.574-.911c.187-.34.338-.7.448-1.08l1.05-.238a.85.85 0 0 0 .662-.829v-.857a.85.85 0 0 0-.662-.83l-1.05-.237c-.11-.38-.26-.74-.447-1.08l.574-.91a.85.85 0 0 0-.118-1.055l-.606-.606a.85.85 0 0 0-1.055-.118l-.91.574a5.323 5.323 0 0 0-1.08-.448l-.239-1.05A.85.85 0 0 0 7.928.65h-.857ZM4.92 3.813a4.476 4.476 0 0 1 1.795-.745L7.071 1.5h.857l.356 1.568a4.48 4.48 0 0 1 1.795.744l1.36-.857l.607.606l-.858 1.36c.37.527.628 1.136.744 1.795l1.568.356v.857l-1.568.355a4.475 4.475 0 0 1-.744 1.796l.857 1.36l-.606.606l-1.36-.857a4.476 4.476 0 0 1-1.795.743L7.928 13.5h-.857l-.356-1.568a4.475 4.475 0 0 1-1.794-.744l-1.36.858l-.607-.606l.858-1.36a4.476 4.476 0 0 1-.744-1.796L1.5 7.93v-.857l1.568-.356a4.476 4.476 0 0 1 .744-1.794L2.954 3.56l.606-.606l1.36.858ZM9.026 7.5a1.525 1.525 0 1 1-3.05 0a1.525 1.525 0 0 1 3.05 0Zm.9 0a2.425 2.425 0 1 1-4.85 0a2.425 2.425 0 0 1 4.85 0Z" clip-rule="evenodd"/>`
	githubLogoInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.25a7.25 7.25 0 0 0-2.292 14.13c.363.066.495-.158.495-.35c0-.172-.006-.628-.01-1.233c-2.016.438-2.442-.972-2.442-.972c-.33-.838-.805-1.06-.805-1.06c-.658-.45.05-.441.05-.441c.728.051 1.11.747 1.11.747c.647 1.108 1.697.788 2.11.602c.066-.468.254-.788.46-.969c-1.61-.183-3.302-.805-3.302-3.583a2.8 2.8 0 0 1 .747-1.945c-.075-.184-.324-.92.07-1.92c0 0 .61-.194 1.994.744A6.963 6.963 0 0 1 7.5 3.756A6.97 6.97 0 0 1 9.315 4c1.384-.938 1.992-.743 1.992-.743c.396.998.147 1.735.072 1.919c.465.507.745 1.153.745 1.945c0 2.785-1.695 3.398-3.31 3.577c.26.224.492.667.492 1.343c0 .97-.009 1.751-.009 1.989c0 .194.131.42.499.349A7.25 7.25 0 0 0 7.499.25Z" clip-rule="evenodd"/>`
	globeInnerSVG                    = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.5 1.8a5.7 5.7 0 1 0 0 11.4a5.7 5.7 0 0 0 0-11.4ZM.9 7.5a6.6 6.6 0 1 1 13.2 0a6.6 6.6 0 0 1-13.2 0Z"/><path d="M13.5 7.9h-12v-.8h12v.8Z"/><path d="M7.1 13.5v-12h.8v12h-.8zm3.275-6c0-2.173-.781-4.322-2.313-5.743l.476-.514c1.702 1.58 2.537 3.93 2.537 6.257c0 2.327-.835 4.678-2.537 6.257l-.476-.514c1.532-1.42 2.313-3.57 2.313-5.743zM4 7.5c0-2.324.808-4.673 2.458-6.253l.484.506C5.458 3.173 4.7 5.323 4.7 7.5c0 2.176.758 4.327 2.242 5.747l-.484.506C4.808 12.173 4 9.823 4 7.5z"/><path d="M7.5 3.958c2.17 0 4.375.401 5.87 1.236a.35.35 0 1 1-.34.612c-1.35-.754-3.422-1.148-5.53-1.148s-4.18.394-5.53 1.148a.35.35 0 1 1-.34-.612c1.495-.835 3.7-1.236 5.87-1.236zm0 6.892c2.17 0 4.375-.401 5.87-1.236a.35.35 0 1 0-.34-.612c-1.35.754-3.422 1.148-5.53 1.148s-4.18-.394-5.53-1.148a.35.35 0 1 0-.34.611c1.495.836 3.7 1.237 5.87 1.237z"/></g>`
	gridInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M12.5 2H8v5h5V2.5a.5.5 0 0 0-.5-.5Zm.5 6H8v5h4.5a.5.5 0 0 0 .5-.5V8ZM7 7V2H2.5a.5.5 0 0 0-.5.5V7h5ZM2 8v4.5a.5.5 0 0 0 .5.5H7V8H2Zm.5-7A1.5 1.5 0 0 0 1 2.5v10A1.5 1.5 0 0 0 2.5 14h10a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 12.5 1h-10Z" clip-rule="evenodd"/>`
	groupInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M1.45.95a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-1.5h1.5a.5.5 0 0 0 0-1h-2Zm4.5 0a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Zm-.5 12.5a.5.5 0 0 1 .5-.5h3a.5.5 0 1 1 0 1h-3a.5.5 0 0 1-.5-.5Zm-3.5-7.5a.5.5 0 0 0-1 0v3a.5.5 0 0 0 1 0v-3Zm11.5-.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .5-.5Zm-2-4.5a.5.5 0 1 0 0 1h1.5v1.5a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.501-.5H11.45Zm-10 10a.5.5 0 0 1 .5.5v1.5h1.5a.5.5 0 1 1 0 1h-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5Zm12.5.5a.5.5 0 0 0-1 0v1.5h-1.5a.5.5 0 1 0 0 1h2a.5.5 0 0 0 .5-.5v-2Z" clip-rule="evenodd"/>`
	halfOneInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877Zm-.5.972A5.673 5.673 0 0 0 7 13.15V1.849ZM8 13.15A5.673 5.673 0 0 0 8 1.849V13.15Z" clip-rule="evenodd"/>`
	halfTwoInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877Zm0 .95a5.673 5.673 0 0 0 0 11.346V1.827Z" clip-rule="evenodd"/>`
	hamburgerMenuInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 3a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1h-12ZM1 7.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 0 1h-12a.5.5 0 0 1-.5-.5Zm0 4a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 0 1h-12a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	handInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M6.811 1.647c-.19 1.232-.128 2.238-.043 3.658c.016.26.033.532.049.823a.5.5 0 0 1-.992.115l-.077-.436c-.104-.642-.264-1.552-.57-2.36c-.316-.834-.73-1.406-1.244-1.587a.608.608 0 0 0-.675.15a.82.82 0 0 0-.196.744c.152.59.383 1.163.635 1.764l.068.162c.23.545.47 1.117.65 1.71v.002c.146.484.24.867.315 1.17l.08.322A.5.5 0 0 1 4 8.387a9.518 9.518 0 0 1-.68-.669a57.738 57.738 0 0 1-.122-.13c-.15-.16-.28-.3-.414-.428c-.323-.309-.599-.48-.977-.516h-.003c-.242-.025-.389.02-.485.082c-.099.063-.198.178-.28.386c-.061.173-.023.518.155.857l.004.008c.204.408.548.839.984 1.316c.216.236.448.475.688.724l.01.011c.236.244.48.497.718.755c.832.903 1.68 1.97 1.866 3.217h6.046c.055-1.565.367-2.732.778-3.865c.158-.435.325-.85.496-1.277c.327-.817.668-1.669.988-2.75c.328-1.11.239-1.738.09-2.068c-.144-.321-.37-.422-.48-.444c-.243-.048-.343.013-.411.083c-.103.104-.2.304-.278.624a8.82 8.82 0 0 0-.159 1.022l-.003.028c-.038.32-.08.677-.163.948a.539.539 0 0 1-.17.274a.502.502 0 0 1-.366.118c-.289-.024-.46-.272-.466-.545c-.009-.393.023-.79.042-1.182c.036-.73.07-1.432.018-2.132V2.83c-.033-.499-.376-.827-.724-.906a.636.636 0 0 0-.466.06c-.132.076-.28.23-.378.53v.004c-.042.123-.084.36-.12.695a21.4 21.4 0 0 0-.085 1.082c-.02.326-.034.664-.048.968c-.017.383-.031.712-.048.891a.528.528 0 0 1-.061.205a.48.48 0 0 1-.424.26a.508.508 0 0 1-.454-.259c-.06-.11-.066-.24-.07-.366a27.762 27.762 0 0 1-.005-.445c-.003-.342-.003-.729-.003-.838c0-1.044 0-2.056-.092-3.066c-.055-.403-.384-.64-.803-.644c-.42-.004-.765.227-.833.647Zm2.601-.285C9.215.448 8.4.008 7.654 0c-.79-.008-1.666.466-1.83 1.49v.003c-.042.265-.072.522-.093.774c-.337-.593-.805-1.118-1.465-1.35a1.606 1.606 0 0 0-1.73.403a1.82 1.82 0 0 0-.446 1.667l.003.01c.172.671.431 1.306.683 1.908l.066.156c.18.427.352.838.497 1.252c-.375-.332-.82-.605-1.436-.665c-.414-.041-.8.027-1.125.236c-.324.208-.537.52-.674.87l-.003.009c-.203.55-.035 1.209.204 1.666c.269.536.691 1.05 1.139 1.54c.226.247.467.496.706.743l.007.007c.238.247.475.492.704.74C3.81 12.49 4.5 13.464 4.5 14.5a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5c0-1.679.302-2.854.727-4.024c.138-.38.295-.77.46-1.184c.337-.843.711-1.779 1.043-2.9c.366-1.239.326-2.133.043-2.762c-.286-.638-.793-.935-1.197-1.015c-.448-.089-.84-.01-1.15.216a7.833 7.833 0 0 0-.004-.07c-.063-.95-.731-1.638-1.499-1.813a1.635 1.635 0 0 0-1.188.17a1.638 1.638 0 0 0-.323.244Z" clip-rule="evenodd"/>`
	headingInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M8.754 2.05a.45.45 0 1 0 0 .9H9.95v4.1h-4.9v-4.1h1.204a.45.45 0 1 0 0-.9h-3.5a.45.45 0 1 0 0 .9H3.95v9.1H2.754a.45.45 0 0 0 0 .9h3.5a.45.45 0 0 0 0-.9H5.05v-4.1h4.9v4.1H8.754a.45.45 0 0 0 0 .9h3.5a.45.45 0 0 0 0-.9H11.05v-9.1h1.204a.45.45 0 0 0 0-.9h-3.5Z" clip-rule="evenodd"/>`
	heartInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4.893 2.352a2.547 2.547 0 0 0-2.54 2.553c0 1.477.857 3.012 2.02 4.43c1.021 1.246 2.222 2.335 3.127 3.143c.905-.808 2.106-1.897 3.127-3.143c1.163-1.418 2.02-2.953 2.02-4.43a2.547 2.547 0 0 0-2.54-2.553c-.836 0-1.288.291-1.567.606c-.261.295-.394.628-.515.932l-.063.156a.5.5 0 0 1-.924 0l-.063-.156c-.121-.304-.254-.637-.515-.932c-.279-.315-.73-.606-1.567-.606Zm-3.54 2.553a3.547 3.547 0 0 1 3.54-3.553c1.115 0 1.842.408 2.316.943c.112.126.208.259.291.39c.083-.131.18-.264.291-.39c.474-.535 1.2-.943 2.316-.943a3.547 3.547 0 0 1 3.54 3.553c0 1.835-1.046 3.6-2.246 5.064c-1.137 1.387-2.48 2.582-3.395 3.397l-.173.155a.5.5 0 0 1-.666 0l-.173-.155c-.916-.815-2.258-2.01-3.395-3.397C2.4 8.505 1.352 6.74 1.352 4.905Z" clip-rule="evenodd"/>`
	heartFilledInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M1.352 4.905a3.547 3.547 0 0 1 3.541-3.553c1.365 0 1.968.571 2.607 1.583c.64-1.012 1.242-1.583 2.607-1.583a3.547 3.547 0 0 1 3.54 3.553c0 1.835-1.046 3.6-2.246 5.064c-1.137 1.387-2.48 2.582-3.395 3.397l-.173.155a.5.5 0 0 1-.666 0l-.173-.155c-.916-.815-2.258-2.01-3.395-3.397C2.4 8.505 1.352 6.74 1.352 4.905Z" clip-rule="evenodd"/>`
	heightInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M7.181 1.682a.45.45 0 0 1 .637 0l2.5 2.5a.45.45 0 0 1-.637.636L7.95 3.086v8.828l1.731-1.732a.45.45 0 0 1 .637.636l-2.5 2.5a.45.45 0 0 1-.637 0l-2.5-2.5a.45.45 0 0 1 .637-.636l1.732 1.732V3.086L5.317 4.818a.45.45 0 0 1-.637-.636l2.5-2.5Z" clip-rule="evenodd"/>`
	hobbyKnifeInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M12.354 13.354a.5.5 0 0 1-.707 0l-5.25-5.25A.5.5 0 0 1 6.316 8H5a.5.5 0 0 1-.472-.336l-2.4-6.9A.5.5 0 0 1 2.936.23l5.4 4.9a.5.5 0 0 1 .164.37v.317a.501.501 0 0 1 .104.08l5.25 5.25a.5.5 0 0 1 0 .707l-1.5 1.5ZM8.25 6.957l-.793.793L12 12.293l.793-.793L8.25 6.957ZM3.717 2.288L5.355 7h.938L7.5 5.793V5.72L3.717 2.288Z" clip-rule="evenodd"/>`
	homeInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.08.222a.6.6 0 0 1 .84 0l6.75 6.64a.6.6 0 0 1-.84.856L13 6.902V12.5a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1-.5-.5V6.902l-.83.816a.6.6 0 1 1-.84-.856L7.08.222Zm.42 1.27L12 5.918V12h-2V8.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5V12H3V5.918l4.5-4.426ZM7 12h2V9H7v3Z" clip-rule="evenodd"/>`
	iconjarLogoInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.003c-.695 0-1.25.278-1.25.747c0 .247.174.37.327.477c.118.084.223.158.223.273c0 .207-.738.243-1.602.284c-1.222.06-2.698.13-2.698.716A.5.5 0 0 0 3 3h9a.5.5 0 0 0 .5-.5c0-.586-1.476-.657-2.698-.716C8.938 1.743 8.2 1.707 8.2 1.5c0-.115.105-.19.223-.273c.153-.107.327-.23.327-.477c0-.469-.555-.747-1.25-.747Zm-4.605 6.12a1.958 1.958 0 0 0-.87 1.933l.81 4.421A1.867 1.867 0 0 0 4.662 14h5.676c.898 0 1.651-.644 1.825-1.523l.812-4.42c.115-.733-.146-1.524-.84-1.963c-.406-.264-.493-.762 0-.966C12.827 4.864 12.693 4 12 4H3c-.695 0-.801.893-.105 1.157c.491.175.404.673 0 .966ZM8.2 6.25c0-.432-.114-.874-.294-1.25H4.111a1.4 1.4 0 0 1 .118.655c-.034.57-.383 1.014-.749 1.278l-.025.018l-.026.017a.96.96 0 0 0-.418.92l.807 4.4a.867.867 0 0 0 .843.712h5.676c.4 0 .756-.285.843-.711l.807-4.4c.061-.407-.092-.763-.388-.95l-.011-.007c-.393-.256-.755-.704-.79-1.284A1.398 1.398 0 0 1 10.912 5H9.055A2.127 2.127 0 0 0 9 5.5c0 .474.217.871.444 1.287c.249.456.51.934.51 1.563c0 .92-.455 1.492-1.332 1.492c-.49 0-1.044-.381-1.044-1.084c0-.405.138-.742.283-1.098c.164-.402.34-.83.34-1.41Z" clip-rule="evenodd"/>`
	idCardInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M14 11V4H1v7h13Zm1-7v7a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1ZM2 5.25A.25.25 0 0 1 2.25 5h3.5a.25.25 0 0 1 .25.25v4.5a.25.25 0 0 1-.25.25h-3.5A.25.25 0 0 1 2 9.75v-4.5ZM7.5 7a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3ZM7 9.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5ZM7.5 5a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4Z" clip-rule="evenodd"/>`
	imageInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M2.5 1h10A1.5 1.5 0 0 1 14 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 12.5v-10A1.5 1.5 0 0 1 2.5 1Zm0 1a.5.5 0 0 0-.5.5v5.864l1.682-1.682a.45.45 0 0 1 .647.01l3.545 3.798l2.808-2.808a.45.45 0 0 1 .636 0L13 9.364V2.5a.5.5 0 0 0-.5-.5h-10ZM2 12.5V9.636l1.989-1.988l3.542 3.794L8.941 13H2.5a.5.5 0 0 1-.5-.5Zm10.5.5h-2.345l-1.672-1.847L11 8.636l2 2V12.5a.5.5 0 0 1-.5.5ZM6.65 5.5a.85.85 0 1 1 1.7 0a.85.85 0 0 1-1.7 0Zm.85-1.75a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5Z" clip-rule="evenodd"/>`
	infoCircledInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0Zm6.423-3a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM6 6h1.5a.5.5 0 0 1 .5.5V10h1v1H6v-1h1V7H6V6Z" clip-rule="evenodd"/>`
	inputInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M6.5 1a.5.5 0 0 0 0 1c.627 0 .957.2 1.156.478C7.878 2.79 8 3.288 8 4v7c0 .712-.122 1.21-.344 1.522c-.199.278-.53.478-1.156.478a.5.5 0 0 0 0 1c.873 0 1.543-.3 1.97-.897l.03-.044l.03.044c.427.597 1.097.897 1.97.897a.5.5 0 0 0 0-1c-.627 0-.957-.2-1.156-.478C9.122 12.21 9 11.712 9 11V4c0-.712.122-1.21.344-1.522C9.543 2.2 9.874 2 10.5 2a.5.5 0 0 0 0-1c-.873 0-1.543.3-1.97.897l-.03.044l-.03-.044C8.042 1.3 7.372 1 6.5 1ZM14 5h-3V4h3a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-3v-1h3V5ZM6 4v1H1v5h5v1H1a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5Z" clip-rule="evenodd"/>`
	instagramLogoInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M12.91 12.909c.326-.327.582-.72.749-1.151c.16-.414.27-.886.302-1.578c.032-.693.04-.915.04-2.68c0-1.765-.008-1.987-.04-2.68c-.032-.692-.142-1.164-.302-1.578a3.185 3.185 0 0 0-.75-1.151a3.187 3.187 0 0 0-1.151-.75c-.414-.16-.886-.27-1.578-.302C9.487 1.007 9.265 1 7.5 1c-1.765 0-1.987.007-2.68.04c-.692.03-1.164.14-1.578.301a3.2 3.2 0 0 0-1.151.75a3.2 3.2 0 0 0-.75 1.151c-.16.414-.27.886-.302 1.578C1.007 5.513 1 5.735 1 7.5c0 1.765.007 1.987.04 2.68c.03.692.14 1.164.301 1.578c.164.434.42.826.75 1.151c.325.33.718.586 1.151.75c.414.16.886.27 1.578.302c.693.031.915.039 2.68.039c1.765 0 1.987-.008 2.68-.04c.692-.03 1.164-.14 1.578-.301a3.323 3.323 0 0 0 1.151-.75ZM2 6.735v1.53c-.002.821-.002 1.034.02 1.5c.026.586.058 1.016.156 1.34c.094.312.199.63.543 1.012c.344.383.675.556 1.097.684c.423.127.954.154 1.415.175c.522.024.73.024 1.826.024H8.24c.842.001 1.054.002 1.526-.02c.585-.027 1.015-.059 1.34-.156c.311-.094.629-.2 1.011-.543c.383-.344.556-.676.684-1.098c.127-.422.155-.953.176-1.414C13 9.247 13 9.04 13 7.947v-.89c0-1.096 0-1.303-.023-1.826c-.021-.461-.049-.992-.176-1.414c-.127-.423-.3-.754-.684-1.098c-.383-.344-.7-.449-1.011-.543c-.325-.097-.755-.13-1.34-.156A27.29 27.29 0 0 0 8.24 2H7.057c-1.096 0-1.304 0-1.826.023c-.461.021-.992.049-1.415.176c-.422.128-.753.301-1.097.684c-.344.383-.45.7-.543 1.012c-.098.324-.13.754-.156 1.34c-.022.466-.022.679-.02 1.5ZM7.5 5.25a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5ZM4.25 7.5a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0Zm6.72-2.72a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z" clip-rule="evenodd"/>`
	justifyCenterInnerSVG            = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7.95 6H11v3H7.95V6zm0-1V1.5a.45.45 0 0 0-.9 0V5h-3.3a.75.75 0 0 0-.75.75v3.5c0 .414.336.75.75.75h3.3v3.5a.45.45 0 1 0 .9 0V10h3.3a.75.75 0 0 0 .75-.75v-3.5a.75.75 0 0 0-.75-.75h-3.3zm-.9 4H4V6h3.05v3z" fill="currentColor"/></g>`
	justifyEndInnerSVG               = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M13.95 1.5a.45.45 0 0 0-.9 0v12a.45.45 0 1 0 .9 0v-12zM4 6h7v3H4V6zm7.25-1a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-.75.75h-7.5A.75.75 0 0 1 3 9.25v-3.5A.75.75 0 0 1 3.75 5h7.5z" fill="currentColor"/></g>`
	justifyStartInnerSVG             = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M1.05 13.5a.45.45 0 0 0 .9 0v-12a.45.45 0 1 0-.9 0v12zM11 9H4V6h7v3zm-7.25 1A.75.75 0 0 1 3 9.25v-3.5A.75.75 0 0 1 3.75 5h7.5a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-.75.75h-7.5z" fill="currentColor"/></g>`
	justifyStretchInnerSVG           = `<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M13.5 1.05a.45.45 0 0 1 .45.45v12a.45.45 0 0 1-.9 0v-12a.45.45 0 0 1 .45-.45zm-12 0a.45.45 0 0 1 .45.45v12a.45.45 0 0 1-.9 0v-12a.45.45 0 0 1 .45-.45zM4 6h7v3H4V6zm7.25-1a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-.75.75h-7.5A.75.75 0 0 1 3 9.25v-3.5A.75.75 0 0 1 3.75 5h7.5z" fill="currentColor"/></g>`
	keyboardInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M13.5 4h-12a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5Zm-12-1A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12h12a1.5 1.5 0 0 0 1.5-1.5v-6A1.5 1.5 0 0 0 13.5 3h-12ZM2 5h1v1H2V5Zm3 0H4v1h1V5Zm1 0h1v1H6V5Zm3 0H8v1h1V5Zm1 0h1v1h-1V5Zm3 0h-1v1h1V5Zm-2 2h1v1h-1V7Zm2 2h-1v1h1V9ZM9 7h1v1H9V7ZM8 7H7v1h1V7ZM5 7h1v1H5V7ZM4 7H3v1h1V7ZM2 9h1v1H2V9Zm9 0H4v1h7V9Z" clip-rule="evenodd"/>`
	lapTimerInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M5.5.5A.5.5 0 0 1 6 0h3a.5.5 0 0 1 0 1H8v1.12a6.363 6.363 0 0 1 2.992 1.016a.638.638 0 0 1 .066-.078l1-1a.625.625 0 0 1 .884.884l-.975.975A6.4 6.4 0 1 1 7 2.119V1H6a.5.5 0 0 1-.5-.5Zm-3.4 8a5.4 5.4 0 1 1 10.8 0a5.4 5.4 0 0 1-10.8 0Zm5.4 0V4.1a4.4 4.4 0 1 0 3.111 7.511L7.5 8.5Z" clip-rule="evenodd"/>`
	laptopInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M2 4.25A.25.25 0 0 1 2.25 4h10.5a.25.25 0 0 1 .25.25v7.25H2V4.25ZM2.25 3C1.56 3 1 3.56 1 4.25V12H0v.5a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5V12h-1V4.25C14 3.56 13.44 3 12.75 3H2.25Z" clip-rule="evenodd"/>`
	layersInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M7.754.82a.5.5 0 0 0-.508 0l-5.5 3.25a.5.5 0 0 0 0 .86l5.5 3.25a.5.5 0 0 0 .508 0l5.5-3.25a.5.5 0 0 0 0-.86L7.754.82ZM7.5 7.17L2.983 4.5L7.5 1.83l4.517 2.67L7.5 7.17Zm-5.93.326a.5.5 0 0 1 .684-.176l5.246 3.1l5.246-3.1a.5.5 0 1 1 .508.86l-5.5 3.25a.5.5 0 0 1-.508 0l-5.5-3.25a.5.5 0 0 1-.177-.684Zm0 3a.5.5 0 0 1 .684-.177l5.246 3.1l5.246-3.1a.5.5 0 0 1 .508.861l-5.5 3.25a.5.5 0 0 1-.508 0l-5.5-3.25a.5.5 0 0 1-.177-.684Z" clip-rule="evenodd"/>`
	layoutInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M9 2H6v11h3V2Zm1 0v11h2.5a.5.5 0 0 0 .5-.5v-10a.5.5 0 0 0-.5-.5H10ZM2.5 2H5v11H2.5a.5.5 0 0 1-.5-.5v-10a.5.5 0 0 1 .5-.5Zm0-1A1.5 1.5 0 0 0 1 2.5v10A1.5 1.5 0 0 0 2.5 14h10a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 12.5 1h-10Z" clip-rule="evenodd"/>`
	letterCaseCapitalizeInnerSVG     = `<path fill="currentColor" fill-rule="evenodd" d="M3.69 2.75a.5.5 0 0 1 .467.32l3.21 8.32a.5.5 0 0 1-.933.36L5.383 9.025H2.01L.967 11.75a.5.5 0 0 1-.934-.358l3.19-8.32a.5.5 0 0 1 .467-.321Zm.002 1.893l1.363 3.532H2.337l1.355-3.532Zm7.207.564c-1.64 0-2.89 1.479-2.89 3.403c0 2.024 1.35 3.402 2.89 3.402a3.06 3.06 0 0 0 2.255-.99v.508a.45.45 0 1 0 .9 0V5.72a.45.45 0 0 0-.9 0v.503A3.01 3.01 0 0 0 10.9 5.207Zm2.255 4.591V7.302c-.39-.721-1.213-1.244-2.067-1.244c-.978 0-2.052.908-2.052 2.552c0 1.543.974 2.552 2.052 2.552c.883 0 1.685-.667 2.067-1.364Z" clip-rule="evenodd"/>`
	letterCaseLowercaseInnerSVG      = `<path fill="currentColor" fill-rule="evenodd" d="M3.699 5.207c-1.64 0-2.89 1.479-2.89 3.403c0 2.024 1.35 3.402 2.89 3.402a3.06 3.06 0 0 0 2.255-.99v.508a.45.45 0 0 0 .9 0V5.72a.45.45 0 0 0-.9 0v.503a3.01 3.01 0 0 0-2.255-1.016Zm2.255 4.592V7.301c-.39-.72-1.213-1.243-2.067-1.243c-.978 0-2.052.908-2.052 2.552c0 1.543.974 2.552 2.052 2.552c.883 0 1.684-.666 2.067-1.363Zm4.845-4.592c-1.64 0-2.89 1.479-2.89 3.403c0 2.024 1.35 3.402 2.89 3.402a3.06 3.06 0 0 0 2.255-.99v.508a.45.45 0 0 0 .9 0V5.72a.45.45 0 1 0-.9 0v.503A3.01 3.01 0 0 0 10.8 5.207Zm2.255 4.591V7.302c-.39-.721-1.213-1.244-2.067-1.244c-.978 0-2.052.908-2.052 2.552c0 1.543.974 2.552 2.052 2.552c.883 0 1.685-.667 2.067-1.364Z" clip-rule="evenodd"/>`
	letterCaseToggleInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M11.29 2.75a.5.5 0 0 1 .467.32l3.21 8.32a.5.5 0 0 1-.934.36l-1.05-2.725H9.61L8.567 11.75a.5.5 0 0 1-.934-.358l3.19-8.32a.5.5 0 0 1 .466-.321Zm.002 1.893l1.362 3.532H9.937l1.354-3.532Zm-8.393.564c-1.64 0-2.89 1.479-2.89 3.403c0 2.024 1.35 3.402 2.89 3.402a3.06 3.06 0 0 0 2.255-.99v.508a.45.45 0 0 0 .9 0V5.72a.45.45 0 1 0-.9 0v.503A3.01 3.01 0 0 0 2.9 5.207Zm2.255 4.591V7.302c-.39-.721-1.213-1.244-2.067-1.244c-.978 0-2.052.908-2.052 2.552c0 1.543.974 2.552 2.052 2.552c.883 0 1.685-.667 2.067-1.364Z" clip-rule="evenodd"/>`
	letterCaseUppercaseInnerSVG      = `<path fill="currentColor" fill-rule="evenodd" d="M3.626 2.75a.5.5 0 0 1 .468.327l3.076 8.32a.5.5 0 0 1-.938.346L5.224 9.016H2.027L1.02 11.743a.5.5 0 1 1-.938-.347l3.076-8.32a.5.5 0 0 1 .469-.326Zm0 1.942L4.91 8.166H2.34l1.284-3.474Zm7.746-1.942a.5.5 0 0 1 .469.327l3.075 8.32a.5.5 0 1 1-.938.346L12.97 9.016H9.773l-1.008 2.727a.5.5 0 1 1-.938-.347l3.076-8.32a.5.5 0 0 1 .469-.326Zm0 1.942l1.284 3.474h-2.568l1.284-3.474Z" clip-rule="evenodd"/>`
	letterSpacingInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M4.553 1c.2 0 .38.12.46.303L8.01 8.3a.5.5 0 1 1-.92.394l-.975-2.277H2.99l-.976 2.277a.5.5 0 0 1-.92-.394l3-6.997A.5.5 0 0 1 4.552 1Zm0 1.77l1.199 2.797H3.354l1.199-2.798Zm6.503 6.232a.5.5 0 0 0 .466-.317l2.751-7.002a.5.5 0 0 0-.93-.366l-2.287 5.818L8.77 1.317a.5.5 0 1 0-.931.366l2.752 7.002a.5.5 0 0 0 .465.317Zm3.898 3.498a.4.4 0 0 1-.118.283l-2 2a.4.4 0 0 1-.565-.566l1.317-1.317H1.519l1.318 1.317a.4.4 0 0 1-.566.566l-2-2a.4.4 0 0 1 0-.566l2-2a.4.4 0 0 1 .566.566L1.519 12.1h12.069l-1.317-1.317a.4.4 0 0 1 .565-.566l2 2a.4.4 0 0 1 .118.283Z" clip-rule="evenodd"/>`
	lightningBoltInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M8.697.04a.5.5 0 0 1 .296.542L8.09 6h4.41a.5.5 0 0 1 .4.8l-6 8a.5.5 0 0 1-.893-.382L6.91 9H2.5a.5.5 0 0 1-.4-.8l6-8a.5.5 0 0 1 .597-.16ZM3.5 8h4a.5.5 0 0 1 .493.582L7.33 12.56L11.5 7h-4a.5.5 0 0 1-.493-.582L7.67 2.44L3.5 8Z" clip-rule="evenodd"/>`
	lineHeightInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M3.782 2.217a.4.4 0 0 0-.565 0l-2 2a.4.4 0 0 0 .565.566L3.1 3.466v8.068l-1.317-1.317a.4.4 0 0 0-.565.566l2 2a.4.4 0 0 0 .565 0l2-2a.4.4 0 0 0-.565-.566L3.9 11.534V3.466l1.318 1.317a.4.4 0 0 0 .565-.566l-2-2ZM8.5 4a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1h-6ZM8 7.5a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1h-6Z" clip-rule="evenodd"/>`
	linkBreakOneInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M13.354 2.354a.5.5 0 0 0-.708-.708l-11 11a.5.5 0 0 0 .708.708l11-11ZM2.037 8.467c.137.66.632 1.204 1.302 1.426l-.76.76a2.91 2.91 0 0 1-1.52-1.983c-.06-.283-.06-.61-.06-1.107v-.126c0-.497 0-.824.06-1.107c.24-1.16 1.179-2.05 2.36-2.275C3.706 4 4.04 4 4.563 4H5.5a.5.5 0 0 1 0 1h-.875c-.604 0-.836.002-1.02.037c-.802.154-1.413.752-1.568 1.496c-.035.17-.037.385-.037.967c0 .583.002.798.037.967Zm10.925-1.934a1.935 1.935 0 0 0-1.301-1.426l.76-.76a2.91 2.91 0 0 1 1.52 1.983c.059.283.059.61.059 1.107v.126c0 .497 0 .824-.059 1.107c-.24 1.16-1.18 2.05-2.36 2.275c-.288.055-.623.055-1.146.055H9.5a.5.5 0 1 1 0-1h.875c.604 0 .835-.002 1.019-.037c.803-.154 1.414-.752 1.568-1.496c.035-.17.038-.384.038-.967c0-.582-.003-.798-.038-.967Z" clip-rule="evenodd"/>`
	linkBreakTwoInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 0a.5.5 0 0 1 .5.5v2a.5.5 0 1 1-1 0v-2a.5.5 0 0 1 .5-.5ZM.646.646a.5.5 0 0 1 .708 0l1.5 1.5a.5.5 0 1 1-.708.708l-1.5-1.5a.5.5 0 0 1 0-.708ZM0 4.5A.5.5 0 0 1 .5 4h2a.5.5 0 1 1 0 1h-2a.5.5 0 0 1-.5-.5Zm12 6a.5.5 0 0 1 .5-.5h2a.5.5 0 1 1 0 1h-2a.5.5 0 0 1-.5-.5ZM10.5 12a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 .5-.5Zm1.646.146a.5.5 0 0 1 .707 0l1.5 1.5a.5.5 0 0 1-.707.707l-1.5-1.5a.5.5 0 0 1 0-.707ZM7.765 3.7c.427-.427.592-.589.747-.694c.676-.46 1.531-.468 2.166-.051c.145.095.299.245.71.657c.412.412.563.566.658.71c.417.636.408 1.49-.052 2.167c-.104.154-.267.32-.694.747l-.619.619a.5.5 0 0 0 .708.707l.618-.62l.043-.042c.37-.37.606-.606.772-.85c.675-.993.71-2.287.06-3.277c-.16-.241-.39-.472-.742-.823l-.044-.045l-.045-.045c-.351-.351-.582-.582-.824-.74c-.99-.651-2.283-.616-3.277.059c-.243.165-.48.402-.85.771l-.042.043l-.619.619a.5.5 0 1 0 .707.707l.619-.619ZM2.992 7.06l-.043.042c-.37.37-.606.606-.771.85c-.676.993-.71 2.287-.06 3.277c.158.241.39.472.74.824l.046.044l.044.045c.351.351.582.582.824.74c.99.651 2.284.616 3.278-.06c.243-.164.48-.4.849-.77l.043-.043l.618-.619a.5.5 0 0 0-.707-.707l-.619.619c-.427.427-.592.589-.746.694c-.677.46-1.532.468-2.167.051c-.144-.095-.299-.245-.71-.657c-.412-.412-.563-.566-.657-.71c-.418-.636-.409-1.49.05-2.167c.106-.154.268-.32.695-.747l.619-.619a.5.5 0 1 0-.707-.707l-.62.62Z" clip-rule="evenodd"/>`
	linkNoneOneInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M4.564 4H5.5a.5.5 0 0 1 0 1h-.875c-.604 0-.836.002-1.019.037c-.803.154-1.414.752-1.568 1.496c-.035.17-.038.384-.038.967c0 .582.003.798.038.967c.154.744.765 1.342 1.568 1.496c.183.035.415.037 1.019.037H5.5a.5.5 0 1 1 0 1h-.935c-.523 0-.858 0-1.147-.055c-1.18-.226-2.12-1.116-2.36-2.275C1 8.387 1 8.06 1 7.563a325.8 325.8 0 0 0 0-.126c0-.497 0-.824.058-1.107c.24-1.16 1.18-2.05 2.36-2.275C3.708 4 4.042 4 4.564 4Zm6.83 1.037C11.21 5.002 10.979 5 10.375 5H9.5a.5.5 0 1 1 0-1h.935c.523 0 .858 0 1.146.055c1.18.225 2.12 1.115 2.36 2.275c.06.283.06.61.059 1.107v.126c0 .497 0 .824-.059 1.107c-.24 1.16-1.18 2.05-2.36 2.275c-.288.055-.623.055-1.145.055H9.5a.5.5 0 0 1 0-1h.875c.604 0 .836-.002 1.019-.037c.803-.154 1.414-.752 1.568-1.496c.035-.17.038-.385.038-.967c0-.583-.003-.798-.038-.967c-.154-.744-.765-1.343-1.568-1.496Z" clip-rule="evenodd"/>`
	linkNoneTwoInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M10.679 2.954c-.636-.417-1.49-.408-2.167.051c-.154.105-.32.268-.747.695l-.619.618a.5.5 0 0 1-.707-.707l.619-.619l.043-.042c.37-.37.606-.607.849-.772c.994-.675 2.288-.71 3.278-.06c.241.159.472.39.824.741l.044.045l.045.045c.351.351.582.582.74.823c.651.99.616 2.284-.059 3.278c-.165.243-.402.48-.771.85l-.043.042l-.619.619a.5.5 0 1 1-.707-.707l.619-.619c.427-.427.589-.592.694-.747c.46-.676.468-1.531.051-2.167c-.095-.144-.245-.298-.657-.71c-.412-.412-.566-.562-.71-.657ZM4.318 6.44a.5.5 0 0 1 0 .707l-.619.618c-.427.427-.59.593-.694.747c-.46.677-.468 1.532-.051 2.167c.095.144.245.298.657.71c.412.412.566.563.71.657c.635.418 1.49.409 2.167-.05c.154-.106.32-.268.747-.695l.618-.619a.5.5 0 1 1 .708.707l-.62.62l-.042.042c-.37.37-.606.606-.85.771c-.993.676-2.287.71-3.277.06c-.241-.158-.472-.39-.824-.74c-.014-.016-.03-.03-.044-.045a91.54 91.54 0 0 1-.045-.045c-.351-.351-.582-.582-.741-.824c-.65-.99-.615-2.284.06-3.278c.165-.243.402-.48.771-.849l.043-.043l.619-.618a.5.5 0 0 1 .707 0Z" clip-rule="evenodd"/>`
	linkOneInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M4.625 4h-.061c-.523 0-.857 0-1.146.055c-1.18.225-2.12 1.116-2.36 2.275C1 6.613 1 6.94 1 7.437v.126c0 .497 0 .824.058 1.107c.24 1.16 1.18 2.05 2.36 2.275c.289.055.623.055 1.146.055H5.5a.5.5 0 1 0 0-1h-.875c-.604 0-.836-.002-1.02-.037c-.802-.154-1.413-.752-1.568-1.496C2.002 8.297 2 8.083 2 7.5c0-.582.002-.798.037-.967c.155-.744.766-1.342 1.569-1.496C3.789 5.002 4.02 5 4.625 5H5.5a.5.5 0 0 0 0-1h-.875Zm5.75 1c.604 0 .835.002 1.019.037c.803.154 1.414.752 1.568 1.496c.035.17.038.385.038.967c0 .583-.003.798-.038.967c-.154.744-.765 1.342-1.568 1.496c-.184.035-.415.037-1.02.037H9.5a.5.5 0 0 0 0 1h.935c.523 0 .857 0 1.146-.055c1.18-.225 2.12-1.116 2.36-2.275C14 8.387 14 8.06 14 7.563v-.126c0-.497 0-.824-.059-1.107c-.24-1.16-1.18-2.05-2.36-2.275C11.293 4 10.958 4 10.435 4H9.5a.5.5 0 0 0 0 1h.875ZM5 7a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1H5Z" clip-rule="evenodd"/>`
	linkTwoInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M8.512 3.005c.676-.46 1.531-.468 2.167-.05c.144.094.298.244.71.656c.412.412.562.566.657.71c.417.636.408 1.49-.051 2.167c-.105.155-.267.32-.694.747l-.62.619a.5.5 0 0 0 .708.707l.619-.619l.043-.043c.37-.37.606-.606.771-.849c.675-.994.71-2.287.06-3.278c-.159-.241-.39-.472-.741-.823l-.045-.045l-.044-.045c-.352-.351-.583-.582-.824-.74c-.99-.65-2.284-.616-3.278.06c-.243.164-.48.4-.85.77l-.042.043l-.619.62a.5.5 0 1 0 .707.706l.62-.618c.426-.427.592-.59.746-.695ZM4.318 7.147a.5.5 0 0 0-.707-.707l-.619.618l-.043.043c-.37.37-.606.606-.771.85c-.675.993-.71 2.287-.06 3.277c.159.242.39.473.741.824l.045.045l.044.044c.352.351.583.583.824.741c.99.65 2.284.616 3.278-.06c.243-.165.48-.401.849-.771l.043-.043l.619-.619a.5.5 0 1 0-.708-.707l-.618.619c-.427.427-.593.59-.747.694c-.676.46-1.532.469-2.167.051c-.144-.094-.298-.245-.71-.657c-.412-.412-.562-.566-.657-.71c-.417-.635-.408-1.49.051-2.167c.105-.154.267-.32.694-.747l.619-.618Zm5.304-1.061a.5.5 0 0 0-.707-.708L5.379 8.914a.5.5 0 1 0 .707.707l3.536-3.535Z" clip-rule="evenodd"/>`
	linkedinLogoInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M2 1a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H2Zm1.05 5h1.9v6h-1.9V6Zm2.025-1.995a1.075 1.075 0 1 1-2.15 0a1.075 1.075 0 0 1 2.15 0ZM12 8.357c0-1.805-1.167-2.507-2.326-2.507a2.206 2.206 0 0 0-1.095.231c-.257.13-.526.424-.734.938h-.053V6H6v6.005h1.906V8.81c-.027-.327.077-.75.291-1.001c.215-.252.52-.312.753-.342h.073c.606 0 1.056.375 1.056 1.32v3.217h1.906L12 8.357Z" clip-rule="evenodd"/>`
	listBulletInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 5.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5ZM4 4.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5ZM4.5 7a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1h-9Zm0 3a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1h-9ZM2.25 7.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-.75 3.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z" clip-rule="evenodd"/>`
	lockClosedInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M5 4.636c0-.876.242-1.53.643-1.962c.396-.427 1.003-.696 1.858-.696c.856 0 1.462.269 1.857.694c.4.431.642 1.085.642 1.961V6H5V4.636ZM4 6V4.636c0-1.055.293-1.978.91-2.643c.623-.67 1.517-1.015 2.591-1.015c1.075 0 1.969.344 2.59 1.014c.617.664.909 1.587.909 2.641V6h1a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1ZM3 7h9v6H3V7Z" clip-rule="evenodd"/>`
	lockOpenOneInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M7.499 0C6.326 0 5.36.39 4.738 1.194C4.238 1.839 4 2.682 4 3.634h1c0-.79.197-1.4.528-1.828c.388-.5 1.024-.806 1.97-.806c.859 0 1.465.265 1.86.686c.4.426.642 1.074.642 1.95V6H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-1V3.636c0-1.055-.293-1.974-.912-2.634C9.465.338 8.57 0 7.498 0ZM3 7h9v6H3V7Z" clip-rule="evenodd"/>`
	lockOpenTwoInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M9 3.636c0-.876.242-1.524.642-1.95c.395-.421 1.001-.686 1.86-.686c.946 0 1.582.306 1.97.806c.331.427.528 1.037.528 1.827h1c0-.95-.237-1.794-.738-2.44C13.64.39 12.674 0 11.502 0c-1.073 0-1.967.338-2.59 1.002C8.294 1.662 8 2.582 8 3.636V6H1a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H9V3.636ZM1 7h9v6H1V7Z" clip-rule="evenodd"/>`
	loopInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M3.354 1.854a.5.5 0 1 0-.708-.708l-2 2a.5.5 0 0 0 0 .708l2 2a.5.5 0 1 0 .708-.708L2.207 4H9.5A3.5 3.5 0 0 1 13 7.5a.5.5 0 0 0 1 0A4.5 4.5 0 0 0 9.5 3H2.207l1.147-1.146ZM2 7.5a.5.5 0 0 0-1 0A4.5 4.5 0 0 0 5.5 12h7.293l-1.147 1.146a.5.5 0 0 0 .708.708l2-2a.5.5 0 0 0 0-.708l-2-2a.5.5 0 0 0-.708.708L12.793 11H5.5A3.5 3.5 0 0 1 2 7.5Z" clip-rule="evenodd"/>`
	magicWandInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M13.9.5a.4.4 0 0 0-.8 0v.6h-.6a.4.4 0 0 0 0 .8h.6v.6a.4.4 0 0 0 .8 0v-.6h.6a.4.4 0 0 0 0-.8h-.6V.5Zm-2.046 2.646a.5.5 0 0 1 0 .708l-1 1a.5.5 0 1 1-.707-.708l1-1a.5.5 0 0 1 .707 0Zm-2 2a.5.5 0 0 1 0 .708l-7 7a.5.5 0 1 1-.708-.708l7-7a.5.5 0 0 1 .708 0ZM13.5 5.1c.22 0 .4.18.4.4v.6h.6a.4.4 0 0 1 0 .8h-.6v.6a.4.4 0 0 1-.8 0v-.6h-.6a.4.4 0 0 1 0-.8h.6v-.6c0-.22.18-.4.4-.4ZM8.9.5a.4.4 0 0 0-.8 0v.6h-.6a.4.4 0 1 0 0 .8h.6v.6a.4.4 0 1 0 .8 0v-.6h.6a.4.4 0 0 0 0-.8h-.6V.5Z" clip-rule="evenodd"/>`
	magnifyingGlassInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M10 6.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm-.691 3.516a4.5 4.5 0 1 1 .707-.707l2.838 2.837a.5.5 0 0 1-.708.708L9.31 10.016Z" clip-rule="evenodd"/>`
	marginInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 2a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm3 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM8 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm2.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm3.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM1.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm.5-3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM1.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 4.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM13.5 11a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm.5-3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM13.5 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM5 13.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm2.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm3.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm2.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM4 5a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5Zm1 0h5v5H5V5Z" clip-rule="evenodd"/>`
	maskOffInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M1 2h13v11H1V2ZM0 2a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V2Zm4.875 5.5a2.625 2.625 0 1 1 5.25 0a2.625 2.625 0 0 1-5.25 0ZM7.5 3.875a3.625 3.625 0 1 0 0 7.25a3.625 3.625 0 0 0 0-7.25Z" clip-rule="evenodd"/>`
	maskOnInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1 1a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H1Zm6.5 9.625a3.125 3.125 0 1 0 0-6.25a3.125 3.125 0 0 0 0 6.25Z" clip-rule="evenodd"/>`
	minusInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M2.25 7.5a.5.5 0 0 1 .5-.5h9.5a.5.5 0 0 1 0 1h-9.5a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	minusCircledInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0ZM4.5 7a.5.5 0 0 0 0 1h6a.5.5 0 1 0 0-1h-6Z" clip-rule="evenodd"/>`
	mixInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M2.15 4a1.85 1.85 0 1 1 3.7 0a1.85 1.85 0 0 1-3.7 0ZM4 1.25a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5ZM5.82 11L2.5 12.837V9.163L5.82 11ZM2.64 8.212a.7.7 0 0 0-1.039.612v4.352a.7.7 0 0 0 1.039.613l3.933-2.176a.7.7 0 0 0 0-1.225L2.64 8.212ZM8.3 9a.7.7 0 0 1 .7-.7h4a.7.7 0 0 1 .7.7v4a.7.7 0 0 1-.7.7H9a.7.7 0 0 1-.7-.7V9Zm.9.2v3.6h3.6V9.2H9.2Zm4.243-7.007a.45.45 0 0 0-.636-.636L11 3.364L9.193 1.557a.45.45 0 1 0-.636.636L10.364 4L8.557 5.807a.45.45 0 1 0 .636.636L11 4.636l1.807 1.807a.45.45 0 0 0 .636-.636L11.636 4l1.807-1.807Z" clip-rule="evenodd"/>`
	mixerHorizontalInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM3 5c.017 0 .033 0 .05-.002a2.5 2.5 0 0 0 4.9 0A.507.507 0 0 0 8 5h5.5a.5.5 0 0 0 0-1H8c-.017 0-.033 0-.05.002a2.5 2.5 0 0 0-4.9 0A.507.507 0 0 0 3 4H1.5a.5.5 0 0 0 0 1H3Zm8.95 5.998a2.5 2.5 0 0 1-4.9 0A.507.507 0 0 1 7 11H1.5a.5.5 0 0 1 0-1H7c.017 0 .033 0 .05.002a2.5 2.5 0 0 1 4.9 0A.506.506 0 0 1 12 10h1.5a.5.5 0 0 1 0 1H12c-.017 0-.033 0-.05-.002ZM8 10.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/>`
	mixerVerticalInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M5 1.5a.5.5 0 0 0-1 0V7c0 .017 0 .033.002.05a2.5 2.5 0 0 0 0 4.9A.506.506 0 0 0 4 12v1.5a.5.5 0 0 0 1 0V12c0-.017 0-.033-.002-.05a2.5 2.5 0 0 0 0-4.9A.507.507 0 0 0 5 7V1.5Zm6 0a.5.5 0 0 0-1 0V3c0 .017 0 .033.002.05a2.5 2.5 0 0 0 0 4.9A.507.507 0 0 0 10 8v5.5a.5.5 0 0 0 1 0V8c0-.017 0-.033-.002-.05a2.5 2.5 0 0 0 0-4.9A.507.507 0 0 0 11 3V1.5ZM4.5 8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM9 5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/>`
	mobileInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M4 2.5a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5h-6a.5.5 0 0 1-.5-.5v-10ZM4.5 1A1.5 1.5 0 0 0 3 2.5v10A1.5 1.5 0 0 0 4.5 14h6a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 10.5 1h-6ZM6 11.65a.35.35 0 1 0 0 .7h3a.35.35 0 1 0 0-.7H6Z" clip-rule="evenodd"/>`
	modulzLogoInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M7.26 3.167L4.37 5.333V1l2.89 2.167ZM1 8.222l2.889-2.166L1 3.889v4.333ZM1 14l2.889-2.167L1 9.667V14Zm6.74-5.778l2.89-2.166l-2.89-2.167v4.333ZM14 3.167l-2.889 2.166V1L14 3.167Zm-2.889 7.944L14 8.944l-2.889-2.166v4.333Zm-7.222 0L1 8.944l2.889-2.166v4.333Zm.481-5.055l2.89 2.166V3.89L4.37 6.056Zm-.481-.723L1 3.167L3.889 1v4.333ZM7.74 3.167l2.889 2.166V1L7.74 3.167ZM14 8.222l-2.889-2.166L14 3.889v4.333Zm-2.889 3.611L14 14V9.667l-2.889 2.166Z" clip-rule="evenodd"/>`
	moonInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M2.9.5a.4.4 0 0 0-.8 0v.6h-.6a.4.4 0 1 0 0 .8h.6v.6a.4.4 0 1 0 .8 0v-.6h.6a.4.4 0 0 0 0-.8h-.6V.5Zm3 3a.4.4 0 1 0-.8 0v.6h-.6a.4.4 0 1 0 0 .8h.6v.6a.4.4 0 1 0 .8 0v-.6h.6a.4.4 0 0 0 0-.8h-.6v-.6Zm-4 3a.4.4 0 1 0-.8 0v.6H.5a.4.4 0 1 0 0 .8h.6v.6a.4.4 0 0 0 .8 0v-.6h.6a.4.4 0 0 0 0-.8h-.6v-.6ZM8.544.982l-.298-.04c-.213-.024-.34.224-.217.4A6.57 6.57 0 0 1 9.203 5.1a6.602 6.602 0 0 1-6.243 6.59c-.214.012-.333.264-.183.417a6.8 6.8 0 0 0 .21.206l.072.066l.26.226l.188.148l.121.09l.187.131l.176.115c.12.076.244.149.37.217l.264.135l.26.12l.303.122l.244.086a6.568 6.568 0 0 0 1.103.26l.317.04l.267.02a6.6 6.6 0 0 0 6.943-7.328l-.037-.277a6.557 6.557 0 0 0-.384-1.415l-.113-.268l-.077-.166l-.074-.148a6.602 6.602 0 0 0-.546-.883l-.153-.2l-.199-.24l-.163-.18l-.12-.124l-.16-.158l-.223-.2l-.32-.26l-.245-.177l-.292-.19l-.321-.186l-.328-.165l-.113-.052l-.24-.101l-.276-.104l-.252-.082l-.325-.09l-.265-.06l-.292-.053Zm1.86 4.318a7.578 7.578 0 0 0-.572-2.894a5.601 5.601 0 1 1-4.748 10.146a7.61 7.61 0 0 0 3.66-2.51a.749.749 0 0 0 1.355-.442a.75.75 0 0 0-.584-.732c.062-.116.122-.235.178-.355A1.25 1.25 0 1 0 10.35 6.2c.034-.295.052-.595.052-.9Z" clip-rule="evenodd"/>`
	moveInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M7.818.932a.45.45 0 0 0-.636 0l-1.75 1.75a.45.45 0 1 0 .636.636L7 2.386V5.5a.5.5 0 0 0 1 0V2.386l.932.932a.45.45 0 0 0 .636-.636L7.818.932ZM8 9.5a.5.5 0 0 0-1 0v3.114l-.932-.932a.45.45 0 0 0-.636.636l1.75 1.75a.45.45 0 0 0 .636 0l1.75-1.75a.45.45 0 0 0-.636-.636L8 12.614V9.5Zm1-2a.5.5 0 0 1 .5-.5h3.114l-.932-.932a.45.45 0 0 1 .636-.636l1.75 1.75a.45.45 0 0 1 0 .636l-1.75 1.75a.45.45 0 0 1-.636-.636L12.614 8H9.5a.5.5 0 0 1-.5-.5ZM3.318 6.068L2.386 7H5.5a.5.5 0 0 1 0 1H2.386l.932.932a.45.45 0 0 1-.636.636l-1.75-1.75a.45.45 0 0 1 0-.636l1.75-1.75a.45.45 0 1 1 .636.636Z" clip-rule="evenodd"/>`
	notionLogoInnerSVG               = `<path fill="currentColor" d="M3.258 3.117c.42.341.577.315 1.366.262l7.433-.446c.158 0 .027-.157-.026-.183l-1.235-.893c-.236-.184-.551-.394-1.155-.341l-7.198.525c-.262.026-.315.157-.21.262l1.025.814Zm.446 1.732v7.821c0 .42.21.578.683.552l8.17-.473c.472-.026.525-.315.525-.656V4.324c0-.34-.131-.525-.42-.499l-8.538.499c-.315.026-.42.184-.42.525Zm8.065.42c.052.236 0 .472-.237.499l-.394.078v5.774c-.341.184-.657.29-.92.29c-.42 0-.525-.132-.84-.526L6.803 7.342v3.911l.815.184s0 .472-.657.472l-1.812.105c-.053-.105 0-.367.184-.42l.472-.13V6.292L5.15 6.24c-.053-.236.078-.577.446-.604l1.944-.13L10.22 9.6V5.978l-.683-.079c-.053-.289.157-.499.42-.525l1.813-.105Zm-9.93-3.937L9.326.781c.919-.08 1.155-.026 1.733.394l2.39 1.68c.395.288.526.367.526.682v9.212c0 .578-.21.92-.946.971l-8.694.525c-.552.027-.815-.052-1.104-.42l-1.76-2.283c-.315-.42-.446-.735-.446-1.103V2.25c0-.472.21-.866.814-.918Z"/>`
	opacityInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 1.5C4.5 4.25 3 6.5 3 9a4.5 4.5 0 1 0 9 0c0-2.5-1.5-4.75-4.5-7.5ZM11 9c0-1.888-1.027-3.728-3.5-6.126C5.027 5.272 4 7.112 4 9a3.5 3.5 0 1 0 7 0Z" clip-rule="evenodd"/>`
	openInNewWindowInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M12 13a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v3.5a.5.5 0 0 0 1 0V3h9v9H8.5a.5.5 0 0 0 0 1H12ZM9 6.5v3a.5.5 0 0 1-1 0V7.707l-5.146 5.147a.5.5 0 0 1-.708-.708L7.293 7H5.5a.5.5 0 0 1 0-1h3a.498.498 0 0 1 .5.497" clip-rule="evenodd"/>`
	overlineInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 1.1a.4.4 0 1 0 0 .8h8a.4.4 0 0 0 0-.8h-8ZM5 4.25a.5.5 0 0 0-1 0v5.3a3.5 3.5 0 0 0 7 0v-5.3a.5.5 0 0 0-1 0v5.3a2.5 2.5 0 0 1-5 0v-5.3Z" clip-rule="evenodd"/>`
	paddingInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M2.857 2h9.286c.473 0 .857.384.857.857v9.286a.857.857 0 0 1-.857.857H2.857A.857.857 0 0 1 2 12.143V2.857C2 2.384 2.384 2 2.857 2ZM1 2.857C1 1.831 1.831 1 2.857 1h9.286C13.168 1 14 1.831 14 2.857v9.286A1.857 1.857 0 0 1 12.143 14H2.857A1.857 1.857 0 0 1 1 12.143V2.857ZM7.5 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-3 6a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM5 7.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM4.5 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm6.5 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM10.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm.5-3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM7.5 11a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/>`
	paperPlaneInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M1.203 1.043a.5.5 0 0 0-.635.709L3.921 7.5L.568 13.248a.5.5 0 0 0 .635.709l13.5-6a.5.5 0 0 0 0-.914l-13.5-6ZM4.846 7.1L2.212 2.586L13.27 7.5L2.212 12.414L4.846 7.9H9a.4.4 0 1 0 0-.8H4.846Z" clip-rule="evenodd"/>`
	pauseInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M6.05 2.75a.55.55 0 0 0-1.1 0v9.5a.55.55 0 0 0 1.1 0v-9.5Zm4 0a.55.55 0 0 0-1.1 0v9.5a.55.55 0 0 0 1.1 0v-9.5Z" clip-rule="evenodd"/>`
	pencilOneInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M11.854 1.146a.5.5 0 0 0-.707 0L3.714 8.578a1 1 0 0 0-.212.314L2.04 12.303a.5.5 0 0 0 .657.657l3.411-1.463a1 1 0 0 0 .314-.211l7.432-7.432a.5.5 0 0 0 0-.708l-2-2Zm-7.432 8.14l7.078-7.08L12.793 3.5l-7.078 7.078l-1.496.641l-.438-.438l.64-1.496Z" clip-rule="evenodd"/>`
	pencilTwoInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M12.146 1.146a.5.5 0 0 1 .707 0l2 2a.5.5 0 0 1 0 .708l-3.942 3.942a1 1 0 0 1-.26.188L6.724 9.947a.5.5 0 0 1-.671-.67l1.963-3.928a1 1 0 0 1 .188-.26l3.942-3.943Zm.354 1.061l-3.59 3.59l-1.037 2.076l.254.254l2.077-1.038L13.793 3.5L12.5 2.207ZM10 2L9 3H4.9c-.428 0-.72 0-.944.019c-.22.018-.332.05-.41.09a1 1 0 0 0-.437.437c-.04.078-.072.19-.09.41C3 4.18 3 4.472 3 4.9v6.2c0 .428 0 .72.019.944c.018.22.05.332.09.41a1 1 0 0 0 .437.437c.078.04.19.072.41.09c.225.019.516.019.944.019h6.2c.428 0 .72 0 .944-.019c.22-.018.332-.05.41-.09a1 1 0 0 0 .437-.437c.04-.078.072-.19.09-.41c.019-.225.019-.516.019-.944V7l1-1v5.12c0 .403 0 .735-.022 1.006c-.023.281-.072.54-.196.782a2 2 0 0 1-.874.874c-.243.124-.501.173-.782.196c-.27.022-.603.022-1.005.022H4.88c-.403 0-.735 0-1.006-.022c-.281-.023-.54-.072-.782-.196a2 2 0 0 1-.874-.874c-.124-.243-.173-.501-.196-.782C2 11.856 2 11.523 2 11.12V4.88c0-.403 0-.735.022-1.006c.023-.281.072-.54.196-.782a2 2 0 0 1 .874-.874c.243-.124.501-.173.782-.196C4.144 2 4.477 2 4.88 2H10Z" clip-rule="evenodd"/>`
	personInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.875a3.625 3.625 0 0 0-1.006 7.109c-1.194.145-2.218.567-2.99 1.328c-.982.967-1.479 2.408-1.479 4.288a.475.475 0 1 0 .95 0c0-1.72.453-2.88 1.196-3.612c.744-.733 1.856-1.113 3.329-1.113s2.585.38 3.33 1.113c.742.733 1.195 1.892 1.195 3.612a.475.475 0 1 0 .95 0c0-1.88-.497-3.32-1.48-4.288c-.77-.76-1.795-1.183-2.989-1.328A3.627 3.627 0 0 0 7.5.875ZM4.825 4.5a2.675 2.675 0 1 1 5.35 0a2.675 2.675 0 0 1-5.35 0Z" clip-rule="evenodd"/>`
	pieChartInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M1.85 7.5a5.65 5.65 0 1 1 11.3 0a5.65 5.65 0 0 1-11.3 0ZM7.5.85a6.65 6.65 0 1 0 0 13.3a6.65 6.65 0 0 0 0-13.3ZM7 8V3.128A4.4 4.4 0 0 1 11.872 8H7Z" clip-rule="evenodd"/>`
	pilcrowInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M3 5.5C3 7.983 4.992 9 7 9v3.5a.5.5 0 0 0 1 0V3.1h1v9.4a.5.5 0 0 0 1 0V3.1h1.5a.55.55 0 1 0 0-1.1H7C4.992 2 3 3.017 3 5.5Z" clip-rule="evenodd"/>`
	pinBottomInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M13.5 13.95a.45.45 0 0 0 0-.9h-12a.45.45 0 1 0 0 .9h12Zm-2.432-6.382a.45.45 0 1 0-.636-.636L7.95 9.414V1.5a.45.45 0 0 0-.9 0v7.914L4.568 6.932a.45.45 0 1 0-.636.636l3.25 3.25a.45.45 0 0 0 .636 0l3.25-3.25Z" clip-rule="evenodd"/>`
	pinLeftInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M2.05 13.5a.45.45 0 0 0 .9 0v-12a.45.45 0 1 0-.9 0v12Zm6.382-2.432a.45.45 0 1 0 .636-.636L6.586 7.95H14.5a.45.45 0 0 0 0-.9H6.586l2.482-2.482a.45.45 0 1 0-.636-.636l-3.25 3.25a.45.45 0 0 0 0 .636l3.25 3.25Z" clip-rule="evenodd"/>`
	pinRightInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M12.95 1.5a.45.45 0 0 0-.9 0v12a.45.45 0 1 0 .9 0v-12ZM6.568 3.932a.45.45 0 1 0-.636.636L8.414 7.05H.5a.45.45 0 0 0 0 .9h7.914l-2.482 2.482a.45.45 0 1 0 .636.636l3.25-3.25a.45.45 0 0 0 0-.636l-3.25-3.25Z" clip-rule="evenodd"/>`
	pinTopInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 1.05a.45.45 0 1 0 0 .9h12a.45.45 0 1 0 0-.9h-12Zm2.432 6.382a.45.45 0 1 0 .636.636L7.05 5.586V13.5a.45.45 0 0 0 .9 0V5.586l2.482 2.482a.45.45 0 1 0 .636-.636l-3.25-3.25a.45.45 0 0 0-.636 0l-3.25 3.25Z" clip-rule="evenodd"/>`
	playInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M3.242 2.322a.5.5 0 0 1 .491-.014l9 4.75a.5.5 0 0 1 0 .884l-9 4.75A.5.5 0 0 1 3 12.25v-9.5a.5.5 0 0 1 .242-.428ZM4 3.579v7.842L11.429 7.5L4 3.58Z" clip-rule="evenodd"/>`
	plusInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M8 2.75a.5.5 0 0 0-1 0V7H2.75a.5.5 0 0 0 0 1H7v4.25a.5.5 0 0 0 1 0V8h4.25a.5.5 0 0 0 0-1H8V2.75Z" clip-rule="evenodd"/>`
	plusCircledInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0ZM7.5 4a.5.5 0 0 1 .5.5V7h2.5a.5.5 0 1 1 0 1H8v2.5a.5.5 0 0 1-1 0V8H4.5a.5.5 0 0 1 0-1H7V4.5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`
	questionMarkInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M5.075 4.1c0-1.189 1.182-2.175 2.425-2.175c1.243 0 2.425.986 2.425 2.175c0 1.099-.557 1.614-1.306 2.279l-.031.027C7.845 7.065 6.925 7.88 6.925 9.5a.575.575 0 1 0 1.15 0c0-1.085.554-1.594 1.307-2.26l.02-.02c.748-.662 1.673-1.482 1.673-3.12C11.075 2.128 9.219.775 7.5.775S3.925 2.128 3.925 4.1a.575.575 0 1 0 1.15 0ZM7.5 13.358a.875.875 0 1 0 0-1.75a.875.875 0 0 0 0 1.75Z" clip-rule="evenodd"/>`
	questionMarkCircledInnerSVG      = `<path fill="currentColor" fill-rule="evenodd" d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 1 0 0 11.346a5.673 5.673 0 0 0 0-11.346Zm.75 8.673a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-2.2-4.25c0-.678.585-1.325 1.45-1.325s1.45.647 1.45 1.325c0 .491-.27.742-.736 1.025c-.051.032-.111.066-.176.104a5.28 5.28 0 0 0-.564.36c-.242.188-.524.493-.524.961a.55.55 0 0 0 1.1.004a.443.443 0 0 1 .1-.098c.102-.079.215-.144.366-.232c.078-.045.167-.097.27-.159c.534-.325 1.264-.861 1.264-1.965c0-1.322-1.115-2.425-2.55-2.425c-1.435 0-2.55 1.103-2.55 2.425a.55.55 0 0 0 1.1 0Z" clip-rule="evenodd"/>`
	quoteInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M9.425 3.441c.631-.204 1.359-.2 1.954.105c1.374.706 1.969 2.526 1.416 4.454c-.248.865-.685 1.705-1.609 2.552c-.924.848-2.206 1.348-2.8 1.348A.38.38 0 0 1 8 11.525c0-.207.176-.375.386-.375c.679 0 1.286-.37 2.005-.914c.55-.417.98-.95 1.217-1.414c.455-.888.47-2.14-.265-2.473a1.8 1.8 0 0 1-1.366.61c-1.2 0-1.907-.965-1.876-1.839c.029-.835.56-1.43 1.324-1.679Zm-6 0c.631-.204 1.359-.2 1.954.105C6.753 4.252 7.348 6.072 6.795 8c-.248.865-.685 1.705-1.609 2.552c-.924.848-2.206 1.348-2.8 1.348A.38.38 0 0 1 2 11.525c0-.207.176-.375.386-.375c.679 0 1.286-.37 2.005-.914c.55-.417.98-.95 1.217-1.414c.455-.888.47-2.14-.265-2.473c-.353.386-.814.61-1.366.61c-1.2 0-1.907-.965-1.876-1.839c.029-.835.56-1.43 1.324-1.679Z" clip-rule="evenodd"/>`
	radiobuttonInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0Zm5.673 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/>`
	readerInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M4.2 1h-.023c-.308 0-.573 0-.79.02a1.527 1.527 0 0 0-.67.201a1.5 1.5 0 0 0-.496.495a1.52 1.52 0 0 0-.2.67C2 2.604 2 2.87 2 3.177v8.646c0 .308 0 .573.02.79c.023.231.071.459.201.67a1.5 1.5 0 0 0 .495.496c.212.13.44.178.67.2c.218.021.483.021.791.021h6.646c.308 0 .573 0 .79-.02c.231-.023.459-.071.67-.201a1.5 1.5 0 0 0 .496-.495c.13-.212.178-.44.2-.67c.021-.218.021-.483.021-.791V3.177c0-.308 0-.573-.02-.79a1.519 1.519 0 0 0-.201-.67a1.5 1.5 0 0 0-.495-.496a1.519 1.519 0 0 0-.67-.2A8.997 8.997 0 0 0 10.823 1H4.2Zm-.961 1.074c.028-.018.085-.043.242-.058C3.645 2.001 3.863 2 4.2 2h6.6c.337 0 .555 0 .72.016c.156.015.213.04.241.058a.5.5 0 0 1 .165.165c.018.028.043.085.058.242c.015.164.016.382.016.719v8.6c0 .337 0 .555-.016.72c-.015.156-.04.213-.058.241a.5.5 0 0 1-.165.165c-.028.018-.085.043-.242.058A8.534 8.534 0 0 1 10.8 13H4.2c-.337 0-.555 0-.72-.016c-.156-.015-.213-.04-.241-.058a.5.5 0 0 1-.165-.165c-.018-.028-.043-.085-.058-.242A8.558 8.558 0 0 1 3 11.8V3.2c0-.337 0-.555.016-.72c.015-.156.04-.213.058-.241a.5.5 0 0 1 .165-.165ZM5 10a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1H5Zm-.5-2.5A.5.5 0 0 1 5 7h5a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5ZM5 4a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1H5Z" clip-rule="evenodd"/>`
	reloadInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1.85 7.5c0-2.835 2.21-5.65 5.65-5.65c2.778 0 4.152 2.056 4.737 3.15H10.5a.5.5 0 0 0 0 1h3a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-1 0v1.813C12.296 3.071 10.666.85 7.5.85C3.437.85.85 4.185.85 7.5c0 3.315 2.587 6.65 6.65 6.65c1.944 0 3.562-.77 4.714-1.942a6.77 6.77 0 0 0 1.428-2.167a.5.5 0 1 0-.925-.38a5.77 5.77 0 0 1-1.216 1.846c-.971.99-2.336 1.643-4.001 1.643c-3.44 0-5.65-2.815-5.65-5.65Z" clip-rule="evenodd"/>`
	resetInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4.854 2.146a.5.5 0 0 1 0 .708L3.707 4H9a4.5 4.5 0 1 1 0 9H5a.5.5 0 0 1 0-1h4a3.5 3.5 0 1 0 0-7H3.707l1.147 1.146a.5.5 0 1 1-.708.708l-2-2a.5.5 0 0 1 0-.708l2-2a.5.5 0 0 1 .708 0Z" clip-rule="evenodd"/>`
	resumeInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M3.05 2.75a.55.55 0 1 0-1.1 0v9.5a.55.55 0 0 0 1.1 0v-9.5Zm2.683-.442A.5.5 0 0 0 5 2.75v9.5a.5.5 0 0 0 .733.442l9-4.75a.5.5 0 0 0 0-.884l-9-4.75ZM6 11.42V3.579L13.429 7.5l-7.43 3.92Z" clip-rule="evenodd"/>`
	rocketInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="m6.854 3.854l.8-.8c.644-.645 1.775-1.092 2.898-1.253a5.342 5.342 0 0 1 1.504-.02c.443.066.714.196.84.323c.127.126.257.397.323.84c.064.427.059.95-.02 1.504c-.16 1.123-.608 2.254-1.253 2.898L7.5 11.793l-1.146-1.146a.5.5 0 1 0-.708.707l1.5 1.5a.5.5 0 0 0 .708 0l.547-.548l1.17 1.951a.5.5 0 0 0 .783.097l2-2a.5.5 0 0 0 .141-.425l-.465-3.252l.624-.623c.855-.856 1.358-2.225 1.535-3.465c.09-.627.1-1.25.019-1.794c-.08-.528-.256-1.05-.604-1.399c-.349-.348-.871-.525-1.4-.604a6.333 6.333 0 0 0-1.793.02C9.17.987 7.8 1.49 6.946 2.345l-.623.624l-3.252-.465a.5.5 0 0 0-.425.141l-2 2a.5.5 0 0 0 .097.783l1.95 1.17l-.547.547a.5.5 0 0 0 0 .708l1.5 1.5a.5.5 0 1 0 .708-.708L3.207 7.5l.647-.646l3-3Zm3.245 9.34l-.97-1.617l2.017-2.016l.324 2.262l-1.37 1.37ZM3.423 5.87l2.016-2.016l-2.262-.324l-1.37 1.37l1.616.97Zm-1.07 4.484a.5.5 0 1 0-.707-.708l-1 1a.5.5 0 1 0 .708.707l1-1Zm1.5 1.5a.5.5 0 1 0-.707-.707l-2 2a.5.5 0 0 0 .708.707l2-2Zm1.5 1.5a.5.5 0 1 0-.707-.708l-1 1a.5.5 0 1 0 .708.707l1-1ZM9.5 6.748a1.249 1.249 0 1 0 0-2.498a1.249 1.249 0 0 0 0 2.498Z" clip-rule="evenodd"/>`
	rotateCounterClockwiseInnerSVG   = `<path fill="currentColor" fill-rule="evenodd" d="M7.597 2.936A.25.25 0 0 0 8 2.74V2c1.981 0 3.185.364 3.91 1.09C12.637 3.814 13 5.018 13 7a.5.5 0 0 0 1 0c0-2.056-.367-3.603-1.382-4.618C11.603 1.368 10.056 1 8 1V.261a.25.25 0 0 0-.403-.197L6.004 1.303a.25.25 0 0 0 0 .394l1.593 1.24ZM9.5 5h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5Zm-7-1A1.5 1.5 0 0 0 1 5.5v7A1.5 1.5 0 0 0 2.5 14h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 9.5 4h-7Z" clip-rule="evenodd"/>`
	rowSpacingInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M7.818.682a.45.45 0 0 0-.636 0l-2 2a.45.45 0 0 0 .636.636L7.05 2.086V5.5a.45.45 0 1 0 .9 0V2.086l1.232 1.232a.45.45 0 0 0 .636-.636l-2-2Zm.132 12.232V9.5a.45.45 0 1 0-.9 0v3.414l-1.232-1.232a.45.45 0 0 0-.636.636l2 2a.45.45 0 0 0 .636 0l2-2a.45.45 0 1 0-.636-.636L7.95 12.914ZM1.5 7a.5.5 0 0 0 0 1h12a.5.5 0 1 0 0-1h-12Z" clip-rule="evenodd"/>`
	rowsInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M14 12.85H1v1.3h13v-1.3Zm0-4H1v1.3h13v-1.3Zm-13-4h13v1.3H1v-1.3Zm13-4H1v1.3h13V.85Z" clip-rule="evenodd"/>`
	rulerHorizontalInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M.5 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5H.5Zm.5 6V5h1.075v2.5a.425.425 0 0 0 .85 0V5h1.15v1.5a.425.425 0 0 0 .85 0V5h1.15v1.5a.425.425 0 0 0 .85 0V5h1.15v2.5a.425.425 0 0 0 .85 0V5h1.15v1.5a.425.425 0 0 0 .85 0V5h1.15v1.5a.425.425 0 0 0 .85 0V5H14v5H1Z" clip-rule="evenodd"/>`
	rulerSquareInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M.5 0a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V5h9.5a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5H.5ZM1 4.075V1h3.075v3.075H1Zm0 .85V14h3v-1.075H2.75a.425.425 0 1 1 0-.85H4v-1.15H2.25a.425.425 0 0 1 0-.85H4v-1.15H2.75a.425.425 0 1 1 0-.85H4v-1.15H2.75a.425.425 0 1 1 0-.85H4v-1.15H1ZM4.925 4h1.15V2.75a.425.425 0 0 1 .85 0V4h1.15V2.75a.425.425 0 0 1 .85 0V4h1.15V2.25a.425.425 0 1 1 .85 0V4h1.15V2.75a.425.425 0 0 1 .85 0V4H14V1H4.925v3Z" clip-rule="evenodd"/>`
	scissorsInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M.95 4.485a1.55 1.55 0 1 0 3.1 0a1.55 1.55 0 0 0-3.1 0Zm1.55 2.45A2.45 2.45 0 1 1 4.773 5.4l.964.644a1.993 1.993 0 0 0-.02.068l-.154.55l-.353.236l-.994-.664c-.442.433-1.048.7-1.716.7ZM.95 10.5a1.55 1.55 0 1 1 3.1 0a1.55 1.55 0 0 1-3.1 0ZM2.5 8.05a2.45 2.45 0 1 0 2.277 1.545L15 2.757l-.951.1a10 10 0 0 0-3.818 1.208l-3.075 1.71a1 1 0 0 0-.476.606l-.253.906L4.224 8.76A2.442 2.442 0 0 0 2.5 8.05Zm4.644 1.165l.012.007l3.075 1.71a10 10 0 0 0 3.818 1.208l.951.1L8.81 8.1L7.145 9.215Z" clip-rule="evenodd"/>`
	sectionInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M2 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM2 5v5h11V5H2Zm0-1a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H2Zm-.5 10a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM4 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM3.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM6 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM5.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM8 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM7.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM10 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM9.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM12 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM11.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM14 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM13.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/>`
	sewingPinInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M6 3.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm2 2.45a2.5 2.5 0 1 0-1 0v7.55a.5.5 0 0 0 1 0V5.95Z" clip-rule="evenodd"/>`
	sewingPinFilledInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M10 3.5a2.5 2.5 0 0 1-2 2.45v7.55a.5.5 0 0 1-1 0V5.95a2.5 2.5 0 1 1 3-2.45Z" clip-rule="evenodd"/>`
	shadowInnerSVG                   = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.783 13.376c1.956-3.423 1.956-8.329 0-11.752l.434-.248c2.044 3.577 2.044 8.671 0 12.248l-.434-.248Z" opacity=".05"/><path d="M7.282 13.477c1.957-3.482 1.957-8.472 0-11.955l.436-.244c2.043 3.634 2.043 8.81 0 12.444l-.436-.245Z" opacity=".1"/><path d="M7.821 13.506c1.904-3.51 1.905-8.492.004-12.005l.44-.238c1.981 3.662 1.98 8.822-.004 12.482l-.44-.239Z" opacity=".15"/><path d="M8.413 13.429c1.782-3.5 1.783-8.354.001-11.855l.446-.227c1.854 3.644 1.853 8.666-.002 12.309l-.445-.227Z" opacity=".2"/><path d="M9.024 13.296c1.633-3.458 1.635-8.119.006-11.58l.452-.212c1.693 3.595 1.69 8.412-.005 12.005l-.453-.213Z" opacity=".25"/><path d="M9.668 13.066c1.442-3.37 1.443-7.754.003-11.125l.46-.196c1.493 3.496 1.492 8.022-.003 11.517l-.46-.197Z" opacity=".3"/><path d="M10.331 12.746c1.224-3.225 1.225-7.255.004-10.482l.467-.177c1.265 3.341 1.264 7.497-.003 10.836l-.468-.177Z" opacity=".35"/><path d="M11.015 12.299c.979-3.002.98-6.586.003-9.588l.476-.155c1.009 3.103 1.008 6.796-.003 9.898l-.476-.155Z" opacity=".4"/><path d="M11.721 11.668c.704-2.655.705-5.671.003-8.327l.483-.128c.725 2.74.724 5.844-.002 8.583l-.484-.128Z" opacity=".45"/><path d="M12.443 10.752c.41-2.114.41-4.391 0-6.505l.49-.095a17.888 17.888 0 0 1 .001 6.695l-.49-.095Z" opacity=".5"/><path d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 1 0 0 11.346a5.673 5.673 0 0 0 0-11.346Z"/></g>`
	shadowInnerInnerSVG              = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12.162 3.852c-3.804 1.037-7.273 4.506-8.31 8.31l-.482-.132c1.084-3.974 4.686-7.577 8.66-8.66l.132.482Z" opacity=".05"/><path d="M11.88 3.427C8.035 4.505 4.507 8.034 3.428 11.88l-.481-.134c1.125-4.015 4.785-7.675 8.8-8.8l.135.481Z" opacity=".1"/><path d="M11.52 3.026c-3.83 1.136-7.352 4.657-8.492 8.486l-.48-.143c1.189-3.99 4.839-7.638 8.83-8.823l.142.48Z" opacity=".15"/><path d="M11.047 2.662C7.31 3.877 3.879 7.308 2.663 11.043l-.475-.154C3.453 7 7.004 3.45 10.892 2.186l.155.476Z" opacity=".2"/><path d="M10.52 2.324c-3.6 1.29-6.896 4.585-8.192 8.183l-.47-.169c1.346-3.739 4.754-7.144 8.493-8.485l.17.47Z" opacity=".25"/><path d="M9.902 2.031C6.5 3.395 3.4 6.494 2.034 9.896l-.464-.187c1.417-3.528 4.617-6.728 8.146-8.142l.186.464Z" opacity=".3"/><path d="M9.207 1.789C6.061 3.203 3.211 6.052 1.793 9.197l-.456-.205c1.468-3.257 4.408-6.194 7.665-7.66l.205.457Z" opacity=".35"/><path d="M8.407 1.62c-2.814 1.431-5.35 3.965-6.782 6.778l-.445-.226c1.48-2.908 4.092-5.519 7-6.997l.227.446Z" opacity=".4"/><path d="M7.462 1.567c-2.375 1.38-4.508 3.512-5.89 5.887l-.432-.252c1.425-2.45 3.62-4.644 6.071-6.067l.251.432Z" opacity=".45"/><path d="M6.304 1.705a17.386 17.386 0 0 0-4.6 4.6l-.414-.28a17.886 17.886 0 0 1 4.734-4.734l.28.414Z" opacity=".5"/><path d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 1 0 0 11.346a5.673 5.673 0 0 0 0-11.346Z"/></g>`
	shadowNoneInnerSVG               = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.783 13.376c1.956-3.423 1.956-8.329 0-11.752l.434-.248c2.044 3.577 2.044 8.671 0 12.248l-.434-.248Z" opacity=".05"/><path d="M7.282 13.477c1.957-3.482 1.957-8.472 0-11.955l.436-.244c2.043 3.634 2.043 8.81 0 12.444l-.436-.245Z" opacity=".1"/><path d="M7.821 13.506c1.904-3.51 1.905-8.492.004-12.005l.44-.238c1.981 3.662 1.98 8.822-.004 12.482l-.44-.239Z" opacity=".15"/><path d="M8.413 13.429c1.782-3.5 1.783-8.354.001-11.855l.446-.227c1.854 3.644 1.853 8.666-.002 12.309l-.445-.227Z" opacity=".2"/><path d="M9.024 13.296c1.633-3.458 1.635-8.119.006-11.58l.452-.212c1.693 3.595 1.69 8.412-.005 12.005l-.453-.213Z" opacity=".25"/><path d="M9.668 13.066c1.442-3.37 1.443-7.754.003-11.125l.46-.196c1.493 3.496 1.492 8.022-.003 11.517l-.46-.197Z" opacity=".3"/><path d="M10.331 12.746c1.224-3.225 1.225-7.255.004-10.482l.467-.177c1.265 3.341 1.264 7.497-.003 10.836l-.468-.177Z" opacity=".35"/><path d="M11.015 12.299c.979-3.002.98-6.586.003-9.588l.476-.155c1.009 3.103 1.008 6.796-.003 9.898l-.476-.155Z" opacity=".4"/><path d="M11.721 11.668c.704-2.655.705-5.671.003-8.327l.483-.128c.725 2.74.724 5.844-.002 8.583l-.484-.128Z" opacity=".45"/><path d="M12.443 10.752c.41-2.114.41-4.391 0-6.505l.49-.095a17.888 17.888 0 0 1 .001 6.695l-.49-.095Z" opacity=".5"/><path d="M7.5.877a6.623 6.623 0 0 0-5.023 10.94l-.83.83a.5.5 0 1 0 .707.707l.83-.83a6.623 6.623 0 0 0 9.34-9.34l.83-.83a.5.5 0 0 0-.707-.708l-.83.83A6.597 6.597 0 0 0 7.5.878Zm3.642 2.274a5.673 5.673 0 0 0-7.992 7.992l7.992-7.992Zm-7.284 8.698a5.673 5.673 0 0 0 7.991-7.992L3.858 11.85Z"/></g>`
	shadowOuterInnerSVG              = `<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12.14 3.886a6 6 0 1 1-8.244 8.269l.425-.263a5.5 5.5 0 1 0 7.557-7.58l.262-.426Z" opacity=".05"/><path d="M12.851 5.073a5.5 5.5 0 1 1-7.778 7.778l.357-.35a5 5 0 1 0 7.07-7.07l.351-.358Z" opacity=".2"/><path d="M13.302 6.45a5 5 0 0 1-6.901 6.822l.26-.427a4.5 4.5 0 0 0 6.21-6.14l.431-.254Z" opacity=".35"/><path d="M13.374 7.94a4.5 4.5 0 0 1-5.502 5.417l.126-.484a4 4 0 0 0 4.89-4.816l.486-.117Z" opacity=".5"/><path d="M12.915 9.821a4.005 4.005 0 0 1-3.123 3.1l-.098-.49a3.504 3.504 0 0 0 2.732-2.712l.49.102Z" opacity=".65"/><path d="M1.277 7.503a6.225 6.225 0 1 1 12.45 0a6.225 6.225 0 0 1-12.45 0Zm6.225-5.275a5.275 5.275 0 1 0 0 10.55a5.275 5.275 0 0 0 0-10.55Z"/></g>`
	shareOneInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M5 7.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm.713 1.164a2.5 2.5 0 1 1 0-2.328l3.391-2.12A2.5 2.5 0 1 1 14 3.5a2.5 2.5 0 0 1-4.484 1.52l-3.53 2.207a2.526 2.526 0 0 1 0 .546l3.53 2.206a2.5 2.5 0 1 1-.411.804l-3.392-2.12ZM11.5 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z" clip-rule="evenodd"/>`
	shareTwoInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 5a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-1.25a.5.5 0 0 1 0-1h1.25A1.5 1.5 0 0 1 13 5.5v6a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 11.5v-6A1.5 1.5 0 0 1 3.5 4h1.25a.5.5 0 0 1 0 1H3.5ZM7 1.636L5.568 3.068a.45.45 0 0 1-.636-.636l2.25-2.25a.45.45 0 0 1 .636 0l2.25 2.25a.45.45 0 0 1-.636.636L8 1.636V8.5a.5.5 0 0 1-1 0V1.636Z" clip-rule="evenodd"/>`
	shuffleInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M12.354 1.146a.5.5 0 0 0-.708.708L12.793 3H12c-1.296 0-2.289.584-3.128 1.39c-.671.644-1.279 1.467-1.877 2.278c-.132.179-.263.357-.395.532C5.109 9.188 3.49 11 .5 11a.5.5 0 0 0 0 1c3.51 0 5.391-2.188 6.9-4.2c.144-.192.283-.38.42-.565c.597-.808 1.14-1.544 1.745-2.124C10.289 4.416 11.046 4 12 4h.793l-1.147 1.146a.5.5 0 0 0 .708.708l2-2a.5.5 0 0 0 0-.708l-2-2ZM.5 3c2.853 0 4.63 1.446 6.005 3.067l-.129.176a78.944 78.944 0 0 1-.484.65C4.573 5.293 3.026 4 .5 4a.5.5 0 0 1 0-1Zm8.372 7.61c-.5-.479-.963-1.057-1.414-1.655c.189-.238.369-.474.542-.705l.09-.12c.494.664.963 1.268 1.475 1.76c.724.694 1.481 1.11 2.435 1.11h.793l-1.147-1.146a.5.5 0 0 1 .708-.708l2 2a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708-.708L12.793 12H12c-1.296 0-2.289-.584-3.128-1.39Z" clip-rule="evenodd"/>`
	sizeInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M11.5 3.05a.45.45 0 0 1 .45.45v4a.45.45 0 0 1-.9 0V4.586L4.586 11.05H7.5a.45.45 0 0 1 0 .9h-4a.45.45 0 0 1-.45-.45v-4a.45.45 0 1 1 .9 0v2.914l6.464-6.464H7.5a.45.45 0 1 1 0-.9h4Z" clip-rule="evenodd"/>`
	sketchLogoInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M7.537.827a.375.375 0 0 0-.074 0l-3.5.35a.375.375 0 0 0-.266.152L.7 5.425a.373.373 0 0 0 .012.465l6.498 7.898a.375.375 0 0 0 .58 0l6.498-7.898a.374.374 0 0 0 .087-.238v-.024a.373.373 0 0 0-.075-.203L11.303 1.33a.375.375 0 0 0-.266-.152l-3.5-.35Zm3.388 4.448v-.023l-.003.023h.003Zm.01-.1h2.253l-1.939-2.649l-.315 2.649Zm-.364-3.291l-2.527-.253l2.13 3.58l.397-3.327Zm-3.615-.253l-2.527.253l.396 3.326l2.13-3.579Zm-3.206.895L1.812 5.175h2.254L3.75 2.526ZM1.794 6.025l4.965 6.034l-2.535-5.992a.301.301 0 0 1-.015-.042H1.794Zm3.357 0L7.5 12.108l2.35-6.083h-4.7Zm5.64 0a.3.3 0 0 1-.015.042l-2.535 5.992l4.965-6.034H10.79ZM7.5 2.183l1.84 3.092H5.66L7.5 2.183Z" clip-rule="evenodd"/>`
	slashInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4.109 14L9.466 1h1.352L5.46 14H4.109Z" clip-rule="evenodd"/>`
	sliderInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M10.3 7.5a1.8 1.8 0 1 1-3.6 0a1.8 1.8 0 0 1 3.6 0Zm.905.5a2.751 2.751 0 0 1-5.41 0H.5a.5.5 0 0 1 0-1h5.295a2.751 2.751 0 0 1 5.41 0H14.5a.5.5 0 0 1 0 1h-3.295Z" clip-rule="evenodd"/>`
	spaceBetweenHorizontallyInnerSVG = `<path fill="currentColor" fill-rule="evenodd" d="M14.5 1a.5.5 0 0 0-.5.5V6h-4a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h4v4.5a.5.5 0 1 0 1 0v-12a.5.5 0 0 0-.5-.5ZM5 6H1V1.5a.5.5 0 0 0-1 0v12a.5.5 0 0 0 1 0V9h4a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1Z" clip-rule="evenodd"/>`
	spaceBetweenVerticallyInnerSVG   = `<path fill="currentColor" fill-rule="evenodd" d="M1 .5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 0 1H9v4a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V1H1.5A.5.5 0 0 1 1 .5ZM7 9a1 1 0 0 0-1 1v4H1.5a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1H9v-4a1 1 0 0 0-1-1H7Z" clip-rule="evenodd"/>`
	spaceEvenlyHorizontallyInnerSVG  = `<path fill="currentColor" fill-rule="evenodd" d="M14.5 1a.5.5 0 0 0-.5.5v12a.5.5 0 1 0 1 0v-12a.5.5 0 0 0-.5-.5ZM.5 1a.5.5 0 0 0-.5.5v12a.5.5 0 0 0 1 0v-12A.5.5 0 0 0 .5 1ZM2 7a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7Zm7-1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H9Z" clip-rule="evenodd"/>`
	spaceEvenlyVerticallyInnerSVG    = `<path fill="currentColor" fill-rule="evenodd" d="M1 .5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 0 1h-12A.5.5 0 0 1 1 .5ZM7 2a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H7Zm0 6a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1H7Zm-5.5 6a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1h-12Z" clip-rule="evenodd"/>`
	speakerLoudInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M7.47 1.05a.5.5 0 0 1 .28.45v12a.5.5 0 0 1-.807.395L3.221 11H1.5A1.5 1.5 0 0 1 0 9.5v-4A1.5 1.5 0 0 1 1.5 4h1.721l3.722-2.895a.5.5 0 0 1 .527-.054Zm-.72 1.472L3.7 4.895A.5.5 0 0 1 3.393 5H1.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h1.893a.5.5 0 0 1 .307.105l3.05 2.373V2.522Zm3.528 1.326a.4.4 0 0 1 .555.111a6.407 6.407 0 0 1 0 7.081a.4.4 0 0 1-.666-.443a5.607 5.607 0 0 0 0-6.194a.4.4 0 0 1 .111-.555Zm2.4-2.418a.4.4 0 0 0-.61.518a8.602 8.602 0 0 1 0 11.104a.4.4 0 0 0 .61.518a9.402 9.402 0 0 0 0-12.14Z" clip-rule="evenodd"/>`
	speakerModerateInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a.5.5 0 0 0-.8-.4L3.333 4H1.5A1.5 1.5 0 0 0 0 5.5v4A1.5 1.5 0 0 0 1.5 11h1.833L7.2 13.9a.5.5 0 0 0 .8-.4v-12ZM3.8 4.9L7 2.5v10l-3.2-2.4a.5.5 0 0 0-.3-.1h-2a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5h2a.5.5 0 0 0 .3-.1Zm7.033-.94a.4.4 0 0 0-.666.443a5.607 5.607 0 0 1 0 6.194a.4.4 0 1 0 .666.444a6.407 6.407 0 0 0 0-7.082Z" clip-rule="evenodd"/>`
	speakerOffInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M7.724 1.053A.5.5 0 0 1 8 1.5v12a.5.5 0 0 1-.8.4L3.333 11H1.5A1.5 1.5 0 0 1 0 9.5v-4A1.5 1.5 0 0 1 1.5 4h1.833L7.2 1.1a.5.5 0 0 1 .524-.047ZM7 2.5L3.8 4.9a.5.5 0 0 1-.3.1h-2a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h2a.5.5 0 0 1 .3.1L7 12.5v-10Zm7.854 2.646a.5.5 0 0 1 0 .708L13.207 7.5l1.647 1.646a.5.5 0 0 1-.708.708L12.5 8.207l-1.646 1.647a.5.5 0 0 1-.708-.708L11.793 7.5l-1.647-1.646a.5.5 0 0 1 .708-.708L12.5 6.793l1.646-1.647a.5.5 0 0 1 .708 0Z" clip-rule="evenodd"/>`
	speakerQuietInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a.5.5 0 0 0-.8-.4L3.333 4H1.5A1.5 1.5 0 0 0 0 5.5v4A1.5 1.5 0 0 0 1.5 11h1.833L7.2 13.9a.5.5 0 0 0 .8-.4v-12ZM3.8 4.9L7 2.5v10l-3.2-2.4a.5.5 0 0 0-.3-.1h-2a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5h2a.5.5 0 0 0 .3-.1Zm6.283.156a.4.4 0 1 0-.666.443a3.623 3.623 0 0 1 0 4.002a.4.4 0 1 0 .666.443a4.423 4.423 0 0 0 0-4.888Z" clip-rule="evenodd"/>`
	squareInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v13H1V1Zm1 1v11h11V2H2Z" clip-rule="evenodd"/>`
	stackInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.754 1.82a.5.5 0 0 0-.508 0l-5.5 3.25a.5.5 0 0 0 0 .86l5.5 3.25a.5.5 0 0 0 .508 0l5.5-3.25a.5.5 0 0 0 0-.86l-5.5-3.25ZM7.5 8.17L2.983 5.5L7.5 2.83l4.517 2.67L7.5 8.17Zm-5.246.15a.5.5 0 0 0-.508.86l5.5 3.25a.5.5 0 0 0 .508 0l5.5-3.25a.5.5 0 1 0-.508-.86L7.5 11.42l-5.246-3.1Z" clip-rule="evenodd"/>`
	starInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="m6.98 1.252l-.022.05L5.588 4.6a.3.3 0 0 1-.253.184l-3.561.286l-.055.004l-.331.027l-.3.024a.3.3 0 0 0-.172.527l.23.196l.252.216l.041.036l2.713 2.324a.3.3 0 0 1 .097.297l-.83 3.475l-.012.053l-.077.323l-.07.294a.3.3 0 0 0 .448.326l.258-.158l.284-.173l.046-.028l3.049-1.863a.3.3 0 0 1 .312 0l3.049 1.863l.046.028l.284.173l.258.158a.3.3 0 0 0 .448-.326l-.07-.293l-.077-.324l-.013-.053l-.829-3.475a.3.3 0 0 1 .097-.297L13.562 6.1l.041-.036l.253-.216l.23-.196a.3.3 0 0 0-.172-.527l-.3-.024l-.332-.027l-.055-.004l-3.56-.286a.3.3 0 0 1-.254-.184L8.042 1.302l-.021-.05l-.128-.307l-.116-.279a.3.3 0 0 0-.554 0l-.116.279l-.128.307Zm.52 1.352l-.99 2.38a1.3 1.3 0 0 1-1.096.797l-2.57.206l1.958 1.677a1.3 1.3 0 0 1 .418 1.29l-.598 2.507l2.2-1.344a1.3 1.3 0 0 1 1.356 0l2.2 1.344l-.598-2.508a1.3 1.3 0 0 1 .418-1.289l1.958-1.677l-2.57-.206a1.3 1.3 0 0 1-1.096-.797l-.99-2.38Z" clip-rule="evenodd"/>`
	starFilledInnerSVG               = `<path fill="currentColor" d="M7.223.666a.3.3 0 0 1 .554 0L9.413 4.6a.3.3 0 0 0 .253.184l4.248.34a.3.3 0 0 1 .171.528L10.85 8.424a.3.3 0 0 0-.097.297l.99 4.145a.3.3 0 0 1-.45.326L7.657 10.97a.3.3 0 0 0-.312 0l-3.637 2.222a.3.3 0 0 1-.448-.326l.989-4.145a.3.3 0 0 0-.097-.297L.915 5.652a.3.3 0 0 1 .171-.527l4.248-.34a.3.3 0 0 0 .253-.185L7.223.666Z"/>`
	stitchesLogoInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M1.85 7.5c0-.287.021-.569.063-.844l7.65-4.417c.3.117.587.26.859.424L3.77 6.503a.447.447 0 0 0-.234.374a.447.447 0 0 0 .196.394l3.592 2.567l-3.453 1.994A5.638 5.638 0 0 1 1.85 7.5Zm6.518 2.775a.42.42 0 0 0 .025-.014l1.448-.836a.448.448 0 0 0 .018-.01l1.451-.838a.45.45 0 0 0 .052-.756L7.828 5.2l3.388-1.957a5.637 5.637 0 0 1 1.849 5.24L5.569 12.81a5.623 5.623 0 0 1-.882-.41l3.681-2.126Zm-.187-.931L4.817 6.939l.692-.4l3.297 2.444l-.625.36ZM6.36 6.048l.62-.357l3.296 2.444l-.62.358L6.36 6.048Zm1.973-4.137l-6.09 3.515a5.652 5.652 0 0 1 6.09-3.515ZM6.847 13.113a5.652 5.652 0 0 0 5.842-3.373l-5.842 3.373ZM7.5.85a6.65 6.65 0 1 0 0 13.3a6.65 6.65 0 0 0 0-13.3Z" clip-rule="evenodd"/>`
	stopInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3Zm10 0H3v9h9V3Z" clip-rule="evenodd"/>`
	stopwatchInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M5.5.5A.5.5 0 0 1 6 0h3a.5.5 0 0 1 0 1H8v1.12a6.363 6.363 0 0 1 2.992 1.016a.638.638 0 0 1 .066-.078l.8-.8a.625.625 0 0 1 .884.884l-.775.775A6.4 6.4 0 1 1 7 2.119V1H6a.5.5 0 0 1-.5-.5Zm-3.4 8a5.4 5.4 0 1 1 10.8 0a5.4 5.4 0 0 1-10.8 0Zm5.9-4a.5.5 0 0 0-1 0v5a.5.5 0 1 0 1 0v-5Z" clip-rule="evenodd"/>`
	stretchHorizontallyInnerSVG      = `<path fill="currentColor" fill-rule="evenodd" d="M14.5 1a.5.5 0 0 0-.5.5V6H1V1.5a.5.5 0 1 0-1 0v12a.5.5 0 0 0 1 0V9h13v4.5a.5.5 0 1 0 1 0v-12a.5.5 0 0 0-.5-.5Z" clip-rule="evenodd"/>`
	stretchVerticallyInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M1 .5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 0 1h-12A.5.5 0 0 1 1 .5ZM9 14V1H6v13H1.5a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1H9Z" clip-rule="evenodd"/>`
	strikethroughInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M5 3.25a.5.5 0 0 0-1 0V7.1H2.5a.4.4 0 1 0 0 .8H4v.65a3.5 3.5 0 1 0 7 0V7.9h1.5a.4.4 0 0 0 0-.8H11V3.25a.5.5 0 1 0-1 0V7.1H5V3.25ZM5 7.9v.65a2.5 2.5 0 0 0 5 0V7.9H5Z" clip-rule="evenodd"/>`
	sunInnerSVG                      = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 .5-.5ZM2.197 2.197a.5.5 0 0 1 .707 0L4.318 3.61a.5.5 0 0 1-.707.707L2.197 2.904a.5.5 0 0 1 0-.707ZM.5 7a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2Zm1.697 5.803a.5.5 0 0 1 0-.707l1.414-1.414a.5.5 0 1 1 .707.707l-1.414 1.414a.5.5 0 0 1-.707 0ZM12.5 7a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2Zm-1.818-2.682a.5.5 0 0 1 0-.707l1.414-1.414a.5.5 0 1 1 .707.707L11.39 4.318a.5.5 0 0 1-.707 0ZM8 12.5a.5.5 0 0 0-1 0v2a.5.5 0 0 0 1 0v-2Zm2.682-1.818a.5.5 0 0 1 .707 0l1.414 1.414a.5.5 0 1 1-.707.707l-1.414-1.414a.5.5 0 0 1 0-.707ZM5.5 7.5a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm2-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z" clip-rule="evenodd"/>`
	switchInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M10.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm-2.829 7A4.491 4.491 0 0 1 6 7.5c0-1.414.652-2.675 1.671-3.5H4.5a3.5 3.5 0 1 0 0 7h3.171ZM0 7.5A4.5 4.5 0 0 1 4.5 3h6a4.5 4.5 0 1 1 0 9h-6A4.5 4.5 0 0 1 0 7.5Z" clip-rule="evenodd"/>`
	symbolInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1.903 7.297c0 3.044 2.207 5.118 4.686 5.547a.521.521 0 1 1-.178 1.027C3.5 13.367.861 10.913.861 7.297c0-1.537.699-2.745 1.515-3.663c.585-.658 1.254-1.193 1.792-1.602H2.532a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0V2.686l-.001.002c-.572.43-1.27.957-1.875 1.638c-.715.804-1.253 1.776-1.253 2.97Zm11.108.406c0-3.012-2.16-5.073-4.607-5.533a.521.521 0 1 1 .192-1.024c2.874.54 5.457 2.98 5.457 6.557c0 1.537-.699 2.744-1.515 3.663c-.585.658-1.254 1.193-1.792 1.602h1.636a.5.5 0 1 1 0 1h-3a.5.5 0 0 1-.5-.5v-3a.5.5 0 1 1 1 0v1.845h.002c.571-.432 1.27-.958 1.874-1.64c.715-.803 1.253-1.775 1.253-2.97Z" clip-rule="evenodd"/>`
	tableInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M8 2h4.5a.5.5 0 0 1 .5.5V5H8V2ZM7 5V2H2.5a.5.5 0 0 0-.5.5V5h5ZM2 6v3h5V6H2Zm6 0h5v3H8V6Zm0 4h5v2.5a.5.5 0 0 1-.5.5H8v-3Zm-6 2.5V10h5v3H2.5a.5.5 0 0 1-.5-.5Zm-1-10A1.5 1.5 0 0 1 2.5 1h10A1.5 1.5 0 0 1 14 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 12.5v-10Z" clip-rule="evenodd"/>`
	targetInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M.9 7.5a6.6 6.6 0 1 1 13.2 0a6.6 6.6 0 0 1-13.2 0Zm6.6-5.7a5.7 5.7 0 1 0 0 11.4a5.7 5.7 0 0 0 0-11.4ZM3.075 7.5a4.425 4.425 0 1 1 8.85 0a4.425 4.425 0 0 1-8.85 0ZM7.5 3.925a3.575 3.575 0 1 0 0 7.15a3.575 3.575 0 0 0 0-7.15Zm0 1.325a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5ZM6.05 7.5a1.45 1.45 0 1 1 2.9 0a1.45 1.45 0 0 1-2.9 0Z" clip-rule="evenodd"/>`
	textInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M3.95 2.95V4.5a.45.45 0 0 1-.9 0v-2a.45.45 0 0 1 .45-.45h8a.45.45 0 0 1 .45.45v2a.45.45 0 1 1-.9 0V2.95h-3v9.1h1.204a.45.45 0 0 1 0 .9h-3.5a.45.45 0 1 1 0-.9H6.95v-9.1h-3Z" clip-rule="evenodd"/>`
	textAlignBottomInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M3.9 5.5a.4.4 0 0 0-.8 0v7.034l-1.318-1.317a.4.4 0 0 0-.565.566l2 2a.4.4 0 0 0 .565 0l2-2a.4.4 0 0 0-.565-.566l-1.318 1.317V5.5ZM8.5 13a.5.5 0 1 1 0-1h6a.5.5 0 0 1 0 1h-6Zm0-3a.5.5 0 1 1 0-1h6a.5.5 0 0 1 0 1h-6ZM8 6.5a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 0-1h-6a.5.5 0 0 0-.5.5Z" clip-rule="evenodd"/>`
	textAlignCenterInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M2 4.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Zm2 3a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5Zm-1 3a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	textAlignJustifyInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M2.5 4a.5.5 0 0 0 0 1h10a.5.5 0 0 0 0-1h-10ZM2 7.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	textAlignLeftInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M2 4.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	textAlignMiddleInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M3.5 14.9a.4.4 0 0 0 .4-.4v-4.034l1.317 1.317a.4.4 0 0 0 .565-.566l-2-2a.4.4 0 0 0-.565 0l-2 2a.4.4 0 0 0 .565.566L3.1 10.466V14.5c0 .22.18.4.4.4ZM8 10.5a.5.5 0 0 0 .5.5h6a.5.5 0 1 0 0-1h-6a.5.5 0 0 0-.5.5Zm0-3a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 0-1h-6a.5.5 0 0 0-.5.5ZM8.5 5a.5.5 0 1 1 0-1h6a.5.5 0 0 1 0 1h-6ZM3.9.5a.4.4 0 0 0-.8 0v4.034L1.781 3.217a.4.4 0 0 0-.565.566l2 2a.4.4 0 0 0 .565 0l2-2a.4.4 0 0 0-.565-.566L3.899 4.534V.5Z" clip-rule="evenodd"/>`
	textAlignRightInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M2 4.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5Zm5 3a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Zm-3 3a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	textAlignTopInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M3.9 9.5a.4.4 0 0 1-.8 0V2.466L1.781 3.783a.4.4 0 0 1-.565-.566l2-2a.4.4 0 0 1 .565 0l2 2a.4.4 0 0 1-.565.566L3.899 2.466V9.5ZM8.5 2a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1h-6Zm0 3a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1h-6ZM8 8.5a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	textNoneInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M13.354 2.354a.5.5 0 0 0-.708-.708l-.717.718a.45.45 0 0 0-.429-.314h-8a.45.45 0 0 0-.45.45v2a.45.45 0 1 0 .9 0V2.95h3v4.393l-5.304 5.303a.5.5 0 0 0 .708.708L6.95 8.757v3.293H5.754a.45.45 0 1 0 0 .9h3.5a.45.45 0 0 0 0-.9H8.05V7.657l5.304-5.303ZM8.05 6.243l3-3V2.95h-3v3.293Z" clip-rule="evenodd"/>`
	thickArrowDownInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M5 3.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V6h2.5a.5.5 0 0 1 .407.79l-5 7a.5.5 0 0 1-.814 0l-5-7A.5.5 0 0 1 2.5 6H5V3.5ZM6 4v2.5a.5.5 0 0 1-.5.5H3.472L7.5 12.64L11.528 7H9.5a.5.5 0 0 1-.5-.5V4H6Z" clip-rule="evenodd"/>`
	thickArrowLeftInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M1 7.5a.5.5 0 0 0 .21.407l7 5A.5.5 0 0 0 9 12.5V10h2.5a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5H9V2.5a.5.5 0 0 0-.79-.407l-7 5A.5.5 0 0 0 1 7.5Zm7-4.028V5.5a.5.5 0 0 0 .5.5H11v3H8.5a.5.5 0 0 0-.5.5v2.028L2.36 7.5L8 3.472Z" clip-rule="evenodd"/>`
	thickArrowRightInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M14 7.5a.5.5 0 0 1-.21.407l-7 5A.5.5 0 0 1 6 12.5V10H3.5a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5H6V2.5a.5.5 0 0 1 .79-.407l7 5A.5.5 0 0 1 14 7.5ZM7 3.472V5.5a.5.5 0 0 1-.5.5H4v3h2.5a.5.5 0 0 1 .5.5v2.028L12.64 7.5L7 3.472Z" clip-rule="evenodd"/>`
	thickArrowUpInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M7.5 1a.5.5 0 0 1 .407.21l5 7A.5.5 0 0 1 12.5 9H10v2.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V9H2.5a.5.5 0 0 1-.407-.79l5-7A.5.5 0 0 1 7.5 1ZM3.472 8H5.5a.5.5 0 0 1 .5.5V11h3V8.5a.5.5 0 0 1 .5-.5h2.028L7.5 2.36L3.472 8Z" clip-rule="evenodd"/>`
	timerInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.85a.5.5 0 0 0-.5.5v2.172a.5.5 0 1 0 1 0v-1.65a5.65 5.65 0 1 1-4.81 1.974a.5.5 0 1 0-.762-.647A6.65 6.65 0 1 0 7.5.85Zm-.76 7.23L4.224 4.573a.25.25 0 0 1 .348-.348L8.081 6.74a.96.96 0 1 1-1.34 1.34Z" clip-rule="evenodd"/>`
	tokensInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M4.5 2a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM3 4.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0ZM10.5 2a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM9 4.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm-7 6a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0ZM4.5 9a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm6-1a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM9 10.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/>`
	trackNextInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M13.05 2.75a.55.55 0 1 0-1.1 0v4.533a.5.5 0 0 0-.217-.225l-9-4.75A.5.5 0 0 0 2 2.75v9.5a.5.5 0 0 0 .733.442l9-4.75a.5.5 0 0 0 .217-.225v4.533a.55.55 0 0 0 1.1 0v-9.5ZM3 11.42V3.58l7.429 3.92L3 11.42Z" clip-rule="evenodd"/>`
	trackPreviousInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M1.95 2.75a.55.55 0 1 1 1.1 0v4.533a.5.5 0 0 1 .217-.225l9-4.75A.5.5 0 0 1 13 2.75v9.5a.5.5 0 0 1-.733.442l-9-4.75a.5.5 0 0 1-.217-.225v4.533a.55.55 0 0 1-1.1 0v-9.5ZM4.57 7.5L12 11.42V3.58L4.571 7.5Z" clip-rule="evenodd"/>`
	transformInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M.85 1.75a.9.9 0 0 1 .9-.9h1.5a.9.9 0 0 1 .9.9v.3h6.7v-.3a.9.9 0 0 1 .9-.9h1.5a.9.9 0 0 1 .9.9v1.5a.9.9 0 0 1-.9.9h-.3v6.7h.3a.9.9 0 0 1 .9.9v1.5a.9.9 0 0 1-.9.9h-1.5a.9.9 0 0 1-.9-.9v-.3h-6.7v.3a.9.9 0 0 1-.9.9h-1.5a.9.9 0 0 1-.9-.9v-1.5a.9.9 0 0 1 .9-.9h.3v-6.7h-.3a.9.9 0 0 1-.9-.9v-1.5Zm2.1 2.4v6.7h.3a.9.9 0 0 1 .9.9v.3h6.7v-.3a.9.9 0 0 1 .9-.9h.3v-6.7h-.3a.9.9 0 0 1-.9-.9v-.3h-6.7v.3a.9.9 0 0 1-.9.9h-.3Zm-.6-2.4h-.6v1.5h1.5v-1.5h-.9ZM5.1 6a.9.9 0 0 1 .9-.9h1a.9.9 0 0 1 .9.9v1a.91.91 0 0 1-.006.106A.908.908 0 0 1 8 7.1h1a.9.9 0 0 1 .9.9v1a.9.9 0 0 1-.9.9H8a.9.9 0 0 1-.9-.9V8c0-.036.002-.071.006-.106A.91.91 0 0 1 7 7.9H6a.9.9 0 0 1-.9-.9V6Zm1 0H6v1h1V6h-.9ZM8 8h1v1H8V8Zm-5.35 3.75h-.9v1.5h1.5v-1.5h-.6Zm9.1-10h1.5v1.5h-1.5v-1.5Zm.9 10h-.9v1.5h1.5v-1.5h-.6Z" clip-rule="evenodd"/>`
	transparencyGridInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M0 0h3v3H0V0Zm6 3H3v3H0v3h3v3H0v3h3v-3h3v3h3v-3h3v3h3v-3h-3V9h3V6h-3V3h3V0h-3v3H9V0H6v3Zm0 3V3h3v3H6Zm0 3H3V6h3v3Zm3 0V6h3v3H9Zm0 0H6v3h3V9Z" clip-rule="evenodd" opacity=".25"/>`
	trashInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M5.5 1a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4ZM3 3.5a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1H11v8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V4h-.5a.5.5 0 0 1-.5-.5ZM5 4h5v8H5V4Z" clip-rule="evenodd"/>`
	triangleDownInnerSVG             = `<path fill="currentColor" d="M4 6h7l-3.5 4.5L4 6Z"/>`
	triangleLeftInnerSVG             = `<path fill="currentColor" d="M9 4v7L4.5 7.5L9 4Z"/>`
	triangleRightInnerSVG            = `<path fill="currentColor" d="M6 11V4l4.5 3.5L6 11Z"/>`
	triangleUpInnerSVG               = `<path fill="currentColor" d="M4 9h7L7.5 4.5L4 9Z"/>`
	twitterLogoInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M7.233 4.696c0-1.727 1.4-3.127 3.127-3.127c1.014 0 1.823.479 2.365 1.175a5.246 5.246 0 0 0 1.626-.629a2.634 2.634 0 0 1-1.148 1.45l.002.003a5.26 5.26 0 0 0 1.5-.413l-.001.002c-.337.505-.76.95-1.248 1.313c.026.177.04.354.04.53c0 3.687-2.809 7.975-7.975 7.975a7.93 7.93 0 0 1-4.296-1.26a.5.5 0 0 1-.108-.748a.45.45 0 0 1 .438-.215c.916.108 1.83-.004 2.637-.356a3.086 3.086 0 0 1-1.69-1.876a.45.45 0 0 1 .103-.448a3.07 3.07 0 0 1-1.045-2.31v-.034a.45.45 0 0 1 .365-.442a3.068 3.068 0 0 1-.344-1.416c0-.468.003-1.058.332-1.59a.45.45 0 0 1 .323-.208a.5.5 0 0 1 .538.161a6.964 6.964 0 0 0 4.46 2.507v-.044Zm-1.712 7.279a6.936 6.936 0 0 1-2.249-.373a5.318 5.318 0 0 0 2.39-1.042a.45.45 0 0 0-.27-.804a2.174 2.174 0 0 1-1.714-.888c.19-.015.376-.048.556-.096a.45.45 0 0 0-.028-.876a2.18 2.18 0 0 1-1.644-1.474c.2.048.409.077.623.084a.45.45 0 0 0 .265-.824a2.177 2.177 0 0 1-.97-1.812c0-.168.003-.317.013-.453a7.95 7.95 0 0 0 5.282 2.376a.5.5 0 0 0 .513-.61a2.127 2.127 0 0 1 2.071-2.614c1.234 0 2.136 1.143 2.136 2.432c0 3.256-2.476 6.974-6.975 6.974Z" clip-rule="evenodd"/>`
	underlineInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M5 2.75a.5.5 0 0 0-1 0v5.3a3.5 3.5 0 0 0 7 0v-5.3a.5.5 0 1 0-1 0v5.3a2.5 2.5 0 1 1-5 0v-5.3ZM3.5 13.1a.4.4 0 1 0 0 .8h8a.4.4 0 0 0 0-.8h-8Z" clip-rule="evenodd"/>`
	updateInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M1.903 7.297c0 3.044 2.207 5.118 4.686 5.547a.521.521 0 1 1-.178 1.027C3.5 13.367.861 10.913.861 7.297c0-1.537.699-2.745 1.515-3.663c.585-.658 1.254-1.193 1.792-1.602H2.532a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0V2.686l-.001.002c-.572.43-1.27.957-1.875 1.638c-.715.804-1.253 1.776-1.253 2.97Zm11.108.406c0-3.012-2.16-5.073-4.607-5.533a.521.521 0 1 1 .192-1.024c2.874.54 5.457 2.98 5.457 6.557c0 1.537-.699 2.744-1.515 3.663c-.585.658-1.254 1.193-1.792 1.602h1.636a.5.5 0 1 1 0 1h-3a.5.5 0 0 1-.5-.5v-3a.5.5 0 1 1 1 0v1.845h.002c.571-.432 1.27-.958 1.874-1.64c.715-.803 1.253-1.775 1.253-2.97Z" clip-rule="evenodd"/>`
	uploadInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M7.818 1.182a.45.45 0 0 0-.636 0l-3 3a.45.45 0 1 0 .636.636L7.05 2.586V9.5a.45.45 0 1 0 .9 0V2.586l2.232 2.232a.45.45 0 1 0 .636-.636l-3-3ZM2.5 10a.5.5 0 0 1 .5.5V12c0 .554.446 1 .996 1h7.005A.999.999 0 0 0 12 12v-1.5a.5.5 0 1 1 1 0V12a2 2 0 0 1-1.999 2H3.996A1.997 1.997 0 0 1 2 12v-1.5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/>`
	valueInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 1 0 0 11.346a5.673 5.673 0 0 0 0-11.346Z" clip-rule="evenodd"/>`
	valueNoneInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 0 0-5.023 10.94l-.83.83a.5.5 0 1 0 .707.707l.83-.83a6.623 6.623 0 0 0 9.34-9.34l.83-.83a.5.5 0 1 0-.708-.708l-.83.83A6.597 6.597 0 0 0 7.5.878Zm3.642 2.274a5.673 5.673 0 0 0-7.992 7.992l7.992-7.992Zm-7.284 8.698a5.673 5.673 0 0 0 7.992-7.992L3.857 11.85Z" clip-rule="evenodd"/>`
	vercelLogoInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="m7.5 1l-.577 1.003L1.175 12L.6 13h13.8l-.575-1l-5.748-9.997L7.5 1Zm0 2.006L2.329 12H12.67L7.5 3.006Z" clip-rule="evenodd"/>`
	videoInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4.764 3.122A32.656 32.656 0 0 1 7.5 3c.94 0 1.868.049 2.736.122c1.044.088 1.72.148 2.236.27c.47.111.733.258.959.489c.024.025.06.063.082.09c.2.23.33.518.405 1.062c.08.583.082 1.343.082 2.492c0 1.135-.002 1.885-.082 2.46c-.074.536-.204.821-.405 1.054a2.276 2.276 0 0 1-.083.09c-.23.234-.49.379-.948.487c-.507.12-1.168.178-2.194.264c-.869.072-1.812.12-2.788.12c-.976 0-1.92-.048-2.788-.12c-1.026-.086-1.687-.144-2.194-.264c-.459-.108-.719-.253-.948-.487a2.299 2.299 0 0 1-.083-.09c-.2-.233-.33-.518-.405-1.054C1.002 9.41 1 8.66 1 7.525c0-1.149.002-1.91.082-2.492c.075-.544.205-.832.405-1.062c.023-.027.058-.065.082-.09c.226-.231.489-.378.959-.489c.517-.122 1.192-.182 2.236-.27ZM0 7.525c0-2.242 0-3.363.73-4.208c.036-.042.085-.095.124-.135c.78-.799 1.796-.885 3.826-1.056C5.57 2.05 6.527 2 7.5 2c.973 0 1.93.05 2.82.126c2.03.171 3.046.257 3.826 1.056c.039.04.087.093.124.135c.73.845.73 1.966.73 4.208c0 2.215 0 3.323-.731 4.168a3.243 3.243 0 0 1-.125.135c-.781.799-1.778.882-3.773 1.048C9.48 12.951 8.508 13 7.5 13s-1.98-.05-2.87-.124c-1.996-.166-2.993-.25-3.774-1.048a3.316 3.316 0 0 1-.125-.135C0 10.848 0 9.74 0 7.525Zm5.25-2.142a.25.25 0 0 1 .35-.23l4.828 2.118c.2.088.2.37 0 .458L5.6 9.846a.25.25 0 0 1-.35-.229V5.383Z" clip-rule="evenodd"/>`
	viewGridInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M7 2H1.5a.5.5 0 0 0-.5.5V7h6V2Zm1 0v5h6V2.5a.5.5 0 0 0-.5-.5H8ZM7 8H1v4.5a.5.5 0 0 0 .5.5H7V8Zm1 5V8h6v4.5a.5.5 0 0 1-.5.5H8ZM1.5 1A1.5 1.5 0 0 0 0 2.5v10A1.5 1.5 0 0 0 1.5 14h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 1h-12Z" clip-rule="evenodd"/>`
	viewHorizontalInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M1.5 2h12a.5.5 0 0 1 .5.5V7H1V2.5a.5.5 0 0 1 .5-.5ZM1 8v4.5a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V8H1ZM0 2.5A1.5 1.5 0 0 1 1.5 1h12A1.5 1.5 0 0 1 15 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10Z" clip-rule="evenodd"/>`
	viewNoneInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M14 2.587L1.852 13H13.5a.5.5 0 0 0 .5-.5V2.587ZM.763 13.807l.062.073l.03-.026c.195.094.414.146.645.146h12a1.5 1.5 0 0 0 1.5-1.5v-10a1.5 1.5 0 0 0-.763-1.307l-.062-.073l-.03.025A1.494 1.494 0 0 0 13.5 1h-12A1.5 1.5 0 0 0 0 2.5v10c0 .56.307 1.05.763 1.307ZM1 12.413L13.148 2H1.5a.5.5 0 0 0-.5.5v9.913Z" clip-rule="evenodd"/>`
	viewVerticalInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M8 2h5.5a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H8V2ZM7 2H1.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5H7V2Zm-7 .5A1.5 1.5 0 0 1 1.5 1h12A1.5 1.5 0 0 1 15 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-10Z" clip-rule="evenodd"/>`
	widthInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M4.818 4.682a.45.45 0 0 1 0 .636L3.086 7.05h8.828l-1.732-1.732a.45.45 0 0 1 .636-.636l2.5 2.5a.45.45 0 0 1 0 .636l-2.5 2.5a.45.45 0 0 1-.636-.636l1.732-1.732H3.086l1.732 1.732a.45.45 0 1 1-.636.636l-2.5-2.5a.45.45 0 0 1 0-.636l2.5-2.5a.45.45 0 0 1 .636 0Z" clip-rule="evenodd"/>`
	zoomInInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M10 6.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm-.691 3.516a4.5 4.5 0 1 1 .707-.707l2.838 2.837a.5.5 0 0 1-.708.708L9.31 10.016ZM4.25 6.5a.5.5 0 0 1 .5-.5H6V4.75a.5.5 0 0 1 1 0V6h1.25a.5.5 0 0 1 0 1H7v1.25a.5.5 0 0 1-1 0V7H4.75a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`
	zoomOutInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M6.5 10a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0 1a4.48 4.48 0 0 0 2.809-.984l2.837 2.838a.5.5 0 0 0 .708-.708L10.016 9.31A4.5 4.5 0 1 0 6.5 11ZM4.75 6a.5.5 0 0 0 0 1h3.5a.5.5 0 0 0 0-1h-3.5Z" clip-rule="evenodd"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Accessibility(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accessibilityInnerSVG).
		Element(children...)
}

func ActivityLog(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(activityLogInnerSVG).
		Element(children...)
}

func AlignBaseline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignBaselineInnerSVG).
		Element(children...)
}

func AlignBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignBottomInnerSVG).
		Element(children...)
}

func AlignCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterInnerSVG).
		Element(children...)
}

func AlignCenterHorizontally(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterHorizontallyInnerSVG).
		Element(children...)
}

func AlignCenterVertically(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignCenterVerticallyInnerSVG).
		Element(children...)
}

func AlignEnd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignEndInnerSVG).
		Element(children...)
}

func AlignHorizontalCenters(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignHorizontalCentersInnerSVG).
		Element(children...)
}

func AlignLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignRightInnerSVG).
		Element(children...)
}

func AlignStart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignStartInnerSVG).
		Element(children...)
}

func AlignStretch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignStretchInnerSVG).
		Element(children...)
}

func AlignTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignTopInnerSVG).
		Element(children...)
}

func AlignVerticalCenters(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alignVerticalCentersInnerSVG).
		Element(children...)
}

func AllSides(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(allSidesInnerSVG).
		Element(children...)
}

func Angle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownInnerSVG).
		Element(children...)
}

func ArrowLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightInnerSVG).
		Element(children...)
}

func ArrowTopLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpInnerSVG).
		Element(children...)
}

func AspectRatio(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aspectRatioInnerSVG).
		Element(children...)
}

func Avatar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(avatarInnerSVG).
		Element(children...)
}

func Backpack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backpackInnerSVG).
		Element(children...)
}

func Badge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(badgeInnerSVG).
		Element(children...)
}

func BarChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartInnerSVG).
		Element(children...)
}

func Bell(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellInnerSVG).
		Element(children...)
}

func BlendingMode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(blendingModeInnerSVG).
		Element(children...)
}

func Bookmark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkInnerSVG).
		Element(children...)
}

func BookmarkFilled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkFilledInnerSVG).
		Element(children...)
}

func BorderAll(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderAllInnerSVG).
		Element(children...)
}

func BorderBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBottomInnerSVG).
		Element(children...)
}

func BorderDashed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderDashedInnerSVG).
		Element(children...)
}

func BorderDotted(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderDottedInnerSVG).
		Element(children...)
}

func BorderLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderLeftInnerSVG).
		Element(children...)
}

func BorderNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderNoneInnerSVG).
		Element(children...)
}

func BorderRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderRightInnerSVG).
		Element(children...)
}

func BorderSolid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderSolidInnerSVG).
		Element(children...)
}

func BorderSplit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderSplitInnerSVG).
		Element(children...)
}

func BorderStyle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderStyleInnerSVG).
		Element(children...)
}

func BorderTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopInnerSVG).
		Element(children...)
}

func BorderWidth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderWidthInnerSVG).
		Element(children...)
}

func Box(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxInnerSVG).
		Element(children...)
}

func BoxModel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxModelInnerSVG).
		Element(children...)
}

func Button(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buttonInnerSVG).
		Element(children...)
}

func Calendar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraInnerSVG).
		Element(children...)
}

func CardStack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardStackInnerSVG).
		Element(children...)
}

func CardStackMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardStackMinusInnerSVG).
		Element(children...)
}

func CardStackPlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardStackPlusInnerSVG).
		Element(children...)
}

func CaretDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretRightInnerSVG).
		Element(children...)
}

func CaretSort(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretSortInnerSVG).
		Element(children...)
}

func CaretUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretUpInnerSVG).
		Element(children...)
}

func ChatBubble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatBubbleInnerSVG).
		Element(children...)
}

func Check(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkInnerSVG).
		Element(children...)
}

func CheckCircled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkCircledInnerSVG).
		Element(children...)
}

func Checkbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxInnerSVG).
		Element(children...)
}

func ChevronDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpInnerSVG).
		Element(children...)
}

func Circle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleInnerSVG).
		Element(children...)
}

func CircleBackslash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleBackslashInnerSVG).
		Element(children...)
}

func Clipboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardInnerSVG).
		Element(children...)
}

func ClipboardCopy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardCopyInnerSVG).
		Element(children...)
}

func Clock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockInnerSVG).
		Element(children...)
}

func Code(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeInnerSVG).
		Element(children...)
}

func CodesandboxLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codesandboxLogoInnerSVG).
		Element(children...)
}

func ColorWheel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(colorWheelInnerSVG).
		Element(children...)
}

func ColumnSpacing(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(columnSpacingInnerSVG).
		Element(children...)
}

func Columns(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(columnsInnerSVG).
		Element(children...)
}

func Commit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commitInnerSVG).
		Element(children...)
}

func ComponentBoolean(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(componentBooleanInnerSVG).
		Element(children...)
}

func ComponentInstance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(componentInstanceInnerSVG).
		Element(children...)
}

func ComponentNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(componentNoneInnerSVG).
		Element(children...)
}

func ComponentOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(componentOneInnerSVG).
		Element(children...)
}

func ComponentPlaceholder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(componentPlaceholderInnerSVG).
		Element(children...)
}

func ComponentTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(componentTwoInnerSVG).
		Element(children...)
}

func Container(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(containerInnerSVG).
		Element(children...)
}

func Cookie(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cookieInnerSVG).
		Element(children...)
}

func Copy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(copyInnerSVG).
		Element(children...)
}

func CornerBottomLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornerTopRightInnerSVG).
		Element(children...)
}

func Corners(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cornersInnerSVG).
		Element(children...)
}

func CountdownTimer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(countdownTimerInnerSVG).
		Element(children...)
}

func CounterClockwiseClock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(counterClockwiseClockInnerSVG).
		Element(children...)
}

func Crop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cropInnerSVG).
		Element(children...)
}

func CrossCircled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crossCircledInnerSVG).
		Element(children...)
}

func CrossOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crossOneInnerSVG).
		Element(children...)
}

func CrossTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crossTwoInnerSVG).
		Element(children...)
}

func CrosshairOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crosshairOneInnerSVG).
		Element(children...)
}

func CrosshairTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crosshairTwoInnerSVG).
		Element(children...)
}

func CrumpledPaper(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crumpledPaperInnerSVG).
		Element(children...)
}

func Cube(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cubeInnerSVG).
		Element(children...)
}

func CursorArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cursorArrowInnerSVG).
		Element(children...)
}

func CursorText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cursorTextInnerSVG).
		Element(children...)
}

func Dash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashInnerSVG).
		Element(children...)
}

func Dashboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashboardInnerSVG).
		Element(children...)
}

func Desktop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(desktopInnerSVG).
		Element(children...)
}

func Dimensions(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dimensionsInnerSVG).
		Element(children...)
}

func Disc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discInnerSVG).
		Element(children...)
}

func DiscordLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discordLogoInnerSVG).
		Element(children...)
}

func DividerHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dividerHorizontalInnerSVG).
		Element(children...)
}

func DividerVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dividerVerticalInnerSVG).
		Element(children...)
}

func Dot(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dotInnerSVG).
		Element(children...)
}

func DotFilled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dotFilledInnerSVG).
		Element(children...)
}

func DotsHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dotsHorizontalInnerSVG).
		Element(children...)
}

func DotsVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dotsVerticalInnerSVG).
		Element(children...)
}

func DoubleArrowDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleArrowDownInnerSVG).
		Element(children...)
}

func DoubleArrowLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleArrowLeftInnerSVG).
		Element(children...)
}

func DoubleArrowRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleArrowRightInnerSVG).
		Element(children...)
}

func DoubleArrowUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doubleArrowUpInnerSVG).
		Element(children...)
}

func Download(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadInnerSVG).
		Element(children...)
}

func DragHandleDotsOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragHandleDotsOneInnerSVG).
		Element(children...)
}

func DragHandleDotsTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragHandleDotsTwoInnerSVG).
		Element(children...)
}

func DragHandleHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragHandleHorizontalInnerSVG).
		Element(children...)
}

func DragHandleVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragHandleVerticalInnerSVG).
		Element(children...)
}

func DrawingPin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(drawingPinInnerSVG).
		Element(children...)
}

func DrawingPinFilled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(drawingPinFilledInnerSVG).
		Element(children...)
}

func DropdownMenu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropdownMenuInnerSVG).
		Element(children...)
}

func Enter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enterInnerSVG).
		Element(children...)
}

func EnterFullScreen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enterFullScreenInnerSVG).
		Element(children...)
}

func EnvelopeClosed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeOpenInnerSVG).
		Element(children...)
}

func Eraser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eraserInnerSVG).
		Element(children...)
}

func ExclamationTriangle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationTriangleInnerSVG).
		Element(children...)
}

func Exit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exitInnerSVG).
		Element(children...)
}

func ExitFullScreen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exitFullScreenInnerSVG).
		Element(children...)
}

func ExternalLink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(externalLinkInnerSVG).
		Element(children...)
}

func EyeClosed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeClosedInnerSVG).
		Element(children...)
}

func EyeNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeNoneInnerSVG).
		Element(children...)
}

func EyeOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeOpenInnerSVG).
		Element(children...)
}

func Face(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceInnerSVG).
		Element(children...)
}

func FigmaLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(figmaLogoInnerSVG).
		Element(children...)
}

func File(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileTextInnerSVG).
		Element(children...)
}

func FontBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fontBoldInnerSVG).
		Element(children...)
}

func FontFamily(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fontFamilyInnerSVG).
		Element(children...)
}

func FontItalic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fontItalicInnerSVG).
		Element(children...)
}

func FontRoman(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fontRomanInnerSVG).
		Element(children...)
}

func FontSize(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fontSizeInnerSVG).
		Element(children...)
}

func FontStyle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fontStyleInnerSVG).
		Element(children...)
}

func Frame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameInnerSVG).
		Element(children...)
}

func FramerLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(framerLogoInnerSVG).
		Element(children...)
}

func Gear(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gearInnerSVG).
		Element(children...)
}

func GithubLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(githubLogoInnerSVG).
		Element(children...)
}

func Globe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gridInnerSVG).
		Element(children...)
}

func Group(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(groupInnerSVG).
		Element(children...)
}

func HalfOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(halfOneInnerSVG).
		Element(children...)
}

func HalfTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(halfTwoInnerSVG).
		Element(children...)
}

func HamburgerMenu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hamburgerMenuInnerSVG).
		Element(children...)
}

func Hand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(handInnerSVG).
		Element(children...)
}

func Heading(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headingInnerSVG).
		Element(children...)
}

func Heart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartInnerSVG).
		Element(children...)
}

func HeartFilled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartFilledInnerSVG).
		Element(children...)
}

func Height(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heightInnerSVG).
		Element(children...)
}

func HobbyKnife(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hobbyKnifeInnerSVG).
		Element(children...)
}

func Home(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeInnerSVG).
		Element(children...)
}

func IconjarLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iconjarLogoInnerSVG).
		Element(children...)
}

func IdCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(idCardInnerSVG).
		Element(children...)
}

func Image(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageInnerSVG).
		Element(children...)
}

func InfoCircled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoCircledInnerSVG).
		Element(children...)
}

func Input(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inputInnerSVG).
		Element(children...)
}

func InstagramLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(instagramLogoInnerSVG).
		Element(children...)
}

func JustifyCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(justifyCenterInnerSVG).
		Element(children...)
}

func JustifyEnd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(justifyEndInnerSVG).
		Element(children...)
}

func JustifyStart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(justifyStartInnerSVG).
		Element(children...)
}

func JustifyStretch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(justifyStretchInnerSVG).
		Element(children...)
}

func Keyboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyboardInnerSVG).
		Element(children...)
}

func LapTimer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lapTimerInnerSVG).
		Element(children...)
}

func Laptop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutInnerSVG).
		Element(children...)
}

func LetterCaseCapitalize(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(letterCaseCapitalizeInnerSVG).
		Element(children...)
}

func LetterCaseLowercase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(letterCaseLowercaseInnerSVG).
		Element(children...)
}

func LetterCaseToggle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(letterCaseToggleInnerSVG).
		Element(children...)
}

func LetterCaseUppercase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(letterCaseUppercaseInnerSVG).
		Element(children...)
}

func LetterSpacing(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(letterSpacingInnerSVG).
		Element(children...)
}

func LightningBolt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightningBoltInnerSVG).
		Element(children...)
}

func LineHeight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineHeightInnerSVG).
		Element(children...)
}

func LinkBreakOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkBreakOneInnerSVG).
		Element(children...)
}

func LinkBreakTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkBreakTwoInnerSVG).
		Element(children...)
}

func LinkNoneOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkNoneOneInnerSVG).
		Element(children...)
}

func LinkNoneTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkNoneTwoInnerSVG).
		Element(children...)
}

func LinkOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkOneInnerSVG).
		Element(children...)
}

func LinkTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkTwoInnerSVG).
		Element(children...)
}

func LinkedinLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkedinLogoInnerSVG).
		Element(children...)
}

func ListBullet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listBulletInnerSVG).
		Element(children...)
}

func LockClosed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockClosedInnerSVG).
		Element(children...)
}

func LockOpenOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockOpenOneInnerSVG).
		Element(children...)
}

func LockOpenTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockOpenTwoInnerSVG).
		Element(children...)
}

func Loop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loopInnerSVG).
		Element(children...)
}

func MagicWand(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magicWandInnerSVG).
		Element(children...)
}

func MagnifyingGlass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magnifyingGlassInnerSVG).
		Element(children...)
}

func Margin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(marginInnerSVG).
		Element(children...)
}

func MaskOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maskOffInnerSVG).
		Element(children...)
}

func MaskOn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maskOnInnerSVG).
		Element(children...)
}

func Minus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusInnerSVG).
		Element(children...)
}

func MinusCircled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircledInnerSVG).
		Element(children...)
}

func Mix(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mixInnerSVG).
		Element(children...)
}

func MixerHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mixerHorizontalInnerSVG).
		Element(children...)
}

func MixerVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mixerVerticalInnerSVG).
		Element(children...)
}

func Mobile(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mobileInnerSVG).
		Element(children...)
}

func ModulzLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(modulzLogoInnerSVG).
		Element(children...)
}

func Moon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveInnerSVG).
		Element(children...)
}

func NotionLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notionLogoInnerSVG).
		Element(children...)
}

func Opacity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(opacityInnerSVG).
		Element(children...)
}

func OpenInNewWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openInNewWindowInnerSVG).
		Element(children...)
}

func Overline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(overlineInnerSVG).
		Element(children...)
}

func Padding(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paddingInnerSVG).
		Element(children...)
}

func PaperPlane(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperPlaneInnerSVG).
		Element(children...)
}

func Pause(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseInnerSVG).
		Element(children...)
}

func PencilOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pencilOneInnerSVG).
		Element(children...)
}

func PencilTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pencilTwoInnerSVG).
		Element(children...)
}

func Person(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(personInnerSVG).
		Element(children...)
}

func PieChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartInnerSVG).
		Element(children...)
}

func Pilcrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pilcrowInnerSVG).
		Element(children...)
}

func PinBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinBottomInnerSVG).
		Element(children...)
}

func PinLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinLeftInnerSVG).
		Element(children...)
}

func PinRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinRightInnerSVG).
		Element(children...)
}

func PinTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinTopInnerSVG).
		Element(children...)
}

func Play(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playInnerSVG).
		Element(children...)
}

func Plus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusInnerSVG).
		Element(children...)
}

func PlusCircled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircledInnerSVG).
		Element(children...)
}

func QuestionMark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionMarkInnerSVG).
		Element(children...)
}

func QuestionMarkCircled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionMarkCircledInnerSVG).
		Element(children...)
}

func Quote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quoteInnerSVG).
		Element(children...)
}

func Radiobutton(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radiobuttonInnerSVG).
		Element(children...)
}

func Reader(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(readerInnerSVG).
		Element(children...)
}

func Reload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reloadInnerSVG).
		Element(children...)
}

func Reset(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resetInnerSVG).
		Element(children...)
}

func Resume(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resumeInnerSVG).
		Element(children...)
}

func Rocket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rocketInnerSVG).
		Element(children...)
}

func RotateCounterClockwise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateCounterClockwiseInnerSVG).
		Element(children...)
}

func RowSpacing(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rowSpacingInnerSVG).
		Element(children...)
}

func Rows(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rowsInnerSVG).
		Element(children...)
}

func RulerHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rulerHorizontalInnerSVG).
		Element(children...)
}

func RulerSquare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rulerSquareInnerSVG).
		Element(children...)
}

func Scissors(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scissorsInnerSVG).
		Element(children...)
}

func Section(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionInnerSVG).
		Element(children...)
}

func SewingPin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sewingPinInnerSVG).
		Element(children...)
}

func SewingPinFilled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sewingPinFilledInnerSVG).
		Element(children...)
}

func Shadow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shadowInnerSVG).
		Element(children...)
}

func ShadowInner(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shadowInnerInnerSVG).
		Element(children...)
}

func ShadowNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shadowNoneInnerSVG).
		Element(children...)
}

func ShadowOuter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shadowOuterInnerSVG).
		Element(children...)
}

func ShareOne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareOneInnerSVG).
		Element(children...)
}

func ShareTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareTwoInnerSVG).
		Element(children...)
}

func Shuffle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shuffleInnerSVG).
		Element(children...)
}

func Size(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sizeInnerSVG).
		Element(children...)
}

func SketchLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sketchLogoInnerSVG).
		Element(children...)
}

func Slash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slashInnerSVG).
		Element(children...)
}

func Slider(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sliderInnerSVG).
		Element(children...)
}

func SpaceBetweenHorizontally(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spaceBetweenHorizontallyInnerSVG).
		Element(children...)
}

func SpaceBetweenVertically(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spaceBetweenVerticallyInnerSVG).
		Element(children...)
}

func SpaceEvenlyHorizontally(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spaceEvenlyHorizontallyInnerSVG).
		Element(children...)
}

func SpaceEvenlyVertically(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spaceEvenlyVerticallyInnerSVG).
		Element(children...)
}

func SpeakerLoud(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speakerLoudInnerSVG).
		Element(children...)
}

func SpeakerModerate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speakerModerateInnerSVG).
		Element(children...)
}

func SpeakerOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speakerOffInnerSVG).
		Element(children...)
}

func SpeakerQuiet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speakerQuietInnerSVG).
		Element(children...)
}

func Square(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareInnerSVG).
		Element(children...)
}

func Stack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stackInnerSVG).
		Element(children...)
}

func Star(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starInnerSVG).
		Element(children...)
}

func StarFilled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starFilledInnerSVG).
		Element(children...)
}

func StitchesLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stitchesLogoInnerSVG).
		Element(children...)
}

func Stop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopInnerSVG).
		Element(children...)
}

func Stopwatch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopwatchInnerSVG).
		Element(children...)
}

func StretchHorizontally(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stretchHorizontallyInnerSVG).
		Element(children...)
}

func StretchVertically(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stretchVerticallyInnerSVG).
		Element(children...)
}

func Strikethrough(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(strikethroughInnerSVG).
		Element(children...)
}

func Sun(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunInnerSVG).
		Element(children...)
}

func Switch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(switchInnerSVG).
		Element(children...)
}

func Symbol(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(symbolInnerSVG).
		Element(children...)
}

func Table(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableInnerSVG).
		Element(children...)
}

func Target(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(targetInnerSVG).
		Element(children...)
}

func Text(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textInnerSVG).
		Element(children...)
}

func TextAlignBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignBottomInnerSVG).
		Element(children...)
}

func TextAlignCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignCenterInnerSVG).
		Element(children...)
}

func TextAlignJustify(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignJustifyInnerSVG).
		Element(children...)
}

func TextAlignLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignLeftInnerSVG).
		Element(children...)
}

func TextAlignMiddle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignMiddleInnerSVG).
		Element(children...)
}

func TextAlignRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignRightInnerSVG).
		Element(children...)
}

func TextAlignTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignTopInnerSVG).
		Element(children...)
}

func TextNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textNoneInnerSVG).
		Element(children...)
}

func ThickArrowDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thickArrowDownInnerSVG).
		Element(children...)
}

func ThickArrowLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thickArrowLeftInnerSVG).
		Element(children...)
}

func ThickArrowRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thickArrowRightInnerSVG).
		Element(children...)
}

func ThickArrowUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thickArrowUpInnerSVG).
		Element(children...)
}

func Timer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timerInnerSVG).
		Element(children...)
}

func Tokens(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tokensInnerSVG).
		Element(children...)
}

func TrackNext(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trackNextInnerSVG).
		Element(children...)
}

func TrackPrevious(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trackPreviousInnerSVG).
		Element(children...)
}

func Transform(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transformInnerSVG).
		Element(children...)
}

func TransparencyGrid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transparencyGridInnerSVG).
		Element(children...)
}

func Trash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashInnerSVG).
		Element(children...)
}

func TriangleDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleDownInnerSVG).
		Element(children...)
}

func TriangleLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleLeftInnerSVG).
		Element(children...)
}

func TriangleRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleRightInnerSVG).
		Element(children...)
}

func TriangleUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleUpInnerSVG).
		Element(children...)
}

func TwitterLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twitterLogoInnerSVG).
		Element(children...)
}

func Underline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(underlineInnerSVG).
		Element(children...)
}

func Update(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(updateInnerSVG).
		Element(children...)
}

func Upload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadInnerSVG).
		Element(children...)
}

func Value(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(valueInnerSVG).
		Element(children...)
}

func ValueNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(valueNoneInnerSVG).
		Element(children...)
}

func VercelLogo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vercelLogoInnerSVG).
		Element(children...)
}

func Video(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(videoInnerSVG).
		Element(children...)
}

func ViewGrid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewGridInnerSVG).
		Element(children...)
}

func ViewHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewHorizontalInnerSVG).
		Element(children...)
}

func ViewNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewNoneInnerSVG).
		Element(children...)
}

func ViewVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewVerticalInnerSVG).
		Element(children...)
}

func Width(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(widthInnerSVG).
		Element(children...)
}

func ZoomIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
			engine.NewAttribute("viewBox", "0 0 15 15"),
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
	case "activity-log":
		return ActivityLog(), nil
	case "align-baseline":
		return AlignBaseline(), nil
	case "align-bottom":
		return AlignBottom(), nil
	case "align-center":
		return AlignCenter(), nil
	case "align-center-horizontally":
		return AlignCenterHorizontally(), nil
	case "align-center-vertically":
		return AlignCenterVertically(), nil
	case "align-end":
		return AlignEnd(), nil
	case "align-horizontal-centers":
		return AlignHorizontalCenters(), nil
	case "align-left":
		return AlignLeft(), nil
	case "align-right":
		return AlignRight(), nil
	case "align-start":
		return AlignStart(), nil
	case "align-stretch":
		return AlignStretch(), nil
	case "align-top":
		return AlignTop(), nil
	case "align-vertical-centers":
		return AlignVerticalCenters(), nil
	case "all-sides":
		return AllSides(), nil
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
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-top-left":
		return ArrowTopLeft(), nil
	case "arrow-top-right":
		return ArrowTopRight(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "aspect-ratio":
		return AspectRatio(), nil
	case "avatar":
		return Avatar(), nil
	case "backpack":
		return Backpack(), nil
	case "badge":
		return Badge(), nil
	case "bar-chart":
		return BarChart(), nil
	case "bell":
		return Bell(), nil
	case "blending-mode":
		return BlendingMode(), nil
	case "bookmark":
		return Bookmark(), nil
	case "bookmark-filled":
		return BookmarkFilled(), nil
	case "border-all":
		return BorderAll(), nil
	case "border-bottom":
		return BorderBottom(), nil
	case "border-dashed":
		return BorderDashed(), nil
	case "border-dotted":
		return BorderDotted(), nil
	case "border-left":
		return BorderLeft(), nil
	case "border-none":
		return BorderNone(), nil
	case "border-right":
		return BorderRight(), nil
	case "border-solid":
		return BorderSolid(), nil
	case "border-split":
		return BorderSplit(), nil
	case "border-style":
		return BorderStyle(), nil
	case "border-top":
		return BorderTop(), nil
	case "border-width":
		return BorderWidth(), nil
	case "box":
		return Box(), nil
	case "box-model":
		return BoxModel(), nil
	case "button":
		return Button(), nil
	case "calendar":
		return Calendar(), nil
	case "camera":
		return Camera(), nil
	case "card-stack":
		return CardStack(), nil
	case "card-stack-minus":
		return CardStackMinus(), nil
	case "card-stack-plus":
		return CardStackPlus(), nil
	case "caret-down":
		return CaretDown(), nil
	case "caret-left":
		return CaretLeft(), nil
	case "caret-right":
		return CaretRight(), nil
	case "caret-sort":
		return CaretSort(), nil
	case "caret-up":
		return CaretUp(), nil
	case "chat-bubble":
		return ChatBubble(), nil
	case "check":
		return Check(), nil
	case "check-circled":
		return CheckCircled(), nil
	case "checkbox":
		return Checkbox(), nil
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
	case "circle-backslash":
		return CircleBackslash(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clipboard-copy":
		return ClipboardCopy(), nil
	case "clock":
		return Clock(), nil
	case "code":
		return Code(), nil
	case "codesandbox-logo":
		return CodesandboxLogo(), nil
	case "color-wheel":
		return ColorWheel(), nil
	case "column-spacing":
		return ColumnSpacing(), nil
	case "columns":
		return Columns(), nil
	case "commit":
		return Commit(), nil
	case "component-boolean":
		return ComponentBoolean(), nil
	case "component-instance":
		return ComponentInstance(), nil
	case "component-none":
		return ComponentNone(), nil
	case "component-1":
		return ComponentOne(), nil
	case "component-placeholder":
		return ComponentPlaceholder(), nil
	case "component-2":
		return ComponentTwo(), nil
	case "container":
		return Container(), nil
	case "cookie":
		return Cookie(), nil
	case "copy":
		return Copy(), nil
	case "corner-bottom-left":
		return CornerBottomLeft(), nil
	case "corner-bottom-right":
		return CornerBottomRight(), nil
	case "corner-top-left":
		return CornerTopLeft(), nil
	case "corner-top-right":
		return CornerTopRight(), nil
	case "corners":
		return Corners(), nil
	case "countdown-timer":
		return CountdownTimer(), nil
	case "counter-clockwise-clock":
		return CounterClockwiseClock(), nil
	case "crop":
		return Crop(), nil
	case "cross-circled":
		return CrossCircled(), nil
	case "cross-1":
		return CrossOne(), nil
	case "cross-2":
		return CrossTwo(), nil
	case "crosshair-1":
		return CrosshairOne(), nil
	case "crosshair-2":
		return CrosshairTwo(), nil
	case "crumpled-paper":
		return CrumpledPaper(), nil
	case "cube":
		return Cube(), nil
	case "cursor-arrow":
		return CursorArrow(), nil
	case "cursor-text":
		return CursorText(), nil
	case "dash":
		return Dash(), nil
	case "dashboard":
		return Dashboard(), nil
	case "desktop":
		return Desktop(), nil
	case "dimensions":
		return Dimensions(), nil
	case "disc":
		return Disc(), nil
	case "discord-logo":
		return DiscordLogo(), nil
	case "divider-horizontal":
		return DividerHorizontal(), nil
	case "divider-vertical":
		return DividerVertical(), nil
	case "dot":
		return Dot(), nil
	case "dot-filled":
		return DotFilled(), nil
	case "dots-horizontal":
		return DotsHorizontal(), nil
	case "dots-vertical":
		return DotsVertical(), nil
	case "double-arrow-down":
		return DoubleArrowDown(), nil
	case "double-arrow-left":
		return DoubleArrowLeft(), nil
	case "double-arrow-right":
		return DoubleArrowRight(), nil
	case "double-arrow-up":
		return DoubleArrowUp(), nil
	case "download":
		return Download(), nil
	case "drag-handle-dots-1":
		return DragHandleDotsOne(), nil
	case "drag-handle-dots-2":
		return DragHandleDotsTwo(), nil
	case "drag-handle-horizontal":
		return DragHandleHorizontal(), nil
	case "drag-handle-vertical":
		return DragHandleVertical(), nil
	case "drawing-pin":
		return DrawingPin(), nil
	case "drawing-pin-filled":
		return DrawingPinFilled(), nil
	case "dropdown-menu":
		return DropdownMenu(), nil
	case "enter":
		return Enter(), nil
	case "enter-full-screen":
		return EnterFullScreen(), nil
	case "envelope-closed":
		return EnvelopeClosed(), nil
	case "envelope-open":
		return EnvelopeOpen(), nil
	case "eraser":
		return Eraser(), nil
	case "exclamation-triangle":
		return ExclamationTriangle(), nil
	case "exit":
		return Exit(), nil
	case "exit-full-screen":
		return ExitFullScreen(), nil
	case "external-link":
		return ExternalLink(), nil
	case "eye-closed":
		return EyeClosed(), nil
	case "eye-none":
		return EyeNone(), nil
	case "eye-open":
		return EyeOpen(), nil
	case "face":
		return Face(), nil
	case "figma-logo":
		return FigmaLogo(), nil
	case "file":
		return File(), nil
	case "file-minus":
		return FileMinus(), nil
	case "file-plus":
		return FilePlus(), nil
	case "file-text":
		return FileText(), nil
	case "font-bold":
		return FontBold(), nil
	case "font-family":
		return FontFamily(), nil
	case "font-italic":
		return FontItalic(), nil
	case "font-roman":
		return FontRoman(), nil
	case "font-size":
		return FontSize(), nil
	case "font-style":
		return FontStyle(), nil
	case "frame":
		return Frame(), nil
	case "framer-logo":
		return FramerLogo(), nil
	case "gear":
		return Gear(), nil
	case "github-logo":
		return GithubLogo(), nil
	case "globe":
		return Globe(), nil
	case "grid":
		return Grid(), nil
	case "group":
		return Group(), nil
	case "half-1":
		return HalfOne(), nil
	case "half-2":
		return HalfTwo(), nil
	case "hamburger-menu":
		return HamburgerMenu(), nil
	case "hand":
		return Hand(), nil
	case "heading":
		return Heading(), nil
	case "heart":
		return Heart(), nil
	case "heart-filled":
		return HeartFilled(), nil
	case "height":
		return Height(), nil
	case "hobby-knife":
		return HobbyKnife(), nil
	case "home":
		return Home(), nil
	case "iconjar-logo":
		return IconjarLogo(), nil
	case "id-card":
		return IdCard(), nil
	case "image":
		return Image(), nil
	case "info-circled":
		return InfoCircled(), nil
	case "input":
		return Input(), nil
	case "instagram-logo":
		return InstagramLogo(), nil
	case "justify-center":
		return JustifyCenter(), nil
	case "justify-end":
		return JustifyEnd(), nil
	case "justify-start":
		return JustifyStart(), nil
	case "justify-stretch":
		return JustifyStretch(), nil
	case "keyboard":
		return Keyboard(), nil
	case "lap-timer":
		return LapTimer(), nil
	case "laptop":
		return Laptop(), nil
	case "layers":
		return Layers(), nil
	case "layout":
		return Layout(), nil
	case "letter-case-capitalize":
		return LetterCaseCapitalize(), nil
	case "letter-case-lowercase":
		return LetterCaseLowercase(), nil
	case "letter-case-toggle":
		return LetterCaseToggle(), nil
	case "letter-case-uppercase":
		return LetterCaseUppercase(), nil
	case "letter-spacing":
		return LetterSpacing(), nil
	case "lightning-bolt":
		return LightningBolt(), nil
	case "line-height":
		return LineHeight(), nil
	case "link-break-1":
		return LinkBreakOne(), nil
	case "link-break-2":
		return LinkBreakTwo(), nil
	case "link-none-1":
		return LinkNoneOne(), nil
	case "link-none-2":
		return LinkNoneTwo(), nil
	case "link-1":
		return LinkOne(), nil
	case "link-2":
		return LinkTwo(), nil
	case "linkedin-logo":
		return LinkedinLogo(), nil
	case "list-bullet":
		return ListBullet(), nil
	case "lock-closed":
		return LockClosed(), nil
	case "lock-open-1":
		return LockOpenOne(), nil
	case "lock-open-2":
		return LockOpenTwo(), nil
	case "loop":
		return Loop(), nil
	case "magic-wand":
		return MagicWand(), nil
	case "magnifying-glass":
		return MagnifyingGlass(), nil
	case "margin":
		return Margin(), nil
	case "mask-off":
		return MaskOff(), nil
	case "mask-on":
		return MaskOn(), nil
	case "minus":
		return Minus(), nil
	case "minus-circled":
		return MinusCircled(), nil
	case "mix":
		return Mix(), nil
	case "mixer-horizontal":
		return MixerHorizontal(), nil
	case "mixer-vertical":
		return MixerVertical(), nil
	case "mobile":
		return Mobile(), nil
	case "modulz-logo":
		return ModulzLogo(), nil
	case "moon":
		return Moon(), nil
	case "move":
		return Move(), nil
	case "notion-logo":
		return NotionLogo(), nil
	case "opacity":
		return Opacity(), nil
	case "open-in-new-window":
		return OpenInNewWindow(), nil
	case "overline":
		return Overline(), nil
	case "padding":
		return Padding(), nil
	case "paper-plane":
		return PaperPlane(), nil
	case "pause":
		return Pause(), nil
	case "pencil-1":
		return PencilOne(), nil
	case "pencil-2":
		return PencilTwo(), nil
	case "person":
		return Person(), nil
	case "pie-chart":
		return PieChart(), nil
	case "pilcrow":
		return Pilcrow(), nil
	case "pin-bottom":
		return PinBottom(), nil
	case "pin-left":
		return PinLeft(), nil
	case "pin-right":
		return PinRight(), nil
	case "pin-top":
		return PinTop(), nil
	case "play":
		return Play(), nil
	case "plus":
		return Plus(), nil
	case "plus-circled":
		return PlusCircled(), nil
	case "question-mark":
		return QuestionMark(), nil
	case "question-mark-circled":
		return QuestionMarkCircled(), nil
	case "quote":
		return Quote(), nil
	case "radiobutton":
		return Radiobutton(), nil
	case "reader":
		return Reader(), nil
	case "reload":
		return Reload(), nil
	case "reset":
		return Reset(), nil
	case "resume":
		return Resume(), nil
	case "rocket":
		return Rocket(), nil
	case "rotate-counter-clockwise":
		return RotateCounterClockwise(), nil
	case "row-spacing":
		return RowSpacing(), nil
	case "rows":
		return Rows(), nil
	case "ruler-horizontal":
		return RulerHorizontal(), nil
	case "ruler-square":
		return RulerSquare(), nil
	case "scissors":
		return Scissors(), nil
	case "section":
		return Section(), nil
	case "sewing-pin":
		return SewingPin(), nil
	case "sewing-pin-filled":
		return SewingPinFilled(), nil
	case "shadow":
		return Shadow(), nil
	case "shadow-inner":
		return ShadowInner(), nil
	case "shadow-none":
		return ShadowNone(), nil
	case "shadow-outer":
		return ShadowOuter(), nil
	case "share-1":
		return ShareOne(), nil
	case "share-2":
		return ShareTwo(), nil
	case "shuffle":
		return Shuffle(), nil
	case "size":
		return Size(), nil
	case "sketch-logo":
		return SketchLogo(), nil
	case "slash":
		return Slash(), nil
	case "slider":
		return Slider(), nil
	case "space-between-horizontally":
		return SpaceBetweenHorizontally(), nil
	case "space-between-vertically":
		return SpaceBetweenVertically(), nil
	case "space-evenly-horizontally":
		return SpaceEvenlyHorizontally(), nil
	case "space-evenly-vertically":
		return SpaceEvenlyVertically(), nil
	case "speaker-loud":
		return SpeakerLoud(), nil
	case "speaker-moderate":
		return SpeakerModerate(), nil
	case "speaker-off":
		return SpeakerOff(), nil
	case "speaker-quiet":
		return SpeakerQuiet(), nil
	case "square":
		return Square(), nil
	case "stack":
		return Stack(), nil
	case "star":
		return Star(), nil
	case "star-filled":
		return StarFilled(), nil
	case "stitches-logo":
		return StitchesLogo(), nil
	case "stop":
		return Stop(), nil
	case "stopwatch":
		return Stopwatch(), nil
	case "stretch-horizontally":
		return StretchHorizontally(), nil
	case "stretch-vertically":
		return StretchVertically(), nil
	case "strikethrough":
		return Strikethrough(), nil
	case "sun":
		return Sun(), nil
	case "switch":
		return Switch(), nil
	case "symbol":
		return Symbol(), nil
	case "table":
		return Table(), nil
	case "target":
		return Target(), nil
	case "text":
		return Text(), nil
	case "text-align-bottom":
		return TextAlignBottom(), nil
	case "text-align-center":
		return TextAlignCenter(), nil
	case "text-align-justify":
		return TextAlignJustify(), nil
	case "text-align-left":
		return TextAlignLeft(), nil
	case "text-align-middle":
		return TextAlignMiddle(), nil
	case "text-align-right":
		return TextAlignRight(), nil
	case "text-align-top":
		return TextAlignTop(), nil
	case "text-none":
		return TextNone(), nil
	case "thick-arrow-down":
		return ThickArrowDown(), nil
	case "thick-arrow-left":
		return ThickArrowLeft(), nil
	case "thick-arrow-right":
		return ThickArrowRight(), nil
	case "thick-arrow-up":
		return ThickArrowUp(), nil
	case "timer":
		return Timer(), nil
	case "tokens":
		return Tokens(), nil
	case "track-next":
		return TrackNext(), nil
	case "track-previous":
		return TrackPrevious(), nil
	case "transform":
		return Transform(), nil
	case "transparency-grid":
		return TransparencyGrid(), nil
	case "trash":
		return Trash(), nil
	case "triangle-down":
		return TriangleDown(), nil
	case "triangle-left":
		return TriangleLeft(), nil
	case "triangle-right":
		return TriangleRight(), nil
	case "triangle-up":
		return TriangleUp(), nil
	case "twitter-logo":
		return TwitterLogo(), nil
	case "underline":
		return Underline(), nil
	case "update":
		return Update(), nil
	case "upload":
		return Upload(), nil
	case "value":
		return Value(), nil
	case "value-none":
		return ValueNone(), nil
	case "vercel-logo":
		return VercelLogo(), nil
	case "video":
		return Video(), nil
	case "view-grid":
		return ViewGrid(), nil
	case "view-horizontal":
		return ViewHorizontal(), nil
	case "view-none":
		return ViewNone(), nil
	case "view-vertical":
		return ViewVertical(), nil
	case "width":
		return Width(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in radix_icons icon set", name)
	}
}
