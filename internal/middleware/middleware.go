package middleware

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type CustomContext struct {
	context.Context
	Templates *template.Template
	StartTime time.Time
}

type CustomHandler func(ctx *CustomContext, w http.ResponseWriter, r *http.Request)
type CustomMiddleware func(ctx *CustomContext, w http.ResponseWriter, r *http.Request) error

func Chain(w http.ResponseWriter, r *http.Request, templates *template.Template, handler CustomHandler, middleware ...CustomMiddleware) {
	customContext := &CustomContext{
		Context:   context.Background(),
		Templates: templates,
		StartTime: time.Now(),
	}
	for _, mw := range middleware {
		err := mw(customContext, w, r)
		if err != nil {
			return
		}
	}
	handler(customContext, w, r)
	Log(customContext, w, r)
}

func Log(ctx *CustomContext, w http.ResponseWriter, r *http.Request) error {
	elapsedTime := time.Since(ctx.StartTime)
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] [%s] [%s]\n", formattedTime, r.Method, r.URL.Path, elapsedTime)
	return nil
}
