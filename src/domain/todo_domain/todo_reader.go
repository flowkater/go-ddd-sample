package todo_domain

import (
	"gorm.io/gorm"
)

type TodoReader interface {
	FindTodoInfoAllByUserId(db *gorm.DB, userId uint) ([]*TodoInfo, error)
	GetTodoById(db *gorm.DB, id uint) (*Todo, error)
}
