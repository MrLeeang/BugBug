package models

import "time"

type FbAdopts struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"` // 用户UID
	Pid       int64     `json:"pid"` // 文章ID
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
