package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

var bookings = make([]UserData, 0)

type UserData struct {
	firstName  string
	lastName   string
	email      string
	numTickets uint
}

var wg = sync.WaitGroup{}

// entrypoint
func main() {

	// greet the users
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		// input validations
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// book the application
			bookTicket(userTickets, firstName, lastName, email)

			// adds num of threads the main thread should wait for (in this case, the send ticket subroutine)
			wg.Add(1)
			// execute in seperate thread
			go sendTicket(userTickets, firstName, lastName, email)

			// fmt.Printf("The whole array is %v\n", bookings)
			// fmt.Printf("size of array is %v\n", len(bookings))

			firstNames := printFirstNames()
			fmt.Printf("These are all the current booking first names so far: %v\n", firstNames)
			fmt.Printf("These are all the bookings: %v\n", bookings)
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is fully booked. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name")

			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket amount")

			}
			if !isValidEmail {
				fmt.Println("Invalid email")

			}
			time.Sleep(2 * time.Second)
		}
		// print out type of variables
		// fmt.Printf("conferenceTickets is %T, remainingTikcets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

		// waits for all threads to be done

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets, and", remainingTickets, "are available")
	fmt.Println("Get your tickets here to attend")
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################################################")

	// removes thread added to waiting list for this routine
	wg.Done()

}

func printFirstNames() []string {
	firstNames := []string{}
	// underscores identify unused variables
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name (the & designates a pointer)
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []UserData) {
	remainingTickets = remainingTickets - userTickets

	// create map
	var userData = UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		numTickets: userTickets,
	}

	// add dynamically
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings
}
