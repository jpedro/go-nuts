package files

import (
	"embed"
	// "io/fs"
)

//go:embed all:nested
var Assets embed.FS

// func Assets() (fs.FS, error) {
// 	return fs.Sub(assets, "nested")
// }

// func ReadFile(path string) ([]byte, error) {
// 	return assets.ReadFile(path)
// }
