package resolver

import "github.com/flowkater/go-ddd-sample/src/application"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todoFacade application.TodoFacade
}

func NewResolver(todoFacade application.TodoFacade) *Resolver {
	return &Resolver{todoFacade: todoFacade}
}
