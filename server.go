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

func getExpensesByIDHandler(c echo.Context) error {
	id := c.Param("id")
	var ex Expense
	row := db.QueryRow("SELECT * from expenses WHERE id = $1", id)

	err := row.Scan(&ex.Id, &ex.Title, &ex.Amount, &ex.Note, &ex.Tags)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query selected expenses statement"})
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

	e.GET("/expenses/:id", getExpensesByIDHandler)

	log.Println("Server start at: 2565")
	log.Fatal(e.Start(":2565"))
	log.Println("Bye")
}
