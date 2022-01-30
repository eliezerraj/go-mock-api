package http_servers

import (
	"time"
	"net/http"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/handlers/http/routers"
	"github.com/go-mock-api/internal/viper"

)

type HttpServer struct {
	start time.Time
}

// NewHttpServer :
func NewHttpServer(start time.Time) HttpServer {
	return HttpServer{start: start}
}

// StartHttpServer :
func (s HttpServer) StartHttpServer() {

	loggers.GetLogger().Named(constants.Server).Info("Start server HTTP")

	r := routers.NewRouter()
	loggers.GetLogger().Named(constants.Server).Info("Created routers")

	duration := time.Since(s.start).Nanoseconds()

	loggers.GetLogger().Named(constants.Server).Info("Server booting",zap.Int("port", viper.Application.Server.Port))
	loggers.GetLogger().Named(constants.Server).Info("Server HTTP started",	zap.Int64("duration", duration))
	
	e := http.ListenAndServe(fmt.Sprintf(":%d", viper.Application.Server.Port), r.Router)
	if e != nil {
		loggers.GetLogger().Named(constants.Server).Panic("Internal error",	zap.Error(e))
	}
}