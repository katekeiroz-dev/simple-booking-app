package main

//Imports the fmt package, which provides functions for formatted I/O (e.g., printing to the console).
import "fmt"

// Declares the main function, where the execution of the program begins.
func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var remainingTickets uint = 50 //This tracks how many tickets are left.
	var bookings []string

	fmt.Printf("Welcome to our %v booking application \n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still availabe \n", confTickets, rremainingTickets)
	fmt.Println("Get your tickets here to attend ")

	//loop to keep asking for user input, and then store the value at the bookings slice
	//Starts an infinite loop so the program keeps asking for user input continuously.
	for {

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

		//Subtracts the number of tickets booked from the remaining tickets.
		remainingTickets = remainingTickets - userTickets

		//store the value at the bookings slice
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v , for booking %v tickets \n. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf(" %v tickets remaing for %v \n", remainingTickets, confName)

		fmt.Printf("These are all the %v in our app \n", bookings)

	}

}
