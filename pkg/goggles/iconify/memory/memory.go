package memory

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	accountInnerSVG                          = `<path fill="currentColor" d="M9 3h4v1h1v1h1v4h-1v1h-1v1H9v-1H8V9H7V5h1V4h1V3m1 5v1h2V8h1V6h-1V5h-2v1H9v2h1m-3 4h8v1h2v1h1v1h1v4H3v-4h1v-1h1v-1h2v-1m-1 4H5v1h12v-1h-1v-1h-2v-1H8v1H6v1Z"/>`
	accountBoxInnerSVG                       = `<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m0 14h1v-1h2v-1h8v1h2v1h1V5h-1V4H5v1H4v11m12 2v-1h-2v-1H8v1H6v1h10M9 5h4v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V7h1V6h1V5m3 3V7h-2v1H9v2h1v1h2v-1h1V8h-1Z"/>`
	alertInnerSVG                            = `<path fill="currentColor" d="M20 20H2v-1H1v-4h1v-2h1v-2h1V9h1V7h1V5h1V3h1V2h6v1h1v2h1v2h1v2h1v2h1v2h1v2h1v4h-1v1M9 6H8v2H7v2H6v2H5v2H4v2H3v2h16v-2h-1v-2h-1v-2h-1v-2h-1V8h-1V6h-1V4H9v2m1 1h2v6h-2V7m0 7h2v2h-2v-2Z"/>`
	alertBoxInnerSVG                         = `<path fill="currentColor" d="M12 12h-2V6h2m0 10h-2v-2h2m6 6H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`
	alertBoxFillInnerSVG                     = `<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-6-8V6h-2v6Zm0 4v-2h-2v2Z"/>`
	alertCircleInnerSVG                      = `<path fill="currentColor" d="M12 12h-2V6h2Zm0 4h-2v-2h2Zm3 5H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-1-2v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1Z"/>`
	alertCircleFillInnerSVG                  = `<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-3-9V6h-2v6Zm0 4v-2h-2v2Z"/>`
	alertRhombusInnerSVG                     = `<path fill="currentColor" d="M12 12h-2V6h2Zm0 4h-2v-2h2Zm0 5h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm0-3v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-2v1H9v1H8v1H7v1H6v1H5v1H4v2h1v1h1v1h1v1h1v1h1v1h1v1Z"/>`
	alertRhombusFillInnerSVG                 = `<path fill="currentColor" d="M12 21h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm0-9V6h-2v6Zm0 4v-2h-2v2Z"/>`
	alphaAInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 3v1h1v9h-2v-4h-2v4H8V7h1V6h4m-1 2h-2v2h2V8Z"/>`
	alphaAFillInnerSVG                       = `<path fill="currentColor" d="M10 8h2v2h-2V8m5-7v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 5H9v1H8v9h2v-4h2v4h2V7h-1V6Z"/>`
	alphaBInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v3h-1v2h1v3h-1v1H8V6m2 2v2h2V8h-2m2 4h-2v2h2v-2Z"/>`
	alphaBFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h5v-1h1v-3h-1v-2h1V7h-1V6H8m2 2h2v2h-2V8m2 4v2h-2v-2h2Z"/>`
	alphaCInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v2h-2V8h-2v6h2v-1h2v2h-1v1H9v-1H8V7h1V6Z"/>`
	alphaCFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1v-2h-2v1h-2V8h2v1h2V7h-1V6H9Z"/>`
	alphaDInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v1h1v6h-1v1h-1v1H8V6m2 2v6h2v-1h1V9h-1V8h-2Z"/>`
	alphaDFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h5v-1h1v-1h1V8h-1V7h-1V6H8m2 2h2v1h1v4h-1v1h-2V8Z"/>`
	alphaEInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-4v2h4v2h-4v2h4v2H8V6Z"/>`
	alphaEFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h6v-2h-4v-2h4v-2h-4V8h4V6H8Z"/>`
	alphaFInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-4v2h3v2h-3v4H8V6Z"/>`
	alphaFFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h3v-2h-3V8h4V6H8Z"/>`
	alphaGInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h5v2h-4v6h2v-2h-1v-2h3v5h-1v1H9v-1H8V7h1V6Z"/>`
	alphaGFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1v-5h-3v2h1v2h-2V8h4V6H9Z"/>`
	alphaHInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v4h2V6h2v10h-2v-4h-2v4H8V6Z"/>`
	alphaHFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h2v4h2V6h-2v4h-2V6H8Z"/>`
	alphaIInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 3v2h-1v6h1v2H9v-2h1V8H9V6h4Z"/>`
	alphaIFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 5H9v2h1v6H9v2h4v-2h-1V8h1V6Z"/>`
	alphaJInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-2 3h2v9h-1v1H9v-1H8v-2h2v1h2V6Z"/>`
	alphaJFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-3 5v8h-2v-1H8v2h1v1h4v-1h1V6h-2Z"/>`
	alphaKInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v2h1V7h1V6h2v2h-1v1h-1v1h-1v1h1v1h1v1h1v3h-2v-2h-1v-1h-1v3H8V6Z"/>`
	alphaKFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-3h1v1h1v2h2v-3h-1v-1h-1v-1h-1v-1h1V9h1V8h1V6h-2v1h-1v1h-1V6H8Z"/>`
	alphaLInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v8h4v2H8V6Z"/>`
	alphaLFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h6v-2h-4V6H8Z"/>`
	alphaMInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h8v1h1v9h-2V8h-2v7h-2V8H8v8H6V7h1V6Z"/>`
	alphaMFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v1H6v9h2V8h2v7h2V8h2v8h2V7h-1V6H7Z"/>`
	alphaNInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v2h1v2h1V6h2v10h-2v-2h-1v-2h-1v4H8V6Z"/>`
	alphaNFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h1v2h1v2h2V6h-2v4h-1V8h-1V6H8Z"/>`
	alphaOInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v8h-1v1H9v-1H8V7h1V6m1 2v6h2V8h-2Z"/>`
	alphaOFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1V7h-1V6H9m1 2h2v6h-2V8Z"/>`
	alphaPInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v4h-1v1h-3v4H8V6m2 2v2h2V8h-2Z"/>`
	alphaPFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h3v-1h1V7h-1V6H8m2 2h2v2h-2V8Z"/>`
	alphaQInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v8h1v1h-1v1h-1v-1h-1v-1h-1v1H9v-1H8V7h1V6m1 2v6h1v-1h1V8h-2Z"/>`
	alphaQFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h2v-1h1v1h1v1h1v-1h1v-1h-1V7h-1V6H9m1 2h2v5h-1v1h-1V8Z"/>`
	alphaRInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v4h-1v2h1v3h-2v-2h-1v-2h-1v4H8V6m2 2v2h2V8h-2Z"/>`
	alphaRFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h1v2h1v2h2v-3h-1v-2h1V7h-1V6H8m2 2h2v2h-2V8Z"/>`
	alphaSInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V7h1V6Z"/>`
	alphaSFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v4h1v1h3v2H8v2h5v-1h1v-4h-1v-1h-3V8h4V6H9Z"/>`
	alphaTInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-2v8h-2V8H8V6Z"/>`
	alphaTFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v2h2v8h2V8h2V6H8Z"/>`
	alphaUInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v8h2V6h2v9h-1v1H9v-1H8V6Z"/>`
	alphaUFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v9h1v1h4v-1h1V6h-2v8h-2V6H8Z"/>`
	alphaVInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h2v3h1v2h2V9h1V6h2v4h-1v2h-1v2h-1v2h-2v-2H9v-2H8v-2H7V6Z"/>`
	alphaVFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v4h1v2h1v2h1v2h2v-2h1v-2h1v-2h1V6h-2v3h-1v2h-2V9H9V6H7Z"/>`
	alphaWInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M6 6h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1H8v-1H7v-2H6V6Z"/>`
	alphaWFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M6 6v7h1v2h1v1h2v-1h2v1h2v-1h1v-2h1V6h-2v6h-1v1h-1V8h-2v5H9v-1H8V6H6Z"/>`
	alphaXInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 13v-2h-1v-1h-2v1H9v2H7v-3h1v-1h1v-2H8V9H7V6h2v2h1v1h2V8h1V6h2v3h-1v1h-1v2h1v1h1v3h-2Z"/>`
	alphaXFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 15h2v-3h-1v-1h-1v-2h1V9h1V6h-2v2h-1v1h-2V8H9V6H7v3h1v1h1v2H8v1H7v3h2v-2h1v-1h2v1h1v2Z"/>`
	alphaYInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h2v2h1v2h2V8h1V6h2v3h-1v2h-1v2h-1v3h-2v-3H9v-2H8V9H7V6Z"/>`
	alphaYFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v3h1v2h1v2h1v3h2v-3h1v-2h1V9h1V6h-2v2h-1v2h-2V8H9V6H7Z"/>`
	alphaZInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v4h-1v1h-1v1h-1v1h-1v1h4v2H8v-4h1v-1h1v-1h1V9h1V8H8V6Z"/>`
	alphaZFillInnerSVG                       = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v2h4v1h-1v1h-1v1H9v1H8v4h6v-2h-4v-1h1v-1h1v-1h1v-1h1V6H8Z"/>`
	applicationInnerSVG                      = `<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h16v1h1v16h-1ZM18 6V4H4v2Zm0 12V8H4v10Z"/>`
	applicationCodeInnerSVG                  = `<path fill="currentColor" d="M11 16H9v-1H8v-4h1v-1h2v2h-1v2h1m4 2h-2v-2h1v-2h-1v-2h2v1h1v4h-1m4 5H3v-1H2V3h1V2h16v1h1v16h-1M18 6V4H4v2m14 12V8H4v10Z"/>`
	archiveInnerSVG                          = `<path fill="currentColor" d="M2 2h18v6h-1v12H3V8H2V2m15 16V8H5v10h12M4 4v2h14V4H4m3 6h8v2H7v-2Z"/>`
	arrowInnerSVG                            = `<path fill="currentColor" d="M20 2h-3v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H3v1H2v1H1v1h1v1h1v1h1v1h1v-1h1v-1h1v-3h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h2"/>`
	arrowBottomLeftInnerSVG                  = `<path fill="currentColor" d="M14 17H5V8h2v5h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1v1h1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1h5v2Z"/>`
	arrowBottomLeftCircleInnerSVG            = `<path fill="currentColor" d="M1 15V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1m4 1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1m9-1H7V8h2v3h1v-1h1V9h1V8h1V7h1v1h1v1h-1v1h-1v1h-1v1h-1v1h3v2Z"/>`
	arrowBottomRightInnerSVG                 = `<path fill="currentColor" d="M17 8v9H8v-2h5v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7h1V6h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1V8h2Z"/>`
	arrowBottomRightCircleInnerSVG           = `<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1m1-4h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1m-1-9v7H8v-2h3v-1h-1v-1H9v-1H8V9H7V8h1V7h1v1h1v1h1v1h1v1h1V8h2Z"/>`
	arrowDownInnerSVG                        = `<path fill="currentColor" d="M12 17h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V4h2v9h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`
	arrowDownBoldInnerSVG                    = `<path fill="currentColor" d="M20 10v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h5V2h8v8h5m-4 2h-3V4H9v8H6v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1Z"/>`
	arrowDownCircleInnerSVG                  = `<path fill="currentColor" d="M16 10v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V6h2v6h1v-1h1v-1h2m5-3v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1Z"/>`
	arrowDownLeftInnerSVG                    = `<path fill="currentColor" d="M3 12v2h1v1h1v1h1v1h1v1h2v-2H8v-1H7v-1h4v-1h2v-1h1v-2h1V3h-2v6h-1v2h-2v1H7v-1h1v-1h1V8H7v1H6v1H5v1H4v1"/>`
	arrowDownRightInnerSVG                   = `<path fill="currentColor" d="M19 12v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h-4v-1H9v-1H8v-2H7V3h2v6h1v2h2v1h3v-1h-1v-1h-1V8h2v1h1v1h1v1h1v1"/>`
	arrowLeftInnerSVG                        = `<path fill="currentColor" d="M5 12v-2h1V9h1V8h1V7h1V6h2v2h-1v1H9v1h9v2H9v1h1v1h1v2H9v-1H8v-1H7v-1H6v-1"/>`
	arrowLeftBoldInnerSVG                    = `<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v5h8v8h-8v5m-2-4v-3h8V9h-8V6H9v1H8v1H7v1H6v1H5v2h1v1h1v1h1v1h1v1h1Z"/>`
	arrowLeftCircleInnerSVG                  = `<path fill="currentColor" d="M12 16h-2v-1H9v-1H8v-1H7v-1H6v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h6v2h-6v1h1v1h1v2m3 5H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1m1-4h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1Z"/>`
	arrowLeftDownInnerSVG                    = `<path fill="currentColor" d="M10 19H8v-1H7v-1H6v-1H5v-1H4v-2h2v1h1v1h1v-4h1V9h1V8h2V7h7v2h-6v1h-2v2h-1v3h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`
	arrowLeftRightInnerSVG                   = `<path fill="currentColor" d="M19 10v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1H3v-2h1V9h1V8h1V7h1V6h2v2H8v1H7v1h8V9h-1V8h-1V6h2v1h1v1h1v1h1v1"/>`
	arrowLeftUpInnerSVG                      = `<path fill="currentColor" d="M10 3H8v1H7v1H6v1H5v1H4v2h2V8h1V7h1v4h1v2h1v1h2v1h7v-2h-6v-1h-2v-2h-1V7h1v1h1v1h2V7h-1V6h-1V5h-1V4h-1"/>`
	arrowRightInnerSVG                       = `<path fill="currentColor" d="M17 10v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H4v-2h9V9h-1V8h-1V6h2v1h1v1h1v1h1v1"/>`
	arrowRightBoldInnerSVG                   = `<path fill="currentColor" d="M10 2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-5H2V7h8V2m2 4v3H4v4h8v3h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6h-1Z"/>`
	arrowRightCircleInnerSVG                 = `<path fill="currentColor" d="M7 1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1M6 5H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1m4 1h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H6v-2h6V9h-1V8h-1V6Z"/>`
	arrowRightDownInnerSVG                   = `<path fill="currentColor" d="M12 19h2v-1h1v-1h1v-1h1v-1h1v-2h-2v1h-1v1h-1v-4h-1V9h-1V8h-2V7H3v2h6v1h2v2h1v3h-1v-1h-1v-1H8v2h1v1h1v1h1v1h1"/>`
	arrowRightUpInnerSVG                     = `<path fill="currentColor" d="M12 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v4h-1v2h-1v1h-2v1H3v-2h6v-1h2v-2h1V7h-1v1h-1v1H8V7h1V6h1V5h1V4h1"/>`
	arrowTopLeftInnerSVG                     = `<path fill="currentColor" d="M5 14V5h9v2H9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7v5H5Z"/>`
	arrowTopLeftCircleInnerSVG               = `<path fill="currentColor" d="M7 1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1M6 5H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1m1 9V7h7v2h-3v1h1v1h1v1h1v1h1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v-1H9v3H7Z"/>`
	arrowTopRightInnerSVG                    = `<path fill="currentColor" d="M8 5h9v9h-2V9h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v-1H6v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7H8V5Z"/>`
	arrowTopRightCircleInnerSVG              = `<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1M8 7h7v7h-2v-3h-1v1h-1v1h-1v1H9v1H8v-1H7v-1h1v-1h1v-1h1v-1h1V9H8V7Z"/>`
	arrowUpInnerSVG                          = `<path fill="currentColor" d="M10 5h2v1h1v1h1v1h1v1h1v2h-2v-1h-1V9h-1v9h-2V9H9v1H8v1H6V9h1V8h1V7h1V6h1"/>`
	arrowUpBoldInnerSVG                      = `<path fill="currentColor" d="M2 12v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-5v8H7v-8H2m4-2h3v8h4v-8h3V9h-1V8h-1V7h-1V6h-1V5h-2v1H9v1H8v1H7v1H6v1Z"/>`
	arrowUpCircleInnerSVG                    = `<path fill="currentColor" d="M6 12v-2h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v6h-2v-6H9v1H8v1H6m-5 3V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1m4 1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1Z"/>`
	arrowUpDownInnerSVG                      = `<path fill="currentColor" d="M10 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v8h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V7H9v1H8v1H6V7h1V6h1V5h1V4h1"/>`
	arrowUpLeftInnerSVG                      = `<path fill="currentColor" d="M3 10V8h1V7h1V6h1V5h1V4h2v2H8v1H7v1h4v1h2v1h1v2h1v7h-2v-6h-1v-2h-2v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1"/>`
	arrowUpRightInnerSVG                     = `<path fill="currentColor" d="M19 10V8h-1V7h-1V6h-1V5h-1V4h-2v2h1v1h1v1h-4v1H9v1H8v2H7v7h2v-6h1v-2h2v-1h3v1h-1v1h-1v2h2v-1h1v-1h1v-1h1v-1"/>`
	axeInnerSVG                              = `<path fill="currentColor" d="M11 3h2v1h1v1h1v1h2v1h2v1h1v1h1v2h-1v2h-1v1h-1v1h-2v1h-1v-1h-1v-1h-1v-2h-1v-1h-1v-1h-1V9H9V8H8V6h1V5h1V4h1m-1 6v1h1v2h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v-1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1"/>`
	bagPersonalInnerSVG                      = `<path fill="currentColor" d="M17 15H9v2H7v-2H5v4h12v-4m0-6h-1V8h-1V7H7v1H6v1H5v4h12V9m-4 2H9v-1h1V9h2v1h1v1M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3V8m6-5v2h4V3H9Z"/>`
	bagPersonalFillInnerSVG                  = `<path fill="currentColor" d="M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3V8m6-5v2h4V3H9m8 11H5v1h2v2h1v-2h9v-1m-5-3h1V9h-1V8h-2v1H9v2h1v1h2v-1Z"/>`
	batteryFiftyInnerSVG                     = `<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7Z"/>`
	batteryOneHundredInnerSVG                = `<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7m-6 1h2v6h-2V8m3 0h2v6h-2V8Z"/>`
	batterySeventyFiveInnerSVG               = `<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7m-6 1h2v6h-2V8Z"/>`
	batteryTwentyFiveInnerSVG                = `<path fill="currentColor" d="M7 8v6H5V8h2m11-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7Z"/>`
	batteryZeroInnerSVG                      = `<path fill="currentColor" d="M3 5h15v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5m1 2v8h13V7H4Z"/>`
	battleAxeInnerSVG                        = `<path fill="currentColor" d="M15 1h-4v1h-1v1H9v1H8v4h4v1h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1H2v1H1v1h1v1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v4h4v-1h1v-1h1v-1h1V7h-4V5h-2"/>`
	bookInnerSVG                             = `<path fill="currentColor" d="M3 2h1V1h14v1h1v18h-1v1H4v-1H3V2m8 7h-1V8H9v1H8v1H7V3H5v16h12V3h-5v7h-1V9Z"/>`
	bookmarkInnerSVG                         = `<path fill="currentColor" d="M5 2h12v1h1v17h-2v-1h-2v-1h-2v-1h-2v1H8v1H6v1H4V3h1V2m1 2v13h1v-1h2v-1h1v-1h2v1h1v1h2v1h1V4H6Z"/>`
	borderBottomInnerSVG                     = `<path fill="currentColor" d="M20 12h-2v-2h2v2M4 12H2v-2h2v2m8-8h-2V2h2v2m4 0h-2V2h2v2m4 4h-2V6h2v2m0-4h-2V2h2v2m0 12h-2v-2h2v2M4 16H2v-2h2v2M8 4H6V2h2v2M4 4H2V2h2v2m0 4H2V6h2v2m16 12H2v-2h18v2Z"/>`
	borderBottomLeftInnerSVG                 = `<path fill="currentColor" d="M10 4V2h2v2h-2m8 8v-2h2v2h-2m0 4v-2h2v2h-2M6 4V2h2v2H6m12 4V6h2v2h-2m0-4V2h2v2h-2m-4 0V2h2v2h-2M2 20V2h2v16h16v2H2Z"/>`
	borderBottomLeftRightInnerSVG            = `<path fill="currentColor" d="M10 4V2h2v2h-2M6 4V2h2v2H6m8 0V2h2v2h-2m4-2h2v18H2V2h2v16h14V2Z"/>`
	borderBottomRightInnerSVG                = `<path fill="currentColor" d="M4 12H2v-2h2v2m8-8h-2V2h2v2m4 0h-2V2h2v2M4 16H2v-2h2v2M8 4H6V2h2v2M4 4H2V2h2v2m0 4H2V6h2v2m16 12H2v-2h16V2h2v18Z"/>`
	borderInsideInnerSVG                     = `<path fill="currentColor" d="M2 10h8V2h2v8h8v2h-8v8h-2v-8H2v-2m4 8h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 2h2v2H2V2m0 4h2v2H2V6m4-4h2v2H6V2m8 0h2v2h-2V2m4 0h2v2h-2V2m0 4h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2Z"/>`
	borderLeftInnerSVG                       = `<path fill="currentColor" d="M10 20v-2h2v2h-2m0-16V2h2v2h-2m8 8v-2h2v2h-2m0 4v-2h2v2h-2m-4 4v-2h2v2h-2m4 0v-2h2v2h-2M6 20v-2h2v2H6M6 4V2h2v2H6m12 4V6h2v2h-2m0-4V2h2v2h-2m-4 0V2h2v2h-2M2 20V2h2v18H2Z"/>`
	borderLeftRightInnerSVG                  = `<path fill="currentColor" d="M12 2v2h-2V2h2m0 16v2h-2v-2h2M8 2v2H6V2h2m8 0v2h-2V2h2m0 16v2h-2v-2h2m-8 0v2H6v-2h2M4 2v18H2V2h2m14 0h2v18h-2V2Z"/>`
	borderNoneInnerSVG                       = `<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2m-8-8h2v2h-2V2m0 16h2v2h-2v-2m-4 0h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 2h2v2H2V2m0 4h2v2H2V6m4-4h2v2H6V2m8 0h2v2h-2V2m4 0h2v2h-2V2m0 4h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2Z"/>`
	borderOutsideInnerSVG                    = `<path fill="currentColor" d="M2 2h18v18H2V2m2 2v14h14V4H4m6 2h2v2h-2V6m0 4h2v2h-2v-2m-4 0h2v2H6v-2m8 0h2v2h-2v-2m-4 4h2v2h-2v-2Z"/>`
	borderRightInnerSVG                      = `<path fill="currentColor" d="M12 2v2h-2V2h2m0 16v2h-2v-2h2m-8-8v2H2v-2h2m0-4v2H2V6h2m4-4v2H6V2h2M4 2v2H2V2h2m12 0v2h-2V2h2m0 16v2h-2v-2h2M4 14v2H2v-2h2m0 4v2H2v-2h2m4 0v2H6v-2h2M20 2v18h-2V2h2Z"/>`
	borderTopInnerSVG                        = `<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2m-8 8h2v2h-2v-2m-4 0h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 6h2v2H2V6m16 0h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2M2 2h18v2H2V2Z"/>`
	borderTopBottomInnerSVG                  = `<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2M2 14h2v2H2v-2m0-8h2v2H2V6m16 0h2v2h-2V6m0 8h2v2h-2v-2M2 18h18v2H2v-2M2 4V2h18v2H2Z"/>`
	borderTopLeftInnerSVG                    = `<path fill="currentColor" d="M18 10h2v2h-2v-2m-8 8h2v2h-2v-2m-4 0h2v2H6v-2M18 6h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2M2 2h18v2H4v16H2V2Z"/>`
	borderTopLeftBottomInnerSVG              = `<path fill="currentColor" d="M18 10h2v2h-2v-2m0-4h2v2h-2V6m0 8h2v2h-2v-2m2 4v2H2V2h18v2H4v14h16Z"/>`
	borderTopLeftRightInnerSVG               = `<path fill="currentColor" d="M12 18v2h-2v-2h2m4 0v2h-2v-2h2m-8 0v2H6v-2h2m-4 2H2V2h18v18h-2V4H4v16Z"/>`
	borderTopRightInnerSVG                   = `<path fill="currentColor" d="M12 18v2h-2v-2h2m-8-8v2H2v-2h2m0-4v2H2V6h2m12 12v2h-2v-2h2M4 14v2H2v-2h2m0 4v2H2v-2h2m4 0v2H6v-2h2M20 2v18h-2V4H2V2h18Z"/>`
	borderTopRightBottomInnerSVG             = `<path fill="currentColor" d="M4 12H2v-2h2v2m0 4H2v-2h2v2m0-8H2V6h2v2M2 4V2h18v18H2v-2h16V4H2Z"/>`
	bowInnerSVG                              = `<path fill="currentColor" d="M1 3h10v1h2v1h2v1h1v1h1v2h1v2h1v10h-2v-2h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6V9H5V8H4V7H3V5H1m15 13h1v-6h-1v-2h-1V8h-1V7h-2V6h-2V5H4v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1"/>`
	bowArrowInnerSVG                         = `<path fill="currentColor" d="M1 3h10v1h2v1h3V4h1V2h3v3h-2v1h-1v3h1v2h1v10h-2v-2h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v1H8v1H7v3H6v1H5v1H4v-1H3v-1H2v-1H1v-1h1v-1h1v-1h3v-1h1v-1h1v-2H7v-1H6V9H5V8H4V7H3V5H1V3m15 15h1v-6h-1v-2h-1V9h-1v1h-1v1h-1v1h-1v1h1v1h1v1h1v1h1v1h1v1M12 7V6h-2V5H4v1h1v1h1v1h1v1h1v1h1v1h1v-1h1V9h1V8h1V7h-1Z"/>`
	boxInnerSVG                              = `<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m13 3V4H5v1H4v12h1v1h12v-1h1V5h-1Z"/>`
	boxLightDashedDownLeftInnerSVG           = `<path fill="currentColor" d="M22 12h-2v-2h2v2m-4 0h-3v-2h3v2m-5 0h-3V9h2v1h1v2M12 0v2h-2V0h2m0 4v3h-2V4h2Z"/>`
	boxLightDashedDownRightInnerSVG          = `<path fill="currentColor" d="M12 0v2h-2V0h2m0 4v3h-2V4h2m0 5v3H9v-2h1V9h2M.002 10H2v2H.002v-2M4 10h3v2H4v-2Z"/>`
	boxLightDashedUpLeftInnerSVG             = `<path fill="currentColor" d="M10 22v-2h2v2h-2m0-4v-3h2v3h-2m0-5v-3h3v2h-1v1h-2m12-1h-2v-2h2v2m-4 0h-3v-2h3v2Z"/>`
	boxLightDashedUpRightInnerSVG            = `<path fill="currentColor" d="M0 10h2v2H0v-2m4 0h3v2H4v-2m5 0h3v3h-2v-1H9v-2m1 12v-2h2v2h-2m0-4v-3h2v3h-2Z"/>`
	boxLightDoubleHorizontalInnerSVG         = `<path fill="currentColor" d="M0 12h22v2H0v-2m0-4h22v2H0V8Z"/>`
	boxLightDoubleRoundDownLeftInnerSVG      = `<path fill="currentColor" d="M10 0v3h1v2h1v2h1v1h1v1h1v1h2v1h2v1h3v2h-4v-1h-2v-1h-2v-1h-1v-1h-1V9h-1V8h-1V6H9V4H8V0h2m12 10h-2V9h-3V8h-1V7h-1V6h-1V5h-1V2h-1V0h2v1h1v3h1v1h1v1h1v1h3v1h1v2Z"/>`
	boxLightDoubleRoundDownRightInnerSVG     = `<path fill="currentColor" d="M0 12h3v-1h2v-1h2V9h1V8h1V7h1V5h1V3h1V0h2v4h-1v2h-1v2h-1v1h-1v1H9v1H8v1H6v1H4v1H0v-2M10 0v2H9v3H8v1H7v1H6v1H5v1H2v1H0V8h1V7h3V6h1V5h1V4h1V1h1V0h2Z"/>`
	boxLightDoubleRoundUpLeftInnerSVG        = `<path fill="currentColor" d="M22 10h-3v1h-2v1h-2v1h-1v1h-1v1h-1v2h-1v2h-1v3H8v-4h1v-2h1v-2h1v-1h1v-1h1v-1h1v-1h2V9h2V8h4v2M12 22v-2h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h2v2h-1v1h-3v1h-1v1h-1v1h-1v3h-1v1h-2Z"/>`
	boxLightDoubleRoundUpRightInnerSVG       = `<path fill="currentColor" d="M12 22v-3h-1v-2h-1v-2H9v-1H8v-1H7v-1H5v-1H3v-1H0V8h4v1h2v1h2v1h1v1h1v1h1v1h1v2h1v2h1v4h-2M0 12h2v1h3v1h1v1h1v1h1v1h1v3h1v2H8v-1H7v-3H6v-1H5v-1H4v-1H1v-1H0v-2Z"/>`
	boxLightDoubleVerticalInnerSVG           = `<path fill="currentColor" d="M12 22V0h2v22h-2m-4 0V0h2v22H8Z"/>`
	boxLightDownLeftInnerSVG                 = `<path fill="currentColor" d="M22 12H10V0h2v10h10v2Z"/>`
	boxLightDownLeftCircleInnerSVG           = `<path fill="currentColor" d="M13 15H9v-1H8v-1H7V9h1V8h1V7h1V0h2v7h1v1h1v1h1v1h7v2h-7v1h-1v1h-1v1m-4-3h1v1h2v-1h1v-2h-1V9h-2v1H9v2Z"/>`
	boxLightDownRightInnerSVG                = `<path fill="currentColor" d="M12 0v12H0v-2h10V0h2Z"/>`
	boxLightDownRightCircleInnerSVG          = `<path fill="currentColor" d="M15 9v4h-1v1h-1v1H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7h1V0h2v7h1v1h1v1h1m-3 4v-1h1v-2h-1V9h-2v1H9v2h1v1h2Z"/>`
	boxLightHorizontalInnerSVG               = `<path fill="currentColor" d="M0 10h22v2H0v-2Z"/>`
	boxLightRoundDownLeftInnerSVG            = `<path fill="currentColor" d="M22 12h-4v-1h-3v-1h-1V9h-1V8h-1V7h-1V4h-1V0h2v3h1v3h1v1h1v1h1v1h3v1h3v2Z"/>`
	boxLightRoundDownRightInnerSVG           = `<path fill="currentColor" d="M12 0v4h-1v3h-1v1H9v1H8v1H7v1H4v1H0v-2h3V9h3V8h1V7h1V6h1V3h1V0h2Z"/>`
	boxLightRoundUpLeftInnerSVG              = `<path fill="currentColor" d="M10 22v-4h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h4v2h-3v1h-3v1h-1v1h-1v1h-1v3h-1v3h-2Z"/>`
	boxLightRoundUpRightInnerSVG             = `<path fill="currentColor" d="M0 10h4v1h3v1h1v1h1v1h1v1h1v3h1v4h-2v-3H9v-3H8v-1H7v-1H6v-1H3v-1H0v-2Z"/>`
	boxLightUpLeftInnerSVG                   = `<path fill="currentColor" d="M10 22V10h12v2H12v10h-2Z"/>`
	boxLightUpLeftCircleInnerSVG             = `<path fill="currentColor" d="M7 13V9h1V8h1V7h4v1h1v1h1v1h7v2h-7v1h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1H7m3-4v1H9v2h1v1h2v-1h1v-2h-1V9h-2Z"/>`
	boxLightUpRightInnerSVG                  = `<path fill="currentColor" d="M0 10h12v12h-2V12H0v-2Z"/>`
	boxLightUpRightCircleInnerSVG            = `<path fill="currentColor" d="M9 7h4v1h1v1h1v4h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7m4 3h-1V9h-2v1H9v2h1v1h2v-1h1v-2Z"/>`
	boxLightVerticalInnerSVG                 = `<path fill="currentColor" d="M12 0v22h-2V0h2Z"/>`
	boxOuterLightAllInnerSVG                 = `<path fill="currentColor" d="M0 0h22v22H0V0m2 2v18h18V2H2Z"/>`
	boxOuterLightDashedAllInnerSVG           = `<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v2h2v2H0v-4h2M6 0h4v2H6V0m6 0h4v2h-4V0m6 0h4v4h-2V2h-2V0m0 22v-2h2v-2h2v4h-4m-2 0h-4v-2h4v2m-6 0H6v-2h4v2M20 6h2v4h-2V6m0 6h2v4h-2v-4Z"/>`
	boxOuterLightDashedDownInnerSVG          = `<path fill="currentColor" d="M1 20h4v2H1v-2m6 0h3v2H7v-2m5 0h4v2h-4v-2m6 0h3v2h-3v-2Z"/>`
	boxOuterLightDashedDownLeftInnerSVG      = `<path fill="currentColor" d="M0 18h2v2h2v2H0v-4m6 2h4v2H6v-2m6 0h4v2h-4v-2m6 0h3v2h-3v-2M0 16v-4h2v4H0m0-6V7h2v3H0m0-5V1h2v4H0Z"/>`
	boxOuterLightDashedDownLeftRightInnerSVG = `<path fill="currentColor" d="M0 18h2v2h2v2H0v-4m6 2h4v2H6v-2m6 0h4v2h-4v-2m6 0h2v-2h2v4h-4v-2M0 16v-4h2v4H0m0-6V7h2v3H0m0-5V1h2v4H0m22-4v4h-2V1h2m0 6v3h-2V7h2m0 5v4h-2v-4h2Z"/>`
	boxOuterLightDashedDownRightInnerSVG     = `<path fill="currentColor" d="M18 22v-2h2v-2h2v4h-4m2-6v-4h2v4h-2m0-6V6h2v4h-2m0-6V1h2v3h-2m-4 18h-4v-2h4v2m-6 0H7v-2h3v2m-5 0H1v-2h4v2Z"/>`
	boxOuterLightDashedLeftInnerSVG          = `<path fill="currentColor" d="M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0m0-6V1h2v3H0Z"/>`
	boxOuterLightDashedLeftRightInnerSVG     = `<path fill="currentColor" d="M20 21v-4h2v4h-2m0-6v-3h2v3h-2m0-5V6h2v4h-2m0-6V1h2v3h-2M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0m0-6V1h2v3H0Z"/>`
	boxOuterLightDashedRightInnerSVG         = `<path fill="currentColor" d="M20 21v-4h2v4h-2m0-6v-3h2v3h-2m0-5V6h2v4h-2m0-6V1h2v3h-2Z"/>`
	boxOuterLightDashedUpInnerSVG            = `<path fill="currentColor" d="M1 0h4v2H1V0m6 0h3v2H7V0m5 0h4v2h-4V0m6 0h3v2h-3V0Z"/>`
	boxOuterLightDashedUpDownInnerSVG        = `<path fill="currentColor" d="M21 2h-4V0h4v2m-6 0h-3V0h3v2m-5 0H6V0h4v2M4 2H1V0h3v2m17 20h-4v-2h4v2m-6 0h-3v-2h3v2m-5 0H6v-2h4v2m-6 0H1v-2h3v2Z"/>`
	boxOuterLightDashedUpDownLeftInnerSVG    = `<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v2h2v2H0v-4h2M6 0h4v2H6V0m6 0h3v2h-3V0m5 0h4v2h-4V0m4 22h-4v-2h4v2m-6 0h-3v-2h3v2m-5 0H6v-2h4v2Z"/>`
	boxOuterLightDashedUpDownRightInnerSVG   = `<path fill="currentColor" d="M18 22v-2h2v-2h2v4h-4m2-6v-4h2v4h-2m0-6V6h2v4h-2m0-6V2h-2V0h4v4h-2m-4 18h-4v-2h4v2m-6 0H7v-2h3v2m-5 0H1v-2h4v2M1 0h4v2H1V0m6 0h3v2H7V0m5 0h4v2h-4V0Z"/>`
	boxOuterLightDashedUpLeftInnerSVG        = `<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v3H0v-3h2M6 0h4v2H6V0m6 0h3v2h-3V0m5 0h4v2h-4V0Z"/>`
	boxOuterLightDashedUpLeftRightInnerSVG   = `<path fill="currentColor" d="M22 4h-2V2h-2V0h4v4m-6-2h-4V0h4v2m-6 0H6V0h4v2M4 2H2v2H0V0h4v2m18 4v4h-2V6h2m0 6v3h-2v-3h2m0 5v4h-2v-4h2M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0Z"/>`
	boxOuterLightDashedUpRightInnerSVG       = `<path fill="currentColor" d="M22 4h-2V2h-2V0h4v4m-6-2h-4V0h4v2m-6 0H6V0h4v2M4 2H1V0h3v2m18 4v4h-2V6h2m0 6v3h-2v-3h2m0 5v4h-2v-4h2Z"/>`
	boxOuterLightDownInnerSVG                = `<path fill="currentColor" d="M0 20h22v2H0v-2Z"/>`
	boxOuterLightDownLeftInnerSVG            = `<path fill="currentColor" d="M0 0h2v20h20v2H0V0Z"/>`
	boxOuterLightDownLeftRightInnerSVG       = `<path fill="currentColor" d="M22 22H0V0h2v20h18V0h2v22Z"/>`
	boxOuterLightDownRightInnerSVG           = `<path fill="currentColor" d="M0 22v-2h20V0h2v22H0Z"/>`
	boxOuterLightLeftInnerSVG                = `<path fill="currentColor" d="M0 22V0h2v22H0Z"/>`
	boxOuterLightLeftRightInnerSVG           = `<path fill="currentColor" d="M0 22V0h2v22H0m20 0V0h2v22h-2Z"/>`
	boxOuterLightRightInnerSVG               = `<path fill="currentColor" d="M20 22V0h2v22h-2Z"/>`
	boxOuterLightRoundDownLeftInnerSVG       = `<path fill="currentColor" d="M0 0h2v5h1v4h1v2h1v2h1v1h1v1h1v1h1v1h2v1h2v1h4v1h5v2h-6v-1h-4v-1h-2v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-2H2v-2H1V6H0V0Z"/>`
	boxOuterLightRoundDownRightInnerSVG      = `<path fill="currentColor" d="M0 22v-2h5v-1h4v-1h2v-1h2v-1h1v-1h1v-1h1v-1h1v-2h1V9h1V5h1V0h2v6h-1v4h-1v2h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-2v1h-2v1H6v1H0Z"/>`
	boxOuterLightRoundUpLeftInnerSVG         = `<path fill="currentColor" d="M22 0v2h-5v1h-4v1h-2v1H9v1H8v1H7v1H6v1H5v2H4v2H3v4H2v5H.004v-6H1v-4h1v-2h1V8h1V7h1V6h1V5h1V4h1V3h2V2h2V1h4V0h6Z"/>`
	boxOuterLightRoundUpRightInnerSVG        = `<path fill="currentColor" d="M22 22h-2v-5h-1v-4h-1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-2V4H9V3H5V2H.004V0H6v1h4v1h2v1h2v1h1v1h1v1h1v1h1v1h1v2h1v2h1v4h1v6Z"/>`
	boxOuterLightUpInnerSVG                  = `<path fill="currentColor" d="M0 0h22v2H0V0Z"/>`
	boxOuterLightUpDownInnerSVG              = `<path fill="currentColor" d="M0 0h22v2H0V0m0 20h22v2H0v-2Z"/>`
	boxOuterLightUpDownLeftInnerSVG          = `<path fill="currentColor" d="M0 22V0h22v2H2v18h20v2H0Z"/>`
	boxOuterLightUpDownRightInnerSVG         = `<path fill="currentColor" d="M22 0v22H0v-2h20V2H0V0h22Z"/>`
	boxOuterLightUpLeftInnerSVG              = `<path fill="currentColor" d="M22 0v2H2v20H0V0h22Z"/>`
	boxOuterLightUpLeftRightInnerSVG         = `<path fill="currentColor" d="M0 0h22v22h-2V2H2v20H0V0Z"/>`
	boxOuterLightUpRightInnerSVG             = `<path fill="currentColor" d="M22 22h-2V2H.002V0H22v22Z"/>`
	briefcaseInnerSVG                        = `<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1V6m7 0h4V4H9v2m10 2H3v10h16V8Z"/>`
	bugInnerSVG                              = `<path fill="currentColor" d="M3 7h3V6h1V5h1V4H7V3H6V2h1V1h1v1h1v1h1v1h2V3h1V2h1V1h1v1h1v1h-1v1h-1v1h1v1h1v1h3v2h-2v2h2v2h-2v2h2v2h-3v1h-1v1h-1v1H8v-1H7v-1H6v-1H3v-2h2v-2H3v-2h2V9H3V7m10 11v-1h1v-1h1V8h-1V7h-1V6H9v1H8v1H7v8h1v1h1v1h4m-4-5h4v2H9v-2m0-4h4v2H9V9Z"/>`
	bugFillInnerSVG                          = `<path fill="currentColor" d="M3 7h3V6h1V5h1V4H7V3H6V2h1V1h1v1h1v1h1v1h2V3h1V2h1V1h1v1h1v1h-1v1h-1v1h1v1h1v1h3v2h-2v2h2v2h-2v2h2v2h-3v1h-1v1h-1v1H8v-1H7v-1H6v-1H3v-2h2v-2H3v-2h2V9H3V7m6 6v2h4v-2H9m0-4v2h4V9H9Z"/>`
	calculatorInnerSVG                       = `<path fill="currentColor" d="M18 21H4v-1H3V3h1V2h14v1h1v17h-1ZM16 7V4H6v3Zm-8 4V9H6v2Zm4 0V9h-2v2Zm4 0V9h-2v2Zm-8 4v-2H6v2Zm4 0v-2h-2v2Zm4 0v-2h-2v2Zm-8 4v-2H6v2Zm4 0v-2h-2v2Zm4 0v-2h-2v2Z"/>`
	calendarInnerSVG                         = `<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1v1M4 4v2h14V4H4m0 4v10h14V8H4m8 4h4v4h-4v-4Z"/>`
	calendarMonthInnerSVG                    = `<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1v1M4 4v2h14V4H4m0 4v10h14V8H4m10 6h2v2h-2v-2m-4 0h2v2h-2v-2m-4 0h2v2H6v-2m0-4h2v2H6v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`
	cancelInnerSVG                           = `<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-2h1V8h-1V6h-1m-3 10v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7H5V6H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h-1v-1h-1Z"/>`
	cardInnerSVG                             = `<path fill="currentColor" d="M2 3h18v1h1v14h-1v1H2v-1H1V4h1V3m1 2v12h16V5H3Z"/>`
	cardTextInnerSVG                         = `<path fill="currentColor" d="M17 8v2H5V8h12M5 12h10v2H5v-2M2 3h18v1h1v14h-1v1H2v-1H1V4h1V3m1 2v12h16V5H3Z"/>`
	cartInnerSVG                             = `<path fill="currentColor" d="M19 14v2H6v-1H5v-4H4V8H3V3H1V1h4v3h16v4h-1v3h-1v1H7v2h12M5 7h1v3h12V7h1V6H5v1m2 10h2v1h1v2H9v1H7v-1H6v-2h1v-1m8 0h2v1h1v2h-1v1h-2v-1h-1v-2h1v-1Z"/>`
	chartBarInnerSVG                         = `<path fill="currentColor" d="M2 2h2v16h16v2H2V2m4 14V8h4v8H6m5 0V4h4v12h-4m5 0v-5h4v5h-4Z"/>`
	chatInnerSVG                             = `<path fill="currentColor" d="M7 4h9v1h2v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-2v1H8v1H4v1H1v-2h2v-1h1v-2H3v-1H2V8h1V7h1V6h1V5h2V4m10 4V7h-2V6H8v1H6v1H5v1H4v4h1v1h1v1h2v1h7v-1h2v-1h1v-1h1V9h-1V8h-1Z"/>`
	checkInnerSVG                            = `<path fill="currentColor" d="M4 11h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-2Z"/>`
	checkboxBlankInnerSVG                    = `<path fill="currentColor" d="M3 4h1V3h14v1h1v14h-1v1H4v-1H3V4m2 13h12V5H5v12Z"/>`
	checkboxCrossInnerSVG                    = `<path fill="currentColor" d="M13 12h1v1h1v2h-2v-1h-1v-1h-2v1H9v1H7v-2h1v-1h1v-2H8V9H7V7h2v1h1v1h2V8h1V7h2v2h-1v1h-1v2m5 7H4v-1H3V4h1V3h14v1h1v14h-1v1M5 5v12h12V5H5Z"/>`
	checkboxMarkedInnerSVG                   = `<path fill="currentColor" d="M3 4h1V3h14v1h-1v1H5v12h12v-6h1v-1h1v8h-1v1H4v-1H3V4m3 5h2v1h1v1h1v1h2v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6V9Z"/>`
	chevronDownInnerSVG                      = `<path fill="currentColor" d="M16 10h1V9h1V7h-2v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8V9H7V8H6V7H4v2h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1"/>`
	chevronDownCircleInnerSVG                = `<path fill="currentColor" d="M16 9v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6V9h2v1h1v1h1v1h2v-1h1v-1h1V9h2m-1-8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`
	chevronLeftInnerSVG                      = `<path fill="currentColor" d="M12 16v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1v-1h-1v-2h1V9h1V8h1V7h1V6h1V4h-2v1h-1v1h-1v1h-1v1H9v1H8v1H7v2h1v1h1v1h1v1h1v1"/>`
	chevronLeftCircleInnerSVG                = `<path fill="currentColor" d="M13 16h-2v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h-1v2h1v1h1v1h1v2m2-15v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`
	chevronRightInnerSVG                     = `<path fill="currentColor" d="M10 6V5H9V4H7v2h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1H9v1H8v1H7v2h2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6"/>`
	chevronRightCircleInnerSVG               = `<path fill="currentColor" d="M9 16h2v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H9v2h1v1h1v1h1v2h-1v1h-1v1H9v2m6-15v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`
	chevronUpInnerSVG                        = `<path fill="currentColor" d="M6 12H5v1H4v2h2v-1h1v-1h1v-1h1v-1h1v-1h2v1h1v1h1v1h1v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-1V7h-2v1H9v1H8v1H7v1H6"/>`
	chevronUpCircleInnerSVG                  = `<path fill="currentColor" d="M6 13v-2h1v-1h1V9h1V8h1V7h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6m9-12v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`
	circleInnerSVG                           = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`
	clipboardInnerSVG                        = `<path fill="currentColor" d="M2 5h1V4h4V2h2V1h4v1h2v2h4v1h1v15h-1v1H3v-1H2V5m8-2v2h2V3h-2m8 3h-2v2H6V6H4v13h14V6Z"/>`
	clockInnerSVG                            = `<path fill="currentColor" d="M10 5h2v6h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-1V5m5-4v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`
	coffeeInnerSVG                           = `<path fill="currentColor" d="M1 20v-2h16v2H1M2 3h17v1h1v6h-1v1h-3v3h-1v1h-1v1H4v-1H3v-1H2V3m14 2v4h2V5h-2M4 5v8h1v1h8v-1h1V5H4Z"/>`
	commentInnerSVG                          = `<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1V2m1 2v12h5v3h1v-1h1v-1h1v-1h8V4H3Z"/>`
	commentTextInnerSVG                      = `<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1V2m1 2v12h5v3h1v-1h1v-1h1v-1h8V4H3m2 3h12v2H5V7m0 4h10v2H5v-2Z"/>`
	compassInnerSVG                          = `<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-4 6V8h2V7h2V6h2v2h-1v2h-1v2h-1v1h-1v1h-2v1H8v1H6v-2h1v-2h1v-2h1V9h1m2 1h-2v2h2v-2Z"/>`
	compassEastArrowInnerSVG                 = `<path fill="currentColor" d="M6 10v2H5v1H4v1H3v1H2v1H0v-2h1v-1h1v-1h1v-2H2V9H1V8H0V6h2v1h1v1h1v1h1v1m5-4h6v2h-4v2h4v2h-4v2h4v2h-6"/>`
	compassNorthArrowInnerSVG                = `<path fill="currentColor" d="M8 4h2v1.5h1V8h1V4h2v10h-2v-2h-1v-2h-1v4H8m2 3v-1h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6v-2h1v-1h1v-1h1v-1"/>`
	compassNorthEastInnerSVG                 = `<path fill="currentColor" d="M4 6h2v1.5h1V10h1V6h2v10H8v-2H7v-2H6v4H4m8-10h6v2h-4v2h4v2h-4v2h4v2h-6"/>`
	compassNorthWestInnerSVG                 = `<path fill="currentColor" d="M2 6h2v1.5h1V10h1V6h2v10H6v-2H5v-2H4v4H2m8-10h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`
	compassSouthArrowInnerSVG                = `<path fill="currentColor" d="M10 6h2V5h1V4h1V3h1V2h1V0h-2v1h-1v1h-1v1h-2V2H9V1H8V0H6v2h1v1h1v1h1v1h1M9 8h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V9h1"/>`
	compassSouthEastInnerSVG                 = `<path fill="currentColor" d="M5 6h5v2H6v2h3v1h1v4H9v1H4v-2h4v-2H5v-1H4V7h1m7-1h6v2h-4v2h4v2h-4v2h4v2h-6"/>`
	compassSouthWestInnerSVG                 = `<path fill="currentColor" d="M3 6h5v2H4v2h3v1h1v4H7v1H2v-2h4v-2H3v-1H2V7h1m7-1h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`
	compassWestArrowInnerSVG                 = `<path fill="currentColor" d="M4 6h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1H8v1H6v-1H5v-2H4m12-3v2h1v1h1v1h1v1h1v1h2v-2h-1v-1h-1v-1h-1v-2h1V9h1V8h1V6h-2v1h-1v1h-1v1h-1v1"/>`
	creditCardInnerSVG                       = `<path fill="currentColor" d="M2 4h18v1h1v12h-1v1H2v-1H1V5h1V4m1 2v2h16V6H3m0 10h16v-5H3v5Z"/>`
	crownInnerSVG                            = `<path fill="currentColor" d="M2 17h18v2H2v-2M4 6v1h1v1h1V7h1V6h1V5h1V4h1V3h2v1h1v1h1v1h1v1h1v1h1V7h1V6h1V5h1v11H2V5h1v1h1m3 8h11v-4h-3V9h-1V8h-1V7h-1V6h-2v1H9v1H8v1H7v1H4v4h3Z"/>`
	cubeUnfoldedInnerSVG                     = `<path fill="currentColor" d="M12 3v5h10v7h-5v5h-7v-5H0V8h5V3h7m-2 2H7v3h3V5m-3 5v3h3v-3H7m-2 0H2v3h3v-3m12 0v3h3v-3h-3m-2 5h-3v3h3v-3m-3-5v3h3v-3h-3Z"/>`
	databaseInnerSVG                         = `<path fill="currentColor" d="M7 2h8v1h2v1h1v1h1v12h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3V5h1V4h1V3h2V2m1 14v-1H6v-1H5v2h1v1h2v1h6v-1h2v-1h1v-2h-1v1h-2v1H8m0-5v-1H6V9H5v3h2v1h2v1h4v-1h2v-1h2V9h-1v1h-2v1H8m1-3v1h4V8h2V7h2V6h-1V5h-2V4H8v1H6v1H5v1h2v1h2Z"/>`
	deviceInnerSVG                           = `<path fill="currentColor" d="M2 1h18v1h.94v18H20v1H2v-1h-.94V2H2V1m1 2v16h16V3H3m1 1h14v8H4V4m1 10h3v3H5v-3m7 1h2v2h-2v-2m3-1h2v2h-2v-2Z"/>`
	diamondInnerSVG                          = `<path fill="currentColor" d="M6 2h10v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2V6h1V5h1V4h1V3h1V2m9 3V4h-1v2h1v1h2V6h-1V5h-1m-3 1V4h-2v2H9v1h4V6h-1M8 6V4H7v1H6v1H5v1h2V6h1m-4 5h1v1h1v1h1v1h1v-2H7V9H4v2m6 1v4h2v-4h1V9H9v3h1m4 0v2h1v-1h1v-1h1v-1h1V9h-3v3h-1Z"/>`
	divisionInnerSVG                         = `<path fill="currentColor" d="M12 8h-2V7H9V5h1V4h2v1h1v2h-1Zm5 4H5v-2h12Zm-5 6h-2v-1H9v-2h1v-1h2v1h1v2h-1Z"/>`
	doorInnerSVG                             = `<path fill="currentColor" d="M12 10h2v2h-2v-2m4-8v1h1v15h2v2H3v-2h2V3h1V2h10m-1 2H7v14h8V4Z"/>`
	doorBoxInnerSVG                          = `<path fill="currentColor" d="M13 14h-2v-2h2Zm3 4h1v-1h1V5h-1V4H5v1H4v12h1v1h1V6h10Zm2 2H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-4-2V8H8v10Z"/>`
	doorOpenInnerSVG                         = `<path fill="currentColor" d="M10 10v2H9v-2h1M6 2h10v1h1v15h2v2H3v-2h2V3h1V2m1 2v14h4V4H7m6 0v1h1V4h-1m1 1v1h1V5h-1m0 1h-1v1h1V6m0 1v1h1V7h-1m0 1h-1v1h1V8m0 1v1h1V9h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1Z"/>`
	downloadInnerSVG                         = `<path fill="currentColor" d="M18 17v2H4v-2h14M14 2v6h4v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V8h4V2h6m-2 2h-2v6H9v1h1v1h2v-1h1v-1h-1V4Z"/>`
	emailInnerSVG                            = `<path fill="currentColor" d="M1 5h1V4h18v1h1v13h-1v1H2v-1H1V5m2 12h16V9h-1v1h-2v1h-2v1h-2v1h-2v-1H8v-1H6v-1H4V9H3v8M19 6H3v1h2v1h2v1h2v1h4V9h2V8h2V7h2V6Z"/>`
	fileInnerSVG                             = `<path fill="currentColor" d="M13 1v1h1v1h1v1h1v1h1v1h1v1h1v13h-1v1H4v-1H3V2h1V1h9m0 3h-1v4h4V7h-1V6h-1V5h-1V4M5 3v16h12v-9h-6V9h-1V3H5Z"/>`
	fireInnerSVG                             = `<path fill="currentColor" d="M14 20H7v-1H6v-1H5v-1H4v-5h1v-2h1V9h1V8h1v1h1v2h1V9h1V5h-1V4H9V3H8V2h3v1h2v1h1v1h1v1h1v1h1v2h1v7h-1v2h-1v1h-2m-2-1v-1h2v-1h1v-2h1v-4h-1V8h-1V7h-1v4h-1v2h-1v1h-1v1H9v-1H8v-3H7v1H6v4h1v1h1v1Z"/>`
	flaskInnerSVG                            = `<path fill="currentColor" d="M11 12h2v2h-2v-2m3-11v1h1v3h-1v2h1v2h1v2h1v2h1v1h1v2h1v4h-1v1H3v-1H2v-4h1v-2h1v-1h1v-2h1V9h1V7h1V5H7V2h1V1h6m-2 2h-2v5H9v2H8v2H7v1H6v2H5v2h1v-1h1v-1h1v-1h1v1h1v1h1v1h1v1h2v-1h1v-2h1v-1h-1v-2h-1v-2h-1V8h-1V3Z"/>`
	flaskEmptyInnerSVG                       = `<path fill="currentColor" d="M8 1h6v1h1v3h-1v2h1v2h1v2h1v2h1v1h1v2h1v4h-1v1H3v-1H2v-4h1v-2h1v-1h1v-2h1V9h1V7h1V5H7V2h1V1m2 2v5H9v2H8v2H7v1H6v2H5v2H4v2h14v-2h-1v-2h-1v-1h-1v-2h-1v-2h-1V8h-1V3h-2Z"/>`
	flaskRoundBottomInnerSVG                 = `<path fill="currentColor" d="M11 11h2v2h-2v-2m2-10v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1V1h4m-1 4h-2v4H9v1H8v1H7v1H6v2h1v-1h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h-1v-1h-1v-1h-1V9h-1V5Z"/>`
	flaskRoundBottomEmptyInnerSVG            = `<path fill="currentColor" d="M9 1h4v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1V1m1 4v4H9v1H8v1H7v1H6v4h1v1h1v1h1v1h4v-1h1v-1h1v-1h1v-4h-1v-1h-1v-1h-1V9h-1V5h-2Z"/>`
	floppyDiskInnerSVG                       = `<path fill="currentColor" d="M2 3h1V2h13v1h1v1h1v1h1v1h1v13h-1v1H3v-1H2V3m16 4h-1V6h-1V5h-1v4H6V4H4v14h2v-5h10v5h2V7m-7-3v3h2V4h-2m3 14v-3H8v3h6Z"/>`
	folderInnerSVG                           = `<path fill="currentColor" d="M2 3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4h1V3m1 4v10h16V7H3Z"/>`
	folderOpenInnerSVG                       = `<path fill="currentColor" d="M1 4h1V3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4m2 5h16V7H9V6H8V5H3v4m0 8h16v-6H3v6Z"/>`
	gamepadCenterInnerSVG                    = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 7H9v1H8v4h1v1h4v-1h1V9h-1V8Z"/>`
	gamepadDownInnerSVG                      = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 14H9v4h4v-4Z"/>`
	gamepadDownLeftInnerSVG                  = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9m2 6v4h4v-4H9Z"/>`
	gamepadDownRightInnerSVG                 = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9M9 15v4h4v-4H9Z"/>`
	gamepadEmptyInnerSVG                     = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 2H9v6H3v4h6v6h4v-6h6V9h-6V3Z"/>`
	gamepadLeftInnerSVG                      = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9Z"/>`
	gamepadRightInnerSVG                     = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9Z"/>`
	gamepadUpInnerSVG                        = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 2H9v4h4V3Z"/>`
	gamepadUpLeftInnerSVG                    = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9m2-6v4h4V3H9Z"/>`
	gamepadUpRightInnerSVG                   = `<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9M9 3v4h4V3H9Z"/>`
	heartInnerSVG                            = `<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2H1V5h1V4h1V3h1V2h5v1h1v1h2V3h1V2h5v1h1v1h1v1h1v5h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1m-7-9v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V6h-1V5h-1V4h-3v1h-1v1h-1v1h-2V6H9V5H8V4H5v1H4v1H3v3h1v2h1Z"/>`
	imageInnerSVG                            = `<path fill="currentColor" d="M1 4h1V3h18v1h1v14h-1v1H2v-1H1V4m2 10h1v-1h1v-1h1v-1h1v-1h1V9h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2V5H3v9m11 3v-1h-1v-1h-1v-1h-1v-1h-1v-1H8v1H7v1H6v1H5v1H4v1h10m-1-9h1V7h2v1h1v2h-1v1h-2v-1h-1V8Z"/>`
	labelInnerSVG                            = `<path fill="currentColor" d="M2 4h13v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-1H1V5h1V4m14 9h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H3v10h11v-1h1v-1h1v-1Z"/>`
	labelVariantInnerSVG                     = `<path fill="currentColor" d="M15 4v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-2H6V9H5V8H4V7H3V6H2V4h13m-1 12v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H6v1h1v1h1v1h1v1h1v2H9v1H8v1H7v1H6v1h8Z"/>`
	ledInnerSVG                              = `<path fill="currentColor" d="M8 21v-6H4v-2h2V6h1V4h1V3h1V2h4v1h1v1h1v2h1v7h2v2h-4v6h-2v-6h-2v6H8m4-16V4h-2v1H9v2H8v6h1v-1h4v1h1V7h-1V5h-1Z"/>`
	lightbulbInnerSVG                        = `<path fill="currentColor" d="M8 19h6v2H8v-2m0-1v-4H7v-1H6v-1H5v-1H4V5h1V4h1V3h1V2h1V1h6v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v4H8m5-6h1v-1h1v-1h1V6h-1V5h-1V4h-1V3H9v1H8v1H7v1H6v4h1v1h1v1h1v1h1v3h2v-3h1v-1Z"/>`
	lockInnerSVG                             = `<path fill="currentColor" d="M10 12h2v1h1v2h-1v1h-2v-1H9v-2h1v-1M8 2h6v1h1v1h1v4h1v1h1v10h-1v1H5v-1H4V9h1V8h1V4h1V3h1V2m1 2v1H8v3h6V5h-1V4H9m7 6H6v8h10v-8Z"/>`
	lockOpenInnerSVG                         = `<path fill="currentColor" d="M10 13h2v1h1v2h-1v1h-2v-1H9v-2h1v-1m4-11v1h1v1h1v5h1v1h1v10h-1v1H5v-1H4V10h1V9h9V5h-1V4H9v1H8v2H6V4h1V3h1V2h6m2 9H6v8h10v-8Z"/>`
	loginInnerSVG                            = `<path fill="currentColor" d="M5 1h12v1h1v18h-1v1H5v-1H4v-6h2v5h10V3H6v5H4V2h1V1m3 5h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1H8v-2h1v-1h1v-1H2v-2h8V9H9V8H8V6Z"/>`
	logoutInnerSVG                           = `<path fill="currentColor" d="M17 1v1h1v4h-1V5h-1V3H6v16h10v-2h1v-1h1v4h-1v1H5v-1H4V2h1V1h12m-4 5h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H8v-2h7V9h-1V8h-1V6Z"/>`
	mapInnerSVG                              = `<path fill="currentColor" d="M2 4h2V3h2V2h4v1h3v1h2V3h2V2h3v16h-2v1h-2v1h-4v-1H9v-1H7v1H5v1H2V4m2 2v11h2v-1h1V5H5v1H4m8-1H9v11h1v1h3V6h-1V5m4 1h-1v11h2v-1h1V5h-2v1Z"/>`
	menuBottomLeftInnerSVG                   = `<path fill="currentColor" d="M7 4h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2H7V4m2 4v5h5v-1h-1v-1h-1v-1h-1V9h-1V8H9Z"/>`
	menuBottomRightInnerSVG                  = `<path fill="currentColor" d="M4 15v-2h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v11H4m4-2h5V8h-1v1h-1v1h-1v1H9v1H8v1Z"/>`
	menuDownInnerSVG                         = `<path fill="currentColor" d="M4 8h14v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4V8m4 2v1h1v1h1v1h2v-1h1v-1h1v-1H8Z"/>`
	menuDownFillInnerSVG                     = `<path fill="currentColor" d="M17 9V8H5v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1V9"/>`
	menuLeftInnerSVG                         = `<path fill="currentColor" d="M15 4v14h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h1V5h1V4h2m-2 4h-1v1h-1v1h-1v2h1v1h1v1h1V8Z"/>`
	menuLeftFillInnerSVG                     = `<path fill="currentColor" d="M13 17h1V5h-1v1h-1v1h-1v1h-1v1H9v1H8v2h1v1h1v1h1v1h1v1h1"/>`
	menuLeftRightInnerSVG                    = `<path fill="currentColor" d="M12 4h2v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2V4m-2 0v14H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4h2m4 4v6h1v-1h1v-1h1v-2h-1V9h-1V8h-1M8 8H7v1H6v1H5v2h1v1h1v1h1V8Z"/>`
	menuRightInnerSVG                        = `<path fill="currentColor" d="M7 18V4h2v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7m2-4h1v-1h1v-1h1v-2h-1V9h-1V8H9v6Z"/>`
	menuRightFillInnerSVG                    = `<path fill="currentColor" d="M9 5H8v12h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H9"/>`
	menuTopLeftInnerSVG                      = `<path fill="currentColor" d="M18 7v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7V7h11m-4 2H9v5h1v-1h1v-1h1v-1h1v-1h1V9Z"/>`
	menuTopRightInnerSVG                     = `<path fill="currentColor" d="M15 18h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V7h11v11m-2-4V9H8v1h1v1h1v1h1v1h1v1h1Z"/>`
	menuUpInnerSVG                           = `<path fill="currentColor" d="M18 14H4v-2h1v-1h1v-1h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v1h1v1h1v2m-4-2v-1h-1v-1h-1V9h-2v1H9v1H8v1h6Z"/>`
	menuUpDownInnerSVG                       = `<path fill="currentColor" d="M4 10V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v2H4m0 2h14v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-2m4-4h6V7h-1V6h-1V5h-2v1H9v1H8v1m0 6v1h1v1h1v1h2v-1h1v-1h1v-1H8Z"/>`
	menuUpFillInnerSVG                       = `<path fill="currentColor" d="M5 13v1h12v-1h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-2v1H9v1H8v1H7v1H6v1"/>`
	messageInnerSVG                          = `<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m1 2v13h1v-1h15V3H3Z"/>`
	messageProcessingInnerSVG                = `<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m2 14h15V3H3v13h1v-1m2-7h2v2H6V8m4 0h2v2h-2V8m4 0h2v2h-2V8Z"/>`
	messageTextInnerSVG                      = `<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m1 2v13h1v-1h15V3H3m2 3h12v2H5V6m0 4h9v2H5v-2Z"/>`
	microphoneInnerSVG                       = `<path fill="currentColor" d="M10 21v-3H8v-1H6v-1H5v-1H4v-2H3V9h2v3h1v2h1v1h2v1h4v-1h2v-1h1v-2h1V9h2v4h-1v2h-1v1h-1v1h-2v1h-2v3h-2m-2-8v-1H7V3h1V2h1V1h4v1h1v1h1v9h-1v1h-1v1H9v-1H8m1-2h1v1h2v-1h1V4h-1V3h-2v1H9v7Z"/>`
	minusInnerSVG                            = `<path fill="currentColor" d="M17 12H5v-2h12Z"/>`
	minusBoxInnerSVG                         = `<path fill="currentColor" d="M16 12H6v-2h10Zm2 8H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`
	minusBoxFillInnerSVG                     = `<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-2-8v-2H6v2Z"/>`
	minusCircleInnerSVG                      = `<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-1 4v2H6v-2h10Z"/>`
	minusCircleFillInnerSVG                  = `<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2m1-8v-2H6v2Z"/>`
	monitorInnerSVG                          = `<path fill="currentColor" d="M2 2h18v1h1v12h-1v1h-7v2h2v2H7v-2h2v-2H2v-1H1V3h1V2m1 2v10h16V4H3Z"/>`
	monitorImageInnerSVG                     = `<path fill="currentColor" d="M14 6h2v1h1v2h-1v1h-2V9h-1V7h1V6M2 2h18v1h1v12h-1v1h-7v2h2v2H7v-2h2v-2H2v-1H1V3h1V2m1 2v5h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h4V4H3m7 7H9v-1H8V9H6v1H5v1H4v1H3v2h9v-1h-1v-1h-1v-1Z"/>`
	multiplyInnerSVG                         = `<path fill="currentColor" d="M16 16h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6v-2h1v-1h1v-1h1v-2H8V9H7V8H6V6h2v1h1v1h1v1h2V8h1V7h1V6h2v2h-1v1h-1v1h-1v2h1v1h1v1h1Z"/>`
	musicNoteInnerSVG                        = `<path fill="currentColor" d="M11 2h7v5h-5v11h-1v1h-1v1H7v-1H6v-1H5v-4h1v-1h1v-1h4V2m0 13h-1v-1H8v1H7v2h1v1h2v-1h1v-2Z"/>`
	necklaceInnerSVG                         = `<path fill="currentColor" d="M9 17h1v-1h2v1h1v2h-1v1h-2v-1H9v-2m1-2v-1H9v-1h-.91v-1H7v-2H6V8H5V6H4V3h1V2h12v1h1v3h-1v2h-1v2h-1v2h-1v1h-1v1h-1v1h-2M7 5v2h1v2h1v2h1.09v1H12v-1h1V9h1V7h1V5h1V4H6v1h1Z"/>`
	noteInnerSVG                             = `<path fill="currentColor" d="M15 3v1h1v1h1v1h1v1h1v1h1v1h1v9h-1v1H2v-1H1V4h1V3h13m0 3h-1v4h4V9h-1V8h-1V7h-1V6M3 5v12h16v-5h-6v-1h-1V5H3Z"/>`
	notebookInnerSVG                         = `<path fill="currentColor" d="M19 2v18h-1v1H4v-1H3v-2H1v-2h2v-4H1v-2h2V6H1V4h2V2h1V1h14v1h1m-5 7h-1V8h-1v1h-1v1h-1V3H7v16h10V3h-2v7h-1V9M3 4v2h2V4H3m2 6H3v2h2v-2m0 6H3v2h2v-2Z"/>`
	notificationInnerSVG                     = `<path fill="currentColor" d="M12 20v1h-2v-1H8v-2h6v2h-2M2 15h1v-1h1V6h1V4h1V3h2V2h2V1h2v1h2v1h2v1h1v2h1v8h1v1h1v2H2v-2m4 0h10V7h-1V5h-2V4H9v1H7v2H6v8Z"/>`
	octagonInnerSVG                          = `<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1V6m13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1v-1Z"/>`
	octagonAlertInnerSVG                     = `<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1V6m13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1v-1M10 6h2v7h-2V6m0 8h2v2h-2v-2Z"/>`
	pauseInnerSVG                            = `<path fill="currentColor" d="M6 4h3v14H6V4m7 14V4h3v14h-3Z"/>`
	pencilInnerSVG                           = `<path fill="currentColor" d="M16 2h1v1h1v1h1v1h1v1h-1v1h-1v1h-1V7h-1V6h-1V5h-1V4h1V3h1m-4 3h2v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H2v-4h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1"/>`
	pickaxeInnerSVG                          = `<path fill="currentColor" d="M8 2h3v1h2v1h2v1h2v2h1v2h1v2h1v3h-2v-2h-1v-1h-1v-1h-2V9h-1V8h-1V6h-1V5h-1V4H8m3 5h1v1h1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1H2v-1H1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1"/>`
	pictogrammersInnerSVG                    = `<path fill="currentColor" d="M4 0h14v1h1v1h1v18h-1v1h-1v1H4v-1H3v-1H2V2h1V1h1V0m0 18v1h1v1h12v-1h1v-1H4M17 2H5v1H4v12h1v1h12v-1h1V3h-1V2m-4 2v1h1v1h1v2h-1v1h-1v1h-3v4H8V4h5m-1 2h-2v2h2V6Z"/>`
	playInnerSVG                             = `<path fill="currentColor" d="M10 5v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7V4h2v1h1m2 5h-1V9h-1V8H9v6h1v-1h1v-1h1v-2Z"/>`
	plusInnerSVG                             = `<path fill="currentColor" d="M12 17h-2v-5H5v-2h5V5h2v5h5v2h-5Z"/>`
	plusBoxInnerSVG                          = `<path fill="currentColor" d="M12 16h-2v-4H6v-2h4V6h2v4h4v2h-4Zm6 4H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`
	plusBoxFillInnerSVG                      = `<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-6-4v-4h4v-2h-4V6h-2v4H6v2h4v4Z"/>`
	plusCircleInnerSVG                       = `<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-7 0h2v4h4v2h-4v4h-2v-4H6v-2h4V6Z"/>`
	plusCircleFillInnerSVG                   = `<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-3-5v-4h4v-2h-4V6h-2v4H6v2h4v4Z"/>`
	radioboxInnerSVG                         = `<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2V2m1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4H9Z"/>`
	radioboxMarkedInnerSVG                   = `<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2V2m1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4H9m4 3v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V9h1V8h1V7h4Z"/>`
	removeCircleInnerSVG                     = `<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-3 1v1h1v1h-1v1h-1v2h1v1h1v1h-1v1h-1v-1h-1v-1h-2v1H9v1H8v-1H7v-1h1v-1h1v-2H8V9H7V8h1V7h1v1h1v1h2V8h1V7h1Z"/>`
	rotateClockwiseInnerSVG                  = `<path fill="currentColor" d="M22 11v1h-1v1h-1v1h-1v1h-2v-1h-1v-1h-1v-1h-1v-1h3V9h-1V7h-1V6h-2V5H9v1H7v1H6v2H5v4h1v2h1v1h2v1h4v-1h3v2h-2v1H8v-1H6v-1H5v-1H4v-2H3V8h1V6h1V5h1V4h2V3h6v1h2v1h1v1h1v2h1v3h3Z"/>`
	rotateCounterclockwiseInnerSVG           = `<path fill="currentColor" d="M0 11v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1H5V9h1V7h1V6h2V5h4v1h2v1h1v2h1v4h-1v2h-1v1h-2v1H9v-1H6v2h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v3H0Z"/>`
	scriptInnerSVG                           = `<path fill="currentColor" d="M20 1H5v1H4v13h2V3h9v16h-2v-1h-1v-1H1v3h1v1h14v-1h1V3h2v2h2V2h-1"/>`
	shieldInnerSVG                           = `<path fill="currentColor" d="M3 4h2V3h2V2h2V1h4v1h2v1h2v1h2v10h-1v2h-1v2h-1v1h-1v1h-2v1H9v-1H7v-1H6v-1H5v-2H4v-2H3V4m7-1v1H8v1H6v1H5v7h1v2h1v2h1v1h2v1h2v-1h2v-1h1v-2h1v-2h1V6h-1V5h-2V4h-2V3h-2Z"/>`
	skullInnerSVG                            = `<path fill="currentColor" d="M6 2h2V1h6v1h2v1h1v1h1v1h1v2h1v7h-1v2h-1v4h-1v1H5v-1H4v-4H3v-2H2V8h1V5h1V4h1V3h1V2m9 3V4h-2V3H9v1H7v1H6v1H5v3H4v4h1v2h1v4h2v-2h2v2h2v-2h2v2h2v-4h1v-2h1V8h-1V6h-1V5h-1M7 8h3v3H7V8m5 3V8h3v3h-3m-2 2h2v2h-2v-2Z"/>`
	speakerInnerSVG                          = `<path fill="currentColor" d="M4 1h14v1h1v18h-1v1H4v-1H3V2h1V1m1 2v16h12V3H5m4 2h1V4h2v1h1v2h-1v1h-2V7H9V5m0 13v-1H8v-1H7v-4h1v-1h1v-1h4v1h1v1h1v4h-1v1h-1v1H9m1-5H9v2h1v1h2v-1h1v-2h-1v-1h-2v1Z"/>`
	stopInnerSVG                             = `<path fill="currentColor" d="M16 5v1h1v10h-1v1H6v-1H5V6h1V5h10M7 7v8h8V7H7Z"/>`
	swordInnerSVG                            = `<path fill="currentColor" d="M8 3v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-1h2v2h-1v1h-1v1h1v1h1v1h1v2h-1v1h-2v-1h-1v-1h-1v-1h-1v1h-1v1h-2v-2h1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V8H3V7H2V2h5v1h1M7 5H6V4H4v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h-1v-1h-1v-1h-1V9h-1V8H9V7H8V6H7V5Z"/>`
	tagInnerSVG                              = `<path fill="currentColor" d="M1 2h1V1h9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V2m2 8h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H3v7m2-6h2v1h1v2H7v1H5V7H4V5h1V4Z"/>`
	tagTextInnerSVG                          = `<path fill="currentColor" d="M1 2h1V1h9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V2m2 8h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H3v7m11 1h1v2h-1v-1h-1v-1h-1v-1h-1V9h-1V7h1v1h1v1h1v1h1v1m-4 1h1v1h1v2h-1v-1h-1v-1H9v-1H8v-2h1v1h1v1M5 4h2v1h1v2H7v1H5V7H4V5h1V4Z"/>`
	targetInnerSVG                           = `<path fill="currentColor" d="M12 13h-2v-1H9v-2h1V9h2v1h1v2h-1Zm2 4H8v-1H7v-1H6v-1H5V8h1V7h1V6h1V5h6v1h1v1h1v1h1v6h-1v1h-1v1h-1Zm1 4H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-2-6v-1h1v-1h1V9h-1V8h-1V7H9v1H8v1H7v4h1v1h1v1Zm1 4v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1Z"/>`
	terminalInnerSVG                         = `<path fill="currentColor" d="M20 3v16h-1v1H3v-1H2V3h1V2h16v1h1m-2 3H4v12h14V6M9 9v1h1v1h1v2h-1v1H9v1H8v1H6v-2h1v-1h1v-2H7v-1H6V8h2v1h1m2 7v-2h5v2h-5Z"/>`
	textBoxInnerSVG                          = `<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m13 3V4H5v1H4v12h1v1h12v-1h1V5h-1M6 8h10v2H6V8m0 4h7v2H6v-2Z"/>`
	textImageInnerSVG                        = `<path fill="currentColor" d="M2 4h10v10H2V4m2 2v6h6V6H4m10-2h6v2h-6V4m0 4h6v2h-6V8m0 4h6v2h-6v-2M2 16h16v2H2v-2Z"/>`
	toggleSwitchOffInnerSVG                  = `<path fill="currentColor" d="M5 8h4v1h1v4H9v1H5v-1H4V9h1V8m14-3v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5h16m-1 2H4v1H3v6h1v1h14v-1h1V8h-1V7Z"/>`
	toggleSwitchOnInnerSVG                   = `<path fill="currentColor" d="M3 5h16v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5m10 3v1h-1v4h1v1h4v-1h1V9h-1V8h-4Z"/>`
	toolboxInnerSVG                          = `<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1V6m7 0h4V4H9v2m10 2H3v4h3v-2h3v2h4v-2h3v2h3V8M3 18h16v-4h-3v2h-3v-2H9v2H6v-2H3v4Z"/>`
	tooltipAboveInnerSVG                     = `<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3Z"/>`
	tooltipAboveAlertInnerSVG                = `<path fill="currentColor" d="M10 5h2v5h-2V5m0 6h2v2h-2v-2M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3Z"/>`
	tooltipAboveTextInnerSVG                 = `<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3m2 3h12v2H5V6m0 4h10v2H5v-2Z"/>`
	tooltipBelowInnerSVG                     = `<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3Z"/>`
	tooltipBelowAlertInnerSVG                = `<path fill="currentColor" d="M10 17h2v-2h-2v2m0-3h2V9h-2v5m-8 7h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3Z"/>`
	tooltipBelowTextInnerSVG                 = `<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3m2-3h10v-2H5v2m0-4h12v-2H5v2Z"/>`
	tooltipEndInnerSVG                       = `<path fill="currentColor" d="M21 20V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1h1m-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12v16Z"/>`
	tooltipEndAlertInnerSVG                  = `<path fill="currentColor" d="M14 15h-2v-2h2v2m0-3h-2V7h2v5m7-10v18h-1v1H6v-1H5v-5H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V2h1V1h14v1h1m-2 1H7v5H6v1H5v1H4v2h1v1h1v1h1v5h12V3Z"/>`
	tooltipEndTextInnerSVG                   = `<path fill="currentColor" d="M17 8V6H9v2h8m-2 8v-2H9v2h6m2-4v-2H9v2h8m4 8V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1h1m-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12v16Z"/>`
	tooltipStartInnerSVG                     = `<path fill="currentColor" d="M1 20V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1H1m2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3v16Z"/>`
	tooltipStartAlertInnerSVG                = `<path fill="currentColor" d="M8 15h2v-2H8v2m0-3h2V7H8v5M1 2v18h1v1h14v-1h1v-5h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V2h-1V1H2v1H1m2 1h12v5h1v1h1v1h1v2h-1v1h-1v1h-1v5H3V3Z"/>`
	tooltipStartTextInnerSVG                 = `<path fill="currentColor" d="M5 8V6h8v2H5m0 8v-2h6v2H5m0-4v-2h8v2H5m-4 8V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1H1m2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3v16Z"/>`
	trashInnerSVG                            = `<path fill="currentColor" d="M10 7v9H8V7h2m2 0h2v9h-2V7M8 2h6v1h5v2h-1v14h-1v1H5v-1H4V5H3V3h5V2M6 5v13h10V5H6Z"/>`
	uploadInnerSVG                           = `<path fill="currentColor" d="M18 17v2H4v-2h14M8 15V9H4V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h-4v6H8m2-2h2V7h1V6h-1V5h-2v1H9v1h1v6Z"/>`
	volumeHighInnerSVG                       = `<path fill="currentColor" d="M13 2h2v1h1v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-1V4h-1V2m1 5v1h1v2h1v2h-1v2h-1v1h-1V7h1M2 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1H9v-1H8v-1H7v-1H6v-1H2V8m2 2v2h3v1h1v1h1V8H8v1H7v1H4Z"/>`
	volumeLowInnerSVG                        = `<path fill="currentColor" d="M6 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H6V8m2 2v2h3v1h1v1h1V8h-1v1h-1v1H8Z"/>`
	volumeMediumInnerSVG                     = `<path fill="currentColor" d="M16 7v1h1v2h1v2h-1v2h-1v1h-1V7h1M8 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1H9v-1H8v-1H4V8h4m-2 2v2h3v1h1v1h1V8h-1v1H9v1H6Z"/>`
	volumeMuteInnerSVG                       = `<path fill="currentColor" d="M12 7h2v1h1v1h2V8h1V7h2v2h-1v1h-1v2h1v1h1v2h-2v-1h-1v-1h-2v1h-1v1h-2v-2h1v-1h1v-2h-1V9h-1V7M6 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1H9v-1H8v-1H7v-1H6v-1H2V8h4m1 2H4v2h3v1h1v1h1V8H8v1H7v1Z"/>`
	wallInnerSVG                             = `<path fill="currentColor" d="M19 20H4v-5H1V7h3V2h15v5h2v8h-2ZM12 7V4H6v3Zm5 0V4h-3v3Zm-9 6V9H3v4Zm11 0V9h-9v4Zm-8 5v-3H6v3Zm6 0v-3h-4v3Z"/>`
	wallFillInnerSVG                         = `<path fill="currentColor" d="M5 3h6v4H5m8-4h5v4h-5m-8 8v4h6v-4m2 0h5v4h-5M8 9H3v4h5m2-4h10v4H10"/>`
	waterInnerSVG                            = `<path fill="currentColor" d="M14 21H8v-1H6v-1H5v-1H4v-2H3v-3h1v-2h1V9h1V7h1V6h1V4h1V3h1V1h2v2h1v1h1v2h1v1h1v2h1v2h1v2h1v3h-1v2h-1v1h-1v1h-2Z"/>`
	wellInnerSVG                             = `<path fill="currentColor" d="M19 21H3v-5H1v-2h20v2h-2v5M5 16v3h12v-3H5M2 7V5h1V4h1V3h1V2h1V1h10v1h1v1h1v1h1v1h1v2h-2v6h-2V7h-4v2h2v4H8V9h2V7H6v6H4V7H2m5-4v1H6v1h10V4h-1V3H7Z"/>`
)

