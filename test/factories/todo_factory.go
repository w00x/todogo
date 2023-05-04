package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"todogo/domain"
	"todogo/infrastructure/repository/gorm"
	"todogo/shared"
)

type Todo struct {
	id        string
	Title     string
	Body      string
	Completed *bool
	CreatedAt shared.DateTime
}

func (i Todo) ToDomain() *domain.Todo {
	return domain.NewTodo(i.id, i.Title, i.Body, i.Completed, i.CreatedAt)
}

func NewTodo() (*Todo, error) {
	todo := &Todo{}
	err := gofakeit.Struct(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func NewTodoFactory() *domain.Todo {
	todo, err := NewTodo()
	if err != nil {
		fmt.Println(err)
	}

	repo := gorm.NewTodoRepository()
	todoDomain, errRepo := repo.Create(todo.ToDomain())
	if errRepo != nil {
		panic(err)
	}
	return todoDomain
}

func NewTodoObjectFactory() map[string]interface{} {
	todo, err := NewTodo()
	if err != nil {
		fmt.Println(err)
	}

	todoMarshal := map[string]interface{}{
		"title":     todo.Title,
		"body":      todo.Body,
		"completed": todo.Completed,
	}

	return todoMarshal
}

func NewTodoFactoryList(count int) []*domain.Todo {
	var todoDomains []*domain.Todo
	repo := gorm.NewTodoRepository()

	for i := 0; i < count; i++ {
		todo, err := NewTodo()
		if err != nil {
			panic(err)
		}

		todoDomain, errRepo := repo.Create(todo.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		todoDomains = append(todoDomains, todoDomain)
	}

	return todoDomains
}
