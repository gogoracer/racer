package main

// func main() {
// 	srv, err := engine.NewPageServer(home)
// 	if err != nil {
// 		panic(err)
// 	}
// 	http.Handle("/", srv)

// 	log.Println("INFO: listing :3000")

// 	if err := http.ListenAndServe(":3000", nil); err != nil {
// 		log.Println("ERRO: http listen and serve: ", err)
// 	}
// }

// const pageStyle engine.HTML = `
// .box {
// 	display: grid;
// 	grid-template-columns: 1fr 1fr 1fr;
// 	gap: 0.5em;
// }

// input, select {
// 	width: 100%;
// }
// `

// func home() *engine.Page {
// 	page := engine.NewPage()
// 	page.DOM().Title().Add("Form Data Initial Sync Example")
// 	page.DOM().Head().Add(engine.NewTag("link",
// 		engine.Attrs{"rel": "stylesheet", "href": "https://cdn.simplecss.org/simple.min.css"}))
// 	page.DOM().Head().Add(engine.NewTag("style", pageStyle))

// 	var (
// 		formValsSync   [9]*engine.NodeBox[string]
// 		formValsNoSync [len(formValsSync)]*engine.NodeBox[string]
// 	)

// 	for i := 0; i < len(formValsSync); i++ {
// 		formValsSync[i] = engine.Box("")
// 		formValsNoSync[i] = engine.Box("")
// 	}

// 	page.DOM().Body().Add(
// 		engine.NewTag("header",
// 			engine.NewTag("h1", "Form Data Initial Sync"),
// 			engine.NewTag("p", "Some browsers, life FireFox, don't clear form field data after a page reload. "+
// 				"HLive will send this data to relevant inputs when this happens."),
// 		),
// 		engine.NewTag("main",
// 			engine.NewTag("p", "To test, using FireFox, change the fields below then reload."),
// 			engine.NewTag("h2", "Form"),
// 			engine.NewTag("form", engine.Class("box"),
// 				//
// 				// Text
// 				//
// 				engine.NewTag("div",
// 					engine.NewTag("label", "Text"),
// 					engine.NewTag("br"),
// 					engine.NewComponent("input", engine.Attrs{"type": "text"},
// 						engine.On("input", func(_ context.Context, e engine.Event) {
// 							if !e.IsInitial {
// 								formValsNoSync[0].Set(e.Value)
// 							}

// 							formValsSync[0].Set(e.Value)
// 						}),
// 					),
// 				),
// 				//
// 				// Password
// 				//
// 				engine.NewTag("div",
// 					engine.NewTag("label", "Password"),
// 					engine.NewTag("br"),
// 					engine.NewComponent("input", engine.Attrs{"type": "password"},
// 						engine.On("input", func(_ context.Context, e engine.Event) {
// 							if !e.IsInitial {
// 								formValsNoSync[1].Set(e.Value)
// 							}

// 							formValsSync[1].Set(e.Value)
// 						})),
// 				),
// 				//
// 				// Range
// 				//
// 				engine.NewTag("div",
// 					engine.NewTag("label", "Range"),
// 					engine.NewTag("br"),
// 					engine.NewComponent("input", engine.Attrs{"type": "range", "min": "0", "max": "1000"},
// 						engine.On("input", func(_ context.Context, e engine.Event) {
// 							if !e.IsInitial {
// 								formValsNoSync[2].Set(e.Value)
// 							}

// 							formValsSync[2].Set(e.Value)
// 						}),
// 					),
// 				),
// 				//
// 				// Multi select
// 				//
// 				engine.NewTag("div",
// 					engine.NewTag("label", "Multi Select"),
// 					engine.NewTag("br"),
// 					engine.NewComponent("select", engine.Attrs{"multiple": ""},
// 						engine.On("input", func(_ context.Context, e engine.Event) {
// 							fmt.Println(">>>", e.Value, e.Values)
// 							if !e.IsInitial {
// 								formValsNoSync[3].Set(strings.Join(e.Values, ", "))
// 							}

