package main

import (
	"BugBug/controller"
	"BugBug/middleware"
	"BugBug/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	r := gin.New()

	// 注册logger
	r.Use(middleware.Logger())
	// 跨域
	r.Use(middleware.Cors())
	// 注册路由
	controller.MakeRouter(r)
	// 启动
	utils.UtilsLogger.Info("Server Run Success: 0.0.0.0:9501")
	_ = r.Run(":9501")
}
