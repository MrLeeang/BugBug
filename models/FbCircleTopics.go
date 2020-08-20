package models

import "time"

type FbCircleTopics struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`   // 用户UID
	Cid       int64     `json:"cid"`   // 圈子ID
	Title     string    `json:"title"` // 话题名称
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
