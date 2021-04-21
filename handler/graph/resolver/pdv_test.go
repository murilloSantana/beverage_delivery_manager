package resolver

import (
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/mocks"
	"beverage_delivery_manager/mocks/helper"
	"beverage_delivery_manager/pdv/domain"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type pdvResolverTestSuite struct {
	resolver        *Resolver
	pdvUseCase      *mocks.PdvUseCase
	pdv             *domain.Pdv
	pdvInput        model.PdvInput
	pdvIDInput      model.PdvIDInput
	pdvAddressInput model.PdvAddressInput
	ctx             context.Context
}

func (suite *pdvResolverTestSuite) setupTest() {
	suite.pdvUseCase = new(mocks.PdvUseCase)
	suite.ctx = context.Background()
	suite.resolver = &Resolver{
		PdvUseCase: suite.pdvUseCase,
	}

	suite.pdv = helper.NewPdv()
	suite.pdvInput = pdvToPdvInput(*suite.pdv)
	suite.pdvIDInput = newPdvIDInput("234343435454")
	suite.pdvAddressInput = newPdvAddressInput(-46.57421, -21.785742)
}

func newPdvAddressInput(longitude, latitude float64) model.PdvAddressInput {
	return model.PdvAddressInput{
		Longitude: longitude,
		Latitude:  latitude,
	}
}

func newPdvIDInput(ID string) model.PdvIDInput {
	return model.PdvIDInput{
		ID: ID,
	}
}

func newPoint(coordinates ...float64) domain.Point {
	return domain.Point{
		Type:        "Point",
		Coordinates: coordinates,
	}
}

func pdvToPdvInput(pdv domain.Pdv) model.PdvInput {
	return model.PdvInput{
		TradingName:  pdv.TradingName,
		OwnerName:    pdv.OwnerName,
		Document:     pdv.Document,
		CoverageArea: pdv.CoverageArea,
		Address:      pdv.Address,
	}
}

func TestSavePdv(t *testing.T) {
	suite := pdvResolverTestSuite{}

	t.Run("Should return error when save fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("save error")

		suite.pdvUseCase.On("Save", mock.Anything, *suite.pdv).Return(nil, expectedErr)

		actual, actualErr := suite.resolver.Mutation().SavePdv(suite.ctx, suite.pdvInput)

		assert.EqualError(t, actualErr, "save error")
		assert.Nil(t, actual)
	})

	t.Run("Should return new pdv created", func(t *testing.T) {
		suite.setupTest()

		expected := helper.NewPdv(helper.WithID("234343435454"))

		suite.pdvUseCase.On("Save", mock.Anything, *suite.pdv).Return(expected, nil)

		actual, actualErr := suite.resolver.Mutation().SavePdv(suite.ctx, suite.pdvInput)

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})
}

func TestFindByID(t *testing.T) {
	suite := pdvResolverTestSuite{}

	t.Run("Should return error when find by id fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("find by id error")

		suite.pdvUseCase.On("FindByID", suite.pdvIDInput.ID).Return(nil, expectedErr)

		actual, actualErr := suite.resolver.Query().FindPdvByID(context.Background(), suite.pdvIDInput)

		assert.EqualError(t, actualErr, "find by id error")
		assert.Nil(t, actual)
	})

	t.Run("Should return empty pdv when id not found", func(t *testing.T) {
		suite.setupTest()

		suite.pdvUseCase.On("FindByID", suite.pdvIDInput.ID).Return(nil, nil)

		actual, actualErr := suite.resolver.Query().FindPdvByID(context.Background(), suite.pdvIDInput)

		assert.NoError(t, actualErr)
		assert.Nil(t, actual)
	})

	t.Run("Should return a valid pdv", func(t *testing.T) {
		suite.setupTest()

		expected := helper.NewPdv(helper.WithID(suite.pdvIDInput.ID))

		suite.pdvUseCase.On("FindByID", suite.pdvIDInput.ID).Return(expected, nil)

		actual, actualErr := suite.resolver.Query().FindPdvByID(context.Background(), suite.pdvIDInput)

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})
}

func TestFindByAddress(t *testing.T) {
	suite := pdvResolverTestSuite{}

	t.Run("Should return error when find by address fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("find by address error")

		point := newPoint(suite.pdvAddressInput.Longitude, suite.pdvAddressInput.Latitude)
		suite.pdvUseCase.On("FindByAddress", point).Return(nil, expectedErr)

		actual, actualErr := suite.resolver.Query().FindPdvByAddress(context.Background(), suite.pdvAddressInput)

		assert.EqualError(t, actualErr, "find by address error")
		assert.Nil(t, actual)
	})

	t.Run("Should return empty PDV when the entered address is not in a coverage area", func(t *testing.T) {
		suite.setupTest()

		point := newPoint(suite.pdvAddressInput.Longitude, suite.pdvAddressInput.Latitude)
		suite.pdvUseCase.On("FindByAddress", point).Return(nil, nil)

		actual, actualErr := suite.resolver.Query().FindPdvByAddress(context.Background(), suite.pdvAddressInput)

		assert.NoError(t, actualErr)
		assert.Nil(t, actual)
	})

	t.Run("Should return a valid pdv", func(t *testing.T) {
		suite.setupTest()

		expected := helper.NewPdv(helper.WithDocument(suite.pdv.Document))

		point := newPoint(suite.pdvAddressInput.Longitude, suite.pdvAddressInput.Latitude)
		suite.pdvUseCase.On("FindByAddress", point).Return(suite.pdv, nil)

		actual, actualErr := suite.resolver.Query().FindPdvByAddress(context.Background(), suite.pdvAddressInput)

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})
}
