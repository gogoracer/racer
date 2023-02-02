package engine

import (
	_ "embed"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/gogoracer/racer/pkg/gas"
)

// //go:embed js/page.js
// var PageJavaScript []byte

// Diff Diffs are from old to new
type diffInfo struct {
	// Root element, where to start the path search from, "doc" is a special case that means the browser document
	Root string
	// Position of each child
	PathIndicies []int
	Type         gas.DiffType
	Tag          Tagger
	Text         *string
	Attribute    *Attribute
	HTML         *HTML
	// Not used for render but for Lifecycle events
	OldNode any
}

func NewDiffer() *Differ {
	// jsB := PageJavaScript

	// m := minify.New()
	// m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	// jsMin, err := m.Bytes("application/javascript", jsB)
	// if err == nil {
	// 	jsB = jsMin
	// } else {
	// 	_ = "TODO: DELANEY deal with logging"
	// 	// LoggerDev.Err(err).Msg("NewDiffer: minify")
	// }

	// rootPath, err := filepath.Abs("../../pkg/headlamp/src/")
	// if err != nil {
	// 	panic(err)
	// }

	// result := esbuild.Build(esbuild.BuildOptions{
	// 	EntryPoints: []string{
	// 		filepath.Join(rootPath, "index.ts"),
	// 	},
	// 	Sourcemap:  esbuild.SourceMapInline,
	// 	SourceRoot: rootPath,
	// 	Target:     esbuild.ESNext,
	// 	Bundle:     true,
	// })

	// if len(result.Errors) > 0 {
	// 	panic(fmt.Sprintf("esbuild: %v", result.Errors))
	// }
	// bundle := result.OutputFiles[0].Contents
	// log.Printf("bundle size: %02f", float64(len(bundle))/1024)

	return &Differ{
		logger: zerolog.Nop(),
		// JavaScript: result.OutputFiles[0].Contents,
	}
}

type Differ struct {
	logger     zerolog.Logger
	JavaScript []byte
}

func (d *Differ) SetLogger(logger zerolog.Logger) {
	d.logger = logger
}

