package repository

import (
	"context"

	"github.com/go-mock-api/internal/core/model"

)

type BalanceRepository interface {
	FindById(ctx context.Context, id string) (model.Balance, error)
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}

type BalanceRepositoryImpl struct {
}

func NewBalanceRepository() BalanceRepository {
	return BalanceRepositoryImpl{
	}
}

func (b BalanceRepositoryImpl) FindById(ctx context.Context, id string) (model.Balance, error) {
	return model.Balance{} , nil
}

func (b BalanceRepositoryImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	return model.Balance{} , nil
}

func (b BalanceRepositoryImpl) List(ctx context.Context) ([]model.Balance, error) {
	return []model.Balance{} , nil
}