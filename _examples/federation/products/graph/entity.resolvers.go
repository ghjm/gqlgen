package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40-dev

import (
	"context"

	"github.com/99designs/gqlgen/_examples/federation/products/graph/model"
)

// FindManufacturerByID is the resolver for the findManufacturerByID field.
func (r *entityResolver) FindManufacturerByID(ctx context.Context, id string) (*model.Manufacturer, error) {
	return &model.Manufacturer{
		ID:   id,
		Name: "Millinery " + id,
	}, nil
}

// FindProductByManufacturerIDAndID is the resolver for the findProductByManufacturerIDAndID field.
func (r *entityResolver) FindProductByManufacturerIDAndID(ctx context.Context, manufacturerID string, id string) (*model.Product, error) {
	for _, hat := range hats {
		if hat.ID == id && hat.Manufacturer.ID == manufacturerID {
			return hat, nil
		}
	}
	return nil, nil
}

// FindProductByUpc is the resolver for the findProductByUpc field.
func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	for _, hat := range hats {
		if hat.Upc == upc {
			return hat, nil
		}
	}
	return nil, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
