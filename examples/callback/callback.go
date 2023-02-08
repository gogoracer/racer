package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gogoracer/racer/pkg/engine"
	"github.com/gogoracer/racer/pkg/frame"
	. "github.com/gogoracer/racer/pkg/goggles/racer_attribute"
	. "github.com/gogoracer/racer/pkg/goggles/racer_event"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"
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

func callback(container *engine.UberElement) {
	eb, attr := frame.OnDiffApply(
		func(ctx context.Context, e engine.Event) {
			container.Element(P().Val().Str("Diff Applied"))
			container.RemoveEventBinding(e.Binding.ID)
		},
	)

	container.On(eb).Attr(attr)

}

func home() engine.Pager {
	container := CODE()

	btn := BUTTON().Val().Str("Trigger Click").On(CLICK(func(ctx context.Context, e engine.Event) {
		container.Element(P().Val().Str("Click"))
		callback(container)
	}))

	page := frame.NewPage()
	page.DOM().Title().Val().Str("Callback Example")
	page.DOM().Head().Element(
		LINK().
			Attr(REL("stylesheet")).
			Attr(HREF("https://cdn.simplecss.org/simple.min.css")))

	page.DOM().Body().
		Element(
			HEADER().
				Element(H1().Val().Str("Callback")).
				Element(P().Val().Str("Get notified when a change has been applied in the browser")))

	page.DOM().Body().
		Element(
			MAIN().
				Element(P().Element(btn)).
				Element(H2().Val().Str("Events")).
				Element(PRE().Element(container)))

	return page
}
