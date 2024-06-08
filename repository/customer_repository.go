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

func (c *customerRepository) DeleteById(id string) error {
	customer, errFind := c.FindById(id)
	if errFind != nil {
		return errFind
	}
	_, err := c.db.Exec("update user_credential set is_active=false where id=$1", customer.UserCredential.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) FindAll() ([]model.Customer, error) {
	rows, err := c.db.Query("select c.id, c.name, c.phone_number, uc.id, uc.username from customer as c join user_credential as uc on uc.id = c.user_id where uc.is_active=true")
	if err != nil {
		return nil, err
	}
	var customers []model.Customer
	for rows.Next() {
		var user model.UserCredential
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &user.Id, &user.Username)
		if err != nil {
			return nil, err
		}
		customer.UserCredential = user
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *customerRepository) FindById(id string) (model.Customer, error) {
	row := c.db.QueryRow("select c.id, c.name, c.phone_number, uc.id, uc.username from customer as c join user_credential as uc on uc.id= c.user_id where c.id=$1 and uc.is_active=true", id)
	var user model.UserCredential
	var customer model.Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &user.Id, &user.Username)
	if err != nil {
		return model.Customer{}, err
	}
	customer.UserCredential = user
	return customer, nil
}

func (c *customerRepository) Save(customer model.Customer) error {
	_, err := c.db.Exec("insert into customer(id, name, phone_number, user_id) values ($1, $2, $3, $4)", customer.Id, customer.Name, customer.PhoneNumber, customer.UserCredential.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Update(customer model.Customer) error {
	_, errFind := c.FindById(customer.Id)
	if errFind != nil {
		return errFind
	}
	_, err := c.db.Exec("update customer set name=$1, phone_number=$2 where id=$3", customer.Name, customer.PhoneNumber, customer.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
