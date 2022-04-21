package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

/*
CSV Format
8 Columns: Name, 9-10, 10-11, 11-12, 1-2, 2-3, 3-4, 4-5
*/

var data [][]string

func (p *stack) tempSaveData() {
	tempStack := &stack{nil, 0}
	data = [][]string{
		{"name", "9am-10am", "10am-11am", "11am-12pm", "1pm-2pm", "2pm-3pm", "3pm-4pm", "4pm-5pm"},
	}
	for p.top != nil {
		item, _ := p.pop()
		tempStack.push(item)
		data = append(data, convDentistToStr(item))
	}
	for tempStack.top != nil {
		item, _ := tempStack.pop()
		p.push(item)
	}
}

func convDentistToStr(dentist Dentist) []string {
	returnData := []string{
		dentist.name}
	for _, v := range dentist.timeSlots {
		if v == true {
			returnData = append(returnData, "1")
		} else {
			returnData = append(returnData, "0")
		}
	}
	return returnData
}

func dentistConversion(input [][]string) {
	dentist := Dentist{}
	for i := len(input) - 1; i >= 1; i-- {
		dentist.name = input[i][0]
		for i := 1; i < 7; i++ {

			b, _ := strconv.ParseBool(input[0][i])
			dentist.timeSlots[i] = b
		}
		addDentist(dentist.name, dentist.timeSlots)
	}

}

func writeToCSV(name string) {
	csvFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	csvWriter := csv.NewWriter(csvFile)
	csvWriter.WriteAll(data)

}

func csvExists(filename string) bool {
	if _, err := os.Stat("./" + filename); err == nil {
		return true
	} else {
		return false
	}
}

func loadCSVData() {
	file, err := os.Open("./" + csvFileName)
	if err != nil {
		panic(err)
	}
	//fmt.Println("file opened")
	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	dentistConversion(rows)

	defer file.Close()
}
