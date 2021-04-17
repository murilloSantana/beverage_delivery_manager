package resolver

import (
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/pdv/domain"
	"context"
)

func (r *mutationResolver) SavePdv(ctx context.Context, input *model.PdvInput) (*domain.Pdv, error) {
	panic("not implemented")
}

func (r *queryResolver) FindPdvByID(ctx context.Context, input *model.PdvIDInput) (*domain.Pdv, error) {
	panic("not implemented")
}

func (r *queryResolver) FindPdvByAddress(ctx context.Context, input *model.PdvAddressInput) (*domain.Pdv, error) {
	panic("not implemented")
}
