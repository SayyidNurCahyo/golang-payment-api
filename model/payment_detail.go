package model

type PaymentDetail struct {
	Id string
	Payment
	Product
	Price int
	Quantity int
}
