package model

import (
	"time"
)

type Author struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Books     []Book    `gorm:"foreignKey:AuthorID"`
}
