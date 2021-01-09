package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"grover-quantum-search/pkg/lib"
	"math"
)

func FixedNQubitTimesAllUsecase() {
	/**
	 * Memory
	 * 25: 2GB
	 * 26: 4GB
	 */
	for i := 4; i <= 27; i++ {
		FixedNQubitTimesUsecase(i)
	}
}

func FixedNQubitTimesUsecase(n int) {
	loop := 100000

	qubits := make(model.Qubits, int(math.Pow(2.0, float64(n))))
	for i := 0; i < len(qubits); i++ {
		qubits[i] = model.Qubit(complex(1, 0))
	}
	qubits = *qubits.Normalize()

	targets := []int{1}
	beforeQubits := qubits

	for i := 1; i < loop; i++ {
		newQubits := *lib.GroverQuantumSearch(&beforeQubits, targets)

		if beforeQubits[targets[0]].Abs() > newQubits[targets[0]].Abs() {
			fmt.Println(n, i-1)
			return
		}

		beforeQubits = newQubits
	}

	fmt.Println("error")
}
