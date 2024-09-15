package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []UserData = make([]UserData, 0)

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, email, userTickets)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email) // sendTicket now blocks for 10 seconds. need to split into threads to maintain concurrency.

			firstNames := printFirstNames()
			fmt.Printf("These are all our bookings: %v\n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year.\n\n")
				break
			}
		} else {
			if !isValidTicketNumber {
				fmt.Printf("There are only %v tickets left. You cannot book %v tickets\n\n", remainingTickets, userTickets)
			}
			if !isValidEmail {
				fmt.Printf("Please enter a valid email.\n\n")
			}
			if !isValidName {
				fmt.Printf("Please enter a valid first or last name.\n\n")
			}
			// fmt.Printf("We only have %v tickets remaining so you cannot book %v tickets.\n\n", remainingTickets, userTickets)
			// break // skips the rest of the code and ends the loop
			// continue // skips the rest of code and restarts the loop
		}

	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n\n")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask the user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // use to get user input. use pointers to populate the vars

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you would like: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - uint(userTickets)

	// create a map for a user
	var userData UserData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData) // Slice
	fmt.Printf("Booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. There are %v remaining tickets left.\n\n", firstName, lastName, userTickets, email, remainingTickets)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // blocks execution for 3 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("############################################\n")
	fmt.Printf("Sending ticket:\n%v to email address: %v\n", ticket, email)
	fmt.Printf("############################################\n")
	wg.Done()
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names []string = strings.Fields(booking) // same as string.split("");
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
