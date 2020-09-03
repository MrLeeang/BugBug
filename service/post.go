package service

import (
	"BugBug/db"
	"BugBug/models"
	"reflect"
)

// AddPost 创建帖子
func AddPost(postData map[string]interface{}) map[string]interface{} {
	postModel := db.AddPost(postData)
	// 拿类型
	postType := postData["type"]
	if postType == 4 || postType == 0 {
		// 纯链接
		return postData
	} else if postType == 2 {
		// 有视频
	}
	// 图片信息处理
	postImgList := []map[string]interface{}{}
	imgs := postData["imgs"]
	imgObj := reflect.ValueOf(imgs)
	if imgObj.Kind() == reflect.Slice {
		if imgObj.Len() != 0 {
			for i := 0; i < imgObj.Len(); i++ {
				imgURL := imgObj.Index(i).Interface()
				postImg := map[string]interface{}{}
				postImg["url"] = imgURL
				postImg["uid"] = postData["uid"]
				postImg["pid"] = postModel.Id
				db.AddPostImg(postImg)
				postImgList = append(postImgList, postImg)
			}
		}
	}

	// 重写图片信息
	postData["imgs"] = postImgList

	// 话题
	postTopicList := []map[string]interface{}{}
	topics := postData["topics"]
	topicObj := reflect.ValueOf(topics)
	if topicObj.Kind() == reflect.Slice {
		if topicObj.Len() != 0 {
			for i := 0; i < topicObj.Len(); i++ {
				cid := topicObj.Index(i).Interface()
				postTopic := map[string]interface{}{}
				postTopic["cid"] = cid
				postTopic["uid"] = postData["uid"]
				postTopic["pid"] = postModel.Id
				db.AddPostTopic(postTopic)
				postImgList = append(postTopicList, postTopic)
			}
		}
	}

	// 重写图片信息
	postData["topics"] = postTopicList
	postData["id"] = postModel.Id
	return postData
}

// GetPostByID 获取帖子
func GetPostByID(id int64) models.FbPosts {
	return db.GetPostByID(id)
}

// DetailPostList 查询postlist
func DetailPostList(queryMap map[string]interface{}, keywords string, page int, size int) []map[string]interface{} {

	// 所有帖子
	allPostInfoList := []map[string]interface{}{}
	// 获取所有采纳的帖子
	posts := db.DetailPostList(queryMap, keywords, page, size)
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
