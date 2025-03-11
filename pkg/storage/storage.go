package storage

import (
	"fmt"
	"go-api/cfg"
	"os"
)

func InitStorage(config *cfg.Config) {
	if _, err := os.Stat(config.Storage.Path); os.IsNotExist(err) {
		err := os.Mkdir(config.Storage.Path, os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Println("Storage initialized successfully")
	}
}
