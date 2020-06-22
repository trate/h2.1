package card

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

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) FindCard(number string) *Card {
	for _, v := range s.Cards {
		if v.Number == number {
			return v
		}
	}
	return nil
}

func (s *Service) Add(cards ...*Card) {
	s.Cards = append(s.Cards, cards...)
}
