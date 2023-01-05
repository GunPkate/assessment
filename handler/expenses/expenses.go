package expenses

import (
	"database/sql"

	"github.com/lib/pq"
)

type handler struct {
	DB *sql.DB
}

func DdbConnection(db *sql.DB) *handler {
	return &handler{db}
}

type Expense struct {
	Id     int            `json:"id"`
	Title  string         `json:"title"`
	Amount float64        `json:"amount"`
	Note   string         `json:"note"`
	Tags   pq.StringArray `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}
