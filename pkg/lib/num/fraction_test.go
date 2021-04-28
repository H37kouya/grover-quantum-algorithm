package num

import (
	"reflect"
	"testing"
)

func TestContinuedFractionExpansion(t *testing.T) {
	type args struct {
		a    float64
		b    float64
		loop int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "連分数展開ができる", args: args{a: 42, b: 11, loop: 100}, want: []float64{3, 1, 4, 2}},
		{name: "連分数展開ができる", args: args{a: 1707, b: 2048, loop: 100}, want: []float64{0, 1, 5, 170, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContinuedFractionExpansion(tt.args.a, tt.args.b, tt.args.loop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContinuedFractionExpansion() = %v, want %v", got, tt.want)
			}
		})
	}
}
