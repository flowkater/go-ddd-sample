//go:build wireinject
// +build wireinject

package module

import (
	"github.com/flowkater/go-ddd-sample/src/application"
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"github.com/flowkater/go-ddd-sample/src/infrastructure/todo_infrastructure"
	"github.com/google/wire"
)

func InitializeTodoFacade() application.TodoFacade {
	wire.Build(
		application.NewTodoFacde,
		todo_domain.NewTodoService,
		wire.NewSet(todo_infrastructure.NewTodoReader, todo_infrastructure.NewTodoStore, todo_infrastructure.NewTodoExecutor),
		todo_infrastructure.NewTodoRepository,
	)

	return application.TodoFacade{}
}
