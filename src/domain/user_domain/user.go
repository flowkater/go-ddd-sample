package user_domain

import (
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"gorm.io/gorm"
)

type User struct {
	ID    uint                `json:"id"`
	Name  string              `json:"name"`
	Todos []*todo_domain.Todo `json:"todos"`
	gorm.Model
}
