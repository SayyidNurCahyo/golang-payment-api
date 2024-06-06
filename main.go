package main

import (
	"fmt"
	"merchant-payment-api/config"
	// "merchant-payment-api/model"
	// "merchant-payment-api/repository"
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
	con.Conn()
}