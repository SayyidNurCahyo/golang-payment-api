package repository

import (
	"database/sql"
	"merchant-payment-api/model"
)

type ProductRepository interface {
	Save(product model.Product) error
	FindById(id string) (model.Product, error)
	FindByName(name string) ([]model.Product, error)
	FindAll() ([]model.Product, error)
	Update(product model.Product) error
	DeleteById(id string) error
}

type productRepository struct {
	db *sql.DB
}

func (p *productRepository) DeleteById(id string) error {
	product, errFind := p.FindById(id)
	if errFind != nil {
		return errFind
	}
	_, err := p.db.Exec("update product set is_available=false where id=$1", product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) FindAll() ([]model.Product, error) {
	rows, err := p.db.Query("select p.id, m.id, m.name, p.name, p.price from product as p join merchant as m on m.id = p.merchant_id where p.is_available=true")
	if err != nil {
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		var merchant model.Merchant
		var product model.Product
		err := rows.Scan(&product.Id, &merchant.Id, &merchant.Name, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		product.Merchant = merchant
		products = append(products, product)
	}
	return products, nil
}

func (p *productRepository) FindById(id string) (model.Product, error) {
	row := p.db.QueryRow("select p.id, m.id, m.name, p.name, p.price from product as p join merchant as m on m.id = p.merchant_id where p.id=$1 and p.is_available=true", id)
	var merchant model.Merchant
	var product model.Product
	err := row.Scan(&product.Id, &merchant.Id, &merchant.Name, &product.Name, &product.Price)
	if err != nil {
		return model.Product{}, err
	}
	product.Merchant = merchant
	return product, nil
}

func (p *productRepository) FindByName(name string) ([]model.Product, error) {
	name = "%"+name+"%"
	rows, err := p.db.Query("select p.id, m.id, m.name, p.name, p.price from product as p join merchant as m on m.id = p.merchant_id where p.is_available=true and p.name ilike $1", name)
	if err != nil {
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		var merchant model.Merchant
		var product model.Product
		err := rows.Scan(&product.Id, &merchant.Id, &merchant.Name, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		product.Merchant = merchant
		products = append(products, product)
	}
	return products, nil
}

func (p *productRepository) Save(product model.Product) error {
	_, err := p.db.Exec("insert into product(id, merchant_id, name, price, is_available) values ($1, $2, $3, $4, $5)", product.Id, product.Merchant.Id, product.Name, product.Price, true)
	if err!=nil{
		return err
	}
	return nil
}

func (p *productRepository) Update(product model.Product) error {
	_, errFind := p.FindById(product.Id)
	if errFind!=nil{
		return errFind
	}
	_, err := p.db.Exec("update product set name=$1, price=$2 where id=$3", product.Name, product.Price, product.Id)
	if err!=nil{
		return err
	}
	return nil
}

func NewProductRepo(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
