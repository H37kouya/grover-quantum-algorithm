package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/collection"
	"grover-quantum-search/pkg/domain/service"
	"grover-quantum-search/pkg/domain/valueObject"
	"grover-quantum-search/pkg/infra"
	"grover-quantum-search/pkg/lib/time"
)

func FixedNQubitCsvUsecase(n int) {
	loop := 3000

	newN, err := valueObject.NewN(n)
	if err != nil {
		panic(err)
	}

	qubits := collection.MakeNQubits(newN)
	for i := 0; i < cap(qubits); i++ {
		qubits = append(qubits, valueObject.Qubit(complex(1, 1)))
	}
	qubits = *qubits.Normalize()

	timeForFileName := time.GetTimeForFileName()
	targets := []int{1}
	newQubitsTransitionData := make(collection.QubitsTransition, 0, loop)
	newQubitsTransitionData = append(newQubitsTransitionData, qubits)

	for i := 1; i < loop; i++ {
		newQubits := service.GroverQuantumSearch(&newQubitsTransitionData[i-1], targets)
		newQubitsTransitionData[i] = *newQubits
	}

	targetQubitTransitionData := newQubitsTransitionData.Column(targets[0])
	qubitTransitionNo2 := newQubitsTransitionData.Column(1)
	qubitTransitionNo3 := newQubitsTransitionData.Column(2)
	qubitTransitionNo4 := newQubitsTransitionData.Column(3)

	parallelQubitArgs := make([]*infra.ParallelQubitArg, 0, 5)
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
			Qubits: qubitTransitionNo2,
			Path:   "./outputs/" + timeForFileName + "_2.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitTransitionNo3,
			Path:   "./outputs/" + timeForFileName + "_3.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitTransitionNo4,
			Path:   "./outputs/" + timeForFileName + "_4.csv",
		},
	)

	infra.ParallelProcessQubitCsv(parallelQubitArgs)

	fmt.Println("done")
}
