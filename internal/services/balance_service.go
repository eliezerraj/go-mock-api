package services

import (
	"context"

	"github.com/go-mock-api/internal/core/model"

)

type BalanceService interface {
	List(ctx context.Context) ([]model.Balance, error)
}

type BalanceServiceImpl struct {
	//repository repository.TravelsRepository
}

func NewBalanceService() BalanceService {
	return BalanceServiceImpl{
	//	repository: repository,
	}
}

func (t BalanceServiceImpl) List(ctx context.Context) ([]model.Balance, error) {

	result := []model.Balance{}
	m1 := model.Balance{ Id: "0099"}
	result = append(result, m1)

	return result, nil
}