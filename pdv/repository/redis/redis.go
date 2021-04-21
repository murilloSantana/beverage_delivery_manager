package redis

import (
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/repository"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type redisRepository struct {
	client        *redis.Client
	jsonUnmarshal func(data []byte, v interface{}) error
	jsonMarshal   func(v interface{}) ([]byte, error)
}

func NewRedisRepository(client *redis.Client) repository.PdvCache {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	return redisRepository{
		client:        client,
		jsonUnmarshal: json.Unmarshal,
		jsonMarshal:   json.Marshal,
	}
}

func (r redisRepository) findByKey(key string) (*domain.Pdv, error) {
	resp, err := r.client.Get(context.Background(), key).Bytes()

	if err != nil {
		return nil, err
	}

	var pdv domain.Pdv
	err = r.jsonUnmarshal(resp, &pdv)

	if err != nil {
		return nil, err
	}

	return &pdv, nil
}

func (r redisRepository) FindByID(ID string) (*domain.Pdv, error) {
	return r.findByKey(ID)
}

func (r redisRepository) FindByAddress(point domain.Point) (*domain.Pdv, error) {
	key := fmt.Sprintf("%v:%v", point.Coordinates[0], point.Coordinates[1])
	return r.findByKey(key)
}

func (r redisRepository) Save(key string, pdv domain.Pdv) error {
	v, err := r.jsonMarshal(pdv)

	if err != nil {
		return err
	}

	err = r.client.Set(context.Background(), key, v, 30*time.Minute).Err()

	if err != nil {
		return err
	}

	return nil
}
