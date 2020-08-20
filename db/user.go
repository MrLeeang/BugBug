package db

import (
	models "BugBug/models"
	"BugBug/utils"
	"fmt"
)

// orm 文档
//https://www.kancloud.cn/xormplus/xorm/167111

// 连表
//type UserGroup struct {
//	models.FbUsers `xorm:"extends"`
//	Name string
//}
//
//func (UserGroup) TableName() string {
//	return "user"
//}
//
//var allUsers = make([]UserGroup, 0)
//_ = Engine.SQL("select user.*, group.name from user, group where user.group_id = group.id").Find(&allUsers)

//engine.Table("user").Join("INNER", "group", "group.id = user.group_id").
//Join("INNER", "type", "type.id = user.type_id").
//Where("user.name like ?", "%"+name+"%").Find(&users, &User{Name:name})

// QueryUsers 查询所有用户
func QueryUsers() []models.FbUsers {
	// 定义一个数组存放结构体
	// allUsers := []*models.UserModel{}
	var allUsers []models.FbUsers
	// 查询
	var err = Engine.Find(&allUsers)

	if err != nil {
		utils.UtilsLogger.Error(err)
	}

	return allUsers
}
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
