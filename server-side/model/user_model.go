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
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
	Name     string `json:"name" validate:"required"`
}

type UserUpdate struct {
	Password *string
	Name     *string
}

func (User) TableName() string {
	return "users"
}
