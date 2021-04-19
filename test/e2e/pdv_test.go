package e2e

import (
	mongoSettings "beverage_delivery_manager/config/mongo"
	"beverage_delivery_manager/config/settings"
	"beverage_delivery_manager/handler/graph/model"
	"beverage_delivery_manager/handler/graph/resolver"
	"beverage_delivery_manager/pdv/domain"
	mongoRepository "beverage_delivery_manager/pdv/repository/mongo"
	"beverage_delivery_manager/pdv/usecase"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os/exec"
	"testing"
	"time"
)

type DefaultPdvOption func(*domain.Pdv)

type pdvE2ETestSuite struct {
	resolver *resolver.Resolver
}

func (p *pdvE2ETestSuite) setupTest() {
	cmd := exec.Command("docker-compose", "up", "-d")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	err := godotenv.Load("../../.env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sts := settings.New()
	mongoCli, err := mongoSettings.NewClient(sts)
	if err != nil {
		log.Fatal(err)
	}

	p.resolver = newResolver(sts, mongoCli)
}

func (p *pdvE2ETestSuite) tearDownTest() {
	cmd := exec.Command("docker-compose", "down")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func newResolver(sts settings.Settings, mongoCli *mongo.Client) *resolver.Resolver {
	mongoSts := sts.MongoSettings
	database := mongoCli.Database(mongoSts.DatabaseName)
	pdvRepository := mongoRepository.NewPdvRepository(database.Collection(mongoSts.CollectionName))

	return &resolver.Resolver{
		PdvUseCase: usecase.NewPdvUseCase(pdvRepository),
	}
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
		Document:    time.Now().String(),
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

func newPdvIDInput(ID string) model.PdvIDInput {
	return model.PdvIDInput{
		ID: ID,
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

func TestSave(t *testing.T) {
	suite := pdvE2ETestSuite{}

	suite.setupTest()
	defer suite.tearDownTest()

	t.Parallel()

	t.Run("Should return new pdv created", func(t *testing.T) {
		pdvInput := pdvToPdvInput(newPdv())

		actual, mutationErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, mutationErr)

		expected, queryErr := suite.resolver.Query().FindPdvByID(context.Background(), newPdvIDInput(actual.ID))

		assert.NoError(t, queryErr)
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return error when document already exists", func(t *testing.T) {
		pdvInput := pdvToPdvInput(newPdv())

		_, mutationErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, mutationErr)

		expected, actualMutationErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, mutationErr)

		assert.EqualError(t, actualMutationErr, "document already exists")
		assert.Nil(t, expected)
	})
}