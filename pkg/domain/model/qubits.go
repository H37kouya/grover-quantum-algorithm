package model

import (
	"crypto/rand"
	"math"
	"math/big"
)

type Qubits []Qubit

// MakeNQubits NQubitsを生成する
func MakeNQubits(n int) Qubits {
	return make(Qubits, 0, int(math.Pow(2.0, float64(n))))
}

// Average 平均値
func (qubits Qubits) Average() Qubit {
	sumQubit := qubits.Sum()
	lenQubit := len(qubits)
	return sumQubit / Qubit(complex(float64(lenQubit), float64(0)))
}

// Sum 合計値を求める
func (qubits Qubits) Sum() Qubit {
	sum := Qubit(complex(float64(0), float64(0)))

	for _, qubit := range qubits {
		sum += qubit
	}

	return sum
}

// SumOfSquares 二乗の和を求める
func (qubits Qubits) SumOfSquares() Qubit {
	sum := Qubit(complex(float64(0), float64(0)))

	for _, qubit := range qubits {
		sum += qubit * qubit
	}

	return sum
}

// SumAbsOfSquares 二乗の絶対値の和を求める
func (qubits Qubits) SumAbsOfSquares() float64 {
	sum := 0.0

	for _, qubit := range qubits {
		sum += qubit.Abs() * qubit.Abs()
	}

	return sum
}

// Max 最大値
func (qubits Qubits) Max() Qubit {
	return qubits[qubits.MaxIdx()]
}

// MaxIdx 最大値の添字を取得
func (qubits Qubits) MaxIdx() int {
	max := Qubit(complex(float64(0), float64(0)))
	maxIdx := 0

	for idx, qubit := range qubits {
		if qubit.Abs() > max.Abs() {
			max = qubit
			maxIdx = idx
		}
	}

	return maxIdx
}

func (qubits Qubits) Min() Qubit {
	return qubits[qubits.MinIdx()]
}

func (qubits Qubits) MinIdx() int {
	min := qubits[0]
	minIdx := 0

	for idx, qubit := range qubits {
		if qubit.Abs() < min.Abs() {
			min = qubit
			minIdx = idx
		}
	}

	return minIdx
}

func (qubits Qubits) Map(callback func(Qubit, int) Qubit) *Qubits {
	newQubits := make(Qubits, len(qubits))

	for idx, qubit := range qubits {
		newQubits[idx] = callback(qubit, idx)
	}

	return &newQubits
}

func (qubits *Qubits) MapPointer(callback func(Qubit, int) Qubit) *Qubits {
	for idx, qubit := range *qubits {
		qubit = callback(qubit, idx)
	}

	return qubits
}

// Normalize 正規化する
func (qubits Qubits) Normalize() *Qubits {
	total := 0.0

	for _, qubit := range qubits {
		abs := qubit.Abs()
		total += abs * abs
	}

	total = math.Sqrt(total)

	return qubits.Map(func(qubit Qubit, i int) Qubit {
		return qubit / Qubit(complex(total, 0))
	})
}

func RandomQubits(len int) Qubits {
	newQubits := make(Qubits, len)

	for i := 0; i < len; i++ {
		r1, _ := rand.Int(rand.Reader, big.NewInt(1e10))
		r2, _ := rand.Int(rand.Reader, big.NewInt(1e10))

		newQubits[i] = Qubit(complex(float64(r1.Int64()-1e10/2), float64(r2.Int64()-1e10/2)))
	}

	return newQubits
}
