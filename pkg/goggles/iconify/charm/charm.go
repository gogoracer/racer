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

var sharedIconAttrs = engine.Attrs{"width": "1em", "height": "1em"}

func Anchor(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		anchorInnerSVG,
		children,
	)
}

func Apps(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		appsInnerSVG,
		children,
	)
}

func AppsMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		appsMinusInnerSVG,
		children,
	)
}

func AppsPlus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		appsPlusInnerSVG,
		children,
	)
}

func Archive(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		archiveInnerSVG,
		children,
	)
}

func ArrowDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		arrowDownInnerSVG,
		children,
	)
}

func ArrowDownLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		arrowDownLeftInnerSVG,
		children,
	)
}

func ArrowDownRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		arrowDownRightInnerSVG,
		children,
	)
}

func ArrowLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		arrowRightInnerSVG,
		children,
	)
}

func ArrowUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		arrowUpInnerSVG,
		children,
	)
}

func ArrowUpLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		arrowUpLeftInnerSVG,
		children,
	)
}

func ArrowUpRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		arrowUpRightInnerSVG,
		children,
	)
}

func AtSign(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		atSignInnerSVG,
		children,
	)
}

func Atom(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		atomInnerSVG,
		children,
	)
}

func Bell(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		bellInnerSVG,
		children,
	)
}

func BellSlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		bellSlashInnerSVG,
		children,
	)
}

func Bin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		binInnerSVG,
		children,
	)
}

func Binary(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		binaryInnerSVG,
		children,
	)
}

func Block(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		blockInnerSVG,
		children,
	)
}

func Bluetooth(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		bluetoothInnerSVG,
		children,
	)
}

func BluetoothConnected(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		bluetoothConnectedInnerSVG,
		children,
	)
}

func BluetoothSearching(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		bluetoothSearchingInnerSVG,
		children,
	)
}

func BluetoothSlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		bluetoothSlashInnerSVG,
		children,
	)
}

func Book(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		bookmarkInnerSVG,
		children,
	)
}

func Briefcase(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		briefcaseInnerSVG,
		children,
	)
}

func Bug(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		bugInnerSVG,
		children,
	)
}

func Calendar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		cameraInnerSVG,
		children,
	)
}

func CameraVideo(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cameraVideoInnerSVG,
		children,
	)
}

func CameraVideoSlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cameraVideoSlashInnerSVG,
		children,
	)
}

func Candy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		candyInnerSVG,
		children,
	)
}

func Cards(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cardsInnerSVG,
		children,
	)
}

func Cast(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		castInnerSVG,
		children,
	)
}

func Certificate(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		certificateInnerSVG,
		children,
	)
}

func ChartBar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		chartBarInnerSVG,
		children,
	)
}

func ChartLine(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		chartLineInnerSVG,
		children,
	)
}

func ChevronDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		chevronUpInnerSVG,
		children,
	)
}

func ChevronsDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		chevronsDownInnerSVG,
		children,
	)
}

func ChevronsLeft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		chevronsLeftInnerSVG,
		children,
	)
}

func ChevronsRight(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		chevronsRightInnerSVG,
		children,
	)
}

func ChevronsUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		chevronsUpInnerSVG,
		children,
	)
}

func ChevronsUpDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		chevronsUpDownInnerSVG,
		children,
	)
}

func Chip(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		chipInnerSVG,
		children,
	)
}

func Circle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		circleInnerSVG,
		children,
	)
}

func CircleCross(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		circleCrossInnerSVG,
		children,
	)
}

func CircleMinus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		circleMinusInnerSVG,
		children,
	)
}

func CircleTick(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		circleTickInnerSVG,
		children,
	)
}

func CircleWarning(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		circleWarningInnerSVG,
		children,
	)
}

func Clipboard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		clipboardInnerSVG,
		children,
	)
}

func ClipboardTick(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		clipboardTickInnerSVG,
		children,
	)
}

func Clock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		clockInnerSVG,
		children,
	)
}

func ClockAlarm(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		clockAlarmInnerSVG,
		children,
	)
}

func Cloud(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cloudInnerSVG,
		children,
	)
}

func Clover(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cloverInnerSVG,
		children,
	)
}

func Code(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		coffeeInnerSVG,
		children,
	)
}

func Cog(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cogInnerSVG,
		children,
	)
}

func Compass(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		compassInnerSVG,
		children,
	)
}

func ConicalFlask(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		conicalFlaskInnerSVG,
		children,
	)
}

func Container(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		containerInnerSVG,
		children,
	)
}

func Copy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		copyInnerSVG,
		children,
	)
}

