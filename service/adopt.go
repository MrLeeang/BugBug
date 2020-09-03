package service

import (
	"BugBug/db"
	"BugBug/utils"
	"fmt"
	"reflect"
	"strconv"
)

// CountAdoptByUserID 采纳数
func CountAdoptByUserID(userID string) int {

	queryStr := fmt.Sprintf("select count(id) from fb_adopts where uid='%s';", userID)
	queryResult, err := db.Engine.QueryString(queryStr)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return 0
	}
	if len(queryResult) < 1 {
		return 0
	}
	queryInfo := queryResult[0]
	adoptCount, err := strconv.Atoi(queryInfo["count(id)"])
	if err != nil {
		utils.UtilsLogger.Error(err)
		return 0
	}
	return adoptCount
}

// AdoptPost 采纳
func AdoptPost(uid int64, pid int64) bool {
	return db.AdoptPost(uid, pid)
}

// UserAdpotPostList 用户采纳帖子列表
func UserAdpotPostList(uid string) []map[string]interface{} {
	// 采纳帖子pid列表
	adoptPidList := []int64{}
	// 获取用户采纳的所有帖子
	userAdoptList := db.GetUserAdoptList(uid)
	// 组装pid
	for _, userAdopt := range userAdoptList {
		adoptPidList = append(adoptPidList, userAdopt.Id)
	}
	// 所有帖子
	allPostInfoList := []map[string]interface{}{}
	// 获取所有采纳的帖子
	posts := db.GetPostByIDs(adoptPidList)
	// 获取用户列表
	userList := db.QueryUsers()
	allUserMapData := map[int64]interface{}{}
	for _, user := range userList {
		allUserMapData[user.Id] = user
	}
	// 组装帖子详细信息
	for _, post := range posts {
		obj1 := reflect.TypeOf(post)
		obj2 := reflect.ValueOf(post)
		// strut 转 map
		var postData = map[string]interface{}{}
		for i := 0; i < obj1.NumField(); i++ {
			postData[obj1.Field(i).Tag.Get("json")] = obj2.Field(i).Interface()
		}
		// 空数据
		emptyList := []string{}
		// 查询帖子图片
		postIms := db.GetPostImagesByPostID(post.Id)
		postData["user"] = allUserMapData[post.Uid]
		if postIms == nil {
			postData["imgs"] = emptyList
		} else {
			postData["imgs"] = postIms
		}
		postData["topics"] = emptyList
		postData["circle"] = emptyList
		// 视频信息
		postData["video"] = db.GetPostVideoByPostID(post.Id)

		allPostInfoList = append(allPostInfoList, postData)
	}
	return allPostInfoList
}

// GetAdoptListByUserPost 根据用户帖子获取采纳列表
func GetAdoptListByUserPost(userID int64, page string, size string) []map[string]interface{} {
	return db.GetAdoptListByUserPost(userID, page, size)
}

// GetAdoptPostListByUser 根据用户帖子获取采纳列表
func GetAdoptPostListByUser(userID string, page string, size string) []map[string]interface{} {
	return db.GetAdoptPostListByUser(userID, page, size)
}
