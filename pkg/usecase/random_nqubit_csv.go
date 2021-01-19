package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"grover-quantum-search/pkg/infra"
	"grover-quantum-search/pkg/lib"
	"math"
)

func RandomNQubitCsvUsecase(n int, qubitPlus model.Qubit, loop int) {

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
	optionals1List := make(model.Qubits, loop)
	optionals2List := make(model.Qubits, loop)
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
		optionals1List[i] = newQubits[targets[0]+1]
		optionals2List[i] = newQubits[targets[0]+2]

		beforeQubits = newQubits
	}
	multipleQubits := make([]*model.Qubits, 0, 5)
	multipleQubits = append(
		multipleQubits,
		&targetQubitList,
		&qubitAbsMaxList,
		&qubitAbsMinList,
		&optionals1List,
		&optionals2List,
	)

	parallelQubitArgs := make([]*infra.ParallelQubitArg, 0, 6)
	parallelQubitArgs = append(parallelQubitArgs,
		&infra.ParallelQubitArg{
			Qubits: qubits,
			Path:   "./outputs/" + timeForFileName + "_original.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: targetQubitList,
			Path:   "./outputs/" + timeForFileName + "_target.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitAbsMaxList,
			Path:   "./outputs/" + timeForFileName + "_max.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitAbsMinList,
			Path:   "./outputs/" + timeForFileName + "_min.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: optionals1List,
			Path:   "./outputs/" + timeForFileName + "_optionals1.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: optionals2List,
			Path:   "./outputs/" + timeForFileName + "_optionals2.csv",
		},
	)

	infra.ParallelProcessQubitCsv(parallelQubitArgs)

	err := infra.WriteMultipleQubitCsv(
		multipleQubits,
		"./outputs/"+timeForFileName+"_multiple.csv",
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
