package interfaces

import domain "github.com/flowkater/go-ddd-sample/src/domain/todo"

type (
	AddTodoRequest struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
		DueDate     *string `json:"due_date"`
	}
)

func (r *AddTodoRequest) toCommand() *domain.TodoCommand {
	return &domain.TodoCommand{
		Name:        r.Name,
		Description: *r.Description,
		DueDate:     r.DueDate,
	}
}
