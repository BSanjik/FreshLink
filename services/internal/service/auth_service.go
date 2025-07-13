package service

import (
	"errors"
	"services/internal/model/dto"
	"services/internal/model/entity"
	"services/internal/storage"
	"services/internal/utils"
)

type AuthService struct {
	userRepo *storage.UserRepository
}

func NewAuthService(userRepo *storage.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: req.Username,
		Phone:    req.Phone,
		Password: hashedPassword,
		Role:     "user",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Phone    string `json:"phone"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		}{
			ID:       user.ID,
			Username: user.Username,
			Phone:    user.Phone,
			Email:    user.Email,
			Role:     user.Role,
		},
	}, nil
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.GetByPhone(req.Phone)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Phone    string `json:"phone"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		}{
			ID:       user.ID,
			Username: user.Username,
			Phone:    user.Phone,
			Email:    user.Email,
			Role:     user.Role,
		},
	}, nil
}
