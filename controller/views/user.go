package views

import (
	"BugBug/service"
	"BugBug/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// ActionUserInfo 用户信息
func ActionUserInfo(c *gin.Context) {

	userId := c.Param("userId")

	// 用户信息
	userInfo := service.GetUserById(userId)
	if userInfo["id"] != nil {
		// 点赞数量
		userInfo["vote_count"] = service.CountVoteByUserId(userId)
		// 采纳数量
		userInfo["adopt_count"] = service.CountAdoptByUserId(userId)
	}

	c.JSON(200, gin.H{
		"data":       userInfo,
		"error_code": 0,
		"msg":        "success.",
	})
}

// ActionUserLogin 登录
func ActionUserLogin(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}
	phone := c.Query("phone")
	code := c.Query("code")
	// 验证验证码
	verifyRet := service.VerifyLoginCode(phone, code)
	if verifyRet == false {
		//c.JSON(200, gin.H{
		//	"error_code": 11002,
		//	"msg": "验证码不正确",
		//})
		//return
	}
	// 获取用户信息
	userInfo := service.GetUserByPhone(phone)
	if userInfo["id"] == nil {
		c.JSON(200, gin.H{
			"error_code": 11002,
			"msg":        "用户不存在",
		})
		return
	}
	uid := userInfo["id"]
	// 生成token
	tokenStr := service.GenerateToken(uid)
	utils.UtilsLogger.Info(tokenStr)
	ret["token"] = tokenStr
	ret["nickname"] = userInfo["nickname"]
	ret["expire_time"] = int64(time.Now().Add(time.Hour * 72).Unix())
	ret["avatar"] = userInfo["avatar"]
	ret["uid"] = userInfo["id"]
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "登录成功",
		"data":       ret,
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
