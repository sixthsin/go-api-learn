package user

import "go-api/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{Database: db}
}
