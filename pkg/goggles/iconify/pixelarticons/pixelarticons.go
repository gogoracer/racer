package pixelarticons

import "github.com/gogoracer/racer/pkg/engine"

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

var sharedIconAttrs = engine.Attrs{"width": "1em", "height": "1em"}

func AbTesting(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		abTestingInnerSVG,
		children,
	)
}

func Ac(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		acInnerSVG,
		children,
	)
}

func AddBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addBoxInnerSVG,
		children,
	)
}

func AddBoxMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addBoxMultipleInnerSVG,
		children,
	)
}

func AddCol(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addColInnerSVG,
		children,
	)
}

func AddGrid(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addGridInnerSVG,
		children,
	)
}

func AddRow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		addRowInnerSVG,
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

func AlignRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alignRightInnerSVG,
		children,
	)
}

func Analytics(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		analyticsInnerSVG,
		children,
	)
}

func Anchor(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		anchorInnerSVG,
		children,
	)
}

func Android(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		androidInnerSVG,
		children,
	)
}

func Animation(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		animationInnerSVG,
		children,
	)
}

func Archive(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		archiveInnerSVG,
		children,
	)
}

func ArrowBarDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowBarDownInnerSVG,
		children,
	)
}

func ArrowBarLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowBarLeftInnerSVG,
		children,
	)
}

func ArrowBarRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowBarRightInnerSVG,
		children,
	)
}

func ArrowBarUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowBarUpInnerSVG,
		children,
	)
}

func ArrowDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownInnerSVG,
		children,
	)
}

func ArrowDownBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownBoxInnerSVG,
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

func ArrowLeftBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLeftBoxInnerSVG,
		children,
	)
}

func ArrowRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowRightInnerSVG,
		children,
	)
}

func ArrowRightBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowRightBoxInnerSVG,
		children,
	)
}

func ArrowUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpInnerSVG,
		children,
	)
}

func ArrowUpBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpBoxInnerSVG,
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

func ArrowsVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowsVerticalInnerSVG,
		children,
	)
}

func ArtText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		artTextInnerSVG,
		children,
	)
}

func Article(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		articleInnerSVG,
		children,
	)
}

func ArticleMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		articleMultipleInnerSVG,
		children,
	)
}

func AspectRatio(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		aspectRatioInnerSVG,
		children,
	)
}

func At(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		atInnerSVG,
		children,
	)
}

func Attachment(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		attachmentInnerSVG,
		children,
	)
}

func AudioDevice(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		audioDeviceInnerSVG,
		children,
	)
}

func Avatar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		avatarInnerSVG,
		children,
	)
}

func Backburger(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backburgerInnerSVG,
		children,
	)
}

func Battery(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryInnerSVG,
		children,
	)
}

func BatteryCharging(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryChargingInnerSVG,
		children,
	)
}

func BatteryFull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryFullInnerSVG,
		children,
	)
}

func BatteryOne(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryOneInnerSVG,
		children,
	)
}

func BatteryTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryTwoInnerSVG,
		children,
	)
}

func Bed(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bedInnerSVG,
		children,
	)
}

func Bitcoin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bitcoinInnerSVG,
		children,
	)
}

func Bluetooth(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bluetoothInnerSVG,
		children,
	)
}

func Book(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookInnerSVG,
		children,
	)
}

func BookOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookOpenInnerSVG,
		children,
	)
}

func Bookmark(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkInnerSVG,
		children,
	)
}

func Bookmarks(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarksInnerSVG,
		children,
	)
}

func Briefcase(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseInnerSVG,
		children,
	)
}

func BriefcaseAccount(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseAccountInnerSVG,
		children,
	)
}

func BriefcaseCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseCheckInnerSVG,
		children,
	)
}

func BriefcaseDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseDeleteInnerSVG,
		children,
	)
}

func BriefcaseDownload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseDownloadInnerSVG,
		children,
	)
}

func BriefcaseMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseMinusInnerSVG,
		children,
	)
}

func BriefcasePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcasePlusInnerSVG,
		children,
	)
}

