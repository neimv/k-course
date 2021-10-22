package main

import "gorm.io/gorm"

type ToDoORM struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Active      bool   `json:"active"`
}
