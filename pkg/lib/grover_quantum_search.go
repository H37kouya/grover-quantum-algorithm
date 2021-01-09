package lib

import (
	"grover-quantum-search/pkg/domain/model"
)

func GroverQuantumSearch(
	qubits *model.Qubits,
	targets []int,
) *model.Qubits {
	// 1 オラクルUf いわゆる、γ
	signInversionQubits := qubitSignInversionByIdx(qubits, targets)

	// 平均をとる
	signInversionQubitsAverage := signInversionQubits.Average()

	// 2 P に γ を入れる
	output := signInversionQubits.Map(func(qubit model.Qubit, _ int) model.Qubit {
		return -1*qubit + 2*signInversionQubitsAverage
	})

	return output
}

func qubitSignInversionByIdx(qubits *model.Qubits, idxes []int) *model.Qubits {
	qubitSignInversionTargets := qubitSignInversionTargetsType(idxes)

	return qubits.Map(func(qubit model.Qubit, idx int) model.Qubit {
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
