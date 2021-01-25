package model

import (
	"testing"
)

func TestQubits_Sum(t *testing.T) {
	type testCase struct {
		name   string
		qubits Qubits
		want   Qubit
	}

	tests := []testCase{
		{
			name: "(1 + i) + (-1 - i) の合計は0である",
			qubits: func() Qubits {
				qubits := make(Qubits, 0, 2)
				qubits = append(qubits, Qubit(complex(float64(1), float64(1))))
				qubits = append(qubits, Qubit(complex(float64(-1), float64(-1))))
				return qubits
			}(),
			want: Qubit(complex(float64(0), float64(0))),
		},
		{
			name: "(1 + i) * 4 の合計は4 + 4iである",
			qubits: func() Qubits {
				qubits := make(Qubits, 4)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(1), float64(1)))
				}
				return qubits
			}(),
			want: Qubit(complex(float64(4), float64(4))),
		},
		{
			name: "2 * 16 の合計は32である",
			qubits: func() Qubits {
				qubits := make(Qubits, 16)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(2), 0))
				}
				return qubits
			}(),
			want: Qubit(complex(float64(32), float64(0))),
		},
		{
			name: "(2 - i2) * 16 の合計は32 - 32iである",
			qubits: func() Qubits {
				qubits := make(Qubits, 16)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(2), float64(-2)))
				}
				return qubits
			}(),
			want: Qubit(complex(float64(32), float64(-32))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.qubits.Sum() != tt.want {
				t.Errorf("Qubits = %+v, Qubit.Sum() = %v, want %v", tt.qubits, tt.qubits.Sum(), tt.want)
			}
		})
	}
}
