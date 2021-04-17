package generated

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"beverage_delivery_manager/domain"
	"context"
)

type Resolver struct{}

func (r *pDVResolver) TradingName(ctx context.Context, obj *domain.PDV) (string, error) {
	panic("not implemented")
}

func (r *queryResolver) Pdvs(ctx context.Context, coverageArea *domain.MultiPolygon) ([]domain.PDV, error) {
	panic("not implemented")
}

// PDV returns PDVResolver implementation.
func (r *Resolver) PDV() PDVResolver { return &pDVResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type pDVResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
