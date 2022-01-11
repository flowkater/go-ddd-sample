package todo_domain

import (
	"gorm.io/gorm"
)

type Todo struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Done        bool    `json:"done"`
	DueDate     *string `json:"due_date"`
	UserID      uint
	gorm.Model
}

func (t *Todo) SetDueDate(dueDate string) {
	t.DueDate = &dueDate
}

func (t *Todo) ToggleDone() {
	t.Done = !t.Done
}

func (t *Todo) SetDescription(description string) {
	t.Description = description
}
