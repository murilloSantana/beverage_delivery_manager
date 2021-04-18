package resolver

import (
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/pdv/domain"
	"context"
)

func (r *mutationResolver) SavePdv(_ context.Context, input model.PdvInput) (*domain.Pdv, error) {
	pdv := domain.Pdv{
		TradingName:  input.TradingName,
		OwnerName:    input.OwnerName,
		Document:     input.Document,
		CoverageArea: input.CoverageArea,
		Address:      input.Address,
	}

	newPdv, err := r.PdvUseCase.Save(pdv)
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
	coordinates := domain.PointCoordinates{input.Longitude, input.Latitude}

	pdv, err := r.PdvUseCase.FindByAddress(coordinates)
	if err != nil {
		return nil, err
	}

	return &pdv, nil
}
