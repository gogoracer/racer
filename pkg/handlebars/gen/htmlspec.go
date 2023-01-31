package gen

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/iancoleman/strcase"

	"github.com/cenkalti/backoff/v4"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type Attribute struct {
	Name            string
	Description     string
	ValidValueTypes []string
}

type EventHandler struct {
	Name        string
	Description string
	TargetsBody bool
}

type Element struct {
	Tag           string
	Description   string
	Interface     string
	Categories    []string
	Attributes    []*Attribute
	EventHandlers []*EventHandler
}

func GenerateElements(ctx context.Context, handlebarsPath string) error {

	elements, err := scrapeHTMLSpec(ctx)
	if err != nil {
		return fmt.Errorf("could not scrape HTML spec: %v", err)
	}

	for _, element := range elements {
		contents := GenerateElement(element)
		filename := fmt.Sprintf("element_%s.go", strcase.ToSnake(element.Tag))
		fullPath := filepath.Join(handlebarsPath, filename)

		if err := os.WriteFile(fullPath, []byte(contents), 0644); err != nil {
			return fmt.Errorf("could not write file %s: %v", fullPath, err)
		}
	}

	return nil
}

// Scrapes the W3C HTML spec for information about HTML elements and attributes.
func scrapeHTMLSpec(ctx context.Context) ([]*Element, error) {
	cachedJSON := "html_spec.json"
	if _, err := os.Stat(cachedJSON); err == nil {
		b, err := os.ReadFile(cachedJSON)
		if err != nil {
			return nil, fmt.Errorf("could not read cached JSON: %v", err)
		}
		var elements []*Element
		if err := json.Unmarshal(b, &elements); err != nil {
			elements = nil
			// If the cached JSON is invalid, just regenerate it.
			if err := os.Remove(cachedJSON); err != nil {
				return nil, fmt.Errorf("could not remove cached JSON: %v", err)
			}
		}

		if len(elements) > 0 {
			return elements, nil
		}
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		// chromedp.UserDataDir("someUserDir"),
		chromedp.Flag("headless", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("restore-on-startup", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(ctx, opts...)
	ctx, _ = chromedp.NewContext(allocCtx)

	defer closeSpec(ctx)

	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://html.spec.whatwg.org/multipage/indices.html`),
	); err != nil {
		return nil, fmt.Errorf("could not navigate to HTML spec: %v", err)
	}

	elementsMap, err := ReadElements(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not read elements: %v", err)
	}

	if err := ApplyAttributes(ctx, elementsMap); err != nil {
		return nil, fmt.Errorf("could not read attributes: %v", err)
	}

	if err := ApplyEventHandlers(ctx, elementsMap); err != nil {
		return nil, fmt.Errorf("could not read event handlers: %v", err)
	}

	b, err := json.MarshalIndent(elementsMap, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("could not marshal elements: %v", err)
	}

	if err := os.WriteFile(cachedJSON, b, 0644); err != nil {
		return nil, fmt.Errorf("could not write cached JSON: %v", err)
	}

	elements := make([]*Element, 0, len(elementsMap))
	for _, element := range elementsMap {
		elements = append(elements, element)
	}
	return elements, nil
}

func closeSpec(ctx context.Context) error {
	if err := chromedp.Run(ctx,
		page.Close(),
		chromedp.Stop()); err != nil {
		log.Fatal(err)
	}
	return nil
}

func ApplyAttributes(ctx context.Context, elements map[string]*Element) error {

	var (
		expectedColumns = int64(4)
		nodes           []*cdp.Node
		table           *cdp.Node
	)
	if err := chromedp.Run(ctx,
		chromedp.Nodes("#attributes-3 ~ table", &nodes, chromedp.ByQuery),
		chromedp.ActionFunc(func(c context.Context) error {
			table = nodes[0]
			if err := dom.RequestChildNodes(table.NodeID).WithDepth(-1).Do(c); err != nil {
				return fmt.Errorf("could not request child nodes: %v", err)
			}

			waitForNodeChildren(table)

			if table.NodeName != "TABLE" {
				return fmt.Errorf("expected table, got %q", table.NodeName)
			}

			thead := table.Children[1]
			if thead.Children[0].ChildNodeCount != expectedColumns {
				return fmt.Errorf("expected %d headings, got %d", expectedColumns, len(thead.Children))
			}

			tbody := table.Children[2]

			for _, tr := range tbody.Children {
				if tr.ChildNodeCount != expectedColumns {
					return fmt.Errorf("expected %d columns, got %d", expectedColumns, tr.ChildNodeCount)
				}

				if err := dom.ScrollIntoViewIfNeeded().WithNodeID(tr.NodeID).Do(c); err != nil {
					return err
				}

				attr := &Attribute{}

				td := tr.Children[0]
				code := td.Children[0]
				inner := code.Children[0]
				name := clean(inner.NodeValue)
				if name == "" {
					continue
				}
				attr.Name = name

				td = tr.Children[2]
				code = td.Children[0]
				attr.Description = clean(code.NodeValue)

				validValueLinks := tr.Children[3].Children
				for _, child := range validValueLinks {
					switch child.NodeName {
					case "A":
						for i := 0; i < len(child.Attributes); i += 2 {
							if child.Attributes[i] != "href" {
								continue
							}
							val := child.Attributes[i+1]
							end := strings.Split(val, "#")[1]
							attr.ValidValueTypes = append(attr.ValidValueTypes, clean(end))
						}
					case "CODE":
						inner := child.Children[0]
						v := clean(inner.NodeValue)
						attr.ValidValueTypes = append(attr.ValidValueTypes, clean(v))
					}
				}

				td = tr.Children[1]
				for _, child := range td.Children {
					v := child
					for v.NodeName != "#text" {
						v = v.Children[0]
					}

					name := clean(v.NodeValue)
					allHTMLElements := name == "HTML elements"
					el, knownElementName := elements[name]

					if !allHTMLElements && !knownElementName {
						continue
					}

					log.Print(attr.Name, " ", name)
					if allHTMLElements {
						for _, element := range elements {
							element.Attributes = append(element.Attributes, attr)
						}
					} else {
						alreadyHasAttributeIndex := -1
						for i, a := range el.Attributes {
							if a.Name == attr.Name {
								alreadyHasAttributeIndex = i
								break
							}
						}

						if alreadyHasAttributeIndex != -1 {
							el.Attributes[alreadyHasAttributeIndex] = attr
						} else {
							el.Attributes = append(el.Attributes, attr)
						}
					}
				}
			}

			return nil
		}),
	); err != nil {
		return fmt.Errorf("could not read elements: %v", err)
	}

	return nil

}

// func LoadGlobalAttributeNames(ctx context.Context) ([]string, error) {
// 	nodes := []*cdp.Node{}
// 	globalAttributes := []string{}
// 	if err := chromedp.Run(ctx,
// 		chromedp.Nodes("#global-attributes ~ ul", &nodes, chromedp.ByQuery),
// 		chromedp.ActionFunc(func(c context.Context) error {
// 			ul := nodes[0]
// 			if err := dom.RequestChildNodes(ul.NodeID).WithDepth(-1).Do(c); err != nil {
// 				return err
// 			}

// 			waitForNodeChildren(ul)

// 			for _, li := range ul.Children {
// 				if li.NodeName != "LI" {
// 					continue
// 				}

// 				code := li.Children[0]
// 				if code.NodeName != "CODE" {
// 					continue
// 				}

// 				a := code.Children[0]
// 				if a.NodeName != "A" {
// 					continue
// 				}

// 				if err := dom.ScrollIntoViewIfNeeded().WithNodeID(a.NodeID).Do(c); err != nil {
// 					return err
// 				}

// 				v := a.Children[0].NodeValue
// 				v = clean(v)

// 				globalAttributes = append(globalAttributes, v)
// 			}

// 			return nil
// 		}),
// 	); err != nil {
// 		return nil, fmt.Errorf("could not load spec: %v", err)
// 	}

// 	return globalAttributes, nil
// }

func ReadElements(ctx context.Context) (map[string]*Element, error) {
	// globalAttributeNames, err := LoadGlobalAttributeNames(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not load global attributes: %v", err)
	// }

	// globalAttributes := []*Attribute{}
	// for _, name := range globalAttributeNames {
	// 	globalAttributes = append(globalAttributes, attributes[name])
	// }

	var (
		expectedColumns = int64(7)
		nodes           []*cdp.Node
		table           *cdp.Node
		elements        = []*Element{}
	)
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible("#elements-3 ~ table"),
		chromedp.Nodes("#elements-3 ~ table", &nodes, chromedp.ByQuery),
		chromedp.ActionFunc(func(c context.Context) error {
			table = nodes[0]
			if err := dom.RequestChildNodes(table.NodeID).WithDepth(-1).Do(c); err != nil {
				return fmt.Errorf("could not request child nodes: %v", err)
			}

			waitForNodeChildren(table)

			if table.NodeName != "TABLE" {
				return fmt.Errorf("expected table, got %q", table.NodeName)
			}

			thead := table.Children[1]
			if thead.Children[0].ChildNodeCount != expectedColumns {
				return fmt.Errorf("expected 7 headings, got %d", len(thead.Children))
			}

			tbody := table.Children[2]

			for _, tr := range tbody.Children {
				if tr.ChildNodeCount != expectedColumns {
					return fmt.Errorf("expected 7 columns, got %d", tr.ChildNodeCount)
				}

				if err := dom.ScrollIntoViewIfNeeded().WithNodeID(tr.NodeID).Do(c); err != nil {
					return err
				}

				el := &Element{}

				td := tr.Children[0]
				tdInner := td.Children[0]
				for _, child := range tdInner.Children {
					if child.ChildNodeCount == 0 {
						continue
					}
					a := tdInner.Children[0]
					if a.NodeName != "A" {
						continue
					}
					aInner := a.Children[0]
					el.Tag = clean(aInner.NodeValue)
					break
				}

				if el.Tag == "" {
					continue
				}

				td = tr.Children[1]
				tdInner = td.Children[0]
				el.Description = clean(tdInner.NodeValue)

				categoryLinks := tr.Children[2].Children
				for _, child := range categoryLinks {
					if child.NodeName != "A" {
						continue
					}
					el.Categories = append(el.Categories, clean(child.Children[0].NodeValue))
				}

				td = tr.Children[6]
				code := td.Children[0]
				a := code.Children[0]
				v := a.Children[0].NodeValue
				el.Interface = clean(v)

				elements = append(elements, el)
			}

			return nil
		}),
	); err != nil {
		return nil, fmt.Errorf("could not read elements: %v", err)
	}

	elementsMap := map[string]*Element{}
	for _, el := range elements {
		elementsMap[el.Tag] = el
	}

	return elementsMap, nil
}

func ApplyEventHandlers(ctx context.Context, elements map[string]*Element) error {

	var (
		expectedColumns = int64(4)
		nodes           []*cdp.Node
		table           *cdp.Node
		eventHandlers   = []*EventHandler{}
	)
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible("#ix-event-handlers"),
		chromedp.Nodes("#ix-event-handlers", &nodes, chromedp.ByQuery),
		chromedp.ActionFunc(func(c context.Context) error {
			table = nodes[0]
			if err := dom.RequestChildNodes(table.NodeID).WithDepth(-1).Do(c); err != nil {
				return fmt.Errorf("could not request child nodes: %v", err)
			}

			waitForNodeChildren(table)

			if table.NodeName != "TABLE" {
				return fmt.Errorf("expected table, got %q", table.NodeName)
			}

			thead := table.Children[1]
			if thead.Children[0].ChildNodeCount != expectedColumns {
				return fmt.Errorf("expected %d headings, got %d", expectedColumns, len(thead.Children))
			}

			tbody := table.Children[2]

			for _, tr := range tbody.Children {
				if tr.ChildNodeCount != expectedColumns {
					return fmt.Errorf("expected %d columns, got %d", expectedColumns, tr.ChildNodeCount)
				}

				if err := dom.ScrollIntoViewIfNeeded().WithNodeID(tr.NodeID).Do(c); err != nil {
					return err
				}

				evtHandler := &EventHandler{}

				td := tr.Children[0]
				tdInner := td.Children[0]
				code := tdInner.Children[0]
				evtHandler.Name = clean(code.NodeValue)

				td = tr.Children[1]
				v := td
				for v.NodeName != "#text" {
					v = v.Children[0]
				}
				evtHandler.TargetsBody = v.NodeValue == "body"

				td = tr.Children[2]
				parts := make([]string, len(td.Children))
				var err error
				for i, child := range td.Children {
					parts[i], err = dom.GetOuterHTML().WithNodeID(child.NodeID).Do(c)
					if err != nil {
						return fmt.Errorf("could not get outer html: %v", err)
					}
				}
				evtHandler.Description = clean(strings.Join(parts, " "))

				for _, el := range elements {
					isBody := el.Tag == "body"
					if !isBody && evtHandler.TargetsBody {
						continue
					}
					el.EventHandlers = append(el.EventHandlers, evtHandler)
				}

				eventHandlers = append(eventHandlers, evtHandler)
			}

			return nil
		}),
	); err != nil {
		return fmt.Errorf("could not read elements: %v", err)
	}

	return nil
}

func waitForNodeChildren(n *cdp.Node) {
	b := backoff.NewExponentialBackOff()
	for len(n.Children) != int(n.ChildNodeCount) {
		d := b.NextBackOff()
		log.Printf("waiting for children to load, got %d, expected %d, waiting %s", len(n.Children), n.ChildNodeCount, d)
		time.Sleep(d)
	}
}
func clean(v string) string {

	v = strings.TrimSpace(v)
	return v

}
