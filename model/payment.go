package model

import "time"

type Payment struct {
	Id          string
	PaymentDate time.Time
	Merchant
	Bank
	Details []PaymentDetail
}
