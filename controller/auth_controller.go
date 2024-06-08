package controller

import (
	"merchant-payment-api/model"
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
	var auth model.UserCredential
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

func (a *AuthController) registerHandler(c *gin.Context){
	var auth model.UserCredential
	if err := c.ShouldBindJSON(&auth); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := a.userService.Register(auth)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully register",
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
	rg.GET("/auth/register", controller.registerHandler)
}