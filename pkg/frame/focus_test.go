package frame_test

// func TestFocus(t *testing.T) {
// 	t.Parallel()

// 	// A button, an input, on button click give input focus
// 	pageFn := func() *engine.Page {
// 		page := engine.NewPage()

// 		input := engine.NewTag("input", engine.Attrs{"id": "in_f"})

// 		page.DOM().Body().Add(
// 			input,
// 			engine.NewComponent("button", engine.Attrs{"id": "btn_f"}, engine.On("click", func(ctx context.Context, e engine.Event) {
// 				input.Add(hlivekit.Focus())
// 			})),
// 		)

// 		return page
// 	}

// 	h := setup(t, pageFn)
// 	defer h.teardown()

// 	response, err := h.pwpage.EvalOnSelector("#in_f", "(el) => el === document.activeElement")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	hasFocus, ok := response.(bool)
// 	if !ok {
// 		t.Fatal("focus eval response not a bool")
// 	}

// 	if hasFocus {
// 		t.Fatal("input already has focus")
// 	}

// 	done := hlivetest.AckWatcher(t, h.pwpage, "#btn_f")

// 	if err := h.pwpage.Click("#btn_f"); err != nil {
// 		t.Fatal(err)
// 	}

// 	<-done

// 	response, err = h.pwpage.EvalOnSelector("#in_f", "(el) => el === document.activeElement")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	hasFocus, ok = response.(bool)
// 	if !ok {
// 		t.Fatal("focus eval response not a bool")
// 	}

// 	if !hasFocus {
// 		t.Fatal("input doesn't have focus")
// 	}
// }
