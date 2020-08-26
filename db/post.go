package db

import (
	models "BugBug/models"
	"BugBug/utils"
	"fmt"
	"strconv"
	"time"
)

// GetPostByID 根据id获取帖子
func GetPostByID(id string) models.FbPosts {
	var Post models.FbPosts
	// 查询
	_, err := Engine.Where("id = ?", id).Get(&Post)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return Post
}

// GetPostByIDs in
func GetPostByIDs(ids []int64) []models.FbPosts {
	var posts []models.FbPosts
	// 查询
	err := Engine.In("id", ids).Find(&posts)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return posts
}

// GetPostImagesByPostID 获取帖子图片
func GetPostImagesByPostID(postID int64) []models.FbPostImgs {
	var postIms []models.FbPostImgs
	err := Engine.Where("pid = ?", postID).Find(&postIms)
	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return postIms

}

// GetPostVideoByPostID 获取帖子视频
func GetPostVideoByPostID(postID int64) models.FbPostVideos {
	var postVideo models.FbPostVideos
	_, err := Engine.Where("pid = ?", postID).Get(&postVideo)
	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return postVideo

}

// DetailPostList 查询
func DetailPostList(key string, value interface{}) []models.FbPosts {
	// 定义一个数组存放结构体
	var postList []models.FbPosts
	// 查询
	sqlString := fmt.Sprintf("%s=?", key)
	var err = Engine.Where(sqlString, value).Find(&postList)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return postList
}

// AddPost 发帖
func AddPost(postData map[string]interface{}) models.FbPosts {
	var post = &models.FbPosts{}

	uid := postData["uid"].(string)
	cid := postData["cid"].(string)

	uidInt64, _ := strconv.ParseInt(uid, 10, 64)
	cidInt64, _ := strconv.ParseInt(cid, 10, 64)
	post.Uid = uidInt64
	post.Cid = cidInt64
	if postData["title"] == nil {
		post.Title = ""
	} else {
		post.Title = postData["title"].(string)
	}
	post.Link = postData["link"].(string)
	post.Content = postData["content"].(string)
	post.Summary = postData["content"].(string)
	post.Type = postData["type"].(int)
	post.Views = 0
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	_, err := Engine.InsertOne(post)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return *post
	}
	return *post
}

// AddPostImg 图片
func AddPostImg(imgData map[string]interface{}) models.FbPostImgs {

	var postImg = &models.FbPostImgs{}

	uid := imgData["uid"].(string)
	pidInt64 := imgData["pid"].(int64)
	url := imgData["url"].(string)

	uidInt64, _ := strconv.ParseInt(uid, 10, 64)

	postImg.Uid = uidInt64
	postImg.Pid = pidInt64
	postImg.Url = url
	postImg.Type = 1
	postImg.Views = 0
	postImg.CreatedAt = time.Now()
	postImg.UpdatedAt = time.Now()
	_, err := Engine.InsertOne(postImg)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return *postImg
	}
	return *postImg
}

// AddPostTopic 话题
func AddPostTopic(topicData map[string]interface{}) models.FbPostTopics {

	var topicImg = &models.FbPostTopics{}

	uid := topicData["uid"].(string)
	pidInt64 := topicData["pid"].(int64)
	ctid := topicData["cid"].(string)

	uidInt64, _ := strconv.ParseInt(uid, 10, 64)
	ctidInt64, _ := strconv.ParseInt(ctid, 10, 64)

	topicImg.Uid = uidInt64
	topicImg.Pid = pidInt64
	topicImg.Ctid = ctidInt64
	topicImg.CreatedAt = time.Now()
	topicImg.UpdatedAt = time.Now()
	_, err := Engine.InsertOne(topicImg)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return *topicImg
	}
	return *topicImg
}
