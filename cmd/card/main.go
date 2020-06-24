package main

import (
	"fmt"
	"github.com/trate/h2.1/pkg/card"
	"github.com/trate/h2.1/pkg/transfer"
	"log"
)

func main() {

	master := &card.Card{
		Id:           1,
		Issuer:       "MasterCard",
		Balance:      65_000,
		Currency:     "RUB",
		Number:       "5106217685644009",
		Transactions: nil,
	}
	visa := &card.Card{
		Id:           2,
		Issuer:       "Visa",
		Balance:      64_000,
		Currency:     "RUB",
		Number:       "5106212265786594",
		Transactions: nil,
	}

	tinkoff := card.NewService("Tinkoff")
	tinkoffTransfers := transfer.NewService(tinkoff, 0.5, 10_00)

	tinkoff.Cards = append(tinkoff.Cards, master)
	tinkoff.Cards = append(tinkoff.Cards, visa)

	total, err := tinkoffTransfers.Card2Card("5106217685644009", "5106212265786594", 50_00)

	if err != nil {
		switch err {
		case transfer.ErrSourceCardBalanceNotEnough:
			fmt.Println("Sorry, can't complete transaction")
		case transfer.ErrTargetCardNotFound:
			fmt.Println("Please check target card number")
		case transfer.ErrSourceCardBalanceNotEnough:
			fmt.Println("Your account has insufficient funds")
		default:
			fmt.Println("Something bad happened. Try again later")
		}
	}

	fmt.Println(total)
	fmt.Println("Check that the cards numbers are valid...")

	from := "4556132133759481"
	to := "4024007104767800"
	err = transfer.Transfer(from, to)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(from, to, "are valid")
}