// Trees diff two node tress
//
// Path: childIndex>childIndex
// Path: 0>1>3
//
// After tree copy you only have Tagger (with []Attribute), HTML, and strings. Then can be grouped in a NodeGroup
func (d *Differ) Trees(selector string, pathIndicies []int, oldNode, newNode any) ([]*diffInfo, error) {
	diffs := []*diffInfo{}

	d.logger.Trace().Str("sel", selector).Any("path", pathIndicies).Msg("diffTrees")

	// More nodes in new node
	if oldNode == nil && newNode != nil {
		diffs = append(diffs, diffCreate(selector, pathIndicies, newNode)...)

		return diffs, nil
	}

	// Old node doesn't exist in new node
	if oldNode != nil && newNode == nil {
		diffs = append(diffs, &diffInfo{
			Type:         gas.DiffType_DELETE,
			Root:         selector,
			PathIndicies: pathIndicies,
			OldNode:      oldNode,
		})

		return diffs, nil
	}

	// Not the same type, remove current node and replace with new
	if !diffTreeNodeTypeMatch(oldNode, newNode) {
		diffs = append(diffs, &diffInfo{
			Root:         selector,
			PathIndicies: pathIndicies,
			Type:         gas.DiffType_DELETE,
			OldNode:      oldNode,
		})

		diffs = append(diffs, diffCreate(selector, pathIndicies, newNode)...)

		return diffs, nil
	}

	switch v := oldNode.(type) {
	case *NodeGroup:
		oldList := v.Get()
		newNG, _ := newNode.(*NodeGroup)
		newList := newNG.group
		indexOffset := 0

		for i := 0; i < len(oldList); i++ {
			var n any
			if len(newList) > i {
				n = newList[i]
			}

			subDiffs, err := d.Trees(selector, append(pathIndicies, i-indexOffset), oldList[i], n)
			if err != nil {
				return nil, fmt.Errorf("diff NodeGroup: %w", err)
			}

			diffs = append(diffs, subDiffs...)
		}
		// Any new elements?
	case string:
		newStr, _ := newNode.(string)

		// content doesn't match, update content
		if v != newStr {
			diffs = append(diffs, &diffInfo{
				Root:         selector,
				PathIndicies: pathIndicies,
				Type:         gas.DiffType_UPDATE,
				Text:         &newStr,
			})
		}
	case HTML:
		newHTML, _ := newNode.(HTML)

		// content doesn't match, update content
		if v != newHTML {
			diffs = append(diffs, &diffInfo{
				Root:         selector,
				PathIndicies: pathIndicies,
				Type:         gas.DiffType_UPDATE,
				HTML:         &newHTML,
			})
		}
	case Tagger:
		newTag, _ := newNode.(Tagger)

		// Different tag?
		if v.GetName() != newTag.GetName() || v.IsVoid() != newTag.IsVoid() {
			diffs = append(diffs, &diffInfo{
				Root:         selector,
				PathIndicies: pathIndicies,
				Type:         gas.DiffType_UPDATE,
				Tag:          newTag,
			})

			return diffs, nil
		}

		// Attributes
		// The browser doesn't care about the order as we use setAttribute and removeAttribute.

		oldAttrs := v.GetAttributes()
		newAttrs := newTag.GetAttributes()

		// exists maps helps us know if we should delete or update
		oldAttrsMap := map[string]Attributer{}
		for i := 0; i < len(oldAttrs); i++ {
			oldAttrsMap[oldAttrs[i].GetName()] = oldAttrs[i]
		}

		newAttrsMap := map[string]Attributer{}

		for i := 0; i < len(newAttrs); i++ {
			newAttrsMap[newAttrs[i].GetName()] = newAttrs[i]
		}

		// Update existing or create new
		for i := 0; i < len(newAttrs); i++ {
			oldAttr, exits := oldAttrsMap[newAttrs[i].GetName()]

			if !exits || newAttrs[i].GetValue() != oldAttr.GetValue() {
				dt := gas.DiffType_UPDATE
				if !exits {
					dt = gas.DiffType_CREATE
				}

				diffs = append(diffs, &diffInfo{
					Root:         selector,
					PathIndicies: pathIndicies,
					Type:         dt,
					Attribute:    newAttrs[i].Clone(),
				})
			}
		}

		// Delete old attrs that have been removed
		for i := 0; i < len(oldAttrs); i++ {
			_, exits := newAttrsMap[oldAttrs[i].GetName()]
			if !exits {
				diffs = append(diffs, &diffInfo{
					Root:         selector,
					PathIndicies: pathIndicies,
					Type:         gas.DiffType_DELETE,
					Attribute:    oldAttrs[i].Clone(),
				})
			}
		}

		// Is this tag a component?
		// TODO: add tests to ensure this always works
		if attr, exits := newAttrsMap[AttrID]; exits {
			selector = attr.GetValue()
			pathIndicies = nil
		}

		oldKids := v.GetNodes().Get()
		newKids := newTag.GetNodes().Get()

		// Loop old kids
		i := 0
		for ; i < len(oldKids); i++ {
			var newKid any
			if i < len(newKids) {
				newKid = newKids[i]
			}

			kidDiffs, err := d.Trees(selector, append(pathIndicies, i), oldKids[i], newKid)
			if err != nil {
				return nil, fmt.Errorf("tag diff kids: %w", err)
			}

			// Reverse order delete batches
			// TODO: make tests
			if len(kidDiffs) > 1 {
				var (
					newKids     = make([]*diffInfo, 0, len(kidDiffs))
					deleteBatch []*diffInfo
				)

				for j := 0; j < len(kidDiffs); j++ {
					// Vars
					var (
						// this diff
						isEndOfLoop   = len(kidDiffs)-1 == j
						thisDiffIsDel = kidDiffs[j].Type == gas.DiffType_DELETE

						// Init for next diff
						nextPathGreater = false
						nextDiffIsDel   = false
						// other
						batchStarted = len(deleteBatch) != 0
					)
					// Next diff vars
					if !isEndOfLoop {
						nextPathGreater = pathGreaterLoop(kidDiffs[j+1].PathIndicies, kidDiffs[j].PathIndicies)
						nextDiffIsDel = kidDiffs[j+1].Type == gas.DiffType_DELETE
					}

					// Next in batch
					if batchStarted {
						deleteBatch = append(deleteBatch, kidDiffs[j])
						// Start of a batch?
					} else if thisDiffIsDel && nextDiffIsDel && nextPathGreater {
						deleteBatch = append(deleteBatch, kidDiffs[j])
					} else {
						newKids = append(newKids, kidDiffs[j])
					}

					// end of a batch?
					if batchStarted && (!nextDiffIsDel || !nextPathGreater) {
						// Reverse
						for k := len(deleteBatch) - 1; k >= 0; k-- {
							newKids = append(newKids, deleteBatch[k])
						}
						// Clear batch
						deleteBatch = nil
						// Add normally
					}
				}

				kidDiffs = newKids
			}

			diffs = append(diffs, kidDiffs...)
		}

		// Any extra new kids?
		for ; i < len(newKids); i++ {
			diffs = append(diffs, diffCreate(selector, append(pathIndicies, i), newKids[i])...)
		}
	}

	return diffs, nil
}

func diffCreate(compID string, pathIndicies []int, el any) (diffs []*diffInfo) {
	createDiff := &diffInfo{
		Type:         gas.DiffType_CREATE,
		Root:         compID,
		PathIndicies: pathIndicies,
	}

	switch v := el.(type) {
	case *NodeGroup:
		g := v.Get()
		for i := 0; i < len(g); i++ {
			diffs = append(diffs, diffCreate(compID, pathIndicies, g[i])...)
		}
		return diffs
	case string:
		createDiff.Text = &v
	case *HTML:
		createDiff.HTML = v
	case Tagger:
		createDiff.Tag = v
	case *Attribute:
		createDiff.Attribute = v
	case nil:
		return nil
	default:
		panic(fmt.Errorf("unexpected type: %#v", el))
	}

	return []*diffInfo{createDiff}
}

func diffTreeNodeTypeMatch(oldNode, newNode any) bool {
	switch oldNode.(type) {
	case *NodeGroup:
		_, ok := newNode.(*NodeGroup)
		return ok

	case string:
		_, ok := newNode.(string)
		return ok

	case *string:
		_, ok := newNode.(*string)
		return ok

	case *HTML:
		_, ok := newNode.(*HTML)
		return ok

	case Tagger:
		_, ok := newNode.(Tagger)
		return ok

	case nil:
		return false

	default:
		panic(fmt.Sprintf("unexpected type: %#v", oldNode))
	}
}

func pathGreaterLoop(pathA, pathB []int) bool {
	if len(pathA) != 0 && len(pathB) == 0 {
		return false
	}

	if len(pathA) == 0 {
		return false
	}

	if pathA[0] > pathB[0] {
		return true
	}

	if pathA[0] < pathB[0] {
		return false
	}

	// If we are here then this level of the path is equal, go to the next level
	return pathGreaterLoop(pathA[1:], pathB[1:])
}