// 							formValsSync[3].Set(strings.Join(e.Values, ", "))
// 						}),
// 						engine.NewTag("option", engine.Attrs{"value": "dog"}, "Dog"),
// 						engine.NewTag("option", engine.Attrs{"value": "cat"}, "Cat"),
// 						engine.NewTag("option", engine.Attrs{"value": "bird"}, "Bird"),
// 						engine.NewTag("option", "Fox"),
// 					),
// 					engine.NewTag("br"),
// 					engine.NewTag("small", "Click + Ctl/Cmd to multi select"),
// 				),
// 				//
// 				// Radio
// 				//
// 				engine.NewTag("div",
// 					engine.NewTag("label", "Radio"),
// 					engine.NewTag("br"),
// 					engine.NewTag("label", engine.Attrs{"for": "radio_1"},
// 						engine.NewComponent("input",
// 							engine.Attrs{"type": "radio", "name": "radio", "value": "orange", "id": "radio_1"},
// 							engine.On("input", func(_ context.Context, e engine.Event) {
// 								if !e.IsInitial {
// 									formValsNoSync[4].Set(e.Value)
// 								}

// 								formValsSync[4].Set(e.Value)
// 							}),
// 						),
// 						" Orange"),
// 					engine.NewTag("label", engine.Attrs{"for": "radio_2"},
// 						engine.NewComponent("input",
// 							engine.Attrs{"type": "radio", "name": "radio", "value": "grape", "id": "radio_2"},
// 							engine.On("input", func(_ context.Context, e engine.Event) {
// 								if !e.IsInitial {
// 									formValsNoSync[4].Set(e.Value)
// 								}

// 								formValsSync[4].Set(e.Value)
// 							}),
// 						),
// 						" Grape"),
// 					engine.NewTag("label", engine.Attrs{"for": "radio_3"},
// 						engine.NewComponent("input",
// 							engine.Attrs{"type": "radio", "name": "radio", "value": "lemon", "id": "radio_3"},
// 							engine.On("input", func(_ context.Context, e engine.Event) {
// 								if !e.IsInitial {
// 									formValsNoSync[4].Set(e.Value)
// 								}

// 								formValsSync[4].Set(e.Value)
// 							}),
// 						),
// 						" Lemon"),
// 					engine.NewTag("label", engine.Attrs{"for": "radio_4"},
// 						engine.NewComponent("input",
// 							engine.Attrs{"type": "radio", "name": "radio", "value": "apple", "id": "radio_4"},
// 							engine.On("input", func(_ context.Context, e engine.Event) {
// 								if !e.IsInitial {
// 									formValsNoSync[4].Set(e.Value)
// 								}

// 								formValsSync[4].Set(e.Value)
// 							}),
// 						),
// 						" Apple"),
// 				),
// 				//
// 				// Checkbox
// 				//
// 				engine.NewTag("div",
// 					engine.NewTag("label", "Checkbox"),
// 					engine.NewTag("br"),

// 					engine.NewTag("label", engine.Attrs{"for": "checkbox_1"},
// 						engine.NewComponent("input", engine.Attrs{"type": "checkbox", "value": "north", "id": "checkbox_1"},
// 							engine.On("input", func(_ context.Context, e engine.Event) {
// 								formValsNoSync[5].Set("")
// 								formValsSync[5].Set("")

// 								if !e.IsInitial && e.Selected {
// 									formValsNoSync[5].Set(e.Value)
// 								}

// 								if e.Selected {
// 									formValsSync[5].Set(e.Value)
// 								}
// 							}),
// 						),
// 						"North"),
// 					engine.NewTag("label", engine.Attrs{"for": "checkbox_2"},
// 						engine.NewComponent("input", engine.Attrs{"type": "checkbox", "value": "east", "id": "checkbox_2"},
// 							engine.On("input", func(_ context.Context, e engine.Event) {
// 								formValsNoSync[6].Set("")
// 								formValsSync[6].Set("")

