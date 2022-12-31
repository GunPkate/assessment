package main

import (
	"database/sql"
	"fmt"
	"log"
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func dbConnection() {
	var err error
	var url string = "postgres://mkzchuoq:loPAe5lWPs4gsdvrMf2aKchys2xsGF0x@tiny.db.elephantsql.com/mkzchuoq"
	// db, err = sql.Open("postgres", os.Getenv(("DATABASE_URL")))
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connection error", err)
	}
	fmt.Println("Server running")
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
}



var db *sql.DB

func main() {

	dbConnection()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/health", serverHealthHandler)
	// e.GET("/users", getServerUsersHandler)
	// e.GET("/users/:id", getServerUsersByIDHandler)
	// e.POST("/users", createServerUsersHandler)

	log.Println("Server start at: 2565")
	log.Fatal(e.Start(":2565"))
	log.Println("Bye")
}
