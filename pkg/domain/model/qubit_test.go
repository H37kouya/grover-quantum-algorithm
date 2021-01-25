package model

import "testing"

func TestQubit_Abs(t *testing.T) {
	type testCase struct {
		name  string
		qubit Qubit
		want  float64
	}
	tests := []testCase{
		{name: "0の絶対値は0である", qubit: Qubit(complex(float64(0), float64(0))), want: 0},
		{name: "1+i1の絶対値は1である", qubit: Qubit(complex(float64(1), float64(1))), want: 1},
		{name: "-1-i1の絶対値は1である", qubit: Qubit(complex(float64(-1), float64(-1))), want: 1},
		{name: "10の絶対値は10である", qubit: Qubit(complex(float64(10), float64(0))), want: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.qubit.Abs() != tt.want {
				t.Errorf("Qubit = %v, Qubit.Abs(), want %v", tt.qubit, tt.want)
			}
		})
	}
}
