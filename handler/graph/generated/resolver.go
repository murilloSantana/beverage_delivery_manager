package generated

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/pdv/domain"
	"context"
)

type Resolver struct{}

func (r *mutationResolver) SavePdv(ctx context.Context, input model.PdvInput) (*domain.Pdv, error) {
	panic("not implemented")
}

func (r *queryResolver) FindPdvByID(ctx context.Context, input model.PdvIDInput) (*domain.Pdv, error) {
	panic("not implemented")
}

func (r *queryResolver) FindPdvByAddress(ctx context.Context, input model.PdvAddressInput) (*domain.Pdv, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
