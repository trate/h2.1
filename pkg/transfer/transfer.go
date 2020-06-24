package transfer

import (
	"errors"
	"github.com/trate/h2.1/pkg/card"
)

var (
	ErrSourceCardBalanceNotEnough = errors.New("source card balance is not enough for the operation")
	ErrSourceCardNotFound         = errors.New("source card not found")
	ErrTargetCardNotFound         = errors.New("target card not found")
)

type Service struct {
	CardSvc *card.Service
	// Commission in percent
	Commission float64
	// Minimal commission in cents
	MinCommission int64
}

func NewService(cardSvc *card.Service, commissionPercent float64, minCommission int64) *Service {
	return &Service{CardSvc: cardSvc, Commission: commissionPercent, MinCommission: minCommission}
}

func (s *Service) Card2Card(from, to string, amount int64) (int64, error) {
	var fromCard *card.Card
	var toCard *card.Card

	commission := s.Commission / 100
	withdraw := float64(amount) + commission*float64(amount)
	total := int64(withdraw)

	fromCard, ok := s.CardSvc.FindCard(from)
	if !ok {
		return total, ErrSourceCardNotFound
	}

	toCard, ok = s.CardSvc.FindCard(to)
	if !ok {
		return total, ErrTargetCardNotFound
	}

	if withdraw < float64(s.MinCommission) {
		withdraw = float64(s.MinCommission)
	}

	if fromCard != nil && withdraw >= float64(fromCard.Balance) {
		return total, ErrSourceCardBalanceNotEnough
	}

	if fromCard != nil {
		fromCard.Balance = fromCard.Balance - int64(withdraw)
	}

	if toCard != nil {
		toCard.Balance = toCard.Balance + amount
	}
	return total, nil
}
