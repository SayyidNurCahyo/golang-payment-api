package middleware

import (
	"merchant-payment-api/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct{
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authHeader authHeader
		if err := c.ShouldBindHeader(&authHeader); err!=nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorizaed",
			})
			return
		}

		// catch token in header
		token := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", 1)
		if token == ""{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorizaed",
			})
			return
		}

		// verifikasi token
		claims, err := security.VerifyToken(token)
		if err!=nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorizaed",
			})
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}