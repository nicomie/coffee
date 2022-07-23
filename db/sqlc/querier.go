// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateCoffee(ctx context.Context, arg CreateCoffeeParams) (Coffee, error)
	DeleteCoffee(ctx context.Context, id int64) error
	GetCoffee(ctx context.Context, id int64) (Coffee, error)
	ListCoffees(ctx context.Context) ([]Coffee, error)
	UpdateCoffee(ctx context.Context, arg UpdateCoffeeParams) (Coffee, error)
}

var _ Querier = (*Queries)(nil)
