package server

import (
	logger "beverage_delivery_manager/cmd/log"
	"beverage_delivery_manager/cmd/settings"
	"beverage_delivery_manager/graph/generated"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
)

type Server interface {
	Run(sts settings.Settings)
}

func New(sts settings.Settings) error {
	log := logger.NewLogger()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &generated.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Info(nil, fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", sts.Port))

	return http.ListenAndServe(fmt.Sprintf(":%s", sts.Port), nil)
}
