package main

import (
	"fmt"
	"go-quickstart/internal/filehandler"
	"go-quickstart/internal/middleware"
	"go-quickstart/internal/stypes"
	"net/http"
	"path/filepath"
)

func main() {

	port := "8080"
	templates, err := filehandler.ParseTemplates()
	if err != nil {
		fmt.Println(err)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		middleware.Chain(w, r, templates, HandleHome)
	})

	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, templates, HandleFavicon)
	})

	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, templates, HandleStatic)
	})

	fmt.Println("Running Development Server on localhost:" + port)
	http.ListenAndServe(":"+port, mux)
}

func HandleFavicon(customContext *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", ".", filePath)
	http.ServeFile(w, r, fullPath)
}

func HandleStatic(customContext *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func HandleHome(customContext *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	err := customContext.Templates.ExecuteTemplate(w, "base.html", stypes.BasePageData{
		Title: "Home",
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
