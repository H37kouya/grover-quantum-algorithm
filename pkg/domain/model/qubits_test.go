package model

import (
	"math"
	"testing"
)

func TestQubits_MakeNQubits(t *testing.T) {
	type testCase struct {
		name         string
		n            int
		wantLength   int
		wantCapacity int
	}

	tests := []testCase{
		{
			name:         "1qubitが生成できる",
			n:            1,
			wantLength:   0,
			wantCapacity: 2,
		},
		{
			name:         "2qubitが生成できる",
			n:            2,
			wantLength:   0,
			wantCapacity: 4,
		},
		{
			name:         "3qubitが生成できる",
			n:            3,
			wantLength:   0,
			wantCapacity: 8,
		},
		{
			name:         "5qubitが生成できる",
			n:            5,
			wantLength:   0,
			wantCapacity: 32,
		},
		{
			name:         "10qubitが生成できる",
			n:            10,
			wantLength:   0,
			wantCapacity: 1024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if len(MakeNQubits(tt.n)) != tt.wantLength {
				t.Errorf("要素数が%vでした。want %v", len(MakeNQubits(tt.n)), tt.wantLength)
			}

			if cap(MakeNQubits(tt.n)) != tt.wantCapacity {
				t.Errorf("容量が%vでした。want %v", cap(MakeNQubits(tt.n)), tt.wantCapacity)
			}
		})
	}
}

func TestQubits_Average(t *testing.T) {
	type testCase struct {
		name   string
		qubits Qubits
		want   Qubit
	}

	tests := []testCase{
		{
			name: "(1 + i) + (-1 - i) の平均値は0である",
			qubits: func() Qubits {
				qubits := make(Qubits, 0, 2)
				qubits = append(qubits, Qubit(complex(float64(1), float64(1))))
				qubits = append(qubits, Qubit(complex(float64(-1), float64(-1))))
				return qubits
			}(),
			want: Qubit(complex(float64(0), float64(0))),
		},
		{
			name: "(1 + i) * 4 の平均値は1 + 1iである",
			qubits: func() Qubits {
				qubits := make(Qubits, 4)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(1), float64(1)))
				}
				return qubits
			}(),
			want: Qubit(complex(float64(1), float64(1))),
		},
		{
			name: "2 * 16 の平均値は2である",
			qubits: func() Qubits {
				qubits := make(Qubits, 16)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(2), 0))
				}
				return qubits
			}(),
			want: Qubit(complex(float64(2), float64(0))),
		},
		{
			name: "(2 - i2) * 16 の平均値は2 - i2である",
			qubits: func() Qubits {
				qubits := make(Qubits, 16)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(2), float64(-2)))
				}
				return qubits
			}(),
			want: Qubit(complex(float64(2), float64(-2))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.qubits.Average() != tt.want {
				t.Errorf("Qubits = %+v, Qubits.Average() = %v, want %v", tt.qubits, tt.qubits.Average(), tt.want)
			}
		})
	}
}

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
			t.Parallel()

			if tt.qubits.Sum() != tt.want {
				t.Errorf("Qubits = %+v, Qubit.Sum() = %v, want %v", tt.qubits, tt.qubits.Sum(), tt.want)
			}
		})
	}
}

func TestQubits_Normalize(t *testing.T) {
	type testCase struct {
		name   string
		qubits Qubits
		want   Qubits
	}

	tests := []testCase{
		{
			name: "{(1 + i), (-1 - i)} を正規化した値である",
			qubits: func() Qubits {
				qubits := make(Qubits, 0, 2)
				qubits = append(qubits, Qubit(complex(float64(1), float64(1))))
				qubits = append(qubits, Qubit(complex(float64(-1), float64(-1))))
				return qubits
			}(),
			want: func() Qubits {
				qubits := make(Qubits, 0, 2)
				qubits = append(qubits, Qubit(complex(0.5, 0.5)))
				qubits = append(qubits, Qubit(complex(-0.5, -0.5)))
				return qubits
			}(),
		},
		{
			name: "1 * 4 を正規化した値である",
			qubits: func() Qubits {
				qubits := make(Qubits, 4)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(1), 0))
				}
				return qubits
			}(),
			want: func() Qubits {
				qubits := make(Qubits, 4)
				for i := range qubits {
					qubits[i] = Qubit(complex(0.5, 0))
				}
				return qubits
			}(),
		},
		{
			name: "(1 + i) * 4 を正規化した値である",
			qubits: func() Qubits {
				qubits := make(Qubits, 4)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(1), float64(1)))
				}
				return qubits
			}(),
			want: func() Qubits {
				qubits := make(Qubits, 4)
				for i := range qubits {
					qubits[i] = Qubit(complex(0.5/math.Sqrt(2.0), 0.5/math.Sqrt(2.0)))
				}
				return qubits
			}(),
		},
		{
			name: "2 * 16 を正規化した値である",
			qubits: func() Qubits {
				qubits := make(Qubits, 16)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(2), 0))
				}
				return qubits
			}(),
			want: func() Qubits {
				qubits := make(Qubits, 16)
				for i := range qubits {
					qubits[i] = Qubit(complex(0.25, 0))
				}
				return qubits
			}(),
		},
		{
			name: "(2 - i2) * 16 を正規化した値である",
			qubits: func() Qubits {
				qubits := make(Qubits, 16)
				for i := range qubits {
					qubits[i] = Qubit(complex(float64(2), float64(-2)))
				}
				return qubits
			}(),
			want: func() Qubits {
				qubits := make(Qubits, 16)
				for i := range qubits {
					qubits[i] = Qubit(complex(math.Sqrt(2)/8.0, -1*math.Sqrt(2)/8.0))
				}
				return qubits
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if !(*tt.qubits.Normalize()).Equal(tt.want) {
				t.Errorf("Qubits = %+v, Qubit.Normalize() = %v, want %v", tt.qubits, *tt.qubits.Normalize(), tt.want)
			}
		})
	}
}
