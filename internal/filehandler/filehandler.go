package filehandler

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

func ParseTemplates() (*template.Template, error) {
	templates := template.New("")
	err := filepath.Walk("./html", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			_, err := templates.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func ExecuteTemplate(t *template.Template, name string, data interface{}) template.HTML {
	var templateContent bytes.Buffer
	err := t.ExecuteTemplate(&templateContent, name, data)
	if err != nil {
		panic(err)
	}
	return template.HTML(templateContent.String())
}
