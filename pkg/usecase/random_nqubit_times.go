package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"grover-quantum-search/pkg/domain/service"
	"math"
)

func RandomNQubitTimesCountUsecase(n int, c int, qubitPlusReal float64, qubitPlusImag float64) {
	for i := 0; i < c; i++ {
		RandomNQubitTimesUsecase(n, model.Qubit(complex(qubitPlusReal, qubitPlusImag)))
	}
}

func RandomNQubitTimesUsecase(n int, qubitPlus model.Qubit) {
	loop := 100000

	qubits := model.RandomQubits(int(math.Pow(2.0, float64(n))))
	qubits = *qubits.Normalize()
	qubits = *qubits.Map(func(qubit model.Qubit, i int) model.Qubit {
		return qubit + qubitPlus
	})

	targets := []int{1}
	beforeQubits := qubits
	result1 := 0
	result2 := 0
	result3 := 0
	isDown := false

	for i := 1; i < loop; i++ {
		newQubits := *service.GroverQuantumSearch(&beforeQubits, targets)

		if isDown == false && beforeQubits[targets[0]].Abs() > newQubits[targets[0]].Abs() {
			if result1 == 0 {
				result1 = i - 1
				isDown = true
			} else if result2 == 0 {
				result2 = i - 1
				isDown = true
			} else if result3 == 0 {
				result3 = i - 1
				isDown = true
			} else {
				fmt.Printf("%d\t%d\t%d\t%d\t%d\n", n, result1, result2, result3, i-1)
				return
			}
		} else {
			if isDown == true && beforeQubits[targets[0]].Abs() < newQubits[targets[0]].Abs() {
				isDown = false
			}
		}

		beforeQubits = newQubits
	}

	fmt.Println("error")
}
