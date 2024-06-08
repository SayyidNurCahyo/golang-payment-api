package dto

import "time"

type PaymentRequest struct {
	CustomerId           string                 `json:"customerId"`
	MerchantId           string                 `json:"merchantId"`
	BankId               string                 `json:"bankId"`
	PaymentDetailRequest []PaymentDetailRequest `json:"details"`
}

type PaymentDetailRequest struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type PaymentResponse struct {
	Id                    string                  `json:"id"`
	PaymentDate           time.Time               `json:"date"`
	CustomerId            string                  `json:"customerId"`
	CustomerName          string                  `json:"customerName"`
	MerchantId            string                  `json:"merchantId"`
	MerchantName          string                  `json:"merchantName"`
	BankId                string                  `json:"bankId"`
	BankName              string                  `json:"bankName"`
	PaymentDetailResponse []PaymentDetailResponse `json:"details"`
}

type PaymentDetailResponse struct {
	DetailId     string `json:"detailId"`
	ProductId    string `json:"productId"`
	ProductName  string `json:"productName"`
	ProductPrice int    `json:"productPrice"`
	Quantity     int    `json:"quantity"`
}
