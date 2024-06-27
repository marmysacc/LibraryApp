package model

import (
    "time"
)

type Fine struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    User      User      `gorm:"foreignKey:UserID"`
    Amount    float64   `gorm:"not null"`
    IssuedAt  time.Time `gorm:"autoCreateTime"`
    PaidAt    *time.Time
}
