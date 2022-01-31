package repository

import (
	"context"
"fmt"

	"github.com/go-mock-api/internal/core/model"

)

type BalanceRepository interface {
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}

type BalanceRepositoryImpl struct {
	//databaseHelper databaseHelper.DatabaseHelper
}

func NewBalanceRepository() BalanceRepository {
	return BalanceRepositoryImpl{
	//	databaseHelper: helper,
	}
}

func (b BalanceRepositoryImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	fmt.Println("========> ",balance)
	return balance , nil
}

func (b BalanceRepositoryImpl) List(ctx context.Context) ([]model.Balance, error) {
	result := []model.Balance{}
	m1 := model.Balance{ Id: "888"}
	result = append(result, m1)
	m2 := model.Balance{ Id: "777"}
	result = append(result, m2)

	return result , nil
}