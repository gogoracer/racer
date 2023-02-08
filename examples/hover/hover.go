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

func home() engine.Pager {
	hoverState := engine.Box(" ")

	hover := H2().
		Styles(engine.Style{"padding": "1em", "text-align": "center", "border": "solid"}).
		On(MOUSEENTER(func(_ context.Context, _ engine.Event) {
			hoverState.Set("Mouse enter")
		})).
		On(MOUSELEAVE(func(_ context.Context, _ engine.Event) {
			hoverState.Set("Mouse leave")
		})).
		Val().Str("Hover over me")

	page := frame.NewPage()
	page.DOM().Title().Val().Str("Hover Example")
	page.DOM().Head().Element(
		LINK().
			Attr(REL("stylesheet")).
			Attr(HREF("https://cdn.simplecss.org/simple.min.css")))

	page.DOM().Body().
		Element(
			HEADER().
				Element(H1().Val().Str("Hover")).
				Element(P().Val().Str("React to hover events on the server")))

	page.DOM().Body().
		Element(
			MAIN().
				Element(DIV().Element(hover)).
				Element(HR()).
				Element(PRE().Element(CODE().Box(hoverState))))

	return page
}
