package repository

import "beverage_delivery_manager/pdv/domain"

//go:generate mockery --name PdvRepository --case=underscore

type PdvRepository interface {
	HasDocument(document string) (bool, error)
	Save(pdv domain.Pdv) (domain.Pdv, error)
}
