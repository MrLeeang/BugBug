package models

import "time"

type FbUserFollows struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`    // 关注者
	Fid       int64     `json:"fid"`    // 被关注者
	Remark    string    `json:"remark"` // 备注
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
