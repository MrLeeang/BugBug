package models

import "time"

// 消息通知表
type FbInfo struct {
	Id        int       `json:"id"`
	Uid       int       `json:"uid"`
	PostId    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	IsRead    int       `json:"is_read"`
	UpdatedAt time.Time `json:"updated_at"`
	Type      int       `json:"type"`
	DeletedAt time.Time `json:"deleted_at"`
}
