package service

import (
	"BugBug/db"
	"BugBug/utils"
	"fmt"
	"strconv"
)

// CountVoteByUserID 统计点赞数
func CountVoteByUserID(userID string) int {

	queryStr := fmt.Sprintf("select count(id) from fb_votes where uid='%s';", userID)
	queryResult, err := db.Engine.QueryString(queryStr)
	if err != nil {
		utils.Logger.Error(err)
		return 0
	}
	if len(queryResult) < 1 {
		return 0
	}
	queryInfo := queryResult[0]
	voteCount, err := strconv.Atoi(queryInfo["count(id)"])
	if err != nil {
		utils.Logger.Error(err)
		return 0
	}
	return voteCount
}

// AddVote 点赞
func AddVote(uid int64, pid int64, pcid int64, result string) bool {
	return db.AddVote(uid, pid, pcid, result)
}

// CancelVote 取消点赞
func CancelVote(pid int64, pcid int64, uid int64) bool {
	return db.CancelVote(pid, pcid, uid)
}

// GetVoteListByUserPost 根据用户帖子获取点赞列表
func GetVoteListByUserPost(userID int64, page string, size string) []map[string]interface{} {
	return db.GetVoteListByUserPost(userID, page, size)
}
