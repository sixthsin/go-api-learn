package fileshare

import (
	"go-api/cfg"
	"io"
	"os"
	"path/filepath"
)

type FileShareService struct {
	*FileshareRepository
	*cfg.Config
}

func NewFileShareService(repo *FileshareRepository, config *cfg.Config) *FileShareService {
	return &FileShareService{
		FileshareRepository: repo,
		Config:              config,
	}
}

func (s *FileShareService) SaveFile(file io.Reader, filename string) error {
	filePath := filepath.Join(s.Config.Storage.Path, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}
	return nil
}
