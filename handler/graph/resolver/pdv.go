package resolver

import (
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/pdv/domain"
	"context"
	"time"
)

func (r *mutationResolver) SavePdv(ctx context.Context, input model.PdvInput) (*domain.Pdv, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pdv := domain.Pdv{
		TradingName:  input.TradingName,
		OwnerName:    input.OwnerName,
		Document:     input.Document,
		CoverageArea: input.CoverageArea,
		Address:      input.Address,
	}

	newPdv, err := r.PdvUseCase.Save(ctx, pdv)
	if err != nil {
		return nil, err
	}

	return &newPdv, nil
}

func (r *queryResolver) FindPdvByID(_ context.Context, input model.PdvIDInput) (*domain.Pdv, error) {
	pdv, err := r.PdvUseCase.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &pdv, nil
}

func (r *queryResolver) FindPdvByAddress(_ context.Context, input model.PdvAddressInput) (*domain.Pdv, error) {
	point := domain.Point{Type: "Point", Coordinates: domain.PointCoordinates{input.Longitude, input.Latitude}}

	pdv, err := r.PdvUseCase.FindByAddress(point)
	if err != nil {
		return nil, err
	}

	return &pdv, nil
}
