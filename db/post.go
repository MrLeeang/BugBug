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
