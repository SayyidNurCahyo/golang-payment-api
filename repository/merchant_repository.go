package repository

import (
	"database/sql"
	"merchant-payment-api/model"
)

type MerchantRepository interface {
	Save(merchant model.Merchant) error
	FindById(id string) (model.Merchant, error)
	FindAll() ([]model.Merchant, error)
	Update(merchant model.Merchant) error
	DeleteById(id string) error
}

type merchantRepository struct {
	db *sql.DB
}

func (m *merchantRepository) DeleteById(id string) error {
	_, errFind := m.FindById(id)
	if errFind != nil {
		return errFind
	}
	_, err := m.db.Exec("delete from merchant where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *merchantRepository) FindAll() ([]model.Merchant, error) {
	rows, err := m.db.Query("select * from merchant")
	if err != nil {
		return nil, err
	}
	var merchants []model.Merchant
	for rows.Next() {
		var merchant model.Merchant
		err := rows.Scan(&merchant.Id, &merchant.Name, &merchant.PhoneNumber, &merchant.Address)
		if err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}
	return merchants, nil
}

func (m *merchantRepository) FindById(id string) (model.Merchant, error) {
	row := m.db.QueryRow("select * from merchant where id=$1", id)
	var merchant model.Merchant
	err := row.Scan(&merchant.Id, &merchant.Name, &merchant.PhoneNumber, &merchant.Address)
	if err != nil {
		return model.Merchant{}, err
	}
	return merchant, nil
}

func (m *merchantRepository) Save(merchant model.Merchant) error {
	_, err := m.db.Exec("insert into merchant(id, name, phone_number, address) values ($1, $2, $3, $4)", merchant.Id, merchant.Name, merchant.PhoneNumber, merchant.Address)
	if err != nil {
		return err
	}
	return nil
}

func (m *merchantRepository) Update(merchant model.Merchant) error {
	_, errFind := m.FindById(merchant.Id)
	if errFind != nil {
		return errFind
	}
	_, err := m.db.Exec("update merchant set name=$1, phone_number=$2, address=$3 where id=$4", merchant.Name, merchant.PhoneNumber, merchant.Address, merchant.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewMerchantRepo(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
