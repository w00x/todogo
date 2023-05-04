package domain

import (
	"github.com/google/uuid"
	"todogo/shared"
)

type Todo struct {
	id        string
	Title     string
	Body      string
	Completed *bool
	CreatedAt shared.DateTime
}

func NewTodo(id string, title string, body string, completed *bool, createdAt shared.DateTime) *Todo {
	if id == "" {
		id = uuid.New().String()
	}

	return &Todo{
		id:        id,
		Title:     title,
		Body:      body,
		Completed: completed,
		CreatedAt: createdAt,
	}
}

func (i Todo) Id() string {
	return i.id
}

func (i *Todo) Toggle() bool {
	if i.Completed == nil {
		i.Completed = new(bool)
	}

	*i.Completed = !*i.Completed
	return *i.Completed
}
