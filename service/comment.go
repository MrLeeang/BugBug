package service

import (
	"BugBug/db"
	"BugBug/models"
)

// CreateComment 发表评论
func CreateComment(params map[string]interface{}) (bool, models.FbPostComments) {
	return db.CreateComment(params)
}

// GetPostCommentByID 通过id获取评论
func GetPostCommentByID(id string) models.FbPostComments {
	return db.GetPostCommentByID(id)
}

// DeleteCommentByID 通过id删除评论
func DeleteCommentByID(id string) bool {
	return db.DeleteCommentByID(id)
}

// DetailPostComments 查询评论
func DetailPostComments(params map[string]interface{}) []models.FbPostComments {
	return db.DetailPostComments(params)
}
