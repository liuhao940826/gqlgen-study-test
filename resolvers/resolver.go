package resolvers

import (
	"context"
	"self_gqlgen/gqlgen"
	"self_gqlgen/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePlayer(ctx context.Context, input models.RegisterUser) (*models.Player, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}
func (r *queryResolver) Players(ctx context.Context) ([]*models.Player, error) {
	panic("not implemented")
}
