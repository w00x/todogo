package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todogohexa/application"
	"todogohexa/infrastructure/api/context"
	"todogohexa/infrastructure/dto"
)

type TodoController struct {
	todoApplication *application.TodoApplication
}

func NewTodoController(todoApplication *application.TodoApplication) *TodoController {
	return &TodoController{todoApplication}
}

func (todoController *TodoController) Index(c context.IContextAdapter) {
	todos, _ := todoController.todoApplication.All()
	c.JSON(http.StatusOK, dto.NewTodoListDtoFromDomains(todos))
}

func (todoController *TodoController) Get(c context.IContextAdapter) {
	id := c.Param("id")
	todo, err := todoController.todoApplication.Show(id)

	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewTodoDtoFromDomain(todo))
}

func (todoController *TodoController) Create(c context.IContextAdapter) {
	var todoDto dto.TodoDto
	if err := c.ShouldBindJSON(&todoDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	todo, err := todoController.todoApplication.Create(todoDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{ "error": err.Error() })
		return
	}
	c.JSON(http.StatusCreated, dto.NewTodoDtoFromDomain(todo))
}

func (todoController *TodoController) Update(c context.IContextAdapter) {
	var todoDto dto.TodoDto
	if err := c.ShouldBindUri(&todoDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&todoDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	todo, err := todoController.todoApplication.Update(todoDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.NewTodoDtoFromDomain(todo))
}

func (todoController *TodoController) Delete(c context.IContextAdapter) {
	var todoDto dto.TodoDto
	if err := c.ShouldBindUri(&todoDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := todoController.todoApplication.Delete(todoDto.ToDomain())
	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (todoController *TodoController) Toggle(c context.IContextAdapter) {
	id := c.Param("id")
	todo, err := todoController.todoApplication.Toggle(id)

	if err != nil {
		c.JSON(err.HttpStatusCode(), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewTodoDtoFromDomain(todo))
}