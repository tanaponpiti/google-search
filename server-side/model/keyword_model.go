package model

import (
	"time"
)

type Keyword struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	KeywordText   string `gorm:"index:idx_keyword_text;unique;type:varchar(255);not null"`
	SearchResults []SearchResult
}

type KeywordCreate struct {
	Keywords []string
}

type KeywordFilter struct {
	KeywordSearch *string
	Status        *[]SearchStatus
}
