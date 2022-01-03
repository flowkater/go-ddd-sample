package domain

import "context"

type TodoService interface {
	AddTodo(ctx context.Context, command *TodoCommand) (*TodoInfo, error)
	UpdateDueDateTodo(ctx context.Context, dueDate string, id uint) (*TodoInfo, error)
	DoneTodo(ictx context.Context, id uint) (*TodoInfo, error)
}
