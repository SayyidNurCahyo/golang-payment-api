package repository

import (
	"database/sql"
	"merchant-payment-api/model"
)

type BankRepository interface {
	Save(bank model.Bank) error
	FindById(id int) (model.Bank, error)
	FindAll() ([]model.Bank, error)
	Update(bank model.Bank) error
	DeleteById(id int) error
}

type bankRepository struct {
	db *sql.DB
}

func (m *bankRepository) DeleteById(id int) error {
	_, errFind := m.FindById(id)
	if errFind != nil {
		return errFind
	}
	_, err := m.db.Exec("delete from bank where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *bankRepository) FindAll() ([]model.Bank, error) {
	rows, err := m.db.Query("select * from bank")
	if err != nil {
		return nil, err
	}
	var banks []model.Bank
	for rows.Next() {
		var bank model.Bank
		err := rows.Scan(&bank.Id, &bank.Name)
		if err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}
	return banks, nil
}

func (m *bankRepository) FindById(id int) (model.Bank, error) {
	row := m.db.QueryRow("select * from bank where id=$1", id)
	var bank model.Bank
	err := row.Scan(&bank.Id, &bank.Name)
	if err != nil {
		return model.Bank{}, err
	}
	return bank, nil
}

func (m *bankRepository) Save(bank model.Bank) error {
	_, err := m.db.Exec("insert into bank(name) values ($1)", bank.Name)
	if err != nil {
		return err
	}
	return nil
}

func (m *bankRepository) Update(bank model.Bank) error {
	_, errFind := m.FindById(bank.Id)
	if errFind != nil {
		return errFind
	}
	_, err := m.db.Exec("update bank set name=$1 where id=$4", bank.Name, bank.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewBankRepo(db *sql.DB) BankRepository {
	return &bankRepository{db: db}
}
