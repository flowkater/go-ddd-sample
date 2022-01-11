package user_domain

import (
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"gorm.io/gorm"
)

type User struct {
	Name  string              `json:"name"`
	Email string              `json:"email"`
	Todos []*todo_domain.Todo `json:"todos"`
	gorm.Model
}