func Copyleft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		copyleftInnerSVG,
		children,
	)
}

func Copyright(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		copyrightInnerSVG,
		children,
	)
}

func CreditCard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		creditCardInnerSVG,
		children,
	)
}

func Crop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cropInnerSVG,
		children,
	)
}

func Cross(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		crossInnerSVG,
		children,
	)
}

func Crosshair(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		crosshairInnerSVG,
		children,
	)
}

func Cube(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cubeInnerSVG,
		children,
	)
}

func Cursor(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		cursorInnerSVG,
		children,
	)
}

func Database(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		databaseInnerSVG,
		children,
	)
}

func Diamond(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		diamondInnerSVG,
		children,
	)
}

func Diff(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		diffInnerSVG,
		children,
	)
}

func Disc(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		discInnerSVG,
		children,
	)
}

func Download(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		downloadInnerSVG,
		children,
	)
}

func Droplet(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		dropletInnerSVG,
		children,
	)
}

func Eraser(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		eraserInnerSVG,
		children,
	)
}

func Extensions(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		extensionsInnerSVG,
		children,
	)
}

func Eye(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		eyeInnerSVG,
		children,
	)
}

func EyeSlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		eyeSlashInnerSVG,
		children,
	)
}

func FaceFrown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		faceFrownInnerSVG,
		children,
	)
}

func FaceNeutral(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		faceNeutralInnerSVG,
		children,
	)
}

func FaceSmile(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		faceSmileInnerSVG,
		children,
	)
}

func File(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		fileInnerSVG,
		children,
	)
}

func FileBinary(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		fileBinaryInnerSVG,
		children,
	)
}

func FileCode(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		fileCodeInnerSVG,
		children,
	)
}

func FileSymlink(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		fileSymlinkInnerSVG,
		children,
	)
}

func Files(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		filesInnerSVG,
		children,
	)
}

func Filter(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		filterInnerSVG,
		children,
	)
}

func Flag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		flagInnerSVG,
		children,
	)
}

func Flame(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		flameInnerSVG,
		children,
	)
}

func FloppyDisk(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		floppyDiskInnerSVG,
		children,
	)
}

func Folder(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		folderInnerSVG,
		children,
	)
}

func FolderSymlink(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		folderSymlinkInnerSVG,
		children,
	)
}

func Folders(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		foldersInnerSVG,
		children,
	)
}

func Forward(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		forwardInnerSVG,
		children,
	)
}

func Gamepad(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gamepadInnerSVG,
		children,
	)
}

func Gem(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gemInnerSVG,
		children,
	)
}

func Gift(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		gitBranchInnerSVG,
		children,
	)
}

func GitCherryPick(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitCherryPickInnerSVG,
		children,
	)
}

func GitCommit(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitCommitInnerSVG,
		children,
	)
}

func GitCompare(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitCompareInnerSVG,
		children,
	)
}

func GitFork(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitForkInnerSVG,
		children,
	)
}

func GitMerge(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitMergeInnerSVG,
		children,
	)
}

func GitRequest(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitRequestInnerSVG,
		children,
	)
}

func GitRequestCross(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitRequestCrossInnerSVG,
		children,
	)
}

func GitRequestDraft(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitRequestDraftInnerSVG,
		children,
	)
}

func Github(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		githubInnerSVG,
		children,
	)
}

func Gitlab(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		gitlabInnerSVG,
		children,
	)
}

func Glasses(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		glassesInnerSVG,
		children,
	)
}

func Globe(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		globeInnerSVG,
		children,
	)
}

func GrabHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		grabHorizontalInnerSVG,
		children,
	)
}

func GrabVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		grabVerticalInnerSVG,
		children,
	)
}

func GraduateCap(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		graduateCapInnerSVG,
		children,
	)
}

func Hash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		hashInnerSVG,
		children,
	)
}

func Headphones(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		headphonesInnerSVG,
		children,
	)
}

func Heart(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		heartInnerSVG,
		children,
	)
}

func Help(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		helpInnerSVG,
		children,
	)
}

func Hexagon(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		hexagonInnerSVG,
		children,
	)
}

func Home(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		hourglassInnerSVG,
		children,
	)
}

func Id(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		idInnerSVG,
		children,
	)
}

func Image(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		imageInnerSVG,
		children,
	)
}

func Inbox(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		inboxInnerSVG,
		children,
	)
}

func Infinity(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		infinityInnerSVG,
		children,
	)
}

func Info(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		infoInnerSVG,
		children,
	)
}

func Key(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		keyInnerSVG,
		children,
	)
}

func Laptop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		laptopInnerSVG,
		children,
	)
}

