package model

import "math"

type Qubits []Qubit

func (qubits Qubits) Sum() Qubit {
	sum := Qubit(complex(float64(0), float64(0)))

	for _, qubit := range qubits {
		sum += qubit
	}

	return sum
}

func (qubits Qubits) SumOfSquares() Qubit {
	sum := Qubit(complex(float64(0), float64(0)))

	for _, qubit := range qubits {
		sum += qubit * qubit
	}

	return sum
}

func (qubits Qubits) SumAbsOfSquares() float64 {
	sum := 0.0

	for _, qubit := range qubits {
		sum += qubit.Abs() * qubit.Abs()
	}

	return sum
}


func (qubits Qubits) Average() Qubit {
	sumQubit := qubits.Sum()
	lenQubit := len(qubits)
	return sumQubit / Qubit(complex(float64(lenQubit), float64(0)))
}

func (qubits Qubits) Max() Qubit {
	return qubits[qubits.MaxIdx()]
}

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
