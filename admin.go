package main

import (
	"errors"
)

var blankTimeSlots [7]bool

func init() {
	dentistList = &stack{nil, 0}

	addDentist("Zane Ping")
	addDentist("Amos Tan")
	addDentist("Kelsey Sim")
	addDentist("Serene Keng")
	addDentist("Serene Keng")
	dentistList = dentistList.sortStack()

}

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

func addDentist(name string) error {
	if name != "" {
		if dentistList.doesNameExist(name) {
			return errors.New("Name exists!")
		} else {
			d := Dentist{name, blankTimeSlots}
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
