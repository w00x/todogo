package test

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"todogo/application"
	repository2 "todogo/domain/repository"
	v1 "todogo/infrastructure/api/controller/v1"
	"todogo/infrastructure/errors"
	"todogo/infrastructure/factory/repository"
	local_gorm "todogo/infrastructure/repository/gorm"
	"todogo/shared"
)

func TodoControllerInjector() *v1.TodoController {
	return v1.NewTodoController(TodoApplicationInjector())
}

func TodoApplicationInjector() *application.TodoApplication {
	factory, err := TodoRepository()
	if err != nil {
		panic(err)
	}
	return application.NewTodoApplication(factory)
}

func TodoRepository() (repository2.ITodoRepository, errors.IBaseError) {
	return repository.TodoFactory(GetCurrentAdapter())
}

func CleanTodo() {
	local_gorm.NewPostgresBase().DB.Exec("DELETE FROM todos")
}

func SetDbConn() {
	uri := shared.GetEnv("DATABASE_TEST_URI")
	var err error
	local_gorm.DbConn, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetCurrentAdapter() string {
	return "gorm"
}
