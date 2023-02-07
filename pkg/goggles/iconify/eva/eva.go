package eva

import (
	"fmt"
	"github.com/gogoracer/racer/pkg/engine"
)

const (
	activityFillInnerSVG                  = `<g id="evaActivityFill0"><g id="evaActivityFill1"><path id="evaActivityFill2" fill="currentColor" d="M14.33 20h-.21a2 2 0 0 1-1.76-1.58L9.68 6l-2.76 6.4A1 1 0 0 1 6 13H3a1 1 0 0 1 0-2h2.34l2.51-5.79a2 2 0 0 1 3.79.38L14.32 18l2.76-6.38A1 1 0 0 1 18 11h3a1 1 0 0 1 0 2h-2.34l-2.51 5.79A2 2 0 0 1 14.33 20Z"/></g></g>`
	activityOutlineInnerSVG               = `<g id="evaActivityOutline0"><g id="evaActivityOutline1"><path id="evaActivityOutline2" fill="currentColor" d="M14.33 20h-.21a2 2 0 0 1-1.76-1.58L9.68 6l-2.76 6.4A1 1 0 0 1 6 13H3a1 1 0 0 1 0-2h2.34l2.51-5.79a2 2 0 0 1 3.79.38L14.32 18l2.76-6.38A1 1 0 0 1 18 11h3a1 1 0 0 1 0 2h-2.34l-2.51 5.79A2 2 0 0 1 14.33 20Z"/></g></g>`
	alertCircleFillInnerSVG               = `<g id="evaAlertCircleFill0"><g id="evaAlertCircleFill1"><path id="evaAlertCircleFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 15a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm1-4a1 1 0 0 1-2 0V8a1 1 0 0 1 2 0Z"/></g></g>`
	alertCircleOutlineInnerSVG            = `<g id="evaAlertCircleOutline0"><g id="evaAlertCircleOutline1"><g id="evaAlertCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><circle cx="12" cy="16" r="1"/><path d="M12 7a1 1 0 0 0-1 1v5a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1Z"/></g></g></g>`
	alertTriangleFillInnerSVG             = `<g id="evaAlertTriangleFill0"><g id="evaAlertTriangleFill1"><path id="evaAlertTriangleFill2" fill="currentColor" d="M22.56 16.3L14.89 3.58a3.43 3.43 0 0 0-5.78 0L1.44 16.3a3 3 0 0 0-.05 3A3.37 3.37 0 0 0 4.33 21h15.34a3.37 3.37 0 0 0 2.94-1.66a3 3 0 0 0-.05-3.04ZM12 17a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm1-4a1 1 0 0 1-2 0V9a1 1 0 0 1 2 0Z"/></g></g>`
	alertTriangleOutlineInnerSVG          = `<g id="evaAlertTriangleOutline0"><g id="evaAlertTriangleOutline1"><g id="evaAlertTriangleOutline2" fill="currentColor"><path d="M22.56 16.3L14.89 3.58a3.43 3.43 0 0 0-5.78 0L1.44 16.3a3 3 0 0 0-.05 3A3.37 3.37 0 0 0 4.33 21h15.34a3.37 3.37 0 0 0 2.94-1.66a3 3 0 0 0-.05-3.04Zm-1.7 2.05a1.31 1.31 0 0 1-1.19.65H4.33a1.31 1.31 0 0 1-1.19-.65a1 1 0 0 1 0-1l7.68-12.73a1.48 1.48 0 0 1 2.36 0l7.67 12.72a1 1 0 0 1 .01 1.01Z"/><circle cx="12" cy="16" r="1"/><path d="M12 8a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1Z"/></g></g></g>`
	archiveFillInnerSVG                   = `<g id="evaArchiveFill0"><g id="evaArchiveFill1"><path id="evaArchiveFill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-2 5.22V18a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8.22A3 3 0 0 0 18 3Zm-3 10.13a.87.87 0 0 1-.87.87H9.87a.87.87 0 0 1-.87-.87v-.26a.87.87 0 0 1 .87-.87h4.26a.87.87 0 0 1 .87.87ZM18 7H6a1 1 0 0 1 0-2h12a1 1 0 0 1 0 2Z"/></g></g>`
	archiveOutlineInnerSVG                = `<g id="evaArchiveOutline0"><g id="evaArchiveOutline1"><g id="evaArchiveOutline2" fill="currentColor"><path d="M21 6a3 3 0 0 0-3-3H6a3 3 0 0 0-2 5.22V18a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8.22A3 3 0 0 0 21 6ZM6 5h12a1 1 0 0 1 0 2H6a1 1 0 0 1 0-2Zm12 13a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V9h12Z"/><rect width="6" height="2" x="9" y="12" rx=".87" ry=".87"/></g></g></g>`
	arrowBackFillInnerSVG                 = `<g id="evaArrowBackFill0"><g id="evaArrowBackFill1"><path id="evaArrowBackFill2" fill="currentColor" d="M19 11H7.14l3.63-4.36a1 1 0 1 0-1.54-1.28l-5 6a1.19 1.19 0 0 0-.09.15c0 .05 0 .08-.07.13A1 1 0 0 0 4 12a1 1 0 0 0 .07.36c0 .05 0 .08.07.13a1.19 1.19 0 0 0 .09.15l5 6A1 1 0 0 0 10 19a1 1 0 0 0 .64-.23a1 1 0 0 0 .13-1.41L7.14 13H19a1 1 0 0 0 0-2Z"/></g></g>`
	arrowBackOutlineInnerSVG              = `<g id="evaArrowBackOutline0"><g id="evaArrowBackOutline1"><path id="evaArrowBackOutline2" fill="currentColor" d="M19 11H7.14l3.63-4.36a1 1 0 1 0-1.54-1.28l-5 6a1.19 1.19 0 0 0-.09.15c0 .05 0 .08-.07.13A1 1 0 0 0 4 12a1 1 0 0 0 .07.36c0 .05 0 .08.07.13a1.19 1.19 0 0 0 .09.15l5 6A1 1 0 0 0 10 19a1 1 0 0 0 .64-.23a1 1 0 0 0 .13-1.41L7.14 13H19a1 1 0 0 0 0-2Z"/></g></g>`
	arrowCircleDownFillInnerSVG           = `<g id="evaArrowCircleDownFill0"><g id="evaArrowCircleDownFill1"><path id="evaArrowCircleDownFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm3.69 11.86l-3 2.86a.49.49 0 0 1-.15.1a.54.54 0 0 1-.16.1a.94.94 0 0 1-.76 0a1 1 0 0 1-.33-.21l-3-3a1 1 0 0 1 1.42-1.42l1.29 1.3V8a1 1 0 0 1 2 0v5.66l1.31-1.25a1 1 0 0 1 1.38 1.45Z"/></g></g>`
	arrowCircleDownOutlineInnerSVG        = `<g id="evaArrowCircleDownOutline0"><g id="evaArrowCircleDownOutline1"><g id="evaArrowCircleDownOutline2" fill="currentColor"><path d="M14.31 12.41L13 13.66V8a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a.54.54 0 0 0 .16-.1a.49.49 0 0 0 .15-.1l3-2.86a1 1 0 0 0-1.38-1.45Z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/></g></g></g>`
	arrowCircleLeftFillInnerSVG           = `<g id="evaArrowCircleLeftFill0"><g id="evaArrowCircleLeftFill1"><path id="evaArrowCircleLeftFill2" fill="currentColor" d="M22 12a10 10 0 1 0-10 10a10 10 0 0 0 10-10Zm-11.86 3.69l-2.86-3a.49.49 0 0 1-.1-.15a.54.54 0 0 1-.1-.16a.94.94 0 0 1 0-.76a1 1 0 0 1 .21-.33l3-3a1 1 0 0 1 1.42 1.42L10.41 11H16a1 1 0 0 1 0 2h-5.66l1.25 1.31a1 1 0 0 1-1.45 1.38Z"/></g></g>`
	arrowCircleLeftOutlineInnerSVG        = `<g id="evaArrowCircleLeftOutline0"><g id="evaArrowCircleLeftOutline1"><g id="evaArrowCircleLeftOutline2" fill="currentColor"><path d="M16 11h-5.66l1.25-1.31a1 1 0 0 0-1.45-1.38l-2.86 3a1 1 0 0 0-.09.13a.72.72 0 0 0-.11.19a.88.88 0 0 0-.06.28S7 12 7 12a1 1 0 0 0 .08.38a1 1 0 0 0 .21.32l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L10.41 13H16a1 1 0 0 0 0-2Z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/></g></g></g>`
	arrowCircleRightFillInnerSVG          = `<g id="evaArrowCircleRightFill0"><g id="evaArrowCircleRightFill1"><path id="evaArrowCircleRightFill2" fill="currentColor" d="M2 12A10 10 0 1 0 12 2A10 10 0 0 0 2 12Zm11.86-3.69l2.86 3a.49.49 0 0 1 .1.15a.54.54 0 0 1 .1.16a.94.94 0 0 1 0 .76a1 1 0 0 1-.21.33l-3 3a1 1 0 0 1-1.42-1.42l1.3-1.29H8a1 1 0 0 1 0-2h5.66l-1.25-1.31a1 1 0 0 1 1.45-1.38Z"/></g></g>`
	arrowCircleRightOutlineInnerSVG       = `<g id="evaArrowCircleRightOutline0"><g id="evaArrowCircleRightOutline1"><g id="evaArrowCircleRightOutline2" fill="currentColor"><path d="M17 12v-.09a.88.88 0 0 0-.06-.28a.72.72 0 0 0-.11-.19a1 1 0 0 0-.09-.13l-2.86-3a1 1 0 0 0-1.45 1.38L13.66 11H8a1 1 0 0 0 0 2h5.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 .21-.32A1 1 0 0 0 17 12Z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/></g></g></g>`
	arrowCircleUpFillInnerSVG             = `<g id="evaArrowCircleUpFill0"><g id="evaArrowCircleUpFill1"><path id="evaArrowCircleUpFill2" fill="currentColor" d="M12 22A10 10 0 1 0 2 12a10 10 0 0 0 10 10ZM8.31 10.14l3-2.86a.49.49 0 0 1 .15-.1a.54.54 0 0 1 .16-.1a.94.94 0 0 1 .76 0a1 1 0 0 1 .33.21l3 3a1 1 0 0 1-1.42 1.42L13 10.41V16a1 1 0 0 1-2 0v-5.66l-1.31 1.25a1 1 0 0 1-1.38-1.45Z"/></g></g>`
	arrowCircleUpOutlineInnerSVG          = `<g id="evaArrowCircleUpOutline0"><g id="evaArrowCircleUpOutline1"><g id="evaArrowCircleUpOutline2" fill="currentColor"><path d="M12.71 7.29a1 1 0 0 0-.32-.21A1 1 0 0 0 12 7h-.1a.82.82 0 0 0-.27.06a.72.72 0 0 0-.19.11a1 1 0 0 0-.13.09l-3 2.86a1 1 0 0 0 1.38 1.45L11 10.34V16a1 1 0 0 0 2 0v-5.59l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/></g></g></g>`
	arrowDownFillInnerSVG                 = `<g id="evaArrowDownFill0"><g id="evaArrowDownFill1"><path id="evaArrowDownFill2" fill="currentColor" d="M12 17a1.72 1.72 0 0 1-1.33-.64l-4.21-5.1a2.1 2.1 0 0 1-.26-2.21A1.76 1.76 0 0 1 7.79 8h8.42a1.76 1.76 0 0 1 1.59 1.05a2.1 2.1 0 0 1-.26 2.21l-4.21 5.1A1.72 1.72 0 0 1 12 17Z"/></g></g>`
	arrowDownOutlineInnerSVG              = `<g id="evaArrowDownOutline0"><g id="evaArrowDownOutline1"><path id="evaArrowDownOutline2" fill="currentColor" d="M12 17a1.72 1.72 0 0 1-1.33-.64l-4.21-5.1a2.1 2.1 0 0 1-.26-2.21A1.76 1.76 0 0 1 7.79 8h8.42a1.76 1.76 0 0 1 1.59 1.05a2.1 2.1 0 0 1-.26 2.21l-4.21 5.1A1.72 1.72 0 0 1 12 17Zm-3.91-7L12 14.82L16 10Z"/></g></g>`
	arrowDownwardFillInnerSVG             = `<g id="evaArrowDownwardFill0"><g id="evaArrowDownwardFill1"><path id="evaArrowDownwardFill2" fill="currentColor" d="M18.77 13.36a1 1 0 0 0-1.41-.13L13 16.86V5a1 1 0 0 0-2 0v11.86l-4.36-3.63a1 1 0 1 0-1.28 1.54l6 5l.15.09l.13.07a1 1 0 0 0 .72 0l.13-.07l.15-.09l6-5a1 1 0 0 0 .13-1.41Z"/></g></g>`
	arrowDownwardOutlineInnerSVG          = `<g id="evaArrowDownwardOutline0"><g id="evaArrowDownwardOutline1"><path id="evaArrowDownwardOutline2" fill="currentColor" d="M18.77 13.36a1 1 0 0 0-1.41-.13L13 16.86V5a1 1 0 0 0-2 0v11.86l-4.36-3.63a1 1 0 1 0-1.28 1.54l6 5l.15.09l.13.07a1 1 0 0 0 .72 0l.13-.07l.15-.09l6-5a1 1 0 0 0 .13-1.41Z"/></g></g>`
	arrowForwardFillInnerSVG              = `<g id="evaArrowForwardFill0"><g id="evaArrowForwardFill1"><path id="evaArrowForwardFill2" fill="currentColor" d="M5 13h11.86l-3.63 4.36a1 1 0 0 0 1.54 1.28l5-6a1.19 1.19 0 0 0 .09-.15c0-.05.05-.08.07-.13A1 1 0 0 0 20 12a1 1 0 0 0-.07-.36c0-.05-.05-.08-.07-.13a1.19 1.19 0 0 0-.09-.15l-5-6A1 1 0 0 0 14 5a1 1 0 0 0-.64.23a1 1 0 0 0-.13 1.41L16.86 11H5a1 1 0 0 0 0 2Z"/></g></g>`
	arrowForwardOutlineInnerSVG           = `<g id="evaArrowForwardOutline0"><g id="evaArrowForwardOutline1"><path id="evaArrowForwardOutline2" fill="currentColor" d="M5 13h11.86l-3.63 4.36a1 1 0 0 0 1.54 1.28l5-6a1.19 1.19 0 0 0 .09-.15c0-.05.05-.08.07-.13A1 1 0 0 0 20 12a1 1 0 0 0-.07-.36c0-.05-.05-.08-.07-.13a1.19 1.19 0 0 0-.09-.15l-5-6A1 1 0 0 0 14 5a1 1 0 0 0-.64.23a1 1 0 0 0-.13 1.41L16.86 11H5a1 1 0 0 0 0 2Z"/></g></g>`
	arrowIosBackFillInnerSVG              = `<g id="evaArrowIosBackFill0"><g id="evaArrowIosBackFill1"><path id="evaArrowIosBackFill2" fill="currentColor" d="M13.83 19a1 1 0 0 1-.78-.37l-4.83-6a1 1 0 0 1 0-1.27l5-6a1 1 0 0 1 1.54 1.28L10.29 12l4.32 5.36a1 1 0 0 1-.78 1.64Z"/></g></g>`
	arrowIosBackOutlineInnerSVG           = `<g id="evaArrowIosBackOutline0"><g id="evaArrowIosBackOutline1"><path id="evaArrowIosBackOutline2" fill="currentColor" d="M13.83 19a1 1 0 0 1-.78-.37l-4.83-6a1 1 0 0 1 0-1.27l5-6a1 1 0 0 1 1.54 1.28L10.29 12l4.32 5.36a1 1 0 0 1-.78 1.64Z"/></g></g>`
	arrowIosDownwardFillInnerSVG          = `<g id="evaArrowIosDownwardFill0"><g id="evaArrowIosDownwardFill1"><path id="evaArrowIosDownwardFill2" fill="currentColor" d="M12 16a1 1 0 0 1-.64-.23l-6-5a1 1 0 1 1 1.28-1.54L12 13.71l5.36-4.32a1 1 0 0 1 1.41.15a1 1 0 0 1-.14 1.46l-6 4.83A1 1 0 0 1 12 16Z"/></g></g>`
	arrowIosDownwardOutlineInnerSVG       = `<g id="evaArrowIosDownwardOutline0"><g id="evaArrowIosDownwardOutline1"><path id="evaArrowIosDownwardOutline2" fill="currentColor" d="M12 16a1 1 0 0 1-.64-.23l-6-5a1 1 0 1 1 1.28-1.54L12 13.71l5.36-4.32a1 1 0 0 1 1.41.15a1 1 0 0 1-.14 1.46l-6 4.83A1 1 0 0 1 12 16Z"/></g></g>`
	arrowIosForwardFillInnerSVG           = `<g id="evaArrowIosForwardFill0"><g id="evaArrowIosForwardFill1"><path id="evaArrowIosForwardFill2" fill="currentColor" d="M10 19a1 1 0 0 1-.64-.23a1 1 0 0 1-.13-1.41L13.71 12L9.39 6.63a1 1 0 0 1 .15-1.41a1 1 0 0 1 1.46.15l4.83 6a1 1 0 0 1 0 1.27l-5 6A1 1 0 0 1 10 19Z"/></g></g>`
	arrowIosForwardOutlineInnerSVG        = `<g id="evaArrowIosForwardOutline0"><g id="evaArrowIosForwardOutline1"><path id="evaArrowIosForwardOutline2" fill="currentColor" d="M10 19a1 1 0 0 1-.64-.23a1 1 0 0 1-.13-1.41L13.71 12L9.39 6.63a1 1 0 0 1 .15-1.41a1 1 0 0 1 1.46.15l4.83 6a1 1 0 0 1 0 1.27l-5 6A1 1 0 0 1 10 19Z"/></g></g>`
	arrowIosUpwardFillInnerSVG            = `<g id="evaArrowIosUpwardFill0"><g id="evaArrowIosUpwardFill1"><path id="evaArrowIosUpwardFill2" fill="currentColor" d="M18 15a1 1 0 0 1-.64-.23L12 10.29l-5.37 4.32a1 1 0 0 1-1.41-.15a1 1 0 0 1 .15-1.41l6-4.83a1 1 0 0 1 1.27 0l6 5a1 1 0 0 1 .13 1.41A1 1 0 0 1 18 15Z"/></g></g>`
	arrowIosUpwardOutlineInnerSVG         = `<g id="evaArrowIosUpwardOutline0"><g id="evaArrowIosUpwardOutline1"><path id="evaArrowIosUpwardOutline2" fill="currentColor" d="M18 15a1 1 0 0 1-.64-.23L12 10.29l-5.37 4.32a1 1 0 0 1-1.41-.15a1 1 0 0 1 .15-1.41l6-4.83a1 1 0 0 1 1.27 0l6 5a1 1 0 0 1 .13 1.41A1 1 0 0 1 18 15Z"/></g></g>`
	arrowLeftFillInnerSVG                 = `<g id="evaArrowLeftFill0"><g id="evaArrowLeftFill1"><path id="evaArrowLeftFill2" fill="currentColor" d="M13.54 18a2.06 2.06 0 0 1-1.3-.46l-5.1-4.21a1.7 1.7 0 0 1 0-2.66l5.1-4.21a2.1 2.1 0 0 1 2.21-.26a1.76 1.76 0 0 1 1.05 1.59v8.42a1.76 1.76 0 0 1-1.05 1.59a2.23 2.23 0 0 1-.91.2Z"/></g></g>`
	arrowLeftOutlineInnerSVG              = `<g id="evaArrowLeftOutline0"><g id="evaArrowLeftOutline1"><path id="evaArrowLeftOutline2" fill="currentColor" d="M13.54 18a2.06 2.06 0 0 1-1.3-.46l-5.1-4.21a1.7 1.7 0 0 1 0-2.66l5.1-4.21a2.1 2.1 0 0 1 2.21-.26a1.76 1.76 0 0 1 1.05 1.59v8.42a1.76 1.76 0 0 1-1.05 1.59a2.23 2.23 0 0 1-.91.2Zm-4.86-6l4.82 4V8.09Z"/></g></g>`
	arrowRightFillInnerSVG                = `<g id="evaArrowRightFill0"><g id="evaArrowRightFill1"><path id="evaArrowRightFill2" fill="currentColor" d="M10.46 18a2.23 2.23 0 0 1-.91-.2a1.76 1.76 0 0 1-1.05-1.59V7.79A1.76 1.76 0 0 1 9.55 6.2a2.1 2.1 0 0 1 2.21.26l5.1 4.21a1.7 1.7 0 0 1 0 2.66l-5.1 4.21a2.06 2.06 0 0 1-1.3.46Z"/></g></g>`
	arrowRightOutlineInnerSVG             = `<g id="evaArrowRightOutline0"><g id="evaArrowRightOutline1"><path id="evaArrowRightOutline2" fill="currentColor" d="M10.46 18a2.23 2.23 0 0 1-.91-.2a1.76 1.76 0 0 1-1.05-1.59V7.79A1.76 1.76 0 0 1 9.55 6.2a2.1 2.1 0 0 1 2.21.26l5.1 4.21a1.7 1.7 0 0 1 0 2.66l-5.1 4.21a2.06 2.06 0 0 1-1.3.46Zm0-10v7.9l4.86-3.9Z"/></g></g>`
	arrowUpFillInnerSVG                   = `<g id="evaArrowUpFill0"><g id="evaArrowUpFill1"><path id="evaArrowUpFill2" fill="currentColor" d="M16.21 16H7.79a1.76 1.76 0 0 1-1.59-1a2.1 2.1 0 0 1 .26-2.21l4.21-5.1a1.76 1.76 0 0 1 2.66 0l4.21 5.1A2.1 2.1 0 0 1 17.8 15a1.76 1.76 0 0 1-1.59 1Z"/></g></g>`
	arrowUpOutlineInnerSVG                = `<g id="evaArrowUpOutline0"><g id="evaArrowUpOutline1"><path id="evaArrowUpOutline2" fill="currentColor" d="M16.21 16H7.79a1.76 1.76 0 0 1-1.59-1a2.1 2.1 0 0 1 .26-2.21l4.21-5.1a1.76 1.76 0 0 1 2.66 0l4.21 5.1A2.1 2.1 0 0 1 17.8 15a1.76 1.76 0 0 1-1.59 1ZM8 14h7.9L12 9.18Z"/></g></g>`
	arrowUpwardFillInnerSVG               = `<g id="evaArrowUpwardFill0"><g id="evaArrowUpwardFill1"><path id="evaArrowUpwardFill2" fill="currentColor" d="M5.23 10.64a1 1 0 0 0 1.41.13L11 7.14V19a1 1 0 0 0 2 0V7.14l4.36 3.63a1 1 0 1 0 1.28-1.54l-6-5l-.15-.09l-.13-.07a1 1 0 0 0-.72 0l-.13.07l-.15.09l-6 5a1 1 0 0 0-.13 1.41Z"/></g></g>`
	arrowUpwardOutlineInnerSVG            = `<g id="evaArrowUpwardOutline0"><g id="evaArrowUpwardOutline1"><path id="evaArrowUpwardOutline2" fill="currentColor" d="M5.23 10.64a1 1 0 0 0 1.41.13L11 7.14V19a1 1 0 0 0 2 0V7.14l4.36 3.63a1 1 0 1 0 1.28-1.54l-6-5l-.15-.09l-.13-.07a1 1 0 0 0-.72 0l-.13.07l-.15.09l-6 5a1 1 0 0 0-.13 1.41Z"/></g></g>`
	arrowheadDownFillInnerSVG             = `<g id="evaArrowheadDownFill0"><g id="evaArrowheadDownFill1"><g id="evaArrowheadDownFill2" fill="currentColor"><path d="M17.37 12.39L12 16.71l-5.36-4.48a1 1 0 1 0-1.28 1.54l6 5a1 1 0 0 0 1.27 0l6-4.83a1 1 0 0 0 .15-1.41a1 1 0 0 0-1.41-.14Z"/><path d="M11.36 11.77a1 1 0 0 0 1.27 0l6-4.83a1 1 0 0 0 .15-1.41a1 1 0 0 0-1.41-.15L12 9.71L6.64 5.23a1 1 0 0 0-1.28 1.54Z"/></g></g></g>`
	arrowheadDownOutlineInnerSVG          = `<g id="evaArrowheadDownOutline0"><g id="evaArrowheadDownOutline1"><g id="evaArrowheadDownOutline2" fill="currentColor"><path d="M17.37 12.39L12 16.71l-5.36-4.48a1 1 0 1 0-1.28 1.54l6 5a1 1 0 0 0 1.27 0l6-4.83a1 1 0 0 0 .15-1.41a1 1 0 0 0-1.41-.14Z"/><path d="M11.36 11.77a1 1 0 0 0 1.27 0l6-4.83a1 1 0 0 0 .15-1.41a1 1 0 0 0-1.41-.15L12 9.71L6.64 5.23a1 1 0 0 0-1.28 1.54Z"/></g></g></g>`
	arrowheadLeftFillInnerSVG             = `<g id="evaArrowheadLeftFill0"><g id="evaArrowheadLeftFill1"><g id="evaArrowheadLeftFill2" fill="currentColor"><path d="M11.64 5.23a1 1 0 0 0-1.41.13l-5 6a1 1 0 0 0 0 1.27l4.83 6a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.63L7.29 12l4.48-5.37a1 1 0 0 0-.13-1.4Z"/><path d="m14.29 12l4.48-5.37a1 1 0 0 0-1.54-1.28l-5 6a1 1 0 0 0 0 1.27l4.83 6a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.63Z"/></g></g></g>`
	arrowheadLeftOutlineInnerSVG          = `<g id="evaArrowheadLeftOutline0"><g id="evaArrowheadLeftOutline1"><g id="evaArrowheadLeftOutline2" fill="currentColor"><path d="M11.64 5.23a1 1 0 0 0-1.41.13l-5 6a1 1 0 0 0 0 1.27l4.83 6a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.63L7.29 12l4.48-5.37a1 1 0 0 0-.13-1.4Z"/><path d="m14.29 12l4.48-5.37a1 1 0 0 0-1.54-1.28l-5 6a1 1 0 0 0 0 1.27l4.83 6a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.63Z"/></g></g></g>`
	arrowheadRightFillInnerSVG            = `<g id="evaArrowheadRightFill0"><g id="evaArrowheadRightFill1"><g id="evaArrowheadRightFill2" fill="currentColor"><path d="m18.78 11.37l-4.78-6a1 1 0 0 0-1.41-.15a1 1 0 0 0-.15 1.41L16.71 12l-4.48 5.37a1 1 0 0 0 .13 1.41A1 1 0 0 0 13 19a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 .01-1.27Z"/><path d="M7 5.37a1 1 0 0 0-1.61 1.26L9.71 12l-4.48 5.36a1 1 0 0 0 .13 1.41A1 1 0 0 0 6 19a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 0-1.27Z"/></g></g></g>`
	arrowheadRightOutlineInnerSVG         = `<g id="evaArrowheadRightOutline0"><g id="evaArrowheadRightOutline1"><g id="evaArrowheadRightOutline2" fill="currentColor"><path d="m18.78 11.37l-4.78-6a1 1 0 0 0-1.41-.15a1 1 0 0 0-.15 1.41L16.71 12l-4.48 5.37a1 1 0 0 0 .13 1.41A1 1 0 0 0 13 19a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 .01-1.27Z"/><path d="M7 5.37a1 1 0 0 0-1.61 1.26L9.71 12l-4.48 5.36a1 1 0 0 0 .13 1.41A1 1 0 0 0 6 19a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 0-1.27Z"/></g></g></g>`
	arrowheadUpFillInnerSVG               = `<g id="evaArrowheadUpFill0"><g id="evaArrowheadUpFill1"><g id="evaArrowheadUpFill2" fill="currentColor"><path d="M6.63 11.61L12 7.29l5.37 4.48A1 1 0 0 0 18 12a1 1 0 0 0 .77-.36a1 1 0 0 0-.13-1.41l-6-5a1 1 0 0 0-1.27 0l-6 4.83a1 1 0 0 0-.15 1.41a1 1 0 0 0 1.41.14Z"/><path d="M12.64 12.23a1 1 0 0 0-1.27 0l-6 4.83a1 1 0 0 0-.15 1.41a1 1 0 0 0 1.41.15L12 14.29l5.37 4.48A1 1 0 0 0 18 19a1 1 0 0 0 .77-.36a1 1 0 0 0-.13-1.41Z"/></g></g></g>`
	arrowheadUpOutlineInnerSVG            = `<g id="evaArrowheadUpOutline0"><g id="evaArrowheadUpOutline1"><g id="evaArrowheadUpOutline2" fill="currentColor"><path d="M6.63 11.61L12 7.29l5.37 4.48A1 1 0 0 0 18 12a1 1 0 0 0 .77-.36a1 1 0 0 0-.13-1.41l-6-5a1 1 0 0 0-1.27 0l-6 4.83a1 1 0 0 0-.15 1.41a1 1 0 0 0 1.41.14Z"/><path d="M12.64 12.23a1 1 0 0 0-1.27 0l-6 4.83a1 1 0 0 0-.15 1.41a1 1 0 0 0 1.41.15L12 14.29l5.37 4.48A1 1 0 0 0 18 19a1 1 0 0 0 .77-.36a1 1 0 0 0-.13-1.41Z"/></g></g></g>`
	atFillInnerSVG                        = `<g id="evaAtFill0"><g id="evaAtFill1"><path id="evaAtFill2" fill="currentColor" d="M13 2a10 10 0 0 0-5 19.1a10.15 10.15 0 0 0 4 .9a10 10 0 0 0 6.08-2.06a1 1 0 0 0 .19-1.4a1 1 0 0 0-1.41-.19A8 8 0 1 1 12.77 4A8.17 8.17 0 0 1 20 12.22v.68a1.71 1.71 0 0 1-1.78 1.7a1.82 1.82 0 0 1-1.62-1.88V8.4a1 1 0 0 0-1-1a1 1 0 0 0-1 .87a5 5 0 0 0-3.44-1.36A5.09 5.09 0 1 0 15.31 15a3.6 3.6 0 0 0 5.55.61A3.67 3.67 0 0 0 22 12.9v-.68A10.2 10.2 0 0 0 13 2Zm-1.82 13.09A3.09 3.09 0 1 1 14.27 12a3.1 3.1 0 0 1-3.09 3.09Z"/></g></g>`
	atOutlineInnerSVG                     = `<g id="evaAtOutline0"><g id="evaAtOutline1"><path id="evaAtOutline2" fill="currentColor" d="M13 2a10 10 0 0 0-5 19.1a10.15 10.15 0 0 0 4 .9a10 10 0 0 0 6.08-2.06a1 1 0 0 0 .19-1.4a1 1 0 0 0-1.41-.19A8 8 0 1 1 12.77 4A8.17 8.17 0 0 1 20 12.22v.68a1.71 1.71 0 0 1-1.78 1.7a1.82 1.82 0 0 1-1.62-1.88V8.4a1 1 0 0 0-1-1a1 1 0 0 0-1 .87a5 5 0 0 0-3.44-1.36A5.09 5.09 0 1 0 15.31 15a3.6 3.6 0 0 0 5.55.61A3.67 3.67 0 0 0 22 12.9v-.68A10.2 10.2 0 0 0 13 2Zm-1.82 13.09A3.09 3.09 0 1 1 14.27 12a3.1 3.1 0 0 1-3.09 3.09Z"/></g></g>`
	attachFillInnerSVG                    = `<g id="evaAttachFill0"><g id="evaAttachFill1"><path id="evaAttachFill2" fill="currentColor" d="M9.29 21a6.23 6.23 0 0 1-4.43-1.88a6 6 0 0 1-.22-8.49L12 3.2A4.11 4.11 0 0 1 15 2a4.48 4.48 0 0 1 3.19 1.35a4.36 4.36 0 0 1 .15 6.13l-7.4 7.43a2.54 2.54 0 0 1-1.81.75a2.72 2.72 0 0 1-1.95-.82a2.68 2.68 0 0 1-.08-3.77l6.83-6.86a1 1 0 0 1 1.37 1.41l-6.83 6.86a.68.68 0 0 0 .08.95a.78.78 0 0 0 .53.23a.56.56 0 0 0 .4-.16l7.39-7.43a2.36 2.36 0 0 0-.15-3.31a2.38 2.38 0 0 0-3.27-.15L6.06 12a4 4 0 0 0 .22 5.67a4.22 4.22 0 0 0 3 1.29a3.67 3.67 0 0 0 2.61-1.06l7.39-7.43a1 1 0 1 1 1.42 1.41l-7.39 7.43A5.65 5.65 0 0 1 9.29 21Z"/></g></g>`
	attachOutlineInnerSVG                 = `<g id="evaAttachOutline0"><g id="evaAttachOutline1"><path id="evaAttachOutline2" fill="currentColor" d="M9.29 21a6.23 6.23 0 0 1-4.43-1.88a6 6 0 0 1-.22-8.49L12 3.2A4.11 4.11 0 0 1 15 2a4.48 4.48 0 0 1 3.19 1.35a4.36 4.36 0 0 1 .15 6.13l-7.4 7.43a2.54 2.54 0 0 1-1.81.75a2.72 2.72 0 0 1-1.95-.82a2.68 2.68 0 0 1-.08-3.77l6.83-6.86a1 1 0 0 1 1.37 1.41l-6.83 6.86a.68.68 0 0 0 .08.95a.78.78 0 0 0 .53.23a.56.56 0 0 0 .4-.16l7.39-7.43a2.36 2.36 0 0 0-.15-3.31a2.38 2.38 0 0 0-3.27-.15L6.06 12a4 4 0 0 0 .22 5.67a4.22 4.22 0 0 0 3 1.29a3.67 3.67 0 0 0 2.61-1.06l7.39-7.43a1 1 0 1 1 1.42 1.41l-7.39 7.43A5.65 5.65 0 0 1 9.29 21Z"/></g></g>`
	attachTwoFillInnerSVG                 = `<g id="evaAttach2Fill0"><g id="evaAttach2Fill1"><path id="evaAttach2Fill2" fill="currentColor" d="M12 22a5.86 5.86 0 0 1-6-5.7V6.13A4.24 4.24 0 0 1 10.33 2a4.24 4.24 0 0 1 4.34 4.13v10.18a2.67 2.67 0 0 1-5.33 0V6.92a1 1 0 0 1 1-1a1 1 0 0 1 1 1v9.39a.67.67 0 0 0 1.33 0V6.13A2.25 2.25 0 0 0 10.33 4A2.25 2.25 0 0 0 8 6.13V16.3a3.86 3.86 0 0 0 4 3.7a3.86 3.86 0 0 0 4-3.7V6.13a1 1 0 1 1 2 0V16.3a5.86 5.86 0 0 1-6 5.7Z"/></g></g>`
	attachTwoOutlineInnerSVG              = `<g id="evaAttach2Outline0"><g id="evaAttach2Outline1"><path id="evaAttach2Outline2" fill="currentColor" d="M12 22a5.86 5.86 0 0 1-6-5.7V6.13A4.24 4.24 0 0 1 10.33 2a4.24 4.24 0 0 1 4.34 4.13v10.18a2.67 2.67 0 0 1-5.33 0V6.92a1 1 0 0 1 1-1a1 1 0 0 1 1 1v9.39a.67.67 0 0 0 1.33 0V6.13A2.25 2.25 0 0 0 10.33 4A2.25 2.25 0 0 0 8 6.13V16.3a3.86 3.86 0 0 0 4 3.7a3.86 3.86 0 0 0 4-3.7V6.13a1 1 0 1 1 2 0V16.3a5.86 5.86 0 0 1-6 5.7Z"/></g></g>`
	awardFillInnerSVG                     = `<g id="evaAwardFill0"><g id="evaAwardFill1"><path id="evaAwardFill2" fill="currentColor" d="m19 20.75l-2.31-9A5.94 5.94 0 0 0 18 8A6 6 0 0 0 6 8a5.94 5.94 0 0 0 1.34 3.77L5 20.75a1 1 0 0 0 1.48 1.11l5.33-3.13l5.68 3.14A.91.91 0 0 0 18 22a1 1 0 0 0 1-1.25ZM12 4a4 4 0 1 1-4 4a4 4 0 0 1 4-4Z"/></g></g>`
	awardOutlineInnerSVG                  = `<g id="evaAwardOutline0"><g id="evaAwardOutline1"><path id="evaAwardOutline2" fill="currentColor" d="m19 20.75l-2.31-9A5.94 5.94 0 0 0 18 8A6 6 0 0 0 6 8a5.94 5.94 0 0 0 1.34 3.77L5 20.75a1 1 0 0 0 1.48 1.11l5.33-3.13l5.68 3.14A.91.91 0 0 0 18 22a1 1 0 0 0 1-1.25ZM12 4a4 4 0 1 1-4 4a4 4 0 0 1 4-4Zm.31 12.71a1 1 0 0 0-1 0l-3.75 2.2L9 13.21a5.94 5.94 0 0 0 5.92 0L16.45 19Z"/></g></g>`
	backspaceFillInnerSVG                 = `<g id="evaBackspaceFill0"><g id="evaBackspaceFill1"><path id="evaBackspaceFill2" fill="currentColor" d="M20.14 4h-9.77a3 3 0 0 0-2 .78l-.1.11l-6 7.48a1 1 0 0 0 .11 1.37l6 5.48a3 3 0 0 0 2 .78h9.77A1.84 1.84 0 0 0 22 18.18V5.82A1.84 1.84 0 0 0 20.14 4Zm-3.43 9.29a1 1 0 0 1 0 1.42a1 1 0 0 1-1.42 0L14 13.41l-1.29 1.3a1 1 0 0 1-1.42 0a1 1 0 0 1 0-1.42l1.3-1.29l-1.3-1.29a1 1 0 0 1 1.42-1.42l1.29 1.3l1.29-1.3a1 1 0 0 1 1.42 1.42L15.41 12Z"/></g></g>`
	backspaceOutlineInnerSVG              = `<g id="evaBackspaceOutline0"><g id="evaBackspaceOutline1"><g id="evaBackspaceOutline2" fill="currentColor"><path d="M20.14 4h-9.77a3 3 0 0 0-2 .78l-.1.11l-6 7.48a1 1 0 0 0 .11 1.37l6 5.48a3 3 0 0 0 2 .78h9.77A1.84 1.84 0 0 0 22 18.18V5.82A1.84 1.84 0 0 0 20.14 4ZM20 18h-9.63a1 1 0 0 1-.67-.26l-5.33-4.85l5.38-6.67a1 1 0 0 1 .62-.22H20Z"/><path d="M11.29 14.71a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L15.41 12l1.3-1.29a1 1 0 0 0-1.42-1.42L14 10.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42Z"/></g></g></g>`
	barChartFillInnerSVG                  = `<g id="evaBarChartFill0"><g id="evaBarChartFill1"><path id="evaBarChartFill2" fill="currentColor" d="M12 4a1 1 0 0 0-1 1v15a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1Zm7 8a1 1 0 0 0-1 1v7a1 1 0 0 0 2 0v-7a1 1 0 0 0-1-1ZM5 8a1 1 0 0 0-1 1v11a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1Z"/></g></g>`
	barChartOutlineInnerSVG               = `<g id="evaBarChartOutline0"><g id="evaBarChartOutline1"><path id="evaBarChartOutline2" fill="currentColor" d="M12 4a1 1 0 0 0-1 1v15a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1Zm7 8a1 1 0 0 0-1 1v7a1 1 0 0 0 2 0v-7a1 1 0 0 0-1-1ZM5 8a1 1 0 0 0-1 1v11a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1Z"/></g></g>`
	barChartTwoFillInnerSVG               = `<g id="evaBarChart2Fill0"><g id="evaBarChart2Fill1"><path id="evaBarChart2Fill2" fill="currentColor" d="M12 8a1 1 0 0 0-1 1v11a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1Zm7-4a1 1 0 0 0-1 1v15a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1ZM5 12a1 1 0 0 0-1 1v7a1 1 0 0 0 2 0v-7a1 1 0 0 0-1-1Z"/></g></g>`
	barChartTwoOutlineInnerSVG            = `<g id="evaBarChart2Outline0"><g id="evaBarChart2Outline1"><path id="evaBarChart2Outline2" fill="currentColor" d="M12 8a1 1 0 0 0-1 1v11a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1Zm7-4a1 1 0 0 0-1 1v15a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1ZM5 12a1 1 0 0 0-1 1v7a1 1 0 0 0 2 0v-7a1 1 0 0 0-1-1Z"/></g></g>`
	batteryFillInnerSVG                   = `<g id="evaBatteryFill0"><g id="evaBatteryFill1"><path id="evaBatteryFill2" fill="currentColor" d="M15.83 6H4.17A2.31 2.31 0 0 0 2 8.43v7.14A2.31 2.31 0 0 0 4.17 18h11.66A2.31 2.31 0 0 0 18 15.57V8.43A2.31 2.31 0 0 0 15.83 6ZM21 9a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1Z"/></g></g>`
	batteryOutlineInnerSVG                = `<g id="evaBatteryOutline0"><g id="evaBatteryOutline1"><path id="evaBatteryOutline2" fill="currentColor" d="M15.83 6H4.17A2.31 2.31 0 0 0 2 8.43v7.14A2.31 2.31 0 0 0 4.17 18h11.66A2.31 2.31 0 0 0 18 15.57V8.43A2.31 2.31 0 0 0 15.83 6Zm.17 9.57a.52.52 0 0 1-.17.43H4.18a.5.5 0 0 1-.18-.43V8.43A.53.53 0 0 1 4.17 8h11.65a.5.5 0 0 1 .18.43ZM21 9a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1Z"/></g></g>`
	behanceFillInnerSVG                   = `<g id="evaBehanceFill0"><g id="evaBehanceFill1"><g id="evaBehanceFill2" fill="currentColor"><path d="M14.76 11.19a1 1 0 0 0-1 1.09h2.06a1 1 0 0 0-1.06-1.09ZM9.49 12.3H8.26v1.94h1c1 0 1.44-.33 1.44-1s-.46-.94-1.21-.94Zm.87-1.78c0-.53-.35-.85-.95-.85H8.26v1.74h.85c.89 0 1.25-.32 1.25-.89Z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2ZM9.7 15.2H7V8.7h2.7c1.17 0 1.94.61 1.94 1.6a1.4 1.4 0 0 1-1.12 1.43A1.52 1.52 0 0 1 12 13.37c0 1.16-1 1.83-2.3 1.83Zm3.55-6h3v.5h-3ZM17 13.05h-3.3v.14a1.07 1.07 0 0 0 1.09 1.19a.9.9 0 0 0 1-.63H17a2 2 0 0 1-2.17 1.55a2.15 2.15 0 0 1-2.36-2.3v-.44a2.11 2.11 0 0 1 2.28-2.25A2.12 2.12 0 0 1 17 12.58Z"/></g></g></g>`
	behanceOutlineInnerSVG                = `<g id="evaBehanceOutline0"><g id="evaBehanceOutline1"><g id="evaBehanceOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M10.52 11.78a1.4 1.4 0 0 0 1.12-1.43c0-1-.77-1.6-1.94-1.6H7v6.5h2.7c1.3-.05 2.3-.72 2.3-1.88a1.52 1.52 0 0 0-1.48-1.59ZM8.26 9.67h1.15c.6 0 .95.32.95.85s-.38.89-1.25.89h-.85Zm1 4.57h-1V12.3h1.23c.75 0 1.17.38 1.17 1s-.42.94-1.44.94Zm5.49-3.94a2.11 2.11 0 0 0-2.28 2.25V13a2.15 2.15 0 0 0 2.34 2.31A2 2 0 0 0 17 13.75h-1.21a.9.9 0 0 1-1 .63a1.07 1.07 0 0 1-1.09-1.19v-.14H17v-.47a2.12 2.12 0 0 0-2.25-2.28Zm1 2h-2.02a1 1 0 0 1 1-1.09a1 1 0 0 1 1 1.09Zm-2.5-3.1h3v.5h-3z"/></g></g></g>`
	bellFillInnerSVG                      = `<g id="evaBellFill0"><g id="evaBellFill1"><path id="evaBellFill2" fill="currentColor" d="m20.52 15.21l-1.8-1.81V8.94a6.86 6.86 0 0 0-5.82-6.88a6.74 6.74 0 0 0-7.62 6.67v4.67l-1.8 1.81A1.64 1.64 0 0 0 4.64 18H8v.34A3.84 3.84 0 0 0 12 22a3.84 3.84 0 0 0 4-3.66V18h3.36a1.64 1.64 0 0 0 1.16-2.79ZM14 18.34A1.88 1.88 0 0 1 12 20a1.88 1.88 0 0 1-2-1.66V18h4Z"/></g></g>`
	bellOffFillInnerSVG                   = `<g id="evaBellOffFill0"><g id="evaBellOffFill1"><path id="evaBellOffFill2" fill="currentColor" d="m15.88 18.71l-.59-.59L14 16.78l-.07-.07L6.58 9.4L5.31 8.14a5.68 5.68 0 0 0 0 .59v4.67l-1.8 1.81A1.64 1.64 0 0 0 4.64 18H8v.34A3.84 3.84 0 0 0 12 22a3.88 3.88 0 0 0 4-3.22ZM14 18.34A1.88 1.88 0 0 1 12 20a1.88 1.88 0 0 1-2-1.66V18h4ZM7.13 4.3l1.46 1.46l9.53 9.53l2 2l.31.3a1.58 1.58 0 0 0 .45-.6a1.62 1.62 0 0 0-.35-1.78l-1.8-1.81V8.94a6.86 6.86 0 0 0-5.83-6.88a6.71 6.71 0 0 0-5.32 1.61a6.88 6.88 0 0 0-.58.54Zm13.58 14.99L19.41 18l-2-2l-9.52-9.53L6.42 5L4.71 3.29a1 1 0 0 0-1.42 1.42L5.53 7l1.75 1.7l7.31 7.3l.07.07L16 17.41l.59.59l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	bellOffOutlineInnerSVG                = `<g id="evaBellOffOutline0"><g id="evaBellOffOutline1"><path id="evaBellOffOutline2" fill="currentColor" d="M8.9 5.17A4.67 4.67 0 0 1 12.64 4a4.86 4.86 0 0 1 4.08 4.9v4.5a1.92 1.92 0 0 0 .1.59l3.6 3.6a1.58 1.58 0 0 0 .45-.6a1.62 1.62 0 0 0-.35-1.78l-1.8-1.81V8.94a6.86 6.86 0 0 0-5.82-6.88a6.71 6.71 0 0 0-5.32 1.61a6.88 6.88 0 0 0-.58.54l1.47 1.43a4.79 4.79 0 0 1 .43-.47ZM14 16.86l-.83-.86H5.51l1.18-1.18a2 2 0 0 0 .59-1.42v-3.29l-2-2a5.68 5.68 0 0 0 0 .59v4.7l-1.8 1.81A1.63 1.63 0 0 0 4.64 18H8v.34A3.84 3.84 0 0 0 12 22a3.88 3.88 0 0 0 4-3.22l-.83-.78ZM12 20a1.88 1.88 0 0 1-2-1.66V18h4v.34A1.88 1.88 0 0 1 12 20Zm8.71-.71L19.41 18l-2-2l-9.52-9.53L6.42 5L4.71 3.29a1 1 0 0 0-1.42 1.42L5.53 7l1.75 1.7l7.31 7.3l.07.07L16 17.41l.59.59l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	bellOutlineInnerSVG                   = `<g id="evaBellOutline0"><g id="evaBellOutline1"><path id="evaBellOutline2" fill="currentColor" d="m20.52 15.21l-1.8-1.81V8.94a6.86 6.86 0 0 0-5.82-6.88a6.74 6.74 0 0 0-7.62 6.67v4.67l-1.8 1.81A1.64 1.64 0 0 0 4.64 18H8v.34A3.84 3.84 0 0 0 12 22a3.84 3.84 0 0 0 4-3.66V18h3.36a1.64 1.64 0 0 0 1.16-2.79ZM14 18.34A1.88 1.88 0 0 1 12 20a1.88 1.88 0 0 1-2-1.66V18h4ZM5.51 16l1.18-1.18a2 2 0 0 0 .59-1.42V8.73A4.73 4.73 0 0 1 8.9 5.17A4.67 4.67 0 0 1 12.64 4a4.86 4.86 0 0 1 4.08 4.9v4.5a2 2 0 0 0 .58 1.42L18.49 16Z"/></g></g>`
	bluetoothFillInnerSVG                 = `<g id="evaBluetoothFill0"><g id="evaBluetoothFill1"><path id="evaBluetoothFill2" fill="currentColor" d="m13.63 12l4-3.79a1.14 1.14 0 0 0-.13-1.77l-4.67-3.23a1.17 1.17 0 0 0-1.21-.08a1.15 1.15 0 0 0-.62 1v6.2l-3.19-4a1 1 0 0 0-1.56 1.3L9.72 12l-3.5 4.43a1 1 0 0 0 .16 1.4A1 1 0 0 0 7 18a1 1 0 0 0 .78-.38L11 13.56v6.29A1.16 1.16 0 0 0 12.16 21a1.16 1.16 0 0 0 .67-.21l4.64-3.18a1.17 1.17 0 0 0 .49-.85a1.15 1.15 0 0 0-.34-.91ZM13 5.76l2.5 1.73L13 9.85Zm0 12.49v-4.07l2.47 2.38Z"/></g></g>`
	bluetoothOutlineInnerSVG              = `<g id="evaBluetoothOutline0"><g id="evaBluetoothOutline1"><path id="evaBluetoothOutline2" fill="currentColor" d="m13.63 12l4-3.79a1.14 1.14 0 0 0-.13-1.77l-4.67-3.23a1.17 1.17 0 0 0-1.21-.08a1.15 1.15 0 0 0-.62 1v6.2l-3.19-4a1 1 0 0 0-1.56 1.3L9.72 12l-3.5 4.43a1 1 0 0 0 .16 1.4A1 1 0 0 0 7 18a1 1 0 0 0 .78-.38L11 13.56v6.29A1.16 1.16 0 0 0 12.16 21a1.16 1.16 0 0 0 .67-.21l4.64-3.18a1.17 1.17 0 0 0 .49-.85a1.15 1.15 0 0 0-.34-.91ZM13 5.76l2.5 1.73L13 9.85Zm0 12.49v-4.07l2.47 2.38Z"/></g></g>`
	bookFillInnerSVG                      = `<g id="evaBookFill0"><g id="evaBookFill1"><path id="evaBookFill2" fill="currentColor" d="M19 3H7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1ZM7 19a1 1 0 0 1 0-2h11v2Z"/></g></g>`
	bookOpenFillInnerSVG                  = `<g id="evaBookOpenFill0"><g id="evaBookOpenFill1"><path id="evaBookOpenFill2" fill="currentColor" d="M21 4.34a1.24 1.24 0 0 0-1.08-.23L13 5.89v14.27l7.56-1.94A1.25 1.25 0 0 0 21.5 17V5.32a1.25 1.25 0 0 0-.5-.98ZM11 5.89L4.06 4.11A1.27 1.27 0 0 0 3 4.34a1.25 1.25 0 0 0-.48 1V17a1.25 1.25 0 0 0 .94 1.21L11 20.16Z"/></g></g>`
	bookOpenOutlineInnerSVG               = `<g id="evaBookOpenOutline0"><g id="evaBookOpenOutline1"><path id="evaBookOpenOutline2" fill="currentColor" d="M20.62 4.22a1 1 0 0 0-.84-.2L12 5.77L4.22 4A1 1 0 0 0 3 5v12.2a1 1 0 0 0 .78 1l8 1.8h.44l8-1.8a1 1 0 0 0 .78-1V5a1 1 0 0 0-.38-.78ZM5 6.25l6 1.35v10.15L5 16.4ZM19 16.4l-6 1.35V7.6l6-1.35Z"/></g></g>`
	bookOutlineInnerSVG                   = `<g id="evaBookOutline0"><g id="evaBookOutline1"><path id="evaBookOutline2" fill="currentColor" d="M19 3H7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1ZM7 5h11v10H7a3 3 0 0 0-1 .18V6a1 1 0 0 1 1-1Zm0 14a1 1 0 0 1 0-2h11v2Z"/></g></g>`
	bookmarkFillInnerSVG                  = `<g id="evaBookmarkFill0"><g id="evaBookmarkFill1"><path id="evaBookmarkFill2" fill="currentColor" d="M6 21a1 1 0 0 1-.49-.13A1 1 0 0 1 5 20V5.33A2.28 2.28 0 0 1 7.2 3h9.6A2.28 2.28 0 0 1 19 5.33V20a1 1 0 0 1-.5.86a1 1 0 0 1-1 0l-5.67-3.21l-5.33 3.2A1 1 0 0 1 6 21Z"/></g></g>`
	bookmarkOutlineInnerSVG               = `<g id="evaBookmarkOutline0"><g id="evaBookmarkOutline1"><path id="evaBookmarkOutline2" fill="currentColor" d="M6.09 21.06a1 1 0 0 1-1-1L4.94 5.4a2.26 2.26 0 0 1 2.18-2.35L16.71 3a2.27 2.27 0 0 1 2.23 2.31l.14 14.66a1 1 0 0 1-.49.87a1 1 0 0 1-1 0l-5.7-3.16l-5.29 3.23a1.2 1.2 0 0 1-.51.15Zm5.76-5.55a1.11 1.11 0 0 1 .5.12l4.71 2.61l-.12-12.95c0-.2-.13-.34-.21-.33l-9.6.09c-.08 0-.19.13-.19.33l.12 12.9l4.28-2.63a1.06 1.06 0 0 1 .51-.14Z"/></g></g>`
	briefcaseFillInnerSVG                 = `<g id="evaBriefcaseFill0"><g id="evaBriefcaseFill1"><path id="evaBriefcaseFill2" fill="currentColor" d="M7 21h10V7h-1V5.5A2.5 2.5 0 0 0 13.5 3h-3A2.5 2.5 0 0 0 8 5.5V7H7Zm3-15.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5V7h-4ZM19 7v14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3ZM5 7a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3Z"/></g></g>`
	briefcaseOutlineInnerSVG              = `<g id="evaBriefcaseOutline0"><g id="evaBriefcaseOutline1"><path id="evaBriefcaseOutline2" fill="currentColor" d="M19 7h-3V5.5A2.5 2.5 0 0 0 13.5 3h-3A2.5 2.5 0 0 0 8 5.5V7H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-4 2v10H9V9Zm-5-3.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5V7h-4ZM4 18v-8a1 1 0 0 1 1-1h2v10H5a1 1 0 0 1-1-1Zm16 0a1 1 0 0 1-1 1h-2V9h2a1 1 0 0 1 1 1Z"/></g></g>`
	browserFillInnerSVG                   = `<g id="evaBrowserFill0"><g id="evaBrowserFill1"><path id="evaBrowserFill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm-6 3a1 1 0 1 1-1 1a1 1 0 0 1 1-1ZM8 6a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm11 12a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7h14Z"/></g></g>`
	browserOutlineInnerSVG                = `<g id="evaBrowserOutline0"><g id="evaBrowserOutline1"><g id="evaBrowserOutline2" fill="currentColor"><path d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm1 15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7h14ZM5 9V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3Z"/><circle cx="8" cy="7.03" r="1"/><circle cx="12" cy="7.03" r="1"/></g></g></g>`
	brushFillInnerSVG                     = `<g id="evaBrushFill0"><g id="evaBrushFill1"><path id="evaBrushFill2" fill="currentColor" d="M7.12 12.55a4 4 0 0 0-3.07 3.86v3.11a.47.47 0 0 0 .48.48l3.24-.06a3.78 3.78 0 0 0 3.44-2.2a3.65 3.65 0 0 0-4.09-5.19Zm12.14-8.09a2.14 2.14 0 0 0-2.88.21L10 11.08a.47.47 0 0 0 0 .66L12.25 14a.47.47 0 0 0 .66 0l6.49-6.47a2.06 2.06 0 0 0 .6-1.47a2 2 0 0 0-.74-1.6Z"/></g></g>`
	brushOutlineInnerSVG                  = `<g id="evaBrushOutline0"><g id="evaBrushOutline1"><path id="evaBrushOutline2" fill="currentColor" d="M20 6.83a2.76 2.76 0 0 0-.82-2a2.89 2.89 0 0 0-4 0l-6.6 6.6h-.22a4.42 4.42 0 0 0-4.3 4.31L4 19a1 1 0 0 0 .29.73A1.05 1.05 0 0 0 5 20l3.26-.06a4.42 4.42 0 0 0 4.31-4.3v-.23l6.61-6.6A2.74 2.74 0 0 0 20 6.83ZM8.25 17.94L6 18v-2.23a2.4 2.4 0 0 1 2.4-2.36a2.15 2.15 0 0 1 2.15 2.19a2.4 2.4 0 0 1-2.3 2.34Zm9.52-10.55l-5.87 5.87a4.55 4.55 0 0 0-.52-.64a3.94 3.94 0 0 0-.64-.52l5.87-5.86a.84.84 0 0 1 1.16 0a.81.81 0 0 1 .23.59a.79.79 0 0 1-.23.56Z"/></g></g>`
	bulbFillInnerSVG                      = `<g id="evaBulbFill0"><g id="evaBulbFill1"><path id="evaBulbFill2" fill="currentColor" d="M12 7a5 5 0 0 0-3 9v4a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-4a5 5 0 0 0-3-9Zm0-1a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1Zm9 5h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2ZM5 11H3a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2Zm2.66-4.58L6.22 5a1 1 0 0 0-1.39 1.47l1.44 1.39a1 1 0 0 0 .73.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.06-1.41Zm11.53-1.37a1 1 0 0 0-1.41 0l-1.44 1.37a1 1 0 0 0 0 1.41a1 1 0 0 0 .72.31a1 1 0 0 0 .69-.28l1.44-1.39a1 1 0 0 0 0-1.42Z"/></g></g>`
	bulbOutlineInnerSVG                   = `<g id="evaBulbOutline0"><g id="evaBulbOutline1"><path id="evaBulbOutline2" fill="currentColor" d="M12 7a5 5 0 0 0-3 9v4a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-4a5 5 0 0 0-3-9Zm1.5 7.59a1 1 0 0 0-.5.87V20h-2v-4.54a1 1 0 0 0-.5-.87A3 3 0 0 1 9 12a3 3 0 0 1 6 0a3 3 0 0 1-1.5 2.59ZM12 6a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1Zm9 5h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2ZM5 11H3a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2Zm2.66-4.58L6.22 5a1 1 0 0 0-1.39 1.47l1.44 1.39a1 1 0 0 0 .73.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.06-1.41Zm11.53-1.37a1 1 0 0 0-1.41 0l-1.44 1.37a1 1 0 0 0 0 1.41a1 1 0 0 0 .72.31a1 1 0 0 0 .69-.28l1.44-1.39a1 1 0 0 0 0-1.42Z"/></g></g>`
	calendarFillInnerSVG                  = `<g id="evaCalendarFill0"><g id="evaCalendarFill1"><path id="evaCalendarFill2" fill="currentColor" d="M18 4h-1V3a1 1 0 0 0-2 0v1H9V3a1 1 0 0 0-2 0v1H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3ZM8 17a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm8 0h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Zm3-6H5V7a1 1 0 0 1 1-1h1v1a1 1 0 0 0 2 0V6h6v1a1 1 0 0 0 2 0V6h1a1 1 0 0 1 1 1Z"/></g></g>`
	calendarOutlineInnerSVG               = `<g id="evaCalendarOutline0"><g id="evaCalendarOutline1"><g id="evaCalendarOutline2" fill="currentColor"><path d="M18 4h-1V3a1 1 0 0 0-2 0v1H9V3a1 1 0 0 0-2 0v1H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3ZM6 6h1v1a1 1 0 0 0 2 0V6h6v1a1 1 0 0 0 2 0V6h1a1 1 0 0 1 1 1v4H5V7a1 1 0 0 1 1-1Zm12 14H6a1 1 0 0 1-1-1v-6h14v6a1 1 0 0 1-1 1Z"/><circle cx="8" cy="16" r="1"/><path d="M16 15h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Z"/></g></g></g>`
	cameraFillInnerSVG                    = `<g id="evaCameraFill0"><g id="evaCameraFill1"><g id="evaCameraFill2" fill="currentColor"><circle cx="12" cy="14" r="1.5"/><path d="M19 7h-3V5.5A2.5 2.5 0 0 0 13.5 3h-3A2.5 2.5 0 0 0 8 5.5V7H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-9-1.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5V7h-4Zm2 12a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5Z"/></g></g></g>`
	cameraOutlineInnerSVG                 = `<g id="evaCameraOutline0"><g id="evaCameraOutline1"><g id="evaCameraOutline2" fill="currentColor"><path d="M19 7h-3V5.5A2.5 2.5 0 0 0 13.5 3h-3A2.5 2.5 0 0 0 8 5.5V7H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-9-1.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5V7h-4ZM20 18a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/><path d="M12 10.5a3.5 3.5 0 1 0 3.5 3.5a3.5 3.5 0 0 0-3.5-3.5Zm0 5a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Z"/></g></g></g>`
	carFillInnerSVG                       = `<g id="evaCarFill0"><g id="evaCarFill1"><path id="evaCarFill2" fill="currentColor" d="M21.6 11.22L17 7.52V5a1.91 1.91 0 0 0-1.81-2H3.79A1.91 1.91 0 0 0 2 5v10a2 2 0 0 0 1.2 1.88a3 3 0 1 0 5.6.12h6.36a3 3 0 1 0 5.64 0h.2a1 1 0 0 0 1-1v-4a1 1 0 0 0-.4-.78ZM20 12.48V15h-3v-4.92ZM7 18a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm12 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/></g></g>`
	carOutlineInnerSVG                    = `<g id="evaCarOutline0"><g id="evaCarOutline1"><path id="evaCarOutline2" fill="currentColor" d="M21.6 11.22L17 7.52V5a1.91 1.91 0 0 0-1.81-2H3.79A1.91 1.91 0 0 0 2 5v10a2 2 0 0 0 1.2 1.88a3 3 0 1 0 5.6.12h6.36a3 3 0 1 0 5.64 0h.2a1 1 0 0 0 1-1v-4a1 1 0 0 0-.4-.78ZM20 12.48V15h-3v-4.92ZM7 18a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm5-3H4V5h11v10Zm7 3a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/></g></g>`
	castFillInnerSVG                      = `<g id="evaCastFill0"><g id="evaCastFill1"><g id="evaCastFill2" fill="currentColor"><path d="M18.4 3H5.6A2.7 2.7 0 0 0 3 5.78V7a1 1 0 0 0 2 0V5.78A.72.72 0 0 1 5.6 5h12.8a.72.72 0 0 1 .6.78v12.44a.72.72 0 0 1-.6.78H17a1 1 0 0 0 0 2h1.4a2.7 2.7 0 0 0 2.6-2.78V5.78A2.7 2.7 0 0 0 18.4 3ZM3.86 14A1 1 0 0 0 3 15.17a1 1 0 0 0 1.14.83a2.49 2.49 0 0 1 2.12.72a2.52 2.52 0 0 1 .51 2.84a1 1 0 0 0 .48 1.33a1.06 1.06 0 0 0 .42.09a1 1 0 0 0 .91-.58A4.52 4.52 0 0 0 3.86 14Z"/><path d="M3.86 10.08a1 1 0 0 0 .28 2a6 6 0 0 1 5.09 1.71a6 6 0 0 1 1.53 5.95a1 1 0 0 0 .68 1.26a.9.9 0 0 0 .28 0a1 1 0 0 0 1-.72a8 8 0 0 0-8.82-10.2Z"/><circle cx="4" cy="19" r="1"/></g></g></g>`
	castOutlineInnerSVG                   = `<g id="evaCastOutline0"><g id="evaCastOutline1"><g id="evaCastOutline2" fill="currentColor"><path d="M18.4 3H5.6A2.7 2.7 0 0 0 3 5.78V7a1 1 0 0 0 2 0V5.78A.72.72 0 0 1 5.6 5h12.8a.72.72 0 0 1 .6.78v12.44a.72.72 0 0 1-.6.78H17a1 1 0 0 0 0 2h1.4a2.7 2.7 0 0 0 2.6-2.78V5.78A2.7 2.7 0 0 0 18.4 3ZM3.86 14A1 1 0 0 0 3 15.17a1 1 0 0 0 1.14.83a2.49 2.49 0 0 1 2.12.72a2.52 2.52 0 0 1 .51 2.84a1 1 0 0 0 .48 1.33a1.06 1.06 0 0 0 .42.09a1 1 0 0 0 .91-.58A4.52 4.52 0 0 0 3.86 14Z"/><path d="M3.86 10.08a1 1 0 0 0 .28 2a6 6 0 0 1 5.09 1.71a6 6 0 0 1 1.53 5.95a1 1 0 0 0 .68 1.26a.9.9 0 0 0 .28 0a1 1 0 0 0 1-.72a8 8 0 0 0-8.82-10.2Z"/><circle cx="4" cy="19" r="1"/></g></g></g>`
	chargingFillInnerSVG                  = `<g id="evaChargingFill0"><g id="evaChargingFill1"><g id="evaChargingFill2" fill="currentColor"><path d="M11.28 13H7a1 1 0 0 1-.86-.5a1 1 0 0 1 0-1L9.28 6H4.17A2.31 2.31 0 0 0 2 8.43v7.14A2.31 2.31 0 0 0 4.17 18h4.25Z"/><path d="M15.83 6h-4.25l-2.86 5H13a1 1 0 0 1 .86.5a1 1 0 0 1 0 1L10.72 18h5.11A2.31 2.31 0 0 0 18 15.57V8.43A2.31 2.31 0 0 0 15.83 6ZM21 9a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1Z"/></g></g></g>`
	chargingOutlineInnerSVG               = `<g id="evaChargingOutline0"><g id="evaChargingOutline1"><g id="evaChargingOutline2" fill="currentColor"><path d="M21 9a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1Zm-5.17-3h-3.1l-1.14 2h4.23a.5.5 0 0 1 .18.43v7.14a.52.52 0 0 1-.17.43H13l-1.15 2h4A2.31 2.31 0 0 0 18 15.57V8.43A2.31 2.31 0 0 0 15.83 6ZM4 15.57V8.43A.53.53 0 0 1 4.17 8H7l1.13-2h-4A2.31 2.31 0 0 0 2 8.43v7.14A2.31 2.31 0 0 0 4.17 18h3.1l1.14-2H4.18a.5.5 0 0 1-.18-.43Z"/><path d="M9 20a1 1 0 0 1-.87-1.5l3.15-5.5H7a1 1 0 0 1-.86-.5a1 1 0 0 1 0-1l4-7a1 1 0 0 1 1.74 1L8.72 11H13a1 1 0 0 1 .86.5a1 1 0 0 1 0 1l-4 7A1 1 0 0 1 9 20Z"/></g></g></g>`
	checkmarkCircleFillInnerSVG           = `<g id="evaCheckmarkCircleFill0"><g id="evaCheckmarkCircleFill1"><g id="evaCheckmarkCircleFill2" fill="currentColor"><path d="M9.71 11.29a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 16a1 1 0 0 0 .72-.34l7-8a1 1 0 0 0-1.5-1.32L12 13.54Z"/><path d="M21 11a1 1 0 0 0-1 1a8 8 0 0 1-8 8A8 8 0 0 1 6.33 6.36A7.93 7.93 0 0 1 12 4a8.79 8.79 0 0 1 1.9.22a1 1 0 1 0 .47-1.94A10.54 10.54 0 0 0 12 2a10 10 0 0 0-7 17.09A9.93 9.93 0 0 0 12 22a10 10 0 0 0 10-10a1 1 0 0 0-1-1Z"/></g></g></g>`
	checkmarkCircleOutlineInnerSVG        = `<g id="evaCheckmarkCircleOutline0"><g id="evaCheckmarkCircleOutline1"><g id="evaCheckmarkCircleOutline2" fill="currentColor"><path d="M9.71 11.29a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 16a1 1 0 0 0 .72-.34l7-8a1 1 0 0 0-1.5-1.32L12 13.54Z"/><path d="M21 11a1 1 0 0 0-1 1a8 8 0 0 1-8 8A8 8 0 0 1 6.33 6.36A7.93 7.93 0 0 1 12 4a8.79 8.79 0 0 1 1.9.22a1 1 0 1 0 .47-1.94A10.54 10.54 0 0 0 12 2a10 10 0 0 0-7 17.09A9.93 9.93 0 0 0 12 22a10 10 0 0 0 10-10a1 1 0 0 0-1-1Z"/></g></g></g>`
	checkmarkCircleTwoFillInnerSVG        = `<g id="evaCheckmarkCircle2Fill0"><g id="evaCheckmarkCircle2Fill1"><path id="evaCheckmarkCircle2Fill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm4.3 7.61l-4.57 6a1 1 0 0 1-.79.39a1 1 0 0 1-.79-.38l-2.44-3.11a1 1 0 0 1 1.58-1.23l1.63 2.08l3.78-5a1 1 0 1 1 1.6 1.22Z"/></g></g>`
	checkmarkCircleTwoOutlineInnerSVG     = `<g id="evaCheckmarkCircle2Outline0"><g id="evaCheckmarkCircle2Outline1"><g id="evaCheckmarkCircle2Outline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="m14.7 8.39l-3.78 5l-1.63-2.11a1 1 0 0 0-1.58 1.23l2.43 3.11a1 1 0 0 0 .79.38a1 1 0 0 0 .79-.39l4.57-6a1 1 0 1 0-1.6-1.22Z"/></g></g></g>`
	checkmarkFillInnerSVG                 = `<g id="evaCheckmarkFill0"><g id="evaCheckmarkFill1"><path id="evaCheckmarkFill2" fill="currentColor" d="M9.86 18a1 1 0 0 1-.73-.32l-4.86-5.17a1 1 0 1 1 1.46-1.37l4.12 4.39l8.41-9.2a1 1 0 1 1 1.48 1.34l-9.14 10a1 1 0 0 1-.73.33Z"/></g></g>`
	checkmarkOutlineInnerSVG              = `<g id="evaCheckmarkOutline0"><g id="evaCheckmarkOutline1"><path id="evaCheckmarkOutline2" fill="currentColor" d="M9.86 18a1 1 0 0 1-.73-.32l-4.86-5.17a1 1 0 1 1 1.46-1.37l4.12 4.39l8.41-9.2a1 1 0 1 1 1.48 1.34l-9.14 10a1 1 0 0 1-.73.33Z"/></g></g>`
	checkmarkSquareFillInnerSVG           = `<g id="evaCheckmarkSquareFill0"><g id="evaCheckmarkSquareFill1"><g id="evaCheckmarkSquareFill2" fill="currentColor"><path d="M20 11.83a1 1 0 0 0-1 1v5.57a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h9.57a1 1 0 1 0 0-2H5.6A2.61 2.61 0 0 0 3 5.6v12.8A2.61 2.61 0 0 0 5.6 21h12.8a2.61 2.61 0 0 0 2.6-2.6v-5.57a1 1 0 0 0-1-1Z"/><path d="M10.72 11a1 1 0 0 0-1.44 1.38l2.22 2.33a1 1 0 0 0 .72.31a1 1 0 0 0 .72-.3l6.78-7a1 1 0 1 0-1.44-1.4l-6.05 6.26Z"/></g></g></g>`
	checkmarkSquareOutlineInnerSVG        = `<g id="evaCheckmarkSquareOutline0"><g id="evaCheckmarkSquareOutline1"><g id="evaCheckmarkSquareOutline2" fill="currentColor"><path d="M20 11.83a1 1 0 0 0-1 1v5.57a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h9.57a1 1 0 1 0 0-2H5.6A2.61 2.61 0 0 0 3 5.6v12.8A2.61 2.61 0 0 0 5.6 21h12.8a2.61 2.61 0 0 0 2.6-2.6v-5.57a1 1 0 0 0-1-1Z"/><path d="M10.72 11a1 1 0 0 0-1.44 1.38l2.22 2.33a1 1 0 0 0 .72.31a1 1 0 0 0 .72-.3l6.78-7a1 1 0 1 0-1.44-1.4l-6.05 6.26Z"/></g></g></g>`
	checkmarkSquareTwoFillInnerSVG        = `<g id="evaCheckmarkSquare2Fill0"><g id="evaCheckmarkSquare2Fill1"><path id="evaCheckmarkSquare2Fill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm-1.7 6.61l-4.57 6a1 1 0 0 1-.79.39a1 1 0 0 1-.79-.38l-2.44-3.11a1 1 0 0 1 1.58-1.23l1.63 2.08l3.78-5a1 1 0 1 1 1.6 1.22Z"/></g></g>`
	checkmarkSquareTwoOutlineInnerSVG     = `<g id="evaCheckmarkSquare2Outline0"><g id="evaCheckmarkSquare2Outline1"><g id="evaCheckmarkSquare2Outline2" fill="currentColor"><path d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm1 15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/><path d="m14.7 8.39l-3.78 5l-1.63-2.11a1 1 0 0 0-1.58 1.23l2.43 3.11a1 1 0 0 0 .79.38a1 1 0 0 0 .79-.39l4.57-6a1 1 0 1 0-1.6-1.22Z"/></g></g></g>`
	chevronDownFillInnerSVG               = `<g id="evaChevronDownFill0"><g id="evaChevronDownFill1"><path id="evaChevronDownFill2" fill="currentColor" d="M12 15.5a1 1 0 0 1-.71-.29l-4-4a1 1 0 1 1 1.42-1.42L12 13.1l3.3-3.18a1 1 0 1 1 1.38 1.44l-4 3.86a1 1 0 0 1-.68.28Z"/></g></g>`
	chevronDownOutlineInnerSVG            = `<g id="evaChevronDownOutline0"><g id="evaChevronDownOutline1"><path id="evaChevronDownOutline2" fill="currentColor" d="M12 15.5a1 1 0 0 1-.71-.29l-4-4a1 1 0 1 1 1.42-1.42L12 13.1l3.3-3.18a1 1 0 1 1 1.38 1.44l-4 3.86a1 1 0 0 1-.68.28Z"/></g></g>`
	chevronLeftFillInnerSVG               = `<g id="evaChevronLeftFill0"><g id="evaChevronLeftFill1"><path id="evaChevronLeftFill2" fill="currentColor" d="M13.36 17a1 1 0 0 1-.72-.31l-3.86-4a1 1 0 0 1 0-1.4l4-4a1 1 0 1 1 1.42 1.42L10.9 12l3.18 3.3a1 1 0 0 1 0 1.41a1 1 0 0 1-.72.29Z"/></g></g>`
	chevronLeftOutlineInnerSVG            = `<g id="evaChevronLeftOutline0"><g id="evaChevronLeftOutline1"><path id="evaChevronLeftOutline2" fill="currentColor" d="M13.36 17a1 1 0 0 1-.72-.31l-3.86-4a1 1 0 0 1 0-1.4l4-4a1 1 0 1 1 1.42 1.42L10.9 12l3.18 3.3a1 1 0 0 1 0 1.41a1 1 0 0 1-.72.29Z"/></g></g>`
	chevronRightFillInnerSVG              = `<g id="evaChevronRightFill0"><g id="evaChevronRightFill1"><path id="evaChevronRightFill2" fill="currentColor" d="M10.5 17a1 1 0 0 1-.71-.29a1 1 0 0 1 0-1.42L13.1 12L9.92 8.69a1 1 0 0 1 0-1.41a1 1 0 0 1 1.42 0l3.86 4a1 1 0 0 1 0 1.4l-4 4a1 1 0 0 1-.7.32Z"/></g></g>`
	chevronRightOutlineInnerSVG           = `<g id="evaChevronRightOutline0"><g id="evaChevronRightOutline1"><path id="evaChevronRightOutline2" fill="currentColor" d="M10.5 17a1 1 0 0 1-.71-.29a1 1 0 0 1 0-1.42L13.1 12L9.92 8.69a1 1 0 0 1 0-1.41a1 1 0 0 1 1.42 0l3.86 4a1 1 0 0 1 0 1.4l-4 4a1 1 0 0 1-.7.32Z"/></g></g>`
	chevronUpFillInnerSVG                 = `<g id="evaChevronUpFill0"><g id="evaChevronUpFill1"><path id="evaChevronUpFill2" fill="currentColor" d="M16 14.5a1 1 0 0 1-.71-.29L12 10.9l-3.3 3.18a1 1 0 0 1-1.41 0a1 1 0 0 1 0-1.42l4-3.86a1 1 0 0 1 1.4 0l4 4a1 1 0 0 1 0 1.42a1 1 0 0 1-.69.28Z"/></g></g>`
	chevronUpOutlineInnerSVG              = `<g id="evaChevronUpOutline0"><g id="evaChevronUpOutline1"><path id="evaChevronUpOutline2" fill="currentColor" d="M16 14.5a1 1 0 0 1-.71-.29L12 10.9l-3.3 3.18a1 1 0 0 1-1.41 0a1 1 0 0 1 0-1.42l4-3.86a1 1 0 0 1 1.4 0l4 4a1 1 0 0 1 0 1.42a1 1 0 0 1-.69.28Z"/></g></g>`
	clipboardFillInnerSVG                 = `<g id="evaClipboardFill0"><g id="evaClipboardFill1"><g id="evaClipboardFill2" fill="currentColor"><path d="M18 4v3a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><rect width="10" height="6" x="7" y="2" rx="1" ry="1"/></g></g></g>`
	clipboardOutlineInnerSVG              = `<g id="evaClipboardOutline0"><g id="evaClipboardOutline1"><path id="evaClipboardOutline2" fill="currentColor" d="M18 5V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v1a3 3 0 0 0-3 3v11a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3ZM8 4h8v4H8V4Zm11 15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1v1a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7a1 1 0 0 1 1 1Z"/></g></g>`
	clockFillInnerSVG                     = `<g id="evaClockFill0"><g id="evaClockFill1"><path id="evaClockFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm4 11h-4a1 1 0 0 1-1-1V8a1 1 0 0 1 2 0v3h3a1 1 0 0 1 0 2Z"/></g></g>`
	clockOutlineInnerSVG                  = `<g id="evaClockOutline0"><g id="evaClockOutline1"><g id="evaClockOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M16 11h-3V8a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1h4a1 1 0 0 0 0-2Z"/></g></g></g>`
	closeCircleFillInnerSVG               = `<g id="evaCloseCircleFill0"><g id="evaCloseCircleFill1"><path id="evaCloseCircleFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm2.71 11.29a1 1 0 0 1 0 1.42a1 1 0 0 1-1.42 0L12 13.41l-1.29 1.3a1 1 0 0 1-1.42 0a1 1 0 0 1 0-1.42l1.3-1.29l-1.3-1.29a1 1 0 0 1 1.42-1.42l1.29 1.3l1.29-1.3a1 1 0 0 1 1.42 1.42L13.41 12Z"/></g></g>`
	closeCircleOutlineInnerSVG            = `<g id="evaCloseCircleOutline0"><g id="evaCloseCircleOutline1"><g id="evaCloseCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M14.71 9.29a1 1 0 0 0-1.42 0L12 10.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 12l1.3-1.29a1 1 0 0 0 0-1.42Z"/></g></g></g>`
	closeFillInnerSVG                     = `<g id="evaCloseFill0"><g id="evaCloseFill1"><path id="evaCloseFill2" fill="currentColor" d="m13.41 12l4.3-4.29a1 1 0 1 0-1.42-1.42L12 10.59l-4.29-4.3a1 1 0 0 0-1.42 1.42l4.3 4.29l-4.3 4.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4.29-4.3l4.29 4.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	closeOutlineInnerSVG                  = `<g id="evaCloseOutline0"><g id="evaCloseOutline1"><path id="evaCloseOutline2" fill="currentColor" d="m13.41 12l4.3-4.29a1 1 0 1 0-1.42-1.42L12 10.59l-4.29-4.3a1 1 0 0 0-1.42 1.42l4.3 4.29l-4.3 4.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4.29-4.3l4.29 4.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	closeSquareFillInnerSVG               = `<g id="evaCloseSquareFill0"><g id="evaCloseSquareFill1"><path id="evaCloseSquareFill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm-3.29 10.29a1 1 0 0 1 0 1.42a1 1 0 0 1-1.42 0L12 13.41l-1.29 1.3a1 1 0 0 1-1.42 0a1 1 0 0 1 0-1.42l1.3-1.29l-1.3-1.29a1 1 0 0 1 1.42-1.42l1.29 1.3l1.29-1.3a1 1 0 0 1 1.42 1.42L13.41 12Z"/></g></g>`
	closeSquareOutlineInnerSVG            = `<g id="evaCloseSquareOutline0"><g id="evaCloseSquareOutline1"><g id="evaCloseSquareOutline2" fill="currentColor"><path d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm1 15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/><path d="M14.71 9.29a1 1 0 0 0-1.42 0L12 10.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 12l1.3-1.29a1 1 0 0 0 0-1.42Z"/></g></g></g>`
	cloudDownloadFillInnerSVG             = `<defs><path id="evaCloudDownloadFill0" d="M21.9 11c0-.11-.06-.22-.09-.33a4.17 4.17 0 0 0-.18-.57c-.05-.12-.12-.24-.18-.37s-.15-.3-.24-.44S21 9.08 21 9s-.2-.25-.31-.37s-.21-.2-.32-.3L20 8l-.36-.24a3.68 3.68 0 0 0-.44-.23l-.39-.18a4.13 4.13 0 0 0-.5-.15a3 3 0 0 0-.41-.09h-.18A6 6 0 0 0 6.33 7h-.18a3 3 0 0 0-.41.09a4.13 4.13 0 0 0-.5.15l-.39.18a3.68 3.68 0 0 0-.44.23L4.05 8l-.37.31c-.11.1-.22.19-.32.3s-.21.25-.31.37s-.18.23-.26.36s-.16.29-.24.44s-.13.25-.18.37a4.17 4.17 0 0 0-.18.57c0 .11-.07.22-.09.33A5.23 5.23 0 0 0 2 12a5.5 5.5 0 0 0 .09.91c0 .1.05.19.07.29a5.58 5.58 0 0 0 .18.58l.12.29a5 5 0 0 0 .3.56l.14.22a.56.56 0 0 0 .05.08L3 15a5 5 0 0 0 4 2a2 2 0 0 1 .59-1.41A2 2 0 0 1 9 15a1.92 1.92 0 0 1 1 .27V12a2 2 0 0 1 4 0v3.37a2 2 0 0 1 1-.27a2.05 2.05 0 0 1 1.44.61A2 2 0 0 1 17 17a5 5 0 0 0 4-2l.05-.05a.56.56 0 0 0 .05-.08l.14-.22a5 5 0 0 0 .3-.56l.12-.29a5.58 5.58 0 0 0 .18-.58c0-.1.05-.19.07-.29A5.5 5.5 0 0 0 22 12a5.23 5.23 0 0 0-.1-1Z"/><path id="evaCloudDownloadFill1" d="M14.31 16.38L13 17.64V12a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 21a1 1 0 0 0 .69-.28l3-2.9a1 1 0 1 0-1.38-1.44Z"/></defs><g id="evaCloudDownloadFill2"><g id="evaCloudDownloadFill3"><g id="evaCloudDownloadFill4" fill="currentColor"><use href="#evaCloudDownloadFill0"/><use href="#evaCloudDownloadFill1"/><use href="#evaCloudDownloadFill0"/><use href="#evaCloudDownloadFill1"/></g></g></g>`
	cloudDownloadOutlineInnerSVG          = `<g id="evaCloudDownloadOutline0"><g id="evaCloudDownloadOutline1"><g id="evaCloudDownloadOutline2" fill="currentColor"><path d="M14.31 16.38L13 17.64V12a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 21a1 1 0 0 0 .69-.28l3-2.9a1 1 0 1 0-1.38-1.44Z"/><path d="M17.67 7A6 6 0 0 0 6.33 7a5 5 0 0 0-3.08 8.27A1 1 0 1 0 4.75 14A3 3 0 0 1 7 9h.1a1 1 0 0 0 1-.8a4 4 0 0 1 7.84 0a1 1 0 0 0 1 .8H17a3 3 0 0 1 2.25 5a1 1 0 0 0 .09 1.42a1 1 0 0 0 .66.25a1 1 0 0 0 .75-.34A5 5 0 0 0 17.67 7Z"/></g></g></g>`
	cloudUploadFillInnerSVG               = `<g id="evaCloudUploadFill0"><g id="evaCloudUploadFill1"><g id="evaCloudUploadFill2" fill="currentColor"><path d="M21.9 12c0-.11-.06-.22-.09-.33a4.17 4.17 0 0 0-.18-.57c-.05-.12-.12-.24-.18-.37s-.15-.3-.24-.44S21 10.08 21 10s-.2-.25-.31-.37s-.21-.2-.32-.3L20 9l-.36-.24a3.68 3.68 0 0 0-.44-.23l-.39-.18a4.13 4.13 0 0 0-.5-.15a3 3 0 0 0-.41-.09L17.67 8A6 6 0 0 0 6.33 8l-.18.05a3 3 0 0 0-.41.09a4.13 4.13 0 0 0-.5.15l-.39.18a3.68 3.68 0 0 0-.44.23l-.36.3l-.37.31c-.11.1-.22.19-.32.3s-.21.25-.31.37s-.18.23-.26.36s-.16.29-.24.44s-.13.25-.18.37a4.17 4.17 0 0 0-.18.57c0 .11-.07.22-.09.33A5.23 5.23 0 0 0 2 13a5.5 5.5 0 0 0 .09.91c0 .1.05.19.07.29a5.58 5.58 0 0 0 .18.58l.12.29a5 5 0 0 0 .3.56l.14.22a.56.56 0 0 0 .05.08L3 16a5 5 0 0 0 4 2h3v-1.37a2 2 0 0 1-1 .27a2.05 2.05 0 0 1-1.44-.61a2 2 0 0 1 .05-2.83l3-2.9A2 2 0 0 1 12 10a2 2 0 0 1 1.41.59l3 3a2 2 0 0 1 0 2.82A2 2 0 0 1 15 17a1.92 1.92 0 0 1-1-.27V18h3a5 5 0 0 0 4-2l.05-.05a.56.56 0 0 0 .05-.08l.14-.22a5 5 0 0 0 .3-.56l.12-.29a5.58 5.58 0 0 0 .18-.58c0-.1.05-.19.07-.29A5.5 5.5 0 0 0 22 13a5.23 5.23 0 0 0-.1-1Z"/><path d="M12.71 11.29a1 1 0 0 0-1.4 0l-3 2.9a1 1 0 1 0 1.38 1.44L11 14.36V20a1 1 0 0 0 2 0v-5.59l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g></g>`
	cloudUploadOutlineInnerSVG            = `<g id="evaCloudUploadOutline0"><g id="evaCloudUploadOutline1"><g id="evaCloudUploadOutline2" fill="currentColor"><path d="M12.71 11.29a1 1 0 0 0-1.4 0l-3 2.9a1 1 0 1 0 1.38 1.44L11 14.36V20a1 1 0 0 0 2 0v-5.59l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/><path d="M17.67 7A6 6 0 0 0 6.33 7a5 5 0 0 0-3.08 8.27A1 1 0 1 0 4.75 14A3 3 0 0 1 7 9h.1a1 1 0 0 0 1-.8a4 4 0 0 1 7.84 0a1 1 0 0 0 1 .8H17a3 3 0 0 1 2.25 5a1 1 0 0 0 .09 1.42a1 1 0 0 0 .66.25a1 1 0 0 0 .75-.34A5 5 0 0 0 17.67 7Z"/></g></g></g>`
	codeDownloadFillInnerSVG              = `<g id="evaCodeDownloadFill0"><g id="evaCodeDownloadFill1"><g id="evaCodeDownloadFill2" fill="currentColor"><path d="m4.29 12l4.48-5.36a1 1 0 1 0-1.54-1.28l-5 6a1 1 0 0 0 0 1.27l4.83 6a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.63Zm17.49-.63l-4.78-6a1 1 0 0 0-1.56 1.26L19.71 12l-4.48 5.37a1 1 0 0 0 .13 1.41A1 1 0 0 0 16 19a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 .01-1.27Z"/><path d="M15.72 11.41a1 1 0 0 0-1.41 0L13 12.64V8a1 1 0 0 0-2 0v4.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 16a1 1 0 0 0 .69-.28l3-2.9a1 1 0 0 0 .03-1.41Z"/></g></g></g>`
	codeDownloadOutlineInnerSVG           = `<g id="evaCodeDownloadOutline0"><g id="evaCodeDownloadOutline1"><g id="evaCodeDownloadOutline2" fill="currentColor"><path d="m4.29 12l4.48-5.36a1 1 0 1 0-1.54-1.28l-5 6a1 1 0 0 0 0 1.27l4.83 6a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.63Zm17.49-.63l-4.78-6a1 1 0 0 0-1.56 1.26L19.71 12l-4.48 5.37a1 1 0 0 0 .13 1.41A1 1 0 0 0 16 19a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 .01-1.27Z"/><path d="M15.72 11.41a1 1 0 0 0-1.41 0L13 12.64V8a1 1 0 0 0-2 0v4.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 16a1 1 0 0 0 .69-.28l3-2.9a1 1 0 0 0 .03-1.41Z"/></g></g></g>`
	codeFillInnerSVG                      = `<g id="evaCodeFill0"><g id="evaCodeFill1"><path id="evaCodeFill2" fill="currentColor" d="M8.64 5.23a1 1 0 0 0-1.41.13l-5 6a1 1 0 0 0 0 1.27l4.83 6a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.63L4.29 12l4.48-5.36a1 1 0 0 0-.13-1.41Zm13.14 6.14l-4.78-6a1 1 0 0 0-1.41-.15a1 1 0 0 0-.15 1.41L19.71 12l-4.48 5.37a1 1 0 0 0 .13 1.41A1 1 0 0 0 16 19a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 .01-1.27Z"/></g></g>`
	codeOutlineInnerSVG                   = `<g id="evaCodeOutline0"><g id="evaCodeOutline1"><path id="evaCodeOutline2" fill="currentColor" d="M8.64 5.23a1 1 0 0 0-1.41.13l-5 6a1 1 0 0 0 0 1.27l4.83 6a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.63L4.29 12l4.48-5.36a1 1 0 0 0-.13-1.41Zm13.14 6.14l-4.78-6a1 1 0 0 0-1.41-.15a1 1 0 0 0-.15 1.41L19.71 12l-4.48 5.37a1 1 0 0 0 .13 1.41A1 1 0 0 0 16 19a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 .01-1.27Z"/></g></g>`
	collapseFillInnerSVG                  = `<g id="evaCollapseFill0"><g id="evaCollapseFill1"><path id="evaCollapseFill2" fill="currentColor" d="M19 9h-2.58l3.29-3.29a1 1 0 1 0-1.42-1.42L15 7.57V5a1 1 0 0 0-1-1a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2Zm-9 4H5a1 1 0 0 0 0 2h2.57l-3.28 3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L9 16.42V19a1 1 0 0 0 1 1a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1Z"/></g></g>`
	collapseOutlineInnerSVG               = `<g id="evaCollapseOutline0"><g id="evaCollapseOutline1"><path id="evaCollapseOutline2" fill="currentColor" d="M19 9h-2.58l3.29-3.29a1 1 0 1 0-1.42-1.42L15 7.57V5a1 1 0 0 0-1-1a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2Zm-9 4H5a1 1 0 0 0 0 2h2.57l-3.28 3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L9 16.42V19a1 1 0 0 0 1 1a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1Z"/></g></g>`
	colorPaletteFillInnerSVG              = `<g id="evaColorPaletteFill0"><g id="evaColorPaletteFill1"><path id="evaColorPaletteFill2" fill="currentColor" d="M19.54 5.08A10.61 10.61 0 0 0 11.91 2a10 10 0 0 0-.05 20a2.58 2.58 0 0 0 2.53-1.89a2.52 2.52 0 0 0-.57-2.28a.5.5 0 0 1 .37-.83h1.65A6.15 6.15 0 0 0 22 11.33a8.48 8.48 0 0 0-2.46-6.25Zm-12.7 9.66a1.5 1.5 0 1 1 .4-2.08a1.49 1.49 0 0 1-.4 2.08ZM8.3 9.25a1.5 1.5 0 1 1-.55-2a1.5 1.5 0 0 1 .55 2ZM11 7a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 11 7Zm5.75.8a1.5 1.5 0 1 1 .55-2a1.5 1.5 0 0 1-.55 2Z"/></g></g>`
	colorPaletteOutlineInnerSVG           = `<g id="evaColorPaletteOutline0"><g id="evaColorPaletteOutline1"><g id="evaColorPaletteOutline2" fill="currentColor"><path d="M19.54 5.08A10.61 10.61 0 0 0 11.91 2a10 10 0 0 0-.05 20a2.58 2.58 0 0 0 2.53-1.89a2.52 2.52 0 0 0-.57-2.28a.5.5 0 0 1 .37-.83h1.65A6.15 6.15 0 0 0 22 11.33a8.48 8.48 0 0 0-2.46-6.25ZM15.88 15h-1.65a2.49 2.49 0 0 0-1.87 4.15a.49.49 0 0 1 .12.49c-.05.21-.28.34-.59.36a8 8 0 0 1-7.82-9.11A8.1 8.1 0 0 1 11.92 4H12a8.47 8.47 0 0 1 6.1 2.48a6.5 6.5 0 0 1 1.9 4.77A4.17 4.17 0 0 1 15.88 15Z"/><circle cx="12" cy="6.5" r="1.5"/><path d="M15.25 7.2a1.5 1.5 0 1 0 2.05.55a1.5 1.5 0 0 0-2.05-.55Zm-6.5 0a1.5 1.5 0 1 0 .55 2.05a1.5 1.5 0 0 0-.55-2.05Zm-2.59 4.06a1.5 1.5 0 1 0 2.08.4a1.49 1.49 0 0 0-2.08-.4Z"/></g></g></g>`
	colorPickerFillInnerSVG               = `<g id="evaColorPickerFill0"><g id="evaColorPickerFill1"><path id="evaColorPickerFill2" fill="currentColor" d="M19.4 7.34L16.66 4.6A1.92 1.92 0 0 0 14 4.53l-2 2l-1.29-1.24a1 1 0 0 0-1.42 1.42L10.53 8L5 13.53a2 2 0 0 0-.57 1.21L4 18.91a1 1 0 0 0 .29.8A1 1 0 0 0 5 20h.09l4.17-.38a2 2 0 0 0 1.21-.57l5.58-5.58l1.24 1.24a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-1.24-1.24l2-2a1.92 1.92 0 0 0-.07-2.71Zm-13 7.6L12 9.36l2.69 2.7l-2.79 2.79"/></g></g>`
	colorPickerOutlineInnerSVG            = `<g id="evaColorPickerOutline0"><g id="evaColorPickerOutline1"><path id="evaColorPickerOutline2" fill="currentColor" d="M19.4 7.34L16.66 4.6A1.92 1.92 0 0 0 14 4.53l-2 2l-1.29-1.24a1 1 0 0 0-1.42 1.42L10.53 8L5 13.53a2 2 0 0 0-.57 1.21L4 18.91a1 1 0 0 0 .29.8A1 1 0 0 0 5 20h.09l4.17-.38a2 2 0 0 0 1.21-.57l5.58-5.58l1.24 1.24a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-1.24-1.24l2-2a1.92 1.92 0 0 0-.07-2.71ZM9.08 17.62l-3 .28l.27-3L12 9.36l2.69 2.7Zm7-7L13.36 8l1.91-2L18 8.73Z"/></g></g>`
	compassFillInnerSVG                   = `<g id="evaCompassFill0"><g id="evaCompassFill1"><g id="evaCompassFill2" fill="currentColor"><path d="m10.8 13.21l1.69-.68l.71-1.74l-1.69.68l-.71 1.74z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm3.93 7.42l-1.75 4.26a1 1 0 0 1-.55.55l-4.21 1.7A1 1 0 0 1 9 16a1 1 0 0 1-.71-.31h-.05a1 1 0 0 1-.18-1l1.75-4.26a1 1 0 0 1 .55-.55l4.21-1.7a1 1 0 0 1 1.1.25a1 1 0 0 1 .26.99Z"/></g></g></g>`
	compassOutlineInnerSVG                = `<g id="evaCompassOutline0"><g id="evaCompassOutline1"><g id="evaCompassOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M15.68 8.32a1 1 0 0 0-1.1-.25l-4.21 1.7a1 1 0 0 0-.55.55l-1.75 4.26a1 1 0 0 0 .18 1h.05A1 1 0 0 0 9 16a1 1 0 0 0 .38-.07l4.21-1.7a1 1 0 0 0 .55-.55l1.75-4.26a1 1 0 0 0-.21-1.1Zm-4.88 4.89l.71-1.74l1.69-.68l-.71 1.74Z"/></g></g></g>`
	copyFillInnerSVG                      = `<g id="evaCopyFill0"><g id="evaCopyFill1"><path id="evaCopyFill2" fill="currentColor" d="M18 9h-3V5.67A2.68 2.68 0 0 0 12.33 3H5.67A2.68 2.68 0 0 0 3 5.67v6.66A2.68 2.68 0 0 0 5.67 15H9v3a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3Zm-9 3v1H5.67a.67.67 0 0 1-.67-.67V5.67A.67.67 0 0 1 5.67 5h6.66a.67.67 0 0 1 .67.67V9h-1a3 3 0 0 0-3 3Z"/></g></g>`
	copyOutlineInnerSVG                   = `<g id="evaCopyOutline0"><g id="evaCopyOutline1"><g id="evaCopyOutline2" fill="currentColor"><path d="M18 21h-6a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3Zm-6-10a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1Z"/><path d="M9.73 15H5.67A2.68 2.68 0 0 1 3 12.33V5.67A2.68 2.68 0 0 1 5.67 3h6.66A2.68 2.68 0 0 1 15 5.67V9.4h-2V5.67a.67.67 0 0 0-.67-.67H5.67a.67.67 0 0 0-.67.67v6.66a.67.67 0 0 0 .67.67h4.06Z"/></g></g></g>`
	cornerDownLeftFillInnerSVG            = `<g id="evaCornerDownLeftFill0"><g id="evaCornerDownLeftFill1"><path id="evaCornerDownLeftFill2" fill="currentColor" d="M20 6a1 1 0 0 0-1-1a1 1 0 0 0-1 1v5a1 1 0 0 1-.29.71A1 1 0 0 1 17 12H8.08l2.69-3.39a1 1 0 0 0-1.52-1.17l-4 5a1 1 0 0 0 0 1.25l4 5a1 1 0 0 0 .78.37a1 1 0 0 0 .62-.22a1 1 0 0 0 .15-1.41l-2.66-3.36h8.92a3 3 0 0 0 3-3Z"/></g></g>`
	cornerDownLeftOutlineInnerSVG         = `<g id="evaCornerDownLeftOutline0"><g id="evaCornerDownLeftOutline1"><path id="evaCornerDownLeftOutline2" fill="currentColor" d="M20 6a1 1 0 0 0-1-1a1 1 0 0 0-1 1v5a1 1 0 0 1-.29.71A1 1 0 0 1 17 12H8.08l2.69-3.39a1 1 0 0 0-1.52-1.17l-4 5a1 1 0 0 0 0 1.25l4 5a1 1 0 0 0 .78.37a1 1 0 0 0 .62-.22a1 1 0 0 0 .15-1.41l-2.66-3.36h8.92a3 3 0 0 0 3-3Z"/></g></g>`
	cornerDownRightFillInnerSVG           = `<g id="evaCornerDownRightFill0"><g id="evaCornerDownRightFill1"><path id="evaCornerDownRightFill2" fill="currentColor" d="m19.78 12.38l-4-5a1 1 0 0 0-1.56 1.24l2.7 3.38H8a1 1 0 0 1-1-1V6a1 1 0 0 0-2 0v5a3 3 0 0 0 3 3h8.92l-2.7 3.38a1 1 0 0 0 .16 1.4A1 1 0 0 0 15 19a1 1 0 0 0 .78-.38l4-5a1 1 0 0 0 0-1.24Z"/></g></g>`
	cornerDownRightOutlineInnerSVG        = `<g id="evaCornerDownRightOutline0"><g id="evaCornerDownRightOutline1"><path id="evaCornerDownRightOutline2" fill="currentColor" d="m19.78 12.38l-4-5a1 1 0 0 0-1.56 1.24l2.7 3.38H8a1 1 0 0 1-1-1V6a1 1 0 0 0-2 0v5a3 3 0 0 0 3 3h8.92l-2.7 3.38a1 1 0 0 0 .16 1.4A1 1 0 0 0 15 19a1 1 0 0 0 .78-.38l4-5a1 1 0 0 0 0-1.24Z"/></g></g>`
	cornerLeftDownFillInnerSVG            = `<g id="evaCornerLeftDownFill0"><g id="evaCornerLeftDownFill1"><path id="evaCornerLeftDownFill2" fill="currentColor" d="M18 5h-5a3 3 0 0 0-3 3v8.92l-3.38-2.7a1 1 0 0 0-1.24 1.56l5 4a1 1 0 0 0 1.24 0l5-4a1 1 0 1 0-1.24-1.56L12 16.92V8a1 1 0 0 1 1-1h5a1 1 0 0 0 0-2Z"/></g></g>`
	cornerLeftDownOutlineInnerSVG         = `<g id="evaCornerLeftDownOutline0"><g id="evaCornerLeftDownOutline1"><path id="evaCornerLeftDownOutline2" fill="currentColor" d="M18 5h-5a3 3 0 0 0-3 3v8.92l-3.38-2.7a1 1 0 0 0-1.24 1.56l5 4a1 1 0 0 0 1.24 0l5-4a1 1 0 1 0-1.24-1.56L12 16.92V8a1 1 0 0 1 1-1h5a1 1 0 0 0 0-2Z"/></g></g>`
	cornerLeftUpFillInnerSVG              = `<g id="evaCornerLeftUpFill0"><g id="evaCornerLeftUpFill1"><path id="evaCornerLeftUpFill2" fill="currentColor" d="M18 17h-5a1 1 0 0 1-1-1V7.08l3.38 2.7A1 1 0 0 0 16 10a1 1 0 0 0 .78-.38a1 1 0 0 0-.16-1.4l-5-4a1 1 0 0 0-1.24 0l-5 4a1 1 0 0 0 1.24 1.56L10 7.08V16a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2Z"/></g></g>`
	cornerLeftUpOutlineInnerSVG           = `<g id="evaCornerLeftUpOutline0"><g id="evaCornerLeftUpOutline1"><path id="evaCornerLeftUpOutline2" fill="currentColor" d="M18 17h-5a1 1 0 0 1-1-1V7.08l3.38 2.7A1 1 0 0 0 16 10a1 1 0 0 0 .78-.38a1 1 0 0 0-.16-1.4l-5-4a1 1 0 0 0-1.24 0l-5 4a1 1 0 0 0 1.24 1.56L10 7.08V16a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2Z"/></g></g>`
	cornerRightDownFillInnerSVG           = `<g id="evaCornerRightDownFill0"><g id="evaCornerRightDownFill1"><path id="evaCornerRightDownFill2" fill="currentColor" d="M18.78 14.38a1 1 0 0 0-1.4-.16L14 16.92V8a3 3 0 0 0-3-3H6a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v8.92l-3.38-2.7a1 1 0 0 0-1.24 1.56l5 4a1 1 0 0 0 1.24 0l5-4a1 1 0 0 0 .16-1.4Z"/></g></g>`
	cornerRightDownOutlineInnerSVG        = `<g id="evaCornerRightDownOutline0"><g id="evaCornerRightDownOutline1"><path id="evaCornerRightDownOutline2" fill="currentColor" d="M18.78 14.38a1 1 0 0 0-1.4-.16L14 16.92V8a3 3 0 0 0-3-3H6a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v8.92l-3.38-2.7a1 1 0 0 0-1.24 1.56l5 4a1 1 0 0 0 1.24 0l5-4a1 1 0 0 0 .16-1.4Z"/></g></g>`
	cornerRightUpFillInnerSVG             = `<g id="evaCornerRightUpFill0"><g id="evaCornerRightUpFill1"><path id="evaCornerRightUpFill2" fill="currentColor" d="m18.62 8.22l-5-4a1 1 0 0 0-1.24 0l-5 4a1 1 0 0 0 1.24 1.56L12 7.08V16a1 1 0 0 1-1 1H6a1 1 0 0 0 0 2h5a3 3 0 0 0 3-3V7.08l3.38 2.7A1 1 0 0 0 18 10a1 1 0 0 0 .78-.38a1 1 0 0 0-.16-1.4Z"/></g></g>`
	cornerRightUpOutlineInnerSVG          = `<g id="evaCornerRightUpOutline0"><g id="evaCornerRightUpOutline1"><path id="evaCornerRightUpOutline2" fill="currentColor" d="m18.62 8.22l-5-4a1 1 0 0 0-1.24 0l-5 4a1 1 0 0 0 1.24 1.56L12 7.08V16a1 1 0 0 1-1 1H6a1 1 0 0 0 0 2h5a3 3 0 0 0 3-3V7.08l3.38 2.7A1 1 0 0 0 18 10a1 1 0 0 0 .78-.38a1 1 0 0 0-.16-1.4Z"/></g></g>`
	cornerUpLeftFillInnerSVG              = `<g id="evaCornerUpLeftFill0"><g id="evaCornerUpLeftFill1"><path id="evaCornerUpLeftFill2" fill="currentColor" d="M16 10H7.08l2.7-3.38a1 1 0 1 0-1.56-1.24l-4 5a1 1 0 0 0 0 1.24l4 5A1 1 0 0 0 9 17a1 1 0 0 0 .62-.22a1 1 0 0 0 .16-1.4L7.08 12H16a1 1 0 0 1 1 1v5a1 1 0 0 0 2 0v-5a3 3 0 0 0-3-3Z"/></g></g>`
	cornerUpLeftOutlineInnerSVG           = `<g id="evaCornerUpLeftOutline0"><g id="evaCornerUpLeftOutline1"><path id="evaCornerUpLeftOutline2" fill="currentColor" d="M16 10H7.08l2.7-3.38a1 1 0 1 0-1.56-1.24l-4 5a1 1 0 0 0 0 1.24l4 5A1 1 0 0 0 9 17a1 1 0 0 0 .62-.22a1 1 0 0 0 .16-1.4L7.08 12H16a1 1 0 0 1 1 1v5a1 1 0 0 0 2 0v-5a3 3 0 0 0-3-3Z"/></g></g>`
	cornerUpRightFillInnerSVG             = `<g id="evaCornerUpRightFill0"><g id="evaCornerUpRightFill1"><path id="evaCornerUpRightFill2" fill="currentColor" d="m19.78 10.38l-4-5a1 1 0 0 0-1.56 1.24l2.7 3.38H8a3 3 0 0 0-3 3v5a1 1 0 0 0 2 0v-5a1 1 0 0 1 1-1h8.92l-2.7 3.38a1 1 0 0 0 .16 1.4A1 1 0 0 0 15 17a1 1 0 0 0 .78-.38l4-5a1 1 0 0 0 0-1.24Z"/></g></g>`
	cornerUpRightOutlineInnerSVG          = `<g id="evaCornerUpRightOutline0"><g id="evaCornerUpRightOutline1"><path id="evaCornerUpRightOutline2" fill="currentColor" d="m19.78 10.38l-4-5a1 1 0 0 0-1.56 1.24l2.7 3.38H8a3 3 0 0 0-3 3v5a1 1 0 0 0 2 0v-5a1 1 0 0 1 1-1h8.92l-2.7 3.38a1 1 0 0 0 .16 1.4A1 1 0 0 0 15 17a1 1 0 0 0 .78-.38l4-5a1 1 0 0 0 0-1.24Z"/></g></g>`
	creditCardFillInnerSVG                = `<g id="evaCreditCardFill0"><g id="evaCreditCardFill1"><path id="evaCreditCardFill2" fill="currentColor" d="M19 5H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3Zm-8 10H7a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Zm6 0h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2Zm3-6H4V8a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/></g></g>`
	creditCardOutlineInnerSVG             = `<g id="evaCreditCardOutline0"><g id="evaCreditCardOutline1"><g id="evaCreditCardOutline2" fill="currentColor"><path d="M19 5H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3ZM4 8a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v1H4Zm16 8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-5h16Z"/><path d="M7 15h4a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2Zm8 0h2a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2Z"/></g></g></g>`
	cropFillInnerSVG                      = `<g id="evaCropFill0"><g id="evaCropFill1"><path id="evaCropFill2" fill="currentColor" d="M21 16h-3V8.56A2.56 2.56 0 0 0 15.44 6H8V3a1 1 0 0 0-2 0v3H3a1 1 0 0 0 0 2h3v7.44A2.56 2.56 0 0 0 8.56 18H16v3a1 1 0 0 0 2 0v-3h3a1 1 0 0 0 0-2ZM8.56 16a.56.56 0 0 1-.56-.56V8h7.44a.56.56 0 0 1 .56.56V16Z"/></g></g>`
	cropOutlineInnerSVG                   = `<g id="evaCropOutline0"><g id="evaCropOutline1"><path id="evaCropOutline2" fill="currentColor" d="M21 16h-3V8.56A2.56 2.56 0 0 0 15.44 6H8V3a1 1 0 0 0-2 0v3H3a1 1 0 0 0 0 2h3v7.44A2.56 2.56 0 0 0 8.56 18H16v3a1 1 0 0 0 2 0v-3h3a1 1 0 0 0 0-2ZM8.56 16a.56.56 0 0 1-.56-.56V8h7.44a.56.56 0 0 1 .56.56V16Z"/></g></g>`
	cubeFillInnerSVG                      = `<g id="evaCubeFill0"><g id="evaCubeFill1"><path id="evaCubeFill2" fill="currentColor" d="M11.25 11.83L3 8.36v7.73a1.69 1.69 0 0 0 1 1.52L11.19 21h.06ZM12 10.5l8.51-3.57a1.62 1.62 0 0 0-.51-.38l-7.2-3.37a1.87 1.87 0 0 0-1.6 0L4 6.55a1.62 1.62 0 0 0-.51.38Zm.75 1.33V21h.05l7.2-3.39a1.69 1.69 0 0 0 1-1.51V8.36Z"/></g></g>`
	cubeOutlineInnerSVG                   = `<g id="evaCubeOutline0"><g id="evaCubeOutline1"><path id="evaCubeOutline2" fill="currentColor" d="M20.66 7.26c0-.07-.1-.14-.15-.21l-.09-.1a2.5 2.5 0 0 0-.86-.68l-6.4-3a2.7 2.7 0 0 0-2.26 0l-6.4 3a2.6 2.6 0 0 0-.86.68L3.52 7a1 1 0 0 0-.15.2A2.39 2.39 0 0 0 3 8.46v7.06a2.49 2.49 0 0 0 1.46 2.26l6.4 3a2.7 2.7 0 0 0 2.27 0l6.4-3A2.49 2.49 0 0 0 21 15.54V8.46a2.39 2.39 0 0 0-.34-1.2Zm-8.95-2.2a.73.73 0 0 1 .58 0l5.33 2.48L12 10.15L6.38 7.54ZM5.3 16a.47.47 0 0 1-.3-.43V9.1l6 2.79v6.72Zm13.39 0L13 18.61v-6.72l6-2.79v6.44a.48.48 0 0 1-.31.46Z"/></g></g>`
	diagonalArrowLeftDownFillInnerSVG     = `<g id="evaDiagonalArrowLeftDownFill0"><g id="evaDiagonalArrowLeftDownFill1"><path id="evaDiagonalArrowLeftDownFill2" fill="currentColor" d="M17.71 6.29a1 1 0 0 0-1.42 0L8 14.59V9a1 1 0 0 0-2 0v8a1 1 0 0 0 1 1h8a1 1 0 0 0 0-2H9.41l8.3-8.29a1 1 0 0 0 0-1.42Z"/></g></g>`
	diagonalArrowLeftDownOutlineInnerSVG  = `<g id="evaDiagonalArrowLeftDownOutline0"><g id="evaDiagonalArrowLeftDownOutline1"><path id="evaDiagonalArrowLeftDownOutline2" fill="currentColor" d="M17.71 6.29a1 1 0 0 0-1.42 0L8 14.59V9a1 1 0 0 0-2 0v8a1 1 0 0 0 1 1h8a1 1 0 0 0 0-2H9.41l8.3-8.29a1 1 0 0 0 0-1.42Z"/></g></g>`
	diagonalArrowLeftUpFillInnerSVG       = `<g id="evaDiagonalArrowLeftUpFill0"><g id="evaDiagonalArrowLeftUpFill1"><path id="evaDiagonalArrowLeftUpFill2" fill="currentColor" d="M17.71 16.29L9.42 8H15a1 1 0 0 0 0-2H7.05a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1H7a1 1 0 0 0 1-1V9.45l8.26 8.26a1 1 0 0 0 1.42 0a1 1 0 0 0 .03-1.42Z"/></g></g>`
	diagonalArrowLeftUpOutlineInnerSVG    = `<g id="evaDiagonalArrowLeftUpOutline0"><g id="evaDiagonalArrowLeftUpOutline1"><path id="evaDiagonalArrowLeftUpOutline2" fill="currentColor" d="M17.71 16.29L9.42 8H15a1 1 0 0 0 0-2H7.05a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1H7a1 1 0 0 0 1-1V9.45l8.26 8.26a1 1 0 0 0 1.42 0a1 1 0 0 0 .03-1.42Z"/></g></g>`
	diagonalArrowRightDownFillInnerSVG    = `<g id="evaDiagonalArrowRightDownFill0"><g id="evaDiagonalArrowRightDownFill1"><path id="evaDiagonalArrowRightDownFill2" fill="currentColor" d="M17 8a1 1 0 0 0-1 1v5.59l-8.29-8.3a1 1 0 0 0-1.42 1.42l8.3 8.29H9a1 1 0 0 0 0 2h8a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1Z"/></g></g>`
	diagonalArrowRightDownOutlineInnerSVG = `<g id="evaDiagonalArrowRightDownOutline0"><g id="evaDiagonalArrowRightDownOutline1"><path id="evaDiagonalArrowRightDownOutline2" fill="currentColor" d="M17 8a1 1 0 0 0-1 1v5.59l-8.29-8.3a1 1 0 0 0-1.42 1.42l8.3 8.29H9a1 1 0 0 0 0 2h8a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1Z"/></g></g>`
	diagonalArrowRightUpFillInnerSVG      = `<g id="evaDiagonalArrowRightUpFill0"><g id="evaDiagonalArrowRightUpFill1"><path id="evaDiagonalArrowRightUpFill2" fill="currentColor" d="M18 7.05a1 1 0 0 0-1-1L9 6a1 1 0 0 0 0 2h5.56l-8.27 8.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L16 9.42V15a1 1 0 0 0 1 1a1 1 0 0 0 1-1Z"/></g></g>`
	diagonalArrowRightUpOutlineInnerSVG   = `<g id="evaDiagonalArrowRightUpOutline0"><g id="evaDiagonalArrowRightUpOutline1"><path id="evaDiagonalArrowRightUpOutline2" fill="currentColor" d="M18 7.05a1 1 0 0 0-1-1L9 6a1 1 0 0 0 0 2h5.56l-8.27 8.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L16 9.42V15a1 1 0 0 0 1 1a1 1 0 0 0 1-1Z"/></g></g>`
	doneAllFillInnerSVG                   = `<g id="evaDoneAllFill0"><g id="evaDoneAllFill1"><g id="evaDoneAllFill2" fill="currentColor"><path d="M16.62 6.21a1 1 0 0 0-1.41.17l-7 9l-3.43-4.18a1 1 0 1 0-1.56 1.25l4.17 5.18a1 1 0 0 0 .78.37a1 1 0 0 0 .83-.38l7.83-10a1 1 0 0 0-.21-1.41Zm5 0a1 1 0 0 0-1.41.17l-7 9l-.61-.75l-1.26 1.62l1.1 1.37a1 1 0 0 0 .78.37a1 1 0 0 0 .78-.38l7.83-10a1 1 0 0 0-.21-1.4Z"/><path d="M8.71 13.06L10 11.44l-.2-.24a1 1 0 0 0-1.43-.2a1 1 0 0 0-.15 1.41Z"/></g></g></g>`
	doneAllOutlineInnerSVG                = `<g id="evaDoneAllOutline0"><g id="evaDoneAllOutline1"><g id="evaDoneAllOutline2" fill="currentColor"><path d="M16.62 6.21a1 1 0 0 0-1.41.17l-7 9l-3.43-4.18a1 1 0 1 0-1.56 1.25l4.17 5.18a1 1 0 0 0 .78.37a1 1 0 0 0 .83-.38l7.83-10a1 1 0 0 0-.21-1.41Zm5 0a1 1 0 0 0-1.41.17l-7 9l-.61-.75l-1.26 1.62l1.1 1.37a1 1 0 0 0 .78.37a1 1 0 0 0 .78-.38l7.83-10a1 1 0 0 0-.21-1.4Z"/><path d="M8.71 13.06L10 11.44l-.2-.24a1 1 0 0 0-1.43-.2a1 1 0 0 0-.15 1.41Z"/></g></g></g>`
	downloadFillInnerSVG                  = `<g id="evaDownloadFill0"><g id="evaDownloadFill1"><g id="evaDownloadFill2" fill="currentColor"><rect width="16" height="2" x="4" y="18" rx="1" ry="1"/><rect width="4" height="2" x="3" y="17" rx="1" ry="1" transform="rotate(-90 5 18)"/><rect width="4" height="2" x="17" y="17" rx="1" ry="1" transform="rotate(-90 19 18)"/><path d="M12 15a1 1 0 0 1-.58-.18l-4-2.82a1 1 0 0 1-.24-1.39a1 1 0 0 1 1.4-.24L12 12.76l3.4-2.56a1 1 0 0 1 1.2 1.6l-4 3a1 1 0 0 1-.6.2Z"/><path d="M12 13a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1Z"/></g></g></g>`
	downloadOutlineInnerSVG               = `<g id="evaDownloadOutline0"><g id="evaDownloadOutline1"><g id="evaDownloadOutline2" fill="currentColor"><rect width="16" height="2" x="4" y="18" rx="1" ry="1"/><rect width="4" height="2" x="3" y="17" rx="1" ry="1" transform="rotate(-90 5 18)"/><rect width="4" height="2" x="17" y="17" rx="1" ry="1" transform="rotate(-90 19 18)"/><path d="M12 15a1 1 0 0 1-.58-.18l-4-2.82a1 1 0 0 1-.24-1.39a1 1 0 0 1 1.4-.24L12 12.76l3.4-2.56a1 1 0 0 1 1.2 1.6l-4 3a1 1 0 0 1-.6.2Z"/><path d="M12 13a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1Z"/></g></g></g>`
	dropletFillInnerSVG                   = `<g id="evaDropletFill0"><g id="evaDropletFill1"><path id="evaDropletFill2" fill="currentColor" d="M12 21.1a7.4 7.4 0 0 1-5.28-2.28a7.73 7.73 0 0 1 .1-10.77l4.64-4.65a.94.94 0 0 1 .71-.3a1 1 0 0 1 .71.31l4.56 4.72a7.73 7.73 0 0 1-.09 10.77A7.33 7.33 0 0 1 12 21.1Z"/></g></g>`
	dropletOffFillInnerSVG                = `<g id="evaDropletOffFill0"><g id="evaDropletOffFill1"><path id="evaDropletOffFill2" fill="currentColor" d="M19 16.14A7.73 7.73 0 0 0 17.34 8l-4.56-4.69a1 1 0 0 0-.71-.31a1 1 0 0 0-.72.3L8.74 5.92ZM6 8.82a7.73 7.73 0 0 0 .64 9.9A7.44 7.44 0 0 0 11.92 21a7.34 7.34 0 0 0 4.64-1.6Zm14.71 10.47l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	dropletOffOutlineInnerSVG             = `<g id="evaDropletOffOutline0"><g id="evaDropletOffOutline1"><path id="evaDropletOffOutline2" fill="currentColor" d="M12 19a5.4 5.4 0 0 1-3.88-1.64a5.73 5.73 0 0 1-.69-7.11L6 8.82a7.74 7.74 0 0 0 .7 9.94A7.37 7.37 0 0 0 12 21a7.36 7.36 0 0 0 4.58-1.59L15.15 18A5.43 5.43 0 0 1 12 19Zm0-13.57l3.88 4a5.71 5.71 0 0 1 1.49 5.15L19 16.15A7.72 7.72 0 0 0 17.31 8l-4.6-4.7A1 1 0 0 0 12 3a1 1 0 0 0-.72.3L8.73 5.9l1.42 1.42Zm8.71 13.86l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	dropletOutlineInnerSVG                = `<g id="evaDropletOutline0"><g id="evaDropletOutline1"><path id="evaDropletOutline2" fill="currentColor" d="M12 21.1a7.4 7.4 0 0 1-5.28-2.28a7.73 7.73 0 0 1 .1-10.77l4.64-4.65a.94.94 0 0 1 .71-.3a1 1 0 0 1 .71.31l4.56 4.72a7.73 7.73 0 0 1-.09 10.77A7.33 7.33 0 0 1 12 21.1Zm.13-15.57L8.24 9.45a5.74 5.74 0 0 0-.07 8A5.43 5.43 0 0 0 12 19.1a5.42 5.42 0 0 0 3.9-1.61a5.72 5.72 0 0 0 .06-8Z"/></g></g>`
	editFillInnerSVG                      = `<g id="evaEditFill0"><g id="evaEditFill1"><path id="evaEditFill2" fill="currentColor" d="M19.4 7.34L16.66 4.6A2 2 0 0 0 14 4.53l-9 9a2 2 0 0 0-.57 1.21L4 18.91a1 1 0 0 0 .29.8A1 1 0 0 0 5 20h.09l4.17-.38a2 2 0 0 0 1.21-.57l9-9a1.92 1.92 0 0 0-.07-2.71ZM16 10.68L13.32 8l1.95-2L18 8.73Z"/></g></g>`
	editOutlineInnerSVG                   = `<g id="evaEditOutline0"><g id="evaEditOutline1"><path id="evaEditOutline2" fill="currentColor" d="M19.4 7.34L16.66 4.6A2 2 0 0 0 14 4.53l-9 9a2 2 0 0 0-.57 1.21L4 18.91a1 1 0 0 0 .29.8A1 1 0 0 0 5 20h.09l4.17-.38a2 2 0 0 0 1.21-.57l9-9a1.92 1.92 0 0 0-.07-2.71ZM9.08 17.62l-3 .28l.27-3L12 9.32l2.7 2.7ZM16 10.68L13.32 8l1.95-2L18 8.73Z"/></g></g>`
	editTwoFillInnerSVG                   = `<g id="evaEdit2Fill0"><g id="evaEdit2Fill1"><path id="evaEdit2Fill2" fill="currentColor" d="M19 20H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2ZM5 18h.09l4.17-.38a2 2 0 0 0 1.21-.57l9-9a1.92 1.92 0 0 0-.07-2.71L16.66 2.6A2 2 0 0 0 14 2.53l-9 9a2 2 0 0 0-.57 1.21L4 16.91a1 1 0 0 0 .29.8A1 1 0 0 0 5 18ZM15.27 4L18 6.73l-2 1.95L13.32 6Z"/></g></g>`
	editTwoOutlineInnerSVG                = `<g id="evaEdit2Outline0"><g id="evaEdit2Outline1"><path id="evaEdit2Outline2" fill="currentColor" d="M19 20H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2ZM5 18h.09l4.17-.38a2 2 0 0 0 1.21-.57l9-9a1.92 1.92 0 0 0-.07-2.71L16.66 2.6A2 2 0 0 0 14 2.53l-9 9a2 2 0 0 0-.57 1.21L4 16.91a1 1 0 0 0 .29.8A1 1 0 0 0 5 18ZM15.27 4L18 6.73l-2 1.95L13.32 6Zm-8.9 8.91L12 7.32l2.7 2.7l-5.6 5.6l-3 .28Z"/></g></g>`
	emailFillInnerSVG                     = `<g id="evaEmailFill0"><g id="evaEmailFill1"><path id="evaEmailFill2" fill="currentColor" d="M19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Zm0 2l-6.5 4.47a1 1 0 0 1-1 0L5 6Z"/></g></g>`
	emailOutlineInnerSVG                  = `<g id="evaEmailOutline0"><g id="evaEmailOutline1"><path id="evaEmailOutline2" fill="currentColor" d="M19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Zm-.67 2L12 10.75L5.67 6ZM19 18H5a1 1 0 0 1-1-1V7.25l7.4 5.55a1 1 0 0 0 .6.2a1 1 0 0 0 .6-.2L20 7.25V17a1 1 0 0 1-1 1Z"/></g></g>`
	expandFillInnerSVG                    = `<g id="evaExpandFill0"><g id="evaExpandFill1"><path id="evaExpandFill2" fill="currentColor" d="M20 5a1 1 0 0 0-1-1h-5a1 1 0 0 0 0 2h2.57l-3.28 3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L18 7.42V10a1 1 0 0 0 1 1a1 1 0 0 0 1-1Zm-9.29 8.29a1 1 0 0 0-1.42 0L6 16.57V14a1 1 0 0 0-1-1a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2H7.42l3.29-3.29a1 1 0 0 0 0-1.42Z"/></g></g>`
	expandOutlineInnerSVG                 = `<g id="evaExpandOutline0"><g id="evaExpandOutline1"><path id="evaExpandOutline2" fill="currentColor" d="M20 5a1 1 0 0 0-1-1h-5a1 1 0 0 0 0 2h2.57l-3.28 3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L18 7.42V10a1 1 0 0 0 1 1a1 1 0 0 0 1-1Zm-9.29 8.29a1 1 0 0 0-1.42 0L6 16.57V14a1 1 0 0 0-1-1a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2H7.42l3.29-3.29a1 1 0 0 0 0-1.42Z"/></g></g>`
	externalLinkFillInnerSVG              = `<g id="evaExternalLinkFill0"><g id="evaExternalLinkFill1"><g id="evaExternalLinkFill2" fill="currentColor"><path d="M20 11a1 1 0 0 0-1 1v6a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h6a1 1 0 0 0 0-2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1Z"/><path d="M16 5h1.58l-6.29 6.28a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L19 6.42V8a1 1 0 0 0 1 1a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-4a1 1 0 0 0 0 2Z"/></g></g></g>`
	externalLinkOutlineInnerSVG           = `<g id="evaExternalLinkOutline0"><g id="evaExternalLinkOutline1"><g id="evaExternalLinkOutline2" fill="currentColor"><path d="M20 11a1 1 0 0 0-1 1v6a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h6a1 1 0 0 0 0-2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1Z"/><path d="M16 5h1.58l-6.29 6.28a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L19 6.42V8a1 1 0 0 0 1 1a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-4a1 1 0 0 0 0 2Z"/></g></g></g>`
	eyeFillInnerSVG                       = `<g id="evaEyeFill0"><g id="evaEyeFill1"><g id="evaEyeFill2" fill="currentColor"><circle cx="12" cy="12" r="1.5"/><path d="M21.87 11.5c-.64-1.11-4.16-6.68-10.14-6.5c-5.53.14-8.73 5-9.6 6.5a1 1 0 0 0 0 1c.63 1.09 4 6.5 9.89 6.5h.25c5.53-.14 8.74-5 9.6-6.5a1 1 0 0 0 0-1Zm-9.87 4a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5Z"/></g></g></g>`
	eyeOffFillInnerSVG                    = `<g id="evaEyeOffFill0"><g id="evaEyeOffFill1"><g id="evaEyeOffFill2" fill="currentColor"><circle cx="12" cy="12" r="1.5"/><path d="M15.29 18.12L14 16.78l-.07-.07l-1.27-1.27a4.07 4.07 0 0 1-.61.06A3.5 3.5 0 0 1 8.5 12a4.07 4.07 0 0 1 .06-.61l-2-2L5 7.87a15.89 15.89 0 0 0-2.87 3.63a1 1 0 0 0 0 1c.63 1.09 4 6.5 9.89 6.5h.25a9.48 9.48 0 0 0 3.23-.67ZM8.59 5.76l2.8 2.8A4.07 4.07 0 0 1 12 8.5a3.5 3.5 0 0 1 3.5 3.5a4.07 4.07 0 0 1-.06.61l2.68 2.68l.84.84a15.89 15.89 0 0 0 2.91-3.63a1 1 0 0 0 0-1c-.64-1.11-4.16-6.68-10.14-6.5a9.48 9.48 0 0 0-3.23.67Zm12.12 13.53L19.41 18l-2-2l-9.52-9.53L6.42 5L4.71 3.29a1 1 0 0 0-1.42 1.42L5.53 7l1.75 1.7l7.31 7.3l.07.07L16 17.41l.59.59l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g></g>`
	eyeOffOutlineInnerSVG                 = `<g id="evaEyeOffOutline0"><g id="evaEyeOffOutline1"><g id="evaEyeOffOutline2" fill="currentColor"><path d="M4.71 3.29a1 1 0 0 0-1.42 1.42l5.63 5.63a3.5 3.5 0 0 0 4.74 4.74l5.63 5.63a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM12 13.5a1.5 1.5 0 0 1-1.5-1.5s0-.05 0-.07l1.56 1.56Z"/><path d="M12.22 17c-4.3.1-7.12-3.59-8-5a13.7 13.7 0 0 1 2.24-2.72L5 7.87a15.89 15.89 0 0 0-2.87 3.63a1 1 0 0 0 0 1c.63 1.09 4 6.5 9.89 6.5h.25a9.48 9.48 0 0 0 3.23-.67l-1.58-1.58a7.74 7.74 0 0 1-1.7.25Zm9.65-5.5c-.64-1.11-4.17-6.68-10.14-6.5a9.48 9.48 0 0 0-3.23.67l1.58 1.58a7.74 7.74 0 0 1 1.7-.25c4.29-.11 7.11 3.59 8 5a13.7 13.7 0 0 1-2.29 2.72L19 16.13a15.89 15.89 0 0 0 2.91-3.63a1 1 0 0 0-.04-1Z"/></g></g></g>`
	eyeOffTwoFillInnerSVG                 = `<g id="evaEyeOff2Fill0"><g id="evaEyeOff2Fill1"><path id="evaEyeOff2Fill2" fill="currentColor" d="M17.81 13.39A8.93 8.93 0 0 0 21 7.62a1 1 0 1 0-2-.24a7.07 7.07 0 0 1-14 0a1 1 0 1 0-2 .24a8.93 8.93 0 0 0 3.18 5.77l-2.3 2.32a1 1 0 0 0 1.41 1.41l2.61-2.6a9.06 9.06 0 0 0 3.1.92V19a1 1 0 0 0 2 0v-3.56a9.06 9.06 0 0 0 3.1-.92l2.61 2.6a1 1 0 0 0 1.41-1.41Z"/></g></g>`
	eyeOffTwoOutlineInnerSVG              = `<g id="evaEyeOff2Outline0"><g id="evaEyeOff2Outline1"><path id="evaEyeOff2Outline2" fill="currentColor" d="M17.81 13.39A8.93 8.93 0 0 0 21 7.62a1 1 0 1 0-2-.24a7.07 7.07 0 0 1-14 0a1 1 0 1 0-2 .24a8.93 8.93 0 0 0 3.18 5.77l-2.3 2.32a1 1 0 0 0 1.41 1.41l2.61-2.6a9.06 9.06 0 0 0 3.1.92V19a1 1 0 0 0 2 0v-3.56a9.06 9.06 0 0 0 3.1-.92l2.61 2.6a1 1 0 0 0 1.41-1.41Z"/></g></g>`
	eyeOutlineInnerSVG                    = `<g id="evaEyeOutline0"><g id="evaEyeOutline1"><g id="evaEyeOutline2" fill="currentColor"><path d="M21.87 11.5c-.64-1.11-4.16-6.68-10.14-6.5c-5.53.14-8.73 5-9.6 6.5a1 1 0 0 0 0 1c.63 1.09 4 6.5 9.89 6.5h.25c5.53-.14 8.74-5 9.6-6.5a1 1 0 0 0 0-1ZM12.22 17c-4.31.1-7.12-3.59-8-5c1-1.61 3.61-4.9 7.61-5c4.29-.11 7.11 3.59 8 5c-1.03 1.61-3.61 4.9-7.61 5Z"/><path d="M12 8.5a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 12 8.5Zm0 5a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Z"/></g></g></g>`
	facebookFillInnerSVG                  = `<g id="evaFacebookFill0"><g id="evaFacebookFill1"><path id="evaFacebookFill2" fill="currentColor" d="M17 3.5a.5.5 0 0 0-.5-.5H14a4.77 4.77 0 0 0-5 4.5v2.7H6.5a.5.5 0 0 0-.5.5v2.6a.5.5 0 0 0 .5.5H9v6.7a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-6.7h2.62a.5.5 0 0 0 .49-.37l.72-2.6a.5.5 0 0 0-.48-.63H13V7.5a1 1 0 0 1 1-.9h2.5a.5.5 0 0 0 .5-.5Z"/></g></g>`
	facebookOutlineInnerSVG               = `<g id="evaFacebookOutline0"><g id="evaFacebookOutline1"><path id="evaFacebookOutline2" fill="currentColor" d="M13 22H9a1 1 0 0 1-1-1v-6.2H6a1 1 0 0 1-1-1v-3.6a1 1 0 0 1 1-1h2V7.5A5.77 5.77 0 0 1 14 2h3a1 1 0 0 1 1 1v3.6a1 1 0 0 1-1 1h-3v1.6h3a1 1 0 0 1 .8.39a1 1 0 0 1 .16.88l-1 3.6a1 1 0 0 1-1 .73H14V21a1 1 0 0 1-1 1Zm-3-2h2v-6.2a1 1 0 0 1 1-1h2.24l.44-1.6H13a1 1 0 0 1-1-1V7.5a2 2 0 0 1 2-1.9h2V4h-2a3.78 3.78 0 0 0-4 3.5v2.7a1 1 0 0 1-1 1H7v1.6h2a1 1 0 0 1 1 1Z"/></g></g>`
	fileAddFillInnerSVG                   = `<g id="evaFileAddFill0"><g id="evaFileAddFill1"><path id="evaFileAddFill2" fill="currentColor" d="m19.74 7.33l-4.44-5a1 1 0 0 0-.74-.33h-8A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V8a1 1 0 0 0-.26-.67ZM14 15h-1v1a1 1 0 0 1-2 0v-1h-1a1 1 0 0 1 0-2h1v-1a1 1 0 0 1 2 0v1h1a1 1 0 0 1 0 2Zm.71-7a.79.79 0 0 1-.71-.85V4l3.74 4Z"/></g></g>`
	fileAddOutlineInnerSVG                = `<g id="evaFileAddOutline0"><g id="evaFileAddOutline1"><g id="evaFileAddOutline2" fill="currentColor"><path d="m19.74 8.33l-5.44-6a1 1 0 0 0-.74-.33h-7A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V9a1 1 0 0 0-.26-.67ZM14 5l2.74 3h-2a.79.79 0 0 1-.74-.85Zm3.44 15H6.56a.53.53 0 0 1-.56-.5v-15a.53.53 0 0 1 .56-.5H12v3.15A2.79 2.79 0 0 0 14.71 10H18v9.5a.53.53 0 0 1-.56.5Z"/><path d="M14 13h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2Z"/></g></g></g>`
	fileFillInnerSVG                      = `<g id="evaFileFill0"><g id="evaFileFill1"><path id="evaFileFill2" fill="currentColor" d="m19.74 7.33l-4.44-5a1 1 0 0 0-.74-.33h-8A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V8a1 1 0 0 0-.26-.67ZM14 4l3.74 4h-3a.79.79 0 0 1-.74-.85Z"/></g></g>`
	fileOutlineInnerSVG                   = `<g id="evaFileOutline0"><g id="evaFileOutline1"><path id="evaFileOutline2" fill="currentColor" d="m19.74 8.33l-5.44-6a1 1 0 0 0-.74-.33h-7A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V9a1 1 0 0 0-.26-.67ZM17.65 9h-3.94a.79.79 0 0 1-.71-.85V4h.11Zm-.21 11H6.56a.53.53 0 0 1-.56-.5v-15a.53.53 0 0 1 .56-.5H11v4.15A2.79 2.79 0 0 0 13.71 11H18v8.5a.53.53 0 0 1-.56.5Z"/></g></g>`
	fileRemoveFillInnerSVG                = `<g id="evaFileRemoveFill0"><g id="evaFileRemoveFill1"><path id="evaFileRemoveFill2" fill="currentColor" d="m19.74 7.33l-4.44-5a1 1 0 0 0-.74-.33h-8A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V8a1 1 0 0 0-.26-.67ZM14 15h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Zm.71-7a.79.79 0 0 1-.71-.85V4l3.74 4Z"/></g></g>`
	fileRemoveOutlineInnerSVG             = `<g id="evaFileRemoveOutline0"><g id="evaFileRemoveOutline1"><g id="evaFileRemoveOutline2" fill="currentColor"><path d="m19.74 8.33l-5.44-6a1 1 0 0 0-.74-.33h-7A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V9a1 1 0 0 0-.26-.67ZM14 5l2.74 3h-2a.79.79 0 0 1-.74-.85Zm3.44 15H6.56a.53.53 0 0 1-.56-.5v-15a.53.53 0 0 1 .56-.5H12v3.15A2.79 2.79 0 0 0 14.71 10H18v9.5a.53.53 0 0 1-.56.5Z"/><path d="M14 13h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Z"/></g></g></g>`
	fileTextFillInnerSVG                  = `<g id="evaFileTextFill0"><g id="evaFileTextFill1"><path id="evaFileTextFill2" fill="currentColor" d="m19.74 7.33l-4.44-5a1 1 0 0 0-.74-.33h-8A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V8a1 1 0 0 0-.26-.67ZM9 12h3a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2Zm6 6H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Zm-.29-10a.79.79 0 0 1-.71-.85V4l3.74 4Z"/></g></g>`
	fileTextOutlineInnerSVG               = `<g id="evaFileTextOutline0"><g id="evaFileTextOutline1"><g id="evaFileTextOutline2" fill="currentColor"><path d="M15 16H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Zm-6-2h3a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2Z"/><path d="m19.74 8.33l-5.44-6a1 1 0 0 0-.74-.33h-7A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V9a1 1 0 0 0-.26-.67ZM14 5l2.74 3h-2a.79.79 0 0 1-.74-.85Zm3.44 15H6.56a.53.53 0 0 1-.56-.5v-15a.53.53 0 0 1 .56-.5H12v3.15A2.79 2.79 0 0 0 14.71 10H18v9.5a.53.53 0 0 1-.56.5Z"/></g></g></g>`
	filmFillInnerSVG                      = `<g id="evaFilmFill0"><g id="evaFilmFill1"><path id="evaFilmFill2" fill="currentColor" d="M18.26 3H5.74A2.74 2.74 0 0 0 3 5.74v12.52A2.74 2.74 0 0 0 5.74 21h12.52A2.74 2.74 0 0 0 21 18.26V5.74A2.74 2.74 0 0 0 18.26 3ZM7 11H5V9h2Zm-2 2h2v2H5Zm14-2h-2V9h2Zm-2 2h2v2h-2Zm2-7.26V7h-2V5h1.26a.74.74 0 0 1 .74.74ZM5.74 5H7v2H5V5.74A.74.74 0 0 1 5.74 5ZM5 18.26V17h2v2H5.74a.74.74 0 0 1-.74-.74Zm14 0a.74.74 0 0 1-.74.74H17v-2h2Z"/></g></g>`
	filmOutlineInnerSVG                   = `<g id="evaFilmOutline0"><g id="evaFilmOutline1"><path id="evaFilmOutline2" fill="currentColor" d="M18.26 3H5.74A2.74 2.74 0 0 0 3 5.74v12.52A2.74 2.74 0 0 0 5.74 21h12.52A2.74 2.74 0 0 0 21 18.26V5.74A2.74 2.74 0 0 0 18.26 3ZM7 11H5V9h2Zm-2 2h2v2H5Zm4-8h6v14H9Zm10 6h-2V9h2Zm-2 2h2v2h-2Zm2-7.26V7h-2V5h1.26a.74.74 0 0 1 .74.74ZM5.74 5H7v2H5V5.74A.74.74 0 0 1 5.74 5ZM5 18.26V17h2v2H5.74a.74.74 0 0 1-.74-.74Zm14 0a.74.74 0 0 1-.74.74H17v-2h2Z"/></g></g>`
	flagFillInnerSVG                      = `<g id="evaFlagFill0"><g id="evaFlagFill1"><path id="evaFlagFill2" fill="currentColor" d="M19.27 4.68a1.79 1.79 0 0 0-1.6-.25a7.53 7.53 0 0 1-2.17.28a8.54 8.54 0 0 1-3.13-.78A10.15 10.15 0 0 0 8.5 3c-2.89 0-4 1-4.2 1.14a1 1 0 0 0-.3.72V20a1 1 0 0 0 2 0v-4.3a6.28 6.28 0 0 1 2.5-.41a8.54 8.54 0 0 1 3.13.78a10.15 10.15 0 0 0 3.87.93a7.66 7.66 0 0 0 3.5-.7a1.74 1.74 0 0 0 1-1.55V6.11a1.77 1.77 0 0 0-.73-1.43Z"/></g></g>`
	flagOutlineInnerSVG                   = `<g id="evaFlagOutline0"><g id="evaFlagOutline1"><path id="evaFlagOutline2" fill="currentColor" d="M19.27 4.68a1.79 1.79 0 0 0-1.6-.25a7.53 7.53 0 0 1-2.17.28a8.54 8.54 0 0 1-3.13-.78A10.15 10.15 0 0 0 8.5 3c-2.89 0-4 1-4.2 1.14a1 1 0 0 0-.3.72V20a1 1 0 0 0 2 0v-4.3a6.28 6.28 0 0 1 2.5-.41a8.54 8.54 0 0 1 3.13.78a10.15 10.15 0 0 0 3.87.93a7.66 7.66 0 0 0 3.5-.7a1.74 1.74 0 0 0 1-1.55V6.11a1.77 1.77 0 0 0-.73-1.43ZM18 14.59a6.32 6.32 0 0 1-2.5.41a8.36 8.36 0 0 1-3.13-.79a10.34 10.34 0 0 0-3.87-.92a9.51 9.51 0 0 0-2.5.29V5.42A6.13 6.13 0 0 1 8.5 5a8.36 8.36 0 0 1 3.13.79a10.34 10.34 0 0 0 3.87.92a9.41 9.41 0 0 0 2.5-.3Z"/></g></g>`
	flashFillInnerSVG                     = `<g id="evaFlashFill0"><g id="evaFlashFill1"><path id="evaFlashFill2" fill="currentColor" d="M11.11 23a1 1 0 0 1-.34-.06a1 1 0 0 1-.65-1.05l.77-7.09H5a1 1 0 0 1-.83-1.56l7.89-11.8a1 1 0 0 1 1.17-.38a1 1 0 0 1 .65 1l-.77 7.14H19a1 1 0 0 1 .83 1.56l-7.89 11.8a1 1 0 0 1-.83.44Z"/></g></g>`
	flashOffFillInnerSVG                  = `<g id="evaFlashOffFill0"><g id="evaFlashOffFill1"><path id="evaFlashOffFill2" fill="currentColor" d="m17.33 14.5l2.5-3.74A1 1 0 0 0 19 9.2h-5.89l.77-7.09a1 1 0 0 0-.65-1a1 1 0 0 0-1.17.38L8.94 6.11Zm-10.66-5l-2.5 3.74A1 1 0 0 0 5 14.8h5.89l-.77 7.09a1 1 0 0 0 .65 1.05a1 1 0 0 0 .34.06a1 1 0 0 0 .83-.44l3.12-4.67Zm14.04 9.79l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	flashOffOutlineInnerSVG               = `<g id="evaFlashOffOutline0"><g id="evaFlashOffOutline1"><path id="evaFlashOffOutline2" fill="currentColor" d="m20.71 19.29l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-8.17-1.23l.27-2.42L10 12.8H6.87l1.24-1.86L6.67 9.5l-2.5 3.74A1 1 0 0 0 5 14.8h5.89l-.77 7.09a1 1 0 0 0 .65 1.05a1 1 0 0 0 .34.06a1 1 0 0 0 .83-.44l3.12-4.67l-1.44-1.44ZM11.46 5.94l-.27 2.42L14 11.2h3.1l-1.24 1.86l1.44 1.44l2.5-3.74A1 1 0 0 0 19 9.2h-5.89l.77-7.09a1 1 0 0 0-.65-1a1 1 0 0 0-1.17.38L8.94 6.11l1.44 1.44Z"/></g></g>`
	flashOutlineInnerSVG                  = `<g id="evaFlashOutline0"><g id="evaFlashOutline1"><path id="evaFlashOutline2" fill="currentColor" d="M11.11 23a1 1 0 0 1-.34-.06a1 1 0 0 1-.65-1.05l.77-7.09H5a1 1 0 0 1-.83-1.56l7.89-11.8a1 1 0 0 1 1.17-.38a1 1 0 0 1 .65 1l-.77 7.14H19a1 1 0 0 1 .83 1.56l-7.89 11.8a1 1 0 0 1-.83.44ZM6.87 12.8H12a1 1 0 0 1 .74.33a1 1 0 0 1 .25.78l-.45 4.15l4.59-6.86H12a1 1 0 0 1-1-1.11l.45-4.15Z"/></g></g>`
	flipFillInnerSVG                      = `<g id="evaFlipFill0"><g id="evaFlipFill1"><path id="evaFlipFill2" fill="currentColor" d="M5 6.09v12l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 0-1.42a1 1 0 0 0-1.42 0L7 18.09v-12A1.56 1.56 0 0 1 8.53 4.5H11a1 1 0 0 0 0-2H8.53A3.56 3.56 0 0 0 5 6.09Zm9.29-.3a1 1 0 0 0 1.42 1.42L17 5.91v12a1.56 1.56 0 0 1-1.53 1.59H13a1 1 0 0 0 0 2h2.47A3.56 3.56 0 0 0 19 17.91v-12l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-1.42 0Z"/></g></g>`
	flipOutlineInnerSVG                   = `<g id="evaFlipOutline0"><g id="evaFlipOutline1"><path id="evaFlipOutline2" fill="currentColor" d="M5 6.09v12l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 0-1.42a1 1 0 0 0-1.42 0L7 18.09v-12A1.56 1.56 0 0 1 8.53 4.5H11a1 1 0 0 0 0-2H8.53A3.56 3.56 0 0 0 5 6.09Zm9.29-.3a1 1 0 0 0 1.42 1.42L17 5.91v12a1.56 1.56 0 0 1-1.53 1.59H13a1 1 0 0 0 0 2h2.47A3.56 3.56 0 0 0 19 17.91v-12l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-1.42 0Z"/></g></g>`
	flipTwoFillInnerSVG                   = `<g id="evaFlip2Fill0"><g id="evaFlip2Fill1"><path id="evaFlip2Fill2" fill="currentColor" d="M6.09 19h12l-1.3 1.29a1 1 0 0 0 1.42 1.42l3-3a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.42l1.3 1.29h-12a1.56 1.56 0 0 1-1.59-1.53V13a1 1 0 0 0-2 0v2.47A3.56 3.56 0 0 0 6.09 19Zm-.3-9.29a1 1 0 1 0 1.42-1.42L5.91 7h12a1.56 1.56 0 0 1 1.59 1.53V11a1 1 0 0 0 2 0V8.53A3.56 3.56 0 0 0 17.91 5h-12l1.3-1.29a1 1 0 0 0 0-1.42a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 0 1.42Z"/></g></g>`
	flipTwoOutlineInnerSVG                = `<g id="evaFlip2Outline0"><g id="evaFlip2Outline1"><path id="evaFlip2Outline2" fill="currentColor" d="M6.09 19h12l-1.3 1.29a1 1 0 0 0 1.42 1.42l3-3a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.42l1.3 1.29h-12a1.56 1.56 0 0 1-1.59-1.53V13a1 1 0 0 0-2 0v2.47A3.56 3.56 0 0 0 6.09 19Zm-.3-9.29a1 1 0 1 0 1.42-1.42L5.91 7h12a1.56 1.56 0 0 1 1.59 1.53V11a1 1 0 0 0 2 0V8.53A3.56 3.56 0 0 0 17.91 5h-12l1.3-1.29a1 1 0 0 0 0-1.42a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 0 1.42Z"/></g></g>`
	folderAddFillInnerSVG                 = `<g id="evaFolderAddFill0"><g id="evaFolderAddFill1"><path id="evaFolderAddFill2" fill="currentColor" d="M19.5 7.05h-7L9.87 3.87a1 1 0 0 0-.77-.37H4.5A2.47 2.47 0 0 0 2 5.93v12.14a2.47 2.47 0 0 0 2.5 2.43h15a2.47 2.47 0 0 0 2.5-2.43V9.48a2.47 2.47 0 0 0-2.5-2.43ZM14 15h-1v1a1 1 0 0 1-2 0v-1h-1a1 1 0 0 1 0-2h1v-1a1 1 0 0 1 2 0v1h1a1 1 0 0 1 0 2Z"/></g></g>`
	folderAddOutlineInnerSVG              = `<g id="evaFolderAddOutline0"><g id="evaFolderAddOutline1"><g id="evaFolderAddOutline2" fill="currentColor"><path d="M14 13h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2Z"/><path d="M19.5 7.05h-7L9.87 3.87a1 1 0 0 0-.77-.37H4.5A2.47 2.47 0 0 0 2 5.93v12.14a2.47 2.47 0 0 0 2.5 2.43h15a2.47 2.47 0 0 0 2.5-2.43V9.48a2.47 2.47 0 0 0-2.5-2.43Zm.5 11a.46.46 0 0 1-.5.43h-15a.46.46 0 0 1-.5-.43V5.93a.46.46 0 0 1 .5-.43h4.13l2.6 3.18a1 1 0 0 0 .77.37h7.5a.46.46 0 0 1 .5.43Z"/></g></g></g>`
	folderFillInnerSVG                    = `<g id="evaFolderFill0"><g id="evaFolderFill1"><path id="evaFolderFill2" fill="currentColor" d="M19.5 20.5h-15A2.47 2.47 0 0 1 2 18.07V5.93A2.47 2.47 0 0 1 4.5 3.5h4.6a1 1 0 0 1 .77.37l2.6 3.18h7A2.47 2.47 0 0 1 22 9.48v8.59a2.47 2.47 0 0 1-2.5 2.43Z"/></g></g>`
	folderOutlineInnerSVG                 = `<g id="evaFolderOutline0"><g id="evaFolderOutline1"><path id="evaFolderOutline2" fill="currentColor" d="M19.5 20.5h-15A2.47 2.47 0 0 1 2 18.07V5.93A2.47 2.47 0 0 1 4.5 3.5h4.6a1 1 0 0 1 .77.37l2.6 3.18h7A2.47 2.47 0 0 1 22 9.48v8.59a2.47 2.47 0 0 1-2.5 2.43ZM4 13.76v4.31a.46.46 0 0 0 .5.43h15a.46.46 0 0 0 .5-.43V9.48a.46.46 0 0 0-.5-.43H12a1 1 0 0 1-.77-.37L8.63 5.5H4.5a.46.46 0 0 0-.5.43Z"/></g></g>`
	folderRemoveFillInnerSVG              = `<g id="evaFolderRemoveFill0"><g id="evaFolderRemoveFill1"><path id="evaFolderRemoveFill2" fill="currentColor" d="M19.5 7.05h-7L9.87 3.87a1 1 0 0 0-.77-.37H4.5A2.47 2.47 0 0 0 2 5.93v12.14a2.47 2.47 0 0 0 2.5 2.43h15a2.47 2.47 0 0 0 2.5-2.43V9.48a2.47 2.47 0 0 0-2.5-2.43ZM14 15h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Z"/></g></g>`
	folderRemoveOutlineInnerSVG           = `<g id="evaFolderRemoveOutline0"><g id="evaFolderRemoveOutline1"><g id="evaFolderRemoveOutline2" fill="currentColor"><path d="M14 13h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Z"/><path d="M19.5 7.05h-7L9.87 3.87a1 1 0 0 0-.77-.37H4.5A2.47 2.47 0 0 0 2 5.93v12.14a2.47 2.47 0 0 0 2.5 2.43h15a2.47 2.47 0 0 0 2.5-2.43V9.48a2.47 2.47 0 0 0-2.5-2.43Zm.5 11a.46.46 0 0 1-.5.43h-15a.46.46 0 0 1-.5-.43V5.93a.46.46 0 0 1 .5-.43h4.13l2.6 3.18a1 1 0 0 0 .77.37h7.5a.46.46 0 0 1 .5.43Z"/></g></g></g>`
	funnelFillInnerSVG                    = `<g id="evaFunnelFill0"><g id="evaFunnelFill1"><path id="evaFunnelFill2" fill="currentColor" d="M13.9 22a1 1 0 0 1-.6-.2l-4-3.05a1 1 0 0 1-.39-.8v-3.27l-4.8-9.22A1 1 0 0 1 5 4h14a1 1 0 0 1 .86.49a1 1 0 0 1 0 1l-5 9.21V21a1 1 0 0 1-.55.9a1 1 0 0 1-.41.1Z"/></g></g>`
	funnelOutlineInnerSVG                 = `<g id="evaFunnelOutline0"><g id="evaFunnelOutline1"><path id="evaFunnelOutline2" fill="currentColor" d="M13.9 22a1 1 0 0 1-.6-.2l-4-3.05a1 1 0 0 1-.39-.8v-3.27l-4.8-9.22A1 1 0 0 1 5 4h14a1 1 0 0 1 .86.49a1 1 0 0 1 0 1l-5 9.21V21a1 1 0 0 1-.55.9a1 1 0 0 1-.41.1Zm-3-4.54l2 1.53v-4.55A1 1 0 0 1 13 14l4.3-8H6.64l4.13 8a1 1 0 0 1 .11.46Z"/></g></g>`
	giftFillInnerSVG                      = `<g id="evaGiftFill0"><g id="evaGiftFill1"><path id="evaGiftFill2" fill="currentColor" d="M4.64 15.27v4.82a.92.92 0 0 0 .92.91h5.62v-5.73ZM12.82 21h5.62a.92.92 0 0 0 .92-.91v-4.82h-6.54ZM20.1 7.09h-1.84a2.82 2.82 0 0 0 .29-1.23A2.87 2.87 0 0 0 15.68 3A4.21 4.21 0 0 0 12 5.57A4.21 4.21 0 0 0 8.32 3a2.87 2.87 0 0 0-2.87 2.86a2.82 2.82 0 0 0 .29 1.23H3.9c-.5 0-.9.59-.9 1.31v3.93c0 .72.4 1.31.9 1.31h7.28V7.09h1.64v6.55h7.28c.5 0 .9-.59.9-1.31V8.4c0-.72-.4-1.31-.9-1.31Zm-11.78 0a1.23 1.23 0 1 1 0-2.45c1.4 0 2.19 1.44 2.58 2.45Zm7.36 0H13.1c.39-1 1.18-2.45 2.58-2.45a1.23 1.23 0 1 1 0 2.45Z"/></g></g>`
	giftOutlineInnerSVG                   = `<g id="evaGiftOutline0"><g id="evaGiftOutline1"><path id="evaGiftOutline2" fill="currentColor" d="M19.2 7h-.39A3 3 0 0 0 19 6a3.08 3.08 0 0 0-3.14-3A4.46 4.46 0 0 0 12 5.4A4.46 4.46 0 0 0 8.14 3A3.08 3.08 0 0 0 5 6a3 3 0 0 0 .19 1H4.8A2 2 0 0 0 3 9.2v3.6A2.08 2.08 0 0 0 4.5 15v4.37A1.75 1.75 0 0 0 6.31 21h11.38a1.75 1.75 0 0 0 1.81-1.67V15a2.08 2.08 0 0 0 1.5-2.2V9.2A2 2 0 0 0 19.2 7ZM19 9.2v3.6a.56.56 0 0 1 0 .2h-6V9h6a.56.56 0 0 1 0 .2ZM15.86 5A1.08 1.08 0 0 1 17 6a1.08 1.08 0 0 1-1.14 1H13.4a2.93 2.93 0 0 1 2.46-2ZM7 6a1.08 1.08 0 0 1 1.14-1a2.93 2.93 0 0 1 2.45 2H8.14A1.08 1.08 0 0 1 7 6ZM5 9.2A.56.56 0 0 1 5 9h6v4H5a.56.56 0 0 1 0-.2ZM6.5 15H11v4H6.5Zm6.5 4v-4h4.5v4Z"/></g></g>`
	githubFillInnerSVG                    = `<g id="evaGithubFill0"><g id="evaGithubFill1"><g id="evaGithubFill2"><g id="evaGithubFill3"><g id="evaGithubFill4"><g id="evaGithubFill5"><path id="evaGithubFill6" fill="currentColor" d="M12 1A10.89 10.89 0 0 0 1 11.77A10.79 10.79 0 0 0 8.52 22c.55.1.75-.23.75-.52v-1.83c-3.06.65-3.71-1.44-3.71-1.44a2.86 2.86 0 0 0-1.22-1.58c-1-.66.08-.65.08-.65a2.31 2.31 0 0 1 1.68 1.11a2.37 2.37 0 0 0 3.2.89a2.33 2.33 0 0 1 .7-1.44c-2.44-.27-5-1.19-5-5.32a4.15 4.15 0 0 1 1.11-2.91a3.78 3.78 0 0 1 .11-2.84s.93-.29 3 1.1a10.68 10.68 0 0 1 5.5 0c2.1-1.39 3-1.1 3-1.1a3.78 3.78 0 0 1 .11 2.84A4.15 4.15 0 0 1 19 11.2c0 4.14-2.58 5.05-5 5.32a2.5 2.5 0 0 1 .75 2v2.95c0 .35.2.63.75.52A10.8 10.8 0 0 0 23 11.77A10.89 10.89 0 0 0 12 1"/></g></g></g></g></g></g>`
	githubOutlineInnerSVG                 = `<g id="evaGithubOutline0"><path id="evaGithubOutline1" fill="currentColor" d="M16.24 22a1 1 0 0 1-1-1v-2.6a2.15 2.15 0 0 0-.54-1.66a1 1 0 0 1 .61-1.67C17.75 14.78 20 14 20 9.77a4 4 0 0 0-.67-2.22a2.75 2.75 0 0 1-.41-2.06a3.71 3.71 0 0 0 0-1.41a7.65 7.65 0 0 0-2.09 1.09a1 1 0 0 1-.84.15a10.15 10.15 0 0 0-5.52 0a1 1 0 0 1-.84-.15a7.4 7.4 0 0 0-2.11-1.09a3.52 3.52 0 0 0 0 1.41a2.84 2.84 0 0 1-.43 2.08a4.07 4.07 0 0 0-.67 2.23c0 3.89 1.88 4.93 4.7 5.29a1 1 0 0 1 .82.66a1 1 0 0 1-.21 1a2.06 2.06 0 0 0-.55 1.56V21a1 1 0 0 1-2 0v-.57a6 6 0 0 1-5.27-2.09a3.9 3.9 0 0 0-1.16-.88a1 1 0 1 1 .5-1.94a4.93 4.93 0 0 1 2 1.36c1 1 2 1.88 3.9 1.52a3.89 3.89 0 0 1 .23-1.58c-2.06-.52-5-2-5-7a6 6 0 0 1 1-3.33a.85.85 0 0 0 .13-.62a5.69 5.69 0 0 1 .33-3.21a1 1 0 0 1 .63-.57c.34-.1 1.56-.3 3.87 1.2a12.16 12.16 0 0 1 5.69 0c2.31-1.5 3.53-1.31 3.86-1.2a1 1 0 0 1 .63.57a5.71 5.71 0 0 1 .33 3.22a.75.75 0 0 0 .11.57a6 6 0 0 1 1 3.34c0 5.07-2.92 6.54-5 7a4.28 4.28 0 0 1 .22 1.67V21a1 1 0 0 1-.94 1Z"/></g>`
	globeFillInnerSVG                     = `<g id="evaGlobeFill0"><g id="evaGlobeFill1"><path id="evaGlobeFill2" fill="currentColor" d="M22 12A10 10 0 0 0 12 2a10 10 0 0 0 0 20a10 10 0 0 0 10-10Zm-2.07-1H17a12.91 12.91 0 0 0-2.33-6.54A8 8 0 0 1 19.93 11ZM9.08 13H15a11.44 11.44 0 0 1-3 6.61A11 11 0 0 1 9.08 13Zm0-2A11.4 11.4 0 0 1 12 4.4a11.19 11.19 0 0 1 3 6.6Zm.36-6.57A13.18 13.18 0 0 0 7.07 11h-3a8 8 0 0 1 5.37-6.57ZM4.07 13h3a12.86 12.86 0 0 0 2.35 6.56A8 8 0 0 1 4.07 13Zm10.55 6.55A13.14 13.14 0 0 0 17 13h2.95a8 8 0 0 1-5.33 6.55Z"/></g></g>`
	globeOutlineInnerSVG                  = `<g id="evaGlobeOutline0"><g id="evaGlobeOutline1"><path id="evaGlobeOutline2" fill="currentColor" d="M22 12A10 10 0 0 0 12 2a10 10 0 0 0 0 20a10 10 0 0 0 10-10Zm-2.07-1H17a12.91 12.91 0 0 0-2.33-6.54A8 8 0 0 1 19.93 11ZM9.08 13H15a11.44 11.44 0 0 1-3 6.61A11 11 0 0 1 9.08 13Zm0-2A11.4 11.4 0 0 1 12 4.4a11.19 11.19 0 0 1 3 6.6Zm.36-6.57A13.18 13.18 0 0 0 7.07 11h-3a8 8 0 0 1 5.37-6.57ZM4.07 13h3a12.86 12.86 0 0 0 2.35 6.56A8 8 0 0 1 4.07 13Zm10.55 6.55A13.14 13.14 0 0 0 17 13h2.95a8 8 0 0 1-5.33 6.55Z"/></g></g>`
	globeThreeFillInnerSVG                = `<g id="evaGlobe3Fill0"><g id="evaGlobe3Fill1"><path id="evaGlobe3Fill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2ZM5 15.8a8.42 8.42 0 0 0 2 .27a5 5 0 0 0 3.14-1c1.71-1.34 1.71-3.06 1.71-4.44a4.76 4.76 0 0 1 .37-2.34a2.86 2.86 0 0 1 1.12-.91a9.75 9.75 0 0 0 .92-.61a4.55 4.55 0 0 0 1.4-1.87A8 8 0 0 1 19 8.12c-1.43.2-3.46.67-3.86 2.53A7 7 0 0 0 15 12a2.93 2.93 0 0 1-.29 1.47l-.1.17c-.65 1.08-1.38 2.31-.39 4c.12.21.25.41.38.61a2.29 2.29 0 0 1 .52 1.08A7.89 7.89 0 0 1 12 20a8 8 0 0 1-7-4.2Z"/></g></g>`
	globeTwoFillInnerSVG                  = `<g id="evaGlobe2Fill0"><g id="evaGlobe2Fill1"><path id="evaGlobe2Fill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 2a8.19 8.19 0 0 1 1.79.21a2.61 2.61 0 0 1-.78 1c-.22.17-.46.31-.7.46a4.56 4.56 0 0 0-1.85 1.67a6.49 6.49 0 0 0-.62 3.3c0 1.36 0 2.16-.95 2.87c-1.37 1.07-3.46.47-4.76-.07A8.33 8.33 0 0 1 4 12a8 8 0 0 1 8-8Zm4.89 14.32a6.79 6.79 0 0 0-.63-1.14c-.11-.16-.22-.32-.32-.49c-.39-.68-.25-1 .38-2l.1-.17a4.77 4.77 0 0 0 .58-2.43a5.42 5.42 0 0 1 .09-1c.16-.73 1.71-.93 2.67-1a7.94 7.94 0 0 1-2.86 8.28Z"/></g></g>`
	globeTwoOutlineInnerSVG               = `<g id="evaGlobe2Outline0"><g id="evaGlobe2Outline1"><path id="evaGlobe2Outline2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 2a8.19 8.19 0 0 1 1.79.21a2.61 2.61 0 0 1-.78 1c-.22.17-.46.31-.7.46a4.56 4.56 0 0 0-1.85 1.67a6.49 6.49 0 0 0-.62 3.3c0 1.36 0 2.16-.95 2.87c-1.37 1.07-3.46.47-4.76-.07A8.33 8.33 0 0 1 4 12a8 8 0 0 1 8-8ZM5 15.8a8.42 8.42 0 0 0 2 .27a5 5 0 0 0 3.14-1c1.71-1.34 1.71-3.06 1.71-4.44a4.76 4.76 0 0 1 .37-2.34a2.86 2.86 0 0 1 1.12-.91a9.75 9.75 0 0 0 .92-.61a4.55 4.55 0 0 0 1.4-1.87A8 8 0 0 1 19 8.12c-1.43.2-3.46.67-3.86 2.53A7 7 0 0 0 15 12a2.93 2.93 0 0 1-.29 1.47l-.1.17c-.65 1.08-1.38 2.31-.39 4c.12.21.25.41.38.61a2.29 2.29 0 0 1 .52 1.08A7.89 7.89 0 0 1 12 20a8 8 0 0 1-7-4.2Zm11.93 2.52a6.79 6.79 0 0 0-.63-1.14c-.11-.16-.22-.32-.32-.49c-.39-.68-.25-1 .38-2l.1-.17a4.77 4.77 0 0 0 .54-2.43a5.42 5.42 0 0 1 .09-1c.16-.73 1.71-.93 2.67-1a7.94 7.94 0 0 1-2.86 8.28Z"/></g></g>`
	googleFillInnerSVG                    = `<g id="evaGoogleFill0"><g id="evaGoogleFill1"><path id="evaGoogleFill2" fill="currentColor" d="M17.5 14a5.51 5.51 0 0 1-4.5 3.93a6.15 6.15 0 0 1-7-5.45A6 6 0 0 1 12 6a6.12 6.12 0 0 1 2.27.44a.5.5 0 0 0 .64-.21l1.44-2.65a.52.52 0 0 0-.23-.7A10 10 0 0 0 2 12.29A10.12 10.12 0 0 0 11.57 22A10 10 0 0 0 22 12.52v-2a.51.51 0 0 0-.5-.5h-9a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h5"/></g></g>`
	googleOutlineInnerSVG                 = `<g id="evaGoogleOutline0"><g id="evaGoogleOutline1"><path id="evaGoogleOutline2" fill="currentColor" d="M12 22h-.43A10.16 10.16 0 0 1 2 12.29a10 10 0 0 1 14.12-9.41a1.48 1.48 0 0 1 .77.86a1.47 1.47 0 0 1-.1 1.16L15.5 7.28a1.44 1.44 0 0 1-1.83.64A4.5 4.5 0 0 0 8.77 9a4.41 4.41 0 0 0-1.16 3.34a4.36 4.36 0 0 0 1.66 3a4.52 4.52 0 0 0 3.45 1a3.89 3.89 0 0 0 2.63-1.57h-2.9A1.45 1.45 0 0 1 11 13.33v-2.68a1.45 1.45 0 0 1 1.45-1.45h8.1A1.46 1.46 0 0 1 22 10.64v1.88A10 10 0 0 1 12 22Zm0-18a8 8 0 0 0-8 8.24A8.12 8.12 0 0 0 11.65 20A8 8 0 0 0 20 12.42V11.2h-7v1.58h5.31l-.41 1.3a6 6 0 0 1-4.9 4.25A6.58 6.58 0 0 1 8 17a6.33 6.33 0 0 1-.72-9.3A6.52 6.52 0 0 1 14 5.91l.77-1.43A7.9 7.9 0 0 0 12 4Z"/></g></g>`
	gridFillInnerSVG                      = `<g id="evaGridFill0"><g id="evaGridFill1"><path id="evaGridFill2" fill="currentColor" d="M9 3H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Zm10 0h-4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2ZM9 13H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2Zm10 0h-4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2Z"/></g></g>`
	gridOutlineInnerSVG                   = `<g id="evaGridOutline0"><g id="evaGridOutline1"><path id="evaGridOutline2" fill="currentColor" d="M9 3H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2ZM5 9V5h4v4Zm14-6h-4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Zm-4 6V5h4v4Zm-6 4H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2Zm-4 6v-4h4v4Zm14-6h-4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2Zm-4 6v-4h4v4Z"/></g></g>`
	hardDriveFillInnerSVG                 = `<g id="evaHardDriveFill0"><g id="evaHardDriveFill1"><path id="evaHardDriveFill2" fill="currentColor" d="m20.79 11.34l-3.34-6.68A3 3 0 0 0 14.76 3H9.24a3 3 0 0 0-2.69 1.66l-3.34 6.68a2 2 0 0 0-.21.9V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5.76a2 2 0 0 0-.21-.9ZM8 17a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm8 0h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2ZM5.62 11l2.72-5.45a1 1 0 0 1 .9-.55h5.52a1 1 0 0 1 .9.55L18.38 11Z"/></g></g>`
	hardDriveOutlineInnerSVG              = `<g id="evaHardDriveOutline0"><g id="evaHardDriveOutline1"><g id="evaHardDriveOutline2" fill="currentColor"><path d="m20.79 11.34l-3.34-6.68A3 3 0 0 0 14.76 3H9.24a3 3 0 0 0-2.69 1.66l-3.34 6.68a2 2 0 0 0-.21.9V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5.76a2 2 0 0 0-.21-.9ZM8.34 5.55a1 1 0 0 1 .9-.55h5.52a1 1 0 0 1 .9.55L18.38 11H5.62ZM18 19H6a1 1 0 0 1-1-1v-5h14v5a1 1 0 0 1-1 1Z"/><path d="M16 15h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Z"/><circle cx="8" cy="16" r="1"/></g></g></g>`
	hashFillInnerSVG                      = `<g id="evaHashFill0"><g id="evaHashFill1"><path id="evaHashFill2" fill="currentColor" d="M20 14h-4.3l.73-4H20a1 1 0 0 0 0-2h-3.21l.69-3.81A1 1 0 0 0 16.64 3a1 1 0 0 0-1.22.82L14.67 8h-3.88l.69-3.81A1 1 0 0 0 10.64 3a1 1 0 0 0-1.22.82L8.67 8H4a1 1 0 0 0 0 2h4.3l-.73 4H4a1 1 0 0 0 0 2h3.21l-.69 3.81A1 1 0 0 0 7.36 21a1 1 0 0 0 1.22-.82L9.33 16h3.88l-.69 3.81a1 1 0 0 0 .84 1.19a1 1 0 0 0 1.22-.82l.75-4.18H20a1 1 0 0 0 0-2ZM9.7 14l.73-4h3.87l-.73 4Z"/></g></g>`
	hashOutlineInnerSVG                   = `<g id="evaHashOutline0"><g id="evaHashOutline1"><path id="evaHashOutline2" fill="currentColor" d="M20 14h-4.3l.73-4H20a1 1 0 0 0 0-2h-3.21l.69-3.81A1 1 0 0 0 16.64 3a1 1 0 0 0-1.22.82L14.67 8h-3.88l.69-3.81A1 1 0 0 0 10.64 3a1 1 0 0 0-1.22.82L8.67 8H4a1 1 0 0 0 0 2h4.3l-.73 4H4a1 1 0 0 0 0 2h3.21l-.69 3.81A1 1 0 0 0 7.36 21a1 1 0 0 0 1.22-.82L9.33 16h3.88l-.69 3.81a1 1 0 0 0 .84 1.19a1 1 0 0 0 1.22-.82l.75-4.18H20a1 1 0 0 0 0-2ZM9.7 14l.73-4h3.87l-.73 4Z"/></g></g>`
	headphonesFillInnerSVG                = `<g id="evaHeadphonesFill0"><g id="evaHeadphonesFill1"><path id="evaHeadphonesFill2" fill="currentColor" d="M12 2A10.2 10.2 0 0 0 2 12.37V17a4 4 0 1 0 4-4a3.91 3.91 0 0 0-2 .56v-1.19A8.2 8.2 0 0 1 12 4a8.2 8.2 0 0 1 8 8.37v1.19a3.91 3.91 0 0 0-2-.56a4 4 0 1 0 4 4v-4.63A10.2 10.2 0 0 0 12 2Z"/></g></g>`
	headphonesOutlineInnerSVG             = `<g id="evaHeadphonesOutline0"><g id="evaHeadphonesOutline1"><path id="evaHeadphonesOutline2" fill="currentColor" d="M12 2A10.2 10.2 0 0 0 2 12.37V17a4 4 0 1 0 4-4a3.91 3.91 0 0 0-2 .56v-1.19A8.2 8.2 0 0 1 12 4a8.2 8.2 0 0 1 8 8.37v1.19a3.91 3.91 0 0 0-2-.56a4 4 0 1 0 4 4v-4.63A10.2 10.2 0 0 0 12 2ZM6 15a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm12 4a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/></g></g>`
	heartFillInnerSVG                     = `<g id="evaHeartFill0"><g id="evaHeartFill1"><path id="evaHeartFill2" fill="currentColor" d="M12 21a1 1 0 0 1-.71-.29l-7.77-7.78a5.26 5.26 0 0 1 0-7.4a5.24 5.24 0 0 1 7.4 0L12 6.61l1.08-1.08a5.24 5.24 0 0 1 7.4 0a5.26 5.26 0 0 1 0 7.4l-7.77 7.78A1 1 0 0 1 12 21Z"/></g></g>`
	heartOutlineInnerSVG                  = `<g id="evaHeartOutline0"><g id="evaHeartOutline1"><path id="evaHeartOutline2" fill="currentColor" d="M12 21a1 1 0 0 1-.71-.29l-7.77-7.78a5.26 5.26 0 0 1 0-7.4a5.24 5.24 0 0 1 7.4 0L12 6.61l1.08-1.08a5.24 5.24 0 0 1 7.4 0a5.26 5.26 0 0 1 0 7.4l-7.77 7.78A1 1 0 0 1 12 21ZM7.22 6a3.2 3.2 0 0 0-2.28.94a3.24 3.24 0 0 0 0 4.57L12 18.58l7.06-7.07a3.24 3.24 0 0 0 0-4.57a3.32 3.32 0 0 0-4.56 0l-1.79 1.8a1 1 0 0 1-1.42 0L9.5 6.94A3.2 3.2 0 0 0 7.22 6Z"/></g></g>`
	homeFillInnerSVG                      = `<g id="evaHomeFill0"><g id="evaHomeFill1"><g id="evaHomeFill2" fill="currentColor"><path d="M10 14h4v7h-4z"/><path d="M20.42 10.18L12.71 2.3a1 1 0 0 0-1.42 0l-7.71 7.89A2 2 0 0 0 3 11.62V20a2 2 0 0 0 1.89 2H8v-9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v9h3.11A2 2 0 0 0 21 20v-8.38a2.07 2.07 0 0 0-.58-1.44Z"/></g></g></g>`
	homeOutlineInnerSVG                   = `<g id="evaHomeOutline0"><g id="evaHomeOutline1"><path id="evaHomeOutline2" fill="currentColor" d="M20.42 10.18L12.71 2.3a1 1 0 0 0-1.42 0l-7.71 7.89A2 2 0 0 0 3 11.62V20a2 2 0 0 0 1.89 2h14.22A2 2 0 0 0 21 20v-8.38a2.07 2.07 0 0 0-.58-1.44ZM10 20v-6h4v6Zm9 0h-3v-7a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v7H5v-8.42l7-7.15l7 7.19Z"/></g></g>`
	imageFillInnerSVG                     = `<g id="evaImageFill0"><g id="evaImageFill1"><g id="evaImageFill2" fill="currentColor"><path d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3ZM6 5h12a1 1 0 0 1 1 1v8.36l-3.2-2.73a2.77 2.77 0 0 0-3.52 0L5 17.7V6a1 1 0 0 1 1-1Z"/><circle cx="8" cy="8.5" r="1.5"/></g></g></g>`
	imageOutlineInnerSVG                  = `<g id="evaImageOutline0"><g id="evaImageOutline1"><g id="evaImageOutline2" fill="currentColor"><path d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3ZM6 5h12a1 1 0 0 1 1 1v8.36l-3.2-2.73a2.77 2.77 0 0 0-3.52 0L5 17.7V6a1 1 0 0 1 1-1Zm12 14H6.56l7-5.84a.78.78 0 0 1 .93 0L19 17v1a1 1 0 0 1-1 1Z"/><circle cx="8" cy="8.5" r="1.5"/></g></g></g>`
	imageTwoFillInnerSVG                  = `<g id="evaImage2Fill0"><g id="evaImage2Fill1"><path id="evaImage2Fill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3ZM8 7a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 8 7Zm11 10.83A1.09 1.09 0 0 1 18 19H6l7.57-6.82a.69.69 0 0 1 .93 0l4.5 4.48Z"/></g></g>`
	inboxFillInnerSVG                     = `<g id="evaInboxFill0"><g id="evaInboxFill1"><path id="evaInboxFill2" fill="currentColor" d="m20.79 11.34l-3.34-6.68A3 3 0 0 0 14.76 3H9.24a3 3 0 0 0-2.69 1.66l-3.34 6.68a2 2 0 0 0-.21.9V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5.76a2 2 0 0 0-.21-.9ZM8.34 5.55a1 1 0 0 1 .9-.55h5.52a1 1 0 0 1 .9.55L18.38 11H16a1 1 0 0 0-1 1v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2a1 1 0 0 0-1-1H5.62Z"/></g></g>`
	inboxOutlineInnerSVG                  = `<g id="evaInboxOutline0"><g id="evaInboxOutline1"><path id="evaInboxOutline2" fill="currentColor" d="m20.79 11.34l-3.34-6.68A3 3 0 0 0 14.76 3H9.24a3 3 0 0 0-2.69 1.66l-3.34 6.68a2 2 0 0 0-.21.9V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5.76a2 2 0 0 0-.21-.9ZM8.34 5.55a1 1 0 0 1 .9-.55h5.52a1 1 0 0 1 .9.55L18.38 11H16a1 1 0 0 0-1 1v3H9v-3a1 1 0 0 0-1-1H5.62ZM18 19H6a1 1 0 0 1-1-1v-5h2v3a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-3h2v5a1 1 0 0 1-1 1Z"/></g></g>`
	infoFillInnerSVG                      = `<g id="evaInfoFill0"><g id="evaInfoFill1"><path id="evaInfoFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm1 14a1 1 0 0 1-2 0v-5a1 1 0 0 1 2 0Zm-1-7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`
	infoOutlineInnerSVG                   = `<g id="evaInfoOutline0"><g id="evaInfoOutline1"><g id="evaInfoOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><circle cx="12" cy="8" r="1"/><path d="M12 10a1 1 0 0 0-1 1v5a1 1 0 0 0 2 0v-5a1 1 0 0 0-1-1Z"/></g></g></g>`
	keypadFillInnerSVG                    = `<g id="evaKeypadFill0"><g id="evaKeypadFill1"><path id="evaKeypadFill2" fill="currentColor" d="M5 2a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 6a3 3 0 1 0-3-3a3 3 0 0 0 3 3ZM5 9a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3ZM5 16a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Z"/></g></g>`
	keypadOutlineInnerSVG                 = `<g id="evaKeypadOutline0"><g id="evaKeypadOutline1"><path id="evaKeypadOutline2" fill="currentColor" d="M5 2a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7-4a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7 2a3 3 0 1 0-3-3a3 3 0 0 0 3 3Zm0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1ZM5 9a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7-4a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7-4a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1ZM5 16a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7-4a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7-4a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`
	layersFillInnerSVG                    = `<g id="evaLayersFill0"><g id="evaLayersFill1"><g id="evaLayersFill2" fill="currentColor"><path d="m3.24 7.29l8.52 4.63a.51.51 0 0 0 .48 0l8.52-4.63a.44.44 0 0 0-.05-.81L12.19 3a.5.5 0 0 0-.38 0L3.29 6.48a.44.44 0 0 0-.05.81Z"/><path d="m20.71 10.66l-1.83-.78l-6.64 3.61a.51.51 0 0 1-.48 0L5.12 9.88l-1.83.78a.48.48 0 0 0 0 .85l8.52 4.9a.46.46 0 0 0 .48 0l8.52-4.9a.48.48 0 0 0-.1-.85Z"/><path d="m20.71 15.1l-1.56-.68l-6.91 3.76a.51.51 0 0 1-.48 0l-6.91-3.76l-1.56.68a.49.49 0 0 0 0 .87l8.52 5a.51.51 0 0 0 .48 0l8.52-5a.49.49 0 0 0-.1-.87Z"/></g></g></g>`
	layersOutlineInnerSVG                 = `<g id="evaLayersOutline0"><g id="evaLayersOutline1"><path id="evaLayersOutline2" fill="currentColor" d="M21 11.35a1 1 0 0 0-.61-.86l-2.15-.92l2.26-1.3a1 1 0 0 0 .5-.92a1 1 0 0 0-.61-.86l-8-3.41a1 1 0 0 0-.78 0l-8 3.41a1 1 0 0 0-.61.86a1 1 0 0 0 .5.92l2.26 1.3l-2.15.92a1 1 0 0 0-.61.86a1 1 0 0 0 .5.92l2.26 1.3l-2.15.92a1 1 0 0 0-.61.86a1 1 0 0 0 .5.92l8 4.6a1 1 0 0 0 1 0l8-4.6a1 1 0 0 0 .5-.92a1 1 0 0 0-.61-.86l-2.15-.92l2.26-1.3a1 1 0 0 0 .5-.92Zm-9-6.26l5.76 2.45L12 10.85L6.24 7.54Zm-.5 7.78a1 1 0 0 0 1 0l3.57-2l1.69.72L12 14.85l-5.76-3.31l1.69-.72Zm6.26 2.67L12 18.85l-5.76-3.31l1.69-.72l3.57 2.05a1 1 0 0 0 1 0l3.57-2.05Z"/></g></g>`
	layoutFillInnerSVG                    = `<g id="evaLayoutFill0"><g id="evaLayoutFill1"><path id="evaLayoutFill2" fill="currentColor" d="M21 8V6a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3v2ZM3 10v8a3 3 0 0 0 3 3h5V10Zm10 0v11h5a3 3 0 0 0 3-3v-8Z"/></g></g>`
	layoutOutlineInnerSVG                 = `<g id="evaLayoutOutline0"><g id="evaLayoutOutline1"><path id="evaLayoutOutline2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3ZM6 5h12a1 1 0 0 1 1 1v2H5V6a1 1 0 0 1 1-1ZM5 18v-8h6v9H6a1 1 0 0 1-1-1Zm13 1h-5v-9h6v8a1 1 0 0 1-1 1Z"/></g></g>`
	linkFillInnerSVG                      = `<g id="evaLinkFill0"><g id="evaLinkFill1"><g id="evaLinkFill2" fill="currentColor"><path d="M8 12a1 1 0 0 0 1 1h6a1 1 0 0 0 0-2H9a1 1 0 0 0-1 1Z"/><path d="M9 16H7.21A4.13 4.13 0 0 1 3 12.37A4 4 0 0 1 7 8h2a1 1 0 0 0 0-2H7.21a6.15 6.15 0 0 0-6.16 5.21A6 6 0 0 0 7 18h2a1 1 0 0 0 0-2Zm14-4.76A6.16 6.16 0 0 0 16.76 6h-1.51C14.44 6 14 6.45 14 7a1 1 0 0 0 1 1h1.79A4.13 4.13 0 0 1 21 11.63A4 4 0 0 1 17 16h-2a1 1 0 0 0 0 2h2a6 6 0 0 0 6-6.76Z"/></g></g></g>`
	linkOutlineInnerSVG                   = `<g id="evaLinkOutline0"><g id="evaLinkOutline1"><g id="evaLinkOutline2" fill="currentColor"><path d="M8 12a1 1 0 0 0 1 1h6a1 1 0 0 0 0-2H9a1 1 0 0 0-1 1Z"/><path d="M9 16H7.21A4.13 4.13 0 0 1 3 12.37A4 4 0 0 1 7 8h2a1 1 0 0 0 0-2H7.21a6.15 6.15 0 0 0-6.16 5.21A6 6 0 0 0 7 18h2a1 1 0 0 0 0-2Zm14-4.76A6.16 6.16 0 0 0 16.76 6h-1.51C14.44 6 14 6.45 14 7a1 1 0 0 0 1 1h1.79A4.13 4.13 0 0 1 21 11.63A4 4 0 0 1 17 16h-2a1 1 0 0 0 0 2h2a6 6 0 0 0 6-6.76Z"/></g></g></g>`
	linkTwoFillInnerSVG                   = `<g id="evaLink2Fill0"><g id="evaLink2Fill1"><g id="evaLink2Fill2" fill="currentColor"><path d="m13.29 9.29l-4 4a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0-1.42-1.42Z"/><path d="M12.28 17.4L11 18.67a4.2 4.2 0 0 1-5.58.4a4 4 0 0 1-.27-5.93l1.42-1.43a1 1 0 0 0 0-1.42a1 1 0 0 0-1.42 0l-1.27 1.28a6.15 6.15 0 0 0-.67 8.07a6.06 6.06 0 0 0 9.07.6l1.42-1.42a1 1 0 0 0-1.42-1.42Zm7.38-14.18a6.18 6.18 0 0 0-8.13.68L10.45 5a1.09 1.09 0 0 0-.17 1.61a1 1 0 0 0 1.42 0L13 5.3a4.17 4.17 0 0 1 5.57-.4a4 4 0 0 1 .27 5.95l-1.42 1.43a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.42-1.42a6.06 6.06 0 0 0-.6-9.06Z"/></g></g></g>`
	linkTwoOutlineInnerSVG                = `<g id="evaLink2Outline0"><g id="evaLink2Outline1"><g id="evaLink2Outline2" fill="currentColor"><path d="m13.29 9.29l-4 4a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0-1.42-1.42Z"/><path d="M12.28 17.4L11 18.67a4.2 4.2 0 0 1-5.58.4a4 4 0 0 1-.27-5.93l1.42-1.43a1 1 0 0 0 0-1.42a1 1 0 0 0-1.42 0l-1.27 1.28a6.15 6.15 0 0 0-.67 8.07a6.06 6.06 0 0 0 9.07.6l1.42-1.42a1 1 0 0 0-1.42-1.42Zm7.38-14.18a6.18 6.18 0 0 0-8.13.68L10.45 5a1.09 1.09 0 0 0-.17 1.61a1 1 0 0 0 1.42 0L13 5.3a4.17 4.17 0 0 1 5.57-.4a4 4 0 0 1 .27 5.95l-1.42 1.43a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.42-1.42a6.06 6.06 0 0 0-.6-9.06Z"/></g></g></g>`
	linkedinFillInnerSVG                  = `<g id="evaLinkedinFill0"><g id="evaLinkedinFill1"><g id="evaLinkedinFill2" fill="currentColor"><path d="M15.15 8.4a5.83 5.83 0 0 0-5.85 5.82v5.88a.9.9 0 0 0 .9.9h2.1a.9.9 0 0 0 .9-.9v-5.88a1.94 1.94 0 0 1 2.15-1.93a2 2 0 0 1 1.75 2v5.81a.9.9 0 0 0 .9.9h2.1a.9.9 0 0 0 .9-.9v-5.88a5.83 5.83 0 0 0-5.85-5.82Z"/><rect width="4.5" height="11.7" x="3" y="9.3" rx=".9" ry=".9"/><circle cx="5.25" cy="5.25" r="2.25"/></g></g></g>`
	linkedinOutlineInnerSVG               = `<g id="evaLinkedinOutline0"><g id="evaLinkedinOutline1"><path id="evaLinkedinOutline2" fill="currentColor" d="M20 22h-1.67a2 2 0 0 1-2-2v-5.37a.92.92 0 0 0-.69-.93a.84.84 0 0 0-.67.19a.85.85 0 0 0-.3.65V20a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2v-5.46a6.5 6.5 0 1 1 13 0V20a2 2 0 0 1-2 2Zm-4.5-10.31a3.73 3.73 0 0 1 .47 0a2.91 2.91 0 0 1 2.36 2.9V20H20v-5.46a4.5 4.5 0 1 0-9 0V20h1.67v-5.46a2.85 2.85 0 0 1 2.83-2.85ZM6 22H4a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2ZM4 10v10h2V10Zm1-3a3 3 0 1 1 3-3a3 3 0 0 1-3 3Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/></g></g>`
	listFillInnerSVG                      = `<g id="evaListFill0"><g id="evaListFill1"><g id="evaListFill2" fill="currentColor"><circle cx="4" cy="7" r="1"/><circle cx="4" cy="12" r="1"/><circle cx="4" cy="17" r="1"/><rect width="14" height="2" x="7" y="11" rx=".94" ry=".94"/><rect width="14" height="2" x="7" y="16" rx=".94" ry=".94"/><rect width="14" height="2" x="7" y="6" rx=".94" ry=".94"/></g></g></g>`
	listOutlineInnerSVG                   = `<g id="evaListOutline0"><g id="evaListOutline1"><g id="evaListOutline2" fill="currentColor"><circle cx="4" cy="7" r="1"/><circle cx="4" cy="12" r="1"/><circle cx="4" cy="17" r="1"/><rect width="14" height="2" x="7" y="11" rx=".94" ry=".94"/><rect width="14" height="2" x="7" y="16" rx=".94" ry=".94"/><rect width="14" height="2" x="7" y="6" rx=".94" ry=".94"/></g></g></g>`
	loaderOutlineInnerSVG                 = `<g id="evaLoaderOutline0"><g id="evaLoaderOutline1"><path id="evaLoaderOutline2" fill="currentColor" d="M12 2a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1Zm9 9h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2ZM6 12a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1Zm.22-7a1 1 0 0 0-1.39 1.47l1.44 1.39a1 1 0 0 0 .73.28a1 1 0 0 0 .72-.31a1 1 0 0 0 0-1.41ZM17 8.14a1 1 0 0 0 .69-.28l1.44-1.39A1 1 0 0 0 17.78 5l-1.44 1.42a1 1 0 0 0 0 1.41a1 1 0 0 0 .66.31ZM12 18a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1Zm5.73-1.86a1 1 0 0 0-1.39 1.44L17.78 19a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.3a1 1 0 0 0 0-1.42Zm-11.46 0l-1.44 1.39a1 1 0 0 0 0 1.42a1 1 0 0 0 .72.3a1 1 0 0 0 .67-.25l1.44-1.39a1 1 0 0 0-1.39-1.44Z"/></g></g>`
	lockFillInnerSVG                      = `<g id="evaLockFill0"><g id="evaLockFill1"><g id="evaLockFill2" fill="currentColor"><circle cx="12" cy="15" r="1"/><path d="M17 8h-1V6.11a4 4 0 1 0-8 0V8H7a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-7-1.89A2.06 2.06 0 0 1 12 4a2.06 2.06 0 0 1 2 2.11V8h-4ZM12 18a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/></g></g></g>`
	lockOutlineInnerSVG                   = `<g id="evaLockOutline0"><g id="evaLockOutline1"><g id="evaLockOutline2" fill="currentColor"><path d="M17 8h-1V6.11a4 4 0 1 0-8 0V8H7a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-7-1.89A2.06 2.06 0 0 1 12 4a2.06 2.06 0 0 1 2 2.11V8h-4ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/><path d="M12 12a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g></g>`
	logInFillInnerSVG                     = `<g id="evaLogInFill0"><g id="evaLogInFill1"><path id="evaLogInFill2" fill="currentColor" d="M19 4h-2a1 1 0 0 0 0 2h1v12h-1a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Zm-7.2 3.4a1 1 0 0 0-1.6 1.2L12 11H4a1 1 0 0 0 0 2h8.09l-1.72 2.44a1 1 0 0 0 .24 1.4a1 1 0 0 0 .58.18a1 1 0 0 0 .81-.42l2.82-4a1 1 0 0 0 0-1.18Z"/></g></g>`
	logInOutlineInnerSVG                  = `<g id="evaLogInOutline0"><g id="evaLogInOutline1"><path id="evaLogInOutline2" fill="currentColor" d="M19 4h-2a1 1 0 0 0 0 2h1v12h-1a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Zm-7.2 3.4a1 1 0 0 0-1.6 1.2L12 11H4a1 1 0 0 0 0 2h8.09l-1.72 2.44a1 1 0 0 0 .24 1.4a1 1 0 0 0 .58.18a1 1 0 0 0 .81-.42l2.82-4a1 1 0 0 0 0-1.18Z"/></g></g>`
	logOutFillInnerSVG                    = `<g id="evaLogOutFill0"><g id="evaLogOutFill1"><path id="evaLogOutFill2" fill="currentColor" d="M7 6a1 1 0 0 0 0-2H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2H6V6Zm13.82 5.42l-2.82-4a1 1 0 0 0-1.39-.24a1 1 0 0 0-.24 1.4L18.09 11H10a1 1 0 0 0 0 2h8l-1.8 2.4a1 1 0 0 0 .2 1.4a1 1 0 0 0 .6.2a1 1 0 0 0 .8-.4l3-4a1 1 0 0 0 .02-1.18Z"/></g></g>`
	logOutOutlineInnerSVG                 = `<g id="evaLogOutOutline0"><g id="evaLogOutOutline1"><path id="evaLogOutOutline2" fill="currentColor" d="M7 6a1 1 0 0 0 0-2H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2H6V6Zm13.82 5.42l-2.82-4a1 1 0 0 0-1.39-.24a1 1 0 0 0-.24 1.4L18.09 11H10a1 1 0 0 0 0 2h8l-1.8 2.4a1 1 0 0 0 .2 1.4a1 1 0 0 0 .6.2a1 1 0 0 0 .8-.4l3-4a1 1 0 0 0 .02-1.18Z"/></g></g>`
	mapFillInnerSVG                       = `<g id="evaMapFill0"><g id="evaMapFill1"><path id="evaMapFill2" fill="currentColor" d="m20.41 5.89l-4-1.8h-.82L12 5.7L8.41 4.09h-.05L8.24 4h-.6l-4 1.8a1 1 0 0 0-.64 1V19a1 1 0 0 0 .46.84A1 1 0 0 0 4 20a1 1 0 0 0 .41-.09L8 18.3l3.59 1.61h.05a.85.85 0 0 0 .72 0h.05L16 18.3l3.59 1.61A1 1 0 0 0 20 20a1 1 0 0 0 .54-.16A1 1 0 0 0 21 19V6.8a1 1 0 0 0-.59-.91ZM9 6.55l2 .89v10l-2-.89Zm10 10.9l-2-.89v-10l2 .89Z"/></g></g>`
	mapOutlineInnerSVG                    = `<g id="evaMapOutline0"><g id="evaMapOutline1"><path id="evaMapOutline2" fill="currentColor" d="m20.41 5.89l-4-1.8h-.82L12 5.7L8.41 4.09h-.05L8.24 4h-.6l-4 1.8a1 1 0 0 0-.64 1V19a1 1 0 0 0 .46.84A1 1 0 0 0 4 20a1 1 0 0 0 .41-.09L8 18.3l3.59 1.61h.05a.85.85 0 0 0 .72 0h.05L16 18.3l3.59 1.61A1 1 0 0 0 20 20a1 1 0 0 0 .54-.16A1 1 0 0 0 21 19V6.8a1 1 0 0 0-.59-.91ZM5 7.44l2-.89v10l-2 .89Zm4-.89l2 .89v10l-2-.89Zm4 .89l2-.89v10l-2 .89Zm6 10l-2-.89v-10l2 .89Z"/></g></g>`
	maximizeFillInnerSVG                  = `<g id="evaMaximizeFill0"><g id="evaMaximizeFill1"><path id="evaMaximizeFill2" fill="currentColor" d="m20.71 19.29l-3.4-3.39A7.92 7.92 0 0 0 19 11a8 8 0 1 0-8 8a7.92 7.92 0 0 0 4.9-1.69l3.39 3.4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM13 12h-1v1a1 1 0 0 1-2 0v-1H9a1 1 0 0 1 0-2h1V9a1 1 0 0 1 2 0v1h1a1 1 0 0 1 0 2Z"/></g></g>`
	maximizeOutlineInnerSVG               = `<g id="evaMaximizeOutline0"><g id="evaMaximizeOutline1"><g id="evaMaximizeOutline2" fill="currentColor"><path d="m20.71 19.29l-3.4-3.39A7.92 7.92 0 0 0 19 11a8 8 0 1 0-8 8a7.92 7.92 0 0 0 4.9-1.69l3.39 3.4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM5 11a6 6 0 1 1 6 6a6 6 0 0 1-6-6Z"/><path d="M13 10h-1V9a1 1 0 0 0-2 0v1H9a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2Z"/></g></g></g>`
	menuArrowFillInnerSVG                 = `<g id="evaMenuArrowFill0"><g id="evaMenuArrowFill1"><g id="evaMenuArrowFill2" fill="currentColor"><path d="M20.05 11H5.91l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0 0 1.42l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.91 13h14.14a1 1 0 0 0 .95-.95V12a1 1 0 0 0-.95-1Z"/><rect width="18" height="2" x="3" y="17" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="5" rx=".95" ry=".95"/></g></g></g>`
	menuArrowOutlineInnerSVG              = `<g id="evaMenuArrowOutline0"><g id="evaMenuArrowOutline1"><g id="evaMenuArrowOutline2" fill="currentColor"><path d="M20.05 11H5.91l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0 0 1.42l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.91 13h14.14a1 1 0 0 0 .95-.95V12a1 1 0 0 0-.95-1Z"/><rect width="18" height="2" x="3" y="17" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="5" rx=".95" ry=".95"/></g></g></g>`
	menuFillInnerSVG                      = `<g id="evaMenuFill0"><g id="evaMenuFill1"><g id="evaMenuFill2" fill="currentColor"><rect width="18" height="2" x="3" y="11" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="16" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="6" rx=".95" ry=".95"/></g></g></g>`
	menuOutlineInnerSVG                   = `<g id="evaMenuOutline0"><g id="evaMenuOutline1"><g id="evaMenuOutline2" fill="currentColor"><rect width="18" height="2" x="3" y="11" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="16" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="6" rx=".95" ry=".95"/></g></g></g>`
	menuTwoFillInnerSVG                   = `<g id="evaMenu2Fill0"><g id="evaMenu2Fill1"><g id="evaMenu2Fill2" fill="currentColor"><circle cx="4" cy="12" r="1"/><rect width="14" height="2" x="7" y="11" rx=".94" ry=".94"/><rect width="18" height="2" x="3" y="16" rx=".94" ry=".94"/><rect width="18" height="2" x="3" y="6" rx=".94" ry=".94"/></g></g></g>`
	menuTwoOutlineInnerSVG                = `<g id="evaMenu2Outline0"><g id="evaMenu2Outline1"><g id="evaMenu2Outline2" fill="currentColor"><circle cx="4" cy="12" r="1"/><rect width="14" height="2" x="7" y="11" rx=".94" ry=".94"/><rect width="18" height="2" x="3" y="16" rx=".94" ry=".94"/><rect width="18" height="2" x="3" y="6" rx=".94" ry=".94"/></g></g></g>`
	messageCircleFillInnerSVG             = `<g id="evaMessageCircleFill0"><g id="evaMessageCircleFill1"><path id="evaMessageCircleFill2" fill="currentColor" d="M19.07 4.93a10 10 0 0 0-16.28 11a1.06 1.06 0 0 1 .09.64L2 20.8a1 1 0 0 0 .27.91A1 1 0 0 0 3 22h.2l4.28-.86a1.26 1.26 0 0 1 .64.09a10 10 0 0 0 11-16.28ZM8 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`
	messageCircleOutlineInnerSVG          = `<g id="evaMessageCircleOutline0"><g id="evaMessageCircleOutline1"><g id="evaMessageCircleOutline2" fill="currentColor"><circle cx="12" cy="12" r="1"/><circle cx="16" cy="12" r="1"/><circle cx="8" cy="12" r="1"/><path d="M19.07 4.93a10 10 0 0 0-16.28 11a1.06 1.06 0 0 1 .09.64L2 20.8a1 1 0 0 0 .27.91A1 1 0 0 0 3 22h.2l4.28-.86a1.26 1.26 0 0 1 .64.09a10 10 0 0 0 11-16.28Zm.83 8.36a8 8 0 0 1-11 6.08a3.26 3.26 0 0 0-1.25-.26a3.43 3.43 0 0 0-.56.05l-2.82.57l.57-2.82a3.09 3.09 0 0 0-.21-1.81a8 8 0 0 1 6.08-11a8 8 0 0 1 9.19 9.19Z"/></g></g></g>`
	messageSquareFillInnerSVG             = `<g id="evaMessageSquareFill0"><g id="evaMessageSquareFill1"><path id="evaMessageSquareFill2" fill="currentColor" d="M19 3H5a3 3 0 0 0-3 3v15a1 1 0 0 0 .51.87A1 1 0 0 0 3 22a1 1 0 0 0 .51-.14L8 19.14a1 1 0 0 1 .55-.14H19a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3ZM8 12a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`
	messageSquareOutlineInnerSVG          = `<g id="evaMessageSquareOutline0"><g id="evaMessageSquareOutline1"><g id="evaMessageSquareOutline2" fill="currentColor"><circle cx="12" cy="11" r="1"/><circle cx="16" cy="11" r="1"/><circle cx="8" cy="11" r="1"/><path d="M19 3H5a3 3 0 0 0-3 3v15a1 1 0 0 0 .51.87A1 1 0 0 0 3 22a1 1 0 0 0 .51-.14L8 19.14a1 1 0 0 1 .55-.14H19a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm1 13a1 1 0 0 1-1 1H8.55a3 3 0 0 0-1.55.43l-3 1.8V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/></g></g></g>`
	micFillInnerSVG                       = `<g id="evaMicFill0"><g id="evaMicFill1"><g id="evaMicFill2" fill="currentColor"><path d="M12 15a4 4 0 0 0 4-4V6a4 4 0 0 0-8 0v5a4 4 0 0 0 4 4Z"/><path d="M19 11a1 1 0 0 0-2 0a5 5 0 0 1-10 0a1 1 0 0 0-2 0a7 7 0 0 0 6 6.92V20H8.89a.89.89 0 0 0-.89.89v.22a.89.89 0 0 0 .89.89h6.22a.89.89 0 0 0 .89-.89v-.22a.89.89 0 0 0-.89-.89H13v-2.08A7 7 0 0 0 19 11Z"/></g></g></g>`
	micOffFillInnerSVG                    = `<g id="evaMicOffFill0"><g id="evaMicOffFill1"><g id="evaMicOffFill2" fill="currentColor"><path d="M15.58 12.75A4 4 0 0 0 16 11V6a4 4 0 0 0-7.92-.75M19 11a1 1 0 0 0-2 0a4.86 4.86 0 0 1-.69 2.48L17.78 15A7 7 0 0 0 19 11Zm-7 4h.16L8 10.83V11a4 4 0 0 0 4 4Zm8.71 4.29l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/><path d="M15 20h-2v-2.08a7 7 0 0 0 1.65-.44l-1.6-1.6A4.57 4.57 0 0 1 12 16a5 5 0 0 1-5-5a1 1 0 0 0-2 0a7 7 0 0 0 6 6.92V20H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z"/></g></g></g>`
	micOffOutlineInnerSVG                 = `<g id="evaMicOffOutline0"><g id="evaMicOffOutline1"><g id="evaMicOffOutline2" fill="currentColor"><path d="M10 6a2 2 0 0 1 4 0v5a1 1 0 0 1 0 .16l1.6 1.59A4 4 0 0 0 16 11V6a4 4 0 0 0-7.92-.75L10 7.17Zm9 5a1 1 0 0 0-2 0a4.86 4.86 0 0 1-.69 2.48L17.78 15A7 7 0 0 0 19 11Zm-7 4h.16L8 10.83V11a4 4 0 0 0 4 4Zm8.71 4.29l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/><path d="M15 20h-2v-2.08a7 7 0 0 0 1.65-.44l-1.6-1.6A4.57 4.57 0 0 1 12 16a5 5 0 0 1-5-5a1 1 0 0 0-2 0a7 7 0 0 0 6 6.92V20H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z"/></g></g></g>`
	micOutlineInnerSVG                    = `<g id="evaMicOutline0"><g id="evaMicOutline1"><g id="evaMicOutline2" fill="currentColor"><path d="M12 15a4 4 0 0 0 4-4V6a4 4 0 0 0-8 0v5a4 4 0 0 0 4 4Zm-2-9a2 2 0 0 1 4 0v5a2 2 0 0 1-4 0Z"/><path d="M19 11a1 1 0 0 0-2 0a5 5 0 0 1-10 0a1 1 0 0 0-2 0a7 7 0 0 0 6 6.92V20H8.89a.89.89 0 0 0-.89.89v.22a.89.89 0 0 0 .89.89h6.22a.89.89 0 0 0 .89-.89v-.22a.89.89 0 0 0-.89-.89H13v-2.08A7 7 0 0 0 19 11Z"/></g></g></g>`
	minimizeFillInnerSVG                  = `<g id="evaMinimizeFill0"><g id="evaMinimizeFill1"><path id="evaMinimizeFill2" fill="currentColor" d="m20.71 19.29l-3.4-3.39A7.92 7.92 0 0 0 19 11a8 8 0 1 0-8 8a7.92 7.92 0 0 0 4.9-1.69l3.39 3.4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM13 12H9a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Z"/></g></g>`
	minimizeOutlineInnerSVG               = `<g id="evaMinimizeOutline0"><g id="evaMinimizeOutline1"><g id="evaMinimizeOutline2" fill="currentColor"><path d="m20.71 19.29l-3.4-3.39A7.92 7.92 0 0 0 19 11a8 8 0 1 0-8 8a7.92 7.92 0 0 0 4.9-1.69l3.39 3.4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM5 11a6 6 0 1 1 6 6a6 6 0 0 1-6-6Z"/><path d="M13 10H9a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Z"/></g></g></g>`
	minusCircleFillInnerSVG               = `<g id="evaMinusCircleFill0"><g id="evaMinusCircleFill1"><path id="evaMinusCircleFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm3 11H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Z"/></g></g>`
	minusCircleOutlineInnerSVG            = `<g id="evaMinusCircleOutline0"><g id="evaMinusCircleOutline1"><g id="evaMinusCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M15 11H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z"/></g></g></g>`
	minusFillInnerSVG                     = `<g id="evaMinusFill0"><g id="evaMinusFill1"><path id="evaMinusFill2" fill="currentColor" d="M19 13H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Z"/></g></g>`
	minusOutlineInnerSVG                  = `<g id="evaMinusOutline0"><g id="evaMinusOutline1"><path id="evaMinusOutline2" fill="currentColor" d="M19 13H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Z"/></g></g>`
	minusSquareFillInnerSVG               = `<g id="evaMinusSquareFill0"><g id="evaMinusSquareFill1"><path id="evaMinusSquareFill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm-3 10H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Z"/></g></g>`
	minusSquareOutlineInnerSVG            = `<g id="evaMinusSquareOutline0"><g id="evaMinusSquareOutline1"><g id="evaMinusSquareOutline2" fill="currentColor"><path d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm1 15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/><path d="M15 11H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z"/></g></g></g>`
	monitorFillInnerSVG                   = `<g id="evaMonitorFill0"><g id="evaMonitorFill1"><path id="evaMonitorFill2" fill="currentColor" d="M19 3H5a3 3 0 0 0-3 3v5h20V6a3 3 0 0 0-3-3ZM2 14a3 3 0 0 0 3 3h6v2H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2h-4v-2h6a3 3 0 0 0 3-3v-1H2Z"/></g></g>`
	monitorOutlineInnerSVG                = `<g id="evaMonitorOutline0"><g id="evaMonitorOutline1"><path id="evaMonitorOutline2" fill="currentColor" d="M19 3H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h6v2H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2h-4v-2h6a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm1 11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/></g></g>`
	moonFillInnerSVG                      = `<g id="evaMoonFill0"><g id="evaMoonFill1"><path id="evaMoonFill2" fill="currentColor" d="M12.3 22h-.1a10.31 10.31 0 0 1-7.34-3.15a10.46 10.46 0 0 1-.26-14a10.13 10.13 0 0 1 4-2.74a1 1 0 0 1 1.06.22a1 1 0 0 1 .24 1a8.4 8.4 0 0 0 1.94 8.81a8.47 8.47 0 0 0 8.83 1.94a1 1 0 0 1 1.27 1.29A10.16 10.16 0 0 1 19.6 19a10.28 10.28 0 0 1-7.3 3Z"/></g></g>`
	moonOutlineInnerSVG                   = `<g id="evaMoonOutline0"><g id="evaMoonOutline1"><path id="evaMoonOutline2" fill="currentColor" d="M12.3 22h-.1a10.31 10.31 0 0 1-7.34-3.15a10.46 10.46 0 0 1-.26-14a10.13 10.13 0 0 1 4-2.74a1 1 0 0 1 1.06.22a1 1 0 0 1 .24 1a8.4 8.4 0 0 0 1.94 8.81a8.47 8.47 0 0 0 8.83 1.94a1 1 0 0 1 1.27 1.29A10.16 10.16 0 0 1 19.6 19a10.28 10.28 0 0 1-7.3 3ZM7.46 4.92a7.93 7.93 0 0 0-1.37 1.22a8.44 8.44 0 0 0 .2 11.32A8.29 8.29 0 0 0 12.22 20h.08a8.34 8.34 0 0 0 6.78-3.49A10.37 10.37 0 0 1 7.46 4.92Z"/></g></g>`
	moreHorizontalFillInnerSVG            = `<g id="evaMoreHorizontalFill0"><g id="evaMoreHorizontalFill1"><g id="evaMoreHorizontalFill2" fill="currentColor"><circle cx="12" cy="12" r="2"/><circle cx="19" cy="12" r="2"/><circle cx="5" cy="12" r="2"/></g></g></g>`
	moreHorizontalOutlineInnerSVG         = `<g id="evaMoreHorizontalOutline0"><g id="evaMoreHorizontalOutline1"><g id="evaMoreHorizontalOutline2" fill="currentColor"><circle cx="12" cy="12" r="2"/><circle cx="19" cy="12" r="2"/><circle cx="5" cy="12" r="2"/></g></g></g>`
	moreVerticalFillInnerSVG              = `<g id="evaMoreVerticalFill0"><g id="evaMoreVerticalFill1"><g id="evaMoreVerticalFill2" fill="currentColor"><circle cx="12" cy="12" r="2"/><circle cx="12" cy="5" r="2"/><circle cx="12" cy="19" r="2"/></g></g></g>`
	moreVerticalOutlineInnerSVG           = `<g id="evaMoreVerticalOutline0"><g id="evaMoreVerticalOutline1"><g id="evaMoreVerticalOutline2" fill="currentColor"><circle cx="12" cy="12" r="2"/><circle cx="12" cy="5" r="2"/><circle cx="12" cy="19" r="2"/></g></g></g>`
	moveFillInnerSVG                      = `<g id="evaMoveFill0"><g id="evaMoveFill1"><path id="evaMoveFill2" fill="currentColor" d="m21.71 11.31l-3-3a1 1 0 0 0-1.42 1.42L18.58 11H13V5.41l1.29 1.3A1 1 0 0 0 15 7a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-3-3A1 1 0 0 0 12 2a1 1 0 0 0-.7.29l-3 3a1 1 0 0 0 1.41 1.42L11 5.42V11H5.41l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3A1 1 0 0 0 2 12a1 1 0 0 0 .29.71l3 3A1 1 0 0 0 6 16a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42L5.42 13H11v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 22a1 1 0 0 0 .7-.29l3-3a1 1 0 0 0-1.42-1.42L13 18.58V13h5.59l-1.3 1.29a1 1 0 0 0 0 1.42A1 1 0 0 0 18 16a1 1 0 0 0 .71-.29l3-3A1 1 0 0 0 22 12a1 1 0 0 0-.29-.69Z"/></g></g>`
	moveOutlineInnerSVG                   = `<g id="evaMoveOutline0"><g id="evaMoveOutline1"><path id="evaMoveOutline2" fill="currentColor" d="m21.71 11.31l-3-3a1 1 0 0 0-1.42 1.42L18.58 11H13V5.41l1.29 1.3A1 1 0 0 0 15 7a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-3-3A1 1 0 0 0 12 2a1 1 0 0 0-.7.29l-3 3a1 1 0 0 0 1.41 1.42L11 5.42V11H5.41l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3A1 1 0 0 0 2 12a1 1 0 0 0 .29.71l3 3A1 1 0 0 0 6 16a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42L5.42 13H11v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 22a1 1 0 0 0 .7-.29l3-3a1 1 0 0 0-1.42-1.42L13 18.58V13h5.59l-1.3 1.29a1 1 0 0 0 0 1.42A1 1 0 0 0 18 16a1 1 0 0 0 .71-.29l3-3A1 1 0 0 0 22 12a1 1 0 0 0-.29-.69Z"/></g></g>`
	musicFillInnerSVG                     = `<g id="evaMusicFill0"><g id="evaMusicFill1"><path id="evaMusicFill2" fill="currentColor" d="M19 15V4a1 1 0 0 0-.38-.78a1 1 0 0 0-.84-.2l-9 2A1 1 0 0 0 8 6v8.34a3.49 3.49 0 1 0 2 3.18a4.36 4.36 0 0 0 0-.52V6.8l7-1.55v7.09a3.49 3.49 0 1 0 2 3.17a4.57 4.57 0 0 0 0-.51Z"/></g></g>`
	musicOutlineInnerSVG                  = `<g id="evaMusicOutline0"><g id="evaMusicOutline1"><path id="evaMusicOutline2" fill="currentColor" d="M19 15V4a1 1 0 0 0-.38-.78a1 1 0 0 0-.84-.2l-9 2A1 1 0 0 0 8 6v8.34a3.49 3.49 0 1 0 2 3.18a4.36 4.36 0 0 0 0-.52V6.8l7-1.55v7.09a3.49 3.49 0 1 0 2 3.17a4.57 4.57 0 0 0 0-.51ZM6.54 19A1.49 1.49 0 1 1 8 17.21a1.53 1.53 0 0 1 0 .3A1.49 1.49 0 0 1 6.54 19Zm9-2A1.5 1.5 0 1 1 17 15.21a1.53 1.53 0 0 1 0 .3A1.5 1.5 0 0 1 15.51 17Z"/></g></g>`
	navigationFillInnerSVG                = `<g id="evaNavigationFill0"><g id="evaNavigationFill1"><path id="evaNavigationFill2" fill="currentColor" d="M20 20a.94.94 0 0 1-.55-.17l-6.9-4.56a1 1 0 0 0-1.1 0l-6.9 4.56a1 1 0 0 1-1.44-1.28l8-16a1 1 0 0 1 1.78 0l8 16a1 1 0 0 1-.23 1.2A1 1 0 0 1 20 20Z"/></g></g>`
	navigationOutlineInnerSVG             = `<g id="evaNavigationOutline0"><g id="evaNavigationOutline1"><path id="evaNavigationOutline2" fill="currentColor" d="M20 20a.94.94 0 0 1-.55-.17L12 14.9l-7.45 4.93a1 1 0 0 1-1.44-1.28l8-16a1 1 0 0 1 1.78 0l8 16a1 1 0 0 1-.23 1.2A1 1 0 0 1 20 20Zm-8-7.3a1 1 0 0 1 .55.17l4.88 3.23L12 5.24L6.57 16.1l4.88-3.23a1 1 0 0 1 .55-.17Z"/></g></g>`
	navigationTwoFillInnerSVG             = `<g id="evaNavigation2Fill0"><g id="evaNavigation2Fill1"><path id="evaNavigation2Fill2" fill="currentColor" d="M13.67 22h-.06a1 1 0 0 1-.92-.8l-1.54-7.57a1 1 0 0 0-.78-.78L2.8 11.31a1 1 0 0 1-.12-1.93l16-5.33A1 1 0 0 1 20 5.32l-5.33 16a1 1 0 0 1-1 .68Z"/></g></g>`
	navigationTwoOutlineInnerSVG          = `<g id="evaNavigation2Outline0"><g id="evaNavigation2Outline1"><path id="evaNavigation2Outline2" fill="currentColor" d="M13.67 22h-.06a1 1 0 0 1-.92-.8L11 13l-8.2-1.69a1 1 0 0 1-.12-1.93l16-5.33A1 1 0 0 1 20 5.32l-5.33 16a1 1 0 0 1-1 .68Zm-6.8-11.9l5.19 1.06a1 1 0 0 1 .79.78l1.05 5.19l3.52-10.55Z"/></g></g>`
	npmFillInnerSVG                       = `<g id="evaNpmFill0"><g id="evaNpmFill1"><path id="evaNpmFill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h7V11h4v10h1a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Z"/></g></g>`
	npmOutlineInnerSVG                    = `<g id="evaNpmOutline0"><g id="evaNpmOutline1"><g id="evaNpmOutline2" fill="currentColor"><path d="M18 21H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3ZM6 5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1Z"/><path d="M12 9h4v10h-4z"/></g></g></g>`
	optionsFillInnerSVG                   = `<g id="evaOptionsFill0"><g id="evaOptionsFill1"><path id="evaOptionsFill2" fill="currentColor" d="M7 14.18V3a1 1 0 0 0-2 0v11.18a3 3 0 0 0 0 5.64V21a1 1 0 0 0 2 0v-1.18a3 3 0 0 0 0-5.64ZM21 13a3 3 0 0 0-2-2.82V3a1 1 0 0 0-2 0v7.18a3 3 0 0 0 0 5.64V21a1 1 0 0 0 2 0v-5.18A3 3 0 0 0 21 13Zm-6-8a3 3 0 1 0-4 2.82V21a1 1 0 0 0 2 0V7.82A3 3 0 0 0 15 5Z"/></g></g>`
	optionsOutlineInnerSVG                = `<g id="evaOptionsOutline0"><g id="evaOptionsOutline1"><path id="evaOptionsOutline2" fill="currentColor" d="M7 14.18V3a1 1 0 0 0-2 0v11.18a3 3 0 0 0 0 5.64V21a1 1 0 0 0 2 0v-1.18a3 3 0 0 0 0-5.64ZM6 18a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm15-5a3 3 0 0 0-2-2.82V3a1 1 0 0 0-2 0v7.18a3 3 0 0 0 0 5.64V21a1 1 0 0 0 2 0v-5.18A3 3 0 0 0 21 13Zm-3 1a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm-3-9a3 3 0 1 0-4 2.82V21a1 1 0 0 0 2 0V7.82A3 3 0 0 0 15 5Zm-3 1a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`
	optionsTwoFillInnerSVG                = `<g id="evaOptions2Fill0"><g id="evaOptions2Fill1"><path id="evaOptions2Fill2" fill="currentColor" d="M19 9a3 3 0 0 0-2.82 2H3a1 1 0 0 0 0 2h13.18A3 3 0 1 0 19 9ZM3 7h1.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2H9.82a3 3 0 0 0-5.64 0H3a1 1 0 0 0 0 2Zm18 10h-7.18a3 3 0 0 0-5.64 0H3a1 1 0 0 0 0 2h5.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2Z"/></g></g>`
	optionsTwoOutlineInnerSVG             = `<g id="evaOptions2Outline0"><g id="evaOptions2Outline1"><path id="evaOptions2Outline2" fill="currentColor" d="M19 9a3 3 0 0 0-2.82 2H3a1 1 0 0 0 0 2h13.18A3 3 0 1 0 19 9Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1ZM3 7h1.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2H9.82a3 3 0 0 0-5.64 0H3a1 1 0 0 0 0 2Zm4-2a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm14 12h-7.18a3 3 0 0 0-5.64 0H3a1 1 0 0 0 0 2h5.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2Zm-10 2a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`
	pantoneFillInnerSVG                   = `<g id="evaPantoneFill0"><g id="evaPantoneFill1"><path id="evaPantoneFill2" fill="currentColor" d="M20 13.18h-2.7l-1.86 2L11.88 19l-1.41 1.52L10 21h10a1 1 0 0 0 1-1v-5.82a1 1 0 0 0-1-1ZM18.19 9.3l-4.14-3.86a.89.89 0 0 0-.71-.26a1 1 0 0 0-.7.31l-.82.89v10.71a5.23 5.23 0 0 1-.06.57l6.48-6.95a1 1 0 0 0-.05-1.41ZM10.82 4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v13.09a3.91 3.91 0 0 0 7.82 0Zm-2 13.09a1.91 1.91 0 0 1-3.82 0V15h3.82Zm0-4.09H5v-3h3.82Zm0-5H5V5h3.82Z"/></g></g>`
	pantoneOutlineInnerSVG                = `<g id="evaPantoneOutline0"><g id="evaPantoneOutline1"><path id="evaPantoneOutline2" fill="currentColor" d="M20 13.18h-4.06l2.3-2.47a1 1 0 0 0 0-1.41l-4.19-3.86a.93.93 0 0 0-.71-.26a1 1 0 0 0-.7.31l-1.82 2V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v13.09A3.91 3.91 0 0 0 6.91 21H20a1 1 0 0 0 1-1v-5.82a1 1 0 0 0-1-1Zm-6.58-5.59l2.67 2.49l-5.27 5.66v-5.36ZM8.82 10v3H5v-3Zm0-5v3H5V5ZM5 17.09V15h3.82v2.09a1.91 1.91 0 0 1-3.82 0ZM19 19h-8.49l3.56-3.82H19Z"/></g></g>`
	paperPlaneFillInnerSVG                = `<g id="evaPaperPlaneFill0"><g id="evaPaperPlaneFill1"><path id="evaPaperPlaneFill2" fill="currentColor" d="M21 4a1.31 1.31 0 0 0-.06-.27v-.09a1 1 0 0 0-.2-.3a1 1 0 0 0-.29-.19h-.09a.86.86 0 0 0-.31-.15H20a1 1 0 0 0-.3 0l-18 6a1 1 0 0 0 0 1.9l8.53 2.84l2.84 8.53a1 1 0 0 0 1.9 0l6-18A1 1 0 0 0 21 4Zm-4.7 2.29l-5.57 5.57L5.16 10ZM14 18.84l-1.86-5.57l5.57-5.57Z"/></g></g>`
	paperPlaneOutlineInnerSVG             = `<g id="evaPaperPlaneOutline0"><g id="evaPaperPlaneOutline1"><path id="evaPaperPlaneOutline2" fill="currentColor" d="M21 4a1.31 1.31 0 0 0-.06-.27v-.09a1 1 0 0 0-.2-.3a1 1 0 0 0-.29-.19h-.09a.86.86 0 0 0-.31-.15H20a1 1 0 0 0-.3 0l-18 6a1 1 0 0 0 0 1.9l8.53 2.84l2.84 8.53a1 1 0 0 0 1.9 0l6-18A1 1 0 0 0 21 4Zm-4.7 2.29l-5.57 5.57L5.16 10ZM14 18.84l-1.86-5.57l5.57-5.57Z"/></g></g>`
	pauseCircleFillInnerSVG               = `<g id="evaPauseCircleFill0"><g id="evaPauseCircleFill1"><path id="evaPauseCircleFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm-2 13a1 1 0 0 1-2 0V9a1 1 0 0 1 2 0Zm6 0a1 1 0 0 1-2 0V9a1 1 0 0 1 2 0Z"/></g></g>`
	pauseCircleOutlineInnerSVG            = `<g id="evaPauseCircleOutline0"><g id="evaPauseCircleOutline1"><g id="evaPauseCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M15 8a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1ZM9 8a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1Z"/></g></g></g>`
	peopleFillInnerSVG                    = `<g id="evaPeopleFill0"><g id="evaPeopleFill1"><path id="evaPeopleFill2" fill="currentColor" d="M9 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm8 2a3 3 0 1 0-3-3a3 3 0 0 0 3 3Zm4 7a1 1 0 0 0 1-1a5 5 0 0 0-8.06-3.95A7 7 0 0 0 2 20a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1"/></g></g>`
	peopleOutlineInnerSVG                 = `<g id="evaPeopleOutline0"><g id="evaPeopleOutline1"><path id="evaPeopleOutline2" fill="currentColor" d="M9 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm8 8a3 3 0 1 0-3-3a3 3 0 0 0 3 3Zm0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm0 5a5 5 0 0 0-3.06 1.05A7 7 0 0 0 2 20a1 1 0 0 0 2 0a5 5 0 0 1 10 0a1 1 0 0 0 2 0a6.9 6.9 0 0 0-.86-3.35A3 3 0 0 1 20 19a1 1 0 0 0 2 0a5 5 0 0 0-5-5Z"/></g></g>`
	percentFillInnerSVG                   = `<g id="evaPercentFill0"><g id="evaPercentFill1"><path id="evaPercentFill2" fill="currentColor" d="M8 11a3.5 3.5 0 1 0-3.5-3.5A3.5 3.5 0 0 0 8 11Zm0-5a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 8 6Zm8 8a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 16 14Zm0 5a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 19Zm3.74-14.74a.89.89 0 0 0-1.26 0L4.26 18.48a.91.91 0 0 0-.26.63a.89.89 0 0 0 1.52.63L19.74 5.52a.89.89 0 0 0 0-1.26Z"/></g></g>`
	percentOutlineInnerSVG                = `<g id="evaPercentOutline0"><g id="evaPercentOutline1"><path id="evaPercentOutline2" fill="currentColor" d="M8 11a3.5 3.5 0 1 0-3.5-3.5A3.5 3.5 0 0 0 8 11Zm0-5a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 8 6Zm8 8a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 16 14Zm0 5a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 19Zm3.74-14.74a.89.89 0 0 0-1.26 0L4.26 18.48a.91.91 0 0 0-.26.63a.89.89 0 0 0 1.52.63L19.74 5.52a.89.89 0 0 0 0-1.26Z"/></g></g>`
	personAddFillInnerSVG                 = `<g id="evaPersonAddFill0"><g id="evaPersonAddFill1"><path id="evaPersonAddFill2" fill="currentColor" d="M21 6h-1V5a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0V8h1a1 1 0 0 0 0-2Zm-11 5a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm6 10a1 1 0 0 0 1-1a7 7 0 0 0-14 0a1 1 0 0 0 1 1"/></g></g>`
	personAddOutlineInnerSVG              = `<g id="evaPersonAddOutline0"><g id="evaPersonAddOutline1"><path id="evaPersonAddOutline2" fill="currentColor" d="M21 6h-1V5a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0V8h1a1 1 0 0 0 0-2Zm-11 5a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm0 8a7 7 0 0 0-7 7a1 1 0 0 0 2 0a5 5 0 0 1 10 0a1 1 0 0 0 2 0a7 7 0 0 0-7-7Z"/></g></g>`
	personDeleteFillInnerSVG              = `<g id="evaPersonDeleteFill0"><g id="evaPersonDeleteFill1"><path id="evaPersonDeleteFill2" fill="currentColor" d="m20.47 7.5l.73-.73a1 1 0 0 0-1.47-1.47L19 6l-.73-.73a1 1 0 0 0-1.47 1.5l.73.73l-.73.73a1 1 0 0 0 1.47 1.47L19 9l.73.73a1 1 0 0 0 1.47-1.5ZM10 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm6 10a1 1 0 0 0 1-1a7 7 0 0 0-14 0a1 1 0 0 0 1 1Z"/></g></g>`
	personDeleteOutlineInnerSVG           = `<g id="evaPersonDeleteOutline0"><g id="evaPersonDeleteOutline1"><path id="evaPersonDeleteOutline2" fill="currentColor" d="m20.47 7.5l.73-.73a1 1 0 0 0-1.47-1.47L19 6l-.73-.73a1 1 0 0 0-1.47 1.5l.73.73l-.73.73a1 1 0 0 0 1.47 1.47L19 9l.73.73a1 1 0 0 0 1.47-1.5ZM10 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm0 8a7 7 0 0 0-7 7a1 1 0 0 0 2 0a5 5 0 0 1 10 0a1 1 0 0 0 2 0a7 7 0 0 0-7-7Z"/></g></g>`
	personDoneFillInnerSVG                = `<g id="evaPersonDoneFill0"><g id="evaPersonDoneFill1"><path id="evaPersonDoneFill2" fill="currentColor" d="M21.66 4.25a1 1 0 0 0-1.41.09l-1.87 2.15l-.63-.71a1 1 0 0 0-1.5 1.33l1.39 1.56a1 1 0 0 0 .75.33a1 1 0 0 0 .74-.34l2.61-3a1 1 0 0 0-.08-1.41ZM10 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm6 10a1 1 0 0 0 1-1a7 7 0 0 0-14 0a1 1 0 0 0 1 1"/></g></g>`
	personDoneOutlineInnerSVG             = `<g id="evaPersonDoneOutline0"><g id="evaPersonDoneOutline1"><path id="evaPersonDoneOutline2" fill="currentColor" d="M21.66 4.25a1 1 0 0 0-1.41.09l-1.87 2.15l-.63-.71a1 1 0 0 0-1.5 1.33l1.39 1.56a1 1 0 0 0 .75.33a1 1 0 0 0 .74-.34l2.61-3a1 1 0 0 0-.08-1.41ZM10 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm0 8a7 7 0 0 0-7 7a1 1 0 0 0 2 0a5 5 0 0 1 10 0a1 1 0 0 0 2 0a7 7 0 0 0-7-7Z"/></g></g>`
	personFillInnerSVG                    = `<g id="evaPersonFill0"><g id="evaPersonFill1"><path id="evaPersonFill2" fill="currentColor" d="M12 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm6 10a1 1 0 0 0 1-1a7 7 0 0 0-14 0a1 1 0 0 0 1 1Z"/></g></g>`
	personOutlineInnerSVG                 = `<g id="evaPersonOutline0"><g id="evaPersonOutline1"><path id="evaPersonOutline2" fill="currentColor" d="M12 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm0 8a7 7 0 0 0-7 7a1 1 0 0 0 2 0a5 5 0 0 1 10 0a1 1 0 0 0 2 0a7 7 0 0 0-7-7Z"/></g></g>`
	personRemoveFillInnerSVG              = `<g id="evaPersonRemoveFill0"><g id="evaPersonRemoveFill1"><path id="evaPersonRemoveFill2" fill="currentColor" d="M21 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Zm-11 5a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm6 10a1 1 0 0 0 1-1a7 7 0 0 0-14 0a1 1 0 0 0 1 1"/></g></g>`
	personRemoveOutlineInnerSVG           = `<g id="evaPersonRemoveOutline0"><g id="evaPersonRemoveOutline1"><path id="evaPersonRemoveOutline2" fill="currentColor" d="M21 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Zm-11 5a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm0 8a7 7 0 0 0-7 7a1 1 0 0 0 2 0a5 5 0 0 1 10 0a1 1 0 0 0 2 0a7 7 0 0 0-7-7Z"/></g></g>`
	phoneCallFillInnerSVG                 = `<g id="evaPhoneCallFill0"><g id="evaPhoneCallFill1"><g id="evaPhoneCallFill2" fill="currentColor"><path d="M13 8a3 3 0 0 1 3 3a1 1 0 0 0 2 0a5 5 0 0 0-5-5a1 1 0 0 0 0 2Z"/><path d="M13 4a7 7 0 0 1 7 7a1 1 0 0 0 2 0a9 9 0 0 0-9-9a1 1 0 0 0 0 2Zm8.75 11.91a1 1 0 0 0-.72-.65l-6-1.37a1 1 0 0 0-.92.26c-.14.13-.15.14-.8 1.38a9.91 9.91 0 0 1-4.87-4.89C9.71 10 9.72 10 9.85 9.85a1 1 0 0 0 .26-.92L8.74 3a1 1 0 0 0-.65-.72a3.79 3.79 0 0 0-.72-.18A3.94 3.94 0 0 0 6.6 2A4.6 4.6 0 0 0 2 6.6A15.42 15.42 0 0 0 17.4 22a4.6 4.6 0 0 0 4.6-4.6a4.77 4.77 0 0 0-.06-.76a4.34 4.34 0 0 0-.19-.73Z"/></g></g></g>`
	phoneCallOutlineInnerSVG              = `<g id="evaPhoneCallOutline0"><g id="evaPhoneCallOutline1"><g id="evaPhoneCallOutline2" fill="currentColor"><path d="M13 8a3 3 0 0 1 3 3a1 1 0 0 0 2 0a5 5 0 0 0-5-5a1 1 0 0 0 0 2Z"/><path d="M13 4a7 7 0 0 1 7 7a1 1 0 0 0 2 0a9 9 0 0 0-9-9a1 1 0 0 0 0 2Zm8.75 11.91a1 1 0 0 0-.72-.65l-6-1.37a1 1 0 0 0-.92.26c-.14.13-.15.14-.8 1.38a9.91 9.91 0 0 1-4.87-4.89C9.71 10 9.72 10 9.85 9.85a1 1 0 0 0 .26-.92L8.74 3a1 1 0 0 0-.65-.72a3.79 3.79 0 0 0-.72-.18A3.94 3.94 0 0 0 6.6 2A4.6 4.6 0 0 0 2 6.6A15.42 15.42 0 0 0 17.4 22a4.6 4.6 0 0 0 4.6-4.6a4.77 4.77 0 0 0-.06-.76a4.34 4.34 0 0 0-.19-.73ZM17.4 20A13.41 13.41 0 0 1 4 6.6A2.61 2.61 0 0 1 6.6 4h.33L8 8.64l-.54.28c-.86.45-1.54.81-1.18 1.59a11.85 11.85 0 0 0 7.18 7.21c.84.34 1.17-.29 1.62-1.16l.29-.55L20 17.07v.33a2.61 2.61 0 0 1-2.6 2.6Z"/></g></g></g>`
	phoneFillInnerSVG                     = `<g id="evaPhoneFill0"><g id="evaPhoneFill1"><path id="evaPhoneFill2" fill="currentColor" d="M17.4 22A15.42 15.42 0 0 1 2 6.6A4.6 4.6 0 0 1 6.6 2a3.94 3.94 0 0 1 .77.07a3.79 3.79 0 0 1 .72.18a1 1 0 0 1 .65.75l1.37 6a1 1 0 0 1-.26.92c-.13.14-.14.15-1.37.79a9.91 9.91 0 0 0 4.87 4.89c.65-1.24.66-1.25.8-1.38a1 1 0 0 1 .92-.26l6 1.37a1 1 0 0 1 .72.65a4.34 4.34 0 0 1 .19.73a4.77 4.77 0 0 1 .06.76A4.6 4.6 0 0 1 17.4 22Z"/></g></g>`
	phoneMissedFillInnerSVG               = `<g id="evaPhoneMissedFill0"><g id="evaPhoneMissedFill1"><path id="evaPhoneMissedFill2" fill="currentColor" d="M21.94 16.64a4.34 4.34 0 0 0-.19-.73a1 1 0 0 0-.72-.65l-6-1.37a1 1 0 0 0-.92.26c-.14.13-.15.14-.8 1.38a10 10 0 0 1-4.88-4.89C9.71 10 9.72 10 9.85 9.85a1 1 0 0 0 .26-.92L8.74 3a1 1 0 0 0-.65-.72a3.79 3.79 0 0 0-.72-.18A3.94 3.94 0 0 0 6.6 2A4.6 4.6 0 0 0 2 6.6A15.42 15.42 0 0 0 17.4 22a4.6 4.6 0 0 0 4.6-4.6a4.77 4.77 0 0 0-.06-.76ZM15.8 8.7a1.05 1.05 0 0 0 1.47 0L18 8l.73.73a1 1 0 0 0 1.47-1.5l-.73-.73l.73-.73a1 1 0 0 0-1.47-1.47L18 5l-.73-.73a1 1 0 0 0-1.47 1.5l.73.73l-.73.73a1.05 1.05 0 0 0 0 1.47Z"/></g></g>`
	phoneMissedOutlineInnerSVG            = `<g id="evaPhoneMissedOutline0"><g id="evaPhoneMissedOutline1"><path id="evaPhoneMissedOutline2" fill="currentColor" d="M21.94 16.64a4.34 4.34 0 0 0-.19-.73a1 1 0 0 0-.72-.65l-6-1.37a1 1 0 0 0-.92.26c-.14.13-.15.14-.8 1.38a10 10 0 0 1-4.88-4.89C9.71 10 9.72 10 9.85 9.85a1 1 0 0 0 .26-.92L8.74 3a1 1 0 0 0-.65-.72a3.79 3.79 0 0 0-.72-.18A3.94 3.94 0 0 0 6.6 2A4.6 4.6 0 0 0 2 6.6A15.42 15.42 0 0 0 17.4 22a4.6 4.6 0 0 0 4.6-4.6a4.77 4.77 0 0 0-.06-.76ZM17.4 20A13.41 13.41 0 0 1 4 6.6A2.61 2.61 0 0 1 6.6 4h.33L8 8.64l-.55.29c-.87.45-1.5.78-1.17 1.58a11.85 11.85 0 0 0 7.18 7.21c.84.34 1.17-.29 1.62-1.16l.29-.55L20 17.07v.33a2.61 2.61 0 0 1-2.6 2.6ZM15.8 8.7a1.05 1.05 0 0 0 1.47 0L18 8l.73.73a1 1 0 0 0 1.47-1.5l-.73-.73l.73-.73a1 1 0 0 0-1.47-1.47L18 5l-.73-.73a1 1 0 0 0-1.47 1.5l.73.73l-.73.73a1.05 1.05 0 0 0 0 1.47Z"/></g></g>`
	phoneOffFillInnerSVG                  = `<g id="evaPhoneOffFill0"><g id="evaPhoneOffFill1"><path id="evaPhoneOffFill2" fill="currentColor" d="M9.27 12.06a10.37 10.37 0 0 1-.8-1.42C9.71 10 9.72 10 9.85 9.85a1 1 0 0 0 .26-.92L8.74 3a1 1 0 0 0-.65-.72a3.79 3.79 0 0 0-.72-.18A3.94 3.94 0 0 0 6.6 2A4.6 4.6 0 0 0 2 6.6a15.33 15.33 0 0 0 3.27 9.46Zm12.67 4.58a4.34 4.34 0 0 0-.19-.73a1 1 0 0 0-.72-.65l-6-1.37a1 1 0 0 0-.92.26c-.14.13-.15.14-.8 1.38a10.88 10.88 0 0 1-1.41-.8l-4 4A15.33 15.33 0 0 0 17.4 22a4.6 4.6 0 0 0 4.6-4.6a4.77 4.77 0 0 0-.06-.76Zm-2.2-12.38a.89.89 0 0 0-1.26 0L4.26 18.48a.91.91 0 0 0-.26.63a.89.89 0 0 0 1.52.63L19.74 5.52a.89.89 0 0 0 0-1.26Z"/></g></g>`
	phoneOffOutlineInnerSVG               = `<g id="evaPhoneOffOutline0"><g id="evaPhoneOffOutline1"><path id="evaPhoneOffOutline2" fill="currentColor" d="M19.74 4.26a.89.89 0 0 0-1.26 0L4.26 18.48a.91.91 0 0 0-.26.63a.89.89 0 0 0 1.52.63L19.74 5.52a.89.89 0 0 0 0-1.26ZM6.7 14.63A13.29 13.29 0 0 1 4 6.6A2.61 2.61 0 0 1 6.6 4h.33L8 8.64l-.55.29c-.87.45-1.5.78-1.17 1.58a11.57 11.57 0 0 0 1.57 3l1.43-1.42a10.37 10.37 0 0 1-.8-1.42C9.71 10 9.72 10 9.85 9.85a1 1 0 0 0 .26-.92L8.74 3a1 1 0 0 0-.65-.72a3.79 3.79 0 0 0-.72-.18A3.94 3.94 0 0 0 6.6 2A4.6 4.6 0 0 0 2 6.6a15.33 15.33 0 0 0 3.27 9.46Zm15.24 2.01a4.34 4.34 0 0 0-.19-.73a1 1 0 0 0-.72-.65l-6-1.37a1 1 0 0 0-.92.26c-.14.13-.15.14-.8 1.38a10.88 10.88 0 0 1-1.41-.8l-1.43 1.43a11.52 11.52 0 0 0 2.94 1.56c.84.34 1.17-.29 1.62-1.16l.29-.55L20 17.07v.33a2.61 2.61 0 0 1-2.6 2.6a13.29 13.29 0 0 1-8-2.7l-1.46 1.43A15.33 15.33 0 0 0 17.4 22a4.6 4.6 0 0 0 4.6-4.6a4.77 4.77 0 0 0-.06-.76Z"/></g></g>`
	phoneOutlineInnerSVG                  = `<g id="evaPhoneOutline0"><g id="evaPhoneOutline1"><path id="evaPhoneOutline2" fill="currentColor" d="M17.4 22A15.42 15.42 0 0 1 2 6.6A4.6 4.6 0 0 1 6.6 2a3.94 3.94 0 0 1 .77.07a3.79 3.79 0 0 1 .72.18a1 1 0 0 1 .65.75l1.37 6a1 1 0 0 1-.26.92c-.13.14-.14.15-1.37.79a9.91 9.91 0 0 0 4.87 4.89c.65-1.24.66-1.25.8-1.38a1 1 0 0 1 .92-.26l6 1.37a1 1 0 0 1 .72.65a4.34 4.34 0 0 1 .19.73a4.77 4.77 0 0 1 .06.76A4.6 4.6 0 0 1 17.4 22ZM6.6 4A2.61 2.61 0 0 0 4 6.6A13.41 13.41 0 0 0 17.4 20a2.61 2.61 0 0 0 2.6-2.6v-.33L15.36 16l-.29.55c-.45.87-.78 1.5-1.62 1.16a11.85 11.85 0 0 1-7.18-7.21c-.36-.78.32-1.14 1.18-1.59L8 8.64L6.93 4Z"/></g></g>`
	pieChartFillInnerSVG                  = `<g id="evaPieChartFill0"><g id="evaPieChartFill1"><g id="evaPieChartFill2" fill="currentColor"><path d="M14.5 10.33h6.67A.83.83 0 0 0 22 9.5A7.5 7.5 0 0 0 14.5 2a.83.83 0 0 0-.83.83V9.5a.83.83 0 0 0 .83.83Z"/><path d="M21.08 12h-8.15a.91.91 0 0 1-.91-.91V2.92A.92.92 0 0 0 11 2a10 10 0 1 0 11 11a.92.92 0 0 0-.92-1Z"/></g></g></g>`
	pieChartOutlineInnerSVG               = `<g id="evaPieChartOutline0"><g id="evaPieChartOutline1"><g id="evaPieChartOutline2" fill="currentColor"><path d="M13 2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1a9 9 0 0 0-9-9Zm1 8V4.07A7 7 0 0 1 19.93 10Z"/><path d="M20.82 14.06a1 1 0 0 0-1.28.61A8 8 0 1 1 9.33 4.46a1 1 0 0 0-.66-1.89a10 10 0 1 0 12.76 12.76a1 1 0 0 0-.61-1.27Z"/></g></g></g>`
	pieChartTwoFillInnerSVG               = `<g id="evaPieChart2Fill0"><g id="evaPieChart2Fill1"><g id="evaPieChart2Fill2" fill="currentColor"><path d="M14.5 10.33h6.67A.83.83 0 0 0 22 9.5A7.5 7.5 0 0 0 14.5 2a.83.83 0 0 0-.83.83V9.5a.83.83 0 0 0 .83.83Zm.83-6.6a5.83 5.83 0 0 1 4.94 4.94h-4.94Z"/><path d="M21.08 12h-8.15a.91.91 0 0 1-.91-.91V2.92A.92.92 0 0 0 11 2a10 10 0 1 0 11 11a.92.92 0 0 0-.92-1Z"/></g></g></g>`
	pinFillInnerSVG                       = `<g id="evaPinFill0"><g id="evaPinFill1"><g id="evaPinFill2" fill="currentColor"><circle cx="12" cy="9.5" r="1.5"/><path d="M12 2a8 8 0 0 0-8 7.92c0 5.48 7.05 11.58 7.35 11.84a1 1 0 0 0 1.3 0C13 21.5 20 15.4 20 9.92A8 8 0 0 0 12 2Zm0 11a3.5 3.5 0 1 1 3.5-3.5A3.5 3.5 0 0 1 12 13Z"/></g></g></g>`
	pinOutlineInnerSVG                    = `<g id="evaPinOutline0"><g id="evaPinOutline1"><g id="evaPinOutline2" fill="currentColor"><path d="M12 2a8 8 0 0 0-8 7.92c0 5.48 7.05 11.58 7.35 11.84a1 1 0 0 0 1.3 0C13 21.5 20 15.4 20 9.92A8 8 0 0 0 12 2Zm0 17.65c-1.67-1.59-6-6-6-9.73a6 6 0 0 1 12 0c0 3.7-4.33 8.14-6 9.73Z"/><path d="M12 6a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 12 6Zm0 5a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 12 11Z"/></g></g></g>`
	playCircleFillInnerSVG                = `<g id="evaPlayCircleFill0"><g id="evaPlayCircleFill1"><g id="evaPlayCircleFill2" fill="currentColor"><path d="m11.5 14.6l2.81-2.6l-2.81-2.6v5.2z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm4 11.18l-3.64 3.37a1.74 1.74 0 0 1-1.16.45a1.68 1.68 0 0 1-.69-.15a1.6 1.6 0 0 1-1-1.48V8.63a1.6 1.6 0 0 1 1-1.48a1.7 1.7 0 0 1 1.85.3L16 10.82a1.6 1.6 0 0 1 0 2.36Z"/></g></g></g>`
	playCircleOutlineInnerSVG             = `<g id="evaPlayCircleOutline0"><g id="evaPlayCircleOutline1"><g id="evaPlayCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M12.34 7.45a1.7 1.7 0 0 0-1.85-.3a1.6 1.6 0 0 0-1 1.48v6.74a1.6 1.6 0 0 0 1 1.48a1.68 1.68 0 0 0 .69.15a1.74 1.74 0 0 0 1.16-.45L16 13.18a1.6 1.6 0 0 0 0-2.36Zm-.84 7.15V9.4l2.81 2.6Z"/></g></g></g>`
	plusCircleFillInnerSVG                = `<g id="evaPlusCircleFill0"><g id="evaPlusCircleFill1"><path id="evaPlusCircleFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm3 11h-2v2a1 1 0 0 1-2 0v-2H9a1 1 0 0 1 0-2h2V9a1 1 0 0 1 2 0v2h2a1 1 0 0 1 0 2Z"/></g></g>`
	plusCircleOutlineInnerSVG             = `<g id="evaPlusCircleOutline0"><g id="evaPlusCircleOutline1"><g id="evaPlusCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M15 11h-2V9a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2Z"/></g></g></g>`
	plusFillInnerSVG                      = `<g id="evaPlusFill0"><g id="evaPlusFill1"><path id="evaPlusFill2" fill="currentColor" d="M19 11h-6V5a1 1 0 0 0-2 0v6H5a1 1 0 0 0 0 2h6v6a1 1 0 0 0 2 0v-6h6a1 1 0 0 0 0-2Z"/></g></g>`
	plusOutlineInnerSVG                   = `<g id="evaPlusOutline0"><g id="evaPlusOutline1"><path id="evaPlusOutline2" fill="currentColor" d="M19 11h-6V5a1 1 0 0 0-2 0v6H5a1 1 0 0 0 0 2h6v6a1 1 0 0 0 2 0v-6h6a1 1 0 0 0 0-2Z"/></g></g>`
	plusSquareFillInnerSVG                = `<g id="evaPlusSquareFill0"><g id="evaPlusSquareFill1"><path id="evaPlusSquareFill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm-3 10h-2v2a1 1 0 0 1-2 0v-2H9a1 1 0 0 1 0-2h2V9a1 1 0 0 1 2 0v2h2a1 1 0 0 1 0 2Z"/></g></g>`
	plusSquareOutlineInnerSVG             = `<g id="evaPlusSquareOutline0"><g id="evaPlusSquareOutline1"><g id="evaPlusSquareOutline2" fill="currentColor"><path d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3Zm1 15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/><path d="M15 11h-2V9a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2Z"/></g></g></g>`
	powerFillInnerSVG                     = `<g id="evaPowerFill0"><g id="evaPowerFill1"><g id="evaPowerFill2" fill="currentColor"><path d="M12 13a1 1 0 0 0 1-1V2a1 1 0 0 0-2 0v10a1 1 0 0 0 1 1Z"/><path d="M16.59 3.11a1 1 0 0 0-.92 1.78a8 8 0 1 1-7.34 0a1 1 0 1 0-.92-1.78a10 10 0 1 0 9.18 0Z"/></g></g></g>`
	powerOutlineInnerSVG                  = `<g id="evaPowerOutline0"><g id="evaPowerOutline1"><g id="evaPowerOutline2" fill="currentColor"><path d="M12 13a1 1 0 0 0 1-1V2a1 1 0 0 0-2 0v10a1 1 0 0 0 1 1Z"/><path d="M16.59 3.11a1 1 0 0 0-.92 1.78a8 8 0 1 1-7.34 0a1 1 0 1 0-.92-1.78a10 10 0 1 0 9.18 0Z"/></g></g></g>`
	pricetagsFillInnerSVG                 = `<g id="evaPricetagsFill0"><g id="evaPricetagsFill1"><path id="evaPricetagsFill2" fill="currentColor" d="m21.47 11.58l-6.42-6.41a1 1 0 0 0-.61-.29L5.09 4a1 1 0 0 0-.8.29a1 1 0 0 0-.29.8l.88 9.35a1 1 0 0 0 .29.61l6.41 6.42a1.84 1.84 0 0 0 1.29.53a1.82 1.82 0 0 0 1.28-.53l7.32-7.32a1.82 1.82 0 0 0 0-2.57Zm-9.91 0a1.5 1.5 0 1 1 0-2.12a1.49 1.49 0 0 1 0 2.1Z"/></g></g>`
	pricetagsOutlineInnerSVG              = `<g id="evaPricetagsOutline0"><g id="evaPricetagsOutline1"><g id="evaPricetagsOutline2" fill="currentColor"><path d="M12.87 22a1.84 1.84 0 0 1-1.29-.53l-6.41-6.42a1 1 0 0 1-.29-.61L4 5.09a1 1 0 0 1 .29-.8a1 1 0 0 1 .8-.29l9.35.88a1 1 0 0 1 .61.29l6.42 6.41a1.82 1.82 0 0 1 0 2.57l-7.32 7.32a1.82 1.82 0 0 1-1.28.53Zm-6-8.11l6 6l7.05-7.05l-6-6l-7.81-.73Z"/><circle cx="10.5" cy="10.5" r="1.5"/></g></g></g>`
	printerFillInnerSVG                   = `<g id="evaPrinterFill0"><g id="evaPrinterFill1"><path id="evaPrinterFill2" fill="currentColor" d="M19.36 7H18V5a1.92 1.92 0 0 0-1.83-2H7.83A1.92 1.92 0 0 0 6 5v2H4.64A2.66 2.66 0 0 0 2 9.67v6.66A2.66 2.66 0 0 0 4.64 19h.86a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2h.86A2.66 2.66 0 0 0 22 16.33V9.67A2.66 2.66 0 0 0 19.36 7ZM8 5h8v2H8Zm-.5 14v-4h9v4Z"/></g></g>`
	printerOutlineInnerSVG                = `<g id="evaPrinterOutline0"><g id="evaPrinterOutline1"><path id="evaPrinterOutline2" fill="currentColor" d="M19.36 7H18V5a1.92 1.92 0 0 0-1.83-2H7.83A1.92 1.92 0 0 0 6 5v2H4.64A2.66 2.66 0 0 0 2 9.67v6.66A2.66 2.66 0 0 0 4.64 19h.86a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2h.86A2.66 2.66 0 0 0 22 16.33V9.67A2.66 2.66 0 0 0 19.36 7ZM8 5h8v2H8Zm-.5 14v-4h9v4ZM20 16.33a.66.66 0 0 1-.64.67h-.86v-2a2 2 0 0 0-2-2h-9a2 2 0 0 0-2 2v2h-.86a.66.66 0 0 1-.64-.67V9.67A.66.66 0 0 1 4.64 9h14.72a.66.66 0 0 1 .64.67Z"/></g></g>`
	questionMarkCircleFillInnerSVG        = `<g id="evaQuestionMarkCircleFill0"><g id="evaQuestionMarkCircleFill1"><path id="evaQuestionMarkCircleFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 16a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm1-5.16V14a1 1 0 0 1-2 0v-2a1 1 0 0 1 1-1a1.5 1.5 0 1 0-1.5-1.5a1 1 0 0 1-2 0a3.5 3.5 0 1 1 4.5 3.34Z"/></g></g>`
	questionMarkCircleOutlineInnerSVG     = `<g id="evaQuestionMarkCircleOutline0"><g id="evaQuestionMarkCircleOutline1"><g id="evaQuestionMarkCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M12 6a3.5 3.5 0 0 0-3.5 3.5a1 1 0 0 0 2 0A1.5 1.5 0 1 1 12 11a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-1.16A3.49 3.49 0 0 0 12 6Z"/><circle cx="12" cy="17" r="1"/></g></g></g>`
	questionMarkFillInnerSVG              = `<g id="evaQuestionMarkFill0"><g id="evaQuestionMarkFill1"><g id="evaQuestionMarkFill2" fill="currentColor"><path d="M17 9A5 5 0 0 0 7 9a1 1 0 0 0 2 0a3 3 0 1 1 3 3a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-1.1A5 5 0 0 0 17 9Z"/><circle cx="12" cy="19" r="1"/></g></g></g>`
	questionMarkOutlineInnerSVG           = `<g id="evaQuestionMarkOutline0"><g id="evaQuestionMarkOutline1"><g id="evaQuestionMarkOutline2" fill="currentColor"><path d="M17 9A5 5 0 0 0 7 9a1 1 0 0 0 2 0a3 3 0 1 1 3 3a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-1.1A5 5 0 0 0 17 9Z"/><circle cx="12" cy="19" r="1"/></g></g></g>`
	radioButtonOffFillInnerSVG            = `<g id="evaRadioButtonOffFill0"><g id="evaRadioButtonOffFill1"><path id="evaRadioButtonOffFill2" fill="currentColor" d="M12 22a10 10 0 1 1 10-10a10 10 0 0 1-10 10Zm0-18a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/></g></g>`
	radioButtonOffOutlineInnerSVG         = `<g id="evaRadioButtonOffOutline0"><g id="evaRadioButtonOffOutline1"><path id="evaRadioButtonOffOutline2" fill="currentColor" d="M12 22a10 10 0 1 1 10-10a10 10 0 0 1-10 10Zm0-18a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/></g></g>`
	radioButtonOnFillInnerSVG             = `<g id="evaRadioButtonOnFill0"><g id="evaRadioButtonOnFill1"><g id="evaRadioButtonOnFill2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M12 7a5 5 0 1 0 5 5a5 5 0 0 0-5-5Z"/></g></g></g>`
	radioButtonOnOutlineInnerSVG          = `<g id="evaRadioButtonOnOutline0"><g id="evaRadioButtonOnOutline1"><g id="evaRadioButtonOnOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M12 7a5 5 0 1 0 5 5a5 5 0 0 0-5-5Zm0 8a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/></g></g></g>`
	radioFillInnerSVG                     = `<g id="evaRadioFill0"><g id="evaRadioFill1"><g id="evaRadioFill2" fill="currentColor"><path d="M12 8a3 3 0 0 0-1 5.83a1 1 0 0 0 0 .17v6a1 1 0 0 0 2 0v-6a1 1 0 0 0 0-.17A3 3 0 0 0 12 8Zm-8.5 3a6.87 6.87 0 0 1 2.64-5.23a1 1 0 1 0-1.28-1.54A8.84 8.84 0 0 0 1.5 11a8.84 8.84 0 0 0 3.36 6.77a1 1 0 1 0 1.28-1.54A6.87 6.87 0 0 1 3.5 11Z"/><path d="M16.64 6.24a1 1 0 0 0-1.28 1.52A4.28 4.28 0 0 1 17 11a4.28 4.28 0 0 1-1.64 3.24A1 1 0 0 0 16 16a1 1 0 0 0 .64-.24A6.2 6.2 0 0 0 19 11a6.2 6.2 0 0 0-2.36-4.76Zm-7.88.12a1 1 0 0 0-1.4-.12A6.2 6.2 0 0 0 5 11a6.2 6.2 0 0 0 2.36 4.76a1 1 0 0 0 1.4-.12a1 1 0 0 0-.12-1.4A4.28 4.28 0 0 1 7 11a4.28 4.28 0 0 1 1.64-3.24a1 1 0 0 0 .12-1.4Z"/><path d="M19.14 4.23a1 1 0 1 0-1.28 1.54A6.87 6.87 0 0 1 20.5 11a6.87 6.87 0 0 1-2.64 5.23a1 1 0 0 0 1.28 1.54A8.84 8.84 0 0 0 22.5 11a8.84 8.84 0 0 0-3.36-6.77Z"/></g></g></g>`
	radioOutlineInnerSVG                  = `<g id="evaRadioOutline0"><g id="evaRadioOutline1"><g id="evaRadioOutline2" fill="currentColor"><path d="M12 8a3 3 0 0 0-1 5.83a1 1 0 0 0 0 .17v6a1 1 0 0 0 2 0v-6a1 1 0 0 0 0-.17A3 3 0 0 0 12 8Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm-8.5-1a6.87 6.87 0 0 1 2.64-5.23a1 1 0 1 0-1.28-1.54A8.84 8.84 0 0 0 1.5 11a8.84 8.84 0 0 0 3.36 6.77a1 1 0 1 0 1.28-1.54A6.87 6.87 0 0 1 3.5 11Z"/><path d="M16.64 6.24a1 1 0 0 0-1.28 1.52A4.28 4.28 0 0 1 17 11a4.28 4.28 0 0 1-1.64 3.24A1 1 0 0 0 16 16a1 1 0 0 0 .64-.24A6.2 6.2 0 0 0 19 11a6.2 6.2 0 0 0-2.36-4.76Zm-7.88.12a1 1 0 0 0-1.4-.12A6.2 6.2 0 0 0 5 11a6.2 6.2 0 0 0 2.36 4.76a1 1 0 0 0 1.4-.12a1 1 0 0 0-.12-1.4A4.28 4.28 0 0 1 7 11a4.28 4.28 0 0 1 1.64-3.24a1 1 0 0 0 .12-1.4Z"/><path d="M19.14 4.23a1 1 0 1 0-1.28 1.54A6.87 6.87 0 0 1 20.5 11a6.87 6.87 0 0 1-2.64 5.23a1 1 0 0 0 1.28 1.54A8.84 8.84 0 0 0 22.5 11a8.84 8.84 0 0 0-3.36-6.77Z"/></g></g></g>`
	recordingFillInnerSVG                 = `<g id="evaRecordingFill0"><g id="evaRecordingFill1"><path id="evaRecordingFill2" fill="currentColor" d="M18 8a4 4 0 0 0-4 4a3.91 3.91 0 0 0 .56 2H9.44a3.91 3.91 0 0 0 .56-2a4 4 0 1 0-4 4h12a4 4 0 0 0 0-8Z"/></g></g>`
	recordingOutlineInnerSVG              = `<g id="evaRecordingOutline0"><g id="evaRecordingOutline1"><path id="evaRecordingOutline2" fill="currentColor" d="M18 8a4 4 0 0 0-4 4a3.91 3.91 0 0 0 .56 2H9.44a3.91 3.91 0 0 0 .56-2a4 4 0 1 0-4 4h12a4 4 0 0 0 0-8ZM4 12a2 2 0 1 1 2 2a2 2 0 0 1-2-2Zm14 2a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/></g></g>`
	refreshFillInnerSVG                   = `<g id="evaRefreshFill0"><g id="evaRefreshFill1"><path id="evaRefreshFill2" fill="currentColor" d="M20.3 13.43a1 1 0 0 0-1.25.65A7.14 7.14 0 0 1 12.18 19A7.1 7.1 0 0 1 5 12a7.1 7.1 0 0 1 7.18-7a7.26 7.26 0 0 1 4.65 1.67l-2.17-.36a1 1 0 0 0-1.15.83a1 1 0 0 0 .83 1.15l4.24.7h.17a1 1 0 0 0 .34-.06a.33.33 0 0 0 .1-.06a.78.78 0 0 0 .2-.11l.09-.11c0-.05.09-.09.13-.15s0-.1.05-.14a1.34 1.34 0 0 0 .07-.18l.75-4a1 1 0 0 0-2-.38l-.27 1.45A9.21 9.21 0 0 0 12.18 3A9.1 9.1 0 0 0 3 12a9.1 9.1 0 0 0 9.18 9A9.12 9.12 0 0 0 21 14.68a1 1 0 0 0-.7-1.25Z"/></g></g>`
	refreshOutlineInnerSVG                = `<g id="evaRefreshOutline0"><g id="evaRefreshOutline1"><path id="evaRefreshOutline2" fill="currentColor" d="M20.3 13.43a1 1 0 0 0-1.25.65A7.14 7.14 0 0 1 12.18 19A7.1 7.1 0 0 1 5 12a7.1 7.1 0 0 1 7.18-7a7.26 7.26 0 0 1 4.65 1.67l-2.17-.36a1 1 0 0 0-1.15.83a1 1 0 0 0 .83 1.15l4.24.7h.17a1 1 0 0 0 .34-.06a.33.33 0 0 0 .1-.06a.78.78 0 0 0 .2-.11l.09-.11c0-.05.09-.09.13-.15s0-.1.05-.14a1.34 1.34 0 0 0 .07-.18l.75-4a1 1 0 0 0-2-.38l-.27 1.45A9.21 9.21 0 0 0 12.18 3A9.1 9.1 0 0 0 3 12a9.1 9.1 0 0 0 9.18 9A9.12 9.12 0 0 0 21 14.68a1 1 0 0 0-.7-1.25Z"/></g></g>`
	repeatFillInnerSVG                    = `<g id="evaRepeatFill0"><g id="evaRepeatFill1"><path id="evaRepeatFill2" fill="currentColor" d="M17.91 5h-12l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0 0 1.42l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.91 7h12a1.56 1.56 0 0 1 1.59 1.53V11a1 1 0 0 0 2 0V8.53A3.56 3.56 0 0 0 17.91 5Zm.3 9.29a1 1 0 0 0-1.42 1.42l1.3 1.29h-12a1.56 1.56 0 0 1-1.59-1.53V13a1 1 0 0 0-2 0v2.47A3.56 3.56 0 0 0 6.09 19h12l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 0-1.42Z"/></g></g>`
	repeatOutlineInnerSVG                 = `<g id="evaRepeatOutline0"><g id="evaRepeatOutline1"><path id="evaRepeatOutline2" fill="currentColor" d="M17.91 5h-12l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0 0 1.42l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.91 7h12a1.56 1.56 0 0 1 1.59 1.53V11a1 1 0 0 0 2 0V8.53A3.56 3.56 0 0 0 17.91 5Zm.3 9.29a1 1 0 0 0-1.42 1.42l1.3 1.29h-12a1.56 1.56 0 0 1-1.59-1.53V13a1 1 0 0 0-2 0v2.47A3.56 3.56 0 0 0 6.09 19h12l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 0-1.42Z"/></g></g>`
	rewindLeftFillInnerSVG                = `<g id="evaRewindLeftFill0"><g id="evaRewindLeftFill1"><path id="evaRewindLeftFill2" fill="currentColor" d="M18.45 6.2a2.1 2.1 0 0 0-2.21.26l-4.74 3.92V7.79a1.76 1.76 0 0 0-1.05-1.59a2.1 2.1 0 0 0-2.21.26l-5.1 4.21a1.7 1.7 0 0 0 0 2.66l5.1 4.21a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59v-2.59l4.74 3.92a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59V7.79a1.76 1.76 0 0 0-1.05-1.59Z"/></g></g>`
	rewindLeftOutlineInnerSVG             = `<g id="evaRewindLeftOutline0"><g id="evaRewindLeftOutline1"><path id="evaRewindLeftOutline2" fill="currentColor" d="M18.45 6.2a2.1 2.1 0 0 0-2.21.26l-4.74 3.92V7.79a1.76 1.76 0 0 0-1.05-1.59a2.1 2.1 0 0 0-2.21.26l-5.1 4.21a1.7 1.7 0 0 0 0 2.66l5.1 4.21a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59v-2.59l4.74 3.92a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59V7.79a1.76 1.76 0 0 0-1.05-1.59ZM9.5 16l-4.82-4L9.5 8.09Zm8 0l-4.82-4l4.82-3.91Z"/></g></g>`
	rewindRightFillInnerSVG               = `<g id="evaRewindRightFill0"><g id="evaRewindRightFill1"><path id="evaRewindRightFill2" fill="currentColor" d="m20.86 10.67l-5.1-4.21a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1.05 1.59v2.59L7.76 6.46a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1 1.59v8.42a1.76 1.76 0 0 0 1 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l4.74-3.92v2.59a1.76 1.76 0 0 0 1.05 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l5.1-4.21a1.7 1.7 0 0 0 0-2.66Z"/></g></g>`
	rewindRightOutlineInnerSVG            = `<g id="evaRewindRightOutline0"><g id="evaRewindRightOutline1"><path id="evaRewindRightOutline2" fill="currentColor" d="m20.86 10.67l-5.1-4.21a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1.05 1.59v2.59L7.76 6.46a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1 1.59v8.42a1.76 1.76 0 0 0 1 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l4.74-3.92v2.59a1.76 1.76 0 0 0 1.05 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l5.1-4.21a1.7 1.7 0 0 0 0-2.66ZM6.5 15.91V8l4.82 4Zm8 0V8l4.82 4Z"/></g></g>`
	saveFillInnerSVG                      = `<g id="evaSaveFill0"><g id="evaSaveFill1"><g id="evaSaveFill2" fill="currentColor"><path d="M10 17h4v4h-4z"/><path d="m20.12 8.71l-4.83-4.83A3 3 0 0 0 13.17 3H10v6h5a1 1 0 0 1 0 2H9a1 1 0 0 1-1-1V3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h2v-4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4h2a3 3 0 0 0 3-3v-7.17a3 3 0 0 0-.88-2.12Z"/></g></g></g>`
	saveOutlineInnerSVG                   = `<g id="evaSaveOutline0"><g id="evaSaveOutline1"><path id="evaSaveOutline2" fill="currentColor" d="m20.12 8.71l-4.83-4.83A3 3 0 0 0 13.17 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-7.17a3 3 0 0 0-.88-2.12ZM10 19v-2h4v2Zm9-1a1 1 0 0 1-1 1h-2v-3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h2v5a1 1 0 0 0 1 1h4a1 1 0 0 0 0-2h-3V5h3.17a1.05 1.05 0 0 1 .71.29l4.83 4.83a1 1 0 0 1 .29.71Z"/></g></g>`
	scissorsFillInnerSVG                  = `<g id="evaScissorsFill0"><g id="evaScissorsFill1"><g id="evaScissorsFill2" fill="currentColor"><path d="M20.21 5.71a1 1 0 1 0-1.42-1.42l-6.28 6.31l-3.3-3.31A3 3 0 0 0 9.5 6a3 3 0 1 0-3 3a3 3 0 0 0 1.29-.3L11.1 12l-3.29 3.3A3 3 0 0 0 6.5 15a3 3 0 1 0 3 3a3 3 0 0 0-.29-1.26ZM6.5 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm0 12a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><path d="M15.21 13.29a1 1 0 0 0-1.42 1.42l5 5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g></g>`
	scissorsOutlineInnerSVG               = `<g id="evaScissorsOutline0"><g id="evaScissorsOutline1"><g id="evaScissorsOutline2" fill="currentColor"><path d="M20.21 5.71a1 1 0 1 0-1.42-1.42l-6.28 6.31l-3.3-3.31A3 3 0 0 0 9.5 6a3 3 0 1 0-3 3a3 3 0 0 0 1.29-.3L11.1 12l-3.29 3.3A3 3 0 0 0 6.5 15a3 3 0 1 0 3 3a3 3 0 0 0-.29-1.26ZM6.5 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm0 12a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><path d="M15.21 13.29a1 1 0 0 0-1.42 1.42l5 5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g></g>`
	searchFillInnerSVG                    = `<g id="evaSearchFill0"><g id="evaSearchFill1"><path id="evaSearchFill2" fill="currentColor" d="m20.71 19.29l-3.4-3.39A7.92 7.92 0 0 0 19 11a8 8 0 1 0-8 8a7.92 7.92 0 0 0 4.9-1.69l3.39 3.4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM5 11a6 6 0 1 1 6 6a6 6 0 0 1-6-6Z"/></g></g>`
	searchOutlineInnerSVG                 = `<g id="evaSearchOutline0"><g id="evaSearchOutline1"><path id="evaSearchOutline2" fill="currentColor" d="m20.71 19.29l-3.4-3.39A7.92 7.92 0 0 0 19 11a8 8 0 1 0-8 8a7.92 7.92 0 0 0 4.9-1.69l3.39 3.4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM5 11a6 6 0 1 1 6 6a6 6 0 0 1-6-6Z"/></g></g>`
	settingsFillInnerSVG                  = `<g id="evaSettingsFill0"><g id="evaSettingsFill1"><g id="evaSettingsFill2" fill="currentColor"><circle cx="12" cy="12" r="1.5"/><path d="M21.89 10.32L21.1 7.8a2.26 2.26 0 0 0-2.88-1.51l-.34.11a1.74 1.74 0 0 1-1.59-.26l-.11-.08a1.76 1.76 0 0 1-.69-1.43v-.28a2.37 2.37 0 0 0-.68-1.68a2.26 2.26 0 0 0-1.6-.67h-2.55a2.32 2.32 0 0 0-2.29 2.33v.24a1.94 1.94 0 0 1-.73 1.51l-.13.1a1.93 1.93 0 0 1-1.78.29a2.14 2.14 0 0 0-1.68.12a2.18 2.18 0 0 0-1.12 1.33l-.82 2.6a2.34 2.34 0 0 0 1.48 2.94h.16a1.83 1.83 0 0 1 1.12 1.22l.06.16a2.06 2.06 0 0 1-.23 1.86a2.37 2.37 0 0 0 .49 3.3l2.07 1.57a2.25 2.25 0 0 0 1.35.43A2 2 0 0 0 9 22a2.25 2.25 0 0 0 1.47-1l.23-.33a1.8 1.8 0 0 1 1.43-.77a1.75 1.75 0 0 1 1.5.78l.12.17a2.24 2.24 0 0 0 3.22.53L19 19.86a2.38 2.38 0 0 0 .5-3.23l-.26-.38A2 2 0 0 1 19 14.6a1.89 1.89 0 0 1 1.21-1.28l.2-.07a2.36 2.36 0 0 0 1.48-2.93ZM12 15.5a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5Z"/></g></g></g>`
	settingsOutlineInnerSVG               = `<g id="evaSettingsOutline0"><g id="evaSettingsOutline1"><g id="evaSettingsOutline2" fill="currentColor"><path id="evaSettingsOutline3" d="M8.61 22a2.25 2.25 0 0 1-1.35-.46L5.19 20a2.37 2.37 0 0 1-.49-3.22a2.06 2.06 0 0 0 .23-1.86l-.06-.16a1.83 1.83 0 0 0-1.12-1.22h-.16a2.34 2.34 0 0 1-1.48-2.94L2.93 8a2.18 2.18 0 0 1 1.12-1.41a2.14 2.14 0 0 1 1.68-.12a1.93 1.93 0 0 0 1.78-.29l.13-.1a1.94 1.94 0 0 0 .73-1.51v-.24A2.32 2.32 0 0 1 10.66 2h2.55a2.26 2.26 0 0 1 1.6.67a2.37 2.37 0 0 1 .68 1.68v.28a1.76 1.76 0 0 0 .69 1.43l.11.08a1.74 1.74 0 0 0 1.59.26l.34-.11A2.26 2.26 0 0 1 21.1 7.8l.79 2.52a2.36 2.36 0 0 1-1.46 2.93l-.2.07A1.89 1.89 0 0 0 19 14.6a2 2 0 0 0 .25 1.65l.26.38a2.38 2.38 0 0 1-.5 3.23L17 21.41a2.24 2.24 0 0 1-3.22-.53l-.12-.17a1.75 1.75 0 0 0-1.5-.78a1.8 1.8 0 0 0-1.43.77l-.23.33A2.25 2.25 0 0 1 9 22a2 2 0 0 1-.39 0ZM4.4 11.62a3.83 3.83 0 0 1 2.38 2.5v.12a4 4 0 0 1-.46 3.62a.38.38 0 0 0 0 .51L8.47 20a.25.25 0 0 0 .37-.07l.23-.33a3.77 3.77 0 0 1 6.2 0l.12.18a.3.3 0 0 0 .18.12a.25.25 0 0 0 .19-.05l2.06-1.56a.36.36 0 0 0 .07-.49l-.26-.38A4 4 0 0 1 17.1 14a3.92 3.92 0 0 1 2.49-2.61l.2-.07a.34.34 0 0 0 .19-.44l-.78-2.49a.35.35 0 0 0-.2-.19a.21.21 0 0 0-.19 0l-.34.11a3.74 3.74 0 0 1-3.43-.57L15 7.65a3.76 3.76 0 0 1-1.49-3v-.31a.37.37 0 0 0-.1-.26a.31.31 0 0 0-.21-.08h-2.54a.31.31 0 0 0-.29.33v.25a3.9 3.9 0 0 1-1.52 3.09l-.13.1a3.91 3.91 0 0 1-3.63.59a.22.22 0 0 0-.14 0a.28.28 0 0 0-.12.15L4 11.12a.36.36 0 0 0 .22.45Z"/><path d="M12 15.5a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5Zm0-5a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5Z"/></g></g></g>`
	settingsTwoFillInnerSVG               = `<g id="evaSettings2Fill0"><g id="evaSettings2Fill1"><g id="evaSettings2Fill2" fill="currentColor"><circle cx="12" cy="12" r="1.5"/><path d="M20.32 9.37h-1.09c-.14 0-.24-.11-.3-.26a.34.34 0 0 1 0-.37l.81-.74a1.63 1.63 0 0 0 .5-1.18a1.67 1.67 0 0 0-.5-1.19L18.4 4.26a1.67 1.67 0 0 0-2.37 0l-.77.74a.38.38 0 0 1-.41 0a.34.34 0 0 1-.22-.29V3.68A1.68 1.68 0 0 0 13 2h-1.94a1.69 1.69 0 0 0-1.69 1.68v1.09c0 .14-.11.24-.26.3a.34.34 0 0 1-.37 0L8 4.26a1.72 1.72 0 0 0-1.19-.5a1.65 1.65 0 0 0-1.18.5L4.26 5.6a1.67 1.67 0 0 0 0 2.4l.74.74a.38.38 0 0 1 0 .41a.34.34 0 0 1-.29.22H3.68A1.68 1.68 0 0 0 2 11.05v1.89a1.69 1.69 0 0 0 1.68 1.69h1.09c.14 0 .24.11.3.26a.34.34 0 0 1 0 .37l-.81.74a1.72 1.72 0 0 0-.5 1.19a1.66 1.66 0 0 0 .5 1.19l1.34 1.36a1.67 1.67 0 0 0 2.37 0l.77-.74a.38.38 0 0 1 .41 0a.34.34 0 0 1 .22.29v1.09A1.68 1.68 0 0 0 11.05 22h1.89a1.69 1.69 0 0 0 1.69-1.68v-1.09c0-.14.11-.24.26-.3a.34.34 0 0 1 .37 0l.76.77a1.72 1.72 0 0 0 1.19.5a1.65 1.65 0 0 0 1.18-.5l1.34-1.34a1.67 1.67 0 0 0 0-2.37l-.73-.73a.34.34 0 0 1 0-.37a.34.34 0 0 1 .29-.22h1.09A1.68 1.68 0 0 0 22 13v-1.94a1.69 1.69 0 0 0-1.68-1.69ZM12 15.5a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5Z"/></g></g></g>`
	settingsTwoOutlineInnerSVG            = `<g id="evaSettings2Outline0"><g id="evaSettings2Outline1"><g id="evaSettings2Outline2" fill="currentColor"><path id="evaSettings2Outline3" d="M12.94 22h-1.89a1.68 1.68 0 0 1-1.68-1.68v-1.09a.34.34 0 0 0-.22-.29a.38.38 0 0 0-.41 0l-.74.8a1.67 1.67 0 0 1-2.37 0L4.26 18.4a1.66 1.66 0 0 1-.5-1.19a1.72 1.72 0 0 1 .5-1.21l.74-.74a.34.34 0 0 0 0-.37c-.06-.15-.16-.26-.3-.26H3.68A1.69 1.69 0 0 1 2 12.94v-1.89a1.68 1.68 0 0 1 1.68-1.68h1.09a.34.34 0 0 0 .29-.22a.38.38 0 0 0 0-.41L4.26 8a1.67 1.67 0 0 1 0-2.37L5.6 4.26a1.65 1.65 0 0 1 1.18-.5a1.72 1.72 0 0 1 1.22.5l.74.74a.34.34 0 0 0 .37 0c.15-.06.26-.16.26-.3V3.68A1.69 1.69 0 0 1 11.06 2H13a1.68 1.68 0 0 1 1.68 1.68v1.09a.34.34 0 0 0 .22.29a.38.38 0 0 0 .41 0l.69-.8a1.67 1.67 0 0 1 2.37 0l1.37 1.34a1.67 1.67 0 0 1 .5 1.19a1.63 1.63 0 0 1-.5 1.21l-.74.74a.34.34 0 0 0 0 .37c.06.15.16.26.3.26h1.09A1.69 1.69 0 0 1 22 11.06V13a1.68 1.68 0 0 1-1.68 1.68h-1.09a.34.34 0 0 0-.29.22a.34.34 0 0 0 0 .37l.77.77a1.67 1.67 0 0 1 0 2.37l-1.31 1.33a1.65 1.65 0 0 1-1.18.5a1.72 1.72 0 0 1-1.19-.5l-.77-.74a.34.34 0 0 0-.37 0c-.15.06-.26.16-.26.3v1.09A1.69 1.69 0 0 1 12.94 22Zm-1.57-2h1.26v-.77a2.33 2.33 0 0 1 1.46-2.14a2.36 2.36 0 0 1 2.59.47l.54.54l.88-.88l-.54-.55a2.34 2.34 0 0 1-.48-2.56a2.33 2.33 0 0 1 2.14-1.45H20v-1.29h-.77a2.33 2.33 0 0 1-2.14-1.46a2.36 2.36 0 0 1 .47-2.59l.54-.54l-.88-.88l-.55.54a2.39 2.39 0 0 1-4-1.67V4h-1.3v.77a2.33 2.33 0 0 1-1.46 2.14a2.36 2.36 0 0 1-2.59-.47l-.54-.54l-.88.88l.54.55a2.39 2.39 0 0 1-1.67 4H4v1.26h.77a2.33 2.33 0 0 1 2.14 1.46a2.36 2.36 0 0 1-.47 2.59l-.54.54l.88.88l.55-.54a2.39 2.39 0 0 1 4 1.67Z"/><path d="M12 15.5a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5Zm0-5a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5Z"/></g></g></g>`
	shakeFillInnerSVG                     = `<g id="evaShakeFill0"><g id="evaShakeFill1"><g id="evaShakeFill2" fill="currentColor"><path d="M5.5 18a1 1 0 0 1-.64-.24A8.81 8.81 0 0 1 1.5 11a8.81 8.81 0 0 1 3.36-6.76a1 1 0 1 1 1.28 1.52A6.9 6.9 0 0 0 3.5 11a6.9 6.9 0 0 0 2.64 5.24a1 1 0 0 1 .13 1.4a1 1 0 0 1-.77.36ZM12 7a4.09 4.09 0 0 1 1 .14V3a1 1 0 0 0-2 0v4.14A4.09 4.09 0 0 1 12 7Zm0 8a4.09 4.09 0 0 1-1-.14V20a1 1 0 0 0 2 0v-5.14a4.09 4.09 0 0 1-1 .14Zm4 1a1 1 0 0 1-.77-.36a1 1 0 0 1 .13-1.4A4.28 4.28 0 0 0 17 11a4.28 4.28 0 0 0-1.64-3.24a1 1 0 1 1 1.28-1.52A6.2 6.2 0 0 1 19 11a6.2 6.2 0 0 1-2.36 4.76A1 1 0 0 1 16 16Z"/><path d="M8 16a1 1 0 0 1-.64-.24A6.2 6.2 0 0 1 5 11a6.2 6.2 0 0 1 2.36-4.76a1 1 0 1 1 1.28 1.52A4.28 4.28 0 0 0 7 11a4.28 4.28 0 0 0 1.64 3.24a1 1 0 0 1 .13 1.4A1 1 0 0 1 8 16Zm10.5 2a1 1 0 0 1-.77-.36a1 1 0 0 1 .13-1.4A6.9 6.9 0 0 0 20.5 11a6.9 6.9 0 0 0-2.64-5.24a1 1 0 1 1 1.28-1.52A8.81 8.81 0 0 1 22.5 11a8.81 8.81 0 0 1-3.36 6.76a1 1 0 0 1-.64.24ZM12 12a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm0-1Zm0 0Zm0 0Zm0 0Zm0 0Zm0 0Zm0 0Z"/></g></g></g>`
	shakeOutlineInnerSVG                  = `<g id="evaShakeOutline0"><g id="evaShakeOutline1"><g id="evaShakeOutline2" fill="currentColor"><path d="M5.5 18a1 1 0 0 1-.64-.24A8.81 8.81 0 0 1 1.5 11a8.81 8.81 0 0 1 3.36-6.76a1 1 0 1 1 1.28 1.52A6.9 6.9 0 0 0 3.5 11a6.9 6.9 0 0 0 2.64 5.24a1 1 0 0 1 .13 1.4a1 1 0 0 1-.77.36ZM12 7a4.09 4.09 0 0 1 1 .14V3a1 1 0 0 0-2 0v4.14A4.09 4.09 0 0 1 12 7Zm0 8a4.09 4.09 0 0 1-1-.14V20a1 1 0 0 0 2 0v-5.14a4.09 4.09 0 0 1-1 .14Zm4 1a1 1 0 0 1-.77-.36a1 1 0 0 1 .13-1.4A4.28 4.28 0 0 0 17 11a4.28 4.28 0 0 0-1.64-3.24a1 1 0 1 1 1.28-1.52A6.2 6.2 0 0 1 19 11a6.2 6.2 0 0 1-2.36 4.76A1 1 0 0 1 16 16Z"/><path d="M8 16a1 1 0 0 1-.64-.24A6.2 6.2 0 0 1 5 11a6.2 6.2 0 0 1 2.36-4.76a1 1 0 1 1 1.28 1.52A4.28 4.28 0 0 0 7 11a4.28 4.28 0 0 0 1.64 3.24a1 1 0 0 1 .13 1.4A1 1 0 0 1 8 16Zm10.5 2a1 1 0 0 1-.77-.36a1 1 0 0 1 .13-1.4A6.9 6.9 0 0 0 20.5 11a6.9 6.9 0 0 0-2.64-5.24a1 1 0 1 1 1.28-1.52A8.81 8.81 0 0 1 22.5 11a8.81 8.81 0 0 1-3.36 6.76a1 1 0 0 1-.64.24ZM12 12a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm0-1Zm0 0Zm0 0Zm0 0Zm0 0Zm0 0Zm0 0Z"/></g></g></g>`
	shareFillInnerSVG                     = `<g id="evaShareFill0"><g id="evaShareFill1"><path id="evaShareFill2" fill="currentColor" d="M18 15a3 3 0 0 0-2.1.86L8 12.34v-.67l7.9-3.53A3 3 0 1 0 15 6v.34L7.1 9.86a3 3 0 1 0 0 4.28l7.9 3.53V18a3 3 0 1 0 3-3Z"/></g></g>`
	shareOutlineInnerSVG                  = `<g id="evaShareOutline0"><g id="evaShareOutline1"><path id="evaShareOutline2" fill="currentColor" d="M18 15a3 3 0 0 0-2.1.86L8 12.34v-.67l7.9-3.53A3 3 0 1 0 15 6v.34L7.1 9.86a3 3 0 1 0 0 4.28l7.9 3.53V18a3 3 0 1 0 3-3Zm0-10a1 1 0 1 1-1 1a1 1 0 0 1 1-1ZM5 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm13 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`
	shieldFillInnerSVG                    = `<g id="evaShieldFill0"><g id="evaShieldFill1"><path id="evaShieldFill2" fill="currentColor" d="M12 21.85a2 2 0 0 1-1-.25l-.3-.17A15.17 15.17 0 0 1 3 8.23v-.14a2 2 0 0 1 1-1.75l7-3.94a2 2 0 0 1 2 0l7 3.94a2 2 0 0 1 1 1.75v.14a15.17 15.17 0 0 1-7.72 13.2l-.3.17a2 2 0 0 1-.98.25Z"/></g></g>`
	shieldOffFillInnerSVG                 = `<g id="evaShieldOffFill0"><g id="evaShieldOffFill1"><path id="evaShieldOffFill2" fill="currentColor" d="M3.73 6.56A2 2 0 0 0 3 8.09v.14a15.17 15.17 0 0 0 7.72 13.2l.3.17a2 2 0 0 0 2 0l.3-.17a15.22 15.22 0 0 0 3-2.27ZM18.84 16A15.08 15.08 0 0 0 21 8.23v-.14a2 2 0 0 0-1-1.75L13 2.4a2 2 0 0 0-2 0L7.32 4.49ZM4.71 3.29a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	shieldOffOutlineInnerSVG              = `<g id="evaShieldOffOutline0"><g id="evaShieldOffOutline1"><path id="evaShieldOffOutline2" fill="currentColor" d="M4.71 3.29a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm7.59 16.39l-.3.17l-.3-.17A13.15 13.15 0 0 1 5 8.23v-.14L5.16 8L3.73 6.56A2 2 0 0 0 3 8.09v.14a15.17 15.17 0 0 0 7.72 13.2l.3.17a2 2 0 0 0 2 0l.3-.17a15.22 15.22 0 0 0 3-2.27l-1.42-1.42a12.56 12.56 0 0 1-2.6 1.94ZM20 6.34L13 2.4a2 2 0 0 0-2 0L7.32 4.49L8.78 6L12 4.15l7 3.94v.14a13 13 0 0 1-1.63 6.31L18.84 16A15.08 15.08 0 0 0 21 8.23v-.14a2 2 0 0 0-1-1.75Z"/></g></g>`
	shieldOutlineInnerSVG                 = `<g id="evaShieldOutline0"><g id="evaShieldOutline1"><path id="evaShieldOutline2" fill="currentColor" d="M12 21.85a2 2 0 0 1-1-.25l-.3-.17A15.17 15.17 0 0 1 3 8.23v-.14a2 2 0 0 1 1-1.75l7-3.94a2 2 0 0 1 2 0l7 3.94a2 2 0 0 1 1 1.75v.14a15.17 15.17 0 0 1-7.72 13.2l-.3.17a2 2 0 0 1-.98.25Zm0-17.7L5 8.09v.14a13.15 13.15 0 0 0 6.7 11.45l.3.17l.3-.17A13.15 13.15 0 0 0 19 8.23v-.14Z"/></g></g>`
	shoppingBagFillInnerSVG               = `<g id="evaShoppingBagFill0"><g id="evaShoppingBagFill1"><path id="evaShoppingBagFill2" fill="currentColor" d="m20.12 6.71l-2.83-2.83A3 3 0 0 0 15.17 3H8.83a3 3 0 0 0-2.12.88L3.88 6.71A3 3 0 0 0 3 8.83V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.83a3 3 0 0 0-.88-2.12ZM12 16a4 4 0 0 1-4-4a1 1 0 0 1 2 0a2 2 0 0 0 4 0a1 1 0 0 1 2 0a4 4 0 0 1-4 4ZM6.41 7l1.71-1.71A1.05 1.05 0 0 1 8.83 5h6.34a1.05 1.05 0 0 1 .71.29L17.59 7Z"/></g></g>`
	shoppingBagOutlineInnerSVG            = `<g id="evaShoppingBagOutline0"><g id="evaShoppingBagOutline1"><g id="evaShoppingBagOutline2" fill="currentColor"><path d="m20.12 6.71l-2.83-2.83A3 3 0 0 0 15.17 3H8.83a3 3 0 0 0-2.12.88L3.88 6.71A3 3 0 0 0 3 8.83V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.83a3 3 0 0 0-.88-2.12Zm-12-1.42A1.05 1.05 0 0 1 8.83 5h6.34a1.05 1.05 0 0 1 .71.29L17.59 7H6.41ZM18 19H6a1 1 0 0 1-1-1V9h14v9a1 1 0 0 1-1 1Z"/><path d="M15 11a1 1 0 0 0-1 1a2 2 0 0 1-4 0a1 1 0 0 0-2 0a4 4 0 0 0 8 0a1 1 0 0 0-1-1Z"/></g></g></g>`
	shoppingCartFillInnerSVG              = `<g id="evaShoppingCartFill0"><g id="evaShoppingCartFill1"><g id="evaShoppingCartFill2" fill="currentColor"><path d="M21.08 7a2 2 0 0 0-1.7-1H6.58L6 3.74A1 1 0 0 0 5 3H3a1 1 0 0 0 0 2h1.24L7 15.26A1 1 0 0 0 8 16h9a1 1 0 0 0 .89-.55l3.28-6.56A2 2 0 0 0 21.08 7Z"/><circle cx="7.5" cy="19.5" r="1.5"/><circle cx="17.5" cy="19.5" r="1.5"/></g></g></g>`
	shoppingCartOutlineInnerSVG           = `<g id="evaShoppingCartOutline0"><g id="evaShoppingCartOutline1"><g id="evaShoppingCartOutline2" fill="currentColor"><path d="M21.08 7a2 2 0 0 0-1.7-1H6.58L6 3.74A1 1 0 0 0 5 3H3a1 1 0 0 0 0 2h1.24L7 15.26A1 1 0 0 0 8 16h9a1 1 0 0 0 .89-.55l3.28-6.56A2 2 0 0 0 21.08 7Zm-4.7 7H8.76L7.13 8h12.25Z"/><circle cx="7.5" cy="19.5" r="1.5"/><circle cx="17.5" cy="19.5" r="1.5"/></g></g></g>`
	shuffleFillInnerSVG                   = `<g id="evaShuffleFill0"><g id="evaShuffleFill1"><g id="evaShuffleFill2" fill="currentColor"><path d="M18 9.31a1 1 0 0 0 1 1a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-4.3a1 1 0 0 0-1 1a1 1 0 0 0 1 1h1.89L12 10.59L6.16 4.76a1 1 0 0 0-1.41 1.41L10.58 12l-6.29 6.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L18 7.42Z"/><path d="M19 13.68a1 1 0 0 0-1 1v1.91l-2.78-2.79a1 1 0 0 0-1.42 1.42L16.57 18h-1.88a1 1 0 0 0 0 2H19a1 1 0 0 0 1-1.11v-4.21a1 1 0 0 0-1-1Z"/></g></g></g>`
	shuffleOutlineInnerSVG                = `<g id="evaShuffleOutline0"><g id="evaShuffleOutline1"><g id="evaShuffleOutline2" fill="currentColor"><path d="M18 9.31a1 1 0 0 0 1 1a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-4.3a1 1 0 0 0-1 1a1 1 0 0 0 1 1h1.89L12 10.59L6.16 4.76a1 1 0 0 0-1.41 1.41L10.58 12l-6.29 6.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L18 7.42Z"/><path d="M19 13.68a1 1 0 0 0-1 1v1.91l-2.78-2.79a1 1 0 0 0-1.42 1.42L16.57 18h-1.88a1 1 0 0 0 0 2H19a1 1 0 0 0 1-1.11v-4.21a1 1 0 0 0-1-1Z"/></g></g></g>`
	shuffleTwoFillInnerSVG                = `<g id="evaShuffle2Fill0"><g id="evaShuffle2Fill1"><path id="evaShuffle2Fill2" fill="currentColor" d="M18.71 14.29a1 1 0 0 0-1.42 1.42l.29.29H16a4 4 0 0 1 0-8h1.59l-.3.29a1 1 0 0 0 0 1.42A1 1 0 0 0 18 10a1 1 0 0 0 .71-.29l2-2A1 1 0 0 0 21 7a1 1 0 0 0-.29-.71l-2-2a1 1 0 0 0-1.42 1.42l.29.29H16a6 6 0 0 0-5 2.69A6 6 0 0 0 6 6H4a1 1 0 0 0 0 2h2a4 4 0 0 1 0 8H4a1 1 0 0 0 0 2h2a6 6 0 0 0 5-2.69A6 6 0 0 0 16 18h1.59l-.3.29a1 1 0 0 0 0 1.42A1 1 0 0 0 18 20a1 1 0 0 0 .71-.29l2-2A1 1 0 0 0 21 17a1 1 0 0 0-.29-.71Z"/></g></g>`
	shuffleTwoOutlineInnerSVG             = `<g id="evaShuffle2Outline0"><g id="evaShuffle2Outline1"><path id="evaShuffle2Outline2" fill="currentColor" d="M18.71 14.29a1 1 0 0 0-1.42 1.42l.29.29H16a4 4 0 0 1 0-8h1.59l-.3.29a1 1 0 0 0 0 1.42A1 1 0 0 0 18 10a1 1 0 0 0 .71-.29l2-2A1 1 0 0 0 21 7a1 1 0 0 0-.29-.71l-2-2a1 1 0 0 0-1.42 1.42l.29.29H16a6 6 0 0 0-5 2.69A6 6 0 0 0 6 6H4a1 1 0 0 0 0 2h2a4 4 0 0 1 0 8H4a1 1 0 0 0 0 2h2a6 6 0 0 0 5-2.69A6 6 0 0 0 16 18h1.59l-.3.29a1 1 0 0 0 0 1.42A1 1 0 0 0 18 20a1 1 0 0 0 .71-.29l2-2A1 1 0 0 0 21 17a1 1 0 0 0-.29-.71Z"/></g></g>`
	skipBackFillInnerSVG                  = `<g id="evaSkipBackFill0"><g id="evaSkipBackFill1"><path id="evaSkipBackFill2" fill="currentColor" d="M16.45 6.2a2.1 2.1 0 0 0-2.21.26l-5.1 4.21l-.14.15V7a1 1 0 0 0-2 0v10a1 1 0 0 0 2 0v-3.82l.14.15l5.1 4.21a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59V7.79a1.76 1.76 0 0 0-1.05-1.59Z"/></g></g>`
	skipBackOutlineInnerSVG               = `<g id="evaSkipBackOutline0"><g id="evaSkipBackOutline1"><path id="evaSkipBackOutline2" fill="currentColor" d="M16.45 6.2a2.1 2.1 0 0 0-2.21.26l-5.1 4.21l-.14.15V7a1 1 0 0 0-2 0v10a1 1 0 0 0 2 0v-3.82l.14.15l5.1 4.21a2.06 2.06 0 0 0 1.3.46a2.23 2.23 0 0 0 .91-.2a1.76 1.76 0 0 0 1.05-1.59V7.79a1.76 1.76 0 0 0-1.05-1.59ZM15.5 16l-4.82-4l4.82-3.91Z"/></g></g>`
	skipForwardFillInnerSVG               = `<g id="evaSkipForwardFill0"><g id="evaSkipForwardFill1"><path id="evaSkipForwardFill2" fill="currentColor" d="M16 6a1 1 0 0 0-1 1v3.82l-.14-.15l-5.1-4.21a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1 1.59v8.42a1.76 1.76 0 0 0 1 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l5.1-4.21l.14-.15V17a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1Z"/></g></g>`
	skipForwardOutlineInnerSVG            = `<g id="evaSkipForwardOutline0"><g id="evaSkipForwardOutline1"><path id="evaSkipForwardOutline2" fill="currentColor" d="M16 6a1 1 0 0 0-1 1v3.82l-.14-.15l-5.1-4.21a2.1 2.1 0 0 0-2.21-.26a1.76 1.76 0 0 0-1 1.59v8.42a1.76 1.76 0 0 0 1 1.59a2.23 2.23 0 0 0 .91.2a2.06 2.06 0 0 0 1.3-.46l5.1-4.21l.14-.15V17a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1Zm-7.5 9.91V8l4.82 4Z"/></g></g>`
	slashFillInnerSVG                     = `<g id="evaSlashFill0"><g id="evaSlashFill1"><path id="evaSlashFill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm8 10a7.92 7.92 0 0 1-1.69 4.9L7.1 5.69A7.92 7.92 0 0 1 12 4a8 8 0 0 1 8 8ZM4 12a7.92 7.92 0 0 1 1.69-4.9L16.9 18.31A7.92 7.92 0 0 1 12 20a8 8 0 0 1-8-8Z"/></g></g>`
	slashOutlineInnerSVG                  = `<g id="evaSlashOutline0"><g id="evaSlashOutline1"><path id="evaSlashOutline2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm8 10a7.92 7.92 0 0 1-1.69 4.9L7.1 5.69A7.92 7.92 0 0 1 12 4a8 8 0 0 1 8 8ZM4 12a7.92 7.92 0 0 1 1.69-4.9L16.9 18.31A7.92 7.92 0 0 1 12 20a8 8 0 0 1-8-8Z"/></g></g>`
	smartphoneFillInnerSVG                = `<g id="evaSmartphoneFill0"><g id="evaSmartphoneFill1"><path id="evaSmartphoneFill2" fill="currentColor" d="M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm-5 16a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 12 18Zm2.5-10h-5a1 1 0 0 1 0-2h5a1 1 0 0 1 0 2Z"/></g></g>`
	smartphoneOutlineInnerSVG             = `<g id="evaSmartphoneOutline0"><g id="evaSmartphoneOutline1"><g id="evaSmartphoneOutline2" fill="currentColor"><path d="M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm1 17a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/><circle cx="12" cy="16.5" r="1.5"/><path d="M14.5 6h-5a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2Z"/></g></g></g>`
	smilingFaceFillInnerSVG               = `<g id="evaSmilingFaceFill0"><g id="evaSmilingFaceFill1"><path id="evaSmilingFaceFill2" fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm0 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm5 9a5 5 0 0 1-10 0Z"/></g></g>`
	smilingFaceOutlineInnerSVG            = `<g id="evaSmilingFaceOutline0"><g id="evaSmilingFaceOutline1"><path id="evaSmilingFaceOutline2" fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm0 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm5 9a5 5 0 0 1-10 0Z"/></g></g>`
	speakerFillInnerSVG                   = `<g id="evaSpeakerFill0"><g id="evaSpeakerFill1"><g id="evaSpeakerFill2" fill="currentColor"><circle cx="12" cy="15.5" r="1.5"/><circle cx="12" cy="8" r="1"/><path d="M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm-5 3a3 3 0 1 1-3 3a3 3 0 0 1 3-3Zm0 14a3.5 3.5 0 1 1 3.5-3.5A3.5 3.5 0 0 1 12 19Z"/></g></g></g>`
	speakerOutlineInnerSVG                = `<g id="evaSpeakerOutline0"><g id="evaSpeakerOutline1"><g id="evaSpeakerOutline2" fill="currentColor"><path d="M12 11a3 3 0 1 0-3-3a3 3 0 0 0 3 3Zm0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm0 5a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 12 12Zm0 5a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 12 17Z"/><path d="M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm1 17a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/></g></g></g>`
	squareFillInnerSVG                    = `<g id="evaSquareFill0"><g id="evaSquareFill1"><path id="evaSquareFill2" fill="currentColor" d="M18 21H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3ZM6 5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1Z"/></g></g>`
	squareOutlineInnerSVG                 = `<g id="evaSquareOutline0"><g id="evaSquareOutline1"><path id="evaSquareOutline2" fill="currentColor" d="M18 21H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3ZM6 5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1Z"/></g></g>`
	starFillInnerSVG                      = `<g id="evaStarFill0"><g id="evaStarFill1"><path id="evaStarFill2" fill="currentColor" d="M17.56 21a1 1 0 0 1-.46-.11L12 18.22l-5.1 2.67a1 1 0 0 1-1.45-1.06l1-5.63l-4.12-4a1 1 0 0 1-.25-1a1 1 0 0 1 .81-.68l5.7-.83l2.51-5.13a1 1 0 0 1 1.8 0l2.54 5.12l5.7.83a1 1 0 0 1 .81.68a1 1 0 0 1-.25 1l-4.12 4l1 5.63a1 1 0 0 1-.4 1a1 1 0 0 1-.62.18Z"/></g></g>`
	starOutlineInnerSVG                   = `<g id="evaStarOutline0"><g id="evaStarOutline1"><path id="evaStarOutline2" fill="currentColor" d="M17.56 21a1 1 0 0 1-.46-.11L12 18.22l-5.1 2.67a1 1 0 0 1-1.45-1.06l1-5.63l-4.12-4a1 1 0 0 1-.25-1a1 1 0 0 1 .81-.68l5.7-.83l2.51-5.13a1 1 0 0 1 1.8 0l2.54 5.12l5.7.83a1 1 0 0 1 .81.68a1 1 0 0 1-.25 1l-4.12 4l1 5.63a1 1 0 0 1-.4 1a1 1 0 0 1-.62.18ZM12 16.1a.92.92 0 0 1 .46.11l3.77 2l-.72-4.21a1 1 0 0 1 .29-.89l3-2.93l-4.2-.62a1 1 0 0 1-.71-.56L12 5.25L10.11 9a1 1 0 0 1-.75.54l-4.2.62l3 2.93a1 1 0 0 1 .29.89l-.72 4.16l3.77-2a.92.92 0 0 1 .5-.04Z"/></g></g>`
	stopCircleFillInnerSVG                = `<g id="evaStopCircleFill0"><g id="evaStopCircleFill1"><g id="evaStopCircleFill2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm4 12.75A1.25 1.25 0 0 1 14.75 16h-5.5A1.25 1.25 0 0 1 8 14.75v-5.5A1.25 1.25 0 0 1 9.25 8h5.5A1.25 1.25 0 0 1 16 9.25Z"/><path d="M10 10h4v4h-4z"/></g></g></g>`
	stopCircleOutlineInnerSVG             = `<g id="evaStopCircleOutline0"><g id="evaStopCircleOutline1"><g id="evaStopCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M14.75 8h-5.5A1.25 1.25 0 0 0 8 9.25v5.5A1.25 1.25 0 0 0 9.25 16h5.5A1.25 1.25 0 0 0 16 14.75v-5.5A1.25 1.25 0 0 0 14.75 8ZM14 14h-4v-4h4Z"/></g></g></g>`
	sunFillInnerSVG                       = `<g id="evaSunFill0"><g id="evaSunFill1"><path id="evaSunFill2" fill="currentColor" d="M12 6a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1Zm9 5h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2ZM6 12a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1Zm.22-7a1 1 0 0 0-1.39 1.47l1.44 1.39a1 1 0 0 0 .73.28a1 1 0 0 0 .72-.31a1 1 0 0 0 0-1.41ZM17 8.14a1 1 0 0 0 .69-.28l1.44-1.39A1 1 0 0 0 17.78 5l-1.44 1.42a1 1 0 0 0 0 1.41a1 1 0 0 0 .66.31ZM12 18a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1Zm5.73-1.86a1 1 0 0 0-1.39 1.44L17.78 19a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.3a1 1 0 0 0 0-1.42Zm-11.46 0l-1.44 1.39a1 1 0 0 0 0 1.42a1 1 0 0 0 .72.3a1 1 0 0 0 .67-.25l1.44-1.39a1 1 0 0 0-1.39-1.44ZM12 8a4 4 0 1 0 4 4a4 4 0 0 0-4-4Z"/></g></g>`
	sunOutlineInnerSVG                    = `<g id="evaSunOutline0"><g id="evaSunOutline1"><path id="evaSunOutline2" fill="currentColor" d="M12 6a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1Zm9 5h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2ZM6 12a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1Zm.22-7a1 1 0 0 0-1.39 1.47l1.44 1.39a1 1 0 0 0 .73.28a1 1 0 0 0 .72-.31a1 1 0 0 0 0-1.41ZM17 8.14a1 1 0 0 0 .69-.28l1.44-1.39A1 1 0 0 0 17.78 5l-1.44 1.42a1 1 0 0 0 0 1.41a1 1 0 0 0 .66.31ZM12 18a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1Zm5.73-1.86a1 1 0 0 0-1.39 1.44L17.78 19a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.3a1 1 0 0 0 0-1.42Zm-11.46 0l-1.44 1.39a1 1 0 0 0 0 1.42a1 1 0 0 0 .72.3a1 1 0 0 0 .67-.25l1.44-1.39a1 1 0 0 0-1.39-1.44ZM12 8a4 4 0 1 0 4 4a4 4 0 0 0-4-4Zm0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/></g></g>`
	swapFillInnerSVG                      = `<g id="evaSwapFill0"><g id="evaSwapFill1"><path id="evaSwapFill2" fill="currentColor" d="M4 9h13l-1.6 1.2a1 1 0 0 0-.2 1.4a1 1 0 0 0 .8.4a1 1 0 0 0 .6-.2l4-3a1 1 0 0 0 0-1.59l-3.86-3a1 1 0 0 0-1.23 1.58L17.08 7H4a1 1 0 0 0 0 2Zm16 7H7l1.6-1.2a1 1 0 0 0-1.2-1.6l-4 3a1 1 0 0 0 0 1.59l3.86 3a1 1 0 0 0 .61.21a1 1 0 0 0 .79-.39a1 1 0 0 0-.17-1.4L6.92 18H20a1 1 0 0 0 0-2Z"/></g></g>`
	swapOutlineInnerSVG                   = `<g id="evaSwapOutline0"><g id="evaSwapOutline1"><path id="evaSwapOutline2" fill="currentColor" d="M4 9h13l-1.6 1.2a1 1 0 0 0-.2 1.4a1 1 0 0 0 .8.4a1 1 0 0 0 .6-.2l4-3a1 1 0 0 0 0-1.59l-3.86-3a1 1 0 0 0-1.23 1.58L17.08 7H4a1 1 0 0 0 0 2Zm16 7H7l1.6-1.2a1 1 0 0 0-1.2-1.6l-4 3a1 1 0 0 0 0 1.59l3.86 3a1 1 0 0 0 .61.21a1 1 0 0 0 .79-.39a1 1 0 0 0-.17-1.4L6.92 18H20a1 1 0 0 0 0-2Z"/></g></g>`
	syncFillInnerSVG                      = `<g id="evaSyncFill0"><g id="evaSyncFill1"><path id="evaSyncFill2" fill="currentColor" d="M21.66 10.37a.62.62 0 0 0 .07-.19l.75-4a1 1 0 0 0-2-.36l-.37 2a9.22 9.22 0 0 0-16.58.84a1 1 0 0 0 .55 1.3a1 1 0 0 0 1.31-.55A7.08 7.08 0 0 1 12.07 5a7.17 7.17 0 0 1 6.24 3.58l-1.65-.27a1 1 0 1 0-.32 2l4.25.71h.16a.93.93 0 0 0 .34-.06a.33.33 0 0 0 .1-.06a.78.78 0 0 0 .2-.11l.08-.1a1.07 1.07 0 0 0 .14-.16a.58.58 0 0 0 .05-.16Zm-1.78 3.7a1 1 0 0 0-1.31.56A7.08 7.08 0 0 1 11.93 19a7.17 7.17 0 0 1-6.24-3.58l1.65.27h.16a1 1 0 0 0 .16-2L3.41 13a.91.91 0 0 0-.33 0H3a1.15 1.15 0 0 0-.32.14a1 1 0 0 0-.18.18l-.09.1a.84.84 0 0 0-.07.19a.44.44 0 0 0-.07.17l-.75 4a1 1 0 0 0 .8 1.22h.18a1 1 0 0 0 1-.82l.37-2a9.22 9.22 0 0 0 16.58-.83a1 1 0 0 0-.57-1.28Z"/></g></g>`
	syncOutlineInnerSVG                   = `<g id="evaSyncOutline0"><g id="evaSyncOutline1"><path id="evaSyncOutline2" fill="currentColor" d="M21.66 10.37a.62.62 0 0 0 .07-.19l.75-4a1 1 0 0 0-2-.36l-.37 2a9.22 9.22 0 0 0-16.58.84a1 1 0 0 0 .55 1.3a1 1 0 0 0 1.31-.55A7.08 7.08 0 0 1 12.07 5a7.17 7.17 0 0 1 6.24 3.58l-1.65-.27a1 1 0 1 0-.32 2l4.25.71h.16a.93.93 0 0 0 .34-.06a.33.33 0 0 0 .1-.06a.78.78 0 0 0 .2-.11l.08-.1a1.07 1.07 0 0 0 .14-.16a.58.58 0 0 0 .05-.16Zm-1.78 3.7a1 1 0 0 0-1.31.56A7.08 7.08 0 0 1 11.93 19a7.17 7.17 0 0 1-6.24-3.58l1.65.27h.16a1 1 0 0 0 .16-2L3.41 13a.91.91 0 0 0-.33 0H3a1.15 1.15 0 0 0-.32.14a1 1 0 0 0-.18.18l-.09.1a.84.84 0 0 0-.07.19a.44.44 0 0 0-.07.17l-.75 4a1 1 0 0 0 .8 1.22h.18a1 1 0 0 0 1-.82l.37-2a9.22 9.22 0 0 0 16.58-.83a1 1 0 0 0-.57-1.28Z"/></g></g>`
	textFillInnerSVG                      = `<g id="evaTextFill0"><g id="evaTextFill1"><path id="evaTextFill2" fill="currentColor" d="M20 4H4a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0V6h6v13H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2V6h6v2a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1Z"/></g></g>`
	textOutlineInnerSVG                   = `<g id="evaTextOutline0"><g id="evaTextOutline1"><path id="evaTextOutline2" fill="currentColor" d="M20 4H4a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0V6h6v13H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2V6h6v2a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1Z"/></g></g>`
	thermometerFillInnerSVG               = `<g id="evaThermometerFill0"><g id="evaThermometerFill1"><path id="evaThermometerFill2" fill="currentColor" d="M12 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm1-12.46V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 12 4a1 1 0 0 0-1 1v4.54Z"/></g></g>`
	thermometerMinusFillInnerSVG          = `<g id="evaThermometerMinusFill0"><g id="evaThermometerMinusFill1"><g id="evaThermometerMinusFill2" fill="currentColor"><rect width="6" height="2" x="2" y="5" rx="1" ry="1"/><path d="M14 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm1-12.46V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 14 4a1 1 0 0 0-1 1v4.54Z"/></g></g></g>`
	thermometerMinusOutlineInnerSVG       = `<g id="evaThermometerMinusOutline0"><g id="evaThermometerMinusOutline1"><g id="evaThermometerMinusOutline2" fill="currentColor"><rect width="6" height="2" x="2" y="5" rx="1" ry="1"/><path d="M14 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm0-18a1 1 0 0 0-1 1v8.54a1 1 0 0 1-.5.87A3 3 0 0 0 11 17a3 3 0 0 0 6 0a3 3 0 0 0-1.5-2.59a1 1 0 0 1-.5-.87V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 14 4Z"/></g></g></g>`
	thermometerOutlineInnerSVG            = `<g id="evaThermometerOutline0"><g id="evaThermometerOutline1"><path id="evaThermometerOutline2" fill="currentColor" d="M12 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm0-18a1 1 0 0 0-1 1v8.54a1 1 0 0 1-.5.87A3 3 0 0 0 9 17a3 3 0 0 0 6 0a3 3 0 0 0-1.5-2.59a1 1 0 0 1-.5-.87V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 12 4Z"/></g></g>`
	thermometerPlusFillInnerSVG           = `<g id="evaThermometerPlusFill0"><g id="evaThermometerPlusFill1"><g id="evaThermometerPlusFill2" fill="currentColor"><rect width="6" height="2" x="2" y="5" rx="1" ry="1"/><rect width="6" height="2" x="2" y="5" rx="1" ry="1" transform="rotate(-90 5 6)"/><path d="M14 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm1-12.46V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 14 4a1 1 0 0 0-1 1v4.54Z"/></g></g></g>`
	thermometerPlusOutlineInnerSVG        = `<g id="evaThermometerPlusOutline0"><g id="evaThermometerPlusOutline1"><g id="evaThermometerPlusOutline2" fill="currentColor"><rect width="6" height="2" x="2" y="5" rx="1" ry="1"/><rect width="6" height="2" x="2" y="5" rx="1" ry="1" transform="rotate(-90 5 6)"/><path d="M14 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm0-18a1 1 0 0 0-1 1v8.54a1 1 0 0 1-.5.87A3 3 0 0 0 11 17a3 3 0 0 0 6 0a3 3 0 0 0-1.5-2.59a1 1 0 0 1-.5-.87V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 14 4Z"/></g></g></g>`
	toggleLeftFillInnerSVG                = `<g id="evaToggleLeftFill0"><g id="evaToggleLeftFill1"><g id="evaToggleLeftFill2" fill="currentColor"><path d="M15 5H9a7 7 0 0 0 0 14h6a7 7 0 0 0 0-14ZM9 15a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/><path d="M9 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/></g></g></g>`
	toggleLeftOutlineInnerSVG             = `<g id="evaToggleLeftOutline0"><g id="evaToggleLeftOutline1"><g id="evaToggleLeftOutline2" fill="currentColor"><path d="M15 5H9a7 7 0 0 0 0 14h6a7 7 0 0 0 0-14Zm0 12H9A5 5 0 0 1 9 7h6a5 5 0 0 1 0 10Z"/><path d="M9 9a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g></g>`
	toggleRightFillInnerSVG               = `<g id="evaToggleRightFill0"><g id="evaToggleRightFill1"><g id="evaToggleRightFill2" fill="currentColor"><circle cx="15" cy="12" r="1"/><path d="M15 5H9a7 7 0 0 0 0 14h6a7 7 0 0 0 0-14Zm0 10a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/></g></g></g>`
	toggleRightOutlineInnerSVG            = `<g id="evaToggleRightOutline0"><g id="evaToggleRightOutline1"><g id="evaToggleRightOutline2" fill="currentColor"><path d="M15 5H9a7 7 0 0 0 0 14h6a7 7 0 0 0 0-14Zm0 12H9A5 5 0 0 1 9 7h6a5 5 0 0 1 0 10Z"/><path d="M15 9a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g></g>`
	trashFillInnerSVG                     = `<g id="evaTrashFill0"><g id="evaTrashFill1"><path id="evaTrashFill2" fill="currentColor" d="M21 6h-5V4.33A2.42 2.42 0 0 0 13.5 2h-3A2.42 2.42 0 0 0 8 4.33V6H3a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8h1a1 1 0 0 0 0-2ZM10 4.33c0-.16.21-.33.5-.33h3c.29 0 .5.17.5.33V6h-4Z"/></g></g>`
	trashOutlineInnerSVG                  = `<g id="evaTrashOutline0"><g id="evaTrashOutline1"><path id="evaTrashOutline2" fill="currentColor" d="M21 6h-5V4.33A2.42 2.42 0 0 0 13.5 2h-3A2.42 2.42 0 0 0 8 4.33V6H3a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8h1a1 1 0 0 0 0-2ZM10 4.33c0-.16.21-.33.5-.33h3c.29 0 .5.17.5.33V6h-4ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V8h12Z"/></g></g>`
	trashTwoFillInnerSVG                  = `<g id="evaTrash2Fill0"><g id="evaTrash2Fill1"><path id="evaTrash2Fill2" fill="currentColor" d="M21 6h-5V4.33A2.42 2.42 0 0 0 13.5 2h-3A2.42 2.42 0 0 0 8 4.33V6H3a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8h1a1 1 0 0 0 0-2ZM10 16a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0Zm0-11.67c0-.16.21-.33.5-.33h3c.29 0 .5.17.5.33V6h-4ZM16 16a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0Z"/></g></g>`
	trashTwoOutlineInnerSVG               = `<g id="evaTrash2Outline0"><g id="evaTrash2Outline1"><g id="evaTrash2Outline2" fill="currentColor"><path d="M21 6h-5V4.33A2.42 2.42 0 0 0 13.5 2h-3A2.42 2.42 0 0 0 8 4.33V6H3a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8h1a1 1 0 0 0 0-2ZM10 4.33c0-.16.21-.33.5-.33h3c.29 0 .5.17.5.33V6h-4ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V8h12Z"/><path d="M9 17a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1Zm6 0a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1Z"/></g></g></g>`
	trendingDownFillInnerSVG              = `<g id="evaTrendingDownFill0"><g id="evaTrendingDownFill1"><path id="evaTrendingDownFill2" fill="currentColor" d="M21 12a1 1 0 0 0-2 0v2.3l-4.24-5a1 1 0 0 0-1.27-.21L9.22 11.7L4.77 6.36a1 1 0 1 0-1.54 1.28l5 6a1 1 0 0 0 1.28.22l4.28-2.57l4 4.71H15a1 1 0 0 0 0 2h5a1.1 1.1 0 0 0 .36-.07l.14-.08a1.19 1.19 0 0 0 .15-.09a.75.75 0 0 0 .14-.17a1.1 1.1 0 0 0 .09-.14a.64.64 0 0 0 .05-.17A.78.78 0 0 0 21 17Z"/></g></g>`
	trendingDownOutlineInnerSVG           = `<g id="evaTrendingDownOutline0"><g id="evaTrendingDownOutline1"><path id="evaTrendingDownOutline2" fill="currentColor" d="M21 12a1 1 0 0 0-2 0v2.3l-4.24-5a1 1 0 0 0-1.27-.21L9.22 11.7L4.77 6.36a1 1 0 1 0-1.54 1.28l5 6a1 1 0 0 0 1.28.22l4.28-2.57l4 4.71H15a1 1 0 0 0 0 2h5a1.1 1.1 0 0 0 .36-.07l.14-.08a1.19 1.19 0 0 0 .15-.09a.75.75 0 0 0 .14-.17a1.1 1.1 0 0 0 .09-.14a.64.64 0 0 0 .05-.17A.78.78 0 0 0 21 17Z"/></g></g>`
	trendingUpFillInnerSVG                = `<g id="evaTrendingUpFill0"><g id="evaTrendingUpFill1"><path id="evaTrendingUpFill2" fill="currentColor" d="M21 7a.78.78 0 0 0 0-.21a.64.64 0 0 0-.05-.17a1.1 1.1 0 0 0-.09-.14a.75.75 0 0 0-.14-.17l-.12-.07a.69.69 0 0 0-.19-.1h-.2A.7.7 0 0 0 20 6h-5a1 1 0 0 0 0 2h2.83l-4 4.71l-4.32-2.57a1 1 0 0 0-1.28.22l-5 6a1 1 0 0 0 .13 1.41A1 1 0 0 0 4 18a1 1 0 0 0 .77-.36l4.45-5.34l4.27 2.56a1 1 0 0 0 1.27-.21L19 9.7V12a1 1 0 0 0 2 0V7Z"/></g></g>`
	trendingUpOutlineInnerSVG             = `<g id="evaTrendingUpOutline0"><g id="evaTrendingUpOutline1"><path id="evaTrendingUpOutline2" fill="currentColor" d="M21 7a.78.78 0 0 0 0-.21a.64.64 0 0 0-.05-.17a1.1 1.1 0 0 0-.09-.14a.75.75 0 0 0-.14-.17l-.12-.07a.69.69 0 0 0-.19-.1h-.2A.7.7 0 0 0 20 6h-5a1 1 0 0 0 0 2h2.83l-4 4.71l-4.32-2.57a1 1 0 0 0-1.28.22l-5 6a1 1 0 0 0 .13 1.41A1 1 0 0 0 4 18a1 1 0 0 0 .77-.36l4.45-5.34l4.27 2.56a1 1 0 0 0 1.27-.21L19 9.7V12a1 1 0 0 0 2 0V7Z"/></g></g>`
	tvFillInnerSVG                        = `<g id="evaTvFill0"><g id="evaTvFill1"><path id="evaTvFill2" fill="currentColor" d="M18 6h-3.59l2.3-2.29a1 1 0 1 0-1.42-1.42L12 5.59l-3.29-3.3a1 1 0 1 0-1.42 1.42L9.59 6H6a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Zm1 13a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/></g></g>`
	tvOutlineInnerSVG                     = `<g id="evaTvOutline0"><g id="evaTvOutline1"><path id="evaTvOutline2" fill="currentColor" d="M18 6h-3.59l2.3-2.29a1 1 0 1 0-1.42-1.42L12 5.59l-3.29-3.3a1 1 0 1 0-1.42 1.42L9.59 6H6a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Zm1 13a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/></g></g>`
	twitterFillInnerSVG                   = `<g id="evaTwitterFill0"><g id="evaTwitterFill1"><path id="evaTwitterFill2" fill="currentColor" d="M8.08 20A11.07 11.07 0 0 0 19.52 9A8.09 8.09 0 0 0 21 6.16a.44.44 0 0 0-.62-.51a1.88 1.88 0 0 1-2.16-.38a3.89 3.89 0 0 0-5.58-.17A4.13 4.13 0 0 0 11.49 9C8.14 9.2 5.84 7.61 4 5.43a.43.43 0 0 0-.75.24a9.68 9.68 0 0 0 4.6 10.05A6.73 6.73 0 0 1 3.38 18a.45.45 0 0 0-.14.84A11 11 0 0 0 8.08 20"/></g></g>`
	twitterOutlineInnerSVG                = `<g id="evaTwitterOutline0"><g id="evaTwitterOutline1"><path id="evaTwitterOutline2" fill="currentColor" d="M8.51 20h-.08a10.87 10.87 0 0 1-4.65-1.09A1.38 1.38 0 0 1 3 17.47a1.41 1.41 0 0 1 1.16-1.18a6.63 6.63 0 0 0 2.54-.89a9.49 9.49 0 0 1-3.51-9.07a1.41 1.41 0 0 1 1-1.15a1.35 1.35 0 0 1 1.43.41a7.09 7.09 0 0 0 4.88 2.75a4.5 4.5 0 0 1 1.41-3.1a4.47 4.47 0 0 1 6.37.19a.7.7 0 0 0 .78.1A1.39 1.39 0 0 1 21 7.13a6.66 6.66 0 0 1-1.28 2.6A10.79 10.79 0 0 1 8.51 20Zm0-2h.08a8.79 8.79 0 0 0 9.09-8.59a1.32 1.32 0 0 1 .37-.85a5.19 5.19 0 0 0 .62-1a2.56 2.56 0 0 1-1.91-.85A2.45 2.45 0 0 0 15 6a2.5 2.5 0 0 0-1.79.69a2.53 2.53 0 0 0-.72 2.42l.26 1.14l-1.17.08a8.3 8.3 0 0 1-6.54-2.4a7.12 7.12 0 0 0 3.73 6.46l.95.54l-.63.9a5.62 5.62 0 0 1-2.68 1.92A8.34 8.34 0 0 0 8.5 18ZM19 6.65Z"/></g></g>`
	umbrellaFillInnerSVG                  = `<g id="evaUmbrellaFill0"><g id="evaUmbrellaFill1"><path id="evaUmbrellaFill2" fill="currentColor" d="M12 2A10 10 0 0 0 2 12a1 1 0 0 0 1 1h8v6a3 3 0 0 0 6 0a1 1 0 0 0-2 0a1 1 0 0 1-2 0v-6h8a1 1 0 0 0 1-1A10 10 0 0 0 12 2Z"/></g></g>`
	umbrellaOutlineInnerSVG               = `<g id="evaUmbrellaOutline0"><g id="evaUmbrellaOutline1"><path id="evaUmbrellaOutline2" fill="currentColor" d="M12 2A10 10 0 0 0 2 12a1 1 0 0 0 1 1h8v6a3 3 0 0 0 6 0a1 1 0 0 0-2 0a1 1 0 0 1-2 0v-6h8a1 1 0 0 0 1-1A10 10 0 0 0 12 2Zm-7.94 9a8 8 0 0 1 15.88 0Z"/></g></g>`
	undoFillInnerSVG                      = `<g id="evaUndoFill0"><g id="evaUndoFill1"><path id="evaUndoFill2" fill="currentColor" d="M20.22 21a1 1 0 0 1-1-.76a8.91 8.91 0 0 0-7.8-6.69v1.12a1.78 1.78 0 0 1-1.09 1.64A2 2 0 0 1 8.18 16l-5.06-4.41a1.76 1.76 0 0 1 0-2.68l5.06-4.42a2 2 0 0 1 2.18-.3a1.78 1.78 0 0 1 1.09 1.64V7A10.89 10.89 0 0 1 21.5 17.75a10.29 10.29 0 0 1-.31 2.49a1 1 0 0 1-1 .76Z"/></g></g>`
	undoOutlineInnerSVG                   = `<g id="evaUndoOutline0"><g id="evaUndoOutline1"><path id="evaUndoOutline2" fill="currentColor" d="M20.22 21a1 1 0 0 1-1-.76a8.91 8.91 0 0 0-7.8-6.69v1.12a1.78 1.78 0 0 1-1.09 1.64A2 2 0 0 1 8.18 16l-5.06-4.41a1.76 1.76 0 0 1 0-2.68l5.06-4.42a2 2 0 0 1 2.18-.3a1.78 1.78 0 0 1 1.09 1.64V7A10.89 10.89 0 0 1 21.5 17.75a10.29 10.29 0 0 1-.31 2.49a1 1 0 0 1-1 .76Zm-9.77-9.5a11.07 11.07 0 0 1 8.81 4.26A9 9 0 0 0 10.45 9a1 1 0 0 1-1-1V6.08l-4.82 4.17l4.82 4.21v-2a1 1 0 0 1 1-.96Z"/></g></g>`
	unlockFillInnerSVG                    = `<g id="evaUnlockFill0"><g id="evaUnlockFill1"><g id="evaUnlockFill2" fill="currentColor"><circle cx="12" cy="15" r="1"/><path d="M17 8h-7V6a2 2 0 0 1 4 0a1 1 0 0 0 2 0a4 4 0 0 0-8 0v2H7a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm-5 10a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/></g></g></g>`
	unlockOutlineInnerSVG                 = `<g id="evaUnlockOutline0"><g id="evaUnlockOutline1"><g id="evaUnlockOutline2" fill="currentColor"><path d="M17 8h-7V6a2 2 0 0 1 4 0a1 1 0 0 0 2 0a4 4 0 0 0-8 0v2H7a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Zm1 11a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/><path d="M12 12a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g></g>`
	uploadFillInnerSVG                    = `<g id="evaUploadFill0"><g id="evaUploadFill1"><g id="evaUploadFill2" fill="currentColor"><rect width="16" height="2" x="4" y="4" rx="1" ry="1" transform="rotate(180 12 5)"/><rect width="4" height="2" x="17" y="5" rx="1" ry="1" transform="rotate(90 19 6)"/><rect width="4" height="2" x="3" y="5" rx="1" ry="1" transform="rotate(90 5 6)"/><path d="M8 14a1 1 0 0 1-.8-.4a1 1 0 0 1 .2-1.4l4-3a1 1 0 0 1 1.18 0l4 2.82a1 1 0 0 1 .24 1.39a1 1 0 0 1-1.4.24L12 11.24L8.6 13.8a1 1 0 0 1-.6.2Z"/><path d="M12 21a1 1 0 0 1-1-1v-8a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1Z"/></g></g></g>`
	uploadOutlineInnerSVG                 = `<g id="evaUploadOutline0"><g id="evaUploadOutline1"><g id="evaUploadOutline2" fill="currentColor"><rect width="16" height="2" x="4" y="4" rx="1" ry="1" transform="rotate(180 12 5)"/><rect width="4" height="2" x="17" y="5" rx="1" ry="1" transform="rotate(90 19 6)"/><rect width="4" height="2" x="3" y="5" rx="1" ry="1" transform="rotate(90 5 6)"/><path d="M8 14a1 1 0 0 1-.8-.4a1 1 0 0 1 .2-1.4l4-3a1 1 0 0 1 1.18 0l4 2.82a1 1 0 0 1 .24 1.39a1 1 0 0 1-1.4.24L12 11.24L8.6 13.8a1 1 0 0 1-.6.2Z"/><path d="M12 21a1 1 0 0 1-1-1v-8a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1Z"/></g></g></g>`
	videoFillInnerSVG                     = `<g id="evaVideoFill0"><g id="evaVideoFill1"><path id="evaVideoFill2" fill="currentColor" d="M21 7.15a1.7 1.7 0 0 0-1.85.3l-2.15 2V8a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3v-1.45l2.16 2a1.74 1.74 0 0 0 1.16.45a1.68 1.68 0 0 0 .69-.15a1.6 1.6 0 0 0 1-1.48V8.63A1.6 1.6 0 0 0 21 7.15Z"/></g></g>`
	videoOffFillInnerSVG                  = `<g id="evaVideoOffFill0"><g id="evaVideoOffFill1"><path id="evaVideoOffFill2" fill="currentColor" d="M14.22 17.05L4.88 7.71L3.12 6L3 5.8A3 3 0 0 0 2 8v8a3 3 0 0 0 3 3h9a2.94 2.94 0 0 0 1.66-.51ZM21 7.15a1.7 1.7 0 0 0-1.85.3l-2.15 2V8a3 3 0 0 0-3-3H7.83l1.29 1.29l6.59 6.59l2 2l2 2a1.73 1.73 0 0 0 .6.11a1.68 1.68 0 0 0 .69-.15a1.6 1.6 0 0 0 1-1.48V8.63a1.6 1.6 0 0 0-1-1.48Zm-4 8.44l-2-2L8.41 7l-2-2l-1.7-1.71a1 1 0 0 0-1.42 1.42l.54.53L5.59 7l9.34 9.34l1.46 1.46l2.9 2.91a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`
	videoOffOutlineInnerSVG               = `<g id="evaVideoOffOutline0"><g id="evaVideoOffOutline1"><path id="evaVideoOffOutline2" fill="currentColor" d="m17 15.59l-2-2L8.41 7l-2-2l-1.7-1.71a1 1 0 0 0-1.42 1.42l.54.53L5.59 7l9.34 9.34l1.46 1.46l2.9 2.91a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM14 17H5a1 1 0 0 1-1-1V8a1 1 0 0 1 .4-.78L3 5.8A3 3 0 0 0 2 8v8a3 3 0 0 0 3 3h9a2.94 2.94 0 0 0 1.66-.51L14.14 17a.7.7 0 0 1-.14 0Zm7-9.85a1.7 1.7 0 0 0-1.85.3l-2.15 2V8a3 3 0 0 0-3-3H7.83l2 2H14a1 1 0 0 1 1 1v4.17l4.72 4.72a1.73 1.73 0 0 0 .6.11a1.68 1.68 0 0 0 .69-.15a1.6 1.6 0 0 0 1-1.48V8.63A1.6 1.6 0 0 0 21 7.15Zm-1 7.45L17.19 12L20 9.4Z"/></g></g>`
	videoOutlineInnerSVG                  = `<g id="evaVideoOutline0"><g id="evaVideoOutline1"><path id="evaVideoOutline2" fill="currentColor" d="M21 7.15a1.7 1.7 0 0 0-1.85.3l-2.15 2V8a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3v-1.45l2.16 2a1.74 1.74 0 0 0 1.16.45a1.68 1.68 0 0 0 .69-.15a1.6 1.6 0 0 0 1-1.48V8.63A1.6 1.6 0 0 0 21 7.15ZM15 16a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1Zm5-1.4L17.19 12L20 9.4Z"/></g></g>`
	volumeDownFillInnerSVG                = `<g id="evaVolumeDownFill0"><g id="evaVolumeDownFill1"><path id="evaVolumeDownFill2" fill="currentColor" d="M20.78 8.37a1 1 0 1 0-1.56 1.26a4 4 0 0 1 0 4.74A1 1 0 0 0 20 16a1 1 0 0 0 .78-.37a6 6 0 0 0 0-7.26Zm-4.31-5.25a1 1 0 0 0-1 0L9 7.57H4a1 1 0 0 0-1 1v6.86a1 1 0 0 0 1 1h5l6.41 4.4A1.06 1.06 0 0 0 16 21a1 1 0 0 0 1-1V4a1 1 0 0 0-.53-.88Z"/></g></g>`
	volumeDownOutlineInnerSVG             = `<g id="evaVolumeDownOutline0"><g id="evaVolumeDownOutline1"><path id="evaVolumeDownOutline2" fill="currentColor" d="M20.78 8.37a1 1 0 1 0-1.56 1.26a4 4 0 0 1 0 4.74A1 1 0 0 0 20 16a1 1 0 0 0 .78-.37a6 6 0 0 0 0-7.26Zm-4.31-5.25a1 1 0 0 0-1 0L9 7.57H4a1 1 0 0 0-1 1v6.86a1 1 0 0 0 1 1h5l6.41 4.4A1.06 1.06 0 0 0 16 21a1 1 0 0 0 1-1V4a1 1 0 0 0-.53-.88ZM15 18.1l-5.1-3.5a1 1 0 0 0-.57-.17H5V9.57h4.33a1 1 0 0 0 .57-.17L15 5.9Z"/></g></g>`
	volumeMuteFillInnerSVG                = `<g id="evaVolumeMuteFill0"><g id="evaVolumeMuteFill1"><path id="evaVolumeMuteFill2" fill="currentColor" d="M17 21a1.06 1.06 0 0 1-.57-.17L10 16.43H5a1 1 0 0 1-1-1V8.57a1 1 0 0 1 1-1h5l6.41-4.4A1 1 0 0 1 18 4v16a1 1 0 0 1-1 1Z"/></g></g>`
	volumeMuteOutlineInnerSVG             = `<g id="evaVolumeMuteOutline0"><g id="evaVolumeMuteOutline1"><path id="evaVolumeMuteOutline2" fill="currentColor" d="M17 21a1.06 1.06 0 0 1-.57-.17L10 16.43H5a1 1 0 0 1-1-1V8.57a1 1 0 0 1 1-1h5l6.41-4.4A1 1 0 0 1 18 4v16a1 1 0 0 1-1 1ZM6 14.43h4.33a1 1 0 0 1 .57.17l5.1 3.5V5.9l-5.1 3.5a1 1 0 0 1-.57.17H6Z"/></g></g>`
	volumeOffFillInnerSVG                 = `<g id="evaVolumeOffFill0"><g id="evaVolumeOffFill1"><g id="evaVolumeOffFill2" fill="currentColor"><path d="m16.91 14.08l1.44 1.44a6 6 0 0 0-.07-7.15a1 1 0 1 0-1.56 1.26a4 4 0 0 1 .19 4.45Z"/><path d="M21 12a6.51 6.51 0 0 1-1.78 4.39l1.42 1.42A8.53 8.53 0 0 0 23 12a8.75 8.75 0 0 0-3.36-6.77a1 1 0 1 0-1.28 1.54A6.8 6.8 0 0 1 21 12Zm-6 .17V4a1 1 0 0 0-1.57-.83L9 6.2ZM4.74 7.57H2a1 1 0 0 0-1 1v6.86a1 1 0 0 0 1 1h5l6.41 4.4A1.06 1.06 0 0 0 14 21a1 1 0 0 0 1-1v-2.17Zm-.03-4.28a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g></g>`
	volumeOffOutlineInnerSVG              = `<g id="evaVolumeOffOutline0"><g id="evaVolumeOffOutline1"><g id="evaVolumeOffOutline2" fill="currentColor"><path d="M4.71 3.29a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm12.2 10.79l1.44 1.44a6 6 0 0 0-.07-7.15a1 1 0 1 0-1.56 1.26a4 4 0 0 1 .19 4.45Z"/><path d="M21 12a6.51 6.51 0 0 1-1.78 4.39l1.42 1.42A8.53 8.53 0 0 0 23 12a8.75 8.75 0 0 0-3.36-6.77a1 1 0 1 0-1.28 1.54A6.8 6.8 0 0 1 21 12Zm-7.5 6.1l-5.1-3.5a1 1 0 0 0-.57-.17H3.5V9.57h3.24l-2-2H2.5a1 1 0 0 0-1 1v6.86a1 1 0 0 0 1 1h5l6.41 4.4a1.06 1.06 0 0 0 .57.17a1 1 0 0 0 1-1v-1.67l-2-2Zm0-12.2v4.77l2 2V4a1 1 0 0 0-1.57-.83L9.23 6.4l1.44 1.44Z"/></g></g></g>`
	volumeUpFillInnerSVG                  = `<g id="evaVolumeUpFill0"><g id="evaVolumeUpFill1"><g id="evaVolumeUpFill2" fill="currentColor"><path d="M18.28 8.37a1 1 0 1 0-1.56 1.26a4 4 0 0 1 0 4.74A1 1 0 0 0 17.5 16a1 1 0 0 0 .78-.37a6 6 0 0 0 0-7.26Z"/><path d="M19.64 5.23a1 1 0 1 0-1.28 1.54A6.8 6.8 0 0 1 21 12a6.8 6.8 0 0 1-2.64 5.23a1 1 0 0 0-.13 1.41A1 1 0 0 0 19 19a1 1 0 0 0 .64-.23A8.75 8.75 0 0 0 23 12a8.75 8.75 0 0 0-3.36-6.77Zm-5.17-2.11a1 1 0 0 0-1 0L7 7.57H2a1 1 0 0 0-1 1v6.86a1 1 0 0 0 1 1h5l6.41 4.4A1.06 1.06 0 0 0 14 21a1 1 0 0 0 1-1V4a1 1 0 0 0-.53-.88Z"/></g></g></g>`
	volumeUpOutlineInnerSVG               = `<g id="evaVolumeUpOutline0"><g id="evaVolumeUpOutline1"><g id="evaVolumeUpOutline2" fill="currentColor"><path d="M18.28 8.37a1 1 0 1 0-1.56 1.26a4 4 0 0 1 0 4.74A1 1 0 0 0 17.5 16a1 1 0 0 0 .78-.37a6 6 0 0 0 0-7.26Z"/><path d="M19.64 5.23a1 1 0 1 0-1.28 1.54A6.8 6.8 0 0 1 21 12a6.8 6.8 0 0 1-2.64 5.23a1 1 0 0 0-.13 1.41A1 1 0 0 0 19 19a1 1 0 0 0 .64-.23A8.75 8.75 0 0 0 23 12a8.75 8.75 0 0 0-3.36-6.77ZM15 3.12a1 1 0 0 0-1 0L7.52 7.57h-5a1 1 0 0 0-1 1v6.86a1 1 0 0 0 1 1h5l6.41 4.4a1.06 1.06 0 0 0 .57.17a1 1 0 0 0 1-1V4a1 1 0 0 0-.5-.88Zm-1.47 15L8.4 14.6a1 1 0 0 0-.57-.17H3.5V9.57h4.33a1 1 0 0 0 .57-.17l5.1-3.5Z"/></g></g></g>`
	wifiFillInnerSVG                      = `<g id="evaWifiFill0"><g id="evaWifiFill1"><g id="evaWifiFill2" fill="currentColor"><circle cx="12" cy="19" r="1"/><path d="M12 14a5 5 0 0 0-3.47 1.4a1 1 0 1 0 1.39 1.44a3.08 3.08 0 0 1 4.16 0a1 1 0 1 0 1.39-1.44A5 5 0 0 0 12 14Zm0-5a9 9 0 0 0-6.47 2.75A1 1 0 0 0 7 13.14a7 7 0 0 1 10.08 0a1 1 0 0 0 .71.3a1 1 0 0 0 .72-1.69A9 9 0 0 0 12 9Z"/><path d="M21.72 7.93a14 14 0 0 0-19.44 0a1 1 0 0 0 1.38 1.44a12 12 0 0 1 16.68 0a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.03-1.41Z"/></g></g></g>`
	wifiOffFillInnerSVG                   = `<g id="evaWifiOffFill0"><g id="evaWifiOffFill1"><g id="evaWifiOffFill2" fill="currentColor"><circle cx="12" cy="19" r="1"/><path d="m12.44 11l-1.9-1.89l-2.46-2.44l-1.55-1.55l-1.82-1.83a1 1 0 0 0-1.42 1.42l1.38 1.37l1.46 1.46l2.23 2.24l1.55 1.54l2.74 2.74l2.79 2.8l3.85 3.85a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm9.28-3.07A13.93 13.93 0 0 0 12 4a14.1 14.1 0 0 0-4.44.73l1.62 1.62a11.89 11.89 0 0 1 11.16 3a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.03-1.39ZM3.82 6.65a14.32 14.32 0 0 0-1.54 1.28a1 1 0 0 0 1.38 1.44a13.09 13.09 0 0 1 1.6-1.29ZM17 13.14a1 1 0 0 0 .71.3a1 1 0 0 0 .72-1.69A9 9 0 0 0 12 9h-.16l2.35 2.35A7 7 0 0 1 17 13.14Zm-9.57-2.88a8.8 8.8 0 0 0-1.9 1.49A1 1 0 0 0 7 13.14a7.3 7.3 0 0 1 2-1.41Zm1.1 5.14a1 1 0 1 0 1.39 1.44a3.06 3.06 0 0 1 3.84-.25l-2.52-2.52a5 5 0 0 0-2.71 1.33Z"/></g></g></g>`
	wifiOffOutlineInnerSVG                = `<g id="evaWifiOffOutline0"><g id="evaWifiOffOutline1"><g id="evaWifiOffOutline2" fill="currentColor"><circle cx="12" cy="19" r="1"/><path d="m12.44 11l-1.9-1.89l-2.46-2.44l-1.55-1.55l-1.82-1.83a1 1 0 0 0-1.42 1.42l1.38 1.37l1.46 1.46l2.23 2.24l1.55 1.54l2.74 2.74l2.79 2.8l3.85 3.85a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm9.28-3.07A13.93 13.93 0 0 0 12 4a14.1 14.1 0 0 0-4.44.73l1.62 1.62a11.89 11.89 0 0 1 11.16 3a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.03-1.39ZM3.82 6.65a14.32 14.32 0 0 0-1.54 1.28a1 1 0 0 0 1.38 1.44a13.09 13.09 0 0 1 1.6-1.29ZM17 13.14a1 1 0 0 0 .71.3a1 1 0 0 0 .72-1.69A9 9 0 0 0 12 9h-.16l2.35 2.35A7 7 0 0 1 17 13.14Zm-9.57-2.88a8.8 8.8 0 0 0-1.9 1.49A1 1 0 0 0 7 13.14a7.3 7.3 0 0 1 2-1.41Zm1.1 5.14a1 1 0 1 0 1.39 1.44a3.06 3.06 0 0 1 3.84-.25l-2.52-2.52a5 5 0 0 0-2.71 1.33Z"/></g></g></g>`
	wifiOutlineInnerSVG                   = `<g id="evaWifiOutline0"><g id="evaWifiOutline1"><g id="evaWifiOutline2" fill="currentColor"><circle cx="12" cy="19" r="1"/><path d="M12 14a5 5 0 0 0-3.47 1.4a1 1 0 1 0 1.39 1.44a3.08 3.08 0 0 1 4.16 0a1 1 0 1 0 1.39-1.44A5 5 0 0 0 12 14Zm0-5a9 9 0 0 0-6.47 2.75A1 1 0 0 0 7 13.14a7 7 0 0 1 10.08 0a1 1 0 0 0 .71.3a1 1 0 0 0 .72-1.69A9 9 0 0 0 12 9Z"/><path d="M21.72 7.93a14 14 0 0 0-19.44 0a1 1 0 0 0 1.38 1.44a12 12 0 0 1 16.68 0a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.03-1.41Z"/></g></g></g>`
)

var sharedIconAttrs = engine.Attrs{"width": "1em", "height": "1em"}

func ActivityFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		activityFillInnerSVG,
		children,
	)
}

func ActivityOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		activityOutlineInnerSVG,
		children,
	)
}

func AlertCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertCircleFillInnerSVG,
		children,
	)
}

func AlertCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertCircleOutlineInnerSVG,
		children,
	)
}

func AlertTriangleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertTriangleFillInnerSVG,
		children,
	)
}

func AlertTriangleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		alertTriangleOutlineInnerSVG,
		children,
	)
}

func ArchiveFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		archiveFillInnerSVG,
		children,
	)
}

func ArchiveOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		archiveOutlineInnerSVG,
		children,
	)
}

func ArrowBackFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowBackFillInnerSVG,
		children,
	)
}

func ArrowBackOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowBackOutlineInnerSVG,
		children,
	)
}

func ArrowCircleDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleDownFillInnerSVG,
		children,
	)
}

func ArrowCircleDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleDownOutlineInnerSVG,
		children,
	)
}

func ArrowCircleLeftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleLeftFillInnerSVG,
		children,
	)
}

func ArrowCircleLeftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleLeftOutlineInnerSVG,
		children,
	)
}

func ArrowCircleRightFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleRightFillInnerSVG,
		children,
	)
}

func ArrowCircleRightOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleRightOutlineInnerSVG,
		children,
	)
}

func ArrowCircleUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleUpFillInnerSVG,
		children,
	)
}

func ArrowCircleUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowCircleUpOutlineInnerSVG,
		children,
	)
}

func ArrowDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownFillInnerSVG,
		children,
	)
}

func ArrowDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownOutlineInnerSVG,
		children,
	)
}

func ArrowDownwardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownwardFillInnerSVG,
		children,
	)
}

func ArrowDownwardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowDownwardOutlineInnerSVG,
		children,
	)
}

func ArrowForwardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowForwardFillInnerSVG,
		children,
	)
}

func ArrowForwardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowForwardOutlineInnerSVG,
		children,
	)
}

func ArrowIosBackFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowIosBackFillInnerSVG,
		children,
	)
}

func ArrowIosBackOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowIosBackOutlineInnerSVG,
		children,
	)
}

func ArrowIosDownwardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowIosDownwardFillInnerSVG,
		children,
	)
}

func ArrowIosDownwardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowIosDownwardOutlineInnerSVG,
		children,
	)
}

func ArrowIosForwardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowIosForwardFillInnerSVG,
		children,
	)
}

func ArrowIosForwardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowIosForwardOutlineInnerSVG,
		children,
	)
}

func ArrowIosUpwardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowIosUpwardFillInnerSVG,
		children,
	)
}

func ArrowIosUpwardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowIosUpwardOutlineInnerSVG,
		children,
	)
}

func ArrowLeftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLeftFillInnerSVG,
		children,
	)
}

func ArrowLeftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowLeftOutlineInnerSVG,
		children,
	)
}

func ArrowRightFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowRightFillInnerSVG,
		children,
	)
}

func ArrowRightOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowRightOutlineInnerSVG,
		children,
	)
}

func ArrowUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpFillInnerSVG,
		children,
	)
}

func ArrowUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpOutlineInnerSVG,
		children,
	)
}

func ArrowUpwardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpwardFillInnerSVG,
		children,
	)
}

func ArrowUpwardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowUpwardOutlineInnerSVG,
		children,
	)
}

func ArrowheadDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowheadDownFillInnerSVG,
		children,
	)
}

func ArrowheadDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowheadDownOutlineInnerSVG,
		children,
	)
}

func ArrowheadLeftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowheadLeftFillInnerSVG,
		children,
	)
}

func ArrowheadLeftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowheadLeftOutlineInnerSVG,
		children,
	)
}

func ArrowheadRightFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowheadRightFillInnerSVG,
		children,
	)
}

func ArrowheadRightOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowheadRightOutlineInnerSVG,
		children,
	)
}

func ArrowheadUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowheadUpFillInnerSVG,
		children,
	)
}

func ArrowheadUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		arrowheadUpOutlineInnerSVG,
		children,
	)
}

func AtFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		atFillInnerSVG,
		children,
	)
}

func AtOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		atOutlineInnerSVG,
		children,
	)
}

func AttachFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		attachFillInnerSVG,
		children,
	)
}

func AttachOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		attachOutlineInnerSVG,
		children,
	)
}

func AttachTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		attachTwoFillInnerSVG,
		children,
	)
}

func AttachTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		attachTwoOutlineInnerSVG,
		children,
	)
}

func AwardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		awardFillInnerSVG,
		children,
	)
}

func AwardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		awardOutlineInnerSVG,
		children,
	)
}

func BackspaceFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backspaceFillInnerSVG,
		children,
	)
}

func BackspaceOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		backspaceOutlineInnerSVG,
		children,
	)
}

func BarChartFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		barChartFillInnerSVG,
		children,
	)
}

func BarChartOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		barChartOutlineInnerSVG,
		children,
	)
}

func BarChartTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		barChartTwoFillInnerSVG,
		children,
	)
}

func BarChartTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		barChartTwoOutlineInnerSVG,
		children,
	)
}

func BatteryFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryFillInnerSVG,
		children,
	)
}

func BatteryOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		batteryOutlineInnerSVG,
		children,
	)
}

func BehanceFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		behanceFillInnerSVG,
		children,
	)
}

func BehanceOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		behanceOutlineInnerSVG,
		children,
	)
}

func BellFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bellFillInnerSVG,
		children,
	)
}

func BellOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bellOffFillInnerSVG,
		children,
	)
}

func BellOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bellOffOutlineInnerSVG,
		children,
	)
}

func BellOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bellOutlineInnerSVG,
		children,
	)
}

func BluetoothFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bluetoothFillInnerSVG,
		children,
	)
}

func BluetoothOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bluetoothOutlineInnerSVG,
		children,
	)
}

func BookFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookFillInnerSVG,
		children,
	)
}

func BookOpenFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookOpenFillInnerSVG,
		children,
	)
}

func BookOpenOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookOpenOutlineInnerSVG,
		children,
	)
}

func BookOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookOutlineInnerSVG,
		children,
	)
}

func BookmarkFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkFillInnerSVG,
		children,
	)
}

func BookmarkOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bookmarkOutlineInnerSVG,
		children,
	)
}

func BriefcaseFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseFillInnerSVG,
		children,
	)
}

func BriefcaseOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		briefcaseOutlineInnerSVG,
		children,
	)
}

func BrowserFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		browserFillInnerSVG,
		children,
	)
}

func BrowserOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		browserOutlineInnerSVG,
		children,
	)
}

func BrushFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		brushFillInnerSVG,
		children,
	)
}

func BrushOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		brushOutlineInnerSVG,
		children,
	)
}

func BulbFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bulbFillInnerSVG,
		children,
	)
}

func BulbOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		bulbOutlineInnerSVG,
		children,
	)
}

func CalendarFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarFillInnerSVG,
		children,
	)
}

func CalendarOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		calendarOutlineInnerSVG,
		children,
	)
}

func CameraFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraFillInnerSVG,
		children,
	)
}

func CameraOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cameraOutlineInnerSVG,
		children,
	)
}

func CarFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		carFillInnerSVG,
		children,
	)
}

func CarOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		carOutlineInnerSVG,
		children,
	)
}

func CastFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		castFillInnerSVG,
		children,
	)
}

func CastOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		castOutlineInnerSVG,
		children,
	)
}

func ChargingFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chargingFillInnerSVG,
		children,
	)
}

func ChargingOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chargingOutlineInnerSVG,
		children,
	)
}

func CheckmarkCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkCircleFillInnerSVG,
		children,
	)
}

func CheckmarkCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkCircleOutlineInnerSVG,
		children,
	)
}

func CheckmarkCircleTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkCircleTwoFillInnerSVG,
		children,
	)
}

func CheckmarkCircleTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkCircleTwoOutlineInnerSVG,
		children,
	)
}

func CheckmarkFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkFillInnerSVG,
		children,
	)
}

func CheckmarkOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkOutlineInnerSVG,
		children,
	)
}

func CheckmarkSquareFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkSquareFillInnerSVG,
		children,
	)
}

func CheckmarkSquareOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkSquareOutlineInnerSVG,
		children,
	)
}

func CheckmarkSquareTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkSquareTwoFillInnerSVG,
		children,
	)
}

func CheckmarkSquareTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		checkmarkSquareTwoOutlineInnerSVG,
		children,
	)
}

func ChevronDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDownFillInnerSVG,
		children,
	)
}

func ChevronDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronDownOutlineInnerSVG,
		children,
	)
}

func ChevronLeftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronLeftFillInnerSVG,
		children,
	)
}

func ChevronLeftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronLeftOutlineInnerSVG,
		children,
	)
}

func ChevronRightFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronRightFillInnerSVG,
		children,
	)
}

func ChevronRightOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronRightOutlineInnerSVG,
		children,
	)
}

func ChevronUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronUpFillInnerSVG,
		children,
	)
}

func ChevronUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		chevronUpOutlineInnerSVG,
		children,
	)
}

func ClipboardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardFillInnerSVG,
		children,
	)
}

func ClipboardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clipboardOutlineInnerSVG,
		children,
	)
}

func ClockFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clockFillInnerSVG,
		children,
	)
}

func ClockOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		clockOutlineInnerSVG,
		children,
	)
}

func CloseCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeCircleFillInnerSVG,
		children,
	)
}

func CloseCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeCircleOutlineInnerSVG,
		children,
	)
}

func CloseFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeFillInnerSVG,
		children,
	)
}

func CloseOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeOutlineInnerSVG,
		children,
	)
}

func CloseSquareFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeSquareFillInnerSVG,
		children,
	)
}

func CloseSquareOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		closeSquareOutlineInnerSVG,
		children,
	)
}

func CloudDownloadFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownloadFillInnerSVG,
		children,
	)
}

func CloudDownloadOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudDownloadOutlineInnerSVG,
		children,
	)
}

func CloudUploadFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUploadFillInnerSVG,
		children,
	)
}

func CloudUploadOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cloudUploadOutlineInnerSVG,
		children,
	)
}

func CodeDownloadFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeDownloadFillInnerSVG,
		children,
	)
}

func CodeDownloadOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeDownloadOutlineInnerSVG,
		children,
	)
}

func CodeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeFillInnerSVG,
		children,
	)
}

func CodeOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		codeOutlineInnerSVG,
		children,
	)
}

func CollapseFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		collapseFillInnerSVG,
		children,
	)
}

func CollapseOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		collapseOutlineInnerSVG,
		children,
	)
}

func ColorPaletteFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		colorPaletteFillInnerSVG,
		children,
	)
}

func ColorPaletteOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		colorPaletteOutlineInnerSVG,
		children,
	)
}

func ColorPickerFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		colorPickerFillInnerSVG,
		children,
	)
}

func ColorPickerOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		colorPickerOutlineInnerSVG,
		children,
	)
}

func CompassFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		compassFillInnerSVG,
		children,
	)
}

func CompassOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		compassOutlineInnerSVG,
		children,
	)
}

func CopyFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		copyFillInnerSVG,
		children,
	)
}

func CopyOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		copyOutlineInnerSVG,
		children,
	)
}

func CornerDownLeftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerDownLeftFillInnerSVG,
		children,
	)
}

func CornerDownLeftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerDownLeftOutlineInnerSVG,
		children,
	)
}

func CornerDownRightFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerDownRightFillInnerSVG,
		children,
	)
}

func CornerDownRightOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerDownRightOutlineInnerSVG,
		children,
	)
}

func CornerLeftDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerLeftDownFillInnerSVG,
		children,
	)
}

func CornerLeftDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerLeftDownOutlineInnerSVG,
		children,
	)
}

func CornerLeftUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerLeftUpFillInnerSVG,
		children,
	)
}

func CornerLeftUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerLeftUpOutlineInnerSVG,
		children,
	)
}

func CornerRightDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerRightDownFillInnerSVG,
		children,
	)
}

func CornerRightDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerRightDownOutlineInnerSVG,
		children,
	)
}

func CornerRightUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerRightUpFillInnerSVG,
		children,
	)
}

func CornerRightUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerRightUpOutlineInnerSVG,
		children,
	)
}

func CornerUpLeftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerUpLeftFillInnerSVG,
		children,
	)
}

func CornerUpLeftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerUpLeftOutlineInnerSVG,
		children,
	)
}

func CornerUpRightFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerUpRightFillInnerSVG,
		children,
	)
}

func CornerUpRightOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cornerUpRightOutlineInnerSVG,
		children,
	)
}

func CreditCardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardFillInnerSVG,
		children,
	)
}

func CreditCardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		creditCardOutlineInnerSVG,
		children,
	)
}

func CropFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cropFillInnerSVG,
		children,
	)
}

func CropOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cropOutlineInnerSVG,
		children,
	)
}

func CubeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cubeFillInnerSVG,
		children,
	)
}

func CubeOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		cubeOutlineInnerSVG,
		children,
	)
}

func DiagonalArrowLeftDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diagonalArrowLeftDownFillInnerSVG,
		children,
	)
}

func DiagonalArrowLeftDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diagonalArrowLeftDownOutlineInnerSVG,
		children,
	)
}

func DiagonalArrowLeftUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diagonalArrowLeftUpFillInnerSVG,
		children,
	)
}

func DiagonalArrowLeftUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diagonalArrowLeftUpOutlineInnerSVG,
		children,
	)
}

func DiagonalArrowRightDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diagonalArrowRightDownFillInnerSVG,
		children,
	)
}

func DiagonalArrowRightDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diagonalArrowRightDownOutlineInnerSVG,
		children,
	)
}

func DiagonalArrowRightUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diagonalArrowRightUpFillInnerSVG,
		children,
	)
}

func DiagonalArrowRightUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		diagonalArrowRightUpOutlineInnerSVG,
		children,
	)
}

func DoneAllFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		doneAllFillInnerSVG,
		children,
	)
}

func DoneAllOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		doneAllOutlineInnerSVG,
		children,
	)
}

func DownloadFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadFillInnerSVG,
		children,
	)
}

func DownloadOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		downloadOutlineInnerSVG,
		children,
	)
}

func DropletFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dropletFillInnerSVG,
		children,
	)
}

func DropletOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dropletOffFillInnerSVG,
		children,
	)
}

func DropletOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dropletOffOutlineInnerSVG,
		children,
	)
}

func DropletOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		dropletOutlineInnerSVG,
		children,
	)
}

func EditFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editFillInnerSVG,
		children,
	)
}

func EditOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editOutlineInnerSVG,
		children,
	)
}

func EditTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editTwoFillInnerSVG,
		children,
	)
}

func EditTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		editTwoOutlineInnerSVG,
		children,
	)
}

func EmailFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emailFillInnerSVG,
		children,
	)
}

func EmailOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		emailOutlineInnerSVG,
		children,
	)
}

func ExpandFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		expandFillInnerSVG,
		children,
	)
}

func ExpandOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		expandOutlineInnerSVG,
		children,
	)
}

func ExternalLinkFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		externalLinkFillInnerSVG,
		children,
	)
}

func ExternalLinkOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		externalLinkOutlineInnerSVG,
		children,
	)
}

func EyeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeFillInnerSVG,
		children,
	)
}

func EyeOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeOffFillInnerSVG,
		children,
	)
}

func EyeOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeOffOutlineInnerSVG,
		children,
	)
}

func EyeOffTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeOffTwoFillInnerSVG,
		children,
	)
}

func EyeOffTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeOffTwoOutlineInnerSVG,
		children,
	)
}

func EyeOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		eyeOutlineInnerSVG,
		children,
	)
}

func FacebookFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		facebookFillInnerSVG,
		children,
	)
}

func FacebookOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		facebookOutlineInnerSVG,
		children,
	)
}

func FileAddFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileAddFillInnerSVG,
		children,
	)
}

func FileAddOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileAddOutlineInnerSVG,
		children,
	)
}

func FileFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileFillInnerSVG,
		children,
	)
}

func FileOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileOutlineInnerSVG,
		children,
	)
}

func FileRemoveFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileRemoveFillInnerSVG,
		children,
	)
}

func FileRemoveOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileRemoveOutlineInnerSVG,
		children,
	)
}

func FileTextFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileTextFillInnerSVG,
		children,
	)
}

func FileTextOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		fileTextOutlineInnerSVG,
		children,
	)
}

func FilmFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filmFillInnerSVG,
		children,
	)
}

func FilmOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		filmOutlineInnerSVG,
		children,
	)
}

func FlagFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flagFillInnerSVG,
		children,
	)
}

func FlagOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flagOutlineInnerSVG,
		children,
	)
}

func FlashFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flashFillInnerSVG,
		children,
	)
}

func FlashOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flashOffFillInnerSVG,
		children,
	)
}

func FlashOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flashOffOutlineInnerSVG,
		children,
	)
}

func FlashOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flashOutlineInnerSVG,
		children,
	)
}

func FlipFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flipFillInnerSVG,
		children,
	)
}

func FlipOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flipOutlineInnerSVG,
		children,
	)
}

func FlipTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flipTwoFillInnerSVG,
		children,
	)
}

func FlipTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		flipTwoOutlineInnerSVG,
		children,
	)
}

func FolderAddFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderAddFillInnerSVG,
		children,
	)
}

func FolderAddOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderAddOutlineInnerSVG,
		children,
	)
}

func FolderFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderFillInnerSVG,
		children,
	)
}

func FolderOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderOutlineInnerSVG,
		children,
	)
}

func FolderRemoveFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderRemoveFillInnerSVG,
		children,
	)
}

func FolderRemoveOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		folderRemoveOutlineInnerSVG,
		children,
	)
}

func FunnelFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		funnelFillInnerSVG,
		children,
	)
}

func FunnelOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		funnelOutlineInnerSVG,
		children,
	)
}

func GiftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		giftFillInnerSVG,
		children,
	)
}

func GiftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		giftOutlineInnerSVG,
		children,
	)
}

func GithubFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		githubFillInnerSVG,
		children,
	)
}

func GithubOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		githubOutlineInnerSVG,
		children,
	)
}

func GlobeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeFillInnerSVG,
		children,
	)
}

func GlobeOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeOutlineInnerSVG,
		children,
	)
}

func GlobeThreeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeThreeFillInnerSVG,
		children,
	)
}

func GlobeTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeTwoFillInnerSVG,
		children,
	)
}

func GlobeTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		globeTwoOutlineInnerSVG,
		children,
	)
}

func GoogleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		googleFillInnerSVG,
		children,
	)
}

func GoogleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		googleOutlineInnerSVG,
		children,
	)
}

func GridFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gridFillInnerSVG,
		children,
	)
}

func GridOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		gridOutlineInnerSVG,
		children,
	)
}

func HardDriveFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hardDriveFillInnerSVG,
		children,
	)
}

func HardDriveOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hardDriveOutlineInnerSVG,
		children,
	)
}

func HashFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hashFillInnerSVG,
		children,
	)
}

func HashOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		hashOutlineInnerSVG,
		children,
	)
}

func HeadphonesFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		headphonesFillInnerSVG,
		children,
	)
}

func HeadphonesOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		headphonesOutlineInnerSVG,
		children,
	)
}

func HeartFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartFillInnerSVG,
		children,
	)
}

func HeartOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		heartOutlineInnerSVG,
		children,
	)
}

func HomeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeFillInnerSVG,
		children,
	)
}

func HomeOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		homeOutlineInnerSVG,
		children,
	)
}

func ImageFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageFillInnerSVG,
		children,
	)
}

func ImageOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageOutlineInnerSVG,
		children,
	)
}

func ImageTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		imageTwoFillInnerSVG,
		children,
	)
}

func InboxFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxFillInnerSVG,
		children,
	)
}

func InboxOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		inboxOutlineInnerSVG,
		children,
	)
}

func InfoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		infoFillInnerSVG,
		children,
	)
}

func InfoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		infoOutlineInnerSVG,
		children,
	)
}

func KeypadFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		keypadFillInnerSVG,
		children,
	)
}

func KeypadOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		keypadOutlineInnerSVG,
		children,
	)
}

func LayersFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layersFillInnerSVG,
		children,
	)
}

func LayersOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layersOutlineInnerSVG,
		children,
	)
}

func LayoutFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutFillInnerSVG,
		children,
	)
}

func LayoutOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		layoutOutlineInnerSVG,
		children,
	)
}

func LinkFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkFillInnerSVG,
		children,
	)
}

func LinkOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkOutlineInnerSVG,
		children,
	)
}

func LinkTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkTwoFillInnerSVG,
		children,
	)
}

func LinkTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkTwoOutlineInnerSVG,
		children,
	)
}

func LinkedinFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkedinFillInnerSVG,
		children,
	)
}

func LinkedinOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		linkedinOutlineInnerSVG,
		children,
	)
}

func ListFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listFillInnerSVG,
		children,
	)
}

func ListOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		listOutlineInnerSVG,
		children,
	)
}

func LoaderOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		loaderOutlineInnerSVG,
		children,
	)
}

func LockFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockFillInnerSVG,
		children,
	)
}

func LockOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		lockOutlineInnerSVG,
		children,
	)
}

func LogInFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logInFillInnerSVG,
		children,
	)
}

func LogInOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logInOutlineInnerSVG,
		children,
	)
}

func LogOutFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logOutFillInnerSVG,
		children,
	)
}

func LogOutOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		logOutOutlineInnerSVG,
		children,
	)
}

func MapFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapFillInnerSVG,
		children,
	)
}

func MapOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		mapOutlineInnerSVG,
		children,
	)
}

func MaximizeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		maximizeFillInnerSVG,
		children,
	)
}

func MaximizeOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		maximizeOutlineInnerSVG,
		children,
	)
}

func MenuArrowFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuArrowFillInnerSVG,
		children,
	)
}

func MenuArrowOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuArrowOutlineInnerSVG,
		children,
	)
}

func MenuFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuFillInnerSVG,
		children,
	)
}

func MenuOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuOutlineInnerSVG,
		children,
	)
}

func MenuTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuTwoFillInnerSVG,
		children,
	)
}

func MenuTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		menuTwoOutlineInnerSVG,
		children,
	)
}

func MessageCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageCircleFillInnerSVG,
		children,
	)
}

func MessageCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageCircleOutlineInnerSVG,
		children,
	)
}

func MessageSquareFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageSquareFillInnerSVG,
		children,
	)
}

func MessageSquareOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		messageSquareOutlineInnerSVG,
		children,
	)
}

func MicFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		micFillInnerSVG,
		children,
	)
}

func MicOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		micOffFillInnerSVG,
		children,
	)
}

func MicOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		micOffOutlineInnerSVG,
		children,
	)
}

func MicOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		micOutlineInnerSVG,
		children,
	)
}

func MinimizeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minimizeFillInnerSVG,
		children,
	)
}

func MinimizeOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minimizeOutlineInnerSVG,
		children,
	)
}

func MinusCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusCircleFillInnerSVG,
		children,
	)
}

func MinusCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusCircleOutlineInnerSVG,
		children,
	)
}

func MinusFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusFillInnerSVG,
		children,
	)
}

func MinusOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusOutlineInnerSVG,
		children,
	)
}

func MinusSquareFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusSquareFillInnerSVG,
		children,
	)
}

func MinusSquareOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		minusSquareOutlineInnerSVG,
		children,
	)
}

func MonitorFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		monitorFillInnerSVG,
		children,
	)
}

func MonitorOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		monitorOutlineInnerSVG,
		children,
	)
}

func MoonFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonFillInnerSVG,
		children,
	)
}

func MoonOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moonOutlineInnerSVG,
		children,
	)
}

func MoreHorizontalFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreHorizontalFillInnerSVG,
		children,
	)
}

func MoreHorizontalOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreHorizontalOutlineInnerSVG,
		children,
	)
}

func MoreVerticalFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreVerticalFillInnerSVG,
		children,
	)
}

func MoreVerticalOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moreVerticalOutlineInnerSVG,
		children,
	)
}

func MoveFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moveFillInnerSVG,
		children,
	)
}

func MoveOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		moveOutlineInnerSVG,
		children,
	)
}

func MusicFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		musicFillInnerSVG,
		children,
	)
}

func MusicOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		musicOutlineInnerSVG,
		children,
	)
}

func NavigationFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		navigationFillInnerSVG,
		children,
	)
}

func NavigationOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		navigationOutlineInnerSVG,
		children,
	)
}

func NavigationTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		navigationTwoFillInnerSVG,
		children,
	)
}

func NavigationTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		navigationTwoOutlineInnerSVG,
		children,
	)
}

func NpmFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		npmFillInnerSVG,
		children,
	)
}

func NpmOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		npmOutlineInnerSVG,
		children,
	)
}

func OptionsFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		optionsFillInnerSVG,
		children,
	)
}

func OptionsOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		optionsOutlineInnerSVG,
		children,
	)
}

func OptionsTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		optionsTwoFillInnerSVG,
		children,
	)
}

func OptionsTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		optionsTwoOutlineInnerSVG,
		children,
	)
}

func PantoneFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pantoneFillInnerSVG,
		children,
	)
}

func PantoneOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pantoneOutlineInnerSVG,
		children,
	)
}

func PaperPlaneFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperPlaneFillInnerSVG,
		children,
	)
}

func PaperPlaneOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		paperPlaneOutlineInnerSVG,
		children,
	)
}

func PauseCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pauseCircleFillInnerSVG,
		children,
	)
}

func PauseCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pauseCircleOutlineInnerSVG,
		children,
	)
}

func PeopleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		peopleFillInnerSVG,
		children,
	)
}

func PeopleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		peopleOutlineInnerSVG,
		children,
	)
}

func PercentFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		percentFillInnerSVG,
		children,
	)
}

func PercentOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		percentOutlineInnerSVG,
		children,
	)
}

func PersonAddFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personAddFillInnerSVG,
		children,
	)
}

func PersonAddOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personAddOutlineInnerSVG,
		children,
	)
}

func PersonDeleteFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personDeleteFillInnerSVG,
		children,
	)
}

func PersonDeleteOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personDeleteOutlineInnerSVG,
		children,
	)
}

func PersonDoneFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personDoneFillInnerSVG,
		children,
	)
}

func PersonDoneOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personDoneOutlineInnerSVG,
		children,
	)
}

func PersonFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personFillInnerSVG,
		children,
	)
}

func PersonOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personOutlineInnerSVG,
		children,
	)
}

func PersonRemoveFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personRemoveFillInnerSVG,
		children,
	)
}

func PersonRemoveOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		personRemoveOutlineInnerSVG,
		children,
	)
}

func PhoneCallFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneCallFillInnerSVG,
		children,
	)
}

func PhoneCallOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneCallOutlineInnerSVG,
		children,
	)
}

func PhoneFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneFillInnerSVG,
		children,
	)
}

func PhoneMissedFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneMissedFillInnerSVG,
		children,
	)
}

func PhoneMissedOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneMissedOutlineInnerSVG,
		children,
	)
}

func PhoneOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneOffFillInnerSVG,
		children,
	)
}

func PhoneOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneOffOutlineInnerSVG,
		children,
	)
}

func PhoneOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		phoneOutlineInnerSVG,
		children,
	)
}

func PieChartFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pieChartFillInnerSVG,
		children,
	)
}

func PieChartOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pieChartOutlineInnerSVG,
		children,
	)
}

func PieChartTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pieChartTwoFillInnerSVG,
		children,
	)
}

func PinFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pinFillInnerSVG,
		children,
	)
}

func PinOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pinOutlineInnerSVG,
		children,
	)
}

func PlayCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playCircleFillInnerSVG,
		children,
	)
}

func PlayCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		playCircleOutlineInnerSVG,
		children,
	)
}

func PlusCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusCircleFillInnerSVG,
		children,
	)
}

func PlusCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusCircleOutlineInnerSVG,
		children,
	)
}

func PlusFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusFillInnerSVG,
		children,
	)
}

func PlusOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusOutlineInnerSVG,
		children,
	)
}

func PlusSquareFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusSquareFillInnerSVG,
		children,
	)
}

func PlusSquareOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		plusSquareOutlineInnerSVG,
		children,
	)
}

func PowerFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		powerFillInnerSVG,
		children,
	)
}

func PowerOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		powerOutlineInnerSVG,
		children,
	)
}

func PricetagsFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pricetagsFillInnerSVG,
		children,
	)
}

func PricetagsOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		pricetagsOutlineInnerSVG,
		children,
	)
}

func PrinterFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		printerFillInnerSVG,
		children,
	)
}

func PrinterOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		printerOutlineInnerSVG,
		children,
	)
}

func QuestionMarkCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionMarkCircleFillInnerSVG,
		children,
	)
}

func QuestionMarkCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionMarkCircleOutlineInnerSVG,
		children,
	)
}

func QuestionMarkFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionMarkFillInnerSVG,
		children,
	)
}

func QuestionMarkOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		questionMarkOutlineInnerSVG,
		children,
	)
}

func RadioButtonOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioButtonOffFillInnerSVG,
		children,
	)
}

func RadioButtonOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioButtonOffOutlineInnerSVG,
		children,
	)
}

func RadioButtonOnFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioButtonOnFillInnerSVG,
		children,
	)
}

func RadioButtonOnOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioButtonOnOutlineInnerSVG,
		children,
	)
}

func RadioFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioFillInnerSVG,
		children,
	)
}

func RadioOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		radioOutlineInnerSVG,
		children,
	)
}

func RecordingFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		recordingFillInnerSVG,
		children,
	)
}

func RecordingOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		recordingOutlineInnerSVG,
		children,
	)
}

func RefreshFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		refreshFillInnerSVG,
		children,
	)
}

func RefreshOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		refreshOutlineInnerSVG,
		children,
	)
}

func RepeatFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		repeatFillInnerSVG,
		children,
	)
}

func RepeatOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		repeatOutlineInnerSVG,
		children,
	)
}

func RewindLeftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rewindLeftFillInnerSVG,
		children,
	)
}

func RewindLeftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rewindLeftOutlineInnerSVG,
		children,
	)
}

func RewindRightFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rewindRightFillInnerSVG,
		children,
	)
}

func RewindRightOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		rewindRightOutlineInnerSVG,
		children,
	)
}

func SaveFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveFillInnerSVG,
		children,
	)
}

func SaveOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		saveOutlineInnerSVG,
		children,
	)
}

func ScissorsFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scissorsFillInnerSVG,
		children,
	)
}

func ScissorsOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		scissorsOutlineInnerSVG,
		children,
	)
}

func SearchFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchFillInnerSVG,
		children,
	)
}

func SearchOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		searchOutlineInnerSVG,
		children,
	)
}

func SettingsFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsFillInnerSVG,
		children,
	)
}

func SettingsOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsOutlineInnerSVG,
		children,
	)
}

func SettingsTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsTwoFillInnerSVG,
		children,
	)
}

func SettingsTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		settingsTwoOutlineInnerSVG,
		children,
	)
}

func ShakeFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shakeFillInnerSVG,
		children,
	)
}

func ShakeOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shakeOutlineInnerSVG,
		children,
	)
}

func ShareFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shareFillInnerSVG,
		children,
	)
}

func ShareOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shareOutlineInnerSVG,
		children,
	)
}

func ShieldFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldFillInnerSVG,
		children,
	)
}

func ShieldOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldOffFillInnerSVG,
		children,
	)
}

func ShieldOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldOffOutlineInnerSVG,
		children,
	)
}

func ShieldOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shieldOutlineInnerSVG,
		children,
	)
}

func ShoppingBagFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingBagFillInnerSVG,
		children,
	)
}

func ShoppingBagOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingBagOutlineInnerSVG,
		children,
	)
}

func ShoppingCartFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingCartFillInnerSVG,
		children,
	)
}

func ShoppingCartOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shoppingCartOutlineInnerSVG,
		children,
	)
}

func ShuffleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shuffleFillInnerSVG,
		children,
	)
}

func ShuffleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shuffleOutlineInnerSVG,
		children,
	)
}

func ShuffleTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shuffleTwoFillInnerSVG,
		children,
	)
}

func ShuffleTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		shuffleTwoOutlineInnerSVG,
		children,
	)
}

func SkipBackFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		skipBackFillInnerSVG,
		children,
	)
}

func SkipBackOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		skipBackOutlineInnerSVG,
		children,
	)
}

func SkipForwardFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		skipForwardFillInnerSVG,
		children,
	)
}

func SkipForwardOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		skipForwardOutlineInnerSVG,
		children,
	)
}

func SlashFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		slashFillInnerSVG,
		children,
	)
}

func SlashOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		slashOutlineInnerSVG,
		children,
	)
}

func SmartphoneFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		smartphoneFillInnerSVG,
		children,
	)
}

func SmartphoneOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		smartphoneOutlineInnerSVG,
		children,
	)
}

func SmilingFaceFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		smilingFaceFillInnerSVG,
		children,
	)
}

func SmilingFaceOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		smilingFaceOutlineInnerSVG,
		children,
	)
}

func SpeakerFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speakerFillInnerSVG,
		children,
	)
}

func SpeakerOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		speakerOutlineInnerSVG,
		children,
	)
}

func SquareFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		squareFillInnerSVG,
		children,
	)
}

func SquareOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		squareOutlineInnerSVG,
		children,
	)
}

func StarFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		starFillInnerSVG,
		children,
	)
}

func StarOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		starOutlineInnerSVG,
		children,
	)
}

func StopCircleFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		stopCircleFillInnerSVG,
		children,
	)
}

func StopCircleOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		stopCircleOutlineInnerSVG,
		children,
	)
}

func SunFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunFillInnerSVG,
		children,
	)
}

func SunOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		sunOutlineInnerSVG,
		children,
	)
}

func SwapFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		swapFillInnerSVG,
		children,
	)
}

func SwapOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		swapOutlineInnerSVG,
		children,
	)
}

func SyncFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		syncFillInnerSVG,
		children,
	)
}

func SyncOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		syncOutlineInnerSVG,
		children,
	)
}

func TextFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textFillInnerSVG,
		children,
	)
}

func TextOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		textOutlineInnerSVG,
		children,
	)
}

func ThermometerFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thermometerFillInnerSVG,
		children,
	)
}

func ThermometerMinusFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thermometerMinusFillInnerSVG,
		children,
	)
}

func ThermometerMinusOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thermometerMinusOutlineInnerSVG,
		children,
	)
}

func ThermometerOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thermometerOutlineInnerSVG,
		children,
	)
}

func ThermometerPlusFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thermometerPlusFillInnerSVG,
		children,
	)
}

func ThermometerPlusOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		thermometerPlusOutlineInnerSVG,
		children,
	)
}

func ToggleLeftFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		toggleLeftFillInnerSVG,
		children,
	)
}

func ToggleLeftOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		toggleLeftOutlineInnerSVG,
		children,
	)
}

func ToggleRightFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		toggleRightFillInnerSVG,
		children,
	)
}

func ToggleRightOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		toggleRightOutlineInnerSVG,
		children,
	)
}

func TrashFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trashFillInnerSVG,
		children,
	)
}

func TrashOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trashOutlineInnerSVG,
		children,
	)
}

func TrashTwoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trashTwoFillInnerSVG,
		children,
	)
}

func TrashTwoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trashTwoOutlineInnerSVG,
		children,
	)
}

func TrendingDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingDownFillInnerSVG,
		children,
	)
}

func TrendingDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingDownOutlineInnerSVG,
		children,
	)
}

func TrendingUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingUpFillInnerSVG,
		children,
	)
}

func TrendingUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		trendingUpOutlineInnerSVG,
		children,
	)
}

func TvFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tvFillInnerSVG,
		children,
	)
}

func TvOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		tvOutlineInnerSVG,
		children,
	)
}

func TwitterFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		twitterFillInnerSVG,
		children,
	)
}

func TwitterOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		twitterOutlineInnerSVG,
		children,
	)
}

func UmbrellaFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		umbrellaFillInnerSVG,
		children,
	)
}

func UmbrellaOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		umbrellaOutlineInnerSVG,
		children,
	)
}

func UndoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		undoFillInnerSVG,
		children,
	)
}

func UndoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		undoOutlineInnerSVG,
		children,
	)
}

func UnlockFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		unlockFillInnerSVG,
		children,
	)
}

func UnlockOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		unlockOutlineInnerSVG,
		children,
	)
}

func UploadFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadFillInnerSVG,
		children,
	)
}

func UploadOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		uploadOutlineInnerSVG,
		children,
	)
}

func VideoFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoFillInnerSVG,
		children,
	)
}

func VideoOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoOffFillInnerSVG,
		children,
	)
}

func VideoOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoOffOutlineInnerSVG,
		children,
	)
}

func VideoOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		videoOutlineInnerSVG,
		children,
	)
}

func VolumeDownFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeDownFillInnerSVG,
		children,
	)
}

func VolumeDownOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeDownOutlineInnerSVG,
		children,
	)
}

func VolumeMuteFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeMuteFillInnerSVG,
		children,
	)
}

func VolumeMuteOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeMuteOutlineInnerSVG,
		children,
	)
}

func VolumeOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeOffFillInnerSVG,
		children,
	)
}

func VolumeOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeOffOutlineInnerSVG,
		children,
	)
}

func VolumeUpFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeUpFillInnerSVG,
		children,
	)
}

func VolumeUpOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		volumeUpOutlineInnerSVG,
		children,
	)
}

func WifiFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		wifiFillInnerSVG,
		children,
	)
}

func WifiOffFill(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		wifiOffFillInnerSVG,
		children,
	)
}

func WifiOffOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		wifiOffOutlineInnerSVG,
		children,
	)
}

func WifiOutline(children ...any) *engine.HTMLElement {
	return engine.Element(
		"svg",
		sharedIconAttrs,
		engine.Attrs{
			"viewBox": "0 0 24 24",
		},
		wifiOutlineInnerSVG,
		children,
	)
}

func ByName(name string) (*engine.HTMLElement, error) {
	switch name {
	case "activity-fill":
		return ActivityFill(), nil
	case "activity-outline":
		return ActivityOutline(), nil
	case "alert-circle-fill":
		return AlertCircleFill(), nil
	case "alert-circle-outline":
		return AlertCircleOutline(), nil
	case "alert-triangle-fill":
		return AlertTriangleFill(), nil
	case "alert-triangle-outline":
		return AlertTriangleOutline(), nil
	case "archive-fill":
		return ArchiveFill(), nil
	case "archive-outline":
		return ArchiveOutline(), nil
	case "arrow-back-fill":
		return ArrowBackFill(), nil
	case "arrow-back-outline":
		return ArrowBackOutline(), nil
	case "arrow-circle-down-fill":
		return ArrowCircleDownFill(), nil
	case "arrow-circle-down-outline":
		return ArrowCircleDownOutline(), nil
	case "arrow-circle-left-fill":
		return ArrowCircleLeftFill(), nil
	case "arrow-circle-left-outline":
		return ArrowCircleLeftOutline(), nil
	case "arrow-circle-right-fill":
		return ArrowCircleRightFill(), nil
	case "arrow-circle-right-outline":
		return ArrowCircleRightOutline(), nil
	case "arrow-circle-up-fill":
		return ArrowCircleUpFill(), nil
	case "arrow-circle-up-outline":
		return ArrowCircleUpOutline(), nil
	case "arrow-down-fill":
		return ArrowDownFill(), nil
	case "arrow-down-outline":
		return ArrowDownOutline(), nil
	case "arrow-downward-fill":
		return ArrowDownwardFill(), nil
	case "arrow-downward-outline":
		return ArrowDownwardOutline(), nil
	case "arrow-forward-fill":
		return ArrowForwardFill(), nil
	case "arrow-forward-outline":
		return ArrowForwardOutline(), nil
	case "arrow-ios-back-fill":
		return ArrowIosBackFill(), nil
	case "arrow-ios-back-outline":
		return ArrowIosBackOutline(), nil
	case "arrow-ios-downward-fill":
		return ArrowIosDownwardFill(), nil
	case "arrow-ios-downward-outline":
		return ArrowIosDownwardOutline(), nil
	case "arrow-ios-forward-fill":
		return ArrowIosForwardFill(), nil
	case "arrow-ios-forward-outline":
		return ArrowIosForwardOutline(), nil
	case "arrow-ios-upward-fill":
		return ArrowIosUpwardFill(), nil
	case "arrow-ios-upward-outline":
		return ArrowIosUpwardOutline(), nil
	case "arrow-left-fill":
		return ArrowLeftFill(), nil
	case "arrow-left-outline":
		return ArrowLeftOutline(), nil
	case "arrow-right-fill":
		return ArrowRightFill(), nil
	case "arrow-right-outline":
		return ArrowRightOutline(), nil
	case "arrow-up-fill":
		return ArrowUpFill(), nil
	case "arrow-up-outline":
		return ArrowUpOutline(), nil
	case "arrow-upward-fill":
		return ArrowUpwardFill(), nil
	case "arrow-upward-outline":
		return ArrowUpwardOutline(), nil
	case "arrowhead-down-fill":
		return ArrowheadDownFill(), nil
	case "arrowhead-down-outline":
		return ArrowheadDownOutline(), nil
	case "arrowhead-left-fill":
		return ArrowheadLeftFill(), nil
	case "arrowhead-left-outline":
		return ArrowheadLeftOutline(), nil
	case "arrowhead-right-fill":
		return ArrowheadRightFill(), nil
	case "arrowhead-right-outline":
		return ArrowheadRightOutline(), nil
	case "arrowhead-up-fill":
		return ArrowheadUpFill(), nil
	case "arrowhead-up-outline":
		return ArrowheadUpOutline(), nil
	case "at-fill":
		return AtFill(), nil
	case "at-outline":
		return AtOutline(), nil
	case "attach-fill":
		return AttachFill(), nil
	case "attach-outline":
		return AttachOutline(), nil
	case "attach-2-fill":
		return AttachTwoFill(), nil
	case "attach-2-outline":
		return AttachTwoOutline(), nil
	case "award-fill":
		return AwardFill(), nil
	case "award-outline":
		return AwardOutline(), nil
	case "backspace-fill":
		return BackspaceFill(), nil
	case "backspace-outline":
		return BackspaceOutline(), nil
	case "bar-chart-fill":
		return BarChartFill(), nil
	case "bar-chart-outline":
		return BarChartOutline(), nil
	case "bar-chart-2-fill":
		return BarChartTwoFill(), nil
	case "bar-chart-2-outline":
		return BarChartTwoOutline(), nil
	case "battery-fill":
		return BatteryFill(), nil
	case "battery-outline":
		return BatteryOutline(), nil
	case "behance-fill":
		return BehanceFill(), nil
	case "behance-outline":
		return BehanceOutline(), nil
	case "bell-fill":
		return BellFill(), nil
	case "bell-off-fill":
		return BellOffFill(), nil
	case "bell-off-outline":
		return BellOffOutline(), nil
	case "bell-outline":
		return BellOutline(), nil
	case "bluetooth-fill":
		return BluetoothFill(), nil
	case "bluetooth-outline":
		return BluetoothOutline(), nil
	case "book-fill":
		return BookFill(), nil
	case "book-open-fill":
		return BookOpenFill(), nil
	case "book-open-outline":
		return BookOpenOutline(), nil
	case "book-outline":
		return BookOutline(), nil
	case "bookmark-fill":
		return BookmarkFill(), nil
	case "bookmark-outline":
		return BookmarkOutline(), nil
	case "briefcase-fill":
		return BriefcaseFill(), nil
	case "briefcase-outline":
		return BriefcaseOutline(), nil
	case "browser-fill":
		return BrowserFill(), nil
	case "browser-outline":
		return BrowserOutline(), nil
	case "brush-fill":
		return BrushFill(), nil
	case "brush-outline":
		return BrushOutline(), nil
	case "bulb-fill":
		return BulbFill(), nil
	case "bulb-outline":
		return BulbOutline(), nil
	case "calendar-fill":
		return CalendarFill(), nil
	case "calendar-outline":
		return CalendarOutline(), nil
	case "camera-fill":
		return CameraFill(), nil
	case "camera-outline":
		return CameraOutline(), nil
	case "car-fill":
		return CarFill(), nil
	case "car-outline":
		return CarOutline(), nil
	case "cast-fill":
		return CastFill(), nil
	case "cast-outline":
		return CastOutline(), nil
	case "charging-fill":
		return ChargingFill(), nil
	case "charging-outline":
		return ChargingOutline(), nil
	case "checkmark-circle-fill":
		return CheckmarkCircleFill(), nil
	case "checkmark-circle-outline":
		return CheckmarkCircleOutline(), nil
	case "checkmark-circle-2-fill":
		return CheckmarkCircleTwoFill(), nil
	case "checkmark-circle-2-outline":
		return CheckmarkCircleTwoOutline(), nil
	case "checkmark-fill":
		return CheckmarkFill(), nil
	case "checkmark-outline":
		return CheckmarkOutline(), nil
	case "checkmark-square-fill":
		return CheckmarkSquareFill(), nil
	case "checkmark-square-outline":
		return CheckmarkSquareOutline(), nil
	case "checkmark-square-2-fill":
		return CheckmarkSquareTwoFill(), nil
	case "checkmark-square-2-outline":
		return CheckmarkSquareTwoOutline(), nil
	case "chevron-down-fill":
		return ChevronDownFill(), nil
	case "chevron-down-outline":
		return ChevronDownOutline(), nil
	case "chevron-left-fill":
		return ChevronLeftFill(), nil
	case "chevron-left-outline":
		return ChevronLeftOutline(), nil
	case "chevron-right-fill":
		return ChevronRightFill(), nil
	case "chevron-right-outline":
		return ChevronRightOutline(), nil
	case "chevron-up-fill":
		return ChevronUpFill(), nil
	case "chevron-up-outline":
		return ChevronUpOutline(), nil
	case "clipboard-fill":
		return ClipboardFill(), nil
	case "clipboard-outline":
		return ClipboardOutline(), nil
	case "clock-fill":
		return ClockFill(), nil
	case "clock-outline":
		return ClockOutline(), nil
	case "close-circle-fill":
		return CloseCircleFill(), nil
	case "close-circle-outline":
		return CloseCircleOutline(), nil
	case "close-fill":
		return CloseFill(), nil
	case "close-outline":
		return CloseOutline(), nil
	case "close-square-fill":
		return CloseSquareFill(), nil
	case "close-square-outline":
		return CloseSquareOutline(), nil
	case "cloud-download-fill":
		return CloudDownloadFill(), nil
	case "cloud-download-outline":
		return CloudDownloadOutline(), nil
	case "cloud-upload-fill":
		return CloudUploadFill(), nil
	case "cloud-upload-outline":
		return CloudUploadOutline(), nil
	case "code-download-fill":
		return CodeDownloadFill(), nil
	case "code-download-outline":
		return CodeDownloadOutline(), nil
	case "code-fill":
		return CodeFill(), nil
	case "code-outline":
		return CodeOutline(), nil
	case "collapse-fill":
		return CollapseFill(), nil
	case "collapse-outline":
		return CollapseOutline(), nil
	case "color-palette-fill":
		return ColorPaletteFill(), nil
	case "color-palette-outline":
		return ColorPaletteOutline(), nil
	case "color-picker-fill":
		return ColorPickerFill(), nil
	case "color-picker-outline":
		return ColorPickerOutline(), nil
	case "compass-fill":
		return CompassFill(), nil
	case "compass-outline":
		return CompassOutline(), nil
	case "copy-fill":
		return CopyFill(), nil
	case "copy-outline":
		return CopyOutline(), nil
	case "corner-down-left-fill":
		return CornerDownLeftFill(), nil
	case "corner-down-left-outline":
		return CornerDownLeftOutline(), nil
	case "corner-down-right-fill":
		return CornerDownRightFill(), nil
	case "corner-down-right-outline":
		return CornerDownRightOutline(), nil
	case "corner-left-down-fill":
		return CornerLeftDownFill(), nil
	case "corner-left-down-outline":
		return CornerLeftDownOutline(), nil
	case "corner-left-up-fill":
		return CornerLeftUpFill(), nil
	case "corner-left-up-outline":
		return CornerLeftUpOutline(), nil
	case "corner-right-down-fill":
		return CornerRightDownFill(), nil
	case "corner-right-down-outline":
		return CornerRightDownOutline(), nil
	case "corner-right-up-fill":
		return CornerRightUpFill(), nil
	case "corner-right-up-outline":
		return CornerRightUpOutline(), nil
	case "corner-up-left-fill":
		return CornerUpLeftFill(), nil
	case "corner-up-left-outline":
		return CornerUpLeftOutline(), nil
	case "corner-up-right-fill":
		return CornerUpRightFill(), nil
	case "corner-up-right-outline":
		return CornerUpRightOutline(), nil
	case "credit-card-fill":
		return CreditCardFill(), nil
	case "credit-card-outline":
		return CreditCardOutline(), nil
	case "crop-fill":
		return CropFill(), nil
	case "crop-outline":
		return CropOutline(), nil
	case "cube-fill":
		return CubeFill(), nil
	case "cube-outline":
		return CubeOutline(), nil
	case "diagonal-arrow-left-down-fill":
		return DiagonalArrowLeftDownFill(), nil
	case "diagonal-arrow-left-down-outline":
		return DiagonalArrowLeftDownOutline(), nil
	case "diagonal-arrow-left-up-fill":
		return DiagonalArrowLeftUpFill(), nil
	case "diagonal-arrow-left-up-outline":
		return DiagonalArrowLeftUpOutline(), nil
	case "diagonal-arrow-right-down-fill":
		return DiagonalArrowRightDownFill(), nil
	case "diagonal-arrow-right-down-outline":
		return DiagonalArrowRightDownOutline(), nil
	case "diagonal-arrow-right-up-fill":
		return DiagonalArrowRightUpFill(), nil
	case "diagonal-arrow-right-up-outline":
		return DiagonalArrowRightUpOutline(), nil
	case "done-all-fill":
		return DoneAllFill(), nil
	case "done-all-outline":
		return DoneAllOutline(), nil
	case "download-fill":
		return DownloadFill(), nil
	case "download-outline":
		return DownloadOutline(), nil
	case "droplet-fill":
		return DropletFill(), nil
	case "droplet-off-fill":
		return DropletOffFill(), nil
	case "droplet-off-outline":
		return DropletOffOutline(), nil
	case "droplet-outline":
		return DropletOutline(), nil
	case "edit-fill":
		return EditFill(), nil
	case "edit-outline":
		return EditOutline(), nil
	case "edit-2-fill":
		return EditTwoFill(), nil
	case "edit-2-outline":
		return EditTwoOutline(), nil
	case "email-fill":
		return EmailFill(), nil
	case "email-outline":
		return EmailOutline(), nil
	case "expand-fill":
		return ExpandFill(), nil
	case "expand-outline":
		return ExpandOutline(), nil
	case "external-link-fill":
		return ExternalLinkFill(), nil
	case "external-link-outline":
		return ExternalLinkOutline(), nil
	case "eye-fill":
		return EyeFill(), nil
	case "eye-off-fill":
		return EyeOffFill(), nil
	case "eye-off-outline":
		return EyeOffOutline(), nil
	case "eye-off-2-fill":
		return EyeOffTwoFill(), nil
	case "eye-off-2-outline":
		return EyeOffTwoOutline(), nil
	case "eye-outline":
		return EyeOutline(), nil
	case "facebook-fill":
		return FacebookFill(), nil
	case "facebook-outline":
		return FacebookOutline(), nil
	case "file-add-fill":
		return FileAddFill(), nil
	case "file-add-outline":
		return FileAddOutline(), nil
	case "file-fill":
		return FileFill(), nil
	case "file-outline":
		return FileOutline(), nil
	case "file-remove-fill":
		return FileRemoveFill(), nil
	case "file-remove-outline":
		return FileRemoveOutline(), nil
	case "file-text-fill":
		return FileTextFill(), nil
	case "file-text-outline":
		return FileTextOutline(), nil
	case "film-fill":
		return FilmFill(), nil
	case "film-outline":
		return FilmOutline(), nil
	case "flag-fill":
		return FlagFill(), nil
	case "flag-outline":
		return FlagOutline(), nil
	case "flash-fill":
		return FlashFill(), nil
	case "flash-off-fill":
		return FlashOffFill(), nil
	case "flash-off-outline":
		return FlashOffOutline(), nil
	case "flash-outline":
		return FlashOutline(), nil
	case "flip-fill":
		return FlipFill(), nil
	case "flip-outline":
		return FlipOutline(), nil
	case "flip-2-fill":
		return FlipTwoFill(), nil
	case "flip-2-outline":
		return FlipTwoOutline(), nil
	case "folder-add-fill":
		return FolderAddFill(), nil
	case "folder-add-outline":
		return FolderAddOutline(), nil
	case "folder-fill":
		return FolderFill(), nil
	case "folder-outline":
		return FolderOutline(), nil
	case "folder-remove-fill":
		return FolderRemoveFill(), nil
	case "folder-remove-outline":
		return FolderRemoveOutline(), nil
	case "funnel-fill":
		return FunnelFill(), nil
	case "funnel-outline":
		return FunnelOutline(), nil
	case "gift-fill":
		return GiftFill(), nil
	case "gift-outline":
		return GiftOutline(), nil
	case "github-fill":
		return GithubFill(), nil
	case "github-outline":
		return GithubOutline(), nil
	case "globe-fill":
		return GlobeFill(), nil
	case "globe-outline":
		return GlobeOutline(), nil
	case "globe-3-fill":
		return GlobeThreeFill(), nil
	case "globe-2-fill":
		return GlobeTwoFill(), nil
	case "globe-2-outline":
		return GlobeTwoOutline(), nil
	case "google-fill":
		return GoogleFill(), nil
	case "google-outline":
		return GoogleOutline(), nil
	case "grid-fill":
		return GridFill(), nil
	case "grid-outline":
		return GridOutline(), nil
	case "hard-drive-fill":
		return HardDriveFill(), nil
	case "hard-drive-outline":
		return HardDriveOutline(), nil
	case "hash-fill":
		return HashFill(), nil
	case "hash-outline":
		return HashOutline(), nil
	case "headphones-fill":
		return HeadphonesFill(), nil
	case "headphones-outline":
		return HeadphonesOutline(), nil
	case "heart-fill":
		return HeartFill(), nil
	case "heart-outline":
		return HeartOutline(), nil
	case "home-fill":
		return HomeFill(), nil
	case "home-outline":
		return HomeOutline(), nil
	case "image-fill":
		return ImageFill(), nil
	case "image-outline":
		return ImageOutline(), nil
	case "image-2-fill":
		return ImageTwoFill(), nil
	case "inbox-fill":
		return InboxFill(), nil
	case "inbox-outline":
		return InboxOutline(), nil
	case "info-fill":
		return InfoFill(), nil
	case "info-outline":
		return InfoOutline(), nil
	case "keypad-fill":
		return KeypadFill(), nil
	case "keypad-outline":
		return KeypadOutline(), nil
	case "layers-fill":
		return LayersFill(), nil
	case "layers-outline":
		return LayersOutline(), nil
	case "layout-fill":
		return LayoutFill(), nil
	case "layout-outline":
		return LayoutOutline(), nil
	case "link-fill":
		return LinkFill(), nil
	case "link-outline":
		return LinkOutline(), nil
	case "link-2-fill":
		return LinkTwoFill(), nil
	case "link-2-outline":
		return LinkTwoOutline(), nil
	case "linkedin-fill":
		return LinkedinFill(), nil
	case "linkedin-outline":
		return LinkedinOutline(), nil
	case "list-fill":
		return ListFill(), nil
	case "list-outline":
		return ListOutline(), nil
	case "loader-outline":
		return LoaderOutline(), nil
	case "lock-fill":
		return LockFill(), nil
	case "lock-outline":
		return LockOutline(), nil
	case "log-in-fill":
		return LogInFill(), nil
	case "log-in-outline":
		return LogInOutline(), nil
	case "log-out-fill":
		return LogOutFill(), nil
	case "log-out-outline":
		return LogOutOutline(), nil
	case "map-fill":
		return MapFill(), nil
	case "map-outline":
		return MapOutline(), nil
	case "maximize-fill":
		return MaximizeFill(), nil
	case "maximize-outline":
		return MaximizeOutline(), nil
	case "menu-arrow-fill":
		return MenuArrowFill(), nil
	case "menu-arrow-outline":
		return MenuArrowOutline(), nil
	case "menu-fill":
		return MenuFill(), nil
	case "menu-outline":
		return MenuOutline(), nil
	case "menu-2-fill":
		return MenuTwoFill(), nil
	case "menu-2-outline":
		return MenuTwoOutline(), nil
	case "message-circle-fill":
		return MessageCircleFill(), nil
	case "message-circle-outline":
		return MessageCircleOutline(), nil
	case "message-square-fill":
		return MessageSquareFill(), nil
	case "message-square-outline":
		return MessageSquareOutline(), nil
	case "mic-fill":
		return MicFill(), nil
	case "mic-off-fill":
		return MicOffFill(), nil
	case "mic-off-outline":
		return MicOffOutline(), nil
	case "mic-outline":
		return MicOutline(), nil
	case "minimize-fill":
		return MinimizeFill(), nil
	case "minimize-outline":
		return MinimizeOutline(), nil
	case "minus-circle-fill":
		return MinusCircleFill(), nil
	case "minus-circle-outline":
		return MinusCircleOutline(), nil
	case "minus-fill":
		return MinusFill(), nil
	case "minus-outline":
		return MinusOutline(), nil
	case "minus-square-fill":
		return MinusSquareFill(), nil
	case "minus-square-outline":
		return MinusSquareOutline(), nil
	case "monitor-fill":
		return MonitorFill(), nil
	case "monitor-outline":
		return MonitorOutline(), nil
	case "moon-fill":
		return MoonFill(), nil
	case "moon-outline":
		return MoonOutline(), nil
	case "more-horizontal-fill":
		return MoreHorizontalFill(), nil
	case "more-horizontal-outline":
		return MoreHorizontalOutline(), nil
	case "more-vertical-fill":
		return MoreVerticalFill(), nil
	case "more-vertical-outline":
		return MoreVerticalOutline(), nil
	case "move-fill":
		return MoveFill(), nil
	case "move-outline":
		return MoveOutline(), nil
	case "music-fill":
		return MusicFill(), nil
	case "music-outline":
		return MusicOutline(), nil
	case "navigation-fill":
		return NavigationFill(), nil
	case "navigation-outline":
		return NavigationOutline(), nil
	case "navigation-2-fill":
		return NavigationTwoFill(), nil
	case "navigation-2-outline":
		return NavigationTwoOutline(), nil
	case "npm-fill":
		return NpmFill(), nil
	case "npm-outline":
		return NpmOutline(), nil
	case "options-fill":
		return OptionsFill(), nil
	case "options-outline":
		return OptionsOutline(), nil
	case "options-2-fill":
		return OptionsTwoFill(), nil
	case "options-2-outline":
		return OptionsTwoOutline(), nil
	case "pantone-fill":
		return PantoneFill(), nil
	case "pantone-outline":
		return PantoneOutline(), nil
	case "paper-plane-fill":
		return PaperPlaneFill(), nil
	case "paper-plane-outline":
		return PaperPlaneOutline(), nil
	case "pause-circle-fill":
		return PauseCircleFill(), nil
	case "pause-circle-outline":
		return PauseCircleOutline(), nil
	case "people-fill":
		return PeopleFill(), nil
	case "people-outline":
		return PeopleOutline(), nil
	case "percent-fill":
		return PercentFill(), nil
	case "percent-outline":
		return PercentOutline(), nil
	case "person-add-fill":
		return PersonAddFill(), nil
	case "person-add-outline":
		return PersonAddOutline(), nil
	case "person-delete-fill":
		return PersonDeleteFill(), nil
	case "person-delete-outline":
		return PersonDeleteOutline(), nil
	case "person-done-fill":
		return PersonDoneFill(), nil
	case "person-done-outline":
		return PersonDoneOutline(), nil
	case "person-fill":
		return PersonFill(), nil
	case "person-outline":
		return PersonOutline(), nil
	case "person-remove-fill":
		return PersonRemoveFill(), nil
	case "person-remove-outline":
		return PersonRemoveOutline(), nil
	case "phone-call-fill":
		return PhoneCallFill(), nil
	case "phone-call-outline":
		return PhoneCallOutline(), nil
	case "phone-fill":
		return PhoneFill(), nil
	case "phone-missed-fill":
		return PhoneMissedFill(), nil
	case "phone-missed-outline":
		return PhoneMissedOutline(), nil
	case "phone-off-fill":
		return PhoneOffFill(), nil
	case "phone-off-outline":
		return PhoneOffOutline(), nil
	case "phone-outline":
		return PhoneOutline(), nil
	case "pie-chart-fill":
		return PieChartFill(), nil
	case "pie-chart-outline":
		return PieChartOutline(), nil
	case "pie-chart-2-fill":
		return PieChartTwoFill(), nil
	case "pin-fill":
		return PinFill(), nil
	case "pin-outline":
		return PinOutline(), nil
	case "play-circle-fill":
		return PlayCircleFill(), nil
	case "play-circle-outline":
		return PlayCircleOutline(), nil
	case "plus-circle-fill":
		return PlusCircleFill(), nil
	case "plus-circle-outline":
		return PlusCircleOutline(), nil
	case "plus-fill":
		return PlusFill(), nil
	case "plus-outline":
		return PlusOutline(), nil
	case "plus-square-fill":
		return PlusSquareFill(), nil
	case "plus-square-outline":
		return PlusSquareOutline(), nil
	case "power-fill":
		return PowerFill(), nil
	case "power-outline":
		return PowerOutline(), nil
	case "pricetags-fill":
		return PricetagsFill(), nil
	case "pricetags-outline":
		return PricetagsOutline(), nil
	case "printer-fill":
		return PrinterFill(), nil
	case "printer-outline":
		return PrinterOutline(), nil
	case "question-mark-circle-fill":
		return QuestionMarkCircleFill(), nil
	case "question-mark-circle-outline":
		return QuestionMarkCircleOutline(), nil
	case "question-mark-fill":
		return QuestionMarkFill(), nil
	case "question-mark-outline":
		return QuestionMarkOutline(), nil
	case "radio-button-off-fill":
		return RadioButtonOffFill(), nil
	case "radio-button-off-outline":
		return RadioButtonOffOutline(), nil
	case "radio-button-on-fill":
		return RadioButtonOnFill(), nil
	case "radio-button-on-outline":
		return RadioButtonOnOutline(), nil
	case "radio-fill":
		return RadioFill(), nil
	case "radio-outline":
		return RadioOutline(), nil
	case "recording-fill":
		return RecordingFill(), nil
	case "recording-outline":
		return RecordingOutline(), nil
	case "refresh-fill":
		return RefreshFill(), nil
	case "refresh-outline":
		return RefreshOutline(), nil
	case "repeat-fill":
		return RepeatFill(), nil
	case "repeat-outline":
		return RepeatOutline(), nil
	case "rewind-left-fill":
		return RewindLeftFill(), nil
	case "rewind-left-outline":
		return RewindLeftOutline(), nil
	case "rewind-right-fill":
		return RewindRightFill(), nil
	case "rewind-right-outline":
		return RewindRightOutline(), nil
	case "save-fill":
		return SaveFill(), nil
	case "save-outline":
		return SaveOutline(), nil
	case "scissors-fill":
		return ScissorsFill(), nil
	case "scissors-outline":
		return ScissorsOutline(), nil
	case "search-fill":
		return SearchFill(), nil
	case "search-outline":
		return SearchOutline(), nil
	case "settings-fill":
		return SettingsFill(), nil
	case "settings-outline":
		return SettingsOutline(), nil
	case "settings-2-fill":
		return SettingsTwoFill(), nil
	case "settings-2-outline":
		return SettingsTwoOutline(), nil
	case "shake-fill":
		return ShakeFill(), nil
	case "shake-outline":
		return ShakeOutline(), nil
	case "share-fill":
		return ShareFill(), nil
	case "share-outline":
		return ShareOutline(), nil
	case "shield-fill":
		return ShieldFill(), nil
	case "shield-off-fill":
		return ShieldOffFill(), nil
	case "shield-off-outline":
		return ShieldOffOutline(), nil
	case "shield-outline":
		return ShieldOutline(), nil
	case "shopping-bag-fill":
		return ShoppingBagFill(), nil
	case "shopping-bag-outline":
		return ShoppingBagOutline(), nil
	case "shopping-cart-fill":
		return ShoppingCartFill(), nil
	case "shopping-cart-outline":
		return ShoppingCartOutline(), nil
	case "shuffle-fill":
		return ShuffleFill(), nil
	case "shuffle-outline":
		return ShuffleOutline(), nil
	case "shuffle-2-fill":
		return ShuffleTwoFill(), nil
	case "shuffle-2-outline":
		return ShuffleTwoOutline(), nil
	case "skip-back-fill":
		return SkipBackFill(), nil
	case "skip-back-outline":
		return SkipBackOutline(), nil
	case "skip-forward-fill":
		return SkipForwardFill(), nil
	case "skip-forward-outline":
		return SkipForwardOutline(), nil
	case "slash-fill":
		return SlashFill(), nil
	case "slash-outline":
		return SlashOutline(), nil
	case "smartphone-fill":
		return SmartphoneFill(), nil
	case "smartphone-outline":
		return SmartphoneOutline(), nil
	case "smiling-face-fill":
		return SmilingFaceFill(), nil
	case "smiling-face-outline":
		return SmilingFaceOutline(), nil
	case "speaker-fill":
		return SpeakerFill(), nil
	case "speaker-outline":
		return SpeakerOutline(), nil
	case "square-fill":
		return SquareFill(), nil
	case "square-outline":
		return SquareOutline(), nil
	case "star-fill":
		return StarFill(), nil
	case "star-outline":
		return StarOutline(), nil
	case "stop-circle-fill":
		return StopCircleFill(), nil
	case "stop-circle-outline":
		return StopCircleOutline(), nil
	case "sun-fill":
		return SunFill(), nil
	case "sun-outline":
		return SunOutline(), nil
	case "swap-fill":
		return SwapFill(), nil
	case "swap-outline":
		return SwapOutline(), nil
	case "sync-fill":
		return SyncFill(), nil
	case "sync-outline":
		return SyncOutline(), nil
	case "text-fill":
		return TextFill(), nil
	case "text-outline":
		return TextOutline(), nil
	case "thermometer-fill":
		return ThermometerFill(), nil
	case "thermometer-minus-fill":
		return ThermometerMinusFill(), nil
	case "thermometer-minus-outline":
		return ThermometerMinusOutline(), nil
	case "thermometer-outline":
		return ThermometerOutline(), nil
	case "thermometer-plus-fill":
		return ThermometerPlusFill(), nil
	case "thermometer-plus-outline":
		return ThermometerPlusOutline(), nil
	case "toggle-left-fill":
		return ToggleLeftFill(), nil
	case "toggle-left-outline":
		return ToggleLeftOutline(), nil
	case "toggle-right-fill":
		return ToggleRightFill(), nil
	case "toggle-right-outline":
		return ToggleRightOutline(), nil
	case "trash-fill":
		return TrashFill(), nil
	case "trash-outline":
		return TrashOutline(), nil
	case "trash-2-fill":
		return TrashTwoFill(), nil
	case "trash-2-outline":
		return TrashTwoOutline(), nil
	case "trending-down-fill":
		return TrendingDownFill(), nil
	case "trending-down-outline":
		return TrendingDownOutline(), nil
	case "trending-up-fill":
		return TrendingUpFill(), nil
	case "trending-up-outline":
		return TrendingUpOutline(), nil
	case "tv-fill":
		return TvFill(), nil
	case "tv-outline":
		return TvOutline(), nil
	case "twitter-fill":
		return TwitterFill(), nil
	case "twitter-outline":
		return TwitterOutline(), nil
	case "umbrella-fill":
		return UmbrellaFill(), nil
	case "umbrella-outline":
		return UmbrellaOutline(), nil
	case "undo-fill":
		return UndoFill(), nil
	case "undo-outline":
		return UndoOutline(), nil
	case "unlock-fill":
		return UnlockFill(), nil
	case "unlock-outline":
		return UnlockOutline(), nil
	case "upload-fill":
		return UploadFill(), nil
	case "upload-outline":
		return UploadOutline(), nil
	case "video-fill":
		return VideoFill(), nil
	case "video-off-fill":
		return VideoOffFill(), nil
	case "video-off-outline":
		return VideoOffOutline(), nil
	case "video-outline":
		return VideoOutline(), nil
	case "volume-down-fill":
		return VolumeDownFill(), nil
	case "volume-down-outline":
		return VolumeDownOutline(), nil
	case "volume-mute-fill":
		return VolumeMuteFill(), nil
	case "volume-mute-outline":
		return VolumeMuteOutline(), nil
	case "volume-off-fill":
		return VolumeOffFill(), nil
	case "volume-off-outline":
		return VolumeOffOutline(), nil
	case "volume-up-fill":
		return VolumeUpFill(), nil
	case "volume-up-outline":
		return VolumeUpOutline(), nil
	case "wifi-fill":
		return WifiFill(), nil
	case "wifi-off-fill":
		return WifiOffFill(), nil
	case "wifi-off-outline":
		return WifiOffOutline(), nil
	case "wifi-outline":
		return WifiOutline(), nil
	default:
		return nil, fmt.Errorf("icon '%s' not found in eva icon set", name)
	}
}
