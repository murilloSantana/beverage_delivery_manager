package server

import (
	"beverage_delivery_manager/config/settings"
	"beverage_delivery_manager/handler/graph/config"
	"beverage_delivery_manager/handler/graph/generated"
	"beverage_delivery_manager/handler/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
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
	apq, _ := config.NewAPQ(sts.RedisSettings)

	resolver := resolver.NewResolver(sts, mongoCli, redisCli)
	srv := handler.
		New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	srv.Use(extension.AutomaticPersistedQuery{Cache: apq})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	if !sts.ApplicationSettings.IsProduction() {
		srv.Use(extension.Introspection{})
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}

	http.Handle("/query", srv)

	return graphServer{
		instance: srv,
		sts:      sts,
	}
}

func (g graphServer) Run() error {
	return http.ListenAndServe(g.sts.ApplicationSettings.Port, nil)
}
