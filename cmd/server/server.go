package server

import (
	logger "beverage_delivery_manager/config/log"
	"beverage_delivery_manager/config/settings"
	"beverage_delivery_manager/pdv/repository/fakedb"
	"beverage_delivery_manager/pdv/usecase"

	"beverage_delivery_manager/handler/graph/generated"
	"beverage_delivery_manager/handler/graph/resolver"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: newResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Info(nil, fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", sts.Port))

	return http.ListenAndServe(fmt.Sprintf(":%s", sts.Port), nil)
}

func newResolver() *resolver.Resolver {
	return &resolver.Resolver{
		PdvUseCase: usecase.NewPdvUseCase(fakedb.NewPdvRepository()),
	}
}
