package service

import (
	"grover-quantum-search/pkg/domain/collection"
	"grover-quantum-search/pkg/domain/valueObject"
)

func GroverQuantumSearch(
	qubits *collection.Qubits,
	targets []int,
) *collection.Qubits {
	// 1 オラクルUf いわゆる、γ
	signInversionQubits := qubitSignInversionByIdx(qubits, targets)

	// 平均をとる
	signInversionQubitsAverage := signInversionQubits.Average()

	// 2 P に γ を入れる
	output := signInversionQubits.Map(func(qubit valueObject.Qubit, _ int) valueObject.Qubit {
		return -1*qubit + 2*signInversionQubitsAverage
	})

	return output
}

func qubitSignInversionByIdx(qubits *collection.Qubits, idxes []int) *collection.Qubits {
	qubitSignInversionTargets := qubitSignInversionTargetsType(idxes)

	return qubits.Map(func(qubit valueObject.Qubit, idx int) valueObject.Qubit {
		if qubitSignInversionTargets.contains(idx) {
			return -1 * qubit
		}

		return qubit
	})
}

type qubitSignInversionTargetsType []int

func (targets qubitSignInversionTargetsType) contains(base int) bool {
	for _, t := range targets {
		if t == base {
			return true
		}
	}
	return false
}
