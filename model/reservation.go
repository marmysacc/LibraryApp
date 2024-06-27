package model

import (
	"time"
)

type Reservation struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	User       User      `gorm:"foreignKey:UserID"`
	BookID     uint      `gorm:"not null"`
	Book       Book      `gorm:"foreignKey:BookID"`
	ReservedAt time.Time `gorm:"autoCreateTime"`
	Status     string    `gorm:"size:20;not null"`
}
