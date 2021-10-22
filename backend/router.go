package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateRouter() *echo.Echo {
	router := echo.New()

	// middlewares
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Use(middleware.CORS())

	// Routes
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	router.GET("/todos", GetTodos)
	router.POST("/todos", CreateTodo)
	router.GET("/todos/:id", GetTodo)
	router.PUT("/todos/:id", UpdateTodo)
	router.DELETE("/todos/:id", DeleteTodo)

	return router
}
