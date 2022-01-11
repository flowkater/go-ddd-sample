package todo_domain

import "context"

type TodoService interface {
	AddTodo(ctx context.Context, command *TodoCommandAddTodoRequest) (*TodoInfo, error)
	UpdateTodo(ctx context.Context, command *TodoCommandUpdateTodoRequest) (*TodoInfo, error)
	DoneTodo(ictx context.Context, id uint) (*TodoInfo, error)
}
