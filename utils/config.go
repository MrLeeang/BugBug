package utils

import "fmt"

const (
	DbType = "mysql"
	DbHost = "49.233.185.3"
	DbPort = "3306"
	DbUser = "root"
	DbPass = "123456"
	DbName = "feedback"

	RedisHost     = "118.31.237.114"
	RedisPort     = "6379"
	RedisPassword = "ccadmin1QAZ"
)

var SessionRedisAddress = fmt.Sprintf("%s:%s", RedisHost, RedisPort)
