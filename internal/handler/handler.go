package handler

import (
	"go-quickstart/internal/component"
	"go-quickstart/internal/httpcontext"
	"go-quickstart/internal/templates"
	"net/http"
	"path/filepath"
)

type HandlerFunc func(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request)

func Favicon(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", ".", filePath)
	http.ServeFile(w, r, fullPath)
}

func Static(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func PageHome(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	loginErr := r.URL.Query().Get("loginErr")
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	w.Write([]byte(templates.Guest("Home", `
		`+component.LoginForm(email, password, loginErr)+`
	`)))
}

func PageAdminPanel(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(templates.Admin("Admin Panel", `
		<h1>Admin Panel</h1>
	`)))
}
