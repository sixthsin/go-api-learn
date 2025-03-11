package fileshare

import (
	"go-api/cfg"
	"go-api/pkg/res"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type FileShareHandler struct {
	*cfg.Config
	*FileShareService
}

type FileshareHandlerDeps struct {
	*cfg.Config
	*FileShareService
}

func NewFileShareHandler(router *http.ServeMux, deps FileshareHandlerDeps) {
	handler := &FileShareHandler{
		Config:           deps.Config,
		FileShareService: deps.FileShareService,
	}
	router.HandleFunc("POST /file/upload", handler.UploadFile())
}

func (h *FileShareHandler) UploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		file, fileHandler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		filePath := filepath.Join(h.Config.Storage.Path, fileHandler.Filename)
		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, fileHandler.Filename+"uploaded successfully", http.StatusOK)
	}
}
