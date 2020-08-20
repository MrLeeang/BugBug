package models

import "time"

type FbUsers struct {
	Id        int64     `json:"id"`
	Phone     string    `json:"phone"`      // 手机号
	Nickname  string    `json:"nickname"`   // 昵称
	Sex       int       `json:"sex"`        // 性别,1:男,0:女
	Avatar    string    `json:"avatar"`     // 头像
	Signature string    `json:"signature"`  // 个性签名
	Status    int       `json:"status"`     // 状态,1:可用,0:不可用
	Level     int       `json:"level"`      // 等级
	Score     int       `json:"score"`      // 积分
	LastLogin time.Time `json:"last_login"` // 最后登录时间
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
