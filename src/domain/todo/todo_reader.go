package domain

import "gorm.io/gorm"

type TodoReader interface {
	GetTodo(db *gorm.DB, id uint) (*Todo, error)
}
