// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package module

import (
	"github.com/flowkater/go-ddd-sample/src/application"
	"github.com/flowkater/go-ddd-sample/src/domain/todo"
	"github.com/flowkater/go-ddd-sample/src/infrastructure/todo"
	"github.com/flowkater/go-ddd-sample/src/interfaces/todo"
)

// Injectors from wire.go:

func InitializeTodoApiController() *interfaces.TodoApiController {
	todoRepository := infrastructure.NewTodoRepository()
	todoReader := infrastructure.NewTodoReader(todoRepository)
	todoStore := infrastructure.NewTodoStore(todoRepository)
	todoExecutor := infrastructure.NewTodoExecutor(todoRepository)
	todoService := todo_domain.NewTodoService(todoReader, todoStore, todoExecutor)
	todoFacade := application.NewTodoFacde(todoService)
	todoApiController := interfaces.NewTodoApiController(todoFacade)
	return todoApiController
}
