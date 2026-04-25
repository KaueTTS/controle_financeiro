package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"not null"`
	Description *string   `gorm:"null"`
	Amount      int64     `gorm:"not null"`
	Category    string    `gorm:"not null"`
	Type        string    `gorm:"not null"` // "income" ou "expense"
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
