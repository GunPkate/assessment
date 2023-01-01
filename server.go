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
	// id     string  `json:"id"`
	// title  string  `json:"title"`
	// amount float64 `json:"amount"`
	// note   string  `json:"note"`
	// tags   string  `json:"tags"`
	id     string
	title  string
	amount float64
	note   string
	tags   pq.StringArray
}

type Err struct {
	Message string `json:"message"`
}

func getExpensesByIDHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT * from expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all expenses statement"})
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all expenses:" + err.Error()})
	}

	expenses := []Expense{}
	for rows.Next() {
		var ex Expense
		err = rows.Scan(&ex.id, &ex.title, &ex.amount, &ex.note, &ex.tags)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all expenses:" + err.Error()})
		}
		expenses = append(expenses, ex)
	}
	return c.JSON(http.StatusOK, expenses)
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
