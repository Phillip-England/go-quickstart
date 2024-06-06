package main

import (
	"fmt"
	"go-quickstart/internal/appform"
	"go-quickstart/internal/database"
	"go-quickstart/internal/handler"
	"go-quickstart/internal/middleware"
	"go-quickstart/internal/route"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	_ = godotenv.Load()

	db, err := sqlx.Connect("sqlite3", "main.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = database.CreateTables(db)
	if err != nil {
		log.Fatalln(err)
	}

	database.PrintTables(db)

	r, err := route.NewRouter()
	if err != nil {
		panic(err)
	}

	r.Add("GET /favicon.ico", handler.Favicon)
	r.Add("GET /static/", handler.Static)

	r.Add("GET /", handler.PageHome, middleware.IsNotGuest)
	r.Add("GET /admin", handler.PageAdminPanel)

	r.Add("POST /", appform.Login)

	port := "8080"
	r.Serve(port, fmt.Sprintf("Server is running on port %s", port))

}
