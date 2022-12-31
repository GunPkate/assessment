package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func serverHealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

type User2 struct {
	ID   int
	AGE  int
	NAME string
}

type Err2 struct {
	Message string `json:"message"`
}

func createServerUsersHandler(c echo.Context) error {
	var u User2
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err2{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO users (name,age) VALUES($1,$2) RETURNING id", u.NAME, u.AGE)
	err = row.Scan(&u.ID) //add ID into u
	if err != nil {
		log.Fatal("can't insert data", err)
	}

	return c.JSON(http.StatusOK, u)
}

func getServerUsersHandler(c echo.Context) error {

	stmt, err := db.Prepare("SELECT * from users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err2{Message: "can't prepare query all users statement"})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err2{Message: "can't query all users:" + err.Error()})
	}

	users := []User2{}
	for rows.Next() {
		var u User2
		err = rows.Scan(&u.ID, &u.NAME, &u.AGE)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err2{Message: "can't query all users:" + err.Error()})
		}
		users = append(users, u)
	}
	return c.JSON(http.StatusOK, users)
}

func getServerUsersByIDHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT * from users WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err2{Message: "can't prepare query all users statement"})
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err2{Message: "can't query all users:" + err.Error()})
	}

	users := []User2{}
	for rows.Next() {
		var u User2
		err = rows.Scan(&u.ID, &u.NAME, &u.AGE)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err2{Message: "can't query all users:" + err.Error()})
		}
		users = append(users, u)
	}
	return c.JSON(http.StatusOK, users)
}

//shadowing
// var db *sql.DB  	new db1
// db, err : =		new db2

var db *sql.DB

// db, err = sql.Open("postgres", os.Getenv(("DATABASE_URL")))
// DATABASE_URL = postgres://mkzchuoq:loPAe5lWPs4gsdvrMf2aKchys2xsGF0x@tiny.db.elephantsql.com/mkzchuoq
func main() {
	var err error
	var url string = "postgres://mkzchuoq:loPAe5lWPs4gsdvrMf2aKchys2xsGF0x@tiny.db.elephantsql.com/mkzchuoq"
	// db, err = sql.Open("postgres", os.Getenv(("DATABASE_URL")))
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connection error", err)
	}
	fmt.Println("Server running")
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT,age INT);`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", serverHealthHandler)
	e.GET("/users", getServerUsersHandler)
	e.GET("/users/:id", getServerUsersByIDHandler)
	e.POST("/users", createServerUsersHandler)

	log.Println("Server start at: 2500")
	log.Fatal(e.Start(":2500"))
	log.Println("Bye")
}
