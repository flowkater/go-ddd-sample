package domain

import "gorm.io/gorm"

type TodoStore interface {
	Store(db *gorm.DB, todo *Todo) (*Todo, error)
}
