package todo_domain

import (
	"gorm.io/gorm"
)

type TodoExecutor interface {
	UpdateDone(db *gorm.DB, todo *Todo) (*Todo, error)
	Update(db *gorm.DB, todo *Todo) (*Todo, error)
}
