package service

import (
	"BugBug/db"
	"BugBug/models"
)

// AddPost 创建帖子
func AddPost(postData map[string]interface{}) map[string]interface{} {

	return map[string]interface{}{}
}

// GetPostByID 获取帖子
func GetPostByID(id string) models.FbPosts {
	return db.GetPostByID(id)
}
