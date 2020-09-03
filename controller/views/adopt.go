package views

import (
	"BugBug/service"
	"BugBug/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// ActionAdoptPost 采纳
func ActionAdoptPost(c *gin.Context) {
	pid := c.PostForm("pid")
	if pid == "" {
		c.JSON(200, gin.H{
			"error_code": 500,
			"msg":        "评论ID或贴子ID是必须的.",
		})
		return
	}

	pidInt64, _ := strconv.ParseInt(pid, 10, 64)

	postModel := service.GetPostByID(pidInt64)
	uidInt64 := postModel.Uid
	if uidInt64 == 0 {
		c.JSON(200, gin.H{
			"error_code": 7000,
			"msg":        "帖子不存在",
		})
		return
	}
	uid := c.Keys["UID"].(int64)

	ok := service.AdoptPost(uid, pidInt64)
	if !ok {
		c.JSON(200, gin.H{
			"error_code": 500,
			"msg":        "操作失败",
		})
		return
	}

	// 更新redis
	if uidInt64 != 0 {
		uidString := strconv.FormatInt(uidInt64, 10)
		voteNum, _ := redis.Int64(utils.RedisClient.Get(uidString + "adopt"))
		voteNum++
		utils.RedisClient.Set(uidString+"adopt", voteNum)
	}
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
	})
	return

}

// ActionUserAdoptList 采纳列表
func ActionUserAdoptList(c *gin.Context) {
	userID := c.Param("userId")
	posts := service.UserAdpotPostList(userID)
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
		"data":       posts,
	})
}

// GetAdoptListByUserPost 根据用户帖子获取点赞列表
func GetAdoptListByUserPost(c *gin.Context) {
	uid := c.Keys["UID"].(int64)
	page := c.Query("page")
	size := c.Query("size")

	if page == "" {
		page = "1"
	}
	if size == "" {
		size = "10"
	}
	uidStr := strconv.FormatInt(uid, 10)
	utils.RedisClient.Set(uidStr+"adopt", 0)
	data := service.GetAdoptListByUserPost(uid, page, size)
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
		"data":       data,
	})
}

// GetAdoptPostListByUser 我得采纳列表
func GetAdoptPostListByUser(c *gin.Context) {
	uid := c.Keys["UID"].(int64)
	page := c.Query("page")
	size := c.Query("size")

	if page == "" {
		page = "1"
	}
	if size == "" {
		size = "10"
	}

	data := service.GetAdoptListByUserPost(uid, page, size)
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
		"data":       data,
	})
}
