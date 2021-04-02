package db

import (
	models "BugBug/models"
	"BugBug/utils"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// GetPostByID 根据id获取帖子
func GetPostByID(id int64) models.FbPosts {
	var Post models.FbPosts
	// 查询
	_, err := Engine.Where("id = ?", id).Get(&Post)

	if err != nil {
		utils.Logger.Error(err)
	}

	return Post
}

// GetPostByIDs in
func GetPostByIDs(ids []int64) []models.FbPosts {
	var posts []models.FbPosts
	// 查询
	err := Engine.In("id", ids).Find(&posts)

	if err != nil {
		utils.Logger.Error(err)
	}

	return posts
}

// GetPostImagesByPostID 获取帖子图片
func GetPostImagesByPostID(postID int64) []models.FbPostImgs {
	var postIms []models.FbPostImgs
	err := Engine.Where("pid = ?", postID).Find(&postIms)
	if err != nil {
		utils.Logger.Error(err)
	}

	return postIms

}

// GetPostVideoByPostID 获取帖子视频
func GetPostVideoByPostID(postID int64) models.FbPostVideos {
	var postVideo models.FbPostVideos
	_, err := Engine.Where("pid = ?", postID).Get(&postVideo)
	if err != nil {
		utils.Logger.Error(err)
	}

	return postVideo

}

// DetailPostList 查询
func DetailPostList(params map[string]interface{}, keywords string, page int, size int) []models.FbPosts {
	// 定义一个数组存放结构体
	var postList []models.FbPosts

	// 反射
	m := reflect.ValueOf(params)
	// 不是map
	if m.Kind() != reflect.Map {
		utils.Logger.Error("params error")
		return postList
	}

	query := Engine.Where("id!=?", 0)
	// 通过反射拿到所有的key
	keys := m.MapKeys()
	for _, key := range keys {
		value := m.MapIndex(key)
		sqlString := fmt.Sprintf("%s=?", key.Interface())
		query.And(sqlString, value.Interface())
	}
	if keywords != "" {
		query.And("content like ? or title like ?", "%"+keywords+"%", "%"+keywords+"%")
	}
	var err = query.Limit(size*page, (page-1)*size).Find(&postList)
	if err != nil {
		utils.Logger.Error(err)
	}

	return postList
}

// AddPost 发帖
func AddPost(postData map[string]interface{}) models.FbPosts {
	var post = &models.FbPosts{}

	uidInt64 := postData["uid"].(int64)
	cid := postData["cid"].(string)

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
		utils.Logger.Error(err)
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
		utils.Logger.Error(err)
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
		utils.Logger.Error(err)
		return *topicImg
	}
	return *topicImg
}
