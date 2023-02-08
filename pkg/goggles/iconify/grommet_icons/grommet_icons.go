package grommet_icons

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	accessibilityInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h7v6l-4 7M20 8h-7v6l4 7M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1 3h2v5h-2V8Z"/>`
	achievementInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M10.325 14.763A6.002 6.002 0 0 1 6 9V1h12v9M6 3H1v4c0 2.509 1.791 4 4 4h1m14.034-.115C21.742 10.49 23 9.103 23 7V3h-5m-8 16H5v4h11.5m0-13a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13Zm3.5 4l-4.5 4.5L13 16m-2.794-.576A4 4 0 0 0 8 19"/>`
	actionInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m1 23l3-3l-3 3ZM20 4l3-3l-3 3ZM9 11l3-3l-3 3Zm4 4l3-3l-3 3ZM10 5l9 9l1-1c2-2 4.053-5 0-9s-7-2-9 0l-1 1Zm-6 6l1-1l9 9l-1 1c-2 2-5 4.087-9 0s-2-7 0-9Z"/>`
	actionsInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 17.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm0-11V1m0 22v-5.5M1 12h5.5m11 0H23M4.437 4.437l4.125 4.125m6.876 6.876l4.124 4.124m0-15.125l-4.125 4.125m-6.874 6.876l-4.126 4.124"/>`
	adInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10h4v4h-4v-4Zm4 8h-4a4 4 0 0 1-4-4v-4a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v4m-4 8h-4a8 8 0 0 1-8-8v-4a8 8 0 0 1 8-8h4a8 8 0 0 1 8 8v4"/>`
	addInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22V2M2 12h20"/>`
	addCircleInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0-4V6m-6 6h12"/>`
	aedInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 8.4C2 4 5 3 7 3s4 2 5 3.5C13 5 15 3 17 3s5 1 5 5.4C22 15 12 21 12 21S2 15 2 8.4ZM12 6c-.5-.5-2 4-2 4h2v2.5L14 9h-2s1-6 5-6c-4 0-5 3-5 3Z"/>`
	aggregateInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 15h7V8a7 7 0 1 0-7 7Zm8-6H9v7a7 7 0 1 0 7-7Z"/>`
	aidInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 22h22V6H1v16ZM8 6h8V2H8v4Zm0 8h8m-4-4v8"/>`
	aidOptionInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 9v10V9Zm5 5H7h10ZM1 6.995C1 5.893 1.89 5 2.991 5H21.01C22.109 5 23 5.893 23 6.995v14.01C23 22.107 22.11 23 21.009 23H2.99A1.992 1.992 0 0 1 1 21.005V6.995ZM7 5V2.01C7 1.451 7.456 1 7.995 1h8.01c.55 0 .995.443.995 1.01V5"/>`
	alarmInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M21 13a9 9 0 0 1-9 9a9 9 0 0 1-9-9a9 9 0 0 1 9-9a9 9 0 0 1 9 9ZM5.5 19.5L2 23l3.5-3.5Zm13 0L22 23l-3.5-3.5ZM9 4c-.71-1.092-2.118-2-4-2c-2.1 0-4 1.9-4 4c0 1.882.908 3.29 2 4m18 0c1.092-.71 2-2.118 2-4c0-2.1-1.9-4-4-4c-1.882 0-3.29.908-4 2m-3 4v5l3 3"/>`
	alertInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 17v2m0-9v6m0-13L2 22h20L12 3Z"/>`
	amazonInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M11.992 23.94c-3.155-.021-5.956-1.061-8.46-2.955a16.663 16.663 0 0 1-2.446-2.282c-.027-.03-.072-.058-.075-.09c-.011-.117-.008-.234-.011-.35c.1.016.218.003.298.053c.67.417 1.314.878 2 1.264a18.612 18.612 0 0 0 6.378 2.192c1.18.18 2.365.24 3.55.17a17.755 17.755 0 0 0 6.486-1.648c.325-.151.644-.316.964-.479c.154-.079.307-.111.438.028c.139.146.133.322.036.485a1.11 1.11 0 0 1-.222.248c-1.841 1.633-3.995 2.636-6.396 3.103c-.832.161-1.684.217-2.527.321l-.013-.06m8.89-5.095c-.456.044-.961.094-1.467.14c-.11.01-.224.017-.334.01c-.128-.008-.148-.086-.078-.18a.658.658 0 0 1 .162-.137c.536-.359 1.141-.516 1.769-.6a4.51 4.51 0 0 1 1.773.096c.394.106.48.153.45.61c-.078 1.206-.466 2.303-1.34 3.183c-.045.046-.087.104-.143.126c-.074.029-.16.028-.24.04c.002-.073-.013-.152.009-.219c.208-.634.426-1.265.633-1.9a1.97 1.97 0 0 0 .073-.397c.051-.428-.12-.643-.574-.71c-.212-.03-.427-.039-.693-.062M13.997 9.42c-.86-.005-1.68.032-2.466.296c-.388.13-.781.295-1.115.527c-.923.639-1.163 1.59-1.087 2.65c.04.562.209 1.082.623 1.489c.588.578 1.579.654 2.374.184c.693-.41 1.107-1.042 1.375-1.779c.395-1.085.27-2.216.296-3.367m.632 6.353c-.309.28-.594.545-.884.804c-1.118.994-2.438 1.438-3.92 1.467c-.805.017-1.596-.044-2.349-.342c-1.342-.53-2.197-1.51-2.544-2.902c-.431-1.728-.286-3.392.737-4.902c.756-1.116 1.849-1.78 3.11-2.186c1.084-.35 2.205-.502 3.33-.628c.618-.07 1.236-.13 1.884-.199c-.02-.773.09-1.55-.181-2.297c-.241-.662-.75-1.003-1.414-1.155c-1.007-.23-2.147.145-2.688.938c-.184.27-.296.597-.395.913c-.139.445-.343.59-.805.535c-.902-.107-1.807-.196-2.71-.302c-.447-.052-.624-.307-.537-.742c.397-1.98 1.543-3.363 3.41-4.09c2.201-.86 4.452-.927 6.684-.106c1.85.68 2.865 2.035 3 4.013c.046.668.052 1.34.055 2.009c.007 1.732.016 3.464-.003 5.196c-.01.965.258 1.82.844 2.582c.141.183.284.367.403.564c.198.327.145.575-.141.822c-.737.636-1.472 1.273-2.208 1.911c-.404.352-.685.347-1.092.004a7.019 7.019 0 0 1-1.357-1.569c-.069-.107-.142-.21-.229-.338"/>`
	amexInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="m4.314 11.965l-.82-1.997l-.815 1.997h1.635Zm7.859 2.161l-.005-3.922l-1.736 3.922h-1.05L7.64 10.2v3.926H5.206l-.46-1.117H2.253l-.465 1.117h-1.3l2.144-5.008H4.41l2.036 4.742V9.118H8.4l1.567 3.397l1.439-3.397H13.4v5.008h-1.227Zm3.133-1.024v-.997h2.628v-1.022h-2.628v-.911h3.001l1.31 1.46l-1.368 1.47h-2.943Zm8.111 1.044h-1.556l-1.474-1.659l-1.532 1.659h-4.742v-5.01h4.815l1.473 1.642l1.523-1.642h1.564l-2.327 2.505l2.256 2.505Z"/>`
	analyticsInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20 7c1.25 1.67 2 3.75 2 6c0 5.52-4.48 10-10 10S2 18.52 2 13S6.48 3 12 3m0-2v12l9.6-7.2C19.41 2.89 15.92 1 12 1Z"/>`
	anchorInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm8 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM4 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8-11v15m-8-5.027C6.194 19.324 8.86 21 12 21c3.14 0 5.807-1.676 8-5.027M16 10H8"/>`
	androidInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M5.685 7.914h12.718v10.364c0 .615-.499 1.115-1.114 1.115h-1.274v3.156A1.44 1.44 0 0 1 14.587 24a1.44 1.44 0 0 1-1.43-1.451v-3.156h-2.225v3.156A1.44 1.44 0 0 1 9.502 24a1.44 1.44 0 0 1-1.429-1.451v-3.156H6.8c-.615 0-1.115-.5-1.115-1.115V7.914Zm-2.492-.085c-.797 0-1.443.656-1.443 1.466v5.727c0 .808.646 1.465 1.443 1.465s1.443-.657 1.443-1.465V9.295c0-.81-.646-1.466-1.443-1.466Zm15.21-.96H5.685C5.842 5.059 7.018 3.5 8.71 2.597L7.5.82a.525.525 0 1 1 .868-.59l1.318 1.936a7.204 7.204 0 0 1 4.717 0L15.721.23a.524.524 0 1 1 .867.59L15.38 2.596c1.692.902 2.866 2.461 3.023 4.274Zm-8.338-2.461a.703.703 0 1 0-1.406-.001a.703.703 0 0 0 1.406 0Zm5.454 0a.704.704 0 1 0-1.408 0a.704.704 0 0 0 1.408 0Zm5.378 3.42c-.797 0-1.444.656-1.444 1.466v5.729c0 .81.647 1.466 1.444 1.466c.797 0 1.441-.657 1.441-1.466v-5.73c0-.809-.644-1.466-1.441-1.466Z"/>`
	announceInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M11 15c3 0 8 4 8 4V3s-5 4-8 4v8Zm-6 0l3 8h4l-3-8m10-1a3 3 0 1 0 0-6m-8 11c1 0 3-1 3-3M2 11c0-3.111 1.791-4 4-4h5v8H6c-2.209 0-4-.889-4-4Z"/>`
	appleInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M15.3 3.832c.84-1.014 1.404-2.427 1.25-3.832c-1.208.049-2.67.805-3.535 1.819c-.777.898-1.457 2.335-1.273 3.712c1.346.105 2.722-.684 3.56-1.699m3.02 8.918c.034 3.632 3.186 4.841 3.22 4.857c-.025.085-.502 1.722-1.66 3.413c-1 1.462-2.038 2.919-3.674 2.949c-1.607.029-2.123-.953-3.961-.953c-1.836 0-2.41.923-3.932.982c-1.578.06-2.78-1.581-3.79-3.037c-2.06-2.98-3.635-8.42-1.52-12.092C4.054 7.045 5.932 5.89 7.97 5.861c1.55-.03 3.013 1.043 3.96 1.043c.948 0 2.726-1.29 4.595-1.101c.783.033 2.979.316 4.39 2.381c-.114.07-2.621 1.53-2.594 4.566"/>`
	appleAppStoreInnerSVG          = `<path fill="currentColor" d="M3 17a1 1 0 0 1 0-2h10.5c1 0 2 2 1.5 2H3Zm14 0a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2h-4ZM12.633 3.501a1 1 0 0 1 1.734.998L7.46 16.495a1 1 0 0 1-1.734-.997L12.633 3.5ZM4 18.5c.5-1 3.5-2 2.5-.28A852.88 852.88 0 0 1 4.867 21a1 1 0 0 1-1.734-.998L4 18.5ZM9.133 4.499a1 1 0 1 1 1.734-.998L12.61 6.53a1 1 0 1 1-1.733.998L9.133 4.499ZM13 11.5c-.898-1.5 0-4.5.716-3.004L20.366 20a1 1 0 0 1-1.733.998L13 11.5Z"/>`
	appsInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 5h2V3h-2v2Zm-8 0h2V3h-2v2ZM3 5h2V3H3v2Zm16 8h2v-2h-2v2Zm-8 0h2v-2h-2v2Zm-8 0h2v-2H3v2Zm16 8h2v-2h-2v2Zm-8 0h2v-2h-2v2Zm-8 0h2v-2H3v2Z"/>`
	appsRoundedInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 6.2c0-1.12 0-1.68.218-2.108a2 2 0 0 1 .874-.874C4.52 3 5.08 3 6.2 3h.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C10 4.52 10 5.08 10 6.2v.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C8.48 10 7.92 10 6.8 10h-.6c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C3 8.48 3 7.92 3 6.8v-.6zm11 0c0-1.12 0-1.68.218-2.108a2 2 0 0 1 .874-.874C15.52 3 16.08 3 17.2 3h.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C21 4.52 21 5.08 21 6.2v.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C19.48 10 18.92 10 17.8 10h-.6c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C14 8.48 14 7.92 14 6.8v-.6zm-11 11c0-1.12 0-1.68.218-2.108a2 2 0 0 1 .874-.874C4.52 14 5.08 14 6.2 14h.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C10 15.52 10 16.08 10 17.2v.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C8.48 21 7.92 21 6.8 21h-.6c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C3 19.48 3 18.92 3 17.8v-.6zm11 0c0-1.12 0-1.68.218-2.108a2 2 0 0 1 .874-.874C15.52 14 16.08 14 17.2 14h.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C21 15.52 21 16.08 21 17.2v.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C19.48 21 18.92 21 17.8 21h-.6c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C14 19.48 14 18.92 14 17.8v-.6z"/>`
	archiveInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 23h18V6H3v17Zm6-9h6v-4H9v4ZM1 6h22V1H1v5Z"/>`
	archlinuxInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M11.995 0c-1.068 2.619-1.712 4.332-2.901 6.873c.729.773 1.624 1.673 3.077 2.69c-1.562-.643-2.628-1.289-3.425-1.959C7.224 10.78 4.84 15.304 0 24c3.804-2.196 6.752-3.55 9.5-4.066a6.964 6.964 0 0 1-.18-1.63l.004-.121c.06-2.437 1.328-4.311 2.83-4.184c1.501.127 2.668 2.207 2.608 4.644c-.011.459-.063.9-.153 1.309c2.717.532 5.634 1.882 9.387 4.048c-.74-1.362-1.4-2.59-2.031-3.76c-.994-.77-2.03-1.771-4.143-2.856c1.452.377 2.493.813 3.303 1.3C14.713 6.746 14.195 5.16 11.995 0Z"/>`
	articleInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 7h3v4h-3V7Zm-7 8h11M9 11h4M9 7h4M6 18.5a2.5 2.5 0 1 1-5 0V7h5.025M6 18.5V3h17v15.5a2.5 2.5 0 0 1-2.5 2.5h-17"/>`
	arubaInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M12.11 17.302c-3.074 0-5.602-2.46-5.602-5.465c0-3.006 2.528-5.465 5.602-5.465c3.074 0 5.601 2.46 5.601 5.465s-2.527 5.465-5.601 5.465ZM12.11 2C6.508 2 2 6.44 2 11.837c0 5.465 4.508 9.836 10.11 9.836c2.323 0 4.44-.751 6.148-2.049c1.025 1.708 3.962 2.05 3.962 2.05v-9.837C22.22 6.44 17.71 2 12.11 2Z"/>`
	ascendInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m2 8l6-6l6 6m-3 13h11m-11-4h8m-8-4h5M8 2v20"/>`
	ascendingInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="m12.08 7.286l.354-.353l.354.353L17 11.498l-.707.708l-3.358-3.359V17h-1V8.847l-3.359 3.359l-.707-.708l4.212-4.212Z" clip-rule="evenodd"/>`
	assistListeningInnerSVG        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 21c.757.667 1.424 1 2 1c2 0 3-1 3-3c0-1.333.667-2.667 2-4c1.267-1.267 2-3.067 2-5a7 7 0 0 0-14 0m11 0a4 4 0 1 0-8 0M3 20l5-6l1 4l5-6"/>`
	atmInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 10H1V6h22v4h-3M6 6h12v14.006A2.003 2.003 0 0 1 15.991 22H8.01A2.002 2.002 0 0 1 6 20.006V6Zm7 12h3m-3-3h3m-3-3h3m-6 10V6M3 2h3m2 0h3m2 0h3m2 0h3"/>`
	attachmentInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="m22 12l-9 9c-6 6-15-3-9-9l9-9c4-4 10 2 6 6l-9 9c-2 2-5-1-3-3l9-9"/>`
	attractionInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="m13 2l1.5.5L13 3V2Zm1 16l1 5h-2l1-5Zm0-1.5l2 6.5h-4l2-6.5ZM12 6l9 5v2H3v-2l9-5Zm-7.5 7h15c0 4.167 1.5 10 1.5 10H3s1.5-5.833 1.5-10Zm0 0h15c0 4.167 1.5 10 1.5 10H3s1.5-5.833 1.5-10Z"/>`
	babyInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 15h4s-1 1.5-2 1.5s-2-1.5-2-1.5Zm2-9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm6 6l-4-3m0 6l3 2.5l-2.5 2.5M6 12l4-3m0 6l-2.5 2.75L10 20m0-11h4v3h-4V9Z"/>`
	backTenInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3.111 7.556C4.67 4.267 8.07 2 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m0-8v4h4m3 8V9l-2 .533M17 12c0-2-1-3.5-2.5-3.5S12 10 12 12s1 3.5 2.5 3.5S17 14 17 12Zm-2.5-3.5C16.925 8.5 17 11 17 12s0 3.5-2.5 3.5S12 13 12 12s.059-3.5 2.5-3.5Z"/>`
	barInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 5h16v18H7V5Zm0 0h16V3c0-1.105-.895-2-1.994-2H8.994A1.992 1.992 0 0 0 7 3v2ZM1 8.009C1 6.899 1.898 6 2.998 6H7v12H2.998A2.005 2.005 0 0 1 1 15.991V8.01ZM11 8v10m4-10v10m4-10v10"/>`
	barChartInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M0 22h24M22 2h-4v16h4V2ZM6 6H2v12h4V6Zm8 12h-4v-8h4v8Z"/>`
	basketInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 12h20l-2 11H4L2 12Zm18-4l-6-7M4 8l6-7M1 8h22v4H1V8Zm7 7v5m8-5v5m-4-5v5"/>`
	beaconInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm0 0v11M7.05 4.05A6.978 6.978 0 0 0 5 9c0 1.933.784 3.683 2.05 4.95m9.9 0A6.978 6.978 0 0 0 19 9a6.978 6.978 0 0 0-2.05-4.95M4.222 1.222A10.966 10.966 0 0 0 1 9c0 3.037 1.231 5.787 3.222 7.778m15.556 0A10.966 10.966 0 0 0 23 9c0-3.038-1.231-5.788-3.222-7.778"/>`
	bikeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM5 6h5m9 9L16 5h-3M9 9l-4 6h7c0-3 2-6 5-6H9Zm0 0L7 6"/>`
	bitcoinInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M23.64 14.905c-1.602 6.429-8.114 10.342-14.544 8.738C2.669 22.041-1.244 15.528.359 9.1C1.962 2.67 8.473-1.244 14.902.36c6.43 1.603 10.342 8.116 8.739 14.546Zm-6.349-4.613c.24-1.597-.977-2.456-2.64-3.028l.54-2.164l-1.317-.328l-.525 2.106a55.083 55.083 0 0 0-1.055-.248l.53-2.12l-1.317-.328l-.54 2.162a44.05 44.05 0 0 1-.84-.197l.001-.007l-1.816-.453l-.35 1.406s.977.224.956.238c.534.133.63.486.614.766l-.615 2.464c.037.01.085.023.137.044l-.139-.035l-.86 3.453c-.066.162-.231.405-.604.312c.013.02-.957-.239-.957-.239l-.654 1.508l1.713.427c.32.08.632.163.94.242l-.546 2.188l1.316.328l.54-2.164c.359.097.707.187 1.049.272l-.538 2.155l1.317.328l.544-2.184c2.246.425 3.934.253 4.645-1.778c.573-1.635-.029-2.578-1.21-3.193c.86-.198 1.508-.764 1.681-1.933Zm-3.008 4.219c-.407 1.635-3.16.75-4.053.53l.723-2.9c.893.223 3.755.664 3.33 2.37Zm.407-4.243c-.371 1.487-2.663.732-3.406.547l.656-2.63c.743.186 3.137.532 2.75 2.083Z"/>`
	bladesHorizontalInnerSVG       = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 17h4m3 0h4m3 0h4M4 21h2a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1zm7 0h2a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1zm7 0h2a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1z"/>`
	bladesVerticalInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17 3v4m0 3v4m0 3v4m4-17v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1zm0 7v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1zm0 7v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1z"/>`
	blockQuoteInnerSVG             = `<path fill="currentColor" d="M.78 8.89c0-3.07 1.53-4.3 4.3-4.34L5.38 6C3.78 6.17 3 7 3.1 8.31h1.44V12H.78Zm5.9 0c0-3.07 1.53-4.3 4.3-4.34l.3 1.45C9.68 6.17 8.89 7 9 8.31h1.44V12H6.68Zm10.26 6.22c0 3.07-1.53 4.3-4.3 4.34L12.35 18c1.6-.16 2.39-1 2.28-2.3h-1.45V12h3.76Zm5.9 0c0 3.07-1.53 4.3-4.3 4.34l-.3-1.45c1.6-.16 2.39-1 2.28-2.3h-1.44V12h3.76Z"/>`
	blogInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 16a3 3 0 1 0 0 6a3 3 0 0 0 0-6ZM5 1c9.925 0 18 8.075 18 18m-5 0c0-7.168-5.832-13-13-13m8 13c0-4.411-3.589-8-8-8m-3 0v8"/>`
	bluetoothInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="m7 7l10 9l-5 4V4l5 4l-10 8"/>`
	boldInnerSVG                   = `<path fill="currentColor" d="M14 11.57a7.93 7.93 0 0 1 3.11 1.25a3.32 3.32 0 0 1 1.28 2.71A3.58 3.58 0 0 1 17 18.42a7.9 7.9 0 0 1-5 1.39H4.07v-.42a3.57 3.57 0 0 0 1.46-.2a1.17 1.17 0 0 0 .54-.52a4.75 4.75 0 0 0 .15-1.58V7a4.81 4.81 0 0 0-.15-1.6a1.13 1.13 0 0 0-.54-.52a3.67 3.67 0 0 0-1.46-.2v-.42h7.45a10.57 10.57 0 0 1 3.78.48a3.94 3.94 0 0 1 1.75 1.42a3.52 3.52 0 0 1 .64 2a2.86 2.86 0 0 1-.81 2A5.84 5.84 0 0 1 14 11.57Zm-4.17.58v5.56a1.27 1.27 0 0 0 .32.93a1.27 1.27 0 0 0 .93.31a3.57 3.57 0 0 0 1.69-.41A2.79 2.79 0 0 0 14 17.37a3.61 3.61 0 0 0 .41-1.73a3.81 3.81 0 0 0-.5-2a2.72 2.72 0 0 0-1.39-1.21a7.52 7.52 0 0 0-2.67-.28Zm0-.89a5.92 5.92 0 0 0 2.4-.37a2.73 2.73 0 0 0 1.19-1a3.17 3.17 0 0 0 .41-1.7a3.18 3.18 0 0 0-.41-1.69a2.59 2.59 0 0 0-1.16-1a6.2 6.2 0 0 0-2.43-.33Z"/>`
	bookInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M10 1v10l3-2l3 2V1M5.5 18a2.5 2.5 0 1 0 0 5H22M3 20.5v-17A2.5 2.5 0 0 1 5.5 1H21v17.007H5.492M20.5 18a2.5 2.5 0 1 0 0 5"/>`
	bookmarkInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 1v21l7-5l7 5V1z"/>`
	bottomCornerInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 20h12V8"/>`
	brailleInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm10 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zM7 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm10 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2zM7 20a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm10-10a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/>`
	briefcaseInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 6h22v15H1V6Zm5 0v15M18 6v15M8 6V3h8v3"/>`
	brushInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9.312 11.73a5.001 5.001 0 0 0-5.362 1.12c-1.953 1.952-1.414 8.485-1.414 8.485s6.532.538 8.485-1.415a5.001 5.001 0 0 0 1.12-5.362L22.334 4.364a1.997 1.997 0 0 0 0-2.828a1.995 1.995 0 0 0-2.828 0L9.312 11.729Zm1.002-1.617l1.838 1.838l1.697 1.697"/>`
	bugInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M23 20c-1.38-2.09-3-3-4-3M5 17c-1 0-2.62.91-4 3M19 9c3 0 4-3 4-3M1 6s1 3 4 3m14 4h5h-5ZM5 13H0h5Zm7 10V12v11Zm0 0c-4 0-7-3-7-7V9s3-2.012 7-2c4 .012 7 2 7 2v7c0 4-3 7-7 7ZM7 8V6c0-2.76 2.24-5 5-5s5 2.24 5 5v2"/>`
	bundleInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17 14H7h10Zm0-11h6v10h-6M1 13v4h6m10 0h6v-4M1 17v4h6m16-4v4h-6M7 22h10V2H7v20Zm0-9H1V3h6"/>`
	busInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 12h18v8H3v-8Zm0-8c0-1.105.893-2 1.995-2h14.01C20.107 2 21 2.887 21 4v8H3V4Zm0 16h3v2.001a.996.996 0 0 1-.999.999H3.999A.996.996 0 0 1 3 22.001V20Zm15 0h3v2.001a.996.996 0 0 1-.999.999h-1.002a.996.996 0 0 1-.999-.999V20ZM7 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm10 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM12 6v6M1 5v8m22-8v8m-13 3h4M3 6h18"/>`
	businessServiceInnerSVG        = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 23H1V5h22v18h-7M8 5V1h8v4M9 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm5.008 1.876a4 4 0 1 0-1.244-7.193m-5.062 5.043A4 4 0 0 0 10 23a4 4 0 0 0 1.401-7.748"/>`
	cafeteriaInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 1v7c0 1.657-1.347 3-2.997 3H5.997A3.002 3.002 0 0 1 3 8V1m3 6V1m3 6V1M6 11v10.504C6 22.33 6.666 23 7.5 23c.828 0 1.5-.68 1.5-1.496V11m6 7v3.5a1.5 1.5 0 1 0 3-.005V15s3 0 3-3V7c0-3-2-5-6-5v16Z"/>`
	calculatorInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M22 23H2V1h20v22Zm-6-6h2v2h-2v-2Zm-5 0h2v2h-2v-2Zm5-5h2v2h-2v-2Zm-5 0h2v2h-2v-2Zm-5 5h2v2H6v-2Zm0-5h2v2H6v-2Zm12-3H6V5h12v4Z"/>`
	calendarInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 5h20v17H2V5Zm16 0V1M6 5V1m-4 9h20"/>`
	cameraInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 21V7h5l2-4h8l2 4h5v14H1Zm11-3a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/>`
	capacityInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M1 19h22V5H1v14Zm3-3h3V8H4v8Zm6 0h3V8h-3v8Z"/>`
	carInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 10.997c0-.55.44-.997 1.002-.997h19.996c.553 0 1.002.453 1.002.997v6.006c0 .55-.44.997-1.002.997H2.002A1.004 1.004 0 0 1 1 17.003v-6.006ZM6 2h12l4 8H2l4-8Zm6 8.5L15 5M3 18h3v2.99c0 .558-.443 1.01-.999 1.01H3.999A.999.999 0 0 1 3 20.99V18Zm15 0h3v2.99c0 .558-.443 1.01-.999 1.01h-1.002A.999.999 0 0 1 18 20.99V18ZM5 15a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm14 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-7-2v2m2-2v2m-4-2v2"/>`
	caretDownInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M22 8L12 20L2 8z"/>`
	caretDownFillInnerSVG          = `<path fill="currentColor" d="M18 9H6l6 6l6-6Z"/>`
	caretLeftFillInnerSVG          = `<path fill="currentColor" d="M15 18V6l-6 6z"/>`
	caretNextInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="m6 2l12 10L6 22z"/>`
	caretPreviousInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 2L6 12l12 10z"/>`
	caretRightFillInnerSVG         = `<path fill="currentColor" d="M9 6v12l6-6z"/>`
	caretUpInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M22 16L12 4L2 16z"/>`
	caretUpFillInnerSVG            = `<path fill="currentColor" d="M6 15h12l-6-6l-6 6Z"/>`
	cartInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 5h17l-2 9H7L4 2H0m7 12l1 4h13m-2 5a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM9 23a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`
	catalogInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 6L1 4.5v13.943L12 23l11-4.557V4l-4 2M5 16V2l7 3l7-3v14l-7 3l-7-3Zm6.95-11v14"/>`
	catalogOptionInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5.5 18a2.5 2.5 0 1 0 0 5H22M3 20.5v-17A2.5 2.5 0 0 1 5.5 1H21v17.007H5.492M20.5 18a2.5 2.5 0 1 0 0 5"/>`
	centosInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M5.161 6.25L3.249 8.164V3.227h4.935L6.24 5.172l-.147.146l.147.147l4.514 4.516v.738h-.83L5.457 6.251l-.15-.15l-.146.15Zm5.594 3.142V3.227H8.772l-2.09 2.091l4.073 4.074ZM5.307 6.693L3.25 8.753v1.966h6.083L5.307 6.693Zm7.964 2.67l4.044-4.045l-2.119-2.12h-1.925v6.165Zm4.636-4.045l-.15.147l-4.486 4.49v.764h.786l4.477-4.478l.146-.146l.15.146l1.96 1.96V3.2h-5.005l1.973 1.973l.15.146Zm-6.736-2.51v7l.828.828l.856-.856V2.782h1.924L12 0L9.192 2.807h1.98Zm3.478 7.911h6.141V8.794l-2.11-2.11l-4.031 4.035Zm-4.025 1.292l-.872-.872h-6.92V9.17L0 12.002l2.832 2.832V12.81h6.993l.8-.799Zm2.647 2.68v6.056h1.964l2.045-2.046l-4.009-4.01ZM5.317 17.32l4.092-4.093h-6.16v2.024l2.068 2.069Zm13.364 0l2.11-2.11v-1.983h-6.2l4.09 4.093Zm2.529-2.53l2.787-2.788l-2.787-2.788v1.925h-6.98l-.863.863l.808.808h7.035v1.98Zm-2.53 3.118l-.146-.146L14 13.227h-.729v.875l4.452 4.452l.146.147l-.146.15l-1.897 1.896h4.964v-4.945l-1.96 1.96l-.15.146Zm-5.825 3.256v-6.89l-.872-.872l-.799.799v6.963h-2.04L11.984 24l2.835-2.836h-1.964Zm-6.762-2.463l.147-.147l4.527-4.528v-.8h-.77l-4.534 4.536l-.146.146l-.147-.146l-1.92-1.922v4.907h4.887L6.24 18.85l-.147-.15Zm4.674-4.083L6.682 18.7l2.046 2.046h2.04v-6.13Z"/>`
	certificateInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M15 19H2V1h16v4m0 0a5 5 0 1 1 0 10a5 5 0 0 1 0-10zm-3 9v8l3-2l3 2v-8M5 8h6m-6 3h5m-5 3h2M5 5h2"/>`
	channelInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 9h20v13H2V9Zm19-7l-8 7h-2L3 2"/>`
	chapterAddInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 14V5h14v15h-8m8-4h4V1H9v4M5 16v8m4-4H1"/>`
	chapterNextInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 3.5V20l9-6v6l11-8l-11-8v6L1 3.5ZM22 2v20V2Z"/>`
	chapterPreviousInnerSVG        = `<path fill="none" stroke="currentColor" stroke-width="2" d="M23 3.5V20l-9-6v6L3 12l11-8v6l9-6.5ZM2 2v20V2Z"/>`
	chatInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 1H5a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h12l6 5V5a4 4 0 0 0-4-4Zm0 6H6m13 4h-9"/>`
	chatOptionInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 7V1h14v10h-3v5l-5-4M1 7h14v11H9l-5 4v-4H1V7Z"/>`
	checkboxInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 2h20v20H2z"/>`
	checkboxSelectedInnerSVG       = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 2h20v20H2V2Zm3 11l5 4l9-11"/>`
	checkmarkInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="m2 14l7 6L22 4"/>`
	chromeInnerSVG                 = `<g fill="currentColor" fill-rule="evenodd"><path d="M11.973 16.414A4.32 4.32 0 0 1 7.66 12.1a4.32 4.32 0 0 1 4.314-4.315a4.32 4.32 0 0 1 4.315 4.315a4.32 4.32 0 0 1-4.315 4.314Z"/><path d="M13.791 17.181a5.507 5.507 0 0 1-5.38-.926a5.496 5.496 0 0 1-1.718-2.505l-.002-.006L1.936 5.51A11.892 11.892 0 0 0 .23 14.35a11.896 11.896 0 0 0 5.08 7.627a11.894 11.894 0 0 0 4.557 1.84l3.924-6.635Z"/><path d="M22.76 6.707a12.08 12.08 0 0 0-20.185-2.25l4.016 6.956a5.393 5.393 0 0 1 5.274-4.706H22.76Z"/><path d="M11.926 24c3.01 0 5.891-1.129 8.11-3.178a11.932 11.932 0 0 0 3.816-7.893a12.05 12.05 0 0 0-.744-5.144h-7.856a5.506 5.506 0 0 1 2.09 4.34a5.529 5.529 0 0 1-1.182 3.381l-5.008 8.47c.258.016.518.024.774.024Z"/></g>`
	circleAlertInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 14V6m0 12v-2m0-14C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2z"/>`
	circleInformationInnerSVG      = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0-12v8m0-12v2"/>`
	circlePlayInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm-2.5-6.5l6-3.5l-6-3.5v7Zm1-2l2-1.5l-2-1.5v3Z"/>`
	circleQuestionInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0-7v-1c0-1 0-1.5 1-2s2-1 2-2.5c0-1-1-2.5-3-2.5s-3 1.264-3 3m3 6v2"/>`
	clearInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10ZM5 5l14 14"/>`
	clearOptionInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M10 4a2 2 0 1 1 4 0v6h6v4H4v-4h6V4ZM4 14h16v8H4v-8Zm12 8v-5.635M8 22v-5.635M12 22v-5.635"/>`
	cliInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h22v22H1V1Zm0 4h22M5 1v4m6 11h8M5 10l3 3l-3 3"/>`
	clipboardInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 3h5v20H3V3h5m0-2h8v5H8V1Z"/>`
	clockInnerSVG                  = `<path fill="currentColor" d="M13 7a1 1 0 1 0-2 0v5a1 1 0 0 0 .4.8l4 3a1 1 0 0 0 1.2-1.6L13 11.5V7z"/><path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1zM3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0z" clip-rule="evenodd"/>`
	cloneInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 23h16V7H7v16ZM17 4V1h-3M1 14v3h3m-3-5V6v6ZM4 1H1v3m5-3h6h-6Z"/>`
	closeInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="m3 3l18 18M3 21L21 3"/>`
	closedCaptionInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 12c0-7 1.5-8 11-8s11 1 11 8s-1.5 8-11 8s-11-1-11-8Zm4.25 2c0 1.5.75 2 2.5 2s2.5-.5 2.5-2h-.271c0 1.25-1 2-2.229 2c-1.23 0-2.229-.75-2.229-2v-4C5.5 8.75 6.5 8 7.75 8s2.25.75 2.229 2h.271c0-1.25-1.021-2-2.5-2s-2.5.75-2.5 2v4Zm8 0c0 1.5.75 2 2.5 2s2.5-.5 2.5-2h-.271c0 1.25-1 2-2.229 2c-1.23 0-2.229-.75-2.229-2v-4c-.021-1.25.979-2 2.229-2s2.25.75 2.229 2h.271c0-1.25-1.021-2-2.5-2s-2.5.75-2.5 2v4Z"/>`
	cloudInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 17v1c0 3-2 4-5 4h-2c-3 0-5-1-5-4v-1A5 5 0 0 1 6 7h6M6 7V6c0-3 2-4 5-4h2c3 0 5 1 5 4v1a5 5 0 0 1 0 10h-6"/>`
	cloudComputerInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 6V5c0-3 1.5-4 4-4h4c2.5 0 4 1.5 4 4v1c3 0 5 2 5 5s-2 5-5 5M14 6H6c-3 0-5 1.5-5 5s2 5 5 5m2 3h8v-7H8v7Zm4 0v4v-4Zm-3 4h6h-6Z"/>`
	cloudDownloadInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 17A5 5 0 0 1 6 7h6M6 7V6c0-3 2-4 5-4h2c3 0 5 1 5 4v1a5 5 0 0 1 0 10m-6-5v9m-4-4l4 4l4-4"/>`
	cloudSoftwareInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 23h8V12H8v11Zm0-7h8m-4-4v4M6 6V5c0-3 1.5-4 4-4h4c2.5 0 4 1.5 4 4v1c3 0 5 2 5 5s-2 5-5 5M14 6H6c-3 0-5 1.5-5 5s2 5 5 5"/>`
	cloudUploadInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 17A5 5 0 0 1 6 7h6M6 7V6c0-3 2-4 5-4h2c3 0 5 1 5 4v1a5 5 0 0 1 0 10m-6-4v9m-4-6l3.965-4.035L16 16"/>`
	cloudlinuxInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M14.068 5.79c.05.046.095.092.142.139c.314.323.609.674.881 1.055l.001.001c.266.381.51.792.744 1.234c.046-1.496-.243-2.81-.914-4.036c-.307-.563-.618-1.141-1.14-1.532c-.749-.562-1.596-.944-2.502-1.185a8.126 8.126 0 0 0-1.242-.229C6.917.809 3.77 2.365 2.093 5.477c.188.05.268-.111.376-.198c3.53-2.871 8.406-2.579 11.539.451l.06.06Zm-3.528 10.9c-.575-.335-1.099-.74-1.638-1.185c.545 1.559 2.637 3.432 4.421 3.92c-.073-.06-.14-.127-.22-.178c-.908-.585-1.654-1.338-2.29-2.206c-.087-.12-.153-.282-.273-.351ZM23.593 9.6c.015-.065.018-.123-.068-.137l-.013.138l-.028.327a8.633 8.633 0 0 1-10.866 7.614c-.174-.047-.346-.097-.519-.145c1.151 1.294 2.568 2.07 4.244 2.387c.563.106 1.108.09 1.662-.06c2.074-.567 3.706-1.75 4.812-3.583c1.235-2.048 1.52-4.243.776-6.542ZM14.47 20.53c-2.821-.515-4.856-2.149-6.25-4.638a7.351 7.351 0 0 0 1.16 4.367c.256.403.498.832.908 1.1c2.679 1.743 5.467 1.921 8.32.483c1.49-.75 2.59-1.927 3.324-3.44c-2.166 1.853-4.644 2.641-7.462 2.128ZM9.766.5c5.73.51 9.29 5.722 7.69 11.262c-.018.063-.062.127.047.168c1.308-1.283 2.121-2.818 2.357-4.65a2.62 2.62 0 0 0-.04-.935C19.26 3.91 17.88 2.08 15.678.922c-1.922-1.01-3.952-1.2-6.038-.523c-.065-.012-.127-.017-.137.074l.141.013c.04.006.081.012.122.016Zm6.969 11.226l-.002-.038l-.002-.016l-.006.016l.01.038Zm3.816-2.163c-.515 2.819-2.154 4.848-4.747 6.253c1.576.038 2.942-.281 4.202-1.015c.508-.295 1.037-.583 1.382-1.076c2.771-3.958 1.356-9.512-2.99-11.658c1.876 2.169 2.671 4.661 2.153 7.496Zm-1.094 1.196a.248.248 0 0 0-.155.105a8.36 8.36 0 0 1-2.433 2.46a.53.53 0 0 0-.174.169c-.315.577-.731 1.082-1.121 1.608c1.866-.936 3.163-2.38 3.883-4.342ZM16.36 5.667c-.117.098-.041.179-.024.247c.235.94.335 1.904.153 2.844c-.196 1.014.162 1.945.242 2.915c.247-.55.329-1.137.388-1.727c.15-1.492-.115-2.912-.76-4.279Zm-1.836 10.702a.692.692 0 0 0-.323.01c-.58.156-1.176.228-1.772.293c-.036-.017-.077-.03-.086.021c-.003.014.052.037.08.056c1.993.643 3.93.468 5.854-.37c-.053-.036-.073-.061-.096-.063a.378.378 0 0 0-.121.01a8.324 8.324 0 0 1-3.536.043Zm-5.65-.855a2.158 2.158 0 0 1-.034-.02c.009.01.018.017.028.026l.006-.006Zm-.486-6.56c.037-.014.076-.034.046-.077c-.008-.01-.063.011-.097.018a.41.41 0 0 0 .05.058ZM5.496 21.848c.034.056.073.1.143.048l-.087-.106a.448.448 0 0 0-.056.058ZM10.65 4.51c.967.61 1.825 1.352 2.447 2.307c.347.534.73.595 1.3.4c-1.003-1.293-2.218-2.219-3.747-2.707ZM8.337 8.895c-1.863.954-3.11 2.448-3.877 4.4c.063-.011.095-.008.112-.023a.38.38 0 0 0 .079-.093a8.322 8.322 0 0 1 2.47-2.53a.69.69 0 0 0 .221-.235c.3-.521.67-.993 1.046-1.46a.41.41 0 0 1-.05-.059ZM7.6 18.103a.25.25 0 0 0 .036-.184a8.364 8.364 0 0 1-.02-3.46a.53.53 0 0 0 .004-.242c-.186-.63-.248-1.282-.345-1.93c-.658 1.982-.554 3.92.325 5.816ZM6.09 7.61c1.056-.229 2.117-.224 3.18-.06c.147.024.308.091.441.056c.644-.17 1.3-.255 1.997-.32c-1.488-.718-4.292-.563-5.898.354c.094-.01.19-.01.28-.03Zm-.538 14.18l-.212-.25a8.633 8.633 0 0 1 2.3-13.068c.155-.09.313-.176.47-.264c-1.73-.101-3.281.352-4.69 1.313c-.472.323-.847.72-1.132 1.218c-1.066 1.867-1.384 3.857-.87 5.936c.576 2.321 1.926 4.074 4.078 5.174a.448.448 0 0 1 .056-.058Zm2.119-2.068c-1.63-2.357-1.906-4.951-1.065-7.779c-1.141 1.09-1.881 2.28-2.254 3.69c-.15.568-.321 1.145-.216 1.738c.839 4.759 5.767 7.686 10.357 6.13c-2.86-.208-5.184-1.408-6.822-3.779ZM.445 14.447C-.105 12.86-.16 11.25.363 9.664c1-3.033 3.098-4.879 6.225-5.54c.479-.101.954.032 1.419.136a7.355 7.355 0 0 1 3.908 2.267c-2.745-.774-5.34-.49-7.697 1.14c-2.356 1.63-3.551 3.94-3.773 6.78Z"/>`
	clusterInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 9h8V1H8v8ZM1 23h8v-8H1v8Zm14 0h8v-8h-8v8ZM5 15l3-6l-3 6Zm5 4h4h-4Zm6-10l3 6l-3-6Z"/>`
	coatCheckInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 11l10.155 6.462c.467.297.845.981.845 1.547v1.982A1 1 0 0 1 21.998 22H2.002A1 1 0 0 1 1 20.99v-1.98c0-.558.379-1.251.845-1.548L12 11Zm3-6a3 3 0 0 0-6 0c0 .932.411 1.802 1.09 2.314C11 8 12 8 12 9.5V11"/>`
	codeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="m9 22l6-20m2 15l5-5l-5-5M7 17l-5-5l5-5"/>`
	codeSandboxInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="m12 1.5l-9 5v11l9 5l9-5v-11l-9-5Zm0 21v-11m9-5l-9 5m0 0l-9-5m18 11V12l-4.5 2.5V20l4.5-2.5Zm-18 0V12l4.5 2.5V20L3 17.5Zm9-16L7.5 4L12 6.5L16.5 4L12 1.5Z"/>`
	codepenInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M12 22.03c-5.53 0-10.03-4.5-10.03-10.03C1.97 6.47 6.47 1.97 12 1.97c5.53 0 10.03 4.5 10.03 10.03c0 5.53-4.5 10.03-10.03 10.03M12 0C5.372 0 0 5.373 0 12c0 6.628 5.372 12 12 12s12-5.372 12-12c0-6.627-5.372-12-12-12m6.144 13.067L16.55 12l1.595-1.067v2.134Zm-5.506 4.524v-2.975l2.764-1.849l2.232 1.493l-4.996 3.33ZM12 13.509L9.745 12L12 10.492L14.255 12L12 13.51Zm-.638 4.082L6.366 14.26l2.232-1.493l2.764 1.85v2.974Zm-5.506-6.658L7.45 12l-1.595 1.067v-2.134Zm5.506-4.523v2.974l-2.764 1.85L6.366 9.74l4.996-3.33Zm1.276 0l4.996 3.33l-2.232 1.493l-2.764-1.85V6.41Zm6.776 3.246l-.005-.027a.624.624 0 0 0-.011-.054l-.01-.03a.617.617 0 0 0-.052-.12l-.019-.03a.568.568 0 0 0-.08-.101l-.026-.025a.728.728 0 0 0-.036-.03l-.029-.022l-.01-.008l-6.782-4.521a.637.637 0 0 0-.708 0l-6.782 4.52l-.01.009l-.03.022a.578.578 0 0 0-.035.03c-.01.008-.017.016-.026.025a.553.553 0 0 0-.099.13a.594.594 0 0 0-.021.043l-.014.03c-.007.016-.012.032-.017.047l-.01.031c-.004.018-.008.036-.01.054l-.006.027a.613.613 0 0 0-.006.083v4.522a.57.57 0 0 0 .006.083l.005.028l.011.053l.01.031c.005.016.01.031.017.047l.014.03a.755.755 0 0 0 .067.111a.422.422 0 0 0 .053.062l.026.025a.545.545 0 0 0 .065.052l.01.008l6.782 4.522a.637.637 0 0 0 .708 0l6.782-4.522l.01-.008a.711.711 0 0 0 .065-.052c.01-.008.017-.016.026-.025a.611.611 0 0 0 .032-.034l.021-.028a.568.568 0 0 0 .027-.039l.019-.029a.574.574 0 0 0 .021-.042l.014-.031a.443.443 0 0 0 .017-.047l.01-.03a.628.628 0 0 0 .01-.054l.006-.028a.66.66 0 0 0 .006-.083V9.739a.648.648 0 0 0-.006-.083Z"/>`
	coffeeInnerSVG                 = `<g fill="none"><path stroke="currentColor" stroke-width="2" d="M6.264 20.192c4.096 2.868 8.602-.081 11.47-4.177c2.868-4.095 4.097-9.338.002-12.206c-4.096-2.868-8.602.08-11.47 4.176c-2.868 4.096-4.098 9.339-.002 12.207Z"/><path fill="currentColor" d="M16.589 5.447c-1.393.246-5.326 2.375-5.408 5.98c-.083 3.604-2.787 6.594-3.77 7.126c1.803.042 5.326-2.375 5.408-5.98c.083-3.604 2.786-6.594 3.77-7.126Z"/></g>`
	columnsInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17 2v20V2Zm-5 0v20V2ZM7 2v20V2ZM2 22h20V2H2v20Z"/>`
	commandInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 9H5.5C3.5 9 2 7.5 2 5.5S3.5 2 5.5 2S9 3.5 9 5.5v13c0 2-1.5 3.5-3.5 3.5S2 20.5 2 18.5S3.5 15 5.5 15h13c2 0 3.5 1.5 3.5 3.5S20.5 22 18.5 22S15 20.5 15 18.5v-13c0-2 1.5-3.5 3.5-3.5S22 3.5 22 5.5S20.5 9 18.5 9H12Z"/>`
	compareInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M11 7H1v10h6V8m4-3v4l2-2l-2-2Zm0 12l2 2v-4l-2 2Zm2 0h10V7h-6v9"/>`
	compassInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm-2-12l5-1l-1 5l-5 1l1-5Zm2 4a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`
	complianceInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 6h8V1H8v5Zm8-3h5v20H3V3h5m0 11l3 3l6-6"/>`
	configureInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 15c4.009-.065 7-3.033 7-7c0-3.012-.997-2.015-2-1c-.991.98-3 3-3 3l-4-4s2.02-2.009 3-3c1.015-1.003 1.015-2-1-2c-3.967 0-6.947 2.991-7 7c.042.976 0 3 0 3c-1.885 1.897-4.34 4.353-6 6c-2.932 2.944 1.056 6.932 4 4c1.65-1.662 4.113-4.125 6-6c0 0 2.024-.042 3 0Z"/>`
	connectInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M10 21c-2.5 2.5-5 2-7 0s-2.5-4.5 0-7l3-3l7 7l-3 3Zm4-18c2.5-2.5 5-2 7.001 0c2.001 2 2.499 4.5 0 7l-3 3L11 6l3-3Zm-3 7l-2.5 2.5L11 10Zm3 3l-2.5 2.5L14 13Z"/>`
	connectivityInnerSVG           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8v4l3 3l6-6l-4-4H5L2.5 3M17 7l3 3v7m-7-6l3 3m-5-1l3 3m-7 1l3 3h10"/>`
	consoleInnerSVG                = `<path fill="currentColor" d="M16 18a1 1 0 1 0 0-2v2zm-8-2a1 1 0 1 0 0 2v-2zm.707-8.707a1 1 0 1 0-1.414 1.414l1.414-1.414zM11 11l.707.707a1 1 0 0 0 0-1.414L11 11zm-3.707 2.293a1 1 0 1 0 1.414 1.414l-1.414-1.414zM7 4h10V2H7v2zm13 3v10h2V7h-2zm-3 13H7v2h10v-2zM4 17V7H2v10h2zm3 3a3 3 0 0 1-3-3H2a5 5 0 0 0 5 5v-2zm13-3a3 3 0 0 1-3 3v2a5 5 0 0 0 5-5h-2zM17 4a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5v2zM7 2a5 5 0 0 0-5 5h2a3 3 0 0 1 3-3V2zm9 14H8v2h8v-2zM7.293 8.707l3 3l1.414-1.414l-3-3l-1.414 1.414zm3 1.586l-3 3l1.414 1.414l3-3l-1.414-1.414z"/>`
	contactInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 2h21v16h-8l-8 4v-4H1V2Zm5 8h1v1H6v-1Zm5 0h1v1h-1v-1Zm5 0h1v1h-1v-1Z"/>`
	contactInfoInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 12a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm4 6v-2c0-2.25-1.787-4-4.036-4h.054C2.768 12 1 13.75 1 16v2M12 7h12M12 17h10m-10-5h7"/>`
	contractInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 14h8v8m-9 1l9-9M23 1l-9 9m8 0h-8V2"/>`
	copyInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 15h8h-8Zm0-4h10H9Zm0-4h4h-4Zm7-6v6h6M6 5H2v18h16v-4m4 0H6V1h11l5 5v13Z"/>`
	cpuInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 18h3m-3-4h3m-3-4h3M1 6h3m16 12h3m-3-4h3m-3-4h3m-3-4h3M6 1v3m4-3v3m4-3v3m4-3v3M6 20v3m4-3v3m4-3v3m4-3v3M5 20h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1zm8-13h4v4h-4V7z"/>`
	creativeCommonsInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M11.984 0c-3.292 0-6.19 1.218-8.428 3.49C1.25 5.827 0 8.856 0 12.016c0 3.194 1.218 6.157 3.523 8.461c2.304 2.305 5.3 3.556 8.46 3.556c3.16 0 6.223-1.251 8.593-3.588C22.815 18.239 24 15.309 24 12.016c0-3.259-1.185-6.222-3.457-8.493C18.24 1.218 15.276 0 11.983 0Zm.032 2.173c2.7 0 5.104 1.02 6.98 2.897c1.843 1.843 2.83 4.28 2.83 6.946c0 2.7-.954 5.07-2.797 6.881c-1.943 1.91-4.445 2.93-7.013 2.93c-2.6 0-5.037-1.02-6.913-2.897c-1.877-1.877-2.93-4.346-2.93-6.914c0-2.6 1.053-5.07 2.93-6.98c1.843-1.875 4.214-2.863 6.913-2.863Zm-.154 7.846c-.68-1.237-1.837-1.73-3.181-1.73c-1.957 0-3.514 1.384-3.514 3.727c0 2.382 1.464 3.727 3.58 3.727c1.358 0 2.516-.745 3.154-1.877l-1.49-.758c-.333.798-.839 1.038-1.478 1.038c-1.105 0-1.61-.919-1.61-2.13c0-1.21.426-2.13 1.61-2.13c.32 0 .959.173 1.331.972l1.598-.839Zm6.932 0c-.68-1.237-1.837-1.73-3.181-1.73c-1.957 0-3.514 1.384-3.514 3.727c0 2.382 1.464 3.727 3.58 3.727c1.358 0 2.516-.745 3.154-1.877l-1.49-.758c-.333.798-.839 1.038-1.477 1.038c-1.105 0-1.611-.919-1.611-2.13c0-1.21.426-2.13 1.61-2.13c.32 0 .959.173 1.331.972l1.598-.839Z"/>`
	creditCardInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 5c0-.552.44-1 1.002-1h19.996A1 1 0 0 1 23 5v14c0 .552-.44 1-1.002 1H2.002A1 1 0 0 1 1 19V5Zm0 3h22v2H1V8Zm4 7h2v.5H5V15Zm5 0h6v.5h-6V15Z"/>`
	cssThreeInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 3h14l-3 15l-8 4l-7-4l1-4m1-5h14"/>`
	cubeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="m12 2l10 5v10l-10 5l-10-5V7l10-5ZM2 7l10 5l10-5m-10 5v10v-10Z"/>`
	cubesInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="m6.5 10.5l5.5 3l5.5-3v-6l-5.5-3l-5.5 3v6Zm0-6l5.5 3l5.5-3m-5.5 3v6v-6Zm-11 12l5.5 3l5.5-3v-6l-5.5-3l-5.5 3v6Zm0-6l5.5 3l5.5-3m-5.5 3v6v-6Zm5.5 3l5.5 3l5.5-3v-6l-5.5-3l-5.5 3v6Zm0-6l5.5 3l5.5-3m-5.5 3v6v-6Z"/>`
	currencyInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 5h22v14H1V5Zm1 4a3 3 0 0 0 3-3m-3 9a3 3 0 0 1 3 3m17-9a3 3 0 0 1-3-3m3 9a3 3 0 0 0-3 3m-7-2c1.657 0 3-1.79 3-4s-1.343-4-3-4s-3 1.79-3 4s1.343 4 3 4Z"/>`
	cursorInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m6 3l12 11l-5 1l3 5.5l-3 1.5l-3-6l-4 3z"/>`
	cutInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M23 4L8 16L23 4Zm0 16L8 8l15 12ZM5 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`
	cycleInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M13 20c6-1 8-6 8-10m-7 6l-2 4l4 3M0 9l4-3l3 4m2 10c-6-3-7-8-5-14m16 1C16 1 10 1 6 4.006M20 2v5h-5"/>`
	dashboardInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm3-6a3 3 0 0 0-6 0M5 5l2 2m5 0v6m0-10v2m7 7h2M3 12h2m12-5l2-2M3 17h18"/>`
	databaseInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 2h22v7H1V2Zm3 10h1v1H4v-1Zm0-7h1v1H4V5Zm0 14h1v1H4v-1Zm-3-3h22v7H1v-7Zm0-7h22v7H1V9Z"/>`
	debianInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M2.656 5.167c.03.167.383-.233.107.39c-.532.371-.064.162-.107-.39Zm.003.032l-.005-.06c0 .023.001.043.005.06Zm-.446 1.76c-.344.44-.159.535-.195.834c.125-.382.147-.612.195-.833ZM13.04 0c-.372.031-.742.05-1.108.097l.161.022c.27-.099.661-.054.947-.119Zm-.611.467l.01-.04l-.15.066l.14-.026ZM12.015.31c.212.038.459.067.424.118c.232-.051.285-.098-.424-.118ZM9.8 13.202c-.138-.153-.22-.338-.312-.521c.088.323.267.6.434.881l-.122-.36Zm4.317-.237c.144-.112.274-.226.39-.336a2.97 2.97 0 0 1-.984.05c-.398.007.075.206.594.286Zm1.54-.817c.237-.327.41-.685.47-1.055c-.053.264-.196.492-.33.732c-.744.468-.07-.278-.001-.561c-.8 1.006-.11.603-.139.884Zm.788-2.05c.048-.716-.141-.49-.205-.217c.075.039.133.505.205.217Zm-4.558 5.147a4.674 4.674 0 0 1-1.527-1.168c.228.335.475.66.794.916c-.54-.183-1.26-1.308-1.47-1.354c.93 1.666 3.773 2.92 5.262 2.298c-.69.026-1.565.014-2.339-.272c-.294-.151-.685-.45-.697-.553c-.015.038-.038.076-.023.133Zm.04-.157a.057.057 0 0 0-.016.05a.367.367 0 0 0 .016-.05Zm-2.8-1.583c.26.355.468.74.802 1.017c-.24-.469-.419-.663-.747-1.296l-.056.279ZM20.46 8.576c.006-.365.101-.191.138-.281c-.071-.041-.26-.321-.374-.858c.083-.126.222.327.335.345c-.073-.426-.198-.752-.203-1.08c-.33-.689-.117.092-.384-.295c-.351-1.095.29-.254.334-.752c.532.771.836 1.965.975 2.46a10.077 10.077 0 0 0-.488-1.753c.162.068-.26-1.241.21-.374c-.502-1.848-2.15-3.575-3.665-4.386c.185.17.42.383.335.416c-.753-.448-.621-.483-.729-.673c-.614-.25-.654.02-1.06 0c-1.159-.613-1.382-.548-2.447-.933l.049.227c-.767-.256-.894.097-1.722 0c-.05-.039.265-.142.525-.18c-.741.098-.706-.146-1.432.027c.179-.125.368-.208.559-.315c-.605.037-1.444.352-1.185.065c-.986.44-2.737 1.058-3.72 1.98l-.03-.207c-.451.54-1.965 1.615-2.085 2.315l-.12.028c-.235.397-.387.846-.573 1.255c-.306.522-.449.2-.405.282c-.603 1.223-.902 2.25-1.161 3.092c.184.276.004 1.659.074 2.766c-.303 5.467 3.837 10.775 8.362 12c.663.238 1.65.229 2.488.253c-.99-.283-1.118-.15-2.082-.486c-.695-.328-.848-.702-1.34-1.13l.195.345c-.966-.342-.562-.423-1.348-.672l.208-.272c-.313-.023-.83-.528-.97-.807l-.343.014c-.412-.508-.631-.874-.615-1.158l-.11.197c-.126-.215-1.515-1.904-.795-1.511c-.134-.122-.312-.2-.505-.55l.147-.167c-.346-.447-.638-1.018-.616-1.209c.185.25.313.297.44.34c-.875-2.173-.924-.12-1.588-2.212l.14-.012c-.107-.162-.172-.337-.259-.51l.061-.609c-.63-.729-.176-3.1-.085-4.4c.063-.528.526-1.09.879-1.973L4.26 7.18c.41-.716 2.343-2.875 3.238-2.764c.433-.544-.086-.002-.171-.139c.952-.985 1.251-.696 1.894-.873c.693-.412-.595.16-.266-.157c1.198-.306.849-.696 2.412-.851c.165.094-.383.145-.52.266c.998-.488 3.16-.377 4.562.271c1.63.762 3.459 3.011 3.53 5.128l.083.022c-.042.842.129 1.815-.166 2.709l.2-.424c.024.644-.188.955-.378 1.508l-.344.172c-.281.546.027.346-.174.78c-.44.391-1.333 1.222-1.619 1.298c-.209-.004.142-.246.188-.34c-.588.403-.472.605-1.371.85l-.026-.058c-2.218 1.043-5.297-1.024-5.256-3.844c-.024.179-.068.134-.117.206c-.114-1.45.67-2.908 1.993-3.503c1.294-.64 2.812-.378 3.739.486c-.51-.667-1.523-1.374-2.724-1.308c-1.177.019-2.278.767-2.645 1.578c-.603.38-.673 1.463-.935 1.662c-.354 2.597.664 3.72 2.387 5.04c.137.092.155.145.144.192l.001-.002c2.032.76 4.131.575 5.89-.834c.447-.348.935-.94 1.077-.95c-.213.32.036.155-.128.437c.446-.719-.193-.293.461-1.241l.242.332c-.09-.596.741-1.32.657-2.264c.19-.289.212.31.01.974c.28-.736.074-.855.146-1.462c.078.204.18.42.233.636c-.183-.711.187-1.198.279-1.612c-.09-.04-.282.315-.326-.526Z"/>`
	deliverInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 18H1V3h13v14m0 1H9m-3 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm11 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM14 8h5l4 5v5h-3"/>`
	deployInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M23 1s-6.528-.458-9 2c-.023.037-4 4-4 4L5 8l-3 2l8 4l4 8l2-3l1-5s3.963-3.977 4-4c2.458-2.472 2-9 2-9Zm-6 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM7 17c-1-1-3-1-4 0s-1 5-1 5s4 0 5-1s1-3 0-4Z"/>`
	descendInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="m2 16l6 6l6-6M11 3h11M11 7h8m-8 4h5M8 22V2"/>`
	descendingInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="m11.92 16.714l-.354.353l-.354-.353L7 12.502l.707-.708l3.359 3.359V7h1v8.153l3.358-3.359l.707.708l-4.212 4.212Z" clip-rule="evenodd"/>`
	desktopInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h22v18H1V1Zm4 22h14H5Zm5-4v4v-4Zm4 0v4v-4Z"/>`
	detachInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m4 4l16 16m2-8l-5.28 5.28M15 19l-2 2c-6 6-15-3-9-9l2-2m2-2l5-5c4-4 10 2 6 6l-5 5m-2 2l-2 2c-2 2-5-1-3-3l2-2m2-2l5-5"/>`
	deviceInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 7h20v14h-6l-4-4l-4 4H2V7Zm4-5l5 5h2l5-5"/>`
	diamondInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 3h12l4 6l-10 12L2 9l4-6ZM2 9h20M11 3L7 9l5 11m1-17l4 6l-5 11"/>`
	directionsInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="m17 11l5-5l-5-5m5 5h-4a6 6 0 0 0-6 6v12M7 6l-5 5l5 5m-5-5h4a6 6 0 0 1 6 6v7"/>`
	disabledOutlineInnerSVG        = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 12H6M4 22h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2z"/>`
	discInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm0-7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`
	dislikeInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M23 1H4C2 1 1 2 1 4v10h7v6c0 2 1 3 3 3h2s.016-6 .016-7.326C13.016 14.348 14 13 16 13h7V1Zm-5 0v12"/>`
	dockerInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M6.942 14.9c.056 0 .11.01.158.03a.179.179 0 1 0 .246.24a.438.438 0 1 1-.404-.27Zm0 1.185a.749.749 0 1 1 .002-1.497a.749.749 0 0 1-.002 1.497Zm13.444-4.901c-2.124 5.628-6.92 8.135-12.576 8.135c-2.672 0-4.803-.92-6.167-2.452l.01-.006c.393.02.745.026 1.101.026c.327 0 .646-.003.941-.02l.084-.006s.091-.006.046-.007a8.545 8.545 0 0 0 1.877-.306a4.82 4.82 0 0 0 .286-.09a.197.197 0 0 0-.128-.371c-.69.239-1.6.37-2.715.395a21.668 21.668 0 0 1-1.86-.045a6.3 6.3 0 0 1-.386-.58l-.187-.34C.15 14.411-.096 13.12.034 11.716h16.363c1.344 0 2.656-.502 3.28-1.055c-1.117-.908-1.006-3.064-.295-3.886c.618.496 1.613 1.54 1.442 2.871c.777-.39 2.127-.583 3.176.022c-.659 1.286-2.107 1.67-3.614 1.516Zm-18.13.135h2.212V9.106H2.255v2.213Zm2.552 0h2.213V9.106H4.808v2.213Zm0-2.553h2.213V6.553H4.808v2.213Zm2.553 2.553h2.213V9.106H7.361v2.213Zm0-2.553h2.213V6.553H7.361v2.213Zm2.553 2.553h2.213V9.106H9.914v2.213Zm0-2.553h2.213V6.553H9.914v2.213Zm0-2.553h2.213V4H9.914v2.213Zm2.553 5.106h2.213V9.106h-2.213v2.213Z"/>`
	documentInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M14 1v7h7m0 15H3V1h12l3 3l3 3v16Z"/>`
	documentCloudInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-4M18 1v5h5m-12 7H6.002A3.003 3.003 0 0 0 3 16c0 1.657 1.343 3 2.99 3H7v1.01A2.993 2.993 0 0 0 10.002 23h1.996A2.999 2.999 0 0 0 15 20.01V19m-4 0h4.998A3.003 3.003 0 0 0 19 16c0-1.657-1.343-3-2.99-3H15v-1.01A2.993 2.993 0 0 0 11.998 9h-1.996A2.999 2.999 0 0 0 7 11.99V13"/>`
	documentConfigInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23h-7m2-22v5h5M9 14v-3m0 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 3v-3m3-3h3M3 17h3m-1-4l2 2m4 4l2 2m0-8l-2 2m-4 4l-2 2"/>`
	documentCsvInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M7 13H5c-1 0-2 .5-2 1.5v3c0 1 1 1.5 2 1.5h2m6.25-6h-2.5c-1.5 0-2 .5-2 1.5s.5 1.5 2 1.5s2 .5 2 1.5s-.5 1.5-2 1.5h-2.5m12.25-7v.5C20.5 13 18 19 18 19h-.5S15 13 15 12.5V12"/>`
	documentDownloadInnerSVG       = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 1H17.5L21 4.5V23H3L2.998 1ZM16 1v5h5m-9 3v9m-4-3l4 4l4-4"/>`
	documentExcelInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M9.25 12l-2 3.25l-2-3.25H5l2.25 3.5l-2.5 3.5H5l2.25-3.25L9.5 19h.25l-2.5-3.5L9.5 12h-.25Z"/>`
	documentImageInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-3M18 1v5h5M3 11h13v12H3V11Zm4 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-2 7l2-3l2 2l4-6l3 4"/>`
	documentLockedInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-6m1-22v5h5M4 15h10v8H4v-8Zm2 0v-2a3 3 0 0 1 6 0v2m-4 4h2"/>`
	documentMissingInnerSVG        = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 1H17.5L21 4.5V23H3L2.998 1ZM16 1v5h5M9 12l6 6m0-6l-6 6"/>`
	documentNotesInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M9.75 12v6.5H9.5l-5-6.5H4v7h.25v-6.5h.25l5 6.5h.5v-7h-.25Z"/>`
	documentOutlookInnerSVG        = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 9V1H17.5L21 4.5V23H2M16 1v5h5M7.75 15.75C7.75 13.5 6.5 12 5 12s-2.75 1.5-2.75 3.75S3.5 19.5 5 19.5s2.75-1.5 2.75-3.75ZM5 12c2.425 0 3 2.5 3 3.75s-.5 3.75-3 3.75s-3-2.5-3-3.75S2.559 12 5 12Z"/>`
	documentPdfInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M3 12h1.5c2 0 2.25 1.25 2.25 2s-.25 2-2.25 2H3.25v2H3v-6Zm6.5 6v-6h1.705c1.137 0 2.295.5 2.295 3s-1.158 3-2.295 3H9.5Zm7 1v-7h4m-4 3.5h3"/>`
	documentPerformanceInnerSVG    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M3 19l5-5l4 4l6.5-6.5M19 17v-6h-6"/>`
	documentPptInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M4 12h1.5c2 0 3.5.5 3.5 2.25S7.5 16.5 5.5 16.5H4.25V19H4v-7Z"/>`
	documentRtfInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5m-2.5 6h-4v7m3-3.5h-3m-8-3.5h6m-3 0v7M3 19v-7h1.5C5 12 7 12 7 14s-2 2-2.5 2H3m2.25 0l2.25 3"/>`
	documentSoundInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 9V1H17.5L21 4.5V23h-3M16 1v5h5M1 14.01v4h3l4 3V11l-4 3.01H1ZM11 18a2 2 0 1 0 0-4m0 8a6 6 0 1 0 0-12"/>`
	documentStoreInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-6m1-22v5h5M3 12s1-2 6-2s6 2 6 2v9s-1 2-6 2s-6-2-6-2v-9Zm0 5s2 2 6 2s6-2 6-2M3 13s2 2 6 2s6-2 6-2"/>`
	documentTestInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 6V1H19.5L23 4.5V23h-3M18 1v5h5M6 9h8M8 9v4.5l-5 8V23h14v-1.581L12 13.5V9m-6.5 8.5s2 1.5 4.5 0s4.5 0 4.5 0"/>`
	documentTextInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 16h10H6Zm0-4h12H6Zm0-4h5h-5Zm8-7v7h7M3 23V1h12l6 6v16H3Z"/>`
	documentThreatInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-6m1-22v5h5M9 23a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0-12V9c0-1 0-2 2-2s2 1 2 2s0 2 2 2h2m-9 0h2v2H8v-2Z"/>`
	documentTimeInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-6m1-22v5h5M10 23a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0-11v4l3 3"/>`
	documentTransferInnerSVG       = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M8 12l-4 4l4 4m-4-4h11"/>`
	documentTxtInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M2 12h5m-2.5 0v7M16 12h5m-2.5 0v7m-4-7.5l-6 7m0-7l6 7"/>`
	documentUpdateInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 7V1H17.5L21 4.5V23h-6m1-22v5h5M8 23A7 7 0 1 0 8 9a7 7 0 0 0 0 14Zm-3.5-6.5L8 13l3.5 3.5m-3.5-3V20"/>`
	documentUploadInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 1H17.5L21 4.5V23H3L2.998 1ZM16 1v5h5m-9 14v-9m-4 3l4-4l4 4"/>`
	documentUserInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23h-7m2-22v5h5M8 11a3 3 0 1 0 0 6a3 3 0 0 0 0-6ZM3 23v-1c0-4 3-5 5-5s5 1 5 5v1H3Z"/>`
	documentVerifiedInnerSVG       = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 1H17.5L21 4.5V23H3L2.998 1ZM16 1v5h5M7.5 15l3 3l6-6"/>`
	documentVideoInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23H4M18 1v5h5M3 10h9v9H3v-9Zm9 3l5-2.5v8L12 16v-3Z"/>`
	documentWindowsInnerSVG        = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5m-8.75 5.5l-6 7m0-7l6 7M20.5 12h-4v6h4m-1-3h-3M7 12H3v6h4m-1-3H3"/>`
	documentWordInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5m-9 6l-1.5 6.75h-.25L9.5 12H9l-2.75 6.75H6L4.5 12H4l2 7h.5L9 12.5h.5L12 19h.5l2-7H14Z"/>`
	documentZipInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5M2 13h5v1l-4 4v1h5m3-7v8v-8Zm4 1v7v-7Zm5 2a2 2 0 0 0-2-2h-3v4h3a2 2 0 0 0 2-2Z"/>`
	domainInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M13 3v4v-4ZM9 3v4v-4ZM5 3v4v-4ZM1 7h22H1Zm0 14h22V3H1v18Z"/>`
	dosInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="m0 4.546l4.934-.124c.958-.02 1.844.082 2.74.392v.175C6.9 5.648 6.345 6.43 5.84 7.306l-2.997.021l.02 8.498l1.772.123l1.813-.02c1.082-.011 2.08-1.741 2.174-2.679l.267-2.955c.114-1.298.865-2.297 2.03-2.833c.73 1.226 1.05 2.585 1.05 4.017c0 4.501-2.791 7.344-7.19 7.18L0 18.471V4.546Zm20.704 4.172c-.237-1.37-1.38-2.05-2.75-2.05c-.938 0-2.473.411-2.524 1.586c-.01.062 0 .113.02.165l.557 1.596c.113.33.134.69.134 1.04c0 .402-.052.804-.103 1.206c-1.978-.556-3.369-1.514-3.369-3.74c0-2.997 2.421-4.398 5.181-4.398c3.07 0 5.47 1.432 5.666 4.595h-2.812Zm-8.24 5.366l2.75.02c.288 1.762 1.452 2.246 3.182 2.246c1.082 0 2.699-.257 2.699-1.638c0-.72-.567-1.092-1.143-1.41l.144-2.75c2.194.596 3.904 1.42 3.904 3.995c0 3.05-2.905 4.337-5.583 4.337c-1.916 0-4.398-.556-5.418-2.369c-.32-.556-.618-1.277-.618-1.926a.78.78 0 0 1 .02-.155l.062-.35Zm7.24-2.627c0 1.576-.422 3.049-1.143 4.43c-.185.01-.37.03-.556.03c-.597 0-2.05-.401-2.05-1.133c0-.123.02-.247.062-.36l.546-1.823c.113-.36.072-.794.072-1.154c0-1.411-.639-2.297-.69-3.265c-.03-.68 1.257-1 1.71-1c.309 0 .618 0 .927.031a8.49 8.49 0 0 1 1.123 4.244ZM10.61 17.03l.648-.927l1.082.02c.515 1.123 1.36 1.885 2.39 2.545a7.166 7.166 0 0 1-2.277.37c-1.143 0-2.215-.257-3.275-.659l1.432-1.35Zm1.926-10.321a8.19 8.19 0 0 0-.495-.02c-1.74 0-3.327 1.523-3.533 3.233l-.453 3.719c-.062.525-.515 1.06-.855 1.442c-.288.309-.628.37-1.05.37h-.052a9.589 9.589 0 0 1-.814-3.862C5.284 7.358 7.654 4 12.062 4c.865 0 1.689.165 2.502.391c-.864.639-1.544 1.35-2.028 2.318Z"/>`
	downInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="m2 8.35l10.173 9.823L21.997 8"/>`
	downloadInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 17v6h22v-6M12 2v17m-7-7l7 7l7-7"/>`
	downloadOptionInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12ZM2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Zm14.293-.707L13 14.586V6h-2v8.586l-3.293-3.293l-1.414 1.414l5 5l.707.707l.707-.707l5-5l-1.414-1.414Z" clip-rule="evenodd"/>`
	dragInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M15 5h2V3h-2v2ZM7 5h2V3H7v2Zm8 8h2v-2h-2v2Zm-8 0h2v-2H7v2Zm8 8h2v-2h-2v2Zm-8 0h2v-2H7v2Z"/>`
	drawerInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 14L6 2h12l5 12l-2 8H3l-2-8Zm0 0h22"/>`
	dribbbleInnerSVG               = `<path fill="currentColor" d="M11.432 8.635c-1.77-3.15-3.666-5.716-3.803-5.904A10.263 10.263 0 0 0 1.97 9.887c.27.004 4.54.056 9.46-1.252Zm1.278 3.443c.135-.041.27-.083.404-.122a34.204 34.204 0 0 0-.832-1.741c-5.278 1.58-10.342 1.464-10.521 1.46c-.003.107-.008.215-.008.325a10.2 10.2 0 0 0 2.63 6.852l-.007-.01s2.804-4.976 8.334-6.764ZM5.701 20.08l.003-.005c-.076-.058-.157-.115-.233-.176c.137.11.23.181.23.181ZM9.62 2.076c-.036.01-.063.02-.1.03a.535.535 0 0 1 .1-.03Zm9.15 2.234A10.198 10.198 0 0 0 12 1.751c-.833 0-1.64.103-2.415.289c.157.206 2.08 2.762 3.83 5.978c3.865-1.447 5.327-3.666 5.354-3.708ZM12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12Zm1.744-10.322c-6.015 2.096-8.001 6.31-8.04 6.396A10.2 10.2 0 0 0 12 22.247c1.42 0 2.772-.29 4.002-.811c-.152-.899-.747-4.038-2.19-7.783c-.024.01-.046.015-.068.025Zm.46-4.132a30.12 30.12 0 0 1 .901 2.016c3.54-.446 7.024.31 7.14.335a10.206 10.206 0 0 0-2.332-6.406c-.02.029-1.663 2.405-5.709 4.055Zm1.528 3.634c1.347 3.698 1.89 6.708 1.994 7.32a10.242 10.242 0 0 0 4.39-6.874c-.203-.066-3.07-.977-6.384-.446Z"/>`
	driveCageInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 23V2h22v21M1 8h22H1Zm0 6h22H1Zm0 6h22H1ZM4 5h12H4Zm14 0h2h-2Zm0 6h2h-2Zm0 6h2h-2ZM4 11h12H4Zm0 6h12H4Z"/>`
	dropboxInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M7.06 1L0 5.61l4.882 3.908L12 5.123L7.06 1ZM0 13.428l7.06 4.61L12 13.914L4.882 9.52L0 13.43Zm12 .486l4.94 4.124l7.06-4.61l-4.882-3.91L12 13.914ZM24 5.61L16.94 1L12 5.124l7.118 4.395L24 5.609ZM12.014 14.8L7.06 18.913l-2.12-1.385v1.552l7.074 4.243l7.075-4.243v-1.552l-2.12 1.385l-4.955-4.112Z"/>`
	duplicateInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.5 17H1V1h16v3.5M7 7h16v16H7V7Zm8 4v8v-8Zm-4 4h8h-8Z"/>`
	dxcInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="m12 14l4 7H8l4-7Zm0-4L8 3h8l-4 7ZM2 18H0V6h2a6 6 0 1 1 0 12Zm20 0a6 6 0 1 1 0-12h2v12h-2Z"/>`
	edgeInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M5.487 6.357C2.632 8.134 1 10.66 1 10.66s.423-5.296 4.487-8.395C7.103 1.033 9.313 0 12.285 0c1.117 0 3.459.195 5.568 1.495c2.109 1.3 2.961 2.39 3.911 3.984c.41.688.744 1.572.952 2.425c.391 1.598.438 3.508.438 3.508v2.517H8.081s-.368 5.065 6.564 5.065c2.411 0 3.255-.38 4.048-.614c1.241-.367 2.44-1.187 2.44-1.187l.002 5.06S18.298 24 14.012 24c-1.207 0-2.479-.101-3.706-.5c-1.072-.347-3.316-1.285-4.819-3.483c-.531-.778-1.107-1.813-1.393-2.824c-.308-1.093-.304-2.155-.304-2.74c0-2.188.747-4.277 2.044-5.787C7.515 6.71 9.638 5.848 9.638 5.848s-.691.807-1.117 1.81a8.288 8.288 0 0 0-.542 2.015h8.511s.498-5.086-4.815-5.086c-2.003 0-4.462.695-6.188 1.77"/>`
	editInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="m14 4l6 6l-6-6Zm8.294 1.294c.39.39.387 1.025-.008 1.42L9 20l-7 2l2-7L17.286 1.714a1 1 0 0 1 1.42-.008l3.588 3.588ZM3 19l2 2m2-4l8-8"/>`
	ejectInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M21 14L12 2L3 14h18ZM2 22h20v-4H2v4Z"/>`
	elevatorInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M1 2.991C1 1.891 1.89 1 2.991 1H21.01C22.109 1 23 1.89 23 2.991V21.01c0 1.1-.89 1.991-1.991 1.991H2.99A1.99 1.99 0 0 1 1 21.009V2.99ZM16.5 8l1.5 2h-3l1.5-2Zm0 8l1.5-2h-3l1.5 2ZM5 13l1.556-3.112C6.801 9.398 7.444 9 8 9h0c.552 0 1.2.398 1.444.888L11 13m-4.5 5l1.25-8h.5l1.25 8M8 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	emergencyInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M10 7.172V2h4v5.172l3.657-3.657l2.828 2.828L16.828 10H22v4h-5.172l3.657 3.657l-2.828 2.828L14 16.828V22h-4v-5.172l-3.657 3.657l-2.828-2.828L7.172 14H2v-4h5.172L3.515 6.343l2.828-2.828L10 7.172Z"/>`
	emojiInnerSVG                  = `<path fill="currentColor" d="M12 1.73A10.27 10.27 0 1 0 22.24 12A10.25 10.25 0 0 0 12 1.73ZM21 12a9 9 0 1 1-9-9a9 9 0 0 1 9 9Z"/><path fill="currentColor" d="M8.8 11.05a1.55 1.55 0 1 0-1.51-1.5a1.56 1.56 0 0 0 1.51 1.5Zm6.64 0a1.55 1.55 0 1 0 0-3.09a1.53 1.53 0 0 0-1.51 1.59a1.51 1.51 0 0 0 1.51 1.5Zm-3.25 5.3A6.58 6.58 0 0 1 6.9 13.5a5.71 5.71 0 0 0 5.3 4a5.54 5.54 0 0 0 5.31-4a6.27 6.27 0 0 1-5.32 2.85Z"/>`
	emptyCircleInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0-6a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`
	eraseInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 21L22 6l-4-4L2 18l3 3h14M6 14l4 4"/>`
	escalatorInnerSVG              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M22 9a2 2 0 0 0-1.998-2H16L6 17H4c-1.105 0-2 .888-2 2h0a2 2 0 0 0 1.998 2H8l10-10h2c1.105 0 2-.888 2-2h0ZM7 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 6v-2.505c0-.273.232-.495.5-.495h0c.276 0 .5.214.5.505V14l-1 1Zm5-11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 6V7.495c0-.273.232-.495.5-.495h0c.276 0 .5.214.5.505V9l-1 1Z"/>`
	expandInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m10 14l-8 8m-1-7v8h8M22 2l-8 8m1-9h8v8"/>`
	ezmeralInnerSVG                = `<path fill="currentColor" d="M7 8h34v8H7V8ZM1 8h6v8H1V8Zm40 0h6v8h-6V8ZM7 16h34v6H7v-6ZM7 2h34v6H7V2ZM1 8l6-6v6H1Zm0 8l6 6v-6H1Zm46-8l-6-6v6h6Zm0 8l-6 6v-6h6Z"/>`
	facebookInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M23 0H1a1 1 0 0 0-1 1v22a1 1 0 0 0 1 1h11.75v-9h-3v-3.75h3v-3c0-3.1 1.963-4.625 4.728-4.625c1.324 0 2.463.099 2.794.142v3.24l-1.917.001c-1.504 0-1.855.715-1.855 1.763v2.479h3.75L19.5 15h-3l.06 9H23a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1"/>`
	facebookOptionInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M9.945 22v-8.834H7V9.485h2.945V6.54c0-3.043 1.926-4.54 4.64-4.54c1.3 0 2.418.097 2.744.14v3.18h-1.883c-1.476 0-1.82.703-1.82 1.732v2.433h3.68l-.736 3.68h-2.944L13.685 22"/>`
	fanInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm3-3c4 3 5 7 5 7m-8 4c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm0-14c3-5 7-6 7-6m-7 12c-3 4-7 5-7 5m4-8C5 9 4 5 4 5"/>`
	fanOptionInnerSVG              = `<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="1"/><rect width="20" height="20" x="2" y="2" rx="10"/><path d="M15 9.5c.5-.333.9-1.7.5-2.5S13.333 5.667 13 5.5m1 5c1.5-2 0-3.5-2.5-5c-1.546-.927 2-1.5 4.5.5c1.875 1.5 1 2.5-2 5.5v-1Zm-5.015 3.902c-.5.333-.9 1.7-.5 2.5s2.167 1.333 2.5 1.5m-1-5c-1.5 2 0 3.5 2.5 5c1.546.927-2 1.5-4.5-.5c-1.875-1.5-1-2.5 2-5.5v1Zm-.443-4.458c-.334-.5-1.7-.9-2.5-.5s-1.334 2.166-1.5 2.5m5-1c-2-1.5-3.5 0-5 2.5c-.928 1.546-1.5-2 .5-4.5c1.5-1.875 2.5-1 5.5 2h-1Zm3.902 5.014c.333.5 1.7.9 2.5.5s1.333-2.166 1.5-2.5m-5 1c2 1.5 3.5 0 5-2.5c.927-1.546 1.5 2-.5 4.5c-1.5 1.876-2.5 1-5.5-2h1Z"/><path d="M3.5 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 17a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm17 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-17a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g>`
	fastForwardInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 2.059V8L1 2.059v20L9 16v6.059l13-10z"/>`
	favoriteInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 8.4C1 4 4.5 3 6.5 3C9 3 11 5 12 6.5C13 5 15 3 17.5 3c2 0 5.5 1 5.5 5.4C23 15 12 21 12 21S1 15 1 8.4Z"/>`
	fedoraInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M12 0C5.375 0 .005 5.368 0 11.992v9.286A2.728 2.728 0 0 0 2.728 24h9.277C18.63 23.997 24 18.626 24 12c0-6.627-5.373-12-12-12Zm4.595 5.577c-.379 0-.517-.073-1.072-.073a2.973 2.973 0 0 0-2.973 2.968v2.583a.42.42 0 0 0 .42.419h1.953c.728 0 1.316.58 1.316 1.31c0 .734-.594 1.312-1.33 1.312H12.55v2.985a5.632 5.632 0 0 1-5.631 5.632c-.472 0-.808-.053-1.245-.167c-.637-.167-1.157-.689-1.157-1.296c0-.734.533-1.269 1.33-1.269c.378 0 .516.073 1.072.073c1.64 0 2.97-1.328 2.972-2.968v-2.583a.42.42 0 0 0-.42-.419H7.518c-.727 0-1.315-.58-1.315-1.31c0-.735.594-1.312 1.33-1.312H9.89V8.476a5.632 5.632 0 0 1 5.632-5.632c.472 0 .807.054 1.244.168c.637.167 1.158.689 1.158 1.296c0 .734-.533 1.269-1.33 1.269Z"/>`
	figmaInnerSVG                  = `<g fill="currentColor" fill-rule="evenodd" transform="translate(4)"><circle cx="12" cy="12" r="4"/><path d="M4 24a4 4 0 0 0 4-4v-4H4a4 4 0 1 0 0 8Zm0-8h4V8H4a4 4 0 1 0 0 8Zm0-8h4V0H4a4 4 0 1 0 0 8Zm8 0H8V0h4a4 4 0 1 1 0 8Z"/></g>`
	filterInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m3 6l7 7v8h4v-8l7-7V3H3z"/>`
	fingerPrintInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M14 15a2 2 0 1 0-2 2h0m0 3a5 5 0 1 1 5-5a1.5 1.5 0 0 0 3 0a8 8 0 1 0-8 8h2M1 15c0 2.672.953 5.122 2.537 7.027M20.52 8.042A10.978 10.978 0 0 0 12 4a10.977 10.977 0 0 0-8.464 3.974m14.99-5.363A13.939 13.939 0 0 0 12 1a13.94 13.94 0 0 0-6.481 1.587"/>`
	fireballInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm0-9l-1.75 1l.75-2l-1.25-1h1.5L12 9.25l.75 1.75h1.5L13 12l.75 2L12 13Z"/>`
	firefoxInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="m4.25 1l-.145 3.684c.624-.154 1.235-.246 1.83.01c1.291-1.768 3.148-2.376 5.194-2.6l.038.06c-.04.045-.08.089-.122.131a9.473 9.473 0 0 0-1.948 2.847c-.134.306-.166.658-.238.99c-.029.129.037.198.168.216c.607.083 1.212.18 1.822.25c.382.043.769.051 1.154.067c.18.008.258.097.262.267c.015.684-.213 1.27-.759 1.69a4.414 4.414 0 0 1-1.781.804c-.068.015-.135.035-.208.054l.198 2.515l-1.854-.888c-.186.392-.2.78-.082 1.176c.316 1.073 1.431 1.695 2.586 1.41c.404-.1.795-.27 1.176-.444c.376-.17.729-.39 1.096-.583c.622-.324 1.249-.296 1.874-.006c.09.041.18.095.249.165c.234.234.417.503.323.857c-.091.342-.34.54-.677.611a2.823 2.823 0 0 1-.594.059c-.132-.001-.202.042-.274.148c-.68 1.009-1.562 1.752-2.79 1.964c-.517.09-1.05.075-1.576.104c-.07.004-.14-.006-.255-.012c.07.064.105.104.147.135c1.39 1.027 2.927 1.378 4.614.974c1.381-.33 2.61-.949 3.623-1.965c.966-.968 1.408-2.15 1.422-3.496c.014-1.353-.288-2.634-.98-3.807c-.121-.205-.279-.39-.454-.63c1.19.54 2.24 1.153 2.727 2.42c.143-1.582-.148-3.08-.758-4.525c-.61-1.444-1.485-2.695-2.682-3.731c.05.01.1.018.149.034c2.694.844 4.763 2.45 6.056 4.991c.626 1.23.946 2.552 1.111 3.913c.177 1.457.156 2.906-.216 4.338c-.622 2.398-1.935 4.359-3.753 6.013c-1.73 1.574-3.745 2.595-6.052 2.887c-5.14.651-9.19-1.198-12.063-5.526C.665 16.849.126 14.92.023 12.859C-.144 9.51.61 6.385 2.272 3.478A9.851 9.851 0 0 1 4.049 1.18c.05-.048.102-.093.2-.181"/>`
	firewallInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M6.006 2.02a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6Zm5.916 2.976a1.1 1.1 0 0 0-1.1 1.1v1.8a1.1 1.1 0 0 0 1.1 1.1h5.8a1.1 1.1 0 0 0 1.1-1.1v-1.8a1.1 1.1 0 0 0-1.1-1.1h-5.8Zm-8.822 0a1.1 1.1 0 0 0-1.1 1.1v1.8a1.1 1.1 0 0 0 1.1 1.1h5.8a1.1 1.1 0 0 0 1.1-1.1v-1.8a1.1 1.1 0 0 0-1.1-1.1H3.1Zm0 9.992a1.1 1.1 0 0 0-1.1 1.1v1.8a1.1 1.1 0 0 0 1.1 1.1h5.8a1.1 1.1 0 0 0 1.1-1.1v-1.8a1.1 1.1 0 0 0-1.1-1.1H3.1Zm7.723 1.1a1.1 1.1 0 0 1 1.1-1.1h5.8a1.1 1.1 0 0 1 1.1 1.1v1.8a1.1 1.1 0 0 1-1.1 1.1h-5.8a1.1 1.1 0 0 1-1.1-1.1v-1.8Zm-5.975-5.015a1.1 1.1 0 0 1 1.1-1.1h5.8a1.1 1.1 0 0 1 1.1 1.1v1.8a1.1 1.1 0 0 1-1.1 1.1h-5.8a1.1 1.1 0 0 1-1.1-1.1v-1.8Zm9.866-1.061a1.1 1.1 0 0 0-1.1 1.1v1.8a1.1 1.1 0 0 0 1.1 1.1h5.8a1.1 1.1 0 0 0 1.1-1.1v-1.8a1.1 1.1 0 0 0-1.1-1.1h-5.8ZM14 3.019a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2h-6a1 1 0 0 1-1-1Zm1 16.962a1 1 0 0 0 0 2h6a1 1 0 1 0 0-2h-6Zm-9.988 1a1 1 0 0 1 1-1h6a1 1 0 0 1 0 2h-6a1 1 0 0 1-1-1ZM3 9.973a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1Zm16.63-3.977a1 1 0 1 1 2 0v2a1 1 0 0 1-2 0v-2Zm1.015 8.992a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1ZM2 20.981a1 1 0 1 1 2 0a1 1 0 0 1-2 0ZM3.011 2.019a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z" clip-rule="evenodd"/>`
	flagInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 24V2c8-3.524 11 4.644 20 0v12c-8 4.895-13-4.103-20 0"/>`
	flagFillInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" stroke="currentColor" stroke-width="2" d="M2 24V2c8-3.524 11 4.644 20 0v12c-8 4.895-13-4.103-20 0"/>`
	flowsInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5zm0 11a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3z"/>`
	folderInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 10V2h8l3 4h9v4H2Zm0 0h20v12H2V10Z"/>`
	folderCycleInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 18a5 5 0 0 1 5-5c1.985 0 3.7 1.156 4.5 3m.5 2a5 5 0 0 1-5 5c-1.985 0-3.699-1.156-4.5-3m5.5-4h4v-4m-6 8H7v4m-3-1H1V1h8l3 4h11v18h-4M1 9h22M4 23H1V1h8l3 4h11v18h-4M1 9h22"/>`
	folderOpenInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 2h6l2 4h6v4H5V2Zm-3 8h20l-3 12H5L2 10Z"/>`
	formAddInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 18V6m-6 6h12"/>`
	formAttachmentInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="m6 13.293l6.36-6.36c2.828-2.828 7.069 1.413 4.242 4.24l-6.361 6.36c-1.414 1.414-3.534-.706-2.12-2.12l6.36-6.36"/>`
	formCalendarInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 19h12V8H6v11Zm9-11V5v3ZM9 8V5v3Zm-3 3.5h12H6Z"/>`
	formCheckmarkInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="m6 13l4.2 3.6L18 7"/>`
	formClockInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0-10v4l3 1"/>`
	formCloseInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="m7 7l10 10M7 17L17 7"/>`
	formCutInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="m18 7.524l-7.857 6.286L18 7.524Zm0 8.38L10.143 9.62L18 15.905Zm-9.429-5.761a1.571 1.571 0 1 0 0-3.143a1.571 1.571 0 0 0 0 3.143Zm0 6.286a1.571 1.571 0 1 0 0-3.143a1.571 1.571 0 0 0 0 3.143Z"/>`
	formDownInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="m18 9l-6 6l-6-6"/>`
	formEditInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="m11.996 8.336l3.497 3.498l-3.497-3.498Zm5.54-.54a.994.994 0 0 1-.004 1.416l-7.451 7.451L6 17.83l1.166-4.08l7.451-7.452a.997.997 0 0 1 1.416-.005l1.504 1.504Z"/>`
	formFilterInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="m6 8l5.667 4.667V18h.666v-5.333L18 8V6H6z"/>`
	formFolderInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 18V6h4.8l1.8 2.4H18V18z"/>`
	formLocationInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 17s-4-3-4-6c0-2.5 2-4 4-4s4 1.5 4 4c0 3-4 6-4 6Zm0-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	formLockInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 11V8a3 3 0 0 1 6 0v3m-3 2v3m5 2v-7H7v7h10Z"/>`
	formNextInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="m9 6l6 6l-6 6"/>`
	formNextLinkInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 12.4h12M12.6 7l5.4 5.4l-5.4 5.4"/>`
	formPreviousInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="m15 6l-6 6l6 6"/>`
	formPreviousLinkInnerSVG       = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 12.4H6M11.4 7L6 12.4l5.4 5.4"/>`
	formRefreshInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17.333 9.333C16.398 7.36 14.358 6 12 6a6 6 0 1 0 6 6m.5-6v4h-4"/>`
	formScheduleInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 19h12V8H6v11Zm2-4h2h-2Zm3 0h5h-5Zm4-7V5v3ZM9 8V5v3Zm-3 3.5h12H6Z"/>`
	formSearchInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M13.8 13.8L18 18l-4.2-4.2ZM10.5 15a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z"/>`
	formSubtractInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 12h12"/>`
	formTrashInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7.5 9h9v10h-9V9ZM5 9h14M9.364 6h5v3h-5V6Zm1.181 5v6m3-6v6"/>`
	formUpInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m18 15l-6-6l-6 6"/>`
	formUploadInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 14.182v3.273h12v-3.273M12 6v8M8.182 9.818L12 6l3.818 3.818"/>`
	formViewInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 17c-2.727 0-6-2.778-6-5s3.273-5 6-5s6 2.778 6 5s-3.273 5-6 5Zm-1-5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`
	formViewHideInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 12h3c.5 2.5 3.273 5 6 5s5.5-2.5 6-5h3m-9 5v3m-4.5-4.5l-2 2m11-2l2 2"/>`
	forwardTenInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20.889 7.556C19.33 4.267 15.93 2 12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10m0-8v4h-4m-9 8V9l-2 .533M17 12c0-2-1-3.5-2.5-3.5S12 10 12 12s1 3.5 2.5 3.5S17 14 17 12Zm-2.5-3.5C16.925 8.5 17 11 17 12s0 3.5-2.5 3.5S12 13 12 12s.059-3.5 2.5-3.5Z"/>`
	freebsdInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M23.725.403c1.273 1.272-2.254 6.862-2.85 7.458c-.597.596-2.111.048-3.383-1.224c-1.272-1.272-1.82-2.787-1.224-3.383c.596-.596 6.185-4.123 7.457-2.85ZM5.885 1.75C3.943.647 1.179-.58.3.3c-.891.89.38 3.717 1.493 5.662A11.945 11.945 0 0 1 5.885 1.75Zm15.9 5.674c.179.606.147 1.108-.143 1.397c-.678.678-2.508-.044-4.158-1.614a8.227 8.227 0 0 1-.341-.323c-.596-.597-1.06-1.232-1.357-1.817c-.578-1.036-.723-1.952-.286-2.388c.238-.238.619-.303 1.083-.22c.303-.19.66-.404 1.053-.623a11.491 11.491 0 0 0-5.33-1.301C5.928.534.757 5.704.757 12.082S5.927 23.63 12.306 23.63c6.378 0 11.548-5.17 11.548-11.548c0-2.06-.54-3.991-1.485-5.666c-.204.373-.403.714-.584 1.007Z"/>`
	galleryInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h18v18H1V1Zm4 18v4h18V5.97h-4M6 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM2 18l5-6l3 3l4-5l5 6"/>`
	gamepadInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 6V2m0 4c2.498.044 4.006 0 5 0c2 0 4 .5 5 4s1 5.5 1 8s-2 3-4 3s-3.054-4-7-4s-5 4-7 4s-4-.5-4-3s0-4.5 1-8s3-4 5-4c.994 0 2.502.044 5 0Zm6 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-4-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM4 12h6M7 9v6"/>`
	gatewayInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 20v-5m0-6V4m-7 8h5m9 0h-5m-.969-4.031L12 9.344l-1.031-1.375h2.062zm-2.062 8.07L12 14.664l1.031 1.375H10.97zM6 13.031L4.625 12L6 10.969v2.062zm12-2.062L19.375 12L18 13.031V10.97zM12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1z"/>`
	gatsbyjsInnerSVG               = `<path fill="currentColor" d="M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12Zm10-11.908h-6.452v1.834h4.424c-.645 2.753-2.674 5.046-5.346 5.964L4.027 9.34c1.106-3.211 4.24-5.505 7.835-5.505c2.765 0 5.254 1.377 6.82 3.487l1.383-1.193C18.22 3.651 15.272 2 11.862 2C7.069 2 3.014 5.395 2 9.89L14.165 22C18.59 20.899 22 16.862 22 12.091v.001ZM2 12c0 2.593 1.018 5.092 2.963 7.037C6.908 20.982 9.5 22 12 22L2 12Z"/>`
	gemInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M7 1h-.414l-.293.293l-3 3L3 4.586v14.828l.293.293l3 3l.293.293h10.828l.293-.293l3-3l.293-.293V4.586l-.293-.293l-3-3L17.414 1H7ZM5 6v12h1V6H5Zm3 15h8v-1H8v1Zm11-3V6h-1v12h1ZM16 3H8v1h8V3Zm0 3v12H8V6h8Z" clip-rule="evenodd"/>`
	giftInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 11h18v12H3V11Zm-1 0V7h20v4H2Zm10 12V7v16ZM7 7h5s-2-5-5-5C3.5 2 3 7 7 7Zm10.184 0h-5s1.816-5 5-5c3.316 0 4 5 0 5Z"/>`
	githubInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M11.999 1C5.926 1 1 5.925 1 12c0 4.86 3.152 8.983 7.523 10.437c.55.102.75-.238.75-.53c0-.26-.009-.952-.014-1.87c-3.06.664-3.706-1.475-3.706-1.475c-.5-1.27-1.221-1.61-1.221-1.61c-.999-.681.075-.668.075-.668c1.105.078 1.685 1.134 1.685 1.134c.981 1.68 2.575 1.195 3.202.914c.1-.71.384-1.195.698-1.47c-2.442-.278-5.01-1.222-5.01-5.437c0-1.2.428-2.183 1.132-2.952c-.114-.278-.491-1.397.108-2.91c0 0 .923-.297 3.025 1.127A10.536 10.536 0 0 1 12 6.32a10.49 10.49 0 0 1 2.754.37c2.1-1.424 3.022-1.128 3.022-1.128c.6 1.514.223 2.633.11 2.911c.705.769 1.13 1.751 1.13 2.952c0 4.226-2.572 5.156-5.022 5.428c.395.34.747 1.01.747 2.037c0 1.47-.014 2.657-.014 3.017c0 .295.199.637.756.53C19.851 20.979 23 16.859 23 12c0-6.075-4.926-11-11.001-11"/>`
	globeInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 8s3.5 1 5 2s.564 2.42 1 3c.436.58 2-1 2 2s3 1 3 4s2.5 2.5 3 1s2.233-3.134 2-5c-.233-1.866-1-3-3-3s-3.5-.5-4-2s3-2 2-5s0-4 0-4m10 11c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Z"/>`
	golangInnerSVG                 = `<g fill="currentColor"><path d="M21.468 3.206c.713 1.474-1.455 1.631-1.344 2.289c.204 1.186.065 2.947-.092 4.68c-.27 2.937 1.927 9.97-2.65 12.796c-.862.538-2.91.834-4.93.88h-.019c-2.02-.046-4.346-.342-5.208-.88c-4.568-2.826-2.372-9.859-2.631-12.796c-.167-1.733-.306-3.494-.093-4.68c.111-.658-2.057-.806-1.343-2.289c.611-1.27 2.122-.148 2.501-.547C7.596.65 9.95.111 11.803.065h1.02c1.843.093 4.197.593 6.143 2.594c.38.399 1.881-.722 2.502.547Zm-10.036 7.7c-.083.038-.222.946.204.992c.288.028 1.075.12 1.26 0c.362-.232.325-.788.121-.927c-.343-.222-1.492-.11-1.585-.064Zm-2.352-6.8c-.955-.111-2.558.778-2.799 2.669c-.26 1.964 2.057 3.91 4.281 1.964c1.196-1.047 1.613-4.272-1.482-4.633Zm6.45 0c-3.096.361-2.679 3.586-1.483 4.633c2.223 1.946 4.54 0 4.28-1.964c-.231-1.89-1.834-2.78-2.798-2.669Z"/><path d="M12.313 8.988c.584-.083 1.983.584 1.89 1.372c-.111.917-3.605 1.038-3.79-.056c-.111-.667.417-1.093 1.9-1.316Zm7.885 7.386c-.287-.01-.472-.39-.472-.621c0-.417.055-.908.361-1.121c.63-.436 1.13 1.751.111 1.742Zm-15.78 0c-1.02.009-.52-2.178.111-1.742c.306.213.361.704.361 1.12c0 .233-.185.612-.472.622Zm13.844 6.718c.185.269.194.454-.102.593c-1.13.519-2.14-.176-1.76-.37c.805-.408 1.306-1.048 1.862-.223Zm-11.908.093c.556-.825 1.057-.186 1.863.222c.38.195-.63.89-1.76.37c-.297-.138-.288-.324-.103-.592Z"/><path d="M19.735 3.429c.046-.195.528-.288.797.046c.324.398-.473.945-.51.723c-.092-.603-.342-.575-.287-.77Zm-14.854 0c.056.194-.195.166-.287.769c-.037.222-.834-.325-.51-.723c.26-.334.741-.241.797-.046Zm10.86 4.512a1.056 1.056 0 1 1 0-2.112a1.056 1.056 0 0 1 0 2.112Zm-.278-1.177a.334.334 0 1 0 0-.667a.334.334 0 0 0 0 .667Zm-4.105 2.67c.111-.427.5-.566.816-.575c.806-.019 1.112.417 1.14.76c.055.583-2.206.722-1.956-.186ZM8.875 7.94a1.056 1.056 0 1 1 0-2.112a1.056 1.056 0 0 1 0 2.112Zm.278-1.177a.334.334 0 1 0 0-.667a.334.334 0 0 0 0 .667Z"/></g>`
	googleInnerSVG                 = `<path fill="currentColor" d="M5.266 9.765A7.077 7.077 0 0 1 12 4.909c1.69 0 3.218.6 4.418 1.582L19.91 3C17.782 1.145 15.055 0 12 0C7.27 0 3.198 2.698 1.24 6.65l4.026 3.115Zm10.774 8.248c-1.09.703-2.474 1.078-4.04 1.078a7.077 7.077 0 0 1-6.723-4.823l-4.04 3.067A11.965 11.965 0 0 0 12 24c2.933 0 5.735-1.043 7.834-3l-3.793-2.987Z"/><path fill="currentColor" d="M19.834 21c2.195-2.048 3.62-5.096 3.62-9c0-.71-.109-1.473-.272-2.182H12v4.637h6.436c-.317 1.559-1.17 2.766-2.395 3.558L19.834 21ZM5.277 14.268A7.12 7.12 0 0 1 4.909 12c0-.782.125-1.533.357-2.235L1.24 6.65A11.934 11.934 0 0 0 0 12c0 1.92.445 3.73 1.237 5.335l4.04-3.067Z"/>`
	googlePlayInnerSVG             = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M3 2v20l18-10z"/><path d="m3 2l11 14M3 22L14 8"/></g>`
	googlePlusInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M.02 11.203c.066-3.906 3.676-7.327 7.603-7.197c1.882-.087 3.65.728 5.091 1.872a27.061 27.061 0 0 1-1.932 1.991C9.05 6.68 6.586 6.338 4.853 7.713c-2.478 1.705-2.59 5.731-.207 7.567c2.318 2.092 6.7 1.053 7.34-2.15c-1.451-.022-2.907 0-4.36-.048c-.003-.861-.007-1.723-.003-2.585a707.94 707.94 0 0 1 7.286.008c.145 2.027-.124 4.185-1.376 5.86c-1.896 2.655-5.702 3.43-8.672 2.292c-2.98-1.13-5.092-4.26-4.84-7.454m19.623-2.882h2.165c.004.721.007 1.445.015 2.165c.724.008 1.452.008 2.176.015v2.154c-.724.007-1.448.011-2.176.018c-.008.724-.011 1.445-.015 2.165H19.64c-.007-.72-.007-1.44-.014-2.161l-2.177-.022v-2.154c.725-.007 1.449-.01 2.177-.015c.003-.724.01-1.444.018-2.165"/>`
	googleWalletInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M12.3 20.206a2.365 2.365 0 0 1-2.722 1.046a2.386 2.386 0 0 1-1.687-2.377c-.34-3.8-2.458-7.06-5.693-8.798a2.258 2.258 0 0 1-.93-3.058c.4-.748 1.162-1.2 2-1.2c.375 0 .749.09 1.071.271c2.145 1.143 4.035 2.785 5.4 4.658c-.185-1.543-.673-3.047-1.49-4.36a2.393 2.393 0 0 1-.193-2.04a2.339 2.339 0 0 1 1.393-1.432a2.05 2.05 0 0 1 .852-.155c.464 0 .916.142 1.29.387c.035.022.068.043.101.066c-.036-.04-.1-.079-.1-.079c2.782 1.701 5.153 4.045 6.791 6.66a19.654 19.654 0 0 0-1.773-6.22a2.497 2.497 0 0 1 1.188-3.33A2.37 2.37 0 0 1 18.868 0a2.53 2.53 0 0 1 2.259 1.432a24.472 24.472 0 0 1 1.767 5.084c.4 1.78.62 3.652.633 5.51c0 1.884-.22 3.716-.62 5.51c-.103.477-.22.929-.361 1.406c-.439 1.587-.942 2.813-1.445 3.755A2.504 2.504 0 0 1 18.907 24c-.374 0-.735-.09-1.07-.245c-.75-.349-1.2-.994-1.368-1.703a2.49 2.49 0 0 1-.065-.581c0-.542-.026-.903-.026-.903c0-2.695-.644-5.261-1.83-7.518a15.707 15.707 0 0 1-2.247 7.156Z"/>`
	graphQlInnerSVG                = `<path fill="currentColor" d="m14.334 2.852l4.71 2.72c.142-.15.308-.28.495-.39A2.234 2.234 0 0 1 22.585 6a2.234 2.234 0 0 1-1.546 3.31v5.38a2.228 2.228 0 1 1-2.166 3.544l-4.626 2.67A2.231 2.231 0 0 1 12.192 24a2.227 2.227 0 0 1-2.12-2.925l-4.678-2.701a2.23 2.23 0 1 1-2.27-3.634l.001-5.48A2.232 2.232 0 0 1 1.798 6a2.232 2.232 0 0 1 3.458-.51l4.765-2.752A2.227 2.227 0 0 1 12.192 0a2.227 2.227 0 0 1 2.142 2.852Zm-.493.88c-.05.054-.103.106-.157.155l6.278 10.875l.069-.02V9.255a2.226 2.226 0 0 1-1.501-2.816L13.84 3.73Zm-3.28.019a2.237 2.237 0 0 1-.093-.107L5.816 6.33a2.227 2.227 0 0 1-1.683 2.976v5.386l.1.02L10.56 3.751Zm2.24.624a2.237 2.237 0 0 1-1.401-.06L5.14 15.154a2.238 2.238 0 0 1 .74 1.12h12.621a2.236 2.236 0 0 1 .595-.996L12.8 4.375Zm.949 15.8l4.739-2.736a2.247 2.247 0 0 1-.035-.157H5.93c-.013.07-.03.14-.049.208l4.716 2.723a2.22 2.22 0 0 1 1.596-.672c.607 0 1.156.242 1.558.634Z"/>`
	gremlinInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm10 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-5 2c.5 0 1-.321 1-1c0-.113-2-.075-2 0c0 .679.5 1 1 1Zm-9 8s.003-8 0-12c.003-4-.447-8 9-8s8.997 4 9 8c-.003 4 0 12 0 12m-.932-17c3.26 0 3.58-3.58 2.046-5.114C20.58.352 17 .673 17 3.932M7 19c0-.203 2.5 1.152 5 1c2.5.152 5-1.499 5-1c0 .802-1.5 2-5 2s-5-1.493-5-2ZM3.945 7C.605 7 .38 3.42 1.908 1.886C3.435.352 7 .673 7 3.932"/>`
	gridInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 1v22m8-22v22M1 8h22M1 16h22M1 1h22v22H1V1Z"/>`
	grommetInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="4" d="M12 2C6.485 2 2 6.485 2 12s4.485 10 10 10s10-4.485 10-10S17.515 2 12 2Z"/>`
	groupInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 13a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm-6 9v-3a6 6 0 1 1 12 0v3M13 5c.404-1.664 2.015-3 4-3c2.172 0 3.98 1.79 4 4c-.02 2.21-1.828 4-4 4h-1h1c3.288 0 6 2.686 6 6v2M11 5c-.404-1.664-2.015-3-4-3c-2.172 0-3.98 1.79-4 4c.02 2.21 1.828 4 4 4h1h-1c-3.288 0-6 2.686-6 6v2"/>`
	growInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 16V7m0 4c0-3 0-6-7-6c0 3 0 6 7 6Zm-8 5h16m-2 0l-2 7H8l-2-7m6-9c0-3 0-6 7-6c0 3 0 6-7 6Z"/>`
	hadoopInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M22.615 8.515c-.162.64-.49 1.186-1.168 1.428c-.332.118-.547-.004-.886-.102c.333.032.524.083.822-.064a2.72 2.72 0 0 0 1.232-1.262Zm1.355 4.215c-.08.777-.287 1.532-.656 2.172a3.548 3.548 0 0 1-1.407 1.385c-.417.222-.902.35-1.403.403a5.83 5.83 0 0 1-1.56-.06l-.07.234c-.03.097-.06.194-.088.29c-.103.337-.347.64-.652.865c-.314.232-.698.383-1.055.4a1.443 1.443 0 0 1-.84-.203c-.226-.133-.42-.318-.631-.522l-.277-.267l-.187-.18l-.02.123c.098.197.167.382.237.572c.045.12.09.241.145.373c.082.195.144.358.187.527c.043.17.066.344.072.557l.015.542c.004.17.008.341.014.512c.075.136.12.23.142.334a1.3 1.3 0 0 1 .01.393c-.028.394-.158.614-.377.739c-.206.116-.478.133-.814.132l-.235-.001c-.327 0-.832 0-.986-.01a3.395 3.395 0 0 1-.525-.066a.762.762 0 0 1-.28-.121c-.112.09-.21.171-.309.252c-.29.24-.584.483-.702.561a4.06 4.06 0 0 0-.125.206c-.121.214-.211.371-.498.527c-.564.304-1 .278-1.382.043c-.365-.224-.666-.635-.98-1.12c-.15-.23-.39-.61-.506-.983c-.129-.412-.11-.812.3-1.023l.058-.03l.123-.064a7.16 7.16 0 0 1-1.393-.05l.012.086c.024.208.04.344.02.48c-.02.14-.075.265-.193.442l-.058.082c-.11.163-.175.26-.326.594c.015.165.02.298.009.43a2.702 2.702 0 0 1-.093.477c-.13.478-.61.588-1.142.554c-.488-.03-1.023-.184-1.332-.268l-.174-.046a4.952 4.952 0 0 1-.778-.248c-.473-.216-.777-.555-.481-1.204a4.04 4.04 0 0 0 .127-.314c.014-.037.024-.073.036-.11a1.48 1.48 0 0 1-.746-.103c-.437-.166-.853-.483-1.108-.738a2.696 2.696 0 0 1-.482-.65a2.495 2.495 0 0 1-.252-.748c-.069-.408-.02-.455.22-.687l.087-.083l.357-.354c.127-.126.256-.253.382-.38l.175-.472l.111-.305a5.999 5.999 0 0 1-.083-.596a9.29 9.29 0 0 1-.037-1.14c.01-.256.034-.492.073-.718c-.2.044-.402.041-.586-.017a.905.905 0 0 1-.55-.489a2.162 2.162 0 0 1-.116-.299a5.379 5.379 0 0 1-.088-.322a.993.993 0 0 1-.286-.525a1.016 1.016 0 0 1 .098-.612a2.04 2.04 0 0 1 .384-.536c.425-.442 1.03-.774 1.31-.823l.212-.037l-.07.203c-.037.102-.08.21-.12.315c-.013.03-.024.061-.037.095a.69.69 0 0 1 .243.406c.057.25.02.537-.028.753l-.058.258l-.163-.207c-.04-.049-.073-.1-.104-.149a1.443 1.443 0 0 0-.086-.122c-.022.336-.063.672-.302.902c.007.028.01.044.015.047c.004.002.044-.02.142-.058c.169-.065.318-.172.46-.298c.144-.13.276-.28.41-.426c.152-.294.325-.573.528-.837c.205-.27.441-.526.713-.766a6.577 6.577 0 0 1 2.047-1.266c.72-.278 1.521-.453 2.506-.611a12.88 12.88 0 0 1 .809-.78a1.96 1.96 0 0 1 .562-.368c.158-.065.313-.09.488-.103c.301-.4.575-.744.898-1.042a4.13 4.13 0 0 1 1.247-.792c1.082-.456 1.999-.694 2.866-.64c.873.055 1.69.405 2.566 1.128c.165.136.334.284.502.43c.367.324.727.64 1.15.904c.16.099.294.185.427.286c.132.1.253.212.38.362c.153.183.292.366.415.562c.103.163.193.337.272.527c.18-.072.369-.152.521-.28c.131-.113.24-.273.35-.432c.097-.143.195-.286.312-.403a.476.476 0 0 1 .128-.095c.178-.094.404-.079.618-.005c.202.07.397.192.53.323c.06.06.11.123.141.185c.266.509.458 1.216.573 1.943c.13.818.164 1.665.102 2.288Zm-3.725-3.348a9.904 9.904 0 0 0-.113.003a.614.614 0 0 1 .13.19a.414.414 0 0 1 .052-.043a6.955 6.955 0 0 0-.069-.15ZM1.499 11.654c.098-.011.16.04.236.116c-.014-.181-.048-.316-.176-.39l-.025.088c-.016.059-.026.12-.035.186Zm.391 1.98a9.669 9.669 0 0 1 .228-.666a1.715 1.715 0 0 1-.518.33c-.524.21-.513.029-.739-.426c.48-.383.258-.853.41-1.382c.036-.122.087-.24.16-.397C.93 11.395.006 12.24.623 12.88c.067.241.118.43.21.62c.189.395.682.294 1.057.135Zm1.389 6.399c-.537-.811-1.021-1.778-1.305-2.76l-.158.421l-.72.73c-.178.18-.213.187-.177.442c.048.35.255.723.564 1.038c.282.288 1.258 1.04 1.641.528c.082-.11.12-.245.155-.384v-.015Zm11.783-1.536c-.085-.202-.131-.39-.19-.582c-.128.457-.3.886-.456 1.38c-.148.467-.909 1.822-1.356 2.151c.087.065.247.09.532.115c.201.017.995.021 1.195.024c.439.008.606-.025.65-.525c.022-.245.002-.29-.12-.505l-.032-1.148c-.01-.362-.08-.572-.223-.91Zm7.89-9.624c-.103-.247-.51-.618-.77-.366c-.262.255-.459.614-.744.84c-.389.305-1.057.2-.936.85c.088.472.114 1.001.012 1.442c-.096.423-.196 1.024-.417 1.308c.065-.24.175-.867.206-1.29a3.72 3.72 0 0 0-.053-.791l-.147.003c-.178.004-.469.26-.547.419c-.218.444-.233.856-.477 1.279c.197-.473.086-.899.266-1.405c.063-.179.223-.333.409-.427c-.104-.015-.207.03-.36.065c-.651.156-.622.668-1.023 1.157c.387-.709.19-1.179 1.028-1.37c.275-.063.477-.088.658.057h.011l.143-.012c-.023-.13-.049-.258-.074-.379a3.908 3.908 0 0 0-1.13.02c.085-.037.168-.071.248-.1c.125-.046.247-.082.367-.103a.238.238 0 0 0 .023-.079a.243.243 0 0 0-.218-.263a.238.238 0 0 0-.259.218a.303.303 0 0 0 .025.13a.594.594 0 0 1 .015-.569a.725.725 0 0 1-.42.006a2.616 2.616 0 0 0 .62-.21a.59.59 0 0 1 .554-.034c.079.023.167.054.273.093a6.158 6.158 0 0 0-.562-.948c-.208-.289-.382-.388-.689-.57a6.177 6.177 0 0 1-.196-.12c-.227-.042-.48-.083-.698-.041c.18-.07.315-.1.473-.112c-.434-.311-.809-.66-1.237-1.02c-1.653-1.39-2.99-1.292-5.095-.378c-.854.372-1.16.746-1.686 1.418h.017a23.346 23.346 0 0 0 1.615-.084c-.568.133-1.112.239-1.67.323c-.615.094-.804.078-1.269.53c-.855.833-1.572 1.774-2.355 2.618c-.453.49-.679.98-.945 1.537c-.264.554-.227.764.127 1.243c.362.491.58.713.734 1.148c.4-.617.882-1.13 1.39-1.743a25.384 25.384 0 0 0-1.136 1.97c-.214.42-.319.635-.312 1.1c.401.462.653.709 1.025.81c.401.111.78.093 1.142-.095c.902-.467 1.759-1.076 2.76-1.153c.517-.92.372-2.079.186-3.127c-.154-.87-.115-1.638.093-2.511c.052.848.12 1.635.288 2.467c.245 1.218.297 2.27-.216 3.528c-1.106.027-2.014.675-2.995 1.177c-.441.225-.854.248-1.348.13c-.503-.119-.889-.507-1.407-1.074c.05-.429.107-.695.26-1.01c-.185-.557-.414-.812-.838-1.381c-.49-.66-.499-.946-.128-1.697c.268-.544.519-1.047.942-1.543c.417-.492.8-.95 1.188-1.384c-1.61.245-2.638.633-3.738 1.598c-.815.715-1.324 1.604-1.687 2.63c-.23.648-.294 1.3-.185 2.36c.139 1.364.864 2.904 1.644 4.027c-.095.7-.219 1.226-.414 1.726c-.221.566.66.74 1.085.856c.345.095 1.656.455 1.786-.01c.077-.278.077-.473.058-.767c.234-.466.305-.528.472-.804c.178-.295.193-.413.195-.76c0-.447-.011-1.245.008-1.534c.06.258.132.618.192.963a7.332 7.332 0 0 0 2.789-.036c.012-.065.033-.136.057-.223c.063-.218.188-.436.25-.654l-.05.644c-.018.221-.019.4.009.62l.044.364c-.067-.11-.172-.218-.233-.33c-.415.274-.643.428-1.05.643c-.431.228.062 1.043.24 1.326c.39.62.857 1.355 1.583.971c.238-.125.378-.417.537-.626c.213-.126 1.343-1.08 1.523-1.183c.254-.146 1.167-1.642 1.283-2.015c.255-.827.555-1.456.64-2.318c-.582-.245-.864-.517-1.281-1.007a5.436 5.436 0 0 0 1.723.867c.23.203.463.408.704.623c.35.312.672.63 1.18.639c.55.01 1.17-.41 1.33-.931l.184-.602a5.48 5.48 0 0 1-.318-.087c-.114.274-.105.38-.158.582c-.093.344-.567.583-1.063.463c-.176-.043-.28-.067-.343-.113c.05.136.147.251.328.284c.2.037.347.053.638-.013c-.36.187-.475.195-.73.138c-.67-.149-.434-.956-.288-1.45c.092-.31.06-.64-.009-.951c.254.16.455.302.744.4c1.354.455 2.96.92 4.322.159c1.068-.596 1.654-1.919 1.785-3.13c.118-1.08-.085-2.765-.527-3.836Zm-6.909 6.503c-.026.168-.062.376-.089.544c.072-.19.157-.41.24-.591c.087-.192.136-.209.32-.31c.133-.072.372-.17.503-.242c-.135.023-.378.072-.512.095c-.362.061-.405.148-.462.504ZM9.865 8.908c-.419.415-.826 1.84-.956 2.409c.204-.474.718-1.797 1.112-2.134a1.39 1.39 0 0 1 .266-.192c-.282.473-.261.59-.162 1.224c.084-.644.307-.893.672-1.373c.4-.101.776-.22 1.187-.381c-.464.052-.926.1-1.39.144c-.39.036-.452.03-.729.303Zm6.964 1.58a.767.767 0 0 0-.432-.392c.167-.09.33-.183.457-.298c-.373.173-.807.13-1.127.347c-.283.19-.672.794-.959 1.05c.207-.08.407-.22.591-.367a.762.762 0 0 0 .286.589a.314.314 0 0 1-.06-.085a.323.323 0 1 1 .6-.231a2.06 2.06 0 0 0-.412.403a1.729 1.729 0 0 0-.25.457a4.386 4.386 0 0 1 2.081-1.345a2.255 2.255 0 0 0-.705.118a.763.763 0 0 0-.07-.247Zm-2.666-.584c.126-.556.349-1.092 1.212-1.506c-1.144.286-1.359.767-1.212 1.506Z"/>`
	haltInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 23h7c3 0 4-2 4-5V6c0-2-1-2-1.5-2S16 4 16 6v7v-9c0-1 0-2-1.5-2S13 3 13 4v9V3c0-1 0-2-1.5-2S10 2 10 3v10v-9c0-1 .03-2-1.5-2C7 2 7 3 7 4v14v-4c0-1-.55-2-2-2H4v6c0 3.962 2 5.024 4 5Z"/>`
	helpInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-5v-3c0-2 1-2 3-3s3-2.842 3-5A6 6 0 1 0 6 7"/>`
	helpOptionInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M2 7a5 5 0 0 1 5-5h10a5 5 0 0 1 5 5v10a5 5 0 0 1-5 5H7a5 5 0 0 1-5-5V7Zm5-3a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H7Zm5 3.5c-1.448 0-2 1.052-2 1.5a1 1 0 1 1-2 0c0-1.552 1.448-3.5 4-3.5c1.156 0 2.17.289 2.909.935C15.664 7.096 16 8.017 16 9c0 1.188-.306 2.028-.882 2.668a4.313 4.313 0 0 1-.828.693a14.92 14.92 0 0 1-.368.235l-.018.011c-.124.078-.236.15-.347.223c-.35.236-.493.412-.55.526c-.046.088-.077.205-.028.438a1 1 0 0 1-1.958.412c-.126-.6-.083-1.197.204-1.759c.273-.535.718-.942 1.218-1.277c.14-.094.277-.18.399-.257l.009-.006c.122-.076.228-.142.328-.21c.205-.136.346-.25.453-.368c.174-.193.368-.516.368-1.329c0-.517-.164-.846-.409-1.06c-.261-.229-.747-.44-1.591-.44Zm0 8.281a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z" clip-rule="evenodd"/>`
	herokuInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M20.443 0H3.162A2.162 2.162 0 0 0 1 2.162V21.84C1 23.034 1.97 24 3.162 24h17.28a2.159 2.159 0 0 0 2.16-2.16V2.162A2.16 2.16 0 0 0 20.442 0Zm.958 21.84a.96.96 0 0 1-.958.96H3.162a.961.961 0 0 1-.962-.96V2.162c0-.532.432-.962.962-.962h17.28c.53 0 .96.43.96.962V21.84Zm-15-1.439l2.701-2.4L6.4 15.6v4.8Zm9.757-9.729c-.486-.488-1.373-1.071-2.856-1.071c-1.627 0-3.303.424-4.5.812V3.6H6.4v10.41l1.697-.769c.028-.013 2.763-1.239 5.205-1.239c1.218 0 1.488.67 1.501 1.231v7.17h2.398v-7.2c.003-.155-.012-1.486-1.043-2.53M13 7.5h2.401c1.085-1.228 1.637-2.536 1.8-3.899h-2.399c-.267 1.36-.858 2.66-1.802 3.9"/>`
	hideInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 17c-2.727 0-6-2.778-6-5s3.273-5 6-5s6 2.778 6 5s-3.273 5-6 5Zm-1-5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm9-7L4 19"/>`
	historyInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1C7.563 1 4 4 2 7.5M1 1v7h7m8 9l-4-4V6"/>`
	homeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="m1 11l11-9l11 9m-8 12v-8H9v8m-5 0V9m16 14V9"/>`
	homeOptionInnerSVG             = `<rect width="18" height="18" x="3" y="3" fill="none" stroke="currentColor" stroke-width="2" rx="4"/>`
	homeRoundedInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 22V9.76a2 2 0 0 1 .851-1.636l9.575-6.72a1 1 0 0 1 1.149 0l9.574 6.72A2 2 0 0 1 23 9.76V22a1 1 0 0 1-1 1h-5.333a1 1 0 0 1-1-1v-5.674a1 1 0 0 0-1-1H9.333a1 1 0 0 0-1 1V22a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1z"/>`
	hortonInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M2.3 12.977h-.157c-.319.989-.755 1.957-.932 2.97c-.222 1.292-.238 2.624-.34 3.936c-.027.339.126.517.465.52c1.526.02 1.352.081 1.602-1.198c.017-.089.029-.178.057-.267c.52-1.658.545-3.285-.27-4.875c-.178-.344-.283-.723-.424-1.086Zm16.606-5.922a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Zm2.176 2.737c.206.44.153 1.021.15 1.538c-.005.924-.057 1.849-.098 2.773a609.04 609.04 0 0 1-.125 2.922c-.577-.355-1.037-.605-1.457-.908c-.141-.1-.242-.335-.254-.517c-.097-1.097-.194-2.2-.226-3.301c-.017-.695-.202-1.207-.925-1.405c-.383-.1-.58-.367-.726-.735c-.614-1.525-1.251-3.047-1.93-4.686c1.163.214 2.237.388 3.298.618c.19.04.375.278.488.472c.622 1.066 1.276 2.115 1.805 3.23Zm.569 7.225c0 .016.016.028.04.08c.771.057 1.615.352 2.172-.718c-.364-.105-.686-.201-1.01-.286c-1.202-.315-1.202-.315-1.202.924Zm-6.3-3.205c-.098-.347-.271-.496-.63-.549c-.436-.068-.933-.117-1.272-.359c-1.497-1.086-2.942-2.244-4.424-3.354c-.4-.303-.496-.618-.33-1.082c.245-.694.455-1.4.718-2.208c-.226.025-.331.029-.432.049c-1.308.29-2.604.646-3.928.86c-.86.14-1.38.577-1.9 1.25C1.888 10.055.92 11.79.564 13.85c-.13.754-.371 1.489-.565 2.232c.06.04.125.08.186.12c.125-.144.302-.27.367-.435c.404-1.09.779-2.188 1.175-3.278c.205-.565.387-.61.762-.129c.243.307.465.658.594 1.021c.69 1.95 1.352 3.908 2.026 5.861c.408 1.19.412 1.17 1.683 1.175c.48 0 .654-.17.614-.638c-.085-1.013-.138-2.026-.23-3.035c-.045-.517.185-.755.662-.884c1.174-.315 2.34-.654 3.507-.99c.735-.205.807-.18 1.17.473c.764 1.36 1.563 2.704 2.261 4.101c.38.767.884 1.062 1.715.973c.61-.065.658-.073.493-.682c-.545-1.974-1.102-3.948-1.635-5.922Zm-.118-7.54c-.38-1.04-.928-1.61-2.054-1.719c-.719-.069-1.417-.36-2.128-.553h-.217c-.097.214-.214.42-.287.638c-.408 1.223-.787 2.454-1.207 3.669c-.157.452-.056.742.327 1.021c1.062.787 2.127 1.578 3.148 2.422c1.635.989 2.24.989 2.975.985c.735-.004 1.32 0 2.059 0c-.093-.263-.15-.448-.222-.626c-.808-1.941-1.68-3.859-2.394-5.837Zm-3.786 9.18c-.287.035-.658.052-.682.536c-.053 1.082-.15 2.163-.234 3.354c.02.25-.073.78.113.912c.274.19.759.154 1.138.101c.153-.02.34-.307.392-.508c.14-.521.153-1.082.327-1.59c.318-.945.12-1.777-.453-2.531c-.12-.162-.411-.3-.601-.275ZM24 15.725c-.093-.566-.145-1.082-.274-1.579c-.033-.125-.291-.238-.46-.258c-.11-.016-.32.113-.36.222c-.286.75.19 1.506 1.094 1.615Z"/>`
	hostInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17 4h1v1h-1V4ZM3 1h18v22H3V1Zm0 12h18H3Zm0 5h18H3ZM3 8h18H3Z"/>`
	hostMaintenanceInnerSVG        = `<path fill="none" stroke="currentColor" stroke-width="2" d="m14 23l6-6m1-3a2 2 0 1 0 2 2M17 4h1v1h-1V4Zm-7 19H3V1h18v10M3 13h14M3 18h10M3 8h18"/>`
	hpInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M8.421 0L5 15.127h2.138L10.56 0H8.421Zm8.993 8.873l-1.496 6.225h2.138l1.496-6.225h-2.138Zm-3.635 0L10.36 24h2.138l3.42-15.127H13.78Zm-3.647 0l-1.497 6.225h2.138l1.496-6.225h-2.137Z"/>`
	hpeInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M2 6h44v12H2V6Zm3 3h38v6H5V9Z" clip-rule="evenodd"/>`
	hpeLabsInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M14 23H5V2h12v13h-4V6H9v13h11"/>`
	hpiInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M15.793 15.333a.5.5 0 0 0 .442-.315l2.196-6.37c.06-.173-.04-.315-.225-.315h-1a.502.502 0 0 0-.443.315l-2.206 6.37c-.06.174.041.315.225.315h1.01ZM24 12c0 6.627-5.373 12-12 12c-.183 0-.36-.02-.542-.027l2.42-6.991a.503.503 0 0 1 .443-.315H16c2.607 0 2.631-.865 2.906-1.66c.642-1.862 1.75-5.073 1.964-5.701C21.178 8.41 21.28 7 19.001 7H15a.504.504 0 0 0-.443.315L8.925 23.587C3.79 22.227 0 17.562 0 12C0 6.694 3.448 2.2 8.223.615L2.776 16.351c-.06.174.04.316.224.316h2a.503.503 0 0 0 .443-.316l2.666-7.703a.501.501 0 0 1 .442-.315h.989c.183 0 .284.142.224.315l-2.656 7.703c-.059.174.042.316.225.316h2a.501.501 0 0 0 .442-.316l2.427-7.036C12.765 7.682 12.312 7 10.346 7H9.013c-.183 0-.284-.142-.224-.315L11.086.046C11.39.023 11.691 0 12 0c6.627 0 12 5.373 12 12Z"/>`
	htmlFiveInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 2h18v16l-9 4l-9-4V2Zm14 4H8v5h8v5l-4 1.5L8 16v-2"/>`
	iceCreamInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17 8c2 0 2-1.5.5-1.5c0-3-2-5.5-5.5-5.5S6.5 3.5 6.5 6.5C5 6.5 5 8 7 8m0 0h10l-5 13L7 8Z"/>`
	imageInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 3h22v18H1V3Zm5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm17 6l-5-6l-6 7l-3-3l-8 8"/>`
	impactInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m13 2l9 4v11l-9 5V2Zm9 4l-9 5l9-5ZM9 22V2v20Zm0-10L3 5l6 7Zm0 0H1h8Zm0 0l-6 7l6-7Z"/>`
	inProgressInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h22M10 4.5h4V6c0 1-2 2-2 2s-2-1-2-2V4.5ZM5 1v5c0 3 5 3.235 5 6s-5 3-5 6v5M19 1v5c0 3-5 3.235-5 6s5 3 5 6v5M1 23h22M8 21c0-2 4-4 4-4s4 2 4 4v2H8v-2Z"/>`
	inboxInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 13L6 2h12l5 11v9H1v-9Zm0 0h7v3h8v-3h7"/>`
	indicatorInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm7-3A7 7 0 0 0 5 8c0 1.933.5 3 2 5s3 3.5 3 6v4h4v-4c0-2.5 1.5-4 3-6s2-3.067 2-5Z"/>`
	infoInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M15 17c0-3 4-5 4-9s-3-7-7-7s-7 3-7 7s4 6 4 9v3c0 2 1 3 3 3s3-1 3-3v-3Zm-6 1h6"/>`
	inheritInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="m17 18l-5-3l5 3ZM7 18l5-3v-4m5 9a3 3 0 1 0 6 0a3 3 0 0 0-6 0ZM4 17a3 3 0 1 0 0 6a3 3 0 0 0 0-6ZM17 6a5 5 0 1 1-10.001-.001A5 5 0 0 1 17 6Z"/>`
	insecureInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 6.919V6a4.724 4.724 0 0 1 5-5a4.724 4.724 0 0 1 5 5v5.052M12 23a7 7 0 1 0-7-7a7 7 0 0 0 7 7zm2.985-7h-5.97"/>`
	inspectInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5.5 21a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM1 16V6.5A4.5 4.5 0 0 1 5.5 2H6m17 14V6.5A4.5 4.5 0 0 0 18.5 2H18m.5 19a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM10 17s0-2 2-2s2 2 2 2"/>`
	instagramInnerSVG              = `<path fill="currentColor" d="M17.318.077c1.218.056 2.06.235 2.838.537a5.36 5.36 0 0 1 1.956 1.274a5.36 5.36 0 0 1 1.274 1.956c.302.779.481 1.62.537 2.838C23.992 8.192 24 8.724 24 12s-.008 3.808-.077 5.318c-.056 1.218-.235 2.06-.537 2.839a5.36 5.36 0 0 1-1.274 1.955a5.359 5.359 0 0 1-1.956 1.274c-.779.302-1.62.481-2.838.537c-1.51.069-2.041.077-5.318.077c-3.277 0-3.809-.008-5.318-.077c-1.218-.056-2.06-.235-2.839-.537a5.359 5.359 0 0 1-1.955-1.274a5.36 5.36 0 0 1-1.274-1.956c-.302-.779-.481-1.62-.537-2.838C.008 15.81 0 15.278 0 12c0-3.277.008-3.81.077-5.318c.056-1.218.235-2.06.537-2.838a5.36 5.36 0 0 1 1.274-1.956A5.36 5.36 0 0 1 3.843.614C4.623.312 5.464.133 6.682.077C8.19.008 8.722 0 12 0c3.277 0 3.81.008 5.318.077ZM12 2.667c-3.24 0-3.736.007-5.197.074c-.927.042-1.483.16-1.994.359a2.73 2.73 0 0 0-1.036.673A2.707 2.707 0 0 0 3.1 4.809c-.198.51-.317 1.067-.359 1.994C2.674 8.264 2.667 8.76 2.667 12s.007 3.736.074 5.197c.042.927.16 1.483.359 1.993c.17.436.35.713.673 1.037c.324.324.601.504 1.036.673c.51.198 1.067.317 1.994.359c1.462.067 1.958.074 5.197.074c3.24 0 3.735-.007 5.197-.074c.927-.042 1.483-.16 1.994-.359a2.73 2.73 0 0 0 1.036-.673c.324-.324.504-.601.673-1.036c.198-.51.317-1.067.359-1.994c.067-1.462.074-1.958.074-5.197s-.007-3.735-.074-5.197c-.042-.927-.16-1.483-.359-1.993a2.709 2.709 0 0 0-.673-1.037A2.708 2.708 0 0 0 19.19 3.1c-.51-.198-1.067-.317-1.994-.359c-1.461-.067-1.957-.074-5.197-.074Zm0 15.555a6.222 6.222 0 1 1 0-12.444a6.222 6.222 0 0 1 0 12.444Zm0-2.666a3.556 3.556 0 1 0 0-7.112a3.556 3.556 0 0 0 0 7.112Zm6.222-8.445a1.333 1.333 0 1 1 0-2.667a1.333 1.333 0 0 1 0 2.667Z"/>`
	installInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 13.5v4L12 22l-7-4.5v-4m7 8.5v-8.5m6.5-5l-6.5-4L15.5 2L22 6l-3.5 2.5Zm-13 0l6.5-4L8.5 2L2 6l3.5 2.5Zm13 .5L12 13l3.5 2.5l6.5-4L18.5 9Zm-13 0l6.5 4l-3.5 2.5l-6.5-4L5.5 9Z"/>`
	installOptionInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 6v10V6Zm0-5c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1Zm5 11l-5 5l-5-5"/>`
	integrationInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 21h18V9H5m14 6H1V3h18"/>`
	internetExplorerInnerSVG       = `<path fill="currentColor" fill-rule="evenodd" d="M24 12.337c0-1.898-.491-3.68-1.351-5.229c3.666-8.298-3.929-7.083-4.352-7c-1.609.315-3.097.82-4.47 1.461a10.868 10.868 0 0 0-.612-.017C8.09 1.552 3.8 5.126 2.702 9.918c2.702-3.03 4.592-4.253 5.724-4.742c-.18.161-.358.324-.532.489l-.17.165c-.115.11-.23.22-.342.332l-.196.2c-.1.101-.199.202-.295.304c-.07.072-.136.144-.204.216a26.855 26.855 0 0 0-1.15 1.31a31.222 31.222 0 0 0-.41.505l-.2.253l-.183.24l-.206.272l-.14.193a33.453 33.453 0 0 0-1.168 1.7l-.002.003c-.093.145-.182.287-.27.428l-.014.023c-.088.141-.173.28-.255.418l-.009.014c-.222.37-.427.727-.613 1.063c-.971 1.76-1.444 2.99-1.464 3.063c-3.068 10.966 6.505 6.335 7.841 5.644a10.74 10.74 0 0 0 4.77 1.11c4.69 0 8.68-2.993 10.165-7.173h-5.666c-.839 1.416-2.453 2.376-4.308 2.376c-2.717 0-4.92-2.059-4.92-4.598h15.426c.059-.455.089-.919.089-1.39ZM21.985 1.724c.929.627 1.674 1.61.394 4.926a10.82 10.82 0 0 0-5.267-4.372c.998-.482 3.47-1.501 4.873-.554ZM2.248 21.989c-.757-.776-.89-2.665.779-6.108a10.816 10.816 0 0 0 4.696 5.739c-1.08.595-3.95 1.934-5.475.369ZM8.46 10.776c.086-2.468 2.235-4.444 4.874-4.444c2.64 0 4.787 1.976 4.874 4.444H8.46Z"/>`
	italicInnerSVG                 = `<path fill="currentColor" d="m13 19.56l-.13.43H6.3l.16-.43a4.05 4.05 0 0 0 1.3-.17a1.6 1.6 0 0 0 .76-.55a7.22 7.22 0 0 0 .8-2l2.77-9.61a7.07 7.07 0 0 0 .35-1.81a.86.86 0 0 0-.15-.52a.94.94 0 0 0-.46-.32a4.28 4.28 0 0 0-1.22-.11l.14-.43h6.16l-.13.43a2.6 2.6 0 0 0-1.12.17a1.78 1.78 0 0 0-.81.67a9.08 9.08 0 0 0-.71 1.93l-2.74 9.63a8.76 8.76 0 0 0-.4 1.69a.83.83 0 0 0 .15.5a.92.92 0 0 0 .47.32a6.35 6.35 0 0 0 1.38.18Z"/>`
	iterationInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 9v14h14M5 5v14h14M9 15h14V1H9v14Z"/>`
	javaInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 18V9h16v9c0 4-2 5-8 5s-8-1-8-5Zm16-9v3a3 3 0 1 0 3-3h-3Zm-4-3V2M5 6V4m4 2V0"/>`
	jsInnerSVG                     = `<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="M0 0h24v24H0z"/><path stroke="currentColor" stroke-width="2" d="M12 11v8c0 .876-.523 2-2 2c-2.385 0-2.5-2-2.5-2m13.29-5.484c-.6-1.01-1.396-1.516-2.386-1.516C16.856 12 16 13 16 14s.5 2 2.508 2.5c1.278.318 2.492 1 2.492 2.5s-1.315 2-2.5 2c-1.514 0-2.514-.667-3-2"/></g>`
	keyInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M10 13v3h3v3h3v2l2 2h5v-4L12.74 8.74C12.91 8.19 13 7.6 13 7c0-3.31-2.69-6-6-6S1 3.69 1 7a6.005 6.005 0 0 0 8.47 5.47L10 13ZM6 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`
	keyboardInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 9h3m-3 3h2m-2 3h1m3 0h10m1 0h1m1 0h1m-3-3h3m-2-3h2m-5 0h2M7 9h2m1 0h2m1 0h2M5 15h1m0-3h2m1 0h2m1 0h2m1 0h2M1 7v10a1 1 0 0 0 1 1h20a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1z"/>`
	languageInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm0 0c3 0 4-5 4-11S15 1 12 1S8 6 8 12s1 11 4 11ZM2 16h20M2 8h20"/>`
	lastfmInnerSVG                 = `<path fill="currentColor" d="M20.297 10.924c-.21-.069-.413-.132-.609-.194c-1.494-.47-2.394-.753-2.394-1.916c0-.943.727-1.627 1.73-1.627c.769 0 1.341.319 1.855 1.038c.048.067.138.09.213.05l1.507-.768a.162.162 0 0 0 .084-.101a.166.166 0 0 0-.014-.131c-.807-1.435-1.972-2.133-3.56-2.133c-2.417 0-3.98 1.462-3.98 3.725c0 2.314 1.511 3.25 4.298 4.168c1.616.538 2.33.824 2.33 1.973c0 1.291-1.165 2.22-2.755 2.166c-1.666-.056-2.17-.94-2.806-2.386a720.097 720.097 0 0 1-2.307-5.337C12.662 6.621 10.232 5 7.22 5C3.239 5 0 8.239 0 12.22c0 3.98 3.238 7.219 7.219 7.219c2.17 0 4.206-.962 5.582-2.641a.167.167 0 0 0 .025-.173l-.91-2.1a.171.171 0 0 0-.149-.1a.165.165 0 0 0-.154.09a4.946 4.946 0 0 1-4.395 2.66a4.96 4.96 0 0 1-4.954-4.955a4.96 4.96 0 0 1 4.954-4.955c1.989 0 3.81 1.18 4.534 2.941l2.251 5.134l.26.577c1.017 2.37 2.512 3.432 4.854 3.44c2.784 0 4.883-1.844 4.883-4.29c0-2.457-1.358-3.378-3.703-4.143Z"/>`
	launchInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 9v5s-3 2-3 5v1h4l2 3h4l2-3h4v-1c0-3-3-5-3-5V9c0-4-3-8-5-8S7 5 7 9Zm1 11h8M12 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	layerInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h16v16H1V1Zm19 6h3v16H7v-3"/>`
	licenseInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M10 13v3h3v3h3v2l2 2h5v-4L12.74 8.74C12.91 8.19 13 7.6 13 7c0-3.31-2.69-6-6-6S1 3.69 1 7a6.005 6.005 0 0 0 8.47 5.47L10 13ZM6 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`
	likeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 23h19c2 0 3-1 3-3V10h-7V4c0-2-1-3-3-3h-2s-.016 6-.016 7.326C10.984 9.652 10 11 8 11H1v12Zm5 0V11"/>`
	lineChartInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="m1 16l7-7l5 5L23 4M0 22h23.999M16 4h7v7"/>`
	linkInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16.125 2.42a2.009 2.009 0 0 1 2.84-.001l2.616 2.617a2 2 0 0 1-.001 2.839l-3.705 3.705a2.009 2.009 0 0 1-2.84.001L12.42 8.964a2 2 0 0 1 .001-2.839l3.705-3.705Zm-10 10a2.009 2.009 0 0 1 2.84-.001l2.616 2.617a2 2 0 0 1-.001 2.839L7.875 21.58a2.009 2.009 0 0 1-2.84.001L2.42 18.964a2 2 0 0 1 .001-2.839l3.705-3.705ZM7 17L17 7"/>`
	linkBottomInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 19V1M4 11l8 8l8-8M2 22h20"/>`
	linkDownInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22V2M3 13l9 9l9-9"/>`
	linkNextInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 12h20m-9-9l9 9l-9 9"/>`
	linkPreviousInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M22 12H2m9-9l-9 9l9 9"/>`
	linkTopInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 5v18M4 13l8-8l8 8M2 2h20"/>`
	linkUpInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 2v20M3 11l9-9l9 9"/>`
	linkedinInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M20.452 20.45h-3.56v-5.57c0-1.328-.022-3.036-1.85-3.036c-1.851 0-2.134 1.447-2.134 2.942v5.664H9.352V8.997h3.413v1.566h.049c.475-.9 1.636-1.85 3.367-1.85c3.605 0 4.27 2.371 4.27 5.456v6.281ZM5.339 7.433a2.063 2.063 0 1 1 0-4.13a2.065 2.065 0 0 1 0 4.13ZM7.12 20.45H3.558V8.997H7.12V20.45ZM23 0H1a1 1 0 0 0-1 1v22a1 1 0 0 0 1 1h22a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1Z"/>`
	linkedinOptionInnerSVG         = `<path fill="currentColor" fill-rule="evenodd" d="M22.037 22h-4.152v-6.496c0-1.55-.026-3.542-2.157-3.542c-2.16 0-2.49 1.688-2.49 3.43V22H9.09V8.64h3.98v1.827h.058c.553-1.05 1.908-2.158 3.928-2.158c4.204 0 4.98 2.766 4.98 6.364V22ZM4.409 6.816A2.407 2.407 0 0 1 2 4.407a2.408 2.408 0 1 1 2.41 2.408ZM6.486 22H2.33V8.64h4.156V22Z"/>`
	listInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M9 6h12M9 12h12M9 18h8M4 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	localInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 4h22v16H1V4Zm10 4h12M1 16h22M1 12h22M11 4v8"/>`
	locationInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22s-8-6-8-12c0-5 4-8 8-8s8 3 8 8c0 6-8 12-8 12Zm0-9a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`
	locationPinInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 10a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 0v12"/>`
	lockInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 23V11H5v12h14Zm-7-8v4m5-8V7c0-3 0-6-5-6S7 4 7 7v4"/>`
	loginInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 15v7h13V2H9v7m9 3H0m13-5l5 5l-5 5"/>`
	logoutInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M13 9V2H1v20h12v-7m9-3H5m12-5l5 5l-5 5"/>`
	loungeInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 5.997C5 5.447 5.45 5 6.007 5h11.986C18.55 5 19 5.453 19 5.997V13H5V5.997ZM22 8v7.003c0 .55-.455.997-.992.997H2.992A.999.999 0 0 1 2 15.003V8m3 8v2v-2Zm14 0v2v-2Z"/>`
	magicInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="m2.5 19.5l17-17l2 2l-17 17l-2-2Zm.5.5L15 8l1 1L4 21l-1-1ZM5.5 3l-.5.5l.5.5l.5-.5l-.5-.5Zm6 0l-.5.5l.5.5l.5-.5l-.5-.5Zm-3 3l-.5.5l.5.5l.5-.5l-.5-.5Zm12 6l-.5.5l.5.5l.5-.5l-.5-.5Zm0 5l-.5.5l.5.5l.5-.5l-.5-.5Z"/>`
	mailInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M23 20V6l-11 9L1 6v14h22Zm-11-8l10-8H2l10 8Z"/>`
	mailOptionInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 4h22v16H1V4Zm0 1l11 8.5L23 5"/>`
	mandrivaInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M14.557 17.2c-3.021 2.72-6.354 4.146-9.075 4.148h-.008c-1.938-.001-3.565-.726-4.544-2.222C-1.43 15.519.858 8.797 6.042 4.11c.44-.398.888-.768 1.339-1.111c-4.343 4.582-6.112 10.546-3.938 13.868c1.952 2.983 6.474 2.899 10.912.13l-1.323-1.332c-.157-.157-.428-.258-.692-.258a.826.826 0 0 0-.292.05l-4.279 1.656a.774.774 0 0 1-.275.06c-.176 0-.259-.093-.292-.148c-.052-.085-.087-.24.042-.495l2.094-4.082c.148-.29.1-.748-.105-1L6.336 7.892c-.195-.239-.18-.406-.134-.504c.033-.07.126-.19.372-.19c.048 0 .1.006.155.015l4.528.729a.68.68 0 0 0 .11.008c.307 0 .655-.18.809-.417l2.487-3.853c.152-.237.302-.287.4-.287c.09 0 .306.044.368.449l.706 4.533c.05.322.358.664.672.747l4.434 1.175c.332.088.399.259.408.352c.01.094-.019.274-.325.43l-4.094 2.07c-.29.147-.52.545-.502.87l.253 4.582c.01.193-.027.337-.113.428a.292.292 0 0 1-.217.093c-.123 0-.253-.068-.387-.203L14.557 17.2ZM24 11.403c-1.245-.672-1.423-.64-2.355.424c.671-1.244.639-1.423-.425-2.356c1.245.672 1.424.64 2.356-.424c-.672 1.245-.64 1.424.424 2.356Z"/>`
	manualInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M14 9v8v-8Zm-4 0v8v-8ZM8 5a4 4 0 1 0 8 0a4 4 0 0 0-8 0ZM4 23h16v-3H4v3Zm3-3h10v-3H7v3Z"/>`
	mapInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M15 15h4l3 7H2l3-7h4m4-7a1 1 0 1 1-2 0a1 1 0 0 1 2 0M6 8c0 5 6 10 6 10s6-5 6-10c0-3.417-2.686-6-6-6S6 4.583 6 8Z"/>`
	mapLocationInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17.5 6.5L23 9v13l-7-3l-8 3l-7-3V6l5 2m10 11v-7M8 22V12m4 4.273S6 11.5 6 7c0-3.75 3-6 6-6s6 2.25 6 6c0 4.5-6 9.273-6 9.273ZM13 7a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/>`
	mastercardInnerSVG             = `<g fill="currentColor" fill-rule="evenodd"><circle cx="7" cy="12" r="7"/><circle cx="17" cy="12" r="7" fill-opacity=".8"/></g>`
	mediumInnerSVG                 = `<path fill="currentColor" d="M2.846 5.887a.93.93 0 0 0-.303-.784l-2.24-2.7V2H7.26l5.378 11.795L17.367 2H24v.403L22.084 4.24a.56.56 0 0 0-.213.538v13.498a.56.56 0 0 0 .213.537l1.871 1.837v.403h-9.412v-.403l1.939-1.882c.19-.19.19-.246.19-.537V7.32l-5.39 13.688h-.727L4.28 7.32v9.174c-.052.385.076.774.347 1.052l2.521 3.058v.404H0v-.404l2.521-3.058c.27-.279.39-.67.325-1.052V5.887Z"/>`
	memoryInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M10 18h4m-4-4h4m-4-4h4m-4-4h4m6 12h3m-3-4h3m-3-4h3m-3-4h3M1 18h3m-3-4h3m-3-4h3M1 6h3m11 14h4a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1zM5 20h4a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1z"/>`
	menuInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 19h20M2 5h20M2 12h20"/>`
	microfocusInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M1 5h4v14h14v4H1V5Zm4-4h18v18h-4V5H5V1Z"/>`
	microphoneInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 11c0 3 1.79 4 4 4s4-1 4-4V5c0-3-1.79-4-4-4S8 2 8 5v6ZM4 9v2c0 5 3.582 8 8 8s8-3 8-8V9m-8 15v-5"/>`
	moneyInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 16c0-1.105-3.134-2-7-2s-7 .895-7 2s3.134 2 7 2s7-.895 7-2ZM2 16v4.937C2 22.077 5.134 23 9 23s7-.924 7-2.063V16M9 5c-4.418 0-8 .895-8 2s3.582 2 8 2M1 7v5c0 1.013 3.582 2 8 2M23 4c0-1.105-3.1-2-6.923-2c-3.824 0-6.923.895-6.923 2s3.1 2 6.923 2S23 5.105 23 4Zm-7 12c3.824 0 7-.987 7-2V4M9.154 4v10.166M9 9c0 1.013 3.253 2 7.077 2C19.9 11 23 10.013 23 9"/>`
	monitorInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 16h22V2H1v14Zm4 6h14H5Zm4 0h6v-6H9v6Z"/>`
	monospaceInnerSVG              = `<path fill="currentColor" d="M11.88 17h-2.3l-.93-2.9H4.44L3.57 17H1.32L5.41 4.17h2.25ZM8 11.93L6.52 7.17l-1.43 4.76ZM14.13 17L12.22 4.17h1.66L15.07 13l1.46-8.82h1.92l1.4 9l1.23-9h1.62L20.78 17h-1.72l-1.6-9.6l-1.58 9.6Zm-2.77.95v1.39H1.89v-1.39h-.51v1.91h10.49v-1.91h-.51zm10.81.05v1.39h-9.48V18h-.5v1.91h10.49V18h-.51z"/>`
	moonInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9.874 5.008c2.728-1.68 6.604-1.014 8.25.197c-2.955.84-5.11 3.267-5.242 6.415c-.18 4.28 3.006 6.588 5.24 7.152c-1.964 1.343-4.36 1.293-5.235 1.172c-3.568-.492-6.902-3.433-7.007-7.711c-.106-4.278 2.573-6.35 3.994-7.225z"/>`
	moreInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 13v-2h2v2H3Zm8 0v-2h2v2h-2Zm8 0v-2h2v2h-2Z"/>`
	moreVerticalInnerSVG           = `<path fill="currentColor" d="M14 14h-4v-4h4v4Zm0-7h-4V3h4v4Zm0 14h-4v-4h4v4Z"/>`
	mouseInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 4a5 5 0 0 1 5 5v6a5 5 0 0 1-10 0V9a5 5 0 0 1 5-5zm0 0v6m-6 0h12"/>`
	multimediaInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8v14h20V9M11 4L7 8M2 4v4h15l4-4H2Zm14 0l-4 4"/>`
	multipleInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 15h4V1H9v4m6 14h4V5H5v4M1 23h14V9H1v14Z"/>`
	musicInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 17.998C1 16.894 1.887 16 2.998 16H9v4.002A1.993 1.993 0 0 1 7.002 22H2.998A2 2 0 0 1 1 20.002v-2.004Zm14 0c0-1.104.887-1.998 1.998-1.998H23v4.002A1.993 1.993 0 0 1 21.002 22h-4.004A2 2 0 0 1 15 20.002v-2.004ZM9 16V2h14v13.5M9 6h14"/>`
	mysqlInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M5.462 4.04a2.65 2.65 0 0 0-.67.074v.038h.037c.13.267.36.44.521.67l.372.781l.038-.037c.23-.162.336-.422.335-.819c-.093-.097-.107-.219-.187-.335c-.106-.154-.312-.242-.446-.372Zm18.017 19.097c.175.129.293.329.521.41v-.038c-.12-.152-.15-.362-.26-.521a73.62 73.62 0 0 1-.484-.484a7.948 7.948 0 0 0-1.713-1.638c-.508-.365-1.649-.859-1.861-1.451l-.038-.038c.361-.04.784-.171 1.117-.26c.56-.15 1.06-.112 1.638-.26c.261-.076.521-.15.782-.224v-.15c-.292-.3-.5-.696-.819-.967c-.834-.71-1.743-1.42-2.68-2.01c-.52-.329-1.162-.541-1.713-.82c-.185-.093-.51-.142-.632-.297c-.29-.37-.447-.837-.67-1.266c-.467-.9-.927-1.883-1.34-2.83c-.283-.645-.467-1.281-.82-1.86c-1.69-2.78-3.51-4.457-6.328-6.106c-.6-.35-1.322-.489-2.084-.67l-1.229-.074c-.25-.105-.51-.41-.744-.559C3.188.434.792-.849.102.838c-.437 1.065.652 2.104 1.042 2.643c.273.379.623.803.819 1.229c.128.28.15.56.26.856c.271.73.506 1.522.856 2.196c.178.341.373.7.596 1.005c.138.187.372.27.409.559c-.23.321-.242.82-.371 1.228c-.582 1.835-.363 4.115.484 5.473c.259.416.87 1.31 1.711.967c.736-.3.572-1.228.782-2.047c.047-.186.019-.323.112-.447v.037l.67 1.34c.496.799 1.376 1.634 2.122 2.197c.386.292.69.797 1.191.968v-.038h-.037c-.098-.15-.25-.213-.372-.335a8.554 8.554 0 0 1-.857-.968c-.678-.92-1.277-1.928-1.823-2.977c-.261-.502-.488-1.054-.708-1.564c-.085-.197-.084-.494-.26-.596c-.241.374-.596.676-.782 1.117c-.298.705-.337 1.565-.447 2.457c-.065.023-.036.007-.075.037c-.518-.125-.7-.659-.893-1.117c-.487-1.157-.578-3.022-.149-4.355c.111-.345.613-1.431.41-1.75c-.098-.318-.417-.501-.596-.744A5.83 5.83 0 0 1 3.6 7.166c-.398-.902-.585-1.916-1.005-2.829c-.2-.436-.54-.877-.819-1.265c-.308-.43-.654-.746-.893-1.266c-.085-.185-.201-.48-.075-.67c.04-.128.097-.182.224-.223c.216-.167.817.055 1.042.148c.597.248 1.095.484 1.6.82c.243.16.489.472.782.558h.335c.525.12 1.112.037 1.601.186c.865.263 1.64.672 2.345 1.117c2.146 1.355 3.9 3.283 5.1 5.584c.193.37.277.724.447 1.117c.343.792.775 1.607 1.116 2.382c.34.773.673 1.553 1.154 2.196c.253.338 1.231.52 1.676.708c.311.131.821.269 1.116.446c.564.34 1.11.745 1.638 1.117c.264.187 1.077.595 1.117.93c-1.31-.034-2.31.087-3.164.448c-.243.102-.63.105-.67.409c.133.14.154.35.26.521c.204.33.549.773.856 1.005c.337.254.683.525 1.043.745c.64.39 1.356.614 1.972 1.005c.365.231.726.521 1.08.782Z"/>`
	navigateInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="m20 11l2-3l-2-3h-8v6h8Zm-8 13V0M4 2L2 5l2 3h8V2H4Z"/>`
	networkInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7.5 7v8.514c0 4.243 5.5 2.829 5.5 6.6V24m-2-14L7.5 7L4 10m12.5-8v8.44c0 4.068-3.5 2.712-3.5 6.328V24m0-19l3.5-3L20 5"/>`
	networkDriveInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 14v4M22 6v6a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1zM12 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4zM6 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`
	newInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 1v22M2 6l20 12m0-12L2 18"/>`
	newWindowInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M11 9h8m-4 4V5m2 12v6H1V7h6m0-6h16v16H7V1Z"/>`
	nextInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="m7 2l10 10L7 22"/>`
	nodeInnerSVG                   = `<path fill="currentColor" d="M11.899 24c-.322 0-.64-.084-.923-.247l-2.935-1.738c-.439-.245-.225-.332-.08-.382c.584-.204.703-.25 1.327-.605c.066-.037.152-.024.219.015l2.255 1.339a.298.298 0 0 0 .273 0l8.794-5.077a.277.277 0 0 0 .134-.237V6.919a.282.282 0 0 0-.136-.242l-8.79-5.072a.27.27 0 0 0-.271 0l-8.79 5.072a.28.28 0 0 0-.139.24v10.148c0 .097.053.19.137.236l2.408 1.391c1.308.654 2.107-.116 2.107-.891V7.785a.25.25 0 0 1 .255-.254h1.114c.139 0 .253.11.253.254v10.02c0 1.744-.95 2.745-2.604 2.745c-.509 0-.91 0-2.028-.55l-2.307-1.33a1.86 1.86 0 0 1-.922-1.605V6.917c0-.66.352-1.277.922-1.602L10.976.236a1.928 1.928 0 0 1 1.849 0l8.792 5.08c.568.329.922.943.922 1.603v10.149a1.86 1.86 0 0 1-.922 1.602l-8.792 5.079a1.848 1.848 0 0 1-.927.246V24Zm2.716-6.993c-3.848 0-4.654-1.766-4.654-3.248c0-.14.113-.253.254-.253h1.136c.126 0 .231.091.251.215c.172 1.158.683 1.742 3.01 1.742c1.853 0 2.641-.419 2.641-1.402c0-.566-.225-.986-3.104-1.268c-2.408-.238-3.896-.768-3.896-2.694c0-1.775 1.497-2.831 4.004-2.831c2.815 0 4.211.977 4.387 3.077a.256.256 0 0 1-.255.278h-1.143a.252.252 0 0 1-.246-.199c-.275-1.217-.94-1.607-2.747-1.607c-2.023 0-2.259.705-2.259 1.233c0 .64.277.828 3.007 1.189c2.703.359 3.987.865 3.987 2.765c0 1.915-1.599 3.014-4.385 3.014l.012-.01Z"/>`
	nodesInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M14 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM7 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm14-8a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`
	nortonInnerSVG                 = `<path fill="currentColor" d="M21.91.816h.797v.797h-.797V.816Zm0 .797h-1.477v.797h1.477v-.797Zm-1.477.797v1.554h-.738V2.41h.738Zm0 1.554h.758v.758h-.758v-.758Zm.758 0v-.748h.836v.748h-.836Zm-.758.758v1.613h-.65v.632h-.69v.923H18.5v.787h-.758V7.9h-.758l-.01.777h.768v.826h-.758l.01.903h-.651l.01.758h-.603l-.01.855c-1.136 1.635-1.923 3.351-2.633 5.063c-.09.217-.123.395-.233.524c-.239.28-.595.368-.885.36c-.542-.015-1.02-.268-1.272-.855c-.85-1.97-1.27-3.35-3.605-5.937c-.376-.417-.289-.949.01-1.117c.286-.163.687-.046 1.136.204c1.258.7 1.958 1.415 3.45 3.4c.543-1.311 1.482-3.084 2.652-5.004h.544V7.9h.515v-.75h.758V5.956h.817l-.01.758h.758v-.758h-.748v-.68h.68v-.554h1.253V5.5h.768v-.778h.738Zm-1.506 0v-.758h.768v.758h-.768Zm0-.758h-.767v-.728h.767v.728Zm1.506-1.55h.757v.798h-.757v-.798Zm2.27-.799h.819v.799h-.82v-.799Zm-6.156 2.929v.742h-.933V6.53h-.933v.734a7.115 7.115 0 0 0-3.313-.812c-3.947 0-7.152 3.195-7.152 7.142c0 3.946 3.205 7.141 7.152 7.141a7.139 7.139 0 0 0 7.141-7.141a7.115 7.115 0 0 0-.823-3.333h.726v-.933h.622v-.933h.622v-.777h.27a10.357 10.357 0 0 1 1.887 5.976C21.813 19.34 17.153 24 11.406 24C5.66 24 1 19.34 1 13.594C1 7.847 5.66 3.187 11.406 3.187c1.87 0 3.624.493 5.14 1.357ZM22.703 0h.819v.817h-.82V0Z"/>`
	noteInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 23h15l7-7V1H1v22Zm14 0v-8h8"/>`
	notesInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 1v22h13l5-5V1H3Zm3 16h5m-5-4h12M6 9h10M3 5h18m0 12h-6v6"/>`
	notificationInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 19V9a8 8 0 0 1 16 0v10M1 19h22m-8 0v1a3 3 0 1 1-6 0v-1"/>`
	npmInnerSVG                    = `<g fill="currentColor" fill-rule="evenodd"><path d="M0 0h24v24H0z"/><path d="M16.718 7.928h-4.513V20.25H4V3h16v17.249h-3.282V7.93Z"/></g>`
	objectGroupInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h3v3H1V1Zm19 0h3v3h-3V1ZM4 2h16M4 22h16M1 20h3v3H1v-3Zm19 0h3v3h-3v-3ZM2 4v16M22 4v16M7 7h7v6H7V7Zm10 3v7h-7v-2"/>`
	objectUngroupInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h3v3H1V1Zm12 0h3v3h-3V1ZM4 2h9m2 7h5M4 15h9M1 13h3v3H1v-3Zm12 0h3v3h-3v-3ZM2 4v9m13-9v9m5-5h3v3h-3V8Zm-9 14h9M8 20h3v3H8v-3Zm12 0h3v3h-3v-3ZM9 16v4m13-9v9"/>`
	offlineStorageInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 21a9 9 0 1 0 0-18a9 9 0 1 0 0 18Zm8-12h-8a3 3 0 0 0 0 6h8"/>`
	onedriveInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M21.692 13.987a2.733 2.733 0 0 1 .602 5.232h-13.7a3.166 3.166 0 1 1 .281-6.321a3.963 3.963 0 0 1 7.482-2.05a3.454 3.454 0 0 1 5.336 3.138ZM8.187 12.209a3.873 3.873 0 0 0-3.44 3.843c0 .81.252 1.563.681 2.186h-2.15a3.279 3.279 0 0 1-.237-6.549a3.692 3.692 0 0 1 5.668-3.86a5.103 5.103 0 0 1 9.648 1.757c-.036-.001-.072-.003-.109-.003c-.568 0-1.125.115-1.64.337a4.644 4.644 0 0 0-3.778-1.929a4.67 4.67 0 0 0-4.643 4.218Z"/>`
	operaInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M12.125 0C5.568 0 1 4.756 1 11.889C1 18.236 5.438 24 12.125 24c6.752 0 11.226-5.763 11.226-12.111C23.35 4.699 18.62 0 12.125 0Zm0 21.32a3.308 3.308 0 0 1-1.425-.298c-1.141-.575-1.828-1.85-2.23-3.41c-.435-1.809-.484-4.004-.484-5.926c0-3.431.25-6.532 1.65-8.08c.612-.65 1.418-1.034 2.476-1.037h.013c1.379 0 2.345.675 3.016 1.734c1 1.688 1.227 4.424 1.227 7.368c0 4.18-.278 9.65-4.243 9.65Z"/>`
	optimizeInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 22h4v-4H2v4ZM22 2L12 12m10-2V2h-8m8 11h-4v9h4v-9Zm-12 9h4v-6h-4v6Z"/>`
	oracleInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M7.957 18.912A6.953 6.953 0 0 1 1 11.962A6.963 6.963 0 0 1 7.957 5h8.087A6.961 6.961 0 0 1 23 11.962a6.952 6.952 0 0 1-6.956 6.95H7.957Zm7.907-2.453a4.497 4.497 0 0 0 4.503-4.497a4.507 4.507 0 0 0-4.503-4.508H8.136a4.507 4.507 0 0 0-4.503 4.508a4.498 4.498 0 0 0 4.503 4.497h7.728Z"/>`
	orderedListInnerSVG            = `<path fill="currentColor" d="M5.77 6.42h18.06v1.75H5.77zm0 5.29h18.06v1.75H5.77zm0 5.28h18.06v1.75H5.77zM3.13 4.87V8a2 2 0 0 0 0 .45a.3.3 0 0 0 .13.16a.62.62 0 0 0 .32.06h.12v.11h-2v-.07h.1a.79.79 0 0 0 .35-.06a.29.29 0 0 0 .14-.16A1.75 1.75 0 0 0 2.3 8V6a1.28 1.28 0 0 0 0-.33a.24.24 0 0 0-.1-.11a.28.28 0 0 0-.16 0a.91.91 0 0 0-.35.09l-.05-.1L3 4.87Zm.52 9.43H1.37v-.06a12.51 12.51 0 0 0 1.27-1.67a1.81 1.81 0 0 0 .22-.84a.7.7 0 0 0-.18-.5a.6.6 0 0 0-.45-.2a.75.75 0 0 0-.68.44h-.11a1.58 1.58 0 0 1 .47-.81a1.09 1.09 0 0 1 .72-.26a1.06 1.06 0 0 1 .54.14a1 1 0 0 1 .38.37a.9.9 0 0 1 .14.45a1.6 1.6 0 0 1-.21.77a7.28 7.28 0 0 1-1.25 1.47h.83a1.87 1.87 0 0 0 .4 0a.33.33 0 0 0 .15-.09a1.16 1.16 0 0 0 .16-.26h.1ZM2.09 18v-.1a1.88 1.88 0 0 0 .45-.17a.67.67 0 0 0 .22-.25a.69.69 0 0 0 .09-.34a.55.55 0 0 0-.17-.41a.58.58 0 0 0-.43-.17a.8.8 0 0 0-.68.42h-.11a1.72 1.72 0 0 1 .54-.79a1.16 1.16 0 0 1 .71-.24a.93.93 0 0 1 .66.24a.76.76 0 0 1 .26.57a.78.78 0 0 1-.12.41a1 1 0 0 1-.38.35a1.34 1.34 0 0 1 .51.4a1 1 0 0 1 .17.6a1.38 1.38 0 0 1-.44 1A1.62 1.62 0 0 1 2.2 20a1.24 1.24 0 0 1-.71-.16a.34.34 0 0 1-.16-.29a.32.32 0 0 1 .31-.32a.41.41 0 0 1 .18 0l.32.25a.84.84 0 0 0 .52.23a.47.47 0 0 0 .34-.16a.61.61 0 0 0 .15-.42a1.1 1.1 0 0 0-.27-.72a1.42 1.42 0 0 0-.79-.41Z"/>`
	organizationInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20 3v20H4V3h16ZM8.042 9h2V7h-2v2ZM14 9h2V7h-2v2Zm-5.958 6h2v-2h-2v2Zm2 8h4v-4h-4v4ZM14 15h2v-2h-2v2ZM2 3h20V1H2v2Z"/>`
	overviewInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18.5 21a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9ZM10 7h4M1.5 14.5S5.5 5 6 4s1.5-1 2-1s2 0 2 2v11m-4.5 5a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm17-6.5S18.5 5 18 4s-1.5-1-2-1s-2 0-2 2v11m-4 0h4"/>`
	packageInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M12.371.571L12 .423l-.371.148l-10 4L1 4.823v14.354l.629.251l10 4l.371.149l.371-.149l10-4l.629-.251V4.823l-.629-.252l-10-4ZM3 6.977v10.846l8 3.2V10.177l-8-3.2Zm10 3.2v10.846l8-3.2V6.977l-8 3.2ZM19.307 5.5L12 2.577L9.943 3.4l7.307 2.923l2.057-.823Zm-14.614 0L7.25 4.477L14.557 7.4L12 8.423L4.693 5.5Z" clip-rule="evenodd"/>`
	paintInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 10H2c0-5.523 0-9 10-9s10 3.477 10 9h-2m-8 0c4.418 0 8-.895 8-2s-3.582-2-8-2s-8 .895-8 2s3.582 2 8 2ZM4 20c0 1.657 3.582 3 8 3s8-1.343 8-3m0-12v12V8ZM4 20V8v12Zm4-7v4m5-4v4m-6-5c-1.5 0-3-1-3-4m14 5.5V12c0-2 2-1 2-4M8 13a1 1 0 0 0-1-1m6 1a1 1 0 0 1 2 0v.5m0 0a1.5 1.5 0 0 0 3 0M8 17a2.5 2.5 0 1 0 5 0"/>`
	panInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8.5 5.5L12 2l3.5 3.5M22 12H2m3.5-3.5L2 12l3.5 3.5m13 0L22 12l-3.5-3.5M12 22V2M8.5 18.5L12 22l3.5-3.5"/>`
	pauseInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 21h6V3H3v18Zm12 0h6V3h-6v18Z"/>`
	pauseFillInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 21h6V3H3v18Zm1-2h4V5H4v14Zm1-2h2V7H5v10Zm10 4h6V3h-6v18Zm1-2h4V5h-4v14Zm1-2h2V7h-2v10Z"/>`
	paypalInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M21.495 7.054c-1.07 4.973-4.48 7.604-9.89 7.604H9.643L8.18 24h3.182c.46 0 .85-.334.923-.788l.037-.198l.732-4.636l.047-.256a.933.933 0 0 1 .922-.788h.581c3.76 0 6.705-1.528 7.565-5.946c.345-1.773.179-3.26-.674-4.334M19.317 1.81C18.206.543 16.197 0 13.627 0H6.169c-.526 0-.973.383-1.055.9L2.008 20.598a.64.64 0 0 0 .633.74h4.604l1.157-7.335l-.036.23c.082-.518.526-.9 1.051-.9h2.188c4.299 0 7.664-1.747 8.648-6.797c.029-.15.076-.438.076-.438c.279-1.869-.002-3.137-1.012-4.287"/>`
	performanceInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 19l-2 3l-3-1l-.5-3.5L3 17l-1-3l3-2l-3-2l1-3l3.5-.5L7 3l3-1l2 3l2-3l3 1l.5 3.5L21 7l1 3l-3 2l3 2l-1 3l-3.5.5L17 21l-3 1l-2-3Zm0-3a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`
	personalComputerInnerSVG       = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 18h18V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v13Zm-1 2h20c1 0 1-1 1-1H1s0 1 1 1Z"/>`
	phoneInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6.375 2C5 2 3 3.5 2.5 4.5c-.715 1.43-.597 1.99-.125 3.5c.625 2 2.457 5.545 5 8c3.625 3.5 7 5 8.5 5.5s3.125 0 4.125-1s2-2 .875-3.5c-.797-1.063-1.959-2.292-3.375-3c-1.288-.644-2.056-.41-2.5.5c-.246.503-.322 1.466-.5 2c-.225.674-1.125.5-2.125 0C11.418 16.021 9 14 7 11c-1.24-1.859.742-1.87 2-2.5c1-.5 1.31-1.65.5-3C8 3 7.5 2 6.375 2Z"/>`
	phoneFlipInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M8 1a1 1 0 0 0-.928.629l-4 10a1 1 0 0 0 0 .742l4 10A1 1 0 0 0 8 23h10a1 1 0 0 0 .928-1.371L15.078 12l3.851-9.629A1 1 0 0 0 18 1H8ZM5.477 11l3.2-8h2.794a1 1 0 0 0 1 1h.057a1 1 0 0 0 1.001-1h2.994l-3.2 8H5.477Zm0 2l3.2 8h7.846l-3.2-8H5.477Zm5.558 6.773a1 1 0 0 1 .937-1.351h.056a1 1 0 1 1 0 2h-.056a1 1 0 0 1-.937-.649Z" clip-rule="evenodd"/>`
	phoneHorizontalInnerSVG        = `<path fill="currentColor" fill-rule="evenodd" d="M23 7a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V7Zm-2 5.996V17H3V7h18v4.004a1.038 1.038 0 0 0-.094-.004h-1a1 1 0 0 0 0 2h1c.032 0 .063-.002.094-.004Z" clip-rule="evenodd"/>`
	phoneVerticalInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M17 23a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h10Zm-6-2H7V3h10v18h-4v-1.031a1 1 0 0 0-2 0V21Z" clip-rule="evenodd"/>`
	pieChartInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M10 23a9 9 0 0 1 0-18v9l1.162 1.162l5.202 5.202A8.972 8.972 0 0 1 10 23Zm4-13V1a9 9 0 0 1 9 9h-9Zm0 3h8a8.964 8.964 0 0 1-2.107 5.787L14 13Z"/>`
	piedPiperInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M0 19.421c2.274 0 4.042-.758 4.042-.758s3.032-7.579 7.326-7.579c3.285 0 3.79 2.527 3.79 2.527S19.958 4.263 24 3c-3.79 3.032-3.284 6.316-5.053 7.832c-1.768 1.515-1.768.006-3.79 3.543c-4.546.505-6.032 2.014-9.094 3.783c5.305-2.526 6.316-2.78 11.116-2.526c.504.026.758.252.505.757c-.733 1.466-1.28 3.673-2.273 3.537c-5.558-.758-8.843.506-11.622.506c-2.778 0-3.789-.506-3.789-1.01Z"/>`
	pinInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="m16 3l-6 6s-4-1-7 2l10 10c3-3 2-7 2-7l6-6l-5-5ZM1 23l7-7m6-15l9 9"/>`
	pinterestInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M12 0C5.372 0 0 5.372 0 12c0 5.084 3.163 9.426 7.627 11.174c-.105-.95-.2-2.406.042-3.442c.217-.936 1.407-5.965 1.407-5.965s-.36-.718-.36-1.781c0-1.669.968-2.915 2.172-2.915c1.024 0 1.518.769 1.518 1.69c0 1.03-.655 2.57-.993 3.996c-.283 1.195.598 2.169 1.777 2.169c2.133 0 3.772-2.25 3.772-5.495c0-2.873-2.065-4.883-5.013-4.883c-3.414 0-5.418 2.562-5.418 5.208c0 1.031.397 2.138.893 2.739a.359.359 0 0 1 .083.344c-.091.38-.293 1.194-.333 1.361c-.053.219-.174.266-.402.16c-1.499-.698-2.436-2.888-2.436-4.649c0-3.785 2.75-7.261 7.93-7.261c4.162 0 7.397 2.966 7.397 6.93c0 4.136-2.607 7.464-6.227 7.464c-1.216 0-2.358-.632-2.75-1.378l-.748 2.853c-.27 1.042-1.002 2.348-1.492 3.146A11.99 11.99 0 0 0 12 24c6.628 0 12-5.373 12-12c0-6.628-5.372-12-12-12"/>`
	planInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 4V0v4ZM7 18H5h2Zm12 0H9h10ZM7 14H5h2Zm12 0H9h10ZM6 4V0v4ZM1 9h22H1Zm0 14h22V4H1v19Z"/>`
	playInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="m3 22l18-10L3 2z"/>`
	playFillInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="m3 22l18-10L3 2v20Zm2-3l12.6-7L5 5v14Zm2-3l7.2-4L7 8v8Zm2-3l1.8-1L9 11v2Z"/>`
	plugInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M15 6V1m-3 23v-9M9 6V1M6 6h12v7a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V6z"/>`
	pocketInnerSVG                 = `<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="M12 2H2a2 2 0 0 0-2 2v8c0 5.982 6 11 12 11s12-5.018 12-11V4a2 2 0 0 0-2-2H12Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="m6 9l6.404 6L18 9"/></g>`
	powerInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 4c3.364 1.43 6 4.99 6 9c0 5.6-4.473 10-10 10S2 18.6 2 13c0-4.01 2.636-7.57 6-9m4-3v10"/>`
	powerCycleInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20 8a8.955 8.955 0 0 0-8.036-5C7.014 3 3 7.03 3 12m1 4a8.955 8.955 0 0 0 8.036 5C16.986 21 21 16.97 21 12M9 16H3v6M21 2v6h-6"/>`
	powerForceShutdownInnerSVG     = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm0-13v8"/>`
	powerResetInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20 8c-1.403-2.96-4.463-5-8-5a9 9 0 1 0 0 18a9 9 0 0 0 9-9m0-9v6h-6"/>`
	powerShutdownInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 1v8M6.994 4.52a9.044 9.044 0 0 0-1.358 1.116a9 9 0 1 0 11.37-1.117"/>`
	previousInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17 2L7 12l10 10"/>`
	printInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 19H1V7h22v12h-5M3 16h18M6 16v7h12v-7m0-9V1H6v6m11 5h2v-1h-2v1Z"/>`
	productHuntInnerSVG            = `<path fill="currentColor" fill-rule="evenodd" d="M13.6 8.4h-3.4V12h3.4a1.8 1.8 0 1 0 0-3.6m0 6h-3.4V18H7.8V6h5.8a4.2 4.2 0 1 1 0 8.4M12 0C5.372 0 0 5.372 0 12s5.372 12 12 12c6.627 0 12-5.372 12-12S18.627 0 12 0"/>`
	projectsInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 15v8H1v-8h8Zm14 0v8h-8v-8h8ZM9 1v8H1V1h8Zm14 0v8h-8V1h8Z"/>`
	qrInnerSVG                     = `<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="M13 14h1v1h-1v-1Zm1 1h1v1h-1v-1Zm0 1h1v1h-1v-1Zm2 0h1v1h-1v-1Zm0 1h1v1h-1v-1Zm-3-1h1v1h-1v-1Zm2 0h1v1h-1v-1Zm0 1h1v1h-1v-1Zm3-1h1v1h-1v-1Zm0-1h1v1h-1v-1Zm1-1h1v1h-1v-1Zm-2 2h1v1h-1v-1Zm0 1h1v1h-1v-1Zm-1 1h1v1h-1v-1Zm-1 0h1v1h-1v-1Zm2 0h1v1h-1v-1Zm1 0h1v1h-1v-1Zm-2 1h1v1h-1v-1Zm-2 0h1v1h-1v-1Zm1 0h1v1h-1v-1Zm-2 0h1v1h-1v-1Zm0 1h1v1h-1v-1Zm1 1h1v1h-1v-1Zm1 0h1v1h-1v-1Zm2 0h1v1h-1v-1Zm1 0h1v1h-1v-1Zm-1-2h1v1h-1v-1Zm1 0h1v1h-1v-1Zm1-1h1v1h-1v-1Zm0-1h1v1h-1v-1Zm0 3h1v1h-1v-1Zm0-1h1v1h-1v-1Zm1-1h1v1h-1v-1Zm0-1h1v1h-1v-1Zm1 3h1v1h-1v-1Zm0-2h1v1h-1v-1Zm0 1h1v1h-1v-1Zm-2-3h1v1h-1v-1Zm-6 1h1v1h-1v-1Zm-1 0h1v1h-1v-1Zm0 1h1v1h-1v-1Zm2 0h1v1h-1v-1Zm-3 0h1v1h-1v-1Zm2 0h1v1h-1v-1Zm-2 1h1v1h-1v-1Zm0 1h1v1h-1v-1Zm0-19h1v1h-1V1Zm1 1h1v1h-1V2Zm-1 2h1v1h-1V4Zm1 1h1v1h-1V5Zm-1 1h1v1h-1V6Zm1 0h1v1h-1V6Zm0 1h1v1h-1V7Zm0 1h1v1h-1V8Zm-1 1h1v1h-1V9Zm1 0h1v1h-1V9Zm-1 1h1v1h-1v-1ZM1 11h1v1H1v-1Zm1 1h1v1H2v-1Zm2-1h1v1H4v-1Zm0 1h1v1H4v-1Zm1-1h1v1H5v-1Zm1 1h1v1H6v-1Zm1-1h1v1H7v-1Zm1 1h1v1H8v-1Zm0-1h1v1H8v-1Zm1 0h1v1H9v-1Zm1 0h1v1h-1v-1Zm1 1h1v1h-1v-1Zm2 0h1v1h-1v-1Zm1-1h1v1h-1v-1Zm1 0h1v1h-1v-1Zm1 0h1v1h-1v-1Zm-1 2h1v1h-1v-1Zm-2 9h1v1h-1v-1Zm-1 0h1v1h-1v-1Zm0-9h1v1h-1v-1Zm-1 0h1v1h-1v-1Zm0 1h1v1h-1v-1Zm0 1h1v1h-1v-1Zm11-1h1v1h-1v-1Zm-1 1h1v1h-1v-1Zm1 2h1v1h-1v-1Zm-5-4h1v1h-1v-1Zm1-1h1v1h-1v-1Zm4 0h1v1h-1v-1Zm0 1h1v1h-1v-1Zm-1 0h1v1h-1v-1Zm1 8h1v1h-1v-1Zm-1 1h1v1h-1v-1Zm-2 0h1v1h-1v-1Zm3 0h1v1h-1v-1Z"/><path stroke="currentColor" stroke-width="2" d="M15 2h7v7h-7V2ZM2 2h7v7H2V2Zm0 13h7v7H2v-7ZM18 5h1v1h-1V5ZM5 5h1v1H5V5Zm0 13h1v1H5v-1Z"/></g>`
	radialInnerSVG                 = `<circle cx="12" cy="12" r="11" fill="none" stroke="currentColor" stroke-width="2"/>`
	radialSelectedInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm0-10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 2a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/>`
	raspberryInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M8.087 0a.667.667 0 0 0-.399.165c-.363-.14-.717-.19-1.032.096c-.486-.063-.643.067-.762.22c-.107-.003-.8-.11-1.117.364c-.798-.095-1.049.47-.763.994c-.163.253-.332.5.048.981c-.134.268-.05.558.267.91c-.084.376.081.641.376.848c-.055.515.472.815.629.922c.06.3.187.584.79.74c.099.447.46.525.811.619c-1.159.673-2.153 1.558-2.146 3.733l-.17.303c-1.328.808-2.524 3.406-.654 5.518c.122.66.326 1.134.508 1.66c.273 2.117 2.054 3.11 2.523 3.227c.688.524 1.422 1.02 2.413 1.369c.935.964 1.946 1.331 3.005 1.331s2.07-.367 3.005-1.331c.992-.348 1.725-.845 2.413-1.37c.47-.117 2.25-1.11 2.523-3.226c.182-.526.386-1 .508-1.66c1.87-2.112.674-4.71-.655-5.518l-.169-.303c.007-2.175-.987-3.06-2.146-3.733c.35-.094.712-.172.812-.619c.602-.156.729-.44.79-.74c.157-.107.683-.407.628-.922c.295-.207.46-.472.376-.848c.318-.352.402-.642.267-.91c.38-.48.212-.728.049-.98c.285-.526.034-1.09-.763-.995c-.318-.474-1.01-.367-1.117-.365c-.12-.152-.277-.282-.763-.22c-.315-.284-.669-.235-1.032-.095A.667.667 0 0 0 16.742 0c-.232-.007-.43.134-.643.202c-.523-.17-.641.061-.898.156c-.569-.12-.742.143-1.015.42l-.32-.007c-.858.506-1.282 1.535-1.452 1.535c-.17 0-.594-1.029-1.453-1.535l-.319.006C10.37.501 10.196.238 9.627.358C9.37.263 9.252.031 8.73.202C8.516.134 8.318-.007 8.087 0Zm.03.609c.436.161.663.37.9.575c.08-.108.202-.188.052-.45c.31.18.544.39.717.625c.191-.122.112-.287.113-.441c.322.262.528.54.778.812c.05-.037.093-.162.132-.358c.768.744 1.852 2.62.279 3.365C9.749 3.633 8.15 2.83 6.38 2.228c2.277 1.174 3.602 2.123 4.327 2.932c-.371 1.49-2.31 1.56-3.019 1.518c.145-.068.266-.15.31-.275c-.179-.127-.81-.013-1.25-.26c.17-.036.249-.069.328-.194c-.416-.132-.864-.248-1.128-.468c.143.002.275.033.46-.095c-.372-.201-.77-.36-1.078-.668c.193-.004.4-.002.46-.073a4.155 4.155 0 0 1-.867-.704c.27.033.382.006.448-.04c-.258-.265-.583-.487-.739-.812c.2.069.383.093.515-.008c-.088-.198-.464-.314-.68-.776c.211.02.436.046.48 0c-.097-.398-.264-.622-.429-.854c.451-.007 1.134.001 1.103-.037l-.279-.285c.44-.118.892.02 1.219.122c.147-.115-.004-.263-.183-.413c.374.05.713.136 1.019.255C7.56.945 7.289.799 7.159.65c.578.11.825.263 1.068.417c.177-.169.01-.312-.11-.46Zm8.594 0c-.12.147-.287.29-.11.46c.244-.155.49-.308 1.068-.418c-.13.148-.401.294-.238.442a4.431 4.431 0 0 1 1.02-.255c-.18.15-.33.298-.184.413c.328-.102.779-.24 1.22-.122l-.28.285c-.03.038.652.03 1.103.037c-.165.232-.331.456-.43.854c.045.046.27.02.48 0c-.215.462-.591.578-.679.776c.132.101.315.077.515.008c-.155.325-.48.547-.738.811c.065.047.178.074.447.041a4.155 4.155 0 0 1-.866.704c.06.07.267.069.46.073c-.31.307-.706.467-1.079.668c.186.128.318.097.46.095c-.264.22-.711.336-1.127.468c.079.125.158.158.327.193c-.44.248-1.071.134-1.249.26c.043.125.164.208.31.276c-.71.041-2.648-.029-3.02-1.518c.726-.81 2.05-1.758 4.328-2.932c-1.772.602-3.37 1.405-4.708 2.509c-1.574-.744-.49-2.62.278-3.365c.04.196.082.32.133.358c.25-.272.455-.55.777-.812c0 .154-.078.32.114.441c.172-.235.406-.446.716-.624c-.15.26-.027.341.053.45c.236-.207.463-.415.899-.576ZM12.414 6.77c1.359 0 2.492.923 2.494 1.466c.004.68-.994 1.38-2.476 1.398h-.036c-1.482-.018-2.48-.717-2.476-1.398c.003-.543 1.136-1.466 2.494-1.466Zm-3.813.447h.084c.22 0 .447.02.677.059c.777.13-3.72 4.058-3.79 3.182c-.062-2.003 1.278-3.213 3.03-3.241Zm7.542 0h.084c1.751.028 3.091 1.238 3.03 3.24c-.07.877-4.567-3.05-3.791-3.181c.23-.04.457-.058.677-.06Zm-6.825 2.5c.259-.002.521.036.777.122c1.363.458 2.052 2.053 1.54 3.56c-.513 1.509-2.034 2.36-3.398 1.903c-1.363-.458-2.052-2.053-1.54-3.56c.417-1.226 1.498-2.017 2.62-2.025Zm6.193 0c1.122.008 2.204.8 2.62 2.024c.513 1.508-.176 3.103-1.54 3.56c-1.364.459-2.885-.393-3.398-1.901c-.512-1.508.177-3.103 1.54-3.56c.256-.087.519-.125.778-.123ZM5.33 11.4c1.062.015.273 5.052-.69 4.624c-1.092-.879-1.444-3.451.582-4.608a.405.405 0 0 1 .082-.014l.026-.002Zm14.168 0l.027.002a.404.404 0 0 1 .081.014c2.026 1.157 1.675 3.73.582 4.608c-.963.428-1.752-4.609-.69-4.624Zm-7.084 3.482a2.8 2.8 0 0 1 1.874.7c.52.47.822 1.14.818 1.811c0 .66-.295 1.313-.8 1.778a2.88 2.88 0 0 1-1.892.739a2.884 2.884 0 0 1-1.892-.739a2.448 2.448 0 0 1-.8-1.778a2.456 2.456 0 0 1 .818-1.81a2.796 2.796 0 0 1 1.874-.7ZM6.148 16.2c.703-.014 1.568.541 2.268 1.353c.78.94 1.136 2.593.485 3.08c-.617.372-2.115.218-3.178-1.31c-.718-1.282-.624-2.587-.12-2.97c.164-.1.348-.149.545-.153Zm12.465 0h.067c.197.004.38.052.546.153c.503.383.597 1.688-.12 2.97c-1.064 1.528-2.562 1.682-3.178 1.31c-.652-.487-.297-2.14.484-3.08c.678-.787 1.51-1.333 2.201-1.353Zm-6.199 4.504c1.09-.012 2.71.447 2.69 1.032c.017.405-1.312 1.574-2.667 1.516h-.045c-1.356.058-2.685-1.11-2.667-1.516c-.02-.585 1.6-1.044 2.69-1.032Z"/>`
	reactjsInnerSVG                = `<g fill="currentColor" fill-rule="evenodd"><circle cx="11.996" cy="11.653" r="2.142"/><path fill-rule="nonzero" d="M11.996 7.81c2.768 0 5.397.39 7.396 1.078c2.22.764 3.575 1.894 3.575 2.765c0 .927-1.463 2.131-3.838 2.918c-1.89.626-4.42.968-7.133.968c-2.85 0-5.409-.335-7.277-.974c-1.17-.4-2.136-.92-2.799-1.486c-.595-.51-.895-1.016-.895-1.426c0-.874 1.302-1.988 3.475-2.746c2-.697 4.693-1.098 7.496-1.098m0-1.025c-2.912 0-5.718.418-7.834 1.156C1.622 8.825 0 10.213 0 11.653c0 1.486 1.741 2.978 4.387 3.882c1.989.68 4.654 1.029 7.609 1.029c2.816 0 5.451-.356 7.456-1.02c2.75-.912 4.54-2.385 4.54-3.891c0-1.445-1.675-2.842-4.266-3.735c-2.113-.727-4.854-1.134-7.73-1.134m-3.348 2.96c1.383-2.399 3.034-4.481 4.63-5.87c1.77-1.54 3.426-2.15 4.18-1.714c.803.463 1.116 2.331.611 4.782c-.402 1.95-1.37 4.312-2.725 6.663c-1.424 2.469-2.992 4.518-4.48 5.817c-.93.814-1.863 1.391-2.685 1.682c-.738.26-1.327.268-1.682.063c-.757-.436-1.072-2.121-.643-4.381c.395-2.081 1.394-4.615 2.794-7.043m-.888-.512c-1.455 2.522-2.495 5.162-2.913 7.363c-.501 2.643-.11 4.742 1.137 5.46c1.289.744 3.45-.02 5.555-1.86c1.583-1.383 3.217-3.518 4.693-6.077c1.406-2.44 2.415-4.9 2.84-6.969c.585-2.837.203-5.124-1.102-5.876c-1.252-.722-3.298.03-5.366 1.83C10.92 4.568 9.197 6.74 7.76 9.231m.891 4.378C7.264 11.215 6.285 8.744 5.879 6.67c-.451-2.305-.152-4.043.602-4.479c.802-.465 2.577.198 4.448 1.859c1.49 1.322 3.052 3.34 4.411 5.689c1.429 2.467 2.42 4.848 2.804 6.785c.24 1.213.274 2.31.115 3.167c-.143.77-.43 1.284-.785 1.49c-.756.437-2.373-.132-4.117-1.633c-1.605-1.38-3.302-3.511-4.706-5.937m-.887.514c1.459 2.52 3.226 4.74 4.925 6.2c2.039 1.755 4.052 2.463 5.298 1.742c1.287-.745 1.705-2.999 1.162-5.742c-.408-2.061-1.441-4.543-2.922-7.1c-1.41-2.437-3.039-4.54-4.618-5.942C9.443 1.36 7.271.548 5.967 1.303c-1.25.724-1.62 2.873-1.094 5.563c.429 2.193 1.45 4.769 2.891 7.258"/></g>`
	redditInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M15.57 15.284c-.897 0-1.65-.728-1.65-1.625s.753-1.648 1.65-1.648c.897 0 1.625.752 1.625 1.649c0 .896-.728 1.624-1.625 1.624m.381 3.153c-.835.835-2.124 1.24-3.939 1.24h-.026c-1.815 0-3.102-.405-3.937-1.24a.655.655 0 1 1 .927-.927c.576.576 1.56.856 3.01.856l.013.001h.013c1.45 0 2.435-.281 3.012-.857a.655.655 0 1 1 .927.927m-9.146-4.778c0-.896.753-1.648 1.649-1.648c.897 0 1.624.752 1.624 1.649c0 .896-.727 1.624-1.624 1.624c-.896 0-1.649-.728-1.649-1.625M19.998 3.311c.607 0 1.102.494 1.102 1.101s-.495 1.102-1.102 1.102a1.103 1.103 0 0 1-1.102-1.102c0-.607.494-1.101 1.102-1.101M24 11.875a2.887 2.887 0 0 0-2.884-2.884c-.689 0-1.321.243-1.818.647c-1.758-1.105-3.99-1.771-6.383-1.912l1.248-3.946l3.43.808a2.415 2.415 0 0 0 2.405 2.237a2.415 2.415 0 0 0 2.412-2.413A2.415 2.415 0 0 0 19.998 2c-.93 0-1.739.53-2.141 1.303l-3.986-.938a.655.655 0 0 0-.774.44l-1.55 4.897c-2.578.063-5.001.732-6.889 1.902a2.87 2.87 0 0 0-1.774-.613A2.887 2.887 0 0 0 0 11.875a2.88 2.88 0 0 0 1.249 2.373a5.063 5.063 0 0 0-.048.693c0 1.988 1.155 3.837 3.254 5.207c2.011 1.313 4.674 2.036 7.496 2.036s5.485-.723 7.497-2.036c2.098-1.37 3.254-3.22 3.254-5.207c0-.213-.015-.424-.04-.633A2.884 2.884 0 0 0 24 11.875"/>`
	redhatInnerSVG                 = `<g fill="currentColor" fill-rule="evenodd"><path d="M5.832 6.125c.645-2.404 1.531-3.365 2.66-2.885c1.691.722 2.416.481 3.141 0c.484-.32 1.209-.32 2.175 0l3.384 1.443c.967.32 1.611 1.442 1.934 3.366c.322 1.923.564 3.205.725 3.846c2.417.962 3.786 2.405 4.109 4.328c.483 2.885-3.384 5.77-11.36 4.327C4.624 19.108-.452 14.54.032 11.895c.322-1.763 1.853-2.644 4.592-2.644l1.208-3.126Z"/><path d="M4.887 8.538c1.462 2.507 4.142 4.074 8.04 4.7c3.897.627 6.171 0 6.82-1.88c.278 1.417.278 2.357 0 2.82c-.914 1.527-3.411 1.969-6.09 1.646c-3.897-.47-6.74-1.724-8.526-3.76c-.487-.627-.731-1.175-.731-1.646c0-.47.162-1.096.487-1.88Z"/></g>`
	redoInnerSVG                   = `<path fill="currentColor" d="M16.82 4L15.4 5.44L17.94 8H8.23a6 6 0 0 0 0 12h2v-2h-2a4 4 0 0 1 0-8h9.71l-2.54 2.51l1.41 1.41L21.77 9Z"/>`
	refreshInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20 8c-1.403-2.96-4.463-5-8-5a9 9 0 1 0 0 18a9 9 0 0 0 9-9m0-9v6h-6"/>`
	resourcesInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="m12 3l9 4.5l-9 4.5l-9-4.5L12 3Zm4.5 7.25L21 12.5L12 17l-9-4.5l4.5-2.25m9 5L21 17.5L12 22l-9-4.5l4.5-2.25"/>`
	restaurantInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 18H5h14Zm-7 0v-6v6Zm3 0v-4v4Zm-6 0v-4v4Zm10 4V11.33a3.001 3.001 0 1 0-2.08-5.63C16.55 3.874 14.46 2 12 2S7.45 3.874 7.08 5.7A3 3 0 1 0 5 11.33V22h14Z"/>`
	restroomInnerSVG               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M6 11h12M6 3h12m-6 13a5 5 0 0 0 5-5H7a5 5 0 0 0 5 5ZM7 3h10v8H7V3Zm0 3h2.5m5 9.5l1.5 6H8l1.5-6"/>`
	restroomMenInnerSVG            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 13.5L11 8l-1 13m7-7.5L13 8l1 13M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1 3h2v5.5h-2V8Z"/>`
	restroomWomenInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 13.5L11 8l1 13m5-7.5L13 8l-1 13m0-16a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1 3h2l1.5 8.5h-5L11 8Z"/>`
	resumeInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 20h5V4H1v16Zm10-1l11-7l-11-7v14Z"/>`
	returnInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m9 19l-5-5l5-5m9-5v10H5"/>`
	revertInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 3L3 8l5 5m4 7h3a6 6 0 1 0 0-12H4"/>`
	rewindInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M22 3.5V20l-9-6v6L2 12l11-8v6z"/>`
	riskInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M14 10h1V9h-1v1Zm4 0h1V9h-1v1Zm0-4h1V5h-1v1Zm-4 0h1V5h-1v1ZM9 19h1v-1H9v1Zm-4-4h1v-1H5v1Zm5-5H1v13h13v-9m-4 0h13V1H10v13Z"/>`
	robotInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18.348 15.954a7 7 0 1 0-12.622.156M12 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 20a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-17V3M9 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm6 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-9 4.988L5 16s.072-.772.5-.5c.93.591 3.074 1.5 6.5 1.5c3.554 0 5.618-.916 6.5-1.5c.359-.238.5.5.5.5l-1 2.988S17.005 21 12 21s-6-2.012-6-2.012Z"/>`
	rotateLeftInnerSVG             = `<path fill="currentColor" d="M11.52 3.43A9.09 9.09 0 0 0 5.7 5.55v-3.2H4.07v6.5h6.5V7.21H6.3a7.46 7.46 0 1 1-1.47 8.65l-1.46.73a9.11 9.11 0 1 0 8.15-13.16Z"/>`
	rotateRightInnerSVG            = `<path fill="currentColor" d="M12.48 3.43a9.09 9.09 0 0 1 5.82 2.12v-3.2h1.64v6.5h-6.5V7.21h4.26a7.46 7.46 0 1 0 1.47 8.65l1.46.73a9.11 9.11 0 1 1-8.15-13.16Z"/>`
	rssInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M22 21C22 10.507 13.493 2 3 2m14 19c0-7.732-6.268-14-14-14m9 14a9 9 0 0 0-9-9m1 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`
	runInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l3 2m0-8a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM9.5 9.5L9.525 6H15L8 17H4m11-9l-3 5l.5 1L17 7.5L15 6m-4 7l5 3.5v5"/>`
	safariOptionInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M12.541 11.844c.022-.414-.33-.772-.775-.79c-.47-.02-.827.325-.835.802c-.007.423.337.758.79.772c.45.014.795-.317.82-.784m4.673-6.722l-.053-.04c-.053.044-.109.087-.16.135c-1.914 1.787-3.831 3.571-5.741 5.364a3.48 3.48 0 0 0-.588.696c-1.401 2.307-2.789 4.623-4.179 6.936c-.057.094-.102.194-.153.292a.242.242 0 0 0 .172-.069c1.94-1.812 3.883-3.62 5.814-5.443a4.03 4.03 0 0 0 .65-.824c1.371-2.256 2.73-4.521 4.092-6.783c.052-.086.098-.176.146-.264m-.07-1.094c1.999 1.394 3.251 3.257 3.89 5.607c-.233.057-.45.093-.652.167c-.087.032-.145.144-.216.219c.094.041.193.124.283.117c.206-.02.408-.083.655-.138c.45 2.407.068 4.629-1.183 6.729c-.195-.118-.357-.242-.537-.317c-.12-.05-.266-.033-.4-.045c.065.109.11.242.202.323c.148.129.327.224.533.36c-1.375 2.01-3.215 3.324-5.6 3.961c-.07-.266-.121-.526-.212-.773c-.033-.09-.157-.147-.241-.22c-.026.1-.087.206-.074.3c.035.242.103.48.17.774c-2.431.453-4.689.095-6.835-1.19c.164-.256.33-.483.458-.73c.053-.101.025-.244.034-.368c-.113.06-.259.095-.332.187c-.167.208-.298.446-.47.713c-2.039-1.396-3.342-3.275-3.983-5.68c.31-.072.574-.12.826-.203c.085-.028.14-.146.21-.224c-.1-.033-.204-.104-.299-.092c-.253.033-.502.1-.797.164c-.448-2.44-.062-4.696 1.251-6.829c.19.12.354.25.54.334c.144.066.314.074.473.108c-.096-.14-.172-.3-.294-.412c-.145-.131-.33-.219-.525-.343c1.425-1.992 3.285-3.301 5.708-3.895c.056.241.087.471.165.683c.045.119.165.208.251.31c.03-.12.092-.245.08-.362c-.022-.22-.087-.434-.145-.698c2.415-.423 4.647-.028 6.751 1.255c-.15.242-.302.46-.42.695c-.042.082.003.208.008.313c.09-.04.208-.056.262-.123c.16-.201.293-.422.465-.677M11.845 22.42c5.876-.014 10.586-4.76 10.577-10.659c-.008-5.772-4.783-10.507-10.589-10.499C5.936 1.27 1.25 6.01 1.264 11.95c.015 5.752 4.797 10.484 10.581 10.47M11.836 0c6.506-.007 11.845 5.307 11.85 11.794A11.86 11.86 0 0 1 11.86 23.683C5.328 23.715 0 18.378 0 11.8C0 5.303 5.312.007 11.836.001"/>`
	sansInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 12h22M2 22h20a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1zM5 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`
	satelliteInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 17C2.97 12.794 2.97 6.118 7 2l15 15c-4.118 4.03-10.794 4.03-15 0Zm0 0c-3.295 0-6 2.95-6 6h12c0-1.139-.37-2.034-1-3m3-11l4-4l-4 4Zm5.5-8a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z"/>`
	saveInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 2v19h18V3h-9v11m-4-3l4 4l4-4"/>`
	scanInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20 10V3H4v7m-3 2h22H1Zm3 1v3v-3Zm16 3v-3v3ZM7 21H4v-3m16 0v3h-3m-8 0h6h-6Z"/>`
	scheduleInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 23h22V4H1v19ZM18 4V0v4ZM6 4V0v4ZM1 8.5h22H1ZM6 14c.556-1.333 1.39-2 2.5-2c1.3 0 2 1 2 2s-1 2-2 3l-2 2v.5h5.405m5.08 1L17 12h-.5c-.5 1.5-2 2-2.743 2"/>`
	scheduleNewInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M14 0v3M1 7h18M6 0v3m-2 8h2m2 0h8M4 15h2m2 0h6m-2 4H1V3h18v10m0 2v9m-4-7l8 5m0-5l-8 5"/>`
	schedulePlayInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M14 0v3M1 7h18M6 0v3m-2 8h2m2 0h8M4 15h2m2 0h6m-1 4H1V3h18v10m-1 10a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm-.5-6l1.5 1l-1.5 1v-2Z"/>`
	schedulesInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17 7h6v16H7v-4m16-8h-6M13 0v3M1 7h16M1 3h16v16H1V3Zm4-3v3m-1 8h2m2 0h6M4 15h2m2 0h6"/>`
	scoInnerSVG                    = `<path fill="currentColor" fill-rule="evenodd" d="M16.16 16.105H24c-.247.746-.99 1.367-1.794 1.367H1.273c-.804 0-1.33-.62-1.268-1.367h14.702a3.784 3.784 0 0 1-.2-1.957h.959c0 .187 0 .343.03.497c.31-.093.62-.28.897-.497h.619c-.526.56-1.02 1.025-1.268 1.212c.114.277.261.528.416.745Zm5.583-8.39C20.692 7.28 19.362 7 17.847 7c-1.546 0-3.03.28-4.206.715l.032.062h8.009l.06-.062Zm-6.617 4.941c.031-.062.062-.124.093-.218l-10.792.032l-.03-.125c2.659-.528 5.967-.839 9.523-.839c.557 0 1.144 0 1.67.031c.03-.124.062-.217.093-.342H7.026l-.031-.093c2.319-.497 5.225-.808 8.348-.808h.34c-.03-.062-.061-.186-.092-.28H9.715l-.03-.093a29.155 29.155 0 0 1 5.38-.684a1.616 1.616 0 0 0-.34-.31h-2.876l-.062-.094c1.484-.466 3.278-.745 5.226-.745c1.887 0 3.556.31 4.885.745l-.03.063l-6.555.03c.092.063.185.156.278.28c.37-.031.741-.031 1.113-.031c2.35 0 4.514.28 6.245.653l-.03.186h-6.896c.031.094.031.218.062.311c2.505.032 4.793.31 6.71.683l-.031.187h-6.617c0 .125 0 .249-.031.342c2.412.124 4.638.373 6.494.745l-.031.156h-4.144a.554.554 0 0 1-.155.218h-.494c.03-.062.093-.125.124-.218h-1.948c-.031.093-.031.156-.063.218h-.773Zm-.62 1.492c.032-.156.063-.31.125-.467H6.839l-.062-.124c2.32-.466 5.195-.777 8.287-.777c0-.031.031-.093.062-.124h.773c0 .031-.031.093-.031.124c.587.031 1.206.031 1.762.093a.958.958 0 0 0 .154-.217h.495c-.03.093-.093.155-.186.249a33.22 33.22 0 0 1 4.36.559l-.03.217h-4.98c-.154.156-.278.31-.432.467h-.618c.217-.125.402-.31.557-.467h-1.423a2.767 2.767 0 0 0-.062.467h-.958Z"/>`
	scorecardInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 18v-6m3 6v-5m3 5v-8m2-7h4v20H3V3h4m0-2h10v4H7V1Z"/>`
	scriptInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M4 1a4 4 0 0 0-4 4v1h5.847L4.01 18.859l-.01.07V19a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4v-1h-5.847L19.99 5.141l.01-.07V5a4 4 0 0 0-4-4H4Zm0 2a2 2 0 0 0-1.732 1h3.464A2 2 0 0 0 4 3Zm3.465 0H16a2 2 0 0 1 1.999 1.936L16.133 18H10v1a2 2 0 0 1-3.999.064L7.99 5.141l.01-.07V5c0-.729-.195-1.412-.535-2ZM20 21h-8.535a3.97 3.97 0 0 0 .409-1h9.858A2 2 0 0 1 20 21Z" clip-rule="evenodd"/>`
	sdInnerSVG                     = `<path fill="currentColor" fill-rule="evenodd" d="M2 2a2 2 0 0 1 2-2h13a1 1 0 0 1 .707.293l4 4A1 1 0 0 1 22 5v17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2Zm14.586 0H4v20h16V5.414L16.586 2ZM5 7V3h2v4H5Zm3-4v4h2V3H8Zm3 4V3h2v4h-2Zm3-2v2h2V5h-2Z" clip-rule="evenodd"/>`
	searchInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m15 15l7 7l-7-7Zm-5.5 2a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Z"/>`
	searchAdvancedInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="m15 16l6 6l-6-6Zm-5 2a7 7 0 1 0 0-14a7 7 0 0 0 0 14ZM20 1v6m-3-3h6"/>`
	secureInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 11V6c0-3 2-5 5-5s5 2 5 5v5m-5 12a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0-8v4m0-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	selectInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 1h6h-6Zm11.188 18.472L16 22l-3.5-4.5l-3 3.5L7 7l13 6.5l-4.5 1.5l3.688 4.472ZM19 4V1h-3M6 1H3v3m0 10v3h3M19 6v4v-4ZM3 12V6v6Z"/>`
	selectionInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 18h3V6H5v12Zm7-16v20V2ZM1 22h22V2H1v20Z"/>`
	semanticsInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="m2 17l10-5l10 5v4l-10-5l-10 5v-4Zm0-9l10-5l10 5v4L12 7L2 12V8Z"/>`
	sendInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M22 3L2 11l18.5 8L22 3ZM10 20.5l3-4.5m2.5-6.5L9 14l.859 6.012c.078.546.216.537.306-.003L11 15l4.5-5.5Z"/>`
	serverInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M14 19h1v-1h-1v1Zm-9 4h14V1H5v22ZM8 5h8h-8Zm0 4h8h-8Zm0 4h8h-8Z"/>`
	serverClusterInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 8h22V1H1v7Zm10-3h1V4h-1v1Zm8 0h1V4h-1v1Zm-4 0h1V4h-1v1Zm-4 7h1v-1h-1v1Zm8 0h1v-1h-1v1Zm-4 0h1v-1h-1v1Zm-4 7h1v-1h-1v1Zm8 0h1v-1h-1v1Zm-4 0h1v-1h-1v1ZM1 15h22V8H1v7Zm0 7h22v-7H1v7Zm20 1H3"/>`
	serversInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 19h1v-1H7v1Zm11 0h1v-1h-1v1ZM1 23h11V1H1v22Zm11 0h11V1H12v22ZM4 5h5h-5Zm11 0h5h-5ZM4 9h5h-5Zm11 0h5h-5ZM4 13h5h-5Zm11 0h5h-5Z"/>`
	servicePlayInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 13A6 6 0 1 0 7 1a6 6 0 0 0 0 12Zm7.995 3.657a6 6 0 1 0-1.89-10.22m-8.281 6.255A6 6 0 0 0 9 23a6 6 0 0 0 2.127-11.612M6.5 6L8 7L6.5 8V6Z"/>`
	servicesInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M6 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-6V0m0 12V9M0 6h3m6 0h3M2 2l2 2m4 4l2 2m0-8L8 4M4 8l-2 2m16 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-6V3m0 12v-3m-6-3h3m6 0h3M14 5l2 2m4 4l2 2m0-8l-2 2m-4 4l-2 2m-5 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-6v-3m0 12v-3m-6-3h3m6 0h3M5 14l2 2m4 4l2 2m0-8l-2 2m-4 4l-2 2"/>`
	settingsOptionInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 9V0m3 12h9M0 12h9m3 12v-9m0 6a9 9 0 1 0 0-18a9 9 0 0 0 0 18ZM3.5 8.5L1 7.5m19.5 8l2.5 1M3 3l2.5 2.5M3 3l2.5 2.5M18 18l2.5 2.5m0-17.5L18 5.5M5.5 18L3 20.5m9-5.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm8.5-6.5l2.5-1m-7.5-4l1-2.5m-1 19.5l1 2.5m-8-2.5l-1 2.5m-4-7.5l-2.5 1m7.5-13L7.5 1"/>`
	shareInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 13v10H1V5h10m3-4h9v9m-13 4L23 1L10 14Z"/>`
	shareOptionInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM6 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm12 7a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm-2-4l-8-5m8-7l-8 5"/>`
	shareRoundedInnerSVG           = `<path fill="currentColor" fill-rule="evenodd" d="M20.924 3.617a.997.997 0 0 0-.215-.322l-.004-.004A.997.997 0 0 0 20 3h-6a1 1 0 1 0 0 2h3.586l-7.293 7.293a1 1 0 1 0 1.414 1.414L19 6.414V10a1 1 0 1 0 2 0V3.997a.999.999 0 0 0-.076-.38ZM3 8a5 5 0 0 1 5-5h1a1 1 0 0 1 0 2H8a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-1a1 1 0 1 1 2 0v1a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8Z" clip-rule="evenodd"/>`
	shieldInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22S3 18 3 5l9-3l9 3c0 13-9 17-9 17ZM4 11h16m-8-9v20"/>`
	shieldSecurityInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22s-9-4-9-11V5l9-3l9 3v6c0 7-9 11-9 11Zm0-8a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-6V5m0 12v-3m-6-3h3m6 0h3M8 7l2 2m4 4l2 2m0-8l-2 2m-4 4l-2 2"/>`
	shiftInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 0v24M2 12h10m10 0H12M6 8l-4 4l4 4m12-8l4 4l-4 4"/>`
	shopInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 7h16v16H4V7Zm4 2V5c0-2.21 1.795-4 4-4h0c2.21 0 4 1.795 4 4v4"/>`
	sidebarInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 22h22V2H1v20ZM16 2v20V2Z"/>`
	signInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 23h7c3 0 4-2 4-5V6c0-2-1-2-1.5-2S16 4 16 6v7c0-1 0-2-1.5-2S13 13 13 13c0-1 0-2-1.5-2S10 12 10 13V4c0-1 .03-2-1.5-2C7 2 7 3 7 4v14v-4c0-1-.55-2-2-2H4v6c0 3.962 2 5.024 4 5Z"/>`
	skypeInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M12.052 18.856c-4.027 0-5.828-1.98-5.828-3.463c0-.761.562-1.295 1.336-1.295c1.724 0 1.277 2.475 4.492 2.475c1.645 0 2.554-.894 2.554-1.809c0-.55-.271-1.159-1.355-1.426l-3.581-.894c-2.884-.723-3.407-2.282-3.407-3.748c0-3.043 2.864-4.185 5.556-4.185c2.477 0 5.4 1.369 5.4 3.194c0 .783-.678 1.238-1.452 1.238c-1.47 0-1.2-2.035-4.162-2.035c-1.47 0-2.283.665-2.283 1.618c0 .95 1.16 1.254 2.168 1.483l2.651.59c2.903.646 3.64 2.34 3.64 3.938c0 2.472-1.898 4.319-5.73 4.319m11.1-4.887l-.022.127l-.04-.241c.021.037.04.076.061.114c.124-.674.19-1.364.19-2.054a11.297 11.297 0 0 0-3.32-8.014A11.287 11.287 0 0 0 12.006.583c-.722 0-1.444.067-2.147.202l-.005.001c.04.021.08.04.118.061L9.736.81l.12-.024A6.722 6.722 0 0 0 6.709 0a6.663 6.663 0 0 0-4.744 1.965A6.666 6.666 0 0 0 0 6.71c0 1.14.293 2.26.844 3.252c.007-.041.012-.083.02-.123l.041.237C.883 10.04.865 10 .844 9.962a11.402 11.402 0 0 0-.171 1.953c0 1.53.3 3.015.891 4.412a11.285 11.285 0 0 0 2.428 3.602a11.268 11.268 0 0 0 3.603 2.428c1.397.592 2.882.892 4.412.892c.665 0 1.332-.061 1.984-.177c-.038-.02-.077-.04-.115-.063l.242.043c-.042.008-.084.013-.127.02c1.004.569 2.14.87 3.3.87a6.66 6.66 0 0 0 4.744-1.965A6.66 6.66 0 0 0 24 17.233a6.707 6.707 0 0 0-.85-3.264"/>`
	slackInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M5.048 15.124a2.512 2.512 0 0 1-2.515 2.514A2.512 2.512 0 0 1 .02 15.124a2.512 2.512 0 0 1 2.514-2.514h2.515v2.514zm1.257 0a2.512 2.512 0 0 1 2.514-2.514a2.512 2.512 0 0 1 2.514 2.514v6.286a2.512 2.512 0 0 1-2.514 2.514a2.512 2.512 0 0 1-2.514-2.514v-6.286zM8.819 5.029a2.512 2.512 0 0 1-2.514-2.515A2.512 2.512 0 0 1 8.819 0a2.512 2.512 0 0 1 2.514 2.514V5.03H8.82zm0 1.276a2.512 2.512 0 0 1 2.514 2.514a2.512 2.512 0 0 1-2.514 2.514H2.514A2.512 2.512 0 0 1 0 8.82a2.512 2.512 0 0 1 2.514-2.514H8.82zm10.076 2.514a2.512 2.512 0 0 1 2.515-2.514a2.512 2.512 0 0 1 2.514 2.514a2.512 2.512 0 0 1-2.514 2.514h-2.515V8.82zm-1.257 0a2.512 2.512 0 0 1-2.514 2.514a2.512 2.512 0 0 1-2.515-2.514V2.514A2.512 2.512 0 0 1 15.124 0a2.512 2.512 0 0 1 2.514 2.514V8.82zm-2.514 10.076a2.512 2.512 0 0 1 2.514 2.514a2.512 2.512 0 0 1-2.514 2.515a2.512 2.512 0 0 1-2.514-2.515v-2.514h2.514zm0-1.257a2.512 2.512 0 0 1-2.514-2.514a2.512 2.512 0 0 1 2.514-2.514h6.305a2.512 2.512 0 0 1 2.514 2.514a2.512 2.512 0 0 1-2.514 2.514h-6.305z"/>`
	snapchatInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M12.151 22.532c-.068 0-.133-.003-.182-.005c-.04.003-.08.005-.12.005c-1.43 0-2.385-.676-3.228-1.272c-.604-.427-1.174-.83-1.842-.941a5.947 5.947 0 0 0-.97-.082c-.567 0-1.016.088-1.344.152c-.201.039-.375.073-.51.073c-.14 0-.311-.031-.383-.275a8.157 8.157 0 0 1-.136-.557c-.098-.447-.169-.72-.336-.746C1.308 18.607.249 18.2.038 17.706A.458.458 0 0 1 0 17.551a.288.288 0 0 1 .241-.3c1.423-.235 2.689-.987 3.762-2.237a8.434 8.434 0 0 0 1.29-2.008c.206-.42.247-.782.122-1.078c-.231-.544-.996-.787-1.502-.948a5.405 5.405 0 0 1-.34-.115c-.448-.177-1.186-.551-1.088-1.068c.072-.377.57-.64.973-.64c.112 0 .211.02.294.06c.456.213.865.321 1.217.321c.438 0 .65-.167.7-.214a74.562 74.562 0 0 0-.042-.717c-.103-1.636-.231-3.67.29-4.838C7.473.276 10.777.005 11.752.005A249 249 0 0 0 12.236 0c.978 0 4.289.272 5.848 3.767c.52 1.168.392 3.205.289 4.842l-.005.078a67.21 67.21 0 0 0-.038.637c.048.044.242.197.635.212c.336-.013.722-.12 1.147-.32a.906.906 0 0 1 .375-.074c.15 0 .301.03.428.082l.007.003c.36.128.596.384.601.652c.005.25-.181.625-1.097.986a5.552 5.552 0 0 1-.34.116c-.506.16-1.27.403-1.501.947c-.126.295-.084.658.123 1.077l.006.014c.064.15 1.605 3.665 5.045 4.231a.288.288 0 0 1 .24.3a.462.462 0 0 1-.038.157c-.209.491-1.268.898-3.06 1.175c-.169.026-.24.298-.337.743c-.04.184-.08.364-.136.553c-.053.179-.169.266-.355.266h-.028a2.83 2.83 0 0 1-.51-.065a6.712 6.712 0 0 0-1.345-.142c-.315 0-.64.027-.97.082c-.666.11-1.236.513-1.84.94c-.844.597-1.8 1.273-3.229 1.273"/>`
	solarisInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M10.18 4.764c-.872-.309-1.27-1.127-1.27-1.127c-.068-.094-.353-.779-.353-.78c-.3-.792-.845-1.035-.845-1.035c-.671-.394-1.5-.404-1.5-.404c.035-.005.493.207.493.207c1.08.515.756 1.744.756 1.744c-.07.331-.086.97-.086.97a1.55 1.55 0 0 0 .659 1.312a7.557 7.557 0 0 1 2.147-.887M5.597 8.17c-.836.399-1.696.102-1.696.102c-.115-.018-.8-.3-.8-.3c-.773-.35-1.33-.136-1.33-.136c-.754.196-1.347.775-1.347.775c.021-.028.495-.202.495-.202c1.128-.4 1.768.698 1.768.698c.184.284.625.747.625.747a1.551 1.551 0 0 0 1.393.462a7.54 7.54 0 0 1 .892-2.146m-.832 5.65c-.31.872-1.127 1.272-1.127 1.272c-.095.067-.78.352-.78.352c-.793.3-1.035.846-1.035.846c-.395.67-.404 1.5-.404 1.5c-.005-.036.206-.494.206-.494c.515-1.08 1.745-.756 1.745-.756c.33.07.969.086.969.086a1.55 1.55 0 0 0 1.312-.658a7.596 7.596 0 0 1-.886-2.148m3.406 4.583c.398.836.102 1.696.102 1.696c-.019.115-.301.801-.302.801c-.349.772-.134 1.33-.134 1.33c.195.754.775 1.347.775 1.347c-.029-.022-.203-.496-.203-.496c-.4-1.127.699-1.767.699-1.767c.283-.185.747-.626.747-.626a1.55 1.55 0 0 0 .461-1.393a7.602 7.602 0 0 1-2.145-.892m5.649.833c.873.309 1.272 1.127 1.272 1.127c.068.094.353.78.352.78c.3.792.846 1.035.846 1.035c.671.395 1.5.404 1.5.404c-.035.005-.493-.207-.493-.207c-1.08-.515-.757-1.744-.757-1.744c.07-.33.086-.97.086-.97a1.55 1.55 0 0 0-.658-1.312a7.587 7.587 0 0 1-2.148.887m4.584-3.406c.835-.399 1.696-.102 1.696-.102c.115.018.8.301.8.301c.773.35 1.33.135 1.33.135c.754-.196 1.347-.775 1.347-.775c-.022.029-.495.202-.495.202c-1.128.4-1.768-.698-1.768-.698c-.185-.284-.625-.747-.625-.747a1.55 1.55 0 0 0-1.394-.462a7.573 7.573 0 0 1-.891 2.146m.832-5.65c.31-.872 1.127-1.272 1.127-1.272c.094-.067.78-.353.78-.352c.792-.3 1.035-.845 1.035-.845c.395-.672.405-1.5.405-1.5c.004.035-.208.493-.208.493c-.514 1.08-1.744.756-1.744.756c-.33-.07-.97-.086-.97-.086a1.551 1.551 0 0 0-1.312.658c.204.335.384.692.536 1.066c.147.358.264.72.351 1.082M15.83 5.597c-.398-.836-.103-1.696-.103-1.696c.02-.115.302-.801.303-.8c.349-.773.134-1.33.134-1.33c-.196-.755-.775-1.348-.775-1.348c.029.022.203.496.203.496c.399 1.127-.7 1.768-.7 1.768c-.282.185-.746.625-.746.625a1.55 1.55 0 0 0-.462 1.393a7.562 7.562 0 0 1 2.146.892m-8.243.387c-.925.042-1.602-.565-1.602-.565c-.099-.061-.62-.588-.62-.589c-.577-.621-1.174-.64-1.174-.64c-.77-.113-1.542.19-1.542.19c.031-.017.535.007.535.007c1.194.07 1.358 1.33 1.358 1.33c.06.333.286.93.286.93a1.55 1.55 0 0 0 1.104.967a7.566 7.566 0 0 1 1.655-1.63m-2.961 4.882c-.624.684-1.533.734-1.533.734c-.113.026-.855.023-.855.022C1.391 11.59.956 12 .956 12C.33 12.464 0 13.225 0 13.225c.01-.035.383-.375.383-.375c.893-.795 1.9-.02 1.9-.02c.278.194.86.457.86.457a1.55 1.55 0 0 0 1.465-.097a7.57 7.57 0 0 1 .018-2.324m1.358 5.547c.042.925-.564 1.603-.564 1.603c-.062.098-.59.62-.59.62c-.621.576-.64 1.173-.64 1.173c-.113.77.19 1.542.19 1.542c-.017-.03.007-.535.007-.535c.07-1.194 1.33-1.358 1.33-1.358c.333-.06.93-.285.93-.285a1.55 1.55 0 0 0 .968-1.105a7.603 7.603 0 0 1-1.631-1.655m4.882 2.962c.684.624.734 1.532.734 1.532c.026.113.023.855.023.855c-.033.847.376 1.283.376 1.283c.465.624 1.226.955 1.226.955c-.035-.01-.375-.382-.375-.382c-.795-.894-.019-1.901-.019-1.901c.193-.278.456-.86.456-.86a1.552 1.552 0 0 0-.097-1.465a7.544 7.544 0 0 1-2.324-.017m5.547-1.358c.925-.043 1.603.564 1.603.564c.098.062.62.589.62.589c.576.622 1.173.64 1.173.64c.77.113 1.542-.19 1.542-.19c-.03.017-.535-.006-.535-.006c-1.194-.07-1.358-1.33-1.358-1.33c-.06-.333-.285-.931-.285-.931a1.55 1.55 0 0 0-1.105-.967a7.55 7.55 0 0 1-1.655 1.63m2.962-4.882c.624-.684 1.532-.734 1.532-.734c.114-.026.856-.023.856-.022c.846.032 1.282-.377 1.282-.377c.624-.465.955-1.225.955-1.225c-.01.034-.382.374-.382.374c-.894.795-1.901.02-1.901.02c-.278-.194-.86-.457-.86-.457a1.55 1.55 0 0 0-1.465.097a7.56 7.56 0 0 1-.017 2.324m-1.358-5.547c-.043-.925.564-1.602.564-1.602c.062-.099.589-.621.59-.62c.62-.577.64-1.174.64-1.174c.112-.77-.191-1.542-.191-1.542c.018.031-.006.535-.006.535c-.07 1.194-1.33 1.358-1.33 1.358c-.333.06-.931.286-.931.286a1.551 1.551 0 0 0-.968 1.104a7.574 7.574 0 0 1 1.632 1.655m-4.883-2.961c-.684-.624-.734-1.533-.734-1.533c-.026-.113-.023-.855-.022-.855C12.41 1.39 12 .956 12 .956C11.536.33 10.776 0 10.776 0c.034.01.374.383.374.383c.795.894.02 1.9.02 1.9c-.193.278-.457.86-.457.86c-.345.78.053 1.399.097 1.466a7.526 7.526 0 0 1 2.324.017"/>`
	sortInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 3h15M7 9h9m-9 6h15M2 2h2v2H2V2Zm0 6h2v2H2V8Zm0 6h2v2H2v-2Zm0 6h2v2H2v-2Zm5 1h9"/>`
	soundcloudInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="M20.727 11.677c1.674 0 3.03 1.368 3.03 3.055c0 1.688-1.356 3.057-3.03 3.057l-8.4-.005a.368.368 0 0 1-.327-.361v-9.7c.002-.178.063-.27.292-.359A5.377 5.377 0 0 1 14.23 7c2.794 0 5.084 2.16 5.325 4.914a3.01 3.01 0 0 1 1.172-.237ZM10.237 18h-.973A.268.268 0 0 1 9 17.73V7.77c0-.148.119-.27.264-.27h.972c.145 0 .264.122.264.27v9.96c0 .148-.119.27-.264.27Zm-9-1.503H.263A.267.267 0 0 1 0 16.228v-2.462c0-.148.119-.269.264-.269h.972c.145 0 .264.121.264.27v2.461c0 .148-.119.27-.264.27Zm6 1.503h-.973A.261.261 0 0 1 6 17.743v-6.486c0-.141.119-.257.264-.257h.972c.145 0 .264.116.264.257v6.486a.261.261 0 0 1-.264.257Zm-3 0h-.973A.258.258 0 0 1 3 17.75v-5c0-.137.119-.25.264-.25h.972c.145 0 .264.113.264.25v5c0 .137-.119.25-.264.25Z"/>`
	spaInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c1.5 0 4-1 4-5.5S12 6 12 6s-4 6-4 10.5s2.5 5.5 4 5.5Zm0 0c-1.5 0-2.953-.22-5-1.5C3 18 2.5 10 2.5 10s4.5.5 6.5 2m3 10c1.5 0 2.953-.22 5-1.5C21 18 21.5 10 21.5 10s-4.5.5-6.5 2"/>`
	spectrumInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="M22.131 23.399h-9.006c-.795 0-1.506-.62-1.503-1.417c.018-4.958-.878-6.584-2.753-8.224c-1.969-1.72-5.414-2.055-7.375-2.094a1.508 1.508 0 0 1-1.479-1.507L0 1.533A1.505 1.505 0 0 1 1.431.023c2.95-.133 9.632.183 15.09 4.956c4.434 3.875 6.824 9.541 7.118 16.859a1.507 1.507 0 0 1-1.508 1.56"/>`
	splitInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 22h22V2H1v20ZM12 2v20V2Z"/>`
	splitsInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 22h22V2H1v20ZM8 2v20V2Zm8 0v20V2Z"/>`
	spotifyInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M19.098 10.638C15.23 8.341 8.85 8.13 5.158 9.251a1.122 1.122 0 1 1-.652-2.148C8.745 5.816 15.79 6.064 20.244 8.708a1.122 1.122 0 1 1-1.146 1.93m-.126 3.403a.936.936 0 0 1-1.287.308c-3.225-1.983-8.142-2.557-11.958-1.399a.937.937 0 0 1-1.167-.623a.937.937 0 0 1 .624-1.167c4.358-1.322 9.776-.682 13.48 1.594c.44.271.578.847.308 1.287m-1.469 3.267a.748.748 0 0 1-1.028.25c-2.818-1.723-6.365-2.112-10.542-1.158a.748.748 0 1 1-.333-1.458c4.571-1.045 8.492-.595 11.655 1.338a.748.748 0 0 1 .248 1.028M12 0C5.373 0 0 5.372 0 12s5.373 12 12 12c6.628 0 12-5.372 12-12S18.628 0 12 0"/>`
	squareInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M14.444 8.333H9.556c-.675 0-1.223.548-1.223 1.223v4.888c0 .675.548 1.223 1.223 1.223h4.888c.675 0 1.223-.548 1.223-1.223V9.556c0-.675-.548-1.223-1.223-1.223M18.111 22H5.89A3.89 3.89 0 0 1 2 18.111V5.89A3.89 3.89 0 0 1 5.889 2H18.11A3.89 3.89 0 0 1 22 5.889V18.11A3.89 3.89 0 0 1 18.111 22Zm0 2A5.89 5.89 0 0 0 24 18.111V5.89A5.89 5.89 0 0 0 18.111 0H5.89A5.89 5.89 0 0 0 0 5.889V18.11A5.89 5.89 0 0 0 5.889 24H18.11Z"/>`
	stackOverflowInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="M11.414.132a1.406 1.406 0 0 1 1.188 0l10.82 5.055a.421.421 0 0 1-.002.764l-10.905 5.017a1.215 1.215 0 0 1-1.016 0L.595 5.951a.42.42 0 0 1-.002-.764L11.413.132Zm12.009 11.526a.42.42 0 0 1-.003.764l-10.904 5.017a1.215 1.215 0 0 1-1.016 0L.595 12.422a.42.42 0 0 1-.002-.764l2.235-1.044a1.216 1.216 0 0 1 1.023-.003l7.649 3.52c.322.148.694.148 1.016 0l7.649-3.52a1.22 1.22 0 0 1 1.023.003l2.235 1.044Zm0 6.5a.42.42 0 0 1-.003.764l-10.904 5.017a1.215 1.215 0 0 1-1.016 0L.595 18.922a.42.42 0 0 1-.002-.764l2.235-1.044a1.216 1.216 0 0 1 1.023-.003l7.649 3.52c.322.148.694.148 1.016 0l7.649-3.52a1.22 1.22 0 0 1 1.023.003l2.235 1.044Z"/>`
	stakeholderInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="m14 9l4.5 2L23 9V4l-4.5-2L14 4v5ZM7 7a4 4 0 1 0 0 8a4 4 0 0 0 0-8ZM1 23v-2c0-4 2.5-6 6-6s6 2 6 6v2H1ZM14 4l4.5 2L23 4m-4.5 2v5v-5Z"/>`
	starInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M12 16.667L5 22l3-8l-6-4.5h7.5L12 2l2.5 7.5H22L16 14l3 8z"/>`
	starHalfInnerSVG               = `<g fill="currentColor" fill-rule="evenodd"><path fill-opacity=".2" d="M12 16.667V2l2.5 7.5H22L16 14l3 8z"/><path d="M12 16.667L5 22l3-8l-6-4.5h7.5L12 2z"/></g>`
	statusCriticalInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12.703 2.703a.99.99 0 0 0-1.406 0l-8.594 8.594a.99.99 0 0 0 0 1.406l8.594 8.594a.99.99 0 0 0 1.406 0l8.594-8.594a.99.99 0 0 0 0-1.406l-8.594-8.594ZM8.983 14.7L14.7 8.983m-5.717 0L14.7 14.7"/>`
	statusCriticalSmallInnerSVG    = `<path fill="currentColor" fill-rule="evenodd" stroke="currentColor" d="M6.712 1.263a1.005 1.005 0 0 0-1.424 0L1.263 5.288a1.005 1.005 0 0 0 0 1.424l4.025 4.025a1.005 1.005 0 0 0 1.424 0l4.025-4.025a1.005 1.005 0 0 0 0-1.424L6.712 1.263Z"/>`
	statusDisabledInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 3.99C2 2.892 2.898 2 3.99 2h16.02C21.108 2 22 2.898 22 3.99v16.02c0 1.099-.898 1.99-1.99 1.99H3.99A1.995 1.995 0 0 1 2 20.01V3.99ZM18 12H6"/>`
	statusDisabledSmallInnerSVG    = `<rect width="10" height="10" x="1" y="1" fill="currentColor" fill-rule="evenodd" stroke="currentColor" rx="1"/>`
	statusGoodInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10ZM7 12l4 3l5-7"/>`
	statusGoodSmallInnerSVG        = `<circle cx="6" cy="6" r="5" fill="currentColor" fill-rule="evenodd" stroke="currentColor"/>`
	statusInfoInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 3.99C2 2.892 2.898 2 3.99 2h16.02C21.108 2 22 2.898 22 3.99v16.02c0 1.099-.898 1.99-1.99 1.99H3.99A1.995 1.995 0 0 1 2 20.01V3.99ZM12 10v8m0-12v2"/>`
	statusInfoSmallInnerSVG        = `<rect width="10" height="10" x="1" y="1" fill="currentColor" fill-rule="evenodd" stroke="currentColor" rx="1"/>`
	statusPlaceholderInnerSVG      = `<rect width="20" height="20" x="2" y="2" fill="none" stroke="currentColor" stroke-width="2" rx="2"/>`
	statusPlaceholderSmallInnerSVG = `<rect width="10" height="10" x="1" y="1" fill="currentColor" fill-rule="evenodd" stroke="currentColor" rx="1"/>`
	statusUnknownInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 3.99C2 2.892 2.898 2 3.99 2h16.02C21.108 2 22 2.898 22 3.99v16.02c0 1.099-.898 1.99-1.99 1.99H3.99A1.995 1.995 0 0 1 2 20.01V3.99ZM12 15v-1c0-1 0-1.5 1-2s2-1 2-2.5c0-1-1-2.5-3-2.5s-3 1.264-3 3m3 6v2"/>`
	statusUnknownSmallInnerSVG     = `<rect width="10" height="10" x="1" y="1" fill="currentColor" fill-rule="evenodd" stroke="currentColor" rx="1"/>`
	statusWarningInnerSVG          = `<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 3l10 18H2L12 3Zm0 6v6m0 1v2"/>`
	statusWarningSmallInnerSVG     = `<path fill="currentColor" fill-rule="evenodd" stroke="currentColor" stroke-linejoin="round" d="m6 1l5 9H1z"/>`
	stepsInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 6h-5v5H6v5H1v7h22V1h-7z"/>`
	stepsOptionInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M24 9h-5v5h-5v5H9v5m-8 0v-7.003c0-.55.313-1.31.703-1.7L15.297 1.703c.388-.388 1.156-.703 1.7-.703H24"/>`
	stopInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 4h16v16H4z"/>`
	stopFillInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 4h16v16H4V4Zm2 2h12v12H6V6Zm2 2h8v8H8V8Zm2 2h4v4h-4v-4Zm1 1h2v2h-2v-2Z"/>`
	storageInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 5.077S3.667 2 12 2s10 3.077 10 3.077v13.846S20.333 22 12 22S2 18.923 2 18.923V5.077ZM2 13s3.333 3 10 3s10-3 10-3M2 7s3.333 3 10 3s10-3 10-3"/>`
	streetViewInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 5a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-1 18v-6h3v-2c0-3.34-2.76-5.97-6-6c-3.21.03-6 2.66-6 6v2h3v6m-5.5 0h17h-17Z"/>`
	strikeThroughInnerSVG          = `<path fill="currentColor" d="M5.77 12.17h12.46v1.09H5.77zm10.93 1.48h-3.81c.59.34 1 .6 1.18.74a3.39 3.39 0 0 1 1 1.07a2.38 2.38 0 0 1 .31 1.14a2.3 2.3 0 0 1-.82 1.76a3.18 3.18 0 0 1-2.22.74a4.7 4.7 0 0 1-2.23-.54a3.77 3.77 0 0 1-1.55-1.36a7.41 7.41 0 0 1-.79-2.46h-.41V20h.41a1.35 1.35 0 0 1 .24-.7a.59.59 0 0 1 .44-.17a6.5 6.5 0 0 1 1.39.35a12.63 12.63 0 0 0 1.45.41a7.26 7.26 0 0 0 1.25.1A4.87 4.87 0 0 0 16 18.72a4 4 0 0 0 1.34-3a3.8 3.8 0 0 0-.42-1.72c-.06-.13-.14-.23-.22-.35Zm-7.81-2.44c.21.17.49.36.81.57h4.87c-.48-.29-1.07-.62-1.76-1a12.49 12.49 0 0 1-3.33-2.24A1.93 1.93 0 0 1 9 7.26a2.18 2.18 0 0 1 .77-1.63a2.72 2.72 0 0 1 1.93-.71a4.14 4.14 0 0 1 2 .53a3.78 3.78 0 0 1 1.49 1.43a6.6 6.6 0 0 1 .73 2.42h.41V4h-.41a1.39 1.39 0 0 1-.3.7a.67.67 0 0 1-.48.17a2.64 2.64 0 0 1-.89-.28A6.45 6.45 0 0 0 11.68 4a4.49 4.49 0 0 0-3.21 1.21A3.75 3.75 0 0 0 7.21 8a3.57 3.57 0 0 0 .43 1.73a4.72 4.72 0 0 0 1.25 1.48Z"/>`
	stripeInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M1 1h22v22H1V1Zm10.12 8.19c0-.604.494-.836 1.314-.836c1.174 0 2.658.356 3.833.99V5.71c-1.283-.509-2.55-.71-3.833-.71c-3.138 0-5.225 1.639-5.225 4.375c0 4.266 5.874 3.586 5.874 5.425c0 .711-.619.943-1.484.943c-1.283 0-2.922-.525-4.22-1.236v3.679c1.437.618 2.89.88 4.22.88c3.215 0 5.426-1.591 5.426-4.358c-.016-4.607-5.905-3.788-5.905-5.519Z"/>`
	subscriptInnerSVG              = `<path fill="currentColor" d="m18.74 3.63l.19 4.07h-.49a4.92 4.92 0 0 0-.38-1.54A2.58 2.58 0 0 0 17 5.07a3.68 3.68 0 0 0-1.73-.35h-2.44V18a3.54 3.54 0 0 0 .35 2a1.92 1.92 0 0 0 1.5.54h.6V21H7.92v-.5h.61a1.76 1.76 0 0 0 1.56-.67a3.88 3.88 0 0 0 .29-1.83V4.72H8.29a5.79 5.79 0 0 0-1.73.18a2.37 2.37 0 0 0-1.14.93a3.78 3.78 0 0 0-.56 1.87h-.48l.21-4.07Zm3.74 13.05h.33v5.59a1.13 1.13 0 0 0 .06.4a.41.41 0 0 0 .17.21a.63.63 0 0 0 .28.08h.4v.27h-2.96V23h.45a.62.62 0 0 0 .29-.1a.38.38 0 0 0 .15-.22a1.4 1.4 0 0 0 0-.37v-3.89a4.45 4.45 0 0 0 0-.64c0-.15-.1-.23-.24-.23a.82.82 0 0 0-.28 0l-.3.13l-.18-.25Z"/>`
	subtractInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 12h20"/>`
	subtractCircleInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10ZM6 12h12"/>`
	sunInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4V2m0 20v-2m8-8h2M2 12h2m13.657-5.657L19.07 4.93M4.93 19.07l1.414-1.414m0-11.314L4.93 4.93m14.14 14.14l-1.414-1.414M12 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10z"/>`
	superscriptInnerSVG            = `<path fill="currentColor" d="m18.74 3.63l.19 4.07h-.49a4.9 4.9 0 0 0-.38-1.54A2.57 2.57 0 0 0 17 5.07a3.68 3.68 0 0 0-1.73-.35h-2.44V18a3.56 3.56 0 0 0 .34 2a1.92 1.92 0 0 0 1.5.54h.6V21H7.92v-.5h.61a1.76 1.76 0 0 0 1.56-.67a3.88 3.88 0 0 0 .29-1.83V4.72H8.29a5.82 5.82 0 0 0-1.73.18a2.37 2.37 0 0 0-1.14.93a3.78 3.78 0 0 0-.56 1.87h-.48l.21-4.07ZM22.48 1h.33v5.62a1.13 1.13 0 0 0 .06.4a.41.41 0 0 0 .17.21a.63.63 0 0 0 .28.08h.4v.29h-2.96v-.27h.45a.62.62 0 0 0 .29-.1a.38.38 0 0 0 .15-.23a1.4 1.4 0 0 0 0-.37V2.77a4.45 4.45 0 0 0 0-.64c0-.15-.1-.23-.24-.23a.82.82 0 0 0-.28 0l-.3.13l-.18-.25Z"/>`
	supportInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4.222 19.778c4.296 4.296 11.26 4.296 15.556 0c4.296-4.296 4.296-11.26 0-15.556c-4.296-4.296-11.26-4.296-15.556 0c-4.296 4.296-4.296 11.26 0 15.556ZM6.343 21.9l4.243-4.242m-8.485 0l4.242-4.243m11.314-2.828l4.242-4.243m-8.485 0l4.243-4.242m-9.9 14.142a6 6 0 1 0 8.486-8.486a6 6 0 0 0-8.486 8.486Zm-5.656-9.9l4.242 4.243m0-8.485l4.243 4.242m2.828 11.314l4.243 4.242m0-8.485l4.242 4.243"/>`
	suseInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M24 8.285c-1.42-2.294-3.663-3.936-6.087-4.354c-2.037-.354-5.752-.391-7.72 3.322c-.822 1.55-.771 3.568.13 5.268c.919 1.74 2.561 2.859 4.504 3.072c1.928.21 3.376-.236 4.19-1.288c.82-1.063.972-2.509.39-3.687c-.624-1.265-1.567-1.937-2.875-2.048c-.999-.089-1.737.327-2.005.75a1.417 1.417 0 0 0-.218.737c0 .95.867 1.21 1.043 1.222c.054-.003.316-.008.662-.095l.168-.056l.2-.056a1.234 1.234 0 0 1 1.463 1.21c0 .408-.207.789-.631 1.071c-.182.115-.35.179-.525.234a4.356 4.356 0 0 1-1.373.236c-1.505-.002-3.6-1.295-3.626-3.724c-.014-1.371.673-2.618 1.885-3.422c1.394-.922 3.803-1.235 6.04.137c2.75 1.682 3.469 4.981 2.75 7.23c-1.043 3.258-3.98 5-7.853 4.653c-2.745-.244-5.295-1.693-6.812-3.874a9.615 9.615 0 0 1-1.44-3.119c-.62-2.456.006-5.066.472-6.222c.8-1.981 1.993-3.813 3.566-5.482H1.001A1 1 0 0 0 0 1v22a1 1 0 0 0 1 1h22a1 1 0 0 0 1-1V8.285Z"/>`
	swiftInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M18.103 21.018c-2.827 1.633-6.714 1.801-10.625.125c-3.166-1.347-5.794-3.705-7.478-6.4c.808.674 1.752 1.213 2.762 1.684c4.039 1.893 8.077 1.764 10.918.005l-.004-.005C9.634 13.328 6.198 9.286 3.638 5.985c-.54-.539-.943-1.212-1.348-1.819c3.1 2.83 8.018 6.4 9.769 7.411C8.354 7.67 5.053 2.82 5.187 2.954c5.861 5.928 11.319 9.297 11.319 9.297c.18.101.32.186.432.262c.118-.3.221-.612.308-.936c.944-3.436-.134-7.343-2.492-10.577c5.456 3.301 8.69 9.499 7.343 14.686c-.035.14-.074.279-.115.414l.048.058c2.694 3.369 1.953 6.94 1.616 6.266c-1.461-2.86-4.167-1.986-5.543-1.406Z"/>`
	swimInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 13c.5.5 2.13-.112 3.262-.5c1.46-.5 3.238 0 2.738-.5l-2-2s-4.5 2.5-4 3Zm-9 7c2 0 3-1 5-1s3 1 5 1s3-1 5-1s3 1 5 1M2 16c2 0 3-1 5-1s3 1 5 1s3-1 5-1s3 1 5 1M17.5 4l-5.278 3l3.278 3.5L12 12m7.222-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	switchInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M5 1h14v22H5V1Zm2.5 10H17v10H7V11h.5ZM15 6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-4 13v-5h2v5h-2Zm1-13.998a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z" clip-rule="evenodd"/>`
	syncInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 19h11a7 7 0 0 0 7-7V9M8 15l-4 4l4 4M19 5H8a7 7 0 0 0-7 7v3M16 1l4 4l-4 4"/>`
	systemInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 19h22V1H1v18Zm4 4h14H5Zm3 0h8v-4H8v4ZM7.757 5.757l2.122 2.122l-2.122-2.122ZM9 10H6h3Zm.879 2.121l-2.122 2.122l2.122-2.122ZM12 13v3v-3Zm2.121-.879l2.122 2.122l-2.122-2.122ZM18 10h-3h3Zm-1.757-4.243l-2.122 2.122l2.122-2.122ZM12 7V4v3Zm0 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/>`
	tableInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 5v18m8-18v18M1 11h22M1 5h22M1 17h22M1 1h22v22H1V1Z"/>`
	tableAddInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 5v18m8-18v6M1 11h16M1 5h22M1 17h10m6 6H1V1h22v16m-6 6a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0-9v6m0-6v6m-3-3h6"/>`
	tagInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M11.706 22.294a.992.992 0 0 1-1.41.003l-8.593-8.594a1 1 0 0 1 .003-1.409L13 1h10v10L11.706 22.294ZM6 12l6 6M9 9l6 6m2-9a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/>`
	tapeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 2h20v20H2V2Zm17 10a7 7 0 1 1-14 0a7 7 0 0 1 14 0Zm-7-3c-1.655 0-3 1.345-3 3s1.345 3 3 3s3-1.345 3-3s-1.345-3-3-3Z"/>`
	tapeOptionInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 2h21v20H2V7m10 12a7 7 0 0 1-7-7m7 7a7 7 0 0 0 0-14H1m11 4c-1.655 0-3 1.345-3 3s1.345 3 3 3s3-1.345 3-3s-1.345-3-3-3Z"/>`
	targetInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-5 0c0-3.309-2.691-6-6-6s-6 2.691-6 6s2.691 6 6 6s6-2.691 6-6Zm-5 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/>`
	taskInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 20h12m-12-8h12M12 4h12M1 19l3 3l5-5m-8-6l3 3l5-5m0-8L4 6L1 3"/>`
	tasksInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 3h22v4H1V3Zm0 7h22v4H1v-4Zm0 7h22v4H1v-4ZM1 4h15v2H1V4Zm0 7h5v2H1v-2Zm0 7h10v2H1v-2Z"/>`
	technologyInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.5 19a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM10 5l2-2m-4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm.5 6l8-8M5.5 21a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm13-13a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM12 21l2-2"/>`
	templateInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 3h22v18H1V3Zm0 5h22M7 8v13"/>`
	terminalInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="m2 5l6 6l-6 6m7 0h14"/>`
	testInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 1v7L2 20v3h20v-3L15 8V1m0 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-6 2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm9-7c-7-3-6 4-12 1M6 1h12"/>`
	testDesktopInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18.218 1H23v18H1V1h5m11 9c-4-3-6 3-10 0M5 23h14H5Zm5-22v4.773l-5 7.182V15h14v-2.045l-5-7.182V1M8 1h8h-8Zm0 22h8v-4H8v4Z"/>`
	textAlignCenterInnerSVG        = `<path fill="currentColor" d="M.46 3.06h23.08v2.18H.46zM4.1 8.29h15.81v2.18H4.1zM.46 13.53h23.08v2.18H.46zm3.64 5.23h15.81v2.18H4.1z"/>`
	textAlignFullInnerSVG          = `<path fill="currentColor" d="M.46 3.06h23.08v2.18H.46zm0 5.23h23.08v2.18H.46zm0 5.24h23.08v2.18H.46zm0 5.23h15.81v2.18H.46z"/>`
	textAlignLeftInnerSVG          = `<path fill="currentColor" d="M.46 3.06h23.08v2.18H.46zm0 5.23h15.81v2.18H.46zm0 5.24h23.08v2.18H.46zm0 5.23h15.81v2.18H.46z"/>`
	textAlignRightInnerSVG         = `<path fill="currentColor" d="M.46 3.06h23.08v2.18H.46zm7.27 5.23h15.81v2.18H7.73zM.46 13.53h23.08v2.18H.46zm7.27 5.23h15.81v2.18H7.73z"/>`
	textWrapInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M17 10h7h-7ZM1 14h13V2H1v12Zm5-8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m11 0h7h-7Zm0-4h7h-7Zm0 12h7h-7ZM0 18h24H0Zm0 4h24H0Zm14-8v-1l-4-5l-3 3l-1-1l-4 4h12Z"/>`
	threatsInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M9 23A7 7 0 1 0 9 9a7 7 0 0 0 0 14ZM9 6V5c0-3 2-4 4-4s4 1 4 4v3c0 1 0 3 2 3s2-2 2-3c0-2 1-2 2-2m-11 7l-6 6m0-6l6 6m-6-9V6h6v4"/>`
	threeDInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 5.5v13l10 4l10-4v-13l-10-4l-10 4ZM13 12h5v5M2 5.5l10 4l10-4M6 14s6 6.5 12-2"/>`
	threeDEffectsInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M11 3h2l9 3v11l-10 3.5L2 17V6.5L4 6l9 2.5l-2 .5l-9-2.5"/>`
	ticketInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M7 16h10V8H7v8Zm13-4c0 2 1 3 3 3v5H1v-5c2 0 3-1 3-3S3 9 1 9V4h22v5c-2 0-3 1-3 3Z"/>`
	tictokInnerSVG                 = `<g fill="currentColor"><path d="M22.459 6.846v3.659c-.197 0-.433.04-.669.04a7.295 7.295 0 0 1-4.682-1.732v7.79a6.987 6.987 0 0 1-1.416 4.25a7.02 7.02 0 0 1-5.626 2.832a6.993 6.993 0 0 1-5.941-3.305c1.259 1.18 2.95 1.928 4.8 1.928a6.893 6.893 0 0 0 5.586-2.833c.866-1.18 1.417-2.636 1.417-4.249v-7.83c1.259 1.102 2.872 1.732 4.682 1.732c.236 0 .433 0 .669-.04v-2.36c.354.079.669.118 1.023.118h.157z"/><path d="M11.05 9.56v4.053a3.277 3.277 0 0 0-.866-.118c-1.732 0-3.148 1.456-3.148 3.226c0 .394.079.748.197 1.102c-.787-.59-1.338-1.535-1.338-2.597c0-1.77 1.416-3.226 3.148-3.226c.314 0 .59.04.865.118V9.521h.236c.315 0 .63 0 .905.04zm6.648-5.626c-.708-.63-1.22-1.495-1.495-2.4h.945v.551a6.25 6.25 0 0 0 .55 1.85z"/><path d="M21.318 6.767v2.36c-.197.04-.433.04-.669.04a7.295 7.295 0 0 1-4.682-1.73v7.79a6.987 6.987 0 0 1-1.416 4.248c-1.299 1.732-3.305 2.833-5.587 2.833c-1.85 0-3.541-.747-4.8-1.928a7.136 7.136 0 0 1-1.062-3.737c0-3.817 3.03-6.925 6.806-7.043v2.597a3.277 3.277 0 0 0-.865-.118c-1.732 0-3.148 1.455-3.148 3.226c0 1.062.512 2.046 1.338 2.597c.433 1.22 1.613 2.124 2.95 2.124c1.732 0 3.148-1.456 3.148-3.226V1.534h2.872c.276.945.787 1.77 1.495 2.4a5.397 5.397 0 0 0 3.62 2.833z"/><path d="M9.908 8.184V9.52c-3.777.118-6.806 3.226-6.806 7.043c0 1.377.393 2.636 1.062 3.738A7.122 7.122 0 0 1 2 15.148c0-3.896 3.148-7.043 7.003-7.043c.315 0 .63.04.905.079z"/><path d="M16.203 1.534h-2.872v15.187c0 1.77-1.416 3.227-3.147 3.227c-1.377 0-2.518-.866-2.951-2.125c.511.354 1.14.59 1.81.59c1.73 0 3.147-1.416 3.147-3.187V0h3.817v.079c0 .157 0 .314.039.472c0 .315.079.669.157.983zm5.115 3.777v1.417c-1.574-.315-2.911-1.377-3.659-2.794a5.11 5.11 0 0 0 3.659 1.377z"/></g>`
	timeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 13h4l2.5-9l5 16.5L17 9l2 4h4"/>`
	tipInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16.007 18H22V2H2v16h6.243l3.882 4l3.882-4Z"/>`
	toastInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 2a4 4 0 0 0-2 7.465V16h12V9.465A4 4 0 0 0 16 2H8Zm3.321 4.874a1.004 1.004 0 0 1 1.38-.37l1.715.991c.483.279.652.889.37 1.38l-.991 1.715a1.004 1.004 0 0 1-1.38.37L10.7 9.968a1.004 1.004 0 0 1-.37-1.379l.991-1.716ZM8 18v2m4-2v5m4-5v3"/>`
	toolsInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="m11 2l11 11l-4.5 4.5l-11-11L11 2Zm5 4l1-1l2 2l-1 1m-5 5l-9 9l-2-2l9-9m-6 7l1 1"/>`
	tooltipInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16.5 18L12 22.5L7.5 18H1V1h22v17h-6.5ZM6 10h1V9H6v1Zm5.5 0h1V9h-1v1Zm5.5 0h1V9h-1v1Z"/>`
	topCornerInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 4H4v12"/>`
	trainInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 1h14a2 2 0 0 1 2 2v15a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2Zm4 1h6M3 5h18M4 23h16M3 12h18M7 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm10 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM12 5v7m-3 8l-2 3m8-3l2 3"/>`
	transactionInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M2 7h18m-4-5l5 5l-5 5m6 5H4m4-5l-5 5l5 5"/>`
	trashInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 5h16v18H4V5ZM1 5h22M9 1h6v4H9V1Zm0 0h6v4H9V1Zm6 8v10M9 9v10"/>`
	treeInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 1h6v6H4V1Zm12 10h4v4h-4v-4Zm0 8h4v4h-4v-4ZM7 7v14h9m-9-8h9"/>`
	treeOptionInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="m7 13l4.375-7H9l3-4l3 4h-2.375L17 13h-2l4 6.667H5L9 13H7Zm5 11v-4"/>`
	triggerInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M4 14h6l-3 9h2L20 9h-6l4-8H7z"/>`
	trophyInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 15a6 6 0 0 1-6-6V1h12v8a6 6 0 0 1-6 6ZM6 3H1v4c0 2.509 1.791 4 4 4h1V3Zm12 8h1c2.209 0 4-1.491 4-4V3h-5v8ZM5 23h14v-4H5v4Zm11-4a4 4 0 1 0-8 0"/>`
	troubleshootInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 5c0-2 1-4 2-4l2 4h2l2-4c1 0 2 2 2 4c0 2.254-1 4-3 5v11c0 1 0 2-2 2s-2-1-2-2V10C2 9 1 7.254 1 5Zm18 7v6m-2 0l1 5h2l1-5h-4Zm-3-6h10h-10Zm7 0V3a2 2 0 1 0-4 0v9"/>`
	ttyInnerSVG                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M7 19h10a1 1 0 0 1 1 1h0a1 1 0 0 1-1 1H7h0a1 1 0 0 1-1-1h0a1 1 0 0 1 1-1Zm0-9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm5-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm5-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	tumblrInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M17.639 19.17c-.446.212-1.3.398-1.937.415c-1.92.05-2.293-1.35-2.31-2.367v-7.47h4.82V6.114H13.41V0H9.894c-.057 0-.158.051-.172.18C9.516 2.051 8.64 5.335 5 6.647v3.1h2.428v7.842c0 2.685 1.981 6.499 7.21 6.41c1.763-.031 3.722-.77 4.155-1.406L17.64 19.17Z"/>`
	turbolinuxInnerSVG             = `<path fill="currentColor" fill-rule="evenodd" d="m9.419 6.222l.547 1.23h-3.35L6 6.223h3.419Zm3.692 5.949L7.094 0l7.042 4.17l.41 1.984h3.351l-.752 2.051h-2.188l1.778 8.274l-4.171-2.052L14.684 24L8.187 10.803l4.923 1.368Z"/>`
	twitterInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M24 4.309a9.83 9.83 0 0 1-2.828.775a4.94 4.94 0 0 0 2.165-2.724a9.865 9.865 0 0 1-3.127 1.196a4.925 4.925 0 0 0-8.39 4.49A13.974 13.974 0 0 1 1.671 2.9a4.902 4.902 0 0 0-.667 2.476c0 1.708.869 3.216 2.191 4.099A4.936 4.936 0 0 1 .964 8.86v.06a4.926 4.926 0 0 0 3.95 4.829a4.964 4.964 0 0 1-2.224.085a4.93 4.93 0 0 0 4.6 3.42a9.886 9.886 0 0 1-6.115 2.107c-.398 0-.79-.023-1.175-.068a13.945 13.945 0 0 0 7.548 2.212c9.057 0 14.009-7.503 14.009-14.01c0-.213-.005-.425-.014-.636A10.012 10.012 0 0 0 24 4.309"/>`
	ubuntuInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M24 12c0 6.627-5.373 12-12 12c-6.628 0-12-5.373-12-12S5.372 0 12 0c6.627 0 12 5.373 12 12ZM3.84 10.398a1.602 1.602 0 1 0 0 3.204a1.602 1.602 0 0 0 0-3.204Zm11.44 7.282a1.601 1.601 0 1 0 1.6 2.774a1.601 1.601 0 0 0-1.6-2.774ZM7.32 12c0-1.583.787-2.982 1.99-3.829L8.14 6.21a6.972 6.972 0 0 0-2.878 4.046c.506.413.829 1.041.829 1.745c0 .704-.323 1.332-.83 1.745A6.97 6.97 0 0 0 8.14 17.79l1.171-1.962A4.672 4.672 0 0 1 7.32 12ZM12 7.32a4.68 4.68 0 0 1 4.66 4.265l2.284-.033a6.938 6.938 0 0 0-2.068-4.515c-.61.23-1.313.195-1.92-.156a2.244 2.244 0 0 1-1.097-1.588a6.96 6.96 0 0 0-4.943.468l1.113 1.994A4.66 4.66 0 0 1 12 7.32Zm0 9.36a4.66 4.66 0 0 1-1.971-.435l-1.114 1.994a6.961 6.961 0 0 0 4.944.467a2.245 2.245 0 0 1 1.096-1.587a2.245 2.245 0 0 1 1.921-.156a6.938 6.938 0 0 0 2.068-4.515l-2.283-.033A4.679 4.679 0 0 1 12 16.68Zm3.279-10.36a1.601 1.601 0 1 0 1.602-2.773A1.601 1.601 0 0 0 15.28 6.32Z"/>`
	underlineInnerSVG              = `<path fill="currentColor" d="M14.41 4.53v-.35h4.66v.36h-.49a1.34 1.34 0 0 0-1.19.65a3 3 0 0 0-.2 1.4v5.33a9.45 9.45 0 0 1-.41 3.08a3.85 3.85 0 0 1-1.54 1.87a5.49 5.49 0 0 1-3.13.78a5.89 5.89 0 0 1-3.27-.75a4 4 0 0 1-1.58-2A11.14 11.14 0 0 1 7 11.64V6.5a2.58 2.58 0 0 0-.33-1.59a1.38 1.38 0 0 0-1.08-.38H5v-.35h5.68v.36h-.5A1.3 1.3 0 0 0 9.06 5a2.87 2.87 0 0 0-.25 1.5v5.73A12.52 12.52 0 0 0 9 14a3.71 3.71 0 0 0 .51 1.54a2.77 2.77 0 0 0 1.06.91a3.68 3.68 0 0 0 1.7.36a4.69 4.69 0 0 0 2.31-.56a3 3 0 0 0 1.39-1.44a8.33 8.33 0 0 0 .37-3V6.5A2.72 2.72 0 0 0 16 5a1.43 1.43 0 0 0-1.12-.43ZM4.93 20v-1H19v1Z"/>`
	undoInnerSVG                   = `<path fill="currentColor" d="M7.18 4L8.6 5.44L6.06 8h9.71a6 6 0 0 1 0 12h-2v-2h2a4 4 0 0 0 0-8H6.06l2.54 2.51l-1.42 1.41L2.23 9Z"/>`
	unlinkInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m17 11l4.586 4.586a1.998 1.998 0 0 1 0 2.828l-3.172 3.172a1.998 1.998 0 0 1-2.828 0L11 17m6-9h4m-5-1V3M8 21v-4m-5-1h4m0-3L2.414 8.414a1.998 1.998 0 0 1 0-2.828l3.172-3.172a1.998 1.998 0 0 1 2.828 0L13 7"/>`
	unlockInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M23 23V11H9v12h14Zm-9-6h4m-7-6V7c0-3 0-6-5-6S1 4 1 7v4"/>`
	unorderedListInnerSVG          = `<path fill="currentColor" d="M5.94 6.42H24v1.75H5.94zm0 5.29H24v1.75H5.94zm0 5.28H24v1.75H5.94z"/><circle cx="1.85" cy="7.29" r="1.52" fill="currentColor"/><circle cx="1.85" cy="12.58" r="1.52" fill="currentColor"/><circle cx="1.85" cy="17.87" r="1.52" fill="currentColor"/>`
	unsortedInnerSVG               = `<path fill="currentColor" fill-rule="evenodd" d="m15.204 15.321l1.597-1.597l.707.707l-2.45 2.45l-.354.354l-.353-.353l-2.45-2.45l.707-.708l1.596 1.597V7.217h1v8.104Zm-5.9-6.407v8.104h1V8.914l1.597 1.597l.707-.707l-2.45-2.45L9.803 7l-.354.354L7 9.804l.707.707l1.597-1.597Z" clip-rule="evenodd"/>`
	upInnerSVG                     = `<path fill="none" stroke="currentColor" stroke-width="2" d="m2 15.999l10.173-9.824l9.824 10.173"/>`
	updateInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1.75 16.002C3.353 20.098 7.338 23 12 23c6.075 0 11-4.925 11-11m-.75-4.002C20.649 3.901 16.663 1 12 1C5.925 1 1 5.925 1 12m8 4H1v8M23 0v8h-8"/>`
	upgradeInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 18V8v10Zm0 5c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm5-11l-5-5l-5 5"/>`
	uploadInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 17v6h22v-6M12 2v17M5 9l7-7l7 7"/>`
	uploadOptionInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="m17 12l-5-5l-5 5m5-4v10m0 5c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11z"/>`
	usbKeyInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h14v-1h7V7h-7V6H3Zm14 3v5h5V9h-5Zm-2 6V8H3a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h12Zm4-3.992h1.01v-1.01H19v1.01Zm.51 2h-.5v-1.01h1.01v1.01h-.51Z" clip-rule="evenodd"/>`
	userInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 24v-5m8 5v-5M3 24v-5c0-4.97 4.03-8 9-8s9 3.03 9 8v5m-9-13a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/>`
	userAddInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M5 24v-5m6 5v-5M1 24v-6c0-4.97 3.134-7 7-7s7 2 7 7v6M8 11A5 5 0 1 0 8 1a5 5 0 0 0 0 10Zm8 0h8m-4-4v8"/>`
	userAdminInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 11A5 5 0 1 0 8 1a5 5 0 0 0 0 10Zm5.023 2.023C11.772 11.76 10.013 11 8 11c-4 0-7 3-7 7v5h7m2-3.5a2.5 2.5 0 1 0 5.002-.002A2.5 2.5 0 0 0 10 19.5ZM23 15l-3-3l-6 6m3.5-3.5l3 3l-3-3Z"/>`
	userExpertInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M8 11A5 5 0 1 0 8 1a5 5 0 0 0 0 10Zm6.643 4.696a6.745 6.745 0 0 0-1.62-2.673C11.772 11.76 10.013 11 8 11c-4 0-7 3-7 7v5h10m1-4.176L16.19 22L23 13"/>`
	userFemaleInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20 24v-5c0-4-4.06-5-5-5c3.948 0 5-2 5-2s-3.057-1.969-3-6c-.057-3-2.15-5-5-5c-2.988 0-5.081 2-5 5c-.081 3.969-3 6-3 6s.914 2 5 2c-1.079 0-5 1-5 5v5m12-5v5m-8-5v5"/>`
	userManagerInnerSVG            = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 12c2.374 1.183 4 3.65 4 7v4H4v-4c0-3.354 1.631-5.825 4-7m4 1a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm6-6c-1.5 0-3 .36-5-2c-2 2.36-4.5 3-7 2m1 6l5.025 5.257L17 13m-5 5v5"/>`
	userNewInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 24V12m5 10l-10-7m10 0l-10 7M8 11A5 5 0 1 0 8 1a5 5 0 0 0 0 10Zm5.023 2.023C11.772 11.76 10.013 11 8 11c-4 0-7 3-7 7v5h10"/>`
	userPoliceInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M16 14c2.374 1.183 4 3.65 4 7v2H4v-2c0-3.354 1.631-5.825 4-7m4 1c3.26 0 5.903-2.686 5.903-5.999c0-.702.218-1.375 0-2.001M6.093 7c-.21.615 0 1.313 0 2.001C6.093 12.314 8.738 15 12 15M6 8h12l3-4c-1.912-1.735-5.21-3-9-3c-3.836 0-7.168 1.296-9 3l3 4Zm6-3a1 1 0 0 0 1-1h-2a1 1 0 0 0 1 1Z"/>`
	userSettingsInnerSVG           = `<path fill="none" stroke="currentColor" stroke-width="2" d="M18 21c-1.655 0-3-1.346-3-3s1.345-3 3-3c1.654 0 3 1.346 3 3s-1.346 3-3 3Zm6-3h-3h3Zm-3.879 2.122l2.121 2.12l-2.12-2.12ZM18.001 24v-3v3Zm-4.244-1.757l2.121-2.122l-2.12 2.122ZM12 18h3h-3Zm3.878-2.121l-2.12-2.121l2.12 2.12Zm2.122-.88v-3v3Zm2.121.88l2.121-2.121l-2.12 2.12ZM12.5 12.5C11.266 11.446 9.775 11 8 11c-3.866 0-7 2.03-7 7v5h10M8 11A5 5 0 1 0 8 1a5 5 0 0 0 0 10Z"/>`
	userWorkerInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 6h18H3Zm7-4v2m4-2v2m2 8c2.374 1.183 4 3.65 4 7v4H4v-4c0-3.354 1.631-5.825 4-7m4 4.5V23m0-10a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm-4-1a4 4 0 1 0 8 0"/>`
	validateInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M20 15c-1 1 1.25 3.75 0 5s-4-1-5 0s-1.5 3-3 3s-2-2-3-3s-3.75 1.25-5 0s1-4 0-5s-3-1.5-3-3s2-2 3-3s-1.25-3.75 0-5s4 1 5 0s1.5-3 3-3s2 2 3 3s3.75-1.25 5 0s-1 4 0 5s3 1.5 3 3s-2 2-3 3ZM7 12l3 3l7-7"/>`
	vendInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M3 2.002A.998.998 0 0 1 3.993 1h16.014c.548 0 .993.44.993 1.002V23H3V2.002ZM15 1v22M3 16h12m-7 1h2m-3-1v-4m0-3V5m4 11v-4m0-3V5m6 7h1m-1-3h2m-2 11h2M3 9h12"/>`
	videoInnerSVG                  = `<path fill="none" stroke="currentColor" stroke-width="2" d="m15 9l8-4v14l-8-4V9ZM1 5h14v14H1V5Z"/>`
	viewInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 21c-5 0-11-5-11-9s6-9 11-9s11 5 11 9s-6 9-11 9Zm0-14a5 5 0 1 0 0 10a5 5 0 0 0 0-10Z"/>`
	vimeoInnerSVG                  = `<path fill="currentColor" fill-rule="evenodd" d="M23.988 6.802c-.107 2.336-1.738 5.534-4.894 9.595c-3.264 4.24-6.024 6.36-8.282 6.36c-1.4 0-2.584-1.29-3.55-3.873c-.646-2.368-1.291-4.736-1.938-7.103c-.718-2.582-1.488-3.874-2.312-3.874c-.18 0-.808.378-1.884 1.13L0 7.583c1.184-1.04 2.352-2.08 3.502-3.123c1.58-1.364 2.767-2.082 3.556-2.155c1.868-.18 3.018 1.097 3.449 3.83c.466 2.948.79 4.782.97 5.5c.54 2.446 1.132 3.668 1.779 3.668c.501 0 1.257-.794 2.262-2.382c1.005-1.586 1.544-2.793 1.616-3.623c.144-1.369-.395-2.055-1.616-2.055c-.574 0-1.167.132-1.776.393c1.18-3.862 3.433-5.738 6.76-5.632c2.466.073 3.629 1.672 3.486 4.798"/>`
	virtualMachineInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 23h13V10H1v13Zm9-4h13V6H10v13Zm-5-5h13V1H5v13Z"/>`
	virtualStorageInnerSVG         = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c4.97 0 9-2.239 9-5s-4.03-5-9-5s-9 2.239-9 5s4.03 5 9 5Zm0-4.5c4.97 0 9-2.239 9-5s-4.03-5-9-5s-9 2.239-9 5s4.03 5 9 5Zm0-5.5c4.97 0 9-2.239 9-5s-4.03-5-9-5s-9 2.239-9 5s4.03 5 9 5Z"/>`
	visaInnerSVG                   = `<path fill="currentColor" fill-rule="evenodd" d="M5.756 6.342C4.344 5.56 2.733 4.93.931 4.494l.059-.35h7.407c.997.036 1.804.355 2.082 1.422l1.61 7.743v.001l.48 2.335l4.497-11.491h4.865L14.7 20.974l-4.86.005L5.756 6.342Z"/>`
	vmMaintenanceInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 10V2H7v12h7V7H2v12h8m4 4l6-6m1-3a2 2 0 1 0 2 2"/>`
	vmwareInnerSVG                 = `<path fill="currentColor" fill-rule="evenodd" d="M5.574 0c-.959 0-1.728.754-1.728 1.68v5.744H1.728C.77 7.424 0 8.18 0 9.104v12.438c0 .926.77 1.68 1.728 1.68H14.58c.96 0 1.753-.754 1.753-1.68v-2.045h5.939c.958 0 1.728-.757 1.728-1.68V5.38c0-.924-.77-1.655-1.728-1.655h-2.118V1.68c0-.926-.771-1.68-1.728-1.68H5.574Zm0 .95h12.852c.42 0 .755.323.755.73v2.044H9.42c-.957 0-1.753.731-1.753 1.655v2.045H4.82V1.68c0-.407.332-.73.755-.73Zm3.846 3.7h9.76v9.443a.759.759 0 0 1-.754.755h-2.093V9.103c0-.923-.794-1.68-1.753-1.68H8.641V5.38c0-.406.36-.73.779-.73Zm10.734 0h2.118c.42 0 .754.323.754.73v12.437c0 .406-.334.73-.754.73h-5.94v-2.75h2.094c.957 0 1.728-.778 1.728-1.704V4.65ZM1.728 8.372h2.118v5.72c0 .926.769 1.704 1.728 1.704h2.093v2.02c0 .923.796 1.68 1.753 1.68h5.939v2.045c0 .406-.356.73-.779.73H1.728a.735.735 0 0 1-.754-.73V9.103c0-.405.329-.73.754-.73Zm3.091 0h2.848v6.475H5.574a.757.757 0 0 1-.755-.755v-5.72Zm3.822 0h5.94c.422 0 .778.325.778.73v5.745H8.641V8.373Zm0 7.424h6.718v2.75h-5.94c-.419 0-.778-.324-.778-.73v-2.02Z"/>`
	volumeInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="M15 16a4 4 0 0 0 0-8m0 12c5 0 8-3.589 8-8s-3.589-8-8-8M1 12V8h5l6-5v18l-6-5H1v-4"/>`
	volumeControlInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12ZM8 8l3 3m1 11a9.99 9.99 0 0 0 8.307-4.43A9.953 9.953 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12"/>`
	volumeLowInnerSVG              = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 8v8h5.099L12 21V3L6 8H1Zm14 8a4 4 0 1 0 0-8"/>`
	volumeMuteInnerSVG             = `<path fill="none" stroke="currentColor" stroke-width="2" d="M1 8v8h5.099L12 21V3L6 8H1Zm14 1l6 6m0-6l-6 6"/>`
	vulnerabilityInnerSVG          = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 0v24V0ZM0 12h24H0Zm17 0c0-2.757-2.243-5-5-5s-5 2.243-5 5s2.243 5 5 5s5-2.243 5-5Zm-5 9c-4.962 0-9-4.037-9-9s4.038-9 9-9s9 4.037 9 9s-4.038 9-9 9Z"/>`
	waypointInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="m3 11l8 2l2 8l8-18z"/>`
	webcamInnerSVG                 = `<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M20 22H4"/><path d="M15.5 18v3m-7.5.5V18"/><path d="M12 19a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z" clip-rule="evenodd"/><path d="M12 16a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z" clip-rule="evenodd"/><path d="M12 14a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/><path d="M12 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/></g>`
	wheelchairInnerSVG             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v9h7l2 6h2m-6-9H9a6 6 0 1 0 6 6M11 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`
	wheelchairActiveInnerSVG       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 6l3-3l7 3v2l-3 3M9 22a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm5-5l5 3l-2 6m2-17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-6 6l4-4m-6 4l4-4"/>`
	wifiInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-width="2" d="M12 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-4.243-6.243a6 6 0 0 1 8.486 0M4.929 10.93c3.905-3.905 10.237-3.905 14.142 0M2.101 8.1c5.467-5.468 14.331-5.468 19.798 0"/>`
	wifiLowInnerSVG                = `<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-width="2"><path d="M12 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-4.243-6.243a6 6 0 0 1 8.486 0"/><path stroke-opacity=".2" d="M4.929 10.929c3.905-3.905 10.237-3.905 14.142 0M2.101 8.1c5.467-5.468 14.331-5.468 19.798 0" opacity=".8"/></g>`
	wifiMediumInnerSVG             = `<g fill="none" stroke="currentColor" stroke-width="2"><path d="M12 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-4.243-6.243a6 6 0 0 1 8.486 0M4.929 10.93c3.905-3.905 10.237-3.905 14.142 0"/><path stroke-opacity=".2" d="M2.1 8.1c5.468-5.467 14.332-5.467 19.8 0" opacity=".8"/></g>`
	wifiNoneInnerSVG               = `<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="18" r="2"/><path stroke-opacity=".2" d="M7.757 13.757a6 6 0 0 1 8.486 0M4.929 10.93c3.905-3.905 10.237-3.905 14.142 0M2.101 8.1c5.467-5.468 14.331-5.468 19.798 0" opacity=".8"/></g>`
	windowsInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M23.923 0L10.959 1.893v9.588l12.964-.102V0ZM0 3.398l.009 8.155l9.773-.055l-.004-9.432L0 3.398Zm.008 17.283l9.773 1.344l-.008-9.44l-9.766-.063l.001 8.159Zm10.951 1.49L23.923 24l.004-11.326l-12.986-.022l.018 9.519Z"/>`
	windowsLegacyInnerSVG          = `<path fill="currentColor" fill-rule="evenodd" d="m2.613 10.096l.282-.985c.666-2.301 1.332-4.603 1.996-6.904c.026-.09.054-.164.157-.205c1.15-.461 2.327-.825 3.565-.959c1.666-.18 3.195.21 4.607 1.094c.184.115.362.24.55.347c.12.069.136.142.098.272c-.553 1.897-1.1 3.795-1.648 5.694c-.192.661-.388 1.322-.574 1.986c-.04.148-.075.173-.211.08c-.814-.553-1.672-1.017-2.633-1.263c-1.212-.312-2.423-.25-3.632.015c-.862.19-1.692.474-2.557.828m14.664 2.929c-1.529.01-2.878-.526-4.134-1.351c-.355-.234-.354-.232-.238-.635c.712-2.458 1.421-4.918 2.134-7.376c.07-.241.019-.285.29-.104c.895.6 1.839 1.094 2.908 1.31c1.128.228 2.248.15 3.362-.099c.7-.157 1.38-.377 2.048-.635c.096-.036.195-.098.294-.007c.097.09.052.19.022.293c-.713 2.465-1.426 4.93-2.134 7.397c-.033.115-.093.171-.2.214c-1.162.465-2.351.831-3.604.95c-.248.025-.498.03-.748.043m4.1.327l-.442 1.533c-.602 2.08-1.206 4.162-1.801 6.245a.437.437 0 0 1-.298.32c-1.026.395-2.069.72-3.163.874c-1.693.238-3.268-.082-4.73-.963c-.242-.146-.475-.306-.72-.45c-.12-.07-.13-.144-.093-.272c.523-1.796 1.04-3.594 1.56-5.392l.635-2.194c.067-.23.068-.23.28-.09c.76.501 1.551.936 2.434 1.186c1.273.36 2.547.3 3.822.018c.849-.187 1.668-.472 2.516-.815m-10.198-1.305l-.366 1.261c-.623 2.156-1.249 4.312-1.866 6.47c-.054.185-.103.19-.253.088c-.857-.585-1.764-1.065-2.787-1.298c-1.157-.264-2.309-.193-3.458.059c-.72.157-1.418.384-2.106.649c-.09.034-.183.085-.277.008c-.105-.086-.06-.191-.03-.291l2.132-7.377a.282.282 0 0 1 .185-.203c1.18-.475 2.389-.844 3.663-.962c1.571-.144 3.026.212 4.372 1.029c.23.14.457.285.684.43c.058.038.138.07.107.137"/>`
	wordpressInnerSVG              = `<path fill="currentColor" fill-rule="evenodd" d="M0 12a12 12 0 0 0 6.763 10.799L1.039 7.116A11.958 11.958 0 0 0 0 11.999m20.1-.605c0-1.483-.533-2.51-.99-3.309c-.607-.99-1.178-1.826-1.178-2.815c0-1.103.836-2.13 2.015-2.13c.053 0 .104.007.155.01A11.954 11.954 0 0 0 12 0C7.807 0 4.12 2.151 1.973 5.408c.282.01.548.014.773.014c1.255 0 3.198-.152 3.198-.152c.646-.038.723.913.076.989c0 0-.65.076-1.373.114l4.37 12.998l2.626-7.876l-1.869-5.121a22.087 22.087 0 0 1-1.26-.115c-.646-.038-.57-1.027.077-.989c0 0 1.982.153 3.16.153c1.256 0 3.199-.153 3.199-.153c.647-.038.723.913.076.989c0 0-.65.076-1.373.114l4.337 12.899l1.197-3.999c.518-1.66.913-2.852.913-3.879m-7.89 1.655L8.61 23.51a11.988 11.988 0 0 0 7.374-.192a1.14 1.14 0 0 1-.085-.165L12.21 13.05Zm10.32-6.807c.051.383.08.792.08 1.234c0 1.217-.228 2.586-.912 4.298l-3.666 10.597A11.993 11.993 0 0 0 24 12c0-2.086-.533-4.047-1.47-5.757"/>`
	workshopInnerSVG               = `<path fill="none" stroke="currentColor" stroke-width="2" d="M19 7s-5 7-12.5 7c-2 0-5.5 1-5.5 5v4h11v-4c0-2.5 3-1 7-8l-1.5-1.5M3 5V2h20v14h-3M11 1h4v2h-4V1ZM6.5 14a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Z"/>`
	yogaInnerSVG                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 11l-1 3l1 3h-1.5L9 14l.5-4.5L12 11Zm1-9L9 6v8l1 3H6l-3 5m17.5 0l-5-3.5L12 17l-1-3l1-3l3.5 2v5.5M14 8.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-3 2L10 17v-3.5l1-3Z"/>`
	youtubeInnerSVG                = `<path fill="currentColor" fill-rule="evenodd" d="M9.522 15.553V8.81l6.484 3.383l-6.484 3.36ZM23.76 7.641s-.235-1.654-.954-2.382c-.913-.956-1.936-.96-2.405-1.016C17.043 4 12.005 4 12.005 4h-.01s-5.038 0-8.396.243c-.47.055-1.492.06-2.406 1.016C.474 5.987.24 7.641.24 7.641S0 9.584 0 11.525v1.822c0 1.942.24 3.884.24 3.884s.234 1.653.953 2.382c.914.956 2.113.926 2.647 1.026c1.92.184 8.16.241 8.16.241s5.043-.007 8.401-.25c.47-.056 1.492-.061 2.405-1.017c.72-.729.954-2.382.954-2.382s.24-1.942.24-3.885v-1.82c0-1.942-.24-3.885-.24-3.885Z"/>`
	zoomInnerSVG                   = `<g fill="currentColor"><path d="M0 8a8 8 0 0 1 8-8h8a8 8 0 0 1 8 8v8a8 8 0 0 1-8 8H8a8 8 0 0 1-8-8V8Z"/><path d="M5 9a1 1 0 0 1 1-1h6a3 3 0 0 1 3 3v4a1 1 0 0 1-1 1H8a3 3 0 0 1-3-3V9Zm10.5 2.752a2 2 0 0 1 .495-1.318l1.69-1.932c.457-.52 1.315-.198 1.315.494v6.008c0 .693-.858 1.015-1.314.494l-1.691-1.932a2 2 0 0 1-.495-1.317v-.498Z"/></g>`
	zoomInInnerSVG                 = `<path fill="none" stroke="currentColor" stroke-width="2" d="m16 16l7 7l-7-7Zm-6 2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0-3V5m-5 5h10"/>`
	zoomOutInnerSVG                = `<path fill="none" stroke="currentColor" stroke-width="2" d="m16 16l7 7l-7-7Zm-6 2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm-5-8h10"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Accessibility(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accessibilityInnerSVG).
		Element(children...)
}

func Achievement(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(achievementInnerSVG).
		Element(children...)
}

func Action(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(actionInnerSVG).
		Element(children...)
}

func Actions(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(actionsInnerSVG).
		Element(children...)
}

func Ad(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(adInnerSVG).
		Element(children...)
}

func Add(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addInnerSVG).
		Element(children...)
}

func AddCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(addCircleInnerSVG).
		Element(children...)
}

func Aed(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aedInnerSVG).
		Element(children...)
}

func Aggregate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aggregateInnerSVG).
		Element(children...)
}

func Aid(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aidInnerSVG).
		Element(children...)
}

func AidOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(aidOptionInnerSVG).
		Element(children...)
}

func Alarm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alarmInnerSVG).
		Element(children...)
}

func Alert(children ...*engine.UberElement) *engine.UberElement {
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

func Amazon(children ...*engine.UberElement) *engine.UberElement {
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

func Amex(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(amexInnerSVG).
		Element(children...)
}

func Analytics(children ...*engine.UberElement) *engine.UberElement {
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

func Anchor(children ...*engine.UberElement) *engine.UberElement {
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

func Android(children ...*engine.UberElement) *engine.UberElement {
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

func Announce(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(announceInnerSVG).
		Element(children...)
}

func Apple(children ...*engine.UberElement) *engine.UberElement {
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

func AppleAppStore(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appleAppStoreInnerSVG).
		Element(children...)
}

func Apps(children ...*engine.UberElement) *engine.UberElement {
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

func AppsRounded(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(appsRoundedInnerSVG).
		Element(children...)
}

func Archive(children ...*engine.UberElement) *engine.UberElement {
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

func Archlinux(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archlinuxInnerSVG).
		Element(children...)
}

func Article(children ...*engine.UberElement) *engine.UberElement {
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

func Aruba(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arubaInnerSVG).
		Element(children...)
}

func Ascend(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ascendInnerSVG).
		Element(children...)
}

func Ascending(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ascendingInnerSVG).
		Element(children...)
}

func AssistListening(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(assistListeningInnerSVG).
		Element(children...)
}

func Atm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(atmInnerSVG).
		Element(children...)
}

func Attachment(children ...*engine.UberElement) *engine.UberElement {
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

func Attraction(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(attractionInnerSVG).
		Element(children...)
}

func Baby(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(babyInnerSVG).
		Element(children...)
}

func BackTen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(backTenInnerSVG).
		Element(children...)
}

func Bar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barInnerSVG).
		Element(children...)
}

func BarChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(barChartInnerSVG).
		Element(children...)
}

func Basket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(basketInnerSVG).
		Element(children...)
}

func Beacon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(beaconInnerSVG).
		Element(children...)
}

func Bike(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bikeInnerSVG).
		Element(children...)
}

func Bitcoin(children ...*engine.UberElement) *engine.UberElement {
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

func BladesHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bladesHorizontalInnerSVG).
		Element(children...)
}

func BladesVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bladesVerticalInnerSVG).
		Element(children...)
}

func BlockQuote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(blockQuoteInnerSVG).
		Element(children...)
}

func Blog(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(blogInnerSVG).
		Element(children...)
}

func Bluetooth(children ...*engine.UberElement) *engine.UberElement {
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

func Bold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boldInnerSVG).
		Element(children...)
}

func Book(children ...*engine.UberElement) *engine.UberElement {
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

func Bookmark(children ...*engine.UberElement) *engine.UberElement {
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

func BottomCorner(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bottomCornerInnerSVG).
		Element(children...)
}

func Braille(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brailleInnerSVG).
		Element(children...)
}

func Briefcase(children ...*engine.UberElement) *engine.UberElement {
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

func Brush(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(brushInnerSVG).
		Element(children...)
}

func Bug(children ...*engine.UberElement) *engine.UberElement {
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

func Bundle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bundleInnerSVG).
		Element(children...)
}

func Bus(children ...*engine.UberElement) *engine.UberElement {
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

func BusinessService(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(businessServiceInnerSVG).
		Element(children...)
}

func Cafeteria(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cafeteriaInnerSVG).
		Element(children...)
}

func Calculator(children ...*engine.UberElement) *engine.UberElement {
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

func Calendar(children ...*engine.UberElement) *engine.UberElement {
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

func Camera(children ...*engine.UberElement) *engine.UberElement {
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

func Capacity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(capacityInnerSVG).
		Element(children...)
}

func Car(children ...*engine.UberElement) *engine.UberElement {
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

func CaretDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretDownInnerSVG).
		Element(children...)
}

func CaretDownFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretDownFillInnerSVG).
		Element(children...)
}

func CaretLeftFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretLeftFillInnerSVG).
		Element(children...)
}

func CaretNext(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretNextInnerSVG).
		Element(children...)
}

func CaretPrevious(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretPreviousInnerSVG).
		Element(children...)
}

func CaretRightFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretRightFillInnerSVG).
		Element(children...)
}

func CaretUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretUpInnerSVG).
		Element(children...)
}

func CaretUpFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(caretUpFillInnerSVG).
		Element(children...)
}

func Cart(children ...*engine.UberElement) *engine.UberElement {
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

func Catalog(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(catalogInnerSVG).
		Element(children...)
}

func CatalogOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(catalogOptionInnerSVG).
		Element(children...)
}

func Centos(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(centosInnerSVG).
		Element(children...)
}

func Certificate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(certificateInnerSVG).
		Element(children...)
}

func Channel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(channelInnerSVG).
		Element(children...)
}

func ChapterAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chapterAddInnerSVG).
		Element(children...)
}

func ChapterNext(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chapterNextInnerSVG).
		Element(children...)
}

func ChapterPrevious(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chapterPreviousInnerSVG).
		Element(children...)
}

func Chat(children ...*engine.UberElement) *engine.UberElement {
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

func ChatOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatOptionInnerSVG).
		Element(children...)
}

func Checkbox(children ...*engine.UberElement) *engine.UberElement {
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

func CheckboxSelected(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxSelectedInnerSVG).
		Element(children...)
}

func Checkmark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkmarkInnerSVG).
		Element(children...)
}

func Chrome(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chromeInnerSVG).
		Element(children...)
}

func CircleAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleAlertInnerSVG).
		Element(children...)
}

func CircleInformation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleInformationInnerSVG).
		Element(children...)
}

func CirclePlay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circlePlayInnerSVG).
		Element(children...)
}

func CircleQuestion(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleQuestionInnerSVG).
		Element(children...)
}

func Clear(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clearInnerSVG).
		Element(children...)
}

func ClearOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clearOptionInnerSVG).
		Element(children...)
}

func Cli(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cliInnerSVG).
		Element(children...)
}

func Clipboard(children ...*engine.UberElement) *engine.UberElement {
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

func Clock(children ...*engine.UberElement) *engine.UberElement {
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

func Clone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloneInnerSVG).
		Element(children...)
}

func Close(children ...*engine.UberElement) *engine.UberElement {
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

func ClosedCaption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(closedCaptionInnerSVG).
		Element(children...)
}

func Cloud(children ...*engine.UberElement) *engine.UberElement {
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

func CloudComputer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudComputerInnerSVG).
		Element(children...)
}

func CloudDownload(children ...*engine.UberElement) *engine.UberElement {
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

func CloudSoftware(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudSoftwareInnerSVG).
		Element(children...)
}

func CloudUpload(children ...*engine.UberElement) *engine.UberElement {
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

func Cloudlinux(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cloudlinuxInnerSVG).
		Element(children...)
}

func Cluster(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clusterInnerSVG).
		Element(children...)
}

func CoatCheck(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coatCheckInnerSVG).
		Element(children...)
}

func Code(children ...*engine.UberElement) *engine.UberElement {
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

func CodeSandbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codeSandboxInnerSVG).
		Element(children...)
}

func Codepen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(codepenInnerSVG).
		Element(children...)
}

func Coffee(children ...*engine.UberElement) *engine.UberElement {
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

func Columns(children ...*engine.UberElement) *engine.UberElement {
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

func Command(children ...*engine.UberElement) *engine.UberElement {
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

func Compare(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compareInnerSVG).
		Element(children...)
}

func Compass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassInnerSVG).
		Element(children...)
}

func Compliance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(complianceInnerSVG).
		Element(children...)
}

func Configure(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(configureInnerSVG).
		Element(children...)
}

func Connect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(connectInnerSVG).
		Element(children...)
}

func Connectivity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(connectivityInnerSVG).
		Element(children...)
}

func Console(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(consoleInnerSVG).
		Element(children...)
}

func Contact(children ...*engine.UberElement) *engine.UberElement {
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

func ContactInfo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contactInfoInnerSVG).
		Element(children...)
}

func Contract(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(contractInnerSVG).
		Element(children...)
}

func Copy(children ...*engine.UberElement) *engine.UberElement {
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

func Cpu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cpuInnerSVG).
		Element(children...)
}

func CreativeCommons(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creativeCommonsInnerSVG).
		Element(children...)
}

func CreditCard(children ...*engine.UberElement) *engine.UberElement {
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

func CssThree(children ...*engine.UberElement) *engine.UberElement {
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

func Cube(children ...*engine.UberElement) *engine.UberElement {
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

func Cubes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cubesInnerSVG).
		Element(children...)
}

func Currency(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(currencyInnerSVG).
		Element(children...)
}

func Cursor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cursorInnerSVG).
		Element(children...)
}

func Cut(children ...*engine.UberElement) *engine.UberElement {
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

func Cycle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cycleInnerSVG).
		Element(children...)
}

func Dashboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dashboardInnerSVG).
		Element(children...)
}

func Database(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseInnerSVG).
		Element(children...)
}

func Debian(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(debianInnerSVG).
		Element(children...)
}

func Deliver(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deliverInnerSVG).
		Element(children...)
}

func Deploy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deployInnerSVG).
		Element(children...)
}

func Descend(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(descendInnerSVG).
		Element(children...)
}

func Descending(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(descendingInnerSVG).
		Element(children...)
}

func Desktop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(desktopInnerSVG).
		Element(children...)
}

func Detach(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(detachInnerSVG).
		Element(children...)
}

func Device(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(deviceInnerSVG).
		Element(children...)
}

func Diamond(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
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
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(directionsInnerSVG).
		Element(children...)
}

func DisabledOutline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(disabledOutlineInnerSVG).
		Element(children...)
}

func Disc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(discInnerSVG).
		Element(children...)
}

func Dislike(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dislikeInnerSVG).
		Element(children...)
}

func Docker(children ...*engine.UberElement) *engine.UberElement {
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

func Document(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentInnerSVG).
		Element(children...)
}

func DocumentCloud(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentCloudInnerSVG).
		Element(children...)
}

func DocumentConfig(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentConfigInnerSVG).
		Element(children...)
}

func DocumentCsv(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentCsvInnerSVG).
		Element(children...)
}

func DocumentDownload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentDownloadInnerSVG).
		Element(children...)
}

func DocumentExcel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentExcelInnerSVG).
		Element(children...)
}

func DocumentImage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentImageInnerSVG).
		Element(children...)
}

func DocumentLocked(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentLockedInnerSVG).
		Element(children...)
}

func DocumentMissing(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentMissingInnerSVG).
		Element(children...)
}

func DocumentNotes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentNotesInnerSVG).
		Element(children...)
}

func DocumentOutlook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentOutlookInnerSVG).
		Element(children...)
}

func DocumentPdf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentPdfInnerSVG).
		Element(children...)
}

func DocumentPerformance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentPerformanceInnerSVG).
		Element(children...)
}

func DocumentPpt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentPptInnerSVG).
		Element(children...)
}

func DocumentRtf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentRtfInnerSVG).
		Element(children...)
}

func DocumentSound(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentSoundInnerSVG).
		Element(children...)
}

func DocumentStore(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentStoreInnerSVG).
		Element(children...)
}

func DocumentTest(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentTestInnerSVG).
		Element(children...)
}

func DocumentText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentTextInnerSVG).
		Element(children...)
}

func DocumentThreat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentThreatInnerSVG).
		Element(children...)
}

func DocumentTime(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentTimeInnerSVG).
		Element(children...)
}

func DocumentTransfer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentTransferInnerSVG).
		Element(children...)
}

func DocumentTxt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentTxtInnerSVG).
		Element(children...)
}

func DocumentUpdate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentUpdateInnerSVG).
		Element(children...)
}

func DocumentUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentUploadInnerSVG).
		Element(children...)
}

func DocumentUser(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentUserInnerSVG).
		Element(children...)
}

func DocumentVerified(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentVerifiedInnerSVG).
		Element(children...)
}

func DocumentVideo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentVideoInnerSVG).
		Element(children...)
}

func DocumentWindows(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentWindowsInnerSVG).
		Element(children...)
}

func DocumentWord(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentWordInnerSVG).
		Element(children...)
}

func DocumentZip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(documentZipInnerSVG).
		Element(children...)
}

func Domain(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(domainInnerSVG).
		Element(children...)
}

func Dos(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dosInnerSVG).
		Element(children...)
}

func Down(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downInnerSVG).
		Element(children...)
}

func Download(children ...*engine.UberElement) *engine.UberElement {
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

func DownloadOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadOptionInnerSVG).
		Element(children...)
}

func Drag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dragInnerSVG).
		Element(children...)
}

func Drawer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(drawerInnerSVG).
		Element(children...)
}

func Dribbble(children ...*engine.UberElement) *engine.UberElement {
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

func DriveCage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(driveCageInnerSVG).
		Element(children...)
}

func Dropbox(children ...*engine.UberElement) *engine.UberElement {
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

func Duplicate(children ...*engine.UberElement) *engine.UberElement {
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

func Dxc(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(dxcInnerSVG).
		Element(children...)
}

func Edge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(edgeInnerSVG).
		Element(children...)
}

func Edit(children ...*engine.UberElement) *engine.UberElement {
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

func Eject(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ejectInnerSVG).
		Element(children...)
}

func Elevator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(elevatorInnerSVG).
		Element(children...)
}

func Emergency(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emergencyInnerSVG).
		Element(children...)
}

func Emoji(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emojiInnerSVG).
		Element(children...)
}

func EmptyCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emptyCircleInnerSVG).
		Element(children...)
}

func Erase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(eraseInnerSVG).
		Element(children...)
}

func Escalator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(escalatorInnerSVG).
		Element(children...)
}

func Expand(children ...*engine.UberElement) *engine.UberElement {
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

func Ezmeral(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ezmeralInnerSVG).
		Element(children...)
}

func Facebook(children ...*engine.UberElement) *engine.UberElement {
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

func FacebookOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(facebookOptionInnerSVG).
		Element(children...)
}

func Fan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fanInnerSVG).
		Element(children...)
}

func FanOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fanOptionInnerSVG).
		Element(children...)
}

func FastForward(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fastForwardInnerSVG).
		Element(children...)
}

func Favorite(children ...*engine.UberElement) *engine.UberElement {
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

func Fedora(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fedoraInnerSVG).
		Element(children...)
}

func Figma(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(figmaInnerSVG).
		Element(children...)
}

func Filter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(filterInnerSVG).
		Element(children...)
}

func FingerPrint(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fingerPrintInnerSVG).
		Element(children...)
}

func Fireball(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fireballInnerSVG).
		Element(children...)
}

func Firefox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(firefoxInnerSVG).
		Element(children...)
}

func Firewall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(firewallInnerSVG).
		Element(children...)
}

func Flag(children ...*engine.UberElement) *engine.UberElement {
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

func FlagFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flagFillInnerSVG).
		Element(children...)
}

func Flows(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flowsInnerSVG).
		Element(children...)
}

func Folder(children ...*engine.UberElement) *engine.UberElement {
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

func FolderCycle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderCycleInnerSVG).
		Element(children...)
}

func FolderOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderOpenInnerSVG).
		Element(children...)
}

func FormAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formAddInnerSVG).
		Element(children...)
}

func FormAttachment(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formAttachmentInnerSVG).
		Element(children...)
}

func FormCalendar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formCalendarInnerSVG).
		Element(children...)
}

func FormCheckmark(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formCheckmarkInnerSVG).
		Element(children...)
}

func FormClock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formClockInnerSVG).
		Element(children...)
}

func FormClose(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formCloseInnerSVG).
		Element(children...)
}

func FormCut(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formCutInnerSVG).
		Element(children...)
}

func FormDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formDownInnerSVG).
		Element(children...)
}

func FormEdit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formEditInnerSVG).
		Element(children...)
}

func FormFilter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formFilterInnerSVG).
		Element(children...)
}

func FormFolder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formFolderInnerSVG).
		Element(children...)
}

func FormLocation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formLocationInnerSVG).
		Element(children...)
}

func FormLock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formLockInnerSVG).
		Element(children...)
}

func FormNext(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formNextInnerSVG).
		Element(children...)
}

func FormNextLink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formNextLinkInnerSVG).
		Element(children...)
}

func FormPrevious(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formPreviousInnerSVG).
		Element(children...)
}

func FormPreviousLink(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formPreviousLinkInnerSVG).
		Element(children...)
}

func FormRefresh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formRefreshInnerSVG).
		Element(children...)
}

func FormSchedule(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formScheduleInnerSVG).
		Element(children...)
}

func FormSearch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formSearchInnerSVG).
		Element(children...)
}

func FormSubtract(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formSubtractInnerSVG).
		Element(children...)
}

func FormTrash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formTrashInnerSVG).
		Element(children...)
}

func FormUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formUpInnerSVG).
		Element(children...)
}

func FormUpload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formUploadInnerSVG).
		Element(children...)
}

func FormView(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formViewInnerSVG).
		Element(children...)
}

func FormViewHide(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(formViewHideInnerSVG).
		Element(children...)
}

func ForwardTen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(forwardTenInnerSVG).
		Element(children...)
}

func Freebsd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(freebsdInnerSVG).
		Element(children...)
}

func Gallery(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(galleryInnerSVG).
		Element(children...)
}

func Gamepad(children ...*engine.UberElement) *engine.UberElement {
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

func Gateway(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gatewayInnerSVG).
		Element(children...)
}

func Gatsbyjs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gatsbyjsInnerSVG).
		Element(children...)
}

func Gem(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
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
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(giftInnerSVG).
		Element(children...)
}

func Github(children ...*engine.UberElement) *engine.UberElement {
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

func Globe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(globeInnerSVG).
		Element(children...)
}

func Golang(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(golangInnerSVG).
		Element(children...)
}

func Google(children ...*engine.UberElement) *engine.UberElement {
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

func GooglePlay(children ...*engine.UberElement) *engine.UberElement {
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

func GooglePlus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googlePlusInnerSVG).
		Element(children...)
}

func GoogleWallet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(googleWalletInnerSVG).
		Element(children...)
}

func GraphQl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(graphQlInnerSVG).
		Element(children...)
}

func Gremlin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gremlinInnerSVG).
		Element(children...)
}

func Grid(children ...*engine.UberElement) *engine.UberElement {
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

func Grommet(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(grommetInnerSVG).
		Element(children...)
}

func Group(children ...*engine.UberElement) *engine.UberElement {
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

func Grow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(growInnerSVG).
		Element(children...)
}

func Hadoop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hadoopInnerSVG).
		Element(children...)
}

func Halt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(haltInnerSVG).
		Element(children...)
}

func Help(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(helpInnerSVG).
		Element(children...)
}

func HelpOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(helpOptionInnerSVG).
		Element(children...)
}

func Heroku(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(herokuInnerSVG).
		Element(children...)
}

func Hide(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hideInnerSVG).
		Element(children...)
}

func History(children ...*engine.UberElement) *engine.UberElement {
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

func Home(children ...*engine.UberElement) *engine.UberElement {
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

func HomeOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeOptionInnerSVG).
		Element(children...)
}

func HomeRounded(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(homeRoundedInnerSVG).
		Element(children...)
}

func Horton(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hortonInnerSVG).
		Element(children...)
}

func Host(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hostInnerSVG).
		Element(children...)
}

func HostMaintenance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hostMaintenanceInnerSVG).
		Element(children...)
}

func Hp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hpInnerSVG).
		Element(children...)
}

func Hpe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hpeInnerSVG).
		Element(children...)
}

func HpeLabs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hpeLabsInnerSVG).
		Element(children...)
}

func Hpi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(hpiInnerSVG).
		Element(children...)
}

func HtmlFive(children ...*engine.UberElement) *engine.UberElement {
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

func IceCream(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iceCreamInnerSVG).
		Element(children...)
}

func Image(children ...*engine.UberElement) *engine.UberElement {
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

func Impact(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(impactInnerSVG).
		Element(children...)
}

func InProgress(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inProgressInnerSVG).
		Element(children...)
}

func Inbox(children ...*engine.UberElement) *engine.UberElement {
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

func Indicator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(indicatorInnerSVG).
		Element(children...)
}

func Info(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(infoInnerSVG).
		Element(children...)
}

func Inherit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inheritInnerSVG).
		Element(children...)
}

func Insecure(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(insecureInnerSVG).
		Element(children...)
}

func Inspect(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(inspectInnerSVG).
		Element(children...)
}

func Instagram(children ...*engine.UberElement) *engine.UberElement {
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

func Install(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(installInnerSVG).
		Element(children...)
}

func InstallOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(installOptionInnerSVG).
		Element(children...)
}

func Integration(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(integrationInnerSVG).
		Element(children...)
}

func InternetExplorer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(internetExplorerInnerSVG).
		Element(children...)
}

func Italic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(italicInnerSVG).
		Element(children...)
}

func Iteration(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(iterationInnerSVG).
		Element(children...)
}

func Java(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(javaInnerSVG).
		Element(children...)
}

func Js(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(jsInnerSVG).
		Element(children...)
}

func Key(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(keyInnerSVG).
		Element(children...)
}

func Keyboard(children ...*engine.UberElement) *engine.UberElement {
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

func Language(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(languageInnerSVG).
		Element(children...)
}

func Lastfm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lastfmInnerSVG).
		Element(children...)
}

func Launch(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(launchInnerSVG).
		Element(children...)
}

func Layer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(layerInnerSVG).
		Element(children...)
}

func License(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(licenseInnerSVG).
		Element(children...)
}

func Like(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(likeInnerSVG).
		Element(children...)
}

func LineChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lineChartInnerSVG).
		Element(children...)
}

func Link(children ...*engine.UberElement) *engine.UberElement {
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

func LinkBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkBottomInnerSVG).
		Element(children...)
}

func LinkDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkDownInnerSVG).
		Element(children...)
}

func LinkNext(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkNextInnerSVG).
		Element(children...)
}

func LinkPrevious(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkPreviousInnerSVG).
		Element(children...)
}

func LinkTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkTopInnerSVG).
		Element(children...)
}

func LinkUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkUpInnerSVG).
		Element(children...)
}

func Linkedin(children ...*engine.UberElement) *engine.UberElement {
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

func LinkedinOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(linkedinOptionInnerSVG).
		Element(children...)
}

func List(children ...*engine.UberElement) *engine.UberElement {
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

func Local(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(localInnerSVG).
		Element(children...)
}

func Location(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationInnerSVG).
		Element(children...)
}

func LocationPin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(locationPinInnerSVG).
		Element(children...)
}

func Lock(children ...*engine.UberElement) *engine.UberElement {
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

func Login(children ...*engine.UberElement) *engine.UberElement {
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

func Logout(children ...*engine.UberElement) *engine.UberElement {
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

func Lounge(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(loungeInnerSVG).
		Element(children...)
}

func Magic(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(magicInnerSVG).
		Element(children...)
}

func Mail(children ...*engine.UberElement) *engine.UberElement {
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

func MailOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mailOptionInnerSVG).
		Element(children...)
}

func Mandriva(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mandrivaInnerSVG).
		Element(children...)
}

func Manual(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(manualInnerSVG).
		Element(children...)
}

func Map(children ...*engine.UberElement) *engine.UberElement {
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

func MapLocation(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapLocationInnerSVG).
		Element(children...)
}

func Mastercard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mastercardInnerSVG).
		Element(children...)
}

func Medium(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mediumInnerSVG).
		Element(children...)
}

func Memory(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(memoryInnerSVG).
		Element(children...)
}

func Menu(children ...*engine.UberElement) *engine.UberElement {
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

func Microfocus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microfocusInnerSVG).
		Element(children...)
}

func Microphone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(microphoneInnerSVG).
		Element(children...)
}

func Money(children ...*engine.UberElement) *engine.UberElement {
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

func Monitor(children ...*engine.UberElement) *engine.UberElement {
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

func Monospace(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(monospaceInnerSVG).
		Element(children...)
}

func Moon(children ...*engine.UberElement) *engine.UberElement {
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

func More(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(moreInnerSVG).
		Element(children...)
}

func MoreVertical(children ...*engine.UberElement) *engine.UberElement {
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

func Mouse(children ...*engine.UberElement) *engine.UberElement {
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

func Multimedia(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multimediaInnerSVG).
		Element(children...)
}

func Multiple(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multipleInnerSVG).
		Element(children...)
}

func Music(children ...*engine.UberElement) *engine.UberElement {
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

func Mysql(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mysqlInnerSVG).
		Element(children...)
}

func Navigate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(navigateInnerSVG).
		Element(children...)
}

func Network(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(networkInnerSVG).
		Element(children...)
}

func NetworkDrive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(networkDriveInnerSVG).
		Element(children...)
}

func New(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(newInnerSVG).
		Element(children...)
}

func NewWindow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(newWindowInnerSVG).
		Element(children...)
}

func Next(children ...*engine.UberElement) *engine.UberElement {
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

func Node(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nodeInnerSVG).
		Element(children...)
}

func Nodes(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nodesInnerSVG).
		Element(children...)
}

func Norton(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(nortonInnerSVG).
		Element(children...)
}

func Note(children ...*engine.UberElement) *engine.UberElement {
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

func Notes(children ...*engine.UberElement) *engine.UberElement {
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

func Notification(children ...*engine.UberElement) *engine.UberElement {
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

func Npm(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(npmInnerSVG).
		Element(children...)
}

func ObjectGroup(children ...*engine.UberElement) *engine.UberElement {
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

func ObjectUngroup(children ...*engine.UberElement) *engine.UberElement {
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

func OfflineStorage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(offlineStorageInnerSVG).
		Element(children...)
}

func Onedrive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(onedriveInnerSVG).
		Element(children...)
}

func Opera(children ...*engine.UberElement) *engine.UberElement {
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

func Optimize(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(optimizeInnerSVG).
		Element(children...)
}

func Oracle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(oracleInnerSVG).
		Element(children...)
}

func OrderedList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(orderedListInnerSVG).
		Element(children...)
}

func Organization(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(organizationInnerSVG).
		Element(children...)
}

func Overview(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(overviewInnerSVG).
		Element(children...)
}

func Package(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(packageInnerSVG).
		Element(children...)
}

func Paint(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(paintInnerSVG).
		Element(children...)
}

func Pan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(panInnerSVG).
		Element(children...)
}

func Pause(children ...*engine.UberElement) *engine.UberElement {
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

func PauseFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseFillInnerSVG).
		Element(children...)
}

func Paypal(children ...*engine.UberElement) *engine.UberElement {
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

func Performance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(performanceInnerSVG).
		Element(children...)
}

func PersonalComputer(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(personalComputerInnerSVG).
		Element(children...)
}

func Phone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneInnerSVG).
		Element(children...)
}

func PhoneFlip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneFlipInnerSVG).
		Element(children...)
}

func PhoneHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneHorizontalInnerSVG).
		Element(children...)
}

func PhoneVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(phoneVerticalInnerSVG).
		Element(children...)
}

func PieChart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pieChartInnerSVG).
		Element(children...)
}

func PiedPiper(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(piedPiperInnerSVG).
		Element(children...)
}

func Pin(children ...*engine.UberElement) *engine.UberElement {
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

func Pinterest(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pinterestInnerSVG).
		Element(children...)
}

func Plan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(planInnerSVG).
		Element(children...)
}

func Play(children ...*engine.UberElement) *engine.UberElement {
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

func PlayFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(playFillInnerSVG).
		Element(children...)
}

func Plug(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plugInnerSVG).
		Element(children...)
}

func Pocket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pocketInnerSVG).
		Element(children...)
}

func Power(children ...*engine.UberElement) *engine.UberElement {
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

func PowerCycle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerCycleInnerSVG).
		Element(children...)
}

func PowerForceShutdown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerForceShutdownInnerSVG).
		Element(children...)
}

func PowerReset(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerResetInnerSVG).
		Element(children...)
}

func PowerShutdown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(powerShutdownInnerSVG).
		Element(children...)
}

func Previous(children ...*engine.UberElement) *engine.UberElement {
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

func Print(children ...*engine.UberElement) *engine.UberElement {
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

func ProductHunt(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(productHuntInnerSVG).
		Element(children...)
}

func Projects(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(projectsInnerSVG).
		Element(children...)
}

func Qr(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(qrInnerSVG).
		Element(children...)
}

func Radial(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radialInnerSVG).
		Element(children...)
}

func RadialSelected(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radialSelectedInnerSVG).
		Element(children...)
}

func Raspberry(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(raspberryInnerSVG).
		Element(children...)
}

func Reactjs(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(reactjsInnerSVG).
		Element(children...)
}

func Reddit(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redditInnerSVG).
		Element(children...)
}

func Redhat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(redhatInnerSVG).
		Element(children...)
}

func Redo(children ...*engine.UberElement) *engine.UberElement {
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

func Refresh(children ...*engine.UberElement) *engine.UberElement {
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

func Resources(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resourcesInnerSVG).
		Element(children...)
}

func Restaurant(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(restaurantInnerSVG).
		Element(children...)
}

func Restroom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(restroomInnerSVG).
		Element(children...)
}

func RestroomMen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(restroomMenInnerSVG).
		Element(children...)
}

func RestroomWomen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(restroomWomenInnerSVG).
		Element(children...)
}

func Resume(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(resumeInnerSVG).
		Element(children...)
}

func Return(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(returnInnerSVG).
		Element(children...)
}

func Revert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(revertInnerSVG).
		Element(children...)
}

func Rewind(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rewindInnerSVG).
		Element(children...)
}

func Risk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(riskInnerSVG).
		Element(children...)
}

func Robot(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(robotInnerSVG).
		Element(children...)
}

func RotateLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateLeftInnerSVG).
		Element(children...)
}

func RotateRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateRightInnerSVG).
		Element(children...)
}

func Rss(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rssInnerSVG).
		Element(children...)
}

func Run(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(runInnerSVG).
		Element(children...)
}

func SafariOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(safariOptionInnerSVG).
		Element(children...)
}

func Sans(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sansInnerSVG).
		Element(children...)
}

func Satellite(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(satelliteInnerSVG).
		Element(children...)
}

func Save(children ...*engine.UberElement) *engine.UberElement {
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

func Scan(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scanInnerSVG).
		Element(children...)
}

func Schedule(children ...*engine.UberElement) *engine.UberElement {
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

func ScheduleNew(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scheduleNewInnerSVG).
		Element(children...)
}

func SchedulePlay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(schedulePlayInnerSVG).
		Element(children...)
}

func Schedules(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(schedulesInnerSVG).
		Element(children...)
}

func Sco(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scoInnerSVG).
		Element(children...)
}

func Scorecard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scorecardInnerSVG).
		Element(children...)
}

func Script(children ...*engine.UberElement) *engine.UberElement {
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

func Sd(children ...*engine.UberElement) *engine.UberElement {
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

func Search(children ...*engine.UberElement) *engine.UberElement {
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

func SearchAdvanced(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(searchAdvancedInnerSVG).
		Element(children...)
}

func Secure(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(secureInnerSVG).
		Element(children...)
}

func Select(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(selectInnerSVG).
		Element(children...)
}

func Selection(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(selectionInnerSVG).
		Element(children...)
}

func Semantics(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(semanticsInnerSVG).
		Element(children...)
}

func Send(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sendInnerSVG).
		Element(children...)
}

func Server(children ...*engine.UberElement) *engine.UberElement {
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

func ServerCluster(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(serverClusterInnerSVG).
		Element(children...)
}

func Servers(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(serversInnerSVG).
		Element(children...)
}

func ServicePlay(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(servicePlayInnerSVG).
		Element(children...)
}

func Services(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(servicesInnerSVG).
		Element(children...)
}

func SettingsOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(settingsOptionInnerSVG).
		Element(children...)
}

func Share(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareInnerSVG).
		Element(children...)
}

func ShareOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareOptionInnerSVG).
		Element(children...)
}

func ShareRounded(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shareRoundedInnerSVG).
		Element(children...)
}

func Shield(children ...*engine.UberElement) *engine.UberElement {
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

func ShieldSecurity(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldSecurityInnerSVG).
		Element(children...)
}

func Shift(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shiftInnerSVG).
		Element(children...)
}

func Shop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shopInnerSVG).
		Element(children...)
}

func Sidebar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(sidebarInnerSVG).
		Element(children...)
}

func Sign(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(signInnerSVG).
		Element(children...)
}

func Skype(children ...*engine.UberElement) *engine.UberElement {
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

func Slack(children ...*engine.UberElement) *engine.UberElement {
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

func Snapchat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(snapchatInnerSVG).
		Element(children...)
}

func Solaris(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(solarisInnerSVG).
		Element(children...)
}

func Sort(children ...*engine.UberElement) *engine.UberElement {
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

func Soundcloud(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(soundcloudInnerSVG).
		Element(children...)
}

func Spa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spaInnerSVG).
		Element(children...)
}

func Spectrum(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spectrumInnerSVG).
		Element(children...)
}

func Split(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splitInnerSVG).
		Element(children...)
}

func Splits(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(splitsInnerSVG).
		Element(children...)
}

func Spotify(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(spotifyInnerSVG).
		Element(children...)
}

func Square(children ...*engine.UberElement) *engine.UberElement {
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

func StackOverflow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stackOverflowInnerSVG).
		Element(children...)
}

func Stakeholder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stakeholderInnerSVG).
		Element(children...)
}

func Star(children ...*engine.UberElement) *engine.UberElement {
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

func StarHalf(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(starHalfInnerSVG).
		Element(children...)
}

func StatusCritical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusCriticalInnerSVG).
		Element(children...)
}

func StatusCriticalSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusCriticalSmallInnerSVG).
		Element(children...)
}

func StatusDisabled(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusDisabledInnerSVG).
		Element(children...)
}

func StatusDisabledSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusDisabledSmallInnerSVG).
		Element(children...)
}

func StatusGood(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusGoodInnerSVG).
		Element(children...)
}

func StatusGoodSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusGoodSmallInnerSVG).
		Element(children...)
}

func StatusInfo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusInfoInnerSVG).
		Element(children...)
}

func StatusInfoSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusInfoSmallInnerSVG).
		Element(children...)
}

func StatusPlaceholder(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusPlaceholderInnerSVG).
		Element(children...)
}

func StatusPlaceholderSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusPlaceholderSmallInnerSVG).
		Element(children...)
}

func StatusUnknown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusUnknownInnerSVG).
		Element(children...)
}

func StatusUnknownSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusUnknownSmallInnerSVG).
		Element(children...)
}

func StatusWarning(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusWarningInnerSVG).
		Element(children...)
}

func StatusWarningSmall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(statusWarningSmallInnerSVG).
		Element(children...)
}

func Steps(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stepsInnerSVG).
		Element(children...)
}

func StepsOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stepsOptionInnerSVG).
		Element(children...)
}

func Stop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopInnerSVG).
		Element(children...)
}

func StopFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopFillInnerSVG).
		Element(children...)
}

func Storage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(storageInnerSVG).
		Element(children...)
}

func StreetView(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(streetViewInnerSVG).
		Element(children...)
}

func StrikeThrough(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(strikeThroughInnerSVG).
		Element(children...)
}

func Stripe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stripeInnerSVG).
		Element(children...)
}

func Subscript(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subscriptInnerSVG).
		Element(children...)
}

func Subtract(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subtractInnerSVG).
		Element(children...)
}

func SubtractCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(subtractCircleInnerSVG).
		Element(children...)
}

func Sun(children ...*engine.UberElement) *engine.UberElement {
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

func Superscript(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(superscriptInnerSVG).
		Element(children...)
}

func Support(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(supportInnerSVG).
		Element(children...)
}

func Suse(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(suseInnerSVG).
		Element(children...)
}

func Swift(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swiftInnerSVG).
		Element(children...)
}

func Swim(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swimInnerSVG).
		Element(children...)
}

func Switch(children ...*engine.UberElement) *engine.UberElement {
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

func Sync(children ...*engine.UberElement) *engine.UberElement {
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

func System(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(systemInnerSVG).
		Element(children...)
}

func Table(children ...*engine.UberElement) *engine.UberElement {
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

func TableAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tableAddInnerSVG).
		Element(children...)
}

func Tag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagInnerSVG).
		Element(children...)
}

func Tape(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tapeInnerSVG).
		Element(children...)
}

func TapeOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tapeOptionInnerSVG).
		Element(children...)
}

func Target(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(targetInnerSVG).
		Element(children...)
}

func Task(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(taskInnerSVG).
		Element(children...)
}

func Tasks(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tasksInnerSVG).
		Element(children...)
}

func Technology(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(technologyInnerSVG).
		Element(children...)
}

func Template(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(templateInnerSVG).
		Element(children...)
}

func Terminal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalInnerSVG).
		Element(children...)
}

func Test(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(testInnerSVG).
		Element(children...)
}

func TestDesktop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(testDesktopInnerSVG).
		Element(children...)
}

func TextAlignCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignCenterInnerSVG).
		Element(children...)
}

func TextAlignFull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignFullInnerSVG).
		Element(children...)
}

func TextAlignLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignLeftInnerSVG).
		Element(children...)
}

func TextAlignRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textAlignRightInnerSVG).
		Element(children...)
}

func TextWrap(children ...*engine.UberElement) *engine.UberElement {
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

func Threats(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threatsInnerSVG).
		Element(children...)
}

func ThreeD(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDInnerSVG).
		Element(children...)
}

func ThreeDEffects(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(threeDEffectsInnerSVG).
		Element(children...)
}

func Ticket(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ticketInnerSVG).
		Element(children...)
}

func Tictok(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tictokInnerSVG).
		Element(children...)
}

func Time(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(timeInnerSVG).
		Element(children...)
}

func Tip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tipInnerSVG).
		Element(children...)
}

func Toast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toastInnerSVG).
		Element(children...)
}

func Tools(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toolsInnerSVG).
		Element(children...)
}

func Tooltip(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipInnerSVG).
		Element(children...)
}

func TopCorner(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(topCornerInnerSVG).
		Element(children...)
}

func Train(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trainInnerSVG).
		Element(children...)
}

func Transaction(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(transactionInnerSVG).
		Element(children...)
}

func Trash(children ...*engine.UberElement) *engine.UberElement {
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

func Tree(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(treeInnerSVG).
		Element(children...)
}

func TreeOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(treeOptionInnerSVG).
		Element(children...)
}

func Trigger(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(triggerInnerSVG).
		Element(children...)
}

func Trophy(children ...*engine.UberElement) *engine.UberElement {
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

func Troubleshoot(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(troubleshootInnerSVG).
		Element(children...)
}

func Tty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ttyInnerSVG).
		Element(children...)
}

func Tumblr(children ...*engine.UberElement) *engine.UberElement {
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

func Turbolinux(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(turbolinuxInnerSVG).
		Element(children...)
}

func Twitter(children ...*engine.UberElement) *engine.UberElement {
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

func Ubuntu(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ubuntuInnerSVG).
		Element(children...)
}

func Underline(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(underlineInnerSVG).
		Element(children...)
}

func Undo(children ...*engine.UberElement) *engine.UberElement {
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

func Unlink(children ...*engine.UberElement) *engine.UberElement {
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

func Unlock(children ...*engine.UberElement) *engine.UberElement {
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

func UnorderedList(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unorderedListInnerSVG).
		Element(children...)
}

func Unsorted(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(unsortedInnerSVG).
		Element(children...)
}

func Up(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upInnerSVG).
		Element(children...)
}

func Update(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(updateInnerSVG).
		Element(children...)
}

func Upgrade(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(upgradeInnerSVG).
		Element(children...)
}

func Upload(children ...*engine.UberElement) *engine.UberElement {
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

func UploadOption(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadOptionInnerSVG).
		Element(children...)
}

func UsbKey(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(usbKeyInnerSVG).
		Element(children...)
}

func User(children ...*engine.UberElement) *engine.UberElement {
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

func UserAdd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userAddInnerSVG).
		Element(children...)
}

func UserAdmin(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userAdminInnerSVG).
		Element(children...)
}

func UserExpert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userExpertInnerSVG).
		Element(children...)
}

func UserFemale(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userFemaleInnerSVG).
		Element(children...)
}

func UserManager(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userManagerInnerSVG).
		Element(children...)
}

func UserNew(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userNewInnerSVG).
		Element(children...)
}

func UserPolice(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userPoliceInnerSVG).
		Element(children...)
}

func UserSettings(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userSettingsInnerSVG).
		Element(children...)
}

func UserWorker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(userWorkerInnerSVG).
		Element(children...)
}

func Validate(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(validateInnerSVG).
		Element(children...)
}

func Vend(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vendInnerSVG).
		Element(children...)
}

func Video(children ...*engine.UberElement) *engine.UberElement {
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

func View(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(viewInnerSVG).
		Element(children...)
}

func Vimeo(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vimeoInnerSVG).
		Element(children...)
}

func VirtualMachine(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(virtualMachineInnerSVG).
		Element(children...)
}

func VirtualStorage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(virtualStorageInnerSVG).
		Element(children...)
}

func Visa(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(visaInnerSVG).
		Element(children...)
}

func VmMaintenance(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vmMaintenanceInnerSVG).
		Element(children...)
}

func Vmware(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vmwareInnerSVG).
		Element(children...)
}

func Volume(children ...*engine.UberElement) *engine.UberElement {
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

func VolumeControl(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeControlInnerSVG).
		Element(children...)
}

func VolumeLow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeLowInnerSVG).
		Element(children...)
}

func VolumeMute(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeMuteInnerSVG).
		Element(children...)
}

func Vulnerability(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(vulnerabilityInnerSVG).
		Element(children...)
}

func Waypoint(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(waypointInnerSVG).
		Element(children...)
}

func Webcam(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(webcamInnerSVG).
		Element(children...)
}

func Wheelchair(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wheelchairInnerSVG).
		Element(children...)
}

func WheelchairActive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wheelchairActiveInnerSVG).
		Element(children...)
}

func Wifi(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiInnerSVG).
		Element(children...)
}

func WifiLow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiLowInnerSVG).
		Element(children...)
}

func WifiMedium(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiMediumInnerSVG).
		Element(children...)
}

func WifiNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wifiNoneInnerSVG).
		Element(children...)
}

func Windows(children ...*engine.UberElement) *engine.UberElement {
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

func WindowsLegacy(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(windowsLegacyInnerSVG).
		Element(children...)
}

func Wordpress(children ...*engine.UberElement) *engine.UberElement {
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

func Workshop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(workshopInnerSVG).
		Element(children...)
}

func Yoga(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(yogaInnerSVG).
		Element(children...)
}

func Youtube(children ...*engine.UberElement) *engine.UberElement {
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

func Zoom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 24 24"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(zoomInnerSVG).
		Element(children...)
}

func ZoomIn(children ...*engine.UberElement) *engine.UberElement {
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

func ZoomOut(children ...*engine.UberElement) *engine.UberElement {
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
	case "accessibility":
		return Accessibility(), nil
	case "achievement":
		return Achievement(), nil
	case "action":
		return Action(), nil
	case "actions":
		return Actions(), nil
	case "ad":
		return Ad(), nil
	case "add":
		return Add(), nil
	case "add-circle":
		return AddCircle(), nil
	case "aed":
		return Aed(), nil
	case "aggregate":
		return Aggregate(), nil
	case "aid":
		return Aid(), nil
	case "aid-option":
		return AidOption(), nil
	case "alarm":
		return Alarm(), nil
	case "alert":
		return Alert(), nil
	case "amazon":
		return Amazon(), nil
	case "amex":
		return Amex(), nil
	case "analytics":
		return Analytics(), nil
	case "anchor":
		return Anchor(), nil
	case "android":
		return Android(), nil
	case "announce":
		return Announce(), nil
	case "apple":
		return Apple(), nil
	case "apple-app-store":
		return AppleAppStore(), nil
	case "apps":
		return Apps(), nil
	case "apps-rounded":
		return AppsRounded(), nil
	case "archive":
		return Archive(), nil
	case "archlinux":
		return Archlinux(), nil
	case "article":
		return Article(), nil
	case "aruba":
		return Aruba(), nil
	case "ascend":
		return Ascend(), nil
	case "ascending":
		return Ascending(), nil
	case "assist-listening":
		return AssistListening(), nil
	case "atm":
		return Atm(), nil
	case "attachment":
		return Attachment(), nil
	case "attraction":
		return Attraction(), nil
	case "baby":
		return Baby(), nil
	case "back-ten":
		return BackTen(), nil
	case "bar":
		return Bar(), nil
	case "bar-chart":
		return BarChart(), nil
	case "basket":
		return Basket(), nil
	case "beacon":
		return Beacon(), nil
	case "bike":
		return Bike(), nil
	case "bitcoin":
		return Bitcoin(), nil
	case "blades-horizontal":
		return BladesHorizontal(), nil
	case "blades-vertical":
		return BladesVertical(), nil
	case "block-quote":
		return BlockQuote(), nil
	case "blog":
		return Blog(), nil
	case "bluetooth":
		return Bluetooth(), nil
	case "bold":
		return Bold(), nil
	case "book":
		return Book(), nil
	case "bookmark":
		return Bookmark(), nil
	case "bottom-corner":
		return BottomCorner(), nil
	case "braille":
		return Braille(), nil
	case "briefcase":
		return Briefcase(), nil
	case "brush":
		return Brush(), nil
	case "bug":
		return Bug(), nil
	case "bundle":
		return Bundle(), nil
	case "bus":
		return Bus(), nil
	case "business-service":
		return BusinessService(), nil
	case "cafeteria":
		return Cafeteria(), nil
	case "calculator":
		return Calculator(), nil
	case "calendar":
		return Calendar(), nil
	case "camera":
		return Camera(), nil
	case "capacity":
		return Capacity(), nil
	case "car":
		return Car(), nil
	case "caret-down":
		return CaretDown(), nil
	case "caret-down-fill":
		return CaretDownFill(), nil
	case "caret-left-fill":
		return CaretLeftFill(), nil
	case "caret-next":
		return CaretNext(), nil
	case "caret-previous":
		return CaretPrevious(), nil
	case "caret-right-fill":
		return CaretRightFill(), nil
	case "caret-up":
		return CaretUp(), nil
	case "caret-up-fill":
		return CaretUpFill(), nil
	case "cart":
		return Cart(), nil
	case "catalog":
		return Catalog(), nil
	case "catalog-option":
		return CatalogOption(), nil
	case "centos":
		return Centos(), nil
	case "certificate":
		return Certificate(), nil
	case "channel":
		return Channel(), nil
	case "chapter-add":
		return ChapterAdd(), nil
	case "chapter-next":
		return ChapterNext(), nil
	case "chapter-previous":
		return ChapterPrevious(), nil
	case "chat":
		return Chat(), nil
	case "chat-option":
		return ChatOption(), nil
	case "checkbox":
		return Checkbox(), nil
	case "checkbox-selected":
		return CheckboxSelected(), nil
	case "checkmark":
		return Checkmark(), nil
	case "chrome":
		return Chrome(), nil
	case "circle-alert":
		return CircleAlert(), nil
	case "circle-information":
		return CircleInformation(), nil
	case "circle-play":
		return CirclePlay(), nil
	case "circle-question":
		return CircleQuestion(), nil
	case "clear":
		return Clear(), nil
	case "clear-option":
		return ClearOption(), nil
	case "cli":
		return Cli(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clock":
		return Clock(), nil
	case "clone":
		return Clone(), nil
	case "close":
		return Close(), nil
	case "closed-caption":
		return ClosedCaption(), nil
	case "cloud":
		return Cloud(), nil
	case "cloud-computer":
		return CloudComputer(), nil
	case "cloud-download":
		return CloudDownload(), nil
	case "cloud-software":
		return CloudSoftware(), nil
	case "cloud-upload":
		return CloudUpload(), nil
	case "cloudlinux":
		return Cloudlinux(), nil
	case "cluster":
		return Cluster(), nil
	case "coat-check":
		return CoatCheck(), nil
	case "code":
		return Code(), nil
	case "code-sandbox":
		return CodeSandbox(), nil
	case "codepen":
		return Codepen(), nil
	case "coffee":
		return Coffee(), nil
	case "columns":
		return Columns(), nil
	case "command":
		return Command(), nil
	case "compare":
		return Compare(), nil
	case "compass":
		return Compass(), nil
	case "compliance":
		return Compliance(), nil
	case "configure":
		return Configure(), nil
	case "connect":
		return Connect(), nil
	case "connectivity":
		return Connectivity(), nil
	case "console":
		return Console(), nil
	case "contact":
		return Contact(), nil
	case "contact-info":
		return ContactInfo(), nil
	case "contract":
		return Contract(), nil
	case "copy":
		return Copy(), nil
	case "cpu":
		return Cpu(), nil
	case "creative-commons":
		return CreativeCommons(), nil
	case "credit-card":
		return CreditCard(), nil
	case "css3":
		return CssThree(), nil
	case "cube":
		return Cube(), nil
	case "cubes":
		return Cubes(), nil
	case "currency":
		return Currency(), nil
	case "cursor":
		return Cursor(), nil
	case "cut":
		return Cut(), nil
	case "cycle":
		return Cycle(), nil
	case "dashboard":
		return Dashboard(), nil
	case "database":
		return Database(), nil
	case "debian":
		return Debian(), nil
	case "deliver":
		return Deliver(), nil
	case "deploy":
		return Deploy(), nil
	case "descend":
		return Descend(), nil
	case "descending":
		return Descending(), nil
	case "desktop":
		return Desktop(), nil
	case "detach":
		return Detach(), nil
	case "device":
		return Device(), nil
	case "diamond":
		return Diamond(), nil
	case "directions":
		return Directions(), nil
	case "disabled-outline":
		return DisabledOutline(), nil
	case "disc":
		return Disc(), nil
	case "dislike":
		return Dislike(), nil
	case "docker":
		return Docker(), nil
	case "document":
		return Document(), nil
	case "document-cloud":
		return DocumentCloud(), nil
	case "document-config":
		return DocumentConfig(), nil
	case "document-csv":
		return DocumentCsv(), nil
	case "document-download":
		return DocumentDownload(), nil
	case "document-excel":
		return DocumentExcel(), nil
	case "document-image":
		return DocumentImage(), nil
	case "document-locked":
		return DocumentLocked(), nil
	case "document-missing":
		return DocumentMissing(), nil
	case "document-notes":
		return DocumentNotes(), nil
	case "document-outlook":
		return DocumentOutlook(), nil
	case "document-pdf":
		return DocumentPdf(), nil
	case "document-performance":
		return DocumentPerformance(), nil
	case "document-ppt":
		return DocumentPpt(), nil
	case "document-rtf":
		return DocumentRtf(), nil
	case "document-sound":
		return DocumentSound(), nil
	case "document-store":
		return DocumentStore(), nil
	case "document-test":
		return DocumentTest(), nil
	case "document-text":
		return DocumentText(), nil
	case "document-threat":
		return DocumentThreat(), nil
	case "document-time":
		return DocumentTime(), nil
	case "document-transfer":
		return DocumentTransfer(), nil
	case "document-txt":
		return DocumentTxt(), nil
	case "document-update":
		return DocumentUpdate(), nil
	case "document-upload":
		return DocumentUpload(), nil
	case "document-user":
		return DocumentUser(), nil
	case "document-verified":
		return DocumentVerified(), nil
	case "document-video":
		return DocumentVideo(), nil
	case "document-windows":
		return DocumentWindows(), nil
	case "document-word":
		return DocumentWord(), nil
	case "document-zip":
		return DocumentZip(), nil
	case "domain":
		return Domain(), nil
	case "dos":
		return Dos(), nil
	case "down":
		return Down(), nil
	case "download":
		return Download(), nil
	case "download-option":
		return DownloadOption(), nil
	case "drag":
		return Drag(), nil
	case "drawer":
		return Drawer(), nil
	case "dribbble":
		return Dribbble(), nil
	case "drive-cage":
		return DriveCage(), nil
	case "dropbox":
		return Dropbox(), nil
	case "duplicate":
		return Duplicate(), nil
	case "dxc":
		return Dxc(), nil
	case "edge":
		return Edge(), nil
	case "edit":
		return Edit(), nil
	case "eject":
		return Eject(), nil
	case "elevator":
		return Elevator(), nil
	case "emergency":
		return Emergency(), nil
	case "emoji":
		return Emoji(), nil
	case "empty-circle":
		return EmptyCircle(), nil
	case "erase":
		return Erase(), nil
	case "escalator":
		return Escalator(), nil
	case "expand":
		return Expand(), nil
	case "ezmeral":
		return Ezmeral(), nil
	case "facebook":
		return Facebook(), nil
	case "facebook-option":
		return FacebookOption(), nil
	case "fan":
		return Fan(), nil
	case "fan-option":
		return FanOption(), nil
	case "fast-forward":
		return FastForward(), nil
	case "favorite":
		return Favorite(), nil
	case "fedora":
		return Fedora(), nil
	case "figma":
		return Figma(), nil
	case "filter":
		return Filter(), nil
	case "finger-print":
		return FingerPrint(), nil
	case "fireball":
		return Fireball(), nil
	case "firefox":
		return Firefox(), nil
	case "firewall":
		return Firewall(), nil
	case "flag":
		return Flag(), nil
	case "flag-fill":
		return FlagFill(), nil
	case "flows":
		return Flows(), nil
	case "folder":
		return Folder(), nil
	case "folder-cycle":
		return FolderCycle(), nil
	case "folder-open":
		return FolderOpen(), nil
	case "form-add":
		return FormAdd(), nil
	case "form-attachment":
		return FormAttachment(), nil
	case "form-calendar":
		return FormCalendar(), nil
	case "form-checkmark":
		return FormCheckmark(), nil
	case "form-clock":
		return FormClock(), nil
	case "form-close":
		return FormClose(), nil
	case "form-cut":
		return FormCut(), nil
	case "form-down":
		return FormDown(), nil
	case "form-edit":
		return FormEdit(), nil
	case "form-filter":
		return FormFilter(), nil
	case "form-folder":
		return FormFolder(), nil
	case "form-location":
		return FormLocation(), nil
	case "form-lock":
		return FormLock(), nil
	case "form-next":
		return FormNext(), nil
	case "form-next-link":
		return FormNextLink(), nil
	case "form-previous":
		return FormPrevious(), nil
	case "form-previous-link":
		return FormPreviousLink(), nil
	case "form-refresh":
		return FormRefresh(), nil
	case "form-schedule":
		return FormSchedule(), nil
	case "form-search":
		return FormSearch(), nil
	case "form-subtract":
		return FormSubtract(), nil
	case "form-trash":
		return FormTrash(), nil
	case "form-up":
		return FormUp(), nil
	case "form-upload":
		return FormUpload(), nil
	case "form-view":
		return FormView(), nil
	case "form-view-hide":
		return FormViewHide(), nil
	case "forward-ten":
		return ForwardTen(), nil
	case "freebsd":
		return Freebsd(), nil
	case "gallery":
		return Gallery(), nil
	case "gamepad":
		return Gamepad(), nil
	case "gateway":
		return Gateway(), nil
	case "gatsbyjs":
		return Gatsbyjs(), nil
	case "gem":
		return Gem(), nil
	case "gift":
		return Gift(), nil
	case "github":
		return Github(), nil
	case "globe":
		return Globe(), nil
	case "golang":
		return Golang(), nil
	case "google":
		return Google(), nil
	case "google-play":
		return GooglePlay(), nil
	case "google-plus":
		return GooglePlus(), nil
	case "google-wallet":
		return GoogleWallet(), nil
	case "graph-ql":
		return GraphQl(), nil
	case "gremlin":
		return Gremlin(), nil
	case "grid":
		return Grid(), nil
	case "grommet":
		return Grommet(), nil
	case "group":
		return Group(), nil
	case "grow":
		return Grow(), nil
	case "hadoop":
		return Hadoop(), nil
	case "halt":
		return Halt(), nil
	case "help":
		return Help(), nil
	case "help-option":
		return HelpOption(), nil
	case "heroku":
		return Heroku(), nil
	case "hide":
		return Hide(), nil
	case "history":
		return History(), nil
	case "home":
		return Home(), nil
	case "home-option":
		return HomeOption(), nil
	case "home-rounded":
		return HomeRounded(), nil
	case "horton":
		return Horton(), nil
	case "host":
		return Host(), nil
	case "host-maintenance":
		return HostMaintenance(), nil
	case "hp":
		return Hp(), nil
	case "hpe":
		return Hpe(), nil
	case "hpe-labs":
		return HpeLabs(), nil
	case "hpi":
		return Hpi(), nil
	case "html5":
		return HtmlFive(), nil
	case "ice-cream":
		return IceCream(), nil
	case "image":
		return Image(), nil
	case "impact":
		return Impact(), nil
	case "in-progress":
		return InProgress(), nil
	case "inbox":
		return Inbox(), nil
	case "indicator":
		return Indicator(), nil
	case "info":
		return Info(), nil
	case "inherit":
		return Inherit(), nil
	case "insecure":
		return Insecure(), nil
	case "inspect":
		return Inspect(), nil
	case "instagram":
		return Instagram(), nil
	case "install":
		return Install(), nil
	case "install-option":
		return InstallOption(), nil
	case "integration":
		return Integration(), nil
	case "internet-explorer":
		return InternetExplorer(), nil
	case "italic":
		return Italic(), nil
	case "iteration":
		return Iteration(), nil
	case "java":
		return Java(), nil
	case "js":
		return Js(), nil
	case "key":
		return Key(), nil
	case "keyboard":
		return Keyboard(), nil
	case "language":
		return Language(), nil
	case "lastfm":
		return Lastfm(), nil
	case "launch":
		return Launch(), nil
	case "layer":
		return Layer(), nil
	case "license":
		return License(), nil
	case "like":
		return Like(), nil
	case "line-chart":
		return LineChart(), nil
	case "link":
		return Link(), nil
	case "link-bottom":
		return LinkBottom(), nil
	case "link-down":
		return LinkDown(), nil
	case "link-next":
		return LinkNext(), nil
	case "link-previous":
		return LinkPrevious(), nil
	case "link-top":
		return LinkTop(), nil
	case "link-up":
		return LinkUp(), nil
	case "linkedin":
		return Linkedin(), nil
	case "linkedin-option":
		return LinkedinOption(), nil
	case "list":
		return List(), nil
	case "local":
		return Local(), nil
	case "location":
		return Location(), nil
	case "location-pin":
		return LocationPin(), nil
	case "lock":
		return Lock(), nil
	case "login":
		return Login(), nil
	case "logout":
		return Logout(), nil
	case "lounge":
		return Lounge(), nil
	case "magic":
		return Magic(), nil
	case "mail":
		return Mail(), nil
	case "mail-option":
		return MailOption(), nil
	case "mandriva":
		return Mandriva(), nil
	case "manual":
		return Manual(), nil
	case "map":
		return Map(), nil
	case "map-location":
		return MapLocation(), nil
	case "mastercard":
		return Mastercard(), nil
	case "medium":
		return Medium(), nil
	case "memory":
		return Memory(), nil
	case "menu":
		return Menu(), nil
	case "microfocus":
		return Microfocus(), nil
	case "microphone":
		return Microphone(), nil
	case "money":
		return Money(), nil
	case "monitor":
		return Monitor(), nil
	case "monospace":
		return Monospace(), nil
	case "moon":
		return Moon(), nil
	case "more":
		return More(), nil
	case "more-vertical":
		return MoreVertical(), nil
	case "mouse":
		return Mouse(), nil
	case "multimedia":
		return Multimedia(), nil
	case "multiple":
		return Multiple(), nil
	case "music":
		return Music(), nil
	case "mysql":
		return Mysql(), nil
	case "navigate":
		return Navigate(), nil
	case "network":
		return Network(), nil
	case "network-drive":
		return NetworkDrive(), nil
	case "new":
		return New(), nil
	case "new-window":
		return NewWindow(), nil
	case "next":
		return Next(), nil
	case "node":
		return Node(), nil
	case "nodes":
		return Nodes(), nil
	case "norton":
		return Norton(), nil
	case "note":
		return Note(), nil
	case "notes":
		return Notes(), nil
	case "notification":
		return Notification(), nil
	case "npm":
		return Npm(), nil
	case "object-group":
		return ObjectGroup(), nil
	case "object-ungroup":
		return ObjectUngroup(), nil
	case "offline-storage":
		return OfflineStorage(), nil
	case "onedrive":
		return Onedrive(), nil
	case "opera":
		return Opera(), nil
	case "optimize":
		return Optimize(), nil
	case "oracle":
		return Oracle(), nil
	case "ordered-list":
		return OrderedList(), nil
	case "organization":
		return Organization(), nil
	case "overview":
		return Overview(), nil
	case "package":
		return Package(), nil
	case "paint":
		return Paint(), nil
	case "pan":
		return Pan(), nil
	case "pause":
		return Pause(), nil
	case "pause-fill":
		return PauseFill(), nil
	case "paypal":
		return Paypal(), nil
	case "performance":
		return Performance(), nil
	case "personal-computer":
		return PersonalComputer(), nil
	case "phone":
		return Phone(), nil
	case "phone-flip":
		return PhoneFlip(), nil
	case "phone-horizontal":
		return PhoneHorizontal(), nil
	case "phone-vertical":
		return PhoneVertical(), nil
	case "pie-chart":
		return PieChart(), nil
	case "pied-piper":
		return PiedPiper(), nil
	case "pin":
		return Pin(), nil
	case "pinterest":
		return Pinterest(), nil
	case "plan":
		return Plan(), nil
	case "play":
		return Play(), nil
	case "play-fill":
		return PlayFill(), nil
	case "plug":
		return Plug(), nil
	case "pocket":
		return Pocket(), nil
	case "power":
		return Power(), nil
	case "power-cycle":
		return PowerCycle(), nil
	case "power-force-shutdown":
		return PowerForceShutdown(), nil
	case "power-reset":
		return PowerReset(), nil
	case "power-shutdown":
		return PowerShutdown(), nil
	case "previous":
		return Previous(), nil
	case "print":
		return Print(), nil
	case "product-hunt":
		return ProductHunt(), nil
	case "projects":
		return Projects(), nil
	case "qr":
		return Qr(), nil
	case "radial":
		return Radial(), nil
	case "radial-selected":
		return RadialSelected(), nil
	case "raspberry":
		return Raspberry(), nil
	case "reactjs":
		return Reactjs(), nil
	case "reddit":
		return Reddit(), nil
	case "redhat":
		return Redhat(), nil
	case "redo":
		return Redo(), nil
	case "refresh":
		return Refresh(), nil
	case "resources":
		return Resources(), nil
	case "restaurant":
		return Restaurant(), nil
	case "restroom":
		return Restroom(), nil
	case "restroom-men":
		return RestroomMen(), nil
	case "restroom-women":
		return RestroomWomen(), nil
	case "resume":
		return Resume(), nil
	case "return":
		return Return(), nil
	case "revert":
		return Revert(), nil
	case "rewind":
		return Rewind(), nil
	case "risk":
		return Risk(), nil
	case "robot":
		return Robot(), nil
	case "rotate-left":
		return RotateLeft(), nil
	case "rotate-right":
		return RotateRight(), nil
	case "rss":
		return Rss(), nil
	case "run":
		return Run(), nil
	case "safari-option":
		return SafariOption(), nil
	case "sans":
		return Sans(), nil
	case "satellite":
		return Satellite(), nil
	case "save":
		return Save(), nil
	case "scan":
		return Scan(), nil
	case "schedule":
		return Schedule(), nil
	case "schedule-new":
		return ScheduleNew(), nil
	case "schedule-play":
		return SchedulePlay(), nil
	case "schedules":
		return Schedules(), nil
	case "sco":
		return Sco(), nil
	case "scorecard":
		return Scorecard(), nil
	case "script":
		return Script(), nil
	case "sd":
		return Sd(), nil
	case "search":
		return Search(), nil
	case "search-advanced":
		return SearchAdvanced(), nil
	case "secure":
		return Secure(), nil
	case "select":
		return Select(), nil
	case "selection":
		return Selection(), nil
	case "semantics":
		return Semantics(), nil
	case "send":
		return Send(), nil
	case "server":
		return Server(), nil
	case "server-cluster":
		return ServerCluster(), nil
	case "servers":
		return Servers(), nil
	case "service-play":
		return ServicePlay(), nil
	case "services":
		return Services(), nil
	case "settings-option":
		return SettingsOption(), nil
	case "share":
		return Share(), nil
	case "share-option":
		return ShareOption(), nil
	case "share-rounded":
		return ShareRounded(), nil
	case "shield":
		return Shield(), nil
	case "shield-security":
		return ShieldSecurity(), nil
	case "shift":
		return Shift(), nil
	case "shop":
		return Shop(), nil
	case "sidebar":
		return Sidebar(), nil
	case "sign":
		return Sign(), nil
	case "skype":
		return Skype(), nil
	case "slack":
		return Slack(), nil
	case "snapchat":
		return Snapchat(), nil
	case "solaris":
		return Solaris(), nil
	case "sort":
		return Sort(), nil
	case "soundcloud":
		return Soundcloud(), nil
	case "spa":
		return Spa(), nil
	case "spectrum":
		return Spectrum(), nil
	case "split":
		return Split(), nil
	case "splits":
		return Splits(), nil
	case "spotify":
		return Spotify(), nil
	case "square":
		return Square(), nil
	case "stack-overflow":
		return StackOverflow(), nil
	case "stakeholder":
		return Stakeholder(), nil
	case "star":
		return Star(), nil
	case "star-half":
		return StarHalf(), nil
	case "status-critical":
		return StatusCritical(), nil
	case "status-critical-small":
		return StatusCriticalSmall(), nil
	case "status-disabled":
		return StatusDisabled(), nil
	case "status-disabled-small":
		return StatusDisabledSmall(), nil
	case "status-good":
		return StatusGood(), nil
	case "status-good-small":
		return StatusGoodSmall(), nil
	case "status-info":
		return StatusInfo(), nil
	case "status-info-small":
		return StatusInfoSmall(), nil
	case "status-placeholder":
		return StatusPlaceholder(), nil
	case "status-placeholder-small":
		return StatusPlaceholderSmall(), nil
	case "status-unknown":
		return StatusUnknown(), nil
	case "status-unknown-small":
		return StatusUnknownSmall(), nil
	case "status-warning":
		return StatusWarning(), nil
	case "status-warning-small":
		return StatusWarningSmall(), nil
	case "steps":
		return Steps(), nil
	case "steps-option":
		return StepsOption(), nil
	case "stop":
		return Stop(), nil
	case "stop-fill":
		return StopFill(), nil
	case "storage":
		return Storage(), nil
	case "street-view":
		return StreetView(), nil
	case "strike-through":
		return StrikeThrough(), nil
	case "stripe":
		return Stripe(), nil
	case "subscript":
		return Subscript(), nil
	case "subtract":
		return Subtract(), nil
	case "subtract-circle":
		return SubtractCircle(), nil
	case "sun":
		return Sun(), nil
	case "superscript":
		return Superscript(), nil
	case "support":
		return Support(), nil
	case "suse":
		return Suse(), nil
	case "swift":
		return Swift(), nil
	case "swim":
		return Swim(), nil
	case "switch":
		return Switch(), nil
	case "sync":
		return Sync(), nil
	case "system":
		return System(), nil
	case "table":
		return Table(), nil
	case "table-add":
		return TableAdd(), nil
	case "tag":
		return Tag(), nil
	case "tape":
		return Tape(), nil
	case "tape-option":
		return TapeOption(), nil
	case "target":
		return Target(), nil
	case "task":
		return Task(), nil
	case "tasks":
		return Tasks(), nil
	case "technology":
		return Technology(), nil
	case "template":
		return Template(), nil
	case "terminal":
		return Terminal(), nil
	case "test":
		return Test(), nil
	case "test-desktop":
		return TestDesktop(), nil
	case "text-align-center":
		return TextAlignCenter(), nil
	case "text-align-full":
		return TextAlignFull(), nil
	case "text-align-left":
		return TextAlignLeft(), nil
	case "text-align-right":
		return TextAlignRight(), nil
	case "text-wrap":
		return TextWrap(), nil
	case "threats":
		return Threats(), nil
	case "three-d":
		return ThreeD(), nil
	case "three-d-effects":
		return ThreeDEffects(), nil
	case "ticket":
		return Ticket(), nil
	case "tictok":
		return Tictok(), nil
	case "time":
		return Time(), nil
	case "tip":
		return Tip(), nil
	case "toast":
		return Toast(), nil
	case "tools":
		return Tools(), nil
	case "tooltip":
		return Tooltip(), nil
	case "top-corner":
		return TopCorner(), nil
	case "train":
		return Train(), nil
	case "transaction":
		return Transaction(), nil
	case "trash":
		return Trash(), nil
	case "tree":
		return Tree(), nil
	case "tree-option":
		return TreeOption(), nil
	case "trigger":
		return Trigger(), nil
	case "trophy":
		return Trophy(), nil
	case "troubleshoot":
		return Troubleshoot(), nil
	case "tty":
		return Tty(), nil
	case "tumblr":
		return Tumblr(), nil
	case "turbolinux":
		return Turbolinux(), nil
	case "twitter":
		return Twitter(), nil
	case "ubuntu":
		return Ubuntu(), nil
	case "underline":
		return Underline(), nil
	case "undo":
		return Undo(), nil
	case "unlink":
		return Unlink(), nil
	case "unlock":
		return Unlock(), nil
	case "unordered-list":
		return UnorderedList(), nil
	case "unsorted":
		return Unsorted(), nil
	case "up":
		return Up(), nil
	case "update":
		return Update(), nil
	case "upgrade":
		return Upgrade(), nil
	case "upload":
		return Upload(), nil
	case "upload-option":
		return UploadOption(), nil
	case "usb-key":
		return UsbKey(), nil
	case "user":
		return User(), nil
	case "user-add":
		return UserAdd(), nil
	case "user-admin":
		return UserAdmin(), nil
	case "user-expert":
		return UserExpert(), nil
	case "user-female":
		return UserFemale(), nil
	case "user-manager":
		return UserManager(), nil
	case "user-new":
		return UserNew(), nil
	case "user-police":
		return UserPolice(), nil
	case "user-settings":
		return UserSettings(), nil
	case "user-worker":
		return UserWorker(), nil
	case "validate":
		return Validate(), nil
	case "vend":
		return Vend(), nil
	case "video":
		return Video(), nil
	case "view":
		return View(), nil
	case "vimeo":
		return Vimeo(), nil
	case "virtual-machine":
		return VirtualMachine(), nil
	case "virtual-storage":
		return VirtualStorage(), nil
	case "visa":
		return Visa(), nil
	case "vm-maintenance":
		return VmMaintenance(), nil
	case "vmware":
		return Vmware(), nil
	case "volume":
		return Volume(), nil
	case "volume-control":
		return VolumeControl(), nil
	case "volume-low":
		return VolumeLow(), nil
	case "volume-mute":
		return VolumeMute(), nil
	case "vulnerability":
		return Vulnerability(), nil
	case "waypoint":
		return Waypoint(), nil
	case "webcam":
		return Webcam(), nil
	case "wheelchair":
		return Wheelchair(), nil
	case "wheelchair-active":
		return WheelchairActive(), nil
	case "wifi":
		return Wifi(), nil
	case "wifi-low":
		return WifiLow(), nil
	case "wifi-medium":
		return WifiMedium(), nil
	case "wifi-none":
		return WifiNone(), nil
	case "windows":
		return Windows(), nil
	case "windows-legacy":
		return WindowsLegacy(), nil
	case "wordpress":
		return Wordpress(), nil
	case "workshop":
		return Workshop(), nil
	case "yoga":
		return Yoga(), nil
	case "youtube":
		return Youtube(), nil
	case "zoom":
		return Zoom(), nil
	case "zoom-in":
		return ZoomIn(), nil
	case "zoom-out":
		return ZoomOut(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in grommet_icons icon set", name)
	}
}
