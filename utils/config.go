package utils

import "fmt"

const (
	// DbType 数据类型
	DbType = "mysql"
	// DbHost 数据库地址
	DbHost = "api.xianglian.group"
	// DbPort DbPort
	DbPort = "3306"
	// DbUser DbUser
	DbUser = "root"
	// DbPass DbPass
	DbPass = "bugbug@2020"
	// DbName DbName
	DbName = "feedback"
	// RedisHost RedisHost
	RedisHost = "118.31.237.114"
	// RedisPort RedisPort
	RedisPort = "6379"
	// RedisPassword RedisPassword
	RedisPassword = "ccadmin1QAZ"

	// QiniuKey QiniuKey
	QiniuKey = "9lgBfFDaat7MKyZG4oh895FAIGsuKdX9UKqiEpaq"
	// QiniuSecret QiniuSecret
	QiniuSecret = "DXK1CUogN_Tyqz5c_buwhpKOkQSkpl5uVJLg7CYf"
	// QiniuBucket QiniuBucket
	QiniuBucket = "bugbug-img"
	// QiniuHost QiniuHost
	QiniuHost = "http://img.xianglian.group"
)

// RedisConnStr RedisConnStr
var RedisConnStr = fmt.Sprintf("%s:%s", RedisHost, RedisPort)
