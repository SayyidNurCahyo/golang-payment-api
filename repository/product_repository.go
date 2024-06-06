package repository

import (
	"database/sql"
	"merchant-payment-api/model"
)

type ProductRepository interface {
	Save(product model.Product) error // insert
	FindById(id int) (model.Product, error)
	FindAll() ([]model.Product, error)
	Update(product model.Product) error
	DeleteById(id int) error
}

type productRepository struct {
	db *sql.DB
}

func (p *productRepository) DeleteById(id int) error {
	_, errFind := p.FindById(id)
	if errFind!=nil{
		return errFind
	}
	_, err := p.db.Exec("delete from product where id=$1", id)
	if err!= nil{
		return err
	}
	return nil
}

func (p *productRepository) FindAll() ([]model.Product, error) {
	rows, err := p.db.Query("select * from product")
	if err!=nil{
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		if err!=nil{
			return nil,err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *productRepository) FindById(id int) (model.Product, error) {
	row := p.db.QueryRow("select * from product where id=$1", id)
	var product model.Product
	// & buat menimpa data di product (di set)
	err := row.Scan(&product.Id, &product.Name, &product.Price)
	if err!= nil{
		return model.Product{}, err
	}
	return product, nil
}

func (p *productRepository) Save(product model.Product) error {
	_, err := p.db.Exec("insert into product(name, price) values ($1, $2)", product.Name, product.Price)
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
