package repository

import (
	"database/sql"
	"merchant-payment-api/model"
)

type CustomerRepository interface {
	Save(customer model.Customer) error
	FindById(id string) (model.Customer, error)
	FindAll() ([]model.Customer, error)
	Update(customer model.Customer) error
	DeleteById(id string) error
}

type customerRepository struct {
	db *sql.DB
}

func (m *customerRepository) DeleteById(id string) error {
	_, errFind := m.FindById(id)
	if errFind != nil {
		return errFind
	}
	_, err := m.db.Exec("delete from customer where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *customerRepository) FindAll() ([]model.Customer, error) {
	rows, err := m.db.Query("select * from customer")
	if err != nil {
		return nil, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (m *customerRepository) FindById(id string) (model.Customer, error) {
	row := m.db.QueryRow("select * from customer where id=$1", id)
	var customer model.Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (m *customerRepository) Save(customer model.Customer) error {
	_, err := m.db.Exec("insert into customer(id, name, phone_number) values ($1, $2, $3)", customer.Id, customer.Name, customer.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (m *customerRepository) Update(customer model.Customer) error {
	_, errFind := m.FindById(customer.Id)
	if errFind != nil {
		return errFind
	}
	_, err := m.db.Exec("update customer set name=$1, phone_number=$2 where id=$3", customer.Name, customer.PhoneNumber, customer.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
