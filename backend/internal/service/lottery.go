package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"lottery/pkg/database"
)

type LotteryService struct {
	winnerPool map[uint]struct{}
	minPool    map[uint]struct{}
	usedPool   map[uint]struct{}
	db         database.DB
	rand       *rand.Rand
	completed  bool
	wait       chan struct{}
}

func NewLotteryService(db database.DB, winnerPool map[uint]struct{}, sizePool uint) *LotteryService {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	minPool := make(map[uint]struct{})
	for k, v := range winnerPool {
		minPool[k] = v
	}

	for len(minPool) < int(sizePool) {
		minPool[uint(r.Intn(500))] = struct{}{}
	}

	lt := &LotteryService{
		winnerPool: winnerPool,
		minPool:    minPool,
		usedPool:   make(map[uint]struct{}),
		db:         db,
		rand:       r,
		completed:  false,
		wait:       make(chan struct{}),
	}
	go func() {
		lt.wait <- struct{}{}
	}()
	return lt
}

func (s *LotteryService) GenerateNumber() (uint, error) {
	if s.completed {
		return 0, errors.New("completed")
	}
	fmt.Println(s.wait)
	<-s.wait
	value, err := s.getNumber()
	go func() {
		s.wait <- struct{}{}
	}()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (s *LotteryService) getNumber() (uint, error) {
	for k := range s.minPool {
		delete(s.minPool, k)
		s.usedPool[k] = struct{}{}
		return k, nil
	}
	return s.generateUnused()
}

func (s *LotteryService) generateUnused() (uint, error) {
	for i := 0; i < 50; i++ {
		value := uint(s.rand.Intn(500))
		if _, exists := s.usedPool[value]; !exists {
			s.usedPool[value] = struct{}{}
			return value, nil
		}
	}
	return 0, errors.New("failed generate new number")
}

func (s *LotteryService) CheckWinner(number uint) bool {
	if !s.completed {
		return false
	}
	_, exists := s.winnerPool[number]
	return exists
}

func (s *LotteryService) LotteryCompleted() map[uint]struct{} {
	s.completed = true
	return s.winnerPool
}
