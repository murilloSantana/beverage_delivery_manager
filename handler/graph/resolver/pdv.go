package resolver

import (
	"beverage_delivery_manager/config/log"
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/pdv/domain"
	"context"
	"time"
)

func (r *mutationResolver) SavePdv(ctx context.Context, input model.PdvInput) (*domain.Pdv, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	r.Log.Info(log.LoggerFields{"input": input}, "input data save pdv")

	pdv := domain.Pdv{
		TradingName:  input.TradingName,
		OwnerName:    input.OwnerName,
		Document:     input.Document,
		CoverageArea: input.CoverageArea,
		Address:      input.Address,
	}

	newPdv, err := r.PdvUseCase.Save(ctx, pdv)
	if err != nil {
		r.Log.Error(log.LoggerFields{"error": err}, "error data save pdv")
		return nil, err
	}

	r.Log.Info(log.LoggerFields{"output": newPdv}, "output data save pdv")

	return newPdv, nil
}

func (r *queryResolver) FindPdvByID(_ context.Context, input model.PdvIDInput) (*domain.Pdv, error) {
	r.Log.Info(log.LoggerFields{"input": input}, "input data find pdv by id")

	pdv, err := r.PdvUseCase.FindByID(input.ID)
	if err != nil {
		r.Log.Error(log.LoggerFields{"error": err}, "error data find pdv by id")
		return nil, err
	}

	r.Log.Info(log.LoggerFields{"output": pdv}, "output data find pdv by id")

	return pdv, nil
}

func (r *queryResolver) FindPdvByAddress(_ context.Context, input model.PdvAddressInput) (*domain.Pdv, error) {
	r.Log.Info(log.LoggerFields{"input": input}, "input data find pdv by address")

	point := domain.Point{Type: "Point", Coordinates: domain.PointCoordinates{input.Longitude, input.Latitude}}

	pdv, err := r.PdvUseCase.FindByAddress(point)
	if err != nil {
		r.Log.Error(log.LoggerFields{"error": err}, "error data find pdv by address")
		return nil, err
	}

	r.Log.Info(log.LoggerFields{"output": pdv}, "output data find pdv by address")

	return pdv, nil
}
