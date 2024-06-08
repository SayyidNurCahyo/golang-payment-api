package service

import (
	"fmt"
	"merchant-payment-api/dto"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(payload dto.SaveProductRequest) error
	Delete(id string) error
	FindAll() ([]dto.GetProductResponse, error)
	FindById(id string) (dto.GetProductResponse, error)
	FindByName(name string) ([]dto.GetProductResponse, error)
	Update(payload dto.UpdateProductRequest) error
}

type productService struct {
	repo            repository.ProductRepository
	merchantService MerchantService
}

// FindByName implements ProductService.
func (p *productService) FindByName(name string) ([]dto.GetProductResponse, error) {
	products, err := p.repo.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}
	responses := make([]dto.GetProductResponse, 0, len(products))
	for _, product := range products {
		response := dto.GetProductResponse{
			Id:           product.Id,
			MerchantId:   product.Merchant.Id,
			MerchantName: product.Merchant.Name,
			Name:         product.Name,
			Price:        product.Price,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

// Create implements ProductService.
func (p *productService) Create(payload dto.SaveProductRequest) error {
	if payload.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.MerchantId == "" {
		return fmt.Errorf("merchant id is required")
	}

	_, err := p.merchantService.FindById(payload.MerchantId)
	if err != nil {
		return err
	}

	product := model.Product{
		Id:          uuid.NewString(),
		Merchant:    model.Merchant{Id: payload.MerchantId},
		Name:        payload.Name,
		Price:       payload.Price,
		IsAvailable: true,
	}
	err = p.repo.Save(product)
	if err != nil {
		return fmt.Errorf("failed to create new product: %v", err)
	}
	return nil
}

// Delete implements ProductService.
func (p *productService) Delete(id string) error {
	product, err := p.repo.FindById(id)
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
func (p *productService) FindAll() ([]dto.GetProductResponse, error) {
	products, err := p.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all product: %v", err)
	}
	responses := make([]dto.GetProductResponse, 0, len(products))
	for _, product := range products {
		response := dto.GetProductResponse{
			Id:           product.Id,
			MerchantId:   product.Merchant.Id,
			MerchantName: product.Merchant.Name,
			Name:         product.Name,
			Price:        product.Price,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

// FindById implements ProductService.
func (p *productService) FindById(id string) (dto.GetProductResponse, error) {
	product, err := p.repo.FindById(id)
	if err != nil {
		return dto.GetProductResponse{}, fmt.Errorf("product not found")
	}
	return dto.GetProductResponse{
		Id:           product.Id,
		MerchantId:   product.Merchant.Id,
		MerchantName: product.Merchant.Name,
		Name:         product.Name,
		Price:        product.Price,
	}, nil
}

// Update implements ProductService.
func (p *productService) Update(payload dto.UpdateProductRequest) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if payload.Price <= 0 {
		return fmt.Errorf("price is required")
	}

	currentProduct, err := p.repo.FindById(payload.Id)
	if err != nil {
		return err
	}

	product := model.Product{
		Id:          currentProduct.Id,
		Merchant:    currentProduct.Merchant,
		Name:        payload.Name,
		Price:       payload.Price,
		IsAvailable: currentProduct.IsAvailable,
	}
	err = p.repo.Update(product)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}
	return nil
}

func NewProductService(repo repository.ProductRepository, merchantService MerchantService) ProductService {
	return &productService{repo: repo, merchantService: merchantService}
}
