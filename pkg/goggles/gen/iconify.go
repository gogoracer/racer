package gen

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/divan/num2words"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/goccy/go-json"
	"github.com/valyala/bytebufferpool"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

func GenerateIconify(ctx context.Context, gentmpDir, gogglesPath string) error {
	tmpParentDir := filepath.Join(gentmpDir, "iconify")
	tmpIconSetsDir := filepath.Join(tmpParentDir, "json")
	if _, err := os.Stat(filepath.Join(tmpParentDir, "collections.json")); os.IsNotExist(err) {
		if err := os.MkdirAll(tmpParentDir, 0755); err != nil {
			return fmt.Errorf("could not create iconify directory: %v", err)
		}

		iconCollections := map[string]iconBasicCollectionInfo{}
		if err := updateIconifyCache(ctx, tmpParentDir, "collections.json", &iconCollections); err != nil {
			return fmt.Errorf("could not get iconify collections: %v", err)
		}
		log.Print("got icon collections")

		eg, ctx := errgroup.WithContext(ctx)

		for collectionName := range iconCollections {
			details := iconCollectionDetailsInfo{}
			collectionName := collectionName

			eg.Go(func() error {
				iconifyAssetPath := filepath.Join(tmpIconSetsDir, collectionName+".json")
				if err := updateIconifyCache(ctx, tmpParentDir, iconifyAssetPath, &details); err != nil {
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

	iconifyPath := filepath.Join(gogglesPath, "iconify")
	if err := os.RemoveAll(iconifyPath); err != nil {
		return fmt.Errorf("could not remove iconify directory: %v", err)
	}
	if err := os.MkdirAll(iconifyPath, 0755); err != nil {
		return fmt.Errorf("could not create iconify directory: %v", err)
	}

	numberWordGroupsRegex := regexp.MustCompile(`(?P<number>\d*)*(?P<word>\D*)`)

	// walk the details and generate the iconify.go file
	if err := filepath.WalkDir(tmpIconSetsDir, func(path string, d os.DirEntry, err error) error {
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

		// some icons have a prefix that is not a valid Go identifier.  Specifically, the
		// use of numbers, dashes, and underscores.  We need to convert these to valid
		// Go identifiers.

		namesUsed := sets.NewString()
		namedIcons := []namedIcon{}
		for iconName, icon := range details.Icons {
			matches := numberWordGroupsRegex.FindAllStringSubmatch(iconName, -1)
			revisedName := iconName

			if len(matches) > 0 {
				parts := []string{}

				for _, match := range matches {
					numbersStr := match[numberWordGroupsRegex.SubexpIndex("number")]
					if numbersStr != "" {
						number, err := strconv.Atoi(numbersStr)
						if err != nil {
							return fmt.Errorf("could not convert %s to number: %v", numbersStr, err)
						}
						numbersStr = num2words.Convert(number) + "_"
					}

					word := match[numberWordGroupsRegex.SubexpIndex("word")]
					parts = append(parts, numbersStr, word)
				}

				revisedName = pascal(strings.Join(parts, " "))
			}

			if namesUsed.Has(revisedName) {
				continue
			}

			namedIcons = append(namedIcons, namedIcon{
				OriginalName: iconName,
				Name:         revisedName,
				Icon:         icon,
			})
			namesUsed.Insert(revisedName)
		}

		slices.SortFunc(namedIcons, func(a, b namedIcon) bool {
			return a.Name < b.Name
		})

		packageName := lower(snake(details.Prefix))
		if goKeywords.Has(packageName) {
			packageName = fmt.Sprintf("%s_icons", packageName)
		}

		iconPkg := iconPackage{
			Name:   packageName,
			Icons:  namedIcons,
			Width:  unknownToDimension(details.Width),
			Height: unknownToDimension(details.Height),
		}
		content := GenerateIcon(iconPkg)

		packagePath := filepath.Join(iconifyPath, packageName)
		if err := os.MkdirAll(packagePath, 0755); err != nil {
			return fmt.Errorf("could not create iconify directory: %v", err)
		}
		fullPath := filepath.Join(packagePath, fmt.Sprintf("%s.go", packageName))
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("could not write file %s: %v", fullPath, err)
		}

		log.Printf("wrote %s", fullPath)

		return nil
	}); err != nil {
		return fmt.Errorf("could not walk iconify directory: %v", err)
	}

	return nil
}

var goKeywords = sets.New(
	"break", "case", "chan", "const", "continue", "default", "defer",
	"else", "fallthrough", "for", "func", "go", "goto", "if", "import",
	"interface", "map", "package", "range", "return", "select", "struct",
	"switch", "type", "var",
)

func updateIconifyCache(ctx context.Context, gentmpDir, iconifyJSONPath string, v interface{}) error {
	fullURL := fmt.Sprintf("https://raw.githubusercontent.com/iconify/icon-sets/master/%s", iconifyJSONPath)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not get iconify %s: %v", fullURL, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("could not get iconify %s: %v", fullURL, res.Status)
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

	jsonPath := filepath.Join(gentmpDir, iconifyJSONPath)
	subDirPath := filepath.Dir(jsonPath)
	if err := os.MkdirAll(subDirPath, 0755); err != nil {
		return fmt.Errorf("could not create directory %s: %v", subDirPath, err)
	}

	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %v", err)
	}
	if err := os.WriteFile(jsonPath, bytes, 0644); err != nil {
		return fmt.Errorf("could not write file %s: %v", jsonPath, err)
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

type iconPackage struct {
	Name          string
	Icons         []namedIcon
	Width, Height int
}
type namedIcon struct {
	OriginalName string
	Name         string
	Icon         icon
}
