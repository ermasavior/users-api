package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

var ErrInvalidToken = errors.New("invalid JWT token")

type JWTService struct {
	secret string
}

func NewJWTService() *JWTService {
	return &JWTService{
		//TODO: store secret code into env variable
		//os.Getenv("JWT_SECRET")
		secret: "secret",
	}
}

func (s *JWTService) GenerateToken(claims jwt.Claims) (string, error) {
	secretKey := []byte(s.secret)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(secretKey)
}

func (s *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken

		}
		return []byte(s.secret), nil
	})
}
