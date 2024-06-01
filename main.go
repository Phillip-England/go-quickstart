package main

import (
	"fmt"
	"go-quickstart/internal/handler"
	"go-quickstart/internal/route"
)

func main() {

	r, err := route.NewRouter()
	if err != nil {
		fmt.Println(err)
		return
	}

	r.Add("GET /", handler.HandleHome)
	r.Add("GET /favicon.ico", handler.HandleFavicon)
	r.Add("GET /static/", handler.HandleStatic)

	port := "8080"
	r.Serve(port, fmt.Sprintf("Server is running on port %s", port))

}
