package model

import (
	"math/cmplx"
)

type Qubit complex128

func (qubit Qubit) ToComplex128() complex128 {
	return complex128(qubit)
}

func (qubit Qubit) Abs() float64 {
	return cmplx.Abs(qubit.ToComplex128())
}

