package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"todogo/application"
	"todogo/infrastructure"
	v1 "todogo/infrastructure/api/controller/v1"
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

func TestNewTodoController(t *testing.T) {
	factory, err := repository.TodoFactory(test.GetCurrentAdapter())
	if err != nil {
		panic(err)
	}
	todoApplication := application.NewTodoApplication(factory)
	todoController := v1.NewTodoController(todoApplication)

	assert.NotNil(t, todoController)
}

func TestTodoController_Create(t *testing.T) {
	router := infrastructure.GinRoutes(test.GetCurrentAdapter())
	server := httptest.NewServer(router)

	title := gofakeit.LoremIpsumWord()
	body := gofakeit.LoremIpsumSentence(100)
	completed := gofakeit.Bool()
	values := map[string]interface{}{"title": title, "body": body, "completed": completed}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Post(fmt.Sprintf("%s/v1/todo", server.URL), "application/json", bytes.NewBuffer(jsonData))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	todoRepository, _ := test.TodoRepository()
	todo, errFindTodo := todoRepository.FindById(response["id"].(string))

	assert.Nil(t, errFindTodo)
	assert.Equal(t, todo.Id(), response["id"])
	assert.Equal(t, title, response["title"])
	assert.Equal(t, body, response["body"])
	assert.Equal(t, completed, response["completed"])
}

func TestTodoController_Delete(t *testing.T) {
	todo := factories.NewTodoFactory()
	router := infrastructure.GinRoutes(test.GetCurrentAdapter())
	server := httptest.NewServer(router)
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/todo/%s", server.URL, todo.Id()), nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	todoRepository, _ := test.TodoRepository()
	_, errFindTodo := todoRepository.FindById(todo.Id())
	assert.NotNil(t, errFindTodo)
	assert.Equal(t, http.StatusNotFound, errFindTodo.HttpStatusCode())
}

func TestTodoController_Get(t *testing.T) {
	todo := factories.NewTodoFactory()
	router := infrastructure.GinRoutes(test.GetCurrentAdapter())
	server := httptest.NewServer(router)
	resp, _ := http.Get(fmt.Sprintf("%s/v1/todo/%s", server.URL, todo.Id()))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, todo.Id(), response["id"])
}

func TestTodoController_Index(t *testing.T) {
	sizeOfTodos := 5
	todos := factories.NewTodoFactoryList(sizeOfTodos)
	router := infrastructure.GinRoutes(test.GetCurrentAdapter())
	server := httptest.NewServer(router)

	resp, _ := http.Get(fmt.Sprintf("%s/v1/todo", server.URL))

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response []gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)
	var ids []string

	for _, responseTodo := range response {
		ids = append(ids, responseTodo["id"].(string))
	}

	assert.Contains(t, ids, todos[0].Id())
}

func TestTodoController_Toggle(t *testing.T) {
	todo := factories.NewTodoFactory()
	todoCompleted := *todo.Completed
	router := infrastructure.GinRoutes(test.GetCurrentAdapter())
	server := httptest.NewServer(router)
	resp, _ := http.Get(fmt.Sprintf("%s/v1/todo/toggle/%s", server.URL, todo.Id()))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, !todoCompleted, response["completed"])
}

func TestTodoController_Update(t *testing.T) {
	todo := factories.NewTodoFactory()
	router := infrastructure.GinRoutes(test.GetCurrentAdapter())
	server := httptest.NewServer(router)

	values := factories.NewTodoObjectFactory()
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v1/todo/%s", server.URL, todo.Id()), bytes.NewBuffer(jsonData))
	client := &http.Client{}
	resp, _ := client.Do(req)

	var response gin.H
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &response)

	assert.Equal(t, values["title"], response["title"])
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
