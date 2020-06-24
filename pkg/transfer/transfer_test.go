package transfer

import (
	"github.com/trate/h2.1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc           *card.Service
		CommissionPercent float64
		MinCommission     int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	cardSvc := card.NewService("Tinkoff")
	cardSvc.Add(
		&card.Card{Balance: 65_000, Number: "4539076789382977"}, &card.Card{Balance: 64_000, Number: "4844649384305716"},
		&card.Card{Balance: 4_000, Number: "4485294233758740055"}, &card.Card{Balance: 34_000, Number: "5594089900819313"},
	)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantError error
	}{
		{
			name:      "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields:    fields{CardSvc: cardSvc, CommissionPercent: 10, MinCommission: 10_00},
			args:      args{from: "4539076789382977", to: "4844649384305716", amount: 50_00},
			wantTotal: 5500,
			wantError: nil,
		},
		{
			name:      "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields:    fields{CardSvc: cardSvc, CommissionPercent: 10, MinCommission: 10_00},
			args:      args{from: "4485294233758740055", to: "5594089900819313", amount: 50_00},
			wantTotal: 5500,
			wantError: ErrSourceCardBalanceNotEnough,
		},
		{
			name:      "Карта своего банка -> Карта не найдена",
			fields:    fields{CardSvc: cardSvc, CommissionPercent: 10, MinCommission: 10_00},
			args:      args{from: "4539076789382977", to: "4844649384305717", amount: 50_00},
			wantTotal: 5500,
			wantError: ErrTargetCardNotFound,
		},
		{
			name:      "Карта не найдена -> Карта своего банка",
			fields:    fields{CardSvc: cardSvc, CommissionPercent: 10, MinCommission: 10_00},
			args:      args{from: "45390767893829778", to: "4844649384305716", amount: 50_00},
			wantTotal: 5500,
			wantError: ErrSourceCardNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:       tt.fields.CardSvc,
				Commission:    tt.fields.CommissionPercent,
				MinCommission: tt.fields.MinCommission,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantError {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantError)
			}
		})
	}
}
