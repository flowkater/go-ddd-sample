package todo_infrastructure

import (
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"gorm.io/gorm"
)

type todoReader struct {
	todoRepository TodoRepository
}

func NewTodoReader(todoRepository TodoRepository) todo_domain.TodoReader {
	return &todoReader{
		todoRepository: todoRepository,
	}
}

func (t *todoReader) GetTodo(db *gorm.DB, id uint) (*todo_domain.Todo, error) {
	return t.todoRepository.FindById(db, id)
}
