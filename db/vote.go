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
