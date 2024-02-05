package model

import "time"

type SearchResult struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	KeywordID    uint64
	AdWordsCount int
	TotalLinks   int
	TotalResults string
	SearchDate   time.Time
	PageData     PageData
}
