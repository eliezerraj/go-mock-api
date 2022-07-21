package config

import (
	"net"
	"os"
	"strconv"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
	"go.uber.org/zap"
	"github.com/go-mock-api/internal/viper"
)

func PodInfo() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		loggers.GetLogger().Named(constants.Config).Panic("Error to get the POD IP address", zap.Error(err))
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				viper.Application.App.IpAdress = ipnet.IP.String()
			}
		}
	}
	viper.Application.App.OSPID = strconv.Itoa(os.Getpid())
}
