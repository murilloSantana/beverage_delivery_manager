package usecase

import "beverage_delivery_manager/pdv/domain"

type PdvUseCase interface {
	Save(pdv domain.Pdv) (domain.Pdv, error)
}

type pdvUseCase struct {
}

func NewPdvUseCase() PdvUseCase {
	return pdvUseCase{}
}

func (p pdvUseCase) Save(pdv domain.Pdv) (domain.Pdv, error) {
	return domain.Pdv{}, nil
}
