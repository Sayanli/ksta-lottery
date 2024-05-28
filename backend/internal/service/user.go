package service

import (
	"errors"
	"lottery/internal/config"
	"lottery/pkg/database"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserService struct {
	db database.DB
}

func NewUserService(db database.DB) *UserService {
	return &UserService{
		db: db,
	}
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user_id"`
}

func (s *UserService) CreateUser(number uint) uint {
	return s.db.Insert(number)
}

func (s *UserService) GetNumberByToken(tokenString string) (uint, error) {
	token, err := stringTokenToJwt(tokenString)
	if err != nil {
		return 0, errors.New("token is invalid")
	}
	claims := token.Claims.(jwt.MapClaims)
	uid := uint(claims["user_id"].(float64))
	return s.db.GetValueByKey(uid)
}

func (s *UserService) GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		id,
	})

	return token.SignedString([]byte(config.EnvConfig.SigningKeyJwt))
}

func stringTokenToJwt(tokenString string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.EnvConfig.SigningKeyJwt), nil
	})
}
