package memory

import (
	"todogo/domain"
	"todogo/infrastructure/errors"
)

type TodoRepository struct {
	Todos *map[string]*domain.Todo
}

func (i *TodoRepository) All() (*[]domain.Todo, errors.IBaseError) {
	var todos []domain.Todo
	for _, value := range *TodoMemoryFactory() {
		todos = append(todos, *value)
	}
	return &todos, nil
}

func (i *TodoRepository) FindById(id string) (*domain.Todo, errors.IBaseError) {
	todo := (*TodoMemoryFactory())[id]
	if todo == nil {
		return nil, errors.NewNotFoundError("Repository not found")
	}
	return todo, nil
}

func (i *TodoRepository) Update(todo *domain.Todo) (*domain.Todo, errors.IBaseError) {
	newTodo := (*TodoMemoryFactory())[todo.Id()]

	if todo.Title != "" {
		newTodo.Title = todo.Title
	}

	if todo.Body != "" {
		newTodo.Body = todo.Body
	}

	if todo.Completed != nil {
		newTodo.Completed = todo.Completed
	}

	(*TodoMemoryFactory())[todo.Id()] = newTodo
	return newTodo, nil
}

func (i *TodoRepository) Create(todo *domain.Todo) (*domain.Todo, errors.IBaseError) {
	(*TodoMemoryFactory())[todo.Id()] = todo
	return todo, nil
}

func (i *TodoRepository) Destroy(todo *domain.Todo) errors.IBaseError {
	delete(*TodoMemoryFactory(), todo.Id())
	return nil
}

var todos map[string]*domain.Todo

func TodoMemoryFactory() *map[string]*domain.Todo {
	if todos != nil {
		return &todos
	}

	todos = make(map[string]*domain.Todo)
	return &todos
}