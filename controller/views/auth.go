package views

import (
	"BugBug/service"
	"BugBug/utils"

	"github.com/gin-gonic/gin"
)

// AuthHandler 登录验证
func AuthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("token")
		verify := service.VerifyToken(tokenStr)
		if verify == nil {
			utils.UtilsLogger.Error(verify)
			var ret = map[string]interface{}{}
			retCode := utils.RetCode.LoginError
			ret["code"] = retCode
			ret["ret_message"] = utils.ErrorCodeMessage[retCode]
			context.JSON(200, gin.H{
				"ret": ret,
			})
			context.Abort()
			return
		}
		// gin上下文存储context.Keys，verifyMap报存到上下文中
		verifyMap := map[string]interface{}{}
		for key, val := range verify {
			if key == "UID" {
				val = int64(val.(float64))
			}
			verifyMap[key] = val
		}
		context.Keys = verifyMap
		context.Next()
	}
}
