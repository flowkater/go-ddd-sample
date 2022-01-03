package domain

import "gorm.io/gorm"

type TodoExecutor interface {
	UpdateDone(db *gorm.DB, todo *Todo) (*Todo, error)
	UpdateDueDate(db *gorm.DB, dueDate string, todo *Todo) (*Todo, error)
}
