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
		utils.UtilsLogger.Error(err)
		return 0
	}
	if len(queryResult) < 1 {
		return 0
	}
	queryInfo := queryResult[0]
	voteCount, err := strconv.Atoi(queryInfo["count(id)"])
	if err != nil {
		utils.UtilsLogger.Error(err)
		return 0
	}
	return voteCount
}

// AddVote 点赞
func AddVote(uid string, pid string, pcid string, result string) bool {
	return db.AddVote(uid, pid, pcid, result)
}

// CancelVote 取消点赞
func CancelVote(pid string, pcid string, uid string) bool {
	return db.CancelVote(pid, pcid, uid)
}
