package usecase

import (
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/repository"
	"fmt"
)

//go:generate mockery --name PdvUseCase --case=underscore

type PdvUseCase interface {
	Save(pdv domain.Pdv) (domain.Pdv, error)
	FindByID(ID string) (domain.Pdv, error)
	FindByAddress(coordinates domain.PointCoordinates) (domain.Pdv, error)
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
	err := p.hasDocument(pdv.Document)
	if err != nil {
		return domain.Pdv{}, err
	}

	newPdv, err := p.repository.Save(pdv)
	if err != nil {
		return domain.Pdv{}, err
	}

	return newPdv, nil
}

func (p pdvUseCase) hasDocument(document string) error {
	hasDoc, err := p.repository.HasDocument(document)
	if err != nil {
		return err
	}

	if hasDoc {
		return fmt.Errorf("document already exists")
	}

	return nil
}

func (p pdvUseCase) FindByID(ID string) (domain.Pdv, error) {
	pdv, err := p.repository.FindByID(ID)
	if err != nil {
		return domain.Pdv{}, err
	}

	return pdv, nil
}

func (p pdvUseCase) FindByAddress(coordinates domain.PointCoordinates) (domain.Pdv, error) {
	pdv, err := p.repository.FindByAddress(coordinates)
	if err != nil {
		return domain.Pdv{}, err
	}

	return pdv, nil
}