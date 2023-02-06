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
	hoverState := engine.Box(" ")

	hover := engine.NewComponent("h2",
		engine.Style{"padding": "1em", "text-align": "center", "border": "solid"},
		engine.On("mouseEnter", func(ctx context.Context, e engine.Event) {
			hoverState.Set("Mouse enter")
		}),
		engine.On("mouseLeave", func(ctx context.Context, e engine.Event) {
			hoverState.Set("Mouse leave")
		}),
		"Hover over me",
	)

	page := engine.NewPage()
	page.DOM().Title().Add("Hover Example")
	page.DOM().Head().Add(
		engine.NewTag("link", engine.Attrs{"rel": "stylesheet", "href": "https://cdn.simplecss.org/simple.min.css"}))

	page.DOM().Body().Add(
		engine.NewTag("header",
			engine.NewTag("h1", "Hover"),
			engine.NewTag("p", "React to hover events on the server"),
		),
		engine.NewTag("main",
			engine.NewTag("div", hover),
			engine.NewTag("hr"),
			engine.NewTag("pre", engine.NewTag("code", hoverState)),
		),
	)

	return page
}
