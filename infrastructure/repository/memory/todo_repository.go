package memory

import (
	"todogohexa/domain"
	"todogohexa/infrastructure/errors"
)

type TodoRepository struct {
	Todos *map[string]*domain.Todo
}

func (i *TodoRepository) All() (*[]domain.Todo, errors.IBaseError) {
	var todos []domain.Todo
	for _, value := range *i.Todos {
		todos = append(todos, *value)
	}
	return &todos, nil
}

func (i *TodoRepository) FindById(id string) (*domain.Todo, errors.IBaseError) {
	return (*i.Todos)[id], nil
}

func (i *TodoRepository) Update(todo *domain.Todo) (*domain.Todo, errors.IBaseError) {
	newTodo := (*i.Todos)[todo.Id()]

	if todo.Title != "" {
		newTodo.Title = todo.Title
	}

	if todo.Body != "" {
		newTodo.Body = todo.Body
	}

	if todo.Completed != nil {
		newTodo.Completed = todo.Completed
	}

	(*i.Todos)[todo.Id()] = newTodo
	return newTodo, nil
}

func (i *TodoRepository) Create(todo *domain.Todo) (*domain.Todo, errors.IBaseError) {
	(*i.Todos)[todo.Id()] = todo
	return todo, nil
}

func (i *TodoRepository) Destroy(todo *domain.Todo) errors.IBaseError {
	delete(*i.Todos, todo.Id())
	return nil
}
