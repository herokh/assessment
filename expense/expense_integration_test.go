//go:build integration
// +build integration

package expense_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/herokh/assessment/expense"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var serverPort = 80
var dbUrl = "postgres://root:root@db/go-test-db?sslmode=disable"

func TestIntegrationGetAll(t *testing.T) {
	// Setup server
	eh := echo.New()
	expense.InitDB(dbUrl)

	// Arrange
	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := eh.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, expense.GetExpensesHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestIntegrationGetOne(t *testing.T) {
	// Setup server
	eh := echo.New()
	expense.InitDB(dbUrl)

	// Arrange
	reqBody := ``
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/", serverPort), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := eh.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, expense.GetExpenseHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestIntegrationCreate(t *testing.T) {
	// Setup server
	eh := echo.New()
	expense.InitDB(dbUrl)

	// Arrange
	var data = strings.NewReader(`
	{
		"title" : "title",
		"amount" : 100,
		"note" : "note",
		"tags" : ["tag1"]
	}
	`)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/", serverPort), data)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := eh.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, expense.CreateExpenseHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func IntegrationUpdate(t *testing.T) {
	// Setup server
	eh := echo.New()
	expense.InitDB(dbUrl)

	// Arrange
	var data = strings.NewReader(`
	{
		"title" : "title",
		"amount" : 100,
		"note" : "note",
		"tags" : ["tag1"]
	}
	`)
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%d/", serverPort), data)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := eh.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, expense.UpdateExpenseHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
