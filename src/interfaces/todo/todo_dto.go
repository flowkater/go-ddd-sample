package interfaces

import domain "github.com/flowkater/go-ddd-sample/src/domain/todo"

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

func (r *AddTodoRequest) toCommand() *domain.TodoCommandAddTodoRequest {
	return &domain.TodoCommandAddTodoRequest{
		Name:        r.Name,
		Description: r.Description,
		DueDate:     r.DueDate,
	}
}

func (r *UpdateTodoRequest) toCommand(id uint) *domain.TodoCommandUpdateTodoRequest {
	return &domain.TodoCommandUpdateTodoRequest{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
		DueDate:     r.DueDate,
	}
}
