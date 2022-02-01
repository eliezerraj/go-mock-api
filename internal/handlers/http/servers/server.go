package http_servers

import (
	"time"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"context"
	"strconv"

	"go.uber.org/zap"

	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/handlers/http/routers"
	"github.com/go-mock-api/internal/viper"

)

type HttpServer struct {
	start time.Time
}

func NewHttpServer(start time.Time) HttpServer {
	return HttpServer{start: start}
}

func (s HttpServer) StartHttpServer() {
	loggers.GetLogger().Named(constants.Server).Info("Start server HTTP")

	r := routers.NewRouter()
	loggers.GetLogger().Named(constants.Server).Info("Created routers")

	duration := time.Since(s.start).Nanoseconds()

	loggers.GetLogger().Named(constants.Server).Info("Server booting",zap.Int("port", viper.Application.Server.Port))
	loggers.GetLogger().Named(constants.Server).Info("Server HTTP started",	zap.Int64("duration", duration))
	
	srv := http.Server{
		Addr:         ":" +  strconv.Itoa(viper.Application.Server.Port),      	
		Handler:      r.Router,                	          
		ReadTimeout:  time.Duration(viper.Application.Server.ReadTimeout) * time.Second,   
		WriteTimeout: time.Duration(viper.Application.Server.WriteTimeout) * time.Second,  
		IdleTimeout:  time.Duration(viper.Application.Server.IdleTimeout) * time.Second, 
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			loggers.GetLogger().Named(constants.Server).Panic("Internal error",	zap.Error(err))
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	loggers.GetLogger().Named(constants.Server).Info("Stopping Server")

	ctx , cancel := context.WithTimeout(context.Background(), time.Duration(viper.Application.Server.CtxTimeout) * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		loggers.GetLogger().Named(constants.Server).Info("WARNING Dirty Shutdown", zap.Error(err))
		return
	}
	loggers.GetLogger().Named(constants.Server).Info("Stop DONE !")
}