package models

import "time"

type FbCircles struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`         // 圈子名称
	Introduction string    `json:"introduction"` // 介绍
	Category     int       `json:"category"`     // 分类
	Owner        int       `json:"owner"`        // 所有者
	Status       int       `json:"status"`       // 状态:0,1
	CoverBg      string    `json:"cover_bg"`     // 封面图
	Logo         string    `json:"logo"`         // logo
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
