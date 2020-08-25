package db

import (
	models "BugBug/models"
	"BugBug/utils"
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
