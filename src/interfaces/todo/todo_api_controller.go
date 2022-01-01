package interfaces

import (
	"net/http"

	"github.com/flowkater/go-ddd-sample/src/application"
	"github.com/labstack/echo/v4"
)

type TodoApiController struct {
	todoFacade application.TodoFacade
}

func (t *TodoApiController) AddTodo(c echo.Context) error {
	request := new(AddTodoRequest)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	command := request.toCommand()
	todoInfo, err := t.todoFacade.AddTodo(command)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, todoInfo)
}

func RegiterHandlerTodos(e *echo.Echo) {
	todoApiController := &TodoApiController{}
	e.POST("/todos", todoApiController.AddTodo)
}
