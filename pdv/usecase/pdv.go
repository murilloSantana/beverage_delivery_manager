package usecase

import (
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/repository"
	"context"
	"fmt"
)

//go:generate mockery --name PdvUseCase --case=underscore --output ../../mocks

type PdvUseCase interface {
	Save(ctx context.Context, pdv domain.Pdv) (*domain.Pdv, error)
	FindByID(ID string) (*domain.Pdv, error)
	FindByAddress(point domain.Point) (*domain.Pdv, error)
}

type pdvUseCase struct {
	repository repository.PdvRepository
}

func NewPdvUseCase(repository repository.PdvRepository) PdvUseCase {
	return pdvUseCase{
		repository: repository,
	}
}

func (p pdvUseCase) Save(ctx context.Context, pdv domain.Pdv) (*domain.Pdv, error) {
	err := p.hasDocument(pdv.Document)
	if err != nil {
		return nil, err
	}

	newPdv, err := p.repository.Save(ctx, pdv, p.repository.GenerateNewID())
	if err != nil {
		return nil, err
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

func (p pdvUseCase) FindByID(ID string) (*domain.Pdv, error) {
	pdv, err := p.repository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return pdv, nil
}

func (p pdvUseCase) FindByAddress(point domain.Point) (*domain.Pdv, error) {
	pdv, err := p.repository.FindByAddress(point)
	if err != nil {
		return nil, err
	}

	return pdv, nil
}
