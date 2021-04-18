package fakedb

import (
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/repository"
)

type pdvRepository struct{}

func NewPdvRepository() repository.PdvRepository {
	return pdvRepository{}
}

func (p pdvRepository) Save(pdv domain.Pdv) (domain.Pdv, error) {
	return domain.Pdv{}, nil
}

func (p pdvRepository) HasDocument(document string) (bool, error) {
	return false, nil
}
