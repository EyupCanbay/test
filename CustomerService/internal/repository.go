package internal

import (
	"context"
	"fmt"
	"tesodev-korpes/CustomerService/internal/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(col *mongo.Collection) *Repository {
	return &Repository{
		collection: col,
	}
}

func (r *Repository) FindByID(ctx context.Context, id string) (*types.Customer, error) {
	var customer *types.Customer
	return customer, nil
}

func (r *Repository) Create(ctx context.Context, customer *types.Customer) (string, error) {
	println(customer.FirstName)
	result, err := r.collection.InsertOne(ctx, customer)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", fmt.Errorf("result is empty")
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("Invalid id type")
	}

	return id.String(), nil
}

func (r *Repository) Update(ctx context.Context, id string, update interface{}) error {
	// Placeholder method
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	// Placeholder method
	return nil
}
