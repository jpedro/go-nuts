package main

import (
	"fmt"
	"log"

	"github.com/jpedro/go-nuts/files/assets"
)

func main() {
	fmt.Println("hi")
	log.Default().Println("log")
	// assets, _ := templates.Assets()
	bytes, err := assets.Assets.ReadFile("nested/directory/some.file")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	// text := os.ReadFile()
	fmt.Println("bytes:", string(bytes))
	// fmt.Printf("file: %T\n", file)
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
