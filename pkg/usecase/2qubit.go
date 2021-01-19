package usecase

import (
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"grover-quantum-search/pkg/infra"
	"grover-quantum-search/pkg/lib"
)

func Fixed2QubitUsecase() {
	qubits := model.Qubits{
		model.Qubit(complex(0.5, 0)),
		model.Qubit(complex(0.5, 0)),
		model.Qubit(complex(0.5, 0)),
		model.Qubit(complex(0.5, 0)),
	}
	loop := 4
	timeForFileName := lib.GetTimeForFileName()
	targets := []int{1}
	newQubitTransitionData := make([]model.Qubits, 0, loop)
	newQubitTransitionData = append(newQubitTransitionData, qubits)

	for i := 1; i < loop; i++ {
		newQubits := lib.GroverQuantumSearch(&newQubitTransitionData[i-1], targets)
		newQubitTransitionData = append(newQubitTransitionData, *newQubits)
	}

	targetQubitList := make(model.Qubits, 0, loop)
	qubitNo2List := make(model.Qubits, 0, loop)
	qubitNo3List := make(model.Qubits, 0, loop)
	qubitNo4List := make(model.Qubits, 0, loop)
	for _, q := range newQubitTransitionData {
		targetQubitList = append(targetQubitList, q[targets[0]])
		qubitNo2List = append(qubitNo2List, q[1])
		qubitNo3List = append(qubitNo3List, q[2])
		qubitNo4List = append(qubitNo4List, q[3])
	}

	pararellQubitArgs := make([]*infra.ParallelQubitArg, 0, 0)
	pararellQubitArgs = append(pararellQubitArgs,
		&infra.ParallelQubitArg{
			Qubits: qubits,
			Path:   "./outputs/" + timeForFileName + "_original.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: targetQubitList,
			Path:   "./outputs/" + timeForFileName + "_target.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitNo2List,
			Path:   "./outputs/" + timeForFileName + "_2.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitNo3List,
			Path:   "./outputs/" + timeForFileName + "_3.csv",
		},
		&infra.ParallelQubitArg{
			Qubits: qubitNo4List,
			Path:   "./outputs/" + timeForFileName + "_4.csv",
		},
	)

	infra.ParallelProcessQubitCsv(pararellQubitArgs)

	fmt.Println("done")
}
