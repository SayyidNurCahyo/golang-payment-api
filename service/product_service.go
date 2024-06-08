package service

import (
	"fmt"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
)

type ProductService interface {
	Create(payload model.Product) error
	Delete(id string) error
	FindAll() ([]model.Product, error)
	FindById(id string) (model.Product, error)
	FindByName(name string) ([]model.Product, error)
	Update(payload model.Product) error
}

type productService struct {
	repo            repository.ProductRepository
	merchantService MerchantService
}

// FindByName implements ProductService.
func (p *productService) FindByName(name string) ([]model.Product, error) {
	products, err := p.repo.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}
	return products, nil
}

// Create implements ProductService.
func (p *productService) Create(payload model.Product) error {
	if payload.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.Merchant.Id == "" {
		return fmt.Errorf("merchant id is required")
	}

	_, err := p.merchantService.FindById(payload.Merchant.Id)
	if err != nil {
		return err
	}

	err = p.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new product: %v", err)
	}
	return nil
}

// Delete implements ProductService.
func (p *productService) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return fmt.Errorf("product not found")
	}

	err = p.repo.DeleteById(product.Id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}
	return nil
}

// FindAll implements ProductService.
func (p *productService) FindAll() ([]model.Product, error) {
	products, err := p.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all product: %v", err)
	}
	return products, nil
}

// FindById implements ProductService.
func (p *productService) FindById(id string) (model.Product, error) {
	product, err := p.repo.FindById(id)
	if err != nil {
		return model.Product{}, fmt.Errorf("product not found")
	}
	return product, nil
}

// Update implements ProductService.
func (p *productService) Update(payload model.Product) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.Price <= 0 {
		return fmt.Errorf("price is required")
	}
	if payload.Merchant.Id == "" {
		return fmt.Errorf("merchant id is required")
	}

	_, err := p.merchantService.FindById(payload.Merchant.Id)
	if err != nil {
		return err
	}

	_, err = p.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = p.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}
	return nil
}

func NewProductService(repo repository.ProductRepository, merchantService MerchantService) ProductService {
	return &productService{repo: repo, merchantService: merchantService}
}
