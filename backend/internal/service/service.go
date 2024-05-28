package service

import (
	"lottery/pkg/database"
)

type User interface {
	CreateUser(number uint) uint

	GenerateToken(id uint) (string, error)

	GetNumberByToken(tokenString string) (uint, error)
}

type Lottery interface {
	GenerateNumber() (uint, error)

	CheckWinner(number uint) bool

	LotteryCompleted() map[uint]struct{}
}

type Service struct {
	User
	Lottery
}

func NewService(db database.DB, winnerPool map[uint]struct{}, sizePool uint) *Service {
	return &Service{
		User:    NewUserService(db),
		Lottery: NewLotteryService(db, winnerPool, sizePool),
	}
}
