package views

import (
	"BugBug/service"

	"github.com/gin-gonic/gin"
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

	uid := c.Keys["UID"].(string)
	ok := service.AddVote(uid, pid, pcid, result)
	if !ok {
		c.JSON(200, gin.H{
			"error_code": 7000,
			"msg":        "点赞失败",
		})
		return
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
	uid := c.Keys["UID"].(string)
	// 取消帖子点赞
	if pid != "" {
		ok = service.CancelVote(pid, pcid, uid)
	} else {
		// 取消帖子点赞
		ok = service.CancelVote("0", pcid, uid)
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
	uid := c.Keys["UID"].(string)
	page := c.Query("page")
	size := c.Query("size")

	if page == "" {
		page = "1"
	}
	if size == "" {
		size = "10"
	}

	data := service.GetVoteListByUserPost(uid, page, size)
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "success",
		"data":       data,
	})
}
