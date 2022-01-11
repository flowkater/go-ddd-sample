// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package module

import (
	"github.com/flowkater/go-ddd-sample/src/application"
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"github.com/flowkater/go-ddd-sample/src/infrastructure/todo_infrastructure"
)

// Injectors from wire.go:

func InitializeTodoFacade() application.TodoFacade {
	todoRepository := todo_infrastructure.NewTodoRepository()
	todoReader := todo_infrastructure.NewTodoReader(todoRepository)
	todoStore := todo_infrastructure.NewTodoStore(todoRepository)
	todoExecutor := todo_infrastructure.NewTodoExecutor(todoRepository)
	todoService := todo_domain.NewTodoService(todoReader, todoStore, todoExecutor)
	todoFacade := application.NewTodoFacde(todoService)
	return todoFacade
}
