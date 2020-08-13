package views

import (
	"github.com/gin-gonic/gin"
)

func ActionRegister(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}
