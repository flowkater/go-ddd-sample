package interfaces

import (
	todo_domain "github.com/flowkater/go-ddd-sample/src/domain/todo"
)

type (
	AddTodoRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		DueDate     *string `json:"due_date"`
	}

	UpdateTodoRequest struct {
		Name        *string `json:"name"`
		Description *string `json:"description"`
		DueDate     *string `json:"due_date"`
	}
)

func (r *AddTodoRequest) toCommand() *todo_domain.TodoCommandAddTodoRequest {
	return &todo_domain.TodoCommandAddTodoRequest{
		Name:        r.Name,
		Description: r.Description,
		DueDate:     r.DueDate,
	}
}

func (r *UpdateTodoRequest) toCommand(id uint) *todo_domain.TodoCommandUpdateTodoRequest {
	return &todo_domain.TodoCommandUpdateTodoRequest{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
		DueDate:     r.DueDate,
	}
}
