package domain

import (
	"context"

	"github.com/flowkater/go-ddd-sample/src/config"
)

type todoService struct {
	todoReader   TodoReader
	todoStore    TodoStore
	todoExecutor TodoExecutor
}

func NewTodoService(
	todoReader TodoReader,
	todoStore TodoStore,
	todoExecutor TodoExecutor,
) TodoService {
	return &todoService{
		todoReader:   todoReader,
		todoStore:    todoStore,
		todoExecutor: todoExecutor,
	}
}

func (t *todoService) AddTodo(ctx context.Context, command *TodoCommandAddTodoRequest) (*TodoInfo, error) {
	db := config.DBWithContext(ctx)

	initTodo := command.ToEntity()
	todo, err := t.todoStore.Store(db, initTodo)
	if err != nil {
		return nil, err
	}

	return &TodoInfo{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Done:        todo.Done,
		DueDate:     todo.DueDate,
	}, nil
}

func (t *todoService) UpdateTodo(ctx context.Context, command *TodoCommandUpdateTodoRequest) (*TodoInfo, error) {
	db := config.DBWithContext(ctx)

	getTodo, err := t.todoReader.GetTodo(db, command.ID)
	if err != nil {
		return nil, err
	}

	updateTodo := command.ToEntity(getTodo)

	todo, err := t.todoExecutor.Update(db, updateTodo)
	if err != nil {
		return nil, err
	}

	return &TodoInfo{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Done:        todo.Done,
		DueDate:     todo.DueDate,
	}, nil
}

func (t *todoService) DoneTodo(ctx context.Context, id uint) (*TodoInfo, error) {
	db := config.DBWithContext(ctx)

	getTodo, err := t.todoReader.GetTodo(db, id)
	if err != nil {
		return nil, err
	}

	todo, err := t.todoExecutor.UpdateDone(db, getTodo)
	if err != nil {
		return nil, err
	}

	return &TodoInfo{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Done:        todo.Done,
		DueDate:     todo.DueDate,
	}, nil
}
