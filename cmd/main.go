package main

import (
	"fmt"
	"go-api/cfg"
	"go-api/internal/auth"
	"go-api/internal/fileshare"
	"go-api/internal/user"
	"go-api/pkg/db"
	"go-api/pkg/middleware"
	"go-api/pkg/storage"
	"net/http"
)

func main() {
	config := cfg.LoadConfig()
	db := db.NewDb(config)
	storage.InitStorage(config)
	router := http.NewServeMux()

	// Repositories
	userRepository := user.NewUserRepository(db)
	fileshareRepository := fileshare.NewFileShareRepository(db)

	// Services
	authService := auth.NewAuthSrevice(userRepository)
	fileshareService := fileshare.NewFileShareService(fileshareRepository, config)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})
	fileshare.NewFileShareHandler(router, fileshare.FileshareHandlerDeps{
		Config:           config,
		FileShareService: fileshareService,
	})

	stack := middleware.Chain()

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
