package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MediumPrice float64        `json:"medium_price"`
	Author      string         `json:"author"`
	ImageUrl    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}
