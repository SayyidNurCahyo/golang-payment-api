package repository

import (
	"database/sql"
	"merchant-payment-api/model"
)

type PaymentRepository interface {
	CreatePayment(payload model.Payment) error
	FindById(id string) (model.Payment, error)
	FindAll() ([]model.Payment, error)
}

type paymentRepository struct {
	db *sql.DB
}

// CreatePayment implements PaymentRepository.
func (p *paymentRepository) CreatePayment(payload model.Payment) error {
	payment, err := p.db.Begin()
	if err!=nil{
		return err
	}
	_, err = payment.Exec("insert into payment(id, date, customer_id, merchant_id, bank_id) values ($1, $2, $3, $4, $5)", payload.Id, payload.PaymentDate, payload.Customer.Id, payload.Merchant.Id, payload.Bank.Id)
	if err!=nil{
		return err
	}

	for _, detail := range payload.Details {
		_, err = payment.Exec("insert into payment_detail(id, payment_id, product_id, price, quantity) values ($1, $2, $3, $4, $5)", detail.Id, detail.Payment.Id, detail.Product.Id, detail.Product.Price, detail.Quantity)
		if err!=nil{
			return err
		}
	}
	if err:= payment.Commit(); err!=nil{
		return err
	}
	return nil
}

// FindAll implements PaymentRepository.
func (p *paymentRepository) FindAll() ([]model.Payment, error) {
	panic("unimplemented")
}

// FindById implements PaymentRepository.
func (p *paymentRepository) FindById(id string) (model.Payment, error) {
	panic("unimplemented")
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}
