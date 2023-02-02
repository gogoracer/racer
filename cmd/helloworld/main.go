package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/gogoracer/racer/pkg/engine"
	"github.com/gogoracer/racer/pkg/handlebars/iconify/academicons"

	. "github.com/gogoracer/racer/pkg/handlebars" //lint:ignore ST1001 we have a DSL
	"github.com/rs/zerolog/log"
)

var loadURLOnStart = "/step5"

// Step 1
func homeStep1() *engine.Page {
	page := engine.NewPage()
	page.DOM().Body().Add("Hello, world!!!!!")

	//  Ignore this, well explain it later
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

	foo := INPUT().
		Type("text").
		Value("message").
		Custom("x-data", fmt.Sprintf("{message:%s}", message)).
		Custom("x-text", "message").
		OnKeyup(func(ctx context.Context, e engine.Event) {
			message = e.Value
		})
	page := NewPage().Body(
		DIV(
			foo,
		),
		HR(),
		Text("Hello, "),
		BoundText(&message),

		BUTTON(Text("Clear")).
			OnClick(func(ctx context.Context, e engine.Event) {
				message = ""
			}),
	).GenerateVDOM()

	addStageButtons(page, 4)

	return page
}

// Add styling
func homeStep5() *engine.Page {
	message := "abc"

	start := time.Now()

	tailwindCustom := SCRIPT(Text(`
	tailwind.config = {
		theme: {
		  extend: {
			colors: {
			  clifford: "#da373d",
			}
		  }
		}
	  }
	`)).Defer("true")

	page := NewPage().
		Head(
			SCRIPT().Src("https://cdn.tailwindcss.com"),
			tailwindCustom,
			SCRIPT().Defer("true").Src("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"),
		).
		Body(
			H1(Text(`HelloWold!!!`)).
				Class("text-5xl font-bold text-clifford"),
			DIV(
				INPUT().
					Type("text").
					Value("message").
					// Custom("x-data", fmt.Sprintf("{message:%s}", message)).
					// Custom("x-text", "message").
					OnKeyup(func(ctx context.Context, e engine.Event) {
						message = e.Value
					}),
			),
			DIV(
				Text("Hello, "),
				BoundText(&message),
				BUTTON(Text("Clear")).
					Class("bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded").
					OnClick(func(ctx context.Context, e engine.Event) {
						message = ""
					}),
				academicons.GoogleScholar(),
			).Class("flex gap-4"),
		).GenerateVDOM()

	addStageButtons(page, 5)

	page.DOM().Body().Add(engine.Attrs{
		"class": "bg-gray-200 min-h-screen w-full",
	})

	log.Info().Msgf("Page generated in %s", time.Since(start))

	return page
}

func main() {
	http.Handle("/step1", engine.MustNewPageServer(homeStep1))
	http.Handle("/step2", engine.MustNewPageServer(homeStep2))
	http.Handle("/step3", engine.MustNewPageServer(homeStep3))
	http.Handle("/step4", engine.MustNewPageServer(homeStep4))
	http.Handle("/step5", engine.MustNewPageServer(homeStep5))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/step1", http.StatusFound)
	}))

	ctx, cancel := context.WithCancel(context.Background())

	port := 3000
	addr := fmt.Sprintf(":%d", port)
	localAddr := fmt.Sprintf("http://localhost:%d", port)

	log.Info().Msg("Listing on " + localAddr)

	srv := &http.Server{Addr: addr}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("Server failed")
			cancel()
		}
	}()

	go func() {
		// Make sure page is ready before we start
		backoff := backoff.NewExponentialBackOff()
		for {
			if _, err := http.Get(localAddr + loadURLOnStart); err == nil {
				break
			}

			d := backoff.NextBackOff()
			log.Printf("Server not ready. Retrying in %v", d)
			time.Sleep(d)
		}

		// path, _ := launcher.LookPath()
		// url := launcher.New().Headless(false).MustLaunch()
		// browser := rod.New().ControlURL(url).MustConnect()

		wsURL := launcher.NewUserMode().MustLaunch()
		browser := rod.New().ControlURL(wsURL).MustConnect().NoDefaultDevice()

		pages, err := browser.Pages()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get pages")
		}

		toLoad := localAddr + loadURLOnStart
		var page *rod.Page
		for _, p := range pages {
			info, err := p.Info()
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to get page info")
			}

			if info.URL == toLoad {
				p.MustActivate().MustReload()
				page = p
				break
			}
		}
		if page == nil {
			page = browser.MustPage(localAddr, loadURLOnStart)
		}

		log.Print("Page loaded", page.FrameID)
	}()

	<-ctx.Done()

	log.Info().Msg("Shutting down server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to shutdown server")
	}

}

func addStageButtons(page *engine.Page, stage int) {
	const gotoFmt = "/step%d"
	buttonContainer := DIV().Class("flex gap-4")

	if stage > 1 {
		buttonContainer.Add(
			A(Text("Previous")).
				Href(fmt.Sprintf(gotoFmt, stage-1)).
				Class(btnClass),
		)
	}

	if stage < 5 {
		buttonContainer.Add(
			A(Text("Next")).
				Href(fmt.Sprintf(gotoFmt, stage+1)).
				Class(btnClass),
		)
	}

	page.DOM().Body().Add(buttonContainer.GenerateVDOM())
}

const btnClass = "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
