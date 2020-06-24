package card

import (
	"log"
	"strconv"
	"strings"
)

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
func  IsOurs(number string) bool {
	if strings.HasPrefix(number, ourCardsPrefix) {
		return true
	}
	return false
}

//Check the card number in accordance with the Luhn's algorithm
func IsValid(number string)  bool {
	var sum int

	number = strings.ReplaceAll(number, " ", "")
	numbers := strings.Split(number, "")
	digits := make([]int, len(numbers))

	for i, v := range numbers {
		d, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("Error converting to int: ", err)
		}
		digits[i] = d
	}

	for i,v := range digits {
		if (len(digits) - i) % 2 == 0 {
			v = v * 2;

			if v > 9 {
				v = v - 9
			}
		}
		sum += v
	}
	if sum % 10 == 0 {
		return true
	}
	return false
}

func (s *Service) FindCard(number string) (*Card, bool) {
	if !IsOurs(number) {
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
