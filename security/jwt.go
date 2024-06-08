package security

import (
	"fmt"
	"merchant-payment-api/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user model.UserCredential) (string, error){
	claims := &JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(60)*time.Minute)),
		},
		Username: user.Username,
		// Role: "",
		// Services: []string{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err!=nil{
		return "", fmt.Errorf("failed to create jwt token: %v", err.Error())
	}
	return tokenString, nil
}