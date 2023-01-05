package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) UpdateExpensesHandler(c echo.Context) error {
	var ex Expense

	id := c.Param("id")

	if err := c.Bind(&ex); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	ex.Id = id
	stmt, err := h.DB.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	_, err = stmt.Exec(ex.Id, ex.Title, ex.Amount, ex.Note, ex.Tags)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ex)
}
