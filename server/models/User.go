package models

import (
	"time"
)

type User struct {
	UserID    int       `gorm:"primaryKey" json:"userID"`
	FullName  string    `json:"fullName"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
