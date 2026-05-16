package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Title       string         `gorm:"type:text;not null"`
	Description *string        `gorm:"type:text"`
	Amount      float64        `gorm:"not null"`
	Category    string         `gorm:"type:varchar(100);not null;index"`
	Type        string         `gorm:"type:varchar(20);not null;index"` // "income" ou "expense"
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
