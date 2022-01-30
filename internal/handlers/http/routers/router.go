package routers

import (
	"sync"

	"github.com/go-chi/chi"

	"github.com/go-mock-api/internal/handlers/http/routers/controllers"
	"github.com/go-mock-api/internal/handlers/http/handlers"
	"github.com/go-mock-api/internal/container"

)

//----------------------------------------
var once sync.Once
var controller *ControllerImpl

type Controller interface {
	ManagerController() controllers.Management
}

type ControllerImpl struct {
	container *container.ServiceContainer
}

func GetInstance() Controller {
	once.Do(func() {
		controller = &ControllerImpl{
			container: container.Container(),
		}
	})
	return controller
}

func (c *ControllerImpl) ManagerController() controllers.Management {
	return controllers.NewManagementController( handlers.GetRequestHandlersInstance(),
												handlers.GetResponseHandlersInstance())
}
//------------------------------------
type Router interface {
	Route(r chi.Router)
	GetPath() string
}

type RouterComponent struct {
	Management     Router
}

func NewRouterComponent() RouterComponent {
	return RouterComponent{
		Management: GetInstance().ManagerController(),
	}
}
//----------------------------------
type ChiRouter struct {
	Router          chi.Router
	managerRouters  []Router
}

func NewRouter() ChiRouter {
	chiRouter := ChiRouter{Router: chi.NewRouter()}

	chiRouter.initializeControllers()
	//chiRouter.configurationRouters()

	return chiRouter
}

func (c *ChiRouter) initializeControllers() {
	routes := NewRouterComponent()

	c.managerRouters = []Router{
		routes.Management,
	}
}

