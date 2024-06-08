package service

import (
	"fmt"
	"merchant-payment-api/dto"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
	"time"

	"github.com/google/uuid"
)

type PaymentService interface {
	CreatePayment(payload dto.PaymentRequest) error
	FindAll() ([]dto.PaymentResponse, error)
	FindById(id string) (dto.PaymentResponse, error)
}

type paymentService struct {
	paymentRepo     repository.PaymentRepository
	customerService CustomerService
	merchantService MerchantService
	bankService     BankService
	productService  ProductService
}

// FindAll implements PaymentService.
func (p *paymentService) FindAll() ([]dto.PaymentResponse, error) {
	payments, err := p.paymentRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all payments: %v", err.Error())
	}
	responses := make([]dto.PaymentResponse, 0, len(payments))
	for _, payment := range payments {
		details := make([]dto.PaymentDetailResponse, 0, len(payment.Details))
		for _, detail := range payment.Details {
			detailResponse := dto.PaymentDetailResponse{
				DetailId:     detail.Id,
				ProductId:    detail.Product.Id,
				ProductName:  detail.Product.Name,
				ProductPrice: detail.Price,
				Quantity:     detail.Quantity,
			}
			details = append(details, detailResponse)
		}
		response := dto.PaymentResponse{
			Id:                    payment.Id,
			PaymentDate:           payment.PaymentDate,
			CustomerId:            payment.Customer.Id,
			CustomerName:          payment.Customer.Name,
			MerchantId:            payment.Merchant.Id,
			MerchantName:          payment.Merchant.Name,
			BankId:                payment.Bank.Id,
			BankName:              payment.Bank.Name,
			PaymentDetailResponse: details,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

// FindById implements PaymentService.
func (p *paymentService) FindById(id string) (dto.PaymentResponse, error) {
	payment, err := p.paymentRepo.FindById(id)
	if err != nil {
		return dto.PaymentResponse{}, fmt.Errorf("failed to get payment: %v", err.Error())
	}

	details := make([]dto.PaymentDetailResponse, 0, len(payment.Details))
	for _, detail := range payment.Details {
		detailResponse := dto.PaymentDetailResponse{
			DetailId:     detail.Id,
			ProductId:    detail.Product.Id,
			ProductName:  detail.Product.Name,
			ProductPrice: detail.Price,
			Quantity:     detail.Quantity,
		}
		details = append(details, detailResponse)
	}
	response := dto.PaymentResponse{
		Id:                    payment.Id,
		PaymentDate:           payment.PaymentDate,
		CustomerId:            payment.Customer.Id,
		CustomerName:          payment.Customer.Name,
		MerchantId:            payment.Merchant.Id,
		MerchantName:          payment.Merchant.Name,
		BankId:                payment.Bank.Id,
		BankName:              payment.Bank.Name,
		PaymentDetailResponse: details,
	}

	return response, nil
}

// CreatePayment implements PaymentService.
func (p *paymentService) CreatePayment(payload dto.PaymentRequest) error {
	if payload.CustomerId == "" {
		return fmt.Errorf("customer id is required")
	}
	if payload.MerchantId == "" {
		return fmt.Errorf("merchant id is required")
	}
	if payload.BankId == "" {
		return fmt.Errorf("bank id is required")
	}
	if len(payload.PaymentDetailRequest) == 0 {
		return fmt.Errorf("payment detail is required")
	}

	var payment model.Payment
	payment.Id = uuid.NewString()
	payment.PaymentDate = time.Now()
	customer, err := p.customerService.FindById(payload.CustomerId)
	if err != nil {
		return fmt.Errorf("customer not found")
	}
	payment.Customer = model.Customer{
		Id:          customer.Id,
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
	}
	merchant, err := p.merchantService.FindById(payload.MerchantId)
	if err != nil {
		return fmt.Errorf("merchant not found")
	}
	payment.Merchant = model.Merchant{
		Id:          merchant.Id,
		Name:        merchant.Name,
		PhoneNumber: merchant.PhoneNumber,
		Address:     merchant.Address,
	}
	bank, err := p.bankService.FindById(payload.BankId)
	if err != nil {
		return fmt.Errorf("bank not found")
	}
	payment.Bank = model.Bank{
		Id:   bank.Id,
		Name: bank.Name,
	}
	paymentDetails := make([]model.PaymentDetail, 0, len(payload.PaymentDetailRequest))
	for _, detailRequest := range payload.PaymentDetailRequest {
		var detail model.PaymentDetail
		product, err := p.productService.FindById(detailRequest.ProductId)
		if err != nil {
			return fmt.Errorf("product not found")
		}
		if product.MerchantId != payload.MerchantId {
			return fmt.Errorf("product not found in merchant store")
		}
		detail.Id = uuid.NewString()
		detail.Payment = payment
		detail.Product = model.Product{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		}
		detail.Price = product.Price
		detail.Quantity = detailRequest.Quantity
		paymentDetails = append(paymentDetails, detail)
	}
	payment.Details = paymentDetails

	err = p.paymentRepo.CreatePayment(payment)
	if err != nil {
		return fmt.Errorf("failed to create payment")
	}
	return nil
}

func NewPaymentService(paymentRepo repository.PaymentRepository, customerService CustomerService, merchantService MerchantService, bankService BankService, productService ProductService) PaymentService {
	return &paymentService{
		paymentRepo:     paymentRepo,
		customerService: customerService,
		merchantService: merchantService,
		bankService:     bankService,
		productService:  productService,
	}
}
