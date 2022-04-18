package domain

import "github.com/google/uuid"

type Todo struct {
	id			string
	Title 		string
	Body 		string
	Completed 	bool
}

func NewTodo(id string, title string, body string, completed bool) *Todo {
	if id == "" {
		id = uuid.New().String()
	}

	return &Todo{
		id:        id,
		Title:     title,
		Body:      body,
		Completed: completed,
	}
}

func (i Todo) Id() string {
	return i.id
}

func (i *Todo) Toggle() bool {
	i.Completed = !i.Completed
	return i.Completed
}