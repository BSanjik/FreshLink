package dto

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required,min=10"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Phone    string `json:"phone" validate:"required,min=10"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	} `json:"user"`
}
