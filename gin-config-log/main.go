package main

import (
	"fmt"
	"gin-config-log/bootstrap"
	"gin-config-log/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志
	global.Logger = bootstrap.InitializeLog()
	global.Slog = global.Logger.Sugar()

	// ---------------------------- 测试 ------------------------------------------
	// [2022-08-17 15:51:33.674]	cus.info	gin-config-log/main.go:18	log init success!
	global.Logger.Info("log init success!")
	//[2022-08-17 15:51:33.708]	cus.info	gin-config-log/main.go:21	log 	{"age": 18}
	global.Logger.Info("log ", zap.Int("age", 18))
	// [2022-08-17 15:51:33.708]	cus.info	gin-config-log/main.go:24	Infof() use Sprintf: 18
	global.Slog.Infof("Infof() use Sprintf: %d", 18)
	// [2022-08-17 15:51:33.708]	cus.info	gin-config-log/main.go:25	Infow() allows tags	{"name": "LeiMing", "type": 1}
	global.Slog.Infow("Infow() allows tags", "name", "LeiMing", "type", 1)

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	fmt.Println(global.Viper.GetString("app.env"), global.Config.App.Env)

	// 启动服务器
	r.Run(":" + global.Config.App.Port)
}
