package card

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Проверяем, что номер карты корректен",
			args: args{number: "4024007104767800"},
			want: true,
		},
		{
			name: "Проверяем, что номер карты некорректен",
			args: args{number: "4124007104767800"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.number); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
