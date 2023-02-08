package tailpipe

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/iancoleman/strcase"
)

func pascal(s string) string {
	return strcase.ToCamel(s)
}

func camel(s string) string {
	return strcase.ToLowerCamel(s)
}

func snake(s string) string {
	return strcase.ToSnake(s)
}

func kebab(s string) string {
	return strcase.ToKebab(s)
}

func atomicName(name string) string {
	if name == "DEFAULT" {
		name = ""
	}

	name = strings.ReplaceAll(name, ".", "TAILPIPEDOT")
	name = strings.ReplaceAll(name, "/", "TAILPIPESSLASH")
	name = pascal(name)
	name = strings.ReplaceAll(name, "TAILPIPEDOT", "dot")
	name = strings.ReplaceAll(name, "TAILPIPESSLASH", "over")
	return name
}

func Generate(ctx context.Context, genTmpDir, tailpipeDir string, cfg TailPipeConfig) error {
	spew.Config.Indent = " "
	res := generateFromConfig(cfg)

	resDir := filepath.Join(genTmpDir, "tailwind")
	if err := os.MkdirAll(resDir, 0755); err != nil {
		return fmt.Errorf("could not create generated directory: %v", err)
	}

	if err := os.WriteFile(filepath.Join(resDir, "tailpipe.go"), []byte(res), 0644); err != nil {
		return fmt.Errorf("could not write tailpipe.go: %v", err)
	}

	return nil
}
