package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 注册logger 跨域
	r.Use(Cors())

	// 获取时间
	r.GET("/time", func(c *gin.Context) {

		// date "+%Y-%m-%d %H:%M:%S"
		cmd := exec.Command("date", "+%Y-%m-%d %H:%M:%S")

		timeByte, err := cmd.Output()

		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"message": err.Error(),
				"code":    -1,
			})
			return
		}

		timeStr := string(timeByte)
		timeStr = strings.Replace(timeStr, "\n", "", 1)

		c.JSON(200, gin.H{
			"message": timeStr,
			"code":    0,
		})
	})

	// 修改时间
	r.PUT("/time", func(c *gin.Context) {
		time := c.PostForm("time")

		fmt.Printf("set time %s \n", time)

		cmd := exec.Command("date", "-s", time)

		ret, err := cmd.Output()

		if err != nil {
			fmt.Println(ret)
			fmt.Println(err)
			c.JSON(200, gin.H{
				"message": "Failed",
				"code":    -1,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Success",
			"code":    0,
		})
	})
	// 启动
	fmt.Println("Server Run Success: 0.0.0.0:9501")
	r.Run(":9501")
}
