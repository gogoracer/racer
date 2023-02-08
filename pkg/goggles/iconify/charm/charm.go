package charm

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	anchorInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 5.75V14M3.25 7.75h-1.5c0 4 2.5 6.5 6.25 6.5s6.25-2.5 6.25-6.5h-1.5"/><circle cx="8" cy="3.5" r="1.75"/></g>`
	appsInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm0-8h4.5v4.5h-4.5z"/>`
	appsMinusInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm5.05-6h-5"/>`
	appsPlusInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm5.05-6h-5m2.5-2.5v5"/>`
	archiveInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v3.5H1.75zm5 6.5h2.5m-6.5-2.5v7.5h10.5v-7.5"/>`
	arrowDownInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.25 8.75l4.5 4.5l4.5-4.5m-4.5-6v10.5"/>`
	arrowDownLeftInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 11.75h-6.5v-6.5m7.5-1l-7.5 7.5"/>`
	arrowDownRightInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 11.75h6.5v-6.5m-7.5-1l7.5 7.5"/>`
	arrowLeftInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.25 3.75l-4.5 4.5l4.5 4.5m6-4.5H2.75"/>`
	arrowRightInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.75 3.25l4.5 4.5l-4.5 4.5m-6-4.5h10.5"/>`
	arrowUpInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.75 7.25l4.5-4.5l4.5 4.5m-4.5 6V2.75"/>`
	arrowUpLeftInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 4.25h-6.5v6.5m7.5 1l-7.5-7.5"/>`
	arrowUpRightInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 4.25h6.5v6.5m-7.5 1l7.5-7.5"/>`
	atSignInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.25 8c0 3.25 4 3.25 4 0A6.25 6.25 0 1 0 8 14.25c2.25 0 3.25-1 3.25-1"/><circle cx="8" cy="8" r="2.25"/></g>`
	atomInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><ellipse cx="11.3" rx="8.28" ry="3.17" transform="rotate(45)"/><ellipse cy="11.3" rx="8.28" ry="3.17" transform="rotate(315)"/><path d="M8 8v0"/></g>`
	bellInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 1.75c-2.468 0-4.25 1.5-4.25 3.5v3l-2 3.5h12.5l-2-3.5v-3c0-2-1.166-3.5-4.25-3.5zm-2.25 10.5c0 3 4.5 3 4.5 0"/>`
	bellSlashInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.75 12.25c0 3 4.5 3 4.5 0m2-4v-3c0-2-1.166-3.5-4.25-3.5m-3.75 2c-.53.585-.5.674-.5 1.5v3l-2 3.5h8.5"/><path d="M5.75 12.25c0 3 4.5 3 4.5 0m-7.5-10.5l10.5 12.5"/></g>`
	binInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 4.25v-2.5h4.5v2.5m-6.5 1v9h8.5v-9m-9.5-.5h10.5"/>`
	binaryInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 1.75h3v4.5h-3zm6.5 4.5h3m-3-4.5h1.5v4m-1.5 4h3v4.5h-3zm-6.5 4.5h3m-3-4.5h1.5v4"/>`
	blockInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="m4.25 11.75l8-8"/></g>`
	bluetoothInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 11.25L12.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25"/>`
	bluetoothConnectedInnerSVG  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 11.25L12.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25M1.75 8h1.5m9.5 0h1.5"/>`
	bluetoothSearchingInnerSVG  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.25L10.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25m11.5 1.5s1 .5 1 1.75s-1 1.75-1 1.75m-2-1.75v0"/>`
	bluetoothSlashInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 6.25L12.25 5l-4.5-3.25v2.5m4.5 6.75l-4.5 3.25v-6l-4 3m-2-8l12.5 9"/>`
	bookInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.75 11.75v2m1.5.5h-9c-.75 0-1.5-.5-1.5-1.5s.75-1.5 1.5-1.5h9v-9.5h-9c-.75 0-1.5.75-1.5 1.5v9.5"/>`
	bookOpenInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 3.75c-1.75-1-2.25-1-6.25-1v9.5c4 0 4.5 0 6.25 1c1.75-1 3.25-1 6.25-1v-9.5c-4 0-4.5 0-6.25 1zm0 .5v8.5"/>`
	bookmarkInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 1.75h8.5v12.5L8 9.75l-4.25 4.5z"/>`
	briefcaseInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75h12.5v9.5H1.75z"/><path d="M1.75 6.25s-.5 3.5 3 3.5h6.5c3.5 0 3-3.5 3-3.5m-8.5-2v-2.5h4.5v2.5"/></g>`
	bugInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="10" r="4.25"/><path d="M14.25 10.25h-1.5m-1 2.5l1.5 1.5m0-8.5l-1.5 1.5m-10 3h1.5m1 2.5l-1.5 1.5m0-8.5l1.5 1.5m1.5-1.5s-.75-3 2.25-3s2.25 3 2.25 3"/></g>`
	calendarInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v10.5H1.75zm9.5-2v1.5m-6.5-1.5v1.5m-2.5 4h11.5"/>`
	cameraInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75v8.5h12.5v-8.5h-3l-1.5-2h-3.5l-1.5 2z"/><circle cx="8" cy="8.5" r="2.25"/></g>`
	cameraVideoInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75h7.5v7.5h-7.5zm8 2.5l4.5-2.5v7.5l-4.5-2.5"/>`
	cameraVideoSlashInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.25 10.75l3 1.5v-7.5l-5 2.5v-2.5h-2.5m1.5 7.5h-6.5v-7.5h1.5m-1.5-2.5l8.5 12"/>`
	candyInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="3.25"/><path d="M7.25 11.25c0 1-.5 2.5-1.5 3c-.75 0-1.5-1-2-2c-1-.5-2-1.5-2-2c.5-1 2-1.5 3-1.5m4-4c0-1 .5-2.5 1.5-3c.75 0 1.5 1 2 2c1 .5 2 1.5 2 2c-.5 1-2 1.5-3 1.5"/></g>`
	cardsInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75H10v11.5H1.75zm8.25 1l4.25 2l-4.25 7.5"/>`
	castInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.25v-2.5h12.5v10.5h-4.5m-8-5a5 5 0 0 1 5 5m-5-2.5a2.5 2.5 0 0 1 2.5 2.5m-2.5 0v0"/>`
	certificateInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.25 1.75h-8.5v11.5h2.5m3.5-3.5l-.5 4.5l2.25-1l2.25 1l-.5-4.5"/><circle cx="10.5" cy="7.5" r="2.75"/></g>`
	chartBarInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75v12.5h12.5m-9-3v-2.5m4 2.5v-5.5m4 5.5v-8.5"/>`
	chartLineInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 11.25l2.5-4.5l2.5 2.5l3.5-6m-11.5-1.5v12.5h12.5"/>`
	chevronDownInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 5.75L8 10.25l4.25-4.5"/>`
	chevronLeftInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 3.75L5.75 8l4.5 4.25"/>`
	chevronRightInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 12.25L10.25 8l-4.5-4.25"/>`
	chevronUpInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 10.25L8 5.75l-4.25 4.5"/>`
	chevronsDownInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 3.75L8 8.25l4.25-4.5m-8.5 5L8 13.25l4.25-4.5"/>`
	chevronsLeftInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 3.75L7.75 8l4.5 4.25m-5-8.5L2.75 8l4.5 4.25"/>`
	chevronsRightInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 12.25L8.25 8l-4.5-4.25m5 8.5L13.25 8l-4.5-4.25"/>`
	chevronsUpInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 12.25L8 7.75l-4.25 4.5m8.5-5L8 2.75l-4.25 4.5"/>`
	chevronsUpDownInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 10.75L8 14.25l-3.25-3.5m6.5-5.5L8 1.75l-3.25 3.5"/>`
	chipInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 2.75h10.5v10.5H2.75z"/><path d="M6.25 6.25h3.5v3.5h-3.5zm-4 4h-1m1-4.5h-1m13.5 4.5h-1m1-4.5h-1m-3.5 8v1m-4.5-1v1m4.5-13.5v1m-4.5-1v1"/></g>`
	circleInnerSVG              = `<circle cx="8" cy="8" r="6.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/>`
	circleCrossInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.25 5.75l-4.5 4.5m0-4.5l4.5 4.5"/><circle cx="8" cy="8" r="6.25"/></g>`
	circleMinusInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M4.75 8h6.5"/></g>`
	circleTickInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 8.75c-.5 2.5-2.385 4.854-5.03 5.38A6.25 6.25 0 0 1 3.373 3.798C5.187 1.8 8.25 1.25 10.75 2.25"/><path d="m5.75 7.75l2.5 2.5l6-6.5"/></g>`
	circleWarningInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8 10.75v0m0-6v3.5"/></g>`
	clipboardInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.75 1.75h4.5v3.5h-4.5z"/><path d="M5.25 2.75h-2.5v11.5h10.5V2.75h-2.5"/></g>`
	clipboardTickInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 1.75h4.5v3.5h-4.5zm4 11.05l1.5 1.5l3-2.5m-9-9h-2.5v11.5h4.5m6-5V2.8h-2.5"/>`
	clockInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8.25 4.75v3.5l-2.5 2"/></g>`
	clockAlarmInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.75 1.75l2.5 2m-10-2l-2.5 2m10.5 9.5l1 1m-9.5-1l-1 1m5.5-7.5v2.5l-1.5 1"/><circle cx="8" cy="9" r="5.25"/></g>`
	cloudInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 3.75A3.25 3.25 0 0 0 3.75 7c.002.255.033.508.094.756h-.002A2.25 2.25 0 0 0 4 12.25h7.5a2.75 2.75 0 1 0-1.252-5.2L10.25 7A3.25 3.25 0 0 0 7 3.75z"/>`
	cloverInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 2.75C4.25 4.25 6 6 8 8c2-2 3.75-3.75 3.25-5.25s-2.5-1-3.25.5c-.75-1.5-2.75-2-3.25-.5zM8 8c2 2 3.75 3.75 5.25 3.25s1-2.5-.5-3.25c1.5-.75 2-2.75.5-3.25S10 6 8 8zm0 0c-2 2-3.75 3.75-3.25 5.25s2.5 1 3.25-.5c.75 1.5 2.75 2 3.25.5S10 10 8 8zm0 0C6 6 4.25 4.25 2.75 4.75s-1 2.5.5 3.25c-1.5.75-2 2.75-.5 3.25S6 10 8 8z"/>`
	codeInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 11.25L1.75 8l3.5-3.25m5.5 6.5L14.25 8l-3.5-3.25"/>`
	coffeeInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 11.25c4.5 0 4.5-5.5 0-5.5h-9v5c0 5 8.5 5 8.5 0v-5m-1.5-4v1.5m-3-1.5v1.5m-3-1.5v1.5"/>`
	cogInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="1.75"/><path d="m6.75 1.75l-.5 1.5l-1.5 1l-2-.5l-1 2l1.5 1.5v1.5l-1.5 1.5l1 2l2-.5l1.5 1l.5 1.5h2.5l.5-1.5l1.5-1l2 .5l1-2l-1.5-1.5v-1.5l1.5-1.5l-1-2l-2 .5l-1.5-1l-.5-1.5z"/></g>`
	compassInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="m6.75 6.75l-1 4l3.5-1.5l1-4z"/></g>`
	conicalFlaskInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 1.75h6.5m-6.5 8h6.5m-5.5-7.5v4.5l-4 7.5h12.5l-4-7.5v-4.5"/>`
	containerInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.75 12.2l5.5 2l7-4.5v-6l-5.5-2l-7 4.5z"/><path d="M10.8 6.25v5.5m-3.5-3.5v6m-5.5-8l5.5 2l7-4.5"/></g>`
	copyInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 4.25v-2.5h-9.5v9.5h2.5m.5-6.5v9.5h9.5v-9.5z"/>`
	copyleftInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M6 6.75s.75-1 2-1s2.25 1 2.25 2.25s-1 2.25-2.25 2.25s-2-1-2-1"/></g>`
	copyrightInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M10 6.75s-.75-1-2-1s-2.25 1-2.25 2.25s1 2.25 2.25 2.25s2-1 2-1"/></g>`
	creditCardInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v9.5H1.75zm8 6.5h1.5m-9-3h11.5"/>`
	cropInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.25 1.75v10h10M11.8 14.2v-2.5m0-2.5v-5h-5m-2.5 0H1.8"/>`
	crossInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.25 4.75l-6.5 6.5m0-6.5l6.5 6.5"/>`
	crosshairInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 5.25v-3m0 11.5v-3M10.75 8h3M2.25 8h3"/><circle cx="8" cy="8" r="6.25"/></g>`
	cubeInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5zM8 14V8m5.75-3L8 8M2 5l6 3"/>`
	cursorInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 1.75l4.5 12.5l2.5-5.5l5.5-2.5zm7.5 7.5l4 4"/>`
	databaseInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 1.75c-3.75 0-5.25 2-5.25 2v8.5s1.5 2 5.25 2s5.25-2 5.25-2v-8.5s-1.5-2-5.25-2z"/><path d="M2.75 8.25s1.5 2 5.25 2s5.25-2 5.25-2m-10.5-4s1.5 2 5.25 2s5.25-2 5.25-2"/></g>`
	diamondInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.25 8L8 14.75L14.75 8L8 1.25z"/>`
	diffInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 13.75h8m0-7.5h-8m4-4v8"/>`
	discInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><circle cx="8" cy="8" r="1.75"/></g>`
	downloadInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 13.25h9m-8.5-6.5l4 3.5l4-3.5m-4-5v8.5"/>`
	dropletInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 9a5.25 5.25 0 1 0 10.5 0C13.25 5.75 8 1.75 8 1.75S2.75 5.75 2.75 9z"/>`
	eraserInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 13.25h-9.5l-3-3l7.5-7.5l5 5l-5.5 5.5m-3.5-6.5l5 5"/>`
	extensionsInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 8.75h5.5v5.5m5-12.5v4m-2-2h4m-12.5-1v11.5h11.5v-5.5h-6v-6z"/></g>`
	eyeInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 8s2-4.25 6.25-4.25S14.25 8 14.25 8s-2 4.25-6.25 4.25S1.75 8 1.75 8z"/><circle cx="8" cy="8" r="1.25" fill="currentColor"/></g>`
	eyeSlashInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.75 3.75c3.5.5 5.5 4.25 5.5 4.25s-.5 1.25-1.5 2.25m-2.5 1.5c-6 2-8.5-3.75-8.5-3.75s.5-1.75 3-3.25"/><path fill="currentColor" d="M8.625 9.083a1.25 1.25 0 0 1-1.649-.366a1.25 1.25 0 0 1 .22-1.675L8 8z"/><path d="m3.75 1.75l8.5 12.5"/></g>`
	faceFrownInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 5s.5-1 2.25-1s2.25 1 2.25 1"/></g>`
	faceNeutralInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 4.5h4.5"/></g>`
	faceSmileInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 4s.5 1.5 2.25 1.5s2.25-1.5 2.25-1.5"/></g>`
	fileInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 1.75h5.5l5 5v7.5H2.75z"/><path d="M7.75 2.25v5h5"/></g>`
	fileBinaryInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5M1.75 10.8h3v3.5h-3z"/><path d="M7.25 14.2h3m-3-3.5h1.5v3m-1-11.45v5h5.05"/></g>`
	fileCodeInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5h-2"/><path d="M7.75 2.25v5h5.05M6.75 10.8l2 1.75l-2 1.75m-3-3.5l-2 1.75l2 1.75"/></g>`
	fileSymlinkInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5h-4"/><path d="M7.75 2.25v5h5m-10 7l3.5-3.5m0 3v-3h-3"/></g>`
	filesInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9.25 1.75l4 4v5.5h-7.5v-9.5z"/><path d="M9.25 2.25v3.5h3.5m-2.5 6v2.5h-7.5v-9.5h2.5"/></g>`
	filterInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h12.5v1.5l-5 5.5v4l-2.5 1.5v-5.5l-5-5.5z"/>`
	flagInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 14.25v-11s2-1.5 4-1.5s2.5 1.5 4.5 1.5s4-1.5 4-1.5v7s-2 1.5-4 1.5s-2.5-1.5-4.5-1.5s-4 1.5-4 1.5"/>`
	flameInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 7.75c2 2 2.5-2.5 3.5-2s1.5 2 1.5 3.25c0 3.25-2.35 5.25-5.25 5.25s-5.25-2.5-5.25-6s3.5-7 5.5-7c0 0-2 4.5 0 6.5z"/>`
	floppyDiskInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 2.75v10.5h10.5v-7.5l-3-3z"/><path d="M5.75 13.25v-3.5h4.5v3.5"/></g>`
	folderInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75v10.5h12.5v-8.5h-6l-1.5-2z"/>`
	folderSymlinkInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.75 13.25l3.5-3.5m0 3v-3h-3"/><path d="M8.25 13.25h6v-8.5h-6l-1.5-2h-5v4"/></g>`
	foldersInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.75 2.25v8h9.5v-6.5h-5l-1.5-1.5z"/><path d="M4.75 5.25h-3v8h9.5v-3"/></g>`
	forwardInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 13.25c.5-6 5.5-7.5 8-7v-3.5L14.25 8l-4.5 5.25v-3.5c-2.5-.5-6.5.5-8 3.5z"/>`
	gamepadInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 3.75c-2 5-2 9 0 9.5s2.5-2 2.5-2h4.5s.5 2.5 2.5 2s2-4.5 0-9.5h-2l-1 1h-3.5l-1-1z"/>`
	gemInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 2.75h6.5l3 3.5l-6.25 7l-6.25-7z"/>`
	giftInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75h12.5v3.5H1.75z"/><path d="M10.25 4.75H8c0-2 .5-3 2.25-3c2 0 2 3 0 3zm-4.5 0H8c0-2-.5-3-2.25-3c-2 0-2 3 0 3zm2.25 9V5M2.75 8.75v5.5h10.5v-5.5"/></g>`
	gitBranchInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="11.5" cy="3.5" r="1.75"/><circle cx="4.5" cy="12.5" r="1.75"/><path d="M5.25 8.25c3 0 6 .5 6-2.5m-6.5 4.5v-4.5"/></g>`
	gitCherryPickInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5" cy="8" r="2.25"/><path d="M5 10.75v3.5m0-12.5v3.5M11.75 8h1.5m-4.5-3.25h1.5l1 3.25l-1 3.25h-1.5"/></g>`
	gitCommitInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="2.25"/><path d="M8 10.75v3.5m0-12.5v3.5"/></g>`
	gitCompareInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M3.75 5.75v5c0 1 .5 1.5 1.5 1.5h2m-.5 2l1.5-2l-1.5-2m5.5 0v-5c0-1-.5-1.5-1.5-1.5h-2m.5-2l-1.5 2l1.5 2"/></g>`
	gitForkInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="12.5" r="1.75"/><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="11.5" cy="3.5" r="1.75"/><path d="M8 8.75v1.5m-3.25-4.5c0 3.5 6.5 3.5 6.5 0"/></g>`
	gitMergeInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="4.5" cy="12.5" r="1.75"/><circle cx="12.5" cy="8.5" r="1.75"/><path d="M4.75 10.25v-4.5c1 2 2 3 5.5 3"/></g>`
	gitRequestInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="m9.25 1.75l-1.5 2l1.5 2m3 4.5v-5c0-1-.5-1.5-1.5-1.5h-2m-5 2v4.5"/></g>`
	gitRequestCrossInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M12.25 7.25v3m-8.5-4.5v4.5m10.5-8.5l-3.5 3.5m0-3.5l3.5 3.5"/></g>`
	gitRequestDraftInnerSVG     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M7.75 2.75h.5m2.5 0h.5m1.5 2.5v-.5m0 3v.5m-9-2.5v4.5"/></g>`
	githubInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 14.25s-.5-2 .5-3c0 0-2 0-3.5-1.5s-1-4.5 0-5.5c-.5-1.5.5-2.5.5-2.5s1.5 0 2.5 1c1-.5 3.5-.5 4.5 0c1-1 2.5-1 2.5-1s1 1 .5 2.5c1 1 1.5 4 0 5.5s-3.5 1.5-3.5 1.5c1 1 .5 3 .5 3m-5-.5c-1.5.5-3-.5-3.5-1"/>`
	gitlabInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 14.25l-6.25-4.5l2-8l2 5.5h4.5l2-5.5l2 8z"/>`
	glassesInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="11" r="2.25"/><circle cx="12" cy="11" r="2.25"/><path d="M14.25 10.75c-1.5-6-2-6.5-3.5-7m-9 7c1.5-6 2-6.5 3.5-7m1 7c1-1 2.5-1 3.5 0"/></g>`
	globeInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M2 8.25h12M8.25 14.2C11 11 11 5 8.25 1.8m-.5 12.4C5 11 5 5 7.75 1.8"/></g>`
	grabHorizontalInnerSVG      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="2.5" cy="5.5" r=".75"/><circle cx="8" cy="5.5" r=".75"/><circle cx="13.5" cy="5.5" r=".75"/><circle cx="2.5" cy="10.5" r=".75"/><circle cx="8" cy="10.5" r=".75"/><circle cx="13.5" cy="10.5" r=".75"/></g>`
	grabVerticalInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5.5" cy="2.5" r=".75"/><circle cx="5.5" cy="8" r=".75"/><circle cx="5.5" cy="13.5" r=".75"/><circle cx="10.496" cy="2.5" r=".75"/><circle cx="10.496" cy="8" r=".75"/><circle cx="10.496" cy="13.5" r=".75"/></g>`
	graduateCapInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 9.25V6L8 2.75L1.75 6L8 9.25l3.25-1.5v3.5c0 1-1.5 2-3.25 2s-3.25-1-3.25-2v-3.5"/>`
	hashInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 10.25h9.5m-8.5-4.5h9.5m-2.5-4l-1.5 12.5m-2.5-12.5l-1.5 12.5"/>`
	headphonesInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.75c0-2.5 3.5-2 3.5-2v4.5s-3.5.5-3.5-2.5v-3.5c0-3 .5-6.5 6.25-6.5s6.25 3.5 6.25 6.5v3.5c0 3-3.5 2.5-3.5 2.5v-4.5s3.5-.5 3.5 2"/>`
	heartInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 9.75c3 3.5 4.75 4.5 4.75 4.5s1.75-1 4.75-4.5s1-7-1.5-7s-3.25 3-3.25 3s-.75-3-3.25-3s-4.5 3.5-1.5 7z"/>`
	helpInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M5.75 6.75c0-1 1-2 2.25-2s2.25 1.034 2.25 2c0 1.5-1.5 1.5-2.25 2m0 2.5v0"/></g>`
	hexagonInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5z"/>`
	homeInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 5.75v7.5h8.5v-7.5m-10.5 1.5L8 1.75l6.25 5.5"/>`
	hourglassInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.75 13.75c0-5-2-4-2-5.75s2-.75 2-5.75m-7.5 11.5c0-5 2-4 2-5.75s-2-.75-2-5.75m-1.5-.5h10.5m-10.5 12.5h10.5"/>`
	idInnerSVG                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><circle cx="8" cy="7.5" r="2.25"/><path d="M4.75 12.75c0-1 .75-3 3.25-3s3.25 2 3.25 3"/></g>`
	imageInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><path d="m3.75 13.2l6.5-5.5l4 3"/><circle cx="5.25" cy="6.25" r=".5" fill="currentColor"/></g>`
	inboxInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 13.25h12.5v-5l-2.5-5.5h-7.5l-2.5 5.5z"/><path d="M2.25 8.75h3l1 1.5h3.5l1-1.5h3"/></g>`
	infinityInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 5c2.5 1 3.5 5 6 6s3.25-1.25 3.25-3S13.5 4 11 5s-3.5 5-6 6s-3.25-1.25-3.25-3S2.5 4 5 5z"/>`
	infoInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8 5.25v0m0 6v-3.5"/></g>`
	keyInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 1.75a4.25 4.25 0 0 0-4.104 5.354L1.75 11.25v3h3v-1.5h1.5v-1.5h1.5L8.9 10.1c.359.098.728.148 1.1.15a4.25 4.25 0 0 0 0-8.5z"/><circle cx="10.75" cy="5.25" r=".5" fill="currentColor"/></g>`
	laptopInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h10.5v7.5H2.75zm0 7.5l-1 3h12.5l-1-3"/>`
	layoutColumnsInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.25.5v9.5"/>`
	layoutDashboardInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.5 4h5.5m-11.5 2.5h5.5m.25-6v9.5"/>`
	layoutGridInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12m-3.75-4.75v9.5m-4.5-9.5v9.5"/>`
	layoutListInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm3.5.5v9.5m-3-6.5h11.5m-11.5 3.5h11.5"/>`
	layoutRowsInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12"/>`
	layoutSidebarInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm4.5.25v9.5"/>`
	layoutStackHInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12M8 8v4.75"/>`
	layoutStackVInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.25.5v9.5M8 8h6"/>`
	lightbulbInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 14.25h2.5M8 1.75c-2.75 0-4.25 2-4.25 4s2 2.5 2 4.5v1h4.5v-1c0-2 2-2.5 2-4.5s-1.5-4-4.25-4z"/>`
	lightningBoltInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.25 1.75l-6.5 7.5l4.5.5l-.5 4.5l6.5-7.5l-4.5-.5z"/>`
	linkInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 4.75c3 0 4.5 1.5 4.5 3.25s-1.5 3.25-4.5 3.25M5.75 8h4.5m-4-3.25c-3 0-4.5 1.5-4.5 3.25s1.5 3.25 4.5 3.25"/>`
	linkExternalInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 2.75h-5.5v10.5h10.5v-5.5m0-5l-5.5 5.5m3-6.5h3.5v3.5"/>`
	linkSlashInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.75 1.75l-5.5 12.5m4.5-9.5c3 0 4.5 1.5 4.5 3.25s-1.5 3.25-4.5 3.25m-3.5-6.5c-3 0-4.5 1.5-4.5 3.25s1.5 3.25 4.5 3.25"/>`
	mailInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 3.75h12.5v9.5H1.75z"/><path d="m2.25 4.25l5.75 5l5.75-5"/></g>`
	mapInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 5.25v8.5m-4.5-10.5v8.5m-4 2.5v-9.5l4-2l4.5 2l4-2v9.5l-4 2l-4.5-2z"/>`
	mapPinInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.25 7c0 3.75-5.25 7.25-5.25 7.25S2.75 10.75 2.75 7a5.25 5.25 0 0 1 10.5 0z"/><circle cx="8" cy="7" r="1.25" fill="currentColor"/></g>`
	mediaBackInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.25 13.25L4.75 8l8.5-5.25zm-11.5-9.5v8.5"/>`
	mediaEjectInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 11.25h10.5L8 2.75zm10.5 3H2.75"/>`
	mediaFastForwardInnerSVG    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 3.75v8.5l6-4.25zm-6.5 0v8.5l6-4.25z"/>`
	mediaPauseInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h3.5v10.5h-3.5zm7 0h3.5v10.5h-3.5z"/>`
	mediaPlayInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75v10.5L12.25 8z"/>`
	mediaRewindInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.75 3.75v8.5L1.75 8zm6.5 0v8.5L8.25 8z"/>`
	mediaSkipInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 13.25L11.25 8l-8.5-5.25zm11.5-9.5v8.5"/>`
	menuHamburgerInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 12.25h10.5m-10.5-4h10.5m-10.5-4h10.5"/>`
	menuKebabInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="2.5" r=".75"/><circle cx="8" cy="8" r=".75"/><circle cx="8" cy="13.5" r=".75"/></g>`
	menuMeatballInnerSVG        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="2.5" cy="8" r=".75"/><circle cx="8" cy="8" r=".75"/><circle cx="13.5" cy="8" r=".75"/></g>`
	messageInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 14.25V2.75h12.5v8.5h-8.5z"/>`
	messagesInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 14.25v-9h-9.5v6h6z"/><path d="m4.75 7.25l-3 3v-8.5h10v3"/></g>`
	microphoneInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 1.75c-2.25 0-2.25 2-2.25 3v1.5c0 1 0 3 2.25 3s2.25-2 2.25-3v-1.5c0-1 0-3-2.25-3z"/><path d="M8 13v1.25m-5.25-6.5s0 4.5 5.25 4.508c5.25.008 5.25-4.508 5.25-4.508"/></g>`
	minusInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.25 7.75H2.75"/>`
	mobileInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 1.75h8.5v12.5h-8.5zm4.5 10h-.5"/>`
	monitorInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h12.5v9.5H1.75zm3 12.5h6.5M8 11.75v2.5"/>`
	monitorArrowInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.2 7.75v3.5H1.7v-9.5h6.5M4.75 14.2h6.5M8 11.7v2.5m1.75-7.95l4.5-4.5m-3.5-.5h4v4"/>`
	monitorCrossInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.2 7.75v3.5H1.7v-9.5h6.5M4.75 14.2h6.5M8 11.7v2.5m6.2-12.45l-3.5 3.5m0-3.5l3.5 3.5"/>`
	moonInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8c0 3.45 2.8 6.25 6.25 6.25c3.41-.003 6.25-3 6.25-6c-1 .5-4 1.5-6-.5s-1-5-.5-6c-3 0-6 2.84-6 6.25z"/>`
	moveInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.25 10.25l2-2.25l-2-2.25m-2-2L8 1.75l-2.25 2m-2 2L1.75 8l2 2.25m2 2l2.25 2l2.25-2M8 1.75v12M13.75 8h-12"/>`
	musicInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="12" r="2.25"/><circle cx="12" cy="11" r="2.25"/><path d="M6.25 12V2.75l8-1V11"/></g>`
	newspaperInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.2 14.2h.5c1.5 0 2.5-1 2.5-2.5v-6h-3m-9.5-4h9.5v12.5h-7c-1.5 0-2.5-1-2.5-2.5V2.26zm3.05 9.5h3.5"/><path d="M4.75 4.75h3.5v3.5h-3.5z"/></g>`
	northStarInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.75 7.75h-12m6-6v12m-3.5-2.5l7-7m0 7l-7-7"/>`
	notesInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 1.75h10.5v12.5H2.75zm3 6h4.5m-4.5 3h2.5m-2.5-6h4.5"/>`
	notesCrossInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 14.25h-5.5V1.75h10.5v6.5m1 2.5l-3.5 3.5m-5-6.5h4.5m-4.5 3h1.5m-1.5-6h4.5m.5 6l3.5 3.5"/>`
	notesTickInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.25 14.25h-4.5V1.75h10.5v7.5m-3.5 3.5l1.5 1.5l3-2.5m-8.5-4h4.5m-4.5 3h1.5m-1.5-6h4.5"/>`
	nutInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5v-6.5z"/><circle cx="8" cy="8" r="2.25"/></g>`
	octagonInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 1.75h5.5l3.5 3.5v5.5l-3.5 3.5h-5.5l-3.5-3.5v-5.5z"/>`
	octagonWarningInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 1.75h5.5l3.5 3.5v5.5l-3.5 3.5h-5.5l-3.5-3.5v-5.5zM8 11.25v0m0-6.5v3.5"/>`
	organisationInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 1.75h3.5v3.5h-3.5zm4 9h3.5v3.5h-3.5zm-8 0h3.5v3.5h-3.5zm5.75-5v2m-3.75 2.5v-2h7.5v2"/>`
	packageInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v8.5h12.5v-8.5l-3.5-4h-5.5zm6.25-4v3.5m-5.75.5h11.5"/>`
	padlockInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 6.75h10.5v7.5H2.75zm2-.5s-1-4.5 3.25-4.5s3.25 4.5 3.25 4.5"/>`
	paperPlaneInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 1.75l12.5 6l-12.5 6.5l1.5-6.5zm2 6h3.5"/>`
	paperclipInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 10.25v-7s0-1.5-1.75-1.5s-1.75 1.5-1.75 1.5v8s0 3 3.25 3s3.25-3 3.25-3v-4.5"/>`
	pencilInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.25v3h3l9.5-9.5l-3-3zm7-6.5l2.5 2.5"/>`
	peopleInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5" cy="9" r="2.25"/><circle cx="11" cy="4" r="2.25"/><path d="M7.75 9.25c0-1 .75-3 3.25-3s3.25 2 3.25 3m-12.5 5c0-1 .75-3 3.25-3s3.25 2 3.25 3"/></g>`
	personInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="6" r="3.25"/><path d="M2.75 14.25c0-2.5 2-5 5.25-5s5.25 2.5 5.25 5"/></g>`
	phoneInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5z"/>`
	phoneCallInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 0c2.5 0 4.5 2 4.5 4.5m-4.5-2c1 0 2 1 2 2"/>`
	phoneCrossInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm11.5 1l-3.5 3.5m0-3.5l3.5 3.5"/>`
	phoneForwardInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 3h4.5m-2 2l2-2l-2-2"/>`
	phoneIncomingInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5z"/><path d="m13.25 2.75l-3.5 3.5m0-3v3h3"/></g>`
	phoneOutgoingInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 4.5l3.5-3.5m0 3v-3h-3"/>`
	pinInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.25 10.25l4 4m-12.5-7.5l5-5s1 2 2 3s4.5 2 4.5 2l-6.5 6.5s-1-3.5-2-4.5s-3-2-3-2z"/>`
	plantPotInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.75 6.75c0 1.25-.75 3-.75 3m.25-2.5s.75-2-1-3.5s-4.5-1-4.5-1s0 2 1.5 3.5s4 1 4 1zm.5-1s-.75-2 1-3.5s4.5-1 4.5-1s0 2-1.5 3.5s-4 1-4 1zm-4 3.5h6.5s.5 4.5-3.25 4.5s-3.25-4.5-3.25-4.5z"/>`
	plusInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.75 7.75h-10m5-5v10"/>`
	powerInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 1.75v6.5m4.25-5s2 1.298 2 4.75a6.25 6.25 0 1 1-12.5 0c0-3.452 2-4.75 2-4.75"/>`
	printerInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.75 9.75h6.5v4.5h-6.5z"/><path d="M4.75 4.25v-2.5h6.5v2.5m-7 8h-2.5v-7.5h12.5v7.5h-2.5"/></g>`
	pulseInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8.25h2.5l2-4.5l3.5 8.5l2-4h2.5"/>`
	quoteInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.25 3.75h-4.5v5.5c0 3.5 2.5 4.5 4.5 4c-1.5-1.5-1.5-2.5-1.5-4h1.5zm7 0h-4.5v5.5c0 3.5 2.5 4.5 4.5 4c-1.5-1.5-1.5-2.5-1.5-4h1.5z"/>`
	refreshInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 10.75h-3m12.5-2c0 3-2.798 5.5-6.25 5.5c-3.75 0-6.25-3.5-6.25-3.5v3.5m9.5-9h3m-12.5 2c0-3 2.798-5.5 6.25-5.5c3.75 0 6.25 3.5 6.25 3.5v-3.5"/>`
	replyInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 13.25c-.5-6-5.5-7.5-8-7v-3.5L1.75 8l4.5 5.25v-3.5c2.5-.5 6.5.5 8 3.5z"/>`
	robotInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 5.75h12.5v7.5H1.75z"/><path d="M10.75 8.75v1.5m-5.5-1.5v1.5m-.5-7.5l3.25 3l3.25-3"/></g>`
	rocketInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.25 9.75l-2-.5s0-1.5.5-3s4-1.5 4-1.5m-.5 7l.5 2s1.5 0 3-.5s1.5-4 1.5-4m-7 .5l2 2s5-2 6.5-4.5s1.5-5.5 1.5-5.5s-3 0-5.5 1.5s-4.5 6.5-4.5 6.5z"/><path fill="currentColor" d="m1.75 14.25l2-1l-1-1z"/><circle cx="10.25" cy="5.75" r=".5" fill="currentColor"/></g>`
	rotateAntiClockwiseInnerSVG = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 5.25h-3m0 3.5c0 2.5 2.798 5.5 6.25 5.5a6.25 6.25 0 1 0 0-12.5c-3.75 0-6.25 3.5-6.25 3.5v-3.5"/>`
	rotateClockwiseInnerSVG     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 5.25h3m0 3.5c0 2.5-2.798 5.5-6.25 5.5a6.25 6.25 0 1 1 0-12.5c3.75 0 6.25 3.5 6.25 3.5v-3.5"/>`
	scalesInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 3.75c1 1 2.5 1.5 4 0h4.5c1.5 1.5 3 1 4 0M8 1.75v12m-3.25.5h6.5"/><path d="m12.75 4.75l-2 5c.5 1 3.5 1 4 0zm-9.5 0l-2 5c.5 1 3.5 1 4 0z"/></g>`
	screenMaximiseInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 14.25h-3.5v-3.5m12.5 0v3.5h-3.5m0-12.5h3.5v3.5m-12.5 0v-3.5h3.5"/>`
	screenMinimiseInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 10.75h3.5v3.5m5.5 0v-3.5h3.5m0-5.5h-3.5v-3.5m-5.5 0v3.5h-3.5"/>`
	searchInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.25 11.25l3 3"/><circle cx="7.5" cy="7.5" r="4.75"/></g>`
	serverInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.25h12.5v10H1.75zm.5 5h11.5m-9 2.5v0m0-5v0m6.5 0h-3m3 5h-3"/>`
	shareInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="8" r="2.25"/><circle cx="12" cy="12" r="2.25"/><circle cx="12" cy="4" r="2.25"/><path d="m6 9l4 2M6 7l4-2"/></g>`
	shieldInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5z"/>`
	shieldCrossInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm1.75 4l-3.5 3.5m0-3.5l3.5 3.5"/>`
	shieldKeyholeInnerSVG       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm0 5.5v3"/><circle cx="8" cy="6.5" r=".75" fill="currentColor"/></g>`
	shieldTickInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5z"/><path d="m5.75 7.75l1.5 1.5l3-3.5"/></g>`
	shieldWarningInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm0 9v0m0-5.5v3"/>`
	shoppingBagInnerSVG         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 4.75h10.5v9.5H2.75z"/><path d="M5.75 7.75c0 1.5 1 2.5 2.25 2.5s2.25-1 2.25-2.5m-7.5-3l1.5-3h7.5l1.5 3"/></g>`
	signInInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 2.25h-3.5v12h3.5m4-9.5l-3.5 3.5l3.5 3.5m5-3.5h-8.5"/>`
	signOutInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 2.25h-3.5v12h3.5m5.5-9.5l3.5 3.5l-3.5 3.5m-5-3.5h8.5"/>`
	signpostInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 9.25h10.5l2-2.25l-2-2.25H1.75zm5.5.5v4.5m0-12.5v2.5"/>`
	skullInnerSVG               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 11.25h3v3h6.5v-3h3s1-9.5-6.25-9.5s-6.25 9.5-6.25 9.5z"/><circle cx="5.25" cy="7.75" r=".5" fill="currentColor"/><circle cx="10.75" cy="7.75" r=".5" fill="currentColor"/></g>`
	snowflakeInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.75 7.75h-12m6-6v12m-2.5-1l2.5-2.5l2.5 2.5m-7.5-7.5l2.5 2.5l-2.5 2.5m7.5-7.5l-2.5 2.5l-2.5-2.5m7.5 7.5l-2.5-2.5l2.5-2.5"/>`
	soundDownInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm9 .5s1 .5 1 1.75s-1 1.75-1 1.75"/>`
	soundMuteInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm12.5 0l-3.5 4.5m0-4.5l3.5 4.5"/>`
	soundUpInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm9 .5s1 .5 1 1.75s-1 1.75-1 1.75m1-6.5c2 1 3 2.5 3 4.75s-1 3.75-3 4.75"/>`
	speakerInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 1.75h9.5v12.5h-9.5zm5 2.5h-.5"/><circle cx="8" cy="9.5" r="2.25"/></g>`
	squareInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h10.5v10.5H2.75z"/>`
	squareCrossInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.25 5.75l-4.5 4.5m0-4.5l4.5 4.5m-7.5-7.5h10.5v10.5H2.75z"/>`
	squareTickInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.25 2.75h-7.5v10.5h10.5v-3.5"/><path d="m5.75 7.75l2.5 2.5l6-6.5"/></g>`
	stackInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11L8 14.25L14.25 11M1.75 8L8 11.25L14.25 8M8 1.75L1.75 5L8 8.25L14.25 5z"/>`
	stackPopInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.25 6.75L1.75 8L8 11.25L14.25 8l-2.5-1.25M1.75 11L8 14.25L14.25 11M8 8.25v-6.5m-2.25 2l2.25-2l2.25 2"/>`
	stackPushInnerSVG           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 7.25L1.75 8L8 11.25L14.25 8l-1.5-.75M1.75 11L8 14.25L14.25 11"/><path d="M8 8.25v-6.5m-2.25 4.5l2.25 2l2.25-2"/></g>`
	starInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l-2.25 4l-4 .5l3 3.5l-1 4.5l4.25-2l4.25 2l-1-4.5l3-3.5l-4-.5z"/>`
	stickyNoteInnerSVG          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.25 13.25h-6.5V2.75h10.5v6.5z"/><path d="M8.75 13.25v-4.5h4.5"/></g>`
	sunInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="3.25"/><path d="m2.75 13.25l.5-.5m9.5 0l.5.5m-.5-10l.5-.5m-10 .5l-.5-.5M2.25 8h-1m13.5 0h-1M8 13.75v1m0-13.5v1"/></g>`
	swapHorizontalInnerSVG      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.75 8.25l-3 3l3 3m7.5-3H2.75m7.5-9.5l3 3l-3 3m-7.5-3h10.5"/>`
	swapVerticalInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.75 5.75l-3-3l-3 3m3 7.5V2.75m9.5 7.5l-3 3l-3-3m3-7.5v10.5"/>`
	swordInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 9.25l1.5 2.5l2 1.5m-4.5 0l1 1m1.5-2.5l-1.5 1.5m3-1l8.5-8.5v-2h-2l-8.5 8.5"/>`
	swordsInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2.75 9.25l1.5 2.5l2 1.5m-4.5 0l1 1m1.5-2.5l-1.5 1.5m3-1l8.5-8.5v-2h-2l-8.5 8.5"/><path d="M10.25 12.25L8 10m2-2l2.25 2.25m1-1l-1.5 2.5l-2 1.5m4.5 0l-1 1m-1.5-2.5l1.5 1.5M6 8L1.75 3.75v-2h2L8 6"/></g>`
	tabletInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 1.75h10.5v12.5H2.75zm5.5 10h-.5"/>`
	tagInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7.25 14.25l-5.5-5.5l7-7h5.5v5.5z"/><circle cx="11" cy="5" r=".5" fill="currentColor"/></g>`
	telescopeInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 5.75l1 2.5m3.5-4.5l1.5 3.5m-9 0l1 2.5l11.5-3.5l-2-4.5zm6 3.95v3m-3-.5L7 11.2l1.75-.5l2.5 3"/>`
	tentInnerSVG                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.25 14.25L8 10l2.75 4.25"/><path d="m9.75 1.75l-8 12.5h12.5l-8-12.5"/></g>`
	terminalInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><path d="M8.75 10.25h2.5m-6.5-4.5L7.25 8l-2.5 2.25"/></g>`
	thumbDownInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 10.25c1.5 0 3 4 4.5 4v-4h4.5s-.5-7.5-3.5-7.5h-5.5zm0 0h-3.5v-7.5h3.5"/>`
	thumbUpInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 5.75c1.5 0 3-4 4.5-4v4h4.5s-.5 7.5-3.5 7.5h-5.5zm0 0h-3.5v7.5h3.5"/>`
	tickInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 8.75l3.5 3.5l7-7.5"/>`
	tickDoubleInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 9.75l2.5 2.5m3.5-4l2.5-2.5m-4.5 4l2.5 2.5l6-6.5"/>`
	ticketInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v3s-2 0-2 1.75s2 1.75 2 1.75v3H1.75v-3s2 0 2-1.75s-2-1.75-2-1.75zm7 8v1.5m0-9.5v1.5m0 2.5v1.5"/>`
	treeFirInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l-4.25 5.5h2.5l-3.5 4h4v3h2.5v-3h4l-3.5-4h2.5z"/>`
	triangleInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 2.75l-6.25 11.5h12.5z"/>`
	trophyInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 10.75h6.5v3.5h-6.5zm3.25-2v2m-3.25-9c-1.5 0-3 .5-3 2.25s1.5 2.25 3 2.25m6.5-4.5c1.5 0 3 .5 3 2.25s-1.5 2.25-3 2.25m-6.5-4.5h6.5v3.5c0 1.5-1 3-3.25 3s-3.25-1.5-3.25-3z"/>`
	umbrellaInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8.25s.5-6.5 6.25-6.5s6.25 6.5 6.25 6.5zm6 .5v4s0 1.5 1.5 1.5s1.5-1.5 1.5-1.5"/>`
	uploadInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 2.75h9m-8.5 6.5l4-3.5l4 3.5m-4 5v-8.5"/>`
	wifiInnerSVG                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/>`
	wifiFairInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/><path d="M4.25 8c2-1.75 5.5-1.75 7.5 0"/></g>`
	wifiPoorInnerSVG            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/><path d="M5 9c.75-1.75 5.25-1.75 6 0"/></g>`
	wifiSlashInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 3.25c-1.5 0-3.5 1.5-3.5 1.5L8 13.25l2.25-3m-1.5-7.5s2.977-.135 5.5 2l-2 2.5m-8-5.5l8 10.5"/>`
	wifiWarningInnerSVG         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 4.75C11 2 5 2 1.75 4.75L8 13.25l1-1.5m3.25 2v0m0-6v3.5"/>`
	zoomInInnerSVG              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="7.5" cy="7.5" r="4.75"/><path d="M9.25 7.5h-3.5M7.5 5.75v3.5m3.75 2l3 3"/></g>`
	zoomOutInnerSVG             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="7.5" cy="7.5" r="4.75"/><path d="M9.25 7.5h-3.5m5.5 3.75l3 3"/></g>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
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

func Apps(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appsInnerSVG).
		Element(children...)
}

func AppsMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appsMinusInnerSVG).
		Element(children...)
}

func AppsPlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appsPlusInnerSVG).
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

func ArrowDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownLeftInnerSVG).
		Element(children...)
}

func ArrowDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownRightInnerSVG).
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

func ArrowUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpLeftInnerSVG).
		Element(children...)
}

func ArrowUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpRightInnerSVG).
		Element(children...)
}

func AtSign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atSignInnerSVG).
		Element(children...)
}

func Atom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atomInnerSVG).
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

func Bin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binInnerSVG).
		Element(children...)
}

func Binary(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(binaryInnerSVG).
		Element(children...)
}

func Block(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(blockInnerSVG).
		Element(children...)
}

func Bluetooth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothInnerSVG).
		Element(children...)
}

func BluetoothConnected(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothConnectedInnerSVG).
		Element(children...)
}

func BluetoothSearching(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothSearchingInnerSVG).
		Element(children...)
}

func BluetoothSlash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bluetoothSlashInnerSVG).
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

func BookOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookOpenInnerSVG).
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

func CameraVideo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraVideoInnerSVG).
		Element(children...)
}

func CameraVideoSlash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraVideoSlashInnerSVG).
		Element(children...)
}

func Candy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(candyInnerSVG).
		Element(children...)
}

func Cards(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardsInnerSVG).
		Element(children...)
}

func Cast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(castInnerSVG).
		Element(children...)
}

func Certificate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(certificateInnerSVG).
		Element(children...)
}

func ChartBar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartBarInnerSVG).
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

func ChevronsDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsDownInnerSVG).
		Element(children...)
}

func ChevronsLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsLeftInnerSVG).
		Element(children...)
}

func ChevronsRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsRightInnerSVG).
		Element(children...)
}

func ChevronsUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsUpInnerSVG).
		Element(children...)
}

func ChevronsUpDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronsUpDownInnerSVG).
		Element(children...)
}

func Chip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chipInnerSVG).
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

func CircleCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleCrossInnerSVG).
		Element(children...)
}

func CircleMinus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleMinusInnerSVG).
		Element(children...)
}

func CircleTick(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleTickInnerSVG).
		Element(children...)
}

func CircleWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleWarningInnerSVG).
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

func ClipboardTick(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clipboardTickInnerSVG).
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

func ClockAlarm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockAlarmInnerSVG).
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

func Clover(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloverInnerSVG).
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

func Compass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassInnerSVG).
		Element(children...)
}

func ConicalFlask(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(conicalFlaskInnerSVG).
		Element(children...)
}

func Container(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(containerInnerSVG).
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

func Copyleft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(copyleftInnerSVG).
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

func Cross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crossInnerSVG).
		Element(children...)
}

func Crosshair(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cubeInnerSVG).
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

func Diff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diffInnerSVG).
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

func Droplet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dropletInnerSVG).
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

func Extensions(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(extensionsInnerSVG).
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

func FaceFrown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceFrownInnerSVG).
		Element(children...)
}

func FaceNeutral(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceNeutralInnerSVG).
		Element(children...)
}

func FaceSmile(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(faceSmileInnerSVG).
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

func FileBinary(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileBinaryInnerSVG).
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

func FileSymlink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileSymlinkInnerSVG).
		Element(children...)
}

func Files(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filesInnerSVG).
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

func Flame(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flameInnerSVG).
		Element(children...)
}

func FloppyDisk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(floppyDiskInnerSVG).
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

func FolderSymlink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderSymlinkInnerSVG).
		Element(children...)
}

func Folders(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(foldersInnerSVG).
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

func Gem(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gemInnerSVG).
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

func GitBranch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitBranchInnerSVG).
		Element(children...)
}

func GitCherryPick(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitCherryPickInnerSVG).
		Element(children...)
}

func GitCommit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitMergeInnerSVG).
		Element(children...)
}

func GitRequest(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitRequestInnerSVG).
		Element(children...)
}

func GitRequestCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitRequestCrossInnerSVG).
		Element(children...)
}

func GitRequestDraft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitRequestDraftInnerSVG).
		Element(children...)
}

func Github(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(githubInnerSVG).
		Element(children...)
}

func Gitlab(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gitlabInnerSVG).
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

func GrabHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(grabHorizontalInnerSVG).
		Element(children...)
}

func GrabVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(grabVerticalInnerSVG).
		Element(children...)
}

func GraduateCap(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graduateCapInnerSVG).
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

func Help(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(helpInnerSVG).
		Element(children...)
}

func Hexagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hexagonInnerSVG).
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

func Id(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(idInnerSVG).
		Element(children...)
}

func Image(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageInnerSVG).
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

func Infinity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infinityInnerSVG).
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

func LayoutColumns(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutColumnsInnerSVG).
		Element(children...)
}

func LayoutDashboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutDashboardInnerSVG).
		Element(children...)
}

func LayoutGrid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutGridInnerSVG).
		Element(children...)
}

func LayoutList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutListInnerSVG).
		Element(children...)
}

func LayoutRows(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutRowsInnerSVG).
		Element(children...)
}

func LayoutSidebar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutSidebarInnerSVG).
		Element(children...)
}

func LayoutStackH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutStackHInnerSVG).
		Element(children...)
}

func LayoutStackV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layoutStackVInnerSVG).
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

func LightningBolt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightningBoltInnerSVG).
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

func LinkExternal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkExternalInnerSVG).
		Element(children...)
}

func LinkSlash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkSlashInnerSVG).
		Element(children...)
}

func Mail(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailInnerSVG).
		Element(children...)
}

func Map(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapInnerSVG).
		Element(children...)
}

func MapPin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapPinInnerSVG).
		Element(children...)
}

func MediaBack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaBackInnerSVG).
		Element(children...)
}

func MediaEject(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaEjectInnerSVG).
		Element(children...)
}

func MediaFastForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaFastForwardInnerSVG).
		Element(children...)
}

func MediaPause(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaPauseInnerSVG).
		Element(children...)
}

func MediaPlay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaPlayInnerSVG).
		Element(children...)
}

func MediaRewind(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaRewindInnerSVG).
		Element(children...)
}

func MediaSkip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediaSkipInnerSVG).
		Element(children...)
}

func MenuHamburger(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuHamburgerInnerSVG).
		Element(children...)
}

func MenuKebab(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuKebabInnerSVG).
		Element(children...)
}

func MenuMeatball(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuMeatballInnerSVG).
		Element(children...)
}

func Message(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageInnerSVG).
		Element(children...)
}

func Messages(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messagesInnerSVG).
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

func Monitor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(monitorInnerSVG).
		Element(children...)
}

func MonitorArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(monitorArrowInnerSVG).
		Element(children...)
}

func MonitorCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(monitorCrossInnerSVG).
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

func Move(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moveInnerSVG).
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

func NorthStar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(northStarInnerSVG).
		Element(children...)
}

func Notes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notesInnerSVG).
		Element(children...)
}

func NotesCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notesCrossInnerSVG).
		Element(children...)
}

func NotesTick(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notesTickInnerSVG).
		Element(children...)
}

func Nut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nutInnerSVG).
		Element(children...)
}

func Octagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(octagonInnerSVG).
		Element(children...)
}

func OctagonWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(octagonWarningInnerSVG).
		Element(children...)
}

func Organisation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(organisationInnerSVG).
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

func Padlock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(padlockInnerSVG).
		Element(children...)
}

func PaperPlane(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperPlaneInnerSVG).
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

func People(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(peopleInnerSVG).
		Element(children...)
}

func Person(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(personInnerSVG).
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

func PhoneCall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneCallInnerSVG).
		Element(children...)
}

func PhoneCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneCrossInnerSVG).
		Element(children...)
}

func PhoneForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneForwardInnerSVG).
		Element(children...)
}

func PhoneIncoming(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneIncomingInnerSVG).
		Element(children...)
}

func PhoneOutgoing(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneOutgoingInnerSVG).
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

func PlantPot(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plantPotInnerSVG).
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

func Power(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerInnerSVG).
		Element(children...)
}

func Printer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(printerInnerSVG).
		Element(children...)
}

func Pulse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pulseInnerSVG).
		Element(children...)
}

func Quote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quoteInnerSVG).
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

func Robot(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(robotInnerSVG).
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

func RotateAntiClockwise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateAntiClockwiseInnerSVG).
		Element(children...)
}

func RotateClockwise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateClockwiseInnerSVG).
		Element(children...)
}

func Scales(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scalesInnerSVG).
		Element(children...)
}

func ScreenMaximise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenMaximiseInnerSVG).
		Element(children...)
}

func ScreenMinimise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(screenMinimiseInnerSVG).
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

func ShieldCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldCrossInnerSVG).
		Element(children...)
}

func ShieldKeyhole(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldKeyholeInnerSVG).
		Element(children...)
}

func ShieldTick(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldTickInnerSVG).
		Element(children...)
}

func ShieldWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldWarningInnerSVG).
		Element(children...)
}

func ShoppingBag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shoppingBagInnerSVG).
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

func Signpost(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signpostInnerSVG).
		Element(children...)
}

func Skull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skullInnerSVG).
		Element(children...)
}

func Snowflake(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snowflakeInnerSVG).
		Element(children...)
}

func SoundDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundDownInnerSVG).
		Element(children...)
}

func SoundMute(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundMuteInnerSVG).
		Element(children...)
}

func SoundUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundUpInnerSVG).
		Element(children...)
}

func Speaker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speakerInnerSVG).
		Element(children...)
}

func Square(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareInnerSVG).
		Element(children...)
}

func SquareCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareCrossInnerSVG).
		Element(children...)
}

func SquareTick(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(squareTickInnerSVG).
		Element(children...)
}

func Stack(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stackInnerSVG).
		Element(children...)
}

func StackPop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stackPopInnerSVG).
		Element(children...)
}

func StackPush(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stackPushInnerSVG).
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

func StickyNote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stickyNoteInnerSVG).
		Element(children...)
}

func Sun(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sunInnerSVG).
		Element(children...)
}

func SwapHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swapHorizontalInnerSVG).
		Element(children...)
}

func SwapVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swapVerticalInnerSVG).
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

func Swords(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swordsInnerSVG).
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

func Telescope(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(telescopeInnerSVG).
		Element(children...)
}

func Tent(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tentInnerSVG).
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

func ThumbDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbDownInnerSVG).
		Element(children...)
}

func ThumbUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thumbUpInnerSVG).
		Element(children...)
}

func Tick(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tickInnerSVG).
		Element(children...)
}

func TickDouble(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tickDoubleInnerSVG).
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

func TreeFir(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(treeFirInnerSVG).
		Element(children...)
}

func Triangle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triangleInnerSVG).
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

func Wifi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiInnerSVG).
		Element(children...)
}

func WifiFair(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiFairInnerSVG).
		Element(children...)
}

func WifiPoor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiPoorInnerSVG).
		Element(children...)
}

func WifiSlash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiSlashInnerSVG).
		Element(children...)
}

func WifiWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiWarningInnerSVG).
		Element(children...)
}

func ZoomIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
			engine.NewAttribute("viewBox", "0 0 0 0"),
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
	case "anchor":
		return Anchor(), nil
	case "apps":
		return Apps(), nil
	case "apps-minus":
		return AppsMinus(), nil
	case "apps-plus":
		return AppsPlus(), nil
	case "archive":
		return Archive(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-down-left":
		return ArrowDownLeft(), nil
	case "arrow-down-right":
		return ArrowDownRight(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrow-up-left":
		return ArrowUpLeft(), nil
	case "arrow-up-right":
		return ArrowUpRight(), nil
	case "at-sign":
		return AtSign(), nil
	case "atom":
		return Atom(), nil
	case "bell":
		return Bell(), nil
	case "bell-slash":
		return BellSlash(), nil
	case "bin":
		return Bin(), nil
	case "binary":
		return Binary(), nil
	case "block":
		return Block(), nil
	case "bluetooth":
		return Bluetooth(), nil
	case "bluetooth-connected":
		return BluetoothConnected(), nil
	case "bluetooth-searching":
		return BluetoothSearching(), nil
	case "bluetooth-slash":
		return BluetoothSlash(), nil
	case "book":
		return Book(), nil
	case "book-open":
		return BookOpen(), nil
	case "bookmark":
		return Bookmark(), nil
	case "briefcase":
		return Briefcase(), nil
	case "bug":
		return Bug(), nil
	case "calendar":
		return Calendar(), nil
	case "camera":
		return Camera(), nil
	case "camera-video":
		return CameraVideo(), nil
	case "camera-video-slash":
		return CameraVideoSlash(), nil
	case "candy":
		return Candy(), nil
	case "cards":
		return Cards(), nil
	case "cast":
		return Cast(), nil
	case "certificate":
		return Certificate(), nil
	case "chart-bar":
		return ChartBar(), nil
	case "chart-line":
		return ChartLine(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "chevrons-down":
		return ChevronsDown(), nil
	case "chevrons-left":
		return ChevronsLeft(), nil
	case "chevrons-right":
		return ChevronsRight(), nil
	case "chevrons-up":
		return ChevronsUp(), nil
	case "chevrons-up-down":
		return ChevronsUpDown(), nil
	case "chip":
		return Chip(), nil
	case "circle":
		return Circle(), nil
	case "circle-cross":
		return CircleCross(), nil
	case "circle-minus":
		return CircleMinus(), nil
	case "circle-tick":
		return CircleTick(), nil
	case "circle-warning":
		return CircleWarning(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clipboard-tick":
		return ClipboardTick(), nil
	case "clock":
		return Clock(), nil
	case "clock-alarm":
		return ClockAlarm(), nil
	case "cloud":
		return Cloud(), nil
	case "clover":
		return Clover(), nil
	case "code":
		return Code(), nil
	case "coffee":
		return Coffee(), nil
	case "cog":
		return Cog(), nil
	case "compass":
		return Compass(), nil
	case "conical-flask":
		return ConicalFlask(), nil
	case "container":
		return Container(), nil
	case "copy":
		return Copy(), nil
	case "copyleft":
		return Copyleft(), nil
	case "copyright":
		return Copyright(), nil
	case "credit-card":
		return CreditCard(), nil
	case "crop":
		return Crop(), nil
	case "cross":
		return Cross(), nil
	case "crosshair":
		return Crosshair(), nil
	case "cube":
		return Cube(), nil
	case "cursor":
		return Cursor(), nil
	case "database":
		return Database(), nil
	case "diamond":
		return Diamond(), nil
	case "diff":
		return Diff(), nil
	case "disc":
		return Disc(), nil
	case "download":
		return Download(), nil
	case "droplet":
		return Droplet(), nil
	case "eraser":
		return Eraser(), nil
	case "extensions":
		return Extensions(), nil
	case "eye":
		return Eye(), nil
	case "eye-slash":
		return EyeSlash(), nil
	case "face-frown":
		return FaceFrown(), nil
	case "face-neutral":
		return FaceNeutral(), nil
	case "face-smile":
		return FaceSmile(), nil
	case "file":
		return File(), nil
	case "file-binary":
		return FileBinary(), nil
	case "file-code":
		return FileCode(), nil
	case "file-symlink":
		return FileSymlink(), nil
	case "files":
		return Files(), nil
	case "filter":
		return Filter(), nil
	case "flag":
		return Flag(), nil
	case "flame":
		return Flame(), nil
	case "floppy-disk":
		return FloppyDisk(), nil
	case "folder":
		return Folder(), nil
	case "folder-symlink":
		return FolderSymlink(), nil
	case "folders":
		return Folders(), nil
	case "forward":
		return Forward(), nil
	case "gamepad":
		return Gamepad(), nil
	case "gem":
		return Gem(), nil
	case "gift":
		return Gift(), nil
	case "git-branch":
		return GitBranch(), nil
	case "git-cherry-pick":
		return GitCherryPick(), nil
	case "git-commit":
		return GitCommit(), nil
	case "git-compare":
		return GitCompare(), nil
	case "git-fork":
		return GitFork(), nil
	case "git-merge":
		return GitMerge(), nil
	case "git-request":
		return GitRequest(), nil
	case "git-request-cross":
		return GitRequestCross(), nil
	case "git-request-draft":
		return GitRequestDraft(), nil
	case "github":
		return Github(), nil
	case "gitlab":
		return Gitlab(), nil
	case "glasses":
		return Glasses(), nil
	case "globe":
		return Globe(), nil
	case "grab-horizontal":
		return GrabHorizontal(), nil
	case "grab-vertical":
		return GrabVertical(), nil
	case "graduate-cap":
		return GraduateCap(), nil
	case "hash":
		return Hash(), nil
	case "headphones":
		return Headphones(), nil
	case "heart":
		return Heart(), nil
	case "help":
		return Help(), nil
	case "hexagon":
		return Hexagon(), nil
	case "home":
		return Home(), nil
	case "hourglass":
		return Hourglass(), nil
	case "id":
		return Id(), nil
	case "image":
		return Image(), nil
	case "inbox":
		return Inbox(), nil
	case "infinity":
		return Infinity(), nil
	case "info":
		return Info(), nil
	case "key":
		return Key(), nil
	case "laptop":
		return Laptop(), nil
	case "layout-columns":
		return LayoutColumns(), nil
	case "layout-dashboard":
		return LayoutDashboard(), nil
	case "layout-grid":
		return LayoutGrid(), nil
	case "layout-list":
		return LayoutList(), nil
	case "layout-rows":
		return LayoutRows(), nil
	case "layout-sidebar":
		return LayoutSidebar(), nil
	case "layout-stack-h":
		return LayoutStackH(), nil
	case "layout-stack-v":
		return LayoutStackV(), nil
	case "lightbulb":
		return Lightbulb(), nil
	case "lightning-bolt":
		return LightningBolt(), nil
	case "link":
		return Link(), nil
	case "link-external":
		return LinkExternal(), nil
	case "link-slash":
		return LinkSlash(), nil
	case "mail":
		return Mail(), nil
	case "map":
		return Map(), nil
	case "map-pin":
		return MapPin(), nil
	case "media-back":
		return MediaBack(), nil
	case "media-eject":
		return MediaEject(), nil
	case "media-fast-forward":
		return MediaFastForward(), nil
	case "media-pause":
		return MediaPause(), nil
	case "media-play":
		return MediaPlay(), nil
	case "media-rewind":
		return MediaRewind(), nil
	case "media-skip":
		return MediaSkip(), nil
	case "menu-hamburger":
		return MenuHamburger(), nil
	case "menu-kebab":
		return MenuKebab(), nil
	case "menu-meatball":
		return MenuMeatball(), nil
	case "message":
		return Message(), nil
	case "messages":
		return Messages(), nil
	case "microphone":
		return Microphone(), nil
	case "minus":
		return Minus(), nil
	case "mobile":
		return Mobile(), nil
	case "monitor":
		return Monitor(), nil
	case "monitor-arrow":
		return MonitorArrow(), nil
	case "monitor-cross":
		return MonitorCross(), nil
	case "moon":
		return Moon(), nil
	case "move":
		return Move(), nil
	case "music":
		return Music(), nil
	case "newspaper":
		return Newspaper(), nil
	case "north-star":
		return NorthStar(), nil
	case "notes":
		return Notes(), nil
	case "notes-cross":
		return NotesCross(), nil
	case "notes-tick":
		return NotesTick(), nil
	case "nut":
		return Nut(), nil
	case "octagon":
		return Octagon(), nil
	case "octagon-warning":
		return OctagonWarning(), nil
	case "organisation":
		return Organisation(), nil
	case "package":
		return Package(), nil
	case "padlock":
		return Padlock(), nil
	case "paper-plane":
		return PaperPlane(), nil
	case "paperclip":
		return Paperclip(), nil
	case "pencil":
		return Pencil(), nil
	case "people":
		return People(), nil
	case "person":
		return Person(), nil
	case "phone":
		return Phone(), nil
	case "phone-call":
		return PhoneCall(), nil
	case "phone-cross":
		return PhoneCross(), nil
	case "phone-forward":
		return PhoneForward(), nil
	case "phone-incoming":
		return PhoneIncoming(), nil
	case "phone-outgoing":
		return PhoneOutgoing(), nil
	case "pin":
		return Pin(), nil
	case "plant-pot":
		return PlantPot(), nil
	case "plus":
		return Plus(), nil
	case "power":
		return Power(), nil
	case "printer":
		return Printer(), nil
	case "pulse":
		return Pulse(), nil
	case "quote":
		return Quote(), nil
	case "refresh":
		return Refresh(), nil
	case "reply":
		return Reply(), nil
	case "robot":
		return Robot(), nil
	case "rocket":
		return Rocket(), nil
	case "rotate-anti-clockwise":
		return RotateAntiClockwise(), nil
	case "rotate-clockwise":
		return RotateClockwise(), nil
	case "scales":
		return Scales(), nil
	case "screen-maximise":
		return ScreenMaximise(), nil
	case "screen-minimise":
		return ScreenMinimise(), nil
	case "search":
		return Search(), nil
	case "server":
		return Server(), nil
	case "share":
		return Share(), nil
	case "shield":
		return Shield(), nil
	case "shield-cross":
		return ShieldCross(), nil
	case "shield-keyhole":
		return ShieldKeyhole(), nil
	case "shield-tick":
		return ShieldTick(), nil
	case "shield-warning":
		return ShieldWarning(), nil
	case "shopping-bag":
		return ShoppingBag(), nil
	case "sign-in":
		return SignIn(), nil
	case "sign-out":
		return SignOut(), nil
	case "signpost":
		return Signpost(), nil
	case "skull":
		return Skull(), nil
	case "snowflake":
		return Snowflake(), nil
	case "sound-down":
		return SoundDown(), nil
	case "sound-mute":
		return SoundMute(), nil
	case "sound-up":
		return SoundUp(), nil
	case "speaker":
		return Speaker(), nil
	case "square":
		return Square(), nil
	case "square-cross":
		return SquareCross(), nil
	case "square-tick":
		return SquareTick(), nil
	case "stack":
		return Stack(), nil
	case "stack-pop":
		return StackPop(), nil
	case "stack-push":
		return StackPush(), nil
	case "star":
		return Star(), nil
	case "sticky-note":
		return StickyNote(), nil
	case "sun":
		return Sun(), nil
	case "swap-horizontal":
		return SwapHorizontal(), nil
	case "swap-vertical":
		return SwapVertical(), nil
	case "sword":
		return Sword(), nil
	case "swords":
		return Swords(), nil
	case "tablet":
		return Tablet(), nil
	case "tag":
		return Tag(), nil
	case "telescope":
		return Telescope(), nil
	case "tent":
		return Tent(), nil
	case "terminal":
		return Terminal(), nil
	case "thumb-down":
		return ThumbDown(), nil
	case "thumb-up":
		return ThumbUp(), nil
	case "tick":
		return Tick(), nil
	case "tick-double":
		return TickDouble(), nil
	case "ticket":
		return Ticket(), nil
	case "tree-fir":
		return TreeFir(), nil
	case "triangle":
		return Triangle(), nil
	case "trophy":
		return Trophy(), nil
	case "umbrella":
		return Umbrella(), nil
	case "upload":
		return Upload(), nil
	case "wifi":
		return Wifi(), nil
	case "wifi-fair":
		return WifiFair(), nil
	case "wifi-poor":
		return WifiPoor(), nil
	case "wifi-slash":
		return WifiSlash(), nil
	case "wifi-warning":
		return WifiWarning(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in charm icon set", name)
	}
}
