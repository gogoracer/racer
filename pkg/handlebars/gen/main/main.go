package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gogoracer/racer/pkg/handlebars/gen"
	"github.com/iancoleman/strcase"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

}

func run(ctx context.Context) error {
	handlebarsPath, err := filepath.Abs("..")
	if err != nil {
		return fmt.Errorf("could not get absolute path for %s: %v", handlebarsPath, err)
	}

	elements, err := gen.ScrapeHTMLSpec(ctx)
	if err != nil {
		return fmt.Errorf("could not scrape HTML spec: %v", err)
	}

	for _, element := range elements {
		contents := gen.GenerateElement(element)
		filename := fmt.Sprintf("element_%s.go", strcase.ToSnake(element.Tag))
		fullPath := filepath.Join(handlebarsPath, filename)

		if err := os.WriteFile(fullPath, []byte(contents), 0644); err != nil {
			return fmt.Errorf("could not write file %s: %v", fullPath, err)
		}
	}

	return nil
}
