package models

import "time"

type FbPostImgs struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`   // 用户UID
	Pid       int64     `json:"pid"`   // 文章ID
	Url       string    `json:"url"`   // 资源url
	Type      int       `json:"type"`  // 类型,1:静态,2:动态
	Views     int64     `json:"views"` // 查看数
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
