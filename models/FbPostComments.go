package models

import "time"

type FbPostComments struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`     // 用户UID
	Pid       int64     `json:"pid"`     // 文章ID
	Cid       int64     `json:"cid"`     // 圈子ID
	Pcid      int64     `json:"pcid"`    // 评论ID
	ToUid     int64     `json:"to_uid"`  // 被评人UID
	Content   string    `json:"content"` // 评论内容
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
