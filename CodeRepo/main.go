package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50 // const - not allowed to change
	var remainingTickets uint = 50
	var bookings = []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("conferenceTickets is %T, ramainingTickets is %T, conferenceName is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// %v - placeholder coming from Formating package and used with Printf
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var emailAddress string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&emailAddress)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(emailAddress, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break

			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is too short")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of Tickets you entered is invalid")

			}

		}

	}

	// Below is Switch Statement that can can work as if else in above if we have mulitple cities
	//	city := "London"

	//	switch city {
	//	    case "New York":
	//		// execute code for booking New York conference tickets
	//	    case "Singapore", "Hong Kong":
	//		// execute code for booking New Singapore & Hong Kong conference tickets
	//	    case "London", "Berlin":
	//		//Some code here
	//	    default:
	//		    fmt.Print("No valid city selected")

}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)

}
