package service

import (
	"BugBug/db"
	"BugBug/utils"
	"fmt"
	"strconv"
)

func CountVoteByUserId(userId string) int {

	queryStr := fmt.Sprintf("select count(id) from fb_votes where uid='%s';", userId)
	queryResult, err := db.Engine.QueryString(queryStr)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return 0
	}
	if len(queryResult) < 1 {
		return 0
	}
	queryInfo := queryResult[0]
	utils.UtilsLogger.Info(queryInfo["count(id)"])
	voteCount, err := strconv.Atoi(queryInfo["count(id)"])
	if err != nil {
		utils.UtilsLogger.Error(err)
		return 0
	}
	return voteCount
}
