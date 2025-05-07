package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var ramainingTicktes = 50

	fmt.Printf("Welcome to our %v booking application \n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still availabe \n", confTickets, ramainingTicktes)
	fmt.Println("Get your tickets here to attend ")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	//ask user for their name ..
	fmt.Println("Enter your fist name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of ticktes: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v , for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
