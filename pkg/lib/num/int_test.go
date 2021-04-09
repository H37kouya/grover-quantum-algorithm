package num

import (
	"math"
	"strconv"
	"testing"
)

func TestInt64ToInt(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{i: 1}, want: 1},
		{args: args{i: -1}, want: -1},
		{args: args{i: math.MaxInt32}, want: math.MaxInt32},
		{args: args{i: math.MinInt32}, want: math.MinInt32},
		{args: args{i: math.MaxInt32 + 1}, want: 0},
		{args: args{i: math.MinInt32 - 1}, want: 0},
		{args: args{i: math.MaxInt64}, want: 0},
		{args: args{i: math.MinInt64}, want: 0},
	}
	for _, tt := range tests {
		t.Run(
			strconv.FormatInt(tt.args.i, 10)+"は"+strconv.Itoa(tt.want)+"に変換できる",
			func(t *testing.T) {
				if got := Int64ToInt(tt.args.i); got != tt.want {
					t.Errorf("Int64ToInt() = %v, want %v", got, tt.want)
				}
			})
	}
}
