package main

import (
	"fmt"
	"go-api/cfg"
	"go-api/internal/auth"
	"go-api/internal/user"
	"go-api/pkg/db"
	"go-api/pkg/middleware"
	"net/http"
)

func main() {
	config := cfg.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()

	// Repositories
	userRepository := user.NewUserRepository(db)

	// Services
	authService := auth.NewAuthSrevice(userRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})

	stack := middleware.Chain()

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
