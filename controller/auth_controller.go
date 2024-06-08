package controller

import (
	"merchant-payment-api/dto"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService service.UserService
	authService service.AuthService
	router *gin.Engine
}

func (a *AuthController) loginHandler(c *gin.Context){
	var auth dto.LoginRequest
	if err := c.ShouldBindJSON(&auth); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := a.authService.Login(auth)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully login",
		"data": response,
	})
}

func (a *AuthController) registerBankHandler(c *gin.Context){
	var auth dto.SaveBankRequest
	if err := c.ShouldBindJSON(&auth); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userService.RegisterBank(auth)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully register new bank",
	})
}

func (a *AuthController) registerCustomerHandler(c *gin.Context){
	var auth dto.SaveCustomerRequest
	if err := c.ShouldBindJSON(&auth); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userService.RegisterCustomer(auth)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully register new customer",
	})
}

func NewAuthController(userService service.UserService, authService service.AuthService, engine  *gin.Engine){
	controller := AuthController{
		userService: userService,
		authService: authService,
		router: engine,
	}
	rg := engine.Group("/api/v1")
	rg.POST("/auth/login", controller.loginHandler)
	rg.POST("/auth/register/bank", controller.registerBankHandler)
	rg.POST("/auth/register/customer", controller.registerCustomerHandler)
}