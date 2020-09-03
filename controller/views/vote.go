package views

import (
	"BugBug/service"
	"BugBug/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// ActionVotePost 点赞
func ActionVotePost(c *gin.Context) {
	// 帖子得ip
	pid := c.PostForm("pid")
	// 回复的评论ID,可以为空
	pcid := c.PostForm("pcid")
	// 1:赞成,0:反对,默认为1
	result := c.PostForm("result")
	if result == "" {
		result = "1"
	}

	if pid == "" && pcid == "" {
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
	// string 转成int64
	pcidInt64, _ := strconv.ParseInt(pcid, 10, 64)
	ok := service.AddVote(uid, pidInt64, pcidInt64, result)
	if !ok {
		c.JSON(200, gin.H{
			"error_code": 7000,
			"msg":        "点赞失败",
		})
		return
	}
	// 更新redis

	utils.UtilsLogger.Info(uidInt64)
	if uidInt64 != 0 {
		uidString := strconv.FormatInt(uidInt64, 10)
		voteNum, _ := redis.Int64(utils.RedisClient.Get(uidString + "vote"))
		voteNum++
		utils.RedisClient.Set(uidString+"vote", voteNum)
	}
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "已点赞",
	})
}

// ActionVoteCancel 取消点赞
func ActionVoteCancel(c *gin.Context) {
	// 帖子得ip
	pid := c.PostForm("pid")
	// 回复的评论ID,可以为空
	pcid := c.PostForm("pcid")
	if pid == "" && pcid == "" {
		c.JSON(200, gin.H{
			"error_code": 500,
			"msg":        "评论ID或贴子ID是必须的.",
		})
		return
	}
	var ok bool
	uid := c.Keys["UID"].(int64)
	pidInt64, _ := strconv.ParseInt(pid, 10, 64)
	pcidInt64, _ := strconv.ParseInt(pcid, 10, 64)
	// 取消帖子点赞
	if pid != "" {
		ok = service.CancelVote(pidInt64, pcidInt64, uid)
	} else {
		// 取消帖子点赞
		ok = service.CancelVote(0, pcidInt64, uid)
	}

	if !ok {
		c.JSON(200, gin.H{
			"error_code": 7000,
			"msg":        "操作失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "已取消点赞",
	})
}

// GetVoteListByUserPost 根据用户帖子获取点赞列表
func GetVoteListByUserPost(c *gin.Context) {
	uid := c.Keys["UID"].(int64)
	page := c.Query("page")
	size := c.Query("size")

	if page == "" {
		page = "1"
	}
	if size == "" {
		size = "10"
	}
	uidString := strconv.FormatInt(uid, 10)
	utils.RedisClient.Set(uidString+"vote", 0)

	data := service.GetVoteListByUserPost(uid, page, size)
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
		"data":       data,
	})
}

// ActionMsgadopt ActionMsgadopt
func ActionMsgadopt(c *gin.Context) {
	uid := c.Keys["UID"].(int64)
	uidString := strconv.FormatInt(uid, 10)
	voteNum, _ := redis.Int64(utils.RedisClient.Get(uidString + "adopt"))
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
		"data": map[string]int64{
			"num": voteNum,
		},
	})
}

// ActionMsglike ActionMsglike
func ActionMsglike(c *gin.Context) {
	uid := c.Keys["UID"].(int64)
	uidString := strconv.FormatInt(uid, 10)
	voteNum, _ := redis.Int64(utils.RedisClient.Get(uidString + "vote"))
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
		"data": map[string]int64{
			"num": voteNum,
		},
	})
}

// ActionMsgcomment ActionMsgcomment
func ActionMsgcomment(c *gin.Context) {
	uid := c.Keys["UID"].(int64)
	uidString := strconv.FormatInt(uid, 10)
	voteNum, _ := redis.Int64(utils.RedisClient.Get(uidString + "comment"))
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
		"data": map[string]int64{
			"num": voteNum,
		},
	})
}