// 								if !e.IsInitial && e.Selected {
// 									formValsNoSync[6].Set(e.Value)
// 								}

// 								if e.Selected {
// 									formValsSync[6].Set(e.Value)
// 								}
// 							}),
// 						),
// 						"East"),
// 					engine.NewTag("label", engine.Attrs{"for": "checkbox_3"},
// 						engine.NewComponent("input", engine.Attrs{"type": "checkbox", "value": "south", "id": "checkbox_3"},
// 							engine.On("input", func(_ context.Context, e engine.Event) {
// 								formValsNoSync[7].Set("")
// 								formValsSync[7].Set("")

// 								if !e.IsInitial && e.Selected {
// 									formValsNoSync[7].Set(e.Value)
// 								}

// 								if e.Selected {
// 									formValsSync[7].Set(e.Value)
// 								}
// 							}),
// 						),
// 						"South"),
// 					engine.NewTag("label", engine.Attrs{"for": "checkbox_4"},
// 						engine.NewComponent("input", engine.Attrs{"type": "checkbox", "value": "west", "id": "checkbox_4"},
// 							engine.On("input", func(_ context.Context, e engine.Event) {
// 								formValsNoSync[8].Set("")
// 								formValsSync[8].Set("")

// 								if !e.IsInitial && e.Selected {
// 									formValsNoSync[8].Set(e.Value)
// 								}

// 								if e.Selected {
// 									formValsSync[8].Set(e.Value)
// 								}
// 							}),
// 						),
// 						"West"),
// 				),
// 			),
// 			//
// 			// Output
// 			//
// 			engine.NewTag("h2", "Server Side Data"),
// 			engine.NewTag("table",
// 				engine.NewTag("thead",
// 					engine.NewTag("tr",
// 						engine.NewTag("th", ""),
// 						engine.NewTag("th", "Sync"),
// 						engine.NewTag("th", "No Sync"),
// 					),
// 				),
// 				engine.NewTag("tbody",
// 					engine.NewTag("tr",
// 						engine.NewTag("td", "Text"),
// 						engine.NewTag("td", formValsSync[0]),
// 						engine.NewTag("td", formValsNoSync[0]),
// 					),
// 					engine.NewTag("tr",
// 						engine.NewTag("td", "Password", engine.NewTag("br"),
// 							engine.NewTag("small", "Browsers won't keep this on refresh")),
// 						engine.NewTag("td", formValsSync[1]),
// 						engine.NewTag("td", formValsNoSync[1]),
// 					),
// 					engine.NewTag("tr",
// 						engine.NewTag("td", "Range", engine.NewTag("br"),
// 							engine.NewTag("small", "Browsers set this to the mid point by default")),
// 						engine.NewTag("td", formValsSync[2]),
// 						engine.NewTag("td", formValsNoSync[2]),
// 					),
// 					engine.NewTag("tr",
// 						engine.NewTag("td", "Multi Select"),
// 						engine.NewTag("td", formValsSync[3]),
// 						engine.NewTag("td", formValsNoSync[3]),
// 					),
// 					engine.NewTag("tr",
// 						engine.NewTag("td", "Radio"),
// 						engine.NewTag("td", formValsSync[4]),
// 						engine.NewTag("td", formValsNoSync[4]),
// 					),
// 					engine.NewTag("tr",
// 						engine.NewTag("td", "Checkbox"),
// 						engine.NewTag("td",
// 							formValsSync[5], " ", formValsSync[6], " ", formValsSync[7], " ", formValsSync[8]),
// 						engine.NewTag("td",
// 							formValsNoSync[5], " ", formValsNoSync[6], " ", formValsNoSync[7], " ", formValsNoSync[8]),
// 					),
// 				),
// 			),
// 		),
// 	)

// 	return page
// }
