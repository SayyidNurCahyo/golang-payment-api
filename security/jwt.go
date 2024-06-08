package security

import (
	"fmt"
	"merchant-payment-api/config"
	"merchant-payment-api/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user model.UserCredential) (string, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return "", err
	}
	claims := &JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.ExpirationToken) * time.Minute)),
		},
		Username: user.Username,
		// Role: "",
		// Services: []string{},
	}

	token := jwt.NewWithClaims(cfg.JWTSigningMethod, claims)
	tokenString, err := token.SignedString(cfg.JWTSignatureKey)
	if err != nil {
		return "", fmt.Errorf("failed to create jwt token: %v", err.Error())
	}
	return tokenString, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	tokenVerified, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != cfg.JWTSigningMethod {
			return nil, fmt.Errorf("invalid token signin method")
		}
		return cfg.JWTSignatureKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenVerified.Claims.(jwt.MapClaims)
	if !ok || !tokenVerified.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
