package main

import (
	"fmt"
	"log"
	// "path/filepath"
	// "bytes"
	// "strings"
	// "text/template"

	"github.com/jpedro/go-nuts/files/templates"
)


func main() {
	fmt.Println("hi")
	log.Default().Println("log")
	assets, _ := templates.Assets()
	fmt.Println("assets:", assets)
}

// func TemplateUpper(text string) string {
// 	return strings.ToUpper(text)
// }

// func RenderTemplate(file string, data any) (string, error) {
// 	funcs := template.FuncMap{
// 		"upper": TemplateUpper,
// 	}

// 	baseName := filepath.Base(file)
// 	tmpl, err := template.New(baseName).Funcs(funcs).ParseFiles(file)
// 	if err != nil {
// 		return "", fmt.Errorf("couldn't render template %s: %w", file, err)
// 	}

// 	buf := &bytes.Buffer{}
// 	err = tmpl.Execute(buf, data)
// 	if err != nil {
// 		return "", err
// 	}

// 	return buf.String(), nil
// }
