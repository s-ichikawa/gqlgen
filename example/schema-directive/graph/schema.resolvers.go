package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/example/schema-directive/graph/generated"
	"github.com/99designs/gqlgen/example/schema-directive/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	t := &model.Todo{
		ID:   fmt.Sprintf("%d", len(r.todos)+1),
		Text: input.Text,
		Done: false,
		User: nil,
	}
	r.todos = append(r.todos, t)
	return t, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*bool, error) {
	delete(r.users, userID)
	res := true
	return &res, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.users[id], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