var sharedIconAttrs = []engine.Attributer{
	engine.NewAttribute("width", "1em"),
	engine.NewAttribute("height", "1em"),
}

func Account(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accountInnerSVG).
		Element(children...)
}

func AccountBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(accountBoxInnerSVG).
		Element(children...)
}

func Alert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertInnerSVG).
		Element(children...)
}

func AlertBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertBoxInnerSVG).
		Element(children...)
}

func AlertBoxFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertBoxFillInnerSVG).
		Element(children...)
}

func AlertCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertCircleInnerSVG).
		Element(children...)
}

func AlertCircleFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertCircleFillInnerSVG).
		Element(children...)
}

func AlertRhombus(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertRhombusInnerSVG).
		Element(children...)
}

func AlertRhombusFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alertRhombusFillInnerSVG).
		Element(children...)
}

func AlphaA(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaAInnerSVG).
		Element(children...)
}

func AlphaAFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaAFillInnerSVG).
		Element(children...)
}

func AlphaB(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaBInnerSVG).
		Element(children...)
}

func AlphaBFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaBFillInnerSVG).
		Element(children...)
}

func AlphaC(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaCInnerSVG).
		Element(children...)
}

func AlphaCFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaCFillInnerSVG).
		Element(children...)
}

func AlphaD(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaDInnerSVG).
		Element(children...)
}

func AlphaDFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaDFillInnerSVG).
		Element(children...)
}

func AlphaE(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaEInnerSVG).
		Element(children...)
}

func AlphaEFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaEFillInnerSVG).
		Element(children...)
}

func AlphaF(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaFInnerSVG).
		Element(children...)
}

func AlphaFFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaFFillInnerSVG).
		Element(children...)
}

func AlphaG(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaGInnerSVG).
		Element(children...)
}

func AlphaGFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaGFillInnerSVG).
		Element(children...)
}

func AlphaH(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaHInnerSVG).
		Element(children...)
}

func AlphaHFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaHFillInnerSVG).
		Element(children...)
}

func AlphaI(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaIInnerSVG).
		Element(children...)
}

func AlphaIFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaIFillInnerSVG).
		Element(children...)
}

func AlphaJ(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaJInnerSVG).
		Element(children...)
}

func AlphaJFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaJFillInnerSVG).
		Element(children...)
}

func AlphaK(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaKInnerSVG).
		Element(children...)
}

func AlphaKFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaKFillInnerSVG).
		Element(children...)
}

func AlphaL(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaLInnerSVG).
		Element(children...)
}

func AlphaLFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaLFillInnerSVG).
		Element(children...)
}

func AlphaM(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaMInnerSVG).
		Element(children...)
}

func AlphaMFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaMFillInnerSVG).
		Element(children...)
}

func AlphaN(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaNInnerSVG).
		Element(children...)
}

func AlphaNFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaNFillInnerSVG).
		Element(children...)
}

func AlphaO(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaOInnerSVG).
		Element(children...)
}

func AlphaOFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaOFillInnerSVG).
		Element(children...)
}

func AlphaP(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaPInnerSVG).
		Element(children...)
}

func AlphaPFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaPFillInnerSVG).
		Element(children...)
}

func AlphaQ(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaQInnerSVG).
		Element(children...)
}

func AlphaQFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaQFillInnerSVG).
		Element(children...)
}

func AlphaR(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaRInnerSVG).
		Element(children...)
}

func AlphaRFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaRFillInnerSVG).
		Element(children...)
}

func AlphaS(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaSInnerSVG).
		Element(children...)
}

func AlphaSFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaSFillInnerSVG).
		Element(children...)
}

func AlphaT(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaTInnerSVG).
		Element(children...)
}

func AlphaTFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaTFillInnerSVG).
		Element(children...)
}

func AlphaU(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaUInnerSVG).
		Element(children...)
}

func AlphaUFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaUFillInnerSVG).
		Element(children...)
}

func AlphaV(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaVInnerSVG).
		Element(children...)
}

func AlphaVFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaVFillInnerSVG).
		Element(children...)
}

func AlphaW(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaWInnerSVG).
		Element(children...)
}

func AlphaWFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaWFillInnerSVG).
		Element(children...)
}

func AlphaX(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaXInnerSVG).
		Element(children...)
}

func AlphaXFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaXFillInnerSVG).
		Element(children...)
}

func AlphaY(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaYInnerSVG).
		Element(children...)
}

func AlphaYFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaYFillInnerSVG).
		Element(children...)
}

func AlphaZ(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaZInnerSVG).
		Element(children...)
}

func AlphaZFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(alphaZFillInnerSVG).
		Element(children...)
}

func Application(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(applicationInnerSVG).
		Element(children...)
}

func ApplicationCode(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(applicationCodeInnerSVG).
		Element(children...)
}

func Archive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(archiveInnerSVG).
		Element(children...)
}

func Arrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowInnerSVG).
		Element(children...)
}

func ArrowBottomLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBottomLeftInnerSVG).
		Element(children...)
}

func ArrowBottomLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBottomLeftCircleInnerSVG).
		Element(children...)
}

func ArrowBottomRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBottomRightInnerSVG).
		Element(children...)
}

func ArrowBottomRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowBottomRightCircleInnerSVG).
		Element(children...)
}

func ArrowDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownInnerSVG).
		Element(children...)
}

func ArrowDownBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownBoldInnerSVG).
		Element(children...)
}

func ArrowDownCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowDownCircleInnerSVG).
		Element(children...)
}

func ArrowDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftInnerSVG).
		Element(children...)
}

func ArrowLeftBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftBoldInnerSVG).
		Element(children...)
}

func ArrowLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftCircleInnerSVG).
		Element(children...)
}

func ArrowLeftDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftDownInnerSVG).
		Element(children...)
}

func ArrowLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftRightInnerSVG).
		Element(children...)
}

func ArrowLeftUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowLeftUpInnerSVG).
		Element(children...)
}

func ArrowRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightInnerSVG).
		Element(children...)
}

func ArrowRightBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightBoldInnerSVG).
		Element(children...)
}

func ArrowRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightCircleInnerSVG).
		Element(children...)
}

func ArrowRightDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightDownInnerSVG).
		Element(children...)
}

func ArrowRightUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowRightUpInnerSVG).
		Element(children...)
}

func ArrowTopLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTopLeftInnerSVG).
		Element(children...)
}

func ArrowTopLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTopLeftCircleInnerSVG).
		Element(children...)
}

func ArrowTopRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTopRightInnerSVG).
		Element(children...)
}

func ArrowTopRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowTopRightCircleInnerSVG).
		Element(children...)
}

func ArrowUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpInnerSVG).
		Element(children...)
}

func ArrowUpBold(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpBoldInnerSVG).
		Element(children...)
}

func ArrowUpCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpCircleInnerSVG).
		Element(children...)
}

func ArrowUpDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpDownInnerSVG).
		Element(children...)
}

func ArrowUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(arrowUpRightInnerSVG).
		Element(children...)
}

func Axe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(axeInnerSVG).
		Element(children...)
}

func BagPersonal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagPersonalInnerSVG).
		Element(children...)
}

func BagPersonalFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bagPersonalFillInnerSVG).
		Element(children...)
}

func BatteryFifty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryFiftyInnerSVG).
		Element(children...)
}

func BatteryOneHundred(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryOneHundredInnerSVG).
		Element(children...)
}

func BatterySeventyFive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batterySeventyFiveInnerSVG).
		Element(children...)
}

func BatteryTwentyFive(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryTwentyFiveInnerSVG).
		Element(children...)
}

func BatteryZero(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(batteryZeroInnerSVG).
		Element(children...)
}

func BattleAxe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(battleAxeInnerSVG).
		Element(children...)
}

func Book(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bookmarkInnerSVG).
		Element(children...)
}

func BorderBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBottomInnerSVG).
		Element(children...)
}

func BorderBottomLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBottomLeftInnerSVG).
		Element(children...)
}

func BorderBottomLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBottomLeftRightInnerSVG).
		Element(children...)
}

func BorderBottomRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderBottomRightInnerSVG).
		Element(children...)
}

func BorderInside(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderInsideInnerSVG).
		Element(children...)
}

func BorderLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderLeftInnerSVG).
		Element(children...)
}

func BorderLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderLeftRightInnerSVG).
		Element(children...)
}

func BorderNone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderNoneInnerSVG).
		Element(children...)
}

func BorderOutside(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderOutsideInnerSVG).
		Element(children...)
}

func BorderRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderRightInnerSVG).
		Element(children...)
}

func BorderTop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopInnerSVG).
		Element(children...)
}

func BorderTopBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopBottomInnerSVG).
		Element(children...)
}

func BorderTopLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopLeftInnerSVG).
		Element(children...)
}

func BorderTopLeftBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopLeftBottomInnerSVG).
		Element(children...)
}

func BorderTopLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopLeftRightInnerSVG).
		Element(children...)
}

func BorderTopRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopRightInnerSVG).
		Element(children...)
}

func BorderTopRightBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(borderTopRightBottomInnerSVG).
		Element(children...)
}

func Bow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bowInnerSVG).
		Element(children...)
}

func BowArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bowArrowInnerSVG).
		Element(children...)
}

func Box(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxInnerSVG).
		Element(children...)
}

func BoxLightDashedDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDashedDownLeftInnerSVG).
		Element(children...)
}

func BoxLightDashedDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDashedDownRightInnerSVG).
		Element(children...)
}

func BoxLightDashedUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDashedUpLeftInnerSVG).
		Element(children...)
}

func BoxLightDashedUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDashedUpRightInnerSVG).
		Element(children...)
}

func BoxLightDoubleHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDoubleHorizontalInnerSVG).
		Element(children...)
}

func BoxLightDoubleRoundDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDoubleRoundDownLeftInnerSVG).
		Element(children...)
}

func BoxLightDoubleRoundDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDoubleRoundDownRightInnerSVG).
		Element(children...)
}

func BoxLightDoubleRoundUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDoubleRoundUpLeftInnerSVG).
		Element(children...)
}

func BoxLightDoubleRoundUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDoubleRoundUpRightInnerSVG).
		Element(children...)
}

func BoxLightDoubleVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDoubleVerticalInnerSVG).
		Element(children...)
}

func BoxLightDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDownLeftInnerSVG).
		Element(children...)
}

func BoxLightDownLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDownLeftCircleInnerSVG).
		Element(children...)
}

func BoxLightDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDownRightInnerSVG).
		Element(children...)
}

func BoxLightDownRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightDownRightCircleInnerSVG).
		Element(children...)
}

func BoxLightHorizontal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightHorizontalInnerSVG).
		Element(children...)
}

func BoxLightRoundDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightRoundDownLeftInnerSVG).
		Element(children...)
}

func BoxLightRoundDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightRoundDownRightInnerSVG).
		Element(children...)
}

func BoxLightRoundUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightRoundUpLeftInnerSVG).
		Element(children...)
}

func BoxLightRoundUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightRoundUpRightInnerSVG).
		Element(children...)
}

func BoxLightUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightUpLeftInnerSVG).
		Element(children...)
}

func BoxLightUpLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightUpLeftCircleInnerSVG).
		Element(children...)
}

func BoxLightUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightUpRightInnerSVG).
		Element(children...)
}

func BoxLightUpRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightUpRightCircleInnerSVG).
		Element(children...)
}

func BoxLightVertical(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxLightVerticalInnerSVG).
		Element(children...)
}

func BoxOuterLightAll(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightAllInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedAll(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedAllInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedDownInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedDownLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedDownLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedDownLeftRightInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedDownRightInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedLeftRightInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedRightInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedUpInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedUpDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedUpDownInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedUpDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedUpDownLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedUpDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedUpDownRightInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedUpLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedUpLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedUpLeftRightInnerSVG).
		Element(children...)
}

func BoxOuterLightDashedUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDashedUpRightInnerSVG).
		Element(children...)
}

func BoxOuterLightDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDownInnerSVG).
		Element(children...)
}

func BoxOuterLightDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDownLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightDownLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDownLeftRightInnerSVG).
		Element(children...)
}

func BoxOuterLightDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightDownRightInnerSVG).
		Element(children...)
}

func BoxOuterLightLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightLeftRightInnerSVG).
		Element(children...)
}

func BoxOuterLightRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightRightInnerSVG).
		Element(children...)
}

func BoxOuterLightRoundDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightRoundDownLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightRoundDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightRoundDownRightInnerSVG).
		Element(children...)
}

func BoxOuterLightRoundUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightRoundUpLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightRoundUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightRoundUpRightInnerSVG).
		Element(children...)
}

func BoxOuterLightUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightUpInnerSVG).
		Element(children...)
}

func BoxOuterLightUpDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightUpDownInnerSVG).
		Element(children...)
}

func BoxOuterLightUpDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightUpDownLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightUpDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightUpDownRightInnerSVG).
		Element(children...)
}

func BoxOuterLightUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightUpLeftInnerSVG).
		Element(children...)
}

func BoxOuterLightUpLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightUpLeftRightInnerSVG).
		Element(children...)
}

func BoxOuterLightUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(boxOuterLightUpRightInnerSVG).
		Element(children...)
}

func Briefcase(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bugInnerSVG).
		Element(children...)
}

func BugFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(bugFillInnerSVG).
		Element(children...)
}

func Calculator(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarInnerSVG).
		Element(children...)
}

func CalendarMonth(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(calendarMonthInnerSVG).
		Element(children...)
}

func Cancel(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cancelInnerSVG).
		Element(children...)
}

func Card(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardInnerSVG).
		Element(children...)
}

func CardText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cardTextInnerSVG).
		Element(children...)
}

func Cart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cartInnerSVG).
		Element(children...)
}

func ChartBar(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chartBarInnerSVG).
		Element(children...)
}

func Chat(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chatInnerSVG).
		Element(children...)
}

func Check(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkInnerSVG).
		Element(children...)
}

func CheckboxBlank(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxBlankInnerSVG).
		Element(children...)
}

func CheckboxCross(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxCrossInnerSVG).
		Element(children...)
}

func CheckboxMarked(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(checkboxMarkedInnerSVG).
		Element(children...)
}

func ChevronDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronDownInnerSVG).
		Element(children...)
}

func ChevronDownCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronDownCircleInnerSVG).
		Element(children...)
}

func ChevronLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronLeftInnerSVG).
		Element(children...)
}

func ChevronLeftCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronLeftCircleInnerSVG).
		Element(children...)
}

func ChevronRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightInnerSVG).
		Element(children...)
}

func ChevronRightCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronRightCircleInnerSVG).
		Element(children...)
}

func ChevronUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpInnerSVG).
		Element(children...)
}

func ChevronUpCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(chevronUpCircleInnerSVG).
		Element(children...)
}

func Circle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(circleInnerSVG).
		Element(children...)
}

func Clipboard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(clockInnerSVG).
		Element(children...)
}

func Coffee(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(coffeeInnerSVG).
		Element(children...)
}

func Comment(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentInnerSVG).
		Element(children...)
}

func CommentText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(commentTextInnerSVG).
		Element(children...)
}

func Compass(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassInnerSVG).
		Element(children...)
}

func CompassEastArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassEastArrowInnerSVG).
		Element(children...)
}

func CompassNorthArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassNorthArrowInnerSVG).
		Element(children...)
}

func CompassNorthEast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassNorthEastInnerSVG).
		Element(children...)
}

func CompassNorthWest(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassNorthWestInnerSVG).
		Element(children...)
}

func CompassSouthArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassSouthArrowInnerSVG).
		Element(children...)
}

func CompassSouthEast(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassSouthEastInnerSVG).
		Element(children...)
}

func CompassSouthWest(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassSouthWestInnerSVG).
		Element(children...)
}

func CompassWestArrow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(compassWestArrowInnerSVG).
		Element(children...)
}

func CreditCard(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(creditCardInnerSVG).
		Element(children...)
}

func Crown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(crownInnerSVG).
		Element(children...)
}

func CubeUnfolded(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(cubeUnfoldedInnerSVG).
		Element(children...)
}

func Database(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(databaseInnerSVG).
		Element(children...)
}

func Device(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(diamondInnerSVG).
		Element(children...)
}

func Division(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(divisionInnerSVG).
		Element(children...)
}

func Door(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doorInnerSVG).
		Element(children...)
}

func DoorBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doorBoxInnerSVG).
		Element(children...)
}

func DoorOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(doorOpenInnerSVG).
		Element(children...)
}

func Download(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(downloadInnerSVG).
		Element(children...)
}

func Email(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(emailInnerSVG).
		Element(children...)
}

func File(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fileInnerSVG).
		Element(children...)
}

func Fire(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(fireInnerSVG).
		Element(children...)
}

func Flask(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flaskInnerSVG).
		Element(children...)
}

func FlaskEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flaskEmptyInnerSVG).
		Element(children...)
}

func FlaskRoundBottom(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flaskRoundBottomInnerSVG).
		Element(children...)
}

func FlaskRoundBottomEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(flaskRoundBottomEmptyInnerSVG).
		Element(children...)
}

func FloppyDisk(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderInnerSVG).
		Element(children...)
}

func FolderOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(folderOpenInnerSVG).
		Element(children...)
}

func GamepadCenter(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadCenterInnerSVG).
		Element(children...)
}

func GamepadDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadDownInnerSVG).
		Element(children...)
}

func GamepadDownLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadDownLeftInnerSVG).
		Element(children...)
}

func GamepadDownRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadDownRightInnerSVG).
		Element(children...)
}

func GamepadEmpty(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadEmptyInnerSVG).
		Element(children...)
}

func GamepadLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadLeftInnerSVG).
		Element(children...)
}

func GamepadRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadRightInnerSVG).
		Element(children...)
}

func GamepadUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadUpInnerSVG).
		Element(children...)
}

func GamepadUpLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadUpLeftInnerSVG).
		Element(children...)
}

func GamepadUpRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(gamepadUpRightInnerSVG).
		Element(children...)
}

func Heart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(heartInnerSVG).
		Element(children...)
}

func Image(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(imageInnerSVG).
		Element(children...)
}

func Label(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(labelInnerSVG).
		Element(children...)
}

func LabelVariant(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(labelVariantInnerSVG).
		Element(children...)
}

func Led(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(ledInnerSVG).
		Element(children...)
}

func Lightbulb(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lightbulbInnerSVG).
		Element(children...)
}

func Lock(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockInnerSVG).
		Element(children...)
}

func LockOpen(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(lockOpenInnerSVG).
		Element(children...)
}

func Login(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(logoutInnerSVG).
		Element(children...)
}

func Map(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(mapInnerSVG).
		Element(children...)
}

func MenuBottomLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuBottomLeftInnerSVG).
		Element(children...)
}

func MenuBottomRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuBottomRightInnerSVG).
		Element(children...)
}

func MenuDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuDownInnerSVG).
		Element(children...)
}

func MenuDownFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuDownFillInnerSVG).
		Element(children...)
}

func MenuLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuLeftInnerSVG).
		Element(children...)
}

func MenuLeftFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuLeftFillInnerSVG).
		Element(children...)
}

func MenuLeftRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuLeftRightInnerSVG).
		Element(children...)
}

func MenuRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuRightInnerSVG).
		Element(children...)
}

func MenuRightFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuRightFillInnerSVG).
		Element(children...)
}

func MenuTopLeft(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuTopLeftInnerSVG).
		Element(children...)
}

func MenuTopRight(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuTopRightInnerSVG).
		Element(children...)
}

func MenuUp(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuUpInnerSVG).
		Element(children...)
}

func MenuUpDown(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuUpDownInnerSVG).
		Element(children...)
}

func MenuUpFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(menuUpFillInnerSVG).
		Element(children...)
}

func Message(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageInnerSVG).
		Element(children...)
}

func MessageProcessing(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageProcessingInnerSVG).
		Element(children...)
}

func MessageText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(messageTextInnerSVG).
		Element(children...)
}

func Microphone(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusInnerSVG).
		Element(children...)
}

func MinusBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusBoxInnerSVG).
		Element(children...)
}

func MinusBoxFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusBoxFillInnerSVG).
		Element(children...)
}

func MinusCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircleInnerSVG).
		Element(children...)
}

func MinusCircleFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(minusCircleFillInnerSVG).
		Element(children...)
}

func Monitor(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(monitorInnerSVG).
		Element(children...)
}

func MonitorImage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(monitorImageInnerSVG).
		Element(children...)
}

func Multiply(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(multiplyInnerSVG).
		Element(children...)
}

func MusicNote(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(musicNoteInnerSVG).
		Element(children...)
}

func Necklace(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(necklaceInnerSVG).
		Element(children...)
}

func Note(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(noteInnerSVG).
		Element(children...)
}

func Notebook(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notebookInnerSVG).
		Element(children...)
}

func Notification(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(notificationInnerSVG).
		Element(children...)
}

func Octagon(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(octagonInnerSVG).
		Element(children...)
}

func OctagonAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(octagonAlertInnerSVG).
		Element(children...)
}

func Pause(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pauseInnerSVG).
		Element(children...)
}

func Pencil(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pencilInnerSVG).
		Element(children...)
}

func Pickaxe(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pickaxeInnerSVG).
		Element(children...)
}

func Pictogrammers(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(pictogrammersInnerSVG).
		Element(children...)
}

func Play(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
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
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusInnerSVG).
		Element(children...)
}

func PlusBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusBoxInnerSVG).
		Element(children...)
}

func PlusBoxFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusBoxFillInnerSVG).
		Element(children...)
}

func PlusCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircleInnerSVG).
		Element(children...)
}

func PlusCircleFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(plusCircleFillInnerSVG).
		Element(children...)
}

func Radiobox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radioboxInnerSVG).
		Element(children...)
}

func RadioboxMarked(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(radioboxMarkedInnerSVG).
		Element(children...)
}

func RemoveCircle(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(removeCircleInnerSVG).
		Element(children...)
}

func RotateClockwise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateClockwiseInnerSVG).
		Element(children...)
}

func RotateCounterclockwise(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(rotateCounterclockwiseInnerSVG).
		Element(children...)
}

func Script(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(scriptInnerSVG).
		Element(children...)
}

func Shield(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(shieldInnerSVG).
		Element(children...)
}

func Skull(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(skullInnerSVG).
		Element(children...)
}

func Speaker(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(speakerInnerSVG).
		Element(children...)
}

func Stop(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(stopInnerSVG).
		Element(children...)
}

func Sword(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(swordInnerSVG).
		Element(children...)
}

func Tag(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagInnerSVG).
		Element(children...)
}

func TagText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tagTextInnerSVG).
		Element(children...)
}

func Target(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(targetInnerSVG).
		Element(children...)
}

func Terminal(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(terminalInnerSVG).
		Element(children...)
}

func TextBox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textBoxInnerSVG).
		Element(children...)
}

func TextImage(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(textImageInnerSVG).
		Element(children...)
}

func ToggleSwitchOff(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleSwitchOffInnerSVG).
		Element(children...)
}

func ToggleSwitchOn(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toggleSwitchOnInnerSVG).
		Element(children...)
}

func Toolbox(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(toolboxInnerSVG).
		Element(children...)
}

func TooltipAbove(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipAboveInnerSVG).
		Element(children...)
}

func TooltipAboveAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipAboveAlertInnerSVG).
		Element(children...)
}

func TooltipAboveText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipAboveTextInnerSVG).
		Element(children...)
}

func TooltipBelow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipBelowInnerSVG).
		Element(children...)
}

func TooltipBelowAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipBelowAlertInnerSVG).
		Element(children...)
}

func TooltipBelowText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipBelowTextInnerSVG).
		Element(children...)
}

func TooltipEnd(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipEndInnerSVG).
		Element(children...)
}

func TooltipEndAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipEndAlertInnerSVG).
		Element(children...)
}

func TooltipEndText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipEndTextInnerSVG).
		Element(children...)
}

func TooltipStart(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipStartInnerSVG).
		Element(children...)
}

func TooltipStartAlert(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipStartAlertInnerSVG).
		Element(children...)
}

func TooltipStartText(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(tooltipStartTextInnerSVG).
		Element(children...)
}

func Trash(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(trashInnerSVG).
		Element(children...)
}

func Upload(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(uploadInnerSVG).
		Element(children...)
}

func VolumeHigh(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeHighInnerSVG).
		Element(children...)
}

func VolumeLow(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeLowInnerSVG).
		Element(children...)
}

func VolumeMedium(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeMediumInnerSVG).
		Element(children...)
}

func VolumeMute(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(volumeMuteInnerSVG).
		Element(children...)
}

func Wall(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wallInnerSVG).
		Element(children...)
}

func WallFill(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wallFillInnerSVG).
		Element(children...)
}

func Water(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(waterInnerSVG).
		Element(children...)
}

func Well(children ...*engine.UberElement) *engine.UberElement {
	attrs := append(
		[]engine.Attributer{
			engine.NewAttribute("viewBox", "0 0 22 22"),
		},
		sharedIconAttrs...,
	)
	return engine.Uber("svg").
		Attr(attrs...).
		HTML(wellInnerSVG).
		Element(children...)
}

