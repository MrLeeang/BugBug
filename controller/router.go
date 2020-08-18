package controller

import (
	"HackInBox.online/controller/views"
	"github.com/gin-gonic/gin"
)

// MakRouter 是路由
func MakRouter(r *gin.Engine) {
	// 用户信息
	r.GET("/v1/user/:userId/info", views.ActionUserInfo)
	// 用户登录
	r.POST("/v1/user/login", views.ActionUserLogin)
	// 修改用户信息
	r.PUT("/v1/user/update", views.AuthHandler(), views.ActionUpdateUserInfo)
	// 关注
	r.POST("/v1/user/follow/:userId", views.AuthHandler(), views.ActionAddUserFollow)
	// 取关
	r.DELETE("/v1/user/follow/:userId", views.AuthHandler(), views.ActionDelUserFollow)
	// 关注列表
	r.GET("/v1/user/:userId/follows", views.AuthHandler(), views.ActionUserFollowList)
	// 粉丝列表
	r.GET("/v1/user/:userId/fans", views.AuthHandler(), views.ActionUserFansList)
	// 删除粉丝
	r.DELETE("/v1/user/fans/:userId", views.AuthHandler(), views.ActionDelUserFans)
}
