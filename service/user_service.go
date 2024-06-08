package service

import (
	"fmt"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
	"merchant-payment-api/security"

	"github.com/google/uuid"
)

type UserService interface {
	FindByUsername(username string) (model.UserCredential, error)
	Register(payload model.UserCredential) error
}

type userService struct {
	userRepo repository.UserRepository
}

// FindByUsername implements UserService.
func (u *userService) FindByUsername(username string) (model.UserCredential, error) {
	return u.userRepo.FindByUsername(username)
}

// Register implements UserService.
func (u *userService) Register(payload model.UserCredential) error {
	hashPassword, err := security.HashPassword(payload.Password)
	if err!=nil{
		return err
	}

	payload.Id = uuid.NewString()
	payload.Password = hashPassword
	payload.IsActive = true
	err = u.userRepo.Save(payload)
	if err!=nil{
		return fmt.Errorf("failed to register user: %v", err.Error())
	}
	return nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}
