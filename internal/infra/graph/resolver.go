package graph

import (
	"context"

	"github.com/cgalimberti/gocleanarc/20-CleanArch/internal/infra/graph/model"
	"github.com/cgalimberti/gocleanarc/20-CleanArch/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return r }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return r }

// CreateOrder is the resolver for the createOrder field.
func (r *Resolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}

	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID:         output.ID,
		Price:      output.Price,
		Tax:        output.Tax,
		FinalPrice: output.FinalPrice,
	}, nil
}

// Orders is the resolver for the orders field.
func (r *Resolver) Orders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var result []*model.Order
	for _, order := range orders {
		result = append(result, &model.Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return result, nil
}
