package route

import (
	"fmt"
	"go-quickstart/internal/handler"
	"go-quickstart/internal/middleware"
	"go-quickstart/internal/templates"
	"html/template"
	"net/http"
)

type Router struct {
	Mux       *http.ServeMux
	Templates *template.Template
}

func NewRouter() (*Router, error) {
	templates, err := templates.ParseTemplates()
	if err != nil {
		return nil, err

	}
	return &Router{
		Mux:       http.NewServeMux(),
		Templates: templates,
	}, nil
}

func (r *Router) Add(path string, handler handler.HandlerFunc, middleware ...middleware.MiddlewareFunc) {
	isIndex := false
	if path == "GET /" || path == "POST /" || path == "PUT /" || path == "DELETE /" || path == "PATCH /" || path == "OPTIONS /" || path == "HEAD /" {
		isIndex = true
	}
	route := &Route{
		mux:        r.Mux,
		path:       path,
		handler:    handler,
		templates:  r.Templates,
		middleware: middleware,
		isIndex:    isIndex,
	}
	if isIndex {
		route.HandleIndex()
		return
	}
	route.Handle()
}

func (r *Router) Serve(port string, message string) {
	fmt.Println(message)
	http.ListenAndServe(":"+port, r.Mux)
}

type Route struct {
	mux        *http.ServeMux
	path       string
	handler    handler.HandlerFunc
	templates  *template.Template
	middleware []middleware.MiddlewareFunc
	isIndex    bool
}

func (r *Route) Handle() {
	r.mux.HandleFunc(r.path, func(w http.ResponseWriter, req *http.Request) {
		middleware.Chain(w, req, r.templates, r.handler, r.middleware...)
	})
}

func (r *Route) HandleIndex() {
	r.mux.HandleFunc(r.path, func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		middleware.Chain(w, req, r.templates, r.handler, r.middleware...)
	})
}
