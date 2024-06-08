package repository

import (
	"database/sql"
	"merchant-payment-api/model"
)

type BankRepository interface {
	Save(bank model.Bank) error
	FindById(id string) (model.Bank, error) 
	FindAll() ([]model.Bank, error)
	Update(bank model.Bank) error
	DeleteById(id string) error
}

type bankRepository struct {
	db *sql.DB
}

func (b *bankRepository) DeleteById(id string) error {
	bank, errFind := b.FindById(id)
	if errFind != nil {
		return errFind
	}
	_, err := b.db.Exec("update user_credential set is_active=false where id=$1", bank.UserCredential.Id)
	if err != nil {
		return err
	}
	return nil
}

func (b *bankRepository) FindAll() ([]model.Bank, error) {
	rows, err := b.db.Query("select b.id, b.name, uc.id, uc.username from bank as b join user_credential as uc on uc.id = b.user_id where uc.is_active=true")
	if err != nil {
		return nil, err
	}
	var banks []model.Bank
	for rows.Next() {
		var user model.UserCredential
		var bank model.Bank
		err := rows.Scan(&bank.Id, &bank.Name, &user.Id, &user.Username)
		if err != nil {
			return nil, err
		}
		bank.UserCredential = user
		banks = append(banks, bank)
	}
	return banks, nil
}

func (b *bankRepository) FindById(id string) (model.Bank, error) {
	row := b.db.QueryRow("select b.id, b.name, uc.id, uc.username from bank as b join user_credential as uc on uc.id=b.user_id where b.id=$1 and uc.is_active=true", id)
	var user model.UserCredential
	var bank model.Bank
	err := row.Scan(&bank.Id, &bank.Name, &user.Id, &user.Username)
	if err != nil {
		return model.Bank{}, err
	}
	bank.UserCredential = user
	return bank, nil
}

func (b *bankRepository) Save(bank model.Bank) error {
	_, err := b.db.Exec("insert into bank(id, name, user_id) values ($1, $2, $3)", bank.Id, bank.Name, bank.UserCredential.Id)
	if err != nil {
		return err
	}
	return nil
}

func (b *bankRepository) Update(bank model.Bank) error {
	_, errFind := b.FindById(bank.Id)
	if errFind != nil {
		return errFind
	}
	_, err := b.db.Exec("update bank set name=$1 where id=$2", bank.Name, bank.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewBankRepo(db *sql.DB) BankRepository {
	return &bankRepository{db: db}
}
