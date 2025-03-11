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
	router.HandleFunc("POST /auth/login", handler.Login())
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(jwt.JWTData{
			Email: body.Email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w,
			RegisterResponse{
				Token:    token,
				Email:    user.Email,
				Username: user.Username,
			},
			http.StatusCreated,
		)
	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userEmail, err := h.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(jwt.JWTData{
			Email: userEmail,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w,
			LoginResponse{
				Token: token,
			},
			http.StatusOK,
		)
	}
}
