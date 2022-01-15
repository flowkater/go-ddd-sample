package application

import (
	"context"

	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
)

type TodoFacade struct {
	todoService todo_domain.TodoService
}

func NewTodoFacade(todoService todo_domain.TodoService) TodoFacade {
	return TodoFacade{
		todoService: todoService,
	}
}

func (t *TodoFacade) FindTodoAllByUserId(ctx context.Context, userID uint) ([]*todo_domain.TodoInfo, error) {
	return t.todoService.FindTodoAllByUserId(ctx, userID)
}

func (t *TodoFacade) GetTodoById(ctx context.Context, id uint) (*todo_domain.TodoInfo, error) {
	return t.todoService.GetTodoById(ctx, id)
}

func (t *TodoFacade) AddTodo(ctx context.Context, command *todo_domain.TodoCommandAddTodoRequest) (*todo_domain.TodoInfo, error) {
	return t.todoService.AddTodo(ctx, command)
}

func (t *TodoFacade) UpdateTodo(ctx context.Context, command *todo_domain.TodoCommandUpdateTodoRequest) (*todo_domain.TodoInfo, error) {
	return t.todoService.UpdateTodo(ctx, command)
}

func (t *TodoFacade) DoneTodo(ctx context.Context, id uint) (*todo_domain.TodoInfo, error) {
	return t.todoService.DoneTodo(ctx, id)
}
