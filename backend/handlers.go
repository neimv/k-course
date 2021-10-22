package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// function to create TODO
func CreateTodo(c echo.Context) error {
	db := GetDB()
	var td ToDoORM

	if err := c.Bind(&td); err != nil {
		fmt.Println(err)
		return err
	}

	td.Active = true
	db.Create(&td)

	return c.JSON(http.StatusCreated, td)
}

// Function to get all TODOs
func GetTodos(c echo.Context) error {
	var todos []ToDoORM
	db := GetDB()

	db.Find(&todos)

	return c.JSON(http.StatusOK, todos)
}

// Function to get one TODOs by id
func GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db := GetDB()
	var todo ToDoORM

	result := db.First(&todo, id)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, todo)
}

// Function to delete one TODO
func DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db := GetDB()
	var todo ToDoORM

	db.Delete(&todo, id)

	return c.NoContent(http.StatusNoContent)
}

// Function to update one TODO
func UpdateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db := GetDB()
	var todo ToDoORM
	var todo_update ToDoORM

	if err := c.Bind(&todo); err != nil {
		return err
	}

	db.First(&todo_update, id)

	if todo.Title != "" {
		todo_update.Title = todo.Title
	}
	if todo.Description != "" {
		todo_update.Description = todo.Description
	}
	if todo.Active != todo_update.Active {
		todo_update.Active = todo.Active
	}

	db.Save(&todo_update)

	return c.NoContent(http.StatusNoContent)
}
