package todo_infrastructure

import (
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"gorm.io/gorm"
)

type todoStore struct {
	todoRepository TodoRepository
}

func NewTodoStore(todoRepository TodoRepository) todo_domain.TodoStore {
	return &todoStore{
		todoRepository: todoRepository,
	}
}

func (t *todoStore) Store(db *gorm.DB, todo *todo_domain.Todo) (*todo_domain.Todo, error) {
	return t.todoRepository.Create(db, todo)
}
