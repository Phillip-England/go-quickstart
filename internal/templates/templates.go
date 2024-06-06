package templates

import (
	"bytes"
	"go-quickstart/internal/component"
	"html/template"
	"os"
	"path/filepath"
)

type BasePageData struct {
	Title   string
	Content template.HTML
}

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

const baseMeta = `
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<script src="https://unpkg.com/htmx.org@1.9.12"></script>
	<script src="/static/js/index.js"></script>
	<link rel="stylesheet" href="/static/css/output.css">
`

func Guest(title string, content string) string {
	return `
	<!DOCTYPE html>
	<html>
	<head>
		` + baseMeta + `
		<title>` + title + ` - CFA Suite</title>
	</head>
	<body hx-boost='true'>
		` + component.GuestHeader() + `
		<main class='p-6'>
			` + content + `
		</main>
	</body>
	</html>
	`
}

func Admin(title string, content string) string {
	return `
	<!DOCTYPE html>
	<html>
	<head>
		` + baseMeta + `
		<title>` + title + ` - CFA Suite</title>
	</head>
	<body hx-boost='true'>
		` + component.GuestHeader() + `
		<main class='p-6'>
			` + content + `
		</main>
	</body>
	</html>
	`
}
