package parts_test

// import (
// 	"github.com/SamHennessy/hlive/hlivetest"
// 	"github.com/playwright-community/playwright-go"
// )

// type harness struct {
// 	server   *hlivetest.Server
// 	pwpage   playwright.Page
// 	teardown func()
// }

// func setup(t *testing.T, pageFn func() *engine.Page) harness {
// 	t.Helper()

// 	if testing.Short() {
// 		t.Skip("skipping test in short mode.")
// 	}

// 	h := harness{
// 		server: hlivetest.NewServer(pageFn),
// 		pwpage: hlivetest.NewBrowserPage(),
// 	}

// 	h.teardown = func() {
// 		if err := h.pwpage.Close(); err != nil {
// 			t.Error(err)
// 		}
// 	}

// 	if _, err := h.pwpage.Goto(h.server.HTTPServer.URL); err != nil {
// 		t.Fatal("goto page:", err)
// 	}

// 	return h
// }
