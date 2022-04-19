package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

/*
7 hourly timeslots ranging from 9am-5pm with 12-1pm as lunch break.
*/
type Dentist struct {
	name      string
	timeSlots [7]bool
}

type node struct {
	dentist Dentist
	next    *node
}

type stack struct {
	top  *node
	size int
}

var dentistList *stack

func (p *stack) push(dentist Dentist) error {
	newNode := &node{dentist: dentist, next: nil}
	if p.top == nil {
		p.top = newNode
	} else {
		newNode.next = p.top
		p.top = newNode
	}
	p.size++
	return nil
}
func (p *stack) pop() (Dentist, error) {
	var item Dentist
	if p.top == nil {
		return Dentist{}, errors.New("Stack empty!")
	}
	item = p.top.dentist
	if p.size == 1 {
		item = p.top.dentist
		p.top = nil
	} else {
		p.top = p.top.next
	}
	p.size--
	return item, nil
}

func (p *stack) removeAt(name string) (Dentist, error) {
	tempStack := &stack{nil, 0}
	d := Dentist{}
	if p.top == nil {
		return Dentist{}, errors.New("Stack is empty.")
	}
	for p.size != 0 {
		temp := p.peek()
		p.pop()
		if temp.name != name {
			tempStack.push(temp)
		} else {
			d = temp
		}
	}
	for tempStack.size != 0 {
		temp := tempStack.peek()
		tempStack.pop()
		p.push(temp)
	}
	if d.name == "" {
		return d, errors.New("Dentist not found.")
	}
	return d, nil
}

func (p *stack) peek() Dentist {
	if p.top == nil {
		return Dentist{}
	}
	return p.top.dentist
}
func (p *stack) sortStack() *stack {
	tempStack := &stack{nil, 0}
	for p.size != 0 {
		temp := p.peek()
		p.pop()
		for tempStack.size != 0 && tempStack.peek().name < temp.name {
			p.push(tempStack.peek())
			tempStack.pop()
		}
		tempStack.push(temp)
	}
	return tempStack
}
func (p *stack) printDentistList() {
	tempStack := &stack{nil, 0}
	for p.top != nil {
		item, _ := p.pop()
		tempStack.push(item)
		printNameAndTime(item)
	}
	for tempStack.top != nil {
		item, _ := tempStack.pop()
		p.push(item)
	}
}

// like printDentistList but just names only.
func (p *stack) printDentistNames() {
	tempStack := &stack{nil, 0}
	for p.top != nil {
		item, _ := p.pop()
		tempStack.push(item)
		fmt.Println(item.name)
	}
	for tempStack.top != nil {
		item, _ := tempStack.pop()
		p.push(item)
	}
}

func (p *stack) availableDoctorsAtTime(index int) {

}

func (p *stack) bookAppointment() {
	fmt.Println("Enter the name of the dentist.")
	p.printDentistNames()
	var personName string
	dentistNode := &node{}
	tempStack := &stack{nil, 0}
	var updated = false
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		personName = scanner.Text()
	}
	fmt.Printf("\nsearching for... %v", personName)
	for p.top != nil {
		item, _ := p.pop()
		tempStack.push(item)
		fmt.Println(item.name)
		if item.name == personName {
			dentistNode = tempStack.top
			addAppointment(dentistNode)
			updated = true
		}
	}
	for tempStack.top != nil {
		item, _ := tempStack.pop()
		p.push(item)
	}
	if !updated {
		fmt.Println("No Such name!")
		p.bookAppointment()
	}
}

func addAppointment(node *node) {
	fmt.Println("Please select a time slot.")
	for i := 0; i < 7; i++ {
		fmt.Printf("\n%v. %v, (%v)", i+1, slotTime(i), printBookedStatus(node.dentist.timeSlots[i]))
	}
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if scanner.Text() == "" {
			mainMenu()
		}
		timeSlotInput, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error. Expected Int.")
			addAppointment(node)
		} else {
			timeSlotInput -= 1
			switch timeSlotInput {
			case 0, 1, 2, 3, 4, 5, 6:
				if node.dentist.timeSlots[timeSlotInput] {
					fmt.Println("Dentist is already booked at " + slotTime(timeSlotInput) + "!")
					addAppointment(node)
				} else {
					node.dentist.timeSlots[timeSlotInput] = true
					fmt.Println("Appointment set at " + slotTime(timeSlotInput) + ".")
				}

			default:
			}
		}
	}
}

func (p *stack) beginRemoveAppointment() {
	fmt.Println("Enter the name of the dentist.")
	p.printDentistNames()
	var personName string
	dentistNode := &node{}
	tempStack := &stack{nil, 0}
	var updated = false
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		personName = scanner.Text()
	}
	fmt.Printf("\nsearching for... %v", personName)
	for p.top != nil {
		item, _ := p.pop()
		tempStack.push(item)
		fmt.Println(item.name)
		if item.name == personName {
			dentistNode = tempStack.top
			removeAppointment(dentistNode)
			updated = true
		}
	}
	for tempStack.top != nil {
		item, _ := tempStack.pop()
		p.push(item)
	}
	if !updated {
		fmt.Println("No Such name!")
		p.beginRemoveAppointment()
	}
}

func removeAppointment(node *node) {
	fmt.Println("Please select a time slot.")
	for i := 0; i < 7; i++ {
		fmt.Printf("\n%v. %v, (%v)", i+1, slotTime(i), printBookedStatus(node.dentist.timeSlots[i]))
	}
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if scanner.Text() == "" {
			mainMenu()
		}
		timeSlotInput, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error. Expected Int.")
			addAppointment(node)
		} else {
			timeSlotInput -= 1
			switch timeSlotInput {
			case 0, 1, 2, 3, 4, 5, 6:
				if !node.dentist.timeSlots[timeSlotInput] {
					fmt.Println("Dentist does not have a booking at " + slotTime(timeSlotInput) + "!")
					removeAppointment(node)
				} else {
					node.dentist.timeSlots[timeSlotInput] = false
					fmt.Println("Appointment removed at " + slotTime(timeSlotInput) + ".")
				}

			default:
			}
		}
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

func printBookedStatus(query bool) string {
	if query {
		return "booked"
	} else {
		return "not booked"
	}

}

func printNameAndTime(dentist Dentist) {
	fmt.Printf("\nDentist %v's timeslots:", dentist.name)
	for i := 0; i < 7; i++ {
		fmt.Printf("\n%v is %v.", slotTime(i), printBookedStatus(dentist.timeSlots[i]))
	}
}
