package models

import "time"

type FbPosts struct {
	Id         int64     `json:"id"`
	Uid        int64     `json:"uid"`        // 用户UID
	Cid        int64     `json:"cid"`        // 圈子ID
	Title      string    `json:"title"`      // 标题
	Link       string    `json:"link"`       // 外部链接
	Gid        int64     `json:"gid"`        // 商品ID
	Content    string    `json:"content"`    // 内容
	Summary    string    `json:"summary"`    // 内容摘要
	Coordinate string    `json:"coordinate"` // 坐标
	Type       int       `json:"type"`       // 类型,0:文本,1:照片,2:视频,4:纯链接
	Views      int64     `json:"views"`      // 阅读数
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	NickName   string    `json:"nick_name"`
}
