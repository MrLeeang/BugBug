package views

import (
	"BugBug/service"
	"BugBug/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ActionCreateComment 发表评论
func ActionCreateComment(c *gin.Context) {
	// 帖子得ip
	pid := c.PostForm("pid")
	// 回复的评论ID,可以为空
	pcid := c.PostForm("pcid")
	// 评论内容
	content := c.PostForm("content")
	if pid == "" || content == "" {
		c.JSON(200, gin.H{
			"error_code": 500,
			"msg":        "缺少贴子ID或评论内容",
		})
		return
	}
	// 获取帖子
	post := service.GetPostByID(pid)

	if post.Id == 0 {
		c.JSON(200, gin.H{
			"error_code": 60003,
			"msg":        "你回复的贴子已经不存在",
		})
		return
	}

	var toUID int64

	if pcid != "" {
		postComment := service.GetPostCommentByID(pcid)
		if postComment.Id == 0 {
			c.JSON(200, gin.H{
				"error_code": 60002,
				"msg":        "你回复的评论已被删除",
			})
			return
		}
		toUID = postComment.Uid
	} else {
		toUID = post.Uid
	}

	utils.UtilsLogger.Info(pid, pcid, content, toUID)

	postCommentParams := map[string]interface{}{
		"uid":     c.Keys["UID"],
		"pid":     pid,
		"pcid":    pcid,
		"toUid":   toUID,
		"content": content,
	}

	ok, postComment := service.CreateComment(postCommentParams)

	if !ok {
		c.JSON(200, gin.H{
			"error_code": 60000,
			"msg":        "评论发表失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "发表成功",
		"data":       postComment,
	})
}

// ActionDeleteComment 删除评论
func ActionDeleteComment(c *gin.Context) {
	postCommentID := c.Param("postCommentID")
	postComment := service.GetPostCommentByID(postCommentID)
	if postComment.Id == 0 {
		c.JSON(200, gin.H{
			"error_code": 60002,
			"msg":        "评论不存在",
		})
		return
	}

	uid := c.Keys["UID"].(string)
	// string 转成int64
	uidInt64, _ := strconv.ParseInt(uid, 10, 64)

	if postComment.Uid != uidInt64 {
		c.JSON(200, gin.H{
			"error_code": 3102,
			"msg":        "你没有权限删除评论",
			"data":       postCommentID,
		})
		return
	}
	ok := service.DeleteCommentByID(postCommentID)
	if !ok {
		c.JSON(200, gin.H{
			"error_code": 500,
			"msg":        "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"error_code": 0,
		"msg":        "删除成功",
	})
}

// ActionCommentReplyList 评论的回复列表
func ActionCommentReplyList(c *gin.Context) {
	postCommentID := c.Param("postCommentID")

	queryMap := map[string]interface{}{
		"pcid": postCommentID,
	}

	PostCommentList := service.DetailPostComments(queryMap)

	c.JSON(200, gin.H{
		"data":       PostCommentList,
		"error_code": 0,
		"msg":        "success.",
	})
}

// ActionCommentListByUser 我得评论列表
func ActionCommentListByUser(c *gin.Context) {
	queryMap := map[string]interface{}{
		"to_uid": c.Keys["UID"].(string),
	}

	PostCommentList := service.DetailPostComments(queryMap)

	c.JSON(200, gin.H{
		"data":       PostCommentList,
		"error_code": 0,
		"msg":        "success.",
	})
}
