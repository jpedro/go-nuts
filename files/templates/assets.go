package files

import (
	"embed"
	"io/fs"
)

//go:embed all:nested
var assets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(assets, "nested")
}
