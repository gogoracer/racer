package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/gogoracer/racer/pkg/handlebars/gen"
	"golang.org/x/sync/errgroup"
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

	eg := errgroup.Group{}

	eg.Go(func() error {
		return gen.GenerateElements(ctx, handlebarsPath)
	})

	eg.Go(func() error {
		return gen.GenerateIconify(ctx, handlebarsPath)
	})

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("could not generate handlebars: %v", err)
	}

	return nil
}
