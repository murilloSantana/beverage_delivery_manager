package helper

import (
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/pdv/domain"
	"time"
)

type DefaultPdvOption func(*domain.Pdv)

func WithID(ID string) DefaultPdvOption {
	return func(pdv *domain.Pdv) {
		pdv.ID = ID
	}
}

func WithDocument(document string) DefaultPdvOption {
	return func(pdv *domain.Pdv) {
		pdv.Document = document
	}
}

func WithAddress(address ...float64) DefaultPdvOption {
	return func(pdv *domain.Pdv) {
		pdv.Address = domain.Point{
			Type:        "Point",
			Coordinates: address,
		}
	}
}

func NewPdv(opts ...DefaultPdvOption) *domain.Pdv {
	pdv := domain.Pdv{
		TradingName: "Mercado Pinheiros",
		OwnerName:   "Luiz Santo",
		Document:    time.Now().String(),
		CoverageArea: domain.MultiPolygon{
			Type: "MultiPolygon",
			Coordinates: [][][][2]float64{{{{-46.623238, -21.785538}, {-46.607616, -21.819803}, {-46.56676, -21.864737},
				{-46.555088, -21.859322}, {-46.552685, -21.848167}, {-46.546677, -21.836536}, {-46.51801, -21.832712},
				{-46.511143, -21.821877}, {-46.489857, -21.81805}, {-46.480587, -21.810083}, {-46.503418, -21.797491},
				{-46.510284, -21.793667}, {-46.518696, -21.794304}, {-46.52831, -21.785538}, {-46.56882, -21.767365},
				{-46.600235, -21.77119}, {-46.619118, -21.768799}, {-46.627872, -21.7739}, {-46.628044, -21.782349},
				{-46.623238, -21.785538}}}},
		},
		Address: domain.Point{
			Type:        "Point",
			Coordinates: []float64{-46.57421, -21.785742},
		},
	}

	for _, opt := range opts {
		opt(&pdv)
	}

	return &pdv
}

func NewPdvIDInput(ID string) model.PdvIDInput {
	return model.PdvIDInput{
		ID: ID,
	}
}

func NewPdvAddressInput(long, lat float64) model.PdvAddressInput {
	return model.PdvAddressInput{
		Longitude: long,
		Latitude:  lat,
	}
}

func PdvToPdvInput(pdv domain.Pdv) model.PdvInput {
	return model.PdvInput{
		TradingName:  pdv.TradingName,
		OwnerName:    pdv.OwnerName,
		Document:     pdv.Document,
		CoverageArea: pdv.CoverageArea,
		Address:      pdv.Address,
	}
}
