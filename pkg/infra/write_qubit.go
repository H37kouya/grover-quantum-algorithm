package infra

import (
	"encoding/csv"
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"os"
	"strconv"
	"sync"
)

type ParallelQubitArg struct {
	Qubits model.Qubits
	Path   string
}

func ParallelProcessQubitCsv(parallelQubitArg []*ParallelQubitArg) {
	wg := &sync.WaitGroup{}
	for _, arg := range parallelQubitArg {
		wg.Add(1)
		go func(a *ParallelQubitArg) {
			_ = WriteQubitCsv(&a.Qubits, a.Path)
			wg.Done()
		}(arg)
	}
	wg.Wait()
}

func WriteQubitCsv(
	qubits *model.Qubits,
	path string,
) error {
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	writer := csv.NewWriter(fp)
	err = writer.WriteAll(writableQubitData(qubits, shouldWriteData{
		No:   "No",
		Real: "real",
		Imag: "imag",
		Abs:  "abs",
	}))
	if err != nil {
		return err
	}
	writer.Flush()
	fmt.Println(path + " is created!")

	return nil
}

type shouldWriteData struct {
	No   string
	Real string
	Imag string
	Abs  string
}

func (d shouldWriteData) toCsvRecord() []string {
	return []string{
		d.No,
		d.Real,
		d.Imag,
		d.Abs,
	}
}

func writableQubitData(qubits *model.Qubits, header shouldWriteData) [][]string {
	newRecords := make([][]string, len(*qubits)+1)
	newRecords[0] = header.toCsvRecord()

	for i, qubit := range *qubits {
		newShouldWriteData := shouldWriteData{
			No:   strconv.Itoa(i),
			Real: strconv.FormatFloat(real(qubit), 'f', -1, 64),
			Imag: strconv.FormatFloat(imag(qubit), 'f', -1, 64),
			Abs:  strconv.FormatFloat(qubit.Abs(), 'f', -1, 64),
		}
		newRecords[i+1] = newShouldWriteData.toCsvRecord()
	}

	return newRecords
}
