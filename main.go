package main

import (
	"fmt"
	"hotswap/internal/filehandler"
	"hotswap/internal/middleware"
	"hotswap/internal/stypes"
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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, templates, HandlerHome)
	})

	fmt.Println("Running Development Server on localhost:" + port)
	http.ListenAndServe(":"+port, mux)

}

func HandlerHome(customContext *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	err := customContext.Templates.ExecuteTemplate(w, "base.html", stypes.BasePageData{
		Title: "Home",
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
