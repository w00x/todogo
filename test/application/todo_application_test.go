package application

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
	"todogo/application"
	"todogo/domain"
	"todogo/infrastructure/factory/repository"
	"todogo/test"
	"todogo/test/factories"
)

func TestMain(m *testing.M) {
	test.SetDbConn()
	ret := m.Run()
	test.CleanTodo()
	os.Exit(ret)
}

func TestNewTodoApplication(t *testing.T) {
	factory, err := repository.TodoFactory("memory")
	if err != nil {
		panic(err)
	}
	todoApplication := application.NewTodoApplication(factory)

	assert.NotNil(t, todoApplication)
}

func TestTodoApplication_All(t *testing.T) {
	todoApplication := test.TodoApplicationInjector()
	todosFactory := factories.NewTodoFactoryList(5)
	todos, err := todoApplication.All()
	if err != nil {
		panic(err)
	}

	assert.Contains(t, *todos, *todosFactory[0])
}

func TestTodoApplication_Create(t *testing.T) {
	todoFactory, err := factories.NewTodo()
	if err != nil {
		panic(err)
	}
	var todoTmp domain.Todo
	todo := todoFactory.ToDomain()
	todoTmp = *todo
	todoApplication := test.TodoApplicationInjector()
	newTodo, errNewTodo := todoApplication.Create(&todoTmp)
	if errNewTodo != nil {
		panic(errNewTodo)
	}

	assert.Equal(t, todo.Id(), newTodo.Id())
	assert.Equal(t, todo.Title, newTodo.Title)
	assert.Equal(t, todo.Body, newTodo.Body)
	assert.Equal(t, todo.Completed, newTodo.Completed)
}

func TestTodoApplication_Delete(t *testing.T) {
	todo := factories.NewTodoFactory()
	todoId := todo.Id()
	todoApplication := test.TodoApplicationInjector()
	errTodo := todoApplication.Delete(todo)
	if errTodo != nil {
		panic(errTodo)
	}
	todoRepository, _ := test.TodoRepository()
	deletedTodo, err := todoRepository.FindById(todoId)

	assert.NotNil(t, err)
	assert.Equal(t, err.HttpStatusCode(), http.StatusNotFound)
	assert.Nil(t, deletedTodo)
}

func TestTodoApplication_Show(t *testing.T) {
	todo := factories.NewTodoFactory()
	todoId := todo.Id()
	todoApplication := test.TodoApplicationInjector()
	showTodo, err := todoApplication.Show(todoId)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, showTodo)
	assert.Nil(t, err)
	assert.Equal(t, showTodo.Id(), todoId)
}

func TestTodoApplication_Toggle(t *testing.T) {
	todo := factories.NewTodoFactory()
	todoId := todo.Id()
	todoCompleted := *todo.Completed
	todoApplication := test.TodoApplicationInjector()
	toggledTodo, err := todoApplication.Toggle(todoId)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, toggledTodo)
	assert.Nil(t, err)
	assert.Equal(t, toggledTodo.Id(), todoId)
	assert.NotEqual(t, toggledTodo.Completed, todoCompleted)
}

func TestTodoApplication_Update(t *testing.T) {
	todo := factories.NewTodoFactory()
	todoId := todo.Id()
	newTitle := gofakeit.LoremIpsumWord()
	todo.Title = newTitle
	todoApplication := test.TodoApplicationInjector()
	updatedTodo, err := todoApplication.Update(todo)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, updatedTodo)
	assert.Nil(t, err)
	assert.Equal(t, updatedTodo.Id(), todoId)
	assert.Equal(t, updatedTodo.Title, newTitle)
}
