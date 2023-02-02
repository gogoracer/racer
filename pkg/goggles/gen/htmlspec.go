package gen

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/goccy/go-json"
	"github.com/iancoleman/strcase"
	"k8s.io/apimachinery/pkg/util/sets"
)

type Attribute struct {
	Name            string
	Description     string
	ValidValueTypes []string
}

type EventHandler struct {
	Name        string
	Description string
	Event       *Event
}

type Event struct {
	Name        string
	Description string
	Interface   string
	Attributes  []EventAttribute
}

type EventAttribute struct {
	Name     string
	Type     string
	Optional bool
	Comment  string
}

type Element struct {
	Tag                       string
	Description               string
	Interface                 string
	Categories                []string
	ChildElementsOrCategories []string
	Attributes                []*Attribute
	EventHandlers             []*EventHandler
}

func GenerateElements(ctx context.Context, gogglesPath string) error {

	elements, err := scrapeHTMLSpec(ctx)
	if err != nil {
		return fmt.Errorf("could not scrape HTML spec: %v", err)
	}

	categoryElements := map[string][]*Element{}
	for _, element := range elements {
		for _, category := range element.Categories {
			categoryElements[category] = append(categoryElements[category], element)
		}
	}

	validChildElements := map[string][]string{}
	for _, element := range elements {
		for _, child := range element.ChildElementsOrCategories {
			if _, ok := categoryElements[child]; ok {
				childElements := categoryElements[child]

				for _, childElement := range childElements {
					validChildElements[element.Tag] = append(validChildElements[element.Tag], childElement.Tag)
				}
			}
		}
	}

	for _, element := range elements {
		contents := GenerateElement(element)
		filename := fmt.Sprintf("element_%s.go", strcase.ToSnake(element.Tag))
		fullPath := filepath.Join(gogglesPath, filename)

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

	htmlSpecDoc, err := getDoc("https://html.spec.whatwg.org")
	if err != nil {
		return nil, fmt.Errorf("could not get HTML spec: %v", err)
	}

	elementsMap, err := readElements(htmlSpecDoc)
	if err != nil {
		return nil, fmt.Errorf("could not read elements: %v", err)
	}

	if err := applyAttributes(htmlSpecDoc, elementsMap); err != nil {
		return nil, fmt.Errorf("could not read attributes: %v", err)
	}

	if err := applyEventHandlers(htmlSpecDoc, elementsMap); err != nil {
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

func readElements(htmlSpecDoc *goquery.Document) (map[string]*Element, error) {
	expectedColumns := 7
	elements := map[string]*Element{}
	var err error
	htmlSpecDoc.Find("#elements-3 ~ table").First().Find("tbody > tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		trChildren := s.Children()
		columnCount := trChildren.Length()
		if columnCount != expectedColumns {
			var html string
			html, err = s.Html()
			if err != nil {
				log.Fatal(err)
			}
			err = fmt.Errorf("unexpected column count for row %d: %d, %s", i, columnCount, html)
			return false
		}

		el := &Element{
			Tag:         trChildren.Eq(0).Find("a").Text(),
			Description: trChildren.Eq(1).Text(),
			Interface:   trChildren.Eq(6).Children().Filter("a").First().Text(),
		}
		trChildren.Eq(2).Children().Filter("a").Each(func(i int, s *goquery.Selection) {
			c := strcase.ToSnake(s.Text())
			el.Categories = append(el.Categories, c)
		})
		trChildren.Eq(4).Children().Filter("a").Each(func(i int, s *goquery.Selection) {
			el.ChildElementsOrCategories = append(el.ChildElementsOrCategories, s.Text())
		})

		elements[el.Tag] = el

		return true
	})
	if err != nil {
		return nil, fmt.Errorf("could not read elements: %v", err)
	}

	return elements, nil
}

func applyAttributes(htmlSpecDoc *goquery.Document, elements map[string]*Element) error {
	if len(elements) == 0 {
		return fmt.Errorf("no elements")
	}
	expectedColumns := 4

	const validAllHTMLElementsText = "HTML elements"

	categoriesElements := map[string]map[string]*Element{}
	for _, element := range elements {
		for _, category := range element.Categories {
			if _, ok := categoriesElements[category]; !ok {
				categoriesElements[category] = map[string]*Element{}
			}
			categoriesElements[category][element.Tag] = element
		}
	}
	attributes := map[string]*Attribute{}
	elementsAttributeNames := map[string]sets.Set[string]{}
	for _, element := range elements {
		elementsAttributeNames[element.Tag] = sets.New[string]()
	}

	var (
		err                error
		knownBadCategories = sets.NewString(
			"form_associated_custom_elements",
		)
	)
	htmlSpecDoc.Find("#attributes-1 > tbody > tr").EachWithBreak(func(i int, rows *goquery.Selection) bool {
		columns := rows.Children()
		if columns.Length() != expectedColumns {
			err = fmt.Errorf("expected %d columns, got %d", expectedColumns, columns.Length())
			return false
		}

		attr := &Attribute{
			Name:        columns.Eq(0).Find("code").Text(),
			Description: strings.TrimSpace(columns.Eq(2).Text()),
		}
		columns.Eq(3).Find("a, code").Each(func(i int, a *goquery.Selection) {
			attr.ValidValueTypes = append(attr.ValidValueTypes, strcase.ToSnake(a.Text()))
		})
		attributes[attr.Name] = attr

		columns.Eq(1).Find("a").EachWithBreak(func(i int, a *goquery.Selection) bool {
			text := a.Text()
			if text == validAllHTMLElementsText {
				for _, set := range elementsAttributeNames {
					set.Insert(attr.Name)
				}
			} else {
				element, ok := elements[text]
				if !ok {
					// not element, but maybe an category
					c := strcase.ToSnake(text)
					categoryElements, ok := categoriesElements[c]
					if !ok {
						if knownBadCategories.Has(c) {
							return true
						}

						// if it's not a category, then it's an error
						err = fmt.Errorf("could not category element %s", text)
						return false
					}

					for _, element := range categoryElements {
						elementsAttributeNames[element.Tag].Insert(attr.Name)
					}
				}

				elementsAttributeNames[element.Tag].Insert(attr.Name)
			}

			return true
		})
		return err == nil
	})
	if err != nil {
		return fmt.Errorf("could not read attributes: %v", err)
	}

	for elementTag, attributeNames := range elementsAttributeNames {
		e, ok := elements[elementTag]
		if !ok {
			return fmt.Errorf("could not find element %s", elementTag)
		}
		for _, attributeName := range attributeNames.UnsortedList() {
			a, ok := attributes[attributeName]
			if !ok {
				return fmt.Errorf("could not find attribute %s", attributeName)
			}
			e.Attributes = append(e.Attributes, a)
		}
	}

	return nil
}

func applyEventHandlers(htmlSpecDoc *goquery.Document, elements map[string]*Element) (err error) {
	const domSpecURL = "https://dom.spec.whatwg.org"
	domSpecDoc, err := getDoc(domSpecURL)
	if err != nil {
		return fmt.Errorf("could not get dom spec: %v", err)
	}

	const pointerSpecURL = "https://w3c.github.io/pointerevents" // hot garbage
	const workingSpecURLWorking = "https://www.w3.org/TR/pointerevents"
	pointerSpecDoc, err := getDoc(workingSpecURLWorking)
	if err != nil {
		return fmt.Errorf("could not get pointer spec: %v", err)
	}

	attributeInfoRegex := regexp.MustCompile(`attribute (?P<type>[\w]*)(?P<optional>\?*) (?P<name>\w*);( // (?P<comment>(.*)))?`)
	pointerAttributeInfoRegex := regexp.MustCompile(`attribute (?P<type>\w*)\s*(?P<name>\w*)`)

	events := map[string]*Event{}
	htmlSpecDoc.Find("#events-2 ~ table").First().Find("tbody > tr").EachWithBreak(func(i int, rows *goquery.Selection) bool {
		columns := rows.Children()

		event := &Event{
			Name:        columns.Eq(0).Find("code").Text(),
			Description: strings.TrimSpace(columns.Eq(3).Text()),
			Interface:   columns.Eq(1).Find("code").Text(),
		}

		link, exists := columns.Eq(1).Find("a").Attr("href")
		if !exists {
			return true
		}
		relativeLink := link[strings.Index(link, "#"):]

		var (
			idlBlock string
			idlRegex = attributeInfoRegex
		)
		switch {
		case strings.HasPrefix(link, domSpecURL):
			idlBlock = domSpecDoc.Find(relativeLink + " ~ pre.idl").First().Text()
		case strings.HasPrefix(link, "#"):
			idlBlock = htmlSpecDoc.Find(link).Parent().Parent().Text()
		case strings.HasPrefix(link, pointerSpecURL):
			idlBlock = pointerSpecDoc.Find(relativeLink).Text()
			idlRegex = pointerAttributeInfoRegex
		default:
			err = fmt.Errorf("expected link to be relative to spec, got %s", link)
			return false
		}
		if idlBlock == "" {
			err = fmt.Errorf("could not find IDL block for %s", link)
			return false
		}

		matches := idlRegex.FindAllStringSubmatch(idlBlock, -1)
		for _, match := range matches {
			evtAttribute := EventAttribute{
				Name: match[idlRegex.SubexpIndex("name")],
				Type: match[idlRegex.SubexpIndex("type")],
			}
			event.Attributes = append(event.Attributes, evtAttribute)

		}
		events[event.Name] = event
		return true
	})
	if err != nil {
		return fmt.Errorf("could not read events: %v", err)
	}

	htmlSpecDoc.Find("#ix-event-handlers > tbody > tr").EachWithBreak(func(i int, rows *goquery.Selection) bool {
		columns := rows.Children()
		if columns.Length() != 4 {
			err = fmt.Errorf("expected %d columns, got %d", 4, columns.Length())
			return false
		}

		evtHandler := &EventHandler{
			Name:        columns.Eq(0).Find("code").Text(),
			Description: strings.TrimSpace(columns.Eq(2).Text()),
		}

		target := columns.Eq(1).Find("a").Text()
		targetsBody := strings.Contains(target, "body")
		if targetsBody {
			elements["body"].EventHandlers = append(elements["body"].EventHandlers, evtHandler)
		} else {
			for _, element := range elements {
				element.EventHandlers = append(element.EventHandlers, evtHandler)
			}
		}

		// TODO: Can't get this in a clean way from the spec, so we'll just hardcode it
		// https://html.spec.whatwg.org/multipage/indices.html#events-2

		return true
	})
	if err != nil {
		return fmt.Errorf("could not read event handlers: %v", err)
	}

	return nil

	// 		thead := table.Children[1]
	// 		if thead.Children[0].ChildNodeCount != expectedColumns {
	// 			return fmt.Errorf("expected %d headings, got %d", expectedColumns, len(thead.Children))
	// 		}

	// 		tbody := table.Children[2]

	// 		for _, tr := range tbody.Children {
	// 			if tr.ChildNodeCount != expectedColumns {
	// 				return fmt.Errorf("expected %d columns, got %d", expectedColumns, tr.ChildNodeCount)
	// 			}

	// 			if err := dom.ScrollIntoViewIfNeeded().WithNodeID(tr.NodeID).Do(c); err != nil {
	// 				return err
	// 			}

	// 			evtHandler := &EventHandler{}

	// 			td := tr.Children[0]
	// 			tdInner := td.Children[0]
	// 			code := tdInner.Children[0]
	// 			evtHandler.Name = code.NodeValue

	// 			td = tr.Children[1]
	// 			v := td
	// 			for v.NodeName != "#text" {
	// 				v = v.Children[0]
	// 			}
	// 			evtHandler.TargetsBody = v.NodeValue == "body"

	// 			td = tr.Children[2]
	// 			parts := make([]string, len(td.Children))
	// 			var err error
	// 			for i, child := range td.Children {
	// 				parts[i], err = dom.GetOuterHTML().WithNodeID(child.NodeID).Do(c)
	// 				if err != nil {
	// 					return fmt.Errorf("could not get outer html: %v", err)
	// 				}
	// 			}
	// 			evtHandler.Description = strings.Join(parts, " ")

	// 			for _, el := range elements {
	// 				isBody := el.Tag == "body"
	// 				if !isBody && evtHandler.TargetsBody {
	// 					continue
	// 				}
	// 				el.EventHandlers = append(el.EventHandlers, evtHandler)
	// 			}

	// 			eventHandlers = append(eventHandlers, evtHandler)
	// 		}

	// 		return nil
	// 	}),
	// ); err != nil {
	// 	return fmt.Errorf("could not read elements: %v", err)
	// }

}

func getDoc(url string) (*goquery.Document, error) {
	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get UI events spec: %v", err)
	}
	defer res.Body.Close()
	domSpecDoc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse UI events spec: %v", err)
	}

	return domSpecDoc, nil
}
