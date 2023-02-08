package tailpipe

import (
	"github.com/facette/natsort"
	"golang.org/x/exp/slices"
)

type StringMap map[string]string

func (s StringMap) Extend(extend ...StringMap) StringMap {
	for _, e := range extend {
		for k, v := range e {
			s[k] = v
		}
	}
	return s
}

type KV struct {
	Key   string
	Value string
}

func (s StringMap) SortedKV() []KV {
	arr := make([]KV, 0, len(s))
	for k, v := range s {
		arr = append(arr, KV{k, v})
	}
	slices.SortFunc(arr, func(a, b KV) bool {
		return natsort.Compare(a.Key, b.Key)
	})
	return arr
}

type StringSlice []string

func (s StringSlice) Extend(extend ...StringSlice) {
	for _, e := range extend {
		s = append(s, e...)
	}
}

type SStringMap map[string]StringMap

func (s SStringMap) Extend(extend ...SStringMap) SStringMap {
	for _, e := range extend {
		for k, v := range e {
			vv, ok := s[k]
			if !ok {
				vv = StringMap{}
				s[k] = vv
			}
			vv.Extend(v)
		}
	}
	return s
}

func (ss SStringMap) Lookup(key1, key2, defaultValue string) string {
	v, ok := ss[key1]
	if !ok {
		return defaultValue
	}
	vv, ok := v[key2]
	if !ok {
		return defaultValue
	}
	return vv
}

func (ss SStringMap) SortedKV() []KV {
	arr := []KV{}
	for kk, vv := range ss {
		for k, v := range vv {
			arr = append(arr, KV{kk + "-" + k, v})
		}
	}
	slices.SortFunc(arr, func(a, b KV) bool {
		return natsort.Compare(a.Key, b.Key)
	})
	return arr
}

type SSStringMap map[string]SStringMap

func (s SSStringMap) Extend(extend ...SSStringMap) SSStringMap {
	for _, e := range extend {
		for k, v := range e {
			vv, ok := s[k]
			if !ok {
				vv = SStringMap{}
				s[k] = vv
			}
			vv.Extend(v)
		}
	}
	return s
}

type StringSliceMap map[string]StringSlice

func (ss StringSliceMap) Extend(extend StringSliceMap) {
	for k, v := range extend {
		vv, ok := ss[k]
		if !ok {
			vv = StringSlice{}
			ss[k] = vv
		}
		vv.Extend(v)
	}
}

func (ss StringSliceMap) SortedKV() []KV {
	arr := []KV{}
	for k, v := range ss {
		for _, vv := range v {
			arr = append(arr, KV{k, vv})
		}
	}
	slices.SortFunc(arr, func(a, b KV) bool {
		return natsort.Compare(a.Key, b.Key)
	})
	return arr
}

type TailPipeConfig struct {
	Animation          StringMap
	Aria               StringMap
	AspectRatio        StringMap
	BackdropFilter     BackdropFilterConfig
	Background         BackgroundConfig
	Blur               StringMap
	Border             BorderConfig
	BoxShadow          BoxShadowConfig
	Brightness         StringMap
	CaretColor         SStringMap
	Columns            StringMap
	Contrast           StringMap
	Color              SStringMap
	Cursor             StringMap
	Divide             DivideConfig
	DropShadow         StringSliceMap
	Fill               SStringMap
	Flex               FlexConfig
	Font               FontConfig
	Gap                StringMap
	GradientColorStops SStringMap
	Grayscale          StringMap
	Grid               GridConfig
	Height             MinMaxValue
	HueRotate          StringMap
	Inset              StringMap
	Invert             StringMap
	Keyframes          SSStringMap
	LetterSpacing      StringMap
	LineHeight         StringMap
	ListStyleType      StringMap
	ListStylePosition  StringMap
	Margin             StringMap
	ObjectPosition     StringMap
	Opacity            StringMap
	Order              StringMap
	Outline            OutlineConfig
	Padding            StringMap
	Placeholder        PlaceholderConfig
	Ring               RingConfig
	Rotate             StringMap
	Saturate           StringMap
	Scale              StringMap
	Screens            StringMap
	Scroll             ScrollConfig
	Sepia              StringMap
	Skew               StringMap
	Space              StringMap
	Spacing            StringMap
	Stroke             SStringMap
	StrokeWidth        StringMap
	Text               TextConfig
	TransformOrigin    StringMap
	Transition         TransitionConfig
	Translate          StringMap
	Width              MinMaxValue
	WillChange         StringMap
	ZIndex             StringMap
}

type MinMaxValue struct {
	Value, Min, Max StringMap
}

type BackgroundConfig struct {
	Color                          SStringMap
	Image, Opacity, Position, Size StringMap
}

type BackdropFilterConfig struct {
	Blur, Brightness, Contrast, Grayscale, HueRotate, Invert, Opacity, Saturate, Sepia StringMap
}

type BorderConfig struct {
	Color                           SStringMap
	Opacity, Radius, Spacing, Width StringMap
}

type BoxShadowConfig struct {
	Color          SStringMap
	Value, Opacity StringMap
}

type DivideConfig struct {
	Color          SStringMap
	Opacity, Width StringMap
}

type FlexConfig struct {
	Value, Basis, Direction, Grow, Shrink, Wrap StringMap
}

type FontConfig struct {
	Family                                   StringSliceMap
	Size                                     FontSizeConfigMap
	Smoothing, Style, Weight, VariantNumeric StringMap
}

type FontSizeConfigMap map[string]FontSizeConfig

type FontSizeConfig struct {
	Value      string
	LineHeight string
}

type GridConfig struct {
	AutoColumns     StringMap
	AutoRows        StringMap
	Columns         GridValueConfig
	Rows            GridValueConfig
	TemplateColumns StringMap
	TemplateRows    StringMap
}

type GridValueConfig struct {
	Value StringMap
	Start StringMap
	End   StringMap
}

