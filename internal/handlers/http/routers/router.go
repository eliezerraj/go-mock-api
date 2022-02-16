package routers

import (
	"sync"
	"net/http"
	"fmt"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/go-chi/render"

	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/handlers/http/routers/controllers"
	"github.com/go-mock-api/internal/handlers/http/handlers"
	"github.com/go-mock-api/internal/container"
	"github.com/go-mock-api/internal/handlers/http/middleware"
)

//----------------------------------------
var once sync.Once
var controller *ControllerImpl

type Controller interface {
	ManagerController() controllers.Management
	BalanceController() controllers.Balance
}

type ControllerImpl struct {
	container *container.ServiceContainer
}

func GetInstance() Controller {
	loggers.GetLogger().Named(constants.Router).Info(" ==> GetInstance() Controller")
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

func (c *ControllerImpl) BalanceController() controllers.Balance {
	return controllers.NewBalanceController(handlers.GetRequestHandlersInstance(), 
											handlers.GetResponseHandlersInstance(),
											c.container.BalanceService,
											middleware.NewValidatorMiddleware(handlers.GetResponseHandlersInstance()))
}

//------------------------------------
type Router interface {
	Route(r chi.Router)
	GetPath() string
}

type RouterComponent struct {
	Management     Router
	Balance        Router
}

func NewRouterComponent() RouterComponent {
	loggers.GetLogger().Named(constants.Router).Info("NewRouterComponent")
	return RouterComponent{
		Management: GetInstance().ManagerController(),
		Balance: 	GetInstance().BalanceController(),
	}
}
//----------------------------------
type ChiRouter struct {
	Router          chi.Router
	managerRouters  []Router
	serviceRouters	[]Router
}

func NewRouter() ChiRouter {
	loggers.GetLogger().Named(constants.Router).Info("NewRouter")
	chiRouter := ChiRouter{Router: chi.NewRouter()}

	chiRouter.initializeControllers()
	chiRouter.configurationRouters()

	return chiRouter
}

func (c *ChiRouter) initializeControllers() {
	loggers.GetLogger().Named(constants.Router).Info("initializeControllers")
	routes := NewRouterComponent()

	c.managerRouters = []Router{
		routes.Management,
	}
	c.serviceRouters = []Router{
		routes.Balance,
	}
}

func (c ChiRouter) Cors() func(next http.Handler) http.Handler {
	loggers.GetLogger().Named(constants.Router).Info("Configuring Cors")
	return cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposedHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials:   true,
		MaxAge:             300, // Maximum value not ignored by any of major browsers
		OptionsPassthrough: false,
	}).Handler
}

func (c ChiRouter) configurationRouters() {
	loggers.GetLogger().Named(constants.Router).Info("configurationRouters")

	c.Router.Use(c.Cors())
	managementMiddleware := middleware.NewManagementMiddleware(handlers.GetResponseHandlersInstance())

	c.Router.Group(func(rManager chi.Router) {
		rManager.Use(render.SetContentType(render.ContentTypeJSON))
		for _, router := range c.managerRouters {
			loggers.GetLogger().Named(constants.Router).Info(fmt.Sprintf("Router %s created", router.GetPath()))
			rManager.Route(router.GetPath(), router.Route)
		}
	})

	c.Router.Group(func(rService chi.Router) {
		rService.Use(render.SetContentType(render.ContentTypeJSON))
		rService.Use(managementMiddleware.Management)
		for _, router := range c.serviceRouters {
			loggers.GetLogger().Named(constants.Router).Info(fmt.Sprintf("Router %s created", router.GetPath()))
			rService.Route(router.GetPath(), router.Route)
		}
	})
}

func (c *ChiRouter) ShutdownControllers() {
	container.Acontainer.ContainerShutdown()
}
