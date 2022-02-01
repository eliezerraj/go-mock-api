package viper

import (
	"go.uber.org/zap"
	"github.com/spf13/viper"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/core/model"

)

var (
	Application model.ManagerInfo
	Viper *viper.Viper
)

const (
	_applicationFileName = "application"
	_extension           = "yml"
	_resourcePath        = "../resources/"
)

//Configuration godoc
func Configuration() {
	var app model.ManagerInfo

	viper.SetConfigName(_applicationFileName)
	viper.SetConfigType(_extension)
	viper.AddConfigPath(_resourcePath)
	viper.ReadInConfig()

	errUnmarshal := viper.Unmarshal(&app)
	if errUnmarshal != nil {
		loggers.GetLogger().Named(constants.Viper).Panic("Parse error for application structure", zap.Error(errUnmarshal))
	}
	loggers.GetLogger().Named(constants.Viper).Info("Data APPLICATION.YML => ",zap.Any("name : ", app))

	Application = app
}