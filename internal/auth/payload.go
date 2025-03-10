package auth

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	// Token string `json:"token"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
