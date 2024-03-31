package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/NaturalBug/patients-api/graph/model"
)

// CreatePatient is the resolver for the createPatient field.
func (r *mutationResolver) CreatePatient(ctx context.Context, input *model.NewPatient) (*model.Patient, error) {
	panic(fmt.Errorf("not implemented: CreatePatient - createPatient"))
}

// UpdatePatient is the resolver for the updatePatient field.
func (r *mutationResolver) UpdatePatient(ctx context.Context, input *model.UpdatePatient) (*model.Patient, error) {
	panic(fmt.Errorf("not implemented: UpdatePatient - updatePatient"))
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.NewOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, input *model.UpdateOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: UpdateOrder - updateOrder"))
}

// Patients is the resolver for the patients field.
func (r *queryResolver) Patients(ctx context.Context) ([]*model.Patient, error) {
	panic(fmt.Errorf("not implemented: Patients - patients"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
