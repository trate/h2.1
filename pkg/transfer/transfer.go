package transfer

import (
	"fmt"
	"github.com/trate/h2.1/pkg/card"
)

type Service struct {
	CardSvc *card.Service
	// Commission in percent
	Commission float64
	// Minimal commission in cents
	MinCommission int64
}

func NewService(cardSvc *card.Service, commission float64, minCommission int64) *Service {
	return &Service{ CardSvc: cardSvc, Commission: commission, MinCommission: minCommission}
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {
	var fromCard *card.Card
	var toCard *card.Card

	fromCard = s.CardSvc.FindCard(from)
	toCard = s.CardSvc.FindCard(to)
	fmt.Println(fromCard)
	fmt.Println(toCard)

	commission := s.Commission / 100
	withdraw := float64(amount) + commission * float64(amount)
	total = int64(withdraw)
	if withdraw < float64(s.MinCommission) {
		withdraw = float64(s.MinCommission)
	}

	if fromCard == nil && toCard == nil {
		ok = true
		return
	}

	if fromCard != nil && withdraw >= float64(fromCard.Balance) {
		ok = false
		return
	}

	if fromCard != nil {
		fromCard.Balance = fromCard.Balance - int64(withdraw)
		ok = true
	}

	if  toCard != nil {
		toCard.Balance = toCard.Balance + amount
		ok = true
	}
	return
}


