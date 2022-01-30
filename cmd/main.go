package main

import (
	"fmt"
	"time"

	"github.com/go-mock-api/internal/configs"
	"github.com/go-mock-api/internal/viper"

)

func main(){
	fmt.Println("teste")
	configs.LoggerConfiguration()
	viper.Configuration()
	configs.ServerConfiguration(time.Now())
}