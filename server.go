package main

import (
	"fmt"
	"merchant-payment-api/config"
	"merchant-payment-api/controller"
	"merchant-payment-api/middleware"
	"merchant-payment-api/repository"
	"merchant-payment-api/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	merchantService service.MerchantService
	productService  service.ProductService
	engine          *gin.Engine
	host            string
	log             *logrus.Logger
}

func (s *Server) Run() {
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

// init controller
func (s *Server) initController() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	controller.NewMerchantController(s.merchantService, s.engine)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	con, err := config.NewDbConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db := con.Conn()
	merchantRepo := repository.NewMerchantRepo(db)
	merchantService := service.NewMerchantService(merchantRepo)
	productRepo := repository.NewProductRepo(db)
	productService := service.NewProductService(productRepo, merchantService)

	host := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	log := logrus.New()
	engine := gin.Default()
	server := Server{
		merchantService: merchantService,
		productService:  productService,
		engine:          engine,
		host:            host,
		log:             log,
	}
	server.initController()

	return &server
}