func LayoutColumns(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		layoutColumnsInnerSVG,
		children,
	)
}

func LayoutDashboard(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		layoutDashboardInnerSVG,
		children,
	)
}

func LayoutGrid(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		layoutGridInnerSVG,
		children,
	)
}

func LayoutList(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		layoutListInnerSVG,
		children,
	)
}

func LayoutRows(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		layoutRowsInnerSVG,
		children,
	)
}

func LayoutSidebar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		layoutSidebarInnerSVG,
		children,
	)
}

func LayoutStackH(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		layoutStackHInnerSVG,
		children,
	)
}

func LayoutStackV(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		layoutStackVInnerSVG,
		children,
	)
}

func Lightbulb(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		lightbulbInnerSVG,
		children,
	)
}

func LightningBolt(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		lightningBoltInnerSVG,
		children,
	)
}

func Link(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		linkInnerSVG,
		children,
	)
}

func LinkExternal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		linkExternalInnerSVG,
		children,
	)
}

func LinkSlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		linkSlashInnerSVG,
		children,
	)
}

func Mail(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mailInnerSVG,
		children,
	)
}

func Map(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mapInnerSVG,
		children,
	)
}

func MapPin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mapPinInnerSVG,
		children,
	)
}

func MediaBack(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mediaBackInnerSVG,
		children,
	)
}

func MediaEject(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mediaEjectInnerSVG,
		children,
	)
}

func MediaFastForward(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mediaFastForwardInnerSVG,
		children,
	)
}

func MediaPause(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mediaPauseInnerSVG,
		children,
	)
}

func MediaPlay(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mediaPlayInnerSVG,
		children,
	)
}

func MediaRewind(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mediaRewindInnerSVG,
		children,
	)
}

func MediaSkip(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		mediaSkipInnerSVG,
		children,
	)
}

func MenuHamburger(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		menuHamburgerInnerSVG,
		children,
	)
}

func MenuKebab(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		menuKebabInnerSVG,
		children,
	)
}

func MenuMeatball(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		menuMeatballInnerSVG,
		children,
	)
}

func Message(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		messageInnerSVG,
		children,
	)
}

func Messages(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		messagesInnerSVG,
		children,
	)
}

func Microphone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		mobileInnerSVG,
		children,
	)
}

func Monitor(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		monitorInnerSVG,
		children,
	)
}

func MonitorArrow(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		monitorArrowInnerSVG,
		children,
	)
}

func MonitorCross(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		monitorCrossInnerSVG,
		children,
	)
}

func Moon(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		musicInnerSVG,
		children,
	)
}

func Newspaper(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		newspaperInnerSVG,
		children,
	)
}

func NorthStar(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		northStarInnerSVG,
		children,
	)
}

func Notes(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		notesInnerSVG,
		children,
	)
}

func NotesCross(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		notesCrossInnerSVG,
		children,
	)
}

func NotesTick(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		notesTickInnerSVG,
		children,
	)
}

func Nut(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		nutInnerSVG,
		children,
	)
}

func Octagon(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		octagonInnerSVG,
		children,
	)
}

func OctagonWarning(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		octagonWarningInnerSVG,
		children,
	)
}

func Organisation(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		organisationInnerSVG,
		children,
	)
}

func Package(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		packageInnerSVG,
		children,
	)
}

func Padlock(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		padlockInnerSVG,
		children,
	)
}

func PaperPlane(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		paperPlaneInnerSVG,
		children,
	)
}

func Paperclip(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		paperclipInnerSVG,
		children,
	)
}

func Pencil(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		pencilInnerSVG,
		children,
	)
}

func People(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		peopleInnerSVG,
		children,
	)
}

func Person(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		personInnerSVG,
		children,
	)
}

func Phone(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		phoneInnerSVG,
		children,
	)
}

func PhoneCall(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		phoneCallInnerSVG,
		children,
	)
}

func PhoneCross(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		phoneCrossInnerSVG,
		children,
	)
}

func PhoneForward(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		phoneForwardInnerSVG,
		children,
	)
}

func PhoneIncoming(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		phoneIncomingInnerSVG,
		children,
	)
}

func PhoneOutgoing(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		phoneOutgoingInnerSVG,
		children,
	)
}

func Pin(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		pinInnerSVG,
		children,
	)
}

func PlantPot(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		plantPotInnerSVG,
		children,
	)
}

func Plus(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		powerInnerSVG,
		children,
	)
}

func Printer(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		printerInnerSVG,
		children,
	)
}

func Pulse(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		pulseInnerSVG,
		children,
	)
}

func Quote(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		quoteInnerSVG,
		children,
	)
}

