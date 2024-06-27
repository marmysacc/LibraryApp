package model

import (
	"time"
)

type Genre struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Books     []Book    `gorm:"foreignKey:GenreID"`
}
