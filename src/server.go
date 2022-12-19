package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/herokh/assessment/expense"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var appPort = os.Getenv("PORT")
	var dbUrl = os.Getenv("DATABASE_URL")

	fmt.Println("application started.")
	fmt.Printf("port number: %s\n", appPort)
	fmt.Printf("database url: %s\n", dbUrl)

	expense.InitDB(dbUrl)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth != "November 10, 2009" {
				return c.NoContent(http.StatusUnauthorized)
			}
			return next(c)
		}
	})

	e.GET("/expenses", expense.GetExpensesHandler)
	e.GET("/expenses/:id", expense.GetExpenseHandler)
	e.POST("/expenses", expense.CreateExpenseHandler)
	e.PUT("/expenses/:id", expense.UpdateExpenseHandler)

	go func() {
		if err := e.Start(appPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
