package routers

import (
	//"fmt"

	"github.com/go-chi/chi"
	//"github.com/go-chi/render"

)

type Router interface {
	Route(r chi.Router)
	GetPath() string
}

type ChiRouter struct {
	Router          chi.Router
	routers         []Router
	managerRouters  []Router
	internalRouters []Router
}

func NewRouter() ChiRouter {
	chiRouter := ChiRouter{Router: chi.NewRouter()}

	chiRouter.initializeControllers()
	//chiRouter.configurationRouters()

	return chiRouter
}

func (c *ChiRouter) initializeControllers() {
	routes := NewRouterComponent()

	c.routers = []Router{
		routes.Health,
		routes.Templates,
	}

	c.managerRouters = []Router{
		routes.Management,
	}

	c.internalRouters = []Router{
		routes.Tenants,
	}
}