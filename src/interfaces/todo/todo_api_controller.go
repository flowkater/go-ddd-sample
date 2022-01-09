package interfaces

import (
	"net/http"
	"strconv"

	"github.com/flowkater/go-ddd-sample/src/application"
	"github.com/labstack/echo/v4"
)

type TodoApiController struct {
	todoFacade application.TodoFacade
}

func NewTodoApiController(todoFacade application.TodoFacade) *TodoApiController {
	return &TodoApiController{
		todoFacade: todoFacade,
	}
}

func (t *TodoApiController) AddTodo(c echo.Context) error {
	request := new(AddTodoRequest)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	todoInfo, err := t.todoFacade.AddTodo(c.Request().Context(), request.toCommand())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, todoInfo)
}

func (t *TodoApiController) UpdateTodo(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	request := new(UpdateTodoRequest)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	todoInfo, err := t.todoFacade.UpdateTodo(c.Request().Context(), request.toCommand(uint(id)))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, todoInfo)
}

func (t *TodoApiController) DoneTodo(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseUint(paramId, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	todoInfo, err := t.todoFacade.DoneTodo(c.Request().Context(), uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, todoInfo)
}

func RegiterHandlerTodos(e *echo.Echo, todoApiController *TodoApiController) {
	e.POST("/todos", todoApiController.AddTodo)
	e.PATCH("/todos/:id", todoApiController.UpdateTodo)
	e.PATCH("/todos/:id/done", todoApiController.DoneTodo)
}