func BriefcaseSearch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseSearchInnerSVG,
		children,
	)
}

func BriefcaseSearchOne(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseSearchOneInnerSVG,
		children,
	)
}

func BriefcaseUpload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseUploadInnerSVG,
		children,
	)
}

func Bug(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bugInnerSVG,
		children,
	)
}

func Building(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		buildingInnerSVG,
		children,
	)
}

func BuildingCommunity(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		buildingCommunityInnerSVG,
		children,
	)
}

func BuildingSkyscraper(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		buildingSkyscraperInnerSVG,
		children,
	)
}

func Buildings(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		buildingsInnerSVG,
		children,
	)
}

func Bulletlist(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bulletlistInnerSVG,
		children,
	)
}

func Bullseye(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bullseyeInnerSVG,
		children,
	)
}

func BullseyeArrow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bullseyeArrowInnerSVG,
		children,
	)
}

func Bus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		busInnerSVG,
		children,
	)
}

func Cake(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cakeInnerSVG,
		children,
	)
}

func Calculator(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calculatorInnerSVG,
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

func CalendarAlert(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarAlertInnerSVG,
		children,
	)
}

func CalendarArrowLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarArrowLeftInnerSVG,
		children,
	)
}

func CalendarArrowRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarArrowRightInnerSVG,
		children,
	)
}

func CalendarCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarCheckInnerSVG,
		children,
	)
}

func CalendarExport(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarExportInnerSVG,
		children,
	)
}

func CalendarGrid(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarGridInnerSVG,
		children,
	)
}

func CalendarImport(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarImportInnerSVG,
		children,
	)
}

func CalendarMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarMinusInnerSVG,
		children,
	)
}

func CalendarMonth(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarMonthInnerSVG,
		children,
	)
}

func CalendarMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarMultipleInnerSVG,
		children,
	)
}

func CalendarMultipleCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarMultipleCheckInnerSVG,
		children,
	)
}

func CalendarPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarPlusInnerSVG,
		children,
	)
}

func CalendarRange(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarRangeInnerSVG,
		children,
	)
}

func CalendarRemove(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarRemoveInnerSVG,
		children,
	)
}

func CalendarSearch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarSearchInnerSVG,
		children,
	)
}

func CalendarSortAscending(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarSortAscendingInnerSVG,
		children,
	)
}

func CalendarSortDescending(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarSortDescendingInnerSVG,
		children,
	)
}

func CalendarText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarTextInnerSVG,
		children,
	)
}

func CalendarToday(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarTodayInnerSVG,
		children,
	)
}

func CalendarTomorrow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarTomorrowInnerSVG,
		children,
	)
}

func CalendarWeek(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarWeekInnerSVG,
		children,
	)
}

func CalendarWeekBegin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarWeekBeginInnerSVG,
		children,
	)
}

func CalendarWeekend(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarWeekendInnerSVG,
		children,
	)
}

func Camera(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraInnerSVG,
		children,
	)
}

func CameraAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraAddInnerSVG,
		children,
	)
}

func CameraAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraAltInnerSVG,
		children,
	)
}

func CameraFace(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraFaceInnerSVG,
		children,
	)
}

func Car(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		carInnerSVG,
		children,
	)
}

func Card(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cardInnerSVG,
		children,
	)
}

func CardId(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cardIdInnerSVG,
		children,
	)
}

func CardPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cardPlusInnerSVG,
		children,
	)
}

func CardStack(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cardStackInnerSVG,
		children,
	)
}

func CardText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cardTextInnerSVG,
		children,
	)
}

func Cart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cartInnerSVG,
		children,
	)
}

func Cast(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		castInnerSVG,
		children,
	)
}

func CellularSignalOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cellularSignalOffInnerSVG,
		children,
	)
}

func CellularSignalOne(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cellularSignalOneInnerSVG,
		children,
	)
}

func CellularSignalThree(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cellularSignalThreeInnerSVG,
		children,
	)
}

func CellularSignalTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cellularSignalTwoInnerSVG,
		children,
	)
}

func CellularSignalZero(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cellularSignalZeroInnerSVG,
		children,
	)
}

