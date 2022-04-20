package main

import (
	"errors"
)

var blankTimeSlots [7]bool
var csvFileName = "main.csv"

func (p *stack) doesNameExist(name string) bool {
	tempStack := &stack{nil, 0}
	exists := false
	for p.top != nil {
		item, _ := p.pop()
		if item.name == name {
			exists = true
		}
		tempStack.push(item)
	}
	for tempStack.top != nil {
		item, _ := tempStack.pop()
		p.push(item)
	}
	return exists
}
func addDentist(name string, slots [7]bool) error {
	if name != "" {
		if dentistList.doesNameExist(name) {
			return errors.New("Name exists!")
		} else {
			d := Dentist{name, slots}
			dentistList.push(d)
			return nil
		}
	} else {
		return errors.New("Name is empty!")
	}
}
func removeDentist(name string) error {
	if name != "" {
		if dentistList.doesNameExist(name) {
			dentistList.removeAt(name)
			return nil
		} else {
			return errors.New("Dentist name does not exist.")
		}
	} else {
		return errors.New("Dentist name is empty!")
	}
}
func init() {

	dentistList = &stack{nil, 0}
	if !csvExists(csvFileName) {
		addDentist("Zane Ping", blankTimeSlots)
		addDentist("Amos Tan", blankTimeSlots)
		addDentist("Kelsey Sim", blankTimeSlots)
		addDentist("Serene Keng", blankTimeSlots)
		dentistList = dentistList.sortStack()
		dentistList.tempSaveData()
	} else {
		loadCSVData()
	}

}
