package main

import (
	"fmt"
	"gin-config/bootstrap"
	"gin-config/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	fmt.Println(global.Viper.GetString("app.env"), global.Config.App.Env)

	// 启动服务器
	r.Run(":" + global.Config.App.Port)
}
