package handler

import (
	"fmt"
	"go-quickstart/internal/filehandler"
	"go-quickstart/internal/middleware"
	"go-quickstart/internal/stypes"
	"net/http"
	"path/filepath"
)

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
		Title:   "Home",
		Content: filehandler.ExecuteTemplate(customContext.Templates, "hello-world.html", nil),
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
