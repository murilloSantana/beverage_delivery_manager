package usecase

import (
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/repository/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DefaultPdvOption func(*domain.Pdv)

type pdvUseCaseTestSuite struct {
	pdvUseCase    PdvUseCase
	pdvRepository *mocks.PdvRepository
	pdv           domain.Pdv
}

func (suite *pdvUseCaseTestSuite) setupTest() {
	suite.pdvRepository = new(mocks.PdvRepository)
	suite.pdvUseCase = NewPdvUseCase(suite.pdvRepository)
	suite.pdv = newPdv()
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

func TestSave(t *testing.T) {
	suite := pdvUseCaseTestSuite{}

	t.Run("Should return error when has document func fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("has document error")

		suite.pdvRepository.On("HasDocument", suite.pdv.Document).Return(false, expectedErr)
		actual, actualErr := suite.pdvUseCase.Save(suite.pdv)

		assert.EqualError(t, actualErr, "has document error")
		assert.Empty(t, actual)
	})

	t.Run("Should return error when document already exists", func(t *testing.T) {
		suite.setupTest()

		suite.pdvRepository.On("HasDocument", suite.pdv.Document).Return(true, nil)
		actual, actualErr := suite.pdvUseCase.Save(suite.pdv)

		assert.EqualError(t, actualErr, "document already exists")
		assert.Empty(t, actual)
	})

	t.Run("Should return error when save fail", func(t *testing.T) {
		suite.setupTest()

		expectedErr := errors.New("save error")

		suite.pdvRepository.On("HasDocument", suite.pdv.Document).Return(false, nil)
		suite.pdvRepository.On("Save", suite.pdv).Return(domain.Pdv{}, expectedErr)

		actual, actualErr := suite.pdvUseCase.Save(suite.pdv)

		assert.EqualError(t, actualErr, "save error")
		assert.Empty(t, actual)
	})

	t.Run("Should return new pdv created", func(t *testing.T) {
		suite.setupTest()

		expected := newPdv(withID("234343435454"))

		suite.pdvRepository.On("HasDocument", suite.pdv.Document).Return(false, nil)
		suite.pdvRepository.On("Save", suite.pdv).Return(expected, nil)

		actual, actualErr := suite.pdvUseCase.Save(suite.pdv)

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})
}
