package integration

import (
	mongoSettings "beverage_delivery_manager/config/mongo"
	redisSettings "beverage_delivery_manager/config/redis"
	"beverage_delivery_manager/config/settings"
	"beverage_delivery_manager/handler/graph/resolver"
	"beverage_delivery_manager/mocks/helper"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os/exec"
	"testing"
)

type pdvIntegrationTestSuite struct {
	resolver *resolver.Resolver
}

func (p *pdvIntegrationTestSuite) setupTest() {
	cmd := exec.Command("docker-compose", "up", "-d")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	err := godotenv.Load("../../.env.test.integration")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sts := settings.New()
	mongoCli, err := mongoSettings.NewClient(sts.MongoSettings)
	if err != nil {
		log.Fatal(err)
	}

	redisCli, err := redisSettings.NewClient(sts.RedisSettings)
	if err != nil {
		log.Fatal(err)
	}

	p.resolver = resolver.NewResolver(sts, mongoCli, redisCli)
}

func (p *pdvIntegrationTestSuite) tearDownTest() {
	cmd := exec.Command("docker-compose", "down")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func TestBusinessRules(t *testing.T) {
	suite := pdvIntegrationTestSuite{}

	suite.setupTest()
	defer suite.tearDownTest()

	t.Parallel()

	t.Run("Should return new pdv created", func(t *testing.T) {
		pdvInput := helper.PdvToPdvInput(*helper.NewPdv())

		expected, expectedErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, expectedErr)

		actual, actualErr := suite.resolver.Query().FindPdvByID(context.Background(), helper.NewPdvIDInput(expected.ID))

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return error when document already exists", func(t *testing.T) {
		pdvInput := helper.PdvToPdvInput(*helper.NewPdv())

		_, expectedErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, expectedErr)

		actual, actualErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.EqualError(t, actualErr, "document already exists")
		assert.Nil(t, actual)
	})

	t.Run("Should return pdv found by address", func(t *testing.T) {
		pdvInput := helper.PdvToPdvInput(*helper.NewPdv(helper.WithAddress(-46.623238, -21.785538)))

		expected, expectedErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, expectedErr)

		actual, actualErr := suite.resolver.Query().FindPdvByAddress(context.Background(),
			helper.NewPdvAddressInput(-46.623238, -21.785538))

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)

		cachedActual, actualErr := suite.resolver.Query().FindPdvByAddress(context.Background(),
			helper.NewPdvAddressInput(-46.623238, -21.785538))

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, cachedActual)
	})

	t.Run("Should return the correct pdv by id", func(t *testing.T) {
		pdvInput := helper.PdvToPdvInput(*helper.NewPdv())

		expected, expectedErr := suite.resolver.Mutation().SavePdv(context.Background(), pdvInput)

		assert.NoError(t, expectedErr)

		nonExistentID := "2345678"
		actual, actualErr := suite.resolver.Query().FindPdvByID(context.Background(), helper.NewPdvIDInput(nonExistentID))

		assert.NoError(t, actualErr)
		assert.Nil(t, actual)

		actual, actualErr = suite.resolver.Query().FindPdvByID(context.Background(), helper.NewPdvIDInput(expected.ID))

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, actual)
	})
}
