package service

import (
	"BugBug/db"
	"fmt"
)
import "BugBug/utils"

func GetUserById(userId string) map[string]interface{} {

	var ret = map[string]interface{}{}

	sqlStr := fmt.Sprintf("select * from fb_users where id='%s' limit 1;", userId)
	queryResult, err := db.Engine.QueryString(sqlStr)
	if err != nil {
		utils.UtilsLogger.Error(err)
		return ret
	}
	if len(queryResult) < 1 {
		return ret
	}
	userInfo := queryResult[0]
	ret["id"] = userInfo["id"]
	ret["phone"] = userInfo["phone"]
	ret["nickname"] = userInfo["nickname"]
	ret["avatar"] = userInfo["avatar"]
	ret["signature"] = userInfo["signature"]
	ret["status"] = userInfo["status"]
	ret["level"] = userInfo["level"]
	ret["score"] = userInfo["score"]
	ret["last_login"] = userInfo["last_login"]
	ret["created_at"] = userInfo["created_at"]
	ret["updated_at"] = userInfo["updated_at"]
	ret["deleted_at"] = userInfo["deleted_at"]
	return ret
}
