package model

import "time"

type PageData struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	SearchResultID uint   `gorm:"unique_index:idx_search_result_id"`
	HtmlData       string `gorm:"type:text"`
}
