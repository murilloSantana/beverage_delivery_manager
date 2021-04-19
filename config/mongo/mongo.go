package mongo

import (
	"beverage_delivery_manager/config/settings"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewClient(sts settings.MongoSettings) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	opts := options.Client().
		ApplyURI(sts.URL).
		SetMinPoolSize(sts.MinPoolSize).
		SetMaxPoolSize(sts.MaxPoolSize).
		SetMaxConnIdleTime(sts.MaxConnIdleTime)

	mongoCli, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return mongoCli, nil
}
