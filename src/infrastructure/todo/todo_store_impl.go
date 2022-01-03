package infrastructure

import (
	domain "github.com/flowkater/go-ddd-sample/src/domain/todo"
	"gorm.io/gorm"
)

type todoStore struct {
	todoRepository TodoRepository
}

func NewTodoStore(todoRepository TodoRepository) domain.TodoStore {
	return &todoStore{
		todoRepository: todoRepository,
	}
}

func (t *todoStore) Store(db *gorm.DB, todo *domain.Todo) (*domain.Todo, error) {
	return t.todoRepository.Create(db, todo)
}
