package mongostorage

import (
	storage2 "backend/storage"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

const dbName = "product_db"
const collectionName = "products"

type storage struct {
	products *mongo.Collection
}

func DatabaseStorage(mongoUrl string) *storage {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		panic(err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	ensureIndexes(ctx, collection)

	return &storage{
		products: collection,
	}
}

func ensureIndexes(ctx context.Context, collection *mongo.Collection) {
	indexModels := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "name", Value: bsonx.Int32(1)},
				{Key: "_id", Value: bsonx.Int32(-1)},
			},
		},
	}
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := collection.Indexes().CreateMany(ctx, indexModels, opts)
	if err != nil {
		panic(fmt.Errorf("failed to ensure indexes - %w", err))
	}
}

func (s *storage) Create(ctx context.Context, data storage2.Product) error {
	_, err := s.products.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("something went wrong - %w", storage2.CommonStorageError)
	}
	return nil
}

func (s *storage) GetProduct(ctx context.Context, id string) (storage2.Product, error) {
	var result storage2.Product
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return storage2.Product{}, fmt.Errorf("invalid id - %w", storage2.CommonStorageError)
	}
	err = s.products.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return storage2.Product{}, fmt.Errorf("no products with id %v - %w", id, storage2.ErrorNotFound)
		}
		return storage2.Product{}, fmt.Errorf("something went wrong - %w", storage2.CommonStorageError)
	}
	return result, nil
}
