package global

import (
	"gin-gorm/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	Slog        *zap.SugaredLogger
	DB          *gorm.DB
}

var app = new(Application)

var Config = app.Config
var Viper = app.ConfigViper
var Logger = app.Log
var Slog = app.Slog
var DB = app.DB
