package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

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

func (r *queryResolver) Todos(ctx context.Context, userID int) ([]*model.Todo, error) {
	todoInfos, err := r.todoFacade.FindTodoAllByUserId(ctx, uint(userID))
	if err != nil {
		return nil, err
	}

	return mapper.TodosOf(todoInfos), nil
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	todoInfo, err := r.todoFacade.GetTodoById(ctx, uint(id))
	if err != nil {
		return nil, err
	}
	return mapper.TodoOf(todoInfo), nil
}

func (r *todoResolver) User(ctx context.Context, todo *model.Todo) (*model.User, error) {
	return r.userFacade.NewUserLoader(ctx).Load(uint(todo.UserID))

	// userInfo, err := r.userFacade.GetUserById(ctx, uint(todo.UserID))
	// if err != nil {
	// 	return nil, err
	// }

	// return mapper.UserOf(userInfo), nil
}

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
