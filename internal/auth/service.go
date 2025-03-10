package auth

import (
	"errors"
	"go-api/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthSrevice(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (s *AuthService) Register(email, password, username string) (*user.User, error) {
	existedUser, _ := s.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return nil, errors.New(ErrUserExists)
	}
	existedUser, _ = s.UserRepository.FindByUsername(username)
	if existedUser != nil {
		return nil, errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Username: username,
	}
	user, err = s.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
