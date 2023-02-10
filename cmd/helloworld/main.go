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
	"github.com/gogoracer/racer/pkg/frame"
	"github.com/gogoracer/racer/pkg/goggles/iconify/academicons"
	. "github.com/gogoracer/racer/pkg/goggles/racer_attribute"
	. "github.com/gogoracer/racer/pkg/goggles/racer_event"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"

	"github.com/rs/zerolog/log"
)

var loadURLOnStart = "/step1"

// Step 1
func homeStep1() engine.Pager {
	page := frame.NewPage()
	page.DOM().Body().Val().Str("Hello, world!!!!!")

	//  Ignore this, well explain it later
	addStageButtons(page, 1)

	return page
}

// Step 2
func homeStep2() engine.Pager {
	var message string

	page := frame.NewPage()

	page.DOM().Body(
		DIV(
			INPUT().
				Attr(TYPE("text")).
				On(KEYUP(func(ctx context.Context, e engine.Event) {
					message = e.Value
				})),
		),
		HR(),
		DIV().
			Val().Str("Hello, ").
			Val().BindStr(&message),
	)

	// Ignore this, well explain it later
	addStageButtons(page, 2)

	return page
}

// Add styling
func homeStep3() engine.Pager {
	message := "abc"

	start := time.Now()

	page := frame.NewPage()
	dom := page.DOM()

	dom.Head(
		SCRIPT().Attr(
			SRC("https://cdn.tailwindcss.com"),
		),
		SCRIPT().Attr(
			DEFER(),
		).Val().Str(`tailwind.config = {
			theme: {
				extend: {
				colors: {
					clifford: "#da373d",
				}
				}
			}
		}`),
		SCRIPT().Attr(
			DEFER(),
			SRC("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"),
		),
	)

	dom.Body(
		H1().Class("text-5xl font-bold text-clifford").Val().Str("HelloWold!!!"),
		DIV(
			INPUT().Attr(
				TYPE("text"),
				VALUE("message"),
			).On(KEYUP(func(ctx context.Context, e engine.Event) {
				message = e.Value
			})),
			// Custom("x-data", fmt.Sprintf("{message:%s}", message)).
			// Custom("x-text", "message").
			DIV(
				SPAN().Val().Str("Hello, "),
				SPAN().Val().BindStr(&message),
				BUTTON().
					Class("bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded").
					Val().Str("Clear").
					On(CLICK(func(ctx context.Context, e engine.Event) {
						message = ""
					})),
			),
			academicons.GoogleScholar(),
		),
	)

	addStageButtons(page, 5)

	page.DOM().Body().Class("bg-gray-200 min-h-screen w-full")

	log.Info().Msgf("Page generated in %s", time.Since(start))

	return page
}

func main() {
	http.Handle("/step1", engine.MustNewPageServer(homeStep1))
	http.Handle("/step2", engine.MustNewPageServer(homeStep2))
	http.Handle("/step3", engine.MustNewPageServer(homeStep3))
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

func addStageButtons(page *frame.Page, stage int) {
	const gotoFmt = "/step%d"
	buttonContainer := DIV().Class("flex gap-4")

	if stage > 1 {
		buttonContainer.Element(
			HYPERLINK().Attr(
				HREF(fmt.Sprintf(gotoFmt, stage-1)),
				CLASS(btnClass),
			).Val().Str("Previous"),
		)
	}

	if stage < 5 {
		buttonContainer.Element(
			HYPERLINK().Val().Str("Next").Attr(
				HREF(fmt.Sprintf(gotoFmt, stage+1)),
				CLASS(btnClass),
			),
		)
	}

	page.DOM().Body().Element(buttonContainer)
}

const btnClass = "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
