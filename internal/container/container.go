package container

import (
	"sync"

	"github.com/go-mock-api/internal/services"

)


var container *ServiceContainer
var once sync.Once

type ServiceContainer struct {
	BalanceService        services.BalanceService
	//LogManager           services.LogManager
}

func Container() *ServiceContainer {
	once.Do(func() {
		container = &ServiceContainer{
			//LogManager:  nil,
			BalanceService:  newBalanceService(),
		}
	})
	return container
}

func newBalanceService() services.BalanceService {
	//travelRepository := repository.NewTravelsRepository(dbh)
	return services.NewBalanceService()
}