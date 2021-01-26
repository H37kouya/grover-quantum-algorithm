package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"grover-quantum-search/pkg/domain/service"
	"grover-quantum-search/pkg/domain/valueObject"
)

func FixedNQubitTimesAllUsecase(min int, max int) {
	if min != 0 {
		min = 4
	}
	if max != 0 {
		max = 25
	}
	/**
	 * Memory
	 * 25: 2GB
	 * 26: 4GB
	 */
	for i := min; i <= max; i++ {
		FixedNQubitTimesUsecase(i)
	}
}

func FixedNQubitTimesUsecase(n int) {
	loop := 100000

	newN, err := valueObject.NewN(n)
	if err != nil {
		panic(err)
	}

	qubits := model.MakeNQubits(newN)
	for i := 0; i < cap(qubits); i++ {
		qubits = append(qubits, model.Qubit(complex(1, 1)))
	}
	qubits = *qubits.Normalize()

	targets := []int{1}
	beforeQubits := qubits

	for i := 1; i < loop; i++ {
		newQubits := *service.GroverQuantumSearch(&beforeQubits, targets)

		if beforeQubits[targets[0]].Abs() > newQubits[targets[0]].Abs() {
			fmt.Println(n, i-1)
			return
		}

		beforeQubits = newQubits
	}

	fmt.Println("error")
}
