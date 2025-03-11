package fileshare

import (
	"go-api/cfg"
	"go-api/pkg/middleware"
	"go-api/pkg/res"
	"net/http"
)

const (
	MsgSuccessfullyUpload = "file uploaded successfully"
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
	router.Handle("POST /file/upload", middleware.IsAuthed(handler.UploadFile(), deps.Config))
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
		err = h.FileShareService.SaveFile(file, fileHandler.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, MsgSuccessfullyUpload, http.StatusOK)
	}
}
