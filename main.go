package main

import (
	"BugBug/controller"
	"BugBug/middleware"
	"BugBug/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	r := gin.New()
	// session redis
	store, _ := redis.NewStore(10, "tcp", utils.SessionRedisAddress, utils.RedisPassword, []byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 300,
		Path:   "/GinSession",
	})
	// 注册session
	r.Use(sessions.Sessions("GinSession", store))
	// 注册logger
	r.Use(middleware.Logger())
	// 注册路由
	controller.MakeRouter(r)
	// 启动
	utils.UtilsLogger.Info("Server Run Success: 0.0.0.0:9501")
	_ = r.Run(":9501")
}
