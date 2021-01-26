package model

import (
	"grover-quantum-search/pkg/lib/num"
	"math/cmplx"
)

type Qubit complex128

func (qubit Qubit) ToComplex128() complex128 {
	return complex128(qubit)
}

// Real 実数を返す
func (qubit Qubit) Real() float64 {
	return real(qubit.ToComplex128())
}

// Imag 虚数を返す
func (qubit Qubit) Imag() float64 {
	return imag(qubit.ToComplex128())
}

// Abs 絶対値を取得
func (qubit Qubit) Abs() float64 {
	return cmplx.Abs(qubit.ToComplex128())
}

// Equal Qubitが一致するか
func (qubit Qubit) Equal(other Qubit) bool {
	if num.Round(qubit.Real(), 8) != num.Round(other.Real(), 8) {
		return false
	}

	return num.Round(qubit.Imag(), 8) == num.Round(other.Imag(), 8)
}
