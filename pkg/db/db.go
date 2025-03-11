package db

import (
	"go-api/cfg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDb(config *cfg.Config) *DB {
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{db}
}
