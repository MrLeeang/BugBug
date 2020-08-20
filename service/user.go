package service

import (
	"BugBug/db"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
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

func GetUserByPhone(phone string) map[string]interface{} {

	var ret = map[string]interface{}{}

	sqlStr := fmt.Sprintf("select * from fb_users where phone='%s' limit 1;", phone)
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

func VerifyLoginCode(phone string, code string) bool {
	// 短信服务请求地址
	verifyUri := "http://webapi.sms.mob.com/sms/verify"
	// app key
	appKey := "2eca699e046f4"
	// 申明并创建一个cookie
	var gCurCookieJar *cookiejar.Jar
	gCurCookieJar = new(cookiejar.Jar)
	// 创建一个http client
	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}
	// 构造请求参数
	var r http.Request
	_ = r.ParseForm()
	r.Form.Add("appkey", appKey)
	r.Form.Add("phone", phone)
	r.Form.Add("code", code)
	r.Form.Add("zone", "86")
	bodyStr := strings.TrimSpace(r.Form.Encode())
	// 构建普通form请求， json 用 bytes.NewBuffer application/json

	//httpReq, _ := http.NewRequest("POST", verifyUri, strings.NewReader(bodyStr))
	//httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 发送请求
	//httpResp, err := httpClient.Do(httpReq)
	// 请求
	httpResp, err := httpClient.Post(
		verifyUri,
		"application/x-www-form-urlencoded",
		strings.NewReader(bodyStr),
	)
	if err != nil {
		utils.UtilsLogger.Error(err.Error())
		return false
	}
	// 关闭请求
	defer httpResp.Body.Close()
	// 获取返回值
	body, _ := ioutil.ReadAll(httpResp.Body)
	// json转成map
	var ret map[string]interface{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		utils.UtilsLogger.Error(err.Error())
		return false
	}
	// 判断返回值
	if ret["status"] != 200 {
		utils.UtilsLogger.Error(ret)
		return false
	}
	return true
}

// SecretKey 认证key
const (
	SecretKey = "fdsafjkldsaklfjkdlasjfkljsdaklfjskdlafklsdjaklfjaslfjiouwiotfdsafasdfjikljsngklnsadvhasjkfghejskd"
	Issuer    = "bugbug"
)

// jwtCustomClaims token签名信息
type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid interface{}
}

// GenerateToken 生成token
func GenerateToken(Uid interface{}) string {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    Issuer,
		},
		Uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		utils.UtilsLogger.Error(err)
		return ""
	}
	return tokenString
}

// VerifyToken 验证token
func VerifyToken(tokenSrt string) jwt.Claims {
	//var verifyToken *jwt.Token
	verifyToken, err := jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil
	}
	claims := verifyToken.Claims
	return claims
}
