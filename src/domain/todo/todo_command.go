package domain

type TodoCommand struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	DueDate     *string `json:"due_date"`
}

func (t *TodoCommand) ToEntity() *Todo {
	return &Todo{
		Name:        t.Name,
		Description: t.Description,
		DueDate:     t.DueDate,
	}
}
