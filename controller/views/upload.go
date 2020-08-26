package views

import (
	"BugBug/utils"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// ActionUploadToken shangchuan
func ActionUploadToken(c *gin.Context) {
	expires := 3600
	// token     := (new Auth($accessKey, $secretKey))->uploadToken($bucket, null, $expires);
	putPolicy := storage.PutPolicy{
		Scope: utils.QiniuBucket,
	}

	mac := qbox.NewMac(utils.QiniuKey, utils.QiniuSecret)
	upToken := putPolicy.UploadToken(mac)

	c.JSON(200, gin.H{
		"host":    utils.QiniuHost,
		"token":   upToken,
		"expires": expires,
		"region":  "1",
	})
}
