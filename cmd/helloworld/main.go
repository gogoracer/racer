package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gogoracer/racer/pkg/gas"
	. "github.com/gogoracer/racer/pkg/handlebars"
	"github.com/rs/zerolog/log"
)

// Step 1
func homeStep1() *gas.Page {
	page := gas.NewPage()
	page.DOM().Body().Add("Hello, world!")

	// Ignore this, well explain it later
	addStageButtons(page, 1)

	return page
}

// Step 2
func homeStep2() *gas.Page {
	var message string

	input := gas.NewComponent("input")
	input.Add(gas.Attrs{"type": "text"})
	input.Add(gas.On("keyup", func(ctx context.Context, e gas.Event) {
		message = e.Value
	}))

	page := gas.NewPage()
	page.DOM().Body().Add(gas.NewTag("div", input))
	page.DOM().Body().Add(gas.NewTag("hr"))
	page.DOM().Body().Add("Hello, ", &message)

	// Ignore this, well explain it later
	addStageButtons(page, 2)

	return page
}

// Step 2.1
func homeStep3() *gas.Page {
	var message string

	page := gas.NewPage()
	page.DOM().Body().Add(
		gas.NewTag(
			"div",
			gas.NewComponent("input",
				gas.Attrs{"type": "text"},
				gas.On("keyup", func(_ context.Context, e gas.Event) {
					message = e.Value
				}),
			)),
		gas.NewTag("hr"),
		"Hello!!, ", &message,
	)

	// Ignore this, well explain it later
	addStageButtons(page, 3)

	return page
}

// Use the DSL
func homeStep4() *gas.Page {
	var message string

	page := gas.NewPage()
	page.DOM().Body().Add(
		Division(
			InputComponent(
				InputAttributes(
					AttrType, "text",
				),
				OnKeyDown(func(ctx context.Context, e gas.Event) {
					message = e.Value
				}),
			),
		),
		HorizontalRule(),
		"Hello, ", &message,
	)

	// Ignore this, well explain it later
	addStageButtons(page, 4)

	return page
}

func main() {
	http.Handle("/step1", gas.NewPageServer(homeStep1))
	http.Handle("/step2", gas.NewPageServer(homeStep2))
	http.Handle("/step3", gas.NewPageServer(homeStep3))
	http.Handle("/step4", gas.NewPageServer(homeStep4))
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

func addStageButtons(page *gas.Page, stage int) {
	body := page.DOM().Body()

	body.Add(HorizontalRule())

	if stage > 1 {
		body.Add(
			ButtonComponent(
				ButtonAttributes(
					"onclick", fmt.Sprintf("window.location.href = '/step%d'", stage-1),
				),
				"Previous",
			),
		)
	}

	if stage < 4 {
		body.Add(
			ButtonComponent(
				ButtonAttributes(
					"onclick", fmt.Sprintf("window.location.href = '/step%d'", stage+1),
				),
				"Next",
			),
		)
	}
}
