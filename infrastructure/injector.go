package infrastructure

import (
	"todogo/application"
	"todogo/infrastructure/api/controller/v1"
	"todogo/infrastructure/factory/repository"
)

func InitializeTodoController(factoryAdapter string) *v1.TodoController {
	factory, err := repository.TodoFactory(factoryAdapter)
	if err != nil {
		panic(err)
	}
	application := application.NewTodoApplication(factory)
	return v1.NewTodoController(application)
}
