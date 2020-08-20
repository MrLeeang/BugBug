package models

import "time"

type FbCircleMembers struct {
	Id        int64     `json:"id"`
	Cid       int64     `json:"cid"` // 圈子ID
	Uid       int64     `json:"uid"` // 用户ID
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
