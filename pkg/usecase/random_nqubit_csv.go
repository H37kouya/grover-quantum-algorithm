package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"grover-quantum-search/pkg/infra"
	"grover-quantum-search/pkg/lib"
	"math"
)

func RandomNQubitCsvUsecase(n int, qubitPlus model.Qubit) {
	loop := 1000

	qubits := model.RandomQubits(int(math.Pow(2.0, float64(n))))
	qubits = *qubits.Normalize()
	qubits = *qubits.Map(func(qubit model.Qubit, i int) model.Qubit {
		return qubit + qubitPlus
	})

	timeForFileName := lib.GetTimeForFileName()
	targets := []int{1}
	beforeQubits := qubits

	targetQubitList := make(model.Qubits, loop)
	qubitAbsMaxList := make(model.Qubits, loop)
	qubitAbsMinList := make(model.Qubits, loop)
	maxIdx := qubits.MaxIdx()
	minIdx := qubits.MinIdx()

	targetQubitList[0] = qubits[targets[0]]
	qubitAbsMaxList[0] = qubits[maxIdx]
	qubitAbsMinList[0] = qubits[minIdx]
	for i := 1; i < loop; i++ {
		newQubits := *lib.GroverQuantumSearch(&beforeQubits, targets)
		targetQubitList[i] = newQubits[targets[0]]
		qubitAbsMaxList[i] = newQubits[maxIdx]
		qubitAbsMinList[i] = newQubits[minIdx]
		beforeQubits = newQubits
	}
	multipleQubits := make([]*model.Qubits, 0, 3)
	multipleQubits = append(multipleQubits, &targetQubitList, &qubitAbsMaxList, &qubitAbsMinList)

	pararellQubitArgs := make([]*infra.PararellQubitArg, 0, 4)
	pararellQubitArgs = append(pararellQubitArgs,
		&infra.PararellQubitArg{
			Qubits: qubits,
			Path:   "./outputs/"+timeForFileName+"_original.csv",
		},
		&infra.PararellQubitArg{
			Qubits: targetQubitList,
			Path:   "./outputs/"+timeForFileName+"_target.csv",
		},
		&infra.PararellQubitArg{
			Qubits: qubitAbsMaxList,
			Path:   "./outputs/"+timeForFileName+"_max.csv",
		},
		&infra.PararellQubitArg{
			Qubits: qubitAbsMinList,
			Path:   "./outputs/" + timeForFileName + "_min.csv",
		},
	)

	infra.PararellProcessQubitCsv(pararellQubitArgs)

	err := infra.WriteMultipleQubitCsv(
		multipleQubits,
		"./outputs/" + timeForFileName + "_multiple.csv",
	)
	if err != nil {
		panic(err)
	}


	fmt.Println("done")
}
