package security

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	jwt.RegisteredClaims
	Username string   `json:"username"`
	Role     string   `json:"role,omitempty"`
	Services []string `json:"services,omitempty"`
}
