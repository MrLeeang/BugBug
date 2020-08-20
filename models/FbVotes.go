package models

import "time"

type FbVotes struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`    // 用户UID
	Pid       int64     `json:"pid"`    // 文章ID
	Pcid      int64     `json:"pcid"`   // 评论ID
	Result    int       `json:"result"` // 结果,0:反对,1:赞成
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
