package service

import (
	"fmt"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
)

type CustomerService interface {
	Create(payload model.Customer) error
	Delete(id string) error
	FindAll() ([]model.Customer, error)
	FindById(id string) (model.Customer, error)
	Update(payload model.Customer) error
}

type customerService struct {
	repo repository.CustomerRepository
}

func (m *customerService) Create(payload model.Customer) error {
	if payload.PhoneNumber==""{
		return fmt.Errorf("phone number is required")
	}
	if payload.Name==""{
		return fmt.Errorf("name is required")
	}

	err := m.repo.Save(payload)
	if err!=nil{
		return fmt.Errorf("failed to create new customer: %v", err)
	}
	return nil
}

func (m *customerService) Delete(id string) error {
	customer, err := m.FindById(id)
	if err!=nil{
		return fmt.Errorf("customer not found")
	}

	err = m.repo.DeleteById(customer.Id)
	if err!=nil{
		return fmt.Errorf("failed to delete customer: %v", err)
	}
	return nil
}

func (m *customerService) FindAll() ([]model.Customer, error) {
	customers, err := m.repo.FindAll()
	if err!= nil{
		return nil, fmt.Errorf("failed to get all customer: %v", err)
	}
	return customers, nil
}

func (m *customerService) FindById(id string) (model.Customer, error) {
	customer, err := m.repo.FindById(id)
	if err!=nil{
		return model.Customer{}, fmt.Errorf("customer not found")
	}
	return customer, nil
}

func (m *customerService) Update(payload model.Customer) error {
	if payload.Id == ""{
		return fmt.Errorf("id is required")
	}
	if payload.Name == ""{
		return fmt.Errorf("name is required")
	}
	if payload.PhoneNumber ==""{
		return fmt.Errorf("phone number is required")
	}
	
	_, err := m.FindById(payload.Id)
	if err!= nil{
		return err
	}

	err = m.repo.Update(payload)
	if err!=nil{
		return fmt.Errorf("failed to update customer: %v", err)
	}
	return nil
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}
