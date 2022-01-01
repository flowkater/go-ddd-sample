package domain

type TodoService interface {
	AddTodo(command *TodoCommand) (todoInfo *TodoInfo, err error)
	DoneTodo(id uint) (todoInfo *TodoInfo, err error)
}
