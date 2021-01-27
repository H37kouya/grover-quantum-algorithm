package collection

import (
	"crypto/rand"
	"grover-quantum-search/pkg/domain/valueObject"
	"math"
	"math/big"
)

type Qubits []valueObject.Qubit

// MakeNQubits NQubitsを生成する  要素数 0, 容量 2^n
func MakeNQubits(n valueObject.N) Qubits {
	return make(Qubits, 0, n.ElementCount())
}

// Average 平均値
func (qubits Qubits) Average() valueObject.Qubit {
	sumQubit := qubits.Sum()
	lenQubit := len(qubits)
	return sumQubit / valueObject.NewQubit(float64(lenQubit), 0.0)
}

func (qubits Qubits) Equal(other Qubits) bool {
	// 要素数と容量が一致するか
	if len(qubits) != len(other) || cap(qubits) != cap(other) {
		return false
	}

	for i := 0; i < len(qubits); i++ {
		if !qubits[i].Equal(other[i]) {
			return false
		}
	}

	return true
}

// Sum 合計値を求める
func (qubits Qubits) Sum() valueObject.Qubit {
	sum := valueObject.NewQubit(0.0, 0.0)

	for _, qubit := range qubits {
		sum += qubit
	}

	return sum
}

// SumOfSquares 二乗の和を求める
func (qubits Qubits) SumOfSquares() valueObject.Qubit {
	sum := valueObject.NewQubit(0.0, 0.0)

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
func (qubits Qubits) Max() valueObject.Qubit {
	return qubits[qubits.MaxIdx()]
}

// MaxIdx 最大値の添字を取得
func (qubits Qubits) MaxIdx() int {
	max := valueObject.NewQubit(0.0, 0.0)
	maxIdx := 0

	for idx, qubit := range qubits {
		if qubit.Abs() > max.Abs() {
			max = qubit
			maxIdx = idx
		}
	}

	return maxIdx
}

func (qubits Qubits) Min() valueObject.Qubit {
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

func (qubits Qubits) Map(callback func(valueObject.Qubit, int) valueObject.Qubit) *Qubits {
	newQubits := make(Qubits, len(qubits))

	for idx, qubit := range qubits {
		newQubits[idx] = callback(qubit, idx)
	}

	return &newQubits
}

func (qubits *Qubits) MapPointer(callback func(valueObject.Qubit, int) valueObject.Qubit) *Qubits {
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

	return qubits.Map(func(qubit valueObject.Qubit, i int) valueObject.Qubit {
		return qubit / valueObject.NewQubit(total, 0)
	})
}

func RandomNQubit(n valueObject.N) Qubits {
	newQubits := make(Qubits, 0, n.ElementCount())

	for i := 0; i < cap(newQubits); i++ {
		r1, _ := rand.Int(rand.Reader, big.NewInt(1e10))
		r2, _ := rand.Int(rand.Reader, big.NewInt(1e10))

		newQubits = append(newQubits, valueObject.NewQubit(float64(r1.Int64()-1e10/2), float64(r2.Int64()-1e10/2)))
	}

	return RandomQubits(n.ElementCount())
}

func RandomQubits(len int) Qubits {
	newQubits := make(Qubits, 0, len)

	for i := 0; i < cap(newQubits); i++ {
		r1, _ := rand.Int(rand.Reader, big.NewInt(1e10))
		r2, _ := rand.Int(rand.Reader, big.NewInt(1e10))

		newQubits = append(newQubits, valueObject.NewQubit(float64(r1.Int64()-1e10/2), float64(r2.Int64()-1e10/2)))
	}

	return newQubits
}
