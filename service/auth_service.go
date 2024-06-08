package service

import (
	"fmt"
	"merchant-payment-api/dto"
	"merchant-payment-api/repository"
	"merchant-payment-api/security"
)

type AuthService interface {
	Login(payload dto.LoginRequest) (dto.LoginResponse, error)
}

type authService struct {
	repo repository.UserRepository
}

// Login implements AuthService.
func (a *authService) Login(payload dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := a.repo.FindByUsername(payload.Username)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unauthorized: invalid credential")
	}

	err = security.VerifyPassword(user.Password, payload.Password)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unauthorized: invalid credential")
	}

	token, err := security.GenerateToken(user)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	return dto.LoginResponse{
		Username: user.Username,
		Token:    token,
	}, err
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}
