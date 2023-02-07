package main

import (
	"context"
	"github.com/gogoracer/racer/pkg/frame"
	. "github.com/gogoracer/racer/pkg/goggles/racer_attribute"
	. "github.com/gogoracer/racer/pkg/goggles/racer_event"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"
	"log"
	"net/http"

	"github.com/gogoracer/racer/pkg/engine"
)

func main() {
	srv, err := engine.NewPageServer(home)
	if err != nil {
		panic(err)
	}

	http.Handle("/", srv)

	log.Println("INFO: listing :3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Println("ERRO: http listen and serve: ", err)
	}
}

func home() engine.Pager {
	page := frame.NewPage()
	page.DOM().Title().Val().Str("Click Example")
	page.DOM().Head().Element(
		LINK().
			Attr(REL("stylesheet")).
			Attr(HREF("https://cdn.simplecss.org/simple.min.css")))

	// A thread safe value container
	count := engine.Box(0)

	page.DOM().Body().
		Element(
			HEADER().
				Element(H1().Val().Str("Click")).
				Element(P().Val().Str("Click the button and see the count increase")))

	page.DOM().Body().
		Element(
			MAIN().
				Element(P().Val().Str("Clicks: ").
					Element(BUTTON().
						Box(count).
						On(CLICK(func(_ context.Context, _ engine.Event) {
							// We need to read and write inside a single lock
							count.Lock(func(v int) int { return v + 1 })
						})),
					),
				),
		)

	return page
}
