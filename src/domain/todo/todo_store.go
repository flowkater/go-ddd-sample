package domain

type TodoStore interface {
	Store(todo *Todo) (*Todo, error)
}
