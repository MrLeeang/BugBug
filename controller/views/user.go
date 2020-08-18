package views

import (
	"github.com/gin-gonic/gin"
)

// ActionUserInfo 用户信息
func ActionUserInfo(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	userId := c.Param("userId")

	ret["userId"] = userId

	c.JSON(200, gin.H{
		"data":       ret,
		"error_code": 0,
		"msg":        "success.",
	})
}

// ActionUserLogin 登录
func ActionUserLogin(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}

// ActionUserLogin 关注
func ActionUpdateUserInfo(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}

// ActionAddUserFollow 关注
func ActionAddUserFollow(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}

// ActionDelUserFollow 取关
func ActionDelUserFollow(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}

// ActionUserFollowList 关注列表
func ActionUserFollowList(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}

// ActionUserFansList 粉丝列表
func ActionUserFansList(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}

// ActionDelUserFans 删除粉丝
func ActionDelUserFans(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}
