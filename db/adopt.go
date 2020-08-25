package db

import (
	"BugBug/models"
	"BugBug/utils"
	"strconv"
	"time"
)

// AdoptPost 采纳
func AdoptPost(uid string, pid string) bool {
	uidInt64, _ := strconv.ParseInt(uid, 10, 64)
	pidInt64, _ := strconv.ParseInt(pid, 10, 64)
	adopt := &models.FbAdopts{}
	adopt.Uid = uidInt64
	adopt.Pid = pidInt64
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
