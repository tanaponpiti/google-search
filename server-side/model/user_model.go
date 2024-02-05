package model

import (
	"time"
)

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Username  string `gorm:"unique_index:idx_username"`
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCreate struct {
	Username string `json:"username" binding:"required"`
	//TODO add password required format
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UserUpdate struct {
	Password *string
	Name     *string
}

func (User) TableName() string {
	return "users"
}
