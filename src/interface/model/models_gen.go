// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Name        string  `json:"name"`
	UserID      string  `json:"userId"`
	DueDate     *string `json:"dueDate"`
	Description *string `json:"description"`
}

type Todo struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Done        bool    `json:"done"`
	DueDate     *string `json:"dueDate"`
	Description *string `json:"description"`
	User        *User   `json:"user"`
}

type UpdateTodo struct {
	Name        *string `json:"name"`
	DueDate     *string `json:"dueDate"`
	Description *string `json:"description"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserInput struct {
	Name string `json:"name"`
}