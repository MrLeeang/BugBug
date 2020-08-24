package db

import (
	models "BugBug/models"
	"BugBug/utils"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// CreateComment 发表评论
func CreateComment(params map[string]interface{}) (bool, models.FbPostComments) {
	var postComment = &models.FbPostComments{}
	uid := params["uid"].(string)
	pid := params["pid"].(string)
	pcid := params["pcid"].(string)

	// string 转成int64
	uidInt64, _ := strconv.ParseInt(uid, 10, 64)
	pidInt64, _ := strconv.ParseInt(pid, 10, 64)
	pcidInt64, _ := strconv.ParseInt(pcid, 10, 64)

	postComment.Uid = uidInt64
	postComment.Pid = pidInt64
	postComment.Pcid = pcidInt64
	postComment.ToUid = params["toUid"].(int64)
	postComment.Content = params["content"].(string)
	postComment.CreatedAt = time.Now()
	postComment.UpdatedAt = time.Now()

	_, err := Engine.InsertOne(postComment)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return false, *postComment
	}

	return true, *postComment
}

// GetPostCommentByID 通过id获取评论
func GetPostCommentByID(id string) models.FbPostComments {
	var PostComment models.FbPostComments
	// 查询
	_, err := Engine.Where("id = ?", id).Get(&PostComment)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return PostComment
}

// DeleteCommentByID 通过id删除评论
func DeleteCommentByID(id string) bool {
	var PostComment models.FbPostComments
	_, err := Engine.ID(id).Delete(&PostComment)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return false
	}
	return true
}

// DetailPostComments 查询
func DetailPostComments(params map[string]interface{}) []models.FbPostComments {
	// 定义一个数组存放结构体
	var PostCommentList []models.FbPostComments
	// 反射
	m := reflect.ValueOf(params)
	// 不是map
	if m.Kind() != reflect.Map {
		utils.UtilsLogger.Error("params error")
		return PostCommentList
	}

	query := Engine.Where("id!=?", 0)
	// 通过反射拿到所有的key
	keys := m.MapKeys()
	for _, key := range keys {
		value := m.MapIndex(key)
		sqlString := fmt.Sprintf("%s=?", key.Interface())
		query.And(sqlString, value.Interface())
	}
	var err = query.Find(&PostCommentList)
	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return PostCommentList
}