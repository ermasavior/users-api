package jwt

import "github.com/golang-jwt/jwt"

type JWTServiceInterface interface {
	GenerateToken(jwt.Claims) (string, error)
	ValidateToken(string) (*jwt.Token, error)
}
