package main

import (
	"errors"
	"fmt"
)

/*
7 hourly timeslots ranging from 9am-5pm with 12-1pm as lunch break.
*/
type Dentist struct {
	name      string
	timeSlots [7]bool
}

func addAppt(dentist Dentist, slot int) error {
	if !dentist.timeSlots[slot] {
		fmt.Printf("\nDentist %v's timeslot at %v", dentist.name, dentist.timeSlots[slot])
		dentist.timeSlots[slot] = true
		return nil
	} else {
		return errors.New("Dentist slot already taken. Please choose another time slot.")
	}

}
func removeAppt(dentist Dentist, slot int) error {
	if dentist.timeSlots[slot] {
		fmt.Printf("\nDentist %v's timeslot at %v has been removed.", dentist.name, dentist.timeSlots[slot])
		dentist.timeSlots[slot] = false
		return nil
	} else {
		return errors.New("Dentist slot is empty!")
	}
}

func slotTime(slot int) string {
	var timeSlot string
	switch slot {
	case 0:
		timeSlot = "9am-10am"
	case 1:
		timeSlot = "10am-11am"
	case 2:
		timeSlot = "11am-12pm"
	case 3:
		timeSlot = "1pm-2pm"
	case 4:
		timeSlot = "2pm-3pm"
	case 5:
		timeSlot = "3pm-4pm"
	case 6:
		timeSlot = "4pm-5pm"
	}
	return timeSlot
}

func isBooked(query bool) string {
	if query {
		return "booked"
	} else {
		return "not booked"
	}

}

func printNameAndTime(dentist Dentist) {
	fmt.Printf("Dentist %v's timeslots:", dentist.name)
	for i := 0; i < 7; i++ {
		fmt.Printf("\n%v is %v.", slotTime(i), isBooked(dentist.timeSlots[i]))
	}
}
