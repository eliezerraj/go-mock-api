package rds

import (
	"context"
	"fmt"

	"github.com/go-mock-api/internal/core/model"
)

type BalanceRepository interface {
	FindById(ctx context.Context, id string) (model.Balance, error)
	ListById(ctx context.Context, balance model.Balance) ([]model.Balance, error)
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}

type BalanceRepositoryImpl struct {
}

func NewBalanceRepository() BalanceRepository {
	return BalanceRepositoryImpl{
	}
}

func (b BalanceRepositoryImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	fmt.Println("** IMPLEMENTAR **")
	return balance , nil
}

func (b BalanceRepositoryImpl) List(ctx context.Context) ([]model.Balance, error) {
	fmt.Println("** IMPLEMENTAR **")
	result := []model.Balance{}
	m1 := model.Balance{ Id: "888"}
	result = append(result, m1)
	m2 := model.Balance{ Id: "777"}
	result = append(result, m2)

	return result , nil
}

func (b BalanceRepositoryImpl) ListById(ctx context.Context, balance model.Balance) ([]model.Balance, error) {
	fmt.Println("** IMPLEMENTAR **")
	result := []model.Balance{}
	m1 := model.Balance{ Id: "888"}
	result = append(result, m1)
	m2 := model.Balance{ Id: "777"}
	result = append(result, m2)

	return result , nil
}