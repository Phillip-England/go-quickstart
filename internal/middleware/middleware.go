package middleware

import (
	"context"
	"fmt"
	"go-quickstart/internal/handler"
	"go-quickstart/internal/httpcontext"
	"html/template"
	"net/http"
	"os"
	"time"
)

type MiddlewareFunc func(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error
type MiddlewareChainFunc func(w http.ResponseWriter, r *http.Request, templates *template.Template, handler handler.HandlerFunc, middleware ...MiddlewareFunc)

func Chain(w http.ResponseWriter, r *http.Request, templates *template.Template, handler handler.HandlerFunc, middleware ...MiddlewareFunc) {
	customContext := &httpcontext.Context{
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

func Log(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error {
	elapsedTime := time.Since(ctx.StartTime)
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] [%s] [%s]\n", formattedTime, r.Method, r.URL.Path, elapsedTime)
	return nil
}

func IsNotGuest(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}
	if cookie.Value == os.Getenv("ADMIN_SESSION_TOKEN") {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
	return nil
}
