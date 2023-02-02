package gen

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/goccy/go-json"
	"github.com/iancoleman/strcase"
	"github.com/valyala/bytebufferpool"
	"golang.org/x/sync/errgroup"
)

func GenerateIconify(ctx context.Context, handlebarsPath string) error {
	detailsDir := "iconify/json"
	if _, err := os.Stat("iconify/collections.json"); os.IsNotExist(err) {
		if err := os.MkdirAll(detailsDir, 0755); err != nil {
			return fmt.Errorf("could not create iconify directory: %v", err)
		}

		iconCollections := map[string]iconBasicCollectionInfo{}
		if err := updateIconifyCache(ctx, "collections", &iconCollections); err != nil {
			return fmt.Errorf("could not get iconify collections: %v", err)
		}
		log.Print("got icon collections")

		eg, ctx := errgroup.WithContext(ctx)

		for collectionName := range iconCollections {
			details := iconCollectionDetailsInfo{}
			collectionName := collectionName

			eg.Go(func() error {
				if err := updateIconifyCache(ctx, "json/"+collectionName, &details); err != nil {
					log.Printf("!!!! could not get iconify collection %s: %v", collectionName, err)
				} else {
					log.Printf("got icon collection %s", collectionName)
				}
				return nil
			})
		}

		if err := eg.Wait(); err != nil {
			return fmt.Errorf("could not get iconify collections: %v", err)
		}
	}

	iconifyPath := filepath.Join(handlebarsPath, "iconify")
	if err := os.MkdirAll(iconifyPath, 0755); err != nil {
		return fmt.Errorf("could not create iconify directory: %v", err)
	}

	// walk the details and generate the iconify.go file
	if err := filepath.WalkDir(detailsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not read file %s: %v", path, err)
		}

		details := iconCollectionDetailsInfo{}
		if err := json.Unmarshal(b, &details); err != nil {
			return fmt.Errorf("could not unmarshal JSON: %v", err)
		}

		content := GenerateIcon(details)
		snaked := strcase.ToSnake(details.Prefix)
		packageName := lower(snaked)
		packageDir := filepath.Join(iconifyPath, packageName)
		if err := os.MkdirAll(packageDir, 0755); err != nil {
			return fmt.Errorf("could not create iconify directory: %v", err)
		}
		fullPath := filepath.Join(packageDir, fmt.Sprintf("%s.go", snaked))
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("could not write file %s: %v", fullPath, err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("could not walk iconify directory: %v", err)
	}

	return nil
}

func updateIconifyCache(ctx context.Context, subpath string, v interface{}) error {
	fullURL := fmt.Sprintf("https://raw.githubusercontent.com/iconify/icon-sets/master/%s.json", subpath)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not get iconify %s: %v", subpath, err)
	}
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return fmt.Errorf("could not read body: %v", err)
	}

	b := buf.Bytes()
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	if err := writeJSON(subpath, v); err != nil {
		return fmt.Errorf("could not write JSON: %v", err)
	}

	return nil
}

var rx *regexp.Regexp

func writeJSON(subpath string, v interface{}) error {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %v", err)
	}

	name := fmt.Sprintf(
		"iconify/%s.json",
		strcase.ToSnake(subpath),
	)

	if err := os.WriteFile(name, bytes, 0644); err != nil {
		return fmt.Errorf("could not write file %s: %v", name, err)
	}

	return nil
}

type author struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type license struct {
	Title string `json:"title,omitempty"`
	SPDX  string `json:"spdx,omitempty"`
	URL   string `json:"url,omitempty"`
}

type iconBasicCollectionInfo struct {
	Name     string  `json:"name,omitempty"`
	Total    int     `json:"total,omitempty"`
	Version  string  `json:"version,omitempty"`
	Author   author  `json:"author,omitempty"`
	License  license `json:"license,omitempty"`
	Category string  `json:"category,omitempty"`
}

type iconCollectionDetailsInfo struct {
	Prefix string `json:"prefix,omitempty"`
	Info   struct {
		Name     string   `json:"name,omitempty"`
		Total    int      `json:"total,omitempty"`
		Version  string   `json:"version,omitempty"`
		Author   author   `json:"author,omitempty"`
		License  license  `json:"license,omitempty"`
		Samples  []string `json:"samples,omitempty"`
		Height   int      `json:"height,omitempty"`
		Category string   `json:"category,omitempty"`
		Palette  bool     `json:"palette,omitempty"`
	} `json:"info,omitempty"`
	LastModified int64             `json:"lastModified,omitempty"`
	Icons        map[string]icon   `json:"icons,omitempty"`
	Suffixes     map[string]string `json:"suffixes,omitempty"`
	Width        interface{}       `json:"width,omitempty"`
	Height       interface{}       `json:"height,omitempty"`
}

type icon struct {
	SvgBody string      `json:"body,omitempty"`
	Width   interface{} `json:"width,omitempty"`
	Height  interface{} `json:"height,omitempty"`
}

func unknownToDimension(x interface{}) int {
	switch v := x.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case []int:
		return v[0]
	default:
		return 0
	}
}
