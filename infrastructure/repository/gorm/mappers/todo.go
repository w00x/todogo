package mappers

import (
	"todogo/domain"
	"todogo/infrastructure/repository/gorm/models"
)

func FromTodoDomainToModel(i *domain.Todo) *models.Todo {
	return models.NewTodo(i.Id(), i.Title, i.Body, i.Completed)
}

func FromTodoModelToDomain(i *models.Todo) *domain.Todo {
	return domain.NewTodo(i.ID, i.Title, i.Body, i.Completed)
}

func NewTodoListDomainFromModel(todos *[]models.Todo) *[]domain.Todo {
	var todosDomain []domain.Todo
	for _, todo := range *todos {
		todosDomain = append(todosDomain, *FromTodoModelToDomain(&todo))
	}
	return &todosDomain
}