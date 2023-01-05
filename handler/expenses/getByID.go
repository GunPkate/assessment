package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetExpensesByIDHandler(c echo.Context) error {
	id := c.Param("id")
	var ex Expense
	row := h.DB.QueryRow("SELECT * from expenses WHERE id = $1", id)

	err := row.Scan(&ex.Id, &ex.Title, &ex.Amount, &ex.Note, &ex.Tags)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query selected expenses statement"})
	}

	return c.JSON(http.StatusOK, ex)
}
