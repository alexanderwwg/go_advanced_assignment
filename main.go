package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

/*
Main Features
- Make Appointment
- List Available Times for Selected Doctor
- Search for Available Doctors at given timeslot
- Edit Appointment to either change timing / remove timeslot.
- Save to main.csv

Admin Features
- Add or Remove Doctor
*/

type mainMenuOptions int

const (
	Exit mainMenuOptions = iota
	ViewDentists
	MakeAppointment
	SearchForTimeslot
	RemoveAppointment
	AdminOptions
	Save
)

var wg sync.WaitGroup

func main() {
	mainMenu()
}

func mainMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nDentist Appointment System\n=================================")
	fmt.Println("Please select an option below:")
	fmt.Println("0.Exit\n1.View Dentists\n2.Make an Appointment\n3.Search for Timeslot\n4.Remove Appointment\n5.Admin Options\n6.Save")
	if scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error, please try again.")
			mainMenu()
		} else {
			switch input {
			case int(Exit):
				exit()
			case int(ViewDentists):
				dentistList.printDentistList()
				fmt.Println("\nPress enter to continue.")
				fmt.Scanln()
				mainMenu()
			case int(MakeAppointment):
				dentistList.bookAppointment()
				mainMenu()
			case int(SearchForTimeslot):
				getAvailableDentistInterface()
				mainMenu()
			case int(RemoveAppointment):
				dentistList.beginRemoveAppointment()
				mainMenu()
			case int(AdminOptions):
				adminOptions()
			case int(Save):
				saveCSV()
				wg.Wait()
				mainMenu()
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
		err := addDentist(s.Text(), blankTimeSlots)
		if err == nil {
			dentistList = dentistList.sortStack()
		}
	}
}

func removeDentistInterface() {
	fmt.Println("Remove Dentist\nPlease put in dentist name.")
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		err := removeDentist(s.Text())
		if err == nil {
			dentistList = dentistList.sortStack()
		}
	}

}
func getAvailableDentistInterface() {
	fmt.Println("Select a time slot.")
	for i, _ := range blankTimeSlots {

		fmt.Printf("%v. %v\n", i+1, getTimeSlot(i))
	}
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		input, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("Error, please try again.")
			getAvailableDentistInterface()
		}
		dentistList.getAvailableDentistsAtTime(input - 1)
	}

}

// queries if user wants to save before exitting.
func exit() {
	fmt.Print("Save? (y/n)\n")
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		if s.Text() == "y" || s.Text() == "Y" {
			saveCSV()
			wg.Wait()
			os.Exit(0)
		} else if s.Text() == "n" || s.Text() == "N" {
			os.Exit(0)
		}
	}
}

func saveCSV() {
	wg.Add(1)
	defer wg.Done()
	fmt.Println("Saving...")
	dentistList = dentistList.sortStack()
	dentistList.tempSaveData()
	writeToCSV(csvFileName)
}
