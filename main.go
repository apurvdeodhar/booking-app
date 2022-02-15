package main

import (
	"fmt"
	"strconv"
	"booking-app/helper"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main()  {

	greetMessage()

	for  {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.UserInputValidate(firstName, lastName,email,userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, email, firstName, lastName)

			fmt.Printf("First names of the booking are: %v\n", getFirstNames())

			if remainingTickets == 0 {
				// end application
				fmt.Printf("The %v has been booked, please come back next year.\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name must have at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("Email must have '@' sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number should be > 0 and < 50.")
			}
		}
	}
}

// Prints greeting message
func greetMessage()  {
	fmt.Printf("Welcome to %v ticket booking platform.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available, Hurry Up !!\n", conferenceTickets, remainingTickets)
}

// Spits out the First Name
func getFirstNames() []string {
	result := []string{}
	for _, booking := range bookings {
		var names = booking["firstName"]
		result = append(result, names)
	}

	return result
}

// Get's input from user
func getUserInput()  (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, email string, firstName string, lastName string)  {
	remainingTickets = remainingTickets - userTickets

	// declare and create empty map
	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["NumberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you for booking %v tickets. You will receive confirmation email at %v\n", userTickets, email)
	fmt.Printf("We have %v tickets remaining for %v\n", remainingTickets, conferenceName)

}
