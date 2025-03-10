package auth

import (
	"go-api/cfg"
	"go-api/pkg/jwt"
	"go-api/pkg/req"
	"go-api/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*cfg.Config
	*AuthService
}

type AuthHandlerDeps struct {
	*cfg.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	// router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := h.AuthService.Register(body.Email, body.Password, body.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(jwt.JWTData{
			Email:    body.Email,
			Username: body.Username,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := RegisterResponse{
			Token:    token,
			Email:    user.Email,
			Username: user.Username,
		}
		res.Json(w, data, http.StatusCreated)
	}
}
