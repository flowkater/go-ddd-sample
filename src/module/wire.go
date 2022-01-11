//go:build wireinject
// +build wireinject

package module

import (
	"github.com/flowkater/go-ddd-sample/src/application"
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"github.com/flowkater/go-ddd-sample/src/infrastructure/todo_infrastructure"
	interfaces "github.com/flowkater/go-ddd-sample/src/interfaces/todo"
	"github.com/google/wire"
)

func InitializeTodoApiController() *interfaces.TodoApiController {
	wire.Build(
		interfaces.NewTodoApiController,
		application.NewTodoFacde,
		todo_domain.NewTodoService,
		wire.NewSet(todo_infrastructure.NewTodoReader, todo_infrastructure.NewTodoStore, todo_infrastructure.NewTodoExecutor),
		todo_infrastructure.NewTodoRepository,
	)

	return &interfaces.TodoApiController{}
}
