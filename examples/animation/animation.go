package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gogoracer/racer/pkg/engine"
)

const pageStyle engine.HTML = `
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

func home() *engine.Page {
	var (
		index    = engine.Box(0)
		current  = engine.Box("")
		playing  = engine.NewLockBox(false)
		btnLabel = engine.Box("Start")
	)

	animationTarget := engine.NewComponent("div", engine.Class("animate__animated text"), "racer")

	nextAnimation := func() {
		index.Lock(func(v int) int {
			animationTarget.Add(engine.ClassOff(animations[v]))

			v++
			if len(animations) == v {
				v = 0
			}

			current.Set(animations[v])
			animationTarget.Add(engine.Class(animations[v]))

			return v
		})
	}

	animationTarget.Add(engine.On("animationend", func(ctx context.Context, e engine.Event) {
		if playing.Get() {
			nextAnimation()
		}
	}))

	animationTarget.Add(engine.On("animationcancel", func(ctx context.Context, e engine.Event) {
		playing.Set(false)
		btnLabel.Set("Start")
		current.Set("")
	}))

	btn := engine.NewComponent("button", btnLabel,
		engine.On("click", func(ctx context.Context, e engine.Event) {
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
		}),
		// You can create multiple event bindings for the same event and component
		engine.On("click", func(ctx context.Context, e engine.Event) {
			log.Println("INFO: Button Clicked")
		}),
	)

	page := engine.NewPage()
	page.DOM().Title().Add("Animation Example")
	page.DOM().Head().Add(
		engine.NewTag("link", engine.Attrs{"rel": "stylesheet", "href": "https://cdn.simplecss.org/simple.min.css"}),
		engine.NewTag("link",
			engine.Attrs{"rel": "stylesheet", "href": "https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css"}),
		engine.NewTag("style", pageStyle),
	)

	page.DOM().Body().Add(
		engine.NewTag("header",
			engine.NewTag("h1", "CSS Animations"),
			engine.NewTag("p", "We can wait for the CSS animation to end before starting the next one"),
		),
		engine.NewTag("main",
			engine.NewTag("p", btn),
			engine.NewTag("p", "Current: ", current),
			engine.NewTag("div", engine.Class("box"), animationTarget),
		),
	)

	return page
}
