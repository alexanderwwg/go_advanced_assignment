package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type mainMenuOptions int

const (
	Exit mainMenuOptions = iota
	ViewDentists
	MakeAppointment
	SearchForTimeslot
	RemoveAppointment
	AdminOptions
)

func main() {
	mainMenu()
}

/*
Main Features
- Make Appointment
- List Available Times for Selected Doctor
- Search for Available Doctors at given timeslot
- Edit Appointment to either change timing / remove timeslot.

Admin Features
- Add or Remove Doctor
*/
func mainMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nDentist Appointment System\n=================================")
	fmt.Println("Please select an option below:")
	fmt.Println("0.Exit\n1.View Dentists\n2.Make an Appointment\n3.Search for Timeslot\n4.Remove Appointment\n5.Admin Options")
	if scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error, please try again.")
			mainMenu()
		} else {
			switch input {
			case int(Exit):
			case int(ViewDentists):
				dentistList.printDentistList()
				fmt.Println("\nPress enter to continue.")
				fmt.Scanln()
				mainMenu()
			case int(MakeAppointment):
				dentistList.bookAppointment()
				mainMenu()
			case int(SearchForTimeslot):
			case int(RemoveAppointment):
				dentistList.beginRemoveAppointment()
				mainMenu()
			case int(AdminOptions):
				adminOptions()
			default:
				mainMenu()
			}
		}
	}
}

func adminOptions() {
	fmt.Println("Admin Options")
	fmt.Println("0. Return to Main Menu\n1.Add dentist.\n2.Remove dentist.")
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		input, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("Error, please try again.")
			adminOptions()
		}
		switch input {
		case 0:
			mainMenu()
		case 1:
			addDentistInterface()
			mainMenu()
		case 2:
			removeDentistInterface()
			mainMenu()
		}
	}
}

func addDentistInterface() {
	fmt.Println("Add Dentist\nPlease put in dentist name.")
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		addDentist(s.Text())
	}
}

func removeDentistInterface() {
	fmt.Println("Remove Dentist\nPlease put in dentist name.")
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		removeDentist(s.Text())
	}

}
