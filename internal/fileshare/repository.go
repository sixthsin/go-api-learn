package fileshare

import "go-api/pkg/db"

type FileshareRepository struct {
	Database *db.DB
}

func NewFileShareRepository(database *db.DB) *FileshareRepository {
	return &FileshareRepository{Database: database}
}
