package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Username string `gorm:"uniqueIndex"`
	Password string
}
