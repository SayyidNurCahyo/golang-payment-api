package controller

import (
	"fmt"
	"merchant-payment-api/config"
	"merchant-payment-api/repository"
	"merchant-payment-api/service"
)

type Console struct {
	merchantService service.MerchantService
}

func NewConsole() *Console{
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	con, err := config.NewDbConnection(cfg)
	if err!=nil {
		fmt.Println(err)
	}
	db := con.Conn()
	merchantRepo := repository.NewMerchantRepo(db)
	merchantService := service.NewMerchantService(merchantRepo)
	return &Console{merchantService: merchantService}
}