func ByName(name string) (*engine.UberElement, error) {
	switch name {
	case "account":
		return Account(), nil
	case "account-box":
		return AccountBox(), nil
	case "alert":
		return Alert(), nil
	case "alert-box":
		return AlertBox(), nil
	case "alert-box-fill":
		return AlertBoxFill(), nil
	case "alert-circle":
		return AlertCircle(), nil
	case "alert-circle-fill":
		return AlertCircleFill(), nil
	case "alert-rhombus":
		return AlertRhombus(), nil
	case "alert-rhombus-fill":
		return AlertRhombusFill(), nil
	case "alpha-a":
		return AlphaA(), nil
	case "alpha-a-fill":
		return AlphaAFill(), nil
	case "alpha-b":
		return AlphaB(), nil
	case "alpha-b-fill":
		return AlphaBFill(), nil
	case "alpha-c":
		return AlphaC(), nil
	case "alpha-c-fill":
		return AlphaCFill(), nil
	case "alpha-d":
		return AlphaD(), nil
	case "alpha-d-fill":
		return AlphaDFill(), nil
	case "alpha-e":
		return AlphaE(), nil
	case "alpha-e-fill":
		return AlphaEFill(), nil
	case "alpha-f":
		return AlphaF(), nil
	case "alpha-f-fill":
		return AlphaFFill(), nil
	case "alpha-g":
		return AlphaG(), nil
	case "alpha-g-fill":
		return AlphaGFill(), nil
	case "alpha-h":
		return AlphaH(), nil
	case "alpha-h-fill":
		return AlphaHFill(), nil
	case "alpha-i":
		return AlphaI(), nil
	case "alpha-i-fill":
		return AlphaIFill(), nil
	case "alpha-j":
		return AlphaJ(), nil
	case "alpha-j-fill":
		return AlphaJFill(), nil
	case "alpha-k":
		return AlphaK(), nil
	case "alpha-k-fill":
		return AlphaKFill(), nil
	case "alpha-l":
		return AlphaL(), nil
	case "alpha-l-fill":
		return AlphaLFill(), nil
	case "alpha-m":
		return AlphaM(), nil
	case "alpha-m-fill":
		return AlphaMFill(), nil
	case "alpha-n":
		return AlphaN(), nil
	case "alpha-n-fill":
		return AlphaNFill(), nil
	case "alpha-o":
		return AlphaO(), nil
	case "alpha-o-fill":
		return AlphaOFill(), nil
	case "alpha-p":
		return AlphaP(), nil
	case "alpha-p-fill":
		return AlphaPFill(), nil
	case "alpha-q":
		return AlphaQ(), nil
	case "alpha-q-fill":
		return AlphaQFill(), nil
	case "alpha-r":
		return AlphaR(), nil
	case "alpha-r-fill":
		return AlphaRFill(), nil
	case "alpha-s":
		return AlphaS(), nil
	case "alpha-s-fill":
		return AlphaSFill(), nil
	case "alpha-t":
		return AlphaT(), nil
	case "alpha-t-fill":
		return AlphaTFill(), nil
	case "alpha-u":
		return AlphaU(), nil
	case "alpha-u-fill":
		return AlphaUFill(), nil
	case "alpha-v":
		return AlphaV(), nil
	case "alpha-v-fill":
		return AlphaVFill(), nil
	case "alpha-w":
		return AlphaW(), nil
	case "alpha-w-fill":
		return AlphaWFill(), nil
	case "alpha-x":
		return AlphaX(), nil
	case "alpha-x-fill":
		return AlphaXFill(), nil
	case "alpha-y":
		return AlphaY(), nil
	case "alpha-y-fill":
		return AlphaYFill(), nil
	case "alpha-z":
		return AlphaZ(), nil
	case "alpha-z-fill":
		return AlphaZFill(), nil
	case "application":
		return Application(), nil
	case "application-code":
		return ApplicationCode(), nil
	case "archive":
		return Archive(), nil
	case "arrow":
		return Arrow(), nil
	case "arrow-bottom-left":
		return ArrowBottomLeft(), nil
	case "arrow-bottom-left-circle":
		return ArrowBottomLeftCircle(), nil
	case "arrow-bottom-right":
		return ArrowBottomRight(), nil
	case "arrow-bottom-right-circle":
		return ArrowBottomRightCircle(), nil
	case "arrow-down":
		return ArrowDown(), nil
	case "arrow-down-bold":
		return ArrowDownBold(), nil
	case "arrow-down-circle":
		return ArrowDownCircle(), nil
	case "arrow-down-left":
		return ArrowDownLeft(), nil
	case "arrow-down-right":
		return ArrowDownRight(), nil
	case "arrow-left":
		return ArrowLeft(), nil
	case "arrow-left-bold":
		return ArrowLeftBold(), nil
	case "arrow-left-circle":
		return ArrowLeftCircle(), nil
	case "arrow-left-down":
		return ArrowLeftDown(), nil
	case "arrow-left-right":
		return ArrowLeftRight(), nil
	case "arrow-left-up":
		return ArrowLeftUp(), nil
	case "arrow-right":
		return ArrowRight(), nil
	case "arrow-right-bold":
		return ArrowRightBold(), nil
	case "arrow-right-circle":
		return ArrowRightCircle(), nil
	case "arrow-right-down":
		return ArrowRightDown(), nil
	case "arrow-right-up":
		return ArrowRightUp(), nil
	case "arrow-top-left":
		return ArrowTopLeft(), nil
	case "arrow-top-left-circle":
		return ArrowTopLeftCircle(), nil
	case "arrow-top-right":
		return ArrowTopRight(), nil
	case "arrow-top-right-circle":
		return ArrowTopRightCircle(), nil
	case "arrow-up":
		return ArrowUp(), nil
	case "arrow-up-bold":
		return ArrowUpBold(), nil
	case "arrow-up-circle":
		return ArrowUpCircle(), nil
	case "arrow-up-down":
		return ArrowUpDown(), nil
	case "arrow-up-left":
		return ArrowUpLeft(), nil
	case "arrow-up-right":
		return ArrowUpRight(), nil
	case "axe":
		return Axe(), nil
	case "bag-personal":
		return BagPersonal(), nil
	case "bag-personal-fill":
		return BagPersonalFill(), nil
	case "battery-50":
		return BatteryFifty(), nil
	case "battery-100":
		return BatteryOneHundred(), nil
	case "battery-75":
		return BatterySeventyFive(), nil
	case "battery-25":
		return BatteryTwentyFive(), nil
	case "battery-0":
		return BatteryZero(), nil
	case "battle-axe":
		return BattleAxe(), nil
	case "book":
		return Book(), nil
	case "bookmark":
		return Bookmark(), nil
	case "border-bottom":
		return BorderBottom(), nil
	case "border-bottom-left":
		return BorderBottomLeft(), nil
	case "border-bottom-left-right":
		return BorderBottomLeftRight(), nil
	case "border-bottom-right":
		return BorderBottomRight(), nil
	case "border-inside":
		return BorderInside(), nil
	case "border-left":
		return BorderLeft(), nil
	case "border-left-right":
		return BorderLeftRight(), nil
	case "border-none":
		return BorderNone(), nil
	case "border-outside":
		return BorderOutside(), nil
	case "border-right":
		return BorderRight(), nil
	case "border-top":
		return BorderTop(), nil
	case "border-top-bottom":
		return BorderTopBottom(), nil
	case "border-top-left":
		return BorderTopLeft(), nil
	case "border-top-left-bottom":
		return BorderTopLeftBottom(), nil
	case "border-top-left-right":
		return BorderTopLeftRight(), nil
	case "border-top-right":
		return BorderTopRight(), nil
	case "border-top-right-bottom":
		return BorderTopRightBottom(), nil
	case "bow":
		return Bow(), nil
	case "bow-arrow":
		return BowArrow(), nil
	case "box":
		return Box(), nil
	case "box-light-dashed-down-left":
		return BoxLightDashedDownLeft(), nil
	case "box-light-dashed-down-right":
		return BoxLightDashedDownRight(), nil
	case "box-light-dashed-up-left":
		return BoxLightDashedUpLeft(), nil
	case "box-light-dashed-up-right":
		return BoxLightDashedUpRight(), nil
	case "box-light-double-horizontal":
		return BoxLightDoubleHorizontal(), nil
	case "box-light-double-round-down-left":
		return BoxLightDoubleRoundDownLeft(), nil
	case "box-light-double-round-down-right":
		return BoxLightDoubleRoundDownRight(), nil
	case "box-light-double-round-up-left":
		return BoxLightDoubleRoundUpLeft(), nil
	case "box-light-double-round-up-right":
		return BoxLightDoubleRoundUpRight(), nil
	case "box-light-double-vertical":
		return BoxLightDoubleVertical(), nil
	case "box-light-down-left":
		return BoxLightDownLeft(), nil
	case "box-light-down-left-circle":
		return BoxLightDownLeftCircle(), nil
	case "box-light-down-right":
		return BoxLightDownRight(), nil
	case "box-light-down-right-circle":
		return BoxLightDownRightCircle(), nil
	case "box-light-horizontal":
		return BoxLightHorizontal(), nil
	case "box-light-round-down-left":
		return BoxLightRoundDownLeft(), nil
	case "box-light-round-down-right":
		return BoxLightRoundDownRight(), nil
	case "box-light-round-up-left":
		return BoxLightRoundUpLeft(), nil
	case "box-light-round-up-right":
		return BoxLightRoundUpRight(), nil
	case "box-light-up-left":
		return BoxLightUpLeft(), nil
	case "box-light-up-left-circle":
		return BoxLightUpLeftCircle(), nil
	case "box-light-up-right":
		return BoxLightUpRight(), nil
	case "box-light-up-right-circle":
		return BoxLightUpRightCircle(), nil
	case "box-light-vertical":
		return BoxLightVertical(), nil
	case "box-outer-light-all":
		return BoxOuterLightAll(), nil
	case "box-outer-light-dashed-all":
		return BoxOuterLightDashedAll(), nil
	case "box-outer-light-dashed-down":
		return BoxOuterLightDashedDown(), nil
	case "box-outer-light-dashed-down-left":
		return BoxOuterLightDashedDownLeft(), nil
	case "box-outer-light-dashed-down-left-right":
		return BoxOuterLightDashedDownLeftRight(), nil
	case "box-outer-light-dashed-down-right":
		return BoxOuterLightDashedDownRight(), nil
	case "box-outer-light-dashed-left":
		return BoxOuterLightDashedLeft(), nil
	case "box-outer-light-dashed-left-right":
		return BoxOuterLightDashedLeftRight(), nil
	case "box-outer-light-dashed-right":
		return BoxOuterLightDashedRight(), nil
	case "box-outer-light-dashed-up":
		return BoxOuterLightDashedUp(), nil
	case "box-outer-light-dashed-up-down":
		return BoxOuterLightDashedUpDown(), nil
	case "box-outer-light-dashed-up-down-left":
		return BoxOuterLightDashedUpDownLeft(), nil
	case "box-outer-light-dashed-up-down-right":
		return BoxOuterLightDashedUpDownRight(), nil
	case "box-outer-light-dashed-up-left":
		return BoxOuterLightDashedUpLeft(), nil
	case "box-outer-light-dashed-up-left-right":
		return BoxOuterLightDashedUpLeftRight(), nil
	case "box-outer-light-dashed-up-right":
		return BoxOuterLightDashedUpRight(), nil
	case "box-outer-light-down":
		return BoxOuterLightDown(), nil
	case "box-outer-light-down-left":
		return BoxOuterLightDownLeft(), nil
	case "box-outer-light-down-left-right":
		return BoxOuterLightDownLeftRight(), nil
	case "box-outer-light-down-right":
		return BoxOuterLightDownRight(), nil
	case "box-outer-light-left":
		return BoxOuterLightLeft(), nil
	case "box-outer-light-left-right":
		return BoxOuterLightLeftRight(), nil
	case "box-outer-light-right":
		return BoxOuterLightRight(), nil
	case "box-outer-light-round-down-left":
		return BoxOuterLightRoundDownLeft(), nil
	case "box-outer-light-round-down-right":
		return BoxOuterLightRoundDownRight(), nil
	case "box-outer-light-round-up-left":
		return BoxOuterLightRoundUpLeft(), nil
	case "box-outer-light-round-up-right":
		return BoxOuterLightRoundUpRight(), nil
	case "box-outer-light-up":
		return BoxOuterLightUp(), nil
	case "box-outer-light-up-down":
		return BoxOuterLightUpDown(), nil
	case "box-outer-light-up-down-left":
		return BoxOuterLightUpDownLeft(), nil
	case "box-outer-light-up-down-right":
		return BoxOuterLightUpDownRight(), nil
	case "box-outer-light-up-left":
		return BoxOuterLightUpLeft(), nil
	case "box-outer-light-up-left-right":
		return BoxOuterLightUpLeftRight(), nil
	case "box-outer-light-up-right":
		return BoxOuterLightUpRight(), nil
	case "briefcase":
		return Briefcase(), nil
	case "bug":
		return Bug(), nil
	case "bug-fill":
		return BugFill(), nil
	case "calculator":
		return Calculator(), nil
	case "calendar":
		return Calendar(), nil
	case "calendar-month":
		return CalendarMonth(), nil
	case "cancel":
		return Cancel(), nil
	case "card":
		return Card(), nil
	case "card-text":
		return CardText(), nil
	case "cart":
		return Cart(), nil
	case "chart-bar":
		return ChartBar(), nil
	case "chat":
		return Chat(), nil
	case "check":
		return Check(), nil
	case "checkbox-blank":
		return CheckboxBlank(), nil
	case "checkbox-cross":
		return CheckboxCross(), nil
	case "checkbox-marked":
		return CheckboxMarked(), nil
	case "chevron-down":
		return ChevronDown(), nil
	case "chevron-down-circle":
		return ChevronDownCircle(), nil
	case "chevron-left":
		return ChevronLeft(), nil
	case "chevron-left-circle":
		return ChevronLeftCircle(), nil
	case "chevron-right":
		return ChevronRight(), nil
	case "chevron-right-circle":
		return ChevronRightCircle(), nil
	case "chevron-up":
		return ChevronUp(), nil
	case "chevron-up-circle":
		return ChevronUpCircle(), nil
	case "circle":
		return Circle(), nil
	case "clipboard":
		return Clipboard(), nil
	case "clock":
		return Clock(), nil
	case "coffee":
		return Coffee(), nil
	case "comment":
		return Comment(), nil
	case "comment-text":
		return CommentText(), nil
	case "compass":
		return Compass(), nil
	case "compass-east-arrow":
		return CompassEastArrow(), nil
	case "compass-north-arrow":
		return CompassNorthArrow(), nil
	case "compass-north-east":
		return CompassNorthEast(), nil
	case "compass-north-west":
		return CompassNorthWest(), nil
	case "compass-south-arrow":
		return CompassSouthArrow(), nil
	case "compass-south-east":
		return CompassSouthEast(), nil
	case "compass-south-west":
		return CompassSouthWest(), nil
	case "compass-west-arrow":
		return CompassWestArrow(), nil
	case "credit-card":
		return CreditCard(), nil
	case "crown":
		return Crown(), nil
	case "cube-unfolded":
		return CubeUnfolded(), nil
	case "database":
		return Database(), nil
	case "device":
		return Device(), nil
	case "diamond":
		return Diamond(), nil
	case "division":
		return Division(), nil
	case "door":
		return Door(), nil
	case "door-box":
		return DoorBox(), nil
	case "door-open":
		return DoorOpen(), nil
	case "download":
		return Download(), nil
	case "email":
		return Email(), nil
	case "file":
		return File(), nil
	case "fire":
		return Fire(), nil
	case "flask":
		return Flask(), nil
	case "flask-empty":
		return FlaskEmpty(), nil
	case "flask-round-bottom":
		return FlaskRoundBottom(), nil
	case "flask-round-bottom-empty":
		return FlaskRoundBottomEmpty(), nil
	case "floppy-disk":
		return FloppyDisk(), nil
	case "folder":
		return Folder(), nil
	case "folder-open":
		return FolderOpen(), nil
	case "gamepad-center":
		return GamepadCenter(), nil
	case "gamepad-down":
		return GamepadDown(), nil
	case "gamepad-down-left":
		return GamepadDownLeft(), nil
	case "gamepad-down-right":
		return GamepadDownRight(), nil
	case "gamepad-empty":
		return GamepadEmpty(), nil
	case "gamepad-left":
		return GamepadLeft(), nil
	case "gamepad-right":
		return GamepadRight(), nil
	case "gamepad-up":
		return GamepadUp(), nil
	case "gamepad-up-left":
		return GamepadUpLeft(), nil
	case "gamepad-up-right":
		return GamepadUpRight(), nil
	case "heart":
		return Heart(), nil
	case "image":
		return Image(), nil
	case "label":
		return Label(), nil
	case "label-variant":
		return LabelVariant(), nil
	case "led":
		return Led(), nil
	case "lightbulb":
		return Lightbulb(), nil
	case "lock":
		return Lock(), nil
	case "lock-open":
		return LockOpen(), nil
	case "login":
		return Login(), nil
	case "logout":
		return Logout(), nil
	case "map":
		return Map(), nil
	case "menu-bottom-left":
		return MenuBottomLeft(), nil
	case "menu-bottom-right":
		return MenuBottomRight(), nil
	case "menu-down":
		return MenuDown(), nil
	case "menu-down-fill":
		return MenuDownFill(), nil
	case "menu-left":
		return MenuLeft(), nil
	case "menu-left-fill":
		return MenuLeftFill(), nil
	case "menu-left-right":
		return MenuLeftRight(), nil
	case "menu-right":
		return MenuRight(), nil
	case "menu-right-fill":
		return MenuRightFill(), nil
	case "menu-top-left":
		return MenuTopLeft(), nil
	case "menu-top-right":
		return MenuTopRight(), nil
	case "menu-up":
		return MenuUp(), nil
	case "menu-up-down":
		return MenuUpDown(), nil
	case "menu-up-fill":
		return MenuUpFill(), nil
	case "message":
		return Message(), nil
	case "message-processing":
		return MessageProcessing(), nil
	case "message-text":
		return MessageText(), nil
	case "microphone":
		return Microphone(), nil
	case "minus":
		return Minus(), nil
	case "minus-box":
		return MinusBox(), nil
	case "minus-box-fill":
		return MinusBoxFill(), nil
	case "minus-circle":
		return MinusCircle(), nil
	case "minus-circle-fill":
		return MinusCircleFill(), nil
	case "monitor":
		return Monitor(), nil
	case "monitor-image":
		return MonitorImage(), nil
	case "multiply":
		return Multiply(), nil
	case "music-note":
		return MusicNote(), nil
	case "necklace":
		return Necklace(), nil
	case "note":
		return Note(), nil
	case "notebook":
		return Notebook(), nil
	case "notification":
		return Notification(), nil
	case "octagon":
		return Octagon(), nil
	case "octagon-alert":
		return OctagonAlert(), nil
	case "pause":
		return Pause(), nil
	case "pencil":
		return Pencil(), nil
	case "pickaxe":
		return Pickaxe(), nil
	case "pictogrammers":
		return Pictogrammers(), nil
	case "play":
		return Play(), nil
	case "plus":
		return Plus(), nil
	case "plus-box":
		return PlusBox(), nil
	case "plus-box-fill":
		return PlusBoxFill(), nil
	case "plus-circle":
		return PlusCircle(), nil
	case "plus-circle-fill":
		return PlusCircleFill(), nil
	case "radiobox":
		return Radiobox(), nil
	case "radiobox-marked":
		return RadioboxMarked(), nil
	case "remove-circle":
		return RemoveCircle(), nil
	case "rotate-clockwise":
		return RotateClockwise(), nil
	case "rotate-counterclockwise":
		return RotateCounterclockwise(), nil
	case "script":
		return Script(), nil
	case "shield":
		return Shield(), nil
	case "skull":
		return Skull(), nil
	case "speaker":
		return Speaker(), nil
	case "stop":
		return Stop(), nil
	case "sword":
		return Sword(), nil
	case "tag":
		return Tag(), nil
	case "tag-text":
		return TagText(), nil
	case "target":
		return Target(), nil
	case "terminal":
		return Terminal(), nil
	case "text-box":
		return TextBox(), nil
	case "text-image":
		return TextImage(), nil
	case "toggle-switch-off":
		return ToggleSwitchOff(), nil
	case "toggle-switch-on":
		return ToggleSwitchOn(), nil
	case "toolbox":
		return Toolbox(), nil
	case "tooltip-above":
		return TooltipAbove(), nil
	case "tooltip-above-alert":
		return TooltipAboveAlert(), nil
	case "tooltip-above-text":
		return TooltipAboveText(), nil
	case "tooltip-below":
		return TooltipBelow(), nil
	case "tooltip-below-alert":
		return TooltipBelowAlert(), nil
	case "tooltip-below-text":
		return TooltipBelowText(), nil
	case "tooltip-end":
		return TooltipEnd(), nil
	case "tooltip-end-alert":
		return TooltipEndAlert(), nil
	case "tooltip-end-text":
		return TooltipEndText(), nil
	case "tooltip-start":
		return TooltipStart(), nil
	case "tooltip-start-alert":
		return TooltipStartAlert(), nil
	case "tooltip-start-text":
		return TooltipStartText(), nil
	case "trash":
		return Trash(), nil
	case "upload":
		return Upload(), nil
	case "volume-high":
		return VolumeHigh(), nil
	case "volume-low":
		return VolumeLow(), nil
	case "volume-medium":
		return VolumeMedium(), nil
	case "volume-mute":
		return VolumeMute(), nil
	case "wall":
		return Wall(), nil
	case "wall-fill":
		return WallFill(), nil
	case "water":
		return Water(), nil
	case "well":
		return Well(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in memory icon set", name)
	}
}
