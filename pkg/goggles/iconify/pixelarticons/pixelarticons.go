package pixelarticons

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	abTestingInnerSVG                  = `<path fill="currentColor" d="M3 3h6v2H5v2h4v2H5v2h4v2H3V3zm6 8h2V9H9v2zm0-4h2V5H9v2zm4 4h8v10h-2v-4h-4v4h-2V11zm2 4h4v-2h-4v2zm0-12h6v6h-2V5h-4V3zM3 15h2v4h4v2H3v-6z"/>`
	acInnerSVG                         = `<path fill="currentColor" d="M13 2h-2v4H9V4H7v2h2v2h2v3H8V9H6V7H4v2h2v2H2v2h4v2H4v2h2v-2h2v-2h3v3H9v2H7v2h2v-2h2v4h2v-4h2v2h2v-2h-2v-2h-2v-3h3v2h2v2h2v-2h-2v-2h4v-2h-4V9h2V7h-2v2h-2v2h-3V8h2V6h2V4h-2v2h-2V2z"/>`
	addBoxInnerSVG                     = `<path fill="currentColor" d="M3 3h18v18H3V3zm16 16V5H5v14h14zm-6-8h4v2h-4v4h-2v-4H7v-2h4V7h2v4z"/>`
	addBoxMultipleInnerSVG             = `<path fill="currentColor" d="M3 3h14v14H3V3zm12 12V5H5v10h10zm-8 6v-2h12V7h2v14H7zm4-12h2v2h-2v2H9v-2H7V9h2V7h2v2z"/>`
	addColInnerSVG                     = `<path fill="currentColor" d="M2 2h10v20H2v-2h8v-4H2v-2h8v-4H2V8h8V4H2V2zm17 9h3v2h-3v3h-2v-3h-3v-2h3V8h2v3z"/>`
	addGridInnerSVG                    = `<path fill="currentColor" d="M3 3h8v8H3V3zm6 6V5H5v4h4zm9 4h-2v3h-3v2h3v3h2v-3h3v-2h-3v-3zM15 3h6v8h-8V3h2zm4 6V5h-4v4h4zM5 13h6v8H3v-8h2zm4 6v-4H5v4h4z"/>`
	addRowInnerSVG                     = `<path fill="currentColor" d="M4 10V2H2v10h20V2h-2v8h-4V2h-2v8h-4V2H8v8H4zm9 9v3h-2v-3H8v-2h3v-3h2v3h3v2h-3z"/>`
	alertInnerSVG                      = `<path fill="currentColor" d="M13 1h-2v2H9v2H7v2H5v2H3v2H1v2h2v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2V5h-2V3h-2V1zm0 2v2h2v2h2v2h2v2h2v2h-2v2h-2v2h-2v2h-2v2h-2v-2H9v-2H7v-2H5v-2H3v-2h2V9h2V7h2V5h2V3h2zm0 4h-2v6h2V7zm0 8h-2v2h2v-2z"/>`
	alignCenterInnerSVG                = `<path fill="currentColor" d="M20 5H4v2h16V5zm-4 4H8v2h8V9zM4 13h16v2H4v-2zm12 4H8v2h8v-2z"/>`
	alignJustifyInnerSVG               = `<path fill="currentColor" d="M20 5H4v2h16V5zm0 4H4v2h16V9zM4 13h16v2H4v-2zm16 4H4v2h16v-2z"/>`
	alignLeftInnerSVG                  = `<path fill="currentColor" d="M20 5H4v2h16V5zm-8 4H4v2h8V9zm8 4v2H4v-2h16zm-8 4H4v2h8v-2z"/>`
	alignRightInnerSVG                 = `<path fill="currentColor" d="M4 5h16v2H4V5zm8 4h8v2h-8V9zm-8 4v2h16v-2H4zm8 4h8v2h-8v-2z"/>`
	analyticsInnerSVG                  = `<path fill="currentColor" d="M3 3h18v18H3V3zm16 2H5v14h14V5zM7 12h2v5H7v-5zm10-5h-2v10h2V7zm-6 3h2v2h-2v-2zm2 4h-2v3h2v-3z"/>`
	anchorInnerSVG                     = `<path fill="currentColor" d="M14 3h-4v2H8v4h2v2h1v8H6v-4h2v-2H4v6h2v2h12v-2h2v-6h-4v2h2v4h-5v-8h1V9h2V5h-2V3zm0 2v4h-4V5h4z"/>`
	androidInnerSVG                    = `<path fill="currentColor" d="M2 5h2v2H2V5zm4 4H4V7h2v2zm2 0H6v2H4v2H2v6h20v-6h-2v-2h-2V9h2V7h2V5h-2v2h-2v2h-2V7H8v2zm0 0h8v2h2v2h2v4H4v-4h2v-2h2V9zm2 4H8v2h2v-2zm4 0h2v2h-2v-2z"/>`
	animationInnerSVG                  = `<path fill="currentColor" d="M4 2H2v12h2V4h10V2H4zm2 4h12v2H8v10H6V6zm4 4h12v12H10V10zm10 10v-8h-8v8h8z"/>`
	archiveInnerSVG                    = `<path fill="currentColor" d="M22 4H2v6h2v10h16V10h2V4zM6 10h12v8H6v-8zm14-4v2H4V6h16zm-5 6H9v2h6v-2z"/>`
	arrowBarDownInnerSVG               = `<path fill="currentColor" d="M11 4h2v8h2v2h-2v2h-2v-2H9v-2h2V4zm-2 8H7v-2h2v2zm6 0v-2h2v2h-2zM4 18h16v2H4v-2z"/>`
	arrowBarLeftInnerSVG               = `<path fill="currentColor" d="M6 4v16H4V4h2zm14 7v2h-8v2h-2v-2H8v-2h2V9h2v2h8zm-8-2V7h2v2h-2zm0 6h2v2h-2v-2z"/>`
	arrowBarRightInnerSVG              = `<path fill="currentColor" d="M18 4v16h2V4h-2zM4 11v2h8v2h-2v2h2v-2h2v-2h2v-2h-2V9h-2V7h-2v2h2v2H4z"/>`
	arrowBarUpInnerSVG                 = `<path fill="currentColor" d="M4 6h16V4H4v2zm7 14h2v-8h2v2h2v-2h-2v-2h-2V8h-2v2H9v2H7v2h2v-2h2v8z"/>`
	arrowDownInnerSVG                  = `<path fill="currentColor" d="M11 4h2v12h2v2h-2v2h-2v-2H9v-2h2V4zM7 14v2h2v-2H7zm0 0v-2H5v2h2zm10 0v2h-2v-2h2zm0 0v-2h2v2h-2z"/>`
	arrowDownBoxInnerSVG               = `<path fill="currentColor" d="M3 3h18v18H3V3zm16 16V5H5v14h14zM11 7h2v6h2v2h-2v2h-2v-2H9v-2h2V7zm-2 4v2H7v-2h2zm8 0h-2v2h2v-2z"/>`
	arrowLeftInnerSVG                  = `<path fill="currentColor" d="M20 11v2H8v2H6v-2H4v-2h2V9h2v2h12zM10 7H8v2h2V7zm0 0h2V5h-2v2zm0 10H8v-2h2v2zm0 0h2v2h-2v-2z"/>`
	arrowLeftBoxInnerSVG               = `<path fill="currentColor" d="M21 3v18H3V3h18zM5 19h14V5H5v14zm12-8v2h-6v2H9v-2H7v-2h2V9h2v2h6zm-4-2h-2V7h2v2zm0 8v-2h-2v2h2z"/>`
	arrowRightInnerSVG                 = `<path fill="currentColor" d="M4 11v2h12v2h2v-2h2v-2h-2V9h-2v2H4zm10-4h2v2h-2V7zm0 0h-2V5h2v2zm0 10h2v-2h-2v2zm0 0h-2v2h2v-2z"/>`
	arrowRightBoxInnerSVG              = `<path fill="currentColor" d="M3 21V3h18v18H3zM19 5H5v14h14V5zM7 13v-2h6V9h2v2h2v2h-2v2h-2v-2H7zm4 2h2v2h-2v-2zm0-8v2h2V7h-2z"/>`
	arrowUpInnerSVG                    = `<path fill="currentColor" d="M11 20h2V8h2V6h-2V4h-2v2H9v2h2v12zM7 10V8h2v2H7zm0 0v2H5v-2h2zm10 0V8h-2v2h2zm0 0v2h2v-2h-2z"/>`
	arrowUpBoxInnerSVG                 = `<path fill="currentColor" d="M3 21h18V3H3v18zM19 5v14H5V5h14zm-8 12h2v-6h2V9h-2V7h-2v2H9v2h2v6zm-2-4v-2H7v2h2zm8 0h-2v-2h2v2z"/>`
	arrowsHorizontalInnerSVG           = `<path fill="currentColor" d="M15 9V7h2v2h-2zm2 6v-2h-4v-2h4V9h2v2h2v2h-2v2h-2zm0 0v2h-2v-2h2zm-6-4v2H7v2H5v-2H3v-2h2V9h2v2h4zm-4 4h2v2H7v-2zm2-8v2H7V7h2z"/>`
	arrowsVerticalInnerSVG             = `<path fill="currentColor" d="M11 11h2V7h2v2h2V7h-2V5h-2V3h-2v2H9v2H7v2h2V7h2v4zm0 2h2v4h2v2h-2v2h-2v-2H9v-2h2v-4zm-2 4v-2H7v2h2zm6 0v-2h2v2h-2z"/>`
	artTextInnerSVG                    = `<path fill="currentColor" d="M2 7h10v10H2V7zm8 8V9H4v6h6zm12-8h-8v2h8V7zm-8 4h8v2h-8v-2zm8 4h-8v2h8v-2z"/>`
	articleInnerSVG                    = `<path fill="currentColor" d="M5 3H3v18h18V3H5zm14 2v14H5V5h14zm-2 2H7v2h10V7zM7 11h10v2H7v-2zm7 4H7v2h7v-2z"/>`
	articleMultipleInnerSVG            = `<path fill="currentColor" d="M3 1H1v18h18V1H3zm14 2v14H3V3h14zm4 18H5v2h18V5h-2v16zM15 5H5v2h10V5zM5 9h10v2H5V9zm7 4H5v2h7v-2z"/>`
	aspectRatioInnerSVG                = `<path fill="currentColor" d="M2 4h20v16H2V4zm2 14h16V6H4v12zM8 8h2v2H8v2H6V8h2zm8 8h-2v-2h2v-2h2v4h-2z"/>`
	atInnerSVG                         = `<path fill="currentColor" d="M4 4h16v12H8V8h8v6h2V6H6v12h14v2H4V4zm10 10v-4h-4v4h4z"/>`
	attachmentInnerSVG                 = `<path fill="currentColor" d="M7 5v14H5V3h14v18H9V7h6v10h-2V9h-2v10h6V5H7z"/>`
	audioDeviceInnerSVG                = `<path fill="currentColor" d="M4 4h4v2H4v8h4v2H2V4h2zm6 0h10v2h-8v12h8v2H10V4zm12 0h-2v16h2V4zm-7 4h2v2h-2V8zm3 4h-4v4h4v-4zM8 18H4v2h4v-2z"/>`
	avatarInnerSVG                     = `<path fill="currentColor" d="M3 3h18v18H3V3zm16 16V5H5v14h14zM14 7h-4v4h4V7zm1 6H9v2H7v2h2v-2h6v2h2v-2h-2v-2z"/>`
	backburgerInnerSVG                 = `<path fill="currentColor" d="M11 7h10v2H11V7zm-8 4h2V9h2v2h14v2H7v2H5v-2H3v-2zm4 4v2h2v-2H7zm0-6V7h2v2H7zm14 6H11v2h10v-2z"/>`
	batteryInnerSVG                    = `<path fill="currentColor" d="M4 5H2v14h18v-4h2V9h-2V5H4zm14 2v10H4V7h14z"/>`
	batteryChargingInnerSVG            = `<path fill="currentColor" d="M4 5H2v14h6v-2H4V7h4V5H4zm10 0h6v4h2v6h-2v4h-6v-2h4V7h-4V5zm-4 2h2v4h4v2h-2v2h-2v2h-2v-4H6v-2h2V9h2V7z"/>`
	batteryFullInnerSVG                = `<path fill="currentColor" d="M18 5H2v14h18v-4h2V9h-2V5h-2zm0 2v10H4V7h14zM8 9H6v6h2V9zm2 0h2v6h-2V9zm6 0h-2v6h2V9z"/>`
	batteryOneInnerSVG                 = `<path fill="currentColor" d="M4 5H2v14h18v-4h2V9h-2V5H4zm14 2v10H4V7h14zM8 9H6v6h2V9z"/>`
	batteryTwoInnerSVG                 = `<path fill="currentColor" d="M4 5H2v14h18v-4h2V9h-2V5H4zm14 2v10H4V7h14zM6 9h2v6H6V9zm6 0h-2v6h2V9z"/>`
	bedInnerSVG                        = `<path fill="currentColor" d="M0 4h2v12h10V8h10v2h-8v6h8v-6h2v10h-2v-2H2v2H0V4zm3 5h2v4H3V9zm6 4v2H5v-2h4zm0-4h2v4H9V9zm0 0H5V7h4v2z"/>`
	bitcoinInnerSVG                    = `<path fill="currentColor" d="M13 3h2v2h2v2H9v4h8v2H9v4h8v2h-2v2h-2v-2h-2v2H9v-2H5v-2h2v-4H5v-2h2V7H5V5h4V3h2v2h2V3zm4 14v-4h2v4h-2zm0-6V7h2v4h-2z"/>`
	bluetoothInnerSVG                  = `<path fill="currentColor" d="M15 3h-2v2h2v2h2v2h-2v2h2V9h2V7h-2V5h-2V3zm-2 0h-2v6H9V7H7V5H5v2h2v2h2v2h2v2H9v2H7v2H5v2h2v-2h2v-2h2v6h2V3zm2 8h-2v2h2v2h2v2h-2v2h-2v2h2v-2h2v-2h2v-2h-2v-2h-2v-2z"/>`
	bookInnerSVG                       = `<path fill="currentColor" d="M8 2h12v20H4V2h4zm4 8h-2v2H8V4H6v16h12V4h-4v8h-2v-2z"/>`
	bookOpenInnerSVG                   = `<path fill="currentColor" d="M3 3h8v2H3v12h8V5h2v12h8V5h-8V3h10v16H13v2h-2v-2H1V3h2zm16 7h-4v2h4v-2zm-4-3h4v2h-4V7zm2 6h-2v2h2v-2z"/>`
	bookmarkInnerSVG                   = `<path fill="currentColor" d="M18 2H6v2h12v16h-2v-2h-2v-2h-4v2H8v2H6V2H4v20h4v-2h2v-2h4v2h2v2h4V2h-2z"/>`
	bookmarksInnerSVG                  = `<path fill="currentColor" d="M21 18V2H7v2h12v14h2zM5 6H3v16h4v-2h2v-2h2v2h2v2h4V6H5zm8 14v-2h-2v-2H9v2H7v2H5V8h10v12h-2z"/>`
	briefcaseInnerSVG                  = `<path fill="currentColor" d="M8 3h8v4h6v14H2V7h6V3zm2 4h4V5h-4v2zM4 9v10h16V9H4z"/>`
	briefcaseAccountInnerSVG           = `<path fill="currentColor" d="M16 3H8v4H2v14h20V7h-6V3zm-2 4h-4V5h4v2zM4 19V9h16v10H4zm6-8h4v3h-4v-3zm-2 4h8v2H8v-2z"/>`
	briefcaseCheckInnerSVG             = `<path fill="currentColor" d="M16 3H8v4H2v14h20V7h-6V3zm-2 4h-4V5h4v2zM4 19V9h16v10H4zm10-8h2v2h-2v-2zm-2 4v-2h2v2h-2zm-2 0h2v2h-2v-2zm0 0H8v-2h2v2z"/>`
	briefcaseDeleteInnerSVG            = `<path fill="currentColor" d="M16 3H8v4H2v14h12v-2H4V9h16v4h2V7h-6V3zm-2 4h-4V5h4v2zm4 8h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2h2v-2h-2v2h-2v-2z"/>`
	briefcaseDownloadInnerSVG          = `<path fill="currentColor" d="M8 3h8v4h6v14h-5v-2h3V9H4v10h3v2H2V7h6V3zm6 2h-4v2h4V5zm-3 6h2v6h2v2h-2v2h-2v-2H9v-2h2v-6zm-2 6H7v-2h2v2zm6 0v-2h2v2h-2z"/>`
	briefcaseMinusInnerSVG             = `<path fill="currentColor" d="M8 3h8v4h6v6h-2V9H4v10h10v2H2V7h6V3zm6 2h-4v2h4V5zm2 12h6v2h-6v-2z"/>`
	briefcasePlusInnerSVG              = `<path fill="currentColor" d="M8 3h8v4h6v4h-2V9H4v10h8v2H2V7h6V3zm2 4h4V5h-4v2zm7 14h2v-3h3v-2h-3v-3h-2v3h-3v2h3v3z"/>`
	briefcaseSearchInnerSVG            = `<path fill="currentColor" d="M16 3H8v4H2v14h10v-2H4V9h16v2h2V7h-6V3zm-2 4h-4V5h4v2zm6 6h-6v6h6v2h2v-2h-2v-6zm-4 4v-2h2v2h-2z"/>`
	briefcaseSearchOneInnerSVG         = `<path fill="currentColor" d="M16 3H8v4H2v14h7v-2H4V9h18V7h-6V3zm-2 4h-4V5h4v2zm0 4h8v2h-8v-2zm0 10h-2v-8h2v8zm8 0v2h-8v-2h8zm0 0h2v-8h-2v8zm-6-6h2v2h2v2h-4v-4z"/>`
	briefcaseUploadInnerSVG            = `<path fill="currentColor" d="M8 3h8v4h6v14h-5v-2h3V9H4v10h3v2H2V7h6V3zm6 2h-4v2h4V5zm-3 16h2v-6h2v2h2v-2h-2v-2h-2v-2h-2v2H9v2H7v2h2v-2h2v6z"/>`
	bugInnerSVG                        = `<path fill="currentColor" d="M8 2h2v4h4V2h2v4h2v3h2v2h-2v2h4v2h-4v2h2v2h-2v3H6v-3H4v-2h2v-2H2v-2h4v-2H4V9h2V6h2V2Zm8 6H8v3h8V8Zm-5 5H8v7h3v-7Zm2 7h3v-7h-3v7ZM4 9H2V7h2v2Zm0 10v2H2v-2h2Zm16 0h2v2h-2v-2Zm0-10V7h2v2h-2Z"/>`
	buildingInnerSVG                   = `<path fill="currentColor" d="M3 2h18v20H3V2zm12 16v2h4V4H5v16h4v-2h6zM7 6h2v2H7V6zm6 0h-2v2h2V6zm2 0h2v2h-2V6zm-6 4H7v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2zM7 14h2v2H7v-2zm6 0h-2v2h2v-2zm4 0h-2v2h2v-2z"/>`
	buildingCommunityInnerSVG          = `<path fill="currentColor" d="M20 2h2v20H2v-8h2v6h4v-4h2v4h4v-6h2v6h4V4H10v2H8V2h12zm-8 10h2v2h-2v-2zm-2-2h2v2h-2v-2zm-2 0V8h2v2H8zm-2 2v-2h2v2H6zm0 0H4v2h2v-2zm10-6h2v2h-2V6zm-2 0h-2v2h2V6zm2 4h2v2h-2v-2z"/>`
	buildingSkyscraperInnerSVG         = `<path fill="currentColor" d="M10 2h4v5h2v2h-2v11h4v-9h2v9h2v2H2v-2h2V8h2v12h6V4h-2V2zM8 6V4h2v2H8zm0 0H6v2h2V6zm10 5h-2V9h2v2zm-8-1H8v2h2v-2zm-2 4h2v2H8v-2z"/>`
	buildingsInnerSVG                  = `<path fill="currentColor" d="M2 2h14v4h6v16H2V2zm18 6h-4v2h2v2h-2v2h2v2h-2v2h2v2h2V8zm-6-4H4v16h2v-2h6v2h2V4zM6 6h2v2H6V6zm6 0h-2v2h2V6zm-6 4h2v2H6v-2zm6 0h-2v2h2v-2zm-6 4h2v2H6v-2zm6 0h-2v2h2v-2z"/>`
	bulletlistInnerSVG                 = `<path fill="currentColor" d="M2 11V5h6v6H2zm4-2V7H4v2h2zm16-4H10v2h12V5zm0 4H10v2h12V9zm-12 4h12v2H10v-2zm12 4H10v2h12v-2zM2 13v6h6v-6H2zm4 2v2H4v-2h2z"/>`
	bullseyeInnerSVG                   = `<path fill="currentColor" d="M18 2H6v2H4v2H2v12h2v2h2v2h12v-2h2v-2h2V6h-2V4h-2V2zm0 2v2h2v12h-2v2H6v-2H4V6h2V4h12zm-8 6h4v4h-4v-4zM8 6h8v2H8V6zm0 10H6V8h2v8zm8 0v2H8v-2h8zm0 0h2V8h-2v8z"/>`
	bullseyeArrowInnerSVG              = `<path fill="currentColor" d="M6 2h10v2H6V2zM4 6V4h2v2H4zm0 12H2V6h2v12zm2 2H4v-2h2v2zm12 0H6v2h12v-2zm2-2v2h-2v-2h2zm0 0h2V8h-2v10zM12 6H8v2H6v8h2v2h8v-2h2v-4h-2v4H8V8h4V6zm2 8v-4h2V8h2V6h4V4h-2V2h-2v4h-2v2h-2v2h-4v4h4z"/>`
	busInnerSVG                        = `<path fill="currentColor" d="M5 2h14v2H5V2zm0 2v6h14V4h2v16h-2v2h-4v-2H9v2H5v-2H3V4h2zm0 14h14v-6H5v6zm2-4h2v2H7v-2zm10 0h-2v2h2v-2z"/>`
	cakeInnerSVG                       = `<path fill="currentColor" d="M6 2h2v2H6V2zm2 3H6v3H2v9h6v-2h2v2h4v-2h2v2h6V8h-4V5h-2v3h-3V5h-2v3H8V5zm12 10h-4v-3h-2v3h-4v-3H8v3H4v-5h16v5zM2 20h20v2H2v-2zM13 2h-2v2h2V2zm3 0h2v2h-2V2zM2 17h2v3H2zm18 0h2v3h-2z"/>`
	calculatorInnerSVG                 = `<path fill="currentColor" d="M5 2H3v20h18V2H5zm14 18H5V4h14v16zM17 6H7v4h10V6zM7 12h2v2H7v-2zm6 0h-2v2h2v-2zm2 0h2v2h-2v-2zm-6 4H7v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2z"/>`
	calendarInnerSVG                   = `<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6V2zM5 8h14V6H5v2zm0 2v10h14V10H5z"/>`
	calendarAlertInnerSVG              = `<path fill="currentColor" d="M7 5V4H5v2H3v14h14V6h-2V4h-2v2H7V5zm-2 5V8h10v2H5zm0 2h10v6H5v-6zm16-3V8h-2v6h2V9zm0 6h-2v2h2v-2z"/>`
	calendarArrowLeftInnerSVG          = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v8h2v-2h14v10h-8v2h10V4h-4V2zm2 6H5V6h14v2zm-6 8H7v-2h2v-2H7v2H5v2H3v2h2v2h2v2h2v-2H7v-2h6v-2z"/>`
	calendarArrowRightInnerSVG         = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h10v-2H5V10h14v2h2V4h-4V2zM7 6h12v2H5V6h2zm14 10h-2v-2h-2v-2h-2v2h2v2h-6v2h6v2h-2v2h2v-2h2v-2h2v-2z"/>`
	calendarCheckInnerSVG              = `<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6V2zm4 6V6H5v2h14zm0 2H5v10h14V10zm-3 2v2h-2v-2h2zm-4 4v-2h2v2h-2zm-2 0h2v2h-2v-2zm0 0H8v-2h2v2z"/>`
	calendarExportInnerSVG             = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h4v-2H5V10h14v10h-2v2h4V4h-4V2zM7 6h12v2H5V6h2zm6 6h-2v6H9v-2H7v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2v-6z"/>`
	calendarGridInnerSVG               = `<path fill="currentColor" d="M3 3h18v18H3V3zm2 2v2h14V5H5zm14 4h-6v2h6V9zm0 4h-6v2h6v-2zm0 4h-6v2h6v-2zm-8 2v-2H5v2h6zm-6-4h6v-2H5v2zm0-4h6V9H5v2z"/>`
	calendarImportInnerSVG             = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h4v-2H5V10h14v10h-2v2h4V4h-4V2zM7 6h12v2H5V6h2zm6 16h-2v-6H9v-2h2v-2h2v2h2v2h-2v6zm2-6v2h2v-2h-2zm-6 0v2H7v-2h2z"/>`
	calendarMinusInnerSVG              = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm10-6H9v2h6v-2z"/>`
	calendarMonthInnerSVG              = `<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6V2zM9 6H5v2h14V6H9zm-4 4v10h14V10H5zm2 2h2v2H7v-2zm6 0h-2v2h2v-2zm2 0h2v2h-2v-2zm-6 4H7v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2z"/>`
	calendarMultipleInnerSVG           = `<path fill="currentColor" d="M17 2h2v2h4v14H5V4h4V2h2v2h6V2zm-6 4H7v2h14V6H11zm-4 4v6h14v-6H7zM3 20h16v2H1V8h2v12z"/>`
	calendarMultipleCheckInnerSVG      = `<path fill="currentColor" d="M17 2h2v2h4v10h-2v-4H7v6h6v2H5V4h4V2h2v2h6V2zm-6 4H7v2h14V6H11zm2 14v2H1V8h2v12h10zm2-2h2v2h-2v-2zm4 2v2h-2v-2h2zm2-2h-2v2h2v-2zm0 0v-2h2v2h-2z"/>`
	calendarPlusInnerSVG               = `<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6V2zM9 6H5v2h14V6H9zm-4 4v10h14V10H5zm6 2h2v2h2v2h-2v2h-2v-2H9v-2h2v-2z"/>`
	calendarRangeInnerSVG              = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm4-8H7v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2z"/>`
	calendarRemoveInnerSVG             = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm6-4H9v2h2v-2zm0-2v-2H9v2h2zm2 0h-2v2h2v2h2v-2h-2v-2zm0 0v-2h2v2h-2z"/>`
	calendarSearchInnerSVG             = `<path fill="currentColor" d="M15 2h2v2h4v8h-2v-2H5v10h6v2H3V4h4V2h2v2h6V2zM9 6H5v2h14V6H9zm8 6v2h-4v-2h4zm-4 6h-2v-4h2v4zm4 0h-4v2h6v2h2v-2h-2v-6h-2v4z"/>`
	calendarSortAscendingInnerSVG      = `<path fill="currentColor" d="M10 5H8v2H4V5H2v2H0v12h12V7h-2V5zM2 9h8v2H2V9zm0 8v-4h8v4H2zM20 7h-2v8h-2v-2h-2v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2V7z"/>`
	calendarSortDescendingInnerSVG     = `<path fill="currentColor" d="M10 5H8v2H4V5H2v2H0v12h12V7h-2V5zM2 9h8v2H2V9zm0 8v-4h8v4H2zm18 2h-2v-8h-2V9h2V7h2v2h2v2h-2v8zm2-8v2h2v-2h-2zm-6 0v2h-2v-2h2z"/>`
	calendarTextInnerSVG               = `<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6V2zM9 6H5v2h14V6H9zm-4 4v10h14V10H5zm2 2h8v2H7v-2zm4 6v-2H7v2h4z"/>`
	calendarTodayInnerSVG              = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm6-4v-4H7v4h4z"/>`
	calendarTomorrowInnerSVG           = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm12-2v-4h-4v4h4z"/>`
	calendarWeekInnerSVG               = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm12-8H7v2h10v-2z"/>`
	calendarWeekBeginInnerSVG          = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm4-8H7v6h2v-6z"/>`
	calendarWeekendInnerSVG            = `<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm12-8h-2v6h2v-6z"/>`
	cameraInnerSVG                     = `<path fill="currentColor" d="M9 3H7v2H2v16h20V5h-5V3H9zm8 4h3v12H4V7h5V5h6v2h2zm-7 2h4v2h-4V9zm4 6h-4v2h4v-2h2v-4h-2v4zm-6-4h2v4H8v-4z"/>`
	cameraAddInnerSVG                  = `<path fill="currentColor" d="M5 2H3v3H0v2h3v3h2V7h3V5H5V2zm12 1h-7v2h5v2h5v12H5v-7H3v9h19V5h-5V3zm-7 6h4v2h2v4h-2v2h-4v-2h4v-4h-4V9zm-2 2h2v4H8v-4z"/>`
	cameraAltInnerSVG                  = `<path fill="currentColor" d="M4 4H2v16h20V4H4zm16 2v12H4V6h16zM8 8H6v2h2V8zm4 0h4v2h-4V8zm-2 2h2v4h-2v-4zm6 4h2v-4h-2v4zm0 0h-4v2h4v-2z"/>`
	cameraFaceInnerSVG                 = `<path fill="currentColor" d="M7 3h10v2h5v16H2V7h2v12h16V7h-5V5H9v2H2V5h5V3zm7 12h-4v2h4v-2zm-4-2v2H8v-2h2zm0-2V9H8v2h2zm6 2v2h-2v-2h2zm0-2V9h-2v2h2z"/>`
	carInnerSVG                        = `<path fill="currentColor" d="M17 4H7v2H5v2H3v12h4v-2h10v2h4V8h-2V6h-2V4zm0 2v2h2v2H5V8h2V6h10zm2 10H5v-4h14v4zm-2-3h-2v2h2v-2zM7 13h2v2H7v-2z"/>`
	cardInnerSVG                       = `<path fill="currentColor" d="M2 4h20v16H2V4zm18 14V6H4v12h16z"/>`
	cardIdInnerSVG                     = `<path fill="currentColor" d="M2 4h20v16H2V4zm2 2v4h16V6H4zm16 6H10v2h10v-2zm0 4h-4v2h4v-2zm-6 2v-2H4v2h10zM4 14h4v-2H4v2z"/>`
	cardPlusInnerSVG                   = `<path fill="currentColor" d="M22 4H2v16h10v-2H4V6h16v4h2V4zm-3 13h3v-2h-3v-3h-2v3h-3v2h3v3h2v-3z"/>`
	cardStackInnerSVG                  = `<path fill="currentColor" d="M4 4h18v12H2V4h2zm16 10V6H4v8h16zm2 4H2v2h20v-2z"/>`
	cardTextInnerSVG                   = `<path fill="currentColor" d="M4 4H2v16h20V4H4zm0 2h16v12H4V6zm2 2h12v2H6V8zm0 4h10v2H6v-2z"/>`
	cartInnerSVG                       = `<path fill="currentColor" d="M2 2h4v4h16v11H4V4H2V2zm4 13h14V8H6v7zm0 4h3v3H6v-3zm14 0h-3v3h3v-3z"/>`
	castInnerSVG                       = `<path fill="currentColor" d="M4 3h18v18h-8v-2h6V5H4v4H2V3h2zm0 16H2v2h2v-2zm-2-4h4v2H2v-2zm8-4H2v2h8v8h2V11h-2zm-4 4h2v6H6v-6z"/>`
	cellularSignalOffInnerSVG          = `<path fill="currentColor" d="M4 2H2v2h2v2H2v2h2V6h2v2h2V6H6V4h2V2H6v2H4V2Zm12 2v16h6V4h-6Zm2 2h2v12h-2V6Zm-9 4v10h6V10H9Zm2 8v-6h2v6h-2Zm-3-4v6H2v-6h6Zm-2 4v-2H4v2h2Z"/>`
	cellularSignalOneInnerSVG          = `<path fill="currentColor" d="M16 4v16h6V4h-6Zm2 2h2v12h-2V6Zm-9 4v10h6V10H9Zm2 8v-6h2v6h-2Zm-3-4H2v6h6v-6Z"/>`
	cellularSignalThreeInnerSVG        = `<path fill="currentColor" d="M16 4h6v16h-6V4ZM2 14h6v6H2v-6Zm13-4H9v10h6V10Z"/>`
	cellularSignalTwoInnerSVG          = `<path fill="currentColor" d="M16 4v16h6V4h-6Zm4 2v12h-2V6h2ZM2 14h6v6H2v-6Zm13-4H9v10h6V10Z"/>`
	cellularSignalZeroInnerSVG         = `<path fill="currentColor" d="M22 4v16h-6V4h6Zm-2 2h-2v12h2V6Zm-5 4v10H9V10h6Zm-2 8v-6h-2v6h2Zm-5-4v6H2v-6h6Zm-2 4v-2H4v2h2Z"/>`
	chartInnerSVG                      = `<path fill="currentColor" d="M5 3H3v18h18V3H5zm14 2v14H5V5h14zM9 11H7v6h2v-6zm2-4h2v10h-2V7zm6 6h-2v4h2v-4z"/>`
	chartAddInnerSVG                   = `<path fill="currentColor" d="M3 3h10v2H5v14h14v-8h2v10H3V3zm6 8H7v6h2v-6zm2-4h2v10h-2V7zm6 6h-2v4h2v-4zm0-10h2v2h2v2h-2v2h-2V7h-2V5h2V3z"/>`
	chartBarInnerSVG                   = `<path fill="currentColor" d="M13 5h2v14h-2V5zm-2 4H9v10h2V9zm-4 4H5v6h2v-6zm12 0h-2v6h2v-6z"/>`
	chartDeleteInnerSVG                = `<path fill="currentColor" d="M13 3H3v18h18V11h-2v8H5V5h8V3zm-6 8h2v6H7v-6zm6-4h-2v10h2V7zm2 6h2v4h-2v-4zm2-6h-2v2h2V7zm0-2V3h-2v2h2zm2 0h-2v2h2v2h2V7h-2V5zm0 0V3h2v2h-2z"/>`
	chartMinusInnerSVG                 = `<path fill="currentColor" d="M13 3H3v18h18V11h-2v8H5V5h8V3zm-6 8h2v6H7v-6zm6-4h-2v10h2V7zm2 6h2v4h-2v-4zm6-8h-6v2h6V5z"/>`
	chartMultipleInnerSVG              = `<path fill="currentColor" d="M3 2H1v16h18V2H3zm0 2h14v12H3V4zm18 2v14H5v2h18V6h-2zM7 8H5v6h2V8zm2-2h2v8H9V6zm6 4h-2v4h2v-4z"/>`
	chatInnerSVG                       = `<path fill="currentColor" d="M20 2H2v20h2V4h16v12H6v2H4v2h2v-2h16V2h-2z"/>`
	checkInnerSVG                      = `<path fill="currentColor" d="M18 6h2v2h-2V6zm-2 4V8h2v2h-2zm-2 2v-2h2v2h-2zm-2 2h2v-2h-2v2zm-2 2h2v-2h-2v2zm-2 0v2h2v-2H8zm-2-2h2v2H6v-2zm0 0H4v-2h2v2z"/>`
	checkDoubleInnerSVG                = `<path fill="currentColor" d="M15 6h2v2h-2V6zm-2 4V8h2v2h-2zm-2 2v-2h2v2h-2zm-2 2v-2h2v2H9zm-2 2v-2h2v2H7zm-2 0h2v2H5v-2zm-2-2h2v2H3v-2zm0 0H1v-2h2v2zm8 2h2v2h-2v-2zm4-2v2h-2v-2h2zm2-2v2h-2v-2h2zm2-2v2h-2v-2h2zm2-2h-2v2h2V8zm0 0h2V6h-2v2z"/>`
	checkboxInnerSVG                   = `<path fill="currentColor" d="M5 3H3v18h18V3H5zm0 2h14v14H5V5zm4 7H7v2h2v2h2v-2h2v-2h2v-2h2V8h-2v2h-2v2h-2v2H9v-2z"/>`
	checkboxOnInnerSVG                 = `<path fill="currentColor" d="M3 3h18v18H3V3zm16 16V5H5v14h14z"/>`
	checklistInnerSVG                  = `<path fill="currentColor" d="M19 4h2v2h-2V4zm-2 4V6h2v2h-2zm-2 0h2v2h-2V8zm0 0h-2V6h2v2zM3 6h8v2H3V6zm8 10H3v2h8v-2zm7 2v-2h2v-2h-2v2h-2v-2h-2v2h2v2h-2v2h2v-2h2zm0 0v2h2v-2h-2z"/>`
	chessInnerSVG                      = `<path fill="currentColor" d="M2 2h20v20H2V2zm2 2v4h4v4H4v4h4v4h4v-4h4v4h4v-4h-4v-4h4V8h-4V4h-4v4H8V4H4zm8 8H8v4h4v-4zm0-4v4h4V8h-4z"/>`
	chevronDownInnerSVG                = `<path fill="currentColor" d="M7 8H5v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2V8h-2v2h-2v2h-2v2h-2v-2H9v-2H7V8z"/>`
	chevronLeftInnerSVG                = `<path fill="currentColor" d="M16 5v2h-2V5h2zm-4 4V7h2v2h-2zm-2 2V9h2v2h-2zm0 2H8v-2h2v2zm2 2v-2h-2v2h2zm0 0h2v2h-2v-2zm4 4v-2h-2v2h2z"/>`
	chevronRightInnerSVG               = `<path fill="currentColor" d="M8 5v2h2V5H8zm4 4V7h-2v2h2zm2 2V9h-2v2h2zm0 2h2v-2h-2v2zm-2 2v-2h2v2h-2zm0 0h-2v2h2v-2zm-4 4v-2h2v2H8z"/>`
	chevronUpInnerSVG                  = `<path fill="currentColor" d="M7 16H5v-2h2v-2h2v-2h2V8h2v2h2v2h2v2h2v2h-2v-2h-2v-2h-2v-2h-2v2H9v2H7v2z"/>`
	chevronsHorizontalInnerSVG         = `<path fill="currentColor" d="M8 9V7h2v2H8zm-2 2V9h2v2H6zm0 2H4v-2h2v2zm2 2v-2H6v2h2zm0 0h2v2H8v-2zm8-6V7h-2v2h2zm2 2V9h-2v2h2zm0 2v-2h2v2h-2zm-2 2v-2h2v2h-2zm0 0v2h-2v-2h2z"/>`
	chevronsVerticalInnerSVG           = `<path fill="currentColor" d="M11 4h2v2h-2V4zM9 8V6h2v2H9zm0 0v2H7V8h2zm6 0h-2V6h2v2zm0 0h2v2h-2V8zm-6 8H7v-2h2v2zm2 2H9v-2h2v2zm2 0v2h-2v-2h2zm2-2h-2v2h2v-2zm0 0v-2h2v2h-2z"/>`
	circleInnerSVG                     = `<path fill="currentColor" d="M17 3H7v2H5v2H3v10h2v2h2v2h10v-2h2v-2h2V7h-2V5h-2V3zm0 2v2h2v10h-2v2H7v-2H5V7h2V5h10z"/>`
	clipboardInnerSVG                  = `<path fill="currentColor" d="M10 2h6v2h4v18H4V4h4V2h2zm6 4v2H8V6H6v14h12V6h-2zm-2 0V4h-4v2h4z"/>`
	clockInnerSVG                      = `<path fill="currentColor" d="M19 3H5v2H3v14h2v2h14v-2h2V5h-2V3zm0 2v14H5V5h14zm-8 2h2v6h4v2h-6V7z"/>`
	closeInnerSVG                      = `<path fill="currentColor" d="M5 5h2v2H5V5zm4 4H7V7h2v2zm2 2H9V9h2v2zm2 0h-2v2H9v2H7v2H5v2h2v-2h2v-2h2v-2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2zm2-2v2h-2V9h2zm2-2v2h-2V7h2zm0 0V5h2v2h-2z"/>`
	closeBoxInnerSVG                   = `<path fill="currentColor" d="M5 3H3v18h18V3H5zm14 2v14H5V5h14zm-8 4H9V7H7v2h2v2h2v2H9v2H7v2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2v-2h2V9h2V7h-2v2h-2v2h-2V9z"/>`
	cloudInnerSVG                      = `<path fill="currentColor" d="M16 4h-6v2H8v2H4v2H2v2H0v6h2v2h20v-2h2v-6h-2v-2h-2V8h-2V6h-2V4zm2 8h4v6H2v-6h2v-2h4v2h2v-2H8V8h2V6h6v2h2v4zm0 0v2h-2v-2h2z"/>`
	cloudDoneInnerSVG                  = `<path fill="currentColor" d="M16 4h-6v2H8v2H4v2H2v2H0v6h2v2h20v-2h2v-6h-2v-2h-2V8h-2V6h-2V4zm0 2v2h2v4h4v6H2v-6h2v-2h4V8h2V6h6zm-6 6H8v2h2v2h2v-2h2v-2h2v-2h-2v2h-2v2h-2v-2z"/>`
	cloudDownloadInnerSVG              = `<path fill="currentColor" d="M10 4h6v2h-6V4zM8 8V6h2v2H8zm-4 2V8h4v2H4zm-2 2v-2h2v2H2zm0 6H0v-6h2v6zm0 0h5v2H2v-2zM18 8h-2V6h2v2zm4 4h-4V8h2v2h2v2zm0 6v-6h2v6h-2zm0 0v2h-5v-2h5zm-11 2h2v-2h2v-2h2v-2h-4V9h-2v5H7v2h2v2h2v2z"/>`
	cloudMoonInnerSVG                  = `<path fill="currentColor" d="M18 2h-8v2H8v2H6v4h2V6h2V4h4v2h-2v4h2v2h4v-2h2v4h-2v2h2v-2h2V6h-2v2h-2v2h-4V6h2V4h2V2ZM8 14v-2h4v2H8Zm0 2v-2H4v2H2v4h2v2h10v-2h2v-4h-2v-2h-2v2h2v4H4v-4h4Zm0 0h2v2H8v-2Z"/>`
	cloudSunInnerSVG                   = `<path fill="currentColor" d="M11 0h2v4h-2V0Zm1 12H8v2H4v2H2v4h2v2h10v-2h2v-4h-2v-2h-2v-2Zm0 2v2h2v4H4v-4h4v2h2v-2H8v-2h4ZM8 6h6v2H8V6Zm0 2v2H6V8h2Zm8 2h-2V8h2v2Zm0 0h2v2h-2v-2Zm4-8h2v2h-2V2Zm0 2v2h-2V4h2ZM2 2h2v2H2V2Zm2 2h2v2H4V4Zm20 7h-4v2h4v-2Z"/>`
	cloudUploadInnerSVG                = `<path fill="currentColor" d="M10 4h6v2h-6V4zM8 8V6h2v2H8zm-4 2V8h4v2H4zm-2 2v-2h2v2H2zm0 6H0v-6h2v6zm0 0h7v2H2v-2zM18 8h-2V6h2v2zm4 4h-4V8h2v2h2v2zm0 6v-6h2v6h-2zm0 0v2h-7v-2h7zM11 9h2v2h2v2h2v2h-4v5h-2v-5H7v-2h2v-2h2V9z"/>`
	cocktailInnerSVG                   = `<path fill="currentColor" d="M19 3H3v4h2v2h2v2h2v2h2v6H7v2h10v-2h-4v-6h2v-2h2V9h2V7h2V3h-2zm0 4H5V5h14v2z"/>`
	codeInnerSVG                       = `<path fill="currentColor" d="M8 5h2v2H8V5zM6 7h2v2H6V7zM4 9h2v2H4V9zm-2 2h2v2H2v-2zm2 2h2v2H4v-2zm2 2h2v2H6v-2zm2 2h2v2H8v-2zm8-12h-2v2h2V5zm2 2h-2v2h2V7zm2 2h-2v2h2V9zm2 2h-2v2h2v-2zm-2 2h-2v2h2v-2zm-2 2h-2v2h2v-2zm-2 2h-2v2h2v-2z"/>`
	coffeeInnerSVG                     = `<path fill="currentColor" d="M4 4h18v7h-4v5H4V4zm14 5h2V6h-2v3zm-2-3H6v8h10V6zm3 14H3v-2h16v2z"/>`
	coffeeAltInnerSVG                  = `<path fill="currentColor" d="M7 3H5v4h2V3zm4 0H9v4h2V3zm2 0h2v4h-2V3zm8 6H3v12h14v-5h4V9zm-2 5h-2v-3h2v3zM5 11h10v8H5v-8z"/>`
	coinInnerSVG                       = `<path fill="currentColor" d="M6 2h12v2H6V2zM4 6V4h2v2H4zm0 12V6H2v12h2zm2 2v-2H4v2h2zm12 0v2H6v-2h12zm2-2v2h-2v-2h2zm0-12h2v12h-2V6zm0 0V4h-2v2h2zm-9-1h2v2h3v2h-6v2h6v6h-3v2h-2v-2H8v-2h6v-2H8V7h3V5z"/>`
	collapseInnerSVG                   = `<path fill="currentColor" d="M17 3h-2v2h-2v2h-2V5H9V3H7v2h2v2h2v2h2V7h2V5h2V3zM4 13h16v-2H4v2zm9 4h-2v-2h2v2zm2 2h-2v-2h2v2zm0 0h2v2h-2v-2zm-6 0h2v-2H9v2zm0 0H7v2h2v-2z"/>`
	colorsSwatchInnerSVG               = `<path fill="currentColor" d="M14 2h8v20H12V2h2zm6 2h-6v16h6V4zM10 20H4v-6h6v-2H6v-2H4V8h2V6h2V4h2V2H8v2H6v2H4v2H2v2h2v2H2v10h8v-2zm8-4h-2v2h2v-2z"/>`
	commandInnerSVG                    = `<path fill="currentColor" d="M4 2H2v8h2V2zm16 0h2v8h-2V2zm-6 6h-4V2H4v2h4v4H4v2h4v4H4v2h4v4H4v2h6v-6h4v6h2v-6h4v-2h-4v-4h4V8h-4V2h-2v6zm-4 6v-4h4v4h-4zM20 2h-4v2h4V2zM2 14h2v8H2v-8zm14 6h4v2h-4v-2zm6-6h-2v8h2v-8z"/>`
	commentInnerSVG                    = `<path fill="currentColor" d="M22 2H2v14h2V4h16v12h-8v2h-2v2H8v-4H2v2h4v4h4v-2h2v-2h10V2z"/>`
	contactInnerSVG                    = `<path fill="currentColor" d="M2 3H0v18h24V3H2zm20 2v14H2V5h20zM10 7H6v4h4V7zm-6 6h8v4H4v-4zm16-6h-6v2h6V7zm-6 4h6v2h-6v-2zm6 4h-6v2h6v-2z"/>`
	contactDeleteInnerSVG              = `<path fill="currentColor" d="M22 3H0v18h16v-2H2V5h20v10h2V3h-2zM6 7h4v4H6V7zm0 8H4v2h2v-2zm4 0H6v-2h4v2zm0 0v2h2v-2h-2zm4-8h6v2h-6V7zm6 4h-6v2h6v-2zm-6 4h2v2h-2v-2zm8 4h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2zm0 0h2v-2h-2v2z"/>`
	contactMultipleInnerSVG            = `<path fill="currentColor" d="M4 3h20v14H4V3zm18 12V5H6v10h16zm-2 4H2V7H0v14h20v-2zM9 7h2v2H9V7zm3 4H8v2h4v-2zm2-4h6v2h-6V7zm6 4h-6v2h6v-2z"/>`
	contactPlusInnerSVG                = `<path fill="currentColor" d="M2 3h22v11h-2V5H2v14h12v2H0V3h2zm8 4H6v4h4V7zm-6 6h8v4H4v-4zm16-6h-6v2h6V7zm-6 4h6v2h-6v-2zm3 4h-3v2h3v-2zm4 6v3h-2v-3h-3v-2h3v-3h2v3h3v2h-3z"/>`
	copyInnerSVG                       = `<path fill="currentColor" d="M4 2h11v2H6v13H4V2zm4 4h12v16H8V6zm2 2v12h8V8h-8z"/>`
	cornerDownLeftInnerSVG             = `<path fill="currentColor" d="M18 16H8v2H6v-2H4v-2h2v-2h2v2h10V4h2v12h-2zM8 12v-2h2v2H8zm0 6v2h2v-2H8z"/>`
	cornerDownRightInnerSVG            = `<path fill="currentColor" d="M6 16h10v2h2v-2h2v-2h-2v-2h-2v2H6V4H4v12h2zm10-4v-2h-2v2h2zm0 6v2h-2v-2h2z"/>`
	cornerLeftDownInnerSVG             = `<path fill="currentColor" d="M8 6v10H6v2h2v2h2v-2h2v-2h-2V6h10V4H8v2zm4 10h2v-2h-2v2zm-6 0H4v-2h2v2z"/>`
	cornerLeftUpInnerSVG               = `<path fill="currentColor" d="M8 18V8H6V6h2V4h2v2h2v2h-2v10h10v2H8v-2zm4-10h2v2h-2V8zM6 8H4v2h2V8z"/>`
	cornerRightDownInnerSVG            = `<path fill="currentColor" d="M16 6v10h2v2h-2v2h-2v-2h-2v-2h2V6H4V4h12v2zm-4 10h-2v-2h2v2zm6 0h2v-2h-2v2z"/>`
	cornerRightUpInnerSVG              = `<path fill="currentColor" d="M16 18V8h2V6h-2V4h-2v2h-2v2h2v10H4v2h12v-2zM12 8h-2v2h2V8zm6 0h2v2h-2V8z"/>`
	cornerUpLeftInnerSVG               = `<path fill="currentColor" d="M18 8H8V6H6v2H4v2h2v2h2v-2h10v10h2V8h-2zM8 12v2h2v-2H8zm0-6V4h2v2H8z"/>`
	cornerUpRightInnerSVG              = `<path fill="currentColor" d="M6 8h10V6h2v2h2v2h-2v2h-2v-2H6v10H4V8h2zm10 4v2h-2v-2h2zm0-6V4h-2v2h2z"/>`
	creditCardInnerSVG                 = `<path fill="currentColor" d="M4 4h16v2H4v2h16v4H4v6h16v2H2V4h2zm18 0h-2v16h2V4z"/>`
	creditCardDeleteInnerSVG           = `<path fill="currentColor" d="M20 4H2v16h12v-2H4v-6h16V8H4V6h16V4zm0 0h2v8h-2V4zm2 14h-2v-2h2v-2h-2v2h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2z"/>`
	creditCardMinusInnerSVG            = `<path fill="currentColor" d="M20 4H2v16h12v-2H4v-6h16V8H4V6h16V4zm0 0h2v8h-2V4zm2 12h-6v2h6v-2z"/>`
	creditCardMultipleInnerSVG         = `<path fill="currentColor" d="M1 3h16v2H3v2h14v4H3v4h14v2H1V3zm18 0h-2v14h2V3zM5 19h16v2H5v-2zM23 7h-2v14h2V7z"/>`
	creditCardPlusInnerSVG             = `<path fill="currentColor" d="M2 4h18v2H4v2h16v4H4v6h10v2H2V4zm20 0h-2v8h2V4zm-4 10h2v2h2v2h-2v2h-2v-2h-2v-2h2v-2z"/>`
	creditCardSettingsInnerSVG         = `<path fill="currentColor" d="M20 4H2v16h18v-2H4v-6h16V8H4V6h16V4zm0 0h2v16h-2V4zm-7 18h-2v2h2v-2zm2 0h2v2h-2v-2zm-6 0H7v2h2v-2z"/>`
	creditCardWirelessInnerSVG         = `<path fill="currentColor" d="M16 2H8v2H6v2h2V4h8v2h2V4h-2V2zM8 8h2v2H8V8zm6 0V6h-4v2h4zm0 0h2v2h-2V8zM4 11h16v12H4V11zm14 10v-3H6v3h12zm0-6v-2H6v2h12z"/>`
	cropInnerSVG                       = `<path fill="currentColor" d="M8 2H6v4H2v2h14v14h2v-4h4v-2h-4V6H8V2zm0 8H6v8h8v-2H8v-6z"/>`
	cutInnerSVG                        = `<path fill="currentColor" d="M2 2h2v2H2V2zm4 4H4V4h2v2zm2 2H6V6h2v2zm2 2V8H8v2h2zm4 0h-4v4H2v8h8v-8h4v8h8v-8h-8v-4zm2-2v2h-2V8h2zm2-2v2h-2V6h2zm2-2h-2v2h2V4zm0 0V2h2v2h-2zM4 20v-4h4v4H4zm12 0v-4h4v4h-4z"/>`
	dashbaordInnerSVG                  = `<path fill="currentColor" d="M3 3h8v10H3V3zm2 2v6h4V5H5zm8-2h8v6h-8V3zm2 2v2h4V5h-4zm-2 6h8v10h-8V11zm2 2v6h4v-6h-4zM3 15h8v6H3v-6zm2 2v2h4v-2H5z"/>`
	debugInnerSVG                      = `<g fill="currentColor"><path d="M6 2h2v2H6V2Zm4 9h4v2h-4v-2Zm4 4h-4v2h4v-2Z"/><path d="M16 4h-2v2h-4V4H8v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h12v-3h2v2h2v-2h-2v-2h-2v-2h4v-2h-4v-2h2V9h2V7h-2v2h-2V6h-2V4ZM8 20V8h8v12H8Zm8-16V2h2v2h-2Z"/></g>`
	debugCheckInnerSVG                 = `<path fill="currentColor" d="M8 2H6v2h2v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h6v-2H8V8h8v6h2v-3h2V9h2V7h-2v2h-2V6h-2V4h2V2h-2v2h-2v2h-4V4H8V2Zm6 9h-4v2h4v-2Zm-4 4h2v2h-2v-2Zm4 3h2v2h-2v-2Zm4 2v2h-2v-2h2Zm2-2h-2v2h2v-2Zm0 0v-2h2v2h-2Z"/>`
	debugOffInnerSVG                   = `<path fill="currentColor" d="M16 2h2v2h-2V2Zm4 7h-2V6h-2V4h-2v2h-2v2h4v5h2v2h4v-2h-4v-2h2V9Zm0 0V7h2v2h-2ZM8 20v-9H6V9H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h10v-2H8Zm2-5h2v2h-2v-2ZM2 2h2v2H2V2Zm4 4H4V4h2v2Zm2 2H6V6h2v2Zm2 2H8V8h2v2Zm0 0v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2h-2v-2h-2Z"/>`
	debugPauseInnerSVG                 = `<path fill="currentColor" d="M8 2H6v2h2v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h8v-2H8V8h8v6h2v-3h2V9h2V7h-2v2h-2V6h-2V4h2V2h-2v2h-2v2h-4V4H8V2Zm6 9h-4v2h4v-2Zm-4 4h4v2h-4v-2Zm6 1h2v6h-2v-6Zm6 0h-2v6h2v-6Z"/>`
	debugPlayInnerSVG                  = `<path fill="currentColor" d="M6 2h2v2H6V2Zm10 2h-2v2h-4V4H8v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h8v-2H8V8h8v3h4V9h2V7h-2v2h-2V6h-2V4Zm0 0V2h2v2h-2Zm-6 7h4v2h-4v-2Zm4 4h-4v2h4v-2Zm4-2h-2v10h2v-2h2v-2h2v-2h-2v-2h-2v-2Z"/>`
	debugStopInnerSVG                  = `<path fill="currentColor" d="M6 2h2v2H6V2Zm10 2h-2v2h-4V4H8v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h8v-2H8V8h8v6h2v-3h2V9h2V7h-2v2h-2V6h-2V4Zm0 0V2h2v2h-2Zm-6 7h4v2h-4v-2Zm4 4h-4v2h4v-2Zm8 1h-6v6h6v-6Z"/>`
	deleteInnerSVG                     = `<path fill="currentColor" d="M21 5H7v2H5v2H3v2H1v2h2v2h2v2h2v2h16V5h-2zM7 17v-2H5v-2H3v-2h2V9h2V7h14v10H7zm8-6h-2V9h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2zm0 0V9h2v2h-2z"/>`
	deskphoneInnerSVG                  = `<path fill="currentColor" d="M3 3h18v18H3V3zm2 2v6h8V5H5zm10 0v14h4V5h-4zm-2 14v-2h-3v2h3zm-5 0v-2H5v2h3zm-3-4h3v-2H5v2zm5-2v2h3v-2h-3z"/>`
	deviceLaptopInnerSVG               = `<path fill="currentColor" d="M6 4H4v12h16V4H6zm12 2v8H6V6h12zm4 12H2v2h20v-2z"/>`
	devicePhoneInnerSVG                = `<path fill="currentColor" d="M6 3h12v18H6V3zm10 16V5h-2v2h-4V5H8v14h8zm-5-4h2v2h-2v-2z"/>`
	deviceTabletInnerSVG               = `<path fill="currentColor" d="M6 2H4v20h16V2H6zm12 2v16H6V4h12zm-5 12h-2v2h2v-2z"/>`
	deviceTvInnerSVG                   = `<path fill="currentColor" d="M2 20h20V6h-7V4h-2v2h-2V4H9v2H2v14zM9 4V2H7v2h2zm6 0h2V2h-2v2zm5 4v10H4V8h16z"/>`
	deviceTvSmartInnerSVG              = `<path fill="currentColor" d="M4 4h18v14h-6v2H8v-2H2V4h2zm16 12V6H4v10h16z"/>`
	deviceVibrateInnerSVG              = `<path fill="currentColor" d="M8 3H6v18h12V3H8zm8 2v14H8V5h8zm-3 10h-2v2h2v-2zm7-8h2v2h-2V7zm2 4V9h2v2h-2zm0 2h-2v-2h2v2zm0 2v-2h2v2h-2zm0 0v2h-2v-2h2zM2 17h2v-2H2v-2h2v-2H2V9h2V7H2v2H0v2h2v2H0v2h2v2z"/>`
	deviceWatchInnerSVG                = `<path fill="currentColor" d="M8 2h8v4h5v12h-5v4H8v-4H3V6h5V2zM5 16h14V8H5v8zm6-6h2v2h2v2h-4v-4z"/>`
	devicesInnerSVG                    = `<path fill="currentColor" d="M2 2h16v6h4v14H12v-6H2V2zm14 6V4H4v10h8V8h4zm-6-2H6v2h4V6zm10 14V10h-6v10h6zm-4-4h2v2h-2v-2zM6 10h4v2H6v-2z"/>`
	diceInnerSVG                       = `<path fill="currentColor" d="M5 3H3v18h18V3H5zm14 2v14H5V5h14zM9 7H7v2h2V7zm6 0h2v2h-2V7zm-6 8H7v2h2v-2zm6 0h2v2h-2v-2zm-2-4h-2v2h2v-2z"/>`
	dollarInnerSVG                     = `<path fill="currentColor" d="M11 2h2v4h6v2H7v3H5V6h6V2zM5 18h6v4h2v-4h6v-2H5v2zm14-7H5v2h12v3h2v-5z"/>`
	downasaurInnerSVG                  = `<path fill="currentColor" d="M6 4h14v2h2v6h-8v2h6v2h-4v2h-2v2H2V8h2V6h2V4zm2 6h2V8H8v2z"/>`
	downloadInnerSVG                   = `<path fill="currentColor" d="M13 17V3h-2v10H9v-2H7v2h2v2h2v2h2zm8 2v-4h-2v4H5v-4H3v6h18v-2zm-8-6v2h2v-2h2v-2h-2v2h-2z"/>`
	draftInnerSVG                      = `<path fill="currentColor" d="M14 2h-4v2H8v2H6v2H4v2H2v12h20V10h-2V8h-2V6h-2V4h-2V2zm0 2v2h2v2h2v4h-2v2h-2v2h-4v-2H8v-2H6V8h2V6h2V4h4zm-8 8v2h2v2h2v2h4v-2h2v-2h2v-2h2v8H4v-8h2z"/>`
	dragAndDropInnerSVG                = `<path fill="currentColor" d="M5 3H3v2h2V3zm14 4h2v6h-2V9H9v10h4v2H7V7h12zM7 3h2v2H7V3zM5 7H3v2h2V7zm-2 4h2v2H3v-2zm2 4H3v2h2v-2zm6-12h2v2h-2V3zm6 0h-2v2h2V3zm-2 14v-2h6v2h-2v2h-2v2h-2v-4zm4 2v2h2v-2h-2z"/>`
	dropInnerSVG                       = `<path fill="currentColor" d="M13 2h-2v2H9v4H7v4H5v6h2v2h2v2h6v-2h2v-2h2v-6h-2V8h-2V4h-2V2zm0 2v4h2v4h2v6h-2v2H9v-2H7v-6h2V8h2V4h2z"/>`
	dropAreaInnerSVG                   = `<path fill="currentColor" d="M5 3H3v2h2V3zm2 0h2v2H7V3zm6 0h-2v2h2V3zm2 0h2v2h-2V3zm4 0h2v2h-2V3zM3 7h2v2H3V7zm2 4H3v2h2v-2zm-2 4h2v2H3v-2zm2 4H3v2h2v-2zm2 0h2v2H7v-2zm6 0h-2v2h2v-2zm6-8h2v2h-2v-2zm2-4h-2v2h2V7zm-6 10v-2h6v2h-2v2h-2v2h-2v-4zm4 2v2h2v-2h-2z"/>`
	dropFullInnerSVG                   = `<path fill="currentColor" d="M11 2h2v2h2v4h2v4h2v6h-2v2h-2v2H9v-2H7v-2H5v-6h2V8h2V4h2V2z"/>`
	dropHalfInnerSVG                   = `<path fill="currentColor" d="M13 2h-2v2H9v4H7v4H5v6h2v2h2v2h6v-2h2v-2h2v-6h-2V8h-2V4h-2V2zm0 2v4h2v4h2v3H7v-3h2V8h2V4h2z"/>`
	duplicateInnerSVG                  = `<path fill="currentColor" d="M5 3h12v4h4v14H7v-4H3V3h2zm10 4V5H5v10h2V7h8zM9 17v2h10V9H9v8z"/>`
	duplicateAltInnerSVG               = `<path fill="currentColor" d="M5 1H3v14h10v2h-2v2h2v-2h2v-2h2v-2h-2v-2h-2V9h-2v2h2v2H5V3h12V1H5zm4 4H7v6h2V7h10v14H9v-4H7v6h14V5H9z"/>`
	editInnerSVG                       = `<path fill="currentColor" d="M18 2h-2v2h-2v2h-2v2h-2v2H8v2H6v2H4v2H2v6h6v-2h2v-2h2v-2h2v-2h2v-2h2v-2h2V8h2V6h-2V4h-2V2zm0 8h-2v2h-2v2h-2v2h-2v2H8v-2H6v-2h2v-2h2v-2h2V8h2V6h2v2h2v2zM6 16H4v4h4v-2H6v-2z"/>`
	editBoxInnerSVG                    = `<path fill="currentColor" d="M18 2h-2v2h2V2zM4 4h6v2H4v14h14v-6h2v8H2V4h2zm4 8H6v6h6v-2h2v-2h-2v2H8v-4zm4-2h-2v2H8v-2h2V8h2V6h2v2h-2v2zm2-6h2v2h-2V4zm4 0h2v2h2v2h-2v2h-2v2h-2v-2h2V8h2V6h-2V4zm-4 8h2v2h-2v-2z"/>`
	euroInnerSVG                       = `<path fill="currentColor" d="M9 4h10v2H9v3h7v2H9v2h7v2H9v3h10v2H7v-5H5v-2h2v-2H5V9h2V4h2z"/>`
	expandInnerSVG                     = `<path fill="currentColor" d="M11 5h2v2h2v2h2V7h-2V5h-2V3h-2v2zM9 7V5h2v2H9zm0 0v2H7V7h2zm-5 6h16v-2H4v2zm9 6h-2v-2H9v-2H7v2h2v2h2v2h2v-2zm2-2h-2v2h2v-2zm0 0h2v-2h-2v2z"/>`
	externalLinkInnerSVG               = `<path fill="currentColor" d="M21 11V3h-8v2h4v2h-2v2h-2v2h-2v2H9v2h2v-2h2v-2h2V9h2V7h2v4h2zM11 5H3v16h16v-8h-2v6H5V7h6V5z"/>`
	eyeInnerSVG                        = `<path fill="currentColor" d="M8 6h8v2H8V6zm-4 4V8h4v2H4zm-2 2v-2h2v2H2zm0 2v-2H0v2h2zm2 2H2v-2h2v2zm4 2H4v-2h4v2zm8 0v2H8v-2h8zm4-2v2h-4v-2h4zm2-2v2h-2v-2h2zm0-2h2v2h-2v-2zm-2-2h2v2h-2v-2zm0 0V8h-4v2h4zm-10 1h4v4h-4v-4z"/>`
	eyeClosedInnerSVG                  = `<path fill="currentColor" d="M0 7h2v2H0V7zm4 4H2V9h2v2zm4 2v-2H4v2H2v2h2v-2h4zm8 0H8v2H6v2h2v-2h8v2h2v-2h-2v-2zm4-2h-4v2h4v2h2v-2h-2v-2zm2-2v2h-2V9h2zm0 0V7h2v2h-2z"/>`
	fileInnerSVG                       = `<path fill="currentColor" d="M3 22h18V8h-2V6h-2v2h-2V6h2V4h-2V2H3v20zm2-2V4h8v6h6v10H5z"/>`
	fileAltInnerSVG                    = `<path fill="currentColor" d="M21 22H3V2h12v2h2v2h2v2h2v14zM17 6h-2v2h2V6zM5 4v16h14V10h-6V4H5zm8 12H7v2h6v-2zm-6-4h10v2H7v-2zm4-4H7v2h4V8z"/>`
	fileDeleteInnerSVG                 = `<path fill="currentColor" d="M11 22h10V8h-2V6h-2v2h-2V6h2V4h-2V2H3v12h2V4h8v6h6v10h-8v2zm-4-2H5v2H3v-2h2v-2H3v-2h2v2h2v-2h2v2H7v2zm0 0h2v2H7v-2z"/>`
	fileFlashInnerSVG                  = `<path fill="currentColor" d="M19 22h-6v-2h6V10h-6V4H5v8H3V2h12v2h2v2h2v2h2v14h-2zM17 6h-2v2h2V6zM7 12h2v4h4v2h-2v2H9v2H7v-4H3v-2h2v-2h2v-2z"/>`
	fileMinusInnerSVG                  = `<path fill="currentColor" d="M13 22h8V8h-2V6h-2v2h-2V6h2V4h-2V2H3v13h2V4h8v6h6v10h-6v2zm-2-3H3v-2h8v2z"/>`
	fileMultipleInnerSVG               = `<path fill="currentColor" d="M21 18H7V2h8v2h2v2h-2v2h2V6h2v2h2v10zM9 4v12h10v-6h-6V4H9zM3 6h2v14h12v2H3V6z"/>`
	fileOffInnerSVG                    = `<path fill="currentColor" d="M5 2H3v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2h-2v-2h6v4h2V8h-2V6h-2V4h-2V2H9v2h4v6h-2V8H9V6H7V4H5V2zm12 4v2h-2V6h2zM3 8h2v12h12v2H3V8z"/>`
	filePlusInnerSVG                   = `<path fill="currentColor" d="M19 22h-7v-2h7V10h-6V4H5v8H3V2h12v2h2v2h2v2h2v14h-2zM17 6h-2v2h2V6zM8 19h3v-2H8v-3H6v3H3v2h3v3h2v-3z"/>`
	fillInnerSVG                       = `<path fill="currentColor" d="M9 2h2v2H9V2zm4 4V4h-2v2H9v2H7v2H5v2H3v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v6h2V12h-2v-2h-2V8h-2V6h-2zm0 0v2h2v2h2v2h2v2h-2v2h-2v2h-2v2h-2v-2H9v-2H7v-2H5v-2h2v-2h2V8h2V6h2z"/>`
	fillHalfInnerSVG                   = `<path fill="currentColor" d="M9 2h2v2H9V2zm4 4V4h-2v2H9v2H7v2H5v2H3v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v6h2V12h-2v-2h-2V8h-2V6h-2zm0 0v2h2v2h2v2h2v2H5v-2h2v-2h2V8h2V6h2z"/>`
	fiveGInnerSVG                      = `<path fill="currentColor" d="M10 7H3v6h5v2H3v2h7v-6H5V9h5V7zm11 0h-9v10h9v-6h-4v2h2v2h-5V9h7V7z"/>`
	flagInnerSVG                       = `<path fill="currentColor" d="M3 2h10v2h8v14H11v-2H5v6H3V2zm2 12h8v2h6V6h-8V4H5v10z"/>`
	flattenInnerSVG                    = `<path fill="currentColor" d="M11 2h2v8h2v2h-2v2h-2v-2H9v-2h2V2zm-2 8H7V8h2v2zm6 0V8h2v2h-2zm5 6H4v2h16v-2zm-4 4H8v2h8v-2z"/>`
	flipToBackInnerSVG                 = `<path fill="currentColor" d="M9 3H7v2h2V3zm0 12H7v2h2v-2zm2-12h2v2h-2V3zm2 12h-2v2h2v-2zm2-12h2v2h-2V3zm2 12h-2v2h2v-2zm2-12h2v2h-2V3zm2 4h-2v2h2V7zM7 7h2v2H7V7zm14 4h-2v2h2v-2zM7 11h2v2H7v-2zm14 4h-2v2h2v-2zM3 7h2v12h12v2H3V7z"/>`
	flipToFrontInnerSVG                = `<path fill="currentColor" d="M21 3H7v14h14V3zm-2 12H9V5h10v10zM5 7H3v2h2V7zm-2 4h2v2H3v-2zm2 4H3v2h2v-2zm-2 4h2v2H3v-2zm6 0H7v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2z"/>`
	floatCenterInnerSVG                = `<path fill="currentColor" d="M10 4h6v8H8V4h2zm4 6V6h-4v4h4zM2 6h4v2H2V6zm20 0h-4v2h4V6zm0 4h-4v2h4v-2zM6 10H2v2h4v-2zm-4 4h20v2H2v-2zm20 4H2v2h20v-2z"/>`
	floatLeftInnerSVG                  = `<path fill="currentColor" d="M4 4h6v8H2V4h2zm4 6V6H4v4h4zm14-4H12v2h10V6zm0 4H12v2h10v-2zm0 4v2H2v-2h20zm0 6v-2H2v2h20z"/>`
	floatRightInnerSVG                 = `<path fill="currentColor" d="M16 4h6v8h-8V4h2zm4 6V6h-4v4h4zm-8-4H2v2h10V6zm0 4H2v2h10v-2zm10 4v2H2v-2h20zm0 6v-2H2v2h20z"/>`
	folderInnerSVG                     = `<path fill="currentColor" d="M4 4h8v2h10v14H2V4h2zm16 4H10V6H4v12h16V8z"/>`
	folderMinusInnerSVG                = `<path fill="currentColor" d="M12 4H2v16h20V6H12V4zm-2 4h10v10H4V6h6v2zm8 6v-2h-6v2h6z"/>`
	folderPlusInnerSVG                 = `<path fill="currentColor" d="M4 4h8v2h10v14H2V4h2zm16 4H10V6H4v12h16V8zm-6 2h2v2h2v2h-2v2h-2v-2h-2v-2h2v-2z"/>`
	folderXInnerSVG                    = `<path fill="currentColor" d="M12 4H2v16h20V6H12V4zm-2 4h10v10H4V6h6v2zm6 4h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2zm0 0h2v-2h-2v2z"/>`
	forwardInnerSVG                    = `<path fill="currentColor" d="M14 5h-2v4H6v2H4v6h2v-2h6v4h2v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2V5z"/>`
	forwardburgerInnerSVG              = `<path fill="currentColor" d="M13 7H3v2h10V7zm8 4h-2V9h-2V7h-2v2h2v2H3v2h14v2h-2v2h2v-2h2v-2h2v-2zM3 15h10v2H3v-2z"/>`
	fourGInnerSVG                      = `<path fill="currentColor" d="M5 7H3v6h5v4h2V7H8v4H5V7zm16 0h-9v10h9v-6h-4v2h2v2h-5V9h7V7z"/>`
	fourKInnerSVG                      = `<path fill="currentColor" d="M3 7h2v4h4V7h2v10H9v-4H3V7zm10 0h2v4h2v2h-2v4h-2V7zm6 8h-2v-2h2v2zm0 0h2v2h-2v-2zm0-6h-2v2h2V9zm0 0V7h2v2h-2z"/>`
	fourKBoxInnerSVG                   = `<path fill="currentColor" d="M3 4H1v16h22V4H3zm18 2v12H3V6h18zM7 8H5v5h4v3h2V8H9v3H7V8zm8 0h-2v8h2v-3h2v3h2v-3h-2v-2h2V8h-2v3h-2V8z"/>`
	frameInnerSVG                      = `<path fill="currentColor" d="M2 3h20v18H2V3zm18 16V7H4v12h16z"/>`
	frameAddInnerSVG                   = `<path fill="currentColor" d="M2 3h20v18H2V3zm18 16V7H4v12h16zm-7-7h3v2h-3v3h-2v-3H8v-2h3V9h2v3z"/>`
	frameCheckInnerSVG                 = `<path fill="currentColor" d="M2 3h20v18H2V3zm18 16V7H4v12h16zm-4-9h-2v2h-2v2h-2v-2H8v2h2v2h2v-2h2v-2h2v-2z"/>`
	frameDeleteInnerSVG                = `<path fill="currentColor" d="M2 3h20v18H2V3zm18 16V7H4v12h16zM9 10h2v2H9v-2zm4 2h-2v2H9v2h2v-2h2v2h2v-2h-2v-2zm0 0v-2h2v2h-2z"/>`
	frameMinusInnerSVG                 = `<path fill="currentColor" d="M2 3h20v18H2V3zm18 16V7H4v12h16zM8 12h8v2H8v-2z"/>`
	gamepadInnerSVG                    = `<path fill="currentColor" d="M2 5h20v14H2V5zm18 12V7H4v10h16zM8 9h2v2h2v2h-2v2H8v-2H6v-2h2V9zm6 0h2v2h-2V9zm4 4h-2v2h2v-2z"/>`
	gifInnerSVG                        = `<path fill="currentColor" d="M3 7h6v2H3v6h4v-2H5v-2h4v6H1V7h2zm14 0h6v2h-6v2h4v2h-4v4h-2V7h2zm-4 0h-2v10h2V7z"/>`
	giftInnerSVG                       = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 12v8h-7m7-8h2V8h-3m1 4H5m13-4V4h-6m6 4H6m0 0V4h6M6 8H3v4h2m0 0v8h7m0 0V4"/>`
	gitBranchInnerSVG                  = `<path fill="currentColor" d="M5 2h2v12h3v3h7v-7h-3V2h8v8h-3v9h-9v3H2v-8h3V2zm15 6V4h-4v4h4zM8 19v-3H4v4h4v-1z"/>`
	gitCommitInnerSVG                  = `<path fill="currentColor" d="M7 7h10v4h5v2h-5v4H7v-4H2v-2h5V7zm2 2v6h6V9H9z"/>`
	gitMergeInnerSVG                   = `<path fill="currentColor" d="M10 2H2v8h3v12h2V10h3v2h2v2h2v8h8v-8h-8v-2h-2v-2h-2V2zM4 8V4h4v4H4zm12 12v-4h4v4h-4z"/>`
	gitPullRequestInnerSVG             = `<path fill="currentColor" d="M2 2h8v8H7v12H5V10H2V2zm2 2v4h4V4H4zm8 1h7.09v9H22v8h-8v-8h3.09V7H12V5zm4 11v4h4v-4h-4z"/>`
	gpsInnerSVG                        = `<path fill="currentColor" d="M13 2v4h5v5h4v2h-4v5h-5v4h-2v-4H6v-5H2v-2h4V6h5V2h2zM8 8v8h8V8H8zm2 2h4v4h-4v-4z"/>`
	gridInnerSVG                       = `<path fill="currentColor" d="M2 2h20v20H2V2zm2 2v4h4V4H4zm6 0v4h4V4h-4zm6 0v4h4V4h-4zm4 6h-4v4h4v-4zm0 6h-4v4h4v-4zm-6 4v-4h-4v4h4zm-6 0v-4H4v4h4zm-4-6h4v-4H4v4zm6-4v4h4v-4h-4z"/>`
	groupInnerSVG                      = `<path fill="currentColor" d="M3 3h18v18H3V3zm2 2v14h14V5H5zm2 2h4v4H7V7zm6 0h4v4h-4V7zm-6 6h4v4H7v-4zm6 0h4v4h-4v-4z"/>`
	hdInnerSVG                         = `<path fill="currentColor" d="M3 7h2v4h4V7h2v10H9v-4H5v4H3V7zm10 8V7h6v2h-4v6h4v2h-6v-2zm6 0V9h2v6h-2z"/>`
	headphoneInnerSVG                  = `<path fill="currentColor" d="M19 4H5v2H3v14h7v-8H5V6h14v6h-5v8h7V6h-2V4zm-3 10h3v4h-3v-4zm-8 0v4H5v-4h3z"/>`
	headsetInnerSVG                    = `<path fill="currentColor" d="M19 2H5v2H3v14h7v-8H5V4h14v6h-5v8h3v2h-6v2h8v-4h2V4h-2V2zm-3 10h3v4h-3v-4zm-8 0v4H5v-4h3z"/>`
	heartInnerSVG                      = `<path fill="currentColor" d="M9 2H5v2H3v2H1v6h2v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v-2h2v-2h2V6h-2V4h-2V2h-4v2h-2v2h-2V4H9V2zm0 2v2h2v2h2V6h2V4h4v2h2v6h-2v2h-2v2h-2v2h-2v2h-2v-2H9v-2H7v-2H5v-2H3V6h2V4h4z"/>`
	hiddenInnerSVG                     = `<path fill="currentColor" d="M8 6h8v2H8V6zm-4 4V8h4v2H4zm-2 2v-2h2v2H2zm0 2v-2H0v2h2zm2 2H2v-2h2v2zm4 2H4v-2h4v2zm8 0v2H8v-2h8zm4-2v2h-4v-2h4zm2-2v2h-2v-2h2zm0-2h2v2h-2v-2zm-2-2h2v2h-2v-2zm0 0V8h-4v2h4zM9 10h2v2H9v-2zm4 2h-2v2H9v2h2v-2h2v2h2v-2h-2v-2zm0 0v-2h2v2h-2z"/>`
	homeInnerSVG                       = `<path fill="currentColor" d="M14 2h-4v2H8v2H6v2H4v2H2v2h2v10h7v-6h2v6h7V12h2v-2h-2V8h-2V6h-2V4h-2V2zm0 2v2h2v2h2v2h2v2h-2v8h-3v-6H9v6H6v-8H4v-2h2V8h2V6h2V4h4z"/>`
	hourglassInnerSVG                  = `<path fill="currentColor" d="M18 2H6v6h2v2h2v4H8v2H6v6h12v-6h-2v-2h-2v-4h2V8h2V2zm-2 6h-2v2h-4V8H8V4h8v4zm-2 6v2h2v4H8v-4h2v-2h4z"/>`
	hqInnerSVG                         = `<path fill="currentColor" d="M3 7h2v4h4V7h2v10H9v-4H5v4H3V7zm10 2h2v6h-2V9zm6 6h-4v2h8v-2h-2V9h-2V7h-4v2h4v6z"/>`
	humanInnerSVG                      = `<path fill="currentColor" d="M10 2h4v4h-4V2zM3 7h18v2h-6v13h-2v-6h-2v6H9V9H3V7z"/>`
	humanHandsdownInnerSVG             = `<path fill="currentColor" d="M10 2h4v4h-4V2zM7 7h10v2h-2v13h-2v-6h-2v6H9V9H7V7zm-2 4h2V9H5v2zm0 0v2H3v-2h2zm14 0h-2V9h2v2zm0 0h2v2h-2v-2z"/>`
	humanHandsupInnerSVG               = `<path fill="currentColor" d="M10 2h4v4h-4V2zM7 7h10v2h-2v13h-2v-6h-2v6H9V9H7V7zM5 5v2h2V5H5zm0 0H3V3h2v2zm14 0v2h-2V5h2zm0 0V3h2v2h-2z"/>`
	humanHeightInnerSVG                = `<path fill="currentColor" d="M6 2h4v4H6V2zM3 7h10v9h-2v6H9v-6H7v6H5v-6H3V7zm18-4h-6v2h6V3zm-4 4h4v2h-4V7zm4 4h-6v2h6v-2zm-6 8h6v2h-6v-2zm6-4h-4v2h4v-2z"/>`
	humanHeightAltInnerSVG             = `<path fill="currentColor" d="M4 2h4v4H4V2zM1 7h10v9H9v6H7v-6H5v6H3v-6H1V7zm18-5h-2v2h-2v2h-2v2h2V6h2v12h-2v-2h-2v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2V6h2v2h2V6h-2V4h-2V2z"/>`
	humanRunInnerSVG                   = `<path fill="currentColor" d="M10 3H8v2H6v2h2V5h2v2h2v2h-2v2H8v2H6v2H4v-2H2v2h2v2h2v-2h4v2h2v2h-2v2h2v-2h2v-2h-2v-4h2v-2h2v2h2v2h2v-2h2v-2h-2v2h-2v-2h-2V9h2V5h-4v2h-2V5h-2V3z"/>`
	imageInnerSVG                      = `<path fill="currentColor" d="M4 3H2v18h20V3H4zm16 2v14H4V5h16zm-6 4h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2V9zM8 7H6v2h2V7z"/>`
	imageArrowRightInnerSVG            = `<path fill="currentColor" d="M19 1h-2v2h2v2h-6v2h6v2h-2v2h2V9h2V7h2V5h-2V3h-2V1zm-8 2H2v18h20v-8h-2v6H4V5h7V3zm1 8V9h2v2h-2zm-2 2v-2h2v2h-2zm-2 2v-2h2v2H8zm0 0v2H6v-2h2zm8-2h-2v-2h2v2zm0 0h2v2h-2v-2zM6 7h2v2H6V7z"/>`
	imageBrokenInnerSVG                = `<path fill="currentColor" d="M22 3H2v18h20v-2h-2v-2h2v-2h-2v-2h2v-2h-2V9h2V7h-2V5h2V3zm-2 4v2h-2v2h2v2h-2v2h2v2h-2v2H4V5h14v2h2zm-6 2h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v-2h-2V9zM6 7h2v2H6V7z"/>`
	imageDeleteInnerSVG                = `<path fill="currentColor" d="M14 3H2v18h20V11h-2v8H4V5h10V3zM6 7h2v2H6V7zm14-2h-2V3h-2v2h2v2h-2v2h2V7h2v2h2V7h-2V5zm0 0V3h2v2h-2zm-8 4h2v2h-2V9zm-2 4v-2h2v2h-2zm-2 2h2v-2H8v2zm0 0v2H6v-2h2zm8-2h-2v-2h2v2zm0 0h2v2h-2v-2z"/>`
	imageFlashInnerSVG                 = `<path fill="currentColor" d="M18 0h2v4h4v2h-2v2h-2v2h-2V6h-4V4h2V2h2V0zM4 3h8v2H4v14h16v-7h2v9H2V3h2zm10 6h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2V9zM8 7H6v2h2V7z"/>`
	imageFrameInnerSVG                 = `<path fill="currentColor" d="M13 1h-2v2H9v2H7v2H2v16h20V7h-5V5h-2V3h-2V1zm2 6H9V5h2V3h2v2h2v2zM4 9h16v12H4V9zm10 6v-2h-2v2h-2v2H8v2h2v-2h2v-2h2zm2 2v-2h-2v2h2zm0 0v2h2v-2h-2zM6 13v-2h2v2H6z"/>`
	imageGalleryInnerSVG               = `<path fill="currentColor" d="M2 2h20v16h-5v2h-2v-2H9v2H7v-2H2V2zm5 18v2H5v-2h2zm10 0v2h2v-2h-2zm3-16H4v12h16V4zm-8 4h2v2h-2V8zm-2 4v-2h2v2h-2zm0 0v2H8v-2h2zm6 0h-2v-2h2v2zm0 0h2v2h-2v-2zM8 6H6v2h2V6z"/>`
	imageMultipleInnerSVG              = `<path fill="currentColor" d="M24 2H4v16h20V2zM6 16V4h16v12H6zM2 4H0v18h20v-2H2V4zm12 2h2v2h-2V6zm-2 4V8h2v2h-2zm-2 2v-2h2v2h-2zm0 0v2H8v-2h2zm8-2h-2V8h2v2zm0 0h2v2h-2v-2zM8 6h2v2H8V6z"/>`
	imageNewInnerSVG                   = `<path fill="currentColor" d="M6 0h12v2H6V0zM4 3H2v18h20V3H4zm16 2v14H4V5h16zm-6 4h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2V9zM8 7H6v2h2V7zm10 17v-2H6v2h12z"/>`
	imagePlusInnerSVG                  = `<path fill="currentColor" d="M4 3h10v2H4v14h16v-8h2v10H2V3h2zm10 6h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2V9zM8 7H6v2h2V7zm10-4h2v2h2v2h-2v2h-2V7h-2V5h2V3z"/>`
	inboxInnerSVG                      = `<path fill="currentColor" d="M3 2h18v20H3V2zm2 2v10h4v2h6v-2h4V4H5zm14 12h-2v2H7v-2H5v4h14v-4z"/>`
	inboxAllInnerSVG                   = `<path fill="currentColor" d="M3 2h18v20H3V2zm2 2v4h4v2h6V8h4V4H5zm14 6h-2v2H7v-2H5v4h14v-4zm0 6h-2v2H7v-2H5v4h14v-4z"/>`
	inboxFullInnerSVG                  = `<path fill="currentColor" d="M3 2h18v20H3V2zm2 2v10h4v2h6v-2h4V4H5zm14 12h-2v2H7v-2H5v4h14v-4zM7 6h10v2H7V6zm0 4h10v2H7v-2z"/>`
	infoBoxInnerSVG                    = `<path fill="currentColor" d="M3 3h2v18H3V3zm16 0H5v2h14v14H5v2h16V3h-2zm-8 6h2V7h-2v2zm2 8h-2v-6h2v6z"/>`
	invertInnerSVG                     = `<path fill="currentColor" d="M3 3h18v18H3V3zm16 4h-2v2h-2v2h-2v2h-2v2H9v2H7v2h12V7z"/>`
	isoInnerSVG                        = `<path fill="currentColor" d="M8 3H6v3H3v2h3v3h2V8h3V6H8V3zm11 2h-2v2h-2v2h-2v2h-2v2H9v2H7v2H5v2h2v-2h2v-2h2v-2h2v-2h2V9h2V7h2V5zm-6 13v-2h8v2h-8z"/>`
	kanbanInnerSVG                     = `<path fill="currentColor" d="M21 3H3v18h18V3zM5 19V5h14v14H5zM9 7H7v8h2V7zm2 0h2v4h-2V7zm6 0h-2v10h2V7z"/>`
	keyboardInnerSVG                   = `<path fill="currentColor" d="M21 3H3v18h18V3zM5 19V5h14v14H5zM9 7H7v2h2V7zm8 8H7v2h10v-2zm-2-8h2v2h-2V7zm-2 0h-2v2h2V7zm-6 4h2v2H7v-2zm10 0h-2v2h2v-2zm-6 0h2v2h-2v-2z"/>`
	labelInnerSVG                      = `<path fill="currentColor" d="M12 2H2v10h2v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v-2h2v-2h-2v-2h-2V8h-2V6h-2V4h-2V2zm0 2v2h2v2h2v2h2v2h2v2h-2v2h-2v2h-2v2h-2v-2h-2v-2H8v-2H6v-2H4V4h8zM6 6h2v2H6V6z"/>`
	labelAltInnerSVG                   = `<path fill="currentColor" d="M16 5H2v14h14v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2V5zm0 2v2h2v2h2v2h-2v2h-2v2H4V7h12z"/>`
	labelAltMultipleInnerSVG           = `<path fill="currentColor" d="M8 5H6v10h12v-2h2v-2h2V9h-2V7h-2V5H8zm10 2v2h2v2h-2v2H8V7h10zM4 9H2v10h12v-2H4V9z"/>`
	labelSharpInnerSVG                 = `<path fill="currentColor" d="M16 5H2v4h2v2h2v2H4v2H2v4h14v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2V5zm0 2v2h2v2h2v2h-2v2h-2v2H4v-2h2v-2h2v-2H6V9H4V7h12z"/>`
	layoutInnerSVG                     = `<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v4h16V7H4zm16 6H10v4h10v-4zM8 17v-4H4v4h4z"/>`
	layoutAlignBottomInnerSVG          = `<path fill="currentColor" d="M16 4H8v12h8V4zm-6 10V6h4v8h-4zm10 6v-2H4v2h16z"/>`
	layoutAlignLeftInnerSVG            = `<path fill="currentColor" d="M20 16V8H8v8h12zm-10-6h8v4h-8v-4zM4 20h2V4H4v16z"/>`
	layoutAlignRightInnerSVG           = `<path fill="currentColor" d="M4 8v8h12V8H4zm10 6H6v-4h8v4zm6-10h-2v16h2V4z"/>`
	layoutAlignTopInnerSVG             = `<path fill="currentColor" d="M16 20H8V8h8v12zm-6-10v8h4v-8h-4zm10-6v2H4V4h16z"/>`
	layoutColumnsInnerSVG              = `<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v10h7V7H4zm9 0v10h7V7h-7z"/>`
	layoutDistributeHorizontalInnerSVG = `<path fill="currentColor" d="M6 4H4v16h2V4zm14 0h-2v16h2V4zM10 7h6v10H8V7h2zm4 8V9h-4v6h4z"/>`
	layoutDistributeVerticalInnerSVG   = `<path fill="currentColor" d="M20 6V4H4v2h16zm0 14v-2H4v2h16zM17 8v8h-2V8h2zm-8 6v-4h6V8H7v8h8v-2H9z"/>`
	layoutFooterInnerSVG               = `<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v6h16V7H4zm16 8H4v2h16v-2z"/>`
	layoutHeaderInnerSVG               = `<path fill="currentColor" d="M2 19h20V5H2v14zm2-2v-6h16v6H4zm16-8H4V7h16v2z"/>`
	layoutRowsInnerSVG                 = `<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v4h16V7H4zm16 6H4v4h16v-4z"/>`
	layoutSidebarLeftInnerSVG          = `<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v10h2V7H4zm4 0v10h12V7H8z"/>`
	layoutSidebarRightInnerSVG         = `<path fill="currentColor" d="M22 5H2v14h20V5zm-2 2v10h-2V7h2zm-4 0v10H4V7h12z"/>`
	linkInnerSVG                       = `<path fill="currentColor" d="M4 6h7v2H4v8h7v2H2V6h2zm16 0h-7v2h7v8h-7v2h9V6h-2zm-3 5H7v2h10v-2z"/>`
	listInnerSVG                       = `<path fill="currentColor" d="M6 6H4v2h2V6zm14 0H8v2h12V6zM4 11h2v2H4v-2zm16 0H8v2h12v-2zM4 16h2v2H4v-2zm16 0H8v2h12v-2z"/>`
	listBoxInnerSVG                    = `<path fill="currentColor" d="M2 3h20v18H2V3zm18 16V5H4v14h16zM8 7H6v2h2V7zm2 0h8v2h-8V7zm-2 4H6v2h2v-2zm2 0h8v2h-8v-2zm-2 4H6v2h2v-2zm2 0h8v2h-8v-2z"/>`
	loaderInnerSVG                     = `<path fill="currentColor" d="M13 2h-2v6h2V2zm0 14h-2v6h2v-6zm9-5v2h-6v-2h6zM8 13v-2H2v2h6zm7-6h2v2h-2V7zm4-2h-2v2h2V5zM9 7H7v2h2V7zM5 5h2v2H5V5zm10 12h2v2h2v-2h-2v-2h-2v2zm-8 0v-2h2v2H7v2H5v-2h2z"/>`
	lockInnerSVG                       = `<path fill="currentColor" d="M15 2H9v2H7v4H4v14h16V8h-3V4h-2V2zm0 2v4H9V4h6zm-6 6h9v10H6V10h3zm4 3h-2v4h2v-4z"/>`
	lockOpenInnerSVG                   = `<path fill="currentColor" d="M15 2H9v2H7v2h2V4h6v4H4v14h16V8h-3V4h-2V2zm0 8h3v10H6V10h9zm-2 3h-2v4h2v-4z"/>`
	loginInnerSVG                      = `<path fill="currentColor" d="M5 3H3v4h2V5h14v14H5v-2H3v4h18V3H5zm12 8h-2V9h-2V7h-2v2h2v2H3v2h10v2h-2v2h2v-2h2v-2h2v-2z"/>`
	logoutInnerSVG                     = `<path fill="currentColor" d="M5 3h16v4h-2V5H5v14h14v-2h2v4H3V3h2zm16 8h-2V9h-2V7h-2v2h2v2H7v2h10v2h-2v2h2v-2h2v-2h2v-2z"/>`
	luggageInnerSVG                    = `<path fill="currentColor" d="M9 2h6v4h4v14h-2v2h-2v-2H9v2H7v-2H5V6h4V2zm2 4h2V4h-2v2zM7 18h10V8H7v10zm4-8v6H9v-6h2zm4 0v6h-2v-6h2z"/>`
	mailInnerSVG                       = `<path fill="currentColor" d="M22 4H2v16h20V4zM4 18V6h16v12H4zM8 8H6v2h2v2h2v2h4v-2h2v-2h2V8h-2v2h-2v2h-4v-2H8V8z"/>`
	mailArrowRightInnerSVG             = `<path fill="currentColor" d="M20 4H2v16h10v-2H4V6h16v6h2V4h-2zM6 8h2v2H6V8zm4 4H8v-2h2v2zm4 0v2h-4v-2h4zm2-2v2h-2v-2h2zm0 0V8h2v2h-2zm8 8h-2v-2h-2v-2h-2v2h2v2h-6v2h6v2h-2v2h2v-2h2v-2h2v-2z"/>`
	mailCheckInnerSVG                  = `<path fill="currentColor" d="M4 4h18v10h-2V6H4v12h8v2H2V4h2zm4 4H6v2h2v2h2v2h4v-2h2v-2h2V8h-2v2h-2v2h-4v-2H8V8zm6 10h2v2h-2v-2zm4 2v2h-2v-2h2zm2-2h-2v2h2v-2zm0 0v-2h2v2h-2z"/>`
	mailDeleteInnerSVG                 = `<path fill="currentColor" d="M20 4H2v16h12v-2H4V6h16v8h2V4h-2zM6 8h2v2H6V8zm4 4H8v-2h2v2zm4 0v2h-4v-2h4zm2-2v2h-2v-2h2zm0 0V8h2v2h-2zm2 6h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2h2v-2h-2v2h-2v-2z"/>`
	mailFlashInnerSVG                  = `<path fill="currentColor" d="M4 4h18v8h-2V6H4v12h8v2H2V4h2zm4 4H6v2h2v2h2v2h4v-2h2v-2h2V8h-2v2h-2v2h-4v-2H8V8zm10 6h2v4h4v2h-2v2h-2v2h-2v-4h-4v-2h2v-2h2v-2z"/>`
	mailMultipleInnerSVG               = `<path fill="currentColor" d="M24 2H4v16h20V2zM6 16V4h16v12H6zM2 7H0v15h19v-2H2V7zm8-1H8v2h2v2h2v2h4v-2h2V8h2V6h-2v2h-2v2h-4V8h-2V6z"/>`
	mailOffInnerSVG                    = `<path fill="currentColor" d="M2 2h2v2H2V2zm4 4H4V4h2v2zm2 2H6V6h2v2zm2 2H8V8h2v2zm2 2h-2v-2h2v2zm2 0h-2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2zm2-2h-2v2h2v-2zm0 0V8h2v2h-2zm-6-6h12v12h-2V6H10V4zm4 14v2H2V8h2v10h10z"/>`
	mailUnreadInnerSVG                 = `<path fill="currentColor" d="M22 2h-6v6h6V2zM4 4h10v2H4v12h16v-8h2v10H2V4h2zm4 4H6v2h2v2h2v2h4v-2h2v-2h-2v2h-4v-2H8V8z"/>`
	mapInnerSVG                        = `<path fill="currentColor" d="M8 2h2v2h2v2h-2v10H8V6H6V4h2V2zM4 8V6h2v2H4zm2 10v2H4v2H2V8h2v10h2zm0 0h2v-2H6v2zm6 0h-2v-2h2v2zm2-10V6h-2v2h2zm2 0h-2v10h-2v2h2v2h2v-2h2v-2h2v-2h2V2h-2v2h-2v2h-2v2zm0 0h2V6h2v10h-2v2h-2V8z"/>`
	membercardInnerSVG                 = `<path fill="currentColor" d="M2 3h20v14h-7v3h-2v-2h-2v2H9v-3H2V3zm2 2v4h16V5H4zm16 8H4v2h16v-2z"/>`
	menuInnerSVG                       = `<path fill="currentColor" d="M4 6h16v2H4V6zm0 5h16v2H4v-2zm16 5H4v2h16v-2z"/>`
	messageInnerSVG                    = `<path fill="currentColor" d="M20 2H2v20h2V4h16v12H6v2H4v2h2v-2h16V2h-2z"/>`
	messageArrowLeftInnerSVG           = `<path fill="currentColor" d="M4 2h18v12h-2V4H4v18H2V2h2zm2 14h4v2H6v2H4v-2h2v-2zm16 0h-6v-2h2v-2h-2v2h-2v2h-2v2h2v2h2v2h2v-2h-2v-2h6v-2z"/>`
	messageArrowRightInnerSVG          = `<path fill="currentColor" d="M4 2h18v10h-2V4H4v18H2V2h2zm2 14h4v2H6v2H4v-2h2v-2zm16 0h-2v-2h-2v-2h-2v2h2v2h-6v2h6v2h-2v2h2v-2h2v-2h2v-2z"/>`
	messageBookmarkInnerSVG            = `<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2h2zm14 4h-6v8h2v-2h2v2h2V6z"/>`
	messageClockInnerSVG               = `<path fill="currentColor" d="M20 2H2v20h2V4h16v4h2V2h-2zM8 16H6v2H4v2h2v-2h2v-2zm6-2h2v2h2v2h-4v-4zm6-4h-8v2h-2v8h2v2h8v-2h2v-8h-2v-2zm0 2v8h-8v-8h8z"/>`
	messageDeleteInnerSVG              = `<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2h2zm9 7h-2V7H9v2h2v2H9v2h2v-2h2v2h2v-2h-2V9zm0 0V7h2v2h-2z"/>`
	messageFlashInnerSVG               = `<path fill="currentColor" d="M20 2H2v20h2V4h16v10h2V2h-2zM10 16H6v2H4v2h2v-2h4v-2zm6-4h2v4h4v2h-2v2h-2v2h-2v-4h-4v-2h2v-2h2v-2z"/>`
	messageImageInnerSVG               = `<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2h2zm10 4h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2V8h2v2h2v2h2v-2h-2V8h-2V6zM6 6h2v2H6V6z"/>`
	messageMinusInnerSVG               = `<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2h2zm12 7H8v2h8V9z"/>`
	messagePlusInnerSVG                = `<path fill="currentColor" d="M20 2H2v20h2V4h16v12H6v2H4v2h2v-2h16V2h-2zm-7 7h3v2h-3v3h-2v-3H8V9h3V6h2v3z"/>`
	messageProcessingInnerSVG          = `<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2h2zm5 7H7v2h2V9zm2 0h2v2h-2V9zm6 0h-2v2h2V9z"/>`
	messageReplyInnerSVG               = `<path fill="currentColor" d="M4 2h18v20h-2V4H4v12h14v2h2v2h-2v-2H2V2h2z"/>`
	messageTextInnerSVG                = `<path fill="currentColor" d="M20 2H2v20h2V4h16v12H6v2H4v2h2v-2h16V2h-2zM6 7h12v2H6V7zm8 4H6v2h8v-2z"/>`
	minusInnerSVG                      = `<path fill="currentColor" d="M4 11h16v2H4z"/>`
	missedCallInnerSVG                 = `<path fill="currentColor" d="M20 6h-4v2h2v2h-2v2h-2v2h-2v2h-2v-2H8v-2H6v-2H4V8H2v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v-2h2v2h2V6h-2z"/>`
	modemInnerSVG                      = `<path fill="currentColor" d="M19 2h-8v2H9v2h2V4h8v2h2V4h-2V2zm-8 6h2v2h-2V8zm6 0V6h-4v2h4zm0 0h2v2h-2V8zm-1 2h-2v2H2v10h20V12h-6v-2zm4 4v6H4v-6h16zm-2 2h-2v2h2v-2zm-6 0h2v2h-2v-2z"/>`
	moneyInnerSVG                      = `<path fill="currentColor" d="M16 4H2v12h4v4h16V8h-4V4h-2zm0 2v2H6v6H4V6h12zm-8 4h12v8H8v-8zm8 2h-4v4h4v-4z"/>`
	monitorInnerSVG                    = `<path fill="currentColor" d="M20 3H2v14h8v2H8v2h8v-2h-2v-2h8V3h-2zm-6 12H4V5h16v10h-6z"/>`
	moodHappyInnerSVG                  = `<path fill="currentColor" d="M5 3h14v2H5V3zm0 16H3V5h2v14zm14 0v2H5v-2h14zm0 0h2V5h-2v14zM10 8H8v2h2V8zm4 0h2v2h-2V8zm-5 6v-2H7v2h2zm6 0v2H9v-2h6zm0 0h2v-2h-2v2z"/>`
	moodNeutralInnerSVG                = `<path fill="currentColor" d="M5 3h14v2H5V3zm0 16H3V5h2v14zm14 0v2H5v-2h14zm0 0h2V5h-2v14zM10 8H8v2h2V8zm4 0h2v2h-2V8zm1 5H9v2h6v-2z"/>`
	moodSadInnerSVG                    = `<path fill="currentColor" d="M5 3h14v2H5V3zm0 16H3V5h2v14zm14 0v2H5v-2h14zm0 0h2V5h-2v14zM10 8H8v2h2V8zm4 0h2v2h-2V8zm-5 8v-2h6v2h2v-2h-2v-2H9v2H7v2h2z"/>`
	moonInnerSVG                       = `<path fill="currentColor" d="M6 2h8v2h-2v2h-2V4H6V2ZM4 6V4h2v2H4Zm0 10H2V6h2v10Zm2 2H4v-2h2v2Zm2 2H6v-2h2v2Zm10 0v2H8v-2h10Zm2-2v2h-2v-2h2Zm-2-4h2v4h2v-8h-2v2h-2v2Zm-6 0v2h6v-2h-6Zm-2-2h2v2h-2v-2Zm0 0V6H8v6h2Z"/>`
	moonStarInnerSVG                   = `<path fill="currentColor" d="M6 2h8v2h-2v2h-2V4H6V2ZM4 6V4h2v2H4Zm0 10H2V6h2v10Zm2 2H4v-2h2v2Zm2 2H6v-2h2v2Zm10 0v2H8v-2h10Zm2-2v2h-2v-2h2Zm-2-4v-2h2v-2h2v8h-2v-4h-2Zm-6 0h6v2h-6v-2Zm-2-2h2v2h-2v-2Zm0 0V6H8v6h2Zm8-10h2v2h2v2h-2v2h-2V6h-2V4h2V2Z"/>`
	moonStarsInnerSVG                  = `<path fill="currentColor" d="M20 0h2v2h2v2h-2v2h-2V4h-2V2h2V0ZM8 4h8v2h-2v2h-2V6H8V4ZM6 8V6h2v2H6Zm0 8H4V8h2v8Zm2 2H6v-2h2v2Zm8 0v2H8v-2h8Zm2-2v2h-2v-2h2Zm-2-4v-2h2V8h2v8h-2v-4h-2Zm-4 0h4v2h-4v-2Zm0 0V8h-2v4h2Zm-8 6H2v2H0v2h2v2h2v-2h2v-2H4v-2Z"/>`
	moreHorizontalInnerSVG             = `<path fill="currentColor" d="M1 9h6v6H1V9zm2 2v2h2v-2H3zm6-2h6v6H9V9zm2 2v2h2v-2h-2zm6-2h6v6h-6V9zm2 2v2h2v-2h-2z"/>`
	moreVerticalInnerSVG               = `<path fill="currentColor" d="M15 1v6H9V1h6zm-2 2h-2v2h2V3zm2 6v6H9V9h6zm-2 2h-2v2h2v-2zm2 6v6H9v-6h6zm-2 2h-2v2h2v-2z"/>`
	mouseInnerSVG                      = `<path fill="currentColor" d="M6 3h12v18H6V3zm2 2v4h3V5H8zm5 0v4h3V5h-3zm3 6H8v8h8v-8z"/>`
	moveInnerSVG                       = `<path fill="currentColor" d="M13 0h-2v2H9v2H7v2h2V4h2v7H4V9h2V7H4v2H2v2H0v2h2v2h2v2h2v-2H4v-2h7v7H9v-2H7v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2v-7h7v2h-2v2h2v-2h2v-2h2v-2h-2V9h-2V7h-2v2h2v2h-7V4h2v2h2V4h-2V2h-2V0z"/>`
	movieInnerSVG                      = `<path fill="currentColor" d="M3 3h18v18H3V3zm2 2v2h2V5H5zm4 0v6h6V5H9zm8 0v2h2V5h-2zm2 4h-2v2h2V9zm0 4h-2v2h2v-2zm0 4h-2v2h2v-2zm-4 2v-6H9v6h6zm-8 0v-2H5v2h2zm-2-4h2v-2H5v2zm0-4h2V9H5v2z"/>`
	musicInnerSVG                      = `<path fill="currentColor" d="M8 4h12v16h-8v-8h6V8h-8v12H2v-8h6V4zm0 10H4v4h4v-4zm10 0h-4v4h4v-4z"/>`
	nextInnerSVG                       = `<path fill="currentColor" d="M6 4h2v2h2v2h2v2h2v4h-2v2h-2v2H8v2H6V4zm12 0h-2v16h2V4z"/>`
	noteInnerSVG                       = `<path fill="currentColor" d="M3 2h18v14h-2v2h-2v-2h-2v2h2v2h-2v2H3V2zm2 2v16h8v-6h6V4H5z"/>`
	noteDeleteInnerSVG                 = `<path fill="currentColor" d="M11 2h10v14h-2v2h-2v-2h-2v2h2v2h-2v2H3V10h2v10h8v-6h6V4h-8V2zM7 4H5V2H3v2h2v2H3v2h2V6h2v2h2V6H7V4zm0 0h2V2H7v2z"/>`
	noteMultipleInnerSVG               = `<path fill="currentColor" d="M21 6H7v16h8v-2h2v-2h-2v-2h2v2h2v-2h2V6zM9 20V8h10v6h-6v6H9zm-6-2h2V4h12V2H3v16z"/>`
	notePlusInnerSVG                   = `<path fill="currentColor" d="M7 1H5v3H2v2h3v3h2V6h3V4H7V1zm12 1h-7v2h7v10h-6v6H5v-9H3v11h12v-2h2v-2h2v-2h2V2h-2zm-2 16h-2v-2h2v2z"/>`
	notesInnerSVG                      = `<path fill="currentColor" d="M5 2h16v20H3V2h2zm14 18V4H5v16h14zM7 6h10v2H7V6zm10 4H7v2h10v-2zM7 14h7v2H7v-2z"/>`
	notesDeleteInnerSVG                = `<path fill="currentColor" d="M19 2H3v20h10v-2H5V4h14v10h2V2h-2zm-2 4H7v2h10V6zM7 10h10v2H7v-2zm6 4H7v2h6v-2zm6 4h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2zm0 0h2v-2h-2v2z"/>`
	notesMultipleInnerSVG              = `<path fill="currentColor" d="M7 0h16v20H5V0h2zm14 18V2H7v16h14zM9 4h10v2H9V4zm10 4H9v2h10V8zM9 12h7v2H9v-2zm10 10H3V4H1v20h18v-2z"/>`
	notesPlusInnerSVG                  = `<path fill="currentColor" d="M5 2h16v12h-2V4H5v16h8v2H3V2h2zm2 4h10v2H7V6zm10 4H7v2h10v-2zM7 14h7v2H7v-2zm13 5h3v2h-3v3h-2v-3h-3v-2h3v-3h2v3z"/>`
	notificationInnerSVG               = `<path fill="currentColor" d="M14 4V2h-4v2H5v2h14V4h-5zm5 12H5v-4H3v6h5v4h2v-4h4v2h-4v2h6v-4h5v-6h-2V6h-2v8h2v2zM5 6v8h2V6H5z"/>`
	notificationOffInnerSVG            = `<path fill="currentColor" d="M14 2v2h5v2h-8V2h3zM5 16h9v2h2v4h-6v-2h4v-2h-4v4H8v-4H3v-6h2v-2h2v4H5v2zm16-2h-2v-2h-2V6h2v6h2v2zM5 2H3v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2h-2V8H9V6H7V4H5V2z"/>`
	openInnerSVG                       = `<path fill="currentColor" d="M5 3h6v2H5v14h14v-6h2v8H3V3h2zm8 0h8v8h-2V7h-2V5h-4V3zm0 8h-2v2H9v2h2v-2h2v-2zm4-4h-2v2h-2v2h2V9h2V7z"/>`
	paintBucketInnerSVG                = `<path fill="currentColor" d="M8 3h8v2H8V3zm0 2H6v4H4v12h16V9h-2V5h-2v4H8V5zm8 6h2v8H6v-8h2v6h2v-4h2v2h2v-2h2v-2z"/>`
	paperclipInnerSVG                  = `<path fill="currentColor" d="M5 5h16v10H7V9h10v2H9v2h10V7H5v10h14v2H3V5h2z"/>`
	pauseInnerSVG                      = `<path fill="currentColor" d="M10 4H5v16h5V4zm9 0h-5v16h5V4z"/>`
	percentInnerSVG                    = `<path fill="currentColor" d="M20 4h-2v2h-2v2h-2v2h-2v2h-2v2H8v2H6v2H4v2h2v-2h2v-2h2v-2h2v-2h2v-2h2V8h2V6h2V4zm-4 10h4v6h-6v-6h2zm2 4v-2h-2v2h2zM6 4h4v6H4V4h2zm2 4V6H6v2h2z"/>`
	pictureInPictureInnerSVG           = `<path fill="currentColor" d="M2 4h20v16H2V4zm2 2v12h16V6H4zm6 2h8v6h-8V8zm2 2v2h4v-2h-4z"/>`
	pictureInPictureAltInnerSVG        = `<path fill="currentColor" d="M2 4h20v16H2V4zm2 2v12h16V6H4zm6 4h8v6h-8v-6zm2 2v2h4v-2h-4z"/>`
	pinInnerSVG                        = `<path fill="currentColor" d="M7 2h10v2H7V2zM5 6V4h2v2H5zm0 8H3V6h2v8zm2 2H5v-2h2v2zm2 2H7v-2h2v2zm2 2H9v-2h2v2zm2 0v2h-2v-2h2zm2-2v2h-2v-2h2zm2-2v2h-2v-2h2zm2-2v2h-2v-2h2zm0-8h2v8h-2V6zm0 0V4h-2v2h2zm-5 2h-4v4h4V8z"/>`
	pixelarticonsInnerSVG              = `<path fill="currentColor" d="M3 3v18h18V3H3zm16 2v14H5V5h14zM7 7h6v6H9v2H7V7zm8 6h-2v2h-2v2h2v-2h2v2h2v-2h-2v-2zm0 0h2v-2h-2v2zM9 9v2h2V9H9z"/>`
	playInnerSVG                       = `<path fill="currentColor" d="M10 20H8V4h2v2h2v3h2v2h2v2h-2v2h-2v3h-2v2z"/>`
	playlistInnerSVG                   = `<path fill="currentColor" d="M10 13h6V5h6v4h-4v10h-8v-6zm2 2v2h4v-2h-4zM2 17h6v2H2v-2zm6-4H2v2h6v-2zM2 9h12v2H2V9zm12-4H2v2h12V5z"/>`
	plusInnerSVG                       = `<path fill="currentColor" d="M11 4h2v7h7v2h-7v7h-2v-7H4v-2h7V4z"/>`
	powerInnerSVG                      = `<path fill="currentColor" d="M20 2h-2v4H6v2H4v8h2v2h2v4h8v-2h4v-2h-4v-2h4v-2h-4v-2H8v4H6V8h12V6h2V2zm-6 18h-4v-6h4v6z"/>`
	prevInnerSVG                       = `<path fill="currentColor" d="M6 4h2v16H6V4zm12 0h-2v2h-2v3h-2v2h-2v2h2v3h2v2h2v2h2V4z"/>`
	printInnerSVG                      = `<path fill="currentColor" d="M6 2h12v6h4v10h-4v4H6v-4H2V8h4V2zm2 6h8V4H8v4zm-2 8v-4h12v4h2v-6H4v6h2zm2-2v6h8v-6H8z"/>`
	radioHandheldInnerSVG              = `<path fill="currentColor" d="M9 2v5h8v15H7V2h2zm0 7v4h6V9H9zm6 6H9v5h6v-5z"/>`
	radioOnInnerSVG                    = `<path fill="currentColor" d="M17 3H7v2H5v2H3v10h2v2h2v2h10v-2h2v-2h2V7h-2V5h-2V3zm0 2v2h2v10h-2v2H7v-2H5V7h2V5h10zm-9 6h2v2h2v2h-2v-2H8v-2zm8-2h-2v2h-2v2h2v-2h2V9z"/>`
	radioSignalInnerSVG                = `<path fill="currentColor" d="M19 2h2v2h-2V2Zm2 14V4h2v12h-2Zm0 0v2h-2v-2h2ZM1 4h2v12H1V4Zm2 12h2v2H3v-2ZM3 4h2V2H3v2Zm2 2h2v8H5V6Zm2 8h2v2H7v-2Zm0-8h2V4H7v2Zm10 0h2v8h-2V6Zm0 0h-2V4h2v2Zm0 8v2h-2v-2h2Zm-6-7h4v6h-2v9h-2v-9H9V7h2Zm0 4h2V9h-2v2Z"/>`
	radioTowerInnerSVG                 = `<path fill="currentColor" d="M22 2h-2v2h2v12h-2v2h2v-2h2V4h-2V2ZM2 4H0v12h2v2h2v-2H2V4Zm0 0V2h2v2H2Zm4 2H4v8h2V6Zm0 0V4h2v2H6Zm4 0h4v2h-4V6Zm0 6H8V8h2v4Zm4 0h-4v2H8v4H6v4h2v-4h2v-4h4v4h2v4h2v-4h-2v-4h-2v-2Zm0 0h2V8h-2v4Zm6-6h-2V4h-2v2h2v8h2V6Z"/>`
	recieptInnerSVG                    = `<path fill="currentColor" d="M3 2h2v2h2v2H5v14h14V6h-2V4h2V2h2v20H3V2zm12 2V2h2v2h-2zm-2 0h2v2h-2V4zm-2 0V2h2v2h-2zM9 4h2v2H9V4zm0 0V2H7v2h2zm8 4H7v2h10V8zM7 12h10v2H7v-2zm10 6v-2h-4v2h4z"/>`
	recieptAltInnerSVG                 = `<path fill="currentColor" d="M5 2H3v20h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2V2h-2v2h-2V2h-2v2h-2V2h-2v2H9V2H7v2H5V2zm2 2h2v2h2V4h2v2h2V4h2v2h2v12h-2v2h-2v-2h-2v2h-2v-2H9v2H7v-2H5V6h2V4zm0 4h10v2H7V8zm10 4H7v2h10v-2z"/>`
	redoInnerSVG                       = `<path fill="currentColor" d="M16 4h-2v2h2v2H6v2H4v8h2v2h6v-2H6v-8h10v2h-2v2h2v-2h2v-2h2V8h-2V6h-2V4z"/>`
	reloadInnerSVG                     = `<path fill="currentColor" d="M16 2h-2v2h2v2H4v2H2v5h2V8h12v2h-2v2h2v-2h2V8h2V6h-2V4h-2V2zM6 20h2v2h2v-2H8v-2h12v-2h2v-5h-2v5H8v-2h2v-2H8v2H6v2H4v2h2v2z"/>`
	removeBoxInnerSVG                  = `<path fill="currentColor" d="M5 3H3v18h18V3H5zm14 2v14H5V5h14zm-3 6H8v2h8v-2z"/>`
	removeBoxMultipleInnerSVG          = `<path fill="currentColor" d="M5 3H3v14h14V3H5zm10 2v10H5V5h10zm4 2v12H7v2h14V7h-2zm-6 2H7v2h6V9z"/>`
	repeatInnerSVG                     = `<path fill="currentColor" d="M11 1H9v2h2v2H5v2H3v10h2v2h2v-2H5V7h6v2H9v2h2V9h2V7h2V5h-2V3h-2V1zm8 4h-2v2h2v10h-6v-2h2v-2h-2v2h-2v2H9v2h2v2h2v2h2v-2h-2v-2h6v-2h2V7h-2V5z"/>`
	replyInnerSVG                      = `<path fill="currentColor" d="M12 19h-2v-2H8v-2H6v-2H4v-2h2V9h2V7h2V5h2v4h8v6h-8v4z"/>`
	replyAllInnerSVG                   = `<path fill="currentColor" d="M13 19h2v-4h7V9h-7V5h-2v2h-2v2H9v2H7v2h2v2h2v2h2v2zM8 7H6v2H4v2H2v2h2v2h2v2h2v2h2v-2H8v-2H6v-2H4v-2h2V9h2V7zm0 0h2V5H8v2z"/>`
	roundedCornerInnerSVG              = `<path fill="currentColor" d="M3 3h2v2H3V3zm0 4h2v2H3V7zm2 4H3v2h2v-2zm-2 4h2v2H3v-2zm2 4H3v2h2v-2zm2 0h2v2H7v-2zm6 0h-2v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2zm-2-4h2v2h-2v-2zM17 5h-2V3h-4v2h4v2h2v2h2v4h2V9h-2V7h-2V5zM7 3h2v2H7V3z"/>`
	saveInnerSVG                       = `<path fill="currentColor" d="M4 2h14v2H4v16h2v-6h12v6h2V6h2v16H2V2h2zm4 18h8v-4H8v4zM20 6h-2V4h2v2zM6 6h9v4H6V6z"/>`
	scaleInnerSVG                      = `<path fill="currentColor" d="M21 3h-8v2h4v2h2v4h2V3zm-4 4h-2v2h-2v2h2V9h2V7zm-8 8h2v-2H9v2H7v2h2v-2zm-4-2v4h2v2H5h6v2H3v-8h2z"/>`
	scriptInnerSVG                     = `<path fill="currentColor" d="M6 3h14v2h2v6h-2v8h-2V5H6V3zm8 14v-2H6V5H4v10H2v4h2v2h14v-2h-2v-2h-2zm0 0v2H4v-2h10z"/>`
	scriptTextInnerSVG                 = `<path fill="currentColor" d="M6 3h14v2h2v6h-2v8h-2V5H6V3zm8 14v-2H6V5H4v10H2v4h2v2h14v-2h-2v-2h-2zm0 0v2H4v-2h10zM8 7h8v2H8V7zm8 4H8v2h8v-2z"/>`
	scrollHorizontalInnerSVG           = `<path fill="currentColor" d="M22 2v2H2V2h20zm0 18v2H2v-2h20zm-6-5v-2H8v2H6v-2H4v-2h2V9h2v2h8V9h2v2h2v2h-2v2h-2zm0 0v2h-2v-2h2zm0-6h-2V7h2v2zM8 9V7h2v2H8zm0 6h2v2H8v-2z"/>`
	scrollVerticalInnerSVG             = `<path fill="currentColor" d="M2 2h2v20H2V2zm9 18h2v-2h2v-2h2v-2h-2v2h-2V8h2v2h2V8h-2V6h-2V4h-2v2H9v2H7v2h2V8h2v8H9v-2H7v2h2v2h2v2zM22 2h-2v20h2V2z"/>`
	sdInnerSVG                         = `<path fill="currentColor" d="M18 2h2v20H4V6h2v14h12V4H8V2h10zM8 4H6v2h2V4zm6 2h2v4h-2V6zm-2 0h-2v4h2V6z"/>`
	searchInnerSVG                     = `<path fill="currentColor" d="M6 2h8v2H6V2zM4 6V4h2v2H4zm0 8H2V6h2v8zm2 2H4v-2h2v2zm8 0v2H6v-2h8zm2-2h-2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2zm0-8h2v8h-2V6zm0 0V4h-2v2h2z"/>`
	sectionInnerSVG                    = `<path fill="currentColor" d="M5 3H3v2h2V3zm4 0H7v2h2V3zM7 19h2v2H7v-2zM5 7H3v2h2V7zm14 0h2v2h-2V7zM5 11H3v2h2v-2zm14 0h2v2h-2v-2zM5 15H3v2h2v-2zm14 0h2v2h-2v-2zM5 19H3v2h2v-2zm6-16h2v2h-2V3zm2 16h-2v2h2v-2zm2-16h2v2h-2V3zm2 16h-2v2h2v-2zm2-16h2v2h-2V3zm2 16h-2v2h2v-2z"/>`
	sectionCopyInnerSVG                = `<path fill="currentColor" d="M5 3H3v2h2V3zm2 4h2v2H7V7zm4 0h2v2h-2V7zm2 12h-2v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2zM7 11h2v2H7v-2zm14 0h-2v2h2v-2zm-2 4h2v2h-2v-2zM7 19h2v2H7v-2zM19 7h2v2h-2V7zM7 3h2v2H7V3zm2 12H7v2h2v-2zM3 7h2v2H3V7zm14 0h-2v2h2V7zM3 11h2v2H3v-2zm2 4H3v2h2v-2zm6-12h2v2h-2V3zm6 0h-2v2h2V3z"/>`
	sectionMinusInnerSVG               = `<path fill="currentColor" d="M5 3H3v2h2V3zm4 0H7v2h2V3zM7 19h2v2H7v-2zm6 0h-2v2h2v-2zM3 7h2v2H3V7zm18 0h-2v2h2V7zm-2 4h2v2h-2v-2zM5 11H3v2h2v-2zm-2 4h2v2H3v-2zm2 4H3v2h2v-2zm6-16h2v2h-2V3zm6 0h-2v2h2V3zm2 0h2v2h-2V3zm2 14h-6v2h6v-2z"/>`
	sectionPlusInnerSVG                = `<path fill="currentColor" d="M3 3h2v2H3V3zm4 0h2v2H7V3zm2 16H7v2h2v-2zm2 0h2v2h-2v-2zM5 7H3v2h2V7zm14 0h2v2h-2V7zm2 4h-2v2h2v-2zM3 11h2v2H3v-2zm2 4H3v2h2v-2zm12 0h2v2h2v2h-2v2h-2v-2h-2v-2h2v-2zM5 19H3v2h2v-2zm6-16h2v2h-2V3zm6 0h-2v2h2V3zm4 0h-2v2h2V3z"/>`
	sectionXInnerSVG                   = `<path fill="currentColor" d="M5 3H3v2h2V3zm4 0H7v2h2V3zM7 19h2v2H7v-2zm6 0h-2v2h2v-2zM3 7h2v2H3V7zm18 0h-2v2h2V7zm-2 4h2v2h-2v-2zm2 8h-2v-2h2v-2h-2v2h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2zM3 11h2v2H3v-2zm2 4H3v2h2v-2zm-2 4h2v2H3v-2zM13 3h-2v2h2V3zm2 0h2v2h-2V3zm6 0h-2v2h2V3z"/>`
	serverInnerSVG                     = `<path fill="currentColor" d="M3 3h18v18H3V3zm2 2v6h14V5H5zm14 8H5v6h14v-6zM7 7h2v2H7V7zm2 8H7v2h2v-2z"/>`
	sharpCornerInnerSVG                = `<path fill="currentColor" d="M3 3h2v2H3V3zm0 4h2v2H3V7zm2 4H3v2h2v-2zm-2 4h2v2H3v-2zm2 4H3v2h2v-2zm2 0h2v2H7v-2zm6 0h-2v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2zm-2-4h2v2h-2v-2zm2-2V3H11v2h8v8h2zM7 3h2v2H7V3z"/>`
	shieldInnerSVG                     = `<path fill="currentColor" d="M22 2H2v12h2V4h16v10h2V2zM6 14H4v2h2v-2zm0 2h2v2h2v2H8v-2H6v-2zm4 4v2h4v-2h2v-2h-2v2h-4zm10-6h-2v2h-2v2h2v-2h2v-2z"/>`
	shieldOffInnerSVG                  = `<path fill="currentColor" d="M8 2h14v12h-2V4H8V2zM2 8h2v6H2V8zm2 6h2v2H4v-2zm4 2H6v2h2v2h2v2h4v-2h-4v-2H8v-2zm10 0h-2v2h2v2h2v2h2v-2h-2v-2h-2v-2zM4 2H2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2V8H8V6H6V4H4V2z"/>`
	shipInnerSVG                       = `<path fill="currentColor" d="M8 4v2h4v2H6v2h6V8h2v2h8v6h-2v-4H4v6h14v-2h2v2h4v2H0v-2h2v-8h2V6h2V4h2z"/>`
	shoppingBagInnerSVG                = `<path fill="currentColor" d="M9 2h6v2H9V2zm6 4V4h2v2h4v16H3V6h4V4h2v2h6zm0 2H9v2H7V8H5v12h14V8h-2v2h-2V8z"/>`
	shuffleInnerSVG                    = `<path fill="currentColor" d="M18 5h-2v2h2v2h-6v2h-2v6H2v2h8v-2h2v-6h6v2h-2v2h2v-2h2v-2h2V9h-2V7h-2V5zM2 9h6v2H2V9zm20 10v-2h-8v2h8z"/>`
	slidersInnerSVG                    = `<path fill="currentColor" d="M17 4h2v10h-2V4zm0 12h-2v2h2v2h2v-2h2v-2h-4zm-4-6h-2v10h2V10zm-8 2H3v2h2v6h2v-6h2v-2H5zm8-8h-2v2H9v2h6V6h-2V4zM5 4h2v6H5V4z"/>`
	slidersTwoInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-width="2" d="M3 8h4m0 0V6h4v2M7 8v2h4V8m0 0h10M3 16h10m0 0v-2h4v2m-4 0v2h4v-2m0 0h4"/>`
	sortInnerSVG                       = `<path fill="currentColor" d="M8 20H6V8H4V6h2V4h2v2h2v2H8v12zm2-12v2h2V8h-2zM4 8v2H2V8h2zm14-4h-2v12h-2v-2h-2v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2V4z"/>`
	sortAlpabeticInnerSVG              = `<path fill="currentColor" d="M11 2h2v2h-2V2zm0 2v2H9V4h2zm2 0h2v2h-2V4zM9 18v2h2v2h2v-2h2v-2h-2v2h-2v-2H9zM8 8H2v8h2v-2h2v2h2V8zm-2 4H4v-2h2v2zm6-1v-1h2v1h-2zm4-3h-6v8h6V8zm-4 6v-1h2v1h-2zm10-6h-4v8h4v-2h-2v-4h2V8z"/>`
	sortNumericInnerSVG                = `<path fill="currentColor" d="M13 2h-2v2H9v2h2V4h2v2h2V4h-2V2zM2 8h4v8H4v-6H2V8zm6 0h6v5h-4v1h4v2H8v-5h4v-1H8V8zm12 0h-4v2h4v1h-4v2h4v1h-4v2h6V8h-2zm-9 10v2H9v-2h2zm2 2h-2v2h2v-2zm0 0v-2h2v2h-2z"/>`
	speakerInnerSVG                    = `<path fill="currentColor" d="M4 2H3v20h18V2H4zm15 2v16H5V4h14zm-6 2h-2v2h2V6zm-5 4h8v6h-2v-4h-4v4H8v-6zm8 6H8v2h8v-2z"/>`
	speedFastInnerSVG                  = `<path fill="currentColor" d="M15 5H9v2H5v2H3v2H1v6h2v2h2v-2H3v-6h2V9h4V7h6V5zm8 6h-2v6h-2v2h2v-2h2v-6zm-13 2h4v4h-4v-4zm6-2h-2v2h2v-2zm2-2v2h-2V9h2zm0 0V7h2v2h-2z"/>`
	speedMediumInnerSVG                = `<path fill="currentColor" d="M13 5h-2v8h-1v4h4v-4h-1V5zM9 7H5v2H3v2H1v6h2v2h2v-2H3v-6h2V9h4V7zm12 4h2v6h-2v-6zm-2-2h2v2h-2V9zm0 0h-4V7h4v2zm2 8v2h-2v-2h2z"/>`
	speedSlowInnerSVG                  = `<path fill="currentColor" d="M9 5h6v2H9V5zm10 4h-4V7h4v2zm2 2h-2V9h2v2zm0 6v-6h2v6h-2zm0 0v2h-2v-2h2zM1 11h2v6H1v-6zm2 6h2v2H3v-2zm11-4h-4v-2H8V9H6V7H4v2h2v2h2v2h2v4h4v-4z"/>`
	spotlightInnerSVG                  = `<path fill="currentColor" d="M5 2h16v20H3V2h2zm14 18V4H5v16h14zM13 6H7v2h6V6zm-6 4h10v8H7v-8z"/>`
	storeInnerSVG                      = `<path fill="currentColor" d="M4 3h16v2H4V3zm0 4h18v8h-2v6h-2v-6h-4v6H4v-6H2V7h2zm8 12v-4H6v4h6zm0-6h8V9H4v4h8z"/>`
	subscriptionsInnerSVG              = `<path fill="currentColor" d="M18 2H6v2h12V2zM4 6h16v2H4V6zm-2 4h20v12H2V10zm18 10v-8H4v8h16z"/>`
	subtitlesInnerSVG                  = `<path fill="currentColor" d="M21 7h-8v10h8v-2h-6V9h6V7zM3 15V7h8v2H5v6h6v2H3v-2z"/>`
	suitcaseInnerSVG                   = `<path fill="currentColor" d="M8 3h8v4h6v14H2V7h6V3zm2 4h4V5h-4v2zM4 9v10h16V9H4zm4 2v6H6v-6h2zm10 0v6h-2v-6h2z"/>`
	sunInnerSVG                        = `<path fill="currentColor" d="M13 3h-2v2h2V3zm4 2h2v2h-2V5zm-6 6h2v2h-2v-2zm-8 0h2v2H3v-2zm18 0h-2v2h2v-2zM5 5h2v2H5V5zm14 14h-2v-2h2v2zm-8 2h2v-2h-2v2zm-4-2H5v-2h2v2zM9 7h6v2H9V7zm0 8H7V9h2v6zm0 0v2h6v-2h2V9h-2v6H9z"/>`
	sunAltInnerSVG                     = `<path fill="currentColor" d="M13 0h-2v4h2V0ZM0 11v2h4v-2H0Zm24 0v2h-4v-2h4ZM13 24h-2v-4h2v4ZM8 6h8v2H8V6ZM6 8h2v8H6V8Zm2 10v-2h8v2H8Zm10-2h-2V8h2v8Zm2-14h2v2h-2V2Zm0 2v2h-2V4h2Zm2 18h-2v-2h2v2Zm-2-2h-2v-2h2v2ZM4 2H2v2h2v2h2V4H4V2ZM2 22h2v-2h2v-2H4v2H2v2Z"/>`
	switchInnerSVG                     = `<path fill="currentColor" d="M3 5V3h2v2H3zm4 2H5V5h2v2zm2 2H7V7h2v2zm2 2H9V9h2v2zm2 0h-2v2h2v2h2v2h2v2h-2v2h6v-6h-2v2h-2v-2h-2v-2h-2v-2zm2-2v2h-2V9h2zm2-2v2h-2V7h2zm0-2v2h2v2h2V3h-6v2h2zM5 19v-2h2v2H5zm0 0v2H3v-2h2zm2-2v-2h2v2H7z"/>`
	syncInnerSVG                       = `<path fill="currentColor" d="M4 9V7h12V5h2v2h2v2h-2v2h-2V9H4zm12 2h-2v2h2v-2zm0-6h-2V3h2v2zm4 12v-2H8v-2h2v-2H8v2H6v2H4v2h2v2h2v2h2v-2H8v-2h12z"/>`
	tabInnerSVG                        = `<path fill="currentColor" d="M2 3h20v18H2V3zm2 2v14h16V9h-8V5H4z"/>`
	tableInnerSVG                      = `<path fill="currentColor" d="M2 3h20v18H2V3zm2 4v5h7V7H4zm9 0v5h7V7h-7zm7 7h-7v5h7v-5zm-9 5v-5H4v5h7z"/>`
	teaInnerSVG                        = `<path fill="currentColor" d="M4 4h18v7h-4v5H4V4zm14 5h2V6h-2v3zm-2-3h-4v2h2v4H8V8h2V6H6v8h10V6zm3 12v2H3v-2h16z"/>`
	teachInnerSVG                      = `<path fill="currentColor" d="M9 2H5v4h4V2zm7 7V7H2v9h2v6h2v-6h2v6h2V9h6zm-5-7h11v14H11v-2h9V4h-9V2z"/>`
	textAddInnerSVG                    = `<path fill="currentColor" d="M19 4H3v2h16V4zm0 4H3v2h16V8zM3 12h8v2H3v-2zm8 4H3v2h8v-2zm7-1h3v2h-3v3h-2v-3h-3v-2h3v-3h2v3z"/>`
	textColumsInnerSVG                 = `<path fill="currentColor" d="M11 5H3v2h8V5zm10 0h-8v2h8V5zM3 9h8v2H3V9zm18 0h-8v2h8V9zM3 13h8v2H3v-2zm18 0h-8v2h8v-2zM3 17h8v2H3v-2zm18 0h-8v2h8v-2z"/>`
	textSearchInnerSVG                 = `<path fill="currentColor" d="M20 4H4v2h16V4zm0 4H4v2h16V8zm-8 4H4v2h8v-2zm8 0h-6v6h6v2h2v-2h-2v-6zm-4 4v-2h2v2h-2zm-4 0H4v2h8v-2z"/>`
	textWrapInnerSVG                   = `<path fill="currentColor" d="M19 5H3v2h16v6h-6v-2h2V9h-2v2h-2v2H9v2h2v2h2v2h2v-2h-2v-2h6v-2h2V7h-2V5zM7 13H3v2h4v-2zM3 9h6v2H3V9z"/>`
	timelineInnerSVG                   = `<path fill="currentColor" d="M7 7h4v4H7V7zm-2 6v-2h2v2H5zm0 0v4H1v-4h4zm8 0h-2v-2h2v2zm4 0h-4v4h4v-4zm2-2v2h-2v-2h2zm0 0h4V7h-4v4z"/>`
	toggleLeftInnerSVG                 = `<path fill="currentColor" d="M4 5h16v2H4V5zm0 12H2V7h2v10zm16 0v2H4v-2h16zm0 0h2V7h-2v10zM10 9H6v6h4V9z"/>`
	toggleRightInnerSVG                = `<path fill="currentColor" d="M4 5h16v2H4V5zm0 12H2V7h2v10zm16 0v2H4v-2h16zm0 0h2V7h-2v10zm-2-8h-4v6h4V9z"/>`
	tournamentInnerSVG                 = `<path fill="currentColor" d="M9 2H2v2h5v4H2v2h7V7h5v10H9v-3H2v2h5v4H2v2h7v-3h7v-6h6v-2h-6V5H9V2z"/>`
	trackChangesInnerSVG               = `<path fill="currentColor" d="M11 2H2v20h20V4h-2v16H4V4h7v2H6v12h12V8h-2v8H8V8h3v2h-1v4h4v-4h-1V2h-2z"/>`
	trashInnerSVG                      = `<path fill="currentColor" d="M16 2v4h6v2h-2v14H4V8H2V6h6V2h8zm-2 2h-4v2h4V4zm0 4H6v12h12V8h-4zm-5 2h2v8H9v-8zm6 0h-2v8h2v-8z"/>`
	trashAltInnerSVG                   = `<path fill="currentColor" d="M16 2v4h6v2h-2v14H4V8H2V6h6V2h8zm-2 2h-4v2h4V4zm0 4H6v12h12V8h-4z"/>`
	trendingInnerSVG                   = `<path fill="currentColor" d="M3 4h2v14h16v2H3V4zm6 10H7v2h2v-2zm2-2v2H9v-2h2zm2 0v-2h-2v2h2zm2 0h-2v2h2v-2zm2-2h-2v2h2v-2zm2-2v2h-2V8h2zm0 0V6h2v2h-2z"/>`
	trendingDownInnerSVG               = `<path fill="currentColor" d="M2 8h2v2h2v2h2v2h2v-2h2v-2h2v2h2v2h2v2h-4v2h8v-8h-2v4h-2v-2h-2v-2h-2V8h-2v2h-2v2H8v-2H6V8H4V6H2v2z"/>`
	trendingUpInnerSVG                 = `<path fill="currentColor" d="M14 6h8v8h-2v-4h-2V8h-4V6zm2 6v-2h2v2h-2zm-2 2v-2h2v2h-2zm-2 0h2v2h-2v-2zm-2-2h2v2h-2v-2zm-2 0v-2h2v2H8zm-2 2v-2h2v2H6zm-2 2v-2h2v2H4zm0 0v2H2v-2h2z"/>`
	trophyInnerSVG                     = `<path fill="currentColor" d="M16 3H6v2H2v10h6V5h8v10h6V5h-4V3h-2zm4 4v6h-2V7h2zM6 13H4V7h2v6zm12 2H6v2h12v-2zm-7 2h2v2h3v2H8v-2h3v-2z"/>`
	truckInnerSVG                      = `<path fill="currentColor" d="M2 4h14v4h4v2h-4v6h6v-4h2v6h-4v2h-4v-2H8v2H4v-2H0V4h2zm20 8h-2v-2h2v2zm-8-2V6H2v10h12v-6z"/>`
	undoInnerSVG                       = `<path fill="currentColor" d="M8 4h2v2H8V4zm10 6V8H8V6H6v2H4v2h2v2h2v2h2v-2H8v-2h10zm0 8v-8h2v8h-2zm0 0v2h-6v-2h6z"/>`
	ungroupInnerSVG                    = `<path fill="currentColor" d="M7 3H3v4h4V3zm0 14H3v4h4v-4zM17 3h4v4h-4V3zm4 14h-4v4h4v-4zM8 8h2v2H8V8zm4 2h-2v4H8v2h2v-2h4v2h2v-2h-2v-4h2V8h-2v2h-2z"/>`
	unlinkInnerSVG                     = `<path fill="currentColor" d="M13 4h-2v16h2V4zM4 6h5v2H4v8h5v2H2V6h2zm11 0h7v12h-7v-2h5V8h-5V6z"/>`
	uploadInnerSVG                     = `<path fill="currentColor" d="M11 5V3h2v2h2v2h2v2h-2V7h-2v10h-2V7H9v2H7V7h2V5h2zM3 15v6h18v-6h-2v4H5v-4H3z"/>`
	userInnerSVG                       = `<path fill="currentColor" d="M15 2H9v2H7v6h2V4h6V2zm0 8H9v2h6v-2zm0-6h2v6h-2V4zM4 16h2v-2h12v2H6v4h12v-4h2v6H4v-6z"/>`
	userMinusInnerSVG                  = `<path fill="currentColor" d="M12 2h6v2h-6v6h-2V4h2V2zm0 8h6v2h-6v-2zm8-6h-2v6h2V4zM9 16H7v6h16v-6h-2v4H9v-4h12v-2H9v2zm-2-6H1v2h6v-2z"/>`
	userPlusInnerSVG                   = `<path fill="currentColor" d="M18 2h-6v2h-2v6h2V4h6V2zm0 8h-6v2h6v-2zm0-6h2v6h-2V4zM7 16h2v-2h12v2H9v4h12v-4h2v6H7v-6zM3 8h2v2h2v2H5v2H3v-2H1v-2h2V8z"/>`
	userXInnerSVG                      = `<path fill="currentColor" d="M12 2h6v2h-6v6h-2V4h2V2zm0 8h6v2h-6v-2zm8-6h-2v6h2V4zM7 16v6h16v-6h-2v4H9v-4h12v-2H9v2H7zm-1-6H4V8H2v2h2v2H2v2h2v-2h2v2h2v-2H6v-2zm0 0h2V8H6v2z"/>`
	usersInnerSVG                      = `<path fill="currentColor" d="M11 0H5v2H3v6h2v2h6V8H5V2h6V0zm0 2h2v6h-2V2zM0 14h2v4h12v2H0v-6zm2 0h12v-2H2v2zm14 0h-2v6h2v-6zM15 0h4v2h-4V0zm4 8h-4v2h4V8zm0-6h2v6h-2V2zm5 12h-2v4h-4v2h6v-6zm-6-2h4v2h-4v-2z"/>`
	videoInnerSVG                      = `<path fill="currentColor" d="M2 5h14v4h2V7h2V5h2v14h-2v-2h-2v-2h-2v4H2V5zm2 12h10V7H4v10z"/>`
	videoOffInnerSVG                   = `<path fill="currentColor" d="M4 5H2v14h14v-4h2v2h2v2h2V5h-2v2h-2v2h-2V5H4zm10 12H4V7h10v10zm-4-6H8V9H6v2h2v2H6v2h2v-2h2v2h2v-2h-2v-2zm0 0V9h2v2h-2z"/>`
	viewColInnerSVG                    = `<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v10h4V7H4zm6 0v10h4V7h-4zm6 0v10h4V7h-4z"/>`
	viewListInnerSVG                   = `<path fill="currentColor" d="M2 5h20v14H2V5zm2 2v2h16V7H4zm16 4H4v2h16v-2zm0 4H4v2h16v-2z"/>`
	viewportNarrowInnerSVG             = `<path fill="currentColor" d="M10 2H8v4h2V4h4v2h2V2h-6zM8 20v-2h2v2h4v-2h2v4H8v-2zm9-9h5v2h-5v2h-2v-2h-2v-2h2V9h2v2zm0-2V7h2v2h-2zm0 6h2v2h-2v-2zM2 11h5V9h2v2h2v2H9v2H7v-2H2v-2zm5 4v2H5v-2h2zm0-6V7H5v2h2z"/>`
	viewportWideInnerSVG               = `<path fill="currentColor" d="M4 2H2v4h2V4h16v2h2V2H4zM2 20v-2h2v2h16v-2h2v4H2v-2zm16-9h-5v2h5v2h-2v2h2v-2h2v-2h2v-2h-2V9h-2V7h-2v2h2v2zm-7 0H6V9h2V7H6v2H4v2H2v2h2v2h2v2h2v-2H6v-2h5v-2z"/>`
	visibleInnerSVG                    = `<path fill="currentColor" d="M8 6h8v2H8V6zm-4 4V8h4v2H4zm-2 2v-2h2v2H2zm0 2v-2H0v2h2zm2 2H2v-2h2v2zm4 2H4v-2h4v2zm8 0v2H8v-2h8zm4-2v2h-4v-2h4zm2-2v2h-2v-2h2zm0-2h2v2h-2v-2zm-2-2h2v2h-2v-2zm0 0V8h-4v2h4zm-10 1h4v4h-4v-4z"/>`
	volumeInnerSVG                     = `<path fill="currentColor" d="M15 2h2v20h-2v-2h-2v-2h2V6h-2V4h2V2zm-4 6V6h2v2h-2zm-2 2h2V8H7v8h4v2h2v-2h-2v-2H9v-4z"/>`
	volumeMinusInnerSVG                = `<path fill="currentColor" d="M12 2h-2v2H8v2H6v2H2v8h4v2h2v2h2v2h2V2zM8 18v-2H6v-2H4v-4h2V8h2V6h2v12H8zm14-7h-8v2h8v-2z"/>`
	volumeOneInnerSVG                  = `<path fill="currentColor" d="M15 2h-2v2h-2v2H9v2H5v8h4v2h2v2h2v2h2V2zm-4 16v-2H9v-2H7v-4h2V8h2V6h2v12h-2zm6-8h2v4h-2v-4z"/>`
	volumePlusInnerSVG                 = `<path fill="currentColor" d="M10 2h2v20h-2v-2H8v-2h2V6H8V4h2V2zM6 8V6h2v2H6zm0 8H2V8h4v2H4v4h2v2zm0 0v2h2v-2H6zm13-5h3v2h-3v3h-2v-3h-3v-2h3V8h2v3z"/>`
	volumeThreeInnerSVG                = `<path fill="currentColor" d="M11 2H9v2H7v2H5v2H1v8h4v2h2v2h2v2h2V2zM7 18v-2H5v-2H3v-4h2V8h2V6h2v12H7zm6-8h2v4h-2v-4zm8-6h-2V2h-6v2h6v2h2v12h-2v2h-6v2h6v-2h2v-2h2V6h-2V4zm-2 4h-2V6h-4v2h4v8h-4v2h4v-2h2V8z"/>`
	volumeTwoInnerSVG                  = `<path fill="currentColor" d="M11 2h2v20h-2v-2H9v-2h2V6H9V4h2V2zM7 8V6h2v2H7zm0 8H3V8h4v2H5v4h2v2zm0 0v2h2v-2H7zm10-6h-2v4h2v-4zm2-2h2v8h-2V8zm0 8v2h-4v-2h4zm0-10v2h-4V6h4z"/>`
	volumeVibrateInnerSVG              = `<path fill="currentColor" d="M14 2h-2v2h-2v2H8v2H4v8h4v2h2v2h2v2h2V2zm-4 16v-2H8v-2H6v-4h2V8h2V6h2v12h-2zm8-15h-2v2h2v2h-2v2h2v2h-2v2h2v2h-2v2h2v2h-2v2h2v-2h2v-2h-2v-2h2v-2h-2v-2h2V9h-2V7h2V5h-2V3z"/>`
	volumeXInnerSVG                    = `<path fill="currentColor" d="M13 2h-2v2H9v2H7v2H3v8h4v2h2v2h2v2h2V2zM9 18v-2H7v-2H5v-4h2V8h2V6h2v12H9zm10-6.777h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2zm0 0h2v-2h-2v2z"/>`
	walletInnerSVG                     = `<path fill="currentColor" d="M18 3H2v18h18v-4h2V7h-2V3h-2zm0 14v2H4V5h14v2h-8v10h8zm2-2h-8V9h8v6zm-4-4h-2v2h2v-2z"/>`
	warningBoxInnerSVG                 = `<path fill="currentColor" d="M3 3h16v2H5v14h14v2H3V3zm18 0h-2v18h2V3zM11 15h2v2h-2v-2zm2-8h-2v6h2V7z"/>`
	windInnerSVG                       = `<path fill="currentColor" d="M12 3H8v2h4v2H2v2h12V3h-2zm10 8V7h-6v2h4v2H2v2h20v-2zM2 17v-2h14v6h-6v-2h4v-2H2z"/>`
	zapInnerSVG                        = `<path fill="currentColor" d="M12 1h2v8h8v4h-2v-2h-8V5h-2V3h2V1zM8 7V5h2v2H8zM6 9V7h2v2H6zm-2 2V9h2v2H4zm10 8v2h-2v2h-2v-8H2v-4h2v2h8v6h2zm2-2v2h-2v-2h2zm2-2v2h-2v-2h2zm0 0h2v-2h-2v2z"/>`
	zoomInInnerSVG                     = `<path fill="currentColor" d="M14 2H6v2H4v2H2v8h2v2h2v2h8v-2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h2V6h-2V4h-2V2zm0 2v2h2v8h-2v2H6v-2H4V6h2V4h8zM9 6h2v3h3v2h-3v3H9v-3H6V9h3V6z"/>`
	zoomOutInnerSVG                    = `<path fill="currentColor" d="M14 2H6v2H4v2H2v8h2v2h2v2h8v-2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h2V6h-2V4h-2V2zm0 2v2h2v8h-2v2H6v-2H4V6h2V4h8zm0 5v2H6V9h8z"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func AbTesting(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(abTestingInnerSVG).
		Element(children...)
}

func Ac(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(acInnerSVG).
		Element(children...)
}

func AddBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addBoxInnerSVG).
		Element(children...)
}

func AddBoxMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addBoxMultipleInnerSVG).
		Element(children...)
}

func AddCol(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addColInnerSVG).
		Element(children...)
}

func AddGrid(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addGridInnerSVG).
		Element(children...)
}

func AddRow(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addRowInnerSVG).
		Element(children...)
}

func Alert(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertInnerSVG).
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

func Animation(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(animationInnerSVG).
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

func ArrowBarDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBarDownInnerSVG).
		Element(children...)
}

func ArrowBarLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBarLeftInnerSVG).
		Element(children...)
}

func ArrowBarRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBarRightInnerSVG).
		Element(children...)
}

func ArrowBarUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBarUpInnerSVG).
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

func ArrowDownBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownBoxInnerSVG).
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

func ArrowLeftBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftBoxInnerSVG).
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

func ArrowRightBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightBoxInnerSVG).
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

func ArrowUpBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpBoxInnerSVG).
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

func ArtText(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(artTextInnerSVG).
		Element(children...)
}

func Article(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(articleInnerSVG).
		Element(children...)
}

func ArticleMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(articleMultipleInnerSVG).
		Element(children...)
}

func AspectRatio(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aspectRatioInnerSVG).
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

func Attachment(children ...engine.UberChild) *engine.UberElement {
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

func AudioDevice(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(audioDeviceInnerSVG).
		Element(children...)
}

func Avatar(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(avatarInnerSVG).
		Element(children...)
}

func Backburger(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backburgerInnerSVG).
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

func BatteryOne(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryOneInnerSVG).
		Element(children...)
}

func BatteryTwo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryTwoInnerSVG).
		Element(children...)
}

func Bed(children ...engine.UberChild) *engine.UberElement {
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

func Bluetooth(children ...engine.UberChild) *engine.UberElement {
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

func Bookmarks(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarksInnerSVG).
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

func BriefcaseAccount(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseAccountInnerSVG).
		Element(children...)
}

func BriefcaseCheck(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseCheckInnerSVG).
		Element(children...)
}

func BriefcaseDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseDeleteInnerSVG).
		Element(children...)
}

func BriefcaseDownload(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseDownloadInnerSVG).
		Element(children...)
}

func BriefcaseMinus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseMinusInnerSVG).
		Element(children...)
}

func BriefcasePlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcasePlusInnerSVG).
		Element(children...)
}

func BriefcaseSearch(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseSearchInnerSVG).
		Element(children...)
}

func BriefcaseSearchOne(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseSearchOneInnerSVG).
		Element(children...)
}

func BriefcaseUpload(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(briefcaseUploadInnerSVG).
		Element(children...)
}

func Bug(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bugInnerSVG).
		Element(children...)
}

func Building(children ...engine.UberChild) *engine.UberElement {
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

func BuildingCommunity(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buildingCommunityInnerSVG).
		Element(children...)
}

func BuildingSkyscraper(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buildingSkyscraperInnerSVG).
		Element(children...)
}

func Buildings(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(buildingsInnerSVG).
		Element(children...)
}

func Bulletlist(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bulletlistInnerSVG).
		Element(children...)
}

func Bullseye(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bullseyeInnerSVG).
		Element(children...)
}

func BullseyeArrow(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bullseyeArrowInnerSVG).
		Element(children...)
}

func Bus(children ...engine.UberChild) *engine.UberElement {
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

func Cake(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cakeInnerSVG).
		Element(children...)
}

func Calculator(children ...engine.UberChild) *engine.UberElement {
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

func CalendarAlert(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarAlertInnerSVG).
		Element(children...)
}

func CalendarArrowLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarArrowLeftInnerSVG).
		Element(children...)
}

func CalendarArrowRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarArrowRightInnerSVG).
		Element(children...)
}

func CalendarCheck(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarCheckInnerSVG).
		Element(children...)
}

func CalendarExport(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarExportInnerSVG).
		Element(children...)
}

func CalendarGrid(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarGridInnerSVG).
		Element(children...)
}

func CalendarImport(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarImportInnerSVG).
		Element(children...)
}

func CalendarMinus(children ...engine.UberChild) *engine.UberElement {
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

func CalendarMonth(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMonthInnerSVG).
		Element(children...)
}

func CalendarMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMultipleInnerSVG).
		Element(children...)
}

func CalendarMultipleCheck(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMultipleCheckInnerSVG).
		Element(children...)
}

func CalendarPlus(children ...engine.UberChild) *engine.UberElement {
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

func CalendarRange(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarRangeInnerSVG).
		Element(children...)
}

func CalendarRemove(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarRemoveInnerSVG).
		Element(children...)
}

func CalendarSearch(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarSearchInnerSVG).
		Element(children...)
}

func CalendarSortAscending(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarSortAscendingInnerSVG).
		Element(children...)
}

func CalendarSortDescending(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarSortDescendingInnerSVG).
		Element(children...)
}

func CalendarText(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarTextInnerSVG).
		Element(children...)
}

func CalendarToday(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarTodayInnerSVG).
		Element(children...)
}

func CalendarTomorrow(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarTomorrowInnerSVG).
		Element(children...)
}

func CalendarWeek(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarWeekInnerSVG).
		Element(children...)
}

func CalendarWeekBegin(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarWeekBeginInnerSVG).
		Element(children...)
}

func CalendarWeekend(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarWeekendInnerSVG).
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

func CameraAdd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraAddInnerSVG).
		Element(children...)
}

func CameraAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraAltInnerSVG).
		Element(children...)
}

func CameraFace(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraFaceInnerSVG).
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

func Card(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardInnerSVG).
		Element(children...)
}

func CardId(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardIdInnerSVG).
		Element(children...)
}

func CardPlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardPlusInnerSVG).
		Element(children...)
}

func CardStack(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardStackInnerSVG).
		Element(children...)
}

func CardText(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardTextInnerSVG).
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

func Cast(children ...engine.UberChild) *engine.UberElement {
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

func CellularSignalOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cellularSignalOffInnerSVG).
		Element(children...)
}

func CellularSignalOne(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cellularSignalOneInnerSVG).
		Element(children...)
}

func CellularSignalThree(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cellularSignalThreeInnerSVG).
		Element(children...)
}

func CellularSignalTwo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cellularSignalTwoInnerSVG).
		Element(children...)
}

func CellularSignalZero(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cellularSignalZeroInnerSVG).
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

func ChartAdd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartAddInnerSVG).
		Element(children...)
}

func ChartBar(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartBarInnerSVG).
		Element(children...)
}

func ChartDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartDeleteInnerSVG).
		Element(children...)
}

func ChartMinus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartMinusInnerSVG).
		Element(children...)
}

func ChartMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartMultipleInnerSVG).
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

func CheckDouble(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkDoubleInnerSVG).
		Element(children...)
}

func Checkbox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxInnerSVG).
		Element(children...)
}

func CheckboxOn(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxOnInnerSVG).
		Element(children...)
}

func Checklist(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checklistInnerSVG).
		Element(children...)
}

func Chess(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chessInnerSVG).
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

func ChevronsHorizontal(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsHorizontalInnerSVG).
		Element(children...)
}

func ChevronsVertical(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsVerticalInnerSVG).
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

func Close(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeInnerSVG).
		Element(children...)
}

func CloseBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeBoxInnerSVG).
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

func CloudDone(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudDoneInnerSVG).
		Element(children...)
}

func CloudDownload(children ...engine.UberChild) *engine.UberElement {
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

func CloudMoon(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudMoonInnerSVG).
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

func CloudUpload(children ...engine.UberChild) *engine.UberElement {
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

func Cocktail(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cocktailInnerSVG).
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

func CoffeeAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coffeeAltInnerSVG).
		Element(children...)
}

func Coin(children ...engine.UberChild) *engine.UberElement {
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

func Collapse(children ...engine.UberChild) *engine.UberElement {
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

func ColorsSwatch(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(colorsSwatchInnerSVG).
		Element(children...)
}

func Command(children ...engine.UberChild) *engine.UberElement {
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

func Contact(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contactInnerSVG).
		Element(children...)
}

func ContactDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contactDeleteInnerSVG).
		Element(children...)
}

func ContactMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contactMultipleInnerSVG).
		Element(children...)
}

func ContactPlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contactPlusInnerSVG).
		Element(children...)
}

func Copy(children ...engine.UberChild) *engine.UberElement {
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

func CreditCardDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardDeleteInnerSVG).
		Element(children...)
}

func CreditCardMinus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardMinusInnerSVG).
		Element(children...)
}

func CreditCardMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardMultipleInnerSVG).
		Element(children...)
}

func CreditCardPlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardPlusInnerSVG).
		Element(children...)
}

func CreditCardSettings(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardSettingsInnerSVG).
		Element(children...)
}

func CreditCardWireless(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardWirelessInnerSVG).
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

func Cut(children ...engine.UberChild) *engine.UberElement {
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

func Dashbaord(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashbaordInnerSVG).
		Element(children...)
}

func Debug(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(debugInnerSVG).
		Element(children...)
}

func DebugCheck(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(debugCheckInnerSVG).
		Element(children...)
}

func DebugOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(debugOffInnerSVG).
		Element(children...)
}

func DebugPause(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(debugPauseInnerSVG).
		Element(children...)
}

func DebugPlay(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(debugPlayInnerSVG).
		Element(children...)
}

func DebugStop(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(debugStopInnerSVG).
		Element(children...)
}

func Delete(children ...engine.UberChild) *engine.UberElement {
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

func Deskphone(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deskphoneInnerSVG).
		Element(children...)
}

func DeviceLaptop(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deviceLaptopInnerSVG).
		Element(children...)
}

func DevicePhone(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(devicePhoneInnerSVG).
		Element(children...)
}

func DeviceTablet(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deviceTabletInnerSVG).
		Element(children...)
}

func DeviceTv(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deviceTvInnerSVG).
		Element(children...)
}

func DeviceTvSmart(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deviceTvSmartInnerSVG).
		Element(children...)
}

func DeviceVibrate(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deviceVibrateInnerSVG).
		Element(children...)
}

func DeviceWatch(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deviceWatchInnerSVG).
		Element(children...)
}

func Devices(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(devicesInnerSVG).
		Element(children...)
}

func Dice(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diceInnerSVG).
		Element(children...)
}

func Dollar(children ...engine.UberChild) *engine.UberElement {
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

func Downasaur(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downasaurInnerSVG).
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

func Draft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(draftInnerSVG).
		Element(children...)
}

func DragAndDrop(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragAndDropInnerSVG).
		Element(children...)
}

func Drop(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropInnerSVG).
		Element(children...)
}

func DropArea(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropAreaInnerSVG).
		Element(children...)
}

func DropFull(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropFullInnerSVG).
		Element(children...)
}

func DropHalf(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropHalfInnerSVG).
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

func DuplicateAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(duplicateAltInnerSVG).
		Element(children...)
}

func Edit(children ...engine.UberChild) *engine.UberElement {
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

func EditBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(editBoxInnerSVG).
		Element(children...)
}

func Euro(children ...engine.UberChild) *engine.UberElement {
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

func EyeClosed(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeClosedInnerSVG).
		Element(children...)
}

func File(children ...engine.UberChild) *engine.UberElement {
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

func FileAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileAltInnerSVG).
		Element(children...)
}

func FileDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileDeleteInnerSVG).
		Element(children...)
}

func FileFlash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileFlashInnerSVG).
		Element(children...)
}

func FileMinus(children ...engine.UberChild) *engine.UberElement {
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

func FileMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileMultipleInnerSVG).
		Element(children...)
}

func FileOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileOffInnerSVG).
		Element(children...)
}

func FilePlus(children ...engine.UberChild) *engine.UberElement {
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

func Fill(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fillInnerSVG).
		Element(children...)
}

func FillHalf(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fillHalfInnerSVG).
		Element(children...)
}

func FiveG(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fiveGInnerSVG).
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

func Flatten(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flattenInnerSVG).
		Element(children...)
}

func FlipToBack(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipToBackInnerSVG).
		Element(children...)
}

func FlipToFront(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flipToFrontInnerSVG).
		Element(children...)
}

func FloatCenter(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatCenterInnerSVG).
		Element(children...)
}

func FloatLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatLeftInnerSVG).
		Element(children...)
}

func FloatRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floatRightInnerSVG).
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

func FolderMinus(children ...engine.UberChild) *engine.UberElement {
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

func FolderPlus(children ...engine.UberChild) *engine.UberElement {
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

func FolderX(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderXInnerSVG).
		Element(children...)
}

func Forward(children ...engine.UberChild) *engine.UberElement {
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

func Forwardburger(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardburgerInnerSVG).
		Element(children...)
}

func FourG(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fourGInnerSVG).
		Element(children...)
}

func FourK(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fourKInnerSVG).
		Element(children...)
}

func FourKBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fourKBoxInnerSVG).
		Element(children...)
}

func Frame(children ...engine.UberChild) *engine.UberElement {
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

func FrameAdd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameAddInnerSVG).
		Element(children...)
}

func FrameCheck(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameCheckInnerSVG).
		Element(children...)
}

func FrameDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameDeleteInnerSVG).
		Element(children...)
}

func FrameMinus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frameMinusInnerSVG).
		Element(children...)
}

func Gamepad(children ...engine.UberChild) *engine.UberElement {
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

func Gif(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gifInnerSVG).
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

func GitBranch(children ...engine.UberChild) *engine.UberElement {
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

func GitCommit(children ...engine.UberChild) *engine.UberElement {
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

func GitMerge(children ...engine.UberChild) *engine.UberElement {
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

func GitPullRequest(children ...engine.UberChild) *engine.UberElement {
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

func Gps(children ...engine.UberChild) *engine.UberElement {
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

func Group(children ...engine.UberChild) *engine.UberElement {
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

func Hd(children ...engine.UberChild) *engine.UberElement {
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

func Headphone(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(headphoneInnerSVG).
		Element(children...)
}

func Headset(children ...engine.UberChild) *engine.UberElement {
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

func Hidden(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hiddenInnerSVG).
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

func Hourglass(children ...engine.UberChild) *engine.UberElement {
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

func Hq(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hqInnerSVG).
		Element(children...)
}

func Human(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(humanInnerSVG).
		Element(children...)
}

func HumanHandsdown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(humanHandsdownInnerSVG).
		Element(children...)
}

func HumanHandsup(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(humanHandsupInnerSVG).
		Element(children...)
}

func HumanHeight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(humanHeightInnerSVG).
		Element(children...)
}

func HumanHeightAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(humanHeightAltInnerSVG).
		Element(children...)
}

func HumanRun(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(humanRunInnerSVG).
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

func ImageArrowRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageArrowRightInnerSVG).
		Element(children...)
}

func ImageBroken(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageBrokenInnerSVG).
		Element(children...)
}

func ImageDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageDeleteInnerSVG).
		Element(children...)
}

func ImageFlash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageFlashInnerSVG).
		Element(children...)
}

func ImageFrame(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageFrameInnerSVG).
		Element(children...)
}

func ImageGallery(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageGalleryInnerSVG).
		Element(children...)
}

func ImageMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageMultipleInnerSVG).
		Element(children...)
}

func ImageNew(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageNewInnerSVG).
		Element(children...)
}

func ImagePlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imagePlusInnerSVG).
		Element(children...)
}

func Inbox(children ...engine.UberChild) *engine.UberElement {
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

func InboxAll(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxAllInnerSVG).
		Element(children...)
}

func InboxFull(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inboxFullInnerSVG).
		Element(children...)
}

func InfoBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoBoxInnerSVG).
		Element(children...)
}

func Invert(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(invertInnerSVG).
		Element(children...)
}

func Iso(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(isoInnerSVG).
		Element(children...)
}

func Kanban(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kanbanInnerSVG).
		Element(children...)
}

func Keyboard(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyboardInnerSVG).
		Element(children...)
}

func Label(children ...engine.UberChild) *engine.UberElement {
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

func LabelAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(labelAltInnerSVG).
		Element(children...)
}

func LabelAltMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(labelAltMultipleInnerSVG).
		Element(children...)
}

func LabelSharp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(labelSharpInnerSVG).
		Element(children...)
}

func Layout(children ...engine.UberChild) *engine.UberElement {
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

func LayoutAlignBottom(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutAlignBottomInnerSVG).
		Element(children...)
}

func LayoutAlignLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutAlignLeftInnerSVG).
		Element(children...)
}

func LayoutAlignRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutAlignRightInnerSVG).
		Element(children...)
}

func LayoutAlignTop(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutAlignTopInnerSVG).
		Element(children...)
}

func LayoutColumns(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutColumnsInnerSVG).
		Element(children...)
}

func LayoutDistributeHorizontal(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutDistributeHorizontalInnerSVG).
		Element(children...)
}

func LayoutDistributeVertical(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutDistributeVerticalInnerSVG).
		Element(children...)
}

func LayoutFooter(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutFooterInnerSVG).
		Element(children...)
}

func LayoutHeader(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutHeaderInnerSVG).
		Element(children...)
}

func LayoutRows(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutRowsInnerSVG).
		Element(children...)
}

func LayoutSidebarLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutSidebarLeftInnerSVG).
		Element(children...)
}

func LayoutSidebarRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutSidebarRightInnerSVG).
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

func List(children ...engine.UberChild) *engine.UberElement {
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

func ListBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(listBoxInnerSVG).
		Element(children...)
}

func Loader(children ...engine.UberChild) *engine.UberElement {
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

func Login(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loginInnerSVG).
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

func Luggage(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(luggageInnerSVG).
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

func MailArrowRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailArrowRightInnerSVG).
		Element(children...)
}

func MailCheck(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailCheckInnerSVG).
		Element(children...)
}

func MailDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailDeleteInnerSVG).
		Element(children...)
}

func MailFlash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailFlashInnerSVG).
		Element(children...)
}

func MailMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailMultipleInnerSVG).
		Element(children...)
}

func MailOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailOffInnerSVG).
		Element(children...)
}

func MailUnread(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailUnreadInnerSVG).
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

func Membercard(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(membercardInnerSVG).
		Element(children...)
}

func Menu(children ...engine.UberChild) *engine.UberElement {
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

func Message(children ...engine.UberChild) *engine.UberElement {
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

func MessageArrowLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageArrowLeftInnerSVG).
		Element(children...)
}

func MessageArrowRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageArrowRightInnerSVG).
		Element(children...)
}

func MessageBookmark(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageBookmarkInnerSVG).
		Element(children...)
}

func MessageClock(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageClockInnerSVG).
		Element(children...)
}

func MessageDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageDeleteInnerSVG).
		Element(children...)
}

func MessageFlash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageFlashInnerSVG).
		Element(children...)
}

func MessageImage(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageImageInnerSVG).
		Element(children...)
}

func MessageMinus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageMinusInnerSVG).
		Element(children...)
}

func MessagePlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messagePlusInnerSVG).
		Element(children...)
}

func MessageProcessing(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageProcessingInnerSVG).
		Element(children...)
}

func MessageReply(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageReplyInnerSVG).
		Element(children...)
}

func MessageText(children ...engine.UberChild) *engine.UberElement {
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

func MissedCall(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(missedCallInnerSVG).
		Element(children...)
}

func Modem(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(modemInnerSVG).
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

func Monitor(children ...engine.UberChild) *engine.UberElement {
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

func MoodHappy(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodHappyInnerSVG).
		Element(children...)
}

func MoodNeutral(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodNeutralInnerSVG).
		Element(children...)
}

func MoodSad(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moodSadInnerSVG).
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

func MoonStar(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonStarInnerSVG).
		Element(children...)
}

func MoonStars(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moonStarsInnerSVG).
		Element(children...)
}

func MoreHorizontal(children ...engine.UberChild) *engine.UberElement {
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

func MoreVertical(children ...engine.UberChild) *engine.UberElement {
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

func Mouse(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mouseInnerSVG).
		Element(children...)
}

func Move(children ...engine.UberChild) *engine.UberElement {
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

func Movie(children ...engine.UberChild) *engine.UberElement {
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

func Music(children ...engine.UberChild) *engine.UberElement {
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

func Next(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nextInnerSVG).
		Element(children...)
}

func Note(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noteInnerSVG).
		Element(children...)
}

func NoteDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noteDeleteInnerSVG).
		Element(children...)
}

func NoteMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noteMultipleInnerSVG).
		Element(children...)
}

func NotePlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notePlusInnerSVG).
		Element(children...)
}

func Notes(children ...engine.UberChild) *engine.UberElement {
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

func NotesDelete(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notesDeleteInnerSVG).
		Element(children...)
}

func NotesMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notesMultipleInnerSVG).
		Element(children...)
}

func NotesPlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notesPlusInnerSVG).
		Element(children...)
}

func Notification(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notificationInnerSVG).
		Element(children...)
}

func NotificationOff(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notificationOffInnerSVG).
		Element(children...)
}

func Open(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(openInnerSVG).
		Element(children...)
}

func PaintBucket(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paintBucketInnerSVG).
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

func Percent(children ...engine.UberChild) *engine.UberElement {
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

func PictureInPicture(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pictureInPictureInnerSVG).
		Element(children...)
}

func PictureInPictureAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pictureInPictureAltInnerSVG).
		Element(children...)
}

func Pin(children ...engine.UberChild) *engine.UberElement {
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

func Pixelarticons(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pixelarticonsInnerSVG).
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

func Playlist(children ...engine.UberChild) *engine.UberElement {
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

func Prev(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(prevInnerSVG).
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

func RadioHandheld(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radioHandheldInnerSVG).
		Element(children...)
}

func RadioOn(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radioOnInnerSVG).
		Element(children...)
}

func RadioSignal(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radioSignalInnerSVG).
		Element(children...)
}

func RadioTower(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radioTowerInnerSVG).
		Element(children...)
}

func Reciept(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(recieptInnerSVG).
		Element(children...)
}

func RecieptAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(recieptAltInnerSVG).
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

func Reload(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reloadInnerSVG).
		Element(children...)
}

func RemoveBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeBoxInnerSVG).
		Element(children...)
}

func RemoveBoxMultiple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeBoxMultipleInnerSVG).
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

func Reply(children ...engine.UberChild) *engine.UberElement {
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

func ReplyAll(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(replyAllInnerSVG).
		Element(children...)
}

func RoundedCorner(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(roundedCornerInnerSVG).
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

func Scale(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scaleInnerSVG).
		Element(children...)
}

func Script(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scriptInnerSVG).
		Element(children...)
}

func ScriptText(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scriptTextInnerSVG).
		Element(children...)
}

func ScrollHorizontal(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scrollHorizontalInnerSVG).
		Element(children...)
}

func ScrollVertical(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scrollVerticalInnerSVG).
		Element(children...)
}

func Sd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sdInnerSVG).
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

func Section(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionInnerSVG).
		Element(children...)
}

func SectionCopy(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionCopyInnerSVG).
		Element(children...)
}

func SectionMinus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionMinusInnerSVG).
		Element(children...)
}

func SectionPlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionPlusInnerSVG).
		Element(children...)
}

func SectionX(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sectionXInnerSVG).
		Element(children...)
}

func Server(children ...engine.UberChild) *engine.UberElement {
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

func SharpCorner(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sharpCornerInnerSVG).
		Element(children...)
}

func Shield(children ...engine.UberChild) *engine.UberElement {
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

func ShieldOff(children ...engine.UberChild) *engine.UberElement {
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

func ShoppingBag(children ...engine.UberChild) *engine.UberElement {
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

func Shuffle(children ...engine.UberChild) *engine.UberElement {
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

func Sliders(children ...engine.UberChild) *engine.UberElement {
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

func SlidersTwo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slidersTwoInnerSVG).
		Element(children...)
}

func Sort(children ...engine.UberChild) *engine.UberElement {
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

func SortAlpabetic(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortAlpabeticInnerSVG).
		Element(children...)
}

func SortNumeric(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sortNumericInnerSVG).
		Element(children...)
}

func Speaker(children ...engine.UberChild) *engine.UberElement {
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

func SpeedFast(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speedFastInnerSVG).
		Element(children...)
}

func SpeedMedium(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speedMediumInnerSVG).
		Element(children...)
}

func SpeedSlow(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speedSlowInnerSVG).
		Element(children...)
}

func Spotlight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spotlightInnerSVG).
		Element(children...)
}

func Store(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(storeInnerSVG).
		Element(children...)
}

func Subscriptions(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subscriptionsInnerSVG).
		Element(children...)
}

func Subtitles(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subtitlesInnerSVG).
		Element(children...)
}

func Suitcase(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(suitcaseInnerSVG).
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

func SunAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunAltInnerSVG).
		Element(children...)
}

func Switch(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(switchInnerSVG).
		Element(children...)
}

func Sync(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(syncInnerSVG).
		Element(children...)
}

func Tab(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tabInnerSVG).
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

func Tea(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(teaInnerSVG).
		Element(children...)
}

func Teach(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(teachInnerSVG).
		Element(children...)
}

func TextAdd(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAddInnerSVG).
		Element(children...)
}

func TextColums(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textColumsInnerSVG).
		Element(children...)
}

func TextSearch(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textSearchInnerSVG).
		Element(children...)
}

func TextWrap(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textWrapInnerSVG).
		Element(children...)
}

func Timeline(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timelineInnerSVG).
		Element(children...)
}

func ToggleLeft(children ...engine.UberChild) *engine.UberElement {
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

func ToggleRight(children ...engine.UberChild) *engine.UberElement {
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

func Tournament(children ...engine.UberChild) *engine.UberElement {
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

func TrackChanges(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trackChangesInnerSVG).
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

func TrashAlt(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashAltInnerSVG).
		Element(children...)
}

func Trending(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trendingInnerSVG).
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

func Trophy(children ...engine.UberChild) *engine.UberElement {
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

func Undo(children ...engine.UberChild) *engine.UberElement {
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

func Ungroup(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ungroupInnerSVG).
		Element(children...)
}

func Unlink(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlinkInnerSVG).
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

func UserMinus(children ...engine.UberChild) *engine.UberElement {
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

func UserPlus(children ...engine.UberChild) *engine.UberElement {
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

func UserX(children ...engine.UberChild) *engine.UberElement {
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

func Video(children ...engine.UberChild) *engine.UberElement {
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

func VideoOff(children ...engine.UberChild) *engine.UberElement {
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

func ViewCol(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewColInnerSVG).
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

func ViewportNarrow(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewportNarrowInnerSVG).
		Element(children...)
}

func ViewportWide(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewportWideInnerSVG).
		Element(children...)
}

func Visible(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(visibleInnerSVG).
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

func VolumeMinus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeMinusInnerSVG).
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

func VolumePlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumePlusInnerSVG).
		Element(children...)
}

func VolumeThree(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeThreeInnerSVG).
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

func VolumeVibrate(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeVibrateInnerSVG).
		Element(children...)
}

func VolumeX(children ...engine.UberChild) *engine.UberElement {
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

func Wallet(children ...engine.UberChild) *engine.UberElement {
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

func WarningBox(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(warningBoxInnerSVG).
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

func Zap(children ...engine.UberChild) *engine.UberElement {
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
	case "ab-testing":
		return AbTesting(), nil
	case "ac":
		return Ac(), nil
	case "add-box":
		return AddBox(), nil
	case "add-box-multiple":
		return AddBoxMultiple(), nil
	case "add-col":
		return AddCol(), nil
	case "add-grid":
		return AddGrid(), nil
	case "add-row":
		return AddRow(), nil
	case "alert":
		return Alert(), nil
	case "align-center":
		return AlignCenter(), nil
	case "align-justify":
		return AlignJustify(), nil
	case "align-left":
		return AlignLeft(), nil
	case "align-right":
		return AlignRight(), nil
	case "analytics":
		return Analytics(), nil
	case "anchor":
		return Anchor(), nil
	case "android":
		return Android(), nil
	case "animation":
		return Animation(), nil
	case "archive":
		return Archive(), nil
	case "arrow-bar-down":
		return ArrowBarDown(), nil
	case "arrow-bar-left":
		return ArrowBarLeft(), nil
	case "arrow-bar-right":
		return ArrowBarRight(), nil
	case "arrow-bar-up":
		return ArrowBarUp(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-down-box":
		return ArrowDownBox(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-left-box":
		return ArrowLeftBox(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-right-box":
		return ArrowRightBox(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrow-up-box":
		return ArrowUpBox(), nil
	case "arrows-horizontal":
		return ArrowsHorizontal(), nil
	case "arrows-vertical":
		return ArrowsVertical(), nil
	case "art-text":
		return ArtText(), nil
	case "article":
		return Article(), nil
	case "article-multiple":
		return ArticleMultiple(), nil
	case "aspect-ratio":
		return AspectRatio(), nil
	case "at":
		return At(), nil
	case "attachment":
		return Attachment(), nil
	case "audio-device":
		return AudioDevice(), nil
	case "avatar":
		return Avatar(), nil
	case "backburger":
		return Backburger(), nil
	case "battery":
		return Battery(), nil
	case "battery-charging":
		return BatteryCharging(), nil
	case "battery-full":
		return BatteryFull(), nil
	case "battery-1":
		return BatteryOne(), nil
	case "battery-2":
		return BatteryTwo(), nil
	case "bed":
		return Bed(), nil
	case "bitcoin":
		return Bitcoin(), nil
	case "bluetooth":
		return Bluetooth(), nil
	case "book":
		return Book(), nil
	case "book-open":
		return BookOpen(), nil
	case "bookmark":
		return Bookmark(), nil
	case "bookmarks":
		return Bookmarks(), nil
	case "briefcase":
		return Briefcase(), nil
	case "briefcase-account":
		return BriefcaseAccount(), nil
	case "briefcase-check":
		return BriefcaseCheck(), nil
	case "briefcase-delete":
		return BriefcaseDelete(), nil
	case "briefcase-download":
		return BriefcaseDownload(), nil
	case "briefcase-minus":
		return BriefcaseMinus(), nil
	case "briefcase-plus":
		return BriefcasePlus(), nil
	case "briefcase-search":
		return BriefcaseSearch(), nil
	case "briefcase-search-1":
		return BriefcaseSearchOne(), nil
	case "briefcase-upload":
		return BriefcaseUpload(), nil
	case "bug":
		return Bug(), nil
	case "building":
		return Building(), nil
	case "building-community":
		return BuildingCommunity(), nil
	case "building-skyscraper":
		return BuildingSkyscraper(), nil
	case "buildings":
		return Buildings(), nil
	case "bulletlist":
		return Bulletlist(), nil
	case "bullseye":
		return Bullseye(), nil
	case "bullseye-arrow":
		return BullseyeArrow(), nil
	case "bus":
		return Bus(), nil
	case "cake":
		return Cake(), nil
	case "calculator":
		return Calculator(), nil
	case "calendar":
		return Calendar(), nil
	case "calendar-alert":
		return CalendarAlert(), nil
	case "calendar-arrow-left":
		return CalendarArrowLeft(), nil
	case "calendar-arrow-right":
		return CalendarArrowRight(), nil
	case "calendar-check":
		return CalendarCheck(), nil
	case "calendar-export":
		return CalendarExport(), nil
	case "calendar-grid":
		return CalendarGrid(), nil
	case "calendar-import":
		return CalendarImport(), nil
	case "calendar-minus":
		return CalendarMinus(), nil
	case "calendar-month":
		return CalendarMonth(), nil
	case "calendar-multiple":
		return CalendarMultiple(), nil
	case "calendar-multiple-check":
		return CalendarMultipleCheck(), nil
	case "calendar-plus":
		return CalendarPlus(), nil
	case "calendar-range":
		return CalendarRange(), nil
	case "calendar-remove":
		return CalendarRemove(), nil
	case "calendar-search":
		return CalendarSearch(), nil
	case "calendar-sort-ascending":
		return CalendarSortAscending(), nil
	case "calendar-sort-descending":
		return CalendarSortDescending(), nil
	case "calendar-text":
		return CalendarText(), nil
	case "calendar-today":
		return CalendarToday(), nil
	case "calendar-tomorrow":
		return CalendarTomorrow(), nil
	case "calendar-week":
		return CalendarWeek(), nil
	case "calendar-week-begin":
		return CalendarWeekBegin(), nil
	case "calendar-weekend":
		return CalendarWeekend(), nil
	case "camera":
		return Camera(), nil
	case "camera-add":
		return CameraAdd(), nil
	case "camera-alt":
		return CameraAlt(), nil
	case "camera-face":
		return CameraFace(), nil
	case "car":
		return Car(), nil
	case "card":
		return Card(), nil
	case "card-id":
		return CardId(), nil
	case "card-plus":
		return CardPlus(), nil
	case "card-stack":
		return CardStack(), nil
	case "card-text":
		return CardText(), nil
	case "cart":
		return Cart(), nil
	case "cast":
		return Cast(), nil
	case "cellular-signal-off":
		return CellularSignalOff(), nil
	case "cellular-signal-1":
		return CellularSignalOne(), nil
	case "cellular-signal-3":
		return CellularSignalThree(), nil
	case "cellular-signal-2":
		return CellularSignalTwo(), nil
	case "cellular-signal-0":
		return CellularSignalZero(), nil
	case "chart":
		return Chart(), nil
	case "chart-add":
		return ChartAdd(), nil
	case "chart-bar":
		return ChartBar(), nil
	case "chart-delete":
		return ChartDelete(), nil
	case "chart-minus":
		return ChartMinus(), nil
	case "chart-multiple":
		return ChartMultiple(), nil
	case "chat":
		return Chat(), nil
	case "check":
		return Check(), nil
	case "check-double":
		return CheckDouble(), nil
	case "checkbox":
		return Checkbox(), nil
	case "checkbox-on":
		return CheckboxOn(), nil
	case "checklist":
		return Checklist(), nil
	case "chess":
		return Chess(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "chevrons-horizontal":
		return ChevronsHorizontal(), nil
	case "chevrons-vertical":
		return ChevronsVertical(), nil
	case "circle":
		return Circle(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clock":
		return Clock(), nil
	case "close":
		return Close(), nil
	case "close-box":
		return CloseBox(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-done":
		return CloudDone(), nil
	case "cloud-download":
		return CloudDownload(), nil
	case "cloud-moon":
		return CloudMoon(), nil
	case "cloud-sun":
		return CloudSun(), nil
	case "cloud-upload":
		return CloudUpload(), nil
	case "cocktail":
		return Cocktail(), nil
	case "code":
		return Code(), nil
	case "coffee":
		return Coffee(), nil
	case "coffee-alt":
		return CoffeeAlt(), nil
	case "coin":
		return Coin(), nil
	case "collapse":
		return Collapse(), nil
	case "colors-swatch":
		return ColorsSwatch(), nil
	case "command":
		return Command(), nil
	case "comment":
		return Comment(), nil
	case "contact":
		return Contact(), nil
	case "contact-delete":
		return ContactDelete(), nil
	case "contact-multiple":
		return ContactMultiple(), nil
	case "contact-plus":
		return ContactPlus(), nil
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
	case "credit-card":
		return CreditCard(), nil
	case "credit-card-delete":
		return CreditCardDelete(), nil
	case "credit-card-minus":
		return CreditCardMinus(), nil
	case "credit-card-multiple":
		return CreditCardMultiple(), nil
	case "credit-card-plus":
		return CreditCardPlus(), nil
	case "credit-card-settings":
		return CreditCardSettings(), nil
	case "credit-card-wireless":
		return CreditCardWireless(), nil
	case "crop":
		return Crop(), nil
	case "cut":
		return Cut(), nil
	case "dashbaord":
		return Dashbaord(), nil
	case "debug":
		return Debug(), nil
	case "debug-check":
		return DebugCheck(), nil
	case "debug-off":
		return DebugOff(), nil
	case "debug-pause":
		return DebugPause(), nil
	case "debug-play":
		return DebugPlay(), nil
	case "debug-stop":
		return DebugStop(), nil
	case "delete":
		return Delete(), nil
	case "deskphone":
		return Deskphone(), nil
	case "device-laptop":
		return DeviceLaptop(), nil
	case "device-phone":
		return DevicePhone(), nil
	case "device-tablet":
		return DeviceTablet(), nil
	case "device-tv":
		return DeviceTv(), nil
	case "device-tv-smart":
		return DeviceTvSmart(), nil
	case "device-vibrate":
		return DeviceVibrate(), nil
	case "device-watch":
		return DeviceWatch(), nil
	case "devices":
		return Devices(), nil
	case "dice":
		return Dice(), nil
	case "dollar":
		return Dollar(), nil
	case "downasaur":
		return Downasaur(), nil
	case "download":
		return Download(), nil
	case "draft":
		return Draft(), nil
	case "drag-and-drop":
		return DragAndDrop(), nil
	case "drop":
		return Drop(), nil
	case "drop-area":
		return DropArea(), nil
	case "drop-full":
		return DropFull(), nil
	case "drop-half":
		return DropHalf(), nil
	case "duplicate":
		return Duplicate(), nil
	case "duplicate-alt":
		return DuplicateAlt(), nil
	case "edit":
		return Edit(), nil
	case "edit-box":
		return EditBox(), nil
	case "euro":
		return Euro(), nil
	case "expand":
		return Expand(), nil
	case "external-link":
		return ExternalLink(), nil
	case "eye":
		return Eye(), nil
	case "eye-closed":
		return EyeClosed(), nil
	case "file":
		return File(), nil
	case "file-alt":
		return FileAlt(), nil
	case "file-delete":
		return FileDelete(), nil
	case "file-flash":
		return FileFlash(), nil
	case "file-minus":
		return FileMinus(), nil
	case "file-multiple":
		return FileMultiple(), nil
	case "file-off":
		return FileOff(), nil
	case "file-plus":
		return FilePlus(), nil
	case "fill":
		return Fill(), nil
	case "fill-half":
		return FillHalf(), nil
	case "5g":
		return FiveG(), nil
	case "flag":
		return Flag(), nil
	case "flatten":
		return Flatten(), nil
	case "flip-to-back":
		return FlipToBack(), nil
	case "flip-to-front":
		return FlipToFront(), nil
	case "float-center":
		return FloatCenter(), nil
	case "float-left":
		return FloatLeft(), nil
	case "float-right":
		return FloatRight(), nil
	case "folder":
		return Folder(), nil
	case "folder-minus":
		return FolderMinus(), nil
	case "folder-plus":
		return FolderPlus(), nil
	case "folder-x":
		return FolderX(), nil
	case "forward":
		return Forward(), nil
	case "forwardburger":
		return Forwardburger(), nil
	case "4g":
		return FourG(), nil
	case "4k":
		return FourK(), nil
	case "4k-box":
		return FourKBox(), nil
	case "frame":
		return Frame(), nil
	case "frame-add":
		return FrameAdd(), nil
	case "frame-check":
		return FrameCheck(), nil
	case "frame-delete":
		return FrameDelete(), nil
	case "frame-minus":
		return FrameMinus(), nil
	case "gamepad":
		return Gamepad(), nil
	case "gif":
		return Gif(), nil
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
	case "gps":
		return Gps(), nil
	case "grid":
		return Grid(), nil
	case "group":
		return Group(), nil
	case "hd":
		return Hd(), nil
	case "headphone":
		return Headphone(), nil
	case "headset":
		return Headset(), nil
	case "heart":
		return Heart(), nil
	case "hidden":
		return Hidden(), nil
	case "home":
		return Home(), nil
	case "hourglass":
		return Hourglass(), nil
	case "hq":
		return Hq(), nil
	case "human":
		return Human(), nil
	case "human-handsdown":
		return HumanHandsdown(), nil
	case "human-handsup":
		return HumanHandsup(), nil
	case "human-height":
		return HumanHeight(), nil
	case "human-height-alt":
		return HumanHeightAlt(), nil
	case "human-run":
		return HumanRun(), nil
	case "image":
		return Image(), nil
	case "image-arrow-right":
		return ImageArrowRight(), nil
	case "image-broken":
		return ImageBroken(), nil
	case "image-delete":
		return ImageDelete(), nil
	case "image-flash":
		return ImageFlash(), nil
	case "image-frame":
		return ImageFrame(), nil
	case "image-gallery":
		return ImageGallery(), nil
	case "image-multiple":
		return ImageMultiple(), nil
	case "image-new":
		return ImageNew(), nil
	case "image-plus":
		return ImagePlus(), nil
	case "inbox":
		return Inbox(), nil
	case "inbox-all":
		return InboxAll(), nil
	case "inbox-full":
		return InboxFull(), nil
	case "info-box":
		return InfoBox(), nil
	case "invert":
		return Invert(), nil
	case "iso":
		return Iso(), nil
	case "kanban":
		return Kanban(), nil
	case "keyboard":
		return Keyboard(), nil
	case "label":
		return Label(), nil
	case "label-alt":
		return LabelAlt(), nil
	case "label-alt-multiple":
		return LabelAltMultiple(), nil
	case "label-sharp":
		return LabelSharp(), nil
	case "layout":
		return Layout(), nil
	case "layout-align-bottom":
		return LayoutAlignBottom(), nil
	case "layout-align-left":
		return LayoutAlignLeft(), nil
	case "layout-align-right":
		return LayoutAlignRight(), nil
	case "layout-align-top":
		return LayoutAlignTop(), nil
	case "layout-columns":
		return LayoutColumns(), nil
	case "layout-distribute-horizontal":
		return LayoutDistributeHorizontal(), nil
	case "layout-distribute-vertical":
		return LayoutDistributeVertical(), nil
	case "layout-footer":
		return LayoutFooter(), nil
	case "layout-header":
		return LayoutHeader(), nil
	case "layout-rows":
		return LayoutRows(), nil
	case "layout-sidebar-left":
		return LayoutSidebarLeft(), nil
	case "layout-sidebar-right":
		return LayoutSidebarRight(), nil
	case "link":
		return Link(), nil
	case "list":
		return List(), nil
	case "list-box":
		return ListBox(), nil
	case "loader":
		return Loader(), nil
	case "lock":
		return Lock(), nil
	case "lock-open":
		return LockOpen(), nil
	case "login":
		return Login(), nil
	case "logout":
		return Logout(), nil
	case "luggage":
		return Luggage(), nil
	case "mail":
		return Mail(), nil
	case "mail-arrow-right":
		return MailArrowRight(), nil
	case "mail-check":
		return MailCheck(), nil
	case "mail-delete":
		return MailDelete(), nil
	case "mail-flash":
		return MailFlash(), nil
	case "mail-multiple":
		return MailMultiple(), nil
	case "mail-off":
		return MailOff(), nil
	case "mail-unread":
		return MailUnread(), nil
	case "map":
		return Map(), nil
	case "membercard":
		return Membercard(), nil
	case "menu":
		return Menu(), nil
	case "message":
		return Message(), nil
	case "message-arrow-left":
		return MessageArrowLeft(), nil
	case "message-arrow-right":
		return MessageArrowRight(), nil
	case "message-bookmark":
		return MessageBookmark(), nil
	case "message-clock":
		return MessageClock(), nil
	case "message-delete":
		return MessageDelete(), nil
	case "message-flash":
		return MessageFlash(), nil
	case "message-image":
		return MessageImage(), nil
	case "message-minus":
		return MessageMinus(), nil
	case "message-plus":
		return MessagePlus(), nil
	case "message-processing":
		return MessageProcessing(), nil
	case "message-reply":
		return MessageReply(), nil
	case "message-text":
		return MessageText(), nil
	case "minus":
		return Minus(), nil
	case "missed-call":
		return MissedCall(), nil
	case "modem":
		return Modem(), nil
	case "money":
		return Money(), nil
	case "monitor":
		return Monitor(), nil
	case "mood-happy":
		return MoodHappy(), nil
	case "mood-neutral":
		return MoodNeutral(), nil
	case "mood-sad":
		return MoodSad(), nil
	case "moon":
		return Moon(), nil
	case "moon-star":
		return MoonStar(), nil
	case "moon-stars":
		return MoonStars(), nil
	case "more-horizontal":
		return MoreHorizontal(), nil
	case "more-vertical":
		return MoreVertical(), nil
	case "mouse":
		return Mouse(), nil
	case "move":
		return Move(), nil
	case "movie":
		return Movie(), nil
	case "music":
		return Music(), nil
	case "next":
		return Next(), nil
	case "note":
		return Note(), nil
	case "note-delete":
		return NoteDelete(), nil
	case "note-multiple":
		return NoteMultiple(), nil
	case "note-plus":
		return NotePlus(), nil
	case "notes":
		return Notes(), nil
	case "notes-delete":
		return NotesDelete(), nil
	case "notes-multiple":
		return NotesMultiple(), nil
	case "notes-plus":
		return NotesPlus(), nil
	case "notification":
		return Notification(), nil
	case "notification-off":
		return NotificationOff(), nil
	case "open":
		return Open(), nil
	case "paint-bucket":
		return PaintBucket(), nil
	case "paperclip":
		return Paperclip(), nil
	case "pause":
		return Pause(), nil
	case "percent":
		return Percent(), nil
	case "picture-in-picture":
		return PictureInPicture(), nil
	case "picture-in-picture-alt":
		return PictureInPictureAlt(), nil
	case "pin":
		return Pin(), nil
	case "pixelarticons":
		return Pixelarticons(), nil
	case "play":
		return Play(), nil
	case "playlist":
		return Playlist(), nil
	case "plus":
		return Plus(), nil
	case "power":
		return Power(), nil
	case "prev":
		return Prev(), nil
	case "print":
		return Print(), nil
	case "radio-handheld":
		return RadioHandheld(), nil
	case "radio-on":
		return RadioOn(), nil
	case "radio-signal":
		return RadioSignal(), nil
	case "radio-tower":
		return RadioTower(), nil
	case "reciept":
		return Reciept(), nil
	case "reciept-alt":
		return RecieptAlt(), nil
	case "redo":
		return Redo(), nil
	case "reload":
		return Reload(), nil
	case "remove-box":
		return RemoveBox(), nil
	case "remove-box-multiple":
		return RemoveBoxMultiple(), nil
	case "repeat":
		return Repeat(), nil
	case "reply":
		return Reply(), nil
	case "reply-all":
		return ReplyAll(), nil
	case "rounded-corner":
		return RoundedCorner(), nil
	case "save":
		return Save(), nil
	case "scale":
		return Scale(), nil
	case "script":
		return Script(), nil
	case "script-text":
		return ScriptText(), nil
	case "scroll-horizontal":
		return ScrollHorizontal(), nil
	case "scroll-vertical":
		return ScrollVertical(), nil
	case "sd":
		return Sd(), nil
	case "search":
		return Search(), nil
	case "section":
		return Section(), nil
	case "section-copy":
		return SectionCopy(), nil
	case "section-minus":
		return SectionMinus(), nil
	case "section-plus":
		return SectionPlus(), nil
	case "section-x":
		return SectionX(), nil
	case "server":
		return Server(), nil
	case "sharp-corner":
		return SharpCorner(), nil
	case "shield":
		return Shield(), nil
	case "shield-off":
		return ShieldOff(), nil
	case "ship":
		return Ship(), nil
	case "shopping-bag":
		return ShoppingBag(), nil
	case "shuffle":
		return Shuffle(), nil
	case "sliders":
		return Sliders(), nil
	case "sliders-2":
		return SlidersTwo(), nil
	case "sort":
		return Sort(), nil
	case "sort-alpabetic":
		return SortAlpabetic(), nil
	case "sort-numeric":
		return SortNumeric(), nil
	case "speaker":
		return Speaker(), nil
	case "speed-fast":
		return SpeedFast(), nil
	case "speed-medium":
		return SpeedMedium(), nil
	case "speed-slow":
		return SpeedSlow(), nil
	case "spotlight":
		return Spotlight(), nil
	case "store":
		return Store(), nil
	case "subscriptions":
		return Subscriptions(), nil
	case "subtitles":
		return Subtitles(), nil
	case "suitcase":
		return Suitcase(), nil
	case "sun":
		return Sun(), nil
	case "sun-alt":
		return SunAlt(), nil
	case "switch":
		return Switch(), nil
	case "sync":
		return Sync(), nil
	case "tab":
		return Tab(), nil
	case "table":
		return Table(), nil
	case "tea":
		return Tea(), nil
	case "teach":
		return Teach(), nil
	case "text-add":
		return TextAdd(), nil
	case "text-colums":
		return TextColums(), nil
	case "text-search":
		return TextSearch(), nil
	case "text-wrap":
		return TextWrap(), nil
	case "timeline":
		return Timeline(), nil
	case "toggle-left":
		return ToggleLeft(), nil
	case "toggle-right":
		return ToggleRight(), nil
	case "tournament":
		return Tournament(), nil
	case "track-changes":
		return TrackChanges(), nil
	case "trash":
		return Trash(), nil
	case "trash-alt":
		return TrashAlt(), nil
	case "trending":
		return Trending(), nil
	case "trending-down":
		return TrendingDown(), nil
	case "trending-up":
		return TrendingUp(), nil
	case "trophy":
		return Trophy(), nil
	case "truck":
		return Truck(), nil
	case "undo":
		return Undo(), nil
	case "ungroup":
		return Ungroup(), nil
	case "unlink":
		return Unlink(), nil
	case "upload":
		return Upload(), nil
	case "user":
		return User(), nil
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
	case "view-col":
		return ViewCol(), nil
	case "view-list":
		return ViewList(), nil
	case "viewport-narrow":
		return ViewportNarrow(), nil
	case "viewport-wide":
		return ViewportWide(), nil
	case "visible":
		return Visible(), nil
	case "volume":
		return Volume(), nil
	case "volume-minus":
		return VolumeMinus(), nil
	case "volume-1":
		return VolumeOne(), nil
	case "volume-plus":
		return VolumePlus(), nil
	case "volume-3":
		return VolumeThree(), nil
	case "volume-2":
		return VolumeTwo(), nil
	case "volume-vibrate":
		return VolumeVibrate(), nil
	case "volume-x":
		return VolumeX(), nil
	case "wallet":
		return Wallet(), nil
	case "warning-box":
		return WarningBox(), nil
	case "wind":
		return Wind(), nil
	case "zap":
		return Zap(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in pixelarticons icon set", name)
	}
}