func Chart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartInnerSVG,
		children,
	)
}

func ChartAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartAddInnerSVG,
		children,
	)
}

func ChartBar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartBarInnerSVG,
		children,
	)
}

func ChartDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartDeleteInnerSVG,
		children,
	)
}

func ChartMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartMinusInnerSVG,
		children,
	)
}

func ChartMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chartMultipleInnerSVG,
		children,
	)
}

func Chat(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chatInnerSVG,
		children,
	)
}

func Check(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkInnerSVG,
		children,
	)
}

func CheckDouble(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkDoubleInnerSVG,
		children,
	)
}

func Checkbox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkboxInnerSVG,
		children,
	)
}

func CheckboxOn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkboxOnInnerSVG,
		children,
	)
}

func Checklist(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checklistInnerSVG,
		children,
	)
}

func Chess(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chessInnerSVG,
		children,
	)
}

func ChevronDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDownInnerSVG,
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

func ChevronRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronRightInnerSVG,
		children,
	)
}

func ChevronUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronUpInnerSVG,
		children,
	)
}

func ChevronsHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronsHorizontalInnerSVG,
		children,
	)
}

func ChevronsVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronsVerticalInnerSVG,
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

func Clock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clockInnerSVG,
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

func CloseBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeBoxInnerSVG,
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

func CloudDone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDoneInnerSVG,
		children,
	)
}

func CloudDownload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownloadInnerSVG,
		children,
	)
}

func CloudMoon(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudMoonInnerSVG,
		children,
	)
}

func CloudSun(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudSunInnerSVG,
		children,
	)
}

func CloudUpload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUploadInnerSVG,
		children,
	)
}

func Cocktail(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cocktailInnerSVG,
		children,
	)
}

func Code(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeInnerSVG,
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

func CoffeeAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coffeeAltInnerSVG,
		children,
	)
}

func Coin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		coinInnerSVG,
		children,
	)
}

func Collapse(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		collapseInnerSVG,
		children,
	)
}

func ColorsSwatch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		colorsSwatchInnerSVG,
		children,
	)
}

func Command(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commandInnerSVG,
		children,
	)
}

func Comment(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		commentInnerSVG,
		children,
	)
}

func Contact(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		contactInnerSVG,
		children,
	)
}

func ContactDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		contactDeleteInnerSVG,
		children,
	)
}

func ContactMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		contactMultipleInnerSVG,
		children,
	)
}

func ContactPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		contactPlusInnerSVG,
		children,
	)
}

func Copy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		copyInnerSVG,
		children,
	)
}

func CornerDownLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerDownLeftInnerSVG,
		children,
	)
}

func CornerDownRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerDownRightInnerSVG,
		children,
	)
}

func CornerLeftDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerLeftDownInnerSVG,
		children,
	)
}

func CornerLeftUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerLeftUpInnerSVG,
		children,
	)
}

func CornerRightDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerRightDownInnerSVG,
		children,
	)
}

func CornerRightUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerRightUpInnerSVG,
		children,
	)
}

func CornerUpLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerUpLeftInnerSVG,
		children,
	)
}

func CornerUpRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerUpRightInnerSVG,
		children,
	)
}

func CreditCard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardInnerSVG,
		children,
	)
}

func CreditCardDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardDeleteInnerSVG,
		children,
	)
}

func CreditCardMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardMinusInnerSVG,
		children,
	)
}

func CreditCardMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardMultipleInnerSVG,
		children,
	)
}

func CreditCardPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardPlusInnerSVG,
		children,
	)
}

func CreditCardSettings(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardSettingsInnerSVG,
		children,
	)
}

func CreditCardWireless(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardWirelessInnerSVG,
		children,
	)
}

func Crop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cropInnerSVG,
		children,
	)
}

func Cut(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cutInnerSVG,
		children,
	)
}

func Dashbaord(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dashbaordInnerSVG,
		children,
	)
}

func Debug(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		debugInnerSVG,
		children,
	)
}

func DebugCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		debugCheckInnerSVG,
		children,
	)
}

func DebugOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		debugOffInnerSVG,
		children,
	)
}

func DebugPause(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		debugPauseInnerSVG,
		children,
	)
}

func DebugPlay(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		debugPlayInnerSVG,
		children,
	)
}

func DebugStop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		debugStopInnerSVG,
		children,
	)
}

func Delete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deleteInnerSVG,
		children,
	)
}

func Deskphone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deskphoneInnerSVG,
		children,
	)
}

func DeviceLaptop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceLaptopInnerSVG,
		children,
	)
}

func DevicePhone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		devicePhoneInnerSVG,
		children,
	)
}

func DeviceTablet(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceTabletInnerSVG,
		children,
	)
}

func DeviceTv(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceTvInnerSVG,
		children,
	)
}

func DeviceTvSmart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceTvSmartInnerSVG,
		children,
	)
}

func DeviceVibrate(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceVibrateInnerSVG,
		children,
	)
}

func DeviceWatch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		deviceWatchInnerSVG,
		children,
	)
}

func Devices(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		devicesInnerSVG,
		children,
	)
}

func Dice(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diceInnerSVG,
		children,
	)
}

func Dollar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dollarInnerSVG,
		children,
	)
}

func Downasaur(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downasaurInnerSVG,
		children,
	)
}

func Download(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadInnerSVG,
		children,
	)
}

func Draft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		draftInnerSVG,
		children,
	)
}

func DragAndDrop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dragAndDropInnerSVG,
		children,
	)
}

func Drop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dropInnerSVG,
		children,
	)
}

func DropArea(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dropAreaInnerSVG,
		children,
	)
}

func DropFull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dropFullInnerSVG,
		children,
	)
}

func DropHalf(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dropHalfInnerSVG,
		children,
	)
}

func Duplicate(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		duplicateInnerSVG,
		children,
	)
}

func DuplicateAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		duplicateAltInnerSVG,
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

func EditBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editBoxInnerSVG,
		children,
	)
}

func Euro(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		euroInnerSVG,
		children,
	)
}

func Expand(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		expandInnerSVG,
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

func Eye(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeInnerSVG,
		children,
	)
}

func EyeClosed(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeClosedInnerSVG,
		children,
	)
}

func File(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileInnerSVG,
		children,
	)
}

func FileAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileAltInnerSVG,
		children,
	)
}

func FileDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileDeleteInnerSVG,
		children,
	)
}

func FileFlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileFlashInnerSVG,
		children,
	)
}

func FileMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileMinusInnerSVG,
		children,
	)
}

func FileMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileMultipleInnerSVG,
		children,
	)
}

func FileOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileOffInnerSVG,
		children,
	)
}

func FilePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filePlusInnerSVG,
		children,
	)
}

func Fill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fillInnerSVG,
		children,
	)
}

func FillHalf(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fillHalfInnerSVG,
		children,
	)
}

func FiveG(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fiveGInnerSVG,
		children,
	)
}

func Flag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flagInnerSVG,
		children,
	)
}

func Flatten(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flattenInnerSVG,
		children,
	)
}

func FlipToBack(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flipToBackInnerSVG,
		children,
	)
}

func FlipToFront(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flipToFrontInnerSVG,
		children,
	)
}

func FloatCenter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		floatCenterInnerSVG,
		children,
	)
}

func FloatLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		floatLeftInnerSVG,
		children,
	)
}

func FloatRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		floatRightInnerSVG,
		children,
	)
}

func Folder(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderInnerSVG,
		children,
	)
}

func FolderMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderMinusInnerSVG,
		children,
	)
}

func FolderPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderPlusInnerSVG,
		children,
	)
}

func FolderX(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderXInnerSVG,
		children,
	)
}

func Forward(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		forwardInnerSVG,
		children,
	)
}

func Forwardburger(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		forwardburgerInnerSVG,
		children,
	)
}

func FourG(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fourGInnerSVG,
		children,
	)
}

func FourK(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fourKInnerSVG,
		children,
	)
}

func FourKBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fourKBoxInnerSVG,
		children,
	)
}

