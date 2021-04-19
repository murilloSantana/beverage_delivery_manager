package resolver

import (
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/usecase/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type DefaultPdvOption func(*domain.Pdv)

type pdvResolverTestSuite struct {
	resolver        *Resolver
	pdvUseCase      *mocks.PdvUseCase
	pdv             domain.Pdv
	pdvInput        model.PdvInput
	pdvIDInput      model.PdvIDInput
	pdvAddressInput model.PdvAddressInput
	ctx context.Context
}

func (suite *pdvResolverTestSuite) setupTest() {
	suite.pdvUseCase = new(mocks.PdvUseCase)
	suite.ctx = context.Background()
	suite.resolver = &Resolver{
		PdvUseCase: suite.pdvUseCase,
	}

	suite.pdv = newPdv()
	suite.pdvInput = pdvToPdvInput(suite.pdv)
	suite.pdvIDInput = newPdvIDInput("234343435454")
	suite.pdvAddressInput = newPdvAddressInput(-46.57421, -21.785742)
}

func withID(ID string) DefaultPdvOption {
	return func(pdv *domain.Pdv) {
		pdv.ID = ID
	}
}

func newPdv(opts ...DefaultPdvOption) domain.Pdv {
	pdv := domain.Pdv{
		TradingName: "Mercado Pinheiros",
		OwnerName:   "Luiz Santo",
		Document:    "06004905000116",
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

	return pdv
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

		suite.pdvUseCase.On("Save", mock.Anything, suite.pdv).Return(domain.Pdv{}, expectedErr)

		actual, actualErr := suite.resolver.Mutation().SavePdv(suite.ctx, suite.pdvInput)

		assert.EqualError(t, actualErr, "save error")
		assert.Empty(t, actual)
	})

	t.Run("Should return new pdv created", func(t *testing.T) {
		suite.setupTest()

		expected := newPdv(withID("234343435454"))

		suite.pdvUseCase.On("Save", mock.Anything, suite.pdv).Return(expected, nil)

		actual, actualErr := suite.resolver.Mutation().SavePdv(suite.ctx, suite.pdvInput)

		assert.NoError(t, actualErr)
		assert.Equal(t, &expected, actual)
	})
}

func TestFindByID(t *testing.T) {
	suite := pdvResolverTestSuite{}

	t.Run("Should return error when find by id fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("find by id error")

		suite.pdvUseCase.On("FindByID", suite.pdvIDInput.ID).Return(domain.Pdv{}, expectedErr)

		actual, actualErr := suite.resolver.Query().FindPdvByID(context.Background(), suite.pdvIDInput)

		assert.EqualError(t, actualErr, "find by id error")
		assert.Empty(t, actual)
	})

	t.Run("Should return empty pdv when id not found", func(t *testing.T) {
		suite.setupTest()

		suite.pdvUseCase.On("FindByID", suite.pdvIDInput.ID).Return(domain.Pdv{}, nil)

		actual, actualErr := suite.resolver.Query().FindPdvByID(context.Background(), suite.pdvIDInput)

		assert.NoError(t, actualErr)
		assert.Empty(t, actual)
	})

	t.Run("Should return a valid pdv", func(t *testing.T) {
		suite.setupTest()

		expected := newPdv(withID(suite.pdvIDInput.ID))

		suite.pdvUseCase.On("FindByID", suite.pdvIDInput.ID).Return(expected, nil)

		actual, actualErr := suite.resolver.Query().FindPdvByID(context.Background(), suite.pdvIDInput)

		assert.NoError(t, actualErr)
		assert.Equal(t, &expected, actual)
	})
}

func TestFindByAddress(t *testing.T) {
	suite := pdvResolverTestSuite{}

	t.Run("Should return error when find by address fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("find by address error")

		point := newPoint(suite.pdvAddressInput.Longitude, suite.pdvAddressInput.Latitude)
		suite.pdvUseCase.On("FindByAddress", point).Return(domain.Pdv{}, expectedErr)

		actual, actualErr := suite.resolver.Query().FindPdvByAddress(context.Background(), suite.pdvAddressInput)

		assert.EqualError(t, actualErr, "find by address error")
		assert.Empty(t, actual)
	})

	t.Run("Should return empty PDV when the entered address is not in a coverage area", func(t *testing.T) {
		suite.setupTest()

		point := newPoint(suite.pdvAddressInput.Longitude, suite.pdvAddressInput.Latitude)
		suite.pdvUseCase.On("FindByAddress", point).Return(domain.Pdv{}, nil)

		actual, actualErr := suite.resolver.Query().FindPdvByAddress(context.Background(), suite.pdvAddressInput)

		assert.NoError(t, actualErr)
		assert.Empty(t, actual)
	})

	t.Run("Should return a valid pdv", func(t *testing.T) {
		suite.setupTest()

		expected := newPdv()

		point := newPoint(suite.pdvAddressInput.Longitude, suite.pdvAddressInput.Latitude)
		suite.pdvUseCase.On("FindByAddress", point).Return(suite.pdv, nil)

		actual, actualErr := suite.resolver.Query().FindPdvByAddress(context.Background(), suite.pdvAddressInput)

		assert.NoError(t, actualErr)
		assert.Equal(t, &expected, actual)
	})
}
