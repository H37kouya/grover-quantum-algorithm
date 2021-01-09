package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"grover-quantum-search/pkg/infra"
	"grover-quantum-search/pkg/lib"
	"math"
)

func FixedNQubitCsvUsecase(n int) {
	loop := 3000

	qubits := make(model.Qubits, int(math.Pow(2.0, float64(n))))
	for i := 0; i < len(qubits); i++ {
		qubits[i] = model.Qubit(complex(1, 0))
	}
	qubits = *qubits.Normalize()

	timeForFileName := lib.GetTimeForFileName()
	targets := []int{1}
	newQubitTransitionData := make([]model.Qubits, loop)
	newQubitTransitionData[0] = qubits

	for i := 1; i < loop; i++ {
		newQubits := lib.GroverQuantumSearch(&newQubitTransitionData[i - 1], targets)
		newQubitTransitionData[i] = *newQubits
	}

	targetQubitList := make(model.Qubits, loop)
	qubitNo2List := make(model.Qubits, loop)
	qubitNo3List := make(model.Qubits, loop)
	qubitNo4List := make(model.Qubits, loop)
	qubitAbsMaxList := make(model.Qubits, loop)
	maxIdx := qubits.MaxIdx()
	for i, q := range newQubitTransitionData {
		targetQubitList[i] = q[targets[0]]
		qubitNo2List[i] = q[1]
		qubitNo3List[i] = q[2]
		qubitNo4List[i] = q[3]
		qubitAbsMaxList[i] = q[maxIdx]
	}

	pararellQubitArgs := make([]*infra.PararellQubitArg, 0, 5)
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
			Qubits: qubitNo2List,
			Path:   "./outputs/"+timeForFileName+"_2.csv",
		},
		&infra.PararellQubitArg{
			Qubits: qubitNo3List,
			Path:   "./outputs/"+timeForFileName+"_3.csv",
		},
		&infra.PararellQubitArg{
			Qubits: qubitNo4List,
			Path:   "./outputs/"+timeForFileName+"_4.csv",
		},
		&infra.PararellQubitArg{
			Qubits: qubitNo4List,
			Path:   "./outputs/"+timeForFileName+"_max.csv",
		},
	)

	infra.PararellProcessQubitCsv(pararellQubitArgs)

	fmt.Println("done")
}
