package resolver

import (
	"beverage_delivery_manager/handler/graph/generated"
	"beverage_delivery_manager/pdv/usecase"
)

type Resolver struct {
	PdvUseCase usecase.PdvUseCase
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
