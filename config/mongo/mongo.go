package mongo

import (
	"beverage_delivery_manager/config/settings"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewClient(sts settings.Settings) (*mongo.Client, error) {
	mongoSts := sts.MongoSettings

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	opts := options.Client().
		ApplyURI(mongoSts.URL).
		SetMinPoolSize(mongoSts.MinPoolSize).
		SetMaxPoolSize(mongoSts.MaxPoolSize).
		SetMaxConnIdleTime(mongoSts.MaxConnIdleTime)

	mongoCli, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return mongoCli, nil
}
