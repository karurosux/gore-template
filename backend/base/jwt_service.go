package base

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
)

type JwtService struct{}

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewJwtService(i *do.Injector) (*JwtService, error) {
	return &JwtService{}, nil
}

func (js *JwtService) GenerateToken(email string) (string, error) {
	claims := &jwtCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (js *JwtService) ValidateToken(token string) (*jwt.Token, *jwtCustomClaims, error) {
	claims := &jwtCustomClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return parsedToken, claims, err
}
