# go-quickstart

go-quickstart is intended to get you a Go server ASAP.

## Serving

To serve the application, simply run:

```bash
go run main.go
```

## Custom Middleware

go-quickstart comes with a custom middleware implemenatation. Instead of deeply nesting all of your middleware logic, go-quickstart enables you to chain middleware to your routes using the `Chain` func found in `/internal/middleware/middleware.go`.

### Creating Custom Middleware

To create a new custom middleware, let's head over to `/internal/middleware/middleware.go`.

```go
type CustomContext struct {
	context.Context
	Templates *template.Template
	StartTime time.Time
}

type CustomHandler func(ctx *CustomContext, w http.ResponseWriter, r *http.Request)
type CustomMiddleware func(ctx *CustomContext, w http.ResponseWriter, r *http.Request) error

func Chain(w http.ResponseWriter, r *http.Request, templates *template.Template, handler CustomHandler, middleware ...CustomMiddleware) {
	customContext := &CustomContext{
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

func Log(ctx *CustomContext, w http.ResponseWriter, r *http.Request) error {
	elapsedTime := time.Since(ctx.StartTime)
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] [%s] [%s]\n", formattedTime, r.Method, r.URL.Path, elapsedTime)
	return nil
}
```

You can already see the middleware func `Log` defined at the bottom of the file. Let's define a new middleware called `HelloWorld`.

At the bottom of the file, write:

```go
func HelloWorld(ctx *CustomContext, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Hello, World")
	return nil
}
```

Now, take note. All custom middleware must take in `*CustomContext`, `http.ResponseWriter`, and `*http.Request` as parameters. This is what defines the func as a custom middleware.

### Using Custom Middleware

To use the custom middleware, lets take a look at `main.go`:

```go
func main() {
    // ...
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        // handles 404 errors
        // only required on the "/" endpoint
        if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		middleware.Chain(w, r, templates, HandlerHome)
	})
}
```

Now, let's add our new `HelloWorld` middleware to the chain.

```go
middleware.Chain(w, r, templates, HandlerHome, middleware.HelloWorld)
```

That's it! You've created a new custom middleware and added it to the "/" endpoint.

Serve the application with `go run main.go` and visit localhost:8080. You should see the following output printed to the console:

```bash
Hello, World
```

### Middleware Errors

If a middleware returns an error, it will be handled in the `Chain` func. This enables you to stop execution and handle errors with ease.

Let's look at the `Chain` func in, `/internal/middleware/middleware.go`:

```go
func Chain(w http.ResponseWriter, r *http.Request, templates *template.Template, handler CustomHandler, middleware ...CustomMiddleware) {
	customContext := &CustomContext{
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
```

Zooming in to where we handle potential middleware errors:

```go
for _, mw := range middleware {
    err := mw(customContext, w, r)
    if err != nil {
        return // if an error occurs, do not run the handler
    }
    handler(customContext, w, r)
}
```

## Parsing Templates

go-quickstart already handles parsing all of your templates at `/html`. It does so with the `ParseTemplates` func found at `/internal/filehandler/filehandler.go`.

Just stash new templates anywhere in `/html` and they will be parsed at the start of your programs execution.


