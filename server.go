package main

import (
	"fmt"
	"merchant-payment-api/config"
	"merchant-payment-api/model"
	"merchant-payment-api/repository"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Server struct {
	merchantService service.MerchantService
	productService  service.ProductService
	engine          *gin.Engine
}

func (s *Server) Run() {
	err := s.engine.Run()
	if err!=nil {
		panic(err)
	}
}

func (s *Server) createMerchantHandler(c *gin.Context){
	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err!=nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	merchant.Id = uuid.NewString()
}

func NewServer() *Server{
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
	productRepo := repository.NewProductRepo(db)
	productService := service.NewProductService(productRepo, merchantService)
	engine := gin.Engine{}

	return &Server{
		merchantService: merchantService,
		productService: productService,
		engine: &engine,
	}
}