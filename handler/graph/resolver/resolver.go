package resolver

import (
	"beverage_delivery_manager/config/settings"
	"beverage_delivery_manager/handler/graph/generated"
	mongoRepo "beverage_delivery_manager/pdv/repository/mongo"
	redisRepo "beverage_delivery_manager/pdv/repository/redis"
	"beverage_delivery_manager/pdv/usecase"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type Resolver struct {
	PdvUseCase usecase.PdvUseCase
}

func NewResolver(sts settings.Settings, mongoCli *mongo.Client, redisCli *redis.Client) *Resolver {
	mongoSts := sts.MongoSettings
	database := mongoCli.Database(mongoSts.DatabaseName)
	cache := redisRepo.NewRedisRepository(redisCli)
	pdvRepository := mongoRepo.NewPdvRepository(database.Collection(mongoSts.CollectionName), cache)

	return &Resolver{
		PdvUseCase: usecase.NewPdvUseCase(pdvRepository),
	}
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
