package main

//Imports the fmt package, which provides functions for formatted I/O (e.g., printing to the console).
import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const confTickets = 50

var confName = "Go Conference"
var remainingTickets uint = 50 //This tracks how many tickets are left.
var bookings []string

// Declares the main function, where the execution of the program begins.
func main() {

	greetUsers()

	//loop to keep asking for user input, and then store the value at the bookings slice
	//Starts an infinite loop so the program keeps asking for user input continuously.
	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//to handle the case of the user input more the the number of tickets available
		//invalid input from the users
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, confName, email)

			firstNames := getFirstNames()                                   // using a fuction that return a value
			fmt.Printf("The first name of bookings are : %v\n", firstNames) //using the valeu firstNames that
			//return from the function and assing into a varible call firstNames that get the function getFirstNames
			//break the loop

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is sold out.")
				break //break the loop
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println(" number of tickets you entered is invalid ")
			}

		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application \n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still availabe \n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend ")
}

func getFirstNames() []string {
	firstNames := []string{} //Initializes an empty slice of strings called firstNames.This will store the first names extracted from each booking.

	//This is a for loop that iterates over the bookings slice.booking represents each item (a full name string like "John Doe").The underscore _ is a blank identifier used to ignore the index (which is normally the first value returned by range).In Go, if you don’t plan to use a value, you must explicitly ignore it with _.

	for _, booking := range bookings { // a blank indentifier "_" ignore a variable you don't want to use
		//so whth Go you need to male unused variables explicit
		var names = strings.Fields(booking) //splits the string

		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their name ..
	fmt.Println("Enter your fist name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickts: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, confName string, email string) {
	//Subtracts the number of tickets booked from the remaining tickets.
	remainingTickets = remainingTickets - userTickets

	//store the value at the bookings slice
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v , for booking %v tickets \n. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf(" %v tickets remaing for %v \n", remainingTickets, confName)
}
