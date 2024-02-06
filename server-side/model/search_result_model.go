package model

import "time"

type SearchResult struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	KeywordID    uint
	AdWordsCount *int
	TotalLinks   *int
	TotalResults *string
	SearchDate   *time.Time
	PageData     *PageData
	Status       SearchStatus `gorm:"index,type:enum('COMPLETED', 'PENDING', 'FAILED')"`
}

type SearchStatus string

const (
	Completed SearchStatus = "COMPLETED"
	Pending   SearchStatus = "PENDING"
	Failed    SearchStatus = "FAILED"
)
