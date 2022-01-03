package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Done        bool    `json:"done"`
	DueDate     *string `json:"due_date"`
}

func (t *Todo) SetDueDate(dueDate string) {
	t.DueDate = &dueDate
}

func (t *Todo) ToggleDone() {
	t.Done = !t.Done
}
