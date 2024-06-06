package service

import (
	"fmt"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
)

type BankService interface {
	Create(payload model.Bank) error
	Delete(id int) error
	FindAll() ([]model.Bank, error)
	FindById(id int) (model.Bank, error)
	Update(payload model.Bank) error
}

type bankService struct {
	repo repository.BankRepository
}

func (m *bankService) Create(payload model.Bank) error {
	if payload.Name==""{
		return fmt.Errorf("name is required")
	}

	err := m.repo.Save(payload)
	if err!=nil{
		return fmt.Errorf("failed to create new bank: %v", err)
	}
	return nil
}

func (m *bankService) Delete(id int) error {
	bank, err := m.FindById(id)
	if err!=nil{
		return fmt.Errorf("bank not found")
	}

	err = m.repo.DeleteById(bank.Id)
	if err!=nil{
		return fmt.Errorf("failed to delete bank: %v", err)
	}
	return nil
}

func (m *bankService) FindAll() ([]model.Bank, error) {
	banks, err := m.repo.FindAll()
	if err!= nil{
		return nil, fmt.Errorf("failed to get all bank: %v", err)
	}
	return banks, nil
}

func (m *bankService) FindById(id int) (model.Bank, error) {
	bank, err := m.repo.FindById(id)
	if err!=nil{
		return model.Bank{}, fmt.Errorf("bank not found")
	}
	return bank, nil
}

func (m *bankService) Update(payload model.Bank) error {
	if payload.Id == 0{
		return fmt.Errorf("id is required")
	}
	if payload.Name == ""{
		return fmt.Errorf("name is required")
	}

	_, err := m.FindById(payload.Id)
	if err!= nil{
		return err
	}

	err = m.repo.Update(payload)
	if err!=nil{
		return fmt.Errorf("failed to update bank: %v", err)
	}
	return nil
}

func NewBankService(repo repository.BankRepository) BankService {
	return &bankService{repo: repo}
}
