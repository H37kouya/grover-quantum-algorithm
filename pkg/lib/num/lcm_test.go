package num

import "testing"

func TestLcm(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "最小公倍数は9である", args: args{x: 1, y: 9}, want: 9},
		{name: "最小公倍数は12である", args: args{x: 4, y: 6}, want: 12},
		{name: "最小公倍数は60である", args: args{x: 5, y: 12}, want: 60},
		{name: "最小公倍数は512である", args: args{x: 2, y: 512}, want: 512},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcm(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}
