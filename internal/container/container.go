package container

import (
	"sync"

)


var container *ServiceContainer
var once sync.Once

type ServiceContainer struct {
//LogManager           services.LogManager
}

func Container() *ServiceContainer {

	once.Do(func() {
		container = &ServiceContainer{
			//LogManager:  nil,
		}
	})
	return container
}

