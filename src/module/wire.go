//go:build wireinject
// +build wireinject

package module

import (
	"github.com/flowkater/go-ddd-sample/src/application"
	domain "github.com/flowkater/go-ddd-sample/src/domain/todo"
	infrastructure "github.com/flowkater/go-ddd-sample/src/infrastructure/todo"
	interfaces "github.com/flowkater/go-ddd-sample/src/interfaces/todo"
	"github.com/google/wire"
)

func InitializeTodoApiController() *interfaces.TodoApiController {
	wire.Build(
		interfaces.NewTodoApiController,
		application.NewTodoFacde,
		domain.NewTodoService,
		wire.NewSet(infrastructure.NewTodoReader, infrastructure.NewTodoStore, infrastructure.NewTodoExecutor),
		infrastructure.NewTodoRepository,
	)

	return &interfaces.TodoApiController{}
}
