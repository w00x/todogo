package repository

import (
	"todogo/domain/repository"
	"todogo/infrastructure/errors"
	"todogo/infrastructure/repository/gorm"
	"todogo/infrastructure/repository/memory"
)

func TodoFactory(adapter string) (repository.ITodoRepository, errors.IBaseError) {
	if adapter == "memory" {
		return &memory.TodoRepository{}, nil
	} else if adapter == "gorm" {
		return gorm.NewTodoRepository(), nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}
