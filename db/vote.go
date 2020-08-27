package db

import (
	"BugBug/models"
	"BugBug/utils"
	"strconv"
	"time"
)

// AddVote 点赞
func AddVote(uid string, pid string, pcid string, result string) bool {
	// string 转成int64
	uidInt64, _ := strconv.ParseInt(uid, 10, 64)
	pidInt64, _ := strconv.ParseInt(pid, 10, 64)
	pcidInt64, _ := strconv.ParseInt(pcid, 10, 64)
	resultInt64, _ := strconv.Atoi(result)
	vote := &models.FbVotes{}
	vote.Uid = uidInt64
	vote.Pid = pidInt64
	vote.Pcid = pcidInt64
	vote.Result = resultInt64
	vote.CreatedAt = time.Now()
	vote.UpdatedAt = time.Now()
	_, err := Engine.InsertOne(vote)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return false
	}
	return true
}

// CancelVote 取消点赞
func CancelVote(pid string, pcid string, uid string) bool {
	sql := "update fb_votes set deleted_at = ? where pid = ? and pcid = ? and uid = ?"
	_, err := Engine.Exec(sql, time.Now(), pid, pcid, uid)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return false
	}
	return true
}

// GetVoteListByUserPost 用户帖子点赞列表
func GetVoteListByUserPost(userID string, page string, size string) []map[string]interface{} {
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	sqlStr := "SELECT vote.id as vid, u.nickname, post.content, post.id as pid from fb_votes vote, fb_posts post, fb_users u  WHERE vote.pid=post.id and vote.uid=u.id and post.id in (SELECT id FROM fb_posts WHERE uid = ? )"
	results, err := Engine.SQL(sqlStr, userID).Query().ListPage(pageInt, sizeInt)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return []map[string]interface{}{}
	}
	return results
}
