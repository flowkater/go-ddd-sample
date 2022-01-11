package todo_domain

type (
	TodoCommandAddTodoRequest struct {
		Name        string  `json:"name"`
		UserID      uint    `json:"user_id"`
		Description string  `json:"description"`
		DueDate     *string `json:"due_date"`
	}

	TodoCommandUpdateTodoRequest struct {
		ID          uint    `json:"id"`
		Name        *string `json:"name"`
		Description *string `json:"description"`
		DueDate     *string `json:"due_date"`
	}
)

func (t *TodoCommandAddTodoRequest) ToEntity() *Todo {
	return &Todo{
		Name:        t.Name,
		UserID:      t.UserID,
		Description: t.Description,
		DueDate:     t.DueDate,
	}
}

func (t *TodoCommandUpdateTodoRequest) ToEntity(todo *Todo) *Todo {
	if t.Name != nil {
		todo.Name = *t.Name
	}

	if t.Description != nil {
		todo.Description = *t.Description
	}

	if t.DueDate != nil {
		todo.DueDate = t.DueDate
	}

	return todo
}
