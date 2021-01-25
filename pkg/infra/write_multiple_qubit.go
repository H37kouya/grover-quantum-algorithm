package infra

import (
	"encoding/csv"
	"fmt"
	"grover-quantum-search/pkg/domain/model"
	"os"
	"strconv"
)

func WriteMultipleQubitCsv(
	multipleQubits [][]model.Qubit,
	path string,
) error {
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	writer := csv.NewWriter(fp)
	data := make([]shouldMultipleChildData, 0, len(multipleQubits))
	for idx := range multipleQubits {
		data = append(data, shouldMultipleChildData{
			Real: "Real" + strconv.Itoa(idx+1),
			Imag: "Imag" + strconv.Itoa(idx+1),
			Abs:  "Abs" + strconv.Itoa(idx+1),
		})
	}

	err = writer.WriteAll(writableMultipleQubitData(multipleQubits, shouldMultipleWriteData{
		No:   "No",
		data: data,
	}))
	if err != nil {
		return err
	}
	writer.Flush()
	fmt.Println(path + " is created!")

	return nil
}

type shouldMultipleWriteData struct {
	No   string
	data []shouldMultipleChildData
}

type shouldMultipleChildData struct {
	Real string
	Imag string
	Abs  string
}

func (d shouldMultipleWriteData) toCsvRecord() []string {
	results := make([]string, 0, 1+len(d.data))
	results = append(results, d.No)

	for _, v := range d.data {
		results = append(results, v.Real, v.Imag, v.Abs)
	}

	return results
}

func writableMultipleQubitData(multipleQubits [][]model.Qubit, header shouldMultipleWriteData) [][]string {
	newRecords := make([][]string, 0, 0)
	newRecords = append(newRecords, header.toCsvRecord())

	for newRecordIdx := 1; newRecordIdx < len(multipleQubits[0]); newRecordIdx++ {
		data := make([]shouldMultipleChildData, 0, len(multipleQubits))
		for multipleQubitsIdx := 0; multipleQubitsIdx < len(multipleQubits); multipleQubitsIdx++ {
			qubit := (multipleQubits[multipleQubitsIdx])[newRecordIdx]
			data = append(data, shouldMultipleChildData{
				Real: strconv.FormatFloat(real(qubit), 'f', -1, 64),
				Imag: strconv.FormatFloat(imag(qubit), 'f', -1, 64),
				Abs:  strconv.FormatFloat(qubit.Abs(), 'f', -1, 64),
			})
		}
		shouldMultipleData := shouldMultipleWriteData{
			No:   strconv.Itoa(newRecordIdx),
			data: data,
		}
		newRecords = append(newRecords, shouldMultipleData.toCsvRecord())
	}

	return newRecords
}
