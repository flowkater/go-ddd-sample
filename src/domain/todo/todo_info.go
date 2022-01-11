package todo_domain

type TodoInfo struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Done        bool    `json:"done"`
	DueDate     *string `json:"due_date"`
}
