package application

import domain "github.com/flowkater/go-ddd-sample/src/domain/todo"

type TodoFacade struct {
	todoService domain.TodoService
}

func (t *TodoFacade) AddTodo(command *domain.TodoCommand) (*domain.TodoInfo, error) {
	return t.todoService.AddTodo(command)
}
