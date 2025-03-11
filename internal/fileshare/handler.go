package fileshare

import (
	"fmt"
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
	router.Handle("POST /file/upload", middleware.IsAuthed(handler.Upload(), deps.Config))
}

func (h *FileShareHandler) Upload() http.HandlerFunc {
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
		emailValue := r.Context().Value(middleware.ContextEmailKey)
		fmt.Println(emailValue)
		email, ok := emailValue.(string)
		if !ok {
			http.Error(w, "email not found", http.StatusInternalServerError)
			return
		}
		user, err := h.FileShareService.SaveFile(file, fileHandler.Filename, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		createdFile := NewFile(fileHandler.Filename, fileHandler.Size, user.ID)
		createdFile.GenerateHash()
		err = h.FileShareService.CreateFile(createdFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, MsgSuccessfullyUpload, http.StatusOK)
	}
}
