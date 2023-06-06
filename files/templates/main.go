package files

import (
	"embed"
	"io/fs"
)

//go:embed all:templates
var assets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(assets, "templates")
}
