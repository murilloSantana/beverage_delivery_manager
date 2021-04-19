package usecase

import (
	"beverage_delivery_manager/mocks"
	"beverage_delivery_manager/mocks/helper"
	"beverage_delivery_manager/pdv/domain"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type pdvUseCaseTestSuite struct {
	pdvUseCase    PdvUseCase
	pdvRepository *mocks.PdvRepository
	pdv           domain.Pdv
	ctx           context.Context
}

func (suite *pdvUseCaseTestSuite) setupTest() {
	suite.ctx = context.Background()
	suite.pdvRepository = new(mocks.PdvRepository)
	suite.pdvUseCase = NewPdvUseCase(suite.pdvRepository)
	suite.pdv = helper.NewPdv()
}

func newPoint(coordinates ...float64) domain.Point {
	return domain.Point{
		Type:        "Point",
		Coordinates: coordinates,
	}
}

func TestSave(t *testing.T) {
	suite := pdvUseCaseTestSuite{}

	generateNewIDMock := func() string {
		return ""
	}

	t.Run("Should return error when has document func fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("has document error")

		suite.pdvRepository.On("HasDocument", suite.pdv.Document).Return(false, expectedErr)
		actual, actualErr := suite.pdvUseCase.Save(suite.ctx, suite.pdv)

		assert.EqualError(t, actualErr, "has document error")
		assert.Empty(t, actual)
	})

	t.Run("Should return error when document already exists", func(t *testing.T) {
		suite.setupTest()

		suite.pdvRepository.On("HasDocument", suite.pdv.Document).Return(true, nil)
		actual, actualErr := suite.pdvUseCase.Save(suite.ctx, suite.pdv)

		assert.EqualError(t, actualErr, "document already exists")
		assert.Empty(t, actual)
	})

	t.Run("Should return error when save fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("save error")

		suite.pdvRepository.On("HasDocument", suite.pdv.Document).Return(false, nil)
		suite.pdvRepository.On("GenerateNewID").Return(generateNewIDMock)
		suite.pdvRepository.On("Save", suite.ctx, suite.pdv, mock.AnythingOfType("func() string")).Return(domain.Pdv{}, expectedErr)

		actual, actualErr := suite.pdvUseCase.Save(suite.ctx, suite.pdv)

		assert.EqualError(t, actualErr, "save error")
		assert.Empty(t, actual)
	})

	t.Run("Should return new pdv created", func(t *testing.T) {
		suite.setupTest()

		expected := helper.NewPdv(helper.WithID("234343435454"))

		suite.pdvRepository.On("HasDocument", suite.pdv.Document).Return(false, nil)
		suite.pdvRepository.On("GenerateNewID").Return(generateNewIDMock)
		suite.pdvRepository.On("Save", suite.ctx, suite.pdv, mock.AnythingOfType("func() string")).Return(expected, nil)

		actual, actualErr := suite.pdvUseCase.Save(suite.ctx, suite.pdv)

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})
}

func TestFindByID(t *testing.T) {
	suite := pdvUseCaseTestSuite{}

	t.Run("Should return error when find by id fail", func(t *testing.T) {
		suite.setupTest()

		ID := "2345678"
		expectedErr := errors.New("find by id error")

		suite.pdvRepository.On("FindByID", ID).Return(domain.Pdv{}, expectedErr)
		actual, actualErr := suite.pdvUseCase.FindByID(ID)

		assert.EqualError(t, actualErr, "find by id error")
		assert.Empty(t, actual)
	})

	t.Run("Should return empty pdv when id not found", func(t *testing.T) {
		suite.setupTest()

		ID := "2345678"

		suite.pdvRepository.On("FindByID", ID).Return(domain.Pdv{}, nil)
		actual, actualErr := suite.pdvUseCase.FindByID(ID)

		assert.NoError(t, actualErr)
		assert.Empty(t, actual)
	})

	t.Run("Should return a valid pdv", func(t *testing.T) {
		suite.setupTest()

		ID := "2345678"
		expected := helper.NewPdv(helper.WithID(ID))

		suite.pdvRepository.On("FindByID", ID).Return(expected, nil)
		actual, actualErr := suite.pdvUseCase.FindByID(ID)

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})
}

func TestFindByAddress(t *testing.T) {
	suite := pdvUseCaseTestSuite{}

	t.Run("Should return error when find by address fail", func(t *testing.T) {
		suite.setupTest()

		point := newPoint(-46.57421, -21.785742)
		expectedErr := errors.New("find by address error")

		suite.pdvRepository.On("FindByAddress", point).Return(domain.Pdv{}, expectedErr)
		actual, actualErr := suite.pdvUseCase.FindByAddress(point)

		assert.EqualError(t, actualErr, "find by address error")
		assert.Empty(t, actual)
	})

	t.Run("Should return empty PDV when the entered address is not in a coverage area", func(t *testing.T) {
		suite.setupTest()

		point := newPoint(-46.57421, -21.785742)

		suite.pdvRepository.On("FindByAddress", point).Return(domain.Pdv{}, nil)
		actual, actualErr := suite.pdvUseCase.FindByAddress(point)

		assert.NoError(t, actualErr)
		assert.Empty(t, actual)
	})

	t.Run("Should return a valid pdv", func(t *testing.T) {
		suite.setupTest()

		ID := "2345678"
		expected := helper.NewPdv(helper.WithID(ID))

		point := newPoint(-46.57421, -21.785742)

		suite.pdvRepository.On("FindByAddress", point).Return(expected, nil)
		actual, actualErr := suite.pdvUseCase.FindByAddress(point)

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})
}
