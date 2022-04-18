package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

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

func (p *stack) getDentist(name string) *node {
	tempStack := &stack{nil, 0}
	dentistNode := &node{}
	for p.top != nil {
		item, _ := p.pop()
		tempStack.push(item)
		fmt.Println(item.name)
		if item.name == name {
			dentistNode = tempStack.top
			fmt.Println("Match!")
		}
	}
	for tempStack.top != nil {
		item, _ := tempStack.pop()
		p.push(item)
	}
	return dentistNode
}

func (p *stack) addAppointment() {
	fmt.Println("Enter the name of the dentist.")
	var personName string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		personName = scanner.Text()
	}
	fmt.Printf("\nsearching for... %v", personName)
	if p.getDentist(personName).dentist.name == "" {
		fmt.Println("No Such name!")
		p.addAppointment()
	} else {
		node := p.getDentist(personName)
		fmt.Println("Please select a time slot.")
		for i := 0; i < 7; i++ {
			fmt.Printf("\n%v. %v, (%v)", i+1, slotTime(i), isBooked(node.dentist.timeSlots[i]))
		}
	}
}

func init() {
	dentistList = &stack{nil, 0}
	var blankTimeSlots [7]bool
	dentistList.push(Dentist{name: "Zane Ping", timeSlots: blankTimeSlots})
	dentistList.push(Dentist{name: "Amos Tan", timeSlots: blankTimeSlots})
	dentistList.push(Dentist{name: "Melvin Oh", timeSlots: blankTimeSlots})
	dentistList.push(Dentist{name: "Kelsey Sim", timeSlots: blankTimeSlots})
	dentistList.push(Dentist{name: "Serene Keng", timeSlots: [7]bool{true, false, false, true, false, true, false}})
	dentistList = dentistList.sortStack()
}
