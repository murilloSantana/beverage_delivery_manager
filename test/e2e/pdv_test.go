package e2e

import (
	mongoSettings "beverage_delivery_manager/config/mongo"
	"beverage_delivery_manager/config/settings"
	"beverage_delivery_manager/handler/graph/resolver"
	"beverage_delivery_manager/mocks/helper"
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
)

type pdvE2ETestSuite struct {
	resolver *resolver.Resolver
}

func (p *pdvE2ETestSuite) setupTest() {
	cmd := exec.Command("docker-compose", "up", "-d")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	err := godotenv.Load("../../.env.test.e2e")
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

func TestSave(t *testing.T) {
	suite := pdvE2ETestSuite{}

	suite.setupTest()
	defer suite.tearDownTest()

	t.Parallel()

	t.Run("Should return new pdv created", func(t *testing.T) {
		pdvInput := helper.PdvToPdvInput(helper.NewPdv())

		actual, mutationErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, mutationErr)

		expected, queryErr := suite.resolver.Query().FindPdvByID(context.Background(), helper.NewPdvIDInput(actual.ID))

		assert.NoError(t, queryErr)
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return error when document already exists", func(t *testing.T) {
		pdvInput := helper.PdvToPdvInput(helper.NewPdv())

		_, mutationErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, mutationErr)

		expected, actualMutationErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, mutationErr)

		assert.EqualError(t, actualMutationErr, "document already exists")
		assert.Nil(t, expected)
	})
}
