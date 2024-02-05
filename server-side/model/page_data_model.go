package model

import "time"

type PageData struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	SearchResultID uint64
	HtmlData       string `gorm:"type:text"`
}
