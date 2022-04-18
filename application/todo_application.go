package application

import (
	"todogohexa/domain"
	"todogohexa/domain/repository"
	"todogohexa/infrastructure/errors"
)

type TodoApplication struct {
	todoRepository repository.ITodoRepository
}

func NewTodoApplication(todoRepository repository.ITodoRepository) *TodoApplication {
	return &TodoApplication{todoRepository: todoRepository}
}

func (ta TodoApplication) All() (*[]domain.Todo, errors.IBaseError) {
	return ta.todoRepository.All()
}

func (ta TodoApplication) Show(id string) (*domain.Todo, errors.IBaseError) {
	return ta.todoRepository.FindById(id)
}

func (ta TodoApplication) Create(todo *domain.Todo) (*domain.Todo, errors.IBaseError) {
	return ta.todoRepository.Create(todo)
}

func (ta TodoApplication) Update(todo *domain.Todo) (*domain.Todo, errors.IBaseError) {
	return ta.todoRepository.Update(todo)
}

func (ta TodoApplication) Delete(todo *domain.Todo) errors.IBaseError {
	return ta.todoRepository.Destroy(todo)
}

func (ta TodoApplication) Toggle(id string) (*domain.Todo, errors.IBaseError) {
	todo, err := ta.todoRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	todo.Toggle()
	return ta.todoRepository.Update(todo)
}