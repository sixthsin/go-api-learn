package fileshare

import (
	"go-api/cfg"
	"go-api/internal/user"
	"io"
	"os"
	"path/filepath"
)

type FileShareService struct {
	*user.UserRepository
	*FileshareRepository
	*cfg.Config
}

func NewFileShareService(userRepo *user.UserRepository, repo *FileshareRepository, config *cfg.Config) *FileShareService {
	return &FileShareService{
		UserRepository:      userRepo,
		FileshareRepository: repo,
		Config:              config,
	}
}

func (s *FileShareService) SaveFile(file io.Reader, filename string, email string) (*user.User, error) {
	filePath := filepath.Join(s.Config.Storage.Path, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	user, err := s.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}
	return user, nil
}
