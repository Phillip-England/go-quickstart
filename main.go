package main

import (
	"fmt"
	"go-quickstart/internal/filehandler"
	"go-quickstart/internal/handler"
	"go-quickstart/internal/middleware"
	"net/http"
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
		middleware.Chain(w, r, templates, handler.HandleHome)
	})

	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, templates, handler.HandleFavicon)
	})

	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, templates, handler.HandleStatic)
	})

	fmt.Println("Running Development Server on localhost:" + port)
	http.ListenAndServe(":"+port, mux)
}
