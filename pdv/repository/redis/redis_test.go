package redis

import (
	"beverage_delivery_manager/mocks/helper"
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/repository"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type pdvRedisTestOption func(*pdvRedisTestSuite)

type pdvRedisTestSuite struct {
	cache         repository.PdvCache
	miniredis     *miniredis.Miniredis
	jsonUnmarshal func(data []byte, v interface{}) error
	jsonMarshal   func(v interface{}) ([]byte, error)
	point         domain.Point
}

func WithJsonUnmarshal(jsonUnmarshal func(data []byte, v interface{}) error) pdvRedisTestOption {
	return func(suite *pdvRedisTestSuite) {
		suite.jsonUnmarshal = jsonUnmarshal
	}
}

func WithJsonMarshal(jsonMarshal func(v interface{}) ([]byte, error)) pdvRedisTestOption {
	return func(suite *pdvRedisTestSuite) {
		suite.jsonMarshal = jsonMarshal
	}
}

func (suite *pdvRedisTestSuite) setupTest(opts ...pdvRedisTestOption) {
	for _, opt := range opts {
		opt(suite)
	}

	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}

	suite.point = domain.Point{Type: "Point", Coordinates: []float64{-46.57421, -21.785742}}
	suite.miniredis = mr
	suite.cache = redisRepository{
		client: redis.NewClient(&redis.Options{
			Addr: suite.miniredis.Addr(),
		}),
		jsonUnmarshal: suite.jsonUnmarshal,
		jsonMarshal:   suite.jsonMarshal,
	}
}

func TestFindByID(t *testing.T) {
	suite := pdvRedisTestSuite{}

	t.Run("Should throw error when pdv is not found", func(t *testing.T) {
		suite.setupTest()
		defer suite.miniredis.Close()

		actual, actualErr := suite.cache.FindByID("2000")

		assert.Error(t, actualErr)
		assert.Empty(t, actual)
	})

	t.Run("Should throw error, because pdv found is invalid", func(t *testing.T) {
		errJsonUnmarshalMock := func(data []byte, v interface{}) error {
			return errors.New("generic error")
		}

		suite.setupTest(WithJsonUnmarshal(errJsonUnmarshalMock))
		defer suite.miniredis.Close()

		ID := "2345678"
		p, _ := json.Marshal(helper.NewPdv())
		_ = suite.miniredis.Set(ID, string(p))

		actual, actualErr := suite.cache.FindByID(ID)

		assert.EqualError(t, actualErr, "generic error")
		assert.Empty(t, actual)
	})

	t.Run("Should return a valid pdv", func(t *testing.T) {
		suite.setupTest(WithJsonUnmarshal(json.Unmarshal))
		defer suite.miniredis.Close()

		ID := "2345678"
		expected, _ := json.Marshal(helper.NewPdv())
		_ = suite.miniredis.Set(ID, string(expected))

		actual, actualErr := suite.cache.FindByID(ID)

		s, _ := json.Marshal(actual)
		suite.miniredis.CheckGet(t, ID, string(s))

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, s)
	})
}

func TestFindAddress(t *testing.T) {
	suite := pdvRedisTestSuite{}

	t.Run("Should throw error when pdv is not found", func(t *testing.T) {
		suite.setupTest()
		defer suite.miniredis.Close()

		actual, actualErr := suite.cache.FindByAddress(suite.point)

		assert.Error(t, actualErr)
		assert.Empty(t, actual)
	})

	t.Run("Should throw error, because pdv found is invalid", func(t *testing.T) {
		errJsonUnmarshalMock := func(data []byte, v interface{}) error {
			return errors.New("generic error")
		}

		suite.setupTest(WithJsonUnmarshal(errJsonUnmarshalMock))
		defer suite.miniredis.Close()

		key := fmt.Sprintf("%v:%v", suite.point.Coordinates[0], suite.point.Coordinates[1])
		p, _ := json.Marshal(helper.NewPdv())
		_ = suite.miniredis.Set(key, string(p))

		actual, actualErr := suite.cache.FindByAddress(suite.point)

		assert.EqualError(t, actualErr, "generic error")
		assert.Empty(t, actual)
	})

	t.Run("Should return a valid pdv", func(t *testing.T) {
		suite.setupTest(WithJsonUnmarshal(json.Unmarshal))
		defer suite.miniredis.Close()

		key := fmt.Sprintf("%v:%v", suite.point.Coordinates[0], suite.point.Coordinates[1])
		expected, _ := json.Marshal(helper.NewPdv())
		_ = suite.miniredis.Set(key, string(expected))

		actual, actualErr := suite.cache.FindByAddress(suite.point)

		s, _ := json.Marshal(actual)
		suite.miniredis.CheckGet(t, key, string(s))

		assert.NoError(t, actualErr)
		assert.Equal(t, expected, s)
	})
}

func TestSavePdv(t *testing.T) {
	suite := pdvRedisTestSuite{}

	t.Run("Should throw error when value is an invalid json", func(t *testing.T) {
		expected := false
		errJsonMarshalMock := func(v interface{}) ([]byte, error) {
			return nil, errors.New("generic error")
		}

		suite.setupTest(WithJsonMarshal(errJsonMarshalMock))
		defer suite.miniredis.Close()

		ID := "2345678"
		actualErr := suite.cache.Save(ID, helper.NewPdv())

		actual := suite.miniredis.Exists(ID)

		assert.EqualError(t, actualErr, "generic error")
		assert.Equal(t, expected, actual)
	})

	t.Run("Should throw error when set function fail", func(t *testing.T) {
		suite.setupTest(WithJsonMarshal(json.Marshal))
		suite.miniredis.Close()

		actualErr := suite.cache.Save("2345", helper.NewPdv())

		assert.Error(t, actualErr)
	})

	t.Run("Should save new pdv in redis", func(t *testing.T) {
		suite.setupTest(WithJsonMarshal(json.Marshal))
		defer suite.miniredis.Close()

		expected := helper.NewPdv()

		ID := "34434343"
		actualErr := suite.cache.Save(ID, expected)

		actual, _ := json.Marshal(expected)

		assert.NoError(t, actualErr)
		suite.miniredis.CheckGet(t, ID, string(actual))
	})
}
