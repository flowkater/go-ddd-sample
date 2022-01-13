package mapper

import (
	"github.com/flowkater/go-ddd-sample/src/domain/todo_domain"
	"github.com/flowkater/go-ddd-sample/src/interfaces/model"
)

func TodoOf(todoInfo *todo_domain.TodoInfo) *model.Todo {
	return &model.Todo{
		ID:          int(todoInfo.ID),
		Name:        todoInfo.Name,
		Description: &todoInfo.Description,
		DueDate:     todoInfo.DueDate,
		Done:        todoInfo.Done,
		UserID:      int(todoInfo.UserID),
	}
}

func TodosOf(todoInfos []*todo_domain.TodoInfo) []*model.Todo {
	todos := make([]*model.Todo, len(todoInfos))
	for i, todo := range todoInfos {
		todos[i] = TodoOf(todo)
	}

	return todos
}
