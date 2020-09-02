package views

import (
	"BugBug/service"
	"BugBug/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ActionUserInfo 用户信息
func ActionUserInfo(c *gin.Context) {

	userID := c.Param("userId")

	// 用户信息
	userInfo := service.GetUserByID(userID)
	if userInfo["id"] != nil {
		// 点赞数量
		userInfo["vote_count"] = service.CountVoteByUserID(userID)
		// 采纳数量
		userInfo["adopt_count"] = service.CountAdoptByUserID(userID)
	}

	c.JSON(200, gin.H{
		"data":       userInfo,
		"error_code": 0,
		"msg":        "success.",
	})
}

// ActionUserList 用户列表
func ActionUserList(c *gin.Context) {

	// 用户信息
	allUserList := service.GetUsers()

	c.JSON(200, gin.H{
		"data":       allUserList,
		"error_code": 0,
		"msg":        "success.",
	})
}

// ActionUserLogin 登录
func ActionUserLogin(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	// 验证验证码
	// verifyRet := service.VerifyLoginCode(phone, code)
	// if verifyRet == false {
	// 	c.JSON(200, gin.H{
	// 		"error_code": 11002,
	// 		"msg":        "验证码不正确",
	// 	})
	// 	return
	// }
	// 获取用户信息
	utils.UtilsLogger.Info(code)
	userInfo := service.GetUserByPhone(phone)
	if userInfo["id"] == nil {
		c.JSON(200, gin.H{
			"error_code": 11002,
			"msg":        "用户不存在",
		})
		return
	}

	tokenStr := service.GenerateToken(userInfo["id"])
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

// ActionUpdateUserInfo 更新用户信息
func ActionUpdateUserInfo(c *gin.Context) {
	// 昵称
	nickName := c.PostForm("nickname")
	// 签名
	signature := c.PostForm("signature")
	// 性别
	sex := c.PostForm("sex")
	// 头像
	avatar := c.PostForm("avatar")

	sexInt, _ := strconv.Atoi(sex)
	// interface 转 string
	uid, _ := c.Keys["UID"].(string)
	// string 转成int64
	uidInt64, _ := strconv.ParseInt(uid, 10, 64)
	ok, userInfo := service.UpdateUserInfoByID(uidInt64, nickName, signature, sexInt, avatar)
	if !ok {
		c.JSON(200, gin.H{
			"error_code": 11002,
			"msg":        "更新失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "更新成功",
		"data":       userInfo,
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
