package dbmodels

import (
	"gorm.io/gorm"
	"time"
)

// Separate GORM model with tags for DB

type UserDB struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
