package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	CommonStorageError = errors.New("storage")
	ErrorCollision     = fmt.Errorf("%w.collision", CommonStorageError)
	ErrorNotFound      = fmt.Errorf("%w.not_found", CommonStorageError)
)

type Product struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Price    float32            `json:"price" bson:"price"`
	Quantity int32              `json:"quantity" bson:"quantity"`
}

type PostsByUser struct {
	Posts      []Product          `json:"posts" bson:"posts"`
	NextPageId primitive.ObjectID `json:"nextPage" bson:"nextPage"`
}

type Storage interface {
	Create(ctx context.Context, data Product) error
	GetProduct(ctx context.Context, name string) (Product, error)
	//Update(ctx context.Context, data Product) error
	//Delete(ctx context.Context, data Product) error
	//List(ctx context.Context, data Product) ([]Product, error)
}
