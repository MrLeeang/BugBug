package views

import (
	"BugBug/service"
	"strconv"

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
	uid := c.Keys["UID"].(string)
	postData["uid"] = uid
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
		"msg":        "发布成功.",
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

// ActionUserPostList 用户帖子列表
func ActionUserPostList(c *gin.Context) {
	userID := c.Param("userId")
	queryMap := map[string]interface{}{}
	queryMap["uid"] = userID
	postList := service.DetailPostList(queryMap, "", 1, 50)
	c.JSON(200, gin.H{
		"data":       postList,
		"error_code": 0,
		"msg":        "success.",
	})
}

// ActionRecommendPostList 推荐的帖子
func ActionRecommendPostList(c *gin.Context) {
	// 找到关注的人发表的
	queryMap := map[string]interface{}{}
	// 所有加入的圈子
	page := c.Query("page")
	size := c.Query("size")
	keywords := c.Query("keywords")
	if page == "" {
		page = "1"
	}
	if size == "" {
		size = "50"
	}

	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)

	postList := service.DetailPostList(queryMap, keywords, pageInt, sizeInt)
	c.JSON(200, gin.H{
		"data":       postList,
		"error_code": 0,
		"msg":        "success.",
	})
}

// ActionPostInfo 帖子详情
func ActionPostInfo(c *gin.Context) {
	postID := c.Param("postID")
	queryMap := map[string]interface{}{}
	queryMap["id"] = postID
	postList := service.DetailPostList(queryMap, "", 1, 50)
	if len(postList) == 0 {
		c.JSON(200, gin.H{
			"data":       nil,
			"error_code": 0,
			"msg":        "success.",
		})
		return
	}
	c.JSON(200, gin.H{
		"data":       postList[0],
		"error_code": 0,
		"msg":        "success.",
	})

}
