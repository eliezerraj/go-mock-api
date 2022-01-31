package memkv

import (
	//"sync"
	"context"
	"fmt"

	"github.com/go-mock-api/internal/core/model"
)

type BalanceRepositoryMemKv interface {
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}

type BalanceRepositoryMemKvImpl struct {
	//kv map[string][]byte
}

func NewBalanceRepositoryMemKv() BalanceRepositoryMemKv {
	return BalanceRepositoryMemKvImpl{
	//	databaseHelper: helper,
	}
}

func (b BalanceRepositoryMemKvImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	fmt.Println("=====www===> ",balance)
	return balance , nil
}

func (b BalanceRepositoryMemKvImpl) List(ctx context.Context) ([]model.Balance, error) {
	result := []model.Balance{}
	m1 := model.Balance{ Id: "888"}
	result = append(result, m1)
	m2 := model.Balance{ Id: "777"}
	result = append(result, m2)

	return result , nil
}