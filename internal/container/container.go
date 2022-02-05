package container

import (
	"sync"

	"github.com/go-mock-api/internal/services"
	"github.com/go-mock-api/internal/adapters/repository/memkv"
	"github.com/go-mock-api/internal/adapters/repository/dynamoDB"
	"github.com/go-mock-api/internal/adapters/repository"
)

var container *ServiceContainer
var once sync.Once
var db_repo repository.BalanceRepository

type ServiceContainer struct {
	BalanceService        services.BalanceService
	//LogManager           services.LogManager
}

func Container() *ServiceContainer {
	once.Do(func() {

		var db_type = "dynamo"

		if db_type == "dynamo" {
			table_name := "test"
			db_repo, _ = dynamoDB.NewBalanceRepositoryDynamoDB(&table_name)
		} else {
			db_repo = memkv.NewBalanceRepositoryMemKv()
		}
		container = &ServiceContainer{
			BalanceService:  newBalanceService(db_repo),
		}
	})
	return container
}

func newBalanceService(db_repo repository.BalanceRepository) services.BalanceService {
	return services.NewBalanceService(db_repo)
}