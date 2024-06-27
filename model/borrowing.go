package model

import (
	"time"
)

type Borrowing struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	User       User      `gorm:"foreignKey:UserID"`
	BookID     uint      `gorm:"not null"`
	Book       Book      `gorm:"foreignKey:BookID"`
	BorrowedAt time.Time `gorm:"autoCreateTime"`
	ReturnedAt *time.Time
}
