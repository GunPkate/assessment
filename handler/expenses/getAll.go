package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetAllExpensesHandler(c echo.Context) error {

	stmt, err := h.DB.Prepare("SELECT * from expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all expenses statement"})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all expenses:" + err.Error()})
	}

	expenses := []Expense{}
	for rows.Next() {
		var ex Expense
		err = rows.Scan(&ex.Id, &ex.Title, &ex.Amount, &ex.Note, &ex.Tags)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all expenses:" + err.Error()})
		}
		expenses = append(expenses, ex)
	}
	return c.JSON(http.StatusOK, expenses)
}
