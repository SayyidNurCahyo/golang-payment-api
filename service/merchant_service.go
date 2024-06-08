package service

import (
	"fmt"
	"merchant-payment-api/dto"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
)

type MerchantService interface {
	Delete(id string) error
	FindAll() ([]dto.GetMerchantResponse, error)
	FindById(id string) (dto.GetMerchantResponse, error)
	Update(payload dto.UpdateMerchantRequest) error
}

type merchantService struct {
	merchantRepo repository.MerchantRepository
}

func (b *merchantService) Delete(id string) error {
	merchant, err := b.merchantRepo.FindById(id)
	if err!=nil{
		return fmt.Errorf("merchant not found")
	}

	err = b.merchantRepo.DeleteById(merchant.Id)
	if err!=nil{
		return fmt.Errorf("failed to delete merchant: %v", err)
	}
	return nil
}

func (b *merchantService) FindAll() ([]dto.GetMerchantResponse, error) {
	merchants, err := b.merchantRepo.FindAll()
	if err!= nil{
		return nil, fmt.Errorf("failed to get all merchant: %v", err)
	}
	responses := make([]dto.GetMerchantResponse, 0, len(merchants))
	for _, merchant := range merchants {
		response := dto.GetMerchantResponse{
			Id: merchant.Id,
			Name: merchant.Name,
			PhoneNumber: merchant.PhoneNumber,
			Address: merchant.Address,
			Username: merchant.UserCredential.Username,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (b *merchantService) FindById(id string) (dto.GetMerchantResponse, error) {
	merchant, err := b.merchantRepo.FindById(id)
	if err!=nil{
		return dto.GetMerchantResponse{}, fmt.Errorf("merchant not found")
	}
	return dto.GetMerchantResponse{
		Id: merchant.Id,
		Name: merchant.Name,
		PhoneNumber: merchant.PhoneNumber,
		Address: merchant.Address,
		Username: merchant.UserCredential.Username,
	}, nil
}

func (b *merchantService) Update(payload dto.UpdateMerchantRequest) error {
	if payload.Id == ""{
		return fmt.Errorf("id is required")
	}
	if payload.Name == ""{
		return fmt.Errorf("name is required")
	}
	if payload.PhoneNumber == ""{
		return fmt.Errorf("phone number is required")
	}
	if payload.Address == ""{
		return fmt.Errorf("address is required")
	}

	currentmerchant, err := b.merchantRepo.FindById(payload.Id)
	if err!= nil{
		return err
	}

	merchant := model.Merchant{
		Id: currentmerchant.Id,
		Name: payload.Name,
		PhoneNumber: payload.PhoneNumber,
		Address: payload.Address,
	}
	err = b.merchantRepo.Update(merchant)
	if err!=nil{
		return fmt.Errorf("failed to update merchant: %v", err)
	}
	return nil
}

func NewMerchantService(repo repository.MerchantRepository) MerchantService {
	return &merchantService{
		merchantRepo: repo,
	}
}
