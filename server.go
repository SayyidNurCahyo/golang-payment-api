package main

import (
	"fmt"
	"merchant-payment-api/config"
	"merchant-payment-api/controller/api"
	"merchant-payment-api/repository"
	"merchant-payment-api/service"
	"github.com/gin-gonic/gin"
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

// init controller
func (s *Server) initController(){
	api.NewMerchantController(s.merchantService, s.engine)
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

	engine := gin.Default()
	server := Server{
		merchantService: merchantService,
		productService: productService,
		engine: engine,
	}
	server.initController()

	return &server
}