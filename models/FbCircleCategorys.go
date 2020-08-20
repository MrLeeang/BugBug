package models

import "time"

type FbCircleCategorys struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"` // 分类名
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
