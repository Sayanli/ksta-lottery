package handler

import (
	"encoding/json"
)

func (h *Handler) GetNumberByToken(body []byte) (uint, error) {
	type Token struct {
		Value string `json:"token"`
	}

	var token Token
	err := json.Unmarshal(body, &token)
	if err != nil {
		return 0, err
	}

	return h.services.User.GetNumberByToken(token.Value)
}

func (h *Handler) GenerateNumberAndToken() (uint, string, error) {
	number, err := h.services.Lottery.GenerateNumber()
	if err != nil {
		return 0, "", err
	}

	uid := h.services.User.CreateUser(number)
	token, err := h.services.User.GenerateToken(uid)
	if err != nil {
		return 0, "", err
	}

	return number, token, nil
}

func (h *Handler) LotteryCompleted() map[uint]struct{} {
	return h.services.Lottery.LotteryCompleted()
}

func (h *Handler) CheckWinner(number uint) bool {
	return h.services.Lottery.CheckWinner(number)
}
