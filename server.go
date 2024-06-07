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
	paymentService  service.PaymentService
	customerService service.CustomerService
	bankService     service.BankService
	engine          *gin.Engine
	host            string
	log             *logrus.Logger
}

func (s *Server) Run() {
	s.initMiddleware()
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddleware() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
}

// init controller
func (s *Server) initController() {
	controller.NewMerchantController(s.merchantService, s.engine)
	controller.NewPaymentController(s.paymentService, s.engine)
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
	customerRepo := repository.NewCustomerRepo(db)
	customerService := service.NewCustomerService(customerRepo)
	bankRepo := repository.NewBankRepo(db)
	bankService := service.NewBankService(bankRepo)
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo, customerService, merchantService, bankService, productService)

	host := fmt.Sprintf(":%s", cfg.ApiPort)
	log := logrus.New()
	engine := gin.Default()
	server := Server{
		merchantService: merchantService,
		productService:  productService,
		bankService:     bankService,
		customerService: customerService,
		paymentService:  paymentService,
		engine:          engine,
		host:            host,
		log:             log,
	}

	return &server
}
