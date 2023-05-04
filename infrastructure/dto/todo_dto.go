package dto

import (
	"todogo/domain"
	"todogo/shared"
)

type TodoDto struct {
	Id        string          `json:"id" uri:"id"`
	Title     string          `json:"title"`
	Body      string          `json:"body"`
	Completed *bool           `json:"completed"`
	CreatedAt shared.DateTime `json:"created_at"`
}

func NewTodoDto(id string, title string, body string, completed *bool, createdAt shared.DateTime) *TodoDto {
	return &TodoDto{id, title, body, completed, createdAt}
}

func NewTodoDtoFromDomain(todo *domain.Todo) *TodoDto {
	return NewTodoDto(todo.Id(), todo.Title, todo.Body, todo.Completed, todo.CreatedAt)
}

func NewTodoListDtoFromDomains(todos *[]domain.Todo) []*TodoDto {
	var todoDtos []*TodoDto
	for _, todo := range *todos {
		todoDtos = append(todoDtos,
			NewTodoDtoFromDomain(&todo))
	}
	return todoDtos
}

func (i TodoDto) ToDomain() *domain.Todo {
	return domain.NewTodo(i.Id, i.Title, i.Body, i.Completed, i.CreatedAt)
}
