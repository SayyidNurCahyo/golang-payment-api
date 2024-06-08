package controller

import (
	"merchant-payment-api/dto"
	"merchant-payment-api/middleware"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentService service.PaymentService
	router         *gin.Engine
}

func (p *PaymentController) createHandler(c *gin.Context) {
	var payment dto.PaymentRequest
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := p.paymentService.CreatePayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully create payment",
	})
}

func (p *PaymentController) getAllHandler(c *gin.Context) {
	payments, err := p.paymentService.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get all payment",
		"data":    payments,
	})
}

func (p *PaymentController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	payment, err := p.paymentService.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get payment",
		"data":    payment,
	})
}

func NewPaymentController(paymentService service.PaymentService, engine *gin.Engine) {
	controller := PaymentController{
		paymentService: paymentService,
		router:         engine,
	}
	rg := engine.Group("/api/v1", middleware.AuthMiddleware())
	rg.POST("/payments", controller.createHandler)
	rg.GET("/payments", controller.getAllHandler)
	rg.GET("/payments/:id", controller.getByIdHandler)
}
