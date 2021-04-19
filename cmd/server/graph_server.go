package server

import (
	"beverage_delivery_manager/config/settings"
	"beverage_delivery_manager/handler/graph/generated"
	"beverage_delivery_manager/handler/graph/resolver"
	mongoRepository "beverage_delivery_manager/pdv/repository/mongo"
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

func New(sts settings.Settings, mongoCli *mongo.Client, _ *redis.Client) Runner {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: newResolver(sts, mongoCli)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	return graphServer{
		instance: srv,
		sts:      sts,
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

func (g graphServer) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", g.sts.Port), nil)
}
