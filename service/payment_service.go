package service

import (
	"fmt"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
	"time"

	"github.com/google/uuid"
)

type PaymentService interface {
	CreatePayment(payload model.Payment) (model.Payment, error)
}

type paymentService struct {
	paymentRepo     repository.PaymentRepository
	customerService CustomerService
	merchantService MerchantService
	bankService     BankService
	productService  ProductService
}

// CreatePayment implements PaymentService.
func (p *paymentService) CreatePayment(payload model.Payment) (model.Payment, error) {
	var payment model.Payment
	payment.Id = uuid.NewString()
	payment.PaymentDate = time.Now()
	// customer, err := p.customerService.FindById(payload.Customer.Id)
	// if err != nil {
	// 	return model.Payment{}, fmt.Errorf("customer not found")
	// }
	// payment.Customer = customer
	merchant, err := p.merchantService.FindById(payload.Merchant.Id)
	if err != nil {
		return model.Payment{}, fmt.Errorf("merchant not found")
	}
	payment.Merchant = merchant
	// bank, err := p.bankService.FindById(payload.Bank.Id)
	// if err != nil {
	// 	return model.Payment{}, fmt.Errorf("bank not found")
	// }
	// payment.Bank = bank
	paymentDetails := make([]model.PaymentDetail, 0, len(payload.Details))
	for _, detail := range payload.Details {
		product, err := p.productService.FindById(detail.Product.Id)
		if err != nil {
			return model.Payment{}, fmt.Errorf("product not found")
		}
		detail.Id = uuid.NewString()
		detail.Payment = payment
		detail.Product = product
		detail.Price = product.Price
		paymentDetails = append(paymentDetails, detail)
	}
	payment.Details = paymentDetails

	err = p.paymentRepo.CreatePayment(payment)
	if err!=nil{
		return model.Payment{}, fmt.Errorf("failed to create payment")
	}
	return payload, nil
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
