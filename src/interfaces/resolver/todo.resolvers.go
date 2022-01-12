package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/flowkater/go-ddd-sample/src/interfaces/mapper"
	"github.com/flowkater/go-ddd-sample/src/interfaces/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	todoInfo, err := r.todoFacade.AddTodo(ctx, input.ToCommand())
	if err != nil {
		return nil, err
	}

	return mapper.TodoOf(todoInfo), nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	todoInfo, err := r.todoFacade.UpdateTodo(ctx, input.ToCommand())
	if err != nil {
		return nil, err
	}

	return mapper.TodoOf(todoInfo), nil
}

func (r *mutationResolver) DoneTodo(ctx context.Context, id int) (*model.Todo, error) {
	todoInfo, err := r.todoFacade.DoneTodo(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	return mapper.TodoOf(todoInfo), nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
