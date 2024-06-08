package controller

import (
	"merchant-payment-api/dto"
	"merchant-payment-api/middleware"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	bankService service.BankService
	router      *gin.Engine
}

func (b *BankController) getAllHandler(c *gin.Context) {
	banks, err := b.bankService.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get all bank",
		"data":    banks,
	})
}

func (b *BankController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	bank, err := b.bankService.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get bank",
		"data":    bank,
	})
}

func (b *BankController) updateHandler(c *gin.Context) {
	var bank dto.UpdateBankRequest
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := b.bankService.Update(bank)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully update bank",
		"data":    bank,
	})
}

func (b *BankController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := b.bankService.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete bank",
	})
}

func NewBankController(bankService service.BankService, engine *gin.Engine) {
	controller := BankController{
		bankService: bankService,
		router:      engine,
	}
	rg := engine.Group("/api/v1", middleware.AuthMiddleware())
	rg.GET("/banks", controller.getAllHandler)
	rg.GET("/banks/:id", controller.getByIdHandler)
	rg.PUT("/banks", controller.updateHandler)
	rg.DELETE("/banks/:id", controller.deleteHandler)
}
