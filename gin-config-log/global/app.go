package global

import (
	"gin-config-log/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	Slog        *zap.SugaredLogger
}

var app = new(Application)

var Config = app.Config
var Viper = app.ConfigViper
var Logger = app.Log
var Slog = app.Slog
