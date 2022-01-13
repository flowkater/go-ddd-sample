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

func (t *todoReader) FindTodoInfoAllByUserId(db *gorm.DB, userId uint) ([]*todo_domain.TodoInfo, error) {
	todos, err := t.todoRepository.FindAllByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	var todoInfos []*todo_domain.TodoInfo
	for _, todo := range todos {
		todoInfos = append(todoInfos, &todo_domain.TodoInfo{
			ID:          todo.ID,
			Name:        todo.Name,
			Description: todo.Description,
			Done:        todo.Done,
			DueDate:     todo.DueDate,
			UserID:      todo.UserID,
		})
	}

	return todoInfos, nil
}

func (t *todoReader) GetTodoById(db *gorm.DB, id uint) (*todo_domain.Todo, error) {
	return t.todoRepository.GetById(db, id)
}
