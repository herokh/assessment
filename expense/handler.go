package expense

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetExpensesHandler(c echo.Context) error {
	expenses, err := GetExpenses(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, expenses)
}
func GetExpenseHandler(c echo.Context) error {
	expense, err := GetExpense(db, c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, expense)
}
func CreateExpenseHandler(c echo.Context) error {
	var expense Expense
	err := c.Bind(&expense)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	err = CreateExpense(db, &expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, expense)
}
func UpdateExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	var expense Expense
	err := c.Bind(&expense)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	err = UpdateExpense(db, id, &expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, expense)
}
