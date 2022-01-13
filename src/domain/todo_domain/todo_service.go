package todo_domain

import "context"

type TodoService interface {
	AddTodo(ctx context.Context, command *TodoCommandAddTodoRequest) (*TodoInfo, error)
	UpdateTodo(ctx context.Context, command *TodoCommandUpdateTodoRequest) (*TodoInfo, error)
	DoneTodo(ctx context.Context, id uint) (*TodoInfo, error)
	FindTodoAllByUserId(ctx context.Context, userId uint) ([]*TodoInfo, error)
	GetTodoById(ctx context.Context, id uint) (*TodoInfo, error)
}
