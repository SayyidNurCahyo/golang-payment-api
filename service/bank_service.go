package service

import (
	"fmt"
	"merchant-payment-api/dto"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
)

type BankService interface {
	Delete(id string) error
	FindAll() ([]dto.GetBankResponse, error)
	FindById(id string) (dto.GetBankResponse, error)
	Update(payload dto.UpdateBankRequest) error
}

type bankService struct {
	bankRepo repository.BankRepository
}

func (b *bankService) Delete(id string) error {
	bank, err := b.bankRepo.FindById(id)
	if err != nil {
		return fmt.Errorf("bank not found")
	}

	err = b.bankRepo.DeleteById(bank.Id)
	if err != nil {
		return fmt.Errorf("failed to delete bank: %v", err)
	}
	return nil
}

func (b *bankService) FindAll() ([]dto.GetBankResponse, error) {
	banks, err := b.bankRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all bank: %v", err)
	}
	responses := make([]dto.GetBankResponse, 0, len(banks))
	for _, bank := range banks {
		response := dto.GetBankResponse{
			Id:       bank.Id,
			Name:     bank.Name,
			Username: bank.UserCredential.Username,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (b *bankService) FindById(id string) (dto.GetBankResponse, error) {
	bank, err := b.bankRepo.FindById(id)
	if err != nil {
		return dto.GetBankResponse{}, fmt.Errorf("bank not found")
	}
	return dto.GetBankResponse{
		Id:       bank.Id,
		Name:     bank.Name,
		Username: bank.UserCredential.Username,
	}, nil
}

func (b *bankService) Update(payload dto.UpdateBankRequest) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	currentBank, err := b.bankRepo.FindById(payload.Id)
	if err != nil {
		return err
	}

	bank := model.Bank{
		Id:   currentBank.Id,
		Name: payload.Name,
	}
	err = b.bankRepo.Update(bank)
	if err != nil {
		return fmt.Errorf("failed to update bank: %v", err)
	}
	return nil
}

func NewBankService(repo repository.BankRepository) BankService {
	return &bankService{
		bankRepo: repo,
	}
}
