package card

import "strings"

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Transactions []Transaction
}

type Transaction struct {
	Id     int64
	Amount int64
	Date   int64
	MCC    string
	Status string
}

type Service struct {
	BankName string
	Cards    []*Card
}

// define the prefix for the card number that belongs to our bank
const ourCardsPrefix = "510621"

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

// check if the card's number belongs to our bank
func  CardIsOurs(number string) bool {
	if strings.HasPrefix(number, ourCardsPrefix) {
		return true
	}
	return false
}

func (s *Service) FindCard(number string) (*Card, bool) {
	if !CardIsOurs(number) {
		return nil, false
	}
	for _, v := range s.Cards {
		if v.Number == number {
			return v, true
		}
	}
	return nil, false
}


func (s *Service) Add(cards ...*Card) {
	s.Cards = append(s.Cards, cards...)
}
