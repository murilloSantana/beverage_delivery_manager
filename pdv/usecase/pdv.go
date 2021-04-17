package usecase

import (
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/repository"
)

type PdvUseCase interface {
	Save(pdv domain.Pdv) (domain.Pdv, error)
}

type pdvUseCase struct {
	repository repository.PdvRepository
}

func NewPdvUseCase(repository repository.PdvRepository) PdvUseCase {
	return pdvUseCase{
		repository: repository,
	}
}

func (p pdvUseCase) Save(pdv domain.Pdv) (domain.Pdv, error) {
	return domain.Pdv{}, nil
}
