package main

import (
	"database/sql"
	"log"

	"github.com/GunPkate/assessment/handler/expenses"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func initConnection() *sql.DB {
	// connStr := os.Getenv("DATABASE_URL")
	connStr := "postgres://mkzchuoq:loPAe5lWPs4gsdvrMf2aKchys2xsGF0x@tiny.db.elephantsql.com/mkzchuoq"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	return db
}

func main() {
	db := initConnection()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := expenses.DdbConnection(db)
	e.GET("/expenses/:id", h.GetExpensesByIDHandler)
	e.POST("/expenses", h.PostExpensesHandler)
	e.GET("/expenses", h.GetAllExpensesHandler)
	e.PUT("/expenses/:id", h.UpdateExpensesHandler)

	log.Println("Server start at: 2565")
	log.Fatal(e.Start(":2565"))
	log.Println("Bye")
}

func dbConnection(db *sql.DB) {
	panic("unimplemented")
}
