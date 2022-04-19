package infrastructure

import (
	"todogohexa/application"
	"todogohexa/domain"
	"todogohexa/infrastructure/api/controller/v1"
	"todogohexa/infrastructure/factory/repository"
)

func InitializeTodoController(factoryAdapter string, todos map[string]*domain.Todo) *v1.TodoController {
	factory, err := repository.TodoFactory(factoryAdapter, &todos)
	if err != nil {
		panic(err)
	}
	application := application.NewTodoApplication(factory)
	return v1.NewTodoController(application)
}
