package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID 			string		`gorm:"type:uuid;primaryKey;column:id"`
	Title 		string
	Body 		string
	Completed 	bool
}

func NewTodo(id string, title string, body string, completed bool) *Todo {
	return &Todo{ID: id, Title: title, Body: body, Completed: completed}
}