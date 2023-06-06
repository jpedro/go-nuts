package files

import (
	"embed"
	"io/fs"
)

//go:embed all:nested
var Assets embed.FS

func OtherAssets() (fs.FS, error) {
	return fs.Sub(Assets, "nested")
}
