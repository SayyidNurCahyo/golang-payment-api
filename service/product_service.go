package service

import "merchant-payment-api/model"

type ProductService interface {
	Create(model.Product) error
	Delete(id int) error
	FindAll() ([]model.Product, error)
	FindById(id int) (model.Product, )
}