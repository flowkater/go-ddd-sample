package application

import (
	"context"

	domain "github.com/flowkater/go-ddd-sample/src/domain/todo"
)

type TodoFacade struct {
	todoService domain.TodoService
}

func NewTodoFacde(todoService domain.TodoService) TodoFacade {
	return TodoFacade{
		todoService: todoService,
	}
}

func (t *TodoFacade) AddTodo(ctx context.Context, command *domain.TodoCommandAddTodoRequest) (*domain.TodoInfo, error) {
	return t.todoService.AddTodo(ctx, command)
}

func (t *TodoFacade) UpdateTodo(ctx context.Context, command *domain.TodoCommandUpdateTodoRequest) (*domain.TodoInfo, error) {
	return t.todoService.UpdateTodo(ctx, command)
}

func (t *TodoFacade) DoneTodo(ctx context.Context, id uint) (*domain.TodoInfo, error) {
	return t.todoService.DoneTodo(ctx, id)
}
