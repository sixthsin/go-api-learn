package fileshare

type FileShareService struct {
	FileshareRepository *FileshareRepository
}

func NewFileShareService(fileshareRepository *FileshareRepository) *FileShareService {
	return &FileShareService{FileshareRepository: fileshareRepository}
}
