package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gogoracer/racer/pkg/engine"
	"github.com/gogoracer/racer/pkg/parts"
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

func callback(container *engine.Component) {
	container.Add(
		parts.OnDiffApply(
			func(ctx context.Context, e engine.Event) {
				container.Add(engine.NewTag("p", "Diff Applied"))
				container.RemoveEventBinding(e.Binding.ID)
			},
		),
	)
}

func home() *engine.Page {
	container := engine.NewComponent("code")

	btn := engine.NewComponent("button", "Trigger Click",
		engine.On("click", func(ctx context.Context, e engine.Event) {
			container.Add(engine.NewTag("p", "Click"))
			callback(container)
		}),
	)

	page := engine.NewPage()
	page.DOM().Title().Add("Callback Example")
	page.DOM().Head().Add(
		engine.NewTag("link", engine.Attrs{"rel": "stylesheet", "href": "https://cdn.simplecss.org/simple.min.css"}))

	page.DOM().Body().Add(
		engine.NewTag("header",
			engine.NewTag("h1", "Callback"),
			engine.NewTag("p", "Get notified when a change has been applied in the browser"),
		),
		engine.NewTag("main",
			engine.NewTag("p", btn),
			engine.NewTag("h2", "Events"),
			engine.NewTag("pre", container),
		),
	)

	return page
}
