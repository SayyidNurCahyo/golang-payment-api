package controller

import (
	"merchant-payment-api/dto"
	"merchant-payment-api/middleware"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService service.CustomerService
	router *gin.Engine
}

func (m *CustomerController) getAllHandler(c *gin.Context){
	customers, err := m.customerService.FindAll()
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get all customer",
		"data": customers,
	})
}

func (m *CustomerController) getByIdHandler(c *gin.Context){
	id := c.Param("id")
	customer, err := m.customerService.FindById(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get customer",
		"data": customer,
	})
}

func (m *CustomerController) updateHandler(c *gin.Context){
	var customer dto.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&customer); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := m.customerService.Update(customer)
	if err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully update customer",
		"data": customer,
	})
}

func (m *CustomerController) deleteHandler(c *gin.Context){
	id := c.Param("id")
	if err := m.customerService.Delete(id); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete customer",
	})
}

func NewCustomerController(customerService service.CustomerService, engine  *gin.Engine){
	controller := CustomerController{
		customerService: customerService,
		router: engine,
	}
	rg := engine.Group("/api/v1", middleware.AuthMiddleware())
	rg.GET("/customers", controller.getAllHandler)
	rg.GET("/customers/:id", controller.getByIdHandler)
	rg.PUT("/customers", controller.updateHandler)
	rg.DELETE("/customers/:id", controller.deleteHandler)
}