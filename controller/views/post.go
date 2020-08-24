package views

import (
	"BugBug/service"

	"github.com/gin-gonic/gin"
)

// ActionPostPublish 发布帖子
func ActionPostPublish(c *gin.Context) {

	// 获取form数据
	cid := c.PostForm("cid")
	content := c.PostForm("content")
	video := c.PostForm("video")
	coordinate := c.PostForm("coordinate")
	link := c.PostForm("link")
	// 接收多个相同的请求参数,拿到一个数组
	imgs, _ := c.GetPostFormArray("imgs[]")
	topics, _ := c.GetPostFormArray("topics[]")

	var postType = 0
	// 纯文本
	if len(imgs) == 0 && video == "" && link == "" {
		postType = 0
	}
	// 照片或者视频
	if video != "" {
		postType = 2
	} else if len(imgs) != 0 {
		postType = 1
	}
	// 纯链接
	if link != "" && len(imgs) == 0 && video == "" && content == "" {
		postType = 4
	}
	var postData = map[string]interface{}{}
	postData["cid"] = cid
	postData["content"] = content
	postData["imgs"] = imgs
	postData["video"] = video
	postData["topics"] = topics
	postData["coordinate"] = coordinate
	postData["link"] = link
	postData["type"] = postType
	// 写入数据库
	postInfo := service.AddPost(postData)

	c.JSON(200, gin.H{
		"data":       postInfo,
		"error_code": 0,
		"msg":        "success.",
	})
}

// ActionPostCommentList 帖子评论列表
func ActionPostCommentList(c *gin.Context) {
	postID := c.Param("postID")

	queryMap := map[string]interface{}{
		"pid":  postID,
		"pcid": 0,
	}

	PostCommentList := service.DetailPostComments(queryMap)

	c.JSON(200, gin.H{
		"data":       PostCommentList,
		"error_code": 0,
		"msg":        "success.",
	})
}
