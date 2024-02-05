package model

import (
	"time"
)

type Keyword struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	KeywordText   string `gorm:"type:varchar(255);not null"`
	SearchResults []SearchResult
}
