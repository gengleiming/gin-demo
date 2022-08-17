package global

import (
	"gin-config/config"
	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
}

var app = new(Application)

var Config = app.Config
var Viper = app.ConfigViper
