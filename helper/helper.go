package helper

import "strings"

// Validates user input

func UserInputValidate(firstName string, lastName string, email string, userTickets uint, remainingTickets uint)  (bool,bool,bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets <= remainingTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTicketNumber
}
