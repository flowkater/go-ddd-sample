package model

import "github.com/flowkater/go-ddd-sample/src/domain/todo_domain"

func (i *CreateTodoInput) ToCommand() *todo_domain.TodoCommandAddTodoRequest {
	todoCommand := &todo_domain.TodoCommandAddTodoRequest{
		Name:    i.Name,
		UserID:  uint(i.UserID),
		DueDate: i.DueDate,
	}

	if i.Description != nil {
		todoCommand.Description = *i.Description
	}

	return todoCommand
}

func (i *UpdateTodoInput) ToCommand() *todo_domain.TodoCommandUpdateTodoRequest {
	return &todo_domain.TodoCommandUpdateTodoRequest{
		ID:          uint(i.ID),
		Name:        i.Name,
		Description: i.Description,
		DueDate:     i.DueDate,
	}
}
