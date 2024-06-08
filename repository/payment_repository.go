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
	if err != nil {
		return err
	}
	_, err = payment.Exec("insert into payment(id, date, customer_id, merchant_id, bank_id) values ($1, $2, $3, $4, $5)", payload.Id, payload.PaymentDate, payload.Customer.Id, payload.Merchant.Id, payload.Bank.Id)
	if err != nil {
		return err
	}

	for _, detail := range payload.Details {
		_, err = payment.Exec("insert into payment_detail(id, payment_id, product_id, price, quantity) values ($1, $2, $3, $4, $5)", detail.Id, detail.Payment.Id, detail.Product.Id, detail.Product.Price, detail.Quantity)
		if err != nil {
			return err
		}
	}
	if err := payment.Commit(); err != nil {
		return err
	}
	return nil
}

// FindAll implements PaymentRepository.
func (p *paymentRepository) FindAll() ([]model.Payment, error) {
	rows, err := p.db.Query("select p.id, p.date, c.id, c.name, m.id, m.name, b.id, b.name, pd.id, pr.id, pr.name, pr.price, pd.quantity from payment as p join customer as c on c.id=p.customer_id join merchant as m on m.id=p.merchant_id join bank as b on b.id=p.bank_id join payment_detail as pd on pd.payment_id=p.id join product as pr on pr.id=pd.product_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []model.Payment

	for rows.Next() {
		var (
			p  model.Payment
			pd model.PaymentDetail
			c  model.Customer
			m  model.Merchant
			b  model.Bank
			pr model.Product
		)
		err := rows.Scan(
			&p.Id, &p.PaymentDate, &c.Id, &c.Name,
			&m.Id, &m.Name, &b.Id, &b.Name,
			&pd.Id, &pr.Id, &pr.Name, &pd.Price, &pd.Quantity,
		)
		if err != nil {
			return nil, err
		}
		p.Customer = c
		p.Merchant = m
		p.Bank = b
		pd.Product = pr

		existingPayment := func() int {
			for i, existing := range payments {
				if existing.Id == p.Id {
					return i
				}
			}
			return -1
		}()

		if existingPayment == -1 {
			p.Details = []model.PaymentDetail{pd}
			payments = append(payments, p)
		} else {
			payments[existingPayment].Details = append(payments[existingPayment].Details, pd)
		}
	}

	return payments, nil
}

// FindById implements PaymentRepository.
func (p *paymentRepository) FindById(id string) (model.Payment, error) {
	rows, err := p.db.Query("select p.id, p.date, c.id, c.name, m.id, m.name, b.id, b.name, pd.id, pr.id, pr.name, pr.price, pd.quantity from payment as p join customer as c on c.id=p.customer_id join merchant as m on m.id=p.merchant_id join bank as b on b.id=p.bank_id join payment_detail as pd on pd.payment_id=p.id join product as pr on pr.id=pd.product_id")
	if err != nil {
		return model.Payment{}, err
	}
	var (
		payment model.Payment
		details []model.PaymentDetail
		pd      model.PaymentDetail
		c       model.Customer
		m       model.Merchant
		b       model.Bank
		pr      model.Product
	)
	for rows.Next() {
		err := rows.Scan(
			&payment.Id, &payment.PaymentDate, &c.Id, &c.Name,
			&m.Id, &m.Name, &b.Id, &b.Name,
			&pd.Id, &pr.Id, &pr.Name, &pr.Price, &pd.Quantity,
		)
		if err != nil {
			return model.Payment{}, err
		}
		details = append(details, pd)
	}
	payment.Customer = c
	payment.Merchant = m
	payment.Bank = b
	payment.Details = details
	return payment, nil
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}
