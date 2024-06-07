package model

import "time"

type Payment struct {
	Id          string
	PaymentDate time.Time
	Customer
	Merchant
	Bank
	Details []PaymentDetail
}
