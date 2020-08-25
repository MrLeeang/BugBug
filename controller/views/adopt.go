package views

import (
	"BugBug/service"

	"github.com/gin-gonic/gin"
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
	uid := c.Keys["UID"].(string)

	ok := service.AdoptPost(uid, pid)
	if !ok {
		c.JSON(200, gin.H{
			"error_code": 500,
			"msg":        "操作失败",
		})
		return
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
