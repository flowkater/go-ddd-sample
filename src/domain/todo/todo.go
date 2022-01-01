package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Done        bool    `json:"done"`
	DueDate     *string `json:"due_date"`
}
