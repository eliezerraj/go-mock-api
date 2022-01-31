package container

import (
	"sync"

	"github.com/go-mock-api/internal/services"
	"github.com/go-mock-api/internal/adapters/repository/memkv"

)

var container *ServiceContainer
var once sync.Once

type ServiceContainer struct {
	BalanceService        services.BalanceService
	//LogManager           services.LogManager
}

func Container() *ServiceContainer {
	
	once.Do(func() {
		//db := tenants.NewDatabaseHelper()

		container = &ServiceContainer{
			//LogManager:  nil,
			BalanceService:  newBalanceService(),
		}
	})
	return container
}

func newBalanceService() services.BalanceService {
	balanceRepository_memkv := memkv.NewBalanceRepositoryMemKv()
	return services.NewBalanceService(balanceRepository_memkv)
}