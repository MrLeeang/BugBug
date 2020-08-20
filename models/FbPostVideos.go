package models

import "time"

type FbPostVideos struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`      // 用户UID
	Pid       int64     `json:"pid"`      // 文章ID
	Url       string    `json:"url"`      // 资源url
	Cover     string    `json:"cover"`    // 封面
	Duration  float32   `json:"duration"` // 视频长度,单位秒
	Width     float32   `json:"width"`    // 宽
	Height    float32   `json:"height"`   // 高
	Views     int64     `json:"views"`    // 播放数
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
