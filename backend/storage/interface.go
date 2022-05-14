package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	CommonStorageError = errors.New("storage")
	ErrorNotFound      = fmt.Errorf("%w.not_found", CommonStorageError)
)

type Product struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Price    string             `json:"price" bson:"price"`
	Quantity string             `json:"quantity" bson:"quantity"`
}

type Storage interface {
	Create(ctx context.Context, data Product) (Product, error)
	GetProduct(ctx context.Context, name string) (Product, error)
	Update(ctx context.Context, data Product) (Product, error)
	//Delete(ctx context.Context, data Product) error
	List(ctx context.Context) ([]Product, error)
}
