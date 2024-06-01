package main

import (
	"errors"
	"fmt"
	"go-quickstart/internal/handler"
	"go-quickstart/internal/httpcontext"
	"go-quickstart/internal/route"
	"net/http"
)

func main() {

	// creating a new router
	r, err := route.NewRouter()
	if err != nil {
		panic(err)
	}

	// handling favicon and static files
	r.Add("GET /favicon.ico", handler.HandleFavicon)
	r.Add("GET /static/", handler.HandleStatic)

	// handling home page
	r.Add("GET /", handler.HandleHome, CustomMiddleware) // chaining middleware

	// handling page to demonstrate exiting early from middleware
	r.Add("GET /exit", handler.HandleHome, CustomMiddleware, ExitMiddleware) // chaining middleware

	// serving at 8080
	port := "8080"
	r.Serve(port, fmt.Sprintf("Server is running on port %s", port))

}

// a custom middleware
func CustomMiddleware(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Executing custom middleware")
	return nil
}

// a custom middleware that exits early
func ExitMiddleware(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("exiting from middleware\n"))
	return errors.New("exit!")
}
