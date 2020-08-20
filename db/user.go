package db

import (
	models "BugBug/models"
	"BugBug/utils"
	"fmt"
)

func DetailUsers(key string, value string) []models.FbUsers {
	// 定义一个数组存放结构体
	var UserList []models.FbUsers
	// 查询
	sqlString := fmt.Sprintf("%s=?", key)
	var err = Engine.Where(sqlString, value).Find(&UserList)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return UserList
}

func GetUserById(id string) models.FbUsers {
	// 定义一个结构体
	var User models.FbUsers
	// 查询
	_, err := Engine.Where("id = ?", id).Get(&User)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return User
}
