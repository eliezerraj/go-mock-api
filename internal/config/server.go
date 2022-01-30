package config

import (
	"time"

	"github.com/go-mock-api/internal/handlers/http_server"

)

// ServerConfiguration :
func ServerConfiguration(start time.Time) {
	httpServer := http_server.NewHttpServer(start)
	httpServer.StartHttpServer()
}
