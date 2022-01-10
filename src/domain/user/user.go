package domain

import (
	domain "github.com/flowkater/go-ddd-sample/src/domain/todo"
	"gorm.io/gorm"
)

type User struct {
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Todos []*domain.Todo `json:"todos"`
	gorm.Model
}
