package main

import (
	"fmt"
	"time"

	"github.com/go-mock-api/internal/config"
	"github.com/go-mock-api/internal/viper"

)

func main(){
	fmt.Println("teste")
	config.LoggerConfiguration()
	viper.Configuration()
	config.ServerConfiguration(time.Now())
}