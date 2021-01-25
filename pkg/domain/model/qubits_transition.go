package model

// QubitsTransition Qubitsの推移
type QubitsTransition []Qubits

// Column 行の取得 (あるqubitの推移だけを取り出す)
func (qubitsTransition QubitsTransition) Column(col int) QubitTransition {
	newQubitsTransition := make(QubitTransition, 0, len(qubitsTransition))

	for _, qubits := range qubitsTransition {
		newQubit := qubits[col]
		newQubitsTransition = append(newQubitsTransition, newQubit)
	}

	return newQubitsTransition
}

// Row ある回数のqubitを取得する
func (qubitsTransition QubitsTransition) Row(row int) QubitTransition {
	return QubitTransition(qubitsTransition[row])
}