type PlaceholderConfig struct {
	Color   SStringMap
	Opacity StringMap
}

type OutlineConfig struct {
	Color         SStringMap
	Offset, Width StringMap
}

type RingConfig struct {
	Color, OffsetColor          SStringMap
	OffsetWidth, Opacity, Width StringMap
}

type ScrollConfig struct {
	Margin  StringMap
	Padding StringMap
}

type TextConfig struct {
	Color, DecorationColor                  SStringMap
	Decoration, DecorationThickness         StringMap
	Align, Indent, Opacity, UnderlineOffset StringMap
}

type TransitionConfig struct {
	Delay, Duration, Property, TimingFunction StringMap
}

func defaultBaseTailPipeConfig() TailPipeConfig {
	return TailPipeConfig{
		Animation: StringMap{
			"none":   "none",
			"spin":   "spin 1s linear infinite",
			"ping":   "ping 1s cubic-bezier(0, 0, 0.2, 1) infinite",
			"pulse":  "pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite",
			"bounce": "bounce 1s infinite",
		},
		Aria: StringMap{
			"checked":  `checked="true"`,
			"disabled": `disabled="true"`,
			"expanded": `expanded="true"`,
			"hidden":   `hidden="true"`,
			"pressed":  `pressed="true"`,
			"readonly": `readonly="true"`,
			"required": `required="true"`,
			"selected": `selected="true"`,
		},
		AspectRatio: StringMap{
			"auto":   "auto",
			"square": "1 / 1",
			"video":  "16 / 9",
		},
		BackdropFilter: BackdropFilterConfig{},
		Background: BackgroundConfig{
			Color: SStringMap{},
			Image: StringMap{
				"none":           "none",
				"gradient-to-t":  "linear-gradient(to top, var(--tw-gradient-stops))",
				"gradient-to-tr": "linear-gradient(to top right, var(--tw-gradient-stops))",
				"gradient-to-r":  "linear-gradient(to right, var(--tw-gradient-stops))",
				"gradient-to-br": "linear-gradient(to bottom right, var(--tw-gradient-stops))",
				"gradient-to-b":  "linear-gradient(to bottom, var(--tw-gradient-stops))",
				"gradient-to-bl": "linear-gradient(to bottom left, var(--tw-gradient-stops))",
				"gradient-to-l":  "linear-gradient(to left, var(--tw-gradient-stops))",
				"gradient-to-tl": "linear-gradient(to top left, var(--tw-gradient-stops))",
			},
			Opacity: StringMap{},
			Position: StringMap{
				"bottom":       "bottom",
				"center":       "center",
				"left":         "left",
				"left-bottom":  "left bottom",
				"left-top":     "left top",
				"right":        "right",
				"right-bottom": "right bottom",
				"right-top":    "right top",
				"top":          "top",
			},
			Size: StringMap{
				"auto":    "auto",
				"cover":   "cover",
				"contain": "contain",
			},
		},
		Blur: StringMap{
			"0":       "0",
			"none":    "0",
			"sm":      "4px",
			"DEFAULT": "8px",
			"md":      "12px",
			"lg":      "16px",
			"xl":      "24px",
			"2xl":     "40px",
			"3xl":     "64px",
		},
		Border: BorderConfig{
			Color:   SStringMap{},
			Opacity: StringMap{},
			Radius: StringMap{
				"none":    "0px",
				"sm":      "0.125rem",
				"DEFAULT": "0.25rem",
				"md":      "0.375rem",
				"lg":      "0.5rem",
				"xl":      "0.75rem",
				"2xl":     "1rem",
				"3xl":     "1.5rem",
				"full":    "9999px",
			},
			Spacing: StringMap{},
			Width: StringMap{
				"DEFAULT": "1px",
				"0":       "0px",
				"2":       "2px",
				"4":       "4px",
				"8":       "8px",
			},
		},
		BoxShadow: BoxShadowConfig{
			Value: StringMap{
				"sm":      "0 1px 2px 0 rgb(0 0 0 / 0.05)",
				"DEFAULT": "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)",
				"md":      "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)",
				"lg":      "0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)",
				"xl":      "0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)",
				"2xl":     "0 25px 50px -12px rgb(0 0 0 / 0.25)",
				"inner":   "inset 0 2px 4px 0 rgb(0 0 0 / 0.05)",
				"none":    "none",
			},
			Color:   SStringMap{},
			Opacity: StringMap{},
		},
		Brightness: StringMap{
			"0":   "0",
			"50":  ".5",
			"75":  ".75",
			"90":  ".9",
			"95":  ".95",
			"100": "1",
			"105": "1.05",
			"110": "1.1",
			"125": "1.25",
			"150": "1.5",
			"200": "2",
		},
		CaretColor: SStringMap{},
		Color: SStringMap{
			"inherit": {
				"DEFAULT": "inherit",
			},
			"current": {
				"DEFAULT": "currentColor",
			},
			"transparent": {
				"DEFAULT": "transparent",
			},
			"black": {
				"DEFAULT": "#000",
			},
			"white": {
				"DEFAULT": "#fff",
			},
			"slate": {
				"50":  "#f8fafc",
				"100": "#f1f5f9",
				"200": "#e2e8f0",
				"300": "#cbd5e1",
				"400": "#94a3b8",
				"500": "#64748b",
				"600": "#475569",
				"700": "#334155",
				"800": "#1e293b",
				"900": "#0f172a",
			},
			"gray": {

				"50":  "#f9fafb",
				"100": "#f3f4f6",
				"200": "#e5e7eb",
				"300": "#d1d5db",
				"400": "#9ca3af",
				"500": "#6b7280",
				"600": "#4b5563",
				"700": "#374151",
				"800": "#1f2937",
				"900": "#111827",
			},
			"zinc": {

				"50":  "#fafafa",
				"100": "#f4f4f5",
				"200": "#e4e4e7",
				"300": "#d4d4d8",
				"400": "#a1a1aa",
				"500": "#71717a",
				"600": "#52525b",
				"700": "#3f3f46",
				"800": "#27272a",
				"900": "#18181b",
			},
			"neutral": {

				"50":  "#fafafa",
				"100": "#f5f5f5",
				"200": "#e5e5e5",
				"300": "#d4d4d4",
				"400": "#a3a3a3",
				"500": "#737373",
				"600": "#525252",
				"700": "#404040",
				"800": "#262626",
				"900": "#171717",
			},
			"stone": {

				"50":  "#fafaf9",
				"100": "#f5f5f4",
				"200": "#e7e5e4",
				"300": "#d6d3d1",
				"400": "#a8a29e",
				"500": "#78716c",
				"600": "#57534e",
				"700": "#44403c",
				"800": "#292524",
				"900": "#1c1917",
			},
			"red": {

				"50":  "#fef2f2",
				"100": "#fee2e2",
				"200": "#fecaca",
				"300": "#fca5a5",
				"400": "#f87171",
				"500": "#ef4444",
				"600": "#dc2626",
				"700": "#b91c1c",
				"800": "#991b1b",
				"900": "#7f1d1d",
			},
			"orange": {

				"50":  "#fff7ed",
				"100": "#ffedd5",
				"200": "#fed7aa",
				"300": "#fdba74",
				"400": "#fb923c",
				"500": "#f97316",
				"600": "#ea580c",
				"700": "#c2410c",
				"800": "#9a3412",
				"900": "#7c2d12",
			},
			"amber": {

				"50":  "#fffbeb",
				"100": "#fef3c7",
				"200": "#fde68a",
				"300": "#fcd34d",
				"400": "#fbbf24",
				"500": "#f59e0b",
				"600": "#d97706",
				"700": "#b45309",
				"800": "#92400e",
				"900": "#78350f",
			},
			"yellow": {

				"50":  "#fefce8",
				"100": "#fef9c3",
				"200": "#fef08a",
				"300": "#fde047",
				"400": "#facc15",
				"500": "#eab308",
				"600": "#ca8a04",
				"700": "#a16207",
				"800": "#854d0e",
				"900": "#713f12",
			},
			"lime": {
				"50":  "#f7fee7",
				"100": "#ecfccb",
				"200": "#d9f99d",
				"300": "#bef264",
				"400": "#a3e635",
				"500": "#84cc16",
				"600": "#65a30d",
				"700": "#4d7c0f",
				"800": "#3f6212",
				"900": "#365314",
			},
			"green": {
				"50":  "#f0fdf4",
				"100": "#dcfce7",
				"200": "#bbf7d0",
				"300": "#86efac",
				"400": "#4ade80",
				"500": "#22c55e",
				"600": "#16a34a",
				"700": "#15803d",
				"800": "#166534",
				"900": "#14532d",
			},
			"emerald": {
				"50":  "#ecfdf5",
				"100": "#d1fae5",
				"200": "#a7f3d0",
				"300": "#6ee7b7",
				"400": "#34d399",
				"500": "#10b981",
				"600": "#059669",
				"700": "#047857",
				"800": "#065f46",
				"900": "#064e3b",
			},
			"teal": {
				"50":  "#f0fdfa",
				"100": "#ccfbf1",
				"200": "#99f6e4",
				"300": "#5eead4",
				"400": "#2dd4bf",
				"500": "#14b8a6",
				"600": "#0d9488",
				"700": "#0f766e",
				"800": "#115e59",
				"900": "#134e4a",
			},
			"cyan": {
				"50":  "#ecfeff",
				"100": "#cffafe",
				"200": "#a5f3fc",
				"300": "#67e8f9",
				"400": "#22d3ee",
				"500": "#06b6d4",
				"600": "#0891b2",
				"700": "#0e7490",
				"800": "#155e75",
				"900": "#164e63",
			},
			"sky": {
				"50":  "#f0f9ff",
				"100": "#e0f2fe",
				"200": "#bae6fd",
				"300": "#7dd3fc",
				"400": "#38bdf8",
				"500": "#0ea5e9",
				"600": "#0284c7",
				"700": "#0369a1",
				"800": "#075985",
				"900": "#0c4a6e",
			},
			"blue": {
				"50":  "#eff6ff",
				"100": "#dbeafe",
				"200": "#bfdbfe",
				"300": "#93c5fd",
				"400": "#60a5fa",
				"500": "#3b82f6",
				"600": "#2563eb",
				"700": "#1d4ed8",
				"800": "#1e40af",
				"900": "#1e3a8a",
			},
			"indigo": {
				"50":  "#eef2ff",
				"100": "#e0e7ff",
				"200": "#c7d2fe",
				"300": "#a5b4fc",
				"400": "#818cf8",
				"500": "#6366f1",
				"600": "#4f46e5",
				"700": "#4338ca",
				"800": "#3730a3",
				"900": "#312e81",
			},
			"violet": {
				"50":  "#f5f3ff",
				"100": "#ede9fe",
				"200": "#ddd6fe",
				"300": "#c4b5fd",
				"400": "#a78bfa",
				"500": "#8b5cf6",
				"600": "#7c3aed",
				"700": "#6d28d9",
				"800": "#5b21b6",
				"900": "#4c1d95",
			},
			"purple": {
				"50":  "#faf5ff",
				"100": "#f3e8ff",
				"200": "#e9d5ff",
				"300": "#d8b4fe",
				"400": "#c084fc",
				"500": "#a855f7",
				"600": "#9333ea",
				"700": "#7e22ce",
				"800": "#6b21a8",
				"900": "#581c87",
			},
			"fuchsia": {
				"50":  "#fdf4ff",
				"100": "#fae8ff",
				"200": "#f5d0fe",
				"300": "#f0abfc",
				"400": "#e879f9",
				"500": "#d946ef",
				"600": "#c026d3",
				"700": "#a21caf",
				"800": "#86198f",
				"900": "#701a75",
			},
			"pink": {
				"50":  "#fdf2f8",
				"100": "#fce7f3",
				"200": "#fbcfe8",
				"300": "#f9a8d4",
				"400": "#f472b6",
				"500": "#ec4899",
				"600": "#db2777",
				"700": "#be185d",
				"800": "#9d174d",
				"900": "#831843",
			},
			"rose": {
				"50":  "#fff1f2",
				"100": "#ffe4e6",
				"200": "#fecdd3",
				"300": "#fda4af",
				"400": "#fb7185",
				"500": "#f43f5e",
				"600": "#e11d48",
				"700": "#be123c",
				"800": "#9f1239",
				"900": "#881337",
			},
		},
		Columns: StringMap{
			"auto": "auto",
			"1":    "1",
			"2":    "2",
			"3":    "3",
			"4":    "4",
			"5":    "5",
			"6":    "6",
			"7":    "7",
			"8":    "8",
			"9":    "9",
			"10":   "10",
			"11":   "11",
			"12":   "12",
			"3xs":  "16rem",
			"2xs":  "18rem",
			"xs":   "20rem",
			"sm":   "24rem",
			"md":   "28rem",
			"lg":   "32rem",
			"xl":   "36rem",
			"2xl":  "42rem",
			"3xl":  "48rem",
			"4xl":  "56rem",
			"5xl":  "64rem",
			"6xl":  "72rem",
			"7xl":  "80rem",
		},
		Contrast: StringMap{
			"0":   "0",
			"50":  ".5",
			"75":  ".75",
			"100": "1",
			"125": "1.25",
			"150": "1.5",
			"200": "2",
		},
		Cursor: StringMap{
			"auto":          "auto",
			"default":       "default",
			"pointer":       "pointer",
			"wait":          "wait",
			"text":          "text",
			"move":          "move",
			"help":          "help",
			"not-allowed":   "not-allowed",
			"none":          "none",
			"context-menu":  "context-menu",
			"progress":      "progress",
			"cell":          "cell",
			"crosshair":     "crosshair",
			"vertical-text": "vertical-text",
			"alias":         "alias",
			"copy":          "copy",
			"no-drop":       "no-drop",
			"grab":          "grab",
			"grabbing":      "grabbing",
			"all-scroll":    "all-scroll",
			"col-resize":    "col-resize",
			"row-resize":    "row-resize",
			"n-resize":      "n-resize",
			"e-resize":      "e-resize",
			"s-resize":      "s-resize",
			"w-resize":      "w-resize",
			"ne-resize":     "ne-resize",
			"nw-resize":     "nw-resize",
			"se-resize":     "se-resize",
			"sw-resize":     "sw-resize",
			"ew-resize":     "ew-resize",
			"ns-resize":     "ns-resize",
			"nesw-resize":   "nesw-resize",
			"nwse-resize":   "nwse-resize",
			"zoom-in":       "zoom-in",
			"zoom-out":      "zoom-out",
		},
		Divide: DivideConfig{
			Color:   SStringMap{},
			Opacity: StringMap{},
			Width:   StringMap{},
		},
		DropShadow: StringSliceMap{
			"sm":      StringSlice{"0 1px 1px rgb(0 0 0 / 0.05)"},
			"DEFAULT": StringSlice{"0 1px 2px rgb(0 0 0 / 0.1)", "0 1px 1px rgb(0 0 0 / 0.06)"},
			"md":      StringSlice{"0 4px 3px rgb(0 0 0 / 0.07)", "0 2px 2px rgb(0 0 0 / 0.06)"},
			"lg":      StringSlice{"0 10px 8px rgb(0 0 0 / 0.04)", "0 4px 3px rgb(0 0 0 / 0.1)"},
			"xl":      StringSlice{"0 20px 13px rgb(0 0 0 / 0.03)", "0 8px 5px rgb(0 0 0 / 0.08)"},
			"2xl":     StringSlice{"0 25px 25px rgb(0 0 0 / 0.15)"},
			"none":    StringSlice{"0 0 #0000"},
		},
		Fill: SStringMap{},
		Flex: FlexConfig{
			Direction: StringMap{
				"row":         "row",
				"row-reverse": "row-reverse",
				"col":         "column",
				"col-reverse": "column-reverse",
			},
			Wrap: StringMap{
				"wrap":         "wrap",
				"wrap-reverse": "wrap-reverse",
				"nowrap":       "nowrap",
			},
			Value: StringMap{
				"1":       "1 1 0%",
				"auto":    "1 1 auto",
				"initial": "0 1 auto",
				"none":    "none",
			},
			Basis: StringMap{},
			Grow: StringMap{
				"0":       "0",
				"DEFAULT": "1",
			},
			Shrink: StringMap{
				"0":       "0",
				"DEFAULT": "1",
			},
		},
		Font: FontConfig{
			Family: StringSliceMap{
				"sans": {
					"ui-sans-serif",
					"system-ui",
					"-apple-system",
					"BlinkMacSystemFont",
					"Segoe UI",
					"Roboto",
					"Helvetica Neue",
					"Arial",
					"Noto Sans",
					"sans-serif",
					"Apple Color Emoji",
					"Segoe UI Emoji",
					"Segoe UI Symbol",
					"Noto Color Emoji",
				},
				"serif": {"ui-serif", "Georgia", "Cambria", "Times New Roman", "Times", "serif"},
				"mono": {
					"ui-monospace",
					"SFMono-Regular",
					"Menlo",
					"Monaco",
					"Consolas",
					"Liberation Mono",
					"Courier New",
					"monospace",
				},
			},
			Size: FontSizeConfigMap{
				"xs": {
					Value:      "0.75rem",
					LineHeight: "1rem",
				},
				"sm": {
					Value:      "0.875rem",
					LineHeight: "1.25rem",
				},
				"base": {
					Value:      "1rem",
					LineHeight: "1.5rem",
				},
				"lg": {
					Value:      "1.125rem",
					LineHeight: "1.75rem",
				},
				"xl": {
					Value:      "1.25rem",
					LineHeight: "1.75rem",
				},
				"2xl": {
					Value:      "1.5rem",
					LineHeight: "2rem",
				},
				"3xl": {
					Value:      "1.875rem",
					LineHeight: "2.25rem",
				},
				"4xl": {
					Value:      "2.25rem",
					LineHeight: "2.5rem",
				},
				"5xl": {
					Value:      "3rem",
					LineHeight: "1",
				},
				"6xl": {
					Value:      "3.75rem",
					LineHeight: "1",
				},
				"7xl": {
					Value:      "4.5rem",
					LineHeight: "1",
				},
				"8xl": {
					Value:      "6rem",
					LineHeight: "1",
				},
				"9xl": {
					Value:      "8rem",
					LineHeight: "1",
				},
			},
			Smoothing: StringMap{
				"antialiased":          "antialiased",
				"subpixel-antialiased": "subpixel-antialiased",
			},
			Style: StringMap{
				"italic":     "italic",
				"not-italic": "normal",
			},
			Weight: StringMap{
				"thin":       "100",
				"extralight": "200",
				"light":      "300",
				"normal":     "400",
				"medium":     "500",
				"semibold":   "600",
				"bold":       "700",
				"extrabold":  "800",
				"black":      "900",
			},
			VariantNumeric: StringMap{
				"normal-nums":        "normal-nums",
				"ordinal":            "ordinal",
				"slashed-zero":       "slashed-zero",
				"lining-nums":        "lining-nums",
				"oldstyle-nums":      "oldstyle-nums",
				"proportional-nums":  "proportional-nums",
				"tabular-nums":       "tabular-nums",
				"diagonal-fractions": "diagonal-fractions",
				"stacked-fractions":  "stacked-fractions",
			},
		},
		Gap:                StringMap{},
		GradientColorStops: SStringMap{},
		Grayscale: StringMap{
			"0":       "0",
			"DEFAULT": "100%",
		},
		Grid: GridConfig{
			AutoColumns: StringMap{
				"auto": "auto",
				"min":  "min-content",
				"max":  "max-content",
				"fr":   "minmax(0, 1fr)",
			},
			AutoRows: StringMap{
				"auto": "auto",
				"min":  "min-content",
				"max":  "max-content",
				"fr":   "minmax(0, 1fr)",
			},
			Columns: GridValueConfig{
				Value: StringMap{
					"auto":      "auto",
					"span-1":    "span 1 / span 1",
					"span-2":    "span 2 / span 2",
					"span-3":    "span 3 / span 3",
					"span-4":    "span 4 / span 4",
					"span-5":    "span 5 / span 5",
					"span-6":    "span 6 / span 6",
					"span-7":    "span 7 / span 7",
					"span-8":    "span 8 / span 8",
					"span-9":    "span 9 / span 9",
					"span-10":   "span 10 / span 10",
					"span-11":   "span 11 / span 11",
					"span-12":   "span 12 / span 12",
					"span-full": "1 / -1",
				},
				Start: StringMap{
					"auto": "auto",
					"1":    "1",
					"2":    "2",
					"3":    "3",
					"4":    "4",
					"5":    "5",
					"6":    "6",
					"7":    "7",
					"8":    "8",
					"9":    "9",
					"10":   "10",
					"11":   "11",
					"12":   "12",
					"13":   "13",
				},
				End: StringMap{
					"auto": "auto",
					"1":    "1",
					"2":    "2",
					"3":    "3",
					"4":    "4",
					"5":    "5",
					"6":    "6",
					"7":    "7",
					"8":    "8",
					"9":    "9",
					"10":   "10",
					"11":   "11",
					"12":   "12",
					"13":   "13",
				},
			},
			Rows: GridValueConfig{
				Value: StringMap{
					"auto":      "auto",
					"span-1":    "span 1 / span 1",
					"span-2":    "span 2 / span 2",
					"span-3":    "span 3 / span 3",
					"span-4":    "span 4 / span 4",
					"span-5":    "span 5 / span 5",
					"span-6":    "span 6 / span 6",
					"span-full": "1 / -1",
				},
				Start: StringMap{
					"auto": "auto",
					"1":    "1",
					"2":    "2",
					"3":    "3",
					"4":    "4",
					"5":    "5",
					"6":    "6",
					"7":    "7",
				},
				End: StringMap{
					"auto": "auto",
					"1":    "1",
					"2":    "2",
					"3":    "3",
					"4":    "4",
					"5":    "5",
					"6":    "6",
					"7":    "7",
				},
			},
			TemplateColumns: StringMap{
				"none": "none",
				"1":    "repeat(1, minmax(0, 1fr))",
				"2":    "repeat(2, minmax(0, 1fr))",
				"3":    "repeat(3, minmax(0, 1fr))",
				"4":    "repeat(4, minmax(0, 1fr))",
				"5":    "repeat(5, minmax(0, 1fr))",
				"6":    "repeat(6, minmax(0, 1fr))",
				"7":    "repeat(7, minmax(0, 1fr))",
				"8":    "repeat(8, minmax(0, 1fr))",
				"9":    "repeat(9, minmax(0, 1fr))",
				"10":   "repeat(10, minmax(0, 1fr))",
				"11":   "repeat(11, minmax(0, 1fr))",
				"12":   "repeat(12, minmax(0, 1fr))",
			},
			TemplateRows: StringMap{
				"none": "none",
				"1":    "repeat(1, minmax(0, 1fr))",
				"2":    "repeat(2, minmax(0, 1fr))",
				"3":    "repeat(3, minmax(0, 1fr))",
				"4":    "repeat(4, minmax(0, 1fr))",
				"5":    "repeat(5, minmax(0, 1fr))",
				"6":    "repeat(6, minmax(0, 1fr))",
			},
		},
		Height: MinMaxValue{
			Value: StringMap{},
			Min: StringMap{
				"0":      "0px",
				"full":   "100%",
				"screen": "100vh",
				"min":    "min-content",
				"max":    "max-content",
				"fit":    "fit-content",
			},
			Max: StringMap{},
		},
		HueRotate: StringMap{
			"0":   "0deg",
			"15":  "15deg",
			"30":  "30deg",
			"60":  "60deg",
			"90":  "90deg",
			"180": "180deg",
		},
		Inset: StringMap{},
		Invert: StringMap{
			"0":       "0",
			"DEFAULT": "100%",
		},
		Keyframes: SSStringMap{
			"spin": SStringMap{
				"to": StringMap{
					"transform": "rotate(360deg)",
				},
			},
			"ping": {
				"75%, 100%": {
					"transform": "scale(2)",
					"opacity":   "0",
				},
			},
			"pulse": {
				"50%": {
					"opacity": ".5",
				},
			},
			"bounce": {
				"0%, 100%": {
					"transform":               "translateY(-25%)",
					"animationTimingFunction": "cubic-bezier(0.8,0,1,1)",
				},
				"50%": {
					"transform":               "none",
					"animationTimingFunction": "cubic-bezier(0,0,0.2,1)",
				},
			},
		},
		LetterSpacing: StringMap{
			"tighter": "-0.05em",
			"tight":   "-0.025em",
			"normal":  "0em",
			"wide":    "0.025em",
			"wider":   "0.05em",
			"widest":  "0.1em",
		},
		LineHeight: StringMap{
			"none":    "1",
			"tight":   "1.25",
			"snug":    "1.375",
			"normal":  "1.5",
			"relaxed": "1.625",
			"loose":   "2",
			"3":       ".75rem",
			"4":       "1rem",
			"5":       "1.25rem",
			"6":       "1.5rem",
			"7":       "1.75rem",
			"8":       "2rem",
			"9":       "2.25rem",
			"10":      "2.5rem",
		},
		ListStyleType: StringMap{
			"none":    "none",
			"disc":    "disc",
			"decimal": "decimal",
		},
		ListStylePosition: StringMap{
			"inside":  "inside",
			"outside": "outside",
		},
		Margin: StringMap{},
		ObjectPosition: StringMap{
			"bottom":       "bottom",
			"center":       "center",
			"left":         "left",
			"left-bottom":  "left bottom",
			"left-top":     "left top",
			"right":        "right",
			"right-bottom": "right bottom",
			"right-top":    "right top",
			"top":          "top",
		},
		Opacity: StringMap{
			"0":   "0",
			"5":   "0.05",
			"10":  "0.1",
			"20":  "0.2",
			"25":  "0.25",
			"30":  "0.3",
			"40":  "0.4",
			"50":  "0.5",
			"60":  "0.6",
			"70":  "0.7",
			"75":  "0.75",
			"80":  "0.8",
			"90":  "0.9",
			"95":  "0.95",
			"100": "1",
		},
		Order: StringMap{
			"first": "-9999",
			"last":  "9999",
			"none":  "0",
			"1":     "1",
			"2":     "2",
			"3":     "3",
			"4":     "4",
			"5":     "5",
			"6":     "6",
			"7":     "7",
			"8":     "8",
			"9":     "9",
			"10":    "10",
			"11":    "11",
			"12":    "12",
		},
		Outline: OutlineConfig{
			Color: SStringMap{},
			Offset: StringMap{
				"0": "0px",
				"1": "1px",
				"2": "2px",
				"4": "4px",
				"8": "8px",
			},
			Width: StringMap{
				"0": "0px",
				"1": "1px",
				"2": "2px",
				"4": "4px",
				"8": "8px",
			},
		},
		Padding: StringMap{},
		Placeholder: PlaceholderConfig{
			Color:   SStringMap{},
			Opacity: StringMap{},
		},

		Ring: RingConfig{
			Color:       SStringMap{},
			OffsetColor: SStringMap{},
			OffsetWidth: StringMap{
				"0": "0px",
				"1": "1px",
				"2": "2px",
				"4": "4px",
				"8": "8px",
			},
			Opacity: StringMap{},
			Width: StringMap{
				"DEFAULT": "3px",
				"0":       "0px",
				"1":       "1px",
				"2":       "2px",
				"4":       "4px",
				"8":       "8px",
			},
		},
		Rotate: StringMap{
			"0":   "0deg",
			"1":   "1deg",
			"2":   "2deg",
			"3":   "3deg",
			"6":   "6deg",
			"12":  "12deg",
			"45":  "45deg",
			"90":  "90deg",
			"180": "180deg",
		},
		Saturate: StringMap{
			"0":   "0",
			"50":  ".5",
			"100": "1",
			"150": "1.5",
			"200": "2",
		},
		Scale: StringMap{
			"0":   "0",
			"50":  ".5",
			"75":  ".75",
			"90":  ".9",
			"95":  ".95",
			"100": "1",
			"105": "1.05",
			"110": "1.1",
			"125": "1.25",
			"150": "1.5",
		},
		Screens: StringMap{
			"sm":  "640px",
			"md":  "768px",
			"lg":  "1024px",
			"xl":  "1280px",
			"2xl": "1536px",
		},
		Scroll: ScrollConfig{
			Margin:  StringMap{},
			Padding: StringMap{},
		},
		Sepia: StringMap{
			"0":       "0",
			"DEFAULT": "100%",
		},
		Skew: StringMap{
			"0":  "0deg",
			"1":  "1deg",
			"2":  "2deg",
			"3":  "3deg",
			"6":  "6deg",
			"12": "12deg",
		},
		Space: StringMap{},
		Spacing: StringMap{
			"px":  "1px",
			"0":   "0px",
			"0.5": "0.125rem",
			"1":   "0.25rem",
			"1.5": "0.375rem",
			"2":   "0.5rem",
			"2.5": "0.625rem",
			"3":   "0.75rem",
			"3.5": "0.875rem",
			"4":   "1rem",
			"5":   "1.25rem",
			"6":   "1.5rem",
			"7":   "1.75rem",
			"8":   "2rem",
			"9":   "2.25rem",
			"10":  "2.5rem",
			"11":  "2.75rem",
			"12":  "3rem",
			"14":  "3.5rem",
			"16":  "4rem",
			"20":  "5rem",
			"24":  "6rem",
			"28":  "7rem",
			"32":  "8rem",
			"36":  "9rem",
			"40":  "10rem",
			"44":  "11rem",
			"48":  "12rem",
			"52":  "13rem",
			"56":  "14rem",
			"60":  "15rem",
			"64":  "16rem",
			"72":  "18rem",
			"80":  "20rem",
			"96":  "24rem",
		},
		Stroke: SStringMap{},
		StrokeWidth: StringMap{
			"0": "0",
			"1": "1",
			"2": "2",
		},
		Text: TextConfig{
			Align: StringMap{
				"left":    "left",
				"center":  "center",
				"right":   "right",
				"justify": "justify",
				"start":   "start",
				"end":     "end",
			},
			Color: SStringMap{},
			Decoration: StringMap{
				"underline":    "underline",
				"overline":     "overline",
				"line-through": "line-through",
				"no-underline": "none",
			},
			DecorationColor: SStringMap{},
			DecorationThickness: StringMap{
				"auto":      "auto",
				"from-font": "from-font",
				"0":         "0px",
				"1":         "1px",
				"2":         "2px",
				"4":         "4px",
				"8":         "8px",
			},
			Indent:  StringMap{},
			Opacity: StringMap{},
			UnderlineOffset: StringMap{
				"auto": "auto",
				"0":    "0px",
				"1":    "1px",
				"2":    "2px",
				"4":    "4px",
				"8":    "8px",
			},
		},
		TransformOrigin: StringMap{
			"center":       "center",
			"top":          "top",
			"top-right":    "top right",
			"right":        "right",
			"bottom-right": "bottom right",
			"bottom":       "bottom",
			"bottom-left":  "bottom left",
			"left":         "left",
			"top-left":     "top left",
		},
		Transition: TransitionConfig{
			Delay: StringMap{
				"0":    "0s",
				"75":   "75ms",
				"100":  "100ms",
				"150":  "150ms",
				"200":  "200ms",
				"300":  "300ms",
				"500":  "500ms",
				"700":  "700ms",
				"1000": "1000ms",
			},
			Duration: StringMap{
				"DEFAULT": "150ms",
				"0":       "0s",
				"75":      "75ms",
				"100":     "100ms",
				"150":     "150ms",
				"200":     "200ms",
				"300":     "300ms",
				"500":     "500ms",
				"700":     "700ms",
				"1000":    "1000ms",
			},
			Property: StringMap{
				"none":      "none",
				"all":       "all",
				"DEFAULT":   "color, background-color, border-color, outline-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter",
				"Color":     "color, background-color, border-color, outline-color, text-decoration-color, fill, stroke",
				"opacity":   "opacity",
				"shadow":    "box-shadow",
				"transform": "transform",
			},
			TimingFunction: StringMap{
				"DEFAULT": "cubic-bezier(0.4, 0, 0.2, 1)",
				"linear":  "linear",
				"in":      "cubic-bezier(0.4, 0, 1, 1)",
				"out":     "cubic-bezier(0, 0, 0.2, 1)",
				"in-out":  "cubic-bezier(0.4, 0, 0.2, 1)",
			},
		},
		Translate: StringMap{},
		Width: MinMaxValue{
			Value: StringMap{},
			Min: StringMap{
				"0":    "0px",
				"full": "100%",
				"min":  "min-content",
				"max":  "max-content",
				"fit":  "fit-content",
			},
			Max: StringMap{},
		},
		WillChange: StringMap{
			"auto":      "auto",
			"scroll":    "scroll-position",
			"contents":  "contents",
			"transform": "transform",
		},
		ZIndex: StringMap{
			"auto": "auto",
			"0":    "0",
			"10":   "10",
			"20":   "20",
			"30":   "30",
			"40":   "40",
			"50":   "50",
		},
	}
}

