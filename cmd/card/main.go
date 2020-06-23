package main

import (
	"fmt"
	"github.com/trate/h2.1/pkg/card"
	"github.com/trate/h2.1/pkg/transfer"
)

func main() {

	master := &card.Card{
		Id:           1,
		Issuer:       "MasterCard",
		Balance:      65_000,
		Currency:     "RUB",
		Number:       "5177827685644009",
		Transactions: nil,
	}
	visa := &card.Card{
		Id:           2,
		Issuer:       "Visa",
		Balance:      64_000,
		Currency:     "RUB",
		Number:       "4716742265786594",
		Transactions: nil,
	}

	tinkoff := card.NewService("Tinkoff")
	tinkoffTransfers := transfer.NewService(tinkoff, 0.5, 10_00)

	tinkoff.Cards = append(tinkoff.Cards, master)
	tinkoff.Cards = append(tinkoff.Cards, visa)


	total, ok := tinkoffTransfers.Card2Card("5177827685644009", "4716742265786594", 50_00 )

	fmt.Println(total, ok)


}