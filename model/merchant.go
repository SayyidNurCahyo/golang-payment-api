package model

type Merchant struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	Address     string `json:"address"`
}
