package transfer

import (
	"github.com/trate/h2.1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc       *card.Service
		Commission    float64
		MinCommission int64
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
		wantOk    bool
	}{
		// TODO: Add test cases.
		{
			name:      "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields:    fields{ CardSvc: cardSvc, Commission: 10, MinCommission: 10_00},
			args:      args{from: "4539076789382977", to: "4844649384305716", amount: 50_00},
			wantTotal: 5500,
			wantOk:    true,
		},
		{
			name:      "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields:    fields{ CardSvc: cardSvc, Commission: 10, MinCommission: 10_00},
			args:      args{from: "4485294233758740055", to: "5594089900819313", amount: 50_00},
			wantTotal: 5500,
			wantOk:    false,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields:    fields{ CardSvc: cardSvc, Commission: 10, MinCommission: 10_00},
			args:      args{from: "4539076789382977", to: "4844649384305717", amount: 50_00},
			wantTotal: 5500,
			wantOk:    true,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields:    fields{ CardSvc: cardSvc, Commission: 10, MinCommission: 10_00},
			args:      args{from: "4485294233758740055", to: "5594089900819318", amount: 50_00},
			wantTotal: 5500,
			wantOk:    false,
		},
		{
			name:      "Карта чужого банка -> Карта своего банка",
			fields:    fields{ CardSvc: cardSvc, Commission: 10, MinCommission: 10_00},
			args:      args{from: "45390767893829778", to: "4844649384305716", amount: 50_00},
			wantTotal: 5500,
			wantOk:    true,
		},
		{
			name:      "Карта чужого банка -> Карта чужого банка",
			fields:    fields{ CardSvc: cardSvc, Commission: 10, MinCommission: 10_00},
			args:      args{from: "45390767893829778", to: "48446493843057166", amount: 50_00},
			wantTotal: 5500,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:       tt.fields.CardSvc,
				Commission:    tt.fields.Commission,
				MinCommission: tt.fields.MinCommission,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
