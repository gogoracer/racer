package ei

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	archiveInnerSVG         = `<path fill="currentColor" d="M42 20h-2v-5c0-.6-.4-1-1-1H11c-.6 0-1 .4-1 1v5H8v-5c0-1.7 1.3-3 3-3h28c1.7 0 3 1.3 3 3v5zm-5 20H13c-1.7 0-3-1.3-3-3V20h2v17c0 .6.4 1 1 1h24c.6 0 1-.4 1-1V20h2v17c0 1.7-1.3 3-3 3z"/><path fill="currentColor" d="M29 26h-8c-.6 0-1-.4-1-1s.4-1 1-1h8c.6 0 1 .4 1 1s-.4 1-1 1zM8 18h34v2H8z"/>`
	arrowDownInnerSVG       = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="m25 34.4l-9.7-9.7l1.4-1.4l8.3 8.3l8.3-8.3l1.4 1.4z"/><path fill="currentColor" d="M24 16h2v17h-2z"/>`
	arrowLeftInnerSVG       = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="M25.3 34.7L15.6 25l9.7-9.7l1.4 1.4l-8.3 8.3l8.3 8.3z"/><path fill="currentColor" d="M17 24h17v2H17z"/>`
	arrowRightInnerSVG      = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="m24.7 34.7l-1.4-1.4l8.3-8.3l-8.3-8.3l1.4-1.4l9.7 9.7z"/><path fill="currentColor" d="M16 24h17v2H16z"/>`
	arrowUpInnerSVG         = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="M33.3 26.7L25 18.4l-8.3 8.3l-1.4-1.4l9.7-9.7l9.7 9.7z"/><path fill="currentColor" d="M24 17h2v17h-2z"/>`
	bellInnerSVG            = `<path fill="currentColor" d="M42 36c-6.5 0-7.4-6.3-8.2-11.9C32.9 17.9 32.1 12 25 12s-7.9 5.9-8.8 12.1C15.4 29.7 14.5 36 8 36v-2c4.6 0 5.3-3.9 6.2-10.1c.9-6.2 2-13.9 10.8-13.9s9.9 7.7 10.8 13.9C36.7 30.1 37.4 34 42 34v2z"/><path fill="currentColor" d="M25 40c-2.8 0-5-2.2-5-5h2c0 1.7 1.3 3 3 3s3-1.3 3-3h2c0 2.8-2.2 5-5 5z"/><path fill="currentColor" d="M8 34h34v2H8zm19-24c0 1.1-.9 1.5-2 1.5s-2-.4-2-1.5s.9-2 2-2s2 .9 2 2z"/>`
	calendarInnerSVG        = `<path fill="currentColor" d="M37 38H13c-1.7 0-3-1.3-3-3V13c0-1.7 1.1-3 2.5-3H14v2h-1.5c-.2 0-.5.4-.5 1v22c0 .6.4 1 1 1h24c.6 0 1-.4 1-1V13c0-.6-.3-1-.5-1H36v-2h1.5c1.4 0 2.5 1.3 2.5 3v22c0 1.7-1.3 3-3 3z"/><path fill="currentColor" d="M17 14c-.6 0-1-.4-1-1V9c0-.6.4-1 1-1s1 .4 1 1v4c0 .6-.4 1-1 1zm16 0c-.6 0-1-.4-1-1V9c0-.6.4-1 1-1s1 .4 1 1v4c0 .6-.4 1-1 1zm-13-4h10v2H20zm-8 6h26v2H12zm22 4h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm16 4h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm20 4h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm16 4h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2z"/>`
	cameraInnerSVG          = `<path fill="currentColor" d="M39 38H11c-1.7 0-3-1.3-3-3V17c0-1.7 1.3-3 3-3h6c.2 0 .5-.2.6-.3l1.1-2.2c.4-.8 1.4-1.4 2.3-1.4h8c.9 0 1.9.6 2.3 1.4l1.1 2.2c.1.2.4.3.6.3h6c1.7 0 3 1.3 3 3v18c0 1.7-1.3 3-3 3zM11 16c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h28c.6 0 1-.4 1-1V17c0-.6-.4-1-1-1h-6c-.9 0-1.9-.6-2.3-1.4l-1.1-2.2c-.1-.2-.4-.4-.6-.4h-8c-.2 0-.5.2-.6.3l-1.1 2.2c-.4.9-1.4 1.5-2.3 1.5h-6z"/><path fill="currentColor" d="M25 34c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9zm0-16c-3.9 0-7 3.1-7 7s3.1 7 7 7s7-3.1 7-7s-3.1-7-7-7z"/><circle cx="35" cy="18" r="1" fill="currentColor"/><path fill="currentColor" d="M12 12h4v1h-4zm13 9v-1c-2.8 0-5 2.2-5 5h1c0-2.2 1.8-4 4-4z"/>`
	cartInnerSVG            = `<path fill="currentColor" d="M35 34H13c-.3 0-.6-.2-.8-.4s-.2-.6-.1-.9l1.9-4.8L12.1 10H6V8h7c.5 0 .9.4 1 .9l2 19c0 .2 0 .3-.1.5L14.5 32H36l-1 2z"/><path fill="currentColor" d="m15.2 29l-.4-2L38 22.2V14H14v-2h25c.6 0 1 .4 1 1v10c0 .5-.3.9-.8 1l-24 5zM36 40c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4zm0-6c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm-24 6c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4zm0-6c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2z"/>`
	chartInnerSVG           = `<path fill="currentColor" d="M18 36h-2V26h-4v10h-2V24h8zm10 0h-2V20h-4v16h-2V18h8zm10 0h-2V14h-4v22h-2V12h8zM8 36h32v2H8z"/>`
	checkInnerSVG           = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="m23 32.4l-8.7-8.7l1.4-1.4l7.3 7.3l11.3-11.3l1.4 1.4z"/>`
	chevronDownInnerSVG     = `<path fill="currentColor" d="m25 32.4l-9.7-9.7l1.4-1.4l8.3 8.3l8.3-8.3l1.4 1.4z"/>`
	chevronLeftInnerSVG     = `<path fill="currentColor" d="M27.3 34.7L17.6 25l9.7-9.7l1.4 1.4l-8.3 8.3l8.3 8.3z"/>`
	chevronRightInnerSVG    = `<path fill="currentColor" d="m22.7 34.7l-1.4-1.4l8.3-8.3l-8.3-8.3l1.4-1.4l9.7 9.7z"/>`
	chevronUpInnerSVG       = `<path fill="currentColor" d="M33.3 28.7L25 20.4l-8.3 8.3l-1.4-1.4l9.7-9.7l9.7 9.7z"/>`
	clockInnerSVG           = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="M30.3 33.7L24 27.4V16h2v10.6l5.7 5.7z"/>`
	closeInnerSVG           = `<path fill="currentColor" d="m37.304 11.282l1.414 1.414l-26.022 26.02l-1.414-1.413z"/><path fill="currentColor" d="m12.696 11.282l26.022 26.02l-1.414 1.415l-26.022-26.02z"/>`
	closeOInnerSVG          = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="m32.283 16.302l1.414 1.415l-15.98 15.98l-1.414-1.414z"/><path fill="currentColor" d="m17.717 16.302l15.98 15.98l-1.414 1.415l-15.98-15.98z"/>`
	commentInnerSVG         = `<path fill="currentColor" d="M15 42h-2l1.2-1.6c.8-1.1 1.3-2.5 1.6-4.2C10.8 33.9 8 29.6 8 24c0-8.6 6.5-14 17-14s17 5.4 17 14c0 8.8-6.4 14-17 14h-.7c-1.6 1.9-4.4 4-9.3 4zm10-30c-9.4 0-15 4.5-15 12c0 6.4 3.9 9.4 7.2 10.7l.7.3l-.1.8c-.2 1.6-.5 3-1.1 4.2c3.3-.4 5.2-2.1 6.3-3.5l.3-.4H25c13.5 0 15-8.4 15-12C40 16.5 34.4 12 25 12z"/>`
	creditCardInnerSVG      = `<path fill="currentColor" d="M39 38H11c-1.7 0-3-1.3-3-3V15c0-1.7 1.3-3 3-3h28c1.7 0 3 1.3 3 3v20c0 1.7-1.3 3-3 3zM11 14c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h28c.6 0 1-.4 1-1V15c0-.6-.4-1-1-1H11z"/><path fill="currentColor" d="M9 16h32v6H9zm3 10h1v2h-1zm2 0h1v2h-1zm2 0h1v2h-1zm3 0h1v2h-1zm2 0h1v2h-1zm2 0h1v2h-1zm3 0h1v2h-1zm2 0h1v2h-1zm2 0h1v2h-1zm3 0h1v2h-1zm2 0h1v2h-1zm2 0h1v2h-1z"/>`
	envelopeInnerSVG        = `<path fill="currentColor" d="m31.796 24.244l9.97 9.97l-1.415 1.414l-9.97-9.97zm-13.518.043l1.414 1.414l-9.9 9.9l-1.414-1.41z" opacity=".9"/><path fill="currentColor" d="M25 29.9c-1.5 0-3.1-.6-4.2-1.8L8.3 15.7l1.4-1.4l12.5 12.5c1.6 1.6 4.1 1.6 5.7 0l12.5-12.5l1.4 1.4l-12.6 12.5c-1.1 1.1-2.7 1.7-4.2 1.7z"/><path fill="currentColor" d="M39 38H11c-1.7 0-3-1.3-3-3V15c0-1.7 1.3-3 3-3h28c1.7 0 3 1.3 3 3v20c0 1.7-1.3 3-3 3zM11 14c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h28c.6 0 1-.4 1-1V15c0-.6-.4-1-1-1H11z"/>`
	exclamationInnerSVG     = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="M24 32h2v2h-2zm1.6-2h-1.2l-.4-8v-6h2v6z"/>`
	externalLinkInnerSVG    = `<path fill="currentColor" d="m38.288 10.297l1.414 1.415l-14.99 14.99l-1.414-1.414z"/><path fill="currentColor" d="M40 20h-2v-8h-8v-2h10zm-5 18H15c-1.7 0-3-1.3-3-3V15c0-1.7 1.3-3 3-3h11v2H15c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h20c.6 0 1-.4 1-1V24h2v11c0 1.7-1.3 3-3 3z"/>`
	eyeInnerSVG             = `<path fill="currentColor" d="M25 36C13.5 36 8.3 25.9 8.1 25.4c-.1-.3-.1-.6 0-.9C8.3 24.1 13.5 14 25 14s16.7 10.1 16.9 10.6c.1.3.1.6 0 .9c-.2.4-5.4 10.5-16.9 10.5zM10.1 25c1.1 1.9 5.9 9 14.9 9s13.7-7.1 14.9-9c-1.1-1.9-5.9-9-14.9-9s-13.7 7.1-14.9 9z"/><path fill="currentColor" d="M25 34c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9zm0-16c-3.9 0-7 3.1-7 7s3.1 7 7 7s7-3.1 7-7s-3.1-7-7-7z"/><path fill="currentColor" d="M25 30c-2.8 0-5-2.2-5-5c0-.6.4-1 1-1s1 .4 1 1c0 1.7 1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3c-.6 0-1-.4-1-1s.4-1 1-1c2.8 0 5 2.2 5 5s-2.2 5-5 5z"/>`
	gearInnerSVG            = `<path fill="currentColor" d="M25 34c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9zm0-16c-3.9 0-7 3.1-7 7s3.1 7 7 7s7-3.1 7-7s-3.1-7-7-7z"/><path fill="currentColor" d="M27.7 44h-5.4l-1.5-4.6c-1-.3-2-.7-2.9-1.2l-4.4 2.2l-3.8-3.8l2.2-4.4c-.5-.9-.9-1.9-1.2-2.9L6 27.7v-5.4l4.6-1.5c.3-1 .7-2 1.2-2.9l-2.2-4.4l3.8-3.8l4.4 2.2c.9-.5 1.9-.9 2.9-1.2L22.3 6h5.4l1.5 4.6c1 .3 2 .7 2.9 1.2l4.4-2.2l3.8 3.8l-2.2 4.4c.5.9.9 1.9 1.2 2.9l4.6 1.5v5.4l-4.6 1.5c-.3 1-.7 2-1.2 2.9l2.2 4.4l-3.8 3.8l-4.4-2.2c-.9.5-1.9.9-2.9 1.2L27.7 44zm-4-2h2.6l1.4-4.3l.5-.1c1.2-.3 2.3-.8 3.4-1.4l.5-.3l4 2l1.8-1.8l-2-4l.3-.5c.6-1 1.1-2.2 1.4-3.4l.1-.5l4.3-1.4v-2.6l-4.3-1.4l-.1-.5c-.3-1.2-.8-2.3-1.4-3.4l-.3-.5l2-4l-1.8-1.8l-4 2l-.5-.3c-1.1-.6-2.2-1.1-3.4-1.4l-.5-.1L26.3 8h-2.6l-1.4 4.3l-.5.1c-1.2.3-2.3.8-3.4 1.4l-.5.3l-4-2l-1.8 1.8l2 4l-.3.5c-.6 1-1.1 2.2-1.4 3.4l-.1.5L8 23.7v2.6l4.3 1.4l.1.5c.3 1.2.8 2.3 1.4 3.4l.3.5l-2 4l1.8 1.8l4-2l.5.3c1.1.6 2.2 1.1 3.4 1.4l.5.1l1.4 4.3z"/>`
	heartInnerSVG           = `<path fill="currentColor" d="m25 39.7l-.6-.5C11.5 28.7 8 25 8 19c0-5 4-9 9-9c4.1 0 6.4 2.3 8 4.1c1.6-1.8 3.9-4.1 8-4.1c5 0 9 4 9 9c0 6-3.5 9.7-16.4 20.2l-.6.5zM17 12c-3.9 0-7 3.1-7 7c0 5.1 3.2 8.5 15 18.1c11.8-9.6 15-13 15-18.1c0-3.9-3.1-7-7-7c-3.5 0-5.4 2.1-6.9 3.8L25 17.1l-1.1-1.3C22.4 14.1 20.5 12 17 12z"/>`
	imageInnerSVG           = `<path fill="currentColor" d="M39 38H11c-1.7 0-3-1.3-3-3V15c0-1.7 1.3-3 3-3h28c1.7 0 3 1.3 3 3v20c0 1.7-1.3 3-3 3zM11 14c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h28c.6 0 1-.4 1-1V15c0-.6-.4-1-1-1H11z"/><path fill="currentColor" d="M30 24c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4zm0-6c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm5.3 19.7L19 22.4L9.7 31l-1.4-1.4l10.7-10l17.7 16.7z"/><path fill="currentColor" d="M40.4 32.7L35 28.3L30.5 32l-1.3-1.6l5.8-4.7l6.6 5.4z"/>`
	likeInnerSVG            = `<path fill="currentColor" d="M40 23.2c0-2.1-1.7-3.2-4-3.2h-6.7c.5-1.8.7-3.5.7-5c0-5.8-1.6-7-3-7c-.9 0-1.6.1-2.5.6c-.3.2-.4.4-.5.7l-1 5.4c-1.1 2.8-3.8 5.3-6 7V36c.8 0 1.6.4 2.6.9c1.1.5 2.2 1.1 3.4 1.1h9.5c2 0 3.5-1.6 3.5-3c0-.3 0-.5-.1-.7c1.2-.5 2.1-1.5 2.1-2.8c0-.6-.1-1.1-.3-1.6c.8-.5 1.5-1.4 1.5-2.4c0-.6-.3-1.2-.6-1.7c.8-.6 1.4-1.6 1.4-2.6zm-2.1 0c0 1.3-1.3 1.4-1.5 2c-.2.7.8.9.8 2.1c0 1.2-1.5 1.2-1.7 1.9c-.2.8.5 1 .5 2.2v.2c-.2 1-1.7 1.1-2 1.5c-.3.5 0 .7 0 1.8c0 .6-.7 1-1.5 1H23c-.8 0-1.6-.4-2.6-.9c-.8-.4-1.6-.8-2.4-1V23.5c2.5-1.9 5.7-4.7 6.9-8.2v-.2l.9-5c.4-.1.7-.1 1.2-.1c.2 0 1 1.2 1 5c0 1.5-.3 3.1-.8 5H27c-.6 0-1 .4-1 1s.4 1 1 1h9c1 0 1.9.5 1.9 1.2z"/><path fill="currentColor" d="M16 38h-6c-1.1 0-2-.9-2-2V22c0-1.1.9-2 2-2h6c1.1 0 2 .9 2 2v14c0 1.1-.9 2-2 2zm-6-16v14h6V22h-6z"/>`
	linkInnerSVG            = `<path fill="currentColor" d="M24 30.2c0 .2.1.5.1.8c0 1.4-.5 2.6-1.5 3.6l-2 2c-1 1-2.2 1.5-3.6 1.5c-2.8 0-5.1-2.3-5.1-5.1c0-1.4.5-2.6 1.5-3.6l2-2c1-1 2.2-1.5 3.6-1.5c.3 0 .5 0 .8.1l1.5-1.5c-.7-.3-1.5-.4-2.3-.4c-1.9 0-3.6.7-4.9 2l-2 2c-1.3 1.3-2 3-2 4.9c0 3.8 3.1 6.9 6.9 6.9c1.9 0 3.6-.7 4.9-2l2-2c1.3-1.3 2-3 2-4.9c0-.8-.1-1.6-.4-2.3L24 30.2zm9-20.1c-1.9 0-3.6.7-4.9 2l-2 2c-1.3 1.3-2 3-2 4.9c0 .8.1 1.6.4 2.3l1.5-1.5c0-.2-.1-.5-.1-.8c0-1.4.5-2.6 1.5-3.6l2-2c1-1 2.2-1.5 3.6-1.5c2.8 0 5.1 2.3 5.1 5.1c0 1.4-.5 2.6-1.5 3.6l-2 2c-1 1-2.2 1.5-3.6 1.5c-.3 0-.5 0-.8-.1l-1.5 1.5c.7.3 1.5.4 2.3.4c1.9 0 3.6-.7 4.9-2l2-2c1.3-1.3 2-3 2-4.9c0-3.8-3.1-6.9-6.9-6.9z"/><path fill="currentColor" d="M20 31c-.3 0-.5-.1-.7-.3c-.4-.4-.4-1 0-1.4l10-10c.4-.4 1-.4 1.4 0s.4 1 0 1.4l-10 10c-.2.2-.4.3-.7.3z"/>`
	locationInnerSVG        = `<path fill="currentColor" d="m25 42.5l-.8-.9C23.7 41.1 12 27.3 12 19c0-7.2 5.8-13 13-13s13 5.8 13 13c0 8.3-11.7 22.1-12.2 22.7l-.8.8zM25 8c-6.1 0-11 4.9-11 11c0 6.4 8.4 17.2 11 20.4c2.6-3.2 11-14 11-20.4c0-6.1-4.9-11-11-11z"/><path fill="currentColor" d="M25 24c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5zm0-8c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3z"/>`
	lockInnerSVG            = `<path fill="currentColor" d="M34 23h-2v-4c0-3.9-3.1-7-7-7s-7 3.1-7 7v4h-2v-4c0-5 4-9 9-9s9 4 9 9v4z"/><path fill="currentColor" d="M33 40H17c-1.7 0-3-1.3-3-3V25c0-1.7 1.3-3 3-3h16c1.7 0 3 1.3 3 3v12c0 1.7-1.3 3-3 3zM17 24c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1h16c.6 0 1-.4 1-1V25c0-.6-.4-1-1-1H17z"/><circle cx="25" cy="28" r="2" fill="currentColor"/><path fill="currentColor" d="M25.5 28h-1l-1 6h3z"/>`
	minusInnerSVG           = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="M16 24h18v2H16z"/>`
	naviconInnerSVG         = `<path fill="currentColor" d="M10 12h30v4H10zm0 10h30v4H10zm0 10h30v4H10z"/>`
	paperclipInnerSVG       = `<path fill="currentColor" d="M13.8 39.6c-1.5 0-3.1-.6-4.2-1.8c-2.3-2.3-2.3-6.1 0-8.5l17-17c3.1-3.1 8.2-3.1 11.3 0c3.1 3.1 3.1 8.2 0 11.3L25.1 36.4L23.7 35l12.7-12.7c2.3-2.3 2.3-6.1 0-8.5c-2.3-2.3-6.1-2.3-8.5 0l-17 17c-.8.8-1.2 1.8-1.2 2.8c0 1.1.4 2.1 1.2 2.8c1.6 1.6 4.1 1.6 5.7 0l12.7-12.7c.8-.8.8-2 0-2.8c-.8-.8-2-.8-2.8 0L18 29.3l-1.4-1.4l8.5-8.5c1.6-1.6 4.1-1.6 5.7 0c1.6 1.6 1.6 4.1 0 5.7L18 37.8c-1.1 1.2-2.7 1.8-4.2 1.8z"/>`
	pencilInnerSVG          = `<path fill="currentColor" d="m9.6 40.4l2.5-9.9L27 15.6l7.4 7.4l-14.9 14.9l-9.9 2.5zm4.3-8.9l-1.5 6.1l6.1-1.5L31.6 23L27 18.4L13.9 31.5z"/><path fill="currentColor" d="M17.8 37.3c-.6-2.5-2.6-4.5-5.1-5.1l.5-1.9c3.2.8 5.7 3.3 6.5 6.5l-1.9.5z"/><path fill="currentColor" d="m29.298 19.287l1.414 1.414l-13.01 13.02l-1.414-1.41zM11 39l2.9-.7c-.3-1.1-1.1-1.9-2.2-2.2L11 39zm24-16.6L27.6 15l3-3l.5.1c3.6.5 6.4 3.3 6.9 6.9l.1.5l-3.1 2.9zM30.4 15l4.6 4.6l.9-.9c-.5-2.3-2.3-4.1-4.6-4.6l-.9.9z"/>`
	playInnerSVG            = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="M20 33.7V16.3L35 25l-15 8.7zm2-14v10.5l9-5.3l-9-5.2z"/>`
	plusInnerSVG            = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="M16 24h18v2H16z"/><path fill="currentColor" d="M24 16h2v18h-2z"/>`
	pointerInnerSVG         = `<path fill="currentColor" d="M33 38H21c-.6 0-1-.4-1-1c0-1.5-.7-2.4-1.8-3.8c-.6-.7-1.3-1.6-2-2.7c-1.9-3-3.6-6.6-4-7.9c-.4-1.3-.1-2.2.3-2.7c.4-.6 1.2-.9 2.1-.9c1.2 0 2.4 1 3.5 2.3V11c0-1.7 1.3-3 3-3s3 1.3 3 3v4.2c.3-.1.6-.2 1-.2c1.1 0 2 .6 2.5 1.4c.4-.3.9-.4 1.4-.4c1.4 0 2.5.9 2.9 2.2c.3-.1.7-.2 1.1-.2c1.7 0 3 1.3 3 3v3c0 2.6-.5 4.7-1 6.7s-1 3.9-1 6.3c0 .6-.4 1-1 1zm-11.1-2H32c.1-2.2.6-4 1-5.8c.5-2 1-3.9 1-6.2v-3c0-.6-.4-1-1-1s-1 .4-1 1v1c0 .6-.4 1-1 1s-1-.4-1-1v-3c0-.6-.4-1-1-1s-1 .4-1 1v2c0 .6-.4 1-1 1s-1-.4-1-1v-3c0-.6-.4-1-1-1s-1 .4-1 1v2c0 .6-.4 1-1 1s-1-.4-1-1v-9c0-.6-.4-1-1-1s-1 .4-1 1v15c0 .6-.4 1-1 1s-1-.4-1-1v-.8c-.9-2.3-2.8-4.2-3.5-4.2c-.2 0-.4 0-.5.1c-.1.1-.1.4 0 .9c.3 1.1 1.8 4.3 3.8 7.5c.6 1 1.2 1.7 1.8 2.5c1.1 1.2 2.1 2.3 2.3 4z"/>`
	questionInnerSVG        = `<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="M19.8 19.6c.3-.8.6-1.4 1.2-1.9c.5-.5 1.1-.9 1.9-1.2s1.6-.4 2.5-.4c.7 0 1.4.1 2 .3c.6.2 1.2.5 1.7.9s.9.9 1.1 1.5c.3.6.4 1.3.4 2c0 1-.2 1.8-.6 2.5s-1 1.3-1.6 2l-1.3 1.3c-.3.3-.6.6-.7.9c-.2.3-.3.7-.3 1.1c-.1.4-.1.7-.1 1.5h-1.6c0-.8 0-1.1.1-1.7c.1-.5.3-1 .5-1.5c.2-.4.5-.8.9-1.2c.4-.4.9-.8 1.4-1.4c.5-.5.9-1 1.2-1.5s.5-1.2.5-1.8c0-.5-.1-1-.3-1.4c-.2-.4-.5-.8-.8-1.1c-.3-.3-.7-.5-1.2-.7c-.5-.2-.9-.3-1.4-.3c-.7 0-1.3.1-1.8.4c-.5.2-1 .6-1.3 1c-.3.4-.6.9-.8 1.5s-.4.9-.4 1.6h-1.6c0-.9.1-1.6.4-2.4zM26 32v2h-2v-2h2z"/>`
	redoInnerSVG            = `<path fill="currentColor" d="M25 38c-7.2 0-13-5.8-13-13s5.8-13 13-13c5.4 0 10.1 3.4 11.9 8.7l-1.9.7c-1.5-4.6-5.4-7.4-10-7.4c-6.1 0-11 4.9-11 11s4.9 11 11 11c4.3 0 8.2-2.5 10-6.4l1.8.8C34.7 35 30.1 38 25 38z"/><path fill="currentColor" d="M38 22h-8v-2h6v-6h2z"/>`
	refreshInnerSVG         = `<path fill="currentColor" d="M25 38c-7.2 0-13-5.8-13-13c0-3.2 1.2-6.2 3.3-8.6l1.5 1.3C15 19.7 14 22.3 14 25c0 6.1 4.9 11 11 11c1.6 0 3.1-.3 4.6-1l.8 1.8c-1.7.8-3.5 1.2-5.4 1.2zm9.7-4.3l-1.5-1.3c1.8-2 2.8-4.6 2.8-7.3c0-6.1-4.9-11-11-11c-1.6 0-3.1.3-4.6 1l-.8-1.8c1.7-.8 3.5-1.2 5.4-1.2c7.2 0 13 5.8 13 13c0 3.1-1.2 6.2-3.3 8.6z"/><path fill="currentColor" d="M18 24h-2v-6h-6v-2h8zm22 10h-8v-8h2v6h6z"/>`
	retweetInnerSVG         = `<path fill="currentColor" d="M38 35h-2V17c0-.6-.4-1-1-1H18v-2h17c1.7 0 3 1.3 3 3v18z"/><path fill="currentColor" d="m37 36.5l-6.8-7.8l1.6-1.4l5.2 6.2l5.2-6.2l1.6 1.4zm-5-.5H15c-1.7 0-3-1.3-3-3V15h2v18c0 .6.4 1 1 1h17v2z"/><path fill="currentColor" d="M18.2 22.7L13 16.5l-5.2 6.2l-1.6-1.4l6.8-7.8l6.8 7.8z"/>`
	scFacebookInnerSVG      = `<path fill="currentColor" d="M26 20v-3c0-1.3.3-2 2.4-2H31v-5h-4c-5 0-7 3.3-7 7v3h-4v5h4v15h6V25h4.4l.6-5h-5z"/>`
	scGithubInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M25 10c-8.3 0-15 6.7-15 15c0 6.6 4.3 12.2 10.3 14.2c.8.1 1-.3 1-.7v-2.6c-4.2.9-5.1-2-5.1-2c-.7-1.7-1.7-2.2-1.7-2.2c-1.4-.9.1-.9.1-.9c1.5.1 2.3 1.5 2.3 1.5c1.3 2.3 3.5 1.6 4.4 1.2c.1-1 .5-1.6 1-2c-3.3-.4-6.8-1.7-6.8-7.4c0-1.6.6-3 1.5-4c-.2-.4-.7-1.9.1-4c0 0 1.3-.4 4.1 1.5c1.2-.3 2.5-.5 3.8-.5c1.3 0 2.6.2 3.8.5c2.9-1.9 4.1-1.5 4.1-1.5c.8 2.1.3 3.6.1 4c1 1 1.5 2.4 1.5 4c0 5.8-3.5 7-6.8 7.4c.5.5 1 1.4 1 2.8v4.1c0 .4.3.9 1 .7c6-2 10.2-7.6 10.2-14.2C40 16.7 33.3 10 25 10z" clip-rule="evenodd"/>`
	scGooglePlusInnerSVG    = `<path fill="currentColor" fill-rule="evenodd" d="M18 23v4.8h7.9c-.3 2.1-2.4 6-7.9 6c-4.8 0-8.7-4-8.7-8.8s3.9-8.8 8.7-8.8c2.7 0 4.5 1.2 5.6 2.2l3.8-3.7C24.9 12.4 21.8 11 18 11c-7.7 0-14 6.3-14 14s6.3 14 14 14c8.1 0 13.4-5.7 13.4-13.7c0-.9-.1-1.6-.2-2.3H18zm30 0h-4v-4h-4v4h-4v4h4v4h4v-4h4z" clip-rule="evenodd"/>`
	scInstagramInnerSVG     = `<g fill="currentColor" fill-rule="evenodd"><path d="M25 12c-3.53 0-3.973.015-5.36.078c-1.384.063-2.329.283-3.156.604a6.372 6.372 0 0 0-2.302 1.5a6.372 6.372 0 0 0-1.5 2.303c-.321.826-.54 1.771-.604 3.155C12.015 21.027 12 21.47 12 25c0 3.53.015 3.973.078 5.36c.063 1.384.283 2.329.604 3.155c.333.855.777 1.58 1.5 2.303a6.372 6.372 0 0 0 2.302 1.5c.827.32 1.772.54 3.156.604c1.387.063 1.83.078 5.36.078c3.53 0 3.973-.015 5.36-.078c1.384-.063 2.329-.283 3.155-.604a6.371 6.371 0 0 0 2.303-1.5a6.372 6.372 0 0 0 1.5-2.303c.32-.826.54-1.771.604-3.155c.063-1.387.078-1.83.078-5.36c0-3.53-.015-3.973-.078-5.36c-.063-1.384-.283-2.329-.605-3.155a6.372 6.372 0 0 0-1.499-2.303a6.371 6.371 0 0 0-2.303-1.5c-.826-.32-1.771-.54-3.155-.604C28.973 12.015 28.53 12 25 12m0 2.342c3.471 0 3.882.014 5.253.076c1.267.058 1.956.27 2.414.448c.607.236 1.04.517 1.495.972c.455.455.736.888.972 1.495c.178.458.39 1.146.448 2.414c.062 1.37.076 1.782.076 5.253s-.014 3.882-.076 5.253c-.058 1.268-.27 1.956-.448 2.414a4.028 4.028 0 0 1-.972 1.495a4.027 4.027 0 0 1-1.495.972c-.458.178-1.147.39-2.414.448c-1.37.062-1.782.076-5.253.076s-3.883-.014-5.253-.076c-1.268-.058-1.956-.27-2.414-.448a4.027 4.027 0 0 1-1.495-.972a4.03 4.03 0 0 1-.972-1.495c-.178-.458-.39-1.146-.448-2.414c-.062-1.37-.076-1.782-.076-5.253s.014-3.882.076-5.253c.058-1.268.27-1.956.448-2.414c.236-.607.517-1.04.972-1.495a4.028 4.028 0 0 1 1.495-.972c.458-.178 1.146-.39 2.414-.448c1.37-.062 1.782-.076 5.253-.076"/><path d="M25 18a7 7 0 1 0 0 14a7 7 0 0 0 0-14m0 11.5a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9m8.7-11.4a1.6 1.6 0 1 1-3.2 0a1.6 1.6 0 0 1 3.2 0"/></g>`
	scLinkedinInnerSVG      = `<path fill="currentColor" d="M36.1 12H13.9c-1.1 0-1.9.8-1.9 1.9v22.2c0 1 .9 1.9 1.9 1.9h22.2c1.1 0 1.9-.8 1.9-1.9V13.9c0-1.1-.9-1.9-1.9-1.9zM20 34h-4V22h4v12zm-2-13.6c-1.3 0-2.4-1.1-2.4-2.4c0-1.3 1.1-2.4 2.4-2.4c1.3 0 2.4 1.1 2.4 2.4c0 1.3-1.1 2.4-2.4 2.4zM34 34h-4v-6c0-1.6-.4-3.2-2-3.2s-2 1.6-2 3.2v6h-4V22h4v1.4h.2c.5-1 1.8-1.8 3.3-1.8c3.7 0 4.5 2.4 4.5 5.4v7z"/>`
	scOdnoklassnikiInnerSVG = `<path fill="currentColor" d="M25 26c-4.4 0-8-3.6-8-8s3.6-8 8-8s8 3.6 8 8s-3.6 8-8 8zm0-12.2c-2.3 0-4.2 1.9-4.2 4.2s1.9 4.2 4.2 4.2s4.2-1.9 4.2-4.2s-1.9-4.2-4.2-4.2z"/><path fill="currentColor" d="M33.6 26.8c-.7-.9-1.9-1-2.8-.4c0 0-2.2 1.6-5.8 1.6c-3.6 0-5.8-1.6-5.8-1.6c-.9-.7-2.1-.5-2.8.4c-.7.9-.5 2.1.4 2.8c.1.1 2.2 1.7 5.7 2.2l-5.3 5.4c-.8.8-.8 2.1 0 2.8c.4.4.9.6 1.4.6c.5 0 1-.2 1.4-.6l5-5.1l5 5.1c.4.4.9.6 1.4.6c.5 0 1-.2 1.4-.6c.8-.8.8-2 0-2.8l-5.3-5.4c3.5-.6 5.6-2.2 5.7-2.2c.9-.7 1.1-2 .4-2.8z"/>`
	scPinterestInnerSVG     = `<path fill="currentColor" d="M25 10c-8.3 0-15 6.7-15 15c0 6.4 4 11.8 9.5 14c-.1-1.2-.2-3 .1-4.3c.3-1.2 1.8-7.5 1.8-7.5s-.4-.9-.4-2.2c0-2.1 1.2-3.6 2.7-3.6c1.3 0 1.9 1 1.9 2.1c0 1.3-.8 3.2-1.2 5c-.4 1.5.7 2.7 2.2 2.7c2.7 0 4.7-2.8 4.7-6.9c0-3.6-2.6-6.1-6.3-6.1c-4.3 0-6.8 3.2-6.8 6.5c0 1.3.5 2.7 1.1 3.4c.1.1.1.3.1.4c-.1.5-.4 1.5-.4 1.7c-.1.3-.2.3-.5.2c-1.9-.9-3-3.6-3-5.8c0-4.7 3.4-9.1 9.9-9.1c5.2 0 9.2 3.7 9.2 8.7c0 5.2-3.3 9.3-7.8 9.3c-1.5 0-2.9-.8-3.4-1.7c0 0-.8 2.9-.9 3.6c-.3 1.3-1.3 2.9-1.9 3.9c1.4.5 2.9.7 4.4.7c8.3 0 15-6.7 15-15s-6.7-15-15-15z"/>`
	scSkypeInnerSVG         = `<path fill="currentColor" d="M38 27.3c.1-.8.2-1.6.2-2.4c0-1.8-.3-3.5-1-5.1c-.7-1.6-1.6-3-2.8-4.2c-1.2-1.2-2.6-2.2-4.2-2.8c-1.6-.7-3.4-1-5.1-1c-.8 0-1.7.1-2.5.2c-1.1-.6-2.4-.9-3.7-.9c-2.1 0-4.1.8-5.5 2.3c-1.5 1.5-2.3 3.4-2.3 5.5c0 1.3.3 2.6 1 3.8c-.1.7-.2 1.5-.2 2.3c0 1.8.3 3.5 1 5.1c.7 1.6 1.6 3 2.8 4.2c1.2 1.2 2.6 2.2 4.2 2.8c1.6.7 3.4 1 5.1 1c.8 0 1.6-.1 2.3-.2c1.2.7 2.5 1 3.9 1c2.1 0 4.1-.8 5.5-2.3c1.5-1.5 2.3-3.4 2.3-5.5c0-1.3-.3-2.6-1-3.8zM25.1 33c-4.7 0-6.8-2.3-6.8-4c0-.9.7-1.5 1.6-1.5c2 0 1.5 2.9 5.2 2.9c1.9 0 3-1 3-2.1c0-.6-.3-1.4-1.6-1.7l-4.2-1c-3.4-.8-4-2.7-4-4.4c0-3.6 3.3-4.9 6.5-4.9c2.9 0 6.3 1.6 6.3 3.7c0 .9-.8 1.4-1.7 1.4c-1.7 0-1.4-2.4-4.9-2.4c-1.7 0-2.7.8-2.7 1.9c0 1.1 1.4 1.5 2.5 1.7l3.1.7c3.4.8 4.2 2.7 4.2 4.6c.1 2.9-2.1 5.1-6.5 5.1z"/>`
	scSoundcloudInnerSVG    = `<path fill="currentColor" d="M40 24h-.2c-.9-4.6-5-8-9.8-8c-3.1 0-5.9 1.4-7.7 3.7c-.2.3-.3.6-.3 1.2l-.4 9.1l.4 5.5c0 .3.3.5.5.5H40c3.3 0 6-2.7 6-6s-2.7-6-6-6zm-21.1-4c-.3 0-.5.2-.5.5l-.8 9v1l.8 5c0 .3.3.5.6.5h.2c.3 0 .5-.2.6-.5l.8-5c0-.3.1-.7 0-1l-.8-9c0-.3-.3-.5-.5-.5h-.4zm-4 1c-.3 0-.5.2-.5.5l-.8 8v1l.8 5c0 .3.3.5.6.5h.2c.3 0 .5-.2.6-.5l.8-5c0-.3.1-.7 0-1l-.8-8c0-.3-.3-.5-.5-.5h-.4zM11 24c-.3 0-.5.2-.6.5l-.8 5v1l.8 5c0 .3.3.5.6.5s.5-.2.6-.5l.8-5v-1l-.8-5c-.1-.3-.3-.5-.6-.5zm-4-1c-.3 0-.5.2-.6.5l-.9 6v1l.8 5c.2.3.4.5.7.5c.3 0 .5-.2.6-.5l.8-5c0-.3.1-.7 0-1l-.9-6c0-.3-.2-.5-.5-.5zm-3.7 3c-.3 0-.5.2-.6.5l-.6 3c-.1.3-.1.7 0 1l.6 4c.1.3.3.5.6.5s.5-.2.6-.5l.6-4v-1l-.6-3c-.1-.3-.3-.5-.6-.5z"/>`
	scTelegramInnerSVG      = `<path fill="currentColor" d="M37.1 13L9.4 24c-.9.3-.8 1.6.1 1.9l7 2.2l2.8 8.8c.2.7 1.1.9 1.6.4l4.1-3.8l7.8 5.7c.6.4 1.4.1 1.6-.6l5.4-23.2c.3-1.7-1.2-3-2.7-2.4zM20.9 29.8L20 35l-2-7.2L37.5 15L20.9 29.8z"/>`
	scTumblrInnerSVG        = `<path fill="currentColor" d="M30.9 32.4c-.5.2-1.5.5-2.3.5c-2.2.1-2.7-1.6-2.7-2.8v-8.7h5.6v-4.2H26V10h-4.1c-.1 0-.2.1-.2.2c-.2 2.2-1.3 6-5.5 7.5v3.6H19v9.1c0 3.1 2.3 7.6 8.4 7.5c2.1 0 4.3-.9 4.8-1.6l-1.3-3.9z"/>`
	scTwitterInnerSVG       = `<path fill="currentColor" d="M39.2 16.8c-1.1.5-2.2.8-3.5 1c1.2-.8 2.2-1.9 2.7-3.3c-1.2.7-2.5 1.2-3.8 1.5c-1.1-1.2-2.7-1.9-4.4-1.9c-3.3 0-6.1 2.7-6.1 6.1c0 .5.1.9.2 1.4c-5-.2-9.5-2.7-12.5-6.3c-.5.7-.8 1.7-.8 2.8c0 2.1 1.1 4 2.7 5c-1 0-1.9-.3-2.7-.8v.1c0 2.9 2.1 5.4 4.9 5.9c-.5.1-1 .2-1.6.2c-.4 0-.8 0-1.1-.1c.8 2.4 3 4.2 5.7 4.2c-2.1 1.6-4.7 2.6-7.5 2.6c-.5 0-1 0-1.4-.1c2.4 1.9 5.6 2.9 9 2.9c11.1 0 17.2-9.2 17.2-17.2V20c1.2-.9 2.2-1.9 3-3.2z"/>`
	scVimeoInnerSVG         = `<path fill="currentColor" d="M38 19.6c-.1 2.7-2 6.4-5.6 11.1c-3.8 4.9-7 7.4-9.6 7.4c-1.6 0-3-1.5-4.1-4.5c-.7-2.7-1.5-5.5-2.2-8.2c-.8-3-1.7-4.5-2.7-4.5c-.2 0-.9.4-2.2 1.3l-1.3-1.7c1.4-1.2 2.7-2.4 4-3.6c1.8-1.6 3.2-2.4 4.1-2.5c2.2-.2 3.5 1.3 4 4.4c.5 3.4.9 5.5 1.1 6.4c.6 2.8 1.3 4.2 2.1 4.2c.6 0 1.5-.9 2.6-2.8c1.2-1.8 1.8-3.2 1.9-4.2c.2-1.6-.5-2.4-1.9-2.4c-.7 0-1.3.2-2 .5c1.4-4.5 4-6.6 7.8-6.5c2.8.1 4.2 1.9 4 5.6z"/>`
	scVkInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M25.1 35.9h2s.6-.1.9-.4c.3-.3.3-.9.3-.9s0-2.6 1.2-3c1.2-.4 2.8 2.6 4.4 3.7c1.2.9 2.1.7 2.1.7l4.4-.1s2.3-.1 1.2-2c-.1-.1-.6-1.3-3.3-3.8c-2.8-2.6-2.4-2.1.9-6.6c2-2.7 2.8-4.3 2.6-5.1c-.2-.7-1.7-.5-1.7-.5h-5s-.4-.1-.6.1c-.3.2-.4.5-.4.5s-.8 2.1-1.8 3.9c-2.2 3.7-3.1 3.9-3.4 3.7c-.8-.5-.6-2.2-.6-3.3c0-3.6.6-5.1-1.1-5.5c-.5-.1-.9-.2-2.3-.2c-1.8 0-3.3 0-4.1.4c-.6.3-1 .9-.7.9c.3 0 1.1.2 1.5.7c.4.9.4 2.4.4 2.4s.3 4.3-.7 4.8c-.7.4-1.6-.4-3.6-3.8c-1-1.7-1.8-3.7-1.8-3.7s-.1-.4-.4-.6c-.3-.2-.8-.3-.8-.3H10s-.7 0-1 .3c-.2.3 0 .8 0 .8s3.7 8.6 7.9 13c3.9 4.2 8.2 3.9 8.2 3.9z" clip-rule="evenodd"/>`
	scYoutubeInnerSVG       = `<path fill="currentColor" d="M39.7 18.6s-.3-2.1-1.2-3c-1.1-1.2-2.4-1.2-3-1.3C31.3 14 25 14 25 14s-6.3 0-10.5.3c-.6.1-1.9.1-3 1.3c-.9.9-1.2 3-1.2 3S10 21 10 23.4v2.2c0 2.4.3 4.9.3 4.9s.3 2.1 1.2 3c1.1 1.2 2.6 1.2 3.3 1.3c2.4.1 10.2.2 10.2.2s6.3 0 10.5-.3c.6-.1 1.9-.1 3-1.3c.9-.9 1.2-3 1.2-3s.3-2.4.3-4.8v-2.2c0-2.4-.3-4.8-.3-4.8zm-17.8 9.8V20l8.1 4.2l-8.1 4.2z"/>`
	searchInnerSVG          = `<path fill="currentColor" d="M23 36c-7.2 0-13-5.8-13-13s5.8-13 13-13s13 5.8 13 13s-5.8 13-13 13zm0-24c-6.1 0-11 4.9-11 11s4.9 11 11 11s11-4.9 11-11s-4.9-11-11-11z"/><path fill="currentColor" d="m32.682 31.267l8.98 8.98l-1.414 1.414l-8.98-8.98z"/>`
	shareAppleInnerSVG      = `<path fill="currentColor" d="M30.3 13.7L25 8.4l-5.3 5.3l-1.4-1.4L25 5.6l6.7 6.7z"/><path fill="currentColor" d="M24 7h2v21h-2z"/><path fill="currentColor" d="M35 40H15c-1.7 0-3-1.3-3-3V19c0-1.7 1.3-3 3-3h7v2h-7c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h20c.6 0 1-.4 1-1V19c0-.6-.4-1-1-1h-7v-2h7c1.7 0 3 1.3 3 3v18c0 1.7-1.3 3-3 3z"/>`
	shareGoogleInnerSVG     = `<path fill="currentColor" d="M15 30c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5zm0-8c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3zm20-2c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5zm0-8c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3zm0 28c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5zm0-8c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3z"/><path fill="currentColor" d="m19.007 25.885l12.88 6.44l-.895 1.788l-12.88-6.44zm11.986-10l.894 1.79l-12.88 6.438l-.894-1.79z"/>`
	spinnerInnerSVG         = `<path fill="currentColor" d="M25 18c-.6 0-1-.4-1-1V9c0-.6.4-1 1-1s1 .4 1 1v8c0 .6-.4 1-1 1z"/><path fill="currentColor" d="M25 42c-.6 0-1-.4-1-1v-8c0-.6.4-1 1-1s1 .4 1 1v8c0 .6-.4 1-1 1zm4-23c-.2 0-.3 0-.5-.1c-.4-.3-.6-.8-.3-1.3l4-6.9c.3-.4.8-.6 1.3-.3c.4.3.6.8.3 1.3l-4 6.9c-.2.2-.5.4-.8.4zM17 39.8c-.2 0-.3 0-.5-.1c-.4-.3-.6-.8-.3-1.3l4-6.9c.3-.4.8-.6 1.3-.3c.4.3.6.8.3 1.3l-4 6.9c-.2.2-.5.4-.8.4z" opacity=".3"/><path fill="currentColor" d="M21 19c-.3 0-.6-.2-.8-.5l-4-6.9c-.3-.4-.1-1 .3-1.3c.4-.3 1-.1 1.3.3l4 6.9c.3.4.1 1-.3 1.3c-.2.2-.3.2-.5.2z" opacity=".93"/><path fill="currentColor" d="M33 39.8c-.3 0-.6-.2-.8-.5l-4-6.9c-.3-.4-.1-1 .3-1.3c.4-.3 1-.1 1.3.3l4 6.9c.3.4.1 1-.3 1.3c-.2.1-.3.2-.5.2z" opacity=".3"/><path fill="currentColor" d="M17 26H9c-.6 0-1-.4-1-1s.4-1 1-1h8c.6 0 1 .4 1 1s-.4 1-1 1z" opacity=".65"/><path fill="currentColor" d="M41 26h-8c-.6 0-1-.4-1-1s.4-1 1-1h8c.6 0 1 .4 1 1s-.4 1-1 1z" opacity=".3"/><path fill="currentColor" d="M18.1 21.9c-.2 0-.3 0-.5-.1l-6.9-4c-.4-.3-.6-.8-.3-1.3c.3-.4.8-.6 1.3-.3l6.9 4c.4.3.6.8.3 1.3c-.2.3-.5.4-.8.4z" opacity=".86"/><path fill="currentColor" d="M38.9 33.9c-.2 0-.3 0-.5-.1l-6.9-4c-.4-.3-.6-.8-.3-1.3c.3-.4.8-.6 1.3-.3l6.9 4c.4.3.6.8.3 1.3c-.2.3-.5.4-.8.4z" opacity=".3"/><path fill="currentColor" d="M11.1 33.9c-.3 0-.6-.2-.8-.5c-.3-.4-.1-1 .3-1.3l6.9-4c.4-.3 1-.1 1.3.3c.3.4.1 1-.3 1.3l-6.9 4c-.1.2-.3.2-.5.2z" opacity=".44"/><path fill="currentColor" d="M31.9 21.9c-.3 0-.6-.2-.8-.5c-.3-.4-.1-1 .3-1.3l6.9-4c.4-.3 1-.1 1.3.3c.3.4.1 1-.3 1.3l-6.9 4c-.2.2-.3.2-.5.2z" opacity=".3"/>`
	spinnerThreeInnerSVG    = `<path fill="currentColor" d="M41.9 23.9c-.3-6.1-4-11.8-9.5-14.4c-6-2.7-13.3-1.6-18.3 2.6c-4.8 4-7 10.5-5.6 16.6c1.3 6 6 10.9 11.9 12.5c7.1 2 13.6-1.4 17.6-7.2c-3.6 4.8-9.1 8-15.2 6.9c-6.1-1.1-11.1-5.7-12.5-11.7c-1.5-6.4 1.5-13.1 7.2-16.4c5.9-3.4 14.2-2.1 18.1 3.7c1 1.4 1.7 3.1 2 4.8c.3 1.4.2 2.9.4 4.3c.2 1.3 1.3 3 2.8 2.1c1.3-.8 1.2-2.5 1.1-3.8c0-.4.1.7 0 0z"/>`
	spinnerTwoInnerSVG      = `<circle cx="25" cy="10" r="2" fill="currentColor"/><circle cx="25" cy="40" r="2" fill="currentColor" opacity=".3"/><circle cx="32.5" cy="12" r="2" fill="currentColor" opacity=".3"/><circle cx="17.5" cy="38" r="2" fill="currentColor" opacity=".3"/><circle cx="17.5" cy="12" r="2" fill="currentColor" opacity=".93"/><circle cx="32.5" cy="38" r="2" fill="currentColor" opacity=".3"/><circle cx="10" cy="25" r="2" fill="currentColor" opacity=".65"/><circle cx="40" cy="25" r="2" fill="currentColor" opacity=".3"/><circle cx="12" cy="17.5" r="2" fill="currentColor" opacity=".86"/><circle cx="38" cy="32.5" r="2" fill="currentColor" opacity=".3"/><circle cx="12" cy="32.5" r="2" fill="currentColor" opacity=".44"/><circle cx="38" cy="17.5" r="2" fill="currentColor" opacity=".3"/>`
	starInnerSVG            = `<path fill="currentColor" d="M15.2 40.6c-.2 0-.4-.1-.6-.2c-.4-.3-.5-.7-.4-1.1l3.9-12l-10.2-7.5c-.4-.3-.5-.7-.4-1.1s.5-.7 1-.7h12.7L25 5.9c.1-.4.5-.7 1-.7s.8.3 1 .7L30.9 18h12.7c.4 0 .8.2 1 .6s0 .9-.4 1.1L34 27.1l3.9 12c.1.4 0 .9-.4 1.1s-.8.3-1.2 0L26 33l-10.2 7.4c-.2.1-.4.2-.6.2zM26 30.7c.2 0 .4.1.6.2l8.3 6.1l-3.2-9.8c-.1-.4 0-.9.4-1.1l8.3-6.1H30.1c-.4 0-.8-.3-1-.7L26 9.5l-3.2 9.8c-.1.4-.5.7-1 .7H11.5l8.3 6.1c.4.3.5.7.4 1.1L17.1 37l8.3-6.1c.2-.1.4-.2.6-.2z"/>`
	tagInnerSVG             = `<path fill="currentColor" d="M22 40.1c-.9 0-1.7-.3-2.3-.9l-8.9-8.9c-1.2-1.2-1.2-3.3 0-4.5l11.9-11.9c1-1 3-1.8 4.5-1.8h7.6c1.8 0 3.2 1.4 3.2 3.2v7.6c0 1.5-.8 3.4-1.8 4.5L24.3 39.2c-.6.6-1.4.9-2.3.9zM27.2 14c-1 0-2.4.6-3 1.3L12.3 27.2c-.5.5-.5 1.2 0 1.7l8.9 8.9c.5.4 1.2.4 1.7 0l11.9-11.9c.7-.7 1.3-2.1 1.3-3v-7.6c0-.7-.5-1.2-1.2-1.2h-7.7z"/><path fill="currentColor" d="M30 24c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4zm0-6c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2z"/>`
	trashInnerSVG           = `<path fill="currentColor" d="M20 18h2v16h-2zm4 0h2v16h-2zm4 0h2v16h-2zm-16-6h26v2H12zm18 0h-2v-1c0-.6-.4-1-1-1h-4c-.6 0-1 .4-1 1v1h-2v-1c0-1.7 1.3-3 3-3h4c1.7 0 3 1.3 3 3v1z"/><path fill="currentColor" d="M31 40H19c-1.6 0-3-1.3-3.2-2.9l-1.8-24l2-.2l1.8 24c0 .6.6 1.1 1.2 1.1h12c.6 0 1.1-.5 1.2-1.1l1.8-24l2 .2l-1.8 24C34 38.7 32.6 40 31 40z"/>`
	trophyInnerSVG          = `<path fill="currentColor" d="M28.6 29.4c3-2.3 7.4-5.7 7.4-18.4v-1H14v1c0 12.7 4.5 16.1 7.4 18.4c1.7 1.3 2.6 2 2.6 3.6v3c-1.6.2-3.2.8-3.8 2H18c-1.1 0-2 .9-2 2h18c0-1.1-.9-2-2-2h-2.2c-.6-1.2-2.1-1.8-3.8-2v-3c0-1.6.8-2.3 2.6-3.6zm-3.6.5c-.6-.8-1.5-1.5-2.3-2.1c-2.7-2.1-6.4-4.9-6.6-15.8h18c-.2 10.8-3.9 13.7-6.6 15.8c-1 .7-1.9 1.3-2.5 2.1z"/><path fill="currentColor" d="M18.8 27C18.7 27 8 24.7 8 13v-1h7v2h-5c.6 9.2 9.1 11 9.2 11l-.4 2zm12.4 0l-.4-2c.4-.1 8.6-1.9 9.2-11h-5v-2h7v1c0 11.7-10.7 14-10.8 14z"/>`
	undoInnerSVG            = `<path fill="currentColor" d="M25 38c-5.1 0-9.7-3-11.8-7.6l1.8-.8c1.8 3.9 5.7 6.4 10 6.4c6.1 0 11-4.9 11-11s-4.9-11-11-11c-4.6 0-8.5 2.8-10.1 7.3l-1.9-.7c1.9-5.2 6.6-8.6 12-8.6c7.2 0 13 5.8 13 13s-5.8 13-13 13z"/><path fill="currentColor" d="M20 22h-8v-8h2v6h6z"/>`
	unlockInnerSVG          = `<path fill="currentColor" d="M18 23h-2v-4c0-5 4-9 9-9c4.5 0 8.4 3.4 8.9 7.9l-2 .2c-.4-3.5-3.4-6.1-6.9-6.1c-3.9 0-7 3.1-7 7v4z"/><path fill="currentColor" d="M33 40H17c-1.7 0-3-1.3-3-3V25c0-1.7 1.3-3 3-3h16c1.7 0 3 1.3 3 3v12c0 1.7-1.3 3-3 3zM17 24c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1h16c.6 0 1-.4 1-1V25c0-.6-.4-1-1-1H17z"/><circle cx="25" cy="28" r="2" fill="currentColor"/><path fill="currentColor" d="M25.5 28h-1l-1 6h3z"/>`
	userInnerSVG            = `<path fill="currentColor" d="M25.1 42c-9.4 0-17-7.6-17-17s7.6-17 17-17s17 7.6 17 17s-7.7 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.8-15-15-15z"/><path fill="currentColor" d="m15.3 37.3l-1.8-.8c.5-1.2 2.1-1.9 3.8-2.7c1.7-.8 3.8-1.7 3.8-2.8v-1.5c-.6-.5-1.6-1.6-1.8-3.2c-.5-.5-1.3-1.4-1.3-2.6c0-.7.3-1.3.5-1.7c-.2-.8-.4-2.3-.4-3.5c0-3.9 2.7-6.5 7-6.5c1.2 0 2.7.3 3.5 1.2c1.9.4 3.5 2.6 3.5 5.3c0 1.7-.3 3.1-.5 3.8c.2.3.4.8.4 1.4c0 1.3-.7 2.2-1.3 2.6c-.2 1.6-1.1 2.6-1.7 3.1V31c0 .9 1.8 1.6 3.4 2.2c1.9.7 3.9 1.5 4.6 3.1l-1.9.7c-.3-.8-1.9-1.4-3.4-1.9c-2.2-.8-4.7-1.7-4.7-4v-2.6l.5-.3s1.2-.8 1.2-2.4v-.7l.6-.3c.1 0 .6-.3.6-1.1c0-.2-.2-.5-.3-.6l-.4-.4l.2-.5s.5-1.6.5-3.6c0-1.9-1.1-3.3-2-3.3h-.6l-.3-.5c0-.4-.7-.8-1.9-.8c-3.1 0-5 1.7-5 4.5c0 1.3.5 3.5.5 3.5l.1.5l-.4.5c-.1 0-.3.3-.3.7c0 .5.6 1.1.9 1.3l.4.3v.5c0 1.5 1.3 2.3 1.3 2.4l.5.3v2.6c0 2.4-2.6 3.6-5 4.6c-1.1.4-2.6 1.1-2.8 1.6z"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Archive(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archiveInnerSVG).
		Element(children...)
}

func ArrowDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownInnerSVG).
		Element(children...)
}

func ArrowLeft(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftInnerSVG).
		Element(children...)
}

func ArrowRight(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightInnerSVG).
		Element(children...)
}

func ArrowUp(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpInnerSVG).
		Element(children...)
}

func Bell(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bellInnerSVG).
		Element(children...)
}

func Calendar(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarInnerSVG).
		Element(children...)
}

func Camera(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cameraInnerSVG).
		Element(children...)
}

func Cart(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartInnerSVG).
		Element(children...)
}

func Chart(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartInnerSVG).
		Element(children...)
}

func Check(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkInnerSVG).
		Element(children...)
}

func ChevronDown(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
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
			engine.NewAttribute("viewBox", "0 0 50 50"),
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
			engine.NewAttribute("viewBox", "0 0 50 50"),
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
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpInnerSVG).
		Element(children...)
}

func Clock(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
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
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeInnerSVG).
		Element(children...)
}

func CloseO(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closeOInnerSVG).
		Element(children...)
}

func Comment(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentInnerSVG).
		Element(children...)
}

func CreditCard(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardInnerSVG).
		Element(children...)
}

func Envelope(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(envelopeInnerSVG).
		Element(children...)
}

func Exclamation(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(exclamationInnerSVG).
		Element(children...)
}

func ExternalLink(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
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
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eyeInnerSVG).
		Element(children...)
}

func Gear(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gearInnerSVG).
		Element(children...)
}

func Heart(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartInnerSVG).
		Element(children...)
}

func Image(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageInnerSVG).
		Element(children...)
}

func Like(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(likeInnerSVG).
		Element(children...)
}

func Link(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkInnerSVG).
		Element(children...)
}

func Location(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationInnerSVG).
		Element(children...)
}

func Lock(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockInnerSVG).
		Element(children...)
}

func Minus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusInnerSVG).
		Element(children...)
}

func Navicon(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(naviconInnerSVG).
		Element(children...)
}

func Paperclip(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paperclipInnerSVG).
		Element(children...)
}

func Pencil(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pencilInnerSVG).
		Element(children...)
}

func Play(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playInnerSVG).
		Element(children...)
}

func Plus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusInnerSVG).
		Element(children...)
}

func Pointer(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pointerInnerSVG).
		Element(children...)
}

func Question(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(questionInnerSVG).
		Element(children...)
}

func Redo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
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
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(refreshInnerSVG).
		Element(children...)
}

func Retweet(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(retweetInnerSVG).
		Element(children...)
}

func ScFacebook(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scFacebookInnerSVG).
		Element(children...)
}

func ScGithub(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scGithubInnerSVG).
		Element(children...)
}

func ScGooglePlus(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scGooglePlusInnerSVG).
		Element(children...)
}

func ScInstagram(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scInstagramInnerSVG).
		Element(children...)
}

func ScLinkedin(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scLinkedinInnerSVG).
		Element(children...)
}

func ScOdnoklassniki(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scOdnoklassnikiInnerSVG).
		Element(children...)
}

func ScPinterest(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scPinterestInnerSVG).
		Element(children...)
}

func ScSkype(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scSkypeInnerSVG).
		Element(children...)
}

func ScSoundcloud(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scSoundcloudInnerSVG).
		Element(children...)
}

func ScTelegram(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scTelegramInnerSVG).
		Element(children...)
}

func ScTumblr(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scTumblrInnerSVG).
		Element(children...)
}

func ScTwitter(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scTwitterInnerSVG).
		Element(children...)
}

func ScVimeo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scVimeoInnerSVG).
		Element(children...)
}

func ScVk(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scVkInnerSVG).
		Element(children...)
}

func ScYoutube(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scYoutubeInnerSVG).
		Element(children...)
}

func Search(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchInnerSVG).
		Element(children...)
}

func ShareApple(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareAppleInnerSVG).
		Element(children...)
}

func ShareGoogle(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareGoogleInnerSVG).
		Element(children...)
}

func Spinner(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spinnerInnerSVG).
		Element(children...)
}

func SpinnerThree(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spinnerThreeInnerSVG).
		Element(children...)
}

func SpinnerTwo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spinnerTwoInnerSVG).
		Element(children...)
}

func Star(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starInnerSVG).
		Element(children...)
}

func Tag(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagInnerSVG).
		Element(children...)
}

func Trash(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashInnerSVG).
		Element(children...)
}

func Trophy(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trophyInnerSVG).
		Element(children...)
}

func Undo(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(undoInnerSVG).
		Element(children...)
}

func Unlock(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unlockInnerSVG).
		Element(children...)
}

func User(children ...engine.UberChild) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 50 50"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userInnerSVG).
		Element(children...)
}

func ByName(name string) (*engine.UberElement, error) {
	switch name {
	case "archive":
		return Archive(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "bell":
		return Bell(), nil
	case "calendar":
		return Calendar(), nil
	case "camera":
		return Camera(), nil
	case "cart":
		return Cart(), nil
	case "chart":
		return Chart(), nil
	case "check":
		return Check(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "clock":
		return Clock(), nil
	case "close":
		return Close(), nil
	case "close-o":
		return CloseO(), nil
	case "comment":
		return Comment(), nil
	case "credit-card":
		return CreditCard(), nil
	case "envelope":
		return Envelope(), nil
	case "exclamation":
		return Exclamation(), nil
	case "external-link":
		return ExternalLink(), nil
	case "eye":
		return Eye(), nil
	case "gear":
		return Gear(), nil
	case "heart":
		return Heart(), nil
	case "image":
		return Image(), nil
	case "like":
		return Like(), nil
	case "link":
		return Link(), nil
	case "location":
		return Location(), nil
	case "lock":
		return Lock(), nil
	case "minus":
		return Minus(), nil
	case "navicon":
		return Navicon(), nil
	case "paperclip":
		return Paperclip(), nil
	case "pencil":
		return Pencil(), nil
	case "play":
		return Play(), nil
	case "plus":
		return Plus(), nil
	case "pointer":
		return Pointer(), nil
	case "question":
		return Question(), nil
	case "redo":
		return Redo(), nil
	case "refresh":
		return Refresh(), nil
	case "retweet":
		return Retweet(), nil
	case "sc-facebook":
		return ScFacebook(), nil
	case "sc-github":
		return ScGithub(), nil
	case "sc-google-plus":
		return ScGooglePlus(), nil
	case "sc-instagram":
		return ScInstagram(), nil
	case "sc-linkedin":
		return ScLinkedin(), nil
	case "sc-odnoklassniki":
		return ScOdnoklassniki(), nil
	case "sc-pinterest":
		return ScPinterest(), nil
	case "sc-skype":
		return ScSkype(), nil
	case "sc-soundcloud":
		return ScSoundcloud(), nil
	case "sc-telegram":
		return ScTelegram(), nil
	case "sc-tumblr":
		return ScTumblr(), nil
	case "sc-twitter":
		return ScTwitter(), nil
	case "sc-vimeo":
		return ScVimeo(), nil
	case "sc-vk":
		return ScVk(), nil
	case "sc-youtube":
		return ScYoutube(), nil
	case "search":
		return Search(), nil
	case "share-apple":
		return ShareApple(), nil
	case "share-google":
		return ShareGoogle(), nil
	case "spinner":
		return Spinner(), nil
	case "spinner-3":
		return SpinnerThree(), nil
	case "spinner-2":
		return SpinnerTwo(), nil
	case "star":
		return Star(), nil
	case "tag":
		return Tag(), nil
	case "trash":
		return Trash(), nil
	case "trophy":
		return Trophy(), nil
	case "undo":
		return Undo(), nil
	case "unlock":
		return Unlock(), nil
	case "user":
		return User(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in ei icon set", name)
	}
}
