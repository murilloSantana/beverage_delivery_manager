package mongo

import (
	"beverage_delivery_manager/pdv/domain"
	"beverage_delivery_manager/pdv/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type pdvRepository struct {
	collection *mongo.Collection
}

func NewPdvRepository(collection *mongo.Collection) repository.PdvRepository {
	return pdvRepository{
		collection: collection,
	}
}

func (p pdvRepository) Save(pdv domain.Pdv) (domain.Pdv, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	pdv.ID = primitive.NewObjectID().Hex()
	resp, err := p.collection.InsertOne(ctx, pdv)
	if err != nil {
		return domain.Pdv{}, err
	}

	return p.FindByID(resp.InsertedID.(string))
}

func (p pdvRepository) HasDocument(document string) (bool, error) {
	filter := bson.M{"document": document}

	qtd, err := p.collection.CountDocuments(context.Background(), filter, withCountOptions())
	if err != nil {
		return false, err
	}

	return qtd > 0, nil
}

func withCountOptions() *options.CountOptions {
	execLimitDuration := 1 * time.Second
	limit := int64(1)
	return &options.CountOptions{Limit: &limit, MaxTime: &execLimitDuration}
}

func (p pdvRepository) FindByID(ID string) (domain.Pdv, error) {
	var pdv domain.Pdv
	filter := bson.M{"_id": ID}

	if err := p.collection.FindOne(context.Background(), filter, withFindOneOptions()).Decode(&pdv); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Pdv{}, nil
		}

		return domain.Pdv{}, err
	}

	return pdv, nil
}

func withFindOneOptions() *options.FindOneOptions {
	execLimitDuration := 1 * time.Second
	return &options.FindOneOptions{MaxTime: &execLimitDuration}
}

func (p pdvRepository) FindByAddress(point domain.Point) (domain.Pdv, error) {
	ctx := context.Background()

	cursor, err := p.collection.Aggregate(ctx, withAddressPipeline(point), withAggregateOptions())
	if err != nil {
		return domain.Pdv{}, err
	}

	defer cursor.Close(ctx)

	var pdvs []domain.Pdv
	if !cursor.TryNext(ctx) {
		return domain.Pdv{}, nil
	}

	if err := cursor.All(ctx, &pdvs); err != nil {
		return domain.Pdv{}, err
	}

	return pdvs[0], nil
}

func withAggregateOptions() *options.AggregateOptions {
	execLimitDuration := 2 * time.Second
	return &options.AggregateOptions{MaxTime: &execLimitDuration}
}

func withAddressPipeline(point domain.Point) mongo.Pipeline {
	coverageAreaFilter := bson.M{"coverageArea": bson.M{"$geoIntersects": bson.M{"$geometry": point}}}
	geoNearFilter := bson.M{"near": point, "distanceField": "calculatedDistance", "key": "address",
		"spherical": true, "query": coverageAreaFilter}

	findStage := bson.D{bson.E{Key: "$geoNear", Value: geoNearFilter}}
	limitStage := bson.D{bson.E{Key: "$limit", Value: 1}}

	return mongo.Pipeline{findStage, limitStage}
}
