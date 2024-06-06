package main

import (
	"fmt"
	"merchant-payment-api/config"
	// "merchant-payment-api/model"
	"merchant-payment-api/repository"

	_ "github.com/lib/pq"
)

type Customer struct{
	Id string
	Name string
	PhoneNumber string
	Address string
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	con, err := config.NewDbConnection(cfg)
	if err!=nil {
		fmt.Println(err)
	}
	db := con.Conn()
	productRepo := repository.NewProductRepo(db)
	// err = productRepo.DeleteById(2)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// err = productRepo.Update(model.Product{Id: 2, Name: "baju", Price: 50000})
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	product, err := productRepo.FindAll()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(product)
	// productRepo.Save(model.Product{
	// 	Name: "laptop",
	// 	Price: 10000000,
	// })
}