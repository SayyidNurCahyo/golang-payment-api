package api

import (
	"merchant-payment-api/model"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MerchantController struct {
	merchantService service.MerchantService
	router *gin.Engine
}

func (s *MerchantController) createMerchantHandler(c *gin.Context){
	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err!=nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	merchant.Id = uuid.NewString()
	if err := s.merchantService.Create(merchant); err!=nil{
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, merchant)
}

func (m *MerchantController) listHandler(c *gin.Context){
	merchants, err := m.merchantService.FindAll()
	if err!=nil{
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, merchants)
}

func (m *MerchantController) getByIdHandler(c *gin.Context){
	id := c.Param("id")
	merchant, err := m.merchantService.FindById(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, merchant)
}

func NewMerchantController(merchantService service.MerchantService, engine  *gin.Engine){
	controller := MerchantController{
		merchantService: merchantService,
		router: engine,
	}
	rg := engine.Group("/api/v1")
	rg.POST("/merchants", controller.createMerchantHandler)
	rg.GET("/merchants", controller.listHandler)
	rg.GET("/merchants/:id", controller.getByIdHandler)
}