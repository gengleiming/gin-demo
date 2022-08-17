package main

import (
	"fmt"
	"gin-gorm/bootstrap"
	"gin-gorm/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志
	global.Logger = bootstrap.InitializeLog()
	global.Slog = global.Logger.Sugar()

	// 初始化数据库
	global.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.DB != nil {
			db, _ := global.DB.DB()
			db.Close()
		}
	}()

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	fmt.Println(global.Viper.GetString("app.env"), global.Config.App.Env)

	// 启动服务器
	r.Run(":" + global.Config.App.Port)
}