func Refresh(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		refreshInnerSVG,
		children,
	)
}

func Reply(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		replyInnerSVG,
		children,
	)
}

func Robot(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		robotInnerSVG,
		children,
	)
}

func Rocket(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		rocketInnerSVG,
		children,
	)
}

func RotateAntiClockwise(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		rotateAntiClockwiseInnerSVG,
		children,
	)
}

func RotateClockwise(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		rotateClockwiseInnerSVG,
		children,
	)
}

func Scales(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		scalesInnerSVG,
		children,
	)
}

func ScreenMaximise(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		screenMaximiseInnerSVG,
		children,
	)
}

func ScreenMinimise(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		screenMinimiseInnerSVG,
		children,
	)
}

func Search(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		searchInnerSVG,
		children,
	)
}

func Server(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		serverInnerSVG,
		children,
	)
}

func Share(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		shareInnerSVG,
		children,
	)
}

func Shield(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		shieldInnerSVG,
		children,
	)
}

func ShieldCross(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		shieldCrossInnerSVG,
		children,
	)
}

func ShieldKeyhole(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		shieldKeyholeInnerSVG,
		children,
	)
}

func ShieldTick(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		shieldTickInnerSVG,
		children,
	)
}

func ShieldWarning(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		shieldWarningInnerSVG,
		children,
	)
}

func ShoppingBag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		shoppingBagInnerSVG,
		children,
	)
}

func SignIn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		signOutInnerSVG,
		children,
	)
}

func Signpost(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		signpostInnerSVG,
		children,
	)
}

func Skull(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		skullInnerSVG,
		children,
	)
}

func Snowflake(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		snowflakeInnerSVG,
		children,
	)
}

func SoundDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		soundDownInnerSVG,
		children,
	)
}

func SoundMute(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		soundMuteInnerSVG,
		children,
	)
}

func SoundUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		soundUpInnerSVG,
		children,
	)
}

func Speaker(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		speakerInnerSVG,
		children,
	)
}

func Square(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		squareInnerSVG,
		children,
	)
}

func SquareCross(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		squareCrossInnerSVG,
		children,
	)
}

func SquareTick(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		squareTickInnerSVG,
		children,
	)
}

func Stack(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		stackInnerSVG,
		children,
	)
}

func StackPop(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		stackPopInnerSVG,
		children,
	)
}

func StackPush(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		stackPushInnerSVG,
		children,
	)
}

func Star(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		starInnerSVG,
		children,
	)
}

func StickyNote(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		stickyNoteInnerSVG,
		children,
	)
}

func Sun(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		sunInnerSVG,
		children,
	)
}

func SwapHorizontal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		swapHorizontalInnerSVG,
		children,
	)
}

func SwapVertical(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		swapVerticalInnerSVG,
		children,
	)
}

func Sword(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		swordInnerSVG,
		children,
	)
}

func Swords(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		swordsInnerSVG,
		children,
	)
}

func Tablet(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		tabletInnerSVG,
		children,
	)
}

func Tag(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		tagInnerSVG,
		children,
	)
}

func Telescope(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		telescopeInnerSVG,
		children,
	)
}

func Tent(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		tentInnerSVG,
		children,
	)
}

func Terminal(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		terminalInnerSVG,
		children,
	)
}

func ThumbDown(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		thumbDownInnerSVG,
		children,
	)
}

func ThumbUp(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		thumbUpInnerSVG,
		children,
	)
}

func Tick(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		tickInnerSVG,
		children,
	)
}

func TickDouble(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		tickDoubleInnerSVG,
		children,
	)
}

func Ticket(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		ticketInnerSVG,
		children,
	)
}

func TreeFir(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		treeFirInnerSVG,
		children,
	)
}

func Triangle(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		triangleInnerSVG,
		children,
	)
}

func Trophy(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		trophyInnerSVG,
		children,
	)
}

func Umbrella(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		umbrellaInnerSVG,
		children,
	)
}

func Upload(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		uploadInnerSVG,
		children,
	)
}

func Wifi(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		wifiInnerSVG,
		children,
	)
}

func WifiFair(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		wifiFairInnerSVG,
		children,
	)
}

func WifiPoor(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		wifiPoorInnerSVG,
		children,
	)
}

func WifiSlash(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		wifiSlashInnerSVG,
		children,
	)
}

func WifiWarning(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
		},
		wifiWarningInnerSVG,
		children,
	)
}

func ZoomIn(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 0 0",
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
			"viewBox": "0 0 0 0",
		},
		zoomOutInnerSVG,
		children,
	)
}

func ByName(name string) (*engine.HTMLElement, error) {
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