func Frame(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		frameInnerSVG,
		children,
	)
}

func FrameAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		frameAddInnerSVG,
		children,
	)
}

func FrameCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		frameCheckInnerSVG,
		children,
	)
}

func FrameDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		frameDeleteInnerSVG,
		children,
	)
}

func FrameMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		frameMinusInnerSVG,
		children,
	)
}

func Gamepad(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gamepadInnerSVG,
		children,
	)
}

func Gif(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gifInnerSVG,
		children,
	)
}

func Gift(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		giftInnerSVG,
		children,
	)
}

func GitBranch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitBranchInnerSVG,
		children,
	)
}

func GitCommit(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitCommitInnerSVG,
		children,
	)
}

func GitMerge(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitMergeInnerSVG,
		children,
	)
}

func GitPullRequest(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gitPullRequestInnerSVG,
		children,
	)
}

func Gps(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gpsInnerSVG,
		children,
	)
}

func Grid(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gridInnerSVG,
		children,
	)
}

func Group(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		groupInnerSVG,
		children,
	)
}

func Hd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hdInnerSVG,
		children,
	)
}

func Headphone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		headphoneInnerSVG,
		children,
	)
}

func Headset(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		headsetInnerSVG,
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

func Hidden(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hiddenInnerSVG,
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

func Hourglass(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hourglassInnerSVG,
		children,
	)
}

func Hq(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hqInnerSVG,
		children,
	)
}

func Human(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		humanInnerSVG,
		children,
	)
}

func HumanHandsdown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		humanHandsdownInnerSVG,
		children,
	)
}

func HumanHandsup(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		humanHandsupInnerSVG,
		children,
	)
}

func HumanHeight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		humanHeightInnerSVG,
		children,
	)
}

func HumanHeightAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		humanHeightAltInnerSVG,
		children,
	)
}

func HumanRun(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		humanRunInnerSVG,
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

func ImageArrowRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageArrowRightInnerSVG,
		children,
	)
}

func ImageBroken(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageBrokenInnerSVG,
		children,
	)
}

func ImageDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageDeleteInnerSVG,
		children,
	)
}

func ImageFlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageFlashInnerSVG,
		children,
	)
}

func ImageFrame(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageFrameInnerSVG,
		children,
	)
}

func ImageGallery(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageGalleryInnerSVG,
		children,
	)
}

func ImageMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageMultipleInnerSVG,
		children,
	)
}

func ImageNew(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageNewInnerSVG,
		children,
	)
}

func ImagePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imagePlusInnerSVG,
		children,
	)
}

func Inbox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxInnerSVG,
		children,
	)
}

func InboxAll(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxAllInnerSVG,
		children,
	)
}

func InboxFull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxFullInnerSVG,
		children,
	)
}

func InfoBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		infoBoxInnerSVG,
		children,
	)
}

func Invert(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		invertInnerSVG,
		children,
	)
}

func Iso(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		isoInnerSVG,
		children,
	)
}

func Kanban(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		kanbanInnerSVG,
		children,
	)
}

func Keyboard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		keyboardInnerSVG,
		children,
	)
}

func Label(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		labelInnerSVG,
		children,
	)
}

func LabelAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		labelAltInnerSVG,
		children,
	)
}

func LabelAltMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		labelAltMultipleInnerSVG,
		children,
	)
}

func LabelSharp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		labelSharpInnerSVG,
		children,
	)
}

func Layout(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutInnerSVG,
		children,
	)
}

func LayoutAlignBottom(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutAlignBottomInnerSVG,
		children,
	)
}

func LayoutAlignLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutAlignLeftInnerSVG,
		children,
	)
}

func LayoutAlignRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutAlignRightInnerSVG,
		children,
	)
}

func LayoutAlignTop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutAlignTopInnerSVG,
		children,
	)
}

func LayoutColumns(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutColumnsInnerSVG,
		children,
	)
}

func LayoutDistributeHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutDistributeHorizontalInnerSVG,
		children,
	)
}

func LayoutDistributeVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutDistributeVerticalInnerSVG,
		children,
	)
}

