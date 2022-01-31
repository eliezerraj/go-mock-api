package services

import (
	"context"

	"github.com/go-mock-api/internal/core/model"
	"github.com/go-mock-api/internal/adapters/repository"
	"github.com/go-mock-api/internal/exceptions"

)

type BalanceService interface {
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}

type BalanceServiceImpl struct {
	repository repository.BalanceRepository
}

func NewBalanceService(repository repository.BalanceRepository) BalanceService {
	return BalanceServiceImpl{
		repository: repository,
	}
}

func (t BalanceServiceImpl) List(ctx context.Context) ([]model.Balance, error) {
	result, err := t.repository.List(ctx)
	if err != nil {
		return []model.Balance{} , exceptions.Throw(err, exceptions.ErrList)
	}
	return result, nil
}

func (t BalanceServiceImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	result, err := t.repository.Save(ctx, balance)
	if err != nil {
		return model.Balance{} , exceptions.Throw(err, exceptions.ErrSave)
	}
	return result, nil
}