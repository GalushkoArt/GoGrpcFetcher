package repository

import (
	"context"
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db: db,
	}
}

func (r *MongoRepository) UpsertItem(ctx context.Context, items []model.Item) error {
	models := make([]mongo.WriteModel, 0, len(items))
	for _, item := range items {
		models = append(models, mongo.NewReplaceOneModel().SetUpsert(true).SetFilter(bson.D{{"name", item.Name}}).SetReplacement(item))
	}
	opts := options.BulkWrite().SetOrdered(false)
	results, err := r.db.Collection("items").BulkWrite(ctx, models, opts)
	if err != nil {
		log.Error().Err(err).Interface("result", results).Msg("Failed to upsert items!")
	} else {
		log.Info().Interface("upserted", results.UpsertedCount).Msg("Upserted items!")
	}
	return err
}

func (r *MongoRepository) GetItems(ctx context.Context, getOptions model.GetOptions) ([]model.Item, error) {
	var result []model.Item
	findOptions := options.Find()
	sorts := make(bson.D, 0, 2)
	if getOptions.NameSort != nil {
		sorts = append(sorts, bson.E{Key: "name", Value: *getOptions.NameSort})
	}
	if getOptions.PriceSort != nil {
		sorts = append(sorts, bson.E{Key: "price", Value: *getOptions.PriceSort})
	}
	if len(sorts) > 0 {
		findOptions.SetSort(sorts)
	}
	if getOptions.Paging != nil {
		findOptions.SetSkip(getOptions.Paging.Page * getOptions.Paging.PageSize)
		findOptions.SetLimit(getOptions.Paging.PageSize)
	}
	cursor, err := r.db.Collection("items").Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *MongoRepository) StartStatus(ctx context.Context, url string) (*primitive.ObjectID, error) {
	fetch := model.FetchStatus{
		URL:         url,
		FetchStatus: model.IN_PROGRESS,
		Started:     time.Now(),
	}
	setUpsert := options.Replace().SetUpsert(true)
	result, err := r.db.Collection("fetches").ReplaceOne(ctx, bson.D{{"url", url}}, fetch, setUpsert)
	if id, ok := result.UpsertedID.(primitive.ObjectID); ok {
		return &id, nil
	}
	return nil, err
}

func (r *MongoRepository) GetStatus(ctx context.Context, url string) (*model.FetchStatus, error) {
	var result *model.FetchStatus
	err := r.db.Collection("fetches").FindOne(ctx, bson.D{{"url", url}}).Decode(result)
	return result, err
}

func (r *MongoRepository) FinishStatus(ctx context.Context, id *primitive.ObjectID) error {
	_, err := r.db.Collection("fetches").UpdateByID(ctx, id, bson.D{{"$set", bson.D{{"fetch_status", model.DONE}, {"finished", time.Now()}}}})
	return err
}
