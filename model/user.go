package model

import (
	"time"
)

type User struct {
	ID           uint          `gorm:"primaryKey"`
	FirstName    string        `gorm:"size:100;not null"`
	LastName     string        `gorm:"size:100;not null"`
	Email        string        `gorm:"size:255;unique;not null"`
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `gorm:"autoUpdateTime"`
	Borrowings   []Borrowing   `gorm:"foreignKey:UserID"`
	Reservations []Reservation `gorm:"foreignKey:UserID"`
	Fines        []Fine        `gorm:"foreignKey:UserID"`
}
