package fileshare

import (
	"go-api/internal/user"
	"math/rand"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Filename string    `gorm:"not null;index"`
	Size     int64     `gorm:"not null"`
	Hash     string    `gorm:"unique;not null"`
	UserID   uint      `gorm:"not null;index"`
	User     user.User `gorm:"foreignKey:UserID"`
}

func NewFile(filename string, size int64, userID uint) *File {
	return &File{
		Filename: filename,
		Size:     size,
		UserID:   userID,
	}
}

func (f *File) GenerateHash() {
	f.Hash = RandStringRunes(6)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
