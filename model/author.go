package model

import (
	"time"
)

type Author struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Books     []Book    `gorm:"foreignKey:AuthorID"`
}
