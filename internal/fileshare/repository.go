package fileshare

import "go-api/pkg/db"

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
