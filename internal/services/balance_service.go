package services

import (
	"context"

	"github.com/go-mock-api/internal/core/model"
	"github.com/go-mock-api/internal/adapters/repository"
	"github.com/go-mock-api/internal/exceptions"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"

)

type BalanceService interface {
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
	FindById(ctx context.Context, id string) (model.Balance, error)
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
	loggers.GetLogger().Named(constants.Service).Info("List") 
	result, err := t.repository.List(ctx)
	if err != nil {
		return []model.Balance{} , exceptions.Throw(err, exceptions.ErrList)
	}
	return result, nil
}

func (t BalanceServiceImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Service).Info("Save") 
	result, err := t.repository.Save(ctx, balance)
	if err != nil {
		return model.Balance{} , exceptions.Throw(err, exceptions.ErrSave)
	}
	return result, nil
}

func (t BalanceServiceImpl) FindById(ctx context.Context, id string) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Service).Info("FindById") 
	result, err := t.repository.FindById(ctx, id)
	if err != nil {
		return model.Balance{} , exceptions.Throw(err, exceptions.ErrNoDataFound)
	}
	return result, nil
}