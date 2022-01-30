package config

import (
	"time"

	_server "github.com/go-mock-api/internal/handlers/http/servers"

)

// ServerConfiguration :
func ServerConfiguration(start time.Time) {
	httpServer := _server.NewHttpServer(start)
	httpServer.StartHttpServer()
}
