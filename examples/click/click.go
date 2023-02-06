package main

import (
	"context"
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

func home() *engine.Page {
	page := engine.NewPage()
	page.DOM().Title().Add("Click Example")
	page.DOM().Head().Add(
		engine.NewTag("link", engine.Attrs{"rel": "stylesheet", "href": "https://cdn.simplecss.org/simple.min.css"}))

	// A thread safe value container
	count := engine.Box(0)

	page.DOM().Body().Add(
		engine.NewTag("header",
			engine.NewTag("h1", "Click"),
			engine.NewTag("p", "Click the button and see the count increase"),
		),
		engine.NewTag("main",
			engine.NewTag("p",
				"Clicks: ",
				engine.NewComponent("button",
					engine.On("click", func(_ context.Context, _ engine.Event) {
						// We need to read and write inside a single lock
						count.Lock(func(v int) int { return v + 1 })
					}),
					count,
				),
			),
		),
	)

	return page
}
