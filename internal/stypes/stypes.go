package stypes

import (
	"html/template"
)

type BasePageData struct {
	Title   string
	Content template.HTML
}
