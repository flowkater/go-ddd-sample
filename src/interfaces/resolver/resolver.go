package resolver

import "github.com/flowkater/go-ddd-sample/src/application"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todoFacade application.TodoFacade
	userFacade application.UserFacade
}

func NewResolver(todoFacade application.TodoFacade, userFacade application.UserFacade) *Resolver {
	return &Resolver{
		todoFacade: todoFacade,
		userFacade: userFacade,
	}
}
