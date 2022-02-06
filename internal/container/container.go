package container

import (
	"sync"
	"os"

	"go.uber.org/zap"

	"github.com/go-mock-api/internal/services"
	"github.com/go-mock-api/internal/adapters/repository/memkv"
	"github.com/go-mock-api/internal/adapters/repository/dynamoDB"
	"github.com/go-mock-api/internal/adapters/repository"
	"github.com/go-mock-api/internal/exceptions"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
)

var container *ServiceContainer
var once sync.Once


type ServiceContainer struct {
	BalanceService        services.BalanceService
	//LogManager           services.LogManager
}

func Container() *ServiceContainer {
	loggers.GetLogger().Named(constants.Container).Info("Container") 

	once.Do(func() {
		var db_type = "dynamo"
		var db_repo repository.BalanceRepository
		
		if db_type == "dynamo" {
			table_name := "balance"
			_db_repo, err := dynamoDB.NewBalanceRepositoryDynamoDB(table_name)
			db_repo = _db_repo
			if err != nil {
				loggers.GetLogger().Named(constants.Viper).Panic("Error open database",zap.Error(exceptions.ErrOpenDataBase))
				os.Exit(1)
			}
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