package db

import (
	"BugBug/models"
	"BugBug/utils"
	"strconv"
	"time"
)

// AdoptPost 采纳
func AdoptPost(uid int64, pid int64) bool {
	adopt := &models.FbAdopts{}
	adopt.Uid = uid
	adopt.Pid = pid
	adopt.CreatedAt = time.Now()
	adopt.UpdatedAt = time.Now()
	_, err := Engine.InsertOne(adopt)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return false
	}

	return true
}

// GetUserAdoptList 采纳列表
func GetUserAdoptList(uid string) []models.FbAdopts {
	var doptList []models.FbAdopts
	var err = Engine.Where("uid=?", uid).Find(&doptList)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return doptList
}

// GetAdoptListByUserPost 用户帖子点赞列表
func GetAdoptListByUserPost(userID int64, page string, size string) []map[string]interface{} {
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	sqlStr := `SELECT
	adopt.id AS aid,
	p.content,
	u.nickname,
	adopt.pid 
FROM
	fb_adopts adopt,
	fb_posts p,
	fb_users u 
WHERE
	adopt.pid = p.id 
	AND adopt.uid = u.id 
	AND p.id IN ( SELECT id FROM fb_posts WHERE uid = ? )`
	results, err := Engine.SQL(sqlStr, userID).Query().ListPage(pageInt, sizeInt)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return []map[string]interface{}{}
	}
	return results
}

// GetAdoptPostListByUser 用户帖子点赞列表
func GetAdoptPostListByUser(userID string, page string, size string) []map[string]interface{} {
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	sqlStr := `SELECT
	adopt.id AS aid,
	p.content,
	u.nickname,
	adopt.pid 
FROM
	fb_adopts adopt,
	fb_posts p,
	fb_users u 
WHERE
	adopt.pid = p.id 
	AND p.uid = u.id 
	AND adopt.uid =?`
	results, err := Engine.SQL(sqlStr, userID).Query().ListPage(pageInt, sizeInt)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return []map[string]interface{}{}
	}
	return results
}
