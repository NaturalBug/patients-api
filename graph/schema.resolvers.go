package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"strconv"

	"github.com/NaturalBug/patients-api/graph/model"
)

// CreatePatient is the resolver for the createPatient field.
func (r *mutationResolver) CreatePatient(ctx context.Context, input *model.CreatePatientInput) (*model.Patient, error) {
	var lastId = 0
	lengthOfPatients := len(r.patients)
	if lengthOfPatients != 0 {
		lastId, _ = strconv.Atoi(r.patients[lengthOfPatients-1].ID)
	}
	newId := strconv.Itoa(lastId + 1)

	patient := &model.Patient{
		ID:   newId,
		Name: input.Name,
	}
	r.patients = append(r.patients, patient)

	return patient, nil
}

// UpdatePatient is the resolver for the updatePatient field.
func (r *mutationResolver) UpdatePatient(ctx context.Context, input *model.UpdatePatientInput) (*model.Patient, error) {
	var patient *model.Patient

	for i := 0; i < len(r.patients); i++ {
		if r.patients[i].ID == input.ID {
			patient = r.patients[i]
			break
		}
	}

	if input.Name != nil {
		patient.Name = *input.Name
	}

	if input.OrderID != nil {
		patient.OrderID = input.OrderID
	}

	return patient, nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.CreateOrderInput) (*model.Order, error) {
	var lastId = 0
	lengthOfOrders := len(r.orders)
	if lengthOfOrders != 0 {
		lastId, _ = strconv.Atoi(r.orders[lengthOfOrders-1].ID)
	}
	newId := strconv.Itoa(lastId + 1)

	order := &model.Order{
		ID:      newId,
		Message: &input.Message,
	}

	r.orders = append(r.orders, order)

	return order, nil
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, input *model.UpdateOrderInput) (*model.Order, error) {
	var order *model.Order

	for i := 0; i < len(r.orders); i++ {
		if r.orders[i].ID == input.ID {
			order = r.orders[i]
			break
		}
	}

	order.Message = &input.Message

	return order, nil
}

// Patients is the resolver for the patients field.
func (r *queryResolver) Patients(ctx context.Context, id *string) ([]*model.Patient, error) {
	if id == nil {
		return r.patients, nil
	}

	for i := 0; i < len(r.patients); i++ {
		if r.patients[i].ID == *id {
			return []*model.Patient{r.patients[i]}, nil
		}
	}

	return nil, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, id *string) ([]*model.Order, error) {
	if id == nil {
		return r.orders, nil
	}

	for i := 0; i < len(r.orders); i++ {
		if r.orders[i].ID == *id {
			return []*model.Order{r.orders[i]}, nil
		}
	}

	return nil, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
