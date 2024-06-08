package repository

import (
	"database/sql"
	"merchant-payment-api/model"
)

type UserRepository interface {
	Save(payload model.UserCredential) error
	FindByUsername(username string) (model.UserCredential, error)
}

type userRepository struct {
	db *sql.DB
}

// FindByUsername implements UserRepository.
func (u *userRepository) FindByUsername(username string) (model.UserCredential, error) {
	row := u.db.QueryRow("select id, username from user_credential where username=$1", username)
	var userCredential model.UserCredential
	err := row.Scan(&userCredential.Id, &userCredential.Username)
	if err!=nil{
		return model.UserCredential{}, err
	}
	return userCredential, nil
}

// Save implements UserRepository.
func (u *userRepository) Save(payload model.UserCredential) error {
	_, err := u.db.Exec("insert into user_credential(id, username, password, is_active) values ($1, $2, $3, $4)", payload.Id, payload.Username, payload.Password, payload.IsActive)
	if err!=nil{
		return err
	}
	return nil
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
