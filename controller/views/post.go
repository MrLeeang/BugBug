package views

import (
	"github.com/gin-gonic/gin"
)

// ActionPostPublish 发布帖子
func ActionPostPublish(c *gin.Context) {

	var ret = map[string]interface{}{}

	cid := c.PostForm("cid")
	content := c.PostForm("content")
	imgs := c.PostForm("imgs")
	video := c.PostForm("video")
	topics := c.PostForm("topics")
	coordinate := c.PostForm("coordinate")
	link := c.PostForm("link")

	ret["cid"] = cid
	ret["content"] = content
	ret["imgs"] = imgs
	ret["video"] = video
	ret["topics"] = topics
	ret["coordinate"] = coordinate
	ret["link"] = link

	c.JSON(200, gin.H{
		"data":       ret,
		"error_code": 0,
		"msg":        "success.",
	})
}
