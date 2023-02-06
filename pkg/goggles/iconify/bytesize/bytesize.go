package bytesize

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	activityInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16h7l3 13l4-26l3 13h7"/>`
	alertInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 3l14 26H2Zm0 8v8m0 4v2"/>`
	archiveInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10v18h24V10M2 4v6h28V4Zm10 11h8"/>`
	arrowBottomInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 22l10 8l10-8m-10 8V2"/>`
	arrowLeftInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6L2 16l8 10M2 16h28"/>`
	arrowRightInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 6l8 10l-8 10m8-10H2"/>`
	arrowTopInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 10l10-8l10 8M16 2v28"/>`
	backwardsInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 2L2 16l14 14V16l14 14V2L16 16Z"/>`
	bagInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9v20h22V9Zm5 0s0-6 6-6s6 6 6 6"/>`
	banInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="16" cy="16" r="14"/><path d="m6 6l20 20"/></g>`
	bellInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 17c0-5 1-11 8-11s8 6 8 11s3 8 3 8H5s3-3 3-8Zm12 8s0 4-4 4s-4-4-4-4m4-22v3"/>`
	bookInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7S9 1 2 6v22c7-5 14 0 14 0s7-5 14 0V6c-7-5-14 1-14 1Zm0 0v21"/>`
	bookmarkInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 2h20v28L16 20L6 30Z"/>`
	calendarInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 6v24h28V6Zm0 9h28M7 3v6m6-6v6m6-6v6m6-6v6"/>`
	cameraInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 8h7l3-4h8l3 4h7v18H2Z"/><circle cx="16" cy="16" r="5"/></g>`
	caretBottomInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M30 10L16 26L2 10Z"/>`
	caretLeftInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 30L6 16L22 2Z"/>`
	caretRightInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 30l16-14L10 2Z"/>`
	caretTopInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M30 22L16 6L2 22Z"/>`
	cartInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 6h24l-3 13H9m18 4H10L5 2H2"/><circle cx="25" cy="27" r="2"/><circle cx="12" cy="27" r="2"/></g>`
	checkmarkInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 20l10 8L30 4"/>`
	chevronBottomInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M30 12L16 24L2 12"/>`
	chevronLeftInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 30L8 16L20 2"/>`
	chevronRightInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 30l12-14L12 2"/>`
	chevronTopInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M30 20L16 8L2 20"/>`
	clipboardInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v4h8V2h-8Zm-1 2H6v26h20V4h-5"/>`
	clockInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="16" cy="16" r="14"/><path d="M16 8v8l4 4"/></g>`
	closeInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 30L30 2m0 28L2 2"/>`
	codeInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 9l-7 8l7 8M22 9l7 8l-7 8M18 7l-4 20"/>`
	composeInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27 15v15H2V5h15m13 1l-4-4L9 19l-2 6l6-2Zm-8 0l4 4ZM9 19l4 4Z"/>`
	creditcardInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 7v18h28V7Zm3 11h4m-4 3h6"/><path fill="currentColor" d="M2 11v2h28v-2Z"/></g>`
	desktopInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 29s0-5 6-5s6 5 6 5H10ZM2 6v17h28V6H2Z"/>`
	downloadInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 22c-9 1-8-10 0-9C6 2 23 2 22 10c10-3 10 13 1 12m-12 4l5 4l5-4m-5-10v14"/>`
	editInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m30 7l-5-5L5 22l-2 7l7-2Zm-9-1l5 5ZM5 22l5 5Z"/>`
	ejectInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M30 18L16 5L2 18ZM2 25h28"/>`
	ellipsisHorizontalInnerSVG = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="16" r="2"/><circle cx="16" cy="16" r="2"/><circle cx="25" cy="16" r="2"/></g>`
	ellipsisVerticalInnerSVG   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="16" cy="7" r="2"/><circle cx="16" cy="16" r="2"/><circle cx="16" cy="25" r="2"/></g>`
	endInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 2v14L10 2v28l14-14v14"/>`
	exportInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M28 22v8H4v-8M16 4v20M8 12l8-8l8 8"/>`
	externalInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 9H3v20h20V18M18 4h10v10m0-10L14 18"/>`
	eyeInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="17" cy="15" r="1"/><circle cx="16" cy="16" r="6"/><path d="M2 16S7 6 16 6s14 10 14 10s-5 10-14 10S2 16 2 16Z"/></g>`
	feedInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="26" r="2" fill="currentColor"/><path d="M4 15c7 0 13 6 13 13M4 6c13 0 22 9 22 22"/></g>`
	fileInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 2v28h20V10l-8-8Zm12 0v8h8"/>`
	filterInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 5s4-2 14-2s14 2 14 2L19 18v9l-6 3V18L2 5Z"/>`
	fireInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 30c-15 0 0-17-3-27 16 10 20 27 3 27zm2 0c-7 0 0-10 0-10s7 10 0 10z"/></g>`
	flagInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 2v28M6 6h20l-6 6l6 6H6"/>`
	folderInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 26h28V7H14l-4-3H2Zm28-14H2"/>`
	folderOpenInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 28h24l2-16H14l-4-4H2Zm24-16V4H4v4"/>`
	forwardsInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 2l14 14l-14 14V16L2 30V2l14 14Z"/>`
	fullscreenInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12V4h8m8 0h8v8M4 20v8h8m16-8v8h-8"/>`
	fullscreenExitInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8V4m8 0v8h8M4 20h8v8m16-8h-8v8"/>`
	giftInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14v16h24V14M2 9v5h28V9Zm14 0s-2-9-8-6s8 6 8 6s2-9 8-6s-8 6-8 6m0 0v21"/>`
	githubInnerSVG             = `<path fill="currentColor" d="M32 0C14 0 0 14 0 32c0 21 19 30 22 30c2 0 2-1 2-2v-5c-7 2-10-2-11-5c0 0 0-1-2-3c-1-1-5-3-1-3c3 0 5 4 5 4c3 4 7 3 9 2c0-2 2-4 2-4c-8-1-14-4-14-15c0-4 1-7 3-9c0 0-2-4 0-9c0 0 5 0 9 4c3-2 13-2 16 0c4-4 9-4 9-4c2 7 0 9 0 9c2 2 3 5 3 9c0 11-7 14-14 15c1 1 2 3 2 6v8c0 1 0 2 2 2c3 0 22-9 22-30C64 14 50 0 32 0Z"/>`
	heartInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16C1 12 2 6 7 4s8 2 9 4c1-2 5-6 10-4s5 8 2 12s-12 12-12 12s-9-8-12-12Z"/>`
	homeInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20v10H4V12L16 2l12 10v18h-8V20Z"/>`
	importInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M28 22v8H4v-8M16 4v20m-8-8l8 8l8-8"/>`
	inboxInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 15v10h28V15l-4-8H6Zm0 0h8s1 5 6 5s6-5 6-5h8"/>`
	infoInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 14v9m0-15v2"/><circle cx="16" cy="16" r="14"/></g>`
	lightningInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 13l8-11L8 13l6 6l-8 11l18-11Z"/>`
	linkInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 8s6-6 9-3s2 7-3 11s-8 5-10 1m0 7s-6 6-9 3s-2-7 3-11s8-5 10-1"/>`
	locationInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="16" cy="11" r="4"/><path d="M24 15c-3 7-8 15-8 15s-5-8-8-15s2-13 8-13s11 6 8 13Z"/></g>`
	lockInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 15v15h22V15Zm4 0C9 9 9 5 16 5s7 4 7 10m-7 5v3"/><circle cx="16" cy="24" r="1"/></g>`
	mailInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 26h28V6H2ZM2 6l14 10L30 6"/>`
	menuInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h24M4 16h24M4 24h24"/>`
	messageInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 4h28v18H16l-8 7v-7H2Z"/>`
	microphoneInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 2c-4 0-4 4-4 4v10s0 4 4 4s4-4 4-4V6s0-4-4-4ZM8 17s0 7 8 7s8-7 8-7M13 29h6m-3-5v5"/>`
	minusInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 16h28"/>`
	mobileInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 2H11c-1 0-2 1-2 2v24c0 1 1 2 2 2h10c1 0 2-1 2-2V4c0-1-1-2-2-2ZM9 5h14M9 27h14"/>`
	moonInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 2C9 2 3 7 3 15s6 14 14 14s13-6 13-11C19 25 7 13 14 2Z"/>`
	moveInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16h26M16 3v26M12 7l4-4l4 4m-8 18l4 4l4-4m5-13l4 4l-4 4M7 12l-4 4l4 4"/>`
	musicInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 25V6l13-3v20M11 13l13-3"/><ellipse cx="7" cy="25" rx="4" ry="5"/><ellipse cx="20" cy="23" rx="4" ry="5"/></g>`
	muteInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 16c0-8-5-14-5-14l-7 8H2v12h6l7 8s5-6 5-14Z"/>`
	optionsInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M28 6H4m24 10H4m24 10H4M24 3v6M8 13v6m12 4v6"/>`
	paperclipInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v15c0 4 3 6 6 6s6-2 6-6V6c0-3-2-4-4-4s-4 1-4 4v17c0 1 1 2 2 2s2-1 2-2V9"/>`
	pauseInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 2v28M9 2v28"/>`
	photoInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20 24l-8-8L2 26V2h28v22m-14-4l6-6l8 8v8H2v-6"/><circle cx="10" cy="9" r="3"/></g>`
	playInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 2v28l14-14Z"/>`
	plusInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 2v28M2 16h28"/>`
	portfolioInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 17v11H3V17M2 8h28v8s-6 4-14 4s-14-4-14-4V8Zm14 14v-4m4-10s0-4-4-4s-4 4-4 4"/>`
	printInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 25H2V9h28v16h-5M7 19v11h18V19ZM25 9V2H7v7m15 5h3"/>`
	reloadInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 16c0 6-5 13-13 13S3 22 3 16S8 3 16 3c5 0 9 3 11 6m-7 1l7-1l1-7"/>`
	replyInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 6l-7 8l7 8m-7-8h15c8 0 12 4 12 12"/>`
	searchInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="14" cy="14" r="12"/><path d="m23 23l7 7"/></g>`
	sendInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 16L30 2L16 30l-4-10ZM30 2L12 20"/>`
	settingsInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 2v4l-2 1l-3-3l-4 4l3 3l-1 2H2v6h4l1 2l-3 3l4 4l3-3l2 1v4h6v-4l2-1l3 3l4-4l-3-3l1-2h4v-6h-4l-1-2l3-3l-4-4l-3 3l-2-1V2Z"/><circle cx="16" cy="16" r="4"/></g>`
	signInInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16h20m-8-8l8 8l-8 8m6-20h8v24h-8"/>`
	signOutInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M28 16H8m12-8l8 8l-8 8m-9 4H3V4h8"/>`
	starInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 2l4 10h10l-8 7l3 11l-9-7l-9 7l3-11l-8-7h10Z"/>`
	startInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 2v14L22 2v28L8 16v14"/>`
	tagInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="24" cy="8" r="2"/><path d="M2 18L18 2h12v12L14 30Z"/></g>`
	telephoneInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12c0-7 7-7 13-7s13 0 13 7c0 8-7-1-7-1H10s-7 9-7 1Zm8 2s-5 5-5 14h20c0-9-5-14-5-14H11Z"/><circle cx="16" cy="21" r="4"/></g>`
	trashInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M28 6H6l2 24h16l2-24H4m12 6v12m5-12l-1 12m-9-12l1 12m0-18l1-4h6l1 4"/>`
	twitterInnerSVG            = `<path fill="currentColor" d="m60 16l-6 1l4-5l-7 2c-9-10-23 1-19 10C16 24 8 12 8 12s-6 9 4 16l-6-2c0 6 4 10 11 12h-7c4 8 11 8 11 8s-6 5-17 5c33 16 53-14 50-30Z"/>`
	unlockInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 15v15h22V15Zm4 0C9 7 9 3 16 3s7 5 7 6m-7 11v3"/><circle cx="16" cy="24" r="1"/></g>`
	uploadInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 22c-9 1-8-10 0-9C6 2 23 2 22 10c10-3 10 13 1 12m-12-4l5-4l5 4m-5-4v15"/>`
	userInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 11c0 5-3 9-6 9s-6-4-6-9s2-8 6-8s6 3 6 8ZM4 30h24c0-9-6-10-12-10S4 21 4 30Z"/>`
	videoInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 13l8-5v16l-8-5ZM2 8v16h20V8Z"/>`
	volumeInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 16c0-8-5-14-5-14l-7 8H2v12h6l7 8s5-6 5-14Zm1-14s4 4 4 14s-4 14-4 14m6-26s3 4 3 12s-3 12-3 12"/>`
	workInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M30 8H2v18h28ZM20 8s0-4-4-4s-4 4-4 4M8 26V8m16 18V8"/>`
	zoomInInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="14" cy="14" r="12"/><path d="m23 23l7 7M14 10v8m-4-4h8"/></g>`
	zoomOutInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="14" cy="14" r="12"/><path d="m23 23l7 7M10 14h8"/></g>`
	zoomResetInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="14" cy="14" r="12"/><path d="m23 23l7 7M9 12V9h3m4 0h3v3M9 16v3h3m7-3v3h-3"/></g>`
)

var sharedIconAttrs = engine.Attrs{"width": "1em", "height": "1em"}

func Activity(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		activityInnerSVG,
		children,
	)
}

func Alert(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		alertInnerSVG,
		children,
	)
}

func Archive(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		archiveInnerSVG,
		children,
	)
}

func ArrowBottom(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		arrowBottomInnerSVG,
		children,
	)
}

func ArrowLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		arrowLeftInnerSVG,
		children,
	)
}

func ArrowRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		arrowRightInnerSVG,
		children,
	)
}

func ArrowTop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		arrowTopInnerSVG,
		children,
	)
}

func Backwards(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		backwardsInnerSVG,
		children,
	)
}

func Bag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		bagInnerSVG,
		children,
	)
}

func Ban(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		banInnerSVG,
		children,
	)
}

func Bell(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		bellInnerSVG,
		children,
	)
}

func Book(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		bookInnerSVG,
		children,
	)
}

func Bookmark(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		bookmarkInnerSVG,
		children,
	)
}

func Calendar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		calendarInnerSVG,
		children,
	)
}

func Camera(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		cameraInnerSVG,
		children,
	)
}

func CaretBottom(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		caretBottomInnerSVG,
		children,
	)
}

func CaretLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		caretLeftInnerSVG,
		children,
	)
}

func CaretRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		caretRightInnerSVG,
		children,
	)
}

func CaretTop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		caretTopInnerSVG,
		children,
	)
}

func Cart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		cartInnerSVG,
		children,
	)
}

func Checkmark(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		checkmarkInnerSVG,
		children,
	)
}

func ChevronBottom(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		chevronBottomInnerSVG,
		children,
	)
}

func ChevronLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
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
			"viewBox": "0 0 32 32",
		},
		chevronRightInnerSVG,
		children,
	)
}

func ChevronTop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		chevronTopInnerSVG,
		children,
	)
}

func Clipboard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
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
			"viewBox": "0 0 32 32",
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
			"viewBox": "0 0 32 32",
		},
		closeInnerSVG,
		children,
	)
}

func Code(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		codeInnerSVG,
		children,
	)
}

func Compose(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		composeInnerSVG,
		children,
	)
}

func Creditcard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		creditcardInnerSVG,
		children,
	)
}

func Desktop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		desktopInnerSVG,
		children,
	)
}

func Download(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		downloadInnerSVG,
		children,
	)
}

func Edit(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		editInnerSVG,
		children,
	)
}

func Eject(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		ejectInnerSVG,
		children,
	)
}

func EllipsisHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		ellipsisHorizontalInnerSVG,
		children,
	)
}

func EllipsisVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		ellipsisVerticalInnerSVG,
		children,
	)
}

func End(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		endInnerSVG,
		children,
	)
}

func Export(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		exportInnerSVG,
		children,
	)
}

func External(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		externalInnerSVG,
		children,
	)
}

func Eye(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		eyeInnerSVG,
		children,
	)
}

func Feed(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		feedInnerSVG,
		children,
	)
}

func File(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		fileInnerSVG,
		children,
	)
}

func Filter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		filterInnerSVG,
		children,
	)
}

func Fire(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		fireInnerSVG,
		children,
	)
}

func Flag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		flagInnerSVG,
		children,
	)
}

func Folder(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		folderInnerSVG,
		children,
	)
}

func FolderOpen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		folderOpenInnerSVG,
		children,
	)
}

func Forwards(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		forwardsInnerSVG,
		children,
	)
}

func Fullscreen(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		fullscreenInnerSVG,
		children,
	)
}

func FullscreenExit(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		fullscreenExitInnerSVG,
		children,
	)
}

func Gift(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		giftInnerSVG,
		children,
	)
}

func Github(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		githubInnerSVG,
		children,
	)
}

func Heart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		heartInnerSVG,
		children,
	)
}

func Home(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		homeInnerSVG,
		children,
	)
}

func Import(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		importInnerSVG,
		children,
	)
}

func Inbox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		inboxInnerSVG,
		children,
	)
}

func Info(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		infoInnerSVG,
		children,
	)
}

func Lightning(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		lightningInnerSVG,
		children,
	)
}

func Link(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		linkInnerSVG,
		children,
	)
}

func Location(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		locationInnerSVG,
		children,
	)
}

func Lock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		lockInnerSVG,
		children,
	)
}

func Mail(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		mailInnerSVG,
		children,
	)
}

func Menu(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
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
			"viewBox": "0 0 32 32",
		},
		messageInnerSVG,
		children,
	)
}

func Microphone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		microphoneInnerSVG,
		children,
	)
}

func Minus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		minusInnerSVG,
		children,
	)
}

func Mobile(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		mobileInnerSVG,
		children,
	)
}

func Moon(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		moonInnerSVG,
		children,
	)
}

func Move(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		moveInnerSVG,
		children,
	)
}

func Music(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		musicInnerSVG,
		children,
	)
}

func Mute(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		muteInnerSVG,
		children,
	)
}

func Options(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		optionsInnerSVG,
		children,
	)
}

func Paperclip(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
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
			"viewBox": "0 0 32 32",
		},
		pauseInnerSVG,
		children,
	)
}

func Photo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		photoInnerSVG,
		children,
	)
}

func Play(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		playInnerSVG,
		children,
	)
}

func Plus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		plusInnerSVG,
		children,
	)
}

func Portfolio(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		portfolioInnerSVG,
		children,
	)
}

func Print(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		printInnerSVG,
		children,
	)
}

func Reload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		reloadInnerSVG,
		children,
	)
}

func Reply(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		replyInnerSVG,
		children,
	)
}

func Search(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		searchInnerSVG,
		children,
	)
}

func Send(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		sendInnerSVG,
		children,
	)
}

func Settings(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		settingsInnerSVG,
		children,
	)
}

func SignIn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		signInInnerSVG,
		children,
	)
}

func SignOut(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		signOutInnerSVG,
		children,
	)
}

func Star(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		starInnerSVG,
		children,
	)
}

func Start(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		startInnerSVG,
		children,
	)
}

func Tag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		tagInnerSVG,
		children,
	)
}

func Telephone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		telephoneInnerSVG,
		children,
	)
}

func Trash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		trashInnerSVG,
		children,
	)
}

func Twitter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		twitterInnerSVG,
		children,
	)
}

func Unlock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		unlockInnerSVG,
		children,
	)
}

func Upload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
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
			"viewBox": "0 0 32 32",
		},
		userInnerSVG,
		children,
	)
}

func Video(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		videoInnerSVG,
		children,
	)
}

func Volume(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		volumeInnerSVG,
		children,
	)
}

func Work(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		workInnerSVG,
		children,
	)
}

func ZoomIn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
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
			"viewBox": "0 0 32 32",
		},
		zoomOutInnerSVG,
		children,
	)
}

func ZoomReset(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 32 32",
		},
		zoomResetInnerSVG,
		children,
	)
}

func ByName(name string) (*engine.HTMLElement, error) {
	switch name {
	case "activity":
		return Activity(), nil
	case "alert":
		return Alert(), nil
	case "archive":
		return Archive(), nil
	case "arrow-bottom":
		return ArrowBottom(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-top":
		return ArrowTop(), nil
	case "backwards":
		return Backwards(), nil
	case "bag":
		return Bag(), nil
	case "ban":
		return Ban(), nil
	case "bell":
		return Bell(), nil
	case "book":
		return Book(), nil
	case "bookmark":
		return Bookmark(), nil
	case "calendar":
		return Calendar(), nil
	case "camera":
		return Camera(), nil
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
	case "checkmark":
		return Checkmark(), nil
	case "chevron-bottom":
		return ChevronBottom(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-top":
		return ChevronTop(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clock":
		return Clock(), nil
	case "close":
		return Close(), nil
	case "code":
		return Code(), nil
	case "compose":
		return Compose(), nil
	case "creditcard":
		return Creditcard(), nil
	case "desktop":
		return Desktop(), nil
	case "download":
		return Download(), nil
	case "edit":
		return Edit(), nil
	case "eject":
		return Eject(), nil
	case "ellipsis-horizontal":
		return EllipsisHorizontal(), nil
	case "ellipsis-vertical":
		return EllipsisVertical(), nil
	case "end":
		return End(), nil
	case "export":
		return Export(), nil
	case "external":
		return External(), nil
	case "eye":
		return Eye(), nil
	case "feed":
		return Feed(), nil
	case "file":
		return File(), nil
	case "filter":
		return Filter(), nil
	case "fire":
		return Fire(), nil
	case "flag":
		return Flag(), nil
	case "folder":
		return Folder(), nil
	case "folder-open":
		return FolderOpen(), nil
	case "forwards":
		return Forwards(), nil
	case "fullscreen":
		return Fullscreen(), nil
	case "fullscreen-exit":
		return FullscreenExit(), nil
	case "gift":
		return Gift(), nil
	case "github":
		return Github(), nil
	case "heart":
		return Heart(), nil
	case "home":
		return Home(), nil
	case "import":
		return Import(), nil
	case "inbox":
		return Inbox(), nil
	case "info":
		return Info(), nil
	case "lightning":
		return Lightning(), nil
	case "link":
		return Link(), nil
	case "location":
		return Location(), nil
	case "lock":
		return Lock(), nil
	case "mail":
		return Mail(), nil
	case "menu":
		return Menu(), nil
	case "message":
		return Message(), nil
	case "microphone":
		return Microphone(), nil
	case "minus":
		return Minus(), nil
	case "mobile":
		return Mobile(), nil
	case "moon":
		return Moon(), nil
	case "move":
		return Move(), nil
	case "music":
		return Music(), nil
	case "mute":
		return Mute(), nil
	case "options":
		return Options(), nil
	case "paperclip":
		return Paperclip(), nil
	case "pause":
		return Pause(), nil
	case "photo":
		return Photo(), nil
	case "play":
		return Play(), nil
	case "plus":
		return Plus(), nil
	case "portfolio":
		return Portfolio(), nil
	case "print":
		return Print(), nil
	case "reload":
		return Reload(), nil
	case "reply":
		return Reply(), nil
	case "search":
		return Search(), nil
	case "send":
		return Send(), nil
	case "settings":
		return Settings(), nil
	case "sign-in":
		return SignIn(), nil
	case "sign-out":
		return SignOut(), nil
	case "star":
		return Star(), nil
	case "start":
		return Start(), nil
	case "tag":
		return Tag(), nil
	case "telephone":
		return Telephone(), nil
	case "trash":
		return Trash(), nil
	case "twitter":
		return Twitter(), nil
	case "unlock":
		return Unlock(), nil
	case "upload":
		return Upload(), nil
	case "user":
		return User(), nil
	case "video":
		return Video(), nil
	case "volume":
		return Volume(), nil
	case "work":
		return Work(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	case "zoom-reset":
		return ZoomReset(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in bytesize icon set", name)
	}
}
