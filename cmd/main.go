package main

import (
	"time"

	_configs "github.com/go-mock-api/internal/configs"
	"github.com/go-mock-api/internal/viper"

)

func main(){
	_configs.LoggerConfiguration()
	viper.Configuration()
	_configs.PodInfo()
	_configs.ServerConfiguration(time.Now())
}