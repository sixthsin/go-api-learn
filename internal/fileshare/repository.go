package fileshare

import (
	"go-api/pkg/db"
)

type FileshareRepository struct {
	Database *db.DB
}

func NewFileShareRepository(database *db.DB) *FileshareRepository {
	return &FileshareRepository{Database: database}
}

func (r *FileshareRepository) CreateFile(file *File) error {
	result := r.Database.DB.Create(file)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *FileshareRepository) GetFileByHash(hash string) (*File, error) {
	var file File
	result := r.Database.DB.Where("hash = ?", hash).First(&file)
	if result.Error != nil {
		return nil, result.Error
	}
	return &file, nil
}
