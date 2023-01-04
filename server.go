package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB

// func dbConnection() {

// }

type Expense struct {
	Id     string         `json:"id"`
	Title  string         `json:"title"`
	Amount float64        `json:"amount"`
	Note   string         `json:"note"`
	Tags   pq.StringArray `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}

func updateExpensesHandler(c echo.Context) error {
	var ex Expense

	id := c.Param("id")

	if err := c.Bind(&ex); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	ex.Id = id
	stmt, err := db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	_, err = stmt.Exec(ex.Id, ex.Title, ex.Amount, ex.Note, ex.Tags)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ex)
}

func main() {

	// dbConnection()
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.PUT("/expenses/:id", updateExpensesHandler)

	log.Println("Server start at: 2565")
	log.Fatal(e.Start(":2565"))
	log.Println("Bye")
}
