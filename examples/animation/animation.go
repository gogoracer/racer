package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gogoracer/racer/pkg/frame"
	. "github.com/gogoracer/racer/pkg/goggles/racer_attribute"
	. "github.com/gogoracer/racer/pkg/goggles/racer_event"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"

	"github.com/gogoracer/racer/pkg/engine"
)

const pageStyle string = `
.box {
	overflow: hidden;
	padding: 3em;
	text-align: center;
	border: solid;
}
.text {
	display: inline-block;
	font-size: 3em;
}
`

var animations = []string{
	"animate__hinge", "animate__jackInTheBox", "animate__rollIn", "animate__rollOut",
	"animate__bounce", "animate__flash", "animate__pulse", "animate__rubberBand", "animate__shakeX",
	"animate__shakeY", "animate__headShake", "animate__swing", "animate__tada", "animate__wobble",
	"animate__jello", "animate__heartBeat", "animate__flip", "animate__backInDown", "animate__backOutDown",
}

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
	var (
		index    = engine.Box(0)
		current  = engine.Box("")
		playing  = engine.NewLockBox(false)
		btnLabel = engine.Box("Start")
	)

	animationTarget := DIV().Class("animate__animated text").Val().Str("racer")

	nextAnimation := func() {
		index.Lock(func(v int) int {
			animationTarget.ClassOff(animations[v])

			v++
			if len(animations) == v {
				v = 0
			}

			current.Set(animations[v])
			animationTarget.Class(animations[v])

			return v
		})
	}

	animationTarget.On(ANIMATIONEND(func(ctx context.Context, e engine.Event) {
		if playing.Get() {
			nextAnimation()
		}
	}))

	animationTarget.On(ANIMATIONCANCEL(func(ctx context.Context, e engine.Event) {
		playing.Set(false)
		btnLabel.Set("Start")
		current.Set("")
	}))

	btn := BUTTON().Box(btnLabel).
		On(CLICK(func(ctx context.Context, e engine.Event) {
			playing.Lock(func(v bool) bool {
				if !v {
					nextAnimation()
					btnLabel.Set("Stop")
				} else {
					btnLabel.Set("Start")
					current.Set("")
				}

				return !v
			})
		})).
		// You can create multiple event bindings for the same event and component
		On(CLICK(func(ctx context.Context, e engine.Event) {
			log.Println("INFO: Button Clicked")
		}))

	page := frame.NewPage()
	page.DOM().Title().Val().Str("Animation Example")

	page.DOM().Head().
		Element(
			LINK().
				Attr(REL("stylesheet")).
				Attr(HREF("https://cdn.simplecss.org/simple.min.css"))).
		Element(
			LINK().
				Attr(REL("stylesheet")).
				Attr(HREF("https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css"))).
		Element(
			STYLETAG().HTML(pageStyle))

	page.DOM().Body().
		Element(
			HEADER().
				Element(H1().Val().Str("CSS Animations")).
				Element(P().Val().Str("We can wait for the CSS animation to end before starting the next one")))

	page.DOM().Body().
		Element(
			MAIN().
				Element(P().Element(btn)).
				Element(P().Val().Str("Current: ").Box(current)).
				Element(DIV().Class("box").Element(animationTarget)))

	return page
}
