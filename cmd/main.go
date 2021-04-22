package main

import (
	"beverage_delivery_manager/cmd/server"
	"beverage_delivery_manager/config/mongo"
	"beverage_delivery_manager/config/redis"
	"beverage_delivery_manager/config/settings"
	"log"
)

func main() {
	sts := settings.New()

	mongoCli, err := mongo.NewClient(sts.MongoSettings)
	if err != nil {
		log.Fatal(err)
	}

	redisCli, err := redis.NewClient(sts.RedisSettings)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.New(sts, mongoCli, redisCli).Run())
}
