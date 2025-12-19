package models

import "time"

type Role struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:50;not null;uniqueIndex"`
	Code        string    `gorm:"size:100;not null;uniqueIndex"`
	Description *string   `gorm:"text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
