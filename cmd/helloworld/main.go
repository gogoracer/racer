package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gogoracer/racer/pkg/engine"

	. "github.com/gogoracer/racer/pkg/handlebars" //lint:ignore ST1001 we have a DSL
	"github.com/rs/zerolog/log"
)

// Step 1
func homeStep1() *engine.Page {
	page := engine.NewPage()
	page.DOM().Body().Add("Hello, world!")

	// Ignore this, well explain it later
	addStageButtons(page, 1)

	return page
}

// Step 2
func homeStep2() *engine.Page {
	var message string

	input := engine.NewComponent("input")
	input.Add(engine.Attrs{"type": "text"})
	input.Add(engine.On("keyup", func(ctx context.Context, e engine.Event) {
		message = e.Value
	}))

	page := engine.NewPage()
	page.DOM().Body().Add(engine.NewTag("div", input))
	page.DOM().Body().Add(engine.NewTag("hr"))
	page.DOM().Body().Add("Hello, ", &message)

	// Ignore this, well explain it later
	addStageButtons(page, 2)

	return page
}

// Step 2.1
func homeStep3() *engine.Page {
	var message string

	page := engine.NewPage()
	page.DOM().Body().Add(
		engine.NewTag(
			"div",
			engine.NewComponent("input",
				engine.Attrs{"type": "text"},
				engine.On("keyup", func(_ context.Context, e engine.Event) {
					message = e.Value
				}),
			)),
		engine.NewTag("hr"),
		"Hello!!, ", &message,
	)

	// Ignore this, well explain it later
	addStageButtons(page, 3)

	return page
}

// Use the DSL
func homeStep4() *engine.Page {
	message := "abc"

	page := engine.NewPage()
	page.DOM().Body().Add(
		Division(
			InputComponent(
				InputAttributes(
					AttrType, "text",
					AttrValue, message,
				),
				OnKeyUp(func(ctx context.Context, e engine.Event) {
					message = e.Value
				}),
			),
		),
		HorizontalRule(),
		"Hello, ", &message,

		ButtonComponent(
			OnClick(func(ctx context.Context, e engine.Event) {
				message = ""
			}),
			"Clear",
		),
	)

	// Ignore this, well explain it later
	addStageButtons(page, 4)

	return page
}

func main() {
	http.Handle("/step1", engine.MustNewPageServer(homeStep1))
	http.Handle("/step2", engine.MustNewPageServer(homeStep2))
	http.Handle("/step3", engine.MustNewPageServer(homeStep3))
	http.Handle("/step4", engine.MustNewPageServer(homeStep4))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/step1", http.StatusFound)
	}))

	port := 3000

	addr := fmt.Sprintf(":%d", port)

	log.Info().Msg("Listing on http://localhost" + addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Info().Err(err)
	}
}

func addStageButtons(page *engine.Page, stage int) {
	body := page.DOM().Body()

	body.Add(HorizontalRule())

	if stage > 1 {
		body.Add(
			ButtonComponent(
				ButtonAttributes(
					AttrOnClick, fmt.Sprintf("window.location.href = '/step%d'", stage-1),
				),
				"Previous",
			),
		)
	}

	if stage < 4 {
		body.Add(
			ButtonComponent(
				ButtonAttributes(
					AttrOnClick, fmt.Sprintf("window.location.href = '/step%d'", stage+1),
				),
				"Next",
			),
		)
	}
}
