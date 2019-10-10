package resolvers

import (
	"context"
	"self_gqlgen/models"
)

func (r *queryResolver) Players(ctx context.Context) ([]*models.Player, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreatePlayer(ctx context.Context, input models.RegisterPlayerReq) (*models.Player, error) {
	panic("not implemented")
}
