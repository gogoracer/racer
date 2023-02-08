package circle_flags

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	acInnerSVG                    = `<mask id="circleFlagsAc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAc0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#6da544" d="m320 144l48-80l48 80z"/><circle cx="368" cy="144" r="48" fill="#acabb1"/><path fill="#338af3" d="M320 144v77c0 36 48 48 48 48s48-12 48-48v-77z"/><rect width="32" height="128" x="288" y="128" fill="#ff9811" rx="16" ry="16"/><rect width="32" height="128" x="416" y="128" fill="#ff9811" rx="16" ry="16"/><path fill="#6da544" d="m368 160l-48 67c2 11 9 19 16 26l32-45l32 45c8-7 14-15 16-26z"/></g>`
	adInnerSVG                    = `<mask id="circleFlagsAd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAd0)"><path fill="#0052b4" d="M0 0h144.7l36 254.6l-36 257.4H0z"/><path fill="#d80027" d="M367.3 0H512v512H367.3l-29.7-257.3z"/><path fill="#ffda44" d="M144.7 0h222.6v512H144.7z"/><path fill="#d80027" d="M256 354.5V256h66.8v47.3zm-66.8-165.3H256V256h-66.8z"/><path fill="#ff9811" d="M289.4 167a22.3 22.3 0 0 0-33.4-19.3a22.1 22.1 0 0 0-11.1-3c-12.3 0-22.3 10-22.3 22.3H167v111.3c0 41.4 32.9 65.4 58.7 77.8a22.1 22.1 0 0 0-3 11.2a22.3 22.3 0 0 0 33.3 19.3a22.1 22.1 0 0 0 11.1 3a22.3 22.3 0 0 0 19.2-33.5c25.8-12.4 58.7-36.4 58.7-77.8V167zm22.3 111.3c0 5.8 0 23.4-27.5 40.9a136.5 136.5 0 0 1-28.2 13.3c-7-2.4-17.8-6.7-28.2-13.3c-27.5-17.5-27.5-35.1-27.5-41v-77.9h111.4z"/></g>`
	aeInnerSVG                    = `<mask id="circleFlagsAe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAe0)"><path fill="#a2001d" d="M0 0h167l52.3 252L167 512H0z"/><path fill="#eee" d="m167 167l170.8-44.6L512 167v178l-173.2 36.9L167 345z"/><path fill="#6da544" d="M167 0h345v167H167z"/><path fill="#333" d="M167 345h345v167H167z"/></g>`
	afInnerSVG                    = `<mask id="circleFlagsAf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAf0)"><path fill="#d80027" d="M144.7 0h222.6l37 257.7l-37 254.3H144.7l-42.4-255.2z"/><path fill="#496e2d" d="M367.3 0H512v512H367.3z"/><path fill="#333" d="M0 0h144.7v512H0z"/><g fill="#ffda44"><path d="M256 167a89 89 0 1 0 0 178a89 89 0 0 0 0-178zm0 144.7a55.7 55.7 0 1 1 0-111.4a55.7 55.7 0 0 1 0 111.4z"/><path d="M256 222.6c-12.3 0-22.3 10-22.3 22.3v33.4h44.6v-33.4c0-12.3-10-22.3-22.3-22.3z"/></g></g>`
	afarInnerSVG                  = `<mask id="circleFlagsAfar0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAfar0)"><path fill="#eee" d="m0 160l256-32l256 32v192l-256 32L0 352Z"/><path fill="#0052b4" d="M0 0h512v160H0Z"/><path fill="#496e2d" d="M0 352h512v160H0Z"/><path fill="#d80027" d="M0 0v512l256-256L0 0z"/><path fill="#ff9811" d="m345 229l-12 11l40 40l-40 40l11 11l40-40l40 40l11-11l-40-40l40-40l-11-11l-40 40z"/><path fill="#d80027" d="m384 181l21 62l-54-38h66l-53 39z"/></g>`
	agInnerSVG                    = `<mask id="circleFlagsAg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAg0)"><path fill="#333" d="M0 .4h512l-34 229H36z"/><path fill="#ffda44" d="m367.3 205.3l-109.7 19.4l-112.9-19.4l45.5-21.3l-24.2-44l49.3 9.4l6.3-49.9l34.4 36.7l34.4-36.6l6.3 50L346 140l-24.2 44z"/><path fill="#0052b4" d="M25.6 205.3h466.8L257 439.5z"/><path fill="#eee" d="M34 307.4h446L256 511.6z"/><path fill="#a2001d" d="m0 511.6l256 .4L0 .4zm256 .4l256-.4V0z"/></g>`
	aiInnerSVG                    = `<mask id="circleFlagsAi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAi0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#496e2d" d="M445.2 256.1zm-155.8 0z"/><path fill="#eee" d="M433 293.6a62.4 62.4 0 0 0 12.2-37.5V144.8a55.4 55.4 0 0 1-33.4 11.1a55.6 55.6 0 0 1-44.5-22.2a55.6 55.6 0 0 1-44.5 22.2a55.4 55.4 0 0 1-33.4-11.1v111.3c0 15 5 27.3 12.3 37.5h131.2z"/><path fill="#ff9811" d="M409.8 235.5a91 91 0 0 0 6.3-27.6c0-10.1-13.2-18.3-13.2-18.3s-13.2 8.2-13.2 18.3a91 91 0 0 0 6.3 27.6l-7.6 17.1a38.3 38.3 0 0 0 29 0zm-51.5-55.6a91 91 0 0 0-27 8.3c-8.8 5-9.3 20.5-9.3 20.5s13.7 7.4 22.4 2.3c5.5-3.1 15-11.8 20.8-19.2l18.6-2a38.4 38.4 0 0 0-4.7-14a38.4 38.4 0 0 0-9.7-11.1zm-22.4 72.2a91 91 0 0 0 20.7 19.3c8.8 5 22.5-2.3 22.5-2.3s-.6-15.5-9.3-20.5a91 91 0 0 0-27-8.4l-11.1-15.1a38.4 38.4 0 0 0-9.7 11a38.4 38.4 0 0 0-4.8 14z"/><path fill="#338af3" d="M299 289.5c20.7 33.3 68.3 44.5 68.3 44.5s47.6-11.2 68.4-44.5H298.9Z"/></g>`
	alInnerSVG                    = `<mask id="circleFlagsAl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAl0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#333" d="M400.7 190H308a33.3 33.3 0 0 0-24.2-56.4a33.3 33.3 0 0 0-27.8 14.9a33.4 33.4 0 1 0-52 41.5h-92.7a45.8 45.8 0 0 0 46 44.5h-1.5c0 24.6 20 44.6 44.5 44.6c0 8 2.1 15.4 5.8 21.8l-37 37l28.4 28.3l40.2-40.2a30.5 30.5 0 0 0 4.9 1.4l-24.3 54.8L256 423l37.7-40.8l-24.3-54.8a30.4 30.4 0 0 0 4.9-1.4l40.2 40.2l28.3-28.3l-37-37a44.2 44.2 0 0 0 5.9-21.8c24.5 0 44.5-20 44.5-44.6h-1.5c24.6 0 46-19.9 46-44.5z"/></g>`
	amInnerSVG                    = `<mask id="circleFlagsAm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAm0)"><path fill="#0052b4" d="m0 166.9l253-26.7L512 167v178l-261.1 26L0 344.8z"/><path fill="#d80027" d="M0 0h512v166.9H0z"/><path fill="#ff9811" d="M0 344.9h512V512H0z"/></g>`
	anInnerSVG                    = `<mask id="circleFlagsAn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAn0)"><path fill="#eee" d="M0 0h171l85 32l85-32h171v171l-32 85l32 85v171H341l-85-32l-85 32H0V341l32-85l-32-85Z"/><path fill="#d80027" d="M171 0h170v512H171z"/><path fill="#0052b4" d="M512 171v170H0V171z"/><path fill="#eee" d="m236 247l52-37h-64l52 37l-20-61zm-45 79l52-37h-64l52 37l-20-61zm90 0l52-37h-64l52 37l-20-61zm74-47l52-37h-64l52 37l-20-61zm-238 0l52-37h-64l52 37l-20-61z"/></g>`
	aoInnerSVG                    = `<mask id="circleFlagsAo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAo0)"><path fill="#d80027" d="M0 0h512v256l-253 36.6L0 256z"/><path fill="#333" d="M0 256h512v256H0z"/><g fill="#ffda44"><path d="m220.9 203.6l21.7 15.8l-8.3 25.5L256 229l21.7 15.7l-8.3-25.5l21.7-15.7h-26.8L256 178l-8.3 25.5z"/><path d="M320 145.1a127.2 127.2 0 0 0-64-17v33.3a94 94 0 0 1 47.3 12.7a94.7 94.7 0 0 1-94.6 163.8a94 94 0 0 1-31.6-29.8l-27.9 18.4a128.1 128.1 0 0 0 217.7-6.5A128.1 128.1 0 0 0 320 145.1z"/><path d="M182.2 233.7a33.4 33.4 0 0 0 13.3 45.4l108.4 59.2c-7.4 13.5-3.4 30 10 37.3l29.3 16a27.8 27.8 0 0 0 37.8-11l16-29.3z"/></g></g>`
	aqInnerSVG                    = `<mask id="circleFlagsAq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAq0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#eee" d="m135 343l-41-70l17-38l-40-51l-9-37l74 51l45-11l19-67l50-29l75 11l87 45l4 74l28 10v76l-53 94l-64 20l-59-14l15-25l-7-26l-8 7z"/></g>`
	arInnerSVG                    = `<mask id="circleFlagsAr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAr0)"><path fill="#338af3" d="M0 0h512v144.7L488 256l24 111.3V512H0V367.3L26 256L0 144.7z"/><path fill="#eee" d="M0 144.7h512v222.6H0z"/><path fill="#ffda44" d="m332.4 256l-31.2 14.7l16.7 30.3l-34-6.5l-4.2 34.3l-23.7-25.2l-23.6 25.2l-4.3-34.3l-34 6.5l16.6-30.3l-31.2-14.7l31.3-14.7L194 211l34 6.5l4.3-34.3l23.6 25.2l23.6-25.2l4.4 34.3l34-6.5l-16.7 30.3z"/></g>`
	asInnerSVG                    = `<mask id="circleFlagsAs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAs0)"><path fill="#eee" d="M512 20.4V490L23.8 255.8z"/><path fill="#a2001d" d="M445.2 246.5h-30.5c8-9.6 7.5-23.7-1.5-32.7a24.2 24.2 0 0 0 0-34.2l-.5.5a25 25 0 0 0 .5-34.8l-137 137a23.9 23.9 0 0 0 34 0l2.6-2.5l65.6-6v28.3h22.3v-30.2l33.4-3z"/><path fill="#ffda44" d="M278.3 311.7L256 300.5l22.3-11.1H423v22.3z"/><path fill="#0052b4" d="M0 0v512h512L28.7 256.2L512 0z"/><path fill="#d80027" d="M512 0L0 256l512 256v-22L43.8 256L512 20.4z"/></g>`
	atInnerSVG                    = `<mask id="circleFlagsAt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAt0)"><path fill="#d80027" d="M0 0h512v167l-23.2 89.7L512 345v167H0V345l29.4-89L0 167z"/><path fill="#eee" d="M0 167h512v178H0z"/></g>`
	auInnerSVG                    = `<mask id="circleFlagsAu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAu0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#eee" d="m154 300l14 30l32-8l-14 30l25 20l-32 7l1 33l-26-21l-26 21l1-33l-33-7l26-20l-14-30l32 8zm222-27h47l-38 27l15-44l14 44zm7-162l7 15l16-4l-7 15l12 10l-15 3v17l-13-11l-13 11v-17l-15-3l12-10l-7-15l16 4zm57 67l7 15l16-4l-7 15l12 10l-15 3v16l-13-10l-13 11v-17l-15-3l12-10l-7-15l16 4zm-122 22l7 15l16-4l-7 15l12 10l-15 3v16l-13-10l-13 11v-17l-15-3l12-10l-7-15l16 4zm65 156l7 15l16-4l-7 15l12 10l-15 3v17l-13-11l-13 11v-17l-15-3l12-10l-7-15l16 4zM0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/></g>`
	auAboriginalInnerSVG          = `<mask id="circleFlagsAuAboriginal0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuAboriginal0)"><path fill="#333" d="M0 0h512v256l-256 32L0 256Z"/><path fill="#d80027" d="M0 256h512v256H0Z"/><circle cx="256" cy="256" r="128" fill="#ffda44"/></g>`
	auActInnerSVG                 = `<mask id="circleFlagsAuAct0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuAct0)"><path fill="#0052b4" d="M0 0h170l64 256l-64 256H0Z"/><path fill="#ffda44" d="M170 0h342v512H170z"/><path fill="#eee" d="m136 189l7 14.7l15.9-3.7l-7.2 14.7l12.7 10l-15.9 3.7v16.3l-12.6-10.2l-12.7 10.2v-16.3l-15.9-3.6l12.8-10.1L113 200l15.9 3.7zm-45-43l7 14.7l15.9-3.7l-7.2 14.7l12.7 10l-15.9 3.7v16.3l-12.6-10.2l-12.7 10.2v-16.3l-15.9-3.6l12.8-10.1L68 157l15.9 3.7zm19 92l5.6 17h17.8l-14.5 10.5l5.5 17l-14.5-10.5l-14.4 10.5l5.5-17L86.5 255h18zm-70-34l7 14.7l15.9-3.7l-7.2 14.7l12.7 10l-15.9 3.7v16.3l-12.6-10.2l-12.7 10.2v-16.3l-15.9-3.6l12.8-10.1L17 215l15.9 3.7zm51 99l-6 17l-17-8l8 17l-17 6l17 6l-8 17l17-8l6 17l6-17l17 8l-8-17l17-6l-17-6l8-17l-17 8z"/><g transform="translate(0 22)"><path fill="#0052b4" d="M276 200a17 17 0 0 0-17 17h-51a17 17 0 0 0 17 17a17 17 0 0 0 17 17a17 17 0 0 0 17 17h71v-68z"/><path fill="#eee" d="M384 200a17 17 0 0 1 17 17h51a17 17 0 0 1-17 17a17 17 0 0 1-17 17a17 17 0 0 1-17 17h-71v-68z"/><path fill="#0052b4" d="M285.5 200v68c0 34 44.5 44 44.5 44s44.5-10 44.5-44v-68z"/><path fill="#eee" d="m294.4 204l-2.5 7.5l26.8 10l-26.8 10l2.5 7.5l35.6-13.3l35.7 13.3l2.5-7.5l-26.9-10l26.9-10l-2.5-7.5l-35.7 13.3zm25 31.2V254h-6.2v-9.2h-15V284h63.7v-38.8h-15v8.8h-7.5v-18.8z"/><circle cx="330" cy="297.2" r="9" fill="#eee"/></g></g>`
	auNswInnerSVG                 = `<mask id="circleFlagsAuNsw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuNsw0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="M402 185h-40a74 74 0 0 0-51 51v40a74 74 0 0 0 51 51h40a74 74 0 0 0 51-51v-40a74 74 0 0 0-51-51z"/><path fill="#d80027" d="M382 182a74 74 0 0 0-20 3v51h-51a74 74 0 0 0-3 20a74 74 0 0 0 3 20h51v51a74 74 0 0 0 20 3a74 74 0 0 0 20-3v-51h51a74 74 0 0 0 3-20a74 74 0 0 0-3-20h-51v-51a74 74 0 0 0-20-3z"/><path fill="#ffda44" d="m382 295l-3 8l-8-4l4 8l-8 3l8 3l-4 8l8-4l3 8l3-8l8 4l-4-8l8-3l-8-3l4-8l-8 4zm54-54l-3 8l-8-4l4 8l-8 3l8 3l-4 8l8-4l3 8l3-8l8 4l-4-8l8-3l-8-3l4-8l-8 4zm-108 0l-3 8l-8-4l4 8l-8 3l8 3l-4 8l8-4l3 8l3-8l8 4l-4-8l8-3l-8-3l4-8l-8 4zm54-54l-3 8l-8-4l4 8l-8 3l8 3l-4 8l8-4l3 8l3-8l8 4l-4-8l8-3l-8-3l4-8l-8 4zm-10 52a7 7 0 0 0-6 6v5h-3v-2c0-3-3-6-6-6l-1 2v12l2 1h8a6 6 0 0 0 6 7v2h-3a5 5 0 0 0-5 6l1 1h12l2-1v-8l17-1v3h-3a5 5 0 0 0-5 5l1 2h13l1-2v-16c3 0 5-2 5-5c0-2-2-5-5-5h-16l-2-1l2-2h16v-3h-16a5 5 0 0 0-5 5c0 2 2 5 5 5h16l2 1l-2 2h-18a6 6 0 0 1-6-7a7 7 0 0 0-7-6z"/></g>`
	auNtInnerSVG                  = `<mask id="circleFlagsAuNt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuNt0)"><path fill="#333" d="M0 0h171l64 256l-64 256H0Z"/><path fill="#ff9811" d="M171 0h341v512H171z"/><path fill="#333" d="M341 282c6-3 13-3 18 0c-5-3-6-11-5-17c2-6 7-13 13-14c-6 1-13-3-17-7c-4-5-6-13-4-19c-2 6-10 9-16 9s-13-3-16-9c2 6 0 14-4 19c-3 4-10 8-16 7c6 1 10 8 12 14c1 6 0 14-5 17c5-3 13-3 18 0c6 2 11 8 11 15c0-7 5-13 11-15z"/><path fill="#eee" d="M358 237c14 18 75 0 45-37c-29-37-60 18-45 37z"/><path fill="#eee" d="M330 224c23 0 47-59 0-59s-23 59 0 59z"/><path fill="#eee" d="M303 237c14-19-17-74-46-37c-30 37 31 55 45 37zm42 53c-21 11-17 73 26 53c42-20-5-63-26-53z"/><path fill="#eee" d="M315 290c-21-10-68 33-26 53c43 20 47-42 26-53zm49-23c-5 22 47 58 57 13c11-46-52-36-57-13z"/><path fill="#eee" d="M296 267c-6-23-68-33-57 13c10 45 62 9 57-13zm-160-78l7 15l16-4l-7 15l12 10l-15 3v17l-13-11l-13 11v-17l-16-3l13-10l-7-15l16 4zm-45-43l7 15l16-4l-7 15l12 10l-15 3v17l-13-11l-13 11v-17l-16-3l13-10l-7-15l16 4zm19 92l6 17h17l-14 11l5 17l-14-11l-14 11l5-17l-14-11h18zm-70-34l7 15l16-4l-7 15l12 10l-16 3v17l-12-11l-13 11v-17l-16-3l13-10l-7-15l16 4zm51 99l-6 17l-17-8l8 17l-17 6l17 6l-8 17l17-8l6 17l6-17l17 8l-8-17l17-6l-17-6l8-17l-17 8z"/></g>`
	auQldInnerSVG                 = `<mask id="circleFlagsAuQld0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuQld0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><circle cx="382" cy="256" r="74" fill="#eee"/><path fill="#338af3" d="M405 322.1L382 312l-23 10.2l23-66.1zm0-132.2L382 200l-23-10.2l23 66.1zM315.9 279l10.2-23l-10.2-23l66.1 23zm132.2 0L438 256l10.2-23l-66.1 23z"/><path fill="#ffda44" d="M377 228v5h-5v10h5v9.4a15.2 15.2 0 0 0-20.7 22.2v9.4H408v-9.4a15.2 15.2 0 0 0-20.7-22.2V243h5v-10h-5v-5z"/></g>`
	auSaInnerSVG                  = `<mask id="circleFlagsAuSa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuSa0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><circle cx="382" cy="256" r="74" fill="#ff9811"/><path fill="#333" d="M382 228a14 14 0 0 0-14 14h-42a14 14 0 0 0 14 14a14 14 0 0 0 14 14a14 14 0 0 0 14 14h28a14 14 0 0 0 14-14a14 14 0 0 0 14-14a14 14 0 0 0 14-14h-42a14 14 0 0 0-14-14z"/></g>`
	auTasInnerSVG                 = `<mask id="circleFlagsAuTas0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuTas0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><circle cx="382" cy="256" r="74" fill="#eee"/><path fill="#d80027" d="M363 221a14 14 0 0 0-14 14v10h-5v-6c0-6-5-11-11-11a3 3 0 0 0-3 3v25c0 1 1 3 3 3h16a13 13 0 0 0 13 12v6h-5c-7 0-12 5-12 11l3 3h25c2 0 3-2 3-3v-17h34v5h-5c-6 0-11 5-11 11c0 2 1 3 3 3h24c2 0 3-1 3-3v-33c5 0 10-4 10-9c0-6-5-10-10-10h-32a3 3 0 0 1-3-3c0-2 1-3 3-3h32v-7h-32a10 10 0 0 0 0 19h32c2 0 3 2 3 4c0 1-1 3-3 3h-35c-7 0-13-6-13-13a14 14 0 0 0-13-14z"/></g>`
	auVicInnerSVG                 = `<mask id="circleFlagsAuVic0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuVic0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="m313 267l7 15l15-4l-7 15l13 10l-16 3v16l-13-10l-12 10v-16l-16-3l13-10l-7-15l16 4zm66-61l7 15l15-4l-7 15l13 10l-16 4v16l-13-10l-12 10v-16l-16-4l13-10l-7-15l16 4zm1 146l40-29h-50l40 29l-15-48Zm-1 29l5 15l14-7l-7 15l15 5l-15 5l7 15l-14-7l-5 15l-6-15l-14 7l7-15l-15-5l15-5l-7-15l14 7zm57-114l-8 14h-16l8 14l-8 14h16l8 13l8-13h16l-8-14l8-14h-16z"/><path fill="#ffda44" d="M370 93v8h-8v17h8v15a25 25 0 0 0-34 37v15h85v-15a25 25 0 0 0-34-37v-15h8v-17h-8v-8z"/></g>`
	auWaInnerSVG                  = `<mask id="circleFlagsAuWa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAuWa0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><circle cx="382" cy="256" r="74" fill="#ffda44"/><path fill="#333" d="M397 299.8c-51.4 7.8-53.2-24-53.2-24c-2-22.9 21.4-29.6 28.4-42.3c7-12.6-9.3-14.5-8.3-6.4c1 8.2-3.8 9.6-4 10.2c-.3.7-7.3 7.6-8.3 6.3c-1-1.3 2.9-6.3 2.9-6.3s-3-1.7-4.2-5.6c-3.2-8.8 13-30.7 29.9-15.9c16.8 14.8-5.6 32-11.4 40.4c-5.8 8.3-5.6 14.6-2 15.7c3.7 1.2 6.5-2.5 8-8.1c1.5-5.3 10.8-22.7 30.5-16.9c14.7 4.4 12.9 10.5 24.7 7.3c4.1 39.6-33 45.6-33 45.6z"/></g>`
	awInnerSVG                    = `<mask id="circleFlagsAw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAw0)"><path fill="#ffda44" d="m0 322.8l253.6-18.4L512 322.8v33.4l-258 15.3L0 356.2zm0 66.8l257.2-13.8L512 389.6V423l-253 16.9L0 423z"/><path fill="#338af3" d="M0 0h512v322.8H0zm0 356.2h512v33.4H0zM0 423h512v89H0z"/><path fill="#eee" d="m117.3 161.5l-50-22.1l50-22l22-50.1l22.2 50l50 22l-50 22.2l-22.1 50z"/><path fill="#d80027" d="m139.4 94.9l13.6 30.9l31 13.6l-31 13.6l-13.6 31l-13.6-31l-31-13.6l31-13.6z"/></g>`
	axInnerSVG                    = `<mask id="circleFlagsAx0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAx0)"><path fill="#0052b4" d="M0 0h100.2l68.3 40.7L233.7 0H512v189.2l-45.5 66l45.5 68.6V512H233.7l-65.8-39.2l-67.7 39.2H0V322.8l45.6-67.5L0 189.2z"/><path fill="#ffda44" d="M100.2 0v189.2H0v33.4l23 34l-23 32.8v33.4h100.2V512h33.4l33.9-22.6l32.8 22.6h33.4V323.8H512v-34.4l-24.2-32.2l24.2-34.6v-33.4H233.7V0h-33.4l-32.6 20l-34.1-20z"/><path fill="#d80027" d="M133.6 0v222.6H0v66.8h133.6V512h66.7V289.4H512v-66.8H200.3V0h-66.7z"/></g>`
	azInnerSVG                    = `<mask id="circleFlagsAz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAz0)"><path fill="#d80027" d="m0 166.9l253-31.8l259 31.8v178l-257.5 37.4L0 345z"/><path fill="#338af3" d="M0 0h512v166.9H0z"/><path fill="#6da544" d="M0 344.9h512V512H0z"/><g fill="#eee"><path d="M261.6 328.2a72.3 72.3 0 1 1 34.4-136a89 89 0 1 0 0 127.3a72 72 0 0 1-34.4 8.7z"/><path d="m317.2 206l9.6 26.8l25.8-12.3l-12.2 25.8l26.9 9.6l-27 9.6l12.3 25.8l-25.8-12.3l-9.6 27l-9.6-27l-25.8 12.3l12.3-25.8l-27-9.6l27-9.6l-12.3-25.8l25.8 12.3z"/></g></g>`
	baInnerSVG                    = `<mask id="circleFlagsBa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBa0)"><path fill="#ffda44" d="M0 0h445.3l33.9 255l-33.9 257l-323.7-134.3L0 66.8z"/><path fill="#0052b4" d="M0 66.8V512h445.4z"/><path fill="#0052b4" d="M445.3 0H512v512h-66.7z"/><path fill="#eee" d="m354.6 456l-8.3 25.6h-26.8l21.7 15.8l-8.3 25.5l21.7-15.8l21.7 15.8l-8.3-25.5l21.7-15.8h-26.8zm-55-55.4l-8.3 25.5h-26.8l21.7 15.8l-8.3 25.5l21.7-15.8l21.7 15.8l-8.3-25.5l21.7-15.8h-26.8zM244.4 345l-8.3 25.5h-26.8l21.7 15.8l-8.3 25.5l21.7-15.8l21.7 15.8l-8.3-25.5l21.7-15.8h-26.8zm-55.1-55.7l-8.3 25.5h-26.8l21.7 15.8l-8.3 25.5l21.7-15.8L211 356l-8.3-25.5l21.7-15.8h-26.8zm-55.4-55.7l-8.3 25.5H98.8l21.7 15.8l-8.3 25.5l21.7-15.8l21.7 15.8l-8.3-25.5L169 259h-26.8zM78.7 178l-8.3 25.5H43.6l21.7 15.8l-8.3 25.5L78.7 229l21.7 15.8l-8.3-25.5l21.7-15.8H87zm-55.2-55.7l-8.3 25.5h-26.8l21.7 15.8L1.8 189l21.7-15.8L45.2 189l-8.3-25.5l21.7-15.8H31.8z"/></g>`
	bbInnerSVG                    = `<mask id="circleFlagsBb0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBb0)"><path fill="#0052b4" d="M0 0h144.8l112.9 36.7L367.4 0H512v512H367.4l-108.9-38.1L144.8 512H0z"/><path fill="#ffda44" d="M144.8 0h222.6v512H144.8z"/><path fill="#333" d="m334.1 155.8l14.8 7.5l-14.9-7.5l-15-7.4c-.8 1.8-20.3 41.4-23.5 102h-22.7v-94.6l-16.7-22.2l-16.7 22.2v94.6h-22.7a278.3 278.3 0 0 0-23.6-102l-29.8 14.9c.2.4 20.5 41.7 20.5 103.8v16.7h55.6v94.6h33.4v-94.6h55.6v-16.7c0-32 5.6-58.6 10.3-75.1c5-18 10.2-28.6 10.3-28.7l-15-7.5z"/></g>`
	bdInnerSVG                    = `<mask id="circleFlagsBd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBd0)"><path fill="#496e2d" d="M0 0h512v512H0z"/><circle cx="200.3" cy="256" r="111.3" fill="#d80027"/></g>`
	beInnerSVG                    = `<mask id="circleFlagsBe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBe0)"><path fill="#333" d="M0 0h167l38.2 252.6L167 512H0z"/><path fill="#d80027" d="M345 0h167v512H345l-36.7-256z"/><path fill="#ffda44" d="M167 0h178v512H167z"/></g>`
	bfInnerSVG                    = `<mask id="circleFlagsBf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBf0)"><path fill="#d80027" d="M0 0h512v256l-255.2 48L0 256z"/><path fill="#6da544" d="M0 256h512v256H0z"/><path fill="#ffda44" d="m256 167l19.3 59.5H338l-50.6 36.8l19.3 59.5L256 286l-50.6 36.8l19.3-59.5l-50.6-36.8h62.6z"/></g>`
	bgInnerSVG                    = `<mask id="circleFlagsBg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBg0)"><path fill="#496e2d" d="m0 166.9l258-31.7l254 31.7v178l-251.4 41.3L0 344.9z"/><path fill="#eee" d="M0 0h512v166.9H0z"/><path fill="#d80027" d="M0 344.9h512V512H0z"/></g>`
	bhInnerSVG                    = `<mask id="circleFlagsBh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBh0)"><path fill="#eee" d="M0 0h182.5l88.1 268.5l-88 243.5H0z"/><path fill="#d80027" d="m182.5 0l-82.3 42.7l82.3 42.7l-82.3 42.6l82.3 42.7l-82.3 42.7l82.3 42.6l-82.3 42.7l82.3 42.7l-82.3 42.6l82.3 42.7l-82.3 42.7l82.3 42.6H512V0H182.5z"/></g>`
	biInnerSVG                    = `<mask id="circleFlagsBi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBi0)"><path fill="#eee" d="M0 0h47.2l207.5 30L464.8 0H512v47.2L477.4 256L512 464.8V512h-47.2l-209.1-35.8L47.2 512H0v-47.2l32.8-202.7L0 47.2z"/><path fill="#d80027" d="M47.2 0L256 208.8L464.8 0H47.2zM256 303.2L47.2 512h417.6L256 303.2z"/><path fill="#6da544" d="M0 47.2v417.6L208.8 256L0 47.2zm512 0L303.2 256L512 464.8V47.2z"/><circle cx="256" cy="256" r="111.3" fill="#eee"/><path fill="#d80027" d="m256 178l9.6 16.8H285l-9.6 16.7l9.6 16.7h-19.3l-9.6 16.7l-9.6-16.7H227l9.6-16.7l-9.6-16.7h19.3zm-49 78l9.6 16.7H236l-9.6 16.7l9.6 16.7h-19.3l-9.6 16.7l-9.6-16.7H178l9.6-16.7l-9.6-16.7h19.3zm98 0l9.6 16.7H334l-9.6 16.7l9.6 16.7h-19.3l-9.6 16.7l-9.6-16.7H276l9.6-16.7l-9.6-16.7h19.3z"/></g>`
	bjInnerSVG                    = `<mask id="circleFlagsBj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBj0)"><path fill="#6da544" d="M0 0h189.2l54 257.6l-54 254.4H0z"/><path fill="#ffda44" d="M189.2 0H512v256l-159 53.5L189.1 256z"/><path fill="#d80027" d="M189.2 256H512v256H189.2z"/></g>`
	blInnerSVG                    = `<mask id="circleFlagsBl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBl0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#acabb1" d="M167 178a28 28 0 0 0-28 28H55a28 28 0 0 0 28 28a28 28 0 0 0 28 28a28 28 0 0 0 28 28h234a28 28 0 0 0 28-28a28 28 0 0 0 28-28a28 28 0 0 0 28-28h-84a28 28 0 0 0-28-28z"/><path fill="#ffda44" d="M123 357h44v44h-44zm222 0h44v44h-44zm-178 11h178v44H167zm67-268v33.5L223 145l-12-6v-17h-44v56l89 14l89-14v-56h-44v17l-12 6l-11-11.5V100z"/><path fill="#0052b4" d="M167 178v112c0 68 89 88.5 89 88.5s89-20.4 89-88.5V178z"/><path fill="#d80027" d="M167 222.1h178v69H167z"/></g>`
	bmInnerSVG                    = `<mask id="circleFlagsBm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBm0)"><path fill="#d80027" d="M256 0h256v512H0V256z"/><path fill="#eee" d="M0 0h33.4l31.8 16.4l35-16.4H256v133.6l-9.3 33.7l9.3 41.5V256h-47.2l-39.3-7l-35.9 7.1L0 256V100.2l15.4-34.5L0 33.4z"/><path fill="#496e2d" d="M445.2 256.1zm-155.8 0z"/><path fill="#d80027" d="m267 235.5l-102-102h-31.4L267 267z"/><path fill="#d80027" d="M33.4 0v33.4H0v66.8h33.4v170.6h66.8V100.2h170.2V33.4H100.2V0z"/><path fill="#0052b4" d="M180.8 133.6H256v75.2zm-47.2 47.2v75.3l75.2-.1z"/><path fill="#eee" d="M289.4 133.6V256c0 59.6 155.8 59.6 155.8 0V133.6z"/><path fill="#6da544" d="M289.4 256c0 59.6 77.9 78 77.9 78s78-18.4 78-78h-156z"/><path fill="#a2001d" d="m367.3 207l-36.2 15.6V256l36.2 22.3l36.2-22.3v-33.4z"/><path fill="#338af3" d="M331.1 189.2h72.4v33.4H331z"/></g>`
	bnInnerSVG                    = `<mask id="circleFlagsBn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBn0)"><path fill="#ffda44" d="M0 0h512v326.7l-19.3 76.5l19.3 77.7V512H0V185.2l21.4-76.5L0 31z"/><path fill="#eee" d="M0 31v117.2l512 295.7V326.7L0 31z"/><path fill="#333" d="M0 108.2v77L512 481v-77L0 108.2z"/><g fill="#d80027"><path d="M328.3 228.2a72.3 72.3 0 1 1-136-34.4a89 89 0 1 0 127.3 0a72 72 0 0 1 8.7 34.4z"/><path d="M239.3 144.7h33.4v167h-33.4z"/><path d="M311.6 178H200.4c0 7.8 6.6 14 14.3 14h-.4a14 14 0 0 0 13.9 14a14 14 0 0 0 13.9 13.8h27.8a14 14 0 0 0 14-13.9a14 14 0 0 0 13.8-13.9h-.4c7.6 0 14.3-6.2 14.3-13.9zM178.1 322.9h155.8v33.4H178.1z"/><path d="M289.4 333.9h66.8v33.4h-66.8zm-133.6 0h66.8v33.4h-66.8z"/></g></g>`
	boInnerSVG                    = `<mask id="circleFlagsBo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBo0)"><path fill="#ffda44" d="m0 167l252.9-29.3L512 167v178l-255.7 25.7L0 345z"/><path fill="#d80027" d="M0 0h512v167H0z"/><path fill="#6da544" d="M0 345h512v167H0z"/></g>`
	bqInnerSVG                    = `<mask id="circleFlagsBq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBq0)"><path fill="#eee" d="M113.7 119.8L276 0h236v31.7L306 289.5L31.6 512H0V276z"/><path fill="#ffda44" d="M0 0v276L276 0H0z"/><path fill="#0052b4" d="M512 31.7L31.7 512H512V31.7z"/><path fill="#333" d="m255 245.7l22.1-12l-22-12a78 78 0 0 0-65-65l-12-22l-12 22a78 78 0 0 0-65 65l-22 12l22 12a78 78 0 0 0 65 65l12 22.1l12-22a78 78 0 0 0 65-65zm-77 32.6a44.5 44.5 0 1 1 0-89a44.5 44.5 0 0 1 0 89z"/><path fill="#d80027" d="m178 200.3l9.7 16.7H207l-9.6 16.7l9.6 16.7h-19.3l-9.6 16.7l-9.7-16.7h-19.2l9.6-16.7l-9.6-16.7h19.2z"/></g>`
	bqBoInnerSVG                  = `<mask id="circleFlagsBqBo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBqBo0)"><path fill="#eee" d="M113.7 119.8L276 0h236v31.7L306 289.5L31.6 512H0V276z"/><path fill="#ffda44" d="M0 0v276L276 0H0z"/><path fill="#0052b4" d="M512 31.7L31.7 512H512V31.7z"/><path fill="#333" d="m255 245.7l22.1-12l-22-12a78 78 0 0 0-65-65l-12-22l-12 22a78 78 0 0 0-65 65l-22 12l22 12a78 78 0 0 0 65 65l12 22.1l12-22a78 78 0 0 0 65-65zm-77 32.6a44.5 44.5 0 1 1 0-89a44.5 44.5 0 0 1 0 89z"/><path fill="#d80027" d="m178 200.3l9.7 16.7H207l-9.6 16.7l9.6 16.7h-19.3l-9.6 16.7l-9.7-16.7h-19.2l9.6-16.7l-9.6-16.7h19.2z"/></g>`
	bqSaInnerSVG                  = `<mask id="circleFlagsBqSa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBqSa0)"><path fill="#d80027" d="M0 0h512v256H0z"/><path fill="#0052b4" d="M0 256h512v256H0z"/><path fill="#eee" d="M0 256L256 0l256 256l-256 256z"/><path fill="#ffda44" d="m256 133.6l27.6 85H373L300.7 271l27.6 85l-72.3-52.5l-72.3 52.6l27.6-85l-72.3-52.6h89.4z"/></g>`
	bqSeInnerSVG                  = `<mask id="circleFlagsBqSe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBqSe0)"><path fill="#0052b4" d="M0 44.6L255.3 18L512 44.6v194.7L493.7 256l18.3 16.7v194.7l-255.5 27.5L0 467.4V272.7l17.5-17.3L0 239.3z"/><path fill="#eee" d="m100.2 256.1l155.8-98l155.8 98l-155.8 98z"/><path fill="#6da544" d="M167 300.6h200.3l-44.5-66.8l-22.3 22.3l-22.2-11l-22.3 33.3h-44.5l-22.3-22.3z"/><path fill="#ffda44" d="m256 189.3l5.5 17h18l-14.6 10.5l5.6 17l-14.5-10.5l-14.5 10.5l5.6-17l-14.5-10.5h17.9z"/><path fill="#d80027" d="M0 0v44.6h239.4v96l-157 98.7H0v33.4h82l157.3 99v95.7H0V512h512v-44.6H272.7v-95.8L430 272.7h82v-33.4h-82.4l-156.9-98.6v-96H512V0H0zm256 169.6l137.6 86.5L256 342.6l-137.6-86.5L256 169.6z"/></g>`
	brInnerSVG                    = `<mask id="circleFlagsBr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBr0)"><path fill="#6da544" d="M0 0h512v512H0z"/><path fill="#ffda44" d="M256 100.2L467.5 256L256 411.8L44.5 256z"/><path fill="#eee" d="M174.2 221a87 87 0 0 0-7.2 36.3l162 49.8a88.5 88.5 0 0 0 14.4-34c-40.6-65.3-119.7-80.3-169.1-52z"/><path fill="#0052b4" d="M255.7 167a89 89 0 0 0-41.9 10.6a89 89 0 0 0-39.6 43.4a181.7 181.7 0 0 1 169.1 52.2a89 89 0 0 0-9-59.4a89 89 0 0 0-78.6-46.8zM212 250.5a149 149 0 0 0-45 6.8a89 89 0 0 0 10.5 40.9a89 89 0 0 0 120.6 36.2a89 89 0 0 0 30.7-27.3A151 151 0 0 0 212 250.5z"/></g>`
	bsInnerSVG                    = `<mask id="circleFlagsBs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBs0)"><path fill="#338af3" d="M0 0h512v167l-37.4 89l37.4 89v167H0l49.6-252z"/><path fill="#ffda44" d="M108.3 167H512v178H108.3z"/><path fill="#333" d="M0 0v512l256-256L0 0z"/></g>`
	btInnerSVG                    = `<mask id="circleFlagsBt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBt0)"><path fill="#ffda44" d="M0 0h512L281 293.2L0 512z"/><path fill="#ff9811" d="M506.7 0L0 506.7v5.3h512V0h-5.3z"/><path fill="#eee" d="M398.2 142V97.7h-85.3l-6.5 6.5c-23.2 23.2-19.3 50.2-16.5 70c2.7 18.3 3 25.3-3.8 32c-6.8 7-14 6.6-32.2 4c-19.7-3-46.8-6.8-70 16.4c-23.2 23.2-19.3 50.3-16.4 70c2.6 18.3 3 25.4-4 32.2c-6.7 6.8-13.7 6.5-32 3.8c-7.2-1-14.5-2-22.4-2.2l-.5 44.6c5 0 10.6.8 16.5 1.7c7 1 15 2.2 23.5 2.2c8.7 0 18-1.3 27-5.2v35.6h66.8v-33.4H209v-22.2h22.3v-33.4h-18c.9-10.8-.6-21.2-2-30c-2.5-18.4-3-25.4 4-32.2c6.8-6.8 13.8-6.5 32.2-3.8c14.3 2 32.6 4.7 50.4-3v35.6h66.8v-33.4h-33.3V231h22.3v-33.4h-17.8a130 130 0 0 0-2-30c-1.8-12.8-2.5-20-.4-25.6z"/></g>`
	bvInnerSVG                    = `<mask id="circleFlagsBv0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBv0)"><path fill="#d80027" d="M0 0h100.2l66.1 53.5L233.7 0H512v189.3L466.3 257l45.7 65.8V512H233.7l-68-50.7l-65.5 50.7H0V322.8l51.4-68.5l-51.4-65z"/><path fill="#eee" d="M100.2 0v189.3H0v33.4l24.6 33L0 289.5v33.4h100.2V512h33.4l30.6-26.3l36.1 26.3h33.4V322.8H512v-33.4l-24.6-33.7l24.6-33v-33.4H233.7V0h-33.4l-33.8 25.3L133.6 0z"/><path fill="#0052b4" d="M133.6 0v222.7H0v66.7h133.6V512h66.7V289.4H512v-66.7H200.3V0z"/></g>`
	bwInnerSVG                    = `<mask id="circleFlagsBw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBw0)"><path fill="#338af3" d="M0 0h512v178l-31 76.9l31 79.1v178H0V334l37-80.7L0 178z"/><path fill="#333" d="m0 211.5l256-19.2l256 19.2v89l-254.6 20.7L0 300.5z"/><path fill="#eee" d="M0 178h512v33.5H0zm0 122.5h512V334H0z"/></g>`
	byInnerSVG                    = `<mask id="circleFlagsBy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBy0)"><path fill="#eee" d="M0 0h155.8l35 254.6l-35 257.4H0z"/><path fill="#a2001d" d="M155.8 0H512v345.1l-183 37.4l-173.2-37.4z"/><path fill="#6da544" d="M155.8 345.1H512V512H155.8z"/><path fill="#a2001d" d="M50 .2L22.3 50l27.8 50.4L77.9 50zm55.8 0L78 50l27.7 50.4l28-50.4zM50 137.5l-27.7 49.6l27.8 50.5l27.7-50.5zm55.8 0L78 187.1l27.7 50.5l28-50.5zM50 274.7l-27.7 49.7l27.8 50.4l27.8-50.4zm55.8 0L78 324.4l27.7 50.4l28-50.4zM50 412l-27.7 49.6l27.8 50.5l27.7-50.5zm55.8 0L78 461.6l27.7 50.5l28-50.5z"/></g>`
	bzInnerSVG                    = `<mask id="circleFlagsBz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBz0)"><path fill="#0052b4" d="m0 44.5l257.8-23.7L512 44.5v423l-252.5 26L0 467.5z"/><circle cx="256" cy="256" r="122.4" fill="#eee"/><circle cx="256" cy="256" r="100.2" fill="#6da544"/><circle cx="256" cy="256" r="66.8" fill="#eee"/><path fill="#0052b4" d="M256 239.3L219.8 256v27.8l36.2 22.3l36.2-22.3V256z"/><path fill="#ffda44" d="M219.8 222.6h72.4V256h-72.4z"/><path fill="#a2001d" d="M0 0h512v44.5H0zm0 467.5h512V512H0z"/></g>`
	caInnerSVG                    = `<mask id="circleFlagsCa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCa0)"><path fill="#d80027" d="M0 0v512h144l112-64l112 64h144V0H368L256 64L144 0Z"/><path fill="#eee" d="M144 0h224v512H144Z"/><path fill="#d80027" d="m301 289l44-22l-22-11v-22l-45 22l23-44h-23l-22-34l-22 33h-23l23 45l-45-22v22l-22 11l45 22l-12 23h45v33h22v-33h45z"/></g>`
	caBcInnerSVG                  = `<mask id="circleFlagsCaBc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCaBc0)"><path fill="#eee" d="M0 61.2L80.7 0h16l18.5 6l12.9-6h94.6l33.8 19.2l33-19.2h63.1l11.9 7.3L384 0h47.2L512 61.2v33.4L491.3 128l20.7 33.4v33.4l-22 29.7l22 31.5v13.1l-9.7 19l9.7 15.3v32.3l-12.7 19.8L512 370v32.8l-28.4 57l-446.2-1l-37.3-56V370l13.4-13.7L0 335.7v-32.3l14.2-14.2L0 269v-13l27.4-30.6L0 194.8v-33.4l25.5-33.5L0 94.6z"/><path fill="#0052b4" d="M0 0v61.2h142.1L80.7 0H0zm431.3 0l-61.2 61.2H512V0h-80.7zM0 194.8V256h80.9l61.2-61.2H0zm370.1 0l61.2 61.2H512v-61.2H370.1zM189.3 242l-14 14h14v-14zm133.6 0v14h14l-14-14zM0 269.1v34.3h.1c21.3 0 21.3 19.5 42.6 19.5c21.4 0 21.4-19.5 42.7-19.5c21.3 0 21.3 19.5 42.7 19.5c21.4 0 21.3-19.5 42.6-19.5c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.3 0 21.3 19.5 42.6 19.5c21.4 0 21.4-19.5 42.7-19.5c21.3 0 21.3 19.5 42.7 19.5c21.4 0 21.3-19.4 42.5-19.5v-34.3c-21.2 0-21.2 19.5-42.5 19.5s-21.4-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.6 19.5s-21.4-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5H0zm0 66.6V370h.1c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.3 0 21.3 19.5 42.6 19.5c21.4 0 21.4-19.5 42.7-19.5c21.3 0 21.3 19.5 42.7 19.5c21.4 0 21.3-19.5 42.6-19.5c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.3 0 21.3 19.5 42.6 19.5s21.4-19.4 42.5-19.5v-34.3c-21.1.1-21.2 19.5-42.5 19.5s-21.3-19.5-42.6-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.6 19.5s-21.4-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5H0zm0 67.1V512h92.6l168.1-21.3L419.4 512H512V402.8c-21.2 0-21.2 19.5-42.5 19.5s-21.4-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.6 19.5s-21.4-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5H0z"/><path fill="#d80027" d="m96.7 0l61.2 61.2h31.4L128.1 0H96.7zm126 0v94.6H0v66.8h222.6V256h66.8v-94.6H512V94.6H289.5V0h-66.8zm129.9 0l-29.7 29.7v31.5L384.1 0h-31.5zM189.4 194.8L128.1 256h31.5l29.8-29.7v-31.5zm133.5 0l61.2 61.2h31.5l-61.2-61.2h-31.5z"/><path fill="#ffda44" d="m201.9 345l-10 78.7l-77.8-14.9l37.7 68.7A268.2 268.2 0 0 0 92.6 512h326.8a268.2 268.2 0 0 0-59-34.5l37.7-68.7l-77.9 15l-9.9-78.8l-54.2 58l-54.2-58z"/></g>`
	ccInnerSVG                    = `<mask id="circleFlagsCc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCc0)"><path fill="#6da544" d="M0 0h512v512H0z"/><g fill="#ffda44"><path d="m393 367.3l7 14.7l15.9-3.7l-7.1 14.7l12.7 10l-15.9 3.7V423L393 412.8L380.3 423v-16.3l-15.9-3.6l12.8-10.1l-7.1-14.7L386 382zm-65.4-155.8l7 14.7l16-3.7l-7.2 14.7l12.8 10l-16 3.6l.1 16.3l-12.7-10.2l-12.7 10.2v-16.3l-15.8-3.5l12.7-10.1l-7-14.7l15.8 3.7zm65.4-89l7 14.6l15.9-3.6l-7.1 14.6l12.7 10.1l-15.9 3.6v16.3L393 167.9l-12.7 10.2v-16.3l-15.9-3.6l12.8-10l-7.1-14.7L386 137zm57 66.7l7 14.7l16-3.7l-7.1 14.7l12.7 10.1l-15.9 3.6v16.3l-12.6-10.2l-12.7 10.2v-16.3l-15.9-3.6l12.8-10.1l-7.1-14.7L443 204zm-40.8 78l5.6 17h17.9l-14.5 10.4l5.5 17l-14.5-10.5l-14.4 10.6l5.5-17l-14.5-10.6h18z"/><path d="M283.8 328.3a72.3 72.3 0 1 1 34.4-136a89 89 0 1 0 0 127.3a72 72 0 0 1-34.4 8.7zM161 100.2a44.5 44.5 0 0 0-77.1 0h77zm.1.2l-38.7 38.7l-38.7-38.7a44.3 44.3 0 0 0-5.8 22a44.5 44.5 0 1 0 89 0a44.3 44.3 0 0 0-5.8-22z"/></g><path fill="#a2001d" d="M111.3 128v37.6a44.6 44.6 0 0 0 11.1 1.4c3.9 0 7.6-.5 11.2-1.4V128h-22.3z"/></g>`
	cdInnerSVG                    = `<mask id="circleFlagsCd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCd0)"><path fill="#338af3" d="M0 0h401.9L512 110.3V512H110.3L0 401.9z"/><path fill="#ffda44" d="M401.9 0L0 401.9V449l63 63h47.3L512 110.3V63L449 0z"/><path fill="#d80027" d="M449 0L0 449v63h63L512 63V0h-63z"/><path fill="#ffda44" d="m136.4 78l13.8 42.4H195l-36 26.3l13.7 42.5l-36.2-26.3l-36 26.3l13.7-42.5L78 120.4h44.7z"/></g>`
	cfInnerSVG                    = `<mask id="circleFlagsCf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCf0)"><path fill="#ffda44" d="m0 378.4l252.9-28.8L512 378.4V512H322.8L256 481l-66.8 31H0z"/><path fill="#6da544" d="m0 256l249.8-28L512 256v122.4H0z"/><path fill="#eee" d="m0 133.6l255.3-28.3L512 133.6V256H0z"/><path fill="#0052b4" d="M0 0h189.2L256 30l66.8-30H512v133.6H0z"/><path fill="#ffda44" d="m137.7 55.7l6.9 21.2H167L148.9 90l6.9 21.3l-18.1-13.1l-18 13.1l6.8-21.3l-18-13h22.3z"/><path fill="#d80027" d="M189.2 0h133.6v512H189.2z"/></g>`
	cgInnerSVG                    = `<mask id="circleFlagsCg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCg0)"><path fill="#ffda44" d="M384 0h128v128L352 352L128 512H0V384l160-224Z"/><path fill="#6da544" d="M0 384L384 0H0Z"/><path fill="#d80027" d="M512 128L128 512h384z"/></g>`
	chInnerSVG                    = `<mask id="circleFlagsCh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCh0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#eee" d="M389.6 211.5h-89v-89h-89.1v89h-89v89h89v89h89v-89h89z"/></g>`
	chGrInnerSVG                  = `<mask id="circleFlagsChGr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsChGr0)"><path fill="#eee" d="M128 0h128l256 256v256H0V256Z"/><path fill="#338af3" d="M256 0h64l32 32l32-32h64l64 64v64l-32 32l32 32v64h-64l-32-32l-32 32h-64l-64-64v-64l32-32l-32-32Z"/><path fill="#333" d="M0 0h128v256H0z"/><path fill="#ffda44" d="M320 0v64h-64v64h128V0h-64zm64 128v128h64v-64h64v-64H384zm-128 64h64v64h-64zM448 0h64v64h-64z"/><path fill="#333" d="M291 291c-32 0-68 8-82 19c-5 4-7 10-7 16c0 10 7 18 17 20l-25 7a99 99 0 0 0-77-44a9 9 0 0 0-3 18c33 12 45 13 57 37c-21-7-35 24-48 37c-3 1-4 4-4 7c0 5 4 9 9 9l6-2c13-10 26-33 44-28c57 16 66 13 105 7c-31 9-30 21-8 36c-11 17-25 17-41 22c-4 1-7 5-7 9c0 5 4 9 9 9l5-1c19-9 40-18 60-33c-1-8-13-17 8-24c9 8 14 19 41 16c4 7 10 12 7 28l-12 3c-5 0-8 4-8 9s4 9 9 9c3 0 6-2 7-4l18-14c2-15-2-29-8-42c-14-2-17-15-23-29c12 6 22 12 32-3c-14-5-28-13-39-4c-39-24-55-19-80-24c9-8 11-20 8-37c26-7 69-23 96 26c3-41-30-54-66-55z"/></g>`
	ciInnerSVG                    = `<mask id="circleFlagsCi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCi0)"><path fill="#eee" d="M167 0h178l31 253.2L345 512H167l-33.4-257.4z"/><path fill="#ff9811" d="M0 0h167v512H0z"/><path fill="#6da544" d="M345 0h167v512H345z"/></g>`
	ckInnerSVG                    = `<mask id="circleFlagsCk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCk0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="m345 256l4.8 14.6H365l-12.4 9l4.7 14.6l-12.4-9l-12.4 9l4.8-14.6l-12.4-9h15.3zm-63 26l13.7 7l10.9-10.8l-2.4 15.1l13.6 7l-15.1 2.4l-2.4 15.1l-7-13.6l-15.1 2.4l10.8-10.9zm-26 63l14.6-4.7V325l9 12.4l14.6-4.8l-9 12.4l9 12.4l-14.6-4.7l-9 12.4v-15.3zm26 63l7-13.6l-10.8-10.9l15.1 2.4l7-13.6l2.4 15l15.1 2.5l-13.6 7l2.4 15l-10.9-10.8zm63 26l-4.7-14.5H325l12.4-9l-4.8-14.6l12.4 9l12.4-9l-4.7 14.6l12.4 9h-15.3zm63-26l-13.6-7l-10.9 10.9l2.4-15.2l-13.6-7l15-2.3l2.5-15.1l7 13.6l15-2.4l-10.8 10.9zm26-63l-14.5 4.8V365l-9-12.4l-14.6 4.7l9-12.4l-9-12.4l14.6 4.8l9-12.4v15.3zm-26-63l-7 13.7l10.9 10.9l-15.2-2.4l-7 13.6l-2.3-15.1l-15.1-2.4l13.6-7l-2.4-15.1l10.9 10.8z"/></g>`
	clInnerSVG                    = `<mask id="circleFlagsCl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCl0)"><path fill="#d80027" d="m0 256l254.5-51.3L512 256v256H0z"/><path fill="#0052b4" d="M0 0h256l52.7 132.8L256 256H0z"/><path fill="#eee" d="M256 0h256v256H256zM152.4 89l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51L82.2 140h53.6z"/></g>`
	cmInnerSVG                    = `<mask id="circleFlagsCm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCm0)"><path fill="#d80027" d="M144.8 0h222.4l32 260l-32 252H144.8l-32.1-256z"/><path fill="#ffda44" d="m256.1 167l22.1 68h71.5L292 277l22 68l-57.8-42l-57.9 42l22.1-68l-57.8-42H234z"/><path fill="#496e2d" d="M0 0h144.8v512H0z"/><path fill="#ffda44" d="M367.2 0H512v512H367.2z"/></g>`
	cnInnerSVG                    = `<mask id="circleFlagsCn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCn0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m140.1 155.8l22.1 68h71.5l-57.8 42.1l22.1 68l-57.9-42l-57.9 42l22.2-68l-57.9-42.1H118zm163.4 240.7l-16.9-20.8l-25 9.7l14.5-22.5l-16.9-20.9l25.9 6.9l14.6-22.5l1.4 26.8l26 6.9l-25.1 9.6zm33.6-61l8-25.6l-21.9-15.5l26.8-.4l7.9-25.6l8.7 25.4l26.8-.3l-21.5 16l8.6 25.4l-21.9-15.5zm45.3-147.6L370.6 212l19.2 18.7l-26.5-3.8l-11.8 24l-4.6-26.4l-26.6-3.8l23.8-12.5l-4.6-26.5l19.2 18.7zm-78.2-73l-2 26.7l24.9 10.1l-26.1 6.4l-1.9 26.8l-14.1-22.8l-26.1 6.4l17.3-20.5l-14.2-22.7l24.9 10.1z"/></g>`
	cnXjInnerSVG                  = `<mask id="circleFlagsCnXj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCnXj0)"><path fill="#338af3" d="M0 0h512v512H0Z"/><path fill="#eee" d="m312 256l116-38l-72 99V195l72 99zm8 69a128 128 0 1 1 0-137a102 102 0 1 0 0 137z"/></g>`
	coInnerSVG                    = `<mask id="circleFlagsCo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCo0)"><path fill="#d80027" d="m0 384l255.8-29.7L512 384v128H0z"/><path fill="#0052b4" d="m0 256l259.5-31L512 256v128H0z"/><path fill="#ffda44" d="M0 0h512v256H0z"/></g>`
	cpInnerSVG                    = `<mask id="circleFlagsCp0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCp0)"><path fill="#0052b4" d="M0 0h512v512H0Z"/><path fill="#fff" d="m256 52l204 204l-204 204L52 256Z"/><path fill="#ff9811" d="m232 463l12-299h24l12 299z"/><path fill="#6da544" d="M293 172c31-14 42-23 90-22c-40-18-77-23-108 0c9-23 27-54 59-77c-55 9-82 36-86 68c-14-32-55-45-100-45c50 27 36 27 73 63c-32-5-68 9-104 50c49-27 72-18 104-14a90 90 0 0 0-41 86c36-45 31-32 77-72c18 40 40 49 49 86c9-37 0-77-18-91c36 14 68 23 86 50c-9-63-50-72-81-82z"/></g>`
	cqInnerSVG                    = `<mask id="circleFlagsCq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCq0)"><path fill="#eee" d="M304 0h208v208l-32 48l32 48v208H304l-48-32l-48 32H0V304Z"/><path fill="#d80027" d="M0 0v304h208v208h96V304h208v-96H304V0Z"/><path fill="#ffda44" d="M120 27a23 23 0 0 0-23 23v16h-9v-9c0-10-8-18-18-18c-2 0-4 2-4 4v42c0 2 2 4 4 4h27a21 21 0 0 0 21 22v9h-8c-11 0-19 8-19 18c0 3 2 5 5 5h41c3 0 5-2 5-5v-28h57v9h-9c-10 0-18 8-18 18c0 2 2 4 4 4h42c2 0 4-2 4-4V82c9 0 16-7 16-16s-7-16-16-16h-54c-3 0-5-2-5-5s2-6 5-6h54V29h-54c-8 0-16 7-16 16c0 8 8 16 16 16h54c3 0 6 2 6 5s-3 5-6 5h-58c-12 0-22-9-22-21a23 23 0 0 0-22-23zm0 135a23 23 0 0 0-23 23v16h-9v-9c0-10-8-18-18-18c-2 0-4 2-4 4v42c0 2 2 4 4 4h27a21 21 0 0 0 21 22v9h-8c-11 0-19 8-19 18c0 3 2 5 5 5h41c3 0 5-2 5-5v-28h57v9h-9c-10 0-18 8-18 18c0 2 2 4 4 4h42c2 0 4-2 4-4v-55c9 0 16-7 16-16s-7-16-16-16h-54c-3 0-5-2-5-5s2-6 5-6h54v-10h-54c-8 0-16 7-16 16c0 8 8 16 16 16h54c3 0 6 2 6 5s-3 5-6 5h-58c-12 0-22-9-22-21a23 23 0 0 0-22-23z"/></g>`
	crInnerSVG                    = `<mask id="circleFlagsCr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCr0)"><path fill="#0052b4" d="M0 0h512v89l-66.3 167.5L512 423v89H0v-89l69.7-167.3L0 89z"/><path fill="#eee" d="M0 89h512v78l-39.7 91.1L512 345v78H0v-78l36.3-85.6L0 167z"/><path fill="#d80027" d="M0 167h512v178H0z"/></g>`
	cuInnerSVG                    = `<mask id="circleFlagsCu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCu0)"><path fill="#0052b4" d="M0 0h512v102.4L483.8 151l28.2 53.8v102.4l-30.5 50.7l30.5 51.7V512H0l39.8-257z"/><path fill="#ff9811" d="M0 44.5v423V256z"/><path fill="#eee" d="M27.2 102.4v102.4H512V102.4H27.2zm0 204.8v102.4H512V307.2H27.2z"/><path fill="#d80027" d="M0 0v512l256-256L0 0z"/><path fill="#eee" d="m103.6 189.2l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6H87z"/></g>`
	cvInnerSVG                    = `<mask id="circleFlagsCv0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCv0)"><path fill="#0052b4" d="M0 0h512v256.2l-41.9 64.3l41.9 63.7V512H0V384.2L41.3 320L0 256.2z"/><path fill="#eee" d="M0 256.2h512v42.9l-15.7 21.6l15.7 21v42.5H0v-42.5l15.1-21.5L0 299z"/><path fill="#d80027" d="M0 299.1h512v42.6H0z"/><path fill="#ffda44" d="m182.8 190.4l5.2 16.4h17.1l-13.8 10l5.3 16.3l-13.8-10l-14 10l5.4-16.3l-13.9-10h17.1zm0 213.3L188 420h17.1l-13.8 10l5.3 16.2l-13.8-10l-14 10L174 430l-14-10h17.2zm-99.2-72.1l5.2 16.2h17.1L92.1 358l5.2 16.2l-13.7-10l-14 10L75 358l-14-10.1h17.2zm37.9-119.8l5 16h17.2l-13.8 10.3l5.2 16.2l-13.7-10l-14 10l5.4-16.3l-14-10.1H116zm-60.4 67h17l5.5-16.2l5.2 16.2h17.1L92.1 289l5.2 16.4L83.6 295l-14 10.3l5.4-16.4zm46.5 143l5.3-16.2L99 395.4h17.1l5.4-16.2l5.2 16.3h17.1L130 405.6l5.3 16.2l-13.8-10zM282 331.6l-5.4 16.2h-17l13.8 10.2l-5.3 16.2l13.9-10l13.8 10l-5.2-16.3l13.7-10.1h-17zm-38-119.8l-5.3 16.2h-17.1l14 10.2l-5.4 16.2l13.9-10l13.8 10l-5.3-16.3l13.8-10.1h-17zm60.3 67h-17l-5.3-16.2l-5.4 16.2h-17l13.8 10.1l-5.3 16.4L282 295l13.8 10.3l-5.2-16.4zm-46.4 143l-5.3-16.2l13.8-10.2h-17l-5.3-16.2l-5.4 16.3h-17.1l14 10.1l-5.4 16.2l13.9-10z"/></g>`
	cwInnerSVG                    = `<mask id="circleFlagsCw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCw0)"><path fill="#0052b4" d="M0 0h512v342.3l-22 34.2l22 32.5v103H0V409l25.4-31L0 342.2z"/><path fill="#eee" d="m175.2 164.2l13.8 42.5h44.7L197.6 233l13.8 42.5l-36.2-26.3l-36.1 26.3l13.8-42.5l-36.2-26.3h44.7zm-76.7-44.5l8.2 25.5h26.9L111.9 161l8.3 25.5l-21.7-15.7l-21.7 15.7L85 161l-21.7-15.7h26.9z"/><path fill="#ffda44" d="M0 342.3h512V409H0z"/></g>`
	cxInnerSVG                    = `<mask id="circleFlagsCx0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCx0)"><path fill="#6da544" d="m0 0l216.9 301.6L512 512V0z"/><path fill="#0052b4" d="m0 0l512 512H0z"/><circle cx="256" cy="256" r="66.8" fill="#ffda44"/><path fill="#eee" d="m95.3 367.3l7 14.7l16-3.7L111 393l13 10l-16 3.6l.1 16.3l-12.7-10.2L82.6 423v-16.3L66.9 403l12.6-10l-7-14.7l15.8 3.7zM49.3 245l7.1 14.7l15.9-3.7l-7.1 14.6l12.7 10.2l-15.9 3.5v16.3l-12.6-10.2l-12.7 10.2v-16.3l-15.9-3.5l12.8-10.2l-7.1-14.6l15.8 3.7zm46-100.2l7 14.7l16-3.7l-7.2 14.7l12.8 10l-16 3.7l.1 16.2l-12.7-10.1l-12.7 10.1v-16.2l-15.8-3.6l12.7-10.1l-7-14.7l15.8 3.7zm57.1 78l7 14.6l15.9-3.7l-7.1 14.7l12.7 10.1l-15.8 3.6v16.3L152.4 268l-12.7 10.2V262l-15.8-3.6l12.7-10.1l-7-14.7l15.8 3.7zm-40.8 66.7l5.5 17H135L120.5 317l5.6 17l-14.5-10.5L97 333.9l5.6-17l-14.5-10.5H106z"/><path fill="#6da544" d="M256 300.5h22.3s9.6-17 0-33.4l22.2-22.2l-11.1-22.3h-11.1s-5.6 16.7-27.9 16.7s-27.8-16.7-27.8-16.7h-11.1l11.1 22.3l-11.1 22.2l11.1 11.2S233.7 256 256 267c0 0 9.5 14 0 33.4z"/><path fill="#ffda44" d="M422.2 140.2a44.7 44.7 0 0 0-30-50.8a83.1 83.1 0 0 1 7.5 49.9a45.6 45.6 0 0 0-10.4-11.6a45.7 45.7 0 0 0-53.7-2.4a85.3 85.3 0 0 1 35 16.7c8.3 6.6 15 15.5 20.2 24.1l-68 21.3c66.8 11.2 122.4-33.4 122.4-33.4c-6.1-12.3-15.7-14.4-23-13.8z"/></g>`
	cyInnerSVG                    = `<mask id="circleFlagsCy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCy0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#6da544" d="M400.7 222.6h-33.4a111.3 111.3 0 0 1-222.6 0h-33.4c0 66.2 44.5 122 105.2 139.2a37 37 0 0 0 3.9 40.5l36.3-29.2l36.4 29.2a37 37 0 0 0 3.7-40.8a144.8 144.8 0 0 0 103.9-138.9z"/><path fill="#ffda44" d="M167 211.5s0 55.6 55.6 55.6l11.1 11.2H256s11.1-33.4 33.4-33.4c0 0 0-22.3 22.3-22.3H345s-11-44.5 44.6-77.9l-22.3-11.1s-78 55.6-133.6 44.5v22.2h-22.2l-11.2-11l-33.3 22.2z"/></g>`
	czInnerSVG                    = `<mask id="circleFlagsCz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCz0)"><path fill="#eee" d="M0 0h512v256l-265 45.2z"/><path fill="#d80027" d="M210 256h302v256H0z"/><path fill="#0052b4" d="M0 0v512l256-256L0 0z"/></g>`
	deInnerSVG                    = `<mask id="circleFlagsDe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDe0)"><path fill="#ffda44" d="m0 345l256.7-25.5L512 345v167H0z"/><path fill="#d80027" d="m0 167l255-23l257 23v178H0z"/><path fill="#333" d="M0 0h512v167H0z"/></g>`
	dgInnerSVG                    = `<mask id="circleFlagsDg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDg0)"><path fill="#eee" d="M0 0h32l32 32L96 0h160l12 8l13-8h36l24 8l25-8h36l25 8l24-8h36l25 22v34l-8 17l8 17v35l-8 17l8 17v34l-8 17l8 17v35l-8 17l8 17v34l-8 17l8 18v34l-8 17l8 17v35l-8 17l8 17v10h-22l-21-8l-20 8h-44l-21-8l-21 8h-44l-20-8l-21 8h-44l-21-8l-20 8h-44l-21-8l-21 8H64l-21-8l-21 8H0v-10l8-17l-8-17v-35l8-17l-8-17v-34l8-18l-8-17v-34l8-17l-8-17V96l32-32L0 32Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#0052b4" d="M256 0v22c21 0 21 19 43 19c21 0 21-19 42-19c22 0 22 19 43 19s21-19 43-19c21 0 21 19 42 19c22 0 22-19 43-19V0h-25c-4 4-9 7-18 7c-8 0-13-3-18-7h-49c-4 4-9 7-18 7c-8 0-14-3-18-7h-49c-5 4-10 7-18 7c-9 0-14-3-18-7h-25zm0 56v34c21 0 21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20V56c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19zm0 69v3h-83l83 83v16c21 0 21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20v-34c-21 0-21 20-43 20c-21 0-21-20-42-20c-22 0-22 20-43 20s-21-20-43-20c-21 0-21 20-42 20c-22 0-22-20-43-20v-34c21 0 21 19 43 19c21 0 21-19 42-19c22 0 22 19 43 19s21-19 43-19c21 0 21 19 42 19c22 0 22-19 43-19v-34c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19zm-128 48v83h83l-83-83zM0 262v34c21 0 21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20s21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20v-34c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19s-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19zm0 68v35c21 0 21 19 43 19c21 0 21-19 42-19c22 0 22 19 43 19s21-19 43-19c21 0 21 19 42 19c22 0 22-19 43-19s21 19 43 19c21 0 21-19 42-19c22 0 22 19 43 19s21-19 43-19c21 0 21 19 42 19c22 0 22-19 43-19v-35c-21 0-21 20-43 20c-21 0-21-20-42-20c-22 0-22 20-43 20s-21-20-43-20c-21 0-21 20-42 20c-22 0-22-20-43-20s-21 20-43 20c-21 0-21-20-42-20c-22 0-22 20-43 20s-21-20-43-20c-21 0-21 20-42 20c-22 0-22-20-43-20zm0 69v34c21 0 21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20s21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20v-34c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19s-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19zm0 69v34c11 0 16 5 22 10h42c5-5 10-10 21-10s17 5 22 10h42c5-5 11-10 22-10s16 5 22 10h41c6-5 11-10 22-10s16 5 22 10h42c5-5 10-10 21-10s17 5 22 10h42c5-5 11-10 22-10s16 5 22 10h41c6-5 11-10 22-10v-34c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19s-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19z"/><path fill="#6da544" d="M334 134h100l-50 50z"/><path fill="#a2001d" d="M373 173h22v50l-11 8l-11-8Zm0 161h22v61h-22z"/><path fill="#ffda44" d="M451 301a34 34 0 0 0-56-25v-20h11v-22h-11v-11h-22v11h-11v22h11v20a34 34 0 0 0-47 3a33 33 0 0 0 2 46v20h112v-20c7-6 11-14 11-24z"/><path fill="#d80027" d="M351 289c-3 0-6 2-8 4c-4 4-4 12 0 16l8 6v8h65v-8l8-6c4-4 4-12 0-16s-10-5-15-1l-15 14h-21l-15-14a11 11 0 0 0-7-3z"/></g>`
	djInnerSVG                    = `<mask id="circleFlagsDj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDj0)"><path fill="#338af3" d="M0 0h512v256l-153.2 35.7L210 256z"/><path fill="#6da544" d="M210 256h302v256H0z"/><path fill="#eee" d="M0 0v512l256-256z"/><path fill="#d80027" d="m103.6 189.2l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6H87z"/></g>`
	dkInnerSVG                    = `<mask id="circleFlagsDk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDk0)"><path fill="#d80027" d="M0 0h133.6l32.7 20.3l34-20.3H512v222.6L491.4 256l20.6 33.4V512H200.3l-31.7-20.4l-35 20.4H0V289.4l29.4-33L0 222.7z"/><path fill="#eee" d="M133.6 0v222.6H0v66.8h133.6V512h66.7V289.4H512v-66.8H200.3V0h-66.7z"/></g>`
	dmInnerSVG                    = `<mask id="circleFlagsDm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDm0)"><path fill="#496e2d" d="M0 0h208l48 32l48-32h208v208l-32 48l32 48v208H304l-48-32l-48 32H0V304l32-48l-32-48Z"/><path fill="#333" d="M512 240v32l-256 32L0 272v-32l256-32z"/><path fill="#ffda44" d="M512 208v32H0v-32z"/><path fill="#333" d="M240 0h32l32 256l-32 256h-32l-32-256Z"/><path fill="#ffda44" d="M208 0h32v512h-32z"/><path fill="#eee" d="M272 0h32v512h-32z"/><path fill="#eee" d="M512 272v32H0v-32z"/><circle cx="256" cy="256" r="122.4" fill="#d80027"/><path fill="#496e2d" d="M284 270c-9-19-21-37-21-37v-13a14 14 0 0 0-27-2a11 11 0 0 0-4 21a16 16 0 0 1 9-9a14 14 0 0 0 2 2h2s-7 19-7 29c0 27 19 36 19 36l-10 9h19v-18l9 9s17-11 9-27zm-28-114l4 13h14l-11 8l4 12l-11-8l-11 8l4-12l-11-8h14Zm-59 19l11 8l11-8l-4 13l11 8h-14l-4 12l-4-12h-14l11-8zm-36 50h13l4-13l4 13h14l-11 8l4 13l-11-8l-11 8l5-13zm0 62l11-8l-5-13l11 8l11-8l-4 13l11 8h-14l-4 13l-4-13Zm36 50l4-13l-11-8h14l4-12l4 12h13l-10 8l4 13l-11-8zm59 19l-4-13h-14l11-8l-4-12l11 8l11-8l-4 12l11 9h-14Zm59-19l-11-8l-11 8l4-13l-10-8h13l4-12l4 12h14l-11 8zm36-50h-13l-4 13l-4-13h-14l11-8l-4-13l11 8l11-8l-5 13zm0-62l-11 8l5 13l-11-8l-11 8l4-13l-11-8h14l4-13l4 13zm-36-50l-4 13l11 8h-14l-4 12l-4-12h-14l11-8l-4-13l11 8z"/></g>`
	doInnerSVG                    = `<mask id="circleFlagsDo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDo0)"><path fill="#d80027" d="M35.5 256h444l32.5-66.8V0H322.8L256 31.8v444.6L189.2 512H0V322.8z"/><path fill="#0052b4" d="M479.5 256h-444L0 189.2V0h189.2L256 31.8v444.6l66.8 35.6H512V322.8z"/><path fill="#eee" d="M189.2 0h133.6v189.2H512v133.6H322.8V512H189.2V322.8H0V189.2h189.2z"/><path fill="#496e2d" d="M322.8 256a66.8 66.8 0 1 1-133.6 0c0-36.9 66.8-66.8 66.8-66.8s66.8 30 66.8 66.8z"/><path fill="#0052b4" d="M189.2 256a66.8 66.8 0 1 1 133.6 0"/><path fill="#d80027" d="M218.4 222.6v41.7a37.6 37.6 0 1 0 75.2 0v-41.7z"/></g>`
	dzInnerSVG                    = `<mask id="circleFlagsDz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDz0)"><path fill="#496e2d" d="M0 0h256l45.3 251.8L256 512H0z"/><path fill="#eee" d="M256 0h256v512H256z"/><g fill="#d80027"><path d="m311 206.9l-21 29l-34-11l21 28.8l-21 29l34-11.1l21 29v-35.8l34-11.1l-34-11z"/><path d="M277.2 328.3a72.3 72.3 0 1 1 34.5-136a89 89 0 1 0 0 127.3a72 72 0 0 1-34.5 8.7z"/></g></g>`
	eaInnerSVG                    = `<mask id="circleFlagsEa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEa0)"><path fill="#eee" d="m0 0l140.5 27.8L256 0h256l-26.3 132.8L512 256v256l-138.7-30.5L256 512H0l30-139.9L0 256z"/><path fill="#333" d="m0 0l256 256V0H0zm256 256h256V0L256 256zm0 0v256h256L256 256zm0 0H0v256l256-256z"/><path fill="#ffda44" d="m167 178l89 11.2l89-11.1v-53.4l-35.6 17.8L256 89l-53.4 53.5l-35.6-17.8z"/><path fill="#d80027" d="M256 389.6c-49.1 0-89-40-89-89v-89.1l33.3-33.4h111.4l33.3 33.4v89a89 89 0 0 1-89 89z"/><path fill="#eee" d="M256 356.2c-30.7 0-55.7-25-55.7-55.7v-89h111.4v89c0 30.7-25 55.7-55.7 55.7z"/><path fill="#ffda44" d="M167 178h33.3v33.5H167zm144.7 0H345v33.5h-33.3zm0 72.4H345v33.4h-33.3zm-144.7 0h33.3v33.4H167zm89 105.8c-5.8 0-11.4-1-16.7-2.6V388a89 89 0 0 0 33.4 0v-34.4a55.5 55.5 0 0 1-16.7 2.6zm-55-46.6l-26.6 26.6a89.4 89.4 0 0 0 19.4 28l23.6-23.7a55.5 55.5 0 0 1-16.3-31zm136.6 26.6l-26.7-26.6c-2 12-7.8 22.8-16.3 31l23.6 23.5a89.4 89.4 0 0 0 19.4-27.9z"/><path fill="#0052b4" d="M244.9 230v18.6a11.1 11.1 0 0 0 22.2 0V230H245zm0 74.2v18.6a11.1 11.1 0 0 0 22.2 0v-18.6H245zm33.4-37v18.5a11.1 11.1 0 0 0 22.2 0V267h-22.2zm-33.4 0v18.5a11.1 11.1 0 0 0 22.2 0V267H245zm-33.4 0v18.5a11.1 11.1 0 0 0 22.2 0V267h-22.2z"/></g>`
	earthInnerSVG                 = `<mask id="circleFlagsEarth0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEarth0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#eee" d="M302.7 233.7a103.1 103.1 0 0 0 0 206a103.1 103.1 0 0 0 0-206zm0 20c46 0 83 37 83 83s-37 83-83 83s-83-37-83-83s37-83 83-83z"/><path fill="#eee" d="M209.4 72.3a103.1 103.1 0 0 0 0 206a103.1 103.1 0 0 0 0-206zm0 20c46 0 83 37 83 83s-37 83-83 83s-83-37-83-83s37-83 83-83z"/><path fill="#eee" d="M302.7 72.3a103.1 103.1 0 0 0 0 206a103.1 103.1 0 0 0 0-206zm0 20c46 0 83 37 83 83s-37 83-83 83s-83-37-83-83s37-83 83-83z"/><path fill="#eee" d="M349.2 153a103.1 103.1 0 0 0 0 206a103.1 103.1 0 0 0 0-206zm0 20c46 0 83 37 83 83s-37 83-83 83s-83-37-83-83s37-83 83-83z"/><path fill="#eee" d="M209.4 233.7a103.1 103.1 0 0 0 0 206a103.1 103.1 0 0 0 0-206zm0 20c46 0 83 37 83 83s-37 83-83 83s-83-37-83-83s37-83 83-83z"/><path fill="#eee" d="M162.8 153a103.1 103.1 0 0 0 0 206a103.1 103.1 0 0 0 0-206zm0 20c46 0 83 37 83 83s-37 83-83 83s-83-37-83-83s37-83 83-83z"/><path fill="#eee" d="M256 153.1a103.1 103.1 0 0 0 0 206a103.1 103.1 0 0 0 0-206zm0 20c46 0 83 37 83 83c0 45.9-37 83-83 83s-83-37.1-83-83c0-46 37-83 83-83z"/></g>`
	eastAfricanFederationInnerSVG = `<mask id="circleFlagsEastAfricanFederation0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEastAfricanFederation0)"><path fill="#338af3" d="M0 0h512v141l-64 115l64 115v141H0V371l64-115L0 141Z"/><path fill="#eee" d="M0 141h512v23l-256 16L0 164Z"/><path fill="#333" d="M0 164h512v38l-256 16L0 202Z"/><path fill="#6da544" d="m0 220l256 16l256-16v-18H0Z"/><path fill="#eee" d="M0 371h512v-23l-256-16L0 348Z"/><path fill="#ff9811" d="M0 348h512v-38l-256-16L0 310Z"/><path fill="#6da544" d="m0 292l256-16l256 16v18H0Z"/><path fill="#ffda44" d="M0 220h512v72H0z"/><circle cx="256" cy="248" r="64" fill="#eee"/><path fill="#eee" d="m178 301l78-29l78 29v27H178Z"/></g>`
	easterIslandInnerSVG          = `<mask id="circleFlagsEasterIsland0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEasterIsland0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#d80027" d="M334 211.5v22.2h30.5a134.2 134.2 0 0 1-49.5 42a27.7 27.7 0 0 0-47.2 13.2a134.8 134.8 0 0 1-23.6 0a27.8 27.8 0 0 0-47.2-13.2a134.2 134.2 0 0 1-49.5-42h30.6v-22.2h-78a155.8 155.8 0 0 0 311.7 0H334z"/></g>`
	ecInnerSVG                    = `<mask id="circleFlagsEc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEc0)"><path fill="#d80027" d="m0 384l254.7-32.7L512 383.9V512H0z"/><path fill="#0052b4" d="m0 256l255-27l257 27v128H0z"/><path fill="#ffda44" d="M0 0h512v256H0z"/><circle cx="256" cy="256" r="89" fill="#ffda44"/><path fill="#338af3" d="M256 311.6c-30.7 0-55.7-25-55.7-55.6v-33.4a55.7 55.7 0 0 1 111.4 0V256c0 30.6-25 55.6-55.7 55.6z"/><path fill="#333" d="M345 122.4h-66.7a22.3 22.3 0 0 0-44.6 0H167a23 23 0 0 0 23 22.3h-.8c0 12.3 10 22.3 22.3 22.3c0 12.3 10 22.2 22.2 22.2h44.6c12.3 0 22.2-10 22.2-22.2c12.3 0 22.3-10 22.3-22.3h-.8a23 23 0 0 0 23-22.3z"/></g>`
	ecWInnerSVG                   = `<mask id="circleFlagsEcW0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEcW0)"><path fill="#eee" d="m0 167l254.6-36.6L512 166.9v178l-254.6 36.4L0 344.9z"/><path fill="#6da544" d="M0 0h512v166.9H0z"/><path fill="#0052b4" d="M0 344.9h512V512H0z"/></g>`
	eeInnerSVG                    = `<mask id="circleFlagsEe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEe0)"><path fill="#333" d="m0 167l254.6-36.6L512 166.9v178l-254.6 36.4L0 344.9z"/><path fill="#0052b4" d="M0 0h512v166.9H0z"/><path fill="#eee" d="M0 344.9h512V512H0z"/></g>`
	egInnerSVG                    = `<mask id="circleFlagsEg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEg0)"><path fill="#eee" d="m0 144l256-32l256 32v224l-256 32L0 368Z"/><path fill="#d80027" d="M0 0h512v144H0Z"/><path fill="#333" d="M0 368h512v144H0Z"/><path fill="#ff9811" d="M250 191c-8 0-17 4-22 14c5-3 16-1 16 13c0 4-2 8-5 10c-8 0-14-14-29-14c-10 0-19 7-19 17v69l46-7l-14 27h66l-14-27l46 7v-69c0-10-9-17-19-17c-15 0-21 14-29 14c8-23-7-37-23-37z"/></g>`
	ehInnerSVG                    = `<mask id="circleFlagsEh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEh0)"><path fill="#eee" d="m90.1 144.8l210.5-27.7L512 144.8v222.6l-209 32.4l-213-32.4z"/><path fill="#333" d="M0 0h512v144.8H28.2z"/><path fill="#496e2d" d="M39.5 367.4H512V512H0z"/><path fill="#d80027" d="M0 .1V512h.1L256 256.1L0 .1zm365.1 178.1a78 78 0 1 0 18.9 154a78 78 0 0 1 0-152.2a78.2 78.2 0 0 0-18.9-1.8z"/><path fill="#d80027" d="m387.3 206.1l11 34h35.8l-29 21l11.1 34l-29-21l-28.9 21l11.1-34l-29-21h35.8z"/></g>`
	enInnerSVG                    = `<clipPath id="svgIDa"><circle cx="256" cy="256" r="256"/></clipPath><g clip-path="url(#svgIDa)"><path fill="#eee" d="m0 0l8 16l-8 15v16l32 65l-32 64v32l32 48l-32 48v32l32 64l-32 65v47l16-8l15 8h16l65-32l64 32h32l48-32l48 32h32l64-32l65 32h47l-8-15l8-16v-16l-32-65l32-64v-32l-32-48l32-48v-32l-32-64l32-65V0l-15 8l-16-8h-16l-65 32l-64-32h-32l-48 32l-48-32h-32l-64 32L47 0H0z"/><path fill="#0052b4" d="m47 0l129 129V0Zm289 0v129L465 0ZM0 47v129h129Zm512 0L383 176h129ZM0 336v129l129-129Zm383 0l129 129V336Zm-47 47v129h129zm-160 0L47 512h129Z"/><path fill="#d80027" d="M208 0v208H0v96h208v208h96V304h208v-96H304V0h-96z"/><path fill="#d80027" d="m336 336l176 176v-31L367 336Zm0-160L512 0h-31L336 145Zm-160 0L0 0v31l145 145zm0 160L0 512h31l145-145Z"/></g>`
	eoInnerSVG                    = `<clipPath id="svgIDa"><circle cx="256" cy="256" r="256"/></clipPath><g clip-path="url(#svgIDa)"><path fill="#6da544" d="M0 256L256 0h256v512H0z"/><path fill="#eee" d="M0 0h256v256H0z"/><path fill="#6da544" d="m152.4 89l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51L82.2 140h53.6z"/></g>`
	erInnerSVG                    = `<mask id="circleFlagsEr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEr0)"><path fill="#6da544" d="M0 0h512v256H62z"/><path fill="#338af3" d="M62 256h450v256H0z"/><path fill="#d80027" d="M0 0v512l512-256z"/><path fill="#ffda44" d="M133.6 150.3c-49.1 0-89 40-89 89v33.4a89.1 89.1 0 0 0 178 0v-33.4c0-49-40-89-89-89zm55.6 122.4c0 24.9-16.4 46-39 53v-36.3l23.7-23.6l-23.6-23.6v-19.6h-33.4V256l-23.6 23.6l23.6 23.6v22.6a55.7 55.7 0 0 1-39-53.1v-33.4a55.7 55.7 0 0 1 111.3 0z"/></g>`
	esInnerSVG                    = `<mask id="circleFlagsEs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEs0)"><path fill="#ffda44" d="m0 128l256-32l256 32v256l-256 32L0 384Z"/><path fill="#d80027" d="M0 0h512v128H0zm0 384h512v128H0z"/><g fill="#eee"><path d="M144 304h-16v-80h16zm128 0h16v-80h-16z"/><ellipse cx="208" cy="296" rx="48" ry="32"/></g><g fill="#d80027"><rect width="16" height="24" x="128" y="192" rx="8"/><rect width="16" height="24" x="272" y="192" rx="8"/><path d="M208 272v24a24 24 0 0 0 24 24a24 24 0 0 0 24-24v-24h-24z"/></g><rect width="32" height="16" x="120" y="208" fill="#ff9811" ry="8"/><rect width="32" height="16" x="264" y="208" fill="#ff9811" ry="8"/><rect width="32" height="16" x="120" y="304" fill="#ff9811" rx="8"/><rect width="32" height="16" x="264" y="304" fill="#ff9811" rx="8"/><path fill="#ff9811" d="M160 272v24c0 8 4 14 9 19l5-6l5 10a21 21 0 0 0 10 0l5-10l5 6c6-5 9-11 9-19v-24h-9l-5 8l-5-8h-10l-5 8l-5-8z"/><path fill="#d80027" d="M122 248a4 4 0 0 0-4 4a4 4 0 0 0 4 4h172a4 4 0 0 0 4-4a4 4 0 0 0-4-4zm0 24a4 4 0 0 0-4 4a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4a4 4 0 0 0-4-4zm144 0a4 4 0 0 0-4 4a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4a4 4 0 0 0-4-4z"/><path fill="#eee" d="M196 168c-7 0-13 5-15 11l-5-1c-9 0-16 7-16 16s7 16 16 16c7 0 13-4 15-11a16 16 0 0 0 17-4a16 16 0 0 0 17 4a16 16 0 1 0 10-20a16 16 0 0 0-27-5c-3-4-7-6-12-6zm0 8c5 0 8 4 8 8c0 5-3 8-8 8c-4 0-8-3-8-8c0-4 4-8 8-8zm24 0c5 0 8 4 8 8c0 5-3 8-8 8c-4 0-8-3-8-8c0-4 4-8 8-8zm-44 10l4 1l4 8c0 4-4 7-8 7s-8-3-8-8c0-4 4-8 8-8zm64 0c5 0 8 4 8 8c0 5-3 8-8 8c-4 0-8-3-8-7l4-8z"/><path fill="none" d="M220 284v12c0 7 5 12 12 12s12-5 12-12v-12z"/><path fill="#ff9811" d="M200 160h16v32h-16z"/><path fill="#eee" d="M208 224h48v48h-48z"/><path fill="#d80027" d="m248 208l-8 8h-64l-8-8c0-13 18-24 40-24s40 11 40 24zm-88 16h48v48h-48z"/><rect width="20" height="32" x="222" y="232" fill="#d80027" rx="10" ry="10"/><path fill="#ff9811" d="M168 232v8h8v16h-8v8h32v-8h-8v-16h8v-8zm8-16h64v8h-64z"/><g fill="#ffda44"><circle cx="186" cy="202" r="6"/><circle cx="208" cy="202" r="6"/><circle cx="230" cy="202" r="6"/></g><path fill="#d80027" d="M169 272v43a24 24 0 0 0 10 4v-47h-10zm20 0v47a24 24 0 0 0 10-4v-43h-10z"/><g fill="#338af3"><circle cx="208" cy="272" r="16"/><rect width="32" height="16" x="264" y="320" ry="8"/><rect width="32" height="16" x="120" y="320" ry="8"/></g></g>`
	esArInnerSVG                  = `<mask id="circleFlagsEsAr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsAr0)"><path fill="#ffda44" d="M0 0v57l32 29l-32 28v57l32 29l-32 28v57l32 28l-32 28v57l32 29l-32 28v57h512v-57l-32-28l32-29v-57l-32-28l32-28v-57l-32-28l32-29v-57l-32-28l32-29V0H0z"/><path fill="#d80027" d="M0 57h512v57H0Zm0 114h512v57H0Zm0 114h512v56H0Zm0 113h512v57H0Z"/><path fill="#ff9811" d="M96 128v160l96 96c53 0 96-43 96-96l-48-48l48-48v-64h-28v32h-27v-32h-27v32h-28v-32h-27v32h-28v-32z"/><path fill="#0052b4" d="M192 192h96v96h-96z"/><path fill="#eee" d="M192 288v96a96 96 0 0 1-96-96Z"/></g>`
	esCeInnerSVG                  = `<mask id="circleFlagsEsCe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsCe0)"><path fill="#eee" d="m0 0l140.5 27.8L256 0h256l-26.3 132.8L512 256v256l-138.7-30.5L256 512H0l30-139.9L0 256z"/><path fill="#333" d="m0 0l256 256V0H0zm256 256h256V0L256 256zm0 0v256h256L256 256zm0 0H0v256l256-256z"/><path fill="#ffda44" d="m167 178l89 11.2l89-11.1v-53.4l-35.6 17.8L256 89l-53.4 53.5l-35.6-17.8z"/><path fill="#d80027" d="M256 389.6c-49.1 0-89-40-89-89v-89.1l33.3-33.4h111.4l33.3 33.4v89a89 89 0 0 1-89 89z"/><path fill="#eee" d="M256 356.2c-30.7 0-55.7-25-55.7-55.7v-89h111.4v89c0 30.7-25 55.7-55.7 55.7z"/><path fill="#ffda44" d="M167 178h33.3v33.5H167zm144.7 0H345v33.5h-33.3zm0 72.4H345v33.4h-33.3zm-144.7 0h33.3v33.4H167zm89 105.8c-5.8 0-11.4-1-16.7-2.6V388a89 89 0 0 0 33.4 0v-34.4a55.5 55.5 0 0 1-16.7 2.6zm-55-46.6l-26.6 26.6a89.4 89.4 0 0 0 19.4 28l23.6-23.7a55.5 55.5 0 0 1-16.3-31zm136.6 26.6l-26.7-26.6c-2 12-7.8 22.8-16.3 31l23.6 23.5a89.4 89.4 0 0 0 19.4-27.9z"/><path fill="#0052b4" d="M244.9 230v18.6a11.1 11.1 0 0 0 22.2 0V230H245zm0 74.2v18.6a11.1 11.1 0 0 0 22.2 0v-18.6H245zm33.4-37v18.5a11.1 11.1 0 0 0 22.2 0V267h-22.2zm-33.4 0v18.5a11.1 11.1 0 0 0 22.2 0V267H245zm-33.4 0v18.5a11.1 11.1 0 0 0 22.2 0V267h-22.2z"/></g>`
	esCnInnerSVG                  = `<mask id="circleFlagsEsCn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsCn0)"><path fill="#338af3" d="M167 0h178l32.3 257L345 512H167l-25.3-256z"/><path fill="#eee" d="M0 0h166.9v512H0z"/><path fill="#ffda44" d="M344.9 0H512v512H344.9z"/></g>`
	esCtInnerSVG                  = `<mask id="circleFlagsEsCt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsCt0)"><path fill="#ffda44" d="M0 0v57l32 29l-32 28v57l32 29l-32 28v57l32 28l-32 28v57l32 29l-32 28v57h512v-57l-32-28l32-29v-57l-32-28l32-28v-57l-32-28l32-29v-57l-32-28l32-29V0H0z"/><path fill="#d80027" d="M0 57h512v57H0Zm0 114h512v57H0Zm0 114h512v56H0Zm0 113h512v57H0Z"/></g>`
	esGaInnerSVG                  = `<mask id="circleFlagsEsGa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsGa0)"><path fill="#eee" d="M0 63L63 0h449v449l-63 63H0z"/><path fill="#338af3" d="M0 0v63l449 449h63v-63L63 0H0z"/><path fill="#d80027" d="m211 181l46 16.4l44-16.4v-38l-18 9l-27-27l-27 27l-18-9z"/><path fill="#0052b4" d="M186 181v103a70 70 0 0 0 140 0V181z"/></g>`
	esIbInnerSVG                  = `<mask id="circleFlagsEsIb0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsIb0)"><path fill="#ffda44" d="M0 256L256 0h256v56.8l-14.7 29.8l14.7 27.2v56.8L496.2 199l15.8 28.6v56.8l-18.4 29.4l18.4 27.5v56.9L495.3 427l16.7 28.1V512H0v-56.9l24.2-27.4L0 398.2v-56.9l21-27.7l-21-29.2z"/><path fill="#d80027" d="M242 56.8v57h270v-57zm0 113.8v57h270v-57zM0 284.4v56.9h512v-56.9zm0 113.8v56.9h512v-56.9z"/><path fill="#4a1f63" d="M0 0h256v256H0z"/><path fill="#eee" d="M211.5 133.6v22.2h-11.2v-22.2h-22.2v22.2H167v-44.5h-44.6v44.5h-11v-22.2H89v22.2H78v-22.2H55.6v66.7h178v-66.7z"/></g>`
	esMlInnerSVG                  = `<mask id="circleFlagsEsMl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsMl0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#acabb1" d="M122.4 278.3h267.2v33.4H122.4z"/><path fill="#ffda44" d="m167 178l89 11.2l89-11.1v-53.4l-35.6 17.8L256 89l-53.4 53.5l-35.6-17.8z"/><path fill="#eee" d="M144.7 345h-33.4v-61.2l11.1-16.7l-11.1-16.7v-39h33.4zm256 0h-33.4V211.6h33.4v39l-11.1 17l11.1 16.3z"/><path fill="#acabb1" d="M111.3 250.4h289.4v33.4H111.3z"/><path fill="#eee" d="M256 389.6c-49.1 0-89-40-89-89V178h178v122.4a89 89 0 0 1-89 89z"/><path fill="#0052b4" d="M256 356.2a55.7 55.7 0 0 1-55.6-55.7v-89h111.2v89a55.7 55.7 0 0 1-55.6 55.7z"/><path fill="#d80027" d="M167 178h33.4v33.5H167zm144.7 0H345v33.5h-33.4zm0 72.4H345v33.4h-33.4zm-144.7 0h33.3v33.4H167zm72.3-72.3h33.4v33.4h-33.4zM256 356c-5.8 0-11.4-.8-16.7-2.5V388a89.4 89.4 0 0 0 33.4 0v-34.4a55.5 55.5 0 0 1-16.7 2.6zm-54.9-46.5l-26.7 26.6a89.3 89.3 0 0 0 19.4 28l23.6-23.7a55.5 55.5 0 0 1-16.3-31zm136.5 26.6l-26.7-26.6a55.6 55.6 0 0 1-16.3 31l23.6 23.5a89.4 89.4 0 0 0 19.4-27.9z"/></g>`
	esPvInnerSVG                  = `<mask id="circleFlagsEsPv0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsPv0)"><path fill="#d80027" d="M0 47.2L47.2 0h175.4L256 27.3L289.4 0H465l47 47.3v175.3l-24.2 35.2l24.2 31.6v175.4L464.8 512H289.4l-32-26.4l-34.8 26.4H47.2L0 464.8V289.4L25 257L0 222.6z"/><path fill="#496e2d" d="M0 0v47.2L208.8 256L0 464.8V512h47.2L256 303.2L464.8 512H512v-47.2L303.2 256L512 47.3V0h-47L256 208.8L47.2 0H0z"/><path fill="#eee" d="M222.6 0v222.6H0v66.8h222.6V512h66.8V289.4H512v-66.8H289.4V0h-66.8z"/></g>`
	esVariantInnerSVG             = `<mask id="circleFlagsEsVariant0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsVariant0)"><path fill="#d80027" d="M0 0h512v128l-39.8 130.3L512 384v128H0V384l37.8-124L0 128z"/><path fill="#ffda44" d="M0 128h512v256H0z"/></g>`
	etInnerSVG                    = `<mask id="circleFlagsEt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEt0)"><path fill="#d80027" d="m0 345l255.7-31L512 345v167H0z"/><path fill="#ffda44" d="m0 167l258-40.7L512 167v178H0z"/><path fill="#6da544" d="M0 0h512v167H0z"/><circle cx="256" cy="256" r="122.4" fill="#0052b4"/><g fill="#ffda44"><path d="m256 161.2l22 68h71.7l-58 42l22.3 68.3l-58-42.3l-58 42.2l22.2-68.1l-58-42H234z"/><path d="m344.1 273l-70-22.9l43.2-59.6l-18-13L256 237l-43.3-59.7l-18 13.1l43.3 59.7l-70.1 22.7l6.9 21.2l70-22.8V345h22.3v-73.7l70.1 22.8z"/></g></g>`
	etOrInnerSVG                  = `<mask id="circleFlagsEtOr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEtOr0)"><path fill="#eee" d="m0 160l256-32l256 32v192l-256 32L0 352Z"/><path fill="#d80027" d="M0 0h512v160H0Z"/><path fill="#333" d="M0 352h512v160H0Z"/><path fill="#6da544" d="M233 294c0 25-26 42-26 42h98s-26-17-26-42z"/><path fill="#496e2d" d="M256 176c-9 0-17 4-23 10a32 32 0 0 0-48 24a32 32 0 0 0 2 59a32 32 0 0 0 46 25c6 6 14 10 23 10s17-4 23-10c4 3 9 4 14 4c17 0 31-13 32-29a32 32 0 0 0 2-59a32 32 0 0 0-48-24c-6-6-14-10-23-10z"/></g>`
	etTiInnerSVG                  = `<mask id="circleFlagsEtTi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEtTi0)"><path fill="#d80027" d="M0 0h512v512H0l64-256Z"/><path fill="#ffda44" d="M0 0v512l256-256Zm404 168v176L300 202l168 54l-168 54Z"/></g>`
	euInnerSVG                    = `<mask id="circleFlagsEu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEu0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m256 100.2l8.3 25.5H291l-21.7 15.7l8.3 25.6l-21.7-15.8l-21.7 15.8l8.3-25.6l-21.7-15.7h26.8zm-110.2 45.6l24 12.2l18.9-19l-4.2 26.5l23.9 12.2l-26.5 4.2l-4.2 26.5l-12.2-24l-26.5 4.3l19-19zM100.2 256l25.5-8.3V221l15.7 21.7l25.6-8.3l-15.8 21.7l15.8 21.7l-25.6-8.3l-15.7 21.7v-26.8zm45.6 110.2l12.2-24l-19-18.9l26.5 4.2l12.2-23.9l4.2 26.5l26.5 4.2l-24 12.2l4.3 26.5l-19-19zM256 411.8l-8.3-25.5H221l21.7-15.7l-8.3-25.6l21.7 15.8l21.7-15.8l-8.3 25.6l21.7 15.7h-26.8zm110.2-45.6l-24-12.2l-18.9 19l4.2-26.5l-23.9-12.2l26.5-4.2l4.2-26.5l12.2 24l26.5-4.3l-19 19zM411.8 256l-25.5 8.3V291l-15.7-21.7l-25.6 8.3l15.8-21.7l-15.8-21.7l25.6 8.3l15.7-21.7v26.8zm-45.6-110.2l-12.2 24l19 18.9l-26.5-4.2l-12.2 23.9l-4.2-26.5l-26.5-4.2l24-12.2l-4.3-26.5l19 19z"/></g>`
	europeanUnionInnerSVG         = `<mask id="circleFlagsEuropeanUnion0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEuropeanUnion0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m256 100.2l8.3 25.5H291l-21.7 15.7l8.3 25.6l-21.7-15.8l-21.7 15.8l8.3-25.6l-21.7-15.7h26.8zm-110.2 45.6l24 12.2l18.9-19l-4.2 26.5l23.9 12.2l-26.5 4.2l-4.2 26.5l-12.2-24l-26.5 4.3l19-19zM100.2 256l25.5-8.3V221l15.7 21.7l25.6-8.3l-15.8 21.7l15.8 21.7l-25.6-8.3l-15.7 21.7v-26.8zm45.6 110.2l12.2-24l-19-18.9l26.5 4.2l12.2-23.9l4.2 26.5l26.5 4.2l-24 12.2l4.3 26.5l-19-19zM256 411.8l-8.3-25.5H221l21.7-15.7l-8.3-25.6l21.7 15.8l21.7-15.8l-8.3 25.6l21.7 15.7h-26.8zm110.2-45.6l-24-12.2l-18.9 19l4.2-26.5l-23.9-12.2l26.5-4.2l4.2-26.5l12.2 24l26.5-4.3l-19 19zM411.8 256l-25.5 8.3V291l-15.7-21.7l-25.6 8.3l15.8-21.7l-15.8-21.7l25.6 8.3l15.7-21.7v26.8zm-45.6-110.2l-12.2 24l19 18.9l-26.5-4.2l-12.2 23.9l-4.2-26.5l-26.5-4.2l24-12.2l-4.3-26.5l19 19z"/></g>`
	eweInnerSVG                   = `<mask id="circleFlagsEwe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEwe0)"><path fill="#d80027" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#6da544" d="M0 0h512v167H0zm0 345h512v167H0z"/><path fill="#ffda44" d="m110 200l36 112l-95-69h117l-94 69zm146 0l36 112l-94-69h117l-95 69zm146 0l36 112l-94-69h117l-95 69z"/></g>`
	fiInnerSVG                    = `<mask id="circleFlagsFi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFi0)"><path fill="#eee" d="M0 0h133.6l35.3 16.7L200.3 0H512v222.6l-22.6 31.7l22.6 35.1V512H200.3l-32-19.8l-34.7 19.8H0V289.4l22.1-33.3L0 222.6z"/><path fill="#0052b4" d="M133.6 0v222.6H0v66.8h133.6V512h66.7V289.4H512v-66.8H200.3V0h-66.7z"/></g>`
	fjInnerSVG                    = `<mask id="circleFlagsFj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFj0)"><path fill="#338af3" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32l42-16l41 16h45l-8-16l8-15v-14l-16-42l16-41V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#0052b4" d="M128 256v-83l83 83zm128-45l-83-83h83z"/><path fill="#eee" d="m367.3 144.7l-78 22.3h.1v66.7l9.2 11l-9.2 11.3c0 45.5 45.3 67 66.8 74.6l11.5-8.8l10.7 8.8c21.5-7.7 66.8-29.1 66.8-74.6l-8-11l8-11.3V167Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97zm317.2 39v-33.4H289.4V167h66.8v66.7h-66.8V256h66.8v74.6a111 111 0 0 0 11.1 3.4s4.4-1 11.1-3.4V256h66.8v-22.3h-66.8V167Z"/></g>`
	fkInnerSVG                    = `<mask id="circleFlagsFk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFk0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#ffda44" d="M411.8 300.5v11.2h-89v-11.2h-33.4V345h22.3v11.2H423V345h22.2v-44.5z"/><path fill="#338af3" d="M289.4 133.6V256c0 59.6 77.9 78 77.9 78s78-18.4 78-78V133.6h-156z"/><path fill="#eee" d="M367.3 224.9c-19.5 0-19.5 17.8-39 17.8s-19.4-17.8-39-17.8V256c19.6 0 19.6 17.8 39 17.8s19.5-17.8 39-17.8s19.5 17.8 39 17.8s19.4-17.8 39-17.8v-31.2c-19.6 0-19.6 17.8-39 17.8c-19.5 0-19.5-17.8-39-17.8zm0-62.4c-19.5 0-19.5 17.8-39 17.8s-19.4-17.8-39-17.8v31.2c19.6 0 19.6 17.8 39 17.8s19.5-17.8 39-17.8s19.5 17.8 39 17.8s19.4-17.8 39-17.8v-31.2c-19.6 0-19.6 17.8-39 17.8c-19.5 0-19.5-17.8-39-17.8z"/></g>`
	fmInnerSVG                    = `<mask id="circleFlagsFm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFm0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#eee" d="m256 111.3l11 34h35.8l-29 21l11.1 34l-28.9-21l-29 21l11.1-34l-29-21H245zM111.3 256l34-11v-35.8l21 29l34-11.1l-21 28.9l21 29l-34-11.1l-21 29V267zM256 400.7l-11-34h-35.8l29-21l-11.1-34l28.9 21l29-21l-11.1 34l29 21H267zM400.7 256l-34 11v35.8l-21-29l-34 11.1l21-28.9l-21-29l34 11.1l21-29V245z"/></g>`
	foInnerSVG                    = `<mask id="circleFlagsFo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFo0)"><path fill="#eee" d="M0 0h100.2l66.1 53.5L233.7 0H512v189.3L466.3 257l45.7 65.8V512H233.7l-68-50.7l-65.5 50.7H0V322.8l51.4-68.5l-51.4-65z"/><path fill="#0052b4" d="M100.2 0v189.3H0v33.4l24.6 33L0 289.5v33.4h100.2V512h33.4l30.6-26.3l36.1 26.3h33.4V322.8H512v-33.4l-24.6-33.7l24.6-33v-33.4H233.7V0h-33.4l-33.8 25.3L133.6 0z"/><path fill="#d80027" d="M133.6 0v222.7H0v66.7h133.6V512h66.7V289.4H512v-66.7H200.3V0h-66.7z"/></g>`
	frInnerSVG                    = `<mask id="circleFlagsFr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFr0)"><path fill="#eee" d="M167 0h178l25.9 252.3L345 512H167l-29.8-253.4z"/><path fill="#0052b4" d="M0 0h167v512H0z"/><path fill="#d80027" d="M345 0h167v512H345z"/></g>`
	frBreInnerSVG                 = `<mask id="circleFlagsFrBre0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFrBre0)"><path fill="#eee" d="M0 0h256l256 57v57l-32 28l32 29v57l-32 28l32 29v56l-32 29l32 28v57l-256 29L0 455v-57l32-28l-32-29v-56l32-29l-32-28Z"/><path fill="#333" d="M256 0h256v57H256zm0 114h256v57H256zM0 228h512v57H0zm0 113h512v57H0zm0 114h512v57H0zM16 72l19 14l19-14l-19-40z"/><circle cx="46.7" cy="32.4" r="6" fill="#333"/><circle cx="22.7" cy="32.4" r="6" fill="#333"/><circle cx="34.7" cy="20.4" r="6" fill="#333"/><path fill="#333" d="m109 72l19 14l19-14l-19-40z"/><circle cx="140.4" cy="32.4" r="6" fill="#333"/><circle cx="116.4" cy="32.4" r="6" fill="#333"/><circle cx="128.4" cy="20.4" r="6" fill="#333"/><path fill="#333" d="m203 72l19 14l19-14l-19-40z"/><circle cx="234.1" cy="32.4" r="6" fill="#333"/><circle cx="210.1" cy="32.4" r="6" fill="#333"/><circle cx="222.1" cy="20.4" r="6" fill="#333"/><path fill="#333" d="m156 138l19 14l19-14l-19-40z"/><circle cx="187.2" cy="98.2" r="6" fill="#333"/><circle cx="163.2" cy="98.2" r="6" fill="#333"/><circle cx="175.2" cy="86.2" r="6" fill="#333"/><path fill="#333" d="m63 138l19 14l18-14l-18-40z"/><circle cx="93.5" cy="98.2" r="6" fill="#333"/><circle cx="69.5" cy="98.2" r="6" fill="#333"/><circle cx="81.5" cy="86.2" r="6" fill="#333"/><path fill="#333" d="m16 204l19 14l19-14l-19-40z"/><circle cx="46.7" cy="164" r="6" fill="#333"/><circle cx="22.7" cy="164" r="6" fill="#333"/><circle cx="34.7" cy="152" r="6" fill="#333"/><path fill="#333" d="m109 204l19 14l19-14l-19-40z"/><circle cx="140.4" cy="164" r="6" fill="#333"/><circle cx="116.4" cy="164" r="6" fill="#333"/><circle cx="128.4" cy="152" r="6" fill="#333"/><path fill="#333" d="m203 204l19 14l19-14l-19-40z"/><circle cx="234.1" cy="164" r="6" fill="#333"/><circle cx="210.1" cy="164" r="6" fill="#333"/><circle cx="222.1" cy="152" r="6" fill="#333"/></g>`
	frCpInnerSVG                  = `<mask id="circleFlagsFrCp0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFrCp0)"><path fill="#0052b4" d="M0 0h512v512H0Z"/><path fill="#fff" d="m256 52l204 204l-204 204L52 256Z"/><path fill="#ff9811" d="m232 463l12-299h24l12 299z"/><path fill="#6da544" d="M293 172c31-14 42-23 90-22c-40-18-77-23-108 0c9-23 27-54 59-77c-55 9-82 36-86 68c-14-32-55-45-100-45c50 27 36 27 73 63c-32-5-68 9-104 50c49-27 72-18 104-14a90 90 0 0 0-41 86c36-45 31-32 77-72c18 40 40 49 49 86c9-37 0-77-18-91c36 14 68 23 86 50c-9-63-50-72-81-82z"/></g>`
	frTwentyRInnerSVG             = `<mask id="circleFlagsFr20r0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFr20r0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#333" d="M345.1 216.8A98 98 0 0 0 157.5 177l-12 43.8l11 11.2l-22.2 44.5l22.3 11.1l-11.2 11.1V310l11.2 11.1v22.3l11.1 11l33.4-11s0 11 11.1 22.2s22.3 33.4 22.3 33.4s33.4 11.1 55.6-22.2c22.3-33.4 44.6-44.6 44.6-44.6l-21.5-42.9a98 98 0 0 0 32-72.4z"/><path fill="#acabb1" d="M353.1 254.3c45.3-7.6 52.8-96.2 52.8-96.2h-62.3c32.2 78-1.5 83.2-1.5 83.2a8110 8110 0 0 0-173.6-83c-1.3 1.7-5.9 6-11.5 19.8c-5.6 13.7-7 26-7 26a6870 6870 0 0 0 181.2 63c9.2 5.9 20.4 23.7 14.1 76.6l60-16.9s-20.4-54.1-52.2-72.5z"/></g>`
	fxInnerSVG                    = `<mask id="circleFlagsFx0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFx0)"><path fill="#eee" d="M167 0h178l25.9 252.3L345 512H167l-29.8-253.4z"/><path fill="#0052b4" d="M0 0h167v512H0z"/><path fill="#d80027" d="M345 0h167v512H345z"/></g>`
	gaInnerSVG                    = `<mask id="circleFlagsGa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGa0)"><path fill="#ffda44" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#6da544" d="M0 0h512v167H0z"/><path fill="#0052b4" d="M0 345h512v167H0z"/></g>`
	gbInnerSVG                    = `<mask id="circleFlagsGb0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGb0)"><path fill="#eee" d="m0 0l8 22l-8 23v23l32 54l-32 54v32l32 48l-32 48v32l32 54l-32 54v68l22-8l23 8h23l54-32l54 32h32l48-32l48 32h32l54-32l54 32h68l-8-22l8-23v-23l-32-54l32-54v-32l-32-48l32-48v-32l-32-54l32-54V0l-22 8l-23-8h-23l-54 32l-54-32h-32l-48 32l-48-32h-32l-54 32L68 0H0z"/><path fill="#0052b4" d="M336 0v108L444 0Zm176 68L404 176h108zM0 176h108L0 68ZM68 0l108 108V0Zm108 512V404L68 512ZM0 444l108-108H0Zm512-108H404l108 108Zm-68 176L336 404v108z"/><path fill="#d80027" d="M0 0v45l131 131h45L0 0zm208 0v208H0v96h208v208h96V304h208v-96H304V0h-96zm259 0L336 131v45L512 0h-45zM176 336L0 512h45l131-131v-45zm160 0l176 176v-45L381 336h-45z"/></g>`
	gbConInnerSVG                 = `<mask id="circleFlagsGbCon0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGbCon0)"><path fill="#333" d="M0 0h208l48 32l48-32h208v208l-32 48l32 48v208H304l-48-32l-48 32H0V304l32-48l-32-48Z"/><path fill="#eee" d="M208 0v208H0v96h208v208h96V304h208v-96H304V0h-96z"/></g>`
	gbEngInnerSVG                 = `<mask id="circleFlagsGbEng0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGbEng0)"><path fill="#eee" d="M0 0h208l48 32l48-32h208v208l-32 48l32 48v208H304l-48-32l-48 32H0V304l32-48l-32-48Z"/><path fill="#d80027" d="M208 0v208H0v96h208v208h96V304h208v-96H304V0h-96z"/></g>`
	gbNirInnerSVG                 = `<mask id="circleFlagsGbNir0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGbNir0)"><path fill="#eee" d="M0 0h224l32 32l32-32h224v224l-32 32l32 32v224H288l-32-32l-32 32H0V288l32-32l-32-32Z"/><path fill="#d80027" d="M224 0v224H0v64h224v224h64V288h224v-64H288V0Z"/><path fill="#acabb1" d="m256 104l-44 76h-87l43 76l-43 76h87l44 76l44-76h87l-43-76l43-76h-87z"/><path fill="#eee" d="m256 131l-36 62h-73l37 63l-37 63h73l36 63l36-63h73l-37-63l37-63h-73l-36-62z"/><path fill="#d80027" d="M259 195c-5 0-5 5-5 5v39h-3v-33s0-6-5-6c-6 0-6 6-6 6v35l-3 1v-25s0-6-5-6c-6 0-6 6-6 6v43c0 10 5 20 12 26v21h36v-17c9-5 16-14 18-25c0-6 2-11 4-16l4-14s1-5-4-7s-7 3-7 3l-8 19h-3v-47s0-5-5-5c-6 0-6 5-6 5v36h-2v-39s0-5-6-5z"/><path fill="#ffda44" d="M245 38v11h-11v22h11v21c-6-6-14-9-22-9a33 33 0 0 0-23 57v20h112v-20c6-7 10-15 10-24a33 33 0 0 0-55-24V71h11V49h-11V38Z"/></g>`
	gbOrkInnerSVG                 = `<mask id="circleFlagsGbOrk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGbOrk0)"><path fill="#d80027" d="M0 0h100.2l66.1 53.5L233.7 0H512v189.3L466.3 257l45.7 65.8V512H233.7l-68-50.7l-65.5 50.7H0V322.8l51.4-68.5l-51.4-65z"/><path fill="#ffda44" d="M100.2 0v189.3H0v33.4l24.6 33L0 289.5v33.4h100.2V512h33.4l30.6-26.3l36.1 26.3h33.4V322.8H512v-33.4l-24.6-33.7l24.6-33v-33.4H233.7V0h-33.4l-33.8 25.3L133.6 0z"/><path fill="#0052b4" d="M133.6 0v222.7H0v66.7h133.6V512h66.7V289.4H512v-66.7H200.3V0z"/></g>`
	gbSctInnerSVG                 = `<mask id="circleFlagsGbSct0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGbSct0)"><path fill="#0052b4" d="M0 68L68 0h376l68 68v376l-68 68H68L0 444Z"/><path fill="#eee" d="M0 0v68l188 188L0 444v68h68l188-188l188 188h68v-68L324 256L512 68V0h-68L256 188L68 0H0z"/></g>`
	gbWlsInnerSVG                 = `<mask id="circleFlagsGbWls0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGbWls0)"><path fill="#6da544" d="m0 256l256-64l256 64v256H0Z"/><path fill="#eee" d="M0 0h512v256H0Z"/><path fill="#d80027" d="M166 149c-12 0-21 9-21 21h-14c-19 0-35 15-35 34h25c3 8 10 14 17 14h7v48l2 14l-16 5c-1-15-13-26-27-26c-4 0-8 3-8 6v56c0 4 4 7 8 7l12-3l56-15c6 3 14 5 22 5h6v14h-13c-16 0-28 12-28 27c0 4 3 7 7 7h62c4 0 7-3 7-7v-41l9-1h49l9 15h-5c-15 0-27 12-27 27c0 4 3 7 7 7h69c4 0 7-3 7-7c0-6-2-11-5-16l-15-26h23c29 0 52-23 52-52v-3l8 2h6c3-3 4-7 2-10l-28-40l-1-1c-3-2-8-2-10 1l-27 40c-1 2-2 4-1 6c1 4 5 6 9 5l8-3v3c0 10-8 18-18 18h-37c12-12 22-26 28-41l25-57l1-5c0-8-6-14-14-14h-95c-15 0-28 12-28 27v17a11 11 0 1 1-22 0v-31c12-4 21-14 21-27h-41z"/></g>`
	gdInnerSVG                    = `<mask id="circleFlagsGd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGd0)"><path fill="#ffda44" d="M23.6 23.7h464.8v464.7H23.6z"/><path fill="#496e2d" d="M0 75.1L38 38l436.3 436.4l37.7-37.5V75l-39.6-35.5L40.2 471.8L0 436.8z"/><circle cx="256" cy="256.1" r="89" fill="#a2001d"/><path fill="#ffda44" d="m256 167l20 61.5h64.6l-52.3 38l20 61.3l-52.3-38l-52.3 38l20-61.3l-52.3-38H236zM256 .2z"/><path fill="#a2001d" d="M0 0h512v75.1H0zm0 436.8h512v75.1H0z"/><path fill="#ffda44" d="m256 28l4.1 12.7h13.5l-10.9 8l4.2 12.6l-10.9-7.9l-10.8 8l4.1-12.8l-10.9-8H252zm-62 0l4 12.7h13.5l-10.9 8l4.2 12.6l-10.9-7.9l-10.8 8l4.1-12.8l-10.8-8h13.4zm124 0l4.2 12.7h13.4l-10.8 8l4.1 12.6l-10.8-7.9l-10.9 8l4.2-12.8l-10.9-8H314zm-62 422.8l4.1 12.7h13.5l-10.9 8l4.2 12.7l-10.9-8l-10.8 8l4.1-12.8l-10.9-7.9H252zm-62 0l4 12.7h13.5l-10.9 8l4.2 12.7l-10.9-8l-10.8 8l4.1-12.8l-10.8-7.9h13.4zm124 0l4.2 12.7h13.4l-10.8 8l4.1 12.7l-10.8-8l-10.9 8l4.2-12.8l-10.9-7.9H314zm-240-199a22.3 22.3 0 0 1-37.6 23.9c-12-18.8-5-49.5-5-49.5S66 233 78 251.9z"/><circle cx="65.1" cy="273.2" r="11.1" fill="#a2001d"/></g>`
	geInnerSVG                    = `<mask id="circleFlagsGe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGe0)"><path fill="#eee" d="M0 0h222.6l31 23.4L289.4 0H512v222.6l-21.5 31l21.5 35.8V512H289.4l-34.2-20.5l-32.6 20.5H0V289.4l22.7-32.6L0 222.6z"/><path fill="#d80027" d="M222.6 0v222.6H0v66.8h222.6V512h66.8V289.4H512v-66.8H289.4V0z"/><path fill="#d80027" d="M155.8 122.4V89h-33.4v33.4H89v33.4h33.4v33.4h33.4v-33.4h33.4v-33.4zm233.8 0V89h-33.4v33.4h-33.4v33.4h33.4v33.4h33.4v-33.4H423v-33.4zM155.8 356.2v-33.4h-33.4v33.4H89v33.4h33.4V423h33.4v-33.4h33.4v-33.4zm233.8 0v-33.4h-33.4v33.4h-33.4v33.4h33.4V423h33.4v-33.4H423v-33.4z"/></g>`
	geAbInnerSVG                  = `<mask id="circleFlagsGeAb0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGeAb0)"><path fill="#eee" d="M256.6 28.8L512 73v73.2l-30.6 37l30.6 36.1v73.2l-37.5 35.9l37.5 37.2v73.2l-242 36.7L0 438.9v-73.2l28.2-35.3L0 292.6v-73.2z"/><path fill="#6da544" d="m256 0l-56.1 73.1H512V0zm-56.1 146.3l56.1 73.1h256v-73.1zM0 292.6v73.1h512v-73.1zm0 146.3V512h512v-73.1z"/><path fill="#d80027" d="M0 0h256v219.4H0z"/><path fill="#eee" d="M128 196.3L116.9 174v-44.5l27.8-22.3l27.8 22.3v33.4l11.2-11.2V174l-22.3 22.3zM68.1 128l3 9h9.5l-7.7 5.7l3 9l-7.8-5.6l-7.7 5.6l3-9l-7.7-5.6h9.5zm22.3-22.3l3 9.1h9.5l-7.7 5.6l3 9l-7.8-5.5l-7.7 5.6l3-9l-7.8-5.7h9.5zm22.3-22.2l2.9 9h9.5l-7.7 5.7l3 9l-7.7-5.6l-7.8 5.6l3-9l-7.7-5.7h9.5zM221.3 128l-3 9h-9.5l7.7 5.7l-3 9l7.8-5.6l7.7 5.6l-3-9l7.7-5.6h-9.5zM199 105.7l-3 9.1h-9.5l7.7 5.6l-3 9l7.8-5.5l7.7 5.6l-3-9l7.8-5.7h-9.6zm-22.3-22.2l-3 9h-9.4l7.7 5.7l-3 9l7.7-5.6l7.8 5.6l-3-9l7.7-5.7h-9.5zm-32-11.2l-3 9.1h-9.5L140 87l-3 9l7.8-5.5l7.7 5.6l-3-9l7.8-5.7h-9.6z"/></g>`
	gfInnerSVG                    = `<mask id="circleFlagsGf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGf0)"><path fill="#6da544" d="m0 0l216.9 301.6L512 512V0z"/><path fill="#ffda44" d="m0 0l512 512H0z"/><path fill="#d80027" d="m256 121l90 270l-234-168h288L166 391z"/></g>`
	ggInnerSVG                    = `<mask id="circleFlagsGg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGg0)"><path fill="#eee" d="M0 0h222.6l31 23.4L289.4 0H512v222.6l-21.5 31l21.5 35.8V512H289.4l-34.2-20.5l-32.6 20.5H0V289.4l22.7-32.6L0 222.6z"/><path fill="#d80027" d="M222.6 0v222.6H0v66.8h222.6V512h66.8V289.4H512v-66.8H289.4V0z"/><path fill="#ffda44" d="m328.3 267l16.7 11.3v-44.7L328.3 245h-61.2v-61.2l11.2-16.8h-44.6l11.2 16.8v61.1h-61.2L167 233.6v44.7l16.7-11.3h61.2v61.3l-11.2 16.6h44.6L267 328.3V267z"/></g>`
	ghInnerSVG                    = `<mask id="circleFlagsGh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGh0)"><path fill="#ffda44" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#d80027" d="M0 0h512v167H0Z"/><path fill="#496e2d" d="M0 345h512v167H0Z"/><path fill="#333" d="m198 345l151-109H163l151 109l-58-178Z"/></g>`
	giInnerSVG                    = `<mask id="circleFlagsGi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGi0)"><path fill="#d80027" d="m0 345l256.3-31.3L512 345v167H0z"/><path fill="#eee" d="M0 0h512v345H0z"/><path fill="#d80027" d="M356.2 211.5V178h11.1v-22.3H345V167h-22.2v-11.2h-22.3v22.3h11.2v33.4h-22.3v-78h11.1v-22.2h-22.2v11.1H267v-11h-22v11h-11.2v-11h-22.2v22.2h11.1v77.9h-22.3V178h11.2v-22.3h-22.3V167H167v-11.2h-22.3v22.3h11.1v33.4h-22.2v89h244.8v-89z"/><path fill="#ffda44" d="M256 289.4a33.4 33.4 0 0 0-11.1 64.9v46.4h-33.4v44.5H267v-91a33.4 33.4 0 0 0-11.1-64.8zm0 44.5a11.1 11.1 0 1 1 0-22.2a11.1 11.1 0 0 1 0 22.2z"/></g>`
	glInnerSVG                    = `<mask id="circleFlagsGl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGl0)"><path fill="#d80027" d="m0 256l259-45.3L512 256v256H0z"/><path fill="#eee" d="M0 0h512v256H0z"/><path fill="#eee" d="M55.7 256a122.4 122.4 0 1 0 244.8 0l-123-24z"/><path fill="#d80027" d="M55.7 256a122.4 122.4 0 1 1 244.8 0z"/></g>`
	gmInnerSVG                    = `<mask id="circleFlagsGm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGm0)"><path fill="#eee" d="m0 155.8l255-30.6l257 30.6v33.4l-27.7 67.2l27.7 66.4v33.4l-256 32.4L0 356.2v-33.4l28.8-68.5L0 189.2z"/><path fill="#a2001d" d="M0 0h512v155.8H0z"/><path fill="#0052b4" d="M0 189.2h512v133.6H0z"/><path fill="#496e2d" d="M0 356.2h512V512H0z"/></g>`
	gnInnerSVG                    = `<mask id="circleFlagsGn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGn0)"><path fill="#ffda44" d="M167 0h178l25.9 252.3L345 512H167l-29.8-253.4z"/><path fill="#d80027" d="M0 0h167v512H0z"/><path fill="#6da544" d="M345 0h167v512H345z"/></g>`
	gpInnerSVG                    = `<mask id="circleFlagsGp0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGp0)"><path fill="#333" d="M0 176h512v336H0z"/><path fill="#0052b4" d="M0 0h512v176H0z"/><path fill="#6da544" d="M94 228a8 8 0 0 0-2 0a8 8 0 0 0-5 4a8 8 0 0 0-1 6l7 25l-39-23a8 8 0 0 0-11 3a8 8 0 0 0 3 11l39 23l-25 6a8 8 0 0 0-5 4a8 8 0 0 0-1 6a8 8 0 0 0 10 6l41-11l8 5l-26 6a8 8 0 0 0-4 4a8 8 0 0 0-1 6a8 8 0 0 0 10 6l40-11l8 5l-25 6a8 8 0 0 0-4 4a8 8 0 0 0-1 6a8 8 0 0 0 10 6l40-11l25 14a8 8 0 0 0 11-3a8 8 0 0 0-3-10l-25-15l-11-40a8 8 0 0 0-10-6a8 8 0 0 0-5 4a8 8 0 0 0 0 6l6 25l-8-5l-11-40a8 8 0 0 0-9-6a8 8 0 0 0-5 4a8 8 0 0 0-1 6l7 25l-8-5l-11-40a8 8 0 0 0-8-6zm180 96l122 107c17-13 32-30 44-48l-166-59z"/><path fill="#ffda44" d="m256 480l-21-29l-31 19l-8-35l-36 5l6-35l-36-9l19-31l-29-21l29-21l-19-31l36-8l-6-36l36 6l8-36l31 19l21-29l21 29l31-19l8 36l36-6l-6 36l36 8l-19 31l29 21l-29 21l19 31l-36 8l6 36l-36-5l-8 35l-31-19zm0-440l-10 11c-7 11-9 22-5 31l7 17c-4-6-9-10-15-12c-4-1-11 0-16 5c-7 6-10 24 11 27c-2-3-2-9 3-11c3-2 7-1 8 0c2 1 4 4 5 7h-5v10h6c-2 6-6 9-15 16c8 1 17-5 19-7l2-5l-1 20c1 5 2 7 6 11c4-4 5-6 6-11l-1-20l2 4c2 3 11 9 19 9c-9-8-13-11-14-17h5v-10h-5c1-3 3-6 5-7c1-1 5-2 9 0s4 8 2 11c21-3 18-21 11-27c-5-5-11-6-16-5c-6 2-11 6-15 12l7-17c4-9 2-20-5-31l-10-11zm96 0l-10 11c-7 11-9 22-5 31l7 17c-4-6-9-10-15-12c-5-1-11 0-17 5c-6 6-9 24 12 27c-2-3-2-9 3-11c3-2 7-1 8 0c2 1 4 4 5 7h-5v10h5c-1 6-5 9-14 16c8 1 17-5 19-7l2-5l-1 20c1 5 2 7 6 11c4-4 5-6 6-11l-1-20l2 4c2 3 11 9 19 9c-9-8-13-11-15-17h6v-10h-5c1-3 3-6 5-7c1-1 5-2 8 0c5 2 5 8 3 11c21-3 18-21 11-27c-5-5-11-6-16-5c-6 2-11 6-15 12l7-17c4-9 2-20-5-31l-10-11zm-192 0l-10 11c-7 11-9 22-5 31l7 17c-4-6-9-10-15-12c-4-1-11 0-16 5c-7 6-10 24 11 27c-2-3-2-9 3-11c3-2 7-1 8 0c2 1 4 4 5 7h-5v10h6c-2 6-6 9-15 16c8 1 17-5 19-7l2-5l-1 20c1 5 2 7 6 11c4-4 5-6 6-11l-1-20l2 4c2 3 11 9 19 9c-9-8-13-11-14-17h5v-10h-5c1-3 3-6 5-7c1-1 5-2 9 0s4 8 2 11c21-3 19-21 12-27c-6-5-12-6-17-5c-6 2-11 6-15 12l7-17c4-9 2-20-5-31l-10-11z"/></g>`
	gqInnerSVG                    = `<mask id="circleFlagsGq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGq0)"><path fill="#eee" d="M41.3 121.9L512 167v178L43.8 391.3z"/><path fill="#6da544" d="M0 0h512v167H111z"/><path fill="#d80027" d="M111 345h401v167H0z"/><path fill="#0052b4" d="M0 0v512l256-256z"/><path fill="#ff9811" d="M334 257.1h22.2v32.3h-22.3z"/><path fill="#6da544" d="M367.3 245a22.3 22.3 0 1 0-44.5 0a11.1 11.1 0 1 0 0 22.1h44.5a11.1 11.1 0 1 0 0-22.2z"/></g>`
	grInnerSVG                    = `<mask id="circleFlagsGr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGr0)"><path fill="#0052b4" d="M0 0h99l29 32l28-32h356v57l-32 28l32 29v57l-32 28l32 29v57l-32 28l32 28v57l-32 29l32 28v57H0v-57l32-28l-32-29v-56l32-29l-32-28V171l32-29l-32-28Z"/><path fill="#eee" d="M99 0v114H0v57h99v114H0v57h512v-57H156V171h100v-57H156V0Zm157 57v57h256V57Zm0 114v57h256v-57ZM0 398v57h512v-57z"/></g>`
	gsInnerSVG                    = `<mask id="circleFlagsGs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGs0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><rect width="64" height="32" x="336" y="256" fill="#6da544" rx="16" ry="16"/><circle cx="368" cy="96" r="32" fill="#ff9811"/><circle cx="368" cy="144" r="48" fill="#acabb1"/><path fill="#338af3" d="M320 144v77c0 36 48 48 48 48s48-12 48-48v-77z"/><rect width="32" height="128" x="288" y="128" fill="#333" rx="16" ry="16"/><rect width="32" height="128" x="416" y="128" fill="#333" rx="16" ry="16"/><path fill="#6da544" d="m320 144l48 112l48-112z"/><path fill="#ffda44" d="M288 288v32h32v8h96v-8h32v-32h-32v8h-96v-8z"/></g>`
	gtInnerSVG                    = `<mask id="circleFlagsGt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGt0)"><path fill="#338af3" d="M0 0h144.7l108.4 41.2L367.3 0H512v512H367.3l-110.2-41.4L144.7 512H0z"/><path fill="#eee" d="M144.7 0h222.6v512H144.7z"/><path fill="#acabb1" d="M322.9 299.3L279.6 256l41.2-41.2L319 193l-11.7-11.8l-51.2 51.2l-51.2-51.2l-11.7 11.8l-2 21.8l41.3 41.2l-43.3 43.3l23.6 23.6l43.3-43.3l43.3 43.3z"/><path fill="#6da544" d="m319 193l-23.6 23.6a55.5 55.5 0 0 1-39.4 95a55.7 55.7 0 0 1-39.4-95L193 193a89 89 0 1 0 126 0z"/></g>`
	guInnerSVG                    = `<mask id="circleFlagsGu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGu0)"><path fill="#d80027" d="M0 0h512v44.6l-41.2 207.2L512 467.6V512H0v-44.4l38-219l-38-204z"/><path fill="#0052b4" d="M0 44.6h512v423H0z"/><path fill="#d80027" d="M241.5 417.7c-3.5-3-85.7-74.5-85.7-161.6c0-87 82.2-158.6 85.7-161.6L256 82.1l14.5 12.4c3.5 3 85.7 74.4 85.7 161.6c0 87.2-82.2 158.6-85.7 161.6L256 430.1z"/><path fill="#338af3" d="M334 256.1c0-78-78-144.7-78-144.7s-78 66.8-78 144.7a150 150 0 0 0 17.5 66.8h121a150 150 0 0 0 17.4-66.8z"/><path fill="#0052b4" d="M334 256.1a150 150 0 0 1-17.5 66.8L256 334l-60.5-11a150 150 0 0 1-17.4-66.9h155.8z"/><path fill="#ffda44" d="M256 400.8s37.5-32.2 60.5-78h-121c23 45.8 60.5 78 60.5 78z"/><path fill="#6da544" d="M206 211.6h100l-50 50z"/><path fill="#a2001d" d="M239.3 245h33.4v100.1h-33.4z"/></g>`
	guaraniInnerSVG               = `<mask id="circleFlagsGuarani0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGuarani0)"><path fill="#eee" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#d80027" d="M0 0h512v167H0Z"/><path fill="#0052b4" d="M0 345h512v167H0Z"/><path fill="#333" d="M256 80A176 176 0 0 0 80 256a176 176 0 0 0 176 176a176 176 0 0 0 176-176A176 176 0 0 0 256 80Zm0 16a160 160 0 0 1 160 160a160 160 0 0 1-160 160A160 160 0 0 1 96 256A160 160 0 0 1 256 96Zm0 7a32 32 0 0 0-32 32a32 32 0 0 0 20 30v28a64 64 0 0 0-24 10l-20-20a32 32 0 0 0-7-35a32 32 0 0 0-23-10a32 32 0 0 0-22 10a32 32 0 0 0 0 45a32 32 0 0 0 35 7l20 20a64 64 0 0 0-10 24h-29a32 32 0 0 0-29-20a32 32 0 0 0-32 32a32 32 0 0 0 32 32a32 32 0 0 0 29-20h29a64 64 0 0 0 10 24l-20 20a32 32 0 0 0-35 7a32 32 0 0 0 0 45a32 32 0 0 0 45 0a32 32 0 0 0 7-35l20-20a64 64 0 0 0 24 10v28a32 32 0 0 0-20 30a32 32 0 0 0 32 32a32 32 0 0 0 32-32a32 32 0 0 0-20-30v-28a64 64 0 0 0 24-10l20 20a32 32 0 0 0 7 35a32 32 0 0 0 45 0a32 32 0 0 0 0-45a32 32 0 0 0-35-7l-20-20a64 64 0 0 0 10-24h28a32 32 0 0 0 30 20a32 32 0 0 0 32-32a32 32 0 0 0-32-32a32 32 0 0 0-30 20h-28a64 64 0 0 0-10-24l20-20a32 32 0 0 0 35-7a32 32 0 0 0 0-45a32 32 0 0 0-22-10a32 32 0 0 0-23 10a32 32 0 0 0-7 35l-20 20a64 64 0 0 0-24-10v-28a32 32 0 0 0 20-30a32 32 0 0 0-32-32Zm0 105a48 48 0 0 1 48 48a48 48 0 0 1-48 48a48 48 0 0 1-48-48a48 48 0 0 1 48-48zm0 16a32 32 0 0 0-32 32a32 32 0 0 0 32 32a32 32 0 0 0 32-32a32 32 0 0 0-32-32z"/></g>`
	gwInnerSVG                    = `<mask id="circleFlagsGw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGw0)"><path fill="#d80027" d="M0 0h189.2l54 257.6l-54 254.4H0z"/><path fill="#ffda44" d="M189.2 0H512v256l-159 53.5L189.1 256z"/><path fill="#6da544" d="M189.2 256H512v256H189.2z"/><path fill="#333" d="m96.7 189.2l16.6 51H167l-43.4 31.6l16.5 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6h53.7z"/></g>`
	gyInnerSVG                    = `<mask id="circleFlagsGy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGy0)"><path fill="#6da544" d="M77.7 0H512v512H77.8z"/><path fill="#eee" d="M425.4 254.7L31.4 512h46.4L512 256L77.7 0H31.4z"/><path fill="#ffda44" d="M256 256L31.4 512l436.8-256L31.4 0z"/><path fill="#333" d="M0 0v1.8l219.6 253.8L0 510v2h31.4l256-256L31.4 0z"/><path fill="#d80027" d="M0 0v512l256-256L0 0z"/></g>`
	hausaInnerSVG                 = `<mask id="circleFlagsHausa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHausa0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#6da544" d="m218 154l38-84l38 84l-140 140l-84-38l84-38l140 140l-38 84l-38-84l140-140l84 38l-84 38z"/><path fill="#333" d="M244.5 29.5c0 40.5-11.2 78.5-30.7 110.8l-49-49a45.1 45.1 0 0 0-63.7 0l-9.9 9.8a45.1 45.1 0 0 0 0 63.7l49.1 49a214.2 214.2 0 0 1-110.8 30.7v23c40.5 0 78.5 11.2 110.8 30.7l-49 49a45.1 45.1 0 0 0 0 63.7l9.8 9.9a45.1 45.1 0 0 0 63.7 0l49-49.1a214.2 214.2 0 0 1 30.7 110.8h23c0-40.5 11.2-78.5 30.7-110.8l49 49a45.1 45.1 0 0 0 63.7 0l9.9-9.8a45.1 45.1 0 0 0 0-63.7l-49.1-49a214.2 214.2 0 0 1 110.8-30.7v-23c-40.5 0-78.5-11.2-110.8-30.7l49-49a45.1 45.1 0 0 0 0-63.7l-9.8-9.9a45.1 45.1 0 0 0-63.7 0l-49 49.1a214.2 214.2 0 0 1-30.7-110.8h-23zM256 92.2a233.8 233.8 0 0 0 27.7 62.6L256 182.5l-27.7-27.7A233.8 233.8 0 0 0 256 92.2zM133 98a25 25 0 0 1 17.6 7.4l52 51.8a215.9 215.9 0 0 1-45.4 45.3l-51.8-51.9a24.7 24.7 0 0 1 0-35.3l9.9-10A25 25 0 0 1 133 98zm246 0c6.4 0 12.8 2.4 17.7 7.4l10 9.9a24.7 24.7 0 0 1 0 35.3l-52 52a215.9 215.9 0 0 1-45.2-45.3l51.9-52A25 25 0 0 1 379 98zm-162.3 73.5l25.2 25.1l-45.3 45.3l-25.2-25.2a236.7 236.7 0 0 0 45.3-45.3zm78.6 0a236.7 236.7 0 0 0 45.2 45.2l-25.1 25.2l-45.3-45.3l25.2-25.1zM256 210.6l45.3 45.3l-45.3 45.3l-45.3-45.3l45.3-45.3zm-101.2 17.6l27.7 27.7l-27.7 27.7A233.8 233.8 0 0 0 92.2 256a233.8 233.8 0 0 0 62.6-27.7zm202.4 0a233.8 233.8 0 0 0 62.6 27.7a233.8 233.8 0 0 0-62.6 27.7L329.5 256l27.7-27.7zM196.6 270l45.3 45.3l-25.2 25.1a236.7 236.7 0 0 0-45.3-45.2l25.2-25.2zm118.8 0l25.1 25.2a236.7 236.7 0 0 0-45.2 45.3l-25.2-25.2l45.3-45.3zm-158.1 39.4a215.9 215.9 0 0 1 45.2 45.3l-51.9 51.8a24.7 24.7 0 0 1-35.3 0l-10-9.9a24.7 24.7 0 0 1 0-35.3l52-51.9zm197.4 0l52 51.9a24.7 24.7 0 0 1 0 35.3l-10 10a24.7 24.7 0 0 1-35.3 0l-52-52a215.9 215.9 0 0 1 45.4-45.2zm-98.7 20l27.7 27.7a233.8 233.8 0 0 0-27.7 62.6a233.8 233.8 0 0 0-27.7-62.6l27.7-27.7z"/></g>`
	hkInnerSVG                    = `<mask id="circleFlagsHk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHk0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#eee" d="M282.4 193.7c-5.8 24.2-16.1 19.6-21.2 40.7a55.7 55.7 0 0 1 26-108.3c-10.1 42.2.4 46-4.8 67.6zM205 211.6c21.2 13 13.6 21.4 32.1 32.8a55.7 55.7 0 0 1-94.9-58.2c37 22.7 43.8 13.8 62.8 25.4zm-7 79.3c19-16.2 24.7-6.4 41.2-20.4a55.7 55.7 0 0 1-84.7 72.2c33-28.2 26.6-37.4 43.6-51.8zm73.4 31c-9.6-23 1.5-25.3-6.8-45.3a55.7 55.7 0 0 1 42.6 102.8c-16.6-40-27.3-36.9-35.8-57.4zm52.2-60c-24.9 2-23.7-9.3-45.3-7.6a55.7 55.7 0 0 1 111-8.7c-43.3 3.4-43.6 14.5-65.7 16.3z"/></g>`
	hmInnerSVG                    = `<mask id="circleFlagsHm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHm0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#eee" d="m154 300l14 30l32-8l-14 30l25 20l-32 7l1 33l-26-21l-26 21l1-33l-33-7l26-20l-14-30l32 8zm222-27h47l-38 27l15-44l14 44zm7-162l7 15l16-4l-7 15l12 10l-15 3v17l-13-11l-13 11v-17l-15-3l12-10l-7-15l16 4zm57 67l7 15l16-4l-7 15l12 10l-15 3v16l-13-10l-13 11v-17l-15-3l12-10l-7-15l16 4zm-122 22l7 15l16-4l-7 15l12 10l-15 3v16l-13-10l-13 11v-17l-15-3l12-10l-7-15l16 4zm65 156l7 15l16-4l-7 15l12 10l-15 3v17l-13-11l-13 11v-17l-15-3l12-10l-7-15l16 4zM0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/></g>`
	hmongInnerSVG                 = `<mask id="circleFlagsHmong0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHmong0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#ffda44" d="M41 32L16 51L6 72l12 33l-14 17l5 27l12-26l13 18l-34 34l54 15l-19-18l24-17l24 17l-19 18l54-15l-34-34l13-18l12 26l5-27l-14-17l12-33l-10-21l-25-19l17 37l-23 24l3-24l-15-25l-15 25l3 24l-23-24l17-37zm394 0l-25 19l-10 21l12 33l-14 17l5 27l12-26l13 18l-34 34l54 15l-19-18l24-17l24 17l-19 18l54-15l-34-34l13-18l12 26l5-27l-14-17l12-33l-10-21l-25-19l17 37l-23 24l3-24l-15-25l-15 25l3 24l-23-24l17-37zM228.3 42v32l-27.7 16l27.7 16v32l27.7-16l27.7 16v-32l27.7-16l-27.7-16V42L256 58l-27.7-16zM256 244a111 111 0 0 0-111 111a111 111 0 0 0 111 111a111 111 0 0 0 111-111a111 111 0 0 0-111-111zM41 322l-25 19l-10 21l28 69l-34 34l54 15l-19-18l24-17l24 17l-19 18l54-15l-34-34l28-69l-10-21l-25-19l17 37l-23 24l3-24l-15-25l-15 25l3 24l-23-24l17-37zm394 0l-25 19l-10 21l28 69l-34 34l54 15l-19-18l24-17l24 17l-19 18l54-15l-34-34l28-69l-10-21l-25-19l17 37l-23 24l3-24l-15-25l-15 25l3 24l-23-24l17-37z"/></g>`
	hnInnerSVG                    = `<mask id="circleFlagsHn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHn0)"><path fill="#338af3" d="M0 0h512v144.7l-40.5 112.6l40.5 110V512H0V367.3l42.2-114L0 144.7z"/><path fill="#eee" d="M0 144.7h512v222.6H0z"/><path fill="#338af3" d="m157.5 167l8.3 25.5h26.9L171 208.2l8.2 25.5l-21.7-15.7l-21.7 15.7l8.3-25.5l-21.7-15.7h26.9zm0 111.3l8.3 25.5h26.9L171 319.5l8.2 25.5l-21.7-15.7l-21.7 15.7l8.3-25.5l-21.7-15.7h26.9zm197-111.3l8.2 25.5h26.9l-21.7 15.7l8.3 25.5l-21.7-15.7l-21.7 15.7l8.2-25.5l-21.7-15.7h26.9zm0 111.3l8.2 25.5h26.9l-21.7 15.7l8.3 25.5l-21.7-15.7l-21.7 15.7l8.2-25.5l-21.7-15.7h26.9zM256 222.6l8.3 25.5H291L269.4 264l8.3 25.5l-21.7-15.8l-21.7 15.8l8.3-25.5l-21.7-15.8h26.8z"/></g>`
	hrInnerSVG                    = `<mask id="circleFlagsHr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHr0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#d80027" d="M0 0h512v167H0z"/><path fill="#0052b4" d="M0 345h512v167H0z"/><path fill="#338af3" d="M322.8 178h-44.5l7.4-55.7l29.7-22.2l29.6 22.2V167zm-133.6 0h44.5l-7.4-55.7l-29.7-22.2l-29.6 22.2V167z"/><path fill="#0052b4" d="M285.7 178h-59.4v-55.7l29.7-22.2l29.7 22.2z"/><path fill="#eee" d="M167 167v122.3a89 89 0 0 0 35.8 71.3l15.5-3.9l19.7 19.8a89.1 89.1 0 0 0 18 1.8a89 89 0 0 0 17.9-1.8l22.4-18.7l13 2.8a89 89 0 0 0 35.7-71.3V167z"/><path fill="#d80027" d="M167 167h35.6v35.5H167zm71.2 0h35.6v35.5h-35.6zm71.2 0H345v35.5h-35.6zm-106.8 35.5h35.6v35.6h-35.6zm71.2 0h35.6v35.6h-35.6zM167 238.1h35.6v35.6H167zm35.6 35.6h35.6v35.6h-35.6zm35.6-35.6h35.6v35.6h-35.6zm71.2 0H345v35.6h-35.6zm-35.6 35.6h35.6v35.6h-35.6zm-35.6 35.6h35.6V345h-35.6zm-35.6 0h-33.3c3 13.3 9 25.4 17.3 35.6h16zM309.4 345h16a88.8 88.8 0 0 0 17.3-35.6h-33.3zm-106.8 0v15.6a88.7 88.7 0 0 0 35.6 16V345zm71.2 0v31.6a88.7 88.7 0 0 0 35.6-16V345z"/></g>`
	htInnerSVG                    = `<mask id="circleFlagsHt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHt0)"><path fill="#a2001d" d="m0 256l254.8-41.8L512 256v256H0z"/><path fill="#0052b4" d="M0 0h512v256H0z"/><path fill="#eee" d="m345 322.8l-89-11.1l-89 11V189.3h178z"/><circle cx="256" cy="267.1" r="44.5" fill="#0052b4"/><circle cx="256" cy="267.1" r="22.3" fill="#a2001d"/><path fill="#6da544" d="M222.6 211.5h66.8L256 244.9z"/><path fill="#ffda44" d="M244.9 233.7H267v66.8h-22z"/><path fill="#6da544" d="M291.6 293.8h-71.2l-53.4 29h178z"/></g>`
	huInnerSVG                    = `<mask id="circleFlagsHu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsHu0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#d80027" d="M0 0h512v167H0z"/><path fill="#6da544" d="M0 345h512v167H0z"/></g>`
	iaInnerSVG                    = `<clipPath id="svgIDa"><circle cx="256" cy="256" r="256"/></clipPath><g clip-path="url(#svgIDa)"><path fill="#0052b4" d="M0 0h512v256H0z"/><path fill="#d80027" d="M0 256h512v256H0z"/><path fill="#eee" d="m256 0l48 208l208 48l-208 48l-48 208l-48-208L0 256l208-48Z"/></g>`
	icInnerSVG                    = `<mask id="circleFlagsIc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIc0)"><path fill="#338af3" d="M167 0h178l32.3 257L345 512H167l-25.3-256z"/><path fill="#eee" d="M0 0h166.9v512H0z"/><path fill="#ffda44" d="M344.9 0H512v512H344.9z"/></g>`
	idInnerSVG                    = `<mask id="circleFlagsId0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsId0)"><path fill="#eee" d="m0 256l249.6-41.3L512 256v256H0z"/><path fill="#a2001d" d="M0 0h512v256H0z"/></g>`
	idJbInnerSVG                  = `<mask id="circleFlagsIdJb0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIdJb0)"><path fill="#496e2d" d="M0 0h512v512H0Z"/><path fill="#ffda44" d="M229 172a4 4 0 0 0-3 0c-5 5-18 17-21 36c-5 25 9 52 34 67a4 4 0 0 0 5-1a4 4 0 0 0-1-6c-22-14-34-37-30-59c3-16 14-27 19-30a4 4 0 0 0 0-6a4 4 0 0 0-3-1z"/><path fill="#eee" d="M283 172a4 4 0 0 0-3 1a4 4 0 0 0 1 6c4 4 15 14 18 30c4 22-7 45-30 59a4 4 0 0 0-1 6a4 4 0 0 0 5 1c25-15 39-42 34-67c-3-19-16-31-21-35a4 4 0 0 0-3-1zm-28 4s-12 6-12 18s19 24 19 33c0 10-9 13-9 13v22h6v-19s19-9 12-35c-5-18-16-14-16-32z"/><path fill="#0052b4" d="m179 300l12 44l65 32l65-32l12-44H179z"/><path fill="#333" d="m243 267l-16 20l-5-9h-32l-10 22h153l-11-22h-32l-5 9l-16-20h-13z"/><path fill="#ffda44" d="M256 141c-24 0-46 13-62 34s-25 49-25 81c0 13 1 26 4 38l-13 14a104 132 0 0 0 96 80a104 132 0 0 0 96-80l-13-14c3-12 5-25 5-38c0-32-10-60-26-81a79 79 0 0 0-62-34zm0 7c22 0 42 11 57 31a135 135 0 0 1 0 154a71 71 0 0 1-57 32c-22 0-42-12-57-32c-14-19-23-47-23-77s9-57 23-77c15-20 35-31 57-31z"/><rect width="8" height="27" x="252" y="260" fill="#d80027" rx="4" ry="4"/></g>`
	idJtInnerSVG                  = `<mask id="circleFlagsIdJt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIdJt0)"><path fill="#ffda44" d="M0 0h512v512H0Z"/><path fill="#d80027" d="m256 114l-6 2l-95 78c-4 3-4 9-1 13l102-33l102 33c3-4 3-10-1-13l-95-78l-6-2z"/><path fill="#eee" d="M278 231h-1zm-43 0z"/><path fill="#6da544" d="m256 134l-99 72l6 22l51-32l42 25l42-25l51 32l6-22z"/><path fill="#333" d="m256 191l-28 60l-9-21l-10 21l-7-16l-10 21l-7-16l-13 27l14 23h140l14-23l-12-27l-8 16l-10-21l-7 16l-10-21l-9 21zm-62 134zm124 0z"/><path fill="#338af3" d="m183 290l11 32h124l11-32h-73z"/><path fill="#eee" d="M256 129a4 4 0 0 0-2 1l-100 73a4 4 0 0 0-2 4l38 117a4 4 0 0 0 4 3h124a4 4 0 0 0 4-3l38-117a4 4 0 0 0-2-4l-100-73a4 4 0 0 0-2-1zm0 9l95 70l-36 111H197l-36-111l95-70z"/><rect width="9" height="142" x="252" y="170" fill="#ff9811" rx="4.5" ry="4.5"/><rect width="128" height="24" x="192" y="338" fill="#d80027" rx="12" ry="12"/><rect width="160" height="24" x="176" y="350" fill="#eee" rx="12" ry="12"/><path fill="#ffda44" d="m242 191l14-42l14 42l-36-26h44z"/></g>`
	ieInnerSVG                    = `<mask id="circleFlagsIe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIe0)"><path fill="#eee" d="M167 0h178l25.9 252.3L345 512H167l-29.8-253.4z"/><path fill="#6da544" d="M0 0h167v512H0z"/><path fill="#ff9811" d="M345 0h167v512H345z"/></g>`
	ilInnerSVG                    = `<mask id="circleFlagsIl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIl0)"><path fill="#eee" d="M0 0h512v55.7l-25 32.7l25 34v267.2l-26 36l26 30.7V512H0v-55.7l24.8-34.1L0 389.6V122.4l27.2-33.2L0 55.7z"/><path fill="#0052b4" d="M0 55.7v66.7h512V55.7zm0 333.9v66.7h512v-66.7zm352.4-189.3H288l-32-55.6l-32.1 55.6h-64.3l32.1 55.7l-32 55.7h64.2l32.1 55.6l32.1-55.6h64.3L320.3 256l32-55.7zm-57 55.7l-19.7 34.2h-39.4L216.5 256l19.8-34.2h39.4l19.8 34.2zM256 187.6l7.3 12.7h-14.6zm-59.2 34.2h14.7l-7.4 12.7zm0 68.4l7.3-12.7l7.4 12.7zm59.2 34.2l-7.3-12.7h14.6zm59.2-34.2h-14.7l7.4-12.7zm-14.7-68.4h14.7l-7.3 12.7z"/></g>`
	imInnerSVG                    = `<mask id="circleFlagsIm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIm0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#eee" d="m350.8 171.6l-18.1 64.6l-54.3-10l-35-72l-94.4 33.4l-7.4-21l-24.7-3l18.6 52.5l65-16.7l18.4 52l-44.9 66.3l76.3 65l-14.5 17l9.7 22.9l36.1-42.3l-46.8-48l35.8-42l79.8 5.8l18.2-98.6l22 4l15-19.8l-54.8-10zM256 256z"/></g>`
	inInnerSVG                    = `<mask id="circleFlagsIn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIn0)"><path fill="#eee" d="m0 160l256-32l256 32v192l-256 32L0 352z"/><path fill="#ff9811" d="M0 0h512v160H0Z"/><path fill="#6da544" d="M0 352h512v160H0Z"/><circle cx="256" cy="256" r="72" fill="#0052b4"/><circle cx="256" cy="256" r="48" fill="#eee"/><circle cx="256" cy="256" r="24" fill="#0052b4"/></g>`
	inAsInnerSVG                  = `<mask id="circleFlagsInAs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInAs0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#333" d="M256 112c-22 0-40 18-40 40c-7-12-20-20-34-20a40 40 0 1 0 0 80h6v126h-2c-10 0-20 8-20 20v16c0 10 10 20 20 20h2v2c0 12 10 20 20 20h96c12 0 20-8 20-20v-2h2c12 0 20-10 20-20v-16c0-12-8-20-20-20h-2V212h6a40 40 0 1 0 0-80c-14 1-27 9-34 22v-2c0-22-18-40-40-40z"/><rect width="224" height="32" x="144" y="448" fill="#acabb1" rx="16" ry="16"/><rect width="192" height="48" x="160" y="32" fill="#acabb1" rx="24" ry="24"/></g>`
	inGjInnerSVG                  = `<mask id="circleFlagsInGj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInGj0)"><path fill="#eee" d="M256 96L0 128v128l256 32l256-32V128L256 96zm0 256L0 384v128h512V384l-256-32z"/><path fill="#d80027" d="M0 0h512v128H0zm0 256h512v128H0z"/></g>`
	inKaInnerSVG                  = `<mask id="circleFlagsInKa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInKa0)"><path fill="#d80027" d="m0 256l256.5-36.4L512 256v256H0z"/><path fill="#ffda44" d="M0 0h512v256H0z"/></g>`
	inOrInnerSVG                  = `<mask id="circleFlagsInOr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInOr0)"><path fill="#ffda44" d="m0 160l256-32l256 32v192l-256 32L0 352Z"/><path fill="#d80027" d="M0 0h512v160H0Z"/><path fill="#eee" d="M0 352h512v160H0Z"/></g>`
	inTgInnerSVG                  = `<mask id="circleFlagsInTg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInTg0)"><path fill="#eee" d="M0 0h512v128l-256 64L0 128Z"/><path fill="#338af3" d="M0 128h512v128l-256 64L0 256Z"/><path fill="#ff9811" d="M0 256h512v128l-256 64L0 384Z"/><path fill="#496e2d" d="M0 384h512v128H0z"/></g>`
	inTnInnerSVG                  = `<mask id="circleFlagsInTn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInTn0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#496e2d" d="M256 32a206 206 0 1 0 0 412a206 206 0 0 0 0-412zm0 8a198 198 0 0 1 182 276h-43a158 158 0 0 0-73-222v8a151 151 0 0 1 64 214h-1v72a197 197 0 0 1-258 0v-72a151 151 0 0 1 63-214v-8a159 159 0 0 0-73 222H74A198 198 0 0 1 256 40zm0 12A186 186 0 0 0 73 271a8 8 0 0 0 9 7a8 8 0 0 0 7-9l-3-31a170 170 0 1 1 337 31a8 8 0 0 0 7 9a8 8 0 0 0 9-7A187 187 0 0 0 256 52zM63 355a8 8 0 0 0-4 1a8 8 0 0 0-3 11a238 238 0 0 0 400 0a8 8 0 0 0-3-11a8 8 0 0 0-11 3a222 222 0 0 1-372 0a8 8 0 0 0-7-4z"/><path fill="#ffda44" d="M199 86v51l-25 153l-38 31h240l-38-31l-25-153V86H199z"/><path fill="#ff9811" d="M136 321h240v24H136z"/><path fill="#eee" d="M136 345h240v24H136z"/><path fill="#6da544" d="M136 369h240v24H136z"/><circle cx="191" cy="357" r="8" fill="#0052b4"/><circle cx="-321" cy="357" r="8" fill="#0052b4" transform="scale(-1 1)"/><path fill="#d80027" d="M256 252a20 20 0 0 0-20 20a20 20 0 0 0-17-10a20 20 0 0 0-20 20a20 20 0 0 0 20 20a20 20 0 0 0 3 0v71h-1c-5 0-10 4-10 10v8c0 5 5 10 10 10h1v1c0 6 5 10 10 10h48c6 0 10-4 10-10v-1h1c6 0 10-5 10-10v-8c0-6-4-10-10-10h-1v-71a20 20 0 0 0 3 0a20 20 0 0 0 20-20a20 20 0 0 0-20-20a20 20 0 0 0-17 11a20 20 0 0 0 0-1a20 20 0 0 0-20-20z"/></g>`
	ioInnerSVG                    = `<mask id="circleFlagsIo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIo0)"><path fill="#eee" d="M0 0h32l32 32L96 0h160l12 8l13-8h36l24 8l25-8h36l25 8l24-8h36l25 22v34l-8 17l8 17v35l-8 17l8 17v34l-8 17l8 17v35l-8 17l8 17v34l-8 17l8 18v34l-8 17l8 17v35l-8 17l8 17v10h-22l-21-8l-20 8h-44l-21-8l-21 8h-44l-20-8l-21 8h-44l-21-8l-20 8h-44l-21-8l-21 8H64l-21-8l-21 8H0v-10l8-17l-8-17v-35l8-17l-8-17v-34l8-18l-8-17v-34l8-17l-8-17V96l32-32L0 32Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#0052b4" d="M256 0v22c21 0 21 19 43 19c21 0 21-19 42-19c22 0 22 19 43 19s21-19 43-19c21 0 21 19 42 19c22 0 22-19 43-19V0h-25c-4 4-9 7-18 7c-8 0-13-3-18-7h-49c-4 4-9 7-18 7c-8 0-14-3-18-7h-49c-5 4-10 7-18 7c-9 0-14-3-18-7h-25zm0 56v34c21 0 21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20V56c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19zm0 69v3h-83l83 83v16c21 0 21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20v-34c-21 0-21 20-43 20c-21 0-21-20-42-20c-22 0-22 20-43 20s-21-20-43-20c-21 0-21 20-42 20c-22 0-22-20-43-20v-34c21 0 21 19 43 19c21 0 21-19 42-19c22 0 22 19 43 19s21-19 43-19c21 0 21 19 42 19c22 0 22-19 43-19v-34c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19zm-128 48v83h83l-83-83zM0 262v34c21 0 21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20s21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20v-34c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19s-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19zm0 68v35c21 0 21 19 43 19c21 0 21-19 42-19c22 0 22 19 43 19s21-19 43-19c21 0 21 19 42 19c22 0 22-19 43-19s21 19 43 19c21 0 21-19 42-19c22 0 22 19 43 19s21-19 43-19c21 0 21 19 42 19c22 0 22-19 43-19v-35c-21 0-21 20-43 20c-21 0-21-20-42-20c-22 0-22 20-43 20s-21-20-43-20c-21 0-21 20-42 20c-22 0-22-20-43-20s-21 20-43 20c-21 0-21-20-42-20c-22 0-22 20-43 20s-21-20-43-20c-21 0-21 20-42 20c-22 0-22-20-43-20zm0 69v34c21 0 21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20s21 20 43 20c21 0 21-20 42-20c22 0 22 20 43 20s21-20 43-20c21 0 21 20 42 20c22 0 22-20 43-20v-34c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19s-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19zm0 69v34c11 0 16 5 22 10h42c5-5 10-10 21-10s17 5 22 10h42c5-5 11-10 22-10s16 5 22 10h41c6-5 11-10 22-10s16 5 22 10h42c5-5 10-10 21-10s17 5 22 10h42c5-5 11-10 22-10s16 5 22 10h41c6-5 11-10 22-10v-34c-21 0-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19s-21 19-43 19c-21 0-21-19-42-19c-22 0-22 19-43 19s-21-19-43-19c-21 0-21 19-42 19c-22 0-22-19-43-19z"/><path fill="#6da544" d="M334 134h100l-50 50z"/><path fill="#a2001d" d="M373 173h22v50l-11 8l-11-8Zm0 161h22v61h-22z"/><path fill="#ffda44" d="M451 301a34 34 0 0 0-56-25v-20h11v-22h-11v-11h-22v11h-11v22h11v20a34 34 0 0 0-47 3a33 33 0 0 0 2 46v20h112v-20c7-6 11-14 11-24z"/><path fill="#d80027" d="M351 289c-3 0-6 2-8 4c-4 4-4 12 0 16l8 6v8h65v-8l8-6c4-4 4-12 0-16s-10-5-15-1l-15 14h-21l-15-14a11 11 0 0 0-7-3z"/></g>`
	iqInnerSVG                    = `<mask id="circleFlagsIq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIq0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#a2001d" d="M0 0h512v167H0z"/><path fill="#333" d="M0 345h512v167H0z"/><path fill="#496e2d" d="M194.8 239.3h-49.4a22.3 22.3 0 0 1 21.6-16.7v-33.4c-30.7 0-55.7 25-55.7 55.7v27.8h83.5a5.6 5.6 0 0 1 5.5 5.6v11H89v33.5h144.7v-44.5a39 39 0 0 0-39-39zm83.5 50v-100h-33.4v133.5h55.6v-33.4zm111.3 0v-100h-33.4v100H345V256h-33.3v66.8h100.1v-33.4z"/></g>`
	irInnerSVG                    = `<mask id="circleFlagsIr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIr0)"><path fill="#eee" d="M0 144.7L258.8 39.6L512 144.7v222.6L257 493L0 367.3z"/><path fill="#6da544" d="M0 0v144.7h105.6v-22.2h33.6v22.2h33.3v-22.2h33.6v22.2h33.3v-22.2H273v22.2h33v-22.2h33.6v22.2h33.2v-22.2h33.6v22.2H512V0z"/><path fill="#d80027" d="M0 367.3V512h512V367.3H406.4v22.4h-33.6v-22.4h-33.2v22.4H306v-22.4h-33v22.4h-33.6v-22.4h-33.3v22.4h-33.6v-22.4h-33.3v22.4h-33.6v-22.4zm339.1-178h-33.4c.2 3.7.4 7.4.4 11.1c0 24.8-6.2 48.8-17 66c-3.3 5.2-9 12.6-16.4 17.6v-94.7h-33.4v94.8c-7.5-5-13-12.4-16.4-17.7c-10.8-17-17-41-17-65.9c0-3.7.2-7.4.4-11H173a190 190 0 0 0-.4 11c0 68.7 36.7 122.5 83.5 122.5s83.5-53.8 83.5-122.5c0-3.7-.1-7.4-.4-11z"/></g>`
	isInnerSVG                    = `<mask id="circleFlagsIs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIs0)"><path fill="#0052b4" d="M0 0h100.2l66.1 53.5L233.7 0H512v189.3L466.3 257l45.7 65.8V512H233.7l-68-50.7l-65.5 50.7H0V322.8l51.4-68.5l-51.4-65z"/><path fill="#eee" d="M100.2 0v189.3H0v33.4l24.6 33L0 289.5v33.4h100.2V512h33.4l30.6-26.3l36.1 26.3h33.4V322.8H512v-33.4l-24.6-33.7l24.6-33v-33.4H233.7V0h-33.4l-33.8 25.3L133.6 0z"/><path fill="#d80027" d="M133.6 0v222.7H0v66.7h133.6V512h66.7V289.4H512v-66.7H200.3V0z"/></g>`
	itInnerSVG                    = `<mask id="circleFlagsIt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIt0)"><path fill="#eee" d="M167 0h178l25.9 252.3L345 512H167l-29.8-253.4z"/><path fill="#6da544" d="M0 0h167v512H0z"/><path fill="#d80027" d="M345 0h167v512H345z"/></g>`
	itEightyEightInnerSVG         = `<mask id="circleFlagsIt880"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIt880)"><path fill="#eee" d="M0 0h222.6l31 23.4L289.4 0H512v222.6l-21.5 31l21.5 35.8V512H289.4l-34.2-20.5l-32.6 20.5H0V289.4l22.7-32.6L0 222.6z"/><path fill="#d80027" d="M222.6 0v222.6H0v66.8h222.6V512h66.8V289.4H512v-66.8H289.4V0z"/><path fill="#333" d="M378.4 89a44.5 44.5 0 0 0-43 33.4l43 11.2l43.2-11.2A44.5 44.5 0 0 0 378.4 89zM334 189.2l55.7 11.1v-23.6c15.6-4 28-16.4 32-32l-43.2-11.1h-55.6L334 167z"/><path fill="#acabb1" d="M421.6 122.4h-98.8V167h11.1v-22.3h87.7a44.6 44.6 0 0 0 1.4-11.1c0-3.9-.5-7.6-1.4-11.2z"/><path fill="#333" d="M144.7 89a44.5 44.5 0 0 0-43.1 33.4l43.1 11.2l43.1-11.2A44.5 44.5 0 0 0 144.7 89zm-44.5 100.2l55.6 11.1v-23.6c15.7-4 28-16.4 32-32l-43.1-11.1H89l11.2 33.4z"/><path fill="#acabb1" d="M187.8 122.4H89V167h11.2v-22.3h87.6a44.6 44.6 0 0 0 1.4-11.1c0-3.9-.5-7.6-1.4-11.2z"/><path fill="#333" d="M378.4 311.7a44.5 44.5 0 0 0-43 33.3l43 11.2l43.2-11.2a44.5 44.5 0 0 0-43.2-33.3zM334 411.8l55.7 11.2v-23.7c15.6-4 28-16.4 32-32l-43.2-11.1h-55.6l11.1 33.4z"/><path fill="#acabb1" d="M421.6 345h-98.8v44.6h11.1v-22.3h87.7a44.6 44.6 0 0 0 1.4-11.1c0-3.9-.5-7.6-1.4-11.2z"/><path fill="#333" d="M144.7 311.7a44.5 44.5 0 0 0-43.1 33.3l43.1 11.2l43.1-11.2a44.5 44.5 0 0 0-43.1-33.3zm-44.5 100.1l55.6 11.2v-23.7c15.7-4 28-16.4 32-32l-43.1-11.1H89l11.2 33.4z"/><path fill="#acabb1" d="M187.8 345H89v44.6h11.2v-22.3h87.6a44.6 44.6 0 0 0 1.4-11.1c0-3.9-.5-7.6-1.4-11.2z"/></g>`
	itEightyTwoInnerSVG           = `<mask id="circleFlagsIt820"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIt820)"><path fill="#d80027" d="m0 0l216.9 301.6L512 512V0z"/><path fill="#ffda44" d="m0 0l512 512H0z"/><path fill="#ff9811" d="M320.4 320.4A38.9 38.9 0 0 1 272 265l-3.2-1.7a72.5 72.5 0 0 0 30.5-125l-23.3 24a38.9 38.9 0 0 1-23.8 69.5v3.6a72.5 72.5 0 0 0-123.2 36l32.4 8.2a38.9 38.9 0 0 1 72-14l3.2-2a72.5 72.5 0 0 0 92.7 89l-9-32.3zM252.7 254z"/><path fill="#eee" d="M348.2 170.8L330 235.4l-54.3-10l-35-72.1l-94.4 33.5l-7.4-21l-24.7-3l18.6 52.4l65-16.6l18.4 52l-45 66.3l76.4 65l-14.5 17l9.7 22.9l36-42.4l-46.7-48l35.8-42l79.8 5.8l18.2-98.5l22 4l15-19.8zM253.4 255z"/></g>`
	itTwentyThreeInnerSVG         = `<mask id="circleFlagsIt230"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsIt230)"><path fill="#333" d="M0 0h256l64 256l-64 256H0Z"/><path fill="#d80027" d="M256 0h256v512H256Z"/></g>`
	jeInnerSVG                    = `<mask id="circleFlagsJe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsJe0)"><path fill="#eee" d="M0 47.1L47 0h417.8L512 47.2v417.7L464.9 512H47L0 464.9z"/><path fill="#d80027" d="M0 0v47.1L208.8 256L0 464.9V512h47.1L256 303.2L464.9 512H512v-47L303.1 256L512 47.2V0h-47.2L256 208.9L47 0z"/><path fill="#ffda44" d="M211.5 78L256 89l44.5-11V40l-17.8 9L256 22.3L229.3 49l-17.8-9z"/><path fill="#d80027" d="M211.5 78v27.7c0 34.1 44.5 44.6 44.5 44.6s44.5-10.5 44.5-44.6V78z"/></g>`
	jmInnerSVG                    = `<mask id="circleFlagsJm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsJm0)"><path fill="#333" d="M23.3 488.6L0 465V47.1l23.4-23.7l464 464l24.6-22.6V47.1l-24.5-22.7z"/><path fill="#6da544" d="M23.3 23.3L47.1 0h417.7l23.8 23.4l-464 464L47 512h418l22.6-24.5z"/><path fill="#ffda44" d="M0 0v47.1L208.8 256L0 464.9V512h47.1L256 303.2L464.9 512H512v-47L303.1 256L512 47.2V0h-47.2L256 208.9L47 0z"/></g>`
	joInnerSVG                    = `<mask id="circleFlagsJo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsJo0)"><path fill="#eee" d="m126 158l127.8-10.3L512 167v178l-254.9 32.3L126 335.9z"/><path fill="#333" d="M0 0h512v167H107z"/><path fill="#6da544" d="M107 345h405v167H0z"/><path fill="#d80027" d="M0 0v512l256-256z"/><path fill="#eee" d="m101.6 200.3l14 29.4l31.8-7.3l-14.2 29.3l25.5 20.2l-31.8 7.2l.1 32.6l-25.4-20.4l-25.4 20.4V279l-31.7-7.2l25.5-20l-14.2-29.4l31.8 7.3z"/></g>`
	jpInnerSVG                    = `<mask id="circleFlagsJp0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsJp0)"><path fill="#eee" d="M0 0h512v512H0z"/><circle cx="256" cy="256" r="111.3" fill="#d80027"/></g>`
	kanuriInnerSVG                = `<mask id="circleFlagsKanuri0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKanuri0)"><path fill="#ffda44" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#496e2d" d="M0 345h512v167H0Z"/><path fill="#0052b4" d="M0 0h512v167H0Z"/><path fill="#ffda44" d="m373 373l36 112l-94-69h117l-95 69z"/></g>`
	keInnerSVG                    = `<mask id="circleFlagsKe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKe0)"><path fill="#eee" d="m0 144.7l253.4-28.2L512 144.7V178l-36 76.4l36 79.6v33.3L254.7 400L0 367.3V334l39-78l-39-78z"/><path fill="#333" d="M0 0h512v144.7H0z"/><path fill="#a2001d" d="M0 178h512v156H0z"/><path fill="#496e2d" d="M0 367.3h512V512H0z"/><path fill="#eee" d="m335.9 118.3l-30.3-14l-49.4 111.9l-49.4-111.9l-30.3 14L237.8 256l-61.3 137.7l30.3 14l49.4-111.9l49.4 111.9l30.3-14L274.6 256z"/><path fill="#eee" d="M256.2 111.3s-7 5.8-16.7 16l-16 127.6l16 129.8c9.7 10.2 16.7 16 16.7 16s7-5.8 16.7-16l15-130.7l-15-126.7c-9.7-10.2-16.7-16-16.7-16z"/><path fill="#333" d="m311.9 179.2l-10 75.5l10 78.1A158.6 158.6 0 0 0 334 256c0-28.2-9.7-54.5-22.2-76.8zm-111.4 0l8.9 74.3l-9 79.3a158.6 158.6 0 0 1-22.1-76.8c0-28.2 9.7-54.5 22.2-76.8z"/><path fill="#a2001d" d="M200.5 179.2v153.6a269 269 0 0 0 39 51.9V127.3c-11.4 12-26.6 30-39 52zm111.4 0a269 269 0 0 0-39-51.9v257.4c11.4-12 26.5-30 39-52V179.3z"/></g>`
	kgInnerSVG                    = `<mask id="circleFlagsKg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKg0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#ffda44" d="M381.2 256L330 280l27.3 49.6l-55.6-10.6l-7 56.1l-38.7-41.3l-38.7 41.3l-7-56.1l-55.6 10.6l27.3-49.5l-51.2-24.1l51.2-24l-27.3-49.6l55.6 10.6l7-56.1l38.7 41.3l38.7-41.3l7 56.1l55.6-10.6l-27.3 49.5z"/><circle cx="256" cy="256" r="77.9" fill="#d80027"/><path fill="#ffda44" d="M217 256c-1.8 0-3.7.1-5.5.3a44.3 44.3 0 0 0 10.4 28.3a78 78 0 0 1 15-24.9A55.4 55.4 0 0 0 217 256zm24 42a44.4 44.4 0 0 0 30 0c-2.6-10-7.8-19-15-26c-7.2 7-12.4 16-15 26zm53.6-64.3a44.5 44.5 0 0 0-77.2 0a77.4 77.4 0 0 1 38.6 10.5a77.4 77.4 0 0 1 38.6-10.5zm-19.6 26a78 78 0 0 1 15.1 25a44.3 44.3 0 0 0 10.4-28.4a55.8 55.8 0 0 0-5.5-.3a55.3 55.3 0 0 0-20 3.7z"/></g>`
	khInnerSVG                    = `<mask id="circleFlagsKh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKh0)"><path fill="#0052b4" d="M0 0h512v133.7l-39 121.6l39 123.2V512H0V378.5L34.2 255L0 133.7z"/><path fill="#d80027" d="M0 133.7h512v244.8H0z"/><path fill="#eee" d="M345 306.1v-33.3h-22.2v-44.5L300.5 206l-22.2 22.3v-44.5L256 161.5l-22.3 22.3v44.5L211.5 206l-22.3 22.3v44.5H167v33.4h-22.3v33.4h222.6v-33.5z"/></g>`
	kiInnerSVG                    = `<mask id="circleFlagsKi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKi0)"><path fill="#d80027" d="M0 0h512l-.1 234.8l-254.7 260.9L.3 234.8z"/><path fill="#ffda44" d="m345 238l-36.3 17.2l19.3 35.2l-39.5-7.5l-5 39.9l-27.5-29.4l-27.5 29.4l-5-40l-39.5 7.6l19.3-35.2L167 238l36.3-17.1l-19.3-35.3l39.5 7.6l5-40l27.5 29.4l27.5-29.3l5 39.9l39.5-7.6l-19.3 35.3zM322.8 83.6h-50.1a16.7 16.7 0 0 0-33.4 0h-50c0 9.2 8 16.7 17.2 16.7h-.6c0 9.2 7.5 16.7 16.7 16.7c0 9.2 7.5 16.7 16.7 16.7h33.4c9.2 0 16.7-7.5 16.7-16.7c9.2 0 16.7-7.5 16.7-16.7h-.6c9.3 0 17.3-7.5 17.3-16.7z"/><path fill="#eee" d="M85.3 234.8c-21.3 0-21.3 19.5-42.7 19.5c-21.2 0-21.3-19.3-42.3-19.5v34.3l12.3 18.2L.3 303.4V336l12.8 18.2L.3 370.3v32.4l39.1 49.4l441.2 3l31.3-52.4v-32.4l-8.2-16.4l8.2-17.9v-32.6l-10.3-14l10.3-20.3v-34.3c-21.2 0-21.2 19.5-42.5 19.5s-21.4-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.6 19.5s-21.4-19.5-42.7-19.5z"/><path fill="#0052b4" d="M85.3 269.1c-21.3 0-21.3 19.5-42.7 19.5c-21.2 0-21.3-19.3-42.3-19.5v34.3c21 .2 21.1 19.5 42.3 19.5c21.4 0 21.4-19.5 42.7-19.5c21.3 0 21.3 19.5 42.7 19.5c21.4 0 21.3-19.5 42.6-19.5c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.3 0 21.3 19.5 42.6 19.5c21.4 0 21.4-19.5 42.7-19.5c21.3 0 21.3 19.5 42.7 19.5c21.4 0 21.3-19.4 42.5-19.5v-34.3c-21.2 0-21.2 19.5-42.5 19.5s-21.4-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.6 19.5s-21.4-19.5-42.7-19.5zm0 66.9C64 336 64 355.5 42.6 355.5c-21.2 0-21.3-19.3-42.3-19.5v34.3c21 .2 21.1 19.5 42.3 19.5c21.4 0 21.4-19.5 42.7-19.5c21.3 0 21.3 19.5 42.7 19.5c21.4 0 21.3-19.5 42.6-19.5c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.4 0 21.4 19.5 42.7 19.5c21.3 0 21.3-19.5 42.7-19.5c21.3 0 21.3 19.5 42.6 19.5c21.4 0 21.4-19.5 42.7-19.5c21.3 0 21.3 19.5 42.7 19.5c21.4 0 21.3-19.4 42.5-19.5V336c-21.2 0-21.2 19.5-42.5 19.5S448 336 426.7 336c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.6 19.5S106.6 336 85.3 336zm0 66.7c-21.3 0-21.3 19.5-42.7 19.5c-21.2 0-21.3-19.3-42.3-19.5L0 512h512l-.1-109.3c-21.2 0-21.2 19.5-42.5 19.5s-21.4-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.6-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.4 0-21.4 19.5-42.7 19.5c-21.3 0-21.3-19.5-42.7-19.5c-21.3 0-21.3 19.5-42.6 19.5s-21.4-19.5-42.7-19.5z"/></g>`
	kikuyuInnerSVG                = `<mask id="circleFlagsKikuyu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKikuyu0)"><path fill="#eee" d="m228 0l32 211l-32 211h-32l-32-211L196 0Z"/><path fill="#496e2d" d="M196 0v450H0V0Z"/><path fill="#eee" d="m178 87l-31-14l-49 112L49 73L18 87l62 138l-62 138l31 14l49-112l49 112l31-14l-62-138Z"/><path fill="#eee" d="M98 80L81 96L65 224l16 130l17 16l17-16l15-131l-15-127l-17-16Z"/><path fill="#333" d="m154 148l-10 76l10 78a159 159 0 0 0 22-77c0-28-10-54-22-77zm-112 0l9 74l-9 80a159 159 0 0 1-22-77c0-28 10-54 22-77z"/><path fill="#a2001d" d="M42 148v154a269 269 0 0 0 39 52V96c-11 12-26 30-39 52zm112 0a269 269 0 0 0-39-52v258c11-12 26-30 39-52V148Z"/><path fill="#d80027" d="M512 0v422H228V0Z"/><path fill="#333" d="m0 450l86-17l75 8l55-51l41 18l106-56l149 70v90H0Z"/></g>`
	klingonInnerSVG               = `<mask id="circleFlagsKlingon0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKlingon0)"><path fill="#d80027" d="M0 0h512v512H0z"/><circle cx="256" cy="277" r="126" fill="#eee"/><path fill="#333" d="M256 355c15.1-25.3 28.9-39.7 44-65c-3.5-15.8-17.3-7.1-43-204c-29 198.4-44.5 189.4-46 202c14.8 25.8 30.2 41.2 45 67zm-53-51c11.5 17.7 32.5 42.3 44 60c-69.2 4.6-83.4 21.3-107 46c-4.2-81.8 50.6-81.5 63-106zm63 59c13.3-19.2 29.7-41.8 43-61c22.6 46.3 52.6 59.7 80 62c-21.4 30.2-76.8 27.2-123-1z"/></g>`
	kmInnerSVG                    = `<mask id="circleFlagsKm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKm0)"><path fill="#ffda44" d="M124.4 149.7L512 132.8V0H0z"/><path fill="#eee" d="M112.5 132.8H512v122.4l-294.2 16.6z"/><path fill="#d80027" d="M512 255.3H234.5l-136 139.8L512 377.7V256z"/><path fill="#0052b4" d="M108.2 377.7L0 511.2l512 .8V377.7z"/><path fill="#6da544" d="M.8 0L0 511.2l256-256z"/><g fill="#eee"><path d="M67.5 255.3a78 78 0 0 1 61.2-76a78.2 78.2 0 0 0-16.7-2a78 78 0 1 0 16.7 154a78 78 0 0 1-61.2-76z"/><path d="m127.9 188.5l4.1 12.7h13.4l-10.8 8l4.1 12.7l-10.8-8l-10.9 8l4.2-12.8l-10.9-7.9h13.4zm0 33.4l4.1 12.7h13.4l-10.8 8l4.1 12.7l-10.8-8l-10.9 8l4.2-12.8l-10.9-7.9h13.4zm0 33.4L132 268h13.4l-10.8 8l4.1 12.6l-10.8-7.8l-10.9 7.8l4.2-12.7l-10.9-7.9h13.4zm0 33.3l4.1 12.8h13.4l-10.8 8l4.1 12.6l-10.8-7.8L117 322l4.2-12.7l-10.9-7.9h13.4z"/></g></g>`
	knInnerSVG                    = `<mask id="circleFlagsKn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKn0)"><path fill="#ffda44" d="m0 401.9l173.6-225.3L401.9 0H449l63 63l-.1 47.3l-167.3 223.5L110.3 512H63L0 449z"/><path fill="#6da544" d="M0 0v401.9L401.9 0z"/><path fill="#d80027" d="M512 512V110.3L110.3 512z"/><path fill="#333" d="M0 512h63L512 63V0h-63L0 449z"/><path fill="#eee" d="m162.8 302l24 12.2l19-19l-4.4 26.5l24 12.2l-26.6 4.2l-4.2 26.5l-12-24L156 345l19-19zM302 162.8l24 12.1l19-19l-4.3 26.6l24 12.1l-26.6 4.2l-4.2 26.5l-12.2-23.9l-26.5 4.2l19-19z"/></g>`
	kongoInnerSVG                 = `<mask id="circleFlagsKongo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKongo0)"><path fill="#eee" d="m0 160l256-32l256 32v192l-256 32L0 352Z"/><path fill="#d80027" d="M0 0h512v160H0Z"/><path fill="#ffda44" d="M0 352h512v160H0Z"/><path fill="#333" d="M144 423L256 79l112 344L75 211h362z"/><path fill="#eee" d="m175 381l81-250l81 250l-213-154h264z"/></g>`
	kpInnerSVG                    = `<mask id="circleFlagsKp0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKp0)"><path fill="#0052b4" d="M0 0h512v89l-61.2 165.6L512 423v89H0v-89l62.5-174.6L0 89z"/><path fill="#eee" d="M0 89h512v33.4l-28.1 130L512 389.7V423H0v-33.4L30 255L0 122.4z"/><path fill="#d80027" d="M0 122.4h512v267.2H0z"/><circle cx="157.5" cy="256" r="98.5" fill="#eee"/><path fill="#d80027" d="m157.5 157.5l22.1 68h71.6l-58 42l22.3 68.3l-58-42.3l-58 42.2l22.2-68.1l-57.9-42h71.6z"/></g>`
	krInnerSVG                    = `<mask id="circleFlagsKr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKr0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#d80027" d="M345 256c0 22.3-39.8 78-89 78s-89-55.7-89-78a89 89 0 1 1 178 0z"/><path fill="#0052b4" d="M345 256a89 89 0 1 1-178 0"/><path fill="#333" d="m350.4 334.7l23.7-23.6l15.7 15.7l-23.6 23.6zm-39.3 39.4l23.6-23.7l15.7 15.8l-23.6 23.6zm86.6 7.8l23.6-23.6L437 374l-23.6 23.7zm-39.4 39.4l23.6-23.6l15.8 15.7L374 437zm15.8-63l23.6-23.6l15.7 15.7l-23.6 23.7zm-39.4 39.4l23.6-23.6l15.8 15.7l-23.7 23.6zm63-220.4l-63-63l15.8-15.7l63 63zm-63-15.7l-23.6-23.7l15.7-15.7l23.7 23.6zm39.4 39.3l-23.7-23.6l15.8-15.7l23.6 23.6zm7.8-86.6l-23.6-23.6L374 75l23.7 23.6zm39.4 39.4L397.7 130l15.7-15.8L437 138zM90.7 358.3l63 63l-15.8 15.7l-63-63zm63 15.7l23.6 23.7l-15.7 15.7l-23.7-23.6zm-39.4-39.3l23.6 23.6l-15.7 15.8l-23.6-23.7zm23.6-23.6l63 63l-15.7 15.7l-63-63zm15.8-220.4l-63 63L75 137.9l63-63zm23.6 23.6l-63 63l-15.7-15.8l63-63zm23.6 23.6l-63 63l-15.7-15.7l63-63z"/></g>`
	kurdistanInnerSVG             = `<mask id="circleFlagsKurdistan0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKurdistan0)"><path fill="#eee" d="m0 144.7l257-22.4l255 22.4v222.6l-254.9 31L0 367.3z"/><path fill="#d80027" d="M0 0h512v144.7H0z"/><path fill="#6da544" d="M0 367.3h512V512H0z"/><path fill="#ffda44" d="m256 95.9l19.4 83.6l56.1-65l-21.6 83l79.8-31.4l-57.7 63.5l85.3 9.2l-80.7 29.3L408 316l-85-11.5l40.8 75.5l-70-49.7L295 416l-39-76.4l-38.9 76.4l1.2-85.8l-70 49.7l40.8-75.5l-85 11.5l71.3-47.9l-80.7-29.3l85.3-9.3l-57.7-63.4l79.8 31.5l-21.6-83.1l56 65z"/></g>`
	kwInnerSVG                    = `<mask id="circleFlagsKw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKw0)"><path fill="#eee" d="M138.4 147L512 167v178l-373.6 20z"/><path fill="#6da544" d="m0 0l138.4 167H512V0z"/><path fill="#d80027" d="m0 512l138.4-167H512v167z"/><path fill="#333" d="M167 167L0 0v512l167-167z"/></g>`
	kyInnerSVG                    = `<mask id="circleFlagsKy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKy0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><circle cx="367.3" cy="144.7" r="44.5" fill="#6da544"/><path fill="#d80027" d="M289.4 133.5h155.8l.1 29l-76.9 79.7l-79.1-79.7z"/><path fill="#496e2d" d="M445.2 256zm-155.8 0z"/><path fill="#ffda44" d="M289.4 300.5V345h23.4a77.6 77.6 0 0 0 109 0h23.4v-44.5z"/><path fill="#338af3" d="M289.3 193.7v31.2l8 16.6l-7.9 14.5c.1 59.6 77.9 78 77.9 78s78-18.4 78-78l-7.7-13.2l7.7-18v-31.1l-78-10.2z"/><path fill="#eee" d="M367.3 224.9c-19.5 0-19.5 17.8-39 17.8s-19.4-17.8-39-17.8V256c19.7 0 19.6 17.8 39 17.8s19.5-17.8 39-17.8s19.5 17.8 39 17.8s19.4-17.8 39-17.8v-31.2c-19.6 0-19.6 17.8-39 17.8c-19.5 0-19.5-17.8-39-17.8zm0-62.4c-19.5 0-19.5 17.8-39 17.8s-19.4-17.8-39-17.8v31.2c19.7 0 19.6 17.8 39 17.8s19.5-17.8 39-17.8s19.5 17.8 39 17.8s19.4-17.8 39-17.8v-31.2c-19.6 0-19.6 17.8-39 17.8c-19.5 0-19.5-17.8-39-17.8z"/></g>`
	kzInnerSVG                    = `<mask id="circleFlagsKz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKz0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#ffda44" d="M400.7 258.8H111.3c0 20 17.4 36.2 37.4 36.2h-1.2c0 20 16.2 36.1 36.2 36.1c0 20 16.1 36.2 36.1 36.2h72.4c20 0 36.1-16.2 36.1-36.2c20 0 36.2-16.2 36.2-36.1h-1.2c20 0 37.4-16.2 37.4-36.2z"/><path fill="#338af3" d="M356.2 211.5a100.2 100.2 0 0 1-200.4 0"/><path fill="#ffda44" d="m332.5 211.5l-31.3 14.7l16.7 30.3l-34-6.5l-4.3 34.3L256 259l-23.6 25.3L228 250l-34 6.5l16.6-30.3l-31.2-14.7l31.2-14.7l-16.6-30.3l34 6.5l4.3-34.3l23.6 25.2l23.6-25.2L284 173l34-6.5l-16.6 30.3z"/></g>`
	laInnerSVG                    = `<mask id="circleFlagsLa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLa0)"><path fill="#d80027" d="M0 0h512v144.8l-45.8 113L512 367.4V512H0V367.4l46.3-111.1L0 144.8z"/><path fill="#0052b4" d="M0 144.8h512v222.6H0z"/><circle cx="256" cy="256.1" r="89" fill="#eee"/></g>`
	lbInnerSVG                    = `<mask id="circleFlagsLb0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLb0)"><path fill="#d80027" d="M0 0h512v144.8l-45.8 113L512 367.4V512H0V367.4l46.3-111.1L0 144.8z"/><path fill="#eee" d="M0 144.8h512v222.6H0z"/><path fill="#6da544" d="M322.8 300.5L256 178.1l-66.8 122.4h50.1V334h33.4v-33.4z"/></g>`
	lcInnerSVG                    = `<mask id="circleFlagsLc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLc0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#eee" d="M161.4 345h189.2L256 122.4z"/><path fill="#333" d="M194.3 322.8L256 182.4l61.7 140.4z"/><path fill="#ffda44" d="M161.4 345h189.2L256 256z"/></g>`
	liInnerSVG                    = `<mask id="circleFlagsLi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLi0)"><path fill="#d80027" d="m0 256l255.2-39.6L512 256v256H0z"/><path fill="#0052b4" d="M0 0h512v256H0z"/><path fill="#ffda44" d="M189.2 178a33.4 33.4 0 0 0-55.6-24.8v-19.6h11.1v-22.3h-11.1v-11.1h-22.3v11.1h-11.1v22.3h11.1v19.6A33.4 33.4 0 0 0 66.8 203v19.6H178V203c6.8-6.1 11.1-15 11.1-25z"/></g>`
	lkInnerSVG                    = `<mask id="circleFlagsLk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLk0)"><path fill="#ff9811" d="M81.4 27.2h135.2v456.5H81.4z"/><path fill="#6da544" d="m0 44.5l100.2-22v463L0 467.5z"/><path fill="#a2001d" d="m218.6 26.5l293.4 18v93.2l-18.6 106.2l18.6 97.4v126.2l-293.4 18z"/><path fill="#ffda44" d="M0 0v44.5h200.3v423H0V512h512v-44.5H233.7v-423H512V0z"/><path fill="#ffda44" d="M300.5 94.6c-18.4 0-33.4 15-33.4 33.4v144.6a33 33 0 0 0 11.2 25v53h22.2v-27.8h100.2v44.5h-22.3v22.3H423V345l20.4-22.2h48.2a72.9 72.9 0 0 0 20.4 18.5V137.7a66 66 0 0 0-22.5 13.1a67.8 67.8 0 0 0-13.2 82.9c-21.3-.1-41 .2-53.3 0V167c0-11.2-16.7-22.3-16.7-22.3s-16.7 11.1-16.7 22.3h-66.8v66.7h33.4v33.4c0 7.1 1.4 14.1 4.1 20.7L345 300.6h-44.5v-206z"/></g>`
	lrInnerSVG                    = `<mask id="circleFlagsLr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLr0)"><path fill="#eee" d="M0 232.7L256 0h256l-19.2 24L512 46.4v46.7l-19.6 22.6l19.6 23.9v46.5L493.8 209l18.2 23.7v46.6L493 304l19 21.8v46.6l-18 23.5l18 23v46.6l-253.3 21L0 465.5v-46.6l18.8-21.6L0 372.4v-46.6l19.6-21.9L0 279.3z"/><path fill="#d80027" d="m256 0l-22 46.5h278V0zm-17.1 93.2v46.5H512V93.2zm-4.9 93l22 46.5h256v-46.5zM0 279.3v46.5h512v-46.5zm0 93.1v46.5h512v-46.5zm0 93.1V512h512v-46.5z"/><path fill="#0052b4" d="M0 0h256v232.7H0z"/><path fill="#eee" d="m152.4 66.8l16.6 51h53.6l-43.4 31.5l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.5h53.6z"/></g>`
	lsInnerSVG                    = `<mask id="circleFlagsLs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLs0)"><path fill="#eee" d="m0 144.7l255.3-36.5L512 144.7v222.6L250.5 407L0 367.3z"/><path fill="#0052b4" d="M0 0h512v144.7H0z"/><path fill="#6da544" d="M0 367.3h512V512H0z"/><path fill="#333" d="M272.7 250.4v-61.2h-33.4v61.2L199 290.8a66.7 66.7 0 0 0 114 0z"/></g>`
	ltInnerSVG                    = `<mask id="circleFlagsLt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLt0)"><path fill="#6da544" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#ffda44" d="M0 0h512v167H0z"/><path fill="#d80027" d="M0 345h512v167H0z"/></g>`
	luInnerSVG                    = `<mask id="circleFlagsLu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLu0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#d80027" d="M0 0h512v167H0z"/><path fill="#338af3" d="M0 345h512v167H0z"/></g>`
	lvInnerSVG                    = `<mask id="circleFlagsLv0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLv0)"><path fill="#a2001d" d="M0 0h512v189.2l-38.5 70l38.5 63.6V512H0V322.8l39.4-63L0 189.1z"/><path fill="#eee" d="M0 189.2h512v133.6H0z"/></g>`
	lyInnerSVG                    = `<mask id="circleFlagsLy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLy0)"><path fill="#333" d="m0 144.7l257-22.4l255 22.4v222.6l-254.9 31L0 367.3z"/><path fill="#d80027" d="M0 0h512v144.7H0z"/><path fill="#496e2d" d="M0 367.3h512V512H0z"/><g fill="#eee"><path d="m315.6 209.3l21 29l34-11l-21 29l21 28.8l-34-11l-21 29v-36l-34-11l34-11z"/><path d="M258.3 328.4a72.3 72.3 0 1 1 34.4-136a89 89 0 1 0 0 127.3a72 72 0 0 1-34.4 8.7z"/></g></g>`
	maInnerSVG                    = `<mask id="circleFlagsMa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMa0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#496e2d" d="M407.3 210H291.7L256 100.3L220.3 210H104.7l93.5 68l-35.7 109.8L256 320l93.5 68l-35.7-110zm-183 59.5l12.2-37.1h39l12.1 37.1l-31.6 23l-31.6-23zm44-59.4h-24.6l12.3-37.9zm38.3 45.7l-7.7-23.4h39.9zM213 232.4l-7.7 23.4l-32.2-23.4zm-8.3 97.3l12.3-38l20 14.5zm70.1-23.4l20-14.5l12.3 37.9z"/></g>`
	malayaliInnerSVG              = `<mask id="circleFlagsMalayali0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMalayali0)"><path fill="#ffda44" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#d80027" d="M0 0h512v167H0z"/><path fill="#338af3" d="M0 345h512v167H0z"/></g>`
	manipurInnerSVG               = `<mask id="circleFlagsManipur0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsManipur0)"><path fill="#d80027" d="M0 0h512v73L256 96L0 73Z"/><path fill="#eee" d="M0 73h512v73l-256 23L0 146Z"/><path fill="#333" d="M0 146h512v73l-256 23L0 219Z"/><path fill="#ffda44" d="M0 219h512v74l-256 22L0 293Z"/><path fill="#4a1f63" d="M0 293h512v73l-256 23L0 366Z"/><path fill="#338af3" d="M0 366h512v73l-256 23L0 439Z"/><path fill="#6da544" d="M0 439h512v73H0z"/></g>`
	maoriInnerSVG                 = `<mask id="circleFlagsMaori0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMaori0)"><path fill="#eee" d="m0 68.6l247.8-24.7L512 299v93.9l-399.2 36.2L0 162.5z"/><path fill="#333" d="M0 0v68.6h178a115.2 115.2 0 1 1 0 230.4h334V0z"/><path fill="#d80027" d="M0 162.5V512h512V392.9H178a93.9 93.9 0 1 1 0-187.8a47 47 0 0 1 0 93.9a68.3 68.3 0 0 0 0-136.5z"/></g>`
	mcInnerSVG                    = `<mask id="circleFlagsMc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMc0)"><path fill="#eee" d="m0 256l258.2-43.3L512 256v256H0z"/><path fill="#d80027" d="M0 0h512v256H0z"/></g>`
	mdInnerSVG                    = `<mask id="circleFlagsMd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMd0)"><path fill="#0052b4" d="M0 0h144.7l36 254.6l-36 257.4H0z"/><path fill="#d80027" d="M367.3 0H512v512H367.3l-29.7-257.3z"/><path fill="#ffda44" d="M144.7 0h222.6v512H144.7z"/><path fill="#ff9811" d="M345.1 201.4H284a27.8 27.8 0 1 0-55.6 0h-61.2a28.2 28.2 0 0 0 28.3 27.4h-1a27.4 27.4 0 0 0 27.5 27.4c0 13.4 9.6 24.5 22.3 27l-21.6 48.7a88.8 88.8 0 0 0 33.5 6.5a88.8 88.8 0 0 0 33.5-6.5L268.1 283a27.4 27.4 0 0 0 22.3-26.9a27.4 27.4 0 0 0 27.4-27.4h-.9a28.2 28.2 0 0 0 28.3-27.4z"/><path fill="#0052b4" d="M256.1 239.3L220 256v33.4l36.2 22.3l36.2-22.3V256z"/><path fill="#d80027" d="M220 222.6h72.3V256H220z"/></g>`
	meInnerSVG                    = `<mask id="circleFlagsMe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMe0)"><path fill="#ffda44" d="M0 0h512v44.5l-43.6 209.7L512 467.5V512H0v-44.5l46.6-212L0 44.5z"/><path fill="#a2001d" d="M0 44.5h512v423H0z"/><path fill="#ffda44" d="M200.3 189.2h111.4v-44.5l-22.3 11.1l-33.4-33.4l-33.4 33.4l-22.3-11.1zM356.2 256h-55.7a25.5 25.5 0 0 0 3.5-12.8a25.7 25.7 0 0 0-48.3-12.1a25.7 25.7 0 1 0-45 24.9h-54.9c0 17 14.9 30.8 31.9 30.8h-1c0 14 9.2 25.8 22 29.6l-25 25l23.6 23.6l33.5-33.5c1.3.5 2.6.9 4 1.1l-20.2 45.7a83.2 83.2 0 0 0 31.4 6.2a83.2 83.2 0 0 0 31.4-6.2l-20.2-45.7a25.4 25.4 0 0 0 4-1.1l33.5 33.5l23.6-23.6l-25-25a30.8 30.8 0 0 0 22-29.6h-1c17 0 31.9-13.8 31.9-30.8z"/><path fill="#6da544" d="m256 272.7l-36.2 16.7v33.4L256 345l36.2-22.2v-33.4z"/><path fill="#0052b4" d="M219.8 256h72.4v44.5h-72.4z"/></g>`
	mfInnerSVG                    = `<mask id="circleFlagsMf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMf0)"><path fill="#eee" d="M0 128V0h512v128L299 512h-86z"/><circle cx="256" cy="213" r="57" fill="#ffda44"/><path fill="#eee" d="M185 213h142l-71 128Z"/><path fill="#d80027" d="M256 341L142 235h228z"/><path fill="#0052b4" d="m0 128l213 213v171H0Zm512 0L299 341v171h213z"/></g>`
	mgInnerSVG                    = `<mask id="circleFlagsMg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMg0)"><path fill="#eee" d="M0 0h167l45.6 257.6L167.1 512H0z"/><path fill="#d80027" d="M167 0h345v256l-176.7 53.5L166.9 256z"/><path fill="#6da544" d="M167 256h345v256H167z"/></g>`
	mhInnerSVG                    = `<mask id="circleFlagsMh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMh0)"><path fill="#0052b4" d="M0 0h397.6L512 114.2V512H0z"/><path fill="#eee" d="M512 0H397.6L0 512L493.4 53.7z"/><path fill="#ff9811" d="m0 512l512-397.8V0z"/><path fill="#eee" d="M222.2 150.1L191 164.8l16.7 30.3l-34-6.5l-4.3 34.3l-23.6-25.2l-23.7 25.2l-4.3-34.3l-33.9 6.5l16.6-30.3l-31.2-14.7l31.2-14.7L84 105.1l34 6.5l4.2-34.3l23.7 25.3l23.6-25.3l4.3 34.3l34-6.5l-16.7 30.3z"/></g>`
	mizoramInnerSVG               = `<mask id="circleFlagsMizoram0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMizoram0)"><path fill="#eee" d="M0 0h512v256l-256 32L0 256Z"/><path fill="#338af3" d="M0 256h512v256H0Z"/><circle cx="256" cy="256" r="128" fill="#d80027"/></g>`
	mkInnerSVG                    = `<mask id="circleFlagsMk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMk0)"><path fill="#ffda44" d="M0 0h68.2l86 33.6L216.5 0h78.8l61.2 34.8L443.8 0H512v68.3l-31.4 88l31.4 60.2v78.8L477.4 356l34.6 87.8V512h-68.2l-77-37.3l-71.3 37.3h-78.8l-61-39.7L68.2 512H0v-68l40.3-88.2L0 295.5v-78.8L40.3 157L0 68.2z"/><path fill="#d80027" d="M295.5 512h148.3L256 256zm-79-512H68.2L256 256zm.2 512L256 256L68.2 512zM0 216.7L256 256L0 68.2zm0 227.2L256 256L0 295.5zM512 68.3L256 256l256-39.5zM443.8 0H295.3L256 256zM512 443.8V295.3L256 256z"/><circle cx="256" cy="256" r="89" fill="#d80027"/><circle cx="256" cy="256" r="66.8" fill="#ffda44"/></g>`
	mlInnerSVG                    = `<mask id="circleFlagsMl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMl0)"><path fill="#ffda44" d="M167 0h178l25.9 252.3L345 512H167l-29.8-253.4z"/><path fill="#6da544" d="M0 0h167v512H0z"/><path fill="#d80027" d="M345 0h167v512H345z"/></g>`
	mmInnerSVG                    = `<mask id="circleFlagsMm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMm0)"><path fill="#6da544" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#ffda44" d="M0 0h512v167H0z"/><path fill="#d80027" d="M0 345h512v167H0z"/><path fill="#eee" d="M431.5 216.5h-134L256 89l-41.4 127.6h-134l108.3 78.8L147.5 423L256 345l108.5 78l-41.4-127.6z"/></g>`
	mnInnerSVG                    = `<mask id="circleFlagsMn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMn0)"><path fill="#a2001d" d="M0 0h167l84.9 45L345 0h167v512H345l-87.7-48.1L167 512H0z"/><path fill="#0052b4" d="M167 0h178v512H167z"/><g fill="#ffda44"><path d="M122.4 256h22.3v89h-22.3zm-89 0h22.3v89H33.4z"/><circle cx="89" cy="289.4" r="22.3"/><circle cx="89" cy="211.5" r="11.1"/><path d="M66.8 322.8h44.5V345H66.8zm0-89h44.5V256H66.8zM89 133.5l8 24.2h25.4l-20.6 15l7.9 24.3L89 182l-20.6 15l7.9-24.3l-20.6-15h25.5z"/></g></g>`
	moInnerSVG                    = `<mask id="circleFlagsMo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMo0)"><path fill="#496e2d" d="M0 0h512v512H0z"/><path fill="#eee" d="M300.5 245.3c-6.3 0-12.2 1.3-17.8 3.3c3-6.8 5.1-14.3 5.1-22.4c0-31.8-31.8-55.6-31.8-55.6s-31.8 23.8-31.8 55.6c0 8 2 15.6 5 22.4c-5.5-2-11.4-3.3-17.7-3.3c-31.8 0-55.7 31.8-55.7 31.8s23.9 31.8 55.7 31.8C230 309 246 298 256 289c10 9.1 25.9 20 44.5 20c31.8 0 55.7-31.8 55.7-31.8s-23.9-31.8-55.7-31.8z"/><path fill="#ffda44" d="m256 100.2l6.6 20.2h21.2l-17.2 12.5l6.6 20.2l-17.2-12.5l-17.2 12.5l6.6-20.2l-17.2-12.5h21.2zm-80.8 35.3l12.6 11.4l14.7-8.5l-6.9 15.5l12.7 11.4l-17-1.8l-6.9 15.6l-3.5-16.7l-17-1.7l14.8-8.5zM116 194.7l16.6 3.5l8.5-14.7l1.8 17l16.6 3.5l-15.5 6.9l1.8 16.9l-11.4-12.6l-15.5 6.9l8.5-14.8zm220.8-59.2L324.2 147l-14.7-8.5l6.9 15.5l-12.7 11.4l17-1.8l6.9 15.6l3.5-16.7l17-1.7l-14.8-8.5zm59.2 59.2l-16.6 3.5l-8.5-14.7l-1.8 17l-16.6 3.5l15.5 6.9l-1.8 16.9l11.4-12.6l15.5 6.9l-8.5-14.8z"/><path fill="#eee" d="M256 398c25.3 0 48.5-8.5 67.2-22.6H188.8c18.7 14.1 42 22.6 67.2 22.6zm-102-66.8a111.3 111.3 0 0 0 13.2 22.6h177.6a111.3 111.3 0 0 0 13.2-22.6z"/></g>`
	mpInnerSVG                    = `<mask id="circleFlagsMp0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMp0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#eee" d="M248 417.2a16.7 16.7 0 0 1-15-11a16.7 16.7 0 0 1-29.7-7.6a16.7 16.7 0 0 1-27.6-13.5a16.7 16.7 0 0 1-24.3-18.9a16.7 16.7 0 0 1-17.8-5a16.7 16.7 0 0 1-2.3-18.3a16.7 16.7 0 0 1-15-26.8a16.7 16.7 0 0 1-9.2-29.4a16.7 16.7 0 0 1-3.1-30.5a16.7 16.7 0 0 1 3-30.6a16.7 16.7 0 0 1 9.2-29.4a16.7 16.7 0 0 1 15-26.9a16.7 16.7 0 0 1 2.2-18.3c4.5-5.2 11.5-7 17.7-5a16.7 16.7 0 0 1 6-17.5a16.7 16.7 0 0 1 18.3-1.4a16.7 16.7 0 0 1 27.6-13.6a16.7 16.7 0 0 1 29.7-7.7a16.7 16.7 0 0 1 30.7-1.6a16.7 16.7 0 0 1 30.4 4.6a16.7 16.7 0 0 1 28.9 10.6a16.7 16.7 0 0 1 26.1 16.2c6-2.6 13.2-1.5 18.2 3.2a16.7 16.7 0 0 1 4.1 18a16.7 16.7 0 0 1 17.7 25.1a16.7 16.7 0 0 1 15.4 10.1a16.7 16.7 0 0 1-3.2 18.2a16.7 16.7 0 0 1 6.2 30c6 2.6 10.2 8.5 10.2 15.4v.4c0 6.9-4.2 12.8-10.2 15.4a16.7 16.7 0 0 1-6.1 30a16.7 16.7 0 0 1-12.2 28.3a16.7 16.7 0 0 1-.4 18.5a16.7 16.7 0 0 1-17.1 6.8a16.7 16.7 0 0 1-4.2 18c-5 4.7-12.2 5.7-18.1 3.2a16.7 16.7 0 0 1-26 16.3a16.7 16.7 0 0 1-28.8 10.7a16.7 16.7 0 0 1-30.5 4.6a16.7 16.7 0 0 1-15.9 9.4z"/><circle cx="256" cy="256" r="111.3" fill="#338af3"/><path fill="#acabb1" d="M280.4 218.2c5.4-4 9-10.5 9-17.9v-11c0-12.4-10-22.3-22.3-22.3H245c-12.3 0-22.3 10-22.3 22.2v11.1c0 7.4 3.6 13.9 9 18L211.5 423s11.1 11 44.5 11s44.5-11 44.5-11z"/><path fill="#eee" d="m256 167l19.3 59.5H338l-50.6 36.8l19.3 59.5L256 286l-50.6 36.8l19.3-59.5l-50.6-36.8h62.6z"/></g>`
	mqInnerSVG                    = `<mask id="circleFlagsMq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMq0)"><path fill="#0052b4" d="M0 0h222.6l31 23.4L289.4 0H512v222.6l-21.5 31l21.5 35.8V512H289.4l-34.2-20.5l-32.6 20.5H0V289.4l22.7-32.6L0 222.6z"/><path fill="#eee" d="M222.6 0v222.6H0v66.8h222.6V512h66.8V289.4H512v-66.8H289.4V0z"/><path fill="#eee" d="M343.4 162.4a6.2 6.2 0 0 1 6.2-6.1h13.2a38 38 0 0 0 0-75.9h-12.6c-4-3-9.6-5-15.6-5c-12.7 0-23 8.3-23 18.4c0 10 10.1 18.2 22.6 18.3h28.6a6.2 6.2 0 0 1 0 12.4h-13.2a38 38 0 0 0 0 75.8h51.1v-31.7h-51.1a6.2 6.2 0 0 1-6.2-6.2zm-200.4 0a6.2 6.2 0 0 1 6.2-6.1h13.2a38 38 0 0 0 0-75.9H150c-4.1-3-9.6-5-15.7-5c-12.6 0-22.9 8.3-22.9 18.4c0 10 10 18.2 22.6 18.3h28.5a6.2 6.2 0 0 1 0 12.4h-13.2a38 38 0 0 0 0 75.8h51.1v-31.7h-51a6.2 6.2 0 0 1-6.3-6.2zm200.4 236.2a6.2 6.2 0 0 1 6.2-6.2h13.2a38 38 0 0 0 0-75.8h-12.6c-4-3-9.6-5-15.6-5c-12.7 0-23 8.3-23 18.4c0 10 10.1 18.1 22.6 18.3h28.6a6.2 6.2 0 0 1 6.2 6.2a6.2 6.2 0 0 1-6.2 6.2h-13.2a38 38 0 0 0 0 75.8h51v-31.7h-51a6.2 6.2 0 0 1-6.2-6.2zm-200.4 0a6.2 6.2 0 0 1 6.2-6.2h13.2a38 38 0 0 0 0-75.8H150c-4.1-3-9.6-5-15.7-5c-12.6 0-22.9 8.3-22.9 18.4c0 10 10 18.1 22.6 18.3h28.5a6.2 6.2 0 0 1 6.2 6.2a6.2 6.2 0 0 1-6.2 6.2h-13.2a38 38 0 0 0 0 75.8h51.1v-31.7h-51a6.2 6.2 0 0 1-6.3-6.2z"/></g>`
	mrInnerSVG                    = `<mask id="circleFlagsMr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMr0)"><path fill="#496e2d" d="M0 0h512v512H0z"/><g fill="#ffda44"><path d="M256 295.8a89 89 0 0 1-87-70a89.4 89.4 0 0 0-2 19a89 89 0 1 0 178 0a89.4 89.4 0 0 0-2-19a89 89 0 0 1-87 70z"/><path d="m256 178l8.3 25.6H291l-21.7 15.8l8.3 25.5L256 229l-21.7 15.8l8.3-25.5l-21.7-15.8h26.8z"/></g></g>`
	msInnerSVG                    = `<mask id="circleFlagsMs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMs0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#338af3" d="M289.4 133.6V256l78 40.4l77.9-40.4V133.6h-156z"/><path fill="#a2001d" d="M289.4 256c0 59.6 77.9 78 77.9 78s78-18.4 78-78h-156z"/><path fill="#333" d="M400.7 189.2h-22.3V167h-22.2v22.2h-22.3v22.3h22.3v66.8h22.2v-66.8h22.3z"/></g>`
	mtInnerSVG                    = `<mask id="circleFlagsMt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMt0)"><path fill="#eee" d="M0 0h256l52 259.2L256 512H0z"/><path fill="#d80027" d="M256 0h256v512H256z"/><path fill="#acabb1" d="M178 100.2V66.8h-33.3v33.4h-33.4v33.4h33.4V167h33.4v-33.4h33.4v-33.4z"/></g>`
	muInnerSVG                    = `<mask id="circleFlagsMu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMu0)"><path fill="#6da544" d="m0 378.3l254-37.1l258 37V512H0z"/><path fill="#ffda44" d="m0 256.1l252.2-33.3L512 256v122.4H0z"/><path fill="#0052b4" d="M0 133.7L249.7 97L512 133.7v122.4H0z"/><path fill="#d80027" d="M0 0h512v133.7H0z"/></g>`
	mvInnerSVG                    = `<mask id="circleFlagsMv0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMv0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#6da544" d="M89 133.6h334v244.8H89z"/><path fill="#eee" d="M297.2 328.3a72.3 72.3 0 1 1 34.4-136a89 89 0 1 0 0 127.3a72 72 0 0 1-34.4 8.7z"/></g>`
	mwInnerSVG                    = `<mask id="circleFlagsMw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMw0)"><path fill="#d80027" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#333" d="M0 0h512v167H0z"/><path fill="#496e2d" d="M0 345h512v167H0z"/><path fill="#d80027" d="m332.5 122.4l-31.2-14.7l16.6-30.3l-34 6.5l-4.3-34.3L256 75l-23.6-25.3L228 84l-34-6.5l16.7 30.3l-31.3 14.7z"/></g>`
	mxInnerSVG                    = `<mask id="circleFlagsMx0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMx0)"><path fill="#eee" d="M144 0h223l33 256l-33 256H144l-32-256z"/><path fill="#496e2d" d="M0 0h144v512H0z"/><path fill="#d80027" d="M368 0h144v512H368z"/><path fill="#ffda44" d="M256 277v10h12l10-22z"/><path fill="#496e2d" d="M160 242a96 96 0 0 0 192 0h-11a85 85 0 0 1-170 0zm39 17l-4 2c-2 2-2 6 1 8c15 14 34 22 54 24v17h12v-17c20-2 39-10 54-24c3-2 3-6 1-8s-6-2-8 0a78 78 0 0 1-53 21c-19 0-38-8-53-21z"/><path fill="#338af3" d="M256 316c-14 0-28-5-40-13l6-9c20 13 48 13 68 0l7 9c-12 8-26 13-41 13z"/><path fill="#751a46" d="M256 174c22 11 12 33 11 34l-2-4c-5-11-18-18-31-18v11c6 0 11 5 11 11c-7 7-9 17-4 26l4 8l-13 23l29-7l18 18v-11l11 11l23-11l-35-21l-5-21l28 16c4 11 12 21 23 26c9-83-42-91-61-91z"/><path fill="#6da544" d="M222 271c-15 0-33-12-38-40l11-2c4 23 18 31 27 31c3 0 5-1 6-3c0-2 0-3-6-5c-3-1-7-2-10-5c-10-12 4-24 11-30c1-1 2-2 1-3c0 0-2-2-5-2c-7 0-12-4-14-11c-2-6 2-13 8-17l5 11c-2 0-2 2-2 4c0 0 1 2 3 2c7 0 14 4 16 9c1 3 2 9-5 15c-7 7-11 12-9 15l5 1c5 2 14 5 13 17c-1 8-8 13-17 13h-1z"/><path fill="#ffda44" d="m234 186l-12 11v11l18-9c3-1 3-5 1-7z"/><circle cx="172" cy="275" r="8" fill="#ffda44"/><circle cx="189" cy="302" r="8" fill="#ffda44"/><circle cx="216" cy="323" r="8" fill="#ffda44"/><circle cx="297" cy="323" r="8" fill="#ffda44"/><circle cx="324" cy="302" r="8" fill="#ffda44"/><circle cx="341" cy="275" r="8" fill="#ffda44"/><rect width="34" height="22" x="239" y="299" fill="#ff9811" rx="11" ry="11"/></g>`
	myInnerSVG                    = `<mask id="circleFlagsMy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMy0)"><path fill="#eee" d="M0 256L256 0h256v55.7l-19.5 33l19.5 33.7v66.8l-22.1 37.7L512 256v66.8l-20.2 38.5l20.2 28.3v66.7l-254.5 28.2L0 456.3v-66.7l26-35.1l-26-31.7z"/><path fill="#d80027" d="M256 256h256v-66.8H222.9zm-33.1-133.6H512V55.7H222.9zM512 512v-55.7H0V512zM0 389.6h512v-66.8H0z"/><path fill="#0052b4" d="M0 0h256v256H0z"/><g fill="#ffda44"><path d="M170.2 219.1a63.3 63.3 0 1 1 30.1-119a78 78 0 1 0 0 111.4a63 63 0 0 1-30 7.6z"/><path d="m188 111.3l11.3 23.5l25.4-5.9l-11.4 23.5l20.4 16.2l-25.4 5.7l.1 26l-20.3-16.2l-20.4 16.2l.1-26l-25.4-5.7l20.4-16.2l-11.4-23.5l25.4 6z"/></g></g>`
	mzInnerSVG                    = `<mask id="circleFlagsMz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMz0)"><path fill="#eee" d="m99 136.8l413 20v31.5l-35.9 66.1l36.2 68.4l-.3 32.4l-413 22z"/><path fill="#496e2d" d="M512 156.8V0H0l122 156.8z"/><path fill="#333" d="M167 188.3v134.5h345.3l-.3-134.5z"/><path fill="#ffda44" d="M512 355.2V512H0l122-156.8z"/><path fill="#a2001d" d="M0 0v512l256-256z"/><path fill="#ffda44" d="m103.6 189.2l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6H87z"/><path fill="#eee" d="M55.1 256h97v44.5h-97z"/><path fill="#333" d="m170.5 205l-15.7-15.8l-51.2 51.2l-51.1-51.2L36.7 205L88 256l-51.2 51.3l15.8 15.5l51.1-51l51.2 51l15.7-15.5l-51.2-51.2z"/></g>`
	naInnerSVG                    = `<mask id="circleFlagsNa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNa0)"><path fill="#eee" d="m0 401.9l160.6-237.1L401.9 0H449l63 63v47.3L350.3 339.2L110.3 512H63L0 449z"/><path fill="#a2001d" d="M0 512h63L512 63V0h-63L0 449z"/><path fill="#0052b4" d="M0 0v401.9L401.9 0z"/><path fill="#496e2d" d="M512 512V110.3L110.3 512z"/><path fill="#ffda44" d="m211.5 144.7l-28.7 13.5L198 186l-31.2-6l-4 31.5l-21.6-23.2l-21.7 23.2l-4-31.5l-31 6l15.2-27.8L71 144.7l28.7-13.5l-15.3-27.8l31.1 6l4-31.5l21.7 23.2L163 78l4 31.5l31-6l-15.2 27.8z"/></g>`
	natoInnerSVG                  = `<mask id="circleFlagsNato0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNato0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#eee" d="m365.9 273.9l46-17.9l-46-17.9l-25.3-9.8l-41-16l-16-41a89.4 89.4 0 0 1 57 57l25.3 9.8a112 112 0 0 0-92-92l-17.9-46l-17.9 46a112 112 0 0 0-92 92l-46 17.9l46 17.9a112 112 0 0 0 92 92l17.9 46l17.9-46a112 112 0 0 0 92-92zm-83.5-44.3L256 256h94.4l-68 26.4L256 256v94.4l-26.4-68L256 256h-94.4l68-26.4L256 256v-94.4zm-54-58.2l-16 41l-41 16a89.4 89.4 0 0 1 57-57zm-57 112.3l41 16l16 41a89.4 89.4 0 0 1-57-57zm112.3 57l16-41l41-16a89.4 89.4 0 0 1-57 57z"/><path fill="#0052b4" d="M350.4 256H256l26.4 26.4zm-188.8 0H256l-26.4-26.4zm94.4 94.4V256l-26.4 26.4zm0-94.4l26.4-26.4l-26.4-68z"/><path fill="#eee" d="M244.9 33.4H267v44.5h-22zM33.4 244.9h44.5V267H33.4zM244.9 434H267v44.5h-22zM434 244.9h44.5V267H434z"/></g>`
	ncInnerSVG                    = `<mask id="circleFlagsNc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNc0)"><path fill="#496e2d" d="M0 352h512v160H0z"/><path fill="#d80027" d="M0 160h512v192H0z"/><path fill="#0052b4" d="M0 0h512v160H0z"/><path fill="#ffda44" d="M398 256a168 168 0 1 1-336 0a168 168 0 0 1 336 0z"/><path fill="#333" d="M230 80a176 176 0 1 0 0 352a176 176 0 0 0 0-352zm-8 16v34c-8 4-14 8-15 16l-1 16l6-3l10-4v8a24 24 0 0 0-16 23a24 24 0 0 0 16 22v8h-20v16h20v8a24 24 0 0 0-16 22a24 24 0 0 0 16 23v7h-20v17h20v7a24 24 0 0 0-16 23a24 24 0 0 0 16 22v8c-10 0-18 3-23 7c-7 5-22 20-30 28a160 160 0 0 1 53-308zm16 0a160 160 0 0 1 51 309l-30-29c-5-4-12-6-21-7v-8a24 24 0 0 0 16-22a24 24 0 0 0-16-23v-7h21v-16h-21v-8a24 24 0 0 0 16-23a24 24 0 0 0-16-22v-8h21v-16h-21v-7a24 24 0 0 0 16-23a24 24 0 0 0-16-23v-20c6-5 11-13 17-23c4-8 2-11 2-11l-2 3l-17 11V96z"/></g>`
	neInnerSVG                    = `<mask id="circleFlagsNe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNe0)"><path fill="#eee" d="m0 144.7l255.3-36.5L512 144.7v222.6L250.5 407L0 367.3z"/><path fill="#ff9811" d="M0 0h512v144.7H0z"/><path fill="#6da544" d="M0 367.3h512V512H0z"/><circle cx="256" cy="256.1" r="89" fill="#ff9811"/></g>`
	nfInnerSVG                    = `<mask id="circleFlagsNf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNf0)"><path fill="#6da544" d="M0 0h144.7l108.4 41.2L367.3 0H512v512H367.3l-110.2-41.4L144.7 512H0z"/><path fill="#eee" d="M144.7 0h222.6v512H144.7z"/><path fill="#6da544" d="m323 334l-67-211.6L189.3 334h50.1v55.7h33.4V334z"/></g>`
	ngInnerSVG                    = `<mask id="circleFlagsNg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNg0)"><path fill="#6da544" d="M0 0v512h160l96-64l96 64h160V0H352l-96 64l-96-64Z"/><path fill="#eee" d="M160 0h192v512H160Z"/></g>`
	niInnerSVG                    = `<mask id="circleFlagsNi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNi0)"><path fill="#338af3" d="M0 0h512v144.7l-41.4 111.7L512 367.3V512H0V367.3l41.5-107.9L0 144.7z"/><path fill="#eee" d="M0 144.7h512v222.6H0z"/><path fill="#ffda44" d="M256 178a78 78 0 1 0 0 156a78 78 0 0 0 0-156zm0 122.5a44.5 44.5 0 1 1 0-89a44.5 44.5 0 0 1 0 89z"/><path fill="#0052b4" d="M294.6 267.1L256 256l-38.6 11.1l-12.8 22.3h102.8z"/><path fill="#338af3" d="M256 200.3L230.3 245l25.7 11l25.7-11.1z"/><path fill="#6da544" d="M217.4 267.1h77.2L281.7 245h-51.4z"/></g>`
	nlInnerSVG                    = `<mask id="circleFlagsNl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNl0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#a2001d" d="M0 0h512v167H0z"/><path fill="#0052b4" d="M0 345h512v167H0z"/></g>`
	nlFrInnerSVG                  = `<mask id="circleFlagsNlFr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNlFr0)"><path fill="#eee" d="M0 0v51.7l32 103.4H0v103.5L32 362H0v150h150v-32l103.4 32h103.5v-32l103.4 32H512v-51.7l-32-103.4h32V253.4L480 150h32V0H362v32L258.6 0H155.1v32L51.7 0H0z"/><path fill="#0052b4" d="M51.7 0L512 460.3V356.9L155.1 0H51.7zm206.9 0L512 253.4V150L362 0H258.6zM0 51.8v103.4L356.8 512h103.4L0 51.8zm0 206.9v103.4L149.9 512h103.5L0 258.7z"/><path fill="#d80027" d="M261.6 222.2a18 18 0 0 0-25.5 0a18 18 0 0 0-.6.6a38 38 0 0 0 53.7 53.7a18 18 0 0 0 .6-.6a18 18 0 0 0 0-25.5a18 18 0 0 0-25.4 0a2 2 0 0 1-2.8 0a2 2 0 0 1 0-2.8a18 18 0 0 0 0-25.4zm-100-100a18 18 0 0 0-25.5 0a18 18 0 0 0-.6.6a38 38 0 0 0 53.7 53.7a18 18 0 0 0 .6-.6a18 18 0 0 0 0-25.5a18 18 0 0 0-25.4 0a2 2 0 0 1-2.8 0a2 2 0 0 1 0-2.8a18 18 0 0 0 0-25.4zm200 200a18 18 0 0 0-25.5 0a18 18 0 0 0-.6.6a38 38 0 0 0 53.7 53.7a18 18 0 0 0 .6-.6a18 18 0 0 0 0-25.5a18 18 0 0 0-25.4 0a2 2 0 0 1-2.8 0a2 2 0 0 1 0-2.8a18 18 0 0 0 0-25.4zm-254-45.9a18 18 0 0 0-25.5 0a18 18 0 0 0-.6.5a38 38 0 0 0 53.7 53.8a18 18 0 0 0 .6-.6a18 18 0 0 0 0-25.5a18 18 0 0 0-25.4 0a2 2 0 0 1-2.8 0a2 2 0 0 1 0-2.9a18 18 0 0 0 0-25.3zm100 100a18 18 0 0 0-25.5 0a18 18 0 0 0-.6.6a38 38 0 0 0 53.7 53.7a18 18 0 0 0 .6-.6a18 18 0 0 0 0-25.5a18 18 0 0 0-25.4 0a2 2 0 0 1-2.8 0a2 2 0 0 1 0-2.8a18 18 0 0 0 0-25.4zm207-207a18 18 0 0 0-25.5 0a18 18 0 0 0-.6.5a38 38 0 0 0 53.7 53.8a18 18 0 0 0 .6-.6a18 18 0 0 0 0-25.5a18 18 0 0 0-25.4 0a2 2 0 0 1-2.8 0a2 2 0 0 1 0-2.9a18 18 0 0 0 0-25.3zm-100-100a18 18 0 0 0-25.5 0a18 18 0 0 0-.6.6a38 38 0 0 0 53.7 53.7a18 18 0 0 0 .6-.6a18 18 0 0 0 0-25.6a18 18 0 0 0-25.4 0a2 2 0 0 1-2.8 0a2 2 0 0 1 0-2.8a18 18 0 0 0 0-25.3z"/></g>`
	noInnerSVG                    = `<mask id="circleFlagsNo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNo0)"><path fill="#d80027" d="M0 0h100.2l66.1 53.5L233.7 0H512v189.3L466.3 257l45.7 65.8V512H233.7l-68-50.7l-65.5 50.7H0V322.8l51.4-68.5l-51.4-65z"/><path fill="#eee" d="M100.2 0v189.3H0v33.4l24.6 33L0 289.5v33.4h100.2V512h33.4l30.6-26.3l36.1 26.3h33.4V322.8H512v-33.4l-24.6-33.7l24.6-33v-33.4H233.7V0h-33.4l-33.8 25.3L133.6 0z"/><path fill="#0052b4" d="M133.6 0v222.7H0v66.7h133.6V512h66.7V289.4H512v-66.7H200.3V0z"/></g>`
	northernCyprusInnerSVG        = `<mask id="circleFlagsNorthernCyprus0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNorthernCyprus0)"><path fill="#eee" d="M0 0v48l32 24L0 96v320l32 24l-32 24v48h512v-48l-32-24l32-24V96l-32-24l32-24V0Z"/><path fill="#d80027" d="M0 464h512v-48H0ZM0 48v48h512V48Zm270 255v-94l55 76l-89-29l89-29z"/><path fill="#d80027" d="M185 167a89 89 0 1 0 62 153a72 72 0 0 1-34 8a72 72 0 1 1 34-136a89 89 0 0 0-62-25Z"/></g>`
	npInnerSVG                    = `<mask id="circleFlagsNp0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNp0)"><path fill="#eee" d="M228.4 0H512v512h-27.4L159.4 258.5z"/><path fill="#0052b4" d="M510.5 283.8L228.5 0h-39.3l-70.9 253.8L445.2 512h39.4L256.4 283.8h254z"/><path fill="#d80027" d="M445.2 256L189.2 0H0v512h445.2l-256-256z"/><path fill="#eee" d="m243.5 378l-31.3-14.7L229 333l-34 6.5l-4.3-34.3l-23.6 25.2l-23.7-25.2l-4.3 34.3l-34-6.5l16.7 30.3L90.4 378l31.3 14.7l-16.6 30.3l34-6.5l4.2 34.3l23.7-25.2l23.6 25.2l4.3-34.3l34 6.5l-16.7-30.3zM149.8 89l-3.2 25l-24.7-4.7l12.1 22l-22.7 10.7h-11.1a66.8 66.8 0 0 0 133.5 0h-11.1L200 131.3l12-22l-24.6 4.7l-3.1-25l-17.2 18.4z"/></g>`
	nrInnerSVG                    = `<mask id="circleFlagsNr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNr0)"><path fill="#0052b4" d="M0 0h512v233.7L494.3 257l17.7 21.3V512H0V278.3l18.7-22.9L0 233.7z"/><path fill="#eee" d="m211.5 345l-28.7 13.5l15.3 27.8l-31.2-6l-4 31.5l-21.6-23.1l-21.7 23.1l-4-31.4l-31.1 6l15.3-27.9L71 345l28.7-13.4l-15.3-27.8l31.1 6l4-31.5l21.7 23.1l21.7-23.1l4 31.4l31-6l-15.2 27.9z"/><path fill="#ffda44" d="M0 233.7h512v44.6H0z"/></g>`
	nuInnerSVG                    = `<mask id="circleFlagsNu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNu0)"><path fill="#ffda44" d="M0 256L256 0h256v512H0z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32l42-16l41 16h45l-8-16l8-15v-14l-16-42l16-41V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#0052b4" d="M128 256v-83l83 83zm128-45l-83-83h83z"/><path fill="#d80027" d="m128 128l128 128v-31l-97-97Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Z"/><circle cx="64" cy="64" r="48" fill="#0052b4"/><path fill="#ffda44" d="m50 198l14-44l15 44l-38-27h47zM162 86l14-44l15 44l-38-27h47ZM64 17l28 86l-73-53h90l-73 53Z"/></g>`
	nzInnerSVG                    = `<mask id="circleFlagsNz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNz0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Zm382 92l-11 35h-37l30 21l-12 35l30-22l30 22l-12-35l30-21h-37l-11-35Zm61 72l-11 35h-37l30 21l-11 35l29-21l30 21l-12-35l30-21h-37Zm-123 10l-11 35h-37l30 22l-11 35l29-22l30 22l-11-35l29-22h-36zm59 130l-11 35h-37l30 21l-11 35l29-21l30 21l-11-35l29-21h-36z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97zm251 201l-5 18h-19l15 10l-6 18l15-11l15 11l-5-18l14-10h-18Zm-59-129l-5 17h-19l15 11l-6 17l15-11l15 11l-6-17l15-11h-18l-6-17zm123-11l-6 18h-18l15 11l-6 17l15-11l15 11l-6-17l15-11h-18l-6-18zm-61-72l-6 17h-18l15 11l-6 17l15-10l15 10l-6-17l15-11h-18z"/></g>`
	occitaniaInnerSVG             = `<mask id="circleFlagsOccitania0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsOccitania0)"><path fill="#d80027" d="M0 0h512v512H0z"/><circle cx="256" cy="64" r="32" fill="#ffda44"/><circle cx="64" cy="256" r="32" fill="#ffda44"/><circle cx="448" cy="256" r="32" fill="#ffda44"/><circle cx="256" cy="448" r="32" fill="#ffda44"/><circle cx="352" cy="400" r="32" fill="#ffda44"/><circle cx="400" cy="352" r="32" fill="#ffda44"/><circle cx="352" cy="112" r="32" fill="#ffda44"/><circle cx="400" cy="160" r="32" fill="#ffda44"/><circle cx="160" cy="112" r="32" fill="#ffda44"/><circle cx="112" cy="160" r="32" fill="#ffda44"/><circle cx="112" cy="352" r="32" fill="#ffda44"/><circle cx="160" cy="400" r="32" fill="#ffda44"/><path fill="#ffda44" d="m256 60l-103 52l20 61l-61-20l-52 103l52 103l61-20l-20 61l103 52l103-52l-20-61l61 20l52-103l-52-103l-61 20l20-61zm0 36l64 32l-32 96l96-32l32 64l-32 64l-96-32l32 96l-64 32l-64-32l32-96l-96 32l-32-64l32-64l96 32l-32-96z"/></g>`
	olympicsInnerSVG              = `<mask id="circleFlagsOlympics0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsOlympics0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#338af3" d="M109.6 153.1A71.8 71.8 0 0 0 38 224.7a71.8 71.8 0 0 0 71.6 71.6a71.8 71.8 0 0 0 71.6-71.6a71.8 71.8 0 0 0-71.6-71.6zm0 20c28.6 0 51.6 23 51.6 51.6s-23 51.6-51.6 51.6s-51.6-23-51.6-51.6s23-51.6 51.6-51.6z"/><path fill="#333" d="M256 153.1a71.8 71.8 0 0 0-71.6 71.6a71.8 71.8 0 0 0 71.6 71.6a71.8 71.8 0 0 0 71.6-71.6a71.8 71.8 0 0 0-71.6-71.6zm0 20c28.6 0 51.6 23 51.6 51.6s-23 51.6-51.6 51.6s-51.6-23-51.6-51.6s23-51.6 51.6-51.6z"/><path fill="#d80027" d="M402.4 153.1a71.8 71.8 0 0 0-71.6 71.6a71.8 71.8 0 0 0 71.6 71.6a71.8 71.8 0 0 0 71.6-71.6a71.8 71.8 0 0 0-71.6-71.6zm0 20c28.6 0 51.6 23 51.6 51.6s-23 51.6-51.6 51.6s-51.6-23-51.6-51.6s23-51.6 51.6-51.6z"/><path fill="#ffda44" d="M182.8 215.7a71.8 71.8 0 0 0-71.6 71.6a71.8 71.8 0 0 0 71.6 71.6a71.8 71.8 0 0 0 71.6-71.6a71.8 71.8 0 0 0-71.6-71.6zm0 20c28.6 0 51.6 23 51.6 51.6s-23 51.6-51.6 51.6s-51.6-23-51.6-51.6s23-51.6 51.6-51.6z"/><path fill="#6da544" d="M329.2 215.7a71.8 71.8 0 0 0-71.6 71.6a71.8 71.8 0 0 0 71.6 71.6a71.8 71.8 0 0 0 71.6-71.6a71.8 71.8 0 0 0-71.6-71.6zm0 20c28.6 0 51.6 23 51.6 51.6s-23 51.6-51.6 51.6s-51.6-23-51.6-51.6s23-51.6 51.6-51.6z"/><path fill="#338af3" d="m166.9 181.7l-16 12a51.5 51.5 0 0 1 0 61.9l16 12a71.6 71.6 0 0 0 0-85.9z"/><path fill="#333" d="m313.3 181.7l-16 12a51.5 51.5 0 0 1 0 61.9l16 12a71.6 71.6 0 0 0 0-85.9zM225 266l-12 16a71.7 71.7 0 0 0 43 14.3v-20a51.5 51.5 0 0 1-31-10.3z"/><path fill="#d80027" d="m371.4 266l-12 16a71.7 71.7 0 0 0 43 14.3v-20a51.5 51.5 0 0 1-31-10.3z"/></g>`
	omInnerSVG                    = `<mask id="circleFlagsOm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsOm0)"><path fill="#eee" d="M189.2 0H512v167l-347.5 24.6z"/><path fill="#6da544" d="m163 320l349 25v167H189.2z"/><path fill="#d80027" d="M0 0h189.2v167H512v178H189.2v167H0z"/><path fill="#eee" d="M156.6 112.7L133 89l-15.7 15.8L101.5 89L78 112.7l15.8 15.7L78 144l23.6 23.6l15.8-15.7l15.7 15.7l23.6-23.6l-15.7-15.7z"/></g>`
	otomiInnerSVG                 = `<mask id="circleFlagsOtomi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsOtomi0)"><path fill="#ff9811" d="m0 160l256-32l256 32v192l-256 32L0 352Z"/><path fill="#eee" d="m0 160l256-32l256 32v32H0Z"/><path fill="#338af3" d="M0 0h512v160H0Z"/><path fill="#eee" d="m0 352l256 32l256-32v-32H0Z"/><path fill="#6da544" d="M0 352h512v160H0Z"/><circle cx="256" cy="256" r="160" fill="#eee"/><circle cx="256" cy="256" r="128" fill="#338af3"/><path fill="#ff9811" d="M207 139v68h-68l49 49l-49 49h68v68l49-49l49 49v-68h68l-49-49l49-49h-68v-68l-49 49z"/></g>`
	paInnerSVG                    = `<mask id="circleFlagsPa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPa0)"><path fill="#eee" d="M0 0h256l256 256v256H256L0 256z"/><path fill="#0052b4" d="M0 256v256h256V256z"/><path fill="#d80027" d="M256 0h256v256H256z"/><path fill="#0052b4" d="m152.4 89l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51L82.2 140h53.6z"/><path fill="#d80027" d="m359.6 289.4l16.6 51h53.6L386.4 372l16.6 51l-43.4-31.5l-43.4 31.6l16.6-51l-43.4-31.6H343z"/></g>`
	peInnerSVG                    = `<mask id="circleFlagsPe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPe0)"><path fill="#d80027" d="M0 0h167l86 41.2L345 0h167v512H345l-87.9-41.4L167 512H0z"/><path fill="#eee" d="M167 0h178v512H167z"/></g>`
	pfInnerSVG                    = `<mask id="circleFlagsPf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPf0)"><path fill="#d80027" d="M0 0h512v133.7l-52 126l52 118.8V512H0V378.5l53-121L0 133.6z"/><path fill="#eee" d="M0 133.7h512v244.8H0z"/><path fill="#ffda44" d="M345 256.1c0 49.2-39.8 78-89 78s-89-28.8-89-78a89 89 0 1 1 178 0z"/><path fill="#0052b4" d="M345 256.1a89 89 0 1 1-178 0"/><path fill="#d80027" d="M200.3 233.8h22.3v44.6h-22.3zm89 0h22.4v44.6h-22.3zM245 200.4h22v78h-22z"/></g>`
	pgInnerSVG                    = `<mask id="circleFlagsPg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPg0)"><path fill="#333" d="M512 512L301.3 226.6L0 0v512z"/><path fill="#a2001d" d="m0 0l512 512V0z"/><path fill="#eee" d="m195 346l4.2 12.9h13.6l-11 8l4.2 12.8l-11-7.9l-11 7.9l4.2-12.9l-11-8h13.6zm-34.6-123.7l7 21.5h22.5L171.7 257l7 21.5l-18.3-13.3l-18.2 13.3l6.9-21.5l-18.2-13.2h22.5zm0 157.4l7 21.5h22.5l-18.2 13.2l7 21.6l-18.3-13.4l-18.2 13.4l6.9-21.6l-18.2-13.2h22.5zm66-101.2l7 21.5H256l-18.2 13.3l6.9 21.4l-18.2-13.2l-18.3 13.2l7-21.4l-18.3-13.3h22.5zm-132 0l7 21.5h22.5l-18.2 13.3l7 21.4l-18.3-13.2l-18.3 13.2l7-21.4L64.8 300h22.5z"/><path fill="#ffda44" d="M430.3 189a67 67 0 0 0-42.6-19.7l48.2-37.6a89.7 89.7 0 0 0-70.3-26A67.2 67.2 0 0 0 348.5 76l-19.2 38.2a28 28 0 0 0-4.8-6.3a28 28 0 1 0-39.7 39.7a28 28 0 0 0 6.3 4.8L253 171.5a67.2 67.2 0 0 0 29.6 17.2a89.7 89.7 0 0 0 26 70.2l42.2-54.2a16.8 16.8 0 0 1 11.9-4.9c4.4 0 8.7 1.7 12 5a16.8 16.8 0 0 1 4.8 11.9c0 4.5-1.7 8.8-4.9 12l15.9 15.9a39.1 39.1 0 0 0 3.7-51.4a44.7 44.7 0 0 1 20.2 11.6a45 45 0 0 1 0 63.6l15.9 15.9a67 67 0 0 0 0-95.4z"/></g>`
	phInnerSVG                    = `<mask id="circleFlagsPh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPh0)"><path fill="#0052b4" d="M0 0h512v256l-265 45.2z"/><path fill="#d80027" d="M210 256h302v256H0z"/><path fill="#eee" d="M0 0v512l256-256z"/><path fill="#ffda44" d="M175.3 256L144 241.3l16.7-30.3l-34 6.5l-4.3-34.3l-23.6 25.2L75 183.2l-4.3 34.3l-34-6.5l16.7 30.3L22.3 256l31.2 14.7L37 301l34-6.5l4.2 34.3l23.7-25.2l23.6 25.2l4.3-34.3l34 6.5l-16.7-30.3zm-107-155.8l10.4 14.5l17-5.4l-10.6 14.4l10.4 14.5l-17-5.6L68 147l.2-17.9l-17-5.6l17-5.4zm0 264.8l10.4 14.6l17-5.4l-10.6 14.3l10.4 14.6l-17-5.7L68 411.8l.2-17.9l-17-5.6l17-5.4zm148.4-132.4L206.3 247l-17-5.4l10.5 14.4l-10.4 14.6l17-5.7l10.6 14.4l-.1-17.9l17-5.6l-17.1-5.4z"/></g>`
	pkInnerSVG                    = `<mask id="circleFlagsPk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPk0)"><path fill="#eee" d="M0 0h133.6l50.2 252.5L133.6 512H0z"/><path fill="#496e2d" d="M133.6 0H512v512H133.6z"/><path fill="#eee" d="M365.5 298.3A72.3 72.3 0 1 1 313.7 168a89 89 0 1 0 74.6 103.2a72 72 0 0 1-22.8 27.2zM364 167l18.2 19.6l24.3-11.3l-13 23.5l18.2 19.6l-26.3-5.1l-13 23.4l-3.3-26.7l-26.1-5l24.4-11.3z"/></g>`
	pkJkInnerSVG                  = `<mask id="circleFlagsPkJk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPkJk0)"><path fill="#ff9811" d="M0 0h224l32 88l-32 88l-96 32L0 176Z"/><path fill="#eee" d="m0 224l256-32l256 32v48l-32 24l32 24v48l-32 24l32 24v48l-256 32L0 464v-48l32-24l-32-24v-48l32-24l-32-24Z"/><path fill="#496e2d" d="M224 0v176H0v48h512V0ZM0 272v48h512v-48zm0 96v48h512v-48zm0 96v48h512v-48z"/><circle cx="355.6" cy="110.2" r="80.5" fill="#eee"/><circle cx="376.2" cy="91.9" r="73.8" fill="#496e2d"/><path fill="#eee" d="m370 76l47-20l-27 44l-4-50l34 38z"/></g>`
	pkSdInnerSVG                  = `<mask id="circleFlagsPkSd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPkSd0)"><path fill="#496e2d" d="M0 0h512v512H0z"/><ellipse cx="256" cy="284" fill="#eee" rx="140" ry="132"/><ellipse cx="256" cy="249" fill="#496e2d" rx="122" ry="111"/><path fill="#eee" d="M256.8 149.6h-1.6c-.3 8.5-4.7 15-11.3 20.9c-6.6 5.9-15.4 11.1-24.1 17a95.8 95.8 0 0 0-24 21.6a55.3 55.3 0 0 0-10.8 34.5c0 14.4 4.5 25.3 11.1 34a99 99 0 0 0 24.2 21.3c8.8 6 17.5 11.2 24 17.2a32 32 0 0 1 10 15.1c.4 2.1.8 4.2 1 6.5h1.5c.1-2.3.5-4.4.9-6.4a32 32 0 0 1 10-15.2c6.5-6 15.2-11.2 24-17.2a99 99 0 0 0 24.2-21.2c6.6-8.8 11-19.7 11-34a55.4 55.4 0 0 0-10.7-34.6a96 96 0 0 0-24-21.5c-8.8-6-17.5-11.2-24.1-17.1c-6.6-6-11-12.4-11.3-21z"/><path fill="#eee" d="M262.6 117.8h29l-23 15.3l9 23.4l-22.3-14.7h-.1l-22 14.6l8.3-24l-21.2-14.6h26.2l7.6-22.2zm56.6 28.9c-2 4.4-3.1 9.6-3 13.8c-2.6-1.5-10.4-5-12.8-5.5c.1 5.8 1.9 12 4.7 16c-6-.7-11.2.3-15.7 2c9.6 13 20.4 14.8 34 10.8a26 26 0 0 0 2.7 12.5a30.6 30.6 0 0 0-15.8-.3c6.8 12.6 17.9 18.5 31 15.7a22.6 22.6 0 0 0-2.3 13a30.7 30.7 0 0 0-14.6-6.1c3.9 16 12.3 22 23.4 26.6c-3.6 1.5-7.8 10.3-8.1 11.4a30.6 30.6 0 0 0-10.7-11.7c-4.2 13-1.5 25.3 9 32.6c-1.1-1.1-8 4.7-9.3 6.4a30.6 30.6 0 0 0-5.4-14.8c-8.7 10.2-11 22.4-4.5 33a26.2 26.2 0 0 0-9 4.4a30.6 30.6 0 0 0-2.6-15.6c-9.5 7.4-13.7 22.4-10.7 31.4c-3.8 0-7.5.7-10.9 2.2a30.7 30.7 0 0 0 1.2-15.8c-13.5 6.5-19.5 17.4-17.2 30.9l-9.7 4l4.7 2.2l6.4-3.6c13.6 6.5 26.6 6.5 36.4-2c-4-2.8-8.8-5-13-5.6a25.9 25.9 0 0 0 6.1-7.5c12.6 6.3 26.3 1.3 34-10.2a31.4 31.4 0 0 0-13.9-2.5a25.9 25.9 0 0 0 4.1-8.3c13.5 1.5 26-4.3 30.3-16.6c-4.8-1-10.1-1-14.1.1c.9-1.2 2.5-9.7 2.6-10.8c14.2-4.4 22.2-12.6 22.6-27.2c-4.8.9-9.7 3-13 5.5a71.3 71.3 0 0 0-2-11.8a30.1 30.1 0 0 0 9.8-36a31 31 0 0 0-9.5 10.4c.4-1-4.3-8.6-6.8-9.9c7.7-12 5.3-26.8-4-36c-2.6 4-4.5 9-4.9 13.2c-1.4-3-7.7-5.8-10-6.7c3.4-14.4 2.7-24.6-9.5-33.6zm-126.4 0c2 4.4 3.1 9.6 3 13.8c2.6-1.5 10.5-5 12.8-5.5c-.1 5.8-1.9 12-4.7 16c6-.7 11.2.3 15.7 2c-9.6 13-20.4 14.8-34 10.8a26 26 0 0 1-2.7 12.5a30.6 30.6 0 0 1 15.8-.3c-6.8 12.6-17.9 18.5-31 15.7a22.6 22.6 0 0 1 2.3 13a30.7 30.7 0 0 1 14.6-6.1c-3.9 16-12.3 22-23.4 26.6c3.6 1.5 7.8 10.3 8.1 11.4a30.6 30.6 0 0 1 10.7-11.7c4.2 13 1.5 25.3-9 32.6c1.1-1.1 8 4.7 9.3 6.4a30.6 30.6 0 0 1 5.4-14.8c8.7 10.2 11 22.4 4.5 33a26.2 26.2 0 0 1 9 4.4a30.6 30.6 0 0 1 2.6-15.6c9.5 7.4 13.7 22.4 10.7 31.4c3.8 0 7.5.7 10.9 2.2a30.6 30.6 0 0 1-1.2-15.8c13.5 6.5 19.5 17.4 17.2 30.9l9.7 4l-4.7 2.2l-6.4-3.6c-13.6 6.5-26.6 6.5-36.4-2c4-2.8 8.8-5 13-5.6a25.9 25.9 0 0 1-6.1-7.5c-12.6 6.3-26.3 1.3-34-10.2c4.5-2 9.7-2.9 13.9-2.5a25.9 25.9 0 0 1-4.1-8.3c-13.5 1.5-26-4.3-30.3-16.6c4.8-1 10.1-1 14.1.1c-.9-1.2-2.5-9.7-2.6-10.8c-14.2-4.4-22.2-12.6-22.6-27.2c4.8.9 9.7 3 13 5.5c.2-2.7 1.4-10.2 2-11.8a30.1 30.1 0 0 1-9.8-36a31 31 0 0 1 9.5 10.4c-.4-1 4.3-8.6 6.8-9.9a28.7 28.7 0 0 1 4-36c2.6 4 4.5 9 4.9 13.2c1.4-3 7.8-5.8 10-6.7c-3.4-14.4-2.7-24.6 9.5-33.6z"/></g>`
	plInnerSVG                    = `<mask id="circleFlagsPl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPl0)"><path fill="#d80027" d="m0 256l256.4-44.3L512 256v256H0z"/><path fill="#eee" d="M0 0h512v256H0z"/></g>`
	pmInnerSVG                    = `<mask id="circleFlagsPm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPm0)"><path fill="#338af3" d="M160 0h352v512H160l-32-256z"/><path fill="#eee" d="m0 160l80-32l80 32v192l-80 32l-80-32z"/><rect width="160" height="160" fill="#d80027" ry="0"/><path fill="#6da544" d="M160 21V0h-20L0 139v21h20z"/><path fill="#6da544" d="M0 21V0h20l140 139v21h-20z"/><path fill="#eee" d="M0 64h160v32H0z"/><path fill="#eee" d="M64 0h32v160H64z"/><path fill="#d80027" d="M0 512h160V352H0z"/><path fill="#ffda44" d="M340 146v158h-40c-24 0-24-32-24-32h-96v32h20l147 32l145-32v-24h-80s0 24-24 24h-32V146z"/><path fill="#ffda44" d="M356 138v32c-31-11-61 8-80-32c33 23 51-10 80 0zm-72 46s-24 24-16 72h-72s32-56 88-72zm16 104h95c-31-48 0-104 0-104h-95s-40 56 0 104zm101-24a96 96 0 0 1 4-64h63s-16 32-1 64zm-101 40h192s-8 0-16 24s-16 48-56 48H276c-39 0-52-50-76-72z"/><rect width="96" height="32" x="32" y="384" fill="#ffda44" rx="16" ry="16"/><rect width="96" height="32" x="32" y="448" fill="#ffda44" rx="16" ry="16"/></g>`
	pnInnerSVG                    = `<mask id="circleFlagsPn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPn0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#acabb1" d="m401 156l-34 11l-33-11l-11-28l11-28h67l-11 28Z"/><path fill="#338af3" d="M289 156v122c0 60 78 78 78 78s78-18 78-78V156H289z"/><path fill="#6da544" d="M296 307c20 37 71 49 71 49s52-12 71-49l-71-118l-71 118z"/><path fill="#ffda44" d="m445 277l-78-121l-78 121v1a62 62 0 0 0 7 29l71-110l71 110a62 62 0 0 0 7-29z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/></g>`
	prInnerSVG                    = `<mask id="circleFlagsPr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPr0)"><path fill="#eee" d="m27 63.3l485 39.1v102.4L477.3 259l34.7 48.2v102.4L27.4 446.9z"/><path fill="#d80027" d="m0 0l51.2 102.4H512V0zm0 512h512V409.6H51.2zm180-204.8h332V204.8H180z"/><path fill="#0052b4" d="M0 0v512l256-256z"/><path fill="#eee" d="m103.6 189.2l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6H87z"/></g>`
	psInnerSVG                    = `<mask id="circleFlagsPs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPs0)"><path fill="#eee" d="M41.3 121.9L512 167v178L43.8 391.3z"/><path fill="#333" d="M0 0h512v167H111z"/><path fill="#6da544" d="M111 345h401v167H0z"/><path fill="#d80027" d="M0 0v512l256-256z"/></g>`
	ptInnerSVG                    = `<mask id="circleFlagsPt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPt0)"><path fill="#6da544" d="M0 512h167l37.9-260.3L167 0H0z"/><path fill="#d80027" d="M512 0H167v512h345z"/><circle cx="167" cy="256" r="89" fill="#ffda44"/><path fill="#d80027" d="M116.9 211.5V267a50 50 0 1 0 100.1 0v-55.6H117z"/><path fill="#eee" d="M167 283.8c-9.2 0-16.7-7.5-16.7-16.7V245h33.4v22c0 9.2-7.5 16.7-16.7 16.7z"/></g>`
	ptThirtyInnerSVG              = `<mask id="circleFlagsPt300"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPt300)"><path fill="#0052b4" d="M0 0h144.7l108.4 41.2L367.3 0H512v512H367.3l-110.2-41.4L144.7 512H0z"/><path fill="#ffda44" d="M144.7 0h222.6v512H144.7z"/><path fill="#d80027" d="m328.4 278.3l16.7 22.2v-89l-16.7 22.2h-50v-50l22.2-16.7h-89l22.2 16.7v50h-50l-16.7-22.2v89l16.7-22.2h50v50L211.6 345h89l-22.2-16.7v-50z"/><path fill="#eee" d="M323 244.9h-56v-55.7h-22V245h-55.7v22H245v55.7h22V267h55.7z"/></g>`
	ptTwentyInnerSVG              = `<mask id="circleFlagsPt200"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPt200)"><path fill="#0052b4" d="M0 512h189.2l41.4-256L189.2 0H0z"/><path fill="#eee" d="M189.2 0H512v512H189.2z"/><path fill="#d80027" d="M91.8 89v41.8a37.6 37.6 0 1 0 75.2 0V89z"/><path fill="#eee" d="M129.4 146c-8.4 0-15.3-6.8-15.3-15.2v-19.5h30.6v19.5c0 8.4-6.9 15.3-15.3 15.3z"/><path fill="#ffda44" d="M278.3 267.1h-66.8a22.3 22.3 0 0 0-44.5 0h-66.8a23 23 0 0 0 23 22.3h-.8c0 12.3 10 22.3 22.3 22.3c0 12.2 10 22.2 22.3 22.2h44.5c12.3 0 22.2-10 22.2-22.2c12.3 0 22.3-10 22.3-22.3h-.7a23 23 0 0 0 23-22.3zm9.1-35.6l-5.8 9.4l7 8.4l-10.6-2.7l-5.9 9.4l-.8-11l-10.7-2.7l10.3-4.1l-.8-11l7 8.4zm-34-28.6l-2.4 10.8l9.6 5.5l-11 1.2l-2.3 10.7l-4.5-10l-11 1.1l8.2-7.4l-4.4-10l9.5 5.5zm-42-15.2l1.6 10.9l10.9 2l-10 4.7l1.6 11l-7.7-8l-9.9 4.9l5.2-9.8l-7.7-8l10.9 2zm-44.4 0l5.1 9.7l10.9-2l-7.7 8l5.2 9.8l-9.9-4.9l-7.7 8l1.6-11l-10-4.8l11-2zm-42 15.2l8.2 7.4l9.6-5.5l-4.5 10l8.2 7.4l-11-1.1l-4.5 10l-2.3-10.7l-11-1.2l9.6-5.5zm-34 28.6l10.2 4.1l7-8.4l-.7 11l10.2 4.1L107 245l-.8 11l-5.8-9.4l-10.7 2.7l7-8.4z"/></g>`
	pwInnerSVG                    = `<mask id="circleFlagsPw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPw0)"><path fill="#338af3" d="M0 0h512v512H0z"/><circle cx="200.3" cy="256" r="111.3" fill="#ffda44"/></g>`
	pyInnerSVG                    = `<mask id="circleFlagsPy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPy0)"><path fill="#eee" d="m0 144.7l255.3-36.5L512 144.7v222.6L250.5 407L0 367.3z"/><path fill="#d80027" d="M0 0h512v144.7H0z"/><path fill="#0052b4" d="M0 367.3h512V512H0z"/><path fill="#6da544" d="m319 182l-23.6 23.5a55.5 55.5 0 0 1-39.4 95a55.7 55.7 0 0 1-39.3-95L193 182a89 89 0 1 0 126 0z"/><path fill="#ffda44" d="m256 211.5l8.3 25.5H291l-21.7 15.8l8.3 25.5l-21.7-15.8l-21.7 15.8l8.3-25.5l-21.7-15.8h26.8z"/></g>`
	qaInnerSVG                    = `<mask id="circleFlagsQa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsQa0)"><path fill="#eee" d="M0 0h173l61 255.8L173.4 512H0z"/><path fill="#751a46" d="m173 0l-72.7 30.8L176 63l-75.7 32.2l75.7 32.1l-75.7 32.2l75.7 32.1l-75.7 32.1l75.7 32.2l-75.7 32.2l75.7 32.1l-75.7 32.2l75.7 32.1l-75.7 32.2l75.7 32.1l-75.7 32.2l73.1 31H512V0z"/></g>`
	quechuaInnerSVG               = `<mask id="circleFlagsQuechua0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsQuechua0)"><path fill="#4a1f63" d="M0 512h512v-70l-256-32L0 442Z"/><path fill="#0052b4" d="M0 442h512v-70l-256-32L0 372Z"/><path fill="#d80027" d="M0 0h512v70l-256 32L0 70Z"/><path fill="#ffda44" d="M0 70h512v70l-256 32L0 140Z"/><path fill="#eee" d="M0 140h512v70l-32 46l32 46v70H0v-70l32-46l-32-46Z"/><path fill="#496e2d" d="M0 210h512v92H0z"/></g>`
	reInnerSVG                    = `<mask id="circleFlagsRe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRe0)"><path fill="#0052b4" d="M64 0L0 64v160l32 32l-32 32v224l256-32l256 32V288l-32-32l32-32V64L448 0H288l-32 32l-32-32H64z"/><path fill="#d80027" d="M256 256L0 512h512z"/><path fill="#ffda44" d="M0 224v64l512-64v64z"/><path fill="#ffda44" d="M256 256L0 64V0h64zm0 0L448 0h64v64zm0 0L224 0h64z"/></g>`
	roInnerSVG                    = `<mask id="circleFlagsRo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRo0)"><path fill="#ffda44" d="M167 0h178l25.9 252.3L345 512H167l-29.8-253.4z"/><path fill="#0052b4" d="M0 0h167v512H0z"/><path fill="#d80027" d="M345 0h167v512H345z"/></g>`
	rsInnerSVG                    = `<mask id="circleFlagsRs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRs0)"><path fill="#0052b4" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#d80027" d="M0 0h512v167H0z"/><path fill="#eee" d="M0 345h512v167H0z"/><path fill="#d80027" d="M66.2 144.7v127.7c0 72.6 94.9 95 94.9 95s94.9-22.4 94.9-95V144.7z"/><path fill="#ffda44" d="M105.4 167h111.4v-44.6l-22.3 11.2l-33.4-33.4l-33.4 33.4l-22.3-11.2zm128.3 123.2l-72.3-72.4L89 290.2l23.7 23.6l48.7-48.7l48.7 48.7z"/><path fill="#eee" d="M233.7 222.6H200a22.1 22.1 0 0 0 3-11.1a22.3 22.3 0 0 0-42-10.5a22.3 22.3 0 0 0-41.9 10.5a22.1 22.1 0 0 0 3 11.1H89a23 23 0 0 0 23 22.3h-.7c0 12.3 10 22.2 22.3 22.2c0 11 7.8 20 18.1 21.9l-17.5 39.6a72.1 72.1 0 0 0 27.2 5.3a72.1 72.1 0 0 0 27.2-5.3L171.1 289c10.3-2 18.1-11 18.1-21.9c12.3 0 22.3-10 22.3-22.2h-.8a23 23 0 0 0 23-22.3z"/></g>`
	ruInnerSVG                    = `<mask id="circleFlagsRu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRu0)"><path fill="#0052b4" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#eee" d="M0 0h512v167H0z"/><path fill="#d80027" d="M0 345h512v167H0z"/></g>`
	ruBaInnerSVG                  = `<mask id="circleFlagsRuBa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuBa0)"><path fill="#eee" d="m0 160l256-32l256 32v192l-256 32L0 352Z"/><path fill="#338af3" d="M0 0h512v160H0Z"/><path fill="#496e2d" d="M0 352h512v160H0Z"/><circle cx="256" cy="256" r="64" fill="#ffda44"/></g>`
	ruCeInnerSVG                  = `<mask id="circleFlagsRuCe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuCe0)"><path fill="#d80027" d="m96 357l208-32l208 32v155H96l-32-78Z"/><path fill="#496e2d" d="M96 0h416v293l-208 32l-208-32l-32-147Z"/><path fill="#eee" d="M0 0v512h96V357h416v-64H96V0Z"/></g>`
	ruCuInnerSVG                  = `<mask id="circleFlagsRuCu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuCu0)"><path fill="#ffda44" d="M0 0h512v385l-256 64L0 385z"/><path fill="#a2001d" d="m375 97l23-24v24h23l-23 23l23 23h-23v23l-23-23l-24 23v-23h-22l22-23l-22-23h22V73zm-238 0l23-24v24h24l-24 23l24 23h-24v23l-23-23l-23 23v-23H91l23-23l-23-23h23V73zm119-50l23-23v23h23l-23 23l23 24h-23v23l-23-23l-23 23V94h-23l23-24l-23-23h23V24Z" class="fil1"/><path fill="#a2001d" d="m218 103l-35 35v28h-28l-35 35l29 29h20v38l17 17l32-32l11 11v121H0v127h512V385H282V264l11-11l32 32l17-17v-38h21l29-29l-35-35l-29 1v-29l-34-35l-29 29v23l26 26l-35 35l-35-35l26-26v-23z" class="fil1"/></g>`
	ruDaInnerSVG                  = `<mask id="circleFlagsRuDa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuDa0)"><path fill="#0052b4" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#d80027" d="M0 345h512v167H0Z"/><path fill="#6da544" d="M0 0h512v167H0Z"/></g>`
	ruKoInnerSVG                  = `<mask id="circleFlagsRuKo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuKo0)"><path fill="#6da544" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#eee" d="M0 345h512v167H0Z"/><path fill="#0052b4" d="M0 0h512v167H0Z"/></g>`
	ruTaInnerSVG                  = `<mask id="circleFlagsRuTa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuTa0)"><path fill="#eee" d="m0 224l256-32l256 32v64l-256 32L0 288Z"/><path fill="#496e2d" d="M0 0h512v224H0z"/><path fill="#d80027" d="M0 288h512v224H0z"/></g>`
	ruUdInnerSVG                  = `<mask id="circleFlagsRuUd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuUd0)"><path fill="#eee" d="M160 0h192l32 256l-32 256H160l-32-256Z"/><path fill="#333" d="M0 0h160v512H0Z"/><path fill="#d80027" d="M352 0h160v512H352ZM229 176l107 107H176l107-107v160L176 229h160L229 336Z"/></g>`
	rwInnerSVG                    = `<mask id="circleFlagsRw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRw0)"><path fill="#496e2d" d="m0 378.5l254.1-22.1L512 378.5V512H0z"/><path fill="#ffda44" d="m0 256.1l255-30.3l257 30.3v122.4H0z"/><path fill="#338af3" d="M0 0h512v256H0z"/><path fill="#ffda44" d="m289.4 150l31.3 14.6L304 195l34-6.5l4.3 34.3l23.6-25.2l23.7 25.2l4.3-34.3l34 6.5l-16.7-30.3l31.2-14.7l-31.2-14.7l16.6-30.3l-34 6.5l-4.2-34.3l-23.7 25.3l-23.6-25.3l-4.3 34.3l-34-6.5l16.7 30.3z"/></g>`
	saInnerSVG                    = `<mask id="circleFlagsSa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSa0)"><path fill="#496e2d" d="M0 0h512v512H0z"/><g fill="#eee"><path d="M144.7 306c0 18.5 15 33.5 33.4 33.5h100.2a27.8 27.8 0 0 0 27.8 27.8h33.4a27.8 27.8 0 0 0 27.8-27.8V306zm225.4-161.3v78c0 12.2-10 22.2-22.3 22.2v33.4c30.7 0 55.7-25 55.7-55.7v-77.9H370zm-239.3 78c0 12.2-10 22.2-22.3 22.2v33.4c30.7 0 55.7-25 55.7-55.7v-77.9h-33.4z"/><path d="M320 144.7h33.4v78H320zm-50 44.5a5.6 5.6 0 0 1-11.2 0v-44.5h-33.4v44.5a5.6 5.6 0 0 1-11.1 0v-44.5h-33.4v44.5a39 39 0 0 0 39 39a38.7 38.7 0 0 0 22.2-7a38.7 38.7 0 0 0 22.2 7c1.7 0 3.4-.1 5-.3a22.3 22.3 0 0 1-21.6 17v33.4c30.6 0 55.6-25 55.6-55.7v-77.9H270z"/><path d="M180.9 244.9h50v33.4h-50z"/></g></g>`
	samiInnerSVG                  = `<mask id="circleFlagsSami0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSami0)"><path fill="#ffda44" d="M199 0h32l32 256l-32 256h-32l-32-256Z"/><path fill="#496e2d" d="M167 0h32v512h-32l-32-256Z"/><path fill="#0052b4" d="M231 0h281v512H231Z"/><path fill="#d80027" d="M0 0h167v512H0Zm199 132l-32 16l32 16a92 92 0 1 1 0 184l-32 16l32 16a124 124 0 0 0 0-248z"/><path fill="#0052b4" d="M199 132a124 124 0 0 0 0 248v-32a92 92 0 1 1 0-184z"/></g>`
	sbInnerSVG                    = `<mask id="circleFlagsSb0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSb0)"><path fill="#496e2d" d="M512 512V23.6L256 255.9L23.6 512z"/><path fill="#0052b4" d="M0 0h488.4L256 256L0 488.4z"/><path fill="#ffda44" d="M488.4 0L0 488.4V512h23.6L512 23.6V0z"/><path fill="#eee" d="m107.8 89l5.5 17h18l-14.6 10.5l5.5 17L108 123l-14.5 10.6l5.5-17L84.4 106h17.8zm91.3 0l5.5 17h18L208 116.4l5.6 17L199 123l-14.5 10.6l5.6-17l-14.5-10.6h18zm-91.3 89l5.5 17h18l-14.6 10.5l5.5 17l-14.4-10.5l-14.5 10.5l5.5-17L84.4 195h17.8zm91.3 0l5.5 17h18L208 205.5l5.6 17l-14.5-10.5l-14.5 10.5l5.6-17l-14.5-10.5h18zm-45.7-44.5l5.6 17h17.8L162.4 161l5.5 17l-14.5-10.5L139 178l5.5-17l-14.5-10.5h18z"/></g>`
	scInnerSVG                    = `<mask id="circleFlagsSc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSc0)"><path fill="#0052b4" d="M0 0v332l150.9-138.5L225.2 0z"/><path fill="#ffda44" d="M273.1 253.3L512 0H225.2L0 332v80.2z"/><path fill="#d80027" d="M512 0L0 412.2v50.4L277.9 390L512 256z"/><path fill="#eee" d="M0 462.6L512 256v133.5l-223.9 78.8L0 488.4z"/><path fill="#6da544" d="m512 389.5l-512 99V512h512z"/></g>`
	sdInnerSVG                    = `<mask id="circleFlagsSd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSd0)"><path fill="#eee" d="M43.6 109.4L512 144.7v222.6L43.8 397.2z"/><path fill="#d80027" d="M0 0h512v144.7H111z"/><path fill="#333" d="M111 367.3h401V512H0z"/><path fill="#496e2d" d="M0 0v512l256-256z"/></g>`
	seInnerSVG                    = `<mask id="circleFlagsSe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSe0)"><path fill="#0052b4" d="M0 0h133.6l35.3 16.7L200.3 0H512v222.6l-22.6 31.7l22.6 35.1V512H200.3l-32-19.8l-34.7 19.8H0V289.4l22.1-33.3L0 222.6z"/><path fill="#ffda44" d="M133.6 0v222.6H0v66.8h133.6V512h66.7V289.4H512v-66.8H200.3V0z"/></g>`
	sgInnerSVG                    = `<mask id="circleFlagsSg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSg0)"><path fill="#eee" d="m0 256l257.7-51L512 256v256H0z"/><path fill="#d80027" d="M0 0h512v256H0z"/><g fill="#eee"><path d="M155.8 133.6A78 78 0 0 1 217 57.5a78.2 78.2 0 0 0-16.7-1.8a78 78 0 1 0 16.7 154a78 78 0 0 1-61.2-76.1zM256 61.2l5.5 17h18l-14.6 10.5l5.6 17L256 95.2l-14.5 10.5l5.6-17l-14.5-10.5h17.9z"/><path d="m212.6 94.6l5.6 17H236l-14.4 10.5l5.5 17l-14.5-10.5l-14.4 10.5l5.5-17l-14.5-10.5h17.9zm86.8 0l5.5 17h17.9l-14.5 10.5l5.5 17l-14.4-10.5l-14.5 10.5l5.5-17l-14.4-10.5h17.8zm-16.7 50.1l5.5 17h17.9l-14.5 10.5l5.5 17l-14.4-10.5l-14.5 10.5l5.5-17l-14.4-10.5h17.9zm-53.4 0l5.5 17h18l-14.5 10.5l5.5 17l-14.5-10.5l-14.4 10.5l5.5-17l-14.5-10.5h17.9z"/></g></g>`
	shInnerSVG                    = `<mask id="circleFlagsSh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSh0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="M288 128h160v80l-80 32l-80-32z"/><path fill="#333" d="M368 140a14 14 0 0 0-14 14h-42a14 14 0 0 0 14 14a14 14 0 0 0 14 14a14 14 0 0 0 14 14h28a14 14 0 0 0 14-14a14 14 0 0 0 14-14a14 14 0 0 0 14-14h-42a14 14 0 0 0-14-14z"/><path fill="#338af3" d="M288 208v48c0 61 80 80 80 80s80-19 80-80v-48z"/><path fill="#eee" d="M372 221c-9 0-15 8-25 1c5 12 14 9 22 10v6h-14s-14 20 0 38h14v5h-14c-9 0-9-11-9-11h-34v11h7c9 8 13 26 27 26h52c15 0 17-9 20-17c3-7 5-8 6-9v-8h-29s0 8-8 8h-12v-6h14c-11-17 0-37 0-37h-14v-16l-4-1zm-23 17c-20 6-31 26-31 26h25c-2-17 6-26 6-26zm44 6a34 34 0 0 0-2 23h24c-5-12 0-23 0-23z"/></g>`
	shAcInnerSVG                  = `<mask id="circleFlagsShAc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsShAc0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#6da544" d="m320 144l48-80l48 80z"/><circle cx="368" cy="144" r="48" fill="#acabb1"/><path fill="#338af3" d="M320 144v77c0 36 48 48 48 48s48-12 48-48v-77z"/><rect width="32" height="128" x="288" y="128" fill="#ff9811" rx="16" ry="16"/><rect width="32" height="128" x="416" y="128" fill="#ff9811" rx="16" ry="16"/><path fill="#6da544" d="m368 160l-48 67c2 11 9 19 16 26l32-45l32 45c8-7 14-15 16-26z"/></g>`
	shHlInnerSVG                  = `<mask id="circleFlagsShHl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsShHl0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="M288 128h160v80l-80 32l-80-32z"/><path fill="#333" d="M368 140a14 14 0 0 0-14 14h-42a14 14 0 0 0 14 14a14 14 0 0 0 14 14a14 14 0 0 0 14 14h28a14 14 0 0 0 14-14a14 14 0 0 0 14-14a14 14 0 0 0 14-14h-42a14 14 0 0 0-14-14z"/><path fill="#338af3" d="M288 208v48c0 61 80 80 80 80s80-19 80-80v-48z"/><path fill="#eee" d="M372 221c-9 0-15 8-25 1c5 12 14 9 22 10v6h-14s-14 20 0 38h14v5h-14c-9 0-9-11-9-11h-34v11h7c9 8 13 26 27 26h52c15 0 17-9 20-17c3-7 5-8 6-9v-8h-29s0 8-8 8h-12v-6h14c-11-17 0-37 0-37h-14v-16l-4-1zm-23 17c-20 6-31 26-31 26h25c-2-17 6-26 6-26zm44 6a34 34 0 0 0-2 23h24c-5-12 0-23 0-23z"/></g>`
	shTaInnerSVG                  = `<mask id="circleFlagsShTa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsShTa0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="M288 288v32h32v8h96v-8h32v-32h-32v8h-96v-8zm32-144l48-80l48 80z"/><circle cx="368" cy="144" r="48" fill="#338af3"/><path fill="#338af3" d="M320 144v48l48 32l48-32v-48z"/><rect width="32" height="128" x="288" y="128" fill="#ff9811" rx="16" ry="16"/><rect width="32" height="128" x="416" y="128" fill="#ff9811" rx="16" ry="16"/><path fill="#eee" d="M320 192v29c0 36 48 48 48 48s48-12 48-48v-29zm48-48l-16 48h32z"/><path fill="#338af3" d="m352 192l16 48l16-48z"/></g>`
	siInnerSVG                    = `<mask id="circleFlagsSi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSi0)"><path fill="#0052b4" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#eee" d="M0 0h512v167H0z"/><path fill="#d80027" d="M0 345h512v167H0z"/><path fill="#0052b4" d="M222.7 167v-66.8H89V167l67 82.6z"/><path fill="#eee" d="M89 167v22.2c0 51.1 66.8 66.8 66.8 66.8s66.8-15.7 66.8-66.8V167l-22.3 22.2l-44.5-33.4l-44.5 33.4z"/></g>`
	sjInnerSVG                    = `<mask id="circleFlagsSj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSj0)"><path fill="#d80027" d="M0 0h100.2l66.1 53.5L233.7 0H512v189.3L466.3 257l45.7 65.8V512H233.7l-68-50.7l-65.5 50.7H0V322.8l51.4-68.5l-51.4-65z"/><path fill="#eee" d="M100.2 0v189.3H0v33.4l24.6 33L0 289.5v33.4h100.2V512h33.4l30.6-26.3l36.1 26.3h33.4V322.8H512v-33.4l-24.6-33.7l24.6-33v-33.4H233.7V0h-33.4l-33.8 25.3L133.6 0z"/><path fill="#0052b4" d="M133.6 0v222.7H0v66.7h133.6V512h66.7V289.4H512v-66.7H200.3V0z"/></g>`
	skInnerSVG                    = `<mask id="circleFlagsSk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSk0)"><path fill="#0052b4" d="m0 160l256-32l256 32v192l-256 32L0 352z"/><path fill="#eee" d="M0 0h512v160H0z"/><path fill="#d80027" d="M0 352h512v160H0z"/><path fill="#eee" d="M64 63v217c0 104 144 137 144 137s144-33 144-137V63z"/><path fill="#d80027" d="M96 95v185a83 78 0 0 0 9 34h206a83 77 0 0 0 9-34V95z"/><path fill="#eee" d="M288 224h-64v-32h32v-32h-32v-32h-32v32h-32v32h32v32h-64v32h64v32h32v-32h64z"/><path fill="#0052b4" d="M152 359a247 231 0 0 0 56 24c12-3 34-11 56-24a123 115 0 0 0 47-45a60 56 0 0 0-34-10l-14 2a60 56 0 0 0-110 0a60 56 0 0 0-14-2c-12 0-24 4-34 10a123 115 0 0 0 47 45z"/></g>`
	slInnerSVG                    = `<mask id="circleFlagsSl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSl0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#6da544" d="M0 0h512v167H0z"/><path fill="#338af3" d="M0 345h512v167H0z"/></g>`
	smInnerSVG                    = `<mask id="circleFlagsSm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSm0)"><path fill="#338af3" d="m0 256l256-52.3L512 256v256H0z"/><path fill="#eee" d="M0 0h512v256H0z"/><path fill="#6da544" d="M357.6 176.6L256 278.3L154.4 176.6a121.7 121.7 0 0 0-20.8 68.3v33.4c0 53.6 34.6 99.2 82.7 115.8a37 37 0 0 0 4 40l36.4-29.2l36.4 29.2a37 37 0 0 0 3.9-40.5a122.6 122.6 0 0 0 81.4-115.3v-33.4c0-25.3-7.6-48.7-20.8-68.3z"/><path fill="#ffda44" d="M256 367.3c-49.1 0-89-40-89-89v-33.4a89.1 89.1 0 0 1 178 0v33.4c0 49-39.9 89-89 89z"/><path fill="#338af3" d="M311.7 278.3v-33.4a55.7 55.7 0 0 0-111.4 0v33.4l55.7 11z"/><path fill="#6da544" d="M200.3 278.3a55.7 55.7 0 0 0 111.4 0z"/><path fill="#ffda44" d="M322.8 155.8A33.4 33.4 0 0 0 267 131v-19.6h11.2V89H267V78h-22v11h-11.2v22.3H245v19.6a33.4 33.4 0 0 0-44.5 49.8v19.6h111.3v-19.6c6.8-6.1 11-15 11-24.9z"/></g>`
	snInnerSVG                    = `<mask id="circleFlagsSn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSn0)"><path fill="#ffda44" d="M144.8 0h222.4l32 260l-32 252H144.8l-32.1-256z"/><path fill="#496e2d" d="M0 0h144.8v512H0z"/><path fill="#d80027" d="M367.2 0H512v512H367.2z"/><path fill="#496e2d" d="m256.1 167l22.1 68h71.5L292 277l22 68l-57.8-42l-57.9 42l22.1-68l-57.8-42H234z"/></g>`
	soInnerSVG                    = `<mask id="circleFlagsSo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSo0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#eee" d="m256 133.6l27.6 85H373L300.7 271l27.6 85l-72.3-52.5l-72.3 52.6l27.6-85l-72.3-52.6h89.4z"/></g>`
	somalilandInnerSVG            = `<mask id="circleFlagsSomaliland0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSomaliland0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#496e2d" d="M0 0h512v167H0z"/><path fill="#d80027" d="M0 345h512v167H0z"/><path fill="#333" d="m256 200.3l13.8 42.6h44.7L278.4 269l13.8 42.6l-36.2-26.3l-36.2 26.3l13.8-42.6l-36.1-26.2h44.7z"/><g fill="#eee"><path d="M332 50v52c0 8.2-6.6 14.9-14.8 14.9V139a37.1 37.1 0 0 0 37.1-37V50h-22.2zm-159.5 52c0 8.2-6.6 14.9-14.8 14.9V139a37.1 37.1 0 0 0 37-37V50h-22.2z"/><path d="M298.7 50h22.2v52h-22.2zm-33.4 29.8a3.7 3.7 0 0 1-3.7 3.7a3.7 3.7 0 0 1-3.7-3.7V50h-22.3v29.7a3.7 3.7 0 0 1-3.7 3.7a3.7 3.7 0 0 1-3.7-3.7V50h-22.3v29.7a26 26 0 0 0 26 26a25.8 25.8 0 0 0 14.8-4.7a25.8 25.8 0 0 0 14.9 4.6c1 0 2.2 0 3.2-.2a14.9 14.9 0 0 1-14.4 11.4V139a37.1 37.1 0 0 0 37.1-37V50h-22.2v29.7z"/><path d="M206 116.9h33.3V139h-33.4z"/></g></g>`
	southOssetiaInnerSVG          = `<mask id="circleFlagsSouthOssetia0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSouthOssetia0)"><path fill="#d80027" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#eee" d="M0 0h512v167H0z"/><path fill="#ffda44" d="M0 345h512v167H0z"/></g>`
	sovietUnionInnerSVG           = `<mask id="circleFlagsSovietUnion0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSovietUnion0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m90 243l25 25l21-21l115 115l18-18l-115-115l29-29l-44-6z"/><path fill="#ffda44" d="M173 166a93 93 0 0 1 44 40c22 38 16 83-14 100c-22 13-52 7-75-15l-58 53l18 18l42-46c27 24 63 30 90 14c36-21 45-73 20-116a95 95 0 0 0-67-48zm-43-6l105-77H105l105 77l-40-124z"/><path fill="#d80027" d="m153 129l45-34h-56l45 34l-17-54z"/></g>`
	srInnerSVG                    = `<mask id="circleFlagsSr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSr0)"><path fill="#6da544" d="M0 0h512v111.3l-85.3 143.1L512 400.7V512H0V400.7l87-149L0 111.3z"/><path fill="#eee" d="M0 111.3h512V167l-41 84.7l41 93.3v55.7H0V345l44.2-86.6L0 167z"/><path fill="#a2001d" d="M0 167h512v178H0z"/><path fill="#ffda44" d="m256 167l22.1 68h71.5l-57.8 42l22 68l-57.8-42l-57.9 42l22.1-68l-57.8-42h71.5z"/></g>`
	ssInnerSVG                    = `<mask id="circleFlagsSs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSs0)"><path fill="#eee" d="M74.1 115L512 156.9v31.5l-42.4 70.3l42.4 64.2v31.5L74.1 386.8z"/><path fill="#333" d="M0 0h512v156.8H50z"/><path fill="#a2001d" d="M150.6 188.3H512v134.5H150.6z"/><path fill="#496e2d" d="M50 354.3h462V512H0z"/><path fill="#0052b4" d="M0 0v512l256-256z"/><path fill="#ffda44" d="m83.4 192.4l31.2 43.6l51.2-16.3l-31.9 43.2l31.3 43.6l-51-16.9l-31.7 43.2l.3-53.7L32 262.2L83 246z"/></g>`
	stInnerSVG                    = `<mask id="circleFlagsSt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSt0)"><path fill="#6da544" d="M0 0h512v167l-52.6 83.8L512 345v167H0l72-264.3z"/><path fill="#ffda44" d="M114.9 166.9H512v178H114.9z"/><path fill="#d80027" d="M0 0v512l256-256z"/><path fill="#333" d="m325 211.5l11.1 34H372l-29 21l11.1 34l-29-21l-28.9 21l11-34l-28.8-21H314zm111.4 0l11 34h35.8l-29 21l11.1 34l-29-21l-28.9 21l11.1-34l-29-21h35.8z"/></g>`
	suInnerSVG                    = `<mask id="circleFlagsSu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSu0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m90 243l25 25l21-21l115 115l18-18l-115-115l29-29l-44-6z"/><path fill="#ffda44" d="M173 166a93 93 0 0 1 44 40c22 38 16 83-14 100c-22 13-52 7-75-15l-58 53l18 18l42-46c27 24 63 30 90 14c36-21 45-73 20-116a95 95 0 0 0-67-48zm-43-6l105-77H105l105 77l-40-124z"/><path fill="#d80027" d="m153 129l45-34h-56l45 34l-17-54z"/></g>`
	svInnerSVG                    = `<mask id="circleFlagsSv0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSv0)"><path fill="#0052b4" d="M0 0h512v144.7l-40.5 112.6l40.5 110V512H0V367.3l42.2-114L0 144.7z"/><path fill="#eee" d="M0 144.7h512v222.6H0z"/><path fill="#ffda44" d="m204.6 267.1l51.4-89l51.4 89z"/><path fill="#6da544" d="M322.8 296.5L256 330l-66.8-33.4V252h133.6z"/><path fill="#ffda44" d="m319 182l-23.6 23.5a55.5 55.5 0 0 1-39.4 95a55.7 55.7 0 0 1-39.3-95L193 182a89 89 0 1 0 126 0z"/></g>`
	sxInnerSVG                    = `<mask id="circleFlagsSx0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSx0)"><path fill="#d80027" d="M0 0h512v256l-265 45.2z"/><path fill="#0052b4" d="M210 256h302v256H0z"/><path fill="#eee" d="M0 0v512l256-256z"/><g fill="#ffda44"><path d="M28 256a73.1 73.1 0 0 0-.2 5.6a72.3 72.3 0 1 0 144.5-5.6z"/><circle cx="100.2" cy="200.3" r="22.3"/></g><path fill="#d80027" d="M50 194.8V267c0 38.4 50.2 50 50.2 50s50-11.6 50-50v-72.3h-100z"/><path fill="#338af3" d="M100.2 294c-9.3-3.4-28-12-28-27v-50H128v50c0 15-18.6 23.6-27.8 26.9z"/><path fill="#eee" d="M111.3 244.9v-11.2l-11.1-5.5l-11.2 5.5V245l-5.5 5.5v22.3h33.4v-22.3z"/></g>`
	syInnerSVG                    = `<mask id="circleFlagsSy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSy0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#d80027" d="M0 0h512v167H0z"/><path fill="#333" d="M0 345h512v167H0z"/><path fill="#6da544" d="m153 194.8l13.8 42.5h44.7l-36.2 26.3l13.8 42.5l-36.1-26.3l-36.2 26.3l13.8-42.5l-36.2-26.3h44.7zm206 0l13.9 42.5h44.7l-36.2 26.3l13.8 42.5l-36.2-26.3l-36.1 26.3l13.8-42.5l-36.2-26.3h44.7z"/></g>`
	szInnerSVG                    = `<mask id="circleFlagsSz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSz0)"><path fill="#ffda44" d="m0 144.7l256-20.5l256 20.5V178l-37.4 79l37.4 77v33.3l-256 24.1L0 367.3V334l37.7-77.3L0 178z"/><path fill="#333" d="M0 0h512v144.7H0z"/><path fill="#0052b4" d="M0 367.3h512V512H0z"/><path fill="#a2001d" d="M0 178h512v156H0z"/><path fill="#ffda44" d="M89.2 244.9h334V267h-334zm44.6-44.6h244.8v22.3H133.8z"/><path fill="#eee" d="m256.2 189.2l-18 65.2l18 68.4c66.8 0 124-66.8 124-66.8s-57.2-66.8-124-66.8z"/><path fill="#333" d="M256.2 322.8c-66.8 0-124-66.8-124-66.8s57.2-66.8 124-66.8"/><path fill="#eee" d="M211.7 233.7h22.2v44.6h-22.2z"/><path fill="#333" d="M278.5 233.7h22.2v44.6h-22.2z"/><g fill="#0052b4" transform="translate(.2)"><circle cx="89" cy="256" r="22.3"/><circle cx="423" cy="256" r="22.3"/></g></g>`
	taInnerSVG                    = `<mask id="circleFlagsTa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTa0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="M288 288v32h32v8h96v-8h32v-32h-32v8h-96v-8zm32-144l48-80l48 80z"/><circle cx="368" cy="144" r="48" fill="#338af3"/><path fill="#338af3" d="M320 144v48l48 32l48-32v-48z"/><rect width="32" height="128" x="288" y="128" fill="#ff9811" rx="16" ry="16"/><rect width="32" height="128" x="416" y="128" fill="#ff9811" rx="16" ry="16"/><path fill="#eee" d="M320 192v29c0 36 48 48 48 48s48-12 48-48v-29zm48-48l-16 48h32z"/><path fill="#338af3" d="m352 192l16 48l16-48z"/></g>`
	tcInnerSVG                    = `<mask id="circleFlagsTc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTc0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#ffda44" d="M289.4 133.6V256c0 59.6 77.9 78 77.9 78s78-18.4 78-78V133.6h-156z"/><path fill="#ff9811" d="M356.2 178c0 12.4-10 44.6-22.3 44.6s-22.2-32.2-22.2-44.5c0-12.3 22.2-22.3 22.2-22.3s22.3 10 22.3 22.3z"/><path fill="#a2001d" d="M415.2 202.3a92.2 92.2 0 0 0 6.4-28c0-10.2-13.3-18.5-13.3-18.5s-13.4 8.3-13.4 18.6c0 6.4 2.8 19.2 6.4 28l-7.7 17.3a38.9 38.9 0 0 0 14.7 3a38.9 38.9 0 0 0 14.7-3z"/><path fill="#6da544" d="M350.6 256s-11.1 22.3-11.1 44.5H395c0-22.2-11.1-44.5-11.1-44.5l-16.7-11.1l-16.7 11.1Z"/><path fill="#d80027" d="M384 256v-5.6a16.7 16.7 0 0 0-33.4 0v5.6z"/></g>`
	tdInnerSVG                    = `<mask id="circleFlagsTd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTd0)"><path fill="#ffda44" d="M167 0h178l25.9 252.3L345 512H167l-29.8-253.4z"/><path fill="#026" d="M0 0h167v512H0z"/><path fill="#d80027" d="M345 0h167v512H345z"/></g>`
	tfInnerSVG                    = `<mask id="circleFlagsTf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTf0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0h256v256H0z"/><path fill="#0052b4" d="M0 0h75v224H0z"/><path fill="#d80027" d="M149 0h75v224h-75z"/><path fill="#eee" d="m323 158l13 21h36v71l-27-43l-37 59h18l19-34l39 70l39-70l19 34h18l-37-59l-27 43v-39h16l12-20h-28v-12h36l13-21zm43 88h-40v14h40zm76 0h-40v14h40z"/><path fill="#eee" d="m301 181l13 38l-34-23h42l-34 23zm166 0l13 38l-34-23h42l-34 23zM339 282l13 38l-34-23h42l-34 23zm90 0l13 38l-34-23h42l-34 23zm-45 34l13 38l-34-23h42l-34 23z"/></g>`
	tgInnerSVG                    = `<mask id="circleFlagsTg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTg0)"><path fill="#496e2d" d="M256 0h256v102.4L483.8 151l28.2 53.8v102.4l-30.5 50.7l30.5 51.7V512H0V409.6l34.7-51L0 307.1z"/><path fill="#ff9811" d="M0 44.5v423V256z"/><path fill="#ffda44" d="M231.7 102.4v102.4H512V102.4zM0 307.2v102.4h512V307.2H256l-128.2-26.4z"/><path fill="#d80027" d="M256 307.2V0H0v307.2z"/><path fill="#eee" d="m141.4 122.4l16.5 51h53.7L168 205l16.6 51l-43.4-31.5L98 256l16.5-51L71 173.5h53.7z"/></g>`
	thInnerSVG                    = `<mask id="circleFlagsTh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTh0)"><path fill="#d80027" d="M0 0h512v89l-79.2 163.7L512 423v89H0v-89l82.7-169.6L0 89z"/><path fill="#eee" d="M0 89h512v78l-42.6 91.2L512 345v78H0v-78l40-92.5L0 167z"/><path fill="#0052b4" d="M0 167h512v178H0z"/></g>`
	tibetInnerSVG                 = `<mask id="circleFlagsTibet0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTibet0)"><path fill="#eee" d="m12.4 481.9l491.1.6l-247-247.3z"/><path fill="#d80027" d="M45 467.4L193.1 319l63-107.5l80.2 105.1L512 362.1V256l-27.8-50l27.8-56V44.7l-260-28l-252 28l38.5 58.8L0 150v106.1l37 50.7l-37 55.4v105.2l23.6 15.8z"/><path fill="#0052b4" d="m512 362l-206.4-85.4L319 319l148.3 148.4l23 18.1l21.6-18.1zm0-105.9v-106l-255.8 106zM467.7 44.6l-60.9-17.8l-63 17.8l-87.6 211.5zm-211.5 0L214.7 27l-46.1 17.7l87.6 211.5zm-211.5 0L23.9 28.8L0 44.6V150l256.2 106.1zM0 256.1v106.1l256.2-106z"/><path fill="#ffda44" d="m256.2 256.1l63 63a89 89 0 1 0-126 0zm39 128l-38.8 19.5l-39.2-19.5a39 39 0 1 1 78 0z"/><path fill="#0052b4" d="M295.2 384.1a39 39 0 0 1-78 0"/><path fill="#ffda44" d="M0 0v44.6h512V0H0zm0 467.4V512h512v-44.6H0z"/></g>`
	tjInnerSVG                    = `<mask id="circleFlagsTj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTj0)"><path fill="#eee" d="m0 144.7l255.3-36.5L512 144.7v222.6L250.5 407L0 367.3z"/><path fill="#d80027" d="M0 0h512v144.7H0z"/><path fill="#6da544" d="M0 367.3h512V512H0z"/><path fill="#ffda44" d="M211.5 323h89v-38l-17.8 9l-26.7-26.9l-26.7 26.7l-17.8-8.9zm-51-44.6l4.2 12.7H178l-10.9 8l4.2 12.7l-10.9-8l-10.8 8l4.1-12.8l-10.7-7.9h13.4zm12.9-44.6l4.1 12.8H191l-10.8 8l4.1 12.6l-10.8-7.9l-10.9 8l4.2-12.8l-10.9-8h13.4zm36.7-33.4l4.2 12.8h13.4l-10.9 8l4.2 12.6l-11-7.7l-10.8 7.8l4.1-12.7l-10.8-7.9H206zm141.4 78l-4.2 12.7H334l10.9 8l-4.2 12.7l10.9-8l10.8 8l-4.1-12.8l10.7-7.9h-13.4zm-12.9-44.6l-4.1 12.8H321l10.8 8l-4.1 12.6l10.8-7.9l10.9 8l-4.2-12.8l10.9-8h-13.4zM302 200.4l-4.2 12.8h-13.4l10.9 8l-4.2 12.6l10.9-7.7l10.8 7.8l-4.1-12.7l10.8-7.9H306zm-46-16.6l4.1 12.7h13.5l-10.9 8l4.1 12.6l-10.8-7.8l-10.8 7.8l4.1-12.7l-10.9-7.9H252z"/></g>`
	tkInnerSVG                    = `<defs><path id="circleFlagsTk0" fill="#eee" d="m188 133.6l5.6 17h17.9L197 161l5.5 17l-14.4-10.5l-14.5 10.5l5.5-17l-14.4-10.5h17.8zM115.5 256l7 21.3h22.3l-18 13l6.8 21.4l-18-13.2l-18.1 13.2l6.9-21.3l-18.1-13.1h22.3zm0-178l7 21.2h22.3l-18 13.1l6.8 21.3l-18-13.2l-18.1 13.2l6.9-21.3l-18.1-13.1h22.3zM59.8 178l6.9 21.3H89l-18 13.2l6.9 21.2l-18.1-13.1l-18.1 13.1l7-21.2l-18.2-13.2H53z"/></defs><mask id="circleFlagsTk1"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTk1)"><path fill="#0052b4" d="M0 0h512v512H0z"/><use href="#circleFlagsTk0"/><use href="#circleFlagsTk0"/><path fill="#ffda44" d="M411.8 122.4L144.7 345h361.7v-35.8a167.4 167.4 0 0 1-97.9-152.3c0-11.7 1-23.5 3.3-34.5zm94.6 244.9H144.7v33.4h361.7z"/></g>`
	tlInnerSVG                    = `<mask id="circleFlagsTl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTl0)"><path fill="#ffda44" d="m0 0l214 251.8L0 512l418-256z"/><path fill="#d80027" d="M512 0H0l367.3 256L0 512h512z"/><path fill="#333" d="M0 0v512l256-256z"/><path fill="#eee" d="m71 197.4l39 36.8l47-25.6l-23 48.4l39 36.9l-53.2-7l-23 48.5l-9.9-52.7l-53.2-7l47.1-25.6z"/></g>`
	tmInnerSVG                    = `<mask id="circleFlagsTm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTm0)"><path fill="#496e2d" d="M0 0h66.8l67.8 32.3L200.3 0H512v512H200.3l-70.2-34.1L66.8 512H0z"/><path fill="#d80027" d="M66.8 0h133.5v512H66.8z"/><g fill="#eee"><path d="M374.4 193.1a78 78 0 0 0-31.9-92.3a78.2 78.2 0 0 1 16.3 4a78 78 0 1 1-52.6 146.6a78.1 78.1 0 0 1-15.1-7.3a78 78 0 0 0 83.3-51zm-69.8-102l-11 14.2l-16.8-6l10 14.8l-10.9 14l17.2-5l10 14.9l.6-18l17.2-5l-16.9-6z"/><path d="m334.1 137.3l-11 14l-16.8-6l10.1 14.8l-11 14.2l17.2-5l10.1 14.7l.5-17.8l17.2-5l-16.8-6zM252.5 108l-11 14.1l-16.8-6l10 14.7l-10.9 14.2l17.2-5l10 14.7l.6-17.8l17.1-5l-16.8-6zm-1.2 52.7l-11 14.2l-16.8-6l10 14.7l-10.9 14.2l17.2-5l10 14.7l.6-17.8l17.1-5l-16.8-6zm50.2 18l-11 14.2l-16.8-6l10 14.7l-10.9 14.2l17.2-5l10 14.7l.6-17.8l17.2-5l-16.9-6z"/></g><path fill="#eee" d="M117.8 134.8L95.4 118v-12l38.2-28l15.8-.1l22.3 16.7v11.5l-38.1 28.7z"/><path fill="#ff9811" d="M133.6 78h-15.8L95.4 94.5V106h38.2zm0 56.8h15.8l22.3-16.7v-12h-38.1z"/><path fill="#496e2d" d="m117.8 284.4l-22.4-16.6v-23.6l22.4-16.6h31.6l22.3 16.6v23.6l-22.3 16.6zm54.7 36.7h-9.7v-9.8h-17.1l-12.1-12l-12.1 12h-17.2v9.8h-9.7v19.5h9.7v9.7h17.2l12 12.1l12.2-12.1h17v-9.7h9.8zm0-149.7h-9.7v-9.7h-17.1l-12.1-12.1l-12.1 12h-17.2v9.8h-9.7V191h9.7v9.8h17.2l12 12l12.2-12h17V191h9.8z"/><g fill="#d80027"><path d="M122.4 244.9h22.3V267h-22.3z"/><circle cx="133.6" cy="181.2" r="11.1"/><circle cx="133.6" cy="330.8" r="11.1"/></g><path fill="#eee" d="m117.8 434l-22.4-16.8v-12l38.2-28l15.8-.1l22.3 16.7v11.5L133.6 434z"/><path fill="#ff9811" d="M133.6 377.2h-15.8l-22.4 16.5v11.5h38.2zm0 56.8h15.8l22.3-16.7v-12h-38.1z"/></g>`
	tnInnerSVG                    = `<mask id="circleFlagsTn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTn0)"><path fill="#d80027" d="M0 0h512v512H0z"/><circle cx="256" cy="256" r="122.4" fill="#eee"/><g fill="#d80027"><path d="m271 209.2l21 29l34.1-11.1l-21 29l21 28.9l-34-11.1l-21 29V267L237 256l34-11z"/><path d="M283.8 328.3a72.3 72.3 0 1 1 34.4-136a89 89 0 1 0 0 127.3a72 72 0 0 1-34.4 8.7z"/></g></g>`
	toInnerSVG                    = `<mask id="circleFlagsTo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTo0)"><path fill="#d80027" d="M0 256L256 0h256v512H0z"/><path fill="#eee" d="M0 0h256v256H0z"/><path fill="#d80027" d="M167 133.6v-33.4h-33.4v33.4h-33.4V167h33.4v33.3H167V167h33.3v-33.4z"/></g>`
	torresStraitIslandsInnerSVG   = `<mask id="circleFlagsTorresStraitIslands0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTorresStraitIslands0)"><path fill="#0052b4" d="m0 128l256-32l256 32v256l-256 32L0 384Z"/><path fill="#333" d="m0 96l256-32l256 32v32H0Z"/><path fill="#6da544" d="M0 0h512v96H0Z"/><path fill="#333" d="m0 416l256 32l256-32v-32H0Z"/><path fill="#6da544" d="M0 512h512v-96H0Z"/><path fill="#eee" d="M245 144c-106 32-101 112-67 186l-40 38l73-27v-98c24-35 66-35 90 0v98l73 27l-40-38c34-74 39-154-67-186l-11 46l-11-46zm11 83l-9 27h-28l23 17l-9 28l23-17l23 17l-9-28l24-17h-29l-9-27z"/></g>`
	trInnerSVG                    = `<mask id="circleFlagsTr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTr0)"><path fill="#d80027" d="M0 0h512v512H0z"/><g fill="#eee"><path d="m245.5 209.2l21 29l34-11.1l-21 29l21 28.9l-34-11.1l-21 29V267l-34-11.1l34-11z"/><path d="M188.2 328.3a72.3 72.3 0 1 1 34.4-136a89 89 0 1 0 0 127.3a72 72 0 0 1-34.4 8.7z"/></g></g>`
	transnistriaInnerSVG          = `<mask id="circleFlagsTransnistria0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTransnistria0)"><path fill="#a2001d" d="M0 0h512v189.5l-39 62l39 71.6V512H0V323l40.8-67L0 189.5z"/><path fill="#6da544" d="M0 189.5h512v133.6H0z"/></g>`
	ttInnerSVG                    = `<mask id="circleFlagsTt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTt0)"><path fill="#d80027" d="M0 110.2L110.2 0H512v401.8L401.8 512H0z"/><path fill="#eee" d="M110.2 0H63L0 63v47.2L401.8 512H449l63-63v-47.2z"/><path fill="#333" d="M512 512v-63L63 0H0v63l449 449z"/></g>`
	tvInnerSVG                    = `<mask id="circleFlagsTv0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTv0)"><path fill="#338af3" d="M0 256L256 0h256v512H0z"/><path fill="#eee" d="M0 0h33.4l32 16.9L100.3 0H256v33.4l-14.8 33.5l14.8 33.3v33.4l-9.3 33.7l9.3 41.5v15.7l-6.2 11.6L256 256h-47.2l-39.3-7l-35.9 7.1l-33.4-.1l-32.6-16.6L33.4 256H0V100.2l14.2-35.8L0 33.4z"/><path fill="#d80027" d="m256 224.5l-91-91h-31.4L256 256z"/><path fill="#d80027" d="M33.4 0v33.4H0v66.8h33.4V256h66.8V100.2H256V33.4H100.2V0z"/><path fill="#ffda44" d="m279.4 423l5.5 17h18l-14.6 10.5l5.6 17l-14.5-10.5l-14.5 10.5l5.6-17L256 440h17.9zm142.4-111.3l5.5 17h18l-14.5 10.5l5.5 17l-14.5-10.5l-14.5 10.5l5.6-17l-14.5-10.5h17.9zm35.7-167l5.5 17h18l-14.5 10.5l5.6 17l-14.5-10.5l-14.5 10.5l5.6-17l-14.5-10.5H452zm8.8 122.4l5.6 17h17.8l-14.4 10.5l5.5 17l-14.5-10.5l-14.4 10.6l5.5-17l-14.5-10.6h18zM305 383l16 8l12.6-12.6l-2.8 17.7l16 8.1L329 407l-2.8 17.7l-8.1-16l-17.7 2.8l12.7-12.6zm0-66.8l16 8.1l12.6-12.6l-2.8 17.6l16 8.1l-17.7 2.8l-2.8 17.7l-8.1-16l-17.7 2.9l12.7-12.7zm55.7 42.8l16 8.1l12.6-12.6l-2.8 17.6l16 8.1l-17.8 2.8l-2.7 17.7l-8.2-16l-17.6 2.9l12.6-12.7zm0-143l16 8.1l12.6-12.6l-2.8 17.6l16 8.2l-17.8 2.8l-2.7 17.6l-8.2-16l-17.6 2.9l12.6-12.7zm54-1.7l15.9 8.1l12.6-12.6l-2.8 17.6l16 8.1l-17.7 2.8l-2.8 17.7l-8.1-16l-17.8 3l12.7-12.7z"/><path fill="#0052b4" d="M180.8 133.6H256v75.2zm-47.2 47.2v75.3l75.2-.1z"/></g>`
	twInnerSVG                    = `<mask id="circleFlagsTw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTw0)"><path fill="#d80027" d="M0 256L256 0h256v512H0z"/><path fill="#0052b4" d="M256 256V0H0v256z"/><path fill="#eee" d="m222.6 149.8l-31.3 14.7l16.7 30.3l-34-6.5l-4.3 34.3l-23.6-25.2l-23.7 25.2l-4.3-34.3l-34 6.5l16.7-30.3l-31.2-14.7l31.2-14.7l-16.6-30.3l34 6.5l4.2-34.3l23.7 25.3L169.7 77l4.3 34.3l34-6.5l-16.7 30.3z"/><circle cx="146.1" cy="149.8" r="47.7" fill="#0052b4"/><circle cx="146.1" cy="149.8" r="41.5" fill="#eee"/></g>`
	tzInnerSVG                    = `<mask id="circleFlagsTz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTz0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#ffda44" d="M399 0L167 167L0 399v45l68 68h45l232-167l167-232V68L444 0Z"/><path fill="#333" d="M444 0L0 444v68h68L512 68V0Z"/><path fill="#338af3" d="m113 512l399-399v399z"/><path fill="#6da544" d="M0 399V0h399Z"/></g>`
	uaInnerSVG                    = `<mask id="circleFlagsUa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUa0)"><path fill="#ffda44" d="m0 256l258-39.4L512 256v256H0z"/><path fill="#338af3" d="M0 0h512v256H0z"/></g>`
	ugInnerSVG                    = `<mask id="circleFlagsUg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUg0)"><path fill="#ffda44" d="M0 85v342l256 31l256-31V85L256 53Z"/><path fill="#333" d="M0 0h512v85H0z"/><path fill="#d80027" d="M0 171h512v85l-256 31L0 256Z"/><path fill="#333" d="M0 256h512v85H0z"/><path fill="#d80027" d="M0 427h512v85H0z"/><circle cx="256" cy="256" r="85.3" fill="#eee"/><path fill="#333" d="m287 260l-31-13l8-26a17 17 0 0 0-5-16l8-8a28 28 0 0 0-19-8a28 28 0 0 0-20 8l8 8a17 17 0 0 0-5 12l1 5l-12 12h21s-9 13-14 24c-4 11 0 25 12 30l6 2l11 10v12l-11 11h22v-23l10-10h21a22 22 0 0 0-11-30z"/></g>`
	ukInnerSVG                    = `<mask id="circleFlagsUk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUk0)"><path fill="#eee" d="m0 0l8 22l-8 23v23l32 54l-32 54v32l32 48l-32 48v32l32 54l-32 54v68l22-8l23 8h23l54-32l54 32h32l48-32l48 32h32l54-32l54 32h68l-8-22l8-23v-23l-32-54l32-54v-32l-32-48l32-48v-32l-32-54l32-54V0l-22 8l-23-8h-23l-54 32l-54-32h-32l-48 32l-48-32h-32l-54 32L68 0H0z"/><path fill="#0052b4" d="M336 0v108L444 0Zm176 68L404 176h108zM0 176h108L0 68ZM68 0l108 108V0Zm108 512V404L68 512ZM0 444l108-108H0Zm512-108H404l108 108Zm-68 176L336 404v108z"/><path fill="#d80027" d="M0 0v45l131 131h45L0 0zm208 0v208H0v96h208v208h96V304h208v-96H304V0h-96zm259 0L336 131v45L512 0h-45zM176 336L0 512h45l131-131v-45zm160 0l176 176v-45L381 336h-45z"/></g>`
	umInnerSVG                    = `<mask id="circleFlagsUm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUm0)"><path fill="#eee" d="M256 0h256v64l-32 32l32 32v64l-32 32l32 32v64l-32 32l32 32v64l-256 32L0 448v-64l32-32l-32-32v-64z"/><path fill="#d80027" d="M224 64h288v64H224Zm0 128h288v64H256ZM0 320h512v64H0Zm0 128h512v64H0Z"/><path fill="#0052b4" d="M0 0h256v256H0Z"/><path fill="#eee" d="m187 243l57-41h-70l57 41l-22-67zm-81 0l57-41H93l57 41l-22-67zm-81 0l57-41H12l57 41l-22-67zm162-81l57-41h-70l57 41l-22-67zm-81 0l57-41H93l57 41l-22-67zm-81 0l57-41H12l57 41l-22-67Zm162-82l57-41h-70l57 41l-22-67Zm-81 0l57-41H93l57 41l-22-67zm-81 0l57-41H12l57 41l-22-67Z"/></g>`
	unInnerSVG                    = `<mask id="circleFlagsUn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUn0)"><path fill="#338af3" d="M0 0h512v512H0z"/><circle cx="256" cy="256" r="145" fill="#eee"/><circle cx="256" cy="256" r="111" fill="#338af3"/><path fill="#338af3" d="M76 76h360L256 256z"/><circle cx="256" cy="256" r="89" fill="#eee"/><circle cx="256" cy="256" r="69" fill="#338af3"/><path fill="#eee" d="M246 178h20v156h-20z"/><path fill="#eee" d="M334 246v20H178v-20z"/><path fill="#eee" d="m304 193.7l14.2 14.2l-110.3 110.3l-14.2-14.1z"/><path fill="#eee" d="m318.2 304l-14.1 14.2l-110.4-110.3l14.2-14.2z"/><circle cx="256" cy="256" r="44" fill="#eee"/><circle cx="256" cy="256" r="22" fill="#338af3"/><ellipse cx="256" cy="412" fill="#eee" rx="44" ry="40"/><path fill="#338af3" d="m256 407l-78 63h156z"/></g>`
	unitedNationsInnerSVG         = `<mask id="circleFlagsUnitedNations0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUnitedNations0)"><path fill="#338af3" d="M0 0h512v512H0z"/><circle cx="256" cy="256" r="145" fill="#eee"/><circle cx="256" cy="256" r="111" fill="#338af3"/><path fill="#338af3" d="M76 76h360L256 256z"/><circle cx="256" cy="256" r="89" fill="#eee"/><circle cx="256" cy="256" r="69" fill="#338af3"/><path fill="#eee" d="M246 178h20v156h-20z"/><path fill="#eee" d="M334 246v20H178v-20z"/><path fill="#eee" d="m304 193.7l14.2 14.2l-110.3 110.3l-14.2-14.1z"/><path fill="#eee" d="m318.2 304l-14.1 14.2l-110.4-110.3l14.2-14.2z"/><circle cx="256" cy="256" r="44" fill="#eee"/><circle cx="256" cy="256" r="22" fill="#338af3"/><ellipse cx="256" cy="412" fill="#eee" rx="44" ry="40"/><path fill="#338af3" d="m256 407l-78 63h156z"/></g>`
	usInnerSVG                    = `<mask id="circleFlagsUs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUs0)"><path fill="#eee" d="M256 0h256v64l-32 32l32 32v64l-32 32l32 32v64l-32 32l32 32v64l-256 32L0 448v-64l32-32l-32-32v-64z"/><path fill="#d80027" d="M224 64h288v64H224Zm0 128h288v64H256ZM0 320h512v64H0Zm0 128h512v64H0Z"/><path fill="#0052b4" d="M0 0h256v256H0Z"/><path fill="#eee" d="m187 243l57-41h-70l57 41l-22-67zm-81 0l57-41H93l57 41l-22-67zm-81 0l57-41H12l57 41l-22-67zm162-81l57-41h-70l57 41l-22-67zm-81 0l57-41H93l57 41l-22-67zm-81 0l57-41H12l57 41l-22-67Zm162-82l57-41h-70l57 41l-22-67Zm-81 0l57-41H93l57 41l-22-67zm-81 0l57-41H12l57 41l-22-67Z"/></g>`
	usAkInnerSVG                  = `<mask id="circleFlagsUsAk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsAk0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m336.8 199.7l104-75.1h-128l104 75.1l-40-122.1zM57 231.7l52-37.4H45l52 37.4l-20-60.9zm85.5 31.7l52-37.4h-64l52 37.4l-20-60.9zM185 309l52-37.4h-64l52 37.4l-20-60.9zm43 47.6l52-37.4h-64l52 37.4l-20-60.9zm-5.6 67.1l52-37.4h-64l52 37.4l-20-60.9zM356 402.2l52-37.4h-64l52 37.4l-20-60.9zm-37.4 53.2l52-37.4h-64l52 37.4l-20-60.9z"/></g>`
	usAlInnerSVG                  = `<mask id="circleFlagsUsAl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsAl0)"><path fill="#eee" d="M0 68L68 0h376l68 68v376l-68 68H68L0 444Z"/><path fill="#d80027" d="M0 0v68l188 188L0 444v68h68l188-188l188 188h68v-68L324 256L512 68V0h-68L256 188L68 0Z"/></g>`
	usArInnerSVG                  = `<mask id="circleFlagsUsAr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsAr0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#0052b4" d="M256 70L5 256l251 186l251-186Z"/><path fill="#eee" d="M256 130L85 256l171 126l171-126Zm-13.4-5.7l34.7-25h-42.7l34.7 25L256 83.7ZM46.5 277.4l27.2-33.1l-41.4 11.1l40.1 15L49 234.6Zm419 0l-27.2-33.1l41.4 11.1l-40.1 15l23.4-35.8zm-145.7-104l27.2-33.1l-41.4 11.1l40.1 15l-23.4-35.8zm70.3 52l27.2-33.1l-41.4 11.1l40.1 15l-23.4-35.8zm-197.9-52L165 140.3l41.4 11.1l-40.1 15l23.4-35.8zm-70.3 52l-27.2-33.1l41.4 11.1l-40.1 15l23.4-35.8zm-13 104l27.2-33.1l-41.5 11.1l40.2 15l-23.4-35.8zm70.3 52l27.2-33.1l-41.4 11.1l40.1 15l-23.4-35.8zm63.5 46.9l34.7-24.9h-42.7l34.7 24.9l-13.3-40.6zm160.4-98.9l-27.2-33.1l41.5 11.1l-40.2 15l23.4-35.8zm-70.3 52l-27.2-33.1l41.4 11.1l-40.1 15l23.4-35.8z"/><path fill="#0052b4" d="m242.7 204.7l34.7-25h-42.7l34.7 25l-13.3-40.7zm-39.2 103l34.7-25h-42.7l34.7 25l-13.3-40.7zm78.2 0l34.7-25h-42.7l34.7 25l-13.3-40.7zm-39.1-4.4l34.7 25h-42.7l34.7-25L256 344ZM160 223h192v32H160z"/></g>`
	usAzInnerSVG                  = `<mask id="circleFlagsUsAz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsAz0)"><path fill="#d80027" d="M71 0L0 34v91l16 35l-16 36v60l256 64l256-64v-60l-16-36l16-35V34L441 0h-85l-32 16l-32-16h-72l-32 16l-32-16Z"/><path fill="#0052b4" d="M0 256h512v256H0Z"/><path fill="#ffda44" d="m292 0l-36 256L356 0h-64zm-36 256L512 34V0h-71L256 256zm0 0l256-60v-71L256 256zm0 0L0 125v71l256 60zm0 0L71 0H0v35l256 221zm0 0L220 0h-64l100 256z"/><path fill="#ff9811" d="m256 128l83 256l-217-159h268L173 384Z"/></g>`
	usCaInnerSVG                  = `<mask id="circleFlagsUsCa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsCa0)"><path fill="#eee" d="M0 0h512v416l-256 32L0 416Z"/><path fill="#6da544" d="m396 280l-67-8l-34 3l-24-2l-54-2l-93 13l-8 12h299z"/><path fill="#584528" d="m389 273l-10-33l-9-44l-23-23l-32-13l-50 3l-41-8l-50 24l-29 9l-9 20l-7 15l11 19l10 2l16-1l6-7l28-8l-12 31l-11 14l-19 9l13 3l15 1l19-12l26-24l11 3l14 13l-16 12l16 1l10-3l9-16l-4-19l11 1l10 3l1 23l-14 8l1 6h26l34-30l15 12l12 4l-6 8l1 5h17l10-7zM84 340h178v32H84z"/><path fill="#d80027" d="M0 416h512v96H0Z"/><path fill="#584528" d="M288 340h140v32H288z"/><path fill="#d80027" d="m71 182l69-50H55l69 50l-26-81z"/></g>`
	usCoInnerSVG                  = `<mask id="circleFlagsUsCo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsCo0)"><path fill="#0052b4" d="M0 0h512v167l-64 89l64 89v167H0V345l64-89l-64-89Z"/><path fill="#eee" d="M0 167h512v178H0z"/><path fill="#d80027" d="M344.3 299.8A128 128 0 0 1 201.8 382A128 128 0 0 1 96 256a128 128 0 0 1 105.8-126a128 128 0 0 1 142.5 82.2L224 256z"/><circle cx="224" cy="256" r="64" fill="#ffda44"/></g>`
	usDcInnerSVG                  = `<mask id="circleFlagsUsDc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsDc0)"><path fill="#eee" d="M0 0h512v186l-64 48l64 48v57l-64 48l64 48v77H0v-77l64-48l-64-48v-57l64-48l-64-48Z"/><path fill="#d80027" d="M0 186h512v96H0zm0 153h512v96H0zm224-181l83-60H205l83 60l-32-98Zm118 0l83-60H323l83 60l-32-98Zm-236 0l83-60H87l83 60l-32-98Z"/></g>`
	usFlInnerSVG                  = `<mask id="circleFlagsUsFl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsFl0)"><path fill="#eee" d="M0 68L68 0h376l68 68v376l-68 68H68L0 444Z"/><path fill="#d80027" d="M0 0v68l188 188L0 444v68h68l188-188l188 188h68v-68L324 256L512 68V0h-68L256 188L68 0Z"/><circle cx="256" cy="256" r="96" fill="#ff9811"/><circle cx="256" cy="256" r="64" fill="#6da544"/></g>`
	usGaInnerSVG                  = `<mask id="circleFlagsUsGa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsGa0)"><path fill="#d80027" d="M512 0v171l-64 85l64 85v171H0V341L341 0Z"/><path fill="#eee" d="M341 171h171v170H341l-64-85z"/><path fill="#0052b4" d="M0 0h341v341H0Z"/><path fill="#ffda44" d="M184 133a40 40 0 0 0-40 40v60h80v-60a40 40 0 0 0-40-40z"/><path fill="#eee" d="m167 105l52-37h-64l52 37l-20-61zm101 62l37 52v-64l-37 52l61-20zm-62 101l-52 37h64l-52-37l20 61zm-101-61l-37-52v64l37-52l-61 20zm125-92l63 10l-45-45l10 63l29-57zm28 115l-10 63l45-45l-63 10l57 29zm-115 28l-63-10l45 45l-10-63l-29 57Zm-28-115l10-63l-45 45l63-10l-57-29Z"/></g>`
	usHiInnerSVG                  = `<mask id="circleFlagsUsHi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsHi0)"><path fill="#eee" d="M0 256V96l32-32L0 32V0h32l32 32L96 0h416v64l-32 64l32 64v64l-32 64l32 64v64l-256 32L0 448v-64l32-64z"/><path fill="#0052b4" d="m173 128l83 83v-19h256v-64L384 96l-128 32Zm-45 45v83h83zM0 320v64h512v-64l-256-32Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160H0v64h512v-64H96V96h160v32h256V64H256V32H96V0H32zm224 256v-31l-97-97h-31l128 128zM0 448v64h512v-64H0z"/></g>`
	usInInnerSVG                  = `<mask id="circleFlagsUsIn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsIn0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#ffda44" d="M224 361h64l-52 37.5l20-61l20 61zm-79.2-32.8h64l-52 37.5l20-61l20 61zM112 249h64l-52 37.5l20-61l20 61zm144 145l-12 8.6V451l12 21l12-21v-48.4zm6.7-207.7c-14.3 5.2-18.9 17.4-17.5 30.5a20.7 20.7 0 0 1-6.3-11.5a26 26 0 0 0-6 15.9a27.8 27.8 0 0 0 8.3 19.8H216v16h16l12 31h-12v16h12v44.4l12-36.6l12 36.6V304h12v-16h-12l12-31h16v-16h-23.3c7.6-8.5 8.6-23.3 1.9-34c-.3 6.4-3.5 11.2-5.6 14.8c-.9-11.8-11.2-18.8-6.3-35.5ZM224 63.5h64L236 101l20-61l20 61zm92.1 29.9h64l-52 37.5l20-61l20 61zm63.5 63.5h64l-52 37.5l20-61l20 61zm29.9 92.1h64l-52 37.5l20-61l20 61zm-29.9 92.1h64l-52 37.5l20-61l20 61zm-63.5 63.5h64l-52 37.5l20-61l20 61zm-184.2 0h64l-52 37.5l20-61l20 61zm-63.5-63.5h64l-52 37.5l20-61l20 61zM38.5 249h64l-52 37.5l20-61l20 61zm29.9-92.1h64l-52 37.5l20-61l20 61zm63.5-63.5h64l-52 37.5l20-61l20 61zM224 137h64l-52 37.5l20-61l20 61zm112 112h64l-52 37.5l20-61l20 61zm-32.8 79.2h64l-52 37.5l20-61l20 61z"/></g>`
	usMoInnerSVG                  = `<mask id="circleFlagsUsMo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsMo0)"><path fill="#eee" d="m0 170l256-64l256 64v172l-256 64L0 342Z"/><path fill="#d80027" d="M0 0h512v170H0Z"/><path fill="#0052b4" d="M0 342h512v170H0Z"/><circle cx="256" cy="256" r="160" fill="#0052b4"/><circle cx="256" cy="256" r="104" fill="#eee"/><path fill="#eee" d="m240 148.4l41.6-30h-51.2l41.6 30l-16-48.8zM363.6 240l30 41.6v-51.2l-30 41.6l48.8-16zM272 363.6l-41.6 30h51.2l-41.6-30l16 48.8zM148.4 272l-30-41.6v51.2l30-41.6l-48.8 16zm147.5-117.2l51.1-5.2l-44.4-25.6l21.1 46.8l10.5-50.3zm61.3 141.1l5.2 51.1l25.6-44.4l-46.8 21.1l50.2 10.5zm-141.1 61.3l-51.1 5.2l44.4 25.6l-21.1-46.8l-10.5 50.2zm-61.3-141.1l-5.2-51.1l-25.6 44.4l46.8-21.1l-50.2-10.5Zm186.4-27.8l46.8 21.1l-25.6-44.4l-5.2 51.1l34.3-38.3zm-17.5 152.9L302.6 388l44.4-25.6l-51.1-5.2l38.3 34.2zm-152.9-17.5L124 302.6l25.6 44.4l5.2-51.1l-34.3 38.3zm17.5-152.9l21.1-46.8l-44.4 25.6l51.1 5.2l-38.3-34.2z"/><path fill="#338af3" d="M256 168a88 88 0 0 0-74.9 42H331a88 88 0 0 0-75-42z"/><circle cx="256" cy="280" r="56" fill="#ff9811"/><path fill="#eee" d="M288 280a32 32 0 0 1-32 32l-16-32l16-32a32 32 0 0 1 32 32z"/><path fill="#0052b4" d="M256 248a32 32 0 0 0-32 32l16 16l16-16z"/><path fill="#d80027" d="M256 312a32 32 0 0 1-32-32h32z"/></g>`
	usMsInnerSVG                  = `<mask id="circleFlagsUsMs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsMs0)"><path fill="#d80027" d="M0 0h104l152 64L408 0h104v512H408l-152-64l-152 64H0Z"/><path fill="#ff9811" d="M104 0h24l128 64L384 0h24v512h-24l-128-64l-128 64h-24z"/><path fill="#0052b4" d="M128 0h256v512H128Z"/><path fill="#eee" d="m198 155l3 16l-14 9l16 1l4 16l7-15l16 2l-12-11l6-15l-14 8l-12-11zm116 0l-12 11l-14-8l7 15l-13 11l17-2l6 15l4-16l16-1l-14-8l3-17zm-134 32l-8 14l-17-3l11 12l-8 14l15-6l11 12l-2-17l15-6l-16-4l-1-16zm153 0l-2 16l-16 4l15 6l-2 17l11-12l15 6l-8-14l11-12l-16 3l-8-14zm-77 21a16 16 0 0 0-16 16a16 16 0 0 0 2 8a16 16 0 0 0-6-6a16 16 0 0 0-8-2a16 16 0 0 0-14 8a16 16 0 0 0 6 22a16 16 0 0 0 8 2a16 16 0 0 0-8 2a16 16 0 0 0-6 22a16 16 0 0 0 22 6a16 16 0 0 0 6-6a16 16 0 0 0-2 8a16 16 0 0 0 16 16a16 16 0 0 0 16-16a16 16 0 0 0-2-8a16 16 0 0 0 6 6a16 16 0 0 0 22-6a16 16 0 0 0-6-22a16 16 0 0 0-8-2a16 16 0 0 0 8-2a16 16 0 0 0 6-22a16 16 0 0 0-13-8a16 16 0 0 0-9 2a16 16 0 0 0-6 6a16 16 0 0 0 2-8a16 16 0 0 0-16-16zm-100 27v16l-16 5l16 5v16l9-13l16 5l-10-13l10-13l-16 5l-9-13zm201 0l-10 13l-15-5l9 13l-9 13l15-5l10 13v-16l15-5l-15-5v-16zm-173 47l-11 12l-15-6l8 14l-11 12l17-3l8 14l1-16l16-4l-15-6l2-17zm144 0l2 17l-15 6l16 4l1 16l9-14l16 3l-11-12l8-14l-15 7l-11-13zm-130 40l-11 12a104 104 0 0 0 138 0l-11-12a88 88 0 0 1-116 0z"/><path fill="#ff9811" d="m256 140l-5 15h-16l13 10l-5 15l13-9l13 9l-5-15l13-10h-16l-5-15zm0 100a16 16 0 0 0-16 16a16 16 0 0 0 16 16a16 16 0 0 0 16-16a16 16 0 0 0-16-16z"/></g>`
	usNcInnerSVG                  = `<mask id="circleFlagsUsNc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsNc0)"><path fill="#0052b4" d="M0 0h232l64 256l-64 256H0Z"/><path fill="#d80027" d="M232 0h280v256l-140 64l-140-64Z"/><path fill="#eee" d="M232 256h280v256H232Z"/><path fill="#0052b4" d="M0 0h232v512H0z"/><path fill="#eee" d="m91 294l65-47H76l65 47l-25-76z"/><path fill="#ffda44" d="M24 224h40v64H24zm144 0h40v64h-40zM64 92v16H24v32h40v-16h104v16h40v-32h-40V92H64zm0 328v-16H24v-32h40v16h104v-16h40v32h-40v16z"/></g>`
	usNmInnerSVG                  = `<mask id="circleFlagsUsNm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsNm0)"><path fill="#ffda44" d="M0 0h512v512H0Z"/><path fill="#d80027" d="M256 56c-7 0-12 5-12 12v128h-24v-92a12 12 0 0 0-24 0v92h-92a12 12 0 0 0 0 24h92v24H68a12 12 0 0 0 0 24h128v24h-92a12 12 0 0 0 0 24h92v92c0 7 5 12 12 12s12-5 12-12v-92h24v128c0 7 5 12 12 12s12-5 12-12V316h24v92c0 7 5 12 12 12s12-5 12-12v-92h92c7 0 12-5 12-12s-5-12-12-12h-92v-24h128c7 0 12-5 12-12s-5-12-12-12H316v-24h92c7 0 12-5 12-12s-5-12-12-12h-92v-92a12 12 0 0 0-24 0v92h-24V68c0-7-5-12-12-12z"/><circle cx="256" cy="256" r="88" fill="#d80027"/><circle cx="256" cy="256" r="64" fill="#ffda44"/></g>`
	usRiInnerSVG                  = `<mask id="circleFlagsUsRi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsRi0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#0052b4" d="M204 296v16h-40v32h40v-16h104v16h40v-32h-40v-16z"/><path fill="#ffda44" d="m236 107l52-38h-64l52 38l-20-61zm170 129l37 52v-64l-37 52l61-20zM276 406l-52 37h64l-52-37l20 61zM107 276l-38-52v64l38-52l-61 20zm206-159l64-7l-55-32l26 59l13-63zm83 196l6 64l32-55l-59 26l63 13zm-197 83l-64 6l55 32l-26-59l-13 63zm-82-197l-7-64l-32 55l59-26l-63-13Zm259-35l58 26l-32-55l-7 64l43-48zm-28 212l-26 58l55-32l-64-7l48 43zm-211-28l-59-26l32 55l6-64l-42 48zm27-211l26-59l-55 32l64 6l-48-42zm92-3a24 24 0 0 0-24 24a24 24 0 0 0 16 23v17h-16v16h16v39c-20-2-30-15-33-18l11-11h-32v32l10-10a68 68 0 0 0 104 0l10 10v-32h-32l11 11c-3 3-13 16-33 18v-39h16v-16h-16v-17a24 24 0 0 0 16-23a24 24 0 0 0-24-24zm0 16a8 8 0 0 1 8 8a8 8 0 0 1-8 8a8 8 0 0 1-8-8a8 8 0 0 1 8-8z"/></g>`
	usTnInnerSVG                  = `<mask id="circleFlagsUsTn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsTn0)"><path fill="#0052b4" d="M448 0h64v512h-64l-16-256Z"/><path fill="#eee" d="M416 0h32v512h-32l-16-256Z"/><path fill="#d80027" d="M0 0h416v512H0z"/><circle cx="208" cy="256" r="160" fill="#eee"/><circle cx="208" cy="256" r="128" fill="#0052b4"/><path fill="#eee" d="m145 284l60 83V265l-60 83l98-32zm15-128l-22 100l76-68l-102 11l89 51zm147 49l-42 93l-21-100l75 69l-102-11z"/></g>`
	usTxInnerSVG                  = `<mask id="circleFlagsUsTx0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsTx0)"><path fill="#0052b4" d="M0 0h167l64 256l-64 256H0Z"/><path fill="#eee" d="m43.5 317l104-75h-128l104 75l-40-122zM167 0h345v256l-173 64l-172-64Z"/><path fill="#d80027" d="M167 256h345v256H167z"/></g>`
	uyInnerSVG                    = `<mask id="circleFlagsUy0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUy0)"><path fill="#338af3" d="M0 256L256 0h256v55.7l-20.7 34.5l20.7 32.2v66.8l-21.2 32.7L512 256v66.8l-24 31.7l24 35.1v66.7l-259.1 28.3L0 456.3v-66.7l27.1-33.3L0 322.8z"/><path fill="#eee" d="M256 256h256v-66.8H236.9zm-19.1-133.6H512V55.7H236.9zM512 512v-55.7H0V512zM0 389.6h512v-66.8H0z"/><path fill="#eee" d="M0 0h256v256H0z"/><path fill="#ffda44" d="m222.6 149.8l-31.3 14.7l16.7 30.3l-34-6.5l-4.3 34.3l-23.6-25.2l-23.7 25.2l-4.3-34.3l-33.9 6.5l16.6-30.3l-31.2-14.7l31.2-14.7l-16.6-30.3l34 6.5l4.2-34.3l23.7 25.3L169.7 77l4.3 34.3l34-6.5l-16.7 30.3z"/></g>`
	uzInnerSVG                    = `<mask id="circleFlagsUz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUz0)"><path fill="#d80027" d="m0 178l254.2-22L512 178v22.3l-40.2 54.1l40.2 57.3V334l-254 23.4L0 334v-22.3l36.7-59.4l-36.7-52z"/><path fill="#338af3" d="M0 0h512v178H0z"/><path fill="#eee" d="M0 200.3h512v111.4H0z"/><path fill="#6da544" d="M0 334h512v178H0z"/><path fill="#eee" d="M117.2 105.7a50 50 0 0 1 39.3-48.9a50.2 50.2 0 0 0-10.7-1.1a50 50 0 1 0 10.7 99c-22.5-5-39.3-25-39.3-49zm69 22.8l3.3 10.4h11l-9 6.5l3.5 10.4l-9-6.4l-8.7 6.4l3.4-10.4l-9-6.5h11zm35 0l3.4 10.4h11l-9 6.5l3.4 10.4l-8.8-6.4l-9 6.4l3.5-10.4l-9-6.5h11zm35 0l3.4 10.4h11l-9 6.5l3.5 10.4l-9-6.4l-8.8 6.4l3.4-10.4l-9-6.5h11zm35 0l3.4 10.4h11l-9 6.5l3.5 10.4l-9-6.4l-8.8 6.4l3.4-10.4l-9-6.5h11zm35 0l3.4 10.4h11l-9 6.5l3.5 10.4l-9-6.4l-8.8 6.4l3.4-10.4l-8.8-6.5h11zm-105-36.4l3.4 10.4h11l-9 6.5l3.4 10.4l-8.8-6.5l-9 6.5l3.5-10.4l-9-6.5h11zm35 0l3.4 10.4h11l-9 6.5l3.5 10.4l-9-6.5l-8.8 6.5l3.4-10.4l-9-6.5h11zm35 0l3.4 10.4h11l-9 6.5l3.5 10.4l-9-6.5l-8.8 6.5l3.4-10.4l-9-6.5h11zm35 0l3.4 10.4h11l-9 6.5l3.5 10.4l-9-6.5l-8.8 6.5l3.4-10.4l-8.8-6.5h11zm-70-36.4l3.4 10.4h11l-9 6.4l3.6 10.5l-9-6.5l-8.8 6.5l3.4-10.5l-9-6.4h11zm35 0l3.4 10.4h11l-9 6.4l3.6 10.5l-9-6.5l-8.8 6.5l3.4-10.5l-9-6.4h11zm35 0l3.4 10.4h11l-9 6.4l3.6 10.5l-9-6.5l-8.8 6.5l3.4-10.5l-8.8-6.4h11z"/></g>`
	vaInnerSVG                    = `<mask id="circleFlagsVa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVa0)"><path fill="#ffda44" d="M0 0h256l51.7 254.7L256 512H0z"/><path fill="#eee" d="M256 0h256v512H256z"/><path fill="#acabb1" d="m354 222.8l48.1 63.6A33.4 33.4 0 1 0 420 273l-75-99.2l-17.7 13.4l-26.7 20.2l26.9 35.5l26.6-20.1zm69.3 73.1a11.1 11.1 0 1 1 13.4 17.8a11.1 11.1 0 0 1-13.4-17.8z"/><path fill="#ffda44" d="m436.6 242.9l26.8-35.5l-26.6-20.2l-17.8-13.4l-75 99.2a33.4 33.4 0 1 0 17.8 13.4l48-63.6l26.8 20zm-93.8 68.6a11.1 11.1 0 1 1-17.8-13.4a11.1 11.1 0 0 1 17.8 13.4z"/></g>`
	vcInnerSVG                    = `<mask id="circleFlagsVc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVc0)"><path fill="#ffda44" d="M378.5 0H133.7l-22.3 256l22.3 256h244.8l22.3-256z"/><path fill="#338af3" d="M133.7 512V0H0v512z"/><path fill="#6da544" d="M512 0H378.5v512H512zM200.4 322.8L156 256l44.5-66.8l44.7 66.8zm111.4 0L267.1 256l44.6-66.8l44.5 66.8zm-55.7 89L211.6 345l44.5-66.7l44.5 66.7z"/></g>`
	veInnerSVG                    = `<mask id="circleFlagsVe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVe0)"><path fill="#0052b4" d="m0 144.7l255.3-36.5L512 144.7v222.6L250.5 407L0 367.3z"/><path fill="#ffda44" d="M0 0h512v144.7H0z"/><path fill="#d80027" d="M0 367.3h512V512H0z"/><path fill="#eee" d="M443.4 306.4L429.8 317l6 16.1l-14.3-9.6l-13.5 10.7l4.7-16.5l-14.2-9.6l17.1-.6l4.7-16.5l6 16.1zm-34.7-60l-9 14.5l11 13.2L394 270l-9 14.6l-1.3-17l-16.6-4.3l15.9-6.4l-1.2-17l11 13zm-53-44.5l-3.6 16.8l14.9 8.4l-17 1.8l-3.6 16.8l-7-15.7l-17 1.8l12.7-11.5l-7-15.6l14.8 8.6zm-65-23.7l2.3 17l17 3l-15.5 7.5l2.4 17l-12-12.4l-15.4 7.6l8-15.2l-11.8-12.3l16.9 3zm-69.3 0l8 15.1l17-3l-12 12.4l8 15.2l-15.4-7.6l-11.9 12.4l2.4-17l-15.4-7.5l16.9-3zm-65 23.7l12.7 11.5l14.8-8.6l-7 15.7l12.8 11.4l-17-1.8l-7 15.7l-3.7-16.7l-17-1.8l14.8-8.5zm-53.1 44.5l15.9 6.4l11-13l-1.2 17l16 6.4l-16.7 4.2l-1.2 17L118 270l-16.7 4.2l11-13.2zm-34.7 60l17.2.5l5.9-16l4.7 16.4l17.1.6l-14.2 9.6l4.7 16.6l-13.5-10.6l-14.2 9.6l5.9-16z"/></g>`
	vgInnerSVG                    = `<mask id="circleFlagsVg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVg0)"><path fill="#0052b4" d="M256 0h256v512H0V256Z"/><path fill="#ffda44" d="m367.3 256l-69.2 69.2a77.9 77.9 0 0 0 138.4 0z"/><path fill="#6da544" d="M289.4 133.6V256c0 59.6 77.9 78 77.9 78s78-18.4 78-78V133.6h-156z"/><path fill="#496e2d" d="M445.2 256zm-155.8 0z"/><path fill="#ffda44" d="M311.7 155.8h22.2v22.3h-22.2zm0 50.1h22.2v22.3h-22.2zm0 50.1h22.2v22.3h-22.2zm89-100.2H423v22.3h-22.3zm0 50.1H423v22.3h-22.3zm0 50.1H423v22.3h-22.3z"/><path fill="#eee" d="M0 0v32l32 32L0 96v160h32l32-32l32 32h32v-83l83 83h45l-8-16l8-15v-14l-83-83h83V96l-32-32l32-32V0H96L64 32L32 0Z"/><path fill="#d80027" d="M32 0v32H0v64h32v160h64V96h160V32H96V0Zm96 128l128 128v-31l-97-97z"/><path fill="#eee" d="M350.6 189.2H384v89h-33.4z"/><circle cx="367.3" cy="189.2" r="16.7" fill="#a2001d"/></g>`
	viInnerSVG                    = `<mask id="circleFlagsVi0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVi0)"><path fill="#eee" d="M0 0h512v512H0z"/><g fill="#ffda44"><path d="M299.5 178.8a43.5 43.5 0 1 0-87 0H117a44 44 0 0 0 44.2 42.9h-1.4a42.8 42.8 0 0 0 42.8 42.8l53.5 42l53.5-42a42.8 42.8 0 0 0 42.8-42.8H351a44 44 0 0 0 44.2-42.9h-95.6z"/><path d="m236.9 302.3l-27 60.9a122 122 0 0 0 46.1 9a122 122 0 0 0 46-9l-26.9-60.9H237z"/></g><path fill="#eee" d="m255.3 214l-55 21.5v39.7a49.1 49.1 0 0 0 22.3 40l11.3-2.3l11 14.3a98.2 98.2 0 0 0 11.1 3.6s4.6-1.1 11.1-3.6l9.4-14.1l13 2.1a49.2 49.2 0 0 0 22.2-40.1v-39.6z"/><path fill="#338af3" d="M62 257.8L92.3 342l29.8-84.2h22.7l-43.4 111.3H83L39.3 257.8h22.8zm349.8 111.3V257.8h21.7v111.3z"/><path fill="#d80027" d="M222.6 224.4v90.8a96 96 0 0 0 22.3 12V224.3zm66.8 0v90.8a98 98 0 0 1-22.3 12V224.3z"/><path fill="#0052b4" d="M200.3 200.2h111.4v35.3H200.3z"/></g>`
	vnInnerSVG                    = `<mask id="circleFlagsVn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVn0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m256 133.6l27.6 85H373L300.7 271l27.6 85l-72.3-52.5l-72.3 52.6l27.6-85l-72.3-52.6h89.4z"/></g>`
	voInnerSVG                    = `<clipPath id="svgIDa"><circle cx="256" cy="256" r="256"/></clipPath><g clip-path="url(#svgIDa)"><path fill="#4a1f63" d="M0 0h512v512H0z"/><path fill="#ffda44" d="M0 0h304v304H0Z"/><path fill="#4a1f63" d="M176 72a104 104 0 1 0 0 208a104 104 0 0 0 0-208zm-25 20c-8 8-16 18-22 31l-17-7c10-11 24-20 39-24zm50 0c15 4 29 13 39 24l-17 7c-6-13-14-23-22-31zm-17 5c7 7 16 17 24 30l-24 4zm-16 0v34l-24-4c8-13 17-23 24-30zm82 31c8 12 12 26 14 40h-28c-1-11-3-21-7-31zm-148 0l21 9c-4 10-6 20-7 31H88c2-14 6-28 14-40zm112 14c3 8 5 17 6 26h-36v-21zm-76 0l30 5v21h-36c1-9 3-18 6-26zm-50 42h28c1 11 3 21 7 31l-21 9c-8-12-12-25-14-40Zm44 0h36v21l-30 5c-3-8-5-17-6-26zm52 0h36c-1 9-3 18-6 26l-30-5zm52 0h28c-2 15-6 28-14 40l-21-9c4-10 6-20 7-31zm-68 37v34c-7-6-16-17-24-30zm16 0l24 4c-8 13-17 24-24 30zm39 8c7 3 13 5 17 8a89 89 0 0 1-39 23c8-8 16-18 22-31zm-94 0c6 13 14 23 22 31c-15-4-29-12-39-23z"/></g>`
	vuInnerSVG                    = `<mask id="circleFlagsVu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVu0)"><path fill="#d80027" d="M512 222.6V0H0l215.6 239.4z"/><path fill="#6da544" d="M512 289.4V512H0l216.5-240z"/><path fill="#333" d="m0 0l222.6 222.6H512v22.2l-10.9 11.8L512 267v22.4H222.6L0 512v-31.5l8-23l-8-8.6V63l9.5-10.4l-9.5-21z"/><g fill="#ffda44"><path d="M512 244.8H213.2L0 31.7V63l192.8 193L0 449v31.5L213.2 267H512z"/><path d="M62.2 310.6V289a38.8 38.8 0 0 0 38.8-38.8A27.5 27.5 0 0 0 73.5 223a18.7 18.7 0 0 0-18.7 18.7a12 12 0 0 0 12 12c3.7 0 6.8-3 6.8-6.8H95A28.2 28.2 0 0 1 66.8 275c-18.4 0-33.4-15-33.4-33.4c0-22.2 18-40.2 40-40.2a49 49 0 0 1 49 49c0 33.1-27 60.2-60.2 60.2z"/></g></g>`
	wfInnerSVG                    = `<mask id="circleFlagsWf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsWf0)"><path fill="#d80027" d="M256 0h256v512H0V256Z"/><path fill="#eee" d="M0 0h256v256H0Z"/><path fill="#0052b4" d="M0 0h75v224H0Z"/><path fill="#d80027" d="M149 0h75v224h-75z"/><path fill="#eee" d="m384 232l-72-72h144zm-24 24l-72-72v144zm24 24l-72 72h144zm24-24l72-72v144z"/></g>`
	wiphalaInnerSVG               = `<mask id="circleFlagsWiphala0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsWiphala0)"><path fill="#eee" d="M0 73V0h73l439 439v73h-73Z"/><path fill="#ffda44" d="M73 0v73h73v73h73v73h74v74h73v73h73v73h73v-73l-73-73l-73-73v-1l-73-73l-74-73l-73-73H73z"/><path fill="#ff9811" d="M146 0v73h73v73h74v73h73v74h73v73h73v-73l-73-74L293 73L219 0h-73z"/><path fill="#d80027" d="M219 0v73h74v73h73v73h73v74h73v-74l-73-73l-73-73l-73-73h-74z"/><path fill="#4a1f63" d="M293 0v73h73v73h73v73h73v-73l-73-73l-73-73h-73z"/><path fill="#0052b4" d="M366 0v73h73v73h73V73L439 0h-73z"/><path fill="#6da544" d="M439 0v73h73V0Zm0 512v-73h-74v-73h-73v-73h-73v-74h-73v-73H73V73H0v73l73 73l73 73v1l73 73l73 73l73 73z"/><path fill="#0052b4" d="M366 512v-73h-74v-73h-73v-73h-73v-74H73v-73H0v73l73 74l146 146l73 73z"/><path fill="#4a1f63" d="M292 512v-73h-73v-73h-73v-73H73v-74H0v74l73 73l73 73l73 73z"/><path fill="#d80027" d="M219 512v-73h-73v-73H73v-74H0v74l73 73l73 73z"/><path fill="#ff9811" d="M146 512v-73H73v-73H0v73l73 73z"/><path fill="#ffda44" d="M73 512v-73H0v73z"/></g>`
	wsInnerSVG                    = `<mask id="circleFlagsWs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsWs0)"><path fill="#d80027" d="M0 256L256 0h256v512H0z"/><path fill="#0052b4" d="M0 0h256v256H0z"/><path fill="#eee" d="m205 167l4.2 12.7h13.4l-10.8 7.9l4.1 12.7l-10.8-7.8l-10.9 7.8l4.1-12.7l-10.8-7.9h13.4zM137.8 66.8l7 21.2H167l-18.1 13.2l6.9 21.2l-18.1-13.1l-18 13.1l6.8-21.2l-18-13.2h22.3zM204.5 89l6.9 21.3h22.3l-18 13.1l6.9 21.3l-18.1-13.1l-18.1 13.1l7-21.3l-18.2-13.1h22.4zm-52.8 89l6.9 21.3H181l-18.1 13.2l6.9 21.2l-18.1-13.1l-18 13.1l6.8-21.2l-18-13.2h22.3zm-58.5-55.6l6.9 21.3h22.3l-18 13.1l6.9 21.3l-18.1-13.2L75 178.1l6.9-21.3l-18-13.1h22.3z"/></g>`
	xkInnerSVG                    = `<mask id="circleFlagsXk0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsXk0)"><path fill="#0052b4" d="M0 0h512v512H0z"/><path fill="#eee" d="m112.4 155.8l5.6 17h17.9l-14.5 10.5l5.5 17l-14.5-10.5L98 200.3l5.5-17L89 172.8h18zm55.7-16.7l5.5 17h18L177 166.6l5.6 17l-14.5-10.5l-14.5 10.6l5.6-17l-14.5-10.6h17.9zm55.7-16.7l5.5 17h17.9L232.7 150l5.5 17l-14.4-10.6l-14.5 10.6l5.5-17l-14.5-10.6h18zm175.8 33.4l-5.6 17h-17.9l14.5 10.5l-5.5 17l14.5-10.5l14.4 10.5l-5.5-17l14.5-10.5h-18zm-55.7-16.7l-5.5 17h-18l14.6 10.5l-5.6 17l14.5-10.5l14.5 10.6l-5.6-17l14.5-10.6h-17.9zm-55.7-16.7l-5.5 17h-17.9l14.5 10.6l-5.5 17l14.4-10.6l14.5 10.6l-5.5-17l14.5-10.6h-18z"/><path fill="#ffda44" d="M300.5 267.1L256 211.5l-22.3 11.1V245l-33.4 22h-22.2v29a89 89 0 0 1 55.6 82.5H256v-22.2l22.3-22.3l22.2 22.3l22.3-22.3v-22.2l22.2-33.4l-44.5-11.2z"/></g>`
	xxInnerSVG                    = `<mask id="circleFlagsXx0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsXx0)"><path fill="#eee" d="M0 0h512v512H0z"/><circle cx="253" cy="380" r="32" fill="#acabb1"/><path fill="#acabb1" d="M322.4 135.5c-15.6-13.6-37.4-20.3-65.5-20.3c-27.9 0-49.9 7.2-66 21.4a74.9 74.9 0 0 0-24.3 55.4h-.2v12.8H224l.1-9a35.2 35.2 0 0 1 9.3-24.8c5.8-6.1 13.7-9 23.5-9c20.7 0 31 11 31 33.4c0 7.4-2 14.5-6 21.1a124.2 124.2 0 0 1-23.9 26a90.4 90.4 0 0 0-24.8 32.3c-4.5 11-6.8 26.7-6.8 45.2h51l.8-13.1a54 54 0 0 1 17.3-33.9l16.2-15.2a131.4 131.4 0 0 0 26.4-33.2a69.5 69.5 0 0 0 7.6-31.8c-.1-24.7-7.8-43.7-23.3-57.3z"/></g>`
	yeInnerSVG                    = `<mask id="circleFlagsYe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsYe0)"><path fill="#eee" d="m0 167l253.8-19.3L512 167v178l-254.9 32.3L0 345z"/><path fill="#d80027" d="M0 0h512v167H0z"/><path fill="#333" d="M0 345h512v167H0z"/></g>`
	yiInnerSVG                    = `<clipPath id="svgIDa"><circle cx="256" cy="256" r="256"/></clipPath><g clip-path="url(#svgIDa)"><path fill="#eee" d="M0 0h512v56l-25 32l25 34v268l-26 36l26 30v56H0v-56l25-34l-25-32V122l27-33L0 56z"/><path fill="#333" d="M0 390v66h512v-66z"/><circle cx="256" cy="186" r="104" fill="#333"/><circle cx="256" cy="186" r="84" fill="#eee"/><circle cx="256" cy="186" r="64" fill="#333"/><circle cx="256" cy="186" r="44" fill="#eee"/><path fill="#333" d="m246 319l-49 11v14h118v-14l-49-11z"/><path fill="#eee" d="M148 78h216v108H148z"/><path fill="#333" d="M246 180h20v150h-20zm87-8h34l-7 15h-20zm-40 0h34l-7 15h-20zm-108 0h34l-7 15h-20zm-40 0h34l-7 15h-20z"/><path fill="#333" d="M239 172h34l-7 15h-20zM0 56v66h512V56z"/></g>`
	yorubalandInnerSVG            = `<mask id="circleFlagsYorubaland0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsYorubaland0)"><path fill="#eee" d="M384 0h128v128L352 352L128 512H0V384l160-224Z"/><path fill="#338af3" d="M0 384L384 0H0Z"/><path fill="#6da544" d="M512 128L128 512h384z"/><path fill="#ffda44" d="m432 234l-6 11h-13l6 11l-6 11h13l6 11l6-11h13l-6-11l6-11h-13zm-14-68l-6 11h-13l6 11l-6 10h13l6 12l6-12h13l-6-10l6-11h-12zm-37-56l-7 11h-12l6 11l-6 10h12l7 12l6-12h13l-7-10l7-11h-13zm-56-38l-7 11h-12l6 11l-6 11h12l7 11l6-11h13l-7-11l7-11h-13zm-69-14l-6 11h-13l6 11l-6 11h13l6 11l6-11h13l-6-11l6-11h-13zm-68 14l-7 11h-13l7 11l-6 11h12l7 11l6-11h13l-7-11l7-11h-13zm-56 38l-7 11h-13l7 11l-7 10h13l7 12l6-12h13l-7-10l7-11h-13zm-38 56l-7 11H75l6 11l-6 10h12l7 12l6-12h13l-7-10l7-11h-13zm-14 68l-6 11H61l6 11l-6 11h13l6 11l6-11h13l-6-11l6-11H86Zm14 69l-7 11H75l6 11l-6 10h12l7 11l6-11h13l-7-11l7-10h-13zm38 56l-7 11h-13l7 11l-7 10h13l7 11l6-11h13l-7-11l7-10h-13zm56 37l-7 11h-13l7 11l-6 11h12l7 11l6-11h13l-7-11l7-11h-13zm68 14l-6 11h-13l6 11l-6 11h13l6 11l6-11h13l-6-11l6-11h-13zm69-14l-7 11h-12l6 11l-6 11h12l7 11l6-11h13l-7-11l7-11h-13zm56-37l-7 11h-12l6 11l-6 10h12l7 11l6-11h13l-7-11l7-10h-13zm37-56l-6 11h-13l6 11l-6 10h13l6 11l6-11h13l-6-11l6-10h-13zM256 165c-3 0-6 3-6 6v20a38 28 0 0 0-32 28a38 28 0 0 0 1 5l3 23a34 54 0 0 0 17 47v22h34v-22a34 54 0 0 0 17-47l3-23a38 28 0 0 0 1-5a38 28 0 0 0-32-28v-20c0-3-3-6-6-6z"/><path fill="#d80027" d="M199 199a80 80 0 1 0 114 0l-17 17a56 56 0 0 1 16 40a56 56 0 0 1-56 56a56 56 0 0 1-56-56a56 56 0 0 1 16-40l-17-17z"/></g>`
	ytInnerSVG                    = `<mask id="circleFlagsYt0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsYt0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#acabb1" d="M113 139c-4 0-11 2-23 8a59 59 0 0 0-13 9c-28 27-20 68-11 79c3 4 16 23 32 30c22 8 36 28 37 51c0 24-15 44-34 46h-4c-19 0-36-19-37-43v2c1 29 20 52 43 52s44-12 55-32l5-8a43 43 0 0 0 4-14c4-19 6-70-47-95c-40-20-45-36-40-43a4 4 0 0 1 3-1a19 19 0 0 0 18 19a15 15 0 0 0 10-4s18 18 24 11s0-7-6-14c-5-7 0-21-11-28c-12-7-18-7-12-15a16 16 0 0 1 12-7s0-3-5-3zm287 0c-5 0-5 3-5 3c5 1 9 3 12 7c5 8 0 8-12 15s-6 21-12 28c-5 7-11 7-5 14c5 7 24-11 24-11a15 15 0 0 0 10 4a19 19 0 0 0 18-19a4 4 0 0 1 2 1c6 7 1 23-40 43c-52 25-50 76-47 95l5 14a189 189 0 0 0 5 8c10 20 32 33 54 32c24 0 43-23 44-52v-2c-2 24-18 43-38 43h-4c-19-2-33-22-33-46c1-23 15-43 36-51c16-7 29-26 33-30c8-11 17-52-11-79a59 59 0 0 0-14-9c-12-6-18-8-22-8zM26 199c-1 0-2 1-2 3c-2 16-1 52 33 64a2 2 0 0 0 3-1l6-8v-4c-5-7-22-31-22-44l-2-2a19 19 0 0 1-15-7l-1-1zm460 0l-1 1a19 19 0 0 1-15 7l-2 2c0 13-17 37-21 44v4l5 8a2 2 0 0 0 3 1c34-12 35-48 34-64l-3-3z"/><path fill="#0052b4" d="M211 147a22.5 22.5 0 0 0 0 45a22.5 22.5 0 0 0 0 45l45 32l45-32a22.5 22.5 0 0 0 0-45a22.5 22.5 0 1 0-22.5-22.5a22.5 22.5 0 0 0-45 0A22.5 22.5 0 0 0 211 147z"/><path fill="#d80027" d="M211 327a22.5 22.5 0 0 1-22.5-22.5A22.5 22.5 0 0 1 211 282a22.5 22.5 0 0 1-22.5-22.5A22.5 22.5 0 0 1 211 237h90a22.5 22.5 0 0 1 22.5 22.5A22.5 22.5 0 0 1 301 282a22.5 22.5 0 0 1 22.5 22.5A22.5 22.5 0 0 1 301 327a22.5 22.5 0 0 1-22.5-22.5A22.5 22.5 0 0 1 256 327a22.5 22.5 0 0 1-22.5-22.5A22.5 22.5 0 0 1 211 327z"/><path fill="#eee" d="M232 188a25 25 0 0 0-1 5a25 25 0 0 0 25 25a25 25 0 0 0 25-25a25 25 0 0 0-1-5a25 25 0 0 1-24 19a25 25 0 0 1-24-19z"/><path fill="#ffda44" d="M283 254a4 4 0 0 0-4 4v8l-7-4a4 4 0 0 0-2-1a4 4 0 0 0-3 2a4 4 0 0 0 1 6l7 4l-7 3a4 4 0 0 0-1 6a4 4 0 0 0 5 1l7-3v7a4 4 0 0 0 4 4a4 4 0 0 0 4-4v-7l7 3a4 4 0 0 0 5-1a4 4 0 0 0-1-6l-7-3l7-4a4 4 0 0 0 1-6a4 4 0 0 0-3-2a4 4 0 0 0-2 1l-7 4v-8a4 4 0 0 0-4-4zm-54 0a4 4 0 0 0-4 4v8l-7-4a4 4 0 0 0-2-1a4 4 0 0 0-3 2a4 4 0 0 0 1 6l7 4l-7 3a4 4 0 0 0-1 6a4 4 0 0 0 5 1l7-3v7a4 4 0 0 0 4 4a4 4 0 0 0 4-4v-7l7 3a4 4 0 0 0 5-1a4 4 0 0 0-1-6l-7-3l7-4a4 4 0 0 0 1-6a4 4 0 0 0-3-2a4 4 0 0 0-2 1l-7 4v-8a4 4 0 0 0-4-4z"/></g>`
	zaInnerSVG                    = `<mask id="circleFlagsZa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsZa0)"><path fill="#eee" d="m0 0l192 256L0 512h47l465-189v-34l-32-33l32-33v-34L47 0Z"/><path fill="#333" d="M0 142v228l140-114z"/><path fill="#ffda44" d="M192 256L0 95v47l114 114L0 370v47z"/><path fill="#6da544" d="M512 223H223L0 0v94l161 162L0 418v94l223-223h289z"/><path fill="#d80027" d="M512 0H47l189 189h276z"/><path fill="#0052b4" d="M512 512H47l189-189h276z"/></g>`
	zmInnerSVG                    = `<mask id="circleFlagsZm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsZm0)"><path fill="#496e2d" d="M0 0h512v256L256 512H0z"/><path fill="#ff9811" d="M473 167h-66.7a22.3 22.3 0 0 0-44.6 0H295a23 23 0 0 0 23 22.2h-.8c0 12.3 10 22.3 22.3 22.3c0 12.3 10 22.2 22.2 22.2h44.6c12.3 0 22.2-10 22.2-22.2c12.3 0 22.3-10 22.3-22.3h-.8a23 23 0 0 0 23-22.2z"/><path fill="#333" d="M341.3 256h85.4l21.1 126.3L426.7 512h-85.4l-23.5-128z"/><path fill="#d80027" d="M256 256h85.3v256H256z"/><path fill="#ff9811" d="M426.7 256H512v256h-85.3z"/></g>`
	zwInnerSVG                    = `<mask id="circleFlagsZw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsZw0)"><path fill="#6da544" d="M31.4 0H512v512H31.4z"/><path fill="#ffda44" d="M57.8 73.3H512v73.3l-65.1 110l65.1 110v73.3H57.8z"/><path fill="#d80027" d="M132 146.6h380v73.3l-27 36l27 37.3v73.3H132z"/><path fill="#eee" d="M0 0v512l279.8-256z"/><path fill="#d80027" d="m103 189.2l16.5 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51l-43.4-31.6h53.6z"/><path fill="#ffda44" d="m148.5 260.2l-43.2-15.3l-3.4-31a16.7 16.7 0 1 0-32.5 7.6l-12 12.1h21.5c0 22.4-16.7 22.4-16.7 44.7l9.2 22.2h55.7l9.3-22.2a22.2 22.2 0 0 0 1.7-6.6c8-3.2 10.4-11.5 10.4-11.5z"/><path fill="#333" d="m31.4 0l220 220H512v73.2H250.2L31.4 512H0l256-256L0 0z"/></g>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Ac(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(acInnerSVG).
		Element(children...)
}

func Ad(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adInnerSVG).
		Element(children...)
}

func Ae(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aeInnerSVG).
		Element(children...)
}

func Af(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(afInnerSVG).
		Element(children...)
}

func Afar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(afarInnerSVG).
		Element(children...)
}

func Ag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(agInnerSVG).
		Element(children...)
}

func Ai(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aiInnerSVG).
		Element(children...)
}

func Al(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alInnerSVG).
		Element(children...)
}

func Am(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(amInnerSVG).
		Element(children...)
}

func An(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(anInnerSVG).
		Element(children...)
}

func Ao(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aoInnerSVG).
		Element(children...)
}

func Aq(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aqInnerSVG).
		Element(children...)
}

func Ar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arInnerSVG).
		Element(children...)
}

func As(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(asInnerSVG).
		Element(children...)
}

func At(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atInnerSVG).
		Element(children...)
}

func Au(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auInnerSVG).
		Element(children...)
}

func AuAboriginal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auAboriginalInnerSVG).
		Element(children...)
}

func AuAct(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auActInnerSVG).
		Element(children...)
}

func AuNsw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auNswInnerSVG).
		Element(children...)
}

func AuNt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auNtInnerSVG).
		Element(children...)
}

func AuQld(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auQldInnerSVG).
		Element(children...)
}

func AuSa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auSaInnerSVG).
		Element(children...)
}

func AuTas(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auTasInnerSVG).
		Element(children...)
}

func AuVic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auVicInnerSVG).
		Element(children...)
}

func AuWa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(auWaInnerSVG).
		Element(children...)
}

func Aw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(awInnerSVG).
		Element(children...)
}

func Ax(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(axInnerSVG).
		Element(children...)
}

func Az(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(azInnerSVG).
		Element(children...)
}

func Ba(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(baInnerSVG).
		Element(children...)
}

func Bb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bbInnerSVG).
		Element(children...)
}

func Bd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bdInnerSVG).
		Element(children...)
}

func Be(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(beInnerSVG).
		Element(children...)
}

func Bf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bfInnerSVG).
		Element(children...)
}

func Bg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bgInnerSVG).
		Element(children...)
}

func Bh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bhInnerSVG).
		Element(children...)
}

func Bi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(biInnerSVG).
		Element(children...)
}

func Bj(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bjInnerSVG).
		Element(children...)
}

func Bl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(blInnerSVG).
		Element(children...)
}

func Bm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bmInnerSVG).
		Element(children...)
}

func Bn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bnInnerSVG).
		Element(children...)
}

func Bo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boInnerSVG).
		Element(children...)
}

func Bq(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bqInnerSVG).
		Element(children...)
}

func BqBo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bqBoInnerSVG).
		Element(children...)
}

func BqSa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bqSaInnerSVG).
		Element(children...)
}

func BqSe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bqSeInnerSVG).
		Element(children...)
}

func Br(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brInnerSVG).
		Element(children...)
}

func Bs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bsInnerSVG).
		Element(children...)
}

func Bt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(btInnerSVG).
		Element(children...)
}

func Bv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bvInnerSVG).
		Element(children...)
}

func Bw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bwInnerSVG).
		Element(children...)
}

func By(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(byInnerSVG).
		Element(children...)
}

func Bz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bzInnerSVG).
		Element(children...)
}

func Ca(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caInnerSVG).
		Element(children...)
}

func CaBc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caBcInnerSVG).
		Element(children...)
}

func Cc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ccInnerSVG).
		Element(children...)
}

func Cd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cdInnerSVG).
		Element(children...)
}

func Cf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cfInnerSVG).
		Element(children...)
}

func Cg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cgInnerSVG).
		Element(children...)
}

func Ch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chInnerSVG).
		Element(children...)
}

func ChGr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chGrInnerSVG).
		Element(children...)
}

func Ci(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ciInnerSVG).
		Element(children...)
}

func Ck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ckInnerSVG).
		Element(children...)
}

func Cl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clInnerSVG).
		Element(children...)
}

func Cm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cmInnerSVG).
		Element(children...)
}

func Cn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cnInnerSVG).
		Element(children...)
}

func CnXj(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cnXjInnerSVG).
		Element(children...)
}

func Co(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coInnerSVG).
		Element(children...)
}

func Cp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cpInnerSVG).
		Element(children...)
}

func Cq(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cqInnerSVG).
		Element(children...)
}

func Cr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crInnerSVG).
		Element(children...)
}

func Cu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cuInnerSVG).
		Element(children...)
}

func Cv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cvInnerSVG).
		Element(children...)
}

func Cw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cwInnerSVG).
		Element(children...)
}

func Cx(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cxInnerSVG).
		Element(children...)
}

func Cy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cyInnerSVG).
		Element(children...)
}

func Cz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(czInnerSVG).
		Element(children...)
}

func De(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deInnerSVG).
		Element(children...)
}

func Dg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dgInnerSVG).
		Element(children...)
}

func Dj(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(djInnerSVG).
		Element(children...)
}

func Dk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dkInnerSVG).
		Element(children...)
}

func Dm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dmInnerSVG).
		Element(children...)
}

func Do(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doInnerSVG).
		Element(children...)
}

func Dz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dzInnerSVG).
		Element(children...)
}

func Ea(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eaInnerSVG).
		Element(children...)
}

func Earth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(earthInnerSVG).
		Element(children...)
}

func EastAfricanFederation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eastAfricanFederationInnerSVG).
		Element(children...)
}

func EasterIsland(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(easterIslandInnerSVG).
		Element(children...)
}

func Ec(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ecInnerSVG).
		Element(children...)
}

func EcW(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ecWInnerSVG).
		Element(children...)
}

func Ee(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eeInnerSVG).
		Element(children...)
}

func Eg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(egInnerSVG).
		Element(children...)
}

func Eh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ehInnerSVG).
		Element(children...)
}

func En(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(enInnerSVG).
		Element(children...)
}

func Eo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eoInnerSVG).
		Element(children...)
}

func Er(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(erInnerSVG).
		Element(children...)
}

func Es(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esInnerSVG).
		Element(children...)
}

func EsAr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esArInnerSVG).
		Element(children...)
}

func EsCe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esCeInnerSVG).
		Element(children...)
}

func EsCn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esCnInnerSVG).
		Element(children...)
}

func EsCt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esCtInnerSVG).
		Element(children...)
}

func EsGa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esGaInnerSVG).
		Element(children...)
}

func EsIb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esIbInnerSVG).
		Element(children...)
}

func EsMl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esMlInnerSVG).
		Element(children...)
}

func EsPv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esPvInnerSVG).
		Element(children...)
}

func EsVariant(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(esVariantInnerSVG).
		Element(children...)
}

func Et(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(etInnerSVG).
		Element(children...)
}

func EtOr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(etOrInnerSVG).
		Element(children...)
}

func EtTi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(etTiInnerSVG).
		Element(children...)
}

func Eu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(euInnerSVG).
		Element(children...)
}

func EuropeanUnion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(europeanUnionInnerSVG).
		Element(children...)
}

func Ewe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eweInnerSVG).
		Element(children...)
}

func Fi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fiInnerSVG).
		Element(children...)
}

func Fj(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fjInnerSVG).
		Element(children...)
}

func Fk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fkInnerSVG).
		Element(children...)
}

func Fm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fmInnerSVG).
		Element(children...)
}

func Fo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(foInnerSVG).
		Element(children...)
}

func Fr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frInnerSVG).
		Element(children...)
}

func FrBre(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frBreInnerSVG).
		Element(children...)
}

func FrCp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frCpInnerSVG).
		Element(children...)
}

func FrTwentyR(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(frTwentyRInnerSVG).
		Element(children...)
}

func Fx(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fxInnerSVG).
		Element(children...)
}

func Ga(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gaInnerSVG).
		Element(children...)
}

func Gb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbInnerSVG).
		Element(children...)
}

func GbCon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbConInnerSVG).
		Element(children...)
}

func GbEng(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbEngInnerSVG).
		Element(children...)
}

func GbNir(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbNirInnerSVG).
		Element(children...)
}

func GbOrk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbOrkInnerSVG).
		Element(children...)
}

func GbSct(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbSctInnerSVG).
		Element(children...)
}

func GbWls(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gbWlsInnerSVG).
		Element(children...)
}

func Gd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gdInnerSVG).
		Element(children...)
}

func Ge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(geInnerSVG).
		Element(children...)
}

func GeAb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(geAbInnerSVG).
		Element(children...)
}

func Gf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gfInnerSVG).
		Element(children...)
}

func Gg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ggInnerSVG).
		Element(children...)
}

func Gh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ghInnerSVG).
		Element(children...)
}

func Gi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(giInnerSVG).
		Element(children...)
}

func Gl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(glInnerSVG).
		Element(children...)
}

func Gm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gmInnerSVG).
		Element(children...)
}

func Gn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gnInnerSVG).
		Element(children...)
}

func Gp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gpInnerSVG).
		Element(children...)
}

func Gq(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gqInnerSVG).
		Element(children...)
}

func Gr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(grInnerSVG).
		Element(children...)
}

func Gs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gsInnerSVG).
		Element(children...)
}

func Gt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gtInnerSVG).
		Element(children...)
}

func Gu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(guInnerSVG).
		Element(children...)
}

func Guarani(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(guaraniInnerSVG).
		Element(children...)
}

func Gw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gwInnerSVG).
		Element(children...)
}

func Gy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gyInnerSVG).
		Element(children...)
}

func Hausa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hausaInnerSVG).
		Element(children...)
}

func Hk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hkInnerSVG).
		Element(children...)
}

func Hm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hmInnerSVG).
		Element(children...)
}

func Hmong(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hmongInnerSVG).
		Element(children...)
}

func Hn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hnInnerSVG).
		Element(children...)
}

func Hr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hrInnerSVG).
		Element(children...)
}

func Ht(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(htInnerSVG).
		Element(children...)
}

func Hu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(huInnerSVG).
		Element(children...)
}

func Ia(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iaInnerSVG).
		Element(children...)
}

func Ic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(icInnerSVG).
		Element(children...)
}

func Id(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(idInnerSVG).
		Element(children...)
}

func IdJb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(idJbInnerSVG).
		Element(children...)
}

func IdJt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(idJtInnerSVG).
		Element(children...)
}

func Ie(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ieInnerSVG).
		Element(children...)
}

func Il(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ilInnerSVG).
		Element(children...)
}

func Im(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imInnerSVG).
		Element(children...)
}

func In(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inInnerSVG).
		Element(children...)
}

func InAs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inAsInnerSVG).
		Element(children...)
}

func InGj(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inGjInnerSVG).
		Element(children...)
}

func InKa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inKaInnerSVG).
		Element(children...)
}

func InOr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inOrInnerSVG).
		Element(children...)
}

func InTg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inTgInnerSVG).
		Element(children...)
}

func InTn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inTnInnerSVG).
		Element(children...)
}

func Io(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ioInnerSVG).
		Element(children...)
}

func Iq(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iqInnerSVG).
		Element(children...)
}

func Ir(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(irInnerSVG).
		Element(children...)
}

func Is(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(isInnerSVG).
		Element(children...)
}

func It(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(itInnerSVG).
		Element(children...)
}

func ItEightyEight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(itEightyEightInnerSVG).
		Element(children...)
}

func ItEightyTwo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(itEightyTwoInnerSVG).
		Element(children...)
}

func ItTwentyThree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(itTwentyThreeInnerSVG).
		Element(children...)
}

func Je(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jeInnerSVG).
		Element(children...)
}

func Jm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jmInnerSVG).
		Element(children...)
}

func Jo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(joInnerSVG).
		Element(children...)
}

func Jp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jpInnerSVG).
		Element(children...)
}

func Kanuri(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kanuriInnerSVG).
		Element(children...)
}

func Ke(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keInnerSVG).
		Element(children...)
}

func Kg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kgInnerSVG).
		Element(children...)
}

func Kh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(khInnerSVG).
		Element(children...)
}

func Ki(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kiInnerSVG).
		Element(children...)
}

func Kikuyu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kikuyuInnerSVG).
		Element(children...)
}

func Klingon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(klingonInnerSVG).
		Element(children...)
}

func Km(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kmInnerSVG).
		Element(children...)
}

func Kn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(knInnerSVG).
		Element(children...)
}

func Kongo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kongoInnerSVG).
		Element(children...)
}

func Kp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kpInnerSVG).
		Element(children...)
}

func Kr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(krInnerSVG).
		Element(children...)
}

func Kurdistan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kurdistanInnerSVG).
		Element(children...)
}

func Kw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kwInnerSVG).
		Element(children...)
}

func Ky(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kyInnerSVG).
		Element(children...)
}

func Kz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(kzInnerSVG).
		Element(children...)
}

func La(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(laInnerSVG).
		Element(children...)
}

func Lb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lbInnerSVG).
		Element(children...)
}

func Lc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lcInnerSVG).
		Element(children...)
}

func Li(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(liInnerSVG).
		Element(children...)
}

func Lk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lkInnerSVG).
		Element(children...)
}

func Lr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lrInnerSVG).
		Element(children...)
}

func Ls(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lsInnerSVG).
		Element(children...)
}

func Lt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ltInnerSVG).
		Element(children...)
}

func Lu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(luInnerSVG).
		Element(children...)
}

func Lv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lvInnerSVG).
		Element(children...)
}

func Ly(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lyInnerSVG).
		Element(children...)
}

func Ma(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maInnerSVG).
		Element(children...)
}

func Malayali(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(malayaliInnerSVG).
		Element(children...)
}

func Manipur(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(manipurInnerSVG).
		Element(children...)
}

func Maori(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(maoriInnerSVG).
		Element(children...)
}

func Mc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mcInnerSVG).
		Element(children...)
}

func Md(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mdInnerSVG).
		Element(children...)
}

func Me(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(meInnerSVG).
		Element(children...)
}

func Mf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mfInnerSVG).
		Element(children...)
}

func Mg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mgInnerSVG).
		Element(children...)
}

func Mh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mhInnerSVG).
		Element(children...)
}

func Mizoram(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mizoramInnerSVG).
		Element(children...)
}

func Mk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mkInnerSVG).
		Element(children...)
}

func Ml(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mlInnerSVG).
		Element(children...)
}

func Mm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mmInnerSVG).
		Element(children...)
}

func Mn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mnInnerSVG).
		Element(children...)
}

func Mo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moInnerSVG).
		Element(children...)
}

func Mp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mpInnerSVG).
		Element(children...)
}

func Mq(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mqInnerSVG).
		Element(children...)
}

func Mr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mrInnerSVG).
		Element(children...)
}

func Ms(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(msInnerSVG).
		Element(children...)
}

func Mt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mtInnerSVG).
		Element(children...)
}

func Mu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(muInnerSVG).
		Element(children...)
}

func Mv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mvInnerSVG).
		Element(children...)
}

func Mw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mwInnerSVG).
		Element(children...)
}

func Mx(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mxInnerSVG).
		Element(children...)
}

func My(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(myInnerSVG).
		Element(children...)
}

func Mz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mzInnerSVG).
		Element(children...)
}

func Na(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(naInnerSVG).
		Element(children...)
}

func Nato(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(natoInnerSVG).
		Element(children...)
}

func Nc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ncInnerSVG).
		Element(children...)
}

func Ne(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(neInnerSVG).
		Element(children...)
}

func Nf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nfInnerSVG).
		Element(children...)
}

func Ng(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ngInnerSVG).
		Element(children...)
}

func Ni(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(niInnerSVG).
		Element(children...)
}

func Nl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nlInnerSVG).
		Element(children...)
}

func NlFr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nlFrInnerSVG).
		Element(children...)
}

func No(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noInnerSVG).
		Element(children...)
}

func NorthernCyprus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(northernCyprusInnerSVG).
		Element(children...)
}

func Np(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(npInnerSVG).
		Element(children...)
}

func Nr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nrInnerSVG).
		Element(children...)
}

func Nu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nuInnerSVG).
		Element(children...)
}

func Nz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nzInnerSVG).
		Element(children...)
}

func Occitania(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(occitaniaInnerSVG).
		Element(children...)
}

func Olympics(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(olympicsInnerSVG).
		Element(children...)
}

func Om(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(omInnerSVG).
		Element(children...)
}

func Otomi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(otomiInnerSVG).
		Element(children...)
}

func Pa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paInnerSVG).
		Element(children...)
}

func Pe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(peInnerSVG).
		Element(children...)
}

func Pf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pfInnerSVG).
		Element(children...)
}

func Pg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pgInnerSVG).
		Element(children...)
}

func Ph(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phInnerSVG).
		Element(children...)
}

func Pk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pkInnerSVG).
		Element(children...)
}

func PkJk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pkJkInnerSVG).
		Element(children...)
}

func PkSd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pkSdInnerSVG).
		Element(children...)
}

func Pl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plInnerSVG).
		Element(children...)
}

func Pm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pmInnerSVG).
		Element(children...)
}

func Pn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pnInnerSVG).
		Element(children...)
}

func Pr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(prInnerSVG).
		Element(children...)
}

func Ps(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(psInnerSVG).
		Element(children...)
}

func Pt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ptInnerSVG).
		Element(children...)
}

func PtThirty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ptThirtyInnerSVG).
		Element(children...)
}

func PtTwenty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ptTwentyInnerSVG).
		Element(children...)
}

func Pw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pwInnerSVG).
		Element(children...)
}

func Py(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pyInnerSVG).
		Element(children...)
}

func Qa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(qaInnerSVG).
		Element(children...)
}

func Quechua(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(quechuaInnerSVG).
		Element(children...)
}

func Re(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reInnerSVG).
		Element(children...)
}

func Ro(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(roInnerSVG).
		Element(children...)
}

func Rs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rsInnerSVG).
		Element(children...)
}

func Ru(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ruInnerSVG).
		Element(children...)
}

func RuBa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ruBaInnerSVG).
		Element(children...)
}

func RuCe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ruCeInnerSVG).
		Element(children...)
}

func RuCu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ruCuInnerSVG).
		Element(children...)
}

func RuDa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ruDaInnerSVG).
		Element(children...)
}

func RuKo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ruKoInnerSVG).
		Element(children...)
}

func RuTa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ruTaInnerSVG).
		Element(children...)
}

func RuUd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ruUdInnerSVG).
		Element(children...)
}

func Rw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rwInnerSVG).
		Element(children...)
}

func Sa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(saInnerSVG).
		Element(children...)
}

func Sami(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(samiInnerSVG).
		Element(children...)
}

func Sb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sbInnerSVG).
		Element(children...)
}

func Sc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scInnerSVG).
		Element(children...)
}

func Sd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sdInnerSVG).
		Element(children...)
}

func Se(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(seInnerSVG).
		Element(children...)
}

func Sg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sgInnerSVG).
		Element(children...)
}

func Sh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shInnerSVG).
		Element(children...)
}

func ShAc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shAcInnerSVG).
		Element(children...)
}

func ShHl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shHlInnerSVG).
		Element(children...)
}

func ShTa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shTaInnerSVG).
		Element(children...)
}

func Si(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(siInnerSVG).
		Element(children...)
}

func Sj(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sjInnerSVG).
		Element(children...)
}

func Sk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skInnerSVG).
		Element(children...)
}

func Sl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(slInnerSVG).
		Element(children...)
}

func Sm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(smInnerSVG).
		Element(children...)
}

func Sn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snInnerSVG).
		Element(children...)
}

func So(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soInnerSVG).
		Element(children...)
}

func Somaliland(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(somalilandInnerSVG).
		Element(children...)
}

func SouthOssetia(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(southOssetiaInnerSVG).
		Element(children...)
}

func SovietUnion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sovietUnionInnerSVG).
		Element(children...)
}

func Sr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(srInnerSVG).
		Element(children...)
}

func Ss(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ssInnerSVG).
		Element(children...)
}

func St(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stInnerSVG).
		Element(children...)
}

func Su(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(suInnerSVG).
		Element(children...)
}

func Sv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(svInnerSVG).
		Element(children...)
}

func Sx(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sxInnerSVG).
		Element(children...)
}

func Sy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(syInnerSVG).
		Element(children...)
}

func Sz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(szInnerSVG).
		Element(children...)
}

func Ta(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(taInnerSVG).
		Element(children...)
}

func Tc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tcInnerSVG).
		Element(children...)
}

func Td(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tdInnerSVG).
		Element(children...)
}

func Tf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tfInnerSVG).
		Element(children...)
}

func Tg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tgInnerSVG).
		Element(children...)
}

func Th(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(thInnerSVG).
		Element(children...)
}

func Tibet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tibetInnerSVG).
		Element(children...)
}

func Tj(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tjInnerSVG).
		Element(children...)
}

func Tk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tkInnerSVG).
		Element(children...)
}

func Tl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tlInnerSVG).
		Element(children...)
}

func Tm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tmInnerSVG).
		Element(children...)
}

func Tn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tnInnerSVG).
		Element(children...)
}

func To(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toInnerSVG).
		Element(children...)
}

func TorresStraitIslands(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(torresStraitIslandsInnerSVG).
		Element(children...)
}

func Tr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trInnerSVG).
		Element(children...)
}

func Transnistria(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transnistriaInnerSVG).
		Element(children...)
}

func Tt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ttInnerSVG).
		Element(children...)
}

func Tv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tvInnerSVG).
		Element(children...)
}

func Tw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(twInnerSVG).
		Element(children...)
}

func Tz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tzInnerSVG).
		Element(children...)
}

func Ua(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uaInnerSVG).
		Element(children...)
}

func Ug(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ugInnerSVG).
		Element(children...)
}

func Uk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ukInnerSVG).
		Element(children...)
}

func Um(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(umInnerSVG).
		Element(children...)
}

func Un(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unInnerSVG).
		Element(children...)
}

func UnitedNations(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unitedNationsInnerSVG).
		Element(children...)
}

func Us(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usInnerSVG).
		Element(children...)
}

func UsAk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usAkInnerSVG).
		Element(children...)
}

func UsAl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usAlInnerSVG).
		Element(children...)
}

func UsAr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usArInnerSVG).
		Element(children...)
}

func UsAz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usAzInnerSVG).
		Element(children...)
}

func UsCa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usCaInnerSVG).
		Element(children...)
}

func UsCo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usCoInnerSVG).
		Element(children...)
}

func UsDc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usDcInnerSVG).
		Element(children...)
}

func UsFl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usFlInnerSVG).
		Element(children...)
}

func UsGa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usGaInnerSVG).
		Element(children...)
}

func UsHi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usHiInnerSVG).
		Element(children...)
}

func UsIn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usInInnerSVG).
		Element(children...)
}

func UsMo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usMoInnerSVG).
		Element(children...)
}

func UsMs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usMsInnerSVG).
		Element(children...)
}

func UsNc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usNcInnerSVG).
		Element(children...)
}

func UsNm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usNmInnerSVG).
		Element(children...)
}

func UsRi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usRiInnerSVG).
		Element(children...)
}

func UsTn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usTnInnerSVG).
		Element(children...)
}

func UsTx(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usTxInnerSVG).
		Element(children...)
}

func Uy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uyInnerSVG).
		Element(children...)
}

func Uz(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uzInnerSVG).
		Element(children...)
}

func Va(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vaInnerSVG).
		Element(children...)
}

func Vc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vcInnerSVG).
		Element(children...)
}

func Ve(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(veInnerSVG).
		Element(children...)
}

func Vg(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vgInnerSVG).
		Element(children...)
}

func Vi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viInnerSVG).
		Element(children...)
}

func Vn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vnInnerSVG).
		Element(children...)
}

func Vo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(voInnerSVG).
		Element(children...)
}

func Vu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vuInnerSVG).
		Element(children...)
}

func Wf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wfInnerSVG).
		Element(children...)
}

func Wiphala(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wiphalaInnerSVG).
		Element(children...)
}

func Ws(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wsInnerSVG).
		Element(children...)
}

func Xk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xkInnerSVG).
		Element(children...)
}

func Xx(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(xxInnerSVG).
		Element(children...)
}

func Ye(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yeInnerSVG).
		Element(children...)
}

func Yi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yiInnerSVG).
		Element(children...)
}

func Yorubaland(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yorubalandInnerSVG).
		Element(children...)
}

func Yt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ytInnerSVG).
		Element(children...)
}

func Za(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zaInnerSVG).
		Element(children...)
}

func Zm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zmInnerSVG).
		Element(children...)
}

func Zw(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 512 512"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zwInnerSVG).
		Element(children...)
}

func ByName(name string) (*engine.UberElement, error) {
	switch name {
	case "ac":
		return Ac(), nil
	case "ad":
		return Ad(), nil
	case "ae":
		return Ae(), nil
	case "af":
		return Af(), nil
	case "afar":
		return Afar(), nil
	case "ag":
		return Ag(), nil
	case "ai":
		return Ai(), nil
	case "al":
		return Al(), nil
	case "am":
		return Am(), nil
	case "an":
		return An(), nil
	case "ao":
		return Ao(), nil
	case "aq":
		return Aq(), nil
	case "ar":
		return Ar(), nil
	case "as":
		return As(), nil
	case "at":
		return At(), nil
	case "au":
		return Au(), nil
	case "au-aboriginal":
		return AuAboriginal(), nil
	case "au-act":
		return AuAct(), nil
	case "au-nsw":
		return AuNsw(), nil
	case "au-nt":
		return AuNt(), nil
	case "au-qld":
		return AuQld(), nil
	case "au-sa":
		return AuSa(), nil
	case "au-tas":
		return AuTas(), nil
	case "au-vic":
		return AuVic(), nil
	case "au-wa":
		return AuWa(), nil
	case "aw":
		return Aw(), nil
	case "ax":
		return Ax(), nil
	case "az":
		return Az(), nil
	case "ba":
		return Ba(), nil
	case "bb":
		return Bb(), nil
	case "bd":
		return Bd(), nil
	case "be":
		return Be(), nil
	case "bf":
		return Bf(), nil
	case "bg":
		return Bg(), nil
	case "bh":
		return Bh(), nil
	case "bi":
		return Bi(), nil
	case "bj":
		return Bj(), nil
	case "bl":
		return Bl(), nil
	case "bm":
		return Bm(), nil
	case "bn":
		return Bn(), nil
	case "bo":
		return Bo(), nil
	case "bq":
		return Bq(), nil
	case "bq-bo":
		return BqBo(), nil
	case "bq-sa":
		return BqSa(), nil
	case "bq-se":
		return BqSe(), nil
	case "br":
		return Br(), nil
	case "bs":
		return Bs(), nil
	case "bt":
		return Bt(), nil
	case "bv":
		return Bv(), nil
	case "bw":
		return Bw(), nil
	case "by":
		return By(), nil
	case "bz":
		return Bz(), nil
	case "ca":
		return Ca(), nil
	case "ca-bc":
		return CaBc(), nil
	case "cc":
		return Cc(), nil
	case "cd":
		return Cd(), nil
	case "cf":
		return Cf(), nil
	case "cg":
		return Cg(), nil
	case "ch":
		return Ch(), nil
	case "ch-gr":
		return ChGr(), nil
	case "ci":
		return Ci(), nil
	case "ck":
		return Ck(), nil
	case "cl":
		return Cl(), nil
	case "cm":
		return Cm(), nil
	case "cn":
		return Cn(), nil
	case "cn-xj":
		return CnXj(), nil
	case "co":
		return Co(), nil
	case "cp":
		return Cp(), nil
	case "cq":
		return Cq(), nil
	case "cr":
		return Cr(), nil
	case "cu":
		return Cu(), nil
	case "cv":
		return Cv(), nil
	case "cw":
		return Cw(), nil
	case "cx":
		return Cx(), nil
	case "cy":
		return Cy(), nil
	case "cz":
		return Cz(), nil
	case "de":
		return De(), nil
	case "dg":
		return Dg(), nil
	case "dj":
		return Dj(), nil
	case "dk":
		return Dk(), nil
	case "dm":
		return Dm(), nil
	case "do":
		return Do(), nil
	case "dz":
		return Dz(), nil
	case "ea":
		return Ea(), nil
	case "earth":
		return Earth(), nil
	case "east-african-federation":
		return EastAfricanFederation(), nil
	case "easter-island":
		return EasterIsland(), nil
	case "ec":
		return Ec(), nil
	case "ec-w":
		return EcW(), nil
	case "ee":
		return Ee(), nil
	case "eg":
		return Eg(), nil
	case "eh":
		return Eh(), nil
	case "en":
		return En(), nil
	case "eo":
		return Eo(), nil
	case "er":
		return Er(), nil
	case "es":
		return Es(), nil
	case "es-ar":
		return EsAr(), nil
	case "es-ce":
		return EsCe(), nil
	case "es-cn":
		return EsCn(), nil
	case "es-ct":
		return EsCt(), nil
	case "es-ga":
		return EsGa(), nil
	case "es-ib":
		return EsIb(), nil
	case "es-ml":
		return EsMl(), nil
	case "es-pv":
		return EsPv(), nil
	case "es-variant":
		return EsVariant(), nil
	case "et":
		return Et(), nil
	case "et-or":
		return EtOr(), nil
	case "et-ti":
		return EtTi(), nil
	case "eu":
		return Eu(), nil
	case "european-union":
		return EuropeanUnion(), nil
	case "ewe":
		return Ewe(), nil
	case "fi":
		return Fi(), nil
	case "fj":
		return Fj(), nil
	case "fk":
		return Fk(), nil
	case "fm":
		return Fm(), nil
	case "fo":
		return Fo(), nil
	case "fr":
		return Fr(), nil
	case "fr-bre":
		return FrBre(), nil
	case "fr-cp":
		return FrCp(), nil
	case "fr-20r":
		return FrTwentyR(), nil
	case "fx":
		return Fx(), nil
	case "ga":
		return Ga(), nil
	case "gb":
		return Gb(), nil
	case "gb-con":
		return GbCon(), nil
	case "gb-eng":
		return GbEng(), nil
	case "gb-nir":
		return GbNir(), nil
	case "gb-ork":
		return GbOrk(), nil
	case "gb-sct":
		return GbSct(), nil
	case "gb-wls":
		return GbWls(), nil
	case "gd":
		return Gd(), nil
	case "ge":
		return Ge(), nil
	case "ge-ab":
		return GeAb(), nil
	case "gf":
		return Gf(), nil
	case "gg":
		return Gg(), nil
	case "gh":
		return Gh(), nil
	case "gi":
		return Gi(), nil
	case "gl":
		return Gl(), nil
	case "gm":
		return Gm(), nil
	case "gn":
		return Gn(), nil
	case "gp":
		return Gp(), nil
	case "gq":
		return Gq(), nil
	case "gr":
		return Gr(), nil
	case "gs":
		return Gs(), nil
	case "gt":
		return Gt(), nil
	case "gu":
		return Gu(), nil
	case "guarani":
		return Guarani(), nil
	case "gw":
		return Gw(), nil
	case "gy":
		return Gy(), nil
	case "hausa":
		return Hausa(), nil
	case "hk":
		return Hk(), nil
	case "hm":
		return Hm(), nil
	case "hmong":
		return Hmong(), nil
	case "hn":
		return Hn(), nil
	case "hr":
		return Hr(), nil
	case "ht":
		return Ht(), nil
	case "hu":
		return Hu(), nil
	case "ia":
		return Ia(), nil
	case "ic":
		return Ic(), nil
	case "id":
		return Id(), nil
	case "id-jb":
		return IdJb(), nil
	case "id-jt":
		return IdJt(), nil
	case "ie":
		return Ie(), nil
	case "il":
		return Il(), nil
	case "im":
		return Im(), nil
	case "in":
		return In(), nil
	case "in-as":
		return InAs(), nil
	case "in-gj":
		return InGj(), nil
	case "in-ka":
		return InKa(), nil
	case "in-or":
		return InOr(), nil
	case "in-tg":
		return InTg(), nil
	case "in-tn":
		return InTn(), nil
	case "io":
		return Io(), nil
	case "iq":
		return Iq(), nil
	case "ir":
		return Ir(), nil
	case "is":
		return Is(), nil
	case "it":
		return It(), nil
	case "it-88":
		return ItEightyEight(), nil
	case "it-82":
		return ItEightyTwo(), nil
	case "it-23":
		return ItTwentyThree(), nil
	case "je":
		return Je(), nil
	case "jm":
		return Jm(), nil
	case "jo":
		return Jo(), nil
	case "jp":
		return Jp(), nil
	case "kanuri":
		return Kanuri(), nil
	case "ke":
		return Ke(), nil
	case "kg":
		return Kg(), nil
	case "kh":
		return Kh(), nil
	case "ki":
		return Ki(), nil
	case "kikuyu":
		return Kikuyu(), nil
	case "klingon":
		return Klingon(), nil
	case "km":
		return Km(), nil
	case "kn":
		return Kn(), nil
	case "kongo":
		return Kongo(), nil
	case "kp":
		return Kp(), nil
	case "kr":
		return Kr(), nil
	case "kurdistan":
		return Kurdistan(), nil
	case "kw":
		return Kw(), nil
	case "ky":
		return Ky(), nil
	case "kz":
		return Kz(), nil
	case "la":
		return La(), nil
	case "lb":
		return Lb(), nil
	case "lc":
		return Lc(), nil
	case "li":
		return Li(), nil
	case "lk":
		return Lk(), nil
	case "lr":
		return Lr(), nil
	case "ls":
		return Ls(), nil
	case "lt":
		return Lt(), nil
	case "lu":
		return Lu(), nil
	case "lv":
		return Lv(), nil
	case "ly":
		return Ly(), nil
	case "ma":
		return Ma(), nil
	case "malayali":
		return Malayali(), nil
	case "manipur":
		return Manipur(), nil
	case "maori":
		return Maori(), nil
	case "mc":
		return Mc(), nil
	case "md":
		return Md(), nil
	case "me":
		return Me(), nil
	case "mf":
		return Mf(), nil
	case "mg":
		return Mg(), nil
	case "mh":
		return Mh(), nil
	case "mizoram":
		return Mizoram(), nil
	case "mk":
		return Mk(), nil
	case "ml":
		return Ml(), nil
	case "mm":
		return Mm(), nil
	case "mn":
		return Mn(), nil
	case "mo":
		return Mo(), nil
	case "mp":
		return Mp(), nil
	case "mq":
		return Mq(), nil
	case "mr":
		return Mr(), nil
	case "ms":
		return Ms(), nil
	case "mt":
		return Mt(), nil
	case "mu":
		return Mu(), nil
	case "mv":
		return Mv(), nil
	case "mw":
		return Mw(), nil
	case "mx":
		return Mx(), nil
	case "my":
		return My(), nil
	case "mz":
		return Mz(), nil
	case "na":
		return Na(), nil
	case "nato":
		return Nato(), nil
	case "nc":
		return Nc(), nil
	case "ne":
		return Ne(), nil
	case "nf":
		return Nf(), nil
	case "ng":
		return Ng(), nil
	case "ni":
		return Ni(), nil
	case "nl":
		return Nl(), nil
	case "nl-fr":
		return NlFr(), nil
	case "no":
		return No(), nil
	case "northern-cyprus":
		return NorthernCyprus(), nil
	case "np":
		return Np(), nil
	case "nr":
		return Nr(), nil
	case "nu":
		return Nu(), nil
	case "nz":
		return Nz(), nil
	case "occitania":
		return Occitania(), nil
	case "olympics":
		return Olympics(), nil
	case "om":
		return Om(), nil
	case "otomi":
		return Otomi(), nil
	case "pa":
		return Pa(), nil
	case "pe":
		return Pe(), nil
	case "pf":
		return Pf(), nil
	case "pg":
		return Pg(), nil
	case "ph":
		return Ph(), nil
	case "pk":
		return Pk(), nil
	case "pk-jk":
		return PkJk(), nil
	case "pk-sd":
		return PkSd(), nil
	case "pl":
		return Pl(), nil
	case "pm":
		return Pm(), nil
	case "pn":
		return Pn(), nil
	case "pr":
		return Pr(), nil
	case "ps":
		return Ps(), nil
	case "pt":
		return Pt(), nil
	case "pt-30":
		return PtThirty(), nil
	case "pt-20":
		return PtTwenty(), nil
	case "pw":
		return Pw(), nil
	case "py":
		return Py(), nil
	case "qa":
		return Qa(), nil
	case "quechua":
		return Quechua(), nil
	case "re":
		return Re(), nil
	case "ro":
		return Ro(), nil
	case "rs":
		return Rs(), nil
	case "ru":
		return Ru(), nil
	case "ru-ba":
		return RuBa(), nil
	case "ru-ce":
		return RuCe(), nil
	case "ru-cu":
		return RuCu(), nil
	case "ru-da":
		return RuDa(), nil
	case "ru-ko":
		return RuKo(), nil
	case "ru-ta":
		return RuTa(), nil
	case "ru-ud":
		return RuUd(), nil
	case "rw":
		return Rw(), nil
	case "sa":
		return Sa(), nil
	case "sami":
		return Sami(), nil
	case "sb":
		return Sb(), nil
	case "sc":
		return Sc(), nil
	case "sd":
		return Sd(), nil
	case "se":
		return Se(), nil
	case "sg":
		return Sg(), nil
	case "sh":
		return Sh(), nil
	case "sh-ac":
		return ShAc(), nil
	case "sh-hl":
		return ShHl(), nil
	case "sh-ta":
		return ShTa(), nil
	case "si":
		return Si(), nil
	case "sj":
		return Sj(), nil
	case "sk":
		return Sk(), nil
	case "sl":
		return Sl(), nil
	case "sm":
		return Sm(), nil
	case "sn":
		return Sn(), nil
	case "so":
		return So(), nil
	case "somaliland":
		return Somaliland(), nil
	case "south-ossetia":
		return SouthOssetia(), nil
	case "soviet-union":
		return SovietUnion(), nil
	case "sr":
		return Sr(), nil
	case "ss":
		return Ss(), nil
	case "st":
		return St(), nil
	case "su":
		return Su(), nil
	case "sv":
		return Sv(), nil
	case "sx":
		return Sx(), nil
	case "sy":
		return Sy(), nil
	case "sz":
		return Sz(), nil
	case "ta":
		return Ta(), nil
	case "tc":
		return Tc(), nil
	case "td":
		return Td(), nil
	case "tf":
		return Tf(), nil
	case "tg":
		return Tg(), nil
	case "th":
		return Th(), nil
	case "tibet":
		return Tibet(), nil
	case "tj":
		return Tj(), nil
	case "tk":
		return Tk(), nil
	case "tl":
		return Tl(), nil
	case "tm":
		return Tm(), nil
	case "tn":
		return Tn(), nil
	case "to":
		return To(), nil
	case "torres-strait-islands":
		return TorresStraitIslands(), nil
	case "tr":
		return Tr(), nil
	case "transnistria":
		return Transnistria(), nil
	case "tt":
		return Tt(), nil
	case "tv":
		return Tv(), nil
	case "tw":
		return Tw(), nil
	case "tz":
		return Tz(), nil
	case "ua":
		return Ua(), nil
	case "ug":
		return Ug(), nil
	case "uk":
		return Uk(), nil
	case "um":
		return Um(), nil
	case "un":
		return Un(), nil
	case "united-nations":
		return UnitedNations(), nil
	case "us":
		return Us(), nil
	case "us-ak":
		return UsAk(), nil
	case "us-al":
		return UsAl(), nil
	case "us-ar":
		return UsAr(), nil
	case "us-az":
		return UsAz(), nil
	case "us-ca":
		return UsCa(), nil
	case "us-co":
		return UsCo(), nil
	case "us-dc":
		return UsDc(), nil
	case "us-fl":
		return UsFl(), nil
	case "us-ga":
		return UsGa(), nil
	case "us-hi":
		return UsHi(), nil
	case "us-in":
		return UsIn(), nil
	case "us-mo":
		return UsMo(), nil
	case "us-ms":
		return UsMs(), nil
	case "us-nc":
		return UsNc(), nil
	case "us-nm":
		return UsNm(), nil
	case "us-ri":
		return UsRi(), nil
	case "us-tn":
		return UsTn(), nil
	case "us-tx":
		return UsTx(), nil
	case "uy":
		return Uy(), nil
	case "uz":
		return Uz(), nil
	case "va":
		return Va(), nil
	case "vc":
		return Vc(), nil
	case "ve":
		return Ve(), nil
	case "vg":
		return Vg(), nil
	case "vi":
		return Vi(), nil
	case "vn":
		return Vn(), nil
	case "vo":
		return Vo(), nil
	case "vu":
		return Vu(), nil
	case "wf":
		return Wf(), nil
	case "wiphala":
		return Wiphala(), nil
	case "ws":
		return Ws(), nil
	case "xk":
		return Xk(), nil
	case "xx":
		return Xx(), nil
	case "ye":
		return Ye(), nil
	case "yi":
		return Yi(), nil
	case "yorubaland":
		return Yorubaland(), nil
	case "yt":
		return Yt(), nil
	case "za":
		return Za(), nil
	case "zm":
		return Zm(), nil
	case "zw":
		return Zw(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in circle_flags icon set", name)
	}
}
