package service

import (
	"fmt"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
)

type MerchantService interface {
	Create(payload model.Merchant) error
	Delete(id string) error
	FindAll() ([]model.Merchant, error)
	FindById(id string) (model.Merchant, error)
	Update(payload model.Merchant) error
}

type merchantService struct {
	repo repository.MerchantRepository
}

func (m *merchantService) Create(payload model.Merchant) error {
	if payload.PhoneNumber==""{
		return fmt.Errorf("phone number is required")
	}
	if payload.Name==""{
		return fmt.Errorf("name is required")
	}
	if payload.Address==""{
		return fmt.Errorf("address is required")
	}

	err := m.repo.Save(payload)
	if err!=nil{
		return fmt.Errorf("failed to create new merchant: %v", err)
	}
	return nil
}

func (m *merchantService) Delete(id string) error {
	merchant, err := m.FindById(id)
	if err!=nil{
		return fmt.Errorf("merchant not found")
	}

	err = m.repo.DeleteById(merchant.Id)
	if err!=nil{
		return fmt.Errorf("failed to delete merchant: %v", err)
	}
	return nil
}

func (m *merchantService) FindAll() ([]model.Merchant, error) {
	merchants, err := m.repo.FindAll()
	if err!= nil{
		return nil, fmt.Errorf("failed to get all merchant: %v", err)
	}
	return merchants, nil
}

func (m *merchantService) FindById(id string) (model.Merchant, error) {
	merchant, err := m.repo.FindById(id)
	if err!=nil{
		return model.Merchant{}, fmt.Errorf("merchant not found")
	}
	return merchant, nil
}

func (m *merchantService) Update(payload model.Merchant) error {
	if payload.Id == ""{
		return fmt.Errorf("id is required")
	}
	if payload.Name == ""{
		return fmt.Errorf("name is required")
	}
	if payload.PhoneNumber ==""{
		return fmt.Errorf("phone number is required")
	}
	if payload.Address == ""{
		return fmt.Errorf("address is required")
	}

	_, err := m.FindById(payload.Id)
	if err!= nil{
		return err
	}

	err = m.repo.Update(payload)
	if err!=nil{
		return fmt.Errorf("failed to update merchant: %v", err)
	}
	return nil
}

func NewMerchantService(repo repository.MerchantRepository) MerchantService {
	return &merchantService{repo: repo}
}
