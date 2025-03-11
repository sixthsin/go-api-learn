package fileshare

import (
	"go-api/internal/user"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Filename    string    `gorm:"not null;index"`
	Size        int64     `gorm:"not null"`
	ContentType string    `gorm:"not null"`
	Hash        string    `gorm:"unique;not null"`
	UserID      uint      `gorm:"not null;index"`
	User        user.User `gorm:"foreignKey:UserID"`
}
