package repository

import (
	"todogohexa/domain"
	"todogohexa/infrastructure/errors"
)

type ITodoRepository interface {
	All() (*[]domain.Todo, errors.IBaseError)
	FindById(id string) (*domain.Todo, errors.IBaseError)
	Update(todo *domain.Todo) (*domain.Todo, errors.IBaseError)
	Create(todo *domain.Todo) (*domain.Todo, errors.IBaseError)
	Destroy(todo *domain.Todo) errors.IBaseError
}