func LayoutFooter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutFooterInnerSVG,
		children,
	)
}

func LayoutHeader(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutHeaderInnerSVG,
		children,
	)
}

func LayoutRows(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutRowsInnerSVG,
		children,
	)
}

func LayoutSidebarLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutSidebarLeftInnerSVG,
		children,
	)
}

func LayoutSidebarRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutSidebarRightInnerSVG,
		children,
	)
}

func Link(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkInnerSVG,
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

func ListBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listBoxInnerSVG,
		children,
	)
}

func Loader(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loaderInnerSVG,
		children,
	)
}

func Lock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockInnerSVG,
		children,
	)
}

func LockOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockOpenInnerSVG,
		children,
	)
}

func Login(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loginInnerSVG,
		children,
	)
}

func Logout(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logoutInnerSVG,
		children,
	)
}

func Luggage(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		luggageInnerSVG,
		children,
	)
}

func Mail(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailInnerSVG,
		children,
	)
}

func MailArrowRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailArrowRightInnerSVG,
		children,
	)
}

func MailCheck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailCheckInnerSVG,
		children,
	)
}

func MailDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailDeleteInnerSVG,
		children,
	)
}

func MailFlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailFlashInnerSVG,
		children,
	)
}

func MailMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailMultipleInnerSVG,
		children,
	)
}

func MailOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailOffInnerSVG,
		children,
	)
}

func MailUnread(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mailUnreadInnerSVG,
		children,
	)
}

func Map(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapInnerSVG,
		children,
	)
}

func Membercard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		membercardInnerSVG,
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

func Message(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageInnerSVG,
		children,
	)
}

func MessageArrowLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageArrowLeftInnerSVG,
		children,
	)
}

func MessageArrowRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageArrowRightInnerSVG,
		children,
	)
}

func MessageBookmark(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageBookmarkInnerSVG,
		children,
	)
}

func MessageClock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageClockInnerSVG,
		children,
	)
}

func MessageDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageDeleteInnerSVG,
		children,
	)
}

func MessageFlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageFlashInnerSVG,
		children,
	)
}

func MessageImage(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageImageInnerSVG,
		children,
	)
}

func MessageMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageMinusInnerSVG,
		children,
	)
}

func MessagePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messagePlusInnerSVG,
		children,
	)
}

func MessageProcessing(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageProcessingInnerSVG,
		children,
	)
}

func MessageReply(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageReplyInnerSVG,
		children,
	)
}

func MessageText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageTextInnerSVG,
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

func MissedCall(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		missedCallInnerSVG,
		children,
	)
}

func Modem(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		modemInnerSVG,
		children,
	)
}

func Money(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moneyInnerSVG,
		children,
	)
}

func Monitor(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		monitorInnerSVG,
		children,
	)
}

func MoodHappy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moodHappyInnerSVG,
		children,
	)
}

func MoodNeutral(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moodNeutralInnerSVG,
		children,
	)
}

func MoodSad(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moodSadInnerSVG,
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

func MoonStar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonStarInnerSVG,
		children,
	)
}

func MoonStars(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonStarsInnerSVG,
		children,
	)
}

func MoreHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreHorizontalInnerSVG,
		children,
	)
}

func MoreVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreVerticalInnerSVG,
		children,
	)
}

func Mouse(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mouseInnerSVG,
		children,
	)
}

func Move(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moveInnerSVG,
		children,
	)
}

func Movie(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		movieInnerSVG,
		children,
	)
}

func Music(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		musicInnerSVG,
		children,
	)
}

func Next(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		nextInnerSVG,
		children,
	)
}

func Note(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteInnerSVG,
		children,
	)
}

func NoteDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteDeleteInnerSVG,
		children,
	)
}

func NoteMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		noteMultipleInnerSVG,
		children,
	)
}

func NotePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		notePlusInnerSVG,
		children,
	)
}

func Notes(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		notesInnerSVG,
		children,
	)
}

func NotesDelete(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		notesDeleteInnerSVG,
		children,
	)
}

func NotesMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		notesMultipleInnerSVG,
		children,
	)
}

func NotesPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		notesPlusInnerSVG,
		children,
	)
}

func Notification(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		notificationInnerSVG,
		children,
	)
}

func NotificationOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		notificationOffInnerSVG,
		children,
	)
}

func Open(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		openInnerSVG,
		children,
	)
}

func PaintBucket(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paintBucketInnerSVG,
		children,
	)
}

func Paperclip(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperclipInnerSVG,
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

func Percent(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		percentInnerSVG,
		children,
	)
}

func PictureInPicture(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pictureInPictureInnerSVG,
		children,
	)
}

func PictureInPictureAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pictureInPictureAltInnerSVG,
		children,
	)
}

func Pin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pinInnerSVG,
		children,
	)
}

func Pixelarticons(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pixelarticonsInnerSVG,
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

func Playlist(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playlistInnerSVG,
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

func Power(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		powerInnerSVG,
		children,
	)
}

func Prev(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		prevInnerSVG,
		children,
	)
}

func Print(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		printInnerSVG,
		children,
	)
}

func RadioHandheld(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioHandheldInnerSVG,
		children,
	)
}

func RadioOn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioOnInnerSVG,
		children,
	)
}

func RadioSignal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioSignalInnerSVG,
		children,
	)
}

func RadioTower(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioTowerInnerSVG,
		children,
	)
}

func Reciept(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		recieptInnerSVG,
		children,
	)
}

func RecieptAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		recieptAltInnerSVG,
		children,
	)
}

func Redo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		redoInnerSVG,
		children,
	)
}

func Reload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		reloadInnerSVG,
		children,
	)
}

func RemoveBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeBoxInnerSVG,
		children,
	)
}

func RemoveBoxMultiple(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		removeBoxMultipleInnerSVG,
		children,
	)
}

func Repeat(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		repeatInnerSVG,
		children,
	)
}

func Reply(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		replyInnerSVG,
		children,
	)
}

func ReplyAll(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		replyAllInnerSVG,
		children,
	)
}

func RoundedCorner(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		roundedCornerInnerSVG,
		children,
	)
}

func Save(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveInnerSVG,
		children,
	)
}

func Scale(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scaleInnerSVG,
		children,
	)
}

func Script(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scriptInnerSVG,
		children,
	)
}

func ScriptText(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scriptTextInnerSVG,
		children,
	)
}

func ScrollHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scrollHorizontalInnerSVG,
		children,
	)
}

func ScrollVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scrollVerticalInnerSVG,
		children,
	)
}

func Sd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sdInnerSVG,
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

func Section(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sectionInnerSVG,
		children,
	)
}

func SectionCopy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sectionCopyInnerSVG,
		children,
	)
}

func SectionMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sectionMinusInnerSVG,
		children,
	)
}

func SectionPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sectionPlusInnerSVG,
		children,
	)
}

func SectionX(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sectionXInnerSVG,
		children,
	)
}

func Server(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		serverInnerSVG,
		children,
	)
}

func SharpCorner(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sharpCornerInnerSVG,
		children,
	)
}

func Shield(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldInnerSVG,
		children,
	)
}

func ShieldOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldOffInnerSVG,
		children,
	)
}

func Ship(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shipInnerSVG,
		children,
	)
}

func ShoppingBag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingBagInnerSVG,
		children,
	)
}

func Shuffle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shuffleInnerSVG,
		children,
	)
}

func Sliders(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		slidersInnerSVG,
		children,
	)
}

func SlidersTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		slidersTwoInnerSVG,
		children,
	)
}

func Sort(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sortInnerSVG,
		children,
	)
}

func SortAlpabetic(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sortAlpabeticInnerSVG,
		children,
	)
}

func SortNumeric(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sortNumericInnerSVG,
		children,
	)
}

func Speaker(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speakerInnerSVG,
		children,
	)
}

func SpeedFast(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speedFastInnerSVG,
		children,
	)
}

func SpeedMedium(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speedMediumInnerSVG,
		children,
	)
}

