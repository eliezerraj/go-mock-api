package repository

import (
	"context"

	"github.com/go-mock-api/internal/core/model"
)

type BalanceRepository interface {
	FindById(ctx context.Context, id string) (model.Balance, error)
	ListById(ctx context.Context, balance model.Balance) ([]model.Balance, error)
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}
