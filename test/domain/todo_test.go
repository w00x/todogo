package domain

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"todogo/domain"
	"todogo/shared"
	"todogo/test/factories"
)

func TestNewTodo(t *testing.T) {
	title := gofakeit.LoremIpsumWord()
	body := gofakeit.LoremIpsumSentence(500)
	completed := gofakeit.Bool()
	createdAt := gofakeit.Date()

	todo := domain.NewTodo("", title, body, &completed, shared.TimeToDateTime(createdAt))

	assert.Equal(t, todo.Title, title)
	assert.Equal(t, todo.Body, body)
	assert.Equal(t, *todo.Completed, completed)
	assert.Equal(t, todo.CreatedAt, createdAt)
	assert.NotNil(t, todo.Id())
	assert.Regexp(t, "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", todo.Id())
}

func TestTodo_Id(t *testing.T) {
	todoFactory, err := factories.NewTodo()
	if err != nil {
		fmt.Println(err)
	}
	todo := todoFactory.ToDomain()

	assert.NotNil(t, todo.Id())
	assert.Regexp(t, "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", todo.Id())
}

func TestTodo_Toggle(t *testing.T) {
	todoFactory, err := factories.NewTodo()
	if err != nil {
		fmt.Println(err)
	}
	todo := todoFactory.ToDomain()
	var completed bool
	completed = *todo.Completed
	todo.Toggle()
	assert.Equal(t, !completed, *todo.Completed)
}
