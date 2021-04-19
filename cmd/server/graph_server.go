package server

import (
	"beverage_delivery_manager/config/settings"
	"beverage_delivery_manager/handler/graph/generated"
	"beverage_delivery_manager/handler/graph/resolver"
	mongoRepo "beverage_delivery_manager/pdv/repository/mongo"
	redisRepo "beverage_delivery_manager/pdv/repository/redis"
	"beverage_delivery_manager/pdv/usecase"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type graphServer struct {
	instance *handler.Server
	sts      settings.Settings
}

func New(sts settings.Settings, mongoCli *mongo.Client, redisCli *redis.Client) Runner {
	srv := handler.
		NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: newResolver(sts, mongoCli, redisCli)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	return graphServer{
		instance: srv,
		sts:      sts,
	}
}

func newResolver(sts settings.Settings, mongoCli *mongo.Client, redisCli *redis.Client) *resolver.Resolver {
	mongoSts := sts.MongoSettings
	database := mongoCli.Database(mongoSts.DatabaseName)
	cache := redisRepo.NewRedisRepository(redisCli)
	pdvRepository := mongoRepo.NewPdvRepository(database.Collection(mongoSts.CollectionName), cache)

	return &resolver.Resolver{
		PdvUseCase: usecase.NewPdvUseCase(pdvRepository),
	}
}

func (g graphServer) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", g.sts.Port), nil)
}
