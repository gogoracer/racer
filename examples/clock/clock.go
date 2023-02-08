package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gogoracer/racer/pkg/engine"
	"github.com/gogoracer/racer/pkg/frame"
	. "github.com/gogoracer/racer/pkg/goggles/racer_attribute"
	. "github.com/gogoracer/racer/pkg/goggles/racer_html"
)

func main() {
	http.Handle("/", home())

	log.Println("INFO: listing :3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Println("ERRO: http listen and serve: ", err)
	}
}

func home() *engine.PageServer {
	f := func() engine.Pager {

		page := frame.NewPage()
		page.DOM().Title().Val().Str("Click Example")
		page.DOM().Head().Element(
			LINK().
				Attr(REL("stylesheet")).
				Attr(HREF("https://cdn.simplecss.org/simple.min.css")))

		page.DOM().Body().
			Element(
				HEADER().
					Element(H1().Val().Str("Clock")).
					Element(P().Val().Str("The time updates are being push from the server every 10ms")))

		page.DOM().Body().
			Element(
				MAIN().Element(PRE().Component(newClock())))

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

	c := &clock{
		UberElement: CODE().Val().Str("Server: ").LockBox(t),
		timeStr:     t,
	}

	c.SetMount(c.mount)
	c.SetUnmount(c.unmount)

	return c
}

type clock struct {
	*engine.UberElement

	timeStr *engine.LockBox[string]
	tick    *time.Ticker
}

func (c *clock) mount(ctx context.Context) {
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

				engine.RenderElement(ctx, c.UberElement)
			}
		}
	}()
}

// Unmount
// Will be called after the page session is deleted
func (c *clock) unmount(_ context.Context) {
	log.Println("DEBU: stop tick")

	if c.tick != nil {
		c.tick.Stop()
	}
}
