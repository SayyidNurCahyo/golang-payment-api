package controller

import (
	"merchant-payment-api/middleware"
	"merchant-payment-api/model"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentService service.PaymentService
	router *gin.Engine
}

func (s *PaymentController) createHandler(c *gin.Context){
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := s.paymentService.CreatePayment(payment)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully update payment",
		"data": response,
	})
}

// func (m *PaymentController) getAllHandler(c *gin.Context){
// 	payments, err := m.paymentService.FindAll()
// 	if err!=nil{
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "successfully update payment",
// 		"data": payments,
// 	})
// }

// func (m *PaymentController) getByIdHandler(c *gin.Context){
// 	id := c.Param("id")
// 	payment, err := m.paymentService.FindById(id)
// 	if err!=nil{
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "successfully update payment",
// 		"data": payment,
// 	})
// }

func NewPaymentController(paymentService service.PaymentService, engine  *gin.Engine){
	controller := PaymentController{
		paymentService: paymentService,
		router: engine,
	}
	rg := engine.Group("/api/v1", middleware.AuthMiddleware())
	rg.POST("/payments", controller.createHandler)
	// rg.GET("/payments", controller.getAllHandler)
	// rg.GET("/payments/:id", controller.getByIdHandler)
}