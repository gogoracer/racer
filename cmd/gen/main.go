package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	gogglesgen "github.com/gogoracer/racer/pkg/goggles/gen"
	"github.com/gogoracer/racer/pkg/tailpipe"
	"golang.org/x/sync/errgroup"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
	log.Printf("done in %s", time.Since(start))
}

func run(ctx context.Context) error {
	racerDir, _ := filepath.Abs(".")
	if !strings.HasSuffix(racerDir, "racer") {
		return errors.New("must be run from the racer directory")
	}

	genTmpDir := filepath.Join(racerDir, "gentmp")

	eg := errgroup.Group{}

	eg.Go(func() error {
		return genGoggles(ctx, racerDir, genTmpDir)
	})

	eg.Go(func() error {
		return genTailpipe(ctx, racerDir, genTmpDir)
	})

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("could not generate goggles: %v", err)
	}

	return nil
}

func genGoggles(ctx context.Context, racerDir, genTmpDir string) error {
	eg := errgroup.Group{}
	gogglesPath := filepath.Join(racerDir, "pkg", "goggles")

	// eg.Go(func() error {
	// 	return gogglesgen.GenerateElements(ctx, genTmpDir, gogglesPath)
	// })

	eg.Go(func() error {
		return gogglesgen.GenerateIconify(ctx, genTmpDir, gogglesPath)
	})
	return eg.Wait()
}

func genTailpipe(ctx context.Context, racerDir, genTmpDir string) error {
	cfg := tailpipe.DefaultTailPipeConfig()
	tailpipePath := filepath.Join(racerDir, "pkg", "tailpipe")
	return tailpipe.Generate(ctx, genTmpDir, tailpipePath, cfg)
}
