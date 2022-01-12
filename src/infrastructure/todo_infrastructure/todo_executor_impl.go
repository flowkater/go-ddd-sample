package todo_infrastructure

import (
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"gorm.io/gorm"
)

type todoExecutor struct {
	todoRepository TodoRepository
}

func NewTodoExecutor(todoRepository TodoRepository) todo_domain.TodoExecutor {
	return &todoExecutor{
		todoRepository: todoRepository,
	}
}

func (t *todoExecutor) Update(db *gorm.DB, todo *todo_domain.Todo) (*todo_domain.Todo, error) {
	return t.todoRepository.Update(db, todo)
}

func (t *todoExecutor) UpdateDone(db *gorm.DB, todo *todo_domain.Todo) (*todo_domain.Todo, error) {
	todo.ToggleDone()

	return t.todoRepository.UpdateDone(db, todo)
}
