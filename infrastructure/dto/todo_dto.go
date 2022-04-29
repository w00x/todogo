package dto

import "todogo/domain"

type TodoDto struct {
	Id			string	`json:"id" uri:"id"`
	Title 		string	`json:"title"`
	Body 		string	`json:"body"`
	Completed 	*bool	`json:"completed"`
}

func NewTodoDto(id string, title string, body string, completed *bool) *TodoDto {
	return &TodoDto{id, title, body, completed}
}

func NewTodoDtoFromDomain(todo *domain.Todo) *TodoDto {
	return NewTodoDto(todo.Id(), todo.Title, todo.Body, todo.Completed)
}

func NewTodoListDtoFromDomains(inventories *[]domain.Todo) []*TodoDto {
	var todoDtos []*TodoDto
	for _, todo := range *inventories {
		todoDtos = append(todoDtos,
			NewTodoDtoFromDomain(&todo))
	}
	return todoDtos
}

func (i TodoDto) ToDomain() *domain.Todo {
	return domain.NewTodo(i.Id, i.Title, i.Body, i.Completed)
}