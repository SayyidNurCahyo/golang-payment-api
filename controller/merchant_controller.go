package controller

import (
	"merchant-payment-api/dto"
	"merchant-payment-api/middleware"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantService service.MerchantService
	router *gin.Engine
}

func (m *MerchantController) getAllHandler(c *gin.Context){
	merchants, err := m.merchantService.FindAll()
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get all merchant",
		"data": merchants,
	})
}

func (m *MerchantController) getByIdHandler(c *gin.Context){
	id := c.Param("id")
	merchant, err := m.merchantService.FindById(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get merchant",
		"data": merchant,
	})
}

func (m *MerchantController) updateHandler(c *gin.Context){
	var merchant dto.UpdateMerchantRequest
	if err := c.ShouldBindJSON(&merchant); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := m.merchantService.Update(merchant)
	if err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully update merchant",
		"data": merchant,
	})
}

func (m *MerchantController) deleteHandler(c *gin.Context){
	id := c.Param("id")
	if err := m.merchantService.Delete(id); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete merchant",
	})
}

func NewMerchantController(merchantService service.MerchantService, engine  *gin.Engine){
	controller := MerchantController{
		merchantService: merchantService,
		router: engine,
	}
	rg := engine.Group("/api/v1", middleware.AuthMiddleware())
	rg.GET("/merchants", controller.getAllHandler)
	rg.GET("/merchants/:id", controller.getByIdHandler)
	rg.PUT("/merchants", controller.updateHandler)
	rg.DELETE("/merchants/:id", controller.deleteHandler)
}