package generated

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"beverage_delivery_manager/graph/model"
	"context"
)

type Resolver struct{}

func (r *queryResolver) Pdvs(ctx context.Context) ([]model.Pdv, error) {
	panic("not implemented")
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
