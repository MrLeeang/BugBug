package service

import (
	"BugBug/db"
	"BugBug/utils"
	"fmt"
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
	utils.UtilsLogger.Info(queryInfo["count(id)"])
	adoptCount, err := strconv.Atoi(queryInfo["count(id)"])
	if err != nil {
		utils.UtilsLogger.Error(err)
		return 0
	}
	return adoptCount
}
