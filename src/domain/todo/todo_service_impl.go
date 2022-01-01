package domain

import "fmt"

type todoService struct{}

func NewTodoService() TodoService {
	return &todoService{}
}

func (t *todoService) AddTodo(command *TodoCommand) (todoInfo *TodoInfo, err error) {
	todo := command.ToEntity()

	fmt.Println(todo)
	// if err := db.Create(todo).Error; err != nil {
	// 	return nil, err
	// }
	return &TodoInfo{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
		Done:        todo.Done,
		DueDate:     todo.DueDate,
	}, nil
}

func (t *todoService) DoneTodo(id uint) (todoInfo *TodoInfo, err error) {
	// todo := &Todo{}
	// if err := db.First(todo, id).Error; err != nil {
	// 	return nil, err
	// }
	// todo.Done = true
	// if err := db.Save(todo).Error; err != nil {
	// 	return nil, err
	// }
	return &TodoInfo{
		ID:          id,
		Name:        "todo",
		Description: "todo",
		Done:        true,
		DueDate:     nil,
	}, nil
}
