package memkv

import (
	"sync"
	"context"
	//"fmt"
	"encoding/json"

	"go.uber.org/zap"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/core/model"
	"github.com/go-mock-api/internal/exceptions"
)

var mutex sync.Mutex

type BalanceRepositoryMemKv interface {
	FindById(ctx context.Context, id string) (model.Balance, error)
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}

type BalanceRepositoryMemKvImpl struct {
	kv map[string][]byte
}

func NewBalanceRepositoryMemKv() BalanceRepositoryMemKv {
	return BalanceRepositoryMemKvImpl{
		kv: map[string][]byte{},
	}
}

func (b BalanceRepositoryMemKvImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("Save") 
	mutex.Lock()
	defer mutex.Unlock()

	bytes, err := json.Marshal(balance)
	if err != nil {
		return model.Balance{}, exceptions.Throw(err, exceptions.ErrSaveDatabase)
	}
	b.kv[balance.Id] = bytes
	loggers.GetLogger().Named(constants.Database).Info("TABLE Balance MEMKV", zap.Any("count :" ,len(b.kv)) )
	
	return balance , nil
}

func (b BalanceRepositoryMemKvImpl) List(ctx context.Context) ([]model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("List") 
	var result []model.Balance
	for _, value := range b.kv {
		balance := model.Balance{}
		err := json.Unmarshal(value, &balance)
		if err != nil {
			return []model.Balance{}, exceptions.Throw(err, exceptions.ErrList)
		}
		result = append(result, balance)
	}
	return result, nil
}

func (b BalanceRepositoryMemKvImpl) FindById(ctx context.Context, id string) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("FindById") 
	if value, ok := b.kv[id]; ok {
		balance := model.Balance{}
		err := json.Unmarshal(value, &balance)
		if err != nil {
			return model.Balance{}, exceptions.Throw(err, exceptions.ErrJsonCode)
		}
		return balance, nil
	}
	return model.Balance{}, exceptions.Throw( exceptions.ErrNoDataFound, exceptions.ErrNoDataFound)
}