func DefaultTailPipeConfig() TailPipeConfig {
	cfg := defaultBaseTailPipeConfig()

	cfg.BackdropFilter = BackdropFilterConfig{
		Blur:       cfg.Blur,
		Brightness: cfg.Brightness,
		Contrast:   cfg.Contrast,
		Grayscale:  cfg.Grayscale,
		HueRotate:  cfg.HueRotate,
		Invert:     cfg.Invert,
		Opacity:    cfg.Opacity,
		Saturate:   cfg.Saturate,
		Sepia:      cfg.Sepia,
	}

	cfg.Background.Color = cfg.Color
	cfg.Background.Opacity = cfg.Opacity

	cfg.Border.Color = SStringMap{
		"DEFAULT": {
			"DEFAULT": cfg.Color.Lookup("gray", "200", "currentColor"),
		},
	}.Extend(cfg.Color)
	cfg.Border.Opacity = cfg.Opacity
	cfg.Border.Spacing = cfg.Spacing

	cfg.BoxShadow.Color = cfg.Color
	cfg.BoxShadow.Opacity = cfg.Opacity

	cfg.CaretColor = cfg.Color

	cfg.Divide.Color = cfg.Border.Color
	cfg.Divide.Opacity = cfg.Border.Opacity
	cfg.Divide.Width = cfg.Border.Width

	cfg.Fill = SStringMap{
		"none": {
			"DEFAULT": "none",
		},
	}.Extend(cfg.Color)

	cfg.Flex.Basis = StringMap{
		"auto":  "auto",
		"1/2":   "50%",
		"1/3":   "33.333333%",
		"2/3":   "66.666667%",
		"1/4":   "25%",
		"2/4":   "50%",
		"3/4":   "75%",
		"1/5":   "20%",
		"2/5":   "40%",
		"3/5":   "60%",
		"4/5":   "80%",
		"1/6":   "16.666667%",
		"2/6":   "33.333333%",
		"3/6":   "50%",
		"4/6":   "66.666667%",
		"5/6":   "83.333333%",
		"1/12":  "8.333333%",
		"2/12":  "16.666667%",
		"3/12":  "25%",
		"4/12":  "33.333333%",
		"5/12":  "41.666667%",
		"6/12":  "50%",
		"7/12":  "58.333333%",
		"8/12":  "66.666667%",
		"9/12":  "75%",
		"10/12": "83.333333%",
		"11/12": "91.666667%",
		"full":  "100%",
	}.Extend(cfg.Spacing)

	cfg.Gap = cfg.Spacing

	cfg.GradientColorStops = cfg.Color

	cfg.Height.Value = StringMap{
		"auto":   "auto",
		"1/2":    "50%",
		"1/3":    "33.333333%",
		"2/3":    "66.666667%",
		"1/4":    "25%",
		"2/4":    "50%",
		"3/4":    "75%",
		"1/5":    "20%",
		"2/5":    "40%",
		"3/5":    "60%",
		"4/5":    "80%",
		"1/6":    "16.666667%",
		"2/6":    "33.333333%",
		"3/6":    "50%",
		"4/6":    "66.666667%",
		"5/6":    "83.333333%",
		"full":   "100%",
		"screen": "100vh",
		"min":    "min-content",
		"max":    "max-content",
		"fit":    "fit-content",
	}.Extend(cfg.Spacing)
	cfg.Height.Max = StringMap{
		"none":   "none",
		"full":   "100%",
		"screen": "100vh",
		"min":    "min-content",
		"max":    "max-content",
		"fit":    "fit-content",
	}.Extend(cfg.Spacing)

	cfg.Inset = StringMap{
		"auto": "auto",
		"1/2":  "50%",
		"1/3":  "33.333333%",
		"2/3":  "66.666667%",
		"1/4":  "25%",
		"2/4":  "50%",
		"3/4":  "75%",
		"full": "100%",
	}.Extend(cfg.Spacing)

	cfg.Margin = StringMap{
		"auto": "auto",
	}.Extend(cfg.Spacing)

	cfg.Outline.Color = cfg.Color
	cfg.Padding = cfg.Spacing
	cfg.Placeholder.Color = cfg.Color
	cfg.Placeholder.Opacity = cfg.Opacity

	cfg.Ring.Color = SStringMap{
		"DEFAULT": {
			"DEFAULT": cfg.Color.Lookup("blue", "500", "#3b82f6"),
		},
	}.Extend(cfg.Color)

	cfg.Ring.OffsetColor = cfg.Color
	cfg.Ring.Opacity = StringMap{
		"DEFAULT": "0.5",
	}.Extend(cfg.Opacity)

	cfg.Scroll.Margin = cfg.Spacing
	cfg.Scroll.Padding = cfg.Spacing

	cfg.Space = cfg.Spacing

	cfg.Stroke = SStringMap{
		"none": {
			"DEFAULT": "none",
		},
	}.Extend(cfg.Color)

	cfg.Text.Color = cfg.Color
	cfg.Text.DecorationColor = cfg.Color
	cfg.Text.Indent = cfg.Spacing
	cfg.Text.Opacity = cfg.Opacity

	cfg.Translate = StringMap{
		"1/2":  "50%",
		"1/3":  "33.333333%",
		"2/3":  "66.666667%",
		"1/4":  "25%",
		"2/4":  "50%",
		"3/4":  "75%",
		"full": "100%",
	}.Extend(cfg.Spacing)

	cfg.Width.Value = StringMap{
		"auto":   "auto",
		"1/2":    "50%",
		"1/3":    "33.333333%",
		"2/3":    "66.666667%",
		"1/4":    "25%",
		"2/4":    "50%",
		"3/4":    "75%",
		"1/5":    "20%",
		"2/5":    "40%",
		"3/5":    "60%",
		"4/5":    "80%",
		"1/6":    "16.666667%",
		"2/6":    "33.333333%",
		"3/6":    "50%",
		"4/6":    "66.666667%",
		"5/6":    "83.333333%",
		"1/12":   "8.333333%",
		"2/12":   "16.666667%",
		"3/12":   "25%",
		"4/12":   "33.333333%",
		"5/12":   "41.666667%",
		"6/12":   "50%",
		"7/12":   "58.333333%",
		"8/12":   "66.666667%",
		"9/12":   "75%",
		"10/12":  "83.333333%",
		"11/12":  "91.666667%",
		"full":   "100%",
		"screen": "100vw",
		"min":    "min-content",
		"max":    "max-content",
		"fit":    "fit-content",
	}.Extend(cfg.Spacing)

	cfg.Width.Max = StringMap{
		"none":  "none",
		"0":     "0rem",
		"xs":    "20rem",
		"sm":    "24rem",
		"md":    "28rem",
		"lg":    "32rem",
		"xl":    "36rem",
		"2xl":   "42rem",
		"3xl":   "48rem",
		"4xl":   "56rem",
		"5xl":   "64rem",
		"6xl":   "72rem",
		"7xl":   "80rem",
		"full":  "100%",
		"min":   "min-content",
		"max":   "max-content",
		"fit":   "fit-content",
		"prose": "65ch",
	}.Extend(cfg.Screens)

	return cfg
}