func SpeedSlow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speedSlowInnerSVG,
		children,
	)
}

func Spotlight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		spotlightInnerSVG,
		children,
	)
}

func Store(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		storeInnerSVG,
		children,
	)
}

func Subscriptions(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		subscriptionsInnerSVG,
		children,
	)
}

func Subtitles(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		subtitlesInnerSVG,
		children,
	)
}

func Suitcase(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		suitcaseInnerSVG,
		children,
	)
}

func Sun(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunInnerSVG,
		children,
	)
}

func SunAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunAltInnerSVG,
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

func Sync(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		syncInnerSVG,
		children,
	)
}

func Tab(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tabInnerSVG,
		children,
	)
}

func Table(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tableInnerSVG,
		children,
	)
}

func Tea(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		teaInnerSVG,
		children,
	)
}

func Teach(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		teachInnerSVG,
		children,
	)
}

func TextAdd(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textAddInnerSVG,
		children,
	)
}

func TextColums(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textColumsInnerSVG,
		children,
	)
}

func TextSearch(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textSearchInnerSVG,
		children,
	)
}

func TextWrap(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textWrapInnerSVG,
		children,
	)
}

func Timeline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		timelineInnerSVG,
		children,
	)
}

func ToggleLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		toggleLeftInnerSVG,
		children,
	)
}

func ToggleRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		toggleRightInnerSVG,
		children,
	)
}

func Tournament(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tournamentInnerSVG,
		children,
	)
}

func TrackChanges(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trackChangesInnerSVG,
		children,
	)
}

func Trash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trashInnerSVG,
		children,
	)
}

func TrashAlt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trashAltInnerSVG,
		children,
	)
}

func Trending(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingInnerSVG,
		children,
	)
}

func TrendingDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingDownInnerSVG,
		children,
	)
}

func TrendingUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingUpInnerSVG,
		children,
	)
}

func Trophy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trophyInnerSVG,
		children,
	)
}

func Truck(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		truckInnerSVG,
		children,
	)
}

func Undo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		undoInnerSVG,
		children,
	)
}

func Ungroup(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		ungroupInnerSVG,
		children,
	)
}

func Unlink(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		unlinkInnerSVG,
		children,
	)
}

func Upload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadInnerSVG,
		children,
	)
}

func User(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userInnerSVG,
		children,
	)
}

func UserMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userMinusInnerSVG,
		children,
	)
}

func UserPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userPlusInnerSVG,
		children,
	)
}

func UserX(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		userXInnerSVG,
		children,
	)
}

func Users(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		usersInnerSVG,
		children,
	)
}

func Video(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoInnerSVG,
		children,
	)
}

func VideoOff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoOffInnerSVG,
		children,
	)
}

func ViewCol(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewColInnerSVG,
		children,
	)
}

func ViewList(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewListInnerSVG,
		children,
	)
}

func ViewportNarrow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewportNarrowInnerSVG,
		children,
	)
}

func ViewportWide(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		viewportWideInnerSVG,
		children,
	)
}

func Visible(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		visibleInnerSVG,
		children,
	)
}

func Volume(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeInnerSVG,
		children,
	)
}

func VolumeMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeMinusInnerSVG,
		children,
	)
}

func VolumeOne(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeOneInnerSVG,
		children,
	)
}

func VolumePlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumePlusInnerSVG,
		children,
	)
}

func VolumeThree(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeThreeInnerSVG,
		children,
	)
}

func VolumeTwo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeTwoInnerSVG,
		children,
	)
}

func VolumeVibrate(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeVibrateInnerSVG,
		children,
	)
}

func VolumeX(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeXInnerSVG,
		children,
	)
}

func Wallet(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		walletInnerSVG,
		children,
	)
}

func WarningBox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		warningBoxInnerSVG,
		children,
	)
}

func Wind(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		windInnerSVG,
		children,
	)
}

func Zap(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		zapInnerSVG,
		children,
	)
}

func ZoomIn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		zoomInInnerSVG,
		children,
	)
}

func ZoomOut(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		zoomOutInnerSVG,
		children,
	)
}
