package handler

import (
	"fmt"
	"go-quickstart/internal/httpcontext"
	"go-quickstart/internal/templates"
	"net/http"
	"path/filepath"
)

type HandlerFunc func(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request)

func HandleFavicon(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", ".", filePath)
	http.ServeFile(w, r, fullPath)
}

func HandleStatic(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func HandleHome(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	err := httpContext.Templates.ExecuteTemplate(w, "base.html", templates.BasePageData{
		Title:   "Home",
		Content: templates.ExecuteTemplate(httpContext.Templates, "hello-world.html", nil),
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
