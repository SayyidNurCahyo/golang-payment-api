package service

import (
	"fmt"
	"merchant-payment-api/dto"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
)

type CustomerService interface {
	Delete(id string) error
	FindAll() ([]dto.GetCustomerResponse, error)
	FindById(id string) (dto.GetCustomerResponse, error)
	Update(payload dto.UpdateCustomerRequest) error
}

type customerService struct {
	customerRepo repository.CustomerRepository
}

func (b *customerService) Delete(id string) error {
	customer, err := b.customerRepo.FindById(id)
	if err!=nil{
		return fmt.Errorf("customer not found")
	}

	err = b.customerRepo.DeleteById(customer.Id)
	if err!=nil{
		return fmt.Errorf("failed to delete customer: %v", err)
	}
	return nil
}

func (b *customerService) FindAll() ([]dto.GetCustomerResponse, error) {
	customers, err := b.customerRepo.FindAll()
	if err!= nil{
		return nil, fmt.Errorf("failed to get all customer: %v", err)
	}
	responses := make([]dto.GetCustomerResponse, 0, len(customers))
	for _, customer := range customers {
		response := dto.GetCustomerResponse{
			Id: customer.Id,
			Name: customer.Name,
			PhoneNumber: customer.PhoneNumber,
			Username: customer.UserCredential.Username,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (b *customerService) FindById(id string) (dto.GetCustomerResponse, error) {
	customer, err := b.customerRepo.FindById(id)
	if err!=nil{
		return dto.GetCustomerResponse{}, fmt.Errorf("customer not found")
	}
	return dto.GetCustomerResponse{
		Id: customer.Id,
		Name: customer.Name,
		PhoneNumber: customer.PhoneNumber,
		Username: customer.UserCredential.Username,
	}, nil
}

func (b *customerService) Update(payload dto.UpdateCustomerRequest) error {
	if payload.Id == ""{
		return fmt.Errorf("id is required")
	}
	if payload.Name == ""{
		return fmt.Errorf("name is required")
	}
	if payload.PhoneNumber == ""{
		return fmt.Errorf("phone number is required")
	}

	currentCustomer, err := b.customerRepo.FindById(payload.Id)
	if err!= nil{
		return err
	}

	customer := model.Customer{
		Id: currentCustomer.Id,
		Name: payload.Name,
		PhoneNumber: payload.PhoneNumber,
	}
	err = b.customerRepo.Update(customer)
	if err!=nil{
		return fmt.Errorf("failed to update customer: %v", err)
	}
	return nil
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: repo,
	}
}
