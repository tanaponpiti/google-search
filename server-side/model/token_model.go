package model

import (
	"time"
)

type Token struct {
	Token     string `gorm:"index:idx_token,unique"`
	UserId    int64
	CreatedAt time.Time
}
