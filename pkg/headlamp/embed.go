package headlamp

import (
	"embed"
)

//go:embed dist/*
var DistFS embed.FS
