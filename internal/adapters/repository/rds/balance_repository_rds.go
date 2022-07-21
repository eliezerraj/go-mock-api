package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/core/model"	
	"github.com/go-mock-api/internal/services"
	"github.com/go-mock-api/internal/adapters/repository"
)

type BalanceRepositoryRDSImpl struct {
	DatabaseHelper services.DatabaseHelper
}

func NewBalanceRepositoryRDS(databaseHelper services.DatabaseHelper) repository.BalanceRepository {
	return BalanceRepositoryRDSImpl{
		DatabaseHelper: databaseHelper,
	}
}

func (b BalanceRepositoryRDSImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Service).Info("Save") 

	client, _ := b.DatabaseHelper.GetConnection(ctx)

	stmt, err := client.Prepare(`INSERT INTO balance2 ( balance_id, 
														 account, 
														 amount, 
														 date_balance, 
														 Description) 
									VALUES( $1, $2, $3, $4, $5) `)
	if err != nil {
		loggers.GetLogger().Named(constants.Service).Info("Error") 
		return model.Balance{}, err
	}
	_, err = stmt.Exec(	balance.Id, 
						balance.Account,
						balance.Amount,
						time.Now(),
						balance.Description)

	return balance , nil
}

func (b BalanceRepositoryRDSImpl) List(ctx context.Context) ([]model.Balance, error) {
	fmt.Println("** IMPLEMENTAR **")
	result := []model.Balance{}
	m1 := model.Balance{ Id: "888"}
	result = append(result, m1)
	m2 := model.Balance{ Id: "777"}
	result = append(result, m2)

	return result , nil
}

func (b BalanceRepositoryRDSImpl) ListById(ctx context.Context, balance model.Balance) ([]model.Balance, error) {
	fmt.Println("** IMPLEMENTAR **")
	result := []model.Balance{}
	m1 := model.Balance{ Id: "888"}
	result = append(result, m1)
	m2 := model.Balance{ Id: "777"}
	result = append(result, m2)

	return result , nil
}

func (b BalanceRepositoryRDSImpl) FindById(ctx context.Context, id string) (model.Balance, error) {
	fmt.Println("** IMPLEMENTAR **")
	result := model.Balance{ Id: "888"}

	return result , nil
}