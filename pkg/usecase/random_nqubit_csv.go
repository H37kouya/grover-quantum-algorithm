package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"grover-quantum-search/pkg/infra"
	"grover-quantum-search/pkg/lib"
	"grover-quantum-search/pkg/lib/time"
	"math"
)

func RandomNQubitCsvUsecase(n int, qubitPlus model.Qubit, loop int) {

	qubits := model.RandomQubits(int(math.Pow(2.0, float64(n))))
	qubits = *qubits.Normalize()
	qubits = *qubits.Map(func(qubit model.Qubit, i int) model.Qubit {
		return qubit + qubitPlus
	})

	timeForFileName := time.GetTimeForFileName()
	targets := []int{1}

	newQubitsTransitionData := make(model.QubitsTransition, 0, loop)
	newQubitsTransitionData = append(newQubitsTransitionData, qubits)
	for i := 1; i < loop; i++ {
		newQubits := lib.GroverQuantumSearch(&newQubitsTransitionData[i-1], targets)
		newQubitsTransitionData[i] = *newQubits
	}

	maxIdx := qubits.MaxIdx()
	minIdx := qubits.MinIdx()

	targetQubitTransitionData := newQubitsTransitionData.Column(targets[0])
	qubitTransitionAbsMax := newQubitsTransitionData.Column(maxIdx)
	qubitTransitionAbsMin := newQubitsTransitionData.Column(minIdx)
	qubitTransitionOptional1 := newQubitsTransitionData.Column(targets[0] + 1)
	qubitTransitionOptional2 := newQubitsTransitionData.Column(targets[0] + 2)

	multipleQubits := make([][]model.Qubit, 0, 5)
	multipleQubits = append(
		multipleQubits,
		targetQubitTransitionData,
		qubitTransitionAbsMax,
		qubitTransitionAbsMin,
		qubitTransitionOptional1,
		qubitTransitionOptional2,
	)

	parallelQubitArgs := make([]*infra.ParallelQubitArg, 0, 6)
	parallelQubitArgs = append(parallelQubitArgs,
		&infra.ParallelQubitArg{
			Qubits: qubits,
			Path:   "./outputs/" + timeForFileName + "_original.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: targetQubitTransitionData,
			Path:   "./outputs/" + timeForFileName + "_target.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitTransitionAbsMax,
			Path:   "./outputs/" + timeForFileName + "_max.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitTransitionAbsMin,
			Path:   "./outputs/" + timeForFileName + "_min.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitTransitionOptional1,
			Path:   "./outputs/" + timeForFileName + "_optionals1.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitTransitionOptional2,
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
