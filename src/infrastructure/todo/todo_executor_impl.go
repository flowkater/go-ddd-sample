package infrastructure

import (
	domain "github.com/flowkater/go-ddd-sample/src/domain/todo"
	"gorm.io/gorm"
)

type todoExecutor struct {
	todoRepository TodoRepository
}

func NewTodoExecutor(todoRepository TodoRepository) domain.TodoExecutor {
	return &todoExecutor{
		todoRepository: todoRepository,
	}
}

func (t *todoExecutor) UpdateDueDate(db *gorm.DB, dueDate string, todo *domain.Todo) (*domain.Todo, error) {
	todo.SetDueDate(dueDate)
	return t.todoRepository.Update(db, todo)
}

func (t *todoExecutor) UpdateDone(db *gorm.DB, todo *domain.Todo) (*domain.Todo, error) {
	todo.ToggleDone()
	return t.todoRepository.UpdateDone(db, todo)
}
