package httpcontext

import (
	"context"
	"html/template"

	"time"
)

type Context struct {
	context.Context
	Templates *template.Template
	StartTime time.Time
}
