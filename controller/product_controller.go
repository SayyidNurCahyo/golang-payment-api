package controller

import (
	"merchant-payment-api/dto"
	"merchant-payment-api/middleware"
	"merchant-payment-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.ProductService
	router         *gin.Engine
}

func (p *ProductController) getAllHandler(c *gin.Context) {
	products, err := p.productService.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get all product",
		"data":    products,
	})
}

func (p *ProductController) getByNameHandler(c *gin.Context) {
	name := c.Param("name")
	products, err := p.productService.FindByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get product",
		"data":    products,
	})
}

func (p *ProductController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	product, err := p.productService.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get product",
		"data":    product,
	})
}

func (p *ProductController) updateHandler(c *gin.Context) {
	var product dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := p.productService.Update(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully update product",
		"data":    product,
	})
}

func (p *ProductController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.productService.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete product",
	})
}

func (p *ProductController) createHandler(c *gin.Context) {
	var request dto.SaveProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := p.productService.Create(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully create new product",
	})
}

func NewProductController(productService service.ProductService, engine *gin.Engine) {
	controller := ProductController{
		productService: productService,
		router:         engine,
	}
	rg := engine.Group("/api/v1", middleware.AuthMiddleware())
	rg.POST("/products", controller.createHandler)
	rg.GET("/products", controller.getAllHandler)
	rg.GET("/products/:id", controller.getByIdHandler)
	rg.GET("/products/name/:name", controller.getByNameHandler)
	rg.PUT("/products", controller.updateHandler)
	rg.DELETE("/products/:id", controller.deleteHandler)
}
