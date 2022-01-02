package domain

type TodoService interface {
	AddTodo(command *TodoCommand) (*TodoInfo, error)
	DoneTodo(id uint) (*TodoInfo, error)
}
