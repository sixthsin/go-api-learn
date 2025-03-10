package auth

import "go-api/internal/user"

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthSrevice(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}
