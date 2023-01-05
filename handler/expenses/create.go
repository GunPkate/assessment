package expenses

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) PostExpensesHandler(c echo.Context) error {
	var ex Expense
	err := c.Bind(&ex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := h.DB.QueryRow("INSERT INTO expenses (title,amount,note,tags) VALUES($1,$2,$3,$4) RETURNING id",
		ex.Title,
		ex.Amount,
		ex.Note,
		ex.Tags,
	)
	err = row.Scan(&ex.Id)
	if err != nil {
		log.Fatal("can't insert data", err)
	}

	return c.JSON(http.StatusCreated, ex)
}
