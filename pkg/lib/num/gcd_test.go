package num

import (
	"strconv"
	"testing"
)

func TestGcd(t *testing.T) {
	type args struct {
		x int
		m int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{12, 18}, want: 6},
		{args: args{18, 12}, want: 6},
		{args: args{13, 8}, want: 1},
		{args: args{16, 8}, want: 8},
		{args: args{45, 135}, want: 45},
	}
	for _, tt := range tests {
		t.Run(
			strconv.Itoa(tt.args.x)+"と"+strconv.Itoa(tt.args.m)+"の最大公約数は"+strconv.Itoa(tt.want)+"である",
			func(t *testing.T) {
				if got := Gcd(tt.args.x, tt.args.m); got != tt.want {
					t.Errorf("Gcd() = %v, want %v", got, tt.want)
				}
			})
	}
}
