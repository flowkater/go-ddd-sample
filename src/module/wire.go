//go:build wireinject
// +build wireinject

package module

import (
	"github.com/flowkater/go-ddd-sample/src/application"
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
	"github.com/flowkater/go-ddd-sample/src/infrastructure/todo_infrastructure"
	"github.com/flowkater/go-ddd-sample/src/infrastructure/user_infrastructure"
	"github.com/google/wire"
)

func InitializeTodoFacade() application.TodoFacade {
	wire.Build(
		application.NewTodoFacade,
		todo_domain.NewTodoService,
		wire.NewSet(todo_infrastructure.NewTodoReader, todo_infrastructure.NewTodoStore, todo_infrastructure.NewTodoExecutor),
		todo_infrastructure.NewTodoRepository,
	)

	return application.TodoFacade{}
}

func InitializeUserFacade() application.UserFacade {
	wire.Build(
		application.NewUserFacade,
		user_domain.NewUserService,
		user_infrastructure.NewUserReader,
		user_infrastructure.NewUserRepository,
	)

	return application.UserFacade{}
}
