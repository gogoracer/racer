package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/gogoracer/racer/pkg/goggles/gen"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

}

func run(ctx context.Context) error {
	racerDir, _ := filepath.Abs(".")
	if !strings.HasSuffix(racerDir, "racer") {
		return errors.New("must be run from the racer directory")
	}

	gogglesPath := filepath.Join(racerDir, "pkg", "goggles")

	eg := errgroup.Group{}

	// eg.Go(func() error {
	// 	return gen.GenerateElements(ctx, gogglesPath)
	// })

	eg.Go(func() error {
		return gen.GenerateIconify(ctx, gogglesPath)
	})

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("could not generate goggles: %v", err)
	}

	return nil
}
