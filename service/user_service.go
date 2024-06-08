package service

import (
	"fmt"
	"merchant-payment-api/dto"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
	"merchant-payment-api/security"

	"github.com/google/uuid"
)

type UserService interface {
	FindByUsername(username string) (model.UserCredential, error)
	RegisterBank(payload dto.SaveBankRequest) error
}

type userService struct {
	userRepo repository.UserRepository
	bankRepo repository.BankRepository
}

// RegisterBank implements UserService.
func (u *userService) RegisterBank(payload dto.SaveBankRequest) error {
	if payload.Name == "" {
		return fmt.Errorf("bank name is required")
	}
	if payload.Username == "" {
		return fmt.Errorf("bank account username is required")
	}
	if payload.Password == "" {
		return fmt.Errorf("bank account password is required")
	}

	checkUser, _ := u.FindByUsername(payload.Username)
	if checkUser.Id != "" {
		return fmt.Errorf("username already used")
	}

	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	user := model.UserCredential{
		Id:       uuid.NewString(),
		Username: payload.Username,
		Password: hashPassword,
		IsActive: true,
	}
	err = u.userRepo.Save(user)
	if err != nil {
		return fmt.Errorf("failed to register user: %v", err.Error())
	}

	bank := model.Bank{
		Id:             uuid.NewString(),
		Name:           payload.Name,
		UserCredential: user,
	}
	err = u.bankRepo.Save(bank)
	if err != nil {
		return fmt.Errorf("failed to register bank: %v", err.Error())
	}
	return nil
}

// FindByUsername implements UserService.
func (u *userService) FindByUsername(username string) (model.UserCredential, error) {
	return u.userRepo.FindByUsername(username)
}

func NewUserService(userRepo repository.UserRepository, bankRepo repository.BankRepository) UserService {
	return &userService{
		userRepo: userRepo,
		bankRepo: bankRepo,
	}
}
