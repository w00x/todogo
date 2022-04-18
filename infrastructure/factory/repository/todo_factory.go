package repository

import (
	"todogohexa/domain"
	"todogohexa/domain/repository"
	"todogohexa/infrastructure/errors"
	"todogohexa/infrastructure/repository/gorm"
	"todogohexa/infrastructure/repository/memory"
)

func TodoFactory(adapter string, db interface{}) (repository.ITodoRepository, errors.IBaseError) {
	if adapter == "memory" {
		return &memory.TodoRepository{Todos: db.(*map[string]*domain.Todo)}, nil
	} else if adapter == "gorm" {
		return gorm.NewTodoRepository(), nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}
