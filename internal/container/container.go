package container

import (
	"sync"
	"os"

	"go.uber.org/zap"

	"github.com/go-mock-api/internal/services"
	"github.com/go-mock-api/internal/adapters/repository/memkv"
	"github.com/go-mock-api/internal/adapters/repository/dynamoDB"
	"github.com/go-mock-api/internal/adapters/repository/rds"
	"github.com/go-mock-api/internal/adapters/repository"
	"github.com/go-mock-api/internal/exceptions"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/viper"
)

var Acontainer *ServiceContainer
var once sync.Once

var db_helper services.DatabaseHelper

type ServiceContainer struct {
	BalanceService        services.BalanceService
}

func Container() *ServiceContainer {
	loggers.GetLogger().Named(constants.Container).Info("Container") 

	once.Do(func() {
		var err error
		var db_repo repository.BalanceRepository
		if viper.Application.Setup.DatabaseType == "dynamo" {
			table_name := "balance"
			db_repo, err = dynamoDB.NewBalanceRepositoryDynamoDB(table_name)
			if err != nil {
				loggers.GetLogger().Named(constants.Viper).Panic("Error open NOSQL database",zap.Error(exceptions.ErrOpenDataBase))
				os.Exit(1)
			}
		} else if viper.Application.Setup.DatabaseType == "rds" {
			db_helper, err = services.NewDatabaseHelper(viper.Application.DatabaseRDS)
			if err != nil {
				loggers.GetLogger().Named(constants.Viper).Panic("Error open RDS database",zap.Error(exceptions.ErrOpenDataBase))
				os.Exit(1)
			}
			db_repo = rds.NewBalanceRepositoryRDS(db_helper)
		} else {
			db_repo = memkv.NewBalanceRepositoryMemKv()
		}
		
		Acontainer = &ServiceContainer{
			BalanceService:  newBalanceService(db_repo),
		}
	})
	return Acontainer
}

func newBalanceService(db_repo repository.BalanceRepository) services.BalanceService {
	return services.NewBalanceService(db_repo)
}

func (c ServiceContainer) ContainerShutdown() {
	loggers.GetLogger().Named(constants.Container).Info("ContainerShutdown !!!") 
	if viper.Application.Setup.DatabaseType == "rds" {
		db_helper.CloseConnection()
	}
}