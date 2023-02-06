package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gogoracer/racer/pkg/engine"
)

func main() {
	http.Handle("/", home())

	log.Println("INFO: listing :3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Println("ERRO: http listen and serve: ", err)
	}
}

func home() *engine.PageServer {
	f := func() *engine.Page {

		page := engine.NewPage()
		page.DOM().Title().Add("Clock Example")
		page.DOM().Head().Add(
			engine.NewTag("link", engine.Attrs{"rel": "stylesheet", "href": "https://cdn.simplecss.org/simple.min.css"}))

		page.DOM().Body().Add(
			engine.NewTag("header",
				engine.NewTag("h1", "Clock"),
				engine.NewTag("p", "The time updates are being push from the server every 10ms"),
			),
			engine.NewTag("main",
				engine.NewTag("pre", newClock()),
			),
		)

		return page
	}

	ps, err := engine.NewPageServer(f)
	if err != nil {
		panic(err)
	}
	// Kill the page session 1 second after the tab is closed
	ps.Sessions.DisconnectTimeout = time.Second

	return ps
}

func newClock() *clock {
	t := engine.NewLockBox("")

	return &clock{
		Component: engine.NewComponent("code", "Server: ", t),
		timeStr:   t,
	}
}

type clock struct {
	*engine.Component

	timeStr *engine.LockBox[string]
	tick    *time.Ticker
}

func (c *clock) Mount(ctx context.Context) {
	log.Println("DEBU: start tick")

	c.tick = time.NewTicker(10 * time.Millisecond)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("DEBU: tick loop stop: ctx")

				return
			case t := <-c.tick.C:
				c.timeStr.Set(t.String())

				engine.RenderComponent(ctx, c)
			}
		}
	}()
}

// Unmount
// Will be called after the page session is deleted
func (c *clock) Unmount(_ context.Context) {
	log.Println("DEBU: stop tick")

	if c.tick != nil {
		c.tick.Stop()
	}
}
